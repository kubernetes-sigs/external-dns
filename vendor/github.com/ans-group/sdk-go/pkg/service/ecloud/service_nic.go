package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNICs retrieves a list of nics
func (s *Service) GetNICs(parameters connection.APIRequestParameters) ([]NIC, error) {
	return connection.InvokeRequestAll(s.GetNICsPaginated, parameters)
}

// GetNICsPaginated retrieves a paginated list of nics
func (s *Service) GetNICsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	body, err := s.getNICsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetNICsPaginated), err
}

func (s *Service) getNICsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NIC], error) {
	body := &connection.APIResponseBodyData[[]NIC]{}

	response, err := s.connection.Get("/ecloud/v2/nics", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNIC retrieves a single nic by id
func (s *Service) GetNIC(nicID string) (NIC, error) {
	body, err := s.getNICResponseBody(nicID)

	return body.Data, err
}

func (s *Service) getNICResponseBody(nicID string) (*connection.APIResponseBodyData[NIC], error) {
	body := &connection.APIResponseBodyData[NIC]{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s", nicID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// GetNICTasks retrieves a list of NIC tasks
func (s *Service) GetNICTasks(nicID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNICTasksPaginated(nicID, p)
	}, parameters)
}

// GetNICTasksPaginated retrieves a paginated list of NIC tasks
func (s *Service) GetNICTasksPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getNICTasksPaginatedResponseBody(nicID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNICTasksPaginated(nicID, p)
	}), err
}

func (s *Service) getNICTasksPaginatedResponseBody(nicID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s/tasks", nicID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// GetNICIPAddress retrieves a list of NIC IP addresses
func (s *Service) GetNICIPAddresses(nicID string, parameters connection.APIRequestParameters) ([]IPAddress, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}, parameters)
}

// GetNICIPAddressPaginated retrieves a paginated list of NIC IP addresses
func (s *Service) GetNICIPAddressesPaginated(nicID string, parameters connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
	body, err := s.getNICIPAddressesPaginatedResponseBody(nicID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[IPAddress], error) {
		return s.GetNICIPAddressesPaginated(nicID, p)
	}), err
}

func (s *Service) getNICIPAddressesPaginatedResponseBody(nicID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]IPAddress], error) {
	body := &connection.APIResponseBodyData[[]IPAddress]{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

func (s *Service) AssignNICIPAddress(nicID string, req AssignIPAddressRequest) (string, error) {
	body, err := s.assignNICIPAddressResponseBody(nicID, req)

	return body.Data.TaskID, err
}

func (s *Service) assignNICIPAddressResponseBody(nicID string, req AssignIPAddressRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses", nicID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}

// UnassignNICIPAddress unassigns an IP Address from a resource
func (s *Service) UnassignNICIPAddress(nicID string, ipID string) (string, error) {
	body, err := s.unassignNICIPAddressResponseBody(nicID, ipID)

	return body.Data.TaskID, err
}

func (s *Service) unassignNICIPAddressResponseBody(nicID string, ipID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if nicID == "" {
		return body, fmt.Errorf("invalid nic id")
	}

	if ipID == "" {
		return body, fmt.Errorf("invalid ip address id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/nics/%s/ip-addresses/%s", nicID, ipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NICNotFoundError{ID: nicID}
		}

		return nil
	})
}
