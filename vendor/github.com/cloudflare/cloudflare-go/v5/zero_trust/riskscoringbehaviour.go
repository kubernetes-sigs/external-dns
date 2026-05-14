// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// RiskScoringBehaviourService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringBehaviourService] method instead.
type RiskScoringBehaviourService struct {
	Options []option.RequestOption
}

// NewRiskScoringBehaviourService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRiskScoringBehaviourService(opts ...option.RequestOption) (r *RiskScoringBehaviourService) {
	r = &RiskScoringBehaviourService{}
	r.Options = opts
	return
}

// Update configuration for risk behaviors
func (r *RiskScoringBehaviourService) Update(ctx context.Context, params RiskScoringBehaviourUpdateParams, opts ...option.RequestOption) (res *RiskScoringBehaviourUpdateResponse, err error) {
	var env RiskScoringBehaviourUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get all behaviors and associated configuration
func (r *RiskScoringBehaviourService) Get(ctx context.Context, query RiskScoringBehaviourGetParams, opts ...option.RequestOption) (res *RiskScoringBehaviourGetResponse, err error) {
	var env RiskScoringBehaviourGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringBehaviourUpdateResponse struct {
	Behaviors map[string]RiskScoringBehaviourUpdateResponseBehavior `json:"behaviors,required"`
	JSON      riskScoringBehaviourUpdateResponseJSON                `json:"-"`
}

// riskScoringBehaviourUpdateResponseJSON contains the JSON metadata for the struct
// [RiskScoringBehaviourUpdateResponse]
type riskScoringBehaviourUpdateResponseJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseBehavior struct {
	Enabled   bool                                                 `json:"enabled,required"`
	RiskLevel RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel `json:"risk_level,required"`
	JSON      riskScoringBehaviourUpdateResponseBehaviorJSON       `json:"-"`
}

// riskScoringBehaviourUpdateResponseBehaviorJSON contains the JSON metadata for
// the struct [RiskScoringBehaviourUpdateResponseBehavior]
type riskScoringBehaviourUpdateResponseBehaviorJSON struct {
	Enabled     apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseBehaviorJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel string

const (
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelLow    RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "low"
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelMedium RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelHigh   RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelLow, RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelMedium, RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourGetResponse struct {
	Behaviors map[string]RiskScoringBehaviourGetResponseBehavior `json:"behaviors,required"`
	JSON      riskScoringBehaviourGetResponseJSON                `json:"-"`
}

// riskScoringBehaviourGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringBehaviourGetResponse]
type riskScoringBehaviourGetResponseJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseBehavior struct {
	Description string                                            `json:"description,required"`
	Enabled     bool                                              `json:"enabled,required"`
	Name        string                                            `json:"name,required"`
	RiskLevel   RiskScoringBehaviourGetResponseBehaviorsRiskLevel `json:"risk_level,required"`
	JSON        riskScoringBehaviourGetResponseBehaviorJSON       `json:"-"`
}

// riskScoringBehaviourGetResponseBehaviorJSON contains the JSON metadata for the
// struct [RiskScoringBehaviourGetResponseBehavior]
type riskScoringBehaviourGetResponseBehaviorJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseBehaviorJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseBehaviorsRiskLevel string

const (
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelLow    RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "low"
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelMedium RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelHigh   RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourGetResponseBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourGetResponseBehaviorsRiskLevelLow, RiskScoringBehaviourGetResponseBehaviorsRiskLevelMedium, RiskScoringBehaviourGetResponseBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourUpdateParams struct {
	AccountID param.Field[string]                                               `path:"account_id,required"`
	Behaviors param.Field[map[string]RiskScoringBehaviourUpdateParamsBehaviors] `json:"behaviors,required"`
}

func (r RiskScoringBehaviourUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringBehaviourUpdateParamsBehaviors struct {
	Enabled   param.Field[bool]                                               `json:"enabled,required"`
	RiskLevel param.Field[RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel] `json:"risk_level,required"`
}

func (r RiskScoringBehaviourUpdateParamsBehaviors) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel string

const (
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelLow    RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "low"
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelMedium RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelHigh   RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelLow, RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelMedium, RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourUpdateResponseEnvelope struct {
	Errors   []RiskScoringBehaviourUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RiskScoringBehaviourUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RiskScoringBehaviourUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringBehaviourUpdateResponse                `json:"result"`
	JSON    riskScoringBehaviourUpdateResponseEnvelopeJSON    `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [RiskScoringBehaviourUpdateResponseEnvelope]
type riskScoringBehaviourUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           RiskScoringBehaviourUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringBehaviourUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [RiskScoringBehaviourUpdateResponseEnvelopeErrors]
type riskScoringBehaviourUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    riskScoringBehaviourUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RiskScoringBehaviourUpdateResponseEnvelopeErrorsSource]
type riskScoringBehaviourUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           RiskScoringBehaviourUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringBehaviourUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [RiskScoringBehaviourUpdateResponseEnvelopeMessages]
type riskScoringBehaviourUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    riskScoringBehaviourUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [RiskScoringBehaviourUpdateResponseEnvelopeMessagesSource]
type riskScoringBehaviourUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RiskScoringBehaviourUpdateResponseEnvelopeSuccess bool

const (
	RiskScoringBehaviourUpdateResponseEnvelopeSuccessTrue RiskScoringBehaviourUpdateResponseEnvelopeSuccess = true
)

func (r RiskScoringBehaviourUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringBehaviourGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringBehaviourGetResponseEnvelope struct {
	Errors   []RiskScoringBehaviourGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RiskScoringBehaviourGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RiskScoringBehaviourGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringBehaviourGetResponse                `json:"result"`
	JSON    riskScoringBehaviourGetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringBehaviourGetResponseEnvelope]
type riskScoringBehaviourGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           RiskScoringBehaviourGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringBehaviourGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RiskScoringBehaviourGetResponseEnvelopeErrors]
type riskScoringBehaviourGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    riskScoringBehaviourGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RiskScoringBehaviourGetResponseEnvelopeErrorsSource]
type riskScoringBehaviourGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           RiskScoringBehaviourGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringBehaviourGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RiskScoringBehaviourGetResponseEnvelopeMessages]
type riskScoringBehaviourGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    riskScoringBehaviourGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RiskScoringBehaviourGetResponseEnvelopeMessagesSource]
type riskScoringBehaviourGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RiskScoringBehaviourGetResponseEnvelopeSuccess bool

const (
	RiskScoringBehaviourGetResponseEnvelopeSuccessTrue RiskScoringBehaviourGetResponseEnvelopeSuccess = true
)

func (r RiskScoringBehaviourGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
