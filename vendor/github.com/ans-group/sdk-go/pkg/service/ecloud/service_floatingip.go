package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFloatingIPs retrieves a list of floating ips
func (s *Service) GetFloatingIPs(parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	return connection.InvokeRequestAll(s.GetFloatingIPsPaginated, parameters)
}

// GetFloatingIPsPaginated retrieves a paginated list of floating ips
func (s *Service) GetFloatingIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
	body, err := s.getFloatingIPsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetFloatingIPsPaginated), err
}

func (s *Service) getFloatingIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FloatingIP], error) {
	body := &connection.APIResponseBodyData[[]FloatingIP]{}

	response, err := s.connection.Get("/ecloud/v2/floating-ips", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFloatingIP retrieves a single floating ip by id
func (s *Service) GetFloatingIP(fipID string) (FloatingIP, error) {
	body, err := s.getFloatingIPResponseBody(fipID)

	return body.Data, err
}

func (s *Service) getFloatingIPResponseBody(fipID string) (*connection.APIResponseBodyData[FloatingIP], error) {
	body := &connection.APIResponseBodyData[FloatingIP]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating ip id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}

// CreateFloatingIP creates a new FloatingIP
func (s *Service) CreateFloatingIP(req CreateFloatingIPRequest) (TaskReference, error) {
	body, err := s.createFloatingIPResponseBody(req)

	return body.Data, err
}

func (s *Service) createFloatingIPResponseBody(req CreateFloatingIPRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/floating-ips", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchFloatingIP patches a floating IP
func (s *Service) PatchFloatingIP(fipID string, req PatchFloatingIPRequest) (TaskReference, error) {
	body, err := s.patchFloatingIPResponseBody(fipID, req)

	return body.Data, err
}

func (s *Service) patchFloatingIPResponseBody(fipID string, req PatchFloatingIPRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}

// DeleteFloatingIP deletes a floating IP
func (s *Service) DeleteFloatingIP(fipID string) (string, error) {
	body, err := s.deleteFloatingIPResponseBody(fipID)

	return body.Data.TaskID, err
}

func (s *Service) deleteFloatingIPResponseBody(fipID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/floating-ips/%s", fipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}

// AssignFloatingIP assigns a floating IP to a resource
func (s *Service) AssignFloatingIP(fipID string, req AssignFloatingIPRequest) (string, error) {
	body, err := s.assignFloatingIPResponseBody(fipID, req)

	return body.Data.TaskID, err
}

func (s *Service) assignFloatingIPResponseBody(fipID string, req AssignFloatingIPRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/floating-ips/%s/assign", fipID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}

// UnassignFloatingIP unassigns a floating IP from a resource
func (s *Service) UnassignFloatingIP(fipID string) (string, error) {
	body, err := s.unassignFloatingIPResponseBody(fipID)

	return body.Data.TaskID, err
}

func (s *Service) unassignFloatingIPResponseBody(fipID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/floating-ips/%s/unassign", fipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}

// GetFloatingIPTasks retrieves a list of FloatingIP tasks
func (s *Service) GetFloatingIPTasks(fipID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFloatingIPTasksPaginated(fipID, p)
	}, parameters)
}

// GetFloatingIPTasksPaginated retrieves a paginated list of FloatingIP tasks
func (s *Service) GetFloatingIPTasksPaginated(fipID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getFloatingIPTasksPaginatedResponseBody(fipID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFloatingIPTasksPaginated(fipID, p)
	}), err
}

func (s *Service) getFloatingIPTasksPaginatedResponseBody(fipID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if fipID == "" {
		return body, fmt.Errorf("invalid floating ip id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/floating-ips/%s/tasks", fipID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FloatingIPNotFoundError{ID: fipID}
		}

		return nil
	})
}
