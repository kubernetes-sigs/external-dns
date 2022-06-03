package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPCs retrieves a list of vpcs
func (s *Service) GetVPCs(parameters connection.APIRequestParameters) ([]VPC, error) {
	return connection.InvokeRequestAll(s.GetVPCsPaginated, parameters)
}

// GetVPCsPaginated retrieves a paginated list of vpcs
func (s *Service) GetVPCsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPC], error) {
	body, err := s.getVPCsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPCsPaginated), err
}

func (s *Service) getVPCsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPC], error) {
	body := &connection.APIResponseBodyData[[]VPC]{}

	response, err := s.connection.Get("/ecloud/v2/vpcs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPC retrieves a single vpc by id
func (s *Service) GetVPC(vpcID string) (VPC, error) {
	body, err := s.getVPCResponseBody(vpcID)

	return body.Data, err
}

func (s *Service) getVPCResponseBody(vpcID string) (*connection.APIResponseBodyData[VPC], error) {
	body := &connection.APIResponseBodyData[VPC]{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// CreateVPC creates a new VPC
func (s *Service) CreateVPC(req CreateVPCRequest) (string, error) {
	body, err := s.createVPCResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createVPCResponseBody(req CreateVPCRequest) (*connection.APIResponseBodyData[VPC], error) {
	body := &connection.APIResponseBodyData[VPC]{}

	response, err := s.connection.Post("/ecloud/v2/vpcs", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPC patches a VPC
func (s *Service) PatchVPC(vpcID string, req PatchVPCRequest) error {
	_, err := s.patchVPCResponseBody(vpcID, req)

	return err
}

func (s *Service) patchVPCResponseBody(vpcID string, req PatchVPCRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// DeleteVPC deletes a VPC
func (s *Service) DeleteVPC(vpcID string) error {
	_, err := s.deleteVPCResponseBody(vpcID)

	return err
}

func (s *Service) deleteVPCResponseBody(vpcID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// DeployVPCDefaults deploys default resources for specified VPC
func (s *Service) DeployVPCDefaults(vpcID string) error {
	_, err := s.deployVPCDefaultsResponseBody(vpcID)

	return err
}

func (s *Service) deployVPCDefaultsResponseBody(vpcID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/vpcs/%s/deploy-defaults", vpcID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// GetVPCVolumes retrieves a list of firewall rule volumes
func (s *Service) GetVPCVolumes(vpcID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVPCVolumesPaginated(vpcID, p)
	}, parameters)
}

// GetVPCVolumesPaginated retrieves a paginated list of firewall rule volumes
func (s *Service) GetVPCVolumesPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	body, err := s.getVPCVolumesPaginatedResponseBody(vpcID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVPCVolumesPaginated(vpcID, p)
	}), err
}

func (s *Service) getVPCVolumesPaginatedResponseBody(vpcID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Volume], error) {
	body := &connection.APIResponseBodyData[[]Volume]{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpcs/%s/volumes", vpcID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// GetVPCInstances retrieves a list of firewall rule instances
func (s *Service) GetVPCInstances(vpcID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVPCInstancesPaginated(vpcID, p)
	}, parameters)
}

// GetVPCInstancesPaginated retrieves a paginated list of firewall rule instances
func (s *Service) GetVPCInstancesPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	body, err := s.getVPCInstancesPaginatedResponseBody(vpcID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVPCInstancesPaginated(vpcID, p)
	}), err
}

func (s *Service) getVPCInstancesPaginatedResponseBody(vpcID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Instance], error) {
	body := &connection.APIResponseBodyData[[]Instance]{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpcs/%s/instances", vpcID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// GetVPCTasks retrieves a list of VPC tasks
func (s *Service) GetVPCTasks(vpcID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPCTasksPaginated(vpcID, p)
	}, parameters)
}

// GetVPCTasksPaginated retrieves a paginated list of VPC tasks
func (s *Service) GetVPCTasksPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getVPCTasksPaginatedResponseBody(vpcID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPCTasksPaginated(vpcID, p)
	}), err
}

func (s *Service) getVPCTasksPaginatedResponseBody(vpcID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpcs/%s/tasks", vpcID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}
