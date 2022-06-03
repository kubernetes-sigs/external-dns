package draas

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBillingTypes retrieves a list of solutions
func (s *Service) GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error) {
	return connection.InvokeRequestAll(s.GetBillingTypesPaginated, parameters)
}

// GetBillingTypesPaginated retrieves a paginated list of solutions
func (s *Service) GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingType], error) {
	body, err := s.getBillingTypesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetBillingTypesPaginated), err
}

func (s *Service) getBillingTypesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]BillingType], error) {
	body := &connection.APIResponseBodyData[[]BillingType]{}

	response, err := s.connection.Get("/draas/v1/billing-types", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetBillingType retrieves a single solution by id
func (s *Service) GetBillingType(billingTypeID string) (BillingType, error) {
	body, err := s.getBillingTypeResponseBody(billingTypeID)

	return body.Data, err
}

func (s *Service) getBillingTypeResponseBody(billingTypeID string) (*connection.APIResponseBodyData[BillingType], error) {
	body := &connection.APIResponseBodyData[BillingType]{}

	if billingTypeID == "" {
		return body, fmt.Errorf("invalid billing type id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/draas/v1/billing-types/%s", billingTypeID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BillingTypeNotFoundError{ID: billingTypeID}
		}

		return nil
	})
}
