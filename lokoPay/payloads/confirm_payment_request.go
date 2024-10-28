package payloads

type ConfirmPaymentRequest struct {
	Cryptocurrency *CryptoCurrency `json:"cryptocurrency,omitempty"`
}

func NewConfirmPaymentRequest(cryptoCurrency *CryptoCurrency) *ConfirmPaymentRequest {
	return &ConfirmPaymentRequest{
		Cryptocurrency: cryptoCurrency,
	}
}
