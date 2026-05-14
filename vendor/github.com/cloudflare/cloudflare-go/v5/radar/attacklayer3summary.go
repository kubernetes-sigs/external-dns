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

// AttackLayer3SummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3SummaryService] method instead.
type AttackLayer3SummaryService struct {
	Options []option.RequestOption
}

// NewAttackLayer3SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer3SummaryService(opts ...option.RequestOption) (r *AttackLayer3SummaryService) {
	r = &AttackLayer3SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 3 attacks by bitrate.
func (r *AttackLayer3SummaryService) Bitrate(ctx context.Context, query AttackLayer3SummaryBitrateParams, opts ...option.RequestOption) (res *AttackLayer3SummaryBitrateResponse, err error) {
	var env AttackLayer3SummaryBitrateResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by duration.
func (r *AttackLayer3SummaryService) Duration(ctx context.Context, query AttackLayer3SummaryDurationParams, opts ...option.RequestOption) (res *AttackLayer3SummaryDurationResponse, err error) {
	var env AttackLayer3SummaryDurationResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by targeted industry.
func (r *AttackLayer3SummaryService) Industry(ctx context.Context, query AttackLayer3SummaryIndustryParams, opts ...option.RequestOption) (res *AttackLayer3SummaryIndustryResponse, err error) {
	var env AttackLayer3SummaryIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by IP version.
func (r *AttackLayer3SummaryService) IPVersion(ctx context.Context, query AttackLayer3SummaryIPVersionParams, opts ...option.RequestOption) (res *AttackLayer3SummaryIPVersionResponse, err error) {
	var env AttackLayer3SummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by protocol.
func (r *AttackLayer3SummaryService) Protocol(ctx context.Context, query AttackLayer3SummaryProtocolParams, opts ...option.RequestOption) (res *AttackLayer3SummaryProtocolResponse, err error) {
	var env AttackLayer3SummaryProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by vector.
func (r *AttackLayer3SummaryService) Vector(ctx context.Context, query AttackLayer3SummaryVectorParams, opts ...option.RequestOption) (res *AttackLayer3SummaryVectorResponse, err error) {
	var env AttackLayer3SummaryVectorResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks by targeted vertical.
func (r *AttackLayer3SummaryService) Vertical(ctx context.Context, query AttackLayer3SummaryVerticalParams, opts ...option.RequestOption) (res *AttackLayer3SummaryVerticalResponse, err error) {
	var env AttackLayer3SummaryVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3SummaryBitrateResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryBitrateResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryBitrateResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryBitrateResponseJSON     `json:"-"`
}

// attackLayer3SummaryBitrateResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryBitrateResponse]
type attackLayer3SummaryBitrateResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryBitrateResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryBitrateResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryBitrateResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryBitrateResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryBitrateResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryBitrateResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryBitrateResponseMeta]
type attackLayer3SummaryBitrateResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryBitrateResponseMetaConfidenceInfo]
type attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryBitrateResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryBitrateResponseMetaDateRange]
type attackLayer3SummaryBitrateResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryBitrateResponseMetaNormalization string

const (
	AttackLayer3SummaryBitrateResponseMetaNormalizationPercentage           AttackLayer3SummaryBitrateResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryBitrateResponseMetaNormalizationMin0Max              AttackLayer3SummaryBitrateResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryBitrateResponseMetaNormalizationMinMax               AttackLayer3SummaryBitrateResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryBitrateResponseMetaNormalizationRawValues            AttackLayer3SummaryBitrateResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryBitrateResponseMetaNormalizationPercentageChange     AttackLayer3SummaryBitrateResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryBitrateResponseMetaNormalizationRollingAverage       AttackLayer3SummaryBitrateResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryBitrateResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryBitrateResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryBitrateResponseMetaNormalizationRatio                AttackLayer3SummaryBitrateResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryBitrateResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateResponseMetaNormalizationPercentage, AttackLayer3SummaryBitrateResponseMetaNormalizationMin0Max, AttackLayer3SummaryBitrateResponseMetaNormalizationMinMax, AttackLayer3SummaryBitrateResponseMetaNormalizationRawValues, AttackLayer3SummaryBitrateResponseMetaNormalizationPercentageChange, AttackLayer3SummaryBitrateResponseMetaNormalizationRollingAverage, AttackLayer3SummaryBitrateResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryBitrateResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  attackLayer3SummaryBitrateResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryBitrateResponseMetaUnit]
type attackLayer3SummaryBitrateResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateResponseSummary0 struct {
	// A numeric string.
	OneGBPSToTenGBPS string `json:"_1_GBPS_TO_10_GBPS,required"`
	// A numeric string.
	TenGBPSToOneHundredGBPS string `json:"_10_GBPS_TO_100_GBPS,required"`
	// A numeric string.
	FiveHundredMBPSToOneGBPS string `json:"_500_MBPS_TO_1_GBPS,required"`
	// A numeric string.
	Over100GBPS string `json:"OVER_100_GBPS,required"`
	// A numeric string.
	Under500MBPS string                                         `json:"UNDER_500_MBPS,required"`
	JSON         attackLayer3SummaryBitrateResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryBitrateResponseSummary0]
type attackLayer3SummaryBitrateResponseSummary0JSON struct {
	OneGBPSToTenGBPS         apijson.Field
	TenGBPSToOneHundredGBPS  apijson.Field
	FiveHundredMBPSToOneGBPS apijson.Field
	Over100GBPS              apijson.Field
	Under500MBPS             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryDurationResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryDurationResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryDurationResponseJSON     `json:"-"`
}

// attackLayer3SummaryDurationResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryDurationResponse]
type attackLayer3SummaryDurationResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryDurationResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryDurationResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryDurationResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryDurationResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryDurationResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryDurationResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryDurationResponseMeta]
type attackLayer3SummaryDurationResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryDurationResponseMetaConfidenceInfo]
type attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryDurationResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryDurationResponseMetaDateRange]
type attackLayer3SummaryDurationResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryDurationResponseMetaNormalization string

const (
	AttackLayer3SummaryDurationResponseMetaNormalizationPercentage           AttackLayer3SummaryDurationResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryDurationResponseMetaNormalizationMin0Max              AttackLayer3SummaryDurationResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryDurationResponseMetaNormalizationMinMax               AttackLayer3SummaryDurationResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryDurationResponseMetaNormalizationRawValues            AttackLayer3SummaryDurationResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryDurationResponseMetaNormalizationPercentageChange     AttackLayer3SummaryDurationResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryDurationResponseMetaNormalizationRollingAverage       AttackLayer3SummaryDurationResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryDurationResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryDurationResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryDurationResponseMetaNormalizationRatio                AttackLayer3SummaryDurationResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryDurationResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationResponseMetaNormalizationPercentage, AttackLayer3SummaryDurationResponseMetaNormalizationMin0Max, AttackLayer3SummaryDurationResponseMetaNormalizationMinMax, AttackLayer3SummaryDurationResponseMetaNormalizationRawValues, AttackLayer3SummaryDurationResponseMetaNormalizationPercentageChange, AttackLayer3SummaryDurationResponseMetaNormalizationRollingAverage, AttackLayer3SummaryDurationResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryDurationResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer3SummaryDurationResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryDurationResponseMetaUnit]
type attackLayer3SummaryDurationResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationResponseSummary0 struct {
	// A numeric string.
	OneHourToThreeHours string `json:"_1_HOUR_TO_3_HOURS,required"`
	// A numeric string.
	TenMinsToTwentyMins string `json:"_10_MINS_TO_20_MINS,required"`
	// A numeric string.
	TwentyMinsToFortyMins string `json:"_20_MINS_TO_40_MINS,required"`
	// A numeric string.
	FortyMinsToOneHour string `json:"_40_MINS_TO_1_HOUR,required"`
	// A numeric string.
	Over3Hours string `json:"OVER_3_HOURS,required"`
	// A numeric string.
	Under10Mins string                                          `json:"UNDER_10_MINS,required"`
	JSON        attackLayer3SummaryDurationResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryDurationResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryDurationResponseSummary0]
type attackLayer3SummaryDurationResponseSummary0JSON struct {
	OneHourToThreeHours   apijson.Field
	TenMinsToTwentyMins   apijson.Field
	TwentyMinsToFortyMins apijson.Field
	FortyMinsToOneHour    apijson.Field
	Over3Hours            apijson.Field
	Under10Mins           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIndustryResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryIndustryResponseMeta `json:"meta,required"`
	Summary0 map[string]string                       `json:"summary_0,required"`
	JSON     attackLayer3SummaryIndustryResponseJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIndustryResponse]
type attackLayer3SummaryIndustryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryIndustryResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIndustryResponseMeta]
type attackLayer3SummaryIndustryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer3SummaryIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryIndustryResponseMetaConfidenceInfo]
type attackLayer3SummaryIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryIndustryResponseMetaDateRange]
type attackLayer3SummaryIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryIndustryResponseMetaNormalization string

const (
	AttackLayer3SummaryIndustryResponseMetaNormalizationPercentage           AttackLayer3SummaryIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryIndustryResponseMetaNormalizationMin0Max              AttackLayer3SummaryIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryIndustryResponseMetaNormalizationMinMax               AttackLayer3SummaryIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryIndustryResponseMetaNormalizationRawValues            AttackLayer3SummaryIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryIndustryResponseMetaNormalizationPercentageChange     AttackLayer3SummaryIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryIndustryResponseMetaNormalizationRollingAverage       AttackLayer3SummaryIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryIndustryResponseMetaNormalizationRatio                AttackLayer3SummaryIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIndustryResponseMetaNormalizationPercentage, AttackLayer3SummaryIndustryResponseMetaNormalizationMin0Max, AttackLayer3SummaryIndustryResponseMetaNormalizationMinMax, AttackLayer3SummaryIndustryResponseMetaNormalizationRawValues, AttackLayer3SummaryIndustryResponseMetaNormalizationPercentageChange, AttackLayer3SummaryIndustryResponseMetaNormalizationRollingAverage, AttackLayer3SummaryIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryIndustryResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer3SummaryIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIndustryResponseMetaUnit]
type attackLayer3SummaryIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryIPVersionResponseJSON     `json:"-"`
}

// attackLayer3SummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIPVersionResponse]
type attackLayer3SummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryIPVersionResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryIPVersionResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryIPVersionResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryIPVersionResponseMeta]
type attackLayer3SummaryIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo]
type attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryIPVersionResponseMetaDateRange]
type attackLayer3SummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryIPVersionResponseMetaNormalization string

const (
	AttackLayer3SummaryIPVersionResponseMetaNormalizationPercentage           AttackLayer3SummaryIPVersionResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationMin0Max              AttackLayer3SummaryIPVersionResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationMinMax               AttackLayer3SummaryIPVersionResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationRawValues            AttackLayer3SummaryIPVersionResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationPercentageChange     AttackLayer3SummaryIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationRollingAverage       AttackLayer3SummaryIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryIPVersionResponseMetaNormalizationRatio                AttackLayer3SummaryIPVersionResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionResponseMetaNormalizationPercentage, AttackLayer3SummaryIPVersionResponseMetaNormalizationMin0Max, AttackLayer3SummaryIPVersionResponseMetaNormalizationMinMax, AttackLayer3SummaryIPVersionResponseMetaNormalizationRawValues, AttackLayer3SummaryIPVersionResponseMetaNormalizationPercentageChange, AttackLayer3SummaryIPVersionResponseMetaNormalizationRollingAverage, AttackLayer3SummaryIPVersionResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryIPVersionResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  attackLayer3SummaryIPVersionResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIPVersionResponseMetaUnit]
type attackLayer3SummaryIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionResponseSummary0 struct {
	// A numeric string.
	IPv4 string `json:"IPv4,required"`
	// A numeric string.
	IPv6 string                                           `json:"IPv6,required"`
	JSON attackLayer3SummaryIPVersionResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIPVersionResponseSummary0]
type attackLayer3SummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryProtocolResponseMeta     `json:"meta,required"`
	Summary0 AttackLayer3SummaryProtocolResponseSummary0 `json:"summary_0,required"`
	JSON     attackLayer3SummaryProtocolResponseJSON     `json:"-"`
}

// attackLayer3SummaryProtocolResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryProtocolResponse]
type attackLayer3SummaryProtocolResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryProtocolResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryProtocolResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryProtocolResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryProtocolResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryProtocolResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryProtocolResponseMeta]
type attackLayer3SummaryProtocolResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryProtocolResponseMetaConfidenceInfo]
type attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryProtocolResponseMetaDateRange]
type attackLayer3SummaryProtocolResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryProtocolResponseMetaNormalization string

const (
	AttackLayer3SummaryProtocolResponseMetaNormalizationPercentage           AttackLayer3SummaryProtocolResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryProtocolResponseMetaNormalizationMin0Max              AttackLayer3SummaryProtocolResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryProtocolResponseMetaNormalizationMinMax               AttackLayer3SummaryProtocolResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryProtocolResponseMetaNormalizationRawValues            AttackLayer3SummaryProtocolResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryProtocolResponseMetaNormalizationPercentageChange     AttackLayer3SummaryProtocolResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryProtocolResponseMetaNormalizationRollingAverage       AttackLayer3SummaryProtocolResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryProtocolResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryProtocolResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryProtocolResponseMetaNormalizationRatio                AttackLayer3SummaryProtocolResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryProtocolResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolResponseMetaNormalizationPercentage, AttackLayer3SummaryProtocolResponseMetaNormalizationMin0Max, AttackLayer3SummaryProtocolResponseMetaNormalizationMinMax, AttackLayer3SummaryProtocolResponseMetaNormalizationRawValues, AttackLayer3SummaryProtocolResponseMetaNormalizationPercentageChange, AttackLayer3SummaryProtocolResponseMetaNormalizationRollingAverage, AttackLayer3SummaryProtocolResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryProtocolResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryProtocolResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer3SummaryProtocolResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryProtocolResponseMetaUnit]
type attackLayer3SummaryProtocolResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolResponseSummary0 struct {
	// A numeric string.
	GRE string `json:"GRE,required"`
	// A numeric string.
	Icmp string `json:"ICMP,required"`
	// A numeric string.
	TCP string `json:"TCP,required"`
	// A numeric string.
	Udp  string                                          `json:"UDP,required"`
	JSON attackLayer3SummaryProtocolResponseSummary0JSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseSummary0JSON contains the JSON metadata for
// the struct [AttackLayer3SummaryProtocolResponseSummary0]
type attackLayer3SummaryProtocolResponseSummary0JSON struct {
	GRE         apijson.Field
	Icmp        apijson.Field
	TCP         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryVectorResponseMeta `json:"meta,required"`
	Summary0 map[string]string                     `json:"summary_0,required"`
	JSON     attackLayer3SummaryVectorResponseJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryVectorResponse]
type attackLayer3SummaryVectorResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryVectorResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryVectorResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryVectorResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryVectorResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryVectorResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryVectorResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVectorResponseMeta]
type attackLayer3SummaryVectorResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryVectorResponseMetaConfidenceInfo]
type attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryVectorResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryVectorResponseMetaDateRange]
type attackLayer3SummaryVectorResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryVectorResponseMetaNormalization string

const (
	AttackLayer3SummaryVectorResponseMetaNormalizationPercentage           AttackLayer3SummaryVectorResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryVectorResponseMetaNormalizationMin0Max              AttackLayer3SummaryVectorResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryVectorResponseMetaNormalizationMinMax               AttackLayer3SummaryVectorResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryVectorResponseMetaNormalizationRawValues            AttackLayer3SummaryVectorResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryVectorResponseMetaNormalizationPercentageChange     AttackLayer3SummaryVectorResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryVectorResponseMetaNormalizationRollingAverage       AttackLayer3SummaryVectorResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryVectorResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryVectorResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryVectorResponseMetaNormalizationRatio                AttackLayer3SummaryVectorResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryVectorResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorResponseMetaNormalizationPercentage, AttackLayer3SummaryVectorResponseMetaNormalizationMin0Max, AttackLayer3SummaryVectorResponseMetaNormalizationMinMax, AttackLayer3SummaryVectorResponseMetaNormalizationRawValues, AttackLayer3SummaryVectorResponseMetaNormalizationPercentageChange, AttackLayer3SummaryVectorResponseMetaNormalizationRollingAverage, AttackLayer3SummaryVectorResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryVectorResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  attackLayer3SummaryVectorResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVectorResponseMetaUnit]
type attackLayer3SummaryVectorResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVerticalResponse struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryVerticalResponseMeta `json:"meta,required"`
	Summary0 map[string]string                       `json:"summary_0,required"`
	JSON     attackLayer3SummaryVerticalResponseJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVerticalResponse]
type attackLayer3SummaryVerticalResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryVerticalResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVerticalResponseMeta]
type attackLayer3SummaryVerticalResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  attackLayer3SummaryVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer3SummaryVerticalResponseMetaConfidenceInfo]
type attackLayer3SummaryVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryVerticalResponseMetaDateRange]
type attackLayer3SummaryVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryVerticalResponseMetaNormalization string

const (
	AttackLayer3SummaryVerticalResponseMetaNormalizationPercentage           AttackLayer3SummaryVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryVerticalResponseMetaNormalizationMin0Max              AttackLayer3SummaryVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryVerticalResponseMetaNormalizationMinMax               AttackLayer3SummaryVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryVerticalResponseMetaNormalizationRawValues            AttackLayer3SummaryVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryVerticalResponseMetaNormalizationPercentageChange     AttackLayer3SummaryVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryVerticalResponseMetaNormalizationRollingAverage       AttackLayer3SummaryVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryVerticalResponseMetaNormalizationRatio                AttackLayer3SummaryVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVerticalResponseMetaNormalizationPercentage, AttackLayer3SummaryVerticalResponseMetaNormalizationMin0Max, AttackLayer3SummaryVerticalResponseMetaNormalizationMinMax, AttackLayer3SummaryVerticalResponseMetaNormalizationRawValues, AttackLayer3SummaryVerticalResponseMetaNormalizationPercentageChange, AttackLayer3SummaryVerticalResponseMetaNormalizationRollingAverage, AttackLayer3SummaryVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryVerticalResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  attackLayer3SummaryVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryVerticalResponseMetaUnit]
type attackLayer3SummaryVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryBitrateParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryBitrateParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryBitrateParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryBitrateParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryBitrateParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryBitrateParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryBitrateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryBitrateParamsDirection string

const (
	AttackLayer3SummaryBitrateParamsDirectionOrigin AttackLayer3SummaryBitrateParamsDirection = "ORIGIN"
	AttackLayer3SummaryBitrateParamsDirectionTarget AttackLayer3SummaryBitrateParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryBitrateParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsDirectionOrigin, AttackLayer3SummaryBitrateParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryBitrateParamsFormat string

const (
	AttackLayer3SummaryBitrateParamsFormatJson AttackLayer3SummaryBitrateParamsFormat = "JSON"
	AttackLayer3SummaryBitrateParamsFormatCsv  AttackLayer3SummaryBitrateParamsFormat = "CSV"
)

func (r AttackLayer3SummaryBitrateParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsFormatJson, AttackLayer3SummaryBitrateParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateParamsIPVersion string

const (
	AttackLayer3SummaryBitrateParamsIPVersionIPv4 AttackLayer3SummaryBitrateParamsIPVersion = "IPv4"
	AttackLayer3SummaryBitrateParamsIPVersionIPv6 AttackLayer3SummaryBitrateParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryBitrateParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsIPVersionIPv4, AttackLayer3SummaryBitrateParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateParamsProtocol string

const (
	AttackLayer3SummaryBitrateParamsProtocolUdp  AttackLayer3SummaryBitrateParamsProtocol = "UDP"
	AttackLayer3SummaryBitrateParamsProtocolTCP  AttackLayer3SummaryBitrateParamsProtocol = "TCP"
	AttackLayer3SummaryBitrateParamsProtocolIcmp AttackLayer3SummaryBitrateParamsProtocol = "ICMP"
	AttackLayer3SummaryBitrateParamsProtocolGRE  AttackLayer3SummaryBitrateParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryBitrateParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryBitrateParamsProtocolUdp, AttackLayer3SummaryBitrateParamsProtocolTCP, AttackLayer3SummaryBitrateParamsProtocolIcmp, AttackLayer3SummaryBitrateParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryBitrateResponseEnvelope struct {
	Result  AttackLayer3SummaryBitrateResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    attackLayer3SummaryBitrateResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryBitrateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryBitrateResponseEnvelope]
type attackLayer3SummaryBitrateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryBitrateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryBitrateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryDurationParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryDurationParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryDurationParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryDurationParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryDurationParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryDurationParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryDurationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryDurationParamsDirection string

const (
	AttackLayer3SummaryDurationParamsDirectionOrigin AttackLayer3SummaryDurationParamsDirection = "ORIGIN"
	AttackLayer3SummaryDurationParamsDirectionTarget AttackLayer3SummaryDurationParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryDurationParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsDirectionOrigin, AttackLayer3SummaryDurationParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryDurationParamsFormat string

const (
	AttackLayer3SummaryDurationParamsFormatJson AttackLayer3SummaryDurationParamsFormat = "JSON"
	AttackLayer3SummaryDurationParamsFormatCsv  AttackLayer3SummaryDurationParamsFormat = "CSV"
)

func (r AttackLayer3SummaryDurationParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsFormatJson, AttackLayer3SummaryDurationParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationParamsIPVersion string

const (
	AttackLayer3SummaryDurationParamsIPVersionIPv4 AttackLayer3SummaryDurationParamsIPVersion = "IPv4"
	AttackLayer3SummaryDurationParamsIPVersionIPv6 AttackLayer3SummaryDurationParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryDurationParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsIPVersionIPv4, AttackLayer3SummaryDurationParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationParamsProtocol string

const (
	AttackLayer3SummaryDurationParamsProtocolUdp  AttackLayer3SummaryDurationParamsProtocol = "UDP"
	AttackLayer3SummaryDurationParamsProtocolTCP  AttackLayer3SummaryDurationParamsProtocol = "TCP"
	AttackLayer3SummaryDurationParamsProtocolIcmp AttackLayer3SummaryDurationParamsProtocol = "ICMP"
	AttackLayer3SummaryDurationParamsProtocolGRE  AttackLayer3SummaryDurationParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryDurationParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryDurationParamsProtocolUdp, AttackLayer3SummaryDurationParamsProtocolTCP, AttackLayer3SummaryDurationParamsProtocolIcmp, AttackLayer3SummaryDurationParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryDurationResponseEnvelope struct {
	Result  AttackLayer3SummaryDurationResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryDurationResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryDurationResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryDurationResponseEnvelope]
type attackLayer3SummaryDurationResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryDurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryDurationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIndustryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryIndustryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryIndustryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryIndustryParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryIndustryParamsDirection string

const (
	AttackLayer3SummaryIndustryParamsDirectionOrigin AttackLayer3SummaryIndustryParamsDirection = "ORIGIN"
	AttackLayer3SummaryIndustryParamsDirectionTarget AttackLayer3SummaryIndustryParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryIndustryParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIndustryParamsDirectionOrigin, AttackLayer3SummaryIndustryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryIndustryParamsFormat string

const (
	AttackLayer3SummaryIndustryParamsFormatJson AttackLayer3SummaryIndustryParamsFormat = "JSON"
	AttackLayer3SummaryIndustryParamsFormatCsv  AttackLayer3SummaryIndustryParamsFormat = "CSV"
)

func (r AttackLayer3SummaryIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIndustryParamsFormatJson, AttackLayer3SummaryIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryIndustryParamsIPVersion string

const (
	AttackLayer3SummaryIndustryParamsIPVersionIPv4 AttackLayer3SummaryIndustryParamsIPVersion = "IPv4"
	AttackLayer3SummaryIndustryParamsIPVersionIPv6 AttackLayer3SummaryIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIndustryParamsIPVersionIPv4, AttackLayer3SummaryIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryIndustryParamsProtocol string

const (
	AttackLayer3SummaryIndustryParamsProtocolUdp  AttackLayer3SummaryIndustryParamsProtocol = "UDP"
	AttackLayer3SummaryIndustryParamsProtocolTCP  AttackLayer3SummaryIndustryParamsProtocol = "TCP"
	AttackLayer3SummaryIndustryParamsProtocolIcmp AttackLayer3SummaryIndustryParamsProtocol = "ICMP"
	AttackLayer3SummaryIndustryParamsProtocolGRE  AttackLayer3SummaryIndustryParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryIndustryParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIndustryParamsProtocolUdp, AttackLayer3SummaryIndustryParamsProtocolTCP, AttackLayer3SummaryIndustryParamsProtocolIcmp, AttackLayer3SummaryIndustryParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryIndustryResponseEnvelope struct {
	Result  AttackLayer3SummaryIndustryResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryIndustryResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIndustryResponseEnvelope]
type attackLayer3SummaryIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryIPVersionParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryIPVersionParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryIPVersionParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryIPVersionParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryIPVersionParamsDirection string

const (
	AttackLayer3SummaryIPVersionParamsDirectionOrigin AttackLayer3SummaryIPVersionParamsDirection = "ORIGIN"
	AttackLayer3SummaryIPVersionParamsDirectionTarget AttackLayer3SummaryIPVersionParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryIPVersionParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsDirectionOrigin, AttackLayer3SummaryIPVersionParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryIPVersionParamsFormat string

const (
	AttackLayer3SummaryIPVersionParamsFormatJson AttackLayer3SummaryIPVersionParamsFormat = "JSON"
	AttackLayer3SummaryIPVersionParamsFormatCsv  AttackLayer3SummaryIPVersionParamsFormat = "CSV"
)

func (r AttackLayer3SummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsFormatJson, AttackLayer3SummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryIPVersionParamsProtocol string

const (
	AttackLayer3SummaryIPVersionParamsProtocolUdp  AttackLayer3SummaryIPVersionParamsProtocol = "UDP"
	AttackLayer3SummaryIPVersionParamsProtocolTCP  AttackLayer3SummaryIPVersionParamsProtocol = "TCP"
	AttackLayer3SummaryIPVersionParamsProtocolIcmp AttackLayer3SummaryIPVersionParamsProtocol = "ICMP"
	AttackLayer3SummaryIPVersionParamsProtocolGRE  AttackLayer3SummaryIPVersionParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryIPVersionParamsProtocolUdp, AttackLayer3SummaryIPVersionParamsProtocolTCP, AttackLayer3SummaryIPVersionParamsProtocolIcmp, AttackLayer3SummaryIPVersionParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryIPVersionResponseEnvelope struct {
	Result  AttackLayer3SummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    attackLayer3SummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryIPVersionResponseEnvelope]
type attackLayer3SummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryProtocolParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryProtocolParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryProtocolParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryProtocolParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer3SummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryProtocolParamsDirection string

const (
	AttackLayer3SummaryProtocolParamsDirectionOrigin AttackLayer3SummaryProtocolParamsDirection = "ORIGIN"
	AttackLayer3SummaryProtocolParamsDirectionTarget AttackLayer3SummaryProtocolParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryProtocolParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsDirectionOrigin, AttackLayer3SummaryProtocolParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryProtocolParamsFormat string

const (
	AttackLayer3SummaryProtocolParamsFormatJson AttackLayer3SummaryProtocolParamsFormat = "JSON"
	AttackLayer3SummaryProtocolParamsFormatCsv  AttackLayer3SummaryProtocolParamsFormat = "CSV"
)

func (r AttackLayer3SummaryProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsFormatJson, AttackLayer3SummaryProtocolParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryProtocolParamsIPVersion string

const (
	AttackLayer3SummaryProtocolParamsIPVersionIPv4 AttackLayer3SummaryProtocolParamsIPVersion = "IPv4"
	AttackLayer3SummaryProtocolParamsIPVersionIPv6 AttackLayer3SummaryProtocolParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryProtocolParamsIPVersionIPv4, AttackLayer3SummaryProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryProtocolResponseEnvelope struct {
	Result  AttackLayer3SummaryProtocolResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryProtocolResponseEnvelope]
type attackLayer3SummaryProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVectorParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryVectorParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryVectorParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryVectorParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryVectorParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryVectorParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryVectorParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryVectorParamsDirection string

const (
	AttackLayer3SummaryVectorParamsDirectionOrigin AttackLayer3SummaryVectorParamsDirection = "ORIGIN"
	AttackLayer3SummaryVectorParamsDirectionTarget AttackLayer3SummaryVectorParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryVectorParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsDirectionOrigin, AttackLayer3SummaryVectorParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryVectorParamsFormat string

const (
	AttackLayer3SummaryVectorParamsFormatJson AttackLayer3SummaryVectorParamsFormat = "JSON"
	AttackLayer3SummaryVectorParamsFormatCsv  AttackLayer3SummaryVectorParamsFormat = "CSV"
)

func (r AttackLayer3SummaryVectorParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsFormatJson, AttackLayer3SummaryVectorParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorParamsIPVersion string

const (
	AttackLayer3SummaryVectorParamsIPVersionIPv4 AttackLayer3SummaryVectorParamsIPVersion = "IPv4"
	AttackLayer3SummaryVectorParamsIPVersionIPv6 AttackLayer3SummaryVectorParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryVectorParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsIPVersionIPv4, AttackLayer3SummaryVectorParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorParamsProtocol string

const (
	AttackLayer3SummaryVectorParamsProtocolUdp  AttackLayer3SummaryVectorParamsProtocol = "UDP"
	AttackLayer3SummaryVectorParamsProtocolTCP  AttackLayer3SummaryVectorParamsProtocol = "TCP"
	AttackLayer3SummaryVectorParamsProtocolIcmp AttackLayer3SummaryVectorParamsProtocol = "ICMP"
	AttackLayer3SummaryVectorParamsProtocolGRE  AttackLayer3SummaryVectorParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryVectorParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVectorParamsProtocolUdp, AttackLayer3SummaryVectorParamsProtocolTCP, AttackLayer3SummaryVectorParamsProtocolIcmp, AttackLayer3SummaryVectorParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryVectorResponseEnvelope struct {
	Result  AttackLayer3SummaryVectorResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    attackLayer3SummaryVectorResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryVectorResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryVectorResponseEnvelope]
type attackLayer3SummaryVectorResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVectorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVectorResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryVerticalParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryVerticalParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryVerticalParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryVerticalParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryVerticalParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryVerticalParamsDirection string

const (
	AttackLayer3SummaryVerticalParamsDirectionOrigin AttackLayer3SummaryVerticalParamsDirection = "ORIGIN"
	AttackLayer3SummaryVerticalParamsDirectionTarget AttackLayer3SummaryVerticalParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryVerticalParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVerticalParamsDirectionOrigin, AttackLayer3SummaryVerticalParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryVerticalParamsFormat string

const (
	AttackLayer3SummaryVerticalParamsFormatJson AttackLayer3SummaryVerticalParamsFormat = "JSON"
	AttackLayer3SummaryVerticalParamsFormatCsv  AttackLayer3SummaryVerticalParamsFormat = "CSV"
)

func (r AttackLayer3SummaryVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVerticalParamsFormatJson, AttackLayer3SummaryVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryVerticalParamsIPVersion string

const (
	AttackLayer3SummaryVerticalParamsIPVersionIPv4 AttackLayer3SummaryVerticalParamsIPVersion = "IPv4"
	AttackLayer3SummaryVerticalParamsIPVersionIPv6 AttackLayer3SummaryVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVerticalParamsIPVersionIPv4, AttackLayer3SummaryVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryVerticalParamsProtocol string

const (
	AttackLayer3SummaryVerticalParamsProtocolUdp  AttackLayer3SummaryVerticalParamsProtocol = "UDP"
	AttackLayer3SummaryVerticalParamsProtocolTCP  AttackLayer3SummaryVerticalParamsProtocol = "TCP"
	AttackLayer3SummaryVerticalParamsProtocolIcmp AttackLayer3SummaryVerticalParamsProtocol = "ICMP"
	AttackLayer3SummaryVerticalParamsProtocolGRE  AttackLayer3SummaryVerticalParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryVerticalParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryVerticalParamsProtocolUdp, AttackLayer3SummaryVerticalParamsProtocolTCP, AttackLayer3SummaryVerticalParamsProtocolIcmp, AttackLayer3SummaryVerticalParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryVerticalResponseEnvelope struct {
	Result  AttackLayer3SummaryVerticalResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    attackLayer3SummaryVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryVerticalResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryVerticalResponseEnvelope]
type attackLayer3SummaryVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
