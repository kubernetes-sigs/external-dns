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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// PlanService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	return
}

// Lists available plans the zone can subscribe to.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AvailableRatePlan], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_plans", query.ZoneID)
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

// Lists available plans the zone can subscribe to.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AvailableRatePlan] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Details of the available plan that the zone can subscribe to.
func (r *PlanService) Get(ctx context.Context, planIdentifier string, query PlanGetParams, opts ...option.RequestOption) (res *AvailableRatePlan, err error) {
	var env PlanGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if planIdentifier == "" {
		err = errors.New("missing required plan_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_plans/%s", query.ZoneID, planIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailableRatePlan struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency AvailableRatePlanFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64               `json:"price"`
	JSON  availableRatePlanJSON `json:"-"`
}

// availableRatePlanJSON contains the JSON metadata for the struct
// [AvailableRatePlan]
type availableRatePlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AvailableRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableRatePlanJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type AvailableRatePlanFrequency string

const (
	AvailableRatePlanFrequencyWeekly    AvailableRatePlanFrequency = "weekly"
	AvailableRatePlanFrequencyMonthly   AvailableRatePlanFrequency = "monthly"
	AvailableRatePlanFrequencyQuarterly AvailableRatePlanFrequency = "quarterly"
	AvailableRatePlanFrequencyYearly    AvailableRatePlanFrequency = "yearly"
)

func (r AvailableRatePlanFrequency) IsKnown() bool {
	switch r {
	case AvailableRatePlanFrequencyWeekly, AvailableRatePlanFrequencyMonthly, AvailableRatePlanFrequencyQuarterly, AvailableRatePlanFrequencyYearly:
		return true
	}
	return false
}

type PlanListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PlanGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PlanGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AvailableRatePlan     `json:"result,required"`
	// Whether the API call was successful
	Success PlanGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    planGetResponseEnvelopeJSON    `json:"-"`
}

// planGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PlanGetResponseEnvelope]
type planGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PlanGetResponseEnvelopeSuccess bool

const (
	PlanGetResponseEnvelopeSuccessTrue PlanGetResponseEnvelopeSuccess = true
)

func (r PlanGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PlanGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
