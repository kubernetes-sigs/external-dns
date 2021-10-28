package ssl

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedCertificate represents a paginated collection of Certificate
type PaginatedCertificate struct {
	*connection.PaginatedBase
	Items []Certificate
}

// NewPaginatedCertificate returns a pointer to an initialized PaginatedCertificate struct
func NewPaginatedCertificate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Certificate) *PaginatedCertificate {
	return &PaginatedCertificate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
