// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// CertificatePackCertificateService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificatePackCertificateService] method instead.
type CertificatePackCertificateService struct {
	Options []option.RequestOption
}

// NewCertificatePackCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCertificatePackCertificateService(opts ...option.RequestOption) (r *CertificatePackCertificateService) {
	r = &CertificatePackCertificateService{}
	r.Options = opts
	return
}

// Replace a single custom certificate within a certificate pack that contains two
// bundled certificates. The replacement must adhere to the following constraints.
// You can only replace an RSA certificate with another RSA certificate or an ECDSA
// certificate with another ECDSA certificate.
func (r *CertificatePackCertificateService) Update(ctx context.Context, customHostnameID string, certificatePackID string, certificateID string, params CertificatePackCertificateUpdateParams, opts ...option.RequestOption) (res *CertificatePackCertificateUpdateResponse, err error) {
	var env CertificatePackCertificateUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s/certificate_pack/%s/certificates/%s", params.ZoneID, customHostnameID, certificatePackID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a single custom certificate from a certificate pack that contains two
// bundled certificates. Deletion is subject to the following constraints. You
// cannot delete a certificate if it is the only remaining certificate in the pack.
// At least one certificate must remain in the pack.
func (r *CertificatePackCertificateService) Delete(ctx context.Context, customHostnameID string, certificatePackID string, certificateID string, body CertificatePackCertificateDeleteParams, opts ...option.RequestOption) (res *CertificatePackCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s/certificate_pack/%s/certificates/%s", body.ZoneID, customHostnameID, certificatePackID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CertificatePackCertificateUpdateResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string                                      `json:"hostname,required"`
	SSL      CertificatePackCertificateUpdateResponseSSL `json:"ssl,required"`
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
	OwnershipVerification CertificatePackCertificateUpdateResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CertificatePackCertificateUpdateResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CertificatePackCertificateUpdateResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string                                     `json:"verification_errors"`
	JSON               certificatePackCertificateUpdateResponseJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [CertificatePackCertificateUpdateResponse]
type certificatePackCertificateUpdateResponseJSON struct {
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

func (r *CertificatePackCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseSSL struct {
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
	SerialNumber string                                              `json:"serial_number"`
	Settings     CertificatePackCertificateUpdateResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CertificatePackCertificateUpdateResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type DomainValidationType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CertificatePackCertificateUpdateResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CertificatePackCertificateUpdateResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                            `json:"wildcard"`
	JSON     certificatePackCertificateUpdateResponseSSLJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseSSLJSON contains the JSON metadata for
// the struct [CertificatePackCertificateUpdateResponseSSL]
type certificatePackCertificateUpdateResponseSSLJSON struct {
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

func (r *CertificatePackCertificateUpdateResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseSSLJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CertificatePackCertificateUpdateResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CertificatePackCertificateUpdateResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   certificatePackCertificateUpdateResponseSSLSettingsJSON   `json:"-"`
}

// certificatePackCertificateUpdateResponseSSLSettingsJSON contains the JSON
// metadata for the struct [CertificatePackCertificateUpdateResponseSSLSettings]
type certificatePackCertificateUpdateResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseSSLSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CertificatePackCertificateUpdateResponseSSLSettingsEarlyHints string

const (
	CertificatePackCertificateUpdateResponseSSLSettingsEarlyHintsOn  CertificatePackCertificateUpdateResponseSSLSettingsEarlyHints = "on"
	CertificatePackCertificateUpdateResponseSSLSettingsEarlyHintsOff CertificatePackCertificateUpdateResponseSSLSettingsEarlyHints = "off"
)

func (r CertificatePackCertificateUpdateResponseSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseSSLSettingsEarlyHintsOn, CertificatePackCertificateUpdateResponseSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CertificatePackCertificateUpdateResponseSSLSettingsHTTP2 string

const (
	CertificatePackCertificateUpdateResponseSSLSettingsHTTP2On  CertificatePackCertificateUpdateResponseSSLSettingsHTTP2 = "on"
	CertificatePackCertificateUpdateResponseSSLSettingsHTTP2Off CertificatePackCertificateUpdateResponseSSLSettingsHTTP2 = "off"
)

func (r CertificatePackCertificateUpdateResponseSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseSSLSettingsHTTP2On, CertificatePackCertificateUpdateResponseSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion string

const (
	CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_0 CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion = "1.0"
	CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_1 CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion = "1.1"
	CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_2 CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion = "1.2"
	CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_3 CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion = "1.3"
)

func (r CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_0, CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_1, CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_2, CertificatePackCertificateUpdateResponseSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3 string

const (
	CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3On  CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3 = "on"
	CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3Off CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3 = "off"
)

func (r CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3On, CertificatePackCertificateUpdateResponseSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CertificatePackCertificateUpdateResponseSSLStatus string

const (
	CertificatePackCertificateUpdateResponseSSLStatusInitializing         CertificatePackCertificateUpdateResponseSSLStatus = "initializing"
	CertificatePackCertificateUpdateResponseSSLStatusPendingValidation    CertificatePackCertificateUpdateResponseSSLStatus = "pending_validation"
	CertificatePackCertificateUpdateResponseSSLStatusDeleted              CertificatePackCertificateUpdateResponseSSLStatus = "deleted"
	CertificatePackCertificateUpdateResponseSSLStatusPendingIssuance      CertificatePackCertificateUpdateResponseSSLStatus = "pending_issuance"
	CertificatePackCertificateUpdateResponseSSLStatusPendingDeployment    CertificatePackCertificateUpdateResponseSSLStatus = "pending_deployment"
	CertificatePackCertificateUpdateResponseSSLStatusPendingDeletion      CertificatePackCertificateUpdateResponseSSLStatus = "pending_deletion"
	CertificatePackCertificateUpdateResponseSSLStatusPendingExpiration    CertificatePackCertificateUpdateResponseSSLStatus = "pending_expiration"
	CertificatePackCertificateUpdateResponseSSLStatusExpired              CertificatePackCertificateUpdateResponseSSLStatus = "expired"
	CertificatePackCertificateUpdateResponseSSLStatusActive               CertificatePackCertificateUpdateResponseSSLStatus = "active"
	CertificatePackCertificateUpdateResponseSSLStatusInitializingTimedOut CertificatePackCertificateUpdateResponseSSLStatus = "initializing_timed_out"
	CertificatePackCertificateUpdateResponseSSLStatusValidationTimedOut   CertificatePackCertificateUpdateResponseSSLStatus = "validation_timed_out"
	CertificatePackCertificateUpdateResponseSSLStatusIssuanceTimedOut     CertificatePackCertificateUpdateResponseSSLStatus = "issuance_timed_out"
	CertificatePackCertificateUpdateResponseSSLStatusDeploymentTimedOut   CertificatePackCertificateUpdateResponseSSLStatus = "deployment_timed_out"
	CertificatePackCertificateUpdateResponseSSLStatusDeletionTimedOut     CertificatePackCertificateUpdateResponseSSLStatus = "deletion_timed_out"
	CertificatePackCertificateUpdateResponseSSLStatusPendingCleanup       CertificatePackCertificateUpdateResponseSSLStatus = "pending_cleanup"
	CertificatePackCertificateUpdateResponseSSLStatusStagingDeployment    CertificatePackCertificateUpdateResponseSSLStatus = "staging_deployment"
	CertificatePackCertificateUpdateResponseSSLStatusStagingActive        CertificatePackCertificateUpdateResponseSSLStatus = "staging_active"
	CertificatePackCertificateUpdateResponseSSLStatusDeactivating         CertificatePackCertificateUpdateResponseSSLStatus = "deactivating"
	CertificatePackCertificateUpdateResponseSSLStatusInactive             CertificatePackCertificateUpdateResponseSSLStatus = "inactive"
	CertificatePackCertificateUpdateResponseSSLStatusBackupIssued         CertificatePackCertificateUpdateResponseSSLStatus = "backup_issued"
	CertificatePackCertificateUpdateResponseSSLStatusHoldingDeployment    CertificatePackCertificateUpdateResponseSSLStatus = "holding_deployment"
)

func (r CertificatePackCertificateUpdateResponseSSLStatus) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseSSLStatusInitializing, CertificatePackCertificateUpdateResponseSSLStatusPendingValidation, CertificatePackCertificateUpdateResponseSSLStatusDeleted, CertificatePackCertificateUpdateResponseSSLStatusPendingIssuance, CertificatePackCertificateUpdateResponseSSLStatusPendingDeployment, CertificatePackCertificateUpdateResponseSSLStatusPendingDeletion, CertificatePackCertificateUpdateResponseSSLStatusPendingExpiration, CertificatePackCertificateUpdateResponseSSLStatusExpired, CertificatePackCertificateUpdateResponseSSLStatusActive, CertificatePackCertificateUpdateResponseSSLStatusInitializingTimedOut, CertificatePackCertificateUpdateResponseSSLStatusValidationTimedOut, CertificatePackCertificateUpdateResponseSSLStatusIssuanceTimedOut, CertificatePackCertificateUpdateResponseSSLStatusDeploymentTimedOut, CertificatePackCertificateUpdateResponseSSLStatusDeletionTimedOut, CertificatePackCertificateUpdateResponseSSLStatusPendingCleanup, CertificatePackCertificateUpdateResponseSSLStatusStagingDeployment, CertificatePackCertificateUpdateResponseSSLStatusStagingActive, CertificatePackCertificateUpdateResponseSSLStatusDeactivating, CertificatePackCertificateUpdateResponseSSLStatusInactive, CertificatePackCertificateUpdateResponseSSLStatusBackupIssued, CertificatePackCertificateUpdateResponseSSLStatusHoldingDeployment:
		return true
	}
	return false
}

type CertificatePackCertificateUpdateResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                                         `json:"message"`
	JSON    certificatePackCertificateUpdateResponseSSLValidationErrorJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseSSLValidationErrorJSON contains the JSON
// metadata for the struct
// [CertificatePackCertificateUpdateResponseSSLValidationError]
type certificatePackCertificateUpdateResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseSSLValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseSSLValidationRecord struct {
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
	TXTValue string                                                          `json:"txt_value"`
	JSON     certificatePackCertificateUpdateResponseSSLValidationRecordJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseSSLValidationRecordJSON contains the
// JSON metadata for the struct
// [CertificatePackCertificateUpdateResponseSSLValidationRecord]
type certificatePackCertificateUpdateResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseSSLValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CertificatePackCertificateUpdateResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CertificatePackCertificateUpdateResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                            `json:"value"`
	JSON  certificatePackCertificateUpdateResponseOwnershipVerificationJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseOwnershipVerificationJSON contains the
// JSON metadata for the struct
// [CertificatePackCertificateUpdateResponseOwnershipVerification]
type certificatePackCertificateUpdateResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CertificatePackCertificateUpdateResponseOwnershipVerificationType string

const (
	CertificatePackCertificateUpdateResponseOwnershipVerificationTypeTXT CertificatePackCertificateUpdateResponseOwnershipVerificationType = "txt"
)

func (r CertificatePackCertificateUpdateResponseOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseOwnershipVerificationTypeTXT:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CertificatePackCertificateUpdateResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                                `json:"http_url"`
	JSON    certificatePackCertificateUpdateResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseOwnershipVerificationHTTPJSON contains
// the JSON metadata for the struct
// [CertificatePackCertificateUpdateResponseOwnershipVerificationHTTP]
type certificatePackCertificateUpdateResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CertificatePackCertificateUpdateResponseStatus string

const (
	CertificatePackCertificateUpdateResponseStatusActive             CertificatePackCertificateUpdateResponseStatus = "active"
	CertificatePackCertificateUpdateResponseStatusPending            CertificatePackCertificateUpdateResponseStatus = "pending"
	CertificatePackCertificateUpdateResponseStatusActiveRedeploying  CertificatePackCertificateUpdateResponseStatus = "active_redeploying"
	CertificatePackCertificateUpdateResponseStatusMoved              CertificatePackCertificateUpdateResponseStatus = "moved"
	CertificatePackCertificateUpdateResponseStatusPendingDeletion    CertificatePackCertificateUpdateResponseStatus = "pending_deletion"
	CertificatePackCertificateUpdateResponseStatusDeleted            CertificatePackCertificateUpdateResponseStatus = "deleted"
	CertificatePackCertificateUpdateResponseStatusPendingBlocked     CertificatePackCertificateUpdateResponseStatus = "pending_blocked"
	CertificatePackCertificateUpdateResponseStatusPendingMigration   CertificatePackCertificateUpdateResponseStatus = "pending_migration"
	CertificatePackCertificateUpdateResponseStatusPendingProvisioned CertificatePackCertificateUpdateResponseStatus = "pending_provisioned"
	CertificatePackCertificateUpdateResponseStatusTestPending        CertificatePackCertificateUpdateResponseStatus = "test_pending"
	CertificatePackCertificateUpdateResponseStatusTestActive         CertificatePackCertificateUpdateResponseStatus = "test_active"
	CertificatePackCertificateUpdateResponseStatusTestActiveApex     CertificatePackCertificateUpdateResponseStatus = "test_active_apex"
	CertificatePackCertificateUpdateResponseStatusTestBlocked        CertificatePackCertificateUpdateResponseStatus = "test_blocked"
	CertificatePackCertificateUpdateResponseStatusTestFailed         CertificatePackCertificateUpdateResponseStatus = "test_failed"
	CertificatePackCertificateUpdateResponseStatusProvisioned        CertificatePackCertificateUpdateResponseStatus = "provisioned"
	CertificatePackCertificateUpdateResponseStatusBlocked            CertificatePackCertificateUpdateResponseStatus = "blocked"
)

func (r CertificatePackCertificateUpdateResponseStatus) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseStatusActive, CertificatePackCertificateUpdateResponseStatusPending, CertificatePackCertificateUpdateResponseStatusActiveRedeploying, CertificatePackCertificateUpdateResponseStatusMoved, CertificatePackCertificateUpdateResponseStatusPendingDeletion, CertificatePackCertificateUpdateResponseStatusDeleted, CertificatePackCertificateUpdateResponseStatusPendingBlocked, CertificatePackCertificateUpdateResponseStatusPendingMigration, CertificatePackCertificateUpdateResponseStatusPendingProvisioned, CertificatePackCertificateUpdateResponseStatusTestPending, CertificatePackCertificateUpdateResponseStatusTestActive, CertificatePackCertificateUpdateResponseStatusTestActiveApex, CertificatePackCertificateUpdateResponseStatusTestBlocked, CertificatePackCertificateUpdateResponseStatusTestFailed, CertificatePackCertificateUpdateResponseStatusProvisioned, CertificatePackCertificateUpdateResponseStatusBlocked:
		return true
	}
	return false
}

type CertificatePackCertificateDeleteResponse struct {
	// Identifier.
	ID   string                                       `json:"id"`
	JSON certificatePackCertificateDeleteResponseJSON `json:"-"`
}

// certificatePackCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [CertificatePackCertificateDeleteResponse]
type certificatePackCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate,required"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key,required"`
}

func (r CertificatePackCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificatePackCertificateUpdateResponseEnvelope struct {
	Errors   []CertificatePackCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificatePackCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificatePackCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackCertificateUpdateResponse                `json:"result"`
	JSON    certificatePackCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// certificatePackCertificateUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [CertificatePackCertificateUpdateResponseEnvelope]
type certificatePackCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           CertificatePackCertificateUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             certificatePackCertificateUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// certificatePackCertificateUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CertificatePackCertificateUpdateResponseEnvelopeErrors]
type certificatePackCertificateUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    certificatePackCertificateUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CertificatePackCertificateUpdateResponseEnvelopeErrorsSource]
type certificatePackCertificateUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           CertificatePackCertificateUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             certificatePackCertificateUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// certificatePackCertificateUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CertificatePackCertificateUpdateResponseEnvelopeMessages]
type certificatePackCertificateUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CertificatePackCertificateUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    certificatePackCertificateUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// certificatePackCertificateUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CertificatePackCertificateUpdateResponseEnvelopeMessagesSource]
type certificatePackCertificateUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackCertificateUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackCertificateUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificatePackCertificateUpdateResponseEnvelopeSuccess bool

const (
	CertificatePackCertificateUpdateResponseEnvelopeSuccessTrue CertificatePackCertificateUpdateResponseEnvelopeSuccess = true
)

func (r CertificatePackCertificateUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackCertificateUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CertificatePackCertificateDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
