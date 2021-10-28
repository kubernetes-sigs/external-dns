package sharedexchange

import "github.com/ukfast/sdk-go/pkg/connection"

// GetDomainSliceResponseBody represents an API response body containing []Domain data
type GetDomainSliceResponseBody struct {
	connection.APIResponseBody
	Data []Domain `json:"data"`
}

// GetDomainResponseBody represents an API response body containing Domain data
type GetDomainResponseBody struct {
	connection.APIResponseBody
	Data Domain `json:"data"`
}
