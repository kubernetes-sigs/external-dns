// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

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
)

// CtAuthorityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCtAuthorityService] method instead.
type CtAuthorityService struct {
	Options []option.RequestOption
}

// NewCtAuthorityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCtAuthorityService(opts ...option.RequestOption) (r *CtAuthorityService) {
	r = &CtAuthorityService{}
	r.Options = opts
	return
}

// Retrieves a list of certificate authorities.
func (r *CtAuthorityService) List(ctx context.Context, query CtAuthorityListParams, opts ...option.RequestOption) (res *CtAuthorityListResponse, err error) {
	var env CtAuthorityListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ct/authorities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested CA information.
func (r *CtAuthorityService) Get(ctx context.Context, caSlug string, query CtAuthorityGetParams, opts ...option.RequestOption) (res *CtAuthorityGetResponse, err error) {
	var env CtAuthorityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if caSlug == "" {
		err = errors.New("missing required ca_slug parameter")
		return
	}
	path := fmt.Sprintf("radar/ct/authorities/%s", caSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CtAuthorityListResponse struct {
	CertificateAuthorities []CtAuthorityListResponseCertificateAuthority `json:"certificateAuthorities,required"`
	JSON                   ctAuthorityListResponseJSON                   `json:"-"`
}

// ctAuthorityListResponseJSON contains the JSON metadata for the struct
// [CtAuthorityListResponse]
type ctAuthorityListResponseJSON struct {
	CertificateAuthorities apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CtAuthorityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseJSON) RawJSON() string {
	return r.raw
}

type CtAuthorityListResponseCertificateAuthority struct {
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CtAuthorityListResponseCertificateAuthoritiesCertificateRecordType `json:"certificateRecordType,required"`
	// The two-letter ISO country code where the CA organization is based.
	Country string `json:"country,required"`
	// The full country name corresponding to the country code.
	CountryName string `json:"countryName,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The organization that owns and operates the CA.
	Owner string `json:"owner,required"`
	// The name of the parent/root certificate authority that issued this intermediate
	// certificate.
	ParentName string `json:"parentName,required"`
	// The SHA-256 fingerprint of the parent certificate.
	ParentSha256Fingerprint string `json:"parentSha256Fingerprint,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CtAuthorityListResponseCertificateAuthoritiesRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string                                          `json:"sha256Fingerprint,required"`
	JSON              ctAuthorityListResponseCertificateAuthorityJSON `json:"-"`
}

// ctAuthorityListResponseCertificateAuthorityJSON contains the JSON metadata for
// the struct [CtAuthorityListResponseCertificateAuthority]
type ctAuthorityListResponseCertificateAuthorityJSON struct {
	CertificateRecordType   apijson.Field
	Country                 apijson.Field
	CountryName             apijson.Field
	Name                    apijson.Field
	Owner                   apijson.Field
	ParentName              apijson.Field
	ParentSha256Fingerprint apijson.Field
	RevocationStatus        apijson.Field
	Sha256Fingerprint       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CtAuthorityListResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of certificate in the trust chain.
type CtAuthorityListResponseCertificateAuthoritiesCertificateRecordType string

const (
	CtAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeRootCertificate         CtAuthorityListResponseCertificateAuthoritiesCertificateRecordType = "ROOT_CERTIFICATE"
	CtAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeIntermediateCertificate CtAuthorityListResponseCertificateAuthoritiesCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CtAuthorityListResponseCertificateAuthoritiesCertificateRecordType) IsKnown() bool {
	switch r {
	case CtAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeRootCertificate, CtAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CtAuthorityListResponseCertificateAuthoritiesRevocationStatus string

const (
	CtAuthorityListResponseCertificateAuthoritiesRevocationStatusNotRevoked        CtAuthorityListResponseCertificateAuthoritiesRevocationStatus = "NOT_REVOKED"
	CtAuthorityListResponseCertificateAuthoritiesRevocationStatusRevoked           CtAuthorityListResponseCertificateAuthoritiesRevocationStatus = "REVOKED"
	CtAuthorityListResponseCertificateAuthoritiesRevocationStatusParentCERTRevoked CtAuthorityListResponseCertificateAuthoritiesRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CtAuthorityListResponseCertificateAuthoritiesRevocationStatus) IsKnown() bool {
	switch r {
	case CtAuthorityListResponseCertificateAuthoritiesRevocationStatusNotRevoked, CtAuthorityListResponseCertificateAuthoritiesRevocationStatusRevoked, CtAuthorityListResponseCertificateAuthoritiesRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

type CtAuthorityGetResponse struct {
	CertificateAuthority CtAuthorityGetResponseCertificateAuthority `json:"certificateAuthority,required"`
	JSON                 ctAuthorityGetResponseJSON                 `json:"-"`
}

// ctAuthorityGetResponseJSON contains the JSON metadata for the struct
// [CtAuthorityGetResponse]
type ctAuthorityGetResponseJSON struct {
	CertificateAuthority apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CtAuthorityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseJSON) RawJSON() string {
	return r.raw
}

type CtAuthorityGetResponseCertificateAuthority struct {
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	AppleStatus CtAuthorityGetResponseCertificateAuthorityAppleStatus `json:"appleStatus,required"`
	// The authorityKeyIdentifier value extracted from the certificate PEM.
	AuthorityKeyIdentifier string `json:"authorityKeyIdentifier,required"`
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CtAuthorityGetResponseCertificateAuthorityCertificateRecordType `json:"certificateRecordType,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	ChromeStatus CtAuthorityGetResponseCertificateAuthorityChromeStatus `json:"chromeStatus,required"`
	// The two-letter ISO country code where the CA organization is based.
	Country string `json:"country,required"`
	// The full country name corresponding to the country code.
	CountryName string `json:"countryName,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	MicrosoftStatus CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus `json:"microsoftStatus,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	MozillaStatus CtAuthorityGetResponseCertificateAuthorityMozillaStatus `json:"mozillaStatus,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The organization that owns and operates the CA.
	Owner string `json:"owner,required"`
	// The name of the parent/root certificate authority that issued this intermediate
	// certificate.
	ParentName string `json:"parentName,required"`
	// The SHA-256 fingerprint of the parent certificate.
	ParentSha256Fingerprint string `json:"parentSha256Fingerprint,required"`
	// CAs from the same owner.
	Related []CtAuthorityGetResponseCertificateAuthorityRelated `json:"related,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CtAuthorityGetResponseCertificateAuthorityRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string `json:"sha256Fingerprint,required"`
	// The subjectKeyIdentifier value extracted from the certificate PEM.
	SubjectKeyIdentifier string `json:"subjectKeyIdentifier,required"`
	// The start date of the certificate’s validity period (ISO format).
	ValidFrom string `json:"validFrom,required"`
	// The end date of the certificate’s validity period (ISO format).
	ValidTo string                                         `json:"validTo,required"`
	JSON    ctAuthorityGetResponseCertificateAuthorityJSON `json:"-"`
}

// ctAuthorityGetResponseCertificateAuthorityJSON contains the JSON metadata for
// the struct [CtAuthorityGetResponseCertificateAuthority]
type ctAuthorityGetResponseCertificateAuthorityJSON struct {
	AppleStatus             apijson.Field
	AuthorityKeyIdentifier  apijson.Field
	CertificateRecordType   apijson.Field
	ChromeStatus            apijson.Field
	Country                 apijson.Field
	CountryName             apijson.Field
	MicrosoftStatus         apijson.Field
	MozillaStatus           apijson.Field
	Name                    apijson.Field
	Owner                   apijson.Field
	ParentName              apijson.Field
	ParentSha256Fingerprint apijson.Field
	Related                 apijson.Field
	RevocationStatus        apijson.Field
	Sha256Fingerprint       apijson.Field
	SubjectKeyIdentifier    apijson.Field
	ValidFrom               apijson.Field
	ValidTo                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CtAuthorityGetResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CtAuthorityGetResponseCertificateAuthorityAppleStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityAppleStatusIncluded       CtAuthorityGetResponseCertificateAuthorityAppleStatus = "INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusNotYetIncluded CtAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_YET_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusNotIncluded    CtAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusNotBefore      CtAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_BEFORE"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusRemoved        CtAuthorityGetResponseCertificateAuthorityAppleStatus = "REMOVED"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusDisabled       CtAuthorityGetResponseCertificateAuthorityAppleStatus = "DISABLED"
	CtAuthorityGetResponseCertificateAuthorityAppleStatusBlocked        CtAuthorityGetResponseCertificateAuthorityAppleStatus = "BLOCKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityAppleStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityAppleStatusIncluded, CtAuthorityGetResponseCertificateAuthorityAppleStatusNotYetIncluded, CtAuthorityGetResponseCertificateAuthorityAppleStatusNotIncluded, CtAuthorityGetResponseCertificateAuthorityAppleStatusNotBefore, CtAuthorityGetResponseCertificateAuthorityAppleStatusRemoved, CtAuthorityGetResponseCertificateAuthorityAppleStatusDisabled, CtAuthorityGetResponseCertificateAuthorityAppleStatusBlocked:
		return true
	}
	return false
}

// Specifies the type of certificate in the trust chain.
type CtAuthorityGetResponseCertificateAuthorityCertificateRecordType string

const (
	CtAuthorityGetResponseCertificateAuthorityCertificateRecordTypeRootCertificate         CtAuthorityGetResponseCertificateAuthorityCertificateRecordType = "ROOT_CERTIFICATE"
	CtAuthorityGetResponseCertificateAuthorityCertificateRecordTypeIntermediateCertificate CtAuthorityGetResponseCertificateAuthorityCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CtAuthorityGetResponseCertificateAuthorityCertificateRecordType) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityCertificateRecordTypeRootCertificate, CtAuthorityGetResponseCertificateAuthorityCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CtAuthorityGetResponseCertificateAuthorityChromeStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityChromeStatusIncluded       CtAuthorityGetResponseCertificateAuthorityChromeStatus = "INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusNotYetIncluded CtAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_YET_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusNotIncluded    CtAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusNotBefore      CtAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_BEFORE"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusRemoved        CtAuthorityGetResponseCertificateAuthorityChromeStatus = "REMOVED"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusDisabled       CtAuthorityGetResponseCertificateAuthorityChromeStatus = "DISABLED"
	CtAuthorityGetResponseCertificateAuthorityChromeStatusBlocked        CtAuthorityGetResponseCertificateAuthorityChromeStatus = "BLOCKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityChromeStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityChromeStatusIncluded, CtAuthorityGetResponseCertificateAuthorityChromeStatusNotYetIncluded, CtAuthorityGetResponseCertificateAuthorityChromeStatusNotIncluded, CtAuthorityGetResponseCertificateAuthorityChromeStatusNotBefore, CtAuthorityGetResponseCertificateAuthorityChromeStatusRemoved, CtAuthorityGetResponseCertificateAuthorityChromeStatusDisabled, CtAuthorityGetResponseCertificateAuthorityChromeStatusBlocked:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusIncluded       CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotYetIncluded CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_YET_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotIncluded    CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotBefore      CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_BEFORE"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusRemoved        CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "REMOVED"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusDisabled       CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "DISABLED"
	CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusBlocked        CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "BLOCKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityMicrosoftStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusIncluded, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotYetIncluded, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotIncluded, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotBefore, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusRemoved, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusDisabled, CtAuthorityGetResponseCertificateAuthorityMicrosoftStatusBlocked:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CtAuthorityGetResponseCertificateAuthorityMozillaStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusIncluded       CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotYetIncluded CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_YET_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotIncluded    CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_INCLUDED"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotBefore      CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_BEFORE"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusRemoved        CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "REMOVED"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusDisabled       CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "DISABLED"
	CtAuthorityGetResponseCertificateAuthorityMozillaStatusBlocked        CtAuthorityGetResponseCertificateAuthorityMozillaStatus = "BLOCKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityMozillaStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityMozillaStatusIncluded, CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotYetIncluded, CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotIncluded, CtAuthorityGetResponseCertificateAuthorityMozillaStatusNotBefore, CtAuthorityGetResponseCertificateAuthorityMozillaStatusRemoved, CtAuthorityGetResponseCertificateAuthorityMozillaStatusDisabled, CtAuthorityGetResponseCertificateAuthorityMozillaStatusBlocked:
		return true
	}
	return false
}

type CtAuthorityGetResponseCertificateAuthorityRelated struct {
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType `json:"certificateRecordType,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string                                                `json:"sha256Fingerprint,required"`
	JSON              ctAuthorityGetResponseCertificateAuthorityRelatedJSON `json:"-"`
}

// ctAuthorityGetResponseCertificateAuthorityRelatedJSON contains the JSON metadata
// for the struct [CtAuthorityGetResponseCertificateAuthorityRelated]
type ctAuthorityGetResponseCertificateAuthorityRelatedJSON struct {
	CertificateRecordType apijson.Field
	Name                  apijson.Field
	RevocationStatus      apijson.Field
	Sha256Fingerprint     apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CtAuthorityGetResponseCertificateAuthorityRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseCertificateAuthorityRelatedJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of certificate in the trust chain.
type CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType string

const (
	CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeRootCertificate         CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType = "ROOT_CERTIFICATE"
	CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeIntermediateCertificate CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeRootCertificate, CtAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusNotRevoked        CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "NOT_REVOKED"
	CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusRevoked           CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "REVOKED"
	CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusParentCERTRevoked CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusNotRevoked, CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusRevoked, CtAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CtAuthorityGetResponseCertificateAuthorityRevocationStatus string

const (
	CtAuthorityGetResponseCertificateAuthorityRevocationStatusNotRevoked        CtAuthorityGetResponseCertificateAuthorityRevocationStatus = "NOT_REVOKED"
	CtAuthorityGetResponseCertificateAuthorityRevocationStatusRevoked           CtAuthorityGetResponseCertificateAuthorityRevocationStatus = "REVOKED"
	CtAuthorityGetResponseCertificateAuthorityRevocationStatusParentCERTRevoked CtAuthorityGetResponseCertificateAuthorityRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CtAuthorityGetResponseCertificateAuthorityRevocationStatus) IsKnown() bool {
	switch r {
	case CtAuthorityGetResponseCertificateAuthorityRevocationStatusNotRevoked, CtAuthorityGetResponseCertificateAuthorityRevocationStatusRevoked, CtAuthorityGetResponseCertificateAuthorityRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

type CtAuthorityListParams struct {
	// Format in which results will be returned.
	Format param.Field[CtAuthorityListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CtAuthorityListParams]'s query parameters as `url.Values`.
func (r CtAuthorityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CtAuthorityListParamsFormat string

const (
	CtAuthorityListParamsFormatJson CtAuthorityListParamsFormat = "JSON"
	CtAuthorityListParamsFormatCsv  CtAuthorityListParamsFormat = "CSV"
)

func (r CtAuthorityListParamsFormat) IsKnown() bool {
	switch r {
	case CtAuthorityListParamsFormatJson, CtAuthorityListParamsFormatCsv:
		return true
	}
	return false
}

type CtAuthorityListResponseEnvelope struct {
	Result  CtAuthorityListResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    ctAuthorityListResponseEnvelopeJSON `json:"-"`
}

// ctAuthorityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtAuthorityListResponseEnvelope]
type ctAuthorityListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtAuthorityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CtAuthorityGetParams struct {
	// Format in which results will be returned.
	Format param.Field[CtAuthorityGetParamsFormat] `query:"format"`
}

// URLQuery serializes [CtAuthorityGetParams]'s query parameters as `url.Values`.
func (r CtAuthorityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CtAuthorityGetParamsFormat string

const (
	CtAuthorityGetParamsFormatJson CtAuthorityGetParamsFormat = "JSON"
	CtAuthorityGetParamsFormatCsv  CtAuthorityGetParamsFormat = "CSV"
)

func (r CtAuthorityGetParamsFormat) IsKnown() bool {
	switch r {
	case CtAuthorityGetParamsFormatJson, CtAuthorityGetParamsFormatCsv:
		return true
	}
	return false
}

type CtAuthorityGetResponseEnvelope struct {
	Result  CtAuthorityGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    ctAuthorityGetResponseEnvelopeJSON `json:"-"`
}

// ctAuthorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtAuthorityGetResponseEnvelope]
type ctAuthorityGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtAuthorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
