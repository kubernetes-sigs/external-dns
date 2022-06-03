package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	return connection.InvokeRequestAll(s.GetHostsPaginated, parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error) {
	body, err := s.getHostsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetHostsPaginated), err
}

func (s *Service) getHostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Host], error) {
	body := &connection.APIResponseBodyData[[]Host]{}

	response, err := s.connection.Get("/ecloud/v2/hosts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID string) (Host, error) {
	body, err := s.getHostResponseBody(hostID)

	return body.Data, err
}

func (s *Service) getHostResponseBody(hostID string) (*connection.APIResponseBodyData[Host], error) {
	body := &connection.APIResponseBodyData[Host]{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// CreateHost creates a host
func (s *Service) CreateHost(req CreateHostRequest) (TaskReference, error) {
	body, err := s.createHostResponseBody(req)

	return body.Data, err
}

func (s *Service) createHostResponseBody(req CreateHostRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/hosts", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchHost patches a host
func (s *Service) PatchHost(hostID string, req PatchHostRequest) (TaskReference, error) {
	body, err := s.patchHostResponseBody(hostID, req)

	return body.Data, err
}

func (s *Service) patchHostResponseBody(hostID string, req PatchHostRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// DeleteHost deletes a host
func (s *Service) DeleteHost(hostID string) (string, error) {
	body, err := s.deleteHostResponseBody(hostID)

	return body.Data.TaskID, err
}

func (s *Service) deleteHostResponseBody(hostID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/hosts/%s", hostID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}

// GetHostTasks retrieves a list of Host tasks
func (s *Service) GetHostTasks(hostID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostTasksPaginated(hostID, p)
	}, parameters)
}

// GetHostTasksPaginated retrieves a paginated list of Host tasks
func (s *Service) GetHostTasksPaginated(hostID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getHostTasksPaginatedResponseBody(hostID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostTasksPaginated(hostID, p)
	}), err
}

func (s *Service) getHostTasksPaginatedResponseBody(hostID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if hostID == "" {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/hosts/%s/tasks", hostID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}
