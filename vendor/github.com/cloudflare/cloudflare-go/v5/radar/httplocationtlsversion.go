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

// HTTPLocationTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationTLSVersionService] method instead.
type HTTPLocationTLSVersionService struct {
	Options []option.RequestOption
}

// NewHTTPLocationTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationTLSVersionService(opts ...option.RequestOption) (r *HTTPLocationTLSVersionService) {
	r = &HTTPLocationTLSVersionService{}
	r.Options = opts
	return
}

// Retrieves the top locations, by HTTP requests, of the requested TLS protocol
// version.
func (r *HTTPLocationTLSVersionService) Get(ctx context.Context, tlsVersion HTTPLocationTLSVersionGetParamsTLSVersion, query HTTPLocationTLSVersionGetParams, opts ...option.RequestOption) (res *HTTPLocationTLSVersionGetResponse, err error) {
	var env HTTPLocationTLSVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationTLSVersionGetResponse struct {
	// Metadata for the results.
	Meta HTTPLocationTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationTLSVersionGetResponseJSON   `json:"-"`
}

// httpLocationTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationTLSVersionGetResponse]
type httpLocationTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPLocationTLSVersionGetResponseMeta struct {
	ConfidenceInfo HTTPLocationTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []HTTPLocationTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPLocationTLSVersionGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPLocationTLSVersionGetResponseMetaUnit `json:"units,required"`
	JSON  httpLocationTLSVersionGetResponseMetaJSON   `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseMeta]
type httpLocationTLSVersionGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationTLSVersionGetResponseMetaConfidenceInfo]
type httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      httpLocationTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationTLSVersionGetResponseMetaDateRange]
type httpLocationTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPLocationTLSVersionGetResponseMetaNormalization string

const (
	HTTPLocationTLSVersionGetResponseMetaNormalizationPercentage           HTTPLocationTLSVersionGetResponseMetaNormalization = "PERCENTAGE"
	HTTPLocationTLSVersionGetResponseMetaNormalizationMin0Max              HTTPLocationTLSVersionGetResponseMetaNormalization = "MIN0_MAX"
	HTTPLocationTLSVersionGetResponseMetaNormalizationMinMax               HTTPLocationTLSVersionGetResponseMetaNormalization = "MIN_MAX"
	HTTPLocationTLSVersionGetResponseMetaNormalizationRawValues            HTTPLocationTLSVersionGetResponseMetaNormalization = "RAW_VALUES"
	HTTPLocationTLSVersionGetResponseMetaNormalizationPercentageChange     HTTPLocationTLSVersionGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPLocationTLSVersionGetResponseMetaNormalizationRollingAverage       HTTPLocationTLSVersionGetResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPLocationTLSVersionGetResponseMetaNormalizationOverlappedPercentage HTTPLocationTLSVersionGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPLocationTLSVersionGetResponseMetaNormalizationRatio                HTTPLocationTLSVersionGetResponseMetaNormalization = "RATIO"
)

func (r HTTPLocationTLSVersionGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetResponseMetaNormalizationPercentage, HTTPLocationTLSVersionGetResponseMetaNormalizationMin0Max, HTTPLocationTLSVersionGetResponseMetaNormalizationMinMax, HTTPLocationTLSVersionGetResponseMetaNormalizationRawValues, HTTPLocationTLSVersionGetResponseMetaNormalizationPercentageChange, HTTPLocationTLSVersionGetResponseMetaNormalizationRollingAverage, HTTPLocationTLSVersionGetResponseMetaNormalizationOverlappedPercentage, HTTPLocationTLSVersionGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  httpLocationTLSVersionGetResponseMetaUnitJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseMetaUnit]
type httpLocationTLSVersionGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                    `json:"value,required"`
	JSON  httpLocationTLSVersionGetResponseTop0JSON `json:"-"`
}

// httpLocationTLSVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseTop0]
type httpLocationTLSVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationTLSVersionGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationTLSVersionGetParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPLocationTLSVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPLocationTLSVersionGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPLocationTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPLocationTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPLocationTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPLocationTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPLocationTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// TLS version.
type HTTPLocationTLSVersionGetParamsTLSVersion string

const (
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_1  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_2  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_3  HTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationTLSVersionGetParamsTLSVersionTlSvQuic HTTPLocationTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationTLSVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_1, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_2, HTTPLocationTLSVersionGetParamsTLSVersionTlSv1_3, HTTPLocationTLSVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsBotClass string

const (
	HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated HTTPLocationTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationTLSVersionGetParamsBotClassLikelyHuman     HTTPLocationTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationTLSVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsBotClassLikelyAutomated, HTTPLocationTLSVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsBrowserFamily string

const (
	HTTPLocationTLSVersionGetParamsBrowserFamilyChrome  HTTPLocationTLSVersionGetParamsBrowserFamily = "CHROME"
	HTTPLocationTLSVersionGetParamsBrowserFamilyEdge    HTTPLocationTLSVersionGetParamsBrowserFamily = "EDGE"
	HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox HTTPLocationTLSVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationTLSVersionGetParamsBrowserFamilySafari  HTTPLocationTLSVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationTLSVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsBrowserFamilyChrome, HTTPLocationTLSVersionGetParamsBrowserFamilyEdge, HTTPLocationTLSVersionGetParamsBrowserFamilyFirefox, HTTPLocationTLSVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsDeviceType string

const (
	HTTPLocationTLSVersionGetParamsDeviceTypeDesktop HTTPLocationTLSVersionGetParamsDeviceType = "DESKTOP"
	HTTPLocationTLSVersionGetParamsDeviceTypeMobile  HTTPLocationTLSVersionGetParamsDeviceType = "MOBILE"
	HTTPLocationTLSVersionGetParamsDeviceTypeOther   HTTPLocationTLSVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationTLSVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsDeviceTypeDesktop, HTTPLocationTLSVersionGetParamsDeviceTypeMobile, HTTPLocationTLSVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPLocationTLSVersionGetParamsFormat string

const (
	HTTPLocationTLSVersionGetParamsFormatJson HTTPLocationTLSVersionGetParamsFormat = "JSON"
	HTTPLocationTLSVersionGetParamsFormatCsv  HTTPLocationTLSVersionGetParamsFormat = "CSV"
)

func (r HTTPLocationTLSVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsFormatJson, HTTPLocationTLSVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsHTTPProtocol string

const (
	HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP  HTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTP"
	HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS HTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationTLSVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsHTTPProtocolHTTP, HTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsHTTPVersion string

const (
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3 HTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationTLSVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsHTTPVersionHttPv1, HTTPLocationTLSVersionGetParamsHTTPVersionHttPv2, HTTPLocationTLSVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsIPVersion string

const (
	HTTPLocationTLSVersionGetParamsIPVersionIPv4 HTTPLocationTLSVersionGetParamsIPVersion = "IPv4"
	HTTPLocationTLSVersionGetParamsIPVersionIPv6 HTTPLocationTLSVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationTLSVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsIPVersionIPv4, HTTPLocationTLSVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetParamsOS string

const (
	HTTPLocationTLSVersionGetParamsOSWindows  HTTPLocationTLSVersionGetParamsOS = "WINDOWS"
	HTTPLocationTLSVersionGetParamsOSMacosx   HTTPLocationTLSVersionGetParamsOS = "MACOSX"
	HTTPLocationTLSVersionGetParamsOSIos      HTTPLocationTLSVersionGetParamsOS = "IOS"
	HTTPLocationTLSVersionGetParamsOSAndroid  HTTPLocationTLSVersionGetParamsOS = "ANDROID"
	HTTPLocationTLSVersionGetParamsOSChromeos HTTPLocationTLSVersionGetParamsOS = "CHROMEOS"
	HTTPLocationTLSVersionGetParamsOSLinux    HTTPLocationTLSVersionGetParamsOS = "LINUX"
	HTTPLocationTLSVersionGetParamsOSSmartTv  HTTPLocationTLSVersionGetParamsOS = "SMART_TV"
)

func (r HTTPLocationTLSVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationTLSVersionGetParamsOSWindows, HTTPLocationTLSVersionGetParamsOSMacosx, HTTPLocationTLSVersionGetParamsOSIos, HTTPLocationTLSVersionGetParamsOSAndroid, HTTPLocationTLSVersionGetParamsOSChromeos, HTTPLocationTLSVersionGetParamsOSLinux, HTTPLocationTLSVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationTLSVersionGetResponseEnvelope struct {
	Result  HTTPLocationTLSVersionGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    httpLocationTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationTLSVersionGetResponseEnvelope]
type httpLocationTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
