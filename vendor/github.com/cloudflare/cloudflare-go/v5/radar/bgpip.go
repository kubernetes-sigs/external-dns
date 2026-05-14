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

// BGPIPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPIPService] method instead.
type BGPIPService struct {
	Options []option.RequestOption
}

// NewBGPIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPIPService(opts ...option.RequestOption) (r *BGPIPService) {
	r = &BGPIPService{}
	r.Options = opts
	return
}

// Retrieves time series data for the announced IP space count, represented as the
// number of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *BGPIPService) Timeseries(ctx context.Context, query BGPIPTimeseriesParams, opts ...option.RequestOption) (res *BgpipTimeseriesResponse, err error) {
	var env BgpipTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/ips/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BgpipTimeseriesResponse struct {
	// Metadata for the results.
	Meta   BgpipTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 BgpipTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   bgpipTimeseriesResponseJSON   `json:"-"`
}

// bgpipTimeseriesResponseJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponse]
type bgpipTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BgpipTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    BgpipTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo BgpipTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BgpipTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BgpipTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BgpipTimeseriesResponseMetaUnit `json:"units,required"`
	Delay BgpipTimeseriesResponseMetaDelay  `json:"delay"`
	JSON  bgpipTimeseriesResponseMetaJSON   `json:"-"`
}

// bgpipTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMeta]
type bgpipTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	Delay          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BgpipTimeseriesResponseMetaAggInterval string

const (
	BgpipTimeseriesResponseMetaAggIntervalFifteenMinutes BgpipTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	BgpipTimeseriesResponseMetaAggIntervalOneHour        BgpipTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	BgpipTimeseriesResponseMetaAggIntervalOneDay         BgpipTimeseriesResponseMetaAggInterval = "ONE_DAY"
	BgpipTimeseriesResponseMetaAggIntervalOneWeek        BgpipTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	BgpipTimeseriesResponseMetaAggIntervalOneMonth       BgpipTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r BgpipTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case BgpipTimeseriesResponseMetaAggIntervalFifteenMinutes, BgpipTimeseriesResponseMetaAggIntervalOneHour, BgpipTimeseriesResponseMetaAggIntervalOneDay, BgpipTimeseriesResponseMetaAggIntervalOneWeek, BgpipTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type BgpipTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []BgpipTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  bgpipTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [BgpipTimeseriesResponseMetaConfidenceInfo]
type bgpipTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BgpipTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            bgpipTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [BgpipTimeseriesResponseMetaConfidenceInfoAnnotation]
type bgpipTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BgpipTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      bgpipTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BgpipTimeseriesResponseMetaDateRange]
type bgpipTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BgpipTimeseriesResponseMetaNormalization string

const (
	BgpipTimeseriesResponseMetaNormalizationPercentage           BgpipTimeseriesResponseMetaNormalization = "PERCENTAGE"
	BgpipTimeseriesResponseMetaNormalizationMin0Max              BgpipTimeseriesResponseMetaNormalization = "MIN0_MAX"
	BgpipTimeseriesResponseMetaNormalizationMinMax               BgpipTimeseriesResponseMetaNormalization = "MIN_MAX"
	BgpipTimeseriesResponseMetaNormalizationRawValues            BgpipTimeseriesResponseMetaNormalization = "RAW_VALUES"
	BgpipTimeseriesResponseMetaNormalizationPercentageChange     BgpipTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BgpipTimeseriesResponseMetaNormalizationRollingAverage       BgpipTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	BgpipTimeseriesResponseMetaNormalizationOverlappedPercentage BgpipTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BgpipTimeseriesResponseMetaNormalizationRatio                BgpipTimeseriesResponseMetaNormalization = "RATIO"
)

func (r BgpipTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BgpipTimeseriesResponseMetaNormalizationPercentage, BgpipTimeseriesResponseMetaNormalizationMin0Max, BgpipTimeseriesResponseMetaNormalizationMinMax, BgpipTimeseriesResponseMetaNormalizationRawValues, BgpipTimeseriesResponseMetaNormalizationPercentageChange, BgpipTimeseriesResponseMetaNormalizationRollingAverage, BgpipTimeseriesResponseMetaNormalizationOverlappedPercentage, BgpipTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BgpipTimeseriesResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  bgpipTimeseriesResponseMetaUnitJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMetaUnit]
type bgpipTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDelay struct {
	ASNData     BgpipTimeseriesResponseMetaDelayASNData     `json:"asn_data,required"`
	CountryData BgpipTimeseriesResponseMetaDelayCountryData `json:"country_data,required"`
	Healthy     bool                                        `json:"healthy,required"`
	NowTs       float64                                     `json:"nowTs,required"`
	JSON        bgpipTimeseriesResponseMetaDelayJSON        `json:"-"`
}

// bgpipTimeseriesResponseMetaDelayJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMetaDelay]
type bgpipTimeseriesResponseMetaDelayJSON struct {
	ASNData     apijson.Field
	CountryData apijson.Field
	Healthy     apijson.Field
	NowTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDelayJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDelayASNData struct {
	DelaySecs float64                                       `json:"delaySecs,required"`
	DelayStr  string                                        `json:"delayStr,required"`
	Healthy   bool                                          `json:"healthy,required"`
	Latest    BgpipTimeseriesResponseMetaDelayASNDataLatest `json:"latest,required"`
	JSON      bgpipTimeseriesResponseMetaDelayASNDataJSON   `json:"-"`
}

// bgpipTimeseriesResponseMetaDelayASNDataJSON contains the JSON metadata for the
// struct [BgpipTimeseriesResponseMetaDelayASNData]
type bgpipTimeseriesResponseMetaDelayASNDataJSON struct {
	DelaySecs   apijson.Field
	DelayStr    apijson.Field
	Healthy     apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDelayASNData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDelayASNDataJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDelayASNDataLatest struct {
	EntriesCount float64                                           `json:"entries_count,required"`
	Path         string                                            `json:"path,required"`
	Timestamp    float64                                           `json:"timestamp,required"`
	JSON         bgpipTimeseriesResponseMetaDelayASNDataLatestJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaDelayASNDataLatestJSON contains the JSON metadata for
// the struct [BgpipTimeseriesResponseMetaDelayASNDataLatest]
type bgpipTimeseriesResponseMetaDelayASNDataLatestJSON struct {
	EntriesCount apijson.Field
	Path         apijson.Field
	Timestamp    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDelayASNDataLatest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDelayASNDataLatestJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDelayCountryData struct {
	DelaySecs float64                                           `json:"delaySecs,required"`
	DelayStr  string                                            `json:"delayStr,required"`
	Healthy   bool                                              `json:"healthy,required"`
	Latest    BgpipTimeseriesResponseMetaDelayCountryDataLatest `json:"latest,required"`
	JSON      bgpipTimeseriesResponseMetaDelayCountryDataJSON   `json:"-"`
}

// bgpipTimeseriesResponseMetaDelayCountryDataJSON contains the JSON metadata for
// the struct [BgpipTimeseriesResponseMetaDelayCountryData]
type bgpipTimeseriesResponseMetaDelayCountryDataJSON struct {
	DelaySecs   apijson.Field
	DelayStr    apijson.Field
	Healthy     apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDelayCountryData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDelayCountryDataJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaDelayCountryDataLatest struct {
	Count     float64                                               `json:"count,required"`
	Timestamp float64                                               `json:"timestamp,required"`
	JSON      bgpipTimeseriesResponseMetaDelayCountryDataLatestJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaDelayCountryDataLatestJSON contains the JSON metadata
// for the struct [BgpipTimeseriesResponseMetaDelayCountryDataLatest]
type bgpipTimeseriesResponseMetaDelayCountryDataLatestJSON struct {
	Count       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaDelayCountryDataLatest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaDelayCountryDataLatestJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseSerie0 struct {
	IPV4       []string                          `json:"ipv4,required"`
	IPV6       []string                          `json:"ipv6,required"`
	Timestamps []time.Time                       `json:"timestamps,required" format:"date-time"`
	JSON       bgpipTimeseriesResponseSerie0JSON `json:"-"`
}

// bgpipTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseSerie0]
type bgpipTimeseriesResponseSerie0JSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPIPTimeseriesParams struct {
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
	Format param.Field[BgpipTimeseriesParamsFormat] `query:"format"`
	// Includes data delay meta information.
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]BgpipTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BGPIPTimeseriesParams]'s query parameters as `url.Values`.
func (r BGPIPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgpipTimeseriesParamsFormat string

const (
	BgpipTimeseriesParamsFormatJson BgpipTimeseriesParamsFormat = "JSON"
	BgpipTimeseriesParamsFormatCsv  BgpipTimeseriesParamsFormat = "CSV"
)

func (r BgpipTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BgpipTimeseriesParamsFormatJson, BgpipTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BgpipTimeseriesParamsIPVersion string

const (
	BgpipTimeseriesParamsIPVersionIPv4 BgpipTimeseriesParamsIPVersion = "IPv4"
	BgpipTimeseriesParamsIPVersionIPv6 BgpipTimeseriesParamsIPVersion = "IPv6"
)

func (r BgpipTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case BgpipTimeseriesParamsIPVersionIPv4, BgpipTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type BgpipTimeseriesResponseEnvelope struct {
	Result  BgpipTimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    bgpipTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpipTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseEnvelope]
type bgpipTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
