package payloads

type CustomerWallet struct {
	Customer                  *Customer         `json:"customer,omitempty"`
	ID                        string            `json:"id,omitempty"`
	Object                    string            `json:"object,omitempty"`
	ObjectSecret              string            `json:"object_secret,omitempty"`
	SupportedCryptocurrencies []*CryptoCurrency `json:"supported_cryptocurrencies"`
	WalletAddresses           []*WalletAddress  `json:"wallet_addresses"`
}
