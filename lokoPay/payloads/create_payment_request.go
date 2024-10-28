package payloads

type CreatePaymentRequest struct {
	Amount      string    `json:"amount,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	Customer    *Customer `json:"customer,omitempty"`
	Description string    `json:"description,omitempty"`
	Order       *Order    `json:"order,omitempty"`
}

func NewCreatePaymentRequest(amount, currency string) *CreatePaymentRequest {
	return &CreatePaymentRequest{
		Amount:   amount,
		Currency: currency,
	}
}

func (c *CreatePaymentRequest) SetCustomer(customer *Customer) *CreatePaymentRequest {
	c.Customer = customer
	return c
}

func (c *CreatePaymentRequest) SetDescription(description string) *CreatePaymentRequest {
	c.Description = description
	return c
}

func (c *CreatePaymentRequest) SetOrder(order *Order) *CreatePaymentRequest {
	c.Order = order
	return c
}
