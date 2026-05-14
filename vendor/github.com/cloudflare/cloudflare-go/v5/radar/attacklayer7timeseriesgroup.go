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

// AttackLayer7TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TimeseriesGroupService] method instead.
type AttackLayer7TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackLayer7TimeseriesGroupService(opts ...option.RequestOption) (r *AttackLayer7TimeseriesGroupService) {
	r = &AttackLayer7TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP method over time.
func (r *AttackLayer7TimeseriesGroupService) HTTPMethod(ctx context.Context, query AttackLayer7TimeseriesGroupHTTPMethodParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupHTTPMethodResponse, err error) {
	var env AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP version over time.
func (r *AttackLayer7TimeseriesGroupService) HTTPVersion(ctx context.Context, query AttackLayer7TimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupHTTPVersionResponse, err error) {
	var env AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by targeted industry over time.
func (r *AttackLayer7TimeseriesGroupService) Industry(ctx context.Context, query AttackLayer7TimeseriesGroupIndustryParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupIndustryResponse, err error) {
	var env AttackLayer7TimeseriesGroupIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by IP version used over time.
func (r *AttackLayer7TimeseriesGroupService) IPVersion(ctx context.Context, query AttackLayer7TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupIPVersionResponse, err error) {
	var env AttackLayer7TimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by managed rules over time.
func (r *AttackLayer7TimeseriesGroupService) ManagedRules(ctx context.Context, query AttackLayer7TimeseriesGroupManagedRulesParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupManagedRulesResponse, err error) {
	var env AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by mitigation product over time.
func (r *AttackLayer7TimeseriesGroupService) MitigationProduct(ctx context.Context, query AttackLayer7TimeseriesGroupMitigationProductParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupMitigationProductResponse, err error) {
	var env AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by targeted vertical over time.
func (r *AttackLayer7TimeseriesGroupService) Vertical(ctx context.Context, query AttackLayer7TimeseriesGroupVerticalParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupVerticalResponse, err error) {
	var env AttackLayer7TimeseriesGroupVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TimeseriesGroupHTTPMethodResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupHTTPMethodResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupHTTPMethodResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupHTTPMethodResponse]
type attackLayer7TimeseriesGroupHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupHTTPMethodResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupHTTPMethodResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupHTTPMethodResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseMeta]
type attackLayer7TimeseriesGroupHTTPMethodResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRange]
type attackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupHTTPMethodResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupHTTPMethodResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseMetaUnit]
type attackLayer7TimeseriesGroupHTTPMethodResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0 struct {
	Timestamps  []time.Time                                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                     `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0]
type attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupHTTPVersionResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponse]
type attackLayer7TimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupHTTPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupHTTPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupHTTPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseMetaJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseMeta]
type attackLayer7TimeseriesGroupHTTPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRange]
type attackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupHTTPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupHTTPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseMetaUnit]
type attackLayer7TimeseriesGroupHTTPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                                 `json:"HTTP/1.x,required"`
	HTTP2      []string                                                 `json:"HTTP/2,required"`
	HTTP3      []string                                                 `json:"HTTP/3,required"`
	Timestamps []time.Time                                              `json:"timestamps,required" format:"date-time"`
	JSON       attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0]
type attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupIndustryResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupIndustryResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupIndustryResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupIndustryResponse]
type attackLayer7TimeseriesGroupIndustryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupIndustryResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupIndustryResponseMeta]
type attackLayer7TimeseriesGroupIndustryResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupIndustryResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupIndustryResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupIndustryResponseMetaDateRange]
type attackLayer7TimeseriesGroupIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIndustryResponseMetaUnit]
type attackLayer7TimeseriesGroupIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryResponseSerie0 struct {
	Timestamps  []time.Time                                           `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupIndustryResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupIndustryResponseSerie0]
type attackLayer7TimeseriesGroupIndustryResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupIPVersionResponse]
type attackLayer7TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupIPVersionResponseMeta]
type attackLayer7TimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupIPVersionResponseMetaDateRange]
type attackLayer7TimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIPVersionResponseMetaUnit]
type attackLayer7TimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                               `json:"IPv4,required"`
	IPv6       []string                                               `json:"IPv6,required"`
	Timestamps []time.Time                                            `json:"timestamps,required" format:"date-time"`
	JSON       attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIPVersionResponseSerie0]
type attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupManagedRulesResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupManagedRulesResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupManagedRulesResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupManagedRulesResponse]
type attackLayer7TimeseriesGroupManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupManagedRulesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupManagedRulesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupManagedRulesResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupManagedRulesResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseMetaJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupManagedRulesResponseMeta]
type attackLayer7TimeseriesGroupManagedRulesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupManagedRulesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                 `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                       `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupManagedRulesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseMetaDateRange]
type attackLayer7TimeseriesGroupManagedRulesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupManagedRulesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesResponseMetaUnit struct {
	Name  string                                                      `json:"name,required"`
	Value string                                                      `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupManagedRulesResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseMetaUnitJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseMetaUnit]
type attackLayer7TimeseriesGroupManagedRulesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesResponseSerie0 struct {
	Timestamps  []time.Time                                               `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                       `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupManagedRulesResponseSerie0]
type attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupMitigationProductResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupMitigationProductResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupMitigationProductResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupMitigationProductResponse]
type attackLayer7TimeseriesGroupMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupMitigationProductResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupMitigationProductResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupMitigationProductResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupMitigationProductResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseMetaJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseMeta]
type attackLayer7TimeseriesGroupMitigationProductResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupMitigationProductResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupMitigationProductResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseMetaDateRangeJSON contains
// the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseMetaDateRange]
type attackLayer7TimeseriesGroupMitigationProductResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupMitigationProductResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductResponseMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupMitigationProductResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseMetaUnitJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseMetaUnit]
type attackLayer7TimeseriesGroupMitigationProductResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductResponseSerie0 struct {
	Timestamps  []time.Time                                                    `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                            `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseSerie0]
type attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupVerticalResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupVerticalResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupVerticalResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupVerticalResponse]
type attackLayer7TimeseriesGroupVerticalResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupVerticalResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupVerticalResponseMeta]
type attackLayer7TimeseriesGroupVerticalResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupVerticalResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupVerticalResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupVerticalResponseMetaDateRange]
type attackLayer7TimeseriesGroupVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalResponseMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupVerticalResponseMetaUnit]
type attackLayer7TimeseriesGroupVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalResponseSerie0 struct {
	Timestamps  []time.Time                                           `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                   `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupVerticalResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseSerie0JSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupVerticalResponseSerie0]
type attackLayer7TimeseriesGroupVerticalResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPMethodParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupHTTPMethodParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1d  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1w  AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval15m, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1h, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1d, AttackLayer7TimeseriesGroupHTTPMethodParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupHTTPMethodParamsFormat string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson AttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupHTTPMethodParamsFormatCsv  AttackLayer7TimeseriesGroupHTTPMethodParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsFormatJson, AttackLayer7TimeseriesGroupHTTPMethodParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupHTTPMethodParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4 AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6 AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv4, AttackLayer7TimeseriesGroupHTTPMethodParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF                AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductWAF, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupHTTPMethodParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupHTTPMethodParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPMethodParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupHTTPMethodResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope]
type attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPMethodResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1d  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1w  AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval15m, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1h, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1d, AttackLayer7TimeseriesGroupHTTPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupHTTPVersionParamsFormat string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson AttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupHTTPVersionParamsFormatCsv  AttackLayer7TimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsFormatJson, AttackLayer7TimeseriesGroupHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPut             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodHead            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodACL             AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLock            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMove            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodReport          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodJson            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCook            AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodGet, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPost, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPut, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodHead, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodACL, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodLock, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodMove, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodReport, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodJson, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodCook, AttackLayer7TimeseriesGroupHTTPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4 AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6 AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv4, AttackLayer7TimeseriesGroupHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF                AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductWAF, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupHTTPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization string

const (
	AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupHTTPVersionParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationPercentage, AttackLayer7TimeseriesGroupHTTPVersionParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope]
type attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIndustryParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupIndustryParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupIndustryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupIndustryParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupIndustryParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIndustryParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1h  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1d  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupIndustryParamsAggInterval1w  AttackLayer7TimeseriesGroupIndustryParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsAggInterval15m, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1h, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1d, AttackLayer7TimeseriesGroupIndustryParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupIndustryParamsFormat string

const (
	AttackLayer7TimeseriesGroupIndustryParamsFormatJson AttackLayer7TimeseriesGroupIndustryParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupIndustryParamsFormatCsv  AttackLayer7TimeseriesGroupIndustryParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsFormatJson, AttackLayer7TimeseriesGroupIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPut             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodHead            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodACL             AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLock            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMove            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodReport          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodJson            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCook            AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodGet, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPost, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPut, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodHead, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodACL, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodLock, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodMove, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodReport, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodJson, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodCook, AttackLayer7TimeseriesGroupIndustryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupIndustryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4 AttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6 AttackLayer7TimeseriesGroupIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv4, AttackLayer7TimeseriesGroupIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF                AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductWAF, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupIndustryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIndustryParamsNormalization string

const (
	AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage AttackLayer7TimeseriesGroupIndustryParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIndustryParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupIndustryParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupIndustryParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIndustryParamsNormalizationPercentage, AttackLayer7TimeseriesGroupIndustryParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIndustryResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupIndustryResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIndustryResponseEnvelope]
type attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupIPVersionParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupIPVersionParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1d  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1w  AttackLayer7TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsAggInterval15m, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1h, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1d, AttackLayer7TimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupIPVersionParamsFormat string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsFormatJson AttackLayer7TimeseriesGroupIPVersionParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupIPVersionParamsFormatCsv  AttackLayer7TimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsFormatJson, AttackLayer7TimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPut             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodHead            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodACL             AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLock            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMove            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodReport          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodJson            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCook            AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodGet, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPost, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPut, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodHead, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodACL, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodLock, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodMove, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodReport, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodJson, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodCook, AttackLayer7TimeseriesGroupIPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF                AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductWAF, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupIPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupIPVersionParamsNormalization string

const (
	AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage AttackLayer7TimeseriesGroupIPVersionParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupIPVersionParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupIPVersionParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupIPVersionParamsNormalizationPercentage, AttackLayer7TimeseriesGroupIPVersionParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupIPVersionResponseEnvelope]
type attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupManagedRulesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupManagedRulesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupManagedRulesParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1d  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1w  AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval15m, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1h, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1d, AttackLayer7TimeseriesGroupManagedRulesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupManagedRulesParamsFormat string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson AttackLayer7TimeseriesGroupManagedRulesParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupManagedRulesParamsFormatCsv  AttackLayer7TimeseriesGroupManagedRulesParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsFormatJson, AttackLayer7TimeseriesGroupManagedRulesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPut             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodHead            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodACL             AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLock            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMove            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodReport          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodJson            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCook            AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodGet, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPost, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPut, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodHead, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodACL, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodLock, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodMove, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodReport, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodJson, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodCook, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupManagedRulesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4 AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6 AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv4, AttackLayer7TimeseriesGroupManagedRulesParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF                AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductWAF, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupManagedRulesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupManagedRulesParamsNormalization string

const (
	AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage AttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupManagedRulesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupManagedRulesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationPercentage, AttackLayer7TimeseriesGroupManagedRulesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupManagedRulesResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope]
type attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupManagedRulesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupMitigationProductParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[AttackLayer7TimeseriesGroupMitigationProductParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupMitigationProductParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1d  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1w  AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval15m, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1h, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1d, AttackLayer7TimeseriesGroupMitigationProductParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupMitigationProductParamsFormat string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson AttackLayer7TimeseriesGroupMitigationProductParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupMitigationProductParamsFormatCsv  AttackLayer7TimeseriesGroupMitigationProductParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsFormatJson, AttackLayer7TimeseriesGroupMitigationProductParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPut             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodHead            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodACL             AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLock            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMove            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodReport          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodJson            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCook            AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodGet, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPost, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPut, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodHead, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodACL, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodLock, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodMove, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodReport, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodJson, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodCook, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupMitigationProductParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4 AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6 AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv4, AttackLayer7TimeseriesGroupMitigationProductParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupMitigationProductParamsNormalization string

const (
	AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage AttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupMitigationProductParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupMitigationProductParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationPercentage, AttackLayer7TimeseriesGroupMitigationProductParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupMitigationProductResponse             `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope]
type attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupMitigationProductResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupVerticalParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupVerticalParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[AttackLayer7TimeseriesGroupVerticalParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupVerticalParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupVerticalParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7TimeseriesGroupVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupVerticalParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1h  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1d  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupVerticalParamsAggInterval1w  AttackLayer7TimeseriesGroupVerticalParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsAggInterval15m, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1h, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1d, AttackLayer7TimeseriesGroupVerticalParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupVerticalParamsFormat string

const (
	AttackLayer7TimeseriesGroupVerticalParamsFormatJson AttackLayer7TimeseriesGroupVerticalParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupVerticalParamsFormatCsv  AttackLayer7TimeseriesGroupVerticalParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsFormatJson, AttackLayer7TimeseriesGroupVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPut             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodHead            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodACL             AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLock            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMove            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodReport          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodJson            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCook            AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodGet, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPost, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPut, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodHead, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodACL, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodLock, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodMove, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodReport, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodJson, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodCook, AttackLayer7TimeseriesGroupVerticalParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupVerticalParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4 AttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6 AttackLayer7TimeseriesGroupVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv4, AttackLayer7TimeseriesGroupVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF                AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductWAF, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupVerticalParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupVerticalParamsNormalization string

const (
	AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage AttackLayer7TimeseriesGroupVerticalParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupVerticalParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupVerticalParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupVerticalParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupVerticalParamsNormalizationPercentage, AttackLayer7TimeseriesGroupVerticalParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupVerticalResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupVerticalResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupVerticalResponseEnvelope]
type attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
