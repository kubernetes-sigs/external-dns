package dnsimple

import (
	"context"
	"fmt"
)

// DomainsService handles communication with the domain related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/domains/
type DomainsService struct {
	client *Client
}

// Domain represents a domain in DNSimple.
type Domain struct {
	ID           int64  `json:"id,omitempty"`
	AccountID    int64  `json:"account_id,omitempty"`
	RegistrantID int64  `json:"registrant_id,omitempty"`
	Name         string `json:"name,omitempty"`
	UnicodeName  string `json:"unicode_name,omitempty"`
	Token        string `json:"token,omitempty"`
	State        string `json:"state,omitempty"`
	AutoRenew    bool   `json:"auto_renew,omitempty"`
	PrivateWhois bool   `json:"private_whois,omitempty"`
	ExpiresAt    string `json:"expires_at,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

func domainPath(accountID string, domainIdentifier string) (path string) {
	path = fmt.Sprintf("/%v/domains", accountID)
	if domainIdentifier != "" {
		path += fmt.Sprintf("/%v", domainIdentifier)
	}
	return
}

// DomainResponse represents a response from an API method that returns a Domain struct.
type DomainResponse struct {
	Response
	Data *Domain `json:"data"`
}

// DomainsResponse represents a response from an API method that returns a collection of Domain struct.
type DomainsResponse struct {
	Response
	Data []Domain `json:"data"`
}

// DomainListOptions specifies the optional parameters you can provide
// to customize the DomainsService.ListDomains method.
type DomainListOptions struct {
	// Select domains where the name contains given string.
	NameLike *string `url:"name_like,omitempty"`

	// Select domains where the registrant matches given ID.
	RegistrantID *int `url:"registrant_id,omitempty"`

	ListOptions
}

// ListDomains lists the domains for an account.
//
// See https://developer.dnsimple.com/v2/domains/#list
func (s *DomainsService) ListDomains(ctx context.Context, accountID string, options *DomainListOptions) (*DomainsResponse, error) {
	path := versioned(domainPath(accountID, ""))
	domainsResponse := &DomainsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, domainsResponse)
	if err != nil {
		return nil, err
	}

	domainsResponse.HTTPResponse = resp
	return domainsResponse, nil
}

// CreateDomain creates a new domain in the account.
//
// See https://developer.dnsimple.com/v2/domains/#create
func (s *DomainsService) CreateDomain(ctx context.Context, accountID string, domainAttributes Domain) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, ""))
	domainResponse := &DomainResponse{}

	resp, err := s.client.post(ctx, path, domainAttributes, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}

// GetDomain fetches a domain.
//
// See https://developer.dnsimple.com/v2/domains/#get
func (s *DomainsService) GetDomain(ctx context.Context, accountID string, domainIdentifier string) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, domainIdentifier))
	domainResponse := &DomainResponse{}

	resp, err := s.client.get(ctx, path, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}

// DeleteDomain PERMANENTLY deletes a domain from the account.
//
// See https://developer.dnsimple.com/v2/domains/#delete
func (s *DomainsService) DeleteDomain(ctx context.Context, accountID string, domainIdentifier string) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, domainIdentifier))
	domainResponse := &DomainResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}
