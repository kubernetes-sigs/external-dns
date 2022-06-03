package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerACLs retrieves a list of ACLs
func (s *Service) GetListenerACLs(listenerID int, parameters connection.APIRequestParameters) ([]ACL, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}, parameters)
}

// GetListenerACLsPaginated retrieves a paginated list of ACLs
func (s *Service) GetListenerACLsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	body, err := s.getListenerACLsPaginatedResponseBody(listenerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
		return s.GetListenerACLsPaginated(listenerID, p)
	}), err
}

func (s *Service) getListenerACLsPaginatedResponseBody(listenerID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ACL], error) {
	body := &connection.APIResponseBodyData[[]ACL]{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/acls", listenerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
