package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewalls retrieves a list of firewalls
func (s *Service) GetFirewalls(parameters connection.APIRequestParameters) ([]Firewall, error) {
	return connection.InvokeRequestAll(s.GetFirewallsPaginated, parameters)
}

// GetFirewallsPaginated retrieves a paginated list of firewalls
func (s *Service) GetFirewallsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
	body, err := s.getFirewallsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallsPaginated), err
}

func (s *Service) getFirewallsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Firewall], error) {
	body := &connection.APIResponseBodyData[[]Firewall]{}

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

func (s *Service) getFirewallResponseBody(firewallID int) (*connection.APIResponseBodyData[Firewall], error) {
	body := &connection.APIResponseBodyData[Firewall]{}

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

func (s *Service) getFirewallConfigResponseBody(firewallID int) (*connection.APIResponseBodyData[FirewallConfig], error) {
	body := &connection.APIResponseBodyData[FirewallConfig]{}

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
