package tokenbalance

import (
	"math/big"

	"github.com/baking-bad/bcdhub/internal/bcd/ast"
	"github.com/baking-bad/bcdhub/internal/bcd/forge"
	"github.com/baking-bad/bcdhub/internal/bcd/types"
)

// SingleAsset -
type SingleAsset struct {
	ReturnType *ast.TypedAst
}

// NewSingleAssetBalance -
func NewSingleAssetBalance() SingleAsset {
	node, _ := ast.NewTypedAstFromString(`{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]}`)
	return SingleAsset{
		ReturnType: node,
	}
}

// NewSingleAssetUpdate -
func NewSingleAssetUpdate() SingleAsset {
	node, _ := ast.NewTypedAstFromString(`{"prim":"map","args":[{"prim":"address"},{"prim":"int"}]}`)
	return SingleAsset{
		ReturnType: node,
	}
}

// GetReturnType -
func (p SingleAsset) GetReturnType() *ast.TypedAst {
	return p.ReturnType
}

// Parse -
func (p SingleAsset) Parse(data []byte) ([]TokenBalance, error) {
	m, err := getMap(p.ReturnType, data)
	if err != nil {
		return nil, err
	}

	balances := make([]TokenBalance, 0)
	err = m.Data.Range(func(key, value ast.Comparable) (bool, error) {
		val := value.(ast.Node)
		k := key.(*ast.Address)

		balance := val.GetValue().(*types.BigInt)

		address := forge.DecodeString(k.GetValue().(string))
		if address == "" {
			return false, nil
		}
		amount, _ := new(big.Float).SetInt(balance.Int).Float64()
		balances = append(balances, TokenBalance{
			Value:   amount,
			Address: address,
		})
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return balances, nil
}
