package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTargetGroups retrieves a list of target groups
func (s *Service) GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error) {
	var groups []TargetGroup

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, group := range response.(*PaginatedTargetGroup).Items {
			groups = append(groups, group)
		}
	}

	return groups, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTargetGroupsPaginated retrieves a paginated list of target groups
func (s *Service) GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedTargetGroup, error) {
	body, err := s.getTargetGroupsPaginatedResponseBody(parameters)

	return NewPaginatedTargetGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTargetGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetTargetGroupSliceResponseBody, error) {
	body := &GetTargetGroupSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/target-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTargetGroup retrieves a single target group by id
func (s *Service) GetTargetGroup(groupID int) (TargetGroup, error) {
	body, err := s.getTargetGroupResponseBody(groupID)

	return body.Data, err
}

func (s *Service) getTargetGroupResponseBody(groupID int) (*GetTargetGroupResponseBody, error) {
	body := &GetTargetGroupResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetGroupNotFoundError{ID: groupID}
		}

		return nil
	})
}

// PatchTargetGroup patches a target group
func (s *Service) CreateTargetGroup(req CreateTargetGroupRequest) (int, error) {
	body, err := s.createTargetGroupResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createTargetGroupResponseBody(req CreateTargetGroupRequest) (*GetTargetGroupResponseBody, error) {
	body := &GetTargetGroupResponseBody{}

	response, err := s.connection.Post("/loadbalancers/v2/target-groups", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}

// PatchTargetGroup patches a target group
func (s *Service) PatchTargetGroup(groupID int, req PatchTargetGroupRequest) error {
	_, err := s.patchTargetGroupResponseBody(groupID, req)

	return err
}

func (s *Service) patchTargetGroupResponseBody(groupID int, req PatchTargetGroupRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetGroupNotFoundError{ID: groupID}
		}

		return nil
	})
}

// PatchTargetGroup patches a target group
func (s *Service) DeleteTargetGroup(groupID int) error {
	_, err := s.deleteTargetGroupResponseBody(groupID)

	return err
}

func (s *Service) deleteTargetGroupResponseBody(groupID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/target-groups/%d", groupID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetGroupNotFoundError{ID: groupID}
		}

		return nil
	})
}
