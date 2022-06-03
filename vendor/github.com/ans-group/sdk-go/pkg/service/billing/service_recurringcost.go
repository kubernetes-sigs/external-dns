package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRecurringCosts retrieves a list of costs
func (s *Service) GetRecurringCosts(parameters connection.APIRequestParameters) ([]RecurringCost, error) {
	return connection.InvokeRequestAll(s.GetRecurringCostsPaginated, parameters)
}

// GetRecurringCostsPaginated retrieves a paginated list of costs
func (s *Service) GetRecurringCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RecurringCost], error) {
	body, err := s.getRecurringCostsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetRecurringCostsPaginated), err
}

func (s *Service) getRecurringCostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]RecurringCost], error) {
	body := &connection.APIResponseBodyData[[]RecurringCost]{}

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

func (s *Service) getRecurringCostResponseBody(costID int) (*connection.APIResponseBodyData[RecurringCost], error) {
	body := &connection.APIResponseBodyData[RecurringCost]{}

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
