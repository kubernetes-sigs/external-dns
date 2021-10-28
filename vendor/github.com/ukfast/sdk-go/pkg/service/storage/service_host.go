package storage

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	var hosts []Host

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, host := range response.(*PaginatedHost).Items {
			hosts = append(hosts, host)
		}
	}

	return hosts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*PaginatedHost, error) {
	body, err := s.getHostsPaginatedResponseBody(parameters)

	return NewPaginatedHost(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetHostSliceResponseBody, error) {
	body := &GetHostSliceResponseBody{}

	response, err := s.connection.Get("/ukfast-storage/v1/hosts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID int) (Host, error) {
	body, err := s.getHostResponseBody(hostID)

	return body.Data, err
}

func (s *Service) getHostResponseBody(hostID int) (*GetHostResponseBody, error) {
	body := &GetHostResponseBody{}

	if hostID < 1 {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ukfast-storage/v1/hosts/%d", hostID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostNotFoundError{ID: hostID}
		}

		return nil
	})
}
