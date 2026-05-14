package redirect

// Certificate represents an NS1 redirect certificate object
type Certificate struct {
	ID          *string `json:"id,omitempty"`
	Domain      string  `json:"domain,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
	ValidFrom   *int64  `json:"valid_from,omitempty"`
	ValidUntil  *int64  `json:"valid_until,omitempty"`
	Processing  *bool   `json:"processing,omitempty"`
	Errors      *string `json:"errors,omitempty"`
	LastUpdated *int64  `json:"last_updated,omitempty"`
}

// CertificateList represents an NS1 redirect certificate list object
type CertificateList struct {
	After   *string        `json:"after,omitempty"`
	Count   int64          `json:"count,omitempty"`
	Limit   *int64         `json:"limit,omitempty"`
	Results []*Certificate `json:"results"`
	Total   int64          `json:"total,omitempty"`
}

// NewCertificate creates a new redirect certificate object with the given domain
func NewCertificate(domain string) *Certificate {
	cert := Certificate{
		Domain: domain,
	}
	return &cert
}
