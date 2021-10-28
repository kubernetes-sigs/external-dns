package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSites retrieves a list of sites
func (s *Service) GetSites(parameters connection.APIRequestParameters) ([]Site, error) {
	var sites []Site

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSitesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedSite).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSitesPaginated retrieves a paginated list of sites
func (s *Service) GetSitesPaginated(parameters connection.APIRequestParameters) (*PaginatedSite, error) {
	body, err := s.getSitesPaginatedResponseBody(parameters)

	return NewPaginatedSite(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSitesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSitesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSiteSliceResponseBody, error) {
	body := &GetSiteSliceResponseBody{}

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

func (s *Service) getSiteResponseBody(siteID int) (*GetSiteResponseBody, error) {
	body := &GetSiteResponseBody{}

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
