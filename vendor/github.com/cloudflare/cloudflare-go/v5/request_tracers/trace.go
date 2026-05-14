// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers

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

// TraceService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTraceService] method instead.
type TraceService struct {
	Options []option.RequestOption
}

// NewTraceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTraceService(opts ...option.RequestOption) (r *TraceService) {
	r = &TraceService{}
	r.Options = opts
	return
}

// Request Trace
func (r *TraceService) New(ctx context.Context, params TraceNewParams, opts ...option.RequestOption) (res *TraceNewResponse, err error) {
	var env TraceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/request-tracer/trace", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Trace []TraceItem

// List of steps acting on request/response
type TraceItem struct {
	// If step type is rule, then action performed by this rule
	Action string `json:"action"`
	// If step type is rule, then action parameters of this rule as JSON
	ActionParameters interface{} `json:"action_parameters"`
	// If step type is rule or ruleset, the description of this entity
	Description string `json:"description"`
	// If step type is rule, then expression used to match for this rule
	Expression string `json:"expression"`
	// If step type is ruleset, then kind of this ruleset
	Kind string `json:"kind"`
	// Whether tracing step affected tracing request/response
	Matched bool `json:"matched"`
	// If step type is ruleset, then name of this ruleset
	Name string `json:"name"`
	// Tracing step identifying name
	StepName string `json:"step_name"`
	Trace    Trace  `json:"trace"`
	// Tracing step type
	Type string        `json:"type"`
	JSON traceItemJSON `json:"-"`
}

// traceItemJSON contains the JSON metadata for the struct [TraceItem]
type traceItemJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Kind             apijson.Field
	Matched          apijson.Field
	Name             apijson.Field
	StepName         apijson.Field
	Trace            apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TraceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceItemJSON) RawJSON() string {
	return r.raw
}

// Trace result with an origin status code
type TraceNewResponse struct {
	// HTTP Status code of zone response
	StatusCode int64                `json:"status_code"`
	Trace      Trace                `json:"trace"`
	JSON       traceNewResponseJSON `json:"-"`
}

// traceNewResponseJSON contains the JSON metadata for the struct
// [TraceNewResponse]
type traceNewResponseJSON struct {
	StatusCode  apijson.Field
	Trace       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseJSON) RawJSON() string {
	return r.raw
}

type TraceNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// HTTP Method of tracing request
	Method param.Field[string] `json:"method,required"`
	// URL to which perform tracing request
	URL  param.Field[string]             `json:"url,required"`
	Body param.Field[TraceNewParamsBody] `json:"body"`
	// Additional request parameters
	Context param.Field[TraceNewParamsContext] `json:"context"`
	// Cookies added to tracing request
	Cookies param.Field[map[string]string] `json:"cookies"`
	// Headers added to tracing request
	Headers param.Field[map[string]string] `json:"headers"`
	// HTTP Protocol of tracing request
	Protocol param.Field[string] `json:"protocol"`
	// Skip sending the request to the Origin server after all rules evaluation
	SkipResponse param.Field[bool] `json:"skip_response"`
}

func (r TraceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TraceNewParamsBody struct {
	// Base64 encoded request body
	Base64 param.Field[string] `json:"base64"`
	// Arbitrary json as request body
	Json param.Field[interface{}] `json:"json"`
	// Request body as plain text
	PlainText param.Field[string] `json:"plain_text"`
}

func (r TraceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional request parameters
type TraceNewParamsContext struct {
	// Bot score used for evaluating tracing request processing
	BotScore param.Field[int64] `json:"bot_score"`
	// Geodata for tracing request
	Geoloc param.Field[TraceNewParamsContextGeoloc] `json:"geoloc"`
	// Whether to skip any challenges for tracing request (e.g.: captcha)
	SkipChallenge param.Field[bool] `json:"skip_challenge"`
	// Threat score used for evaluating tracing request processing
	ThreatScore param.Field[int64] `json:"threat_score"`
}

func (r TraceNewParamsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geodata for tracing request
type TraceNewParamsContextGeoloc struct {
	City                param.Field[string]  `json:"city"`
	Continent           param.Field[string]  `json:"continent"`
	IsEuCountry         param.Field[bool]    `json:"is_eu_country"`
	ISOCode             param.Field[string]  `json:"iso_code"`
	Latitude            param.Field[float64] `json:"latitude"`
	Longitude           param.Field[float64] `json:"longitude"`
	PostalCode          param.Field[string]  `json:"postal_code"`
	RegionCode          param.Field[string]  `json:"region_code"`
	Subdivision2ISOCode param.Field[string]  `json:"subdivision_2_iso_code"`
	Timezone            param.Field[string]  `json:"timezone"`
}

func (r TraceNewParamsContextGeoloc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TraceNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TraceNewResponseEnvelopeSuccess `json:"success,required"`
	// Trace result with an origin status code
	Result TraceNewResponse             `json:"result"`
	JSON   traceNewResponseEnvelopeJSON `json:"-"`
}

// traceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TraceNewResponseEnvelope]
type traceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TraceNewResponseEnvelopeSuccess bool

const (
	TraceNewResponseEnvelopeSuccessTrue TraceNewResponseEnvelopeSuccess = true
)

func (r TraceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TraceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
