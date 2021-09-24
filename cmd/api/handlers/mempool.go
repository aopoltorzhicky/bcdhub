package handlers

import (
	"net/http"
	"time"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/bcd/tezerrors"
	"github.com/baking-bad/bcdhub/internal/bcd/types"
	"github.com/baking-bad/bcdhub/internal/helpers"
	modelTypes "github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/services/mempool"
	"github.com/gin-gonic/gin"
)

// GetMempool godoc
// @Summary Get contract mempool operations
// @Description Get contract mempool operations
// @Tags contract
// @ID get-contract-mempool
// @Param network path string true "Network"
// @Param address path string true "KT address" minlength(36) maxlength(36)
// @Accept  json
// @Produce  json
// @Success 200 {array} Operation
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /v1/contract/{network}/{address}/mempool [get]
func (ctx *Context) GetMempool(c *gin.Context) {
	var req getContractRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusNotFound) {
		return
	}

	mempoolService, err := ctx.GetMempoolService(req.NetworkID())
	if err != nil {
		c.SecureJSON(http.StatusNoContent, []Operation{})
		return
	}

	res, err := mempoolService.Get(req.Address)
	if ctx.handleError(c, err, 0) {
		return
	}

	c.SecureJSON(http.StatusOK, ctx.mempoolPostprocessing(res, req.NetworkID()))
}

func (ctx *Context) mempoolPostprocessing(res mempool.PendingOperations, network modelTypes.Network) []Operation {
	ret := make([]Operation, 0)
	if len(res.Originations)+len(res.Transactions) == 0 {
		return ret
	}

	for _, origination := range res.Originations {
		op := ctx.prepareMempoolOrigination(network, origination)
		if op != nil {
			ret = append(ret, *op)
		}
	}

	for _, tx := range res.Transactions {
		op := ctx.prepareMempoolTransaction(network, tx)
		if op != nil {
			ret = append(ret, *op)
		}
	}

	return ret
}

func (ctx *Context) prepareMempoolTransaction(network modelTypes.Network, tx mempool.PendingTransaction) *Operation {
	status := tx.Status
	if status == consts.Applied {
		status = consts.Pending
	}
	if !helpers.StringInArray(tx.Kind, []string{consts.Transaction, consts.Origination, consts.OriginationNew}) {
		return nil
	}

	amount, err := tx.Amount.Int64()
	if err != nil {
		return nil
	}

	op := Operation{
		Hash:             tx.Hash,
		Network:          network.String(),
		Timestamp:        time.Unix(tx.UpdatedAt, 0).UTC(),
		SourceAlias:      ctx.CachedAlias(network, tx.Source),
		DestinationAlias: ctx.CachedAlias(network, tx.Destination),
		Kind:             tx.Kind,
		Source:           tx.Source,
		Fee:              tx.Fee,
		Counter:          tx.Counter,
		GasLimit:         tx.GasLimit,
		StorageLimit:     tx.StorageLimit,
		Amount:           amount,
		Destination:      tx.Destination,
		Mempool:          true,
		Status:           status,
		RawMempool:       tx.Raw,
		Protocol:         tx.Protocol,
	}

	errs, err := tezerrors.ParseArray(tx.Errors)
	if err != nil {
		return nil
	}
	op.Errors = errs

	if bcd.IsContract(op.Destination) && op.Protocol != "" && op.Status == consts.Pending {
		if len(tx.Parameters) > 0 {
			_ = ctx.buildMempoolOperationParameters(tx.Parameters, &op)
		} else {
			op.Entrypoint = consts.DefaultEntrypoint
		}
	}

	return &op
}

func (ctx *Context) prepareMempoolOrigination(network modelTypes.Network, origination mempool.PendingOrigination) *Operation {
	status := origination.Status
	if status == consts.Applied {
		status = consts.Pending
	}
	if !helpers.StringInArray(origination.Kind, []string{consts.Transaction, consts.Origination, consts.OriginationNew}) {
		return nil
	}

	op := Operation{
		Hash:         origination.Hash,
		Network:      network.String(),
		Timestamp:    time.Unix(origination.UpdatedAt, 0).UTC(),
		SourceAlias:  ctx.CachedAlias(network, origination.Source),
		Kind:         origination.Kind,
		Source:       origination.Source,
		Fee:          origination.Fee,
		Counter:      origination.Counter,
		GasLimit:     origination.GasLimit,
		StorageLimit: origination.StorageLimit,
		Mempool:      true,
		Status:       status,
		RawMempool:   origination.Raw,
		Protocol:     origination.Protocol,
	}

	errs, err := tezerrors.ParseArray(origination.Errors)
	if err != nil {
		return nil
	}
	op.Errors = errs
	return &op
}

func (ctx *Context) buildMempoolOperationParameters(data []byte, op *Operation) error {
	network := modelTypes.NewNetwork(op.Network)
	proto, err := ctx.CachedProtocolByHash(network, op.Protocol)
	if err != nil {
		return err
	}
	script, err := ctx.getScript(network, op.Destination, proto.SymLink)
	if err != nil {
		return err
	}
	parameter, err := script.ParameterType()
	if err != nil {
		return err
	}
	params := types.NewParameters(data)
	op.Entrypoint = params.Entrypoint

	tree, err := parameter.FromParameters(params)
	if err != nil {
		return err
	}

	op.Parameters, err = tree.ToMiguel()
	if err != nil && !tezerrors.HasParametersError(op.Errors) {
		return err
	}
	return nil
}
