package dnsimple

import (
	"context"
	"fmt"
)

// DomainPush represents a domain push in DNSimple.
type DomainPush struct {
	ID         int64  `json:"id,omitempty"`
	DomainID   int64  `json:"domain_id,omitempty"`
	ContactID  int64  `json:"contact_id,omitempty"`
	AccountID  int64  `json:"account_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	AcceptedAt string `json:"accepted_at,omitempty"`
}

func domainPushPath(accountID string, pushID int64) (path string) {
	path = fmt.Sprintf("/%v/pushes", accountID)
	if pushID != 0 {
		path += fmt.Sprintf("/%v", pushID)
	}
	return
}

// DomainPushResponse represents a response from an API method that returns a DomainPush struct.
type DomainPushResponse struct {
	Response
	Data *DomainPush `json:"data"`
}

// DomainPushesResponse represents a response from an API method that returns a collection of DomainPush struct.
type DomainPushesResponse struct {
	Response
	Data []DomainPush `json:"data"`
}

// DomainPushAttributes represent a domain push payload (see initiate).
type DomainPushAttributes struct {
	NewAccountEmail string `json:"new_account_email,omitempty"`
	ContactID       int64  `json:"contact_id,omitempty"`
}

// InitiatePush initiate a new domain push.
//
// See https://developer.dnsimple.com/v2/domains/pushes/#initiateDomainPush
func (s *DomainsService) InitiatePush(ctx context.Context, accountID, domainID string, pushAttributes DomainPushAttributes) (*DomainPushResponse, error) {
	path := versioned(fmt.Sprintf("/%v/pushes", domainPath(accountID, domainID)))
	pushResponse := &DomainPushResponse{}

	resp, err := s.client.post(ctx, path, pushAttributes, pushResponse)
	if err != nil {
		return nil, err
	}

	pushResponse.HTTPResponse = resp
	return pushResponse, nil
}

// ListPushes lists the pushes for an account.
//
// See https://developer.dnsimple.com/v2/domains/pushes/#listPushes
func (s *DomainsService) ListPushes(ctx context.Context, accountID string, options *ListOptions) (*DomainPushesResponse, error) {
	path := versioned(domainPushPath(accountID, 0))
	pushesResponse := &DomainPushesResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, pushesResponse)
	if err != nil {
		return nil, err
	}

	pushesResponse.HTTPResponse = resp
	return pushesResponse, nil
}

// AcceptPush accept a push for a domain.
//
// See https://developer.dnsimple.com/v2/domains/pushes/#acceptPush
func (s *DomainsService) AcceptPush(ctx context.Context, accountID string, pushID int64, pushAttributes DomainPushAttributes) (*DomainPushResponse, error) {
	path := versioned(domainPushPath(accountID, pushID))
	pushResponse := &DomainPushResponse{}

	resp, err := s.client.post(ctx, path, pushAttributes, nil)
	if err != nil {
		return nil, err
	}

	pushResponse.HTTPResponse = resp
	return pushResponse, nil
}

// RejectPush reject a push for a domain.
//
// See https://developer.dnsimple.com/v2/domains/pushes/#rejectPush
func (s *DomainsService) RejectPush(ctx context.Context, accountID string, pushID int64) (*DomainPushResponse, error) {
	path := versioned(domainPushPath(accountID, pushID))
	pushResponse := &DomainPushResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	pushResponse.HTTPResponse = resp
	return pushResponse, nil
}
