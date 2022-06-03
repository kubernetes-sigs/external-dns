package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDHCPs retrieves a list of dhcps
func (s *Service) GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error) {
	return connection.InvokeRequestAll(s.GetDHCPsPaginated, parameters)
}

// GetDHCPsPaginated retrieves a paginated list of dhcps
func (s *Service) GetDHCPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DHCP], error) {
	body, err := s.getDHCPsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDHCPsPaginated), err
}

func (s *Service) getDHCPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]DHCP], error) {
	body := &connection.APIResponseBodyData[[]DHCP]{}

	response, err := s.connection.Get("/ecloud/v2/dhcps", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDHCP retrieves a single dhcp by id
func (s *Service) GetDHCP(dhcpID string) (DHCP, error) {
	body, err := s.getDHCPResponseBody(dhcpID)

	return body.Data, err
}

func (s *Service) getDHCPResponseBody(dhcpID string) (*connection.APIResponseBodyData[DHCP], error) {
	body := &connection.APIResponseBodyData[DHCP]{}

	if dhcpID == "" {
		return body, fmt.Errorf("invalid dhcp id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/dhcps/%s", dhcpID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DHCPNotFoundError{ID: dhcpID}
		}

		return nil
	})
}

// GetDHCPTasks retrieves a list of DHCP tasks
func (s *Service) GetDHCPTasks(dhcpID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetDHCPTasksPaginated(dhcpID, p)
	}, parameters)
}

// GetDHCPTasksPaginated retrieves a paginated list of DHCP tasks
func (s *Service) GetDHCPTasksPaginated(dhcpID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getDHCPTasksPaginatedResponseBody(dhcpID, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetDHCPTasksPaginated(dhcpID, p)
	}), err
}

func (s *Service) getDHCPTasksPaginatedResponseBody(dhcpID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if dhcpID == "" {
		return body, fmt.Errorf("invalid dhcp id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/dhcps/%s/tasks", dhcpID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DHCPNotFoundError{ID: dhcpID}
		}

		return nil
	})
}
