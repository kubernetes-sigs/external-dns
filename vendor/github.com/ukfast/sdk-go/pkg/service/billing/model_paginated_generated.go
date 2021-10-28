package billing

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedCard represents a paginated collection of Card
type PaginatedCard struct {
	*connection.PaginatedBase
	Items []Card
}

// NewPaginatedCard returns a pointer to an initialized PaginatedCard struct
func NewPaginatedCard(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Card) *PaginatedCard {
	return &PaginatedCard{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCloudCost represents a paginated collection of CloudCost
type PaginatedCloudCost struct {
	*connection.PaginatedBase
	Items []CloudCost
}

// NewPaginatedCloudCost returns a pointer to an initialized PaginatedCloudCost struct
func NewPaginatedCloudCost(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []CloudCost) *PaginatedCloudCost {
	return &PaginatedCloudCost{
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

// PaginatedPayment represents a paginated collection of Payment
type PaginatedPayment struct {
	*connection.PaginatedBase
	Items []Payment
}

// NewPaginatedPayment returns a pointer to an initialized PaginatedPayment struct
func NewPaginatedPayment(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Payment) *PaginatedPayment {
	return &PaginatedPayment{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRecurringCost represents a paginated collection of RecurringCost
type PaginatedRecurringCost struct {
	*connection.PaginatedBase
	Items []RecurringCost
}

// NewPaginatedRecurringCost returns a pointer to an initialized PaginatedRecurringCost struct
func NewPaginatedRecurringCost(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []RecurringCost) *PaginatedRecurringCost {
	return &PaginatedRecurringCost{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
