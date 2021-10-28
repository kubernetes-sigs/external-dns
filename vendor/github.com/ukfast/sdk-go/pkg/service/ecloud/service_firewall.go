package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewalls retrieves a list of firewalls
func (s *Service) GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error) {
	var firewalls []Firewall

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, firewall := range response.(*PaginatedFirewall).Items {
			firewalls = append(firewalls, firewall)
		}
	}

	return firewalls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallsPaginated retrieves a paginated list of firewalls
func (s *Service) GetFirewallsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewall, error) {
	body, err := s.getFirewallsPaginatedResponseBody(parameters)

	return NewPaginatedFirewall(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallSliceResponseBody, error) {
	body := &GetFirewallSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v1/firewalls", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewall retrieves a single firewall by ID
func (s *Service) GetFirewall(firewallID int) (Firewall, error) {
	body, err := s.getFirewallResponseBody(firewallID)

	return body.Data, err
}

func (s *Service) getFirewallResponseBody(firewallID int) (*GetFirewallResponseBody, error) {
	body := &GetFirewallResponseBody{}

	if firewallID < 1 {
		return body, fmt.Errorf("invalid firewall id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/firewalls/%d", firewallID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallNotFoundError{ID: firewallID}
		}

		return nil
	})
}

// GetFirewallConfig retrieves a single firewall config by ID
func (s *Service) GetFirewallConfig(firewallID int) (FirewallConfig, error) {
	body, err := s.getFirewallConfigResponseBody(firewallID)

	return body.Data, err
}

func (s *Service) getFirewallConfigResponseBody(firewallID int) (*GetFirewallConfigResponseBody, error) {
	body := &GetFirewallConfigResponseBody{}

	if firewallID < 1 {
		return body, fmt.Errorf("invalid firewall id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/firewalls/%d/config", firewallID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallNotFoundError{ID: firewallID}
		}

		return nil
	})
}
