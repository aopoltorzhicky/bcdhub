package index

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/tzkt"
)

// TzKT -
type TzKT struct {
	api *tzkt.TzKT

	lastContractsPage int64
}

// NewTzKT -
func NewTzKT(host string, timeout time.Duration) *TzKT {
	return &TzKT{
		api: tzkt.NewTzKT(host, timeout),
	}
}

// GetHead -
func (t *TzKT) GetHead() (Head, error) {
	resp := Head{}
	head, err := t.api.GetHead()
	if err != nil {
		return resp, err
	}
	resp.Level = head.Level
	resp.Hash = head.Hash
	resp.Timestamp = head.Timestamp.UTC()
	return resp, err
}

// GetContracts -
func (t *TzKT) GetContracts(startLevel int64) ([]Contract, error) {
	resp := make([]Contract, 0)

	end := false
	for !end {
		contracts, err := t.api.GetAccounts(tzkt.ContractKindSmart, t.lastContractsPage, 1000)
		if err != nil {
			return nil, err
		}
		for _, contract := range contracts {
			if contract.FirstActivity <= startLevel {
				continue
			}

			data := Contract{
				Address:   contract.Address,
				Level:     contract.FirstActivity,
				Timestamp: contract.FirstActivityTime.UTC(),
				Balance:   contract.Balance,
			}
			if contract.Creator != nil {
				data.Manager = contract.Creator.Address
			}
			if contract.Delegate != nil {
				data.Delegate = contract.Delegate.Address
			}
			resp = append(resp, data)

		}
		if len(contracts) == 1000 {
			t.lastContractsPage++
		} else {
			end = true
		}
	}

	return resp, nil
}

// GetContractOperationBlocks -
func (t *TzKT) GetContractOperationBlocks(startBlock, endBlock int64) ([]int64, error) {
	start := int64(startBlock)
	end := false

	result := make([]int64, 0)
	for !end {
		blocks, err := t.api.GetContractOperationBlocks(start, 10000)
		if err != nil {
			return nil, err
		}

		if len(blocks) == 0 {
			end = true
			continue
		}

		for i := range blocks {
			if blocks[i] <= int64(endBlock) {
				result = append(result, blocks[i])
			} else {
				return result, nil
			}
		}

		start = blocks[len(blocks)-1]
	}

	return result, nil
}

// GetProtocols -
func (t *TzKT) GetProtocols() ([]Protocol, error) {
	protocols, err := t.api.GetProtocols()
	if err != nil {
		return nil, err
	}
	res := make([]Protocol, len(protocols))
	for i := range protocols {
		res[i] = Protocol{
			Hash:       protocols[i].Hash,
			StartLevel: protocols[i].StartLevel,
			LastLevel:  protocols[i].LastLevel,
			Alias:      protocols[i].Metadata.Alias,
		}
	}
	return res, nil
}
