package service

import (
	"encoding/json"
	"github.com/lokopay/loko-golang/lokoPay/client"
	"github.com/lokopay/loko-golang/lokoPay/payloads"
)

type CustomerWalletService struct {
	client *client.Client
}

func NewCustomerWalletService(client *client.Client) *CustomerWalletService {
	return &CustomerWalletService{
		client: client,
	}
}

func (c *CustomerWalletService) Create(createCustomerWalletParams *payloads.CreateCustomerWalletRequest) (*payloads.CustomerWallet, error) {
	path := "/v1/customer_wallets"
	res, err := c.client.Post(path, createCustomerWalletParams)
	if err != nil {
		return nil, err
	}
	var customerWallet *payloads.CustomerWallet
	err = json.Unmarshal([]byte(res), &customerWallet)
	if err != nil {
		return nil, err
	}
	c.client.Decrypt(customerWallet)
	return customerWallet, nil
}
