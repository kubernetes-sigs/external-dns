package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type LinodeInterface struct {
	ID           int                    `json:"id"`
	Version      int                    `json:"version"`
	MACAddress   string                 `json:"mac_address"`
	Created      *time.Time             `json:"-"`
	Updated      *time.Time             `json:"-"`
	DefaultRoute *InterfaceDefaultRoute `json:"default_route"`
	Public       *PublicInterface       `json:"public"`
	VPC          *VPCInterface          `json:"vpc"`
	VLAN         *VLANInterface         `json:"vlan"`
}

type InterfaceDefaultRoute struct {
	IPv4 *bool `json:"ipv4,omitempty"`
	IPv6 *bool `json:"ipv6,omitempty"`
}

type PublicInterface struct {
	IPv4 *PublicInterfaceIPv4 `json:"ipv4"`
	IPv6 *PublicInterfaceIPv6 `json:"ipv6"`
}

type PublicInterfaceIPv4 struct {
	Addresses []PublicInterfaceIPv4Address `json:"addresses"`
	Shared    []PublicInterfaceIPv4Shared  `json:"shared"`
}

type PublicInterfaceIPv6 struct {
	Ranges []PublicInterfaceIPv6Range `json:"ranges"`
	Shared []PublicInterfaceIPv6Range `json:"shared"`
	SLAAC  []PublicInterfaceIPv6SLAAC `json:"slaac"`
}

type PublicInterfaceIPv4Address struct {
	Address string `json:"address"`
	Primary bool   `json:"primary"`
}

type PublicInterfaceIPv4Shared struct {
	Address  string `json:"address"`
	LinodeID int    `json:"linode_id"`
}

type PublicInterfaceIPv6Range struct {
	Range       string  `json:"range"`
	RouteTarget *string `json:"route_target"`
}

type PublicInterfaceIPv6SLAAC struct {
	Prefix  int    `json:"prefix"`
	Address string `json:"address"`
}

type VPCInterface struct {
	VPCID    int              `json:"vpc_id"`
	SubnetID int              `json:"subnet_id"`
	IPv4     VPCInterfaceIPv4 `json:"ipv4"`

	// NOTE: IPv6 VPC interfaces may not currently be available to all users.
	IPv6 VPCInterfaceIPv6 `json:"ipv6"`
}

type VPCInterfaceIPv4 struct {
	Addresses []VPCInterfaceIPv4Address `json:"addresses"`
	Ranges    []VPCInterfaceIPv4Range   `json:"ranges"`
}

type VPCInterfaceIPv4Address struct {
	Address        string  `json:"address"`
	Primary        bool    `json:"primary"`
	NAT1To1Address *string `json:"nat_1_1_address"`
}

type VPCInterfaceIPv4Range struct {
	Range string `json:"range"`
}

// VPCInterfaceIPv6 contains the IPv6 configuration for a VPC.
// NOTE: IPv6 VPC interfaces may not currently be available to all users.
type VPCInterfaceIPv6 struct {
	SLAAC    []VPCInterfaceIPv6SLAAC `json:"slaac"`
	Ranges   []VPCInterfaceIPv6Range `json:"ranges"`
	IsPublic *bool                   `json:"is_public"`
}

// VPCInterfaceIPv6SLAAC contains the information for a single IPv6 SLAAC under a VPC.
// NOTE: IPv6 VPC interfaces may not currently be available to all users.
type VPCInterfaceIPv6SLAAC struct {
	Range   string `json:"range"`
	Address string `json:"address"`
}

// VPCInterfaceIPv6Range contains the information for a single IPv6 range under a VPC.
// NOTE: IPv6 VPC interfaces may not currently be available to all users.
type VPCInterfaceIPv6Range struct {
	Range string `json:"range"`
}

type VLANInterface struct {
	VLANLabel   string  `json:"vlan_label"`
	IPAMAddress *string `json:"ipam_address,omitempty"`
}

type LinodeInterfaceCreateOptions struct {
	FirewallID   *int                          `json:"firewall_id,omitempty"`
	DefaultRoute *InterfaceDefaultRoute        `json:"default_route,omitempty"`
	Public       *PublicInterfaceCreateOptions `json:"public,omitempty"`
	VPC          *VPCInterfaceCreateOptions    `json:"vpc,omitempty"`
	VLAN         *VLANInterface                `json:"vlan,omitempty"`
}

type LinodeInterfaceUpdateOptions struct {
	DefaultRoute *InterfaceDefaultRoute        `json:"default_route,omitempty"`
	Public       *PublicInterfaceCreateOptions `json:"public,omitempty"`
	VPC          *VPCInterfaceUpdateOptions    `json:"vpc,omitempty"`
}

type PublicInterfaceCreateOptions struct {
	IPv4 *PublicInterfaceIPv4CreateOptions `json:"ipv4,omitempty"`
	IPv6 *PublicInterfaceIPv6CreateOptions `json:"ipv6,omitempty"`
}

type PublicInterfaceIPv4CreateOptions struct {
	Addresses *[]PublicInterfaceIPv4AddressCreateOptions `json:"addresses,omitempty"`
}

type PublicInterfaceIPv4AddressCreateOptions struct {
	Address *string `json:"address,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
}

type PublicInterfaceIPv6CreateOptions struct {
	Ranges *[]PublicInterfaceIPv6RangeCreateOptions `json:"ranges,omitempty"`
}

type PublicInterfaceIPv6RangeCreateOptions struct {
	Range string `json:"range"`
}

type VPCInterfaceCreateOptions struct {
	SubnetID int                            `json:"subnet_id"`
	IPv4     *VPCInterfaceIPv4CreateOptions `json:"ipv4,omitempty"`
	IPv6     *VPCInterfaceIPv6CreateOptions `json:"ipv6,omitempty"`
}

type VPCInterfaceIPv4CreateOptions struct {
	Addresses *[]VPCInterfaceIPv4AddressCreateOptions `json:"addresses,omitempty"`
	Ranges    *[]VPCInterfaceIPv4RangeCreateOptions   `json:"ranges,omitempty"`
}

type VPCInterfaceIPv4AddressCreateOptions struct {
	Address        *string `json:"address,omitempty"`
	Primary        *bool   `json:"primary,omitempty"`
	NAT1To1Address *string `json:"nat_1_1_address,omitempty"`
}

type VPCInterfaceIPv4RangeCreateOptions struct {
	Range string `json:"range"`
}

// VPCInterfaceIPv6CreateOptions specifies IPv6 configuration parameters for VPC creation.
// NOTE: IPv6 interfaces may not currently be available to all users.
type VPCInterfaceIPv6CreateOptions struct {
	SLAAC    *[]VPCInterfaceIPv6SLAACCreateOptions `json:"slaac,omitempty"`
	Ranges   *[]VPCInterfaceIPv6RangeCreateOptions `json:"ranges,omitempty"`
	IsPublic *bool                                 `json:"is_public"`
}

// VPCInterfaceIPv6SLAACCreateOptions defines the IPv6 SLAAC configuration parameters for VPC creation.
// NOTE: IPv6 interfaces may not currently be available to all users.
type VPCInterfaceIPv6SLAACCreateOptions struct {
	Range string `json:"range"`
}

// VPCInterfaceIPv6RangeCreateOptions defines the IPv6 range configuration parameters for VPC creation.
// NOTE: IPv6 interfaces may not currently be available to all users.
type VPCInterfaceIPv6RangeCreateOptions struct {
	Range string `json:"range"`
}

type VPCInterfaceUpdateOptions struct {
	IPv4 *VPCInterfaceIPv4CreateOptions `json:"ipv4,omitempty"`
	IPv6 *VPCInterfaceIPv6CreateOptions `json:"ipv6,omitempty"`
}

type LinodeInterfacesUpgrade struct {
	ConfigID   int               `json:"config_id"`
	DryRun     bool              `json:"dry_run"`
	Interfaces []LinodeInterface `json:"interfaces"`
}

type LinodeInterfacesUpgradeOptions struct {
	ConfigID *int  `json:"config_id,omitempty"`
	DryRun   *bool `json:"dry_run,omitempty"`
}

type InterfaceSettings struct {
	NetworkHelper bool                         `json:"network_helper"`
	DefaultRoute  InterfaceDefaultRouteSetting `json:"default_route"`
}

type InterfaceSettingsUpdateOptions struct {
	NetworkHelper *bool                                      `json:"network_helper,omitempty"`
	DefaultRoute  *InterfaceDefaultRouteSettingUpdateOptions `json:"default_route,omitempty"`
}

type InterfaceDefaultRouteSettingUpdateOptions struct {
	IPv4InterfaceID *int `json:"ipv4_interface_id,omitempty"`
	IPv6InterfaceID *int `json:"ipv6_interface_id,omitempty"`
}

type InterfaceDefaultRouteSetting struct {
	IPv4InterfaceID          *int  `json:"ipv4_interface_id"`
	IPv4EligibleInterfaceIDs []int `json:"ipv4_eligible_interface_ids"`
	IPv6InterfaceID          *int  `json:"ipv6_interface_id"`
	IPv6EligibleInterfaceIDs []int `json:"ipv6_eligible_interface_ids"`
}

func (i *LinodeInterface) UnmarshalJSON(b []byte) error {
	type Mask LinodeInterface

	p := struct {
		*Mask

		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

func (c *Client) ListInterfaces(ctx context.Context, linodeID int, opts *ListOptions) ([]LinodeInterface, error) {
	e := formatAPIPath("linode/instances/%d/interfaces", linodeID)
	return getPaginatedResults[LinodeInterface](ctx, c, e, opts)
}

func (c *Client) GetInterface(ctx context.Context, linodeID int, interfaceID int) (*LinodeInterface, error) {
	e := formatAPIPath("linode/instances/%d/interfaces/%d", linodeID, interfaceID)
	return doGETRequest[LinodeInterface](ctx, c, e)
}

func (c *Client) CreateInterface(ctx context.Context, linodeID int, opts LinodeInterfaceCreateOptions) (*LinodeInterface, error) {
	e := formatAPIPath("linode/instances/%d/interfaces", linodeID)
	return doPOSTRequest[LinodeInterface](ctx, c, e, opts)
}

func (c *Client) UpdateInterface(ctx context.Context, linodeID int, interfaceID int, opts LinodeInterfaceUpdateOptions) (*LinodeInterface, error) {
	e := formatAPIPath("linode/instances/%d/interfaces/%d", linodeID, interfaceID)
	return doPUTRequest[LinodeInterface](ctx, c, e, opts)
}

func (c *Client) DeleteInterface(ctx context.Context, linodeID int, interfaceID int) error {
	e := formatAPIPath("linode/instances/%d/interfaces/%d", linodeID, interfaceID)
	return doDELETERequest(ctx, c, e)
}

func (c *Client) UpgradeInterfaces(ctx context.Context, linodeID int, opts LinodeInterfacesUpgradeOptions) (*LinodeInterfacesUpgrade, error) {
	e := formatAPIPath("linode/instances/%d/upgrade-interfaces", linodeID)
	return doPOSTRequest[LinodeInterfacesUpgrade](ctx, c, e, opts)
}

func (c *Client) ListInterfaceFirewalls(ctx context.Context, linodeID int, interfaceID int, opts *ListOptions) ([]Firewall, error) {
	e := formatAPIPath("linode/instances/%d/interfaces/%d/firewalls", linodeID, interfaceID)
	return getPaginatedResults[Firewall](ctx, c, e, opts)
}

func (c *Client) GetInterfaceSettings(ctx context.Context, linodeID int) (*InterfaceSettings, error) {
	e := formatAPIPath("linode/instances/%d/interfaces/settings", linodeID)
	return doGETRequest[InterfaceSettings](ctx, c, e)
}

func (c *Client) UpdateInterfaceSettings(ctx context.Context, linodeID int, opts InterfaceSettingsUpdateOptions) (*InterfaceSettings, error) {
	e := formatAPIPath("linode/instances/%d/interfaces/settings", linodeID)
	return doPUTRequest[InterfaceSettings](ctx, c, e, opts)
}
