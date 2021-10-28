package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetCloudCosts retrieves a list of costs
func (s *Service) GetCloudCosts(parameters connection.APIRequestParameters) ([]CloudCost, error) {
	var costs []CloudCost

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCloudCostsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, cost := range response.(*PaginatedCloudCost).Items {
			costs = append(costs, cost)
		}
	}

	return costs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetCloudCostsPaginated retrieves a paginated list of costs
func (s *Service) GetCloudCostsPaginated(parameters connection.APIRequestParameters) (*PaginatedCloudCost, error) {
	body, err := s.getCloudCostsPaginatedResponseBody(parameters)

	return NewPaginatedCloudCost(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCloudCostsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getCloudCostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetCloudCostSliceResponseBody, error) {
	body := &GetCloudCostSliceResponseBody{}

	response, err := s.connection.Get("/billing/v1/cloud-costs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCloudCost retrieves a single cost by id
func (s *Service) GetCloudCost(costID int) (CloudCost, error) {
	body, err := s.getCloudCostResponseBody(costID)

	return body.Data, err
}

func (s *Service) getCloudCostResponseBody(costID int) (*GetCloudCostResponseBody, error) {
	body := &GetCloudCostResponseBody{}

	if costID < 1 {
		return body, fmt.Errorf("invalid cost id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/cloud-costs/%d", costID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CloudCostNotFoundError{ID: costID}
		}

		return nil
	})
}
