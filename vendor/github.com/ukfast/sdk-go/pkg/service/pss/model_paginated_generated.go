package pss

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedRequest represents a paginated collection of Request
type PaginatedRequest struct {
	*connection.PaginatedBase
	Items []Request
}

// NewPaginatedRequest returns a pointer to an initialized PaginatedRequest struct
func NewPaginatedRequest(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Request) *PaginatedRequest {
	return &PaginatedRequest{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedReply represents a paginated collection of Reply
type PaginatedReply struct {
	*connection.PaginatedBase
	Items []Reply
}

// NewPaginatedReply returns a pointer to an initialized PaginatedReply struct
func NewPaginatedReply(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Reply) *PaginatedReply {
	return &PaginatedReply{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
