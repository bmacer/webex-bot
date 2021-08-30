package api

import (
	"bytes"
	"fmt"
)

type Webhooks struct {
	Items []Webhook `json:"items"`
}

type Webhook struct {
	Id        string
	Name      string
	TargetUrl string
}

// Get webhooks
func GetWebhooks() Webhooks {
	var wh Webhooks
	rooms_url := "https://webexapis.com/v1/webhooks"
	r1 := sendRequest("GET", rooms_url, nil)
	extract(r1, &wh)
	return wh
}

type WebhookEvent string

const (
	WebhookEventCreated WebhookEvent = "created"
	WebhookEventUpdated WebhookEvent = "updated"
	WebhookEventDeleted WebhookEvent = "deleted"
	WebhookEventStarted WebhookEvent = "started"
	WebhookEventEnded   WebhookEvent = "ended"
)

// Create webhook
func CreateWebhook(name string, targetUrl string, resource string, event WebhookEvent) Webhook {
	var wh Webhook
	url := "https://webexapis.com/v1/webhooks"
	fmt.Println(event)
	whb := []byte(fmt.Sprintf(`{
		"name": "%v",
		"targetUrl": "%v",
		"resource": "%v",
		"event": "%v"
		}`, name, targetUrl, resource, event))
	resp := sendRequest("POST", url, bytes.NewBuffer(whb))
	extract(resp, &wh)
	return wh
}

// Delete webhook
func DeleteWebhook(webhookId string) []byte {
	rooms_url := fmt.Sprintf("https://webexapis.com/v1/webhooks/%v", webhookId)
	req := sendRequest("DELETE", rooms_url, nil)
	return req
}
