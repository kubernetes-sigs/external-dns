package account

import "github.com/ukfast/sdk-go/pkg/connection"

// GetContactSliceResponseBody represents an API response body containing []Contact data
type GetContactSliceResponseBody struct {
	connection.APIResponseBody
	Data []Contact `json:"data"`
}

// GetContactResponseBody represents an API response body containing Contact data
type GetContactResponseBody struct {
	connection.APIResponseBody
	Data Contact `json:"data"`
}

// GetDetailsSliceResponseBody represents an API response body containing []Details data
type GetDetailsSliceResponseBody struct {
	connection.APIResponseBody
	Data []Details `json:"data"`
}

// GetDetailsResponseBody represents an API response body containing Details data
type GetDetailsResponseBody struct {
	connection.APIResponseBody
	Data Details `json:"data"`
}

// GetCreditSliceResponseBody represents an API response body containing []Credit data
type GetCreditSliceResponseBody struct {
	connection.APIResponseBody
	Data []Credit `json:"data"`
}

// GetCreditResponseBody represents an API response body containing Credit data
type GetCreditResponseBody struct {
	connection.APIResponseBody
	Data Credit `json:"data"`
}

// GetInvoiceSliceResponseBody represents an API response body containing []Invoice data
type GetInvoiceSliceResponseBody struct {
	connection.APIResponseBody
	Data []Invoice `json:"data"`
}

// GetInvoiceResponseBody represents an API response body containing Invoice data
type GetInvoiceResponseBody struct {
	connection.APIResponseBody
	Data Invoice `json:"data"`
}

// GetInvoiceQuerySliceResponseBody represents an API response body containing []InvoiceQuery data
type GetInvoiceQuerySliceResponseBody struct {
	connection.APIResponseBody
	Data []InvoiceQuery `json:"data"`
}

// GetInvoiceQueryResponseBody represents an API response body containing InvoiceQuery data
type GetInvoiceQueryResponseBody struct {
	connection.APIResponseBody
	Data InvoiceQuery `json:"data"`
}
