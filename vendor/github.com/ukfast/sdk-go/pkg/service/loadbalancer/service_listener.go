package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetListeners retrieves a list of listeners
func (s *Service) GetListeners(parameters connection.APIRequestParameters) ([]Listener, error) {
	var listeners []Listener

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, listener := range response.(*PaginatedListener).Items {
			listeners = append(listeners, listener)
		}
	}

	return listeners, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetListenersPaginated retrieves a paginated list of listeners
func (s *Service) GetListenersPaginated(parameters connection.APIRequestParameters) (*PaginatedListener, error) {
	body, err := s.getListenersPaginatedResponseBody(parameters)

	return NewPaginatedListener(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getListenersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetListenerSliceResponseBody, error) {
	body := &GetListenerSliceResponseBody{}

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

func (s *Service) getListenerResponseBody(listenerID int) (*GetListenerResponseBody, error) {
	body := &GetListenerResponseBody{}

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

func (s *Service) createListenerResponseBody(req CreateListenerRequest) (*GetListenerResponseBody, error) {
	body := &GetListenerResponseBody{}

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
