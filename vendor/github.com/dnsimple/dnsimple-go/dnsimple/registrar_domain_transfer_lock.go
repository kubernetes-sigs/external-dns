package dnsimple

import (
	"context"
	"fmt"
)

type DomainTransferLock struct {
	Enabled bool `json:"enabled"`
}

type DomainTransferLockResponse struct {
	Response
	Data *DomainTransferLock `json:"data"`
}

// GetDomainTransferLock gets the domain transfer lock for a domain.
//
// See https://developer.dnsimple.com/v2/registrar/#getDomainTransferLock
func (s *RegistrarService) GetDomainTransferLock(ctx context.Context, accountID string, domainIdentifier string) (*DomainTransferLockResponse, error) {
	path := versioned(fmt.Sprintf("%v/registrar/domains/%v/transfer_lock", accountID, domainIdentifier))
	res := &DomainTransferLockResponse{}

	httpRes, err := s.client.get(ctx, path, res)
	if err != nil {
		return nil, err
	}

	res.HTTPResponse = httpRes
	return res, nil
}

// EnableDomainTransferLock gets the domain transfer lock for a domain.
//
// See https://developer.dnsimple.com/v2/registrar/#enableDomainTransferLock
func (s *RegistrarService) EnableDomainTransferLock(ctx context.Context, accountID string, domainIdentifier string) (*DomainTransferLockResponse, error) {
	path := versioned(fmt.Sprintf("%v/registrar/domains/%v/transfer_lock", accountID, domainIdentifier))
	res := &DomainTransferLockResponse{}

	httpRes, err := s.client.post(ctx, path, nil, res)
	if err != nil {
		return nil, err
	}

	res.HTTPResponse = httpRes
	return res, nil
}

// DisableDomainTransferLock gets the domain transfer lock for a domain.
//
// See https://developer.dnsimple.com/v2/registrar/#disableDomainTransferLock
func (s *RegistrarService) DisableDomainTransferLock(ctx context.Context, accountID string, domainIdentifier string) (*DomainTransferLockResponse, error) {
	path := versioned(fmt.Sprintf("%v/registrar/domains/%v/transfer_lock", accountID, domainIdentifier))
	res := &DomainTransferLockResponse{}

	httpRes, err := s.client.delete(ctx, path, nil, res)
	if err != nil {
		return nil, err
	}

	res.HTTPResponse = httpRes
	return res, nil
}
