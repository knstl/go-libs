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

	TxSend                    TransactionType = "send"
	TxDelegate                TransactionType = "delegate"
	TxBeginRedelegate         TransactionType = "begin_redelegate"
	TxUndelegate              TransactionType = "undelegate"
	TxWithdrawDelegatorReward TransactionType = "withdraw_delegator_reward"
	TxVote                    TransactionType = "vote"
	TxCreateValidator         TransactionType = "create_validator"
	TxContractCall            TransactionType = "contract_call"
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
		Memo           string          `json:"memo,omitempty"`
		Fee            Fee             `json:"fee"`
	}

	Fee struct {
		Asset string `json:"asset"`
		Value Amount `json:"value"`
	}

	Txs []Tx
)

type (
	TransferMetadata struct {
		Asset string `json:"asset"`
		Value Amount `json:"value"`
	}

	ContractCallMetadata struct {
		Asset   string `json:"asset"`
		Value   Amount `json:"value"`
		Address string `json:"address"`
	}
)
