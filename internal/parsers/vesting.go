package parsers

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/contractparser/consts"
	"github.com/baking-bad/bcdhub/internal/contractparser/kinds"
	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/balanceupdate"
	"github.com/baking-bad/bcdhub/internal/models/migration"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/parsers/contract"
	"github.com/tidwall/gjson"
)

// VestingParser -
type VestingParser struct {
	filesDirectory string
	interfaces     map[string]kinds.ContractKind
}

// NewVestingParser -
func NewVestingParser(filesDirectory string, interfaces map[string]kinds.ContractKind) *VestingParser {
	return &VestingParser{
		filesDirectory: filesDirectory,
		interfaces:     interfaces,
	}
}

// Parse -
func (p *VestingParser) Parse(data gjson.Result, head noderpc.Header, network, address string) ([]models.Model, error) {
	migration := &migration.Migration{
		ID:          helpers.GenerateID(),
		IndexedTime: time.Now().UnixNano() / 1000,

		Level:     head.Level,
		Network:   network,
		Protocol:  head.Protocol,
		Address:   address,
		Timestamp: head.Timestamp,
		Kind:      consts.MigrationBootstrap,
	}
	parsedModels := []models.Model{migration}

	script := data.Get("script")
	op := operation.Operation{
		ID:          helpers.GenerateID(),
		Network:     network,
		Protocol:    head.Protocol,
		Status:      "applied",
		Kind:        consts.Migration,
		Amount:      data.Get("balance").Int(),
		Counter:     data.Get("counter").Int(),
		Source:      data.Get("manager").String(),
		Destination: address,
		Delegate:    data.Get("delegate.value").String(),
		Level:       head.Level,
		Timestamp:   head.Timestamp,
		IndexedTime: time.Now().UnixNano() / 1000,
		Script:      script,
	}

	parser := contract.NewParser(p.interfaces, contract.WithShareDirContractParser(p.filesDirectory))
	contractModels, err := parser.Parse(op)
	if err != nil {
		return nil, err
	}
	if len(contractModels) > 0 {
		parsedModels = append(parsedModels, contractModels...)
	}

	parsedModels = append(parsedModels, &balanceupdate.BalanceUpdate{
		ID:       helpers.GenerateID(),
		Change:   op.Amount,
		Network:  op.Network,
		Contract: address,
		Level:    head.Level,
	})

	return parsedModels, nil
}
