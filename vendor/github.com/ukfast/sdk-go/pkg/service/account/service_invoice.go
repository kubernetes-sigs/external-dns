package account

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

	response, err := s.connection.Get("/account/v1/invoices", parameters)
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

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/invoices/%d", invoiceID), connection.APIRequestParameters{})
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

// GetInvoiceQueries retrieves a list of invoice queries
func (s *Service) GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error) {
	var queries []InvoiceQuery

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInvoiceQueriesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, query := range response.(*PaginatedInvoiceQuery).Items {
			queries = append(queries, query)
		}
	}

	return queries, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInvoiceQueriesPaginated retrieves a paginated list of invoice queries
func (s *Service) GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*PaginatedInvoiceQuery, error) {
	body, err := s.getInvoiceQueriesPaginatedResponseBody(parameters)

	return NewPaginatedInvoiceQuery(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInvoiceQueriesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInvoiceQueriesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetInvoiceQuerySliceResponseBody, error) {
	body := &GetInvoiceQuerySliceResponseBody{}

	response, err := s.connection.Get("/account/v1/invoice-queries", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetInvoiceQuery retrieves a single invoice query by id
func (s *Service) GetInvoiceQuery(queryID int) (InvoiceQuery, error) {
	body, err := s.getInvoiceQueryResponseBody(queryID)

	return body.Data, err
}

func (s *Service) getInvoiceQueryResponseBody(queryID int) (*GetInvoiceQueryResponseBody, error) {
	body := &GetInvoiceQueryResponseBody{}

	if queryID < 1 {
		return body, fmt.Errorf("invalid invoice query id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/invoice-queries/%d", queryID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InvoiceQueryNotFoundError{ID: queryID}
		}

		return nil
	})
}

// CreateInvoiceQuery retrieves creates an InvoiceQuery
func (s *Service) CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error) {
	body, err := s.createInvoiceQueryResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createInvoiceQueryResponseBody(req CreateInvoiceQueryRequest) (*GetInvoiceQueryResponseBody, error) {
	body := &GetInvoiceQueryResponseBody{}

	response, err := s.connection.Post("/account/v1/invoice-queries", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
