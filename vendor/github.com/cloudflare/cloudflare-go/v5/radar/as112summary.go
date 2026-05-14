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

// AS112SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112SummaryService] method instead.
type AS112SummaryService struct {
	Options []option.RequestOption
}

// NewAS112SummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAS112SummaryService(opts ...option.RequestOption) (r *AS112SummaryService) {
	r = &AS112SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries to AS112 by DNSSEC (DNS Security
// Extensions) support.
func (r *AS112SummaryService) DNSSEC(ctx context.Context, query AS112SummaryDNSSECParams, opts ...option.RequestOption) (res *AS112SummaryDNSSECResponse, err error) {
	var env AS112SummaryDNSSECResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries to AS112 by EDNS (Extension Mechanisms
// for DNS) support.
func (r *AS112SummaryService) Edns(ctx context.Context, query AS112SummaryEdnsParams, opts ...option.RequestOption) (res *AS112SummaryEdnsResponse, err error) {
	var env AS112SummaryEdnsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries to AS112 by IP version.
func (r *AS112SummaryService) IPVersion(ctx context.Context, query AS112SummaryIPVersionParams, opts ...option.RequestOption) (res *AS112SummaryIPVersionResponse, err error) {
	var env AS112SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries to AS112 by protocol.
func (r *AS112SummaryService) Protocol(ctx context.Context, query AS112SummaryProtocolParams, opts ...option.RequestOption) (res *AS112SummaryProtocolResponse, err error) {
	var env AS112SummaryProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries to AS112 by type.
func (r *AS112SummaryService) QueryType(ctx context.Context, query AS112SummaryQueryTypeParams, opts ...option.RequestOption) (res *AS112SummaryQueryTypeResponse, err error) {
	var env AS112SummaryQueryTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 DNS requests classified by response code.
func (r *AS112SummaryService) ResponseCodes(ctx context.Context, query AS112SummaryResponseCodesParams, opts ...option.RequestOption) (res *AS112SummaryResponseCodesResponse, err error) {
	var env AS112SummaryResponseCodesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112SummaryDNSSECResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryDNSSECResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryDNSSECResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryDNSSECResponseJSON     `json:"-"`
}

// as112SummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponse]
type as112SummaryDNSSECResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryDNSSECResponseMeta struct {
	ConfidenceInfo AS112SummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryDNSSECResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryDNSSECResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryDNSSECResponseMetaJSON   `json:"-"`
}

// as112SummaryDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseMeta]
type as112SummaryDNSSECResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  as112SummaryDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112SummaryDNSSECResponseMetaConfidenceInfo]
type as112SummaryDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryDNSSECResponseMetaDateRange]
type as112SummaryDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryDNSSECResponseMetaNormalization string

const (
	AS112SummaryDNSSECResponseMetaNormalizationPercentage           AS112SummaryDNSSECResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryDNSSECResponseMetaNormalizationMin0Max              AS112SummaryDNSSECResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryDNSSECResponseMetaNormalizationMinMax               AS112SummaryDNSSECResponseMetaNormalization = "MIN_MAX"
	AS112SummaryDNSSECResponseMetaNormalizationRawValues            AS112SummaryDNSSECResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryDNSSECResponseMetaNormalizationPercentageChange     AS112SummaryDNSSECResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryDNSSECResponseMetaNormalizationRollingAverage       AS112SummaryDNSSECResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryDNSSECResponseMetaNormalizationOverlappedPercentage AS112SummaryDNSSECResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryDNSSECResponseMetaNormalizationRatio                AS112SummaryDNSSECResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryDNSSECResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECResponseMetaNormalizationPercentage, AS112SummaryDNSSECResponseMetaNormalizationMin0Max, AS112SummaryDNSSECResponseMetaNormalizationMinMax, AS112SummaryDNSSECResponseMetaNormalizationRawValues, AS112SummaryDNSSECResponseMetaNormalizationPercentageChange, AS112SummaryDNSSECResponseMetaNormalizationRollingAverage, AS112SummaryDNSSECResponseMetaNormalizationOverlappedPercentage, AS112SummaryDNSSECResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryDNSSECResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  as112SummaryDNSSECResponseMetaUnitJSON `json:"-"`
}

// as112SummaryDNSSECResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseMetaUnit]
type as112SummaryDNSSECResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECResponseSummary0 struct {
	// A numeric string.
	NotSupported string `json:"NOT_SUPPORTED,required"`
	// A numeric string.
	Supported string                                 `json:"SUPPORTED,required"`
	JSON      as112SummaryDNSSECResponseSummary0JSON `json:"-"`
}

// as112SummaryDNSSECResponseSummary0JSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseSummary0]
type as112SummaryDNSSECResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryEdnsResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryEdnsResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryEdnsResponseJSON     `json:"-"`
}

// as112SummaryEdnsResponseJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponse]
type as112SummaryEdnsResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryEdnsResponseMeta struct {
	ConfidenceInfo AS112SummaryEdnsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryEdnsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryEdnsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryEdnsResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryEdnsResponseMetaJSON   `json:"-"`
}

// as112SummaryEdnsResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseMeta]
type as112SummaryEdnsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                          `json:"level,required"`
	JSON  as112SummaryEdnsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112SummaryEdnsResponseMetaConfidenceInfo]
type as112SummaryEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation]
type as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                 `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryEdnsResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryEdnsResponseMetaDateRange]
type as112SummaryEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryEdnsResponseMetaNormalization string

const (
	AS112SummaryEdnsResponseMetaNormalizationPercentage           AS112SummaryEdnsResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryEdnsResponseMetaNormalizationMin0Max              AS112SummaryEdnsResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryEdnsResponseMetaNormalizationMinMax               AS112SummaryEdnsResponseMetaNormalization = "MIN_MAX"
	AS112SummaryEdnsResponseMetaNormalizationRawValues            AS112SummaryEdnsResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryEdnsResponseMetaNormalizationPercentageChange     AS112SummaryEdnsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryEdnsResponseMetaNormalizationRollingAverage       AS112SummaryEdnsResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryEdnsResponseMetaNormalizationOverlappedPercentage AS112SummaryEdnsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryEdnsResponseMetaNormalizationRatio                AS112SummaryEdnsResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryEdnsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsResponseMetaNormalizationPercentage, AS112SummaryEdnsResponseMetaNormalizationMin0Max, AS112SummaryEdnsResponseMetaNormalizationMinMax, AS112SummaryEdnsResponseMetaNormalizationRawValues, AS112SummaryEdnsResponseMetaNormalizationPercentageChange, AS112SummaryEdnsResponseMetaNormalizationRollingAverage, AS112SummaryEdnsResponseMetaNormalizationOverlappedPercentage, AS112SummaryEdnsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryEdnsResponseMetaUnit struct {
	Name  string                               `json:"name,required"`
	Value string                               `json:"value,required"`
	JSON  as112SummaryEdnsResponseMetaUnitJSON `json:"-"`
}

// as112SummaryEdnsResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseMetaUnit]
type as112SummaryEdnsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsResponseSummary0 struct {
	// A numeric string.
	NotSupported string `json:"NOT_SUPPORTED,required"`
	// A numeric string.
	Supported string                               `json:"SUPPORTED,required"`
	JSON      as112SummaryEdnsResponseSummary0JSON `json:"-"`
}

// as112SummaryEdnsResponseSummary0JSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseSummary0]
type as112SummaryEdnsResponseSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryIPVersionResponseJSON     `json:"-"`
}

// as112SummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [AS112SummaryIPVersionResponse]
type as112SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryIPVersionResponseMeta struct {
	ConfidenceInfo AS112SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryIPVersionResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryIPVersionResponseMetaJSON   `json:"-"`
}

// as112SummaryIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryIPVersionResponseMeta]
type as112SummaryIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  as112SummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryIPVersionResponseMetaConfidenceInfo]
type as112SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112SummaryIPVersionResponseMetaDateRange]
type as112SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryIPVersionResponseMetaNormalization string

const (
	AS112SummaryIPVersionResponseMetaNormalizationPercentage           AS112SummaryIPVersionResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryIPVersionResponseMetaNormalizationMin0Max              AS112SummaryIPVersionResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryIPVersionResponseMetaNormalizationMinMax               AS112SummaryIPVersionResponseMetaNormalization = "MIN_MAX"
	AS112SummaryIPVersionResponseMetaNormalizationRawValues            AS112SummaryIPVersionResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryIPVersionResponseMetaNormalizationPercentageChange     AS112SummaryIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryIPVersionResponseMetaNormalizationRollingAverage       AS112SummaryIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryIPVersionResponseMetaNormalizationOverlappedPercentage AS112SummaryIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryIPVersionResponseMetaNormalizationRatio                AS112SummaryIPVersionResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionResponseMetaNormalizationPercentage, AS112SummaryIPVersionResponseMetaNormalizationMin0Max, AS112SummaryIPVersionResponseMetaNormalizationMinMax, AS112SummaryIPVersionResponseMetaNormalizationRawValues, AS112SummaryIPVersionResponseMetaNormalizationPercentageChange, AS112SummaryIPVersionResponseMetaNormalizationRollingAverage, AS112SummaryIPVersionResponseMetaNormalizationOverlappedPercentage, AS112SummaryIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryIPVersionResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  as112SummaryIPVersionResponseMetaUnitJSON `json:"-"`
}

// as112SummaryIPVersionResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112SummaryIPVersionResponseMetaUnit]
type as112SummaryIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionResponseSummary0 struct {
	// A numeric string.
	IPv4 string `json:"IPv4,required"`
	// A numeric string.
	IPv6 string                                    `json:"IPv6,required"`
	JSON as112SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// as112SummaryIPVersionResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryIPVersionResponseSummary0]
type as112SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 AS112SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     as112SummaryProtocolResponseJSON     `json:"-"`
}

// as112SummaryProtocolResponseJSON contains the JSON metadata for the struct
// [AS112SummaryProtocolResponse]
type as112SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryProtocolResponseMeta struct {
	ConfidenceInfo AS112SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryProtocolResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryProtocolResponseMetaJSON   `json:"-"`
}

// as112SummaryProtocolResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryProtocolResponseMeta]
type as112SummaryProtocolResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  as112SummaryProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryProtocolResponseMetaConfidenceInfo]
type as112SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseMetaDateRange]
type as112SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryProtocolResponseMetaNormalization string

const (
	AS112SummaryProtocolResponseMetaNormalizationPercentage           AS112SummaryProtocolResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryProtocolResponseMetaNormalizationMin0Max              AS112SummaryProtocolResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryProtocolResponseMetaNormalizationMinMax               AS112SummaryProtocolResponseMetaNormalization = "MIN_MAX"
	AS112SummaryProtocolResponseMetaNormalizationRawValues            AS112SummaryProtocolResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryProtocolResponseMetaNormalizationPercentageChange     AS112SummaryProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryProtocolResponseMetaNormalizationRollingAverage       AS112SummaryProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryProtocolResponseMetaNormalizationOverlappedPercentage AS112SummaryProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryProtocolResponseMetaNormalizationRatio                AS112SummaryProtocolResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryProtocolResponseMetaNormalizationPercentage, AS112SummaryProtocolResponseMetaNormalizationMin0Max, AS112SummaryProtocolResponseMetaNormalizationMinMax, AS112SummaryProtocolResponseMetaNormalizationRawValues, AS112SummaryProtocolResponseMetaNormalizationPercentageChange, AS112SummaryProtocolResponseMetaNormalizationRollingAverage, AS112SummaryProtocolResponseMetaNormalizationOverlappedPercentage, AS112SummaryProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryProtocolResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  as112SummaryProtocolResponseMetaUnitJSON `json:"-"`
}

// as112SummaryProtocolResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseMetaUnit]
type as112SummaryProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolResponseSummary0 struct {
	// A numeric string.
	HTTPS string `json:"HTTPS,required"`
	// A numeric string.
	TCP string `json:"TCP,required"`
	// A numeric string.
	TLS string `json:"TLS,required"`
	// A numeric string.
	Udp  string                                   `json:"UDP,required"`
	JSON as112SummaryProtocolResponseSummary0JSON `json:"-"`
}

// as112SummaryProtocolResponseSummary0JSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseSummary0]
type as112SummaryProtocolResponseSummary0JSON struct {
	HTTPS       apijson.Field
	TCP         apijson.Field
	TLS         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryQueryTypeResponseMeta `json:"meta,required"`
	Summary0 map[string]string                 `json:"summary_0,required"`
	JSON     as112SummaryQueryTypeResponseJSON `json:"-"`
}

// as112SummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [AS112SummaryQueryTypeResponse]
type as112SummaryQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryQueryTypeResponseMeta struct {
	ConfidenceInfo AS112SummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryQueryTypeResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryQueryTypeResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryQueryTypeResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryQueryTypeResponseMetaJSON   `json:"-"`
}

// as112SummaryQueryTypeResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryQueryTypeResponseMeta]
type as112SummaryQueryTypeResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  as112SummaryQueryTypeResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112SummaryQueryTypeResponseMetaConfidenceInfo]
type as112SummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112SummaryQueryTypeResponseMetaDateRange]
type as112SummaryQueryTypeResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryQueryTypeResponseMetaNormalization string

const (
	AS112SummaryQueryTypeResponseMetaNormalizationPercentage           AS112SummaryQueryTypeResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryQueryTypeResponseMetaNormalizationMin0Max              AS112SummaryQueryTypeResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryQueryTypeResponseMetaNormalizationMinMax               AS112SummaryQueryTypeResponseMetaNormalization = "MIN_MAX"
	AS112SummaryQueryTypeResponseMetaNormalizationRawValues            AS112SummaryQueryTypeResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryQueryTypeResponseMetaNormalizationPercentageChange     AS112SummaryQueryTypeResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryQueryTypeResponseMetaNormalizationRollingAverage       AS112SummaryQueryTypeResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryQueryTypeResponseMetaNormalizationOverlappedPercentage AS112SummaryQueryTypeResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryQueryTypeResponseMetaNormalizationRatio                AS112SummaryQueryTypeResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryQueryTypeResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryQueryTypeResponseMetaNormalizationPercentage, AS112SummaryQueryTypeResponseMetaNormalizationMin0Max, AS112SummaryQueryTypeResponseMetaNormalizationMinMax, AS112SummaryQueryTypeResponseMetaNormalizationRawValues, AS112SummaryQueryTypeResponseMetaNormalizationPercentageChange, AS112SummaryQueryTypeResponseMetaNormalizationRollingAverage, AS112SummaryQueryTypeResponseMetaNormalizationOverlappedPercentage, AS112SummaryQueryTypeResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryQueryTypeResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  as112SummaryQueryTypeResponseMetaUnitJSON `json:"-"`
}

// as112SummaryQueryTypeResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112SummaryQueryTypeResponseMetaUnit]
type as112SummaryQueryTypeResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponse struct {
	// Metadata for the results.
	Meta     AS112SummaryResponseCodesResponseMeta `json:"meta,required"`
	Summary0 map[string]string                     `json:"summary_0,required"`
	JSON     as112SummaryResponseCodesResponseJSON `json:"-"`
}

// as112SummaryResponseCodesResponseJSON contains the JSON metadata for the struct
// [AS112SummaryResponseCodesResponse]
type as112SummaryResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryResponseCodesResponseMeta struct {
	ConfidenceInfo AS112SummaryResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryResponseCodesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryResponseCodesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryResponseCodesResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryResponseCodesResponseMetaJSON   `json:"-"`
}

// as112SummaryResponseCodesResponseMetaJSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseMeta]
type as112SummaryResponseCodesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  as112SummaryResponseCodesResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AS112SummaryResponseCodesResponseMetaConfidenceInfo]
type as112SummaryResponseCodesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation]
type as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AS112SummaryResponseCodesResponseMetaDateRange]
type as112SummaryResponseCodesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryResponseCodesResponseMetaNormalization string

const (
	AS112SummaryResponseCodesResponseMetaNormalizationPercentage           AS112SummaryResponseCodesResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryResponseCodesResponseMetaNormalizationMin0Max              AS112SummaryResponseCodesResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryResponseCodesResponseMetaNormalizationMinMax               AS112SummaryResponseCodesResponseMetaNormalization = "MIN_MAX"
	AS112SummaryResponseCodesResponseMetaNormalizationRawValues            AS112SummaryResponseCodesResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryResponseCodesResponseMetaNormalizationPercentageChange     AS112SummaryResponseCodesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryResponseCodesResponseMetaNormalizationRollingAverage       AS112SummaryResponseCodesResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryResponseCodesResponseMetaNormalizationOverlappedPercentage AS112SummaryResponseCodesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryResponseCodesResponseMetaNormalizationRatio                AS112SummaryResponseCodesResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryResponseCodesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryResponseCodesResponseMetaNormalizationPercentage, AS112SummaryResponseCodesResponseMetaNormalizationMin0Max, AS112SummaryResponseCodesResponseMetaNormalizationMinMax, AS112SummaryResponseCodesResponseMetaNormalizationRawValues, AS112SummaryResponseCodesResponseMetaNormalizationPercentageChange, AS112SummaryResponseCodesResponseMetaNormalizationRollingAverage, AS112SummaryResponseCodesResponseMetaNormalizationOverlappedPercentage, AS112SummaryResponseCodesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryResponseCodesResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  as112SummaryResponseCodesResponseMetaUnitJSON `json:"-"`
}

// as112SummaryResponseCodesResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseMetaUnit]
type as112SummaryResponseCodesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryDNSSECParams struct {
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
	Format param.Field[AS112SummaryDNSSECParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112SummaryDNSSECParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112SummaryDNSSECParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112SummaryDNSSECParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryDNSSECParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryDNSSECParamsFormat string

const (
	AS112SummaryDNSSECParamsFormatJson AS112SummaryDNSSECParamsFormat = "JSON"
	AS112SummaryDNSSECParamsFormatCsv  AS112SummaryDNSSECParamsFormat = "CSV"
)

func (r AS112SummaryDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECParamsFormatJson, AS112SummaryDNSSECParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112SummaryDNSSECParamsProtocol string

const (
	AS112SummaryDNSSECParamsProtocolUdp   AS112SummaryDNSSECParamsProtocol = "UDP"
	AS112SummaryDNSSECParamsProtocolTCP   AS112SummaryDNSSECParamsProtocol = "TCP"
	AS112SummaryDNSSECParamsProtocolHTTPS AS112SummaryDNSSECParamsProtocol = "HTTPS"
	AS112SummaryDNSSECParamsProtocolTLS   AS112SummaryDNSSECParamsProtocol = "TLS"
)

func (r AS112SummaryDNSSECParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECParamsProtocolUdp, AS112SummaryDNSSECParamsProtocolTCP, AS112SummaryDNSSECParamsProtocolHTTPS, AS112SummaryDNSSECParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112SummaryDNSSECParamsQueryType string

const (
	AS112SummaryDNSSECParamsQueryTypeA          AS112SummaryDNSSECParamsQueryType = "A"
	AS112SummaryDNSSECParamsQueryTypeAAAA       AS112SummaryDNSSECParamsQueryType = "AAAA"
	AS112SummaryDNSSECParamsQueryTypeA6         AS112SummaryDNSSECParamsQueryType = "A6"
	AS112SummaryDNSSECParamsQueryTypeAfsdb      AS112SummaryDNSSECParamsQueryType = "AFSDB"
	AS112SummaryDNSSECParamsQueryTypeAny        AS112SummaryDNSSECParamsQueryType = "ANY"
	AS112SummaryDNSSECParamsQueryTypeApl        AS112SummaryDNSSECParamsQueryType = "APL"
	AS112SummaryDNSSECParamsQueryTypeAtma       AS112SummaryDNSSECParamsQueryType = "ATMA"
	AS112SummaryDNSSECParamsQueryTypeAXFR       AS112SummaryDNSSECParamsQueryType = "AXFR"
	AS112SummaryDNSSECParamsQueryTypeCAA        AS112SummaryDNSSECParamsQueryType = "CAA"
	AS112SummaryDNSSECParamsQueryTypeCdnskey    AS112SummaryDNSSECParamsQueryType = "CDNSKEY"
	AS112SummaryDNSSECParamsQueryTypeCds        AS112SummaryDNSSECParamsQueryType = "CDS"
	AS112SummaryDNSSECParamsQueryTypeCERT       AS112SummaryDNSSECParamsQueryType = "CERT"
	AS112SummaryDNSSECParamsQueryTypeCNAME      AS112SummaryDNSSECParamsQueryType = "CNAME"
	AS112SummaryDNSSECParamsQueryTypeCsync      AS112SummaryDNSSECParamsQueryType = "CSYNC"
	AS112SummaryDNSSECParamsQueryTypeDhcid      AS112SummaryDNSSECParamsQueryType = "DHCID"
	AS112SummaryDNSSECParamsQueryTypeDlv        AS112SummaryDNSSECParamsQueryType = "DLV"
	AS112SummaryDNSSECParamsQueryTypeDname      AS112SummaryDNSSECParamsQueryType = "DNAME"
	AS112SummaryDNSSECParamsQueryTypeDNSKEY     AS112SummaryDNSSECParamsQueryType = "DNSKEY"
	AS112SummaryDNSSECParamsQueryTypeDoa        AS112SummaryDNSSECParamsQueryType = "DOA"
	AS112SummaryDNSSECParamsQueryTypeDS         AS112SummaryDNSSECParamsQueryType = "DS"
	AS112SummaryDNSSECParamsQueryTypeEid        AS112SummaryDNSSECParamsQueryType = "EID"
	AS112SummaryDNSSECParamsQueryTypeEui48      AS112SummaryDNSSECParamsQueryType = "EUI48"
	AS112SummaryDNSSECParamsQueryTypeEui64      AS112SummaryDNSSECParamsQueryType = "EUI64"
	AS112SummaryDNSSECParamsQueryTypeGpos       AS112SummaryDNSSECParamsQueryType = "GPOS"
	AS112SummaryDNSSECParamsQueryTypeGid        AS112SummaryDNSSECParamsQueryType = "GID"
	AS112SummaryDNSSECParamsQueryTypeHinfo      AS112SummaryDNSSECParamsQueryType = "HINFO"
	AS112SummaryDNSSECParamsQueryTypeHip        AS112SummaryDNSSECParamsQueryType = "HIP"
	AS112SummaryDNSSECParamsQueryTypeHTTPS      AS112SummaryDNSSECParamsQueryType = "HTTPS"
	AS112SummaryDNSSECParamsQueryTypeIpseckey   AS112SummaryDNSSECParamsQueryType = "IPSECKEY"
	AS112SummaryDNSSECParamsQueryTypeIsdn       AS112SummaryDNSSECParamsQueryType = "ISDN"
	AS112SummaryDNSSECParamsQueryTypeIxfr       AS112SummaryDNSSECParamsQueryType = "IXFR"
	AS112SummaryDNSSECParamsQueryTypeKey        AS112SummaryDNSSECParamsQueryType = "KEY"
	AS112SummaryDNSSECParamsQueryTypeKx         AS112SummaryDNSSECParamsQueryType = "KX"
	AS112SummaryDNSSECParamsQueryTypeL32        AS112SummaryDNSSECParamsQueryType = "L32"
	AS112SummaryDNSSECParamsQueryTypeL64        AS112SummaryDNSSECParamsQueryType = "L64"
	AS112SummaryDNSSECParamsQueryTypeLOC        AS112SummaryDNSSECParamsQueryType = "LOC"
	AS112SummaryDNSSECParamsQueryTypeLp         AS112SummaryDNSSECParamsQueryType = "LP"
	AS112SummaryDNSSECParamsQueryTypeMaila      AS112SummaryDNSSECParamsQueryType = "MAILA"
	AS112SummaryDNSSECParamsQueryTypeMailb      AS112SummaryDNSSECParamsQueryType = "MAILB"
	AS112SummaryDNSSECParamsQueryTypeMB         AS112SummaryDNSSECParamsQueryType = "MB"
	AS112SummaryDNSSECParamsQueryTypeMd         AS112SummaryDNSSECParamsQueryType = "MD"
	AS112SummaryDNSSECParamsQueryTypeMf         AS112SummaryDNSSECParamsQueryType = "MF"
	AS112SummaryDNSSECParamsQueryTypeMg         AS112SummaryDNSSECParamsQueryType = "MG"
	AS112SummaryDNSSECParamsQueryTypeMinfo      AS112SummaryDNSSECParamsQueryType = "MINFO"
	AS112SummaryDNSSECParamsQueryTypeMr         AS112SummaryDNSSECParamsQueryType = "MR"
	AS112SummaryDNSSECParamsQueryTypeMX         AS112SummaryDNSSECParamsQueryType = "MX"
	AS112SummaryDNSSECParamsQueryTypeNAPTR      AS112SummaryDNSSECParamsQueryType = "NAPTR"
	AS112SummaryDNSSECParamsQueryTypeNb         AS112SummaryDNSSECParamsQueryType = "NB"
	AS112SummaryDNSSECParamsQueryTypeNbstat     AS112SummaryDNSSECParamsQueryType = "NBSTAT"
	AS112SummaryDNSSECParamsQueryTypeNid        AS112SummaryDNSSECParamsQueryType = "NID"
	AS112SummaryDNSSECParamsQueryTypeNimloc     AS112SummaryDNSSECParamsQueryType = "NIMLOC"
	AS112SummaryDNSSECParamsQueryTypeNinfo      AS112SummaryDNSSECParamsQueryType = "NINFO"
	AS112SummaryDNSSECParamsQueryTypeNS         AS112SummaryDNSSECParamsQueryType = "NS"
	AS112SummaryDNSSECParamsQueryTypeNsap       AS112SummaryDNSSECParamsQueryType = "NSAP"
	AS112SummaryDNSSECParamsQueryTypeNsec       AS112SummaryDNSSECParamsQueryType = "NSEC"
	AS112SummaryDNSSECParamsQueryTypeNsec3      AS112SummaryDNSSECParamsQueryType = "NSEC3"
	AS112SummaryDNSSECParamsQueryTypeNsec3Param AS112SummaryDNSSECParamsQueryType = "NSEC3PARAM"
	AS112SummaryDNSSECParamsQueryTypeNull       AS112SummaryDNSSECParamsQueryType = "NULL"
	AS112SummaryDNSSECParamsQueryTypeNxt        AS112SummaryDNSSECParamsQueryType = "NXT"
	AS112SummaryDNSSECParamsQueryTypeOpenpgpkey AS112SummaryDNSSECParamsQueryType = "OPENPGPKEY"
	AS112SummaryDNSSECParamsQueryTypeOpt        AS112SummaryDNSSECParamsQueryType = "OPT"
	AS112SummaryDNSSECParamsQueryTypePTR        AS112SummaryDNSSECParamsQueryType = "PTR"
	AS112SummaryDNSSECParamsQueryTypePx         AS112SummaryDNSSECParamsQueryType = "PX"
	AS112SummaryDNSSECParamsQueryTypeRkey       AS112SummaryDNSSECParamsQueryType = "RKEY"
	AS112SummaryDNSSECParamsQueryTypeRp         AS112SummaryDNSSECParamsQueryType = "RP"
	AS112SummaryDNSSECParamsQueryTypeRrsig      AS112SummaryDNSSECParamsQueryType = "RRSIG"
	AS112SummaryDNSSECParamsQueryTypeRt         AS112SummaryDNSSECParamsQueryType = "RT"
	AS112SummaryDNSSECParamsQueryTypeSig        AS112SummaryDNSSECParamsQueryType = "SIG"
	AS112SummaryDNSSECParamsQueryTypeSink       AS112SummaryDNSSECParamsQueryType = "SINK"
	AS112SummaryDNSSECParamsQueryTypeSMIMEA     AS112SummaryDNSSECParamsQueryType = "SMIMEA"
	AS112SummaryDNSSECParamsQueryTypeSOA        AS112SummaryDNSSECParamsQueryType = "SOA"
	AS112SummaryDNSSECParamsQueryTypeSPF        AS112SummaryDNSSECParamsQueryType = "SPF"
	AS112SummaryDNSSECParamsQueryTypeSRV        AS112SummaryDNSSECParamsQueryType = "SRV"
	AS112SummaryDNSSECParamsQueryTypeSSHFP      AS112SummaryDNSSECParamsQueryType = "SSHFP"
	AS112SummaryDNSSECParamsQueryTypeSVCB       AS112SummaryDNSSECParamsQueryType = "SVCB"
	AS112SummaryDNSSECParamsQueryTypeTa         AS112SummaryDNSSECParamsQueryType = "TA"
	AS112SummaryDNSSECParamsQueryTypeTalink     AS112SummaryDNSSECParamsQueryType = "TALINK"
	AS112SummaryDNSSECParamsQueryTypeTkey       AS112SummaryDNSSECParamsQueryType = "TKEY"
	AS112SummaryDNSSECParamsQueryTypeTLSA       AS112SummaryDNSSECParamsQueryType = "TLSA"
	AS112SummaryDNSSECParamsQueryTypeTSIG       AS112SummaryDNSSECParamsQueryType = "TSIG"
	AS112SummaryDNSSECParamsQueryTypeTXT        AS112SummaryDNSSECParamsQueryType = "TXT"
	AS112SummaryDNSSECParamsQueryTypeUinfo      AS112SummaryDNSSECParamsQueryType = "UINFO"
	AS112SummaryDNSSECParamsQueryTypeUID        AS112SummaryDNSSECParamsQueryType = "UID"
	AS112SummaryDNSSECParamsQueryTypeUnspec     AS112SummaryDNSSECParamsQueryType = "UNSPEC"
	AS112SummaryDNSSECParamsQueryTypeURI        AS112SummaryDNSSECParamsQueryType = "URI"
	AS112SummaryDNSSECParamsQueryTypeWks        AS112SummaryDNSSECParamsQueryType = "WKS"
	AS112SummaryDNSSECParamsQueryTypeX25        AS112SummaryDNSSECParamsQueryType = "X25"
	AS112SummaryDNSSECParamsQueryTypeZonemd     AS112SummaryDNSSECParamsQueryType = "ZONEMD"
)

func (r AS112SummaryDNSSECParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECParamsQueryTypeA, AS112SummaryDNSSECParamsQueryTypeAAAA, AS112SummaryDNSSECParamsQueryTypeA6, AS112SummaryDNSSECParamsQueryTypeAfsdb, AS112SummaryDNSSECParamsQueryTypeAny, AS112SummaryDNSSECParamsQueryTypeApl, AS112SummaryDNSSECParamsQueryTypeAtma, AS112SummaryDNSSECParamsQueryTypeAXFR, AS112SummaryDNSSECParamsQueryTypeCAA, AS112SummaryDNSSECParamsQueryTypeCdnskey, AS112SummaryDNSSECParamsQueryTypeCds, AS112SummaryDNSSECParamsQueryTypeCERT, AS112SummaryDNSSECParamsQueryTypeCNAME, AS112SummaryDNSSECParamsQueryTypeCsync, AS112SummaryDNSSECParamsQueryTypeDhcid, AS112SummaryDNSSECParamsQueryTypeDlv, AS112SummaryDNSSECParamsQueryTypeDname, AS112SummaryDNSSECParamsQueryTypeDNSKEY, AS112SummaryDNSSECParamsQueryTypeDoa, AS112SummaryDNSSECParamsQueryTypeDS, AS112SummaryDNSSECParamsQueryTypeEid, AS112SummaryDNSSECParamsQueryTypeEui48, AS112SummaryDNSSECParamsQueryTypeEui64, AS112SummaryDNSSECParamsQueryTypeGpos, AS112SummaryDNSSECParamsQueryTypeGid, AS112SummaryDNSSECParamsQueryTypeHinfo, AS112SummaryDNSSECParamsQueryTypeHip, AS112SummaryDNSSECParamsQueryTypeHTTPS, AS112SummaryDNSSECParamsQueryTypeIpseckey, AS112SummaryDNSSECParamsQueryTypeIsdn, AS112SummaryDNSSECParamsQueryTypeIxfr, AS112SummaryDNSSECParamsQueryTypeKey, AS112SummaryDNSSECParamsQueryTypeKx, AS112SummaryDNSSECParamsQueryTypeL32, AS112SummaryDNSSECParamsQueryTypeL64, AS112SummaryDNSSECParamsQueryTypeLOC, AS112SummaryDNSSECParamsQueryTypeLp, AS112SummaryDNSSECParamsQueryTypeMaila, AS112SummaryDNSSECParamsQueryTypeMailb, AS112SummaryDNSSECParamsQueryTypeMB, AS112SummaryDNSSECParamsQueryTypeMd, AS112SummaryDNSSECParamsQueryTypeMf, AS112SummaryDNSSECParamsQueryTypeMg, AS112SummaryDNSSECParamsQueryTypeMinfo, AS112SummaryDNSSECParamsQueryTypeMr, AS112SummaryDNSSECParamsQueryTypeMX, AS112SummaryDNSSECParamsQueryTypeNAPTR, AS112SummaryDNSSECParamsQueryTypeNb, AS112SummaryDNSSECParamsQueryTypeNbstat, AS112SummaryDNSSECParamsQueryTypeNid, AS112SummaryDNSSECParamsQueryTypeNimloc, AS112SummaryDNSSECParamsQueryTypeNinfo, AS112SummaryDNSSECParamsQueryTypeNS, AS112SummaryDNSSECParamsQueryTypeNsap, AS112SummaryDNSSECParamsQueryTypeNsec, AS112SummaryDNSSECParamsQueryTypeNsec3, AS112SummaryDNSSECParamsQueryTypeNsec3Param, AS112SummaryDNSSECParamsQueryTypeNull, AS112SummaryDNSSECParamsQueryTypeNxt, AS112SummaryDNSSECParamsQueryTypeOpenpgpkey, AS112SummaryDNSSECParamsQueryTypeOpt, AS112SummaryDNSSECParamsQueryTypePTR, AS112SummaryDNSSECParamsQueryTypePx, AS112SummaryDNSSECParamsQueryTypeRkey, AS112SummaryDNSSECParamsQueryTypeRp, AS112SummaryDNSSECParamsQueryTypeRrsig, AS112SummaryDNSSECParamsQueryTypeRt, AS112SummaryDNSSECParamsQueryTypeSig, AS112SummaryDNSSECParamsQueryTypeSink, AS112SummaryDNSSECParamsQueryTypeSMIMEA, AS112SummaryDNSSECParamsQueryTypeSOA, AS112SummaryDNSSECParamsQueryTypeSPF, AS112SummaryDNSSECParamsQueryTypeSRV, AS112SummaryDNSSECParamsQueryTypeSSHFP, AS112SummaryDNSSECParamsQueryTypeSVCB, AS112SummaryDNSSECParamsQueryTypeTa, AS112SummaryDNSSECParamsQueryTypeTalink, AS112SummaryDNSSECParamsQueryTypeTkey, AS112SummaryDNSSECParamsQueryTypeTLSA, AS112SummaryDNSSECParamsQueryTypeTSIG, AS112SummaryDNSSECParamsQueryTypeTXT, AS112SummaryDNSSECParamsQueryTypeUinfo, AS112SummaryDNSSECParamsQueryTypeUID, AS112SummaryDNSSECParamsQueryTypeUnspec, AS112SummaryDNSSECParamsQueryTypeURI, AS112SummaryDNSSECParamsQueryTypeWks, AS112SummaryDNSSECParamsQueryTypeX25, AS112SummaryDNSSECParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112SummaryDNSSECParamsResponseCode string

const (
	AS112SummaryDNSSECParamsResponseCodeNoerror   AS112SummaryDNSSECParamsResponseCode = "NOERROR"
	AS112SummaryDNSSECParamsResponseCodeFormerr   AS112SummaryDNSSECParamsResponseCode = "FORMERR"
	AS112SummaryDNSSECParamsResponseCodeServfail  AS112SummaryDNSSECParamsResponseCode = "SERVFAIL"
	AS112SummaryDNSSECParamsResponseCodeNxdomain  AS112SummaryDNSSECParamsResponseCode = "NXDOMAIN"
	AS112SummaryDNSSECParamsResponseCodeNotimp    AS112SummaryDNSSECParamsResponseCode = "NOTIMP"
	AS112SummaryDNSSECParamsResponseCodeRefused   AS112SummaryDNSSECParamsResponseCode = "REFUSED"
	AS112SummaryDNSSECParamsResponseCodeYxdomain  AS112SummaryDNSSECParamsResponseCode = "YXDOMAIN"
	AS112SummaryDNSSECParamsResponseCodeYxrrset   AS112SummaryDNSSECParamsResponseCode = "YXRRSET"
	AS112SummaryDNSSECParamsResponseCodeNxrrset   AS112SummaryDNSSECParamsResponseCode = "NXRRSET"
	AS112SummaryDNSSECParamsResponseCodeNotauth   AS112SummaryDNSSECParamsResponseCode = "NOTAUTH"
	AS112SummaryDNSSECParamsResponseCodeNotzone   AS112SummaryDNSSECParamsResponseCode = "NOTZONE"
	AS112SummaryDNSSECParamsResponseCodeBadsig    AS112SummaryDNSSECParamsResponseCode = "BADSIG"
	AS112SummaryDNSSECParamsResponseCodeBadkey    AS112SummaryDNSSECParamsResponseCode = "BADKEY"
	AS112SummaryDNSSECParamsResponseCodeBadtime   AS112SummaryDNSSECParamsResponseCode = "BADTIME"
	AS112SummaryDNSSECParamsResponseCodeBadmode   AS112SummaryDNSSECParamsResponseCode = "BADMODE"
	AS112SummaryDNSSECParamsResponseCodeBadname   AS112SummaryDNSSECParamsResponseCode = "BADNAME"
	AS112SummaryDNSSECParamsResponseCodeBadalg    AS112SummaryDNSSECParamsResponseCode = "BADALG"
	AS112SummaryDNSSECParamsResponseCodeBadtrunc  AS112SummaryDNSSECParamsResponseCode = "BADTRUNC"
	AS112SummaryDNSSECParamsResponseCodeBadcookie AS112SummaryDNSSECParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryDNSSECParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryDNSSECParamsResponseCodeNoerror, AS112SummaryDNSSECParamsResponseCodeFormerr, AS112SummaryDNSSECParamsResponseCodeServfail, AS112SummaryDNSSECParamsResponseCodeNxdomain, AS112SummaryDNSSECParamsResponseCodeNotimp, AS112SummaryDNSSECParamsResponseCodeRefused, AS112SummaryDNSSECParamsResponseCodeYxdomain, AS112SummaryDNSSECParamsResponseCodeYxrrset, AS112SummaryDNSSECParamsResponseCodeNxrrset, AS112SummaryDNSSECParamsResponseCodeNotauth, AS112SummaryDNSSECParamsResponseCodeNotzone, AS112SummaryDNSSECParamsResponseCodeBadsig, AS112SummaryDNSSECParamsResponseCodeBadkey, AS112SummaryDNSSECParamsResponseCodeBadtime, AS112SummaryDNSSECParamsResponseCodeBadmode, AS112SummaryDNSSECParamsResponseCodeBadname, AS112SummaryDNSSECParamsResponseCodeBadalg, AS112SummaryDNSSECParamsResponseCodeBadtrunc, AS112SummaryDNSSECParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryDNSSECResponseEnvelope struct {
	Result  AS112SummaryDNSSECResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    as112SummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112SummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112SummaryDNSSECResponseEnvelope]
type as112SummaryDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryEdnsParams struct {
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
	Format param.Field[AS112SummaryEdnsParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112SummaryEdnsParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112SummaryEdnsParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112SummaryEdnsParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryEdnsParams]'s query parameters as `url.Values`.
func (r AS112SummaryEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryEdnsParamsFormat string

const (
	AS112SummaryEdnsParamsFormatJson AS112SummaryEdnsParamsFormat = "JSON"
	AS112SummaryEdnsParamsFormatCsv  AS112SummaryEdnsParamsFormat = "CSV"
)

func (r AS112SummaryEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsParamsFormatJson, AS112SummaryEdnsParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112SummaryEdnsParamsProtocol string

const (
	AS112SummaryEdnsParamsProtocolUdp   AS112SummaryEdnsParamsProtocol = "UDP"
	AS112SummaryEdnsParamsProtocolTCP   AS112SummaryEdnsParamsProtocol = "TCP"
	AS112SummaryEdnsParamsProtocolHTTPS AS112SummaryEdnsParamsProtocol = "HTTPS"
	AS112SummaryEdnsParamsProtocolTLS   AS112SummaryEdnsParamsProtocol = "TLS"
)

func (r AS112SummaryEdnsParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsParamsProtocolUdp, AS112SummaryEdnsParamsProtocolTCP, AS112SummaryEdnsParamsProtocolHTTPS, AS112SummaryEdnsParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112SummaryEdnsParamsQueryType string

const (
	AS112SummaryEdnsParamsQueryTypeA          AS112SummaryEdnsParamsQueryType = "A"
	AS112SummaryEdnsParamsQueryTypeAAAA       AS112SummaryEdnsParamsQueryType = "AAAA"
	AS112SummaryEdnsParamsQueryTypeA6         AS112SummaryEdnsParamsQueryType = "A6"
	AS112SummaryEdnsParamsQueryTypeAfsdb      AS112SummaryEdnsParamsQueryType = "AFSDB"
	AS112SummaryEdnsParamsQueryTypeAny        AS112SummaryEdnsParamsQueryType = "ANY"
	AS112SummaryEdnsParamsQueryTypeApl        AS112SummaryEdnsParamsQueryType = "APL"
	AS112SummaryEdnsParamsQueryTypeAtma       AS112SummaryEdnsParamsQueryType = "ATMA"
	AS112SummaryEdnsParamsQueryTypeAXFR       AS112SummaryEdnsParamsQueryType = "AXFR"
	AS112SummaryEdnsParamsQueryTypeCAA        AS112SummaryEdnsParamsQueryType = "CAA"
	AS112SummaryEdnsParamsQueryTypeCdnskey    AS112SummaryEdnsParamsQueryType = "CDNSKEY"
	AS112SummaryEdnsParamsQueryTypeCds        AS112SummaryEdnsParamsQueryType = "CDS"
	AS112SummaryEdnsParamsQueryTypeCERT       AS112SummaryEdnsParamsQueryType = "CERT"
	AS112SummaryEdnsParamsQueryTypeCNAME      AS112SummaryEdnsParamsQueryType = "CNAME"
	AS112SummaryEdnsParamsQueryTypeCsync      AS112SummaryEdnsParamsQueryType = "CSYNC"
	AS112SummaryEdnsParamsQueryTypeDhcid      AS112SummaryEdnsParamsQueryType = "DHCID"
	AS112SummaryEdnsParamsQueryTypeDlv        AS112SummaryEdnsParamsQueryType = "DLV"
	AS112SummaryEdnsParamsQueryTypeDname      AS112SummaryEdnsParamsQueryType = "DNAME"
	AS112SummaryEdnsParamsQueryTypeDNSKEY     AS112SummaryEdnsParamsQueryType = "DNSKEY"
	AS112SummaryEdnsParamsQueryTypeDoa        AS112SummaryEdnsParamsQueryType = "DOA"
	AS112SummaryEdnsParamsQueryTypeDS         AS112SummaryEdnsParamsQueryType = "DS"
	AS112SummaryEdnsParamsQueryTypeEid        AS112SummaryEdnsParamsQueryType = "EID"
	AS112SummaryEdnsParamsQueryTypeEui48      AS112SummaryEdnsParamsQueryType = "EUI48"
	AS112SummaryEdnsParamsQueryTypeEui64      AS112SummaryEdnsParamsQueryType = "EUI64"
	AS112SummaryEdnsParamsQueryTypeGpos       AS112SummaryEdnsParamsQueryType = "GPOS"
	AS112SummaryEdnsParamsQueryTypeGid        AS112SummaryEdnsParamsQueryType = "GID"
	AS112SummaryEdnsParamsQueryTypeHinfo      AS112SummaryEdnsParamsQueryType = "HINFO"
	AS112SummaryEdnsParamsQueryTypeHip        AS112SummaryEdnsParamsQueryType = "HIP"
	AS112SummaryEdnsParamsQueryTypeHTTPS      AS112SummaryEdnsParamsQueryType = "HTTPS"
	AS112SummaryEdnsParamsQueryTypeIpseckey   AS112SummaryEdnsParamsQueryType = "IPSECKEY"
	AS112SummaryEdnsParamsQueryTypeIsdn       AS112SummaryEdnsParamsQueryType = "ISDN"
	AS112SummaryEdnsParamsQueryTypeIxfr       AS112SummaryEdnsParamsQueryType = "IXFR"
	AS112SummaryEdnsParamsQueryTypeKey        AS112SummaryEdnsParamsQueryType = "KEY"
	AS112SummaryEdnsParamsQueryTypeKx         AS112SummaryEdnsParamsQueryType = "KX"
	AS112SummaryEdnsParamsQueryTypeL32        AS112SummaryEdnsParamsQueryType = "L32"
	AS112SummaryEdnsParamsQueryTypeL64        AS112SummaryEdnsParamsQueryType = "L64"
	AS112SummaryEdnsParamsQueryTypeLOC        AS112SummaryEdnsParamsQueryType = "LOC"
	AS112SummaryEdnsParamsQueryTypeLp         AS112SummaryEdnsParamsQueryType = "LP"
	AS112SummaryEdnsParamsQueryTypeMaila      AS112SummaryEdnsParamsQueryType = "MAILA"
	AS112SummaryEdnsParamsQueryTypeMailb      AS112SummaryEdnsParamsQueryType = "MAILB"
	AS112SummaryEdnsParamsQueryTypeMB         AS112SummaryEdnsParamsQueryType = "MB"
	AS112SummaryEdnsParamsQueryTypeMd         AS112SummaryEdnsParamsQueryType = "MD"
	AS112SummaryEdnsParamsQueryTypeMf         AS112SummaryEdnsParamsQueryType = "MF"
	AS112SummaryEdnsParamsQueryTypeMg         AS112SummaryEdnsParamsQueryType = "MG"
	AS112SummaryEdnsParamsQueryTypeMinfo      AS112SummaryEdnsParamsQueryType = "MINFO"
	AS112SummaryEdnsParamsQueryTypeMr         AS112SummaryEdnsParamsQueryType = "MR"
	AS112SummaryEdnsParamsQueryTypeMX         AS112SummaryEdnsParamsQueryType = "MX"
	AS112SummaryEdnsParamsQueryTypeNAPTR      AS112SummaryEdnsParamsQueryType = "NAPTR"
	AS112SummaryEdnsParamsQueryTypeNb         AS112SummaryEdnsParamsQueryType = "NB"
	AS112SummaryEdnsParamsQueryTypeNbstat     AS112SummaryEdnsParamsQueryType = "NBSTAT"
	AS112SummaryEdnsParamsQueryTypeNid        AS112SummaryEdnsParamsQueryType = "NID"
	AS112SummaryEdnsParamsQueryTypeNimloc     AS112SummaryEdnsParamsQueryType = "NIMLOC"
	AS112SummaryEdnsParamsQueryTypeNinfo      AS112SummaryEdnsParamsQueryType = "NINFO"
	AS112SummaryEdnsParamsQueryTypeNS         AS112SummaryEdnsParamsQueryType = "NS"
	AS112SummaryEdnsParamsQueryTypeNsap       AS112SummaryEdnsParamsQueryType = "NSAP"
	AS112SummaryEdnsParamsQueryTypeNsec       AS112SummaryEdnsParamsQueryType = "NSEC"
	AS112SummaryEdnsParamsQueryTypeNsec3      AS112SummaryEdnsParamsQueryType = "NSEC3"
	AS112SummaryEdnsParamsQueryTypeNsec3Param AS112SummaryEdnsParamsQueryType = "NSEC3PARAM"
	AS112SummaryEdnsParamsQueryTypeNull       AS112SummaryEdnsParamsQueryType = "NULL"
	AS112SummaryEdnsParamsQueryTypeNxt        AS112SummaryEdnsParamsQueryType = "NXT"
	AS112SummaryEdnsParamsQueryTypeOpenpgpkey AS112SummaryEdnsParamsQueryType = "OPENPGPKEY"
	AS112SummaryEdnsParamsQueryTypeOpt        AS112SummaryEdnsParamsQueryType = "OPT"
	AS112SummaryEdnsParamsQueryTypePTR        AS112SummaryEdnsParamsQueryType = "PTR"
	AS112SummaryEdnsParamsQueryTypePx         AS112SummaryEdnsParamsQueryType = "PX"
	AS112SummaryEdnsParamsQueryTypeRkey       AS112SummaryEdnsParamsQueryType = "RKEY"
	AS112SummaryEdnsParamsQueryTypeRp         AS112SummaryEdnsParamsQueryType = "RP"
	AS112SummaryEdnsParamsQueryTypeRrsig      AS112SummaryEdnsParamsQueryType = "RRSIG"
	AS112SummaryEdnsParamsQueryTypeRt         AS112SummaryEdnsParamsQueryType = "RT"
	AS112SummaryEdnsParamsQueryTypeSig        AS112SummaryEdnsParamsQueryType = "SIG"
	AS112SummaryEdnsParamsQueryTypeSink       AS112SummaryEdnsParamsQueryType = "SINK"
	AS112SummaryEdnsParamsQueryTypeSMIMEA     AS112SummaryEdnsParamsQueryType = "SMIMEA"
	AS112SummaryEdnsParamsQueryTypeSOA        AS112SummaryEdnsParamsQueryType = "SOA"
	AS112SummaryEdnsParamsQueryTypeSPF        AS112SummaryEdnsParamsQueryType = "SPF"
	AS112SummaryEdnsParamsQueryTypeSRV        AS112SummaryEdnsParamsQueryType = "SRV"
	AS112SummaryEdnsParamsQueryTypeSSHFP      AS112SummaryEdnsParamsQueryType = "SSHFP"
	AS112SummaryEdnsParamsQueryTypeSVCB       AS112SummaryEdnsParamsQueryType = "SVCB"
	AS112SummaryEdnsParamsQueryTypeTa         AS112SummaryEdnsParamsQueryType = "TA"
	AS112SummaryEdnsParamsQueryTypeTalink     AS112SummaryEdnsParamsQueryType = "TALINK"
	AS112SummaryEdnsParamsQueryTypeTkey       AS112SummaryEdnsParamsQueryType = "TKEY"
	AS112SummaryEdnsParamsQueryTypeTLSA       AS112SummaryEdnsParamsQueryType = "TLSA"
	AS112SummaryEdnsParamsQueryTypeTSIG       AS112SummaryEdnsParamsQueryType = "TSIG"
	AS112SummaryEdnsParamsQueryTypeTXT        AS112SummaryEdnsParamsQueryType = "TXT"
	AS112SummaryEdnsParamsQueryTypeUinfo      AS112SummaryEdnsParamsQueryType = "UINFO"
	AS112SummaryEdnsParamsQueryTypeUID        AS112SummaryEdnsParamsQueryType = "UID"
	AS112SummaryEdnsParamsQueryTypeUnspec     AS112SummaryEdnsParamsQueryType = "UNSPEC"
	AS112SummaryEdnsParamsQueryTypeURI        AS112SummaryEdnsParamsQueryType = "URI"
	AS112SummaryEdnsParamsQueryTypeWks        AS112SummaryEdnsParamsQueryType = "WKS"
	AS112SummaryEdnsParamsQueryTypeX25        AS112SummaryEdnsParamsQueryType = "X25"
	AS112SummaryEdnsParamsQueryTypeZonemd     AS112SummaryEdnsParamsQueryType = "ZONEMD"
)

func (r AS112SummaryEdnsParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsParamsQueryTypeA, AS112SummaryEdnsParamsQueryTypeAAAA, AS112SummaryEdnsParamsQueryTypeA6, AS112SummaryEdnsParamsQueryTypeAfsdb, AS112SummaryEdnsParamsQueryTypeAny, AS112SummaryEdnsParamsQueryTypeApl, AS112SummaryEdnsParamsQueryTypeAtma, AS112SummaryEdnsParamsQueryTypeAXFR, AS112SummaryEdnsParamsQueryTypeCAA, AS112SummaryEdnsParamsQueryTypeCdnskey, AS112SummaryEdnsParamsQueryTypeCds, AS112SummaryEdnsParamsQueryTypeCERT, AS112SummaryEdnsParamsQueryTypeCNAME, AS112SummaryEdnsParamsQueryTypeCsync, AS112SummaryEdnsParamsQueryTypeDhcid, AS112SummaryEdnsParamsQueryTypeDlv, AS112SummaryEdnsParamsQueryTypeDname, AS112SummaryEdnsParamsQueryTypeDNSKEY, AS112SummaryEdnsParamsQueryTypeDoa, AS112SummaryEdnsParamsQueryTypeDS, AS112SummaryEdnsParamsQueryTypeEid, AS112SummaryEdnsParamsQueryTypeEui48, AS112SummaryEdnsParamsQueryTypeEui64, AS112SummaryEdnsParamsQueryTypeGpos, AS112SummaryEdnsParamsQueryTypeGid, AS112SummaryEdnsParamsQueryTypeHinfo, AS112SummaryEdnsParamsQueryTypeHip, AS112SummaryEdnsParamsQueryTypeHTTPS, AS112SummaryEdnsParamsQueryTypeIpseckey, AS112SummaryEdnsParamsQueryTypeIsdn, AS112SummaryEdnsParamsQueryTypeIxfr, AS112SummaryEdnsParamsQueryTypeKey, AS112SummaryEdnsParamsQueryTypeKx, AS112SummaryEdnsParamsQueryTypeL32, AS112SummaryEdnsParamsQueryTypeL64, AS112SummaryEdnsParamsQueryTypeLOC, AS112SummaryEdnsParamsQueryTypeLp, AS112SummaryEdnsParamsQueryTypeMaila, AS112SummaryEdnsParamsQueryTypeMailb, AS112SummaryEdnsParamsQueryTypeMB, AS112SummaryEdnsParamsQueryTypeMd, AS112SummaryEdnsParamsQueryTypeMf, AS112SummaryEdnsParamsQueryTypeMg, AS112SummaryEdnsParamsQueryTypeMinfo, AS112SummaryEdnsParamsQueryTypeMr, AS112SummaryEdnsParamsQueryTypeMX, AS112SummaryEdnsParamsQueryTypeNAPTR, AS112SummaryEdnsParamsQueryTypeNb, AS112SummaryEdnsParamsQueryTypeNbstat, AS112SummaryEdnsParamsQueryTypeNid, AS112SummaryEdnsParamsQueryTypeNimloc, AS112SummaryEdnsParamsQueryTypeNinfo, AS112SummaryEdnsParamsQueryTypeNS, AS112SummaryEdnsParamsQueryTypeNsap, AS112SummaryEdnsParamsQueryTypeNsec, AS112SummaryEdnsParamsQueryTypeNsec3, AS112SummaryEdnsParamsQueryTypeNsec3Param, AS112SummaryEdnsParamsQueryTypeNull, AS112SummaryEdnsParamsQueryTypeNxt, AS112SummaryEdnsParamsQueryTypeOpenpgpkey, AS112SummaryEdnsParamsQueryTypeOpt, AS112SummaryEdnsParamsQueryTypePTR, AS112SummaryEdnsParamsQueryTypePx, AS112SummaryEdnsParamsQueryTypeRkey, AS112SummaryEdnsParamsQueryTypeRp, AS112SummaryEdnsParamsQueryTypeRrsig, AS112SummaryEdnsParamsQueryTypeRt, AS112SummaryEdnsParamsQueryTypeSig, AS112SummaryEdnsParamsQueryTypeSink, AS112SummaryEdnsParamsQueryTypeSMIMEA, AS112SummaryEdnsParamsQueryTypeSOA, AS112SummaryEdnsParamsQueryTypeSPF, AS112SummaryEdnsParamsQueryTypeSRV, AS112SummaryEdnsParamsQueryTypeSSHFP, AS112SummaryEdnsParamsQueryTypeSVCB, AS112SummaryEdnsParamsQueryTypeTa, AS112SummaryEdnsParamsQueryTypeTalink, AS112SummaryEdnsParamsQueryTypeTkey, AS112SummaryEdnsParamsQueryTypeTLSA, AS112SummaryEdnsParamsQueryTypeTSIG, AS112SummaryEdnsParamsQueryTypeTXT, AS112SummaryEdnsParamsQueryTypeUinfo, AS112SummaryEdnsParamsQueryTypeUID, AS112SummaryEdnsParamsQueryTypeUnspec, AS112SummaryEdnsParamsQueryTypeURI, AS112SummaryEdnsParamsQueryTypeWks, AS112SummaryEdnsParamsQueryTypeX25, AS112SummaryEdnsParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112SummaryEdnsParamsResponseCode string

const (
	AS112SummaryEdnsParamsResponseCodeNoerror   AS112SummaryEdnsParamsResponseCode = "NOERROR"
	AS112SummaryEdnsParamsResponseCodeFormerr   AS112SummaryEdnsParamsResponseCode = "FORMERR"
	AS112SummaryEdnsParamsResponseCodeServfail  AS112SummaryEdnsParamsResponseCode = "SERVFAIL"
	AS112SummaryEdnsParamsResponseCodeNxdomain  AS112SummaryEdnsParamsResponseCode = "NXDOMAIN"
	AS112SummaryEdnsParamsResponseCodeNotimp    AS112SummaryEdnsParamsResponseCode = "NOTIMP"
	AS112SummaryEdnsParamsResponseCodeRefused   AS112SummaryEdnsParamsResponseCode = "REFUSED"
	AS112SummaryEdnsParamsResponseCodeYxdomain  AS112SummaryEdnsParamsResponseCode = "YXDOMAIN"
	AS112SummaryEdnsParamsResponseCodeYxrrset   AS112SummaryEdnsParamsResponseCode = "YXRRSET"
	AS112SummaryEdnsParamsResponseCodeNxrrset   AS112SummaryEdnsParamsResponseCode = "NXRRSET"
	AS112SummaryEdnsParamsResponseCodeNotauth   AS112SummaryEdnsParamsResponseCode = "NOTAUTH"
	AS112SummaryEdnsParamsResponseCodeNotzone   AS112SummaryEdnsParamsResponseCode = "NOTZONE"
	AS112SummaryEdnsParamsResponseCodeBadsig    AS112SummaryEdnsParamsResponseCode = "BADSIG"
	AS112SummaryEdnsParamsResponseCodeBadkey    AS112SummaryEdnsParamsResponseCode = "BADKEY"
	AS112SummaryEdnsParamsResponseCodeBadtime   AS112SummaryEdnsParamsResponseCode = "BADTIME"
	AS112SummaryEdnsParamsResponseCodeBadmode   AS112SummaryEdnsParamsResponseCode = "BADMODE"
	AS112SummaryEdnsParamsResponseCodeBadname   AS112SummaryEdnsParamsResponseCode = "BADNAME"
	AS112SummaryEdnsParamsResponseCodeBadalg    AS112SummaryEdnsParamsResponseCode = "BADALG"
	AS112SummaryEdnsParamsResponseCodeBadtrunc  AS112SummaryEdnsParamsResponseCode = "BADTRUNC"
	AS112SummaryEdnsParamsResponseCodeBadcookie AS112SummaryEdnsParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryEdnsParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryEdnsParamsResponseCodeNoerror, AS112SummaryEdnsParamsResponseCodeFormerr, AS112SummaryEdnsParamsResponseCodeServfail, AS112SummaryEdnsParamsResponseCodeNxdomain, AS112SummaryEdnsParamsResponseCodeNotimp, AS112SummaryEdnsParamsResponseCodeRefused, AS112SummaryEdnsParamsResponseCodeYxdomain, AS112SummaryEdnsParamsResponseCodeYxrrset, AS112SummaryEdnsParamsResponseCodeNxrrset, AS112SummaryEdnsParamsResponseCodeNotauth, AS112SummaryEdnsParamsResponseCodeNotzone, AS112SummaryEdnsParamsResponseCodeBadsig, AS112SummaryEdnsParamsResponseCodeBadkey, AS112SummaryEdnsParamsResponseCodeBadtime, AS112SummaryEdnsParamsResponseCodeBadmode, AS112SummaryEdnsParamsResponseCodeBadname, AS112SummaryEdnsParamsResponseCodeBadalg, AS112SummaryEdnsParamsResponseCodeBadtrunc, AS112SummaryEdnsParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryEdnsResponseEnvelope struct {
	Result  AS112SummaryEdnsResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    as112SummaryEdnsResponseEnvelopeJSON `json:"-"`
}

// as112SummaryEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112SummaryEdnsResponseEnvelope]
type as112SummaryEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryIPVersionParams struct {
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
	Format param.Field[AS112SummaryIPVersionParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112SummaryIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112SummaryIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112SummaryIPVersionParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryIPVersionParamsFormat string

const (
	AS112SummaryIPVersionParamsFormatJson AS112SummaryIPVersionParamsFormat = "JSON"
	AS112SummaryIPVersionParamsFormatCsv  AS112SummaryIPVersionParamsFormat = "CSV"
)

func (r AS112SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionParamsFormatJson, AS112SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112SummaryIPVersionParamsProtocol string

const (
	AS112SummaryIPVersionParamsProtocolUdp   AS112SummaryIPVersionParamsProtocol = "UDP"
	AS112SummaryIPVersionParamsProtocolTCP   AS112SummaryIPVersionParamsProtocol = "TCP"
	AS112SummaryIPVersionParamsProtocolHTTPS AS112SummaryIPVersionParamsProtocol = "HTTPS"
	AS112SummaryIPVersionParamsProtocolTLS   AS112SummaryIPVersionParamsProtocol = "TLS"
)

func (r AS112SummaryIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionParamsProtocolUdp, AS112SummaryIPVersionParamsProtocolTCP, AS112SummaryIPVersionParamsProtocolHTTPS, AS112SummaryIPVersionParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112SummaryIPVersionParamsQueryType string

const (
	AS112SummaryIPVersionParamsQueryTypeA          AS112SummaryIPVersionParamsQueryType = "A"
	AS112SummaryIPVersionParamsQueryTypeAAAA       AS112SummaryIPVersionParamsQueryType = "AAAA"
	AS112SummaryIPVersionParamsQueryTypeA6         AS112SummaryIPVersionParamsQueryType = "A6"
	AS112SummaryIPVersionParamsQueryTypeAfsdb      AS112SummaryIPVersionParamsQueryType = "AFSDB"
	AS112SummaryIPVersionParamsQueryTypeAny        AS112SummaryIPVersionParamsQueryType = "ANY"
	AS112SummaryIPVersionParamsQueryTypeApl        AS112SummaryIPVersionParamsQueryType = "APL"
	AS112SummaryIPVersionParamsQueryTypeAtma       AS112SummaryIPVersionParamsQueryType = "ATMA"
	AS112SummaryIPVersionParamsQueryTypeAXFR       AS112SummaryIPVersionParamsQueryType = "AXFR"
	AS112SummaryIPVersionParamsQueryTypeCAA        AS112SummaryIPVersionParamsQueryType = "CAA"
	AS112SummaryIPVersionParamsQueryTypeCdnskey    AS112SummaryIPVersionParamsQueryType = "CDNSKEY"
	AS112SummaryIPVersionParamsQueryTypeCds        AS112SummaryIPVersionParamsQueryType = "CDS"
	AS112SummaryIPVersionParamsQueryTypeCERT       AS112SummaryIPVersionParamsQueryType = "CERT"
	AS112SummaryIPVersionParamsQueryTypeCNAME      AS112SummaryIPVersionParamsQueryType = "CNAME"
	AS112SummaryIPVersionParamsQueryTypeCsync      AS112SummaryIPVersionParamsQueryType = "CSYNC"
	AS112SummaryIPVersionParamsQueryTypeDhcid      AS112SummaryIPVersionParamsQueryType = "DHCID"
	AS112SummaryIPVersionParamsQueryTypeDlv        AS112SummaryIPVersionParamsQueryType = "DLV"
	AS112SummaryIPVersionParamsQueryTypeDname      AS112SummaryIPVersionParamsQueryType = "DNAME"
	AS112SummaryIPVersionParamsQueryTypeDNSKEY     AS112SummaryIPVersionParamsQueryType = "DNSKEY"
	AS112SummaryIPVersionParamsQueryTypeDoa        AS112SummaryIPVersionParamsQueryType = "DOA"
	AS112SummaryIPVersionParamsQueryTypeDS         AS112SummaryIPVersionParamsQueryType = "DS"
	AS112SummaryIPVersionParamsQueryTypeEid        AS112SummaryIPVersionParamsQueryType = "EID"
	AS112SummaryIPVersionParamsQueryTypeEui48      AS112SummaryIPVersionParamsQueryType = "EUI48"
	AS112SummaryIPVersionParamsQueryTypeEui64      AS112SummaryIPVersionParamsQueryType = "EUI64"
	AS112SummaryIPVersionParamsQueryTypeGpos       AS112SummaryIPVersionParamsQueryType = "GPOS"
	AS112SummaryIPVersionParamsQueryTypeGid        AS112SummaryIPVersionParamsQueryType = "GID"
	AS112SummaryIPVersionParamsQueryTypeHinfo      AS112SummaryIPVersionParamsQueryType = "HINFO"
	AS112SummaryIPVersionParamsQueryTypeHip        AS112SummaryIPVersionParamsQueryType = "HIP"
	AS112SummaryIPVersionParamsQueryTypeHTTPS      AS112SummaryIPVersionParamsQueryType = "HTTPS"
	AS112SummaryIPVersionParamsQueryTypeIpseckey   AS112SummaryIPVersionParamsQueryType = "IPSECKEY"
	AS112SummaryIPVersionParamsQueryTypeIsdn       AS112SummaryIPVersionParamsQueryType = "ISDN"
	AS112SummaryIPVersionParamsQueryTypeIxfr       AS112SummaryIPVersionParamsQueryType = "IXFR"
	AS112SummaryIPVersionParamsQueryTypeKey        AS112SummaryIPVersionParamsQueryType = "KEY"
	AS112SummaryIPVersionParamsQueryTypeKx         AS112SummaryIPVersionParamsQueryType = "KX"
	AS112SummaryIPVersionParamsQueryTypeL32        AS112SummaryIPVersionParamsQueryType = "L32"
	AS112SummaryIPVersionParamsQueryTypeL64        AS112SummaryIPVersionParamsQueryType = "L64"
	AS112SummaryIPVersionParamsQueryTypeLOC        AS112SummaryIPVersionParamsQueryType = "LOC"
	AS112SummaryIPVersionParamsQueryTypeLp         AS112SummaryIPVersionParamsQueryType = "LP"
	AS112SummaryIPVersionParamsQueryTypeMaila      AS112SummaryIPVersionParamsQueryType = "MAILA"
	AS112SummaryIPVersionParamsQueryTypeMailb      AS112SummaryIPVersionParamsQueryType = "MAILB"
	AS112SummaryIPVersionParamsQueryTypeMB         AS112SummaryIPVersionParamsQueryType = "MB"
	AS112SummaryIPVersionParamsQueryTypeMd         AS112SummaryIPVersionParamsQueryType = "MD"
	AS112SummaryIPVersionParamsQueryTypeMf         AS112SummaryIPVersionParamsQueryType = "MF"
	AS112SummaryIPVersionParamsQueryTypeMg         AS112SummaryIPVersionParamsQueryType = "MG"
	AS112SummaryIPVersionParamsQueryTypeMinfo      AS112SummaryIPVersionParamsQueryType = "MINFO"
	AS112SummaryIPVersionParamsQueryTypeMr         AS112SummaryIPVersionParamsQueryType = "MR"
	AS112SummaryIPVersionParamsQueryTypeMX         AS112SummaryIPVersionParamsQueryType = "MX"
	AS112SummaryIPVersionParamsQueryTypeNAPTR      AS112SummaryIPVersionParamsQueryType = "NAPTR"
	AS112SummaryIPVersionParamsQueryTypeNb         AS112SummaryIPVersionParamsQueryType = "NB"
	AS112SummaryIPVersionParamsQueryTypeNbstat     AS112SummaryIPVersionParamsQueryType = "NBSTAT"
	AS112SummaryIPVersionParamsQueryTypeNid        AS112SummaryIPVersionParamsQueryType = "NID"
	AS112SummaryIPVersionParamsQueryTypeNimloc     AS112SummaryIPVersionParamsQueryType = "NIMLOC"
	AS112SummaryIPVersionParamsQueryTypeNinfo      AS112SummaryIPVersionParamsQueryType = "NINFO"
	AS112SummaryIPVersionParamsQueryTypeNS         AS112SummaryIPVersionParamsQueryType = "NS"
	AS112SummaryIPVersionParamsQueryTypeNsap       AS112SummaryIPVersionParamsQueryType = "NSAP"
	AS112SummaryIPVersionParamsQueryTypeNsec       AS112SummaryIPVersionParamsQueryType = "NSEC"
	AS112SummaryIPVersionParamsQueryTypeNsec3      AS112SummaryIPVersionParamsQueryType = "NSEC3"
	AS112SummaryIPVersionParamsQueryTypeNsec3Param AS112SummaryIPVersionParamsQueryType = "NSEC3PARAM"
	AS112SummaryIPVersionParamsQueryTypeNull       AS112SummaryIPVersionParamsQueryType = "NULL"
	AS112SummaryIPVersionParamsQueryTypeNxt        AS112SummaryIPVersionParamsQueryType = "NXT"
	AS112SummaryIPVersionParamsQueryTypeOpenpgpkey AS112SummaryIPVersionParamsQueryType = "OPENPGPKEY"
	AS112SummaryIPVersionParamsQueryTypeOpt        AS112SummaryIPVersionParamsQueryType = "OPT"
	AS112SummaryIPVersionParamsQueryTypePTR        AS112SummaryIPVersionParamsQueryType = "PTR"
	AS112SummaryIPVersionParamsQueryTypePx         AS112SummaryIPVersionParamsQueryType = "PX"
	AS112SummaryIPVersionParamsQueryTypeRkey       AS112SummaryIPVersionParamsQueryType = "RKEY"
	AS112SummaryIPVersionParamsQueryTypeRp         AS112SummaryIPVersionParamsQueryType = "RP"
	AS112SummaryIPVersionParamsQueryTypeRrsig      AS112SummaryIPVersionParamsQueryType = "RRSIG"
	AS112SummaryIPVersionParamsQueryTypeRt         AS112SummaryIPVersionParamsQueryType = "RT"
	AS112SummaryIPVersionParamsQueryTypeSig        AS112SummaryIPVersionParamsQueryType = "SIG"
	AS112SummaryIPVersionParamsQueryTypeSink       AS112SummaryIPVersionParamsQueryType = "SINK"
	AS112SummaryIPVersionParamsQueryTypeSMIMEA     AS112SummaryIPVersionParamsQueryType = "SMIMEA"
	AS112SummaryIPVersionParamsQueryTypeSOA        AS112SummaryIPVersionParamsQueryType = "SOA"
	AS112SummaryIPVersionParamsQueryTypeSPF        AS112SummaryIPVersionParamsQueryType = "SPF"
	AS112SummaryIPVersionParamsQueryTypeSRV        AS112SummaryIPVersionParamsQueryType = "SRV"
	AS112SummaryIPVersionParamsQueryTypeSSHFP      AS112SummaryIPVersionParamsQueryType = "SSHFP"
	AS112SummaryIPVersionParamsQueryTypeSVCB       AS112SummaryIPVersionParamsQueryType = "SVCB"
	AS112SummaryIPVersionParamsQueryTypeTa         AS112SummaryIPVersionParamsQueryType = "TA"
	AS112SummaryIPVersionParamsQueryTypeTalink     AS112SummaryIPVersionParamsQueryType = "TALINK"
	AS112SummaryIPVersionParamsQueryTypeTkey       AS112SummaryIPVersionParamsQueryType = "TKEY"
	AS112SummaryIPVersionParamsQueryTypeTLSA       AS112SummaryIPVersionParamsQueryType = "TLSA"
	AS112SummaryIPVersionParamsQueryTypeTSIG       AS112SummaryIPVersionParamsQueryType = "TSIG"
	AS112SummaryIPVersionParamsQueryTypeTXT        AS112SummaryIPVersionParamsQueryType = "TXT"
	AS112SummaryIPVersionParamsQueryTypeUinfo      AS112SummaryIPVersionParamsQueryType = "UINFO"
	AS112SummaryIPVersionParamsQueryTypeUID        AS112SummaryIPVersionParamsQueryType = "UID"
	AS112SummaryIPVersionParamsQueryTypeUnspec     AS112SummaryIPVersionParamsQueryType = "UNSPEC"
	AS112SummaryIPVersionParamsQueryTypeURI        AS112SummaryIPVersionParamsQueryType = "URI"
	AS112SummaryIPVersionParamsQueryTypeWks        AS112SummaryIPVersionParamsQueryType = "WKS"
	AS112SummaryIPVersionParamsQueryTypeX25        AS112SummaryIPVersionParamsQueryType = "X25"
	AS112SummaryIPVersionParamsQueryTypeZonemd     AS112SummaryIPVersionParamsQueryType = "ZONEMD"
)

func (r AS112SummaryIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionParamsQueryTypeA, AS112SummaryIPVersionParamsQueryTypeAAAA, AS112SummaryIPVersionParamsQueryTypeA6, AS112SummaryIPVersionParamsQueryTypeAfsdb, AS112SummaryIPVersionParamsQueryTypeAny, AS112SummaryIPVersionParamsQueryTypeApl, AS112SummaryIPVersionParamsQueryTypeAtma, AS112SummaryIPVersionParamsQueryTypeAXFR, AS112SummaryIPVersionParamsQueryTypeCAA, AS112SummaryIPVersionParamsQueryTypeCdnskey, AS112SummaryIPVersionParamsQueryTypeCds, AS112SummaryIPVersionParamsQueryTypeCERT, AS112SummaryIPVersionParamsQueryTypeCNAME, AS112SummaryIPVersionParamsQueryTypeCsync, AS112SummaryIPVersionParamsQueryTypeDhcid, AS112SummaryIPVersionParamsQueryTypeDlv, AS112SummaryIPVersionParamsQueryTypeDname, AS112SummaryIPVersionParamsQueryTypeDNSKEY, AS112SummaryIPVersionParamsQueryTypeDoa, AS112SummaryIPVersionParamsQueryTypeDS, AS112SummaryIPVersionParamsQueryTypeEid, AS112SummaryIPVersionParamsQueryTypeEui48, AS112SummaryIPVersionParamsQueryTypeEui64, AS112SummaryIPVersionParamsQueryTypeGpos, AS112SummaryIPVersionParamsQueryTypeGid, AS112SummaryIPVersionParamsQueryTypeHinfo, AS112SummaryIPVersionParamsQueryTypeHip, AS112SummaryIPVersionParamsQueryTypeHTTPS, AS112SummaryIPVersionParamsQueryTypeIpseckey, AS112SummaryIPVersionParamsQueryTypeIsdn, AS112SummaryIPVersionParamsQueryTypeIxfr, AS112SummaryIPVersionParamsQueryTypeKey, AS112SummaryIPVersionParamsQueryTypeKx, AS112SummaryIPVersionParamsQueryTypeL32, AS112SummaryIPVersionParamsQueryTypeL64, AS112SummaryIPVersionParamsQueryTypeLOC, AS112SummaryIPVersionParamsQueryTypeLp, AS112SummaryIPVersionParamsQueryTypeMaila, AS112SummaryIPVersionParamsQueryTypeMailb, AS112SummaryIPVersionParamsQueryTypeMB, AS112SummaryIPVersionParamsQueryTypeMd, AS112SummaryIPVersionParamsQueryTypeMf, AS112SummaryIPVersionParamsQueryTypeMg, AS112SummaryIPVersionParamsQueryTypeMinfo, AS112SummaryIPVersionParamsQueryTypeMr, AS112SummaryIPVersionParamsQueryTypeMX, AS112SummaryIPVersionParamsQueryTypeNAPTR, AS112SummaryIPVersionParamsQueryTypeNb, AS112SummaryIPVersionParamsQueryTypeNbstat, AS112SummaryIPVersionParamsQueryTypeNid, AS112SummaryIPVersionParamsQueryTypeNimloc, AS112SummaryIPVersionParamsQueryTypeNinfo, AS112SummaryIPVersionParamsQueryTypeNS, AS112SummaryIPVersionParamsQueryTypeNsap, AS112SummaryIPVersionParamsQueryTypeNsec, AS112SummaryIPVersionParamsQueryTypeNsec3, AS112SummaryIPVersionParamsQueryTypeNsec3Param, AS112SummaryIPVersionParamsQueryTypeNull, AS112SummaryIPVersionParamsQueryTypeNxt, AS112SummaryIPVersionParamsQueryTypeOpenpgpkey, AS112SummaryIPVersionParamsQueryTypeOpt, AS112SummaryIPVersionParamsQueryTypePTR, AS112SummaryIPVersionParamsQueryTypePx, AS112SummaryIPVersionParamsQueryTypeRkey, AS112SummaryIPVersionParamsQueryTypeRp, AS112SummaryIPVersionParamsQueryTypeRrsig, AS112SummaryIPVersionParamsQueryTypeRt, AS112SummaryIPVersionParamsQueryTypeSig, AS112SummaryIPVersionParamsQueryTypeSink, AS112SummaryIPVersionParamsQueryTypeSMIMEA, AS112SummaryIPVersionParamsQueryTypeSOA, AS112SummaryIPVersionParamsQueryTypeSPF, AS112SummaryIPVersionParamsQueryTypeSRV, AS112SummaryIPVersionParamsQueryTypeSSHFP, AS112SummaryIPVersionParamsQueryTypeSVCB, AS112SummaryIPVersionParamsQueryTypeTa, AS112SummaryIPVersionParamsQueryTypeTalink, AS112SummaryIPVersionParamsQueryTypeTkey, AS112SummaryIPVersionParamsQueryTypeTLSA, AS112SummaryIPVersionParamsQueryTypeTSIG, AS112SummaryIPVersionParamsQueryTypeTXT, AS112SummaryIPVersionParamsQueryTypeUinfo, AS112SummaryIPVersionParamsQueryTypeUID, AS112SummaryIPVersionParamsQueryTypeUnspec, AS112SummaryIPVersionParamsQueryTypeURI, AS112SummaryIPVersionParamsQueryTypeWks, AS112SummaryIPVersionParamsQueryTypeX25, AS112SummaryIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112SummaryIPVersionParamsResponseCode string

const (
	AS112SummaryIPVersionParamsResponseCodeNoerror   AS112SummaryIPVersionParamsResponseCode = "NOERROR"
	AS112SummaryIPVersionParamsResponseCodeFormerr   AS112SummaryIPVersionParamsResponseCode = "FORMERR"
	AS112SummaryIPVersionParamsResponseCodeServfail  AS112SummaryIPVersionParamsResponseCode = "SERVFAIL"
	AS112SummaryIPVersionParamsResponseCodeNxdomain  AS112SummaryIPVersionParamsResponseCode = "NXDOMAIN"
	AS112SummaryIPVersionParamsResponseCodeNotimp    AS112SummaryIPVersionParamsResponseCode = "NOTIMP"
	AS112SummaryIPVersionParamsResponseCodeRefused   AS112SummaryIPVersionParamsResponseCode = "REFUSED"
	AS112SummaryIPVersionParamsResponseCodeYxdomain  AS112SummaryIPVersionParamsResponseCode = "YXDOMAIN"
	AS112SummaryIPVersionParamsResponseCodeYxrrset   AS112SummaryIPVersionParamsResponseCode = "YXRRSET"
	AS112SummaryIPVersionParamsResponseCodeNxrrset   AS112SummaryIPVersionParamsResponseCode = "NXRRSET"
	AS112SummaryIPVersionParamsResponseCodeNotauth   AS112SummaryIPVersionParamsResponseCode = "NOTAUTH"
	AS112SummaryIPVersionParamsResponseCodeNotzone   AS112SummaryIPVersionParamsResponseCode = "NOTZONE"
	AS112SummaryIPVersionParamsResponseCodeBadsig    AS112SummaryIPVersionParamsResponseCode = "BADSIG"
	AS112SummaryIPVersionParamsResponseCodeBadkey    AS112SummaryIPVersionParamsResponseCode = "BADKEY"
	AS112SummaryIPVersionParamsResponseCodeBadtime   AS112SummaryIPVersionParamsResponseCode = "BADTIME"
	AS112SummaryIPVersionParamsResponseCodeBadmode   AS112SummaryIPVersionParamsResponseCode = "BADMODE"
	AS112SummaryIPVersionParamsResponseCodeBadname   AS112SummaryIPVersionParamsResponseCode = "BADNAME"
	AS112SummaryIPVersionParamsResponseCodeBadalg    AS112SummaryIPVersionParamsResponseCode = "BADALG"
	AS112SummaryIPVersionParamsResponseCodeBadtrunc  AS112SummaryIPVersionParamsResponseCode = "BADTRUNC"
	AS112SummaryIPVersionParamsResponseCodeBadcookie AS112SummaryIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryIPVersionParamsResponseCodeNoerror, AS112SummaryIPVersionParamsResponseCodeFormerr, AS112SummaryIPVersionParamsResponseCodeServfail, AS112SummaryIPVersionParamsResponseCodeNxdomain, AS112SummaryIPVersionParamsResponseCodeNotimp, AS112SummaryIPVersionParamsResponseCodeRefused, AS112SummaryIPVersionParamsResponseCodeYxdomain, AS112SummaryIPVersionParamsResponseCodeYxrrset, AS112SummaryIPVersionParamsResponseCodeNxrrset, AS112SummaryIPVersionParamsResponseCodeNotauth, AS112SummaryIPVersionParamsResponseCodeNotzone, AS112SummaryIPVersionParamsResponseCodeBadsig, AS112SummaryIPVersionParamsResponseCodeBadkey, AS112SummaryIPVersionParamsResponseCodeBadtime, AS112SummaryIPVersionParamsResponseCodeBadmode, AS112SummaryIPVersionParamsResponseCodeBadname, AS112SummaryIPVersionParamsResponseCodeBadalg, AS112SummaryIPVersionParamsResponseCodeBadtrunc, AS112SummaryIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryIPVersionResponseEnvelope struct {
	Result  AS112SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    as112SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryIPVersionResponseEnvelope]
type as112SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryProtocolParams struct {
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
	Format param.Field[AS112SummaryProtocolParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112SummaryProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112SummaryProtocolParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryProtocolParamsFormat string

const (
	AS112SummaryProtocolParamsFormatJson AS112SummaryProtocolParamsFormat = "JSON"
	AS112SummaryProtocolParamsFormatCsv  AS112SummaryProtocolParamsFormat = "CSV"
)

func (r AS112SummaryProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryProtocolParamsFormatJson, AS112SummaryProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112SummaryProtocolParamsQueryType string

const (
	AS112SummaryProtocolParamsQueryTypeA          AS112SummaryProtocolParamsQueryType = "A"
	AS112SummaryProtocolParamsQueryTypeAAAA       AS112SummaryProtocolParamsQueryType = "AAAA"
	AS112SummaryProtocolParamsQueryTypeA6         AS112SummaryProtocolParamsQueryType = "A6"
	AS112SummaryProtocolParamsQueryTypeAfsdb      AS112SummaryProtocolParamsQueryType = "AFSDB"
	AS112SummaryProtocolParamsQueryTypeAny        AS112SummaryProtocolParamsQueryType = "ANY"
	AS112SummaryProtocolParamsQueryTypeApl        AS112SummaryProtocolParamsQueryType = "APL"
	AS112SummaryProtocolParamsQueryTypeAtma       AS112SummaryProtocolParamsQueryType = "ATMA"
	AS112SummaryProtocolParamsQueryTypeAXFR       AS112SummaryProtocolParamsQueryType = "AXFR"
	AS112SummaryProtocolParamsQueryTypeCAA        AS112SummaryProtocolParamsQueryType = "CAA"
	AS112SummaryProtocolParamsQueryTypeCdnskey    AS112SummaryProtocolParamsQueryType = "CDNSKEY"
	AS112SummaryProtocolParamsQueryTypeCds        AS112SummaryProtocolParamsQueryType = "CDS"
	AS112SummaryProtocolParamsQueryTypeCERT       AS112SummaryProtocolParamsQueryType = "CERT"
	AS112SummaryProtocolParamsQueryTypeCNAME      AS112SummaryProtocolParamsQueryType = "CNAME"
	AS112SummaryProtocolParamsQueryTypeCsync      AS112SummaryProtocolParamsQueryType = "CSYNC"
	AS112SummaryProtocolParamsQueryTypeDhcid      AS112SummaryProtocolParamsQueryType = "DHCID"
	AS112SummaryProtocolParamsQueryTypeDlv        AS112SummaryProtocolParamsQueryType = "DLV"
	AS112SummaryProtocolParamsQueryTypeDname      AS112SummaryProtocolParamsQueryType = "DNAME"
	AS112SummaryProtocolParamsQueryTypeDNSKEY     AS112SummaryProtocolParamsQueryType = "DNSKEY"
	AS112SummaryProtocolParamsQueryTypeDoa        AS112SummaryProtocolParamsQueryType = "DOA"
	AS112SummaryProtocolParamsQueryTypeDS         AS112SummaryProtocolParamsQueryType = "DS"
	AS112SummaryProtocolParamsQueryTypeEid        AS112SummaryProtocolParamsQueryType = "EID"
	AS112SummaryProtocolParamsQueryTypeEui48      AS112SummaryProtocolParamsQueryType = "EUI48"
	AS112SummaryProtocolParamsQueryTypeEui64      AS112SummaryProtocolParamsQueryType = "EUI64"
	AS112SummaryProtocolParamsQueryTypeGpos       AS112SummaryProtocolParamsQueryType = "GPOS"
	AS112SummaryProtocolParamsQueryTypeGid        AS112SummaryProtocolParamsQueryType = "GID"
	AS112SummaryProtocolParamsQueryTypeHinfo      AS112SummaryProtocolParamsQueryType = "HINFO"
	AS112SummaryProtocolParamsQueryTypeHip        AS112SummaryProtocolParamsQueryType = "HIP"
	AS112SummaryProtocolParamsQueryTypeHTTPS      AS112SummaryProtocolParamsQueryType = "HTTPS"
	AS112SummaryProtocolParamsQueryTypeIpseckey   AS112SummaryProtocolParamsQueryType = "IPSECKEY"
	AS112SummaryProtocolParamsQueryTypeIsdn       AS112SummaryProtocolParamsQueryType = "ISDN"
	AS112SummaryProtocolParamsQueryTypeIxfr       AS112SummaryProtocolParamsQueryType = "IXFR"
	AS112SummaryProtocolParamsQueryTypeKey        AS112SummaryProtocolParamsQueryType = "KEY"
	AS112SummaryProtocolParamsQueryTypeKx         AS112SummaryProtocolParamsQueryType = "KX"
	AS112SummaryProtocolParamsQueryTypeL32        AS112SummaryProtocolParamsQueryType = "L32"
	AS112SummaryProtocolParamsQueryTypeL64        AS112SummaryProtocolParamsQueryType = "L64"
	AS112SummaryProtocolParamsQueryTypeLOC        AS112SummaryProtocolParamsQueryType = "LOC"
	AS112SummaryProtocolParamsQueryTypeLp         AS112SummaryProtocolParamsQueryType = "LP"
	AS112SummaryProtocolParamsQueryTypeMaila      AS112SummaryProtocolParamsQueryType = "MAILA"
	AS112SummaryProtocolParamsQueryTypeMailb      AS112SummaryProtocolParamsQueryType = "MAILB"
	AS112SummaryProtocolParamsQueryTypeMB         AS112SummaryProtocolParamsQueryType = "MB"
	AS112SummaryProtocolParamsQueryTypeMd         AS112SummaryProtocolParamsQueryType = "MD"
	AS112SummaryProtocolParamsQueryTypeMf         AS112SummaryProtocolParamsQueryType = "MF"
	AS112SummaryProtocolParamsQueryTypeMg         AS112SummaryProtocolParamsQueryType = "MG"
	AS112SummaryProtocolParamsQueryTypeMinfo      AS112SummaryProtocolParamsQueryType = "MINFO"
	AS112SummaryProtocolParamsQueryTypeMr         AS112SummaryProtocolParamsQueryType = "MR"
	AS112SummaryProtocolParamsQueryTypeMX         AS112SummaryProtocolParamsQueryType = "MX"
	AS112SummaryProtocolParamsQueryTypeNAPTR      AS112SummaryProtocolParamsQueryType = "NAPTR"
	AS112SummaryProtocolParamsQueryTypeNb         AS112SummaryProtocolParamsQueryType = "NB"
	AS112SummaryProtocolParamsQueryTypeNbstat     AS112SummaryProtocolParamsQueryType = "NBSTAT"
	AS112SummaryProtocolParamsQueryTypeNid        AS112SummaryProtocolParamsQueryType = "NID"
	AS112SummaryProtocolParamsQueryTypeNimloc     AS112SummaryProtocolParamsQueryType = "NIMLOC"
	AS112SummaryProtocolParamsQueryTypeNinfo      AS112SummaryProtocolParamsQueryType = "NINFO"
	AS112SummaryProtocolParamsQueryTypeNS         AS112SummaryProtocolParamsQueryType = "NS"
	AS112SummaryProtocolParamsQueryTypeNsap       AS112SummaryProtocolParamsQueryType = "NSAP"
	AS112SummaryProtocolParamsQueryTypeNsec       AS112SummaryProtocolParamsQueryType = "NSEC"
	AS112SummaryProtocolParamsQueryTypeNsec3      AS112SummaryProtocolParamsQueryType = "NSEC3"
	AS112SummaryProtocolParamsQueryTypeNsec3Param AS112SummaryProtocolParamsQueryType = "NSEC3PARAM"
	AS112SummaryProtocolParamsQueryTypeNull       AS112SummaryProtocolParamsQueryType = "NULL"
	AS112SummaryProtocolParamsQueryTypeNxt        AS112SummaryProtocolParamsQueryType = "NXT"
	AS112SummaryProtocolParamsQueryTypeOpenpgpkey AS112SummaryProtocolParamsQueryType = "OPENPGPKEY"
	AS112SummaryProtocolParamsQueryTypeOpt        AS112SummaryProtocolParamsQueryType = "OPT"
	AS112SummaryProtocolParamsQueryTypePTR        AS112SummaryProtocolParamsQueryType = "PTR"
	AS112SummaryProtocolParamsQueryTypePx         AS112SummaryProtocolParamsQueryType = "PX"
	AS112SummaryProtocolParamsQueryTypeRkey       AS112SummaryProtocolParamsQueryType = "RKEY"
	AS112SummaryProtocolParamsQueryTypeRp         AS112SummaryProtocolParamsQueryType = "RP"
	AS112SummaryProtocolParamsQueryTypeRrsig      AS112SummaryProtocolParamsQueryType = "RRSIG"
	AS112SummaryProtocolParamsQueryTypeRt         AS112SummaryProtocolParamsQueryType = "RT"
	AS112SummaryProtocolParamsQueryTypeSig        AS112SummaryProtocolParamsQueryType = "SIG"
	AS112SummaryProtocolParamsQueryTypeSink       AS112SummaryProtocolParamsQueryType = "SINK"
	AS112SummaryProtocolParamsQueryTypeSMIMEA     AS112SummaryProtocolParamsQueryType = "SMIMEA"
	AS112SummaryProtocolParamsQueryTypeSOA        AS112SummaryProtocolParamsQueryType = "SOA"
	AS112SummaryProtocolParamsQueryTypeSPF        AS112SummaryProtocolParamsQueryType = "SPF"
	AS112SummaryProtocolParamsQueryTypeSRV        AS112SummaryProtocolParamsQueryType = "SRV"
	AS112SummaryProtocolParamsQueryTypeSSHFP      AS112SummaryProtocolParamsQueryType = "SSHFP"
	AS112SummaryProtocolParamsQueryTypeSVCB       AS112SummaryProtocolParamsQueryType = "SVCB"
	AS112SummaryProtocolParamsQueryTypeTa         AS112SummaryProtocolParamsQueryType = "TA"
	AS112SummaryProtocolParamsQueryTypeTalink     AS112SummaryProtocolParamsQueryType = "TALINK"
	AS112SummaryProtocolParamsQueryTypeTkey       AS112SummaryProtocolParamsQueryType = "TKEY"
	AS112SummaryProtocolParamsQueryTypeTLSA       AS112SummaryProtocolParamsQueryType = "TLSA"
	AS112SummaryProtocolParamsQueryTypeTSIG       AS112SummaryProtocolParamsQueryType = "TSIG"
	AS112SummaryProtocolParamsQueryTypeTXT        AS112SummaryProtocolParamsQueryType = "TXT"
	AS112SummaryProtocolParamsQueryTypeUinfo      AS112SummaryProtocolParamsQueryType = "UINFO"
	AS112SummaryProtocolParamsQueryTypeUID        AS112SummaryProtocolParamsQueryType = "UID"
	AS112SummaryProtocolParamsQueryTypeUnspec     AS112SummaryProtocolParamsQueryType = "UNSPEC"
	AS112SummaryProtocolParamsQueryTypeURI        AS112SummaryProtocolParamsQueryType = "URI"
	AS112SummaryProtocolParamsQueryTypeWks        AS112SummaryProtocolParamsQueryType = "WKS"
	AS112SummaryProtocolParamsQueryTypeX25        AS112SummaryProtocolParamsQueryType = "X25"
	AS112SummaryProtocolParamsQueryTypeZonemd     AS112SummaryProtocolParamsQueryType = "ZONEMD"
)

func (r AS112SummaryProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryProtocolParamsQueryTypeA, AS112SummaryProtocolParamsQueryTypeAAAA, AS112SummaryProtocolParamsQueryTypeA6, AS112SummaryProtocolParamsQueryTypeAfsdb, AS112SummaryProtocolParamsQueryTypeAny, AS112SummaryProtocolParamsQueryTypeApl, AS112SummaryProtocolParamsQueryTypeAtma, AS112SummaryProtocolParamsQueryTypeAXFR, AS112SummaryProtocolParamsQueryTypeCAA, AS112SummaryProtocolParamsQueryTypeCdnskey, AS112SummaryProtocolParamsQueryTypeCds, AS112SummaryProtocolParamsQueryTypeCERT, AS112SummaryProtocolParamsQueryTypeCNAME, AS112SummaryProtocolParamsQueryTypeCsync, AS112SummaryProtocolParamsQueryTypeDhcid, AS112SummaryProtocolParamsQueryTypeDlv, AS112SummaryProtocolParamsQueryTypeDname, AS112SummaryProtocolParamsQueryTypeDNSKEY, AS112SummaryProtocolParamsQueryTypeDoa, AS112SummaryProtocolParamsQueryTypeDS, AS112SummaryProtocolParamsQueryTypeEid, AS112SummaryProtocolParamsQueryTypeEui48, AS112SummaryProtocolParamsQueryTypeEui64, AS112SummaryProtocolParamsQueryTypeGpos, AS112SummaryProtocolParamsQueryTypeGid, AS112SummaryProtocolParamsQueryTypeHinfo, AS112SummaryProtocolParamsQueryTypeHip, AS112SummaryProtocolParamsQueryTypeHTTPS, AS112SummaryProtocolParamsQueryTypeIpseckey, AS112SummaryProtocolParamsQueryTypeIsdn, AS112SummaryProtocolParamsQueryTypeIxfr, AS112SummaryProtocolParamsQueryTypeKey, AS112SummaryProtocolParamsQueryTypeKx, AS112SummaryProtocolParamsQueryTypeL32, AS112SummaryProtocolParamsQueryTypeL64, AS112SummaryProtocolParamsQueryTypeLOC, AS112SummaryProtocolParamsQueryTypeLp, AS112SummaryProtocolParamsQueryTypeMaila, AS112SummaryProtocolParamsQueryTypeMailb, AS112SummaryProtocolParamsQueryTypeMB, AS112SummaryProtocolParamsQueryTypeMd, AS112SummaryProtocolParamsQueryTypeMf, AS112SummaryProtocolParamsQueryTypeMg, AS112SummaryProtocolParamsQueryTypeMinfo, AS112SummaryProtocolParamsQueryTypeMr, AS112SummaryProtocolParamsQueryTypeMX, AS112SummaryProtocolParamsQueryTypeNAPTR, AS112SummaryProtocolParamsQueryTypeNb, AS112SummaryProtocolParamsQueryTypeNbstat, AS112SummaryProtocolParamsQueryTypeNid, AS112SummaryProtocolParamsQueryTypeNimloc, AS112SummaryProtocolParamsQueryTypeNinfo, AS112SummaryProtocolParamsQueryTypeNS, AS112SummaryProtocolParamsQueryTypeNsap, AS112SummaryProtocolParamsQueryTypeNsec, AS112SummaryProtocolParamsQueryTypeNsec3, AS112SummaryProtocolParamsQueryTypeNsec3Param, AS112SummaryProtocolParamsQueryTypeNull, AS112SummaryProtocolParamsQueryTypeNxt, AS112SummaryProtocolParamsQueryTypeOpenpgpkey, AS112SummaryProtocolParamsQueryTypeOpt, AS112SummaryProtocolParamsQueryTypePTR, AS112SummaryProtocolParamsQueryTypePx, AS112SummaryProtocolParamsQueryTypeRkey, AS112SummaryProtocolParamsQueryTypeRp, AS112SummaryProtocolParamsQueryTypeRrsig, AS112SummaryProtocolParamsQueryTypeRt, AS112SummaryProtocolParamsQueryTypeSig, AS112SummaryProtocolParamsQueryTypeSink, AS112SummaryProtocolParamsQueryTypeSMIMEA, AS112SummaryProtocolParamsQueryTypeSOA, AS112SummaryProtocolParamsQueryTypeSPF, AS112SummaryProtocolParamsQueryTypeSRV, AS112SummaryProtocolParamsQueryTypeSSHFP, AS112SummaryProtocolParamsQueryTypeSVCB, AS112SummaryProtocolParamsQueryTypeTa, AS112SummaryProtocolParamsQueryTypeTalink, AS112SummaryProtocolParamsQueryTypeTkey, AS112SummaryProtocolParamsQueryTypeTLSA, AS112SummaryProtocolParamsQueryTypeTSIG, AS112SummaryProtocolParamsQueryTypeTXT, AS112SummaryProtocolParamsQueryTypeUinfo, AS112SummaryProtocolParamsQueryTypeUID, AS112SummaryProtocolParamsQueryTypeUnspec, AS112SummaryProtocolParamsQueryTypeURI, AS112SummaryProtocolParamsQueryTypeWks, AS112SummaryProtocolParamsQueryTypeX25, AS112SummaryProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112SummaryProtocolParamsResponseCode string

const (
	AS112SummaryProtocolParamsResponseCodeNoerror   AS112SummaryProtocolParamsResponseCode = "NOERROR"
	AS112SummaryProtocolParamsResponseCodeFormerr   AS112SummaryProtocolParamsResponseCode = "FORMERR"
	AS112SummaryProtocolParamsResponseCodeServfail  AS112SummaryProtocolParamsResponseCode = "SERVFAIL"
	AS112SummaryProtocolParamsResponseCodeNxdomain  AS112SummaryProtocolParamsResponseCode = "NXDOMAIN"
	AS112SummaryProtocolParamsResponseCodeNotimp    AS112SummaryProtocolParamsResponseCode = "NOTIMP"
	AS112SummaryProtocolParamsResponseCodeRefused   AS112SummaryProtocolParamsResponseCode = "REFUSED"
	AS112SummaryProtocolParamsResponseCodeYxdomain  AS112SummaryProtocolParamsResponseCode = "YXDOMAIN"
	AS112SummaryProtocolParamsResponseCodeYxrrset   AS112SummaryProtocolParamsResponseCode = "YXRRSET"
	AS112SummaryProtocolParamsResponseCodeNxrrset   AS112SummaryProtocolParamsResponseCode = "NXRRSET"
	AS112SummaryProtocolParamsResponseCodeNotauth   AS112SummaryProtocolParamsResponseCode = "NOTAUTH"
	AS112SummaryProtocolParamsResponseCodeNotzone   AS112SummaryProtocolParamsResponseCode = "NOTZONE"
	AS112SummaryProtocolParamsResponseCodeBadsig    AS112SummaryProtocolParamsResponseCode = "BADSIG"
	AS112SummaryProtocolParamsResponseCodeBadkey    AS112SummaryProtocolParamsResponseCode = "BADKEY"
	AS112SummaryProtocolParamsResponseCodeBadtime   AS112SummaryProtocolParamsResponseCode = "BADTIME"
	AS112SummaryProtocolParamsResponseCodeBadmode   AS112SummaryProtocolParamsResponseCode = "BADMODE"
	AS112SummaryProtocolParamsResponseCodeBadname   AS112SummaryProtocolParamsResponseCode = "BADNAME"
	AS112SummaryProtocolParamsResponseCodeBadalg    AS112SummaryProtocolParamsResponseCode = "BADALG"
	AS112SummaryProtocolParamsResponseCodeBadtrunc  AS112SummaryProtocolParamsResponseCode = "BADTRUNC"
	AS112SummaryProtocolParamsResponseCodeBadcookie AS112SummaryProtocolParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryProtocolParamsResponseCodeNoerror, AS112SummaryProtocolParamsResponseCodeFormerr, AS112SummaryProtocolParamsResponseCodeServfail, AS112SummaryProtocolParamsResponseCodeNxdomain, AS112SummaryProtocolParamsResponseCodeNotimp, AS112SummaryProtocolParamsResponseCodeRefused, AS112SummaryProtocolParamsResponseCodeYxdomain, AS112SummaryProtocolParamsResponseCodeYxrrset, AS112SummaryProtocolParamsResponseCodeNxrrset, AS112SummaryProtocolParamsResponseCodeNotauth, AS112SummaryProtocolParamsResponseCodeNotzone, AS112SummaryProtocolParamsResponseCodeBadsig, AS112SummaryProtocolParamsResponseCodeBadkey, AS112SummaryProtocolParamsResponseCodeBadtime, AS112SummaryProtocolParamsResponseCodeBadmode, AS112SummaryProtocolParamsResponseCodeBadname, AS112SummaryProtocolParamsResponseCodeBadalg, AS112SummaryProtocolParamsResponseCodeBadtrunc, AS112SummaryProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryProtocolResponseEnvelope struct {
	Result  AS112SummaryProtocolResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    as112SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// as112SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryProtocolResponseEnvelope]
type as112SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryQueryTypeParams struct {
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
	Format param.Field[AS112SummaryQueryTypeParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112SummaryQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[AS112SummaryQueryTypeParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryQueryTypeParamsFormat string

const (
	AS112SummaryQueryTypeParamsFormatJson AS112SummaryQueryTypeParamsFormat = "JSON"
	AS112SummaryQueryTypeParamsFormatCsv  AS112SummaryQueryTypeParamsFormat = "CSV"
)

func (r AS112SummaryQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryQueryTypeParamsFormatJson, AS112SummaryQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112SummaryQueryTypeParamsProtocol string

const (
	AS112SummaryQueryTypeParamsProtocolUdp   AS112SummaryQueryTypeParamsProtocol = "UDP"
	AS112SummaryQueryTypeParamsProtocolTCP   AS112SummaryQueryTypeParamsProtocol = "TCP"
	AS112SummaryQueryTypeParamsProtocolHTTPS AS112SummaryQueryTypeParamsProtocol = "HTTPS"
	AS112SummaryQueryTypeParamsProtocolTLS   AS112SummaryQueryTypeParamsProtocol = "TLS"
)

func (r AS112SummaryQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryQueryTypeParamsProtocolUdp, AS112SummaryQueryTypeParamsProtocolTCP, AS112SummaryQueryTypeParamsProtocolHTTPS, AS112SummaryQueryTypeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS response code.
type AS112SummaryQueryTypeParamsResponseCode string

const (
	AS112SummaryQueryTypeParamsResponseCodeNoerror   AS112SummaryQueryTypeParamsResponseCode = "NOERROR"
	AS112SummaryQueryTypeParamsResponseCodeFormerr   AS112SummaryQueryTypeParamsResponseCode = "FORMERR"
	AS112SummaryQueryTypeParamsResponseCodeServfail  AS112SummaryQueryTypeParamsResponseCode = "SERVFAIL"
	AS112SummaryQueryTypeParamsResponseCodeNxdomain  AS112SummaryQueryTypeParamsResponseCode = "NXDOMAIN"
	AS112SummaryQueryTypeParamsResponseCodeNotimp    AS112SummaryQueryTypeParamsResponseCode = "NOTIMP"
	AS112SummaryQueryTypeParamsResponseCodeRefused   AS112SummaryQueryTypeParamsResponseCode = "REFUSED"
	AS112SummaryQueryTypeParamsResponseCodeYxdomain  AS112SummaryQueryTypeParamsResponseCode = "YXDOMAIN"
	AS112SummaryQueryTypeParamsResponseCodeYxrrset   AS112SummaryQueryTypeParamsResponseCode = "YXRRSET"
	AS112SummaryQueryTypeParamsResponseCodeNxrrset   AS112SummaryQueryTypeParamsResponseCode = "NXRRSET"
	AS112SummaryQueryTypeParamsResponseCodeNotauth   AS112SummaryQueryTypeParamsResponseCode = "NOTAUTH"
	AS112SummaryQueryTypeParamsResponseCodeNotzone   AS112SummaryQueryTypeParamsResponseCode = "NOTZONE"
	AS112SummaryQueryTypeParamsResponseCodeBadsig    AS112SummaryQueryTypeParamsResponseCode = "BADSIG"
	AS112SummaryQueryTypeParamsResponseCodeBadkey    AS112SummaryQueryTypeParamsResponseCode = "BADKEY"
	AS112SummaryQueryTypeParamsResponseCodeBadtime   AS112SummaryQueryTypeParamsResponseCode = "BADTIME"
	AS112SummaryQueryTypeParamsResponseCodeBadmode   AS112SummaryQueryTypeParamsResponseCode = "BADMODE"
	AS112SummaryQueryTypeParamsResponseCodeBadname   AS112SummaryQueryTypeParamsResponseCode = "BADNAME"
	AS112SummaryQueryTypeParamsResponseCodeBadalg    AS112SummaryQueryTypeParamsResponseCode = "BADALG"
	AS112SummaryQueryTypeParamsResponseCodeBadtrunc  AS112SummaryQueryTypeParamsResponseCode = "BADTRUNC"
	AS112SummaryQueryTypeParamsResponseCodeBadcookie AS112SummaryQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryQueryTypeParamsResponseCodeNoerror, AS112SummaryQueryTypeParamsResponseCodeFormerr, AS112SummaryQueryTypeParamsResponseCodeServfail, AS112SummaryQueryTypeParamsResponseCodeNxdomain, AS112SummaryQueryTypeParamsResponseCodeNotimp, AS112SummaryQueryTypeParamsResponseCodeRefused, AS112SummaryQueryTypeParamsResponseCodeYxdomain, AS112SummaryQueryTypeParamsResponseCodeYxrrset, AS112SummaryQueryTypeParamsResponseCodeNxrrset, AS112SummaryQueryTypeParamsResponseCodeNotauth, AS112SummaryQueryTypeParamsResponseCodeNotzone, AS112SummaryQueryTypeParamsResponseCodeBadsig, AS112SummaryQueryTypeParamsResponseCodeBadkey, AS112SummaryQueryTypeParamsResponseCodeBadtime, AS112SummaryQueryTypeParamsResponseCodeBadmode, AS112SummaryQueryTypeParamsResponseCodeBadname, AS112SummaryQueryTypeParamsResponseCodeBadalg, AS112SummaryQueryTypeParamsResponseCodeBadtrunc, AS112SummaryQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryQueryTypeResponseEnvelope struct {
	Result  AS112SummaryQueryTypeResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    as112SummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// as112SummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryQueryTypeResponseEnvelope]
type as112SummaryQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryResponseCodesParams struct {
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
	Format param.Field[AS112SummaryResponseCodesParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[AS112SummaryResponseCodesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[AS112SummaryResponseCodesParamsQueryType] `query:"queryType"`
}

// URLQuery serializes [AS112SummaryResponseCodesParams]'s query parameters as
// `url.Values`.
func (r AS112SummaryResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112SummaryResponseCodesParamsFormat string

const (
	AS112SummaryResponseCodesParamsFormatJson AS112SummaryResponseCodesParamsFormat = "JSON"
	AS112SummaryResponseCodesParamsFormatCsv  AS112SummaryResponseCodesParamsFormat = "CSV"
)

func (r AS112SummaryResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryResponseCodesParamsFormatJson, AS112SummaryResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type AS112SummaryResponseCodesParamsProtocol string

const (
	AS112SummaryResponseCodesParamsProtocolUdp   AS112SummaryResponseCodesParamsProtocol = "UDP"
	AS112SummaryResponseCodesParamsProtocolTCP   AS112SummaryResponseCodesParamsProtocol = "TCP"
	AS112SummaryResponseCodesParamsProtocolHTTPS AS112SummaryResponseCodesParamsProtocol = "HTTPS"
	AS112SummaryResponseCodesParamsProtocolTLS   AS112SummaryResponseCodesParamsProtocol = "TLS"
)

func (r AS112SummaryResponseCodesParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryResponseCodesParamsProtocolUdp, AS112SummaryResponseCodesParamsProtocolTCP, AS112SummaryResponseCodesParamsProtocolHTTPS, AS112SummaryResponseCodesParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type AS112SummaryResponseCodesParamsQueryType string

const (
	AS112SummaryResponseCodesParamsQueryTypeA          AS112SummaryResponseCodesParamsQueryType = "A"
	AS112SummaryResponseCodesParamsQueryTypeAAAA       AS112SummaryResponseCodesParamsQueryType = "AAAA"
	AS112SummaryResponseCodesParamsQueryTypeA6         AS112SummaryResponseCodesParamsQueryType = "A6"
	AS112SummaryResponseCodesParamsQueryTypeAfsdb      AS112SummaryResponseCodesParamsQueryType = "AFSDB"
	AS112SummaryResponseCodesParamsQueryTypeAny        AS112SummaryResponseCodesParamsQueryType = "ANY"
	AS112SummaryResponseCodesParamsQueryTypeApl        AS112SummaryResponseCodesParamsQueryType = "APL"
	AS112SummaryResponseCodesParamsQueryTypeAtma       AS112SummaryResponseCodesParamsQueryType = "ATMA"
	AS112SummaryResponseCodesParamsQueryTypeAXFR       AS112SummaryResponseCodesParamsQueryType = "AXFR"
	AS112SummaryResponseCodesParamsQueryTypeCAA        AS112SummaryResponseCodesParamsQueryType = "CAA"
	AS112SummaryResponseCodesParamsQueryTypeCdnskey    AS112SummaryResponseCodesParamsQueryType = "CDNSKEY"
	AS112SummaryResponseCodesParamsQueryTypeCds        AS112SummaryResponseCodesParamsQueryType = "CDS"
	AS112SummaryResponseCodesParamsQueryTypeCERT       AS112SummaryResponseCodesParamsQueryType = "CERT"
	AS112SummaryResponseCodesParamsQueryTypeCNAME      AS112SummaryResponseCodesParamsQueryType = "CNAME"
	AS112SummaryResponseCodesParamsQueryTypeCsync      AS112SummaryResponseCodesParamsQueryType = "CSYNC"
	AS112SummaryResponseCodesParamsQueryTypeDhcid      AS112SummaryResponseCodesParamsQueryType = "DHCID"
	AS112SummaryResponseCodesParamsQueryTypeDlv        AS112SummaryResponseCodesParamsQueryType = "DLV"
	AS112SummaryResponseCodesParamsQueryTypeDname      AS112SummaryResponseCodesParamsQueryType = "DNAME"
	AS112SummaryResponseCodesParamsQueryTypeDNSKEY     AS112SummaryResponseCodesParamsQueryType = "DNSKEY"
	AS112SummaryResponseCodesParamsQueryTypeDoa        AS112SummaryResponseCodesParamsQueryType = "DOA"
	AS112SummaryResponseCodesParamsQueryTypeDS         AS112SummaryResponseCodesParamsQueryType = "DS"
	AS112SummaryResponseCodesParamsQueryTypeEid        AS112SummaryResponseCodesParamsQueryType = "EID"
	AS112SummaryResponseCodesParamsQueryTypeEui48      AS112SummaryResponseCodesParamsQueryType = "EUI48"
	AS112SummaryResponseCodesParamsQueryTypeEui64      AS112SummaryResponseCodesParamsQueryType = "EUI64"
	AS112SummaryResponseCodesParamsQueryTypeGpos       AS112SummaryResponseCodesParamsQueryType = "GPOS"
	AS112SummaryResponseCodesParamsQueryTypeGid        AS112SummaryResponseCodesParamsQueryType = "GID"
	AS112SummaryResponseCodesParamsQueryTypeHinfo      AS112SummaryResponseCodesParamsQueryType = "HINFO"
	AS112SummaryResponseCodesParamsQueryTypeHip        AS112SummaryResponseCodesParamsQueryType = "HIP"
	AS112SummaryResponseCodesParamsQueryTypeHTTPS      AS112SummaryResponseCodesParamsQueryType = "HTTPS"
	AS112SummaryResponseCodesParamsQueryTypeIpseckey   AS112SummaryResponseCodesParamsQueryType = "IPSECKEY"
	AS112SummaryResponseCodesParamsQueryTypeIsdn       AS112SummaryResponseCodesParamsQueryType = "ISDN"
	AS112SummaryResponseCodesParamsQueryTypeIxfr       AS112SummaryResponseCodesParamsQueryType = "IXFR"
	AS112SummaryResponseCodesParamsQueryTypeKey        AS112SummaryResponseCodesParamsQueryType = "KEY"
	AS112SummaryResponseCodesParamsQueryTypeKx         AS112SummaryResponseCodesParamsQueryType = "KX"
	AS112SummaryResponseCodesParamsQueryTypeL32        AS112SummaryResponseCodesParamsQueryType = "L32"
	AS112SummaryResponseCodesParamsQueryTypeL64        AS112SummaryResponseCodesParamsQueryType = "L64"
	AS112SummaryResponseCodesParamsQueryTypeLOC        AS112SummaryResponseCodesParamsQueryType = "LOC"
	AS112SummaryResponseCodesParamsQueryTypeLp         AS112SummaryResponseCodesParamsQueryType = "LP"
	AS112SummaryResponseCodesParamsQueryTypeMaila      AS112SummaryResponseCodesParamsQueryType = "MAILA"
	AS112SummaryResponseCodesParamsQueryTypeMailb      AS112SummaryResponseCodesParamsQueryType = "MAILB"
	AS112SummaryResponseCodesParamsQueryTypeMB         AS112SummaryResponseCodesParamsQueryType = "MB"
	AS112SummaryResponseCodesParamsQueryTypeMd         AS112SummaryResponseCodesParamsQueryType = "MD"
	AS112SummaryResponseCodesParamsQueryTypeMf         AS112SummaryResponseCodesParamsQueryType = "MF"
	AS112SummaryResponseCodesParamsQueryTypeMg         AS112SummaryResponseCodesParamsQueryType = "MG"
	AS112SummaryResponseCodesParamsQueryTypeMinfo      AS112SummaryResponseCodesParamsQueryType = "MINFO"
	AS112SummaryResponseCodesParamsQueryTypeMr         AS112SummaryResponseCodesParamsQueryType = "MR"
	AS112SummaryResponseCodesParamsQueryTypeMX         AS112SummaryResponseCodesParamsQueryType = "MX"
	AS112SummaryResponseCodesParamsQueryTypeNAPTR      AS112SummaryResponseCodesParamsQueryType = "NAPTR"
	AS112SummaryResponseCodesParamsQueryTypeNb         AS112SummaryResponseCodesParamsQueryType = "NB"
	AS112SummaryResponseCodesParamsQueryTypeNbstat     AS112SummaryResponseCodesParamsQueryType = "NBSTAT"
	AS112SummaryResponseCodesParamsQueryTypeNid        AS112SummaryResponseCodesParamsQueryType = "NID"
	AS112SummaryResponseCodesParamsQueryTypeNimloc     AS112SummaryResponseCodesParamsQueryType = "NIMLOC"
	AS112SummaryResponseCodesParamsQueryTypeNinfo      AS112SummaryResponseCodesParamsQueryType = "NINFO"
	AS112SummaryResponseCodesParamsQueryTypeNS         AS112SummaryResponseCodesParamsQueryType = "NS"
	AS112SummaryResponseCodesParamsQueryTypeNsap       AS112SummaryResponseCodesParamsQueryType = "NSAP"
	AS112SummaryResponseCodesParamsQueryTypeNsec       AS112SummaryResponseCodesParamsQueryType = "NSEC"
	AS112SummaryResponseCodesParamsQueryTypeNsec3      AS112SummaryResponseCodesParamsQueryType = "NSEC3"
	AS112SummaryResponseCodesParamsQueryTypeNsec3Param AS112SummaryResponseCodesParamsQueryType = "NSEC3PARAM"
	AS112SummaryResponseCodesParamsQueryTypeNull       AS112SummaryResponseCodesParamsQueryType = "NULL"
	AS112SummaryResponseCodesParamsQueryTypeNxt        AS112SummaryResponseCodesParamsQueryType = "NXT"
	AS112SummaryResponseCodesParamsQueryTypeOpenpgpkey AS112SummaryResponseCodesParamsQueryType = "OPENPGPKEY"
	AS112SummaryResponseCodesParamsQueryTypeOpt        AS112SummaryResponseCodesParamsQueryType = "OPT"
	AS112SummaryResponseCodesParamsQueryTypePTR        AS112SummaryResponseCodesParamsQueryType = "PTR"
	AS112SummaryResponseCodesParamsQueryTypePx         AS112SummaryResponseCodesParamsQueryType = "PX"
	AS112SummaryResponseCodesParamsQueryTypeRkey       AS112SummaryResponseCodesParamsQueryType = "RKEY"
	AS112SummaryResponseCodesParamsQueryTypeRp         AS112SummaryResponseCodesParamsQueryType = "RP"
	AS112SummaryResponseCodesParamsQueryTypeRrsig      AS112SummaryResponseCodesParamsQueryType = "RRSIG"
	AS112SummaryResponseCodesParamsQueryTypeRt         AS112SummaryResponseCodesParamsQueryType = "RT"
	AS112SummaryResponseCodesParamsQueryTypeSig        AS112SummaryResponseCodesParamsQueryType = "SIG"
	AS112SummaryResponseCodesParamsQueryTypeSink       AS112SummaryResponseCodesParamsQueryType = "SINK"
	AS112SummaryResponseCodesParamsQueryTypeSMIMEA     AS112SummaryResponseCodesParamsQueryType = "SMIMEA"
	AS112SummaryResponseCodesParamsQueryTypeSOA        AS112SummaryResponseCodesParamsQueryType = "SOA"
	AS112SummaryResponseCodesParamsQueryTypeSPF        AS112SummaryResponseCodesParamsQueryType = "SPF"
	AS112SummaryResponseCodesParamsQueryTypeSRV        AS112SummaryResponseCodesParamsQueryType = "SRV"
	AS112SummaryResponseCodesParamsQueryTypeSSHFP      AS112SummaryResponseCodesParamsQueryType = "SSHFP"
	AS112SummaryResponseCodesParamsQueryTypeSVCB       AS112SummaryResponseCodesParamsQueryType = "SVCB"
	AS112SummaryResponseCodesParamsQueryTypeTa         AS112SummaryResponseCodesParamsQueryType = "TA"
	AS112SummaryResponseCodesParamsQueryTypeTalink     AS112SummaryResponseCodesParamsQueryType = "TALINK"
	AS112SummaryResponseCodesParamsQueryTypeTkey       AS112SummaryResponseCodesParamsQueryType = "TKEY"
	AS112SummaryResponseCodesParamsQueryTypeTLSA       AS112SummaryResponseCodesParamsQueryType = "TLSA"
	AS112SummaryResponseCodesParamsQueryTypeTSIG       AS112SummaryResponseCodesParamsQueryType = "TSIG"
	AS112SummaryResponseCodesParamsQueryTypeTXT        AS112SummaryResponseCodesParamsQueryType = "TXT"
	AS112SummaryResponseCodesParamsQueryTypeUinfo      AS112SummaryResponseCodesParamsQueryType = "UINFO"
	AS112SummaryResponseCodesParamsQueryTypeUID        AS112SummaryResponseCodesParamsQueryType = "UID"
	AS112SummaryResponseCodesParamsQueryTypeUnspec     AS112SummaryResponseCodesParamsQueryType = "UNSPEC"
	AS112SummaryResponseCodesParamsQueryTypeURI        AS112SummaryResponseCodesParamsQueryType = "URI"
	AS112SummaryResponseCodesParamsQueryTypeWks        AS112SummaryResponseCodesParamsQueryType = "WKS"
	AS112SummaryResponseCodesParamsQueryTypeX25        AS112SummaryResponseCodesParamsQueryType = "X25"
	AS112SummaryResponseCodesParamsQueryTypeZonemd     AS112SummaryResponseCodesParamsQueryType = "ZONEMD"
)

func (r AS112SummaryResponseCodesParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryResponseCodesParamsQueryTypeA, AS112SummaryResponseCodesParamsQueryTypeAAAA, AS112SummaryResponseCodesParamsQueryTypeA6, AS112SummaryResponseCodesParamsQueryTypeAfsdb, AS112SummaryResponseCodesParamsQueryTypeAny, AS112SummaryResponseCodesParamsQueryTypeApl, AS112SummaryResponseCodesParamsQueryTypeAtma, AS112SummaryResponseCodesParamsQueryTypeAXFR, AS112SummaryResponseCodesParamsQueryTypeCAA, AS112SummaryResponseCodesParamsQueryTypeCdnskey, AS112SummaryResponseCodesParamsQueryTypeCds, AS112SummaryResponseCodesParamsQueryTypeCERT, AS112SummaryResponseCodesParamsQueryTypeCNAME, AS112SummaryResponseCodesParamsQueryTypeCsync, AS112SummaryResponseCodesParamsQueryTypeDhcid, AS112SummaryResponseCodesParamsQueryTypeDlv, AS112SummaryResponseCodesParamsQueryTypeDname, AS112SummaryResponseCodesParamsQueryTypeDNSKEY, AS112SummaryResponseCodesParamsQueryTypeDoa, AS112SummaryResponseCodesParamsQueryTypeDS, AS112SummaryResponseCodesParamsQueryTypeEid, AS112SummaryResponseCodesParamsQueryTypeEui48, AS112SummaryResponseCodesParamsQueryTypeEui64, AS112SummaryResponseCodesParamsQueryTypeGpos, AS112SummaryResponseCodesParamsQueryTypeGid, AS112SummaryResponseCodesParamsQueryTypeHinfo, AS112SummaryResponseCodesParamsQueryTypeHip, AS112SummaryResponseCodesParamsQueryTypeHTTPS, AS112SummaryResponseCodesParamsQueryTypeIpseckey, AS112SummaryResponseCodesParamsQueryTypeIsdn, AS112SummaryResponseCodesParamsQueryTypeIxfr, AS112SummaryResponseCodesParamsQueryTypeKey, AS112SummaryResponseCodesParamsQueryTypeKx, AS112SummaryResponseCodesParamsQueryTypeL32, AS112SummaryResponseCodesParamsQueryTypeL64, AS112SummaryResponseCodesParamsQueryTypeLOC, AS112SummaryResponseCodesParamsQueryTypeLp, AS112SummaryResponseCodesParamsQueryTypeMaila, AS112SummaryResponseCodesParamsQueryTypeMailb, AS112SummaryResponseCodesParamsQueryTypeMB, AS112SummaryResponseCodesParamsQueryTypeMd, AS112SummaryResponseCodesParamsQueryTypeMf, AS112SummaryResponseCodesParamsQueryTypeMg, AS112SummaryResponseCodesParamsQueryTypeMinfo, AS112SummaryResponseCodesParamsQueryTypeMr, AS112SummaryResponseCodesParamsQueryTypeMX, AS112SummaryResponseCodesParamsQueryTypeNAPTR, AS112SummaryResponseCodesParamsQueryTypeNb, AS112SummaryResponseCodesParamsQueryTypeNbstat, AS112SummaryResponseCodesParamsQueryTypeNid, AS112SummaryResponseCodesParamsQueryTypeNimloc, AS112SummaryResponseCodesParamsQueryTypeNinfo, AS112SummaryResponseCodesParamsQueryTypeNS, AS112SummaryResponseCodesParamsQueryTypeNsap, AS112SummaryResponseCodesParamsQueryTypeNsec, AS112SummaryResponseCodesParamsQueryTypeNsec3, AS112SummaryResponseCodesParamsQueryTypeNsec3Param, AS112SummaryResponseCodesParamsQueryTypeNull, AS112SummaryResponseCodesParamsQueryTypeNxt, AS112SummaryResponseCodesParamsQueryTypeOpenpgpkey, AS112SummaryResponseCodesParamsQueryTypeOpt, AS112SummaryResponseCodesParamsQueryTypePTR, AS112SummaryResponseCodesParamsQueryTypePx, AS112SummaryResponseCodesParamsQueryTypeRkey, AS112SummaryResponseCodesParamsQueryTypeRp, AS112SummaryResponseCodesParamsQueryTypeRrsig, AS112SummaryResponseCodesParamsQueryTypeRt, AS112SummaryResponseCodesParamsQueryTypeSig, AS112SummaryResponseCodesParamsQueryTypeSink, AS112SummaryResponseCodesParamsQueryTypeSMIMEA, AS112SummaryResponseCodesParamsQueryTypeSOA, AS112SummaryResponseCodesParamsQueryTypeSPF, AS112SummaryResponseCodesParamsQueryTypeSRV, AS112SummaryResponseCodesParamsQueryTypeSSHFP, AS112SummaryResponseCodesParamsQueryTypeSVCB, AS112SummaryResponseCodesParamsQueryTypeTa, AS112SummaryResponseCodesParamsQueryTypeTalink, AS112SummaryResponseCodesParamsQueryTypeTkey, AS112SummaryResponseCodesParamsQueryTypeTLSA, AS112SummaryResponseCodesParamsQueryTypeTSIG, AS112SummaryResponseCodesParamsQueryTypeTXT, AS112SummaryResponseCodesParamsQueryTypeUinfo, AS112SummaryResponseCodesParamsQueryTypeUID, AS112SummaryResponseCodesParamsQueryTypeUnspec, AS112SummaryResponseCodesParamsQueryTypeURI, AS112SummaryResponseCodesParamsQueryTypeWks, AS112SummaryResponseCodesParamsQueryTypeX25, AS112SummaryResponseCodesParamsQueryTypeZonemd:
		return true
	}
	return false
}

type AS112SummaryResponseCodesResponseEnvelope struct {
	Result  AS112SummaryResponseCodesResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    as112SummaryResponseCodesResponseEnvelopeJSON `json:"-"`
}

// as112SummaryResponseCodesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112SummaryResponseCodesResponseEnvelope]
type as112SummaryResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryResponseCodesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
