// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AttackLayer7TopAseService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TopAseService] method instead.
type AttackLayer7TopAseService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7TopAseService(opts ...option.RequestOption) (r *AttackLayer7TopAseService) {
	r = &AttackLayer7TopAseService{}
	r.Options = opts
	return
}

// Retrieves the top origin autonomous systems of layer 7 attacks. Values are
// percentages of the total layer 7 attacks, with the origin autonomous systems
// determined by the client IP address.
func (r *AttackLayer7TopAseService) Origin(ctx context.Context, query AttackLayer7TopAseOriginParams, opts ...option.RequestOption) (res *AttackLayer7TopAseOriginResponse, err error) {
	var env AttackLayer7TopAseOriginResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopAseOriginResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopAseOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopAseOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopAseOriginResponseJSON   `json:"-"`
}

// attackLayer7TopAseOriginResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopAseOriginResponse]
type attackLayer7TopAseOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopAseOriginResponseMeta struct {
	ConfidenceInfo AttackLayer7TopAseOriginResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopAseOriginResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopAseOriginResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopAseOriginResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopAseOriginResponseMetaJSON   `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseMeta]
type attackLayer7TopAseOriginResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7TopAseOriginResponseMetaConfidenceInfo]
type attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopAseOriginResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopAseOriginResponseMetaDateRange]
type attackLayer7TopAseOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopAseOriginResponseMetaNormalization string

const (
	AttackLayer7TopAseOriginResponseMetaNormalizationPercentage           AttackLayer7TopAseOriginResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopAseOriginResponseMetaNormalizationMin0Max              AttackLayer7TopAseOriginResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopAseOriginResponseMetaNormalizationMinMax               AttackLayer7TopAseOriginResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopAseOriginResponseMetaNormalizationRawValues            AttackLayer7TopAseOriginResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopAseOriginResponseMetaNormalizationPercentageChange     AttackLayer7TopAseOriginResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopAseOriginResponseMetaNormalizationRollingAverage       AttackLayer7TopAseOriginResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopAseOriginResponseMetaNormalizationOverlappedPercentage AttackLayer7TopAseOriginResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopAseOriginResponseMetaNormalizationRatio                AttackLayer7TopAseOriginResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopAseOriginResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginResponseMetaNormalizationPercentage, AttackLayer7TopAseOriginResponseMetaNormalizationMin0Max, AttackLayer7TopAseOriginResponseMetaNormalizationMinMax, AttackLayer7TopAseOriginResponseMetaNormalizationRawValues, AttackLayer7TopAseOriginResponseMetaNormalizationPercentageChange, AttackLayer7TopAseOriginResponseMetaNormalizationRollingAverage, AttackLayer7TopAseOriginResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopAseOriginResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  attackLayer7TopAseOriginResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseMetaUnit]
type attackLayer7TopAseOriginResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseTop0 struct {
	OriginASN     string                                   `json:"originAsn,required"`
	OriginASNName string                                   `json:"originAsnName,required"`
	Rank          float64                                  `json:"rank,required"`
	Value         string                                   `json:"value,required"`
	JSON          attackLayer7TopAseOriginResponseTop0JSON `json:"-"`
}

// attackLayer7TopAseOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseTop0]
type attackLayer7TopAseOriginResponseTop0JSON struct {
	OriginASN     apijson.Field
	OriginASNName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginParams struct {
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
	Format param.Field[AttackLayer7TopAseOriginParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TopAseOriginParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TopAseOriginParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TopAseOriginParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopAseOriginParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopAseOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopAseOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopAseOriginParamsFormat string

const (
	AttackLayer7TopAseOriginParamsFormatJson AttackLayer7TopAseOriginParamsFormat = "JSON"
	AttackLayer7TopAseOriginParamsFormatCsv  AttackLayer7TopAseOriginParamsFormat = "CSV"
)

func (r AttackLayer7TopAseOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsFormatJson, AttackLayer7TopAseOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginParamsHTTPMethod string

const (
	AttackLayer7TopAseOriginParamsHTTPMethodGet             AttackLayer7TopAseOriginParamsHTTPMethod = "GET"
	AttackLayer7TopAseOriginParamsHTTPMethodPost            AttackLayer7TopAseOriginParamsHTTPMethod = "POST"
	AttackLayer7TopAseOriginParamsHTTPMethodDelete          AttackLayer7TopAseOriginParamsHTTPMethod = "DELETE"
	AttackLayer7TopAseOriginParamsHTTPMethodPut             AttackLayer7TopAseOriginParamsHTTPMethod = "PUT"
	AttackLayer7TopAseOriginParamsHTTPMethodHead            AttackLayer7TopAseOriginParamsHTTPMethod = "HEAD"
	AttackLayer7TopAseOriginParamsHTTPMethodPurge           AttackLayer7TopAseOriginParamsHTTPMethod = "PURGE"
	AttackLayer7TopAseOriginParamsHTTPMethodOptions         AttackLayer7TopAseOriginParamsHTTPMethod = "OPTIONS"
	AttackLayer7TopAseOriginParamsHTTPMethodPropfind        AttackLayer7TopAseOriginParamsHTTPMethod = "PROPFIND"
	AttackLayer7TopAseOriginParamsHTTPMethodMkcol           AttackLayer7TopAseOriginParamsHTTPMethod = "MKCOL"
	AttackLayer7TopAseOriginParamsHTTPMethodPatch           AttackLayer7TopAseOriginParamsHTTPMethod = "PATCH"
	AttackLayer7TopAseOriginParamsHTTPMethodACL             AttackLayer7TopAseOriginParamsHTTPMethod = "ACL"
	AttackLayer7TopAseOriginParamsHTTPMethodBcopy           AttackLayer7TopAseOriginParamsHTTPMethod = "BCOPY"
	AttackLayer7TopAseOriginParamsHTTPMethodBdelete         AttackLayer7TopAseOriginParamsHTTPMethod = "BDELETE"
	AttackLayer7TopAseOriginParamsHTTPMethodBmove           AttackLayer7TopAseOriginParamsHTTPMethod = "BMOVE"
	AttackLayer7TopAseOriginParamsHTTPMethodBpropfind       AttackLayer7TopAseOriginParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TopAseOriginParamsHTTPMethodBproppatch      AttackLayer7TopAseOriginParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TopAseOriginParamsHTTPMethodCheckin         AttackLayer7TopAseOriginParamsHTTPMethod = "CHECKIN"
	AttackLayer7TopAseOriginParamsHTTPMethodCheckout        AttackLayer7TopAseOriginParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TopAseOriginParamsHTTPMethodConnect         AttackLayer7TopAseOriginParamsHTTPMethod = "CONNECT"
	AttackLayer7TopAseOriginParamsHTTPMethodCopy            AttackLayer7TopAseOriginParamsHTTPMethod = "COPY"
	AttackLayer7TopAseOriginParamsHTTPMethodLabel           AttackLayer7TopAseOriginParamsHTTPMethod = "LABEL"
	AttackLayer7TopAseOriginParamsHTTPMethodLock            AttackLayer7TopAseOriginParamsHTTPMethod = "LOCK"
	AttackLayer7TopAseOriginParamsHTTPMethodMerge           AttackLayer7TopAseOriginParamsHTTPMethod = "MERGE"
	AttackLayer7TopAseOriginParamsHTTPMethodMkactivity      AttackLayer7TopAseOriginParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TopAseOriginParamsHTTPMethodMkworkspace     AttackLayer7TopAseOriginParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TopAseOriginParamsHTTPMethodMove            AttackLayer7TopAseOriginParamsHTTPMethod = "MOVE"
	AttackLayer7TopAseOriginParamsHTTPMethodNotify          AttackLayer7TopAseOriginParamsHTTPMethod = "NOTIFY"
	AttackLayer7TopAseOriginParamsHTTPMethodOrderpatch      AttackLayer7TopAseOriginParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TopAseOriginParamsHTTPMethodPoll            AttackLayer7TopAseOriginParamsHTTPMethod = "POLL"
	AttackLayer7TopAseOriginParamsHTTPMethodProppatch       AttackLayer7TopAseOriginParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TopAseOriginParamsHTTPMethodReport          AttackLayer7TopAseOriginParamsHTTPMethod = "REPORT"
	AttackLayer7TopAseOriginParamsHTTPMethodSearch          AttackLayer7TopAseOriginParamsHTTPMethod = "SEARCH"
	AttackLayer7TopAseOriginParamsHTTPMethodSubscribe       AttackLayer7TopAseOriginParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TopAseOriginParamsHTTPMethodTrace           AttackLayer7TopAseOriginParamsHTTPMethod = "TRACE"
	AttackLayer7TopAseOriginParamsHTTPMethodUncheckout      AttackLayer7TopAseOriginParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TopAseOriginParamsHTTPMethodUnlock          AttackLayer7TopAseOriginParamsHTTPMethod = "UNLOCK"
	AttackLayer7TopAseOriginParamsHTTPMethodUnsubscribe     AttackLayer7TopAseOriginParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TopAseOriginParamsHTTPMethodUpdate          AttackLayer7TopAseOriginParamsHTTPMethod = "UPDATE"
	AttackLayer7TopAseOriginParamsHTTPMethodVersioncontrol  AttackLayer7TopAseOriginParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TopAseOriginParamsHTTPMethodBaselinecontrol AttackLayer7TopAseOriginParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TopAseOriginParamsHTTPMethodXmsenumatts     AttackLayer7TopAseOriginParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TopAseOriginParamsHTTPMethodRpcOutData      AttackLayer7TopAseOriginParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TopAseOriginParamsHTTPMethodRpcInData       AttackLayer7TopAseOriginParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TopAseOriginParamsHTTPMethodJson            AttackLayer7TopAseOriginParamsHTTPMethod = "JSON"
	AttackLayer7TopAseOriginParamsHTTPMethodCook            AttackLayer7TopAseOriginParamsHTTPMethod = "COOK"
	AttackLayer7TopAseOriginParamsHTTPMethodTrack           AttackLayer7TopAseOriginParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TopAseOriginParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsHTTPMethodGet, AttackLayer7TopAseOriginParamsHTTPMethodPost, AttackLayer7TopAseOriginParamsHTTPMethodDelete, AttackLayer7TopAseOriginParamsHTTPMethodPut, AttackLayer7TopAseOriginParamsHTTPMethodHead, AttackLayer7TopAseOriginParamsHTTPMethodPurge, AttackLayer7TopAseOriginParamsHTTPMethodOptions, AttackLayer7TopAseOriginParamsHTTPMethodPropfind, AttackLayer7TopAseOriginParamsHTTPMethodMkcol, AttackLayer7TopAseOriginParamsHTTPMethodPatch, AttackLayer7TopAseOriginParamsHTTPMethodACL, AttackLayer7TopAseOriginParamsHTTPMethodBcopy, AttackLayer7TopAseOriginParamsHTTPMethodBdelete, AttackLayer7TopAseOriginParamsHTTPMethodBmove, AttackLayer7TopAseOriginParamsHTTPMethodBpropfind, AttackLayer7TopAseOriginParamsHTTPMethodBproppatch, AttackLayer7TopAseOriginParamsHTTPMethodCheckin, AttackLayer7TopAseOriginParamsHTTPMethodCheckout, AttackLayer7TopAseOriginParamsHTTPMethodConnect, AttackLayer7TopAseOriginParamsHTTPMethodCopy, AttackLayer7TopAseOriginParamsHTTPMethodLabel, AttackLayer7TopAseOriginParamsHTTPMethodLock, AttackLayer7TopAseOriginParamsHTTPMethodMerge, AttackLayer7TopAseOriginParamsHTTPMethodMkactivity, AttackLayer7TopAseOriginParamsHTTPMethodMkworkspace, AttackLayer7TopAseOriginParamsHTTPMethodMove, AttackLayer7TopAseOriginParamsHTTPMethodNotify, AttackLayer7TopAseOriginParamsHTTPMethodOrderpatch, AttackLayer7TopAseOriginParamsHTTPMethodPoll, AttackLayer7TopAseOriginParamsHTTPMethodProppatch, AttackLayer7TopAseOriginParamsHTTPMethodReport, AttackLayer7TopAseOriginParamsHTTPMethodSearch, AttackLayer7TopAseOriginParamsHTTPMethodSubscribe, AttackLayer7TopAseOriginParamsHTTPMethodTrace, AttackLayer7TopAseOriginParamsHTTPMethodUncheckout, AttackLayer7TopAseOriginParamsHTTPMethodUnlock, AttackLayer7TopAseOriginParamsHTTPMethodUnsubscribe, AttackLayer7TopAseOriginParamsHTTPMethodUpdate, AttackLayer7TopAseOriginParamsHTTPMethodVersioncontrol, AttackLayer7TopAseOriginParamsHTTPMethodBaselinecontrol, AttackLayer7TopAseOriginParamsHTTPMethodXmsenumatts, AttackLayer7TopAseOriginParamsHTTPMethodRpcOutData, AttackLayer7TopAseOriginParamsHTTPMethodRpcInData, AttackLayer7TopAseOriginParamsHTTPMethodJson, AttackLayer7TopAseOriginParamsHTTPMethodCook, AttackLayer7TopAseOriginParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginParamsHTTPVersion string

const (
	AttackLayer7TopAseOriginParamsHTTPVersionHttPv1 AttackLayer7TopAseOriginParamsHTTPVersion = "HTTPv1"
	AttackLayer7TopAseOriginParamsHTTPVersionHttPv2 AttackLayer7TopAseOriginParamsHTTPVersion = "HTTPv2"
	AttackLayer7TopAseOriginParamsHTTPVersionHttPv3 AttackLayer7TopAseOriginParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TopAseOriginParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsHTTPVersionHttPv1, AttackLayer7TopAseOriginParamsHTTPVersionHttPv2, AttackLayer7TopAseOriginParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginParamsIPVersion string

const (
	AttackLayer7TopAseOriginParamsIPVersionIPv4 AttackLayer7TopAseOriginParamsIPVersion = "IPv4"
	AttackLayer7TopAseOriginParamsIPVersionIPv6 AttackLayer7TopAseOriginParamsIPVersion = "IPv6"
)

func (r AttackLayer7TopAseOriginParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsIPVersionIPv4, AttackLayer7TopAseOriginParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginParamsMitigationProduct string

const (
	AttackLayer7TopAseOriginParamsMitigationProductDDoS               AttackLayer7TopAseOriginParamsMitigationProduct = "DDOS"
	AttackLayer7TopAseOriginParamsMitigationProductWAF                AttackLayer7TopAseOriginParamsMitigationProduct = "WAF"
	AttackLayer7TopAseOriginParamsMitigationProductBotManagement      AttackLayer7TopAseOriginParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopAseOriginParamsMitigationProductAccessRules        AttackLayer7TopAseOriginParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopAseOriginParamsMitigationProductIPReputation       AttackLayer7TopAseOriginParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopAseOriginParamsMitigationProductAPIShield          AttackLayer7TopAseOriginParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopAseOriginParamsMitigationProductDataLossPrevention AttackLayer7TopAseOriginParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopAseOriginParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsMitigationProductDDoS, AttackLayer7TopAseOriginParamsMitigationProductWAF, AttackLayer7TopAseOriginParamsMitigationProductBotManagement, AttackLayer7TopAseOriginParamsMitigationProductAccessRules, AttackLayer7TopAseOriginParamsMitigationProductIPReputation, AttackLayer7TopAseOriginParamsMitigationProductAPIShield, AttackLayer7TopAseOriginParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginResponseEnvelope struct {
	Result  AttackLayer7TopAseOriginResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    attackLayer7TopAseOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseEnvelope]
type attackLayer7TopAseOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
