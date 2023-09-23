package bigmapdiff

import (
	"context"
	"time"

	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/uptrace/bun"
)

// BigMapDiff -
type BigMapDiff struct {
	bun.BaseModel `bun:"big_map_diffs"`

	ID          int64       `bun:"id,pk,notnull,autoincrement"`
	Ptr         int64       `bun:"ptr"`
	Key         types.Bytes `bun:",notnull,type:bytea"`
	KeyHash     string
	Value       types.Bytes `bun:",type:bytea"`
	Level       int64
	Contract    string
	Timestamp   time.Time `bun:",pk"`
	ProtocolID  int64     `bun:",type:SMALLINT"`
	OperationID int64
}

func (BigMapDiff) PartitionBy() string {
	return "RANGE(timestamp)"
}

// GetID -
func (b *BigMapDiff) GetID() int64 {
	return b.ID
}

// GetIndex -
func (b *BigMapDiff) GetIndex() string {
	return "big_map_diffs"
}

// Save -
func (b *BigMapDiff) Save(ctx context.Context, tx bun.IDB) error {
	_, err := tx.NewInsert().Model(b).
		On("CONFLICT (id, timestamp) DO UPDATE").
		Set("ptr = excluded.ptr").
		Set("key = excluded.key").
		Set("key_hash = excluded.key_hash").
		Set("value = excluded.value").
		Set("level = excluded.level").
		Set("contract = excluded.contract").
		Set("timestamp = excluded.timestamp").
		Set("protocol_id = excluded.protocol_id").
		Set("operation_id = excluded.operation_id").
		Returning("id").
		Exec(ctx)
	return err
}

// LogFields -
func (b *BigMapDiff) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"contract": b.Contract,
		"ptr":      b.Ptr,
		"block":    b.Level,
		"key_hash": b.KeyHash,
	}
}

// KeyBytes -
func (b *BigMapDiff) KeyBytes() []byte {
	if len(b.Key) >= 2 {
		if b.Key[0] == 34 && b.Key[len(b.Key)-1] == 34 {
			return b.Key[1 : len(b.Key)-1]
		}
	}
	return b.Key
}

// ValueBytes -
func (b *BigMapDiff) ValueBytes() []byte {
	if len(b.Value) >= 2 {
		if b.Value[0] == 34 && b.Value[len(b.Value)-1] == 34 {
			return b.Value[1 : len(b.Value)-1]
		}
	}
	return b.Value
}

// ToState -
func (b *BigMapDiff) ToState() *BigMapState {
	state := &BigMapState{
		Contract:        b.Contract,
		Ptr:             b.Ptr,
		LastUpdateLevel: b.Level,
		LastUpdateTime:  b.Timestamp,
		KeyHash:         b.KeyHash,
		Key:             b.KeyBytes(),
	}

	val := b.ValueBytes()
	if len(val) == 0 {
		state.Removed = true
	} else {
		state.Value = val
	}

	return state
}
