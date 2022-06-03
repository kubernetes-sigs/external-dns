package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVolumes retrieves a list of volumes
func (s *Service) GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error) {
	return connection.InvokeRequestAll(s.GetVolumesPaginated, parameters)
}

// GetVolumesPaginated retrieves a paginated list of volumes
func (s *Service) GetVolumesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	body, err := s.getVolumesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVolumesPaginated), err
}

func (s *Service) getVolumesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Volume], error) {
	body := &connection.APIResponseBodyData[[]Volume]{}

	response, err := s.connection.Get("/ecloud/v2/volumes", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID string) (Volume, error) {
	body, err := s.getVolumeResponseBody(volumeID)

	return body.Data, err
}

func (s *Service) getVolumeResponseBody(volumeID string) (*connection.APIResponseBodyData[Volume], error) {
	body := &connection.APIResponseBodyData[Volume]{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// CreateVolume creates a volume
func (s *Service) CreateVolume(req CreateVolumeRequest) (TaskReference, error) {
	body, err := s.createVolumeResponseBody(req)

	return body.Data, err
}

func (s *Service) createVolumeResponseBody(req CreateVolumeRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/volumes", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVolume patches a Volume
func (s *Service) PatchVolume(volumeID string, req PatchVolumeRequest) (TaskReference, error) {
	body, err := s.patchVolumeResponseBody(volumeID, req)

	return body.Data, err
}

func (s *Service) patchVolumeResponseBody(volumeID string, req PatchVolumeRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// DeleteVolume deletes a Volume
func (s *Service) DeleteVolume(volumeID string) (string, error) {
	body, err := s.deleteVolumeResponseBody(volumeID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVolumeResponseBody(volumeID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// GetVolumeInstances retrieves a list of volume instances
func (s *Service) GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}, parameters)
}

// GetVolumeInstancesPaginated retrieves a paginated list of volume instances
func (s *Service) GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	body, err := s.getVolumeInstancesPaginatedResponseBody(volumeID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}), err
}

func (s *Service) getVolumeInstancesPaginatedResponseBody(volumeID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Instance], error) {
	body := &connection.APIResponseBodyData[[]Instance]{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s/instances", volumeID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// GetVolumeTasks retrieves a list of Volume tasks
func (s *Service) GetVolumeTasks(volumeID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}, parameters)
}

// GetVolumeTasksPaginated retrieves a paginated list of Volume tasks
func (s *Service) GetVolumeTasksPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getVolumeTasksPaginatedResponseBody(volumeID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}), err
}

func (s *Service) getVolumeTasksPaginatedResponseBody(volumeID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s/tasks", volumeID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}
