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

// HTTPTimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPTimeseriesGroupService] method instead.
type HTTPTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewHTTPTimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPTimeseriesGroupService(opts ...option.RequestOption) (r *HTTPTimeseriesGroupService) {
	r = &HTTPTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP requests classified as automated or human
// over time. Visit https://developers.cloudflare.com/radar/concepts/bot-classes/
// for more information.
func (r *HTTPTimeseriesGroupService) BotClass(ctx context.Context, query HTTPTimeseriesGroupBotClassParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBotClassResponse, err error) {
	var env HTTPTimeseriesGroupBotClassResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by user agent over time.
func (r *HTTPTimeseriesGroupService) Browser(ctx context.Context, query HTTPTimeseriesGroupBrowserParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBrowserResponse, err error) {
	var env HTTPTimeseriesGroupBrowserResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by user agent family over time.
func (r *HTTPTimeseriesGroupService) BrowserFamily(ctx context.Context, query HTTPTimeseriesGroupBrowserFamilyParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupBrowserFamilyResponse, err error) {
	var env HTTPTimeseriesGroupBrowserFamilyResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by device type over time.
func (r *HTTPTimeseriesGroupService) DeviceType(ctx context.Context, query HTTPTimeseriesGroupDeviceTypeParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupDeviceTypeResponse, err error) {
	var env HTTPTimeseriesGroupDeviceTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by HTTP protocol (HTTP vs. HTTPS)
// over time.
func (r *HTTPTimeseriesGroupService) HTTPProtocol(ctx context.Context, query HTTPTimeseriesGroupHTTPProtocolParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupHTTPProtocolResponse, err error) {
	var env HTTPTimeseriesGroupHTTPProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by HTTP version over time.
func (r *HTTPTimeseriesGroupService) HTTPVersion(ctx context.Context, query HTTPTimeseriesGroupHTTPVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupHTTPVersionResponse, err error) {
	var env HTTPTimeseriesGroupHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by IP version over time.
func (r *HTTPTimeseriesGroupService) IPVersion(ctx context.Context, query HTTPTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupIPVersionResponse, err error) {
	var env HTTPTimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by operating system over time.
func (r *HTTPTimeseriesGroupService) OS(ctx context.Context, query HTTPTimeseriesGroupOSParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupOSResponse, err error) {
	var env HTTPTimeseriesGroupOSResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by post-quantum support over time.
func (r *HTTPTimeseriesGroupService) PostQuantum(ctx context.Context, query HTTPTimeseriesGroupPostQuantumParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupPostQuantumResponse, err error) {
	var env HTTPTimeseriesGroupPostQuantumResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/post_quantum"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests by TLS version over time.
func (r *HTTPTimeseriesGroupService) TLSVersion(ctx context.Context, query HTTPTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *HTTPTimeseriesGroupTLSVersionResponse, err error) {
	var env HTTPTimeseriesGroupTLSVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTimeseriesGroupBotClassResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupBotClassResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBotClassResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBotClassResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBotClassResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBotClassResponse]
type httpTimeseriesGroupBotClassResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupBotClassResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupBotClassResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupBotClassResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupBotClassResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupBotClassResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupBotClassResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupBotClassResponseMetaJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBotClassResponseMeta]
type httpTimeseriesGroupBotClassResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBotClassResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupBotClassResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupBotClassResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupBotClassResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupBotClassResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupBotClassResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupBotClassResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupBotClassResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupBotClassResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  httpTimeseriesGroupBotClassResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfo]
type httpTimeseriesGroupBotClassResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBotClassResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupBotClassResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBotClassResponseMetaDateRange]
type httpTimeseriesGroupBotClassResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupBotClassResponseMetaNormalization string

const (
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationPercentage           HTTPTimeseriesGroupBotClassResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupBotClassResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationMinMax               HTTPTimeseriesGroupBotClassResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationRawValues            HTTPTimeseriesGroupBotClassResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupBotClassResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupBotClassResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupBotClassResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupBotClassResponseMetaNormalizationRatio                HTTPTimeseriesGroupBotClassResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupBotClassResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassResponseMetaNormalizationPercentage, HTTPTimeseriesGroupBotClassResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupBotClassResponseMetaNormalizationMinMax, HTTPTimeseriesGroupBotClassResponseMetaNormalizationRawValues, HTTPTimeseriesGroupBotClassResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupBotClassResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupBotClassResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupBotClassResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  httpTimeseriesGroupBotClassResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBotClassResponseMetaUnit]
type httpTimeseriesGroupBotClassResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBotClassResponseSerie0 struct {
	Bot        []string                                      `json:"bot,required"`
	Human      []string                                      `json:"human,required"`
	Timestamps []time.Time                                   `json:"timestamps,required" format:"date-time"`
	JSON       httpTimeseriesGroupBotClassResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBotClassResponseSerie0]
type httpTimeseriesGroupBotClassResponseSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupBrowserResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBrowserResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBrowserResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupBrowserResponse]
type httpTimeseriesGroupBrowserResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupBrowserResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupBrowserResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupBrowserResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupBrowserResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupBrowserResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupBrowserResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserResponseMetaJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBrowserResponseMeta]
type httpTimeseriesGroupBrowserResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupBrowserResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupBrowserResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupBrowserResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupBrowserResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupBrowserResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupBrowserResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupBrowserResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupBrowserResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  httpTimeseriesGroupBrowserResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfo]
type httpTimeseriesGroupBrowserResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupBrowserResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserResponseMetaDateRange]
type httpTimeseriesGroupBrowserResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupBrowserResponseMetaNormalization string

const (
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationPercentage           HTTPTimeseriesGroupBrowserResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupBrowserResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationMinMax               HTTPTimeseriesGroupBrowserResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationRawValues            HTTPTimeseriesGroupBrowserResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupBrowserResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupBrowserResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupBrowserResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupBrowserResponseMetaNormalizationRatio                HTTPTimeseriesGroupBrowserResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupBrowserResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserResponseMetaNormalizationPercentage, HTTPTimeseriesGroupBrowserResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupBrowserResponseMetaNormalizationMinMax, HTTPTimeseriesGroupBrowserResponseMetaNormalizationRawValues, HTTPTimeseriesGroupBrowserResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupBrowserResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupBrowserResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupBrowserResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  httpTimeseriesGroupBrowserResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBrowserResponseMetaUnit]
type httpTimeseriesGroupBrowserResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserResponseSerie0 struct {
	Timestamps  []time.Time                                  `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                          `json:"-,extras"`
	JSON        httpTimeseriesGroupBrowserResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBrowserResponseSerie0]
type httpTimeseriesGroupBrowserResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupBrowserFamilyResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupBrowserFamilyResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupBrowserFamilyResponseJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupBrowserFamilyResponse]
type httpTimeseriesGroupBrowserFamilyResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupBrowserFamilyResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupBrowserFamilyResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupBrowserFamilyResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupBrowserFamilyResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseMetaJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBrowserFamilyResponseMeta]
type httpTimeseriesGroupBrowserFamilyResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupBrowserFamilyResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupBrowserFamilyResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfo]
type httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupBrowserFamilyResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupBrowserFamilyResponseMetaDateRange]
type httpTimeseriesGroupBrowserFamilyResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization string

const (
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationPercentage           HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationMinMax               HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRawValues            HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRatio                HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationPercentage, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationMinMax, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRawValues, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupBrowserFamilyResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  httpTimeseriesGroupBrowserFamilyResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseMetaUnitJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserFamilyResponseMetaUnit]
type httpTimeseriesGroupBrowserFamilyResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyResponseSerie0 struct {
	Timestamps  []time.Time                                        `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                `json:"-,extras"`
	JSON        httpTimeseriesGroupBrowserFamilyResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseSerie0JSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserFamilyResponseSerie0]
type httpTimeseriesGroupBrowserFamilyResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupDeviceTypeResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupDeviceTypeResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupDeviceTypeResponseJSON   `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupDeviceTypeResponse]
type httpTimeseriesGroupDeviceTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupDeviceTypeResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupDeviceTypeResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupDeviceTypeResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupDeviceTypeResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseMetaJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupDeviceTypeResponseMeta]
type httpTimeseriesGroupDeviceTypeResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupDeviceTypeResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupDeviceTypeResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfo]
type httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupDeviceTypeResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupDeviceTypeResponseMetaDateRange]
type httpTimeseriesGroupDeviceTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization string

const (
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationPercentage           HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationMinMax               HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRawValues            HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRatio                HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupDeviceTypeResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationPercentage, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationMinMax, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRawValues, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupDeviceTypeResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  httpTimeseriesGroupDeviceTypeResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupDeviceTypeResponseMetaUnit]
type httpTimeseriesGroupDeviceTypeResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeResponseSerie0 struct {
	Desktop    []string                                        `json:"desktop,required"`
	Mobile     []string                                        `json:"mobile,required"`
	Other      []string                                        `json:"other,required"`
	Timestamps []time.Time                                     `json:"timestamps,required" format:"date-time"`
	JSON       httpTimeseriesGroupDeviceTypeResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupDeviceTypeResponseSerie0]
type httpTimeseriesGroupDeviceTypeResponseSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupHTTPProtocolResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupHTTPProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupHTTPProtocolResponseJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupHTTPProtocolResponse]
type httpTimeseriesGroupHTTPProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupHTTPProtocolResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupHTTPProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupHTTPProtocolResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupHTTPProtocolResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseMetaJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPProtocolResponseMeta]
type httpTimeseriesGroupHTTPProtocolResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupHTTPProtocolResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupHTTPProtocolResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfo]
type httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupHTTPProtocolResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupHTTPProtocolResponseMetaDateRange]
type httpTimeseriesGroupHTTPProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization string

const (
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationPercentage           HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationMinMax               HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRawValues            HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRatio                HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationPercentage, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationMinMax, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRawValues, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupHTTPProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  httpTimeseriesGroupHTTPProtocolResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseMetaUnitJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPProtocolResponseMetaUnit]
type httpTimeseriesGroupHTTPProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolResponseSerie0 struct {
	HTTP       []string                                          `json:"http,required"`
	HTTPS      []string                                          `json:"https,required"`
	Timestamps []time.Time                                       `json:"timestamps,required" format:"date-time"`
	JSON       httpTimeseriesGroupHTTPProtocolResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPProtocolResponseSerie0]
type httpTimeseriesGroupHTTPProtocolResponseSerie0JSON struct {
	HTTP        apijson.Field
	HTTPS       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupHTTPVersionResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupHTTPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupHTTPVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupHTTPVersionResponse]
type httpTimeseriesGroupHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupHTTPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupHTTPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupHTTPVersionResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupHTTPVersionResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseMetaJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPVersionResponseMeta]
type httpTimeseriesGroupHTTPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupHTTPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupHTTPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfo]
type httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupHTTPVersionResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupHTTPVersionResponseMetaDateRange]
type httpTimeseriesGroupHTTPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization string

const (
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationPercentage           HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationMinMax               HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRawValues            HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRatio                HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupHTTPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationPercentage, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationMinMax, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRawValues, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupHTTPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  httpTimeseriesGroupHTTPVersionResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseMetaUnitJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPVersionResponseMetaUnit]
type httpTimeseriesGroupHTTPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionResponseSerie0 struct {
	HTTP1X     []string                                         `json:"HTTP/1.x,required"`
	HTTP2      []string                                         `json:"HTTP/2,required"`
	HTTP3      []string                                         `json:"HTTP/3,required"`
	Timestamps []time.Time                                      `json:"timestamps,required" format:"date-time"`
	JSON       httpTimeseriesGroupHTTPVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupHTTPVersionResponseSerie0]
type httpTimeseriesGroupHTTPVersionResponseSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupIPVersionResponse]
type httpTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupIPVersionResponseMeta]
type httpTimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupIPVersionResponseMetaDateRange]
type httpTimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupIPVersionResponseMetaNormalization string

const (
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationPercentage           HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationMinMax               HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRawValues            HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRatio                HTTPTimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionResponseMetaNormalizationPercentage, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationMinMax, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRawValues, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  httpTimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupIPVersionResponseMetaUnit]
type httpTimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4       []string                                       `json:"IPv4,required"`
	IPv6       []string                                       `json:"IPv6,required"`
	Timestamps []time.Time                                    `json:"timestamps,required" format:"date-time"`
	JSON       httpTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupIPVersionResponseSerie0]
type httpTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupOSResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupOSResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupOSResponseJSON   `json:"-"`
}

// httpTimeseriesGroupOSResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupOSResponse]
type httpTimeseriesGroupOSResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupOSResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupOSResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupOSResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupOSResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupOSResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupOSResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupOSResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupOSResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupOSResponseMeta]
type httpTimeseriesGroupOSResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupOSResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupOSResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupOSResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupOSResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupOSResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupOSResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupOSResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupOSResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupOSResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupOSResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupOSResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupOSResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupOSResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupOSResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupOSResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupOSResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupOSResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  httpTimeseriesGroupOSResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupOSResponseMetaConfidenceInfo]
type httpTimeseriesGroupOSResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupOSResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupOSResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupOSResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupOSResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupOSResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupOSResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupOSResponseMetaDateRange]
type httpTimeseriesGroupOSResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupOSResponseMetaNormalization string

const (
	HTTPTimeseriesGroupOSResponseMetaNormalizationPercentage           HTTPTimeseriesGroupOSResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupOSResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupOSResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupOSResponseMetaNormalizationMinMax               HTTPTimeseriesGroupOSResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupOSResponseMetaNormalizationRawValues            HTTPTimeseriesGroupOSResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupOSResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupOSResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupOSResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupOSResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupOSResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupOSResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupOSResponseMetaNormalizationRatio                HTTPTimeseriesGroupOSResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupOSResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSResponseMetaNormalizationPercentage, HTTPTimeseriesGroupOSResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupOSResponseMetaNormalizationMinMax, HTTPTimeseriesGroupOSResponseMetaNormalizationRawValues, HTTPTimeseriesGroupOSResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupOSResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupOSResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupOSResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  httpTimeseriesGroupOSResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupOSResponseMetaUnit]
type httpTimeseriesGroupOSResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSResponseSerie0 struct {
	Timestamps  []time.Time                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                     `json:"-,extras"`
	JSON        httpTimeseriesGroupOSResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupOSResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupOSResponseSerie0]
type httpTimeseriesGroupOSResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupPostQuantumResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupPostQuantumResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupPostQuantumResponseJSON   `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupPostQuantumResponse]
type httpTimeseriesGroupPostQuantumResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupPostQuantumResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupPostQuantumResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupPostQuantumResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupPostQuantumResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupPostQuantumResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseMetaJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupPostQuantumResponseMeta]
type httpTimeseriesGroupPostQuantumResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupPostQuantumResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupPostQuantumResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfo]
type httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupPostQuantumResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupPostQuantumResponseMetaDateRange]
type httpTimeseriesGroupPostQuantumResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupPostQuantumResponseMetaNormalization string

const (
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationPercentage           HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationMinMax               HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRawValues            HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRatio                HTTPTimeseriesGroupPostQuantumResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupPostQuantumResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationPercentage, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationMinMax, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRawValues, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupPostQuantumResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  httpTimeseriesGroupPostQuantumResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseMetaUnitJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupPostQuantumResponseMetaUnit]
type httpTimeseriesGroupPostQuantumResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumResponseSerie0 struct {
	NotSupported []string                                         `json:"NOT_SUPPORTED,required"`
	Supported    []string                                         `json:"SUPPORTED,required"`
	Timestamps   []time.Time                                      `json:"timestamps,required" format:"date-time"`
	JSON         httpTimeseriesGroupPostQuantumResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupPostQuantumResponseSerie0]
type httpTimeseriesGroupPostQuantumResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	Timestamps   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionResponse struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupTLSVersionResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupTLSVersionResponse]
type httpTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupTLSVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupTLSVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupTLSVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupTLSVersionResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupTLSVersionResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseMetaJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupTLSVersionResponseMeta]
type httpTimeseriesGroupTLSVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupTLSVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupTLSVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfo]
type httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupTLSVersionResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [HTTPTimeseriesGroupTLSVersionResponseMetaDateRange]
type httpTimeseriesGroupTLSVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupTLSVersionResponseMetaNormalization string

const (
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationPercentage           HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationMinMax               HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRawValues            HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRatio                HTTPTimeseriesGroupTLSVersionResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupTLSVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationPercentage, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationMinMax, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRawValues, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupTLSVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  httpTimeseriesGroupTLSVersionResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupTLSVersionResponseMetaUnit]
type httpTimeseriesGroupTLSVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionResponseSerie0 struct {
	Timestamps []time.Time                                     `json:"timestamps,required" format:"date-time"`
	TLS1_0     []string                                        `json:"TLS 1.0,required"`
	TLS1_1     []string                                        `json:"TLS 1.1,required"`
	TLS1_2     []string                                        `json:"TLS 1.2,required"`
	TLS1_3     []string                                        `json:"TLS 1.3,required"`
	TLSQuic    []string                                        `json:"TLS QUIC,required"`
	JSON       httpTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupTLSVersionResponseSerie0]
type httpTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	Timestamps  apijson.Field
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	TLSQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBotClassParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBotClassParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupBotClassParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBotClassParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupBotClassParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupBotClassParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBotClassParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBotClassParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBotClassParamsAggInterval string

const (
	HTTPTimeseriesGroupBotClassParamsAggInterval15m HTTPTimeseriesGroupBotClassParamsAggInterval = "15m"
	HTTPTimeseriesGroupBotClassParamsAggInterval1h  HTTPTimeseriesGroupBotClassParamsAggInterval = "1h"
	HTTPTimeseriesGroupBotClassParamsAggInterval1d  HTTPTimeseriesGroupBotClassParamsAggInterval = "1d"
	HTTPTimeseriesGroupBotClassParamsAggInterval1w  HTTPTimeseriesGroupBotClassParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBotClassParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsAggInterval15m, HTTPTimeseriesGroupBotClassParamsAggInterval1h, HTTPTimeseriesGroupBotClassParamsAggInterval1d, HTTPTimeseriesGroupBotClassParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsBrowserFamily string

const (
	HTTPTimeseriesGroupBotClassParamsBrowserFamilyChrome  HTTPTimeseriesGroupBotClassParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupBotClassParamsBrowserFamilyEdge    HTTPTimeseriesGroupBotClassParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupBotClassParamsBrowserFamilyFirefox HTTPTimeseriesGroupBotClassParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupBotClassParamsBrowserFamilySafari  HTTPTimeseriesGroupBotClassParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupBotClassParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsBrowserFamilyChrome, HTTPTimeseriesGroupBotClassParamsBrowserFamilyEdge, HTTPTimeseriesGroupBotClassParamsBrowserFamilyFirefox, HTTPTimeseriesGroupBotClassParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsDeviceType string

const (
	HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop HTTPTimeseriesGroupBotClassParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile  HTTPTimeseriesGroupBotClassParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBotClassParamsDeviceTypeOther   HTTPTimeseriesGroupBotClassParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBotClassParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsDeviceTypeDesktop, HTTPTimeseriesGroupBotClassParamsDeviceTypeMobile, HTTPTimeseriesGroupBotClassParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupBotClassParamsFormat string

const (
	HTTPTimeseriesGroupBotClassParamsFormatJson HTTPTimeseriesGroupBotClassParamsFormat = "JSON"
	HTTPTimeseriesGroupBotClassParamsFormatCsv  HTTPTimeseriesGroupBotClassParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBotClassParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsFormatJson, HTTPTimeseriesGroupBotClassParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBotClassParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBotClassParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBotClassParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBotClassParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBotClassParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBotClassParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsIPVersion string

const (
	HTTPTimeseriesGroupBotClassParamsIPVersionIPv4 HTTPTimeseriesGroupBotClassParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBotClassParamsIPVersionIPv6 HTTPTimeseriesGroupBotClassParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBotClassParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsIPVersionIPv4, HTTPTimeseriesGroupBotClassParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsOS string

const (
	HTTPTimeseriesGroupBotClassParamsOSWindows  HTTPTimeseriesGroupBotClassParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBotClassParamsOSMacosx   HTTPTimeseriesGroupBotClassParamsOS = "MACOSX"
	HTTPTimeseriesGroupBotClassParamsOSIos      HTTPTimeseriesGroupBotClassParamsOS = "IOS"
	HTTPTimeseriesGroupBotClassParamsOSAndroid  HTTPTimeseriesGroupBotClassParamsOS = "ANDROID"
	HTTPTimeseriesGroupBotClassParamsOSChromeos HTTPTimeseriesGroupBotClassParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBotClassParamsOSLinux    HTTPTimeseriesGroupBotClassParamsOS = "LINUX"
	HTTPTimeseriesGroupBotClassParamsOSSmartTv  HTTPTimeseriesGroupBotClassParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBotClassParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsOSWindows, HTTPTimeseriesGroupBotClassParamsOSMacosx, HTTPTimeseriesGroupBotClassParamsOSIos, HTTPTimeseriesGroupBotClassParamsOSAndroid, HTTPTimeseriesGroupBotClassParamsOSChromeos, HTTPTimeseriesGroupBotClassParamsOSLinux, HTTPTimeseriesGroupBotClassParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassParamsTLSVersion string

const (
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBotClassParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBotClassParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBotClassParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBotClassParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBotClassResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBotClassResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    httpTimeseriesGroupBotClassResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBotClassResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBotClassResponseEnvelope]
type httpTimeseriesGroupBotClassResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBotClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBotClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBrowserParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupBrowserParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupBrowserParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBrowserParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupBrowserParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsIPVersion] `query:"ipVersion"`
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
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupBrowserParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBrowserParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBrowserParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserParamsAggInterval string

const (
	HTTPTimeseriesGroupBrowserParamsAggInterval15m HTTPTimeseriesGroupBrowserParamsAggInterval = "15m"
	HTTPTimeseriesGroupBrowserParamsAggInterval1h  HTTPTimeseriesGroupBrowserParamsAggInterval = "1h"
	HTTPTimeseriesGroupBrowserParamsAggInterval1d  HTTPTimeseriesGroupBrowserParamsAggInterval = "1d"
	HTTPTimeseriesGroupBrowserParamsAggInterval1w  HTTPTimeseriesGroupBrowserParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBrowserParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsAggInterval15m, HTTPTimeseriesGroupBrowserParamsAggInterval1h, HTTPTimeseriesGroupBrowserParamsAggInterval1d, HTTPTimeseriesGroupBrowserParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsBotClass string

const (
	HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated HTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman     HTTPTimeseriesGroupBrowserParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupBrowserParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsBotClassLikelyAutomated, HTTPTimeseriesGroupBrowserParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsBrowserFamily string

const (
	HTTPTimeseriesGroupBrowserParamsBrowserFamilyChrome  HTTPTimeseriesGroupBrowserParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupBrowserParamsBrowserFamilyEdge    HTTPTimeseriesGroupBrowserParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupBrowserParamsBrowserFamilyFirefox HTTPTimeseriesGroupBrowserParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupBrowserParamsBrowserFamilySafari  HTTPTimeseriesGroupBrowserParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupBrowserParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsBrowserFamilyChrome, HTTPTimeseriesGroupBrowserParamsBrowserFamilyEdge, HTTPTimeseriesGroupBrowserParamsBrowserFamilyFirefox, HTTPTimeseriesGroupBrowserParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsDeviceType string

const (
	HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop HTTPTimeseriesGroupBrowserParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile  HTTPTimeseriesGroupBrowserParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBrowserParamsDeviceTypeOther   HTTPTimeseriesGroupBrowserParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBrowserParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsDeviceTypeDesktop, HTTPTimeseriesGroupBrowserParamsDeviceTypeMobile, HTTPTimeseriesGroupBrowserParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupBrowserParamsFormat string

const (
	HTTPTimeseriesGroupBrowserParamsFormatJson HTTPTimeseriesGroupBrowserParamsFormat = "JSON"
	HTTPTimeseriesGroupBrowserParamsFormatCsv  HTTPTimeseriesGroupBrowserParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBrowserParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsFormatJson, HTTPTimeseriesGroupBrowserParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBrowserParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBrowserParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBrowserParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBrowserParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBrowserParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBrowserParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsIPVersion string

const (
	HTTPTimeseriesGroupBrowserParamsIPVersionIPv4 HTTPTimeseriesGroupBrowserParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBrowserParamsIPVersionIPv6 HTTPTimeseriesGroupBrowserParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBrowserParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsIPVersionIPv4, HTTPTimeseriesGroupBrowserParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsOS string

const (
	HTTPTimeseriesGroupBrowserParamsOSWindows  HTTPTimeseriesGroupBrowserParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBrowserParamsOSMacosx   HTTPTimeseriesGroupBrowserParamsOS = "MACOSX"
	HTTPTimeseriesGroupBrowserParamsOSIos      HTTPTimeseriesGroupBrowserParamsOS = "IOS"
	HTTPTimeseriesGroupBrowserParamsOSAndroid  HTTPTimeseriesGroupBrowserParamsOS = "ANDROID"
	HTTPTimeseriesGroupBrowserParamsOSChromeos HTTPTimeseriesGroupBrowserParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBrowserParamsOSLinux    HTTPTimeseriesGroupBrowserParamsOS = "LINUX"
	HTTPTimeseriesGroupBrowserParamsOSSmartTv  HTTPTimeseriesGroupBrowserParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBrowserParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsOSWindows, HTTPTimeseriesGroupBrowserParamsOSMacosx, HTTPTimeseriesGroupBrowserParamsOSIos, HTTPTimeseriesGroupBrowserParamsOSAndroid, HTTPTimeseriesGroupBrowserParamsOSChromeos, HTTPTimeseriesGroupBrowserParamsOSLinux, HTTPTimeseriesGroupBrowserParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserParamsTLSVersion string

const (
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBrowserParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBrowserParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBrowserParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBrowserParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBrowserResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    httpTimeseriesGroupBrowserResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupBrowserResponseEnvelope]
type httpTimeseriesGroupBrowserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupBrowserFamilyParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupBrowserFamilyParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupBrowserFamilyParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsIPVersion] `query:"ipVersion"`
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
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupBrowserFamilyParams]'s query parameters
// as `url.Values`.
func (r HTTPTimeseriesGroupBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupBrowserFamilyParamsAggInterval string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "15m"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1h"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1d  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1d"
	HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1w  HTTPTimeseriesGroupBrowserFamilyParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsAggInterval15m, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1h, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1d, HTTPTimeseriesGroupBrowserFamilyParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsBotClass string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated HTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman     HTTPTimeseriesGroupBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyAutomated, HTTPTimeseriesGroupBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsDeviceType string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile  HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther   HTTPTimeseriesGroupBrowserFamilyParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeDesktop, HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeMobile, HTTPTimeseriesGroupBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupBrowserFamilyParamsFormat string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsFormatJson HTTPTimeseriesGroupBrowserFamilyParamsFormat = "JSON"
	HTTPTimeseriesGroupBrowserFamilyParamsFormatCsv  HTTPTimeseriesGroupBrowserFamilyParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsFormatJson, HTTPTimeseriesGroupBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP  HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTP, HTTPTimeseriesGroupBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3 HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv1, HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv2, HTTPTimeseriesGroupBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsIPVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4 HTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6 HTTPTimeseriesGroupBrowserFamilyParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv4, HTTPTimeseriesGroupBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsOS string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsOSWindows  HTTPTimeseriesGroupBrowserFamilyParamsOS = "WINDOWS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx   HTTPTimeseriesGroupBrowserFamilyParamsOS = "MACOSX"
	HTTPTimeseriesGroupBrowserFamilyParamsOSIos      HTTPTimeseriesGroupBrowserFamilyParamsOS = "IOS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSAndroid  HTTPTimeseriesGroupBrowserFamilyParamsOS = "ANDROID"
	HTTPTimeseriesGroupBrowserFamilyParamsOSChromeos HTTPTimeseriesGroupBrowserFamilyParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupBrowserFamilyParamsOSLinux    HTTPTimeseriesGroupBrowserFamilyParamsOS = "LINUX"
	HTTPTimeseriesGroupBrowserFamilyParamsOSSmartTv  HTTPTimeseriesGroupBrowserFamilyParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsOSWindows, HTTPTimeseriesGroupBrowserFamilyParamsOSMacosx, HTTPTimeseriesGroupBrowserFamilyParamsOSIos, HTTPTimeseriesGroupBrowserFamilyParamsOSAndroid, HTTPTimeseriesGroupBrowserFamilyParamsOSChromeos, HTTPTimeseriesGroupBrowserFamilyParamsOSLinux, HTTPTimeseriesGroupBrowserFamilyParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion string

const (
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSvQuic HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupBrowserFamilyParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupBrowserFamilyParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupBrowserFamilyResponseEnvelope struct {
	Result  HTTPTimeseriesGroupBrowserFamilyResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupBrowserFamilyResponseEnvelope]
type httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupBrowserFamilyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupBrowserFamilyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupDeviceTypeParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupDeviceTypeParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily] `query:"browserFamily"`
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
	Format param.Field[HTTPTimeseriesGroupDeviceTypeParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupDeviceTypeParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupDeviceTypeParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupDeviceTypeParamsAggInterval string

const (
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "15m"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1h  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1h"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1d  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1d"
	HTTPTimeseriesGroupDeviceTypeParamsAggInterval1w  HTTPTimeseriesGroupDeviceTypeParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsAggInterval15m, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1h, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1d, HTTPTimeseriesGroupDeviceTypeParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsBotClass string

const (
	HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated HTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman     HTTPTimeseriesGroupDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyAutomated, HTTPTimeseriesGroupDeviceTypeParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily string

const (
	HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyChrome  HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyEdge    HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyFirefox HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilySafari  HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyChrome, HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyEdge, HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilyFirefox, HTTPTimeseriesGroupDeviceTypeParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupDeviceTypeParamsFormat string

const (
	HTTPTimeseriesGroupDeviceTypeParamsFormatJson HTTPTimeseriesGroupDeviceTypeParamsFormat = "JSON"
	HTTPTimeseriesGroupDeviceTypeParamsFormatCsv  HTTPTimeseriesGroupDeviceTypeParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsFormatJson, HTTPTimeseriesGroupDeviceTypeParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP  HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTP, HTTPTimeseriesGroupDeviceTypeParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3 HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv1, HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv2, HTTPTimeseriesGroupDeviceTypeParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsIPVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4 HTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6 HTTPTimeseriesGroupDeviceTypeParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv4, HTTPTimeseriesGroupDeviceTypeParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsOS string

const (
	HTTPTimeseriesGroupDeviceTypeParamsOSWindows  HTTPTimeseriesGroupDeviceTypeParamsOS = "WINDOWS"
	HTTPTimeseriesGroupDeviceTypeParamsOSMacosx   HTTPTimeseriesGroupDeviceTypeParamsOS = "MACOSX"
	HTTPTimeseriesGroupDeviceTypeParamsOSIos      HTTPTimeseriesGroupDeviceTypeParamsOS = "IOS"
	HTTPTimeseriesGroupDeviceTypeParamsOSAndroid  HTTPTimeseriesGroupDeviceTypeParamsOS = "ANDROID"
	HTTPTimeseriesGroupDeviceTypeParamsOSChromeos HTTPTimeseriesGroupDeviceTypeParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupDeviceTypeParamsOSLinux    HTTPTimeseriesGroupDeviceTypeParamsOS = "LINUX"
	HTTPTimeseriesGroupDeviceTypeParamsOSSmartTv  HTTPTimeseriesGroupDeviceTypeParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsOSWindows, HTTPTimeseriesGroupDeviceTypeParamsOSMacosx, HTTPTimeseriesGroupDeviceTypeParamsOSIos, HTTPTimeseriesGroupDeviceTypeParamsOSAndroid, HTTPTimeseriesGroupDeviceTypeParamsOSChromeos, HTTPTimeseriesGroupDeviceTypeParamsOSLinux, HTTPTimeseriesGroupDeviceTypeParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeParamsTLSVersion string

const (
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSvQuic HTTPTimeseriesGroupDeviceTypeParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupDeviceTypeParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupDeviceTypeParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupDeviceTypeResponseEnvelope struct {
	Result  HTTPTimeseriesGroupDeviceTypeResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupDeviceTypeResponseEnvelope]
type httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupDeviceTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupDeviceTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPProtocolParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupHTTPProtocolParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupHTTPProtocolParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupHTTPProtocolParams]'s query parameters
// as `url.Values`.
func (r HTTPTimeseriesGroupHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPProtocolParamsAggInterval string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "15m"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1h"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1d  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1d"
	HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1w  HTTPTimeseriesGroupHTTPProtocolParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsAggInterval15m, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1h, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1d, HTTPTimeseriesGroupHTTPProtocolParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsBotClass string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated HTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman     HTTPTimeseriesGroupHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyAutomated, HTTPTimeseriesGroupHTTPProtocolParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyChrome  HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyEdge    HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyFirefox HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilySafari  HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyChrome, HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyEdge, HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilyFirefox, HTTPTimeseriesGroupHTTPProtocolParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsDeviceType string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile  HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther   HTTPTimeseriesGroupHTTPProtocolParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeDesktop, HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeMobile, HTTPTimeseriesGroupHTTPProtocolParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupHTTPProtocolParamsFormat string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsFormatJson HTTPTimeseriesGroupHTTPProtocolParamsFormat = "JSON"
	HTTPTimeseriesGroupHTTPProtocolParamsFormatCsv  HTTPTimeseriesGroupHTTPProtocolParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsFormatJson, HTTPTimeseriesGroupHTTPProtocolParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3 HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv1, HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv2, HTTPTimeseriesGroupHTTPProtocolParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsIPVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4 HTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6 HTTPTimeseriesGroupHTTPProtocolParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv4, HTTPTimeseriesGroupHTTPProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsOS string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsOSWindows  HTTPTimeseriesGroupHTTPProtocolParamsOS = "WINDOWS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx   HTTPTimeseriesGroupHTTPProtocolParamsOS = "MACOSX"
	HTTPTimeseriesGroupHTTPProtocolParamsOSIos      HTTPTimeseriesGroupHTTPProtocolParamsOS = "IOS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSAndroid  HTTPTimeseriesGroupHTTPProtocolParamsOS = "ANDROID"
	HTTPTimeseriesGroupHTTPProtocolParamsOSChromeos HTTPTimeseriesGroupHTTPProtocolParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupHTTPProtocolParamsOSLinux    HTTPTimeseriesGroupHTTPProtocolParamsOS = "LINUX"
	HTTPTimeseriesGroupHTTPProtocolParamsOSSmartTv  HTTPTimeseriesGroupHTTPProtocolParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsOSWindows, HTTPTimeseriesGroupHTTPProtocolParamsOSMacosx, HTTPTimeseriesGroupHTTPProtocolParamsOSIos, HTTPTimeseriesGroupHTTPProtocolParamsOSAndroid, HTTPTimeseriesGroupHTTPProtocolParamsOSChromeos, HTTPTimeseriesGroupHTTPProtocolParamsOSLinux, HTTPTimeseriesGroupHTTPProtocolParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion string

const (
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSvQuic HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupHTTPProtocolParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupHTTPProtocolParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPProtocolResponseEnvelope struct {
	Result  HTTPTimeseriesGroupHTTPProtocolResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPProtocolResponseEnvelope]
type httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupHTTPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupHTTPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupHTTPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupHTTPVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupHTTPVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1h  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1d  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupHTTPVersionParamsAggInterval1w  HTTPTimeseriesGroupHTTPVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsAggInterval15m, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1h, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1d, HTTPTimeseriesGroupHTTPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsBotClass string

const (
	HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupHTTPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily string

const (
	HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyChrome  HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyEdge    HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyFirefox HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilySafari  HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyChrome, HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyEdge, HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilyFirefox, HTTPTimeseriesGroupHTTPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther   HTTPTimeseriesGroupHTTPVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupHTTPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupHTTPVersionParamsFormat string

const (
	HTTPTimeseriesGroupHTTPVersionParamsFormatJson HTTPTimeseriesGroupHTTPVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupHTTPVersionParamsFormatCsv  HTTPTimeseriesGroupHTTPVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsFormatJson, HTTPTimeseriesGroupHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupHTTPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsIPVersion string

const (
	HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4 HTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6 HTTPTimeseriesGroupHTTPVersionParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv4, HTTPTimeseriesGroupHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsOS string

const (
	HTTPTimeseriesGroupHTTPVersionParamsOSWindows  HTTPTimeseriesGroupHTTPVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupHTTPVersionParamsOSMacosx   HTTPTimeseriesGroupHTTPVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupHTTPVersionParamsOSIos      HTTPTimeseriesGroupHTTPVersionParamsOS = "IOS"
	HTTPTimeseriesGroupHTTPVersionParamsOSAndroid  HTTPTimeseriesGroupHTTPVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupHTTPVersionParamsOSChromeos HTTPTimeseriesGroupHTTPVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupHTTPVersionParamsOSLinux    HTTPTimeseriesGroupHTTPVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupHTTPVersionParamsOSSmartTv  HTTPTimeseriesGroupHTTPVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsOSWindows, HTTPTimeseriesGroupHTTPVersionParamsOSMacosx, HTTPTimeseriesGroupHTTPVersionParamsOSIos, HTTPTimeseriesGroupHTTPVersionParamsOSAndroid, HTTPTimeseriesGroupHTTPVersionParamsOSChromeos, HTTPTimeseriesGroupHTTPVersionParamsOSLinux, HTTPTimeseriesGroupHTTPVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionParamsTLSVersion string

const (
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSvQuic HTTPTimeseriesGroupHTTPVersionParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupHTTPVersionParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupHTTPVersionParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupHTTPVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupHTTPVersionResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupHTTPVersionResponseEnvelope]
type httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupIPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupIPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupIPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupIPVersionParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupIPVersionParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupIPVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupIPVersionParamsAggInterval15m HTTPTimeseriesGroupIPVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1h  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1d  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupIPVersionParamsAggInterval1w  HTTPTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsAggInterval15m, HTTPTimeseriesGroupIPVersionParamsAggInterval1h, HTTPTimeseriesGroupIPVersionParamsAggInterval1d, HTTPTimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsBotClass string

const (
	HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupIPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupIPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupIPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsBrowserFamily string

const (
	HTTPTimeseriesGroupIPVersionParamsBrowserFamilyChrome  HTTPTimeseriesGroupIPVersionParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupIPVersionParamsBrowserFamilyEdge    HTTPTimeseriesGroupIPVersionParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupIPVersionParamsBrowserFamilyFirefox HTTPTimeseriesGroupIPVersionParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupIPVersionParamsBrowserFamilySafari  HTTPTimeseriesGroupIPVersionParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupIPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsBrowserFamilyChrome, HTTPTimeseriesGroupIPVersionParamsBrowserFamilyEdge, HTTPTimeseriesGroupIPVersionParamsBrowserFamilyFirefox, HTTPTimeseriesGroupIPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupIPVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupIPVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther   HTTPTimeseriesGroupIPVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupIPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupIPVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupIPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupIPVersionParamsFormat string

const (
	HTTPTimeseriesGroupIPVersionParamsFormatJson HTTPTimeseriesGroupIPVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupIPVersionParamsFormatCsv  HTTPTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsFormatJson, HTTPTimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupIPVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupIPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupIPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsHTTPVersion string

const (
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3 HTTPTimeseriesGroupIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv1, HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv2, HTTPTimeseriesGroupIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsOS string

const (
	HTTPTimeseriesGroupIPVersionParamsOSWindows  HTTPTimeseriesGroupIPVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupIPVersionParamsOSMacosx   HTTPTimeseriesGroupIPVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupIPVersionParamsOSIos      HTTPTimeseriesGroupIPVersionParamsOS = "IOS"
	HTTPTimeseriesGroupIPVersionParamsOSAndroid  HTTPTimeseriesGroupIPVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupIPVersionParamsOSChromeos HTTPTimeseriesGroupIPVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupIPVersionParamsOSLinux    HTTPTimeseriesGroupIPVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupIPVersionParamsOSSmartTv  HTTPTimeseriesGroupIPVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupIPVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsOSWindows, HTTPTimeseriesGroupIPVersionParamsOSMacosx, HTTPTimeseriesGroupIPVersionParamsOSIos, HTTPTimeseriesGroupIPVersionParamsOSAndroid, HTTPTimeseriesGroupIPVersionParamsOSChromeos, HTTPTimeseriesGroupIPVersionParamsOSLinux, HTTPTimeseriesGroupIPVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionParamsTLSVersion string

const (
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSvQuic HTTPTimeseriesGroupIPVersionParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupIPVersionParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupIPVersionParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    httpTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupIPVersionResponseEnvelope]
type httpTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupOSParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupOSParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupOSParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupOSParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupOSParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupOSParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupOSParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupOSParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupOSParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupOSParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupOSParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupOSParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupOSParamsAggInterval string

const (
	HTTPTimeseriesGroupOSParamsAggInterval15m HTTPTimeseriesGroupOSParamsAggInterval = "15m"
	HTTPTimeseriesGroupOSParamsAggInterval1h  HTTPTimeseriesGroupOSParamsAggInterval = "1h"
	HTTPTimeseriesGroupOSParamsAggInterval1d  HTTPTimeseriesGroupOSParamsAggInterval = "1d"
	HTTPTimeseriesGroupOSParamsAggInterval1w  HTTPTimeseriesGroupOSParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupOSParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsAggInterval15m, HTTPTimeseriesGroupOSParamsAggInterval1h, HTTPTimeseriesGroupOSParamsAggInterval1d, HTTPTimeseriesGroupOSParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsBotClass string

const (
	HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated HTTPTimeseriesGroupOSParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupOSParamsBotClassLikelyHuman     HTTPTimeseriesGroupOSParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupOSParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsBotClassLikelyAutomated, HTTPTimeseriesGroupOSParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsBrowserFamily string

const (
	HTTPTimeseriesGroupOSParamsBrowserFamilyChrome  HTTPTimeseriesGroupOSParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupOSParamsBrowserFamilyEdge    HTTPTimeseriesGroupOSParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupOSParamsBrowserFamilyFirefox HTTPTimeseriesGroupOSParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupOSParamsBrowserFamilySafari  HTTPTimeseriesGroupOSParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupOSParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsBrowserFamilyChrome, HTTPTimeseriesGroupOSParamsBrowserFamilyEdge, HTTPTimeseriesGroupOSParamsBrowserFamilyFirefox, HTTPTimeseriesGroupOSParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsDeviceType string

const (
	HTTPTimeseriesGroupOSParamsDeviceTypeDesktop HTTPTimeseriesGroupOSParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupOSParamsDeviceTypeMobile  HTTPTimeseriesGroupOSParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupOSParamsDeviceTypeOther   HTTPTimeseriesGroupOSParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupOSParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsDeviceTypeDesktop, HTTPTimeseriesGroupOSParamsDeviceTypeMobile, HTTPTimeseriesGroupOSParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupOSParamsFormat string

const (
	HTTPTimeseriesGroupOSParamsFormatJson HTTPTimeseriesGroupOSParamsFormat = "JSON"
	HTTPTimeseriesGroupOSParamsFormatCsv  HTTPTimeseriesGroupOSParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupOSParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsFormatJson, HTTPTimeseriesGroupOSParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP  HTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS HTTPTimeseriesGroupOSParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupOSParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsHTTPProtocolHTTP, HTTPTimeseriesGroupOSParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsHTTPVersion string

const (
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3 HTTPTimeseriesGroupOSParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupOSParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsHTTPVersionHttPv1, HTTPTimeseriesGroupOSParamsHTTPVersionHttPv2, HTTPTimeseriesGroupOSParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsIPVersion string

const (
	HTTPTimeseriesGroupOSParamsIPVersionIPv4 HTTPTimeseriesGroupOSParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupOSParamsIPVersionIPv6 HTTPTimeseriesGroupOSParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupOSParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsIPVersionIPv4, HTTPTimeseriesGroupOSParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSParamsTLSVersion string

const (
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupOSParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupOSParamsTLSVersionTlSvQuic HTTPTimeseriesGroupOSParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupOSParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupOSParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupOSParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupOSResponseEnvelope struct {
	Result  HTTPTimeseriesGroupOSResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    httpTimeseriesGroupOSResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupOSResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupOSResponseEnvelope]
type httpTimeseriesGroupOSResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupOSResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupOSResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupPostQuantumParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupPostQuantumParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupPostQuantumParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupPostQuantumParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupPostQuantumParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupPostQuantumParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupPostQuantumParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupPostQuantumParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupPostQuantumParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupPostQuantumParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupPostQuantumParamsAggInterval string

const (
	HTTPTimeseriesGroupPostQuantumParamsAggInterval15m HTTPTimeseriesGroupPostQuantumParamsAggInterval = "15m"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1h  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1h"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1d  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1d"
	HTTPTimeseriesGroupPostQuantumParamsAggInterval1w  HTTPTimeseriesGroupPostQuantumParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupPostQuantumParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsAggInterval15m, HTTPTimeseriesGroupPostQuantumParamsAggInterval1h, HTTPTimeseriesGroupPostQuantumParamsAggInterval1d, HTTPTimeseriesGroupPostQuantumParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsBotClass string

const (
	HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated HTTPTimeseriesGroupPostQuantumParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman     HTTPTimeseriesGroupPostQuantumParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupPostQuantumParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyAutomated, HTTPTimeseriesGroupPostQuantumParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsBrowserFamily string

const (
	HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyChrome  HTTPTimeseriesGroupPostQuantumParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyEdge    HTTPTimeseriesGroupPostQuantumParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyFirefox HTTPTimeseriesGroupPostQuantumParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupPostQuantumParamsBrowserFamilySafari  HTTPTimeseriesGroupPostQuantumParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupPostQuantumParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyChrome, HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyEdge, HTTPTimeseriesGroupPostQuantumParamsBrowserFamilyFirefox, HTTPTimeseriesGroupPostQuantumParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsDeviceType string

const (
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop HTTPTimeseriesGroupPostQuantumParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile  HTTPTimeseriesGroupPostQuantumParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther   HTTPTimeseriesGroupPostQuantumParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupPostQuantumParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsDeviceTypeDesktop, HTTPTimeseriesGroupPostQuantumParamsDeviceTypeMobile, HTTPTimeseriesGroupPostQuantumParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupPostQuantumParamsFormat string

const (
	HTTPTimeseriesGroupPostQuantumParamsFormatJson HTTPTimeseriesGroupPostQuantumParamsFormat = "JSON"
	HTTPTimeseriesGroupPostQuantumParamsFormatCsv  HTTPTimeseriesGroupPostQuantumParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupPostQuantumParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsFormatJson, HTTPTimeseriesGroupPostQuantumParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP  HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupPostQuantumParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTP, HTTPTimeseriesGroupPostQuantumParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsHTTPVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3 HTTPTimeseriesGroupPostQuantumParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupPostQuantumParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv1, HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv2, HTTPTimeseriesGroupPostQuantumParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsIPVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4 HTTPTimeseriesGroupPostQuantumParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6 HTTPTimeseriesGroupPostQuantumParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupPostQuantumParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv4, HTTPTimeseriesGroupPostQuantumParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsOS string

const (
	HTTPTimeseriesGroupPostQuantumParamsOSWindows  HTTPTimeseriesGroupPostQuantumParamsOS = "WINDOWS"
	HTTPTimeseriesGroupPostQuantumParamsOSMacosx   HTTPTimeseriesGroupPostQuantumParamsOS = "MACOSX"
	HTTPTimeseriesGroupPostQuantumParamsOSIos      HTTPTimeseriesGroupPostQuantumParamsOS = "IOS"
	HTTPTimeseriesGroupPostQuantumParamsOSAndroid  HTTPTimeseriesGroupPostQuantumParamsOS = "ANDROID"
	HTTPTimeseriesGroupPostQuantumParamsOSChromeos HTTPTimeseriesGroupPostQuantumParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupPostQuantumParamsOSLinux    HTTPTimeseriesGroupPostQuantumParamsOS = "LINUX"
	HTTPTimeseriesGroupPostQuantumParamsOSSmartTv  HTTPTimeseriesGroupPostQuantumParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupPostQuantumParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsOSWindows, HTTPTimeseriesGroupPostQuantumParamsOSMacosx, HTTPTimeseriesGroupPostQuantumParamsOSIos, HTTPTimeseriesGroupPostQuantumParamsOSAndroid, HTTPTimeseriesGroupPostQuantumParamsOSChromeos, HTTPTimeseriesGroupPostQuantumParamsOSLinux, HTTPTimeseriesGroupPostQuantumParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumParamsTLSVersion string

const (
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSvQuic HTTPTimeseriesGroupPostQuantumParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupPostQuantumParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupPostQuantumParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupPostQuantumResponseEnvelope struct {
	Result  HTTPTimeseriesGroupPostQuantumResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    httpTimeseriesGroupPostQuantumResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupPostQuantumResponseEnvelopeJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupPostQuantumResponseEnvelope]
type httpTimeseriesGroupPostQuantumResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupPostQuantumResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupPostQuantumResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupTLSVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesGroupTLSVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupTLSVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupTLSVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupTLSVersionParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupTLSVersionParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPTimeseriesGroupTLSVersionParams]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupTLSVersionParamsAggInterval string

const (
	HTTPTimeseriesGroupTLSVersionParamsAggInterval15m HTTPTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1h  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1d  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	HTTPTimeseriesGroupTLSVersionParamsAggInterval1w  HTTPTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupTLSVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsAggInterval15m, HTTPTimeseriesGroupTLSVersionParamsAggInterval1h, HTTPTimeseriesGroupTLSVersionParamsAggInterval1d, HTTPTimeseriesGroupTLSVersionParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsBotClass string

const (
	HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated HTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman     HTTPTimeseriesGroupTLSVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupTLSVersionParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyAutomated, HTTPTimeseriesGroupTLSVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsBrowserFamily string

const (
	HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyChrome  HTTPTimeseriesGroupTLSVersionParamsBrowserFamily = "CHROME"
	HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyEdge    HTTPTimeseriesGroupTLSVersionParamsBrowserFamily = "EDGE"
	HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyFirefox HTTPTimeseriesGroupTLSVersionParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesGroupTLSVersionParamsBrowserFamilySafari  HTTPTimeseriesGroupTLSVersionParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesGroupTLSVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyChrome, HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyEdge, HTTPTimeseriesGroupTLSVersionParamsBrowserFamilyFirefox, HTTPTimeseriesGroupTLSVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsDeviceType string

const (
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop HTTPTimeseriesGroupTLSVersionParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile  HTTPTimeseriesGroupTLSVersionParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther   HTTPTimeseriesGroupTLSVersionParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupTLSVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsDeviceTypeDesktop, HTTPTimeseriesGroupTLSVersionParamsDeviceTypeMobile, HTTPTimeseriesGroupTLSVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupTLSVersionParamsFormat string

const (
	HTTPTimeseriesGroupTLSVersionParamsFormatJson HTTPTimeseriesGroupTLSVersionParamsFormat = "JSON"
	HTTPTimeseriesGroupTLSVersionParamsFormatCsv  HTTPTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupTLSVersionParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsFormatJson, HTTPTimeseriesGroupTLSVersionParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP  HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupTLSVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTP, HTTPTimeseriesGroupTLSVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsHTTPVersion string

const (
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3 HTTPTimeseriesGroupTLSVersionParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupTLSVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv1, HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv2, HTTPTimeseriesGroupTLSVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsIPVersion string

const (
	HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4 HTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6 HTTPTimeseriesGroupTLSVersionParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupTLSVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv4, HTTPTimeseriesGroupTLSVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionParamsOS string

const (
	HTTPTimeseriesGroupTLSVersionParamsOSWindows  HTTPTimeseriesGroupTLSVersionParamsOS = "WINDOWS"
	HTTPTimeseriesGroupTLSVersionParamsOSMacosx   HTTPTimeseriesGroupTLSVersionParamsOS = "MACOSX"
	HTTPTimeseriesGroupTLSVersionParamsOSIos      HTTPTimeseriesGroupTLSVersionParamsOS = "IOS"
	HTTPTimeseriesGroupTLSVersionParamsOSAndroid  HTTPTimeseriesGroupTLSVersionParamsOS = "ANDROID"
	HTTPTimeseriesGroupTLSVersionParamsOSChromeos HTTPTimeseriesGroupTLSVersionParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupTLSVersionParamsOSLinux    HTTPTimeseriesGroupTLSVersionParamsOS = "LINUX"
	HTTPTimeseriesGroupTLSVersionParamsOSSmartTv  HTTPTimeseriesGroupTLSVersionParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupTLSVersionParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupTLSVersionParamsOSWindows, HTTPTimeseriesGroupTLSVersionParamsOSMacosx, HTTPTimeseriesGroupTLSVersionParamsOSIos, HTTPTimeseriesGroupTLSVersionParamsOSAndroid, HTTPTimeseriesGroupTLSVersionParamsOSChromeos, HTTPTimeseriesGroupTLSVersionParamsOSLinux, HTTPTimeseriesGroupTLSVersionParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  HTTPTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    httpTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupTLSVersionResponseEnvelope]
type httpTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
