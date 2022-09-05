package main

import (
	"fmt"

	"github.com/radkomih/gosemble/dev"
)

func main() {
	wasmBytes, _ := dev.ReadBytes("build/dev_runtime.wasm")

	in, err := dev.NewInstance(wasmBytes, dev.Config{})
	dev.Check(err)

	res, err := in.Exec("Core_version", []byte{'H', 'o', 's', 't', 'D', 'a', 't', 'a'})
	dev.Check(err)

	fmt.Printf("\n%q\n", res)
}
