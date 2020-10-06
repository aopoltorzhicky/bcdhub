package elastic

// Document names
const (
	DocContracts     = "contract"
	DocBlocks        = "block"
	DocOperations    = "operation"
	DocBigMapDiff    = "bigmapdiff"
	DocBigMapActions = "bigmapaction"
	DocMetadata      = "metadata"
	DocMigrations    = "migration"
	DocProtocol      = "protocol"
	DocTransfers     = "transfer"
	DocTokenMetadata = "token_metadata"
	DocTZIP          = "tzip"
)

// Index names
const (
	IndexName = "bcd"
)

// Errors
const (
	IndexNotFoundError = "index_not_found_exception"
	RecordNotFound     = "Record is not found:"
)
