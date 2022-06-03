package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAffinityRules retrieves a list of affinity rules
func (s *Service) GetAffinityRules(parameters connection.APIRequestParameters) ([]AffinityRule, error) {
	return connection.InvokeRequestAll(s.GetAffinityRulesPaginated, parameters)
}

// GetAffinityRulesPaginated retrieves a paginated list of affinity rules
func (s *Service) GetAffinityRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRule], error) {
	body, err := s.getAffinityRulesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetAffinityRulesPaginated), err
}

func (s *Service) getAffinityRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AffinityRule], error) {
	body := &connection.APIResponseBodyData[[]AffinityRule]{}

	response, err := s.connection.Get("/ecloud/v2/affinity-rules", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAffinityRule retrieves a single AffinityRule by id
func (s *Service) GetAffinityRule(affinityruleID string) (AffinityRule, error) {
	body, err := s.getAffinityRuleResponseBody(affinityruleID)

	return body.Data, err
}

func (s *Service) getAffinityRuleResponseBody(affinityruleID string) (*connection.APIResponseBodyData[AffinityRule], error) {
	body := &connection.APIResponseBodyData[AffinityRule]{}

	if affinityruleID == "" {
		return body, fmt.Errorf("invalid affinity rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleNotFoundError{ID: affinityruleID}
		}

		return nil
	})
}

// CreateAffinityRule creates a new AffinityRule
func (s *Service) CreateAffinityRule(req CreateAffinityRuleRequest) (TaskReference, error) {
	body, err := s.createAffinityRuleResponseBody(req)

	return body.Data, err
}

func (s *Service) createAffinityRuleResponseBody(req CreateAffinityRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/affinity-rules", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchAffinityRule patches a AffinityRule
func (s *Service) PatchAffinityRule(affinityruleID string, req PatchAffinityRuleRequest) (TaskReference, error) {
	body, err := s.patchAffinityRuleResponseBody(affinityruleID, req)

	return body.Data, err
}

func (s *Service) patchAffinityRuleResponseBody(affinityruleID string, req PatchAffinityRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if affinityruleID == "" {
		return body, fmt.Errorf("invalid affinity rule id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleNotFoundError{ID: affinityruleID}
		}

		return nil
	})
}

// DeleteAffinityRule deletes a AffinityRule
func (s *Service) DeleteAffinityRule(affinityruleID string) (string, error) {
	body, err := s.deleteAffinityRuleResponseBody(affinityruleID)

	return body.Data.TaskID, err
}

func (s *Service) deleteAffinityRuleResponseBody(affinityruleID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if affinityruleID == "" {
		return body, fmt.Errorf("invalid affinity rule id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/affinity-rules/%s", affinityruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleNotFoundError{ID: affinityruleID}
		}

		return nil
	})
}
