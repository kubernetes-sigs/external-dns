package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetListenerBinds retrieves a list of binds
func (s *Service) GetListenerBinds(listenerID int, parameters connection.APIRequestParameters) ([]Bind, error) {
	var binds []Bind

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerBindsPaginated(listenerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, bind := range response.(*PaginatedBind).Items {
			binds = append(binds, bind)
		}
	}

	return binds, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetListenerBindsPaginated retrieves a paginated list of binds
func (s *Service) GetListenerBindsPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedBind, error) {
	body, err := s.getListenerBindsPaginatedResponseBody(listenerID, parameters)

	return NewPaginatedBind(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerBindsPaginated(listenerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getListenerBindsPaginatedResponseBody(listenerID int, parameters connection.APIRequestParameters) (*GetBindSliceResponseBody, error) {
	body := &GetBindSliceResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds", listenerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetListenerBind retrieves a single bind by id
func (s *Service) GetListenerBind(listenerID int, bindID int) (Bind, error) {
	body, err := s.getListenerBindResponseBody(listenerID, bindID)

	return body.Data, err
}

func (s *Service) getListenerBindResponseBody(listenerID int, bindID int) (*GetBindResponseBody, error) {
	body := &GetBindResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if bindID < 1 {
		return body, fmt.Errorf("invalid bind id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BindNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// CreateListenerBind creates an bind
func (s *Service) CreateListenerBind(listenerID int, req CreateBindRequest) (int, error) {
	body, err := s.createListenerBindResponseBody(listenerID, req)

	return body.Data.ID, err
}

func (s *Service) createListenerBindResponseBody(listenerID int, req CreateBindRequest) (*GetBindResponseBody, error) {
	body := &GetBindResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds", listenerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BindNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// PatchListenerBind patches an bind
func (s *Service) PatchListenerBind(listenerID int, bindID int, req PatchBindRequest) error {
	_, err := s.patchListenerBindResponseBody(listenerID, bindID, req)

	return err
}

func (s *Service) patchListenerBindResponseBody(listenerID int, bindID int, req PatchBindRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if bindID < 1 {
		return body, fmt.Errorf("invalid bind id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BindNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// DeleteListenerBind deletes a bind
func (s *Service) DeleteListenerBind(listenerID int, bindID int) error {
	_, err := s.deleteListenerBindResponseBody(listenerID, bindID)

	return err
}

func (s *Service) deleteListenerBindResponseBody(listenerID int, bindID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if bindID < 1 {
		return body, fmt.Errorf("invalid bind id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/listeners/%d/binds/%d", listenerID, bindID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BindNotFoundError{ID: listenerID}
		}

		return nil
	})
}
