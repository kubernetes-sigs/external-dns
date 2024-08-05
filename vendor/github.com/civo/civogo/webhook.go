package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Webhook is a representation of a saved webhook callback from changes in Civo
type Webhook struct {
	ID                string   `json:"id"`
	Events            []string `json:"events"`
	URL               string   `json:"url"`
	Secret            string   `json:"secret"`
	Disabled          bool     `json:"disabled"`
	Failures          int      `json:"failures"`
	LasrFailureReason string   `json:"last_failure_reason"`
}

// WebhookConfig represents the options required for creating a new webhook
type WebhookConfig struct {
	Events []string `json:"events"`
	URL    string   `json:"url"`
	Secret string   `json:"secret"`
}

// CreateWebhook creates a new webhook
func (c *Client) CreateWebhook(r *WebhookConfig) (*Webhook, error) {
	body, err := c.SendPostRequest("/v2/webhooks", r)
	if err != nil {
		return nil, decodeError(err)
	}

	var n = &Webhook{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(n); err != nil {
		return nil, err
	}

	return n, nil
}

// ListWebhooks returns a list of all webhook within the current account
func (c *Client) ListWebhooks() ([]Webhook, error) {
	resp, err := c.SendGetRequest("/v2/webhooks")
	if err != nil {
		return nil, decodeError(err)
	}

	webhook := make([]Webhook, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&webhook); err != nil {
		return nil, err
	}

	return webhook, nil
}

// FindWebhook finds a webhook by either part of the ID or part of the name
func (c *Client) FindWebhook(search string) (*Webhook, error) {
	webhooks, err := c.ListWebhooks()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Webhook{}

	for _, value := range webhooks {
		if value.URL == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.URL, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// UpdateWebhook updates a webhook
func (c *Client) UpdateWebhook(id string, r *WebhookConfig) (*Webhook, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/webhooks/%s", id), r)
	if err != nil {
		return nil, decodeError(err)
	}

	var n = &Webhook{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(n); err != nil {
		return nil, err
	}

	return n, nil
}

// DeleteWebhook deletes a webhook
func (c *Client) DeleteWebhook(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/webhooks/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
