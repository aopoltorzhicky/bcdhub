package handlers

import (
	"net/http"

	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/gin-gonic/gin"
)

// GetContractTransfers godoc
// @Summary Show contract`s tokens transfers
// @Description Show contract`s tokens transfers
// @Tags contract
// @ID get-contract-transfers
// @Param size query integer false "Transfers count" mininum(1)
// @Param offset query integer false "Offset" mininum(1)
// @Param token_id query integer false "Token ID" mininum(1)
// @Accept  json
// @Produce  json
// @Success 200 {object} TransferResponse
// @Failure 500 {object} Error
// @Router /{network}/{address}/transfers [get]
func (ctx *Context) GetContractTransfers(c *gin.Context) {
	var req getContractTransfers
	if err := c.BindQuery(&req); handleError(c, err, http.StatusBadRequest) {
		return
	}
	var contractRequest getContractRequest
	if err := c.BindUri(&contractRequest); handleError(c, err, http.StatusBadRequest) {
		return
	}

	tokenID := int64(-1)
	if req.TokenID != nil {
		tokenID = int64(*req.TokenID)
	}

	transfers, err := ctx.ES.GetContractTransfers(contractRequest.Network, contractRequest.Address, tokenID, req.Size, req.Offset)
	if handleError(c, err, 0) {
		return
	}
	response, err := ctx.transfersPostprocessing(transfers)
	if handleError(c, err, 0) {
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctx *Context) transfersPostprocessing(transfers elastic.TransfersResponse) (response TransferResponse, err error) {
	response.Total = transfers.Total
	response.Transfers = make([]Transfer, len(transfers.Transfers))

	for i := range transfers.Transfers {
		response.Transfers[i] = Transfer{&transfers.Transfers[i], nil}
		token, ok := ctx.Tokens[tokenKey{
			Network:  transfers.Transfers[i].Network,
			Contract: transfers.Transfers[i].Contract,
			TokenID:  transfers.Transfers[i].TokenID,
		}]
		if ok {
			response.Transfers[i].Token = &token
		}

	}
	return
}
