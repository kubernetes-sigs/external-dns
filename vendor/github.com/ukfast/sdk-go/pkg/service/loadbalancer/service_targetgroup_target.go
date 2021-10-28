package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTargetGroupTargets retrieves a list of targets
func (s *Service) GetTargetGroupTargets(groupID int, parameters connection.APIRequestParameters) ([]Target, error) {
	var targets []Target

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupTargetsPaginated(groupID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, target := range response.(*PaginatedTarget).Items {
			targets = append(targets, target)
		}
	}

	return targets, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTargetGroupTargetsPaginated retrieves a paginated list of targets
func (s *Service) GetTargetGroupTargetsPaginated(groupID int, parameters connection.APIRequestParameters) (*PaginatedTarget, error) {
	body, err := s.getTargetGroupTargetsPaginatedResponseBody(groupID, parameters)

	return NewPaginatedTarget(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupTargetsPaginated(groupID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTargetGroupTargetsPaginatedResponseBody(groupID int, parameters connection.APIRequestParameters) (*GetTargetSliceResponseBody, error) {
	body := &GetTargetSliceResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets", groupID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTargetGroupTarget retrieves a single target by id
func (s *Service) GetTargetGroupTarget(groupID int, targetID int) (Target, error) {
	body, err := s.getTargetGroupTargetResponseBody(groupID, targetID)

	return body.Data, err
}

func (s *Service) getTargetGroupTargetResponseBody(groupID int, targetID int) (*GetTargetResponseBody, error) {
	body := &GetTargetResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	if targetID < 1 {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: groupID}
		}

		return nil
	})
}

// CreateTargetGroupTarget creates a target
func (s *Service) CreateTargetGroupTarget(groupID int, req CreateTargetRequest) (int, error) {
	body, err := s.createTargetGroupTargetResponseBody(groupID, req)

	return body.Data.ID, err
}

func (s *Service) createTargetGroupTargetResponseBody(groupID int, req CreateTargetRequest) (*GetTargetResponseBody, error) {
	body := &GetTargetResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets", groupID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: groupID}
		}

		return nil
	})
}

// PatchTargetGroupTarget patches a target
func (s *Service) PatchTargetGroupTarget(groupID int, targetID int, req PatchTargetRequest) error {
	_, err := s.patchTargetGroupTargetResponseBody(groupID, targetID, req)

	return err
}

func (s *Service) patchTargetGroupTargetResponseBody(groupID int, targetID int, req PatchTargetRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	if targetID < 1 {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: groupID}
		}

		return nil
	})
}

// DeleteTargetGroupTarget deletes a target
func (s *Service) DeleteTargetGroupTarget(groupID int, targetID int) error {
	_, err := s.deleteTargetGroupTargetResponseBody(groupID, targetID)

	return err
}

func (s *Service) deleteTargetGroupTargetResponseBody(groupID int, targetID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if groupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	if targetID < 1 {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/targets/%d", groupID, targetID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: groupID}
		}

		return nil
	})
}
