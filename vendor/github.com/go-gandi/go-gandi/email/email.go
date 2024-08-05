package email

import (
	"encoding/json"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Email API client
func New(config config.Config) *Email {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
	client.SetEndpoint("email/")
	return &Email{client: *client}
}

// NewFromClient returns an instance of the Email API client
func NewFromClient(g client.Gandi) *Email {
	g.SetEndpoint("email/")
	return &Email{client: g}
}

// ListMailboxes list mailboxes attached to domain
func (e *Email) ListMailboxes(domain string) (mailboxes []ListMailboxResponse, err error) {
	_, elements, err := e.client.GetCollection("/mailboxes/"+domain, nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var mailbox ListMailboxResponse
		err := json.Unmarshal(element, &mailbox)
		if err != nil {
			return nil, err
		}
		mailboxes = append(mailboxes, mailbox)
	}
	return mailboxes, nil
}

// GetMailbox returns all the parameters linked to a specific mailbox
func (e *Email) GetMailbox(domain, mailbox_id string) (mailbox MailboxResponse, err error) {
	_, err = e.client.Get("mailboxes/"+domain+"/"+mailbox_id, nil, &mailbox)
	return
}

// CreateEmail creates a new mailbox for the given domain
func (e *Email) CreateEmail(domain string, req CreateEmailRequest) (err error) {
	_, err = e.client.Post("mailboxes/"+domain, req, nil)
	return
}

// UpdateEmail update mailbox parameters
func (e *Email) UpdateEmail(domain, mailbox_id string, req UpdateEmailRequest) (err error) {
	_, err = e.client.Patch("mailboxes/"+domain+"/"+mailbox_id, req, nil)
	return
}

// DeleteEmail remove mailbox
func (e *Email) DeleteEmail(domain, mailbox_id string) (err error) {
	_, err = e.client.Delete("mailboxes/"+domain+"/"+mailbox_id, nil, nil)
	return
}

// CreateForward creates forwarding
func (e *Email) CreateForward(domain string, req CreateForwardRequest) (err error) {
	_, err = e.client.Post("forwards/"+domain, req, nil)
	return
}

// GetForwards retrieves all forwardings for domain
func (e *Email) GetForwards(domain string) (forwards []GetForwardRequest, err error) {
	_, err = e.client.Get("forwards/"+domain, nil, &forwards)
	return
}

// UpdateForward update forwarding
func (e *Email) UpdateForward(domain, source string, req UpdateForwardRequest) (err error) {
	_, err = e.client.Put("forwards/"+domain+"/"+source, req, nil)
	return
}

// DeleteForward delete forwarding
func (e *Email) DeleteForward(domain, source string) (err error) {
	_, err = e.client.Delete("forwards/"+domain+"/"+source, nil, nil)
	return
}
