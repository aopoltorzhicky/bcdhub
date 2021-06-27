package services

import (
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/handlers"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/pkg/errors"
)

// ContractMetadataHandler -
type ContractMetadataHandler struct {
	*config.Context
	handler *handlers.ContractMetadata
}

// NewContractMetadataHandler -
func NewContractMetadataHandler(ctx *config.Context) *ContractMetadataHandler {
	return &ContractMetadataHandler{
		ctx,
		handlers.NewContractMetadata(ctx.BigMapDiffs, ctx.Blocks, ctx.Storage, ctx.TZIP, ctx.RPC, ctx.SharePath, ctx.Config.IPFSGateways),
	}
}

// Handle -
func (cm *ContractMetadataHandler) Handle(items []models.Model) error {
	if len(items) == 0 {
		return nil
	}

	updates := make([]models.Model, 0)
	for i := range items {
		bmd, ok := items[i].(*bigmapdiff.BigMapDiff)
		if !ok {
			return errors.Errorf("[ContractMetadata.Handle] invalid type: expected *bigmapdiff.BigMapDiff got %T", items[i])
		}

		protocol, err := cm.CachedProtocolByID(bmd.Network, bmd.ProtocolID)
		if err != nil {
			return errors.Errorf("[ContractMetadata.Handle] can't get protocol by ID %d in %s: %s", bmd.ProtocolID, bmd.Network.String(), err)
		}

		storageType, err := cm.CachedStorageType(bmd.Network, bmd.Contract, protocol.SymLink)
		if err != nil {
			return errors.Errorf("[ContractMetadata.Handle] can't get storage type for '%s' in %s: %s", bmd.Contract, bmd.Network.String(), err)
		}

		res, err := cm.handler.Do(bmd, storageType)
		if err != nil {
			return errors.Errorf("[ContractMetadata.Handle] compute error message: %s", err)
		}

		updates = append(updates, res...)
	}

	if len(updates) == 0 {
		return nil
	}

	logger.Info("%2d contract metadata are processed", len(updates))

	if err := cm.Storage.Save(updates); err != nil {
		return err
	}

	return saveSearchModels(cm.Context, updates)
}

// Chunk -
func (cm *ContractMetadataHandler) Chunk(lastID, size int64) ([]models.Model, error) {
	var diff []bigmapdiff.BigMapDiff
	if err := getModels(cm.StorageDB.DB, models.DocBigMapDiff, lastID, size, &diff); err != nil {
		return nil, err
	}

	data := make([]models.Model, len(diff))
	for i := range diff {
		data[i] = &diff[i]
	}
	return data, nil
}
