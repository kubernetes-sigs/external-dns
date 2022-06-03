package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworkPolicies retrieves a list of network policies
func (s *Service) GetNetworkPolicies(parameters connection.APIRequestParameters) ([]NetworkPolicy, error) {
	return connection.InvokeRequestAll(s.GetNetworkPoliciesPaginated, parameters)
}

// GetNetworkPoliciesPaginated retrieves a paginated list of network policies
func (s *Service) GetNetworkPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkPolicy], error) {
	body, err := s.getNetworkPoliciesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworkPoliciesPaginated), err
}

func (s *Service) getNetworkPoliciesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NetworkPolicy], error) {
	body := &connection.APIResponseBodyData[[]NetworkPolicy]{}

	response, err := s.connection.Get("/ecloud/v2/network-policies", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetworkPolicy retrieves a single network policy by id
func (s *Service) GetNetworkPolicy(policyID string) (NetworkPolicy, error) {
	body, err := s.getNetworkPolicyResponseBody(policyID)

	return body.Data, err
}

func (s *Service) getNetworkPolicyResponseBody(policyID string) (*connection.APIResponseBodyData[NetworkPolicy], error) {
	body := &connection.APIResponseBodyData[NetworkPolicy]{}

	if policyID == "" {
		return body, fmt.Errorf("invalid network policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// CreateNetworkPolicy creates a new NetworkPolicy
func (s *Service) CreateNetworkPolicy(req CreateNetworkPolicyRequest) (TaskReference, error) {
	body, err := s.createNetworkPolicyResponseBody(req)

	return body.Data, err
}

func (s *Service) createNetworkPolicyResponseBody(req CreateNetworkPolicyRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/network-policies", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchNetworkPolicy patches a NetworkPolicy
func (s *Service) PatchNetworkPolicy(policyID string, req PatchNetworkPolicyRequest) (TaskReference, error) {
	body, err := s.patchNetworkPolicyResponseBody(policyID, req)

	return body.Data, err
}

func (s *Service) patchNetworkPolicyResponseBody(policyID string, req PatchNetworkPolicyRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// DeleteNetworkPolicy deletes a NetworkPolicy
func (s *Service) DeleteNetworkPolicy(policyID string) (string, error) {
	body, err := s.deleteNetworkPolicyResponseBody(policyID)

	return body.Data.TaskID, err
}

func (s *Service) deleteNetworkPolicyResponseBody(policyID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/network-policies/%s", policyID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// GetNetworkPolicyNetworkRules retrieves a list of network policy rules
func (s *Service) GetNetworkPolicyNetworkRules(policyID string, parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
		return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
	}, parameters)
}

// GetNetworkPolicyNetworkRulesPaginated retrieves a paginated list of network policy NetworkRules
func (s *Service) GetNetworkPolicyNetworkRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	body, err := s.getNetworkPolicyNetworkRulesPaginatedResponseBody(policyID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
		return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
	}), err
}

func (s *Service) getNetworkPolicyNetworkRulesPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NetworkRule], error) {
	body := &connection.APIResponseBodyData[[]NetworkRule]{}

	if policyID == "" {
		return body, fmt.Errorf("invalid network policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-policies/%s/network-rules", policyID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// GetNetworkPolicyTasks retrieves a list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}, parameters)
}

// GetNetworkPolicyTasksPaginated retrieves a paginated list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getNetworkPolicyTasksPaginatedResponseBody(policyID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}), err
}

func (s *Service) getNetworkPolicyTasksPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if policyID == "" {
		return body, fmt.Errorf("invalid network policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-policies/%s/tasks", policyID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}
