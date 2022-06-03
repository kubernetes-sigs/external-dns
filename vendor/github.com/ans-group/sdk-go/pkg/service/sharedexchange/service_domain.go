package sharedexchange

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return connection.InvokeRequestAll(s.GetDomainsPaginated, parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	body, err := s.getDomainsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDomainsPaginated), err
}

func (s *Service) getDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Domain], error) {
	body := &connection.APIResponseBodyData[[]Domain]{}

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

func (s *Service) getDomainResponseBody(domainID int) (*connection.APIResponseBodyData[Domain], error) {
	body := &connection.APIResponseBodyData[Domain]{}

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
