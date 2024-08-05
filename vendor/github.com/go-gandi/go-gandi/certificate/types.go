package certificate

import (
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
)

type Certificate struct {
	client client.Gandi
}

type Package struct {
	Name       string `json:"name"`
	NameLabel  string `json:"name_label"`
	Href       string `json:"href"`
	MaxDomains int    `json:"max_domains"`
	Type       string `json:"type"`
	TypeLabel  string `json:"type_label"`
	Wildcard   bool   `json:"wildcard"`
}

type CertificateType struct {
	ID              string              `json:"id"`
	CN              string              `json:"cn"`
	CNUnicode       string              `json:"cn_unicode"`
	AltNames        []string            `json:"altnames"`
	AltNamesUnicode []string            `json:"altnames_unicode"`
	Contact         *CertificateContact `json:"contact"`
	Dates           *CertificateDates   `json:"dates"`
	Package         *Package            `json:"package"`
	Software        int                 `json:"software"`
	Status          string              `json:"status"`
}

type CertificateDates struct {
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	EndsAt            *time.Time `json:"ends_at,omitempty"`
	StartedAt         *time.Time `json:"started_at,omitempty"`
	SubcriptionEndsAt *time.Time `json:"subscription_ends_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

type CertificateContact struct {
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Email      string `json:"email,omitempty"`
	Family     string `json:"family,omitempty"`
	Given      string `json:"given,omitempty"`
	OrgName    string `json:"orgname,omitempty"`
	State      string `json:"state,omitempty"`
	StreetAddr string `json:"street_addr,omitempty"`
	Zip        string `json:"zip,omitempty"`
}

type CreateCertificateRequest struct {
	CN      string `json:"cn"`
	Package string `json:"package"`
}

type CreateCertificateResponse struct {
	Href    string `json:"href"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Cause   string `json:"cause,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Object  string `json:"object,omitempty"`
}
