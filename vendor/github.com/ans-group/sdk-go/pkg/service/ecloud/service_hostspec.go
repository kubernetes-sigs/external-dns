package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHostSpecs retrieves a list of host specs
func (s *Service) GetHostSpecs(parameters connection.APIRequestParameters) ([]HostSpec, error) {
	return connection.InvokeRequestAll(s.GetHostSpecsPaginated, parameters)
}

// GetHostSpecsPaginated retrieves a paginated list of host specs
func (s *Service) GetHostSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostSpec], error) {
	body, err := s.getHostSpecsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetHostSpecsPaginated), err
}

func (s *Service) getHostSpecsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]HostSpec], error) {
	body := &connection.APIResponseBodyData[[]HostSpec]{}

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

func (s *Service) getHostSpecResponseBody(specID string) (*connection.APIResponseBodyData[HostSpec], error) {
	body := &connection.APIResponseBodyData[HostSpec]{}

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
