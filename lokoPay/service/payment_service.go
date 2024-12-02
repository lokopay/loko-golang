package service

import (
	"encoding/json"
	"fmt"
	"github.com/lokopay/loko-golang/lokoPay/client"
	"github.com/lokopay/loko-golang/lokoPay/payloads"
)

type PaymentService struct {
	client *client.Client
}

func NewPaymentService(client *client.Client) *PaymentService {
	return &PaymentService{
		client: client,
	}
}

func (p *PaymentService) DecodePayment(res string) (*payloads.Payment, error) {
	var payment *payloads.Payment
	err := json.Unmarshal([]byte(res), &payment)
	if err != nil {
		return nil, err
	}
	p.client.Decrypt(payment)
	return payment, nil
}

func (p *PaymentService) Create(createPaymentParams *payloads.CreatePaymentRequest) (*payloads.Payment, error) {
	path := "/v1/payments"
	res, err := p.client.Post(path, createPaymentParams)
	if err != nil {
		return nil, err
	}
	return p.DecodePayment(res)
}

func (p *PaymentService) Retrieve(id string) (*payloads.Payment, error) {
	path := fmt.Sprintf("/v1/payments/%s", id)
	res, err := p.client.Get(path, "")
	if err != nil {
		return nil, err
	}
	return p.DecodePayment(res)
}

func (p *PaymentService) Confirm(id string, confirmPaymentParams *payloads.ConfirmPaymentRequest) (*payloads.Payment, error) {
	path := fmt.Sprintf("/v1/payments/%s/confirm", id)
	res, err := p.client.Post(path, confirmPaymentParams)
	if err != nil {
		return nil, err
	}
	return p.DecodePayment(res)
}

func (p *PaymentService) Cancel(id string) (*payloads.Payment, error) {
	path := fmt.Sprintf("/v1/payments/%s/cancel", id)
	res, err := p.client.Post(path, nil)
	if err != nil {
		return nil, err
	}
	return p.DecodePayment(res)
}

func (p *PaymentService) List(queryParams *payloads.QueryParam) ([]*payloads.Payment, error) {
	path := "/v1/payments"
	queryParamsString := queryParams.String()
	if queryParamsString != "" {
		path = fmt.Sprintf("%s?%s", path, queryParamsString)
	}
	res, err := p.client.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var payments struct {
		Payments []*payloads.Payment `json:"data"`
	}
	err = json.Unmarshal([]byte(res), &payments)
	if err != nil {
		return nil, err
	}
	for _, payment := range payments.Payments {
		p.client.Decrypt(payment)
	}
	return payments.Payments, nil
}
