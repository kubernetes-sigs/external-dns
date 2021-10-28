package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetHostSpecs retrieves a list of host specs
func (s *Service) GetHostSpecs(parameters connection.APIRequestParameters) ([]HostSpec, error) {
	var specs []HostSpec

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostSpecsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, spec := range response.(*PaginatedHostSpec).Items {
			specs = append(specs, spec)
		}
	}

	return specs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHostSpecsPaginated retrieves a paginated list of host specs
func (s *Service) GetHostSpecsPaginated(parameters connection.APIRequestParameters) (*PaginatedHostSpec, error) {
	body, err := s.getHostSpecsPaginatedResponseBody(parameters)

	return NewPaginatedHostSpec(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHostSpecsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHostSpecsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetHostSpecSliceResponseBody, error) {
	body := &GetHostSpecSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/host-specs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetHostSpec retrieves a single host spec by id
func (s *Service) GetHostSpec(specID string) (HostSpec, error) {
	body, err := s.getHostSpecResponseBody(specID)

	return body.Data, err
}

func (s *Service) getHostSpecResponseBody(specID string) (*GetHostSpecResponseBody, error) {
	body := &GetHostSpecResponseBody{}

	if specID == "" {
		return body, fmt.Errorf("invalid spec id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/host-specs/%s", specID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HostSpecNotFoundError{ID: specID}
		}

		return nil
	})
}
