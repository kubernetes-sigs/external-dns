// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/intel"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// InsightSeverityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightSeverityService] method instead.
type InsightSeverityService struct {
	Options []option.RequestOption
}

// NewInsightSeverityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightSeverityService(opts ...option.RequestOption) (r *InsightSeverityService) {
	r = &InsightSeverityService{}
	r.Options = opts
	return
}

// Get Security Center Insight Counts by Severity
func (r *InsightSeverityService) Get(ctx context.Context, params InsightSeverityGetParams, opts ...option.RequestOption) (res *[]InsightSeverityGetResponse, err error) {
	var env InsightSeverityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/severity", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InsightSeverityGetResponse struct {
	Count int64                          `json:"count"`
	Value string                         `json:"value"`
	JSON  insightSeverityGetResponseJSON `json:"-"`
}

// insightSeverityGetResponseJSON contains the JSON metadata for the struct
// [InsightSeverityGetResponse]
type insightSeverityGetResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightSeverityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseJSON) RawJSON() string {
	return r.raw
}

type InsightSeverityGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]                     `path:"zone_id"`
	Dismissed     param.Field[bool]                       `query:"dismissed"`
	IssueClass    param.Field[[]string]                   `query:"issue_class"`
	IssueClassNeq param.Field[[]string]                   `query:"issue_class~neq"`
	IssueType     param.Field[[]intel.IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]intel.IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]                   `query:"product"`
	ProductNeq    param.Field[[]string]                   `query:"product~neq"`
	Severity      param.Field[[]intel.SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]intel.SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]                   `query:"subject"`
	SubjectNeq    param.Field[[]string]                   `query:"subject~neq"`
}

// URLQuery serializes [InsightSeverityGetParams]'s query parameters as
// `url.Values`.
func (r InsightSeverityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InsightSeverityGetResponseEnvelope struct {
	Errors   []InsightSeverityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InsightSeverityGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success InsightSeverityGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []InsightSeverityGetResponse              `json:"result"`
	JSON    insightSeverityGetResponseEnvelopeJSON    `json:"-"`
}

// insightSeverityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InsightSeverityGetResponseEnvelope]
type insightSeverityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightSeverityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InsightSeverityGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           InsightSeverityGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             insightSeverityGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// insightSeverityGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InsightSeverityGetResponseEnvelopeErrors]
type insightSeverityGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightSeverityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InsightSeverityGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    insightSeverityGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// insightSeverityGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InsightSeverityGetResponseEnvelopeErrorsSource]
type insightSeverityGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightSeverityGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InsightSeverityGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           InsightSeverityGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             insightSeverityGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// insightSeverityGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InsightSeverityGetResponseEnvelopeMessages]
type insightSeverityGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightSeverityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InsightSeverityGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    insightSeverityGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// insightSeverityGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InsightSeverityGetResponseEnvelopeMessagesSource]
type insightSeverityGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightSeverityGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightSeverityGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InsightSeverityGetResponseEnvelopeSuccess bool

const (
	InsightSeverityGetResponseEnvelopeSuccessTrue InsightSeverityGetResponseEnvelopeSuccess = true
)

func (r InsightSeverityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InsightSeverityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
