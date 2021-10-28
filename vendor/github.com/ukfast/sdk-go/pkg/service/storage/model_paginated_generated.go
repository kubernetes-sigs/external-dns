package storage

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedSolution represents a paginated collection of Solution
type PaginatedSolution struct {
	*connection.PaginatedBase
	Items []Solution
}

// NewPaginatedSolution returns a pointer to an initialized PaginatedSolution struct
func NewPaginatedSolution(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Solution) *PaginatedSolution {
	return &PaginatedSolution{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHost represents a paginated collection of Host
type PaginatedHost struct {
	*connection.PaginatedBase
	Items []Host
}

// NewPaginatedHost returns a pointer to an initialized PaginatedHost struct
func NewPaginatedHost(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Host) *PaginatedHost {
	return &PaginatedHost{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVolume represents a paginated collection of Volume
type PaginatedVolume struct {
	*connection.PaginatedBase
	Items []Volume
}

// NewPaginatedVolume returns a pointer to an initialized PaginatedVolume struct
func NewPaginatedVolume(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Volume) *PaginatedVolume {
	return &PaginatedVolume{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
