package rollback

import (
	"context"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
)

func (rm Manager) rollbackOperations(ctx context.Context, level int64, rCtx *rollbackContext) error {
	logger.Info().Msg("rollback operations...")

	ops, err := rm.rollback.GetOperations(ctx, level)
	if err != nil {
		return err
	}
	if len(ops) == 0 {
		return nil
	}

	count, err := rm.rollback.DeleteAll(ctx, (*operation.Operation)(nil), level)
	if err != nil {
		return err
	}
	rCtx.generalStats.OperationsCount -= count

	for i := range ops {
		if !ops[i].Destination.IsEmpty() {
			rCtx.applyOperationsCount(ops[i].DestinationID, 1)
			rCtx.applyTicketUpdates(ops[i].DestinationID, int64(ops[i].TicketUpdatesCount))
		}

		if !ops[i].Source.IsEmpty() {
			rCtx.applyOperationsCount(ops[i].SourceID, 1)
		}

		switch ops[i].Kind {
		case types.OperationKindEvent:
			rCtx.generalStats.EventsCount -= 1
			rCtx.applyEvent(ops[i].SourceID)

		case types.OperationKindOrigination:
			rCtx.generalStats.OriginationsCount -= 1

		case types.OperationKindSrOrigination:
			rCtx.generalStats.SrOriginationsCount -= 1

		case types.OperationKindTransaction:
			rCtx.generalStats.TransactionsCount -= 1

		case types.OperationKindRegisterGlobalConstant:
			rCtx.generalStats.RegisterGlobalConstantCount -= 1

		case types.OperationKindSrExecuteOutboxMessage:
			rCtx.generalStats.SrExecutesCount -= 1

		case types.OperationKindTransferTicket:
			rCtx.generalStats.TransferTicketsCount -= 1
		}
	}

	if err := rCtx.getLastActions(ctx, rm.rollback); err != nil {
		return err
	}

	return nil
}
