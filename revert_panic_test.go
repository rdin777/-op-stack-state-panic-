package state

import (
        "testing"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/rawdb"
        "github.com/ethereum/go-ethereum/triedb"
        "github.com/holiman/uint256"
)

func TestOpStackStatePanic(t *testing.T) {
        // 1. Создаем базу в памяти
        rawDB := rawdb.NewMemoryDatabase()

        // 2. Инициализируем TrieDB (это то, что требовал стэк-трейс)
        tdb := triedb.NewDatabase(rawDB, nil)

        // 3. Теперь создаем Database для StateDB
        db := NewDatabase(tdb, nil)
        sdb, _ := New(common.Hash{}, db)

        addr := common.HexToAddress("0xdeadbeef")
        val := uint256.NewInt(100)

        for i := 0; i < 1000; i++ {
                snap := sdb.Snapshot()

                sdb.AddBalance(addr, val, 0)
                sdb.SetNonce(addr, uint64(i), 0)

                sdb.CreateAccount(addr)
                sdb.SelfDestruct(addr)

                // Вот здесь мы ищем настоящий баг
                sdb.RevertToSnapshot(snap)

                if sdb.GetNonce(addr) != 0 {
                        t.Errorf("Iteration %d: Nonce not reverted!", i)
                }
        }
}

