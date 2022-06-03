package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetInvoiceQueries retrieves a list of invoice queries
func (s *Service) GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error) {
	return connection.InvokeRequestAll(s.GetInvoiceQueriesPaginated, parameters)
}

// GetInvoiceQueriesPaginated retrieves a paginated list of invoice queries
func (s *Service) GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error) {
	body, err := s.getInvoiceQueriesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoiceQueriesPaginated), err
}

func (s *Service) getInvoiceQueriesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]InvoiceQuery], error) {
	body := &connection.APIResponseBodyData[[]InvoiceQuery]{}

	response, err := s.connection.Get("/billing/v1/invoice-queries", parameters)
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

func (s *Service) getInvoiceQueryResponseBody(queryID int) (*connection.APIResponseBodyData[InvoiceQuery], error) {
	body := &connection.APIResponseBodyData[InvoiceQuery]{}

	if queryID < 1 {
		return body, fmt.Errorf("invalid invoice query id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/invoice-queries/%d", queryID), connection.APIRequestParameters{})
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

func (s *Service) createInvoiceQueryResponseBody(req CreateInvoiceQueryRequest) (*connection.APIResponseBodyData[InvoiceQuery], error) {
	body := &connection.APIResponseBodyData[InvoiceQuery]{}

	response, err := s.connection.Post("/billing/v1/invoice-queries", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
