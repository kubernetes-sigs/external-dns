package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type VPC struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Region      string `json:"region"`

	// NOTE: IPv6 VPCs may not currently be available to all users.
	IPv6 []VPCIPv6Range `json:"ipv6"`

	Subnets []VPCSubnet `json:"subnets"`
	Created *time.Time  `json:"-"`
	Updated *time.Time  `json:"-"`
}

// VPCIPv6Range represents a single IPv6 range assigned to a VPC.
// NOTE: IPv6 VPCs may not currently be available to all users.
type VPCIPv6Range struct {
	Range string `json:"range"`
}

type VPCCreateOptions struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Region      string `json:"region"`

	// NOTE: IPv6 VPCs may not currently be available to all users.
	IPv6 []VPCCreateOptionsIPv6 `json:"ipv6,omitempty"`

	Subnets []VPCSubnetCreateOptions `json:"subnets,omitempty"`
}

// VPCCreateOptionsIPv6 represents a single IPv6 range assigned to a VPC
// which is specified during a VPC's creation.
// NOTE: IPv6 VPCs may not currently be available to all users.
type VPCCreateOptionsIPv6 struct {
	Range           *string `json:"range,omitempty"`
	AllocationClass *string `json:"allocation_class,omitempty"`
}

type VPCUpdateOptions struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
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
		IPv6: mapSlice(v.IPv6, func(i VPCIPv6Range) VPCCreateOptionsIPv6 {
			return VPCCreateOptionsIPv6{
				Range: copyValue(&i.Range),
			}
		}),
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
	return doPOSTRequest[VPC](ctx, c, "vpcs", opts)
}

func (c *Client) GetVPC(ctx context.Context, vpcID int) (*VPC, error) {
	e := formatAPIPath("/vpcs/%d", vpcID)
	return doGETRequest[VPC](ctx, c, e)
}

func (c *Client) ListVPCs(ctx context.Context, opts *ListOptions) ([]VPC, error) {
	return getPaginatedResults[VPC](ctx, c, "vpcs", opts)
}

func (c *Client) UpdateVPC(
	ctx context.Context,
	vpcID int,
	opts VPCUpdateOptions,
) (*VPC, error) {
	e := formatAPIPath("vpcs/%d", vpcID)
	return doPUTRequest[VPC](ctx, c, e, opts)
}

func (c *Client) DeleteVPC(ctx context.Context, vpcID int) error {
	e := formatAPIPath("vpcs/%d", vpcID)
	return doDELETERequest(ctx, c, e)
}
