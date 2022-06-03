package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetFirewallRulePorts retrieves a list of firewall rules
func (s *Service) GetFirewallRulePorts(parameters connection.APIRequestParameters) ([]FirewallRulePort, error) {
	return connection.InvokeRequestAll(s.GetFirewallRulePortsPaginated, parameters)
}

// GetFirewallRulePortsPaginated retrieves a paginated list of firewall rules
func (s *Service) GetFirewallRulePortsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[FirewallRulePort], error) {
	body, err := s.getFirewallRulePortsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetFirewallRulePortsPaginated), err
}

func (s *Service) getFirewallRulePortsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallRulePort], error) {
	body := &connection.APIResponseBodyData[[]FirewallRulePort]{}

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

func (s *Service) getFirewallRulePortResponseBody(ruleID string) (*connection.APIResponseBodyData[FirewallRulePort], error) {
	body := &connection.APIResponseBodyData[FirewallRulePort]{}

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

func (s *Service) createFirewallRulePortResponseBody(req CreateFirewallRulePortRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) patchFirewallRulePortResponseBody(ruleID string, req PatchFirewallRulePortRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) deleteFirewallRulePortResponseBody(ruleID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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
