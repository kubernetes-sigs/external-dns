package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLoadBalancerSpecs retrieves a list of load balancer specs
func (s *Service) GetLoadBalancerSpecs(parameters connection.APIRequestParameters) ([]LoadBalancerSpec, error) {
	var specs []LoadBalancerSpec

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerSpecsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, spec := range response.(*PaginatedLoadBalancerSpec).Items {
			specs = append(specs, spec)
		}
	}

	return specs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetLoadBalancerSpecsPaginated retrieves a paginated list of load balancer specs
func (s *Service) GetLoadBalancerSpecsPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancerSpec, error) {
	body, err := s.getLoadBalancerSpecsPaginatedResponseBody(parameters)

	return NewPaginatedLoadBalancerSpec(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancerSpecsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getLoadBalancerSpecsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetLoadBalancerSpecSliceResponseBody, error) {
	body := &GetLoadBalancerSpecSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/load-balancer-specs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetLoadBalancerSpec retrieves a single spec by id
func (s *Service) GetLoadBalancerSpec(lbSpecID string) (LoadBalancerSpec, error) {
	body, err := s.getLoadBalancerSpecResponseBody(lbSpecID)

	return body.Data, err
}

func (s *Service) getLoadBalancerSpecResponseBody(lbSpecID string) (*GetLoadBalancerSpecResponseBody, error) {
	body := &GetLoadBalancerSpecResponseBody{}

	if lbSpecID == "" {
		return body, fmt.Errorf("invalid load balancer spec id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/load-balancer-specs/%s", lbSpecID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerSpecNotFoundError{ID: lbSpecID}
		}

		return nil
	})
}
