package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRouterThroughputs retrieves a list of router throughputs
func (s *Service) GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error) {
	return connection.InvokeRequestAll(s.GetRouterThroughputsPaginated, parameters)
}

// GetRouterThroughputsPaginated retrieves a paginated list of router throughputs
func (s *Service) GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RouterThroughput], error) {
	body, err := s.getRouterThroughputsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetRouterThroughputsPaginated), err
}

func (s *Service) getRouterThroughputsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]RouterThroughput], error) {
	body := &connection.APIResponseBodyData[[]RouterThroughput]{}

	response, err := s.connection.Get("/ecloud/v2/router-throughputs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRouterThroughput retrieves a single router throughput by id
func (s *Service) GetRouterThroughput(throughputID string) (RouterThroughput, error) {
	body, err := s.getRouterThroughputResponseBody(throughputID)

	return body.Data, err
}

func (s *Service) getRouterThroughputResponseBody(throughputID string) (*connection.APIResponseBodyData[RouterThroughput], error) {
	body := &connection.APIResponseBodyData[RouterThroughput]{}

	if throughputID == "" {
		return body, fmt.Errorf("invalid router throughput id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/router-throughputs/%s", throughputID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterThroughputNotFoundError{ID: throughputID}
		}

		return nil
	})
}
