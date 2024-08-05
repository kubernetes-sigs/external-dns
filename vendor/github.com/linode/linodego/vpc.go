package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

type VPC struct {
	ID          int         `json:"id"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Region      string      `json:"region"`
	Subnets     []VPCSubnet `json:"subnets"`
	Created     *time.Time  `json:"-"`
	Updated     *time.Time  `json:"-"`
}

type VPCCreateOptions struct {
	Label       string                   `json:"label"`
	Description string                   `json:"description,omitempty"`
	Region      string                   `json:"region"`
	Subnets     []VPCSubnetCreateOptions `json:"subnets,omitempty"`
}

type VPCUpdateOptions struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

type VPCsPagedResponse struct {
	*PageOptions
	Data []VPC `json:"data"`
}

func (VPCsPagedResponse) endpoint(_ ...any) string {
	return "vpcs"
}

func (resp *VPCsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(VPCsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*VPCsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

func (v VPC) GetCreateOptions() VPCCreateOptions {
	subnetCreations := make([]VPCSubnetCreateOptions, len(v.Subnets))
	for i, s := range v.Subnets {
		subnetCreations[i] = s.GetCreateOptions()
	}

	return VPCCreateOptions{
		Label:       v.Label,
		Description: v.Description,
		Region:      v.Region,
		Subnets:     subnetCreations,
	}
}

func (v VPC) GetUpdateOptions() VPCUpdateOptions {
	return VPCUpdateOptions{
		Label:       v.Label,
		Description: v.Description,
	}
}

func (v *VPC) UnmarshalJSON(b []byte) error {
	type Mask VPC
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

func (c *Client) CreateVPC(
	ctx context.Context,
	opts VPCCreateOptions,
) (*VPC, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&VPC{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post("vpcs"))
	if err != nil {
		return nil, err
	}

	return r.Result().(*VPC), nil
}

func (c *Client) GetVPC(ctx context.Context, vpcID int) (*VPC, error) {
	e := fmt.Sprintf("/vpcs/%d", vpcID)
	req := c.R(ctx).SetResult(&VPC{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*VPC), nil
}

func (c *Client) ListVPCs(ctx context.Context, opts *ListOptions) ([]VPC, error) {
	response := VPCsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *Client) UpdateVPC(
	ctx context.Context,
	vpcID int,
	opts VPCUpdateOptions,
) (*VPC, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("vpcs/%d", vpcID)
	req := c.R(ctx).SetResult(&VPC{}).SetBody(body)
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*VPC), nil
}

func (c *Client) DeleteVPC(ctx context.Context, vpcID int) error {
	e := fmt.Sprintf("vpcs/%d", vpcID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
