package handlers

import (
	"net/http"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/gin-gonic/gin"
)

// GetBySlug godoc
// @Summary Get contract by slug
// @Description Get contract by slug
// @Tags contract
// @ID get-contract-by-slug
// @Param slug path string true "Slug"
// @Accept  json
// @Produce  json
// @Success 200 {object} Alias
// @Success 204 {object} gin.H
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /v1/slug/{slug} [get]
func GetBySlug() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("context").(*config.Context)

		var req getBySlugRequest
		if err := c.BindUri(&req); handleError(c, ctx.Storage, err, http.StatusBadRequest) {
			return
		}

		a, err := ctx.ContractMetadata.GetBySlug(req.Slug)
		if handleError(c, ctx.Storage, err, 0) {
			return
		}
		if a == nil {
			c.SecureJSON(http.StatusNoContent, gin.H{})
			return
		}
		var alias Alias
		alias.FromModel(a)
		c.SecureJSON(http.StatusOK, alias)
	}
}
