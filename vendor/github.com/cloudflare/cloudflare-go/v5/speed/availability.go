// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// AvailabilityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailabilityService] method instead.
type AvailabilityService struct {
	Options []option.RequestOption
}

// NewAvailabilityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailabilityService(opts ...option.RequestOption) (r *AvailabilityService) {
	r = &AvailabilityService{}
	r.Options = opts
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *AvailabilityService) List(ctx context.Context, query AvailabilityListParams, opts ...option.RequestOption) (res *Availability, err error) {
	var env AvailabilityListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Availability struct {
	Quota   AvailabilityQuota `json:"quota"`
	Regions []LabeledRegion   `json:"regions"`
	// Available regions.
	RegionsPerPlan AvailabilityRegionsPerPlan `json:"regionsPerPlan"`
	JSON           availabilityJSON           `json:"-"`
}

// availabilityJSON contains the JSON metadata for the struct [Availability]
type availabilityJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Availability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityJSON) RawJSON() string {
	return r.raw
}

type AvailabilityQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan AvailabilityQuotaQuotasPerPlan `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan AvailabilityQuotaScheduleQuotasPerPlan `json:"scheduleQuotasPerPlan"`
	JSON                  availabilityQuotaJSON                  `json:"-"`
}

// availabilityQuotaJSON contains the JSON metadata for the struct
// [AvailabilityQuota]
type availabilityQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AvailabilityQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaJSON) RawJSON() string {
	return r.raw
}

// The number of tests available per plan.
type AvailabilityQuotaQuotasPerPlan struct {
	// Counts per account plan.
	Value AvailabilityQuotaQuotasPerPlanValue `json:"value"`
	JSON  availabilityQuotaQuotasPerPlanJSON  `json:"-"`
}

// availabilityQuotaQuotasPerPlanJSON contains the JSON metadata for the struct
// [AvailabilityQuotaQuotasPerPlan]
type availabilityQuotaQuotasPerPlanJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityQuotaQuotasPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaQuotasPerPlanJSON) RawJSON() string {
	return r.raw
}

// Counts per account plan.
type AvailabilityQuotaQuotasPerPlanValue struct {
	Business   int64                                   `json:"business"`
	Enterprise int64                                   `json:"enterprise"`
	Free       int64                                   `json:"free"`
	Pro        int64                                   `json:"pro"`
	JSON       availabilityQuotaQuotasPerPlanValueJSON `json:"-"`
}

// availabilityQuotaQuotasPerPlanValueJSON contains the JSON metadata for the
// struct [AvailabilityQuotaQuotasPerPlanValue]
type availabilityQuotaQuotasPerPlanValueJSON struct {
	Business    apijson.Field
	Enterprise  apijson.Field
	Free        apijson.Field
	Pro         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityQuotaQuotasPerPlanValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaQuotasPerPlanValueJSON) RawJSON() string {
	return r.raw
}

// The number of schedules available per plan.
type AvailabilityQuotaScheduleQuotasPerPlan struct {
	// Counts per account plan.
	Value AvailabilityQuotaScheduleQuotasPerPlanValue `json:"value"`
	JSON  availabilityQuotaScheduleQuotasPerPlanJSON  `json:"-"`
}

// availabilityQuotaScheduleQuotasPerPlanJSON contains the JSON metadata for the
// struct [AvailabilityQuotaScheduleQuotasPerPlan]
type availabilityQuotaScheduleQuotasPerPlanJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityQuotaScheduleQuotasPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaScheduleQuotasPerPlanJSON) RawJSON() string {
	return r.raw
}

// Counts per account plan.
type AvailabilityQuotaScheduleQuotasPerPlanValue struct {
	Business   int64                                           `json:"business"`
	Enterprise int64                                           `json:"enterprise"`
	Free       int64                                           `json:"free"`
	Pro        int64                                           `json:"pro"`
	JSON       availabilityQuotaScheduleQuotasPerPlanValueJSON `json:"-"`
}

// availabilityQuotaScheduleQuotasPerPlanValueJSON contains the JSON metadata for
// the struct [AvailabilityQuotaScheduleQuotasPerPlanValue]
type availabilityQuotaScheduleQuotasPerPlanValueJSON struct {
	Business    apijson.Field
	Enterprise  apijson.Field
	Free        apijson.Field
	Pro         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityQuotaScheduleQuotasPerPlanValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityQuotaScheduleQuotasPerPlanValueJSON) RawJSON() string {
	return r.raw
}

// Available regions.
type AvailabilityRegionsPerPlan struct {
	Business   []LabeledRegion                `json:"business"`
	Enterprise []LabeledRegion                `json:"enterprise"`
	Free       []LabeledRegion                `json:"free"`
	Pro        []LabeledRegion                `json:"pro"`
	JSON       availabilityRegionsPerPlanJSON `json:"-"`
}

// availabilityRegionsPerPlanJSON contains the JSON metadata for the struct
// [AvailabilityRegionsPerPlan]
type availabilityRegionsPerPlanJSON struct {
	Business    apijson.Field
	Enterprise  apijson.Field
	Free        apijson.Field
	Pro         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityRegionsPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityRegionsPerPlanJSON) RawJSON() string {
	return r.raw
}

type AvailabilityListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AvailabilityListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                 `json:"success,required"`
	Result  Availability                         `json:"result"`
	JSON    availabilityListResponseEnvelopeJSON `json:"-"`
}

// availabilityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailabilityListResponseEnvelope]
type availabilityListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availabilityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
