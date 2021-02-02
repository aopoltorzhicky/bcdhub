package transfer

import (
	"math/big"

	"github.com/baking-bad/bcdhub/internal/events"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/tokenbalance"
	"github.com/baking-bad/bcdhub/internal/models/transfer"
)

// DefaultBalanceParser -
type DefaultBalanceParser struct {
	repo models.GeneralRepository
}

// NewDefaultBalanceParser -
func NewDefaultBalanceParser(repo models.GeneralRepository) *DefaultBalanceParser {
	return &DefaultBalanceParser{repo}
}

// Parse -
func (parser *DefaultBalanceParser) Parse(balances []events.TokenBalance, operation operation.Operation) ([]*transfer.Transfer, error) {
	transfers := make([]*transfer.Transfer, 0)
	for _, balance := range balances {
		transfer := transfer.EmptyTransfer(operation)
		if balance.Value.Cmp(big.NewInt(0)) == 1 {
			transfer.To = balance.Address
		} else {
			transfer.From = balance.Address
		}
		transfer.Amount = bigIntToFloat64(balance.Value)
		transfer.TokenID = balance.TokenID

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

// ParseBalances -
func (parser *DefaultBalanceParser) ParseBalances(network, contract string, balances []events.TokenBalance, operation operation.Operation) ([]*transfer.Transfer, error) {
	transfers := make([]*transfer.Transfer, 0)
	for _, balance := range balances {
		transfer := transfer.EmptyTransfer(operation)

		tb := tokenbalance.TokenBalance{
			Network:  network,
			Contract: contract,
			Address:  balance.Address,
			TokenID:  balance.TokenID,
		}
		if err := parser.repo.GetByID(&tb); err != nil {
			if !parser.repo.IsRecordNotFound(err) {
				return nil, err
			}
		}

		delta := big.NewInt(0)
		delta.Sub(balance.Value, tb.Value)

		if delta.Cmp(big.NewInt(0)) == 1 {
			transfer.To = balance.Address
		} else {
			transfer.From = balance.Address
		}

		transfer.Amount = bigIntToFloat64(delta)
		transfer.TokenID = balance.TokenID

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func bigIntToFloat64(x *big.Int) float64 {
	f := new(big.Float).SetInt(x)
	ret, _ := f.Abs(f).Float64()
	return ret
}
