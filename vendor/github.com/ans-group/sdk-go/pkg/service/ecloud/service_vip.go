package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of vips
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return connection.InvokeRequestAll(s.GetVIPsPaginated, parameters)
}

// GetVIPsPaginated retrieves a paginated list of vips
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	body, err := s.getVIPsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVIPsPaginated), err
}

func (s *Service) getVIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VIP], error) {
	body := &connection.APIResponseBodyData[[]VIP]{}

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

func (s *Service) getVIPResponseBody(vipID string) (*connection.APIResponseBodyData[VIP], error) {
	body := &connection.APIResponseBodyData[VIP]{}

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

func (s *Service) createVIPResponseBody(req CreateVIPRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) patchVIPResponseBody(vipID string, req PatchVIPRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) deleteVIPResponseBody(vipID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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
