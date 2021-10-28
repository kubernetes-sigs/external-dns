package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetACLs retrieves a list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLs(parameters connection.APIRequestParameters) ([]ACL, error) {
	var acls []ACL

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetACLsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, acl := range response.(*PaginatedACL).Items {
			acls = append(acls, acl)
		}
	}

	return acls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetACLsPaginated retrieves a paginated list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLsPaginated(parameters connection.APIRequestParameters) (*PaginatedACL, error) {
	body, err := s.getACLsPaginatedResponseBody(parameters)

	return NewPaginatedACL(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetACLsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getACLsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetACLSliceResponseBody, error) {
	body := &GetACLSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/acls", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetACL retrieves a single ACL by id
func (s *Service) GetACL(aclID int) (ACL, error) {
	body, err := s.getACLResponseBody(aclID)

	return body.Data, err
}

func (s *Service) getACLResponseBody(aclID int) (*GetACLResponseBody, error) {
	body := &GetACLResponseBody{}

	if aclID < 1 {
		return body, fmt.Errorf("invalid acl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLNotFoundError{ID: aclID}
		}

		return nil
	})
}

// CreateACL creates an ACL
func (s *Service) CreateACL(req CreateACLRequest) (int, error) {
	body, err := s.createACLResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createACLResponseBody(req CreateACLRequest) (*GetACLResponseBody, error) {
	body := &GetACLResponseBody{}

	response, err := s.connection.Post("/loadbalancers/v2/acls", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}

// PatchACL patches an ACL
func (s *Service) PatchACL(aclID int, req PatchACLRequest) error {
	_, err := s.patchACLResponseBody(aclID, req)

	return err
}

func (s *Service) patchACLResponseBody(aclID int, req PatchACLRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if aclID < 1 {
		return body, fmt.Errorf("invalid acl id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLNotFoundError{ID: aclID}
		}

		return nil
	})
}

// DeleteACL deletes an ACL
func (s *Service) DeleteACL(aclID int) error {
	_, err := s.deleteACLResponseBody(aclID)

	return err
}

func (s *Service) deleteACLResponseBody(aclID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if aclID < 1 {
		return body, fmt.Errorf("invalid acl id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLNotFoundError{ID: aclID}
		}

		return nil
	})
}
