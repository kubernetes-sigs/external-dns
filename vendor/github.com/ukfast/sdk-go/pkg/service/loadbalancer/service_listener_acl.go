package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetListenerACLs retrieves a list of ACLs
func (s *Service) GetListenerACLs(listenerID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	var acls []ACL

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, acl := range response.(*PaginatedACL).Items {
			acls = append(acls, acl)
		}
	}

	return acls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetListenerACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetListenerACLsPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedACL, error) {
	body, err := s.getListenerACLsPaginatedResponseBody(listenerID, parameters)

	return NewPaginatedACL(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getListenerACLsPaginatedResponseBody(listenerID int, parameters connection.APIRequestParameters) (*GetACLSliceResponseBody, error) {
	body := &GetACLSliceResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/acls", listenerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
