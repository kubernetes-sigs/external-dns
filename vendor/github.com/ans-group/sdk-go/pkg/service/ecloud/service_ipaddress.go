package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIPAddresses retrieves a list of ips
func (s *Service) GetIPAddresses(parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return connection.InvokeRequestAll(s.GetIPAddressesPaginated, parameters)
}

// GetIPAddressesPaginated retrieves a paginated list of ips
func (s *Service) GetIPAddressesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	body, err := s.getIPAddressesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIPAddressesPaginated), err
}

func (s *Service) getIPAddressesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]IPAddress], error) {
	body := &connection.APIResponseBodyData[[]IPAddress]{}

	response, err := s.connection.Get("/ecloud/v2/ip-addresses", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetIPAddress retrieves a single ip by id
func (s *Service) GetIPAddress(ipID string) (IPAddress, error) {
	body, err := s.getIPAddressResponseBody(ipID)

	return body.Data, err
}

func (s *Service) getIPAddressResponseBody(ipID string) (*connection.APIResponseBodyData[IPAddress], error) {
	body := &connection.APIResponseBodyData[IPAddress]{}

	if ipID == "" {
		return body, fmt.Errorf("invalid ip address id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &IPAddressNotFoundError{ID: ipID}
		}

		return nil
	})
}

// CreateIPAddress creates a new IPAddress
func (s *Service) CreateIPAddress(req CreateIPAddressRequest) (TaskReference, error) {
	body, err := s.createIPAddressResponseBody(req)

	return body.Data, err
}

func (s *Service) createIPAddressResponseBody(req CreateIPAddressRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/ip-addresses", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchIPAddress patches a IPAddress
func (s *Service) PatchIPAddress(ipID string, req PatchIPAddressRequest) (TaskReference, error) {
	body, err := s.patchIPAddressResponseBody(ipID, req)

	return body.Data, err
}

func (s *Service) patchIPAddressResponseBody(ipID string, req PatchIPAddressRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ipID == "" {
		return body, fmt.Errorf("invalid ip address id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &IPAddressNotFoundError{ID: ipID}
		}

		return nil
	})
}

// DeleteIPAddress deletes a IPAddress
func (s *Service) DeleteIPAddress(ipID string) (string, error) {
	body, err := s.deleteIPAddressResponseBody(ipID)

	return body.Data.TaskID, err
}

func (s *Service) deleteIPAddressResponseBody(ipID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ipID == "" {
		return body, fmt.Errorf("invalid ip address id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/ip-addresses/%s", ipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &IPAddressNotFoundError{ID: ipID}
		}

		return nil
	})
}
