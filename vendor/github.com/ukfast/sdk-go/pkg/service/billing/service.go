package billing

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// BillingService is an interface for managing billing
type BillingService interface {
	GetCards(parameters connection.APIRequestParameters) ([]Card, error)
	GetCardsPaginated(parameters connection.APIRequestParameters) (*PaginatedCard, error)
	GetCard(cardID int) (Card, error)
	CreateCard(req CreateCardRequest) (int, error)
	PatchCard(cardID int, patch PatchCardRequest) error
	DeleteCard(cardID int) error

	GetCloudCosts(parameters connection.APIRequestParameters) ([]CloudCost, error)
	GetCloudCostsPaginated(parameters connection.APIRequestParameters) (*PaginatedCloudCost, error)
	GetCloudCost(costID int) (CloudCost, error)

	GetDirectDebit() (DirectDebit, error)

	GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error)
	GetInvoicesPaginated(parameters connection.APIRequestParameters) (*PaginatedInvoice, error)
	GetInvoice(invoiceID int) (Invoice, error)

	GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error)
	GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*PaginatedInvoiceQuery, error)
	GetInvoiceQuery(queryID int) (InvoiceQuery, error)
	CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error)

	GetPayments(parameters connection.APIRequestParameters) ([]Payment, error)
	GetPaymentsPaginated(parameters connection.APIRequestParameters) (*PaginatedPayment, error)
	GetPayment(paymentID int) (Payment, error)

	GetRecurringCosts(parameters connection.APIRequestParameters) ([]RecurringCost, error)
	GetRecurringCostsPaginated(parameters connection.APIRequestParameters) (*PaginatedRecurringCost, error)
	GetRecurringCost(costID int) (RecurringCost, error)
}

// Service implements BillingService for managing
// Billing certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of BillingService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
