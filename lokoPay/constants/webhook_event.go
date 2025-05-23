package constants

type WebhookEventType string

func (w WebhookEventType) String() string {
	return string(w)
}

const (
	WebhookEventTypePaymentDeposited WebhookEventType = "payment.deposited"
	WebhookEventTypePaymentFailed    WebhookEventType = "payment.failed"
	WebhookEventTypePaymentExpired   WebhookEventType = "payment.expired"
	WebhookEventTypePaymentSucceeded WebhookEventType = "payment.succeeded"
	WebhookEventTypePayoutPending    WebhookEventType = "payout.pending"
	WebhookEventTypePayoutFailed     WebhookEventType = "payout.failed"
	WebhookEventTypePayoutSucceeded  WebhookEventType = "payout.succeeded"
)
