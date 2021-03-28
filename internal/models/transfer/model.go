package transfer

import (
	"fmt"
	"time"

	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/tokenbalance"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Transfer -
type Transfer struct {
	ID         int64     `json:"-"`
	Network    string    `json:"network"`
	Contract   string    `json:"contract"`
	Initiator  string    `json:"initiator"`
	Hash       string    `json:"hash"`
	Status     string    `json:"status"`
	Timestamp  time.Time `json:"timestamp"`
	Level      int64     `json:"level"`
	From       string    `json:"from"`
	To         string    `json:"to"`
	TokenID    uint64    `json:"token_id" gorm:"type:numeric(50,0)"`
	Amount     float64   `json:"amount,string" gorm:"type:numeric(100,0)"`
	Counter    int64     `json:"counter"`
	Nonce      *int64    `json:"nonce,omitempty"`
	Parent     string    `json:"parent,omitempty"`
	Entrypoint string    `json:"entrypoint,omitempty"`
}

// GetID -
func (t *Transfer) GetID() int64 {
	return t.ID
}

// GetIndex -
func (t *Transfer) GetIndex() string {
	return "transfers"
}

// Save -
func (t *Transfer) Save(tx *gorm.DB) error {
	return tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Save(t).Error
}

// GetQueues -
func (t *Transfer) GetQueues() []string {
	return nil
}

// MarshalToQueue -
func (t *Transfer) MarshalToQueue() ([]byte, error) {
	return nil, nil
}

// LogFields -
func (t *Transfer) LogFields() logrus.Fields {
	return logrus.Fields{
		"network":  t.Network,
		"contract": t.Contract,
		"block":    t.Level,
		"from":     t.From,
		"to":       t.To,
	}
}

// EmptyTransfer -
func EmptyTransfer(o operation.Operation) *Transfer {
	return &Transfer{
		Network:    o.Network,
		Contract:   o.Destination,
		Hash:       o.Hash,
		Status:     o.Status,
		Timestamp:  o.Timestamp,
		Level:      o.Level,
		Initiator:  o.Source,
		Counter:    o.Counter,
		Nonce:      o.Nonce,
		Entrypoint: o.Entrypoint,
	}
}

// GetFromTokenBalanceID -
func (t *Transfer) GetFromTokenBalanceID() string {
	if t.From != "" {
		return fmt.Sprintf("%s_%s_%s_%d", t.Network, t.From, t.Contract, t.TokenID)
	}
	return ""
}

// GetToTokenBalanceID -
func (t *Transfer) GetToTokenBalanceID() string {
	if t.To != "" {
		return fmt.Sprintf("%s_%s_%s_%d", t.Network, t.To, t.Contract, t.TokenID)
	}
	return ""
}

// MakeTokenBalanceUpdate -
func (t *Transfer) MakeTokenBalanceUpdate(from, rollback bool) *tokenbalance.TokenBalance {
	tb := &tokenbalance.TokenBalance{
		Network:  t.Network,
		Contract: t.Contract,
		TokenID:  t.TokenID,
		Balance:  0,
	}
	switch {
	case from && rollback:
		tb.Address = t.From
		tb.Balance = t.Amount
	case !from && rollback:
		tb.Address = t.To
		tb.Balance = -t.Amount
	case from && !rollback:
		tb.Address = t.From
		tb.Balance = -t.Amount
	case !from && !rollback:
		tb.Address = t.To
		tb.Balance = t.Amount
	}
	return tb
}

// TokenSupply -
type TokenSupply struct {
	Supply     float64 `json:"supply"`
	Transfered float64 `json:"transfered"`
}

// Pageable -
type Pageable struct {
	Transfers []Transfer `json:"transfers"`
	Total     int64      `json:"total"`
	LastID    string     `json:"last_id"`
}

// Balance -
type Balance struct {
	Balance float64
	Address string
	TokenID uint64
}
