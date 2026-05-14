// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities

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

// HostnameAssociationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameAssociationService] method instead.
type HostnameAssociationService struct {
	Options []option.RequestOption
}

// NewHostnameAssociationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHostnameAssociationService(opts ...option.RequestOption) (r *HostnameAssociationService) {
	r = &HostnameAssociationService{}
	r.Options = opts
	return
}

// Replace Hostname Associations
func (r *HostnameAssociationService) Update(ctx context.Context, params HostnameAssociationUpdateParams, opts ...option.RequestOption) (res *HostnameAssociationUpdateResponse, err error) {
	var env HostnameAssociationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Hostname Associations
func (r *HostnameAssociationService) Get(ctx context.Context, params HostnameAssociationGetParams, opts ...option.RequestOption) (res *HostnameAssociationGetResponse, err error) {
	var env HostnameAssociationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameAssociation = string

type HostnameAssociationParam = string

type TLSHostnameAssociationParam struct {
	Hostnames param.Field[[]HostnameAssociationParam] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
}

func (r TLSHostnameAssociationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameAssociationUpdateResponse struct {
	Hostnames []HostnameAssociation                 `json:"hostnames"`
	JSON      hostnameAssociationUpdateResponseJSON `json:"-"`
}

// hostnameAssociationUpdateResponseJSON contains the JSON metadata for the struct
// [HostnameAssociationUpdateResponse]
type hostnameAssociationUpdateResponseJSON struct {
	Hostnames   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationGetResponse struct {
	Hostnames []HostnameAssociation              `json:"hostnames"`
	JSON      hostnameAssociationGetResponseJSON `json:"-"`
}

// hostnameAssociationGetResponseJSON contains the JSON metadata for the struct
// [HostnameAssociationGetResponse]
type hostnameAssociationGetResponseJSON struct {
	Hostnames   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationUpdateParams struct {
	// Identifier.
	ZoneID                 param.Field[string]         `path:"zone_id,required"`
	TLSHostnameAssociation TLSHostnameAssociationParam `json:"tls_hostname_association,required"`
}

func (r HostnameAssociationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TLSHostnameAssociation)
}

type HostnameAssociationUpdateResponseEnvelope struct {
	Errors   []HostnameAssociationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameAssociationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameAssociationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  HostnameAssociationUpdateResponse                `json:"result"`
	JSON    hostnameAssociationUpdateResponseEnvelopeJSON    `json:"-"`
}

// hostnameAssociationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameAssociationUpdateResponseEnvelope]
type hostnameAssociationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           HostnameAssociationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             hostnameAssociationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// hostnameAssociationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HostnameAssociationUpdateResponseEnvelopeErrors]
type hostnameAssociationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    hostnameAssociationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// hostnameAssociationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [HostnameAssociationUpdateResponseEnvelopeErrorsSource]
type hostnameAssociationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           HostnameAssociationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             hostnameAssociationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// hostnameAssociationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameAssociationUpdateResponseEnvelopeMessages]
type hostnameAssociationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    hostnameAssociationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// hostnameAssociationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [HostnameAssociationUpdateResponseEnvelopeMessagesSource]
type hostnameAssociationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameAssociationUpdateResponseEnvelopeSuccess bool

const (
	HostnameAssociationUpdateResponseEnvelopeSuccessTrue HostnameAssociationUpdateResponseEnvelopeSuccess = true
)

func (r HostnameAssociationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameAssociationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameAssociationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MTLSCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes [HostnameAssociationGetParams]'s query parameters as
// `url.Values`.
func (r HostnameAssociationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HostnameAssociationGetResponseEnvelope struct {
	Errors   []HostnameAssociationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameAssociationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameAssociationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  HostnameAssociationGetResponse                `json:"result"`
	JSON    hostnameAssociationGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameAssociationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameAssociationGetResponseEnvelope]
type hostnameAssociationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           HostnameAssociationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             hostnameAssociationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// hostnameAssociationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameAssociationGetResponseEnvelopeErrors]
type hostnameAssociationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameAssociationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    hostnameAssociationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// hostnameAssociationGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [HostnameAssociationGetResponseEnvelopeErrorsSource]
type hostnameAssociationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           HostnameAssociationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             hostnameAssociationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// hostnameAssociationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameAssociationGetResponseEnvelopeMessages]
type hostnameAssociationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameAssociationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HostnameAssociationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    hostnameAssociationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// hostnameAssociationGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [HostnameAssociationGetResponseEnvelopeMessagesSource]
type hostnameAssociationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameAssociationGetResponseEnvelopeSuccess bool

const (
	HostnameAssociationGetResponseEnvelopeSuccessTrue HostnameAssociationGetResponseEnvelopeSuccess = true
)

func (r HostnameAssociationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameAssociationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
