package handlers

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/contractparser/cerrors"
	"github.com/baking-bad/bcdhub/internal/contractparser/docstring"
	"github.com/baking-bad/bcdhub/internal/contractparser/formatter"
	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/baking-bad/bcdhub/internal/jsonschema"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/tidwall/gjson"
)

// Error -
type Error struct {
	Message string `json:"message" example:"text"`
}

// Operation -
type Operation struct {
	Level                              int64            `json:"level,omitempty"`
	Fee                                int64            `json:"fee,omitempty"`
	Counter                            int64            `json:"counter,omitempty"`
	GasLimit                           int64            `json:"gas_limit,omitempty"`
	StorageLimit                       int64            `json:"storage_limit,omitempty"`
	Amount                             int64            `json:"amount,omitempty"`
	Balance                            int64            `json:"balance,omitempty"`
	Burned                             int64            `json:"burned,omitempty"`
	AllocatedDestinationContractBurned int64            `json:"allocated_destination_contract_burned,omitempty"`
	IndexedTime                        int64            `json:"-"`
	ContentIndex                       int64            `json:"content_index"`
	Errors                             []*cerrors.Error `json:"errors,omitempty"`
	Result                             *OperationResult `json:"result,omitempty"`
	Parameters                         interface{}      `json:"parameters,omitempty"`
	StorageDiff                        interface{}      `json:"storage_diff,omitempty"`
	RawMempool                         interface{}      `json:"rawMempool,omitempty"`
	Timestamp                          time.Time        `json:"timestamp"`
	ID                                 string           `json:"id,omitempty"`
	Protocol                           string           `json:"protocol"`
	Hash                               string           `json:"hash,omitempty"`
	Network                            string           `json:"network"`
	Kind                               string           `json:"kind"`
	Source                             string           `json:"source,omitempty"`
	SourceAlias                        string           `json:"source_alias,omitempty"`
	Destination                        string           `json:"destination,omitempty"`
	DestinationAlias                   string           `json:"destination_alias,omitempty"`
	PublicKey                          string           `json:"public_key,omitempty"`
	ManagerPubKey                      string           `json:"manager_pubkey,omitempty"`
	Delegate                           string           `json:"delegate,omitempty"`
	Status                             string           `json:"status"`
	Entrypoint                         string           `json:"entrypoint,omitempty"`
	Internal                           bool             `json:"internal"`
	Mempool                            bool             `json:"mempool"`
}

// ParseJSON -
func (o *Operation) ParseJSON(raw gjson.Result) {
	o.Status = raw.Get("status").String()
	o.Kind = raw.Get("kind").String()
	o.Source = raw.Get("source").String()
	o.Fee = raw.Get("fee").Int()
	o.Counter = raw.Get("counter").Int()
	o.GasLimit = raw.Get("gas_limit").Int()
	o.StorageLimit = raw.Get("storage_limit").Int()
	o.Amount = raw.Get("amount").Int()
	o.Destination = raw.Get("destination").String()
	o.PublicKey = raw.Get("public_key").String()
	o.ManagerPubKey = raw.Get("manager_pubkey").String()
	o.Delegate = raw.Get("delegate").String()
}

// FromModel -
func (o *Operation) FromModel(operation models.Operation) {
	o.ID = operation.ID
	o.Protocol = operation.Protocol
	o.Hash = operation.Hash
	o.Network = operation.Network
	o.Internal = operation.Internal
	o.Timestamp = operation.Timestamp

	o.Level = operation.Level
	o.Kind = operation.Kind
	o.Source = operation.Source
	o.SourceAlias = operation.SourceAlias
	o.Fee = operation.Fee
	o.Counter = operation.Counter
	o.GasLimit = operation.GasLimit
	o.StorageLimit = operation.StorageLimit
	o.Amount = operation.Amount
	o.Destination = operation.Destination
	o.DestinationAlias = operation.DestinationAlias
	o.PublicKey = operation.PublicKey
	o.ManagerPubKey = operation.ManagerPubKey
	o.Delegate = operation.Delegate
	o.Status = operation.Status
	o.Burned = operation.Burned
	o.Entrypoint = operation.Entrypoint
	o.IndexedTime = operation.IndexedTime
	o.ContentIndex = operation.ContentIndex
	o.AllocatedDestinationContractBurned = operation.AllocatedDestinationContractBurned
}

// ToModel -
func (o *Operation) ToModel() models.Operation {
	var result *models.OperationResult
	if o.Result != nil {
		result = o.Result.ToModel()
	}
	return models.Operation{
		ID:        o.ID,
		Protocol:  o.Protocol,
		Hash:      o.Hash,
		Network:   o.Network,
		Internal:  o.Internal,
		Timestamp: o.Timestamp,

		Level:            o.Level,
		Kind:             o.Kind,
		Source:           o.Source,
		SourceAlias:      o.SourceAlias,
		Fee:              o.Fee,
		Counter:          o.Counter,
		GasLimit:         o.GasLimit,
		StorageLimit:     o.StorageLimit,
		Amount:           o.Amount,
		Destination:      o.Destination,
		DestinationAlias: o.DestinationAlias,
		PublicKey:        o.PublicKey,
		ManagerPubKey:    o.ManagerPubKey,
		Delegate:         o.Delegate,
		Status:           o.Status,
		Burned:           o.Burned,
		Entrypoint:       o.Entrypoint,
		IndexedTime:      o.IndexedTime,

		Result: result,
	}
}

// OperationResult -
type OperationResult struct {
	ConsumedGas                  int64 `json:"consumed_gas,omitempty" example:"100"`
	StorageSize                  int64 `json:"storage_size,omitempty" example:"200"`
	PaidStorageSizeDiff          int64 `json:"paid_storage_size_diff,omitempty" example:"300"`
	AllocatedDestinationContract bool  `json:"allocated_destination_contract,omitempty" example:"true"`
}

// FromModel -
func (r *OperationResult) FromModel(result *models.OperationResult) {
	if result == nil || r == nil {
		return
	}
	r.AllocatedDestinationContract = result.AllocatedDestinationContract
	r.ConsumedGas = result.ConsumedGas
	r.PaidStorageSizeDiff = result.PaidStorageSizeDiff
	r.StorageSize = result.StorageSize
}

// ToModel -
func (r *OperationResult) ToModel() *models.OperationResult {
	if r == nil {
		return nil
	}

	return &models.OperationResult{
		AllocatedDestinationContract: r.AllocatedDestinationContract,
		ConsumedGas:                  r.ConsumedGas,
		PaidStorageSizeDiff:          r.PaidStorageSizeDiff,
		StorageSize:                  r.StorageSize,
	}
}

// Contract -
type Contract struct {
	ID        string    `json:"id"`
	Network   string    `json:"network"`
	Level     int64     `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Language  string    `json:"language,omitempty"`

	Hash        string   `json:"hash"`
	Tags        []string `json:"tags,omitempty"`
	Hardcoded   []string `json:"hardcoded,omitempty"`
	FailStrings []string `json:"fail_strings,omitempty"`
	Annotations []string `json:"annotations,omitempty"`
	Entrypoints []string `json:"entrypoints,omitempty"`

	Address  string `json:"address"`
	Manager  string `json:"manager,omitempty"`
	Delegate string `json:"delegate,omitempty"`

	ProjectID       string    `json:"project_id,omitempty"`
	FoundBy         string    `json:"found_by,omitempty"`
	LastAction      time.Time `json:"last_action,omitempty"`
	TxCount         int64     `json:"tx_count,omitempty"`
	MigrationsCount int64     `json:"migrations_count,omitempty"`
	Alias           string    `json:"alias,omitempty"`
	DelegateAlias   string    `json:"delegate_alias,omitempty"`

	Subscription       *Subscription `json:"subscription,omitempty"`
	TotalSubscribed    int           `json:"total_subscribed"`
	Slug               string        `json:"slug,omitempty"`
	Verified           bool          `json:"verified,omitempty"`
	VerificationSource string        `json:"verification_source,omitempty"`

	Tokens []TokenBalance `json:"tokens"`
}

// FromModel -
func (c *Contract) FromModel(contract models.Contract) {
	c.Address = contract.Address
	c.Alias = contract.Alias
	c.Annotations = contract.Annotations
	c.Delegate = contract.Delegate
	c.DelegateAlias = contract.DelegateAlias
	c.Entrypoints = contract.Entrypoints
	c.FailStrings = contract.FailStrings
	c.FoundBy = contract.FoundBy
	c.Hardcoded = contract.Hardcoded
	c.Hash = contract.Hash
	c.ID = contract.GetID()
	c.Language = contract.Language
	c.TxCount = contract.TxCount
	c.LastAction = contract.LastAction

	c.Level = contract.Level
	c.Manager = contract.Manager
	c.MigrationsCount = contract.MigrationsCount
	c.Network = contract.Network
	c.ProjectID = contract.ProjectID
	c.Tags = contract.Tags
	c.Timestamp = contract.Timestamp
	c.Verified = contract.Verified
	c.VerificationSource = contract.VerificationSource
}

// Subscription -
type Subscription struct {
	Address          string    `json:"address"`
	Network          string    `json:"network"`
	Alias            string    `json:"alias,omitempty"`
	SubscribedAt     time.Time `json:"subscribed_at"`
	WatchSame        bool      `json:"watch_same"`
	WatchSimilar     bool      `json:"watch_similar"`
	WatchMempool     bool      `json:"watch_mempool"`
	WatchMigrations  bool      `json:"watch_migrations"`
	WatchDeployments bool      `json:"watch_deployments"`
	WatchCalls       bool      `json:"watch_calls"`
	WatchErrors      bool      `json:"watch_errors"`
	SentryEnabled    bool      `json:"sentry_enabled"`
	SentryDSN        string    `json:"sentry_dsn,omitempty"`
}

// Event -
type Event struct {
	Event string    `json:"event"`
	Date  time.Time `json:"date"`
}

// OperationResponse -
type OperationResponse struct {
	Operations []Operation `json:"operations"`
	LastID     string      `json:"last_id,omitempty" example:"1588640276994159"`
}

type userProfile struct {
	Login            string    `json:"login"`
	AvatarURL        string    `json:"avatar_url"`
	MarkReadAt       time.Time `json:"mark_read_at"`
	RegisteredAt     time.Time `json:"registered_at"`
	MarkedContracts  int       `json:"marked_contracts"`
	CompilationTasks int64     `json:"compilation_tasks"`
	Verifications    int64     `json:"verifications"`
	Deployments      int64     `json:"deployments"`

	Subscriptions []Subscription `json:"subscriptions"`
}

// BigMapItem -
type BigMapItem struct {
	Key       interface{} `json:"key"`
	Value     interface{} `json:"value"`
	KeyHash   string      `json:"key_hash"`
	KeyString string      `json:"key_string"`
	Level     int64       `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
}

// BigMapResponseItem -
type BigMapResponseItem struct {
	Item  BigMapItem `json:"data"`
	Count int64      `json:"count"`
}

// GetBigMapResponse -
type GetBigMapResponse struct {
	Address       string              `json:"address"`
	Network       string              `json:"network"`
	Ptr           int64               `json:"ptr"`
	ActiveKeys    uint                `json:"active_keys"`
	TotalKeys     uint                `json:"total_keys"`
	ContractAlias string              `json:"contract_alias,omitempty"`
	Typedef       []docstring.Typedef `json:"typedef,omitempty"`
}

// Migration -
type Migration struct {
	Level        int64     `json:"level"`
	Timestamp    time.Time `json:"timestamp"`
	Hash         string    `json:"hash,omitempty"`
	Protocol     string    `json:"protocol"`
	PrevProtocol string    `json:"prev_protocol"`
	Kind         string    `json:"kind"`
}

// TokenContract -
type TokenContract struct {
	Network       string                      `json:"network"`
	Level         int64                       `json:"level"`
	Timestamp     time.Time                   `json:"timestamp"`
	LastAction    time.Time                   `json:"last_action"`
	Address       string                      `json:"address"`
	Manager       string                      `json:"manager,omitempty"`
	Delegate      string                      `json:"delegate,omitempty"`
	Alias         string                      `json:"alias,omitempty"`
	DelegateAlias string                      `json:"delegate_alias,omitempty"`
	Type          string                      `json:"type"`
	Balance       int64                       `json:"balance"`
	TxCount       int64                       `json:"tx_count"`
	Methods       map[string]TokenMethodStats `json:"methods,omitempty"`
}

// TokenMethodStats -
type TokenMethodStats struct {
	CallCount          int64 `json:"call_count"`
	AverageConsumedGas int64 `json:"average_consumed_gas"`
}

// PageableTokenContracts -
type PageableTokenContracts struct {
	Total  int64           `json:"total"`
	Tokens []TokenContract `json:"tokens"`
}

// TokenTransfer -
type TokenTransfer struct {
	Contract  string    `json:"contract"`
	Network   string    `json:"network"`
	Protocol  string    `json:"protocol"`
	Hash      string    `json:"hash"`
	Counter   int64     `json:"counter,omitempty"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Level     int64     `json:"level"`
	From      string    `json:"from,omitempty"`
	To        string    `json:"to"`
	Amount    int64     `json:"amount"`
	Source    string    `json:"source"`
	Nonce     *int64    `json:"nonce,omitempty"`
}

// PageableTokenTransfers -
type PageableTokenTransfers struct {
	Transfers []TokenTransfer `json:"transfers"`
	LastID    string          `json:"last_id,omitempty"`
}

// BigMapDiffItem -
type BigMapDiffItem struct {
	Value     interface{} `json:"value"`
	Level     int64       `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
}

// BigMapDiffByKeyResponse -
type BigMapDiffByKeyResponse struct {
	Key     interface{}      `json:"key,omitempty"`
	KeyHash string           `json:"key_hash"`
	Values  []BigMapDiffItem `json:"values,omitempty"`
	Total   int64            `json:"total"`
}

// CodeDiffResponse -
type CodeDiffResponse struct {
	Left  CodeDiffLeg          `json:"left"`
	Right CodeDiffLeg          `json:"right"`
	Diff  formatter.DiffResult `json:"diff"`
}

// NetworkStats -
type NetworkStats struct {
	ContractsCount  int64            `json:"contracts_count" example:"10"`
	OperationsCount int64            `json:"operations_count" example:"100"`
	Protocols       []Protocol       `json:"protocols"`
	Languages       map[string]int64 `json:"languages"`
}

// SearchBigMapDiff -
type SearchBigMapDiff struct {
	Ptr       int64     `json:"ptr"`
	Key       string    `json:"key"`
	KeyHash   string    `json:"key_hash"`
	Value     string    `json:"value"`
	Level     int64     `json:"level"`
	Address   string    `json:"address"`
	Network   string    `json:"network"`
	Timestamp time.Time `json:"timestamp"`
	FoundBy   string    `json:"found_by"`
}

// EntrypointSchema ;
type EntrypointSchema struct {
	docstring.EntrypointType
	Schema       jsonschema.Schema       `json:"schema"`
	DefaultModel jsonschema.DefaultModel `json:"default_model,omitempty"`
}

// GetErrorLocationResponse -
type GetErrorLocationResponse struct {
	Text        string `json:"text"`
	FailedRow   int    `json:"failed_row"`
	FirstRow    int    `json:"first_row"`
	StartColumn int    `json:"start_col"`
	EndColumn   int    `json:"end_col"`
}

// Alias -
type Alias struct {
	Alias   string `json:"alias" example:"Contract alias"`
	Network string `json:"network" example:"babylonnet"`
	Address string `json:"address" example:"KT1CeekjGVRc5ASmgWDc658NBExetoKNuiqz"`
	Slug    string `json:"slug" example:"contract_slug"`
}

// FromModel -
func (a *Alias) FromModel(alias *models.TZIP) {
	a.Alias = alias.Name
	a.Address = alias.Address
	a.Network = alias.Network
	a.Slug = alias.Slug
}

// Protocol -
type Protocol struct {
	Hash       string `json:"hash" example:"PsCARTHAGazKbHtnKfLzQg3kms52kSRpgnDY982a9oYsSXRLQEb"`
	Network    string `json:"network" example:"mainnet"`
	StartLevel int64  `json:"start_level" example:"851969"`
	EndLevel   int64  `json:"end_level" example:"0"`
	Alias      string `json:"alias" example:"Carthage"`
}

// FromModel -
func (p *Protocol) FromModel(protocol models.Protocol) {
	p.Hash = protocol.Hash
	p.Network = protocol.Network
	p.StartLevel = protocol.StartLevel
	p.EndLevel = protocol.EndLevel
	p.Alias = protocol.Alias
}

// Block -
type Block struct {
	Network     string    `json:"network" example:"mainnet"`
	Hash        string    `json:"hash" example:"BLyAEwaXShJuZasvUezHUfLqzZ48V8XrPvXF2wRaH15tmzEpsHT"`
	Level       int64     `json:"level" example:"24"`
	Predecessor string    `json:"predecessor" example:"BMWVEwEYw9m5iaHzqxDfkPzZTV4rhkSouRh3DkVMVGkxZ3EVaNs"`
	ChainID     string    `json:"chain_id" example:"NetXdQprcVkpaWU"`
	Protocol    string    `json:"protocol" example:"PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY"`
	Timestamp   time.Time `json:"timestamp" example:"2018-06-30T18:05:27Z"`
}

// FromModel -
func (b *Block) FromModel(block models.Block) {
	b.Network = block.Network
	b.Hash = block.Hash
	b.Level = block.Level
	b.Predecessor = block.Predecessor
	b.ChainID = block.ChainID
	b.Protocol = block.Protocol
	b.Timestamp = block.Timestamp
}

// LightContract -
type LightContract struct {
	Address  string    `json:"address"`
	Network  string    `json:"network"`
	Deployed time.Time `json:"deploy_time"`
}

// FromModel -
func (c *LightContract) FromModel(light elastic.LightContract) {
	c.Address = light.Address
	c.Network = light.Network
	c.Deployed = light.Deployed
}

// SimilarContractsResponse -
type SimilarContractsResponse struct {
	Count     int               `json:"count"`
	Contracts []SimilarContract `json:"contracts"`
}

// SimilarContract -
type SimilarContract struct {
	*Contract
	Count   int64 `json:"count"`
	Added   int64 `json:"added,omitempty"`
	Removed int64 `json:"removed,omitempty"`
}

// FromModel -
func (c *SimilarContract) FromModel(similar elastic.SimilarContract, diff CodeDiffResponse) {
	var contract Contract
	contract.FromModel(*similar.Contract)
	c.Contract = &contract

	c.Count = similar.Count
	c.Added = diff.Diff.Added
	c.Removed = diff.Diff.Removed
}

// SameContractsResponse -
type SameContractsResponse struct {
	Count     int64      `json:"count"`
	Contracts []Contract `json:"contracts"`
}

// FromModel -
func (c *SameContractsResponse) FromModel(same elastic.SameContractsResponse) {
	c.Count = same.Count

	c.Contracts = make([]Contract, len(same.Contracts))
	for i := range same.Contracts {
		var contract Contract
		contract.FromModel(same.Contracts[i])
		c.Contracts[i] = contract
	}
}

// Series -
type Series [][]int64

// BigMapHistoryResponse -
type BigMapHistoryResponse struct {
	Address string              `json:"address"`
	Network string              `json:"network"`
	Ptr     int64               `json:"ptr"`
	Items   []BigMapHistoryItem `json:"items,omitempty"`
}

// BigMapHistoryItem -
type BigMapHistoryItem struct {
	Action         string    `json:"action"`
	SourcePtr      *int64    `json:"source_ptr,omitempty"`
	DestinationPtr *int64    `json:"destination_ptr,omitempty"`
	Timestamp      time.Time `json:"timestamp"`
}

// Transfer -
type Transfer struct {
	*models.Transfer
	Token *TokenMetadata `json:"token,omitempty"`
}

// TransferResponse -
type TransferResponse struct {
	Transfers []Transfer `json:"transfers"`
	Total     int64      `json:"total"`
}

// ConfigResponse -
type ConfigResponse struct {
	Networks      []string          `json:"networks"`
	OauthEnabled  bool              `json:"oauth_enabled"`
	RPCEndpoints  map[string]string `json:"rpc_endpoints"`
	TzKTEndpoints map[string]string `json:"tzkt_endpoints"`
	SentryDSN     string            `json:"sentry_dsn"`
}

// DApp -
type DApp struct {
	Name              string   `json:"name"`
	ShortDescription  string   `json:"short_description"`
	FullDescription   string   `json:"full_description"`
	WebSite           string   `json:"website"`
	Slug              string   `json:"slug,omitempty"`
	AgoraReviewPostID int64    `json:"agora_review_post_id,omitempty"`
	AgoraQAPostID     int64    `json:"agora_qa_post_id,omitempty"`
	Authors           []string `json:"authors"`
	SocialLinks       []string `json:"social_links"`
	Interfaces        []string `json:"interfaces"`
	Categories        []string `json:"categories"`
	Soon              bool     `json:"soon"`
	Logo              string   `json:"logo"`
	Cover             string   `json:"cover,omitempty"`
	Volume24Hours     float64  `json:"volume_24_hours"`

	Screenshots []Screenshot    `json:"screenshots,omitempty"`
	Contracts   []DAppContract  `json:"contracts,omitempty"`
	DexTokens   []TokenMetadata `json:"dex_tokens,omitempty"`
	Tokens      []Token         `json:"tokens,omitempty"`
}

// DAppContract -
type DAppContract struct {
	Network     string    `json:"network"`
	Address     string    `json:"address"`
	Alias       string    `json:"alias,omitempty"`
	ReleaseDate time.Time `json:"release_date"`
}

// Screenshot -
type Screenshot struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

// Token -
type Token struct {
	TokenMetadata
	elastic.TokenSupply
}

// AccountInfo -
type AccountInfo struct {
	Address    string         `json:"address"`
	Network    string         `json:"network"`
	Alias      string         `json:"alias,omitempty"`
	Balance    int64          `json:"balance"`
	TxCount    int64          `json:"tx_count"`
	LastAction time.Time      `json:"last_action"`
	Tokens     []TokenBalance `json:"tokens"`
}

// TokenBalance -
type TokenBalance struct {
	TokenMetadata
	Balance int64 `json:"balance"`
}

// TokenMetadata -
type TokenMetadata struct {
	Contract        string                 `json:"contract"`
	Network         string                 `json:"network"`
	RegistryAddress string                 `json:"registry_address,omitempty"`
	Level           int64                  `json:"level,omitempty"`
	TokenID         int64                  `json:"token_id"`
	Symbol          string                 `json:"symbol,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Decimals        *int64                 `json:"decimals,omitempty"`
	Extras          map[string]interface{} `json:"extras,omitempty"`
	Volume24Hours   float64                `json:"volume_24_hours"`
}

// TokenMetadataFromElasticModel -
func TokenMetadataFromElasticModel(model elastic.TokenMetadata) (tm TokenMetadata) {
	tm.TokenID = model.TokenID
	tm.Symbol = model.Symbol
	tm.Name = model.Name
	tm.RegistryAddress = model.RegistryAddress
	tm.Decimals = model.Decimals
	tm.Contract = model.Address
	tm.Level = model.Level
	tm.Network = model.Network
	tm.Extras = model.Extras
	return
}

// DomainsResponse -
type DomainsResponse struct {
	Domains []models.TezosDomain `json:"domains"`
	Total   int64                `json:"total"`
}

// CountResponse -
type CountResponse struct {
	Count int64 `json:"count"`
}

// MetadataResponse -
type MetadataResponse struct {
	Hash string `json:"hash"`
}
