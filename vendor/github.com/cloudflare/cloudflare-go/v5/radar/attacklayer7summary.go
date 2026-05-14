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

// AttackLayer7SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7SummaryService] method instead.
type AttackLayer7SummaryService struct {
	Options []option.RequestOption
}

// NewAttackLayer7SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7SummaryService(opts ...option.RequestOption) (r *AttackLayer7SummaryService) {
	r = &AttackLayer7SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP method.
func (r *AttackLayer7SummaryService) HTTPMethod(ctx context.Context, query AttackLayer7SummaryHTTPMethodParams, opts ...option.RequestOption) (res *AttackLayer7SummaryHTTPMethodResponse, err error) {
	var env AttackLayer7SummaryHTTPMethodResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP version.
func (r *AttackLayer7SummaryService) HTTPVersion(ctx context.Context, query AttackLayer7SummaryHTTPVersionParams, opts ...option.RequestOption) (res *AttackLayer7SummaryHTTPVersionResponse, err error) {
	var env AttackLayer7SummaryHTTPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by targeted industry.
func (r *AttackLayer7SummaryService) Industry(ctx context.Context, query AttackLayer7SummaryIndustryParams, opts ...option.RequestOption) (res *AttackLayer7SummaryIndustryResponse, err error) {
	var env AttackLayer7SummaryIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by IP version.
func (r *AttackLayer7SummaryService) IPVersion(ctx context.Context, query AttackLayer7SummaryIPVersionParams, opts ...option.RequestOption) (res *AttackLayer7SummaryIPVersionResponse, err error) {
	var env AttackLayer7SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by managed rules.
func (r *AttackLayer7SummaryService) ManagedRules(ctx context.Context, query AttackLayer7SummaryManagedRulesParams, opts ...option.RequestOption) (res *AttackLayer7SummaryManagedRulesResponse, err error) {
	var env AttackLayer7SummaryManagedRulesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by mitigation product.
func (r *AttackLayer7SummaryService) MitigationProduct(ctx context.Context, query AttackLayer7SummaryMitigationProductParams, opts ...option.RequestOption) (res *AttackLayer7SummaryMitigationProductResponse, err error) {
	var env AttackLayer7SummaryMitigationProductResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks by targeted vertical.
func (r *AttackLayer7SummaryService) Vertical(ctx context.Context, query AttackLayer7SummaryVerticalParams, opts ...option.RequestOption) (res *AttackLayer7SummaryVerticalResponse, err error) {
	var env AttackLayer7SummaryVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7SummaryHTTPMethodResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryHTTPMethodResponseMeta `json:"meta,required"`
	Summary0 map[string]string                         `json:"summary_0,required"`
	JSON     attackLayer7SummaryHTTPMethodResponseJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPMethodResponse]
type attackLayer7SummaryHTTPMethodResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryHTTPMethodResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryHTTPMethodResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryHTTPMethodResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryHTTPMethodResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryHTTPMethodResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPMethodResponseMeta]
type attackLayer7SummaryHTTPMethodResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo]
type attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryHTTPMethodResponseMetaDateRange]
type attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryHTTPMethodResponseMetaNormalization string

const (
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationPercentage           AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationMin0Max              AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationMinMax               AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRawValues            AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationPercentageChange     AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRollingAverage       AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRatio                AttackLayer7SummaryHTTPMethodResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryHTTPMethodResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodResponseMetaNormalizationPercentage, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationMin0Max, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationMinMax, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRawValues, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationPercentageChange, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRollingAverage, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryHTTPMethodResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  attackLayer7SummaryHTTPMethodResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPMethodResponseMetaUnit]
type attackLayer7SummaryHTTPMethodResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryHTTPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryHTTPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryHTTPVersionResponseJSON     `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryHTTPVersionResponse]
type attackLayer7SummaryHTTPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryHTTPVersionResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryHTTPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryHTTPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryHTTPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryHTTPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPVersionResponseMeta]
type attackLayer7SummaryHTTPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo]
type attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryHTTPVersionResponseMetaDateRange]
type attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryHTTPVersionResponseMetaNormalization string

const (
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationPercentage           AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationMin0Max              AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationMinMax               AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRawValues            AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationPercentageChange     AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRollingAverage       AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRatio                AttackLayer7SummaryHTTPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryHTTPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionResponseMetaNormalizationPercentage, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationMin0Max, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationMinMax, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRawValues, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationPercentageChange, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRollingAverage, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryHTTPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  attackLayer7SummaryHTTPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseMetaUnitJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryHTTPVersionResponseMetaUnit]
type attackLayer7SummaryHTTPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionResponseSummary0 struct {
	HTTP1X string                                             `json:"HTTP/1.x,required"`
	HTTP2  string                                             `json:"HTTP/2,required"`
	HTTP3  string                                             `json:"HTTP/3,required"`
	JSON   attackLayer7SummaryHTTPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseSummary0JSON contains the JSON metadata
// for the struct [AttackLayer7SummaryHTTPVersionResponseSummary0]
type attackLayer7SummaryHTTPVersionResponseSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIndustryResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryIndustryResponseMeta `json:"meta,required"`
	Summary0 map[string]string                       `json:"summary_0,required"`
	JSON     attackLayer7SummaryIndustryResponseJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIndustryResponse]
type attackLayer7SummaryIndustryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryIndustryResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIndustryResponseMeta]
type attackLayer7SummaryIndustryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer7SummaryIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryIndustryResponseMetaConfidenceInfo]
type attackLayer7SummaryIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryIndustryResponseMetaDateRange]
type attackLayer7SummaryIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryIndustryResponseMetaNormalization string

const (
	AttackLayer7SummaryIndustryResponseMetaNormalizationPercentage           AttackLayer7SummaryIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryIndustryResponseMetaNormalizationMin0Max              AttackLayer7SummaryIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryIndustryResponseMetaNormalizationMinMax               AttackLayer7SummaryIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryIndustryResponseMetaNormalizationRawValues            AttackLayer7SummaryIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryIndustryResponseMetaNormalizationPercentageChange     AttackLayer7SummaryIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryIndustryResponseMetaNormalizationRollingAverage       AttackLayer7SummaryIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryIndustryResponseMetaNormalizationRatio                AttackLayer7SummaryIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryResponseMetaNormalizationPercentage, AttackLayer7SummaryIndustryResponseMetaNormalizationMin0Max, AttackLayer7SummaryIndustryResponseMetaNormalizationMinMax, AttackLayer7SummaryIndustryResponseMetaNormalizationRawValues, AttackLayer7SummaryIndustryResponseMetaNormalizationPercentageChange, AttackLayer7SummaryIndustryResponseMetaNormalizationRollingAverage, AttackLayer7SummaryIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer7SummaryIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIndustryResponseMetaUnit]
type attackLayer7SummaryIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer7SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer7SummaryIPVersionResponseJSON     `json:"-"`
}

// attackLayer7SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIPVersionResponse]
type attackLayer7SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryIPVersionResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryIPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryIPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryIPVersionResponseMeta]
type attackLayer7SummaryIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo]
type attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryIPVersionResponseMetaDateRange]
type attackLayer7SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryIPVersionResponseMetaNormalization string

const (
	AttackLayer7SummaryIPVersionResponseMetaNormalizationPercentage           AttackLayer7SummaryIPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationMin0Max              AttackLayer7SummaryIPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationMinMax               AttackLayer7SummaryIPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationRawValues            AttackLayer7SummaryIPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationPercentageChange     AttackLayer7SummaryIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationRollingAverage       AttackLayer7SummaryIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryIPVersionResponseMetaNormalizationRatio                AttackLayer7SummaryIPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionResponseMetaNormalizationPercentage, AttackLayer7SummaryIPVersionResponseMetaNormalizationMin0Max, AttackLayer7SummaryIPVersionResponseMetaNormalizationMinMax, AttackLayer7SummaryIPVersionResponseMetaNormalizationRawValues, AttackLayer7SummaryIPVersionResponseMetaNormalizationPercentageChange, AttackLayer7SummaryIPVersionResponseMetaNormalizationRollingAverage, AttackLayer7SummaryIPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  attackLayer7SummaryIPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIPVersionResponseMetaUnit]
type attackLayer7SummaryIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionResponseSummary0 struct {
	IPv4 string                                           `json:"IPv4,required"`
	IPv6 string                                           `json:"IPv6,required"`
	JSON attackLayer7SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIPVersionResponseSummary0]
type attackLayer7SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryManagedRulesResponseMeta `json:"meta,required"`
	Summary0 map[string]string                           `json:"summary_0,required"`
	JSON     attackLayer7SummaryManagedRulesResponseJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryManagedRulesResponse]
type attackLayer7SummaryManagedRulesResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryManagedRulesResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryManagedRulesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryManagedRulesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryManagedRulesResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryManagedRulesResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryManagedRulesResponseMeta]
type attackLayer7SummaryManagedRulesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo]
type attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryManagedRulesResponseMetaDateRange]
type attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryManagedRulesResponseMetaNormalization string

const (
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationPercentage           AttackLayer7SummaryManagedRulesResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationMin0Max              AttackLayer7SummaryManagedRulesResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationMinMax               AttackLayer7SummaryManagedRulesResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationRawValues            AttackLayer7SummaryManagedRulesResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationPercentageChange     AttackLayer7SummaryManagedRulesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationRollingAverage       AttackLayer7SummaryManagedRulesResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryManagedRulesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryManagedRulesResponseMetaNormalizationRatio                AttackLayer7SummaryManagedRulesResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryManagedRulesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesResponseMetaNormalizationPercentage, AttackLayer7SummaryManagedRulesResponseMetaNormalizationMin0Max, AttackLayer7SummaryManagedRulesResponseMetaNormalizationMinMax, AttackLayer7SummaryManagedRulesResponseMetaNormalizationRawValues, AttackLayer7SummaryManagedRulesResponseMetaNormalizationPercentageChange, AttackLayer7SummaryManagedRulesResponseMetaNormalizationRollingAverage, AttackLayer7SummaryManagedRulesResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryManagedRulesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  attackLayer7SummaryManagedRulesResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseMetaUnitJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryManagedRulesResponseMetaUnit]
type attackLayer7SummaryManagedRulesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryMitigationProductResponseMeta `json:"meta,required"`
	Summary0 map[string]string                                `json:"summary_0,required"`
	JSON     attackLayer7SummaryMitigationProductResponseJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryMitigationProductResponse]
type attackLayer7SummaryMitigationProductResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryMitigationProductResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryMitigationProductResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryMitigationProductResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryMitigationProductResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryMitigationProductResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryMitigationProductResponseMeta]
type attackLayer7SummaryMitigationProductResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo]
type attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryMitigationProductResponseMetaDateRange]
type attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryMitigationProductResponseMetaNormalization string

const (
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationPercentage           AttackLayer7SummaryMitigationProductResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationMin0Max              AttackLayer7SummaryMitigationProductResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationMinMax               AttackLayer7SummaryMitigationProductResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationRawValues            AttackLayer7SummaryMitigationProductResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationPercentageChange     AttackLayer7SummaryMitigationProductResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationRollingAverage       AttackLayer7SummaryMitigationProductResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryMitigationProductResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryMitigationProductResponseMetaNormalizationRatio                AttackLayer7SummaryMitigationProductResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryMitigationProductResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductResponseMetaNormalizationPercentage, AttackLayer7SummaryMitigationProductResponseMetaNormalizationMin0Max, AttackLayer7SummaryMitigationProductResponseMetaNormalizationMinMax, AttackLayer7SummaryMitigationProductResponseMetaNormalizationRawValues, AttackLayer7SummaryMitigationProductResponseMetaNormalizationPercentageChange, AttackLayer7SummaryMitigationProductResponseMetaNormalizationRollingAverage, AttackLayer7SummaryMitigationProductResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryMitigationProductResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  attackLayer7SummaryMitigationProductResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseMetaUnitJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryMitigationProductResponseMetaUnit]
type attackLayer7SummaryMitigationProductResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryVerticalResponse struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryVerticalResponseMeta `json:"meta,required"`
	Summary0 map[string]string                       `json:"summary_0,required"`
	JSON     attackLayer7SummaryVerticalResponseJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryVerticalResponse]
type attackLayer7SummaryVerticalResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryVerticalResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryVerticalResponseMeta]
type attackLayer7SummaryVerticalResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer7SummaryVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryVerticalResponseMetaConfidenceInfo]
type attackLayer7SummaryVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryVerticalResponseMetaDateRange]
type attackLayer7SummaryVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryVerticalResponseMetaNormalization string

const (
	AttackLayer7SummaryVerticalResponseMetaNormalizationPercentage           AttackLayer7SummaryVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryVerticalResponseMetaNormalizationMin0Max              AttackLayer7SummaryVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryVerticalResponseMetaNormalizationMinMax               AttackLayer7SummaryVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryVerticalResponseMetaNormalizationRawValues            AttackLayer7SummaryVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryVerticalResponseMetaNormalizationPercentageChange     AttackLayer7SummaryVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryVerticalResponseMetaNormalizationRollingAverage       AttackLayer7SummaryVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryVerticalResponseMetaNormalizationRatio                AttackLayer7SummaryVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalResponseMetaNormalizationPercentage, AttackLayer7SummaryVerticalResponseMetaNormalizationMin0Max, AttackLayer7SummaryVerticalResponseMetaNormalizationMinMax, AttackLayer7SummaryVerticalResponseMetaNormalizationRawValues, AttackLayer7SummaryVerticalResponseMetaNormalizationPercentageChange, AttackLayer7SummaryVerticalResponseMetaNormalizationRollingAverage, AttackLayer7SummaryVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer7SummaryVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryVerticalResponseMetaUnit]
type attackLayer7SummaryVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPMethodParams struct {
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
	Format param.Field[AttackLayer7SummaryHTTPMethodParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryHTTPMethodParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryHTTPMethodParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryHTTPMethodParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryHTTPMethodParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryHTTPMethodParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryHTTPMethodParamsFormat string

const (
	AttackLayer7SummaryHTTPMethodParamsFormatJson AttackLayer7SummaryHTTPMethodParamsFormat = "JSON"
	AttackLayer7SummaryHTTPMethodParamsFormatCsv  AttackLayer7SummaryHTTPMethodParamsFormat = "CSV"
)

func (r AttackLayer7SummaryHTTPMethodParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsFormatJson, AttackLayer7SummaryHTTPMethodParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsHTTPVersion string

const (
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3 AttackLayer7SummaryHTTPMethodParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryHTTPMethodParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv1, AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv2, AttackLayer7SummaryHTTPMethodParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsIPVersion string

const (
	AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4 AttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv4"
	AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6 AttackLayer7SummaryHTTPMethodParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryHTTPMethodParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsIPVersionIPv4, AttackLayer7SummaryHTTPMethodParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodParamsMitigationProduct string

const (
	AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS               AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF                AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "WAF"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement      AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductAccessRules        AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductIPReputation       AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductAPIShield          AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryHTTPMethodParamsMitigationProductDataLossPrevention AttackLayer7SummaryHTTPMethodParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryHTTPMethodParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPMethodParamsMitigationProductDDoS, AttackLayer7SummaryHTTPMethodParamsMitigationProductWAF, AttackLayer7SummaryHTTPMethodParamsMitigationProductBotManagement, AttackLayer7SummaryHTTPMethodParamsMitigationProductAccessRules, AttackLayer7SummaryHTTPMethodParamsMitigationProductIPReputation, AttackLayer7SummaryHTTPMethodParamsMitigationProductAPIShield, AttackLayer7SummaryHTTPMethodParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPMethodResponseEnvelope struct {
	Result  AttackLayer7SummaryHTTPMethodResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7SummaryHTTPMethodResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryHTTPMethodResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryHTTPMethodResponseEnvelope]
type attackLayer7SummaryHTTPMethodResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPMethodResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPMethodResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryHTTPVersionParams struct {
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
	Format param.Field[AttackLayer7SummaryHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryHTTPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryHTTPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryHTTPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryHTTPVersionParamsFormat string

const (
	AttackLayer7SummaryHTTPVersionParamsFormatJson AttackLayer7SummaryHTTPVersionParamsFormat = "JSON"
	AttackLayer7SummaryHTTPVersionParamsFormatCsv  AttackLayer7SummaryHTTPVersionParamsFormat = "CSV"
)

func (r AttackLayer7SummaryHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsFormatJson, AttackLayer7SummaryHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsHTTPMethod string

const (
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "GET"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POST"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPut             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodHead            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPurge           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodOptions         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPropfind        AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkcol           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPatch           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodACL             AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ACL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBcopy           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBdelete         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBmove           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBpropfind       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBproppatch      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckin         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckout        AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodConnect         AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCopy            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COPY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodLabel           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodLock            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMerge           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkactivity      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkworkspace     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodMove            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodNotify          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodOrderpatch      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodPoll            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "POLL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodProppatch       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodReport          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodSearch          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodSubscribe       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrace           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUncheckout      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnlock          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnsubscribe     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodUpdate          AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodVersioncontrol  AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodBaselinecontrol AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodXmsenumatts     AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcOutData      AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcInData       AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodJson            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "JSON"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodCook            AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "COOK"
	AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrack           AttackLayer7SummaryHTTPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryHTTPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsHTTPMethodGet, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPost, AttackLayer7SummaryHTTPVersionParamsHTTPMethodDelete, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPut, AttackLayer7SummaryHTTPVersionParamsHTTPMethodHead, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPurge, AttackLayer7SummaryHTTPVersionParamsHTTPMethodOptions, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPropfind, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkcol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodACL, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBcopy, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBdelete, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBmove, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBpropfind, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBproppatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckin, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCheckout, AttackLayer7SummaryHTTPVersionParamsHTTPMethodConnect, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCopy, AttackLayer7SummaryHTTPVersionParamsHTTPMethodLabel, AttackLayer7SummaryHTTPVersionParamsHTTPMethodLock, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMerge, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkactivity, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMkworkspace, AttackLayer7SummaryHTTPVersionParamsHTTPMethodMove, AttackLayer7SummaryHTTPVersionParamsHTTPMethodNotify, AttackLayer7SummaryHTTPVersionParamsHTTPMethodOrderpatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodPoll, AttackLayer7SummaryHTTPVersionParamsHTTPMethodProppatch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodReport, AttackLayer7SummaryHTTPVersionParamsHTTPMethodSearch, AttackLayer7SummaryHTTPVersionParamsHTTPMethodSubscribe, AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrace, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUncheckout, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnlock, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUnsubscribe, AttackLayer7SummaryHTTPVersionParamsHTTPMethodUpdate, AttackLayer7SummaryHTTPVersionParamsHTTPMethodVersioncontrol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryHTTPVersionParamsHTTPMethodXmsenumatts, AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcOutData, AttackLayer7SummaryHTTPVersionParamsHTTPMethodRpcInData, AttackLayer7SummaryHTTPVersionParamsHTTPMethodJson, AttackLayer7SummaryHTTPVersionParamsHTTPMethodCook, AttackLayer7SummaryHTTPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsIPVersion string

const (
	AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4 AttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv4"
	AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6 AttackLayer7SummaryHTTPVersionParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsIPVersionIPv4, AttackLayer7SummaryHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionParamsMitigationProduct string

const (
	AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS               AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF                AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "WAF"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement      AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductAccessRules        AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductIPReputation       AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductAPIShield          AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryHTTPVersionParamsMitigationProductDataLossPrevention AttackLayer7SummaryHTTPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryHTTPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryHTTPVersionParamsMitigationProductDDoS, AttackLayer7SummaryHTTPVersionParamsMitigationProductWAF, AttackLayer7SummaryHTTPVersionParamsMitigationProductBotManagement, AttackLayer7SummaryHTTPVersionParamsMitigationProductAccessRules, AttackLayer7SummaryHTTPVersionParamsMitigationProductIPReputation, AttackLayer7SummaryHTTPVersionParamsMitigationProductAPIShield, AttackLayer7SummaryHTTPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryHTTPVersionResponseEnvelope struct {
	Result  AttackLayer7SummaryHTTPVersionResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer7SummaryHTTPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryHTTPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryHTTPVersionResponseEnvelope]
type attackLayer7SummaryHTTPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryHTTPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryHTTPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIndustryParams struct {
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
	Format param.Field[AttackLayer7SummaryIndustryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryIndustryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryIndustryParamsFormat string

const (
	AttackLayer7SummaryIndustryParamsFormatJson AttackLayer7SummaryIndustryParamsFormat = "JSON"
	AttackLayer7SummaryIndustryParamsFormatCsv  AttackLayer7SummaryIndustryParamsFormat = "CSV"
)

func (r AttackLayer7SummaryIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryParamsFormatJson, AttackLayer7SummaryIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryParamsHTTPMethod string

const (
	AttackLayer7SummaryIndustryParamsHTTPMethodGet             AttackLayer7SummaryIndustryParamsHTTPMethod = "GET"
	AttackLayer7SummaryIndustryParamsHTTPMethodPost            AttackLayer7SummaryIndustryParamsHTTPMethod = "POST"
	AttackLayer7SummaryIndustryParamsHTTPMethodDelete          AttackLayer7SummaryIndustryParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryIndustryParamsHTTPMethodPut             AttackLayer7SummaryIndustryParamsHTTPMethod = "PUT"
	AttackLayer7SummaryIndustryParamsHTTPMethodHead            AttackLayer7SummaryIndustryParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryIndustryParamsHTTPMethodPurge           AttackLayer7SummaryIndustryParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryIndustryParamsHTTPMethodOptions         AttackLayer7SummaryIndustryParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryIndustryParamsHTTPMethodPropfind        AttackLayer7SummaryIndustryParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryIndustryParamsHTTPMethodMkcol           AttackLayer7SummaryIndustryParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryIndustryParamsHTTPMethodPatch           AttackLayer7SummaryIndustryParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryIndustryParamsHTTPMethodACL             AttackLayer7SummaryIndustryParamsHTTPMethod = "ACL"
	AttackLayer7SummaryIndustryParamsHTTPMethodBcopy           AttackLayer7SummaryIndustryParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryIndustryParamsHTTPMethodBdelete         AttackLayer7SummaryIndustryParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryIndustryParamsHTTPMethodBmove           AttackLayer7SummaryIndustryParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryIndustryParamsHTTPMethodBpropfind       AttackLayer7SummaryIndustryParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryIndustryParamsHTTPMethodBproppatch      AttackLayer7SummaryIndustryParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryIndustryParamsHTTPMethodCheckin         AttackLayer7SummaryIndustryParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryIndustryParamsHTTPMethodCheckout        AttackLayer7SummaryIndustryParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryIndustryParamsHTTPMethodConnect         AttackLayer7SummaryIndustryParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryIndustryParamsHTTPMethodCopy            AttackLayer7SummaryIndustryParamsHTTPMethod = "COPY"
	AttackLayer7SummaryIndustryParamsHTTPMethodLabel           AttackLayer7SummaryIndustryParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryIndustryParamsHTTPMethodLock            AttackLayer7SummaryIndustryParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryIndustryParamsHTTPMethodMerge           AttackLayer7SummaryIndustryParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryIndustryParamsHTTPMethodMkactivity      AttackLayer7SummaryIndustryParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryIndustryParamsHTTPMethodMkworkspace     AttackLayer7SummaryIndustryParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryIndustryParamsHTTPMethodMove            AttackLayer7SummaryIndustryParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryIndustryParamsHTTPMethodNotify          AttackLayer7SummaryIndustryParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryIndustryParamsHTTPMethodOrderpatch      AttackLayer7SummaryIndustryParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryIndustryParamsHTTPMethodPoll            AttackLayer7SummaryIndustryParamsHTTPMethod = "POLL"
	AttackLayer7SummaryIndustryParamsHTTPMethodProppatch       AttackLayer7SummaryIndustryParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryIndustryParamsHTTPMethodReport          AttackLayer7SummaryIndustryParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryIndustryParamsHTTPMethodSearch          AttackLayer7SummaryIndustryParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryIndustryParamsHTTPMethodSubscribe       AttackLayer7SummaryIndustryParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryIndustryParamsHTTPMethodTrace           AttackLayer7SummaryIndustryParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryIndustryParamsHTTPMethodUncheckout      AttackLayer7SummaryIndustryParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryIndustryParamsHTTPMethodUnlock          AttackLayer7SummaryIndustryParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryIndustryParamsHTTPMethodUnsubscribe     AttackLayer7SummaryIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryIndustryParamsHTTPMethodUpdate          AttackLayer7SummaryIndustryParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryIndustryParamsHTTPMethodVersioncontrol  AttackLayer7SummaryIndustryParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryIndustryParamsHTTPMethodBaselinecontrol AttackLayer7SummaryIndustryParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryIndustryParamsHTTPMethodXmsenumatts     AttackLayer7SummaryIndustryParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryIndustryParamsHTTPMethodRpcOutData      AttackLayer7SummaryIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryIndustryParamsHTTPMethodRpcInData       AttackLayer7SummaryIndustryParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryIndustryParamsHTTPMethodJson            AttackLayer7SummaryIndustryParamsHTTPMethod = "JSON"
	AttackLayer7SummaryIndustryParamsHTTPMethodCook            AttackLayer7SummaryIndustryParamsHTTPMethod = "COOK"
	AttackLayer7SummaryIndustryParamsHTTPMethodTrack           AttackLayer7SummaryIndustryParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryIndustryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryParamsHTTPMethodGet, AttackLayer7SummaryIndustryParamsHTTPMethodPost, AttackLayer7SummaryIndustryParamsHTTPMethodDelete, AttackLayer7SummaryIndustryParamsHTTPMethodPut, AttackLayer7SummaryIndustryParamsHTTPMethodHead, AttackLayer7SummaryIndustryParamsHTTPMethodPurge, AttackLayer7SummaryIndustryParamsHTTPMethodOptions, AttackLayer7SummaryIndustryParamsHTTPMethodPropfind, AttackLayer7SummaryIndustryParamsHTTPMethodMkcol, AttackLayer7SummaryIndustryParamsHTTPMethodPatch, AttackLayer7SummaryIndustryParamsHTTPMethodACL, AttackLayer7SummaryIndustryParamsHTTPMethodBcopy, AttackLayer7SummaryIndustryParamsHTTPMethodBdelete, AttackLayer7SummaryIndustryParamsHTTPMethodBmove, AttackLayer7SummaryIndustryParamsHTTPMethodBpropfind, AttackLayer7SummaryIndustryParamsHTTPMethodBproppatch, AttackLayer7SummaryIndustryParamsHTTPMethodCheckin, AttackLayer7SummaryIndustryParamsHTTPMethodCheckout, AttackLayer7SummaryIndustryParamsHTTPMethodConnect, AttackLayer7SummaryIndustryParamsHTTPMethodCopy, AttackLayer7SummaryIndustryParamsHTTPMethodLabel, AttackLayer7SummaryIndustryParamsHTTPMethodLock, AttackLayer7SummaryIndustryParamsHTTPMethodMerge, AttackLayer7SummaryIndustryParamsHTTPMethodMkactivity, AttackLayer7SummaryIndustryParamsHTTPMethodMkworkspace, AttackLayer7SummaryIndustryParamsHTTPMethodMove, AttackLayer7SummaryIndustryParamsHTTPMethodNotify, AttackLayer7SummaryIndustryParamsHTTPMethodOrderpatch, AttackLayer7SummaryIndustryParamsHTTPMethodPoll, AttackLayer7SummaryIndustryParamsHTTPMethodProppatch, AttackLayer7SummaryIndustryParamsHTTPMethodReport, AttackLayer7SummaryIndustryParamsHTTPMethodSearch, AttackLayer7SummaryIndustryParamsHTTPMethodSubscribe, AttackLayer7SummaryIndustryParamsHTTPMethodTrace, AttackLayer7SummaryIndustryParamsHTTPMethodUncheckout, AttackLayer7SummaryIndustryParamsHTTPMethodUnlock, AttackLayer7SummaryIndustryParamsHTTPMethodUnsubscribe, AttackLayer7SummaryIndustryParamsHTTPMethodUpdate, AttackLayer7SummaryIndustryParamsHTTPMethodVersioncontrol, AttackLayer7SummaryIndustryParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryIndustryParamsHTTPMethodXmsenumatts, AttackLayer7SummaryIndustryParamsHTTPMethodRpcOutData, AttackLayer7SummaryIndustryParamsHTTPMethodRpcInData, AttackLayer7SummaryIndustryParamsHTTPMethodJson, AttackLayer7SummaryIndustryParamsHTTPMethodCook, AttackLayer7SummaryIndustryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryParamsHTTPVersion string

const (
	AttackLayer7SummaryIndustryParamsHTTPVersionHttPv1 AttackLayer7SummaryIndustryParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryIndustryParamsHTTPVersionHttPv2 AttackLayer7SummaryIndustryParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryIndustryParamsHTTPVersionHttPv3 AttackLayer7SummaryIndustryParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryIndustryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryParamsHTTPVersionHttPv1, AttackLayer7SummaryIndustryParamsHTTPVersionHttPv2, AttackLayer7SummaryIndustryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryParamsIPVersion string

const (
	AttackLayer7SummaryIndustryParamsIPVersionIPv4 AttackLayer7SummaryIndustryParamsIPVersion = "IPv4"
	AttackLayer7SummaryIndustryParamsIPVersionIPv6 AttackLayer7SummaryIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryParamsIPVersionIPv4, AttackLayer7SummaryIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryParamsMitigationProduct string

const (
	AttackLayer7SummaryIndustryParamsMitigationProductDDoS               AttackLayer7SummaryIndustryParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryIndustryParamsMitigationProductWAF                AttackLayer7SummaryIndustryParamsMitigationProduct = "WAF"
	AttackLayer7SummaryIndustryParamsMitigationProductBotManagement      AttackLayer7SummaryIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryIndustryParamsMitigationProductAccessRules        AttackLayer7SummaryIndustryParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryIndustryParamsMitigationProductIPReputation       AttackLayer7SummaryIndustryParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryIndustryParamsMitigationProductAPIShield          AttackLayer7SummaryIndustryParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryIndustryParamsMitigationProductDataLossPrevention AttackLayer7SummaryIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryIndustryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIndustryParamsMitigationProductDDoS, AttackLayer7SummaryIndustryParamsMitigationProductWAF, AttackLayer7SummaryIndustryParamsMitigationProductBotManagement, AttackLayer7SummaryIndustryParamsMitigationProductAccessRules, AttackLayer7SummaryIndustryParamsMitigationProductIPReputation, AttackLayer7SummaryIndustryParamsMitigationProductAPIShield, AttackLayer7SummaryIndustryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryIndustryResponseEnvelope struct {
	Result  AttackLayer7SummaryIndustryResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer7SummaryIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryIndustryResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIndustryResponseEnvelope]
type attackLayer7SummaryIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryIPVersionParams struct {
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
	Format param.Field[AttackLayer7SummaryIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryIPVersionParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryIPVersionParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryIPVersionParamsFormat string

const (
	AttackLayer7SummaryIPVersionParamsFormatJson AttackLayer7SummaryIPVersionParamsFormat = "JSON"
	AttackLayer7SummaryIPVersionParamsFormatCsv  AttackLayer7SummaryIPVersionParamsFormat = "CSV"
)

func (r AttackLayer7SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsFormatJson, AttackLayer7SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsHTTPMethod string

const (
	AttackLayer7SummaryIPVersionParamsHTTPMethodGet             AttackLayer7SummaryIPVersionParamsHTTPMethod = "GET"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPost            AttackLayer7SummaryIPVersionParamsHTTPMethod = "POST"
	AttackLayer7SummaryIPVersionParamsHTTPMethodDelete          AttackLayer7SummaryIPVersionParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPut             AttackLayer7SummaryIPVersionParamsHTTPMethod = "PUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodHead            AttackLayer7SummaryIPVersionParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPurge           AttackLayer7SummaryIPVersionParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodOptions         AttackLayer7SummaryIPVersionParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPropfind        AttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkcol           AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPatch           AttackLayer7SummaryIPVersionParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodACL             AttackLayer7SummaryIPVersionParamsHTTPMethod = "ACL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBcopy           AttackLayer7SummaryIPVersionParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBdelete         AttackLayer7SummaryIPVersionParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBmove           AttackLayer7SummaryIPVersionParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBpropfind       AttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBproppatch      AttackLayer7SummaryIPVersionParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCheckin         AttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCheckout        AttackLayer7SummaryIPVersionParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodConnect         AttackLayer7SummaryIPVersionParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCopy            AttackLayer7SummaryIPVersionParamsHTTPMethod = "COPY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodLabel           AttackLayer7SummaryIPVersionParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodLock            AttackLayer7SummaryIPVersionParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMerge           AttackLayer7SummaryIPVersionParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkactivity      AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMkworkspace     AttackLayer7SummaryIPVersionParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodMove            AttackLayer7SummaryIPVersionParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodNotify          AttackLayer7SummaryIPVersionParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryIPVersionParamsHTTPMethodOrderpatch      AttackLayer7SummaryIPVersionParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodPoll            AttackLayer7SummaryIPVersionParamsHTTPMethod = "POLL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodProppatch       AttackLayer7SummaryIPVersionParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodReport          AttackLayer7SummaryIPVersionParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodSearch          AttackLayer7SummaryIPVersionParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryIPVersionParamsHTTPMethodSubscribe       AttackLayer7SummaryIPVersionParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodTrace           AttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUncheckout      AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUnlock          AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUnsubscribe     AttackLayer7SummaryIPVersionParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodUpdate          AttackLayer7SummaryIPVersionParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryIPVersionParamsHTTPMethodVersioncontrol  AttackLayer7SummaryIPVersionParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodBaselinecontrol AttackLayer7SummaryIPVersionParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryIPVersionParamsHTTPMethodXmsenumatts     AttackLayer7SummaryIPVersionParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryIPVersionParamsHTTPMethodRpcOutData      AttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryIPVersionParamsHTTPMethodRpcInData       AttackLayer7SummaryIPVersionParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryIPVersionParamsHTTPMethodJson            AttackLayer7SummaryIPVersionParamsHTTPMethod = "JSON"
	AttackLayer7SummaryIPVersionParamsHTTPMethodCook            AttackLayer7SummaryIPVersionParamsHTTPMethod = "COOK"
	AttackLayer7SummaryIPVersionParamsHTTPMethodTrack           AttackLayer7SummaryIPVersionParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryIPVersionParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsHTTPMethodGet, AttackLayer7SummaryIPVersionParamsHTTPMethodPost, AttackLayer7SummaryIPVersionParamsHTTPMethodDelete, AttackLayer7SummaryIPVersionParamsHTTPMethodPut, AttackLayer7SummaryIPVersionParamsHTTPMethodHead, AttackLayer7SummaryIPVersionParamsHTTPMethodPurge, AttackLayer7SummaryIPVersionParamsHTTPMethodOptions, AttackLayer7SummaryIPVersionParamsHTTPMethodPropfind, AttackLayer7SummaryIPVersionParamsHTTPMethodMkcol, AttackLayer7SummaryIPVersionParamsHTTPMethodPatch, AttackLayer7SummaryIPVersionParamsHTTPMethodACL, AttackLayer7SummaryIPVersionParamsHTTPMethodBcopy, AttackLayer7SummaryIPVersionParamsHTTPMethodBdelete, AttackLayer7SummaryIPVersionParamsHTTPMethodBmove, AttackLayer7SummaryIPVersionParamsHTTPMethodBpropfind, AttackLayer7SummaryIPVersionParamsHTTPMethodBproppatch, AttackLayer7SummaryIPVersionParamsHTTPMethodCheckin, AttackLayer7SummaryIPVersionParamsHTTPMethodCheckout, AttackLayer7SummaryIPVersionParamsHTTPMethodConnect, AttackLayer7SummaryIPVersionParamsHTTPMethodCopy, AttackLayer7SummaryIPVersionParamsHTTPMethodLabel, AttackLayer7SummaryIPVersionParamsHTTPMethodLock, AttackLayer7SummaryIPVersionParamsHTTPMethodMerge, AttackLayer7SummaryIPVersionParamsHTTPMethodMkactivity, AttackLayer7SummaryIPVersionParamsHTTPMethodMkworkspace, AttackLayer7SummaryIPVersionParamsHTTPMethodMove, AttackLayer7SummaryIPVersionParamsHTTPMethodNotify, AttackLayer7SummaryIPVersionParamsHTTPMethodOrderpatch, AttackLayer7SummaryIPVersionParamsHTTPMethodPoll, AttackLayer7SummaryIPVersionParamsHTTPMethodProppatch, AttackLayer7SummaryIPVersionParamsHTTPMethodReport, AttackLayer7SummaryIPVersionParamsHTTPMethodSearch, AttackLayer7SummaryIPVersionParamsHTTPMethodSubscribe, AttackLayer7SummaryIPVersionParamsHTTPMethodTrace, AttackLayer7SummaryIPVersionParamsHTTPMethodUncheckout, AttackLayer7SummaryIPVersionParamsHTTPMethodUnlock, AttackLayer7SummaryIPVersionParamsHTTPMethodUnsubscribe, AttackLayer7SummaryIPVersionParamsHTTPMethodUpdate, AttackLayer7SummaryIPVersionParamsHTTPMethodVersioncontrol, AttackLayer7SummaryIPVersionParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryIPVersionParamsHTTPMethodXmsenumatts, AttackLayer7SummaryIPVersionParamsHTTPMethodRpcOutData, AttackLayer7SummaryIPVersionParamsHTTPMethodRpcInData, AttackLayer7SummaryIPVersionParamsHTTPMethodJson, AttackLayer7SummaryIPVersionParamsHTTPMethodCook, AttackLayer7SummaryIPVersionParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsHTTPVersion string

const (
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3 AttackLayer7SummaryIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv1, AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv2, AttackLayer7SummaryIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionParamsMitigationProduct string

const (
	AttackLayer7SummaryIPVersionParamsMitigationProductDDoS               AttackLayer7SummaryIPVersionParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryIPVersionParamsMitigationProductWAF                AttackLayer7SummaryIPVersionParamsMitigationProduct = "WAF"
	AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement      AttackLayer7SummaryIPVersionParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryIPVersionParamsMitigationProductAccessRules        AttackLayer7SummaryIPVersionParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryIPVersionParamsMitigationProductIPReputation       AttackLayer7SummaryIPVersionParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryIPVersionParamsMitigationProductAPIShield          AttackLayer7SummaryIPVersionParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryIPVersionParamsMitigationProductDataLossPrevention AttackLayer7SummaryIPVersionParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryIPVersionParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryIPVersionParamsMitigationProductDDoS, AttackLayer7SummaryIPVersionParamsMitigationProductWAF, AttackLayer7SummaryIPVersionParamsMitigationProductBotManagement, AttackLayer7SummaryIPVersionParamsMitigationProductAccessRules, AttackLayer7SummaryIPVersionParamsMitigationProductIPReputation, AttackLayer7SummaryIPVersionParamsMitigationProductAPIShield, AttackLayer7SummaryIPVersionParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryIPVersionResponseEnvelope struct {
	Result  AttackLayer7SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    attackLayer7SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryIPVersionResponseEnvelope]
type attackLayer7SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryManagedRulesParams struct {
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
	Format param.Field[AttackLayer7SummaryManagedRulesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryManagedRulesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryManagedRulesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryManagedRulesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryManagedRulesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryManagedRulesParams]'s query parameters
// as `url.Values`.
func (r AttackLayer7SummaryManagedRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryManagedRulesParamsFormat string

const (
	AttackLayer7SummaryManagedRulesParamsFormatJson AttackLayer7SummaryManagedRulesParamsFormat = "JSON"
	AttackLayer7SummaryManagedRulesParamsFormatCsv  AttackLayer7SummaryManagedRulesParamsFormat = "CSV"
)

func (r AttackLayer7SummaryManagedRulesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsFormatJson, AttackLayer7SummaryManagedRulesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsHTTPMethod string

const (
	AttackLayer7SummaryManagedRulesParamsHTTPMethodGet             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "GET"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPost            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "POST"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPut             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodHead            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPurge           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodOptions         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPropfind        AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkcol           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPatch           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodACL             AttackLayer7SummaryManagedRulesParamsHTTPMethod = "ACL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBcopy           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBdelete         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBmove           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBpropfind       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBproppatch      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckin         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckout        AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodConnect         AttackLayer7SummaryManagedRulesParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCopy            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "COPY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodLabel           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodLock            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMerge           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkactivity      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMkworkspace     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodMove            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodNotify          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodOrderpatch      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodPoll            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "POLL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodProppatch       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodReport          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodSearch          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodSubscribe       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodTrace           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUncheckout      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUnlock          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUnsubscribe     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodUpdate          AttackLayer7SummaryManagedRulesParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodVersioncontrol  AttackLayer7SummaryManagedRulesParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodBaselinecontrol AttackLayer7SummaryManagedRulesParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodXmsenumatts     AttackLayer7SummaryManagedRulesParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcOutData      AttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcInData       AttackLayer7SummaryManagedRulesParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodJson            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "JSON"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodCook            AttackLayer7SummaryManagedRulesParamsHTTPMethod = "COOK"
	AttackLayer7SummaryManagedRulesParamsHTTPMethodTrack           AttackLayer7SummaryManagedRulesParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryManagedRulesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsHTTPMethodGet, AttackLayer7SummaryManagedRulesParamsHTTPMethodPost, AttackLayer7SummaryManagedRulesParamsHTTPMethodDelete, AttackLayer7SummaryManagedRulesParamsHTTPMethodPut, AttackLayer7SummaryManagedRulesParamsHTTPMethodHead, AttackLayer7SummaryManagedRulesParamsHTTPMethodPurge, AttackLayer7SummaryManagedRulesParamsHTTPMethodOptions, AttackLayer7SummaryManagedRulesParamsHTTPMethodPropfind, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkcol, AttackLayer7SummaryManagedRulesParamsHTTPMethodPatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodACL, AttackLayer7SummaryManagedRulesParamsHTTPMethodBcopy, AttackLayer7SummaryManagedRulesParamsHTTPMethodBdelete, AttackLayer7SummaryManagedRulesParamsHTTPMethodBmove, AttackLayer7SummaryManagedRulesParamsHTTPMethodBpropfind, AttackLayer7SummaryManagedRulesParamsHTTPMethodBproppatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckin, AttackLayer7SummaryManagedRulesParamsHTTPMethodCheckout, AttackLayer7SummaryManagedRulesParamsHTTPMethodConnect, AttackLayer7SummaryManagedRulesParamsHTTPMethodCopy, AttackLayer7SummaryManagedRulesParamsHTTPMethodLabel, AttackLayer7SummaryManagedRulesParamsHTTPMethodLock, AttackLayer7SummaryManagedRulesParamsHTTPMethodMerge, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkactivity, AttackLayer7SummaryManagedRulesParamsHTTPMethodMkworkspace, AttackLayer7SummaryManagedRulesParamsHTTPMethodMove, AttackLayer7SummaryManagedRulesParamsHTTPMethodNotify, AttackLayer7SummaryManagedRulesParamsHTTPMethodOrderpatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodPoll, AttackLayer7SummaryManagedRulesParamsHTTPMethodProppatch, AttackLayer7SummaryManagedRulesParamsHTTPMethodReport, AttackLayer7SummaryManagedRulesParamsHTTPMethodSearch, AttackLayer7SummaryManagedRulesParamsHTTPMethodSubscribe, AttackLayer7SummaryManagedRulesParamsHTTPMethodTrace, AttackLayer7SummaryManagedRulesParamsHTTPMethodUncheckout, AttackLayer7SummaryManagedRulesParamsHTTPMethodUnlock, AttackLayer7SummaryManagedRulesParamsHTTPMethodUnsubscribe, AttackLayer7SummaryManagedRulesParamsHTTPMethodUpdate, AttackLayer7SummaryManagedRulesParamsHTTPMethodVersioncontrol, AttackLayer7SummaryManagedRulesParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryManagedRulesParamsHTTPMethodXmsenumatts, AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcOutData, AttackLayer7SummaryManagedRulesParamsHTTPMethodRpcInData, AttackLayer7SummaryManagedRulesParamsHTTPMethodJson, AttackLayer7SummaryManagedRulesParamsHTTPMethodCook, AttackLayer7SummaryManagedRulesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsHTTPVersion string

const (
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3 AttackLayer7SummaryManagedRulesParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryManagedRulesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv1, AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv2, AttackLayer7SummaryManagedRulesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsIPVersion string

const (
	AttackLayer7SummaryManagedRulesParamsIPVersionIPv4 AttackLayer7SummaryManagedRulesParamsIPVersion = "IPv4"
	AttackLayer7SummaryManagedRulesParamsIPVersionIPv6 AttackLayer7SummaryManagedRulesParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryManagedRulesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsIPVersionIPv4, AttackLayer7SummaryManagedRulesParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesParamsMitigationProduct string

const (
	AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS               AttackLayer7SummaryManagedRulesParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryManagedRulesParamsMitigationProductWAF                AttackLayer7SummaryManagedRulesParamsMitigationProduct = "WAF"
	AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement      AttackLayer7SummaryManagedRulesParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryManagedRulesParamsMitigationProductAccessRules        AttackLayer7SummaryManagedRulesParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryManagedRulesParamsMitigationProductIPReputation       AttackLayer7SummaryManagedRulesParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryManagedRulesParamsMitigationProductAPIShield          AttackLayer7SummaryManagedRulesParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryManagedRulesParamsMitigationProductDataLossPrevention AttackLayer7SummaryManagedRulesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryManagedRulesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryManagedRulesParamsMitigationProductDDoS, AttackLayer7SummaryManagedRulesParamsMitigationProductWAF, AttackLayer7SummaryManagedRulesParamsMitigationProductBotManagement, AttackLayer7SummaryManagedRulesParamsMitigationProductAccessRules, AttackLayer7SummaryManagedRulesParamsMitigationProductIPReputation, AttackLayer7SummaryManagedRulesParamsMitigationProductAPIShield, AttackLayer7SummaryManagedRulesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryManagedRulesResponseEnvelope struct {
	Result  AttackLayer7SummaryManagedRulesResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    attackLayer7SummaryManagedRulesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryManagedRulesResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryManagedRulesResponseEnvelope]
type attackLayer7SummaryManagedRulesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryManagedRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryManagedRulesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryMitigationProductParams struct {
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
	Format param.Field[AttackLayer7SummaryMitigationProductParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryMitigationProductParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryMitigationProductParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryMitigationProductParamsIPVersion] `query:"ipVersion"`
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
}

// URLQuery serializes [AttackLayer7SummaryMitigationProductParams]'s query
// parameters as `url.Values`.
func (r AttackLayer7SummaryMitigationProductParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryMitigationProductParamsFormat string

const (
	AttackLayer7SummaryMitigationProductParamsFormatJson AttackLayer7SummaryMitigationProductParamsFormat = "JSON"
	AttackLayer7SummaryMitigationProductParamsFormatCsv  AttackLayer7SummaryMitigationProductParamsFormat = "CSV"
)

func (r AttackLayer7SummaryMitigationProductParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsFormatJson, AttackLayer7SummaryMitigationProductParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsHTTPMethod string

const (
	AttackLayer7SummaryMitigationProductParamsHTTPMethodGet             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "GET"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPost            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "POST"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPut             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodHead            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPurge           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodOptions         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPropfind        AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkcol           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPatch           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodACL             AttackLayer7SummaryMitigationProductParamsHTTPMethod = "ACL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBcopy           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBdelete         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBmove           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBpropfind       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBproppatch      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckin         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckout        AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodConnect         AttackLayer7SummaryMitigationProductParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCopy            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "COPY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodLabel           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodLock            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMerge           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkactivity      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMkworkspace     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodMove            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodNotify          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodOrderpatch      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodPoll            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "POLL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodProppatch       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodReport          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodSearch          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodSubscribe       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodTrace           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUncheckout      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUnlock          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUnsubscribe     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodUpdate          AttackLayer7SummaryMitigationProductParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodVersioncontrol  AttackLayer7SummaryMitigationProductParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodBaselinecontrol AttackLayer7SummaryMitigationProductParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodXmsenumatts     AttackLayer7SummaryMitigationProductParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcOutData      AttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcInData       AttackLayer7SummaryMitigationProductParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodJson            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "JSON"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodCook            AttackLayer7SummaryMitigationProductParamsHTTPMethod = "COOK"
	AttackLayer7SummaryMitigationProductParamsHTTPMethodTrack           AttackLayer7SummaryMitigationProductParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryMitigationProductParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsHTTPMethodGet, AttackLayer7SummaryMitigationProductParamsHTTPMethodPost, AttackLayer7SummaryMitigationProductParamsHTTPMethodDelete, AttackLayer7SummaryMitigationProductParamsHTTPMethodPut, AttackLayer7SummaryMitigationProductParamsHTTPMethodHead, AttackLayer7SummaryMitigationProductParamsHTTPMethodPurge, AttackLayer7SummaryMitigationProductParamsHTTPMethodOptions, AttackLayer7SummaryMitigationProductParamsHTTPMethodPropfind, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkcol, AttackLayer7SummaryMitigationProductParamsHTTPMethodPatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodACL, AttackLayer7SummaryMitigationProductParamsHTTPMethodBcopy, AttackLayer7SummaryMitigationProductParamsHTTPMethodBdelete, AttackLayer7SummaryMitigationProductParamsHTTPMethodBmove, AttackLayer7SummaryMitigationProductParamsHTTPMethodBpropfind, AttackLayer7SummaryMitigationProductParamsHTTPMethodBproppatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckin, AttackLayer7SummaryMitigationProductParamsHTTPMethodCheckout, AttackLayer7SummaryMitigationProductParamsHTTPMethodConnect, AttackLayer7SummaryMitigationProductParamsHTTPMethodCopy, AttackLayer7SummaryMitigationProductParamsHTTPMethodLabel, AttackLayer7SummaryMitigationProductParamsHTTPMethodLock, AttackLayer7SummaryMitigationProductParamsHTTPMethodMerge, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkactivity, AttackLayer7SummaryMitigationProductParamsHTTPMethodMkworkspace, AttackLayer7SummaryMitigationProductParamsHTTPMethodMove, AttackLayer7SummaryMitigationProductParamsHTTPMethodNotify, AttackLayer7SummaryMitigationProductParamsHTTPMethodOrderpatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodPoll, AttackLayer7SummaryMitigationProductParamsHTTPMethodProppatch, AttackLayer7SummaryMitigationProductParamsHTTPMethodReport, AttackLayer7SummaryMitigationProductParamsHTTPMethodSearch, AttackLayer7SummaryMitigationProductParamsHTTPMethodSubscribe, AttackLayer7SummaryMitigationProductParamsHTTPMethodTrace, AttackLayer7SummaryMitigationProductParamsHTTPMethodUncheckout, AttackLayer7SummaryMitigationProductParamsHTTPMethodUnlock, AttackLayer7SummaryMitigationProductParamsHTTPMethodUnsubscribe, AttackLayer7SummaryMitigationProductParamsHTTPMethodUpdate, AttackLayer7SummaryMitigationProductParamsHTTPMethodVersioncontrol, AttackLayer7SummaryMitigationProductParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryMitigationProductParamsHTTPMethodXmsenumatts, AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcOutData, AttackLayer7SummaryMitigationProductParamsHTTPMethodRpcInData, AttackLayer7SummaryMitigationProductParamsHTTPMethodJson, AttackLayer7SummaryMitigationProductParamsHTTPMethodCook, AttackLayer7SummaryMitigationProductParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsHTTPVersion string

const (
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3 AttackLayer7SummaryMitigationProductParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryMitigationProductParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv1, AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv2, AttackLayer7SummaryMitigationProductParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductParamsIPVersion string

const (
	AttackLayer7SummaryMitigationProductParamsIPVersionIPv4 AttackLayer7SummaryMitigationProductParamsIPVersion = "IPv4"
	AttackLayer7SummaryMitigationProductParamsIPVersionIPv6 AttackLayer7SummaryMitigationProductParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryMitigationProductParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryMitigationProductParamsIPVersionIPv4, AttackLayer7SummaryMitigationProductParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryMitigationProductResponseEnvelope struct {
	Result  AttackLayer7SummaryMitigationProductResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    attackLayer7SummaryMitigationProductResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryMitigationProductResponseEnvelopeJSON contains the JSON
// metadata for the struct [AttackLayer7SummaryMitigationProductResponseEnvelope]
type attackLayer7SummaryMitigationProductResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryMitigationProductResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryMitigationProductResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryVerticalParams struct {
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
	Format param.Field[AttackLayer7SummaryVerticalParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryVerticalParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryVerticalParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryVerticalParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryVerticalParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7SummaryVerticalParamsFormat string

const (
	AttackLayer7SummaryVerticalParamsFormatJson AttackLayer7SummaryVerticalParamsFormat = "JSON"
	AttackLayer7SummaryVerticalParamsFormatCsv  AttackLayer7SummaryVerticalParamsFormat = "CSV"
)

func (r AttackLayer7SummaryVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalParamsFormatJson, AttackLayer7SummaryVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalParamsHTTPMethod string

const (
	AttackLayer7SummaryVerticalParamsHTTPMethodGet             AttackLayer7SummaryVerticalParamsHTTPMethod = "GET"
	AttackLayer7SummaryVerticalParamsHTTPMethodPost            AttackLayer7SummaryVerticalParamsHTTPMethod = "POST"
	AttackLayer7SummaryVerticalParamsHTTPMethodDelete          AttackLayer7SummaryVerticalParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryVerticalParamsHTTPMethodPut             AttackLayer7SummaryVerticalParamsHTTPMethod = "PUT"
	AttackLayer7SummaryVerticalParamsHTTPMethodHead            AttackLayer7SummaryVerticalParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryVerticalParamsHTTPMethodPurge           AttackLayer7SummaryVerticalParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryVerticalParamsHTTPMethodOptions         AttackLayer7SummaryVerticalParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryVerticalParamsHTTPMethodPropfind        AttackLayer7SummaryVerticalParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryVerticalParamsHTTPMethodMkcol           AttackLayer7SummaryVerticalParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryVerticalParamsHTTPMethodPatch           AttackLayer7SummaryVerticalParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryVerticalParamsHTTPMethodACL             AttackLayer7SummaryVerticalParamsHTTPMethod = "ACL"
	AttackLayer7SummaryVerticalParamsHTTPMethodBcopy           AttackLayer7SummaryVerticalParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryVerticalParamsHTTPMethodBdelete         AttackLayer7SummaryVerticalParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryVerticalParamsHTTPMethodBmove           AttackLayer7SummaryVerticalParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryVerticalParamsHTTPMethodBpropfind       AttackLayer7SummaryVerticalParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryVerticalParamsHTTPMethodBproppatch      AttackLayer7SummaryVerticalParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryVerticalParamsHTTPMethodCheckin         AttackLayer7SummaryVerticalParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryVerticalParamsHTTPMethodCheckout        AttackLayer7SummaryVerticalParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryVerticalParamsHTTPMethodConnect         AttackLayer7SummaryVerticalParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryVerticalParamsHTTPMethodCopy            AttackLayer7SummaryVerticalParamsHTTPMethod = "COPY"
	AttackLayer7SummaryVerticalParamsHTTPMethodLabel           AttackLayer7SummaryVerticalParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryVerticalParamsHTTPMethodLock            AttackLayer7SummaryVerticalParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryVerticalParamsHTTPMethodMerge           AttackLayer7SummaryVerticalParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryVerticalParamsHTTPMethodMkactivity      AttackLayer7SummaryVerticalParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryVerticalParamsHTTPMethodMkworkspace     AttackLayer7SummaryVerticalParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryVerticalParamsHTTPMethodMove            AttackLayer7SummaryVerticalParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryVerticalParamsHTTPMethodNotify          AttackLayer7SummaryVerticalParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryVerticalParamsHTTPMethodOrderpatch      AttackLayer7SummaryVerticalParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryVerticalParamsHTTPMethodPoll            AttackLayer7SummaryVerticalParamsHTTPMethod = "POLL"
	AttackLayer7SummaryVerticalParamsHTTPMethodProppatch       AttackLayer7SummaryVerticalParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryVerticalParamsHTTPMethodReport          AttackLayer7SummaryVerticalParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryVerticalParamsHTTPMethodSearch          AttackLayer7SummaryVerticalParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryVerticalParamsHTTPMethodSubscribe       AttackLayer7SummaryVerticalParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryVerticalParamsHTTPMethodTrace           AttackLayer7SummaryVerticalParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryVerticalParamsHTTPMethodUncheckout      AttackLayer7SummaryVerticalParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryVerticalParamsHTTPMethodUnlock          AttackLayer7SummaryVerticalParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryVerticalParamsHTTPMethodUnsubscribe     AttackLayer7SummaryVerticalParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryVerticalParamsHTTPMethodUpdate          AttackLayer7SummaryVerticalParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryVerticalParamsHTTPMethodVersioncontrol  AttackLayer7SummaryVerticalParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryVerticalParamsHTTPMethodBaselinecontrol AttackLayer7SummaryVerticalParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryVerticalParamsHTTPMethodXmsenumatts     AttackLayer7SummaryVerticalParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryVerticalParamsHTTPMethodRpcOutData      AttackLayer7SummaryVerticalParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryVerticalParamsHTTPMethodRpcInData       AttackLayer7SummaryVerticalParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryVerticalParamsHTTPMethodJson            AttackLayer7SummaryVerticalParamsHTTPMethod = "JSON"
	AttackLayer7SummaryVerticalParamsHTTPMethodCook            AttackLayer7SummaryVerticalParamsHTTPMethod = "COOK"
	AttackLayer7SummaryVerticalParamsHTTPMethodTrack           AttackLayer7SummaryVerticalParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryVerticalParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalParamsHTTPMethodGet, AttackLayer7SummaryVerticalParamsHTTPMethodPost, AttackLayer7SummaryVerticalParamsHTTPMethodDelete, AttackLayer7SummaryVerticalParamsHTTPMethodPut, AttackLayer7SummaryVerticalParamsHTTPMethodHead, AttackLayer7SummaryVerticalParamsHTTPMethodPurge, AttackLayer7SummaryVerticalParamsHTTPMethodOptions, AttackLayer7SummaryVerticalParamsHTTPMethodPropfind, AttackLayer7SummaryVerticalParamsHTTPMethodMkcol, AttackLayer7SummaryVerticalParamsHTTPMethodPatch, AttackLayer7SummaryVerticalParamsHTTPMethodACL, AttackLayer7SummaryVerticalParamsHTTPMethodBcopy, AttackLayer7SummaryVerticalParamsHTTPMethodBdelete, AttackLayer7SummaryVerticalParamsHTTPMethodBmove, AttackLayer7SummaryVerticalParamsHTTPMethodBpropfind, AttackLayer7SummaryVerticalParamsHTTPMethodBproppatch, AttackLayer7SummaryVerticalParamsHTTPMethodCheckin, AttackLayer7SummaryVerticalParamsHTTPMethodCheckout, AttackLayer7SummaryVerticalParamsHTTPMethodConnect, AttackLayer7SummaryVerticalParamsHTTPMethodCopy, AttackLayer7SummaryVerticalParamsHTTPMethodLabel, AttackLayer7SummaryVerticalParamsHTTPMethodLock, AttackLayer7SummaryVerticalParamsHTTPMethodMerge, AttackLayer7SummaryVerticalParamsHTTPMethodMkactivity, AttackLayer7SummaryVerticalParamsHTTPMethodMkworkspace, AttackLayer7SummaryVerticalParamsHTTPMethodMove, AttackLayer7SummaryVerticalParamsHTTPMethodNotify, AttackLayer7SummaryVerticalParamsHTTPMethodOrderpatch, AttackLayer7SummaryVerticalParamsHTTPMethodPoll, AttackLayer7SummaryVerticalParamsHTTPMethodProppatch, AttackLayer7SummaryVerticalParamsHTTPMethodReport, AttackLayer7SummaryVerticalParamsHTTPMethodSearch, AttackLayer7SummaryVerticalParamsHTTPMethodSubscribe, AttackLayer7SummaryVerticalParamsHTTPMethodTrace, AttackLayer7SummaryVerticalParamsHTTPMethodUncheckout, AttackLayer7SummaryVerticalParamsHTTPMethodUnlock, AttackLayer7SummaryVerticalParamsHTTPMethodUnsubscribe, AttackLayer7SummaryVerticalParamsHTTPMethodUpdate, AttackLayer7SummaryVerticalParamsHTTPMethodVersioncontrol, AttackLayer7SummaryVerticalParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryVerticalParamsHTTPMethodXmsenumatts, AttackLayer7SummaryVerticalParamsHTTPMethodRpcOutData, AttackLayer7SummaryVerticalParamsHTTPMethodRpcInData, AttackLayer7SummaryVerticalParamsHTTPMethodJson, AttackLayer7SummaryVerticalParamsHTTPMethodCook, AttackLayer7SummaryVerticalParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalParamsHTTPVersion string

const (
	AttackLayer7SummaryVerticalParamsHTTPVersionHttPv1 AttackLayer7SummaryVerticalParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryVerticalParamsHTTPVersionHttPv2 AttackLayer7SummaryVerticalParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryVerticalParamsHTTPVersionHttPv3 AttackLayer7SummaryVerticalParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryVerticalParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalParamsHTTPVersionHttPv1, AttackLayer7SummaryVerticalParamsHTTPVersionHttPv2, AttackLayer7SummaryVerticalParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalParamsIPVersion string

const (
	AttackLayer7SummaryVerticalParamsIPVersionIPv4 AttackLayer7SummaryVerticalParamsIPVersion = "IPv4"
	AttackLayer7SummaryVerticalParamsIPVersionIPv6 AttackLayer7SummaryVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalParamsIPVersionIPv4, AttackLayer7SummaryVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalParamsMitigationProduct string

const (
	AttackLayer7SummaryVerticalParamsMitigationProductDDoS               AttackLayer7SummaryVerticalParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryVerticalParamsMitigationProductWAF                AttackLayer7SummaryVerticalParamsMitigationProduct = "WAF"
	AttackLayer7SummaryVerticalParamsMitigationProductBotManagement      AttackLayer7SummaryVerticalParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryVerticalParamsMitigationProductAccessRules        AttackLayer7SummaryVerticalParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryVerticalParamsMitigationProductIPReputation       AttackLayer7SummaryVerticalParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryVerticalParamsMitigationProductAPIShield          AttackLayer7SummaryVerticalParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryVerticalParamsMitigationProductDataLossPrevention AttackLayer7SummaryVerticalParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryVerticalParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryVerticalParamsMitigationProductDDoS, AttackLayer7SummaryVerticalParamsMitigationProductWAF, AttackLayer7SummaryVerticalParamsMitigationProductBotManagement, AttackLayer7SummaryVerticalParamsMitigationProductAccessRules, AttackLayer7SummaryVerticalParamsMitigationProductIPReputation, AttackLayer7SummaryVerticalParamsMitigationProductAPIShield, AttackLayer7SummaryVerticalParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryVerticalResponseEnvelope struct {
	Result  AttackLayer7SummaryVerticalResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer7SummaryVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryVerticalResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryVerticalResponseEnvelope]
type attackLayer7SummaryVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
