package draas

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIOPSTiers retrieves a list of solutions
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(s.GetIOPSTiersPaginated, parameters)
}

// GetIOPSTiersPaginated retrieves a paginated list of solutions
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	body, err := s.getIOPSTiersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIOPSTiersPaginated), err
}

func (s *Service) getIOPSTiersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]IOPSTier], error) {
	body := &connection.APIResponseBodyData[[]IOPSTier]{}

	response, err := s.connection.Get("/draas/v1/iops-tiers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetIOPSTier retrieves a single solution by id
func (s *Service) GetIOPSTier(iopsTierID string) (IOPSTier, error) {
	body, err := s.getIOPSTierResponseBody(iopsTierID)

	return body.Data, err
}

func (s *Service) getIOPSTierResponseBody(iopsTierID string) (*connection.APIResponseBodyData[IOPSTier], error) {
	body := &connection.APIResponseBodyData[IOPSTier]{}

	if iopsTierID == "" {
		return body, fmt.Errorf("invalid iops tier id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/iops-tiers/%s", iopsTierID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &IOPSTierNotFoundError{ID: iopsTierID}
		}

		return nil
	})
}
