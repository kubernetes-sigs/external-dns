// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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

// SiteLANService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteLANService] method instead.
type SiteLANService struct {
	Options []option.RequestOption
}

// NewSiteLANService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteLANService(opts ...option.RequestOption) (r *SiteLANService) {
	r = &SiteLANService{}
	r.Options = opts
	return
}

// Creates a new Site LAN. If the site is in high availability mode,
// static_addressing is required along with secondary and virtual address.
func (r *SiteLANService) New(ctx context.Context, siteID string, params SiteLANNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[LAN], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", params.AccountID, siteID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Creates a new Site LAN. If the site is in high availability mode,
// static_addressing is required along with secondary and virtual address.
func (r *SiteLANService) NewAutoPaging(ctx context.Context, siteID string, params SiteLANNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LAN] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, siteID, params, opts...))
}

// Update a specific Site LAN.
func (r *SiteLANService) Update(ctx context.Context, siteID string, lanID string, params SiteLANUpdateParams, opts ...option.RequestOption) (res *LAN, err error) {
	var env SiteLANUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", params.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Site LANs associated with an account.
func (r *SiteLANService) List(ctx context.Context, siteID string, query SiteLANListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LAN], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", query.AccountID, siteID)
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

// Lists Site LANs associated with an account.
func (r *SiteLANService) ListAutoPaging(ctx context.Context, siteID string, query SiteLANListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LAN] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, siteID, query, opts...))
}

// Remove a specific Site LAN.
func (r *SiteLANService) Delete(ctx context.Context, siteID string, lanID string, body SiteLANDeleteParams, opts ...option.RequestOption) (res *LAN, err error) {
	var env SiteLANDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", body.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch a specific Site LAN.
func (r *SiteLANService) Edit(ctx context.Context, siteID string, lanID string, params SiteLANEditParams, opts ...option.RequestOption) (res *LAN, err error) {
	var env SiteLANEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", params.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site LAN.
func (r *SiteLANService) Get(ctx context.Context, siteID string, lanID string, query SiteLANGetParams, opts ...option.RequestOption) (res *LAN, err error) {
	var env SiteLANGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", query.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string      `json:"server_addresses"`
	JSON            dhcpRelayJSON `json:"-"`
}

// dhcpRelayJSON contains the JSON metadata for the struct [DHCPRelay]
type dhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dhcpRelayJSON) RawJSON() string {
	return r.raw
}

type DHCPRelayParam struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r DHCPRelayParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer  string   `json:"dns_server"`
	DNSServers []string `json:"dns_servers"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string `json:"reservations"`
	JSON         dhcpServerJSON    `json:"-"`
}

// dhcpServerJSON contains the JSON metadata for the struct [DHCPServer]
type dhcpServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	DNSServers    apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dhcpServerJSON) RawJSON() string {
	return r.raw
}

type DHCPServerParam struct {
	// A valid IPv4 address.
	DHCPPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer  param.Field[string]   `json:"dns_server"`
	DNSServers param.Field[[]string] `json:"dns_servers"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r DHCPServerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LAN struct {
	// Identifier
	ID string `json:"id"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool           `json:"ha_link"`
	Name          string         `json:"name"`
	Nat           Nat            `json:"nat"`
	Physport      int64          `json:"physport"`
	RoutedSubnets []RoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing LANStaticAddressing `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag int64   `json:"vlan_tag"`
	JSON    lanJSON `json:"-"`
}

// lanJSON contains the JSON metadata for the struct [LAN]
type lanJSON struct {
	ID               apijson.Field
	HaLink           apijson.Field
	Name             apijson.Field
	Nat              apijson.Field
	Physport         apijson.Field
	RoutedSubnets    apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lanJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type LANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string     `json:"address,required"`
	DHCPRelay  DHCPRelay  `json:"dhcp_relay"`
	DHCPServer DHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                  `json:"virtual_address"`
	JSON           lanStaticAddressingJSON `json:"-"`
}

// lanStaticAddressingJSON contains the JSON metadata for the struct
// [LANStaticAddressing]
type lanStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type LANStaticAddressingParam struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]          `json:"address,required"`
	DHCPRelay  param.Field[DHCPRelayParam]  `json:"dhcp_relay"`
	DHCPServer param.Field[DHCPServerParam] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r LANStaticAddressingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Nat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string  `json:"static_prefix"`
	JSON         natJSON `json:"-"`
}

// natJSON contains the JSON metadata for the struct [Nat]
type natJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Nat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r natJSON) RawJSON() string {
	return r.raw
}

type NatParam struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r NatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string           `json:"prefix,required"`
	Nat    Nat              `json:"nat"`
	JSON   routedSubnetJSON `json:"-"`
}

// routedSubnetJSON contains the JSON metadata for the struct [RoutedSubnet]
type routedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routedSubnetJSON) RawJSON() string {
	return r.raw
}

type RoutedSubnetParam struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]   `json:"prefix,required"`
	Nat    param.Field[NatParam] `json:"nat"`
}

func (r RoutedSubnetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Physport  param.Field[int64]  `json:"physport,required"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        param.Field[bool]                `json:"ha_link"`
	Name          param.Field[string]              `json:"name"`
	Nat           param.Field[NatParam]            `json:"nat"`
	RoutedSubnets param.Field[[]RoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[LANStaticAddressingParam] `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLANNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParams struct {
	// Identifier
	AccountID     param.Field[string]              `path:"account_id,required"`
	Name          param.Field[string]              `json:"name"`
	Nat           param.Field[NatParam]            `json:"nat"`
	Physport      param.Field[int64]               `json:"physport"`
	RoutedSubnets param.Field[[]RoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[LANStaticAddressingParam] `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLANUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LAN                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteLANUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANUpdateResponseEnvelope]
type siteLANUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANUpdateResponseEnvelopeSuccess bool

const (
	SiteLANUpdateResponseEnvelopeSuccessTrue SiteLANUpdateResponseEnvelopeSuccess = true
)

func (r SiteLANUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteLANDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteLANDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LAN                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteLANDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANDeleteResponseEnvelope]
type siteLANDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANDeleteResponseEnvelopeSuccess bool

const (
	SiteLANDeleteResponseEnvelopeSuccessTrue SiteLANDeleteResponseEnvelopeSuccess = true
)

func (r SiteLANDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANEditParams struct {
	// Identifier
	AccountID     param.Field[string]              `path:"account_id,required"`
	Name          param.Field[string]              `json:"name"`
	Nat           param.Field[NatParam]            `json:"nat"`
	Physport      param.Field[int64]               `json:"physport"`
	RoutedSubnets param.Field[[]RoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[LANStaticAddressingParam] `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLANEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LAN                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANEditResponseEnvelopeJSON    `json:"-"`
}

// siteLANEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANEditResponseEnvelope]
type siteLANEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANEditResponseEnvelopeSuccess bool

const (
	SiteLANEditResponseEnvelopeSuccessTrue SiteLANEditResponseEnvelopeSuccess = true
)

func (r SiteLANEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteLANGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LAN                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANGetResponseEnvelopeJSON    `json:"-"`
}

// siteLANGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANGetResponseEnvelope]
type siteLANGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANGetResponseEnvelopeSuccess bool

const (
	SiteLANGetResponseEnvelopeSuccessTrue SiteLANGetResponseEnvelopeSuccess = true
)

func (r SiteLANGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
