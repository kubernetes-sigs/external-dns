package linodego

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// NodeBalancerFirewallsPagedResponse represents a Linode API response for listing of Cloud Firewalls
type NodeBalancerFirewallsPagedResponse struct {
	*PageOptions
	Data []Firewall `json:"data"`
}

func (NodeBalancerFirewallsPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("nodebalancers/%d/firewalls", id)
}

func (resp *NodeBalancerFirewallsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(NodeBalancerFirewallsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*NodeBalancerFirewallsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListNodeBalancerFirewalls returns a paginated list of Cloud Firewalls for nodebalancerID
func (c *Client) ListNodeBalancerFirewalls(ctx context.Context, nodebalancerID int, opts *ListOptions) ([]Firewall, error) {
	response := NodeBalancerFirewallsPagedResponse{}

	err := c.listHelper(ctx, &response, opts, nodebalancerID)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
