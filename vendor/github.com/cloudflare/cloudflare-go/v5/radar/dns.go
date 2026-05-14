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

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options          []option.RequestOption
	Top              *DNSTopService
	Summary          *DNSSummaryService
	TimeseriesGroups *DNSTimeseriesGroupService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Top = NewDNSTopService(opts...)
	r.Summary = NewDNSSummaryService(opts...)
	r.TimeseriesGroups = NewDNSTimeseriesGroupService(opts...)
	return
}

// Retrieves normalized query volume to the 1.1.1.1 DNS resolver over time.
func (r *DNSService) Timeseries(ctx context.Context, query DNSTimeseriesParams, opts ...option.RequestOption) (res *DNSTimeseriesResponse, err error) {
	var env DNSTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSTimeseriesResponse struct {
	// Metadata for the results.
	Meta        DNSTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]DNSTimeseriesResponse `json:"-,extras"`
	JSON        dnsTimeseriesResponseJSON        `json:"-"`
}

// dnsTimeseriesResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponse]
type dnsTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    DNSTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo DNSTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  dnsTimeseriesResponseMetaJSON   `json:"-"`
}

// dnsTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMeta]
type dnsTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesResponseMetaAggInterval string

const (
	DNSTimeseriesResponseMetaAggIntervalFifteenMinutes DNSTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	DNSTimeseriesResponseMetaAggIntervalOneHour        DNSTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	DNSTimeseriesResponseMetaAggIntervalOneDay         DNSTimeseriesResponseMetaAggInterval = "ONE_DAY"
	DNSTimeseriesResponseMetaAggIntervalOneWeek        DNSTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	DNSTimeseriesResponseMetaAggIntervalOneMonth       DNSTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r DNSTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaAggIntervalFifteenMinutes, DNSTimeseriesResponseMetaAggIntervalOneHour, DNSTimeseriesResponseMetaAggIntervalOneDay, DNSTimeseriesResponseMetaAggIntervalOneWeek, DNSTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []DNSTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level,required"`
	JSON  dnsTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTimeseriesResponseMetaConfidenceInfo]
type dnsTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate,required" format:"date-time"`
	JSON            dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTimeseriesResponseMetaConfidenceInfoAnnotation]
type dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      dnsTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaDateRange]
type dnsTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTimeseriesResponseMetaNormalization string

const (
	DNSTimeseriesResponseMetaNormalizationPercentage           DNSTimeseriesResponseMetaNormalization = "PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationMin0Max              DNSTimeseriesResponseMetaNormalization = "MIN0_MAX"
	DNSTimeseriesResponseMetaNormalizationMinMax               DNSTimeseriesResponseMetaNormalization = "MIN_MAX"
	DNSTimeseriesResponseMetaNormalizationRawValues            DNSTimeseriesResponseMetaNormalization = "RAW_VALUES"
	DNSTimeseriesResponseMetaNormalizationPercentageChange     DNSTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTimeseriesResponseMetaNormalizationRollingAverage       DNSTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTimeseriesResponseMetaNormalizationOverlappedPercentage DNSTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationRatio                DNSTimeseriesResponseMetaNormalization = "RATIO"
)

func (r DNSTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaNormalizationPercentage, DNSTimeseriesResponseMetaNormalizationMin0Max, DNSTimeseriesResponseMetaNormalizationMinMax, DNSTimeseriesResponseMetaNormalizationRawValues, DNSTimeseriesResponseMetaNormalizationPercentageChange, DNSTimeseriesResponseMetaNormalizationRollingAverage, DNSTimeseriesResponseMetaNormalizationOverlappedPercentage, DNSTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaUnit struct {
	Name  string                            `json:"name,required"`
	Value string                            `json:"value,required"`
	JSON  dnsTimeseriesResponseMetaUnitJSON `json:"-"`
}

// dnsTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaUnit]
type dnsTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
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
	Format param.Field[DNSTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesParams]'s query parameters as `url.Values`.
func (r DNSTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesParamsAggInterval string

const (
	DNSTimeseriesParamsAggInterval15m DNSTimeseriesParamsAggInterval = "15m"
	DNSTimeseriesParamsAggInterval1h  DNSTimeseriesParamsAggInterval = "1h"
	DNSTimeseriesParamsAggInterval1d  DNSTimeseriesParamsAggInterval = "1d"
	DNSTimeseriesParamsAggInterval1w  DNSTimeseriesParamsAggInterval = "1w"
)

func (r DNSTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsAggInterval15m, DNSTimeseriesParamsAggInterval1h, DNSTimeseriesParamsAggInterval1d, DNSTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesParamsFormat string

const (
	DNSTimeseriesParamsFormatJson DNSTimeseriesParamsFormat = "JSON"
	DNSTimeseriesParamsFormatCsv  DNSTimeseriesParamsFormat = "CSV"
)

func (r DNSTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsFormatJson, DNSTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesParamsProtocol string

const (
	DNSTimeseriesParamsProtocolUdp   DNSTimeseriesParamsProtocol = "UDP"
	DNSTimeseriesParamsProtocolTCP   DNSTimeseriesParamsProtocol = "TCP"
	DNSTimeseriesParamsProtocolHTTPS DNSTimeseriesParamsProtocol = "HTTPS"
	DNSTimeseriesParamsProtocolTLS   DNSTimeseriesParamsProtocol = "TLS"
)

func (r DNSTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsProtocolUdp, DNSTimeseriesParamsProtocolTCP, DNSTimeseriesParamsProtocolHTTPS, DNSTimeseriesParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesParamsQueryType string

const (
	DNSTimeseriesParamsQueryTypeA          DNSTimeseriesParamsQueryType = "A"
	DNSTimeseriesParamsQueryTypeAAAA       DNSTimeseriesParamsQueryType = "AAAA"
	DNSTimeseriesParamsQueryTypeA6         DNSTimeseriesParamsQueryType = "A6"
	DNSTimeseriesParamsQueryTypeAfsdb      DNSTimeseriesParamsQueryType = "AFSDB"
	DNSTimeseriesParamsQueryTypeAny        DNSTimeseriesParamsQueryType = "ANY"
	DNSTimeseriesParamsQueryTypeApl        DNSTimeseriesParamsQueryType = "APL"
	DNSTimeseriesParamsQueryTypeAtma       DNSTimeseriesParamsQueryType = "ATMA"
	DNSTimeseriesParamsQueryTypeAXFR       DNSTimeseriesParamsQueryType = "AXFR"
	DNSTimeseriesParamsQueryTypeCAA        DNSTimeseriesParamsQueryType = "CAA"
	DNSTimeseriesParamsQueryTypeCdnskey    DNSTimeseriesParamsQueryType = "CDNSKEY"
	DNSTimeseriesParamsQueryTypeCds        DNSTimeseriesParamsQueryType = "CDS"
	DNSTimeseriesParamsQueryTypeCERT       DNSTimeseriesParamsQueryType = "CERT"
	DNSTimeseriesParamsQueryTypeCNAME      DNSTimeseriesParamsQueryType = "CNAME"
	DNSTimeseriesParamsQueryTypeCsync      DNSTimeseriesParamsQueryType = "CSYNC"
	DNSTimeseriesParamsQueryTypeDhcid      DNSTimeseriesParamsQueryType = "DHCID"
	DNSTimeseriesParamsQueryTypeDlv        DNSTimeseriesParamsQueryType = "DLV"
	DNSTimeseriesParamsQueryTypeDname      DNSTimeseriesParamsQueryType = "DNAME"
	DNSTimeseriesParamsQueryTypeDNSKEY     DNSTimeseriesParamsQueryType = "DNSKEY"
	DNSTimeseriesParamsQueryTypeDoa        DNSTimeseriesParamsQueryType = "DOA"
	DNSTimeseriesParamsQueryTypeDS         DNSTimeseriesParamsQueryType = "DS"
	DNSTimeseriesParamsQueryTypeEid        DNSTimeseriesParamsQueryType = "EID"
	DNSTimeseriesParamsQueryTypeEui48      DNSTimeseriesParamsQueryType = "EUI48"
	DNSTimeseriesParamsQueryTypeEui64      DNSTimeseriesParamsQueryType = "EUI64"
	DNSTimeseriesParamsQueryTypeGpos       DNSTimeseriesParamsQueryType = "GPOS"
	DNSTimeseriesParamsQueryTypeGid        DNSTimeseriesParamsQueryType = "GID"
	DNSTimeseriesParamsQueryTypeHinfo      DNSTimeseriesParamsQueryType = "HINFO"
	DNSTimeseriesParamsQueryTypeHip        DNSTimeseriesParamsQueryType = "HIP"
	DNSTimeseriesParamsQueryTypeHTTPS      DNSTimeseriesParamsQueryType = "HTTPS"
	DNSTimeseriesParamsQueryTypeIpseckey   DNSTimeseriesParamsQueryType = "IPSECKEY"
	DNSTimeseriesParamsQueryTypeIsdn       DNSTimeseriesParamsQueryType = "ISDN"
	DNSTimeseriesParamsQueryTypeIxfr       DNSTimeseriesParamsQueryType = "IXFR"
	DNSTimeseriesParamsQueryTypeKey        DNSTimeseriesParamsQueryType = "KEY"
	DNSTimeseriesParamsQueryTypeKx         DNSTimeseriesParamsQueryType = "KX"
	DNSTimeseriesParamsQueryTypeL32        DNSTimeseriesParamsQueryType = "L32"
	DNSTimeseriesParamsQueryTypeL64        DNSTimeseriesParamsQueryType = "L64"
	DNSTimeseriesParamsQueryTypeLOC        DNSTimeseriesParamsQueryType = "LOC"
	DNSTimeseriesParamsQueryTypeLp         DNSTimeseriesParamsQueryType = "LP"
	DNSTimeseriesParamsQueryTypeMaila      DNSTimeseriesParamsQueryType = "MAILA"
	DNSTimeseriesParamsQueryTypeMailb      DNSTimeseriesParamsQueryType = "MAILB"
	DNSTimeseriesParamsQueryTypeMB         DNSTimeseriesParamsQueryType = "MB"
	DNSTimeseriesParamsQueryTypeMd         DNSTimeseriesParamsQueryType = "MD"
	DNSTimeseriesParamsQueryTypeMf         DNSTimeseriesParamsQueryType = "MF"
	DNSTimeseriesParamsQueryTypeMg         DNSTimeseriesParamsQueryType = "MG"
	DNSTimeseriesParamsQueryTypeMinfo      DNSTimeseriesParamsQueryType = "MINFO"
	DNSTimeseriesParamsQueryTypeMr         DNSTimeseriesParamsQueryType = "MR"
	DNSTimeseriesParamsQueryTypeMX         DNSTimeseriesParamsQueryType = "MX"
	DNSTimeseriesParamsQueryTypeNAPTR      DNSTimeseriesParamsQueryType = "NAPTR"
	DNSTimeseriesParamsQueryTypeNb         DNSTimeseriesParamsQueryType = "NB"
	DNSTimeseriesParamsQueryTypeNbstat     DNSTimeseriesParamsQueryType = "NBSTAT"
	DNSTimeseriesParamsQueryTypeNid        DNSTimeseriesParamsQueryType = "NID"
	DNSTimeseriesParamsQueryTypeNimloc     DNSTimeseriesParamsQueryType = "NIMLOC"
	DNSTimeseriesParamsQueryTypeNinfo      DNSTimeseriesParamsQueryType = "NINFO"
	DNSTimeseriesParamsQueryTypeNS         DNSTimeseriesParamsQueryType = "NS"
	DNSTimeseriesParamsQueryTypeNsap       DNSTimeseriesParamsQueryType = "NSAP"
	DNSTimeseriesParamsQueryTypeNsec       DNSTimeseriesParamsQueryType = "NSEC"
	DNSTimeseriesParamsQueryTypeNsec3      DNSTimeseriesParamsQueryType = "NSEC3"
	DNSTimeseriesParamsQueryTypeNsec3Param DNSTimeseriesParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesParamsQueryTypeNull       DNSTimeseriesParamsQueryType = "NULL"
	DNSTimeseriesParamsQueryTypeNxt        DNSTimeseriesParamsQueryType = "NXT"
	DNSTimeseriesParamsQueryTypeOpenpgpkey DNSTimeseriesParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesParamsQueryTypeOpt        DNSTimeseriesParamsQueryType = "OPT"
	DNSTimeseriesParamsQueryTypePTR        DNSTimeseriesParamsQueryType = "PTR"
	DNSTimeseriesParamsQueryTypePx         DNSTimeseriesParamsQueryType = "PX"
	DNSTimeseriesParamsQueryTypeRkey       DNSTimeseriesParamsQueryType = "RKEY"
	DNSTimeseriesParamsQueryTypeRp         DNSTimeseriesParamsQueryType = "RP"
	DNSTimeseriesParamsQueryTypeRrsig      DNSTimeseriesParamsQueryType = "RRSIG"
	DNSTimeseriesParamsQueryTypeRt         DNSTimeseriesParamsQueryType = "RT"
	DNSTimeseriesParamsQueryTypeSig        DNSTimeseriesParamsQueryType = "SIG"
	DNSTimeseriesParamsQueryTypeSink       DNSTimeseriesParamsQueryType = "SINK"
	DNSTimeseriesParamsQueryTypeSMIMEA     DNSTimeseriesParamsQueryType = "SMIMEA"
	DNSTimeseriesParamsQueryTypeSOA        DNSTimeseriesParamsQueryType = "SOA"
	DNSTimeseriesParamsQueryTypeSPF        DNSTimeseriesParamsQueryType = "SPF"
	DNSTimeseriesParamsQueryTypeSRV        DNSTimeseriesParamsQueryType = "SRV"
	DNSTimeseriesParamsQueryTypeSSHFP      DNSTimeseriesParamsQueryType = "SSHFP"
	DNSTimeseriesParamsQueryTypeSVCB       DNSTimeseriesParamsQueryType = "SVCB"
	DNSTimeseriesParamsQueryTypeTa         DNSTimeseriesParamsQueryType = "TA"
	DNSTimeseriesParamsQueryTypeTalink     DNSTimeseriesParamsQueryType = "TALINK"
	DNSTimeseriesParamsQueryTypeTkey       DNSTimeseriesParamsQueryType = "TKEY"
	DNSTimeseriesParamsQueryTypeTLSA       DNSTimeseriesParamsQueryType = "TLSA"
	DNSTimeseriesParamsQueryTypeTSIG       DNSTimeseriesParamsQueryType = "TSIG"
	DNSTimeseriesParamsQueryTypeTXT        DNSTimeseriesParamsQueryType = "TXT"
	DNSTimeseriesParamsQueryTypeUinfo      DNSTimeseriesParamsQueryType = "UINFO"
	DNSTimeseriesParamsQueryTypeUID        DNSTimeseriesParamsQueryType = "UID"
	DNSTimeseriesParamsQueryTypeUnspec     DNSTimeseriesParamsQueryType = "UNSPEC"
	DNSTimeseriesParamsQueryTypeURI        DNSTimeseriesParamsQueryType = "URI"
	DNSTimeseriesParamsQueryTypeWks        DNSTimeseriesParamsQueryType = "WKS"
	DNSTimeseriesParamsQueryTypeX25        DNSTimeseriesParamsQueryType = "X25"
	DNSTimeseriesParamsQueryTypeZonemd     DNSTimeseriesParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsQueryTypeA, DNSTimeseriesParamsQueryTypeAAAA, DNSTimeseriesParamsQueryTypeA6, DNSTimeseriesParamsQueryTypeAfsdb, DNSTimeseriesParamsQueryTypeAny, DNSTimeseriesParamsQueryTypeApl, DNSTimeseriesParamsQueryTypeAtma, DNSTimeseriesParamsQueryTypeAXFR, DNSTimeseriesParamsQueryTypeCAA, DNSTimeseriesParamsQueryTypeCdnskey, DNSTimeseriesParamsQueryTypeCds, DNSTimeseriesParamsQueryTypeCERT, DNSTimeseriesParamsQueryTypeCNAME, DNSTimeseriesParamsQueryTypeCsync, DNSTimeseriesParamsQueryTypeDhcid, DNSTimeseriesParamsQueryTypeDlv, DNSTimeseriesParamsQueryTypeDname, DNSTimeseriesParamsQueryTypeDNSKEY, DNSTimeseriesParamsQueryTypeDoa, DNSTimeseriesParamsQueryTypeDS, DNSTimeseriesParamsQueryTypeEid, DNSTimeseriesParamsQueryTypeEui48, DNSTimeseriesParamsQueryTypeEui64, DNSTimeseriesParamsQueryTypeGpos, DNSTimeseriesParamsQueryTypeGid, DNSTimeseriesParamsQueryTypeHinfo, DNSTimeseriesParamsQueryTypeHip, DNSTimeseriesParamsQueryTypeHTTPS, DNSTimeseriesParamsQueryTypeIpseckey, DNSTimeseriesParamsQueryTypeIsdn, DNSTimeseriesParamsQueryTypeIxfr, DNSTimeseriesParamsQueryTypeKey, DNSTimeseriesParamsQueryTypeKx, DNSTimeseriesParamsQueryTypeL32, DNSTimeseriesParamsQueryTypeL64, DNSTimeseriesParamsQueryTypeLOC, DNSTimeseriesParamsQueryTypeLp, DNSTimeseriesParamsQueryTypeMaila, DNSTimeseriesParamsQueryTypeMailb, DNSTimeseriesParamsQueryTypeMB, DNSTimeseriesParamsQueryTypeMd, DNSTimeseriesParamsQueryTypeMf, DNSTimeseriesParamsQueryTypeMg, DNSTimeseriesParamsQueryTypeMinfo, DNSTimeseriesParamsQueryTypeMr, DNSTimeseriesParamsQueryTypeMX, DNSTimeseriesParamsQueryTypeNAPTR, DNSTimeseriesParamsQueryTypeNb, DNSTimeseriesParamsQueryTypeNbstat, DNSTimeseriesParamsQueryTypeNid, DNSTimeseriesParamsQueryTypeNimloc, DNSTimeseriesParamsQueryTypeNinfo, DNSTimeseriesParamsQueryTypeNS, DNSTimeseriesParamsQueryTypeNsap, DNSTimeseriesParamsQueryTypeNsec, DNSTimeseriesParamsQueryTypeNsec3, DNSTimeseriesParamsQueryTypeNsec3Param, DNSTimeseriesParamsQueryTypeNull, DNSTimeseriesParamsQueryTypeNxt, DNSTimeseriesParamsQueryTypeOpenpgpkey, DNSTimeseriesParamsQueryTypeOpt, DNSTimeseriesParamsQueryTypePTR, DNSTimeseriesParamsQueryTypePx, DNSTimeseriesParamsQueryTypeRkey, DNSTimeseriesParamsQueryTypeRp, DNSTimeseriesParamsQueryTypeRrsig, DNSTimeseriesParamsQueryTypeRt, DNSTimeseriesParamsQueryTypeSig, DNSTimeseriesParamsQueryTypeSink, DNSTimeseriesParamsQueryTypeSMIMEA, DNSTimeseriesParamsQueryTypeSOA, DNSTimeseriesParamsQueryTypeSPF, DNSTimeseriesParamsQueryTypeSRV, DNSTimeseriesParamsQueryTypeSSHFP, DNSTimeseriesParamsQueryTypeSVCB, DNSTimeseriesParamsQueryTypeTa, DNSTimeseriesParamsQueryTypeTalink, DNSTimeseriesParamsQueryTypeTkey, DNSTimeseriesParamsQueryTypeTLSA, DNSTimeseriesParamsQueryTypeTSIG, DNSTimeseriesParamsQueryTypeTXT, DNSTimeseriesParamsQueryTypeUinfo, DNSTimeseriesParamsQueryTypeUID, DNSTimeseriesParamsQueryTypeUnspec, DNSTimeseriesParamsQueryTypeURI, DNSTimeseriesParamsQueryTypeWks, DNSTimeseriesParamsQueryTypeX25, DNSTimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesParamsResponseCode string

const (
	DNSTimeseriesParamsResponseCodeNoerror   DNSTimeseriesParamsResponseCode = "NOERROR"
	DNSTimeseriesParamsResponseCodeFormerr   DNSTimeseriesParamsResponseCode = "FORMERR"
	DNSTimeseriesParamsResponseCodeServfail  DNSTimeseriesParamsResponseCode = "SERVFAIL"
	DNSTimeseriesParamsResponseCodeNxdomain  DNSTimeseriesParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesParamsResponseCodeNotimp    DNSTimeseriesParamsResponseCode = "NOTIMP"
	DNSTimeseriesParamsResponseCodeRefused   DNSTimeseriesParamsResponseCode = "REFUSED"
	DNSTimeseriesParamsResponseCodeYxdomain  DNSTimeseriesParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesParamsResponseCodeYxrrset   DNSTimeseriesParamsResponseCode = "YXRRSET"
	DNSTimeseriesParamsResponseCodeNxrrset   DNSTimeseriesParamsResponseCode = "NXRRSET"
	DNSTimeseriesParamsResponseCodeNotauth   DNSTimeseriesParamsResponseCode = "NOTAUTH"
	DNSTimeseriesParamsResponseCodeNotzone   DNSTimeseriesParamsResponseCode = "NOTZONE"
	DNSTimeseriesParamsResponseCodeBadsig    DNSTimeseriesParamsResponseCode = "BADSIG"
	DNSTimeseriesParamsResponseCodeBadkey    DNSTimeseriesParamsResponseCode = "BADKEY"
	DNSTimeseriesParamsResponseCodeBadtime   DNSTimeseriesParamsResponseCode = "BADTIME"
	DNSTimeseriesParamsResponseCodeBadmode   DNSTimeseriesParamsResponseCode = "BADMODE"
	DNSTimeseriesParamsResponseCodeBadname   DNSTimeseriesParamsResponseCode = "BADNAME"
	DNSTimeseriesParamsResponseCodeBadalg    DNSTimeseriesParamsResponseCode = "BADALG"
	DNSTimeseriesParamsResponseCodeBadtrunc  DNSTimeseriesParamsResponseCode = "BADTRUNC"
	DNSTimeseriesParamsResponseCodeBadcookie DNSTimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsResponseCodeNoerror, DNSTimeseriesParamsResponseCodeFormerr, DNSTimeseriesParamsResponseCodeServfail, DNSTimeseriesParamsResponseCodeNxdomain, DNSTimeseriesParamsResponseCodeNotimp, DNSTimeseriesParamsResponseCodeRefused, DNSTimeseriesParamsResponseCodeYxdomain, DNSTimeseriesParamsResponseCodeYxrrset, DNSTimeseriesParamsResponseCodeNxrrset, DNSTimeseriesParamsResponseCodeNotauth, DNSTimeseriesParamsResponseCodeNotzone, DNSTimeseriesParamsResponseCodeBadsig, DNSTimeseriesParamsResponseCodeBadkey, DNSTimeseriesParamsResponseCodeBadtime, DNSTimeseriesParamsResponseCodeBadmode, DNSTimeseriesParamsResponseCodeBadname, DNSTimeseriesParamsResponseCodeBadalg, DNSTimeseriesParamsResponseCodeBadtrunc, DNSTimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesResponseEnvelope struct {
	Result  DNSTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    dnsTimeseriesResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseEnvelope]
type dnsTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
