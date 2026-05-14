// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// AccessLogAccessRequestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogAccessRequestService] method instead.
type AccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewAccessLogAccessRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogAccessRequestService(opts ...option.RequestOption) (r *AccessLogAccessRequestService) {
	r = &AccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccessLogAccessRequestService) List(ctx context.Context, params AccessLogAccessRequestListParams, opts ...option.RequestOption) (res *[]AccessRequest, err error) {
	var env AccessLogAccessRequestListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessLogAccessRequestListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The chronological sorting order for the logs.
	Direction param.Field[AccessLogAccessRequestListParamsDirection] `query:"direction"`
	// The maximum number of log entries to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The earliest event timestamp to query.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The latest event timestamp to query.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccessLogAccessRequestListParams]'s query parameters as
// `url.Values`.
func (r AccessLogAccessRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The chronological sorting order for the logs.
type AccessLogAccessRequestListParamsDirection string

const (
	AccessLogAccessRequestListParamsDirectionDesc AccessLogAccessRequestListParamsDirection = "desc"
	AccessLogAccessRequestListParamsDirectionAsc  AccessLogAccessRequestListParamsDirection = "asc"
)

func (r AccessLogAccessRequestListParamsDirection) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsDirectionDesc, AccessLogAccessRequestListParamsDirectionAsc:
		return true
	}
	return false
}

type AccessLogAccessRequestListResponseEnvelope struct {
	Errors   []AccessLogAccessRequestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessLogAccessRequestListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessLogAccessRequestListResponseEnvelopeSuccess `json:"success,required"`
	Result  []AccessRequest                                   `json:"result"`
	JSON    accessLogAccessRequestListResponseEnvelopeJSON    `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessLogAccessRequestListResponseEnvelope]
type accessLogAccessRequestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessLogAccessRequestListResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessLogAccessRequestListResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessLogAccessRequestListResponseEnvelopeErrors]
type accessLogAccessRequestListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessLogAccessRequestListResponseEnvelopeErrorsSource]
type accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessLogAccessRequestListResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessLogAccessRequestListResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessLogAccessRequestListResponseEnvelopeMessages]
type accessLogAccessRequestListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessLogAccessRequestListResponseEnvelopeMessagesSource]
type accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessLogAccessRequestListResponseEnvelopeSuccess bool

const (
	AccessLogAccessRequestListResponseEnvelopeSuccessTrue AccessLogAccessRequestListResponseEnvelopeSuccess = true
)

func (r AccessLogAccessRequestListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
