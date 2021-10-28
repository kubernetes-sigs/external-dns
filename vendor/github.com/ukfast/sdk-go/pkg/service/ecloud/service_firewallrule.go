package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallRules retrieves a list of firewall rules
func (s *Service) GetFirewallRules(parameters connection.APIRequestParameters) ([]FirewallRule, error) {
	var rules []FirewallRule

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, rule := range response.(*PaginatedFirewallRule).Items {
			rules = append(rules, rule)
		}
	}

	return rules, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallRulesPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRule, error) {
	body, err := s.getFirewallRulesPaginatedResponseBody(parameters)

	return NewPaginatedFirewallRule(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallRuleSliceResponseBody, error) {
	body := &GetFirewallRuleSliceResponseBody{}

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

func (s *Service) getFirewallRuleResponseBody(ruleID string) (*GetFirewallRuleResponseBody, error) {
	body := &GetFirewallRuleResponseBody{}

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

func (s *Service) createFirewallRuleResponseBody(req CreateFirewallRuleRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) patchFirewallRuleResponseBody(ruleID string, req PatchFirewallRuleRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) deleteFirewallRuleResponseBody(ruleID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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
	var ports []FirewallRulePort

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, port := range response.(*PaginatedFirewallRulePort).Items {
			ports = append(ports, port)
		}
	}

	return ports, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallRuleFirewallRulePortsPaginated retrieves a paginated list of firewall rule ports
func (s *Service) GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID string, parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error) {
	body, err := s.getFirewallRuleFirewallRulePortsPaginatedResponseBody(firewallRuleID, parameters)

	return NewPaginatedFirewallRulePort(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRuleFirewallRulePortsPaginated(firewallRuleID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallRuleFirewallRulePortsPaginatedResponseBody(firewallRuleID string, parameters connection.APIRequestParameters) (*GetFirewallRulePortSliceResponseBody, error) {
	body := &GetFirewallRulePortSliceResponseBody{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rules/%s/ports", firewallRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
