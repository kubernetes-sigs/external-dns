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
)

// VerificationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationService] method instead.
type VerificationService struct {
	Options []option.RequestOption
}

// NewVerificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerificationService(opts ...option.RequestOption) (r *VerificationService) {
	r = &VerificationService{}
	r.Options = opts
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *VerificationService) Edit(ctx context.Context, certificatePackID string, params VerificationEditParams, opts ...option.RequestOption) (res *VerificationEditResponse, err error) {
	var env VerificationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", params.ZoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get SSL Verification Info for a Zone.
func (r *VerificationService) Get(ctx context.Context, params VerificationGetParams, opts ...option.RequestOption) (res *[]Verification, err error) {
	var env VerificationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/verification", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Verification struct {
	// Current status of certificate.
	CertificateStatus VerificationCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CERTPackUUID string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature VerificationSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod ValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo VerificationVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType VerificationVerificationType `json:"verification_type"`
	JSON             verificationJSON             `json:"-"`
}

// verificationJSON contains the JSON metadata for the struct [Verification]
type verificationJSON struct {
	CertificateStatus  apijson.Field
	BrandCheck         apijson.Field
	CERTPackUUID       apijson.Field
	Signature          apijson.Field
	ValidationMethod   apijson.Field
	VerificationInfo   apijson.Field
	VerificationStatus apijson.Field
	VerificationType   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Verification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type VerificationCertificateStatus string

const (
	VerificationCertificateStatusInitializing      VerificationCertificateStatus = "initializing"
	VerificationCertificateStatusAuthorizing       VerificationCertificateStatus = "authorizing"
	VerificationCertificateStatusActive            VerificationCertificateStatus = "active"
	VerificationCertificateStatusExpired           VerificationCertificateStatus = "expired"
	VerificationCertificateStatusIssuing           VerificationCertificateStatus = "issuing"
	VerificationCertificateStatusTimingOut         VerificationCertificateStatus = "timing_out"
	VerificationCertificateStatusPendingDeployment VerificationCertificateStatus = "pending_deployment"
)

func (r VerificationCertificateStatus) IsKnown() bool {
	switch r {
	case VerificationCertificateStatusInitializing, VerificationCertificateStatusAuthorizing, VerificationCertificateStatusActive, VerificationCertificateStatusExpired, VerificationCertificateStatusIssuing, VerificationCertificateStatusTimingOut, VerificationCertificateStatusPendingDeployment:
		return true
	}
	return false
}

// Certificate's signature algorithm.
type VerificationSignature string

const (
	VerificationSignatureEcdsaWithSha256 VerificationSignature = "ECDSAWithSHA256"
	VerificationSignatureSha1WithRSA     VerificationSignature = "SHA1WithRSA"
	VerificationSignatureSha256WithRSA   VerificationSignature = "SHA256WithRSA"
)

func (r VerificationSignature) IsKnown() bool {
	switch r {
	case VerificationSignatureEcdsaWithSha256, VerificationSignatureSha1WithRSA, VerificationSignatureSha256WithRSA:
		return true
	}
	return false
}

// Certificate's required verification information.
type VerificationVerificationInfo struct {
	// Name of CNAME record.
	RecordName VerificationVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget VerificationVerificationInfoRecordTarget `json:"record_target"`
	JSON         verificationVerificationInfoJSON         `json:"-"`
}

// verificationVerificationInfoJSON contains the JSON metadata for the struct
// [VerificationVerificationInfo]
type verificationVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VerificationVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type VerificationVerificationInfoRecordName string

const (
	VerificationVerificationInfoRecordNameRecordName VerificationVerificationInfoRecordName = "record_name"
	VerificationVerificationInfoRecordNameHTTPURL    VerificationVerificationInfoRecordName = "http_url"
	VerificationVerificationInfoRecordNameCNAME      VerificationVerificationInfoRecordName = "cname"
	VerificationVerificationInfoRecordNameTXTName    VerificationVerificationInfoRecordName = "txt_name"
)

func (r VerificationVerificationInfoRecordName) IsKnown() bool {
	switch r {
	case VerificationVerificationInfoRecordNameRecordName, VerificationVerificationInfoRecordNameHTTPURL, VerificationVerificationInfoRecordNameCNAME, VerificationVerificationInfoRecordNameTXTName:
		return true
	}
	return false
}

// Target of CNAME record.
type VerificationVerificationInfoRecordTarget string

const (
	VerificationVerificationInfoRecordTargetRecordValue VerificationVerificationInfoRecordTarget = "record_value"
	VerificationVerificationInfoRecordTargetHTTPBody    VerificationVerificationInfoRecordTarget = "http_body"
	VerificationVerificationInfoRecordTargetCNAMETarget VerificationVerificationInfoRecordTarget = "cname_target"
	VerificationVerificationInfoRecordTargetTXTValue    VerificationVerificationInfoRecordTarget = "txt_value"
)

func (r VerificationVerificationInfoRecordTarget) IsKnown() bool {
	switch r {
	case VerificationVerificationInfoRecordTargetRecordValue, VerificationVerificationInfoRecordTargetHTTPBody, VerificationVerificationInfoRecordTargetCNAMETarget, VerificationVerificationInfoRecordTargetTXTValue:
		return true
	}
	return false
}

// Method of verification.
type VerificationVerificationType string

const (
	VerificationVerificationTypeCNAME   VerificationVerificationType = "cname"
	VerificationVerificationTypeMetaTag VerificationVerificationType = "meta tag"
)

func (r VerificationVerificationType) IsKnown() bool {
	switch r {
	case VerificationVerificationTypeCNAME, VerificationVerificationTypeMetaTag:
		return true
	}
	return false
}

type VerificationEditResponse struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod VerificationEditResponseValidationMethod `json:"validation_method"`
	JSON             verificationEditResponseJSON             `json:"-"`
}

// verificationEditResponseJSON contains the JSON metadata for the struct
// [VerificationEditResponse]
type verificationEditResponseJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VerificationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseJSON) RawJSON() string {
	return r.raw
}

// Desired validation method.
type VerificationEditResponseValidationMethod string

const (
	VerificationEditResponseValidationMethodHTTP  VerificationEditResponseValidationMethod = "http"
	VerificationEditResponseValidationMethodCNAME VerificationEditResponseValidationMethod = "cname"
	VerificationEditResponseValidationMethodTXT   VerificationEditResponseValidationMethod = "txt"
	VerificationEditResponseValidationMethodEmail VerificationEditResponseValidationMethod = "email"
)

func (r VerificationEditResponseValidationMethod) IsKnown() bool {
	switch r {
	case VerificationEditResponseValidationMethodHTTP, VerificationEditResponseValidationMethodCNAME, VerificationEditResponseValidationMethodTXT, VerificationEditResponseValidationMethodEmail:
		return true
	}
	return false
}

type VerificationEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Desired validation method.
	ValidationMethod param.Field[VerificationEditParamsValidationMethod] `json:"validation_method,required"`
}

func (r VerificationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired validation method.
type VerificationEditParamsValidationMethod string

const (
	VerificationEditParamsValidationMethodHTTP  VerificationEditParamsValidationMethod = "http"
	VerificationEditParamsValidationMethodCNAME VerificationEditParamsValidationMethod = "cname"
	VerificationEditParamsValidationMethodTXT   VerificationEditParamsValidationMethod = "txt"
	VerificationEditParamsValidationMethodEmail VerificationEditParamsValidationMethod = "email"
)

func (r VerificationEditParamsValidationMethod) IsKnown() bool {
	switch r {
	case VerificationEditParamsValidationMethodHTTP, VerificationEditParamsValidationMethodCNAME, VerificationEditParamsValidationMethodTXT, VerificationEditParamsValidationMethodEmail:
		return true
	}
	return false
}

type VerificationEditResponseEnvelope struct {
	Errors   []VerificationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VerificationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success VerificationEditResponseEnvelopeSuccess `json:"success,required"`
	Result  VerificationEditResponse                `json:"result"`
	JSON    verificationEditResponseEnvelopeJSON    `json:"-"`
}

// verificationEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [VerificationEditResponseEnvelope]
type verificationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           VerificationEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             verificationEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// verificationEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VerificationEditResponseEnvelopeErrors]
type verificationEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    verificationEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// verificationEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VerificationEditResponseEnvelopeErrorsSource]
type verificationEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           VerificationEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             verificationEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// verificationEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VerificationEditResponseEnvelopeMessages]
type verificationEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type VerificationEditResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    verificationEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// verificationEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [VerificationEditResponseEnvelopeMessagesSource]
type verificationEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VerificationEditResponseEnvelopeSuccess bool

const (
	VerificationEditResponseEnvelopeSuccessTrue VerificationEditResponseEnvelopeSuccess = true
)

func (r VerificationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VerificationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type VerificationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Immediately retry SSL Verification.
	Retry param.Field[VerificationGetParamsRetry] `query:"retry"`
}

// URLQuery serializes [VerificationGetParams]'s query parameters as `url.Values`.
func (r VerificationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Immediately retry SSL Verification.
type VerificationGetParamsRetry bool

const (
	VerificationGetParamsRetryTrue VerificationGetParamsRetry = true
)

func (r VerificationGetParamsRetry) IsKnown() bool {
	switch r {
	case VerificationGetParamsRetryTrue:
		return true
	}
	return false
}

type VerificationGetResponseEnvelope struct {
	Result []Verification                      `json:"result"`
	JSON   verificationGetResponseEnvelopeJSON `json:"-"`
}

// verificationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VerificationGetResponseEnvelope]
type verificationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
