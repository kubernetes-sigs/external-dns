package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHostGroups retrieves a list of host groups
func (s *Service) GetHostGroups(parameters connection.APIRequestParameters) ([]HostGroup, error) {
	return connection.InvokeRequestAll(s.GetHostGroupsPaginated, parameters)
}

// GetHostGroupsPaginated retrieves a paginated list of host groups
func (s *Service) GetHostGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostGroup], error) {
	body, err := s.getHostGroupsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetHostGroupsPaginated), err
}

func (s *Service) getHostGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]HostGroup], error) {
	body := &connection.APIResponseBodyData[[]HostGroup]{}

	response, err := s.connection.Get("/ecloud/v2/host-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHostGroup retrieves a single host group by id
func (s *Service) GetHostGroup(hostGroupID string) (HostGroup, error) {
	body, err := s.getHostGroupResponseBody(hostGroupID)

	return body.Data, err
}

func (s *Service) getHostGroupResponseBody(hostGroupID string) (*connection.APIResponseBodyData[HostGroup], error) {
	body := &connection.APIResponseBodyData[HostGroup]{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}

// CreateHostGroup creates a host group
func (s *Service) CreateHostGroup(req CreateHostGroupRequest) (TaskReference, error) {
	body, err := s.createHostGroupResponseBody(req)

	return body.Data, err
}

func (s *Service) createHostGroupResponseBody(req CreateHostGroupRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/host-groups", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchHostGroup patches a host group
func (s *Service) PatchHostGroup(hostGroupID string, req PatchHostGroupRequest) (TaskReference, error) {
	body, err := s.patchHostGroupResponseBody(hostGroupID, req)

	return body.Data, err
}

func (s *Service) patchHostGroupResponseBody(hostGroupID string, req PatchHostGroupRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}

// DeleteHostGroup deletes a host group
func (s *Service) DeleteHostGroup(hostGroupID string) (string, error) {
	body, err := s.deleteHostGroupResponseBody(hostGroupID)

	return body.Data.TaskID, err
}

func (s *Service) deleteHostGroupResponseBody(hostGroupID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/host-groups/%s", hostGroupID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}

// GetHostGroupTasks retrieves a list of HostGroup tasks
func (s *Service) GetHostGroupTasks(hostGroupID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostGroupTasksPaginated(hostGroupID, p)
	}, parameters)
}

// GetHostGroupTasksPaginated retrieves a paginated list of HostGroup tasks
func (s *Service) GetHostGroupTasksPaginated(hostGroupID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getHostGroupTasksPaginatedResponseBody(hostGroupID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetHostGroupTasksPaginated(hostGroupID, p)
	}), err
}

func (s *Service) getHostGroupTasksPaginatedResponseBody(hostGroupID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if hostGroupID == "" {
		return body, fmt.Errorf("invalid host group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/host-groups/%s/tasks", hostGroupID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostGroupNotFoundError{ID: hostGroupID}
		}

		return nil
	})
}
