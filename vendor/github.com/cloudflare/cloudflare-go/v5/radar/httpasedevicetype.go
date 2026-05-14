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

// HTTPAseDeviceTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseDeviceTypeService] method instead.
type HTTPAseDeviceTypeService struct {
	Options []option.RequestOption
}

// NewHTTPAseDeviceTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseDeviceTypeService(opts ...option.RequestOption) (r *HTTPAseDeviceTypeService) {
	r = &HTTPAseDeviceTypeService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested device
// type.
func (r *HTTPAseDeviceTypeService) Get(ctx context.Context, deviceType HTTPAseDeviceTypeGetParamsDeviceType, query HTTPAseDeviceTypeGetParams, opts ...option.RequestOption) (res *HTTPAseDeviceTypeGetResponse, err error) {
	var env HTTPAseDeviceTypeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseDeviceTypeGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseDeviceTypeGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseDeviceTypeGetResponseTop0 `json:"top_0,required"`
	JSON httpAseDeviceTypeGetResponseJSON   `json:"-"`
}

// httpAseDeviceTypeGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseDeviceTypeGetResponse]
type httpAseDeviceTypeGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseDeviceTypeGetResponseMeta struct {
	ConfidenceInfo HTTPAseDeviceTypeGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseDeviceTypeGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseDeviceTypeGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseDeviceTypeGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseDeviceTypeGetResponseMetaJSON   `json:"-"`
}

// httpAseDeviceTypeGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseDeviceTypeGetResponseMeta]
type httpAseDeviceTypeGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseDeviceTypeGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  httpAseDeviceTypeGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseDeviceTypeGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseDeviceTypeGetResponseMetaConfidenceInfo]
type httpAseDeviceTypeGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            httpAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation]
type httpAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseDeviceTypeGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseDeviceTypeGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpAseDeviceTypeGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseDeviceTypeGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseDeviceTypeGetResponseMetaDateRange]
type httpAseDeviceTypeGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseDeviceTypeGetResponseMetaNormalization string

const (
	HTTPAseDeviceTypeGetResponseMetaNormalizationPercentage           HTTPAseDeviceTypeGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseDeviceTypeGetResponseMetaNormalizationMin0Max              HTTPAseDeviceTypeGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseDeviceTypeGetResponseMetaNormalizationMinMax               HTTPAseDeviceTypeGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseDeviceTypeGetResponseMetaNormalizationRawValues            HTTPAseDeviceTypeGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseDeviceTypeGetResponseMetaNormalizationPercentageChange     HTTPAseDeviceTypeGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseDeviceTypeGetResponseMetaNormalizationRollingAverage       HTTPAseDeviceTypeGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseDeviceTypeGetResponseMetaNormalizationOverlappedPercentage HTTPAseDeviceTypeGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseDeviceTypeGetResponseMetaNormalizationRatio                HTTPAseDeviceTypeGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseDeviceTypeGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetResponseMetaNormalizationPercentage, HTTPAseDeviceTypeGetResponseMetaNormalizationMin0Max, HTTPAseDeviceTypeGetResponseMetaNormalizationMinMax, HTTPAseDeviceTypeGetResponseMetaNormalizationRawValues, HTTPAseDeviceTypeGetResponseMetaNormalizationPercentageChange, HTTPAseDeviceTypeGetResponseMetaNormalizationRollingAverage, HTTPAseDeviceTypeGetResponseMetaNormalizationOverlappedPercentage, HTTPAseDeviceTypeGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  httpAseDeviceTypeGetResponseMetaUnitJSON `json:"-"`
}

// httpAseDeviceTypeGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPAseDeviceTypeGetResponseMetaUnit]
type httpAseDeviceTypeGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseDeviceTypeGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                               `json:"value,required"`
	JSON  httpAseDeviceTypeGetResponseTop0JSON `json:"-"`
}

// httpAseDeviceTypeGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseDeviceTypeGetResponseTop0]
type httpAseDeviceTypeGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseDeviceTypeGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseDeviceTypeGetParamsBrowserFamily] `query:"browserFamily"`
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
	Format param.Field[HTTPAseDeviceTypeGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseDeviceTypeGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseDeviceTypeGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseDeviceTypeGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Device type.
type HTTPAseDeviceTypeGetParamsDeviceType string

const (
	HTTPAseDeviceTypeGetParamsDeviceTypeDesktop HTTPAseDeviceTypeGetParamsDeviceType = "DESKTOP"
	HTTPAseDeviceTypeGetParamsDeviceTypeMobile  HTTPAseDeviceTypeGetParamsDeviceType = "MOBILE"
	HTTPAseDeviceTypeGetParamsDeviceTypeOther   HTTPAseDeviceTypeGetParamsDeviceType = "OTHER"
)

func (r HTTPAseDeviceTypeGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsDeviceTypeDesktop, HTTPAseDeviceTypeGetParamsDeviceTypeMobile, HTTPAseDeviceTypeGetParamsDeviceTypeOther:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsBotClass string

const (
	HTTPAseDeviceTypeGetParamsBotClassLikelyAutomated HTTPAseDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseDeviceTypeGetParamsBotClassLikelyHuman     HTTPAseDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseDeviceTypeGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsBotClassLikelyAutomated, HTTPAseDeviceTypeGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsBrowserFamily string

const (
	HTTPAseDeviceTypeGetParamsBrowserFamilyChrome  HTTPAseDeviceTypeGetParamsBrowserFamily = "CHROME"
	HTTPAseDeviceTypeGetParamsBrowserFamilyEdge    HTTPAseDeviceTypeGetParamsBrowserFamily = "EDGE"
	HTTPAseDeviceTypeGetParamsBrowserFamilyFirefox HTTPAseDeviceTypeGetParamsBrowserFamily = "FIREFOX"
	HTTPAseDeviceTypeGetParamsBrowserFamilySafari  HTTPAseDeviceTypeGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseDeviceTypeGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsBrowserFamilyChrome, HTTPAseDeviceTypeGetParamsBrowserFamilyEdge, HTTPAseDeviceTypeGetParamsBrowserFamilyFirefox, HTTPAseDeviceTypeGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseDeviceTypeGetParamsFormat string

const (
	HTTPAseDeviceTypeGetParamsFormatJson HTTPAseDeviceTypeGetParamsFormat = "JSON"
	HTTPAseDeviceTypeGetParamsFormatCsv  HTTPAseDeviceTypeGetParamsFormat = "CSV"
)

func (r HTTPAseDeviceTypeGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsFormatJson, HTTPAseDeviceTypeGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsHTTPProtocol string

const (
	HTTPAseDeviceTypeGetParamsHTTPProtocolHTTP  HTTPAseDeviceTypeGetParamsHTTPProtocol = "HTTP"
	HTTPAseDeviceTypeGetParamsHTTPProtocolHTTPS HTTPAseDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseDeviceTypeGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsHTTPProtocolHTTP, HTTPAseDeviceTypeGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsHTTPVersion string

const (
	HTTPAseDeviceTypeGetParamsHTTPVersionHttPv1 HTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	HTTPAseDeviceTypeGetParamsHTTPVersionHttPv2 HTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	HTTPAseDeviceTypeGetParamsHTTPVersionHttPv3 HTTPAseDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseDeviceTypeGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsHTTPVersionHttPv1, HTTPAseDeviceTypeGetParamsHTTPVersionHttPv2, HTTPAseDeviceTypeGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsIPVersion string

const (
	HTTPAseDeviceTypeGetParamsIPVersionIPv4 HTTPAseDeviceTypeGetParamsIPVersion = "IPv4"
	HTTPAseDeviceTypeGetParamsIPVersionIPv6 HTTPAseDeviceTypeGetParamsIPVersion = "IPv6"
)

func (r HTTPAseDeviceTypeGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsIPVersionIPv4, HTTPAseDeviceTypeGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsOS string

const (
	HTTPAseDeviceTypeGetParamsOSWindows  HTTPAseDeviceTypeGetParamsOS = "WINDOWS"
	HTTPAseDeviceTypeGetParamsOSMacosx   HTTPAseDeviceTypeGetParamsOS = "MACOSX"
	HTTPAseDeviceTypeGetParamsOSIos      HTTPAseDeviceTypeGetParamsOS = "IOS"
	HTTPAseDeviceTypeGetParamsOSAndroid  HTTPAseDeviceTypeGetParamsOS = "ANDROID"
	HTTPAseDeviceTypeGetParamsOSChromeos HTTPAseDeviceTypeGetParamsOS = "CHROMEOS"
	HTTPAseDeviceTypeGetParamsOSLinux    HTTPAseDeviceTypeGetParamsOS = "LINUX"
	HTTPAseDeviceTypeGetParamsOSSmartTv  HTTPAseDeviceTypeGetParamsOS = "SMART_TV"
)

func (r HTTPAseDeviceTypeGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsOSWindows, HTTPAseDeviceTypeGetParamsOSMacosx, HTTPAseDeviceTypeGetParamsOSIos, HTTPAseDeviceTypeGetParamsOSAndroid, HTTPAseDeviceTypeGetParamsOSChromeos, HTTPAseDeviceTypeGetParamsOSLinux, HTTPAseDeviceTypeGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetParamsTLSVersion string

const (
	HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_0  HTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_0"
	HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_1  HTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_1"
	HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_2  HTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_2"
	HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_3  HTTPAseDeviceTypeGetParamsTLSVersion = "TLSv1_3"
	HTTPAseDeviceTypeGetParamsTLSVersionTlSvQuic HTTPAseDeviceTypeGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseDeviceTypeGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_0, HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_1, HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_2, HTTPAseDeviceTypeGetParamsTLSVersionTlSv1_3, HTTPAseDeviceTypeGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseDeviceTypeGetResponseEnvelope struct {
	Result  HTTPAseDeviceTypeGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseDeviceTypeGetResponseEnvelopeJSON `json:"-"`
}

// httpAseDeviceTypeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseDeviceTypeGetResponseEnvelope]
type httpAseDeviceTypeGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseDeviceTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseDeviceTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
