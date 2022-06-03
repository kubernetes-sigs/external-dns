package registrar

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

	response, err := s.connection.Get("/registrar/v1/domains", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDomain retrieves a single domain by name
func (s *Service) GetDomain(domainName string) (Domain, error) {
	body, err := s.getDomainResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getDomainResponseBody(domainName string) (*connection.APIResponseBodyData[Domain], error) {
	body := &connection.APIResponseBodyData[Domain]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/registrar/v1/domains/%s", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainNameservers retrieves the nameservers for a domain
func (s *Service) GetDomainNameservers(domainName string) ([]Nameserver, error) {
	body, err := s.getDomainNameserversResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getDomainNameserversResponseBody(domainName string) (*connection.APIResponseBodyData[[]Nameserver], error) {
	body := &connection.APIResponseBodyData[[]Nameserver]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/registrar/v1/domains/%s/nameservers", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}
