package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetFirewallRulePorts retrieves a list of firewall rules
func (s *Service) GetFirewallRulePorts(parameters connection.APIRequestParameters) ([]FirewallRulePort, error) {
	var rules []FirewallRulePort

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulePortsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, rule := range response.(*PaginatedFirewallRulePort).Items {
			rules = append(rules, rule)
		}
	}

	return rules, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetFirewallRulePortsPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulePortsPaginated(parameters connection.APIRequestParameters) (*PaginatedFirewallRulePort, error) {
	body, err := s.getFirewallRulePortsPaginatedResponseBody(parameters)

	return NewPaginatedFirewallRulePort(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetFirewallRulePortsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getFirewallRulePortsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetFirewallRulePortSliceResponseBody, error) {
	body := &GetFirewallRulePortSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/firewall-rule-ports", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetFirewallRulePort retrieves a single rule by id
func (s *Service) GetFirewallRulePort(ruleID string) (FirewallRulePort, error) {
	body, err := s.getFirewallRulePortResponseBody(ruleID)

	return body.Data, err
}

func (s *Service) getFirewallRulePortResponseBody(ruleID string) (*GetFirewallRulePortResponseBody, error) {
	body := &GetFirewallRulePortResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/firewall-rule-ports/%s", ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateFirewallRulePort creates a new FirewallRulePort
func (s *Service) CreateFirewallRulePort(req CreateFirewallRulePortRequest) (TaskReference, error) {
	body, err := s.createFirewallRulePortResponseBody(req)

	return body.Data, err
}

func (s *Service) createFirewallRulePortResponseBody(req CreateFirewallRulePortRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/firewall-rule-ports", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchFirewallRulePort patches a FirewallRulePort
func (s *Service) PatchFirewallRulePort(ruleID string, req PatchFirewallRulePortRequest) (TaskReference, error) {
	body, err := s.patchFirewallRulePortResponseBody(ruleID, req)

	return body.Data, err
}

func (s *Service) patchFirewallRulePortResponseBody(ruleID string, req PatchFirewallRulePortRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/firewall-rule-ports/%s", ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteFirewallRulePort deletes a FirewallRulePort
func (s *Service) DeleteFirewallRulePort(ruleID string) (string, error) {
	body, err := s.deleteFirewallRulePortResponseBody(ruleID)

	return body.Data.TaskID, err
}

func (s *Service) deleteFirewallRulePortResponseBody(ruleID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid firewall rule id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/firewall-rule-ports/%s", ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &FirewallRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}
