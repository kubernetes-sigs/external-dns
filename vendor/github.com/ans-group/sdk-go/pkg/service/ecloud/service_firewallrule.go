package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewallRules retrieves a list of firewall rules
func (s *Service) GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	return connection.InvokeRequestAll(s.GetFirewallRulesPaginated, parameters)
}

// GetFirewallRulesPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRule], error) {
	body, err := s.getFirewallRulesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallRulesPaginated), err
}

func (s *Service) getFirewallRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallRule], error) {
	body := &connection.APIResponseBodyData[[]FirewallRule]{}

	response, err := s.connection.Get("/ecloud/v2/firewall-rules", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewallRule retrieves a single rule by id
func (s *Service) GetFirewallRule(ruleID string) (FirewallRule, error) {
	body, err := s.getFirewallRuleResponseBody(ruleID)

	return body.Data, err
}

func (s *Service) getFirewallRuleResponseBody(ruleID string) (*connection.APIResponseBodyData[FirewallRule], error) {
	body := &connection.APIResponseBodyData[FirewallRule]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateFirewallRule creates a new FirewallRule
func (s *Service) CreateFirewallRule(req CreateFirewallRuleRequest) (TaskReference, error) {
	body, err := s.createFirewallRuleResponseBody(req)

	return body.Data, err
}

func (s *Service) createFirewallRuleResponseBody(req CreateFirewallRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/firewall-rules", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchFirewallRule patches a FirewallRule
func (s *Service) PatchFirewallRule(ruleID string, req PatchFirewallRuleRequest) (TaskReference, error) {
	body, err := s.patchFirewallRuleResponseBody(ruleID, req)

	return body.Data, err
}

func (s *Service) patchFirewallRuleResponseBody(ruleID string, req PatchFirewallRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteFirewallRule deletes a FirewallRule
func (s *Service) DeleteFirewallRule(ruleID string) (string, error) {
	body, err := s.deleteFirewallRuleResponseBody(ruleID)

	return body.Data.TaskID, err
}

func (s *Service) deleteFirewallRuleResponseBody(ruleID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/firewall-rules/%s", ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// GetFirewallRuleFirewallRulePorts retrieves a list of firewall rule ports
func (s *Service) GetFirewallRuleFirewallRulePorts(firewallRuleID string, parameters connection.APIRequestParameters) ([]FirewallRulePort, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}, parameters)
}

// GetFirewallRuleFirewallRulePortsPaginated retrieves a paginated list of firewall rule ports
func (s *Service) GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
	body, err := s.getFirewallRuleFirewallRulePortsPaginatedResponseBody(firewallRuleID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}), err
}

func (s *Service) getFirewallRuleFirewallRulePortsPaginatedResponseBody(firewallRuleID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallRulePort], error) {
	body := &connection.APIResponseBodyData[[]FirewallRulePort]{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rules/%s/ports", firewallRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
