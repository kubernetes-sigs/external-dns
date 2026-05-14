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

// TCPResetsTimeoutService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTCPResetsTimeoutService] method instead.
type TCPResetsTimeoutService struct {
	Options []option.RequestOption
}

// NewTCPResetsTimeoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTCPResetsTimeoutService(opts ...option.RequestOption) (r *TCPResetsTimeoutService) {
	r = &TCPResetsTimeoutService{}
	r.Options = opts
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout.
func (r *TCPResetsTimeoutService) Summary(ctx context.Context, query TCPResetsTimeoutSummaryParams, opts ...option.RequestOption) (res *TCPResetsTimeoutSummaryResponse, err error) {
	var env TCPResetsTimeoutSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout over time.
func (r *TCPResetsTimeoutService) TimeseriesGroups(ctx context.Context, query TCPResetsTimeoutTimeseriesGroupsParams, opts ...option.RequestOption) (res *TCPResetsTimeoutTimeseriesGroupsResponse, err error) {
	var env TCPResetsTimeoutTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TCPResetsTimeoutSummaryResponse struct {
	// Metadata for the results.
	Meta     TCPResetsTimeoutSummaryResponseMeta     `json:"meta,required"`
	Summary0 TCPResetsTimeoutSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     tcpResetsTimeoutSummaryResponseJSON     `json:"-"`
}

// tcpResetsTimeoutSummaryResponseJSON contains the JSON metadata for the struct
// [TCPResetsTimeoutSummaryResponse]
type tcpResetsTimeoutSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type TCPResetsTimeoutSummaryResponseMeta struct {
	ConfidenceInfo TCPResetsTimeoutSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []TCPResetsTimeoutSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization TCPResetsTimeoutSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []TCPResetsTimeoutSummaryResponseMetaUnit `json:"units,required"`
	JSON  tcpResetsTimeoutSummaryResponseMetaJSON   `json:"-"`
}

// tcpResetsTimeoutSummaryResponseMetaJSON contains the JSON metadata for the
// struct [TCPResetsTimeoutSummaryResponseMeta]
type tcpResetsTimeoutSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutSummaryResponseMetaConfidenceInfo struct {
	Annotations []TCPResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  tcpResetsTimeoutSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [TCPResetsTimeoutSummaryResponseMetaConfidenceInfo]
type tcpResetsTimeoutSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type TCPResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            tcpResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [TCPResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotation]
type tcpResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *TCPResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      tcpResetsTimeoutSummaryResponseMetaDateRangeJSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [TCPResetsTimeoutSummaryResponseMetaDateRange]
type tcpResetsTimeoutSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type TCPResetsTimeoutSummaryResponseMetaNormalization string

const (
	TCPResetsTimeoutSummaryResponseMetaNormalizationPercentage           TCPResetsTimeoutSummaryResponseMetaNormalization = "PERCENTAGE"
	TCPResetsTimeoutSummaryResponseMetaNormalizationMin0Max              TCPResetsTimeoutSummaryResponseMetaNormalization = "MIN0_MAX"
	TCPResetsTimeoutSummaryResponseMetaNormalizationMinMax               TCPResetsTimeoutSummaryResponseMetaNormalization = "MIN_MAX"
	TCPResetsTimeoutSummaryResponseMetaNormalizationRawValues            TCPResetsTimeoutSummaryResponseMetaNormalization = "RAW_VALUES"
	TCPResetsTimeoutSummaryResponseMetaNormalizationPercentageChange     TCPResetsTimeoutSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	TCPResetsTimeoutSummaryResponseMetaNormalizationRollingAverage       TCPResetsTimeoutSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	TCPResetsTimeoutSummaryResponseMetaNormalizationOverlappedPercentage TCPResetsTimeoutSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	TCPResetsTimeoutSummaryResponseMetaNormalizationRatio                TCPResetsTimeoutSummaryResponseMetaNormalization = "RATIO"
)

func (r TCPResetsTimeoutSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutSummaryResponseMetaNormalizationPercentage, TCPResetsTimeoutSummaryResponseMetaNormalizationMin0Max, TCPResetsTimeoutSummaryResponseMetaNormalizationMinMax, TCPResetsTimeoutSummaryResponseMetaNormalizationRawValues, TCPResetsTimeoutSummaryResponseMetaNormalizationPercentageChange, TCPResetsTimeoutSummaryResponseMetaNormalizationRollingAverage, TCPResetsTimeoutSummaryResponseMetaNormalizationOverlappedPercentage, TCPResetsTimeoutSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type TCPResetsTimeoutSummaryResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  tcpResetsTimeoutSummaryResponseMetaUnitJSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [TCPResetsTimeoutSummaryResponseMetaUnit]
type tcpResetsTimeoutSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutSummaryResponseSummary0 struct {
	// Connection resets within the first 10 packets from the client, but after the
	// server has received multiple data packets.
	LaterInFlow string `json:"later_in_flow,required"`
	// All other connections.
	NoMatch string `json:"no_match,required"`
	// Connection resets or timeouts after the server received both a SYN packet and an
	// ACK packet, meaning the connection was successfully established.
	PostAck string `json:"post_ack,required"`
	// Connection resets or timeouts after the server received a packet with PSH flag
	// set, following connection establishment.
	PostPsh string `json:"post_psh,required"`
	// Connection resets or timeouts after the server received only a single SYN
	// packet.
	PostSyn string                                      `json:"post_syn,required"`
	JSON    tcpResetsTimeoutSummaryResponseSummary0JSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [TCPResetsTimeoutSummaryResponseSummary0]
type tcpResetsTimeoutSummaryResponseSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   TCPResetsTimeoutTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 TCPResetsTimeoutTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   tcpResetsTimeoutTimeseriesGroupsResponseJSON   `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [TCPResetsTimeoutTimeseriesGroupsResponse]
type tcpResetsTimeoutTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type TCPResetsTimeoutTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []TCPResetsTimeoutTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []TCPResetsTimeoutTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  tcpResetsTimeoutTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseMetaJSON contains the JSON metadata for
// the struct [TCPResetsTimeoutTimeseriesGroupsResponseMeta]
type tcpResetsTimeoutTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval string

const (
	TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneHour        TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneDay         TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneWeek        TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneMonth       TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r TCPResetsTimeoutTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneHour, TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneDay, TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneWeek, TCPResetsTimeoutTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfo]
type tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *TCPResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      tcpResetsTimeoutTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [TCPResetsTimeoutTimeseriesGroupsResponseMetaDateRange]
type tcpResetsTimeoutTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization string

const (
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationPercentage           TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationMin0Max              TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationMinMax               TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRawValues            TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationPercentageChange     TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRollingAverage       TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRatio                TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationPercentage, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationMin0Max, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationMinMax, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRawValues, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationPercentageChange, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRollingAverage, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, TCPResetsTimeoutTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type TCPResetsTimeoutTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  tcpResetsTimeoutTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata
// for the struct [TCPResetsTimeoutTimeseriesGroupsResponseMetaUnit]
type tcpResetsTimeoutTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutTimeseriesGroupsResponseSerie0 struct {
	LaterInFlow []string                                           `json:"later_in_flow,required"`
	NoMatch     []string                                           `json:"no_match,required"`
	PostAck     []string                                           `json:"post_ack,required"`
	PostPsh     []string                                           `json:"post_psh,required"`
	PostSyn     []string                                           `json:"post_syn,required"`
	Timestamps  []time.Time                                        `json:"timestamps,required" format:"date-time"`
	JSON        tcpResetsTimeoutTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseSerie0JSON contains the JSON metadata
// for the struct [TCPResetsTimeoutTimeseriesGroupsResponseSerie0]
type tcpResetsTimeoutTimeseriesGroupsResponseSerie0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutSummaryParams struct {
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
	Format param.Field[TCPResetsTimeoutSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [TCPResetsTimeoutSummaryParams]'s query parameters as
// `url.Values`.
func (r TCPResetsTimeoutSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type TCPResetsTimeoutSummaryParamsFormat string

const (
	TCPResetsTimeoutSummaryParamsFormatJson TCPResetsTimeoutSummaryParamsFormat = "JSON"
	TCPResetsTimeoutSummaryParamsFormatCsv  TCPResetsTimeoutSummaryParamsFormat = "CSV"
)

func (r TCPResetsTimeoutSummaryParamsFormat) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutSummaryParamsFormatJson, TCPResetsTimeoutSummaryParamsFormatCsv:
		return true
	}
	return false
}

type TCPResetsTimeoutSummaryResponseEnvelope struct {
	Result  TCPResetsTimeoutSummaryResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    tcpResetsTimeoutSummaryResponseEnvelopeJSON `json:"-"`
}

// tcpResetsTimeoutSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [TCPResetsTimeoutSummaryResponseEnvelope]
type tcpResetsTimeoutSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TCPResetsTimeoutTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[TCPResetsTimeoutTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[TCPResetsTimeoutTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [TCPResetsTimeoutTimeseriesGroupsParams]'s query parameters
// as `url.Values`.
func (r TCPResetsTimeoutTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type TCPResetsTimeoutTimeseriesGroupsParamsAggInterval string

const (
	TCPResetsTimeoutTimeseriesGroupsParamsAggInterval15m TCPResetsTimeoutTimeseriesGroupsParamsAggInterval = "15m"
	TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1h  TCPResetsTimeoutTimeseriesGroupsParamsAggInterval = "1h"
	TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1d  TCPResetsTimeoutTimeseriesGroupsParamsAggInterval = "1d"
	TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1w  TCPResetsTimeoutTimeseriesGroupsParamsAggInterval = "1w"
)

func (r TCPResetsTimeoutTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutTimeseriesGroupsParamsAggInterval15m, TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1h, TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1d, TCPResetsTimeoutTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type TCPResetsTimeoutTimeseriesGroupsParamsFormat string

const (
	TCPResetsTimeoutTimeseriesGroupsParamsFormatJson TCPResetsTimeoutTimeseriesGroupsParamsFormat = "JSON"
	TCPResetsTimeoutTimeseriesGroupsParamsFormatCsv  TCPResetsTimeoutTimeseriesGroupsParamsFormat = "CSV"
)

func (r TCPResetsTimeoutTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case TCPResetsTimeoutTimeseriesGroupsParamsFormatJson, TCPResetsTimeoutTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type TCPResetsTimeoutTimeseriesGroupsResponseEnvelope struct {
	Result  TCPResetsTimeoutTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    tcpResetsTimeoutTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// tcpResetsTimeoutTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata
// for the struct [TCPResetsTimeoutTimeseriesGroupsResponseEnvelope]
type tcpResetsTimeoutTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TCPResetsTimeoutTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tcpResetsTimeoutTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
