package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetPayments retrieves a list of payments
func (s *Service) GetPayments(parameters connection.APIRequestParameters) ([]Payment, error) {
	var payments []Payment

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetPaymentsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, payment := range response.(*PaginatedPayment).Items {
			payments = append(payments, payment)
		}
	}

	return payments, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetPaymentsPaginated retrieves a paginated list of payments
func (s *Service) GetPaymentsPaginated(parameters connection.APIRequestParameters) (*PaginatedPayment, error) {
	body, err := s.getPaymentsPaginatedResponseBody(parameters)

	return NewPaginatedPayment(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetPaymentsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getPaymentsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetPaymentSliceResponseBody, error) {
	body := &GetPaymentSliceResponseBody{}

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

func (s *Service) getPaymentResponseBody(paymentID int) (*GetPaymentResponseBody, error) {
	body := &GetPaymentResponseBody{}

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
