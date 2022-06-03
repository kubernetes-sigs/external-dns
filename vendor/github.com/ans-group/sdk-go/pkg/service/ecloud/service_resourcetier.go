package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetResourceTiers retrieves a list of resource tiers
func (s *Service) GetResourceTiers(parameters connection.APIRequestParameters) ([]ResourceTier, error) {
	return connection.InvokeRequestAll(s.GetResourceTiersPaginated, parameters)
}

// GetResourceTiersPaginated retrieves a paginated list of resource tiers
func (s *Service) GetResourceTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ResourceTier], error) {
	body, err := s.getResourceTiersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetResourceTiersPaginated), err
}

func (s *Service) getResourceTiersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ResourceTier], error) {
	body := &connection.APIResponseBodyData[[]ResourceTier]{}

	response, err := s.connection.Get("/ecloud/v2/resource-tiers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetResourceTier retrieves a single resource tier by id
func (s *Service) GetResourceTier(tierID string) (ResourceTier, error) {
	body, err := s.getResourceTierResponseBody(tierID)

	return body.Data, err
}

func (s *Service) getResourceTierResponseBody(tierID string) (*connection.APIResponseBodyData[ResourceTier], error) {
	body := &connection.APIResponseBodyData[ResourceTier]{}

	if tierID == "" {
		return body, fmt.Errorf("invalid resource tier id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/resource-tiers/%s", tierID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ResourceTierNotFoundError{ID: tierID}
		}

		return nil
	})
}
