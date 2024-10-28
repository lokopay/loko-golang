package payloads

type WalletAddress struct {
	Address         string `json:"address,omitempty" encryptAbleField:"true"`
	AddressCurrency string `json:"address_currency,omitempty"`
	AddressNetwork  string `json:"address_network,omitempty"`
	Description     string `json:"description,omitempty"`
	ID              string `json:"id,omitempty"`
}
