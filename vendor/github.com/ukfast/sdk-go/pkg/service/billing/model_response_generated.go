package billing

import "github.com/ukfast/sdk-go/pkg/connection"

// GetCardSliceResponseBody represents an API response body containing []Card data
type GetCardSliceResponseBody struct {
	connection.APIResponseBody
	Data []Card `json:"data"`
}

// GetCardResponseBody represents an API response body containing Card data
type GetCardResponseBody struct {
	connection.APIResponseBody
	Data Card `json:"data"`
}

// GetCloudCostSliceResponseBody represents an API response body containing []CloudCost data
type GetCloudCostSliceResponseBody struct {
	connection.APIResponseBody
	Data []CloudCost `json:"data"`
}

// GetCloudCostResponseBody represents an API response body containing CloudCost data
type GetCloudCostResponseBody struct {
	connection.APIResponseBody
	Data CloudCost `json:"data"`
}

// GetDirectDebitSliceResponseBody represents an API response body containing []DirectDebit data
type GetDirectDebitSliceResponseBody struct {
	connection.APIResponseBody
	Data []DirectDebit `json:"data"`
}

// GetDirectDebitResponseBody represents an API response body containing DirectDebit data
type GetDirectDebitResponseBody struct {
	connection.APIResponseBody
	Data DirectDebit `json:"data"`
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

// GetPaymentSliceResponseBody represents an API response body containing []Payment data
type GetPaymentSliceResponseBody struct {
	connection.APIResponseBody
	Data []Payment `json:"data"`
}

// GetPaymentResponseBody represents an API response body containing Payment data
type GetPaymentResponseBody struct {
	connection.APIResponseBody
	Data Payment `json:"data"`
}

// GetRecurringCostSliceResponseBody represents an API response body containing []RecurringCost data
type GetRecurringCostSliceResponseBody struct {
	connection.APIResponseBody
	Data []RecurringCost `json:"data"`
}

// GetRecurringCostResponseBody represents an API response body containing RecurringCost data
type GetRecurringCostResponseBody struct {
	connection.APIResponseBody
	Data RecurringCost `json:"data"`
}
