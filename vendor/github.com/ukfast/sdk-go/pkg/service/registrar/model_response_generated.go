package registrar

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

// GetNameserverSliceResponseBody represents an API response body containing []Nameserver data
type GetNameserverSliceResponseBody struct {
	connection.APIResponseBody
	Data []Nameserver `json:"data"`
}

// GetNameserverResponseBody represents an API response body containing Nameserver data
type GetNameserverResponseBody struct {
	connection.APIResponseBody
	Data Nameserver `json:"data"`
}

// GetWhoisSliceResponseBody represents an API response body containing []Whois data
type GetWhoisSliceResponseBody struct {
	connection.APIResponseBody
	Data []Whois `json:"data"`
}

// GetWhoisResponseBody represents an API response body containing Whois data
type GetWhoisResponseBody struct {
	connection.APIResponseBody
	Data Whois `json:"data"`
}
