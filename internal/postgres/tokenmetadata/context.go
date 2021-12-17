package tokenmetadata

import (
	"fmt"

	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/go-pg/pg/v10/orm"
)

func (storage *Storage) buildGetTokenMetadataContext(query *orm.Query, ctx ...tokenmetadata.GetContext) {
	if len(ctx) == 0 {
		return
	}

	query = query.WhereOrGroup(func(q *orm.Query) (*orm.Query, error) {
		for i := range ctx {
			q = query.WhereGroup(func(subQuery *orm.Query) (*orm.Query, error) {
				if ctx[i].Network != types.Empty {
					subQuery.Where("network = ?", ctx[i].Network)
				}
				if ctx[i].Contract != "" {
					subQuery.Where("contract = ?", ctx[i].Contract)
				}
				if ctx[i].TokenID != nil {
					subQuery.Where("token_id = ?", *ctx[i].TokenID)
				}
				if ctx[i].MaxLevel > 0 {
					subQuery.Where(fmt.Sprintf("level <= %d", ctx[i].MaxLevel))
				}
				if ctx[i].MinLevel > 0 {
					subQuery.Where(fmt.Sprintf("level > %d", ctx[i].MinLevel))
				}
				if ctx[i].Creator != "" {
					subQuery.Where("? = ANY(creators)", ctx[i].Creator)
				}
				if ctx[i].Name != "" {
					subQuery.Where("name = ?", ctx[i].Name)
				}
				return subQuery, nil
			})
		}
		return q, nil
	})
}
