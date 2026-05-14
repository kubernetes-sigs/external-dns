// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// DEXCommandQuotaService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandQuotaService] method instead.
type DEXCommandQuotaService struct {
	Options []option.RequestOption
}

// NewDEXCommandQuotaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXCommandQuotaService(opts ...option.RequestOption) (r *DEXCommandQuotaService) {
	r = &DEXCommandQuotaService{}
	r.Options = opts
	return
}

// Retrieves the current quota usage and limits for device commands within a
// specific account, including the time when the quota will reset
func (r *DEXCommandQuotaService) Get(ctx context.Context, query DEXCommandQuotaGetParams, opts ...option.RequestOption) (res *DEXCommandQuotaGetResponse, err error) {
	var env DEXCommandQuotaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/quota", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXCommandQuotaGetResponse struct {
	// The remaining number of commands that can be initiated for an account
	Quota float64 `json:"quota,required"`
	// The number of commands that have been initiated for an account
	QuotaUsage float64 `json:"quota_usage,required"`
	// The time when the quota resets
	ResetTime time.Time                      `json:"reset_time,required" format:"date-time"`
	JSON      dexCommandQuotaGetResponseJSON `json:"-"`
}

// dexCommandQuotaGetResponseJSON contains the JSON metadata for the struct
// [DEXCommandQuotaGetResponse]
type dexCommandQuotaGetResponseJSON struct {
	Quota       apijson.Field
	QuotaUsage  apijson.Field
	ResetTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandQuotaGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXCommandQuotaGetResponseEnvelope struct {
	Errors   []DEXCommandQuotaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXCommandQuotaGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    DEXCommandQuotaGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     DEXCommandQuotaGetResponse                   `json:"result"`
	ResultInfo DEXCommandQuotaGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexCommandQuotaGetResponseEnvelopeJSON       `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXCommandQuotaGetResponseEnvelope]
type dexCommandQuotaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXCommandQuotaGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DEXCommandQuotaGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexCommandQuotaGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXCommandQuotaGetResponseEnvelopeErrors]
type dexCommandQuotaGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXCommandQuotaGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dexCommandQuotaGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DEXCommandQuotaGetResponseEnvelopeErrorsSource]
type dexCommandQuotaGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXCommandQuotaGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DEXCommandQuotaGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexCommandQuotaGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DEXCommandQuotaGetResponseEnvelopeMessages]
type dexCommandQuotaGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXCommandQuotaGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dexCommandQuotaGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DEXCommandQuotaGetResponseEnvelopeMessagesSource]
type dexCommandQuotaGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXCommandQuotaGetResponseEnvelopeSuccess bool

const (
	DEXCommandQuotaGetResponseEnvelopeSuccessTrue DEXCommandQuotaGetResponseEnvelopeSuccess = true
)

func (r DEXCommandQuotaGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXCommandQuotaGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXCommandQuotaGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                          `json:"total_count"`
	JSON       dexCommandQuotaGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexCommandQuotaGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DEXCommandQuotaGetResponseEnvelopeResultInfo]
type dexCommandQuotaGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandQuotaGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandQuotaGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
