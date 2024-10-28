package payloads

type BlockchainTransactionDetail struct {
	Address       string `json:"address,omitempty" encryptAbleField:"true"`
	Amount        string `json:"amount,omitempty"`
	BlockHeight   int64  `json:"block_height,omitempty"`
	BlockTime     int64  `json:"block_time,omitempty"`
	Confirmations int64  `json:"confirmations,omitempty"`
	Currency      string `json:"currency,omitempty"`
	ID            string `json:"id,omitempty"`
	Network       string `json:"network,omitempty"`
	TxHash        string `json:"tx_hash,omitempty" encryptAbleField:"true"`
}
