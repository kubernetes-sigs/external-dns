package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewallPolicies retrieves a list of firewall policies
func (s *Service) GetFirewallPolicies(parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	return connection.InvokeRequestAll(s.GetFirewallPoliciesPaginated, parameters)
}

// GetFirewallPoliciesPaginated retrieves a paginated list of firewall policies
func (s *Service) GetFirewallPoliciesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	body, err := s.getFirewallPoliciesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallPoliciesPaginated), err
}

func (s *Service) getFirewallPoliciesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallPolicy], error) {
	body := &connection.APIResponseBodyData[[]FirewallPolicy]{}

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

func (s *Service) getFirewallPolicyResponseBody(policyID string) (*connection.APIResponseBodyData[FirewallPolicy], error) {
	body := &connection.APIResponseBodyData[FirewallPolicy]{}

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

func (s *Service) createFirewallPolicyResponseBody(req CreateFirewallPolicyRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) patchFirewallPolicyResponseBody(policyID string, req PatchFirewallPolicyRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) deleteFirewallPolicyResponseBody(policyID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
		return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
	}, parameters)
}

// GetFirewallPolicyFirewallRulesPaginated retrieves a paginated list of firewall policy FirewallRules
func (s *Service) GetFirewallPolicyFirewallRulesPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
	body, err := s.getFirewallPolicyFirewallRulesPaginatedResponseBody(policyID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
		return s.GetFirewallPolicyFirewallRulesPaginated(policyID, p)
	}), err
}

func (s *Service) getFirewallPolicyFirewallRulesPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallRule], error) {
	body := &connection.APIResponseBodyData[[]FirewallRule]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}, parameters)
}

// GetFirewallPolicyTasksPaginated retrieves a paginated list of FirewallPolicy tasks
func (s *Service) GetFirewallPolicyTasksPaginated(policyID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getFirewallPolicyTasksPaginatedResponseBody(policyID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetFirewallPolicyTasksPaginated(policyID, p)
	}), err
}

func (s *Service) getFirewallPolicyTasksPaginatedResponseBody(policyID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

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
