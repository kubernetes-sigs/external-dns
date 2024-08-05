package linodego

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// InstanceFirewallsPagedResponse represents a Linode API response for listing of Cloud Firewalls
type InstanceFirewallsPagedResponse struct {
	*PageOptions
	Data []Firewall `json:"data"`
}

func (InstanceFirewallsPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("linode/instances/%d/firewalls", id)
}

func (resp *InstanceFirewallsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(InstanceFirewallsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*InstanceFirewallsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListInstanceFirewalls returns a paginated list of Cloud Firewalls for linodeID
func (c *Client) ListInstanceFirewalls(ctx context.Context, linodeID int, opts *ListOptions) ([]Firewall, error) {
	response := InstanceFirewallsPagedResponse{}

	err := c.listHelper(ctx, &response, opts, linodeID)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
