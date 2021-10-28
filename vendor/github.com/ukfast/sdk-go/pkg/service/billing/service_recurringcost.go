package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetRecurringCosts retrieves a list of costs
func (s *Service) GetRecurringCosts(parameters connection.APIRequestParameters) ([]RecurringCost, error) {
	var costs []RecurringCost

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRecurringCostsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, cost := range response.(*PaginatedRecurringCost).Items {
			costs = append(costs, cost)
		}
	}

	return costs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetRecurringCostsPaginated retrieves a paginated list of costs
func (s *Service) GetRecurringCostsPaginated(parameters connection.APIRequestParameters) (*PaginatedRecurringCost, error) {
	body, err := s.getRecurringCostsPaginatedResponseBody(parameters)

	return NewPaginatedRecurringCost(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetRecurringCostsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getRecurringCostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetRecurringCostSliceResponseBody, error) {
	body := &GetRecurringCostSliceResponseBody{}

	response, err := s.connection.Get("/billing/v1/recurring-costs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetRecurringCost retrieves a single cost by id
func (s *Service) GetRecurringCost(costID int) (RecurringCost, error) {
	body, err := s.getRecurringCostResponseBody(costID)

	return body.Data, err
}

func (s *Service) getRecurringCostResponseBody(costID int) (*GetRecurringCostResponseBody, error) {
	body := &GetRecurringCostResponseBody{}

	if costID < 1 {
		return body, fmt.Errorf("invalid cost id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/recurring-costs/%d", costID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &RecurringCostNotFoundError{ID: costID}
		}

		return nil
	})
}
