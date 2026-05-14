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

// HTTPAseOSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseOSService] method instead.
type HTTPAseOSService struct {
	Options []option.RequestOption
}

// NewHTTPAseOSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPAseOSService(opts ...option.RequestOption) (r *HTTPAseOSService) {
	r = &HTTPAseOSService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested
// operating system.
func (r *HTTPAseOSService) Get(ctx context.Context, os HTTPAseOSGetParamsOS, query HTTPAseOSGetParams, opts ...option.RequestOption) (res *HTTPAseOSGetResponse, err error) {
	var env HTTPAseOSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseOSGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseOSGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseOSGetResponseTop0 `json:"top_0,required"`
	JSON httpAseOSGetResponseJSON   `json:"-"`
}

// httpAseOSGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponse]
type httpAseOSGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseOSGetResponseMeta struct {
	ConfidenceInfo HTTPAseOSGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseOSGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseOSGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseOSGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseOSGetResponseMetaJSON   `json:"-"`
}

// httpAseOSGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponseMeta]
type httpAseOSGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseOSGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseOSGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  httpAseOSGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseOSGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPAseOSGetResponseMetaConfidenceInfo]
type httpAseOSGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseOSGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            httpAseOSGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseOSGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [HTTPAseOSGetResponseMetaConfidenceInfoAnnotation]
type httpAseOSGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseOSGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseOSGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      httpAseOSGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseOSGetResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponseMetaDateRange]
type httpAseOSGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseOSGetResponseMetaNormalization string

const (
	HTTPAseOSGetResponseMetaNormalizationPercentage           HTTPAseOSGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseOSGetResponseMetaNormalizationMin0Max              HTTPAseOSGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseOSGetResponseMetaNormalizationMinMax               HTTPAseOSGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseOSGetResponseMetaNormalizationRawValues            HTTPAseOSGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseOSGetResponseMetaNormalizationPercentageChange     HTTPAseOSGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseOSGetResponseMetaNormalizationRollingAverage       HTTPAseOSGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseOSGetResponseMetaNormalizationOverlappedPercentage HTTPAseOSGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseOSGetResponseMetaNormalizationRatio                HTTPAseOSGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseOSGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseOSGetResponseMetaNormalizationPercentage, HTTPAseOSGetResponseMetaNormalizationMin0Max, HTTPAseOSGetResponseMetaNormalizationMinMax, HTTPAseOSGetResponseMetaNormalizationRawValues, HTTPAseOSGetResponseMetaNormalizationPercentageChange, HTTPAseOSGetResponseMetaNormalizationRollingAverage, HTTPAseOSGetResponseMetaNormalizationOverlappedPercentage, HTTPAseOSGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseOSGetResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  httpAseOSGetResponseMetaUnitJSON `json:"-"`
}

// httpAseOSGetResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponseMetaUnit]
type httpAseOSGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseOSGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                       `json:"value,required"`
	JSON  httpAseOSGetResponseTop0JSON `json:"-"`
}

// httpAseOSGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponseTop0]
type httpAseOSGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseOSGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseOSGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseOSGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseOSGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseOSGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseOSGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseOSGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseOSGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseOSGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseOSGetParams]'s query parameters as `url.Values`.
func (r HTTPAseOSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Operating system.
type HTTPAseOSGetParamsOS string

const (
	HTTPAseOSGetParamsOSWindows  HTTPAseOSGetParamsOS = "WINDOWS"
	HTTPAseOSGetParamsOSMacosx   HTTPAseOSGetParamsOS = "MACOSX"
	HTTPAseOSGetParamsOSIos      HTTPAseOSGetParamsOS = "IOS"
	HTTPAseOSGetParamsOSAndroid  HTTPAseOSGetParamsOS = "ANDROID"
	HTTPAseOSGetParamsOSChromeos HTTPAseOSGetParamsOS = "CHROMEOS"
	HTTPAseOSGetParamsOSLinux    HTTPAseOSGetParamsOS = "LINUX"
	HTTPAseOSGetParamsOSSmartTv  HTTPAseOSGetParamsOS = "SMART_TV"
)

func (r HTTPAseOSGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsOSWindows, HTTPAseOSGetParamsOSMacosx, HTTPAseOSGetParamsOSIos, HTTPAseOSGetParamsOSAndroid, HTTPAseOSGetParamsOSChromeos, HTTPAseOSGetParamsOSLinux, HTTPAseOSGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseOSGetParamsBotClass string

const (
	HTTPAseOSGetParamsBotClassLikelyAutomated HTTPAseOSGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseOSGetParamsBotClassLikelyHuman     HTTPAseOSGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseOSGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsBotClassLikelyAutomated, HTTPAseOSGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseOSGetParamsBrowserFamily string

const (
	HTTPAseOSGetParamsBrowserFamilyChrome  HTTPAseOSGetParamsBrowserFamily = "CHROME"
	HTTPAseOSGetParamsBrowserFamilyEdge    HTTPAseOSGetParamsBrowserFamily = "EDGE"
	HTTPAseOSGetParamsBrowserFamilyFirefox HTTPAseOSGetParamsBrowserFamily = "FIREFOX"
	HTTPAseOSGetParamsBrowserFamilySafari  HTTPAseOSGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseOSGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsBrowserFamilyChrome, HTTPAseOSGetParamsBrowserFamilyEdge, HTTPAseOSGetParamsBrowserFamilyFirefox, HTTPAseOSGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseOSGetParamsDeviceType string

const (
	HTTPAseOSGetParamsDeviceTypeDesktop HTTPAseOSGetParamsDeviceType = "DESKTOP"
	HTTPAseOSGetParamsDeviceTypeMobile  HTTPAseOSGetParamsDeviceType = "MOBILE"
	HTTPAseOSGetParamsDeviceTypeOther   HTTPAseOSGetParamsDeviceType = "OTHER"
)

func (r HTTPAseOSGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsDeviceTypeDesktop, HTTPAseOSGetParamsDeviceTypeMobile, HTTPAseOSGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseOSGetParamsFormat string

const (
	HTTPAseOSGetParamsFormatJson HTTPAseOSGetParamsFormat = "JSON"
	HTTPAseOSGetParamsFormatCsv  HTTPAseOSGetParamsFormat = "CSV"
)

func (r HTTPAseOSGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsFormatJson, HTTPAseOSGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseOSGetParamsHTTPProtocol string

const (
	HTTPAseOSGetParamsHTTPProtocolHTTP  HTTPAseOSGetParamsHTTPProtocol = "HTTP"
	HTTPAseOSGetParamsHTTPProtocolHTTPS HTTPAseOSGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseOSGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsHTTPProtocolHTTP, HTTPAseOSGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseOSGetParamsHTTPVersion string

const (
	HTTPAseOSGetParamsHTTPVersionHttPv1 HTTPAseOSGetParamsHTTPVersion = "HTTPv1"
	HTTPAseOSGetParamsHTTPVersionHttPv2 HTTPAseOSGetParamsHTTPVersion = "HTTPv2"
	HTTPAseOSGetParamsHTTPVersionHttPv3 HTTPAseOSGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseOSGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsHTTPVersionHttPv1, HTTPAseOSGetParamsHTTPVersionHttPv2, HTTPAseOSGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseOSGetParamsIPVersion string

const (
	HTTPAseOSGetParamsIPVersionIPv4 HTTPAseOSGetParamsIPVersion = "IPv4"
	HTTPAseOSGetParamsIPVersionIPv6 HTTPAseOSGetParamsIPVersion = "IPv6"
)

func (r HTTPAseOSGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsIPVersionIPv4, HTTPAseOSGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseOSGetParamsTLSVersion string

const (
	HTTPAseOSGetParamsTLSVersionTlSv1_0  HTTPAseOSGetParamsTLSVersion = "TLSv1_0"
	HTTPAseOSGetParamsTLSVersionTlSv1_1  HTTPAseOSGetParamsTLSVersion = "TLSv1_1"
	HTTPAseOSGetParamsTLSVersionTlSv1_2  HTTPAseOSGetParamsTLSVersion = "TLSv1_2"
	HTTPAseOSGetParamsTLSVersionTlSv1_3  HTTPAseOSGetParamsTLSVersion = "TLSv1_3"
	HTTPAseOSGetParamsTLSVersionTlSvQuic HTTPAseOSGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseOSGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseOSGetParamsTLSVersionTlSv1_0, HTTPAseOSGetParamsTLSVersionTlSv1_1, HTTPAseOSGetParamsTLSVersionTlSv1_2, HTTPAseOSGetParamsTLSVersionTlSv1_3, HTTPAseOSGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseOSGetResponseEnvelope struct {
	Result  HTTPAseOSGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    httpAseOSGetResponseEnvelopeJSON `json:"-"`
}

// httpAseOSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPAseOSGetResponseEnvelope]
type httpAseOSGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseOSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
