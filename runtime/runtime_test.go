package main

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/radkomih/gosemble/helpers"
	"github.com/stretchr/testify/require"
)

func Test_Core_version(t *testing.T) {
	tt := trie.NewEmptyTrie()
	rt := helpers.NewTestInstanceWithTrie(t, tt)

	res, err := rt.Exec("Core_version", []byte{'W', 'a', 's', 'm'})
	t.Log(res)
	t.Log(tt)

	require.NoError(t, err)
}
