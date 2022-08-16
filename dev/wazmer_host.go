package dev

// #include <stdlib.h>
//
// extern void ext_allocator_free_version_1(void *context, int32_t a);
// extern int32_t ext_allocator_malloc_version_1(void *context, int32_t a);
// extern void ext_logging_log_version_1(void *context, int32_t a);
//
import "C"

import (
	"fmt"
	"path/filepath"
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

// TODO implement bump allocator

//export ext_allocator_malloc_version_1
func ext_allocator_malloc_version_1(context unsafe.Pointer, size C.int32_t) C.int32_t {
	return C.int32_t(10)
}

//export ext_allocator_free_version_1
func ext_allocator_free_version_1(context unsafe.Pointer, addr C.int32_t) {}

//export ext_logging_log_version_1
func ext_logging_log_version_1(context unsafe.Pointer, size C.int32_t) {}

func RunInWazmer(wasmRuntimeFile string) {
	modulePath, err := filepath.Abs(wasmRuntimeFile)
	check(err)

	bytes, err := wasmer.ReadBytes(modulePath)
	check(err)

	// Compile bytes into wasm binary
	module, err := wasmer.Compile(bytes)
	check(err)

	// Get current wasi version and corresponded import objects
	// wasiVersion := wasmer.WasiGetVersion(module)
	// if wasiVersion == 0 {
	// 	// wasiVersion is unknow, use Latest instead
	// 	wasiVersion = wasmer.Latest
	// }
	// importObject := wasmer.NewDefaultWasiImportObjectForVersion(wasiVersion)

	// Instantiate WebAssembly module using derived import objects.
	importObject := wasmer.NewDefaultWasiImportObject()

	// Allocate memory from the host (the Wasm module expects to import 20 pages)
	memory, err := wasmer.NewMemory(20, 0)
	check(err)

	// Import host provided memory and functions into the Wasm module
	imports := wasmer.NewImports()
	imports.Namespace("env").AppendMemory("memory", memory)
	imports.Namespace("env").AppendFunction("ext_allocator_malloc_version_1", ext_allocator_malloc_version_1, C.ext_allocator_malloc_version_1)
	imports.Namespace("env").AppendFunction("ext_allocator_free_version_1", ext_allocator_free_version_1, C.ext_allocator_free_version_1)
	// imports.Namespace("env").AppendFunction("ext_logging_log_version_1", ext_logging_log_version_1, C.ext_logging_log_version_1)
	importObject.Extend(*imports)

	// Instantiate new module
	instance, err := module.InstantiateWithImportObject(importObject)
	check(err)
	defer importObject.Close()
	defer instance.Close()

	mem := memory.Data()

	// Write some data into the shared memory (from the host)
	data := []byte("HostData")

	dataOffset := int32(1)
	dataSize := int32(len(data))
	for i := int32(0); i < dataSize; i++ {
		mem[i+dataOffset] = data[i]
	}
	fmt.Printf("%s\n", mem[dataOffset:dataSize+dataOffset])

	// Call an exported function from the Wasm module by
	// passing an offset and size to the allocated data
	coreVersion := instance.Exports["Core_version"]
	// TODO check/fix panic with 0 offset
	result, err := coreVersion(dataOffset, dataSize)
	check(err)
	fmt.Println(result)
	// offset, size := utils.Int64ToPointerAndSize(uint64(result.ToI64()))

	// Wasm function overwrites the "HostData" with "RuntimeData"
	fmt.Printf("%s\n", mem[dataOffset:dataSize+dataOffset])
}
