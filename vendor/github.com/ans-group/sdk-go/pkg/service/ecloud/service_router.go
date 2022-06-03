package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRouters retrieves a list of routers
func (s *Service) GetRouters(parameters connection.APIRequestParameters) ([]Router, error) {
	return connection.InvokeRequestAll(s.GetRoutersPaginated, parameters)
}

// GetRoutersPaginated retrieves a paginated list of routers
func (s *Service) GetRoutersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Router], error) {
	body, err := s.getRoutersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetRoutersPaginated), err
}

func (s *Service) getRoutersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Router], error) {
	body := &connection.APIResponseBodyData[[]Router]{}

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

func (s *Service) getRouterResponseBody(routerID string) (*connection.APIResponseBodyData[Router], error) {
	body := &connection.APIResponseBodyData[Router]{}

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

func (s *Service) createRouterResponseBody(req CreateRouterRequest) (*connection.APIResponseBodyData[Router], error) {
	body := &connection.APIResponseBodyData[Router]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}, parameters)
}

// GetRouterFirewallPoliciesPaginated retrieves a paginated list of firewall rule policies
func (s *Service) GetRouterFirewallPoliciesPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
	body, err := s.getRouterFirewallPoliciesPaginatedResponseBody(routerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FirewallPolicy], error) {
		return s.GetRouterFirewallPoliciesPaginated(routerID, p)
	}), err
}

func (s *Service) getRouterFirewallPoliciesPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FirewallPolicy], error) {
	body := &connection.APIResponseBodyData[[]FirewallPolicy]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Network], error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}, parameters)
}

// GetRouterNetworksPaginated retrieves a paginated list of router networks
func (s *Service) GetRouterNetworksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Network], error) {
	body, err := s.getRouterNetworksPaginatedResponseBody(routerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Network], error) {
		return s.GetRouterNetworksPaginated(routerID, p)
	}), err
}

func (s *Service) getRouterNetworksPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Network], error) {
	body := &connection.APIResponseBodyData[[]Network]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}, parameters)
}

// GetRouterVPNsPaginated retrieves a paginated list of router VPNs
func (s *Service) GetRouterVPNsPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
	body, err := s.getRouterVPNsPaginatedResponseBody(routerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[VPN], error) {
		return s.GetRouterVPNsPaginated(routerID, p)
	}), err
}

func (s *Service) getRouterVPNsPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPN], error) {
	body := &connection.APIResponseBodyData[[]VPN]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}, parameters)
}

// GetRouterTasksPaginated retrieves a paginated list of Router tasks
func (s *Service) GetRouterTasksPaginated(routerID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getRouterTasksPaginatedResponseBody(routerID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetRouterTasksPaginated(routerID, p)
	}), err
}

func (s *Service) getRouterTasksPaginatedResponseBody(routerID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

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
