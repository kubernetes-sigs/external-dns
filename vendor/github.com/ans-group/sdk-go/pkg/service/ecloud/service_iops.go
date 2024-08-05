package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIOPSs retrieves a list of iops
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(s.GetIOPSTiersPaginated, parameters)
}

// GetIOPSsPaginated retrieves a paginated list of iops
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	body, err := s.getIOPSsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIOPSTiersPaginated), err
}

func (s *Service) getIOPSsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]IOPSTier], error) {
	body := &connection.APIResponseBodyData[[]IOPSTier]{}

	response, err := s.connection.Get("/ecloud/v2/iops", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetIOPS retrieves a single IOPS by ID
func (s *Service) GetIOPSTier(iopsID string) (IOPSTier, error) {
	body, err := s.getIOPSTierResponseBody(iopsID)

	return body.Data, err
}

func (s *Service) getIOPSTierResponseBody(iopsID string) (*connection.APIResponseBodyData[IOPSTier], error) {
	body := &connection.APIResponseBodyData[IOPSTier]{}

	if iopsID == "" {
		return body, fmt.Errorf("invalid IOPS id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/iops/%s", iopsID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &IOPSNotFoundError{ID: iopsID}
		}

		return nil
	})
}
