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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DevicePolicyCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomService] method instead.
type DevicePolicyCustomService struct {
	Options         []option.RequestOption
	Excludes        *DevicePolicyCustomExcludeService
	Includes        *DevicePolicyCustomIncludeService
	FallbackDomains *DevicePolicyCustomFallbackDomainService
}

// NewDevicePolicyCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomService(opts ...option.RequestOption) (r *DevicePolicyCustomService) {
	r = &DevicePolicyCustomService{}
	r.Options = opts
	r.Excludes = NewDevicePolicyCustomExcludeService(opts...)
	r.Includes = NewDevicePolicyCustomIncludeService(opts...)
	r.FallbackDomains = NewDevicePolicyCustomFallbackDomainService(opts...)
	return
}

// Creates a device settings profile to be applied to certain devices matching the
// criteria.
func (r *DevicePolicyCustomService) New(ctx context.Context, params DevicePolicyCustomNewParams, opts ...option.RequestOption) (res *SettingsPolicy, err error) {
	var env DevicePolicyCustomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of the device settings profiles for an account.
func (r *DevicePolicyCustomService) List(ctx context.Context, query DevicePolicyCustomListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SettingsPolicy], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policies", query.AccountID)
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

// Fetches a list of the device settings profiles for an account.
func (r *DevicePolicyCustomService) ListAutoPaging(ctx context.Context, query DevicePolicyCustomListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SettingsPolicy] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *DevicePolicyCustomService) Delete(ctx context.Context, policyID string, body DevicePolicyCustomDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[SettingsPolicy], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", body.AccountID, policyID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes a device settings profile and fetches a list of the remaining profiles
// for an account.
func (r *DevicePolicyCustomService) DeleteAutoPaging(ctx context.Context, policyID string, body DevicePolicyCustomDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SettingsPolicy] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, policyID, body, opts...))
}

// Updates a configured device settings profile.
func (r *DevicePolicyCustomService) Edit(ctx context.Context, policyID string, params DevicePolicyCustomEditParams, opts ...option.RequestOption) (res *SettingsPolicy, err error) {
	var env DevicePolicyCustomEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a device settings profile by ID.
func (r *DevicePolicyCustomService) Get(ctx context.Context, policyID string, query DevicePolicyCustomGetParams, opts ...option.RequestOption) (res *SettingsPolicy, err error) {
	var env DevicePolicyCustomGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyCustomNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match param.Field[string] `json:"match,required"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name,required"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence param.Field[float64] `json:"precedence,required"`
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Exclude param.Field[[]SplitTunnelExcludeUnionParam] `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// List of routes included in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Include param.Field[[]SplitTunnelIncludeUnionParam] `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LANAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LANAllowSubnetSize param.Field[float64] `json:"lan_allow_subnet_size"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS param.Field[bool] `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport param.Field[bool]                                     `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[DevicePolicyCustomNewParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r DevicePolicyCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyCustomNewParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyCustomNewParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SettingsPolicy        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCustomNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyCustomNewResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyCustomNewResponseEnvelope]
type devicePolicyCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomNewResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomNewResponseEnvelopeSuccessTrue DevicePolicyCustomNewResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch param.Field[bool] `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates param.Field[bool] `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave param.Field[bool] `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect param.Field[float64] `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal param.Field[float64] `json:"captive_portal"`
	// A description of the policy.
	Description param.Field[string] `json:"description"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Exclude param.Field[[]SplitTunnelExcludeUnionParam] `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs param.Field[bool] `json:"exclude_office_ips"`
	// List of routes included in the WARP client's tunnel. Both 'exclude' and
	// 'include' cannot be set in the same request.
	Include param.Field[[]SplitTunnelIncludeUnionParam] `json:"include"`
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LANAllowMinutes param.Field[float64] `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LANAllowSubnetSize param.Field[float64] `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match param.Field[string] `json:"match"`
	// The name of the device settings profile.
	Name param.Field[string] `json:"name"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence param.Field[float64] `json:"precedence"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS param.Field[bool] `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport param.Field[bool]                                      `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[DevicePolicyCustomEditParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r DevicePolicyCustomEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyCustomEditParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyCustomEditParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyCustomEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SettingsPolicy        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCustomEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyCustomEditResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyCustomEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyCustomEditResponseEnvelope]
type devicePolicyCustomEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomEditResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomEditResponseEnvelopeSuccessTrue DevicePolicyCustomEditResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SettingsPolicy        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCustomGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyCustomGetResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyCustomGetResponseEnvelope]
type devicePolicyCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomGetResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomGetResponseEnvelopeSuccessTrue DevicePolicyCustomGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
