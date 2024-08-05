package linodego

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// AccountAvailability returns the resources availability in a region to an account.
type AccountAvailability struct {
	// region id
	Region string `json:"region"`

	// the unavailable resources in a region to the customer
	Unavailable []string `json:"unavailable"`

	// the available resources in a region to the customer
	Available []string `json:"available"`
}

// AccountAvailabilityPagedResponse represents a paginated Account Availability API response
type AccountAvailabilityPagedResponse struct {
	*PageOptions
	Data []AccountAvailability `json:"data"`
}

// endpoint gets the endpoint URL for AccountAvailability
func (AccountAvailabilityPagedResponse) endpoint(_ ...any) string {
	return "/account/availability"
}

func (resp *AccountAvailabilityPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(AccountAvailabilityPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*AccountAvailabilityPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListAccountAvailabilities lists all regions and the resource availabilities to the account.
func (c *Client) ListAccountAvailabilities(ctx context.Context, opts *ListOptions) ([]AccountAvailability, error) {
	response := AccountAvailabilityPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetAccountAvailability gets the resources availability in a region to the customer.
func (c *Client) GetAccountAvailability(ctx context.Context, regionID string) (*AccountAvailability, error) {
	req := c.R(ctx).SetResult(&AccountAvailability{})
	regionID = url.PathEscape(regionID)
	b := fmt.Sprintf("account/availability/%s", regionID)
	r, err := coupleAPIErrors(req.Get(b))
	if err != nil {
		return nil, err
	}

	return r.Result().(*AccountAvailability), nil
}
