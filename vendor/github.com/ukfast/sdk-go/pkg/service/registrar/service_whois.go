package registrar

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetWhois retrieves WHOIS information for a single domain
func (s *Service) GetWhois(domainName string) (Whois, error) {
	body, err := s.getWhoisResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getWhoisResponseBody(domainName string) (*GetWhoisResponseBody, error) {
	body := &GetWhoisResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/registrar/v1/whois/%s", domainName), connection.APIRequestParameters{})
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

// GetWhoisRaw retrieves raw WHOIS information for a single domain
func (s *Service) GetWhoisRaw(domainName string) (string, error) {
	body, err := s.getWhoisRawResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getWhoisRawResponseBody(domainName string) (*connection.APIResponseBodyStringData, error) {
	body := &connection.APIResponseBodyStringData{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/registrar/v1/whois/%s/raw", domainName), connection.APIRequestParameters{})
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
