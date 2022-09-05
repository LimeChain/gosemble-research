/*
Targets WebAssembly MVP
*/
package main

import (
	"github.com/radkomih/gosemble/types"

	"github.com/radkomih/gosemble/utils"
)

const SPEC_NAME = "gosemble"
const IMPL_NAME = "go"
const AUTHORING_VERSION = 1
const SPEC_VERSION = 2
const IMPL_VERSION = 3
const TRANSACTION_VERSION = 4
const STATE_VERSION = 5

var versionDataConfig = types.VersionData{
	SpecName:         []byte(SPEC_NAME),
	ImplName:         []byte(IMPL_NAME),
	AuthoringVersion: uint32(AUTHORING_VERSION),
	SpecVersion:      uint32(SPEC_VERSION),
	ImplVersion:      uint32(IMPL_VERSION),
	Apis: []types.ApiItem{
		{Name: [8]byte{1, 1, 1, 1, 1, 1, 1, 1}, Version: 0},
	},
	TransactionVersion: uint32(TRANSACTION_VERSION),
	StateVersion:       uint32(STATE_VERSION),
}

//go:wasm-module env
//export ext_storage_set_version_1
func ext_storage_set_version_1(key int64, value int64)

/*
	SCALE encoded arguments () allocated in the Wasm VM memory, passed as:
	dataPtr - i32 pointer to the memory location.
	dataLen - i32 length (in bytes) of the encoded arguments.
	returns a pointer-size to the SCALE-encoded (version types.VersionData) data.
*/
//export Core_version
func CoreVersion(dataPtr int32, dataLen int32) int64 {
	scaleEncVersion, err := versionDataConfig.Encode()
	if err != nil {
		// TODO: handle or log
	}
	// TODO: retain the pointer to the scaleEncVersion
	return utils.BytesToOffsetAndSize(scaleEncVersion)
}

/*
	SCALE encoded arguments (block types.Block) allocated in the Wasm VM memory, passed as:
	dataPtr - i32 pointer to the memory location.
	dataLen - i32 length (in bytes) of the encoded arguments.
*/
//export Core_execute_block
func ExecuteBlock(dataPtr int32, dataLen int32) {

}

/*
SCALE encoded arguments (header *types.Header) allocated in the Wasm VM memory, passed as:
dataPtr - i32 pointer to the memory location.
dataLen - i32 length (in bytes) of the encoded arguments.
*/
//export Core_initialize_block
func CoreInitializeBlock(dataPtr int32, dataLen int32) {
	input := utils.ToWasmMemorySlice(dataPtr, dataLen)
	header := (&types.Header{}).Decode(input)
	_ = header
	ext_storage_set_version_1(int64(123), int64(456))
}

// TODO:
// Remove the _start export and find a way to call it from the runtime to initialize the memory
// TinyGo requires to have a main function to compile to Wasm.
func main() {}
