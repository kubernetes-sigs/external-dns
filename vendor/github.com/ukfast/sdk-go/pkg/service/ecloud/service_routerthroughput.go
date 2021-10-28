package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRouterThroughputs retrieves a list of router throughputs
func (s *Service) GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error) {
	var throughputs []RouterThroughput

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterThroughputsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, throughput := range response.(*PaginatedRouterThroughput).Items {
			throughputs = append(throughputs, throughput)
		}
	}

	return throughputs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRouterThroughputsPaginated retrieves a paginated list of router throughputs
func (s *Service) GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*PaginatedRouterThroughput, error) {
	body, err := s.getRouterThroughputsPaginatedResponseBody(parameters)

	return NewPaginatedRouterThroughput(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterThroughputsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRouterThroughputsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRouterThroughputSliceResponseBody, error) {
	body := &GetRouterThroughputSliceResponseBody{}

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

func (s *Service) getRouterThroughputResponseBody(throughputID string) (*GetRouterThroughputResponseBody, error) {
	body := &GetRouterThroughputResponseBody{}

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
