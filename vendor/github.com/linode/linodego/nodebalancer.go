package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// NodeBalancer represents a NodeBalancer object
type NodeBalancer struct {
	// This NodeBalancer's unique ID.
	ID int `json:"id"`
	// This NodeBalancer's label. These must be unique on your Account.
	Label *string `json:"label"`
	// The Region where this NodeBalancer is located. NodeBalancers only support backends in the same Region.
	Region string `json:"region"`
	// This NodeBalancer's hostname, ending with .nodebalancer.linode.com
	Hostname *string `json:"hostname"`
	// This NodeBalancer's public IPv4 address.
	IPv4 *string `json:"ipv4"`
	// This NodeBalancer's public IPv6 address.
	IPv6 *string `json:"ipv6"`
	// Throttle connections per second (0-20). Set to 0 (zero) to disable throttling.
	ClientConnThrottle int `json:"client_conn_throttle"`
	// Information about the amount of transfer this NodeBalancer has had so far this month.
	Transfer NodeBalancerTransfer `json:"transfer"`

	// An array of tags applied to this object. Tags are for organizational purposes only.
	Tags []string `json:"tags"`

	Created *time.Time `json:"-"`
	Updated *time.Time `json:"-"`
}

// NodeBalancerTransfer contains information about the amount of transfer a NodeBalancer has had in the current month
type NodeBalancerTransfer struct {
	// The total transfer, in MB, used by this NodeBalancer this month.
	Total *float64 `json:"total"`
	// The total inbound transfer, in MB, used for this NodeBalancer this month.
	Out *float64 `json:"out"`
	// The total outbound transfer, in MB, used for this NodeBalancer this month.
	In *float64 `json:"in"`
}

// NodeBalancerCreateOptions are the options permitted for CreateNodeBalancer
type NodeBalancerCreateOptions struct {
	Label              *string                            `json:"label,omitempty"`
	Region             string                             `json:"region,omitempty"`
	ClientConnThrottle *int                               `json:"client_conn_throttle,omitempty"`
	Configs            []*NodeBalancerConfigCreateOptions `json:"configs,omitempty"`
	Tags               []string                           `json:"tags"`
	FirewallID         int                                `json:"firewall_id,omitempty"`
}

// NodeBalancerUpdateOptions are the options permitted for UpdateNodeBalancer
type NodeBalancerUpdateOptions struct {
	Label              *string   `json:"label,omitempty"`
	ClientConnThrottle *int      `json:"client_conn_throttle,omitempty"`
	Tags               *[]string `json:"tags,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *NodeBalancer) UnmarshalJSON(b []byte) error {
	type Mask NodeBalancer

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

// GetCreateOptions converts a NodeBalancer to NodeBalancerCreateOptions for use in CreateNodeBalancer
func (i NodeBalancer) GetCreateOptions() NodeBalancerCreateOptions {
	return NodeBalancerCreateOptions{
		Label:              i.Label,
		Region:             i.Region,
		ClientConnThrottle: &i.ClientConnThrottle,
		Tags:               i.Tags,
	}
}

// GetUpdateOptions converts a NodeBalancer to NodeBalancerUpdateOptions for use in UpdateNodeBalancer
func (i NodeBalancer) GetUpdateOptions() NodeBalancerUpdateOptions {
	return NodeBalancerUpdateOptions{
		Label:              i.Label,
		ClientConnThrottle: &i.ClientConnThrottle,
		Tags:               &i.Tags,
	}
}

// NodeBalancersPagedResponse represents a paginated NodeBalancer API response
type NodeBalancersPagedResponse struct {
	*PageOptions
	Data []NodeBalancer `json:"data"`
}

func (*NodeBalancersPagedResponse) endpoint(_ ...any) string {
	return "nodebalancers"
}

func (resp *NodeBalancersPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(NodeBalancersPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*NodeBalancersPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListNodeBalancers lists NodeBalancers
func (c *Client) ListNodeBalancers(ctx context.Context, opts *ListOptions) ([]NodeBalancer, error) {
	response := NodeBalancersPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetNodeBalancer gets the NodeBalancer with the provided ID
func (c *Client) GetNodeBalancer(ctx context.Context, nodebalancerID int) (*NodeBalancer, error) {
	e := fmt.Sprintf("nodebalancers/%d", nodebalancerID)
	req := c.R(ctx).SetResult(&NodeBalancer{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// CreateNodeBalancer creates a NodeBalancer
func (c *Client) CreateNodeBalancer(ctx context.Context, opts NodeBalancerCreateOptions) (*NodeBalancer, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(nodebalancer); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// UpdateNodeBalancer updates the NodeBalancer with the specified id
func (c *Client) UpdateNodeBalancer(ctx context.Context, id int, updateOpts NodeBalancerUpdateOptions) (*NodeBalancer, error) {
	var body string
	e, err := c.NodeBalancers.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// UpdateNodeBalancer updates the NodeBalancer with the specified id
func (c *Client) UpdateNodeBalancer(ctx context.Context, id int, updateOpts NodeBalancerUpdateOptions) (*NodeBalancer, error) {
	var body string
	e, err := c.NodeBalancers.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// UpdateNodeBalancer updates the NodeBalancer with the specified id
func (c *Client) UpdateNodeBalancer(ctx context.Context, id int, updateOpts NodeBalancerUpdateOptions) (*NodeBalancer, error) {
	var body string
	e, err := c.NodeBalancers.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// UpdateNodeBalancer updates the NodeBalancer with the specified id
func (c *Client) UpdateNodeBalancer(ctx context.Context, id int, updateOpts NodeBalancerUpdateOptions) (*NodeBalancer, error) {
	var body string
	e, err := c.NodeBalancers.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
<<<<<<< HEAD

>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	req := c.R(ctx).SetResult(&NodeBalancer{})

	if bodyData, err := json.Marshal(nodebalancer); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))

=======
	e := "nodebalancers"
	req := c.R(ctx).SetResult(&NodeBalancer{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// UpdateNodeBalancer updates the NodeBalancer with the specified id
func (c *Client) UpdateNodeBalancer(ctx context.Context, nodebalancerID int, opts NodeBalancerUpdateOptions) (*NodeBalancer, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	e := fmt.Sprintf("nodebalancers/%d", nodebalancerID)
	req := c.R(ctx).SetResult(&NodeBalancer{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancer), nil
}

// DeleteNodeBalancer deletes the NodeBalancer with the specified id
func (c *Client) DeleteNodeBalancer(ctx context.Context, nodebalancerID int) error {
	e := fmt.Sprintf("nodebalancers/%d", nodebalancerID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
