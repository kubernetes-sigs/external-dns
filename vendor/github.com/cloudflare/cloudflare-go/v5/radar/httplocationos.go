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

// HTTPLocationOSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationOSService] method instead.
type HTTPLocationOSService struct {
	Options []option.RequestOption
}

// NewHTTPLocationOSService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPLocationOSService(opts ...option.RequestOption) (r *HTTPLocationOSService) {
	r = &HTTPLocationOSService{}
	r.Options = opts
	return
}

// Retrieves the top locations, by HTTP requests, of the requested operating
// system.
func (r *HTTPLocationOSService) Get(ctx context.Context, os HTTPLocationOSGetParamsOS, query HTTPLocationOSGetParams, opts ...option.RequestOption) (res *HTTPLocationOSGetResponse, err error) {
	var env HTTPLocationOSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationOSGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationOSGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationOSGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationOSGetResponseJSON   `json:"-"`
}

// httpLocationOSGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationOSGetResponse]
type httpLocationOSGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationOSGetResponseMeta struct {
	ConfidenceInfo HTTPLocationOSGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationOSGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationOSGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationOSGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationOSGetResponseMetaJSON   `json:"-"`
}

// httpLocationOSGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPLocationOSGetResponseMeta]
type httpLocationOSGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationOSGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationOSGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                           `json:"level,required"`
	JSON  httpLocationOSGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationOSGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [HTTPLocationOSGetResponseMetaConfidenceInfo]
type httpLocationOSGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationOSGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                 `json:"startDate,required" format:"date-time"`
	JSON            httpLocationOSGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationOSGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPLocationOSGetResponseMetaConfidenceInfoAnnotation]
type httpLocationOSGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationOSGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationOSGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      httpLocationOSGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationOSGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPLocationOSGetResponseMetaDateRange]
type httpLocationOSGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationOSGetResponseMetaNormalization string

const (
	HTTPLocationOSGetResponseMetaNormalizationPercentage           HTTPLocationOSGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationOSGetResponseMetaNormalizationMin0Max              HTTPLocationOSGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationOSGetResponseMetaNormalizationMinMax               HTTPLocationOSGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationOSGetResponseMetaNormalizationRawValues            HTTPLocationOSGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationOSGetResponseMetaNormalizationPercentageChange     HTTPLocationOSGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationOSGetResponseMetaNormalizationRollingAverage       HTTPLocationOSGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationOSGetResponseMetaNormalizationOverlappedPercentage HTTPLocationOSGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationOSGetResponseMetaNormalizationRatio                HTTPLocationOSGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationOSGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetResponseMetaNormalizationPercentage, HTTPLocationOSGetResponseMetaNormalizationMin0Max, HTTPLocationOSGetResponseMetaNormalizationMinMax, HTTPLocationOSGetResponseMetaNormalizationRawValues, HTTPLocationOSGetResponseMetaNormalizationPercentageChange, HTTPLocationOSGetResponseMetaNormalizationRollingAverage, HTTPLocationOSGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationOSGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationOSGetResponseMetaUnit struct {
	Name  string                                `json:"name,required"`
	Value string                                `json:"value,required"`
	JSON  httpLocationOSGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationOSGetResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPLocationOSGetResponseMetaUnit]
type httpLocationOSGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationOSGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                            `json:"value,required"`
	JSON  httpLocationOSGetResponseTop0JSON `json:"-"`
}

// httpLocationOSGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPLocationOSGetResponseTop0]
type httpLocationOSGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationOSGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationOSGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPLocationOSGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationOSGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationOSGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPLocationOSGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationOSGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPLocationOSGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPLocationOSGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationOSGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationOSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Operating system.
type HTTPLocationOSGetParamsOS string

const (
	HTTPLocationOSGetParamsOSWindows  HTTPLocationOSGetParamsOS = "WINDOWS"
	HTTPLocationOSGetParamsOSMacosx   HTTPLocationOSGetParamsOS = "MACOSX"
	HTTPLocationOSGetParamsOSIos      HTTPLocationOSGetParamsOS = "IOS"
	HTTPLocationOSGetParamsOSAndroid  HTTPLocationOSGetParamsOS = "ANDROID"
	HTTPLocationOSGetParamsOSChromeos HTTPLocationOSGetParamsOS = "CHROMEOS"
	HTTPLocationOSGetParamsOSLinux    HTTPLocationOSGetParamsOS = "LINUX"
	HTTPLocationOSGetParamsOSSmartTv  HTTPLocationOSGetParamsOS = "SMART_TV"
)

func (r HTTPLocationOSGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsOSWindows, HTTPLocationOSGetParamsOSMacosx, HTTPLocationOSGetParamsOSIos, HTTPLocationOSGetParamsOSAndroid, HTTPLocationOSGetParamsOSChromeos, HTTPLocationOSGetParamsOSLinux, HTTPLocationOSGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsBotClass string

const (
	HTTPLocationOSGetParamsBotClassLikelyAutomated HTTPLocationOSGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationOSGetParamsBotClassLikelyHuman     HTTPLocationOSGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationOSGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsBotClassLikelyAutomated, HTTPLocationOSGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsBrowserFamily string

const (
	HTTPLocationOSGetParamsBrowserFamilyChrome  HTTPLocationOSGetParamsBrowserFamily = "CHROME"
	HTTPLocationOSGetParamsBrowserFamilyEdge    HTTPLocationOSGetParamsBrowserFamily = "EDGE"
	HTTPLocationOSGetParamsBrowserFamilyFirefox HTTPLocationOSGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationOSGetParamsBrowserFamilySafari  HTTPLocationOSGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationOSGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsBrowserFamilyChrome, HTTPLocationOSGetParamsBrowserFamilyEdge, HTTPLocationOSGetParamsBrowserFamilyFirefox, HTTPLocationOSGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsDeviceType string

const (
	HTTPLocationOSGetParamsDeviceTypeDesktop HTTPLocationOSGetParamsDeviceType = "DESKTOP"
	HTTPLocationOSGetParamsDeviceTypeMobile  HTTPLocationOSGetParamsDeviceType = "MOBILE"
	HTTPLocationOSGetParamsDeviceTypeOther   HTTPLocationOSGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationOSGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsDeviceTypeDesktop, HTTPLocationOSGetParamsDeviceTypeMobile, HTTPLocationOSGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationOSGetParamsFormat string

const (
	HTTPLocationOSGetParamsFormatJson HTTPLocationOSGetParamsFormat = "JSON"
	HTTPLocationOSGetParamsFormatCsv  HTTPLocationOSGetParamsFormat = "CSV"
)

func (r HTTPLocationOSGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsFormatJson, HTTPLocationOSGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsHTTPProtocol string

const (
	HTTPLocationOSGetParamsHTTPProtocolHTTP  HTTPLocationOSGetParamsHTTPProtocol = "HTTP"
	HTTPLocationOSGetParamsHTTPProtocolHTTPS HTTPLocationOSGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationOSGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsHTTPProtocolHTTP, HTTPLocationOSGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsHTTPVersion string

const (
	HTTPLocationOSGetParamsHTTPVersionHttPv1 HTTPLocationOSGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationOSGetParamsHTTPVersionHttPv2 HTTPLocationOSGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationOSGetParamsHTTPVersionHttPv3 HTTPLocationOSGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationOSGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsHTTPVersionHttPv1, HTTPLocationOSGetParamsHTTPVersionHttPv2, HTTPLocationOSGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsIPVersion string

const (
	HTTPLocationOSGetParamsIPVersionIPv4 HTTPLocationOSGetParamsIPVersion = "IPv4"
	HTTPLocationOSGetParamsIPVersionIPv6 HTTPLocationOSGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationOSGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsIPVersionIPv4, HTTPLocationOSGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationOSGetParamsTLSVersion string

const (
	HTTPLocationOSGetParamsTLSVersionTlSv1_0  HTTPLocationOSGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationOSGetParamsTLSVersionTlSv1_1  HTTPLocationOSGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationOSGetParamsTLSVersionTlSv1_2  HTTPLocationOSGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationOSGetParamsTLSVersionTlSv1_3  HTTPLocationOSGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationOSGetParamsTLSVersionTlSvQuic HTTPLocationOSGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationOSGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationOSGetParamsTLSVersionTlSv1_0, HTTPLocationOSGetParamsTLSVersionTlSv1_1, HTTPLocationOSGetParamsTLSVersionTlSv1_2, HTTPLocationOSGetParamsTLSVersionTlSv1_3, HTTPLocationOSGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationOSGetResponseEnvelope struct {
	Result  HTTPLocationOSGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    httpLocationOSGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationOSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPLocationOSGetResponseEnvelope]
type httpLocationOSGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationOSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
