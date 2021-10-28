package dnsimple

import (
	"context"
	"fmt"
)

// RegistrarService handles communication with the registrar related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/registrar/
type RegistrarService struct {
	client *Client
}

// DomainCheck represents the result of a domain check.
type DomainCheck struct {
	Domain    string `json:"domain"`
	Available bool   `json:"available"`
	Premium   bool   `json:"premium"`
}

// DomainCheckResponse represents a response from a domain check request.
type DomainCheckResponse struct {
	Response
	Data *DomainCheck `json:"data"`
}

// CheckDomain checks a domain name.
//
// See https://developer.dnsimple.com/v2/registrar/#check
func (s *RegistrarService) CheckDomain(ctx context.Context, accountID string, domainName string) (*DomainCheckResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/check", accountID, domainName))
	checkResponse := &DomainCheckResponse{}

	resp, err := s.client.get(ctx, path, checkResponse)
	if err != nil {
		return nil, err
	}

	checkResponse.HTTPResponse = resp
	return checkResponse, nil
}

// DomainPremiumPrice represents the premium price for a premium domain.
type DomainPremiumPrice struct {
	// The domain premium price
	PremiumPrice string `json:"premium_price"`
	// The registrar action.
	// Possible values are registration|transfer|renewal
	Action string `json:"action"`
}

// DomainPremiumPriceResponse represents a response from a domain premium price request.
type DomainPremiumPriceResponse struct {
	Response
	Data *DomainPremiumPrice `json:"data"`
}

// DomainPremiumPriceOptions specifies the optional parameters you can provide
// to customize the RegistrarService.GetDomainPremiumPrice method.
type DomainPremiumPriceOptions struct {
	Action string `url:"action,omitempty"`
}

// GetDomainPremiumPrice gets the premium price for a domain.
//
// Deprecated: GetDomainPremiumPrice has been deprecated, use GetDomainPrices instead.
//
// You must specify an action to get the price for. Valid actions are:
// - registration
// - transfer
// - renewal
//
// See https://developer.dnsimple.com/v2/registrar/#premium-price
func (s *RegistrarService) GetDomainPremiumPrice(ctx context.Context, accountID string, domainName string, options *DomainPremiumPriceOptions) (*DomainPremiumPriceResponse, error) {
	var err error
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/premium_price", accountID, domainName))
	priceResponse := &DomainPremiumPriceResponse{}

	if options != nil {
		path, err = addURLQueryOptions(path, options)
		if err != nil {
			return nil, err
		}
	}

	resp, err := s.client.get(ctx, path, priceResponse)
	if err != nil {
		return nil, err
	}

	priceResponse.HTTPResponse = resp
	return priceResponse, nil
}

// DomainPrice represents the result of a domain prices call.
type DomainPrice struct {
	Domain            string  `json:"domain"`
	Premium           bool    `json:"premium"`
	RegistrationPrice float64 `json:"registration_price"`
	RenewalPrice      float64 `json:"renewal_price"`
	TransferPrice     float64 `json:"transfer_price"`
}

// DomainPriceResponse represents a response from an API method that returns a DomainPrice struct.
type DomainPriceResponse struct {
	Response
	Data *DomainPrice `json:"data"`
}

// GetDomainPrices get prices for a domain.
//
// See https://developer.dnsimple.com/v2/registrar/#getDomainPrices
func (s *RegistrarService) GetDomainPrices(ctx context.Context, accountID string, domainName string) (*DomainPriceResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/prices", accountID, domainName))
	pricesResponse := &DomainPriceResponse{}

	resp, err := s.client.get(ctx, path, pricesResponse)
	if err != nil {
		return nil, err
	}

	pricesResponse.HTTPResponse = resp
	return pricesResponse, nil
}

// DomainRegistration represents the result of a domain registration call.
type DomainRegistration struct {
	ID           int64  `json:"id"`
	DomainID     int64  `json:"domain_id"`
	RegistrantID int64  `json:"registrant_id"`
	Period       int    `json:"period"`
	State        string `json:"state"`
	AutoRenew    bool   `json:"auto_renew"`
	WhoisPrivacy bool   `json:"whois_privacy"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// DomainRegistrationResponse represents a response from an API method that results in a domain registration.
type DomainRegistrationResponse struct {
	Response
	Data *DomainRegistration `json:"data"`
}

// RegisterDomainInput represents the attributes you can pass to a register API request.
// Some attributes are mandatory.
type RegisterDomainInput struct {
	// The ID of the Contact to use as registrant for the domain
	RegistrantID int `json:"registrant_id"`
	// Set to true to enable the whois privacy service. An extra cost may apply.
	// Default to false.
	EnableWhoisPrivacy bool `json:"whois_privacy,omitempty"`
	// Set to true to enable the auto-renewal of the domain.
	// Default to true.
	EnableAutoRenewal bool `json:"auto_renew,omitempty"`
	// Required by some TLDs. Use Tlds.GetTldExtendedAttributes() to get the required entries.
	ExtendedAttributes map[string]string `json:"extended_attributes,omitempty"`
	// Required as confirmation of the price, only if the domain is premium.
	PremiumPrice string `json:"premium_price,omitempty"`
}

// RegisterDomain registers a domain name.
//
// See https://developer.dnsimple.com/v2/registrar/#registerDomain
func (s *RegistrarService) RegisterDomain(ctx context.Context, accountID string, domainName string, input *RegisterDomainInput) (*DomainRegistrationResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/registrations", accountID, domainName))
	registrationResponse := &DomainRegistrationResponse{}

	// TODO: validate mandatory attributes RegistrantID

	resp, err := s.client.post(ctx, path, input, registrationResponse)
	if err != nil {
		return nil, err
	}

	registrationResponse.HTTPResponse = resp
	return registrationResponse, nil
}

// DomainTransfer represents the result of a domain transfer call.
type DomainTransfer struct {
	ID                int64  `json:"id"`
	DomainID          int64  `json:"domain_id"`
	RegistrantID      int64  `json:"registrant_id"`
	State             string `json:"state"`
	AutoRenew         bool   `json:"auto_renew"`
	WhoisPrivacy      bool   `json:"whois_privacy"`
	StatusDescription string `json:"status_description"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
}

// DomainTransferResponse represents a response from an API method that results in a domain transfer.
type DomainTransferResponse struct {
	Response
	Data *DomainTransfer `json:"data"`
}

// TransferDomainInput represents the attributes you can pass to a transfer API request.
// Some attributes are mandatory.
type TransferDomainInput struct {
	// The ID of the Contact to use as registrant for the domain
	RegistrantID int `json:"registrant_id"`
	// The Auth-Code required to transfer the domain.
	// This is provided by the current registrar of the domain.
	AuthCode string `json:"auth_code,omitempty"`
	// Set to true to enable the whois privacy service. An extra cost may apply.
	// Default to false.
	EnableWhoisPrivacy bool `json:"whois_privacy,omitempty"`
	// Set to true to enable the auto-renewal of the domain.
	// Default to true.
	EnableAutoRenewal bool `json:"auto_renew,omitempty"`
	// Required by some TLDs. Use Tlds.GetTldExtendedAttributes() to get the required entries.
	ExtendedAttributes map[string]string `json:"extended_attributes,omitempty"`
	// Required as confirmation of the price, only if the domain is premium.
	PremiumPrice string `json:"premium_price,omitempty"`
}

// TransferDomain transfers a domain name.
//
// See https://developer.dnsimple.com/v2/registrar/#transferDomain
func (s *RegistrarService) TransferDomain(ctx context.Context, accountID string, domainName string, input *TransferDomainInput) (*DomainTransferResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/transfers", accountID, domainName))
	transferResponse := &DomainTransferResponse{}

	// TODO: validate mandatory attributes RegistrantID

	resp, err := s.client.post(ctx, path, input, transferResponse)
	if err != nil {
		return nil, err
	}

	transferResponse.HTTPResponse = resp
	return transferResponse, nil
}

// GetDomainTransfer fetches a domain transfer.
//
// See https://developer.dnsimple.com/v2/registrar/#getDomainTransfer
func (s *RegistrarService) GetDomainTransfer(ctx context.Context, accountID string, domainName string, domainTransferID int64) (*DomainTransferResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/transfers/%v", accountID, domainName, domainTransferID))
	transferResponse := &DomainTransferResponse{}

	resp, err := s.client.get(ctx, path, transferResponse)
	if err != nil {
		return nil, err
	}

	transferResponse.HTTPResponse = resp
	return transferResponse, nil
}

// CancelDomainTransfer cancels an in progress domain transfer.
//
// See https://developer.dnsimple.com/v2/registrar/#cancelDomainTransfer
func (s *RegistrarService) CancelDomainTransfer(ctx context.Context, accountID string, domainName string, domainTransferID int64) (*DomainTransferResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/transfers/%v", accountID, domainName, domainTransferID))
	transferResponse := &DomainTransferResponse{}

	resp, err := s.client.delete(ctx, path, nil, transferResponse)
	if err != nil {
		return nil, err
	}

	transferResponse.HTTPResponse = resp
	return transferResponse, nil
}

// DomainTransferOutResponse represents a response from an API method that results in a domain transfer out.
type DomainTransferOutResponse struct {
	Response
	Data *Domain `json:"data"`
}

// TransferDomainOut prepares a domain for outbound transfer.
//
// See https://developer.dnsimple.com/v2/registrar/#authorizeDomainTransferOut
func (s *RegistrarService) TransferDomainOut(ctx context.Context, accountID string, domainName string) (*DomainTransferOutResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/authorize_transfer_out", accountID, domainName))
	transferResponse := &DomainTransferOutResponse{}

	resp, err := s.client.post(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	transferResponse.HTTPResponse = resp
	return transferResponse, nil
}

// DomainRenewal represents the result of a domain renewal call.
type DomainRenewal struct {
	ID        int64  `json:"id"`
	DomainID  int64  `json:"domain_id"`
	Period    int    `json:"period"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// DomainRenewalResponse represents a response from an API method that returns a domain renewal.
type DomainRenewalResponse struct {
	Response
	Data *DomainRenewal `json:"data"`
}

// RenewDomainInput represents the attributes you can pass to a renew API request.
// Some attributes are mandatory.
type RenewDomainInput struct {
	// The number of years
	Period int `json:"period"`
	// Required as confirmation of the price, only if the domain is premium.
	PremiumPrice string `json:"premium_price,omitempty"`
}

// RenewDomain renews a domain name.
//
// See https://developer.dnsimple.com/v2/registrar/#renewDomain
func (s *RegistrarService) RenewDomain(ctx context.Context, accountID string, domainName string, input *RenewDomainInput) (*DomainRenewalResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/renewals", accountID, domainName))
	renewalResponse := &DomainRenewalResponse{}

	resp, err := s.client.post(ctx, path, input, renewalResponse)
	if err != nil {
		return nil, err
	}

	renewalResponse.HTTPResponse = resp
	return renewalResponse, nil
}
