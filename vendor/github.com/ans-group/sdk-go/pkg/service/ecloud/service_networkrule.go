package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworkRules retrieves a list of network rules
func (s *Service) GetNetworkRules(parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	return connection.InvokeRequestAll(s.GetNetworkRulesPaginated, parameters)
}

// GetNetworkRulesPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRule], error) {
	body, err := s.getNetworkRulesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworkRulesPaginated), err
}

func (s *Service) getNetworkRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NetworkRule], error) {
	body := &connection.APIResponseBodyData[[]NetworkRule]{}

	response, err := s.connection.Get("/ecloud/v2/network-rules", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetworkRule retrieves a single rule by id
func (s *Service) GetNetworkRule(ruleID string) (NetworkRule, error) {
	body, err := s.getNetworkRuleResponseBody(ruleID)

	return body.Data, err
}

func (s *Service) getNetworkRuleResponseBody(ruleID string) (*connection.APIResponseBodyData[NetworkRule], error) {
	body := &connection.APIResponseBodyData[NetworkRule]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateNetworkRule creates a new NetworkRule
func (s *Service) CreateNetworkRule(req CreateNetworkRuleRequest) (TaskReference, error) {
	body, err := s.createNetworkRuleResponseBody(req)

	return body.Data, err
}

func (s *Service) createNetworkRuleResponseBody(req CreateNetworkRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/network-rules", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchNetworkRule patches a NetworkRule
func (s *Service) PatchNetworkRule(ruleID string, req PatchNetworkRuleRequest) (TaskReference, error) {
	body, err := s.patchNetworkRuleResponseBody(ruleID, req)

	return body.Data, err
}

func (s *Service) patchNetworkRuleResponseBody(ruleID string, req PatchNetworkRuleRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteNetworkRule deletes a NetworkRule
func (s *Service) DeleteNetworkRule(ruleID string) (string, error) {
	body, err := s.deleteNetworkRuleResponseBody(ruleID)

	return body.Data.TaskID, err
}

func (s *Service) deleteNetworkRuleResponseBody(ruleID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/network-rules/%s", ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// GetNetworkRuleNetworkRulePorts retrieves a list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePorts(networkRuleID string, parameters connection.APIRequestParameters) ([]NetworkRulePort, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}, parameters)
}

// GetNetworkRuleNetworkRulePortsPaginated retrieves a paginated list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePortsPaginated(networkRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
	body, err := s.getNetworkRuleNetworkRulePortsPaginatedResponseBody(networkRuleID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NetworkRulePort], error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}), err
}

func (s *Service) getNetworkRuleNetworkRulePortsPaginatedResponseBody(networkRuleID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NetworkRulePort], error) {
	body := &connection.APIResponseBodyData[[]NetworkRulePort]{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-rules/%s/ports", networkRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
