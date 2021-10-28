package sharedexchange

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	var domains []Domain

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDomainsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, domain := range response.(*PaginatedDomain).Items {
			domains = append(domains, domain)
		}
	}

	return domains, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error) {
	body, err := s.getDomainsPaginatedResponseBody(parameters)

	return NewPaginatedDomain(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDomainsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetDomainSliceResponseBody, error) {
	body := &GetDomainSliceResponseBody{}

	response, err := s.connection.Get("/shared-exchange/v1/domains", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDomain retrieves a single domain by id
func (s *Service) GetDomain(domainID int) (Domain, error) {
	body, err := s.getDomainResponseBody(domainID)

	return body.Data, err
}

func (s *Service) getDomainResponseBody(domainID int) (*GetDomainResponseBody, error) {
	body := &GetDomainResponseBody{}

	if domainID < 1 {
		return body, fmt.Errorf("invalid domain id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/shared-exchange/v1/domains/%d", domainID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{ID: domainID}
		}

		return nil
	})
}
