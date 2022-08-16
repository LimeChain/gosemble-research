package dev

import (
	"path/filepath"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadBytes(wasmRuntimeFile string) (wasmBytes []byte, err error) {
	modulePath, err := filepath.Abs(wasmRuntimeFile)
	Check(err)

	return wasmer.ReadBytes(modulePath)
}
