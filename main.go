package main

import (
	"fmt"

	"github.com/radkomih/gosemble/dev"
)

func main() {
	wasmBytes, _ := dev.ReadBytes("build/dev_runtime.wasm")

	in, err := dev.NewInstance(wasmBytes, dev.Config{})
	dev.Check(err)

	// res, err := in.Exec("Core_version", []byte{})
	res, err := in.Exec("Core_initialize_block", []byte{'G', 'o'})
	dev.Check(err)

	fmt.Printf("%q\n", res)
}
