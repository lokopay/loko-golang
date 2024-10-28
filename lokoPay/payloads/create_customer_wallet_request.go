package payloads

type CreateCustomerWalletRequest struct {
	Currency string    `json:"currency,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Network  string    `json:"network,omitempty"`
}
