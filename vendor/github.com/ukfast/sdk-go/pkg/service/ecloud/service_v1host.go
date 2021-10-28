package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetV1Hosts retrieves a list of v1 hosts
func (s *Service) GetV1Hosts(parameters connection.APIRequestParameters) ([]V1Host, error) {
	var hosts []V1Host

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetV1HostsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, host := range response.(*PaginatedV1Host).Items {
			hosts = append(hosts, host)
		}
	}

	return hosts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetV1HostsPaginated retrieves a paginated list of v1 hosts
func (s *Service) GetV1HostsPaginated(parameters connection.APIRequestParameters) (*PaginatedV1Host, error) {
	body, err := s.getV1HostsPaginatedResponseBody(parameters)

	return NewPaginatedV1Host(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetV1HostsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getV1HostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetV1HostSliceResponseBody, error) {
	body := &GetV1HostSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v1/hosts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetV1Host retrieves a single v1 host by ID
func (s *Service) GetV1Host(hostID int) (V1Host, error) {
	body, err := s.getV1HostResponseBody(hostID)

	return body.Data, err
}

func (s *Service) getV1HostResponseBody(hostID int) (*GetV1HostResponseBody, error) {
	body := &GetV1HostResponseBody{}

	if hostID < 1 {
		return body, fmt.Errorf("invalid host id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/hosts/%d", hostID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &V1HostNotFoundError{ID: hostID}
		}

		return nil
	})
}
