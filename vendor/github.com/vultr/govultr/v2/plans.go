package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// PlanService is the interface to interact with the Plans endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/plans
type PlanService interface {
	List(ctx context.Context, planType string, options *ListOptions) ([]Plan, *Meta, error)
	ListBareMetal(ctx context.Context, options *ListOptions) ([]BareMetalPlan, *Meta, error)
}

// PlanServiceHandler handles interaction with the Plans methods for the Vultr API
type PlanServiceHandler struct {
	client *Client
}

// BareMetalPlan represents bare metal plans
type BareMetalPlan struct {
	ID          string   `json:"id"`
	CPUCount    int      `json:"cpu_count"`
	CPUModel    string   `json:"cpu_model"`
	CPUThreads  int      `json:"cpu_threads"`
	RAM         int      `json:"ram"`
	Disk        int      `json:"disk"`
	DiskCount   int      `json:"disk_count"`
	Bandwidth   int      `json:"bandwidth"`
	MonthlyCost float32  `json:"monthly_cost"`
	Type        string   `json:"type"`
	Locations   []string `json:"locations"`
}

// Plan represents vc2, vdc, or vhf
type Plan struct {
	ID          string   `json:"id"`
	VCPUCount   int      `json:"vcpu_count"`
	RAM         int      `json:"ram"`
	Disk        int      `json:"disk"`
	DiskCount   int      `json:"disk_count"`
	Bandwidth   int      `json:"bandwidth"`
	MonthlyCost float32  `json:"monthly_cost"`
	Type        string   `json:"type"`
	Locations   []string `json:"locations"`
}

type plansBase struct {
	Plans []Plan `json:"plans"`
	Meta  *Meta  `json:"meta"`
}

type bareMetalPlansBase struct {
	Plans []BareMetalPlan `json:"plans_metal"`
	Meta  *Meta           `json:"meta"`
}

// List retrieves a list of all active plans.
// planType is optional - pass an empty string to get all plans
func (p *PlanServiceHandler) List(ctx context.Context, planType string, options *ListOptions) ([]Plan, *Meta, error) {
	uri := "/v2/plans"

	req, err := p.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	if planType != "" {
		newValues.Add("type", planType)
	}

	req.URL.RawQuery = newValues.Encode()

	plans := new(plansBase)
	if err = p.client.DoWithContext(ctx, req, plans); err != nil {
		return nil, nil, err
	}

	return plans.Plans, plans.Meta, nil
}

// ListBareMetal all active bare metal plans.
func (p *PlanServiceHandler) ListBareMetal(ctx context.Context, options *ListOptions) ([]BareMetalPlan, *Meta, error) {
	uri := "/v2/plans-metal"

	req, err := p.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	bmPlans := new(bareMetalPlansBase)
	if err = p.client.DoWithContext(ctx, req, bmPlans); err != nil {
		return nil, nil, err
	}

	return bmPlans.Plans, bmPlans.Meta, nil
}
