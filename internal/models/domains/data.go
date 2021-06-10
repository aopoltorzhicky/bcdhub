package domains

import (
	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/baking-bad/bcdhub/internal/models/transfer"
)

// TokenBalance -
type TokenBalance struct {
	tokenmetadata.TokenMetadata

	Balance string
}

// TokenBalanceResponse -
type TokenBalanceResponse struct {
	Balances []TokenBalance
	Count    int64
}

// Transfer -
type Transfer struct {
	*transfer.Transfer
	Hash     string
	Symbol   string
	Name     string
	Counter  int64
	Nonce    *int64
	Decimals *int64
}

// TransfersResponse -
type TransfersResponse struct {
	Total     int64
	LastID    string
	Transfers []Transfer
}
