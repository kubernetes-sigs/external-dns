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

// BGPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPService] method instead.
type BGPService struct {
	Options []option.RequestOption
	Leaks   *BGPLeakService
	Top     *BGPTopService
	Hijacks *BGPHijackService
	Routes  *BGPRouteService
	IPs     *BGPIPService
}

// NewBGPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPService(opts ...option.RequestOption) (r *BGPService) {
	r = &BGPService{}
	r.Options = opts
	r.Leaks = NewBGPLeakService(opts...)
	r.Top = NewBGPTopService(opts...)
	r.Hijacks = NewBGPHijackService(opts...)
	r.Routes = NewBGPRouteService(opts...)
	r.IPs = NewBGPIPService(opts...)
	return
}

// Retrieves BGP updates over time. When requesting updates for an autonomous
// system, only BGP updates of type announcement are returned.
func (r *BGPService) Timeseries(ctx context.Context, query BGPTimeseriesParams, opts ...option.RequestOption) (res *BGPTimeseriesResponse, err error) {
	var env BGPTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPTimeseriesResponse struct {
	Meta   BGPTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 BGPTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   bgpTimeseriesResponseJSON   `json:"-"`
}

// bgpTimeseriesResponseJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponse]
type bgpTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMeta struct {
	AggInterval    BGPTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo BGPTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BGPTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                               `json:"lastUpdated,required" format:"date-time"`
	JSON           bgpTimeseriesResponseMetaJSON           `json:"-"`
}

// bgpTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseMeta]
type bgpTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMetaAggInterval string

const (
	BGPTimeseriesResponseMetaAggInterval15m BGPTimeseriesResponseMetaAggInterval = "15m"
	BGPTimeseriesResponseMetaAggInterval1h  BGPTimeseriesResponseMetaAggInterval = "1h"
	BGPTimeseriesResponseMetaAggInterval1d  BGPTimeseriesResponseMetaAggInterval = "1d"
	BGPTimeseriesResponseMetaAggInterval1w  BGPTimeseriesResponseMetaAggInterval = "1w"
)

func (r BGPTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case BGPTimeseriesResponseMetaAggInterval15m, BGPTimeseriesResponseMetaAggInterval1h, BGPTimeseriesResponseMetaAggInterval1d, BGPTimeseriesResponseMetaAggInterval1w:
		return true
	}
	return false
}

type BGPTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []BGPTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level,required"`
	JSON  bgpTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// bgpTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [BGPTimeseriesResponseMetaConfidenceInfo]
type bgpTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BGPTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate,required" format:"date-time"`
	JSON            bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [BGPTimeseriesResponseMetaConfidenceInfoAnnotation]
type bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BGPTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      bgpTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// bgpTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseMetaDateRange]
type bgpTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                     `json:"timestamps,required" format:"date-time"`
	Values     []string                        `json:"values,required"`
	JSON       bgpTimeseriesResponseSerie0JSON `json:"-"`
}

// bgpTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseSerie0]
type bgpTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[BGPTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BGPTimeseriesParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by BGP network prefix.
	Prefix param.Field[[]string] `query:"prefix"`
	// Filters results by BGP update type.
	UpdateType param.Field[[]BGPTimeseriesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [BGPTimeseriesParams]'s query parameters as `url.Values`.
func (r BGPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BGPTimeseriesParamsAggInterval string

const (
	BGPTimeseriesParamsAggInterval15m BGPTimeseriesParamsAggInterval = "15m"
	BGPTimeseriesParamsAggInterval1h  BGPTimeseriesParamsAggInterval = "1h"
	BGPTimeseriesParamsAggInterval1d  BGPTimeseriesParamsAggInterval = "1d"
	BGPTimeseriesParamsAggInterval1w  BGPTimeseriesParamsAggInterval = "1w"
)

func (r BGPTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsAggInterval15m, BGPTimeseriesParamsAggInterval1h, BGPTimeseriesParamsAggInterval1d, BGPTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type BGPTimeseriesParamsFormat string

const (
	BGPTimeseriesParamsFormatJson BGPTimeseriesParamsFormat = "JSON"
	BGPTimeseriesParamsFormatCsv  BGPTimeseriesParamsFormat = "CSV"
)

func (r BGPTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsFormatJson, BGPTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BGPTimeseriesParamsUpdateType string

const (
	BGPTimeseriesParamsUpdateTypeAnnouncement BGPTimeseriesParamsUpdateType = "ANNOUNCEMENT"
	BGPTimeseriesParamsUpdateTypeWithdrawal   BGPTimeseriesParamsUpdateType = "WITHDRAWAL"
)

func (r BGPTimeseriesParamsUpdateType) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsUpdateTypeAnnouncement, BGPTimeseriesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}

type BGPTimeseriesResponseEnvelope struct {
	Result  BGPTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    bgpTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseEnvelope]
type bgpTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
