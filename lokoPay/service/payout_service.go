package service

import (
	"encoding/json"
	"fmt"
	"github.com/lokopay/loko-golang/lokoPay/client"
	"github.com/lokopay/loko-golang/lokoPay/payloads"
)

type PayoutService struct {
	client *client.Client
}

func NewPayoutService(client *client.Client) *PayoutService {
	return &PayoutService{
		client: client,
	}
}

func (p *PayoutService) DecodePayout(res string) (*payloads.Payout, error) {
	var payout *payloads.Payout
	err := json.Unmarshal([]byte(res), &payout)
	if err != nil {
		return nil, err
	}
	p.client.Decrypt(payout)
	return payout, nil
}

func (p *PayoutService) Create(createPayoutParams *payloads.CreatePayoutRequest) (*payloads.Payout, error) {
	path := "/v1/payouts"
	p.client.Encrypt(createPayoutParams)
	payload, err := json.Marshal(createPayoutParams)
	if err != nil {
		return nil, err
	}
	res, err := p.client.Post(path, string(payload))
	if err != nil {
		return nil, err
	}
	return p.DecodePayout(res)
}

func (p *PayoutService) Retrieve(id string) (*payloads.Payout, error) {
	path := fmt.Sprintf("/v1/payouts/%s", id)
	res, err := p.client.Get(path, nil)
	if err != nil {
		return nil, err
	}
	return p.DecodePayout(res)
}

func (p *PayoutService) Confirm(id string, confirmPayoutParams *payloads.ConfirmPayoutRequest) (*payloads.Payout, error) {
	path := fmt.Sprintf("/v1/payouts/%s/confirm", id)
	payload, err := json.Marshal(confirmPayoutParams)
	if err != nil {
		return nil, err
	}
	res, err := p.client.Post(path, string(payload))
	if err != nil {
		return nil, err
	}
	return p.DecodePayout(res)
}

func (p *PayoutService) Cancel(id string) (*payloads.Payout, error) {
	path := fmt.Sprintf("/v1/payouts/%s/cancel", id)
	res, err := p.client.Post(path, nil)
	if err != nil {
		return nil, err
	}
	return p.DecodePayout(res)
}

func (p *PayoutService) List(queryParams *payloads.QueryParam) ([]*payloads.Payout, error) {
	path := "/v1/payouts"
	queryParamsString := queryParams.String()
	if queryParamsString != "" {
		path = fmt.Sprintf("%s?%s", path, queryParamsString)
	}
	res, err := p.client.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var payouts struct {
		Payouts []*payloads.Payout `json:"data"`
	}
	err = json.Unmarshal([]byte(res), &payouts)
	if err != nil {
		return nil, err
	}
	return payouts.Payouts, nil
}

func (p *PayoutService) Networkfees() ([]*payloads.DestinationNetworkDetail, error) {
	path := "/v1/payouts/networkfees"
	res, err := p.client.Get(path, nil)
	if err != nil {
		return nil, err
	}
	var networkFees struct {
		DestinationNetworkDetails []*payloads.DestinationNetworkDetail `json:"destination_network_details"`
	}
	err = json.Unmarshal([]byte(res), &networkFees)
	if err != nil {
		return nil, err
	}
	return networkFees.DestinationNetworkDetails, nil
}
