package migrations

import (
	"context"

	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/migration"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/protocol"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/parsers"
	"github.com/baking-bad/bcdhub/internal/parsers/contract"
	"github.com/pkg/errors"
)

// ImplicitParser -
type ImplicitParser struct {
	ctx      *config.Context
	rpc      noderpc.INode
	protocol protocol.Protocol
}

// NewImplicitParser -
func NewImplicitParser(ctx *config.Context, rpc noderpc.INode, protocol protocol.Protocol) *ImplicitParser {
	return &ImplicitParser{ctx, rpc, protocol}
}

// Parse -
func (p *ImplicitParser) Parse(ctx context.Context, metadata noderpc.Metadata, head noderpc.Header) (*parsers.Result, error) {
	if len(metadata.ImplicitOperationsResults) == 0 {
		return nil, nil
	}

	parserResult := parsers.NewResult()
	for i := range metadata.ImplicitOperationsResults {
		switch metadata.ImplicitOperationsResults[i].Kind {
		case consts.Origination:
			if err := p.origination(ctx, metadata.ImplicitOperationsResults[i], head, p.protocol.ID, parserResult); err != nil {
				return nil, err
			}
		case consts.Transaction:
			if err := p.transaction(metadata.ImplicitOperationsResults[i], head, p.protocol.ID, parserResult); err != nil {
				return nil, err
			}
		}
	}
	return parserResult, nil
}

func (p *ImplicitParser) origination(ctx context.Context, implicit noderpc.ImplicitOperationsResult, head noderpc.Header, protocolID int64, result *parsers.Result) error {
	origination := operation.Operation{
		ProtocolID: protocolID,
		Level:      head.Level,
		Timestamp:  head.Timestamp,
		Kind:       types.OperationKindOrigination,
		Destination: account.Account{
			Address: implicit.OriginatedContracts[0],
			Type:    types.AccountTypeContract,
		},
		ConsumedGas:         implicit.ConsumedGas,
		PaidStorageSizeDiff: implicit.PaidStorageSizeDiff,
		StorageSize:         implicit.StorageSize,
		DeffatedStorage:     implicit.Storage,
	}

	script, err := p.rpc.GetRawScript(ctx, origination.Destination.Address, origination.Level)
	if err != nil {
		return err
	}
	origination.Script = script

	contractParser := contract.NewParser(p.ctx)
	if err := contractParser.Parse(&origination, p.protocol.SymLink, result); err != nil {
		return err
	}

	for i := range result.Contracts {
		if result.Contracts[i].Account.Address == implicit.OriginatedContracts[0] {
			result.Migrations = append(result.Migrations, &migration.Migration{
				ProtocolID: protocolID,
				Level:      head.Level,
				Timestamp:  head.Timestamp,
				Kind:       types.MigrationKindBootstrap,
				Contract:   result.Contracts[i],
			})
			break
		}
	}

	logger.Info().Msg("Implicit bootstrap migration found")

	return nil
}

func (p *ImplicitParser) transaction(implicit noderpc.ImplicitOperationsResult, head noderpc.Header, protocolID int64, result *parsers.Result) error {
	tx := operation.Operation{
		ProtocolID:      protocolID,
		Level:           head.Level,
		Timestamp:       head.Timestamp,
		Kind:            types.OperationKindTransaction,
		ConsumedGas:     implicit.ConsumedGas,
		StorageSize:     implicit.StorageSize,
		DeffatedStorage: implicit.Storage,
		Status:          types.OperationStatusApplied,
		Tags:            types.NewTags([]string{types.ImplicitOperationStringTag}),
		Counter:         head.Level,
	}

	for i := range implicit.BalanceUpdates {
		if implicit.BalanceUpdates[i].Kind == "contract" && implicit.BalanceUpdates[i].Origin == "subsidy" {
			tx.Destination = account.Account{
				Type:    types.NewAccountType(implicit.BalanceUpdates[i].Contract),
				Address: implicit.BalanceUpdates[i].Contract,
			}
			break
		}
	}

	if tx.Destination.Address == "" {
		return errors.Errorf("empty destination in implicit transaction at level %d", head.Level)
	}

	result.Operations = append(result.Operations, &tx)

	return nil
}
