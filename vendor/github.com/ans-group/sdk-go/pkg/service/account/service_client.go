package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetClients retrieves a list of clients
func (s *Service) GetClients(parameters connection.APIRequestParameters) ([]Client, error) {
	return connection.InvokeRequestAll(s.GetClientsPaginated, parameters)
}

// GetClientsPaginated retrieves a paginated list of clients
func (s *Service) GetClientsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Client], error) {
	body, err := s.getClientsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetClientsPaginated), err
}

func (s *Service) getClientsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Client], error) {
	body := &connection.APIResponseBodyData[[]Client]{}

	response, err := s.connection.Get("/account/v1/clients", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetClient retrieves a single client by id
func (s *Service) GetClient(clientID int) (Client, error) {
	body, err := s.getClientResponseBody(clientID)

	return body.Data, err
}

func (s *Service) getClientResponseBody(clientID int) (*connection.APIResponseBodyData[Client], error) {
	body := &connection.APIResponseBodyData[Client]{}

	if clientID < 1 {
		return body, fmt.Errorf("invalid client id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/clients/%d", clientID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClientNotFoundError{ID: clientID}
		}

		return nil
	})
}

// CreateClient creates a new client
func (s *Service) CreateClient(req CreateClientRequest) (int, error) {
	body, err := s.createClientResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createClientResponseBody(req CreateClientRequest) (*connection.APIResponseBodyData[Client], error) {
	body := &connection.APIResponseBodyData[Client]{}

	response, err := s.connection.Post("/account/v1/clients", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchClient patches a client
func (s *Service) PatchClient(clientID int, patch PatchClientRequest) error {
	_, err := s.patchClientResponseBody(clientID, patch)

	return err
}

func (s *Service) patchClientResponseBody(clientID int, patch PatchClientRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clientID < 1 {
		return body, fmt.Errorf("invalid client id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/account/v1/clients/%d", clientID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClientNotFoundError{ID: clientID}
		}

		return nil
	})
}

// DeleteClient removes a client
func (s *Service) DeleteClient(clientID int) error {
	_, err := s.deleteClientResponseBody(clientID)

	return err
}

func (s *Service) deleteClientResponseBody(clientID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if clientID < 1 {
		return body, fmt.Errorf("invalid client id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/account/v1/clients/%d", clientID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ClientNotFoundError{ID: clientID}
		}

		return nil
	})
}
