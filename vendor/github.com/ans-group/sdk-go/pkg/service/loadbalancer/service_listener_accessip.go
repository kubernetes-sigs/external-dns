package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerAccessIPs retrieves a list of access IPs
func (s *Service) GetListenerAccessIPs(listenerID int, parameters connection.APIRequestParameters) ([]AccessIP, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
		return s.GetListenerAccessIPsPaginated(listenerID, p)
	}, parameters)
}

// GetListenerAccessIPsPaginated retrieves a paginated list of access IPs
func (s *Service) GetListenerAccessIPsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
	body, err := s.getListenerAccessIPsPaginatedResponseBody(listenerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
		return s.GetListenerAccessIPsPaginated(listenerID, p)
	}), err
}

func (s *Service) getListenerAccessIPsPaginatedResponseBody(listenerID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AccessIP], error) {
	body := &connection.APIResponseBodyData[[]AccessIP]{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips", listenerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetListenerAccessIP retrieves a single access IP by id
func (s *Service) GetListenerAccessIP(listenerID int, accessID int) (AccessIP, error) {
	body, err := s.getListenerAccessIPResponseBody(listenerID, accessID)

	return body.Data, err
}

func (s *Service) getListenerAccessIPResponseBody(listenerID int, accessID int) (*connection.APIResponseBodyData[AccessIP], error) {
	body := &connection.APIResponseBodyData[AccessIP]{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if accessID < 1 {
		return body, fmt.Errorf("invalid access id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips/%d", listenerID, accessID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccessIPNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// CreateListenerAccessIP creates an access IP
func (s *Service) CreateListenerAccessIP(listenerID int, req CreateAccessIPRequest) (int, error) {
	body, err := s.createListenerAccessIPResponseBody(listenerID, req)

	return body.Data.ID, err
}

func (s *Service) createListenerAccessIPResponseBody(listenerID int, req CreateAccessIPRequest) (*connection.APIResponseBodyData[AccessIP], error) {
	body := &connection.APIResponseBodyData[AccessIP]{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips", listenerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccessIPNotFoundError{ID: listenerID}
		}

		return nil
	})
}
