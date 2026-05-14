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

// HTTPLocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationService] method instead.
type HTTPLocationService struct {
	Options       []option.RequestOption
	BotClass      *HTTPLocationBotClassService
	DeviceType    *HTTPLocationDeviceTypeService
	HTTPProtocol  *HTTPLocationHTTPProtocolService
	HTTPMethod    *HTTPLocationHTTPMethodService
	IPVersion     *HTTPLocationIPVersionService
	OS            *HTTPLocationOSService
	TLSVersion    *HTTPLocationTLSVersionService
	BrowserFamily *HTTPLocationBrowserFamilyService
}

// NewHTTPLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPLocationService(opts ...option.RequestOption) (r *HTTPLocationService) {
	r = &HTTPLocationService{}
	r.Options = opts
	r.BotClass = NewHTTPLocationBotClassService(opts...)
	r.DeviceType = NewHTTPLocationDeviceTypeService(opts...)
	r.HTTPProtocol = NewHTTPLocationHTTPProtocolService(opts...)
	r.HTTPMethod = NewHTTPLocationHTTPMethodService(opts...)
	r.IPVersion = NewHTTPLocationIPVersionService(opts...)
	r.OS = NewHTTPLocationOSService(opts...)
	r.TLSVersion = NewHTTPLocationTLSVersionService(opts...)
	r.BrowserFamily = NewHTTPLocationBrowserFamilyService(opts...)
	return
}

// Retrieves the top locations by HTTP requests.
func (r *HTTPLocationService) Get(ctx context.Context, query HTTPLocationGetParams, opts ...option.RequestOption) (res *HTTPLocationGetResponse, err error) {
	var env HTTPLocationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationGetResponseJSON   `json:"-"`
}

// httpLocationGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponse]
type httpLocationGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationGetResponseMeta struct {
	ConfidenceInfo HTTPLocationGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationGetResponseMetaJSON   `json:"-"`
}

// httpLocationGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseMeta]
type httpLocationGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  httpLocationGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPLocationGetResponseMetaConfidenceInfo]
type httpLocationGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            httpLocationGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPLocationGetResponseMetaConfidenceInfoAnnotation]
type httpLocationGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      httpLocationGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPLocationGetResponseMetaDateRange]
type httpLocationGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationGetResponseMetaNormalization string

const (
	HTTPLocationGetResponseMetaNormalizationPercentage           HTTPLocationGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationGetResponseMetaNormalizationMin0Max              HTTPLocationGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationGetResponseMetaNormalizationMinMax               HTTPLocationGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationGetResponseMetaNormalizationRawValues            HTTPLocationGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationGetResponseMetaNormalizationPercentageChange     HTTPLocationGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationGetResponseMetaNormalizationRollingAverage       HTTPLocationGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationGetResponseMetaNormalizationOverlappedPercentage HTTPLocationGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationGetResponseMetaNormalizationRatio                HTTPLocationGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationGetResponseMetaNormalizationPercentage, HTTPLocationGetResponseMetaNormalizationMin0Max, HTTPLocationGetResponseMetaNormalizationMinMax, HTTPLocationGetResponseMetaNormalizationRawValues, HTTPLocationGetResponseMetaNormalizationPercentageChange, HTTPLocationGetResponseMetaNormalizationRollingAverage, HTTPLocationGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationGetResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  httpLocationGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationGetResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseMetaUnit]
type httpLocationGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                          `json:"value,required"`
	JSON  httpLocationGetResponseTop0JSON `json:"-"`
}

// httpLocationGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseTop0]
type httpLocationGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPLocationGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPLocationGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPLocationGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPLocationGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPLocationGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationGetParams]'s query parameters as `url.Values`.
func (r HTTPLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPLocationGetParamsBotClass string

const (
	HTTPLocationGetParamsBotClassLikelyAutomated HTTPLocationGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationGetParamsBotClassLikelyHuman     HTTPLocationGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsBotClassLikelyAutomated, HTTPLocationGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationGetParamsBrowserFamily string

const (
	HTTPLocationGetParamsBrowserFamilyChrome  HTTPLocationGetParamsBrowserFamily = "CHROME"
	HTTPLocationGetParamsBrowserFamilyEdge    HTTPLocationGetParamsBrowserFamily = "EDGE"
	HTTPLocationGetParamsBrowserFamilyFirefox HTTPLocationGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationGetParamsBrowserFamilySafari  HTTPLocationGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsBrowserFamilyChrome, HTTPLocationGetParamsBrowserFamilyEdge, HTTPLocationGetParamsBrowserFamilyFirefox, HTTPLocationGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationGetParamsDeviceType string

const (
	HTTPLocationGetParamsDeviceTypeDesktop HTTPLocationGetParamsDeviceType = "DESKTOP"
	HTTPLocationGetParamsDeviceTypeMobile  HTTPLocationGetParamsDeviceType = "MOBILE"
	HTTPLocationGetParamsDeviceTypeOther   HTTPLocationGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsDeviceTypeDesktop, HTTPLocationGetParamsDeviceTypeMobile, HTTPLocationGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationGetParamsFormat string

const (
	HTTPLocationGetParamsFormatJson HTTPLocationGetParamsFormat = "JSON"
	HTTPLocationGetParamsFormatCsv  HTTPLocationGetParamsFormat = "CSV"
)

func (r HTTPLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsFormatJson, HTTPLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationGetParamsHTTPProtocol string

const (
	HTTPLocationGetParamsHTTPProtocolHTTP  HTTPLocationGetParamsHTTPProtocol = "HTTP"
	HTTPLocationGetParamsHTTPProtocolHTTPS HTTPLocationGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsHTTPProtocolHTTP, HTTPLocationGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationGetParamsHTTPVersion string

const (
	HTTPLocationGetParamsHTTPVersionHttPv1 HTTPLocationGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationGetParamsHTTPVersionHttPv2 HTTPLocationGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationGetParamsHTTPVersionHttPv3 HTTPLocationGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsHTTPVersionHttPv1, HTTPLocationGetParamsHTTPVersionHttPv2, HTTPLocationGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationGetParamsIPVersion string

const (
	HTTPLocationGetParamsIPVersionIPv4 HTTPLocationGetParamsIPVersion = "IPv4"
	HTTPLocationGetParamsIPVersionIPv6 HTTPLocationGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsIPVersionIPv4, HTTPLocationGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationGetParamsOS string

const (
	HTTPLocationGetParamsOSWindows  HTTPLocationGetParamsOS = "WINDOWS"
	HTTPLocationGetParamsOSMacosx   HTTPLocationGetParamsOS = "MACOSX"
	HTTPLocationGetParamsOSIos      HTTPLocationGetParamsOS = "IOS"
	HTTPLocationGetParamsOSAndroid  HTTPLocationGetParamsOS = "ANDROID"
	HTTPLocationGetParamsOSChromeos HTTPLocationGetParamsOS = "CHROMEOS"
	HTTPLocationGetParamsOSLinux    HTTPLocationGetParamsOS = "LINUX"
	HTTPLocationGetParamsOSSmartTv  HTTPLocationGetParamsOS = "SMART_TV"
)

func (r HTTPLocationGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsOSWindows, HTTPLocationGetParamsOSMacosx, HTTPLocationGetParamsOSIos, HTTPLocationGetParamsOSAndroid, HTTPLocationGetParamsOSChromeos, HTTPLocationGetParamsOSLinux, HTTPLocationGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationGetParamsTLSVersion string

const (
	HTTPLocationGetParamsTLSVersionTlSv1_0  HTTPLocationGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationGetParamsTLSVersionTlSv1_1  HTTPLocationGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationGetParamsTLSVersionTlSv1_2  HTTPLocationGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationGetParamsTLSVersionTlSv1_3  HTTPLocationGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationGetParamsTLSVersionTlSvQuic HTTPLocationGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsTLSVersionTlSv1_0, HTTPLocationGetParamsTLSVersionTlSv1_1, HTTPLocationGetParamsTLSVersionTlSv1_2, HTTPLocationGetParamsTLSVersionTlSv1_3, HTTPLocationGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationGetResponseEnvelope struct {
	Result  HTTPLocationGetResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    httpLocationGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseEnvelope]
type httpLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
