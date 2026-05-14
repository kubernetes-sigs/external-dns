// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// RulesetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRulesetService] method instead.
type RulesetService struct {
	Options  []option.RequestOption
	Phases   *PhaseService
	Rules    *RuleService
	Versions *VersionService
}

// NewRulesetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRulesetService(opts ...option.RequestOption) (r *RulesetService) {
	r = &RulesetService{}
	r.Options = opts
	r.Phases = NewPhaseService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Versions = NewVersionService(opts...)
	return
}

// Creates a ruleset.
func (r *RulesetService) New(ctx context.Context, params RulesetNewParams, opts ...option.RequestOption) (res *RulesetNewResponse, err error) {
	var env RulesetNewResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an account or zone ruleset, creating a new version.
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetUpdateResponse, err error) {
	var env RulesetUpdateResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all rulesets.
func (r *RulesetService) List(ctx context.Context, params RulesetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RulesetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches all rulesets.
func (r *RulesetService) ListAutoPaging(ctx context.Context, params RulesetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RulesetListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes all versions of an existing account or zone ruleset.
func (r *RulesetService) Delete(ctx context.Context, rulesetID string, body RulesetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, rulesetID string, query RulesetGetParams, opts ...option.RequestOption) (res *RulesetGetResponse, err error) {
	var env RulesetGetResponseEnvelope
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The kind of the ruleset.
type Kind string

const (
	KindManaged Kind = "managed"
	KindCustom  Kind = "custom"
	KindRoot    Kind = "root"
	KindZone    Kind = "zone"
)

func (r Kind) IsKnown() bool {
	switch r {
	case KindManaged, KindCustom, KindRoot, KindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type Phase string

const (
	PhaseDDoSL4                       Phase = "ddos_l4"
	PhaseDDoSL7                       Phase = "ddos_l7"
	PhaseHTTPConfigSettings           Phase = "http_config_settings"
	PhaseHTTPCustomErrors             Phase = "http_custom_errors"
	PhaseHTTPLogCustomFields          Phase = "http_log_custom_fields"
	PhaseHTTPRatelimit                Phase = "http_ratelimit"
	PhaseHTTPRequestCacheSettings     Phase = "http_request_cache_settings"
	PhaseHTTPRequestDynamicRedirect   Phase = "http_request_dynamic_redirect"
	PhaseHTTPRequestFirewallCustom    Phase = "http_request_firewall_custom"
	PhaseHTTPRequestFirewallManaged   Phase = "http_request_firewall_managed"
	PhaseHTTPRequestLateTransform     Phase = "http_request_late_transform"
	PhaseHTTPRequestOrigin            Phase = "http_request_origin"
	PhaseHTTPRequestRedirect          Phase = "http_request_redirect"
	PhaseHTTPRequestSanitize          Phase = "http_request_sanitize"
	PhaseHTTPRequestSBFM              Phase = "http_request_sbfm"
	PhaseHTTPRequestTransform         Phase = "http_request_transform"
	PhaseHTTPResponseCompression      Phase = "http_response_compression"
	PhaseHTTPResponseFirewallManaged  Phase = "http_response_firewall_managed"
	PhaseHTTPResponseHeadersTransform Phase = "http_response_headers_transform"
	PhaseMagicTransit                 Phase = "magic_transit"
	PhaseMagicTransitIDsManaged       Phase = "magic_transit_ids_managed"
	PhaseMagicTransitManaged          Phase = "magic_transit_managed"
	PhaseMagicTransitRatelimit        Phase = "magic_transit_ratelimit"
)

func (r Phase) IsKnown() bool {
	switch r {
	case PhaseDDoSL4, PhaseDDoSL7, PhaseHTTPConfigSettings, PhaseHTTPCustomErrors, PhaseHTTPLogCustomFields, PhaseHTTPRatelimit, PhaseHTTPRequestCacheSettings, PhaseHTTPRequestDynamicRedirect, PhaseHTTPRequestFirewallCustom, PhaseHTTPRequestFirewallManaged, PhaseHTTPRequestLateTransform, PhaseHTTPRequestOrigin, PhaseHTTPRequestRedirect, PhaseHTTPRequestSanitize, PhaseHTTPRequestSBFM, PhaseHTTPRequestTransform, PhaseHTTPResponseCompression, PhaseHTTPResponseFirewallManaged, PhaseHTTPResponseHeadersTransform, PhaseMagicTransit, PhaseMagicTransitIDsManaged, PhaseMagicTransitManaged, PhaseMagicTransitRatelimit:
		return true
	}
	return false
}

// A ruleset object.
type RulesetNewResponse struct {
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
	Rules []RulesetNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetNewResponseJSON `json:"-"`
}

// rulesetNewResponseJSON contains the JSON metadata for the struct
// [RulesetNewResponse]
type rulesetNewResponseJSON struct {
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

func (r *RulesetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesAction `json:"action"`
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
	// [RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
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
	// [RulesetNewResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                     `json:"ref"`
	JSON  rulesetNewResponseRuleJSON `json:"-"`
	union RulesetNewResponseRulesUnion
}

// rulesetNewResponseRuleJSON contains the JSON metadata for the struct
// [RulesetNewResponseRule]
type rulesetNewResponseRuleJSON struct {
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

func (r rulesetNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetNewResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetNewResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RulesetNewResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RulesetNewResponseRule) AsUnion() RulesetNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RulesetNewResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [RulesetNewResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RulesetNewResponseRulesUnion interface {
	implementsRulesetNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsJSChallengeRule{}),
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

type RulesetNewResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsChallengeRule]
type rulesetNewResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsChallengeRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge RulesetNewResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression,required"`
	JSON               rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetNewResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	JSON                    rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsJSChallengeRule]
type rulesetNewResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsJSChallengeRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetNewResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetNewResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression,required"`
	JSON               rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesAction string

const (
	RulesetNewResponseRulesActionBlock                RulesetNewResponseRulesAction = "block"
	RulesetNewResponseRulesActionChallenge            RulesetNewResponseRulesAction = "challenge"
	RulesetNewResponseRulesActionCompressResponse     RulesetNewResponseRulesAction = "compress_response"
	RulesetNewResponseRulesActionExecute              RulesetNewResponseRulesAction = "execute"
	RulesetNewResponseRulesActionJSChallenge          RulesetNewResponseRulesAction = "js_challenge"
	RulesetNewResponseRulesActionLog                  RulesetNewResponseRulesAction = "log"
	RulesetNewResponseRulesActionManagedChallenge     RulesetNewResponseRulesAction = "managed_challenge"
	RulesetNewResponseRulesActionRedirect             RulesetNewResponseRulesAction = "redirect"
	RulesetNewResponseRulesActionRewrite              RulesetNewResponseRulesAction = "rewrite"
	RulesetNewResponseRulesActionRoute                RulesetNewResponseRulesAction = "route"
	RulesetNewResponseRulesActionScore                RulesetNewResponseRulesAction = "score"
	RulesetNewResponseRulesActionServeError           RulesetNewResponseRulesAction = "serve_error"
	RulesetNewResponseRulesActionSetConfig            RulesetNewResponseRulesAction = "set_config"
	RulesetNewResponseRulesActionSkip                 RulesetNewResponseRulesAction = "skip"
	RulesetNewResponseRulesActionSetCacheSettings     RulesetNewResponseRulesAction = "set_cache_settings"
	RulesetNewResponseRulesActionLogCustomField       RulesetNewResponseRulesAction = "log_custom_field"
	RulesetNewResponseRulesActionDDoSDynamic          RulesetNewResponseRulesAction = "ddos_dynamic"
	RulesetNewResponseRulesActionForceConnectionClose RulesetNewResponseRulesAction = "force_connection_close"
)

func (r RulesetNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesActionBlock, RulesetNewResponseRulesActionChallenge, RulesetNewResponseRulesActionCompressResponse, RulesetNewResponseRulesActionExecute, RulesetNewResponseRulesActionJSChallenge, RulesetNewResponseRulesActionLog, RulesetNewResponseRulesActionManagedChallenge, RulesetNewResponseRulesActionRedirect, RulesetNewResponseRulesActionRewrite, RulesetNewResponseRulesActionRoute, RulesetNewResponseRulesActionScore, RulesetNewResponseRulesActionServeError, RulesetNewResponseRulesActionSetConfig, RulesetNewResponseRulesActionSkip, RulesetNewResponseRulesActionSetCacheSettings, RulesetNewResponseRulesActionLogCustomField, RulesetNewResponseRulesActionDDoSDynamic, RulesetNewResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A ruleset object.
type RulesetUpdateResponse struct {
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
	Rules []RulesetUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                    `json:"description"`
	JSON        rulesetUpdateResponseJSON `json:"-"`
}

// rulesetUpdateResponseJSON contains the JSON metadata for the struct
// [RulesetUpdateResponse]
type rulesetUpdateResponseJSON struct {
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

func (r *RulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesAction `json:"action"`
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
	// [RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
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
	// [RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                        `json:"ref"`
	JSON  rulesetUpdateResponseRuleJSON `json:"-"`
	union RulesetUpdateResponseRulesUnion
}

// rulesetUpdateResponseRuleJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseRule]
type rulesetUpdateResponseRuleJSON struct {
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

func (r rulesetUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetUpdateResponseRulesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RulesetUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RulesetUpdateResponseRule) AsUnion() RulesetUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule],
// [RulesetUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RulesetUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RulesetUpdateResponseRulesUnion interface {
	implementsRulesetUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsJSChallengeRule{}),
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

type RulesetUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsChallengeRule]
type rulesetUpdateResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsChallengeRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge RulesetUpdateResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression,required"`
	JSON               rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	JSON                    rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsJSChallengeRule]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsJSChallengeRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                      `json:"username_expression,required"`
	JSON               rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                         `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesAction string

const (
	RulesetUpdateResponseRulesActionBlock                RulesetUpdateResponseRulesAction = "block"
	RulesetUpdateResponseRulesActionChallenge            RulesetUpdateResponseRulesAction = "challenge"
	RulesetUpdateResponseRulesActionCompressResponse     RulesetUpdateResponseRulesAction = "compress_response"
	RulesetUpdateResponseRulesActionExecute              RulesetUpdateResponseRulesAction = "execute"
	RulesetUpdateResponseRulesActionJSChallenge          RulesetUpdateResponseRulesAction = "js_challenge"
	RulesetUpdateResponseRulesActionLog                  RulesetUpdateResponseRulesAction = "log"
	RulesetUpdateResponseRulesActionManagedChallenge     RulesetUpdateResponseRulesAction = "managed_challenge"
	RulesetUpdateResponseRulesActionRedirect             RulesetUpdateResponseRulesAction = "redirect"
	RulesetUpdateResponseRulesActionRewrite              RulesetUpdateResponseRulesAction = "rewrite"
	RulesetUpdateResponseRulesActionRoute                RulesetUpdateResponseRulesAction = "route"
	RulesetUpdateResponseRulesActionScore                RulesetUpdateResponseRulesAction = "score"
	RulesetUpdateResponseRulesActionServeError           RulesetUpdateResponseRulesAction = "serve_error"
	RulesetUpdateResponseRulesActionSetConfig            RulesetUpdateResponseRulesAction = "set_config"
	RulesetUpdateResponseRulesActionSkip                 RulesetUpdateResponseRulesAction = "skip"
	RulesetUpdateResponseRulesActionSetCacheSettings     RulesetUpdateResponseRulesAction = "set_cache_settings"
	RulesetUpdateResponseRulesActionLogCustomField       RulesetUpdateResponseRulesAction = "log_custom_field"
	RulesetUpdateResponseRulesActionDDoSDynamic          RulesetUpdateResponseRulesAction = "ddos_dynamic"
	RulesetUpdateResponseRulesActionForceConnectionClose RulesetUpdateResponseRulesAction = "force_connection_close"
)

func (r RulesetUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesActionBlock, RulesetUpdateResponseRulesActionChallenge, RulesetUpdateResponseRulesActionCompressResponse, RulesetUpdateResponseRulesActionExecute, RulesetUpdateResponseRulesActionJSChallenge, RulesetUpdateResponseRulesActionLog, RulesetUpdateResponseRulesActionManagedChallenge, RulesetUpdateResponseRulesActionRedirect, RulesetUpdateResponseRulesActionRewrite, RulesetUpdateResponseRulesActionRoute, RulesetUpdateResponseRulesActionScore, RulesetUpdateResponseRulesActionServeError, RulesetUpdateResponseRulesActionSetConfig, RulesetUpdateResponseRulesActionSkip, RulesetUpdateResponseRulesActionSetCacheSettings, RulesetUpdateResponseRulesActionLogCustomField, RulesetUpdateResponseRulesActionDDoSDynamic, RulesetUpdateResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A ruleset object.
type RulesetListResponse struct {
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
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                  `json:"description"`
	JSON        rulesetListResponseJSON `json:"-"`
}

// rulesetListResponseJSON contains the JSON metadata for the struct
// [RulesetListResponse]
type rulesetListResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetListResponseJSON) RawJSON() string {
	return r.raw
}

// A ruleset object.
type RulesetGetResponse struct {
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
	Rules []RulesetGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetGetResponseJSON `json:"-"`
}

// rulesetGetResponseJSON contains the JSON metadata for the struct
// [RulesetGetResponse]
type rulesetGetResponseJSON struct {
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

func (r *RulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesAction `json:"action"`
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
	// [RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [ExecuteRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
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
	// [RulesetGetResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [ExecuteRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit], [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [DDoSDynamicRuleRatelimit], [ForceConnectionCloseRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref   string                     `json:"ref"`
	JSON  rulesetGetResponseRuleJSON `json:"-"`
	union RulesetGetResponseRulesUnion
}

// rulesetGetResponseRuleJSON contains the JSON metadata for the struct
// [RulesetGetResponseRule]
type rulesetGetResponseRuleJSON struct {
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

func (r rulesetGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetGetResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [ExecuteRule], [RulesetGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule], [ForceConnectionCloseRule].
func (r RulesetGetResponseRule) AsUnion() RulesetGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RulesetGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [ExecuteRule],
// [RulesetGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [SetConfigRule], [SkipRule], [SetCacheSettingsRule],
// [LogCustomFieldRule], [DDoSDynamicRule] or [ForceConnectionCloseRule].
type RulesetGetResponseRulesUnion interface {
	implementsRulesetGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsJSChallengeRule{}),
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

type RulesetGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsChallengeRule]
type rulesetGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsChallengeRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge RulesetGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression,required"`
	JSON               rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetGetResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	JSON                    rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsJSChallengeRule]
type rulesetGetResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsJSChallengeRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression,required"`
	JSON               rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's ratelimit behavior.
type RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesAction string

const (
	RulesetGetResponseRulesActionBlock                RulesetGetResponseRulesAction = "block"
	RulesetGetResponseRulesActionChallenge            RulesetGetResponseRulesAction = "challenge"
	RulesetGetResponseRulesActionCompressResponse     RulesetGetResponseRulesAction = "compress_response"
	RulesetGetResponseRulesActionExecute              RulesetGetResponseRulesAction = "execute"
	RulesetGetResponseRulesActionJSChallenge          RulesetGetResponseRulesAction = "js_challenge"
	RulesetGetResponseRulesActionLog                  RulesetGetResponseRulesAction = "log"
	RulesetGetResponseRulesActionManagedChallenge     RulesetGetResponseRulesAction = "managed_challenge"
	RulesetGetResponseRulesActionRedirect             RulesetGetResponseRulesAction = "redirect"
	RulesetGetResponseRulesActionRewrite              RulesetGetResponseRulesAction = "rewrite"
	RulesetGetResponseRulesActionRoute                RulesetGetResponseRulesAction = "route"
	RulesetGetResponseRulesActionScore                RulesetGetResponseRulesAction = "score"
	RulesetGetResponseRulesActionServeError           RulesetGetResponseRulesAction = "serve_error"
	RulesetGetResponseRulesActionSetConfig            RulesetGetResponseRulesAction = "set_config"
	RulesetGetResponseRulesActionSkip                 RulesetGetResponseRulesAction = "skip"
	RulesetGetResponseRulesActionSetCacheSettings     RulesetGetResponseRulesAction = "set_cache_settings"
	RulesetGetResponseRulesActionLogCustomField       RulesetGetResponseRulesAction = "log_custom_field"
	RulesetGetResponseRulesActionDDoSDynamic          RulesetGetResponseRulesAction = "ddos_dynamic"
	RulesetGetResponseRulesActionForceConnectionClose RulesetGetResponseRulesAction = "force_connection_close"
)

func (r RulesetGetResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesActionBlock, RulesetGetResponseRulesActionChallenge, RulesetGetResponseRulesActionCompressResponse, RulesetGetResponseRulesActionExecute, RulesetGetResponseRulesActionJSChallenge, RulesetGetResponseRulesActionLog, RulesetGetResponseRulesActionManagedChallenge, RulesetGetResponseRulesActionRedirect, RulesetGetResponseRulesActionRewrite, RulesetGetResponseRulesActionRoute, RulesetGetResponseRulesActionScore, RulesetGetResponseRulesActionServeError, RulesetGetResponseRulesActionSetConfig, RulesetGetResponseRulesActionSkip, RulesetGetResponseRulesActionSetCacheSettings, RulesetGetResponseRulesActionLogCustomField, RulesetGetResponseRulesActionDDoSDynamic, RulesetGetResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

type RulesetNewParams struct {
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetNewParamsRuleUnion] `json:"rules"`
}

func (r RulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                 `json:"action_parameters"`
	Categories       param.Field[interface{}]                 `json:"categories"`
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

func (r RulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRule) implementsRulesetNewParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.ManagedChallengeRuleParam],
// [rulesets.RedirectRuleParam], [rulesets.RewriteRuleParam],
// [rulesets.RouteRuleParam], [rulesets.ScoreRuleParam],
// [rulesets.ServeErrorRuleParam], [rulesets.SetConfigRuleParam],
// [rulesets.SkipRuleParam], [rulesets.SetCacheSettingsRuleParam],
// [rulesets.LogCustomFieldRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ForceConnectionCloseRuleParam], [RulesetNewParamsRule].
type RulesetNewParamsRuleUnion interface {
	implementsRulesetNewParamsRuleUnion()
}

type RulesetNewParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsChallengeRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge RulesetNewParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RulesetNewParamsRulesRulesetsChallengeRuleRatelimit struct {
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

func (r RulesetNewParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsJSChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsJSChallengeRuleActionJSChallenge RulesetNewParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit struct {
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

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesAction string

const (
	RulesetNewParamsRulesActionBlock                RulesetNewParamsRulesAction = "block"
	RulesetNewParamsRulesActionChallenge            RulesetNewParamsRulesAction = "challenge"
	RulesetNewParamsRulesActionCompressResponse     RulesetNewParamsRulesAction = "compress_response"
	RulesetNewParamsRulesActionExecute              RulesetNewParamsRulesAction = "execute"
	RulesetNewParamsRulesActionJSChallenge          RulesetNewParamsRulesAction = "js_challenge"
	RulesetNewParamsRulesActionLog                  RulesetNewParamsRulesAction = "log"
	RulesetNewParamsRulesActionManagedChallenge     RulesetNewParamsRulesAction = "managed_challenge"
	RulesetNewParamsRulesActionRedirect             RulesetNewParamsRulesAction = "redirect"
	RulesetNewParamsRulesActionRewrite              RulesetNewParamsRulesAction = "rewrite"
	RulesetNewParamsRulesActionRoute                RulesetNewParamsRulesAction = "route"
	RulesetNewParamsRulesActionScore                RulesetNewParamsRulesAction = "score"
	RulesetNewParamsRulesActionServeError           RulesetNewParamsRulesAction = "serve_error"
	RulesetNewParamsRulesActionSetConfig            RulesetNewParamsRulesAction = "set_config"
	RulesetNewParamsRulesActionSkip                 RulesetNewParamsRulesAction = "skip"
	RulesetNewParamsRulesActionSetCacheSettings     RulesetNewParamsRulesAction = "set_cache_settings"
	RulesetNewParamsRulesActionLogCustomField       RulesetNewParamsRulesAction = "log_custom_field"
	RulesetNewParamsRulesActionDDoSDynamic          RulesetNewParamsRulesAction = "ddos_dynamic"
	RulesetNewParamsRulesActionForceConnectionClose RulesetNewParamsRulesAction = "force_connection_close"
)

func (r RulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesActionBlock, RulesetNewParamsRulesActionChallenge, RulesetNewParamsRulesActionCompressResponse, RulesetNewParamsRulesActionExecute, RulesetNewParamsRulesActionJSChallenge, RulesetNewParamsRulesActionLog, RulesetNewParamsRulesActionManagedChallenge, RulesetNewParamsRulesActionRedirect, RulesetNewParamsRulesActionRewrite, RulesetNewParamsRulesActionRoute, RulesetNewParamsRulesActionScore, RulesetNewParamsRulesActionServeError, RulesetNewParamsRulesActionSetConfig, RulesetNewParamsRulesActionSkip, RulesetNewParamsRulesActionSetCacheSettings, RulesetNewParamsRulesActionLogCustomField, RulesetNewParamsRulesActionDDoSDynamic, RulesetNewParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A response object.
type RulesetNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetNewResponseEnvelopeJSON    `json:"-"`
}

// rulesetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelope]
type rulesetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeErrors]
type rulesetNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    rulesetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeErrorsSource]
type rulesetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeMessages]
type rulesetNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    rulesetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeMessagesSource]
type rulesetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetNewResponseEnvelopeSuccess bool

const (
	RulesetNewResponseEnvelopeSuccessTrue RulesetNewResponseEnvelopeSuccess = true
)

func (r RulesetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RulesetUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetUpdateParamsRuleUnion] `json:"rules"`
}

func (r RulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RulesetUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                    `json:"action_parameters"`
	Categories       param.Field[interface{}]                    `json:"categories"`
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

func (r RulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRule) implementsRulesetUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.ManagedChallengeRuleParam],
// [rulesets.RedirectRuleParam], [rulesets.RewriteRuleParam],
// [rulesets.RouteRuleParam], [rulesets.ScoreRuleParam],
// [rulesets.ServeErrorRuleParam], [rulesets.SetConfigRuleParam],
// [rulesets.SkipRuleParam], [rulesets.SetCacheSettingsRuleParam],
// [rulesets.LogCustomFieldRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ForceConnectionCloseRuleParam], [RulesetUpdateParamsRule].
type RulesetUpdateParamsRuleUnion interface {
	implementsRulesetUpdateParamsRuleUnion()
}

type RulesetUpdateParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) implementsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge RulesetUpdateParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit struct {
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

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRule) implementsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configure checks for exposed credentials.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit struct {
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

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesAction string

const (
	RulesetUpdateParamsRulesActionBlock                RulesetUpdateParamsRulesAction = "block"
	RulesetUpdateParamsRulesActionChallenge            RulesetUpdateParamsRulesAction = "challenge"
	RulesetUpdateParamsRulesActionCompressResponse     RulesetUpdateParamsRulesAction = "compress_response"
	RulesetUpdateParamsRulesActionExecute              RulesetUpdateParamsRulesAction = "execute"
	RulesetUpdateParamsRulesActionJSChallenge          RulesetUpdateParamsRulesAction = "js_challenge"
	RulesetUpdateParamsRulesActionLog                  RulesetUpdateParamsRulesAction = "log"
	RulesetUpdateParamsRulesActionManagedChallenge     RulesetUpdateParamsRulesAction = "managed_challenge"
	RulesetUpdateParamsRulesActionRedirect             RulesetUpdateParamsRulesAction = "redirect"
	RulesetUpdateParamsRulesActionRewrite              RulesetUpdateParamsRulesAction = "rewrite"
	RulesetUpdateParamsRulesActionRoute                RulesetUpdateParamsRulesAction = "route"
	RulesetUpdateParamsRulesActionScore                RulesetUpdateParamsRulesAction = "score"
	RulesetUpdateParamsRulesActionServeError           RulesetUpdateParamsRulesAction = "serve_error"
	RulesetUpdateParamsRulesActionSetConfig            RulesetUpdateParamsRulesAction = "set_config"
	RulesetUpdateParamsRulesActionSkip                 RulesetUpdateParamsRulesAction = "skip"
	RulesetUpdateParamsRulesActionSetCacheSettings     RulesetUpdateParamsRulesAction = "set_cache_settings"
	RulesetUpdateParamsRulesActionLogCustomField       RulesetUpdateParamsRulesAction = "log_custom_field"
	RulesetUpdateParamsRulesActionDDoSDynamic          RulesetUpdateParamsRulesAction = "ddos_dynamic"
	RulesetUpdateParamsRulesActionForceConnectionClose RulesetUpdateParamsRulesAction = "force_connection_close"
)

func (r RulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesActionBlock, RulesetUpdateParamsRulesActionChallenge, RulesetUpdateParamsRulesActionCompressResponse, RulesetUpdateParamsRulesActionExecute, RulesetUpdateParamsRulesActionJSChallenge, RulesetUpdateParamsRulesActionLog, RulesetUpdateParamsRulesActionManagedChallenge, RulesetUpdateParamsRulesActionRedirect, RulesetUpdateParamsRulesActionRewrite, RulesetUpdateParamsRulesActionRoute, RulesetUpdateParamsRulesActionScore, RulesetUpdateParamsRulesActionServeError, RulesetUpdateParamsRulesActionSetConfig, RulesetUpdateParamsRulesActionSkip, RulesetUpdateParamsRulesActionSetCacheSettings, RulesetUpdateParamsRulesActionLogCustomField, RulesetUpdateParamsRulesActionDDoSDynamic, RulesetUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

// A response object.
type RulesetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetUpdateResponseEnvelopeJSON    `json:"-"`
}

// rulesetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelope]
type rulesetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrors]
type rulesetUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    rulesetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrorsSource]
type rulesetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeMessages]
type rulesetUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    rulesetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseEnvelopeMessagesSource]
type rulesetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetUpdateResponseEnvelopeSuccess bool

const (
	RulesetUpdateResponseEnvelopeSuccessTrue RulesetUpdateResponseEnvelopeSuccess = true
)

func (r RulesetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RulesetListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Cursor to use for the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Number of rulesets to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [RulesetListParams]'s query parameters as `url.Values`.
func (r RulesetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RulesetDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type RulesetGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RulesetGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelope]
type rulesetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeErrors]
type rulesetGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    rulesetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeErrorsSource]
type rulesetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeMessages]
type rulesetGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    rulesetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeMessagesSource]
type rulesetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetGetResponseEnvelopeSuccess bool

const (
	RulesetGetResponseEnvelopeSuccessTrue RulesetGetResponseEnvelopeSuccess = true
)

func (r RulesetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
