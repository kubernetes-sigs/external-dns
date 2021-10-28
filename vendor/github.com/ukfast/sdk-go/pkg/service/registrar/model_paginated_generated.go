package registrar

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedDomain represents a paginated collection of Domain
type PaginatedDomain struct {
	*connection.PaginatedBase
	Items []Domain
}

// NewPaginatedDomain returns a pointer to an initialized PaginatedDomain struct
func NewPaginatedDomain(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Domain) *PaginatedDomain {
	return &PaginatedDomain{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
