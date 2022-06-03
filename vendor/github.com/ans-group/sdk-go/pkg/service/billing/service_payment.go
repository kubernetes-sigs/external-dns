package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetPayments retrieves a list of payments
func (s *Service) GetPayments(parameters connection.APIRequestParameters) ([]Payment, error) {
	return connection.InvokeRequestAll(s.GetPaymentsPaginated, parameters)
}

// GetPaymentsPaginated retrieves a paginated list of payments
func (s *Service) GetPaymentsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Payment], error) {
	body, err := s.getPaymentsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetPaymentsPaginated), err
}

func (s *Service) getPaymentsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Payment], error) {
	body := &connection.APIResponseBodyData[[]Payment]{}

	response, err := s.connection.Get("/billing/v1/payments", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetPayment retrieves a single payment by id
func (s *Service) GetPayment(paymentID int) (Payment, error) {
	body, err := s.getPaymentResponseBody(paymentID)

	return body.Data, err
}

func (s *Service) getPaymentResponseBody(paymentID int) (*connection.APIResponseBodyData[Payment], error) {
	body := &connection.APIResponseBodyData[Payment]{}

	if paymentID < 1 {
		return body, fmt.Errorf("invalid payment id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/payments/%d", paymentID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &PaymentNotFoundError{ID: paymentID}
		}

		return nil
	})
}
