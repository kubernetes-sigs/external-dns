package account

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedContact represents a paginated collection of Contact
type PaginatedContact struct {
	*connection.PaginatedBase
	Items []Contact
}

// NewPaginatedContact returns a pointer to an initialized PaginatedContact struct
func NewPaginatedContact(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Contact) *PaginatedContact {
	return &PaginatedContact{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCredit represents a paginated collection of Credit
type PaginatedCredit struct {
	*connection.PaginatedBase
	Items []Credit
}

// NewPaginatedCredit returns a pointer to an initialized PaginatedCredit struct
func NewPaginatedCredit(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Credit) *PaginatedCredit {
	return &PaginatedCredit{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedInvoice represents a paginated collection of Invoice
type PaginatedInvoice struct {
	*connection.PaginatedBase
	Items []Invoice
}

// NewPaginatedInvoice returns a pointer to an initialized PaginatedInvoice struct
func NewPaginatedInvoice(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Invoice) *PaginatedInvoice {
	return &PaginatedInvoice{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedInvoiceQuery represents a paginated collection of InvoiceQuery
type PaginatedInvoiceQuery struct {
	*connection.PaginatedBase
	Items []InvoiceQuery
}

// NewPaginatedInvoiceQuery returns a pointer to an initialized PaginatedInvoiceQuery struct
func NewPaginatedInvoiceQuery(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []InvoiceQuery) *PaginatedInvoiceQuery {
	return &PaginatedInvoiceQuery{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
