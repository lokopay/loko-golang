package payloads

type Order struct {
	Currency string       `json:"currency,omitempty"`
	Discount int64        `json:"discount,omitempty"`
	Items    []*OrderItem `json:"items"`
	Name     string       `json:"name,omitempty"`
	OrderID  string       `json:"order_id,omitempty"`
	SalesTax int64        `json:"sales_tax,omitempty"`
	Shipping int64        `json:"shipping,omitempty"`
	Subtotal int64        `json:"subtotal,omitempty"`
	Total    int64        `json:"total,omitempty"`
}

func NewOrder(id string) *Order {
	return &Order{
		OrderID: id,
	}
}

func (o *Order) SetCurrency(currency string) *Order {
	o.Currency = currency
	return o
}

func (o *Order) SetDiscount(discount int64) *Order {
	o.Discount = discount
	return o
}

func (o *Order) SetItems(items []*OrderItem) *Order {
	o.Items = items
	return o
}

func (o *Order) SetName(name string) *Order {
	o.Name = name
	return o
}

func (o *Order) SetSalesTax(salesTax int64) *Order {
	o.SalesTax = salesTax
	return o
}

func (o *Order) SetShipping(shipping int64) *Order {
	o.Shipping = shipping
	return o
}

func (o *Order) SetSubtotal(subtotal int64) *Order {
	o.Subtotal = subtotal
	return o
}

func (o *Order) SetTotal(total int64) *Order {
	o.Total = total
	return o
}
