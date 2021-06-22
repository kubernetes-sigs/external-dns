package dnsimple

import (
	"context"
	"fmt"
)

// WhoisPrivacy represents a whois privacy in DNSimple.
type WhoisPrivacy struct {
	ID        int64  `json:"id,omitempty"`
	DomainID  int64  `json:"domain_id,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	ExpiresOn string `json:"expires_on,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// WhoisPrivacyRenewal represents a whois privacy renewal in DNSimple.
type WhoisPrivacyRenewal struct {
	ID             int64  `json:"id,omitempty"`
	DomainID       int64  `json:"domain_id,omitempty"`
	WhoisPrivacyID int64  `json:"whois_privacy_id,omitempty"`
	State          string `json:"string,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	ExpiresOn      string `json:"expires_on,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

// WhoisPrivacyResponse represents a response from an API method that returns a WhoisPrivacy struct.
type WhoisPrivacyResponse struct {
	Response
	Data *WhoisPrivacy `json:"data"`
}

// WhoisPrivacyRenewalResponse represents a response from an API method that returns a WhoisPrivacyRenewal struct.
type WhoisPrivacyRenewalResponse struct {
	Response
	Data *WhoisPrivacyRenewal `json:"data"`
}

// GetWhoisPrivacy gets the whois privacy for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/whois-privacy/#get
func (s *RegistrarService) GetWhoisPrivacy(ctx context.Context, accountID string, domainName string) (*WhoisPrivacyResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/whois_privacy", accountID, domainName))
	privacyResponse := &WhoisPrivacyResponse{}

	resp, err := s.client.get(ctx, path, privacyResponse)
	if err != nil {
		return nil, err
	}

	privacyResponse.HTTPResponse = resp
	return privacyResponse, nil
}

// EnableWhoisPrivacy enables the whois privacy for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/whois-privacy/#enable
func (s *RegistrarService) EnableWhoisPrivacy(ctx context.Context, accountID string, domainName string) (*WhoisPrivacyResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/whois_privacy", accountID, domainName))
	privacyResponse := &WhoisPrivacyResponse{}

	resp, err := s.client.put(ctx, path, nil, privacyResponse)
	if err != nil {
		return nil, err
	}

	privacyResponse.HTTPResponse = resp
	return privacyResponse, nil
}

// DisableWhoisPrivacy disables the whois privacy for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/whois-privacy/#enable
func (s *RegistrarService) DisableWhoisPrivacy(ctx context.Context, accountID string, domainName string) (*WhoisPrivacyResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/whois_privacy", accountID, domainName))
	privacyResponse := &WhoisPrivacyResponse{}

	resp, err := s.client.delete(ctx, path, nil, privacyResponse)
	if err != nil {
		return nil, err
	}

	privacyResponse.HTTPResponse = resp
	return privacyResponse, nil
}

// RenewWhoisPrivacy renews the whois privacy for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/whois-privacy/#renew
func (s *RegistrarService) RenewWhoisPrivacy(ctx context.Context, accountID string, domainName string) (*WhoisPrivacyRenewalResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/whois_privacy/renewals", accountID, domainName))
	privacyRenewalResponse := &WhoisPrivacyRenewalResponse{}

	resp, err := s.client.post(ctx, path, nil, privacyRenewalResponse)
	if err != nil {
		return nil, err
	}

	privacyRenewalResponse.HTTPResponse = resp
	return privacyRenewalResponse, nil
}
