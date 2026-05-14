// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// UARuleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUARuleService] method instead.
type UARuleService struct {
	Options []option.RequestOption
}

// NewUARuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUARuleService(opts ...option.RequestOption) (r *UARuleService) {
	r = &UARuleService{}
	r.Options = opts
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *UARuleService) New(ctx context.Context, params UARuleNewParams, opts ...option.RequestOption) (res *UARuleNewResponse, err error) {
	var env UARuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing User Agent Blocking rule.
func (r *UARuleService) Update(ctx context.Context, uaRuleID string, params UARuleUpdateParams, opts ...option.RequestOption) (res *UARuleUpdateResponse, err error) {
	var env UARuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", params.ZoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *UARuleService) List(ctx context.Context, params UARuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UARuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", params.ZoneID)
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

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *UARuleService) ListAutoPaging(ctx context.Context, params UARuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UARuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing User Agent Blocking rule.
func (r *UARuleService) Delete(ctx context.Context, uaRuleID string, body UARuleDeleteParams, opts ...option.RequestOption) (res *UARuleDeleteResponse, err error) {
	var env UARuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", body.ZoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *UARuleService) Get(ctx context.Context, uaRuleID string, query UARuleGetParams, opts ...option.RequestOption) (res *UARuleGetResponse, err error) {
	var env UARuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", query.ZoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UARuleNewResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UARuleNewResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UARuleNewResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                  `json:"paused"`
	JSON   uaRuleNewResponseJSON `json:"-"`
}

// uaRuleNewResponseJSON contains the JSON metadata for the struct
// [UARuleNewResponse]
type uaRuleNewResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UARuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UARuleNewResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                             `json:"value"`
	JSON  uaRuleNewResponseConfigurationJSON `json:"-"`
}

// uaRuleNewResponseConfigurationJSON contains the JSON metadata for the struct
// [UARuleNewResponseConfiguration]
type uaRuleNewResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleNewResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleNewResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UARuleNewResponseMode string

const (
	UARuleNewResponseModeBlock            UARuleNewResponseMode = "block"
	UARuleNewResponseModeChallenge        UARuleNewResponseMode = "challenge"
	UARuleNewResponseModeJSChallenge      UARuleNewResponseMode = "js_challenge"
	UARuleNewResponseModeManagedChallenge UARuleNewResponseMode = "managed_challenge"
)

func (r UARuleNewResponseMode) IsKnown() bool {
	switch r {
	case UARuleNewResponseModeBlock, UARuleNewResponseModeChallenge, UARuleNewResponseModeJSChallenge, UARuleNewResponseModeManagedChallenge:
		return true
	}
	return false
}

type UARuleUpdateResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UARuleUpdateResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UARuleUpdateResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                     `json:"paused"`
	JSON   uaRuleUpdateResponseJSON `json:"-"`
}

// uaRuleUpdateResponseJSON contains the JSON metadata for the struct
// [UARuleUpdateResponse]
type uaRuleUpdateResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UARuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UARuleUpdateResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                `json:"value"`
	JSON  uaRuleUpdateResponseConfigurationJSON `json:"-"`
}

// uaRuleUpdateResponseConfigurationJSON contains the JSON metadata for the struct
// [UARuleUpdateResponseConfiguration]
type uaRuleUpdateResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleUpdateResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleUpdateResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UARuleUpdateResponseMode string

const (
	UARuleUpdateResponseModeBlock            UARuleUpdateResponseMode = "block"
	UARuleUpdateResponseModeChallenge        UARuleUpdateResponseMode = "challenge"
	UARuleUpdateResponseModeJSChallenge      UARuleUpdateResponseMode = "js_challenge"
	UARuleUpdateResponseModeManagedChallenge UARuleUpdateResponseMode = "managed_challenge"
)

func (r UARuleUpdateResponseMode) IsKnown() bool {
	switch r {
	case UARuleUpdateResponseModeBlock, UARuleUpdateResponseModeChallenge, UARuleUpdateResponseModeJSChallenge, UARuleUpdateResponseModeManagedChallenge:
		return true
	}
	return false
}

type UARuleListResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UARuleListResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UARuleListResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                   `json:"paused"`
	JSON   uaRuleListResponseJSON `json:"-"`
}

// uaRuleListResponseJSON contains the JSON metadata for the struct
// [UARuleListResponse]
type uaRuleListResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UARuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UARuleListResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                              `json:"value"`
	JSON  uaRuleListResponseConfigurationJSON `json:"-"`
}

// uaRuleListResponseConfigurationJSON contains the JSON metadata for the struct
// [UARuleListResponseConfiguration]
type uaRuleListResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UARuleListResponseMode string

const (
	UARuleListResponseModeBlock            UARuleListResponseMode = "block"
	UARuleListResponseModeChallenge        UARuleListResponseMode = "challenge"
	UARuleListResponseModeJSChallenge      UARuleListResponseMode = "js_challenge"
	UARuleListResponseModeManagedChallenge UARuleListResponseMode = "managed_challenge"
)

func (r UARuleListResponseMode) IsKnown() bool {
	switch r {
	case UARuleListResponseModeBlock, UARuleListResponseModeChallenge, UARuleListResponseModeJSChallenge, UARuleListResponseModeManagedChallenge:
		return true
	}
	return false
}

type UARuleDeleteResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UARuleDeleteResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UARuleDeleteResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                     `json:"paused"`
	JSON   uaRuleDeleteResponseJSON `json:"-"`
}

// uaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [UARuleDeleteResponse]
type uaRuleDeleteResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UARuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UARuleDeleteResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                `json:"value"`
	JSON  uaRuleDeleteResponseConfigurationJSON `json:"-"`
}

// uaRuleDeleteResponseConfigurationJSON contains the JSON metadata for the struct
// [UARuleDeleteResponseConfiguration]
type uaRuleDeleteResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleDeleteResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleDeleteResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UARuleDeleteResponseMode string

const (
	UARuleDeleteResponseModeBlock            UARuleDeleteResponseMode = "block"
	UARuleDeleteResponseModeChallenge        UARuleDeleteResponseMode = "challenge"
	UARuleDeleteResponseModeJSChallenge      UARuleDeleteResponseMode = "js_challenge"
	UARuleDeleteResponseModeManagedChallenge UARuleDeleteResponseMode = "managed_challenge"
)

func (r UARuleDeleteResponseMode) IsKnown() bool {
	switch r {
	case UARuleDeleteResponseModeBlock, UARuleDeleteResponseModeChallenge, UARuleDeleteResponseModeJSChallenge, UARuleDeleteResponseModeManagedChallenge:
		return true
	}
	return false
}

type UARuleGetResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UARuleGetResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UARuleGetResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                  `json:"paused"`
	JSON   uaRuleGetResponseJSON `json:"-"`
}

// uaRuleGetResponseJSON contains the JSON metadata for the struct
// [UARuleGetResponse]
type uaRuleGetResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UARuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UARuleGetResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                             `json:"value"`
	JSON  uaRuleGetResponseConfigurationJSON `json:"-"`
}

// uaRuleGetResponseConfigurationJSON contains the JSON metadata for the struct
// [UARuleGetResponseConfiguration]
type uaRuleGetResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleGetResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleGetResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UARuleGetResponseMode string

const (
	UARuleGetResponseModeBlock            UARuleGetResponseMode = "block"
	UARuleGetResponseModeChallenge        UARuleGetResponseMode = "challenge"
	UARuleGetResponseModeJSChallenge      UARuleGetResponseMode = "js_challenge"
	UARuleGetResponseModeManagedChallenge UARuleGetResponseMode = "managed_challenge"
)

func (r UARuleGetResponseMode) IsKnown() bool {
	switch r {
	case UARuleGetResponseModeBlock, UARuleGetResponseModeChallenge, UARuleGetResponseModeJSChallenge, UARuleGetResponseModeManagedChallenge:
		return true
	}
	return false
}

type UARuleNewParams struct {
	// Defines an identifier.
	ZoneID        param.Field[string]                       `path:"zone_id,required"`
	Configuration param.Field[UARuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[UARuleNewParamsMode] `json:"mode,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
}

func (r UARuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UARuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ua` when specifying a user
	// agent in the rule.
	Target param.Field[UARuleNewParamsConfigurationTarget] `json:"target"`
	// the user agent to exactly match
	Value param.Field[string] `json:"value"`
}

func (r UARuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration target. You must set the target to `ua` when specifying a user
// agent in the rule.
type UARuleNewParamsConfigurationTarget string

const (
	UARuleNewParamsConfigurationTargetUA UARuleNewParamsConfigurationTarget = "ua"
)

func (r UARuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case UARuleNewParamsConfigurationTargetUA:
		return true
	}
	return false
}

// The action to apply to a matched request.
type UARuleNewParamsMode string

const (
	UARuleNewParamsModeBlock            UARuleNewParamsMode = "block"
	UARuleNewParamsModeChallenge        UARuleNewParamsMode = "challenge"
	UARuleNewParamsModeWhitelist        UARuleNewParamsMode = "whitelist"
	UARuleNewParamsModeJSChallenge      UARuleNewParamsMode = "js_challenge"
	UARuleNewParamsModeManagedChallenge UARuleNewParamsMode = "managed_challenge"
)

func (r UARuleNewParamsMode) IsKnown() bool {
	switch r {
	case UARuleNewParamsModeBlock, UARuleNewParamsModeChallenge, UARuleNewParamsModeWhitelist, UARuleNewParamsModeJSChallenge, UARuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type UARuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UARuleNewResponse     `json:"result,required"`
	// Defines whether the API call was successful.
	Success UARuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleNewResponseEnvelopeJSON    `json:"-"`
}

// uaRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleNewResponseEnvelope]
type uaRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type UARuleNewResponseEnvelopeSuccess bool

const (
	UARuleNewResponseEnvelopeSuccessTrue UARuleNewResponseEnvelopeSuccess = true
)

func (r UARuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The rule configuration.
	Configuration param.Field[UARuleUpdateParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[UARuleUpdateParamsMode] `json:"mode,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
}

func (r UARuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type UARuleUpdateParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[UARuleUpdateParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r UARuleUpdateParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UARuleUpdateParamsConfiguration) implementsUARuleUpdateParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by [firewall.AccessRuleIPConfigurationParam],
// [firewall.IPV6ConfigurationParam], [firewall.AccessRuleCIDRConfigurationParam],
// [firewall.ASNConfigurationParam], [firewall.CountryConfigurationParam],
// [UARuleUpdateParamsConfiguration].
type UARuleUpdateParamsConfigurationUnion interface {
	implementsUARuleUpdateParamsConfigurationUnion()
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type UARuleUpdateParamsConfigurationTarget string

const (
	UARuleUpdateParamsConfigurationTargetIP      UARuleUpdateParamsConfigurationTarget = "ip"
	UARuleUpdateParamsConfigurationTargetIp6     UARuleUpdateParamsConfigurationTarget = "ip6"
	UARuleUpdateParamsConfigurationTargetIPRange UARuleUpdateParamsConfigurationTarget = "ip_range"
	UARuleUpdateParamsConfigurationTargetASN     UARuleUpdateParamsConfigurationTarget = "asn"
	UARuleUpdateParamsConfigurationTargetCountry UARuleUpdateParamsConfigurationTarget = "country"
)

func (r UARuleUpdateParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case UARuleUpdateParamsConfigurationTargetIP, UARuleUpdateParamsConfigurationTargetIp6, UARuleUpdateParamsConfigurationTargetIPRange, UARuleUpdateParamsConfigurationTargetASN, UARuleUpdateParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type UARuleUpdateParamsMode string

const (
	UARuleUpdateParamsModeBlock            UARuleUpdateParamsMode = "block"
	UARuleUpdateParamsModeChallenge        UARuleUpdateParamsMode = "challenge"
	UARuleUpdateParamsModeWhitelist        UARuleUpdateParamsMode = "whitelist"
	UARuleUpdateParamsModeJSChallenge      UARuleUpdateParamsMode = "js_challenge"
	UARuleUpdateParamsModeManagedChallenge UARuleUpdateParamsMode = "managed_challenge"
)

func (r UARuleUpdateParamsMode) IsKnown() bool {
	switch r {
	case UARuleUpdateParamsModeBlock, UARuleUpdateParamsModeChallenge, UARuleUpdateParamsModeWhitelist, UARuleUpdateParamsModeJSChallenge, UARuleUpdateParamsModeManagedChallenge:
		return true
	}
	return false
}

type UARuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UARuleUpdateResponse  `json:"result,required"`
	// Defines whether the API call was successful.
	Success UARuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// uaRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleUpdateResponseEnvelope]
type uaRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type UARuleUpdateResponseEnvelopeSuccess bool

const (
	UARuleUpdateResponseEnvelopeSuccessTrue UARuleUpdateResponseEnvelopeSuccess = true
)

func (r UARuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// A string to search for in the user agent values of existing rules.
	UserAgent param.Field[string] `query:"user_agent"`
}

// URLQuery serializes [UARuleListParams]'s query parameters as `url.Values`.
func (r UARuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UARuleDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UARuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UARuleDeleteResponse  `json:"result,required"`
	// Defines whether the API call was successful.
	Success UARuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// uaRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleDeleteResponseEnvelope]
type uaRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type UARuleDeleteResponseEnvelopeSuccess bool

const (
	UARuleDeleteResponseEnvelopeSuccessTrue UARuleDeleteResponseEnvelopeSuccess = true
)

func (r UARuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UARuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UARuleGetResponse     `json:"result,required"`
	// Defines whether the API call was successful.
	Success UARuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleGetResponseEnvelopeJSON    `json:"-"`
}

// uaRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleGetResponseEnvelope]
type uaRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type UARuleGetResponseEnvelopeSuccess bool

const (
	UARuleGetResponseEnvelopeSuccessTrue UARuleGetResponseEnvelopeSuccess = true
)

func (r UARuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
