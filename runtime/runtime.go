/*
Targets WebAssembly MVP
*/
package main

import (
	"github.com/radkomih/gosemble/constants"
	"github.com/radkomih/gosemble/types"

	"github.com/radkomih/gosemble/utils"
)

/*
	SCALE encoded arguments () allocated in the Wasm VM memory, passed as:
	dataPtr - i32 pointer to the memory location.
	dataLen - i32 length (in bytes) of the encoded arguments.
	returns a pointer-size to the SCALE-encoded (version types.VersionData) data.
*/
//export Core_version
func CoreVersion(dataPtr int32, dataLen int32) int64 {
	scaleEncVersion, err := constants.VersionDataConfig.Encode()
	if err != nil {
		panic(err)
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

// TODO: remove the _start export and find a way to call it from the runtime to initialize the memory.
// TinyGo requires to have a main function to compile to Wasm.
func main() {}
