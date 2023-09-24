package rollback

import (
	"context"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapaction"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/block"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/migration"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	smartrollup "github.com/baking-bad/bcdhub/internal/models/smart_rollup"
	"github.com/baking-bad/bcdhub/internal/models/ticket"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/pkg/errors"
)

// Manager -
type Manager struct {
	storage   models.GeneralRepository
	blockRepo block.Repository
	saver     Saver
}

// NewManager -
func NewManager(storage models.GeneralRepository, blockRepo block.Repository, saver Saver) Manager {
	return Manager{
		storage, blockRepo, saver,
	}
}

// Rollback - rollback indexer state to level
func (rm Manager) Rollback(ctx context.Context, network types.Network, fromState block.Block, toLevel int64) error {
	if toLevel >= fromState.Level {
		return errors.Errorf("To level must be less than from level: %d >= %d", toLevel, fromState.Level)
	}

	for level := fromState.Level; level > toLevel; level-- {
		logger.Info().Str("network", network.String()).Msgf("Rollback to %d block", level)

		if _, err := rm.blockRepo.Get(ctx, level); err != nil {
			if rm.storage.IsRecordNotFound(err) {
				continue
			}
			return err
		}

		if err := rm.rollback(ctx, level); err != nil {
			logger.Error().Err(err).Str("network", network.String()).Msg("rollback error")
			return rm.saver.Rollback()
		}
	}
	return nil
}

// TODO: rollback protocol and m2m entities
func (rm Manager) rollback(ctx context.Context, level int64) error {
	if err := rm.rollbackOperations(ctx, level); err != nil {
		return err
	}
	if err := rm.rollbackBigMapState(ctx, level); err != nil {
		return err
	}
	if err := rm.rollbackAll(ctx, level); err != nil {
		return err
	}
	if err := rm.saver.Commit(); err != nil {
		return err
	}
	return nil
}

func (rm Manager) rollbackAll(ctx context.Context, level int64) error {
	for _, model := range []models.Model{
		&block.Block{}, &contract.Contract{}, &bigmapdiff.BigMapDiff{},
		&bigmapaction.BigMapAction{}, &contract.GlobalConstant{},
		&smartrollup.SmartRollup{}, &ticket.TicketUpdate{},
		&migration.Migration{},
	} {
		if err := rm.saver.DeleteAll(ctx, model, level); err != nil {
			return err
		}
		logger.Info().
			Str("model", model.GetIndex()).
			Msg("rollback")
	}
	return nil
}

func (rm Manager) rollbackBigMapState(ctx context.Context, level int64) error {
	logger.Info().Msg("rollback big map states...")
	states, err := rm.saver.StatesChangedAtLevel(ctx, level)
	if err != nil {
		return err
	}

	for i, state := range states {
		diff, err := rm.saver.LastDiff(ctx, state.Ptr, state.KeyHash, false)
		if err != nil {
			if rm.storage.IsRecordNotFound(err) {
				if err := rm.saver.DeleteBigMapState(ctx, states[i]); err != nil {
					return err
				}
				continue
			}
			return err
		}
		states[i].LastUpdateLevel = diff.Level
		states[i].LastUpdateTime = diff.Timestamp
		states[i].IsRollback = true

		if len(diff.Value) > 0 {
			states[i].Value = diff.ValueBytes()
			states[i].Removed = false
		} else {
			states[i].Removed = true
			valuedDiff, err := rm.saver.LastDiff(ctx, state.Ptr, state.KeyHash, true)
			if err != nil {
				if !rm.storage.IsRecordNotFound(err) {
					return err
				}
			} else {
				states[i].Value = valuedDiff.ValueBytes()
			}
		}

		if err := rm.saver.SaveBigMapState(ctx, states[i]); err != nil {
			return err
		}
	}

	return nil
}

func (rm Manager) rollbackOperations(ctx context.Context, level int64) error {
	logger.Info().Msg("rollback operations...")

	ops, err := rm.saver.GetOperations(ctx, level)
	if err != nil {
		return err
	}
	if len(ops) == 0 {
		return nil
	}

	if err := rm.saver.DeleteAll(ctx, new(operation.Operation), level); err != nil {
		return err
	}

	contracts := make(map[int64]int64)
	for i := range ops {
		if ops[i].IsOrigination() {
			continue
		}
		if ops[i].Destination.Type == types.AccountTypeContract {
			if _, ok := contracts[ops[i].DestinationID]; !ok {
				contracts[ops[i].DestinationID] = 1
			} else {
				contracts[ops[i].DestinationID] += 1
			}
		}
		if ops[i].Source.Type == types.AccountTypeContract {
			if _, ok := contracts[ops[i].SourceID]; !ok {
				contracts[ops[i].SourceID] = 1
			} else {
				contracts[ops[i].SourceID] += 1
			}
		}
	}

	if len(contracts) == 0 {
		return nil
	}

	addresses := make([]int64, 0, len(contracts))
	for address := range contracts {
		addresses = append(addresses, address)
	}

	actions, err := rm.saver.GetContractsLastAction(ctx, addresses...)
	if err != nil {
		return err
	}

	for i := range actions {
		count, ok := contracts[actions[i].AccountId]
		if !ok {
			count = 1
		}

		if err := rm.saver.UpdateContractStats(ctx, actions[i].AccountId, actions[i].Time, count); err != nil {
			return err
		}
	}

	return nil
}
