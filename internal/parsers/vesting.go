package parsers

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/contractparser/consts"
	"github.com/baking-bad/bcdhub/internal/contractparser/kinds"
	"github.com/baking-bad/bcdhub/internal/contractparser/meta"
	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/tidwall/gjson"
)

// VestingParser -
type VestingParser struct {
	rpc            noderpc.INode
	es             elastic.IElastic
	filesDirectory string
	interfaces     map[string][]kinds.Entrypoint
}

// NewVestingParser -
func NewVestingParser(rpc noderpc.INode, es elastic.IElastic, filesDirectory string, interfaces map[string][]kinds.Entrypoint) *VestingParser {
	return &VestingParser{
		rpc:            rpc,
		es:             es,
		filesDirectory: filesDirectory,
		interfaces:     interfaces,
	}
}

// Parse -
func (p *VestingParser) Parse(data gjson.Result, head noderpc.Header, network, address string) ([]elastic.Model, error) {
	migration := &models.Migration{
		ID:          helpers.GenerateID(),
		IndexedTime: time.Now().UnixNano() / 1000,

		Level:     head.Level,
		Network:   network,
		Protocol:  head.Protocol,
		Address:   address,
		Timestamp: head.Timestamp,
		Kind:      consts.MigrationBootstrap,
	}
	parsedModels := []elastic.Model{migration}

	script := data.Get("script")
	op := models.Operation{
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
		Timestamp:   head.Timestamp,
		IndexedTime: time.Now().UnixNano() / 1000,
		Script:      script,
	}

	protoSymLink, err := meta.GetProtoSymLink(migration.Protocol)
	if err != nil {
		return nil, err
	}
	contractModels, err := createNewContract(p.es, p.interfaces, op, p.filesDirectory, protoSymLink)
	if err != nil {
		return nil, err
	}
	if len(contractModels) > 0 {
		parsedModels = append(parsedModels, contractModels...)
	}

	return parsedModels, nil
}
