package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPNServices retrieves a list of VPN services
func (s *Service) GetVPNServices(parameters connection.APIRequestParameters) ([]VPNService, error) {
	var services []VPNService

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNServicesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, service := range response.(*PaginatedVPNService).Items {
			services = append(services, service)
		}
	}

	return services, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNServicesPaginated retrieves a paginated list of VPN services
func (s *Service) GetVPNServicesPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNService, error) {
	body, err := s.getVPNServicesPaginatedResponseBody(parameters)

	return NewPaginatedVPNService(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNServicesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNServicesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPNServiceSliceResponseBody, error) {
	body := &GetVPNServiceSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/vpn-services", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNService retrieves a single VPN service by id
func (s *Service) GetVPNService(serviceID string) (VPNService, error) {
	body, err := s.getVPNServiceResponseBody(serviceID)

	return body.Data, err
}

func (s *Service) getVPNServiceResponseBody(serviceID string) (*GetVPNServiceResponseBody, error) {
	body := &GetVPNServiceResponseBody{}

	if serviceID == "" {
		return body, fmt.Errorf("invalid vpn service id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNServiceNotFoundError{ID: serviceID}
		}

		return nil
	})
}

// CreateVPNService creates a new VPN service
func (s *Service) CreateVPNService(req CreateVPNServiceRequest) (TaskReference, error) {
	body, err := s.createVPNServiceResponseBody(req)

	return body.Data, err
}

func (s *Service) createVPNServiceResponseBody(req CreateVPNServiceRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/vpn-services", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPNService patches a VPN service
func (s *Service) PatchVPNService(serviceID string, req PatchVPNServiceRequest) (TaskReference, error) {
	body, err := s.patchVPNServiceResponseBody(serviceID, req)

	return body.Data, err
}

func (s *Service) patchVPNServiceResponseBody(serviceID string, req PatchVPNServiceRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if serviceID == "" {
		return body, fmt.Errorf("invalid service id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNServiceNotFoundError{ID: serviceID}
		}

		return nil
	})
}

// DeleteVPNService deletes a VPN service
func (s *Service) DeleteVPNService(serviceID string) (string, error) {
	body, err := s.deleteVPNServiceResponseBody(serviceID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVPNServiceResponseBody(serviceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if serviceID == "" {
		return body, fmt.Errorf("invalid service id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNServiceNotFoundError{ID: serviceID}
		}

		return nil
	})
}

// GetVPNServiceTasks retrieves a list of VPN service tasks
func (s *Service) GetVPNServiceTasks(serviceID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNServiceTasksPaginated(serviceID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNServiceTasksPaginated retrieves a paginated list of VPN service tasks
func (s *Service) GetVPNServiceTasksPaginated(serviceID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getVPNServiceTasksPaginatedResponseBody(serviceID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNServiceTasksPaginated(serviceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNServiceTasksPaginatedResponseBody(serviceID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if serviceID == "" {
		return body, fmt.Errorf("invalid vpn service id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-services/%s/tasks", serviceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNServiceNotFoundError{ID: serviceID}
		}

		return nil
	})
}
