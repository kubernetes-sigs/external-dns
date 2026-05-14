// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// LockdownService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockdownService] method instead.
type LockdownService struct {
	Options []option.RequestOption
}

// NewLockdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockdownService(opts ...option.RequestOption) (r *LockdownService) {
	r = &LockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *LockdownService) New(ctx context.Context, params LockdownNewParams, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *LockdownService) Update(ctx context.Context, lockDownsID string, params LockdownUpdateParams, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", params.ZoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) List(ctx context.Context, params LockdownListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Lockdown], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", params.ZoneID)
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

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) ListAutoPaging(ctx context.Context, params LockdownListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Lockdown] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *LockdownService) Delete(ctx context.Context, lockDownsID string, body LockdownDeleteParams, opts ...option.RequestOption) (res *LockdownDeleteResponse, err error) {
	var env LockdownDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", body.ZoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *LockdownService) Get(ctx context.Context, lockDownsID string, query LockdownGetParams, opts ...option.RequestOption) (res *Lockdown, err error) {
	var env LockdownGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if lockDownsID == "" {
		err = errors.New("missing required lock_downs_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", query.ZoneID, lockDownsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Configuration []ConfigurationItem

type ConfigurationItem struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ConfigurationItemTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                `json:"value"`
	JSON  configurationItemJSON `json:"-"`
	union ConfigurationItemUnion
}

// configurationItemJSON contains the JSON metadata for the struct
// [ConfigurationItem]
type configurationItemJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationItemJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationItem) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationItemUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [LockdownIPConfiguration],
// [LockdownCIDRConfiguration].
func (r ConfigurationItem) AsUnion() ConfigurationItemUnion {
	return r.union
}

// Union satisfied by [LockdownIPConfiguration] or [LockdownCIDRConfiguration].
type ConfigurationItemUnion interface {
	implementsConfigurationItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownCIDRConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ConfigurationItemTarget string

const (
	ConfigurationItemTargetIP      ConfigurationItemTarget = "ip"
	ConfigurationItemTargetIPRange ConfigurationItemTarget = "ip_range"
)

func (r ConfigurationItemTarget) IsKnown() bool {
	switch r {
	case ConfigurationItemTargetIP, ConfigurationItemTargetIPRange:
		return true
	}
	return false
}

type ConfigurationParam []ConfigurationItemUnionParam

type ConfigurationItemParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target param.Field[ConfigurationItemTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r ConfigurationItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationItemParam) implementsConfigurationItemUnionParam() {}

// Satisfied by [firewall.LockdownIPConfigurationParam],
// [firewall.LockdownCIDRConfigurationParam], [ConfigurationItemParam].
type ConfigurationItemUnionParam interface {
	implementsConfigurationItemUnionParam()
}

type Lockdown struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations Configuration `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []LockdownURL `json:"urls,required"`
	JSON lockdownJSON  `json:"-"`
}

// lockdownJSON contains the JSON metadata for the struct [Lockdown]
type lockdownJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Lockdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownJSON) RawJSON() string {
	return r.raw
}

type LockdownCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                        `json:"value"`
	JSON  lockdownCIDRConfigurationJSON `json:"-"`
}

// lockdownCIDRConfigurationJSON contains the JSON metadata for the struct
// [LockdownCIDRConfiguration]
type lockdownCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownCIDRConfiguration) implementsConfigurationItem() {}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownCIDRConfigurationTarget string

const (
	LockdownCIDRConfigurationTargetIPRange LockdownCIDRConfigurationTarget = "ip_range"
)

func (r LockdownCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type LockdownCIDRConfigurationParam struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target param.Field[LockdownCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value param.Field[string] `json:"value"`
}

func (r LockdownCIDRConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LockdownCIDRConfigurationParam) implementsConfigurationItemUnionParam() {}

type LockdownIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                      `json:"value"`
	JSON  lockdownIPConfigurationJSON `json:"-"`
}

// lockdownIPConfigurationJSON contains the JSON metadata for the struct
// [LockdownIPConfiguration]
type lockdownIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownIPConfiguration) implementsConfigurationItem() {}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownIPConfigurationTarget string

const (
	LockdownIPConfigurationTargetIP LockdownIPConfigurationTarget = "ip"
)

func (r LockdownIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownIPConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target param.Field[LockdownIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r LockdownIPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LockdownIPConfigurationParam) implementsConfigurationItemUnionParam() {}

type LockdownURL = string

type LockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                     `json:"id"`
	JSON lockdownDeleteResponseJSON `json:"-"`
}

// lockdownDeleteResponseJSON contains the JSON metadata for the struct
// [LockdownDeleteResponse]
type lockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LockdownNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations param.Field[ConfigurationParam] `json:"configurations,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `json:"priority"`
}

func (r LockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LockdownNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Defines whether the API call was successful.
	Success LockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownNewResponseEnvelopeJSON    `json:"-"`
}

// lockdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownNewResponseEnvelope]
type lockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type LockdownNewResponseEnvelopeSuccess bool

const (
	LockdownNewResponseEnvelopeSuccessTrue LockdownNewResponseEnvelopeSuccess = true
)

func (r LockdownNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations param.Field[ConfigurationParam] `json:"configurations,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls,required"`
}

func (r LockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LockdownUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Defines whether the API call was successful.
	Success LockdownUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownUpdateResponseEnvelopeJSON    `json:"-"`
}

// lockdownUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownUpdateResponseEnvelope]
type lockdownUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type LockdownUpdateResponseEnvelopeSuccess bool

const (
	LockdownUpdateResponseEnvelopeSuccessTrue LockdownUpdateResponseEnvelopeSuccess = true
)

func (r LockdownUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The timestamp of when the rule was created.
	CreatedOn param.Field[time.Time] `query:"created_on" format:"date-time"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// The timestamp of when the rule was last modified.
	ModifiedOn param.Field[time.Time] `query:"modified_on" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	URISearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [LockdownListParams]'s query parameters as `url.Values`.
func (r LockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LockdownDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LockdownDeleteResponseEnvelope struct {
	Result LockdownDeleteResponse             `json:"result"`
	JSON   lockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// lockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownDeleteResponseEnvelope]
type lockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LockdownGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LockdownGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Lockdown              `json:"result,required"`
	// Defines whether the API call was successful.
	Success LockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownGetResponseEnvelopeJSON    `json:"-"`
}

// lockdownGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownGetResponseEnvelope]
type lockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type LockdownGetResponseEnvelopeSuccess bool

const (
	LockdownGetResponseEnvelopeSuccessTrue LockdownGetResponseEnvelopeSuccess = true
)

func (r LockdownGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
