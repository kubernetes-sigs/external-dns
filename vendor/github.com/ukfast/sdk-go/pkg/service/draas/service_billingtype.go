package draas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetBillingTypes retrieves a list of solutions
func (s *Service) GetBillingTypes(parameters connection.APIRequestParameters) ([]BillingType, error) {
	var billingTypes []BillingType

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBillingTypesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, billingType := range response.(*PaginatedBillingType).Items {
			billingTypes = append(billingTypes, billingType)
		}
	}

	return billingTypes, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetBillingTypesPaginated retrieves a paginated list of solutions
func (s *Service) GetBillingTypesPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingType, error) {
	body, err := s.getBillingTypesPaginatedResponseBody(parameters)

	return NewPaginatedBillingType(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBillingTypesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getBillingTypesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetBillingTypeSliceResponseBody, error) {
	body := &GetBillingTypeSliceResponseBody{}

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

func (s *Service) getBillingTypeResponseBody(billingTypeID string) (*GetBillingTypeResponseBody, error) {
	body := &GetBillingTypeResponseBody{}

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
