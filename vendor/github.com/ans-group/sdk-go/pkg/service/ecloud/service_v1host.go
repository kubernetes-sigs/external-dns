package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetV1Hosts retrieves a list of v1 hosts
func (s *Service) GetV1Hosts(parameters connection.APIRequestParameters) ([]V1Host, error) {
	return connection.InvokeRequestAll(s.GetV1HostsPaginated, parameters)
}

// GetV1HostsPaginated retrieves a paginated list of v1 hosts
func (s *Service) GetV1HostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
	body, err := s.getV1HostsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetV1HostsPaginated), err
}

func (s *Service) getV1HostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]V1Host], error) {
	body := &connection.APIResponseBodyData[[]V1Host]{}

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

func (s *Service) getV1HostResponseBody(hostID int) (*connection.APIResponseBodyData[V1Host], error) {
	body := &connection.APIResponseBodyData[V1Host]{}

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
