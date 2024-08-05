package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNATOverloadRules retrieves a list of NAT overload rules
func (s *Service) GetNATOverloadRules(parameters connection.APIRequestParameters) ([]NATOverloadRule, error) {
	return connection.InvokeRequestAll(s.GetNATOverloadRulesPaginated, parameters)
}

// GetNATOverloadRulesPaginated retrieves a paginated list of NAT overload rules
func (s *Service) GetNATOverloadRulesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[NATOverloadRule], error) {
	body, err := s.getNATOverloadRulesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetNATOverloadRulesPaginated), err
}

func (s *Service) getNATOverloadRulesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NATOverloadRule], error) {
	return connection.Get[[]NATOverloadRule](s.connection, "/ecloud/v2/nat-overload-rules", parameters)
}

// GetNATOverloadRule retrieves a single NAT overload rule by id
func (s *Service) GetNATOverloadRule(ruleID string) (NATOverloadRule, error) {
	body, err := s.getNATOverloadRuleResponseBody(ruleID)
	if err != nil {
		return NATOverloadRule{}, err
	}

	return body.Data, err
}

func (s *Service) getNATOverloadRuleResponseBody(ruleID string) (*connection.APIResponseBodyData[NATOverloadRule], error) {
	if ruleID == "" {
		return nil, fmt.Errorf("invalid nat overload rule id")
	}

	return connection.Get[NATOverloadRule](s.connection, fmt.Sprintf("/ecloud/v2/nat-overload-rules/%s", ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&NATOverloadRuleNotFoundError{ID: ruleID}))
}

// CreateNATOverloadRule creates a new NAT overload
func (s *Service) CreateNATOverloadRule(req CreateNATOverloadRuleRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/nat-overload-rules", &req)
	if err != nil {
		return TaskReference{}, err
	}

	return body.Data, err
}

// PatchNATOverloadRule patches a NAT overload
func (s *Service) PatchNATOverloadRule(ruleID string, req PatchNATOverloadRuleRequest) (TaskReference, error) {
	if ruleID == "" {
		return TaskReference{}, fmt.Errorf("invalid nat overload rule id")
	}

	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nat-overload-rules/%s", ruleID), &req, connection.NotFoundResponseHandler(&NATOverloadRuleNotFoundError{ID: ruleID}))
	if err != nil {
		return TaskReference{}, err
	}

	return body.Data, err
}

// DeleteNATOverloadRule deletes a NAT overload
func (s *Service) DeleteNATOverloadRule(ruleID string) (string, error) {
	if ruleID == "" {
		return "", fmt.Errorf("invalid nat overload rule id")
	}

	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/nat-overload-rules/%s", ruleID), nil, connection.NotFoundResponseHandler(&NATOverloadRuleNotFoundError{ID: ruleID}))
	if err != nil {
		return "", err
	}

	return body.Data.TaskID, err
}
