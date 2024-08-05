package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

<<<<<<< HEAD
	"github.com/linode/linodego/internal/parseabletime"
)

type VLAN struct {
	Label   string     `json:"label"`
	Linodes []int      `json:"linodes"`
	Region  string     `json:"region"`
	Created *time.Time `json:"-"`
}

// UnmarshalJSON for VLAN responses
func (v *VLAN) UnmarshalJSON(b []byte) error {
	type Mask VLAN

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(v),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	v.Created = (*time.Time)(p.Created)
	return nil
}

// VLANsPagedResponse represents a Linode API response for listing of VLANs
type VLANsPagedResponse struct {
	*PageOptions
	Data []VLAN `json:"data"`
}

func (VLANsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.VLANs.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

func (resp *VLANsPagedResponse) appendData(r *VLANsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListVLANs returns a paginated list of VLANs
func (c *Client) ListVLANs(ctx context.Context, opts *ListOptions) ([]VLAN, error) {
	response := VLANsPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetVLANIPAMAddress returns the IPAM Address for a given VLAN Label as a string (10.0.0.1/24)
func (c *Client) GetVLANIPAMAddress(ctx context.Context, linodeID int, vlanLabel string) (string, error) {
	f := Filter{}
	f.AddField(Eq, "interfaces", vlanLabel)
	vlanFilter, err := f.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("Unable to convert VLAN label: %s to a filterable object: %s", vlanLabel, err)
	}

	cfgs, err := c.ListInstanceConfigs(ctx, linodeID, &ListOptions{Filter: string(vlanFilter)})
	if err != nil {
		return "", fmt.Errorf("Fetching configs for instance %v failed: %s", linodeID, err)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

type VLAN struct {
	Label   string     `json:"label"`
	Linodes []int      `json:"linodes"`
	Region  string     `json:"region"`
	Created *time.Time `json:"-"`
}

// UnmarshalJSON for VLAN responses
func (v *VLAN) UnmarshalJSON(b []byte) error {
	type Mask VLAN

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(v),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	v.Created = (*time.Time)(p.Created)
	return nil
}

// VLANsPagedResponse represents a Linode API response for listing of VLANs
type VLANsPagedResponse struct {
	*PageOptions
	Data []VLAN `json:"data"`
}

func (VLANsPagedResponse) endpoint(_ ...any) string {
	return "networking/vlans"
}

func (resp *VLANsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(VLANsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*VLANsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListVLANs returns a paginated list of VLANs
func (c *Client) ListVLANs(ctx context.Context, opts *ListOptions) ([]VLAN, error) {
	response := VLANsPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetVLANIPAMAddress returns the IPAM Address for a given VLAN Label as a string (10.0.0.1/24)
func (c *Client) GetVLANIPAMAddress(ctx context.Context, linodeID int, vlanLabel string) (string, error) {
	f := Filter{}
	f.AddField(Eq, "interfaces", vlanLabel)
	vlanFilter, err := f.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("Unable to convert VLAN label: %s to a filterable object: %w", vlanLabel, err)
	}

	cfgs, err := c.ListInstanceConfigs(ctx, linodeID, &ListOptions{Filter: string(vlanFilter)})
	if err != nil {
		return "", fmt.Errorf("Fetching configs for instance %v failed: %w", linodeID, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	interfaces := cfgs[0].Interfaces
	for _, face := range interfaces {
		if face.Label == vlanLabel {
			return face.IPAMAddress, nil
		}
	}

	return "", fmt.Errorf("Failed to find IPAMAddress for VLAN: %s", vlanLabel)
}
