package lokoPay

import (
	"loko-golang/lokoPay/client"
	"loko-golang/lokoPay/service"
	"time"
)

const (
	LiveApiBaseUrl = "https://api.lokopay.io"
	TestApiBaseUrl = "https://api-test.lokopay.io"
)

type LokoPay struct {
	client *client.Client
}

func NewLokoPay(apiPublicKey, apiSecretKey string, liveMode bool) *LokoPay {
	var baseUrl string
	if liveMode {
		baseUrl = LiveApiBaseUrl
	} else {
		baseUrl = TestApiBaseUrl
	}
	client := client.NewClient(apiPublicKey, apiSecretKey, baseUrl)
	return &LokoPay{
		client: client,
	}
}

func (l *LokoPay) SetBaseUrl(baseUrl string) {
	l.client.SetBaseUrl(baseUrl)
}

func (l *LokoPay) SetTimeout(timeout time.Duration) {
	l.client.SetTimeout(timeout)
}

func (l *LokoPay) Payment() *service.PaymentService {
	return service.NewPaymentService(l.client)
}

func (l *LokoPay) Payout() *service.PayoutService {
	return service.NewPayoutService(l.client)
}

func (l *LokoPay) CustomerWallet() *service.CustomerWalletService {
	return service.NewCustomerWalletService(l.client)
}

func (l *LokoPay) WebhookEvent() *service.WebhookEventService {
	return service.NewWebhookEventService(l.client)
}
