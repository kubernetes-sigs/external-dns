package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAvailabilityZones retrieves a list of azs
func (s *Service) GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	var azs []AvailabilityZone

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAvailabilityZonesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, az := range response.(*PaginatedAvailabilityZone).Items {
			azs = append(azs, az)
		}
	}

	return azs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetAvailabilityZonesPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*PaginatedAvailabilityZone, error) {
	body, err := s.getAvailabilityZonesPaginatedResponseBody(parameters)

	return NewPaginatedAvailabilityZone(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAvailabilityZonesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getAvailabilityZonesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetAvailabilityZoneSliceResponseBody, error) {
	body := &GetAvailabilityZoneSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/availability-zones", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAvailabilityZone retrieves a single az by id
func (s *Service) GetAvailabilityZone(azID string) (AvailabilityZone, error) {
	body, err := s.getAvailabilityZoneResponseBody(azID)

	return body.Data, err
}

func (s *Service) getAvailabilityZoneResponseBody(azID string) (*GetAvailabilityZoneResponseBody, error) {
	body := &GetAvailabilityZoneResponseBody{}

	if azID == "" {
		return body, fmt.Errorf("invalid az id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/availability-zones/%s", azID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AvailabilityZoneNotFoundError{ID: azID}
		}

		return nil
	})
}
