package storage

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

func (s *Service) getVolumeResponseBody(volumeID int) (*connection.APIResponseBodyData[Volume], error) {
	body := &connection.APIResponseBodyData[Volume]{}

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
