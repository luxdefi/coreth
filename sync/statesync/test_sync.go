// (c) 2021-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package statesync

import (
	"bytes"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/coreth/core/rawdb"
	"github.com/ava-labs/coreth/core/state/snapshot"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

// assertDBConsistency checks [serverTrieDB] and [clientTrieDB] have the same EVM state trie at [root],
// and that [clientTrieDB.DiskDB] has corresponding account & snapshot values.
// Also verifies any code referenced by [clientTrieDB] is present the hash is correct.
// TODO ensure snapshot does not contain any extra data
func assertDBConsistency(t testing.TB, root common.Hash, serverTrieDB, clientTrieDB *trie.Database) {
	trie.AssertTrieConsistency(t, root, serverTrieDB, clientTrieDB, func(key, val []byte) error {
		accHash := common.BytesToHash(key)
		var acc types.StateAccount
		if err := rlp.DecodeBytes(val, &acc); err != nil {
			return err
		}
		// check snapshot consistency
		snapshotVal := rawdb.ReadAccountSnapshot(clientTrieDB.DiskDB(), accHash)
		expectedSnapshotVal := snapshot.SlimAccountRLP(acc.Nonce, acc.Balance, acc.Root, acc.CodeHash, acc.IsMultiCoin)
		assert.Equal(t, expectedSnapshotVal, snapshotVal)

		// check code consistency
		if !bytes.Equal(acc.CodeHash, types.EmptyCodeHash[:]) {
			codeHash := common.BytesToHash(acc.CodeHash)
			code := rawdb.ReadCode(clientTrieDB.DiskDB(), codeHash)
			actualHash := crypto.Keccak256Hash(code)
			assert.NotZero(t, len(code))
			assert.Equal(t, codeHash, actualHash)
		}
		if acc.Root == types.EmptyRootHash {
			return nil
		}
		// check storage trie and storage snapshot consistency
		trie.AssertTrieConsistency(t, acc.Root, serverTrieDB, clientTrieDB, func(key, val []byte) error {
			snapshotVal := rawdb.ReadStorageSnapshot(clientTrieDB.DiskDB(), accHash, common.BytesToHash(key))
			assert.Equal(t, val, snapshotVal)
			return nil
		})
		return nil
	})
}

func fillAccountsWithStorage(t *testing.T, serverTrieDB *trie.Database, root common.Hash, numAccounts int64) common.Hash {
	return fillAccounts(t, serverTrieDB, root, numAccounts, func(t *testing.T, index int64, account types.StateAccount, tr *trie.Trie) types.StateAccount {
		// Add code and storage for every third account
		if index%3 == 0 {
			codeBytes := make([]byte, 256)
			_, err := rand.Read(codeBytes)
			if err != nil {
				t.Fatalf("error reading random code bytes: %v", err)
			}

			codeHash := crypto.Keccak256Hash(codeBytes)
			rawdb.WriteCode(serverTrieDB.DiskDB(), codeHash, codeBytes)
			account.CodeHash = codeHash[:]

			// now create state trie
			numKeys := 16
			account.Root, _, _ = trie.GenerateTrie(t, serverTrieDB, numKeys, wrappers.LongLen+1)
		}
		return account
	})
}

func fillAccounts(
	t *testing.T, trieDB *trie.Database, root common.Hash, numAccounts int64,
	onAccount func(*testing.T, int64, types.StateAccount, *trie.Trie) types.StateAccount,
) common.Hash {
	tr, err := trie.New(root, trieDB)
	if err != nil {
		t.Fatalf("error opening trie: %v", err)
	}
	for i := int64(0); i < numAccounts; i++ {
		acc := types.StateAccount{
			Nonce:    uint64(i),
			Balance:  big.NewInt(i % 1337),
			CodeHash: types.EmptyCodeHash[:],
			Root:     types.EmptyRootHash,
		}

		if i%5 == 0 {
			acc.Nonce += rand.Uint64()
		}

		if onAccount != nil {
			acc = onAccount(t, i, acc, tr)
		}

		accBytes, err := rlp.EncodeToBytes(acc)
		if err != nil {
			t.Fatalf("failed to rlp encode account: %v", err)
		}

		accHash := make([]byte, common.HashLength)
		if _, err := rand.Read(accHash); err != nil {
			t.Fatalf("error reading random bytes: %v", err)
		}
		if err = tr.TryUpdate(accHash, accBytes); err != nil {
			t.Fatalf("error updating trie with account, hash=%s, err=%v", common.BytesToHash(accHash), err)
		}
	}
	newRoot, _, err := tr.Commit(nil)
	if err != nil {
		t.Fatalf("error committing trie: %v", err)
	}
	if err := trieDB.Commit(newRoot, false, nil); err != nil {
		t.Fatalf("error committing trieDB: %v", err)
	}
	return newRoot
}
