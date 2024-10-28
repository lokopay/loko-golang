package client_test

import (
	"loko-golang/lokoPay/client"
	"testing"
)

var clientServ *client.Client

func init() {
	clientServ = client.NewClient(
		"xxx",
		"xxx",
		"https://api-test.lokopay.io")
}

func TestClient_Post(t *testing.T) {
	t.Log("TestClient_Post")
	payload := `{
		"amount": "10000",
		"currency": "USDC",
		"description": "order #123",
		"customer_ip_address": "null",
		"customer_email": "ryo@lokofi.io",
		"customer":{
			"id":"test-xx-2"
		}
	}`
	res, err := clientServ.Post("/v1/payments", payload)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
	t.Log(res)
}

func TestClient_Get(t *testing.T) {
	t.Log("TestClient_Get")
	res, err := clientServ.Get("/v1/networkfees", "")
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
	t.Log(res)
}
