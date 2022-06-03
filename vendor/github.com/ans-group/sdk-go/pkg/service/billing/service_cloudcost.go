package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCloudCosts retrieves a list of costs
func (s *Service) GetCloudCosts(parameters connection.APIRequestParameters) ([]CloudCost, error) {
	return connection.InvokeRequestAll(s.GetCloudCostsPaginated, parameters)
}

// GetCloudCostsPaginated retrieves a paginated list of costs
func (s *Service) GetCloudCostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CloudCost], error) {
	body, err := s.getCloudCostsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetCloudCostsPaginated), err
}

func (s *Service) getCloudCostsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CloudCost], error) {
	body := &connection.APIResponseBodyData[[]CloudCost]{}

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

func (s *Service) getCloudCostResponseBody(costID int) (*connection.APIResponseBodyData[CloudCost], error) {
	body := &connection.APIResponseBodyData[CloudCost]{}

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
