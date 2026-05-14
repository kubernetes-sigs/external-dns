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

// QualitySpeedTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualitySpeedTopService] method instead.
type QualitySpeedTopService struct {
	Options []option.RequestOption
}

// NewQualitySpeedTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualitySpeedTopService(opts ...option.RequestOption) (r *QualitySpeedTopService) {
	r = &QualitySpeedTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by bandwidth, latency, jitter, or packet
// loss, from the previous 90 days of Cloudflare Speed Test data.
func (r *QualitySpeedTopService) Ases(ctx context.Context, query QualitySpeedTopAsesParams, opts ...option.RequestOption) (res *QualitySpeedTopAsesResponse, err error) {
	var env QualitySpeedTopAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by bandwidth, latency, jitter, or packet loss, from
// the previous 90 days of Cloudflare Speed Test data.
func (r *QualitySpeedTopService) Locations(ctx context.Context, query QualitySpeedTopLocationsParams, opts ...option.RequestOption) (res *QualitySpeedTopLocationsResponse, err error) {
	var env QualitySpeedTopLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QualitySpeedTopAsesResponse struct {
	// Metadata for the results.
	Meta QualitySpeedTopAsesResponseMeta   `json:"meta,required"`
	Top0 []QualitySpeedTopAsesResponseTop0 `json:"top_0,required"`
	JSON qualitySpeedTopAsesResponseJSON   `json:"-"`
}

// qualitySpeedTopAsesResponseJSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponse]
type qualitySpeedTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type QualitySpeedTopAsesResponseMeta struct {
	ConfidenceInfo QualitySpeedTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []QualitySpeedTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization QualitySpeedTopAsesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []QualitySpeedTopAsesResponseMetaUnit `json:"units,required"`
	JSON  qualitySpeedTopAsesResponseMetaJSON   `json:"-"`
}

// qualitySpeedTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponseMeta]
type qualitySpeedTopAsesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  qualitySpeedTopAsesResponseMetaConfidenceInfoJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [QualitySpeedTopAsesResponseMetaConfidenceInfo]
type qualitySpeedTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation]
type qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [QualitySpeedTopAsesResponseMetaDateRange]
type qualitySpeedTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type QualitySpeedTopAsesResponseMetaNormalization string

const (
	QualitySpeedTopAsesResponseMetaNormalizationPercentage           QualitySpeedTopAsesResponseMetaNormalization = "PERCENTAGE"
	QualitySpeedTopAsesResponseMetaNormalizationMin0Max              QualitySpeedTopAsesResponseMetaNormalization = "MIN0_MAX"
	QualitySpeedTopAsesResponseMetaNormalizationMinMax               QualitySpeedTopAsesResponseMetaNormalization = "MIN_MAX"
	QualitySpeedTopAsesResponseMetaNormalizationRawValues            QualitySpeedTopAsesResponseMetaNormalization = "RAW_VALUES"
	QualitySpeedTopAsesResponseMetaNormalizationPercentageChange     QualitySpeedTopAsesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	QualitySpeedTopAsesResponseMetaNormalizationRollingAverage       QualitySpeedTopAsesResponseMetaNormalization = "ROLLING_AVERAGE"
	QualitySpeedTopAsesResponseMetaNormalizationOverlappedPercentage QualitySpeedTopAsesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	QualitySpeedTopAsesResponseMetaNormalizationRatio                QualitySpeedTopAsesResponseMetaNormalization = "RATIO"
)

func (r QualitySpeedTopAsesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case QualitySpeedTopAsesResponseMetaNormalizationPercentage, QualitySpeedTopAsesResponseMetaNormalizationMin0Max, QualitySpeedTopAsesResponseMetaNormalizationMinMax, QualitySpeedTopAsesResponseMetaNormalizationRawValues, QualitySpeedTopAsesResponseMetaNormalizationPercentageChange, QualitySpeedTopAsesResponseMetaNormalizationRollingAverage, QualitySpeedTopAsesResponseMetaNormalizationOverlappedPercentage, QualitySpeedTopAsesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type QualitySpeedTopAsesResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  qualitySpeedTopAsesResponseMetaUnitJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaUnitJSON contains the JSON metadata for the
// struct [QualitySpeedTopAsesResponseMetaUnit]
type qualitySpeedTopAsesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseTop0 struct {
	BandwidthDownload string                              `json:"bandwidthDownload,required"`
	BandwidthUpload   string                              `json:"bandwidthUpload,required"`
	ClientASN         float64                             `json:"clientASN,required"`
	ClientAsName      string                              `json:"clientASName,required"`
	JitterIdle        string                              `json:"jitterIdle,required"`
	JitterLoaded      string                              `json:"jitterLoaded,required"`
	LatencyIdle       string                              `json:"latencyIdle,required"`
	LatencyLoaded     string                              `json:"latencyLoaded,required"`
	NumTests          float64                             `json:"numTests,required"`
	RankPower         float64                             `json:"rankPower,required"`
	JSON              qualitySpeedTopAsesResponseTop0JSON `json:"-"`
}

// qualitySpeedTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponseTop0]
type qualitySpeedTopAsesResponseTop0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	ClientASN         apijson.Field
	ClientAsName      apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	NumTests          apijson.Field
	RankPower         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponse struct {
	// Metadata for the results.
	Meta QualitySpeedTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []QualitySpeedTopLocationsResponseTop0 `json:"top_0,required"`
	JSON qualitySpeedTopLocationsResponseJSON   `json:"-"`
}

// qualitySpeedTopLocationsResponseJSON contains the JSON metadata for the struct
// [QualitySpeedTopLocationsResponse]
type qualitySpeedTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type QualitySpeedTopLocationsResponseMeta struct {
	ConfidenceInfo QualitySpeedTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []QualitySpeedTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization QualitySpeedTopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []QualitySpeedTopLocationsResponseMetaUnit `json:"units,required"`
	JSON  qualitySpeedTopLocationsResponseMetaJSON   `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaJSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseMeta]
type qualitySpeedTopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [QualitySpeedTopLocationsResponseMetaConfidenceInfo]
type qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation]
type qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [QualitySpeedTopLocationsResponseMetaDateRange]
type qualitySpeedTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type QualitySpeedTopLocationsResponseMetaNormalization string

const (
	QualitySpeedTopLocationsResponseMetaNormalizationPercentage           QualitySpeedTopLocationsResponseMetaNormalization = "PERCENTAGE"
	QualitySpeedTopLocationsResponseMetaNormalizationMin0Max              QualitySpeedTopLocationsResponseMetaNormalization = "MIN0_MAX"
	QualitySpeedTopLocationsResponseMetaNormalizationMinMax               QualitySpeedTopLocationsResponseMetaNormalization = "MIN_MAX"
	QualitySpeedTopLocationsResponseMetaNormalizationRawValues            QualitySpeedTopLocationsResponseMetaNormalization = "RAW_VALUES"
	QualitySpeedTopLocationsResponseMetaNormalizationPercentageChange     QualitySpeedTopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	QualitySpeedTopLocationsResponseMetaNormalizationRollingAverage       QualitySpeedTopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	QualitySpeedTopLocationsResponseMetaNormalizationOverlappedPercentage QualitySpeedTopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	QualitySpeedTopLocationsResponseMetaNormalizationRatio                QualitySpeedTopLocationsResponseMetaNormalization = "RATIO"
)

func (r QualitySpeedTopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case QualitySpeedTopLocationsResponseMetaNormalizationPercentage, QualitySpeedTopLocationsResponseMetaNormalizationMin0Max, QualitySpeedTopLocationsResponseMetaNormalizationMinMax, QualitySpeedTopLocationsResponseMetaNormalizationRawValues, QualitySpeedTopLocationsResponseMetaNormalizationPercentageChange, QualitySpeedTopLocationsResponseMetaNormalizationRollingAverage, QualitySpeedTopLocationsResponseMetaNormalizationOverlappedPercentage, QualitySpeedTopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type QualitySpeedTopLocationsResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  qualitySpeedTopLocationsResponseMetaUnitJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaUnitJSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseMetaUnit]
type qualitySpeedTopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseTop0 struct {
	BandwidthDownload   string                                   `json:"bandwidthDownload,required"`
	BandwidthUpload     string                                   `json:"bandwidthUpload,required"`
	ClientCountryAlpha2 string                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                   `json:"clientCountryName,required"`
	JitterIdle          string                                   `json:"jitterIdle,required"`
	JitterLoaded        string                                   `json:"jitterLoaded,required"`
	LatencyIdle         string                                   `json:"latencyIdle,required"`
	LatencyLoaded       string                                   `json:"latencyLoaded,required"`
	NumTests            float64                                  `json:"numTests,required"`
	RankPower           float64                                  `json:"rankPower,required"`
	JSON                qualitySpeedTopLocationsResponseTop0JSON `json:"-"`
}

// qualitySpeedTopLocationsResponseTop0JSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseTop0]
type qualitySpeedTopLocationsResponseTop0JSON struct {
	BandwidthDownload   apijson.Field
	BandwidthUpload     apijson.Field
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	JitterIdle          apijson.Field
	JitterLoaded        apijson.Field
	LatencyIdle         apijson.Field
	LatencyLoaded       apijson.Field
	NumTests            apijson.Field
	RankPower           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesParams struct {
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
	// Format in which results will be returned.
	Format param.Field[QualitySpeedTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies the metric to order the results by.
	OrderBy param.Field[QualitySpeedTopAsesParamsOrderBy] `query:"orderBy"`
	// Reverses the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [QualitySpeedTopAsesParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type QualitySpeedTopAsesParamsFormat string

const (
	QualitySpeedTopAsesParamsFormatJson QualitySpeedTopAsesParamsFormat = "JSON"
	QualitySpeedTopAsesParamsFormatCsv  QualitySpeedTopAsesParamsFormat = "CSV"
)

func (r QualitySpeedTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedTopAsesParamsFormatJson, QualitySpeedTopAsesParamsFormatCsv:
		return true
	}
	return false
}

// Specifies the metric to order the results by.
type QualitySpeedTopAsesParamsOrderBy string

const (
	QualitySpeedTopAsesParamsOrderByBandwidthDownload QualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	QualitySpeedTopAsesParamsOrderByBandwidthUpload   QualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_UPLOAD"
	QualitySpeedTopAsesParamsOrderByLatencyIdle       QualitySpeedTopAsesParamsOrderBy = "LATENCY_IDLE"
	QualitySpeedTopAsesParamsOrderByLatencyLoaded     QualitySpeedTopAsesParamsOrderBy = "LATENCY_LOADED"
	QualitySpeedTopAsesParamsOrderByJitterIdle        QualitySpeedTopAsesParamsOrderBy = "JITTER_IDLE"
	QualitySpeedTopAsesParamsOrderByJitterLoaded      QualitySpeedTopAsesParamsOrderBy = "JITTER_LOADED"
)

func (r QualitySpeedTopAsesParamsOrderBy) IsKnown() bool {
	switch r {
	case QualitySpeedTopAsesParamsOrderByBandwidthDownload, QualitySpeedTopAsesParamsOrderByBandwidthUpload, QualitySpeedTopAsesParamsOrderByLatencyIdle, QualitySpeedTopAsesParamsOrderByLatencyLoaded, QualitySpeedTopAsesParamsOrderByJitterIdle, QualitySpeedTopAsesParamsOrderByJitterLoaded:
		return true
	}
	return false
}

type QualitySpeedTopAsesResponseEnvelope struct {
	Result  QualitySpeedTopAsesResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    qualitySpeedTopAsesResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedTopAsesResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedTopAsesResponseEnvelope]
type qualitySpeedTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsParams struct {
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
	// Format in which results will be returned.
	Format param.Field[QualitySpeedTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies the metric to order the results by.
	OrderBy param.Field[QualitySpeedTopLocationsParamsOrderBy] `query:"orderBy"`
	// Reverses the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [QualitySpeedTopLocationsParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type QualitySpeedTopLocationsParamsFormat string

const (
	QualitySpeedTopLocationsParamsFormatJson QualitySpeedTopLocationsParamsFormat = "JSON"
	QualitySpeedTopLocationsParamsFormatCsv  QualitySpeedTopLocationsParamsFormat = "CSV"
)

func (r QualitySpeedTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedTopLocationsParamsFormatJson, QualitySpeedTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

// Specifies the metric to order the results by.
type QualitySpeedTopLocationsParamsOrderBy string

const (
	QualitySpeedTopLocationsParamsOrderByBandwidthDownload QualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	QualitySpeedTopLocationsParamsOrderByBandwidthUpload   QualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_UPLOAD"
	QualitySpeedTopLocationsParamsOrderByLatencyIdle       QualitySpeedTopLocationsParamsOrderBy = "LATENCY_IDLE"
	QualitySpeedTopLocationsParamsOrderByLatencyLoaded     QualitySpeedTopLocationsParamsOrderBy = "LATENCY_LOADED"
	QualitySpeedTopLocationsParamsOrderByJitterIdle        QualitySpeedTopLocationsParamsOrderBy = "JITTER_IDLE"
	QualitySpeedTopLocationsParamsOrderByJitterLoaded      QualitySpeedTopLocationsParamsOrderBy = "JITTER_LOADED"
)

func (r QualitySpeedTopLocationsParamsOrderBy) IsKnown() bool {
	switch r {
	case QualitySpeedTopLocationsParamsOrderByBandwidthDownload, QualitySpeedTopLocationsParamsOrderByBandwidthUpload, QualitySpeedTopLocationsParamsOrderByLatencyIdle, QualitySpeedTopLocationsParamsOrderByLatencyLoaded, QualitySpeedTopLocationsParamsOrderByJitterIdle, QualitySpeedTopLocationsParamsOrderByJitterLoaded:
		return true
	}
	return false
}

type QualitySpeedTopLocationsResponseEnvelope struct {
	Result  QualitySpeedTopLocationsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    qualitySpeedTopLocationsResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseEnvelope]
type qualitySpeedTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
