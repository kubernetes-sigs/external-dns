// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// DevicePolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyService] method instead.
type DevicePolicyService struct {
	Options []option.RequestOption
	Default *DevicePolicyDefaultService
	Custom  *DevicePolicyCustomService
}

// NewDevicePolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePolicyService(opts ...option.RequestOption) (r *DevicePolicyService) {
	r = &DevicePolicyService{}
	r.Options = opts
	r.Default = NewDevicePolicyDefaultService(opts...)
	r.Custom = NewDevicePolicyCustomService(opts...)
	return
}

type DevicePolicyCertificates struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled bool                         `json:"enabled,required"`
	JSON    devicePolicyCertificatesJSON `json:"-"`
}

// devicePolicyCertificatesJSON contains the JSON metadata for the struct
// [DevicePolicyCertificates]
type devicePolicyCertificatesJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificatesJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificatesParam struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r DevicePolicyCertificatesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []string           `json:"dns_server"`
	JSON      fallbackDomainJSON `json:"-"`
}

// fallbackDomainJSON contains the JSON metadata for the struct [FallbackDomain]
type fallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainParam struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]string] `json:"dns_server"`
}

func (r FallbackDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingsPolicy struct {
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
	// Whether the policy is the default policy for an account.
	Default bool `json:"default"`
	// A description of the policy.
	Description string `json:"description"`
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
	// The amount of time in minutes a user is allowed access to their LAN. A value of
	// 0 will allow LAN access until the next WARP reconnection, such as a reboot or a
	// laptop waking from sleep. Note that this field is omitted from the response if
	// null or unset.
	LANAllowMinutes float64 `json:"lan_allow_minutes"`
	// The size of the subnet for the local access network. Note that this field is
	// omitted from the response if null or unset.
	LANAllowSubnetSize float64 `json:"lan_allow_subnet_size"`
	// The wirefilter expression to match devices. Available values: "identity.email",
	// "identity.groups.id", "identity.groups.name", "identity.groups.email",
	// "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name",
	// "os.version".
	Match string `json:"match"`
	// The name of the device settings profile.
	Name     string `json:"name"`
	PolicyID string `json:"policy_id"`
	// The precedence of the policy. Lower values indicate higher precedence. Policies
	// will be evaluated in ascending order of this field.
	Precedence float64 `json:"precedence"`
	// Determines if the operating system will register WARP's local interface IP with
	// your on-premises DNS server.
	RegisterInterfaceIPWithDNS bool `json:"register_interface_ip_with_dns"`
	// Determines whether the WARP client indicates to SCCM that it is inside a VPN
	// boundary. (Windows only).
	SccmVpnBoundarySupport bool                        `json:"sccm_vpn_boundary_support"`
	ServiceModeV2          SettingsPolicyServiceModeV2 `json:"service_mode_v2"`
	// The URL to launch when the Send Feedback button is clicked.
	SupportURL string `json:"support_url"`
	// Whether to allow the user to turn off the WARP switch and disconnect the client.
	SwitchLocked bool                       `json:"switch_locked"`
	TargetTests  []SettingsPolicyTargetTest `json:"target_tests"`
	// Determines which tunnel protocol to use.
	TunnelProtocol string             `json:"tunnel_protocol"`
	JSON           settingsPolicyJSON `json:"-"`
}

// settingsPolicyJSON contains the JSON metadata for the struct [SettingsPolicy]
type settingsPolicyJSON struct {
	AllowModeSwitch            apijson.Field
	AllowUpdates               apijson.Field
	AllowedToLeave             apijson.Field
	AutoConnect                apijson.Field
	CaptivePortal              apijson.Field
	Default                    apijson.Field
	Description                apijson.Field
	DisableAutoFallback        apijson.Field
	Enabled                    apijson.Field
	Exclude                    apijson.Field
	ExcludeOfficeIPs           apijson.Field
	FallbackDomains            apijson.Field
	GatewayUniqueID            apijson.Field
	Include                    apijson.Field
	LANAllowMinutes            apijson.Field
	LANAllowSubnetSize         apijson.Field
	Match                      apijson.Field
	Name                       apijson.Field
	PolicyID                   apijson.Field
	Precedence                 apijson.Field
	RegisterInterfaceIPWithDNS apijson.Field
	SccmVpnBoundarySupport     apijson.Field
	ServiceModeV2              apijson.Field
	SupportURL                 apijson.Field
	SwitchLocked               apijson.Field
	TargetTests                apijson.Field
	TunnelProtocol             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SettingsPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsPolicyJSON) RawJSON() string {
	return r.raw
}

type SettingsPolicyServiceModeV2 struct {
	// The mode to run the WARP client under.
	Mode string `json:"mode"`
	// The port number when used with proxy mode.
	Port float64                         `json:"port"`
	JSON settingsPolicyServiceModeV2JSON `json:"-"`
}

// settingsPolicyServiceModeV2JSON contains the JSON metadata for the struct
// [SettingsPolicyServiceModeV2]
type settingsPolicyServiceModeV2JSON struct {
	Mode        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsPolicyServiceModeV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsPolicyServiceModeV2JSON) RawJSON() string {
	return r.raw
}

type SettingsPolicyTargetTest struct {
	// The id of the DEX test targeting this policy.
	ID string `json:"id"`
	// The name of the DEX test targeting this policy.
	Name string                       `json:"name"`
	JSON settingsPolicyTargetTestJSON `json:"-"`
}

// settingsPolicyTargetTestJSON contains the JSON metadata for the struct
// [SettingsPolicyTargetTest]
type settingsPolicyTargetTestJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsPolicyTargetTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsPolicyTargetTestJSON) RawJSON() string {
	return r.raw
}

type SplitTunnelExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host  string                 `json:"host"`
	JSON  splitTunnelExcludeJSON `json:"-"`
	union SplitTunnelExcludeUnion
}

// splitTunnelExcludeJSON contains the JSON metadata for the struct
// [SplitTunnelExclude]
type splitTunnelExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r splitTunnelExcludeJSON) RawJSON() string {
	return r.raw
}

func (r *SplitTunnelExclude) UnmarshalJSON(data []byte) (err error) {
	*r = SplitTunnelExclude{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplitTunnelExcludeUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress],
// [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost].
func (r SplitTunnelExclude) AsUnion() SplitTunnelExcludeUnion {
	return r.union
}

// Union satisfied by [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress]
// or [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost].
type SplitTunnelExcludeUnion interface {
	implementsSplitTunnelExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplitTunnelExcludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost{}),
		},
	)
}

type SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                          `json:"description"`
	JSON        splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressJSON `json:"-"`
}

// splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressJSON contains the
// JSON metadata for the struct
// [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress]
type splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressJSON struct {
	Address     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddress) implementsSplitTunnelExclude() {}

type SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost struct {
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                       `json:"description"`
	JSON        splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostJSON `json:"-"`
}

// splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostJSON contains the JSON
// metadata for the struct
// [SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost]
type splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostJSON struct {
	Host        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHost) implementsSplitTunnelExclude() {}

type SplitTunnelExcludeParam struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelExcludeParam) implementsSplitTunnelExcludeUnionParam() {}

// Satisfied by
// [zero_trust.SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressParam],
// [zero_trust.SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostParam],
// [SplitTunnelExcludeParam].
type SplitTunnelExcludeUnionParam interface {
	implementsSplitTunnelExcludeUnionParam()
}

type SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressParam struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithAddressParam) implementsSplitTunnelExcludeUnionParam() {
}

type SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostParam struct {
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelExcludeTeamsDevicesExcludeSplitTunnelWithHostParam) implementsSplitTunnelExcludeUnionParam() {
}

type SplitTunnelInclude struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description"`
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host  string                 `json:"host"`
	JSON  splitTunnelIncludeJSON `json:"-"`
	union SplitTunnelIncludeUnion
}

// splitTunnelIncludeJSON contains the JSON metadata for the struct
// [SplitTunnelInclude]
type splitTunnelIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r splitTunnelIncludeJSON) RawJSON() string {
	return r.raw
}

func (r *SplitTunnelInclude) UnmarshalJSON(data []byte) (err error) {
	*r = SplitTunnelInclude{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplitTunnelIncludeUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress],
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost].
func (r SplitTunnelInclude) AsUnion() SplitTunnelIncludeUnion {
	return r.union
}

// Union satisfied by [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress]
// or [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost].
type SplitTunnelIncludeUnion interface {
	implementsSplitTunnelInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplitTunnelIncludeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost{}),
		},
	)
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                          `json:"description"`
	JSON        splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON `json:"-"`
}

// splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON contains the
// JSON metadata for the struct
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress]
type splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON struct {
	Address     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddress) implementsSplitTunnelInclude() {}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost struct {
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host string `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string                                                       `json:"description"`
	JSON        splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON `json:"-"`
}

// splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON contains the JSON
// metadata for the struct
// [SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost]
type splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON struct {
	Host        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostJSON) RawJSON() string {
	return r.raw
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHost) implementsSplitTunnelInclude() {}

type SplitTunnelIncludeParam struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelIncludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeParam) implementsSplitTunnelIncludeUnionParam() {}

// Satisfied by
// [zero_trust.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam],
// [zero_trust.SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam],
// [SplitTunnelIncludeParam].
type SplitTunnelIncludeUnionParam interface {
	implementsSplitTunnelIncludeUnionParam()
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam struct {
	// The address in CIDR format to include in the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithAddressParam) implementsSplitTunnelIncludeUnionParam() {
}

type SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam struct {
	// The domain name to include in the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description"`
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitTunnelIncludeTeamsDevicesIncludeSplitTunnelWithHostParam) implementsSplitTunnelIncludeUnionParam() {
}
