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

// PhaseService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhaseService] method instead.
type PhaseService struct {
	Options  []option.RequestOption
	Versions *PhaseVersionService
}

// NewPhaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPhaseService(opts ...option.RequestOption) (r *PhaseService) {
	r = &PhaseService{}
	r.Options = opts
	r.Versions = NewPhaseVersionService(opts...)
	return
}

// Updates an account or zone entry point ruleset, creating a new version.
func (r *PhaseService) Update(ctx context.Context, rulesetPhase Phase, params PhaseUpdateParams, opts ...option.RequestOption) (res *PhaseUpdateResponse, err error) {
	var env PhaseUpdateResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *PhaseService) Get(ctx context.Context, rulesetPhase Phase, query PhaseGetParams, opts ...option.RequestOption) (res *PhaseGetResponse, err error) {
	var env PhaseGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseUpdateResponse struct {
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
	Rules []PhaseUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                  `json:"description"`
	JSON        phaseUpdateResponseJSON `json:"-"`
}

// phaseUpdateResponseJSON contains the JSON metadata for the struct
// [PhaseUpdateResponse]
type phaseUpdateResponseJSON struct {
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

func (r *PhaseUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesAction `json:"action"`
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
	// [PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
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
	// [PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                      `json:"ref"`
	JSON  phaseUpdateResponseRuleJSON `json:"-"`
	union PhaseUpdateResponseRulesUnion
}

// phaseUpdateResponseRuleJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseRule]
type phaseUpdateResponseRuleJSON struct {
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

func (r phaseUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseUpdateResponseRulesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [PhaseUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [PhaseUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r PhaseUpdateResponseRule) AsUnion() PhaseUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [PhaseUpdateResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [PhaseUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type PhaseUpdateResponseRulesUnion interface {
	implementsPhaseUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsJSChallengeRule{}),
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

type PhaseUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsChallengeRule]
type phaseUpdateResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsChallengeRule) implementsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsChallengeRuleActionChallenge PhaseUpdateResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseUpdateResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                  `json:"username_expression,required"`
	JSON               phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                     `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit]
type phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseUpdateResponseRulesRulesetsJSChallengeRule]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsJSChallengeRule) implementsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression,required"`
	JSON               phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                       `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesAction string

const (
	PhaseUpdateResponseRulesActionBlock                PhaseUpdateResponseRulesAction = "block"
	PhaseUpdateResponseRulesActionChallenge            PhaseUpdateResponseRulesAction = "challenge"
	PhaseUpdateResponseRulesActionCompressResponse     PhaseUpdateResponseRulesAction = "compress_response"
	PhaseUpdateResponseRulesActionExecute              PhaseUpdateResponseRulesAction = "execute"
	PhaseUpdateResponseRulesActionJSChallenge          PhaseUpdateResponseRulesAction = "js_challenge"
	PhaseUpdateResponseRulesActionLog                  PhaseUpdateResponseRulesAction = "log"
	PhaseUpdateResponseRulesActionManagedChallenge     PhaseUpdateResponseRulesAction = "managed_challenge"
	PhaseUpdateResponseRulesActionRedirect             PhaseUpdateResponseRulesAction = "redirect"
	PhaseUpdateResponseRulesActionRewrite              PhaseUpdateResponseRulesAction = "rewrite"
	PhaseUpdateResponseRulesActionRoute                PhaseUpdateResponseRulesAction = "route"
	PhaseUpdateResponseRulesActionScore                PhaseUpdateResponseRulesAction = "score"
	PhaseUpdateResponseRulesActionServeError           PhaseUpdateResponseRulesAction = "serve_error"
	PhaseUpdateResponseRulesActionSetConfig            PhaseUpdateResponseRulesAction = "set_config"
	PhaseUpdateResponseRulesActionSkip                 PhaseUpdateResponseRulesAction = "skip"
	PhaseUpdateResponseRulesActionSetCacheSettings     PhaseUpdateResponseRulesAction = "set_cache_settings"
	PhaseUpdateResponseRulesActionLogCustomField       PhaseUpdateResponseRulesAction = "log_custom_field"
	PhaseUpdateResponseRulesActionDDoSDynamic          PhaseUpdateResponseRulesAction = "ddos_dynamic"
	PhaseUpdateResponseRulesActionForceConnectionClose PhaseUpdateResponseRulesAction = "force_connection_close"
)

func (r PhaseUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesActionBlock, PhaseUpdateResponseRulesActionChallenge, PhaseUpdateResponseRulesActionCompressResponse, PhaseUpdateResponseRulesActionExecute, PhaseUpdateResponseRulesActionJSChallenge, PhaseUpdateResponseRulesActionLog, PhaseUpdateResponseRulesActionManagedChallenge, PhaseUpdateResponseRulesActionRedirect, PhaseUpdateResponseRulesActionRewrite, PhaseUpdateResponseRulesActionRoute, PhaseUpdateResponseRulesActionScore, PhaseUpdateResponseRulesActionServeError, PhaseUpdateResponseRulesActionSetConfig, PhaseUpdateResponseRulesActionSkip, PhaseUpdateResponseRulesActionSetCacheSettings, PhaseUpdateResponseRulesActionLogCustomField, PhaseUpdateResponseRulesActionDDoSDynamic, PhaseUpdateResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A ruleset object.
type PhaseGetResponse struct {
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
	Rules []PhaseGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        phaseGetResponseJSON `json:"-"`
}

// phaseGetResponseJSON contains the JSON metadata for the struct
// [PhaseGetResponse]
type phaseGetResponseJSON struct {
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

func (r *PhaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesAction `json:"action"`
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
	// [PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
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
	// [PhaseGetResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                   `json:"ref"`
	JSON  phaseGetResponseRuleJSON `json:"-"`
	union PhaseGetResponseRulesUnion
}

// phaseGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseGetResponseRule]
type phaseGetResponseRuleJSON struct {
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

func (r phaseGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseGetResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [PhaseGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [PhaseGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [PhaseGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [PhaseGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type PhaseGetResponseRulesUnion interface {
	implementsPhaseGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsJSChallengeRule{}),
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

type PhaseGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON phaseGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsChallengeRule]
type phaseGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsChallengeRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsChallengeRuleActionChallenge PhaseGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                               `json:"username_expression,required"`
	JSON               phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type PhaseGetResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	JSON                    phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsChallengeRuleRatelimit]
type phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON phaseGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsJSChallengeRule]
type phaseGetResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsJSChallengeRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge PhaseGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression,required"`
	JSON               phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit]
type phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock                PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionChallenge            PhaseGetResponseRulesAction = "challenge"
	PhaseGetResponseRulesActionCompressResponse     PhaseGetResponseRulesAction = "compress_response"
	PhaseGetResponseRulesActionExecute              PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionJSChallenge          PhaseGetResponseRulesAction = "js_challenge"
	PhaseGetResponseRulesActionLog                  PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionManagedChallenge     PhaseGetResponseRulesAction = "managed_challenge"
	PhaseGetResponseRulesActionRedirect             PhaseGetResponseRulesAction = "redirect"
	PhaseGetResponseRulesActionRewrite              PhaseGetResponseRulesAction = "rewrite"
	PhaseGetResponseRulesActionRoute                PhaseGetResponseRulesAction = "route"
	PhaseGetResponseRulesActionScore                PhaseGetResponseRulesAction = "score"
	PhaseGetResponseRulesActionServeError           PhaseGetResponseRulesAction = "serve_error"
	PhaseGetResponseRulesActionSetConfig            PhaseGetResponseRulesAction = "set_config"
	PhaseGetResponseRulesActionSkip                 PhaseGetResponseRulesAction = "skip"
	PhaseGetResponseRulesActionSetCacheSettings     PhaseGetResponseRulesAction = "set_cache_settings"
	PhaseGetResponseRulesActionLogCustomField       PhaseGetResponseRulesAction = "log_custom_field"
	PhaseGetResponseRulesActionDDoSDynamic          PhaseGetResponseRulesAction = "ddos_dynamic"
	PhaseGetResponseRulesActionForceConnectionClose PhaseGetResponseRulesAction = "force_connection_close"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionChallenge, PhaseGetResponseRulesActionCompressResponse, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionJSChallenge, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionManagedChallenge, PhaseGetResponseRulesActionRedirect, PhaseGetResponseRulesActionRewrite, PhaseGetResponseRulesActionRoute, PhaseGetResponseRulesActionScore, PhaseGetResponseRulesActionServeError, PhaseGetResponseRulesActionSetConfig, PhaseGetResponseRulesActionSkip, PhaseGetResponseRulesActionSetCacheSettings, PhaseGetResponseRulesActionLogCustomField, PhaseGetResponseRulesActionDDoSDynamic, PhaseGetResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRuleUnion] `json:"rules"`
}

func (r PhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[PhaseUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                  `json:"action_parameters"`
	Categories       param.Field[interface{}]                  `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRule) implementsPhaseUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.ManagedChallengeRuleParam],
// [rulesets.RedirectRuleParam], [rulesets.RewriteRuleParam],
// [rulesets.RouteRuleParam], [rulesets.ScoreRuleParam],
// [rulesets.ServeErrorRuleParam], [rulesets.SetConfigRuleParam],
// [rulesets.SkipRuleParam], [rulesets.SetCacheSettingsRuleParam],
// [rulesets.LogCustomFieldRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ForceConnectionCloseRuleParam], [PhaseUpdateParamsRule].
type PhaseUpdateParamsRuleUnion interface {
	implementsPhaseUpdateParamsRuleUnion()
}

type PhaseUpdateParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) implementsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsChallengeRuleActionChallenge PhaseUpdateParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit struct {
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

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRule) implementsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit struct {
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

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock                PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionChallenge            PhaseUpdateParamsRulesAction = "challenge"
	PhaseUpdateParamsRulesActionCompressResponse     PhaseUpdateParamsRulesAction = "compress_response"
	PhaseUpdateParamsRulesActionExecute              PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionJSChallenge          PhaseUpdateParamsRulesAction = "js_challenge"
	PhaseUpdateParamsRulesActionLog                  PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionManagedChallenge     PhaseUpdateParamsRulesAction = "managed_challenge"
	PhaseUpdateParamsRulesActionRedirect             PhaseUpdateParamsRulesAction = "redirect"
	PhaseUpdateParamsRulesActionRewrite              PhaseUpdateParamsRulesAction = "rewrite"
	PhaseUpdateParamsRulesActionRoute                PhaseUpdateParamsRulesAction = "route"
	PhaseUpdateParamsRulesActionScore                PhaseUpdateParamsRulesAction = "score"
	PhaseUpdateParamsRulesActionServeError           PhaseUpdateParamsRulesAction = "serve_error"
	PhaseUpdateParamsRulesActionSetConfig            PhaseUpdateParamsRulesAction = "set_config"
	PhaseUpdateParamsRulesActionSkip                 PhaseUpdateParamsRulesAction = "skip"
	PhaseUpdateParamsRulesActionSetCacheSettings     PhaseUpdateParamsRulesAction = "set_cache_settings"
	PhaseUpdateParamsRulesActionLogCustomField       PhaseUpdateParamsRulesAction = "log_custom_field"
	PhaseUpdateParamsRulesActionDDoSDynamic          PhaseUpdateParamsRulesAction = "ddos_dynamic"
	PhaseUpdateParamsRulesActionForceConnectionClose PhaseUpdateParamsRulesAction = "force_connection_close"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionChallenge, PhaseUpdateParamsRulesActionCompressResponse, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionJSChallenge, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionManagedChallenge, PhaseUpdateParamsRulesActionRedirect, PhaseUpdateParamsRulesActionRewrite, PhaseUpdateParamsRulesActionRoute, PhaseUpdateParamsRulesActionScore, PhaseUpdateParamsRulesActionServeError, PhaseUpdateParamsRulesActionSetConfig, PhaseUpdateParamsRulesActionSkip, PhaseUpdateParamsRulesActionSetCacheSettings, PhaseUpdateParamsRulesActionLogCustomField, PhaseUpdateParamsRulesActionDDoSDynamic, PhaseUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseUpdateResponseEnvelopeJSON    `json:"-"`
}

// phaseUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelope]
type phaseUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeErrors]
type phaseUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeErrorsSource]
type phaseUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessages]
type phaseUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessagesSource]
type phaseUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseUpdateResponseEnvelopeSuccess bool

const (
	PhaseUpdateResponseEnvelopeSuccessTrue PhaseUpdateResponseEnvelopeSuccess = true
)

func (r PhaseUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type PhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseGetResponseEnvelopeJSON    `json:"-"`
}

// phaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelope]
type phaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeErrors]
type phaseGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeErrorsSource]
type phaseGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeMessages]
type phaseGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeMessagesSource]
type phaseGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseGetResponseEnvelopeSuccess bool

const (
	PhaseGetResponseEnvelopeSuccessTrue PhaseGetResponseEnvelopeSuccess = true
)

func (r PhaseGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
