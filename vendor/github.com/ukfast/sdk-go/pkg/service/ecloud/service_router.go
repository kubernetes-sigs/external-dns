package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	var routers []Router

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRoutersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, router := range response.(*PaginatedRouter).Items {
			routers = append(routers, router)
		}
	}

	return routers, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRoutersPaginated retrieves a paginated list of routers
func (s *Service) GetRoutersPaginated(parameters connection.APIRequestParameters) (*PaginatedRouter, error) {
	body, err := s.getRoutersPaginatedResponseBody(parameters)

	return NewPaginatedRouter(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRoutersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRoutersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRouterSliceResponseBody, error) {
	body := &GetRouterSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/routers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRouter retrieves a single router by id
func (s *Service) GetRouter(routerID string) (Router, error) {
	body, err := s.getRouterResponseBody(routerID)

	return body.Data, err
}

func (s *Service) getRouterResponseBody(routerID string) (*GetRouterResponseBody, error) {
	body := &GetRouterResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// CreateRouter creates a new Router
func (s *Service) CreateRouter(req CreateRouterRequest) (string, error) {
	body, err := s.createRouterResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createRouterResponseBody(req CreateRouterRequest) (*GetRouterResponseBody, error) {
	body := &GetRouterResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/routers", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchRouter patches a Router
func (s *Service) PatchRouter(routerID string, req PatchRouterRequest) error {
	_, err := s.patchRouterResponseBody(routerID, req)

	return err
}

func (s *Service) patchRouterResponseBody(routerID string, req PatchRouterRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// DeleteRouter deletes a Router
func (s *Service) DeleteRouter(routerID string) error {
	_, err := s.deleteRouterResponseBody(routerID)

	return err
}

func (s *Service) deleteRouterResponseBody(routerID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/routers/%s", routerID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// GetRouterFirewallPolicies retrieves a list of firewall rule policies
func (s *Service) GetRouterFirewallPolicies(routerID string, parameters connection.APIRequestParameters) ([]FirewallPolicy, error) {
	var policies []FirewallPolicy

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedFirewallPolicy).Items {
			policies = append(policies, policy)
		}
	}

	return policies, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRouterFirewallPoliciesPaginated retrieves a paginated list of firewall rule policies
func (s *Service) GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedFirewallPolicy, error) {
	body, err := s.getRouterFirewallPoliciesPaginatedResponseBody(routerID, parameters)

	return NewPaginatedFirewallPolicy(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRouterFirewallPoliciesPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*GetFirewallPolicySliceResponseBody, error) {
	body := &GetFirewallPolicySliceResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s/firewall-policies", routerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// GetRouterNetworks retrieves a list of router networks
func (s *Service) GetRouterNetworks(routerID string, parameters connection.APIRequestParameters) ([]Network, error) {
	var policies []Network

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedNetwork).Items {
			policies = append(policies, policy)
		}
	}

	return policies, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRouterNetworksPaginated retrieves a paginated list of router networks
func (s *Service) GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedNetwork, error) {
	body, err := s.getRouterNetworksPaginatedResponseBody(routerID, parameters)

	return NewPaginatedNetwork(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRouterNetworksPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*GetNetworkSliceResponseBody, error) {
	body := &GetNetworkSliceResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s/networks", routerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// GetRouterVPNs retrieves a list of router VPNs
func (s *Service) GetRouterVPNs(routerID string, parameters connection.APIRequestParameters) ([]VPN, error) {
	var policies []VPN

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, policy := range response.(*PaginatedVPN).Items {
			policies = append(policies, policy)
		}
	}

	return policies, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRouterVPNsPaginated retrieves a paginated list of router VPNs
func (s *Service) GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedVPN, error) {
	body, err := s.getRouterVPNsPaginatedResponseBody(routerID, parameters)

	return NewPaginatedVPN(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRouterVPNsPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*GetVPNSliceResponseBody, error) {
	body := &GetVPNSliceResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s/vpns", routerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// DeployRouterDefaultFirewallPolicies deploys default firewall policy resources for specified router
func (s *Service) DeployRouterDefaultFirewallPolicies(routerID string) error {
	_, err := s.deployRouterDefaultFirewallPolicies(routerID)

	return err
}

func (s *Service) deployRouterDefaultFirewallPolicies(routerID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/routers/%s/configure-default-policies", routerID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}

// GetRouterTasks retrieves a list of Router tasks
func (s *Service) GetRouterTasks(routerID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRouterTasksPaginated retrieves a paginated list of Router tasks
func (s *Service) GetRouterTasksPaginated(routerID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getRouterTasksPaginatedResponseBody(routerID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRouterTasksPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if routerID == "" {
		return body, fmt.Errorf("invalid router id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/routers/%s/tasks", routerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RouterNotFoundError{ID: routerID}
		}

		return nil
	})
}
