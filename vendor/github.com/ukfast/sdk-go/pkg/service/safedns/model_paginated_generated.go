package safedns

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedZone represents a paginated collection of Zone
type PaginatedZone struct {
	*connection.PaginatedBase
	Items []Zone
}

// NewPaginatedZone returns a pointer to an initialized PaginatedZone struct
func NewPaginatedZone(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Zone) *PaginatedZone {
	return &PaginatedZone{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRecord represents a paginated collection of Record
type PaginatedRecord struct {
	*connection.PaginatedBase
	Items []Record
}

// NewPaginatedRecord returns a pointer to an initialized PaginatedRecord struct
func NewPaginatedRecord(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Record) *PaginatedRecord {
	return &PaginatedRecord{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNote represents a paginated collection of Note
type PaginatedNote struct {
	*connection.PaginatedBase
	Items []Note
}

// NewPaginatedNote returns a pointer to an initialized PaginatedNote struct
func NewPaginatedNote(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Note) *PaginatedNote {
	return &PaginatedNote{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTemplate represents a paginated collection of Template
type PaginatedTemplate struct {
	*connection.PaginatedBase
	Items []Template
}

// NewPaginatedTemplate returns a pointer to an initialized PaginatedTemplate struct
func NewPaginatedTemplate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Template) *PaginatedTemplate {
	return &PaginatedTemplate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
