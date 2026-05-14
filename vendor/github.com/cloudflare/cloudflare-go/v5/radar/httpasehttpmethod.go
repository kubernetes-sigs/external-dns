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

// HTTPAseHTTPMethodService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseHTTPMethodService] method instead.
type HTTPAseHTTPMethodService struct {
	Options []option.RequestOption
}

// NewHTTPAseHTTPMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseHTTPMethodService(opts ...option.RequestOption) (r *HTTPAseHTTPMethodService) {
	r = &HTTPAseHTTPMethodService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested HTTP
// version.
func (r *HTTPAseHTTPMethodService) Get(ctx context.Context, httpVersion HTTPAseHTTPMethodGetParamsHTTPVersion, query HTTPAseHTTPMethodGetParams, opts ...option.RequestOption) (res *HTTPAseHTTPMethodGetResponse, err error) {
	var env HTTPAseHTTPMethodGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseHTTPMethodGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseHTTPMethodGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseHTTPMethodGetResponseTop0 `json:"top_0,required"`
	JSON httpAseHTTPMethodGetResponseJSON   `json:"-"`
}

// httpAseHTTPMethodGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponse]
type httpAseHTTPMethodGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseHTTPMethodGetResponseMeta struct {
	ConfidenceInfo HTTPAseHTTPMethodGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseHTTPMethodGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseHTTPMethodGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseHTTPMethodGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseHTTPMethodGetResponseMetaJSON   `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponseMeta]
type httpAseHTTPMethodGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseHTTPMethodGetResponseMetaConfidenceInfo]
type httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            httpAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation]
type httpAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpAseHTTPMethodGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseHTTPMethodGetResponseMetaDateRange]
type httpAseHTTPMethodGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseHTTPMethodGetResponseMetaNormalization string

const (
	HTTPAseHTTPMethodGetResponseMetaNormalizationPercentage           HTTPAseHTTPMethodGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseHTTPMethodGetResponseMetaNormalizationMin0Max              HTTPAseHTTPMethodGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseHTTPMethodGetResponseMetaNormalizationMinMax               HTTPAseHTTPMethodGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseHTTPMethodGetResponseMetaNormalizationRawValues            HTTPAseHTTPMethodGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseHTTPMethodGetResponseMetaNormalizationPercentageChange     HTTPAseHTTPMethodGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseHTTPMethodGetResponseMetaNormalizationRollingAverage       HTTPAseHTTPMethodGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseHTTPMethodGetResponseMetaNormalizationOverlappedPercentage HTTPAseHTTPMethodGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseHTTPMethodGetResponseMetaNormalizationRatio                HTTPAseHTTPMethodGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseHTTPMethodGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetResponseMetaNormalizationPercentage, HTTPAseHTTPMethodGetResponseMetaNormalizationMin0Max, HTTPAseHTTPMethodGetResponseMetaNormalizationMinMax, HTTPAseHTTPMethodGetResponseMetaNormalizationRawValues, HTTPAseHTTPMethodGetResponseMetaNormalizationPercentageChange, HTTPAseHTTPMethodGetResponseMetaNormalizationRollingAverage, HTTPAseHTTPMethodGetResponseMetaNormalizationOverlappedPercentage, HTTPAseHTTPMethodGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  httpAseHTTPMethodGetResponseMetaUnitJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPAseHTTPMethodGetResponseMetaUnit]
type httpAseHTTPMethodGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                               `json:"value,required"`
	JSON  httpAseHTTPMethodGetResponseTop0JSON `json:"-"`
}

// httpAseHTTPMethodGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseHTTPMethodGetResponseTop0]
type httpAseHTTPMethodGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseHTTPMethodGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseHTTPMethodGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseHTTPMethodGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseHTTPMethodGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseHTTPMethodGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseHTTPMethodGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseHTTPMethodGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseHTTPMethodGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseHTTPMethodGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseHTTPMethodGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseHTTPMethodGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// HTTP version.
type HTTPAseHTTPMethodGetParamsHTTPVersion string

const (
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv1"
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv2 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv2"
	HTTPAseHTTPMethodGetParamsHTTPVersionHttPv3 HTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseHTTPMethodGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsHTTPVersionHttPv1, HTTPAseHTTPMethodGetParamsHTTPVersionHttPv2, HTTPAseHTTPMethodGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsBotClass string

const (
	HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated HTTPAseHTTPMethodGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseHTTPMethodGetParamsBotClassLikelyHuman     HTTPAseHTTPMethodGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseHTTPMethodGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsBotClassLikelyAutomated, HTTPAseHTTPMethodGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsBrowserFamily string

const (
	HTTPAseHTTPMethodGetParamsBrowserFamilyChrome  HTTPAseHTTPMethodGetParamsBrowserFamily = "CHROME"
	HTTPAseHTTPMethodGetParamsBrowserFamilyEdge    HTTPAseHTTPMethodGetParamsBrowserFamily = "EDGE"
	HTTPAseHTTPMethodGetParamsBrowserFamilyFirefox HTTPAseHTTPMethodGetParamsBrowserFamily = "FIREFOX"
	HTTPAseHTTPMethodGetParamsBrowserFamilySafari  HTTPAseHTTPMethodGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseHTTPMethodGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsBrowserFamilyChrome, HTTPAseHTTPMethodGetParamsBrowserFamilyEdge, HTTPAseHTTPMethodGetParamsBrowserFamilyFirefox, HTTPAseHTTPMethodGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsDeviceType string

const (
	HTTPAseHTTPMethodGetParamsDeviceTypeDesktop HTTPAseHTTPMethodGetParamsDeviceType = "DESKTOP"
	HTTPAseHTTPMethodGetParamsDeviceTypeMobile  HTTPAseHTTPMethodGetParamsDeviceType = "MOBILE"
	HTTPAseHTTPMethodGetParamsDeviceTypeOther   HTTPAseHTTPMethodGetParamsDeviceType = "OTHER"
)

func (r HTTPAseHTTPMethodGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsDeviceTypeDesktop, HTTPAseHTTPMethodGetParamsDeviceTypeMobile, HTTPAseHTTPMethodGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseHTTPMethodGetParamsFormat string

const (
	HTTPAseHTTPMethodGetParamsFormatJson HTTPAseHTTPMethodGetParamsFormat = "JSON"
	HTTPAseHTTPMethodGetParamsFormatCsv  HTTPAseHTTPMethodGetParamsFormat = "CSV"
)

func (r HTTPAseHTTPMethodGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsFormatJson, HTTPAseHTTPMethodGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsHTTPProtocol string

const (
	HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP  HTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTP"
	HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS HTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseHTTPMethodGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsHTTPProtocolHTTP, HTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsIPVersion string

const (
	HTTPAseHTTPMethodGetParamsIPVersionIPv4 HTTPAseHTTPMethodGetParamsIPVersion = "IPv4"
	HTTPAseHTTPMethodGetParamsIPVersionIPv6 HTTPAseHTTPMethodGetParamsIPVersion = "IPv6"
)

func (r HTTPAseHTTPMethodGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsIPVersionIPv4, HTTPAseHTTPMethodGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsOS string

const (
	HTTPAseHTTPMethodGetParamsOSWindows  HTTPAseHTTPMethodGetParamsOS = "WINDOWS"
	HTTPAseHTTPMethodGetParamsOSMacosx   HTTPAseHTTPMethodGetParamsOS = "MACOSX"
	HTTPAseHTTPMethodGetParamsOSIos      HTTPAseHTTPMethodGetParamsOS = "IOS"
	HTTPAseHTTPMethodGetParamsOSAndroid  HTTPAseHTTPMethodGetParamsOS = "ANDROID"
	HTTPAseHTTPMethodGetParamsOSChromeos HTTPAseHTTPMethodGetParamsOS = "CHROMEOS"
	HTTPAseHTTPMethodGetParamsOSLinux    HTTPAseHTTPMethodGetParamsOS = "LINUX"
	HTTPAseHTTPMethodGetParamsOSSmartTv  HTTPAseHTTPMethodGetParamsOS = "SMART_TV"
)

func (r HTTPAseHTTPMethodGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsOSWindows, HTTPAseHTTPMethodGetParamsOSMacosx, HTTPAseHTTPMethodGetParamsOSIos, HTTPAseHTTPMethodGetParamsOSAndroid, HTTPAseHTTPMethodGetParamsOSChromeos, HTTPAseHTTPMethodGetParamsOSLinux, HTTPAseHTTPMethodGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetParamsTLSVersion string

const (
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_0"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_1"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_2"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_3  HTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_3"
	HTTPAseHTTPMethodGetParamsTLSVersionTlSvQuic HTTPAseHTTPMethodGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseHTTPMethodGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2, HTTPAseHTTPMethodGetParamsTLSVersionTlSv1_3, HTTPAseHTTPMethodGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseHTTPMethodGetResponseEnvelope struct {
	Result  HTTPAseHTTPMethodGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseHTTPMethodGetResponseEnvelopeJSON `json:"-"`
}

// httpAseHTTPMethodGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseHTTPMethodGetResponseEnvelope]
type httpAseHTTPMethodGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseHTTPMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseHTTPMethodGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
