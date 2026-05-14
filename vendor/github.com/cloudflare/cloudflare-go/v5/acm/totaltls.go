// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acm

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TotalTLSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTotalTLSService] method instead.
type TotalTLSService struct {
	Options []option.RequestOption
}

// NewTotalTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTotalTLSService(opts ...option.RequestOption) (r *TotalTLSService) {
	r = &TotalTLSService{}
	r.Options = opts
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *TotalTLSService) New(ctx context.Context, params TotalTLSNewParams, opts ...option.RequestOption) (res *TotalTLSNewResponse, err error) {
	var env TotalTLSNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/total_tls", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Total TLS Settings for a Zone.
func (r *TotalTLSService) Get(ctx context.Context, query TotalTLSGetParams, opts ...option.RequestOption) (res *TotalTLSGetResponse, err error) {
	var env TotalTLSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/total_tls", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The Certificate Authority that Total TLS certificates will be issued through.
type CertificateAuthority string

const (
	CertificateAuthorityGoogle      CertificateAuthority = "google"
	CertificateAuthorityLetsEncrypt CertificateAuthority = "lets_encrypt"
	CertificateAuthoritySSLCom      CertificateAuthority = "ssl_com"
)

func (r CertificateAuthority) IsKnown() bool {
	switch r {
	case CertificateAuthorityGoogle, CertificateAuthorityLetsEncrypt, CertificateAuthoritySSLCom:
		return true
	}
	return false
}

type TotalTLSNewResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityPeriod TotalTLSNewResponseValidityPeriod `json:"validity_period"`
	JSON           totalTLSNewResponseJSON           `json:"-"`
}

// totalTLSNewResponseJSON contains the JSON metadata for the struct
// [TotalTLSNewResponse]
type totalTLSNewResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSNewResponseValidityPeriod int64

const (
	TotalTLSNewResponseValidityPeriod90 TotalTLSNewResponseValidityPeriod = 90
)

func (r TotalTLSNewResponseValidityPeriod) IsKnown() bool {
	switch r {
	case TotalTLSNewResponseValidityPeriod90:
		return true
	}
	return false
}

type TotalTLSGetResponse struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthority `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityPeriod TotalTLSGetResponseValidityPeriod `json:"validity_period"`
	JSON           totalTLSGetResponseJSON           `json:"-"`
}

// totalTLSGetResponseJSON contains the JSON metadata for the struct
// [TotalTLSGetResponse]
type totalTLSGetResponseJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TotalTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type TotalTLSGetResponseValidityPeriod int64

const (
	TotalTLSGetResponseValidityPeriod90 TotalTLSGetResponseValidityPeriod = 90
)

func (r TotalTLSGetResponseValidityPeriod) IsKnown() bool {
	switch r {
	case TotalTLSGetResponseValidityPeriod90:
		return true
	}
	return false
}

type TotalTLSNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[CertificateAuthority] `json:"certificate_authority"`
}

func (r TotalTLSNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TotalTLSNewResponseEnvelope struct {
	Errors   []TotalTLSNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TotalTLSNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TotalTLSNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSNewResponse                `json:"result"`
	JSON    totalTLSNewResponseEnvelopeJSON    `json:"-"`
}

// totalTLSNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSNewResponseEnvelope]
type totalTLSNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TotalTLSNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TotalTLSNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             totalTLSNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// totalTLSNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TotalTLSNewResponseEnvelopeErrors]
type totalTLSNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TotalTLSNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    totalTLSNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// totalTLSNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TotalTLSNewResponseEnvelopeErrorsSource]
type totalTLSNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TotalTLSNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TotalTLSNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             totalTLSNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// totalTLSNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TotalTLSNewResponseEnvelopeMessages]
type totalTLSNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TotalTLSNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    totalTLSNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// totalTLSNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TotalTLSNewResponseEnvelopeMessagesSource]
type totalTLSNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TotalTLSNewResponseEnvelopeSuccess bool

const (
	TotalTLSNewResponseEnvelopeSuccessTrue TotalTLSNewResponseEnvelopeSuccess = true
)

func (r TotalTLSNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TotalTLSGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TotalTLSGetResponseEnvelope struct {
	Errors   []TotalTLSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TotalTLSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TotalTLSGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TotalTLSGetResponse                `json:"result"`
	JSON    totalTLSGetResponseEnvelopeJSON    `json:"-"`
}

// totalTLSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TotalTLSGetResponseEnvelope]
type totalTLSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TotalTLSGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             totalTLSGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// totalTLSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TotalTLSGetResponseEnvelopeErrors]
type totalTLSGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    totalTLSGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// totalTLSGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeErrorsSource]
type totalTLSGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TotalTLSGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             totalTLSGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// totalTLSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeMessages]
type totalTLSGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TotalTLSGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    totalTLSGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// totalTLSGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TotalTLSGetResponseEnvelopeMessagesSource]
type totalTLSGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalTLSGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalTLSGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TotalTLSGetResponseEnvelopeSuccess bool

const (
	TotalTLSGetResponseEnvelopeSuccessTrue TotalTLSGetResponseEnvelopeSuccess = true
)

func (r TotalTLSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TotalTLSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
