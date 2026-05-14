// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// CertificatePackService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificatePackService] method instead.
type CertificatePackService struct {
	Options []option.RequestOption
	Quota   *CertificatePackQuotaService
}

// NewCertificatePackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificatePackService(opts ...option.RequestOption) (r *CertificatePackService) {
	r = &CertificatePackService{}
	r.Options = opts
	r.Quota = NewCertificatePackQuotaService(opts...)
	return
}

// For a given zone, order an advanced certificate pack.
func (r *CertificatePackService) New(ctx context.Context, params CertificatePackNewParams, opts ...option.RequestOption) (res *CertificatePackNewResponse, err error) {
	var env CertificatePackNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, list all active certificate packs.
func (r *CertificatePackService) List(ctx context.Context, params CertificatePackListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CertificatePackListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// For a given zone, list all active certificate packs.
func (r *CertificatePackService) ListAutoPaging(ctx context.Context, params CertificatePackListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CertificatePackListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// For a given zone, delete an advanced certificate pack.
func (r *CertificatePackService) Delete(ctx context.Context, certificatePackID string, body CertificatePackDeleteParams, opts ...option.RequestOption) (res *CertificatePackDeleteResponse, err error) {
	var env CertificatePackDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", body.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, restart validation or add cloudflare branding for an advanced
// certificate pack. The former is only a validation operation for a Certificate
// Pack in a validation_timed_out status.
func (r *CertificatePackService) Edit(ctx context.Context, certificatePackID string, params CertificatePackEditParams, opts ...option.RequestOption) (res *CertificatePackEditResponse, err error) {
	var env CertificatePackEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", params.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// For a given zone, get a certificate pack.
func (r *CertificatePackService) Get(ctx context.Context, certificatePackID string, query CertificatePackGetParams, opts ...option.RequestOption) (res *CertificatePackGetResponse, err error) {
	var env CertificatePackGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", query.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Host = string

type HostParam = string

// The number of days for which the certificate should be valid.
type RequestValidity float64

const (
	RequestValidity7    RequestValidity = 7
	RequestValidity30   RequestValidity = 30
	RequestValidity90   RequestValidity = 90
	RequestValidity365  RequestValidity = 365
	RequestValidity730  RequestValidity = 730
	RequestValidity1095 RequestValidity = 1095
	RequestValidity5475 RequestValidity = 5475
)

func (r RequestValidity) IsKnown() bool {
	switch r {
	case RequestValidity7, RequestValidity30, RequestValidity90, RequestValidity365, RequestValidity730, RequestValidity1095, RequestValidity5475:
		return true
	}
	return false
}

// Status of certificate pack.
type Status string

const (
	StatusInitializing         Status = "initializing"
	StatusPendingValidation    Status = "pending_validation"
	StatusDeleted              Status = "deleted"
	StatusPendingIssuance      Status = "pending_issuance"
	StatusPendingDeployment    Status = "pending_deployment"
	StatusPendingDeletion      Status = "pending_deletion"
	StatusPendingExpiration    Status = "pending_expiration"
	StatusExpired              Status = "expired"
	StatusActive               Status = "active"
	StatusInitializingTimedOut Status = "initializing_timed_out"
	StatusValidationTimedOut   Status = "validation_timed_out"
	StatusIssuanceTimedOut     Status = "issuance_timed_out"
	StatusDeploymentTimedOut   Status = "deployment_timed_out"
	StatusDeletionTimedOut     Status = "deletion_timed_out"
	StatusPendingCleanup       Status = "pending_cleanup"
	StatusStagingDeployment    Status = "staging_deployment"
	StatusStagingActive        Status = "staging_active"
	StatusDeactivating         Status = "deactivating"
	StatusInactive             Status = "inactive"
	StatusBackupIssued         Status = "backup_issued"
	StatusHoldingDeployment    Status = "holding_deployment"
)

func (r Status) IsKnown() bool {
	switch r {
	case StatusInitializing, StatusPendingValidation, StatusDeleted, StatusPendingIssuance, StatusPendingDeployment, StatusPendingDeletion, StatusPendingExpiration, StatusExpired, StatusActive, StatusInitializingTimedOut, StatusValidationTimedOut, StatusIssuanceTimedOut, StatusDeploymentTimedOut, StatusDeletionTimedOut, StatusPendingCleanup, StatusStagingDeployment, StatusStagingActive, StatusDeactivating, StatusInactive, StatusBackupIssued, StatusHoldingDeployment:
		return true
	}
	return false
}

// Validation method in use for a certificate pack order.
type ValidationMethod string

const (
	ValidationMethodHTTP  ValidationMethod = "http"
	ValidationMethodCNAME ValidationMethod = "cname"
	ValidationMethodTXT   ValidationMethod = "txt"
)

func (r ValidationMethod) IsKnown() bool {
	switch r {
	case ValidationMethodHTTP, ValidationMethodCNAME, ValidationMethodTXT:
		return true
	}
	return false
}

type CertificatePackNewResponse struct {
	// Identifier.
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificatePackNewResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []Host `json:"hosts"`
	// Status of certificate pack.
	Status Status `json:"status"`
	// Type of certificate pack.
	Type CertificatePackNewResponseType `json:"type"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors []CertificatePackNewResponseValidationError `json:"validation_errors"`
	// Validation Method selected for the order.
	ValidationMethod CertificatePackNewResponseValidationMethod `json:"validation_method"`
	// Certificates' validation records. Only present when certificate pack is in
	// "pending_validation" status
	ValidationRecords []CertificatePackNewResponseValidationRecord `json:"validation_records"`
	// Validity Days selected for the order.
	ValidityDays CertificatePackNewResponseValidityDays `json:"validity_days"`
	JSON         certificatePackNewResponseJSON         `json:"-"`
}

// certificatePackNewResponseJSON contains the JSON metadata for the struct
// [CertificatePackNewResponse]
type certificatePackNewResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CloudflareBranding   apijson.Field
	Hosts                apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	ValidationErrors     apijson.Field
	ValidationMethod     apijson.Field
	ValidationRecords    apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CertificatePackNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackNewResponseCertificateAuthority string

const (
	CertificatePackNewResponseCertificateAuthorityGoogle      CertificatePackNewResponseCertificateAuthority = "google"
	CertificatePackNewResponseCertificateAuthorityLetsEncrypt CertificatePackNewResponseCertificateAuthority = "lets_encrypt"
	CertificatePackNewResponseCertificateAuthoritySSLCom      CertificatePackNewResponseCertificateAuthority = "ssl_com"
)

func (r CertificatePackNewResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackNewResponseCertificateAuthorityGoogle, CertificatePackNewResponseCertificateAuthorityLetsEncrypt, CertificatePackNewResponseCertificateAuthoritySSLCom:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackNewResponseType string

const (
	CertificatePackNewResponseTypeMhCustom        CertificatePackNewResponseType = "mh_custom"
	CertificatePackNewResponseTypeManagedHostname CertificatePackNewResponseType = "managed_hostname"
	CertificatePackNewResponseTypeSNICustom       CertificatePackNewResponseType = "sni_custom"
	CertificatePackNewResponseTypeUniversal       CertificatePackNewResponseType = "universal"
	CertificatePackNewResponseTypeAdvanced        CertificatePackNewResponseType = "advanced"
	CertificatePackNewResponseTypeTotalTLS        CertificatePackNewResponseType = "total_tls"
	CertificatePackNewResponseTypeKeyless         CertificatePackNewResponseType = "keyless"
	CertificatePackNewResponseTypeLegacyCustom    CertificatePackNewResponseType = "legacy_custom"
)

func (r CertificatePackNewResponseType) IsKnown() bool {
	switch r {
	case CertificatePackNewResponseTypeMhCustom, CertificatePackNewResponseTypeManagedHostname, CertificatePackNewResponseTypeSNICustom, CertificatePackNewResponseTypeUniversal, CertificatePackNewResponseTypeAdvanced, CertificatePackNewResponseTypeTotalTLS, CertificatePackNewResponseTypeKeyless, CertificatePackNewResponseTypeLegacyCustom:
		return true
	}
	return false
}

type CertificatePackNewResponseValidationError struct {
	// A domain validation error.
	Message string                                        `json:"message"`
	JSON    certificatePackNewResponseValidationErrorJSON `json:"-"`
}

// certificatePackNewResponseValidationErrorJSON contains the JSON metadata for the
// struct [CertificatePackNewResponseValidationError]
type certificatePackNewResponseValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackNewResponseValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseValidationErrorJSON) RawJSON() string {
	return r.raw
}

// Validation Method selected for the order.
type CertificatePackNewResponseValidationMethod string

const (
	CertificatePackNewResponseValidationMethodTXT   CertificatePackNewResponseValidationMethod = "txt"
	CertificatePackNewResponseValidationMethodHTTP  CertificatePackNewResponseValidationMethod = "http"
	CertificatePackNewResponseValidationMethodEmail CertificatePackNewResponseValidationMethod = "email"
)

func (r CertificatePackNewResponseValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackNewResponseValidationMethodTXT, CertificatePackNewResponseValidationMethodHTTP, CertificatePackNewResponseValidationMethodEmail:
		return true
	}
	return false
}

// Certificate's required validation record.
type CertificatePackNewResponseValidationRecord struct {
	// The set of email addresses that the certificate authority (CA) will use to
	// complete domain validation.
	Emails []string `json:"emails"`
	// The content that the certificate authority (CA) will expect to find at the
	// http_url during the domain validation.
	HTTPBody string `json:"http_body"`
	// The url that will be checked during domain validation.
	HTTPURL string `json:"http_url"`
	// The hostname that the certificate authority (CA) will check for a TXT record
	// during domain validation .
	TXTName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TXTValue string                                         `json:"txt_value"`
	JSON     certificatePackNewResponseValidationRecordJSON `json:"-"`
}

// certificatePackNewResponseValidationRecordJSON contains the JSON metadata for
// the struct [CertificatePackNewResponseValidationRecord]
type certificatePackNewResponseValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackNewResponseValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseValidationRecordJSON) RawJSON() string {
	return r.raw
}

// Validity Days selected for the order.
type CertificatePackNewResponseValidityDays int64

const (
	CertificatePackNewResponseValidityDays14  CertificatePackNewResponseValidityDays = 14
	CertificatePackNewResponseValidityDays30  CertificatePackNewResponseValidityDays = 30
	CertificatePackNewResponseValidityDays90  CertificatePackNewResponseValidityDays = 90
	CertificatePackNewResponseValidityDays365 CertificatePackNewResponseValidityDays = 365
)

func (r CertificatePackNewResponseValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackNewResponseValidityDays14, CertificatePackNewResponseValidityDays30, CertificatePackNewResponseValidityDays90, CertificatePackNewResponseValidityDays365:
		return true
	}
	return false
}

type CertificatePackListResponse = interface{}

type CertificatePackDeleteResponse struct {
	// Identifier.
	ID   string                            `json:"id"`
	JSON certificatePackDeleteResponseJSON `json:"-"`
}

// certificatePackDeleteResponseJSON contains the JSON metadata for the struct
// [CertificatePackDeleteResponse]
type certificatePackDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponse struct {
	// Identifier.
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificatePackEditResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []Host `json:"hosts"`
	// Status of certificate pack.
	Status Status `json:"status"`
	// Type of certificate pack.
	Type CertificatePackEditResponseType `json:"type"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors []CertificatePackEditResponseValidationError `json:"validation_errors"`
	// Validation Method selected for the order.
	ValidationMethod CertificatePackEditResponseValidationMethod `json:"validation_method"`
	// Certificates' validation records. Only present when certificate pack is in
	// "pending_validation" status
	ValidationRecords []CertificatePackEditResponseValidationRecord `json:"validation_records"`
	// Validity Days selected for the order.
	ValidityDays CertificatePackEditResponseValidityDays `json:"validity_days"`
	JSON         certificatePackEditResponseJSON         `json:"-"`
}

// certificatePackEditResponseJSON contains the JSON metadata for the struct
// [CertificatePackEditResponse]
type certificatePackEditResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CloudflareBranding   apijson.Field
	Hosts                apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	ValidationErrors     apijson.Field
	ValidationMethod     apijson.Field
	ValidationRecords    apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CertificatePackEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackEditResponseCertificateAuthority string

const (
	CertificatePackEditResponseCertificateAuthorityGoogle      CertificatePackEditResponseCertificateAuthority = "google"
	CertificatePackEditResponseCertificateAuthorityLetsEncrypt CertificatePackEditResponseCertificateAuthority = "lets_encrypt"
	CertificatePackEditResponseCertificateAuthoritySSLCom      CertificatePackEditResponseCertificateAuthority = "ssl_com"
)

func (r CertificatePackEditResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseCertificateAuthorityGoogle, CertificatePackEditResponseCertificateAuthorityLetsEncrypt, CertificatePackEditResponseCertificateAuthoritySSLCom:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackEditResponseType string

const (
	CertificatePackEditResponseTypeMhCustom        CertificatePackEditResponseType = "mh_custom"
	CertificatePackEditResponseTypeManagedHostname CertificatePackEditResponseType = "managed_hostname"
	CertificatePackEditResponseTypeSNICustom       CertificatePackEditResponseType = "sni_custom"
	CertificatePackEditResponseTypeUniversal       CertificatePackEditResponseType = "universal"
	CertificatePackEditResponseTypeAdvanced        CertificatePackEditResponseType = "advanced"
	CertificatePackEditResponseTypeTotalTLS        CertificatePackEditResponseType = "total_tls"
	CertificatePackEditResponseTypeKeyless         CertificatePackEditResponseType = "keyless"
	CertificatePackEditResponseTypeLegacyCustom    CertificatePackEditResponseType = "legacy_custom"
)

func (r CertificatePackEditResponseType) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseTypeMhCustom, CertificatePackEditResponseTypeManagedHostname, CertificatePackEditResponseTypeSNICustom, CertificatePackEditResponseTypeUniversal, CertificatePackEditResponseTypeAdvanced, CertificatePackEditResponseTypeTotalTLS, CertificatePackEditResponseTypeKeyless, CertificatePackEditResponseTypeLegacyCustom:
		return true
	}
	return false
}

type CertificatePackEditResponseValidationError struct {
	// A domain validation error.
	Message string                                         `json:"message"`
	JSON    certificatePackEditResponseValidationErrorJSON `json:"-"`
}

// certificatePackEditResponseValidationErrorJSON contains the JSON metadata for
// the struct [CertificatePackEditResponseValidationError]
type certificatePackEditResponseValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseValidationErrorJSON) RawJSON() string {
	return r.raw
}

// Validation Method selected for the order.
type CertificatePackEditResponseValidationMethod string

const (
	CertificatePackEditResponseValidationMethodTXT   CertificatePackEditResponseValidationMethod = "txt"
	CertificatePackEditResponseValidationMethodHTTP  CertificatePackEditResponseValidationMethod = "http"
	CertificatePackEditResponseValidationMethodEmail CertificatePackEditResponseValidationMethod = "email"
)

func (r CertificatePackEditResponseValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseValidationMethodTXT, CertificatePackEditResponseValidationMethodHTTP, CertificatePackEditResponseValidationMethodEmail:
		return true
	}
	return false
}

// Certificate's required validation record.
type CertificatePackEditResponseValidationRecord struct {
	// The set of email addresses that the certificate authority (CA) will use to
	// complete domain validation.
	Emails []string `json:"emails"`
	// The content that the certificate authority (CA) will expect to find at the
	// http_url during the domain validation.
	HTTPBody string `json:"http_body"`
	// The url that will be checked during domain validation.
	HTTPURL string `json:"http_url"`
	// The hostname that the certificate authority (CA) will check for a TXT record
	// during domain validation .
	TXTName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TXTValue string                                          `json:"txt_value"`
	JSON     certificatePackEditResponseValidationRecordJSON `json:"-"`
}

// certificatePackEditResponseValidationRecordJSON contains the JSON metadata for
// the struct [CertificatePackEditResponseValidationRecord]
type certificatePackEditResponseValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseValidationRecordJSON) RawJSON() string {
	return r.raw
}

// Validity Days selected for the order.
type CertificatePackEditResponseValidityDays int64

const (
	CertificatePackEditResponseValidityDays14  CertificatePackEditResponseValidityDays = 14
	CertificatePackEditResponseValidityDays30  CertificatePackEditResponseValidityDays = 30
	CertificatePackEditResponseValidityDays90  CertificatePackEditResponseValidityDays = 90
	CertificatePackEditResponseValidityDays365 CertificatePackEditResponseValidityDays = 365
)

func (r CertificatePackEditResponseValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseValidityDays14, CertificatePackEditResponseValidityDays30, CertificatePackEditResponseValidityDays90, CertificatePackEditResponseValidityDays365:
		return true
	}
	return false
}

type CertificatePackGetResponse = interface{}

type CertificatePackNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[CertificatePackNewParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]HostParam] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[CertificatePackNewParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[CertificatePackNewParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[CertificatePackNewParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r CertificatePackNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackNewParamsCertificateAuthority string

const (
	CertificatePackNewParamsCertificateAuthorityGoogle      CertificatePackNewParamsCertificateAuthority = "google"
	CertificatePackNewParamsCertificateAuthorityLetsEncrypt CertificatePackNewParamsCertificateAuthority = "lets_encrypt"
	CertificatePackNewParamsCertificateAuthoritySSLCom      CertificatePackNewParamsCertificateAuthority = "ssl_com"
)

func (r CertificatePackNewParamsCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackNewParamsCertificateAuthorityGoogle, CertificatePackNewParamsCertificateAuthorityLetsEncrypt, CertificatePackNewParamsCertificateAuthoritySSLCom:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackNewParamsType string

const (
	CertificatePackNewParamsTypeAdvanced CertificatePackNewParamsType = "advanced"
)

func (r CertificatePackNewParamsType) IsKnown() bool {
	switch r {
	case CertificatePackNewParamsTypeAdvanced:
		return true
	}
	return false
}

// Validation Method selected for the order.
type CertificatePackNewParamsValidationMethod string

const (
	CertificatePackNewParamsValidationMethodTXT   CertificatePackNewParamsValidationMethod = "txt"
	CertificatePackNewParamsValidationMethodHTTP  CertificatePackNewParamsValidationMethod = "http"
	CertificatePackNewParamsValidationMethodEmail CertificatePackNewParamsValidationMethod = "email"
)

func (r CertificatePackNewParamsValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackNewParamsValidationMethodTXT, CertificatePackNewParamsValidationMethodHTTP, CertificatePackNewParamsValidationMethodEmail:
		return true
	}
	return false
}

// Validity Days selected for the order.
type CertificatePackNewParamsValidityDays int64

const (
	CertificatePackNewParamsValidityDays14  CertificatePackNewParamsValidityDays = 14
	CertificatePackNewParamsValidityDays30  CertificatePackNewParamsValidityDays = 30
	CertificatePackNewParamsValidityDays90  CertificatePackNewParamsValidityDays = 90
	CertificatePackNewParamsValidityDays365 CertificatePackNewParamsValidityDays = 365
)

func (r CertificatePackNewParamsValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackNewParamsValidityDays14, CertificatePackNewParamsValidityDays30, CertificatePackNewParamsValidityDays90, CertificatePackNewParamsValidityDays365:
		return true
	}
	return false
}

type CertificatePackNewResponseEnvelope struct {
	Errors   []CertificatePackNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificatePackNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackNewResponse                `json:"result"`
	JSON    certificatePackNewResponseEnvelopeJSON    `json:"-"`
}

// certificatePackNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificatePackNewResponseEnvelope]
type certificatePackNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CertificatePackNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             certificatePackNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// certificatePackNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackNewResponseEnvelopeErrors]
type certificatePackNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    certificatePackNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// certificatePackNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CertificatePackNewResponseEnvelopeErrorsSource]
type certificatePackNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CertificatePackNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           CertificatePackNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             certificatePackNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// certificatePackNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackNewResponseEnvelopeMessages]
type certificatePackNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CertificatePackNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    certificatePackNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// certificatePackNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CertificatePackNewResponseEnvelopeMessagesSource]
type certificatePackNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificatePackNewResponseEnvelopeSuccess bool

const (
	CertificatePackNewResponseEnvelopeSuccessTrue CertificatePackNewResponseEnvelopeSuccess = true
)

func (r CertificatePackNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Include Certificate Packs of all statuses, not just active ones.
	Status param.Field[CertificatePackListParamsStatus] `query:"status"`
}

// URLQuery serializes [CertificatePackListParams]'s query parameters as
// `url.Values`.
func (r CertificatePackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type CertificatePackListParamsStatus string

const (
	CertificatePackListParamsStatusAll CertificatePackListParamsStatus = "all"
)

func (r CertificatePackListParamsStatus) IsKnown() bool {
	switch r {
	case CertificatePackListParamsStatusAll:
		return true
	}
	return false
}

type CertificatePackDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackDeleteResponseEnvelope struct {
	Errors   []CertificatePackDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificatePackDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackDeleteResponse                `json:"result"`
	JSON    certificatePackDeleteResponseEnvelopeJSON    `json:"-"`
}

// certificatePackDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackDeleteResponseEnvelope]
type certificatePackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CertificatePackDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             certificatePackDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// certificatePackDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CertificatePackDeleteResponseEnvelopeErrors]
type certificatePackDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    certificatePackDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// certificatePackDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CertificatePackDeleteResponseEnvelopeErrorsSource]
type certificatePackDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           CertificatePackDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             certificatePackDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// certificatePackDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackDeleteResponseEnvelopeMessages]
type certificatePackDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CertificatePackDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    certificatePackDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// certificatePackDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CertificatePackDeleteResponseEnvelopeMessagesSource]
type certificatePackDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificatePackDeleteResponseEnvelopeSuccess bool

const (
	CertificatePackDeleteResponseEnvelopeSuccessTrue CertificatePackDeleteResponseEnvelopeSuccess = true
)

func (r CertificatePackDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r CertificatePackEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificatePackEditResponseEnvelope struct {
	Errors   []CertificatePackEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificatePackEditResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackEditResponse                `json:"result"`
	JSON    certificatePackEditResponseEnvelopeJSON    `json:"-"`
}

// certificatePackEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackEditResponseEnvelope]
type certificatePackEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CertificatePackEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             certificatePackEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// certificatePackEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackEditResponseEnvelopeErrors]
type certificatePackEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    certificatePackEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// certificatePackEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CertificatePackEditResponseEnvelopeErrorsSource]
type certificatePackEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CertificatePackEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             certificatePackEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// certificatePackEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackEditResponseEnvelopeMessages]
type certificatePackEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CertificatePackEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    certificatePackEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// certificatePackEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CertificatePackEditResponseEnvelopeMessagesSource]
type certificatePackEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificatePackEditResponseEnvelopeSuccess bool

const (
	CertificatePackEditResponseEnvelopeSuccessTrue CertificatePackEditResponseEnvelopeSuccess = true
)

func (r CertificatePackEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CertificatePackGetResponseEnvelope struct {
	Errors   []CertificatePackGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificatePackGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackGetResponse                `json:"result"`
	JSON    certificatePackGetResponseEnvelopeJSON    `json:"-"`
}

// certificatePackGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CertificatePackGetResponseEnvelope]
type certificatePackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CertificatePackGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             certificatePackGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// certificatePackGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CertificatePackGetResponseEnvelopeErrors]
type certificatePackGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    certificatePackGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// certificatePackGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CertificatePackGetResponseEnvelopeErrorsSource]
type certificatePackGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           CertificatePackGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             certificatePackGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// certificatePackGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CertificatePackGetResponseEnvelopeMessages]
type certificatePackGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CertificatePackGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    certificatePackGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// certificatePackGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CertificatePackGetResponseEnvelopeMessagesSource]
type certificatePackGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificatePackGetResponseEnvelopeSuccess bool

const (
	CertificatePackGetResponseEnvelopeSuccessTrue CertificatePackGetResponseEnvelopeSuccess = true
)

func (r CertificatePackGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
