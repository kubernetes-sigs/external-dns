package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetLoadBalancerSpecs retrieves a list of load balancer specs
func (s *Service) GetLoadBalancerSpecs(parameters connection.APIRequestParameters) ([]LoadBalancerSpec, error) {
	return connection.InvokeRequestAll(s.GetLoadBalancerSpecsPaginated, parameters)
}

// GetLoadBalancerSpecsPaginated retrieves a paginated list of load balancer specs
func (s *Service) GetLoadBalancerSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[LoadBalancerSpec], error) {
	body, err := s.getLoadBalancerSpecsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetLoadBalancerSpecsPaginated), err
}

func (s *Service) getLoadBalancerSpecsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]LoadBalancerSpec], error) {
	body := &connection.APIResponseBodyData[[]LoadBalancerSpec]{}

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

func (s *Service) getLoadBalancerSpecResponseBody(lbSpecID string) (*connection.APIResponseBodyData[LoadBalancerSpec], error) {
	body := &connection.APIResponseBodyData[LoadBalancerSpec]{}

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
