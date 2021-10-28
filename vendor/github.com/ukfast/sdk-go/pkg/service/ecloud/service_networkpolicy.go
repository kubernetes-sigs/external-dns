package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetNetworkPolicies retrieves a list of network policies
func (s *Service) GetNetworkPolicies(parameters connection.APIRequestParameters) ([]NetworkPolicy, error) {
	var policys []NetworkPolicy

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkPoliciesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedNetworkPolicy).Items {
			policys = append(policys, policy)
		}
	}

	return policys, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworkPoliciesPaginated retrieves a paginated list of network policies
func (s *Service) GetNetworkPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkPolicy, error) {
	body, err := s.getNetworkPoliciesPaginatedResponseBody(parameters)

	return NewPaginatedNetworkPolicy(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkPoliciesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkPoliciesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetNetworkPolicySliceResponseBody, error) {
	body := &GetNetworkPolicySliceResponseBody{}

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

func (s *Service) getNetworkPolicyResponseBody(policyID string) (*GetNetworkPolicyResponseBody, error) {
	body := &GetNetworkPolicyResponseBody{}

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

func (s *Service) createNetworkPolicyResponseBody(req CreateNetworkPolicyRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) patchNetworkPolicyResponseBody(policyID string, req PatchNetworkPolicyRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) deleteNetworkPolicyResponseBody(policyID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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
	var networkRules []NetworkRule

	return networkRules, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
		},
		func(response connection.Paginated) {
			for _, networkRule := range response.(*PaginatedNetworkRule).Items {
				networkRules = append(networkRules, networkRule)
			}
		},
		parameters,
	)
}

// GetNetworkPolicyNetworkRulesPaginated retrieves a paginated list of network policy NetworkRules
func (s *Service) GetNetworkPolicyNetworkRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedNetworkRule, error) {
	body, err := s.getNetworkPolicyNetworkRulesPaginatedResponseBody(policyID, parameters)

	return NewPaginatedNetworkRule(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkPolicyNetworkRulesPaginated(policyID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkPolicyNetworkRulesPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*GetNetworkRuleSliceResponseBody, error) {
	body := &GetNetworkRuleSliceResponseBody{}

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
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworkPolicyTasksPaginated retrieves a paginated list of NetworkPolicy tasks
func (s *Service) GetNetworkPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getNetworkPolicyTasksPaginatedResponseBody(policyID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkPolicyTasksPaginated(policyID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkPolicyTasksPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

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
