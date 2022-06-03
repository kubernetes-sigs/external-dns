package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListeners retrieves a list of listeners
func (s *Service) GetListeners(parameters connection.APIRequestParameters) ([]Listener, error) {
	return connection.InvokeRequestAll(s.GetListenersPaginated, parameters)
}

// GetListenersPaginated retrieves a paginated list of listeners
func (s *Service) GetListenersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Listener], error) {
	body, err := s.getListenersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetListenersPaginated), err
}

func (s *Service) getListenersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Listener], error) {
	body := &connection.APIResponseBodyData[[]Listener]{}

	response, err := s.connection.Get("/loadbalancers/v2/listeners", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetListener retrieves a single listener by id
func (s *Service) GetListener(listenerID int) (Listener, error) {
	body, err := s.getListenerResponseBody(listenerID)

	return body.Data, err
}

func (s *Service) getListenerResponseBody(listenerID int) (*connection.APIResponseBodyData[Listener], error) {
	body := &connection.APIResponseBodyData[Listener]{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ListenerNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// CreateListener creates a listener
func (s *Service) CreateListener(req CreateListenerRequest) (int, error) {
	body, err := s.createListenerResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createListenerResponseBody(req CreateListenerRequest) (*connection.APIResponseBodyData[Listener], error) {
	body := &connection.APIResponseBodyData[Listener]{}

	response, err := s.connection.Post("/loadbalancers/v2/listeners", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body)
}

// PatchListener patches a listener
func (s *Service) PatchListener(listenerID int, req PatchListenerRequest) error {
	_, err := s.patchListenerResponseBody(listenerID, req)

	return err
}

func (s *Service) patchListenerResponseBody(listenerID int, req PatchListenerRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ListenerNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// DeleteListener deletes a listener
func (s *Service) DeleteListener(listenerID int) error {
	_, err := s.deleteListenerResponseBody(listenerID)

	return err
}

func (s *Service) deleteListenerResponseBody(listenerID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/listeners/%d", listenerID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ListenerNotFoundError{ID: listenerID}
		}

		return nil
	})
}
