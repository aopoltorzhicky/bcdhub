package handlers

import (
	"github.com/baking-bad/bcdhub/internal/bcd/ast"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/block"
	"github.com/baking-bad/bcdhub/internal/models/domains"
	"github.com/baking-bad/bcdhub/internal/models/types"
	tzipModel "github.com/baking-bad/bcdhub/internal/models/tzip"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/parsers/tzip"
	"github.com/pkg/errors"
)

// TZIP -
type TZIP struct {
	repo    tzipModel.Repository
	parsers map[types.Network]tzip.Parser
}

// NewTZIP -
func NewTZIP(bigMapRepo bigmapdiff.Repository, blockRepo block.Repository, storage models.GeneralRepository, repo tzipModel.Repository, rpcs map[types.Network]noderpc.INode, sharePath string, ipfs []string) *TZIP {
	parsers := make(map[types.Network]tzip.Parser)
	for network, rpc := range rpcs {
		parsers[network] = tzip.NewParser(bigMapRepo, blockRepo, storage, rpc, tzip.ParserConfig{
			IPFSGateways: ipfs,
			SharePath:    sharePath,
		})
	}
	return &TZIP{
		repo, parsers,
	}
}

// Do -
func (t *TZIP) Do(bmd *domains.BigMapDiff, storage *ast.TypedAst) (bool, []models.Model, error) {
	if bmd.KeyHash != tzip.EmptyStringKey {
		return false, nil, nil
	}
	res, err := t.handle(bmd)
	return true, res, err
}

func (t *TZIP) handle(bmd *domains.BigMapDiff) ([]models.Model, error) {
	tzipParser, ok := t.parsers[bmd.Network]
	if !ok {
		return nil, errors.Errorf("Unknown network for tzip parser: %s", bmd.Network)
	}

	model, err := tzipParser.Parse(tzip.ParseContext{
		BigMapDiff: *bmd.BigMapDiff,
	})
	if err != nil {
		logger.With(bmd).Warn(err)
		return nil, nil
	}
	if model == nil {
		return nil, nil
	}

	m, err := t.repo.Get(model.Network, model.Address)
	if err == nil && m.OffChain {
		return nil, nil
	}

	return []models.Model{model}, nil
}
