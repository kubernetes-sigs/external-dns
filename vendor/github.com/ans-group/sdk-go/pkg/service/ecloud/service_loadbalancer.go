package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetLoadBalancers retrieves a list of load balancers
func (s *Service) GetLoadBalancers(parameters connection.APIRequestParameters) ([]LoadBalancer, error) {
	return connection.InvokeRequestAll(s.GetLoadBalancersPaginated, parameters)
}

// GetLoadBalancersPaginated retrieves a paginated list of lbs
func (s *Service) GetLoadBalancersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[LoadBalancer], error) {
	body, err := s.getLoadBalancersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetLoadBalancersPaginated), err
}

func (s *Service) getLoadBalancersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]LoadBalancer], error) {
	body := &connection.APIResponseBodyData[[]LoadBalancer]{}

	response, err := s.connection.Get("/ecloud/v2/load-balancers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetLoadBalancer retrieves a single lb by id
func (s *Service) GetLoadBalancer(loadbalancerID string) (LoadBalancer, error) {
	body, err := s.getLoadBalancerResponseBody(loadbalancerID)

	return body.Data, err
}

func (s *Service) getLoadBalancerResponseBody(loadbalancerID string) (*connection.APIResponseBodyData[LoadBalancer], error) {
	body := &connection.APIResponseBodyData[LoadBalancer]{}

	if loadbalancerID == "" {
		return body, fmt.Errorf("invalid load balancer id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNotFoundError{ID: loadbalancerID}
		}

		return nil
	})
}

// CreateLoadBalancer creates a new LoadBalancer
func (s *Service) CreateLoadBalancer(req CreateLoadBalancerRequest) (TaskReference, error) {
	body, err := s.createLoadBalancerResponseBody(req)

	return body.Data, err
}

func (s *Service) createLoadBalancerResponseBody(req CreateLoadBalancerRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/load-balancers", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchLoadBalancer patches a LoadBalancer
func (s *Service) PatchLoadBalancer(loadbalancerID string, req PatchLoadBalancerRequest) (TaskReference, error) {
	body, err := s.patchLoadBalancerResponseBody(loadbalancerID, req)

	return body.Data, err
}

func (s *Service) patchLoadBalancerResponseBody(loadbalancerID string, req PatchLoadBalancerRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if loadbalancerID == "" {
		return body, fmt.Errorf("invalid load balancer id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNotFoundError{ID: loadbalancerID}
		}

		return nil
	})
}

// DeleteLoadBalancer deletes a LoadBalancer
func (s *Service) DeleteLoadBalancer(loadbalancerID string) (string, error) {
	body, err := s.deleteLoadBalancerResponseBody(loadbalancerID)

	return body.Data.TaskID, err
}

func (s *Service) deleteLoadBalancerResponseBody(loadbalancerID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if loadbalancerID == "" {
		return body, fmt.Errorf("invalid load balancer id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/load-balancers/%s", loadbalancerID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &LoadBalancerNotFoundError{ID: loadbalancerID}
		}

		return nil
	})
}
