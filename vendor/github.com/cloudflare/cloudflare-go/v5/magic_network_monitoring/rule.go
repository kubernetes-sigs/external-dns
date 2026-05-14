// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options        []option.RequestOption
	Advertisements *RuleAdvertisementService
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	r.Advertisements = NewRuleAdvertisementService(opts...)
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringRule, err error) {
	var env RuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update network monitoring rules for account.
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringRule, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists network monitoring rules for account.
func (r *RuleService) List(ctx context.Context, query RuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[MagicNetworkMonitoringRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists network monitoring rules for account.
func (r *RuleService) ListAutoPaging(ctx context.Context, query RuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MagicNetworkMonitoringRule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a network monitoring rule for account.
func (r *RuleService) Delete(ctx context.Context, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringRule, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a network monitoring rule for account.
func (r *RuleService) Edit(ctx context.Context, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringRule, err error) {
	var env RuleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single network monitoring rule for account.
func (r *RuleService) Get(ctx context.Context, ruleID string, query RuleGetParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringRule, err error) {
	var env RuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringRule struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string   `json:"name,required"`
	Prefixes []string `json:"prefixes,required"`
	// MNM rule type.
	Type MagicNetworkMonitoringRuleType `json:"type,required"`
	// The id of the rule. Must be unique.
	ID string `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration MagicNetworkMonitoringRuleDuration `json:"duration"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64 `json:"packet_threshold"`
	// Prefix match type to be applied for a prefix auto advertisement when using an
	// advanced_ddos rule.
	PrefixMatch MagicNetworkMonitoringRulePrefixMatch `json:"prefix_match,nullable"`
	// Level of sensitivity set for zscore rules.
	ZscoreSensitivity MagicNetworkMonitoringRuleZscoreSensitivity `json:"zscore_sensitivity,nullable"`
	// Target of the zscore rule analysis.
	ZscoreTarget MagicNetworkMonitoringRuleZscoreTarget `json:"zscore_target,nullable"`
	JSON         magicNetworkMonitoringRuleJSON         `json:"-"`
}

// magicNetworkMonitoringRuleJSON contains the JSON metadata for the struct
// [MagicNetworkMonitoringRule]
type magicNetworkMonitoringRuleJSON struct {
	AutomaticAdvertisement apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	Type                   apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	Duration               apijson.Field
	PacketThreshold        apijson.Field
	PrefixMatch            apijson.Field
	ZscoreSensitivity      apijson.Field
	ZscoreTarget           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleJSON) RawJSON() string {
	return r.raw
}

// MNM rule type.
type MagicNetworkMonitoringRuleType string

const (
	MagicNetworkMonitoringRuleTypeThreshold    MagicNetworkMonitoringRuleType = "threshold"
	MagicNetworkMonitoringRuleTypeZscore       MagicNetworkMonitoringRuleType = "zscore"
	MagicNetworkMonitoringRuleTypeAdvancedDDoS MagicNetworkMonitoringRuleType = "advanced_ddos"
)

func (r MagicNetworkMonitoringRuleType) IsKnown() bool {
	switch r {
	case MagicNetworkMonitoringRuleTypeThreshold, MagicNetworkMonitoringRuleTypeZscore, MagicNetworkMonitoringRuleTypeAdvancedDDoS:
		return true
	}
	return false
}

// The amount of time that the rule threshold must be exceeded to send an alert
// notification. The final value must be equivalent to one of the following 8
// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
type MagicNetworkMonitoringRuleDuration string

const (
	MagicNetworkMonitoringRuleDuration1m  MagicNetworkMonitoringRuleDuration = "1m"
	MagicNetworkMonitoringRuleDuration5m  MagicNetworkMonitoringRuleDuration = "5m"
	MagicNetworkMonitoringRuleDuration10m MagicNetworkMonitoringRuleDuration = "10m"
	MagicNetworkMonitoringRuleDuration15m MagicNetworkMonitoringRuleDuration = "15m"
	MagicNetworkMonitoringRuleDuration20m MagicNetworkMonitoringRuleDuration = "20m"
	MagicNetworkMonitoringRuleDuration30m MagicNetworkMonitoringRuleDuration = "30m"
	MagicNetworkMonitoringRuleDuration45m MagicNetworkMonitoringRuleDuration = "45m"
	MagicNetworkMonitoringRuleDuration60m MagicNetworkMonitoringRuleDuration = "60m"
)

func (r MagicNetworkMonitoringRuleDuration) IsKnown() bool {
	switch r {
	case MagicNetworkMonitoringRuleDuration1m, MagicNetworkMonitoringRuleDuration5m, MagicNetworkMonitoringRuleDuration10m, MagicNetworkMonitoringRuleDuration15m, MagicNetworkMonitoringRuleDuration20m, MagicNetworkMonitoringRuleDuration30m, MagicNetworkMonitoringRuleDuration45m, MagicNetworkMonitoringRuleDuration60m:
		return true
	}
	return false
}

// Prefix match type to be applied for a prefix auto advertisement when using an
// advanced_ddos rule.
type MagicNetworkMonitoringRulePrefixMatch string

const (
	MagicNetworkMonitoringRulePrefixMatchExact    MagicNetworkMonitoringRulePrefixMatch = "exact"
	MagicNetworkMonitoringRulePrefixMatchSubnet   MagicNetworkMonitoringRulePrefixMatch = "subnet"
	MagicNetworkMonitoringRulePrefixMatchSupernet MagicNetworkMonitoringRulePrefixMatch = "supernet"
)

func (r MagicNetworkMonitoringRulePrefixMatch) IsKnown() bool {
	switch r {
	case MagicNetworkMonitoringRulePrefixMatchExact, MagicNetworkMonitoringRulePrefixMatchSubnet, MagicNetworkMonitoringRulePrefixMatchSupernet:
		return true
	}
	return false
}

// Level of sensitivity set for zscore rules.
type MagicNetworkMonitoringRuleZscoreSensitivity string

const (
	MagicNetworkMonitoringRuleZscoreSensitivityLow    MagicNetworkMonitoringRuleZscoreSensitivity = "low"
	MagicNetworkMonitoringRuleZscoreSensitivityMedium MagicNetworkMonitoringRuleZscoreSensitivity = "medium"
	MagicNetworkMonitoringRuleZscoreSensitivityHigh   MagicNetworkMonitoringRuleZscoreSensitivity = "high"
)

func (r MagicNetworkMonitoringRuleZscoreSensitivity) IsKnown() bool {
	switch r {
	case MagicNetworkMonitoringRuleZscoreSensitivityLow, MagicNetworkMonitoringRuleZscoreSensitivityMedium, MagicNetworkMonitoringRuleZscoreSensitivityHigh:
		return true
	}
	return false
}

// Target of the zscore rule analysis.
type MagicNetworkMonitoringRuleZscoreTarget string

const (
	MagicNetworkMonitoringRuleZscoreTargetBits    MagicNetworkMonitoringRuleZscoreTarget = "bits"
	MagicNetworkMonitoringRuleZscoreTargetPackets MagicNetworkMonitoringRuleZscoreTarget = "packets"
)

func (r MagicNetworkMonitoringRuleZscoreTarget) IsKnown() bool {
	switch r {
	case MagicNetworkMonitoringRuleZscoreTargetBits, MagicNetworkMonitoringRuleZscoreTargetPackets:
		return true
	}
	return false
}

type RuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleNewParamsDuration] `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name,required"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time that the rule threshold must be exceeded to send an alert
// notification. The final value must be equivalent to one of the following 8
// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
type RuleNewParamsDuration string

const (
	RuleNewParamsDuration1m  RuleNewParamsDuration = "1m"
	RuleNewParamsDuration5m  RuleNewParamsDuration = "5m"
	RuleNewParamsDuration10m RuleNewParamsDuration = "10m"
	RuleNewParamsDuration15m RuleNewParamsDuration = "15m"
	RuleNewParamsDuration20m RuleNewParamsDuration = "20m"
	RuleNewParamsDuration30m RuleNewParamsDuration = "30m"
	RuleNewParamsDuration45m RuleNewParamsDuration = "45m"
	RuleNewParamsDuration60m RuleNewParamsDuration = "60m"
)

func (r RuleNewParamsDuration) IsKnown() bool {
	switch r {
	case RuleNewParamsDuration1m, RuleNewParamsDuration5m, RuleNewParamsDuration10m, RuleNewParamsDuration15m, RuleNewParamsDuration20m, RuleNewParamsDuration30m, RuleNewParamsDuration45m, RuleNewParamsDuration60m:
		return true
	}
	return false
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   MagicNetworkMonitoringRule `json:"result,required,nullable"`
	// Whether the API call was successful
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

// Whether the API call was successful
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

type RuleUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleUpdateParamsDuration] `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name,required"`
	// The id of the rule. Must be unique.
	ID param.Field[string] `json:"id"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time that the rule threshold must be exceeded to send an alert
// notification. The final value must be equivalent to one of the following 8
// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
type RuleUpdateParamsDuration string

const (
	RuleUpdateParamsDuration1m  RuleUpdateParamsDuration = "1m"
	RuleUpdateParamsDuration5m  RuleUpdateParamsDuration = "5m"
	RuleUpdateParamsDuration10m RuleUpdateParamsDuration = "10m"
	RuleUpdateParamsDuration15m RuleUpdateParamsDuration = "15m"
	RuleUpdateParamsDuration20m RuleUpdateParamsDuration = "20m"
	RuleUpdateParamsDuration30m RuleUpdateParamsDuration = "30m"
	RuleUpdateParamsDuration45m RuleUpdateParamsDuration = "45m"
	RuleUpdateParamsDuration60m RuleUpdateParamsDuration = "60m"
)

func (r RuleUpdateParamsDuration) IsKnown() bool {
	switch r {
	case RuleUpdateParamsDuration1m, RuleUpdateParamsDuration5m, RuleUpdateParamsDuration10m, RuleUpdateParamsDuration15m, RuleUpdateParamsDuration20m, RuleUpdateParamsDuration30m, RuleUpdateParamsDuration45m, RuleUpdateParamsDuration60m:
		return true
	}
	return false
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   MagicNetworkMonitoringRule `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RuleDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   MagicNetworkMonitoringRule `json:"result,required,nullable"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
	AccountID param.Field[string] `path:"account_id,required"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleEditParamsDuration] `json:"duration"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time that the rule threshold must be exceeded to send an alert
// notification. The final value must be equivalent to one of the following 8
// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
type RuleEditParamsDuration string

const (
	RuleEditParamsDuration1m  RuleEditParamsDuration = "1m"
	RuleEditParamsDuration5m  RuleEditParamsDuration = "5m"
	RuleEditParamsDuration10m RuleEditParamsDuration = "10m"
	RuleEditParamsDuration15m RuleEditParamsDuration = "15m"
	RuleEditParamsDuration20m RuleEditParamsDuration = "20m"
	RuleEditParamsDuration30m RuleEditParamsDuration = "30m"
	RuleEditParamsDuration45m RuleEditParamsDuration = "45m"
	RuleEditParamsDuration60m RuleEditParamsDuration = "60m"
)

func (r RuleEditParamsDuration) IsKnown() bool {
	switch r {
	case RuleEditParamsDuration1m, RuleEditParamsDuration5m, RuleEditParamsDuration10m, RuleEditParamsDuration15m, RuleEditParamsDuration20m, RuleEditParamsDuration30m, RuleEditParamsDuration45m, RuleEditParamsDuration60m:
		return true
	}
	return false
}

type RuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   MagicNetworkMonitoringRule `json:"result,required,nullable"`
	// Whether the API call was successful
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

// Whether the API call was successful
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

type RuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   MagicNetworkMonitoringRule `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
