package cache

import (
	"fmt"
	"time"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/block"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/contract_metadata"
	"github.com/baking-bad/bcdhub/internal/models/protocol"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/karlseguin/ccache"
	"github.com/microcosm-cc/bluemonday"
)

// Cache -
type Cache struct {
	*ccache.Cache
	network types.Network
	rpc     noderpc.INode

	blocks    block.Repository
	accounts  account.Repository
	contracts contract.Repository
	protocols protocol.Repository
	sanitizer *bluemonday.Policy
	tzip      contract_metadata.Repository
}

// NewCache -
func NewCache(network types.Network, rpc noderpc.INode, blocks block.Repository, accounts account.Repository, contracts contract.Repository, protocols protocol.Repository, cm contract_metadata.Repository, sanitizer *bluemonday.Policy) *Cache {
	return &Cache{
		ccache.New(ccache.Configure().MaxSize(1000)),
		network,
		rpc,
		blocks,
		accounts,
		contracts,
		protocols,
		sanitizer,
		cm,
	}
}

// Alias -
func (cache *Cache) Alias(address string) string {
	if !bcd.IsContract(address) {
		return ""
	}
	key := fmt.Sprintf("alias:%s", address)
	item, err := cache.Fetch(key, time.Minute*30, func() (interface{}, error) {
		acc, err := cache.accounts.Get(cache.network, address)
		if err == nil && acc.Alias != "" {
			return acc.Alias, nil
		}

		cm, err := cache.tzip.Get(cache.network, address)
		if err == nil {
			return cm.Name, nil
		}

		return "", err
	})
	if err != nil {
		return ""
	}

	if data, ok := item.Value().(string); ok && data != "" {
		return cache.sanitizer.Sanitize(data)
	}
	return ""
}

// ContractMetadata -
func (cache *Cache) ContractMetadata(address string) (*contract_metadata.ContractMetadata, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}
	key := fmt.Sprintf("contract_metadata:%s", address)
	item, err := cache.Fetch(key, time.Minute*30, func() (interface{}, error) {
		return cache.tzip.Get(cache.network, address)
	})
	if err != nil {
		return nil, err
	}

	return item.Value().(*contract_metadata.ContractMetadata), nil
}

// Events -
func (cache *Cache) Events(address string) (contract_metadata.Events, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}
	key := fmt.Sprintf("contract_metadata:%s", address)
	item, err := cache.Fetch(key, time.Hour, func() (interface{}, error) {
		return cache.tzip.Events(cache.network, address)
	})
	if err != nil {
		return nil, err
	}

	return item.Value().(contract_metadata.Events), nil
}

// Contract -
func (cache *Cache) Contract(address string) (*contract.Contract, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}

	key := fmt.Sprintf("contract:%s", address)
	item, err := cache.Fetch(key, time.Minute*10, func() (interface{}, error) {
		return cache.contracts.Get(cache.network, address)
	})
	if err != nil {
		return nil, err
	}
	cntr := item.Value().(contract.Contract)
	return &cntr, nil
}

// ContractTags -
func (cache *Cache) ContractTags(address string) (types.Tags, error) {
	if !bcd.IsContract(address) {
		return 0, nil
	}

	key := fmt.Sprintf("contract:%s", address)
	item, err := cache.Fetch(key, time.Minute*10, func() (interface{}, error) {
		c, err := cache.contracts.Get(cache.network, address)
		if err != nil {
			return 0, err
		}
		return c.Tags, nil
	})
	if err != nil {
		return 0, err
	}
	return item.Value().(types.Tags), nil
}

// ProjectIDByHash -
func (cache *Cache) ProjectIDByHash(hash string) string {
	return fmt.Sprintf("project_id:%s", hash)
}

// CurrentBlock -
func (cache *Cache) CurrentBlock() (block.Block, error) {
	item, err := cache.Fetch("block", time.Second*15, func() (interface{}, error) {
		return cache.blocks.Last(cache.network)
	})
	if err != nil {
		return block.Block{}, err
	}
	return item.Value().(block.Block), nil
}

//nolint
// TezosBalance -
func (cache *Cache) TezosBalance(address string, level int64) (int64, error) {
	key := fmt.Sprintf("tezos_balance:%s:%d", address, level)
	item, err := cache.Fetch(key, 30*time.Second, func() (interface{}, error) {
		return cache.rpc.GetContractBalance(address, level)
	})
	if err != nil {
		return 0, err
	}
	return item.Value().(int64), nil
}

// ScriptBytes -
func (cache *Cache) ScriptBytes(address, symLink string) ([]byte, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}

	key := fmt.Sprintf("script_bytes:%s", address)
	item, err := cache.Fetch(key, time.Hour, func() (interface{}, error) {
		script, err := cache.contracts.Script(cache.network, address, symLink)
		if err != nil {
			return nil, err
		}
		return script.Full()
	})
	if err != nil {
		return nil, err
	}
	return item.Value().([]byte), nil
}

// StorageTypeBytes -
func (cache *Cache) StorageTypeBytes(address, symLink string) ([]byte, error) {
	if !bcd.IsContract(address) {
		return nil, nil
	}

	key := fmt.Sprintf("storage:%s", address)
	item, err := cache.Fetch(key, 5*time.Minute, func() (interface{}, error) {
		return cache.contracts.ScriptPart(cache.network, address, symLink, consts.STORAGE)
	})
	if err != nil {
		return nil, err
	}
	return item.Value().([]byte), nil
}

// ProtocolByID -
func (cache *Cache) ProtocolByID(id int64) (protocol.Protocol, error) {
	key := fmt.Sprintf("protocol_id:%d", id)
	item, err := cache.Fetch(key, time.Hour, func() (interface{}, error) {
		return cache.protocols.GetByID(id)
	})
	if err != nil {
		return protocol.Protocol{}, err
	}
	return item.Value().(protocol.Protocol), nil
}

// ProtocolByHash -
func (cache *Cache) ProtocolByHash(hash string) (protocol.Protocol, error) {
	key := fmt.Sprintf("protocol_hash:%s", hash)
	item, err := cache.Fetch(key, time.Hour, func() (interface{}, error) {
		return cache.protocols.Get(cache.network, hash, -1)
	})
	if err != nil {
		return protocol.Protocol{}, err
	}
	return item.Value().(protocol.Protocol), nil
}
