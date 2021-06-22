package dnsimple

import (
	"context"
	"fmt"
)

// EmailForward represents an email forward in DNSimple.
type EmailForward struct {
	ID        int64  `json:"id,omitempty"`
	DomainID  int64  `json:"domain_id,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func emailForwardPath(accountID string, domainIdentifier string, forwardID int64) (path string) {
	path = fmt.Sprintf("%v/email_forwards", domainPath(accountID, domainIdentifier))
	if forwardID != 0 {
		path += fmt.Sprintf("/%v", forwardID)
	}
	return
}

// EmailForwardResponse represents a response from an API method that returns an EmailForward struct.
type EmailForwardResponse struct {
	Response
	Data *EmailForward `json:"data"`
}

// EmailForwardsResponse represents a response from an API method that returns a collection of EmailForward struct.
type EmailForwardsResponse struct {
	Response
	Data []EmailForward `json:"data"`
}

// ListEmailForwards lists the email forwards for a domain.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#list
func (s *DomainsService) ListEmailForwards(ctx context.Context, accountID string, domainIdentifier string, options *ListOptions) (*EmailForwardsResponse, error) {
	path := versioned(emailForwardPath(accountID, domainIdentifier, 0))
	forwardsResponse := &EmailForwardsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, forwardsResponse)
	if err != nil {
		return nil, err
	}

	forwardsResponse.HTTPResponse = resp
	return forwardsResponse, nil
}

// CreateEmailForward creates a new email forward.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#create
func (s *DomainsService) CreateEmailForward(ctx context.Context, accountID string, domainIdentifier string, forwardAttributes EmailForward) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domainIdentifier, 0))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.post(ctx, path, forwardAttributes, forwardResponse)
	if err != nil {
		return nil, err
	}

	forwardResponse.HTTPResponse = resp
	return forwardResponse, nil
}

// GetEmailForward fetches an email forward.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#get
func (s *DomainsService) GetEmailForward(ctx context.Context, accountID string, domainIdentifier string, forwardID int64) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domainIdentifier, forwardID))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.get(ctx, path, forwardResponse)
	if err != nil {
		return nil, err
	}

	forwardResponse.HTTPResponse = resp
	return forwardResponse, nil
}

// DeleteEmailForward PERMANENTLY deletes an email forward from the domain.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#delete
func (s *DomainsService) DeleteEmailForward(ctx context.Context, accountID string, domainIdentifier string, forwardID int64) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domainIdentifier, forwardID))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	forwardResponse.HTTPResponse = resp
	return forwardResponse, nil
}
