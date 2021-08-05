package config

import (
	"github.com/baking-bad/bcdhub/internal/aws"
	"github.com/baking-bad/bcdhub/internal/cache"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapaction"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/block"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/dapp"
	"github.com/baking-bad/bcdhub/internal/models/domains"
	"github.com/baking-bad/bcdhub/internal/models/migration"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/protocol"
	"github.com/baking-bad/bcdhub/internal/models/service"
	"github.com/baking-bad/bcdhub/internal/models/tezosdomain"
	"github.com/baking-bad/bcdhub/internal/models/tokenbalance"
	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/baking-bad/bcdhub/internal/models/transfer"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/models/tzip"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/pinata"
	"github.com/baking-bad/bcdhub/internal/postgres/core"
	"github.com/baking-bad/bcdhub/internal/search"
	"github.com/baking-bad/bcdhub/internal/tzkt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

// Context -
type Context struct {
	AWS          *aws.Client
	RPC          map[types.Network]noderpc.INode
	TzKTServices map[types.Network]tzkt.Service
	Pinata       pinata.Service

	StorageDB *core.Postgres

	Config     Config
	SharePath  string
	TzipSchema string

	TezosDomainsContracts map[types.Network]string

	Storage       models.GeneralRepository
	BigMapActions bigmapaction.Repository
	BigMapDiffs   bigmapdiff.Repository
	Blocks        block.Repository
	Contracts     contract.Repository
	DApps         dapp.Repository
	Migrations    migration.Repository
	Operations    operation.Repository
	Protocols     protocol.Repository
	TezosDomains  tezosdomain.Repository
	TokenBalances tokenbalance.Repository
	TokenMetadata tokenmetadata.Repository
	Transfers     transfer.Repository
	TZIP          tzip.Repository
	Domains       domains.Repository
	Services      service.Repository

	Searcher search.Searcher

	Cache     *cache.Cache
	Sanitizer *bluemonday.Policy
}

// NewContext -
func NewContext(opts ...ContextOption) *Context {
	ctx := &Context{
		Cache:     cache.NewCache(),
		Sanitizer: bluemonday.UGCPolicy(),
	}
	ctx.Sanitizer.AllowAttrs("em")

	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

// GetRPC -
func (ctx *Context) GetRPC(network types.Network) (noderpc.INode, error) {
	if rpc, ok := ctx.RPC[network]; ok {
		return rpc, nil
	}
	return nil, errors.Errorf("Unknown rpc network %s", network)
}

// GetTzKTService -
func (ctx *Context) GetTzKTService(network types.Network) (tzkt.Service, error) {
	if rpc, ok := ctx.TzKTServices[network]; ok {
		return rpc, nil
	}
	return nil, errors.Errorf("Unknown tzkt service network %s", network)
}

// Close -
func (ctx *Context) Close() {
	if ctx.StorageDB != nil {
		ctx.StorageDB.Close()
	}
}
