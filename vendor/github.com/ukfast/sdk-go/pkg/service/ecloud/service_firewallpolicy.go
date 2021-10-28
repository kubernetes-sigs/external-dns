package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallPolicies retrieves a list of firewall policies
func (s *Service) GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	var policys []FirewallPolicy

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPoliciesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedFirewallPolicy).Items {
			policys = append(policys, policy)
		}
	}

	return policys, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallPoliciesPaginated retrieves a paginated list of firewall policies
func (s *Service) GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error) {
	body, err := s.getFirewallPoliciesPaginatedResponseBody(parameters)

	return NewPaginatedFirewallPolicy(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPoliciesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallPoliciesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallPolicySliceResponseBody, error) {
	body := &GetFirewallPolicySliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/firewall-policies", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewallPolicy retrieves a single firewall policy by id
func (s *Service) GetFirewallPolicy(policyID string) (FirewallPolicy, error) {
	body, err := s.getFirewallPolicyResponseBody(policyID)

	return body.Data, err
}

func (s *Service) getFirewallPolicyResponseBody(policyID string) (*GetFirewallPolicyResponseBody, error) {
	body := &GetFirewallPolicyResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid firewall policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// CreateFirewallPolicy creates a new FirewallPolicy
func (s *Service) CreateFirewallPolicy(req CreateFirewallPolicyRequest) (TaskReference, error) {
	body, err := s.createFirewallPolicyResponseBody(req)

	return body.Data, err
}

func (s *Service) createFirewallPolicyResponseBody(req CreateFirewallPolicyRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/firewall-policies", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchFirewallPolicy patches a FirewallPolicy
func (s *Service) PatchFirewallPolicy(policyID string, req PatchFirewallPolicyRequest) (TaskReference, error) {
	body, err := s.patchFirewallPolicyResponseBody(policyID, req)

	return body.Data, err
}

func (s *Service) patchFirewallPolicyResponseBody(policyID string, req PatchFirewallPolicyRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// DeleteFirewallPolicy deletes a FirewallPolicy
func (s *Service) DeleteFirewallPolicy(policyID string) (string, error) {
	body, err := s.deleteFirewallPolicyResponseBody(policyID)

	return body.Data.TaskID, err
}

func (s *Service) deleteFirewallPolicyResponseBody(policyID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid policy id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/firewall-policies/%s", policyID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// GetFirewallPolicyFirewallRules retrieves a list of firewall policy rules
func (s *Service) GetFirewallPolicyFirewallRules(policyID string, parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	var firewallRules []FirewallRule

	return firewallRules, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
		},
		func(response connection.Paginated) {
			for _, firewallRule := range response.(*PaginatedFirewallRule).Items {
				firewallRules = append(firewallRules, firewallRule)
			}
		},
		parameters,
	)
}

// GetFirewallPolicyFirewallRulesPaginated retrieves a paginated list of firewall policy FirewallRules
func (s *Service) GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error) {
	body, err := s.getFirewallPolicyFirewallRulesPaginatedResponseBody(policyID, parameters)

	return NewPaginatedFirewallRule(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallPolicyFirewallRulesPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*GetFirewallRuleSliceResponseBody, error) {
	body := &GetFirewallRuleSliceResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid firewall policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-policies/%s/firewall-rules", policyID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}

// GetFirewallPolicyTasks retrieves a list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasks(policyID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallPolicyTasksPaginated retrieves a paginated list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getFirewallPolicyTasksPaginatedResponseBody(policyID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallPolicyTasksPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if policyID == "" {
		return body, fmt.Errorf("invalid firewall policy id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-policies/%s/tasks", policyID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallPolicyNotFoundError{ID: policyID}
		}

		return nil
	})
}
