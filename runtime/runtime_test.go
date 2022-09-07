package main

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/runtime/wasmer"
	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/radkomih/gosemble/types"
	"github.com/stretchr/testify/assert"
)

const WASM_RUNTIME = "../build/runtime.wasm"

func Test_Core_version(t *testing.T) {
	rt := wasmer.NewTestInstanceWithTrieLocal(t, WASM_RUNTIME, trie.NewEmptyTrie())
	res, err := rt.Exec("Core_version", []byte{})
	assert.Nil(t, err)

	resultVersion := types.VersionData{}
	err = resultVersion.Decode(res)
	assert.Nil(t, err)

	assert.Equal(t, versionDataConfig, resultVersion)
}

func Test_Core_initialize_block(t *testing.T) {
	rt := wasmer.NewTestInstanceWithTrieLocal(t, WASM_RUNTIME, trie.NewEmptyTrie())

	scaleEncHeader, err := (&types.Header{}).Encode()
	assert.Nil(t, err)

	res, err := rt.Exec("Core_initialize_block", scaleEncHeader)
	assert.Nil(t, err)

	t.Logf("%q", res)
}
