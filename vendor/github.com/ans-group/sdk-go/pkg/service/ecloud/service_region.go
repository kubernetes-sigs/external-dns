package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRegions retrieves a list of regions
func (s *Service) GetRegions(parameters connection.APIRequestParameters) ([]Region, error) {
	return connection.InvokeRequestAll(s.GetRegionsPaginated, parameters)
}

// GetRegionsPaginated retrieves a paginated list of regions
func (s *Service) GetRegionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Region], error) {
	body, err := s.getRegionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetRegionsPaginated), err
}

func (s *Service) getRegionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Region], error) {
	body := &connection.APIResponseBodyData[[]Region]{}

	response, err := s.connection.Get("/ecloud/v2/regions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRegion retrieves a single region by id
func (s *Service) GetRegion(regionID string) (Region, error) {
	body, err := s.getRegionResponseBody(regionID)

	return body.Data, err
}

func (s *Service) getRegionResponseBody(regionID string) (*connection.APIResponseBodyData[Region], error) {
	body := &connection.APIResponseBodyData[Region]{}

	if regionID == "" {
		return body, fmt.Errorf("invalid region id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/regions/%s", regionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RegionNotFoundError{ID: regionID}
		}

		return nil
	})
}
