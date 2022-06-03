package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSites retrieves a list of sites
func (s *Service) GetSites(parameters connection.APIRequestParameters) ([]Site, error) {
	return connection.InvokeRequestAll(s.GetSitesPaginated, parameters)
}

// GetSitesPaginated retrieves a paginated list of sites
func (s *Service) GetSitesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Site], error) {
	body, err := s.getSitesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSitesPaginated), err
}

func (s *Service) getSitesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Site], error) {
	body := &connection.APIResponseBodyData[[]Site]{}

	response, err := s.connection.Get("/ecloud/v1/sites", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSite retrieves a single site by ID
func (s *Service) GetSite(siteID int) (Site, error) {
	body, err := s.getSiteResponseBody(siteID)

	return body.Data, err
}

func (s *Service) getSiteResponseBody(siteID int) (*connection.APIResponseBodyData[Site], error) {
	body := &connection.APIResponseBodyData[Site]{}

	if siteID < 1 {
		return body, fmt.Errorf("invalid site id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/sites/%d", siteID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SiteNotFoundError{ID: siteID}
		}

		return nil
	})
}
