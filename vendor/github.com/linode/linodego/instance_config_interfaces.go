package linodego

import (
	"context"
	"encoding/json"
	"fmt"
)

// InstanceConfigInterface contains information about a configuration's network interface
type InstanceConfigInterface struct {
	ID          int                    `json:"id"`
	IPAMAddress string                 `json:"ipam_address"`
	Label       string                 `json:"label"`
	Purpose     ConfigInterfacePurpose `json:"purpose"`
	Primary     bool                   `json:"primary"`
	Active      bool                   `json:"active"`
	VPCID       *int                   `json:"vpc_id"`
	SubnetID    *int                   `json:"subnet_id"`
	IPv4        *VPCIPv4               `json:"ipv4"`
	IPRanges    []string               `json:"ip_ranges"`
}

type VPCIPv4 struct {
	VPC     string  `json:"vpc,omitempty"`
	NAT1To1 *string `json:"nat_1_1,omitempty"`
}

type InstanceConfigInterfaceCreateOptions struct {
	IPAMAddress string                 `json:"ipam_address,omitempty"`
	Label       string                 `json:"label,omitempty"`
	Purpose     ConfigInterfacePurpose `json:"purpose,omitempty"`
	Primary     bool                   `json:"primary,omitempty"`
	SubnetID    *int                   `json:"subnet_id,omitempty"`
	IPv4        *VPCIPv4               `json:"ipv4,omitempty"`
	IPRanges    []string               `json:"ip_ranges,omitempty"`
}

type InstanceConfigInterfaceUpdateOptions struct {
	Primary  bool      `json:"primary,omitempty"`
	IPv4     *VPCIPv4  `json:"ipv4,omitempty"`
	IPRanges *[]string `json:"ip_ranges,omitempty"`
}

type InstanceConfigInterfacesReorderOptions struct {
	IDs []int `json:"ids"`
}

func getInstanceConfigInterfacesCreateOptionsList(
	interfaces []InstanceConfigInterface,
) []InstanceConfigInterfaceCreateOptions {
	interfaceOptsList := make([]InstanceConfigInterfaceCreateOptions, len(interfaces))
	for index, configInterface := range interfaces {
		interfaceOptsList[index] = configInterface.GetCreateOptions()
	}
	return interfaceOptsList
}

func (i InstanceConfigInterface) GetCreateOptions() InstanceConfigInterfaceCreateOptions {
	opts := InstanceConfigInterfaceCreateOptions{
		Label:    i.Label,
		Purpose:  i.Purpose,
		Primary:  i.Primary,
		SubnetID: i.SubnetID,
	}

	if len(i.IPRanges) > 0 {
		opts.IPRanges = i.IPRanges
	}

	if i.Purpose == InterfacePurposeVPC && i.IPv4 != nil {
		opts.IPv4 = &VPCIPv4{
			VPC:     i.IPv4.VPC,
			NAT1To1: i.IPv4.NAT1To1,
		}
	}

	opts.IPAMAddress = i.IPAMAddress

	return opts
}

func (i InstanceConfigInterface) GetUpdateOptions() InstanceConfigInterfaceUpdateOptions {
	opts := InstanceConfigInterfaceUpdateOptions{
		Primary: i.Primary,
	}

	if i.Purpose == InterfacePurposeVPC && i.IPv4 != nil {
		opts.IPv4 = &VPCIPv4{
			VPC:     i.IPv4.VPC,
			NAT1To1: i.IPv4.NAT1To1,
		}
	}

	if i.IPRanges != nil {
		// Copy the slice to prevent accidental
		// mutations
		copiedIPRanges := make([]string, len(i.IPRanges))
		copy(copiedIPRanges, i.IPRanges)

		opts.IPRanges = &copiedIPRanges
	}

	return opts
}

func (c *Client) AppendInstanceConfigInterface(
	ctx context.Context,
	linodeID int,
	configID int,
	opts InstanceConfigInterfaceCreateOptions,
) (*InstanceConfigInterface, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceConfigInterface{}).SetBody(string(body))
	e := fmt.Sprintf("/linode/instances/%d/configs/%d/interfaces", linodeID, configID)
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceConfigInterface), nil
}

func (c *Client) GetInstanceConfigInterface(
	ctx context.Context,
	linodeID int,
	configID int,
	interfaceID int,
) (*InstanceConfigInterface, error) {
	e := fmt.Sprintf(
		"linode/instances/%d/configs/%d/interfaces/%d",
		linodeID,
		configID,
		interfaceID,
	)
	req := c.R(ctx).SetResult(&InstanceConfigInterface{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceConfigInterface), nil
}

func (c *Client) ListInstanceConfigInterfaces(
	ctx context.Context,
	linodeID int,
	configID int,
) ([]InstanceConfigInterface, error) {
	e := fmt.Sprintf(
		"linode/instances/%d/configs/%d/interfaces",
		linodeID,
		configID,
	)
	req := c.R(ctx).SetResult([]InstanceConfigInterface{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return *r.Result().(*[]InstanceConfigInterface), nil
}

func (c *Client) UpdateInstanceConfigInterface(
	ctx context.Context,
	linodeID int,
	configID int,
	interfaceID int,
	opts InstanceConfigInterfaceUpdateOptions,
) (*InstanceConfigInterface, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf(
		"linode/instances/%d/configs/%d/interfaces/%d",
		linodeID,
		configID,
		interfaceID,
	)
	req := c.R(ctx).SetResult(&InstanceConfigInterface{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceConfigInterface), nil
}

func (c *Client) DeleteInstanceConfigInterface(
	ctx context.Context,
	linodeID int,
	configID int,
	interfaceID int,
) error {
	e := fmt.Sprintf(
		"linode/instances/%d/configs/%d/interfaces/%d",
		linodeID,
		configID,
		interfaceID,
	)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}

func (c *Client) ReorderInstanceConfigInterfaces(
	ctx context.Context,
	linodeID int,
	configID int,
	opts InstanceConfigInterfacesReorderOptions,
) error {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	e := fmt.Sprintf(
		"linode/instances/%d/configs/%d/interfaces/order",
		linodeID,
		configID,
	)

	req := c.R(ctx).SetBody(string(body))
	_, err = coupleAPIErrors(req.Post(e))

	return err
}
