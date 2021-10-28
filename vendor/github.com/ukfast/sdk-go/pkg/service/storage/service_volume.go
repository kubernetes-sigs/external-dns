package storage

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVolumes retrieves a list of volumes
func (s *Service) GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error) {
	var volumes []Volume

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, volume := range response.(*PaginatedVolume).Items {
			volumes = append(volumes, volume)
		}
	}

	return volumes, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumesPaginated retrieves a paginated list of volumes
func (s *Service) GetVolumesPaginated(parameters connection.APIRequestParameters) (*PaginatedVolume, error) {
	body, err := s.getVolumesPaginatedResponseBody(parameters)

	return NewPaginatedVolume(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVolumeSliceResponseBody, error) {
	body := &GetVolumeSliceResponseBody{}

	response, err := s.connection.Get("/ukfast-storage/v1/volumes", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID int) (Volume, error) {
	body, err := s.getVolumeResponseBody(volumeID)

	return body.Data, err
}

func (s *Service) getVolumeResponseBody(volumeID int) (*GetVolumeResponseBody, error) {
	body := &GetVolumeResponseBody{}

	if volumeID < 1 {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ukfast-storage/v1/volumes/%d", volumeID), connection.APIRequestParameters{})
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
