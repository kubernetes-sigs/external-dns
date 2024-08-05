package linodego

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Region represents a linode region object
type RegionAvailability struct {
	Region    string `json:"region"`
	Plan      string `json:"plan"`
	Available bool   `json:"available"`
}

// RegionsAvailabilityPagedResponse represents a linode API response for listing
type RegionsAvailabilityPagedResponse struct {
	*PageOptions
	Data []RegionAvailability `json:"data"`
}

// endpoint gets the endpoint URL for Region
func (RegionsAvailabilityPagedResponse) endpoint(_ ...any) string {
	return "regions/availability"
}

func (resp *RegionsAvailabilityPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(RegionsAvailabilityPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*RegionsAvailabilityPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListRegionsAvailability lists Regions. This endpoint is cached by default.
func (c *Client) ListRegionsAvailability(ctx context.Context, opts *ListOptions) ([]RegionAvailability, error) {
	response := RegionsAvailabilityPagedResponse{}

	endpoint, err := generateListCacheURL(response.endpoint(), opts)
	if err != nil {
		return nil, err
	}

	if result := c.getCachedResponse(endpoint); result != nil {
		return result.([]RegionAvailability), nil
	}

	err = c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(endpoint, response.Data, &cacheExpiryTime)

	return response.Data, nil
}

// GetRegionAvailability gets the template with the provided ID. This endpoint is cached by default.
func (c *Client) GetRegionAvailability(ctx context.Context, regionID string) (*RegionAvailability, error) {
	e := fmt.Sprintf("regions/%s/availability", url.PathEscape(regionID))

	if result := c.getCachedResponse(e); result != nil {
		result := result.(RegionAvailability)
		return &result, nil
	}

	req := c.R(ctx).SetResult(&RegionAvailability{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(e, r.Result(), &cacheExpiryTime)

	return r.Result().(*RegionAvailability), nil
}
