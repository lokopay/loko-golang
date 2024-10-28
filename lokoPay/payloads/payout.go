package payloads

type Payout struct {
	Amount                       string                         `json:"amount,omitempty"`
	BlockchainTransactionDetails []*BlockchainTransactionDetail `json:"blockchain_transaction_details"`
	CanceledAt                   int64                          `json:"canceled_at,omitempty"`
	CreatedAt                    int64                          `json:"created_at,omitempty"`
	Currency                     string                         `json:"currency,omitempty"`
	Customer                     *Customer                      `json:"customer,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	DestinationNetworkDetails    []*DestinationNetworkDetail    `json:"destination_network_details"`
	FailedReason                 string                         `json:"failed_reason,omitempty"`
	ID                           string                         `json:"id,omitempty"`
	ObjSecret                    string                         `json:"obj_secret,omitempty"`
	Object                       string                         `json:"object,omitempty"`
	Status                       string                         `json:"status,omitempty"`
	TransferWithNativeToken      *TransferWithNativeToken       `json:"transfer_with_native_token,omitempty"`
}
