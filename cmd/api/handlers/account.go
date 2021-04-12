package handlers

import (
	"net/http"
	"strings"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GetInfo godoc
// @Summary Get account info
// @Description Get account info
// @Tags account
// @ID get-account-info
// @Param network path string true "Network"
// @Param address path string true "Address" minlength(36) maxlength(36)
// @Accept  json
// @Produce  json
// @Success 200 {object} AccountInfo
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /v1/account/{network}/{address} [get]
func (ctx *Context) GetInfo(c *gin.Context) {
	var req getContractRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusNotFound) {
		return
	}

	stats, err := ctx.Operations.GetStats(req.Network, req.Address)
	if ctx.handleError(c, err, 0) {
		return
	}
	block, err := ctx.Blocks.Last(req.Network)
	if ctx.handleError(c, err, 0) {
		return
	}

	rpc, err := ctx.GetRPC(req.Network)
	if ctx.handleError(c, err, 0) {
		return
	}
	balance, err := rpc.GetContractBalance(req.Address, block.Level)
	if ctx.handleError(c, err, 0) {
		return
	}

	accountInfo := AccountInfo{
		Address:    req.Address,
		Network:    req.Network,
		TxCount:    stats.Count,
		Balance:    balance,
		LastAction: stats.LastAction,
	}

	alias, err := ctx.TZIP.Get(req.Network, req.Address)
	if err != nil {
		if !ctx.Storage.IsRecordNotFound(err) {
			ctx.handleError(c, err, 0)
			return
		}
	} else {
		accountInfo.Alias = alias.Name
	}

	c.JSON(http.StatusOK, accountInfo)
}

// GetBatchTokenBalances godoc
// @Summary Batch account token balances
// @Description Batch account token balances
// @Tags account
// @ID get-batch-token-balances
// @Param network path string true "Network"
// @Param address query string false "Comma-separated list of addresses (e.g. addr1,addr2,addr3)"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]TokenBalance
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /v1/account/{network} [get]
func (ctx *Context) GetBatchTokenBalances(c *gin.Context) {
	var req getByNetwork
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}
	var queryParams batchAddressRequest
	if err := c.BindQuery(&queryParams); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}
	address := strings.Split(queryParams.Address, ",")
	for i := range address {
		if !bcd.IsAddress(address[i]) {
			ctx.handleError(c, errors.Errorf("Invalid address: %s", address[i]), http.StatusBadRequest)
			return
		}
	}

	balances, err := ctx.TokenBalances.Batch(req.Network, address)
	if ctx.handleError(c, err, 0) {
		return
	}

	result := make(map[string][]TokenBalance)
	for a, b := range balances {
		result[a] = make([]TokenBalance, len(b))
		for i := range b {
			result[a][i] = TokenBalance{
				Balance: b[i].BalanceString,
				TokenMetadata: TokenMetadata{
					TokenID:  b[i].TokenID,
					Contract: b[i].Contract,
					Network:  b[i].Network,
				},
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetAccountTokenBalances godoc
// @Summary Get account token balances
// @Description Get account token balances
// @Tags account
// @ID get-account-token-balances
// @Param network path string true "Network"
// @Param address path string true "Address" minlength(36) maxlength(36)
// @Param offset query integer false "Offset"
// @Param size query integer false "Requested count" minimum(0) maximum(10)
// @Param contract query string false "Contract address"
// @Accept  json
// @Produce  json
// @Success 200 {object} TokenBalances
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /v1/account/{network}/{address}/token_balances [get]
func (ctx *Context) GetAccountTokenBalances(c *gin.Context) {
	var req getContractRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusNotFound) {
		return
	}
	var queryParams tokenBalanceRequest
	if err := c.BindQuery(&queryParams); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}
	balances, err := ctx.getAccountBalances(req.Network, req.Address, queryParams)
	if ctx.handleError(c, err, 0) {
		return
	}

	c.JSON(http.StatusOK, balances)
}

func (ctx *Context) getAccountBalances(network, address string, req tokenBalanceRequest) (*TokenBalances, error) {
	tokenBalances, total, err := ctx.TokenBalances.GetAccountBalances(network, address, req.Contract, req.Size, req.Offset)
	if err != nil {
		return nil, err
	}

	response := TokenBalances{
		Balances: make([]TokenBalance, 0),
		Total:    total,
	}

	contextes := make([]tokenmetadata.GetContext, 0)
	balances := make(map[tokenmetadata.GetContext]string)

	for i := range tokenBalances {
		c := tokenmetadata.GetContext{
			TokenID:  &tokenBalances[i].TokenID,
			Contract: tokenBalances[i].Contract,
			Network:  network,
		}
		balances[c] = tokenBalances[i].BalanceString
		contextes = append(contextes, c)
	}

	tokens, err := ctx.TokenMetadata.GetAll(contextes...)
	if err != nil {
		return nil, err
	}

	for _, token := range tokens {
		c := tokenmetadata.GetContext{
			TokenID:  &token.TokenID,
			Contract: token.Contract,
			Network:  network,
		}

		balance, ok := balances[c]
		if !ok {
			continue
		}

		delete(balances, c)

		tb := TokenBalance{
			Balance:       balance,
			TokenMetadata: TokenMetadataFromElasticModel(token, false),
		}

		response.Balances = append(response.Balances, tb)
	}

	for c, balance := range balances {
		response.Balances = append(response.Balances, TokenBalance{
			Balance: balance,
			TokenMetadata: TokenMetadata{
				Contract: c.Contract,
				TokenID:  *c.TokenID,
				Network:  c.Network,
			},
		})
	}

	return &response, nil
}

// GetAccountTokenBalancesGroupedCount godoc
// @Summary Get account token balances count grouped by count
// @Description Get account token balances count grouped by count
// @Tags account
// @ID get-account-token-balances-count
// @Param network path string true "Network"
// @Param address path string true "Address" minlength(36) maxlength(36)
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int64
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /v1/account/{network}/{address}/count [get]
func (ctx *Context) GetAccountTokenBalancesGroupedCount(c *gin.Context) {
	var req getContractRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusNotFound) {
		return
	}
	res, err := ctx.TokenBalances.CountByContract(req.Network, req.Address)
	if ctx.handleError(c, err, 0) {
		return
	}
	c.JSON(http.StatusOK, res)
}
