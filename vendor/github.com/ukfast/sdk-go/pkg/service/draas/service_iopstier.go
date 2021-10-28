package draas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetIOPSTiers retrieves a list of solutions
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	var tiers []IOPSTier

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetIOPSTiersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, tier := range response.(*PaginatedIOPSTier).Items {
			tiers = append(tiers, tier)
		}
	}

	return tiers, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetIOPSTiersPaginated retrieves a paginated list of solutions
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*PaginatedIOPSTier, error) {
	body, err := s.getIOPSTiersPaginatedResponseBody(parameters)

	return NewPaginatedIOPSTier(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetIOPSTiersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getIOPSTiersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetIOPSTierSliceResponseBody, error) {
	body := &GetIOPSTierSliceResponseBody{}

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

func (s *Service) getIOPSTierResponseBody(iopsTierID string) (*GetIOPSTierResponseBody, error) {
	body := &GetIOPSTierResponseBody{}

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
