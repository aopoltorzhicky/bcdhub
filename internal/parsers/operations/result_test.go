package operations

import (
	"reflect"
	"testing"

	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
)

func Test_parseOperationResult(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     operation.Operation
	}{
		{
			name:     "test 1",
			fileName: "./data/result/test1.json",
			want: operation.Operation{
				Status:      types.OperationStatusApplied,
				ConsumedGas: 1020700,
			},
		}, {
			name:     "test 2",
			fileName: "./data/result/test2.json",
			want: operation.Operation{
				Status:      types.OperationStatusApplied,
				ConsumedGas: 1020700,
			},
		}, {
			name:     "test 3",
			fileName: "./data/result/test3.json",
			want: operation.Operation{
				Status:              types.OperationStatusApplied,
				ConsumedGas:         1555500,
				StorageSize:         232,
				PaidStorageSizeDiff: 232,
				Destination: account.Account{
					Address: "KT1FVhijNC7ZBL5EjcetiKddDQ2n98t8w4jo",
					Type:    types.AccountTypeContract,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var op noderpc.Operation
			if err := readJSONFile(tt.fileName, &op); err != nil {
				t.Errorf(`readJSONFile("%s") = error %v`, tt.fileName, err)
				return
			}

			var res operation.Operation
			parseOperationResult(op, &res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Result.Parse() = %v, want %v", res, tt.want)
			}
		})
	}
}
