package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTargetGroupACLs retrieves a list of ACLs
func (s *Service) GetTargetGroupACLs(targetGroupID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	var acls []ACL

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupACLsPaginated(targetGroupID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, acl := range response.(*PaginatedACL).Items {
			acls = append(acls, acl)
		}
	}

	return acls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTargetGroupACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetTargetGroupACLsPaginated(targetGroupID int, parameters connection.APIRequestParameters) (*PaginatedACL, error) {
	body, err := s.getTargetGroupACLsPaginatedResponseBody(targetGroupID, parameters)

	return NewPaginatedACL(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetGroupACLsPaginated(targetGroupID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTargetGroupACLsPaginatedResponseBody(targetGroupID int, parameters connection.APIRequestParameters) (*GetACLSliceResponseBody, error) {
	body := &GetACLSliceResponseBody{}

	if targetGroupID < 1 {
		return body, fmt.Errorf("invalid target group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/target-groups/%d/acls", targetGroupID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
