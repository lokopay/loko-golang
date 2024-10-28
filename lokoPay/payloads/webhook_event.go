package payloads

import "encoding/json"

type WebhookEvent struct {
	ID        string          `json:"id"`
	Object    string          `json:"object"`
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
	CreatedAt int64           `json:"created_at"`
}
