package payloads

type Payment struct {
	Amount                       string                         `json:"amount,omitempty"`
	AmountDue                    string                         `json:"amount_due,omitempty"`
	AmountPaid                   string                         `json:"amount_paid,omitempty"`
	BlockchainTransactionDetails []*BlockchainTransactionDetail `json:"blockchain_transaction_details"`
	CanceledAt                   int64                          `json:"canceled_at,omitempty"`
	CreatedAt                    int64                          `json:"created_at,omitempty"`
	Currency                     string                         `json:"currency,omitempty"`
	CurrencyType                 string                         `json:"currency_type,omitempty"`
	CurrencyDue                  string                         `json:"currency_due,omitempty"`
	CurrencyDueAddress           string                         `json:"currency_due_address,omitempty" encryptAbleField:"true"`
	CurrencyDueNetwork           string                         `json:"currency_due_network,omitempty"`
	CurrencyPaid                 string                         `json:"currency_paid,omitempty"`
	Customer                     *Customer                      `json:"customer,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	ExpiresAt                    int64                          `json:"expires_at,omitempty"`
	FailedReason                 string                         `json:"failed_reason,omitempty"`
	ID                           string                         `json:"id,omitempty"`
	ObjSecret                    string                         `json:"obj_secret,omitempty"`
	Object                       string                         `json:"object,omitempty"`
	PriceExpiresAt               int64                          `json:"price_expires_at,omitempty"`
	AddressExpiresAt             int64                          `json:"address_expires_at,omitempty"`
	Status                       string                         `json:"status,omitempty"`
	SupportedCryptocurrencies    []*CryptoCurrency              `json:"supported_cryptocurrencies"`
}
