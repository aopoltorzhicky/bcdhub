package bigmapdiff

import (
	"fmt"

	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"gorm.io/gorm"
)

func (storage *Storage) buildGetContext(ctx bigmapdiff.GetContext) *gorm.DB {
	query := storage.DB.Table(models.DocBigMapDiff).Select("max(id) as id, count(id) as keys_count")

	if ctx.Network != 0 {
		query.Where("network = ?", ctx.Network)
	}
	if ctx.Contract != "" {
		query.Where("contract = ?", ctx.Contract)
	}
	if ctx.Ptr != nil {
		query.Where("ptr = ?", *ctx.Ptr)
	}
	if ctx.MaxLevel != nil {
		query.Where("level < ?", *ctx.MaxLevel)
	}
	if ctx.MinLevel != nil {
		query.Where("level >= ?", *ctx.MinLevel)
	}
	if ctx.CurrentLevel != nil {
		query.Where("level = ?", *ctx.CurrentLevel)
	}
	if ctx.Query != "" {
		query.Where("key_hash LIKE %?%", ctx.Query)
	}

	query.Limit(storage.GetPageSize(ctx.Size))

	if ctx.Offset > 0 {
		query.Offset(int(ctx.Offset))
	}

	return query.Group("key_hash").Order("id desc")
}

func (storage *Storage) buildGetContextForState(ctx bigmapdiff.GetContext) *gorm.DB {
	query := storage.DB.Table(models.DocBigMapState)

	if ctx.Network != 0 {
		query.Where("network = ?", ctx.Network)
	}
	if ctx.Contract != "" {
		query.Where("contract = ?", ctx.Contract)
	}
	if ctx.Ptr != nil {
		query.Where("ptr = ?", *ctx.Ptr)
	}
	if ctx.MaxLevel != nil {
		query.Where("level < ?", *ctx.MaxLevel)
	}
	if ctx.MinLevel != nil {
		query.Where("level >= ?", *ctx.MinLevel)
	}
	if ctx.CurrentLevel != nil {
		query.Where("level = ?", *ctx.CurrentLevel)
	}
	if ctx.Query != "" {
		query.Where("key_hash LIKE ?", fmt.Sprintf("%%%s%%", ctx.Query))
	}

	query.Limit(storage.GetPageSize(ctx.Size))

	if ctx.Offset > 0 {
		query.Offset(int(ctx.Offset))
	}

	return query.Order("id desc")
}
