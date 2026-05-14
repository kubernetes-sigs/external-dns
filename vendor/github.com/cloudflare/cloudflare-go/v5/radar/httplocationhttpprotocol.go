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

// HTTPLocationHTTPProtocolService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationHTTPProtocolService] method instead.
type HTTPLocationHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewHTTPLocationHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHTTPLocationHTTPProtocolService(opts ...option.RequestOption) (r *HTTPLocationHTTPProtocolService) {
	r = &HTTPLocationHTTPProtocolService{}
	r.Options = opts
	return
}

// Retrieves the top locations, by HTTP requests, of the requested HTTP protocol.
func (r *HTTPLocationHTTPProtocolService) Get(ctx context.Context, httpProtocol HTTPLocationHTTPProtocolGetParamsHTTPProtocol, query HTTPLocationHTTPProtocolGetParams, opts ...option.RequestOption) (res *HTTPLocationHTTPProtocolGetResponse, err error) {
	var env HTTPLocationHTTPProtocolGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationHTTPProtocolGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationHTTPProtocolGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationHTTPProtocolGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationHTTPProtocolGetResponseJSON   `json:"-"`
}

// httpLocationHTTPProtocolGetResponseJSON contains the JSON metadata for the
// struct [HTTPLocationHTTPProtocolGetResponse]
type httpLocationHTTPProtocolGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationHTTPProtocolGetResponseMeta struct {
	ConfidenceInfo HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationHTTPProtocolGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationHTTPProtocolGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationHTTPProtocolGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationHTTPProtocolGetResponseMetaJSON   `json:"-"`
}

// httpLocationHTTPProtocolGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationHTTPProtocolGetResponseMeta]
type httpLocationHTTPProtocolGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  httpLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo]
type httpLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            httpLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation]
type httpLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPProtocolGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      httpLocationHTTPProtocolGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationHTTPProtocolGetResponseMetaDateRange]
type httpLocationHTTPProtocolGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationHTTPProtocolGetResponseMetaNormalization string

const (
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationPercentage           HTTPLocationHTTPProtocolGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationMin0Max              HTTPLocationHTTPProtocolGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationMinMax               HTTPLocationHTTPProtocolGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationRawValues            HTTPLocationHTTPProtocolGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationPercentageChange     HTTPLocationHTTPProtocolGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationRollingAverage       HTTPLocationHTTPProtocolGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationOverlappedPercentage HTTPLocationHTTPProtocolGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationHTTPProtocolGetResponseMetaNormalizationRatio                HTTPLocationHTTPProtocolGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationHTTPProtocolGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetResponseMetaNormalizationPercentage, HTTPLocationHTTPProtocolGetResponseMetaNormalizationMin0Max, HTTPLocationHTTPProtocolGetResponseMetaNormalizationMinMax, HTTPLocationHTTPProtocolGetResponseMetaNormalizationRawValues, HTTPLocationHTTPProtocolGetResponseMetaNormalizationPercentageChange, HTTPLocationHTTPProtocolGetResponseMetaNormalizationRollingAverage, HTTPLocationHTTPProtocolGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationHTTPProtocolGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  httpLocationHTTPProtocolGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [HTTPLocationHTTPProtocolGetResponseMetaUnit]
type httpLocationHTTPProtocolGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPProtocolGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                      `json:"value,required"`
	JSON  httpLocationHTTPProtocolGetResponseTop0JSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationHTTPProtocolGetResponseTop0]
type httpLocationHTTPProtocolGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPProtocolGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPLocationHTTPProtocolGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationHTTPProtocolGetParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationHTTPProtocolGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPLocationHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPLocationHTTPProtocolGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPLocationHTTPProtocolGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationHTTPProtocolGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// HTTP protocol (HTTP vs. HTTPS).
type HTTPLocationHTTPProtocolGetParamsHTTPProtocol string

const (
	HTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTP  HTTPLocationHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	HTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTPS HTTPLocationHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationHTTPProtocolGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTP, HTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsBotClass string

const (
	HTTPLocationHTTPProtocolGetParamsBotClassLikelyAutomated HTTPLocationHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationHTTPProtocolGetParamsBotClassLikelyHuman     HTTPLocationHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationHTTPProtocolGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsBotClassLikelyAutomated, HTTPLocationHTTPProtocolGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsBrowserFamily string

const (
	HTTPLocationHTTPProtocolGetParamsBrowserFamilyChrome  HTTPLocationHTTPProtocolGetParamsBrowserFamily = "CHROME"
	HTTPLocationHTTPProtocolGetParamsBrowserFamilyEdge    HTTPLocationHTTPProtocolGetParamsBrowserFamily = "EDGE"
	HTTPLocationHTTPProtocolGetParamsBrowserFamilyFirefox HTTPLocationHTTPProtocolGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationHTTPProtocolGetParamsBrowserFamilySafari  HTTPLocationHTTPProtocolGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationHTTPProtocolGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsBrowserFamilyChrome, HTTPLocationHTTPProtocolGetParamsBrowserFamilyEdge, HTTPLocationHTTPProtocolGetParamsBrowserFamilyFirefox, HTTPLocationHTTPProtocolGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsDeviceType string

const (
	HTTPLocationHTTPProtocolGetParamsDeviceTypeDesktop HTTPLocationHTTPProtocolGetParamsDeviceType = "DESKTOP"
	HTTPLocationHTTPProtocolGetParamsDeviceTypeMobile  HTTPLocationHTTPProtocolGetParamsDeviceType = "MOBILE"
	HTTPLocationHTTPProtocolGetParamsDeviceTypeOther   HTTPLocationHTTPProtocolGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationHTTPProtocolGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsDeviceTypeDesktop, HTTPLocationHTTPProtocolGetParamsDeviceTypeMobile, HTTPLocationHTTPProtocolGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationHTTPProtocolGetParamsFormat string

const (
	HTTPLocationHTTPProtocolGetParamsFormatJson HTTPLocationHTTPProtocolGetParamsFormat = "JSON"
	HTTPLocationHTTPProtocolGetParamsFormatCsv  HTTPLocationHTTPProtocolGetParamsFormat = "CSV"
)

func (r HTTPLocationHTTPProtocolGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsFormatJson, HTTPLocationHTTPProtocolGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsHTTPVersion string

const (
	HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv1 HTTPLocationHTTPProtocolGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv2 HTTPLocationHTTPProtocolGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv3 HTTPLocationHTTPProtocolGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationHTTPProtocolGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv1, HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv2, HTTPLocationHTTPProtocolGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsIPVersion string

const (
	HTTPLocationHTTPProtocolGetParamsIPVersionIPv4 HTTPLocationHTTPProtocolGetParamsIPVersion = "IPv4"
	HTTPLocationHTTPProtocolGetParamsIPVersionIPv6 HTTPLocationHTTPProtocolGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationHTTPProtocolGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsIPVersionIPv4, HTTPLocationHTTPProtocolGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsOS string

const (
	HTTPLocationHTTPProtocolGetParamsOSWindows  HTTPLocationHTTPProtocolGetParamsOS = "WINDOWS"
	HTTPLocationHTTPProtocolGetParamsOSMacosx   HTTPLocationHTTPProtocolGetParamsOS = "MACOSX"
	HTTPLocationHTTPProtocolGetParamsOSIos      HTTPLocationHTTPProtocolGetParamsOS = "IOS"
	HTTPLocationHTTPProtocolGetParamsOSAndroid  HTTPLocationHTTPProtocolGetParamsOS = "ANDROID"
	HTTPLocationHTTPProtocolGetParamsOSChromeos HTTPLocationHTTPProtocolGetParamsOS = "CHROMEOS"
	HTTPLocationHTTPProtocolGetParamsOSLinux    HTTPLocationHTTPProtocolGetParamsOS = "LINUX"
	HTTPLocationHTTPProtocolGetParamsOSSmartTv  HTTPLocationHTTPProtocolGetParamsOS = "SMART_TV"
)

func (r HTTPLocationHTTPProtocolGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsOSWindows, HTTPLocationHTTPProtocolGetParamsOSMacosx, HTTPLocationHTTPProtocolGetParamsOSIos, HTTPLocationHTTPProtocolGetParamsOSAndroid, HTTPLocationHTTPProtocolGetParamsOSChromeos, HTTPLocationHTTPProtocolGetParamsOSLinux, HTTPLocationHTTPProtocolGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetParamsTLSVersion string

const (
	HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_0  HTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_1  HTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_2  HTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_3  HTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationHTTPProtocolGetParamsTLSVersionTlSvQuic HTTPLocationHTTPProtocolGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationHTTPProtocolGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_0, HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_1, HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_2, HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_3, HTTPLocationHTTPProtocolGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationHTTPProtocolGetResponseEnvelope struct {
	Result  HTTPLocationHTTPProtocolGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    httpLocationHTTPProtocolGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationHTTPProtocolGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPLocationHTTPProtocolGetResponseEnvelope]
type httpLocationHTTPProtocolGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPProtocolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPProtocolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
