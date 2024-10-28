package payloads

type OrderItem struct {
	Currency string `json:"currency,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    int64  `json:"price,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

func NewOrderItem(name string, price, quantity int64) *OrderItem {
	return &OrderItem{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func (o *OrderItem) SetCurrency(currency string) *OrderItem {
	o.Currency = currency
	return o
}

func (o *OrderItem) SetMetadata(metadata string) *OrderItem {
	o.Metadata = metadata
	return o
}
