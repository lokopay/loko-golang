package service

import (
	"bytes"
	"encoding/json"
	"github.com/lokopay/loko-golang/lokoPay/client"
	"github.com/lokopay/loko-golang/lokoPay/payloads"
	"io/ioutil"
	"net/http"
)

type WebhookEventService struct {
	client *client.Client
}

func NewWebhookEventService(client *client.Client) *WebhookEventService {
	return &WebhookEventService{
		client: client,
	}
}

func (w *WebhookEventService) Retrieve(requestUrl string, r *http.Request) (*payloads.WebhookEvent, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	lokoSignature := r.Header.Get("loko-signature")
	bodyString := string(bodyBytes)
	if !w.client.VerifySignature(requestUrl, bodyString, lokoSignature) {
		return nil, err
	}
	var webhookEvent *payloads.WebhookEvent
	err = json.Unmarshal(bodyBytes, &webhookEvent)
	if err != nil {
		return nil, err
	}
	return webhookEvent, nil
}
