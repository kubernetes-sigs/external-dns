// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) Get(ctx context.Context, query RatePlanGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[RatePlanGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_rate_plans", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) GetAutoPaging(ctx context.Context, query RatePlanGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RatePlanGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type RatePlanGetResponse struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []RatePlanGetResponseComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency RatePlanGetResponseFrequency `json:"frequency"`
	// The plan name.
	Name string                  `json:"name"`
	JSON ratePlanGetResponseJSON `json:"-"`
}

// ratePlanGetResponseJSON contains the JSON metadata for the struct
// [RatePlanGetResponse]
type ratePlanGetResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseJSON) RawJSON() string {
	return r.raw
}

type RatePlanGetResponseComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name RatePlanGetResponseComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                          `json:"unit_price"`
	JSON      ratePlanGetResponseComponentJSON `json:"-"`
}

// ratePlanGetResponseComponentJSON contains the JSON metadata for the struct
// [RatePlanGetResponseComponent]
type ratePlanGetResponseComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseComponentJSON) RawJSON() string {
	return r.raw
}

// The unique component.
type RatePlanGetResponseComponentsName string

const (
	RatePlanGetResponseComponentsNameZones                       RatePlanGetResponseComponentsName = "zones"
	RatePlanGetResponseComponentsNamePageRules                   RatePlanGetResponseComponentsName = "page_rules"
	RatePlanGetResponseComponentsNameDedicatedCertificates       RatePlanGetResponseComponentsName = "dedicated_certificates"
	RatePlanGetResponseComponentsNameDedicatedCertificatesCustom RatePlanGetResponseComponentsName = "dedicated_certificates_custom"
)

func (r RatePlanGetResponseComponentsName) IsKnown() bool {
	switch r {
	case RatePlanGetResponseComponentsNameZones, RatePlanGetResponseComponentsNamePageRules, RatePlanGetResponseComponentsNameDedicatedCertificates, RatePlanGetResponseComponentsNameDedicatedCertificatesCustom:
		return true
	}
	return false
}

// The frequency at which you will be billed for this plan.
type RatePlanGetResponseFrequency string

const (
	RatePlanGetResponseFrequencyWeekly    RatePlanGetResponseFrequency = "weekly"
	RatePlanGetResponseFrequencyMonthly   RatePlanGetResponseFrequency = "monthly"
	RatePlanGetResponseFrequencyQuarterly RatePlanGetResponseFrequency = "quarterly"
	RatePlanGetResponseFrequencyYearly    RatePlanGetResponseFrequency = "yearly"
)

func (r RatePlanGetResponseFrequency) IsKnown() bool {
	switch r {
	case RatePlanGetResponseFrequencyWeekly, RatePlanGetResponseFrequencyMonthly, RatePlanGetResponseFrequencyQuarterly, RatePlanGetResponseFrequencyYearly:
		return true
	}
	return false
}

type RatePlanGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
