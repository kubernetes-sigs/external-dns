package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// RegionService is the interface to interact with Region endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/region
type RegionService interface {
	Availability(ctx context.Context, regionID string, planType string) (*PlanAvailability, error)
	List(ctx context.Context, options *ListOptions) ([]Region, *Meta, error)
}

var _ RegionService = &RegionServiceHandler{}

// RegionServiceHandler handles interaction with the region methods for the Vultr API
type RegionServiceHandler struct {
	Client *Client
}

// Region represents a Vultr region
type Region struct {
	ID        string   `json:"id"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Continent string   `json:"continent,omitempty"`
	Options   []string `json:"options"`
}

type regionBase struct {
	Regions []Region `json:"regions"`
	Meta    *Meta
}

// PlanAvailability contains all available plans.
type PlanAvailability struct {
	AvailablePlans []string `json:"available_plans"`
}

// List returns all available regions
func (r *RegionServiceHandler) List(ctx context.Context, options *ListOptions) ([]Region, *Meta, error) {
	uri := "/v2/regions"

	req, err := r.Client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	regions := new(regionBase)
	if err = r.Client.DoWithContext(ctx, req, &regions); err != nil {
		return nil, nil, err
	}

	return regions.Regions, regions.Meta, nil
}

// Availability retrieves a list of the plan IDs currently available for a given location.
func (r *RegionServiceHandler) Availability(ctx context.Context, regionID string, planType string) (*PlanAvailability, error) {
	uri := fmt.Sprintf("/v2/regions/%s/availability", regionID)

	req, err := r.Client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	// Optional planType filter
	if planType != "" {
		q := req.URL.Query()
		q.Add("type", planType)
		req.URL.RawQuery = q.Encode()
	}

	plans := new(PlanAvailability)
	if err = r.Client.DoWithContext(ctx, req, plans); err != nil {
		return nil, err
	}

	return plans, nil
}
