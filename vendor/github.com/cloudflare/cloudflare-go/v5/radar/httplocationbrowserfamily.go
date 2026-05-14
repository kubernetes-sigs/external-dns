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

// HTTPLocationBrowserFamilyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationBrowserFamilyService] method instead.
type HTTPLocationBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewHTTPLocationBrowserFamilyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHTTPLocationBrowserFamilyService(opts ...option.RequestOption) (r *HTTPLocationBrowserFamilyService) {
	r = &HTTPLocationBrowserFamilyService{}
	r.Options = opts
	return
}

// Retrieves the top locations, by HTTP requests, of the requested browser family.
func (r *HTTPLocationBrowserFamilyService) Get(ctx context.Context, browserFamily HTTPLocationBrowserFamilyGetParamsBrowserFamily, query HTTPLocationBrowserFamilyGetParams, opts ...option.RequestOption) (res *HTTPLocationBrowserFamilyGetResponse, err error) {
	var env HTTPLocationBrowserFamilyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/browser_family/%v", browserFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationBrowserFamilyGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationBrowserFamilyGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationBrowserFamilyGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationBrowserFamilyGetResponseJSON   `json:"-"`
}

// httpLocationBrowserFamilyGetResponseJSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponse]
type httpLocationBrowserFamilyGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationBrowserFamilyGetResponseMeta struct {
	ConfidenceInfo HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationBrowserFamilyGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationBrowserFamilyGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationBrowserFamilyGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationBrowserFamilyGetResponseMetaJSON   `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponseMeta]
type httpLocationBrowserFamilyGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo]
type httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation]
type httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      httpLocationBrowserFamilyGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationBrowserFamilyGetResponseMetaDateRange]
type httpLocationBrowserFamilyGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationBrowserFamilyGetResponseMetaNormalization string

const (
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationPercentage           HTTPLocationBrowserFamilyGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationMin0Max              HTTPLocationBrowserFamilyGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationMinMax               HTTPLocationBrowserFamilyGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationRawValues            HTTPLocationBrowserFamilyGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationPercentageChange     HTTPLocationBrowserFamilyGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationRollingAverage       HTTPLocationBrowserFamilyGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationOverlappedPercentage HTTPLocationBrowserFamilyGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationBrowserFamilyGetResponseMetaNormalizationRatio                HTTPLocationBrowserFamilyGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationBrowserFamilyGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetResponseMetaNormalizationPercentage, HTTPLocationBrowserFamilyGetResponseMetaNormalizationMin0Max, HTTPLocationBrowserFamilyGetResponseMetaNormalizationMinMax, HTTPLocationBrowserFamilyGetResponseMetaNormalizationRawValues, HTTPLocationBrowserFamilyGetResponseMetaNormalizationPercentageChange, HTTPLocationBrowserFamilyGetResponseMetaNormalizationRollingAverage, HTTPLocationBrowserFamilyGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationBrowserFamilyGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  httpLocationBrowserFamilyGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPLocationBrowserFamilyGetResponseMetaUnit]
type httpLocationBrowserFamilyGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                       `json:"value,required"`
	JSON  httpLocationBrowserFamilyGetResponseTop0JSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponseTop0]
type httpLocationBrowserFamilyGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationBrowserFamilyGetParamsBotClass] `query:"botClass"`
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
	DeviceType param.Field[[]HTTPLocationBrowserFamilyGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationBrowserFamilyGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPLocationBrowserFamilyGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPLocationBrowserFamilyGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationBrowserFamilyGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationBrowserFamilyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Browser family.
type HTTPLocationBrowserFamilyGetParamsBrowserFamily string

const (
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome  HTTPLocationBrowserFamilyGetParamsBrowserFamily = "CHROME"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyEdge    HTTPLocationBrowserFamilyGetParamsBrowserFamily = "EDGE"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyFirefox HTTPLocationBrowserFamilyGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilySafari  HTTPLocationBrowserFamilyGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationBrowserFamilyGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome, HTTPLocationBrowserFamilyGetParamsBrowserFamilyEdge, HTTPLocationBrowserFamilyGetParamsBrowserFamilyFirefox, HTTPLocationBrowserFamilyGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsBotClass string

const (
	HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated HTTPLocationBrowserFamilyGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman     HTTPLocationBrowserFamilyGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationBrowserFamilyGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated, HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsDeviceType string

const (
	HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop HTTPLocationBrowserFamilyGetParamsDeviceType = "DESKTOP"
	HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile  HTTPLocationBrowserFamilyGetParamsDeviceType = "MOBILE"
	HTTPLocationBrowserFamilyGetParamsDeviceTypeOther   HTTPLocationBrowserFamilyGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationBrowserFamilyGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop, HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile, HTTPLocationBrowserFamilyGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationBrowserFamilyGetParamsFormat string

const (
	HTTPLocationBrowserFamilyGetParamsFormatJson HTTPLocationBrowserFamilyGetParamsFormat = "JSON"
	HTTPLocationBrowserFamilyGetParamsFormatCsv  HTTPLocationBrowserFamilyGetParamsFormat = "CSV"
)

func (r HTTPLocationBrowserFamilyGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsFormatJson, HTTPLocationBrowserFamilyGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsHTTPProtocol string

const (
	HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP  HTTPLocationBrowserFamilyGetParamsHTTPProtocol = "HTTP"
	HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS HTTPLocationBrowserFamilyGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationBrowserFamilyGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP, HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsHTTPVersion string

const (
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationBrowserFamilyGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1, HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2, HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsIPVersion string

const (
	HTTPLocationBrowserFamilyGetParamsIPVersionIPv4 HTTPLocationBrowserFamilyGetParamsIPVersion = "IPv4"
	HTTPLocationBrowserFamilyGetParamsIPVersionIPv6 HTTPLocationBrowserFamilyGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationBrowserFamilyGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsIPVersionIPv4, HTTPLocationBrowserFamilyGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsOS string

const (
	HTTPLocationBrowserFamilyGetParamsOSWindows  HTTPLocationBrowserFamilyGetParamsOS = "WINDOWS"
	HTTPLocationBrowserFamilyGetParamsOSMacosx   HTTPLocationBrowserFamilyGetParamsOS = "MACOSX"
	HTTPLocationBrowserFamilyGetParamsOSIos      HTTPLocationBrowserFamilyGetParamsOS = "IOS"
	HTTPLocationBrowserFamilyGetParamsOSAndroid  HTTPLocationBrowserFamilyGetParamsOS = "ANDROID"
	HTTPLocationBrowserFamilyGetParamsOSChromeos HTTPLocationBrowserFamilyGetParamsOS = "CHROMEOS"
	HTTPLocationBrowserFamilyGetParamsOSLinux    HTTPLocationBrowserFamilyGetParamsOS = "LINUX"
	HTTPLocationBrowserFamilyGetParamsOSSmartTv  HTTPLocationBrowserFamilyGetParamsOS = "SMART_TV"
)

func (r HTTPLocationBrowserFamilyGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsOSWindows, HTTPLocationBrowserFamilyGetParamsOSMacosx, HTTPLocationBrowserFamilyGetParamsOSIos, HTTPLocationBrowserFamilyGetParamsOSAndroid, HTTPLocationBrowserFamilyGetParamsOSChromeos, HTTPLocationBrowserFamilyGetParamsOSLinux, HTTPLocationBrowserFamilyGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsTLSVersion string

const (
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_3  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSvQuic HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationBrowserFamilyGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_3, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetResponseEnvelope struct {
	Result  HTTPLocationBrowserFamilyGetResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    httpLocationBrowserFamilyGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPLocationBrowserFamilyGetResponseEnvelope]
type httpLocationBrowserFamilyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
