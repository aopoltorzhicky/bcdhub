package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStats -
func (ctx *Context) GetStats(c *gin.Context) {
	stats, err := ctx.ES.GetAllStates()
	if handleError(c, err, 0) {
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetNetworkStats -
func (ctx *Context) GetNetworkStats(c *gin.Context) {
	var req getByNetwork
	if err := c.BindUri(&req); handleError(c, err, http.StatusBadRequest) {
		return
	}

	var stats NetworkStats
	counts, err := ctx.ES.GetItemsCountForNetwork(req.Network)
	if handleError(c, err, 0) {
		return
	}
	stats.ContractsCount = counts.Contracts
	stats.OperationsCount = counts.Operations

	protocols, err := ctx.ES.GetProtocolsByNetwork(req.Network)
	if handleError(c, err, 0) {
		return
	}
	stats.Protocols = protocols

	c.JSON(http.StatusOK, stats)
}

// GetSeries -
func (ctx *Context) GetSeries(c *gin.Context) {
	var req getByNetwork
	if err := c.BindUri(&req); handleError(c, err, http.StatusBadRequest) {
		return
	}

	var reqArgs getSeriesRequest
	if err := c.BindQuery(&reqArgs); handleError(c, err, http.StatusBadRequest) {
		return
	}

	series, err := ctx.ES.GetDateHistogram(req.Network, reqArgs.Index, reqArgs.Period)
	if handleError(c, err, 0) {
		return
	}

	c.JSON(http.StatusOK, series)
}
