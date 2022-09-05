/*
Targets WebAssembly MVP
*/
package main

import (
	"bytes"

	"github.com/radkomih/gosemble/scale"
	"github.com/radkomih/gosemble/utils"
)

const BYTES_SIZE = 1

func newBytesOnHeap() *[BYTES_SIZE]byte {
	bs := [BYTES_SIZE]byte{}
	for i := 0; i < BYTES_SIZE; i++ {
		bs[i] = '-'
	}
	return &bs
}

const SPEC_NAME = "gosemble"
const IMPL_NAME = "Go"
const AUTHORING_VERSION = 1
const SPEC_VERSION = 1
const IMPL_VERSION = 1
const TRANSACTION_VERSION = 1
const STATE_VERSION = 1

/*
	SCALE encoded arguments () allocated in the Wasm VM memory, passed as:
	dataPtr - i32 pointer to the memory location.
	dataLen - i32 length (in bytes) of the encoded arguments.
	returns a pointer-size to the SCALE-encoded (version types.VersionData) data.
*/
//export Core_version
func CoreVersion(dataPtr int32, dataLen int32) int64 {
	// version := &types.VersionData{
	// 	SpecName:         []byte(SPEC_NAME),
	// 	ImplName:         []byte(IMPL_NAME),
	// 	AuthoringVersion: uint32(AUTHORING_VERSION),
	// 	SpecVersion:      uint32(SPEC_VERSION),
	// 	ImplVersion:      uint32(IMPL_VERSION),
	// 	Apis: []types.ApiItem{
	// 		{Name: [8]byte{1, 1, 1, 1, 1, 1, 1, 1}, Version: 1},
	// 	},
	// 	TransactionVersion: uint32(TRANSACTION_VERSION),
	// 	StateVersion:       uint32(STATE_VERSION),
	// }

	// scaleEncVersion, err := version.Encode()
	// if err != nil {
	// 	// TODO: handle or log
	// }

	// // TODO: retain the pointer to the scaleEncVersion
	// return utils.BytesToOffsetAndSize(scaleEncVersion)

	var buffer = bytes.Buffer{}
	var encoder = scale.Encoder{Writer: &buffer}
	encoder.EncodeBool(true)
	buf := buffer.Bytes()
	return utils.BytesToOffsetAndSize(buf)

	// // allocated on the heap from host's allocator
	// data := [256]byte{'W', 'a', 's', 'm', 'D', 'a', 't', 'a'}
	// _ = data
	// utils.WriteToMemory(dataPtr, dataLen, data)

	// // allocated on the heap from runtime's allocator
	// for i := 0; i < 2*64*1024; i++ {
	// 	_ = newBytesOnHeap()
	// }

	// return utils.OffsetAndSizeToInt64(dataPtr, int32(dataLen))
}

// TODO:
// Remove the _start export and find a way to call it from the runtime to initialize the memory
// TinyGo requires to have a main function to compile to Wasm.
func main() {}
