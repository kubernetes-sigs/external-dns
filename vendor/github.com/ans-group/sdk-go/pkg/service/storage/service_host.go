package storage

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	return connection.InvokeRequestAll(s.GetHostsPaginated, parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error) {
	body, err := s.getHostsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetHostsPaginated), err
}

func (s *Service) getHostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Host], error) {
	body := &connection.APIResponseBodyData[[]Host]{}

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

func (s *Service) getHostResponseBody(hostID int) (*connection.APIResponseBodyData[Host], error) {
	body := &connection.APIResponseBodyData[Host]{}

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
