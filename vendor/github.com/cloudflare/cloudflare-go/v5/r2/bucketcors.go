// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// BucketCORSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketCORSService] method instead.
type BucketCORSService struct {
	Options []option.RequestOption
}

// NewBucketCORSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketCORSService(opts ...option.RequestOption) (r *BucketCORSService) {
	r = &BucketCORSService{}
	r.Options = opts
	return
}

// Set the CORS policy for a bucket.
func (r *BucketCORSService) Update(ctx context.Context, bucketName string, params BucketCORSUpdateParams, opts ...option.RequestOption) (res *BucketCORSUpdateResponse, err error) {
	var env BucketCORSUpdateResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the CORS policy for a bucket.
func (r *BucketCORSService) Delete(ctx context.Context, bucketName string, params BucketCORSDeleteParams, opts ...option.RequestOption) (res *BucketCORSDeleteResponse, err error) {
	var env BucketCORSDeleteResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the CORS policy for a bucket.
func (r *BucketCORSService) Get(ctx context.Context, bucketName string, params BucketCORSGetParams, opts ...option.RequestOption) (res *BucketCORSGetResponse, err error) {
	var env BucketCORSGetResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/cors", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketCORSUpdateResponse = interface{}

type BucketCORSDeleteResponse = interface{}

type BucketCORSGetResponse struct {
	Rules []BucketCORSGetResponseRule `json:"rules"`
	JSON  bucketCORSGetResponseJSON   `json:"-"`
}

// bucketCORSGetResponseJSON contains the JSON metadata for the struct
// [BucketCORSGetResponse]
type bucketCORSGetResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketCORSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketCORSGetResponseRule struct {
	// Object specifying allowed origins, methods and headers for this CORS rule.
	Allowed BucketCORSGetResponseRulesAllowed `json:"allowed,required"`
	// Identifier for this rule.
	ID string `json:"id"`
	// Specifies the headers that can be exposed back, and accessed by, the JavaScript
	// making the cross-origin request. If you need to access headers beyond the
	// safelisted response headers, such as Content-Encoding or cf-cache-status, you
	// must specify it here.
	ExposeHeaders []string `json:"exposeHeaders"`
	// Specifies the amount of time (in seconds) browsers are allowed to cache CORS
	// preflight responses. Browsers may limit this to 2 hours or less, even if the
	// maximum value (86400) is specified.
	MaxAgeSeconds float64                       `json:"maxAgeSeconds"`
	JSON          bucketCORSGetResponseRuleJSON `json:"-"`
}

// bucketCORSGetResponseRuleJSON contains the JSON metadata for the struct
// [BucketCORSGetResponseRule]
type bucketCORSGetResponseRuleJSON struct {
	Allowed       apijson.Field
	ID            apijson.Field
	ExposeHeaders apijson.Field
	MaxAgeSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BucketCORSGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

// Object specifying allowed origins, methods and headers for this CORS rule.
type BucketCORSGetResponseRulesAllowed struct {
	// Specifies the value for the Access-Control-Allow-Methods header R2 sets when
	// requesting objects in a bucket from a browser.
	Methods []BucketCORSGetResponseRulesAllowedMethod `json:"methods,required"`
	// Specifies the value for the Access-Control-Allow-Origin header R2 sets when
	// requesting objects in a bucket from a browser.
	Origins []string `json:"origins,required"`
	// Specifies the value for the Access-Control-Allow-Headers header R2 sets when
	// requesting objects in this bucket from a browser. Cross-origin requests that
	// include custom headers (e.g. x-user-id) should specify these headers as
	// AllowedHeaders.
	Headers []string                              `json:"headers"`
	JSON    bucketCORSGetResponseRulesAllowedJSON `json:"-"`
}

// bucketCORSGetResponseRulesAllowedJSON contains the JSON metadata for the struct
// [BucketCORSGetResponseRulesAllowed]
type bucketCORSGetResponseRulesAllowedJSON struct {
	Methods     apijson.Field
	Origins     apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketCORSGetResponseRulesAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSGetResponseRulesAllowedJSON) RawJSON() string {
	return r.raw
}

type BucketCORSGetResponseRulesAllowedMethod string

const (
	BucketCORSGetResponseRulesAllowedMethodGet    BucketCORSGetResponseRulesAllowedMethod = "GET"
	BucketCORSGetResponseRulesAllowedMethodPut    BucketCORSGetResponseRulesAllowedMethod = "PUT"
	BucketCORSGetResponseRulesAllowedMethodPost   BucketCORSGetResponseRulesAllowedMethod = "POST"
	BucketCORSGetResponseRulesAllowedMethodDelete BucketCORSGetResponseRulesAllowedMethod = "DELETE"
	BucketCORSGetResponseRulesAllowedMethodHead   BucketCORSGetResponseRulesAllowedMethod = "HEAD"
)

func (r BucketCORSGetResponseRulesAllowedMethod) IsKnown() bool {
	switch r {
	case BucketCORSGetResponseRulesAllowedMethodGet, BucketCORSGetResponseRulesAllowedMethodPut, BucketCORSGetResponseRulesAllowedMethodPost, BucketCORSGetResponseRulesAllowedMethodDelete, BucketCORSGetResponseRulesAllowedMethodHead:
		return true
	}
	return false
}

type BucketCORSUpdateParams struct {
	// Account ID.
	AccountID param.Field[string]                       `path:"account_id,required"`
	Rules     param.Field[[]BucketCORSUpdateParamsRule] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketCORSUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketCORSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketCORSUpdateParamsRule struct {
	// Object specifying allowed origins, methods and headers for this CORS rule.
	Allowed param.Field[BucketCORSUpdateParamsRulesAllowed] `json:"allowed,required"`
	// Identifier for this rule.
	ID param.Field[string] `json:"id"`
	// Specifies the headers that can be exposed back, and accessed by, the JavaScript
	// making the cross-origin request. If you need to access headers beyond the
	// safelisted response headers, such as Content-Encoding or cf-cache-status, you
	// must specify it here.
	ExposeHeaders param.Field[[]string] `json:"exposeHeaders"`
	// Specifies the amount of time (in seconds) browsers are allowed to cache CORS
	// preflight responses. Browsers may limit this to 2 hours or less, even if the
	// maximum value (86400) is specified.
	MaxAgeSeconds param.Field[float64] `json:"maxAgeSeconds"`
}

func (r BucketCORSUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object specifying allowed origins, methods and headers for this CORS rule.
type BucketCORSUpdateParamsRulesAllowed struct {
	// Specifies the value for the Access-Control-Allow-Methods header R2 sets when
	// requesting objects in a bucket from a browser.
	Methods param.Field[[]BucketCORSUpdateParamsRulesAllowedMethod] `json:"methods,required"`
	// Specifies the value for the Access-Control-Allow-Origin header R2 sets when
	// requesting objects in a bucket from a browser.
	Origins param.Field[[]string] `json:"origins,required"`
	// Specifies the value for the Access-Control-Allow-Headers header R2 sets when
	// requesting objects in this bucket from a browser. Cross-origin requests that
	// include custom headers (e.g. x-user-id) should specify these headers as
	// AllowedHeaders.
	Headers param.Field[[]string] `json:"headers"`
}

func (r BucketCORSUpdateParamsRulesAllowed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketCORSUpdateParamsRulesAllowedMethod string

const (
	BucketCORSUpdateParamsRulesAllowedMethodGet    BucketCORSUpdateParamsRulesAllowedMethod = "GET"
	BucketCORSUpdateParamsRulesAllowedMethodPut    BucketCORSUpdateParamsRulesAllowedMethod = "PUT"
	BucketCORSUpdateParamsRulesAllowedMethodPost   BucketCORSUpdateParamsRulesAllowedMethod = "POST"
	BucketCORSUpdateParamsRulesAllowedMethodDelete BucketCORSUpdateParamsRulesAllowedMethod = "DELETE"
	BucketCORSUpdateParamsRulesAllowedMethodHead   BucketCORSUpdateParamsRulesAllowedMethod = "HEAD"
)

func (r BucketCORSUpdateParamsRulesAllowedMethod) IsKnown() bool {
	switch r {
	case BucketCORSUpdateParamsRulesAllowedMethodGet, BucketCORSUpdateParamsRulesAllowedMethodPut, BucketCORSUpdateParamsRulesAllowedMethodPost, BucketCORSUpdateParamsRulesAllowedMethodDelete, BucketCORSUpdateParamsRulesAllowedMethodHead:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketCORSUpdateParamsCfR2Jurisdiction string

const (
	BucketCORSUpdateParamsCfR2JurisdictionDefault BucketCORSUpdateParamsCfR2Jurisdiction = "default"
	BucketCORSUpdateParamsCfR2JurisdictionEu      BucketCORSUpdateParamsCfR2Jurisdiction = "eu"
	BucketCORSUpdateParamsCfR2JurisdictionFedramp BucketCORSUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketCORSUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketCORSUpdateParamsCfR2JurisdictionDefault, BucketCORSUpdateParamsCfR2JurisdictionEu, BucketCORSUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketCORSUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []string                 `json:"messages,required"`
	Result   BucketCORSUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketCORSUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketCORSUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketCORSUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketCORSUpdateResponseEnvelope]
type bucketCORSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketCORSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketCORSUpdateResponseEnvelopeSuccess bool

const (
	BucketCORSUpdateResponseEnvelopeSuccessTrue BucketCORSUpdateResponseEnvelopeSuccess = true
)

func (r BucketCORSUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketCORSUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketCORSDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketCORSDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketCORSDeleteParamsCfR2Jurisdiction string

const (
	BucketCORSDeleteParamsCfR2JurisdictionDefault BucketCORSDeleteParamsCfR2Jurisdiction = "default"
	BucketCORSDeleteParamsCfR2JurisdictionEu      BucketCORSDeleteParamsCfR2Jurisdiction = "eu"
	BucketCORSDeleteParamsCfR2JurisdictionFedramp BucketCORSDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketCORSDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketCORSDeleteParamsCfR2JurisdictionDefault, BucketCORSDeleteParamsCfR2JurisdictionEu, BucketCORSDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketCORSDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []string                 `json:"messages,required"`
	Result   BucketCORSDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketCORSDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketCORSDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketCORSDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketCORSDeleteResponseEnvelope]
type bucketCORSDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketCORSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketCORSDeleteResponseEnvelopeSuccess bool

const (
	BucketCORSDeleteResponseEnvelopeSuccessTrue BucketCORSDeleteResponseEnvelopeSuccess = true
)

func (r BucketCORSDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketCORSDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketCORSGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketCORSGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketCORSGetParamsCfR2Jurisdiction string

const (
	BucketCORSGetParamsCfR2JurisdictionDefault BucketCORSGetParamsCfR2Jurisdiction = "default"
	BucketCORSGetParamsCfR2JurisdictionEu      BucketCORSGetParamsCfR2Jurisdiction = "eu"
	BucketCORSGetParamsCfR2JurisdictionFedramp BucketCORSGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketCORSGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketCORSGetParamsCfR2JurisdictionDefault, BucketCORSGetParamsCfR2JurisdictionEu, BucketCORSGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketCORSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   BucketCORSGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketCORSGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketCORSGetResponseEnvelopeJSON    `json:"-"`
}

// bucketCORSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketCORSGetResponseEnvelope]
type bucketCORSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketCORSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketCORSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketCORSGetResponseEnvelopeSuccess bool

const (
	BucketCORSGetResponseEnvelopeSuccessTrue BucketCORSGetResponseEnvelopeSuccess = true
)

func (r BucketCORSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketCORSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
