package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLoadBalancers retrieves a list of load balancers
func (s *Service) GetLoadBalancers(parameters connection.APIRequestParameters) ([]LoadBalancer, error) {
	var lbs []LoadBalancer

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		lbs = append(lbs, response.(*PaginatedLoadBalancer).Items...)
	}

	return lbs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetLoadBalancersPaginated retrieves a paginated list of lbs
func (s *Service) GetLoadBalancersPaginated(parameters connection.APIRequestParameters) (*PaginatedLoadBalancer, error) {
	body, err := s.getLoadBalancersPaginatedResponseBody(parameters)

	return NewPaginatedLoadBalancer(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetLoadBalancersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getLoadBalancersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetLoadBalancerSliceResponseBody, error) {
	body := &GetLoadBalancerSliceResponseBody{}

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

func (s *Service) getLoadBalancerResponseBody(loadbalancerID string) (*GetLoadBalancerResponseBody, error) {
	body := &GetLoadBalancerResponseBody{}

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

func (s *Service) createLoadBalancerResponseBody(req CreateLoadBalancerRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) patchLoadBalancerResponseBody(loadbalancerID string, req PatchLoadBalancerRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) deleteLoadBalancerResponseBody(loadbalancerID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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
