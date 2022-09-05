package main

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/radkomih/gosemble/helpers"
	"github.com/radkomih/gosemble/types"
	"github.com/stretchr/testify/assert"
)

func Test_Core_version(t *testing.T) {
	// TODO add host provided storage functions

	tt := trie.NewEmptyTrie()
	rt := helpers.NewTestInstanceWithTrie(t, tt)

	res, err := rt.Exec("Core_version", []byte{})
	t.Log(res)
	t.Log(tt)

	assert.Nil(t, err)

	resultVersion := types.VersionData{}
	err = resultVersion.Decode(res)

	assert.Nil(t, err)
	assert.Equal(t, versionDataConfig, resultVersion)
}
