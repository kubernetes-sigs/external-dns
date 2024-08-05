package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// VPCSubnetLinodeInterface represents an interface on a Linode that is currently
// assigned to this VPC subnet.
type VPCSubnetLinodeInterface struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}

// VPCSubnetLinode represents a Linode currently assigned to a VPC subnet.
type VPCSubnetLinode struct {
	ID         int                        `json:"id"`
	Interfaces []VPCSubnetLinodeInterface `json:"interfaces"`
}

type VPCSubnet struct {
	ID      int               `json:"id"`
	Label   string            `json:"label"`
	IPv4    string            `json:"ipv4"`
	Linodes []VPCSubnetLinode `json:"linodes"`
	Created *time.Time        `json:"-"`
	Updated *time.Time        `json:"-"`
}

type VPCSubnetCreateOptions struct {
	Label string `json:"label"`
	IPv4  string `json:"ipv4"`
}

type VPCSubnetUpdateOptions struct {
	Label string `json:"label"`
}

type VPCSubnetsPagedResponse struct {
	*PageOptions
	Data []VPCSubnet `json:"data"`
}

func (VPCSubnetsPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("vpcs/%d/subnets", id)
}

func (resp *VPCSubnetsPagedResponse) castResult(
	r *resty.Request,
	e string,
) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(VPCSubnetsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*VPCSubnetsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

func (v *VPCSubnet) UnmarshalJSON(b []byte) error {
	type Mask VPCSubnet
	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(v),
	}
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	v.Created = (*time.Time)(p.Created)
	v.Updated = (*time.Time)(p.Updated)

	return nil
}

func (v VPCSubnet) GetCreateOptions() VPCSubnetCreateOptions {
	return VPCSubnetCreateOptions{
		Label: v.Label,
		IPv4:  v.IPv4,
	}
}

func (v VPCSubnet) GetUpdateOptions() VPCSubnetUpdateOptions {
	return VPCSubnetUpdateOptions{Label: v.Label}
}

func (c *Client) CreateVPCSubnet(
	ctx context.Context,
	opts VPCSubnetCreateOptions,
	vpcID int,
) (*VPCSubnet, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&VPCSubnet{}).SetBody(string(body))
	e := fmt.Sprintf("vpcs/%d/subnets", vpcID)
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*VPCSubnet), nil
}

func (c *Client) GetVPCSubnet(
	ctx context.Context,
	vpcID int,
	subnetID int,
) (*VPCSubnet, error) {
	req := c.R(ctx).SetResult(&VPCSubnet{})

	e := fmt.Sprintf("vpcs/%d/subnets/%d", vpcID, subnetID)
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*VPCSubnet), nil
}

func (c *Client) ListVPCSubnets(
	ctx context.Context,
	vpcID int,
	opts *ListOptions,
) ([]VPCSubnet, error) {
	response := VPCSubnetsPagedResponse{}
	err := c.listHelper(ctx, &response, opts, vpcID)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *Client) UpdateVPCSubnet(
	ctx context.Context,
	vpcID int,
	subnetID int,
	opts VPCSubnetUpdateOptions,
) (*VPCSubnet, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&VPCSubnet{}).SetBody(body)
	e := fmt.Sprintf("vpcs/%d/subnets/%d", vpcID, subnetID)
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*VPCSubnet), nil
}

func (c *Client) DeleteVPCSubnet(ctx context.Context, vpcID int, subnetID int) error {
	e := fmt.Sprintf("vpcs/%d/subnets/%d", vpcID, subnetID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
