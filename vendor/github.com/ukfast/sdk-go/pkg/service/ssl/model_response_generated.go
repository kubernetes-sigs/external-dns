package ssl

import "github.com/ukfast/sdk-go/pkg/connection"

// GetCertificateSliceResponseBody represents an API response body containing []Certificate data
type GetCertificateSliceResponseBody struct {
	connection.APIResponseBody
	Data []Certificate `json:"data"`
}

// GetCertificateResponseBody represents an API response body containing Certificate data
type GetCertificateResponseBody struct {
	connection.APIResponseBody
	Data Certificate `json:"data"`
}

// GetCertificateContentSliceResponseBody represents an API response body containing []CertificateContent data
type GetCertificateContentSliceResponseBody struct {
	connection.APIResponseBody
	Data []CertificateContent `json:"data"`
}

// GetCertificateContentResponseBody represents an API response body containing CertificateContent data
type GetCertificateContentResponseBody struct {
	connection.APIResponseBody
	Data CertificateContent `json:"data"`
}

// GetCertificatePrivateKeySliceResponseBody represents an API response body containing []CertificatePrivateKey data
type GetCertificatePrivateKeySliceResponseBody struct {
	connection.APIResponseBody
	Data []CertificatePrivateKey `json:"data"`
}

// GetCertificatePrivateKeyResponseBody represents an API response body containing CertificatePrivateKey data
type GetCertificatePrivateKeyResponseBody struct {
	connection.APIResponseBody
	Data CertificatePrivateKey `json:"data"`
}

// GetCertificateValidationSliceResponseBody represents an API response body containing []CertificateValidation data
type GetCertificateValidationSliceResponseBody struct {
	connection.APIResponseBody
	Data []CertificateValidation `json:"data"`
}

// GetCertificateValidationResponseBody represents an API response body containing CertificateValidation data
type GetCertificateValidationResponseBody struct {
	connection.APIResponseBody
	Data CertificateValidation `json:"data"`
}

// GetRecommendationsSliceResponseBody represents an API response body containing []Recommendations data
type GetRecommendationsSliceResponseBody struct {
	connection.APIResponseBody
	Data []Recommendations `json:"data"`
}

// GetRecommendationsResponseBody represents an API response body containing Recommendations data
type GetRecommendationsResponseBody struct {
	connection.APIResponseBody
	Data Recommendations `json:"data"`
}

// GetReportSliceResponseBody represents an API response body containing []Report data
type GetReportSliceResponseBody struct {
	connection.APIResponseBody
	Data []Report `json:"data"`
}

// GetReportResponseBody represents an API response body containing Report data
type GetReportResponseBody struct {
	connection.APIResponseBody
	Data Report `json:"data"`
}
