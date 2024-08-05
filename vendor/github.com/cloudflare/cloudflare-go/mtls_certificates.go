package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// MTLSAssociation represents the metadata for an existing association
// between a user-uploaded mTLS certificate and a Cloudflare service.
type MTLSAssociation struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// MTLSAssociationResponse represents the response from the retrieval endpoint
// for mTLS certificate associations.
type MTLSAssociationResponse struct {
	Response
	Result []MTLSAssociation `json:"result"`
}

// MTLSCertificate represents the metadata for a user-uploaded mTLS
// certificate.
type MTLSCertificate struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Issuer       string    `json:"issuer"`
	Signature    string    `json:"signature"`
	SerialNumber string    `json:"serial_number"`
	Certificates string    `json:"certificates"`
	CA           bool      `json:"ca"`
	UploadedOn   time.Time `json:"uploaded_on"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExpiresOn    time.Time `json:"expires_on"`
}

// MTLSCertificateResponse represents the response from endpoints relating to
// retrieving, creating, and deleting an mTLS certificate.
type MTLSCertificateResponse struct {
	Response
	Result MTLSCertificate `json:"result"`
}

// MTLSCertificatesResponse represents the response from the mTLS certificate
// list endpoint.
type MTLSCertificatesResponse struct {
	Response
	Result     []MTLSCertificate `json:"result"`
	ResultInfo `json:"result_info"`
}

// MTLSCertificateParams represents the data related to the mTLS certificate
// being uploaded. Name is an optional field.
type CreateMTLSCertificateParams struct {
	Name         string `json:"name"`
	Certificates string `json:"certificates"`
	PrivateKey   string `json:"private_key"`
	CA           bool   `json:"ca"`
}

type ListMTLSCertificatesParams struct {
	PaginationOptions
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Name   string `url:"name,omitempty"`
	CA     bool   `url:"ca,omitempty"`
}

type ListMTLSCertificateAssociationsParams struct {
	CertificateID string
}

var (
	ErrMissingCertificateID = errors.New("missing required certificate ID")
)

// ListMTLSCertificates returns a list of all user-uploaded mTLS certificates.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-list-mtls-certificates
func (api *API) ListMTLSCertificates(ctx context.Context, rc *ResourceContainer, params ListMTLSCertificatesParams) ([]MTLSCertificate, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []MTLSCertificate{}, ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []MTLSCertificate{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return []MTLSCertificate{}, ResultInfo{}, err
	}
	var r MTLSCertificatesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []MTLSCertificate{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, err
}

// GetMTLSCertificate returns the metadata associated with a user-uploaded mTLS
// certificate.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-get-mtls-certificate
func (api *API) GetMTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) (MTLSCertificate, error) {
	if rc.Level != AccountRouteLevel {
		return MTLSCertificate{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return MTLSCertificate{}, ErrMissingAccountID
	}

	if certificateID == "" {
		return MTLSCertificate{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", rc.Identifier, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return MTLSCertificate{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListMTLSCertificateAssociations returns a list of all existing associations
// between the mTLS certificate and Cloudflare services.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-list-mtls-certificate-associations
func (api *API) ListMTLSCertificateAssociations(ctx context.Context, rc *ResourceContainer, params ListMTLSCertificateAssociationsParams) ([]MTLSAssociation, error) {
	if rc.Level != AccountRouteLevel {
		return []MTLSAssociation{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []MTLSAssociation{}, ErrMissingAccountID
	}

	if params.CertificateID == "" {
		return []MTLSAssociation{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s/associations", rc.Identifier, params.CertificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MTLSAssociation{}, err
	}
	var r MTLSAssociationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []MTLSAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateMTLSCertificate will create the provided certificate for use with mTLS
// enabled Cloudflare services.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-upload-mtls-certificate
func (api *API) CreateMTLSCertificate(ctx context.Context, rc *ResourceContainer, params CreateMTLSCertificateParams) (MTLSCertificate, error) {
	if rc.Level != AccountRouteLevel {
		return MTLSCertificate{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return MTLSCertificate{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return MTLSCertificate{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteMTLSCertificate will delete the specified mTLS certificate.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-delete-mtls-certificate
func (api *API) DeleteMTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) (MTLSCertificate, error) {
	if rc.Level != AccountRouteLevel {
		return MTLSCertificate{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return MTLSCertificate{}, ErrMissingAccountID
	}

	if certificateID == "" {
		return MTLSCertificate{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", rc.Identifier, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return MTLSCertificate{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
