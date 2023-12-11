// (c) 2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"testing"

	"github.com/luxdefi/coreth/core/rawdb"
	"github.com/luxdefi/coreth/precompile/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func NewTestStateDB(t testing.TB) contract.StateDB {
	db := rawdb.NewMemoryDatabase()
	stateDB, err := New(common.Hash{}, NewDatabase(db), nil)
	require.NoError(t, err)
	return stateDB
}
