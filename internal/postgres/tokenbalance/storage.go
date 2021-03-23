package tokenbalance

import (
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/tokenbalance"
	"github.com/baking-bad/bcdhub/internal/postgres/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Storage -
type Storage struct {
	*core.Postgres
}

// NewStorage -
func NewStorage(pg *core.Postgres) *Storage {
	return &Storage{pg}
}

// Update -
func (storage *Storage) Update(updates []*tokenbalance.TokenBalance) error {
	if len(updates) == 0 {
		return nil
	}

	return storage.DB.Transaction(func(tx *gorm.DB) error {
		for i := range updates {
			if err := tx.Table(models.DocTokenBalances).Clauses(clause.OnConflict{
				Columns: []clause.Column{
					{Name: "network"},
					{Name: "contract"},
					{Name: "address"},
					{Name: "token_id"},
				},
				DoUpdates: clause.AssignmentColumns([]string{"balance"}),
			}).Create(&updates[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetHolders -
func (storage *Storage) GetHolders(network, contract string, tokenID uint64) ([]tokenbalance.TokenBalance, error) {
	var balances []tokenbalance.TokenBalance
	err := storage.DB.Table(models.DocTokenBalances).
		Scopes(core.Token(network, contract, tokenID)).
		Where("balance != '0'").
		Find(&balances).Error
	return balances, err
}

// GetAccountBalances -
func (storage *Storage) GetAccountBalances(network, address, contract string, size, offset int64) ([]tokenbalance.TokenBalance, int64, error) {
	var balances []tokenbalance.TokenBalance

	query := storage.DB.Table(models.DocTokenBalances).Scopes(core.NetworkAndAddress(network, address))

	if contract != "" {
		query.Where("contract = ?", contract)
	}

	limit := core.GetPageSize(size)
	if err := query.
		Limit(limit).
		Offset(int(offset)).
		Find(&balances).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return balances, count, nil
}

// NFTHolders -
func (storage *Storage) NFTHolders(network, contract string, tokenID uint64) (tokens []tokenbalance.TokenBalance, err error) {
	err = storage.DB.
		Scopes(core.Token(network, contract, tokenID)).
		Where("balance != '0'").
		Find(&tokens).Error
	return
}

// Batch -
func (storage *Storage) Batch(network string, addresses []string) (map[string][]tokenbalance.TokenBalance, error) {
	var balances []tokenbalance.TokenBalance

	query := storage.DB.Table(models.DocTokenBalances).Scopes(core.Network(network)).Where("balance != '0'")

	for i := range addresses {
		if i == 0 {
			query.Where("address = ?", addresses[i])
		} else {
			query.Or("address = ?", addresses[i])
		}
	}

	if err := query.Find(&balances).Error; err != nil {
		return nil, err
	}

	result := make(map[string][]tokenbalance.TokenBalance)

	for _, b := range balances {
		if _, ok := result[b.Address]; !ok {
			result[b.Address] = make([]tokenbalance.TokenBalance, 0)
		}
		result[b.Address] = append(result[b.Address], b)
	}

	return result, nil
}

type tokensByContract struct {
	Contract    string
	TokensCount int64
}

// CountByContract -
func (storage *Storage) CountByContract(network, address string) (map[string]int64, error) {
	var resp []tokensByContract
	query := storage.DB.Table(models.DocTokenBalances).
		Select("contract, count(*) as tokens_count").
		Scopes(core.NetworkAndAddress(network, address)).
		Group("contract").
		Scan(&resp)

	if query.Error != nil {
		return nil, query.Error
	}

	result := make(map[string]int64)
	for i := range resp {
		result[resp[i].Contract] = resp[i].TokensCount
	}
	return result, nil
}
