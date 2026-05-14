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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DevicePolicyDefaultService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultService] method instead.
type DevicePolicyDefaultService struct {
	Options         []option.RequestOption
	Excludes        *DevicePolicyDefaultExcludeService
	Includes        *DevicePolicyDefaultIncludeService
	FallbackDomains *DevicePolicyDefaultFallbackDomainService
	Certificates    *DevicePolicyDefaultCertificateService
}

// NewDevicePolicyDefaultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultService(opts ...option.RequestOption) (r *DevicePolicyDefaultService) {
	r = &DevicePolicyDefaultService{}
	r.Options = opts
	r.Excludes = NewDevicePolicyDefaultExcludeService(opts...)
	r.Includes = NewDevicePolicyDefaultIncludeService(opts...)
	r.FallbackDomains = NewDevicePolicyDefaultFallbackDomainService(opts...)
	r.Certificates = NewDevicePolicyDefaultCertificateService(opts...)
	return
}

// Updates the default device settings profile for an account.
func (r *DevicePolicyDefaultService) Edit(ctx context.Context, params DevicePolicyDefaultEditParams, opts ...option.RequestOption) (res *DevicePolicyDefaultEditResponse, err error) {
	var env DevicePolicyDefaultEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the default device settings profile for an account.
func (r *DevicePolicyDefaultService) Get(ctx context.Context, query DevicePolicyDefaultGetParams, opts ...option.RequestOption) (res *DevicePolicyDefaultGetResponse, err error) {
	var env DevicePolicyDefaultGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultEditResponse struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy will be applied to matching devices.
	Default bool `json:"default"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel.
	Exclude []SplitTunnelExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool             `json:"exclude_office_ips"`
	FallbackDomains  []FallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string           `json:"gateway_unique_id"`
	// List of routes included in the WARP client's tunnel.
	Include []SplitTunnelInclude `json:"include"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS bool `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport bool                                         `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          DevicePolicyDefaultEditResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string                              `json:"tunnel_protocol"`
	JSON           devicePolicyDefaultEditResponseJSON `json:"-"`
}

// devicePolicyDefaultEditResponseJSON contains the JSON metadata for the struct
// [DevicePolicyDefaultEditResponse]
type devicePolicyDefaultEditResponseJSON struct {
	AllowModeSwitch            apijson.Field
	AllowUpdates               apijson.Field
	AllowedToLeave             apijson.Field
	AutoConnect                apijson.Field
	CaptivePortal              apijson.Field
	Default                    apijson.Field
	DisableAutoFallback        apijson.Field
	Enabled                    apijson.Field
	Exclude                    apijson.Field
	ExcludeOfficeIPs           apijson.Field
	FallbackDomains            apijson.Field
	GatewayUniqueID            apijson.Field
	Include                    apijson.Field
	RegisterInterfaceIPWithDNS apijson.Field
	SccmVpnBoundarySupport     apijson.Field
	ServiceModeV2              apijson.Field
	SupportURL                 apijson.Field
	SwitchLocked               apijson.Field
	TunnelProtocol             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DevicePolicyDefaultEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultEditResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultEditResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                          `json:"port"`
	JSON devicePolicyDefaultEditResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyDefaultEditResponseServiceModeV2JSON contains the JSON metadata for
// the struct [DevicePolicyDefaultEditResponseServiceModeV2]
type devicePolicyDefaultEditResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultEditResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultEditResponseServiceModeV2JSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultGetResponse struct {
	// Whether to allow the user to switch WARP between modes.
	AllowModeSwitch bool `json:"allow_mode_switch"`
	// Whether to receive update notifications when a new version of the client is
	// available.
	AllowUpdates bool `json:"allow_updates"`
	// Whether to allow devices to leave the organization.
	AllowedToLeave bool `json:"allowed_to_leave"`
	// The amount of time in seconds to reconnect after having been disabled.
	AutoConnect float64 `json:"auto_connect"`
	// Turn on the captive portal after the specified amount of time.
	CaptivePortal float64 `json:"captive_portal"`
	// Whether the policy will be applied to matching devices.
	Default bool `json:"default"`
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback bool `json:"disable_auto_fallback"`
	// Whether the policy will be applied to matching devices.
	Enabled bool `json:"enabled"`
	// List of routes excluded in the WARP client's tunnel.
	Exclude []SplitTunnelExclude `json:"exclude"`
	// Whether to add Microsoft IPs to Split Tunnel exclusions.
	ExcludeOfficeIPs bool             `json:"exclude_office_ips"`
	FallbackDomains  []FallbackDomain `json:"fallback_domains"`
	GatewayUniqueID  string           `json:"gateway_unique_id"`
	// List of routes included in the WARP client's tunnel.
	Include []SplitTunnelInclude `json:"include"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS bool `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport bool                                        `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          DevicePolicyDefaultGetResponseServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string                             `json:"tunnel_protocol"`
	JSON           devicePolicyDefaultGetResponseJSON `json:"-"`
}

// devicePolicyDefaultGetResponseJSON contains the JSON metadata for the struct
// [DevicePolicyDefaultGetResponse]
type devicePolicyDefaultGetResponseJSON struct {
	AllowModeSwitch            apijson.Field
	AllowUpdates               apijson.Field
	AllowedToLeave             apijson.Field
	AutoConnect                apijson.Field
	CaptivePortal              apijson.Field
	Default                    apijson.Field
	DisableAutoFallback        apijson.Field
	Enabled                    apijson.Field
	Exclude                    apijson.Field
	ExcludeOfficeIPs           apijson.Field
	FallbackDomains            apijson.Field
	GatewayUniqueID            apijson.Field
	Include                    apijson.Field
	RegisterInterfaceIPWithDNS apijson.Field
	SccmVpnBoundarySupport     apijson.Field
	ServiceModeV2              apijson.Field
	SupportURL                 apijson.Field
	SwitchLocked               apijson.Field
	TunnelProtocol             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DevicePolicyDefaultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultGetResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultGetResponseServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                                         `json:"port"`
	JSON devicePolicyDefaultGetResponseServiceModeV2JSON `json:"-"`
}

// devicePolicyDefaultGetResponseServiceModeV2JSON contains the JSON metadata for
// the struct [DevicePolicyDefaultGetResponseServiceModeV2]
type devicePolicyDefaultGetResponseServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultGetResponseServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultGetResponseServiceModeV2JSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultEditParams struct {
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
	// If the `dns_server` field of a fallback domain is not present, the client will
	// fall back to a best guess of the default/system DNS resolvers unless this policy
	// option is set to `true`.
	DisableAutoFallback param.Field[bool] `json:"disable_auto_fallback"`
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
	SccmVpnBoundarySupport param.Field[bool]                                       `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          param.Field[DevicePolicyDefaultEditParamsServiceModeV2] `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL param.Field[string] `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked param.Field[bool] `json:"switch_locked"`
	// Determines which tunnel protocol to use.
	TunnelProtocol param.Field[string] `json:"tunnel_protocol"`
}

func (r DevicePolicyDefaultEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDefaultEditParamsServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode param.Field[string] `json:"mode"`
	// The port number when used with proxy mode.
	Port param.Field[float64] `json:"port"`
}

func (r DevicePolicyDefaultEditParamsServiceModeV2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyDefaultEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   DevicePolicyDefaultEditResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyDefaultEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyDefaultEditResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyDefaultEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyDefaultEditResponseEnvelope]
type devicePolicyDefaultEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultEditResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultEditResponseEnvelopeSuccessTrue DevicePolicyDefaultEditResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   DevicePolicyDefaultGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyDefaultGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyDefaultGetResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyDefaultGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyDefaultGetResponseEnvelope]
type devicePolicyDefaultGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultGetResponseEnvelopeSuccessTrue DevicePolicyDefaultGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
