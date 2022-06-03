package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetInvoices retrieves a list of invoices
func (s *Service) GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error) {
	return connection.InvokeRequestAll(s.GetInvoicesPaginated, parameters)
}

// GetInvoicesPaginated retrieves a paginated list of invoices
func (s *Service) GetInvoicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Invoice], error) {
	body, err := s.getInvoicesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoicesPaginated), err
}

func (s *Service) getInvoicesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Invoice], error) {
	body := &connection.APIResponseBodyData[[]Invoice]{}

	response, err := s.connection.Get("/billing/v1/invoices", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetInvoice retrieves a single invoice by id
func (s *Service) GetInvoice(invoiceID int) (Invoice, error) {
	body, err := s.getInvoiceResponseBody(invoiceID)

	return body.Data, err
}

func (s *Service) getInvoiceResponseBody(invoiceID int) (*connection.APIResponseBodyData[Invoice], error) {
	body := &connection.APIResponseBodyData[Invoice]{}

	if invoiceID < 1 {
		return body, fmt.Errorf("invalid invoice id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/invoices/%d", invoiceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InvoiceNotFoundError{ID: invoiceID}
		}

		return nil
	})
}
