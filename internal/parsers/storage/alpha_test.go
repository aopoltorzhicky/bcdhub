package storage

import (
	"testing"
	"time"

	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestAlpha_ParseOrigination(t *testing.T) {
	type args struct {
		content   string
		operation operation.Operation
	}
	tests := []struct {
		name    string
		args    args
		want    RichStorage
		wantErr bool
	}{
		{
			name: "mainnet/KT1Fv5xCoUqEeb2TycB7ijXdAXUFH4uPnRNN",
			args: args{
				content: `{"kind":"origination","source":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6","fee":"2790","counter":"9843781","gas_limit":"10726","storage_limit":"1899","balance":"0","script":{"code":[{"prim":"parameter","args":[{"prim":"or","args":[{"prim":"or","args":[{"prim":"or","args":[{"prim":"pair","args":[{"prim":"address","annots":["%spender"]},{"prim":"nat","annots":["%value"]}],"annots":["%approve"]},{"prim":"pair","args":[{"prim":"pair","args":[{"prim":"address","annots":["%owner"]},{"prim":"address","annots":["%spender"]}]},{"prim":"contract","args":[{"prim":"nat"}]}],"annots":["%getAllowance"]}]},{"prim":"or","args":[{"prim":"pair","args":[{"prim":"address","annots":["%owner"]},{"prim":"contract","args":[{"prim":"nat"}]}],"annots":["%getBalance"]},{"prim":"pair","args":[{"prim":"unit"},{"prim":"contract","args":[{"prim":"nat"}]}],"annots":["%getTotalSupply"]}]}]},{"prim":"pair","args":[{"prim":"address","annots":["%from"]},{"prim":"pair","args":[{"prim":"address","annots":["%to"]},{"prim":"nat","annots":["%value"]}]}],"annots":["%transfer"]}]}]},{"prim":"storage","args":[{"prim":"pair","args":[{"prim":"big_map","args":[{"prim":"address"},{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}],"annots":["%allowances"]},{"prim":"nat","annots":["%balance"]}]}],"annots":["%ledger"]},{"prim":"nat","annots":["%totalSupply"]}]}]},{"prim":"code","args":[[{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"LAMBDA","args":[{"prim":"pair","args":[{"prim":"address"},{"prim":"pair","args":[{"prim":"big_map","args":[{"prim":"address"},{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]},{"prim":"nat"}]}]},{"prim":"nat"}]}]},{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]},{"prim":"nat"}]},[{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"GET"},{"prim":"IF_NONE","args":[[{"prim":"PUSH","args":[{"prim":"nat"},{"int":"0"}]},{"prim":"EMPTY_MAP","args":[{"prim":"address"},{"prim":"nat"}]},{"prim":"PAIR"}],[]]}]]},{"prim":"LAMBDA","args":[{"prim":"pair","args":[{"prim":"pair","args":[{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]},{"prim":"nat"}]},{"prim":"address"}]},{"prim":"pair","args":[{"prim":"big_map","args":[{"prim":"address"},{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]},{"prim":"nat"}]}]},{"prim":"nat"}]}]},{"prim":"nat"},[{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"DROP"},{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"GET"},{"prim":"IF_NONE","args":[[{"prim":"PUSH","args":[{"prim":"nat"},{"int":"0"}]}],[]]}]]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"IF_LEFT","args":[[{"prim":"IF_LEFT","args":[[{"prim":"IF_LEFT","args":[[{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"SENDER"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"5"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"4"}]},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"PAIR"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"5"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"PUSH","args":[{"prim":"nat"},{"int":"0"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"5"}]},{"prim":"COMPARE"},{"prim":"GT"},{"prim":"PUSH","args":[{"prim":"nat"},{"int":"0"}]},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"COMPARE"},{"prim":"GT"},{"prim":"AND"},{"prim":"IF","args":[[{"prim":"PUSH","args":[{"prim":"string"},{"string":"UnsafeAllowanceChange"}]},{"prim":"FAILWITH"}],[]]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"4"}]},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"5"}]},{"prim":"DIG","args":[{"int":"5"}]},{"prim":"SWAP"},{"prim":"SOME"},{"prim":"SWAP"},{"prim":"UPDATE"},{"prim":"PAIR"},{"prim":"SOME"},{"prim":"SENDER"},{"prim":"UPDATE"},{"prim":"PAIR"},{"prim":"SWAP"},{"prim":"PAIR"}],[{"prim":"DIG","args":[{"int":"4"}]},{"prim":"DROP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"PAIR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CAR"},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"CAR"},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"5"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"PAIR"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"PUSH","args":[{"prim":"mutez"},{"int":"0"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"TRANSFER_TOKENS"},{"prim":"CONS"},{"prim":"PAIR"}]]}],[{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DROP"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DROP"},{"prim":"IF_LEFT","args":[[{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"SWAP"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"PUSH","args":[{"prim":"mutez"},{"int":"0"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"CDR"},{"prim":"TRANSFER_TOKENS"},{"prim":"CONS"},{"prim":"PAIR"}],[{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DROP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"PUSH","args":[{"prim":"mutez"},{"int":"0"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"CDR"},{"prim":"TRANSFER_TOKENS"},{"prim":"CONS"},{"prim":"PAIR"}]]}]]}],[{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"CDR"},{"prim":"PAIR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"CDR"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"4"}]},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"6"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"7"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"COMPARE"},{"prim":"LT"},{"prim":"IF","args":[[{"prim":"PUSH","args":[{"prim":"string"},{"string":"NotEnoughBalance"}]},{"prim":"FAILWITH"}],[]]},{"prim":"SENDER"},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"5"}]},{"prim":"COMPARE"},{"prim":"NEQ"},{"prim":"IF","args":[[{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"SENDER"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"PAIR"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"6"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"COMPARE"},{"prim":"LT"},{"prim":"IF","args":[[{"prim":"PUSH","args":[{"prim":"string"},{"string":"NotEnoughAllowance"}]},{"prim":"FAILWITH"}],[]]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"4"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"SUB"},{"prim":"ABS"},{"prim":"SOME"},{"prim":"SENDER"},{"prim":"UPDATE"},{"prim":"PAIR"}],[{"prim":"DIG","args":[{"int":"5"}]},{"prim":"DROP"}]]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"SUB"},{"prim":"ABS"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"3"}]},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"SWAP"},{"prim":"SOME"},{"prim":"SWAP"},{"prim":"UPDATE"},{"prim":"PAIR"},{"prim":"DUP"},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"4"}]},{"prim":"PAIR"},{"prim":"DIG","args":[{"int":"4"}]},{"prim":"SWAP"},{"prim":"EXEC"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"ADD"},{"prim":"SWAP"},{"prim":"CAR"},{"prim":"PAIR"},{"prim":"SWAP"},{"prim":"DUP"},{"prim":"DUG","args":[{"int":"2"}]},{"prim":"CDR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"CAR"},{"prim":"DIG","args":[{"int":"2"}]},{"prim":"DIG","args":[{"int":"3"}]},{"prim":"SWAP"},{"prim":"SOME"},{"prim":"SWAP"},{"prim":"UPDATE"},{"prim":"PAIR"},{"prim":"SWAP"},{"prim":"PAIR"}]]}]]}],"storage":{"prim":"Pair","args":[[{"prim":"Elt","args":[{"string":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6"},{"prim":"Pair","args":[[{"prim":"Elt","args":[{"string":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6"},{"int":"1000000"}]}],{"int":"1000000"}]}]}],{"int":"1000000000000"}]}},"metadata":{"balance_updates":[{"kind":"contract","contract":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6","change":"-2790"},{"kind":"freezer","category":"fees","delegate":"tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r","cycle":320,"change":"2790"}],"operation_result":{"status":"applied","big_map_diff":[{"action":"alloc","big_map":"325","key_type":{"prim":"address"},"value_type":{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}],"annots":["%allowances"]},{"prim":"nat","annots":["%balance"]}]}},{"action":"update","big_map":"325","key_hash":"exprudn2kdsp9N7P4ZP6wu22AACpnLE5N1YdDW5zSCqb55fTwSnsdz","key":{"bytes":"0000170b66dca1c7fb751c81d3c66149df164c2e4fe8"},"value":{"prim":"Pair","args":[[{"prim":"Elt","args":[{"bytes":"0000170b66dca1c7fb751c81d3c66149df164c2e4fe8"},{"int":"1000000"}]}],{"int":"1000000"}]}}],"balance_updates":[{"kind":"contract","contract":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6","change":"-405500"},{"kind":"contract","contract":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6","change":"-64250"}],"originated_contracts":["KT1Fv5xCoUqEeb2TycB7ijXdAXUFH4uPnRNN"],"consumed_gas":"10232","consumed_milligas":"10231762","storage_size":"1622","paid_storage_size_diff":"1622"}}}]}]`,
				operation: operation.Operation{
					Level:     1311215,
					Protocol:  "PsDELPH1Kxsxt8f9eWbxQeRxkjfbxoqM52jvs5Y5fBxWWh4ifpo",
					Timestamp: time.Date(2018, 06, 30, 0, 0, 0, 0, time.Local),
					Network:   "mainnet",
					Script:    gjson.Parse(`{"code":[{"prim":"storage","args":[{"prim":"pair","args":[{"prim":"big_map","args":[{"prim":"address"},{"prim":"pair","args":[{"prim":"map","args":[{"prim":"address"},{"prim":"nat"}],"annots":["%allowances"]},{"prim":"nat","annots":["%balance"]}]}],"annots":["%ledger"]},{"prim":"nat","annots":["%totalSupply"]}]}]}`),
				},
			},
			want: RichStorage{
				DeffatedStorage: `{"prim":"Pair","args":[[],{"int":"1000000000000"}]}`,
				Models: []models.Model{
					&bigmapdiff.BigMapDiff{
						Ptr:       -1,
						Address:   "KT1Fv5xCoUqEeb2TycB7ijXdAXUFH4uPnRNN",
						Protocol:  "PsDELPH1Kxsxt8f9eWbxQeRxkjfbxoqM52jvs5Y5fBxWWh4ifpo",
						Timestamp: time.Date(2018, 06, 30, 0, 0, 0, 0, time.Local),
						Level:     1311215,
						KeyHash:   "exprudn2kdsp9N7P4ZP6wu22AACpnLE5N1YdDW5zSCqb55fTwSnsdz",
						Network:   "mainnet",
						Key:       []byte(`{"string":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6"}`),
						Value:     []byte(`{"prim":"Pair","args":[[{"prim":"Elt","args":[{"string":"tz1Mjstk27ppU7SH8eQHh8HU9wrg6dwvoFd6"},{"int":"1000000"}]}],{"int":"1000000"}]}`),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAlpha()

			content := gjson.Parse(tt.args.content)
			got, err := a.ParseOrigination(content, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("Alpha.ParseOrigination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Empty, got.Empty)
			assert.Equal(t, tt.want.DeffatedStorage, got.DeffatedStorage)
			assert.Len(t, got.Models, len(tt.want.Models))

			for i := range tt.want.Models {
				bmd := got.Models[i].(*bigmapdiff.BigMapDiff)
				newBmd := tt.want.Models[i].(*bigmapdiff.BigMapDiff)
				newBmd.ID = bmd.ID
				newBmd.IndexedTime = bmd.IndexedTime
			}
			assert.Equal(t, tt.want.Models, got.Models)
		})
	}
}
