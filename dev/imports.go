package dev

// #include <stdlib.h>
//
// extern int32_t ext_allocator_malloc_version_1(void *context, int32_t a);
// extern void ext_allocator_free_version_1(void *context, int32_t a);
//
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export ext_allocator_malloc_version_1
func ext_allocator_malloc_version_1(context unsafe.Pointer, size C.int32_t) C.int32_t {
	// fmt.Printf("executing malloc with size %d...", int64(size))

	instanceContext := wasmer.IntoInstanceContext(context)
	ctx := instanceContext.Data().(*Context)

	// Allocate memory
	res, err := ctx.Allocator.Allocate(uint32(size))
	if err != nil {
		fmt.Errorf("failed to allocate memory: %s", err)
		panic(err)
	}

	return C.int32_t(res)
}

//export ext_allocator_free_version_1
func ext_allocator_free_version_1(context unsafe.Pointer, addr C.int32_t) {
	// fmt.Printf("executing free...")

	instanceContext := wasmer.IntoInstanceContext(context)
	runtimeCtx := instanceContext.Data().(*Context)

	// Deallocate memory
	err := runtimeCtx.Allocator.Deallocate(uint32(addr))

	if err != nil {
		fmt.Errorf("failed to free memory: %s", err)
	}
}

// importsNodeRuntime returns the WASM imports for the node runtime.
func importsNodeRuntime() (imports *wasmer.Imports, err error) {
	imports = wasmer.NewImports()
	// Note imports are closed by the call to wasm.Instance.Close()

	for _, toRegister := range []struct {
		importName     string
		implementation interface{}
		cgoPointer     unsafe.Pointer
	}{
		{"ext_allocator_free_version_1", ext_allocator_free_version_1, C.ext_allocator_free_version_1},
		{"ext_allocator_malloc_version_1", ext_allocator_malloc_version_1, C.ext_allocator_malloc_version_1},
	} {
		_, err = imports.AppendFunction(toRegister.importName, toRegister.implementation, toRegister.cgoPointer)
		if err != nil {
			return nil, fmt.Errorf("importing function: %w", err)
		}
	}

	return imports, nil
}
