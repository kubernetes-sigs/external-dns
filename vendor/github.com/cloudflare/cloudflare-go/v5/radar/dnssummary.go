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

// DNSSummaryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSSummaryService] method instead.
type DNSSummaryService struct {
	Options []option.RequestOption
}

// NewDNSSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSummaryService(opts ...option.RequestOption) (r *DNSSummaryService) {
	r = &DNSSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries by cache status.
func (r *DNSSummaryService) CacheHit(ctx context.Context, query DNSSummaryCacheHitParams, opts ...option.RequestOption) (res *DNSSummaryCacheHitResponse, err error) {
	var env DNSSummaryCacheHitResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/cache_hit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS responses by DNSSEC (DNS Security Extensions)
// support.
func (r *DNSSummaryService) DNSSEC(ctx context.Context, query DNSSummaryDNSSECParams, opts ...option.RequestOption) (res *DNSSummaryDNSSECResponse, err error) {
	var env DNSSummaryDNSSECResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by DNSSEC (DNS Security Extensions)
// client awareness.
func (r *DNSSummaryService) DNSSECAware(ctx context.Context, query DNSSummaryDNSSECAwareParams, opts ...option.RequestOption) (res *DNSSummaryDNSSECAwareResponse, err error) {
	var env DNSSummaryDNSSECAwareResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec_aware"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNSSEC-validated answers by end-to-end security
// status.
func (r *DNSSummaryService) DNSSECE2E(ctx context.Context, query DNSSummaryDNSSECE2EParams, opts ...option.RequestOption) (res *DNSSummaryDnssece2EResponse, err error) {
	var env DNSSummaryDnssece2EResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec_e2e"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by IP version.
func (r *DNSSummaryService) IPVersion(ctx context.Context, query DNSSummaryIPVersionParams, opts ...option.RequestOption) (res *DNSSummaryIPVersionResponse, err error) {
	var env DNSSummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by matching answers.
func (r *DNSSummaryService) MatchingAnswer(ctx context.Context, query DNSSummaryMatchingAnswerParams, opts ...option.RequestOption) (res *DNSSummaryMatchingAnswerResponse, err error) {
	var env DNSSummaryMatchingAnswerResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/matching_answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by DNS transport protocol.
func (r *DNSSummaryService) Protocol(ctx context.Context, query DNSSummaryProtocolParams, opts ...option.RequestOption) (res *DNSSummaryProtocolResponse, err error) {
	var env DNSSummaryProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by type.
func (r *DNSSummaryService) QueryType(ctx context.Context, query DNSSummaryQueryTypeParams, opts ...option.RequestOption) (res *DNSSummaryQueryTypeResponse, err error) {
	var env DNSSummaryQueryTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by response code.
func (r *DNSSummaryService) ResponseCode(ctx context.Context, query DNSSummaryResponseCodeParams, opts ...option.RequestOption) (res *DNSSummaryResponseCodeResponse, err error) {
	var env DNSSummaryResponseCodeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/response_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by minimum response TTL.
func (r *DNSSummaryService) ResponseTTL(ctx context.Context, query DNSSummaryResponseTTLParams, opts ...option.RequestOption) (res *DNSSummaryResponseTTLResponse, err error) {
	var env DNSSummaryResponseTTLResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/response_ttl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSSummaryCacheHitResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryCacheHitResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryCacheHitResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryCacheHitResponseJSON     `json:"-"`
}

// dnsSummaryCacheHitResponseJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponse]
type dnsSummaryCacheHitResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryCacheHitResponseMeta struct {
	ConfidenceInfo DNSSummaryCacheHitResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryCacheHitResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryCacheHitResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryCacheHitResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryCacheHitResponseMetaJSON   `json:"-"`
}

// dnsSummaryCacheHitResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseMeta]
type dnsSummaryCacheHitResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryCacheHitResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  dnsSummaryCacheHitResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryCacheHitResponseMetaConfidenceInfo]
type dnsSummaryCacheHitResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation]
type dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryCacheHitResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryCacheHitResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryCacheHitResponseMetaDateRange]
type dnsSummaryCacheHitResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryCacheHitResponseMetaNormalization string

const (
	DNSSummaryCacheHitResponseMetaNormalizationPercentage           DNSSummaryCacheHitResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryCacheHitResponseMetaNormalizationMin0Max              DNSSummaryCacheHitResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryCacheHitResponseMetaNormalizationMinMax               DNSSummaryCacheHitResponseMetaNormalization = "MIN_MAX"
	DNSSummaryCacheHitResponseMetaNormalizationRawValues            DNSSummaryCacheHitResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryCacheHitResponseMetaNormalizationPercentageChange     DNSSummaryCacheHitResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryCacheHitResponseMetaNormalizationRollingAverage       DNSSummaryCacheHitResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryCacheHitResponseMetaNormalizationOverlappedPercentage DNSSummaryCacheHitResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryCacheHitResponseMetaNormalizationRatio                DNSSummaryCacheHitResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryCacheHitResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryCacheHitResponseMetaNormalizationPercentage, DNSSummaryCacheHitResponseMetaNormalizationMin0Max, DNSSummaryCacheHitResponseMetaNormalizationMinMax, DNSSummaryCacheHitResponseMetaNormalizationRawValues, DNSSummaryCacheHitResponseMetaNormalizationPercentageChange, DNSSummaryCacheHitResponseMetaNormalizationRollingAverage, DNSSummaryCacheHitResponseMetaNormalizationOverlappedPercentage, DNSSummaryCacheHitResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryCacheHitResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  dnsSummaryCacheHitResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseMetaUnit]
type dnsSummaryCacheHitResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryCacheHitResponseSummary0 struct {
	// A numeric string.
	Negative string `json:"NEGATIVE,required"`
	// A numeric string.
	Positive string                                 `json:"POSITIVE,required"`
	JSON     dnsSummaryCacheHitResponseSummary0JSON `json:"-"`
}

// dnsSummaryCacheHitResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseSummary0]
type dnsSummaryCacheHitResponseSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryDNSSECResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryDNSSECResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryDNSSECResponseJSON     `json:"-"`
}

// dnsSummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponse]
type dnsSummaryDNSSECResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryDNSSECResponseMeta struct {
	ConfidenceInfo DNSSummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryDNSSECResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryDNSSECResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryDNSSECResponseMetaJSON   `json:"-"`
}

// dnsSummaryDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseMeta]
type dnsSummaryDNSSECResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                          `json:"level,required"`
	JSON  dnsSummaryDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryDNSSECResponseMetaConfidenceInfo]
type dnsSummaryDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                 `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECResponseMetaDateRange]
type dnsSummaryDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryDNSSECResponseMetaNormalization string

const (
	DNSSummaryDNSSECResponseMetaNormalizationPercentage           DNSSummaryDNSSECResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryDNSSECResponseMetaNormalizationMin0Max              DNSSummaryDNSSECResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryDNSSECResponseMetaNormalizationMinMax               DNSSummaryDNSSECResponseMetaNormalization = "MIN_MAX"
	DNSSummaryDNSSECResponseMetaNormalizationRawValues            DNSSummaryDNSSECResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryDNSSECResponseMetaNormalizationPercentageChange     DNSSummaryDNSSECResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryDNSSECResponseMetaNormalizationRollingAverage       DNSSummaryDNSSECResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryDNSSECResponseMetaNormalizationOverlappedPercentage DNSSummaryDNSSECResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryDNSSECResponseMetaNormalizationRatio                DNSSummaryDNSSECResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryDNSSECResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECResponseMetaNormalizationPercentage, DNSSummaryDNSSECResponseMetaNormalizationMin0Max, DNSSummaryDNSSECResponseMetaNormalizationMinMax, DNSSummaryDNSSECResponseMetaNormalizationRawValues, DNSSummaryDNSSECResponseMetaNormalizationPercentageChange, DNSSummaryDNSSECResponseMetaNormalizationRollingAverage, DNSSummaryDNSSECResponseMetaNormalizationOverlappedPercentage, DNSSummaryDNSSECResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryDNSSECResponseMetaUnit struct {
	Name  string                               `json:"name,required"`
	Value string                               `json:"value,required"`
	JSON  dnsSummaryDNSSECResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseMetaUnit]
type dnsSummaryDNSSECResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECResponseSummary0 struct {
	// A numeric string.
	Insecure string `json:"INSECURE,required"`
	// A numeric string.
	Invalid string `json:"INVALID,required"`
	// A numeric string.
	Other string `json:"OTHER,required"`
	// A numeric string.
	Secure string                               `json:"SECURE,required"`
	JSON   dnsSummaryDNSSECResponseSummary0JSON `json:"-"`
}

// dnsSummaryDNSSECResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseSummary0]
type dnsSummaryDNSSECResponseSummary0JSON struct {
	Insecure    apijson.Field
	Invalid     apijson.Field
	Other       apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECAwareResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryDNSSECAwareResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryDNSSECAwareResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryDNSSECAwareResponseJSON     `json:"-"`
}

// dnsSummaryDNSSECAwareResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECAwareResponse]
type dnsSummaryDNSSECAwareResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryDNSSECAwareResponseMeta struct {
	ConfidenceInfo DNSSummaryDNSSECAwareResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryDNSSECAwareResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryDNSSECAwareResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryDNSSECAwareResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryDNSSECAwareResponseMetaJSON   `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECAwareResponseMeta]
type dnsSummaryDNSSECAwareResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECAwareResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryDNSSECAwareResponseMetaConfidenceInfo]
type dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECAwareResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryDNSSECAwareResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryDNSSECAwareResponseMetaDateRange]
type dnsSummaryDNSSECAwareResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryDNSSECAwareResponseMetaNormalization string

const (
	DNSSummaryDNSSECAwareResponseMetaNormalizationPercentage           DNSSummaryDNSSECAwareResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryDNSSECAwareResponseMetaNormalizationMin0Max              DNSSummaryDNSSECAwareResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryDNSSECAwareResponseMetaNormalizationMinMax               DNSSummaryDNSSECAwareResponseMetaNormalization = "MIN_MAX"
	DNSSummaryDNSSECAwareResponseMetaNormalizationRawValues            DNSSummaryDNSSECAwareResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryDNSSECAwareResponseMetaNormalizationPercentageChange     DNSSummaryDNSSECAwareResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryDNSSECAwareResponseMetaNormalizationRollingAverage       DNSSummaryDNSSECAwareResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryDNSSECAwareResponseMetaNormalizationOverlappedPercentage DNSSummaryDNSSECAwareResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryDNSSECAwareResponseMetaNormalizationRatio                DNSSummaryDNSSECAwareResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryDNSSECAwareResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECAwareResponseMetaNormalizationPercentage, DNSSummaryDNSSECAwareResponseMetaNormalizationMin0Max, DNSSummaryDNSSECAwareResponseMetaNormalizationMinMax, DNSSummaryDNSSECAwareResponseMetaNormalizationRawValues, DNSSummaryDNSSECAwareResponseMetaNormalizationPercentageChange, DNSSummaryDNSSECAwareResponseMetaNormalizationRollingAverage, DNSSummaryDNSSECAwareResponseMetaNormalizationOverlappedPercentage, DNSSummaryDNSSECAwareResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryDNSSECAwareResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  dnsSummaryDNSSECAwareResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECAwareResponseMetaUnit]
type dnsSummaryDNSSECAwareResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECAwareResponseSummary0 struct {
	// A numeric string.
	NotSupported string `json:"NOT_SUPPORTED,required"`
	// A numeric string.
	Supported string                                    `json:"SUPPORTED,required"`
	JSON      dnsSummaryDNSSECAwareResponseSummary0JSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECAwareResponseSummary0]
type dnsSummaryDNSSECAwareResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDnssece2EResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryDnssece2EResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryDnssece2EResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryDnssece2EResponseJSON     `json:"-"`
}

// dnsSummaryDnssece2EResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDnssece2EResponse]
type dnsSummaryDnssece2EResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryDnssece2EResponseMeta struct {
	ConfidenceInfo DNSSummaryDnssece2EResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryDnssece2EResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryDnssece2EResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryDnssece2EResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryDnssece2EResponseMetaJSON   `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDnssece2EResponseMeta]
type dnsSummaryDnssece2EResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDnssece2EResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryDnssece2EResponseMetaConfidenceInfo]
type dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDnssece2EResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryDnssece2EResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseMetaDateRange]
type dnsSummaryDnssece2EResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryDnssece2EResponseMetaNormalization string

const (
	DNSSummaryDnssece2EResponseMetaNormalizationPercentage           DNSSummaryDnssece2EResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryDnssece2EResponseMetaNormalizationMin0Max              DNSSummaryDnssece2EResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryDnssece2EResponseMetaNormalizationMinMax               DNSSummaryDnssece2EResponseMetaNormalization = "MIN_MAX"
	DNSSummaryDnssece2EResponseMetaNormalizationRawValues            DNSSummaryDnssece2EResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryDnssece2EResponseMetaNormalizationPercentageChange     DNSSummaryDnssece2EResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryDnssece2EResponseMetaNormalizationRollingAverage       DNSSummaryDnssece2EResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryDnssece2EResponseMetaNormalizationOverlappedPercentage DNSSummaryDnssece2EResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryDnssece2EResponseMetaNormalizationRatio                DNSSummaryDnssece2EResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryDnssece2EResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryDnssece2EResponseMetaNormalizationPercentage, DNSSummaryDnssece2EResponseMetaNormalizationMin0Max, DNSSummaryDnssece2EResponseMetaNormalizationMinMax, DNSSummaryDnssece2EResponseMetaNormalizationRawValues, DNSSummaryDnssece2EResponseMetaNormalizationPercentageChange, DNSSummaryDnssece2EResponseMetaNormalizationRollingAverage, DNSSummaryDnssece2EResponseMetaNormalizationOverlappedPercentage, DNSSummaryDnssece2EResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryDnssece2EResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  dnsSummaryDnssece2EResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseMetaUnit]
type dnsSummaryDnssece2EResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDnssece2EResponseSummary0 struct {
	// A numeric string.
	Negative string `json:"NEGATIVE,required"`
	// A numeric string.
	Positive string                                  `json:"POSITIVE,required"`
	JSON     dnsSummaryDnssece2EResponseSummary0JSON `json:"-"`
}

// dnsSummaryDnssece2EResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseSummary0]
type dnsSummaryDnssece2EResponseSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryIPVersionResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryIPVersionResponseJSON     `json:"-"`
}

// dnsSummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [DNSSummaryIPVersionResponse]
type dnsSummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryIPVersionResponseMeta struct {
	ConfidenceInfo DNSSummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryIPVersionResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryIPVersionResponseMetaJSON   `json:"-"`
}

// dnsSummaryIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryIPVersionResponseMeta]
type dnsSummaryIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  dnsSummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryIPVersionResponseMetaConfidenceInfo]
type dnsSummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseMetaDateRange]
type dnsSummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryIPVersionResponseMetaNormalization string

const (
	DNSSummaryIPVersionResponseMetaNormalizationPercentage           DNSSummaryIPVersionResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryIPVersionResponseMetaNormalizationMin0Max              DNSSummaryIPVersionResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryIPVersionResponseMetaNormalizationMinMax               DNSSummaryIPVersionResponseMetaNormalization = "MIN_MAX"
	DNSSummaryIPVersionResponseMetaNormalizationRawValues            DNSSummaryIPVersionResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryIPVersionResponseMetaNormalizationPercentageChange     DNSSummaryIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryIPVersionResponseMetaNormalizationRollingAverage       DNSSummaryIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryIPVersionResponseMetaNormalizationOverlappedPercentage DNSSummaryIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryIPVersionResponseMetaNormalizationRatio                DNSSummaryIPVersionResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryIPVersionResponseMetaNormalizationPercentage, DNSSummaryIPVersionResponseMetaNormalizationMin0Max, DNSSummaryIPVersionResponseMetaNormalizationMinMax, DNSSummaryIPVersionResponseMetaNormalizationRawValues, DNSSummaryIPVersionResponseMetaNormalizationPercentageChange, DNSSummaryIPVersionResponseMetaNormalizationRollingAverage, DNSSummaryIPVersionResponseMetaNormalizationOverlappedPercentage, DNSSummaryIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryIPVersionResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  dnsSummaryIPVersionResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseMetaUnit]
type dnsSummaryIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryIPVersionResponseSummary0 struct {
	// A numeric string.
	IPv4 string `json:"IPv4,required"`
	// A numeric string.
	IPv6 string                                  `json:"IPv6,required"`
	JSON dnsSummaryIPVersionResponseSummary0JSON `json:"-"`
}

// dnsSummaryIPVersionResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseSummary0]
type dnsSummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryMatchingAnswerResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryMatchingAnswerResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryMatchingAnswerResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryMatchingAnswerResponseJSON     `json:"-"`
}

// dnsSummaryMatchingAnswerResponseJSON contains the JSON metadata for the struct
// [DNSSummaryMatchingAnswerResponse]
type dnsSummaryMatchingAnswerResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryMatchingAnswerResponseMeta struct {
	ConfidenceInfo DNSSummaryMatchingAnswerResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryMatchingAnswerResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryMatchingAnswerResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryMatchingAnswerResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryMatchingAnswerResponseMetaJSON   `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaJSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseMeta]
type dnsSummaryMatchingAnswerResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryMatchingAnswerResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [DNSSummaryMatchingAnswerResponseMetaConfidenceInfo]
type dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation]
type dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryMatchingAnswerResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryMatchingAnswerResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryMatchingAnswerResponseMetaDateRange]
type dnsSummaryMatchingAnswerResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryMatchingAnswerResponseMetaNormalization string

const (
	DNSSummaryMatchingAnswerResponseMetaNormalizationPercentage           DNSSummaryMatchingAnswerResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryMatchingAnswerResponseMetaNormalizationMin0Max              DNSSummaryMatchingAnswerResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryMatchingAnswerResponseMetaNormalizationMinMax               DNSSummaryMatchingAnswerResponseMetaNormalization = "MIN_MAX"
	DNSSummaryMatchingAnswerResponseMetaNormalizationRawValues            DNSSummaryMatchingAnswerResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryMatchingAnswerResponseMetaNormalizationPercentageChange     DNSSummaryMatchingAnswerResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryMatchingAnswerResponseMetaNormalizationRollingAverage       DNSSummaryMatchingAnswerResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryMatchingAnswerResponseMetaNormalizationOverlappedPercentage DNSSummaryMatchingAnswerResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryMatchingAnswerResponseMetaNormalizationRatio                DNSSummaryMatchingAnswerResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryMatchingAnswerResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryMatchingAnswerResponseMetaNormalizationPercentage, DNSSummaryMatchingAnswerResponseMetaNormalizationMin0Max, DNSSummaryMatchingAnswerResponseMetaNormalizationMinMax, DNSSummaryMatchingAnswerResponseMetaNormalizationRawValues, DNSSummaryMatchingAnswerResponseMetaNormalizationPercentageChange, DNSSummaryMatchingAnswerResponseMetaNormalizationRollingAverage, DNSSummaryMatchingAnswerResponseMetaNormalizationOverlappedPercentage, DNSSummaryMatchingAnswerResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryMatchingAnswerResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  dnsSummaryMatchingAnswerResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseMetaUnit]
type dnsSummaryMatchingAnswerResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryMatchingAnswerResponseSummary0 struct {
	// A numeric string.
	Negative string `json:"NEGATIVE,required"`
	// A numeric string.
	Positive string                                       `json:"POSITIVE,required"`
	JSON     dnsSummaryMatchingAnswerResponseSummary0JSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseSummary0]
type dnsSummaryMatchingAnswerResponseSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryProtocolResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryProtocolResponseJSON     `json:"-"`
}

// dnsSummaryProtocolResponseJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponse]
type dnsSummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryProtocolResponseMeta struct {
	ConfidenceInfo DNSSummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryProtocolResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryProtocolResponseMetaJSON   `json:"-"`
}

// dnsSummaryProtocolResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseMeta]
type dnsSummaryProtocolResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  dnsSummaryProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryProtocolResponseMetaConfidenceInfo]
type dnsSummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation]
type dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryProtocolResponseMetaDateRange]
type dnsSummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryProtocolResponseMetaNormalization string

const (
	DNSSummaryProtocolResponseMetaNormalizationPercentage           DNSSummaryProtocolResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryProtocolResponseMetaNormalizationMin0Max              DNSSummaryProtocolResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryProtocolResponseMetaNormalizationMinMax               DNSSummaryProtocolResponseMetaNormalization = "MIN_MAX"
	DNSSummaryProtocolResponseMetaNormalizationRawValues            DNSSummaryProtocolResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryProtocolResponseMetaNormalizationPercentageChange     DNSSummaryProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryProtocolResponseMetaNormalizationRollingAverage       DNSSummaryProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryProtocolResponseMetaNormalizationOverlappedPercentage DNSSummaryProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryProtocolResponseMetaNormalizationRatio                DNSSummaryProtocolResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryProtocolResponseMetaNormalizationPercentage, DNSSummaryProtocolResponseMetaNormalizationMin0Max, DNSSummaryProtocolResponseMetaNormalizationMinMax, DNSSummaryProtocolResponseMetaNormalizationRawValues, DNSSummaryProtocolResponseMetaNormalizationPercentageChange, DNSSummaryProtocolResponseMetaNormalizationRollingAverage, DNSSummaryProtocolResponseMetaNormalizationOverlappedPercentage, DNSSummaryProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryProtocolResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  dnsSummaryProtocolResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseMetaUnit]
type dnsSummaryProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryProtocolResponseSummary0 struct {
	// A numeric string.
	HTTPS string `json:"HTTPS,required"`
	// A numeric string.
	TCP string `json:"TCP,required"`
	// A numeric string.
	TLS string `json:"TLS,required"`
	// A numeric string.
	Udp  string                                 `json:"UDP,required"`
	JSON dnsSummaryProtocolResponseSummary0JSON `json:"-"`
}

// dnsSummaryProtocolResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseSummary0]
type dnsSummaryProtocolResponseSummary0JSON struct {
	HTTPS       apijson.Field
	TCP         apijson.Field
	TLS         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryQueryTypeResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryQueryTypeResponseMeta `json:"meta,required"`
	Summary0 map[string]string               `json:"summary_0,required"`
	JSON     dnsSummaryQueryTypeResponseJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [DNSSummaryQueryTypeResponse]
type dnsSummaryQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryQueryTypeResponseMeta struct {
	ConfidenceInfo DNSSummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryQueryTypeResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryQueryTypeResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryQueryTypeResponseMetaJSON   `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryQueryTypeResponseMeta]
type dnsSummaryQueryTypeResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryQueryTypeResponseMetaConfidenceInfo]
type dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseMetaDateRange]
type dnsSummaryQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryQueryTypeResponseMetaNormalization string

const (
	DNSSummaryQueryTypeResponseMetaNormalizationPercentage           DNSSummaryQueryTypeResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryQueryTypeResponseMetaNormalizationMin0Max              DNSSummaryQueryTypeResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryQueryTypeResponseMetaNormalizationMinMax               DNSSummaryQueryTypeResponseMetaNormalization = "MIN_MAX"
	DNSSummaryQueryTypeResponseMetaNormalizationRawValues            DNSSummaryQueryTypeResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryQueryTypeResponseMetaNormalizationPercentageChange     DNSSummaryQueryTypeResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryQueryTypeResponseMetaNormalizationRollingAverage       DNSSummaryQueryTypeResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryQueryTypeResponseMetaNormalizationOverlappedPercentage DNSSummaryQueryTypeResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryQueryTypeResponseMetaNormalizationRatio                DNSSummaryQueryTypeResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryQueryTypeResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryQueryTypeResponseMetaNormalizationPercentage, DNSSummaryQueryTypeResponseMetaNormalizationMin0Max, DNSSummaryQueryTypeResponseMetaNormalizationMinMax, DNSSummaryQueryTypeResponseMetaNormalizationRawValues, DNSSummaryQueryTypeResponseMetaNormalizationPercentageChange, DNSSummaryQueryTypeResponseMetaNormalizationRollingAverage, DNSSummaryQueryTypeResponseMetaNormalizationOverlappedPercentage, DNSSummaryQueryTypeResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryQueryTypeResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  dnsSummaryQueryTypeResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseMetaUnit]
type dnsSummaryQueryTypeResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseCodeResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryResponseCodeResponseMeta `json:"meta,required"`
	Summary0 map[string]string                  `json:"summary_0,required"`
	JSON     dnsSummaryResponseCodeResponseJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseJSON contains the JSON metadata for the struct
// [DNSSummaryResponseCodeResponse]
type dnsSummaryResponseCodeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryResponseCodeResponseMeta struct {
	ConfidenceInfo DNSSummaryResponseCodeResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryResponseCodeResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryResponseCodeResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryResponseCodeResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryResponseCodeResponseMetaJSON   `json:"-"`
}

// dnsSummaryResponseCodeResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryResponseCodeResponseMeta]
type dnsSummaryResponseCodeResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseCodeResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryResponseCodeResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  dnsSummaryResponseCodeResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryResponseCodeResponseMetaConfidenceInfo]
type dnsSummaryResponseCodeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryResponseCodeResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryResponseCodeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryResponseCodeResponseMetaConfidenceInfoAnnotation]
type dnsSummaryResponseCodeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryResponseCodeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseCodeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryResponseCodeResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryResponseCodeResponseMetaDateRange]
type dnsSummaryResponseCodeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryResponseCodeResponseMetaNormalization string

const (
	DNSSummaryResponseCodeResponseMetaNormalizationPercentage           DNSSummaryResponseCodeResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryResponseCodeResponseMetaNormalizationMin0Max              DNSSummaryResponseCodeResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryResponseCodeResponseMetaNormalizationMinMax               DNSSummaryResponseCodeResponseMetaNormalization = "MIN_MAX"
	DNSSummaryResponseCodeResponseMetaNormalizationRawValues            DNSSummaryResponseCodeResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryResponseCodeResponseMetaNormalizationPercentageChange     DNSSummaryResponseCodeResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryResponseCodeResponseMetaNormalizationRollingAverage       DNSSummaryResponseCodeResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryResponseCodeResponseMetaNormalizationOverlappedPercentage DNSSummaryResponseCodeResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryResponseCodeResponseMetaNormalizationRatio                DNSSummaryResponseCodeResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryResponseCodeResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryResponseCodeResponseMetaNormalizationPercentage, DNSSummaryResponseCodeResponseMetaNormalizationMin0Max, DNSSummaryResponseCodeResponseMetaNormalizationMinMax, DNSSummaryResponseCodeResponseMetaNormalizationRawValues, DNSSummaryResponseCodeResponseMetaNormalizationPercentageChange, DNSSummaryResponseCodeResponseMetaNormalizationRollingAverage, DNSSummaryResponseCodeResponseMetaNormalizationOverlappedPercentage, DNSSummaryResponseCodeResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryResponseCodeResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  dnsSummaryResponseCodeResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryResponseCodeResponseMetaUnit]
type dnsSummaryResponseCodeResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseTTLResponse struct {
	// Metadata for the results.
	Meta     DNSSummaryResponseTTLResponseMeta     `json:"meta,required"`
	Summary0 DNSSummaryResponseTTLResponseSummary0 `json:"summary_0,required"`
	JSON     dnsSummaryResponseTTLResponseJSON     `json:"-"`
}

// dnsSummaryResponseTTLResponseJSON contains the JSON metadata for the struct
// [DNSSummaryResponseTTLResponse]
type dnsSummaryResponseTTLResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSSummaryResponseTTLResponseMeta struct {
	ConfidenceInfo DNSSummaryResponseTTLResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []DNSSummaryResponseTTLResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSSummaryResponseTTLResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSSummaryResponseTTLResponseMetaUnit `json:"units,required"`
	JSON  dnsSummaryResponseTTLResponseMetaJSON   `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryResponseTTLResponseMeta]
type dnsSummaryResponseTTLResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseTTLResponseMetaConfidenceInfo struct {
	Annotations []DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryResponseTTLResponseMetaConfidenceInfo]
type dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation]
type dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseTTLResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      dnsSummaryResponseTTLResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryResponseTTLResponseMetaDateRange]
type dnsSummaryResponseTTLResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSSummaryResponseTTLResponseMetaNormalization string

const (
	DNSSummaryResponseTTLResponseMetaNormalizationPercentage           DNSSummaryResponseTTLResponseMetaNormalization = "PERCENTAGE"
	DNSSummaryResponseTTLResponseMetaNormalizationMin0Max              DNSSummaryResponseTTLResponseMetaNormalization = "MIN0_MAX"
	DNSSummaryResponseTTLResponseMetaNormalizationMinMax               DNSSummaryResponseTTLResponseMetaNormalization = "MIN_MAX"
	DNSSummaryResponseTTLResponseMetaNormalizationRawValues            DNSSummaryResponseTTLResponseMetaNormalization = "RAW_VALUES"
	DNSSummaryResponseTTLResponseMetaNormalizationPercentageChange     DNSSummaryResponseTTLResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSSummaryResponseTTLResponseMetaNormalizationRollingAverage       DNSSummaryResponseTTLResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSSummaryResponseTTLResponseMetaNormalizationOverlappedPercentage DNSSummaryResponseTTLResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSSummaryResponseTTLResponseMetaNormalizationRatio                DNSSummaryResponseTTLResponseMetaNormalization = "RATIO"
)

func (r DNSSummaryResponseTTLResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSSummaryResponseTTLResponseMetaNormalizationPercentage, DNSSummaryResponseTTLResponseMetaNormalizationMin0Max, DNSSummaryResponseTTLResponseMetaNormalizationMinMax, DNSSummaryResponseTTLResponseMetaNormalizationRawValues, DNSSummaryResponseTTLResponseMetaNormalizationPercentageChange, DNSSummaryResponseTTLResponseMetaNormalizationRollingAverage, DNSSummaryResponseTTLResponseMetaNormalizationOverlappedPercentage, DNSSummaryResponseTTLResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSSummaryResponseTTLResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  dnsSummaryResponseTTLResponseMetaUnitJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaUnitJSON contains the JSON metadata for the
// struct [DNSSummaryResponseTTLResponseMetaUnit]
type dnsSummaryResponseTTLResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseTTLResponseSummary0 struct {
	// A numeric string.
	Gt15mLte1h string `json:"gt_15m_lte_1h,required"`
	// A numeric string.
	Gt1dLte1w string `json:"gt_1d_lte_1w,required"`
	// A numeric string.
	Gt1hLte1d string `json:"gt_1h_lte_1d,required"`
	// A numeric string.
	Gt1mLte5m string `json:"gt_1m_lte_5m,required"`
	// A numeric string.
	Gt1w string `json:"gt_1w,required"`
	// A numeric string.
	Gt5mLte15m string `json:"gt_5m_lte_15m,required"`
	// A numeric string.
	Lte1m string                                    `json:"lte_1m,required"`
	JSON  dnsSummaryResponseTTLResponseSummary0JSON `json:"-"`
}

// dnsSummaryResponseTTLResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryResponseTTLResponseSummary0]
type dnsSummaryResponseTTLResponseSummary0JSON struct {
	Gt15mLte1h  apijson.Field
	Gt1dLte1w   apijson.Field
	Gt1hLte1d   apijson.Field
	Gt1mLte5m   apijson.Field
	Gt1w        apijson.Field
	Gt5mLte15m  apijson.Field
	Lte1m       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type DNSSummaryCacheHitParams struct {
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
	Format param.Field[DNSSummaryCacheHitParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryCacheHitParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryCacheHitParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryCacheHitParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryCacheHitParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryCacheHitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryCacheHitParamsFormat string

const (
	DNSSummaryCacheHitParamsFormatJson DNSSummaryCacheHitParamsFormat = "JSON"
	DNSSummaryCacheHitParamsFormatCsv  DNSSummaryCacheHitParamsFormat = "CSV"
)

func (r DNSSummaryCacheHitParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryCacheHitParamsFormatJson, DNSSummaryCacheHitParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryCacheHitParamsProtocol string

const (
	DNSSummaryCacheHitParamsProtocolUdp   DNSSummaryCacheHitParamsProtocol = "UDP"
	DNSSummaryCacheHitParamsProtocolTCP   DNSSummaryCacheHitParamsProtocol = "TCP"
	DNSSummaryCacheHitParamsProtocolHTTPS DNSSummaryCacheHitParamsProtocol = "HTTPS"
	DNSSummaryCacheHitParamsProtocolTLS   DNSSummaryCacheHitParamsProtocol = "TLS"
)

func (r DNSSummaryCacheHitParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryCacheHitParamsProtocolUdp, DNSSummaryCacheHitParamsProtocolTCP, DNSSummaryCacheHitParamsProtocolHTTPS, DNSSummaryCacheHitParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryCacheHitParamsQueryType string

const (
	DNSSummaryCacheHitParamsQueryTypeA          DNSSummaryCacheHitParamsQueryType = "A"
	DNSSummaryCacheHitParamsQueryTypeAAAA       DNSSummaryCacheHitParamsQueryType = "AAAA"
	DNSSummaryCacheHitParamsQueryTypeA6         DNSSummaryCacheHitParamsQueryType = "A6"
	DNSSummaryCacheHitParamsQueryTypeAfsdb      DNSSummaryCacheHitParamsQueryType = "AFSDB"
	DNSSummaryCacheHitParamsQueryTypeAny        DNSSummaryCacheHitParamsQueryType = "ANY"
	DNSSummaryCacheHitParamsQueryTypeApl        DNSSummaryCacheHitParamsQueryType = "APL"
	DNSSummaryCacheHitParamsQueryTypeAtma       DNSSummaryCacheHitParamsQueryType = "ATMA"
	DNSSummaryCacheHitParamsQueryTypeAXFR       DNSSummaryCacheHitParamsQueryType = "AXFR"
	DNSSummaryCacheHitParamsQueryTypeCAA        DNSSummaryCacheHitParamsQueryType = "CAA"
	DNSSummaryCacheHitParamsQueryTypeCdnskey    DNSSummaryCacheHitParamsQueryType = "CDNSKEY"
	DNSSummaryCacheHitParamsQueryTypeCds        DNSSummaryCacheHitParamsQueryType = "CDS"
	DNSSummaryCacheHitParamsQueryTypeCERT       DNSSummaryCacheHitParamsQueryType = "CERT"
	DNSSummaryCacheHitParamsQueryTypeCNAME      DNSSummaryCacheHitParamsQueryType = "CNAME"
	DNSSummaryCacheHitParamsQueryTypeCsync      DNSSummaryCacheHitParamsQueryType = "CSYNC"
	DNSSummaryCacheHitParamsQueryTypeDhcid      DNSSummaryCacheHitParamsQueryType = "DHCID"
	DNSSummaryCacheHitParamsQueryTypeDlv        DNSSummaryCacheHitParamsQueryType = "DLV"
	DNSSummaryCacheHitParamsQueryTypeDname      DNSSummaryCacheHitParamsQueryType = "DNAME"
	DNSSummaryCacheHitParamsQueryTypeDNSKEY     DNSSummaryCacheHitParamsQueryType = "DNSKEY"
	DNSSummaryCacheHitParamsQueryTypeDoa        DNSSummaryCacheHitParamsQueryType = "DOA"
	DNSSummaryCacheHitParamsQueryTypeDS         DNSSummaryCacheHitParamsQueryType = "DS"
	DNSSummaryCacheHitParamsQueryTypeEid        DNSSummaryCacheHitParamsQueryType = "EID"
	DNSSummaryCacheHitParamsQueryTypeEui48      DNSSummaryCacheHitParamsQueryType = "EUI48"
	DNSSummaryCacheHitParamsQueryTypeEui64      DNSSummaryCacheHitParamsQueryType = "EUI64"
	DNSSummaryCacheHitParamsQueryTypeGpos       DNSSummaryCacheHitParamsQueryType = "GPOS"
	DNSSummaryCacheHitParamsQueryTypeGid        DNSSummaryCacheHitParamsQueryType = "GID"
	DNSSummaryCacheHitParamsQueryTypeHinfo      DNSSummaryCacheHitParamsQueryType = "HINFO"
	DNSSummaryCacheHitParamsQueryTypeHip        DNSSummaryCacheHitParamsQueryType = "HIP"
	DNSSummaryCacheHitParamsQueryTypeHTTPS      DNSSummaryCacheHitParamsQueryType = "HTTPS"
	DNSSummaryCacheHitParamsQueryTypeIpseckey   DNSSummaryCacheHitParamsQueryType = "IPSECKEY"
	DNSSummaryCacheHitParamsQueryTypeIsdn       DNSSummaryCacheHitParamsQueryType = "ISDN"
	DNSSummaryCacheHitParamsQueryTypeIxfr       DNSSummaryCacheHitParamsQueryType = "IXFR"
	DNSSummaryCacheHitParamsQueryTypeKey        DNSSummaryCacheHitParamsQueryType = "KEY"
	DNSSummaryCacheHitParamsQueryTypeKx         DNSSummaryCacheHitParamsQueryType = "KX"
	DNSSummaryCacheHitParamsQueryTypeL32        DNSSummaryCacheHitParamsQueryType = "L32"
	DNSSummaryCacheHitParamsQueryTypeL64        DNSSummaryCacheHitParamsQueryType = "L64"
	DNSSummaryCacheHitParamsQueryTypeLOC        DNSSummaryCacheHitParamsQueryType = "LOC"
	DNSSummaryCacheHitParamsQueryTypeLp         DNSSummaryCacheHitParamsQueryType = "LP"
	DNSSummaryCacheHitParamsQueryTypeMaila      DNSSummaryCacheHitParamsQueryType = "MAILA"
	DNSSummaryCacheHitParamsQueryTypeMailb      DNSSummaryCacheHitParamsQueryType = "MAILB"
	DNSSummaryCacheHitParamsQueryTypeMB         DNSSummaryCacheHitParamsQueryType = "MB"
	DNSSummaryCacheHitParamsQueryTypeMd         DNSSummaryCacheHitParamsQueryType = "MD"
	DNSSummaryCacheHitParamsQueryTypeMf         DNSSummaryCacheHitParamsQueryType = "MF"
	DNSSummaryCacheHitParamsQueryTypeMg         DNSSummaryCacheHitParamsQueryType = "MG"
	DNSSummaryCacheHitParamsQueryTypeMinfo      DNSSummaryCacheHitParamsQueryType = "MINFO"
	DNSSummaryCacheHitParamsQueryTypeMr         DNSSummaryCacheHitParamsQueryType = "MR"
	DNSSummaryCacheHitParamsQueryTypeMX         DNSSummaryCacheHitParamsQueryType = "MX"
	DNSSummaryCacheHitParamsQueryTypeNAPTR      DNSSummaryCacheHitParamsQueryType = "NAPTR"
	DNSSummaryCacheHitParamsQueryTypeNb         DNSSummaryCacheHitParamsQueryType = "NB"
	DNSSummaryCacheHitParamsQueryTypeNbstat     DNSSummaryCacheHitParamsQueryType = "NBSTAT"
	DNSSummaryCacheHitParamsQueryTypeNid        DNSSummaryCacheHitParamsQueryType = "NID"
	DNSSummaryCacheHitParamsQueryTypeNimloc     DNSSummaryCacheHitParamsQueryType = "NIMLOC"
	DNSSummaryCacheHitParamsQueryTypeNinfo      DNSSummaryCacheHitParamsQueryType = "NINFO"
	DNSSummaryCacheHitParamsQueryTypeNS         DNSSummaryCacheHitParamsQueryType = "NS"
	DNSSummaryCacheHitParamsQueryTypeNsap       DNSSummaryCacheHitParamsQueryType = "NSAP"
	DNSSummaryCacheHitParamsQueryTypeNsec       DNSSummaryCacheHitParamsQueryType = "NSEC"
	DNSSummaryCacheHitParamsQueryTypeNsec3      DNSSummaryCacheHitParamsQueryType = "NSEC3"
	DNSSummaryCacheHitParamsQueryTypeNsec3Param DNSSummaryCacheHitParamsQueryType = "NSEC3PARAM"
	DNSSummaryCacheHitParamsQueryTypeNull       DNSSummaryCacheHitParamsQueryType = "NULL"
	DNSSummaryCacheHitParamsQueryTypeNxt        DNSSummaryCacheHitParamsQueryType = "NXT"
	DNSSummaryCacheHitParamsQueryTypeOpenpgpkey DNSSummaryCacheHitParamsQueryType = "OPENPGPKEY"
	DNSSummaryCacheHitParamsQueryTypeOpt        DNSSummaryCacheHitParamsQueryType = "OPT"
	DNSSummaryCacheHitParamsQueryTypePTR        DNSSummaryCacheHitParamsQueryType = "PTR"
	DNSSummaryCacheHitParamsQueryTypePx         DNSSummaryCacheHitParamsQueryType = "PX"
	DNSSummaryCacheHitParamsQueryTypeRkey       DNSSummaryCacheHitParamsQueryType = "RKEY"
	DNSSummaryCacheHitParamsQueryTypeRp         DNSSummaryCacheHitParamsQueryType = "RP"
	DNSSummaryCacheHitParamsQueryTypeRrsig      DNSSummaryCacheHitParamsQueryType = "RRSIG"
	DNSSummaryCacheHitParamsQueryTypeRt         DNSSummaryCacheHitParamsQueryType = "RT"
	DNSSummaryCacheHitParamsQueryTypeSig        DNSSummaryCacheHitParamsQueryType = "SIG"
	DNSSummaryCacheHitParamsQueryTypeSink       DNSSummaryCacheHitParamsQueryType = "SINK"
	DNSSummaryCacheHitParamsQueryTypeSMIMEA     DNSSummaryCacheHitParamsQueryType = "SMIMEA"
	DNSSummaryCacheHitParamsQueryTypeSOA        DNSSummaryCacheHitParamsQueryType = "SOA"
	DNSSummaryCacheHitParamsQueryTypeSPF        DNSSummaryCacheHitParamsQueryType = "SPF"
	DNSSummaryCacheHitParamsQueryTypeSRV        DNSSummaryCacheHitParamsQueryType = "SRV"
	DNSSummaryCacheHitParamsQueryTypeSSHFP      DNSSummaryCacheHitParamsQueryType = "SSHFP"
	DNSSummaryCacheHitParamsQueryTypeSVCB       DNSSummaryCacheHitParamsQueryType = "SVCB"
	DNSSummaryCacheHitParamsQueryTypeTa         DNSSummaryCacheHitParamsQueryType = "TA"
	DNSSummaryCacheHitParamsQueryTypeTalink     DNSSummaryCacheHitParamsQueryType = "TALINK"
	DNSSummaryCacheHitParamsQueryTypeTkey       DNSSummaryCacheHitParamsQueryType = "TKEY"
	DNSSummaryCacheHitParamsQueryTypeTLSA       DNSSummaryCacheHitParamsQueryType = "TLSA"
	DNSSummaryCacheHitParamsQueryTypeTSIG       DNSSummaryCacheHitParamsQueryType = "TSIG"
	DNSSummaryCacheHitParamsQueryTypeTXT        DNSSummaryCacheHitParamsQueryType = "TXT"
	DNSSummaryCacheHitParamsQueryTypeUinfo      DNSSummaryCacheHitParamsQueryType = "UINFO"
	DNSSummaryCacheHitParamsQueryTypeUID        DNSSummaryCacheHitParamsQueryType = "UID"
	DNSSummaryCacheHitParamsQueryTypeUnspec     DNSSummaryCacheHitParamsQueryType = "UNSPEC"
	DNSSummaryCacheHitParamsQueryTypeURI        DNSSummaryCacheHitParamsQueryType = "URI"
	DNSSummaryCacheHitParamsQueryTypeWks        DNSSummaryCacheHitParamsQueryType = "WKS"
	DNSSummaryCacheHitParamsQueryTypeX25        DNSSummaryCacheHitParamsQueryType = "X25"
	DNSSummaryCacheHitParamsQueryTypeZonemd     DNSSummaryCacheHitParamsQueryType = "ZONEMD"
)

func (r DNSSummaryCacheHitParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryCacheHitParamsQueryTypeA, DNSSummaryCacheHitParamsQueryTypeAAAA, DNSSummaryCacheHitParamsQueryTypeA6, DNSSummaryCacheHitParamsQueryTypeAfsdb, DNSSummaryCacheHitParamsQueryTypeAny, DNSSummaryCacheHitParamsQueryTypeApl, DNSSummaryCacheHitParamsQueryTypeAtma, DNSSummaryCacheHitParamsQueryTypeAXFR, DNSSummaryCacheHitParamsQueryTypeCAA, DNSSummaryCacheHitParamsQueryTypeCdnskey, DNSSummaryCacheHitParamsQueryTypeCds, DNSSummaryCacheHitParamsQueryTypeCERT, DNSSummaryCacheHitParamsQueryTypeCNAME, DNSSummaryCacheHitParamsQueryTypeCsync, DNSSummaryCacheHitParamsQueryTypeDhcid, DNSSummaryCacheHitParamsQueryTypeDlv, DNSSummaryCacheHitParamsQueryTypeDname, DNSSummaryCacheHitParamsQueryTypeDNSKEY, DNSSummaryCacheHitParamsQueryTypeDoa, DNSSummaryCacheHitParamsQueryTypeDS, DNSSummaryCacheHitParamsQueryTypeEid, DNSSummaryCacheHitParamsQueryTypeEui48, DNSSummaryCacheHitParamsQueryTypeEui64, DNSSummaryCacheHitParamsQueryTypeGpos, DNSSummaryCacheHitParamsQueryTypeGid, DNSSummaryCacheHitParamsQueryTypeHinfo, DNSSummaryCacheHitParamsQueryTypeHip, DNSSummaryCacheHitParamsQueryTypeHTTPS, DNSSummaryCacheHitParamsQueryTypeIpseckey, DNSSummaryCacheHitParamsQueryTypeIsdn, DNSSummaryCacheHitParamsQueryTypeIxfr, DNSSummaryCacheHitParamsQueryTypeKey, DNSSummaryCacheHitParamsQueryTypeKx, DNSSummaryCacheHitParamsQueryTypeL32, DNSSummaryCacheHitParamsQueryTypeL64, DNSSummaryCacheHitParamsQueryTypeLOC, DNSSummaryCacheHitParamsQueryTypeLp, DNSSummaryCacheHitParamsQueryTypeMaila, DNSSummaryCacheHitParamsQueryTypeMailb, DNSSummaryCacheHitParamsQueryTypeMB, DNSSummaryCacheHitParamsQueryTypeMd, DNSSummaryCacheHitParamsQueryTypeMf, DNSSummaryCacheHitParamsQueryTypeMg, DNSSummaryCacheHitParamsQueryTypeMinfo, DNSSummaryCacheHitParamsQueryTypeMr, DNSSummaryCacheHitParamsQueryTypeMX, DNSSummaryCacheHitParamsQueryTypeNAPTR, DNSSummaryCacheHitParamsQueryTypeNb, DNSSummaryCacheHitParamsQueryTypeNbstat, DNSSummaryCacheHitParamsQueryTypeNid, DNSSummaryCacheHitParamsQueryTypeNimloc, DNSSummaryCacheHitParamsQueryTypeNinfo, DNSSummaryCacheHitParamsQueryTypeNS, DNSSummaryCacheHitParamsQueryTypeNsap, DNSSummaryCacheHitParamsQueryTypeNsec, DNSSummaryCacheHitParamsQueryTypeNsec3, DNSSummaryCacheHitParamsQueryTypeNsec3Param, DNSSummaryCacheHitParamsQueryTypeNull, DNSSummaryCacheHitParamsQueryTypeNxt, DNSSummaryCacheHitParamsQueryTypeOpenpgpkey, DNSSummaryCacheHitParamsQueryTypeOpt, DNSSummaryCacheHitParamsQueryTypePTR, DNSSummaryCacheHitParamsQueryTypePx, DNSSummaryCacheHitParamsQueryTypeRkey, DNSSummaryCacheHitParamsQueryTypeRp, DNSSummaryCacheHitParamsQueryTypeRrsig, DNSSummaryCacheHitParamsQueryTypeRt, DNSSummaryCacheHitParamsQueryTypeSig, DNSSummaryCacheHitParamsQueryTypeSink, DNSSummaryCacheHitParamsQueryTypeSMIMEA, DNSSummaryCacheHitParamsQueryTypeSOA, DNSSummaryCacheHitParamsQueryTypeSPF, DNSSummaryCacheHitParamsQueryTypeSRV, DNSSummaryCacheHitParamsQueryTypeSSHFP, DNSSummaryCacheHitParamsQueryTypeSVCB, DNSSummaryCacheHitParamsQueryTypeTa, DNSSummaryCacheHitParamsQueryTypeTalink, DNSSummaryCacheHitParamsQueryTypeTkey, DNSSummaryCacheHitParamsQueryTypeTLSA, DNSSummaryCacheHitParamsQueryTypeTSIG, DNSSummaryCacheHitParamsQueryTypeTXT, DNSSummaryCacheHitParamsQueryTypeUinfo, DNSSummaryCacheHitParamsQueryTypeUID, DNSSummaryCacheHitParamsQueryTypeUnspec, DNSSummaryCacheHitParamsQueryTypeURI, DNSSummaryCacheHitParamsQueryTypeWks, DNSSummaryCacheHitParamsQueryTypeX25, DNSSummaryCacheHitParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryCacheHitParamsResponseCode string

const (
	DNSSummaryCacheHitParamsResponseCodeNoerror   DNSSummaryCacheHitParamsResponseCode = "NOERROR"
	DNSSummaryCacheHitParamsResponseCodeFormerr   DNSSummaryCacheHitParamsResponseCode = "FORMERR"
	DNSSummaryCacheHitParamsResponseCodeServfail  DNSSummaryCacheHitParamsResponseCode = "SERVFAIL"
	DNSSummaryCacheHitParamsResponseCodeNxdomain  DNSSummaryCacheHitParamsResponseCode = "NXDOMAIN"
	DNSSummaryCacheHitParamsResponseCodeNotimp    DNSSummaryCacheHitParamsResponseCode = "NOTIMP"
	DNSSummaryCacheHitParamsResponseCodeRefused   DNSSummaryCacheHitParamsResponseCode = "REFUSED"
	DNSSummaryCacheHitParamsResponseCodeYxdomain  DNSSummaryCacheHitParamsResponseCode = "YXDOMAIN"
	DNSSummaryCacheHitParamsResponseCodeYxrrset   DNSSummaryCacheHitParamsResponseCode = "YXRRSET"
	DNSSummaryCacheHitParamsResponseCodeNxrrset   DNSSummaryCacheHitParamsResponseCode = "NXRRSET"
	DNSSummaryCacheHitParamsResponseCodeNotauth   DNSSummaryCacheHitParamsResponseCode = "NOTAUTH"
	DNSSummaryCacheHitParamsResponseCodeNotzone   DNSSummaryCacheHitParamsResponseCode = "NOTZONE"
	DNSSummaryCacheHitParamsResponseCodeBadsig    DNSSummaryCacheHitParamsResponseCode = "BADSIG"
	DNSSummaryCacheHitParamsResponseCodeBadkey    DNSSummaryCacheHitParamsResponseCode = "BADKEY"
	DNSSummaryCacheHitParamsResponseCodeBadtime   DNSSummaryCacheHitParamsResponseCode = "BADTIME"
	DNSSummaryCacheHitParamsResponseCodeBadmode   DNSSummaryCacheHitParamsResponseCode = "BADMODE"
	DNSSummaryCacheHitParamsResponseCodeBadname   DNSSummaryCacheHitParamsResponseCode = "BADNAME"
	DNSSummaryCacheHitParamsResponseCodeBadalg    DNSSummaryCacheHitParamsResponseCode = "BADALG"
	DNSSummaryCacheHitParamsResponseCodeBadtrunc  DNSSummaryCacheHitParamsResponseCode = "BADTRUNC"
	DNSSummaryCacheHitParamsResponseCodeBadcookie DNSSummaryCacheHitParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryCacheHitParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryCacheHitParamsResponseCodeNoerror, DNSSummaryCacheHitParamsResponseCodeFormerr, DNSSummaryCacheHitParamsResponseCodeServfail, DNSSummaryCacheHitParamsResponseCodeNxdomain, DNSSummaryCacheHitParamsResponseCodeNotimp, DNSSummaryCacheHitParamsResponseCodeRefused, DNSSummaryCacheHitParamsResponseCodeYxdomain, DNSSummaryCacheHitParamsResponseCodeYxrrset, DNSSummaryCacheHitParamsResponseCodeNxrrset, DNSSummaryCacheHitParamsResponseCodeNotauth, DNSSummaryCacheHitParamsResponseCodeNotzone, DNSSummaryCacheHitParamsResponseCodeBadsig, DNSSummaryCacheHitParamsResponseCodeBadkey, DNSSummaryCacheHitParamsResponseCodeBadtime, DNSSummaryCacheHitParamsResponseCodeBadmode, DNSSummaryCacheHitParamsResponseCodeBadname, DNSSummaryCacheHitParamsResponseCodeBadalg, DNSSummaryCacheHitParamsResponseCodeBadtrunc, DNSSummaryCacheHitParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryCacheHitResponseEnvelope struct {
	Result  DNSSummaryCacheHitResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    dnsSummaryCacheHitResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryCacheHitResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseEnvelope]
type dnsSummaryCacheHitResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECParams struct {
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
	Format param.Field[DNSSummaryDNSSECParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryDNSSECParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryDNSSECParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryDNSSECParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECParams]'s query parameters as `url.Values`.
func (r DNSSummaryDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryDNSSECParamsFormat string

const (
	DNSSummaryDNSSECParamsFormatJson DNSSummaryDNSSECParamsFormat = "JSON"
	DNSSummaryDNSSECParamsFormatCsv  DNSSummaryDNSSECParamsFormat = "CSV"
)

func (r DNSSummaryDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECParamsFormatJson, DNSSummaryDNSSECParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryDNSSECParamsProtocol string

const (
	DNSSummaryDNSSECParamsProtocolUdp   DNSSummaryDNSSECParamsProtocol = "UDP"
	DNSSummaryDNSSECParamsProtocolTCP   DNSSummaryDNSSECParamsProtocol = "TCP"
	DNSSummaryDNSSECParamsProtocolHTTPS DNSSummaryDNSSECParamsProtocol = "HTTPS"
	DNSSummaryDNSSECParamsProtocolTLS   DNSSummaryDNSSECParamsProtocol = "TLS"
)

func (r DNSSummaryDNSSECParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECParamsProtocolUdp, DNSSummaryDNSSECParamsProtocolTCP, DNSSummaryDNSSECParamsProtocolHTTPS, DNSSummaryDNSSECParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryDNSSECParamsQueryType string

const (
	DNSSummaryDNSSECParamsQueryTypeA          DNSSummaryDNSSECParamsQueryType = "A"
	DNSSummaryDNSSECParamsQueryTypeAAAA       DNSSummaryDNSSECParamsQueryType = "AAAA"
	DNSSummaryDNSSECParamsQueryTypeA6         DNSSummaryDNSSECParamsQueryType = "A6"
	DNSSummaryDNSSECParamsQueryTypeAfsdb      DNSSummaryDNSSECParamsQueryType = "AFSDB"
	DNSSummaryDNSSECParamsQueryTypeAny        DNSSummaryDNSSECParamsQueryType = "ANY"
	DNSSummaryDNSSECParamsQueryTypeApl        DNSSummaryDNSSECParamsQueryType = "APL"
	DNSSummaryDNSSECParamsQueryTypeAtma       DNSSummaryDNSSECParamsQueryType = "ATMA"
	DNSSummaryDNSSECParamsQueryTypeAXFR       DNSSummaryDNSSECParamsQueryType = "AXFR"
	DNSSummaryDNSSECParamsQueryTypeCAA        DNSSummaryDNSSECParamsQueryType = "CAA"
	DNSSummaryDNSSECParamsQueryTypeCdnskey    DNSSummaryDNSSECParamsQueryType = "CDNSKEY"
	DNSSummaryDNSSECParamsQueryTypeCds        DNSSummaryDNSSECParamsQueryType = "CDS"
	DNSSummaryDNSSECParamsQueryTypeCERT       DNSSummaryDNSSECParamsQueryType = "CERT"
	DNSSummaryDNSSECParamsQueryTypeCNAME      DNSSummaryDNSSECParamsQueryType = "CNAME"
	DNSSummaryDNSSECParamsQueryTypeCsync      DNSSummaryDNSSECParamsQueryType = "CSYNC"
	DNSSummaryDNSSECParamsQueryTypeDhcid      DNSSummaryDNSSECParamsQueryType = "DHCID"
	DNSSummaryDNSSECParamsQueryTypeDlv        DNSSummaryDNSSECParamsQueryType = "DLV"
	DNSSummaryDNSSECParamsQueryTypeDname      DNSSummaryDNSSECParamsQueryType = "DNAME"
	DNSSummaryDNSSECParamsQueryTypeDNSKEY     DNSSummaryDNSSECParamsQueryType = "DNSKEY"
	DNSSummaryDNSSECParamsQueryTypeDoa        DNSSummaryDNSSECParamsQueryType = "DOA"
	DNSSummaryDNSSECParamsQueryTypeDS         DNSSummaryDNSSECParamsQueryType = "DS"
	DNSSummaryDNSSECParamsQueryTypeEid        DNSSummaryDNSSECParamsQueryType = "EID"
	DNSSummaryDNSSECParamsQueryTypeEui48      DNSSummaryDNSSECParamsQueryType = "EUI48"
	DNSSummaryDNSSECParamsQueryTypeEui64      DNSSummaryDNSSECParamsQueryType = "EUI64"
	DNSSummaryDNSSECParamsQueryTypeGpos       DNSSummaryDNSSECParamsQueryType = "GPOS"
	DNSSummaryDNSSECParamsQueryTypeGid        DNSSummaryDNSSECParamsQueryType = "GID"
	DNSSummaryDNSSECParamsQueryTypeHinfo      DNSSummaryDNSSECParamsQueryType = "HINFO"
	DNSSummaryDNSSECParamsQueryTypeHip        DNSSummaryDNSSECParamsQueryType = "HIP"
	DNSSummaryDNSSECParamsQueryTypeHTTPS      DNSSummaryDNSSECParamsQueryType = "HTTPS"
	DNSSummaryDNSSECParamsQueryTypeIpseckey   DNSSummaryDNSSECParamsQueryType = "IPSECKEY"
	DNSSummaryDNSSECParamsQueryTypeIsdn       DNSSummaryDNSSECParamsQueryType = "ISDN"
	DNSSummaryDNSSECParamsQueryTypeIxfr       DNSSummaryDNSSECParamsQueryType = "IXFR"
	DNSSummaryDNSSECParamsQueryTypeKey        DNSSummaryDNSSECParamsQueryType = "KEY"
	DNSSummaryDNSSECParamsQueryTypeKx         DNSSummaryDNSSECParamsQueryType = "KX"
	DNSSummaryDNSSECParamsQueryTypeL32        DNSSummaryDNSSECParamsQueryType = "L32"
	DNSSummaryDNSSECParamsQueryTypeL64        DNSSummaryDNSSECParamsQueryType = "L64"
	DNSSummaryDNSSECParamsQueryTypeLOC        DNSSummaryDNSSECParamsQueryType = "LOC"
	DNSSummaryDNSSECParamsQueryTypeLp         DNSSummaryDNSSECParamsQueryType = "LP"
	DNSSummaryDNSSECParamsQueryTypeMaila      DNSSummaryDNSSECParamsQueryType = "MAILA"
	DNSSummaryDNSSECParamsQueryTypeMailb      DNSSummaryDNSSECParamsQueryType = "MAILB"
	DNSSummaryDNSSECParamsQueryTypeMB         DNSSummaryDNSSECParamsQueryType = "MB"
	DNSSummaryDNSSECParamsQueryTypeMd         DNSSummaryDNSSECParamsQueryType = "MD"
	DNSSummaryDNSSECParamsQueryTypeMf         DNSSummaryDNSSECParamsQueryType = "MF"
	DNSSummaryDNSSECParamsQueryTypeMg         DNSSummaryDNSSECParamsQueryType = "MG"
	DNSSummaryDNSSECParamsQueryTypeMinfo      DNSSummaryDNSSECParamsQueryType = "MINFO"
	DNSSummaryDNSSECParamsQueryTypeMr         DNSSummaryDNSSECParamsQueryType = "MR"
	DNSSummaryDNSSECParamsQueryTypeMX         DNSSummaryDNSSECParamsQueryType = "MX"
	DNSSummaryDNSSECParamsQueryTypeNAPTR      DNSSummaryDNSSECParamsQueryType = "NAPTR"
	DNSSummaryDNSSECParamsQueryTypeNb         DNSSummaryDNSSECParamsQueryType = "NB"
	DNSSummaryDNSSECParamsQueryTypeNbstat     DNSSummaryDNSSECParamsQueryType = "NBSTAT"
	DNSSummaryDNSSECParamsQueryTypeNid        DNSSummaryDNSSECParamsQueryType = "NID"
	DNSSummaryDNSSECParamsQueryTypeNimloc     DNSSummaryDNSSECParamsQueryType = "NIMLOC"
	DNSSummaryDNSSECParamsQueryTypeNinfo      DNSSummaryDNSSECParamsQueryType = "NINFO"
	DNSSummaryDNSSECParamsQueryTypeNS         DNSSummaryDNSSECParamsQueryType = "NS"
	DNSSummaryDNSSECParamsQueryTypeNsap       DNSSummaryDNSSECParamsQueryType = "NSAP"
	DNSSummaryDNSSECParamsQueryTypeNsec       DNSSummaryDNSSECParamsQueryType = "NSEC"
	DNSSummaryDNSSECParamsQueryTypeNsec3      DNSSummaryDNSSECParamsQueryType = "NSEC3"
	DNSSummaryDNSSECParamsQueryTypeNsec3Param DNSSummaryDNSSECParamsQueryType = "NSEC3PARAM"
	DNSSummaryDNSSECParamsQueryTypeNull       DNSSummaryDNSSECParamsQueryType = "NULL"
	DNSSummaryDNSSECParamsQueryTypeNxt        DNSSummaryDNSSECParamsQueryType = "NXT"
	DNSSummaryDNSSECParamsQueryTypeOpenpgpkey DNSSummaryDNSSECParamsQueryType = "OPENPGPKEY"
	DNSSummaryDNSSECParamsQueryTypeOpt        DNSSummaryDNSSECParamsQueryType = "OPT"
	DNSSummaryDNSSECParamsQueryTypePTR        DNSSummaryDNSSECParamsQueryType = "PTR"
	DNSSummaryDNSSECParamsQueryTypePx         DNSSummaryDNSSECParamsQueryType = "PX"
	DNSSummaryDNSSECParamsQueryTypeRkey       DNSSummaryDNSSECParamsQueryType = "RKEY"
	DNSSummaryDNSSECParamsQueryTypeRp         DNSSummaryDNSSECParamsQueryType = "RP"
	DNSSummaryDNSSECParamsQueryTypeRrsig      DNSSummaryDNSSECParamsQueryType = "RRSIG"
	DNSSummaryDNSSECParamsQueryTypeRt         DNSSummaryDNSSECParamsQueryType = "RT"
	DNSSummaryDNSSECParamsQueryTypeSig        DNSSummaryDNSSECParamsQueryType = "SIG"
	DNSSummaryDNSSECParamsQueryTypeSink       DNSSummaryDNSSECParamsQueryType = "SINK"
	DNSSummaryDNSSECParamsQueryTypeSMIMEA     DNSSummaryDNSSECParamsQueryType = "SMIMEA"
	DNSSummaryDNSSECParamsQueryTypeSOA        DNSSummaryDNSSECParamsQueryType = "SOA"
	DNSSummaryDNSSECParamsQueryTypeSPF        DNSSummaryDNSSECParamsQueryType = "SPF"
	DNSSummaryDNSSECParamsQueryTypeSRV        DNSSummaryDNSSECParamsQueryType = "SRV"
	DNSSummaryDNSSECParamsQueryTypeSSHFP      DNSSummaryDNSSECParamsQueryType = "SSHFP"
	DNSSummaryDNSSECParamsQueryTypeSVCB       DNSSummaryDNSSECParamsQueryType = "SVCB"
	DNSSummaryDNSSECParamsQueryTypeTa         DNSSummaryDNSSECParamsQueryType = "TA"
	DNSSummaryDNSSECParamsQueryTypeTalink     DNSSummaryDNSSECParamsQueryType = "TALINK"
	DNSSummaryDNSSECParamsQueryTypeTkey       DNSSummaryDNSSECParamsQueryType = "TKEY"
	DNSSummaryDNSSECParamsQueryTypeTLSA       DNSSummaryDNSSECParamsQueryType = "TLSA"
	DNSSummaryDNSSECParamsQueryTypeTSIG       DNSSummaryDNSSECParamsQueryType = "TSIG"
	DNSSummaryDNSSECParamsQueryTypeTXT        DNSSummaryDNSSECParamsQueryType = "TXT"
	DNSSummaryDNSSECParamsQueryTypeUinfo      DNSSummaryDNSSECParamsQueryType = "UINFO"
	DNSSummaryDNSSECParamsQueryTypeUID        DNSSummaryDNSSECParamsQueryType = "UID"
	DNSSummaryDNSSECParamsQueryTypeUnspec     DNSSummaryDNSSECParamsQueryType = "UNSPEC"
	DNSSummaryDNSSECParamsQueryTypeURI        DNSSummaryDNSSECParamsQueryType = "URI"
	DNSSummaryDNSSECParamsQueryTypeWks        DNSSummaryDNSSECParamsQueryType = "WKS"
	DNSSummaryDNSSECParamsQueryTypeX25        DNSSummaryDNSSECParamsQueryType = "X25"
	DNSSummaryDNSSECParamsQueryTypeZonemd     DNSSummaryDNSSECParamsQueryType = "ZONEMD"
)

func (r DNSSummaryDNSSECParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECParamsQueryTypeA, DNSSummaryDNSSECParamsQueryTypeAAAA, DNSSummaryDNSSECParamsQueryTypeA6, DNSSummaryDNSSECParamsQueryTypeAfsdb, DNSSummaryDNSSECParamsQueryTypeAny, DNSSummaryDNSSECParamsQueryTypeApl, DNSSummaryDNSSECParamsQueryTypeAtma, DNSSummaryDNSSECParamsQueryTypeAXFR, DNSSummaryDNSSECParamsQueryTypeCAA, DNSSummaryDNSSECParamsQueryTypeCdnskey, DNSSummaryDNSSECParamsQueryTypeCds, DNSSummaryDNSSECParamsQueryTypeCERT, DNSSummaryDNSSECParamsQueryTypeCNAME, DNSSummaryDNSSECParamsQueryTypeCsync, DNSSummaryDNSSECParamsQueryTypeDhcid, DNSSummaryDNSSECParamsQueryTypeDlv, DNSSummaryDNSSECParamsQueryTypeDname, DNSSummaryDNSSECParamsQueryTypeDNSKEY, DNSSummaryDNSSECParamsQueryTypeDoa, DNSSummaryDNSSECParamsQueryTypeDS, DNSSummaryDNSSECParamsQueryTypeEid, DNSSummaryDNSSECParamsQueryTypeEui48, DNSSummaryDNSSECParamsQueryTypeEui64, DNSSummaryDNSSECParamsQueryTypeGpos, DNSSummaryDNSSECParamsQueryTypeGid, DNSSummaryDNSSECParamsQueryTypeHinfo, DNSSummaryDNSSECParamsQueryTypeHip, DNSSummaryDNSSECParamsQueryTypeHTTPS, DNSSummaryDNSSECParamsQueryTypeIpseckey, DNSSummaryDNSSECParamsQueryTypeIsdn, DNSSummaryDNSSECParamsQueryTypeIxfr, DNSSummaryDNSSECParamsQueryTypeKey, DNSSummaryDNSSECParamsQueryTypeKx, DNSSummaryDNSSECParamsQueryTypeL32, DNSSummaryDNSSECParamsQueryTypeL64, DNSSummaryDNSSECParamsQueryTypeLOC, DNSSummaryDNSSECParamsQueryTypeLp, DNSSummaryDNSSECParamsQueryTypeMaila, DNSSummaryDNSSECParamsQueryTypeMailb, DNSSummaryDNSSECParamsQueryTypeMB, DNSSummaryDNSSECParamsQueryTypeMd, DNSSummaryDNSSECParamsQueryTypeMf, DNSSummaryDNSSECParamsQueryTypeMg, DNSSummaryDNSSECParamsQueryTypeMinfo, DNSSummaryDNSSECParamsQueryTypeMr, DNSSummaryDNSSECParamsQueryTypeMX, DNSSummaryDNSSECParamsQueryTypeNAPTR, DNSSummaryDNSSECParamsQueryTypeNb, DNSSummaryDNSSECParamsQueryTypeNbstat, DNSSummaryDNSSECParamsQueryTypeNid, DNSSummaryDNSSECParamsQueryTypeNimloc, DNSSummaryDNSSECParamsQueryTypeNinfo, DNSSummaryDNSSECParamsQueryTypeNS, DNSSummaryDNSSECParamsQueryTypeNsap, DNSSummaryDNSSECParamsQueryTypeNsec, DNSSummaryDNSSECParamsQueryTypeNsec3, DNSSummaryDNSSECParamsQueryTypeNsec3Param, DNSSummaryDNSSECParamsQueryTypeNull, DNSSummaryDNSSECParamsQueryTypeNxt, DNSSummaryDNSSECParamsQueryTypeOpenpgpkey, DNSSummaryDNSSECParamsQueryTypeOpt, DNSSummaryDNSSECParamsQueryTypePTR, DNSSummaryDNSSECParamsQueryTypePx, DNSSummaryDNSSECParamsQueryTypeRkey, DNSSummaryDNSSECParamsQueryTypeRp, DNSSummaryDNSSECParamsQueryTypeRrsig, DNSSummaryDNSSECParamsQueryTypeRt, DNSSummaryDNSSECParamsQueryTypeSig, DNSSummaryDNSSECParamsQueryTypeSink, DNSSummaryDNSSECParamsQueryTypeSMIMEA, DNSSummaryDNSSECParamsQueryTypeSOA, DNSSummaryDNSSECParamsQueryTypeSPF, DNSSummaryDNSSECParamsQueryTypeSRV, DNSSummaryDNSSECParamsQueryTypeSSHFP, DNSSummaryDNSSECParamsQueryTypeSVCB, DNSSummaryDNSSECParamsQueryTypeTa, DNSSummaryDNSSECParamsQueryTypeTalink, DNSSummaryDNSSECParamsQueryTypeTkey, DNSSummaryDNSSECParamsQueryTypeTLSA, DNSSummaryDNSSECParamsQueryTypeTSIG, DNSSummaryDNSSECParamsQueryTypeTXT, DNSSummaryDNSSECParamsQueryTypeUinfo, DNSSummaryDNSSECParamsQueryTypeUID, DNSSummaryDNSSECParamsQueryTypeUnspec, DNSSummaryDNSSECParamsQueryTypeURI, DNSSummaryDNSSECParamsQueryTypeWks, DNSSummaryDNSSECParamsQueryTypeX25, DNSSummaryDNSSECParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryDNSSECParamsResponseCode string

const (
	DNSSummaryDNSSECParamsResponseCodeNoerror   DNSSummaryDNSSECParamsResponseCode = "NOERROR"
	DNSSummaryDNSSECParamsResponseCodeFormerr   DNSSummaryDNSSECParamsResponseCode = "FORMERR"
	DNSSummaryDNSSECParamsResponseCodeServfail  DNSSummaryDNSSECParamsResponseCode = "SERVFAIL"
	DNSSummaryDNSSECParamsResponseCodeNxdomain  DNSSummaryDNSSECParamsResponseCode = "NXDOMAIN"
	DNSSummaryDNSSECParamsResponseCodeNotimp    DNSSummaryDNSSECParamsResponseCode = "NOTIMP"
	DNSSummaryDNSSECParamsResponseCodeRefused   DNSSummaryDNSSECParamsResponseCode = "REFUSED"
	DNSSummaryDNSSECParamsResponseCodeYxdomain  DNSSummaryDNSSECParamsResponseCode = "YXDOMAIN"
	DNSSummaryDNSSECParamsResponseCodeYxrrset   DNSSummaryDNSSECParamsResponseCode = "YXRRSET"
	DNSSummaryDNSSECParamsResponseCodeNxrrset   DNSSummaryDNSSECParamsResponseCode = "NXRRSET"
	DNSSummaryDNSSECParamsResponseCodeNotauth   DNSSummaryDNSSECParamsResponseCode = "NOTAUTH"
	DNSSummaryDNSSECParamsResponseCodeNotzone   DNSSummaryDNSSECParamsResponseCode = "NOTZONE"
	DNSSummaryDNSSECParamsResponseCodeBadsig    DNSSummaryDNSSECParamsResponseCode = "BADSIG"
	DNSSummaryDNSSECParamsResponseCodeBadkey    DNSSummaryDNSSECParamsResponseCode = "BADKEY"
	DNSSummaryDNSSECParamsResponseCodeBadtime   DNSSummaryDNSSECParamsResponseCode = "BADTIME"
	DNSSummaryDNSSECParamsResponseCodeBadmode   DNSSummaryDNSSECParamsResponseCode = "BADMODE"
	DNSSummaryDNSSECParamsResponseCodeBadname   DNSSummaryDNSSECParamsResponseCode = "BADNAME"
	DNSSummaryDNSSECParamsResponseCodeBadalg    DNSSummaryDNSSECParamsResponseCode = "BADALG"
	DNSSummaryDNSSECParamsResponseCodeBadtrunc  DNSSummaryDNSSECParamsResponseCode = "BADTRUNC"
	DNSSummaryDNSSECParamsResponseCodeBadcookie DNSSummaryDNSSECParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryDNSSECParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECParamsResponseCodeNoerror, DNSSummaryDNSSECParamsResponseCodeFormerr, DNSSummaryDNSSECParamsResponseCodeServfail, DNSSummaryDNSSECParamsResponseCodeNxdomain, DNSSummaryDNSSECParamsResponseCodeNotimp, DNSSummaryDNSSECParamsResponseCodeRefused, DNSSummaryDNSSECParamsResponseCodeYxdomain, DNSSummaryDNSSECParamsResponseCodeYxrrset, DNSSummaryDNSSECParamsResponseCodeNxrrset, DNSSummaryDNSSECParamsResponseCodeNotauth, DNSSummaryDNSSECParamsResponseCodeNotzone, DNSSummaryDNSSECParamsResponseCodeBadsig, DNSSummaryDNSSECParamsResponseCodeBadkey, DNSSummaryDNSSECParamsResponseCodeBadtime, DNSSummaryDNSSECParamsResponseCodeBadmode, DNSSummaryDNSSECParamsResponseCodeBadname, DNSSummaryDNSSECParamsResponseCodeBadalg, DNSSummaryDNSSECParamsResponseCodeBadtrunc, DNSSummaryDNSSECParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryDNSSECResponseEnvelope struct {
	Result  DNSSummaryDNSSECResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    dnsSummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseEnvelope]
type dnsSummaryDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECAwareParams struct {
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
	Format param.Field[DNSSummaryDNSSECAwareParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryDNSSECAwareParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryDNSSECAwareParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryDNSSECAwareParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECAwareParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryDNSSECAwareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryDNSSECAwareParamsFormat string

const (
	DNSSummaryDNSSECAwareParamsFormatJson DNSSummaryDNSSECAwareParamsFormat = "JSON"
	DNSSummaryDNSSECAwareParamsFormatCsv  DNSSummaryDNSSECAwareParamsFormat = "CSV"
)

func (r DNSSummaryDNSSECAwareParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECAwareParamsFormatJson, DNSSummaryDNSSECAwareParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryDNSSECAwareParamsProtocol string

const (
	DNSSummaryDNSSECAwareParamsProtocolUdp   DNSSummaryDNSSECAwareParamsProtocol = "UDP"
	DNSSummaryDNSSECAwareParamsProtocolTCP   DNSSummaryDNSSECAwareParamsProtocol = "TCP"
	DNSSummaryDNSSECAwareParamsProtocolHTTPS DNSSummaryDNSSECAwareParamsProtocol = "HTTPS"
	DNSSummaryDNSSECAwareParamsProtocolTLS   DNSSummaryDNSSECAwareParamsProtocol = "TLS"
)

func (r DNSSummaryDNSSECAwareParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECAwareParamsProtocolUdp, DNSSummaryDNSSECAwareParamsProtocolTCP, DNSSummaryDNSSECAwareParamsProtocolHTTPS, DNSSummaryDNSSECAwareParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryDNSSECAwareParamsQueryType string

const (
	DNSSummaryDNSSECAwareParamsQueryTypeA          DNSSummaryDNSSECAwareParamsQueryType = "A"
	DNSSummaryDNSSECAwareParamsQueryTypeAAAA       DNSSummaryDNSSECAwareParamsQueryType = "AAAA"
	DNSSummaryDNSSECAwareParamsQueryTypeA6         DNSSummaryDNSSECAwareParamsQueryType = "A6"
	DNSSummaryDNSSECAwareParamsQueryTypeAfsdb      DNSSummaryDNSSECAwareParamsQueryType = "AFSDB"
	DNSSummaryDNSSECAwareParamsQueryTypeAny        DNSSummaryDNSSECAwareParamsQueryType = "ANY"
	DNSSummaryDNSSECAwareParamsQueryTypeApl        DNSSummaryDNSSECAwareParamsQueryType = "APL"
	DNSSummaryDNSSECAwareParamsQueryTypeAtma       DNSSummaryDNSSECAwareParamsQueryType = "ATMA"
	DNSSummaryDNSSECAwareParamsQueryTypeAXFR       DNSSummaryDNSSECAwareParamsQueryType = "AXFR"
	DNSSummaryDNSSECAwareParamsQueryTypeCAA        DNSSummaryDNSSECAwareParamsQueryType = "CAA"
	DNSSummaryDNSSECAwareParamsQueryTypeCdnskey    DNSSummaryDNSSECAwareParamsQueryType = "CDNSKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeCds        DNSSummaryDNSSECAwareParamsQueryType = "CDS"
	DNSSummaryDNSSECAwareParamsQueryTypeCERT       DNSSummaryDNSSECAwareParamsQueryType = "CERT"
	DNSSummaryDNSSECAwareParamsQueryTypeCNAME      DNSSummaryDNSSECAwareParamsQueryType = "CNAME"
	DNSSummaryDNSSECAwareParamsQueryTypeCsync      DNSSummaryDNSSECAwareParamsQueryType = "CSYNC"
	DNSSummaryDNSSECAwareParamsQueryTypeDhcid      DNSSummaryDNSSECAwareParamsQueryType = "DHCID"
	DNSSummaryDNSSECAwareParamsQueryTypeDlv        DNSSummaryDNSSECAwareParamsQueryType = "DLV"
	DNSSummaryDNSSECAwareParamsQueryTypeDname      DNSSummaryDNSSECAwareParamsQueryType = "DNAME"
	DNSSummaryDNSSECAwareParamsQueryTypeDNSKEY     DNSSummaryDNSSECAwareParamsQueryType = "DNSKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeDoa        DNSSummaryDNSSECAwareParamsQueryType = "DOA"
	DNSSummaryDNSSECAwareParamsQueryTypeDS         DNSSummaryDNSSECAwareParamsQueryType = "DS"
	DNSSummaryDNSSECAwareParamsQueryTypeEid        DNSSummaryDNSSECAwareParamsQueryType = "EID"
	DNSSummaryDNSSECAwareParamsQueryTypeEui48      DNSSummaryDNSSECAwareParamsQueryType = "EUI48"
	DNSSummaryDNSSECAwareParamsQueryTypeEui64      DNSSummaryDNSSECAwareParamsQueryType = "EUI64"
	DNSSummaryDNSSECAwareParamsQueryTypeGpos       DNSSummaryDNSSECAwareParamsQueryType = "GPOS"
	DNSSummaryDNSSECAwareParamsQueryTypeGid        DNSSummaryDNSSECAwareParamsQueryType = "GID"
	DNSSummaryDNSSECAwareParamsQueryTypeHinfo      DNSSummaryDNSSECAwareParamsQueryType = "HINFO"
	DNSSummaryDNSSECAwareParamsQueryTypeHip        DNSSummaryDNSSECAwareParamsQueryType = "HIP"
	DNSSummaryDNSSECAwareParamsQueryTypeHTTPS      DNSSummaryDNSSECAwareParamsQueryType = "HTTPS"
	DNSSummaryDNSSECAwareParamsQueryTypeIpseckey   DNSSummaryDNSSECAwareParamsQueryType = "IPSECKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeIsdn       DNSSummaryDNSSECAwareParamsQueryType = "ISDN"
	DNSSummaryDNSSECAwareParamsQueryTypeIxfr       DNSSummaryDNSSECAwareParamsQueryType = "IXFR"
	DNSSummaryDNSSECAwareParamsQueryTypeKey        DNSSummaryDNSSECAwareParamsQueryType = "KEY"
	DNSSummaryDNSSECAwareParamsQueryTypeKx         DNSSummaryDNSSECAwareParamsQueryType = "KX"
	DNSSummaryDNSSECAwareParamsQueryTypeL32        DNSSummaryDNSSECAwareParamsQueryType = "L32"
	DNSSummaryDNSSECAwareParamsQueryTypeL64        DNSSummaryDNSSECAwareParamsQueryType = "L64"
	DNSSummaryDNSSECAwareParamsQueryTypeLOC        DNSSummaryDNSSECAwareParamsQueryType = "LOC"
	DNSSummaryDNSSECAwareParamsQueryTypeLp         DNSSummaryDNSSECAwareParamsQueryType = "LP"
	DNSSummaryDNSSECAwareParamsQueryTypeMaila      DNSSummaryDNSSECAwareParamsQueryType = "MAILA"
	DNSSummaryDNSSECAwareParamsQueryTypeMailb      DNSSummaryDNSSECAwareParamsQueryType = "MAILB"
	DNSSummaryDNSSECAwareParamsQueryTypeMB         DNSSummaryDNSSECAwareParamsQueryType = "MB"
	DNSSummaryDNSSECAwareParamsQueryTypeMd         DNSSummaryDNSSECAwareParamsQueryType = "MD"
	DNSSummaryDNSSECAwareParamsQueryTypeMf         DNSSummaryDNSSECAwareParamsQueryType = "MF"
	DNSSummaryDNSSECAwareParamsQueryTypeMg         DNSSummaryDNSSECAwareParamsQueryType = "MG"
	DNSSummaryDNSSECAwareParamsQueryTypeMinfo      DNSSummaryDNSSECAwareParamsQueryType = "MINFO"
	DNSSummaryDNSSECAwareParamsQueryTypeMr         DNSSummaryDNSSECAwareParamsQueryType = "MR"
	DNSSummaryDNSSECAwareParamsQueryTypeMX         DNSSummaryDNSSECAwareParamsQueryType = "MX"
	DNSSummaryDNSSECAwareParamsQueryTypeNAPTR      DNSSummaryDNSSECAwareParamsQueryType = "NAPTR"
	DNSSummaryDNSSECAwareParamsQueryTypeNb         DNSSummaryDNSSECAwareParamsQueryType = "NB"
	DNSSummaryDNSSECAwareParamsQueryTypeNbstat     DNSSummaryDNSSECAwareParamsQueryType = "NBSTAT"
	DNSSummaryDNSSECAwareParamsQueryTypeNid        DNSSummaryDNSSECAwareParamsQueryType = "NID"
	DNSSummaryDNSSECAwareParamsQueryTypeNimloc     DNSSummaryDNSSECAwareParamsQueryType = "NIMLOC"
	DNSSummaryDNSSECAwareParamsQueryTypeNinfo      DNSSummaryDNSSECAwareParamsQueryType = "NINFO"
	DNSSummaryDNSSECAwareParamsQueryTypeNS         DNSSummaryDNSSECAwareParamsQueryType = "NS"
	DNSSummaryDNSSECAwareParamsQueryTypeNsap       DNSSummaryDNSSECAwareParamsQueryType = "NSAP"
	DNSSummaryDNSSECAwareParamsQueryTypeNsec       DNSSummaryDNSSECAwareParamsQueryType = "NSEC"
	DNSSummaryDNSSECAwareParamsQueryTypeNsec3      DNSSummaryDNSSECAwareParamsQueryType = "NSEC3"
	DNSSummaryDNSSECAwareParamsQueryTypeNsec3Param DNSSummaryDNSSECAwareParamsQueryType = "NSEC3PARAM"
	DNSSummaryDNSSECAwareParamsQueryTypeNull       DNSSummaryDNSSECAwareParamsQueryType = "NULL"
	DNSSummaryDNSSECAwareParamsQueryTypeNxt        DNSSummaryDNSSECAwareParamsQueryType = "NXT"
	DNSSummaryDNSSECAwareParamsQueryTypeOpenpgpkey DNSSummaryDNSSECAwareParamsQueryType = "OPENPGPKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeOpt        DNSSummaryDNSSECAwareParamsQueryType = "OPT"
	DNSSummaryDNSSECAwareParamsQueryTypePTR        DNSSummaryDNSSECAwareParamsQueryType = "PTR"
	DNSSummaryDNSSECAwareParamsQueryTypePx         DNSSummaryDNSSECAwareParamsQueryType = "PX"
	DNSSummaryDNSSECAwareParamsQueryTypeRkey       DNSSummaryDNSSECAwareParamsQueryType = "RKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeRp         DNSSummaryDNSSECAwareParamsQueryType = "RP"
	DNSSummaryDNSSECAwareParamsQueryTypeRrsig      DNSSummaryDNSSECAwareParamsQueryType = "RRSIG"
	DNSSummaryDNSSECAwareParamsQueryTypeRt         DNSSummaryDNSSECAwareParamsQueryType = "RT"
	DNSSummaryDNSSECAwareParamsQueryTypeSig        DNSSummaryDNSSECAwareParamsQueryType = "SIG"
	DNSSummaryDNSSECAwareParamsQueryTypeSink       DNSSummaryDNSSECAwareParamsQueryType = "SINK"
	DNSSummaryDNSSECAwareParamsQueryTypeSMIMEA     DNSSummaryDNSSECAwareParamsQueryType = "SMIMEA"
	DNSSummaryDNSSECAwareParamsQueryTypeSOA        DNSSummaryDNSSECAwareParamsQueryType = "SOA"
	DNSSummaryDNSSECAwareParamsQueryTypeSPF        DNSSummaryDNSSECAwareParamsQueryType = "SPF"
	DNSSummaryDNSSECAwareParamsQueryTypeSRV        DNSSummaryDNSSECAwareParamsQueryType = "SRV"
	DNSSummaryDNSSECAwareParamsQueryTypeSSHFP      DNSSummaryDNSSECAwareParamsQueryType = "SSHFP"
	DNSSummaryDNSSECAwareParamsQueryTypeSVCB       DNSSummaryDNSSECAwareParamsQueryType = "SVCB"
	DNSSummaryDNSSECAwareParamsQueryTypeTa         DNSSummaryDNSSECAwareParamsQueryType = "TA"
	DNSSummaryDNSSECAwareParamsQueryTypeTalink     DNSSummaryDNSSECAwareParamsQueryType = "TALINK"
	DNSSummaryDNSSECAwareParamsQueryTypeTkey       DNSSummaryDNSSECAwareParamsQueryType = "TKEY"
	DNSSummaryDNSSECAwareParamsQueryTypeTLSA       DNSSummaryDNSSECAwareParamsQueryType = "TLSA"
	DNSSummaryDNSSECAwareParamsQueryTypeTSIG       DNSSummaryDNSSECAwareParamsQueryType = "TSIG"
	DNSSummaryDNSSECAwareParamsQueryTypeTXT        DNSSummaryDNSSECAwareParamsQueryType = "TXT"
	DNSSummaryDNSSECAwareParamsQueryTypeUinfo      DNSSummaryDNSSECAwareParamsQueryType = "UINFO"
	DNSSummaryDNSSECAwareParamsQueryTypeUID        DNSSummaryDNSSECAwareParamsQueryType = "UID"
	DNSSummaryDNSSECAwareParamsQueryTypeUnspec     DNSSummaryDNSSECAwareParamsQueryType = "UNSPEC"
	DNSSummaryDNSSECAwareParamsQueryTypeURI        DNSSummaryDNSSECAwareParamsQueryType = "URI"
	DNSSummaryDNSSECAwareParamsQueryTypeWks        DNSSummaryDNSSECAwareParamsQueryType = "WKS"
	DNSSummaryDNSSECAwareParamsQueryTypeX25        DNSSummaryDNSSECAwareParamsQueryType = "X25"
	DNSSummaryDNSSECAwareParamsQueryTypeZonemd     DNSSummaryDNSSECAwareParamsQueryType = "ZONEMD"
)

func (r DNSSummaryDNSSECAwareParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECAwareParamsQueryTypeA, DNSSummaryDNSSECAwareParamsQueryTypeAAAA, DNSSummaryDNSSECAwareParamsQueryTypeA6, DNSSummaryDNSSECAwareParamsQueryTypeAfsdb, DNSSummaryDNSSECAwareParamsQueryTypeAny, DNSSummaryDNSSECAwareParamsQueryTypeApl, DNSSummaryDNSSECAwareParamsQueryTypeAtma, DNSSummaryDNSSECAwareParamsQueryTypeAXFR, DNSSummaryDNSSECAwareParamsQueryTypeCAA, DNSSummaryDNSSECAwareParamsQueryTypeCdnskey, DNSSummaryDNSSECAwareParamsQueryTypeCds, DNSSummaryDNSSECAwareParamsQueryTypeCERT, DNSSummaryDNSSECAwareParamsQueryTypeCNAME, DNSSummaryDNSSECAwareParamsQueryTypeCsync, DNSSummaryDNSSECAwareParamsQueryTypeDhcid, DNSSummaryDNSSECAwareParamsQueryTypeDlv, DNSSummaryDNSSECAwareParamsQueryTypeDname, DNSSummaryDNSSECAwareParamsQueryTypeDNSKEY, DNSSummaryDNSSECAwareParamsQueryTypeDoa, DNSSummaryDNSSECAwareParamsQueryTypeDS, DNSSummaryDNSSECAwareParamsQueryTypeEid, DNSSummaryDNSSECAwareParamsQueryTypeEui48, DNSSummaryDNSSECAwareParamsQueryTypeEui64, DNSSummaryDNSSECAwareParamsQueryTypeGpos, DNSSummaryDNSSECAwareParamsQueryTypeGid, DNSSummaryDNSSECAwareParamsQueryTypeHinfo, DNSSummaryDNSSECAwareParamsQueryTypeHip, DNSSummaryDNSSECAwareParamsQueryTypeHTTPS, DNSSummaryDNSSECAwareParamsQueryTypeIpseckey, DNSSummaryDNSSECAwareParamsQueryTypeIsdn, DNSSummaryDNSSECAwareParamsQueryTypeIxfr, DNSSummaryDNSSECAwareParamsQueryTypeKey, DNSSummaryDNSSECAwareParamsQueryTypeKx, DNSSummaryDNSSECAwareParamsQueryTypeL32, DNSSummaryDNSSECAwareParamsQueryTypeL64, DNSSummaryDNSSECAwareParamsQueryTypeLOC, DNSSummaryDNSSECAwareParamsQueryTypeLp, DNSSummaryDNSSECAwareParamsQueryTypeMaila, DNSSummaryDNSSECAwareParamsQueryTypeMailb, DNSSummaryDNSSECAwareParamsQueryTypeMB, DNSSummaryDNSSECAwareParamsQueryTypeMd, DNSSummaryDNSSECAwareParamsQueryTypeMf, DNSSummaryDNSSECAwareParamsQueryTypeMg, DNSSummaryDNSSECAwareParamsQueryTypeMinfo, DNSSummaryDNSSECAwareParamsQueryTypeMr, DNSSummaryDNSSECAwareParamsQueryTypeMX, DNSSummaryDNSSECAwareParamsQueryTypeNAPTR, DNSSummaryDNSSECAwareParamsQueryTypeNb, DNSSummaryDNSSECAwareParamsQueryTypeNbstat, DNSSummaryDNSSECAwareParamsQueryTypeNid, DNSSummaryDNSSECAwareParamsQueryTypeNimloc, DNSSummaryDNSSECAwareParamsQueryTypeNinfo, DNSSummaryDNSSECAwareParamsQueryTypeNS, DNSSummaryDNSSECAwareParamsQueryTypeNsap, DNSSummaryDNSSECAwareParamsQueryTypeNsec, DNSSummaryDNSSECAwareParamsQueryTypeNsec3, DNSSummaryDNSSECAwareParamsQueryTypeNsec3Param, DNSSummaryDNSSECAwareParamsQueryTypeNull, DNSSummaryDNSSECAwareParamsQueryTypeNxt, DNSSummaryDNSSECAwareParamsQueryTypeOpenpgpkey, DNSSummaryDNSSECAwareParamsQueryTypeOpt, DNSSummaryDNSSECAwareParamsQueryTypePTR, DNSSummaryDNSSECAwareParamsQueryTypePx, DNSSummaryDNSSECAwareParamsQueryTypeRkey, DNSSummaryDNSSECAwareParamsQueryTypeRp, DNSSummaryDNSSECAwareParamsQueryTypeRrsig, DNSSummaryDNSSECAwareParamsQueryTypeRt, DNSSummaryDNSSECAwareParamsQueryTypeSig, DNSSummaryDNSSECAwareParamsQueryTypeSink, DNSSummaryDNSSECAwareParamsQueryTypeSMIMEA, DNSSummaryDNSSECAwareParamsQueryTypeSOA, DNSSummaryDNSSECAwareParamsQueryTypeSPF, DNSSummaryDNSSECAwareParamsQueryTypeSRV, DNSSummaryDNSSECAwareParamsQueryTypeSSHFP, DNSSummaryDNSSECAwareParamsQueryTypeSVCB, DNSSummaryDNSSECAwareParamsQueryTypeTa, DNSSummaryDNSSECAwareParamsQueryTypeTalink, DNSSummaryDNSSECAwareParamsQueryTypeTkey, DNSSummaryDNSSECAwareParamsQueryTypeTLSA, DNSSummaryDNSSECAwareParamsQueryTypeTSIG, DNSSummaryDNSSECAwareParamsQueryTypeTXT, DNSSummaryDNSSECAwareParamsQueryTypeUinfo, DNSSummaryDNSSECAwareParamsQueryTypeUID, DNSSummaryDNSSECAwareParamsQueryTypeUnspec, DNSSummaryDNSSECAwareParamsQueryTypeURI, DNSSummaryDNSSECAwareParamsQueryTypeWks, DNSSummaryDNSSECAwareParamsQueryTypeX25, DNSSummaryDNSSECAwareParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryDNSSECAwareParamsResponseCode string

const (
	DNSSummaryDNSSECAwareParamsResponseCodeNoerror   DNSSummaryDNSSECAwareParamsResponseCode = "NOERROR"
	DNSSummaryDNSSECAwareParamsResponseCodeFormerr   DNSSummaryDNSSECAwareParamsResponseCode = "FORMERR"
	DNSSummaryDNSSECAwareParamsResponseCodeServfail  DNSSummaryDNSSECAwareParamsResponseCode = "SERVFAIL"
	DNSSummaryDNSSECAwareParamsResponseCodeNxdomain  DNSSummaryDNSSECAwareParamsResponseCode = "NXDOMAIN"
	DNSSummaryDNSSECAwareParamsResponseCodeNotimp    DNSSummaryDNSSECAwareParamsResponseCode = "NOTIMP"
	DNSSummaryDNSSECAwareParamsResponseCodeRefused   DNSSummaryDNSSECAwareParamsResponseCode = "REFUSED"
	DNSSummaryDNSSECAwareParamsResponseCodeYxdomain  DNSSummaryDNSSECAwareParamsResponseCode = "YXDOMAIN"
	DNSSummaryDNSSECAwareParamsResponseCodeYxrrset   DNSSummaryDNSSECAwareParamsResponseCode = "YXRRSET"
	DNSSummaryDNSSECAwareParamsResponseCodeNxrrset   DNSSummaryDNSSECAwareParamsResponseCode = "NXRRSET"
	DNSSummaryDNSSECAwareParamsResponseCodeNotauth   DNSSummaryDNSSECAwareParamsResponseCode = "NOTAUTH"
	DNSSummaryDNSSECAwareParamsResponseCodeNotzone   DNSSummaryDNSSECAwareParamsResponseCode = "NOTZONE"
	DNSSummaryDNSSECAwareParamsResponseCodeBadsig    DNSSummaryDNSSECAwareParamsResponseCode = "BADSIG"
	DNSSummaryDNSSECAwareParamsResponseCodeBadkey    DNSSummaryDNSSECAwareParamsResponseCode = "BADKEY"
	DNSSummaryDNSSECAwareParamsResponseCodeBadtime   DNSSummaryDNSSECAwareParamsResponseCode = "BADTIME"
	DNSSummaryDNSSECAwareParamsResponseCodeBadmode   DNSSummaryDNSSECAwareParamsResponseCode = "BADMODE"
	DNSSummaryDNSSECAwareParamsResponseCodeBadname   DNSSummaryDNSSECAwareParamsResponseCode = "BADNAME"
	DNSSummaryDNSSECAwareParamsResponseCodeBadalg    DNSSummaryDNSSECAwareParamsResponseCode = "BADALG"
	DNSSummaryDNSSECAwareParamsResponseCodeBadtrunc  DNSSummaryDNSSECAwareParamsResponseCode = "BADTRUNC"
	DNSSummaryDNSSECAwareParamsResponseCodeBadcookie DNSSummaryDNSSECAwareParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryDNSSECAwareParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryDNSSECAwareParamsResponseCodeNoerror, DNSSummaryDNSSECAwareParamsResponseCodeFormerr, DNSSummaryDNSSECAwareParamsResponseCodeServfail, DNSSummaryDNSSECAwareParamsResponseCodeNxdomain, DNSSummaryDNSSECAwareParamsResponseCodeNotimp, DNSSummaryDNSSECAwareParamsResponseCodeRefused, DNSSummaryDNSSECAwareParamsResponseCodeYxdomain, DNSSummaryDNSSECAwareParamsResponseCodeYxrrset, DNSSummaryDNSSECAwareParamsResponseCodeNxrrset, DNSSummaryDNSSECAwareParamsResponseCodeNotauth, DNSSummaryDNSSECAwareParamsResponseCodeNotzone, DNSSummaryDNSSECAwareParamsResponseCodeBadsig, DNSSummaryDNSSECAwareParamsResponseCodeBadkey, DNSSummaryDNSSECAwareParamsResponseCodeBadtime, DNSSummaryDNSSECAwareParamsResponseCodeBadmode, DNSSummaryDNSSECAwareParamsResponseCodeBadname, DNSSummaryDNSSECAwareParamsResponseCodeBadalg, DNSSummaryDNSSECAwareParamsResponseCodeBadtrunc, DNSSummaryDNSSECAwareParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryDNSSECAwareResponseEnvelope struct {
	Result  DNSSummaryDNSSECAwareResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    dnsSummaryDNSSECAwareResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECAwareResponseEnvelope]
type dnsSummaryDNSSECAwareResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryDNSSECE2EParams struct {
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
	Format param.Field[DNSSummaryDnssece2EParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryDnssece2EParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryDnssece2EParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryDnssece2EParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECE2EParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryDNSSECE2EParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryDnssece2EParamsFormat string

const (
	DNSSummaryDnssece2EParamsFormatJson DNSSummaryDnssece2EParamsFormat = "JSON"
	DNSSummaryDnssece2EParamsFormatCsv  DNSSummaryDnssece2EParamsFormat = "CSV"
)

func (r DNSSummaryDnssece2EParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryDnssece2EParamsFormatJson, DNSSummaryDnssece2EParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryDnssece2EParamsProtocol string

const (
	DNSSummaryDnssece2EParamsProtocolUdp   DNSSummaryDnssece2EParamsProtocol = "UDP"
	DNSSummaryDnssece2EParamsProtocolTCP   DNSSummaryDnssece2EParamsProtocol = "TCP"
	DNSSummaryDnssece2EParamsProtocolHTTPS DNSSummaryDnssece2EParamsProtocol = "HTTPS"
	DNSSummaryDnssece2EParamsProtocolTLS   DNSSummaryDnssece2EParamsProtocol = "TLS"
)

func (r DNSSummaryDnssece2EParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryDnssece2EParamsProtocolUdp, DNSSummaryDnssece2EParamsProtocolTCP, DNSSummaryDnssece2EParamsProtocolHTTPS, DNSSummaryDnssece2EParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryDnssece2EParamsQueryType string

const (
	DNSSummaryDnssece2EParamsQueryTypeA          DNSSummaryDnssece2EParamsQueryType = "A"
	DNSSummaryDnssece2EParamsQueryTypeAAAA       DNSSummaryDnssece2EParamsQueryType = "AAAA"
	DNSSummaryDnssece2EParamsQueryTypeA6         DNSSummaryDnssece2EParamsQueryType = "A6"
	DNSSummaryDnssece2EParamsQueryTypeAfsdb      DNSSummaryDnssece2EParamsQueryType = "AFSDB"
	DNSSummaryDnssece2EParamsQueryTypeAny        DNSSummaryDnssece2EParamsQueryType = "ANY"
	DNSSummaryDnssece2EParamsQueryTypeApl        DNSSummaryDnssece2EParamsQueryType = "APL"
	DNSSummaryDnssece2EParamsQueryTypeAtma       DNSSummaryDnssece2EParamsQueryType = "ATMA"
	DNSSummaryDnssece2EParamsQueryTypeAXFR       DNSSummaryDnssece2EParamsQueryType = "AXFR"
	DNSSummaryDnssece2EParamsQueryTypeCAA        DNSSummaryDnssece2EParamsQueryType = "CAA"
	DNSSummaryDnssece2EParamsQueryTypeCdnskey    DNSSummaryDnssece2EParamsQueryType = "CDNSKEY"
	DNSSummaryDnssece2EParamsQueryTypeCds        DNSSummaryDnssece2EParamsQueryType = "CDS"
	DNSSummaryDnssece2EParamsQueryTypeCERT       DNSSummaryDnssece2EParamsQueryType = "CERT"
	DNSSummaryDnssece2EParamsQueryTypeCNAME      DNSSummaryDnssece2EParamsQueryType = "CNAME"
	DNSSummaryDnssece2EParamsQueryTypeCsync      DNSSummaryDnssece2EParamsQueryType = "CSYNC"
	DNSSummaryDnssece2EParamsQueryTypeDhcid      DNSSummaryDnssece2EParamsQueryType = "DHCID"
	DNSSummaryDnssece2EParamsQueryTypeDlv        DNSSummaryDnssece2EParamsQueryType = "DLV"
	DNSSummaryDnssece2EParamsQueryTypeDname      DNSSummaryDnssece2EParamsQueryType = "DNAME"
	DNSSummaryDnssece2EParamsQueryTypeDNSKEY     DNSSummaryDnssece2EParamsQueryType = "DNSKEY"
	DNSSummaryDnssece2EParamsQueryTypeDoa        DNSSummaryDnssece2EParamsQueryType = "DOA"
	DNSSummaryDnssece2EParamsQueryTypeDS         DNSSummaryDnssece2EParamsQueryType = "DS"
	DNSSummaryDnssece2EParamsQueryTypeEid        DNSSummaryDnssece2EParamsQueryType = "EID"
	DNSSummaryDnssece2EParamsQueryTypeEui48      DNSSummaryDnssece2EParamsQueryType = "EUI48"
	DNSSummaryDnssece2EParamsQueryTypeEui64      DNSSummaryDnssece2EParamsQueryType = "EUI64"
	DNSSummaryDnssece2EParamsQueryTypeGpos       DNSSummaryDnssece2EParamsQueryType = "GPOS"
	DNSSummaryDnssece2EParamsQueryTypeGid        DNSSummaryDnssece2EParamsQueryType = "GID"
	DNSSummaryDnssece2EParamsQueryTypeHinfo      DNSSummaryDnssece2EParamsQueryType = "HINFO"
	DNSSummaryDnssece2EParamsQueryTypeHip        DNSSummaryDnssece2EParamsQueryType = "HIP"
	DNSSummaryDnssece2EParamsQueryTypeHTTPS      DNSSummaryDnssece2EParamsQueryType = "HTTPS"
	DNSSummaryDnssece2EParamsQueryTypeIpseckey   DNSSummaryDnssece2EParamsQueryType = "IPSECKEY"
	DNSSummaryDnssece2EParamsQueryTypeIsdn       DNSSummaryDnssece2EParamsQueryType = "ISDN"
	DNSSummaryDnssece2EParamsQueryTypeIxfr       DNSSummaryDnssece2EParamsQueryType = "IXFR"
	DNSSummaryDnssece2EParamsQueryTypeKey        DNSSummaryDnssece2EParamsQueryType = "KEY"
	DNSSummaryDnssece2EParamsQueryTypeKx         DNSSummaryDnssece2EParamsQueryType = "KX"
	DNSSummaryDnssece2EParamsQueryTypeL32        DNSSummaryDnssece2EParamsQueryType = "L32"
	DNSSummaryDnssece2EParamsQueryTypeL64        DNSSummaryDnssece2EParamsQueryType = "L64"
	DNSSummaryDnssece2EParamsQueryTypeLOC        DNSSummaryDnssece2EParamsQueryType = "LOC"
	DNSSummaryDnssece2EParamsQueryTypeLp         DNSSummaryDnssece2EParamsQueryType = "LP"
	DNSSummaryDnssece2EParamsQueryTypeMaila      DNSSummaryDnssece2EParamsQueryType = "MAILA"
	DNSSummaryDnssece2EParamsQueryTypeMailb      DNSSummaryDnssece2EParamsQueryType = "MAILB"
	DNSSummaryDnssece2EParamsQueryTypeMB         DNSSummaryDnssece2EParamsQueryType = "MB"
	DNSSummaryDnssece2EParamsQueryTypeMd         DNSSummaryDnssece2EParamsQueryType = "MD"
	DNSSummaryDnssece2EParamsQueryTypeMf         DNSSummaryDnssece2EParamsQueryType = "MF"
	DNSSummaryDnssece2EParamsQueryTypeMg         DNSSummaryDnssece2EParamsQueryType = "MG"
	DNSSummaryDnssece2EParamsQueryTypeMinfo      DNSSummaryDnssece2EParamsQueryType = "MINFO"
	DNSSummaryDnssece2EParamsQueryTypeMr         DNSSummaryDnssece2EParamsQueryType = "MR"
	DNSSummaryDnssece2EParamsQueryTypeMX         DNSSummaryDnssece2EParamsQueryType = "MX"
	DNSSummaryDnssece2EParamsQueryTypeNAPTR      DNSSummaryDnssece2EParamsQueryType = "NAPTR"
	DNSSummaryDnssece2EParamsQueryTypeNb         DNSSummaryDnssece2EParamsQueryType = "NB"
	DNSSummaryDnssece2EParamsQueryTypeNbstat     DNSSummaryDnssece2EParamsQueryType = "NBSTAT"
	DNSSummaryDnssece2EParamsQueryTypeNid        DNSSummaryDnssece2EParamsQueryType = "NID"
	DNSSummaryDnssece2EParamsQueryTypeNimloc     DNSSummaryDnssece2EParamsQueryType = "NIMLOC"
	DNSSummaryDnssece2EParamsQueryTypeNinfo      DNSSummaryDnssece2EParamsQueryType = "NINFO"
	DNSSummaryDnssece2EParamsQueryTypeNS         DNSSummaryDnssece2EParamsQueryType = "NS"
	DNSSummaryDnssece2EParamsQueryTypeNsap       DNSSummaryDnssece2EParamsQueryType = "NSAP"
	DNSSummaryDnssece2EParamsQueryTypeNsec       DNSSummaryDnssece2EParamsQueryType = "NSEC"
	DNSSummaryDnssece2EParamsQueryTypeNsec3      DNSSummaryDnssece2EParamsQueryType = "NSEC3"
	DNSSummaryDnssece2EParamsQueryTypeNsec3Param DNSSummaryDnssece2EParamsQueryType = "NSEC3PARAM"
	DNSSummaryDnssece2EParamsQueryTypeNull       DNSSummaryDnssece2EParamsQueryType = "NULL"
	DNSSummaryDnssece2EParamsQueryTypeNxt        DNSSummaryDnssece2EParamsQueryType = "NXT"
	DNSSummaryDnssece2EParamsQueryTypeOpenpgpkey DNSSummaryDnssece2EParamsQueryType = "OPENPGPKEY"
	DNSSummaryDnssece2EParamsQueryTypeOpt        DNSSummaryDnssece2EParamsQueryType = "OPT"
	DNSSummaryDnssece2EParamsQueryTypePTR        DNSSummaryDnssece2EParamsQueryType = "PTR"
	DNSSummaryDnssece2EParamsQueryTypePx         DNSSummaryDnssece2EParamsQueryType = "PX"
	DNSSummaryDnssece2EParamsQueryTypeRkey       DNSSummaryDnssece2EParamsQueryType = "RKEY"
	DNSSummaryDnssece2EParamsQueryTypeRp         DNSSummaryDnssece2EParamsQueryType = "RP"
	DNSSummaryDnssece2EParamsQueryTypeRrsig      DNSSummaryDnssece2EParamsQueryType = "RRSIG"
	DNSSummaryDnssece2EParamsQueryTypeRt         DNSSummaryDnssece2EParamsQueryType = "RT"
	DNSSummaryDnssece2EParamsQueryTypeSig        DNSSummaryDnssece2EParamsQueryType = "SIG"
	DNSSummaryDnssece2EParamsQueryTypeSink       DNSSummaryDnssece2EParamsQueryType = "SINK"
	DNSSummaryDnssece2EParamsQueryTypeSMIMEA     DNSSummaryDnssece2EParamsQueryType = "SMIMEA"
	DNSSummaryDnssece2EParamsQueryTypeSOA        DNSSummaryDnssece2EParamsQueryType = "SOA"
	DNSSummaryDnssece2EParamsQueryTypeSPF        DNSSummaryDnssece2EParamsQueryType = "SPF"
	DNSSummaryDnssece2EParamsQueryTypeSRV        DNSSummaryDnssece2EParamsQueryType = "SRV"
	DNSSummaryDnssece2EParamsQueryTypeSSHFP      DNSSummaryDnssece2EParamsQueryType = "SSHFP"
	DNSSummaryDnssece2EParamsQueryTypeSVCB       DNSSummaryDnssece2EParamsQueryType = "SVCB"
	DNSSummaryDnssece2EParamsQueryTypeTa         DNSSummaryDnssece2EParamsQueryType = "TA"
	DNSSummaryDnssece2EParamsQueryTypeTalink     DNSSummaryDnssece2EParamsQueryType = "TALINK"
	DNSSummaryDnssece2EParamsQueryTypeTkey       DNSSummaryDnssece2EParamsQueryType = "TKEY"
	DNSSummaryDnssece2EParamsQueryTypeTLSA       DNSSummaryDnssece2EParamsQueryType = "TLSA"
	DNSSummaryDnssece2EParamsQueryTypeTSIG       DNSSummaryDnssece2EParamsQueryType = "TSIG"
	DNSSummaryDnssece2EParamsQueryTypeTXT        DNSSummaryDnssece2EParamsQueryType = "TXT"
	DNSSummaryDnssece2EParamsQueryTypeUinfo      DNSSummaryDnssece2EParamsQueryType = "UINFO"
	DNSSummaryDnssece2EParamsQueryTypeUID        DNSSummaryDnssece2EParamsQueryType = "UID"
	DNSSummaryDnssece2EParamsQueryTypeUnspec     DNSSummaryDnssece2EParamsQueryType = "UNSPEC"
	DNSSummaryDnssece2EParamsQueryTypeURI        DNSSummaryDnssece2EParamsQueryType = "URI"
	DNSSummaryDnssece2EParamsQueryTypeWks        DNSSummaryDnssece2EParamsQueryType = "WKS"
	DNSSummaryDnssece2EParamsQueryTypeX25        DNSSummaryDnssece2EParamsQueryType = "X25"
	DNSSummaryDnssece2EParamsQueryTypeZonemd     DNSSummaryDnssece2EParamsQueryType = "ZONEMD"
)

func (r DNSSummaryDnssece2EParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryDnssece2EParamsQueryTypeA, DNSSummaryDnssece2EParamsQueryTypeAAAA, DNSSummaryDnssece2EParamsQueryTypeA6, DNSSummaryDnssece2EParamsQueryTypeAfsdb, DNSSummaryDnssece2EParamsQueryTypeAny, DNSSummaryDnssece2EParamsQueryTypeApl, DNSSummaryDnssece2EParamsQueryTypeAtma, DNSSummaryDnssece2EParamsQueryTypeAXFR, DNSSummaryDnssece2EParamsQueryTypeCAA, DNSSummaryDnssece2EParamsQueryTypeCdnskey, DNSSummaryDnssece2EParamsQueryTypeCds, DNSSummaryDnssece2EParamsQueryTypeCERT, DNSSummaryDnssece2EParamsQueryTypeCNAME, DNSSummaryDnssece2EParamsQueryTypeCsync, DNSSummaryDnssece2EParamsQueryTypeDhcid, DNSSummaryDnssece2EParamsQueryTypeDlv, DNSSummaryDnssece2EParamsQueryTypeDname, DNSSummaryDnssece2EParamsQueryTypeDNSKEY, DNSSummaryDnssece2EParamsQueryTypeDoa, DNSSummaryDnssece2EParamsQueryTypeDS, DNSSummaryDnssece2EParamsQueryTypeEid, DNSSummaryDnssece2EParamsQueryTypeEui48, DNSSummaryDnssece2EParamsQueryTypeEui64, DNSSummaryDnssece2EParamsQueryTypeGpos, DNSSummaryDnssece2EParamsQueryTypeGid, DNSSummaryDnssece2EParamsQueryTypeHinfo, DNSSummaryDnssece2EParamsQueryTypeHip, DNSSummaryDnssece2EParamsQueryTypeHTTPS, DNSSummaryDnssece2EParamsQueryTypeIpseckey, DNSSummaryDnssece2EParamsQueryTypeIsdn, DNSSummaryDnssece2EParamsQueryTypeIxfr, DNSSummaryDnssece2EParamsQueryTypeKey, DNSSummaryDnssece2EParamsQueryTypeKx, DNSSummaryDnssece2EParamsQueryTypeL32, DNSSummaryDnssece2EParamsQueryTypeL64, DNSSummaryDnssece2EParamsQueryTypeLOC, DNSSummaryDnssece2EParamsQueryTypeLp, DNSSummaryDnssece2EParamsQueryTypeMaila, DNSSummaryDnssece2EParamsQueryTypeMailb, DNSSummaryDnssece2EParamsQueryTypeMB, DNSSummaryDnssece2EParamsQueryTypeMd, DNSSummaryDnssece2EParamsQueryTypeMf, DNSSummaryDnssece2EParamsQueryTypeMg, DNSSummaryDnssece2EParamsQueryTypeMinfo, DNSSummaryDnssece2EParamsQueryTypeMr, DNSSummaryDnssece2EParamsQueryTypeMX, DNSSummaryDnssece2EParamsQueryTypeNAPTR, DNSSummaryDnssece2EParamsQueryTypeNb, DNSSummaryDnssece2EParamsQueryTypeNbstat, DNSSummaryDnssece2EParamsQueryTypeNid, DNSSummaryDnssece2EParamsQueryTypeNimloc, DNSSummaryDnssece2EParamsQueryTypeNinfo, DNSSummaryDnssece2EParamsQueryTypeNS, DNSSummaryDnssece2EParamsQueryTypeNsap, DNSSummaryDnssece2EParamsQueryTypeNsec, DNSSummaryDnssece2EParamsQueryTypeNsec3, DNSSummaryDnssece2EParamsQueryTypeNsec3Param, DNSSummaryDnssece2EParamsQueryTypeNull, DNSSummaryDnssece2EParamsQueryTypeNxt, DNSSummaryDnssece2EParamsQueryTypeOpenpgpkey, DNSSummaryDnssece2EParamsQueryTypeOpt, DNSSummaryDnssece2EParamsQueryTypePTR, DNSSummaryDnssece2EParamsQueryTypePx, DNSSummaryDnssece2EParamsQueryTypeRkey, DNSSummaryDnssece2EParamsQueryTypeRp, DNSSummaryDnssece2EParamsQueryTypeRrsig, DNSSummaryDnssece2EParamsQueryTypeRt, DNSSummaryDnssece2EParamsQueryTypeSig, DNSSummaryDnssece2EParamsQueryTypeSink, DNSSummaryDnssece2EParamsQueryTypeSMIMEA, DNSSummaryDnssece2EParamsQueryTypeSOA, DNSSummaryDnssece2EParamsQueryTypeSPF, DNSSummaryDnssece2EParamsQueryTypeSRV, DNSSummaryDnssece2EParamsQueryTypeSSHFP, DNSSummaryDnssece2EParamsQueryTypeSVCB, DNSSummaryDnssece2EParamsQueryTypeTa, DNSSummaryDnssece2EParamsQueryTypeTalink, DNSSummaryDnssece2EParamsQueryTypeTkey, DNSSummaryDnssece2EParamsQueryTypeTLSA, DNSSummaryDnssece2EParamsQueryTypeTSIG, DNSSummaryDnssece2EParamsQueryTypeTXT, DNSSummaryDnssece2EParamsQueryTypeUinfo, DNSSummaryDnssece2EParamsQueryTypeUID, DNSSummaryDnssece2EParamsQueryTypeUnspec, DNSSummaryDnssece2EParamsQueryTypeURI, DNSSummaryDnssece2EParamsQueryTypeWks, DNSSummaryDnssece2EParamsQueryTypeX25, DNSSummaryDnssece2EParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryDnssece2EParamsResponseCode string

const (
	DNSSummaryDnssece2EParamsResponseCodeNoerror   DNSSummaryDnssece2EParamsResponseCode = "NOERROR"
	DNSSummaryDnssece2EParamsResponseCodeFormerr   DNSSummaryDnssece2EParamsResponseCode = "FORMERR"
	DNSSummaryDnssece2EParamsResponseCodeServfail  DNSSummaryDnssece2EParamsResponseCode = "SERVFAIL"
	DNSSummaryDnssece2EParamsResponseCodeNxdomain  DNSSummaryDnssece2EParamsResponseCode = "NXDOMAIN"
	DNSSummaryDnssece2EParamsResponseCodeNotimp    DNSSummaryDnssece2EParamsResponseCode = "NOTIMP"
	DNSSummaryDnssece2EParamsResponseCodeRefused   DNSSummaryDnssece2EParamsResponseCode = "REFUSED"
	DNSSummaryDnssece2EParamsResponseCodeYxdomain  DNSSummaryDnssece2EParamsResponseCode = "YXDOMAIN"
	DNSSummaryDnssece2EParamsResponseCodeYxrrset   DNSSummaryDnssece2EParamsResponseCode = "YXRRSET"
	DNSSummaryDnssece2EParamsResponseCodeNxrrset   DNSSummaryDnssece2EParamsResponseCode = "NXRRSET"
	DNSSummaryDnssece2EParamsResponseCodeNotauth   DNSSummaryDnssece2EParamsResponseCode = "NOTAUTH"
	DNSSummaryDnssece2EParamsResponseCodeNotzone   DNSSummaryDnssece2EParamsResponseCode = "NOTZONE"
	DNSSummaryDnssece2EParamsResponseCodeBadsig    DNSSummaryDnssece2EParamsResponseCode = "BADSIG"
	DNSSummaryDnssece2EParamsResponseCodeBadkey    DNSSummaryDnssece2EParamsResponseCode = "BADKEY"
	DNSSummaryDnssece2EParamsResponseCodeBadtime   DNSSummaryDnssece2EParamsResponseCode = "BADTIME"
	DNSSummaryDnssece2EParamsResponseCodeBadmode   DNSSummaryDnssece2EParamsResponseCode = "BADMODE"
	DNSSummaryDnssece2EParamsResponseCodeBadname   DNSSummaryDnssece2EParamsResponseCode = "BADNAME"
	DNSSummaryDnssece2EParamsResponseCodeBadalg    DNSSummaryDnssece2EParamsResponseCode = "BADALG"
	DNSSummaryDnssece2EParamsResponseCodeBadtrunc  DNSSummaryDnssece2EParamsResponseCode = "BADTRUNC"
	DNSSummaryDnssece2EParamsResponseCodeBadcookie DNSSummaryDnssece2EParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryDnssece2EParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryDnssece2EParamsResponseCodeNoerror, DNSSummaryDnssece2EParamsResponseCodeFormerr, DNSSummaryDnssece2EParamsResponseCodeServfail, DNSSummaryDnssece2EParamsResponseCodeNxdomain, DNSSummaryDnssece2EParamsResponseCodeNotimp, DNSSummaryDnssece2EParamsResponseCodeRefused, DNSSummaryDnssece2EParamsResponseCodeYxdomain, DNSSummaryDnssece2EParamsResponseCodeYxrrset, DNSSummaryDnssece2EParamsResponseCodeNxrrset, DNSSummaryDnssece2EParamsResponseCodeNotauth, DNSSummaryDnssece2EParamsResponseCodeNotzone, DNSSummaryDnssece2EParamsResponseCodeBadsig, DNSSummaryDnssece2EParamsResponseCodeBadkey, DNSSummaryDnssece2EParamsResponseCodeBadtime, DNSSummaryDnssece2EParamsResponseCodeBadmode, DNSSummaryDnssece2EParamsResponseCodeBadname, DNSSummaryDnssece2EParamsResponseCodeBadalg, DNSSummaryDnssece2EParamsResponseCodeBadtrunc, DNSSummaryDnssece2EParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryDnssece2EResponseEnvelope struct {
	Result  DNSSummaryDnssece2EResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    dnsSummaryDnssece2EResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseEnvelope]
type dnsSummaryDnssece2EResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryIPVersionParams struct {
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
	Format param.Field[DNSSummaryIPVersionParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryIPVersionParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryIPVersionParamsFormat string

const (
	DNSSummaryIPVersionParamsFormatJson DNSSummaryIPVersionParamsFormat = "JSON"
	DNSSummaryIPVersionParamsFormatCsv  DNSSummaryIPVersionParamsFormat = "CSV"
)

func (r DNSSummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryIPVersionParamsFormatJson, DNSSummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryIPVersionParamsProtocol string

const (
	DNSSummaryIPVersionParamsProtocolUdp   DNSSummaryIPVersionParamsProtocol = "UDP"
	DNSSummaryIPVersionParamsProtocolTCP   DNSSummaryIPVersionParamsProtocol = "TCP"
	DNSSummaryIPVersionParamsProtocolHTTPS DNSSummaryIPVersionParamsProtocol = "HTTPS"
	DNSSummaryIPVersionParamsProtocolTLS   DNSSummaryIPVersionParamsProtocol = "TLS"
)

func (r DNSSummaryIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryIPVersionParamsProtocolUdp, DNSSummaryIPVersionParamsProtocolTCP, DNSSummaryIPVersionParamsProtocolHTTPS, DNSSummaryIPVersionParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryIPVersionParamsQueryType string

const (
	DNSSummaryIPVersionParamsQueryTypeA          DNSSummaryIPVersionParamsQueryType = "A"
	DNSSummaryIPVersionParamsQueryTypeAAAA       DNSSummaryIPVersionParamsQueryType = "AAAA"
	DNSSummaryIPVersionParamsQueryTypeA6         DNSSummaryIPVersionParamsQueryType = "A6"
	DNSSummaryIPVersionParamsQueryTypeAfsdb      DNSSummaryIPVersionParamsQueryType = "AFSDB"
	DNSSummaryIPVersionParamsQueryTypeAny        DNSSummaryIPVersionParamsQueryType = "ANY"
	DNSSummaryIPVersionParamsQueryTypeApl        DNSSummaryIPVersionParamsQueryType = "APL"
	DNSSummaryIPVersionParamsQueryTypeAtma       DNSSummaryIPVersionParamsQueryType = "ATMA"
	DNSSummaryIPVersionParamsQueryTypeAXFR       DNSSummaryIPVersionParamsQueryType = "AXFR"
	DNSSummaryIPVersionParamsQueryTypeCAA        DNSSummaryIPVersionParamsQueryType = "CAA"
	DNSSummaryIPVersionParamsQueryTypeCdnskey    DNSSummaryIPVersionParamsQueryType = "CDNSKEY"
	DNSSummaryIPVersionParamsQueryTypeCds        DNSSummaryIPVersionParamsQueryType = "CDS"
	DNSSummaryIPVersionParamsQueryTypeCERT       DNSSummaryIPVersionParamsQueryType = "CERT"
	DNSSummaryIPVersionParamsQueryTypeCNAME      DNSSummaryIPVersionParamsQueryType = "CNAME"
	DNSSummaryIPVersionParamsQueryTypeCsync      DNSSummaryIPVersionParamsQueryType = "CSYNC"
	DNSSummaryIPVersionParamsQueryTypeDhcid      DNSSummaryIPVersionParamsQueryType = "DHCID"
	DNSSummaryIPVersionParamsQueryTypeDlv        DNSSummaryIPVersionParamsQueryType = "DLV"
	DNSSummaryIPVersionParamsQueryTypeDname      DNSSummaryIPVersionParamsQueryType = "DNAME"
	DNSSummaryIPVersionParamsQueryTypeDNSKEY     DNSSummaryIPVersionParamsQueryType = "DNSKEY"
	DNSSummaryIPVersionParamsQueryTypeDoa        DNSSummaryIPVersionParamsQueryType = "DOA"
	DNSSummaryIPVersionParamsQueryTypeDS         DNSSummaryIPVersionParamsQueryType = "DS"
	DNSSummaryIPVersionParamsQueryTypeEid        DNSSummaryIPVersionParamsQueryType = "EID"
	DNSSummaryIPVersionParamsQueryTypeEui48      DNSSummaryIPVersionParamsQueryType = "EUI48"
	DNSSummaryIPVersionParamsQueryTypeEui64      DNSSummaryIPVersionParamsQueryType = "EUI64"
	DNSSummaryIPVersionParamsQueryTypeGpos       DNSSummaryIPVersionParamsQueryType = "GPOS"
	DNSSummaryIPVersionParamsQueryTypeGid        DNSSummaryIPVersionParamsQueryType = "GID"
	DNSSummaryIPVersionParamsQueryTypeHinfo      DNSSummaryIPVersionParamsQueryType = "HINFO"
	DNSSummaryIPVersionParamsQueryTypeHip        DNSSummaryIPVersionParamsQueryType = "HIP"
	DNSSummaryIPVersionParamsQueryTypeHTTPS      DNSSummaryIPVersionParamsQueryType = "HTTPS"
	DNSSummaryIPVersionParamsQueryTypeIpseckey   DNSSummaryIPVersionParamsQueryType = "IPSECKEY"
	DNSSummaryIPVersionParamsQueryTypeIsdn       DNSSummaryIPVersionParamsQueryType = "ISDN"
	DNSSummaryIPVersionParamsQueryTypeIxfr       DNSSummaryIPVersionParamsQueryType = "IXFR"
	DNSSummaryIPVersionParamsQueryTypeKey        DNSSummaryIPVersionParamsQueryType = "KEY"
	DNSSummaryIPVersionParamsQueryTypeKx         DNSSummaryIPVersionParamsQueryType = "KX"
	DNSSummaryIPVersionParamsQueryTypeL32        DNSSummaryIPVersionParamsQueryType = "L32"
	DNSSummaryIPVersionParamsQueryTypeL64        DNSSummaryIPVersionParamsQueryType = "L64"
	DNSSummaryIPVersionParamsQueryTypeLOC        DNSSummaryIPVersionParamsQueryType = "LOC"
	DNSSummaryIPVersionParamsQueryTypeLp         DNSSummaryIPVersionParamsQueryType = "LP"
	DNSSummaryIPVersionParamsQueryTypeMaila      DNSSummaryIPVersionParamsQueryType = "MAILA"
	DNSSummaryIPVersionParamsQueryTypeMailb      DNSSummaryIPVersionParamsQueryType = "MAILB"
	DNSSummaryIPVersionParamsQueryTypeMB         DNSSummaryIPVersionParamsQueryType = "MB"
	DNSSummaryIPVersionParamsQueryTypeMd         DNSSummaryIPVersionParamsQueryType = "MD"
	DNSSummaryIPVersionParamsQueryTypeMf         DNSSummaryIPVersionParamsQueryType = "MF"
	DNSSummaryIPVersionParamsQueryTypeMg         DNSSummaryIPVersionParamsQueryType = "MG"
	DNSSummaryIPVersionParamsQueryTypeMinfo      DNSSummaryIPVersionParamsQueryType = "MINFO"
	DNSSummaryIPVersionParamsQueryTypeMr         DNSSummaryIPVersionParamsQueryType = "MR"
	DNSSummaryIPVersionParamsQueryTypeMX         DNSSummaryIPVersionParamsQueryType = "MX"
	DNSSummaryIPVersionParamsQueryTypeNAPTR      DNSSummaryIPVersionParamsQueryType = "NAPTR"
	DNSSummaryIPVersionParamsQueryTypeNb         DNSSummaryIPVersionParamsQueryType = "NB"
	DNSSummaryIPVersionParamsQueryTypeNbstat     DNSSummaryIPVersionParamsQueryType = "NBSTAT"
	DNSSummaryIPVersionParamsQueryTypeNid        DNSSummaryIPVersionParamsQueryType = "NID"
	DNSSummaryIPVersionParamsQueryTypeNimloc     DNSSummaryIPVersionParamsQueryType = "NIMLOC"
	DNSSummaryIPVersionParamsQueryTypeNinfo      DNSSummaryIPVersionParamsQueryType = "NINFO"
	DNSSummaryIPVersionParamsQueryTypeNS         DNSSummaryIPVersionParamsQueryType = "NS"
	DNSSummaryIPVersionParamsQueryTypeNsap       DNSSummaryIPVersionParamsQueryType = "NSAP"
	DNSSummaryIPVersionParamsQueryTypeNsec       DNSSummaryIPVersionParamsQueryType = "NSEC"
	DNSSummaryIPVersionParamsQueryTypeNsec3      DNSSummaryIPVersionParamsQueryType = "NSEC3"
	DNSSummaryIPVersionParamsQueryTypeNsec3Param DNSSummaryIPVersionParamsQueryType = "NSEC3PARAM"
	DNSSummaryIPVersionParamsQueryTypeNull       DNSSummaryIPVersionParamsQueryType = "NULL"
	DNSSummaryIPVersionParamsQueryTypeNxt        DNSSummaryIPVersionParamsQueryType = "NXT"
	DNSSummaryIPVersionParamsQueryTypeOpenpgpkey DNSSummaryIPVersionParamsQueryType = "OPENPGPKEY"
	DNSSummaryIPVersionParamsQueryTypeOpt        DNSSummaryIPVersionParamsQueryType = "OPT"
	DNSSummaryIPVersionParamsQueryTypePTR        DNSSummaryIPVersionParamsQueryType = "PTR"
	DNSSummaryIPVersionParamsQueryTypePx         DNSSummaryIPVersionParamsQueryType = "PX"
	DNSSummaryIPVersionParamsQueryTypeRkey       DNSSummaryIPVersionParamsQueryType = "RKEY"
	DNSSummaryIPVersionParamsQueryTypeRp         DNSSummaryIPVersionParamsQueryType = "RP"
	DNSSummaryIPVersionParamsQueryTypeRrsig      DNSSummaryIPVersionParamsQueryType = "RRSIG"
	DNSSummaryIPVersionParamsQueryTypeRt         DNSSummaryIPVersionParamsQueryType = "RT"
	DNSSummaryIPVersionParamsQueryTypeSig        DNSSummaryIPVersionParamsQueryType = "SIG"
	DNSSummaryIPVersionParamsQueryTypeSink       DNSSummaryIPVersionParamsQueryType = "SINK"
	DNSSummaryIPVersionParamsQueryTypeSMIMEA     DNSSummaryIPVersionParamsQueryType = "SMIMEA"
	DNSSummaryIPVersionParamsQueryTypeSOA        DNSSummaryIPVersionParamsQueryType = "SOA"
	DNSSummaryIPVersionParamsQueryTypeSPF        DNSSummaryIPVersionParamsQueryType = "SPF"
	DNSSummaryIPVersionParamsQueryTypeSRV        DNSSummaryIPVersionParamsQueryType = "SRV"
	DNSSummaryIPVersionParamsQueryTypeSSHFP      DNSSummaryIPVersionParamsQueryType = "SSHFP"
	DNSSummaryIPVersionParamsQueryTypeSVCB       DNSSummaryIPVersionParamsQueryType = "SVCB"
	DNSSummaryIPVersionParamsQueryTypeTa         DNSSummaryIPVersionParamsQueryType = "TA"
	DNSSummaryIPVersionParamsQueryTypeTalink     DNSSummaryIPVersionParamsQueryType = "TALINK"
	DNSSummaryIPVersionParamsQueryTypeTkey       DNSSummaryIPVersionParamsQueryType = "TKEY"
	DNSSummaryIPVersionParamsQueryTypeTLSA       DNSSummaryIPVersionParamsQueryType = "TLSA"
	DNSSummaryIPVersionParamsQueryTypeTSIG       DNSSummaryIPVersionParamsQueryType = "TSIG"
	DNSSummaryIPVersionParamsQueryTypeTXT        DNSSummaryIPVersionParamsQueryType = "TXT"
	DNSSummaryIPVersionParamsQueryTypeUinfo      DNSSummaryIPVersionParamsQueryType = "UINFO"
	DNSSummaryIPVersionParamsQueryTypeUID        DNSSummaryIPVersionParamsQueryType = "UID"
	DNSSummaryIPVersionParamsQueryTypeUnspec     DNSSummaryIPVersionParamsQueryType = "UNSPEC"
	DNSSummaryIPVersionParamsQueryTypeURI        DNSSummaryIPVersionParamsQueryType = "URI"
	DNSSummaryIPVersionParamsQueryTypeWks        DNSSummaryIPVersionParamsQueryType = "WKS"
	DNSSummaryIPVersionParamsQueryTypeX25        DNSSummaryIPVersionParamsQueryType = "X25"
	DNSSummaryIPVersionParamsQueryTypeZonemd     DNSSummaryIPVersionParamsQueryType = "ZONEMD"
)

func (r DNSSummaryIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryIPVersionParamsQueryTypeA, DNSSummaryIPVersionParamsQueryTypeAAAA, DNSSummaryIPVersionParamsQueryTypeA6, DNSSummaryIPVersionParamsQueryTypeAfsdb, DNSSummaryIPVersionParamsQueryTypeAny, DNSSummaryIPVersionParamsQueryTypeApl, DNSSummaryIPVersionParamsQueryTypeAtma, DNSSummaryIPVersionParamsQueryTypeAXFR, DNSSummaryIPVersionParamsQueryTypeCAA, DNSSummaryIPVersionParamsQueryTypeCdnskey, DNSSummaryIPVersionParamsQueryTypeCds, DNSSummaryIPVersionParamsQueryTypeCERT, DNSSummaryIPVersionParamsQueryTypeCNAME, DNSSummaryIPVersionParamsQueryTypeCsync, DNSSummaryIPVersionParamsQueryTypeDhcid, DNSSummaryIPVersionParamsQueryTypeDlv, DNSSummaryIPVersionParamsQueryTypeDname, DNSSummaryIPVersionParamsQueryTypeDNSKEY, DNSSummaryIPVersionParamsQueryTypeDoa, DNSSummaryIPVersionParamsQueryTypeDS, DNSSummaryIPVersionParamsQueryTypeEid, DNSSummaryIPVersionParamsQueryTypeEui48, DNSSummaryIPVersionParamsQueryTypeEui64, DNSSummaryIPVersionParamsQueryTypeGpos, DNSSummaryIPVersionParamsQueryTypeGid, DNSSummaryIPVersionParamsQueryTypeHinfo, DNSSummaryIPVersionParamsQueryTypeHip, DNSSummaryIPVersionParamsQueryTypeHTTPS, DNSSummaryIPVersionParamsQueryTypeIpseckey, DNSSummaryIPVersionParamsQueryTypeIsdn, DNSSummaryIPVersionParamsQueryTypeIxfr, DNSSummaryIPVersionParamsQueryTypeKey, DNSSummaryIPVersionParamsQueryTypeKx, DNSSummaryIPVersionParamsQueryTypeL32, DNSSummaryIPVersionParamsQueryTypeL64, DNSSummaryIPVersionParamsQueryTypeLOC, DNSSummaryIPVersionParamsQueryTypeLp, DNSSummaryIPVersionParamsQueryTypeMaila, DNSSummaryIPVersionParamsQueryTypeMailb, DNSSummaryIPVersionParamsQueryTypeMB, DNSSummaryIPVersionParamsQueryTypeMd, DNSSummaryIPVersionParamsQueryTypeMf, DNSSummaryIPVersionParamsQueryTypeMg, DNSSummaryIPVersionParamsQueryTypeMinfo, DNSSummaryIPVersionParamsQueryTypeMr, DNSSummaryIPVersionParamsQueryTypeMX, DNSSummaryIPVersionParamsQueryTypeNAPTR, DNSSummaryIPVersionParamsQueryTypeNb, DNSSummaryIPVersionParamsQueryTypeNbstat, DNSSummaryIPVersionParamsQueryTypeNid, DNSSummaryIPVersionParamsQueryTypeNimloc, DNSSummaryIPVersionParamsQueryTypeNinfo, DNSSummaryIPVersionParamsQueryTypeNS, DNSSummaryIPVersionParamsQueryTypeNsap, DNSSummaryIPVersionParamsQueryTypeNsec, DNSSummaryIPVersionParamsQueryTypeNsec3, DNSSummaryIPVersionParamsQueryTypeNsec3Param, DNSSummaryIPVersionParamsQueryTypeNull, DNSSummaryIPVersionParamsQueryTypeNxt, DNSSummaryIPVersionParamsQueryTypeOpenpgpkey, DNSSummaryIPVersionParamsQueryTypeOpt, DNSSummaryIPVersionParamsQueryTypePTR, DNSSummaryIPVersionParamsQueryTypePx, DNSSummaryIPVersionParamsQueryTypeRkey, DNSSummaryIPVersionParamsQueryTypeRp, DNSSummaryIPVersionParamsQueryTypeRrsig, DNSSummaryIPVersionParamsQueryTypeRt, DNSSummaryIPVersionParamsQueryTypeSig, DNSSummaryIPVersionParamsQueryTypeSink, DNSSummaryIPVersionParamsQueryTypeSMIMEA, DNSSummaryIPVersionParamsQueryTypeSOA, DNSSummaryIPVersionParamsQueryTypeSPF, DNSSummaryIPVersionParamsQueryTypeSRV, DNSSummaryIPVersionParamsQueryTypeSSHFP, DNSSummaryIPVersionParamsQueryTypeSVCB, DNSSummaryIPVersionParamsQueryTypeTa, DNSSummaryIPVersionParamsQueryTypeTalink, DNSSummaryIPVersionParamsQueryTypeTkey, DNSSummaryIPVersionParamsQueryTypeTLSA, DNSSummaryIPVersionParamsQueryTypeTSIG, DNSSummaryIPVersionParamsQueryTypeTXT, DNSSummaryIPVersionParamsQueryTypeUinfo, DNSSummaryIPVersionParamsQueryTypeUID, DNSSummaryIPVersionParamsQueryTypeUnspec, DNSSummaryIPVersionParamsQueryTypeURI, DNSSummaryIPVersionParamsQueryTypeWks, DNSSummaryIPVersionParamsQueryTypeX25, DNSSummaryIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryIPVersionParamsResponseCode string

const (
	DNSSummaryIPVersionParamsResponseCodeNoerror   DNSSummaryIPVersionParamsResponseCode = "NOERROR"
	DNSSummaryIPVersionParamsResponseCodeFormerr   DNSSummaryIPVersionParamsResponseCode = "FORMERR"
	DNSSummaryIPVersionParamsResponseCodeServfail  DNSSummaryIPVersionParamsResponseCode = "SERVFAIL"
	DNSSummaryIPVersionParamsResponseCodeNxdomain  DNSSummaryIPVersionParamsResponseCode = "NXDOMAIN"
	DNSSummaryIPVersionParamsResponseCodeNotimp    DNSSummaryIPVersionParamsResponseCode = "NOTIMP"
	DNSSummaryIPVersionParamsResponseCodeRefused   DNSSummaryIPVersionParamsResponseCode = "REFUSED"
	DNSSummaryIPVersionParamsResponseCodeYxdomain  DNSSummaryIPVersionParamsResponseCode = "YXDOMAIN"
	DNSSummaryIPVersionParamsResponseCodeYxrrset   DNSSummaryIPVersionParamsResponseCode = "YXRRSET"
	DNSSummaryIPVersionParamsResponseCodeNxrrset   DNSSummaryIPVersionParamsResponseCode = "NXRRSET"
	DNSSummaryIPVersionParamsResponseCodeNotauth   DNSSummaryIPVersionParamsResponseCode = "NOTAUTH"
	DNSSummaryIPVersionParamsResponseCodeNotzone   DNSSummaryIPVersionParamsResponseCode = "NOTZONE"
	DNSSummaryIPVersionParamsResponseCodeBadsig    DNSSummaryIPVersionParamsResponseCode = "BADSIG"
	DNSSummaryIPVersionParamsResponseCodeBadkey    DNSSummaryIPVersionParamsResponseCode = "BADKEY"
	DNSSummaryIPVersionParamsResponseCodeBadtime   DNSSummaryIPVersionParamsResponseCode = "BADTIME"
	DNSSummaryIPVersionParamsResponseCodeBadmode   DNSSummaryIPVersionParamsResponseCode = "BADMODE"
	DNSSummaryIPVersionParamsResponseCodeBadname   DNSSummaryIPVersionParamsResponseCode = "BADNAME"
	DNSSummaryIPVersionParamsResponseCodeBadalg    DNSSummaryIPVersionParamsResponseCode = "BADALG"
	DNSSummaryIPVersionParamsResponseCodeBadtrunc  DNSSummaryIPVersionParamsResponseCode = "BADTRUNC"
	DNSSummaryIPVersionParamsResponseCodeBadcookie DNSSummaryIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryIPVersionParamsResponseCodeNoerror, DNSSummaryIPVersionParamsResponseCodeFormerr, DNSSummaryIPVersionParamsResponseCodeServfail, DNSSummaryIPVersionParamsResponseCodeNxdomain, DNSSummaryIPVersionParamsResponseCodeNotimp, DNSSummaryIPVersionParamsResponseCodeRefused, DNSSummaryIPVersionParamsResponseCodeYxdomain, DNSSummaryIPVersionParamsResponseCodeYxrrset, DNSSummaryIPVersionParamsResponseCodeNxrrset, DNSSummaryIPVersionParamsResponseCodeNotauth, DNSSummaryIPVersionParamsResponseCodeNotzone, DNSSummaryIPVersionParamsResponseCodeBadsig, DNSSummaryIPVersionParamsResponseCodeBadkey, DNSSummaryIPVersionParamsResponseCodeBadtime, DNSSummaryIPVersionParamsResponseCodeBadmode, DNSSummaryIPVersionParamsResponseCodeBadname, DNSSummaryIPVersionParamsResponseCodeBadalg, DNSSummaryIPVersionParamsResponseCodeBadtrunc, DNSSummaryIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryIPVersionResponseEnvelope struct {
	Result  DNSSummaryIPVersionResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    dnsSummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseEnvelope]
type dnsSummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryMatchingAnswerParams struct {
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
	Format param.Field[DNSSummaryMatchingAnswerParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryMatchingAnswerParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryMatchingAnswerParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryMatchingAnswerParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryMatchingAnswerParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryMatchingAnswerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryMatchingAnswerParamsFormat string

const (
	DNSSummaryMatchingAnswerParamsFormatJson DNSSummaryMatchingAnswerParamsFormat = "JSON"
	DNSSummaryMatchingAnswerParamsFormatCsv  DNSSummaryMatchingAnswerParamsFormat = "CSV"
)

func (r DNSSummaryMatchingAnswerParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryMatchingAnswerParamsFormatJson, DNSSummaryMatchingAnswerParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryMatchingAnswerParamsProtocol string

const (
	DNSSummaryMatchingAnswerParamsProtocolUdp   DNSSummaryMatchingAnswerParamsProtocol = "UDP"
	DNSSummaryMatchingAnswerParamsProtocolTCP   DNSSummaryMatchingAnswerParamsProtocol = "TCP"
	DNSSummaryMatchingAnswerParamsProtocolHTTPS DNSSummaryMatchingAnswerParamsProtocol = "HTTPS"
	DNSSummaryMatchingAnswerParamsProtocolTLS   DNSSummaryMatchingAnswerParamsProtocol = "TLS"
)

func (r DNSSummaryMatchingAnswerParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryMatchingAnswerParamsProtocolUdp, DNSSummaryMatchingAnswerParamsProtocolTCP, DNSSummaryMatchingAnswerParamsProtocolHTTPS, DNSSummaryMatchingAnswerParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryMatchingAnswerParamsQueryType string

const (
	DNSSummaryMatchingAnswerParamsQueryTypeA          DNSSummaryMatchingAnswerParamsQueryType = "A"
	DNSSummaryMatchingAnswerParamsQueryTypeAAAA       DNSSummaryMatchingAnswerParamsQueryType = "AAAA"
	DNSSummaryMatchingAnswerParamsQueryTypeA6         DNSSummaryMatchingAnswerParamsQueryType = "A6"
	DNSSummaryMatchingAnswerParamsQueryTypeAfsdb      DNSSummaryMatchingAnswerParamsQueryType = "AFSDB"
	DNSSummaryMatchingAnswerParamsQueryTypeAny        DNSSummaryMatchingAnswerParamsQueryType = "ANY"
	DNSSummaryMatchingAnswerParamsQueryTypeApl        DNSSummaryMatchingAnswerParamsQueryType = "APL"
	DNSSummaryMatchingAnswerParamsQueryTypeAtma       DNSSummaryMatchingAnswerParamsQueryType = "ATMA"
	DNSSummaryMatchingAnswerParamsQueryTypeAXFR       DNSSummaryMatchingAnswerParamsQueryType = "AXFR"
	DNSSummaryMatchingAnswerParamsQueryTypeCAA        DNSSummaryMatchingAnswerParamsQueryType = "CAA"
	DNSSummaryMatchingAnswerParamsQueryTypeCdnskey    DNSSummaryMatchingAnswerParamsQueryType = "CDNSKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeCds        DNSSummaryMatchingAnswerParamsQueryType = "CDS"
	DNSSummaryMatchingAnswerParamsQueryTypeCERT       DNSSummaryMatchingAnswerParamsQueryType = "CERT"
	DNSSummaryMatchingAnswerParamsQueryTypeCNAME      DNSSummaryMatchingAnswerParamsQueryType = "CNAME"
	DNSSummaryMatchingAnswerParamsQueryTypeCsync      DNSSummaryMatchingAnswerParamsQueryType = "CSYNC"
	DNSSummaryMatchingAnswerParamsQueryTypeDhcid      DNSSummaryMatchingAnswerParamsQueryType = "DHCID"
	DNSSummaryMatchingAnswerParamsQueryTypeDlv        DNSSummaryMatchingAnswerParamsQueryType = "DLV"
	DNSSummaryMatchingAnswerParamsQueryTypeDname      DNSSummaryMatchingAnswerParamsQueryType = "DNAME"
	DNSSummaryMatchingAnswerParamsQueryTypeDNSKEY     DNSSummaryMatchingAnswerParamsQueryType = "DNSKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeDoa        DNSSummaryMatchingAnswerParamsQueryType = "DOA"
	DNSSummaryMatchingAnswerParamsQueryTypeDS         DNSSummaryMatchingAnswerParamsQueryType = "DS"
	DNSSummaryMatchingAnswerParamsQueryTypeEid        DNSSummaryMatchingAnswerParamsQueryType = "EID"
	DNSSummaryMatchingAnswerParamsQueryTypeEui48      DNSSummaryMatchingAnswerParamsQueryType = "EUI48"
	DNSSummaryMatchingAnswerParamsQueryTypeEui64      DNSSummaryMatchingAnswerParamsQueryType = "EUI64"
	DNSSummaryMatchingAnswerParamsQueryTypeGpos       DNSSummaryMatchingAnswerParamsQueryType = "GPOS"
	DNSSummaryMatchingAnswerParamsQueryTypeGid        DNSSummaryMatchingAnswerParamsQueryType = "GID"
	DNSSummaryMatchingAnswerParamsQueryTypeHinfo      DNSSummaryMatchingAnswerParamsQueryType = "HINFO"
	DNSSummaryMatchingAnswerParamsQueryTypeHip        DNSSummaryMatchingAnswerParamsQueryType = "HIP"
	DNSSummaryMatchingAnswerParamsQueryTypeHTTPS      DNSSummaryMatchingAnswerParamsQueryType = "HTTPS"
	DNSSummaryMatchingAnswerParamsQueryTypeIpseckey   DNSSummaryMatchingAnswerParamsQueryType = "IPSECKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeIsdn       DNSSummaryMatchingAnswerParamsQueryType = "ISDN"
	DNSSummaryMatchingAnswerParamsQueryTypeIxfr       DNSSummaryMatchingAnswerParamsQueryType = "IXFR"
	DNSSummaryMatchingAnswerParamsQueryTypeKey        DNSSummaryMatchingAnswerParamsQueryType = "KEY"
	DNSSummaryMatchingAnswerParamsQueryTypeKx         DNSSummaryMatchingAnswerParamsQueryType = "KX"
	DNSSummaryMatchingAnswerParamsQueryTypeL32        DNSSummaryMatchingAnswerParamsQueryType = "L32"
	DNSSummaryMatchingAnswerParamsQueryTypeL64        DNSSummaryMatchingAnswerParamsQueryType = "L64"
	DNSSummaryMatchingAnswerParamsQueryTypeLOC        DNSSummaryMatchingAnswerParamsQueryType = "LOC"
	DNSSummaryMatchingAnswerParamsQueryTypeLp         DNSSummaryMatchingAnswerParamsQueryType = "LP"
	DNSSummaryMatchingAnswerParamsQueryTypeMaila      DNSSummaryMatchingAnswerParamsQueryType = "MAILA"
	DNSSummaryMatchingAnswerParamsQueryTypeMailb      DNSSummaryMatchingAnswerParamsQueryType = "MAILB"
	DNSSummaryMatchingAnswerParamsQueryTypeMB         DNSSummaryMatchingAnswerParamsQueryType = "MB"
	DNSSummaryMatchingAnswerParamsQueryTypeMd         DNSSummaryMatchingAnswerParamsQueryType = "MD"
	DNSSummaryMatchingAnswerParamsQueryTypeMf         DNSSummaryMatchingAnswerParamsQueryType = "MF"
	DNSSummaryMatchingAnswerParamsQueryTypeMg         DNSSummaryMatchingAnswerParamsQueryType = "MG"
	DNSSummaryMatchingAnswerParamsQueryTypeMinfo      DNSSummaryMatchingAnswerParamsQueryType = "MINFO"
	DNSSummaryMatchingAnswerParamsQueryTypeMr         DNSSummaryMatchingAnswerParamsQueryType = "MR"
	DNSSummaryMatchingAnswerParamsQueryTypeMX         DNSSummaryMatchingAnswerParamsQueryType = "MX"
	DNSSummaryMatchingAnswerParamsQueryTypeNAPTR      DNSSummaryMatchingAnswerParamsQueryType = "NAPTR"
	DNSSummaryMatchingAnswerParamsQueryTypeNb         DNSSummaryMatchingAnswerParamsQueryType = "NB"
	DNSSummaryMatchingAnswerParamsQueryTypeNbstat     DNSSummaryMatchingAnswerParamsQueryType = "NBSTAT"
	DNSSummaryMatchingAnswerParamsQueryTypeNid        DNSSummaryMatchingAnswerParamsQueryType = "NID"
	DNSSummaryMatchingAnswerParamsQueryTypeNimloc     DNSSummaryMatchingAnswerParamsQueryType = "NIMLOC"
	DNSSummaryMatchingAnswerParamsQueryTypeNinfo      DNSSummaryMatchingAnswerParamsQueryType = "NINFO"
	DNSSummaryMatchingAnswerParamsQueryTypeNS         DNSSummaryMatchingAnswerParamsQueryType = "NS"
	DNSSummaryMatchingAnswerParamsQueryTypeNsap       DNSSummaryMatchingAnswerParamsQueryType = "NSAP"
	DNSSummaryMatchingAnswerParamsQueryTypeNsec       DNSSummaryMatchingAnswerParamsQueryType = "NSEC"
	DNSSummaryMatchingAnswerParamsQueryTypeNsec3      DNSSummaryMatchingAnswerParamsQueryType = "NSEC3"
	DNSSummaryMatchingAnswerParamsQueryTypeNsec3Param DNSSummaryMatchingAnswerParamsQueryType = "NSEC3PARAM"
	DNSSummaryMatchingAnswerParamsQueryTypeNull       DNSSummaryMatchingAnswerParamsQueryType = "NULL"
	DNSSummaryMatchingAnswerParamsQueryTypeNxt        DNSSummaryMatchingAnswerParamsQueryType = "NXT"
	DNSSummaryMatchingAnswerParamsQueryTypeOpenpgpkey DNSSummaryMatchingAnswerParamsQueryType = "OPENPGPKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeOpt        DNSSummaryMatchingAnswerParamsQueryType = "OPT"
	DNSSummaryMatchingAnswerParamsQueryTypePTR        DNSSummaryMatchingAnswerParamsQueryType = "PTR"
	DNSSummaryMatchingAnswerParamsQueryTypePx         DNSSummaryMatchingAnswerParamsQueryType = "PX"
	DNSSummaryMatchingAnswerParamsQueryTypeRkey       DNSSummaryMatchingAnswerParamsQueryType = "RKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeRp         DNSSummaryMatchingAnswerParamsQueryType = "RP"
	DNSSummaryMatchingAnswerParamsQueryTypeRrsig      DNSSummaryMatchingAnswerParamsQueryType = "RRSIG"
	DNSSummaryMatchingAnswerParamsQueryTypeRt         DNSSummaryMatchingAnswerParamsQueryType = "RT"
	DNSSummaryMatchingAnswerParamsQueryTypeSig        DNSSummaryMatchingAnswerParamsQueryType = "SIG"
	DNSSummaryMatchingAnswerParamsQueryTypeSink       DNSSummaryMatchingAnswerParamsQueryType = "SINK"
	DNSSummaryMatchingAnswerParamsQueryTypeSMIMEA     DNSSummaryMatchingAnswerParamsQueryType = "SMIMEA"
	DNSSummaryMatchingAnswerParamsQueryTypeSOA        DNSSummaryMatchingAnswerParamsQueryType = "SOA"
	DNSSummaryMatchingAnswerParamsQueryTypeSPF        DNSSummaryMatchingAnswerParamsQueryType = "SPF"
	DNSSummaryMatchingAnswerParamsQueryTypeSRV        DNSSummaryMatchingAnswerParamsQueryType = "SRV"
	DNSSummaryMatchingAnswerParamsQueryTypeSSHFP      DNSSummaryMatchingAnswerParamsQueryType = "SSHFP"
	DNSSummaryMatchingAnswerParamsQueryTypeSVCB       DNSSummaryMatchingAnswerParamsQueryType = "SVCB"
	DNSSummaryMatchingAnswerParamsQueryTypeTa         DNSSummaryMatchingAnswerParamsQueryType = "TA"
	DNSSummaryMatchingAnswerParamsQueryTypeTalink     DNSSummaryMatchingAnswerParamsQueryType = "TALINK"
	DNSSummaryMatchingAnswerParamsQueryTypeTkey       DNSSummaryMatchingAnswerParamsQueryType = "TKEY"
	DNSSummaryMatchingAnswerParamsQueryTypeTLSA       DNSSummaryMatchingAnswerParamsQueryType = "TLSA"
	DNSSummaryMatchingAnswerParamsQueryTypeTSIG       DNSSummaryMatchingAnswerParamsQueryType = "TSIG"
	DNSSummaryMatchingAnswerParamsQueryTypeTXT        DNSSummaryMatchingAnswerParamsQueryType = "TXT"
	DNSSummaryMatchingAnswerParamsQueryTypeUinfo      DNSSummaryMatchingAnswerParamsQueryType = "UINFO"
	DNSSummaryMatchingAnswerParamsQueryTypeUID        DNSSummaryMatchingAnswerParamsQueryType = "UID"
	DNSSummaryMatchingAnswerParamsQueryTypeUnspec     DNSSummaryMatchingAnswerParamsQueryType = "UNSPEC"
	DNSSummaryMatchingAnswerParamsQueryTypeURI        DNSSummaryMatchingAnswerParamsQueryType = "URI"
	DNSSummaryMatchingAnswerParamsQueryTypeWks        DNSSummaryMatchingAnswerParamsQueryType = "WKS"
	DNSSummaryMatchingAnswerParamsQueryTypeX25        DNSSummaryMatchingAnswerParamsQueryType = "X25"
	DNSSummaryMatchingAnswerParamsQueryTypeZonemd     DNSSummaryMatchingAnswerParamsQueryType = "ZONEMD"
)

func (r DNSSummaryMatchingAnswerParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryMatchingAnswerParamsQueryTypeA, DNSSummaryMatchingAnswerParamsQueryTypeAAAA, DNSSummaryMatchingAnswerParamsQueryTypeA6, DNSSummaryMatchingAnswerParamsQueryTypeAfsdb, DNSSummaryMatchingAnswerParamsQueryTypeAny, DNSSummaryMatchingAnswerParamsQueryTypeApl, DNSSummaryMatchingAnswerParamsQueryTypeAtma, DNSSummaryMatchingAnswerParamsQueryTypeAXFR, DNSSummaryMatchingAnswerParamsQueryTypeCAA, DNSSummaryMatchingAnswerParamsQueryTypeCdnskey, DNSSummaryMatchingAnswerParamsQueryTypeCds, DNSSummaryMatchingAnswerParamsQueryTypeCERT, DNSSummaryMatchingAnswerParamsQueryTypeCNAME, DNSSummaryMatchingAnswerParamsQueryTypeCsync, DNSSummaryMatchingAnswerParamsQueryTypeDhcid, DNSSummaryMatchingAnswerParamsQueryTypeDlv, DNSSummaryMatchingAnswerParamsQueryTypeDname, DNSSummaryMatchingAnswerParamsQueryTypeDNSKEY, DNSSummaryMatchingAnswerParamsQueryTypeDoa, DNSSummaryMatchingAnswerParamsQueryTypeDS, DNSSummaryMatchingAnswerParamsQueryTypeEid, DNSSummaryMatchingAnswerParamsQueryTypeEui48, DNSSummaryMatchingAnswerParamsQueryTypeEui64, DNSSummaryMatchingAnswerParamsQueryTypeGpos, DNSSummaryMatchingAnswerParamsQueryTypeGid, DNSSummaryMatchingAnswerParamsQueryTypeHinfo, DNSSummaryMatchingAnswerParamsQueryTypeHip, DNSSummaryMatchingAnswerParamsQueryTypeHTTPS, DNSSummaryMatchingAnswerParamsQueryTypeIpseckey, DNSSummaryMatchingAnswerParamsQueryTypeIsdn, DNSSummaryMatchingAnswerParamsQueryTypeIxfr, DNSSummaryMatchingAnswerParamsQueryTypeKey, DNSSummaryMatchingAnswerParamsQueryTypeKx, DNSSummaryMatchingAnswerParamsQueryTypeL32, DNSSummaryMatchingAnswerParamsQueryTypeL64, DNSSummaryMatchingAnswerParamsQueryTypeLOC, DNSSummaryMatchingAnswerParamsQueryTypeLp, DNSSummaryMatchingAnswerParamsQueryTypeMaila, DNSSummaryMatchingAnswerParamsQueryTypeMailb, DNSSummaryMatchingAnswerParamsQueryTypeMB, DNSSummaryMatchingAnswerParamsQueryTypeMd, DNSSummaryMatchingAnswerParamsQueryTypeMf, DNSSummaryMatchingAnswerParamsQueryTypeMg, DNSSummaryMatchingAnswerParamsQueryTypeMinfo, DNSSummaryMatchingAnswerParamsQueryTypeMr, DNSSummaryMatchingAnswerParamsQueryTypeMX, DNSSummaryMatchingAnswerParamsQueryTypeNAPTR, DNSSummaryMatchingAnswerParamsQueryTypeNb, DNSSummaryMatchingAnswerParamsQueryTypeNbstat, DNSSummaryMatchingAnswerParamsQueryTypeNid, DNSSummaryMatchingAnswerParamsQueryTypeNimloc, DNSSummaryMatchingAnswerParamsQueryTypeNinfo, DNSSummaryMatchingAnswerParamsQueryTypeNS, DNSSummaryMatchingAnswerParamsQueryTypeNsap, DNSSummaryMatchingAnswerParamsQueryTypeNsec, DNSSummaryMatchingAnswerParamsQueryTypeNsec3, DNSSummaryMatchingAnswerParamsQueryTypeNsec3Param, DNSSummaryMatchingAnswerParamsQueryTypeNull, DNSSummaryMatchingAnswerParamsQueryTypeNxt, DNSSummaryMatchingAnswerParamsQueryTypeOpenpgpkey, DNSSummaryMatchingAnswerParamsQueryTypeOpt, DNSSummaryMatchingAnswerParamsQueryTypePTR, DNSSummaryMatchingAnswerParamsQueryTypePx, DNSSummaryMatchingAnswerParamsQueryTypeRkey, DNSSummaryMatchingAnswerParamsQueryTypeRp, DNSSummaryMatchingAnswerParamsQueryTypeRrsig, DNSSummaryMatchingAnswerParamsQueryTypeRt, DNSSummaryMatchingAnswerParamsQueryTypeSig, DNSSummaryMatchingAnswerParamsQueryTypeSink, DNSSummaryMatchingAnswerParamsQueryTypeSMIMEA, DNSSummaryMatchingAnswerParamsQueryTypeSOA, DNSSummaryMatchingAnswerParamsQueryTypeSPF, DNSSummaryMatchingAnswerParamsQueryTypeSRV, DNSSummaryMatchingAnswerParamsQueryTypeSSHFP, DNSSummaryMatchingAnswerParamsQueryTypeSVCB, DNSSummaryMatchingAnswerParamsQueryTypeTa, DNSSummaryMatchingAnswerParamsQueryTypeTalink, DNSSummaryMatchingAnswerParamsQueryTypeTkey, DNSSummaryMatchingAnswerParamsQueryTypeTLSA, DNSSummaryMatchingAnswerParamsQueryTypeTSIG, DNSSummaryMatchingAnswerParamsQueryTypeTXT, DNSSummaryMatchingAnswerParamsQueryTypeUinfo, DNSSummaryMatchingAnswerParamsQueryTypeUID, DNSSummaryMatchingAnswerParamsQueryTypeUnspec, DNSSummaryMatchingAnswerParamsQueryTypeURI, DNSSummaryMatchingAnswerParamsQueryTypeWks, DNSSummaryMatchingAnswerParamsQueryTypeX25, DNSSummaryMatchingAnswerParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryMatchingAnswerParamsResponseCode string

const (
	DNSSummaryMatchingAnswerParamsResponseCodeNoerror   DNSSummaryMatchingAnswerParamsResponseCode = "NOERROR"
	DNSSummaryMatchingAnswerParamsResponseCodeFormerr   DNSSummaryMatchingAnswerParamsResponseCode = "FORMERR"
	DNSSummaryMatchingAnswerParamsResponseCodeServfail  DNSSummaryMatchingAnswerParamsResponseCode = "SERVFAIL"
	DNSSummaryMatchingAnswerParamsResponseCodeNxdomain  DNSSummaryMatchingAnswerParamsResponseCode = "NXDOMAIN"
	DNSSummaryMatchingAnswerParamsResponseCodeNotimp    DNSSummaryMatchingAnswerParamsResponseCode = "NOTIMP"
	DNSSummaryMatchingAnswerParamsResponseCodeRefused   DNSSummaryMatchingAnswerParamsResponseCode = "REFUSED"
	DNSSummaryMatchingAnswerParamsResponseCodeYxdomain  DNSSummaryMatchingAnswerParamsResponseCode = "YXDOMAIN"
	DNSSummaryMatchingAnswerParamsResponseCodeYxrrset   DNSSummaryMatchingAnswerParamsResponseCode = "YXRRSET"
	DNSSummaryMatchingAnswerParamsResponseCodeNxrrset   DNSSummaryMatchingAnswerParamsResponseCode = "NXRRSET"
	DNSSummaryMatchingAnswerParamsResponseCodeNotauth   DNSSummaryMatchingAnswerParamsResponseCode = "NOTAUTH"
	DNSSummaryMatchingAnswerParamsResponseCodeNotzone   DNSSummaryMatchingAnswerParamsResponseCode = "NOTZONE"
	DNSSummaryMatchingAnswerParamsResponseCodeBadsig    DNSSummaryMatchingAnswerParamsResponseCode = "BADSIG"
	DNSSummaryMatchingAnswerParamsResponseCodeBadkey    DNSSummaryMatchingAnswerParamsResponseCode = "BADKEY"
	DNSSummaryMatchingAnswerParamsResponseCodeBadtime   DNSSummaryMatchingAnswerParamsResponseCode = "BADTIME"
	DNSSummaryMatchingAnswerParamsResponseCodeBadmode   DNSSummaryMatchingAnswerParamsResponseCode = "BADMODE"
	DNSSummaryMatchingAnswerParamsResponseCodeBadname   DNSSummaryMatchingAnswerParamsResponseCode = "BADNAME"
	DNSSummaryMatchingAnswerParamsResponseCodeBadalg    DNSSummaryMatchingAnswerParamsResponseCode = "BADALG"
	DNSSummaryMatchingAnswerParamsResponseCodeBadtrunc  DNSSummaryMatchingAnswerParamsResponseCode = "BADTRUNC"
	DNSSummaryMatchingAnswerParamsResponseCodeBadcookie DNSSummaryMatchingAnswerParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryMatchingAnswerParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryMatchingAnswerParamsResponseCodeNoerror, DNSSummaryMatchingAnswerParamsResponseCodeFormerr, DNSSummaryMatchingAnswerParamsResponseCodeServfail, DNSSummaryMatchingAnswerParamsResponseCodeNxdomain, DNSSummaryMatchingAnswerParamsResponseCodeNotimp, DNSSummaryMatchingAnswerParamsResponseCodeRefused, DNSSummaryMatchingAnswerParamsResponseCodeYxdomain, DNSSummaryMatchingAnswerParamsResponseCodeYxrrset, DNSSummaryMatchingAnswerParamsResponseCodeNxrrset, DNSSummaryMatchingAnswerParamsResponseCodeNotauth, DNSSummaryMatchingAnswerParamsResponseCodeNotzone, DNSSummaryMatchingAnswerParamsResponseCodeBadsig, DNSSummaryMatchingAnswerParamsResponseCodeBadkey, DNSSummaryMatchingAnswerParamsResponseCodeBadtime, DNSSummaryMatchingAnswerParamsResponseCodeBadmode, DNSSummaryMatchingAnswerParamsResponseCodeBadname, DNSSummaryMatchingAnswerParamsResponseCodeBadalg, DNSSummaryMatchingAnswerParamsResponseCodeBadtrunc, DNSSummaryMatchingAnswerParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryMatchingAnswerResponseEnvelope struct {
	Result  DNSSummaryMatchingAnswerResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    dnsSummaryMatchingAnswerResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseEnvelope]
type dnsSummaryMatchingAnswerResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryProtocolParams struct {
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
	Format param.Field[DNSSummaryProtocolParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryProtocolParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryProtocolParamsFormat string

const (
	DNSSummaryProtocolParamsFormatJson DNSSummaryProtocolParamsFormat = "JSON"
	DNSSummaryProtocolParamsFormatCsv  DNSSummaryProtocolParamsFormat = "CSV"
)

func (r DNSSummaryProtocolParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryProtocolParamsFormatJson, DNSSummaryProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryProtocolParamsQueryType string

const (
	DNSSummaryProtocolParamsQueryTypeA          DNSSummaryProtocolParamsQueryType = "A"
	DNSSummaryProtocolParamsQueryTypeAAAA       DNSSummaryProtocolParamsQueryType = "AAAA"
	DNSSummaryProtocolParamsQueryTypeA6         DNSSummaryProtocolParamsQueryType = "A6"
	DNSSummaryProtocolParamsQueryTypeAfsdb      DNSSummaryProtocolParamsQueryType = "AFSDB"
	DNSSummaryProtocolParamsQueryTypeAny        DNSSummaryProtocolParamsQueryType = "ANY"
	DNSSummaryProtocolParamsQueryTypeApl        DNSSummaryProtocolParamsQueryType = "APL"
	DNSSummaryProtocolParamsQueryTypeAtma       DNSSummaryProtocolParamsQueryType = "ATMA"
	DNSSummaryProtocolParamsQueryTypeAXFR       DNSSummaryProtocolParamsQueryType = "AXFR"
	DNSSummaryProtocolParamsQueryTypeCAA        DNSSummaryProtocolParamsQueryType = "CAA"
	DNSSummaryProtocolParamsQueryTypeCdnskey    DNSSummaryProtocolParamsQueryType = "CDNSKEY"
	DNSSummaryProtocolParamsQueryTypeCds        DNSSummaryProtocolParamsQueryType = "CDS"
	DNSSummaryProtocolParamsQueryTypeCERT       DNSSummaryProtocolParamsQueryType = "CERT"
	DNSSummaryProtocolParamsQueryTypeCNAME      DNSSummaryProtocolParamsQueryType = "CNAME"
	DNSSummaryProtocolParamsQueryTypeCsync      DNSSummaryProtocolParamsQueryType = "CSYNC"
	DNSSummaryProtocolParamsQueryTypeDhcid      DNSSummaryProtocolParamsQueryType = "DHCID"
	DNSSummaryProtocolParamsQueryTypeDlv        DNSSummaryProtocolParamsQueryType = "DLV"
	DNSSummaryProtocolParamsQueryTypeDname      DNSSummaryProtocolParamsQueryType = "DNAME"
	DNSSummaryProtocolParamsQueryTypeDNSKEY     DNSSummaryProtocolParamsQueryType = "DNSKEY"
	DNSSummaryProtocolParamsQueryTypeDoa        DNSSummaryProtocolParamsQueryType = "DOA"
	DNSSummaryProtocolParamsQueryTypeDS         DNSSummaryProtocolParamsQueryType = "DS"
	DNSSummaryProtocolParamsQueryTypeEid        DNSSummaryProtocolParamsQueryType = "EID"
	DNSSummaryProtocolParamsQueryTypeEui48      DNSSummaryProtocolParamsQueryType = "EUI48"
	DNSSummaryProtocolParamsQueryTypeEui64      DNSSummaryProtocolParamsQueryType = "EUI64"
	DNSSummaryProtocolParamsQueryTypeGpos       DNSSummaryProtocolParamsQueryType = "GPOS"
	DNSSummaryProtocolParamsQueryTypeGid        DNSSummaryProtocolParamsQueryType = "GID"
	DNSSummaryProtocolParamsQueryTypeHinfo      DNSSummaryProtocolParamsQueryType = "HINFO"
	DNSSummaryProtocolParamsQueryTypeHip        DNSSummaryProtocolParamsQueryType = "HIP"
	DNSSummaryProtocolParamsQueryTypeHTTPS      DNSSummaryProtocolParamsQueryType = "HTTPS"
	DNSSummaryProtocolParamsQueryTypeIpseckey   DNSSummaryProtocolParamsQueryType = "IPSECKEY"
	DNSSummaryProtocolParamsQueryTypeIsdn       DNSSummaryProtocolParamsQueryType = "ISDN"
	DNSSummaryProtocolParamsQueryTypeIxfr       DNSSummaryProtocolParamsQueryType = "IXFR"
	DNSSummaryProtocolParamsQueryTypeKey        DNSSummaryProtocolParamsQueryType = "KEY"
	DNSSummaryProtocolParamsQueryTypeKx         DNSSummaryProtocolParamsQueryType = "KX"
	DNSSummaryProtocolParamsQueryTypeL32        DNSSummaryProtocolParamsQueryType = "L32"
	DNSSummaryProtocolParamsQueryTypeL64        DNSSummaryProtocolParamsQueryType = "L64"
	DNSSummaryProtocolParamsQueryTypeLOC        DNSSummaryProtocolParamsQueryType = "LOC"
	DNSSummaryProtocolParamsQueryTypeLp         DNSSummaryProtocolParamsQueryType = "LP"
	DNSSummaryProtocolParamsQueryTypeMaila      DNSSummaryProtocolParamsQueryType = "MAILA"
	DNSSummaryProtocolParamsQueryTypeMailb      DNSSummaryProtocolParamsQueryType = "MAILB"
	DNSSummaryProtocolParamsQueryTypeMB         DNSSummaryProtocolParamsQueryType = "MB"
	DNSSummaryProtocolParamsQueryTypeMd         DNSSummaryProtocolParamsQueryType = "MD"
	DNSSummaryProtocolParamsQueryTypeMf         DNSSummaryProtocolParamsQueryType = "MF"
	DNSSummaryProtocolParamsQueryTypeMg         DNSSummaryProtocolParamsQueryType = "MG"
	DNSSummaryProtocolParamsQueryTypeMinfo      DNSSummaryProtocolParamsQueryType = "MINFO"
	DNSSummaryProtocolParamsQueryTypeMr         DNSSummaryProtocolParamsQueryType = "MR"
	DNSSummaryProtocolParamsQueryTypeMX         DNSSummaryProtocolParamsQueryType = "MX"
	DNSSummaryProtocolParamsQueryTypeNAPTR      DNSSummaryProtocolParamsQueryType = "NAPTR"
	DNSSummaryProtocolParamsQueryTypeNb         DNSSummaryProtocolParamsQueryType = "NB"
	DNSSummaryProtocolParamsQueryTypeNbstat     DNSSummaryProtocolParamsQueryType = "NBSTAT"
	DNSSummaryProtocolParamsQueryTypeNid        DNSSummaryProtocolParamsQueryType = "NID"
	DNSSummaryProtocolParamsQueryTypeNimloc     DNSSummaryProtocolParamsQueryType = "NIMLOC"
	DNSSummaryProtocolParamsQueryTypeNinfo      DNSSummaryProtocolParamsQueryType = "NINFO"
	DNSSummaryProtocolParamsQueryTypeNS         DNSSummaryProtocolParamsQueryType = "NS"
	DNSSummaryProtocolParamsQueryTypeNsap       DNSSummaryProtocolParamsQueryType = "NSAP"
	DNSSummaryProtocolParamsQueryTypeNsec       DNSSummaryProtocolParamsQueryType = "NSEC"
	DNSSummaryProtocolParamsQueryTypeNsec3      DNSSummaryProtocolParamsQueryType = "NSEC3"
	DNSSummaryProtocolParamsQueryTypeNsec3Param DNSSummaryProtocolParamsQueryType = "NSEC3PARAM"
	DNSSummaryProtocolParamsQueryTypeNull       DNSSummaryProtocolParamsQueryType = "NULL"
	DNSSummaryProtocolParamsQueryTypeNxt        DNSSummaryProtocolParamsQueryType = "NXT"
	DNSSummaryProtocolParamsQueryTypeOpenpgpkey DNSSummaryProtocolParamsQueryType = "OPENPGPKEY"
	DNSSummaryProtocolParamsQueryTypeOpt        DNSSummaryProtocolParamsQueryType = "OPT"
	DNSSummaryProtocolParamsQueryTypePTR        DNSSummaryProtocolParamsQueryType = "PTR"
	DNSSummaryProtocolParamsQueryTypePx         DNSSummaryProtocolParamsQueryType = "PX"
	DNSSummaryProtocolParamsQueryTypeRkey       DNSSummaryProtocolParamsQueryType = "RKEY"
	DNSSummaryProtocolParamsQueryTypeRp         DNSSummaryProtocolParamsQueryType = "RP"
	DNSSummaryProtocolParamsQueryTypeRrsig      DNSSummaryProtocolParamsQueryType = "RRSIG"
	DNSSummaryProtocolParamsQueryTypeRt         DNSSummaryProtocolParamsQueryType = "RT"
	DNSSummaryProtocolParamsQueryTypeSig        DNSSummaryProtocolParamsQueryType = "SIG"
	DNSSummaryProtocolParamsQueryTypeSink       DNSSummaryProtocolParamsQueryType = "SINK"
	DNSSummaryProtocolParamsQueryTypeSMIMEA     DNSSummaryProtocolParamsQueryType = "SMIMEA"
	DNSSummaryProtocolParamsQueryTypeSOA        DNSSummaryProtocolParamsQueryType = "SOA"
	DNSSummaryProtocolParamsQueryTypeSPF        DNSSummaryProtocolParamsQueryType = "SPF"
	DNSSummaryProtocolParamsQueryTypeSRV        DNSSummaryProtocolParamsQueryType = "SRV"
	DNSSummaryProtocolParamsQueryTypeSSHFP      DNSSummaryProtocolParamsQueryType = "SSHFP"
	DNSSummaryProtocolParamsQueryTypeSVCB       DNSSummaryProtocolParamsQueryType = "SVCB"
	DNSSummaryProtocolParamsQueryTypeTa         DNSSummaryProtocolParamsQueryType = "TA"
	DNSSummaryProtocolParamsQueryTypeTalink     DNSSummaryProtocolParamsQueryType = "TALINK"
	DNSSummaryProtocolParamsQueryTypeTkey       DNSSummaryProtocolParamsQueryType = "TKEY"
	DNSSummaryProtocolParamsQueryTypeTLSA       DNSSummaryProtocolParamsQueryType = "TLSA"
	DNSSummaryProtocolParamsQueryTypeTSIG       DNSSummaryProtocolParamsQueryType = "TSIG"
	DNSSummaryProtocolParamsQueryTypeTXT        DNSSummaryProtocolParamsQueryType = "TXT"
	DNSSummaryProtocolParamsQueryTypeUinfo      DNSSummaryProtocolParamsQueryType = "UINFO"
	DNSSummaryProtocolParamsQueryTypeUID        DNSSummaryProtocolParamsQueryType = "UID"
	DNSSummaryProtocolParamsQueryTypeUnspec     DNSSummaryProtocolParamsQueryType = "UNSPEC"
	DNSSummaryProtocolParamsQueryTypeURI        DNSSummaryProtocolParamsQueryType = "URI"
	DNSSummaryProtocolParamsQueryTypeWks        DNSSummaryProtocolParamsQueryType = "WKS"
	DNSSummaryProtocolParamsQueryTypeX25        DNSSummaryProtocolParamsQueryType = "X25"
	DNSSummaryProtocolParamsQueryTypeZonemd     DNSSummaryProtocolParamsQueryType = "ZONEMD"
)

func (r DNSSummaryProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryProtocolParamsQueryTypeA, DNSSummaryProtocolParamsQueryTypeAAAA, DNSSummaryProtocolParamsQueryTypeA6, DNSSummaryProtocolParamsQueryTypeAfsdb, DNSSummaryProtocolParamsQueryTypeAny, DNSSummaryProtocolParamsQueryTypeApl, DNSSummaryProtocolParamsQueryTypeAtma, DNSSummaryProtocolParamsQueryTypeAXFR, DNSSummaryProtocolParamsQueryTypeCAA, DNSSummaryProtocolParamsQueryTypeCdnskey, DNSSummaryProtocolParamsQueryTypeCds, DNSSummaryProtocolParamsQueryTypeCERT, DNSSummaryProtocolParamsQueryTypeCNAME, DNSSummaryProtocolParamsQueryTypeCsync, DNSSummaryProtocolParamsQueryTypeDhcid, DNSSummaryProtocolParamsQueryTypeDlv, DNSSummaryProtocolParamsQueryTypeDname, DNSSummaryProtocolParamsQueryTypeDNSKEY, DNSSummaryProtocolParamsQueryTypeDoa, DNSSummaryProtocolParamsQueryTypeDS, DNSSummaryProtocolParamsQueryTypeEid, DNSSummaryProtocolParamsQueryTypeEui48, DNSSummaryProtocolParamsQueryTypeEui64, DNSSummaryProtocolParamsQueryTypeGpos, DNSSummaryProtocolParamsQueryTypeGid, DNSSummaryProtocolParamsQueryTypeHinfo, DNSSummaryProtocolParamsQueryTypeHip, DNSSummaryProtocolParamsQueryTypeHTTPS, DNSSummaryProtocolParamsQueryTypeIpseckey, DNSSummaryProtocolParamsQueryTypeIsdn, DNSSummaryProtocolParamsQueryTypeIxfr, DNSSummaryProtocolParamsQueryTypeKey, DNSSummaryProtocolParamsQueryTypeKx, DNSSummaryProtocolParamsQueryTypeL32, DNSSummaryProtocolParamsQueryTypeL64, DNSSummaryProtocolParamsQueryTypeLOC, DNSSummaryProtocolParamsQueryTypeLp, DNSSummaryProtocolParamsQueryTypeMaila, DNSSummaryProtocolParamsQueryTypeMailb, DNSSummaryProtocolParamsQueryTypeMB, DNSSummaryProtocolParamsQueryTypeMd, DNSSummaryProtocolParamsQueryTypeMf, DNSSummaryProtocolParamsQueryTypeMg, DNSSummaryProtocolParamsQueryTypeMinfo, DNSSummaryProtocolParamsQueryTypeMr, DNSSummaryProtocolParamsQueryTypeMX, DNSSummaryProtocolParamsQueryTypeNAPTR, DNSSummaryProtocolParamsQueryTypeNb, DNSSummaryProtocolParamsQueryTypeNbstat, DNSSummaryProtocolParamsQueryTypeNid, DNSSummaryProtocolParamsQueryTypeNimloc, DNSSummaryProtocolParamsQueryTypeNinfo, DNSSummaryProtocolParamsQueryTypeNS, DNSSummaryProtocolParamsQueryTypeNsap, DNSSummaryProtocolParamsQueryTypeNsec, DNSSummaryProtocolParamsQueryTypeNsec3, DNSSummaryProtocolParamsQueryTypeNsec3Param, DNSSummaryProtocolParamsQueryTypeNull, DNSSummaryProtocolParamsQueryTypeNxt, DNSSummaryProtocolParamsQueryTypeOpenpgpkey, DNSSummaryProtocolParamsQueryTypeOpt, DNSSummaryProtocolParamsQueryTypePTR, DNSSummaryProtocolParamsQueryTypePx, DNSSummaryProtocolParamsQueryTypeRkey, DNSSummaryProtocolParamsQueryTypeRp, DNSSummaryProtocolParamsQueryTypeRrsig, DNSSummaryProtocolParamsQueryTypeRt, DNSSummaryProtocolParamsQueryTypeSig, DNSSummaryProtocolParamsQueryTypeSink, DNSSummaryProtocolParamsQueryTypeSMIMEA, DNSSummaryProtocolParamsQueryTypeSOA, DNSSummaryProtocolParamsQueryTypeSPF, DNSSummaryProtocolParamsQueryTypeSRV, DNSSummaryProtocolParamsQueryTypeSSHFP, DNSSummaryProtocolParamsQueryTypeSVCB, DNSSummaryProtocolParamsQueryTypeTa, DNSSummaryProtocolParamsQueryTypeTalink, DNSSummaryProtocolParamsQueryTypeTkey, DNSSummaryProtocolParamsQueryTypeTLSA, DNSSummaryProtocolParamsQueryTypeTSIG, DNSSummaryProtocolParamsQueryTypeTXT, DNSSummaryProtocolParamsQueryTypeUinfo, DNSSummaryProtocolParamsQueryTypeUID, DNSSummaryProtocolParamsQueryTypeUnspec, DNSSummaryProtocolParamsQueryTypeURI, DNSSummaryProtocolParamsQueryTypeWks, DNSSummaryProtocolParamsQueryTypeX25, DNSSummaryProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryProtocolParamsResponseCode string

const (
	DNSSummaryProtocolParamsResponseCodeNoerror   DNSSummaryProtocolParamsResponseCode = "NOERROR"
	DNSSummaryProtocolParamsResponseCodeFormerr   DNSSummaryProtocolParamsResponseCode = "FORMERR"
	DNSSummaryProtocolParamsResponseCodeServfail  DNSSummaryProtocolParamsResponseCode = "SERVFAIL"
	DNSSummaryProtocolParamsResponseCodeNxdomain  DNSSummaryProtocolParamsResponseCode = "NXDOMAIN"
	DNSSummaryProtocolParamsResponseCodeNotimp    DNSSummaryProtocolParamsResponseCode = "NOTIMP"
	DNSSummaryProtocolParamsResponseCodeRefused   DNSSummaryProtocolParamsResponseCode = "REFUSED"
	DNSSummaryProtocolParamsResponseCodeYxdomain  DNSSummaryProtocolParamsResponseCode = "YXDOMAIN"
	DNSSummaryProtocolParamsResponseCodeYxrrset   DNSSummaryProtocolParamsResponseCode = "YXRRSET"
	DNSSummaryProtocolParamsResponseCodeNxrrset   DNSSummaryProtocolParamsResponseCode = "NXRRSET"
	DNSSummaryProtocolParamsResponseCodeNotauth   DNSSummaryProtocolParamsResponseCode = "NOTAUTH"
	DNSSummaryProtocolParamsResponseCodeNotzone   DNSSummaryProtocolParamsResponseCode = "NOTZONE"
	DNSSummaryProtocolParamsResponseCodeBadsig    DNSSummaryProtocolParamsResponseCode = "BADSIG"
	DNSSummaryProtocolParamsResponseCodeBadkey    DNSSummaryProtocolParamsResponseCode = "BADKEY"
	DNSSummaryProtocolParamsResponseCodeBadtime   DNSSummaryProtocolParamsResponseCode = "BADTIME"
	DNSSummaryProtocolParamsResponseCodeBadmode   DNSSummaryProtocolParamsResponseCode = "BADMODE"
	DNSSummaryProtocolParamsResponseCodeBadname   DNSSummaryProtocolParamsResponseCode = "BADNAME"
	DNSSummaryProtocolParamsResponseCodeBadalg    DNSSummaryProtocolParamsResponseCode = "BADALG"
	DNSSummaryProtocolParamsResponseCodeBadtrunc  DNSSummaryProtocolParamsResponseCode = "BADTRUNC"
	DNSSummaryProtocolParamsResponseCodeBadcookie DNSSummaryProtocolParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryProtocolParamsResponseCodeNoerror, DNSSummaryProtocolParamsResponseCodeFormerr, DNSSummaryProtocolParamsResponseCodeServfail, DNSSummaryProtocolParamsResponseCodeNxdomain, DNSSummaryProtocolParamsResponseCodeNotimp, DNSSummaryProtocolParamsResponseCodeRefused, DNSSummaryProtocolParamsResponseCodeYxdomain, DNSSummaryProtocolParamsResponseCodeYxrrset, DNSSummaryProtocolParamsResponseCodeNxrrset, DNSSummaryProtocolParamsResponseCodeNotauth, DNSSummaryProtocolParamsResponseCodeNotzone, DNSSummaryProtocolParamsResponseCodeBadsig, DNSSummaryProtocolParamsResponseCodeBadkey, DNSSummaryProtocolParamsResponseCodeBadtime, DNSSummaryProtocolParamsResponseCodeBadmode, DNSSummaryProtocolParamsResponseCodeBadname, DNSSummaryProtocolParamsResponseCodeBadalg, DNSSummaryProtocolParamsResponseCodeBadtrunc, DNSSummaryProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryProtocolResponseEnvelope struct {
	Result  DNSSummaryProtocolResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    dnsSummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseEnvelope]
type dnsSummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryQueryTypeParams struct {
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
	Format param.Field[DNSSummaryQueryTypeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryQueryTypeParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryQueryTypeParamsFormat string

const (
	DNSSummaryQueryTypeParamsFormatJson DNSSummaryQueryTypeParamsFormat = "JSON"
	DNSSummaryQueryTypeParamsFormatCsv  DNSSummaryQueryTypeParamsFormat = "CSV"
)

func (r DNSSummaryQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryQueryTypeParamsFormatJson, DNSSummaryQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryQueryTypeParamsProtocol string

const (
	DNSSummaryQueryTypeParamsProtocolUdp   DNSSummaryQueryTypeParamsProtocol = "UDP"
	DNSSummaryQueryTypeParamsProtocolTCP   DNSSummaryQueryTypeParamsProtocol = "TCP"
	DNSSummaryQueryTypeParamsProtocolHTTPS DNSSummaryQueryTypeParamsProtocol = "HTTPS"
	DNSSummaryQueryTypeParamsProtocolTLS   DNSSummaryQueryTypeParamsProtocol = "TLS"
)

func (r DNSSummaryQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryQueryTypeParamsProtocolUdp, DNSSummaryQueryTypeParamsProtocolTCP, DNSSummaryQueryTypeParamsProtocolHTTPS, DNSSummaryQueryTypeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryQueryTypeParamsResponseCode string

const (
	DNSSummaryQueryTypeParamsResponseCodeNoerror   DNSSummaryQueryTypeParamsResponseCode = "NOERROR"
	DNSSummaryQueryTypeParamsResponseCodeFormerr   DNSSummaryQueryTypeParamsResponseCode = "FORMERR"
	DNSSummaryQueryTypeParamsResponseCodeServfail  DNSSummaryQueryTypeParamsResponseCode = "SERVFAIL"
	DNSSummaryQueryTypeParamsResponseCodeNxdomain  DNSSummaryQueryTypeParamsResponseCode = "NXDOMAIN"
	DNSSummaryQueryTypeParamsResponseCodeNotimp    DNSSummaryQueryTypeParamsResponseCode = "NOTIMP"
	DNSSummaryQueryTypeParamsResponseCodeRefused   DNSSummaryQueryTypeParamsResponseCode = "REFUSED"
	DNSSummaryQueryTypeParamsResponseCodeYxdomain  DNSSummaryQueryTypeParamsResponseCode = "YXDOMAIN"
	DNSSummaryQueryTypeParamsResponseCodeYxrrset   DNSSummaryQueryTypeParamsResponseCode = "YXRRSET"
	DNSSummaryQueryTypeParamsResponseCodeNxrrset   DNSSummaryQueryTypeParamsResponseCode = "NXRRSET"
	DNSSummaryQueryTypeParamsResponseCodeNotauth   DNSSummaryQueryTypeParamsResponseCode = "NOTAUTH"
	DNSSummaryQueryTypeParamsResponseCodeNotzone   DNSSummaryQueryTypeParamsResponseCode = "NOTZONE"
	DNSSummaryQueryTypeParamsResponseCodeBadsig    DNSSummaryQueryTypeParamsResponseCode = "BADSIG"
	DNSSummaryQueryTypeParamsResponseCodeBadkey    DNSSummaryQueryTypeParamsResponseCode = "BADKEY"
	DNSSummaryQueryTypeParamsResponseCodeBadtime   DNSSummaryQueryTypeParamsResponseCode = "BADTIME"
	DNSSummaryQueryTypeParamsResponseCodeBadmode   DNSSummaryQueryTypeParamsResponseCode = "BADMODE"
	DNSSummaryQueryTypeParamsResponseCodeBadname   DNSSummaryQueryTypeParamsResponseCode = "BADNAME"
	DNSSummaryQueryTypeParamsResponseCodeBadalg    DNSSummaryQueryTypeParamsResponseCode = "BADALG"
	DNSSummaryQueryTypeParamsResponseCodeBadtrunc  DNSSummaryQueryTypeParamsResponseCode = "BADTRUNC"
	DNSSummaryQueryTypeParamsResponseCodeBadcookie DNSSummaryQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryQueryTypeParamsResponseCodeNoerror, DNSSummaryQueryTypeParamsResponseCodeFormerr, DNSSummaryQueryTypeParamsResponseCodeServfail, DNSSummaryQueryTypeParamsResponseCodeNxdomain, DNSSummaryQueryTypeParamsResponseCodeNotimp, DNSSummaryQueryTypeParamsResponseCodeRefused, DNSSummaryQueryTypeParamsResponseCodeYxdomain, DNSSummaryQueryTypeParamsResponseCodeYxrrset, DNSSummaryQueryTypeParamsResponseCodeNxrrset, DNSSummaryQueryTypeParamsResponseCodeNotauth, DNSSummaryQueryTypeParamsResponseCodeNotzone, DNSSummaryQueryTypeParamsResponseCodeBadsig, DNSSummaryQueryTypeParamsResponseCodeBadkey, DNSSummaryQueryTypeParamsResponseCodeBadtime, DNSSummaryQueryTypeParamsResponseCodeBadmode, DNSSummaryQueryTypeParamsResponseCodeBadname, DNSSummaryQueryTypeParamsResponseCodeBadalg, DNSSummaryQueryTypeParamsResponseCodeBadtrunc, DNSSummaryQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryQueryTypeResponseEnvelope struct {
	Result  DNSSummaryQueryTypeResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    dnsSummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseEnvelope]
type dnsSummaryQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseCodeParams struct {
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
	Format param.Field[DNSSummaryResponseCodeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryResponseCodeParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryResponseCodeParamsQueryType] `query:"queryType"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryResponseCodeParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryResponseCodeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryResponseCodeParamsFormat string

const (
	DNSSummaryResponseCodeParamsFormatJson DNSSummaryResponseCodeParamsFormat = "JSON"
	DNSSummaryResponseCodeParamsFormatCsv  DNSSummaryResponseCodeParamsFormat = "CSV"
)

func (r DNSSummaryResponseCodeParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryResponseCodeParamsFormatJson, DNSSummaryResponseCodeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryResponseCodeParamsProtocol string

const (
	DNSSummaryResponseCodeParamsProtocolUdp   DNSSummaryResponseCodeParamsProtocol = "UDP"
	DNSSummaryResponseCodeParamsProtocolTCP   DNSSummaryResponseCodeParamsProtocol = "TCP"
	DNSSummaryResponseCodeParamsProtocolHTTPS DNSSummaryResponseCodeParamsProtocol = "HTTPS"
	DNSSummaryResponseCodeParamsProtocolTLS   DNSSummaryResponseCodeParamsProtocol = "TLS"
)

func (r DNSSummaryResponseCodeParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryResponseCodeParamsProtocolUdp, DNSSummaryResponseCodeParamsProtocolTCP, DNSSummaryResponseCodeParamsProtocolHTTPS, DNSSummaryResponseCodeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryResponseCodeParamsQueryType string

const (
	DNSSummaryResponseCodeParamsQueryTypeA          DNSSummaryResponseCodeParamsQueryType = "A"
	DNSSummaryResponseCodeParamsQueryTypeAAAA       DNSSummaryResponseCodeParamsQueryType = "AAAA"
	DNSSummaryResponseCodeParamsQueryTypeA6         DNSSummaryResponseCodeParamsQueryType = "A6"
	DNSSummaryResponseCodeParamsQueryTypeAfsdb      DNSSummaryResponseCodeParamsQueryType = "AFSDB"
	DNSSummaryResponseCodeParamsQueryTypeAny        DNSSummaryResponseCodeParamsQueryType = "ANY"
	DNSSummaryResponseCodeParamsQueryTypeApl        DNSSummaryResponseCodeParamsQueryType = "APL"
	DNSSummaryResponseCodeParamsQueryTypeAtma       DNSSummaryResponseCodeParamsQueryType = "ATMA"
	DNSSummaryResponseCodeParamsQueryTypeAXFR       DNSSummaryResponseCodeParamsQueryType = "AXFR"
	DNSSummaryResponseCodeParamsQueryTypeCAA        DNSSummaryResponseCodeParamsQueryType = "CAA"
	DNSSummaryResponseCodeParamsQueryTypeCdnskey    DNSSummaryResponseCodeParamsQueryType = "CDNSKEY"
	DNSSummaryResponseCodeParamsQueryTypeCds        DNSSummaryResponseCodeParamsQueryType = "CDS"
	DNSSummaryResponseCodeParamsQueryTypeCERT       DNSSummaryResponseCodeParamsQueryType = "CERT"
	DNSSummaryResponseCodeParamsQueryTypeCNAME      DNSSummaryResponseCodeParamsQueryType = "CNAME"
	DNSSummaryResponseCodeParamsQueryTypeCsync      DNSSummaryResponseCodeParamsQueryType = "CSYNC"
	DNSSummaryResponseCodeParamsQueryTypeDhcid      DNSSummaryResponseCodeParamsQueryType = "DHCID"
	DNSSummaryResponseCodeParamsQueryTypeDlv        DNSSummaryResponseCodeParamsQueryType = "DLV"
	DNSSummaryResponseCodeParamsQueryTypeDname      DNSSummaryResponseCodeParamsQueryType = "DNAME"
	DNSSummaryResponseCodeParamsQueryTypeDNSKEY     DNSSummaryResponseCodeParamsQueryType = "DNSKEY"
	DNSSummaryResponseCodeParamsQueryTypeDoa        DNSSummaryResponseCodeParamsQueryType = "DOA"
	DNSSummaryResponseCodeParamsQueryTypeDS         DNSSummaryResponseCodeParamsQueryType = "DS"
	DNSSummaryResponseCodeParamsQueryTypeEid        DNSSummaryResponseCodeParamsQueryType = "EID"
	DNSSummaryResponseCodeParamsQueryTypeEui48      DNSSummaryResponseCodeParamsQueryType = "EUI48"
	DNSSummaryResponseCodeParamsQueryTypeEui64      DNSSummaryResponseCodeParamsQueryType = "EUI64"
	DNSSummaryResponseCodeParamsQueryTypeGpos       DNSSummaryResponseCodeParamsQueryType = "GPOS"
	DNSSummaryResponseCodeParamsQueryTypeGid        DNSSummaryResponseCodeParamsQueryType = "GID"
	DNSSummaryResponseCodeParamsQueryTypeHinfo      DNSSummaryResponseCodeParamsQueryType = "HINFO"
	DNSSummaryResponseCodeParamsQueryTypeHip        DNSSummaryResponseCodeParamsQueryType = "HIP"
	DNSSummaryResponseCodeParamsQueryTypeHTTPS      DNSSummaryResponseCodeParamsQueryType = "HTTPS"
	DNSSummaryResponseCodeParamsQueryTypeIpseckey   DNSSummaryResponseCodeParamsQueryType = "IPSECKEY"
	DNSSummaryResponseCodeParamsQueryTypeIsdn       DNSSummaryResponseCodeParamsQueryType = "ISDN"
	DNSSummaryResponseCodeParamsQueryTypeIxfr       DNSSummaryResponseCodeParamsQueryType = "IXFR"
	DNSSummaryResponseCodeParamsQueryTypeKey        DNSSummaryResponseCodeParamsQueryType = "KEY"
	DNSSummaryResponseCodeParamsQueryTypeKx         DNSSummaryResponseCodeParamsQueryType = "KX"
	DNSSummaryResponseCodeParamsQueryTypeL32        DNSSummaryResponseCodeParamsQueryType = "L32"
	DNSSummaryResponseCodeParamsQueryTypeL64        DNSSummaryResponseCodeParamsQueryType = "L64"
	DNSSummaryResponseCodeParamsQueryTypeLOC        DNSSummaryResponseCodeParamsQueryType = "LOC"
	DNSSummaryResponseCodeParamsQueryTypeLp         DNSSummaryResponseCodeParamsQueryType = "LP"
	DNSSummaryResponseCodeParamsQueryTypeMaila      DNSSummaryResponseCodeParamsQueryType = "MAILA"
	DNSSummaryResponseCodeParamsQueryTypeMailb      DNSSummaryResponseCodeParamsQueryType = "MAILB"
	DNSSummaryResponseCodeParamsQueryTypeMB         DNSSummaryResponseCodeParamsQueryType = "MB"
	DNSSummaryResponseCodeParamsQueryTypeMd         DNSSummaryResponseCodeParamsQueryType = "MD"
	DNSSummaryResponseCodeParamsQueryTypeMf         DNSSummaryResponseCodeParamsQueryType = "MF"
	DNSSummaryResponseCodeParamsQueryTypeMg         DNSSummaryResponseCodeParamsQueryType = "MG"
	DNSSummaryResponseCodeParamsQueryTypeMinfo      DNSSummaryResponseCodeParamsQueryType = "MINFO"
	DNSSummaryResponseCodeParamsQueryTypeMr         DNSSummaryResponseCodeParamsQueryType = "MR"
	DNSSummaryResponseCodeParamsQueryTypeMX         DNSSummaryResponseCodeParamsQueryType = "MX"
	DNSSummaryResponseCodeParamsQueryTypeNAPTR      DNSSummaryResponseCodeParamsQueryType = "NAPTR"
	DNSSummaryResponseCodeParamsQueryTypeNb         DNSSummaryResponseCodeParamsQueryType = "NB"
	DNSSummaryResponseCodeParamsQueryTypeNbstat     DNSSummaryResponseCodeParamsQueryType = "NBSTAT"
	DNSSummaryResponseCodeParamsQueryTypeNid        DNSSummaryResponseCodeParamsQueryType = "NID"
	DNSSummaryResponseCodeParamsQueryTypeNimloc     DNSSummaryResponseCodeParamsQueryType = "NIMLOC"
	DNSSummaryResponseCodeParamsQueryTypeNinfo      DNSSummaryResponseCodeParamsQueryType = "NINFO"
	DNSSummaryResponseCodeParamsQueryTypeNS         DNSSummaryResponseCodeParamsQueryType = "NS"
	DNSSummaryResponseCodeParamsQueryTypeNsap       DNSSummaryResponseCodeParamsQueryType = "NSAP"
	DNSSummaryResponseCodeParamsQueryTypeNsec       DNSSummaryResponseCodeParamsQueryType = "NSEC"
	DNSSummaryResponseCodeParamsQueryTypeNsec3      DNSSummaryResponseCodeParamsQueryType = "NSEC3"
	DNSSummaryResponseCodeParamsQueryTypeNsec3Param DNSSummaryResponseCodeParamsQueryType = "NSEC3PARAM"
	DNSSummaryResponseCodeParamsQueryTypeNull       DNSSummaryResponseCodeParamsQueryType = "NULL"
	DNSSummaryResponseCodeParamsQueryTypeNxt        DNSSummaryResponseCodeParamsQueryType = "NXT"
	DNSSummaryResponseCodeParamsQueryTypeOpenpgpkey DNSSummaryResponseCodeParamsQueryType = "OPENPGPKEY"
	DNSSummaryResponseCodeParamsQueryTypeOpt        DNSSummaryResponseCodeParamsQueryType = "OPT"
	DNSSummaryResponseCodeParamsQueryTypePTR        DNSSummaryResponseCodeParamsQueryType = "PTR"
	DNSSummaryResponseCodeParamsQueryTypePx         DNSSummaryResponseCodeParamsQueryType = "PX"
	DNSSummaryResponseCodeParamsQueryTypeRkey       DNSSummaryResponseCodeParamsQueryType = "RKEY"
	DNSSummaryResponseCodeParamsQueryTypeRp         DNSSummaryResponseCodeParamsQueryType = "RP"
	DNSSummaryResponseCodeParamsQueryTypeRrsig      DNSSummaryResponseCodeParamsQueryType = "RRSIG"
	DNSSummaryResponseCodeParamsQueryTypeRt         DNSSummaryResponseCodeParamsQueryType = "RT"
	DNSSummaryResponseCodeParamsQueryTypeSig        DNSSummaryResponseCodeParamsQueryType = "SIG"
	DNSSummaryResponseCodeParamsQueryTypeSink       DNSSummaryResponseCodeParamsQueryType = "SINK"
	DNSSummaryResponseCodeParamsQueryTypeSMIMEA     DNSSummaryResponseCodeParamsQueryType = "SMIMEA"
	DNSSummaryResponseCodeParamsQueryTypeSOA        DNSSummaryResponseCodeParamsQueryType = "SOA"
	DNSSummaryResponseCodeParamsQueryTypeSPF        DNSSummaryResponseCodeParamsQueryType = "SPF"
	DNSSummaryResponseCodeParamsQueryTypeSRV        DNSSummaryResponseCodeParamsQueryType = "SRV"
	DNSSummaryResponseCodeParamsQueryTypeSSHFP      DNSSummaryResponseCodeParamsQueryType = "SSHFP"
	DNSSummaryResponseCodeParamsQueryTypeSVCB       DNSSummaryResponseCodeParamsQueryType = "SVCB"
	DNSSummaryResponseCodeParamsQueryTypeTa         DNSSummaryResponseCodeParamsQueryType = "TA"
	DNSSummaryResponseCodeParamsQueryTypeTalink     DNSSummaryResponseCodeParamsQueryType = "TALINK"
	DNSSummaryResponseCodeParamsQueryTypeTkey       DNSSummaryResponseCodeParamsQueryType = "TKEY"
	DNSSummaryResponseCodeParamsQueryTypeTLSA       DNSSummaryResponseCodeParamsQueryType = "TLSA"
	DNSSummaryResponseCodeParamsQueryTypeTSIG       DNSSummaryResponseCodeParamsQueryType = "TSIG"
	DNSSummaryResponseCodeParamsQueryTypeTXT        DNSSummaryResponseCodeParamsQueryType = "TXT"
	DNSSummaryResponseCodeParamsQueryTypeUinfo      DNSSummaryResponseCodeParamsQueryType = "UINFO"
	DNSSummaryResponseCodeParamsQueryTypeUID        DNSSummaryResponseCodeParamsQueryType = "UID"
	DNSSummaryResponseCodeParamsQueryTypeUnspec     DNSSummaryResponseCodeParamsQueryType = "UNSPEC"
	DNSSummaryResponseCodeParamsQueryTypeURI        DNSSummaryResponseCodeParamsQueryType = "URI"
	DNSSummaryResponseCodeParamsQueryTypeWks        DNSSummaryResponseCodeParamsQueryType = "WKS"
	DNSSummaryResponseCodeParamsQueryTypeX25        DNSSummaryResponseCodeParamsQueryType = "X25"
	DNSSummaryResponseCodeParamsQueryTypeZonemd     DNSSummaryResponseCodeParamsQueryType = "ZONEMD"
)

func (r DNSSummaryResponseCodeParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryResponseCodeParamsQueryTypeA, DNSSummaryResponseCodeParamsQueryTypeAAAA, DNSSummaryResponseCodeParamsQueryTypeA6, DNSSummaryResponseCodeParamsQueryTypeAfsdb, DNSSummaryResponseCodeParamsQueryTypeAny, DNSSummaryResponseCodeParamsQueryTypeApl, DNSSummaryResponseCodeParamsQueryTypeAtma, DNSSummaryResponseCodeParamsQueryTypeAXFR, DNSSummaryResponseCodeParamsQueryTypeCAA, DNSSummaryResponseCodeParamsQueryTypeCdnskey, DNSSummaryResponseCodeParamsQueryTypeCds, DNSSummaryResponseCodeParamsQueryTypeCERT, DNSSummaryResponseCodeParamsQueryTypeCNAME, DNSSummaryResponseCodeParamsQueryTypeCsync, DNSSummaryResponseCodeParamsQueryTypeDhcid, DNSSummaryResponseCodeParamsQueryTypeDlv, DNSSummaryResponseCodeParamsQueryTypeDname, DNSSummaryResponseCodeParamsQueryTypeDNSKEY, DNSSummaryResponseCodeParamsQueryTypeDoa, DNSSummaryResponseCodeParamsQueryTypeDS, DNSSummaryResponseCodeParamsQueryTypeEid, DNSSummaryResponseCodeParamsQueryTypeEui48, DNSSummaryResponseCodeParamsQueryTypeEui64, DNSSummaryResponseCodeParamsQueryTypeGpos, DNSSummaryResponseCodeParamsQueryTypeGid, DNSSummaryResponseCodeParamsQueryTypeHinfo, DNSSummaryResponseCodeParamsQueryTypeHip, DNSSummaryResponseCodeParamsQueryTypeHTTPS, DNSSummaryResponseCodeParamsQueryTypeIpseckey, DNSSummaryResponseCodeParamsQueryTypeIsdn, DNSSummaryResponseCodeParamsQueryTypeIxfr, DNSSummaryResponseCodeParamsQueryTypeKey, DNSSummaryResponseCodeParamsQueryTypeKx, DNSSummaryResponseCodeParamsQueryTypeL32, DNSSummaryResponseCodeParamsQueryTypeL64, DNSSummaryResponseCodeParamsQueryTypeLOC, DNSSummaryResponseCodeParamsQueryTypeLp, DNSSummaryResponseCodeParamsQueryTypeMaila, DNSSummaryResponseCodeParamsQueryTypeMailb, DNSSummaryResponseCodeParamsQueryTypeMB, DNSSummaryResponseCodeParamsQueryTypeMd, DNSSummaryResponseCodeParamsQueryTypeMf, DNSSummaryResponseCodeParamsQueryTypeMg, DNSSummaryResponseCodeParamsQueryTypeMinfo, DNSSummaryResponseCodeParamsQueryTypeMr, DNSSummaryResponseCodeParamsQueryTypeMX, DNSSummaryResponseCodeParamsQueryTypeNAPTR, DNSSummaryResponseCodeParamsQueryTypeNb, DNSSummaryResponseCodeParamsQueryTypeNbstat, DNSSummaryResponseCodeParamsQueryTypeNid, DNSSummaryResponseCodeParamsQueryTypeNimloc, DNSSummaryResponseCodeParamsQueryTypeNinfo, DNSSummaryResponseCodeParamsQueryTypeNS, DNSSummaryResponseCodeParamsQueryTypeNsap, DNSSummaryResponseCodeParamsQueryTypeNsec, DNSSummaryResponseCodeParamsQueryTypeNsec3, DNSSummaryResponseCodeParamsQueryTypeNsec3Param, DNSSummaryResponseCodeParamsQueryTypeNull, DNSSummaryResponseCodeParamsQueryTypeNxt, DNSSummaryResponseCodeParamsQueryTypeOpenpgpkey, DNSSummaryResponseCodeParamsQueryTypeOpt, DNSSummaryResponseCodeParamsQueryTypePTR, DNSSummaryResponseCodeParamsQueryTypePx, DNSSummaryResponseCodeParamsQueryTypeRkey, DNSSummaryResponseCodeParamsQueryTypeRp, DNSSummaryResponseCodeParamsQueryTypeRrsig, DNSSummaryResponseCodeParamsQueryTypeRt, DNSSummaryResponseCodeParamsQueryTypeSig, DNSSummaryResponseCodeParamsQueryTypeSink, DNSSummaryResponseCodeParamsQueryTypeSMIMEA, DNSSummaryResponseCodeParamsQueryTypeSOA, DNSSummaryResponseCodeParamsQueryTypeSPF, DNSSummaryResponseCodeParamsQueryTypeSRV, DNSSummaryResponseCodeParamsQueryTypeSSHFP, DNSSummaryResponseCodeParamsQueryTypeSVCB, DNSSummaryResponseCodeParamsQueryTypeTa, DNSSummaryResponseCodeParamsQueryTypeTalink, DNSSummaryResponseCodeParamsQueryTypeTkey, DNSSummaryResponseCodeParamsQueryTypeTLSA, DNSSummaryResponseCodeParamsQueryTypeTSIG, DNSSummaryResponseCodeParamsQueryTypeTXT, DNSSummaryResponseCodeParamsQueryTypeUinfo, DNSSummaryResponseCodeParamsQueryTypeUID, DNSSummaryResponseCodeParamsQueryTypeUnspec, DNSSummaryResponseCodeParamsQueryTypeURI, DNSSummaryResponseCodeParamsQueryTypeWks, DNSSummaryResponseCodeParamsQueryTypeX25, DNSSummaryResponseCodeParamsQueryTypeZonemd:
		return true
	}
	return false
}

type DNSSummaryResponseCodeResponseEnvelope struct {
	Result  DNSSummaryResponseCodeResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    dnsSummaryResponseCodeResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryResponseCodeResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryResponseCodeResponseEnvelope]
type dnsSummaryResponseCodeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSummaryResponseTTLParams struct {
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
	Format param.Field[DNSSummaryResponseTTLParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSSummaryResponseTTLParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSSummaryResponseTTLParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSSummaryResponseTTLParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryResponseTTLParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryResponseTTLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSSummaryResponseTTLParamsFormat string

const (
	DNSSummaryResponseTTLParamsFormatJson DNSSummaryResponseTTLParamsFormat = "JSON"
	DNSSummaryResponseTTLParamsFormatCsv  DNSSummaryResponseTTLParamsFormat = "CSV"
)

func (r DNSSummaryResponseTTLParamsFormat) IsKnown() bool {
	switch r {
	case DNSSummaryResponseTTLParamsFormatJson, DNSSummaryResponseTTLParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSSummaryResponseTTLParamsProtocol string

const (
	DNSSummaryResponseTTLParamsProtocolUdp   DNSSummaryResponseTTLParamsProtocol = "UDP"
	DNSSummaryResponseTTLParamsProtocolTCP   DNSSummaryResponseTTLParamsProtocol = "TCP"
	DNSSummaryResponseTTLParamsProtocolHTTPS DNSSummaryResponseTTLParamsProtocol = "HTTPS"
	DNSSummaryResponseTTLParamsProtocolTLS   DNSSummaryResponseTTLParamsProtocol = "TLS"
)

func (r DNSSummaryResponseTTLParamsProtocol) IsKnown() bool {
	switch r {
	case DNSSummaryResponseTTLParamsProtocolUdp, DNSSummaryResponseTTLParamsProtocolTCP, DNSSummaryResponseTTLParamsProtocolHTTPS, DNSSummaryResponseTTLParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSSummaryResponseTTLParamsQueryType string

const (
	DNSSummaryResponseTTLParamsQueryTypeA          DNSSummaryResponseTTLParamsQueryType = "A"
	DNSSummaryResponseTTLParamsQueryTypeAAAA       DNSSummaryResponseTTLParamsQueryType = "AAAA"
	DNSSummaryResponseTTLParamsQueryTypeA6         DNSSummaryResponseTTLParamsQueryType = "A6"
	DNSSummaryResponseTTLParamsQueryTypeAfsdb      DNSSummaryResponseTTLParamsQueryType = "AFSDB"
	DNSSummaryResponseTTLParamsQueryTypeAny        DNSSummaryResponseTTLParamsQueryType = "ANY"
	DNSSummaryResponseTTLParamsQueryTypeApl        DNSSummaryResponseTTLParamsQueryType = "APL"
	DNSSummaryResponseTTLParamsQueryTypeAtma       DNSSummaryResponseTTLParamsQueryType = "ATMA"
	DNSSummaryResponseTTLParamsQueryTypeAXFR       DNSSummaryResponseTTLParamsQueryType = "AXFR"
	DNSSummaryResponseTTLParamsQueryTypeCAA        DNSSummaryResponseTTLParamsQueryType = "CAA"
	DNSSummaryResponseTTLParamsQueryTypeCdnskey    DNSSummaryResponseTTLParamsQueryType = "CDNSKEY"
	DNSSummaryResponseTTLParamsQueryTypeCds        DNSSummaryResponseTTLParamsQueryType = "CDS"
	DNSSummaryResponseTTLParamsQueryTypeCERT       DNSSummaryResponseTTLParamsQueryType = "CERT"
	DNSSummaryResponseTTLParamsQueryTypeCNAME      DNSSummaryResponseTTLParamsQueryType = "CNAME"
	DNSSummaryResponseTTLParamsQueryTypeCsync      DNSSummaryResponseTTLParamsQueryType = "CSYNC"
	DNSSummaryResponseTTLParamsQueryTypeDhcid      DNSSummaryResponseTTLParamsQueryType = "DHCID"
	DNSSummaryResponseTTLParamsQueryTypeDlv        DNSSummaryResponseTTLParamsQueryType = "DLV"
	DNSSummaryResponseTTLParamsQueryTypeDname      DNSSummaryResponseTTLParamsQueryType = "DNAME"
	DNSSummaryResponseTTLParamsQueryTypeDNSKEY     DNSSummaryResponseTTLParamsQueryType = "DNSKEY"
	DNSSummaryResponseTTLParamsQueryTypeDoa        DNSSummaryResponseTTLParamsQueryType = "DOA"
	DNSSummaryResponseTTLParamsQueryTypeDS         DNSSummaryResponseTTLParamsQueryType = "DS"
	DNSSummaryResponseTTLParamsQueryTypeEid        DNSSummaryResponseTTLParamsQueryType = "EID"
	DNSSummaryResponseTTLParamsQueryTypeEui48      DNSSummaryResponseTTLParamsQueryType = "EUI48"
	DNSSummaryResponseTTLParamsQueryTypeEui64      DNSSummaryResponseTTLParamsQueryType = "EUI64"
	DNSSummaryResponseTTLParamsQueryTypeGpos       DNSSummaryResponseTTLParamsQueryType = "GPOS"
	DNSSummaryResponseTTLParamsQueryTypeGid        DNSSummaryResponseTTLParamsQueryType = "GID"
	DNSSummaryResponseTTLParamsQueryTypeHinfo      DNSSummaryResponseTTLParamsQueryType = "HINFO"
	DNSSummaryResponseTTLParamsQueryTypeHip        DNSSummaryResponseTTLParamsQueryType = "HIP"
	DNSSummaryResponseTTLParamsQueryTypeHTTPS      DNSSummaryResponseTTLParamsQueryType = "HTTPS"
	DNSSummaryResponseTTLParamsQueryTypeIpseckey   DNSSummaryResponseTTLParamsQueryType = "IPSECKEY"
	DNSSummaryResponseTTLParamsQueryTypeIsdn       DNSSummaryResponseTTLParamsQueryType = "ISDN"
	DNSSummaryResponseTTLParamsQueryTypeIxfr       DNSSummaryResponseTTLParamsQueryType = "IXFR"
	DNSSummaryResponseTTLParamsQueryTypeKey        DNSSummaryResponseTTLParamsQueryType = "KEY"
	DNSSummaryResponseTTLParamsQueryTypeKx         DNSSummaryResponseTTLParamsQueryType = "KX"
	DNSSummaryResponseTTLParamsQueryTypeL32        DNSSummaryResponseTTLParamsQueryType = "L32"
	DNSSummaryResponseTTLParamsQueryTypeL64        DNSSummaryResponseTTLParamsQueryType = "L64"
	DNSSummaryResponseTTLParamsQueryTypeLOC        DNSSummaryResponseTTLParamsQueryType = "LOC"
	DNSSummaryResponseTTLParamsQueryTypeLp         DNSSummaryResponseTTLParamsQueryType = "LP"
	DNSSummaryResponseTTLParamsQueryTypeMaila      DNSSummaryResponseTTLParamsQueryType = "MAILA"
	DNSSummaryResponseTTLParamsQueryTypeMailb      DNSSummaryResponseTTLParamsQueryType = "MAILB"
	DNSSummaryResponseTTLParamsQueryTypeMB         DNSSummaryResponseTTLParamsQueryType = "MB"
	DNSSummaryResponseTTLParamsQueryTypeMd         DNSSummaryResponseTTLParamsQueryType = "MD"
	DNSSummaryResponseTTLParamsQueryTypeMf         DNSSummaryResponseTTLParamsQueryType = "MF"
	DNSSummaryResponseTTLParamsQueryTypeMg         DNSSummaryResponseTTLParamsQueryType = "MG"
	DNSSummaryResponseTTLParamsQueryTypeMinfo      DNSSummaryResponseTTLParamsQueryType = "MINFO"
	DNSSummaryResponseTTLParamsQueryTypeMr         DNSSummaryResponseTTLParamsQueryType = "MR"
	DNSSummaryResponseTTLParamsQueryTypeMX         DNSSummaryResponseTTLParamsQueryType = "MX"
	DNSSummaryResponseTTLParamsQueryTypeNAPTR      DNSSummaryResponseTTLParamsQueryType = "NAPTR"
	DNSSummaryResponseTTLParamsQueryTypeNb         DNSSummaryResponseTTLParamsQueryType = "NB"
	DNSSummaryResponseTTLParamsQueryTypeNbstat     DNSSummaryResponseTTLParamsQueryType = "NBSTAT"
	DNSSummaryResponseTTLParamsQueryTypeNid        DNSSummaryResponseTTLParamsQueryType = "NID"
	DNSSummaryResponseTTLParamsQueryTypeNimloc     DNSSummaryResponseTTLParamsQueryType = "NIMLOC"
	DNSSummaryResponseTTLParamsQueryTypeNinfo      DNSSummaryResponseTTLParamsQueryType = "NINFO"
	DNSSummaryResponseTTLParamsQueryTypeNS         DNSSummaryResponseTTLParamsQueryType = "NS"
	DNSSummaryResponseTTLParamsQueryTypeNsap       DNSSummaryResponseTTLParamsQueryType = "NSAP"
	DNSSummaryResponseTTLParamsQueryTypeNsec       DNSSummaryResponseTTLParamsQueryType = "NSEC"
	DNSSummaryResponseTTLParamsQueryTypeNsec3      DNSSummaryResponseTTLParamsQueryType = "NSEC3"
	DNSSummaryResponseTTLParamsQueryTypeNsec3Param DNSSummaryResponseTTLParamsQueryType = "NSEC3PARAM"
	DNSSummaryResponseTTLParamsQueryTypeNull       DNSSummaryResponseTTLParamsQueryType = "NULL"
	DNSSummaryResponseTTLParamsQueryTypeNxt        DNSSummaryResponseTTLParamsQueryType = "NXT"
	DNSSummaryResponseTTLParamsQueryTypeOpenpgpkey DNSSummaryResponseTTLParamsQueryType = "OPENPGPKEY"
	DNSSummaryResponseTTLParamsQueryTypeOpt        DNSSummaryResponseTTLParamsQueryType = "OPT"
	DNSSummaryResponseTTLParamsQueryTypePTR        DNSSummaryResponseTTLParamsQueryType = "PTR"
	DNSSummaryResponseTTLParamsQueryTypePx         DNSSummaryResponseTTLParamsQueryType = "PX"
	DNSSummaryResponseTTLParamsQueryTypeRkey       DNSSummaryResponseTTLParamsQueryType = "RKEY"
	DNSSummaryResponseTTLParamsQueryTypeRp         DNSSummaryResponseTTLParamsQueryType = "RP"
	DNSSummaryResponseTTLParamsQueryTypeRrsig      DNSSummaryResponseTTLParamsQueryType = "RRSIG"
	DNSSummaryResponseTTLParamsQueryTypeRt         DNSSummaryResponseTTLParamsQueryType = "RT"
	DNSSummaryResponseTTLParamsQueryTypeSig        DNSSummaryResponseTTLParamsQueryType = "SIG"
	DNSSummaryResponseTTLParamsQueryTypeSink       DNSSummaryResponseTTLParamsQueryType = "SINK"
	DNSSummaryResponseTTLParamsQueryTypeSMIMEA     DNSSummaryResponseTTLParamsQueryType = "SMIMEA"
	DNSSummaryResponseTTLParamsQueryTypeSOA        DNSSummaryResponseTTLParamsQueryType = "SOA"
	DNSSummaryResponseTTLParamsQueryTypeSPF        DNSSummaryResponseTTLParamsQueryType = "SPF"
	DNSSummaryResponseTTLParamsQueryTypeSRV        DNSSummaryResponseTTLParamsQueryType = "SRV"
	DNSSummaryResponseTTLParamsQueryTypeSSHFP      DNSSummaryResponseTTLParamsQueryType = "SSHFP"
	DNSSummaryResponseTTLParamsQueryTypeSVCB       DNSSummaryResponseTTLParamsQueryType = "SVCB"
	DNSSummaryResponseTTLParamsQueryTypeTa         DNSSummaryResponseTTLParamsQueryType = "TA"
	DNSSummaryResponseTTLParamsQueryTypeTalink     DNSSummaryResponseTTLParamsQueryType = "TALINK"
	DNSSummaryResponseTTLParamsQueryTypeTkey       DNSSummaryResponseTTLParamsQueryType = "TKEY"
	DNSSummaryResponseTTLParamsQueryTypeTLSA       DNSSummaryResponseTTLParamsQueryType = "TLSA"
	DNSSummaryResponseTTLParamsQueryTypeTSIG       DNSSummaryResponseTTLParamsQueryType = "TSIG"
	DNSSummaryResponseTTLParamsQueryTypeTXT        DNSSummaryResponseTTLParamsQueryType = "TXT"
	DNSSummaryResponseTTLParamsQueryTypeUinfo      DNSSummaryResponseTTLParamsQueryType = "UINFO"
	DNSSummaryResponseTTLParamsQueryTypeUID        DNSSummaryResponseTTLParamsQueryType = "UID"
	DNSSummaryResponseTTLParamsQueryTypeUnspec     DNSSummaryResponseTTLParamsQueryType = "UNSPEC"
	DNSSummaryResponseTTLParamsQueryTypeURI        DNSSummaryResponseTTLParamsQueryType = "URI"
	DNSSummaryResponseTTLParamsQueryTypeWks        DNSSummaryResponseTTLParamsQueryType = "WKS"
	DNSSummaryResponseTTLParamsQueryTypeX25        DNSSummaryResponseTTLParamsQueryType = "X25"
	DNSSummaryResponseTTLParamsQueryTypeZonemd     DNSSummaryResponseTTLParamsQueryType = "ZONEMD"
)

func (r DNSSummaryResponseTTLParamsQueryType) IsKnown() bool {
	switch r {
	case DNSSummaryResponseTTLParamsQueryTypeA, DNSSummaryResponseTTLParamsQueryTypeAAAA, DNSSummaryResponseTTLParamsQueryTypeA6, DNSSummaryResponseTTLParamsQueryTypeAfsdb, DNSSummaryResponseTTLParamsQueryTypeAny, DNSSummaryResponseTTLParamsQueryTypeApl, DNSSummaryResponseTTLParamsQueryTypeAtma, DNSSummaryResponseTTLParamsQueryTypeAXFR, DNSSummaryResponseTTLParamsQueryTypeCAA, DNSSummaryResponseTTLParamsQueryTypeCdnskey, DNSSummaryResponseTTLParamsQueryTypeCds, DNSSummaryResponseTTLParamsQueryTypeCERT, DNSSummaryResponseTTLParamsQueryTypeCNAME, DNSSummaryResponseTTLParamsQueryTypeCsync, DNSSummaryResponseTTLParamsQueryTypeDhcid, DNSSummaryResponseTTLParamsQueryTypeDlv, DNSSummaryResponseTTLParamsQueryTypeDname, DNSSummaryResponseTTLParamsQueryTypeDNSKEY, DNSSummaryResponseTTLParamsQueryTypeDoa, DNSSummaryResponseTTLParamsQueryTypeDS, DNSSummaryResponseTTLParamsQueryTypeEid, DNSSummaryResponseTTLParamsQueryTypeEui48, DNSSummaryResponseTTLParamsQueryTypeEui64, DNSSummaryResponseTTLParamsQueryTypeGpos, DNSSummaryResponseTTLParamsQueryTypeGid, DNSSummaryResponseTTLParamsQueryTypeHinfo, DNSSummaryResponseTTLParamsQueryTypeHip, DNSSummaryResponseTTLParamsQueryTypeHTTPS, DNSSummaryResponseTTLParamsQueryTypeIpseckey, DNSSummaryResponseTTLParamsQueryTypeIsdn, DNSSummaryResponseTTLParamsQueryTypeIxfr, DNSSummaryResponseTTLParamsQueryTypeKey, DNSSummaryResponseTTLParamsQueryTypeKx, DNSSummaryResponseTTLParamsQueryTypeL32, DNSSummaryResponseTTLParamsQueryTypeL64, DNSSummaryResponseTTLParamsQueryTypeLOC, DNSSummaryResponseTTLParamsQueryTypeLp, DNSSummaryResponseTTLParamsQueryTypeMaila, DNSSummaryResponseTTLParamsQueryTypeMailb, DNSSummaryResponseTTLParamsQueryTypeMB, DNSSummaryResponseTTLParamsQueryTypeMd, DNSSummaryResponseTTLParamsQueryTypeMf, DNSSummaryResponseTTLParamsQueryTypeMg, DNSSummaryResponseTTLParamsQueryTypeMinfo, DNSSummaryResponseTTLParamsQueryTypeMr, DNSSummaryResponseTTLParamsQueryTypeMX, DNSSummaryResponseTTLParamsQueryTypeNAPTR, DNSSummaryResponseTTLParamsQueryTypeNb, DNSSummaryResponseTTLParamsQueryTypeNbstat, DNSSummaryResponseTTLParamsQueryTypeNid, DNSSummaryResponseTTLParamsQueryTypeNimloc, DNSSummaryResponseTTLParamsQueryTypeNinfo, DNSSummaryResponseTTLParamsQueryTypeNS, DNSSummaryResponseTTLParamsQueryTypeNsap, DNSSummaryResponseTTLParamsQueryTypeNsec, DNSSummaryResponseTTLParamsQueryTypeNsec3, DNSSummaryResponseTTLParamsQueryTypeNsec3Param, DNSSummaryResponseTTLParamsQueryTypeNull, DNSSummaryResponseTTLParamsQueryTypeNxt, DNSSummaryResponseTTLParamsQueryTypeOpenpgpkey, DNSSummaryResponseTTLParamsQueryTypeOpt, DNSSummaryResponseTTLParamsQueryTypePTR, DNSSummaryResponseTTLParamsQueryTypePx, DNSSummaryResponseTTLParamsQueryTypeRkey, DNSSummaryResponseTTLParamsQueryTypeRp, DNSSummaryResponseTTLParamsQueryTypeRrsig, DNSSummaryResponseTTLParamsQueryTypeRt, DNSSummaryResponseTTLParamsQueryTypeSig, DNSSummaryResponseTTLParamsQueryTypeSink, DNSSummaryResponseTTLParamsQueryTypeSMIMEA, DNSSummaryResponseTTLParamsQueryTypeSOA, DNSSummaryResponseTTLParamsQueryTypeSPF, DNSSummaryResponseTTLParamsQueryTypeSRV, DNSSummaryResponseTTLParamsQueryTypeSSHFP, DNSSummaryResponseTTLParamsQueryTypeSVCB, DNSSummaryResponseTTLParamsQueryTypeTa, DNSSummaryResponseTTLParamsQueryTypeTalink, DNSSummaryResponseTTLParamsQueryTypeTkey, DNSSummaryResponseTTLParamsQueryTypeTLSA, DNSSummaryResponseTTLParamsQueryTypeTSIG, DNSSummaryResponseTTLParamsQueryTypeTXT, DNSSummaryResponseTTLParamsQueryTypeUinfo, DNSSummaryResponseTTLParamsQueryTypeUID, DNSSummaryResponseTTLParamsQueryTypeUnspec, DNSSummaryResponseTTLParamsQueryTypeURI, DNSSummaryResponseTTLParamsQueryTypeWks, DNSSummaryResponseTTLParamsQueryTypeX25, DNSSummaryResponseTTLParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSSummaryResponseTTLParamsResponseCode string

const (
	DNSSummaryResponseTTLParamsResponseCodeNoerror   DNSSummaryResponseTTLParamsResponseCode = "NOERROR"
	DNSSummaryResponseTTLParamsResponseCodeFormerr   DNSSummaryResponseTTLParamsResponseCode = "FORMERR"
	DNSSummaryResponseTTLParamsResponseCodeServfail  DNSSummaryResponseTTLParamsResponseCode = "SERVFAIL"
	DNSSummaryResponseTTLParamsResponseCodeNxdomain  DNSSummaryResponseTTLParamsResponseCode = "NXDOMAIN"
	DNSSummaryResponseTTLParamsResponseCodeNotimp    DNSSummaryResponseTTLParamsResponseCode = "NOTIMP"
	DNSSummaryResponseTTLParamsResponseCodeRefused   DNSSummaryResponseTTLParamsResponseCode = "REFUSED"
	DNSSummaryResponseTTLParamsResponseCodeYxdomain  DNSSummaryResponseTTLParamsResponseCode = "YXDOMAIN"
	DNSSummaryResponseTTLParamsResponseCodeYxrrset   DNSSummaryResponseTTLParamsResponseCode = "YXRRSET"
	DNSSummaryResponseTTLParamsResponseCodeNxrrset   DNSSummaryResponseTTLParamsResponseCode = "NXRRSET"
	DNSSummaryResponseTTLParamsResponseCodeNotauth   DNSSummaryResponseTTLParamsResponseCode = "NOTAUTH"
	DNSSummaryResponseTTLParamsResponseCodeNotzone   DNSSummaryResponseTTLParamsResponseCode = "NOTZONE"
	DNSSummaryResponseTTLParamsResponseCodeBadsig    DNSSummaryResponseTTLParamsResponseCode = "BADSIG"
	DNSSummaryResponseTTLParamsResponseCodeBadkey    DNSSummaryResponseTTLParamsResponseCode = "BADKEY"
	DNSSummaryResponseTTLParamsResponseCodeBadtime   DNSSummaryResponseTTLParamsResponseCode = "BADTIME"
	DNSSummaryResponseTTLParamsResponseCodeBadmode   DNSSummaryResponseTTLParamsResponseCode = "BADMODE"
	DNSSummaryResponseTTLParamsResponseCodeBadname   DNSSummaryResponseTTLParamsResponseCode = "BADNAME"
	DNSSummaryResponseTTLParamsResponseCodeBadalg    DNSSummaryResponseTTLParamsResponseCode = "BADALG"
	DNSSummaryResponseTTLParamsResponseCodeBadtrunc  DNSSummaryResponseTTLParamsResponseCode = "BADTRUNC"
	DNSSummaryResponseTTLParamsResponseCodeBadcookie DNSSummaryResponseTTLParamsResponseCode = "BADCOOKIE"
)

func (r DNSSummaryResponseTTLParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSSummaryResponseTTLParamsResponseCodeNoerror, DNSSummaryResponseTTLParamsResponseCodeFormerr, DNSSummaryResponseTTLParamsResponseCodeServfail, DNSSummaryResponseTTLParamsResponseCodeNxdomain, DNSSummaryResponseTTLParamsResponseCodeNotimp, DNSSummaryResponseTTLParamsResponseCodeRefused, DNSSummaryResponseTTLParamsResponseCodeYxdomain, DNSSummaryResponseTTLParamsResponseCodeYxrrset, DNSSummaryResponseTTLParamsResponseCodeNxrrset, DNSSummaryResponseTTLParamsResponseCodeNotauth, DNSSummaryResponseTTLParamsResponseCodeNotzone, DNSSummaryResponseTTLParamsResponseCodeBadsig, DNSSummaryResponseTTLParamsResponseCodeBadkey, DNSSummaryResponseTTLParamsResponseCodeBadtime, DNSSummaryResponseTTLParamsResponseCodeBadmode, DNSSummaryResponseTTLParamsResponseCodeBadname, DNSSummaryResponseTTLParamsResponseCodeBadalg, DNSSummaryResponseTTLParamsResponseCodeBadtrunc, DNSSummaryResponseTTLParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSSummaryResponseTTLResponseEnvelope struct {
	Result  DNSSummaryResponseTTLResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    dnsSummaryResponseTTLResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryResponseTTLResponseEnvelope]
type dnsSummaryResponseTTLResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
