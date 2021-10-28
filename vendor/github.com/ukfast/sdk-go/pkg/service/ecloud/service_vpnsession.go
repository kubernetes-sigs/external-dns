package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPNSessions retrieves a list of VPN sessions
func (s *Service) GetVPNSessions(parameters connection.APIRequestParameters) ([]VPNSession, error) {
	var sessions []VPNSession

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNSessionsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, session := range response.(*PaginatedVPNSession).Items {
			sessions = append(sessions, session)
		}
	}

	return sessions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNSessionsPaginated retrieves a paginated list of VPN sessions
func (s *Service) GetVPNSessionsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNSession, error) {
	body, err := s.getVPNSessionsPaginatedResponseBody(parameters)

	return NewPaginatedVPNSession(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNSessionsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNSessionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPNSessionSliceResponseBody, error) {
	body := &GetVPNSessionSliceResponseBody{}

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

func (s *Service) getVPNSessionResponseBody(sessionID string) (*GetVPNSessionResponseBody, error) {
	body := &GetVPNSessionResponseBody{}

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

func (s *Service) createVPNSessionResponseBody(req CreateVPNSessionRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

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

func (s *Service) patchVPNSessionResponseBody(sessionID string, req PatchVPNSessionRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid session id")
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

func (s *Service) deleteVPNSessionResponseBody(sessionID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if sessionID == "" {
		return body, fmt.Errorf("invalid session id")
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
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNSessionTasksPaginated retrieves a paginated list of VPN session tasks
func (s *Service) GetVPNSessionTasksPaginated(sessionID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getVPNSessionTasksPaginatedResponseBody(sessionID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNSessionTasksPaginatedResponseBody(sessionID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

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

func (s *Service) getVPNSessionPreSharedKeyResponseBody(sessionID string) (*GetVPNSessionPreSharedKeyResponseBody, error) {
	body := &GetVPNSessionPreSharedKeyResponseBody{}

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
