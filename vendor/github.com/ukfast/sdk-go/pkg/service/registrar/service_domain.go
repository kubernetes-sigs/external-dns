package registrar

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

func (s *Service) getDomainResponseBody(domainName string) (*GetDomainResponseBody, error) {
	body := &GetDomainResponseBody{}

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

func (s *Service) getDomainNameserversResponseBody(domainName string) (*GetNameserverSliceResponseBody, error) {
	body := &GetNameserverSliceResponseBody{}

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
