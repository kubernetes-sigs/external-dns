package account

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// AccountService is an interface for managing account
type AccountService interface {
	GetContacts(parameters connection.APIRequestParameters) ([]Contact, error)
	GetContactsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Contact], error)
	GetContact(contactID int) (Contact, error)

	GetDetails() (Details, error)

	GetCredits(parameters connection.APIRequestParameters) ([]Credit, error)

	GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error)
	GetInvoicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Invoice], error)
	GetInvoice(invoiceID int) (Invoice, error)

	GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error)
	GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error)
	GetInvoiceQuery(invoiceQueryID int) (InvoiceQuery, error)
	CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error)
}

// Service implements AccountService for managing
// Account certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of AccountService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
