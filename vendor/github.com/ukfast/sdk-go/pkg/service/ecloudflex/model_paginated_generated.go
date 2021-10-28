package ecloudflex

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedProject represents a paginated collection of Project
type PaginatedProject struct {
	*connection.PaginatedBase
	Items []Project
}

// NewPaginatedProject returns a pointer to an initialized PaginatedProject struct
func NewPaginatedProject(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Project) *PaginatedProject {
	return &PaginatedProject{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
