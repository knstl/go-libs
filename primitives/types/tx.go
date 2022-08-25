package types

// Transaction fields
type (
	Status          string
	TransactionType string
	Amount          string
)

const (
	StatusCompleted Status = "completed"
	StatusError     Status = "error"

	TxTransfer TransactionType = "transfer"
)

type (
	Block struct {
		Number int64 `json:"number"`
		Txs    []Tx  `json:"txs"`
	}

	Tx struct {
		// Unique identifier
		ID string `json:"id" bson:"id"`
		// Address of the transaction sender
		From string `json:"from" bson:"from"`
		// Address of the transaction recipient
		To string `json:"to" bson:"to"`
		// Unix timestamp of the block the transaction was included in
		BlockCreatedAt int64 `json:"block_created_at" bson:"block_created_at"`
		// Height of the block the transaction was included in
		Block uint64 `json:"block_num" bson:"block_num"`
		// Status of the transaction e.g: "completed", "pending", "error"
		Status Status `json:"status" bson:"status"`
		// Empty if the transaction "completed" or "pending", else error explaining why the transaction failed (optional)
		Error string `json:"error,omitempty" bson:"error"`
		// Type of metadata
		Type TransactionType `json:"type" bson:"type"`
		// Metadata data object
		Metadata interface{} `json:"metadata" bson:"metadata"`
	}

	Txs []Tx

	// Transfer describes the transfer of currency
	Transfer struct {
		Asset string `json:"asset"`
		Value Amount `json:"value"`
	}
)
