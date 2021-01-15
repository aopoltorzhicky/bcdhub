package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestTokenMetadata_Parse(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
		want    *TokenMetadata
	}{
		{
			name:    "test 1",
			value:   `{"prim":"Pair","args":[{"int":"0"},[{"prim":"Elt","args":[{"string":""},{"bytes":"697066733a2f2f516d543633634b35584a6943645047436a586a526162634167646a787875714d4d67553679416d4c6e6178455a35"}]}]]}`,
			wantErr: false,
			want: &TokenMetadata{
				Link:    "ipfs://QmT63cK5XJiCdPGCjXjRabcAgdjxxuqMMgU6yAmLnaxEZ5",
				TokenID: 0,
				Extras:  make(map[string]interface{}),
			},
		}, {
			name:    "test 2",
			value:   `{"prim":"Pair","args":[{"int":"1"},[{"prim":"Elt","args":[{"string":"decimals"},{"bytes":"36"}]},{"prim":"Elt","args":[{"string":"name"},{"bytes":"4e616d65"}]},{"prim":"Elt","args":[{"string":"symbol"},{"bytes":"534d42"}]}]]}`,
			wantErr: false,
			want: &TokenMetadata{
				TokenID:  1,
				Decimals: getIntPtr(36),
				Name:     "Name",
				Symbol:   "SMB",
				Extras:   make(map[string]interface{}),
			},
		}, {
			name:    "test 3",
			value:   `{"prim":"Pair","args":[{"int":"2"},[{"prim":"Elt","args":[{"string":""},{"bytes":"74657a6f732d73746f726167653a636f6e74656e74"}]},{"prim":"Elt","args":[{"string":"content"},{"bytes":"7b226e616d65223a20224e616d65222c202273796d626f6c223a2022534d42222c2022646563696d616c73223a20367d"}]}]]}`,
			wantErr: false,
			want: &TokenMetadata{
				TokenID: 2,
				Extras: map[string]interface{}{
					"content": "7b226e616d65223a20224e616d65222c202273796d626f6c223a2022534d42222c2022646563696d616c73223a20367d",
				},
				Link: "tezos-storage:content",
			},
		}, {
			name:    "test 4: invalid prim",
			value:   `{"prim":"list","args":[{"int":"2"},[{"prim":"Elt","args":[{"string":""},{"bytes":"74657a6f732d73746f726167653a636f6e74656e74"}]},{"prim":"Elt","args":[{"string":"content"},{"bytes":"7b226e616d65223a20224e616d65222c202273796d626f6c223a2022534d42222c2022646563696d616c73223a20367d"}]}]]}`,
			wantErr: true,
			want:    &TokenMetadata{},
		}, {
			name:    "test 5: invalid token ID",
			value:   `{"prim":"Pair","args":[{"string":"2"},[{"prim":"Elt","args":[{"string":""},{"bytes":"74657a6f732d73746f726167653a636f6e74656e74"}]},{"prim":"Elt","args":[{"string":"content"},{"bytes":"7b226e616d65223a20224e616d65222c202273796d626f6c223a2022534d42222c2022646563696d616c73223a20367d"}]}]]}`,
			wantErr: true,
			want:    &TokenMetadata{},
		}, {
			name:    "test 6: invalid metadata map",
			value:   `{"prim":"Pair","args":[{"int":"2"},{"prim":"Elt","args":[{"string":""},{"bytes":"74657a6f732d73746f726167653a636f6e74656e74"}]}]}`,
			wantErr: true,
			want:    &TokenMetadata{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := gjson.Parse(tt.value)

			m := &TokenMetadata{}
			if err := m.Parse(value, "", 0); (err != nil) != tt.wantErr {
				t.Errorf("TokenMetadataParser.parseMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, m, tt.want) {
				t.Errorf("assert")
			}
		})
	}
}

func getIntPtr(value int64) *int64 {
	return &value
}

func TestTokenMetadata_Merge(t *testing.T) {
	tests := []struct {
		name   string
		one    *TokenMetadata
		second *TokenMetadata
		want   *TokenMetadata
	}{
		{
			name: "test 1",
			one:  &TokenMetadata{},
			second: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  10,
			},
			want: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
			},
		}, {
			name:   "test 2",
			second: &TokenMetadata{},
			one: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  10,
			},
			want: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  10,
			},
		}, {
			name: "test 2",
			one: &TokenMetadata{
				Symbol:   "symbol old",
				Name:     "name old",
				Decimals: getIntPtr(9),
				TokenID:  11,
			},
			second: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  10,
			},
			want: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  11,
			},
		}, {
			name: "test 2",
			one: &TokenMetadata{
				Symbol:   "symbol old",
				Name:     "name old",
				Decimals: getIntPtr(9),
				TokenID:  11,
				Extras: map[string]interface{}{
					"test": "1234",
					"a":    "234",
				},
			},
			second: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  10,
				Extras: map[string]interface{}{
					"test": "12345",
					"b":    "234",
				},
			},
			want: &TokenMetadata{
				Symbol:   "symbol",
				Name:     "name",
				Decimals: getIntPtr(10),
				TokenID:  11,
				Extras: map[string]interface{}{
					"test": "12345",
					"a":    "234",
					"b":    "234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.one.Merge(tt.second)
			if !assert.Equal(t, tt.one, tt.want) {
				t.Errorf("assert")
			}
		})
	}
}