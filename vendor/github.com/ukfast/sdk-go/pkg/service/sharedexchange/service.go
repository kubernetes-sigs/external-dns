package sharedexchange

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// SharedExchangeService is an interface for managing Shared Exchange
type SharedExchangeService interface {
	GetDomains(parameters connection.APIRequestParameters) ([]Domain, error)
	GetDomainsPaginated(parameters connection.APIRequestParameters) (*PaginatedDomain, error)
	GetDomain(domainID int) (Domain, error)
}

// Service implements SharedExchangeService for managing the Shared Exchange service
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of SharedExchangeService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
