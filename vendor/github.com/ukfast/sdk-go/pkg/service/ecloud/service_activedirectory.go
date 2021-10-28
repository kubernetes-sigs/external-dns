package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetActiveDirectoryDomains retrieves a list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error) {
	var domains []ActiveDirectoryDomain

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetActiveDirectoryDomainsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, domain := range response.(*PaginatedActiveDirectoryDomain).Items {
			domains = append(domains, domain)
		}
	}

	return domains, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetActiveDirectoryDomainsPaginated retrieves a paginated list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedActiveDirectoryDomain, error) {
	body, err := s.getActiveDirectoryDomainsPaginatedResponseBody(parameters)

	return NewPaginatedActiveDirectoryDomain(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetActiveDirectoryDomainsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getActiveDirectoryDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetActiveDirectoryDomainSliceResponseBody, error) {
	body := &GetActiveDirectoryDomainSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v1/active-directory/domains", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetActiveDirectoryDomain retrieves a single domain by ID
func (s *Service) GetActiveDirectoryDomain(domainID int) (ActiveDirectoryDomain, error) {
	body, err := s.getActiveDirectoryDomainResponseBody(domainID)

	return body.Data, err
}

func (s *Service) getActiveDirectoryDomainResponseBody(domainID int) (*GetActiveDirectoryDomainResponseBody, error) {
	body := &GetActiveDirectoryDomainResponseBody{}

	if domainID < 1 {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/active-directory/domains/%d", domainID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ActiveDirectoryDomainNotFoundError{ID: domainID}
		}

		return nil
	})
}
