package linodego

import (
	"context"
)

// IPAddressUpdateOptionsV2 fields are those accepted by UpdateIPAddress.
// NOTE: An IP's RDNS can be reset to default using the following pattern:
//
//	IPAddressUpdateOptionsV2{
//		RDNS: linodego.Pointer[*string](nil),
//	}
type IPAddressUpdateOptionsV2 struct {
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if set to nil.
	Reserved *bool    `json:"reserved,omitempty"`
	RDNS     **string `json:"rdns,omitempty"`
}

// IPAddressUpdateOptions fields are those accepted by UpdateIPAddress.
//
// Deprecated: Please use IPAddressUpdateOptionsV2 for all new implementations.
type IPAddressUpdateOptions struct {
	RDNS *string `json:"rdns"`
}

// LinodeIPAssignment stores an assignment between an IP address and a Linode instance.
type LinodeIPAssignment struct {
	Address  string `json:"address"`
	LinodeID int    `json:"linode_id"`
}

type AllocateReserveIPOptions struct {
	Type     string `json:"type"`
	Public   bool   `json:"public"`
	Reserved bool   `json:"reserved,omitempty"`
	Region   string `json:"region,omitempty"`
	LinodeID int    `json:"linode_id,omitempty"`
}

// LinodesAssignIPsOptions fields are those accepted by InstancesAssignIPs.
type LinodesAssignIPsOptions struct {
	Region string `json:"region"`

	Assignments []LinodeIPAssignment `json:"assignments"`
}

// IPAddressesShareOptions fields are those accepted by ShareIPAddresses.
type IPAddressesShareOptions struct {
	IPs      []string `json:"ips"`
	LinodeID int      `json:"linode_id"`
}

// ListIPAddressesQuery fields are those accepted as query params for the
// ListIPAddresses function.
type ListIPAddressesQuery struct {
	SkipIPv6RDNS bool `query:"skip_ipv6_rdns"`
}

// GetUpdateOptionsV2 converts a IPAddress to IPAddressUpdateOptionsV2 for use in UpdateIPAddressV2.
func (i InstanceIP) GetUpdateOptionsV2() IPAddressUpdateOptionsV2 {
	rdns := copyString(&i.RDNS)

	return IPAddressUpdateOptionsV2{
		RDNS:     &rdns,
		Reserved: copyBool(&i.Reserved),
	}
}

// GetUpdateOptions converts a IPAddress to IPAddressUpdateOptions for use in UpdateIPAddress.
//
// Deprecated: Please use GetUpdateOptionsV2 for all new implementations.
func (i InstanceIP) GetUpdateOptions() (o IPAddressUpdateOptions) {
	o.RDNS = copyString(&i.RDNS)
	return o
}

// ListIPAddresses lists IPAddresses.
func (c *Client) ListIPAddresses(ctx context.Context, opts *ListOptions) ([]InstanceIP, error) {
	return getPaginatedResults[InstanceIP](ctx, c, "networking/ips", opts)
}

// GetIPAddress gets the IPAddress with the provided IP.
func (c *Client) GetIPAddress(ctx context.Context, id string) (*InstanceIP, error) {
	e := formatAPIPath("networking/ips/%s", id)
	return doGETRequest[InstanceIP](ctx, c, e)
}

// UpdateIPAddressV2 updates the IP address with the specified address.
func (c *Client) UpdateIPAddressV2(ctx context.Context, address string, opts IPAddressUpdateOptionsV2) (*InstanceIP, error) {
	e := formatAPIPath("networking/ips/%s", address)
	return doPUTRequest[InstanceIP](ctx, c, e, opts)
}

// UpdateIPAddress updates the IP address with the specified id.
//
// Deprecated: Please use UpdateIPAddressV2 for all new implementation.
func (c *Client) UpdateIPAddress(ctx context.Context, id string, opts IPAddressUpdateOptions) (*InstanceIP, error) {
	e := formatAPIPath("networking/ips/%s", id)
	return doPUTRequest[InstanceIP](ctx, c, e, opts)
}

// InstancesAssignIPs assigns multiple IPv4 addresses and/or IPv6 ranges to multiple Linodes in one Region.
// This allows swapping, shuffling, or otherwise reorganizing IPs to your Linodes.
func (c *Client) InstancesAssignIPs(ctx context.Context, opts LinodesAssignIPsOptions) error {
	return doPOSTRequestNoResponseBody(ctx, c, "networking/ips/assign", opts)
}

// ShareIPAddresses allows IP address reassignment (also referred to as IP failover)
// from one Linode to another if the primary Linode becomes unresponsive.
func (c *Client) ShareIPAddresses(ctx context.Context, opts IPAddressesShareOptions) error {
	return doPOSTRequestNoResponseBody(ctx, c, "networking/ips/share", opts)
}

// AllocateReserveIP allocates a new IPv4 address to the Account, with the option to reserve it
// and optionally assign it to a Linode.
func (c *Client) AllocateReserveIP(ctx context.Context, opts AllocateReserveIPOptions) (*InstanceIP, error) {
	return doPOSTRequest[InstanceIP](ctx, c, "networking/ips", opts)
}
