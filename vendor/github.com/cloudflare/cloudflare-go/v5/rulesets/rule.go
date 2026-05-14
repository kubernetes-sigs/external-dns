// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RuleNewResponse, err error) {
	var env RuleNewResponseEnvelope
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account or zone ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *RuleEditResponse, err error) {
	var env RuleEditResponseEnvelope
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action BlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters BlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck BlockRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit BlockRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string        `json:"ref"`
	JSON blockRuleJSON `json:"-"`
}

// blockRuleJSON contains the JSON metadata for the struct [BlockRule]
type blockRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleJSON) RawJSON() string {
	return r.raw
}

func (r BlockRule) implementsRulesetNewResponseRule() {}

func (r BlockRule) implementsRulesetUpdateResponseRule() {}

func (r BlockRule) implementsRulesetGetResponseRule() {}

func (r BlockRule) implementsPhaseUpdateResponseRule() {}

func (r BlockRule) implementsPhaseGetResponseRule() {}

func (r BlockRule) implementsPhaseVersionGetResponseRule() {}

func (r BlockRule) implementsRuleNewResponseRule() {}

func (r BlockRule) implementsRuleDeleteResponseRule() {}

func (r BlockRule) implementsRuleEditResponseRule() {}

func (r BlockRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type BlockRuleAction string

const (
	BlockRuleActionBlock BlockRuleAction = "block"
)

func (r BlockRuleAction) IsKnown() bool {
	switch r {
	case BlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type BlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response BlockRuleActionParametersResponse `json:"response"`
	JSON     blockRuleActionParametersJSON     `json:"-"`
}

// blockRuleActionParametersJSON contains the JSON metadata for the struct
// [BlockRuleActionParameters]
type blockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type BlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                 `json:"status_code,required"`
	JSON       blockRuleActionParametersResponseJSON `json:"-"`
}

// blockRuleActionParametersResponseJSON contains the JSON metadata for the struct
// [BlockRuleActionParametersResponse]
type blockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type BlockRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                              `json:"username_expression,required"`
	JSON               blockRuleExposedCredentialCheckJSON `json:"-"`
}

// blockRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [BlockRuleExposedCredentialCheck]
type blockRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BlockRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type BlockRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                 `json:"score_response_header_name"`
	JSON                    blockRuleRatelimitJSON `json:"-"`
}

// blockRuleRatelimitJSON contains the JSON metadata for the struct
// [BlockRuleRatelimit]
type blockRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BlockRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type BlockRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[BlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[BlockRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[BlockRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[BlockRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r BlockRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BlockRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r BlockRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r BlockRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type BlockRuleActionParametersParam struct {
	// The response to show when the block is applied.
	Response param.Field[BlockRuleActionParametersResponseParam] `json:"response"`
}

func (r BlockRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type BlockRuleActionParametersResponseParam struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r BlockRuleActionParametersResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type BlockRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r BlockRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type BlockRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r BlockRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action CompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters CompressResponseRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck CompressResponseRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit CompressResponseRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                   `json:"ref"`
	JSON compressResponseRuleJSON `json:"-"`
}

// compressResponseRuleJSON contains the JSON metadata for the struct
// [CompressResponseRule]
type compressResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r CompressResponseRule) implementsRulesetNewResponseRule() {}

func (r CompressResponseRule) implementsRulesetUpdateResponseRule() {}

func (r CompressResponseRule) implementsRulesetGetResponseRule() {}

func (r CompressResponseRule) implementsPhaseUpdateResponseRule() {}

func (r CompressResponseRule) implementsPhaseGetResponseRule() {}

func (r CompressResponseRule) implementsPhaseVersionGetResponseRule() {}

func (r CompressResponseRule) implementsRuleNewResponseRule() {}

func (r CompressResponseRule) implementsRuleDeleteResponseRule() {}

func (r CompressResponseRule) implementsRuleEditResponseRule() {}

func (r CompressResponseRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type CompressResponseRuleAction string

const (
	CompressResponseRuleActionCompressResponse CompressResponseRuleAction = "compress_response"
)

func (r CompressResponseRuleAction) IsKnown() bool {
	switch r {
	case CompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type CompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []CompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       compressResponseRuleActionParametersJSON        `json:"-"`
}

// compressResponseRuleActionParametersJSON contains the JSON metadata for the
// struct [CompressResponseRuleActionParameters]
type compressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name CompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON compressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// compressResponseRuleActionParametersAlgorithmJSON contains the JSON metadata for
// the struct [CompressResponseRuleActionParametersAlgorithm]
type compressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithmsName string

const (
	CompressResponseRuleActionParametersAlgorithmsNameNone    CompressResponseRuleActionParametersAlgorithmsName = "none"
	CompressResponseRuleActionParametersAlgorithmsNameAuto    CompressResponseRuleActionParametersAlgorithmsName = "auto"
	CompressResponseRuleActionParametersAlgorithmsNameDefault CompressResponseRuleActionParametersAlgorithmsName = "default"
	CompressResponseRuleActionParametersAlgorithmsNameGzip    CompressResponseRuleActionParametersAlgorithmsName = "gzip"
	CompressResponseRuleActionParametersAlgorithmsNameBrotli  CompressResponseRuleActionParametersAlgorithmsName = "brotli"
	CompressResponseRuleActionParametersAlgorithmsNameZstd    CompressResponseRuleActionParametersAlgorithmsName = "zstd"
)

func (r CompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case CompressResponseRuleActionParametersAlgorithmsNameNone, CompressResponseRuleActionParametersAlgorithmsNameAuto, CompressResponseRuleActionParametersAlgorithmsNameDefault, CompressResponseRuleActionParametersAlgorithmsNameGzip, CompressResponseRuleActionParametersAlgorithmsNameBrotli, CompressResponseRuleActionParametersAlgorithmsNameZstd:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type CompressResponseRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                         `json:"username_expression,required"`
	JSON               compressResponseRuleExposedCredentialCheckJSON `json:"-"`
}

// compressResponseRuleExposedCredentialCheckJSON contains the JSON metadata for
// the struct [CompressResponseRuleExposedCredentialCheck]
type compressResponseRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CompressResponseRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type CompressResponseRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                            `json:"score_response_header_name"`
	JSON                    compressResponseRuleRatelimitJSON `json:"-"`
}

// compressResponseRuleRatelimitJSON contains the JSON metadata for the struct
// [CompressResponseRuleRatelimit]
type compressResponseRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CompressResponseRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type CompressResponseRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[CompressResponseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[CompressResponseRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[CompressResponseRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[CompressResponseRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r CompressResponseRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CompressResponseRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r CompressResponseRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r CompressResponseRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type CompressResponseRuleActionParametersParam struct {
	// Custom order for compression algorithms.
	Algorithms param.Field[[]CompressResponseRuleActionParametersAlgorithmParam] `json:"algorithms"`
}

func (r CompressResponseRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithmParam struct {
	// Name of compression algorithm to enable.
	Name param.Field[CompressResponseRuleActionParametersAlgorithmsName] `json:"name"`
}

func (r CompressResponseRuleActionParametersAlgorithmParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type CompressResponseRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r CompressResponseRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type CompressResponseRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r CompressResponseRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DDoSDynamicRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action DDoSDynamicRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck DDoSDynamicRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit DDoSDynamicRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string              `json:"ref"`
	JSON DDoSDynamicRuleJSON `json:"-"`
}

// DDoSDynamicRuleJSON contains the JSON metadata for the struct [DDoSDynamicRule]
type DDoSDynamicRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r DDoSDynamicRule) implementsRulesetNewResponseRule() {}

func (r DDoSDynamicRule) implementsRulesetUpdateResponseRule() {}

func (r DDoSDynamicRule) implementsRulesetGetResponseRule() {}

func (r DDoSDynamicRule) implementsPhaseUpdateResponseRule() {}

func (r DDoSDynamicRule) implementsPhaseGetResponseRule() {}

func (r DDoSDynamicRule) implementsPhaseVersionGetResponseRule() {}

func (r DDoSDynamicRule) implementsRuleNewResponseRule() {}

func (r DDoSDynamicRule) implementsRuleDeleteResponseRule() {}

func (r DDoSDynamicRule) implementsRuleEditResponseRule() {}

func (r DDoSDynamicRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type DDoSDynamicRuleAction string

const (
	DDoSDynamicRuleActionDDoSDynamic DDoSDynamicRuleAction = "ddos_dynamic"
)

func (r DDoSDynamicRuleAction) IsKnown() bool {
	switch r {
	case DDoSDynamicRuleActionDDoSDynamic:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type DDoSDynamicRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                    `json:"username_expression,required"`
	JSON               DDoSDynamicRuleExposedCredentialCheckJSON `json:"-"`
}

// DDoSDynamicRuleExposedCredentialCheckJSON contains the JSON metadata for the
// struct [DDoSDynamicRuleExposedCredentialCheck]
type DDoSDynamicRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DDoSDynamicRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DDoSDynamicRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type DDoSDynamicRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                       `json:"score_response_header_name"`
	JSON                    DDoSDynamicRuleRatelimitJSON `json:"-"`
}

// DDoSDynamicRuleRatelimitJSON contains the JSON metadata for the struct
// [DDoSDynamicRuleRatelimit]
type DDoSDynamicRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DDoSDynamicRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DDoSDynamicRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type DDoSDynamicRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[DDoSDynamicRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[DDoSDynamicRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[DDoSDynamicRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r DDoSDynamicRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DDoSDynamicRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r DDoSDynamicRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r DDoSDynamicRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// Configure checks for exposed credentials.
type DDoSDynamicRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r DDoSDynamicRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type DDoSDynamicRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r DDoSDynamicRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck ExecuteRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit ExecuteRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string          `json:"ref"`
	JSON executeRuleJSON `json:"-"`
}

// executeRuleJSON contains the JSON metadata for the struct [ExecuteRule]
type executeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteRule) implementsRulesetNewResponseRule() {}

func (r ExecuteRule) implementsRulesetUpdateResponseRule() {}

func (r ExecuteRule) implementsRulesetGetResponseRule() {}

func (r ExecuteRule) implementsPhaseUpdateResponseRule() {}

func (r ExecuteRule) implementsPhaseGetResponseRule() {}

func (r ExecuteRule) implementsPhaseVersionGetResponseRule() {}

func (r ExecuteRule) implementsRuleNewResponseRule() {}

func (r ExecuteRule) implementsRuleDeleteResponseRule() {}

func (r ExecuteRule) implementsRuleEditResponseRule() {}

func (r ExecuteRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type ExecuteRuleAction string

const (
	ExecuteRuleActionExecute ExecuteRuleAction = "execute"
)

func (r ExecuteRuleAction) IsKnown() bool {
	switch r {
	case ExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData ExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides ExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      executeRuleActionParametersJSON      `json:"-"`
}

// executeRuleActionParametersJSON contains the JSON metadata for the struct
// [ExecuteRuleActionParameters]
type executeRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type ExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                     `json:"public_key,required"`
	JSON      executeRuleActionParametersMatchedDataJSON `json:"-"`
}

// executeRuleActionParametersMatchedDataJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersMatchedData]
type executeRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type ExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []ExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []ExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel ExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesJSON             `json:"-"`
}

// executeRuleActionParametersOverridesJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersOverrides]
type executeRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override.
type ExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// executeRuleActionParametersOverridesCategoryJSON contains the JSON metadata for
// the struct [ExecuteRuleActionParametersOverridesCategory]
type executeRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override.
type ExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel ExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// executeRuleActionParametersOverridesRuleJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersOverridesRule]
type executeRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type ExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, ExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type ExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesSensitivityLevelDefault ExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesSensitivityLevelLow     ExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesSensitivityLevelLow, ExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type ExecuteRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                `json:"username_expression,required"`
	JSON               executeRuleExposedCredentialCheckJSON `json:"-"`
}

// executeRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [ExecuteRuleExposedCredentialCheck]
type executeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExecuteRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type ExecuteRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                   `json:"score_response_header_name"`
	JSON                    executeRuleRatelimitJSON `json:"-"`
}

// executeRuleRatelimitJSON contains the JSON metadata for the struct
// [ExecuteRuleRatelimit]
type executeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ExecuteRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type ExecuteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ExecuteRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ExecuteRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ExecuteRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ExecuteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExecuteRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r ExecuteRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r ExecuteRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type ExecuteRuleActionParametersParam struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[ExecuteRuleActionParametersMatchedDataParam] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[ExecuteRuleActionParametersOverridesParam] `json:"overrides"`
}

func (r ExecuteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type ExecuteRuleActionParametersMatchedDataParam struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ExecuteRuleActionParametersMatchedDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type ExecuteRuleActionParametersOverridesParam struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]ExecuteRuleActionParametersOverridesCategoryParam] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]ExecuteRuleActionParametersOverridesRuleParam] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override.
type ExecuteRuleActionParametersOverridesCategoryParam struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override.
type ExecuteRuleActionParametersOverridesRuleParam struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ExecuteRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ExecuteRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ExecuteRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r ExecuteRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ForceConnectionCloseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ForceConnectionCloseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck ForceConnectionCloseRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit ForceConnectionCloseRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                       `json:"ref"`
	JSON forceConnectionCloseRuleJSON `json:"-"`
}

// forceConnectionCloseRuleJSON contains the JSON metadata for the struct
// [ForceConnectionCloseRule]
type forceConnectionCloseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r ForceConnectionCloseRule) implementsRulesetNewResponseRule() {}

func (r ForceConnectionCloseRule) implementsRulesetUpdateResponseRule() {}

func (r ForceConnectionCloseRule) implementsRulesetGetResponseRule() {}

func (r ForceConnectionCloseRule) implementsPhaseUpdateResponseRule() {}

func (r ForceConnectionCloseRule) implementsPhaseGetResponseRule() {}

func (r ForceConnectionCloseRule) implementsPhaseVersionGetResponseRule() {}

func (r ForceConnectionCloseRule) implementsRuleNewResponseRule() {}

func (r ForceConnectionCloseRule) implementsRuleDeleteResponseRule() {}

func (r ForceConnectionCloseRule) implementsRuleEditResponseRule() {}

func (r ForceConnectionCloseRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type ForceConnectionCloseRuleAction string

const (
	ForceConnectionCloseRuleActionForceConnectionClose ForceConnectionCloseRuleAction = "force_connection_close"
)

func (r ForceConnectionCloseRuleAction) IsKnown() bool {
	switch r {
	case ForceConnectionCloseRuleActionForceConnectionClose:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type ForceConnectionCloseRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                             `json:"username_expression,required"`
	JSON               forceConnectionCloseRuleExposedCredentialCheckJSON `json:"-"`
}

// forceConnectionCloseRuleExposedCredentialCheckJSON contains the JSON metadata
// for the struct [ForceConnectionCloseRuleExposedCredentialCheck]
type forceConnectionCloseRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ForceConnectionCloseRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceConnectionCloseRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type ForceConnectionCloseRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                `json:"score_response_header_name"`
	JSON                    forceConnectionCloseRuleRatelimitJSON `json:"-"`
}

// forceConnectionCloseRuleRatelimitJSON contains the JSON metadata for the struct
// [ForceConnectionCloseRuleRatelimit]
type forceConnectionCloseRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ForceConnectionCloseRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forceConnectionCloseRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type ForceConnectionCloseRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ForceConnectionCloseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ForceConnectionCloseRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ForceConnectionCloseRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ForceConnectionCloseRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ForceConnectionCloseRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r ForceConnectionCloseRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r ForceConnectionCloseRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// Configure checks for exposed credentials.
type ForceConnectionCloseRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ForceConnectionCloseRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ForceConnectionCloseRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r ForceConnectionCloseRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action LogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters LogCustomFieldRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck LogCustomFieldRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit LogCustomFieldRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                 `json:"ref"`
	JSON logCustomFieldRuleJSON `json:"-"`
}

// logCustomFieldRuleJSON contains the JSON metadata for the struct
// [LogCustomFieldRule]
type logCustomFieldRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *LogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r LogCustomFieldRule) implementsRulesetNewResponseRule() {}

func (r LogCustomFieldRule) implementsRulesetUpdateResponseRule() {}

func (r LogCustomFieldRule) implementsRulesetGetResponseRule() {}

func (r LogCustomFieldRule) implementsPhaseUpdateResponseRule() {}

func (r LogCustomFieldRule) implementsPhaseGetResponseRule() {}

func (r LogCustomFieldRule) implementsPhaseVersionGetResponseRule() {}

func (r LogCustomFieldRule) implementsRuleNewResponseRule() {}

func (r LogCustomFieldRule) implementsRuleDeleteResponseRule() {}

func (r LogCustomFieldRule) implementsRuleEditResponseRule() {}

func (r LogCustomFieldRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type LogCustomFieldRuleAction string

const (
	LogCustomFieldRuleActionLogCustomField LogCustomFieldRuleAction = "log_custom_field"
)

func (r LogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case LogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type LogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []LogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The raw response fields to log.
	RawResponseFields []LogCustomFieldRuleActionParametersRawResponseField `json:"raw_response_fields"`
	// The raw request fields to log.
	RequestFields []LogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The transformed response fields to log.
	ResponseFields []LogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	// The transformed request fields to log.
	TransformedRequestFields []LogCustomFieldRuleActionParametersTransformedRequestField `json:"transformed_request_fields"`
	JSON                     logCustomFieldRuleActionParametersJSON                      `json:"-"`
}

// logCustomFieldRuleActionParametersJSON contains the JSON metadata for the struct
// [LogCustomFieldRuleActionParameters]
type logCustomFieldRuleActionParametersJSON struct {
	CookieFields             apijson.Field
	RawResponseFields        apijson.Field
	RequestFields            apijson.Field
	ResponseFields           apijson.Field
	TransformedRequestFields apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type LogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                            `json:"name,required"`
	JSON logCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// logCustomFieldRuleActionParametersCookieFieldJSON contains the JSON metadata for
// the struct [LogCustomFieldRuleActionParametersCookieField]
type logCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The raw response field to log.
type LogCustomFieldRuleActionParametersRawResponseField struct {
	// The name of the field.
	Name string `json:"name,required"`
	// Whether to log duplicate values of the same header.
	PreserveDuplicates bool                                                   `json:"preserve_duplicates"`
	JSON               logCustomFieldRuleActionParametersRawResponseFieldJSON `json:"-"`
}

// logCustomFieldRuleActionParametersRawResponseFieldJSON contains the JSON
// metadata for the struct [LogCustomFieldRuleActionParametersRawResponseField]
type logCustomFieldRuleActionParametersRawResponseFieldJSON struct {
	Name               apijson.Field
	PreserveDuplicates apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParametersRawResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersRawResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The raw request field to log.
type LogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                             `json:"name,required"`
	JSON logCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// logCustomFieldRuleActionParametersRequestFieldJSON contains the JSON metadata
// for the struct [LogCustomFieldRuleActionParametersRequestField]
type logCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The transformed response field to log.
type LogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string `json:"name,required"`
	// Whether to log duplicate values of the same header.
	PreserveDuplicates bool                                                `json:"preserve_duplicates"`
	JSON               logCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// logCustomFieldRuleActionParametersResponseFieldJSON contains the JSON metadata
// for the struct [LogCustomFieldRuleActionParametersResponseField]
type logCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name               apijson.Field
	PreserveDuplicates apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The transformed request field to log.
type LogCustomFieldRuleActionParametersTransformedRequestField struct {
	// The name of the field.
	Name string                                                        `json:"name,required"`
	JSON logCustomFieldRuleActionParametersTransformedRequestFieldJSON `json:"-"`
}

// logCustomFieldRuleActionParametersTransformedRequestFieldJSON contains the JSON
// metadata for the struct
// [LogCustomFieldRuleActionParametersTransformedRequestField]
type logCustomFieldRuleActionParametersTransformedRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogCustomFieldRuleActionParametersTransformedRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleActionParametersTransformedRequestFieldJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type LogCustomFieldRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                       `json:"username_expression,required"`
	JSON               logCustomFieldRuleExposedCredentialCheckJSON `json:"-"`
}

// logCustomFieldRuleExposedCredentialCheckJSON contains the JSON metadata for the
// struct [LogCustomFieldRuleExposedCredentialCheck]
type logCustomFieldRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LogCustomFieldRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type LogCustomFieldRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                          `json:"score_response_header_name"`
	JSON                    logCustomFieldRuleRatelimitJSON `json:"-"`
}

// logCustomFieldRuleRatelimitJSON contains the JSON metadata for the struct
// [LogCustomFieldRuleRatelimit]
type logCustomFieldRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LogCustomFieldRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logCustomFieldRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type LogCustomFieldRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[LogCustomFieldRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[LogCustomFieldRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[LogCustomFieldRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[LogCustomFieldRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r LogCustomFieldRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LogCustomFieldRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r LogCustomFieldRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r LogCustomFieldRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type LogCustomFieldRuleActionParametersParam struct {
	// The cookie fields to log.
	CookieFields param.Field[[]LogCustomFieldRuleActionParametersCookieFieldParam] `json:"cookie_fields"`
	// The raw response fields to log.
	RawResponseFields param.Field[[]LogCustomFieldRuleActionParametersRawResponseFieldParam] `json:"raw_response_fields"`
	// The raw request fields to log.
	RequestFields param.Field[[]LogCustomFieldRuleActionParametersRequestFieldParam] `json:"request_fields"`
	// The transformed response fields to log.
	ResponseFields param.Field[[]LogCustomFieldRuleActionParametersResponseFieldParam] `json:"response_fields"`
	// The transformed request fields to log.
	TransformedRequestFields param.Field[[]LogCustomFieldRuleActionParametersTransformedRequestFieldParam] `json:"transformed_request_fields"`
}

func (r LogCustomFieldRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookie field to log.
type LogCustomFieldRuleActionParametersCookieFieldParam struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r LogCustomFieldRuleActionParametersCookieFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The raw response field to log.
type LogCustomFieldRuleActionParametersRawResponseFieldParam struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
	// Whether to log duplicate values of the same header.
	PreserveDuplicates param.Field[bool] `json:"preserve_duplicates"`
}

func (r LogCustomFieldRuleActionParametersRawResponseFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The raw request field to log.
type LogCustomFieldRuleActionParametersRequestFieldParam struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r LogCustomFieldRuleActionParametersRequestFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The transformed response field to log.
type LogCustomFieldRuleActionParametersResponseFieldParam struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
	// Whether to log duplicate values of the same header.
	PreserveDuplicates param.Field[bool] `json:"preserve_duplicates"`
}

func (r LogCustomFieldRuleActionParametersResponseFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The transformed request field to log.
type LogCustomFieldRuleActionParametersTransformedRequestFieldParam struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r LogCustomFieldRuleActionParametersTransformedRequestFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type LogCustomFieldRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r LogCustomFieldRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type LogCustomFieldRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r LogCustomFieldRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action LogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck LogRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit LogRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string      `json:"ref"`
	JSON logRuleJSON `json:"-"`
}

// logRuleJSON contains the JSON metadata for the struct [LogRule]
type logRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *LogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logRuleJSON) RawJSON() string {
	return r.raw
}

func (r LogRule) implementsRulesetNewResponseRule() {}

func (r LogRule) implementsRulesetUpdateResponseRule() {}

func (r LogRule) implementsRulesetGetResponseRule() {}

func (r LogRule) implementsPhaseUpdateResponseRule() {}

func (r LogRule) implementsPhaseGetResponseRule() {}

func (r LogRule) implementsPhaseVersionGetResponseRule() {}

func (r LogRule) implementsRuleNewResponseRule() {}

func (r LogRule) implementsRuleDeleteResponseRule() {}

func (r LogRule) implementsRuleEditResponseRule() {}

func (r LogRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type LogRuleAction string

const (
	LogRuleActionLog LogRuleAction = "log"
)

func (r LogRuleAction) IsKnown() bool {
	switch r {
	case LogRuleActionLog:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type LogRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                            `json:"username_expression,required"`
	JSON               logRuleExposedCredentialCheckJSON `json:"-"`
}

// logRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [LogRuleExposedCredentialCheck]
type logRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LogRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type LogRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string               `json:"score_response_header_name"`
	JSON                    logRuleRatelimitJSON `json:"-"`
}

// logRuleRatelimitJSON contains the JSON metadata for the struct
// [LogRuleRatelimit]
type logRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LogRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type LogRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[LogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[LogRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[LogRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r LogRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LogRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r LogRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r LogRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// Configure checks for exposed credentials.
type LogRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r LogRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type LogRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r LogRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type Logging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool        `json:"enabled,required"`
	JSON    loggingJSON `json:"-"`
}

// loggingJSON contains the JSON metadata for the struct [Logging]
type loggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Logging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loggingJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type LoggingParam struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r LoggingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ManagedChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck ManagedChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit ManagedChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                   `json:"ref"`
	JSON managedChallengeRuleJSON `json:"-"`
}

// managedChallengeRuleJSON contains the JSON metadata for the struct
// [ManagedChallengeRule]
type managedChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r ManagedChallengeRule) implementsRulesetNewResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetUpdateResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetGetResponseRule() {}

func (r ManagedChallengeRule) implementsPhaseUpdateResponseRule() {}

func (r ManagedChallengeRule) implementsPhaseGetResponseRule() {}

func (r ManagedChallengeRule) implementsPhaseVersionGetResponseRule() {}

func (r ManagedChallengeRule) implementsRuleNewResponseRule() {}

func (r ManagedChallengeRule) implementsRuleDeleteResponseRule() {}

func (r ManagedChallengeRule) implementsRuleEditResponseRule() {}

func (r ManagedChallengeRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type ManagedChallengeRuleAction string

const (
	ManagedChallengeRuleActionManagedChallenge ManagedChallengeRuleAction = "managed_challenge"
)

func (r ManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case ManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type ManagedChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                         `json:"username_expression,required"`
	JSON               managedChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// managedChallengeRuleExposedCredentialCheckJSON contains the JSON metadata for
// the struct [ManagedChallengeRuleExposedCredentialCheck]
type managedChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ManagedChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type ManagedChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                            `json:"score_response_header_name"`
	JSON                    managedChallengeRuleRatelimitJSON `json:"-"`
}

// managedChallengeRuleRatelimitJSON contains the JSON metadata for the struct
// [ManagedChallengeRuleRatelimit]
type managedChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ManagedChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type ManagedChallengeRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ManagedChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ManagedChallengeRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ManagedChallengeRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ManagedChallengeRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ManagedChallengeRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r ManagedChallengeRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r ManagedChallengeRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// Configure checks for exposed credentials.
type ManagedChallengeRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ManagedChallengeRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ManagedChallengeRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r ManagedChallengeRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RedirectRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RedirectRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RedirectRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string           `json:"ref"`
	JSON redirectRuleJSON `json:"-"`
}

// redirectRuleJSON contains the JSON metadata for the struct [RedirectRule]
type redirectRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRule) implementsRulesetNewResponseRule() {}

func (r RedirectRule) implementsRulesetUpdateResponseRule() {}

func (r RedirectRule) implementsRulesetGetResponseRule() {}

func (r RedirectRule) implementsPhaseUpdateResponseRule() {}

func (r RedirectRule) implementsPhaseGetResponseRule() {}

func (r RedirectRule) implementsPhaseVersionGetResponseRule() {}

func (r RedirectRule) implementsRuleNewResponseRule() {}

func (r RedirectRule) implementsRuleDeleteResponseRule() {}

func (r RedirectRule) implementsRuleEditResponseRule() {}

func (r RedirectRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RedirectRuleAction string

const (
	RedirectRuleActionRedirect RedirectRuleAction = "redirect"
)

func (r RedirectRuleAction) IsKnown() bool {
	switch r {
	case RedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList RedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue RedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      redirectRuleActionParametersJSON      `json:"-"`
}

// redirectRuleActionParametersJSON contains the JSON metadata for the struct
// [RedirectRuleActionParameters]
type redirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type RedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                   `json:"name"`
	JSON redirectRuleActionParametersFromListJSON `json:"-"`
}

// redirectRuleActionParametersFromListJSON contains the JSON metadata for the
// struct [RedirectRuleActionParametersFromList]
type redirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type RedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode RedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL RedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      redirectRuleActionParametersFromValueJSON      `json:"-"`
}

// redirectRuleActionParametersFromValueJSON contains the JSON metadata for the
// struct [RedirectRuleActionParametersFromValue]
type redirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type RedirectRuleActionParametersFromValueStatusCode float64

const (
	RedirectRuleActionParametersFromValueStatusCode301 RedirectRuleActionParametersFromValueStatusCode = 301
	RedirectRuleActionParametersFromValueStatusCode302 RedirectRuleActionParametersFromValueStatusCode = 302
	RedirectRuleActionParametersFromValueStatusCode303 RedirectRuleActionParametersFromValueStatusCode = 303
	RedirectRuleActionParametersFromValueStatusCode307 RedirectRuleActionParametersFromValueStatusCode = 307
	RedirectRuleActionParametersFromValueStatusCode308 RedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RedirectRuleActionParametersFromValueStatusCode301, RedirectRuleActionParametersFromValueStatusCode302, RedirectRuleActionParametersFromValueStatusCode303, RedirectRuleActionParametersFromValueStatusCode307, RedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RedirectRuleActionParametersFromValueTargetURL struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string `json:"expression"`
	// The URL to redirect the request to.
	Value string                                             `json:"value"`
	JSON  redirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union RedirectRuleActionParametersFromValueTargetURLUnion
}

// redirectRuleActionParametersFromValueTargetURLJSON contains the JSON metadata
// for the struct [RedirectRuleActionParametersFromValueTargetURL]
type redirectRuleActionParametersFromValueTargetURLJSON struct {
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r redirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *RedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	*r = RedirectRuleActionParametersFromValueTargetURL{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RedirectRuleActionParametersFromValueTargetURLUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect],
// [RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
func (r RedirectRuleActionParametersFromValueTargetURL) AsUnion() RedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect] or
// [RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type RedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                              `json:"value"`
	JSON  redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON contains the
// JSON metadata for the struct
// [RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRedirectRuleActionParametersFromValueTargetURL() {
}

type RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                               `json:"expression"`
	JSON       redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON contains
// the JSON metadata for the struct
// [RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRedirectRuleActionParametersFromValueTargetURL() {
}

// Configure checks for exposed credentials.
type RedirectRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                 `json:"username_expression,required"`
	JSON               redirectRuleExposedCredentialCheckJSON `json:"-"`
}

// redirectRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [RedirectRuleExposedCredentialCheck]
type redirectRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RedirectRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RedirectRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                    `json:"score_response_header_name"`
	JSON                    redirectRuleRatelimitJSON `json:"-"`
}

// redirectRuleRatelimitJSON contains the JSON metadata for the struct
// [RedirectRuleRatelimit]
type redirectRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RedirectRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RedirectRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RedirectRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RedirectRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RedirectRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RedirectRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RedirectRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r RedirectRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r RedirectRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RedirectRuleActionParametersParam struct {
	// Serve a redirect based on a bulk list lookup.
	FromList param.Field[RedirectRuleActionParametersFromListParam] `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue param.Field[RedirectRuleActionParametersFromValueParam] `json:"from_value"`
}

func (r RedirectRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on a bulk list lookup.
type RedirectRuleActionParametersFromListParam struct {
	// Expression that evaluates to the list lookup key.
	Key param.Field[string] `json:"key"`
	// The name of the list to match against.
	Name param.Field[string] `json:"name"`
}

func (r RedirectRuleActionParametersFromListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on the request properties.
type RedirectRuleActionParametersFromValueParam struct {
	// Keep the query string of the original request.
	PreserveQueryString param.Field[bool] `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode param.Field[RedirectRuleActionParametersFromValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL param.Field[RedirectRuleActionParametersFromValueTargetURLUnionParam] `json:"target_url"`
}

func (r RedirectRuleActionParametersFromValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL to redirect the request to.
type RedirectRuleActionParametersFromValueTargetURLParam struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r RedirectRuleActionParametersFromValueTargetURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLParam) implementsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

// The URL to redirect the request to.
//
// Satisfied by
// [rulesets.RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam],
// [rulesets.RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam],
// [RedirectRuleActionParametersFromValueTargetURLParam].
type RedirectRuleActionParametersFromValueTargetURLUnionParam interface {
	implementsRedirectRuleActionParametersFromValueTargetURLUnionParam()
}

type RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam) implementsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

type RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam) implementsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

// Configure checks for exposed credentials.
type RedirectRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RedirectRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RedirectRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RedirectRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RewriteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RewriteRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RewriteRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string          `json:"ref"`
	JSON rewriteRuleJSON `json:"-"`
}

// rewriteRuleJSON contains the JSON metadata for the struct [RewriteRule]
type rewriteRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRule) implementsRulesetNewResponseRule() {}

func (r RewriteRule) implementsRulesetUpdateResponseRule() {}

func (r RewriteRule) implementsRulesetGetResponseRule() {}

func (r RewriteRule) implementsPhaseUpdateResponseRule() {}

func (r RewriteRule) implementsPhaseGetResponseRule() {}

func (r RewriteRule) implementsPhaseVersionGetResponseRule() {}

func (r RewriteRule) implementsRuleNewResponseRule() {}

func (r RewriteRule) implementsRuleDeleteResponseRule() {}

func (r RewriteRule) implementsRuleEditResponseRule() {}

func (r RewriteRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RewriteRuleAction string

const (
	RewriteRuleActionRewrite RewriteRuleAction = "rewrite"
)

func (r RewriteRuleAction) IsKnown() bool {
	switch r {
	case RewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]RewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  RewriteRuleActionParametersURI  `json:"uri"`
	JSON rewriteRuleActionParametersJSON `json:"-"`
}

// rewriteRuleActionParametersJSON contains the JSON metadata for the struct
// [RewriteRuleActionParameters]
type rewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type RewriteRuleActionParametersHeader struct {
	Operation RewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Expression for the header value.
	Expression string `json:"expression"`
	// Static value for the header.
	Value string                                `json:"value"`
	JSON  rewriteRuleActionParametersHeaderJSON `json:"-"`
	union RewriteRuleActionParametersHeadersUnion
}

// rewriteRuleActionParametersHeaderJSON contains the JSON metadata for the struct
// [RewriteRuleActionParametersHeader]
type rewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *RewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	*r = RewriteRuleActionParametersHeader{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RewriteRuleActionParametersHeadersUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RewriteRuleActionParametersHeadersRemoveHeader],
// [RewriteRuleActionParametersHeadersAddStaticHeader],
// [RewriteRuleActionParametersHeadersSetStaticHeader],
// [RewriteRuleActionParametersHeadersAddDynamicHeader],
// [RewriteRuleActionParametersHeadersSetDynamicHeader].
func (r RewriteRuleActionParametersHeader) AsUnion() RewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by [RewriteRuleActionParametersHeadersRemoveHeader],
// [RewriteRuleActionParametersHeadersAddStaticHeader],
// [RewriteRuleActionParametersHeadersSetStaticHeader],
// [RewriteRuleActionParametersHeadersAddDynamicHeader] or
// [RewriteRuleActionParametersHeadersSetDynamicHeader].
type RewriteRuleActionParametersHeadersUnion interface {
	implementsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersAddStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersSetStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersAddDynamicHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersSetDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type RewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation RewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      rewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// rewriteRuleActionParametersHeadersRemoveHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersRemoveHeader]
type rewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersRemoveHeader) implementsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Add a request header with a static value.
type RewriteRuleActionParametersHeadersAddStaticHeader struct {
	Operation RewriteRuleActionParametersHeadersAddStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                `json:"value,required"`
	JSON  rewriteRuleActionParametersHeadersAddStaticHeaderJSON `json:"-"`
}

// rewriteRuleActionParametersHeadersAddStaticHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersAddStaticHeader]
type rewriteRuleActionParametersHeadersAddStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersAddStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersAddStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersAddStaticHeader) implementsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersAddStaticHeaderOperation string

const (
	RewriteRuleActionParametersHeadersAddStaticHeaderOperationAdd RewriteRuleActionParametersHeadersAddStaticHeaderOperation = "add"
)

func (r RewriteRuleActionParametersHeadersAddStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersAddStaticHeaderOperationAdd:
		return true
	}
	return false
}

// Set a request header with a static value.
type RewriteRuleActionParametersHeadersSetStaticHeader struct {
	Operation RewriteRuleActionParametersHeadersSetStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                `json:"value,required"`
	JSON  rewriteRuleActionParametersHeadersSetStaticHeaderJSON `json:"-"`
}

// rewriteRuleActionParametersHeadersSetStaticHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersSetStaticHeader]
type rewriteRuleActionParametersHeadersSetStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersSetStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersSetStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersSetStaticHeader) implementsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersSetStaticHeaderOperation string

const (
	RewriteRuleActionParametersHeadersSetStaticHeaderOperationSet RewriteRuleActionParametersHeadersSetStaticHeaderOperation = "set"
)

func (r RewriteRuleActionParametersHeadersSetStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersSetStaticHeaderOperationSet:
		return true
	}
	return false
}

// Add a request header with a dynamic value.
type RewriteRuleActionParametersHeadersAddDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                      `json:"expression,required"`
	Operation  RewriteRuleActionParametersHeadersAddDynamicHeaderOperation `json:"operation,required"`
	JSON       rewriteRuleActionParametersHeadersAddDynamicHeaderJSON      `json:"-"`
}

// rewriteRuleActionParametersHeadersAddDynamicHeaderJSON contains the JSON
// metadata for the struct [RewriteRuleActionParametersHeadersAddDynamicHeader]
type rewriteRuleActionParametersHeadersAddDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersAddDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersAddDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersAddDynamicHeader) implementsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersAddDynamicHeaderOperation string

const (
	RewriteRuleActionParametersHeadersAddDynamicHeaderOperationAdd RewriteRuleActionParametersHeadersAddDynamicHeaderOperation = "add"
)

func (r RewriteRuleActionParametersHeadersAddDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersAddDynamicHeaderOperationAdd:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RewriteRuleActionParametersHeadersSetDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                      `json:"expression,required"`
	Operation  RewriteRuleActionParametersHeadersSetDynamicHeaderOperation `json:"operation,required"`
	JSON       rewriteRuleActionParametersHeadersSetDynamicHeaderJSON      `json:"-"`
}

// rewriteRuleActionParametersHeadersSetDynamicHeaderJSON contains the JSON
// metadata for the struct [RewriteRuleActionParametersHeadersSetDynamicHeader]
type rewriteRuleActionParametersHeadersSetDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersSetDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersSetDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersSetDynamicHeader) implementsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersSetDynamicHeaderOperation string

const (
	RewriteRuleActionParametersHeadersSetDynamicHeaderOperationSet RewriteRuleActionParametersHeadersSetDynamicHeaderOperation = "set"
)

func (r RewriteRuleActionParametersHeadersSetDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersSetDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RewriteRuleActionParametersHeadersOperation string

const (
	RewriteRuleActionParametersHeadersOperationRemove RewriteRuleActionParametersHeadersOperation = "remove"
	RewriteRuleActionParametersHeadersOperationAdd    RewriteRuleActionParametersHeadersOperation = "add"
	RewriteRuleActionParametersHeadersOperationSet    RewriteRuleActionParametersHeadersOperation = "set"
)

func (r RewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersOperationRemove, RewriteRuleActionParametersHeadersOperationAdd, RewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path RewriteURIPart `json:"path"`
	// Query portion rewrite.
	Query RewriteURIPart                     `json:"query"`
	JSON  rewriteRuleActionParametersURIJSON `json:"-"`
}

// rewriteRuleActionParametersURIJSON contains the JSON metadata for the struct
// [RewriteRuleActionParametersURI]
type rewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type RewriteRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                `json:"username_expression,required"`
	JSON               rewriteRuleExposedCredentialCheckJSON `json:"-"`
}

// rewriteRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [RewriteRuleExposedCredentialCheck]
type rewriteRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RewriteRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RewriteRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                   `json:"score_response_header_name"`
	JSON                    rewriteRuleRatelimitJSON `json:"-"`
}

// rewriteRuleRatelimitJSON contains the JSON metadata for the struct
// [RewriteRuleRatelimit]
type rewriteRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RewriteRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RewriteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RewriteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RewriteRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RewriteRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RewriteRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RewriteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r RewriteRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r RewriteRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RewriteRuleActionParametersParam struct {
	// Map of request headers to modify.
	Headers param.Field[map[string]RewriteRuleActionParametersHeadersUnionParam] `json:"headers"`
	// URI to rewrite the request to.
	URI param.Field[RewriteRuleActionParametersURIParam] `json:"uri"`
}

func (r RewriteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Remove the header from the request.
type RewriteRuleActionParametersHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersOperation] `json:"operation,required"`
	// Expression for the header value.
	Expression param.Field[string] `json:"expression"`
	// Static value for the header.
	Value param.Field[string] `json:"value"`
}

func (r RewriteRuleActionParametersHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// Remove the header from the request.
//
// Satisfied by [rulesets.RewriteRuleActionParametersHeadersRemoveHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersAddStaticHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersSetStaticHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersAddDynamicHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersSetDynamicHeaderParam],
// [RewriteRuleActionParametersHeaderParam].
type RewriteRuleActionParametersHeadersUnionParam interface {
	implementsRewriteRuleActionParametersHeadersUnionParam()
}

// Remove the header from the request.
type RewriteRuleActionParametersHeadersRemoveHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersRemoveHeaderOperation] `json:"operation,required"`
}

func (r RewriteRuleActionParametersHeadersRemoveHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersRemoveHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// Add a request header with a static value.
type RewriteRuleActionParametersHeadersAddStaticHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersAddStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r RewriteRuleActionParametersHeadersAddStaticHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersAddStaticHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// Set a request header with a static value.
type RewriteRuleActionParametersHeadersSetStaticHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersSetStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r RewriteRuleActionParametersHeadersSetStaticHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersSetStaticHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// Add a request header with a dynamic value.
type RewriteRuleActionParametersHeadersAddDynamicHeaderParam struct {
	// Expression for the header value.
	Expression param.Field[string]                                                      `json:"expression,required"`
	Operation  param.Field[RewriteRuleActionParametersHeadersAddDynamicHeaderOperation] `json:"operation,required"`
}

func (r RewriteRuleActionParametersHeadersAddDynamicHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersAddDynamicHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// Set a request header with a dynamic value.
type RewriteRuleActionParametersHeadersSetDynamicHeaderParam struct {
	// Expression for the header value.
	Expression param.Field[string]                                                      `json:"expression,required"`
	Operation  param.Field[RewriteRuleActionParametersHeadersSetDynamicHeaderOperation] `json:"operation,required"`
}

func (r RewriteRuleActionParametersHeadersSetDynamicHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersSetDynamicHeaderParam) implementsRewriteRuleActionParametersHeadersUnionParam() {
}

// URI to rewrite the request to.
type RewriteRuleActionParametersURIParam struct {
	// Path portion rewrite.
	Path param.Field[RewriteURIPartUnionParam] `json:"path"`
	// Query portion rewrite.
	Query param.Field[RewriteURIPartUnionParam] `json:"query"`
}

func (r RewriteRuleActionParametersURIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type RewriteRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RewriteRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RewriteRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RewriteRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RewriteURIPart struct {
	// Expression to evaluate for the replacement value.
	Expression string `json:"expression"`
	// Predefined replacement value.
	Value string             `json:"value"`
	JSON  rewriteURIPartJSON `json:"-"`
	union RewriteURIPartUnion
}

// rewriteURIPartJSON contains the JSON metadata for the struct [RewriteURIPart]
type rewriteURIPartJSON struct {
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rewriteURIPartJSON) RawJSON() string {
	return r.raw
}

func (r *RewriteURIPart) UnmarshalJSON(data []byte) (err error) {
	*r = RewriteURIPart{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RewriteURIPartUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [RewriteURIPartStaticValue],
// [RewriteURIPartDynamicValue].
func (r RewriteURIPart) AsUnion() RewriteURIPartUnion {
	return r.union
}

// Union satisfied by [RewriteURIPartStaticValue] or [RewriteURIPartDynamicValue].
type RewriteURIPartUnion interface {
	implementsRewriteURIPart()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RewriteURIPartUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteURIPartStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteURIPartDynamicValue{}),
		},
	)
}

type RewriteURIPartStaticValue struct {
	// Predefined replacement value.
	Value string                        `json:"value,required"`
	JSON  rewriteURIPartStaticValueJSON `json:"-"`
}

// rewriteURIPartStaticValueJSON contains the JSON metadata for the struct
// [RewriteURIPartStaticValue]
type rewriteURIPartStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteURIPartStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteURIPartStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RewriteURIPartStaticValue) implementsRewriteURIPart() {}

type RewriteURIPartDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                         `json:"expression,required"`
	JSON       rewriteURIPartDynamicValueJSON `json:"-"`
}

// rewriteURIPartDynamicValueJSON contains the JSON metadata for the struct
// [RewriteURIPartDynamicValue]
type rewriteURIPartDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteURIPartDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteURIPartDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RewriteURIPartDynamicValue) implementsRewriteURIPart() {}

type RewriteURIPartParam struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
}

func (r RewriteURIPartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartParam) implementsRewriteURIPartUnionParam() {}

// Satisfied by [rulesets.RewriteURIPartStaticValueParam],
// [rulesets.RewriteURIPartDynamicValueParam], [RewriteURIPartParam].
type RewriteURIPartUnionParam interface {
	implementsRewriteURIPartUnionParam()
}

type RewriteURIPartStaticValueParam struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RewriteURIPartStaticValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartStaticValueParam) implementsRewriteURIPartUnionParam() {}

type RewriteURIPartDynamicValueParam struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RewriteURIPartDynamicValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartDynamicValueParam) implementsRewriteURIPartUnionParam() {}

type RouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RouteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RouteRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RouteRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string        `json:"ref"`
	JSON routeRuleJSON `json:"-"`
}

// routeRuleJSON contains the JSON metadata for the struct [RouteRule]
type routeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RouteRule) implementsRulesetNewResponseRule() {}

func (r RouteRule) implementsRulesetUpdateResponseRule() {}

func (r RouteRule) implementsRulesetGetResponseRule() {}

func (r RouteRule) implementsPhaseUpdateResponseRule() {}

func (r RouteRule) implementsPhaseGetResponseRule() {}

func (r RouteRule) implementsPhaseVersionGetResponseRule() {}

func (r RouteRule) implementsRuleNewResponseRule() {}

func (r RouteRule) implementsRuleDeleteResponseRule() {}

func (r RouteRule) implementsRuleEditResponseRule() {}

func (r RouteRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RouteRuleAction string

const (
	RouteRuleActionRoute RouteRuleAction = "route"
)

func (r RouteRuleAction) IsKnown() bool {
	switch r {
	case RouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin RouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	SNI  RouteRuleActionParametersSNI  `json:"sni"`
	JSON routeRuleActionParametersJSON `json:"-"`
}

// routeRuleActionParametersJSON contains the JSON metadata for the struct
// [RouteRuleActionParameters]
type routeRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	SNI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type RouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                             `json:"port"`
	JSON routeRuleActionParametersOriginJSON `json:"-"`
}

// routeRuleActionParametersOriginJSON contains the JSON metadata for the struct
// [RouteRuleActionParametersOrigin]
type routeRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type RouteRuleActionParametersSNI struct {
	// The SNI override.
	Value string                           `json:"value,required"`
	JSON  routeRuleActionParametersSNIJSON `json:"-"`
}

// routeRuleActionParametersSNIJSON contains the JSON metadata for the struct
// [RouteRuleActionParametersSNI]
type routeRuleActionParametersSNIJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParametersSNI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersSNIJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type RouteRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                              `json:"username_expression,required"`
	JSON               routeRuleExposedCredentialCheckJSON `json:"-"`
}

// routeRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [RouteRuleExposedCredentialCheck]
type routeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RouteRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RouteRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                 `json:"score_response_header_name"`
	JSON                    routeRuleRatelimitJSON `json:"-"`
}

// routeRuleRatelimitJSON contains the JSON metadata for the struct
// [RouteRuleRatelimit]
type routeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RouteRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RouteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RouteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RouteRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RouteRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RouteRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RouteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RouteRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r RouteRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r RouteRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RouteRuleActionParametersParam struct {
	// Rewrite the HTTP Host header.
	HostHeader param.Field[string] `json:"host_header"`
	// Override the IP/TCP destination.
	Origin param.Field[RouteRuleActionParametersOriginParam] `json:"origin"`
	// Override the Server Name Indication (SNI).
	SNI param.Field[RouteRuleActionParametersSNIParam] `json:"sni"`
}

func (r RouteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the IP/TCP destination.
type RouteRuleActionParametersOriginParam struct {
	// Override the resolved hostname.
	Host param.Field[string] `json:"host"`
	// Override the destination port.
	Port param.Field[float64] `json:"port"`
}

func (r RouteRuleActionParametersOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the Server Name Indication (SNI).
type RouteRuleActionParametersSNIParam struct {
	// The SNI override.
	Value param.Field[string] `json:"value,required"`
}

func (r RouteRuleActionParametersSNIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type RouteRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RouteRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RouteRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RouteRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ScoreRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck ScoreRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit ScoreRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string        `json:"ref"`
	JSON scoreRuleJSON `json:"-"`
}

// scoreRuleJSON contains the JSON metadata for the struct [ScoreRule]
type scoreRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r ScoreRule) implementsRulesetNewResponseRule() {}

func (r ScoreRule) implementsRulesetUpdateResponseRule() {}

func (r ScoreRule) implementsRulesetGetResponseRule() {}

func (r ScoreRule) implementsPhaseUpdateResponseRule() {}

func (r ScoreRule) implementsPhaseGetResponseRule() {}

func (r ScoreRule) implementsPhaseVersionGetResponseRule() {}

func (r ScoreRule) implementsRuleNewResponseRule() {}

func (r ScoreRule) implementsRuleDeleteResponseRule() {}

func (r ScoreRule) implementsRuleEditResponseRule() {}

func (r ScoreRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type ScoreRuleAction string

const (
	ScoreRuleActionScore ScoreRuleAction = "score"
)

func (r ScoreRuleAction) IsKnown() bool {
	switch r {
	case ScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                         `json:"increment"`
	JSON      scoreRuleActionParametersJSON `json:"-"`
}

// scoreRuleActionParametersJSON contains the JSON metadata for the struct
// [ScoreRuleActionParameters]
type scoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type ScoreRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                              `json:"username_expression,required"`
	JSON               scoreRuleExposedCredentialCheckJSON `json:"-"`
}

// scoreRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [ScoreRuleExposedCredentialCheck]
type scoreRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScoreRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type ScoreRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                 `json:"score_response_header_name"`
	JSON                    scoreRuleRatelimitJSON `json:"-"`
}

// scoreRuleRatelimitJSON contains the JSON metadata for the struct
// [ScoreRuleRatelimit]
type scoreRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ScoreRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type ScoreRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ScoreRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ScoreRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ScoreRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ScoreRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ScoreRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScoreRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r ScoreRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r ScoreRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type ScoreRuleActionParametersParam struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment param.Field[int64] `json:"increment"`
}

func (r ScoreRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ScoreRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ScoreRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ScoreRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r ScoreRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ServeErrorRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck ServeErrorRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit ServeErrorRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string             `json:"ref"`
	JSON serveErrorRuleJSON `json:"-"`
}

// serveErrorRuleJSON contains the JSON metadata for the struct [ServeErrorRule]
type serveErrorRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r ServeErrorRule) implementsRulesetNewResponseRule() {}

func (r ServeErrorRule) implementsRulesetUpdateResponseRule() {}

func (r ServeErrorRule) implementsRulesetGetResponseRule() {}

func (r ServeErrorRule) implementsPhaseUpdateResponseRule() {}

func (r ServeErrorRule) implementsPhaseGetResponseRule() {}

func (r ServeErrorRule) implementsPhaseVersionGetResponseRule() {}

func (r ServeErrorRule) implementsRuleNewResponseRule() {}

func (r ServeErrorRule) implementsRuleDeleteResponseRule() {}

func (r ServeErrorRule) implementsRuleEditResponseRule() {}

func (r ServeErrorRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type ServeErrorRuleAction string

const (
	ServeErrorRuleActionServeError ServeErrorRuleAction = "serve_error"
)

func (r ServeErrorRuleAction) IsKnown() bool {
	switch r {
	case ServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType ServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                            `json:"status_code"`
	JSON       serveErrorRuleActionParametersJSON `json:"-"`
}

// serveErrorRuleActionParametersJSON contains the JSON metadata for the struct
// [ServeErrorRuleActionParameters]
type serveErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type ServeErrorRuleActionParametersContentType string

const (
	ServeErrorRuleActionParametersContentTypeApplicationJson ServeErrorRuleActionParametersContentType = "application/json"
	ServeErrorRuleActionParametersContentTypeTextXml         ServeErrorRuleActionParametersContentType = "text/xml"
	ServeErrorRuleActionParametersContentTypeTextPlain       ServeErrorRuleActionParametersContentType = "text/plain"
	ServeErrorRuleActionParametersContentTypeTextHTML        ServeErrorRuleActionParametersContentType = "text/html"
)

func (r ServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case ServeErrorRuleActionParametersContentTypeApplicationJson, ServeErrorRuleActionParametersContentTypeTextXml, ServeErrorRuleActionParametersContentTypeTextPlain, ServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type ServeErrorRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                   `json:"username_expression,required"`
	JSON               serveErrorRuleExposedCredentialCheckJSON `json:"-"`
}

// serveErrorRuleExposedCredentialCheckJSON contains the JSON metadata for the
// struct [ServeErrorRuleExposedCredentialCheck]
type serveErrorRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ServeErrorRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type ServeErrorRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                      `json:"score_response_header_name"`
	JSON                    serveErrorRuleRatelimitJSON `json:"-"`
}

// serveErrorRuleRatelimitJSON contains the JSON metadata for the struct
// [ServeErrorRuleRatelimit]
type serveErrorRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ServeErrorRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type ServeErrorRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ServeErrorRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ServeErrorRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ServeErrorRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ServeErrorRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ServeErrorRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServeErrorRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r ServeErrorRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r ServeErrorRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type ServeErrorRuleActionParametersParam struct {
	// Error response content.
	Content param.Field[string] `json:"content"`
	// Content-type header to set with the response.
	ContentType param.Field[ServeErrorRuleActionParametersContentType] `json:"content_type"`
	// The status code to use for the error.
	StatusCode param.Field[float64] `json:"status_code"`
}

func (r ServeErrorRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ServeErrorRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ServeErrorRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ServeErrorRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r ServeErrorRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SetCacheSettingsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck SetCacheSettingsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit SetCacheSettingsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                   `json:"ref"`
	JSON setCacheSettingsRuleJSON `json:"-"`
}

// setCacheSettingsRuleJSON contains the JSON metadata for the struct
// [SetCacheSettingsRule]
type setCacheSettingsRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRule) implementsRulesetNewResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetUpdateResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetGetResponseRule() {}

func (r SetCacheSettingsRule) implementsPhaseUpdateResponseRule() {}

func (r SetCacheSettingsRule) implementsPhaseGetResponseRule() {}

func (r SetCacheSettingsRule) implementsPhaseVersionGetResponseRule() {}

func (r SetCacheSettingsRule) implementsRuleNewResponseRule() {}

func (r SetCacheSettingsRule) implementsRuleDeleteResponseRule() {}

func (r SetCacheSettingsRule) implementsRuleEditResponseRule() {}

func (r SetCacheSettingsRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type SetCacheSettingsRuleAction string

const (
	SetCacheSettingsRuleActionSetCacheSettings SetCacheSettingsRuleAction = "set_cache_settings"
)

func (r SetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL SetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey SetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve SetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL SetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl bool `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin.
	OriginErrorPagePassthru bool `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout int64 `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags bool `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale SetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       setCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// setCacheSettingsRuleActionParametersJSON contains the JSON metadata for the
// struct [SetCacheSettingsRuleActionParameters]
type setCacheSettingsRuleActionParametersJSON struct {
	AdditionalCacheablePorts apijson.Field
	BrowserTTL               apijson.Field
	Cache                    apijson.Field
	CacheKey                 apijson.Field
	CacheReserve             apijson.Field
	EdgeTTL                  apijson.Field
	OriginCacheControl       apijson.Field
	OriginErrorPagePassthru  apijson.Field
	ReadTimeout              apijson.Field
	RespectStrongEtags       apijson.Field
	ServeStale               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type SetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode SetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                              `json:"default"`
	JSON    setCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersBrowserTTLJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersBrowserTTL]
type setCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type SetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	SetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   SetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	SetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault SetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	SetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  SetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
	SetCacheSettingsRuleActionParametersBrowserTTLModeBypass          SetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass"
)

func (r SetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, SetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, SetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin, SetCacheSettingsRuleActionParametersBrowserTTLModeBypass:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type SetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type.
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached.
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey SetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                             `json:"ignore_query_strings_order"`
	JSON                    setCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyJSON contains the JSON metadata for
// the struct [SetCacheSettingsRuleActionParametersCacheKey]
type setCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence of parameters in the query string to build the cache key.
	QueryString SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON contains the JSON
// metadata for the struct [SetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                        `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON contains the
// JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// For each header name and list of values combination, check if the request header
	// contains any of the values provided. The presence of the request header and
	// whether any of the values provided are contained in the request header value is
	// used in building the cache key.
	Contains map[string][]string `json:"contains"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                        `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON contains the
// JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	Contains      apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                          `json:"resolved"`
	JSON     setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// This field can have the runtime type of
	// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude].
	Exclude interface{} `json:"exclude"`
	// This field can have the runtime type of
	// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude].
	Include interface{}                                                          `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON `json:"-"`
	union   SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnion
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON contains
// the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	*r = SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters].
func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) AsUnion() SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnion {
	return r.union
}

// Use the presence of parameters in the query string to build the cache key.
//
// Union satisfied by
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters]
// or
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnion interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters{}),
		},
	)
}

// Use the presence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters struct {
	// A list of query string parameters used to build the cache key.
	Include SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersJSON    `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersJSON struct {
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString() {
}

// A list of query string parameters used to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude struct {
	// Determines whether to include all query string parameters in the cache key.
	All bool `json:"all"`
	// This field can have the runtime type of [[]string].
	List  interface{}                                                                                              `json:"list"`
	JSON  setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeJSON `json:"-"`
	union SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnion
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeJSON) RawJSON() string {
	return r.raw
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude) UnmarshalJSON(data []byte) (err error) {
	*r = SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters].
func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude) AsUnion() SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnion {
	return r.union
}

// A list of query string parameters used to build the cache key.
//
// Union satisfied by
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters]
// or
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnion interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters{}),
		},
	)
}

// A list of query string parameters used to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters struct {
	List []string                                                                                                                          `json:"list"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersJSON struct {
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude() {
}

// Build the cache key using ALL query string parameters that are in the request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters struct {
	// Determines whether to include all query string parameters in the cache key.
	All  bool                                                                                                                             `json:"all"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersInclude() {
}

// Use the absence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters struct {
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	Exclude SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude `json:"exclude"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersJSON    `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersJSON struct {
	Exclude     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString() {
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude struct {
	// Determines whether to exclude all query string parameters from the cache key.
	All bool `json:"all"`
	// This field can have the runtime type of [[]string].
	List  interface{}                                                                                              `json:"list"`
	JSON  setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeJSON `json:"-"`
	union SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnion
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeJSON) RawJSON() string {
	return r.raw
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude) UnmarshalJSON(data []byte) (err error) {
	*r = SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters].
func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude) AsUnion() SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnion {
	return r.union
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
//
// Union satisfied by
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters]
// or
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnion interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters{}),
		},
	)
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters struct {
	List []string                                                                                                                          `json:"list"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersJSON struct {
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude() {
}

// Build the cache key excluding ALL query string parameters that are in the
// request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters struct {
	// Determines whether to exclude all query string parameters from the cache key.
	All  bool                                                                                                                             `json:"all"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParameters) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExclude() {
}

// Characteristics of the request user agent used in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                          `json:"lang"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type SetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinimumFileSize int64                                                `json:"minimum_file_size,required"`
	JSON            setCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheReserveJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersCacheReserve]
type setCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible        apijson.Field
	MinimumFileSize apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type SetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// Edge TTL options.
	Mode SetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode.
	StatusCodeTTL []SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          setCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLJSON contains the JSON metadata for
// the struct [SetCacheSettingsRuleActionParametersEdgeTTL]
type setCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// Edge TTL options.
type SetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	SetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   SetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	SetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault SetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	SetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  SetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r SetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, SetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, SetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the TTL for responses with this specific status code.
	StatusCodeValue int64                                                        `json:"status_code_value"`
	JSON            setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// Response status code lower bound.
	From int64 `json:"from,required"`
	// Response status code upper bound.
	To   int64                                                                       `json:"to,required"`
	JSON setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type SetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                               `json:"disable_stale_while_updating,required"`
	JSON                      setCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersServeStaleJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersServeStale]
type setCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// Configure checks for exposed credentials.
type SetCacheSettingsRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                         `json:"username_expression,required"`
	JSON               setCacheSettingsRuleExposedCredentialCheckJSON `json:"-"`
}

// setCacheSettingsRuleExposedCredentialCheckJSON contains the JSON metadata for
// the struct [SetCacheSettingsRuleExposedCredentialCheck]
type setCacheSettingsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SetCacheSettingsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type SetCacheSettingsRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                            `json:"score_response_header_name"`
	JSON                    setCacheSettingsRuleRatelimitJSON `json:"-"`
}

// setCacheSettingsRuleRatelimitJSON contains the JSON metadata for the struct
// [SetCacheSettingsRuleRatelimit]
type setCacheSettingsRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetCacheSettingsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type SetCacheSettingsRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SetCacheSettingsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SetCacheSettingsRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[SetCacheSettingsRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[SetCacheSettingsRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r SetCacheSettingsRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r SetCacheSettingsRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r SetCacheSettingsRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SetCacheSettingsRuleActionParametersParam struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts param.Field[[]int64] `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL param.Field[SetCacheSettingsRuleActionParametersBrowserTTLParam] `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache param.Field[bool] `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey param.Field[SetCacheSettingsRuleActionParametersCacheKeyParam] `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve param.Field[SetCacheSettingsRuleActionParametersCacheReserveParam] `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL param.Field[SetCacheSettingsRuleActionParametersEdgeTTLParam] `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl param.Field[bool] `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin.
	OriginErrorPagePassthru param.Field[bool] `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout param.Field[int64] `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags param.Field[bool] `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale param.Field[SetCacheSettingsRuleActionParametersServeStaleParam] `json:"serve_stale"`
}

func (r SetCacheSettingsRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type SetCacheSettingsRuleActionParametersBrowserTTLParam struct {
	// Determines which browser ttl mode to use.
	Mode param.Field[SetCacheSettingsRuleActionParametersBrowserTTLMode] `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default"`
}

func (r SetCacheSettingsRuleActionParametersBrowserTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type SetCacheSettingsRuleActionParametersCacheKeyParam struct {
	// Separate cached content based on the visitor’s device type.
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached.
	CacheDeceptionArmor param.Field[bool] `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam] `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder param.Field[bool] `json:"ignore_query_strings_order"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customize which components of the request are included or excluded from the
// cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam struct {
	// The cookies to include in building the cache key.
	Cookie param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam] `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam] `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam] `json:"host"`
	// Use the presence of parameters in the query string to build the cache key.
	QueryString param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam] `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam] `json:"user"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookies to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Include these cookies' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The header names and values to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// For each header name and list of values combination, check if the request header
	// contains any of the values provided. The presence of the request header and
	// whether any of the values provided are contained in the request header value is
	// used in building the cache key.
	Contains param.Field[map[string][]string] `json:"contains"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin param.Field[bool] `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to use the original host or the resolved host in the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Use the presence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam struct {
	Exclude param.Field[interface{}] `json:"exclude"`
	Include param.Field[interface{}] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam() {
}

// Use the presence of parameters in the query string to build the cache key.
//
// Satisfied by
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersParam],
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersParam],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam()
}

// Use the presence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersParam struct {
	// A list of query string parameters used to build the cache key.
	Include param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam() {
}

// A list of query string parameters used to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeParam struct {
	// Determines whether to include all query string parameters in the cache key.
	All  param.Field[bool]        `json:"all"`
	List param.Field[interface{}] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam() {
}

// A list of query string parameters used to build the cache key.
//
// Satisfied by
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersParam],
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersParam],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeParam].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam()
}

// A list of query string parameters used to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersParam struct {
	List param.Field[[]string] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeSomeQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam() {
}

// Build the cache key using ALL query string parameters that are in the request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersParam struct {
	// Determines whether to include all query string parameters in the cache key.
	All param.Field[bool] `json:"all"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeAllQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludedQueryStringParametersIncludeUnionParam() {
}

// Use the absence of parameters in the query string to build the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersParam struct {
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	Exclude param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam] `json:"exclude"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringUnionParam() {
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeParam struct {
	// Determines whether to exclude all query string parameters from the cache key.
	All  param.Field[bool]        `json:"all"`
	List param.Field[interface{}] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam() {
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
//
// Satisfied by
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersParam],
// [rulesets.SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersParam],
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeParam].
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam interface {
	implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam()
}

// A list of query string parameters NOT used to build the cache key. All
// parameters present in the request but missing in this list will be used to build
// the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersParam struct {
	List param.Field[[]string] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeSomeQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam() {
}

// Build the cache key excluding ALL query string parameters that are in the
// request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersParam struct {
	// Determines whether to exclude all query string parameters from the cache key.
	All param.Field[bool] `json:"all"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeAllQueryStringParametersParam) implementsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludedQueryStringParametersExcludeUnionParam() {
}

// Characteristics of the request user agent used in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam struct {
	// Use the user agent's device type in the cache key.
	DeviceType param.Field[bool] `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo param.Field[bool] `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang param.Field[bool] `json:"lang"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type SetCacheSettingsRuleActionParametersCacheReserveParam struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible param.Field[bool] `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinimumFileSize param.Field[int64] `json:"minimum_file_size,required"`
}

func (r SetCacheSettingsRuleActionParametersCacheReserveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type SetCacheSettingsRuleActionParametersEdgeTTLParam struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default,required"`
	// Edge TTL options.
	Mode param.Field[SetCacheSettingsRuleActionParametersEdgeTTLMode] `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode.
	StatusCodeTTL param.Field[[]SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam] `json:"status_code_ttl,required"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value param.Field[int64] `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange param.Field[SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam] `json:"status_code_range"`
	// Set the TTL for responses with this specific status code.
	StatusCodeValue param.Field[int64] `json:"status_code_value"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The range of status codes used to apply the selected mode.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam struct {
	// Response status code lower bound.
	From param.Field[int64] `json:"from,required"`
	// Response status code upper bound.
	To param.Field[int64] `json:"to,required"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type SetCacheSettingsRuleActionParametersServeStaleParam struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating param.Field[bool] `json:"disable_stale_while_updating,required"`
}

func (r SetCacheSettingsRuleActionParametersServeStaleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type SetCacheSettingsRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r SetCacheSettingsRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type SetCacheSettingsRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r SetCacheSettingsRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SetConfigRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck SetConfigRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit SetConfigRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string            `json:"ref"`
	JSON setConfigRuleJSON `json:"-"`
}

// setConfigRuleJSON contains the JSON metadata for the struct [SetConfigRule]
type setConfigRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r SetConfigRule) implementsRulesetNewResponseRule() {}

func (r SetConfigRule) implementsRulesetUpdateResponseRule() {}

func (r SetConfigRule) implementsRulesetGetResponseRule() {}

func (r SetConfigRule) implementsPhaseUpdateResponseRule() {}

func (r SetConfigRule) implementsPhaseGetResponseRule() {}

func (r SetConfigRule) implementsPhaseVersionGetResponseRule() {}

func (r SetConfigRule) implementsRuleNewResponseRule() {}

func (r SetConfigRule) implementsRuleDeleteResponseRule() {}

func (r SetConfigRule) implementsRuleEditResponseRule() {}

func (r SetConfigRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type SetConfigRuleAction string

const (
	SetConfigRuleActionSetConfig SetConfigRuleAction = "set_config"
)

func (r SetConfigRuleAction) IsKnown() bool {
	switch r {
	case SetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify SetConfigRuleActionParametersAutominify `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	BIC bool `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps SetConfigRuleActionParametersDisableApps `json:"disable_apps"`
	// Turn off Real User Monitoring (RUM).
	DisableRUM SetConfigRuleActionParametersDisableRUM `json:"disable_rum"`
	// Turn off Zaraz.
	DisableZaraz SetConfigRuleActionParametersDisableZaraz `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation bool `json:"email_obfuscation"`
	// Turn on or off Cloudflare Fonts.
	Fonts bool `json:"fonts"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection bool `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage bool `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption bool `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish SetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader.
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel SetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL SetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	SXG  bool                              `json:"sxg"`
	JSON setConfigRuleActionParametersJSON `json:"-"`
}

// setConfigRuleActionParametersJSON contains the JSON metadata for the struct
// [SetConfigRuleActionParameters]
type setConfigRuleActionParametersJSON struct {
	AutomaticHTTPSRewrites  apijson.Field
	Autominify              apijson.Field
	BIC                     apijson.Field
	DisableApps             apijson.Field
	DisableRUM              apijson.Field
	DisableZaraz            apijson.Field
	EmailObfuscation        apijson.Field
	Fonts                   apijson.Field
	HotlinkProtection       apijson.Field
	Mirage                  apijson.Field
	OpportunisticEncryption apijson.Field
	Polish                  apijson.Field
	RocketLoader            apijson.Field
	SecurityLevel           apijson.Field
	ServerSideExcludes      apijson.Field
	SSL                     apijson.Field
	SXG                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type SetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	CSS bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	JS   bool                                        `json:"js"`
	JSON setConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// setConfigRuleActionParametersAutominifyJSON contains the JSON metadata for the
// struct [SetConfigRuleActionParametersAutominify]
type setConfigRuleActionParametersAutominifyJSON struct {
	CSS         apijson.Field
	HTML        apijson.Field
	JS          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Turn off all active Cloudflare Apps.
type SetConfigRuleActionParametersDisableApps bool

const (
	SetConfigRuleActionParametersDisableAppsTrue SetConfigRuleActionParametersDisableApps = true
)

func (r SetConfigRuleActionParametersDisableApps) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableAppsTrue:
		return true
	}
	return false
}

// Turn off Real User Monitoring (RUM).
type SetConfigRuleActionParametersDisableRUM bool

const (
	SetConfigRuleActionParametersDisableRUMTrue SetConfigRuleActionParametersDisableRUM = true
)

func (r SetConfigRuleActionParametersDisableRUM) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableRUMTrue:
		return true
	}
	return false
}

// Turn off Zaraz.
type SetConfigRuleActionParametersDisableZaraz bool

const (
	SetConfigRuleActionParametersDisableZarazTrue SetConfigRuleActionParametersDisableZaraz = true
)

func (r SetConfigRuleActionParametersDisableZaraz) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableZarazTrue:
		return true
	}
	return false
}

// Configure the Polish level.
type SetConfigRuleActionParametersPolish string

const (
	SetConfigRuleActionParametersPolishOff      SetConfigRuleActionParametersPolish = "off"
	SetConfigRuleActionParametersPolishLossless SetConfigRuleActionParametersPolish = "lossless"
	SetConfigRuleActionParametersPolishLossy    SetConfigRuleActionParametersPolish = "lossy"
	SetConfigRuleActionParametersPolishWebP     SetConfigRuleActionParametersPolish = "webp"
)

func (r SetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersPolishOff, SetConfigRuleActionParametersPolishLossless, SetConfigRuleActionParametersPolishLossy, SetConfigRuleActionParametersPolishWebP:
		return true
	}
	return false
}

// Configure the Security Level.
type SetConfigRuleActionParametersSecurityLevel string

const (
	SetConfigRuleActionParametersSecurityLevelOff            SetConfigRuleActionParametersSecurityLevel = "off"
	SetConfigRuleActionParametersSecurityLevelEssentiallyOff SetConfigRuleActionParametersSecurityLevel = "essentially_off"
	SetConfigRuleActionParametersSecurityLevelLow            SetConfigRuleActionParametersSecurityLevel = "low"
	SetConfigRuleActionParametersSecurityLevelMedium         SetConfigRuleActionParametersSecurityLevel = "medium"
	SetConfigRuleActionParametersSecurityLevelHigh           SetConfigRuleActionParametersSecurityLevel = "high"
	SetConfigRuleActionParametersSecurityLevelUnderAttack    SetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r SetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersSecurityLevelOff, SetConfigRuleActionParametersSecurityLevelEssentiallyOff, SetConfigRuleActionParametersSecurityLevelLow, SetConfigRuleActionParametersSecurityLevelMedium, SetConfigRuleActionParametersSecurityLevelHigh, SetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type SetConfigRuleActionParametersSSL string

const (
	SetConfigRuleActionParametersSSLOff        SetConfigRuleActionParametersSSL = "off"
	SetConfigRuleActionParametersSSLFlexible   SetConfigRuleActionParametersSSL = "flexible"
	SetConfigRuleActionParametersSSLFull       SetConfigRuleActionParametersSSL = "full"
	SetConfigRuleActionParametersSSLStrict     SetConfigRuleActionParametersSSL = "strict"
	SetConfigRuleActionParametersSSLOriginPull SetConfigRuleActionParametersSSL = "origin_pull"
)

func (r SetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersSSLOff, SetConfigRuleActionParametersSSLFlexible, SetConfigRuleActionParametersSSLFull, SetConfigRuleActionParametersSSLStrict, SetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type SetConfigRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                  `json:"username_expression,required"`
	JSON               setConfigRuleExposedCredentialCheckJSON `json:"-"`
}

// setConfigRuleExposedCredentialCheckJSON contains the JSON metadata for the
// struct [SetConfigRuleExposedCredentialCheck]
type setConfigRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SetConfigRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type SetConfigRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                     `json:"score_response_header_name"`
	JSON                    setConfigRuleRatelimitJSON `json:"-"`
}

// setConfigRuleRatelimitJSON contains the JSON metadata for the struct
// [SetConfigRuleRatelimit]
type setConfigRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetConfigRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type SetConfigRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SetConfigRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SetConfigRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[SetConfigRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[SetConfigRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r SetConfigRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetConfigRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r SetConfigRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r SetConfigRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SetConfigRuleActionParametersParam struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites param.Field[bool] `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify param.Field[SetConfigRuleActionParametersAutominifyParam] `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	BIC param.Field[bool] `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps param.Field[SetConfigRuleActionParametersDisableApps] `json:"disable_apps"`
	// Turn off Real User Monitoring (RUM).
	DisableRUM param.Field[SetConfigRuleActionParametersDisableRUM] `json:"disable_rum"`
	// Turn off Zaraz.
	DisableZaraz param.Field[SetConfigRuleActionParametersDisableZaraz] `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation param.Field[bool] `json:"email_obfuscation"`
	// Turn on or off Cloudflare Fonts.
	Fonts param.Field[bool] `json:"fonts"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection param.Field[bool] `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage param.Field[bool] `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption param.Field[bool] `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish param.Field[SetConfigRuleActionParametersPolish] `json:"polish"`
	// Turn on or off Rocket Loader.
	RocketLoader param.Field[bool] `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel param.Field[SetConfigRuleActionParametersSecurityLevel] `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes param.Field[bool] `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL param.Field[SetConfigRuleActionParametersSSL] `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	SXG param.Field[bool] `json:"sxg"`
}

func (r SetConfigRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Select which file extensions to minify automatically.
type SetConfigRuleActionParametersAutominifyParam struct {
	// Minify CSS files.
	CSS param.Field[bool] `json:"css"`
	// Minify HTML files.
	HTML param.Field[bool] `json:"html"`
	// Minify JS files.
	JS param.Field[bool] `json:"js"`
}

func (r SetConfigRuleActionParametersAutominifyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type SetConfigRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r SetConfigRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type SetConfigRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r SetConfigRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck SkipRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit SkipRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string       `json:"ref"`
	JSON skipRuleJSON `json:"-"`
}

// skipRuleJSON contains the JSON metadata for the struct [SkipRule]
type skipRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleJSON) RawJSON() string {
	return r.raw
}

func (r SkipRule) implementsRulesetNewResponseRule() {}

func (r SkipRule) implementsRulesetUpdateResponseRule() {}

func (r SkipRule) implementsRulesetGetResponseRule() {}

func (r SkipRule) implementsPhaseUpdateResponseRule() {}

func (r SkipRule) implementsPhaseGetResponseRule() {}

func (r SkipRule) implementsPhaseVersionGetResponseRule() {}

func (r SkipRule) implementsRuleNewResponseRule() {}

func (r SkipRule) implementsRuleDeleteResponseRule() {}

func (r SkipRule) implementsRuleEditResponseRule() {}

func (r SkipRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type SkipRuleAction string

const (
	SkipRuleActionSkip SkipRuleAction = "skip"
)

func (r SkipRuleAction) IsKnown() bool {
	switch r {
	case SkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SkipRuleActionParameters struct {
	// A phase to skip the execution of. This property is only compatible with
	// products.
	Phase SkipRuleActionParametersPhase `json:"phase"`
	// A list of phases to skip the execution of. This option is incompatible with the
	// rulesets option.
	Phases []Phase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []SkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets option.
	Ruleset SkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                     `json:"rulesets"`
	JSON     skipRuleActionParametersJSON `json:"-"`
}

// skipRuleActionParametersJSON contains the JSON metadata for the struct
// [SkipRuleActionParameters]
type skipRuleActionParametersJSON struct {
	Phase       apijson.Field
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of. This property is only compatible with
// products.
type SkipRuleActionParametersPhase string

const (
	SkipRuleActionParametersPhaseCurrent SkipRuleActionParametersPhase = "current"
)

func (r SkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersPhaseCurrent:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type SkipRuleActionParametersProduct string

const (
	SkipRuleActionParametersProductBIC           SkipRuleActionParametersProduct = "bic"
	SkipRuleActionParametersProductHot           SkipRuleActionParametersProduct = "hot"
	SkipRuleActionParametersProductRateLimit     SkipRuleActionParametersProduct = "rateLimit"
	SkipRuleActionParametersProductSecurityLevel SkipRuleActionParametersProduct = "securityLevel"
	SkipRuleActionParametersProductUABlock       SkipRuleActionParametersProduct = "uaBlock"
	SkipRuleActionParametersProductWAF           SkipRuleActionParametersProduct = "waf"
	SkipRuleActionParametersProductZoneLockdown  SkipRuleActionParametersProduct = "zoneLockdown"
)

func (r SkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersProductBIC, SkipRuleActionParametersProductHot, SkipRuleActionParametersProductRateLimit, SkipRuleActionParametersProductSecurityLevel, SkipRuleActionParametersProductUABlock, SkipRuleActionParametersProductWAF, SkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets option.
type SkipRuleActionParametersRuleset string

const (
	SkipRuleActionParametersRulesetCurrent SkipRuleActionParametersRuleset = "current"
)

func (r SkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type SkipRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                             `json:"username_expression,required"`
	JSON               skipRuleExposedCredentialCheckJSON `json:"-"`
}

// skipRuleExposedCredentialCheckJSON contains the JSON metadata for the struct
// [SkipRuleExposedCredentialCheck]
type skipRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SkipRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type SkipRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                `json:"score_response_header_name"`
	JSON                    skipRuleRatelimitJSON `json:"-"`
}

// skipRuleRatelimitJSON contains the JSON metadata for the struct
// [SkipRuleRatelimit]
type skipRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SkipRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type SkipRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SkipRuleActionParametersParam] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[SkipRuleExposedCredentialCheckParam] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[SkipRuleRatelimitParam] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r SkipRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SkipRuleParam) implementsRulesetNewParamsRuleUnion() {}

func (r SkipRuleParam) implementsRulesetUpdateParamsRuleUnion() {}

func (r SkipRuleParam) implementsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SkipRuleActionParametersParam struct {
	// A phase to skip the execution of. This property is only compatible with
	// products.
	Phase param.Field[SkipRuleActionParametersPhase] `json:"phase"`
	// A list of phases to skip the execution of. This option is incompatible with the
	// rulesets option.
	Phases param.Field[[]Phase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]SkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets option.
	Ruleset param.Field[SkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r SkipRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type SkipRuleExposedCredentialCheckParam struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r SkipRuleExposedCredentialCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type SkipRuleRatelimitParam struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r SkipRuleRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ruleset object.
type RuleNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string              `json:"description"`
	JSON        ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [RedirectRuleActionParameters],
	// [RewriteRuleActionParameters], [RouteRuleActionParameters],
	// [ScoreRuleActionParameters], [ServeErrorRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters], [LogCustomFieldRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RuleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RuleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck], [SetConfigRuleExposedCredentialCheck],
	// [SkipRuleExposedCredentialCheck], [SetCacheSettingsRuleExposedCredentialCheck],
	// [LogCustomFieldRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RuleNewResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RuleNewResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                  `json:"ref"`
	JSON  ruleNewResponseRuleJSON `json:"-"`
	union RuleNewResponseRulesUnion
}

// ruleNewResponseRuleJSON contains the JSON metadata for the struct
// [RuleNewResponseRule]
type ruleNewResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r ruleNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RuleNewResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RuleNewResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RuleNewResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RuleNewResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RuleNewResponseRule) AsUnion() RuleNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RuleNewResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [RuleNewResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RuleNewResponseRulesUnion interface {
	implementsRuleNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
	)
}

type RuleNewResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleNewResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                        `json:"ref"`
	JSON ruleNewResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for the
// struct [RuleNewResponseRulesRulesetsChallengeRule]
type ruleNewResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsChallengeRule) implementsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsChallengeRuleAction string

const (
	RuleNewResponseRulesRulesetsChallengeRuleActionChallenge RuleNewResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RuleNewResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                              `json:"username_expression,required"`
	JSON               ruleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains the
// JSON metadata for the struct
// [RuleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type ruleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleNewResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                 `json:"score_response_header_name"`
	JSON                    ruleNewResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RuleNewResponseRulesRulesetsChallengeRuleRatelimit]
type ruleNewResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleNewResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                          `json:"ref"`
	JSON ruleNewResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata for
// the struct [RuleNewResponseRulesRulesetsJSChallengeRule]
type ruleNewResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsJSChallengeRule) implementsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsJSChallengeRuleAction string

const (
	RuleNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge RuleNewResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RuleNewResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                `json:"username_expression,required"`
	JSON               ruleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type ruleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleNewResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                   `json:"score_response_header_name"`
	JSON                    ruleNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RuleNewResponseRulesRulesetsJSChallengeRuleRatelimit]
type ruleNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleNewResponseRulesAction string

const (
	RuleNewResponseRulesActionBlock                RuleNewResponseRulesAction = "block"
	RuleNewResponseRulesActionChallenge            RuleNewResponseRulesAction = "challenge"
	RuleNewResponseRulesActionCompressResponse     RuleNewResponseRulesAction = "compress_response"
	RuleNewResponseRulesActionExecute              RuleNewResponseRulesAction = "execute"
	RuleNewResponseRulesActionJSChallenge          RuleNewResponseRulesAction = "js_challenge"
	RuleNewResponseRulesActionLog                  RuleNewResponseRulesAction = "log"
	RuleNewResponseRulesActionManagedChallenge     RuleNewResponseRulesAction = "managed_challenge"
	RuleNewResponseRulesActionRedirect             RuleNewResponseRulesAction = "redirect"
	RuleNewResponseRulesActionRewrite              RuleNewResponseRulesAction = "rewrite"
	RuleNewResponseRulesActionRoute                RuleNewResponseRulesAction = "route"
	RuleNewResponseRulesActionScore                RuleNewResponseRulesAction = "score"
	RuleNewResponseRulesActionServeError           RuleNewResponseRulesAction = "serve_error"
	RuleNewResponseRulesActionSetConfig            RuleNewResponseRulesAction = "set_config"
	RuleNewResponseRulesActionSkip                 RuleNewResponseRulesAction = "skip"
	RuleNewResponseRulesActionSetCacheSettings     RuleNewResponseRulesAction = "set_cache_settings"
	RuleNewResponseRulesActionLogCustomField       RuleNewResponseRulesAction = "log_custom_field"
	RuleNewResponseRulesActionDDoSDynamic          RuleNewResponseRulesAction = "ddos_dynamic"
	RuleNewResponseRulesActionForceConnectionClose RuleNewResponseRulesAction = "force_connection_close"
)

func (r RuleNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesActionBlock, RuleNewResponseRulesActionChallenge, RuleNewResponseRulesActionCompressResponse, RuleNewResponseRulesActionExecute, RuleNewResponseRulesActionJSChallenge, RuleNewResponseRulesActionLog, RuleNewResponseRulesActionManagedChallenge, RuleNewResponseRulesActionRedirect, RuleNewResponseRulesActionRewrite, RuleNewResponseRulesActionRoute, RuleNewResponseRulesActionScore, RuleNewResponseRulesActionServeError, RuleNewResponseRulesActionSetConfig, RuleNewResponseRulesActionSkip, RuleNewResponseRulesActionSetCacheSettings, RuleNewResponseRulesActionLogCustomField, RuleNewResponseRulesActionDDoSDynamic, RuleNewResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A ruleset object.
type RuleDeleteResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleDeleteResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [RedirectRuleActionParameters],
	// [RewriteRuleActionParameters], [RouteRuleActionParameters],
	// [ScoreRuleActionParameters], [ServeErrorRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters], [LogCustomFieldRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RuleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RuleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck], [SetConfigRuleExposedCredentialCheck],
	// [SkipRuleExposedCredentialCheck], [SetCacheSettingsRuleExposedCredentialCheck],
	// [LogCustomFieldRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RuleDeleteResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RuleDeleteResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                     `json:"ref"`
	JSON  ruleDeleteResponseRuleJSON `json:"-"`
	union RuleDeleteResponseRulesUnion
}

// ruleDeleteResponseRuleJSON contains the JSON metadata for the struct
// [RuleDeleteResponseRule]
type ruleDeleteResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r ruleDeleteResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleDeleteResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RuleDeleteResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RuleDeleteResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RuleDeleteResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RuleDeleteResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RuleDeleteResponseRule) AsUnion() RuleDeleteResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RuleDeleteResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [RuleDeleteResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RuleDeleteResponseRulesUnion interface {
	implementsRuleDeleteResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
	)
}

type RuleDeleteResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleDeleteResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RuleDeleteResponseRulesRulesetsChallengeRule]
type ruleDeleteResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsChallengeRule) implementsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsChallengeRuleAction string

const (
	RuleDeleteResponseRulesRulesetsChallengeRuleActionChallenge RuleDeleteResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RuleDeleteResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression,required"`
	JSON               ruleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type ruleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleDeleteResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                    `json:"score_response_header_name"`
	JSON                    ruleDeleteResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RuleDeleteResponseRulesRulesetsChallengeRuleRatelimit]
type ruleDeleteResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleDeleteResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RuleDeleteResponseRulesRulesetsJSChallengeRule]
type ruleDeleteResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsJSChallengeRule) implementsRuleDeleteResponseRule() {}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsJSChallengeRuleAction string

const (
	RuleDeleteResponseRulesRulesetsJSChallengeRuleActionJSChallenge RuleDeleteResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RuleDeleteResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression,required"`
	JSON               ruleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type ruleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleDeleteResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                      `json:"score_response_header_name"`
	JSON                    ruleDeleteResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RuleDeleteResponseRulesRulesetsJSChallengeRuleRatelimit]
type ruleDeleteResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesAction string

const (
	RuleDeleteResponseRulesActionBlock                RuleDeleteResponseRulesAction = "block"
	RuleDeleteResponseRulesActionChallenge            RuleDeleteResponseRulesAction = "challenge"
	RuleDeleteResponseRulesActionCompressResponse     RuleDeleteResponseRulesAction = "compress_response"
	RuleDeleteResponseRulesActionExecute              RuleDeleteResponseRulesAction = "execute"
	RuleDeleteResponseRulesActionJSChallenge          RuleDeleteResponseRulesAction = "js_challenge"
	RuleDeleteResponseRulesActionLog                  RuleDeleteResponseRulesAction = "log"
	RuleDeleteResponseRulesActionManagedChallenge     RuleDeleteResponseRulesAction = "managed_challenge"
	RuleDeleteResponseRulesActionRedirect             RuleDeleteResponseRulesAction = "redirect"
	RuleDeleteResponseRulesActionRewrite              RuleDeleteResponseRulesAction = "rewrite"
	RuleDeleteResponseRulesActionRoute                RuleDeleteResponseRulesAction = "route"
	RuleDeleteResponseRulesActionScore                RuleDeleteResponseRulesAction = "score"
	RuleDeleteResponseRulesActionServeError           RuleDeleteResponseRulesAction = "serve_error"
	RuleDeleteResponseRulesActionSetConfig            RuleDeleteResponseRulesAction = "set_config"
	RuleDeleteResponseRulesActionSkip                 RuleDeleteResponseRulesAction = "skip"
	RuleDeleteResponseRulesActionSetCacheSettings     RuleDeleteResponseRulesAction = "set_cache_settings"
	RuleDeleteResponseRulesActionLogCustomField       RuleDeleteResponseRulesAction = "log_custom_field"
	RuleDeleteResponseRulesActionDDoSDynamic          RuleDeleteResponseRulesAction = "ddos_dynamic"
	RuleDeleteResponseRulesActionForceConnectionClose RuleDeleteResponseRulesAction = "force_connection_close"
)

func (r RuleDeleteResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesActionBlock, RuleDeleteResponseRulesActionChallenge, RuleDeleteResponseRulesActionCompressResponse, RuleDeleteResponseRulesActionExecute, RuleDeleteResponseRulesActionJSChallenge, RuleDeleteResponseRulesActionLog, RuleDeleteResponseRulesActionManagedChallenge, RuleDeleteResponseRulesActionRedirect, RuleDeleteResponseRulesActionRewrite, RuleDeleteResponseRulesActionRoute, RuleDeleteResponseRulesActionScore, RuleDeleteResponseRulesActionServeError, RuleDeleteResponseRulesActionSetConfig, RuleDeleteResponseRulesActionSkip, RuleDeleteResponseRulesActionSetCacheSettings, RuleDeleteResponseRulesActionLogCustomField, RuleDeleteResponseRulesActionDDoSDynamic, RuleDeleteResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A ruleset object.
type RuleEditResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleEditResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [RedirectRuleActionParameters],
	// [RewriteRuleActionParameters], [RouteRuleActionParameters],
	// [ScoreRuleActionParameters], [ServeErrorRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters], [LogCustomFieldRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RuleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RuleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck], [SetConfigRuleExposedCredentialCheck],
	// [SkipRuleExposedCredentialCheck], [SetCacheSettingsRuleExposedCredentialCheck],
	// [LogCustomFieldRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RuleEditResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RuleEditResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                   `json:"ref"`
	JSON  ruleEditResponseRuleJSON `json:"-"`
	union RuleEditResponseRulesUnion
}

// ruleEditResponseRuleJSON contains the JSON metadata for the struct
// [RuleEditResponseRule]
type ruleEditResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r ruleEditResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleEditResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RuleEditResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RuleEditResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RuleEditResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RuleEditResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RuleEditResponseRule) AsUnion() RuleEditResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RuleEditResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [RuleEditResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RuleEditResponseRulesUnion interface {
	implementsRuleEditResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
	)
}

type RuleEditResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleEditResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON ruleEditResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RuleEditResponseRulesRulesetsChallengeRule]
type ruleEditResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsChallengeRule) implementsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsChallengeRuleAction string

const (
	RuleEditResponseRulesRulesetsChallengeRuleActionChallenge RuleEditResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RuleEditResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                               `json:"username_expression,required"`
	JSON               ruleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type ruleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleEditResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                  `json:"score_response_header_name"`
	JSON                    ruleEditResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RuleEditResponseRulesRulesetsChallengeRuleRatelimit]
type ruleEditResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RuleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RuleEditResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON ruleEditResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata for
// the struct [RuleEditResponseRulesRulesetsJSChallengeRule]
type ruleEditResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsJSChallengeRule) implementsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsJSChallengeRuleAction string

const (
	RuleEditResponseRulesRulesetsJSChallengeRuleActionJSChallenge RuleEditResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RuleEditResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression,required"`
	JSON               ruleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type ruleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RuleEditResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics []string `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName string                                                    `json:"score_response_header_name"`
	JSON                    ruleEditResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RuleEditResponseRulesRulesetsJSChallengeRuleRatelimit]
type ruleEditResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleEditResponseRulesAction string

const (
	RuleEditResponseRulesActionBlock                RuleEditResponseRulesAction = "block"
	RuleEditResponseRulesActionChallenge            RuleEditResponseRulesAction = "challenge"
	RuleEditResponseRulesActionCompressResponse     RuleEditResponseRulesAction = "compress_response"
	RuleEditResponseRulesActionExecute              RuleEditResponseRulesAction = "execute"
	RuleEditResponseRulesActionJSChallenge          RuleEditResponseRulesAction = "js_challenge"
	RuleEditResponseRulesActionLog                  RuleEditResponseRulesAction = "log"
	RuleEditResponseRulesActionManagedChallenge     RuleEditResponseRulesAction = "managed_challenge"
	RuleEditResponseRulesActionRedirect             RuleEditResponseRulesAction = "redirect"
	RuleEditResponseRulesActionRewrite              RuleEditResponseRulesAction = "rewrite"
	RuleEditResponseRulesActionRoute                RuleEditResponseRulesAction = "route"
	RuleEditResponseRulesActionScore                RuleEditResponseRulesAction = "score"
	RuleEditResponseRulesActionServeError           RuleEditResponseRulesAction = "serve_error"
	RuleEditResponseRulesActionSetConfig            RuleEditResponseRulesAction = "set_config"
	RuleEditResponseRulesActionSkip                 RuleEditResponseRulesAction = "skip"
	RuleEditResponseRulesActionSetCacheSettings     RuleEditResponseRulesAction = "set_cache_settings"
	RuleEditResponseRulesActionLogCustomField       RuleEditResponseRulesAction = "log_custom_field"
	RuleEditResponseRulesActionDDoSDynamic          RuleEditResponseRulesAction = "ddos_dynamic"
	RuleEditResponseRulesActionForceConnectionClose RuleEditResponseRulesAction = "force_connection_close"
)

func (r RuleEditResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesActionBlock, RuleEditResponseRulesActionChallenge, RuleEditResponseRulesActionCompressResponse, RuleEditResponseRulesActionExecute, RuleEditResponseRulesActionJSChallenge, RuleEditResponseRulesActionLog, RuleEditResponseRulesActionManagedChallenge, RuleEditResponseRulesActionRedirect, RuleEditResponseRulesActionRewrite, RuleEditResponseRulesActionRoute, RuleEditResponseRulesActionScore, RuleEditResponseRulesActionServeError, RuleEditResponseRulesActionSetConfig, RuleEditResponseRulesActionSkip, RuleEditResponseRulesActionSetCacheSettings, RuleEditResponseRulesActionLogCustomField, RuleEditResponseRulesActionDDoSDynamic, RuleEditResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

type RuleNewParams struct {
	Body RuleNewParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleNewParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RuleNewParamsBodyAction] `json:"action"`
	ActionParameters param.Field[interface{}]             `json:"action_parameters"`
	Categories       param.Field[interface{}]             `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Position  param.Field[interface{}]  `json:"position"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBody) implementsRuleNewParamsBodyUnion() {}

// Satisfied by [rulesets.RuleNewParamsBodyBlockRule],
// [rulesets.RuleNewParamsBodyChallengeRule],
// [rulesets.RuleNewParamsBodyCompressionRule],
// [rulesets.RuleNewParamsBodyExecuteRule],
// [rulesets.RuleNewParamsBodyJavascriptChallengeRule],
// [rulesets.RuleNewParamsBodyLogRule],
// [rulesets.RuleNewParamsBodyManagedChallengeRule],
// [rulesets.RuleNewParamsBodyRedirectRule],
// [rulesets.RuleNewParamsBodyRewriteRule], [rulesets.RuleNewParamsBodyOriginRule],
// [rulesets.RuleNewParamsBodyScoreRule],
// [rulesets.RuleNewParamsBodyServeErrorRule],
// [rulesets.RuleNewParamsBodySetConfigRule], [rulesets.RuleNewParamsBodySkipRule],
// [rulesets.RuleNewParamsBodySetCacheSettingsRule],
// [rulesets.RuleNewParamsBodyLogCustomFieldRule],
// [rulesets.RuleNewParamsBodyDDoSDynamicRule],
// [rulesets.RuleNewParamsBodyForceConnectionCloseRule], [RuleNewParamsBody].
type RuleNewParamsBodyUnion interface {
	implementsRuleNewParamsBodyUnion()
}

type RuleNewParamsBodyBlockRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyBlockRulePositionUnion] `json:"position"`
	BlockRuleParam
}

func (r RuleNewParamsBodyBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyBlockRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyBlockRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyBlockRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyBlockRulePosition) implementsRuleNewParamsBodyBlockRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyBlockRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyBlockRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyBlockRulePositionIndexPosition],
// [RuleNewParamsBodyBlockRulePosition].
type RuleNewParamsBodyBlockRulePositionUnion interface {
	implementsRuleNewParamsBodyBlockRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyBlockRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyBlockRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyBlockRulePositionBeforePosition) implementsRuleNewParamsBodyBlockRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyBlockRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyBlockRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyBlockRulePositionAfterPosition) implementsRuleNewParamsBodyBlockRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyBlockRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyBlockRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyBlockRulePositionIndexPosition) implementsRuleNewParamsBodyBlockRulePositionUnion() {
}

type RuleNewParamsBodyChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsBodyChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RuleNewParamsBodyChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyChallengeRulePositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RuleNewParamsBodyChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsBodyChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyChallengeRule) implementsRuleNewParamsBodyUnion() {}

// The action to perform when the rule matches.
type RuleNewParamsBodyChallengeRuleAction string

const (
	RuleNewParamsBodyChallengeRuleActionChallenge RuleNewParamsBodyChallengeRuleAction = "challenge"
)

func (r RuleNewParamsBodyChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsBodyChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleNewParamsBodyChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RuleNewParamsBodyChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyChallengeRulePosition) implementsRuleNewParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyChallengeRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyChallengeRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyChallengeRulePositionIndexPosition],
// [RuleNewParamsBodyChallengeRulePosition].
type RuleNewParamsBodyChallengeRulePositionUnion interface {
	implementsRuleNewParamsBodyChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyChallengeRulePositionBeforePosition) implementsRuleNewParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyChallengeRulePositionAfterPosition) implementsRuleNewParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyChallengeRulePositionIndexPosition) implementsRuleNewParamsBodyChallengeRulePositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type RuleNewParamsBodyChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RuleNewParamsBodyChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsBodyCompressionRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyCompressionRulePositionUnion] `json:"position"`
	CompressResponseRuleParam
}

func (r RuleNewParamsBodyCompressionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyCompressionRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyCompressionRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyCompressionRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyCompressionRulePosition) implementsRuleNewParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyCompressionRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyCompressionRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyCompressionRulePositionIndexPosition],
// [RuleNewParamsBodyCompressionRulePosition].
type RuleNewParamsBodyCompressionRulePositionUnion interface {
	implementsRuleNewParamsBodyCompressionRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyCompressionRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyCompressionRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyCompressionRulePositionBeforePosition) implementsRuleNewParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyCompressionRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyCompressionRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyCompressionRulePositionAfterPosition) implementsRuleNewParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyCompressionRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyCompressionRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyCompressionRulePositionIndexPosition) implementsRuleNewParamsBodyCompressionRulePositionUnion() {
}

type RuleNewParamsBodyExecuteRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyExecuteRulePositionUnion] `json:"position"`
	ExecuteRuleParam
}

func (r RuleNewParamsBodyExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyExecuteRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyExecuteRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyExecuteRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyExecuteRulePosition) implementsRuleNewParamsBodyExecuteRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyExecuteRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyExecuteRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyExecuteRulePositionIndexPosition],
// [RuleNewParamsBodyExecuteRulePosition].
type RuleNewParamsBodyExecuteRulePositionUnion interface {
	implementsRuleNewParamsBodyExecuteRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyExecuteRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyExecuteRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyExecuteRulePositionBeforePosition) implementsRuleNewParamsBodyExecuteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyExecuteRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyExecuteRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyExecuteRulePositionAfterPosition) implementsRuleNewParamsBodyExecuteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyExecuteRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyExecuteRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyExecuteRulePositionIndexPosition) implementsRuleNewParamsBodyExecuteRulePositionUnion() {
}

type RuleNewParamsBodyJavascriptChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsBodyJavascriptChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RuleNewParamsBodyJavascriptChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyJavascriptChallengeRulePositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RuleNewParamsBodyJavascriptChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsBodyJavascriptChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyJavascriptChallengeRule) implementsRuleNewParamsBodyUnion() {}

// The action to perform when the rule matches.
type RuleNewParamsBodyJavascriptChallengeRuleAction string

const (
	RuleNewParamsBodyJavascriptChallengeRuleActionJSChallenge RuleNewParamsBodyJavascriptChallengeRuleAction = "js_challenge"
)

func (r RuleNewParamsBodyJavascriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsBodyJavascriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleNewParamsBodyJavascriptChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RuleNewParamsBodyJavascriptChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyJavascriptChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyJavascriptChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyJavascriptChallengeRulePosition) implementsRuleNewParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleNewParamsBodyJavascriptChallengeRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyJavascriptChallengeRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyJavascriptChallengeRulePositionIndexPosition],
// [RuleNewParamsBodyJavascriptChallengeRulePosition].
type RuleNewParamsBodyJavascriptChallengeRulePositionUnion interface {
	implementsRuleNewParamsBodyJavascriptChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyJavascriptChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionBeforePosition) implementsRuleNewParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyJavascriptChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionAfterPosition) implementsRuleNewParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyJavascriptChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyJavascriptChallengeRulePositionIndexPosition) implementsRuleNewParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type RuleNewParamsBodyJavascriptChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RuleNewParamsBodyJavascriptChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsBodyLogRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyLogRulePositionUnion] `json:"position"`
	LogRuleParam
}

func (r RuleNewParamsBodyLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyLogRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogRulePosition) implementsRuleNewParamsBodyLogRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyLogRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyLogRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyLogRulePositionIndexPosition],
// [RuleNewParamsBodyLogRulePosition].
type RuleNewParamsBodyLogRulePositionUnion interface {
	implementsRuleNewParamsBodyLogRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyLogRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogRulePositionBeforePosition) implementsRuleNewParamsBodyLogRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyLogRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogRulePositionAfterPosition) implementsRuleNewParamsBodyLogRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyLogRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogRulePositionIndexPosition) implementsRuleNewParamsBodyLogRulePositionUnion() {
}

type RuleNewParamsBodyManagedChallengeRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyManagedChallengeRulePositionUnion] `json:"position"`
	ManagedChallengeRuleParam
}

func (r RuleNewParamsBodyManagedChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyManagedChallengeRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyManagedChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyManagedChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyManagedChallengeRulePosition) implementsRuleNewParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleNewParamsBodyManagedChallengeRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyManagedChallengeRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyManagedChallengeRulePositionIndexPosition],
// [RuleNewParamsBodyManagedChallengeRulePosition].
type RuleNewParamsBodyManagedChallengeRulePositionUnion interface {
	implementsRuleNewParamsBodyManagedChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyManagedChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyManagedChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyManagedChallengeRulePositionBeforePosition) implementsRuleNewParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyManagedChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyManagedChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyManagedChallengeRulePositionAfterPosition) implementsRuleNewParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyManagedChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyManagedChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyManagedChallengeRulePositionIndexPosition) implementsRuleNewParamsBodyManagedChallengeRulePositionUnion() {
}

type RuleNewParamsBodyRedirectRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyRedirectRulePositionUnion] `json:"position"`
	RedirectRuleParam
}

func (r RuleNewParamsBodyRedirectRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRedirectRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRedirectRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyRedirectRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRedirectRulePosition) implementsRuleNewParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyRedirectRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyRedirectRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyRedirectRulePositionIndexPosition],
// [RuleNewParamsBodyRedirectRulePosition].
type RuleNewParamsBodyRedirectRulePositionUnion interface {
	implementsRuleNewParamsBodyRedirectRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRedirectRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyRedirectRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRedirectRulePositionBeforePosition) implementsRuleNewParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRedirectRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyRedirectRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRedirectRulePositionAfterPosition) implementsRuleNewParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRedirectRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyRedirectRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRedirectRulePositionIndexPosition) implementsRuleNewParamsBodyRedirectRulePositionUnion() {
}

type RuleNewParamsBodyRewriteRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyRewriteRulePositionUnion] `json:"position"`
	RewriteRuleParam
}

func (r RuleNewParamsBodyRewriteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRewriteRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRewriteRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyRewriteRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRewriteRulePosition) implementsRuleNewParamsBodyRewriteRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyRewriteRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyRewriteRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyRewriteRulePositionIndexPosition],
// [RuleNewParamsBodyRewriteRulePosition].
type RuleNewParamsBodyRewriteRulePositionUnion interface {
	implementsRuleNewParamsBodyRewriteRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRewriteRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyRewriteRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRewriteRulePositionBeforePosition) implementsRuleNewParamsBodyRewriteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRewriteRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyRewriteRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRewriteRulePositionAfterPosition) implementsRuleNewParamsBodyRewriteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyRewriteRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyRewriteRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyRewriteRulePositionIndexPosition) implementsRuleNewParamsBodyRewriteRulePositionUnion() {
}

type RuleNewParamsBodyOriginRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyOriginRulePositionUnion] `json:"position"`
	RouteRuleParam
}

func (r RuleNewParamsBodyOriginRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyOriginRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyOriginRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyOriginRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyOriginRulePosition) implementsRuleNewParamsBodyOriginRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyOriginRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyOriginRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyOriginRulePositionIndexPosition],
// [RuleNewParamsBodyOriginRulePosition].
type RuleNewParamsBodyOriginRulePositionUnion interface {
	implementsRuleNewParamsBodyOriginRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyOriginRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyOriginRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyOriginRulePositionBeforePosition) implementsRuleNewParamsBodyOriginRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyOriginRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyOriginRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyOriginRulePositionAfterPosition) implementsRuleNewParamsBodyOriginRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyOriginRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyOriginRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyOriginRulePositionIndexPosition) implementsRuleNewParamsBodyOriginRulePositionUnion() {
}

type RuleNewParamsBodyScoreRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyScoreRulePositionUnion] `json:"position"`
	ScoreRuleParam
}

func (r RuleNewParamsBodyScoreRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyScoreRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyScoreRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyScoreRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyScoreRulePosition) implementsRuleNewParamsBodyScoreRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyScoreRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyScoreRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyScoreRulePositionIndexPosition],
// [RuleNewParamsBodyScoreRulePosition].
type RuleNewParamsBodyScoreRulePositionUnion interface {
	implementsRuleNewParamsBodyScoreRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyScoreRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyScoreRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyScoreRulePositionBeforePosition) implementsRuleNewParamsBodyScoreRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyScoreRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyScoreRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyScoreRulePositionAfterPosition) implementsRuleNewParamsBodyScoreRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyScoreRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyScoreRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyScoreRulePositionIndexPosition) implementsRuleNewParamsBodyScoreRulePositionUnion() {
}

type RuleNewParamsBodyServeErrorRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyServeErrorRulePositionUnion] `json:"position"`
	ServeErrorRuleParam
}

func (r RuleNewParamsBodyServeErrorRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyServeErrorRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyServeErrorRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyServeErrorRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyServeErrorRulePosition) implementsRuleNewParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyServeErrorRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyServeErrorRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyServeErrorRulePositionIndexPosition],
// [RuleNewParamsBodyServeErrorRulePosition].
type RuleNewParamsBodyServeErrorRulePositionUnion interface {
	implementsRuleNewParamsBodyServeErrorRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyServeErrorRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyServeErrorRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyServeErrorRulePositionBeforePosition) implementsRuleNewParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyServeErrorRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyServeErrorRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyServeErrorRulePositionAfterPosition) implementsRuleNewParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyServeErrorRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyServeErrorRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyServeErrorRulePositionIndexPosition) implementsRuleNewParamsBodyServeErrorRulePositionUnion() {
}

type RuleNewParamsBodySetConfigRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodySetConfigRulePositionUnion] `json:"position"`
	SetConfigRuleParam
}

func (r RuleNewParamsBodySetConfigRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetConfigRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetConfigRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySetConfigRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetConfigRulePosition) implementsRuleNewParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodySetConfigRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodySetConfigRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodySetConfigRulePositionIndexPosition],
// [RuleNewParamsBodySetConfigRulePosition].
type RuleNewParamsBodySetConfigRulePositionUnion interface {
	implementsRuleNewParamsBodySetConfigRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetConfigRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodySetConfigRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetConfigRulePositionBeforePosition) implementsRuleNewParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetConfigRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodySetConfigRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetConfigRulePositionAfterPosition) implementsRuleNewParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetConfigRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySetConfigRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetConfigRulePositionIndexPosition) implementsRuleNewParamsBodySetConfigRulePositionUnion() {
}

type RuleNewParamsBodySkipRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodySkipRulePositionUnion] `json:"position"`
	SkipRuleParam
}

func (r RuleNewParamsBodySkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySkipRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySkipRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySkipRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySkipRulePosition) implementsRuleNewParamsBodySkipRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodySkipRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodySkipRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodySkipRulePositionIndexPosition],
// [RuleNewParamsBodySkipRulePosition].
type RuleNewParamsBodySkipRulePositionUnion interface {
	implementsRuleNewParamsBodySkipRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySkipRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodySkipRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySkipRulePositionBeforePosition) implementsRuleNewParamsBodySkipRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySkipRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodySkipRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySkipRulePositionAfterPosition) implementsRuleNewParamsBodySkipRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySkipRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySkipRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySkipRulePositionIndexPosition) implementsRuleNewParamsBodySkipRulePositionUnion() {
}

type RuleNewParamsBodySetCacheSettingsRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodySetCacheSettingsRulePositionUnion] `json:"position"`
	SetCacheSettingsRuleParam
}

func (r RuleNewParamsBodySetCacheSettingsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetCacheSettingsRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetCacheSettingsRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySetCacheSettingsRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetCacheSettingsRulePosition) implementsRuleNewParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleNewParamsBodySetCacheSettingsRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodySetCacheSettingsRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodySetCacheSettingsRulePositionIndexPosition],
// [RuleNewParamsBodySetCacheSettingsRulePosition].
type RuleNewParamsBodySetCacheSettingsRulePositionUnion interface {
	implementsRuleNewParamsBodySetCacheSettingsRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetCacheSettingsRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionBeforePosition) implementsRuleNewParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetCacheSettingsRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionAfterPosition) implementsRuleNewParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodySetCacheSettingsRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodySetCacheSettingsRulePositionIndexPosition) implementsRuleNewParamsBodySetCacheSettingsRulePositionUnion() {
}

type RuleNewParamsBodyLogCustomFieldRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyLogCustomFieldRulePositionUnion] `json:"position"`
	LogCustomFieldRuleParam
}

func (r RuleNewParamsBodyLogCustomFieldRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogCustomFieldRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogCustomFieldRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyLogCustomFieldRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogCustomFieldRulePosition) implementsRuleNewParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleNewParamsBodyLogCustomFieldRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyLogCustomFieldRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyLogCustomFieldRulePositionIndexPosition],
// [RuleNewParamsBodyLogCustomFieldRulePosition].
type RuleNewParamsBodyLogCustomFieldRulePositionUnion interface {
	implementsRuleNewParamsBodyLogCustomFieldRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogCustomFieldRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionBeforePosition) implementsRuleNewParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogCustomFieldRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionAfterPosition) implementsRuleNewParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyLogCustomFieldRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyLogCustomFieldRulePositionIndexPosition) implementsRuleNewParamsBodyLogCustomFieldRulePositionUnion() {
}

type RuleNewParamsBodyDDoSDynamicRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyDDoSDynamicRulePositionUnion] `json:"position"`
	DDoSDynamicRuleParam
}

func (r RuleNewParamsBodyDDoSDynamicRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyDDoSDynamicRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyDDoSDynamicRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyDDoSDynamicRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyDDoSDynamicRulePosition) implementsRuleNewParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsBodyDDoSDynamicRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyDDoSDynamicRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyDDoSDynamicRulePositionIndexPosition],
// [RuleNewParamsBodyDDoSDynamicRulePosition].
type RuleNewParamsBodyDDoSDynamicRulePositionUnion interface {
	implementsRuleNewParamsBodyDDoSDynamicRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyDDoSDynamicRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionBeforePosition) implementsRuleNewParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyDDoSDynamicRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionAfterPosition) implementsRuleNewParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyDDoSDynamicRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyDDoSDynamicRulePositionIndexPosition) implementsRuleNewParamsBodyDDoSDynamicRulePositionUnion() {
}

type RuleNewParamsBodyForceConnectionCloseRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsBodyForceConnectionCloseRulePositionUnion] `json:"position"`
	ForceConnectionCloseRuleParam
}

func (r RuleNewParamsBodyForceConnectionCloseRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyForceConnectionCloseRule) implementsRuleNewParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyForceConnectionCloseRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyForceConnectionCloseRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyForceConnectionCloseRulePosition) implementsRuleNewParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleNewParamsBodyForceConnectionCloseRulePositionBeforePosition],
// [rulesets.RuleNewParamsBodyForceConnectionCloseRulePositionAfterPosition],
// [rulesets.RuleNewParamsBodyForceConnectionCloseRulePositionIndexPosition],
// [RuleNewParamsBodyForceConnectionCloseRulePosition].
type RuleNewParamsBodyForceConnectionCloseRulePositionUnion interface {
	implementsRuleNewParamsBodyForceConnectionCloseRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyForceConnectionCloseRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionBeforePosition) implementsRuleNewParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyForceConnectionCloseRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionAfterPosition) implementsRuleNewParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleNewParamsBodyForceConnectionCloseRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBodyForceConnectionCloseRulePositionIndexPosition) implementsRuleNewParamsBodyForceConnectionCloseRulePositionUnion() {
}

// The action to perform when the rule matches.
type RuleNewParamsBodyAction string

const (
	RuleNewParamsBodyActionBlock                RuleNewParamsBodyAction = "block"
	RuleNewParamsBodyActionChallenge            RuleNewParamsBodyAction = "challenge"
	RuleNewParamsBodyActionCompressResponse     RuleNewParamsBodyAction = "compress_response"
	RuleNewParamsBodyActionExecute              RuleNewParamsBodyAction = "execute"
	RuleNewParamsBodyActionJSChallenge          RuleNewParamsBodyAction = "js_challenge"
	RuleNewParamsBodyActionLog                  RuleNewParamsBodyAction = "log"
	RuleNewParamsBodyActionManagedChallenge     RuleNewParamsBodyAction = "managed_challenge"
	RuleNewParamsBodyActionRedirect             RuleNewParamsBodyAction = "redirect"
	RuleNewParamsBodyActionRewrite              RuleNewParamsBodyAction = "rewrite"
	RuleNewParamsBodyActionRoute                RuleNewParamsBodyAction = "route"
	RuleNewParamsBodyActionScore                RuleNewParamsBodyAction = "score"
	RuleNewParamsBodyActionServeError           RuleNewParamsBodyAction = "serve_error"
	RuleNewParamsBodyActionSetConfig            RuleNewParamsBodyAction = "set_config"
	RuleNewParamsBodyActionSkip                 RuleNewParamsBodyAction = "skip"
	RuleNewParamsBodyActionSetCacheSettings     RuleNewParamsBodyAction = "set_cache_settings"
	RuleNewParamsBodyActionLogCustomField       RuleNewParamsBodyAction = "log_custom_field"
	RuleNewParamsBodyActionDDoSDynamic          RuleNewParamsBodyAction = "ddos_dynamic"
	RuleNewParamsBodyActionForceConnectionClose RuleNewParamsBodyAction = "force_connection_close"
)

func (r RuleNewParamsBodyAction) IsKnown() bool {
	switch r {
	case RuleNewParamsBodyActionBlock, RuleNewParamsBodyActionChallenge, RuleNewParamsBodyActionCompressResponse, RuleNewParamsBodyActionExecute, RuleNewParamsBodyActionJSChallenge, RuleNewParamsBodyActionLog, RuleNewParamsBodyActionManagedChallenge, RuleNewParamsBodyActionRedirect, RuleNewParamsBodyActionRewrite, RuleNewParamsBodyActionRoute, RuleNewParamsBodyActionScore, RuleNewParamsBodyActionServeError, RuleNewParamsBodyActionSetConfig, RuleNewParamsBodyActionSkip, RuleNewParamsBodyActionSetCacheSettings, RuleNewParamsBodyActionLogCustomField, RuleNewParamsBodyActionDDoSDynamic, RuleNewParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                  `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeErrorsSource]
type ruleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeMessagesSource]
type ruleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeErrorsSource]
type ruleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeMessagesSource]
type ruleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditParams struct {
	Body RuleEditParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleEditParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RuleEditParamsBodyAction] `json:"action"`
	ActionParameters param.Field[interface{}]              `json:"action_parameters"`
	Categories       param.Field[interface{}]              `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Position  param.Field[interface{}]  `json:"position"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBody) implementsRuleEditParamsBodyUnion() {}

// Satisfied by [rulesets.RuleEditParamsBodyBlockRule],
// [rulesets.RuleEditParamsBodyChallengeRule],
// [rulesets.RuleEditParamsBodyCompressionRule],
// [rulesets.RuleEditParamsBodyExecuteRule],
// [rulesets.RuleEditParamsBodyJavascriptChallengeRule],
// [rulesets.RuleEditParamsBodyLogRule],
// [rulesets.RuleEditParamsBodyManagedChallengeRule],
// [rulesets.RuleEditParamsBodyRedirectRule],
// [rulesets.RuleEditParamsBodyRewriteRule],
// [rulesets.RuleEditParamsBodyOriginRule], [rulesets.RuleEditParamsBodyScoreRule],
// [rulesets.RuleEditParamsBodyServeErrorRule],
// [rulesets.RuleEditParamsBodySetConfigRule],
// [rulesets.RuleEditParamsBodySkipRule],
// [rulesets.RuleEditParamsBodySetCacheSettingsRule],
// [rulesets.RuleEditParamsBodyLogCustomFieldRule],
// [rulesets.RuleEditParamsBodyDDoSDynamicRule],
// [rulesets.RuleEditParamsBodyForceConnectionCloseRule], [RuleEditParamsBody].
type RuleEditParamsBodyUnion interface {
	implementsRuleEditParamsBodyUnion()
}

type RuleEditParamsBodyBlockRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyBlockRulePositionUnion] `json:"position"`
	BlockRuleParam
}

func (r RuleEditParamsBodyBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyBlockRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyBlockRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyBlockRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyBlockRulePosition) implementsRuleEditParamsBodyBlockRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyBlockRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyBlockRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyBlockRulePositionIndexPosition],
// [RuleEditParamsBodyBlockRulePosition].
type RuleEditParamsBodyBlockRulePositionUnion interface {
	implementsRuleEditParamsBodyBlockRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyBlockRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyBlockRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyBlockRulePositionBeforePosition) implementsRuleEditParamsBodyBlockRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyBlockRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyBlockRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyBlockRulePositionAfterPosition) implementsRuleEditParamsBodyBlockRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyBlockRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyBlockRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyBlockRulePositionIndexPosition) implementsRuleEditParamsBodyBlockRulePositionUnion() {
}

type RuleEditParamsBodyChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsBodyChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RuleEditParamsBodyChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyChallengeRulePositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RuleEditParamsBodyChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsBodyChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyChallengeRule) implementsRuleEditParamsBodyUnion() {}

// The action to perform when the rule matches.
type RuleEditParamsBodyChallengeRuleAction string

const (
	RuleEditParamsBodyChallengeRuleActionChallenge RuleEditParamsBodyChallengeRuleAction = "challenge"
)

func (r RuleEditParamsBodyChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsBodyChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleEditParamsBodyChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RuleEditParamsBodyChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyChallengeRulePosition) implementsRuleEditParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyChallengeRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyChallengeRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyChallengeRulePositionIndexPosition],
// [RuleEditParamsBodyChallengeRulePosition].
type RuleEditParamsBodyChallengeRulePositionUnion interface {
	implementsRuleEditParamsBodyChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyChallengeRulePositionBeforePosition) implementsRuleEditParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyChallengeRulePositionAfterPosition) implementsRuleEditParamsBodyChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyChallengeRulePositionIndexPosition) implementsRuleEditParamsBodyChallengeRulePositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type RuleEditParamsBodyChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RuleEditParamsBodyChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsBodyCompressionRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyCompressionRulePositionUnion] `json:"position"`
	CompressResponseRuleParam
}

func (r RuleEditParamsBodyCompressionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyCompressionRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyCompressionRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyCompressionRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyCompressionRulePosition) implementsRuleEditParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyCompressionRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyCompressionRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyCompressionRulePositionIndexPosition],
// [RuleEditParamsBodyCompressionRulePosition].
type RuleEditParamsBodyCompressionRulePositionUnion interface {
	implementsRuleEditParamsBodyCompressionRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyCompressionRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyCompressionRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyCompressionRulePositionBeforePosition) implementsRuleEditParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyCompressionRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyCompressionRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyCompressionRulePositionAfterPosition) implementsRuleEditParamsBodyCompressionRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyCompressionRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyCompressionRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyCompressionRulePositionIndexPosition) implementsRuleEditParamsBodyCompressionRulePositionUnion() {
}

type RuleEditParamsBodyExecuteRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyExecuteRulePositionUnion] `json:"position"`
	ExecuteRuleParam
}

func (r RuleEditParamsBodyExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyExecuteRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyExecuteRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyExecuteRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyExecuteRulePosition) implementsRuleEditParamsBodyExecuteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyExecuteRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyExecuteRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyExecuteRulePositionIndexPosition],
// [RuleEditParamsBodyExecuteRulePosition].
type RuleEditParamsBodyExecuteRulePositionUnion interface {
	implementsRuleEditParamsBodyExecuteRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyExecuteRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyExecuteRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyExecuteRulePositionBeforePosition) implementsRuleEditParamsBodyExecuteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyExecuteRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyExecuteRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyExecuteRulePositionAfterPosition) implementsRuleEditParamsBodyExecuteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyExecuteRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyExecuteRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyExecuteRulePositionIndexPosition) implementsRuleEditParamsBodyExecuteRulePositionUnion() {
}

type RuleEditParamsBodyJavascriptChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsBodyJavascriptChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RuleEditParamsBodyJavascriptChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyJavascriptChallengeRulePositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RuleEditParamsBodyJavascriptChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsBodyJavascriptChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyJavascriptChallengeRule) implementsRuleEditParamsBodyUnion() {}

// The action to perform when the rule matches.
type RuleEditParamsBodyJavascriptChallengeRuleAction string

const (
	RuleEditParamsBodyJavascriptChallengeRuleActionJSChallenge RuleEditParamsBodyJavascriptChallengeRuleAction = "js_challenge"
)

func (r RuleEditParamsBodyJavascriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsBodyJavascriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RuleEditParamsBodyJavascriptChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RuleEditParamsBodyJavascriptChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyJavascriptChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyJavascriptChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyJavascriptChallengeRulePosition) implementsRuleEditParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleEditParamsBodyJavascriptChallengeRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyJavascriptChallengeRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyJavascriptChallengeRulePositionIndexPosition],
// [RuleEditParamsBodyJavascriptChallengeRulePosition].
type RuleEditParamsBodyJavascriptChallengeRulePositionUnion interface {
	implementsRuleEditParamsBodyJavascriptChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyJavascriptChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionBeforePosition) implementsRuleEditParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyJavascriptChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionAfterPosition) implementsRuleEditParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyJavascriptChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyJavascriptChallengeRulePositionIndexPosition) implementsRuleEditParamsBodyJavascriptChallengeRulePositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type RuleEditParamsBodyJavascriptChallengeRuleRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RuleEditParamsBodyJavascriptChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsBodyLogRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyLogRulePositionUnion] `json:"position"`
	LogRuleParam
}

func (r RuleEditParamsBodyLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyLogRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogRulePosition) implementsRuleEditParamsBodyLogRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyLogRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyLogRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyLogRulePositionIndexPosition],
// [RuleEditParamsBodyLogRulePosition].
type RuleEditParamsBodyLogRulePositionUnion interface {
	implementsRuleEditParamsBodyLogRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyLogRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogRulePositionBeforePosition) implementsRuleEditParamsBodyLogRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyLogRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogRulePositionAfterPosition) implementsRuleEditParamsBodyLogRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyLogRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogRulePositionIndexPosition) implementsRuleEditParamsBodyLogRulePositionUnion() {
}

type RuleEditParamsBodyManagedChallengeRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyManagedChallengeRulePositionUnion] `json:"position"`
	ManagedChallengeRuleParam
}

func (r RuleEditParamsBodyManagedChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyManagedChallengeRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyManagedChallengeRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyManagedChallengeRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyManagedChallengeRulePosition) implementsRuleEditParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleEditParamsBodyManagedChallengeRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyManagedChallengeRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyManagedChallengeRulePositionIndexPosition],
// [RuleEditParamsBodyManagedChallengeRulePosition].
type RuleEditParamsBodyManagedChallengeRulePositionUnion interface {
	implementsRuleEditParamsBodyManagedChallengeRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyManagedChallengeRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyManagedChallengeRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyManagedChallengeRulePositionBeforePosition) implementsRuleEditParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyManagedChallengeRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyManagedChallengeRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyManagedChallengeRulePositionAfterPosition) implementsRuleEditParamsBodyManagedChallengeRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyManagedChallengeRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyManagedChallengeRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyManagedChallengeRulePositionIndexPosition) implementsRuleEditParamsBodyManagedChallengeRulePositionUnion() {
}

type RuleEditParamsBodyRedirectRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyRedirectRulePositionUnion] `json:"position"`
	RedirectRuleParam
}

func (r RuleEditParamsBodyRedirectRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRedirectRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRedirectRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyRedirectRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRedirectRulePosition) implementsRuleEditParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyRedirectRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyRedirectRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyRedirectRulePositionIndexPosition],
// [RuleEditParamsBodyRedirectRulePosition].
type RuleEditParamsBodyRedirectRulePositionUnion interface {
	implementsRuleEditParamsBodyRedirectRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRedirectRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyRedirectRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRedirectRulePositionBeforePosition) implementsRuleEditParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRedirectRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyRedirectRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRedirectRulePositionAfterPosition) implementsRuleEditParamsBodyRedirectRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRedirectRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyRedirectRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRedirectRulePositionIndexPosition) implementsRuleEditParamsBodyRedirectRulePositionUnion() {
}

type RuleEditParamsBodyRewriteRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyRewriteRulePositionUnion] `json:"position"`
	RewriteRuleParam
}

func (r RuleEditParamsBodyRewriteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRewriteRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRewriteRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyRewriteRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRewriteRulePosition) implementsRuleEditParamsBodyRewriteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyRewriteRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyRewriteRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyRewriteRulePositionIndexPosition],
// [RuleEditParamsBodyRewriteRulePosition].
type RuleEditParamsBodyRewriteRulePositionUnion interface {
	implementsRuleEditParamsBodyRewriteRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRewriteRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyRewriteRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRewriteRulePositionBeforePosition) implementsRuleEditParamsBodyRewriteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRewriteRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyRewriteRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRewriteRulePositionAfterPosition) implementsRuleEditParamsBodyRewriteRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyRewriteRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyRewriteRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyRewriteRulePositionIndexPosition) implementsRuleEditParamsBodyRewriteRulePositionUnion() {
}

type RuleEditParamsBodyOriginRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyOriginRulePositionUnion] `json:"position"`
	RouteRuleParam
}

func (r RuleEditParamsBodyOriginRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyOriginRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyOriginRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyOriginRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyOriginRulePosition) implementsRuleEditParamsBodyOriginRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyOriginRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyOriginRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyOriginRulePositionIndexPosition],
// [RuleEditParamsBodyOriginRulePosition].
type RuleEditParamsBodyOriginRulePositionUnion interface {
	implementsRuleEditParamsBodyOriginRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyOriginRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyOriginRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyOriginRulePositionBeforePosition) implementsRuleEditParamsBodyOriginRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyOriginRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyOriginRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyOriginRulePositionAfterPosition) implementsRuleEditParamsBodyOriginRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyOriginRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyOriginRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyOriginRulePositionIndexPosition) implementsRuleEditParamsBodyOriginRulePositionUnion() {
}

type RuleEditParamsBodyScoreRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyScoreRulePositionUnion] `json:"position"`
	ScoreRuleParam
}

func (r RuleEditParamsBodyScoreRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyScoreRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyScoreRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyScoreRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyScoreRulePosition) implementsRuleEditParamsBodyScoreRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyScoreRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyScoreRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyScoreRulePositionIndexPosition],
// [RuleEditParamsBodyScoreRulePosition].
type RuleEditParamsBodyScoreRulePositionUnion interface {
	implementsRuleEditParamsBodyScoreRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyScoreRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyScoreRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyScoreRulePositionBeforePosition) implementsRuleEditParamsBodyScoreRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyScoreRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyScoreRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyScoreRulePositionAfterPosition) implementsRuleEditParamsBodyScoreRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyScoreRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyScoreRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyScoreRulePositionIndexPosition) implementsRuleEditParamsBodyScoreRulePositionUnion() {
}

type RuleEditParamsBodyServeErrorRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyServeErrorRulePositionUnion] `json:"position"`
	ServeErrorRuleParam
}

func (r RuleEditParamsBodyServeErrorRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyServeErrorRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyServeErrorRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyServeErrorRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyServeErrorRulePosition) implementsRuleEditParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyServeErrorRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyServeErrorRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyServeErrorRulePositionIndexPosition],
// [RuleEditParamsBodyServeErrorRulePosition].
type RuleEditParamsBodyServeErrorRulePositionUnion interface {
	implementsRuleEditParamsBodyServeErrorRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyServeErrorRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyServeErrorRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyServeErrorRulePositionBeforePosition) implementsRuleEditParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyServeErrorRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyServeErrorRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyServeErrorRulePositionAfterPosition) implementsRuleEditParamsBodyServeErrorRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyServeErrorRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyServeErrorRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyServeErrorRulePositionIndexPosition) implementsRuleEditParamsBodyServeErrorRulePositionUnion() {
}

type RuleEditParamsBodySetConfigRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodySetConfigRulePositionUnion] `json:"position"`
	SetConfigRuleParam
}

func (r RuleEditParamsBodySetConfigRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetConfigRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetConfigRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySetConfigRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetConfigRulePosition) implementsRuleEditParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodySetConfigRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodySetConfigRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodySetConfigRulePositionIndexPosition],
// [RuleEditParamsBodySetConfigRulePosition].
type RuleEditParamsBodySetConfigRulePositionUnion interface {
	implementsRuleEditParamsBodySetConfigRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetConfigRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodySetConfigRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetConfigRulePositionBeforePosition) implementsRuleEditParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetConfigRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodySetConfigRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetConfigRulePositionAfterPosition) implementsRuleEditParamsBodySetConfigRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetConfigRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySetConfigRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetConfigRulePositionIndexPosition) implementsRuleEditParamsBodySetConfigRulePositionUnion() {
}

type RuleEditParamsBodySkipRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodySkipRulePositionUnion] `json:"position"`
	SkipRuleParam
}

func (r RuleEditParamsBodySkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySkipRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySkipRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySkipRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySkipRulePosition) implementsRuleEditParamsBodySkipRulePositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodySkipRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodySkipRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodySkipRulePositionIndexPosition],
// [RuleEditParamsBodySkipRulePosition].
type RuleEditParamsBodySkipRulePositionUnion interface {
	implementsRuleEditParamsBodySkipRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySkipRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodySkipRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySkipRulePositionBeforePosition) implementsRuleEditParamsBodySkipRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySkipRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodySkipRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySkipRulePositionAfterPosition) implementsRuleEditParamsBodySkipRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySkipRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySkipRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySkipRulePositionIndexPosition) implementsRuleEditParamsBodySkipRulePositionUnion() {
}

type RuleEditParamsBodySetCacheSettingsRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodySetCacheSettingsRulePositionUnion] `json:"position"`
	SetCacheSettingsRuleParam
}

func (r RuleEditParamsBodySetCacheSettingsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetCacheSettingsRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetCacheSettingsRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySetCacheSettingsRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetCacheSettingsRulePosition) implementsRuleEditParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleEditParamsBodySetCacheSettingsRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodySetCacheSettingsRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodySetCacheSettingsRulePositionIndexPosition],
// [RuleEditParamsBodySetCacheSettingsRulePosition].
type RuleEditParamsBodySetCacheSettingsRulePositionUnion interface {
	implementsRuleEditParamsBodySetCacheSettingsRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetCacheSettingsRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionBeforePosition) implementsRuleEditParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetCacheSettingsRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionAfterPosition) implementsRuleEditParamsBodySetCacheSettingsRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodySetCacheSettingsRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodySetCacheSettingsRulePositionIndexPosition) implementsRuleEditParamsBodySetCacheSettingsRulePositionUnion() {
}

type RuleEditParamsBodyLogCustomFieldRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyLogCustomFieldRulePositionUnion] `json:"position"`
	LogCustomFieldRuleParam
}

func (r RuleEditParamsBodyLogCustomFieldRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogCustomFieldRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogCustomFieldRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyLogCustomFieldRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogCustomFieldRulePosition) implementsRuleEditParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleEditParamsBodyLogCustomFieldRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyLogCustomFieldRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyLogCustomFieldRulePositionIndexPosition],
// [RuleEditParamsBodyLogCustomFieldRulePosition].
type RuleEditParamsBodyLogCustomFieldRulePositionUnion interface {
	implementsRuleEditParamsBodyLogCustomFieldRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogCustomFieldRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionBeforePosition) implementsRuleEditParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogCustomFieldRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionAfterPosition) implementsRuleEditParamsBodyLogCustomFieldRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyLogCustomFieldRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyLogCustomFieldRulePositionIndexPosition) implementsRuleEditParamsBodyLogCustomFieldRulePositionUnion() {
}

type RuleEditParamsBodyDDoSDynamicRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyDDoSDynamicRulePositionUnion] `json:"position"`
	DDoSDynamicRuleParam
}

func (r RuleEditParamsBodyDDoSDynamicRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyDDoSDynamicRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyDDoSDynamicRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyDDoSDynamicRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyDDoSDynamicRulePosition) implementsRuleEditParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsBodyDDoSDynamicRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyDDoSDynamicRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyDDoSDynamicRulePositionIndexPosition],
// [RuleEditParamsBodyDDoSDynamicRulePosition].
type RuleEditParamsBodyDDoSDynamicRulePositionUnion interface {
	implementsRuleEditParamsBodyDDoSDynamicRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyDDoSDynamicRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionBeforePosition) implementsRuleEditParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyDDoSDynamicRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionAfterPosition) implementsRuleEditParamsBodyDDoSDynamicRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyDDoSDynamicRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyDDoSDynamicRulePositionIndexPosition) implementsRuleEditParamsBodyDDoSDynamicRulePositionUnion() {
}

type RuleEditParamsBodyForceConnectionCloseRule struct {
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsBodyForceConnectionCloseRulePositionUnion] `json:"position"`
	ForceConnectionCloseRuleParam
}

func (r RuleEditParamsBodyForceConnectionCloseRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyForceConnectionCloseRule) implementsRuleEditParamsBodyUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyForceConnectionCloseRulePosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyForceConnectionCloseRulePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyForceConnectionCloseRulePosition) implementsRuleEditParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
//
// Satisfied by
// [rulesets.RuleEditParamsBodyForceConnectionCloseRulePositionBeforePosition],
// [rulesets.RuleEditParamsBodyForceConnectionCloseRulePositionAfterPosition],
// [rulesets.RuleEditParamsBodyForceConnectionCloseRulePositionIndexPosition],
// [RuleEditParamsBodyForceConnectionCloseRulePosition].
type RuleEditParamsBodyForceConnectionCloseRulePositionUnion interface {
	implementsRuleEditParamsBodyForceConnectionCloseRulePositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyForceConnectionCloseRulePositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionBeforePosition) implementsRuleEditParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyForceConnectionCloseRulePositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionAfterPosition) implementsRuleEditParamsBodyForceConnectionCloseRulePositionUnion() {
}

// An object configuring where the rule will be placed.
type RuleEditParamsBodyForceConnectionCloseRulePositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBodyForceConnectionCloseRulePositionIndexPosition) implementsRuleEditParamsBodyForceConnectionCloseRulePositionUnion() {
}

// The action to perform when the rule matches.
type RuleEditParamsBodyAction string

const (
	RuleEditParamsBodyActionBlock                RuleEditParamsBodyAction = "block"
	RuleEditParamsBodyActionChallenge            RuleEditParamsBodyAction = "challenge"
	RuleEditParamsBodyActionCompressResponse     RuleEditParamsBodyAction = "compress_response"
	RuleEditParamsBodyActionExecute              RuleEditParamsBodyAction = "execute"
	RuleEditParamsBodyActionJSChallenge          RuleEditParamsBodyAction = "js_challenge"
	RuleEditParamsBodyActionLog                  RuleEditParamsBodyAction = "log"
	RuleEditParamsBodyActionManagedChallenge     RuleEditParamsBodyAction = "managed_challenge"
	RuleEditParamsBodyActionRedirect             RuleEditParamsBodyAction = "redirect"
	RuleEditParamsBodyActionRewrite              RuleEditParamsBodyAction = "rewrite"
	RuleEditParamsBodyActionRoute                RuleEditParamsBodyAction = "route"
	RuleEditParamsBodyActionScore                RuleEditParamsBodyAction = "score"
	RuleEditParamsBodyActionServeError           RuleEditParamsBodyAction = "serve_error"
	RuleEditParamsBodyActionSetConfig            RuleEditParamsBodyAction = "set_config"
	RuleEditParamsBodyActionSkip                 RuleEditParamsBodyAction = "skip"
	RuleEditParamsBodyActionSetCacheSettings     RuleEditParamsBodyAction = "set_cache_settings"
	RuleEditParamsBodyActionLogCustomField       RuleEditParamsBodyAction = "log_custom_field"
	RuleEditParamsBodyActionDDoSDynamic          RuleEditParamsBodyAction = "ddos_dynamic"
	RuleEditParamsBodyActionForceConnectionClose RuleEditParamsBodyAction = "force_connection_close"
)

func (r RuleEditParamsBodyAction) IsKnown() bool {
	switch r {
	case RuleEditParamsBodyActionBlock, RuleEditParamsBodyActionChallenge, RuleEditParamsBodyActionCompressResponse, RuleEditParamsBodyActionExecute, RuleEditParamsBodyActionJSChallenge, RuleEditParamsBodyActionLog, RuleEditParamsBodyActionManagedChallenge, RuleEditParamsBodyActionRedirect, RuleEditParamsBodyActionRewrite, RuleEditParamsBodyActionRoute, RuleEditParamsBodyActionScore, RuleEditParamsBodyActionServeError, RuleEditParamsBodyActionSetConfig, RuleEditParamsBodyActionSkip, RuleEditParamsBodyActionSetCacheSettings, RuleEditParamsBodyActionLogCustomField, RuleEditParamsBodyActionDDoSDynamic, RuleEditParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleEditResponseEnvelopeJSON    `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeErrorsSource]
type ruleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeMessagesSource]
type ruleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
