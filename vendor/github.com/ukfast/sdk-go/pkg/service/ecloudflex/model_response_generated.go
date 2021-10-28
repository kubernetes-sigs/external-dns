package ecloudflex

import "github.com/ukfast/sdk-go/pkg/connection"

// GetProjectSliceResponseBody represents an API response body containing []Project data
type GetProjectSliceResponseBody struct {
	connection.APIResponseBody
	Data []Project `json:"data"`
}

// GetProjectResponseBody represents an API response body containing Project data
type GetProjectResponseBody struct {
	connection.APIResponseBody
	Data Project `json:"data"`
}
