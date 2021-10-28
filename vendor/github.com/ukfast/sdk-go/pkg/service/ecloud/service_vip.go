package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of vips
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	var vips []VIP

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVIPsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, vip := range response.(*PaginatedVIP).Items {
			vips = append(vips, vip)
		}
	}

	return vips, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVIPsPaginated retrieves a paginated list of vips
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedVIP, error) {
	body, err := s.getVIPsPaginatedResponseBody(parameters)

	return NewPaginatedVIP(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVIPsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVIPSliceResponseBody, error) {
	body := &GetVIPSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/vips", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVIP retrieves a single vip by id
func (s *Service) GetVIP(vipID string) (VIP, error) {
	body, err := s.getVIPResponseBody(vipID)

	return body.Data, err
}

func (s *Service) getVIPResponseBody(vipID string) (*GetVIPResponseBody, error) {
	body := &GetVIPResponseBody{}

	if vipID == "" {
		return body, fmt.Errorf("invalid vip id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vips/%s", vipID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VIPNotFoundError{ID: vipID}
		}

		return nil
	})
}

// CreateVIP creates a new VIP
func (s *Service) CreateVIP(req CreateVIPRequest) (TaskReference, error) {
	body, err := s.createVIPResponseBody(req)

	return body.Data, err
}

func (s *Service) createVIPResponseBody(req CreateVIPRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/vips", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVIP patches a VIP
func (s *Service) PatchVIP(vipID string, req PatchVIPRequest) (TaskReference, error) {
	body, err := s.patchVIPResponseBody(vipID, req)

	return body.Data, err
}

func (s *Service) patchVIPResponseBody(vipID string, req PatchVIPRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if vipID == "" {
		return body, fmt.Errorf("invalid vip id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vips/%s", vipID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VIPNotFoundError{ID: vipID}
		}

		return nil
	})
}

// DeleteVIP deletes a VIP
func (s *Service) DeleteVIP(vipID string) (string, error) {
	body, err := s.deleteVIPResponseBody(vipID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVIPResponseBody(vipID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if vipID == "" {
		return body, fmt.Errorf("invalid vip id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vips/%s", vipID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VIPNotFoundError{ID: vipID}
		}

		return nil
	})
}
