package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetNetworks retrieves a list of networks
func (s *Service) GetNetworks(parameters connection.APIRequestParameters) ([]Network, error) {
	return connection.InvokeRequestAll(s.GetNetworksPaginated, parameters)
}

// GetNetworksPaginated retrieves a paginated list of networks
func (s *Service) GetNetworksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	body, err := s.getNetworksPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetNetworksPaginated), err
}

func (s *Service) getNetworksPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Network], error) {
	body := &connection.APIResponseBodyData[[]Network]{}

	response, err := s.connection.Get("/ecloud/v2/networks", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetwork retrieves a single network by id
func (s *Service) GetNetwork(networkID string) (Network, error) {
	body, err := s.getNetworkResponseBody(networkID)

	return body.Data, err
}

func (s *Service) getNetworkResponseBody(networkID string) (*connection.APIResponseBodyData[Network], error) {
	body := &connection.APIResponseBodyData[Network]{}

	if networkID == "" {
		return body, fmt.Errorf("invalid network id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/networks/%s", networkID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkNotFoundError{ID: networkID}
		}

		return nil
	})
}

// CreateNetwork creates a new Network
func (s *Service) CreateNetwork(req CreateNetworkRequest) (string, error) {
	body, err := s.createNetworkResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createNetworkResponseBody(req CreateNetworkRequest) (*connection.APIResponseBodyData[Network], error) {
	body := &connection.APIResponseBodyData[Network]{}

	response, err := s.connection.Post("/ecloud/v2/networks", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchNetwork patches a Network
func (s *Service) PatchNetwork(networkID string, req PatchNetworkRequest) error {
	_, err := s.patchNetworkResponseBody(networkID, req)

	return err
}

func (s *Service) patchNetworkResponseBody(networkID string, req PatchNetworkRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if networkID == "" {
		return body, fmt.Errorf("invalid network id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/networks/%s", networkID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkNotFoundError{ID: networkID}
		}

		return nil
	})
}

// DeleteNetwork deletes a Network
func (s *Service) DeleteNetwork(networkID string) error {
	_, err := s.deleteNetworkResponseBody(networkID)

	return err
}

func (s *Service) deleteNetworkResponseBody(networkID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if networkID == "" {
		return body, fmt.Errorf("invalid network id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/networks/%s", networkID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkNotFoundError{ID: networkID}
		}

		return nil
	})
}

// GetNetworkNICs retrieves a list of firewall rule nics
func (s *Service) GetNetworkNICs(networkID string, parameters connection.APIRequestParameters) ([]NIC, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetNetworkNICsPaginated(networkID, p)
	}, parameters)
}

// GetNetworkNICsPaginated retrieves a paginated list of firewall rule nics
func (s *Service) GetNetworkNICsPaginated(networkID string, parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	body, err := s.getNetworkNICsPaginatedResponseBody(networkID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetNetworkNICsPaginated(networkID, p)
	}), err
}

func (s *Service) getNetworkNICsPaginatedResponseBody(networkID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NIC], error) {
	body := &connection.APIResponseBodyData[[]NIC]{}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/networks/%s/nics", networkID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetNetworkTasks retrieves a list of Network tasks
func (s *Service) GetNetworkTasks(networkID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkTasksPaginated(networkID, p)
	}, parameters)
}

// GetNetworkTasksPaginated retrieves a paginated list of Network tasks
func (s *Service) GetNetworkTasksPaginated(networkID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getNetworkTasksPaginatedResponseBody(networkID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetNetworkTasksPaginated(networkID, p)
	}), err
}

func (s *Service) getNetworkTasksPaginatedResponseBody(networkID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if networkID == "" {
		return body, fmt.Errorf("invalid network id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/networks/%s/tasks", networkID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &NetworkNotFoundError{ID: networkID}
		}

		return nil
	})
}
