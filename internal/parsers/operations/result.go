package operations

import (
	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/bcd/tezerrors"
	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
)

func parseOperationResult(data noderpc.Operation, tx *operation.Operation) {
	result := data.GetResult()
	if result == nil {
		return
	}

	tx.Status = types.NewOperationStatus(result.Status)
	tx.ConsumedGas = result.ConsumedGas
	if result.StorageSize != nil {
		tx.StorageSize = *result.StorageSize
	}
	if result.PaidStorageSizeDiff != nil {
		tx.PaidStorageSizeDiff = *result.PaidStorageSizeDiff
	}
	if len(result.Originated) > 0 {
		tx.Destination = account.Account{
			Network: tx.Network,
			Address: result.Originated[0],
			Type:    types.AccountTypeContract,
		}
	}

	tx.AllocatedDestinationContract = data.Kind == consts.Origination
	if result.AllocatedDestinationContract != nil {
		tx.AllocatedDestinationContract = *result.AllocatedDestinationContract
	}

	if errs, err := tezerrors.ParseArray(result.Errors); err == nil {
		tx.Errors = errs
	}
}
