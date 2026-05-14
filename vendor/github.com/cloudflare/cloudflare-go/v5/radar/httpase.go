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

// HTTPAseService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseService] method instead.
type HTTPAseService struct {
	Options       []option.RequestOption
	BotClass      *HTTPAseBotClassService
	DeviceType    *HTTPAseDeviceTypeService
	HTTPProtocol  *HTTPAseHTTPProtocolService
	HTTPMethod    *HTTPAseHTTPMethodService
	IPVersion     *HTTPAseIPVersionService
	OS            *HTTPAseOSService
	TLSVersion    *HTTPAseTLSVersionService
	BrowserFamily *HTTPAseBrowserFamilyService
}

// NewHTTPAseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPAseService(opts ...option.RequestOption) (r *HTTPAseService) {
	r = &HTTPAseService{}
	r.Options = opts
	r.BotClass = NewHTTPAseBotClassService(opts...)
	r.DeviceType = NewHTTPAseDeviceTypeService(opts...)
	r.HTTPProtocol = NewHTTPAseHTTPProtocolService(opts...)
	r.HTTPMethod = NewHTTPAseHTTPMethodService(opts...)
	r.IPVersion = NewHTTPAseIPVersionService(opts...)
	r.OS = NewHTTPAseOSService(opts...)
	r.TLSVersion = NewHTTPAseTLSVersionService(opts...)
	r.BrowserFamily = NewHTTPAseBrowserFamilyService(opts...)
	return
}

// Retrieves the top autonomous systems by HTTP requests.
func (r *HTTPAseService) Get(ctx context.Context, query HTTPAseGetParams, opts ...option.RequestOption) (res *HTTPAseGetResponse, err error) {
	var env HTTPAseGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseGetResponseTop0 `json:"top_0,required"`
	JSON httpAseGetResponseJSON   `json:"-"`
}

// httpAseGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseGetResponse]
type httpAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseGetResponseMeta struct {
	ConfidenceInfo HTTPAseGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseGetResponseMetaJSON   `json:"-"`
}

// httpAseGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseMeta]
type httpAseGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                    `json:"level,required"`
	JSON  httpAseGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPAseGetResponseMetaConfidenceInfo]
type httpAseGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                          `json:"startDate,required" format:"date-time"`
	JSON            httpAseGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [HTTPAseGetResponseMetaConfidenceInfoAnnotation]
type httpAseGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                           `json:"startTime,required" format:"date-time"`
	JSON      httpAseGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseGetResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseMetaDateRange]
type httpAseGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseGetResponseMetaNormalization string

const (
	HTTPAseGetResponseMetaNormalizationPercentage           HTTPAseGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseGetResponseMetaNormalizationMin0Max              HTTPAseGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseGetResponseMetaNormalizationMinMax               HTTPAseGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseGetResponseMetaNormalizationRawValues            HTTPAseGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseGetResponseMetaNormalizationPercentageChange     HTTPAseGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseGetResponseMetaNormalizationRollingAverage       HTTPAseGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseGetResponseMetaNormalizationOverlappedPercentage HTTPAseGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseGetResponseMetaNormalizationRatio                HTTPAseGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseGetResponseMetaNormalizationPercentage, HTTPAseGetResponseMetaNormalizationMin0Max, HTTPAseGetResponseMetaNormalizationMinMax, HTTPAseGetResponseMetaNormalizationRawValues, HTTPAseGetResponseMetaNormalizationPercentageChange, HTTPAseGetResponseMetaNormalizationRollingAverage, HTTPAseGetResponseMetaNormalizationOverlappedPercentage, HTTPAseGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseGetResponseMetaUnit struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  httpAseGetResponseMetaUnitJSON `json:"-"`
}

// httpAseGetResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseMetaUnit]
type httpAseGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                     `json:"value,required"`
	JSON  httpAseGetResponseTop0JSON `json:"-"`
}

// httpAseGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseGetResponseTop0]
type httpAseGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseGetParams]'s query parameters as `url.Values`.
func (r HTTPAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPAseGetParamsBotClass string

const (
	HTTPAseGetParamsBotClassLikelyAutomated HTTPAseGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseGetParamsBotClassLikelyHuman     HTTPAseGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsBotClassLikelyAutomated, HTTPAseGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseGetParamsBrowserFamily string

const (
	HTTPAseGetParamsBrowserFamilyChrome  HTTPAseGetParamsBrowserFamily = "CHROME"
	HTTPAseGetParamsBrowserFamilyEdge    HTTPAseGetParamsBrowserFamily = "EDGE"
	HTTPAseGetParamsBrowserFamilyFirefox HTTPAseGetParamsBrowserFamily = "FIREFOX"
	HTTPAseGetParamsBrowserFamilySafari  HTTPAseGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsBrowserFamilyChrome, HTTPAseGetParamsBrowserFamilyEdge, HTTPAseGetParamsBrowserFamilyFirefox, HTTPAseGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseGetParamsDeviceType string

const (
	HTTPAseGetParamsDeviceTypeDesktop HTTPAseGetParamsDeviceType = "DESKTOP"
	HTTPAseGetParamsDeviceTypeMobile  HTTPAseGetParamsDeviceType = "MOBILE"
	HTTPAseGetParamsDeviceTypeOther   HTTPAseGetParamsDeviceType = "OTHER"
)

func (r HTTPAseGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsDeviceTypeDesktop, HTTPAseGetParamsDeviceTypeMobile, HTTPAseGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseGetParamsFormat string

const (
	HTTPAseGetParamsFormatJson HTTPAseGetParamsFormat = "JSON"
	HTTPAseGetParamsFormatCsv  HTTPAseGetParamsFormat = "CSV"
)

func (r HTTPAseGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsFormatJson, HTTPAseGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseGetParamsHTTPProtocol string

const (
	HTTPAseGetParamsHTTPProtocolHTTP  HTTPAseGetParamsHTTPProtocol = "HTTP"
	HTTPAseGetParamsHTTPProtocolHTTPS HTTPAseGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsHTTPProtocolHTTP, HTTPAseGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseGetParamsHTTPVersion string

const (
	HTTPAseGetParamsHTTPVersionHttPv1 HTTPAseGetParamsHTTPVersion = "HTTPv1"
	HTTPAseGetParamsHTTPVersionHttPv2 HTTPAseGetParamsHTTPVersion = "HTTPv2"
	HTTPAseGetParamsHTTPVersionHttPv3 HTTPAseGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsHTTPVersionHttPv1, HTTPAseGetParamsHTTPVersionHttPv2, HTTPAseGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseGetParamsIPVersion string

const (
	HTTPAseGetParamsIPVersionIPv4 HTTPAseGetParamsIPVersion = "IPv4"
	HTTPAseGetParamsIPVersionIPv6 HTTPAseGetParamsIPVersion = "IPv6"
)

func (r HTTPAseGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsIPVersionIPv4, HTTPAseGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseGetParamsOS string

const (
	HTTPAseGetParamsOSWindows  HTTPAseGetParamsOS = "WINDOWS"
	HTTPAseGetParamsOSMacosx   HTTPAseGetParamsOS = "MACOSX"
	HTTPAseGetParamsOSIos      HTTPAseGetParamsOS = "IOS"
	HTTPAseGetParamsOSAndroid  HTTPAseGetParamsOS = "ANDROID"
	HTTPAseGetParamsOSChromeos HTTPAseGetParamsOS = "CHROMEOS"
	HTTPAseGetParamsOSLinux    HTTPAseGetParamsOS = "LINUX"
	HTTPAseGetParamsOSSmartTv  HTTPAseGetParamsOS = "SMART_TV"
)

func (r HTTPAseGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsOSWindows, HTTPAseGetParamsOSMacosx, HTTPAseGetParamsOSIos, HTTPAseGetParamsOSAndroid, HTTPAseGetParamsOSChromeos, HTTPAseGetParamsOSLinux, HTTPAseGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseGetParamsTLSVersion string

const (
	HTTPAseGetParamsTLSVersionTlSv1_0  HTTPAseGetParamsTLSVersion = "TLSv1_0"
	HTTPAseGetParamsTLSVersionTlSv1_1  HTTPAseGetParamsTLSVersion = "TLSv1_1"
	HTTPAseGetParamsTLSVersionTlSv1_2  HTTPAseGetParamsTLSVersion = "TLSv1_2"
	HTTPAseGetParamsTLSVersionTlSv1_3  HTTPAseGetParamsTLSVersion = "TLSv1_3"
	HTTPAseGetParamsTLSVersionTlSvQuic HTTPAseGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsTLSVersionTlSv1_0, HTTPAseGetParamsTLSVersionTlSv1_1, HTTPAseGetParamsTLSVersionTlSv1_2, HTTPAseGetParamsTLSVersionTlSv1_3, HTTPAseGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseGetResponseEnvelope struct {
	Result  HTTPAseGetResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    httpAseGetResponseEnvelopeJSON `json:"-"`
}

// httpAseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseEnvelope]
type httpAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
