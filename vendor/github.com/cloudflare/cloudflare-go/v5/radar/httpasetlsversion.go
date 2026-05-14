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

// HTTPAseTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseTLSVersionService] method instead.
type HTTPAseTLSVersionService struct {
	Options []option.RequestOption
}

// NewHTTPAseTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseTLSVersionService(opts ...option.RequestOption) (r *HTTPAseTLSVersionService) {
	r = &HTTPAseTLSVersionService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested TLS
// protocol version.
func (r *HTTPAseTLSVersionService) Get(ctx context.Context, tlsVersion HTTPAseTLSVersionGetParamsTLSVersion, query HTTPAseTLSVersionGetParams, opts ...option.RequestOption) (res *HTTPAseTLSVersionGetResponse, err error) {
	var env HTTPAseTLSVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseTLSVersionGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpAseTLSVersionGetResponseJSON   `json:"-"`
}

// httpAseTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponse]
type httpAseTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseTLSVersionGetResponseMeta struct {
	ConfidenceInfo HTTPAseTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseTLSVersionGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseTLSVersionGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseTLSVersionGetResponseMetaJSON   `json:"-"`
}

// httpAseTLSVersionGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponseMeta]
type httpAseTLSVersionGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  httpAseTLSVersionGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseTLSVersionGetResponseMetaConfidenceInfo]
type httpAseTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpAseTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseMetaDateRange]
type httpAseTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseTLSVersionGetResponseMetaNormalization string

const (
	HTTPAseTLSVersionGetResponseMetaNormalizationPercentage           HTTPAseTLSVersionGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseTLSVersionGetResponseMetaNormalizationMin0Max              HTTPAseTLSVersionGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseTLSVersionGetResponseMetaNormalizationMinMax               HTTPAseTLSVersionGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseTLSVersionGetResponseMetaNormalizationRawValues            HTTPAseTLSVersionGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseTLSVersionGetResponseMetaNormalizationPercentageChange     HTTPAseTLSVersionGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseTLSVersionGetResponseMetaNormalizationRollingAverage       HTTPAseTLSVersionGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseTLSVersionGetResponseMetaNormalizationOverlappedPercentage HTTPAseTLSVersionGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseTLSVersionGetResponseMetaNormalizationRatio                HTTPAseTLSVersionGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseTLSVersionGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetResponseMetaNormalizationPercentage, HTTPAseTLSVersionGetResponseMetaNormalizationMin0Max, HTTPAseTLSVersionGetResponseMetaNormalizationMinMax, HTTPAseTLSVersionGetResponseMetaNormalizationRawValues, HTTPAseTLSVersionGetResponseMetaNormalizationPercentageChange, HTTPAseTLSVersionGetResponseMetaNormalizationRollingAverage, HTTPAseTLSVersionGetResponseMetaNormalizationOverlappedPercentage, HTTPAseTLSVersionGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  httpAseTLSVersionGetResponseMetaUnitJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseMetaUnit]
type httpAseTLSVersionGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                               `json:"value,required"`
	JSON  httpAseTLSVersionGetResponseTop0JSON `json:"-"`
}

// httpAseTLSVersionGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponseTop0]
type httpAseTLSVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseTLSVersionGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseTLSVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseTLSVersionGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPAseTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// TLS version.
type HTTPAseTLSVersionGetParamsTLSVersion string

const (
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic HTTPAseTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseTLSVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3, HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsBotClass string

const (
	HTTPAseTLSVersionGetParamsBotClassLikelyAutomated HTTPAseTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseTLSVersionGetParamsBotClassLikelyHuman     HTTPAseTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseTLSVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsBotClassLikelyAutomated, HTTPAseTLSVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsBrowserFamily string

const (
	HTTPAseTLSVersionGetParamsBrowserFamilyChrome  HTTPAseTLSVersionGetParamsBrowserFamily = "CHROME"
	HTTPAseTLSVersionGetParamsBrowserFamilyEdge    HTTPAseTLSVersionGetParamsBrowserFamily = "EDGE"
	HTTPAseTLSVersionGetParamsBrowserFamilyFirefox HTTPAseTLSVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPAseTLSVersionGetParamsBrowserFamilySafari  HTTPAseTLSVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseTLSVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsBrowserFamilyChrome, HTTPAseTLSVersionGetParamsBrowserFamilyEdge, HTTPAseTLSVersionGetParamsBrowserFamilyFirefox, HTTPAseTLSVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsDeviceType string

const (
	HTTPAseTLSVersionGetParamsDeviceTypeDesktop HTTPAseTLSVersionGetParamsDeviceType = "DESKTOP"
	HTTPAseTLSVersionGetParamsDeviceTypeMobile  HTTPAseTLSVersionGetParamsDeviceType = "MOBILE"
	HTTPAseTLSVersionGetParamsDeviceTypeOther   HTTPAseTLSVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPAseTLSVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsDeviceTypeDesktop, HTTPAseTLSVersionGetParamsDeviceTypeMobile, HTTPAseTLSVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseTLSVersionGetParamsFormat string

const (
	HTTPAseTLSVersionGetParamsFormatJson HTTPAseTLSVersionGetParamsFormat = "JSON"
	HTTPAseTLSVersionGetParamsFormatCsv  HTTPAseTLSVersionGetParamsFormat = "CSV"
)

func (r HTTPAseTLSVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsFormatJson, HTTPAseTLSVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPProtocol string

const (
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTP  HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTP"
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseTLSVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPProtocolHTTP, HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPVersion string

const (
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv1 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv2 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv3 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseTLSVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPVersionHttPv1, HTTPAseTLSVersionGetParamsHTTPVersionHttPv2, HTTPAseTLSVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsIPVersion string

const (
	HTTPAseTLSVersionGetParamsIPVersionIPv4 HTTPAseTLSVersionGetParamsIPVersion = "IPv4"
	HTTPAseTLSVersionGetParamsIPVersionIPv6 HTTPAseTLSVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPAseTLSVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsIPVersionIPv4, HTTPAseTLSVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsOS string

const (
	HTTPAseTLSVersionGetParamsOSWindows  HTTPAseTLSVersionGetParamsOS = "WINDOWS"
	HTTPAseTLSVersionGetParamsOSMacosx   HTTPAseTLSVersionGetParamsOS = "MACOSX"
	HTTPAseTLSVersionGetParamsOSIos      HTTPAseTLSVersionGetParamsOS = "IOS"
	HTTPAseTLSVersionGetParamsOSAndroid  HTTPAseTLSVersionGetParamsOS = "ANDROID"
	HTTPAseTLSVersionGetParamsOSChromeos HTTPAseTLSVersionGetParamsOS = "CHROMEOS"
	HTTPAseTLSVersionGetParamsOSLinux    HTTPAseTLSVersionGetParamsOS = "LINUX"
	HTTPAseTLSVersionGetParamsOSSmartTv  HTTPAseTLSVersionGetParamsOS = "SMART_TV"
)

func (r HTTPAseTLSVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsOSWindows, HTTPAseTLSVersionGetParamsOSMacosx, HTTPAseTLSVersionGetParamsOSIos, HTTPAseTLSVersionGetParamsOSAndroid, HTTPAseTLSVersionGetParamsOSChromeos, HTTPAseTLSVersionGetParamsOSLinux, HTTPAseTLSVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetResponseEnvelope struct {
	Result  HTTPAseTLSVersionGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpAseTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseEnvelope]
type httpAseTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
