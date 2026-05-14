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

// AS112Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112Service] method instead.
type AS112Service struct {
	Options          []option.RequestOption
	Summary          *AS112SummaryService
	TimeseriesGroups *AS112TimeseriesGroupService
	Top              *AS112TopService
}

// NewAS112Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAS112Service(opts ...option.RequestOption) (r *AS112Service) {
	r = &AS112Service{}
	r.Options = opts
	r.Summary = NewAS112SummaryService(opts...)
	r.TimeseriesGroups = NewAS112TimeseriesGroupService(opts...)
	r.Top = NewAS112TopService(opts...)
	return
}

// Retrieves the AS112 DNS queries over time.
func (r *AS112Service) Timeseries(ctx context.Context, query AS112TimeseriesParams, opts ...option.RequestOption) (res *AS112TimeseriesResponse, err error) {
	var env AS112TimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TimeseriesResponse struct {
	// Metadata for the results.
	Meta        AS112TimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AS112TimeseriesResponse `json:"-,extras"`
	JSON        as112TimeseriesResponseJSON        `json:"-"`
}

// as112TimeseriesResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponse]
type as112TimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesResponseMetaJSON   `json:"-"`
}

// as112TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseMeta]
type as112TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesResponseMetaAggInterval string

const (
	AS112TimeseriesResponseMetaAggIntervalFifteenMinutes AS112TimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesResponseMetaAggIntervalOneHour        AS112TimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesResponseMetaAggIntervalOneDay         AS112TimeseriesResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesResponseMetaAggIntervalOneWeek        AS112TimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesResponseMetaAggIntervalOneMonth       AS112TimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesResponseMetaAggIntervalOneHour, AS112TimeseriesResponseMetaAggIntervalOneDay, AS112TimeseriesResponseMetaAggIntervalOneWeek, AS112TimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  as112TimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaConfidenceInfo]
type as112TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TimeseriesResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaDateRange]
type as112TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesResponseMetaNormalization string

const (
	AS112TimeseriesResponseMetaNormalizationPercentage           AS112TimeseriesResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesResponseMetaNormalizationMin0Max              AS112TimeseriesResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesResponseMetaNormalizationMinMax               AS112TimeseriesResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesResponseMetaNormalizationRawValues            AS112TimeseriesResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesResponseMetaNormalizationPercentageChange     AS112TimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesResponseMetaNormalizationRollingAverage       AS112TimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesResponseMetaNormalizationOverlappedPercentage AS112TimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesResponseMetaNormalizationRatio                AS112TimeseriesResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaNormalizationPercentage, AS112TimeseriesResponseMetaNormalizationMin0Max, AS112TimeseriesResponseMetaNormalizationMinMax, AS112TimeseriesResponseMetaNormalizationRawValues, AS112TimeseriesResponseMetaNormalizationPercentageChange, AS112TimeseriesResponseMetaNormalizationRollingAverage, AS112TimeseriesResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  as112TimeseriesResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseMetaUnit]
type as112TimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112TimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112TimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112TimeseriesParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesParams]'s query parameters as `url.Values`.
func (r AS112TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesParamsAggInterval string

const (
	AS112TimeseriesParamsAggInterval15m AS112TimeseriesParamsAggInterval = "15m"
	AS112TimeseriesParamsAggInterval1h  AS112TimeseriesParamsAggInterval = "1h"
	AS112TimeseriesParamsAggInterval1d  AS112TimeseriesParamsAggInterval = "1d"
	AS112TimeseriesParamsAggInterval1w  AS112TimeseriesParamsAggInterval = "1w"
)

func (r AS112TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsAggInterval15m, AS112TimeseriesParamsAggInterval1h, AS112TimeseriesParamsAggInterval1d, AS112TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesParamsFormat string

const (
	AS112TimeseriesParamsFormatJson AS112TimeseriesParamsFormat = "JSON"
	AS112TimeseriesParamsFormatCsv  AS112TimeseriesParamsFormat = "CSV"
)

func (r AS112TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsFormatJson, AS112TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112TimeseriesParamsProtocol string

const (
	AS112TimeseriesParamsProtocolUdp   AS112TimeseriesParamsProtocol = "UDP"
	AS112TimeseriesParamsProtocolTCP   AS112TimeseriesParamsProtocol = "TCP"
	AS112TimeseriesParamsProtocolHTTPS AS112TimeseriesParamsProtocol = "HTTPS"
	AS112TimeseriesParamsProtocolTLS   AS112TimeseriesParamsProtocol = "TLS"
)

func (r AS112TimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsProtocolUdp, AS112TimeseriesParamsProtocolTCP, AS112TimeseriesParamsProtocolHTTPS, AS112TimeseriesParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112TimeseriesParamsQueryType string

const (
	AS112TimeseriesParamsQueryTypeA          AS112TimeseriesParamsQueryType = "A"
	AS112TimeseriesParamsQueryTypeAAAA       AS112TimeseriesParamsQueryType = "AAAA"
	AS112TimeseriesParamsQueryTypeA6         AS112TimeseriesParamsQueryType = "A6"
	AS112TimeseriesParamsQueryTypeAfsdb      AS112TimeseriesParamsQueryType = "AFSDB"
	AS112TimeseriesParamsQueryTypeAny        AS112TimeseriesParamsQueryType = "ANY"
	AS112TimeseriesParamsQueryTypeApl        AS112TimeseriesParamsQueryType = "APL"
	AS112TimeseriesParamsQueryTypeAtma       AS112TimeseriesParamsQueryType = "ATMA"
	AS112TimeseriesParamsQueryTypeAXFR       AS112TimeseriesParamsQueryType = "AXFR"
	AS112TimeseriesParamsQueryTypeCAA        AS112TimeseriesParamsQueryType = "CAA"
	AS112TimeseriesParamsQueryTypeCdnskey    AS112TimeseriesParamsQueryType = "CDNSKEY"
	AS112TimeseriesParamsQueryTypeCds        AS112TimeseriesParamsQueryType = "CDS"
	AS112TimeseriesParamsQueryTypeCERT       AS112TimeseriesParamsQueryType = "CERT"
	AS112TimeseriesParamsQueryTypeCNAME      AS112TimeseriesParamsQueryType = "CNAME"
	AS112TimeseriesParamsQueryTypeCsync      AS112TimeseriesParamsQueryType = "CSYNC"
	AS112TimeseriesParamsQueryTypeDhcid      AS112TimeseriesParamsQueryType = "DHCID"
	AS112TimeseriesParamsQueryTypeDlv        AS112TimeseriesParamsQueryType = "DLV"
	AS112TimeseriesParamsQueryTypeDname      AS112TimeseriesParamsQueryType = "DNAME"
	AS112TimeseriesParamsQueryTypeDNSKEY     AS112TimeseriesParamsQueryType = "DNSKEY"
	AS112TimeseriesParamsQueryTypeDoa        AS112TimeseriesParamsQueryType = "DOA"
	AS112TimeseriesParamsQueryTypeDS         AS112TimeseriesParamsQueryType = "DS"
	AS112TimeseriesParamsQueryTypeEid        AS112TimeseriesParamsQueryType = "EID"
	AS112TimeseriesParamsQueryTypeEui48      AS112TimeseriesParamsQueryType = "EUI48"
	AS112TimeseriesParamsQueryTypeEui64      AS112TimeseriesParamsQueryType = "EUI64"
	AS112TimeseriesParamsQueryTypeGpos       AS112TimeseriesParamsQueryType = "GPOS"
	AS112TimeseriesParamsQueryTypeGid        AS112TimeseriesParamsQueryType = "GID"
	AS112TimeseriesParamsQueryTypeHinfo      AS112TimeseriesParamsQueryType = "HINFO"
	AS112TimeseriesParamsQueryTypeHip        AS112TimeseriesParamsQueryType = "HIP"
	AS112TimeseriesParamsQueryTypeHTTPS      AS112TimeseriesParamsQueryType = "HTTPS"
	AS112TimeseriesParamsQueryTypeIpseckey   AS112TimeseriesParamsQueryType = "IPSECKEY"
	AS112TimeseriesParamsQueryTypeIsdn       AS112TimeseriesParamsQueryType = "ISDN"
	AS112TimeseriesParamsQueryTypeIxfr       AS112TimeseriesParamsQueryType = "IXFR"
	AS112TimeseriesParamsQueryTypeKey        AS112TimeseriesParamsQueryType = "KEY"
	AS112TimeseriesParamsQueryTypeKx         AS112TimeseriesParamsQueryType = "KX"
	AS112TimeseriesParamsQueryTypeL32        AS112TimeseriesParamsQueryType = "L32"
	AS112TimeseriesParamsQueryTypeL64        AS112TimeseriesParamsQueryType = "L64"
	AS112TimeseriesParamsQueryTypeLOC        AS112TimeseriesParamsQueryType = "LOC"
	AS112TimeseriesParamsQueryTypeLp         AS112TimeseriesParamsQueryType = "LP"
	AS112TimeseriesParamsQueryTypeMaila      AS112TimeseriesParamsQueryType = "MAILA"
	AS112TimeseriesParamsQueryTypeMailb      AS112TimeseriesParamsQueryType = "MAILB"
	AS112TimeseriesParamsQueryTypeMB         AS112TimeseriesParamsQueryType = "MB"
	AS112TimeseriesParamsQueryTypeMd         AS112TimeseriesParamsQueryType = "MD"
	AS112TimeseriesParamsQueryTypeMf         AS112TimeseriesParamsQueryType = "MF"
	AS112TimeseriesParamsQueryTypeMg         AS112TimeseriesParamsQueryType = "MG"
	AS112TimeseriesParamsQueryTypeMinfo      AS112TimeseriesParamsQueryType = "MINFO"
	AS112TimeseriesParamsQueryTypeMr         AS112TimeseriesParamsQueryType = "MR"
	AS112TimeseriesParamsQueryTypeMX         AS112TimeseriesParamsQueryType = "MX"
	AS112TimeseriesParamsQueryTypeNAPTR      AS112TimeseriesParamsQueryType = "NAPTR"
	AS112TimeseriesParamsQueryTypeNb         AS112TimeseriesParamsQueryType = "NB"
	AS112TimeseriesParamsQueryTypeNbstat     AS112TimeseriesParamsQueryType = "NBSTAT"
	AS112TimeseriesParamsQueryTypeNid        AS112TimeseriesParamsQueryType = "NID"
	AS112TimeseriesParamsQueryTypeNimloc     AS112TimeseriesParamsQueryType = "NIMLOC"
	AS112TimeseriesParamsQueryTypeNinfo      AS112TimeseriesParamsQueryType = "NINFO"
	AS112TimeseriesParamsQueryTypeNS         AS112TimeseriesParamsQueryType = "NS"
	AS112TimeseriesParamsQueryTypeNsap       AS112TimeseriesParamsQueryType = "NSAP"
	AS112TimeseriesParamsQueryTypeNsec       AS112TimeseriesParamsQueryType = "NSEC"
	AS112TimeseriesParamsQueryTypeNsec3      AS112TimeseriesParamsQueryType = "NSEC3"
	AS112TimeseriesParamsQueryTypeNsec3Param AS112TimeseriesParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesParamsQueryTypeNull       AS112TimeseriesParamsQueryType = "NULL"
	AS112TimeseriesParamsQueryTypeNxt        AS112TimeseriesParamsQueryType = "NXT"
	AS112TimeseriesParamsQueryTypeOpenpgpkey AS112TimeseriesParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesParamsQueryTypeOpt        AS112TimeseriesParamsQueryType = "OPT"
	AS112TimeseriesParamsQueryTypePTR        AS112TimeseriesParamsQueryType = "PTR"
	AS112TimeseriesParamsQueryTypePx         AS112TimeseriesParamsQueryType = "PX"
	AS112TimeseriesParamsQueryTypeRkey       AS112TimeseriesParamsQueryType = "RKEY"
	AS112TimeseriesParamsQueryTypeRp         AS112TimeseriesParamsQueryType = "RP"
	AS112TimeseriesParamsQueryTypeRrsig      AS112TimeseriesParamsQueryType = "RRSIG"
	AS112TimeseriesParamsQueryTypeRt         AS112TimeseriesParamsQueryType = "RT"
	AS112TimeseriesParamsQueryTypeSig        AS112TimeseriesParamsQueryType = "SIG"
	AS112TimeseriesParamsQueryTypeSink       AS112TimeseriesParamsQueryType = "SINK"
	AS112TimeseriesParamsQueryTypeSMIMEA     AS112TimeseriesParamsQueryType = "SMIMEA"
	AS112TimeseriesParamsQueryTypeSOA        AS112TimeseriesParamsQueryType = "SOA"
	AS112TimeseriesParamsQueryTypeSPF        AS112TimeseriesParamsQueryType = "SPF"
	AS112TimeseriesParamsQueryTypeSRV        AS112TimeseriesParamsQueryType = "SRV"
	AS112TimeseriesParamsQueryTypeSSHFP      AS112TimeseriesParamsQueryType = "SSHFP"
	AS112TimeseriesParamsQueryTypeSVCB       AS112TimeseriesParamsQueryType = "SVCB"
	AS112TimeseriesParamsQueryTypeTa         AS112TimeseriesParamsQueryType = "TA"
	AS112TimeseriesParamsQueryTypeTalink     AS112TimeseriesParamsQueryType = "TALINK"
	AS112TimeseriesParamsQueryTypeTkey       AS112TimeseriesParamsQueryType = "TKEY"
	AS112TimeseriesParamsQueryTypeTLSA       AS112TimeseriesParamsQueryType = "TLSA"
	AS112TimeseriesParamsQueryTypeTSIG       AS112TimeseriesParamsQueryType = "TSIG"
	AS112TimeseriesParamsQueryTypeTXT        AS112TimeseriesParamsQueryType = "TXT"
	AS112TimeseriesParamsQueryTypeUinfo      AS112TimeseriesParamsQueryType = "UINFO"
	AS112TimeseriesParamsQueryTypeUID        AS112TimeseriesParamsQueryType = "UID"
	AS112TimeseriesParamsQueryTypeUnspec     AS112TimeseriesParamsQueryType = "UNSPEC"
	AS112TimeseriesParamsQueryTypeURI        AS112TimeseriesParamsQueryType = "URI"
	AS112TimeseriesParamsQueryTypeWks        AS112TimeseriesParamsQueryType = "WKS"
	AS112TimeseriesParamsQueryTypeX25        AS112TimeseriesParamsQueryType = "X25"
	AS112TimeseriesParamsQueryTypeZonemd     AS112TimeseriesParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsQueryTypeA, AS112TimeseriesParamsQueryTypeAAAA, AS112TimeseriesParamsQueryTypeA6, AS112TimeseriesParamsQueryTypeAfsdb, AS112TimeseriesParamsQueryTypeAny, AS112TimeseriesParamsQueryTypeApl, AS112TimeseriesParamsQueryTypeAtma, AS112TimeseriesParamsQueryTypeAXFR, AS112TimeseriesParamsQueryTypeCAA, AS112TimeseriesParamsQueryTypeCdnskey, AS112TimeseriesParamsQueryTypeCds, AS112TimeseriesParamsQueryTypeCERT, AS112TimeseriesParamsQueryTypeCNAME, AS112TimeseriesParamsQueryTypeCsync, AS112TimeseriesParamsQueryTypeDhcid, AS112TimeseriesParamsQueryTypeDlv, AS112TimeseriesParamsQueryTypeDname, AS112TimeseriesParamsQueryTypeDNSKEY, AS112TimeseriesParamsQueryTypeDoa, AS112TimeseriesParamsQueryTypeDS, AS112TimeseriesParamsQueryTypeEid, AS112TimeseriesParamsQueryTypeEui48, AS112TimeseriesParamsQueryTypeEui64, AS112TimeseriesParamsQueryTypeGpos, AS112TimeseriesParamsQueryTypeGid, AS112TimeseriesParamsQueryTypeHinfo, AS112TimeseriesParamsQueryTypeHip, AS112TimeseriesParamsQueryTypeHTTPS, AS112TimeseriesParamsQueryTypeIpseckey, AS112TimeseriesParamsQueryTypeIsdn, AS112TimeseriesParamsQueryTypeIxfr, AS112TimeseriesParamsQueryTypeKey, AS112TimeseriesParamsQueryTypeKx, AS112TimeseriesParamsQueryTypeL32, AS112TimeseriesParamsQueryTypeL64, AS112TimeseriesParamsQueryTypeLOC, AS112TimeseriesParamsQueryTypeLp, AS112TimeseriesParamsQueryTypeMaila, AS112TimeseriesParamsQueryTypeMailb, AS112TimeseriesParamsQueryTypeMB, AS112TimeseriesParamsQueryTypeMd, AS112TimeseriesParamsQueryTypeMf, AS112TimeseriesParamsQueryTypeMg, AS112TimeseriesParamsQueryTypeMinfo, AS112TimeseriesParamsQueryTypeMr, AS112TimeseriesParamsQueryTypeMX, AS112TimeseriesParamsQueryTypeNAPTR, AS112TimeseriesParamsQueryTypeNb, AS112TimeseriesParamsQueryTypeNbstat, AS112TimeseriesParamsQueryTypeNid, AS112TimeseriesParamsQueryTypeNimloc, AS112TimeseriesParamsQueryTypeNinfo, AS112TimeseriesParamsQueryTypeNS, AS112TimeseriesParamsQueryTypeNsap, AS112TimeseriesParamsQueryTypeNsec, AS112TimeseriesParamsQueryTypeNsec3, AS112TimeseriesParamsQueryTypeNsec3Param, AS112TimeseriesParamsQueryTypeNull, AS112TimeseriesParamsQueryTypeNxt, AS112TimeseriesParamsQueryTypeOpenpgpkey, AS112TimeseriesParamsQueryTypeOpt, AS112TimeseriesParamsQueryTypePTR, AS112TimeseriesParamsQueryTypePx, AS112TimeseriesParamsQueryTypeRkey, AS112TimeseriesParamsQueryTypeRp, AS112TimeseriesParamsQueryTypeRrsig, AS112TimeseriesParamsQueryTypeRt, AS112TimeseriesParamsQueryTypeSig, AS112TimeseriesParamsQueryTypeSink, AS112TimeseriesParamsQueryTypeSMIMEA, AS112TimeseriesParamsQueryTypeSOA, AS112TimeseriesParamsQueryTypeSPF, AS112TimeseriesParamsQueryTypeSRV, AS112TimeseriesParamsQueryTypeSSHFP, AS112TimeseriesParamsQueryTypeSVCB, AS112TimeseriesParamsQueryTypeTa, AS112TimeseriesParamsQueryTypeTalink, AS112TimeseriesParamsQueryTypeTkey, AS112TimeseriesParamsQueryTypeTLSA, AS112TimeseriesParamsQueryTypeTSIG, AS112TimeseriesParamsQueryTypeTXT, AS112TimeseriesParamsQueryTypeUinfo, AS112TimeseriesParamsQueryTypeUID, AS112TimeseriesParamsQueryTypeUnspec, AS112TimeseriesParamsQueryTypeURI, AS112TimeseriesParamsQueryTypeWks, AS112TimeseriesParamsQueryTypeX25, AS112TimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112TimeseriesParamsResponseCode string

const (
	AS112TimeseriesParamsResponseCodeNoerror   AS112TimeseriesParamsResponseCode = "NOERROR"
	AS112TimeseriesParamsResponseCodeFormerr   AS112TimeseriesParamsResponseCode = "FORMERR"
	AS112TimeseriesParamsResponseCodeServfail  AS112TimeseriesParamsResponseCode = "SERVFAIL"
	AS112TimeseriesParamsResponseCodeNxdomain  AS112TimeseriesParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesParamsResponseCodeNotimp    AS112TimeseriesParamsResponseCode = "NOTIMP"
	AS112TimeseriesParamsResponseCodeRefused   AS112TimeseriesParamsResponseCode = "REFUSED"
	AS112TimeseriesParamsResponseCodeYxdomain  AS112TimeseriesParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesParamsResponseCodeYxrrset   AS112TimeseriesParamsResponseCode = "YXRRSET"
	AS112TimeseriesParamsResponseCodeNxrrset   AS112TimeseriesParamsResponseCode = "NXRRSET"
	AS112TimeseriesParamsResponseCodeNotauth   AS112TimeseriesParamsResponseCode = "NOTAUTH"
	AS112TimeseriesParamsResponseCodeNotzone   AS112TimeseriesParamsResponseCode = "NOTZONE"
	AS112TimeseriesParamsResponseCodeBadsig    AS112TimeseriesParamsResponseCode = "BADSIG"
	AS112TimeseriesParamsResponseCodeBadkey    AS112TimeseriesParamsResponseCode = "BADKEY"
	AS112TimeseriesParamsResponseCodeBadtime   AS112TimeseriesParamsResponseCode = "BADTIME"
	AS112TimeseriesParamsResponseCodeBadmode   AS112TimeseriesParamsResponseCode = "BADMODE"
	AS112TimeseriesParamsResponseCodeBadname   AS112TimeseriesParamsResponseCode = "BADNAME"
	AS112TimeseriesParamsResponseCodeBadalg    AS112TimeseriesParamsResponseCode = "BADALG"
	AS112TimeseriesParamsResponseCodeBadtrunc  AS112TimeseriesParamsResponseCode = "BADTRUNC"
	AS112TimeseriesParamsResponseCodeBadcookie AS112TimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsResponseCodeNoerror, AS112TimeseriesParamsResponseCodeFormerr, AS112TimeseriesParamsResponseCodeServfail, AS112TimeseriesParamsResponseCodeNxdomain, AS112TimeseriesParamsResponseCodeNotimp, AS112TimeseriesParamsResponseCodeRefused, AS112TimeseriesParamsResponseCodeYxdomain, AS112TimeseriesParamsResponseCodeYxrrset, AS112TimeseriesParamsResponseCodeNxrrset, AS112TimeseriesParamsResponseCodeNotauth, AS112TimeseriesParamsResponseCodeNotzone, AS112TimeseriesParamsResponseCodeBadsig, AS112TimeseriesParamsResponseCodeBadkey, AS112TimeseriesParamsResponseCodeBadtime, AS112TimeseriesParamsResponseCodeBadmode, AS112TimeseriesParamsResponseCodeBadname, AS112TimeseriesParamsResponseCodeBadalg, AS112TimeseriesParamsResponseCodeBadtrunc, AS112TimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesResponseEnvelope struct {
	Result  AS112TimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    as112TimeseriesResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseEnvelope]
type as112TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
