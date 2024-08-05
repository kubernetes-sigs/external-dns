package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// InstanceIPAddressResponse contains the IPv4 and IPv6 details for an Instance
type InstanceIPAddressResponse struct {
	IPv4 *InstanceIPv4Response `json:"ipv4"`
	IPv6 *InstanceIPv6Response `json:"ipv6"`
}

// InstanceIPv4Response contains the details of all IPv4 addresses associated with an Instance
type InstanceIPv4Response struct {
	Public   []*InstanceIP `json:"public"`
	Private  []*InstanceIP `json:"private"`
	Shared   []*InstanceIP `json:"shared"`
	Reserved []*InstanceIP `json:"reserved"`
	VPC      []*VPCIP      `json:"vpc"`
}

// InstanceIP represents an Instance IP with additional DNS and networking details
type InstanceIP struct {
	Address    string             `json:"address"`
	Gateway    string             `json:"gateway"`
	SubnetMask string             `json:"subnet_mask"`
	Prefix     int                `json:"prefix"`
	Type       InstanceIPType     `json:"type"`
	Public     bool               `json:"public"`
	RDNS       string             `json:"rdns"`
	LinodeID   int                `json:"linode_id"`
	Region     string             `json:"region"`
	VPCNAT1To1 *InstanceIPNAT1To1 `json:"vpc_nat_1_1"`
}

// VPCIP represents a private IP address in a VPC subnet with additional networking details
type VPCIP struct {
	Address      *string `json:"address"`
	AddressRange *string `json:"address_range"`
	Gateway      string  `json:"gateway"`
	SubnetMask   string  `json:"subnet_mask"`
	Prefix       int     `json:"prefix"`
	LinodeID     int     `json:"linode_id"`
	Region       string  `json:"region"`
	Active       bool    `json:"active"`
	NAT1To1      *string `json:"nat_1_1"`
	VPCID        int     `json:"vpc_id"`
	SubnetID     int     `json:"subnet_id"`
	ConfigID     int     `json:"config_id"`
	InterfaceID  int     `json:"interface_id"`
}

// InstanceIPv6Response contains the IPv6 addresses and ranges for an Instance
type InstanceIPv6Response struct {
<<<<<<< HEAD
<<<<<<< HEAD
	LinkLocal *InstanceIP `json:"link_local"`
	SLAAC     *InstanceIP `json:"slaac"`
	Global    []IPv6Range `json:"global"`
}

// IPv6Range represents a range of IPv6 addresses routed to a single Linode in a given Region
type IPv6Range struct {
	Range  string `json:"range"`
	Region string `json:"region"`
	Prefix int    `json:"prefix"`

	RouteTarget string `json:"route_target"`

	// These fields are only returned by GetIPv6Range(...)
	IsBGP   bool  `json:"is_bgp"`
	Linodes []int `json:"linodes"`
}

// InstanceIPType constants start with IPType and include Linode Instance IP Types
type InstanceIPType string

// InstanceIPType constants represent the IP types an Instance IP may be
const (
	IPTypeIPv4      InstanceIPType = "ipv4"
	IPTypeIPv6      InstanceIPType = "ipv6"
	IPTypeIPv6Pool  InstanceIPType = "ipv6/pool"
	IPTypeIPv6Range InstanceIPType = "ipv6/range"
)

// GetInstanceIPAddresses gets the IPAddresses for a Linode instance
func (c *Client) GetInstanceIPAddresses(ctx context.Context, linodeID int) (*InstanceIPAddressResponse, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIPAddressResponse{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIPAddressResponse), nil
}

// GetInstanceIPAddress gets the IPAddress for a Linode instance matching a supplied IP address
func (c *Client) GetInstanceIPAddress(ctx context.Context, linodeID int, ipaddress string) (*InstanceIP, error) {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipaddress)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// AddInstanceIPAddress adds a public or private IP to a Linode instance
func (c *Client) AddInstanceIPAddress(ctx context.Context, linodeID int, public bool) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceIP{})

	instanceipRequest := struct {
		Type   string `json:"type"`
		Public bool   `json:"public"`
	}{"ipv4", public}

	if bodyData, err := json.Marshal(instanceipRequest); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// UpdateInstanceIPAddress updates the IPAddress with the specified instance id and IP address
func (c *Client) UpdateInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipAddress)

	req := c.R(ctx).SetResult(&InstanceIP{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

func (c *Client) DeleteInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string) error {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s", e, ipAddress)
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.InstanceIPs.endpointWithID(linodeID)
||||||| parent of 5ce8c7613 (update vendored files)
	e, err := c.InstanceIPs.endpointWithID(linodeID)
=======
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIPAddressResponse{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIPAddressResponse), nil
}

// GetInstanceIPAddress gets the IPAddress for a Linode instance matching a supplied IP address
func (c *Client) GetInstanceIPAddress(ctx context.Context, linodeID int, ipaddress string) (*InstanceIP, error) {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipaddress)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// AddInstanceIPAddress adds a public or private IP to a Linode instance
func (c *Client) AddInstanceIPAddress(ctx context.Context, linodeID int, public bool) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceIP{})

	instanceipRequest := struct {
		Type   string `json:"type"`
		Public bool   `json:"public"`
	}{"ipv4", public}

	if bodyData, err := json.Marshal(instanceipRequest); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// UpdateInstanceIPAddress updates the IPAddress with the specified instance id and IP address
func (c *Client) UpdateInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipAddress)

	req := c.R(ctx).SetResult(&InstanceIP{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func (c *Client) DeleteInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string) error {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s", e, ipAddress)
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.InstanceIPs.endpointWithID(linodeID)
||||||| parent of 6b7ce455e (update vendored files)
	e, err := c.InstanceIPs.endpointWithID(linodeID)
=======
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIPAddressResponse{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIPAddressResponse), nil
}

// GetInstanceIPAddress gets the IPAddress for a Linode instance matching a supplied IP address
func (c *Client) GetInstanceIPAddress(ctx context.Context, linodeID int, ipaddress string) (*InstanceIP, error) {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipaddress)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// AddInstanceIPAddress adds a public or private IP to a Linode instance
func (c *Client) AddInstanceIPAddress(ctx context.Context, linodeID int, public bool) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceIP{})

	instanceipRequest := struct {
		Type   string `json:"type"`
		Public bool   `json:"public"`
	}{"ipv4", public}

	if bodyData, err := json.Marshal(instanceipRequest); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// UpdateInstanceIPAddress updates the IPAddress with the specified instance id and IP address
func (c *Client) UpdateInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipAddress)

	req := c.R(ctx).SetResult(&InstanceIP{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func (c *Client) DeleteInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string) error {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s", e, ipAddress)
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.InstanceIPs.endpointWithID(linodeID)
||||||| parent of 4d7e5ad26 (update vendored files)
	e, err := c.InstanceIPs.endpointWithID(linodeID)
=======
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIPAddressResponse{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIPAddressResponse), nil
}

// GetInstanceIPAddress gets the IPAddress for a Linode instance matching a supplied IP address
func (c *Client) GetInstanceIPAddress(ctx context.Context, linodeID int, ipaddress string) (*InstanceIP, error) {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipaddress)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// AddInstanceIPAddress adds a public or private IP to a Linode instance
func (c *Client) AddInstanceIPAddress(ctx context.Context, linodeID int, public bool) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceIP{})

	instanceipRequest := struct {
		Type   string `json:"type"`
		Public bool   `json:"public"`
	}{"ipv4", public}

	if bodyData, err := json.Marshal(instanceipRequest); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// UpdateInstanceIPAddress updates the IPAddress with the specified instance id and IP address
func (c *Client) UpdateInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, ipAddress)

	req := c.R(ctx).SetResult(&InstanceIP{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func (c *Client) DeleteInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string) error {
	e, err := c.InstanceIPs.endpointWithParams(linodeID)
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/%s", e, ipAddress)
	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
	return err
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	LinkLocal *InstanceIP  `json:"link_local"`
	SLAAC     *InstanceIP  `json:"slaac"`
	Global    []*IPv6Range `json:"global"`
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	LinkLocal *InstanceIP  `json:"link_local"`
	SLAAC     *InstanceIP  `json:"slaac"`
	Global    []*IPv6Range `json:"global"`
=======
	LinkLocal *InstanceIP `json:"link_local"`
	SLAAC     *InstanceIP `json:"slaac"`
	Global    []IPv6Range `json:"global"`
}

// InstanceIPNAT1To1 contains information about the NAT 1:1 mapping
// of a public IP address to a VPC subnet.
type InstanceIPNAT1To1 struct {
	Address  string `json:"address"`
	SubnetID int    `json:"subnet_id"`
	VPCID    int    `json:"vpc_id"`
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

// IPv6Range represents a range of IPv6 addresses routed to a single Linode in a given Region
type IPv6Range struct {
	Range  string `json:"range"`
	Region string `json:"region"`
	Prefix int    `json:"prefix"`

	RouteTarget string `json:"route_target"`

	// These fields are only returned by GetIPv6Range(...)
	IsBGP   bool  `json:"is_bgp"`
	Linodes []int `json:"linodes"`
}

// InstanceIPType constants start with IPType and include Linode Instance IP Types
type InstanceIPType string

// InstanceIPType constants represent the IP types an Instance IP may be
const (
	IPTypeIPv4      InstanceIPType = "ipv4"
	IPTypeIPv6      InstanceIPType = "ipv6"
	IPTypeIPv6Pool  InstanceIPType = "ipv6/pool"
	IPTypeIPv6Range InstanceIPType = "ipv6/range"
)

// GetInstanceIPAddresses gets the IPAddresses for a Linode instance
func (c *Client) GetInstanceIPAddresses(ctx context.Context, linodeID int) (*InstanceIPAddressResponse, error) {
	e := fmt.Sprintf("linode/instances/%d/ips", linodeID)
	req := c.R(ctx).SetResult(&InstanceIPAddressResponse{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIPAddressResponse), nil
}

// GetInstanceIPAddress gets the IPAddress for a Linode instance matching a supplied IP address
func (c *Client) GetInstanceIPAddress(ctx context.Context, linodeID int, ipaddress string) (*InstanceIP, error) {
	ipaddress = url.PathEscape(ipaddress)
	e := fmt.Sprintf("linode/instances/%d/ips/%s", linodeID, ipaddress)
	req := c.R(ctx).SetResult(&InstanceIP{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// AddInstanceIPAddress adds a public or private IP to a Linode instance
func (c *Client) AddInstanceIPAddress(ctx context.Context, linodeID int, public bool) (*InstanceIP, error) {
	instanceipRequest := struct {
		Type   string `json:"type"`
		Public bool   `json:"public"`
	}{"ipv4", public}

	body, err := json.Marshal(instanceipRequest)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("linode/instances/%d/ips", linodeID)
	req := c.R(ctx).SetResult(&InstanceIP{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceIP), nil
}

// UpdateInstanceIPAddress updates the IPAddress with the specified instance id and IP address
func (c *Client) UpdateInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string, opts IPAddressUpdateOptions) (*InstanceIP, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	ipAddress = url.PathEscape(ipAddress)

	e := fmt.Sprintf("linode/instances/%d/ips/%s", linodeID, ipAddress)
	req := c.R(ctx).SetResult(&InstanceIP{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func (c *Client) DeleteInstanceIPAddress(ctx context.Context, linodeID int, ipAddress string) error {
	ipAddress = url.PathEscape(ipAddress)
	e := fmt.Sprintf("linode/instances/%d/ips/%s", linodeID, ipAddress)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
