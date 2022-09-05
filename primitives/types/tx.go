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

	TxTransfer   TransactionType = "transfer"
	TxDelegate   TransactionType = "delegate"
	TxRedelegate TransactionType = "redelegate"
	TxUnbound    TransactionType = "unbound"
	TxReward     TransactionType = "reward"
)

type (
	Block struct {
		Number int64 `json:"number"`
		Txs    []Tx  `json:"txs"`
	}

	Tx struct {
		ID             string          `json:"id" bson:"id"`
		From           string          `json:"from" bson:"from"`
		To             string          `json:"to" bson:"to"`
		BlockCreatedAt int64           `json:"block_created_at" bson:"block_created_at"`
		Block          uint64          `json:"block_num" bson:"block_num"`
		Status         Status          `json:"status" bson:"status"`
		Error          string          `json:"error,omitempty" bson:"error"`
		Type           TransactionType `json:"type" bson:"type"`
		Metadata       interface{}     `json:"metadata" bson:"metadata"`
		GasUsed        string          `json:"gas_used" bson:"gas_used"`
		GasWanted      string          `json:"gas_wanted" bson:"gas_wanted"`
		Fee            Fee             `json:"fee"`
	}

	Fee struct {
		Asset string `json:"asset"`
		Value Amount `json:"value"`
	}

	Txs []Tx

	// Transfer describes the transfer metadata
	Transfer struct {
		FromAddr string `json:"from_address"`
		ToAddr   string `json:"to_address"`
		Asset    string `json:"asset"`
		Value    Amount `json:"value"`
	}

	// Delegate describes the delegate metadata
	Delegate struct {
		DelegatorAddr string `json:"delegator_address"`
		ValidatorAddr string `json:"validator_address"`
		Asset         string `json:"asset"`
		Value         Amount `json:"value"`
	}

	// Redelegate describes the redelegate metadata
	Redelegate struct {
		DelegatorAddr    string `json:"delegator_address"`
		ValidatorSrcAddr string `json:"validator_src_address"`
		ValidatorDstAddr string `json:"validator_dst_address"`
		Asset            string `json:"asset"`
		Value            Amount `json:"value"`
	}

	// Unbound describes the unbound metadata
	Unbound struct {
		DelegatorAddr string `json:"delegator_address"`
		ValidatorAddr string `json:"validator_address"`
		Asset         string `json:"asset"`
		Value         Amount `json:"value"`
	}

	// Reward describes the reward metadata
	Reward struct {
		DelegatorAddr string `json:"delegator_address"`
		ValidatorAddr string `json:"validator_address"`
	}
)
