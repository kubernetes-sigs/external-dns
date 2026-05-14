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

// HTTPAseIPVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseIPVersionService] method instead.
type HTTPAseIPVersionService struct {
	Options []option.RequestOption
}

// NewHTTPAseIPVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseIPVersionService(opts ...option.RequestOption) (r *HTTPAseIPVersionService) {
	r = &HTTPAseIPVersionService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested IP
// version.
func (r *HTTPAseIPVersionService) Get(ctx context.Context, ipVersion HTTPAseIPVersionGetParamsIPVersion, query HTTPAseIPVersionGetParams, opts ...option.RequestOption) (res *HTTPAseIPVersionGetResponse, err error) {
	var env HTTPAseIPVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseIPVersionGetResponse struct {
	// Metadata for the results.
	Meta HTTPAseIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpAseIPVersionGetResponseJSON   `json:"-"`
}

// httpAseIPVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponse]
type httpAseIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPAseIPVersionGetResponseMeta struct {
	ConfidenceInfo HTTPAseIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPAseIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPAseIPVersionGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPAseIPVersionGetResponseMetaUnit `json:"units,required"`
	JSON  httpAseIPVersionGetResponseMetaJSON   `json:"-"`
}

// httpAseIPVersionGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponseMeta]
type httpAseIPVersionGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  httpAseIPVersionGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [HTTPAseIPVersionGetResponseMetaConfidenceInfo]
type httpAseIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation]
type httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      httpAseIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseIPVersionGetResponseMetaDateRange]
type httpAseIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPAseIPVersionGetResponseMetaNormalization string

const (
	HTTPAseIPVersionGetResponseMetaNormalizationPercentage           HTTPAseIPVersionGetResponseMetaNormalization = "PERCENTAGE"
	HTTPAseIPVersionGetResponseMetaNormalizationMin0Max              HTTPAseIPVersionGetResponseMetaNormalization = "MIN0_MAX"
	HTTPAseIPVersionGetResponseMetaNormalizationMinMax               HTTPAseIPVersionGetResponseMetaNormalization = "MIN_MAX"
	HTTPAseIPVersionGetResponseMetaNormalizationRawValues            HTTPAseIPVersionGetResponseMetaNormalization = "RAW_VALUES"
	HTTPAseIPVersionGetResponseMetaNormalizationPercentageChange     HTTPAseIPVersionGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPAseIPVersionGetResponseMetaNormalizationRollingAverage       HTTPAseIPVersionGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPAseIPVersionGetResponseMetaNormalizationOverlappedPercentage HTTPAseIPVersionGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPAseIPVersionGetResponseMetaNormalizationRatio                HTTPAseIPVersionGetResponseMetaNormalization = "RATIO"
)

func (r HTTPAseIPVersionGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetResponseMetaNormalizationPercentage, HTTPAseIPVersionGetResponseMetaNormalizationMin0Max, HTTPAseIPVersionGetResponseMetaNormalizationMinMax, HTTPAseIPVersionGetResponseMetaNormalizationRawValues, HTTPAseIPVersionGetResponseMetaNormalizationPercentageChange, HTTPAseIPVersionGetResponseMetaNormalizationRollingAverage, HTTPAseIPVersionGetResponseMetaNormalizationOverlappedPercentage, HTTPAseIPVersionGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPAseIPVersionGetResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  httpAseIPVersionGetResponseMetaUnitJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPAseIPVersionGetResponseMetaUnit]
type httpAseIPVersionGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                              `json:"value,required"`
	JSON  httpAseIPVersionGetResponseTop0JSON `json:"-"`
}

// httpAseIPVersionGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponseTop0]
type httpAseIPVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseIPVersionGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPAseIPVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseIPVersionGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseIPVersionGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// IP version.
type HTTPAseIPVersionGetParamsIPVersion string

const (
	HTTPAseIPVersionGetParamsIPVersionIPv4 HTTPAseIPVersionGetParamsIPVersion = "IPv4"
	HTTPAseIPVersionGetParamsIPVersionIPv6 HTTPAseIPVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPAseIPVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsIPVersionIPv4, HTTPAseIPVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsBotClass string

const (
	HTTPAseIPVersionGetParamsBotClassLikelyAutomated HTTPAseIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseIPVersionGetParamsBotClassLikelyHuman     HTTPAseIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseIPVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsBotClassLikelyAutomated, HTTPAseIPVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsBrowserFamily string

const (
	HTTPAseIPVersionGetParamsBrowserFamilyChrome  HTTPAseIPVersionGetParamsBrowserFamily = "CHROME"
	HTTPAseIPVersionGetParamsBrowserFamilyEdge    HTTPAseIPVersionGetParamsBrowserFamily = "EDGE"
	HTTPAseIPVersionGetParamsBrowserFamilyFirefox HTTPAseIPVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPAseIPVersionGetParamsBrowserFamilySafari  HTTPAseIPVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseIPVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsBrowserFamilyChrome, HTTPAseIPVersionGetParamsBrowserFamilyEdge, HTTPAseIPVersionGetParamsBrowserFamilyFirefox, HTTPAseIPVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsDeviceType string

const (
	HTTPAseIPVersionGetParamsDeviceTypeDesktop HTTPAseIPVersionGetParamsDeviceType = "DESKTOP"
	HTTPAseIPVersionGetParamsDeviceTypeMobile  HTTPAseIPVersionGetParamsDeviceType = "MOBILE"
	HTTPAseIPVersionGetParamsDeviceTypeOther   HTTPAseIPVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPAseIPVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsDeviceTypeDesktop, HTTPAseIPVersionGetParamsDeviceTypeMobile, HTTPAseIPVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseIPVersionGetParamsFormat string

const (
	HTTPAseIPVersionGetParamsFormatJson HTTPAseIPVersionGetParamsFormat = "JSON"
	HTTPAseIPVersionGetParamsFormatCsv  HTTPAseIPVersionGetParamsFormat = "CSV"
)

func (r HTTPAseIPVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsFormatJson, HTTPAseIPVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsHTTPProtocol string

const (
	HTTPAseIPVersionGetParamsHTTPProtocolHTTP  HTTPAseIPVersionGetParamsHTTPProtocol = "HTTP"
	HTTPAseIPVersionGetParamsHTTPProtocolHTTPS HTTPAseIPVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseIPVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsHTTPProtocolHTTP, HTTPAseIPVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsHTTPVersion string

const (
	HTTPAseIPVersionGetParamsHTTPVersionHttPv1 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPAseIPVersionGetParamsHTTPVersionHttPv2 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPAseIPVersionGetParamsHTTPVersionHttPv3 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseIPVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsHTTPVersionHttPv1, HTTPAseIPVersionGetParamsHTTPVersionHttPv2, HTTPAseIPVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsOS string

const (
	HTTPAseIPVersionGetParamsOSWindows  HTTPAseIPVersionGetParamsOS = "WINDOWS"
	HTTPAseIPVersionGetParamsOSMacosx   HTTPAseIPVersionGetParamsOS = "MACOSX"
	HTTPAseIPVersionGetParamsOSIos      HTTPAseIPVersionGetParamsOS = "IOS"
	HTTPAseIPVersionGetParamsOSAndroid  HTTPAseIPVersionGetParamsOS = "ANDROID"
	HTTPAseIPVersionGetParamsOSChromeos HTTPAseIPVersionGetParamsOS = "CHROMEOS"
	HTTPAseIPVersionGetParamsOSLinux    HTTPAseIPVersionGetParamsOS = "LINUX"
	HTTPAseIPVersionGetParamsOSSmartTv  HTTPAseIPVersionGetParamsOS = "SMART_TV"
)

func (r HTTPAseIPVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsOSWindows, HTTPAseIPVersionGetParamsOSMacosx, HTTPAseIPVersionGetParamsOSIos, HTTPAseIPVersionGetParamsOSAndroid, HTTPAseIPVersionGetParamsOSChromeos, HTTPAseIPVersionGetParamsOSLinux, HTTPAseIPVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsTLSVersion string

const (
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_0  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_1  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_2  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_3  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPAseIPVersionGetParamsTLSVersionTlSvQuic HTTPAseIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseIPVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsTLSVersionTlSv1_0, HTTPAseIPVersionGetParamsTLSVersionTlSv1_1, HTTPAseIPVersionGetParamsTLSVersionTlSv1_2, HTTPAseIPVersionGetParamsTLSVersionTlSv1_3, HTTPAseIPVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseIPVersionGetResponseEnvelope struct {
	Result  HTTPAseIPVersionGetResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    httpAseIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpAseIPVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseIPVersionGetResponseEnvelope]
type httpAseIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
