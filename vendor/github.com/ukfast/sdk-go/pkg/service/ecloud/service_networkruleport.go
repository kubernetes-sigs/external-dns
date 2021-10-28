package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetNetworkRulePorts retrieves a list of network rules
func (s *Service) GetNetworkRulePorts(parameters connection.APIRequestParameters) ([]NetworkRulePort, error) {
	var rules []NetworkRulePort

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRulePortsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, rule := range response.(*PaginatedNetworkRulePort).Items {
			rules = append(rules, rule)
		}
	}

	return rules, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetNetworkRulePortsPaginated retrieves a paginated list of network rules
func (s *Service) GetNetworkRulePortsPaginated(parameters connection.APIRequestParameters) (*PaginatedNetworkRulePort, error) {
	body, err := s.getNetworkRulePortsPaginatedResponseBody(parameters)

	return NewPaginatedNetworkRulePort(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetNetworkRulePortsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getNetworkRulePortsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetNetworkRulePortSliceResponseBody, error) {
	body := &GetNetworkRulePortSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/network-rule-ports", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetworkRulePort retrieves a single rule by id
func (s *Service) GetNetworkRulePort(ruleID string) (NetworkRulePort, error) {
	body, err := s.getNetworkRulePortResponseBody(ruleID)

	return body.Data, err
}

func (s *Service) getNetworkRulePortResponseBody(ruleID string) (*GetNetworkRulePortResponseBody, error) {
	body := &GetNetworkRulePortResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/network-rule-ports/%s", ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateNetworkRulePort creates a new NetworkRulePort
func (s *Service) CreateNetworkRulePort(req CreateNetworkRulePortRequest) (TaskReference, error) {
	body, err := s.createNetworkRulePortResponseBody(req)

	return body.Data, err
}

func (s *Service) createNetworkRulePortResponseBody(req CreateNetworkRulePortRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/network-rule-ports", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchNetworkRulePort patches a NetworkRulePort
func (s *Service) PatchNetworkRulePort(ruleID string, req PatchNetworkRulePortRequest) (TaskReference, error) {
	body, err := s.patchNetworkRulePortResponseBody(ruleID, req)

	return body.Data, err
}

func (s *Service) patchNetworkRulePortResponseBody(ruleID string, req PatchNetworkRulePortRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/network-rule-ports/%s", ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteNetworkRulePort deletes a NetworkRulePort
func (s *Service) DeleteNetworkRulePort(ruleID string) (string, error) {
	body, err := s.deleteNetworkRulePortResponseBody(ruleID)

	return body.Data.TaskID, err
}

func (s *Service) deleteNetworkRulePortResponseBody(ruleID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if ruleID == "" {
		return body, fmt.Errorf("invalid network rule id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/network-rule-ports/%s", ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkRulePortNotFoundError{ID: ruleID}
		}

		return nil
	})
}
