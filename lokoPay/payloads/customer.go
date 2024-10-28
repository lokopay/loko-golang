package payloads

type Customer struct {
	DestinationAddress  string `json:"destination_address,omitempty" encryptAbleField:"true"`
	DestinationCurrency string `json:"destination_currency,omitempty"`
	DestinationNetwork  string `json:"destination_network,omitempty"`
	Email               string `json:"email,omitempty"`
	ID                  string `json:"id,omitempty"`
	IPAddress           string `json:"ip_address,omitempty"`
}

func NewCustomer(id string) *Customer {
	return &Customer{ID: id}
}

func (c *Customer) SetDestinationAddress(destinationAddress string) *Customer {
	c.DestinationAddress = destinationAddress
	return c
}

func (c *Customer) SetDestinationCurrency(destinationCurrency string) *Customer {
	c.DestinationCurrency = destinationCurrency
	return c
}

func (c *Customer) SetDestinationNetwork(destinationNetwork string) *Customer {
	c.DestinationNetwork = destinationNetwork
	return c
}

func (c *Customer) SetEmail(email string) *Customer {
	c.Email = email
	return c
}

func (c *Customer) SetIPAddress(ipAddress string) *Customer {
	c.IPAddress = ipAddress
	return c
}
