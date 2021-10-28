package ssl

// ValidateRequest represents a request to validate a certificate
type ValidateRequest struct {
	Key         string `json:"key"`
	Certificate string `json:"certificate"`
	CABundle    string `json:"ca_bundle"`
}
