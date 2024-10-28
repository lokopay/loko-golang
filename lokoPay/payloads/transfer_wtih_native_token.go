package payloads

type TransferWithNativeToken struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
	Network  string `json:"network,omitempty"`
}
