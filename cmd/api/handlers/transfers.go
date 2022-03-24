package handlers

import (
	"net/http"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/models/domains"
	"github.com/baking-bad/bcdhub/internal/models/transfer"
	"github.com/gin-gonic/gin"
)

// GetContractTransfers godoc
// @Summary Show contract`s tokens transfers
// @Description Show contract`s tokens transfers.
// @Tags contract
// @ID get-contract-transfers
// @Param network path string true "Network"
// @Param address path string true "KT address" minlength(36) maxlength(36)
// @Param size query integer false "Transfers count" mininum(1) maximum(10)
// @Param offset query integer false "Offset" mininum(1)
// @Param token_id query integer false "Token ID" mininum(1)
// @Accept  json
// @Produce  json
// @Success 200 {object} TransferResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /v1/contract/{network}/{address}/transfers [get]
func GetContractTransfers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("context").(*config.Context)

		var contractRequest getContractRequest
		if err := c.BindUri(&contractRequest); handleError(c, ctx.Storage, err, http.StatusNotFound) {
			return
		}
		var req getContractTransfers
		if err := c.BindQuery(&req); handleError(c, ctx.Storage, err, http.StatusBadRequest) {
			return
		}

		transfers, err := ctx.Domains.Transfers(transfer.GetContext{
			Contracts: []string{contractRequest.Address},
			Size:      req.Size,
			Offset:    req.Offset,
			TokenID:   req.TokenID,
			AccountID: -1,
		})
		if handleError(c, ctx.Storage, err, 0) {
			return
		}
		c.SecureJSON(http.StatusOK, transfersPostprocessing(ctx, transfers, false))
	}
}

func transfersPostprocessing(ctx *config.Context, transfers domains.TransfersResponse, withLastID bool) (response TransferResponse) {
	response.Total = transfers.Total
	response.Transfers = make([]Transfer, len(transfers.Transfers))
	if withLastID {
		response.LastID = transfers.LastID
	}

	for i := range transfers.Transfers {
		token := TokenMetadata{
			Contract: transfers.Transfers[i].Contract,
			TokenID:  transfers.Transfers[i].TokenID,
			Symbol:   transfers.Transfers[i].Symbol,
			Decimals: transfers.Transfers[i].Decimals,
			Name:     transfers.Transfers[i].Name,
		}

		response.Transfers[i] = TransferFromModel(transfers.Transfers[i])
		response.Transfers[i].Token = &token
		response.Transfers[i].Alias = ctx.Cache.Alias(transfers.Transfers[i].Contract)
		response.Transfers[i].InitiatorAlias = transfers.Transfers[i].Initiator.Alias
		response.Transfers[i].FromAlias = transfers.Transfers[i].From.Alias
		response.Transfers[i].ToAlias = transfers.Transfers[i].To.Alias
	}
	return
}
