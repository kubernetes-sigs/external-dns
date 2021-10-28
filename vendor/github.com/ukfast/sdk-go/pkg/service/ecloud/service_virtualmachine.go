package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVirtualMachines retrieves a list of vms
func (s *Service) GetVirtualMachines(parameters connection.APIRequestParameters) ([]VirtualMachine, error) {
	var vms []VirtualMachine

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVirtualMachinesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, vm := range response.(*PaginatedVirtualMachine).Items {
			vms = append(vms, vm)
		}
	}

	return vms, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVirtualMachinesPaginated retrieves a paginated list of vms
func (s *Service) GetVirtualMachinesPaginated(parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error) {
	body, err := s.getVirtualMachinesPaginatedResponseBody(parameters)

	return NewPaginatedVirtualMachine(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVirtualMachinesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVirtualMachinesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVirtualMachineSliceResponseBody, error) {
	body := &GetVirtualMachineSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v1/vms", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVirtualMachine retrieves a single virtual machine by ID
func (s *Service) GetVirtualMachine(vmID int) (VirtualMachine, error) {
	body, err := s.getVirtualMachineResponseBody(vmID)

	return body.Data, err
}

func (s *Service) getVirtualMachineResponseBody(vmID int) (*GetVirtualMachineResponseBody, error) {
	body := &GetVirtualMachineResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/vms/%d", vmID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// DeleteVirtualMachine removes a virtual machine
func (s *Service) DeleteVirtualMachine(vmID int) error {
	_, err := s.deleteVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) deleteVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v1/vms/%d", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// CreateVirtualMachine creates a new virtual machine
func (s *Service) CreateVirtualMachine(req CreateVirtualMachineRequest) (int, error) {
	body, err := s.createVirtualMachineResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createVirtualMachineResponseBody(req CreateVirtualMachineRequest) (*GetVirtualMachineResponseBody, error) {
	body := &GetVirtualMachineResponseBody{}

	response, err := s.connection.Post("/ecloud/v1/vms", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVirtualMachine patches an eCloud virtual machine
func (s *Service) PatchVirtualMachine(vmID int, patch PatchVirtualMachineRequest) error {
	_, err := s.patchVirtualMachineResponseBody(vmID, patch)

	return err
}

func (s *Service) patchVirtualMachineResponseBody(vmID int, patch PatchVirtualMachineRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v1/vms/%d", vmID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// CloneVirtualMachine clones a virtual machine
func (s *Service) CloneVirtualMachine(vmID int, req CloneVirtualMachineRequest) (int, error) {
	body, err := s.cloneVirtualMachineResponseBody(vmID, req)

	return body.Data.ID, err
}

func (s *Service) cloneVirtualMachineResponseBody(vmID int, req CloneVirtualMachineRequest) (*GetVirtualMachineResponseBody, error) {
	body := &GetVirtualMachineResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/vms/%d/clone", vmID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PowerOnVirtualMachine powers on a virtual machine
func (s *Service) PowerOnVirtualMachine(vmID int) error {
	_, err := s.powerOnVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) powerOnVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/power-on", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PowerOffVirtualMachine powers off a virtual machine
func (s *Service) PowerOffVirtualMachine(vmID int) error {
	_, err := s.powerOffVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) powerOffVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/power-off", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PowerResetVirtualMachine resets a virtual machine (hard power off)
func (s *Service) PowerResetVirtualMachine(vmID int) error {
	_, err := s.powerResetVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) powerResetVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/power-reset", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PowerShutdownVirtualMachine shuts down a virtual machine
func (s *Service) PowerShutdownVirtualMachine(vmID int) error {
	_, err := s.powerShutdownVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) powerShutdownVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/power-shutdown", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PowerRestartVirtualMachine resets a virtual machine (graceful power off)
func (s *Service) PowerRestartVirtualMachine(vmID int) error {
	_, err := s.powerRestartVirtualMachineResponseBody(vmID)

	return err
}

func (s *Service) powerRestartVirtualMachineResponseBody(vmID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/power-restart", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// CreateVirtualMachineTemplate creates a virtual machine template
func (s *Service) CreateVirtualMachineTemplate(vmID int, req CreateVirtualMachineTemplateRequest) error {
	_, err := s.createVirtualMachineTemplateResponseBody(vmID, req)

	return err
}

func (s *Service) createVirtualMachineTemplateResponseBody(vmID int, req CreateVirtualMachineTemplateRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/vms/%d/clone-to-template", vmID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// GetVirtualMachineTags retrieves a list of tags
func (s *Service) GetVirtualMachineTags(vmID int, parameters connection.APIRequestParameters) ([]Tag, error) {
	var tags []Tag

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVirtualMachineTagsPaginated(vmID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, tag := range response.(*PaginatedTag).Items {
			tags = append(tags, tag)
		}
	}

	return tags, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVirtualMachineTagsPaginated retrieves a paginated list of domains
func (s *Service) GetVirtualMachineTagsPaginated(vmID int, parameters connection.APIRequestParameters) (*PaginatedTag, error) {
	body, err := s.getVirtualMachineTagsPaginatedResponseBody(vmID, parameters)

	return NewPaginatedTag(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVirtualMachineTagsPaginated(vmID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVirtualMachineTagsPaginatedResponseBody(vmID int, parameters connection.APIRequestParameters) (*GetTagSliceResponseBody, error) {
	body := &GetTagSliceResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/vms/%d/tags", vmID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// GetVirtualMachineTag retrieves a single virtual machine tag by key
func (s *Service) GetVirtualMachineTag(vmID int, tagKey string) (Tag, error) {
	body, err := s.getVirtualMachineTagResponseBody(vmID, tagKey)

	return body.Data, err
}

func (s *Service) getVirtualMachineTagResponseBody(vmID int, tagKey string) (*GetTagResponseBody, error) {
	body := &GetTagResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}

// CreateVirtualMachineTag creates a new virtual machine tag
func (s *Service) CreateVirtualMachineTag(vmID int, req CreateTagRequest) error {
	_, err := s.createVirtualMachineTagResponseBody(vmID, req)

	return err
}

func (s *Service) createVirtualMachineTagResponseBody(vmID int, req CreateTagRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/vms/%d/tags", vmID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}

// PatchVirtualMachineTag patches an eCloud virtual machine tag
func (s *Service) PatchVirtualMachineTag(vmID int, tagKey string, patch PatchTagRequest) error {
	_, err := s.patchVirtualMachineTagResponseBody(vmID, tagKey, patch)

	return err
}

func (s *Service) patchVirtualMachineTagResponseBody(vmID int, tagKey string, patch PatchTagRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}

// DeleteVirtualMachineTag removes a virtual machine tag
func (s *Service) DeleteVirtualMachineTag(vmID int, tagKey string) error {
	_, err := s.deleteVirtualMachineTagResponseBody(vmID, tagKey)

	return err
}

func (s *Service) deleteVirtualMachineTagResponseBody(vmID int, tagKey string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}

// CreateVirtualMachineConsoleSession creates a virtual machine console session
func (s *Service) CreateVirtualMachineConsoleSession(vmID int) (ConsoleSession, error) {
	body, err := s.createVirtualMachineConsoleSessionResponseBody(vmID)

	return body.Data, err
}

func (s *Service) createVirtualMachineConsoleSessionResponseBody(vmID int) (*GetConsoleSessionResponseBody, error) {
	body := &GetConsoleSessionResponseBody{}

	if vmID < 1 {
		return body, fmt.Errorf("invalid virtual machine id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v1/vms/%d/console-session", vmID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VirtualMachineNotFoundError{ID: vmID}
		}

		return nil
	})
}
