package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetActiveDirectoryDomains retrieves a list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomains(parameters connection.APIRequestParameters) ([]ActiveDirectoryDomain, error) {
	return connection.InvokeRequestAll(s.GetActiveDirectoryDomainsPaginated, parameters)
}

// GetActiveDirectoryDomainsPaginated retrieves a paginated list of Active Directory Domains
func (s *Service) GetActiveDirectoryDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ActiveDirectoryDomain], error) {
	body, err := s.getActiveDirectoryDomainsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetActiveDirectoryDomainsPaginated), err
}

func (s *Service) getActiveDirectoryDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ActiveDirectoryDomain], error) {
	body := &connection.APIResponseBodyData[[]ActiveDirectoryDomain]{}

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

func (s *Service) getActiveDirectoryDomainResponseBody(domainID int) (*connection.APIResponseBodyData[ActiveDirectoryDomain], error) {
	body := &connection.APIResponseBodyData[ActiveDirectoryDomain]{}

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
