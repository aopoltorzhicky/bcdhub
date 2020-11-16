package search

import (
	"encoding/json"

	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/models"
)

// Contract -
type Contract struct{}

// GetIndex -
func (c Contract) GetIndex() string {
	return "contract"
}

// GetScores -
func (c Contract) GetScores(search string) []string {
	if helpers.IsAddress(search) {
		return []string{
			"address^10",
			"alias^9",
			"tags^9",
			"entrypoints^8",
			"fail_strings^6",
			"language^4",
			"annotations^3",
			"delegate^2",
			"hardcoded^2",
			"manager",
		}
	}
	return []string{
		"alias^10",
		"tags^9",
		"entrypoints^8",
		"fail_strings^6",
		"language^4",
		"annotations^3",
		"delegate^2",
		"hardcoded^2",
		"manager",
		"address",
	}
}

// GetFields -
func (c Contract) GetFields() []string {
	return []string{
		"address",
		"alias",
		"tags",
		"entrypoints",
		"fail_strings",
		"language",
		"annotations",
		"delegate",
		"hardcoded",
		"manager",
	}
}

// Parse  -
func (c Contract) Parse(highlight map[string][]string, data []byte) (interface{}, error) {
	var contract models.Contract
	if err := json.Unmarshal(data, &contract); err != nil {
		return nil, err
	}
	return Item{
		Type:       c.GetIndex(),
		Value:      contract.Address,
		Body:       contract,
		Highlights: highlight,
		Network:    contract.Network,
	}, nil
}
