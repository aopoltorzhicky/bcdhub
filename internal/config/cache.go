package config

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/baking-bad/bcdhub/internal/models/block"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/baking-bad/bcdhub/internal/models/tzip"
)

// CachedAlias -
func (ctx *Context) CachedAlias(network, address string) string {
	if !bcd.IsContract(address) {
		return ""
	}
	key := ctx.Cache.AliasKey(network, address)
	item, err := ctx.Cache.Fetch(key, time.Minute*30, func() (interface{}, error) {
		return ctx.TZIP.Get(network, address)
	})
	if err != nil {
		return ""
	}

	if data, ok := item.Value().(*tzip.TZIP); ok && data != nil {
		return data.Name
	}
	return ""
}

// CachedTokenMetadata -
func (ctx *Context) CachedTokenMetadata(network, address string, tokenID uint64) (*tokenmetadata.TokenMetadata, error) {
	key := ctx.Cache.TokenMetadataKey(network, address, tokenID)
	item, err := ctx.Cache.Fetch(key, time.Minute*30, func() (interface{}, error) {
		return ctx.TokenMetadata.GetOne(network, address, tokenID)
	})
	if err != nil {
		return nil, err
	}
	return item.Value().(*tokenmetadata.TokenMetadata), nil
}

// CachedCurrentBlock -
func (ctx *Context) CachedCurrentBlock(network string) (block.Block, error) {
	key := ctx.Cache.BlockKey(network)
	item, err := ctx.Cache.Fetch(key, time.Second*15, func() (interface{}, error) {
		return ctx.Blocks.Last(network)
	})
	if err != nil {
		return block.Block{}, err
	}
	return item.Value().(block.Block), nil
}

// CachedTezosBalance -
func (ctx *Context) CachedTezosBalance(network, address string, level int64) (int64, error) {
	key := ctx.Cache.TezosBalanceKey(network, address, level)
	item, err := ctx.Cache.Fetch(key, 30*time.Second, func() (interface{}, error) {
		rpc, err := ctx.GetRPC(network)
		if err != nil {
			return 0, err
		}
		return rpc.GetContractBalance(address, level)
	})
	if err != nil {
		return 0, err
	}
	return item.Value().(int64), nil
}

// CachedContract -
func (ctx *Context) CachedContract(network, address string) (*contract.Contract, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}

	key := ctx.Cache.ContractKey(network, address)
	item, err := ctx.Cache.Fetch(key, time.Minute*10, func() (interface{}, error) {
		return ctx.Contracts.Get(network, address)
	})
	if err != nil {
		return nil, err
	}
	cntr := item.Value().(contract.Contract)
	return &cntr, nil
}