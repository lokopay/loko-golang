package payloads

type CryptoCurrency struct {
	Amount      string  `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Description string  `json:"description,omitempty"`
	ID          string  `json:"id,omitempty"`
	Network     string  `json:"network,omitempty"`
	ChainId     string  `json:"chain_id,omitempty"`
	Price       float64 `json:"price,omitempty"`
	PricePair   string  `json:"price_pair,omitempty"`
}
