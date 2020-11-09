package models

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/contractparser/cerrors"
	"github.com/baking-bad/bcdhub/internal/contractparser/consts"
	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/models/utils"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Operation -
type Operation struct {
	ID string `json:"-"`

	IndexedTime  int64 `json:"indexed_time"`
	ContentIndex int64 `json:"content_index,omitempty"`

	Network  string `json:"network"`
	Protocol string `json:"protocol"`
	Hash     string `json:"hash"`
	Internal bool   `json:"internal"`
	Nonce    *int64 `json:"nonce,omitempty"`

	Status           string    `json:"status"`
	Timestamp        time.Time `json:"timestamp"`
	Level            int64     `json:"level"`
	Kind             string    `json:"kind"`
	Initiator        string    `json:"initiator"`
	Source           string    `json:"source"`
	Fee              int64     `json:"fee,omitempty"`
	Counter          int64     `json:"counter,omitempty"`
	GasLimit         int64     `json:"gas_limit,omitempty"`
	StorageLimit     int64     `json:"storage_limit,omitempty"`
	Amount           int64     `json:"amount,omitempty"`
	Destination      string    `json:"destination,omitempty"`
	PublicKey        string    `json:"public_key,omitempty"`
	ManagerPubKey    string    `json:"manager_pubkey,omitempty"`
	Delegate         string    `json:"delegate,omitempty"`
	Parameters       string    `json:"parameters,omitempty"`
	FoundBy          string    `json:"found_by,omitempty"`
	Entrypoint       string    `json:"entrypoint,omitempty"`
	SourceAlias      string    `json:"source_alias,omitempty"`
	DestinationAlias string    `json:"destination_alias,omitempty"`

	Result                             *OperationResult `json:"result,omitempty"`
	Errors                             []*cerrors.Error `json:"errors,omitempty"`
	Burned                             int64            `json:"burned,omitempty"`
	AllocatedDestinationContractBurned int64            `json:"allocated_destination_contract_burned,omitempty"`

	DeffatedStorage string       `json:"deffated_storage"`
	Script          gjson.Result `json:"-"`

	DelegateAlias string `json:"delegate_alias,omitempty"`

	ParameterStrings []string `json:"parameter_strings,omitempty"`
	StorageStrings   []string `json:"storage_strings,omitempty"`
	Tags             []string `json:"tags,omitempty"`
}

// GetID -
func (o *Operation) GetID() string {
	return o.ID
}

// GetIndex -
func (o *Operation) GetIndex() string {
	return "operation"
}

// GetQueues -
func (o *Operation) GetQueues() []string {
	return []string{"operations"}
}

// MarshalToQueue -
func (o *Operation) MarshalToQueue() ([]byte, error) {
	return []byte(o.ID), nil
}

// LogFields -
func (o *Operation) LogFields() logrus.Fields {
	return logrus.Fields{
		"network": o.Network,
		"hash":    o.Hash,
		"block":   o.Level,
	}
}

// GetScores -
func (o *Operation) GetScores(search string) []string {
	return []string{
		"entrypoint^8",
		"parameter_strings^7",
		"storage_strings^7",
		"errors.with^6",
		"errors.id^5",
		"source_alias^3",
		"hash",
		"source",
	}
}

// FoundByName -
func (o *Operation) FoundByName(hit gjson.Result) string {
	keys := hit.Get("highlight").Map()
	categories := o.GetScores("")
	return utils.GetFoundBy(keys, categories)
}

// SetAllocationBurn -
func (o *Operation) SetAllocationBurn(constants Constants) {
	o.AllocatedDestinationContractBurned = 257 * constants.CostPerByte
}

// SetBurned -
func (o *Operation) SetBurned(constants Constants) {
	if o.Status != consts.Applied {
		return
	}

	if o.Result == nil {
		return
	}

	var burned int64

	if o.Result.PaidStorageSizeDiff != 0 {
		burned += o.Result.PaidStorageSizeDiff * constants.CostPerByte
	}

	if o.Result.AllocatedDestinationContract {
		o.SetAllocationBurn(constants)
		burned += o.AllocatedDestinationContractBurned
	}

	o.Burned = burned
}

// IsOrigination -
func (o *Operation) IsOrigination() bool {
	return o.Kind == consts.Origination || o.Kind == consts.OriginationNew
}

// IsTransaction -
func (o *Operation) IsTransaction() bool {
	return o.Kind == consts.Transaction
}

// IsApplied -
func (o *Operation) IsApplied() bool {
	return o.Status == consts.Applied
}

// IsCall -
func (o *Operation) IsCall() bool {
	return helpers.IsContract(o.Destination)
}

// OperationResult -
type OperationResult struct {
	Status                       string           `json:"-"`
	ConsumedGas                  int64            `json:"consumed_gas,omitempty"`
	StorageSize                  int64            `json:"storage_size,omitempty"`
	PaidStorageSizeDiff          int64            `json:"paid_storage_size_diff,omitempty"`
	AllocatedDestinationContract bool             `json:"allocated_destination_contract,omitempty"`
	Originated                   string           `json:"-"`
	Errors                       []*cerrors.Error `json:"-"`
}
