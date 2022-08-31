package main

import (
	"fmt"

	"github.com/radkomih/gosemble/dev"
)

func main() {
	// dev.Run("build/dev_runtime.wasm")

	wasmBytes, _ := dev.ReadBytes("build/dev_runtime.wasm")

	in, err := dev.NewInstance(wasmBytes, dev.Config{})
	dev.Check(err)

	// HostData -> RuntimeD

	res, err := in.Exec("Core_version", []byte{})
	dev.Check(err)
	fmt.Printf("%q\n", res)
}
