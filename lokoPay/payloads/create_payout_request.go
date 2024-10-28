package payloads

type CreatePayoutRequest struct {
	Amount                  string                   `json:"amount,omitempty"`
	Currency                string                   `json:"currency,omitempty"`
	Customer                *Customer                `json:"customer,omitempty"`
	Description             string                   `json:"description,omitempty"`
	TransferWithNativeToken *TransferWithNativeToken `json:"transfer_with_native_token,omitempty"`
}

func NewCreatePayoutRequest(amount, currency string) *CreatePayoutRequest {
	return &CreatePayoutRequest{
		Amount:   amount,
		Currency: currency,
	}
}

func (c *CreatePayoutRequest) SetCustomer(customer *Customer) *CreatePayoutRequest {
	c.Customer = customer
	return c
}

func (c *CreatePayoutRequest) SetDescription(description string) *CreatePayoutRequest {
	c.Description = description
	return c
}

func (c *CreatePayoutRequest) SetTransferWithNativeToken(enabled bool) *CreatePayoutRequest {
	c.TransferWithNativeToken = &TransferWithNativeToken{
		Enabled: enabled,
	}
	return c
}
