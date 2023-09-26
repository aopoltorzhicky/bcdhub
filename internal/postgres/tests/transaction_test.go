package tests

import (
	"context"
	"time"

	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/migration"
	"github.com/baking-bad/bcdhub/internal/models/protocol"
	"github.com/baking-bad/bcdhub/internal/models/stats"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/postgres/core"
)

func (s *StorageTestSuite) TestSave() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	account := account.Account{
		Address: "address",
		Type:    types.AccountTypeContract,
		Level:   100,
	}
	err = tx.Save(ctx, &account)
	s.Require().NoError(err)
	s.Require().Positive(account.ID)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestMigrations() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	m := migration.Migration{
		ProtocolID:     1,
		PrevProtocolID: 0,
		Hash:           []byte{0, 1, 2, 3, 4},
		Timestamp:      time.Now(),
		Level:          100,
		Kind:           types.MigrationKindBootstrap,
		ContractID:     1,
	}
	err = tx.Migrations(ctx, &m)
	s.Require().NoError(err)
	s.Require().Positive(m.ID)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestProtocol() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	p := protocol.Protocol{
		Hash:       "protocol_hash",
		StartLevel: 100,
		EndLevel:   200,
		SymLink:    "symlink",
		Alias:      "alias",
		ChainID:    "chain_id",
		Constants:  &protocol.Constants{},
	}
	err = tx.Protocol(ctx, &p)
	s.Require().NoError(err)
	s.Require().Positive(p.ID)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestScriptConstants() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	sc := []*contract.ScriptConstants{
		{
			ScriptId:         1,
			GlobalConstantId: 1,
		}, {
			ScriptId:         2,
			GlobalConstantId: 1,
		}, {
			ScriptId:         1,
			GlobalConstantId: 2,
		},
	}
	err = tx.ScriptConstant(ctx, sc...)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestScripts() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	sc := []*contract.Script{
		{
			Hash: "hash_1",
		}, {
			Hash: "hash_2",
		}, {
			Hash: "hash_3",
		},
	}
	err = tx.Scripts(ctx, sc...)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestScriptsConflict() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	update := contract.Script{
		Hash: "8436dde35bd56644cd4f40c5f26839cb8f4b51052e415da2b9fadcd9bddcb03e",
	}
	err = tx.Scripts(ctx, &update)
	s.Require().NoError(err)
	s.Require().EqualValues(5, update.ID)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestAccounts() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	sc := []*account.Account{
		{
			Address: "address_1",
			Type:    types.AccountTypeContract,
			Level:   100,
		}, {
			Address: "address_12",
			Type:    types.AccountTypeSmartRollup,
			Level:   100,
		}, {
			Address: "address_2",
			Type:    types.AccountTypeTz,
			Level:   100,
		},
	}
	err = tx.Accounts(ctx, sc...)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)
}

func (s *StorageTestSuite) TestBigMapStates() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	sc := []*bigmapdiff.BigMapState{
		{
			Key:             []byte{0, 1, 2, 3},
			KeyHash:         "hash 1",
			Ptr:             100000,
			LastUpdateLevel: 100,
			Count:           1,
			Removed:         false,
			Contract:        "contract 1",
		}, {
			Key:             []byte{0, 1, 2, 3, 4},
			KeyHash:         "hash 2",
			Ptr:             100000,
			LastUpdateLevel: 100,
			Count:           1,
			Removed:         false,
			Contract:        "contract 2",
		}, {
			Key:             []byte{0, 1, 2, 3, 5},
			KeyHash:         "hash 3",
			Ptr:             100000,
			LastUpdateLevel: 100,
			Count:           1,
			Removed:         false,
			Contract:        "contract 3"},
	}
	err = tx.BigMapStates(ctx, sc...)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var result []bigmapdiff.BigMapState
	err = s.storage.DB.NewSelect().Model(&result).Where("ptr = 100000").Scan(ctx)
	s.Require().NoError(err)
	s.Require().Len(result, 3)
}

func (s *StorageTestSuite) TestBabylonUpdateNonDelegator() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	c := contract.Contract{
		ID:        2,
		BabylonID: 10,
	}

	err = tx.BabylonUpdateNonDelegator(ctx, &c)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var newContract contract.Contract
	err = s.storage.DB.NewSelect().Model(&newContract).Where("id = 2").Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(10, newContract.BabylonID)
}

func (s *StorageTestSuite) TestJakartaVesting() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	c := contract.Contract{
		ID: 2,
	}

	err = tx.JakartaVesting(ctx, &c)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var newContract contract.Contract
	err = s.storage.DB.NewSelect().Model(&newContract).Where("id = 2").Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(5, newContract.JakartaID)
}

func (s *StorageTestSuite) TestJakartaUpdateNonDelegator() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	c := contract.Contract{
		ID:        2,
		JakartaID: 100,
	}

	err = tx.JakartaUpdateNonDelegator(ctx, &c)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var newContract contract.Contract
	err = s.storage.DB.NewSelect().Model(&newContract).Where("id = 2").Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(100, newContract.JakartaID)
}

func (s *StorageTestSuite) TestToJakarta() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	err = tx.ToJakarta(ctx)
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var newContract contract.Contract
	err = s.storage.DB.NewSelect().Model(&newContract).Where("id = 16").Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(14, newContract.JakartaID)
}

func (s *StorageTestSuite) TestBabylonBigMapStates() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	err = tx.BabylonBigMapStates(ctx, &bigmapdiff.BigMapState{
		ID:       3,
		Ptr:      1000,
		KeyHash:  "expruDuAZnFKqmLoisJqUGqrNzXTvw7PJM2rYk97JErM5FHCerQqgn",
		Contract: "KT1Pz65ssbPF7Zv9Dh7ggqUkgAYNSuJ9iia7",
	})
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var state bigmapdiff.BigMapState
	err = s.storage.DB.NewSelect().Model(&state).Where("id = 3").Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(1000, state.Ptr)
}

func (s *StorageTestSuite) TestUpdateStats() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := core.NewTransaction(ctx, s.storage.DB)
	s.Require().NoError(err)

	err = tx.UpdateStats(ctx, stats.Stats{
		ID:                  1,
		ContractsCount:      1,
		OperationsCount:     4,
		OriginationsCount:   1,
		TransactionsCount:   1,
		EventsCount:         1,
		SrOriginationsCount: 1,
	})
	s.Require().NoError(err)

	err = tx.Commit()
	s.Require().NoError(err)

	var stats stats.Stats
	err = s.storage.DB.NewSelect().Model(&stats).Limit(1).Scan(ctx)
	s.Require().NoError(err)
	s.Require().EqualValues(121, stats.ContractsCount)
	s.Require().EqualValues(196, stats.OperationsCount)
	s.Require().EqualValues(73, stats.TransactionsCount)
	s.Require().EqualValues(119, stats.OriginationsCount)
	s.Require().EqualValues(3, stats.EventsCount)
	s.Require().EqualValues(1, stats.SrOriginationsCount)
}