package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetNetworkRules retrieves a list of network rules
func (s *Service) GetNetworkRules(parameters connection.APIRequestParameters) ([]NetworkRule, error) {
	var rules []NetworkRule

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRulesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, rule := range response.(*PaginatedNetworkRule).Items {
			rules = append(rules, rule)
		}
	}

	return rules, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworkRulesPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulesPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkRule, error) {
	body, err := s.getNetworkRulesPaginatedResponseBody(parameters)

	return NewPaginatedNetworkRule(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRulesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetNetworkRuleSliceResponseBody, error) {
	body := &GetNetworkRuleSliceResponseBody{}

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

func (s *Service) getNetworkRuleResponseBody(ruleID string) (*GetNetworkRuleResponseBody, error) {
	body := &GetNetworkRuleResponseBody{}

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

func (s *Service) createNetworkRuleResponseBody(req CreateNetworkRuleRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) patchNetworkRuleResponseBody(ruleID string, req PatchNetworkRuleRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) deleteNetworkRuleResponseBody(ruleID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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
	var ports []NetworkRulePort

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, port := range response.(*PaginatedNetworkRulePort).Items {
			ports = append(ports, port)
		}
	}

	return ports, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworkRuleNetworkRulePortsPaginated retrieves a paginated list of network rule ports
func (s *Service) GetNetworkRuleNetworkRulePortsPaginated(networkRuleID string, parameters connection.APIRequestParameters) (*PaginatedNetworkRulePort, error) {
	body, err := s.getNetworkRuleNetworkRulePortsPaginatedResponseBody(networkRuleID, parameters)

	return NewPaginatedNetworkRulePort(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRuleNetworkRulePortsPaginated(networkRuleID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkRuleNetworkRulePortsPaginatedResponseBody(networkRuleID string, parameters connection.APIRequestParameters) (*GetNetworkRulePortSliceResponseBody, error) {
	body := &GetNetworkRulePortSliceResponseBody{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-rules/%s/ports", networkRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
