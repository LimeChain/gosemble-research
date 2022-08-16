/*
	Targets WebAssembly MVP
*/
package main

import "github.com/radkomih/gosemble/utils"

const SPEC_NAME = "gosemble"
const IMPL_NAME = "Go"
const AUTHORING_VERSION = 1
const SPEC_VERSION = 1
const IMPL_VERSION = 1
const TRANSACTION_VERSION = 1
const STATE_VERSION = 1

// var global *int8
func newInt8Ref() *int8 {
	n := int8(127)
	return &n
}

/*
	SCALE encoded arguments () allocated in the Wasm VM memory, passed as:
	dataPtr - i32 pointer to the memory location.
	dataLen - i32 length (in bytes) of the encoded arguments.
	returns a pointer-size to the SCALE-encoded (version types.VersionData) data.
*/
//export Core_version
func CoreVersion(dataPtr int32, dataLen int32) int64 {
	data := [11]byte{'R', 'u', 'n', 't', 'i', 'm', 'e', 'D', 'a', 't', 'a'}
	utils.WriteMemory(dataPtr, dataLen, data)

	n := newInt8Ref()
	*n += int8(127)

	return utils.PointerAndSizeToInt64(dataPtr, int32(dataLen))

	// TODO fix/add support of the reflect package
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
	// scaleEncVersion, _ := version.Encode()
	// return utils.BytesToPointerAndSize(scaleEncVersion)
}

// TinyGo requires to have a main function to compile to Wasm.
func main() {}
