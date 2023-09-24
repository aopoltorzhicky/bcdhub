package protocol

import (
	"github.com/uptrace/bun"
)

// Protocol -
type Protocol struct {
	bun.BaseModel `bun:"protocols"`

	ID int64 `bun:"id,pk,notnull,autoincrement"`

	Hash       string `bun:"hash,unique:protocol_hash_idx"`
	StartLevel int64
	EndLevel   int64
	SymLink    string
	Alias      string
	ChainID    string
	*Constants
}

// Constants -
type Constants struct {
	CostPerByte                  int64
	HardGasLimitPerOperation     int64
	HardStorageLimitPerOperation int64
	TimeBetweenBlocks            int64
}

// GetID -
func (p *Protocol) GetID() int64 {
	return p.ID
}

// GetIndex -
func (p *Protocol) GetIndex() string {
	return "protocols"
}

// ValidateChainID -
func (p *Protocol) ValidateChainID(chainID string) bool {
	if p.ChainID == "" {
		return p.StartLevel == 0
	}
	return p.ChainID == chainID
}

func (Protocol) PartitionBy() string {
	return ""
}
