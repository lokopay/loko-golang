package lokoPay

import (
	"github.com/lokopay/loko-golang/lokoPay/client"
	"github.com/lokopay/loko-golang/lokoPay/service"
	"time"
)

type LokoPay struct {
	client *client.Client
}

func NewLokoPay(apiPublicKey, apiSecretKey string, liveMode bool) *LokoPay {
	client := client.NewClient(apiPublicKey, apiSecretKey, liveMode)
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
