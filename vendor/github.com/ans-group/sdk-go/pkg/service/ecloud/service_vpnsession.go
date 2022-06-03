package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNSessions retrieves a list of VPN sessions
func (s *Service) GetVPNSessions(parameters connection.APIRequestParameters) ([]VPNSession, error) {
	return connection.InvokeRequestAll(s.GetVPNSessionsPaginated, parameters)
}

// GetVPNSessionsPaginated retrieves a paginated list of VPN sessions
func (s *Service) GetVPNSessionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNSession], error) {
	body, err := s.getVPNSessionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNSessionsPaginated), err
}

func (s *Service) getVPNSessionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPNSession], error) {
	body := &connection.APIResponseBodyData[[]VPNSession]{}

	response, err := s.connection.Get("/ecloud/v2/vpn-sessions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNSession retrieves a single VPN session by id
func (s *Service) GetVPNSession(sessionID string) (VPNSession, error) {
	body, err := s.getVPNSessionResponseBody(sessionID)

	return body.Data, err
}

func (s *Service) getVPNSessionResponseBody(sessionID string) (*connection.APIResponseBodyData[VPNSession], error) {
	body := &connection.APIResponseBodyData[VPNSession]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}

// CreateVPNSession creates a new VPN session
func (s *Service) CreateVPNSession(req CreateVPNSessionRequest) (TaskReference, error) {
	body, err := s.createVPNSessionResponseBody(req)

	return body.Data, err
}

func (s *Service) createVPNSessionResponseBody(req CreateVPNSessionRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/vpn-sessions", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPNSession patches a VPN session
func (s *Service) PatchVPNSession(sessionID string, req PatchVPNSessionRequest) (TaskReference, error) {
	body, err := s.patchVPNSessionResponseBody(sessionID, req)

	return body.Data, err
}

func (s *Service) patchVPNSessionResponseBody(sessionID string, req PatchVPNSessionRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}

// DeleteVPNSession deletes a VPN session
func (s *Service) DeleteVPNSession(sessionID string) (string, error) {
	body, err := s.deleteVPNSessionResponseBody(sessionID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVPNSessionResponseBody(sessionID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}

// GetVPNSessionTasks retrieves a list of VPN session tasks
func (s *Service) GetVPNSessionTasks(sessionID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}, parameters)
}

// GetVPNSessionTasksPaginated retrieves a paginated list of VPN session tasks
func (s *Service) GetVPNSessionTasksPaginated(sessionID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getVPNSessionTasksPaginatedResponseBody(sessionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}), err
}

func (s *Service) getVPNSessionTasksPaginatedResponseBody(sessionID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/tasks", sessionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}

// GetVPNSession retrieves a single VPN session by id
func (s *Service) GetVPNSessionPreSharedKey(sessionID string) (VPNSessionPreSharedKey, error) {
	body, err := s.getVPNSessionPreSharedKeyResponseBody(sessionID)

	return body.Data, err
}

func (s *Service) getVPNSessionPreSharedKeyResponseBody(sessionID string) (*connection.APIResponseBodyData[VPNSessionPreSharedKey], error) {
	body := &connection.APIResponseBodyData[VPNSessionPreSharedKey]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/pre-shared-key", sessionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}

// UpdateVPNSession retrieves a single VPN session by id
func (s *Service) UpdateVPNSessionPreSharedKey(sessionID string, req UpdateVPNSessionPreSharedKeyRequest) (TaskReference, error) {
	body, err := s.updateVPNSessionPreSharedKeyResponseBody(sessionID, req)

	return body.Data, err
}

func (s *Service) updateVPNSessionPreSharedKeyResponseBody(sessionID string, req UpdateVPNSessionPreSharedKeyRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid vpn session id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/pre-shared-key", sessionID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNSessionNotFoundError{ID: sessionID}
		}

		return nil
	})
}
