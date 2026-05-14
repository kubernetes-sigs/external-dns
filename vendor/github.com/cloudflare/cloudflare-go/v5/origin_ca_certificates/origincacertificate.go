// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificates

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
	"github.com/cloudflare/cloudflare-go/v5/ssl"
)

// OriginCACertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginCACertificateService] method instead.
type OriginCACertificateService struct {
	Options []option.RequestOption
}

// NewOriginCACertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginCACertificateService(opts ...option.RequestOption) (r *OriginCACertificateService) {
	r = &OriginCACertificateService{}
	r.Options = opts
	return
}

// Create an Origin CA certificate. You can use an Origin CA Key as your User
// Service Key or an API token when calling this endpoint ([see above](#requests)).
func (r *OriginCACertificateService) New(ctx context.Context, body OriginCACertificateNewParams, opts ...option.RequestOption) (res *OriginCACertificate, err error) {
	var env OriginCACertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all existing Origin CA certificates for a given zone. You can use an Origin
// CA Key as your User Service Key or an API token when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) List(ctx context.Context, query OriginCACertificateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[OriginCACertificate], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "certificates"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List all existing Origin CA certificates for a given zone. You can use an Origin
// CA Key as your User Service Key or an API token when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) ListAutoPaging(ctx context.Context, query OriginCACertificateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[OriginCACertificate] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Revoke an existing Origin CA certificate by its serial number. You can use an
// Origin CA Key as your User Service Key or an API token when calling this
// endpoint ([see above](#requests)).
func (r *OriginCACertificateService) Delete(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *OriginCACertificateDeleteResponse, err error) {
	var env OriginCACertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an existing Origin CA certificate by its serial number. You can use an
// Origin CA Key as your User Service Key or an API token when calling this
// endpoint ([see above](#requests)).
func (r *OriginCACertificateService) Get(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *OriginCACertificate, err error) {
	var env OriginCACertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginCACertificate struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []string `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType shared.CertificateRequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity ssl.RequestValidity `json:"requested_validity,required"`
	// Identifier.
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn string                  `json:"expires_on"`
	JSON      originCACertificateJSON `json:"-"`
}

// originCACertificateJSON contains the JSON metadata for the struct
// [OriginCACertificate]
type originCACertificateJSON struct {
	Csr               apijson.Field
	Hostnames         apijson.Field
	RequestType       apijson.Field
	RequestedValidity apijson.Field
	ID                apijson.Field
	Certificate       apijson.Field
	ExpiresOn         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OriginCACertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateDeleteResponse struct {
	// Identifier.
	ID string `json:"id"`
	// When the certificate was revoked.
	RevokedAt time.Time                             `json:"revoked_at" format:"date-time"`
	JSON      originCACertificateDeleteResponseJSON `json:"-"`
}

// originCACertificateDeleteResponseJSON contains the JSON metadata for the struct
// [OriginCACertificateDeleteResponse]
type originCACertificateDeleteResponseJSON struct {
	ID          apijson.Field
	RevokedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType param.Field[shared.CertificateRequestType] `json:"request_type"`
	// The number of days for which the certificate should be valid.
	RequestedValidity param.Field[ssl.RequestValidity] `json:"requested_validity"`
}

func (r OriginCACertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginCACertificateNewResponseEnvelope struct {
	Errors   []OriginCACertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginCACertificateNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OriginCACertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginCACertificate                           `json:"result"`
	JSON    originCACertificateNewResponseEnvelopeJSON    `json:"-"`
}

// originCACertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateNewResponseEnvelope]
type originCACertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           OriginCACertificateNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             originCACertificateNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// originCACertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginCACertificateNewResponseEnvelopeErrors]
type originCACertificateNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    originCACertificateNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// originCACertificateNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OriginCACertificateNewResponseEnvelopeErrorsSource]
type originCACertificateNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           OriginCACertificateNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             originCACertificateNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// originCACertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginCACertificateNewResponseEnvelopeMessages]
type originCACertificateNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    originCACertificateNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// originCACertificateNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OriginCACertificateNewResponseEnvelopeMessagesSource]
type originCACertificateNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCACertificateNewResponseEnvelopeSuccess bool

const (
	OriginCACertificateNewResponseEnvelopeSuccessTrue OriginCACertificateNewResponseEnvelopeSuccess = true
)

func (r OriginCACertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCACertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCACertificateListParams struct {
	// Identifier.
	ZoneID param.Field[string] `query:"zone_id,required"`
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [OriginCACertificateListParams]'s query parameters as
// `url.Values`.
func (r OriginCACertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OriginCACertificateDeleteResponseEnvelope struct {
	Result OriginCACertificateDeleteResponse             `json:"result"`
	JSON   originCACertificateDeleteResponseEnvelopeJSON `json:"-"`
}

// originCACertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateDeleteResponseEnvelope]
type originCACertificateDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateGetResponseEnvelope struct {
	Errors   []OriginCACertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginCACertificateGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OriginCACertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginCACertificate                           `json:"result"`
	JSON    originCACertificateGetResponseEnvelopeJSON    `json:"-"`
}

// originCACertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateGetResponseEnvelope]
type originCACertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           OriginCACertificateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             originCACertificateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// originCACertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginCACertificateGetResponseEnvelopeErrors]
type originCACertificateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    originCACertificateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// originCACertificateGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OriginCACertificateGetResponseEnvelopeErrorsSource]
type originCACertificateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           OriginCACertificateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             originCACertificateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// originCACertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginCACertificateGetResponseEnvelopeMessages]
type originCACertificateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OriginCACertificateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    originCACertificateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// originCACertificateGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OriginCACertificateGetResponseEnvelopeMessagesSource]
type originCACertificateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCACertificateGetResponseEnvelopeSuccess bool

const (
	OriginCACertificateGetResponseEnvelopeSuccessTrue OriginCACertificateGetResponseEnvelopeSuccess = true
)

func (r OriginCACertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCACertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
