package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetInvoices retrieves a list of invoices
func (s *Service) GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error) {
	var invoices []Invoice

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInvoicesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, invoice := range response.(*PaginatedInvoice).Items {
			invoices = append(invoices, invoice)
		}
	}

	return invoices, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInvoicesPaginated retrieves a paginated list of invoices
func (s *Service) GetInvoicesPaginated(parameters connection.APIRequestParameters) (*PaginatedInvoice, error) {
	body, err := s.getInvoicesPaginatedResponseBody(parameters)

	return NewPaginatedInvoice(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInvoicesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInvoicesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetInvoiceSliceResponseBody, error) {
	body := &GetInvoiceSliceResponseBody{}

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

func (s *Service) getInvoiceResponseBody(invoiceID int) (*GetInvoiceResponseBody, error) {
	body := &GetInvoiceResponseBody{}

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
