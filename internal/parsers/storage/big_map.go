package storage

import (
	"github.com/baking-bad/bcdhub/internal/bcd/ast"
	"github.com/baking-bad/bcdhub/internal/fetch"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/pkg/errors"
)

// errors -
var (
	ErrBigMapNotFound = errors.New("Big map is not found")
)

// GetBigMapPtr -
func GetBigMapPtr(rpc noderpc.INode, address, key, network, protocol, sharePath string, level int64) (int64, error) {
	data, err := fetch.Contract(address, network, protocol, sharePath)
	if err != nil {
		return 0, err
	}
	script, err := ast.NewScript(data)
	if err != nil {
		return 0, err
	}
	storage, err := script.StorageType()
	if err != nil {
		return 0, err
	}

	node := storage.FindByName(key, false)
	if node == nil {
		return 0, errors.Wrap(ErrBigMapNotFound, key)
	}

	storageJSON, err := rpc.GetScriptStorageRaw(address, level)
	if err != nil {
		return 0, err
	}
	var storageData ast.UntypedAST
	if err := json.Unmarshal(storageJSON, &storageData); err != nil {
		return 0, err
	}
	if err := storage.Settle(storageData); err != nil {
		return 0, err
	}

	if bm, ok := node.(*ast.BigMap); ok {
		return *bm.Ptr, nil
	}

	return 0, errors.Wrap(ErrBigMapNotFound, key)
}

// FindByName -
func FindByName(address, key, network, protocol, sharePath string) *ast.BigMap {
	data, err := fetch.Contract(address, network, protocol, sharePath)
	if err != nil {
		return nil
	}

	script, err := ast.NewScript(data)
	if err != nil {
		return nil
	}

	storage, err := script.StorageType()
	if err != nil {
		return nil
	}

	node := storage.FindByName(key, false)
	if node == nil {
		return nil
	}

	if bm, ok := node.(*ast.BigMap); ok {
		return bm
	}

	return nil
}
