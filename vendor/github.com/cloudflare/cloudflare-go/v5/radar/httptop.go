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

// HTTPTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPTopService] method instead.
type HTTPTopService struct {
	Options []option.RequestOption
}

// NewHTTPTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPTopService(opts ...option.RequestOption) (r *HTTPTopService) {
	r = &HTTPTopService{}
	r.Options = opts
	return
}

// Retrieves the top user agents by HTTP requests.
func (r *HTTPTopService) Browser(ctx context.Context, query HTTPTopBrowserParams, opts ...option.RequestOption) (res *HTTPTopBrowserResponse, err error) {
	var env HTTPTopBrowserResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top user agents, aggregated in families, by HTTP requests.
func (r *HTTPTopService) BrowserFamily(ctx context.Context, query HTTPTopBrowserFamilyParams, opts ...option.RequestOption) (res *HTTPTopBrowserFamilyResponse, err error) {
	var env HTTPTopBrowserFamilyResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTopBrowserResponse struct {
	// Metadata for the results.
	Meta HTTPTopBrowserResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowserResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowserResponseJSON   `json:"-"`
}

// httpTopBrowserResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponse]
type httpTopBrowserResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTopBrowserResponseMeta struct {
	ConfidenceInfo HTTPTopBrowserResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPTopBrowserResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTopBrowserResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTopBrowserResponseMetaUnit `json:"units,required"`
	JSON  httpTopBrowserResponseMetaJSON   `json:"-"`
}

// httpTopBrowserResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseMeta]
type httpTopBrowserResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowserResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  httpTopBrowserResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTopBrowserResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPTopBrowserResponseMetaConfidenceInfo]
type httpTopBrowserResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTopBrowserResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPTopBrowserResponseMetaConfidenceInfoAnnotation]
type httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowserResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowserResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowserResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserResponseMetaDateRange]
type httpTopBrowserResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTopBrowserResponseMetaNormalization string

const (
	HTTPTopBrowserResponseMetaNormalizationPercentage           HTTPTopBrowserResponseMetaNormalization = "PERCENTAGE"
	HTTPTopBrowserResponseMetaNormalizationMin0Max              HTTPTopBrowserResponseMetaNormalization = "MIN0_MAX"
	HTTPTopBrowserResponseMetaNormalizationMinMax               HTTPTopBrowserResponseMetaNormalization = "MIN_MAX"
	HTTPTopBrowserResponseMetaNormalizationRawValues            HTTPTopBrowserResponseMetaNormalization = "RAW_VALUES"
	HTTPTopBrowserResponseMetaNormalizationPercentageChange     HTTPTopBrowserResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTopBrowserResponseMetaNormalizationRollingAverage       HTTPTopBrowserResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTopBrowserResponseMetaNormalizationOverlappedPercentage HTTPTopBrowserResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTopBrowserResponseMetaNormalizationRatio                HTTPTopBrowserResponseMetaNormalization = "RATIO"
)

func (r HTTPTopBrowserResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTopBrowserResponseMetaNormalizationPercentage, HTTPTopBrowserResponseMetaNormalizationMin0Max, HTTPTopBrowserResponseMetaNormalizationMinMax, HTTPTopBrowserResponseMetaNormalizationRawValues, HTTPTopBrowserResponseMetaNormalizationPercentageChange, HTTPTopBrowserResponseMetaNormalizationRollingAverage, HTTPTopBrowserResponseMetaNormalizationOverlappedPercentage, HTTPTopBrowserResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTopBrowserResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  httpTopBrowserResponseMetaUnitJSON `json:"-"`
}

// httpTopBrowserResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseMetaUnit]
type httpTopBrowserResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserResponseTop0 struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  httpTopBrowserResponseTop0JSON `json:"-"`
}

// httpTopBrowserResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseTop0]
type httpTopBrowserResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponse struct {
	// Metadata for the results.
	Meta HTTPTopBrowserFamilyResponseMeta   `json:"meta,required"`
	Top0 []HTTPTopBrowserFamilyResponseTop0 `json:"top_0,required"`
	JSON httpTopBrowserFamilyResponseJSON   `json:"-"`
}

// httpTopBrowserFamilyResponseJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponse]
type httpTopBrowserFamilyResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTopBrowserFamilyResponseMeta struct {
	ConfidenceInfo HTTPTopBrowserFamilyResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPTopBrowserFamilyResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTopBrowserFamilyResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTopBrowserFamilyResponseMetaUnit `json:"units,required"`
	JSON  httpTopBrowserFamilyResponseMetaJSON   `json:"-"`
}

// httpTopBrowserFamilyResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponseMeta]
type httpTopBrowserFamilyResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMetaConfidenceInfo struct {
	Annotations []HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  httpTopBrowserFamilyResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPTopBrowserFamilyResponseMetaConfidenceInfo]
type httpTopBrowserFamilyResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation]
type httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTopBrowserFamilyResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpTopBrowserFamilyResponseMetaDateRangeJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamilyResponseMetaDateRange]
type httpTopBrowserFamilyResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTopBrowserFamilyResponseMetaNormalization string

const (
	HTTPTopBrowserFamilyResponseMetaNormalizationPercentage           HTTPTopBrowserFamilyResponseMetaNormalization = "PERCENTAGE"
	HTTPTopBrowserFamilyResponseMetaNormalizationMin0Max              HTTPTopBrowserFamilyResponseMetaNormalization = "MIN0_MAX"
	HTTPTopBrowserFamilyResponseMetaNormalizationMinMax               HTTPTopBrowserFamilyResponseMetaNormalization = "MIN_MAX"
	HTTPTopBrowserFamilyResponseMetaNormalizationRawValues            HTTPTopBrowserFamilyResponseMetaNormalization = "RAW_VALUES"
	HTTPTopBrowserFamilyResponseMetaNormalizationPercentageChange     HTTPTopBrowserFamilyResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTopBrowserFamilyResponseMetaNormalizationRollingAverage       HTTPTopBrowserFamilyResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTopBrowserFamilyResponseMetaNormalizationOverlappedPercentage HTTPTopBrowserFamilyResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTopBrowserFamilyResponseMetaNormalizationRatio                HTTPTopBrowserFamilyResponseMetaNormalization = "RATIO"
)

func (r HTTPTopBrowserFamilyResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyResponseMetaNormalizationPercentage, HTTPTopBrowserFamilyResponseMetaNormalizationMin0Max, HTTPTopBrowserFamilyResponseMetaNormalizationMinMax, HTTPTopBrowserFamilyResponseMetaNormalizationRawValues, HTTPTopBrowserFamilyResponseMetaNormalizationPercentageChange, HTTPTopBrowserFamilyResponseMetaNormalizationRollingAverage, HTTPTopBrowserFamilyResponseMetaNormalizationOverlappedPercentage, HTTPTopBrowserFamilyResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  httpTopBrowserFamilyResponseMetaUnitJSON `json:"-"`
}

// httpTopBrowserFamilyResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamilyResponseMetaUnit]
type httpTopBrowserFamilyResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyResponseTop0 struct {
	Name  string                               `json:"name,required"`
	Value string                               `json:"value,required"`
	JSON  httpTopBrowserFamilyResponseTop0JSON `json:"-"`
}

// httpTopBrowserFamilyResponseTop0JSON contains the JSON metadata for the struct
// [HTTPTopBrowserFamilyResponseTop0]
type httpTopBrowserFamilyResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowserParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTopBrowserParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPTopBrowserParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTopBrowserParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTopBrowserParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTopBrowserParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTopBrowserParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTopBrowserParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTopBrowserParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowserParams]'s query parameters as `url.Values`.
func (r HTTPTopBrowserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPTopBrowserParamsBotClass string

const (
	HTTPTopBrowserParamsBotClassLikelyAutomated HTTPTopBrowserParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowserParamsBotClassLikelyHuman     HTTPTopBrowserParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowserParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsBotClassLikelyAutomated, HTTPTopBrowserParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowserParamsBrowserFamily string

const (
	HTTPTopBrowserParamsBrowserFamilyChrome  HTTPTopBrowserParamsBrowserFamily = "CHROME"
	HTTPTopBrowserParamsBrowserFamilyEdge    HTTPTopBrowserParamsBrowserFamily = "EDGE"
	HTTPTopBrowserParamsBrowserFamilyFirefox HTTPTopBrowserParamsBrowserFamily = "FIREFOX"
	HTTPTopBrowserParamsBrowserFamilySafari  HTTPTopBrowserParamsBrowserFamily = "SAFARI"
)

func (r HTTPTopBrowserParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsBrowserFamilyChrome, HTTPTopBrowserParamsBrowserFamilyEdge, HTTPTopBrowserParamsBrowserFamilyFirefox, HTTPTopBrowserParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTopBrowserParamsDeviceType string

const (
	HTTPTopBrowserParamsDeviceTypeDesktop HTTPTopBrowserParamsDeviceType = "DESKTOP"
	HTTPTopBrowserParamsDeviceTypeMobile  HTTPTopBrowserParamsDeviceType = "MOBILE"
	HTTPTopBrowserParamsDeviceTypeOther   HTTPTopBrowserParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowserParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsDeviceTypeDesktop, HTTPTopBrowserParamsDeviceTypeMobile, HTTPTopBrowserParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTopBrowserParamsFormat string

const (
	HTTPTopBrowserParamsFormatJson HTTPTopBrowserParamsFormat = "JSON"
	HTTPTopBrowserParamsFormatCsv  HTTPTopBrowserParamsFormat = "CSV"
)

func (r HTTPTopBrowserParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsFormatJson, HTTPTopBrowserParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowserParamsHTTPProtocol string

const (
	HTTPTopBrowserParamsHTTPProtocolHTTP  HTTPTopBrowserParamsHTTPProtocol = "HTTP"
	HTTPTopBrowserParamsHTTPProtocolHTTPS HTTPTopBrowserParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowserParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsHTTPProtocolHTTP, HTTPTopBrowserParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowserParamsHTTPVersion string

const (
	HTTPTopBrowserParamsHTTPVersionHttPv1 HTTPTopBrowserParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowserParamsHTTPVersionHttPv2 HTTPTopBrowserParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowserParamsHTTPVersionHttPv3 HTTPTopBrowserParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowserParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsHTTPVersionHttPv1, HTTPTopBrowserParamsHTTPVersionHttPv2, HTTPTopBrowserParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowserParamsIPVersion string

const (
	HTTPTopBrowserParamsIPVersionIPv4 HTTPTopBrowserParamsIPVersion = "IPv4"
	HTTPTopBrowserParamsIPVersionIPv6 HTTPTopBrowserParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowserParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsIPVersionIPv4, HTTPTopBrowserParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowserParamsOS string

const (
	HTTPTopBrowserParamsOSWindows  HTTPTopBrowserParamsOS = "WINDOWS"
	HTTPTopBrowserParamsOSMacosx   HTTPTopBrowserParamsOS = "MACOSX"
	HTTPTopBrowserParamsOSIos      HTTPTopBrowserParamsOS = "IOS"
	HTTPTopBrowserParamsOSAndroid  HTTPTopBrowserParamsOS = "ANDROID"
	HTTPTopBrowserParamsOSChromeos HTTPTopBrowserParamsOS = "CHROMEOS"
	HTTPTopBrowserParamsOSLinux    HTTPTopBrowserParamsOS = "LINUX"
	HTTPTopBrowserParamsOSSmartTv  HTTPTopBrowserParamsOS = "SMART_TV"
)

func (r HTTPTopBrowserParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsOSWindows, HTTPTopBrowserParamsOSMacosx, HTTPTopBrowserParamsOSIos, HTTPTopBrowserParamsOSAndroid, HTTPTopBrowserParamsOSChromeos, HTTPTopBrowserParamsOSLinux, HTTPTopBrowserParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowserParamsTLSVersion string

const (
	HTTPTopBrowserParamsTLSVersionTlSv1_0  HTTPTopBrowserParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowserParamsTLSVersionTlSv1_1  HTTPTopBrowserParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowserParamsTLSVersionTlSv1_2  HTTPTopBrowserParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowserParamsTLSVersionTlSv1_3  HTTPTopBrowserParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowserParamsTLSVersionTlSvQuic HTTPTopBrowserParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowserParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserParamsTLSVersionTlSv1_0, HTTPTopBrowserParamsTLSVersionTlSv1_1, HTTPTopBrowserParamsTLSVersionTlSv1_2, HTTPTopBrowserParamsTLSVersionTlSv1_3, HTTPTopBrowserParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowserResponseEnvelope struct {
	Result  HTTPTopBrowserResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    httpTopBrowserResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowserResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPTopBrowserResponseEnvelope]
type httpTopBrowserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTopBrowserFamilyParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTopBrowserFamilyParamsBotClass] `query:"botClass"`
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
	DeviceType param.Field[[]HTTPTopBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTopBrowserFamilyParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTopBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTopBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTopBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTopBrowserFamilyParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTopBrowserFamilyParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTopBrowserFamilyParams]'s query parameters as
// `url.Values`.
func (r HTTPTopBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPTopBrowserFamilyParamsBotClass string

const (
	HTTPTopBrowserFamilyParamsBotClassLikelyAutomated HTTPTopBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTopBrowserFamilyParamsBotClassLikelyHuman     HTTPTopBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTopBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsBotClassLikelyAutomated, HTTPTopBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsDeviceType string

const (
	HTTPTopBrowserFamilyParamsDeviceTypeDesktop HTTPTopBrowserFamilyParamsDeviceType = "DESKTOP"
	HTTPTopBrowserFamilyParamsDeviceTypeMobile  HTTPTopBrowserFamilyParamsDeviceType = "MOBILE"
	HTTPTopBrowserFamilyParamsDeviceTypeOther   HTTPTopBrowserFamilyParamsDeviceType = "OTHER"
)

func (r HTTPTopBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsDeviceTypeDesktop, HTTPTopBrowserFamilyParamsDeviceTypeMobile, HTTPTopBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTopBrowserFamilyParamsFormat string

const (
	HTTPTopBrowserFamilyParamsFormatJson HTTPTopBrowserFamilyParamsFormat = "JSON"
	HTTPTopBrowserFamilyParamsFormatCsv  HTTPTopBrowserFamilyParamsFormat = "CSV"
)

func (r HTTPTopBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsFormatJson, HTTPTopBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsHTTPProtocol string

const (
	HTTPTopBrowserFamilyParamsHTTPProtocolHTTP  HTTPTopBrowserFamilyParamsHTTPProtocol = "HTTP"
	HTTPTopBrowserFamilyParamsHTTPProtocolHTTPS HTTPTopBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTopBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsHTTPProtocolHTTP, HTTPTopBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsHTTPVersion string

const (
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv1 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv1"
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv2 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv2"
	HTTPTopBrowserFamilyParamsHTTPVersionHttPv3 HTTPTopBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTopBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsHTTPVersionHttPv1, HTTPTopBrowserFamilyParamsHTTPVersionHttPv2, HTTPTopBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsIPVersion string

const (
	HTTPTopBrowserFamilyParamsIPVersionIPv4 HTTPTopBrowserFamilyParamsIPVersion = "IPv4"
	HTTPTopBrowserFamilyParamsIPVersionIPv6 HTTPTopBrowserFamilyParamsIPVersion = "IPv6"
)

func (r HTTPTopBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsIPVersionIPv4, HTTPTopBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsOS string

const (
	HTTPTopBrowserFamilyParamsOSWindows  HTTPTopBrowserFamilyParamsOS = "WINDOWS"
	HTTPTopBrowserFamilyParamsOSMacosx   HTTPTopBrowserFamilyParamsOS = "MACOSX"
	HTTPTopBrowserFamilyParamsOSIos      HTTPTopBrowserFamilyParamsOS = "IOS"
	HTTPTopBrowserFamilyParamsOSAndroid  HTTPTopBrowserFamilyParamsOS = "ANDROID"
	HTTPTopBrowserFamilyParamsOSChromeos HTTPTopBrowserFamilyParamsOS = "CHROMEOS"
	HTTPTopBrowserFamilyParamsOSLinux    HTTPTopBrowserFamilyParamsOS = "LINUX"
	HTTPTopBrowserFamilyParamsOSSmartTv  HTTPTopBrowserFamilyParamsOS = "SMART_TV"
)

func (r HTTPTopBrowserFamilyParamsOS) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsOSWindows, HTTPTopBrowserFamilyParamsOSMacosx, HTTPTopBrowserFamilyParamsOSIos, HTTPTopBrowserFamilyParamsOSAndroid, HTTPTopBrowserFamilyParamsOSChromeos, HTTPTopBrowserFamilyParamsOSLinux, HTTPTopBrowserFamilyParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyParamsTLSVersion string

const (
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_0"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_1  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_1"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_2  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_2"
	HTTPTopBrowserFamilyParamsTLSVersionTlSv1_3  HTTPTopBrowserFamilyParamsTLSVersion = "TLSv1_3"
	HTTPTopBrowserFamilyParamsTLSVersionTlSvQuic HTTPTopBrowserFamilyParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTopBrowserFamilyParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTopBrowserFamilyParamsTLSVersionTlSv1_0, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_1, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_2, HTTPTopBrowserFamilyParamsTLSVersionTlSv1_3, HTTPTopBrowserFamilyParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTopBrowserFamilyResponseEnvelope struct {
	Result  HTTPTopBrowserFamilyResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpTopBrowserFamilyResponseEnvelopeJSON `json:"-"`
}

// httpTopBrowserFamilyResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTopBrowserFamilyResponseEnvelope]
type httpTopBrowserFamilyResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTopBrowserFamilyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTopBrowserFamilyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
