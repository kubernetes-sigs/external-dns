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

// RiskScoringSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringSummaryService] method instead.
type RiskScoringSummaryService struct {
	Options []option.RequestOption
}

// NewRiskScoringSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRiskScoringSummaryService(opts ...option.RequestOption) (r *RiskScoringSummaryService) {
	r = &RiskScoringSummaryService{}
	r.Options = opts
	return
}

// Get risk score info for all users in the account
func (r *RiskScoringSummaryService) Get(ctx context.Context, query RiskScoringSummaryGetParams, opts ...option.RequestOption) (res *RiskScoringSummaryGetResponse, err error) {
	var env RiskScoringSummaryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/summary", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringSummaryGetResponse struct {
	Users []RiskScoringSummaryGetResponseUser `json:"users,required"`
	JSON  riskScoringSummaryGetResponseJSON   `json:"-"`
}

// riskScoringSummaryGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringSummaryGetResponse]
type riskScoringSummaryGetResponseJSON struct {
	Users       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseUser struct {
	Email        string                                         `json:"email,required"`
	EventCount   int64                                          `json:"event_count,required"`
	LastEvent    time.Time                                      `json:"last_event,required" format:"date-time"`
	MaxRiskLevel RiskScoringSummaryGetResponseUsersMaxRiskLevel `json:"max_risk_level,required"`
	Name         string                                         `json:"name,required"`
	UserID       string                                         `json:"user_id,required" format:"uuid"`
	JSON         riskScoringSummaryGetResponseUserJSON          `json:"-"`
}

// riskScoringSummaryGetResponseUserJSON contains the JSON metadata for the struct
// [RiskScoringSummaryGetResponseUser]
type riskScoringSummaryGetResponseUserJSON struct {
	Email        apijson.Field
	EventCount   apijson.Field
	LastEvent    apijson.Field
	MaxRiskLevel apijson.Field
	Name         apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseUsersMaxRiskLevel string

const (
	RiskScoringSummaryGetResponseUsersMaxRiskLevelLow    RiskScoringSummaryGetResponseUsersMaxRiskLevel = "low"
	RiskScoringSummaryGetResponseUsersMaxRiskLevelMedium RiskScoringSummaryGetResponseUsersMaxRiskLevel = "medium"
	RiskScoringSummaryGetResponseUsersMaxRiskLevelHigh   RiskScoringSummaryGetResponseUsersMaxRiskLevel = "high"
)

func (r RiskScoringSummaryGetResponseUsersMaxRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetResponseUsersMaxRiskLevelLow, RiskScoringSummaryGetResponseUsersMaxRiskLevelMedium, RiskScoringSummaryGetResponseUsersMaxRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringSummaryGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringSummaryGetResponseEnvelope struct {
	Errors   []RiskScoringSummaryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RiskScoringSummaryGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    RiskScoringSummaryGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     RiskScoringSummaryGetResponse                   `json:"result"`
	ResultInfo RiskScoringSummaryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringSummaryGetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringSummaryGetResponseEnvelope]
type riskScoringSummaryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           RiskScoringSummaryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringSummaryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RiskScoringSummaryGetResponseEnvelopeErrors]
type riskScoringSummaryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    riskScoringSummaryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RiskScoringSummaryGetResponseEnvelopeErrorsSource]
type riskScoringSummaryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           RiskScoringSummaryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringSummaryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RiskScoringSummaryGetResponseEnvelopeMessages]
type riskScoringSummaryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringSummaryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    riskScoringSummaryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RiskScoringSummaryGetResponseEnvelopeMessagesSource]
type riskScoringSummaryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RiskScoringSummaryGetResponseEnvelopeSuccess bool

const (
	RiskScoringSummaryGetResponseEnvelopeSuccessTrue RiskScoringSummaryGetResponseEnvelopeSuccess = true
)

func (r RiskScoringSummaryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringSummaryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringSummaryGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                             `json:"total_count"`
	JSON       riskScoringSummaryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringSummaryGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [RiskScoringSummaryGetResponseEnvelopeResultInfo]
type riskScoringSummaryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringSummaryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringSummaryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
