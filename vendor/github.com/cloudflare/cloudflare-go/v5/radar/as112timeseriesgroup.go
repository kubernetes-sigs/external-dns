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

// AS112TimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112TimeseriesGroupService] method instead.
type AS112TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAS112TimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAS112TimeseriesGroupService(opts ...option.RequestOption) (r *AS112TimeseriesGroupService) {
	r = &AS112TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of AS112 DNS queries by DNSSEC (DNS Security
// Extensions) support over time.
func (r *AS112TimeseriesGroupService) DNSSEC(ctx context.Context, query AS112TimeseriesGroupDNSSECParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupDNSSECResponse, err error) {
	var env AS112TimeseriesGroupDNSSECResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS queries by EDNS (Extension Mechanisms
// for DNS) support over time.
func (r *AS112TimeseriesGroupService) Edns(ctx context.Context, query AS112TimeseriesGroupEdnsParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupEdnsResponse, err error) {
	var env AS112TimeseriesGroupEdnsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS queries by IP version over time.
func (r *AS112TimeseriesGroupService) IPVersion(ctx context.Context, query AS112TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupIPVersionResponse, err error) {
	var env AS112TimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS requests classified by protocol over
// time.
func (r *AS112TimeseriesGroupService) Protocol(ctx context.Context, query AS112TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupProtocolResponse, err error) {
	var env AS112TimeseriesGroupProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS queries by type over time.
func (r *AS112TimeseriesGroupService) QueryType(ctx context.Context, query AS112TimeseriesGroupQueryTypeParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupQueryTypeResponse, err error) {
	var env AS112TimeseriesGroupQueryTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS requests classified by response code
// over time.
func (r *AS112TimeseriesGroupService) ResponseCodes(ctx context.Context, query AS112TimeseriesGroupResponseCodesParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupResponseCodesResponse, err error) {
	var env AS112TimeseriesGroupResponseCodesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TimeseriesGroupDNSSECResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupDNSSECResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupDNSSECResponseJSON   `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesGroupDNSSECResponse]
type as112TimeseriesGroupDNSSECResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupDNSSECResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupDNSSECResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupDNSSECResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupDNSSECResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupDNSSECResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupDNSSECResponseMeta]
type as112TimeseriesGroupDNSSECResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupDNSSECResponseMetaAggInterval string

const (
	AS112TimeseriesGroupDNSSECResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupDNSSECResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneHour        AS112TimeseriesGroupDNSSECResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneDay         AS112TimeseriesGroupDNSSECResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupDNSSECResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupDNSSECResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupDNSSECResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneHour, AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneDay, AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupDNSSECResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfo]
type as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupDNSSECResponseMetaDateRange]
type as112TimeseriesGroupDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupDNSSECResponseMetaNormalization string

const (
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationPercentage           AS112TimeseriesGroupDNSSECResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationMin0Max              AS112TimeseriesGroupDNSSECResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationMinMax               AS112TimeseriesGroupDNSSECResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationRawValues            AS112TimeseriesGroupDNSSECResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupDNSSECResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupDNSSECResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupDNSSECResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupDNSSECResponseMetaNormalizationRatio                AS112TimeseriesGroupDNSSECResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupDNSSECResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECResponseMetaNormalizationPercentage, AS112TimeseriesGroupDNSSECResponseMetaNormalizationMin0Max, AS112TimeseriesGroupDNSSECResponseMetaNormalizationMinMax, AS112TimeseriesGroupDNSSECResponseMetaNormalizationRawValues, AS112TimeseriesGroupDNSSECResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupDNSSECResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupDNSSECResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupDNSSECResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupDNSSECResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  as112TimeseriesGroupDNSSECResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseMetaUnitJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupDNSSECResponseMetaUnit]
type as112TimeseriesGroupDNSSECResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupDNSSECResponseSerie0 struct {
	NotSupported []string                                     `json:"NOT_SUPPORTED,required"`
	Supported    []string                                     `json:"SUPPORTED,required"`
	JSON         as112TimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupDNSSECResponseSerie0]
type as112TimeseriesGroupDNSSECResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupEdnsResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupEdnsResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupEdnsResponseJSON   `json:"-"`
}

// as112TimeseriesGroupEdnsResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesGroupEdnsResponse]
type as112TimeseriesGroupEdnsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupEdnsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupEdnsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupEdnsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupEdnsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupEdnsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupEdnsResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupEdnsResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupEdnsResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseMeta]
type as112TimeseriesGroupEdnsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupEdnsResponseMetaAggInterval string

const (
	AS112TimeseriesGroupEdnsResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupEdnsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneHour        AS112TimeseriesGroupEdnsResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneDay         AS112TimeseriesGroupEdnsResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupEdnsResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupEdnsResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupEdnsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneHour, AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneDay, AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupEdnsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupEdnsResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  as112TimeseriesGroupEdnsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupEdnsResponseMetaConfidenceInfo]
type as112TimeseriesGroupEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AS112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupEdnsResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupEdnsResponseMetaDateRange]
type as112TimeseriesGroupEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupEdnsResponseMetaNormalization string

const (
	AS112TimeseriesGroupEdnsResponseMetaNormalizationPercentage           AS112TimeseriesGroupEdnsResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationMin0Max              AS112TimeseriesGroupEdnsResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationMinMax               AS112TimeseriesGroupEdnsResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationRawValues            AS112TimeseriesGroupEdnsResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupEdnsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupEdnsResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupEdnsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupEdnsResponseMetaNormalizationRatio                AS112TimeseriesGroupEdnsResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupEdnsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsResponseMetaNormalizationPercentage, AS112TimeseriesGroupEdnsResponseMetaNormalizationMin0Max, AS112TimeseriesGroupEdnsResponseMetaNormalizationMinMax, AS112TimeseriesGroupEdnsResponseMetaNormalizationRawValues, AS112TimeseriesGroupEdnsResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupEdnsResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupEdnsResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupEdnsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupEdnsResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  as112TimeseriesGroupEdnsResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseMetaUnit]
type as112TimeseriesGroupEdnsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsResponseSerie0 struct {
	NotSupported []string                                   `json:"NOT_SUPPORTED,required"`
	Supported    []string                                   `json:"SUPPORTED,required"`
	JSON         as112TimeseriesGroupEdnsResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseSerie0JSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseSerie0]
type as112TimeseriesGroupEdnsResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupIPVersionResponse]
type as112TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupIPVersionResponseMeta]
type as112TimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	AS112TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        AS112TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         AS112TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupIPVersionResponseMetaDateRange]
type as112TimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupIPVersionResponseMetaNormalization string

const (
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationPercentage           AS112TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              AS112TimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationMinMax               AS112TimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationRawValues            AS112TimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupIPVersionResponseMetaNormalizationRatio                AS112TimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionResponseMetaNormalizationPercentage, AS112TimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, AS112TimeseriesGroupIPVersionResponseMetaNormalizationMinMax, AS112TimeseriesGroupIPVersionResponseMetaNormalizationRawValues, AS112TimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  as112TimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupIPVersionResponseMetaUnit]
type as112TimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                        `json:"IPv4,required"`
	IPv6 []string                                        `json:"IPv6,required"`
	JSON as112TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupIPVersionResponseSerie0]
type as112TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupProtocolResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// as112TimeseriesGroupProtocolResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupProtocolResponse]
type as112TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupProtocolResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupProtocolResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupProtocolResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupProtocolResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupProtocolResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupProtocolResponseMeta]
type as112TimeseriesGroupProtocolResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupProtocolResponseMetaAggInterval string

const (
	AS112TimeseriesGroupProtocolResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupProtocolResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneHour        AS112TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneDay         AS112TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupProtocolResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupProtocolResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneHour, AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneDay, AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupProtocolResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupProtocolResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  as112TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupProtocolResponseMetaConfidenceInfo]
type as112TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AS112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupProtocolResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupProtocolResponseMetaDateRange]
type as112TimeseriesGroupProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupProtocolResponseMetaNormalization string

const (
	AS112TimeseriesGroupProtocolResponseMetaNormalizationPercentage           AS112TimeseriesGroupProtocolResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationMin0Max              AS112TimeseriesGroupProtocolResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationMinMax               AS112TimeseriesGroupProtocolResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationRawValues            AS112TimeseriesGroupProtocolResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupProtocolResponseMetaNormalizationRatio                AS112TimeseriesGroupProtocolResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolResponseMetaNormalizationPercentage, AS112TimeseriesGroupProtocolResponseMetaNormalizationMin0Max, AS112TimeseriesGroupProtocolResponseMetaNormalizationMinMax, AS112TimeseriesGroupProtocolResponseMetaNormalizationRawValues, AS112TimeseriesGroupProtocolResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupProtocolResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupProtocolResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupProtocolResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  as112TimeseriesGroupProtocolResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseMetaUnitJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupProtocolResponseMetaUnit]
type as112TimeseriesGroupProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolResponseSerie0 struct {
	HTTPS []string                                       `json:"HTTPS,required"`
	TCP   []string                                       `json:"TCP,required"`
	TLS   []string                                       `json:"TLS,required"`
	Udp   []string                                       `json:"UDP,required"`
	JSON  as112TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupProtocolResponseSerie0]
type as112TimeseriesGroupProtocolResponseSerie0JSON struct {
	HTTPS       apijson.Field
	TCP         apijson.Field
	TLS         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupQueryTypeResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupQueryTypeResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupQueryTypeResponseJSON   `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupQueryTypeResponse]
type as112TimeseriesGroupQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupQueryTypeResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupQueryTypeResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupQueryTypeResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupQueryTypeResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupQueryTypeResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupQueryTypeResponseMeta]
type as112TimeseriesGroupQueryTypeResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupQueryTypeResponseMetaAggInterval string

const (
	AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupQueryTypeResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneHour        AS112TimeseriesGroupQueryTypeResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneDay         AS112TimeseriesGroupQueryTypeResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupQueryTypeResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupQueryTypeResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupQueryTypeResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneHour, AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneDay, AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupQueryTypeResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfo]
type as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupQueryTypeResponseMetaDateRange]
type as112TimeseriesGroupQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupQueryTypeResponseMetaNormalization string

const (
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationPercentage           AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationMin0Max              AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationMinMax               AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRawValues            AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRatio                AS112TimeseriesGroupQueryTypeResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupQueryTypeResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeResponseMetaNormalizationPercentage, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationMin0Max, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationMinMax, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRawValues, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupQueryTypeResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupQueryTypeResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  as112TimeseriesGroupQueryTypeResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseMetaUnitJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupQueryTypeResponseMetaUnit]
type as112TimeseriesGroupQueryTypeResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeResponseSerie0 struct {
	Timestamps  []time.Time                                     `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                             `json:"-,extras"`
	JSON        as112TimeseriesGroupQueryTypeResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupQueryTypeResponseSerie0]
type as112TimeseriesGroupQueryTypeResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesResponse struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupResponseCodesResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupResponseCodesResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupResponseCodesResponseJSON   `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupResponseCodesResponse]
type as112TimeseriesGroupResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupResponseCodesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupResponseCodesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupResponseCodesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupResponseCodesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupResponseCodesResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupResponseCodesResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseMetaJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupResponseCodesResponseMeta]
type as112TimeseriesGroupResponseCodesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupResponseCodesResponseMetaAggInterval string

const (
	AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupResponseCodesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneHour        AS112TimeseriesGroupResponseCodesResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneDay         AS112TimeseriesGroupResponseCodesResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupResponseCodesResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupResponseCodesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupResponseCodesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneHour, AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneDay, AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupResponseCodesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfo]
type as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AS112TimeseriesGroupResponseCodesResponseMetaDateRange]
type as112TimeseriesGroupResponseCodesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupResponseCodesResponseMetaNormalization string

const (
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationPercentage           AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationMin0Max              AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationMinMax               AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRawValues            AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRatio                AS112TimeseriesGroupResponseCodesResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupResponseCodesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesResponseMetaNormalizationPercentage, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationMin0Max, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationMinMax, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRawValues, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupResponseCodesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupResponseCodesResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  as112TimeseriesGroupResponseCodesResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseMetaUnitJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupResponseCodesResponseMetaUnit]
type as112TimeseriesGroupResponseCodesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesResponseSerie0 struct {
	Timestamps  []time.Time                                         `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                 `json:"-,extras"`
	JSON        as112TimeseriesGroupResponseCodesResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseSerie0JSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupResponseCodesResponseSerie0]
type as112TimeseriesGroupResponseCodesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupDNSSECParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupDNSSECParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesGroupDNSSECParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesGroupDNSSECParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesGroupDNSSECParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupDNSSECParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupDNSSECParamsAggInterval string

const (
	AS112TimeseriesGroupDNSSECParamsAggInterval15m AS112TimeseriesGroupDNSSECParamsAggInterval = "15m"
	AS112TimeseriesGroupDNSSECParamsAggInterval1h  AS112TimeseriesGroupDNSSECParamsAggInterval = "1h"
	AS112TimeseriesGroupDNSSECParamsAggInterval1d  AS112TimeseriesGroupDNSSECParamsAggInterval = "1d"
	AS112TimeseriesGroupDNSSECParamsAggInterval1w  AS112TimeseriesGroupDNSSECParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupDNSSECParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsAggInterval15m, AS112TimeseriesGroupDNSSECParamsAggInterval1h, AS112TimeseriesGroupDNSSECParamsAggInterval1d, AS112TimeseriesGroupDNSSECParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupDNSSECParamsFormat string

const (
	AS112TimeseriesGroupDNSSECParamsFormatJson AS112TimeseriesGroupDNSSECParamsFormat = "JSON"
	AS112TimeseriesGroupDNSSECParamsFormatCsv  AS112TimeseriesGroupDNSSECParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsFormatJson, AS112TimeseriesGroupDNSSECParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesGroupDNSSECParamsProtocol string

const (
	AS112TimeseriesGroupDNSSECParamsProtocolUdp   AS112TimeseriesGroupDNSSECParamsProtocol = "UDP"
	AS112TimeseriesGroupDNSSECParamsProtocolTCP   AS112TimeseriesGroupDNSSECParamsProtocol = "TCP"
	AS112TimeseriesGroupDNSSECParamsProtocolHTTPS AS112TimeseriesGroupDNSSECParamsProtocol = "HTTPS"
	AS112TimeseriesGroupDNSSECParamsProtocolTLS   AS112TimeseriesGroupDNSSECParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupDNSSECParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsProtocolUdp, AS112TimeseriesGroupDNSSECParamsProtocolTCP, AS112TimeseriesGroupDNSSECParamsProtocolHTTPS, AS112TimeseriesGroupDNSSECParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesGroupDNSSECParamsQueryType string

const (
	AS112TimeseriesGroupDNSSECParamsQueryTypeA          AS112TimeseriesGroupDNSSECParamsQueryType = "A"
	AS112TimeseriesGroupDNSSECParamsQueryTypeAAAA       AS112TimeseriesGroupDNSSECParamsQueryType = "AAAA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeA6         AS112TimeseriesGroupDNSSECParamsQueryType = "A6"
	AS112TimeseriesGroupDNSSECParamsQueryTypeAfsdb      AS112TimeseriesGroupDNSSECParamsQueryType = "AFSDB"
	AS112TimeseriesGroupDNSSECParamsQueryTypeAny        AS112TimeseriesGroupDNSSECParamsQueryType = "ANY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeApl        AS112TimeseriesGroupDNSSECParamsQueryType = "APL"
	AS112TimeseriesGroupDNSSECParamsQueryTypeAtma       AS112TimeseriesGroupDNSSECParamsQueryType = "ATMA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeAXFR       AS112TimeseriesGroupDNSSECParamsQueryType = "AXFR"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCAA        AS112TimeseriesGroupDNSSECParamsQueryType = "CAA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCdnskey    AS112TimeseriesGroupDNSSECParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCds        AS112TimeseriesGroupDNSSECParamsQueryType = "CDS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCERT       AS112TimeseriesGroupDNSSECParamsQueryType = "CERT"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCNAME      AS112TimeseriesGroupDNSSECParamsQueryType = "CNAME"
	AS112TimeseriesGroupDNSSECParamsQueryTypeCsync      AS112TimeseriesGroupDNSSECParamsQueryType = "CSYNC"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDhcid      AS112TimeseriesGroupDNSSECParamsQueryType = "DHCID"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDlv        AS112TimeseriesGroupDNSSECParamsQueryType = "DLV"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDname      AS112TimeseriesGroupDNSSECParamsQueryType = "DNAME"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDNSKEY     AS112TimeseriesGroupDNSSECParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDoa        AS112TimeseriesGroupDNSSECParamsQueryType = "DOA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeDS         AS112TimeseriesGroupDNSSECParamsQueryType = "DS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeEid        AS112TimeseriesGroupDNSSECParamsQueryType = "EID"
	AS112TimeseriesGroupDNSSECParamsQueryTypeEui48      AS112TimeseriesGroupDNSSECParamsQueryType = "EUI48"
	AS112TimeseriesGroupDNSSECParamsQueryTypeEui64      AS112TimeseriesGroupDNSSECParamsQueryType = "EUI64"
	AS112TimeseriesGroupDNSSECParamsQueryTypeGpos       AS112TimeseriesGroupDNSSECParamsQueryType = "GPOS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeGid        AS112TimeseriesGroupDNSSECParamsQueryType = "GID"
	AS112TimeseriesGroupDNSSECParamsQueryTypeHinfo      AS112TimeseriesGroupDNSSECParamsQueryType = "HINFO"
	AS112TimeseriesGroupDNSSECParamsQueryTypeHip        AS112TimeseriesGroupDNSSECParamsQueryType = "HIP"
	AS112TimeseriesGroupDNSSECParamsQueryTypeHTTPS      AS112TimeseriesGroupDNSSECParamsQueryType = "HTTPS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeIpseckey   AS112TimeseriesGroupDNSSECParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeIsdn       AS112TimeseriesGroupDNSSECParamsQueryType = "ISDN"
	AS112TimeseriesGroupDNSSECParamsQueryTypeIxfr       AS112TimeseriesGroupDNSSECParamsQueryType = "IXFR"
	AS112TimeseriesGroupDNSSECParamsQueryTypeKey        AS112TimeseriesGroupDNSSECParamsQueryType = "KEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeKx         AS112TimeseriesGroupDNSSECParamsQueryType = "KX"
	AS112TimeseriesGroupDNSSECParamsQueryTypeL32        AS112TimeseriesGroupDNSSECParamsQueryType = "L32"
	AS112TimeseriesGroupDNSSECParamsQueryTypeL64        AS112TimeseriesGroupDNSSECParamsQueryType = "L64"
	AS112TimeseriesGroupDNSSECParamsQueryTypeLOC        AS112TimeseriesGroupDNSSECParamsQueryType = "LOC"
	AS112TimeseriesGroupDNSSECParamsQueryTypeLp         AS112TimeseriesGroupDNSSECParamsQueryType = "LP"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMaila      AS112TimeseriesGroupDNSSECParamsQueryType = "MAILA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMailb      AS112TimeseriesGroupDNSSECParamsQueryType = "MAILB"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMB         AS112TimeseriesGroupDNSSECParamsQueryType = "MB"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMd         AS112TimeseriesGroupDNSSECParamsQueryType = "MD"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMf         AS112TimeseriesGroupDNSSECParamsQueryType = "MF"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMg         AS112TimeseriesGroupDNSSECParamsQueryType = "MG"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMinfo      AS112TimeseriesGroupDNSSECParamsQueryType = "MINFO"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMr         AS112TimeseriesGroupDNSSECParamsQueryType = "MR"
	AS112TimeseriesGroupDNSSECParamsQueryTypeMX         AS112TimeseriesGroupDNSSECParamsQueryType = "MX"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNAPTR      AS112TimeseriesGroupDNSSECParamsQueryType = "NAPTR"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNb         AS112TimeseriesGroupDNSSECParamsQueryType = "NB"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNbstat     AS112TimeseriesGroupDNSSECParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNid        AS112TimeseriesGroupDNSSECParamsQueryType = "NID"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNimloc     AS112TimeseriesGroupDNSSECParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNinfo      AS112TimeseriesGroupDNSSECParamsQueryType = "NINFO"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNS         AS112TimeseriesGroupDNSSECParamsQueryType = "NS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNsap       AS112TimeseriesGroupDNSSECParamsQueryType = "NSAP"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNsec       AS112TimeseriesGroupDNSSECParamsQueryType = "NSEC"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNsec3      AS112TimeseriesGroupDNSSECParamsQueryType = "NSEC3"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNsec3Param AS112TimeseriesGroupDNSSECParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNull       AS112TimeseriesGroupDNSSECParamsQueryType = "NULL"
	AS112TimeseriesGroupDNSSECParamsQueryTypeNxt        AS112TimeseriesGroupDNSSECParamsQueryType = "NXT"
	AS112TimeseriesGroupDNSSECParamsQueryTypeOpenpgpkey AS112TimeseriesGroupDNSSECParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeOpt        AS112TimeseriesGroupDNSSECParamsQueryType = "OPT"
	AS112TimeseriesGroupDNSSECParamsQueryTypePTR        AS112TimeseriesGroupDNSSECParamsQueryType = "PTR"
	AS112TimeseriesGroupDNSSECParamsQueryTypePx         AS112TimeseriesGroupDNSSECParamsQueryType = "PX"
	AS112TimeseriesGroupDNSSECParamsQueryTypeRkey       AS112TimeseriesGroupDNSSECParamsQueryType = "RKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeRp         AS112TimeseriesGroupDNSSECParamsQueryType = "RP"
	AS112TimeseriesGroupDNSSECParamsQueryTypeRrsig      AS112TimeseriesGroupDNSSECParamsQueryType = "RRSIG"
	AS112TimeseriesGroupDNSSECParamsQueryTypeRt         AS112TimeseriesGroupDNSSECParamsQueryType = "RT"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSig        AS112TimeseriesGroupDNSSECParamsQueryType = "SIG"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSink       AS112TimeseriesGroupDNSSECParamsQueryType = "SINK"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSMIMEA     AS112TimeseriesGroupDNSSECParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSOA        AS112TimeseriesGroupDNSSECParamsQueryType = "SOA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSPF        AS112TimeseriesGroupDNSSECParamsQueryType = "SPF"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSRV        AS112TimeseriesGroupDNSSECParamsQueryType = "SRV"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSSHFP      AS112TimeseriesGroupDNSSECParamsQueryType = "SSHFP"
	AS112TimeseriesGroupDNSSECParamsQueryTypeSVCB       AS112TimeseriesGroupDNSSECParamsQueryType = "SVCB"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTa         AS112TimeseriesGroupDNSSECParamsQueryType = "TA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTalink     AS112TimeseriesGroupDNSSECParamsQueryType = "TALINK"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTkey       AS112TimeseriesGroupDNSSECParamsQueryType = "TKEY"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTLSA       AS112TimeseriesGroupDNSSECParamsQueryType = "TLSA"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTSIG       AS112TimeseriesGroupDNSSECParamsQueryType = "TSIG"
	AS112TimeseriesGroupDNSSECParamsQueryTypeTXT        AS112TimeseriesGroupDNSSECParamsQueryType = "TXT"
	AS112TimeseriesGroupDNSSECParamsQueryTypeUinfo      AS112TimeseriesGroupDNSSECParamsQueryType = "UINFO"
	AS112TimeseriesGroupDNSSECParamsQueryTypeUID        AS112TimeseriesGroupDNSSECParamsQueryType = "UID"
	AS112TimeseriesGroupDNSSECParamsQueryTypeUnspec     AS112TimeseriesGroupDNSSECParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupDNSSECParamsQueryTypeURI        AS112TimeseriesGroupDNSSECParamsQueryType = "URI"
	AS112TimeseriesGroupDNSSECParamsQueryTypeWks        AS112TimeseriesGroupDNSSECParamsQueryType = "WKS"
	AS112TimeseriesGroupDNSSECParamsQueryTypeX25        AS112TimeseriesGroupDNSSECParamsQueryType = "X25"
	AS112TimeseriesGroupDNSSECParamsQueryTypeZonemd     AS112TimeseriesGroupDNSSECParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupDNSSECParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsQueryTypeA, AS112TimeseriesGroupDNSSECParamsQueryTypeAAAA, AS112TimeseriesGroupDNSSECParamsQueryTypeA6, AS112TimeseriesGroupDNSSECParamsQueryTypeAfsdb, AS112TimeseriesGroupDNSSECParamsQueryTypeAny, AS112TimeseriesGroupDNSSECParamsQueryTypeApl, AS112TimeseriesGroupDNSSECParamsQueryTypeAtma, AS112TimeseriesGroupDNSSECParamsQueryTypeAXFR, AS112TimeseriesGroupDNSSECParamsQueryTypeCAA, AS112TimeseriesGroupDNSSECParamsQueryTypeCdnskey, AS112TimeseriesGroupDNSSECParamsQueryTypeCds, AS112TimeseriesGroupDNSSECParamsQueryTypeCERT, AS112TimeseriesGroupDNSSECParamsQueryTypeCNAME, AS112TimeseriesGroupDNSSECParamsQueryTypeCsync, AS112TimeseriesGroupDNSSECParamsQueryTypeDhcid, AS112TimeseriesGroupDNSSECParamsQueryTypeDlv, AS112TimeseriesGroupDNSSECParamsQueryTypeDname, AS112TimeseriesGroupDNSSECParamsQueryTypeDNSKEY, AS112TimeseriesGroupDNSSECParamsQueryTypeDoa, AS112TimeseriesGroupDNSSECParamsQueryTypeDS, AS112TimeseriesGroupDNSSECParamsQueryTypeEid, AS112TimeseriesGroupDNSSECParamsQueryTypeEui48, AS112TimeseriesGroupDNSSECParamsQueryTypeEui64, AS112TimeseriesGroupDNSSECParamsQueryTypeGpos, AS112TimeseriesGroupDNSSECParamsQueryTypeGid, AS112TimeseriesGroupDNSSECParamsQueryTypeHinfo, AS112TimeseriesGroupDNSSECParamsQueryTypeHip, AS112TimeseriesGroupDNSSECParamsQueryTypeHTTPS, AS112TimeseriesGroupDNSSECParamsQueryTypeIpseckey, AS112TimeseriesGroupDNSSECParamsQueryTypeIsdn, AS112TimeseriesGroupDNSSECParamsQueryTypeIxfr, AS112TimeseriesGroupDNSSECParamsQueryTypeKey, AS112TimeseriesGroupDNSSECParamsQueryTypeKx, AS112TimeseriesGroupDNSSECParamsQueryTypeL32, AS112TimeseriesGroupDNSSECParamsQueryTypeL64, AS112TimeseriesGroupDNSSECParamsQueryTypeLOC, AS112TimeseriesGroupDNSSECParamsQueryTypeLp, AS112TimeseriesGroupDNSSECParamsQueryTypeMaila, AS112TimeseriesGroupDNSSECParamsQueryTypeMailb, AS112TimeseriesGroupDNSSECParamsQueryTypeMB, AS112TimeseriesGroupDNSSECParamsQueryTypeMd, AS112TimeseriesGroupDNSSECParamsQueryTypeMf, AS112TimeseriesGroupDNSSECParamsQueryTypeMg, AS112TimeseriesGroupDNSSECParamsQueryTypeMinfo, AS112TimeseriesGroupDNSSECParamsQueryTypeMr, AS112TimeseriesGroupDNSSECParamsQueryTypeMX, AS112TimeseriesGroupDNSSECParamsQueryTypeNAPTR, AS112TimeseriesGroupDNSSECParamsQueryTypeNb, AS112TimeseriesGroupDNSSECParamsQueryTypeNbstat, AS112TimeseriesGroupDNSSECParamsQueryTypeNid, AS112TimeseriesGroupDNSSECParamsQueryTypeNimloc, AS112TimeseriesGroupDNSSECParamsQueryTypeNinfo, AS112TimeseriesGroupDNSSECParamsQueryTypeNS, AS112TimeseriesGroupDNSSECParamsQueryTypeNsap, AS112TimeseriesGroupDNSSECParamsQueryTypeNsec, AS112TimeseriesGroupDNSSECParamsQueryTypeNsec3, AS112TimeseriesGroupDNSSECParamsQueryTypeNsec3Param, AS112TimeseriesGroupDNSSECParamsQueryTypeNull, AS112TimeseriesGroupDNSSECParamsQueryTypeNxt, AS112TimeseriesGroupDNSSECParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupDNSSECParamsQueryTypeOpt, AS112TimeseriesGroupDNSSECParamsQueryTypePTR, AS112TimeseriesGroupDNSSECParamsQueryTypePx, AS112TimeseriesGroupDNSSECParamsQueryTypeRkey, AS112TimeseriesGroupDNSSECParamsQueryTypeRp, AS112TimeseriesGroupDNSSECParamsQueryTypeRrsig, AS112TimeseriesGroupDNSSECParamsQueryTypeRt, AS112TimeseriesGroupDNSSECParamsQueryTypeSig, AS112TimeseriesGroupDNSSECParamsQueryTypeSink, AS112TimeseriesGroupDNSSECParamsQueryTypeSMIMEA, AS112TimeseriesGroupDNSSECParamsQueryTypeSOA, AS112TimeseriesGroupDNSSECParamsQueryTypeSPF, AS112TimeseriesGroupDNSSECParamsQueryTypeSRV, AS112TimeseriesGroupDNSSECParamsQueryTypeSSHFP, AS112TimeseriesGroupDNSSECParamsQueryTypeSVCB, AS112TimeseriesGroupDNSSECParamsQueryTypeTa, AS112TimeseriesGroupDNSSECParamsQueryTypeTalink, AS112TimeseriesGroupDNSSECParamsQueryTypeTkey, AS112TimeseriesGroupDNSSECParamsQueryTypeTLSA, AS112TimeseriesGroupDNSSECParamsQueryTypeTSIG, AS112TimeseriesGroupDNSSECParamsQueryTypeTXT, AS112TimeseriesGroupDNSSECParamsQueryTypeUinfo, AS112TimeseriesGroupDNSSECParamsQueryTypeUID, AS112TimeseriesGroupDNSSECParamsQueryTypeUnspec, AS112TimeseriesGroupDNSSECParamsQueryTypeURI, AS112TimeseriesGroupDNSSECParamsQueryTypeWks, AS112TimeseriesGroupDNSSECParamsQueryTypeX25, AS112TimeseriesGroupDNSSECParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesGroupDNSSECParamsResponseCode string

const (
	AS112TimeseriesGroupDNSSECParamsResponseCodeNoerror   AS112TimeseriesGroupDNSSECParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupDNSSECParamsResponseCodeFormerr   AS112TimeseriesGroupDNSSECParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupDNSSECParamsResponseCodeServfail  AS112TimeseriesGroupDNSSECParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupDNSSECParamsResponseCodeNxdomain  AS112TimeseriesGroupDNSSECParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupDNSSECParamsResponseCodeNotimp    AS112TimeseriesGroupDNSSECParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupDNSSECParamsResponseCodeRefused   AS112TimeseriesGroupDNSSECParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupDNSSECParamsResponseCodeYxdomain  AS112TimeseriesGroupDNSSECParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupDNSSECParamsResponseCodeYxrrset   AS112TimeseriesGroupDNSSECParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupDNSSECParamsResponseCodeNxrrset   AS112TimeseriesGroupDNSSECParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupDNSSECParamsResponseCodeNotauth   AS112TimeseriesGroupDNSSECParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupDNSSECParamsResponseCodeNotzone   AS112TimeseriesGroupDNSSECParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadsig    AS112TimeseriesGroupDNSSECParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadkey    AS112TimeseriesGroupDNSSECParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadtime   AS112TimeseriesGroupDNSSECParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadmode   AS112TimeseriesGroupDNSSECParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadname   AS112TimeseriesGroupDNSSECParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadalg    AS112TimeseriesGroupDNSSECParamsResponseCode = "BADALG"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadtrunc  AS112TimeseriesGroupDNSSECParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupDNSSECParamsResponseCodeBadcookie AS112TimeseriesGroupDNSSECParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupDNSSECParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsResponseCodeNoerror, AS112TimeseriesGroupDNSSECParamsResponseCodeFormerr, AS112TimeseriesGroupDNSSECParamsResponseCodeServfail, AS112TimeseriesGroupDNSSECParamsResponseCodeNxdomain, AS112TimeseriesGroupDNSSECParamsResponseCodeNotimp, AS112TimeseriesGroupDNSSECParamsResponseCodeRefused, AS112TimeseriesGroupDNSSECParamsResponseCodeYxdomain, AS112TimeseriesGroupDNSSECParamsResponseCodeYxrrset, AS112TimeseriesGroupDNSSECParamsResponseCodeNxrrset, AS112TimeseriesGroupDNSSECParamsResponseCodeNotauth, AS112TimeseriesGroupDNSSECParamsResponseCodeNotzone, AS112TimeseriesGroupDNSSECParamsResponseCodeBadsig, AS112TimeseriesGroupDNSSECParamsResponseCodeBadkey, AS112TimeseriesGroupDNSSECParamsResponseCodeBadtime, AS112TimeseriesGroupDNSSECParamsResponseCodeBadmode, AS112TimeseriesGroupDNSSECParamsResponseCodeBadname, AS112TimeseriesGroupDNSSECParamsResponseCodeBadalg, AS112TimeseriesGroupDNSSECParamsResponseCodeBadtrunc, AS112TimeseriesGroupDNSSECParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupDNSSECResponseEnvelope struct {
	Result  AS112TimeseriesGroupDNSSECResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    as112TimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupDNSSECResponseEnvelope]
type as112TimeseriesGroupDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupEdnsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupEdnsParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesGroupEdnsParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesGroupEdnsParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesGroupEdnsParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupEdnsParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupEdnsParamsAggInterval string

const (
	AS112TimeseriesGroupEdnsParamsAggInterval15m AS112TimeseriesGroupEdnsParamsAggInterval = "15m"
	AS112TimeseriesGroupEdnsParamsAggInterval1h  AS112TimeseriesGroupEdnsParamsAggInterval = "1h"
	AS112TimeseriesGroupEdnsParamsAggInterval1d  AS112TimeseriesGroupEdnsParamsAggInterval = "1d"
	AS112TimeseriesGroupEdnsParamsAggInterval1w  AS112TimeseriesGroupEdnsParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupEdnsParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsAggInterval15m, AS112TimeseriesGroupEdnsParamsAggInterval1h, AS112TimeseriesGroupEdnsParamsAggInterval1d, AS112TimeseriesGroupEdnsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupEdnsParamsFormat string

const (
	AS112TimeseriesGroupEdnsParamsFormatJson AS112TimeseriesGroupEdnsParamsFormat = "JSON"
	AS112TimeseriesGroupEdnsParamsFormatCsv  AS112TimeseriesGroupEdnsParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsFormatJson, AS112TimeseriesGroupEdnsParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesGroupEdnsParamsProtocol string

const (
	AS112TimeseriesGroupEdnsParamsProtocolUdp   AS112TimeseriesGroupEdnsParamsProtocol = "UDP"
	AS112TimeseriesGroupEdnsParamsProtocolTCP   AS112TimeseriesGroupEdnsParamsProtocol = "TCP"
	AS112TimeseriesGroupEdnsParamsProtocolHTTPS AS112TimeseriesGroupEdnsParamsProtocol = "HTTPS"
	AS112TimeseriesGroupEdnsParamsProtocolTLS   AS112TimeseriesGroupEdnsParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupEdnsParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsProtocolUdp, AS112TimeseriesGroupEdnsParamsProtocolTCP, AS112TimeseriesGroupEdnsParamsProtocolHTTPS, AS112TimeseriesGroupEdnsParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesGroupEdnsParamsQueryType string

const (
	AS112TimeseriesGroupEdnsParamsQueryTypeA          AS112TimeseriesGroupEdnsParamsQueryType = "A"
	AS112TimeseriesGroupEdnsParamsQueryTypeAAAA       AS112TimeseriesGroupEdnsParamsQueryType = "AAAA"
	AS112TimeseriesGroupEdnsParamsQueryTypeA6         AS112TimeseriesGroupEdnsParamsQueryType = "A6"
	AS112TimeseriesGroupEdnsParamsQueryTypeAfsdb      AS112TimeseriesGroupEdnsParamsQueryType = "AFSDB"
	AS112TimeseriesGroupEdnsParamsQueryTypeAny        AS112TimeseriesGroupEdnsParamsQueryType = "ANY"
	AS112TimeseriesGroupEdnsParamsQueryTypeApl        AS112TimeseriesGroupEdnsParamsQueryType = "APL"
	AS112TimeseriesGroupEdnsParamsQueryTypeAtma       AS112TimeseriesGroupEdnsParamsQueryType = "ATMA"
	AS112TimeseriesGroupEdnsParamsQueryTypeAXFR       AS112TimeseriesGroupEdnsParamsQueryType = "AXFR"
	AS112TimeseriesGroupEdnsParamsQueryTypeCAA        AS112TimeseriesGroupEdnsParamsQueryType = "CAA"
	AS112TimeseriesGroupEdnsParamsQueryTypeCdnskey    AS112TimeseriesGroupEdnsParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeCds        AS112TimeseriesGroupEdnsParamsQueryType = "CDS"
	AS112TimeseriesGroupEdnsParamsQueryTypeCERT       AS112TimeseriesGroupEdnsParamsQueryType = "CERT"
	AS112TimeseriesGroupEdnsParamsQueryTypeCNAME      AS112TimeseriesGroupEdnsParamsQueryType = "CNAME"
	AS112TimeseriesGroupEdnsParamsQueryTypeCsync      AS112TimeseriesGroupEdnsParamsQueryType = "CSYNC"
	AS112TimeseriesGroupEdnsParamsQueryTypeDhcid      AS112TimeseriesGroupEdnsParamsQueryType = "DHCID"
	AS112TimeseriesGroupEdnsParamsQueryTypeDlv        AS112TimeseriesGroupEdnsParamsQueryType = "DLV"
	AS112TimeseriesGroupEdnsParamsQueryTypeDname      AS112TimeseriesGroupEdnsParamsQueryType = "DNAME"
	AS112TimeseriesGroupEdnsParamsQueryTypeDNSKEY     AS112TimeseriesGroupEdnsParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeDoa        AS112TimeseriesGroupEdnsParamsQueryType = "DOA"
	AS112TimeseriesGroupEdnsParamsQueryTypeDS         AS112TimeseriesGroupEdnsParamsQueryType = "DS"
	AS112TimeseriesGroupEdnsParamsQueryTypeEid        AS112TimeseriesGroupEdnsParamsQueryType = "EID"
	AS112TimeseriesGroupEdnsParamsQueryTypeEui48      AS112TimeseriesGroupEdnsParamsQueryType = "EUI48"
	AS112TimeseriesGroupEdnsParamsQueryTypeEui64      AS112TimeseriesGroupEdnsParamsQueryType = "EUI64"
	AS112TimeseriesGroupEdnsParamsQueryTypeGpos       AS112TimeseriesGroupEdnsParamsQueryType = "GPOS"
	AS112TimeseriesGroupEdnsParamsQueryTypeGid        AS112TimeseriesGroupEdnsParamsQueryType = "GID"
	AS112TimeseriesGroupEdnsParamsQueryTypeHinfo      AS112TimeseriesGroupEdnsParamsQueryType = "HINFO"
	AS112TimeseriesGroupEdnsParamsQueryTypeHip        AS112TimeseriesGroupEdnsParamsQueryType = "HIP"
	AS112TimeseriesGroupEdnsParamsQueryTypeHTTPS      AS112TimeseriesGroupEdnsParamsQueryType = "HTTPS"
	AS112TimeseriesGroupEdnsParamsQueryTypeIpseckey   AS112TimeseriesGroupEdnsParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeIsdn       AS112TimeseriesGroupEdnsParamsQueryType = "ISDN"
	AS112TimeseriesGroupEdnsParamsQueryTypeIxfr       AS112TimeseriesGroupEdnsParamsQueryType = "IXFR"
	AS112TimeseriesGroupEdnsParamsQueryTypeKey        AS112TimeseriesGroupEdnsParamsQueryType = "KEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeKx         AS112TimeseriesGroupEdnsParamsQueryType = "KX"
	AS112TimeseriesGroupEdnsParamsQueryTypeL32        AS112TimeseriesGroupEdnsParamsQueryType = "L32"
	AS112TimeseriesGroupEdnsParamsQueryTypeL64        AS112TimeseriesGroupEdnsParamsQueryType = "L64"
	AS112TimeseriesGroupEdnsParamsQueryTypeLOC        AS112TimeseriesGroupEdnsParamsQueryType = "LOC"
	AS112TimeseriesGroupEdnsParamsQueryTypeLp         AS112TimeseriesGroupEdnsParamsQueryType = "LP"
	AS112TimeseriesGroupEdnsParamsQueryTypeMaila      AS112TimeseriesGroupEdnsParamsQueryType = "MAILA"
	AS112TimeseriesGroupEdnsParamsQueryTypeMailb      AS112TimeseriesGroupEdnsParamsQueryType = "MAILB"
	AS112TimeseriesGroupEdnsParamsQueryTypeMB         AS112TimeseriesGroupEdnsParamsQueryType = "MB"
	AS112TimeseriesGroupEdnsParamsQueryTypeMd         AS112TimeseriesGroupEdnsParamsQueryType = "MD"
	AS112TimeseriesGroupEdnsParamsQueryTypeMf         AS112TimeseriesGroupEdnsParamsQueryType = "MF"
	AS112TimeseriesGroupEdnsParamsQueryTypeMg         AS112TimeseriesGroupEdnsParamsQueryType = "MG"
	AS112TimeseriesGroupEdnsParamsQueryTypeMinfo      AS112TimeseriesGroupEdnsParamsQueryType = "MINFO"
	AS112TimeseriesGroupEdnsParamsQueryTypeMr         AS112TimeseriesGroupEdnsParamsQueryType = "MR"
	AS112TimeseriesGroupEdnsParamsQueryTypeMX         AS112TimeseriesGroupEdnsParamsQueryType = "MX"
	AS112TimeseriesGroupEdnsParamsQueryTypeNAPTR      AS112TimeseriesGroupEdnsParamsQueryType = "NAPTR"
	AS112TimeseriesGroupEdnsParamsQueryTypeNb         AS112TimeseriesGroupEdnsParamsQueryType = "NB"
	AS112TimeseriesGroupEdnsParamsQueryTypeNbstat     AS112TimeseriesGroupEdnsParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupEdnsParamsQueryTypeNid        AS112TimeseriesGroupEdnsParamsQueryType = "NID"
	AS112TimeseriesGroupEdnsParamsQueryTypeNimloc     AS112TimeseriesGroupEdnsParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupEdnsParamsQueryTypeNinfo      AS112TimeseriesGroupEdnsParamsQueryType = "NINFO"
	AS112TimeseriesGroupEdnsParamsQueryTypeNS         AS112TimeseriesGroupEdnsParamsQueryType = "NS"
	AS112TimeseriesGroupEdnsParamsQueryTypeNsap       AS112TimeseriesGroupEdnsParamsQueryType = "NSAP"
	AS112TimeseriesGroupEdnsParamsQueryTypeNsec       AS112TimeseriesGroupEdnsParamsQueryType = "NSEC"
	AS112TimeseriesGroupEdnsParamsQueryTypeNsec3      AS112TimeseriesGroupEdnsParamsQueryType = "NSEC3"
	AS112TimeseriesGroupEdnsParamsQueryTypeNsec3Param AS112TimeseriesGroupEdnsParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupEdnsParamsQueryTypeNull       AS112TimeseriesGroupEdnsParamsQueryType = "NULL"
	AS112TimeseriesGroupEdnsParamsQueryTypeNxt        AS112TimeseriesGroupEdnsParamsQueryType = "NXT"
	AS112TimeseriesGroupEdnsParamsQueryTypeOpenpgpkey AS112TimeseriesGroupEdnsParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeOpt        AS112TimeseriesGroupEdnsParamsQueryType = "OPT"
	AS112TimeseriesGroupEdnsParamsQueryTypePTR        AS112TimeseriesGroupEdnsParamsQueryType = "PTR"
	AS112TimeseriesGroupEdnsParamsQueryTypePx         AS112TimeseriesGroupEdnsParamsQueryType = "PX"
	AS112TimeseriesGroupEdnsParamsQueryTypeRkey       AS112TimeseriesGroupEdnsParamsQueryType = "RKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeRp         AS112TimeseriesGroupEdnsParamsQueryType = "RP"
	AS112TimeseriesGroupEdnsParamsQueryTypeRrsig      AS112TimeseriesGroupEdnsParamsQueryType = "RRSIG"
	AS112TimeseriesGroupEdnsParamsQueryTypeRt         AS112TimeseriesGroupEdnsParamsQueryType = "RT"
	AS112TimeseriesGroupEdnsParamsQueryTypeSig        AS112TimeseriesGroupEdnsParamsQueryType = "SIG"
	AS112TimeseriesGroupEdnsParamsQueryTypeSink       AS112TimeseriesGroupEdnsParamsQueryType = "SINK"
	AS112TimeseriesGroupEdnsParamsQueryTypeSMIMEA     AS112TimeseriesGroupEdnsParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupEdnsParamsQueryTypeSOA        AS112TimeseriesGroupEdnsParamsQueryType = "SOA"
	AS112TimeseriesGroupEdnsParamsQueryTypeSPF        AS112TimeseriesGroupEdnsParamsQueryType = "SPF"
	AS112TimeseriesGroupEdnsParamsQueryTypeSRV        AS112TimeseriesGroupEdnsParamsQueryType = "SRV"
	AS112TimeseriesGroupEdnsParamsQueryTypeSSHFP      AS112TimeseriesGroupEdnsParamsQueryType = "SSHFP"
	AS112TimeseriesGroupEdnsParamsQueryTypeSVCB       AS112TimeseriesGroupEdnsParamsQueryType = "SVCB"
	AS112TimeseriesGroupEdnsParamsQueryTypeTa         AS112TimeseriesGroupEdnsParamsQueryType = "TA"
	AS112TimeseriesGroupEdnsParamsQueryTypeTalink     AS112TimeseriesGroupEdnsParamsQueryType = "TALINK"
	AS112TimeseriesGroupEdnsParamsQueryTypeTkey       AS112TimeseriesGroupEdnsParamsQueryType = "TKEY"
	AS112TimeseriesGroupEdnsParamsQueryTypeTLSA       AS112TimeseriesGroupEdnsParamsQueryType = "TLSA"
	AS112TimeseriesGroupEdnsParamsQueryTypeTSIG       AS112TimeseriesGroupEdnsParamsQueryType = "TSIG"
	AS112TimeseriesGroupEdnsParamsQueryTypeTXT        AS112TimeseriesGroupEdnsParamsQueryType = "TXT"
	AS112TimeseriesGroupEdnsParamsQueryTypeUinfo      AS112TimeseriesGroupEdnsParamsQueryType = "UINFO"
	AS112TimeseriesGroupEdnsParamsQueryTypeUID        AS112TimeseriesGroupEdnsParamsQueryType = "UID"
	AS112TimeseriesGroupEdnsParamsQueryTypeUnspec     AS112TimeseriesGroupEdnsParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupEdnsParamsQueryTypeURI        AS112TimeseriesGroupEdnsParamsQueryType = "URI"
	AS112TimeseriesGroupEdnsParamsQueryTypeWks        AS112TimeseriesGroupEdnsParamsQueryType = "WKS"
	AS112TimeseriesGroupEdnsParamsQueryTypeX25        AS112TimeseriesGroupEdnsParamsQueryType = "X25"
	AS112TimeseriesGroupEdnsParamsQueryTypeZonemd     AS112TimeseriesGroupEdnsParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupEdnsParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsQueryTypeA, AS112TimeseriesGroupEdnsParamsQueryTypeAAAA, AS112TimeseriesGroupEdnsParamsQueryTypeA6, AS112TimeseriesGroupEdnsParamsQueryTypeAfsdb, AS112TimeseriesGroupEdnsParamsQueryTypeAny, AS112TimeseriesGroupEdnsParamsQueryTypeApl, AS112TimeseriesGroupEdnsParamsQueryTypeAtma, AS112TimeseriesGroupEdnsParamsQueryTypeAXFR, AS112TimeseriesGroupEdnsParamsQueryTypeCAA, AS112TimeseriesGroupEdnsParamsQueryTypeCdnskey, AS112TimeseriesGroupEdnsParamsQueryTypeCds, AS112TimeseriesGroupEdnsParamsQueryTypeCERT, AS112TimeseriesGroupEdnsParamsQueryTypeCNAME, AS112TimeseriesGroupEdnsParamsQueryTypeCsync, AS112TimeseriesGroupEdnsParamsQueryTypeDhcid, AS112TimeseriesGroupEdnsParamsQueryTypeDlv, AS112TimeseriesGroupEdnsParamsQueryTypeDname, AS112TimeseriesGroupEdnsParamsQueryTypeDNSKEY, AS112TimeseriesGroupEdnsParamsQueryTypeDoa, AS112TimeseriesGroupEdnsParamsQueryTypeDS, AS112TimeseriesGroupEdnsParamsQueryTypeEid, AS112TimeseriesGroupEdnsParamsQueryTypeEui48, AS112TimeseriesGroupEdnsParamsQueryTypeEui64, AS112TimeseriesGroupEdnsParamsQueryTypeGpos, AS112TimeseriesGroupEdnsParamsQueryTypeGid, AS112TimeseriesGroupEdnsParamsQueryTypeHinfo, AS112TimeseriesGroupEdnsParamsQueryTypeHip, AS112TimeseriesGroupEdnsParamsQueryTypeHTTPS, AS112TimeseriesGroupEdnsParamsQueryTypeIpseckey, AS112TimeseriesGroupEdnsParamsQueryTypeIsdn, AS112TimeseriesGroupEdnsParamsQueryTypeIxfr, AS112TimeseriesGroupEdnsParamsQueryTypeKey, AS112TimeseriesGroupEdnsParamsQueryTypeKx, AS112TimeseriesGroupEdnsParamsQueryTypeL32, AS112TimeseriesGroupEdnsParamsQueryTypeL64, AS112TimeseriesGroupEdnsParamsQueryTypeLOC, AS112TimeseriesGroupEdnsParamsQueryTypeLp, AS112TimeseriesGroupEdnsParamsQueryTypeMaila, AS112TimeseriesGroupEdnsParamsQueryTypeMailb, AS112TimeseriesGroupEdnsParamsQueryTypeMB, AS112TimeseriesGroupEdnsParamsQueryTypeMd, AS112TimeseriesGroupEdnsParamsQueryTypeMf, AS112TimeseriesGroupEdnsParamsQueryTypeMg, AS112TimeseriesGroupEdnsParamsQueryTypeMinfo, AS112TimeseriesGroupEdnsParamsQueryTypeMr, AS112TimeseriesGroupEdnsParamsQueryTypeMX, AS112TimeseriesGroupEdnsParamsQueryTypeNAPTR, AS112TimeseriesGroupEdnsParamsQueryTypeNb, AS112TimeseriesGroupEdnsParamsQueryTypeNbstat, AS112TimeseriesGroupEdnsParamsQueryTypeNid, AS112TimeseriesGroupEdnsParamsQueryTypeNimloc, AS112TimeseriesGroupEdnsParamsQueryTypeNinfo, AS112TimeseriesGroupEdnsParamsQueryTypeNS, AS112TimeseriesGroupEdnsParamsQueryTypeNsap, AS112TimeseriesGroupEdnsParamsQueryTypeNsec, AS112TimeseriesGroupEdnsParamsQueryTypeNsec3, AS112TimeseriesGroupEdnsParamsQueryTypeNsec3Param, AS112TimeseriesGroupEdnsParamsQueryTypeNull, AS112TimeseriesGroupEdnsParamsQueryTypeNxt, AS112TimeseriesGroupEdnsParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupEdnsParamsQueryTypeOpt, AS112TimeseriesGroupEdnsParamsQueryTypePTR, AS112TimeseriesGroupEdnsParamsQueryTypePx, AS112TimeseriesGroupEdnsParamsQueryTypeRkey, AS112TimeseriesGroupEdnsParamsQueryTypeRp, AS112TimeseriesGroupEdnsParamsQueryTypeRrsig, AS112TimeseriesGroupEdnsParamsQueryTypeRt, AS112TimeseriesGroupEdnsParamsQueryTypeSig, AS112TimeseriesGroupEdnsParamsQueryTypeSink, AS112TimeseriesGroupEdnsParamsQueryTypeSMIMEA, AS112TimeseriesGroupEdnsParamsQueryTypeSOA, AS112TimeseriesGroupEdnsParamsQueryTypeSPF, AS112TimeseriesGroupEdnsParamsQueryTypeSRV, AS112TimeseriesGroupEdnsParamsQueryTypeSSHFP, AS112TimeseriesGroupEdnsParamsQueryTypeSVCB, AS112TimeseriesGroupEdnsParamsQueryTypeTa, AS112TimeseriesGroupEdnsParamsQueryTypeTalink, AS112TimeseriesGroupEdnsParamsQueryTypeTkey, AS112TimeseriesGroupEdnsParamsQueryTypeTLSA, AS112TimeseriesGroupEdnsParamsQueryTypeTSIG, AS112TimeseriesGroupEdnsParamsQueryTypeTXT, AS112TimeseriesGroupEdnsParamsQueryTypeUinfo, AS112TimeseriesGroupEdnsParamsQueryTypeUID, AS112TimeseriesGroupEdnsParamsQueryTypeUnspec, AS112TimeseriesGroupEdnsParamsQueryTypeURI, AS112TimeseriesGroupEdnsParamsQueryTypeWks, AS112TimeseriesGroupEdnsParamsQueryTypeX25, AS112TimeseriesGroupEdnsParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesGroupEdnsParamsResponseCode string

const (
	AS112TimeseriesGroupEdnsParamsResponseCodeNoerror   AS112TimeseriesGroupEdnsParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupEdnsParamsResponseCodeFormerr   AS112TimeseriesGroupEdnsParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupEdnsParamsResponseCodeServfail  AS112TimeseriesGroupEdnsParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupEdnsParamsResponseCodeNxdomain  AS112TimeseriesGroupEdnsParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupEdnsParamsResponseCodeNotimp    AS112TimeseriesGroupEdnsParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupEdnsParamsResponseCodeRefused   AS112TimeseriesGroupEdnsParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupEdnsParamsResponseCodeYxdomain  AS112TimeseriesGroupEdnsParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupEdnsParamsResponseCodeYxrrset   AS112TimeseriesGroupEdnsParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupEdnsParamsResponseCodeNxrrset   AS112TimeseriesGroupEdnsParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupEdnsParamsResponseCodeNotauth   AS112TimeseriesGroupEdnsParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupEdnsParamsResponseCodeNotzone   AS112TimeseriesGroupEdnsParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadsig    AS112TimeseriesGroupEdnsParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadkey    AS112TimeseriesGroupEdnsParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadtime   AS112TimeseriesGroupEdnsParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadmode   AS112TimeseriesGroupEdnsParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadname   AS112TimeseriesGroupEdnsParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadalg    AS112TimeseriesGroupEdnsParamsResponseCode = "BADALG"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadtrunc  AS112TimeseriesGroupEdnsParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupEdnsParamsResponseCodeBadcookie AS112TimeseriesGroupEdnsParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupEdnsParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsResponseCodeNoerror, AS112TimeseriesGroupEdnsParamsResponseCodeFormerr, AS112TimeseriesGroupEdnsParamsResponseCodeServfail, AS112TimeseriesGroupEdnsParamsResponseCodeNxdomain, AS112TimeseriesGroupEdnsParamsResponseCodeNotimp, AS112TimeseriesGroupEdnsParamsResponseCodeRefused, AS112TimeseriesGroupEdnsParamsResponseCodeYxdomain, AS112TimeseriesGroupEdnsParamsResponseCodeYxrrset, AS112TimeseriesGroupEdnsParamsResponseCodeNxrrset, AS112TimeseriesGroupEdnsParamsResponseCodeNotauth, AS112TimeseriesGroupEdnsParamsResponseCodeNotzone, AS112TimeseriesGroupEdnsParamsResponseCodeBadsig, AS112TimeseriesGroupEdnsParamsResponseCodeBadkey, AS112TimeseriesGroupEdnsParamsResponseCodeBadtime, AS112TimeseriesGroupEdnsParamsResponseCodeBadmode, AS112TimeseriesGroupEdnsParamsResponseCodeBadname, AS112TimeseriesGroupEdnsParamsResponseCodeBadalg, AS112TimeseriesGroupEdnsParamsResponseCodeBadtrunc, AS112TimeseriesGroupEdnsParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupEdnsResponseEnvelope struct {
	Result  AS112TimeseriesGroupEdnsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    as112TimeseriesGroupEdnsResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseEnvelope]
type as112TimeseriesGroupEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesGroupIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesGroupIPVersionParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupIPVersionParamsAggInterval string

const (
	AS112TimeseriesGroupIPVersionParamsAggInterval15m AS112TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AS112TimeseriesGroupIPVersionParamsAggInterval1h  AS112TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AS112TimeseriesGroupIPVersionParamsAggInterval1d  AS112TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AS112TimeseriesGroupIPVersionParamsAggInterval1w  AS112TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsAggInterval15m, AS112TimeseriesGroupIPVersionParamsAggInterval1h, AS112TimeseriesGroupIPVersionParamsAggInterval1d, AS112TimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupIPVersionParamsFormat string

const (
	AS112TimeseriesGroupIPVersionParamsFormatJson AS112TimeseriesGroupIPVersionParamsFormat = "JSON"
	AS112TimeseriesGroupIPVersionParamsFormatCsv  AS112TimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsFormatJson, AS112TimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesGroupIPVersionParamsProtocol string

const (
	AS112TimeseriesGroupIPVersionParamsProtocolUdp   AS112TimeseriesGroupIPVersionParamsProtocol = "UDP"
	AS112TimeseriesGroupIPVersionParamsProtocolTCP   AS112TimeseriesGroupIPVersionParamsProtocol = "TCP"
	AS112TimeseriesGroupIPVersionParamsProtocolHTTPS AS112TimeseriesGroupIPVersionParamsProtocol = "HTTPS"
	AS112TimeseriesGroupIPVersionParamsProtocolTLS   AS112TimeseriesGroupIPVersionParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsProtocolUdp, AS112TimeseriesGroupIPVersionParamsProtocolTCP, AS112TimeseriesGroupIPVersionParamsProtocolHTTPS, AS112TimeseriesGroupIPVersionParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesGroupIPVersionParamsQueryType string

const (
	AS112TimeseriesGroupIPVersionParamsQueryTypeA          AS112TimeseriesGroupIPVersionParamsQueryType = "A"
	AS112TimeseriesGroupIPVersionParamsQueryTypeAAAA       AS112TimeseriesGroupIPVersionParamsQueryType = "AAAA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeA6         AS112TimeseriesGroupIPVersionParamsQueryType = "A6"
	AS112TimeseriesGroupIPVersionParamsQueryTypeAfsdb      AS112TimeseriesGroupIPVersionParamsQueryType = "AFSDB"
	AS112TimeseriesGroupIPVersionParamsQueryTypeAny        AS112TimeseriesGroupIPVersionParamsQueryType = "ANY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeApl        AS112TimeseriesGroupIPVersionParamsQueryType = "APL"
	AS112TimeseriesGroupIPVersionParamsQueryTypeAtma       AS112TimeseriesGroupIPVersionParamsQueryType = "ATMA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeAXFR       AS112TimeseriesGroupIPVersionParamsQueryType = "AXFR"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCAA        AS112TimeseriesGroupIPVersionParamsQueryType = "CAA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCdnskey    AS112TimeseriesGroupIPVersionParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCds        AS112TimeseriesGroupIPVersionParamsQueryType = "CDS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCERT       AS112TimeseriesGroupIPVersionParamsQueryType = "CERT"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCNAME      AS112TimeseriesGroupIPVersionParamsQueryType = "CNAME"
	AS112TimeseriesGroupIPVersionParamsQueryTypeCsync      AS112TimeseriesGroupIPVersionParamsQueryType = "CSYNC"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDhcid      AS112TimeseriesGroupIPVersionParamsQueryType = "DHCID"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDlv        AS112TimeseriesGroupIPVersionParamsQueryType = "DLV"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDname      AS112TimeseriesGroupIPVersionParamsQueryType = "DNAME"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDNSKEY     AS112TimeseriesGroupIPVersionParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDoa        AS112TimeseriesGroupIPVersionParamsQueryType = "DOA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeDS         AS112TimeseriesGroupIPVersionParamsQueryType = "DS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeEid        AS112TimeseriesGroupIPVersionParamsQueryType = "EID"
	AS112TimeseriesGroupIPVersionParamsQueryTypeEui48      AS112TimeseriesGroupIPVersionParamsQueryType = "EUI48"
	AS112TimeseriesGroupIPVersionParamsQueryTypeEui64      AS112TimeseriesGroupIPVersionParamsQueryType = "EUI64"
	AS112TimeseriesGroupIPVersionParamsQueryTypeGpos       AS112TimeseriesGroupIPVersionParamsQueryType = "GPOS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeGid        AS112TimeseriesGroupIPVersionParamsQueryType = "GID"
	AS112TimeseriesGroupIPVersionParamsQueryTypeHinfo      AS112TimeseriesGroupIPVersionParamsQueryType = "HINFO"
	AS112TimeseriesGroupIPVersionParamsQueryTypeHip        AS112TimeseriesGroupIPVersionParamsQueryType = "HIP"
	AS112TimeseriesGroupIPVersionParamsQueryTypeHTTPS      AS112TimeseriesGroupIPVersionParamsQueryType = "HTTPS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeIpseckey   AS112TimeseriesGroupIPVersionParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeIsdn       AS112TimeseriesGroupIPVersionParamsQueryType = "ISDN"
	AS112TimeseriesGroupIPVersionParamsQueryTypeIxfr       AS112TimeseriesGroupIPVersionParamsQueryType = "IXFR"
	AS112TimeseriesGroupIPVersionParamsQueryTypeKey        AS112TimeseriesGroupIPVersionParamsQueryType = "KEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeKx         AS112TimeseriesGroupIPVersionParamsQueryType = "KX"
	AS112TimeseriesGroupIPVersionParamsQueryTypeL32        AS112TimeseriesGroupIPVersionParamsQueryType = "L32"
	AS112TimeseriesGroupIPVersionParamsQueryTypeL64        AS112TimeseriesGroupIPVersionParamsQueryType = "L64"
	AS112TimeseriesGroupIPVersionParamsQueryTypeLOC        AS112TimeseriesGroupIPVersionParamsQueryType = "LOC"
	AS112TimeseriesGroupIPVersionParamsQueryTypeLp         AS112TimeseriesGroupIPVersionParamsQueryType = "LP"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMaila      AS112TimeseriesGroupIPVersionParamsQueryType = "MAILA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMailb      AS112TimeseriesGroupIPVersionParamsQueryType = "MAILB"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMB         AS112TimeseriesGroupIPVersionParamsQueryType = "MB"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMd         AS112TimeseriesGroupIPVersionParamsQueryType = "MD"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMf         AS112TimeseriesGroupIPVersionParamsQueryType = "MF"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMg         AS112TimeseriesGroupIPVersionParamsQueryType = "MG"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMinfo      AS112TimeseriesGroupIPVersionParamsQueryType = "MINFO"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMr         AS112TimeseriesGroupIPVersionParamsQueryType = "MR"
	AS112TimeseriesGroupIPVersionParamsQueryTypeMX         AS112TimeseriesGroupIPVersionParamsQueryType = "MX"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNAPTR      AS112TimeseriesGroupIPVersionParamsQueryType = "NAPTR"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNb         AS112TimeseriesGroupIPVersionParamsQueryType = "NB"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNbstat     AS112TimeseriesGroupIPVersionParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNid        AS112TimeseriesGroupIPVersionParamsQueryType = "NID"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNimloc     AS112TimeseriesGroupIPVersionParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNinfo      AS112TimeseriesGroupIPVersionParamsQueryType = "NINFO"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNS         AS112TimeseriesGroupIPVersionParamsQueryType = "NS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNsap       AS112TimeseriesGroupIPVersionParamsQueryType = "NSAP"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNsec       AS112TimeseriesGroupIPVersionParamsQueryType = "NSEC"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNsec3      AS112TimeseriesGroupIPVersionParamsQueryType = "NSEC3"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNsec3Param AS112TimeseriesGroupIPVersionParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNull       AS112TimeseriesGroupIPVersionParamsQueryType = "NULL"
	AS112TimeseriesGroupIPVersionParamsQueryTypeNxt        AS112TimeseriesGroupIPVersionParamsQueryType = "NXT"
	AS112TimeseriesGroupIPVersionParamsQueryTypeOpenpgpkey AS112TimeseriesGroupIPVersionParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeOpt        AS112TimeseriesGroupIPVersionParamsQueryType = "OPT"
	AS112TimeseriesGroupIPVersionParamsQueryTypePTR        AS112TimeseriesGroupIPVersionParamsQueryType = "PTR"
	AS112TimeseriesGroupIPVersionParamsQueryTypePx         AS112TimeseriesGroupIPVersionParamsQueryType = "PX"
	AS112TimeseriesGroupIPVersionParamsQueryTypeRkey       AS112TimeseriesGroupIPVersionParamsQueryType = "RKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeRp         AS112TimeseriesGroupIPVersionParamsQueryType = "RP"
	AS112TimeseriesGroupIPVersionParamsQueryTypeRrsig      AS112TimeseriesGroupIPVersionParamsQueryType = "RRSIG"
	AS112TimeseriesGroupIPVersionParamsQueryTypeRt         AS112TimeseriesGroupIPVersionParamsQueryType = "RT"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSig        AS112TimeseriesGroupIPVersionParamsQueryType = "SIG"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSink       AS112TimeseriesGroupIPVersionParamsQueryType = "SINK"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSMIMEA     AS112TimeseriesGroupIPVersionParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSOA        AS112TimeseriesGroupIPVersionParamsQueryType = "SOA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSPF        AS112TimeseriesGroupIPVersionParamsQueryType = "SPF"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSRV        AS112TimeseriesGroupIPVersionParamsQueryType = "SRV"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSSHFP      AS112TimeseriesGroupIPVersionParamsQueryType = "SSHFP"
	AS112TimeseriesGroupIPVersionParamsQueryTypeSVCB       AS112TimeseriesGroupIPVersionParamsQueryType = "SVCB"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTa         AS112TimeseriesGroupIPVersionParamsQueryType = "TA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTalink     AS112TimeseriesGroupIPVersionParamsQueryType = "TALINK"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTkey       AS112TimeseriesGroupIPVersionParamsQueryType = "TKEY"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTLSA       AS112TimeseriesGroupIPVersionParamsQueryType = "TLSA"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTSIG       AS112TimeseriesGroupIPVersionParamsQueryType = "TSIG"
	AS112TimeseriesGroupIPVersionParamsQueryTypeTXT        AS112TimeseriesGroupIPVersionParamsQueryType = "TXT"
	AS112TimeseriesGroupIPVersionParamsQueryTypeUinfo      AS112TimeseriesGroupIPVersionParamsQueryType = "UINFO"
	AS112TimeseriesGroupIPVersionParamsQueryTypeUID        AS112TimeseriesGroupIPVersionParamsQueryType = "UID"
	AS112TimeseriesGroupIPVersionParamsQueryTypeUnspec     AS112TimeseriesGroupIPVersionParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupIPVersionParamsQueryTypeURI        AS112TimeseriesGroupIPVersionParamsQueryType = "URI"
	AS112TimeseriesGroupIPVersionParamsQueryTypeWks        AS112TimeseriesGroupIPVersionParamsQueryType = "WKS"
	AS112TimeseriesGroupIPVersionParamsQueryTypeX25        AS112TimeseriesGroupIPVersionParamsQueryType = "X25"
	AS112TimeseriesGroupIPVersionParamsQueryTypeZonemd     AS112TimeseriesGroupIPVersionParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsQueryTypeA, AS112TimeseriesGroupIPVersionParamsQueryTypeAAAA, AS112TimeseriesGroupIPVersionParamsQueryTypeA6, AS112TimeseriesGroupIPVersionParamsQueryTypeAfsdb, AS112TimeseriesGroupIPVersionParamsQueryTypeAny, AS112TimeseriesGroupIPVersionParamsQueryTypeApl, AS112TimeseriesGroupIPVersionParamsQueryTypeAtma, AS112TimeseriesGroupIPVersionParamsQueryTypeAXFR, AS112TimeseriesGroupIPVersionParamsQueryTypeCAA, AS112TimeseriesGroupIPVersionParamsQueryTypeCdnskey, AS112TimeseriesGroupIPVersionParamsQueryTypeCds, AS112TimeseriesGroupIPVersionParamsQueryTypeCERT, AS112TimeseriesGroupIPVersionParamsQueryTypeCNAME, AS112TimeseriesGroupIPVersionParamsQueryTypeCsync, AS112TimeseriesGroupIPVersionParamsQueryTypeDhcid, AS112TimeseriesGroupIPVersionParamsQueryTypeDlv, AS112TimeseriesGroupIPVersionParamsQueryTypeDname, AS112TimeseriesGroupIPVersionParamsQueryTypeDNSKEY, AS112TimeseriesGroupIPVersionParamsQueryTypeDoa, AS112TimeseriesGroupIPVersionParamsQueryTypeDS, AS112TimeseriesGroupIPVersionParamsQueryTypeEid, AS112TimeseriesGroupIPVersionParamsQueryTypeEui48, AS112TimeseriesGroupIPVersionParamsQueryTypeEui64, AS112TimeseriesGroupIPVersionParamsQueryTypeGpos, AS112TimeseriesGroupIPVersionParamsQueryTypeGid, AS112TimeseriesGroupIPVersionParamsQueryTypeHinfo, AS112TimeseriesGroupIPVersionParamsQueryTypeHip, AS112TimeseriesGroupIPVersionParamsQueryTypeHTTPS, AS112TimeseriesGroupIPVersionParamsQueryTypeIpseckey, AS112TimeseriesGroupIPVersionParamsQueryTypeIsdn, AS112TimeseriesGroupIPVersionParamsQueryTypeIxfr, AS112TimeseriesGroupIPVersionParamsQueryTypeKey, AS112TimeseriesGroupIPVersionParamsQueryTypeKx, AS112TimeseriesGroupIPVersionParamsQueryTypeL32, AS112TimeseriesGroupIPVersionParamsQueryTypeL64, AS112TimeseriesGroupIPVersionParamsQueryTypeLOC, AS112TimeseriesGroupIPVersionParamsQueryTypeLp, AS112TimeseriesGroupIPVersionParamsQueryTypeMaila, AS112TimeseriesGroupIPVersionParamsQueryTypeMailb, AS112TimeseriesGroupIPVersionParamsQueryTypeMB, AS112TimeseriesGroupIPVersionParamsQueryTypeMd, AS112TimeseriesGroupIPVersionParamsQueryTypeMf, AS112TimeseriesGroupIPVersionParamsQueryTypeMg, AS112TimeseriesGroupIPVersionParamsQueryTypeMinfo, AS112TimeseriesGroupIPVersionParamsQueryTypeMr, AS112TimeseriesGroupIPVersionParamsQueryTypeMX, AS112TimeseriesGroupIPVersionParamsQueryTypeNAPTR, AS112TimeseriesGroupIPVersionParamsQueryTypeNb, AS112TimeseriesGroupIPVersionParamsQueryTypeNbstat, AS112TimeseriesGroupIPVersionParamsQueryTypeNid, AS112TimeseriesGroupIPVersionParamsQueryTypeNimloc, AS112TimeseriesGroupIPVersionParamsQueryTypeNinfo, AS112TimeseriesGroupIPVersionParamsQueryTypeNS, AS112TimeseriesGroupIPVersionParamsQueryTypeNsap, AS112TimeseriesGroupIPVersionParamsQueryTypeNsec, AS112TimeseriesGroupIPVersionParamsQueryTypeNsec3, AS112TimeseriesGroupIPVersionParamsQueryTypeNsec3Param, AS112TimeseriesGroupIPVersionParamsQueryTypeNull, AS112TimeseriesGroupIPVersionParamsQueryTypeNxt, AS112TimeseriesGroupIPVersionParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupIPVersionParamsQueryTypeOpt, AS112TimeseriesGroupIPVersionParamsQueryTypePTR, AS112TimeseriesGroupIPVersionParamsQueryTypePx, AS112TimeseriesGroupIPVersionParamsQueryTypeRkey, AS112TimeseriesGroupIPVersionParamsQueryTypeRp, AS112TimeseriesGroupIPVersionParamsQueryTypeRrsig, AS112TimeseriesGroupIPVersionParamsQueryTypeRt, AS112TimeseriesGroupIPVersionParamsQueryTypeSig, AS112TimeseriesGroupIPVersionParamsQueryTypeSink, AS112TimeseriesGroupIPVersionParamsQueryTypeSMIMEA, AS112TimeseriesGroupIPVersionParamsQueryTypeSOA, AS112TimeseriesGroupIPVersionParamsQueryTypeSPF, AS112TimeseriesGroupIPVersionParamsQueryTypeSRV, AS112TimeseriesGroupIPVersionParamsQueryTypeSSHFP, AS112TimeseriesGroupIPVersionParamsQueryTypeSVCB, AS112TimeseriesGroupIPVersionParamsQueryTypeTa, AS112TimeseriesGroupIPVersionParamsQueryTypeTalink, AS112TimeseriesGroupIPVersionParamsQueryTypeTkey, AS112TimeseriesGroupIPVersionParamsQueryTypeTLSA, AS112TimeseriesGroupIPVersionParamsQueryTypeTSIG, AS112TimeseriesGroupIPVersionParamsQueryTypeTXT, AS112TimeseriesGroupIPVersionParamsQueryTypeUinfo, AS112TimeseriesGroupIPVersionParamsQueryTypeUID, AS112TimeseriesGroupIPVersionParamsQueryTypeUnspec, AS112TimeseriesGroupIPVersionParamsQueryTypeURI, AS112TimeseriesGroupIPVersionParamsQueryTypeWks, AS112TimeseriesGroupIPVersionParamsQueryTypeX25, AS112TimeseriesGroupIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesGroupIPVersionParamsResponseCode string

const (
	AS112TimeseriesGroupIPVersionParamsResponseCodeNoerror   AS112TimeseriesGroupIPVersionParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupIPVersionParamsResponseCodeFormerr   AS112TimeseriesGroupIPVersionParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupIPVersionParamsResponseCodeServfail  AS112TimeseriesGroupIPVersionParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupIPVersionParamsResponseCodeNxdomain  AS112TimeseriesGroupIPVersionParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupIPVersionParamsResponseCodeNotimp    AS112TimeseriesGroupIPVersionParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupIPVersionParamsResponseCodeRefused   AS112TimeseriesGroupIPVersionParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupIPVersionParamsResponseCodeYxdomain  AS112TimeseriesGroupIPVersionParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupIPVersionParamsResponseCodeYxrrset   AS112TimeseriesGroupIPVersionParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupIPVersionParamsResponseCodeNxrrset   AS112TimeseriesGroupIPVersionParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupIPVersionParamsResponseCodeNotauth   AS112TimeseriesGroupIPVersionParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupIPVersionParamsResponseCodeNotzone   AS112TimeseriesGroupIPVersionParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadsig    AS112TimeseriesGroupIPVersionParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadkey    AS112TimeseriesGroupIPVersionParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadtime   AS112TimeseriesGroupIPVersionParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadmode   AS112TimeseriesGroupIPVersionParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadname   AS112TimeseriesGroupIPVersionParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadalg    AS112TimeseriesGroupIPVersionParamsResponseCode = "BADALG"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadtrunc  AS112TimeseriesGroupIPVersionParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupIPVersionParamsResponseCodeBadcookie AS112TimeseriesGroupIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsResponseCodeNoerror, AS112TimeseriesGroupIPVersionParamsResponseCodeFormerr, AS112TimeseriesGroupIPVersionParamsResponseCodeServfail, AS112TimeseriesGroupIPVersionParamsResponseCodeNxdomain, AS112TimeseriesGroupIPVersionParamsResponseCodeNotimp, AS112TimeseriesGroupIPVersionParamsResponseCodeRefused, AS112TimeseriesGroupIPVersionParamsResponseCodeYxdomain, AS112TimeseriesGroupIPVersionParamsResponseCodeYxrrset, AS112TimeseriesGroupIPVersionParamsResponseCodeNxrrset, AS112TimeseriesGroupIPVersionParamsResponseCodeNotauth, AS112TimeseriesGroupIPVersionParamsResponseCodeNotzone, AS112TimeseriesGroupIPVersionParamsResponseCodeBadsig, AS112TimeseriesGroupIPVersionParamsResponseCodeBadkey, AS112TimeseriesGroupIPVersionParamsResponseCodeBadtime, AS112TimeseriesGroupIPVersionParamsResponseCodeBadmode, AS112TimeseriesGroupIPVersionParamsResponseCodeBadname, AS112TimeseriesGroupIPVersionParamsResponseCodeBadalg, AS112TimeseriesGroupIPVersionParamsResponseCodeBadtrunc, AS112TimeseriesGroupIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AS112TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    as112TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupIPVersionResponseEnvelope]
type as112TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesGroupProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesGroupProtocolParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupProtocolParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupProtocolParamsAggInterval string

const (
	AS112TimeseriesGroupProtocolParamsAggInterval15m AS112TimeseriesGroupProtocolParamsAggInterval = "15m"
	AS112TimeseriesGroupProtocolParamsAggInterval1h  AS112TimeseriesGroupProtocolParamsAggInterval = "1h"
	AS112TimeseriesGroupProtocolParamsAggInterval1d  AS112TimeseriesGroupProtocolParamsAggInterval = "1d"
	AS112TimeseriesGroupProtocolParamsAggInterval1w  AS112TimeseriesGroupProtocolParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsAggInterval15m, AS112TimeseriesGroupProtocolParamsAggInterval1h, AS112TimeseriesGroupProtocolParamsAggInterval1d, AS112TimeseriesGroupProtocolParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupProtocolParamsFormat string

const (
	AS112TimeseriesGroupProtocolParamsFormatJson AS112TimeseriesGroupProtocolParamsFormat = "JSON"
	AS112TimeseriesGroupProtocolParamsFormatCsv  AS112TimeseriesGroupProtocolParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsFormatJson, AS112TimeseriesGroupProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesGroupProtocolParamsQueryType string

const (
	AS112TimeseriesGroupProtocolParamsQueryTypeA          AS112TimeseriesGroupProtocolParamsQueryType = "A"
	AS112TimeseriesGroupProtocolParamsQueryTypeAAAA       AS112TimeseriesGroupProtocolParamsQueryType = "AAAA"
	AS112TimeseriesGroupProtocolParamsQueryTypeA6         AS112TimeseriesGroupProtocolParamsQueryType = "A6"
	AS112TimeseriesGroupProtocolParamsQueryTypeAfsdb      AS112TimeseriesGroupProtocolParamsQueryType = "AFSDB"
	AS112TimeseriesGroupProtocolParamsQueryTypeAny        AS112TimeseriesGroupProtocolParamsQueryType = "ANY"
	AS112TimeseriesGroupProtocolParamsQueryTypeApl        AS112TimeseriesGroupProtocolParamsQueryType = "APL"
	AS112TimeseriesGroupProtocolParamsQueryTypeAtma       AS112TimeseriesGroupProtocolParamsQueryType = "ATMA"
	AS112TimeseriesGroupProtocolParamsQueryTypeAXFR       AS112TimeseriesGroupProtocolParamsQueryType = "AXFR"
	AS112TimeseriesGroupProtocolParamsQueryTypeCAA        AS112TimeseriesGroupProtocolParamsQueryType = "CAA"
	AS112TimeseriesGroupProtocolParamsQueryTypeCdnskey    AS112TimeseriesGroupProtocolParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeCds        AS112TimeseriesGroupProtocolParamsQueryType = "CDS"
	AS112TimeseriesGroupProtocolParamsQueryTypeCERT       AS112TimeseriesGroupProtocolParamsQueryType = "CERT"
	AS112TimeseriesGroupProtocolParamsQueryTypeCNAME      AS112TimeseriesGroupProtocolParamsQueryType = "CNAME"
	AS112TimeseriesGroupProtocolParamsQueryTypeCsync      AS112TimeseriesGroupProtocolParamsQueryType = "CSYNC"
	AS112TimeseriesGroupProtocolParamsQueryTypeDhcid      AS112TimeseriesGroupProtocolParamsQueryType = "DHCID"
	AS112TimeseriesGroupProtocolParamsQueryTypeDlv        AS112TimeseriesGroupProtocolParamsQueryType = "DLV"
	AS112TimeseriesGroupProtocolParamsQueryTypeDname      AS112TimeseriesGroupProtocolParamsQueryType = "DNAME"
	AS112TimeseriesGroupProtocolParamsQueryTypeDNSKEY     AS112TimeseriesGroupProtocolParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeDoa        AS112TimeseriesGroupProtocolParamsQueryType = "DOA"
	AS112TimeseriesGroupProtocolParamsQueryTypeDS         AS112TimeseriesGroupProtocolParamsQueryType = "DS"
	AS112TimeseriesGroupProtocolParamsQueryTypeEid        AS112TimeseriesGroupProtocolParamsQueryType = "EID"
	AS112TimeseriesGroupProtocolParamsQueryTypeEui48      AS112TimeseriesGroupProtocolParamsQueryType = "EUI48"
	AS112TimeseriesGroupProtocolParamsQueryTypeEui64      AS112TimeseriesGroupProtocolParamsQueryType = "EUI64"
	AS112TimeseriesGroupProtocolParamsQueryTypeGpos       AS112TimeseriesGroupProtocolParamsQueryType = "GPOS"
	AS112TimeseriesGroupProtocolParamsQueryTypeGid        AS112TimeseriesGroupProtocolParamsQueryType = "GID"
	AS112TimeseriesGroupProtocolParamsQueryTypeHinfo      AS112TimeseriesGroupProtocolParamsQueryType = "HINFO"
	AS112TimeseriesGroupProtocolParamsQueryTypeHip        AS112TimeseriesGroupProtocolParamsQueryType = "HIP"
	AS112TimeseriesGroupProtocolParamsQueryTypeHTTPS      AS112TimeseriesGroupProtocolParamsQueryType = "HTTPS"
	AS112TimeseriesGroupProtocolParamsQueryTypeIpseckey   AS112TimeseriesGroupProtocolParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeIsdn       AS112TimeseriesGroupProtocolParamsQueryType = "ISDN"
	AS112TimeseriesGroupProtocolParamsQueryTypeIxfr       AS112TimeseriesGroupProtocolParamsQueryType = "IXFR"
	AS112TimeseriesGroupProtocolParamsQueryTypeKey        AS112TimeseriesGroupProtocolParamsQueryType = "KEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeKx         AS112TimeseriesGroupProtocolParamsQueryType = "KX"
	AS112TimeseriesGroupProtocolParamsQueryTypeL32        AS112TimeseriesGroupProtocolParamsQueryType = "L32"
	AS112TimeseriesGroupProtocolParamsQueryTypeL64        AS112TimeseriesGroupProtocolParamsQueryType = "L64"
	AS112TimeseriesGroupProtocolParamsQueryTypeLOC        AS112TimeseriesGroupProtocolParamsQueryType = "LOC"
	AS112TimeseriesGroupProtocolParamsQueryTypeLp         AS112TimeseriesGroupProtocolParamsQueryType = "LP"
	AS112TimeseriesGroupProtocolParamsQueryTypeMaila      AS112TimeseriesGroupProtocolParamsQueryType = "MAILA"
	AS112TimeseriesGroupProtocolParamsQueryTypeMailb      AS112TimeseriesGroupProtocolParamsQueryType = "MAILB"
	AS112TimeseriesGroupProtocolParamsQueryTypeMB         AS112TimeseriesGroupProtocolParamsQueryType = "MB"
	AS112TimeseriesGroupProtocolParamsQueryTypeMd         AS112TimeseriesGroupProtocolParamsQueryType = "MD"
	AS112TimeseriesGroupProtocolParamsQueryTypeMf         AS112TimeseriesGroupProtocolParamsQueryType = "MF"
	AS112TimeseriesGroupProtocolParamsQueryTypeMg         AS112TimeseriesGroupProtocolParamsQueryType = "MG"
	AS112TimeseriesGroupProtocolParamsQueryTypeMinfo      AS112TimeseriesGroupProtocolParamsQueryType = "MINFO"
	AS112TimeseriesGroupProtocolParamsQueryTypeMr         AS112TimeseriesGroupProtocolParamsQueryType = "MR"
	AS112TimeseriesGroupProtocolParamsQueryTypeMX         AS112TimeseriesGroupProtocolParamsQueryType = "MX"
	AS112TimeseriesGroupProtocolParamsQueryTypeNAPTR      AS112TimeseriesGroupProtocolParamsQueryType = "NAPTR"
	AS112TimeseriesGroupProtocolParamsQueryTypeNb         AS112TimeseriesGroupProtocolParamsQueryType = "NB"
	AS112TimeseriesGroupProtocolParamsQueryTypeNbstat     AS112TimeseriesGroupProtocolParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupProtocolParamsQueryTypeNid        AS112TimeseriesGroupProtocolParamsQueryType = "NID"
	AS112TimeseriesGroupProtocolParamsQueryTypeNimloc     AS112TimeseriesGroupProtocolParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupProtocolParamsQueryTypeNinfo      AS112TimeseriesGroupProtocolParamsQueryType = "NINFO"
	AS112TimeseriesGroupProtocolParamsQueryTypeNS         AS112TimeseriesGroupProtocolParamsQueryType = "NS"
	AS112TimeseriesGroupProtocolParamsQueryTypeNsap       AS112TimeseriesGroupProtocolParamsQueryType = "NSAP"
	AS112TimeseriesGroupProtocolParamsQueryTypeNsec       AS112TimeseriesGroupProtocolParamsQueryType = "NSEC"
	AS112TimeseriesGroupProtocolParamsQueryTypeNsec3      AS112TimeseriesGroupProtocolParamsQueryType = "NSEC3"
	AS112TimeseriesGroupProtocolParamsQueryTypeNsec3Param AS112TimeseriesGroupProtocolParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupProtocolParamsQueryTypeNull       AS112TimeseriesGroupProtocolParamsQueryType = "NULL"
	AS112TimeseriesGroupProtocolParamsQueryTypeNxt        AS112TimeseriesGroupProtocolParamsQueryType = "NXT"
	AS112TimeseriesGroupProtocolParamsQueryTypeOpenpgpkey AS112TimeseriesGroupProtocolParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeOpt        AS112TimeseriesGroupProtocolParamsQueryType = "OPT"
	AS112TimeseriesGroupProtocolParamsQueryTypePTR        AS112TimeseriesGroupProtocolParamsQueryType = "PTR"
	AS112TimeseriesGroupProtocolParamsQueryTypePx         AS112TimeseriesGroupProtocolParamsQueryType = "PX"
	AS112TimeseriesGroupProtocolParamsQueryTypeRkey       AS112TimeseriesGroupProtocolParamsQueryType = "RKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeRp         AS112TimeseriesGroupProtocolParamsQueryType = "RP"
	AS112TimeseriesGroupProtocolParamsQueryTypeRrsig      AS112TimeseriesGroupProtocolParamsQueryType = "RRSIG"
	AS112TimeseriesGroupProtocolParamsQueryTypeRt         AS112TimeseriesGroupProtocolParamsQueryType = "RT"
	AS112TimeseriesGroupProtocolParamsQueryTypeSig        AS112TimeseriesGroupProtocolParamsQueryType = "SIG"
	AS112TimeseriesGroupProtocolParamsQueryTypeSink       AS112TimeseriesGroupProtocolParamsQueryType = "SINK"
	AS112TimeseriesGroupProtocolParamsQueryTypeSMIMEA     AS112TimeseriesGroupProtocolParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupProtocolParamsQueryTypeSOA        AS112TimeseriesGroupProtocolParamsQueryType = "SOA"
	AS112TimeseriesGroupProtocolParamsQueryTypeSPF        AS112TimeseriesGroupProtocolParamsQueryType = "SPF"
	AS112TimeseriesGroupProtocolParamsQueryTypeSRV        AS112TimeseriesGroupProtocolParamsQueryType = "SRV"
	AS112TimeseriesGroupProtocolParamsQueryTypeSSHFP      AS112TimeseriesGroupProtocolParamsQueryType = "SSHFP"
	AS112TimeseriesGroupProtocolParamsQueryTypeSVCB       AS112TimeseriesGroupProtocolParamsQueryType = "SVCB"
	AS112TimeseriesGroupProtocolParamsQueryTypeTa         AS112TimeseriesGroupProtocolParamsQueryType = "TA"
	AS112TimeseriesGroupProtocolParamsQueryTypeTalink     AS112TimeseriesGroupProtocolParamsQueryType = "TALINK"
	AS112TimeseriesGroupProtocolParamsQueryTypeTkey       AS112TimeseriesGroupProtocolParamsQueryType = "TKEY"
	AS112TimeseriesGroupProtocolParamsQueryTypeTLSA       AS112TimeseriesGroupProtocolParamsQueryType = "TLSA"
	AS112TimeseriesGroupProtocolParamsQueryTypeTSIG       AS112TimeseriesGroupProtocolParamsQueryType = "TSIG"
	AS112TimeseriesGroupProtocolParamsQueryTypeTXT        AS112TimeseriesGroupProtocolParamsQueryType = "TXT"
	AS112TimeseriesGroupProtocolParamsQueryTypeUinfo      AS112TimeseriesGroupProtocolParamsQueryType = "UINFO"
	AS112TimeseriesGroupProtocolParamsQueryTypeUID        AS112TimeseriesGroupProtocolParamsQueryType = "UID"
	AS112TimeseriesGroupProtocolParamsQueryTypeUnspec     AS112TimeseriesGroupProtocolParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupProtocolParamsQueryTypeURI        AS112TimeseriesGroupProtocolParamsQueryType = "URI"
	AS112TimeseriesGroupProtocolParamsQueryTypeWks        AS112TimeseriesGroupProtocolParamsQueryType = "WKS"
	AS112TimeseriesGroupProtocolParamsQueryTypeX25        AS112TimeseriesGroupProtocolParamsQueryType = "X25"
	AS112TimeseriesGroupProtocolParamsQueryTypeZonemd     AS112TimeseriesGroupProtocolParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsQueryTypeA, AS112TimeseriesGroupProtocolParamsQueryTypeAAAA, AS112TimeseriesGroupProtocolParamsQueryTypeA6, AS112TimeseriesGroupProtocolParamsQueryTypeAfsdb, AS112TimeseriesGroupProtocolParamsQueryTypeAny, AS112TimeseriesGroupProtocolParamsQueryTypeApl, AS112TimeseriesGroupProtocolParamsQueryTypeAtma, AS112TimeseriesGroupProtocolParamsQueryTypeAXFR, AS112TimeseriesGroupProtocolParamsQueryTypeCAA, AS112TimeseriesGroupProtocolParamsQueryTypeCdnskey, AS112TimeseriesGroupProtocolParamsQueryTypeCds, AS112TimeseriesGroupProtocolParamsQueryTypeCERT, AS112TimeseriesGroupProtocolParamsQueryTypeCNAME, AS112TimeseriesGroupProtocolParamsQueryTypeCsync, AS112TimeseriesGroupProtocolParamsQueryTypeDhcid, AS112TimeseriesGroupProtocolParamsQueryTypeDlv, AS112TimeseriesGroupProtocolParamsQueryTypeDname, AS112TimeseriesGroupProtocolParamsQueryTypeDNSKEY, AS112TimeseriesGroupProtocolParamsQueryTypeDoa, AS112TimeseriesGroupProtocolParamsQueryTypeDS, AS112TimeseriesGroupProtocolParamsQueryTypeEid, AS112TimeseriesGroupProtocolParamsQueryTypeEui48, AS112TimeseriesGroupProtocolParamsQueryTypeEui64, AS112TimeseriesGroupProtocolParamsQueryTypeGpos, AS112TimeseriesGroupProtocolParamsQueryTypeGid, AS112TimeseriesGroupProtocolParamsQueryTypeHinfo, AS112TimeseriesGroupProtocolParamsQueryTypeHip, AS112TimeseriesGroupProtocolParamsQueryTypeHTTPS, AS112TimeseriesGroupProtocolParamsQueryTypeIpseckey, AS112TimeseriesGroupProtocolParamsQueryTypeIsdn, AS112TimeseriesGroupProtocolParamsQueryTypeIxfr, AS112TimeseriesGroupProtocolParamsQueryTypeKey, AS112TimeseriesGroupProtocolParamsQueryTypeKx, AS112TimeseriesGroupProtocolParamsQueryTypeL32, AS112TimeseriesGroupProtocolParamsQueryTypeL64, AS112TimeseriesGroupProtocolParamsQueryTypeLOC, AS112TimeseriesGroupProtocolParamsQueryTypeLp, AS112TimeseriesGroupProtocolParamsQueryTypeMaila, AS112TimeseriesGroupProtocolParamsQueryTypeMailb, AS112TimeseriesGroupProtocolParamsQueryTypeMB, AS112TimeseriesGroupProtocolParamsQueryTypeMd, AS112TimeseriesGroupProtocolParamsQueryTypeMf, AS112TimeseriesGroupProtocolParamsQueryTypeMg, AS112TimeseriesGroupProtocolParamsQueryTypeMinfo, AS112TimeseriesGroupProtocolParamsQueryTypeMr, AS112TimeseriesGroupProtocolParamsQueryTypeMX, AS112TimeseriesGroupProtocolParamsQueryTypeNAPTR, AS112TimeseriesGroupProtocolParamsQueryTypeNb, AS112TimeseriesGroupProtocolParamsQueryTypeNbstat, AS112TimeseriesGroupProtocolParamsQueryTypeNid, AS112TimeseriesGroupProtocolParamsQueryTypeNimloc, AS112TimeseriesGroupProtocolParamsQueryTypeNinfo, AS112TimeseriesGroupProtocolParamsQueryTypeNS, AS112TimeseriesGroupProtocolParamsQueryTypeNsap, AS112TimeseriesGroupProtocolParamsQueryTypeNsec, AS112TimeseriesGroupProtocolParamsQueryTypeNsec3, AS112TimeseriesGroupProtocolParamsQueryTypeNsec3Param, AS112TimeseriesGroupProtocolParamsQueryTypeNull, AS112TimeseriesGroupProtocolParamsQueryTypeNxt, AS112TimeseriesGroupProtocolParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupProtocolParamsQueryTypeOpt, AS112TimeseriesGroupProtocolParamsQueryTypePTR, AS112TimeseriesGroupProtocolParamsQueryTypePx, AS112TimeseriesGroupProtocolParamsQueryTypeRkey, AS112TimeseriesGroupProtocolParamsQueryTypeRp, AS112TimeseriesGroupProtocolParamsQueryTypeRrsig, AS112TimeseriesGroupProtocolParamsQueryTypeRt, AS112TimeseriesGroupProtocolParamsQueryTypeSig, AS112TimeseriesGroupProtocolParamsQueryTypeSink, AS112TimeseriesGroupProtocolParamsQueryTypeSMIMEA, AS112TimeseriesGroupProtocolParamsQueryTypeSOA, AS112TimeseriesGroupProtocolParamsQueryTypeSPF, AS112TimeseriesGroupProtocolParamsQueryTypeSRV, AS112TimeseriesGroupProtocolParamsQueryTypeSSHFP, AS112TimeseriesGroupProtocolParamsQueryTypeSVCB, AS112TimeseriesGroupProtocolParamsQueryTypeTa, AS112TimeseriesGroupProtocolParamsQueryTypeTalink, AS112TimeseriesGroupProtocolParamsQueryTypeTkey, AS112TimeseriesGroupProtocolParamsQueryTypeTLSA, AS112TimeseriesGroupProtocolParamsQueryTypeTSIG, AS112TimeseriesGroupProtocolParamsQueryTypeTXT, AS112TimeseriesGroupProtocolParamsQueryTypeUinfo, AS112TimeseriesGroupProtocolParamsQueryTypeUID, AS112TimeseriesGroupProtocolParamsQueryTypeUnspec, AS112TimeseriesGroupProtocolParamsQueryTypeURI, AS112TimeseriesGroupProtocolParamsQueryTypeWks, AS112TimeseriesGroupProtocolParamsQueryTypeX25, AS112TimeseriesGroupProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesGroupProtocolParamsResponseCode string

const (
	AS112TimeseriesGroupProtocolParamsResponseCodeNoerror   AS112TimeseriesGroupProtocolParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupProtocolParamsResponseCodeFormerr   AS112TimeseriesGroupProtocolParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupProtocolParamsResponseCodeServfail  AS112TimeseriesGroupProtocolParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupProtocolParamsResponseCodeNxdomain  AS112TimeseriesGroupProtocolParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupProtocolParamsResponseCodeNotimp    AS112TimeseriesGroupProtocolParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupProtocolParamsResponseCodeRefused   AS112TimeseriesGroupProtocolParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupProtocolParamsResponseCodeYxdomain  AS112TimeseriesGroupProtocolParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupProtocolParamsResponseCodeYxrrset   AS112TimeseriesGroupProtocolParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupProtocolParamsResponseCodeNxrrset   AS112TimeseriesGroupProtocolParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupProtocolParamsResponseCodeNotauth   AS112TimeseriesGroupProtocolParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupProtocolParamsResponseCodeNotzone   AS112TimeseriesGroupProtocolParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadsig    AS112TimeseriesGroupProtocolParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadkey    AS112TimeseriesGroupProtocolParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadtime   AS112TimeseriesGroupProtocolParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadmode   AS112TimeseriesGroupProtocolParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadname   AS112TimeseriesGroupProtocolParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadalg    AS112TimeseriesGroupProtocolParamsResponseCode = "BADALG"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadtrunc  AS112TimeseriesGroupProtocolParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupProtocolParamsResponseCodeBadcookie AS112TimeseriesGroupProtocolParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsResponseCodeNoerror, AS112TimeseriesGroupProtocolParamsResponseCodeFormerr, AS112TimeseriesGroupProtocolParamsResponseCodeServfail, AS112TimeseriesGroupProtocolParamsResponseCodeNxdomain, AS112TimeseriesGroupProtocolParamsResponseCodeNotimp, AS112TimeseriesGroupProtocolParamsResponseCodeRefused, AS112TimeseriesGroupProtocolParamsResponseCodeYxdomain, AS112TimeseriesGroupProtocolParamsResponseCodeYxrrset, AS112TimeseriesGroupProtocolParamsResponseCodeNxrrset, AS112TimeseriesGroupProtocolParamsResponseCodeNotauth, AS112TimeseriesGroupProtocolParamsResponseCodeNotzone, AS112TimeseriesGroupProtocolParamsResponseCodeBadsig, AS112TimeseriesGroupProtocolParamsResponseCodeBadkey, AS112TimeseriesGroupProtocolParamsResponseCodeBadtime, AS112TimeseriesGroupProtocolParamsResponseCodeBadmode, AS112TimeseriesGroupProtocolParamsResponseCodeBadname, AS112TimeseriesGroupProtocolParamsResponseCodeBadalg, AS112TimeseriesGroupProtocolParamsResponseCodeBadtrunc, AS112TimeseriesGroupProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupProtocolResponseEnvelope struct {
	Result  AS112TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    as112TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupProtocolResponseEnvelope]
type as112TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupQueryTypeParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesGroupQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesGroupQueryTypeParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupQueryTypeParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupQueryTypeParamsAggInterval string

const (
	AS112TimeseriesGroupQueryTypeParamsAggInterval15m AS112TimeseriesGroupQueryTypeParamsAggInterval = "15m"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1h  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1h"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1d  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1d"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1w  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupQueryTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsAggInterval15m, AS112TimeseriesGroupQueryTypeParamsAggInterval1h, AS112TimeseriesGroupQueryTypeParamsAggInterval1d, AS112TimeseriesGroupQueryTypeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupQueryTypeParamsFormat string

const (
	AS112TimeseriesGroupQueryTypeParamsFormatJson AS112TimeseriesGroupQueryTypeParamsFormat = "JSON"
	AS112TimeseriesGroupQueryTypeParamsFormatCsv  AS112TimeseriesGroupQueryTypeParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsFormatJson, AS112TimeseriesGroupQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesGroupQueryTypeParamsProtocol string

const (
	AS112TimeseriesGroupQueryTypeParamsProtocolUdp   AS112TimeseriesGroupQueryTypeParamsProtocol = "UDP"
	AS112TimeseriesGroupQueryTypeParamsProtocolTCP   AS112TimeseriesGroupQueryTypeParamsProtocol = "TCP"
	AS112TimeseriesGroupQueryTypeParamsProtocolHTTPS AS112TimeseriesGroupQueryTypeParamsProtocol = "HTTPS"
	AS112TimeseriesGroupQueryTypeParamsProtocolTLS   AS112TimeseriesGroupQueryTypeParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsProtocolUdp, AS112TimeseriesGroupQueryTypeParamsProtocolTCP, AS112TimeseriesGroupQueryTypeParamsProtocolHTTPS, AS112TimeseriesGroupQueryTypeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesGroupQueryTypeParamsResponseCode string

const (
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNoerror   AS112TimeseriesGroupQueryTypeParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeFormerr   AS112TimeseriesGroupQueryTypeParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeServfail  AS112TimeseriesGroupQueryTypeParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNxdomain  AS112TimeseriesGroupQueryTypeParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNotimp    AS112TimeseriesGroupQueryTypeParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeRefused   AS112TimeseriesGroupQueryTypeParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeYxdomain  AS112TimeseriesGroupQueryTypeParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeYxrrset   AS112TimeseriesGroupQueryTypeParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNxrrset   AS112TimeseriesGroupQueryTypeParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNotauth   AS112TimeseriesGroupQueryTypeParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeNotzone   AS112TimeseriesGroupQueryTypeParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadsig    AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadkey    AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadtime   AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadmode   AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadname   AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadalg    AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADALG"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadtrunc  AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupQueryTypeParamsResponseCodeBadcookie AS112TimeseriesGroupQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsResponseCodeNoerror, AS112TimeseriesGroupQueryTypeParamsResponseCodeFormerr, AS112TimeseriesGroupQueryTypeParamsResponseCodeServfail, AS112TimeseriesGroupQueryTypeParamsResponseCodeNxdomain, AS112TimeseriesGroupQueryTypeParamsResponseCodeNotimp, AS112TimeseriesGroupQueryTypeParamsResponseCodeRefused, AS112TimeseriesGroupQueryTypeParamsResponseCodeYxdomain, AS112TimeseriesGroupQueryTypeParamsResponseCodeYxrrset, AS112TimeseriesGroupQueryTypeParamsResponseCodeNxrrset, AS112TimeseriesGroupQueryTypeParamsResponseCodeNotauth, AS112TimeseriesGroupQueryTypeParamsResponseCodeNotzone, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadsig, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadkey, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadtime, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadmode, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadname, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadalg, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadtrunc, AS112TimeseriesGroupQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupQueryTypeResponseEnvelope struct {
	Result  AS112TimeseriesGroupQueryTypeResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    as112TimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupQueryTypeResponseEnvelope]
type as112TimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupResponseCodesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupResponseCodesParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesGroupResponseCodesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesGroupResponseCodesParamsQueryType] `query:"queryType"`
}

// URLQuery serializes [AS112TimeseriesGroupResponseCodesParams]'s query parameters
// as `url.Values`.
func (r AS112TimeseriesGroupResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupResponseCodesParamsAggInterval string

const (
	AS112TimeseriesGroupResponseCodesParamsAggInterval15m AS112TimeseriesGroupResponseCodesParamsAggInterval = "15m"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1h  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1h"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1d  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1d"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1w  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupResponseCodesParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsAggInterval15m, AS112TimeseriesGroupResponseCodesParamsAggInterval1h, AS112TimeseriesGroupResponseCodesParamsAggInterval1d, AS112TimeseriesGroupResponseCodesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupResponseCodesParamsFormat string

const (
	AS112TimeseriesGroupResponseCodesParamsFormatJson AS112TimeseriesGroupResponseCodesParamsFormat = "JSON"
	AS112TimeseriesGroupResponseCodesParamsFormatCsv  AS112TimeseriesGroupResponseCodesParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsFormatJson, AS112TimeseriesGroupResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesGroupResponseCodesParamsProtocol string

const (
	AS112TimeseriesGroupResponseCodesParamsProtocolUdp   AS112TimeseriesGroupResponseCodesParamsProtocol = "UDP"
	AS112TimeseriesGroupResponseCodesParamsProtocolTCP   AS112TimeseriesGroupResponseCodesParamsProtocol = "TCP"
	AS112TimeseriesGroupResponseCodesParamsProtocolHTTPS AS112TimeseriesGroupResponseCodesParamsProtocol = "HTTPS"
	AS112TimeseriesGroupResponseCodesParamsProtocolTLS   AS112TimeseriesGroupResponseCodesParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupResponseCodesParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsProtocolUdp, AS112TimeseriesGroupResponseCodesParamsProtocolTCP, AS112TimeseriesGroupResponseCodesParamsProtocolHTTPS, AS112TimeseriesGroupResponseCodesParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesGroupResponseCodesParamsQueryType string

const (
	AS112TimeseriesGroupResponseCodesParamsQueryTypeA          AS112TimeseriesGroupResponseCodesParamsQueryType = "A"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeAAAA       AS112TimeseriesGroupResponseCodesParamsQueryType = "AAAA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeA6         AS112TimeseriesGroupResponseCodesParamsQueryType = "A6"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeAfsdb      AS112TimeseriesGroupResponseCodesParamsQueryType = "AFSDB"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeAny        AS112TimeseriesGroupResponseCodesParamsQueryType = "ANY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeApl        AS112TimeseriesGroupResponseCodesParamsQueryType = "APL"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeAtma       AS112TimeseriesGroupResponseCodesParamsQueryType = "ATMA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeAXFR       AS112TimeseriesGroupResponseCodesParamsQueryType = "AXFR"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCAA        AS112TimeseriesGroupResponseCodesParamsQueryType = "CAA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCdnskey    AS112TimeseriesGroupResponseCodesParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCds        AS112TimeseriesGroupResponseCodesParamsQueryType = "CDS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCERT       AS112TimeseriesGroupResponseCodesParamsQueryType = "CERT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCNAME      AS112TimeseriesGroupResponseCodesParamsQueryType = "CNAME"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeCsync      AS112TimeseriesGroupResponseCodesParamsQueryType = "CSYNC"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDhcid      AS112TimeseriesGroupResponseCodesParamsQueryType = "DHCID"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDlv        AS112TimeseriesGroupResponseCodesParamsQueryType = "DLV"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDname      AS112TimeseriesGroupResponseCodesParamsQueryType = "DNAME"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDNSKEY     AS112TimeseriesGroupResponseCodesParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDoa        AS112TimeseriesGroupResponseCodesParamsQueryType = "DOA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeDS         AS112TimeseriesGroupResponseCodesParamsQueryType = "DS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeEid        AS112TimeseriesGroupResponseCodesParamsQueryType = "EID"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeEui48      AS112TimeseriesGroupResponseCodesParamsQueryType = "EUI48"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeEui64      AS112TimeseriesGroupResponseCodesParamsQueryType = "EUI64"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeGpos       AS112TimeseriesGroupResponseCodesParamsQueryType = "GPOS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeGid        AS112TimeseriesGroupResponseCodesParamsQueryType = "GID"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeHinfo      AS112TimeseriesGroupResponseCodesParamsQueryType = "HINFO"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeHip        AS112TimeseriesGroupResponseCodesParamsQueryType = "HIP"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeHTTPS      AS112TimeseriesGroupResponseCodesParamsQueryType = "HTTPS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeIpseckey   AS112TimeseriesGroupResponseCodesParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeIsdn       AS112TimeseriesGroupResponseCodesParamsQueryType = "ISDN"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeIxfr       AS112TimeseriesGroupResponseCodesParamsQueryType = "IXFR"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeKey        AS112TimeseriesGroupResponseCodesParamsQueryType = "KEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeKx         AS112TimeseriesGroupResponseCodesParamsQueryType = "KX"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeL32        AS112TimeseriesGroupResponseCodesParamsQueryType = "L32"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeL64        AS112TimeseriesGroupResponseCodesParamsQueryType = "L64"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeLOC        AS112TimeseriesGroupResponseCodesParamsQueryType = "LOC"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeLp         AS112TimeseriesGroupResponseCodesParamsQueryType = "LP"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMaila      AS112TimeseriesGroupResponseCodesParamsQueryType = "MAILA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMailb      AS112TimeseriesGroupResponseCodesParamsQueryType = "MAILB"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMB         AS112TimeseriesGroupResponseCodesParamsQueryType = "MB"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMd         AS112TimeseriesGroupResponseCodesParamsQueryType = "MD"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMf         AS112TimeseriesGroupResponseCodesParamsQueryType = "MF"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMg         AS112TimeseriesGroupResponseCodesParamsQueryType = "MG"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMinfo      AS112TimeseriesGroupResponseCodesParamsQueryType = "MINFO"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMr         AS112TimeseriesGroupResponseCodesParamsQueryType = "MR"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeMX         AS112TimeseriesGroupResponseCodesParamsQueryType = "MX"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNAPTR      AS112TimeseriesGroupResponseCodesParamsQueryType = "NAPTR"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNb         AS112TimeseriesGroupResponseCodesParamsQueryType = "NB"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNbstat     AS112TimeseriesGroupResponseCodesParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNid        AS112TimeseriesGroupResponseCodesParamsQueryType = "NID"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNimloc     AS112TimeseriesGroupResponseCodesParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNinfo      AS112TimeseriesGroupResponseCodesParamsQueryType = "NINFO"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNS         AS112TimeseriesGroupResponseCodesParamsQueryType = "NS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNsap       AS112TimeseriesGroupResponseCodesParamsQueryType = "NSAP"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec       AS112TimeseriesGroupResponseCodesParamsQueryType = "NSEC"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec3      AS112TimeseriesGroupResponseCodesParamsQueryType = "NSEC3"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec3Param AS112TimeseriesGroupResponseCodesParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNull       AS112TimeseriesGroupResponseCodesParamsQueryType = "NULL"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeNxt        AS112TimeseriesGroupResponseCodesParamsQueryType = "NXT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeOpenpgpkey AS112TimeseriesGroupResponseCodesParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeOpt        AS112TimeseriesGroupResponseCodesParamsQueryType = "OPT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypePTR        AS112TimeseriesGroupResponseCodesParamsQueryType = "PTR"
	AS112TimeseriesGroupResponseCodesParamsQueryTypePx         AS112TimeseriesGroupResponseCodesParamsQueryType = "PX"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeRkey       AS112TimeseriesGroupResponseCodesParamsQueryType = "RKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeRp         AS112TimeseriesGroupResponseCodesParamsQueryType = "RP"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeRrsig      AS112TimeseriesGroupResponseCodesParamsQueryType = "RRSIG"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeRt         AS112TimeseriesGroupResponseCodesParamsQueryType = "RT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSig        AS112TimeseriesGroupResponseCodesParamsQueryType = "SIG"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSink       AS112TimeseriesGroupResponseCodesParamsQueryType = "SINK"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSMIMEA     AS112TimeseriesGroupResponseCodesParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSOA        AS112TimeseriesGroupResponseCodesParamsQueryType = "SOA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSPF        AS112TimeseriesGroupResponseCodesParamsQueryType = "SPF"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSRV        AS112TimeseriesGroupResponseCodesParamsQueryType = "SRV"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSSHFP      AS112TimeseriesGroupResponseCodesParamsQueryType = "SSHFP"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeSVCB       AS112TimeseriesGroupResponseCodesParamsQueryType = "SVCB"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTa         AS112TimeseriesGroupResponseCodesParamsQueryType = "TA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTalink     AS112TimeseriesGroupResponseCodesParamsQueryType = "TALINK"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTkey       AS112TimeseriesGroupResponseCodesParamsQueryType = "TKEY"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTLSA       AS112TimeseriesGroupResponseCodesParamsQueryType = "TLSA"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTSIG       AS112TimeseriesGroupResponseCodesParamsQueryType = "TSIG"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeTXT        AS112TimeseriesGroupResponseCodesParamsQueryType = "TXT"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeUinfo      AS112TimeseriesGroupResponseCodesParamsQueryType = "UINFO"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeUID        AS112TimeseriesGroupResponseCodesParamsQueryType = "UID"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeUnspec     AS112TimeseriesGroupResponseCodesParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeURI        AS112TimeseriesGroupResponseCodesParamsQueryType = "URI"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeWks        AS112TimeseriesGroupResponseCodesParamsQueryType = "WKS"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeX25        AS112TimeseriesGroupResponseCodesParamsQueryType = "X25"
	AS112TimeseriesGroupResponseCodesParamsQueryTypeZonemd     AS112TimeseriesGroupResponseCodesParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupResponseCodesParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsQueryTypeA, AS112TimeseriesGroupResponseCodesParamsQueryTypeAAAA, AS112TimeseriesGroupResponseCodesParamsQueryTypeA6, AS112TimeseriesGroupResponseCodesParamsQueryTypeAfsdb, AS112TimeseriesGroupResponseCodesParamsQueryTypeAny, AS112TimeseriesGroupResponseCodesParamsQueryTypeApl, AS112TimeseriesGroupResponseCodesParamsQueryTypeAtma, AS112TimeseriesGroupResponseCodesParamsQueryTypeAXFR, AS112TimeseriesGroupResponseCodesParamsQueryTypeCAA, AS112TimeseriesGroupResponseCodesParamsQueryTypeCdnskey, AS112TimeseriesGroupResponseCodesParamsQueryTypeCds, AS112TimeseriesGroupResponseCodesParamsQueryTypeCERT, AS112TimeseriesGroupResponseCodesParamsQueryTypeCNAME, AS112TimeseriesGroupResponseCodesParamsQueryTypeCsync, AS112TimeseriesGroupResponseCodesParamsQueryTypeDhcid, AS112TimeseriesGroupResponseCodesParamsQueryTypeDlv, AS112TimeseriesGroupResponseCodesParamsQueryTypeDname, AS112TimeseriesGroupResponseCodesParamsQueryTypeDNSKEY, AS112TimeseriesGroupResponseCodesParamsQueryTypeDoa, AS112TimeseriesGroupResponseCodesParamsQueryTypeDS, AS112TimeseriesGroupResponseCodesParamsQueryTypeEid, AS112TimeseriesGroupResponseCodesParamsQueryTypeEui48, AS112TimeseriesGroupResponseCodesParamsQueryTypeEui64, AS112TimeseriesGroupResponseCodesParamsQueryTypeGpos, AS112TimeseriesGroupResponseCodesParamsQueryTypeGid, AS112TimeseriesGroupResponseCodesParamsQueryTypeHinfo, AS112TimeseriesGroupResponseCodesParamsQueryTypeHip, AS112TimeseriesGroupResponseCodesParamsQueryTypeHTTPS, AS112TimeseriesGroupResponseCodesParamsQueryTypeIpseckey, AS112TimeseriesGroupResponseCodesParamsQueryTypeIsdn, AS112TimeseriesGroupResponseCodesParamsQueryTypeIxfr, AS112TimeseriesGroupResponseCodesParamsQueryTypeKey, AS112TimeseriesGroupResponseCodesParamsQueryTypeKx, AS112TimeseriesGroupResponseCodesParamsQueryTypeL32, AS112TimeseriesGroupResponseCodesParamsQueryTypeL64, AS112TimeseriesGroupResponseCodesParamsQueryTypeLOC, AS112TimeseriesGroupResponseCodesParamsQueryTypeLp, AS112TimeseriesGroupResponseCodesParamsQueryTypeMaila, AS112TimeseriesGroupResponseCodesParamsQueryTypeMailb, AS112TimeseriesGroupResponseCodesParamsQueryTypeMB, AS112TimeseriesGroupResponseCodesParamsQueryTypeMd, AS112TimeseriesGroupResponseCodesParamsQueryTypeMf, AS112TimeseriesGroupResponseCodesParamsQueryTypeMg, AS112TimeseriesGroupResponseCodesParamsQueryTypeMinfo, AS112TimeseriesGroupResponseCodesParamsQueryTypeMr, AS112TimeseriesGroupResponseCodesParamsQueryTypeMX, AS112TimeseriesGroupResponseCodesParamsQueryTypeNAPTR, AS112TimeseriesGroupResponseCodesParamsQueryTypeNb, AS112TimeseriesGroupResponseCodesParamsQueryTypeNbstat, AS112TimeseriesGroupResponseCodesParamsQueryTypeNid, AS112TimeseriesGroupResponseCodesParamsQueryTypeNimloc, AS112TimeseriesGroupResponseCodesParamsQueryTypeNinfo, AS112TimeseriesGroupResponseCodesParamsQueryTypeNS, AS112TimeseriesGroupResponseCodesParamsQueryTypeNsap, AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec, AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec3, AS112TimeseriesGroupResponseCodesParamsQueryTypeNsec3Param, AS112TimeseriesGroupResponseCodesParamsQueryTypeNull, AS112TimeseriesGroupResponseCodesParamsQueryTypeNxt, AS112TimeseriesGroupResponseCodesParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupResponseCodesParamsQueryTypeOpt, AS112TimeseriesGroupResponseCodesParamsQueryTypePTR, AS112TimeseriesGroupResponseCodesParamsQueryTypePx, AS112TimeseriesGroupResponseCodesParamsQueryTypeRkey, AS112TimeseriesGroupResponseCodesParamsQueryTypeRp, AS112TimeseriesGroupResponseCodesParamsQueryTypeRrsig, AS112TimeseriesGroupResponseCodesParamsQueryTypeRt, AS112TimeseriesGroupResponseCodesParamsQueryTypeSig, AS112TimeseriesGroupResponseCodesParamsQueryTypeSink, AS112TimeseriesGroupResponseCodesParamsQueryTypeSMIMEA, AS112TimeseriesGroupResponseCodesParamsQueryTypeSOA, AS112TimeseriesGroupResponseCodesParamsQueryTypeSPF, AS112TimeseriesGroupResponseCodesParamsQueryTypeSRV, AS112TimeseriesGroupResponseCodesParamsQueryTypeSSHFP, AS112TimeseriesGroupResponseCodesParamsQueryTypeSVCB, AS112TimeseriesGroupResponseCodesParamsQueryTypeTa, AS112TimeseriesGroupResponseCodesParamsQueryTypeTalink, AS112TimeseriesGroupResponseCodesParamsQueryTypeTkey, AS112TimeseriesGroupResponseCodesParamsQueryTypeTLSA, AS112TimeseriesGroupResponseCodesParamsQueryTypeTSIG, AS112TimeseriesGroupResponseCodesParamsQueryTypeTXT, AS112TimeseriesGroupResponseCodesParamsQueryTypeUinfo, AS112TimeseriesGroupResponseCodesParamsQueryTypeUID, AS112TimeseriesGroupResponseCodesParamsQueryTypeUnspec, AS112TimeseriesGroupResponseCodesParamsQueryTypeURI, AS112TimeseriesGroupResponseCodesParamsQueryTypeWks, AS112TimeseriesGroupResponseCodesParamsQueryTypeX25, AS112TimeseriesGroupResponseCodesParamsQueryTypeZonemd:
		return true
	}
	return false
}

type AS112TimeseriesGroupResponseCodesResponseEnvelope struct {
	Result  AS112TimeseriesGroupResponseCodesResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    as112TimeseriesGroupResponseCodesResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseEnvelopeJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupResponseCodesResponseEnvelope]
type as112TimeseriesGroupResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
