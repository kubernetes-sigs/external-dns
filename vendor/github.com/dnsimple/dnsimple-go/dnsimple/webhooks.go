package dnsimple

import (
	"context"
	"fmt"
)

// WebhooksService handles communication with the webhook related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/webhooks
type WebhooksService struct {
	client *Client
}

// Webhook represents a DNSimple webhook.
type Webhook struct {
	ID  int64  `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

func webhookPath(accountID string, webhookID int64) (path string) {
	path = fmt.Sprintf("/%v/webhooks", accountID)
	if webhookID != 0 {
		path = fmt.Sprintf("%v/%v", path, webhookID)
	}
	return
}

// WebhookResponse represents a response from an API method that returns a Webhook struct.
type WebhookResponse struct {
	Response
	Data *Webhook `json:"data"`
}

// WebhooksResponse represents a response from an API method that returns a collection of Webhook struct.
type WebhooksResponse struct {
	Response
	Data []Webhook `json:"data"`
}

// ListWebhooks lists the webhooks for an account.
//
// See https://developer.dnsimple.com/v2/webhooks/#listWebhooks
func (s *WebhooksService) ListWebhooks(ctx context.Context, accountID string, _ *ListOptions) (*WebhooksResponse, error) {
	path := versioned(webhookPath(accountID, 0))
	webhooksResponse := &WebhooksResponse{}

	resp, err := s.client.get(ctx, path, webhooksResponse)
	if err != nil {
		return webhooksResponse, err
	}

	webhooksResponse.HTTPResponse = resp
	return webhooksResponse, nil
}

// CreateWebhook creates a new webhook.
//
// See https://developer.dnsimple.com/v2/webhooks/#createWebhook
func (s *WebhooksService) CreateWebhook(ctx context.Context, accountID string, webhookAttributes Webhook) (*WebhookResponse, error) {
	path := versioned(webhookPath(accountID, 0))
	webhookResponse := &WebhookResponse{}

	resp, err := s.client.post(ctx, path, webhookAttributes, webhookResponse)
	if err != nil {
		return nil, err
	}

	webhookResponse.HTTPResponse = resp
	return webhookResponse, nil
}

// GetWebhook fetches a webhook.
//
// See https://developer.dnsimple.com/v2/webhooks/#getWebhook
func (s *WebhooksService) GetWebhook(ctx context.Context, accountID string, webhookID int64) (*WebhookResponse, error) {
	path := versioned(webhookPath(accountID, webhookID))
	webhookResponse := &WebhookResponse{}

	resp, err := s.client.get(ctx, path, webhookResponse)
	if err != nil {
		return nil, err
	}

	webhookResponse.HTTPResponse = resp
	return webhookResponse, nil
}

// DeleteWebhook PERMANENTLY deletes the webhook.
//
// See https://developer.dnsimple.com/v2/webhooks/#deleteWebhook
func (s *WebhooksService) DeleteWebhook(ctx context.Context, accountID string, webhookID int64) (*WebhookResponse, error) {
	path := versioned(webhookPath(accountID, webhookID))
	webhookResponse := &WebhookResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	webhookResponse.HTTPResponse = resp
	return webhookResponse, nil
}
