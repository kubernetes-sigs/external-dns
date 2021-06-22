package linodego

import (
	"context"
	"encoding/json"
	"fmt"
)

// IPAddressesPagedResponse represents a paginated IPAddress API response
type IPAddressesPagedResponse struct {
	*PageOptions
	Data []InstanceIP `json:"data"`
}

// IPAddressUpdateOptions fields are those accepted by UpdateToken
type IPAddressUpdateOptions struct {
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if set to nil.
	RDNS *string `json:"rdns"`
}

<<<<<<< HEAD
// LinodeIPAssignment stores an assignment between an IP address and a Linode instance.
type LinodeIPAssignment struct {
	Address  string `json:"address"`
	LinodeID int    `json:"linode_id"`
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

// GetUpdateOptions converts a IPAddress to IPAddressUpdateOptions for use in UpdateIPAddress
func (i InstanceIP) GetUpdateOptions() (o IPAddressUpdateOptions) {
	o.RDNS = copyString(&i.RDNS)
	return
}

// endpoint gets the endpoint URL for IPAddress
func (IPAddressesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.IPAddresses.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends IPAddresses when processing paginated InstanceIPAddress responses
func (resp *IPAddressesPagedResponse) appendData(r *IPAddressesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListIPAddresses lists IPAddresses
func (c *Client) ListIPAddresses(ctx context.Context, opts *ListOptions) ([]InstanceIP, error) {
	response := IPAddressesPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetIPAddress gets the template with the provided ID
func (c *Client) GetIPAddress(ctx context.Context, id string) (*InstanceIP, error) {
	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// UpdateIPAddress updates the IPAddress with the specified id
func (c *Client) UpdateIPAddress(ctx context.Context, id string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)

	req := c.R(ctx).SetResult(&InstanceIP{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// InstancesAssignIPs assigns multiple IPv4 addresses and/or IPv6 ranges to multiple Linodes in one Region.
// This allows swapping, shuffling, or otherwise reorganizing IPs to your Linodes.
func (c *Client) InstancesAssignIPs(ctx context.Context, updateOpts LinodesAssignIPsOptions) error {
	var body string

	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/assign", e)

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))

	return err
}

// ShareIPAddresses allows IP address reassignment (also referred to as IP failover)
// from one Linode to another if the primary Linode becomes unresponsive.
func (c *Client) ShareIPAddresses(ctx context.Context, shareOpts IPAddressesShareOptions) error {
	var body string

	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return err
	}

	e = fmt.Sprintf("%s/share", e)

	if bodyData, err := json.Marshal(shareOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))

	return err
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// GetUpdateOptions converts a IPAddress to IPAddressUpdateOptions for use in UpdateIPAddress
func (i InstanceIP) GetUpdateOptions() (o IPAddressUpdateOptions) {
	o.RDNS = copyString(&i.RDNS)
	return
}

// endpoint gets the endpoint URL for IPAddress
func (IPAddressesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.IPAddresses.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends IPAddresses when processing paginated InstanceIPAddress responses
func (resp *IPAddressesPagedResponse) appendData(r *IPAddressesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListIPAddresses lists IPAddresses
func (c *Client) ListIPAddresses(ctx context.Context, opts *ListOptions) ([]InstanceIP, error) {
	response := IPAddressesPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetIPAddress gets the template with the provided ID
func (c *Client) GetIPAddress(ctx context.Context, id string) (*InstanceIP, error) {
	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceIP{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceIP), nil
}

// UpdateIPAddress updates the IPAddress with the specified id
func (c *Client) UpdateIPAddress(ctx context.Context, id string, updateOpts IPAddressUpdateOptions) (*InstanceIP, error) {
	var body string
	e, err := c.IPAddresses.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)

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
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
