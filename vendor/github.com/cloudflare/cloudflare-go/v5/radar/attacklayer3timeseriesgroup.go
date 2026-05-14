// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AttackLayer3TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3TimeseriesGroupService] method instead.
type AttackLayer3TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAttackLayer3TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *AttackLayer3TimeseriesGroupService) {
	r = &AttackLayer3TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 3 attacks by bitrate over time.
func (r *AttackLayer3TimeseriesGroupService) Bitrate(ctx context.Context, query AttackLayer3TimeseriesGroupBitrateParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupBitrateResponse, err error) {
	var env AttackLayer3TimeseriesGroupBitrateResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by duration over time.
func (r *AttackLayer3TimeseriesGroupService) Duration(ctx context.Context, query AttackLayer3TimeseriesGroupDurationParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupDurationResponse, err error) {
	var env AttackLayer3TimeseriesGroupDurationResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by targeted industry over time.
func (r *AttackLayer3TimeseriesGroupService) Industry(ctx context.Context, query AttackLayer3TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupIndustryResponse, err error) {
	var env AttackLayer3TimeseriesGroupIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by IP version over time.
func (r *AttackLayer3TimeseriesGroupService) IPVersion(ctx context.Context, query AttackLayer3TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupIPVersionResponse, err error) {
	var env AttackLayer3TimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by protocol over time.
func (r *AttackLayer3TimeseriesGroupService) Protocol(ctx context.Context, query AttackLayer3TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupProtocolResponse, err error) {
	var env AttackLayer3TimeseriesGroupProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by vector over time.
func (r *AttackLayer3TimeseriesGroupService) Vector(ctx context.Context, query AttackLayer3TimeseriesGroupVectorParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupVectorResponse, err error) {
	var env AttackLayer3TimeseriesGroupVectorResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by targeted vertical over time.
func (r *AttackLayer3TimeseriesGroupService) Vertical(ctx context.Context, query AttackLayer3TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupVerticalResponse, err error) {
	var env AttackLayer3TimeseriesGroupVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TimeseriesGroupBitrateResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupBitrateResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupBitrateResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupBitrateResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupBitrateResponse]
type attackLayer3TimeseriesGroupBitrateResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupBitrateResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupBitrateResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupBitrateResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupBitrateResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupBitrateResponseMeta]
type attackLayer3TimeseriesGroupBitrateResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupBitrateResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupBitrateResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupBitrateResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupBitrateResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupBitrateResponseMetaDateRange]
type attackLayer3TimeseriesGroupBitrateResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupBitrateResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupBitrateResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupBitrateResponseMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupBitrateResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupBitrateResponseMetaUnit]
type attackLayer3TimeseriesGroupBitrateResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupBitrateResponseSerie0 struct {
	OneGBPSToTenGBPS         []string                                             `json:"_1_GBPS_TO_10_GBPS,required"`
	TenGBPSToOneHundredGBPS  []string                                             `json:"_10_GBPS_TO_100_GBPS,required"`
	FiveHundredMBPSToOneGBPS []string                                             `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100GBPS              []string                                             `json:"OVER_100_GBPS,required"`
	Timestamps               []time.Time                                          `json:"timestamps,required" format:"date-time"`
	Under500MBPS             []string                                             `json:"UNDER_500_MBPS,required"`
	JSON                     attackLayer3TimeseriesGroupBitrateResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupBitrateResponseSerie0]
type attackLayer3TimeseriesGroupBitrateResponseSerie0JSON struct {
	OneGBPSToTenGBPS         apijson.Field
	TenGBPSToOneHundredGBPS  apijson.Field
	FiveHundredMBPSToOneGBPS apijson.Field
	Over100GBPS              apijson.Field
	Timestamps               apijson.Field
	Under500MBPS             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupDurationResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupDurationResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupDurationResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupDurationResponse]
type attackLayer3TimeseriesGroupDurationResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupDurationResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupDurationResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupDurationResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupDurationResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupDurationResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupDurationResponseMeta]
type attackLayer3TimeseriesGroupDurationResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupDurationResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupDurationResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupDurationResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupDurationResponseMetaDateRange]
type attackLayer3TimeseriesGroupDurationResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupDurationResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupDurationResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupDurationResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupDurationResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupDurationResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupDurationResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupDurationResponseMetaUnit]
type attackLayer3TimeseriesGroupDurationResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationResponseSerie0 struct {
	OneHourToThreeHours   []string                                              `json:"_1_HOUR_TO_3_HOURS,required"`
	TenMinsToTwentyMins   []string                                              `json:"_10_MINS_TO_20_MINS,required"`
	TwentyMinsToFortyMins []string                                              `json:"_20_MINS_TO_40_MINS,required"`
	FortyMinsToOneHour    []string                                              `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours            []string                                              `json:"OVER_3_HOURS,required"`
	Timestamps            []time.Time                                           `json:"timestamps,required" format:"date-time"`
	Under10Mins           []string                                              `json:"UNDER_10_MINS,required"`
	JSON                  attackLayer3TimeseriesGroupDurationResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupDurationResponseSerie0]
type attackLayer3TimeseriesGroupDurationResponseSerie0JSON struct {
	OneHourToThreeHours   apijson.Field
	TenMinsToTwentyMins   apijson.Field
	TwentyMinsToFortyMins apijson.Field
	FortyMinsToOneHour    apijson.Field
	Over3Hours            apijson.Field
	Timestamps            apijson.Field
	Under10Mins           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupIndustryResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupIndustryResponse]
type attackLayer3TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupIndustryResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupIndustryResponseMeta]
type attackLayer3TimeseriesGroupIndustryResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupIndustryResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupIndustryResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupIndustryResponseMetaDateRange]
type attackLayer3TimeseriesGroupIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIndustryResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIndustryResponseMetaUnit]
type attackLayer3TimeseriesGroupIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []time.Time                                           `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupIndustryResponseSerie0]
type attackLayer3TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupIPVersionResponse]
type attackLayer3TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupIPVersionResponseMeta]
type attackLayer3TimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupIPVersionResponseMetaDateRange]
type attackLayer3TimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIPVersionResponseMetaUnit]
type attackLayer3TimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                               `json:"IPv4,required"`
	IPv6       []string                                               `json:"IPv6,required"`
	Timestamps []time.Time                                            `json:"timestamps,required" format:"date-time"`
	JSON       attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIPVersionResponseSerie0]
type attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupProtocolResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupProtocolResponse]
type attackLayer3TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupProtocolResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupProtocolResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupProtocolResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupProtocolResponseMeta]
type attackLayer3TimeseriesGroupProtocolResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupProtocolResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupProtocolResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupProtocolResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupProtocolResponseMetaDateRange]
type attackLayer3TimeseriesGroupProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupProtocolResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupProtocolResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupProtocolResponseMetaUnit]
type attackLayer3TimeseriesGroupProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolResponseSerie0 struct {
	GRE        []string                                              `json:"GRE,required"`
	Icmp       []string                                              `json:"ICMP,required"`
	TCP        []string                                              `json:"TCP,required"`
	Timestamps []time.Time                                           `json:"timestamps,required" format:"date-time"`
	Udp        []string                                              `json:"UDP,required"`
	JSON       attackLayer3TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupProtocolResponseSerie0]
type attackLayer3TimeseriesGroupProtocolResponseSerie0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	TCP         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupVectorResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupVectorResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupVectorResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesGroupVectorResponse]
type attackLayer3TimeseriesGroupVectorResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupVectorResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupVectorResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupVectorResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupVectorResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupVectorResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupVectorResponseMeta]
type attackLayer3TimeseriesGroupVectorResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupVectorResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupVectorResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupVectorResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupVectorResponseMetaDateRange]
type attackLayer3TimeseriesGroupVectorResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVectorResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupVectorResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupVectorResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupVectorResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVectorResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupVectorResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseMetaUnitJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVectorResponseMetaUnit]
type attackLayer3TimeseriesGroupVectorResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorResponseSerie0 struct {
	Timestamps  []time.Time                                         `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                 `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupVectorResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVectorResponseSerie0]
type attackLayer3TimeseriesGroupVectorResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalResponse struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupVerticalResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupVerticalResponse]
type attackLayer3TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupVerticalResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVerticalResponseMeta]
type attackLayer3TimeseriesGroupVerticalResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupVerticalResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupVerticalResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupVerticalResponseMetaDateRange]
type attackLayer3TimeseriesGroupVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVerticalResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupVerticalResponseMetaUnit]
type attackLayer3TimeseriesGroupVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []time.Time                                           `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVerticalResponseSerie0]
type attackLayer3TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupBitrateParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupBitrateParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupBitrateParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupBitrateParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupBitrateParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupBitrateParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupBitrateParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupBitrateParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval15m AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1h  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1d  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupBitrateParamsAggInterval1w  AttackLayer3TimeseriesGroupBitrateParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsAggInterval15m, AttackLayer3TimeseriesGroupBitrateParamsAggInterval1h, AttackLayer3TimeseriesGroupBitrateParamsAggInterval1d, AttackLayer3TimeseriesGroupBitrateParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupBitrateParamsDirection string

const (
	AttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin AttackLayer3TimeseriesGroupBitrateParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupBitrateParamsDirectionTarget AttackLayer3TimeseriesGroupBitrateParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsDirectionOrigin, AttackLayer3TimeseriesGroupBitrateParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupBitrateParamsFormat string

const (
	AttackLayer3TimeseriesGroupBitrateParamsFormatJson AttackLayer3TimeseriesGroupBitrateParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupBitrateParamsFormatCsv  AttackLayer3TimeseriesGroupBitrateParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsFormatJson, AttackLayer3TimeseriesGroupBitrateParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupBitrateParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4 AttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv6 AttackLayer3TimeseriesGroupBitrateParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv4, AttackLayer3TimeseriesGroupBitrateParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupBitrateParamsNormalization string

const (
	AttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage AttackLayer3TimeseriesGroupBitrateParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupBitrateParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupBitrateParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsNormalizationPercentage, AttackLayer3TimeseriesGroupBitrateParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupBitrateParamsProtocol string

const (
	AttackLayer3TimeseriesGroupBitrateParamsProtocolUdp  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolTCP  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolIcmp AttackLayer3TimeseriesGroupBitrateParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupBitrateParamsProtocolGRE  AttackLayer3TimeseriesGroupBitrateParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupBitrateParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupBitrateParamsProtocolUdp, AttackLayer3TimeseriesGroupBitrateParamsProtocolTCP, AttackLayer3TimeseriesGroupBitrateParamsProtocolIcmp, AttackLayer3TimeseriesGroupBitrateParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupBitrateResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupBitrateResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupBitrateResponseEnvelope]
type attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupBitrateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupDurationParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupDurationParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupDurationParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupDurationParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupDurationParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupDurationParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupDurationParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupDurationParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupDurationParamsAggInterval15m AttackLayer3TimeseriesGroupDurationParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1h  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1d  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupDurationParamsAggInterval1w  AttackLayer3TimeseriesGroupDurationParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupDurationParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsAggInterval15m, AttackLayer3TimeseriesGroupDurationParamsAggInterval1h, AttackLayer3TimeseriesGroupDurationParamsAggInterval1d, AttackLayer3TimeseriesGroupDurationParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupDurationParamsDirection string

const (
	AttackLayer3TimeseriesGroupDurationParamsDirectionOrigin AttackLayer3TimeseriesGroupDurationParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupDurationParamsDirectionTarget AttackLayer3TimeseriesGroupDurationParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupDurationParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsDirectionOrigin, AttackLayer3TimeseriesGroupDurationParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupDurationParamsFormat string

const (
	AttackLayer3TimeseriesGroupDurationParamsFormatJson AttackLayer3TimeseriesGroupDurationParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupDurationParamsFormatCsv  AttackLayer3TimeseriesGroupDurationParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupDurationParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsFormatJson, AttackLayer3TimeseriesGroupDurationParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupDurationParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4 AttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv6 AttackLayer3TimeseriesGroupDurationParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupDurationParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv4, AttackLayer3TimeseriesGroupDurationParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupDurationParamsNormalization string

const (
	AttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage AttackLayer3TimeseriesGroupDurationParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupDurationParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupDurationParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupDurationParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsNormalizationPercentage, AttackLayer3TimeseriesGroupDurationParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupDurationParamsProtocol string

const (
	AttackLayer3TimeseriesGroupDurationParamsProtocolUdp  AttackLayer3TimeseriesGroupDurationParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolTCP  AttackLayer3TimeseriesGroupDurationParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolIcmp AttackLayer3TimeseriesGroupDurationParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupDurationParamsProtocolGRE  AttackLayer3TimeseriesGroupDurationParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupDurationParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupDurationParamsProtocolUdp, AttackLayer3TimeseriesGroupDurationParamsProtocolTCP, AttackLayer3TimeseriesGroupDurationParamsProtocolIcmp, AttackLayer3TimeseriesGroupDurationParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupDurationResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupDurationResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupDurationResponseEnvelope]
type attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupDurationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIndustryParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupIndustryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIndustryParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval15m AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1h  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1d  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupIndustryParamsAggInterval1w  AttackLayer3TimeseriesGroupIndustryParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsAggInterval15m, AttackLayer3TimeseriesGroupIndustryParamsAggInterval1h, AttackLayer3TimeseriesGroupIndustryParamsAggInterval1d, AttackLayer3TimeseriesGroupIndustryParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupIndustryParamsDirection string

const (
	AttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin AttackLayer3TimeseriesGroupIndustryParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupIndustryParamsDirectionTarget AttackLayer3TimeseriesGroupIndustryParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsDirectionOrigin, AttackLayer3TimeseriesGroupIndustryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupIndustryParamsFormat string

const (
	AttackLayer3TimeseriesGroupIndustryParamsFormatJson AttackLayer3TimeseriesGroupIndustryParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupIndustryParamsFormatCsv  AttackLayer3TimeseriesGroupIndustryParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsFormatJson, AttackLayer3TimeseriesGroupIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIndustryParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4 AttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv6 AttackLayer3TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv4, AttackLayer3TimeseriesGroupIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIndustryParamsNormalization string

const (
	AttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage AttackLayer3TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIndustryParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsNormalizationPercentage, AttackLayer3TimeseriesGroupIndustryParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIndustryParamsProtocol string

const (
	AttackLayer3TimeseriesGroupIndustryParamsProtocolUdp  AttackLayer3TimeseriesGroupIndustryParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupIndustryParamsProtocolTCP  AttackLayer3TimeseriesGroupIndustryParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupIndustryParamsProtocolIcmp AttackLayer3TimeseriesGroupIndustryParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupIndustryParamsProtocolGRE  AttackLayer3TimeseriesGroupIndustryParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupIndustryParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIndustryParamsProtocolUdp, AttackLayer3TimeseriesGroupIndustryParamsProtocolTCP, AttackLayer3TimeseriesGroupIndustryParamsProtocolIcmp, AttackLayer3TimeseriesGroupIndustryParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIndustryResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIndustryResponseEnvelope]
type attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupIPVersionParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupIPVersionParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1h  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1d  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1w  AttackLayer3TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionParamsAggInterval15m, AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1h, AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1d, AttackLayer3TimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupIPVersionParamsDirection string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin AttackLayer3TimeseriesGroupIPVersionParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupIPVersionParamsDirectionTarget AttackLayer3TimeseriesGroupIPVersionParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupIPVersionParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionParamsDirectionOrigin, AttackLayer3TimeseriesGroupIPVersionParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupIPVersionParamsFormat string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsFormatJson AttackLayer3TimeseriesGroupIPVersionParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupIPVersionParamsFormatCsv  AttackLayer3TimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionParamsFormatJson, AttackLayer3TimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupIPVersionParamsNormalization string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage AttackLayer3TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupIPVersionParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupIPVersionParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionParamsNormalizationPercentage, AttackLayer3TimeseriesGroupIPVersionParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIPVersionParamsProtocol string

const (
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolTCP  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolIcmp AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupIPVersionParamsProtocolGRE  AttackLayer3TimeseriesGroupIPVersionParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupIPVersionParamsProtocolUdp, AttackLayer3TimeseriesGroupIPVersionParamsProtocolTCP, AttackLayer3TimeseriesGroupIPVersionParamsProtocolIcmp, AttackLayer3TimeseriesGroupIPVersionParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupIPVersionResponseEnvelope]
type attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupProtocolParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupProtocolParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupProtocolParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupProtocolParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupProtocolParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupProtocolParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval15m AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1h  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1d  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupProtocolParamsAggInterval1w  AttackLayer3TimeseriesGroupProtocolParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolParamsAggInterval15m, AttackLayer3TimeseriesGroupProtocolParamsAggInterval1h, AttackLayer3TimeseriesGroupProtocolParamsAggInterval1d, AttackLayer3TimeseriesGroupProtocolParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupProtocolParamsDirection string

const (
	AttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin AttackLayer3TimeseriesGroupProtocolParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupProtocolParamsDirectionTarget AttackLayer3TimeseriesGroupProtocolParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupProtocolParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolParamsDirectionOrigin, AttackLayer3TimeseriesGroupProtocolParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupProtocolParamsFormat string

const (
	AttackLayer3TimeseriesGroupProtocolParamsFormatJson AttackLayer3TimeseriesGroupProtocolParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupProtocolParamsFormatCsv  AttackLayer3TimeseriesGroupProtocolParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolParamsFormatJson, AttackLayer3TimeseriesGroupProtocolParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupProtocolParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4 AttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv6 AttackLayer3TimeseriesGroupProtocolParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv4, AttackLayer3TimeseriesGroupProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupProtocolParamsNormalization string

const (
	AttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage AttackLayer3TimeseriesGroupProtocolParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupProtocolParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupProtocolParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupProtocolParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupProtocolParamsNormalizationPercentage, AttackLayer3TimeseriesGroupProtocolParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupProtocolResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupProtocolResponseEnvelope]
type attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVectorParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupVectorParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupVectorParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupVectorParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupVectorParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupVectorParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupVectorParams]'s query parameters
// as `url.Values`.
func (r AttackLayer3TimeseriesGroupVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVectorParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupVectorParamsAggInterval15m AttackLayer3TimeseriesGroupVectorParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1h  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1d  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupVectorParamsAggInterval1w  AttackLayer3TimeseriesGroupVectorParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupVectorParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsAggInterval15m, AttackLayer3TimeseriesGroupVectorParamsAggInterval1h, AttackLayer3TimeseriesGroupVectorParamsAggInterval1d, AttackLayer3TimeseriesGroupVectorParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupVectorParamsDirection string

const (
	AttackLayer3TimeseriesGroupVectorParamsDirectionOrigin AttackLayer3TimeseriesGroupVectorParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupVectorParamsDirectionTarget AttackLayer3TimeseriesGroupVectorParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupVectorParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsDirectionOrigin, AttackLayer3TimeseriesGroupVectorParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupVectorParamsFormat string

const (
	AttackLayer3TimeseriesGroupVectorParamsFormatJson AttackLayer3TimeseriesGroupVectorParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupVectorParamsFormatCsv  AttackLayer3TimeseriesGroupVectorParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupVectorParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsFormatJson, AttackLayer3TimeseriesGroupVectorParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVectorParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4 AttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv6 AttackLayer3TimeseriesGroupVectorParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupVectorParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv4, AttackLayer3TimeseriesGroupVectorParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVectorParamsNormalization string

const (
	AttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage AttackLayer3TimeseriesGroupVectorParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVectorParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupVectorParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupVectorParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsNormalizationPercentage, AttackLayer3TimeseriesGroupVectorParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVectorParamsProtocol string

const (
	AttackLayer3TimeseriesGroupVectorParamsProtocolUdp  AttackLayer3TimeseriesGroupVectorParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolTCP  AttackLayer3TimeseriesGroupVectorParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolIcmp AttackLayer3TimeseriesGroupVectorParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupVectorParamsProtocolGRE  AttackLayer3TimeseriesGroupVectorParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupVectorParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVectorParamsProtocolUdp, AttackLayer3TimeseriesGroupVectorParamsProtocolTCP, AttackLayer3TimeseriesGroupVectorParamsProtocolIcmp, AttackLayer3TimeseriesGroupVectorParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVectorResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupVectorResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupVectorResponseEnvelope]
type attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVectorResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupVerticalParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupVerticalParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupVerticalParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r AttackLayer3TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupVerticalParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval15m AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1h  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1d  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupVerticalParamsAggInterval1w  AttackLayer3TimeseriesGroupVerticalParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsAggInterval15m, AttackLayer3TimeseriesGroupVerticalParamsAggInterval1h, AttackLayer3TimeseriesGroupVerticalParamsAggInterval1d, AttackLayer3TimeseriesGroupVerticalParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupVerticalParamsDirection string

const (
	AttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin AttackLayer3TimeseriesGroupVerticalParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupVerticalParamsDirectionTarget AttackLayer3TimeseriesGroupVerticalParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsDirectionOrigin, AttackLayer3TimeseriesGroupVerticalParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupVerticalParamsFormat string

const (
	AttackLayer3TimeseriesGroupVerticalParamsFormatJson AttackLayer3TimeseriesGroupVerticalParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupVerticalParamsFormatCsv  AttackLayer3TimeseriesGroupVerticalParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsFormatJson, AttackLayer3TimeseriesGroupVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVerticalParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4 AttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv6 AttackLayer3TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv4, AttackLayer3TimeseriesGroupVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupVerticalParamsNormalization string

const (
	AttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage AttackLayer3TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupVerticalParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsNormalizationPercentage, AttackLayer3TimeseriesGroupVerticalParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVerticalParamsProtocol string

const (
	AttackLayer3TimeseriesGroupVerticalParamsProtocolUdp  AttackLayer3TimeseriesGroupVerticalParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupVerticalParamsProtocolTCP  AttackLayer3TimeseriesGroupVerticalParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupVerticalParamsProtocolIcmp AttackLayer3TimeseriesGroupVerticalParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupVerticalParamsProtocolGRE  AttackLayer3TimeseriesGroupVerticalParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupVerticalParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupVerticalParamsProtocolUdp, AttackLayer3TimeseriesGroupVerticalParamsProtocolTCP, AttackLayer3TimeseriesGroupVerticalParamsProtocolIcmp, AttackLayer3TimeseriesGroupVerticalParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupVerticalResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupVerticalResponseEnvelope]
type attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
