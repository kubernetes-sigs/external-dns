package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRegions retrieves a list of regions
func (s *Service) GetRegions(parameters connection.APIRequestParameters) ([]Region, error) {
	var regions []Region

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRegionsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, region := range response.(*PaginatedRegion).Items {
			regions = append(regions, region)
		}
	}

	return regions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRegionsPaginated retrieves a paginated list of regions
func (s *Service) GetRegionsPaginated(parameters connection.APIRequestParameters) (*PaginatedRegion, error) {
	body, err := s.getRegionsPaginatedResponseBody(parameters)

	return NewPaginatedRegion(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRegionsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRegionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRegionSliceResponseBody, error) {
	body := &GetRegionSliceResponseBody{}

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

func (s *Service) getRegionResponseBody(regionID string) (*GetRegionResponseBody, error) {
	body := &GetRegionResponseBody{}

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
