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

// QualitySpeedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualitySpeedService] method instead.
type QualitySpeedService struct {
	Options []option.RequestOption
	Top     *QualitySpeedTopService
}

// NewQualitySpeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualitySpeedService(opts ...option.RequestOption) (r *QualitySpeedService) {
	r = &QualitySpeedService{}
	r.Options = opts
	r.Top = NewQualitySpeedTopService(opts...)
	return
}

// Retrieves a histogram from the previous 90 days of Cloudflare Speed Test data,
// split into fixed bandwidth (Mbps), latency (ms), or jitter (ms) buckets.
func (r *QualitySpeedService) Histogram(ctx context.Context, query QualitySpeedHistogramParams, opts ...option.RequestOption) (res *QualitySpeedHistogramResponse, err error) {
	var env QualitySpeedHistogramResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a summary of bandwidth, latency, jitter, and packet loss, from the
// previous 90 days of Cloudflare Speed Test data.
func (r *QualitySpeedService) Summary(ctx context.Context, query QualitySpeedSummaryParams, opts ...option.RequestOption) (res *QualitySpeedSummaryResponse, err error) {
	var env QualitySpeedSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QualitySpeedHistogramResponse struct {
	Histogram0 QualitySpeedHistogramResponseHistogram0 `json:"histogram_0,required"`
	// Metadata for the results.
	Meta QualitySpeedHistogramResponseMeta `json:"meta,required"`
	JSON qualitySpeedHistogramResponseJSON `json:"-"`
}

// qualitySpeedHistogramResponseJSON contains the JSON metadata for the struct
// [QualitySpeedHistogramResponse]
type qualitySpeedHistogramResponseJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramResponseHistogram0 struct {
	BandwidthDownload []string                                    `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                    `json:"bandwidthUpload,required"`
	BucketMin         []string                                    `json:"bucketMin,required"`
	JSON              qualitySpeedHistogramResponseHistogram0JSON `json:"-"`
}

// qualitySpeedHistogramResponseHistogram0JSON contains the JSON metadata for the
// struct [QualitySpeedHistogramResponseHistogram0]
type qualitySpeedHistogramResponseHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseHistogram0JSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type QualitySpeedHistogramResponseMeta struct {
	// The width for every bucket in the histogram.
	BucketSize     int64                                           `json:"bucketSize,required"`
	ConfidenceInfo QualitySpeedHistogramResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []QualitySpeedHistogramResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization QualitySpeedHistogramResponseMetaNormalization `json:"normalization,required"`
	TotalTests    []int64                                        `json:"totalTests,required"`
	// Measurement units for the results.
	Units []QualitySpeedHistogramResponseMetaUnit `json:"units,required"`
	JSON  qualitySpeedHistogramResponseMetaJSON   `json:"-"`
}

// qualitySpeedHistogramResponseMetaJSON contains the JSON metadata for the struct
// [QualitySpeedHistogramResponseMeta]
type qualitySpeedHistogramResponseMetaJSON struct {
	BucketSize     apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	TotalTests     apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedHistogramResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  qualitySpeedHistogramResponseMetaConfidenceInfoJSON `json:"-"`
}

// qualitySpeedHistogramResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [QualitySpeedHistogramResponseMetaConfidenceInfo]
type qualitySpeedHistogramResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type QualitySpeedHistogramResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            qualitySpeedHistogramResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedHistogramResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [QualitySpeedHistogramResponseMetaConfidenceInfoAnnotation]
type qualitySpeedHistogramResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedHistogramResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedHistogramResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedHistogramResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [QualitySpeedHistogramResponseMetaDateRange]
type qualitySpeedHistogramResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type QualitySpeedHistogramResponseMetaNormalization string

const (
	QualitySpeedHistogramResponseMetaNormalizationPercentage           QualitySpeedHistogramResponseMetaNormalization = "PERCENTAGE"
	QualitySpeedHistogramResponseMetaNormalizationMin0Max              QualitySpeedHistogramResponseMetaNormalization = "MIN0_MAX"
	QualitySpeedHistogramResponseMetaNormalizationMinMax               QualitySpeedHistogramResponseMetaNormalization = "MIN_MAX"
	QualitySpeedHistogramResponseMetaNormalizationRawValues            QualitySpeedHistogramResponseMetaNormalization = "RAW_VALUES"
	QualitySpeedHistogramResponseMetaNormalizationPercentageChange     QualitySpeedHistogramResponseMetaNormalization = "PERCENTAGE_CHANGE"
	QualitySpeedHistogramResponseMetaNormalizationRollingAverage       QualitySpeedHistogramResponseMetaNormalization = "ROLLING_AVERAGE"
	QualitySpeedHistogramResponseMetaNormalizationOverlappedPercentage QualitySpeedHistogramResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	QualitySpeedHistogramResponseMetaNormalizationRatio                QualitySpeedHistogramResponseMetaNormalization = "RATIO"
)

func (r QualitySpeedHistogramResponseMetaNormalization) IsKnown() bool {
	switch r {
	case QualitySpeedHistogramResponseMetaNormalizationPercentage, QualitySpeedHistogramResponseMetaNormalizationMin0Max, QualitySpeedHistogramResponseMetaNormalizationMinMax, QualitySpeedHistogramResponseMetaNormalizationRawValues, QualitySpeedHistogramResponseMetaNormalizationPercentageChange, QualitySpeedHistogramResponseMetaNormalizationRollingAverage, QualitySpeedHistogramResponseMetaNormalizationOverlappedPercentage, QualitySpeedHistogramResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type QualitySpeedHistogramResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  qualitySpeedHistogramResponseMetaUnitJSON `json:"-"`
}

// qualitySpeedHistogramResponseMetaUnitJSON contains the JSON metadata for the
// struct [QualitySpeedHistogramResponseMetaUnit]
type qualitySpeedHistogramResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponse struct {
	// Metadata for the results.
	Meta     QualitySpeedSummaryResponseMeta     `json:"meta,required"`
	Summary0 QualitySpeedSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     qualitySpeedSummaryResponseJSON     `json:"-"`
}

// qualitySpeedSummaryResponseJSON contains the JSON metadata for the struct
// [QualitySpeedSummaryResponse]
type qualitySpeedSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type QualitySpeedSummaryResponseMeta struct {
	ConfidenceInfo QualitySpeedSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []QualitySpeedSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization QualitySpeedSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []QualitySpeedSummaryResponseMetaUnit `json:"units,required"`
	JSON  qualitySpeedSummaryResponseMetaJSON   `json:"-"`
}

// qualitySpeedSummaryResponseMetaJSON contains the JSON metadata for the struct
// [QualitySpeedSummaryResponseMeta]
type qualitySpeedSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  qualitySpeedSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [QualitySpeedSummaryResponseMetaConfidenceInfo]
type qualitySpeedSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation]
type qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedSummaryResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseMetaDateRange]
type qualitySpeedSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type QualitySpeedSummaryResponseMetaNormalization string

const (
	QualitySpeedSummaryResponseMetaNormalizationPercentage           QualitySpeedSummaryResponseMetaNormalization = "PERCENTAGE"
	QualitySpeedSummaryResponseMetaNormalizationMin0Max              QualitySpeedSummaryResponseMetaNormalization = "MIN0_MAX"
	QualitySpeedSummaryResponseMetaNormalizationMinMax               QualitySpeedSummaryResponseMetaNormalization = "MIN_MAX"
	QualitySpeedSummaryResponseMetaNormalizationRawValues            QualitySpeedSummaryResponseMetaNormalization = "RAW_VALUES"
	QualitySpeedSummaryResponseMetaNormalizationPercentageChange     QualitySpeedSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	QualitySpeedSummaryResponseMetaNormalizationRollingAverage       QualitySpeedSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	QualitySpeedSummaryResponseMetaNormalizationOverlappedPercentage QualitySpeedSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	QualitySpeedSummaryResponseMetaNormalizationRatio                QualitySpeedSummaryResponseMetaNormalization = "RATIO"
)

func (r QualitySpeedSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case QualitySpeedSummaryResponseMetaNormalizationPercentage, QualitySpeedSummaryResponseMetaNormalizationMin0Max, QualitySpeedSummaryResponseMetaNormalizationMinMax, QualitySpeedSummaryResponseMetaNormalizationRawValues, QualitySpeedSummaryResponseMetaNormalizationPercentageChange, QualitySpeedSummaryResponseMetaNormalizationRollingAverage, QualitySpeedSummaryResponseMetaNormalizationOverlappedPercentage, QualitySpeedSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type QualitySpeedSummaryResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  qualitySpeedSummaryResponseMetaUnitJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseMetaUnit]
type qualitySpeedSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseSummary0 struct {
	BandwidthDownload string                                  `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                  `json:"bandwidthUpload,required"`
	JitterIdle        string                                  `json:"jitterIdle,required"`
	JitterLoaded      string                                  `json:"jitterLoaded,required"`
	LatencyIdle       string                                  `json:"latencyIdle,required"`
	LatencyLoaded     string                                  `json:"latencyLoaded,required"`
	PacketLoss        string                                  `json:"packetLoss,required"`
	JSON              qualitySpeedSummaryResponseSummary0JSON `json:"-"`
}

// qualitySpeedSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseSummary0]
type qualitySpeedSummaryResponseSummary0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	PacketLoss        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Specifies the width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[QualitySpeedHistogramParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[QualitySpeedHistogramParamsMetricGroup] `query:"metricGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualitySpeedHistogramParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedHistogramParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type QualitySpeedHistogramParamsFormat string

const (
	QualitySpeedHistogramParamsFormatJson QualitySpeedHistogramParamsFormat = "JSON"
	QualitySpeedHistogramParamsFormatCsv  QualitySpeedHistogramParamsFormat = "CSV"
)

func (r QualitySpeedHistogramParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedHistogramParamsFormatJson, QualitySpeedHistogramParamsFormatCsv:
		return true
	}
	return false
}

// Metrics to be returned.
type QualitySpeedHistogramParamsMetricGroup string

const (
	QualitySpeedHistogramParamsMetricGroupBandwidth QualitySpeedHistogramParamsMetricGroup = "BANDWIDTH"
	QualitySpeedHistogramParamsMetricGroupLatency   QualitySpeedHistogramParamsMetricGroup = "LATENCY"
	QualitySpeedHistogramParamsMetricGroupJitter    QualitySpeedHistogramParamsMetricGroup = "JITTER"
)

func (r QualitySpeedHistogramParamsMetricGroup) IsKnown() bool {
	switch r {
	case QualitySpeedHistogramParamsMetricGroupBandwidth, QualitySpeedHistogramParamsMetricGroupLatency, QualitySpeedHistogramParamsMetricGroupJitter:
		return true
	}
	return false
}

type QualitySpeedHistogramResponseEnvelope struct {
	Result  QualitySpeedHistogramResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    qualitySpeedHistogramResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedHistogramResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedHistogramResponseEnvelope]
type qualitySpeedHistogramResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryParams struct {
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
	Format param.Field[QualitySpeedSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualitySpeedSummaryParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type QualitySpeedSummaryParamsFormat string

const (
	QualitySpeedSummaryParamsFormatJson QualitySpeedSummaryParamsFormat = "JSON"
	QualitySpeedSummaryParamsFormatCsv  QualitySpeedSummaryParamsFormat = "CSV"
)

func (r QualitySpeedSummaryParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedSummaryParamsFormatJson, QualitySpeedSummaryParamsFormatCsv:
		return true
	}
	return false
}

type QualitySpeedSummaryResponseEnvelope struct {
	Result  QualitySpeedSummaryResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    qualitySpeedSummaryResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseEnvelope]
type qualitySpeedSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
