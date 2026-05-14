// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// CustomHostnameService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomHostnameService] method instead.
type CustomHostnameService struct {
	Options         []option.RequestOption
	FallbackOrigin  *FallbackOriginService
	CertificatePack *CertificatePackService
}

// NewCustomHostnameService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomHostnameService(opts ...option.RequestOption) (r *CustomHostnameService) {
	r = &CustomHostnameService{}
	r.Options = opts
	r.FallbackOrigin = NewFallbackOriginService(opts...)
	r.CertificatePack = NewCertificatePackService(opts...)
	return
}

// Add a new custom hostname and request that an SSL certificate be issued for it.
// One of three validation methods—http, txt, email—should be used, with 'http'
// recommended if the CNAME is already in place (or will be soon). Specifying
// 'email' will send an email to the WHOIS contacts on file for the base domain
// plus hostmaster, postmaster, webmaster, admin, administrator. If http is used
// and the domain is not already pointing to the Managed CNAME host, the PATCH
// method must be used once it is (to complete validation). Enable bundling of
// certificates using the custom_cert_bundle field. The bundling process requires
// the following condition One certificate in the bundle must use an RSA, and the
// other must use an ECDSA.
func (r *CustomHostnameService) New(ctx context.Context, params CustomHostnameNewParams, opts ...option.RequestOption) (res *CustomHostnameNewResponse, err error) {
	var env CustomHostnameNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) List(ctx context.Context, params CustomHostnameListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomHostnameListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames", params.ZoneID)
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

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) ListAutoPaging(ctx context.Context, params CustomHostnameListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomHostnameListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Custom Hostname (and any issued SSL certificates)
func (r *CustomHostnameService) Delete(ctx context.Context, customHostnameID string, body CustomHostnameDeleteParams, opts ...option.RequestOption) (res *CustomHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", body.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'. Bundle an existing certificate with another certificate by
// using the "custom_cert_bundle" field. The bundling process supports combining
// certificates as long as the following condition is met. One certificate must use
// the RSA algorithm, and the other must use the ECDSA algorithm.
func (r *CustomHostnameService) Edit(ctx context.Context, customHostnameID string, params CustomHostnameEditParams, opts ...option.RequestOption) (res *CustomHostnameEditResponse, err error) {
	var env CustomHostnameEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", params.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Custom Hostname Details
func (r *CustomHostnameService) Get(ctx context.Context, customHostnameID string, query CustomHostnameGetParams, opts ...option.RequestOption) (res *CustomHostnameGetResponse, err error) {
	var env CustomHostnameGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", query.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type BundleMethod string

const (
	BundleMethodUbiquitous BundleMethod = "ubiquitous"
	BundleMethodOptimal    BundleMethod = "optimal"
	BundleMethodForce      BundleMethod = "force"
)

func (r BundleMethod) IsKnown() bool {
	switch r {
	case BundleMethodUbiquitous, BundleMethodOptimal, BundleMethodForce:
		return true
	}
	return false
}

// Domain control validation (DCV) method used for this hostname.
type DCVMethod string

const (
	DCVMethodHTTP  DCVMethod = "http"
	DCVMethodTXT   DCVMethod = "txt"
	DCVMethodEmail DCVMethod = "email"
)

func (r DCVMethod) IsKnown() bool {
	switch r {
	case DCVMethodHTTP, DCVMethodTXT, DCVMethodEmail:
		return true
	}
	return false
}

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type DomainValidationType string

const (
	DomainValidationTypeDv DomainValidationType = "dv"
)

func (r DomainValidationType) IsKnown() bool {
	switch r {
	case DomainValidationTypeDv:
		return true
	}
	return false
}

type CustomHostnameNewResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string                       `json:"hostname,required"`
	SSL      CustomHostnameNewResponseSSL `json:"ssl,required"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSNI string `json:"custom_origin_sni"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameNewResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameNewResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CustomHostnameNewResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string                      `json:"verification_errors"`
	JSON               customHostnameNewResponseJSON `json:"-"`
}

// customHostnameNewResponseJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponse]
type customHostnameNewResponseJSON struct {
	ID                        apijson.Field
	Hostname                  apijson.Field
	SSL                       apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSNI           apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod BundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority shared.CertificateCA `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []string `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method DCVMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string                               `json:"serial_number"`
	Settings     CustomHostnameNewResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameNewResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type DomainValidationType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameNewResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameNewResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                             `json:"wildcard"`
	JSON     customHostnameNewResponseSSLJSON `json:"-"`
}

// customHostnameNewResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponseSSL]
type customHostnameNewResponseSSLJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseSSLJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameNewResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameNewResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameNewResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameNewResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameNewResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameNewResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseSSLSettings]
type customHostnameNewResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseSSLSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CustomHostnameNewResponseSSLSettingsEarlyHints string

const (
	CustomHostnameNewResponseSSLSettingsEarlyHintsOn  CustomHostnameNewResponseSSLSettingsEarlyHints = "on"
	CustomHostnameNewResponseSSLSettingsEarlyHintsOff CustomHostnameNewResponseSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameNewResponseSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseSSLSettingsEarlyHintsOn, CustomHostnameNewResponseSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameNewResponseSSLSettingsHTTP2 string

const (
	CustomHostnameNewResponseSSLSettingsHTTP2On  CustomHostnameNewResponseSSLSettingsHTTP2 = "on"
	CustomHostnameNewResponseSSLSettingsHTTP2Off CustomHostnameNewResponseSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameNewResponseSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseSSLSettingsHTTP2On, CustomHostnameNewResponseSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameNewResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_0 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_1 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_2 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_3 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameNewResponseSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseSSLSettingsMinTLSVersion1_0, CustomHostnameNewResponseSSLSettingsMinTLSVersion1_1, CustomHostnameNewResponseSSLSettingsMinTLSVersion1_2, CustomHostnameNewResponseSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameNewResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameNewResponseSSLSettingsTLS1_3On  CustomHostnameNewResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameNewResponseSSLSettingsTLS1_3Off CustomHostnameNewResponseSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameNewResponseSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseSSLSettingsTLS1_3On, CustomHostnameNewResponseSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CustomHostnameNewResponseSSLStatus string

const (
	CustomHostnameNewResponseSSLStatusInitializing         CustomHostnameNewResponseSSLStatus = "initializing"
	CustomHostnameNewResponseSSLStatusPendingValidation    CustomHostnameNewResponseSSLStatus = "pending_validation"
	CustomHostnameNewResponseSSLStatusDeleted              CustomHostnameNewResponseSSLStatus = "deleted"
	CustomHostnameNewResponseSSLStatusPendingIssuance      CustomHostnameNewResponseSSLStatus = "pending_issuance"
	CustomHostnameNewResponseSSLStatusPendingDeployment    CustomHostnameNewResponseSSLStatus = "pending_deployment"
	CustomHostnameNewResponseSSLStatusPendingDeletion      CustomHostnameNewResponseSSLStatus = "pending_deletion"
	CustomHostnameNewResponseSSLStatusPendingExpiration    CustomHostnameNewResponseSSLStatus = "pending_expiration"
	CustomHostnameNewResponseSSLStatusExpired              CustomHostnameNewResponseSSLStatus = "expired"
	CustomHostnameNewResponseSSLStatusActive               CustomHostnameNewResponseSSLStatus = "active"
	CustomHostnameNewResponseSSLStatusInitializingTimedOut CustomHostnameNewResponseSSLStatus = "initializing_timed_out"
	CustomHostnameNewResponseSSLStatusValidationTimedOut   CustomHostnameNewResponseSSLStatus = "validation_timed_out"
	CustomHostnameNewResponseSSLStatusIssuanceTimedOut     CustomHostnameNewResponseSSLStatus = "issuance_timed_out"
	CustomHostnameNewResponseSSLStatusDeploymentTimedOut   CustomHostnameNewResponseSSLStatus = "deployment_timed_out"
	CustomHostnameNewResponseSSLStatusDeletionTimedOut     CustomHostnameNewResponseSSLStatus = "deletion_timed_out"
	CustomHostnameNewResponseSSLStatusPendingCleanup       CustomHostnameNewResponseSSLStatus = "pending_cleanup"
	CustomHostnameNewResponseSSLStatusStagingDeployment    CustomHostnameNewResponseSSLStatus = "staging_deployment"
	CustomHostnameNewResponseSSLStatusStagingActive        CustomHostnameNewResponseSSLStatus = "staging_active"
	CustomHostnameNewResponseSSLStatusDeactivating         CustomHostnameNewResponseSSLStatus = "deactivating"
	CustomHostnameNewResponseSSLStatusInactive             CustomHostnameNewResponseSSLStatus = "inactive"
	CustomHostnameNewResponseSSLStatusBackupIssued         CustomHostnameNewResponseSSLStatus = "backup_issued"
	CustomHostnameNewResponseSSLStatusHoldingDeployment    CustomHostnameNewResponseSSLStatus = "holding_deployment"
)

func (r CustomHostnameNewResponseSSLStatus) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseSSLStatusInitializing, CustomHostnameNewResponseSSLStatusPendingValidation, CustomHostnameNewResponseSSLStatusDeleted, CustomHostnameNewResponseSSLStatusPendingIssuance, CustomHostnameNewResponseSSLStatusPendingDeployment, CustomHostnameNewResponseSSLStatusPendingDeletion, CustomHostnameNewResponseSSLStatusPendingExpiration, CustomHostnameNewResponseSSLStatusExpired, CustomHostnameNewResponseSSLStatusActive, CustomHostnameNewResponseSSLStatusInitializingTimedOut, CustomHostnameNewResponseSSLStatusValidationTimedOut, CustomHostnameNewResponseSSLStatusIssuanceTimedOut, CustomHostnameNewResponseSSLStatusDeploymentTimedOut, CustomHostnameNewResponseSSLStatusDeletionTimedOut, CustomHostnameNewResponseSSLStatusPendingCleanup, CustomHostnameNewResponseSSLStatusStagingDeployment, CustomHostnameNewResponseSSLStatusStagingActive, CustomHostnameNewResponseSSLStatusDeactivating, CustomHostnameNewResponseSSLStatusInactive, CustomHostnameNewResponseSSLStatusBackupIssued, CustomHostnameNewResponseSSLStatusHoldingDeployment:
		return true
	}
	return false
}

type CustomHostnameNewResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                          `json:"message"`
	JSON    customHostnameNewResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameNewResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameNewResponseSSLValidationError]
type customHostnameNewResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseSSLValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseSSLValidationRecord struct {
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
	TXTValue string                                           `json:"txt_value"`
	JSON     customHostnameNewResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameNewResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameNewResponseSSLValidationRecord]
type customHostnameNewResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseSSLValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameNewResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameNewResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                             `json:"value"`
	JSON  customHostnameNewResponseOwnershipVerificationJSON `json:"-"`
}

// customHostnameNewResponseOwnershipVerificationJSON contains the JSON metadata
// for the struct [CustomHostnameNewResponseOwnershipVerification]
type customHostnameNewResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CustomHostnameNewResponseOwnershipVerificationType string

const (
	CustomHostnameNewResponseOwnershipVerificationTypeTXT CustomHostnameNewResponseOwnershipVerificationType = "txt"
)

func (r CustomHostnameNewResponseOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseOwnershipVerificationTypeTXT:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameNewResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                 `json:"http_url"`
	JSON    customHostnameNewResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameNewResponseOwnershipVerificationHTTPJSON contains the JSON
// metadata for the struct [CustomHostnameNewResponseOwnershipVerificationHTTP]
type customHostnameNewResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CustomHostnameNewResponseStatus string

const (
	CustomHostnameNewResponseStatusActive             CustomHostnameNewResponseStatus = "active"
	CustomHostnameNewResponseStatusPending            CustomHostnameNewResponseStatus = "pending"
	CustomHostnameNewResponseStatusActiveRedeploying  CustomHostnameNewResponseStatus = "active_redeploying"
	CustomHostnameNewResponseStatusMoved              CustomHostnameNewResponseStatus = "moved"
	CustomHostnameNewResponseStatusPendingDeletion    CustomHostnameNewResponseStatus = "pending_deletion"
	CustomHostnameNewResponseStatusDeleted            CustomHostnameNewResponseStatus = "deleted"
	CustomHostnameNewResponseStatusPendingBlocked     CustomHostnameNewResponseStatus = "pending_blocked"
	CustomHostnameNewResponseStatusPendingMigration   CustomHostnameNewResponseStatus = "pending_migration"
	CustomHostnameNewResponseStatusPendingProvisioned CustomHostnameNewResponseStatus = "pending_provisioned"
	CustomHostnameNewResponseStatusTestPending        CustomHostnameNewResponseStatus = "test_pending"
	CustomHostnameNewResponseStatusTestActive         CustomHostnameNewResponseStatus = "test_active"
	CustomHostnameNewResponseStatusTestActiveApex     CustomHostnameNewResponseStatus = "test_active_apex"
	CustomHostnameNewResponseStatusTestBlocked        CustomHostnameNewResponseStatus = "test_blocked"
	CustomHostnameNewResponseStatusTestFailed         CustomHostnameNewResponseStatus = "test_failed"
	CustomHostnameNewResponseStatusProvisioned        CustomHostnameNewResponseStatus = "provisioned"
	CustomHostnameNewResponseStatusBlocked            CustomHostnameNewResponseStatus = "blocked"
)

func (r CustomHostnameNewResponseStatus) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseStatusActive, CustomHostnameNewResponseStatusPending, CustomHostnameNewResponseStatusActiveRedeploying, CustomHostnameNewResponseStatusMoved, CustomHostnameNewResponseStatusPendingDeletion, CustomHostnameNewResponseStatusDeleted, CustomHostnameNewResponseStatusPendingBlocked, CustomHostnameNewResponseStatusPendingMigration, CustomHostnameNewResponseStatusPendingProvisioned, CustomHostnameNewResponseStatusTestPending, CustomHostnameNewResponseStatusTestActive, CustomHostnameNewResponseStatusTestActiveApex, CustomHostnameNewResponseStatusTestBlocked, CustomHostnameNewResponseStatusTestFailed, CustomHostnameNewResponseStatusProvisioned, CustomHostnameNewResponseStatusBlocked:
		return true
	}
	return false
}

type CustomHostnameListResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string                        `json:"hostname,required"`
	SSL      CustomHostnameListResponseSSL `json:"ssl,required"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSNI string `json:"custom_origin_sni"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameListResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameListResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CustomHostnameListResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string                       `json:"verification_errors"`
	JSON               customHostnameListResponseJSON `json:"-"`
}

// customHostnameListResponseJSON contains the JSON metadata for the struct
// [CustomHostnameListResponse]
type customHostnameListResponseJSON struct {
	ID                        apijson.Field
	Hostname                  apijson.Field
	SSL                       apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSNI           apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameListResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod BundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority shared.CertificateCA `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []string `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method DCVMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string                                `json:"serial_number"`
	Settings     CustomHostnameListResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameListResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type DomainValidationType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameListResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameListResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                              `json:"wildcard"`
	JSON     customHostnameListResponseSSLJSON `json:"-"`
}

// customHostnameListResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameListResponseSSL]
type customHostnameListResponseSSLJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseSSLJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameListResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameListResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameListResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameListResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameListResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameListResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameListResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameListResponseSSLSettings]
type customHostnameListResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseSSLSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CustomHostnameListResponseSSLSettingsEarlyHints string

const (
	CustomHostnameListResponseSSLSettingsEarlyHintsOn  CustomHostnameListResponseSSLSettingsEarlyHints = "on"
	CustomHostnameListResponseSSLSettingsEarlyHintsOff CustomHostnameListResponseSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameListResponseSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseSSLSettingsEarlyHintsOn, CustomHostnameListResponseSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameListResponseSSLSettingsHTTP2 string

const (
	CustomHostnameListResponseSSLSettingsHTTP2On  CustomHostnameListResponseSSLSettingsHTTP2 = "on"
	CustomHostnameListResponseSSLSettingsHTTP2Off CustomHostnameListResponseSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameListResponseSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseSSLSettingsHTTP2On, CustomHostnameListResponseSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameListResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_0 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_1 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_2 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_3 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameListResponseSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseSSLSettingsMinTLSVersion1_0, CustomHostnameListResponseSSLSettingsMinTLSVersion1_1, CustomHostnameListResponseSSLSettingsMinTLSVersion1_2, CustomHostnameListResponseSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameListResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameListResponseSSLSettingsTLS1_3On  CustomHostnameListResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameListResponseSSLSettingsTLS1_3Off CustomHostnameListResponseSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameListResponseSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseSSLSettingsTLS1_3On, CustomHostnameListResponseSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CustomHostnameListResponseSSLStatus string

const (
	CustomHostnameListResponseSSLStatusInitializing         CustomHostnameListResponseSSLStatus = "initializing"
	CustomHostnameListResponseSSLStatusPendingValidation    CustomHostnameListResponseSSLStatus = "pending_validation"
	CustomHostnameListResponseSSLStatusDeleted              CustomHostnameListResponseSSLStatus = "deleted"
	CustomHostnameListResponseSSLStatusPendingIssuance      CustomHostnameListResponseSSLStatus = "pending_issuance"
	CustomHostnameListResponseSSLStatusPendingDeployment    CustomHostnameListResponseSSLStatus = "pending_deployment"
	CustomHostnameListResponseSSLStatusPendingDeletion      CustomHostnameListResponseSSLStatus = "pending_deletion"
	CustomHostnameListResponseSSLStatusPendingExpiration    CustomHostnameListResponseSSLStatus = "pending_expiration"
	CustomHostnameListResponseSSLStatusExpired              CustomHostnameListResponseSSLStatus = "expired"
	CustomHostnameListResponseSSLStatusActive               CustomHostnameListResponseSSLStatus = "active"
	CustomHostnameListResponseSSLStatusInitializingTimedOut CustomHostnameListResponseSSLStatus = "initializing_timed_out"
	CustomHostnameListResponseSSLStatusValidationTimedOut   CustomHostnameListResponseSSLStatus = "validation_timed_out"
	CustomHostnameListResponseSSLStatusIssuanceTimedOut     CustomHostnameListResponseSSLStatus = "issuance_timed_out"
	CustomHostnameListResponseSSLStatusDeploymentTimedOut   CustomHostnameListResponseSSLStatus = "deployment_timed_out"
	CustomHostnameListResponseSSLStatusDeletionTimedOut     CustomHostnameListResponseSSLStatus = "deletion_timed_out"
	CustomHostnameListResponseSSLStatusPendingCleanup       CustomHostnameListResponseSSLStatus = "pending_cleanup"
	CustomHostnameListResponseSSLStatusStagingDeployment    CustomHostnameListResponseSSLStatus = "staging_deployment"
	CustomHostnameListResponseSSLStatusStagingActive        CustomHostnameListResponseSSLStatus = "staging_active"
	CustomHostnameListResponseSSLStatusDeactivating         CustomHostnameListResponseSSLStatus = "deactivating"
	CustomHostnameListResponseSSLStatusInactive             CustomHostnameListResponseSSLStatus = "inactive"
	CustomHostnameListResponseSSLStatusBackupIssued         CustomHostnameListResponseSSLStatus = "backup_issued"
	CustomHostnameListResponseSSLStatusHoldingDeployment    CustomHostnameListResponseSSLStatus = "holding_deployment"
)

func (r CustomHostnameListResponseSSLStatus) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseSSLStatusInitializing, CustomHostnameListResponseSSLStatusPendingValidation, CustomHostnameListResponseSSLStatusDeleted, CustomHostnameListResponseSSLStatusPendingIssuance, CustomHostnameListResponseSSLStatusPendingDeployment, CustomHostnameListResponseSSLStatusPendingDeletion, CustomHostnameListResponseSSLStatusPendingExpiration, CustomHostnameListResponseSSLStatusExpired, CustomHostnameListResponseSSLStatusActive, CustomHostnameListResponseSSLStatusInitializingTimedOut, CustomHostnameListResponseSSLStatusValidationTimedOut, CustomHostnameListResponseSSLStatusIssuanceTimedOut, CustomHostnameListResponseSSLStatusDeploymentTimedOut, CustomHostnameListResponseSSLStatusDeletionTimedOut, CustomHostnameListResponseSSLStatusPendingCleanup, CustomHostnameListResponseSSLStatusStagingDeployment, CustomHostnameListResponseSSLStatusStagingActive, CustomHostnameListResponseSSLStatusDeactivating, CustomHostnameListResponseSSLStatusInactive, CustomHostnameListResponseSSLStatusBackupIssued, CustomHostnameListResponseSSLStatusHoldingDeployment:
		return true
	}
	return false
}

type CustomHostnameListResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                           `json:"message"`
	JSON    customHostnameListResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameListResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameListResponseSSLValidationError]
type customHostnameListResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseSSLValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameListResponseSSLValidationRecord struct {
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
	TXTValue string                                            `json:"txt_value"`
	JSON     customHostnameListResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameListResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameListResponseSSLValidationRecord]
type customHostnameListResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseSSLValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameListResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameListResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                              `json:"value"`
	JSON  customHostnameListResponseOwnershipVerificationJSON `json:"-"`
}

// customHostnameListResponseOwnershipVerificationJSON contains the JSON metadata
// for the struct [CustomHostnameListResponseOwnershipVerification]
type customHostnameListResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CustomHostnameListResponseOwnershipVerificationType string

const (
	CustomHostnameListResponseOwnershipVerificationTypeTXT CustomHostnameListResponseOwnershipVerificationType = "txt"
)

func (r CustomHostnameListResponseOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseOwnershipVerificationTypeTXT:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameListResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                  `json:"http_url"`
	JSON    customHostnameListResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameListResponseOwnershipVerificationHTTPJSON contains the JSON
// metadata for the struct [CustomHostnameListResponseOwnershipVerificationHTTP]
type customHostnameListResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CustomHostnameListResponseStatus string

const (
	CustomHostnameListResponseStatusActive             CustomHostnameListResponseStatus = "active"
	CustomHostnameListResponseStatusPending            CustomHostnameListResponseStatus = "pending"
	CustomHostnameListResponseStatusActiveRedeploying  CustomHostnameListResponseStatus = "active_redeploying"
	CustomHostnameListResponseStatusMoved              CustomHostnameListResponseStatus = "moved"
	CustomHostnameListResponseStatusPendingDeletion    CustomHostnameListResponseStatus = "pending_deletion"
	CustomHostnameListResponseStatusDeleted            CustomHostnameListResponseStatus = "deleted"
	CustomHostnameListResponseStatusPendingBlocked     CustomHostnameListResponseStatus = "pending_blocked"
	CustomHostnameListResponseStatusPendingMigration   CustomHostnameListResponseStatus = "pending_migration"
	CustomHostnameListResponseStatusPendingProvisioned CustomHostnameListResponseStatus = "pending_provisioned"
	CustomHostnameListResponseStatusTestPending        CustomHostnameListResponseStatus = "test_pending"
	CustomHostnameListResponseStatusTestActive         CustomHostnameListResponseStatus = "test_active"
	CustomHostnameListResponseStatusTestActiveApex     CustomHostnameListResponseStatus = "test_active_apex"
	CustomHostnameListResponseStatusTestBlocked        CustomHostnameListResponseStatus = "test_blocked"
	CustomHostnameListResponseStatusTestFailed         CustomHostnameListResponseStatus = "test_failed"
	CustomHostnameListResponseStatusProvisioned        CustomHostnameListResponseStatus = "provisioned"
	CustomHostnameListResponseStatusBlocked            CustomHostnameListResponseStatus = "blocked"
)

func (r CustomHostnameListResponseStatus) IsKnown() bool {
	switch r {
	case CustomHostnameListResponseStatusActive, CustomHostnameListResponseStatusPending, CustomHostnameListResponseStatusActiveRedeploying, CustomHostnameListResponseStatusMoved, CustomHostnameListResponseStatusPendingDeletion, CustomHostnameListResponseStatusDeleted, CustomHostnameListResponseStatusPendingBlocked, CustomHostnameListResponseStatusPendingMigration, CustomHostnameListResponseStatusPendingProvisioned, CustomHostnameListResponseStatusTestPending, CustomHostnameListResponseStatusTestActive, CustomHostnameListResponseStatusTestActiveApex, CustomHostnameListResponseStatusTestBlocked, CustomHostnameListResponseStatusTestFailed, CustomHostnameListResponseStatusProvisioned, CustomHostnameListResponseStatusBlocked:
		return true
	}
	return false
}

type CustomHostnameDeleteResponse struct {
	// Identifier.
	ID   string                           `json:"id"`
	JSON customHostnameDeleteResponseJSON `json:"-"`
}

// customHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [CustomHostnameDeleteResponse]
type customHostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string                        `json:"hostname,required"`
	SSL      CustomHostnameEditResponseSSL `json:"ssl,required"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSNI string `json:"custom_origin_sni"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameEditResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameEditResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CustomHostnameEditResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string                       `json:"verification_errors"`
	JSON               customHostnameEditResponseJSON `json:"-"`
}

// customHostnameEditResponseJSON contains the JSON metadata for the struct
// [CustomHostnameEditResponse]
type customHostnameEditResponseJSON struct {
	ID                        apijson.Field
	Hostname                  apijson.Field
	SSL                       apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSNI           apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod BundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority shared.CertificateCA `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []string `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method DCVMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string                                `json:"serial_number"`
	Settings     CustomHostnameEditResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameEditResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type DomainValidationType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameEditResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameEditResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                              `json:"wildcard"`
	JSON     customHostnameEditResponseSSLJSON `json:"-"`
}

// customHostnameEditResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameEditResponseSSL]
type customHostnameEditResponseSSLJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomHostnameEditResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseSSLJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameEditResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameEditResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameEditResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameEditResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameEditResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameEditResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameEditResponseSSLSettings]
type customHostnameEditResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameEditResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseSSLSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CustomHostnameEditResponseSSLSettingsEarlyHints string

const (
	CustomHostnameEditResponseSSLSettingsEarlyHintsOn  CustomHostnameEditResponseSSLSettingsEarlyHints = "on"
	CustomHostnameEditResponseSSLSettingsEarlyHintsOff CustomHostnameEditResponseSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameEditResponseSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseSSLSettingsEarlyHintsOn, CustomHostnameEditResponseSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameEditResponseSSLSettingsHTTP2 string

const (
	CustomHostnameEditResponseSSLSettingsHTTP2On  CustomHostnameEditResponseSSLSettingsHTTP2 = "on"
	CustomHostnameEditResponseSSLSettingsHTTP2Off CustomHostnameEditResponseSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameEditResponseSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseSSLSettingsHTTP2On, CustomHostnameEditResponseSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameEditResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameEditResponseSSLSettingsMinTLSVersion1_0 CustomHostnameEditResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameEditResponseSSLSettingsMinTLSVersion1_1 CustomHostnameEditResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameEditResponseSSLSettingsMinTLSVersion1_2 CustomHostnameEditResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameEditResponseSSLSettingsMinTLSVersion1_3 CustomHostnameEditResponseSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameEditResponseSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseSSLSettingsMinTLSVersion1_0, CustomHostnameEditResponseSSLSettingsMinTLSVersion1_1, CustomHostnameEditResponseSSLSettingsMinTLSVersion1_2, CustomHostnameEditResponseSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameEditResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameEditResponseSSLSettingsTLS1_3On  CustomHostnameEditResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameEditResponseSSLSettingsTLS1_3Off CustomHostnameEditResponseSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameEditResponseSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseSSLSettingsTLS1_3On, CustomHostnameEditResponseSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CustomHostnameEditResponseSSLStatus string

const (
	CustomHostnameEditResponseSSLStatusInitializing         CustomHostnameEditResponseSSLStatus = "initializing"
	CustomHostnameEditResponseSSLStatusPendingValidation    CustomHostnameEditResponseSSLStatus = "pending_validation"
	CustomHostnameEditResponseSSLStatusDeleted              CustomHostnameEditResponseSSLStatus = "deleted"
	CustomHostnameEditResponseSSLStatusPendingIssuance      CustomHostnameEditResponseSSLStatus = "pending_issuance"
	CustomHostnameEditResponseSSLStatusPendingDeployment    CustomHostnameEditResponseSSLStatus = "pending_deployment"
	CustomHostnameEditResponseSSLStatusPendingDeletion      CustomHostnameEditResponseSSLStatus = "pending_deletion"
	CustomHostnameEditResponseSSLStatusPendingExpiration    CustomHostnameEditResponseSSLStatus = "pending_expiration"
	CustomHostnameEditResponseSSLStatusExpired              CustomHostnameEditResponseSSLStatus = "expired"
	CustomHostnameEditResponseSSLStatusActive               CustomHostnameEditResponseSSLStatus = "active"
	CustomHostnameEditResponseSSLStatusInitializingTimedOut CustomHostnameEditResponseSSLStatus = "initializing_timed_out"
	CustomHostnameEditResponseSSLStatusValidationTimedOut   CustomHostnameEditResponseSSLStatus = "validation_timed_out"
	CustomHostnameEditResponseSSLStatusIssuanceTimedOut     CustomHostnameEditResponseSSLStatus = "issuance_timed_out"
	CustomHostnameEditResponseSSLStatusDeploymentTimedOut   CustomHostnameEditResponseSSLStatus = "deployment_timed_out"
	CustomHostnameEditResponseSSLStatusDeletionTimedOut     CustomHostnameEditResponseSSLStatus = "deletion_timed_out"
	CustomHostnameEditResponseSSLStatusPendingCleanup       CustomHostnameEditResponseSSLStatus = "pending_cleanup"
	CustomHostnameEditResponseSSLStatusStagingDeployment    CustomHostnameEditResponseSSLStatus = "staging_deployment"
	CustomHostnameEditResponseSSLStatusStagingActive        CustomHostnameEditResponseSSLStatus = "staging_active"
	CustomHostnameEditResponseSSLStatusDeactivating         CustomHostnameEditResponseSSLStatus = "deactivating"
	CustomHostnameEditResponseSSLStatusInactive             CustomHostnameEditResponseSSLStatus = "inactive"
	CustomHostnameEditResponseSSLStatusBackupIssued         CustomHostnameEditResponseSSLStatus = "backup_issued"
	CustomHostnameEditResponseSSLStatusHoldingDeployment    CustomHostnameEditResponseSSLStatus = "holding_deployment"
)

func (r CustomHostnameEditResponseSSLStatus) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseSSLStatusInitializing, CustomHostnameEditResponseSSLStatusPendingValidation, CustomHostnameEditResponseSSLStatusDeleted, CustomHostnameEditResponseSSLStatusPendingIssuance, CustomHostnameEditResponseSSLStatusPendingDeployment, CustomHostnameEditResponseSSLStatusPendingDeletion, CustomHostnameEditResponseSSLStatusPendingExpiration, CustomHostnameEditResponseSSLStatusExpired, CustomHostnameEditResponseSSLStatusActive, CustomHostnameEditResponseSSLStatusInitializingTimedOut, CustomHostnameEditResponseSSLStatusValidationTimedOut, CustomHostnameEditResponseSSLStatusIssuanceTimedOut, CustomHostnameEditResponseSSLStatusDeploymentTimedOut, CustomHostnameEditResponseSSLStatusDeletionTimedOut, CustomHostnameEditResponseSSLStatusPendingCleanup, CustomHostnameEditResponseSSLStatusStagingDeployment, CustomHostnameEditResponseSSLStatusStagingActive, CustomHostnameEditResponseSSLStatusDeactivating, CustomHostnameEditResponseSSLStatusInactive, CustomHostnameEditResponseSSLStatusBackupIssued, CustomHostnameEditResponseSSLStatusHoldingDeployment:
		return true
	}
	return false
}

type CustomHostnameEditResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                           `json:"message"`
	JSON    customHostnameEditResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameEditResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameEditResponseSSLValidationError]
type customHostnameEditResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseSSLValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseSSLValidationRecord struct {
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
	TXTValue string                                            `json:"txt_value"`
	JSON     customHostnameEditResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameEditResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameEditResponseSSLValidationRecord]
type customHostnameEditResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseSSLValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameEditResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameEditResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                              `json:"value"`
	JSON  customHostnameEditResponseOwnershipVerificationJSON `json:"-"`
}

// customHostnameEditResponseOwnershipVerificationJSON contains the JSON metadata
// for the struct [CustomHostnameEditResponseOwnershipVerification]
type customHostnameEditResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CustomHostnameEditResponseOwnershipVerificationType string

const (
	CustomHostnameEditResponseOwnershipVerificationTypeTXT CustomHostnameEditResponseOwnershipVerificationType = "txt"
)

func (r CustomHostnameEditResponseOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseOwnershipVerificationTypeTXT:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameEditResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                  `json:"http_url"`
	JSON    customHostnameEditResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameEditResponseOwnershipVerificationHTTPJSON contains the JSON
// metadata for the struct [CustomHostnameEditResponseOwnershipVerificationHTTP]
type customHostnameEditResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CustomHostnameEditResponseStatus string

const (
	CustomHostnameEditResponseStatusActive             CustomHostnameEditResponseStatus = "active"
	CustomHostnameEditResponseStatusPending            CustomHostnameEditResponseStatus = "pending"
	CustomHostnameEditResponseStatusActiveRedeploying  CustomHostnameEditResponseStatus = "active_redeploying"
	CustomHostnameEditResponseStatusMoved              CustomHostnameEditResponseStatus = "moved"
	CustomHostnameEditResponseStatusPendingDeletion    CustomHostnameEditResponseStatus = "pending_deletion"
	CustomHostnameEditResponseStatusDeleted            CustomHostnameEditResponseStatus = "deleted"
	CustomHostnameEditResponseStatusPendingBlocked     CustomHostnameEditResponseStatus = "pending_blocked"
	CustomHostnameEditResponseStatusPendingMigration   CustomHostnameEditResponseStatus = "pending_migration"
	CustomHostnameEditResponseStatusPendingProvisioned CustomHostnameEditResponseStatus = "pending_provisioned"
	CustomHostnameEditResponseStatusTestPending        CustomHostnameEditResponseStatus = "test_pending"
	CustomHostnameEditResponseStatusTestActive         CustomHostnameEditResponseStatus = "test_active"
	CustomHostnameEditResponseStatusTestActiveApex     CustomHostnameEditResponseStatus = "test_active_apex"
	CustomHostnameEditResponseStatusTestBlocked        CustomHostnameEditResponseStatus = "test_blocked"
	CustomHostnameEditResponseStatusTestFailed         CustomHostnameEditResponseStatus = "test_failed"
	CustomHostnameEditResponseStatusProvisioned        CustomHostnameEditResponseStatus = "provisioned"
	CustomHostnameEditResponseStatusBlocked            CustomHostnameEditResponseStatus = "blocked"
)

func (r CustomHostnameEditResponseStatus) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseStatusActive, CustomHostnameEditResponseStatusPending, CustomHostnameEditResponseStatusActiveRedeploying, CustomHostnameEditResponseStatusMoved, CustomHostnameEditResponseStatusPendingDeletion, CustomHostnameEditResponseStatusDeleted, CustomHostnameEditResponseStatusPendingBlocked, CustomHostnameEditResponseStatusPendingMigration, CustomHostnameEditResponseStatusPendingProvisioned, CustomHostnameEditResponseStatusTestPending, CustomHostnameEditResponseStatusTestActive, CustomHostnameEditResponseStatusTestActiveApex, CustomHostnameEditResponseStatusTestBlocked, CustomHostnameEditResponseStatusTestFailed, CustomHostnameEditResponseStatusProvisioned, CustomHostnameEditResponseStatusBlocked:
		return true
	}
	return false
}

type CustomHostnameGetResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string                       `json:"hostname,required"`
	SSL      CustomHostnameGetResponseSSL `json:"ssl,required"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSNI string `json:"custom_origin_sni"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameGetResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameGetResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CustomHostnameGetResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string                      `json:"verification_errors"`
	JSON               customHostnameGetResponseJSON `json:"-"`
}

// customHostnameGetResponseJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponse]
type customHostnameGetResponseJSON struct {
	ID                        apijson.Field
	Hostname                  apijson.Field
	SSL                       apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSNI           apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod BundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority shared.CertificateCA `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []string `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method DCVMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string                               `json:"serial_number"`
	Settings     CustomHostnameGetResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameGetResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type DomainValidationType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameGetResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameGetResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                             `json:"wildcard"`
	JSON     customHostnameGetResponseSSLJSON `json:"-"`
}

// customHostnameGetResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponseSSL]
type customHostnameGetResponseSSLJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseSSLJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameGetResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameGetResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameGetResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameGetResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameGetResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameGetResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseSSLSettings]
type customHostnameGetResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseSSLSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CustomHostnameGetResponseSSLSettingsEarlyHints string

const (
	CustomHostnameGetResponseSSLSettingsEarlyHintsOn  CustomHostnameGetResponseSSLSettingsEarlyHints = "on"
	CustomHostnameGetResponseSSLSettingsEarlyHintsOff CustomHostnameGetResponseSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameGetResponseSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseSSLSettingsEarlyHintsOn, CustomHostnameGetResponseSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameGetResponseSSLSettingsHTTP2 string

const (
	CustomHostnameGetResponseSSLSettingsHTTP2On  CustomHostnameGetResponseSSLSettingsHTTP2 = "on"
	CustomHostnameGetResponseSSLSettingsHTTP2Off CustomHostnameGetResponseSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameGetResponseSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseSSLSettingsHTTP2On, CustomHostnameGetResponseSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameGetResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_0 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_1 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_2 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_3 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameGetResponseSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseSSLSettingsMinTLSVersion1_0, CustomHostnameGetResponseSSLSettingsMinTLSVersion1_1, CustomHostnameGetResponseSSLSettingsMinTLSVersion1_2, CustomHostnameGetResponseSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameGetResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameGetResponseSSLSettingsTLS1_3On  CustomHostnameGetResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameGetResponseSSLSettingsTLS1_3Off CustomHostnameGetResponseSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameGetResponseSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseSSLSettingsTLS1_3On, CustomHostnameGetResponseSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CustomHostnameGetResponseSSLStatus string

const (
	CustomHostnameGetResponseSSLStatusInitializing         CustomHostnameGetResponseSSLStatus = "initializing"
	CustomHostnameGetResponseSSLStatusPendingValidation    CustomHostnameGetResponseSSLStatus = "pending_validation"
	CustomHostnameGetResponseSSLStatusDeleted              CustomHostnameGetResponseSSLStatus = "deleted"
	CustomHostnameGetResponseSSLStatusPendingIssuance      CustomHostnameGetResponseSSLStatus = "pending_issuance"
	CustomHostnameGetResponseSSLStatusPendingDeployment    CustomHostnameGetResponseSSLStatus = "pending_deployment"
	CustomHostnameGetResponseSSLStatusPendingDeletion      CustomHostnameGetResponseSSLStatus = "pending_deletion"
	CustomHostnameGetResponseSSLStatusPendingExpiration    CustomHostnameGetResponseSSLStatus = "pending_expiration"
	CustomHostnameGetResponseSSLStatusExpired              CustomHostnameGetResponseSSLStatus = "expired"
	CustomHostnameGetResponseSSLStatusActive               CustomHostnameGetResponseSSLStatus = "active"
	CustomHostnameGetResponseSSLStatusInitializingTimedOut CustomHostnameGetResponseSSLStatus = "initializing_timed_out"
	CustomHostnameGetResponseSSLStatusValidationTimedOut   CustomHostnameGetResponseSSLStatus = "validation_timed_out"
	CustomHostnameGetResponseSSLStatusIssuanceTimedOut     CustomHostnameGetResponseSSLStatus = "issuance_timed_out"
	CustomHostnameGetResponseSSLStatusDeploymentTimedOut   CustomHostnameGetResponseSSLStatus = "deployment_timed_out"
	CustomHostnameGetResponseSSLStatusDeletionTimedOut     CustomHostnameGetResponseSSLStatus = "deletion_timed_out"
	CustomHostnameGetResponseSSLStatusPendingCleanup       CustomHostnameGetResponseSSLStatus = "pending_cleanup"
	CustomHostnameGetResponseSSLStatusStagingDeployment    CustomHostnameGetResponseSSLStatus = "staging_deployment"
	CustomHostnameGetResponseSSLStatusStagingActive        CustomHostnameGetResponseSSLStatus = "staging_active"
	CustomHostnameGetResponseSSLStatusDeactivating         CustomHostnameGetResponseSSLStatus = "deactivating"
	CustomHostnameGetResponseSSLStatusInactive             CustomHostnameGetResponseSSLStatus = "inactive"
	CustomHostnameGetResponseSSLStatusBackupIssued         CustomHostnameGetResponseSSLStatus = "backup_issued"
	CustomHostnameGetResponseSSLStatusHoldingDeployment    CustomHostnameGetResponseSSLStatus = "holding_deployment"
)

func (r CustomHostnameGetResponseSSLStatus) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseSSLStatusInitializing, CustomHostnameGetResponseSSLStatusPendingValidation, CustomHostnameGetResponseSSLStatusDeleted, CustomHostnameGetResponseSSLStatusPendingIssuance, CustomHostnameGetResponseSSLStatusPendingDeployment, CustomHostnameGetResponseSSLStatusPendingDeletion, CustomHostnameGetResponseSSLStatusPendingExpiration, CustomHostnameGetResponseSSLStatusExpired, CustomHostnameGetResponseSSLStatusActive, CustomHostnameGetResponseSSLStatusInitializingTimedOut, CustomHostnameGetResponseSSLStatusValidationTimedOut, CustomHostnameGetResponseSSLStatusIssuanceTimedOut, CustomHostnameGetResponseSSLStatusDeploymentTimedOut, CustomHostnameGetResponseSSLStatusDeletionTimedOut, CustomHostnameGetResponseSSLStatusPendingCleanup, CustomHostnameGetResponseSSLStatusStagingDeployment, CustomHostnameGetResponseSSLStatusStagingActive, CustomHostnameGetResponseSSLStatusDeactivating, CustomHostnameGetResponseSSLStatusInactive, CustomHostnameGetResponseSSLStatusBackupIssued, CustomHostnameGetResponseSSLStatusHoldingDeployment:
		return true
	}
	return false
}

type CustomHostnameGetResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                          `json:"message"`
	JSON    customHostnameGetResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameGetResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameGetResponseSSLValidationError]
type customHostnameGetResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseSSLValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseSSLValidationRecord struct {
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
	TXTValue string                                           `json:"txt_value"`
	JSON     customHostnameGetResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameGetResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameGetResponseSSLValidationRecord]
type customHostnameGetResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseSSLValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameGetResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameGetResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                             `json:"value"`
	JSON  customHostnameGetResponseOwnershipVerificationJSON `json:"-"`
}

// customHostnameGetResponseOwnershipVerificationJSON contains the JSON metadata
// for the struct [CustomHostnameGetResponseOwnershipVerification]
type customHostnameGetResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CustomHostnameGetResponseOwnershipVerificationType string

const (
	CustomHostnameGetResponseOwnershipVerificationTypeTXT CustomHostnameGetResponseOwnershipVerificationType = "txt"
)

func (r CustomHostnameGetResponseOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseOwnershipVerificationTypeTXT:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameGetResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                 `json:"http_url"`
	JSON    customHostnameGetResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameGetResponseOwnershipVerificationHTTPJSON contains the JSON
// metadata for the struct [CustomHostnameGetResponseOwnershipVerificationHTTP]
type customHostnameGetResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CustomHostnameGetResponseStatus string

const (
	CustomHostnameGetResponseStatusActive             CustomHostnameGetResponseStatus = "active"
	CustomHostnameGetResponseStatusPending            CustomHostnameGetResponseStatus = "pending"
	CustomHostnameGetResponseStatusActiveRedeploying  CustomHostnameGetResponseStatus = "active_redeploying"
	CustomHostnameGetResponseStatusMoved              CustomHostnameGetResponseStatus = "moved"
	CustomHostnameGetResponseStatusPendingDeletion    CustomHostnameGetResponseStatus = "pending_deletion"
	CustomHostnameGetResponseStatusDeleted            CustomHostnameGetResponseStatus = "deleted"
	CustomHostnameGetResponseStatusPendingBlocked     CustomHostnameGetResponseStatus = "pending_blocked"
	CustomHostnameGetResponseStatusPendingMigration   CustomHostnameGetResponseStatus = "pending_migration"
	CustomHostnameGetResponseStatusPendingProvisioned CustomHostnameGetResponseStatus = "pending_provisioned"
	CustomHostnameGetResponseStatusTestPending        CustomHostnameGetResponseStatus = "test_pending"
	CustomHostnameGetResponseStatusTestActive         CustomHostnameGetResponseStatus = "test_active"
	CustomHostnameGetResponseStatusTestActiveApex     CustomHostnameGetResponseStatus = "test_active_apex"
	CustomHostnameGetResponseStatusTestBlocked        CustomHostnameGetResponseStatus = "test_blocked"
	CustomHostnameGetResponseStatusTestFailed         CustomHostnameGetResponseStatus = "test_failed"
	CustomHostnameGetResponseStatusProvisioned        CustomHostnameGetResponseStatus = "provisioned"
	CustomHostnameGetResponseStatusBlocked            CustomHostnameGetResponseStatus = "blocked"
)

func (r CustomHostnameGetResponseStatus) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseStatusActive, CustomHostnameGetResponseStatusPending, CustomHostnameGetResponseStatusActiveRedeploying, CustomHostnameGetResponseStatusMoved, CustomHostnameGetResponseStatusPendingDeletion, CustomHostnameGetResponseStatusDeleted, CustomHostnameGetResponseStatusPendingBlocked, CustomHostnameGetResponseStatusPendingMigration, CustomHostnameGetResponseStatusPendingProvisioned, CustomHostnameGetResponseStatusTestPending, CustomHostnameGetResponseStatusTestActive, CustomHostnameGetResponseStatusTestActiveApex, CustomHostnameGetResponseStatusTestBlocked, CustomHostnameGetResponseStatusTestFailed, CustomHostnameGetResponseStatusProvisioned, CustomHostnameGetResponseStatusBlocked:
		return true
	}
	return false
}

type CustomHostnameNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname param.Field[string] `json:"hostname,required"`
	// SSL properties used when creating the custom hostname.
	SSL param.Field[CustomHostnameNewParamsSSL] `json:"ssl,required"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata param.Field[map[string]string] `json:"custom_metadata"`
}

func (r CustomHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameNewParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[shared.CertificateCA] `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
	// Array of custom certificate and key pairs (1 or 2 pairs allowed)
	CustomCERTBundle param.Field[[]CustomHostnameNewParamsSSLCustomCERTBundle] `json:"custom_cert_bundle"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[DCVMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameNewParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[DomainValidationType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameNewParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameNewParamsSSLCustomCERTBundle struct {
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate,required"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key,required"`
}

func (r CustomHostnameNewParamsSSLCustomCERTBundle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL specific settings.
type CustomHostnameNewParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameNewParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameNewParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameNewParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameNewParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameNewParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameNewParamsSSLSettingsEarlyHints string

const (
	CustomHostnameNewParamsSSLSettingsEarlyHintsOn  CustomHostnameNewParamsSSLSettingsEarlyHints = "on"
	CustomHostnameNewParamsSSLSettingsEarlyHintsOff CustomHostnameNewParamsSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameNewParamsSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsEarlyHintsOn, CustomHostnameNewParamsSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameNewParamsSSLSettingsHTTP2 string

const (
	CustomHostnameNewParamsSSLSettingsHTTP2On  CustomHostnameNewParamsSSLSettingsHTTP2 = "on"
	CustomHostnameNewParamsSSLSettingsHTTP2Off CustomHostnameNewParamsSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameNewParamsSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsHTTP2On, CustomHostnameNewParamsSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameNewParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_0 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_1 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_2 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_3 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameNewParamsSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsMinTLSVersion1_0, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_1, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_2, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameNewParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameNewParamsSSLSettingsTLS1_3On  CustomHostnameNewParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameNewParamsSSLSettingsTLS1_3Off CustomHostnameNewParamsSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameNewParamsSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsTLS1_3On, CustomHostnameNewParamsSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

type CustomHostnameNewResponseEnvelope struct {
	Errors   []CustomHostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomHostnameNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomHostnameNewResponse                `json:"result"`
	JSON    customHostnameNewResponseEnvelopeJSON    `json:"-"`
}

// customHostnameNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponseEnvelope]
type customHostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           CustomHostnameNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             customHostnameNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// customHostnameNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseEnvelopeErrors]
type customHostnameNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    customHostnameNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customHostnameNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CustomHostnameNewResponseEnvelopeErrorsSource]
type customHostnameNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CustomHostnameNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             customHostnameNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// customHostnameNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseEnvelopeMessages]
type customHostnameNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    customHostnameNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customHostnameNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomHostnameNewResponseEnvelopeMessagesSource]
type customHostnameNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomHostnameNewResponseEnvelopeSuccess bool

const (
	CustomHostnameNewResponseEnvelopeSuccessTrue CustomHostnameNewResponseEnvelopeSuccess = true
)

func (r CustomHostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomHostnameListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Hostname ID to match against. This ID was generated and returned during the
	// initial custom_hostname creation. This parameter cannot be used with the
	// 'hostname' parameter.
	ID param.Field[string] `query:"id"`
	// Direction to order hostnames.
	Direction param.Field[CustomHostnameListParamsDirection] `query:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with
	// the 'id' parameter.
	Hostname param.Field[string] `query:"hostname"`
	// Field to order hostnames by.
	Order param.Field[CustomHostnameListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of hostnames per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether to filter hostnames based on if they have SSL enabled.
	SSL param.Field[CustomHostnameListParamsSSL] `query:"ssl"`
}

// URLQuery serializes [CustomHostnameListParams]'s query parameters as
// `url.Values`.
func (r CustomHostnameListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order hostnames.
type CustomHostnameListParamsDirection string

const (
	CustomHostnameListParamsDirectionAsc  CustomHostnameListParamsDirection = "asc"
	CustomHostnameListParamsDirectionDesc CustomHostnameListParamsDirection = "desc"
)

func (r CustomHostnameListParamsDirection) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsDirectionAsc, CustomHostnameListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order hostnames by.
type CustomHostnameListParamsOrder string

const (
	CustomHostnameListParamsOrderSSL       CustomHostnameListParamsOrder = "ssl"
	CustomHostnameListParamsOrderSSLStatus CustomHostnameListParamsOrder = "ssl_status"
)

func (r CustomHostnameListParamsOrder) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsOrderSSL, CustomHostnameListParamsOrderSSLStatus:
		return true
	}
	return false
}

// Whether to filter hostnames based on if they have SSL enabled.
type CustomHostnameListParamsSSL float64

const (
	CustomHostnameListParamsSSL0 CustomHostnameListParamsSSL = 0
	CustomHostnameListParamsSSL1 CustomHostnameListParamsSSL = 1
)

func (r CustomHostnameListParamsSSL) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsSSL0, CustomHostnameListParamsSSL1:
		return true
	}
	return false
}

type CustomHostnameDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomHostnameEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata param.Field[map[string]string] `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer param.Field[string] `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSNI param.Field[string] `json:"custom_origin_sni"`
	// SSL properties used when creating the custom hostname.
	SSL param.Field[CustomHostnameEditParamsSSL] `json:"ssl"`
}

func (r CustomHostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameEditParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[shared.CertificateCA] `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
	// Array of custom certificate and key pairs (1 or 2 pairs allowed)
	CustomCERTBundle param.Field[[]CustomHostnameEditParamsSSLCustomCERTBundle] `json:"custom_cert_bundle"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[DCVMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameEditParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[DomainValidationType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameEditParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameEditParamsSSLCustomCERTBundle struct {
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate,required"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key,required"`
}

func (r CustomHostnameEditParamsSSLCustomCERTBundle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL specific settings.
type CustomHostnameEditParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameEditParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameEditParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameEditParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameEditParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameEditParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameEditParamsSSLSettingsEarlyHints string

const (
	CustomHostnameEditParamsSSLSettingsEarlyHintsOn  CustomHostnameEditParamsSSLSettingsEarlyHints = "on"
	CustomHostnameEditParamsSSLSettingsEarlyHintsOff CustomHostnameEditParamsSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameEditParamsSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsEarlyHintsOn, CustomHostnameEditParamsSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameEditParamsSSLSettingsHTTP2 string

const (
	CustomHostnameEditParamsSSLSettingsHTTP2On  CustomHostnameEditParamsSSLSettingsHTTP2 = "on"
	CustomHostnameEditParamsSSLSettingsHTTP2Off CustomHostnameEditParamsSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameEditParamsSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsHTTP2On, CustomHostnameEditParamsSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameEditParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_0 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_1 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_2 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_3 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameEditParamsSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsMinTLSVersion1_0, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_1, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_2, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameEditParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameEditParamsSSLSettingsTLS1_3On  CustomHostnameEditParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameEditParamsSSLSettingsTLS1_3Off CustomHostnameEditParamsSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameEditParamsSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsTLS1_3On, CustomHostnameEditParamsSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

type CustomHostnameEditResponseEnvelope struct {
	Errors   []CustomHostnameEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomHostnameEditResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomHostnameEditResponse                `json:"result"`
	JSON    customHostnameEditResponseEnvelopeJSON    `json:"-"`
}

// customHostnameEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameEditResponseEnvelope]
type customHostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CustomHostnameEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             customHostnameEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// customHostnameEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomHostnameEditResponseEnvelopeErrors]
type customHostnameEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    customHostnameEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customHostnameEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CustomHostnameEditResponseEnvelopeErrorsSource]
type customHostnameEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           CustomHostnameEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             customHostnameEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// customHostnameEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomHostnameEditResponseEnvelopeMessages]
type customHostnameEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    customHostnameEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customHostnameEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomHostnameEditResponseEnvelopeMessagesSource]
type customHostnameEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomHostnameEditResponseEnvelopeSuccess bool

const (
	CustomHostnameEditResponseEnvelopeSuccessTrue CustomHostnameEditResponseEnvelopeSuccess = true
)

func (r CustomHostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomHostnameGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomHostnameGetResponseEnvelope struct {
	Errors   []CustomHostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomHostnameGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomHostnameGetResponse                `json:"result"`
	JSON    customHostnameGetResponseEnvelopeJSON    `json:"-"`
}

// customHostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponseEnvelope]
type customHostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           CustomHostnameGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             customHostnameGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// customHostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseEnvelopeErrors]
type customHostnameGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    customHostnameGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customHostnameGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CustomHostnameGetResponseEnvelopeErrorsSource]
type customHostnameGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CustomHostnameGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             customHostnameGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// customHostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseEnvelopeMessages]
type customHostnameGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    customHostnameGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customHostnameGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomHostnameGetResponseEnvelopeMessagesSource]
type customHostnameGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomHostnameGetResponseEnvelopeSuccess bool

const (
	CustomHostnameGetResponseEnvelopeSuccessTrue CustomHostnameGetResponseEnvelopeSuccess = true
)

func (r CustomHostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
