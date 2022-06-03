package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTargetGroups retrieves a list of target groups
func (s *Service) GetTargetGroups(parameters connection.APIRequestParameters) ([]TargetGroup, error) {
	return connection.InvokeRequestAll(s.GetTargetGroupsPaginated, parameters)
}

// GetTargetGroupsPaginated retrieves a paginated list of target groups
func (s *Service) GetTargetGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[TargetGroup], error) {
	body, err := s.getTargetGroupsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetTargetGroupsPaginated), err
}

func (s *Service) getTargetGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]TargetGroup], error) {
	body := &connection.APIResponseBodyData[[]TargetGroup]{}

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

func (s *Service) getTargetGroupResponseBody(groupID int) (*connection.APIResponseBodyData[TargetGroup], error) {
	body := &connection.APIResponseBodyData[TargetGroup]{}

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

func (s *Service) createTargetGroupResponseBody(req CreateTargetGroupRequest) (*connection.APIResponseBodyData[TargetGroup], error) {
	body := &connection.APIResponseBodyData[TargetGroup]{}

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
