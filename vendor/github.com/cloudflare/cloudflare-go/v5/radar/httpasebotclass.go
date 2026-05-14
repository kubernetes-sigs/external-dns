// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// HTTPAseBotClassService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseBotClassService] method instead.
type HTTPAseBotClassService struct {
	Options []option.RequestOption
}

// NewHTTPAseBotClassService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPAseBotClassService(opts ...option.RequestOption) (r *HTTPAseBotClassService) {
	r = &HTTPAseBotClassService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested bot
// class.
func (r *HTTPAseBotClassService) Get(ctx context.Context, botClass HTTPAseBotClassGetParamsBotClass, query HTTPAseBotClassGetParams, opts ...option.RequestOption) (res *HTTPAseBotClassGetResponse, err error) {
	var env HTTPAseBotClassGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseBotClassGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseBotClassGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseBotClassGetResponseTop0 `json:"top_0,required"`
	JSON httpAseBotClassGetResponseJSON   `json:"-"`
}

// httpAseBotClassGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponse]
type httpAseBotClassGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseBotClassGetResponseMeta struct {
	ConfidenceInfo HTTPAseBotClassGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseBotClassGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseBotClassGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseBotClassGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseBotClassGetResponseMetaJSON   `json:"-"`
}

// httpAseBotClassGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseMeta]
type httpAseBotClassGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  httpAseBotClassGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [HTTPAseBotClassGetResponseMetaConfidenceInfo]
type httpAseBotClassGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation]
type httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      httpAseBotClassGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseBotClassGetResponseMetaDateRange]
type httpAseBotClassGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseBotClassGetResponseMetaNormalization string

const (
	HTTPAseBotClassGetResponseMetaNormalizationPercentage           HTTPAseBotClassGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseBotClassGetResponseMetaNormalizationMin0Max              HTTPAseBotClassGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseBotClassGetResponseMetaNormalizationMinMax               HTTPAseBotClassGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseBotClassGetResponseMetaNormalizationRawValues            HTTPAseBotClassGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseBotClassGetResponseMetaNormalizationPercentageChange     HTTPAseBotClassGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseBotClassGetResponseMetaNormalizationRollingAverage       HTTPAseBotClassGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseBotClassGetResponseMetaNormalizationOverlappedPercentage HTTPAseBotClassGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseBotClassGetResponseMetaNormalizationRatio                HTTPAseBotClassGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseBotClassGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetResponseMetaNormalizationPercentage, HTTPAseBotClassGetResponseMetaNormalizationMin0Max, HTTPAseBotClassGetResponseMetaNormalizationMinMax, HTTPAseBotClassGetResponseMetaNormalizationRawValues, HTTPAseBotClassGetResponseMetaNormalizationPercentageChange, HTTPAseBotClassGetResponseMetaNormalizationRollingAverage, HTTPAseBotClassGetResponseMetaNormalizationOverlappedPercentage, HTTPAseBotClassGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseBotClassGetResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  httpAseBotClassGetResponseMetaUnitJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseMetaUnit]
type httpAseBotClassGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                             `json:"value,required"`
	JSON  httpAseBotClassGetResponseTop0JSON `json:"-"`
}

// httpAseBotClassGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseTop0]
type httpAseBotClassGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseBotClassGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseBotClassGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseBotClassGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseBotClassGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseBotClassGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Bot class. Refer to
// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
type HTTPAseBotClassGetParamsBotClass string

const (
	HTTPAseBotClassGetParamsBotClassLikelyAutomated HTTPAseBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseBotClassGetParamsBotClassLikelyHuman     HTTPAseBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseBotClassGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsBotClassLikelyAutomated, HTTPAseBotClassGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsBrowserFamily string

const (
	HTTPAseBotClassGetParamsBrowserFamilyChrome  HTTPAseBotClassGetParamsBrowserFamily = "CHROME"
	HTTPAseBotClassGetParamsBrowserFamilyEdge    HTTPAseBotClassGetParamsBrowserFamily = "EDGE"
	HTTPAseBotClassGetParamsBrowserFamilyFirefox HTTPAseBotClassGetParamsBrowserFamily = "FIREFOX"
	HTTPAseBotClassGetParamsBrowserFamilySafari  HTTPAseBotClassGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseBotClassGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsBrowserFamilyChrome, HTTPAseBotClassGetParamsBrowserFamilyEdge, HTTPAseBotClassGetParamsBrowserFamilyFirefox, HTTPAseBotClassGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsDeviceType string

const (
	HTTPAseBotClassGetParamsDeviceTypeDesktop HTTPAseBotClassGetParamsDeviceType = "DESKTOP"
	HTTPAseBotClassGetParamsDeviceTypeMobile  HTTPAseBotClassGetParamsDeviceType = "MOBILE"
	HTTPAseBotClassGetParamsDeviceTypeOther   HTTPAseBotClassGetParamsDeviceType = "OTHER"
)

func (r HTTPAseBotClassGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsDeviceTypeDesktop, HTTPAseBotClassGetParamsDeviceTypeMobile, HTTPAseBotClassGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseBotClassGetParamsFormat string

const (
	HTTPAseBotClassGetParamsFormatJson HTTPAseBotClassGetParamsFormat = "JSON"
	HTTPAseBotClassGetParamsFormatCsv  HTTPAseBotClassGetParamsFormat = "CSV"
)

func (r HTTPAseBotClassGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsFormatJson, HTTPAseBotClassGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsHTTPProtocol string

const (
	HTTPAseBotClassGetParamsHTTPProtocolHTTP  HTTPAseBotClassGetParamsHTTPProtocol = "HTTP"
	HTTPAseBotClassGetParamsHTTPProtocolHTTPS HTTPAseBotClassGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseBotClassGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsHTTPProtocolHTTP, HTTPAseBotClassGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsHTTPVersion string

const (
	HTTPAseBotClassGetParamsHTTPVersionHttPv1 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv1"
	HTTPAseBotClassGetParamsHTTPVersionHttPv2 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv2"
	HTTPAseBotClassGetParamsHTTPVersionHttPv3 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseBotClassGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsHTTPVersionHttPv1, HTTPAseBotClassGetParamsHTTPVersionHttPv2, HTTPAseBotClassGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsIPVersion string

const (
	HTTPAseBotClassGetParamsIPVersionIPv4 HTTPAseBotClassGetParamsIPVersion = "IPv4"
	HTTPAseBotClassGetParamsIPVersionIPv6 HTTPAseBotClassGetParamsIPVersion = "IPv6"
)

func (r HTTPAseBotClassGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsIPVersionIPv4, HTTPAseBotClassGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsOS string

const (
	HTTPAseBotClassGetParamsOSWindows  HTTPAseBotClassGetParamsOS = "WINDOWS"
	HTTPAseBotClassGetParamsOSMacosx   HTTPAseBotClassGetParamsOS = "MACOSX"
	HTTPAseBotClassGetParamsOSIos      HTTPAseBotClassGetParamsOS = "IOS"
	HTTPAseBotClassGetParamsOSAndroid  HTTPAseBotClassGetParamsOS = "ANDROID"
	HTTPAseBotClassGetParamsOSChromeos HTTPAseBotClassGetParamsOS = "CHROMEOS"
	HTTPAseBotClassGetParamsOSLinux    HTTPAseBotClassGetParamsOS = "LINUX"
	HTTPAseBotClassGetParamsOSSmartTv  HTTPAseBotClassGetParamsOS = "SMART_TV"
)

func (r HTTPAseBotClassGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsOSWindows, HTTPAseBotClassGetParamsOSMacosx, HTTPAseBotClassGetParamsOSIos, HTTPAseBotClassGetParamsOSAndroid, HTTPAseBotClassGetParamsOSChromeos, HTTPAseBotClassGetParamsOSLinux, HTTPAseBotClassGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsTLSVersion string

const (
	HTTPAseBotClassGetParamsTLSVersionTlSv1_0  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_0"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_1  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_1"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_2  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_2"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_3  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_3"
	HTTPAseBotClassGetParamsTLSVersionTlSvQuic HTTPAseBotClassGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseBotClassGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsTLSVersionTlSv1_0, HTTPAseBotClassGetParamsTLSVersionTlSv1_1, HTTPAseBotClassGetParamsTLSVersionTlSv1_2, HTTPAseBotClassGetParamsTLSVersionTlSv1_3, HTTPAseBotClassGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseBotClassGetResponseEnvelope struct {
	Result  HTTPAseBotClassGetResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    httpAseBotClassGetResponseEnvelopeJSON `json:"-"`
}

// httpAseBotClassGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseEnvelope]
type httpAseBotClassGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
