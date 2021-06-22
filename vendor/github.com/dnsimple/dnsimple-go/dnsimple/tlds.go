package dnsimple

import (
	"context"
	"fmt"
)

// TldsService handles communication with the Tld related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/tlds/
type TldsService struct {
	client *Client
}

// Tld represents a TLD in DNSimple.
type Tld struct {
	Tld                 string `json:"tld"`
	TldType             int    `json:"tld_type"`
	WhoisPrivacy        bool   `json:"whois_privacy"`
	AutoRenewOnly       bool   `json:"auto_renew_only"`
	MinimumRegistration int    `json:"minimum_registration"`
	RegistrationEnabled bool   `json:"registration_enabled"`
	RenewalEnabled      bool   `json:"renewal_enabled"`
	TransferEnabled     bool   `json:"transfer_enabled"`
}

// TldExtendedAttribute represents an extended attributes supported or required
// by a specific TLD.
//
// See https://developer.dnsimple.com/v2/tlds/
type TldExtendedAttribute struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Required    bool                         `json:"required"`
	Options     []TldExtendedAttributeOption `json:"options"`
}

// TldExtendedAttributeOption represents a single option you can assign to an extended attributes.
//
// See https://developer.dnsimple.com/v2/tlds/
type TldExtendedAttributeOption struct {
	Title       string `json:"title"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

// TldResponse represents a response from an API method that returns a Tld struct.
type TldResponse struct {
	Response
	Data *Tld `json:"data"`
}

// TldsResponse represents a response from an API method that returns a collection of Tld struct.
type TldsResponse struct {
	Response
	Data []Tld `json:"data"`
}

// TldExtendedAttributesResponse represents a response from an API method that returns
// a collection of Tld extended attributes.
type TldExtendedAttributesResponse struct {
	Response
	Data []TldExtendedAttribute `json:"data"`
}

// ListTlds lists the supported TLDs.
//
// See https://developer.dnsimple.com/v2/tlds/#list
func (s *TldsService) ListTlds(ctx context.Context, options *ListOptions) (*TldsResponse, error) {
	path := versioned("/tlds")
	tldsResponse := &TldsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, tldsResponse)
	if err != nil {
		return tldsResponse, err
	}

	tldsResponse.HTTPResponse = resp
	return tldsResponse, nil
}

// GetTld fetches a TLD.
//
// See https://developer.dnsimple.com/v2/tlds/#get
func (s *TldsService) GetTld(ctx context.Context, tld string) (*TldResponse, error) {
	path := versioned(fmt.Sprintf("/tlds/%s", tld))
	tldResponse := &TldResponse{}

	resp, err := s.client.get(ctx, path, tldResponse)
	if err != nil {
		return nil, err
	}

	tldResponse.HTTPResponse = resp
	return tldResponse, nil
}

// GetTldExtendedAttributes fetches the extended attributes of a TLD.
//
// See https://developer.dnsimple.com/v2/tlds/#get
func (s *TldsService) GetTldExtendedAttributes(ctx context.Context, tld string) (*TldExtendedAttributesResponse, error) {
	path := versioned(fmt.Sprintf("/tlds/%s/extended_attributes", tld))
	tldResponse := &TldExtendedAttributesResponse{}

	resp, err := s.client.get(ctx, path, tldResponse)
	if err != nil {
		return nil, err
	}

	tldResponse.HTTPResponse = resp
	return tldResponse, nil
}
