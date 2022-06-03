package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAvailabilityZones retrieves a list of azs
func (s *Service) GetAvailabilityZones(parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return connection.InvokeRequestAll(s.GetAvailabilityZonesPaginated, parameters)
}

// GetAvailabilityZonesPaginated retrieves a paginated list of azs
func (s *Service) GetAvailabilityZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	body, err := s.getAvailabilityZonesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetAvailabilityZonesPaginated), err
}

func (s *Service) getAvailabilityZonesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AvailabilityZone], error) {
	body := &connection.APIResponseBodyData[[]AvailabilityZone]{}

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

func (s *Service) getAvailabilityZoneResponseBody(azID string) (*connection.APIResponseBodyData[AvailabilityZone], error) {
	body := &connection.APIResponseBodyData[AvailabilityZone]{}

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
