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

// HTTPLocationIPVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationIPVersionService] method instead.
type HTTPLocationIPVersionService struct {
	Options []option.RequestOption
}

// NewHTTPLocationIPVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationIPVersionService(opts ...option.RequestOption) (r *HTTPLocationIPVersionService) {
	r = &HTTPLocationIPVersionService{}
	r.Options = opts
	return
}

// Retrieves the top locations, by HTTP requests, of the requested IP version.
func (r *HTTPLocationIPVersionService) Get(ctx context.Context, ipVersion HTTPLocationIPVersionGetParamsIPVersion, query HTTPLocationIPVersionGetParams, opts ...option.RequestOption) (res *HTTPLocationIPVersionGetResponse, err error) {
	var env HTTPLocationIPVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationIPVersionGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationIPVersionGetResponseJSON   `json:"-"`
}

// httpLocationIPVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationIPVersionGetResponse]
type httpLocationIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationIPVersionGetResponseMeta struct {
	ConfidenceInfo HTTPLocationIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationIPVersionGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationIPVersionGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationIPVersionGetResponseMetaJSON   `json:"-"`
}

// httpLocationIPVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseMeta]
type httpLocationIPVersionGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  httpLocationIPVersionGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationIPVersionGetResponseMetaConfidenceInfo]
type httpLocationIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation]
type httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      httpLocationIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPLocationIPVersionGetResponseMetaDateRange]
type httpLocationIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationIPVersionGetResponseMetaNormalization string

const (
	HTTPLocationIPVersionGetResponseMetaNormalizationPercentage           HTTPLocationIPVersionGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationIPVersionGetResponseMetaNormalizationMin0Max              HTTPLocationIPVersionGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationIPVersionGetResponseMetaNormalizationMinMax               HTTPLocationIPVersionGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationIPVersionGetResponseMetaNormalizationRawValues            HTTPLocationIPVersionGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationIPVersionGetResponseMetaNormalizationPercentageChange     HTTPLocationIPVersionGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationIPVersionGetResponseMetaNormalizationRollingAverage       HTTPLocationIPVersionGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationIPVersionGetResponseMetaNormalizationOverlappedPercentage HTTPLocationIPVersionGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationIPVersionGetResponseMetaNormalizationRatio                HTTPLocationIPVersionGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationIPVersionGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetResponseMetaNormalizationPercentage, HTTPLocationIPVersionGetResponseMetaNormalizationMin0Max, HTTPLocationIPVersionGetResponseMetaNormalizationMinMax, HTTPLocationIPVersionGetResponseMetaNormalizationRawValues, HTTPLocationIPVersionGetResponseMetaNormalizationPercentageChange, HTTPLocationIPVersionGetResponseMetaNormalizationRollingAverage, HTTPLocationIPVersionGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationIPVersionGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  httpLocationIPVersionGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseMetaUnit]
type httpLocationIPVersionGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                   `json:"value,required"`
	JSON  httpLocationIPVersionGetResponseTop0JSON `json:"-"`
}

// httpLocationIPVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseTop0]
type httpLocationIPVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationIPVersionGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPLocationIPVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationIPVersionGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPLocationIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPLocationIPVersionGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPLocationIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// IP version.
type HTTPLocationIPVersionGetParamsIPVersion string

const (
	HTTPLocationIPVersionGetParamsIPVersionIPv4 HTTPLocationIPVersionGetParamsIPVersion = "IPv4"
	HTTPLocationIPVersionGetParamsIPVersionIPv6 HTTPLocationIPVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationIPVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsIPVersionIPv4, HTTPLocationIPVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsBotClass string

const (
	HTTPLocationIPVersionGetParamsBotClassLikelyAutomated HTTPLocationIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationIPVersionGetParamsBotClassLikelyHuman     HTTPLocationIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationIPVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsBotClassLikelyAutomated, HTTPLocationIPVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsBrowserFamily string

const (
	HTTPLocationIPVersionGetParamsBrowserFamilyChrome  HTTPLocationIPVersionGetParamsBrowserFamily = "CHROME"
	HTTPLocationIPVersionGetParamsBrowserFamilyEdge    HTTPLocationIPVersionGetParamsBrowserFamily = "EDGE"
	HTTPLocationIPVersionGetParamsBrowserFamilyFirefox HTTPLocationIPVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationIPVersionGetParamsBrowserFamilySafari  HTTPLocationIPVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationIPVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsBrowserFamilyChrome, HTTPLocationIPVersionGetParamsBrowserFamilyEdge, HTTPLocationIPVersionGetParamsBrowserFamilyFirefox, HTTPLocationIPVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsDeviceType string

const (
	HTTPLocationIPVersionGetParamsDeviceTypeDesktop HTTPLocationIPVersionGetParamsDeviceType = "DESKTOP"
	HTTPLocationIPVersionGetParamsDeviceTypeMobile  HTTPLocationIPVersionGetParamsDeviceType = "MOBILE"
	HTTPLocationIPVersionGetParamsDeviceTypeOther   HTTPLocationIPVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationIPVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsDeviceTypeDesktop, HTTPLocationIPVersionGetParamsDeviceTypeMobile, HTTPLocationIPVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationIPVersionGetParamsFormat string

const (
	HTTPLocationIPVersionGetParamsFormatJson HTTPLocationIPVersionGetParamsFormat = "JSON"
	HTTPLocationIPVersionGetParamsFormatCsv  HTTPLocationIPVersionGetParamsFormat = "CSV"
)

func (r HTTPLocationIPVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsFormatJson, HTTPLocationIPVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsHTTPProtocol string

const (
	HTTPLocationIPVersionGetParamsHTTPProtocolHTTP  HTTPLocationIPVersionGetParamsHTTPProtocol = "HTTP"
	HTTPLocationIPVersionGetParamsHTTPProtocolHTTPS HTTPLocationIPVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationIPVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsHTTPProtocolHTTP, HTTPLocationIPVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsHTTPVersion string

const (
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv1 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv2 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv3 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationIPVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsHTTPVersionHttPv1, HTTPLocationIPVersionGetParamsHTTPVersionHttPv2, HTTPLocationIPVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsOS string

const (
	HTTPLocationIPVersionGetParamsOSWindows  HTTPLocationIPVersionGetParamsOS = "WINDOWS"
	HTTPLocationIPVersionGetParamsOSMacosx   HTTPLocationIPVersionGetParamsOS = "MACOSX"
	HTTPLocationIPVersionGetParamsOSIos      HTTPLocationIPVersionGetParamsOS = "IOS"
	HTTPLocationIPVersionGetParamsOSAndroid  HTTPLocationIPVersionGetParamsOS = "ANDROID"
	HTTPLocationIPVersionGetParamsOSChromeos HTTPLocationIPVersionGetParamsOS = "CHROMEOS"
	HTTPLocationIPVersionGetParamsOSLinux    HTTPLocationIPVersionGetParamsOS = "LINUX"
	HTTPLocationIPVersionGetParamsOSSmartTv  HTTPLocationIPVersionGetParamsOS = "SMART_TV"
)

func (r HTTPLocationIPVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsOSWindows, HTTPLocationIPVersionGetParamsOSMacosx, HTTPLocationIPVersionGetParamsOSIos, HTTPLocationIPVersionGetParamsOSAndroid, HTTPLocationIPVersionGetParamsOSChromeos, HTTPLocationIPVersionGetParamsOSLinux, HTTPLocationIPVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsTLSVersion string

const (
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_0  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_1  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_2  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_3  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationIPVersionGetParamsTLSVersionTlSvQuic HTTPLocationIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationIPVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsTLSVersionTlSv1_0, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_1, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_2, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_3, HTTPLocationIPVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetResponseEnvelope struct {
	Result  HTTPLocationIPVersionGetResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    httpLocationIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationIPVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseEnvelope]
type httpLocationIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
