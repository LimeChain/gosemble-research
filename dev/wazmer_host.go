package dev

// #include <stdlib.h>
//
// extern int32_t ext_allocator_malloc_version_1(void *context, int32_t a);
// extern void ext_allocator_free_version_1(void *context, int32_t a);
//
import "C"

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func Run(wasmRuntimeFile string) {
	wasmBytes, err := ReadBytes(wasmRuntimeFile)
	Check(err)

	if len(wasmBytes) == 0 {
		Check(errors.New("code is empty"))
	}

	// Compile bytes into Wasm binary
	module, err := wasmer.Compile(wasmBytes)
	Check(err)

	// Import host provided memory and functions into the Wasm module
	imports := wasmer.NewImports()

	// Host provided functions
	for _, toRegister := range []struct {
		importName     string
		implementation interface{}
		cgoPointer     unsafe.Pointer
	}{
		{"ext_allocator_malloc_version_1", ext_allocator_malloc_version_1, C.ext_allocator_malloc_version_1},
		{"ext_allocator_free_version_1", ext_allocator_free_version_1, C.ext_allocator_free_version_1},
	} {
		_, err = imports.Namespace("env").AppendFunction(toRegister.importName, toRegister.implementation, toRegister.cgoPointer)
		Check(err)
	}

	// Host provided memory (the Wasm module expects to import 20 pages)
	memory, err := wasmer.NewMemory(20, 0)
	Check(err)
	imports.Namespace("env").AppendMemory("memory", memory)

	// Instantiate new WebAssembly module using derived import objects.
	importObject := wasmer.NewDefaultWasiImportObject()
	importObject.Extend(*imports)
	defer importObject.Close()
	instance, err := module.InstantiateWithImportObject(importObject) // instance, err := wasmer.NewInstanceWithImports(wasmBytes, imports)
	Check(err)
	defer instance.Close()

	if !instance.HasMemory() {
		instance.Memory = memory
	}

	allocator := NewAllocator(instance.Memory, DefaultHeapBase)
	_ = allocator

	mem := memory.Data()

	// Write some data into the shared memory (from the host)
	data := []byte("HostData")

	// TODO fix 0 offset
	dataOffset := int32(1)
	dataSize := int32(len(data))
	copy(mem[dataOffset:dataOffset+int32(len(data))], data)

	fmt.Printf("%s\n", mem[dataOffset:dataOffset+dataSize])

	// Call an exported function from the Wasm module by
	// passing an offset and size to the allocated data
	coreVersion := instance.Exports["Core_version"]
	// TODO check/fix panic with 0 offset
	result, err := coreVersion(dataOffset, dataSize)
	Check(err)

	fmt.Println(result)

	// Wasm function overwrites the "HostData" with "RuntimeData"
	// offset, size := utils.Int64ToOffsetAndSize(uint64(result.ToI64()))
	fmt.Printf("%s\n", mem[dataOffset:dataOffset+dataSize])
}
