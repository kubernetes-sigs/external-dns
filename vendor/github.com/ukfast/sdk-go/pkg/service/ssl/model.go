//go:generate go run ../../gen/model_response/main.go -package ssl -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package ssl -source model.go -destination model_paginated_generated.go

package ssl

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

type CertificateStatus string

func (s CertificateStatus) String() string {
	return string(s)
}

const (
	CertificateStatusCompleted      CertificateStatus = "Completed"
	CertificateStatusProcessing     CertificateStatus = "Processing"
	CertificateStatusExpired        CertificateStatus = "Expired"
	CertificateStatusExpiring       CertificateStatus = "Expiring"
	CertificateStatusPendingInstall CertificateStatus = "Pending Install"
)

// Certificate represents an SSL certificate
// +genie:model_response
// +genie:model_paginated
type Certificate struct {
	ID               int                 `json:"id"`
	Name             string              `json:"name"`
	Status           CertificateStatus   `json:"status"`
	CommonName       string              `json:"common_name"`
	AlternativeNames []string            `json:"alternative_names"`
	ValidDays        int                 `json:"valid_days"`
	OrderedDate      connection.DateTime `json:"ordered_date"`
	RenewalDate      connection.DateTime `json:"renewal_date"`
}

// CertificateContent represents the content of an SSL certificate
// +genie:model_response
type CertificateContent struct {
	Server       string `json:"server"`
	Intermediate string `json:"intermediate"`
}

// CertificatePrivateKey represents an SSL certificate private key
// +genie:model_response
type CertificatePrivateKey struct {
	Key string `json:"key"`
}

// CertificateValidation represents the results of certificate validation
// +genie:model_response
type CertificateValidation struct {
	Domains   []string            `json:"domains"`
	ExpiresAt connection.DateTime `json:"expires_at"`
}

type RecommendationLevel string

func (s RecommendationLevel) String() string {
	return string(s)
}

const (
	RecommendationLevelLow    RecommendationLevel = "low"
	RecommendationLevelMedium RecommendationLevel = "medium"
	RecommendationLevelHigh   RecommendationLevel = "high"
)

// Recommendations represents SSL recommendations
// +genie:model_response
type Recommendations struct {
	Level    RecommendationLevel `json:"level"`
	Messages []string            `json:"messages"`
}

// Report represents an SSL report
// +genie:model_response
type Report struct {
	Certificate struct {
		Name               string              `json:"name"`
		ValidFrom          connection.DateTime `json:"valid_from"`
		ValidTo            connection.DateTime `json:"valid_to"`
		Issuer             string              `json:"issuer"`
		SerialNumber       string              `json:"serial_number"`
		SignatureAlgorithm string              `json:"signature_algorithm"`
		CoversDomain       bool                `json:"covers_domain"`
		DomainsSecured     []string            `json:"domains_secured"`
		MultiDomain        bool                `json:"multi_domain"`
		Wildcard           bool                `json:"wildcard"`
		Expiring           bool                `json:"expiring"`
		Expired            bool                `json:"expired"`
		SecureSha          bool                `json:"secure_sha"`
	} `json:"certificate"`
	Server struct {
		IP             string              `json:"ip"`
		Hostname       string              `json:"hostname"`
		Port           string              `json:"port"`
		CurrentTime    connection.DateTime `json:"current_time"`
		ServertTime    connection.DateTime `json:"server_time"`
		Software       string              `json:"software"`
		OpenSSLVersion string              `json:"openssl_version"`
		SSLVersions    struct {
			TLS struct {
				_1   bool `json:"1"`
				_1_1 bool `json:"1.1"`
				_1_2 bool `json:"1.2"`
			} `json:"tls"`
			SSL struct {
				_2 bool `json:"2"`
				_3 bool `json:"3"`
			} `json:"ssl"`
		} `json:"ssl_versions"`
	} `json:"server"`
	Vulnerabilities struct {
		Heartbleed bool `json:"heartbleed"`
		Poodle     bool `json:"poodle"`
	} `json:"vulnerabilities"`
	Findings []string `json:"findings"`
	Chain    struct {
		Certificates []struct {
			Name               string              `json:"name"`
			ValidFrom          connection.DateTime `json:"valid_from"`
			ValidTo            connection.DateTime `json:"valid_to"`
			Issuer             string              `json:"issuer"`
			SerialNumber       string              `json:"serial_number"`
			SignatureAlgorithm string              `json:"signature_algorithm"`
			ChainIntact        bool                `json:"chain_intact"`
			CertificateType    string              `json:"certificate_type"`
		} `json:"certificates"`
	} `json:"chain"`
	ChainIntact bool `json:"chain_intact"`
}
