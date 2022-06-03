package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAffinityRuleMembers retrieves a list of affinity rule members
func (s *Service) GetAffinityRuleMembers(affinityRuleID string, parameters connection.APIRequestParameters) ([]AffinityRuleMember, error) {
	return connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
			return s.GetAffinityRuleMembersPaginated(affinityRuleID, p)
		}, parameters)
}

// GetAffinityRuleMembersPaginated retrieves a paginated list of affinity rule members
func (s *Service) GetAffinityRuleMembersPaginated(affinityRuleID string, parameters connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
	body, err := s.getAffinityRuleMembersPaginatedResponseBody(affinityRuleID, parameters)
	return connection.NewPaginated(
		body,
		parameters,
		func(p connection.APIRequestParameters) (*connection.Paginated[AffinityRuleMember], error) {
			return s.GetAffinityRuleMembersPaginated(affinityRuleID, p)
		}), err
}

func (s *Service) getAffinityRuleMembersPaginatedResponseBody(affinityRuleID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AffinityRuleMember], error) {
	body := &connection.APIResponseBodyData[[]AffinityRuleMember]{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/affinity-rules/%s/members", affinityRuleID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAffinityRuleMember retrieves a single AffinityRuleMember by id
func (s *Service) GetAffinityRuleMember(memberID string) (AffinityRuleMember, error) {
	body, err := s.getAffinityRuleMemberResponseBody(memberID)

	return body.Data, err
}

func (s *Service) getAffinityRuleMemberResponseBody(memberID string) (*connection.APIResponseBodyData[AffinityRuleMember], error) {
	body := &connection.APIResponseBodyData[AffinityRuleMember]{}

	if memberID == "" {
		return body, fmt.Errorf("invalid affinity rule member id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/affinity-rule-members/%s", memberID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleMemberNotFoundError{ID: memberID}
		}

		return nil
	})
}

// CreateAffinityRuleMember creates a new AffinityRuleMember
func (s *Service) CreateAffinityRuleMember(req CreateAffinityRuleMemberRequest) (TaskReference, error) {
	body, err := s.createAffinityRuleMemberResponseBody(req)

	return body.Data, err
}

func (s *Service) createAffinityRuleMemberResponseBody(req CreateAffinityRuleMemberRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/affinity-rule-members", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteAffinityRuleMember deletes a AffinityRuleMember
func (s *Service) DeleteAffinityRuleMember(memberID string) (string, error) {
	body, err := s.deleteAffinityRuleMemberResponseBody(memberID)

	return body.Data.TaskID, err
}

func (s *Service) deleteAffinityRuleMemberResponseBody(memberID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if memberID == "" {
		return body, fmt.Errorf("invalid affinity rule member id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/affinity-rule-members/%s", memberID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AffinityRuleMemberNotFoundError{ID: memberID}
		}

		return nil
	})
}
