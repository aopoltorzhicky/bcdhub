package models

// SubscriptionRequest -
type SubscriptionRequest struct {
	Address         string
	Network         string
	Alias           string
	Hash            string
	ProjectID       string
	WithSame        bool
	WithSimilar     bool
	WithMempool     bool
	WithMigrations  bool
	WithErrors      bool
	WithCalls       bool
	WithDeployments bool
}

// EventType -
const (
	EventTypeError     = "error"
	EventTypeMigration = "migration"
	EventTypeCall      = "call"
	EventTypeInvoke    = "invoke"
	EventTypeDeploy    = "deploy"
	EventTypeSame      = "same"
	EventTypeSimilar   = "similar"
	EventTypeMempool   = "mempool"
)

// Event -
type Event struct {
	Type    string      `json:"type"`
	Address string      `json:"address"`
	Network string      `json:"network"`
	Alias   string      `json:"alias"`
	Body    interface{} `json:"body,omitempty"`
}

// Repository -
type Repository struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ContractCountStats -
type ContractCountStats struct {
	Total          int64
	SameCount      int64
	TotalWithdrawn int64
	Balance        int64
}
