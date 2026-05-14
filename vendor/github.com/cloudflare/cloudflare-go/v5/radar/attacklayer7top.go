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

// AttackLayer7TopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TopService] method instead.
type AttackLayer7TopService struct {
	Options   []option.RequestOption
	Locations *AttackLayer7TopLocationService
	Ases      *AttackLayer7TopAseService
}

// NewAttackLayer7TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer7TopService(opts ...option.RequestOption) (r *AttackLayer7TopService) {
	r = &AttackLayer7TopService{}
	r.Options = opts
	r.Locations = NewAttackLayer7TopLocationService(opts...)
	r.Ases = NewAttackLayer7TopAseService(opts...)
	return
}

// Retrieves the top attacks from origin to target location. Values are percentages
// of the total layer 7 attacks (with billing country). The attack magnitude can be
// defined by the number of mitigated requests or by the number of zones affected.
// You can optionally limit the number of attacks by origin/target location (useful
// if all the top attacks are from or to the same location).
func (r *AttackLayer7TopService) Attacks(ctx context.Context, query AttackLayer7TopAttacksParams, opts ...option.RequestOption) (res *AttackLayer7TopAttacksResponse, err error) {
	var env AttackLayer7TopAttacksResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by industry endpoint.
//
// Deprecated: deprecated
func (r *AttackLayer7TopService) Industry(ctx context.Context, query AttackLayer7TopIndustryParams, opts ...option.RequestOption) (res *AttackLayer7TopIndustryResponse, err error) {
	var env AttackLayer7TopIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by vertical endpoint.
//
// Deprecated: deprecated
func (r *AttackLayer7TopService) Vertical(ctx context.Context, query AttackLayer7TopVerticalParams, opts ...option.RequestOption) (res *AttackLayer7TopVerticalResponse, err error) {
	var env AttackLayer7TopVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopAttacksResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopAttacksResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopAttacksResponseJSON   `json:"-"`
}

// attackLayer7TopAttacksResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponse]
type attackLayer7TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopAttacksResponseMeta struct {
	ConfidenceInfo AttackLayer7TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopAttacksResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopAttacksResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopAttacksResponseMetaJSON   `json:"-"`
}

// attackLayer7TopAttacksResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponseMeta]
type attackLayer7TopAttacksResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  attackLayer7TopAttacksResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopAttacksResponseMetaConfidenceInfo]
type attackLayer7TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopAttacksResponseMetaDateRange]
type attackLayer7TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopAttacksResponseMetaNormalization string

const (
	AttackLayer7TopAttacksResponseMetaNormalizationPercentage           AttackLayer7TopAttacksResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopAttacksResponseMetaNormalizationMin0Max              AttackLayer7TopAttacksResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopAttacksResponseMetaNormalizationMinMax               AttackLayer7TopAttacksResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopAttacksResponseMetaNormalizationRawValues            AttackLayer7TopAttacksResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopAttacksResponseMetaNormalizationPercentageChange     AttackLayer7TopAttacksResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopAttacksResponseMetaNormalizationRollingAverage       AttackLayer7TopAttacksResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopAttacksResponseMetaNormalizationOverlappedPercentage AttackLayer7TopAttacksResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopAttacksResponseMetaNormalizationRatio                AttackLayer7TopAttacksResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopAttacksResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksResponseMetaNormalizationPercentage, AttackLayer7TopAttacksResponseMetaNormalizationMin0Max, AttackLayer7TopAttacksResponseMetaNormalizationMinMax, AttackLayer7TopAttacksResponseMetaNormalizationRawValues, AttackLayer7TopAttacksResponseMetaNormalizationPercentageChange, AttackLayer7TopAttacksResponseMetaNormalizationRollingAverage, AttackLayer7TopAttacksResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopAttacksResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopAttacksResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  attackLayer7TopAttacksResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopAttacksResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7TopAttacksResponseMetaUnit]
type attackLayer7TopAttacksResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                 `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                 `json:"originCountryName,required"`
	TargetCountryAlpha2 string                                 `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                 `json:"targetCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                attackLayer7TopAttacksResponseTop0JSON `json:"-"`
}

// attackLayer7TopAttacksResponseTop0JSON contains the JSON metadata for the struct
// [AttackLayer7TopAttacksResponseTop0]
type attackLayer7TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopIndustryResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopIndustryResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopIndustryResponseJSON   `json:"-"`
}

// attackLayer7TopIndustryResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopIndustryResponse]
type attackLayer7TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopIndustryResponseMeta struct {
	ConfidenceInfo AttackLayer7TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer7TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseMeta]
type attackLayer7TopIndustryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  attackLayer7TopIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopIndustryResponseMetaConfidenceInfo]
type attackLayer7TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopIndustryResponseMetaDateRange]
type attackLayer7TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopIndustryResponseMetaNormalization string

const (
	AttackLayer7TopIndustryResponseMetaNormalizationPercentage           AttackLayer7TopIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopIndustryResponseMetaNormalizationMin0Max              AttackLayer7TopIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopIndustryResponseMetaNormalizationMinMax               AttackLayer7TopIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopIndustryResponseMetaNormalizationRawValues            AttackLayer7TopIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopIndustryResponseMetaNormalizationPercentageChange     AttackLayer7TopIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopIndustryResponseMetaNormalizationRollingAverage       AttackLayer7TopIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer7TopIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopIndustryResponseMetaNormalizationRatio                AttackLayer7TopIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryResponseMetaNormalizationPercentage, AttackLayer7TopIndustryResponseMetaNormalizationMin0Max, AttackLayer7TopIndustryResponseMetaNormalizationMinMax, AttackLayer7TopIndustryResponseMetaNormalizationRawValues, AttackLayer7TopIndustryResponseMetaNormalizationPercentageChange, AttackLayer7TopIndustryResponseMetaNormalizationRollingAverage, AttackLayer7TopIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopIndustryResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  attackLayer7TopIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopIndustryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseMetaUnit]
type attackLayer7TopIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer7TopIndustryResponseTop0JSON `json:"-"`
}

// attackLayer7TopIndustryResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseTop0]
type attackLayer7TopIndustryResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopVerticalResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopVerticalResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopVerticalResponseJSON   `json:"-"`
}

// attackLayer7TopVerticalResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopVerticalResponse]
type attackLayer7TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopVerticalResponseMeta struct {
	ConfidenceInfo AttackLayer7TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer7TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseMeta]
type attackLayer7TopVerticalResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  attackLayer7TopVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TopVerticalResponseMetaConfidenceInfo]
type attackLayer7TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopVerticalResponseMetaDateRange]
type attackLayer7TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopVerticalResponseMetaNormalization string

const (
	AttackLayer7TopVerticalResponseMetaNormalizationPercentage           AttackLayer7TopVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopVerticalResponseMetaNormalizationMin0Max              AttackLayer7TopVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopVerticalResponseMetaNormalizationMinMax               AttackLayer7TopVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopVerticalResponseMetaNormalizationRawValues            AttackLayer7TopVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopVerticalResponseMetaNormalizationPercentageChange     AttackLayer7TopVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopVerticalResponseMetaNormalizationRollingAverage       AttackLayer7TopVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer7TopVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopVerticalResponseMetaNormalizationRatio                AttackLayer7TopVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalResponseMetaNormalizationPercentage, AttackLayer7TopVerticalResponseMetaNormalizationMin0Max, AttackLayer7TopVerticalResponseMetaNormalizationMinMax, AttackLayer7TopVerticalResponseMetaNormalizationRawValues, AttackLayer7TopVerticalResponseMetaNormalizationPercentageChange, AttackLayer7TopVerticalResponseMetaNormalizationRollingAverage, AttackLayer7TopVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopVerticalResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  attackLayer7TopVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopVerticalResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseMetaUnit]
type attackLayer7TopVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer7TopVerticalResponseTop0JSON `json:"-"`
}

// attackLayer7TopVerticalResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseTop0]
type attackLayer7TopVerticalResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAttacksParams struct {
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
	Format param.Field[AttackLayer7TopAttacksParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Specifies whether the `limitPerLocation` applies to the source or target
	// location.
	LimitDirection param.Field[AttackLayer7TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limits the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopAttacksParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TopAttacksParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TopAttacksParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopAttacksParamsFormat string

const (
	AttackLayer7TopAttacksParamsFormatJson AttackLayer7TopAttacksParamsFormat = "JSON"
	AttackLayer7TopAttacksParamsFormatCsv  AttackLayer7TopAttacksParamsFormat = "CSV"
)

func (r AttackLayer7TopAttacksParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsFormatJson, AttackLayer7TopAttacksParamsFormatCsv:
		return true
	}
	return false
}

// Specifies whether the `limitPerLocation` applies to the source or target
// location.
type AttackLayer7TopAttacksParamsLimitDirection string

const (
	AttackLayer7TopAttacksParamsLimitDirectionOrigin AttackLayer7TopAttacksParamsLimitDirection = "ORIGIN"
	AttackLayer7TopAttacksParamsLimitDirectionTarget AttackLayer7TopAttacksParamsLimitDirection = "TARGET"
)

func (r AttackLayer7TopAttacksParamsLimitDirection) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsLimitDirectionOrigin, AttackLayer7TopAttacksParamsLimitDirectionTarget:
		return true
	}
	return false
}

type AttackLayer7TopAttacksParamsMitigationProduct string

const (
	AttackLayer7TopAttacksParamsMitigationProductDDoS               AttackLayer7TopAttacksParamsMitigationProduct = "DDOS"
	AttackLayer7TopAttacksParamsMitigationProductWAF                AttackLayer7TopAttacksParamsMitigationProduct = "WAF"
	AttackLayer7TopAttacksParamsMitigationProductBotManagement      AttackLayer7TopAttacksParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopAttacksParamsMitigationProductAccessRules        AttackLayer7TopAttacksParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopAttacksParamsMitigationProductIPReputation       AttackLayer7TopAttacksParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopAttacksParamsMitigationProductAPIShield          AttackLayer7TopAttacksParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopAttacksParamsMitigationProductDataLossPrevention AttackLayer7TopAttacksParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopAttacksParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsMitigationProductDDoS, AttackLayer7TopAttacksParamsMitigationProductWAF, AttackLayer7TopAttacksParamsMitigationProductBotManagement, AttackLayer7TopAttacksParamsMitigationProductAccessRules, AttackLayer7TopAttacksParamsMitigationProductIPReputation, AttackLayer7TopAttacksParamsMitigationProductAPIShield, AttackLayer7TopAttacksParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopAttacksParamsNormalization string

const (
	AttackLayer7TopAttacksParamsNormalizationPercentage AttackLayer7TopAttacksParamsNormalization = "PERCENTAGE"
	AttackLayer7TopAttacksParamsNormalizationMinMax     AttackLayer7TopAttacksParamsNormalization = "MIN_MAX"
)

func (r AttackLayer7TopAttacksParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopAttacksParamsNormalizationPercentage, AttackLayer7TopAttacksParamsNormalizationMinMax:
		return true
	}
	return false
}

type AttackLayer7TopAttacksResponseEnvelope struct {
	Result  AttackLayer7TopAttacksResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer7TopAttacksResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopAttacksResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopAttacksResponseEnvelope]
type attackLayer7TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAttacksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopIndustryParams struct {
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
	Format param.Field[AttackLayer7TopIndustryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TopIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TopIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopIndustryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopIndustryParamsFormat string

const (
	AttackLayer7TopIndustryParamsFormatJson AttackLayer7TopIndustryParamsFormat = "JSON"
	AttackLayer7TopIndustryParamsFormatCsv  AttackLayer7TopIndustryParamsFormat = "CSV"
)

func (r AttackLayer7TopIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsFormatJson, AttackLayer7TopIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopIndustryParamsHTTPMethod string

const (
	AttackLayer7TopIndustryParamsHTTPMethodGet             AttackLayer7TopIndustryParamsHTTPMethod = "GET"
	AttackLayer7TopIndustryParamsHTTPMethodPost            AttackLayer7TopIndustryParamsHTTPMethod = "POST"
	AttackLayer7TopIndustryParamsHTTPMethodDelete          AttackLayer7TopIndustryParamsHTTPMethod = "DELETE"
	AttackLayer7TopIndustryParamsHTTPMethodPut             AttackLayer7TopIndustryParamsHTTPMethod = "PUT"
	AttackLayer7TopIndustryParamsHTTPMethodHead            AttackLayer7TopIndustryParamsHTTPMethod = "HEAD"
	AttackLayer7TopIndustryParamsHTTPMethodPurge           AttackLayer7TopIndustryParamsHTTPMethod = "PURGE"
	AttackLayer7TopIndustryParamsHTTPMethodOptions         AttackLayer7TopIndustryParamsHTTPMethod = "OPTIONS"
	AttackLayer7TopIndustryParamsHTTPMethodPropfind        AttackLayer7TopIndustryParamsHTTPMethod = "PROPFIND"
	AttackLayer7TopIndustryParamsHTTPMethodMkcol           AttackLayer7TopIndustryParamsHTTPMethod = "MKCOL"
	AttackLayer7TopIndustryParamsHTTPMethodPatch           AttackLayer7TopIndustryParamsHTTPMethod = "PATCH"
	AttackLayer7TopIndustryParamsHTTPMethodACL             AttackLayer7TopIndustryParamsHTTPMethod = "ACL"
	AttackLayer7TopIndustryParamsHTTPMethodBcopy           AttackLayer7TopIndustryParamsHTTPMethod = "BCOPY"
	AttackLayer7TopIndustryParamsHTTPMethodBdelete         AttackLayer7TopIndustryParamsHTTPMethod = "BDELETE"
	AttackLayer7TopIndustryParamsHTTPMethodBmove           AttackLayer7TopIndustryParamsHTTPMethod = "BMOVE"
	AttackLayer7TopIndustryParamsHTTPMethodBpropfind       AttackLayer7TopIndustryParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TopIndustryParamsHTTPMethodBproppatch      AttackLayer7TopIndustryParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TopIndustryParamsHTTPMethodCheckin         AttackLayer7TopIndustryParamsHTTPMethod = "CHECKIN"
	AttackLayer7TopIndustryParamsHTTPMethodCheckout        AttackLayer7TopIndustryParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TopIndustryParamsHTTPMethodConnect         AttackLayer7TopIndustryParamsHTTPMethod = "CONNECT"
	AttackLayer7TopIndustryParamsHTTPMethodCopy            AttackLayer7TopIndustryParamsHTTPMethod = "COPY"
	AttackLayer7TopIndustryParamsHTTPMethodLabel           AttackLayer7TopIndustryParamsHTTPMethod = "LABEL"
	AttackLayer7TopIndustryParamsHTTPMethodLock            AttackLayer7TopIndustryParamsHTTPMethod = "LOCK"
	AttackLayer7TopIndustryParamsHTTPMethodMerge           AttackLayer7TopIndustryParamsHTTPMethod = "MERGE"
	AttackLayer7TopIndustryParamsHTTPMethodMkactivity      AttackLayer7TopIndustryParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TopIndustryParamsHTTPMethodMkworkspace     AttackLayer7TopIndustryParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TopIndustryParamsHTTPMethodMove            AttackLayer7TopIndustryParamsHTTPMethod = "MOVE"
	AttackLayer7TopIndustryParamsHTTPMethodNotify          AttackLayer7TopIndustryParamsHTTPMethod = "NOTIFY"
	AttackLayer7TopIndustryParamsHTTPMethodOrderpatch      AttackLayer7TopIndustryParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TopIndustryParamsHTTPMethodPoll            AttackLayer7TopIndustryParamsHTTPMethod = "POLL"
	AttackLayer7TopIndustryParamsHTTPMethodProppatch       AttackLayer7TopIndustryParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TopIndustryParamsHTTPMethodReport          AttackLayer7TopIndustryParamsHTTPMethod = "REPORT"
	AttackLayer7TopIndustryParamsHTTPMethodSearch          AttackLayer7TopIndustryParamsHTTPMethod = "SEARCH"
	AttackLayer7TopIndustryParamsHTTPMethodSubscribe       AttackLayer7TopIndustryParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TopIndustryParamsHTTPMethodTrace           AttackLayer7TopIndustryParamsHTTPMethod = "TRACE"
	AttackLayer7TopIndustryParamsHTTPMethodUncheckout      AttackLayer7TopIndustryParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TopIndustryParamsHTTPMethodUnlock          AttackLayer7TopIndustryParamsHTTPMethod = "UNLOCK"
	AttackLayer7TopIndustryParamsHTTPMethodUnsubscribe     AttackLayer7TopIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TopIndustryParamsHTTPMethodUpdate          AttackLayer7TopIndustryParamsHTTPMethod = "UPDATE"
	AttackLayer7TopIndustryParamsHTTPMethodVersioncontrol  AttackLayer7TopIndustryParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TopIndustryParamsHTTPMethodBaselinecontrol AttackLayer7TopIndustryParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TopIndustryParamsHTTPMethodXmsenumatts     AttackLayer7TopIndustryParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TopIndustryParamsHTTPMethodRpcOutData      AttackLayer7TopIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TopIndustryParamsHTTPMethodRpcInData       AttackLayer7TopIndustryParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TopIndustryParamsHTTPMethodJson            AttackLayer7TopIndustryParamsHTTPMethod = "JSON"
	AttackLayer7TopIndustryParamsHTTPMethodCook            AttackLayer7TopIndustryParamsHTTPMethod = "COOK"
	AttackLayer7TopIndustryParamsHTTPMethodTrack           AttackLayer7TopIndustryParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TopIndustryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsHTTPMethodGet, AttackLayer7TopIndustryParamsHTTPMethodPost, AttackLayer7TopIndustryParamsHTTPMethodDelete, AttackLayer7TopIndustryParamsHTTPMethodPut, AttackLayer7TopIndustryParamsHTTPMethodHead, AttackLayer7TopIndustryParamsHTTPMethodPurge, AttackLayer7TopIndustryParamsHTTPMethodOptions, AttackLayer7TopIndustryParamsHTTPMethodPropfind, AttackLayer7TopIndustryParamsHTTPMethodMkcol, AttackLayer7TopIndustryParamsHTTPMethodPatch, AttackLayer7TopIndustryParamsHTTPMethodACL, AttackLayer7TopIndustryParamsHTTPMethodBcopy, AttackLayer7TopIndustryParamsHTTPMethodBdelete, AttackLayer7TopIndustryParamsHTTPMethodBmove, AttackLayer7TopIndustryParamsHTTPMethodBpropfind, AttackLayer7TopIndustryParamsHTTPMethodBproppatch, AttackLayer7TopIndustryParamsHTTPMethodCheckin, AttackLayer7TopIndustryParamsHTTPMethodCheckout, AttackLayer7TopIndustryParamsHTTPMethodConnect, AttackLayer7TopIndustryParamsHTTPMethodCopy, AttackLayer7TopIndustryParamsHTTPMethodLabel, AttackLayer7TopIndustryParamsHTTPMethodLock, AttackLayer7TopIndustryParamsHTTPMethodMerge, AttackLayer7TopIndustryParamsHTTPMethodMkactivity, AttackLayer7TopIndustryParamsHTTPMethodMkworkspace, AttackLayer7TopIndustryParamsHTTPMethodMove, AttackLayer7TopIndustryParamsHTTPMethodNotify, AttackLayer7TopIndustryParamsHTTPMethodOrderpatch, AttackLayer7TopIndustryParamsHTTPMethodPoll, AttackLayer7TopIndustryParamsHTTPMethodProppatch, AttackLayer7TopIndustryParamsHTTPMethodReport, AttackLayer7TopIndustryParamsHTTPMethodSearch, AttackLayer7TopIndustryParamsHTTPMethodSubscribe, AttackLayer7TopIndustryParamsHTTPMethodTrace, AttackLayer7TopIndustryParamsHTTPMethodUncheckout, AttackLayer7TopIndustryParamsHTTPMethodUnlock, AttackLayer7TopIndustryParamsHTTPMethodUnsubscribe, AttackLayer7TopIndustryParamsHTTPMethodUpdate, AttackLayer7TopIndustryParamsHTTPMethodVersioncontrol, AttackLayer7TopIndustryParamsHTTPMethodBaselinecontrol, AttackLayer7TopIndustryParamsHTTPMethodXmsenumatts, AttackLayer7TopIndustryParamsHTTPMethodRpcOutData, AttackLayer7TopIndustryParamsHTTPMethodRpcInData, AttackLayer7TopIndustryParamsHTTPMethodJson, AttackLayer7TopIndustryParamsHTTPMethodCook, AttackLayer7TopIndustryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TopIndustryParamsHTTPVersion string

const (
	AttackLayer7TopIndustryParamsHTTPVersionHttPv1 AttackLayer7TopIndustryParamsHTTPVersion = "HTTPv1"
	AttackLayer7TopIndustryParamsHTTPVersionHttPv2 AttackLayer7TopIndustryParamsHTTPVersion = "HTTPv2"
	AttackLayer7TopIndustryParamsHTTPVersionHttPv3 AttackLayer7TopIndustryParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TopIndustryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsHTTPVersionHttPv1, AttackLayer7TopIndustryParamsHTTPVersionHttPv2, AttackLayer7TopIndustryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TopIndustryParamsIPVersion string

const (
	AttackLayer7TopIndustryParamsIPVersionIPv4 AttackLayer7TopIndustryParamsIPVersion = "IPv4"
	AttackLayer7TopIndustryParamsIPVersionIPv6 AttackLayer7TopIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer7TopIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsIPVersionIPv4, AttackLayer7TopIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TopIndustryParamsMitigationProduct string

const (
	AttackLayer7TopIndustryParamsMitigationProductDDoS               AttackLayer7TopIndustryParamsMitigationProduct = "DDOS"
	AttackLayer7TopIndustryParamsMitigationProductWAF                AttackLayer7TopIndustryParamsMitigationProduct = "WAF"
	AttackLayer7TopIndustryParamsMitigationProductBotManagement      AttackLayer7TopIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopIndustryParamsMitigationProductAccessRules        AttackLayer7TopIndustryParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopIndustryParamsMitigationProductIPReputation       AttackLayer7TopIndustryParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopIndustryParamsMitigationProductAPIShield          AttackLayer7TopIndustryParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopIndustryParamsMitigationProductDataLossPrevention AttackLayer7TopIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopIndustryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopIndustryParamsMitigationProductDDoS, AttackLayer7TopIndustryParamsMitigationProductWAF, AttackLayer7TopIndustryParamsMitigationProductBotManagement, AttackLayer7TopIndustryParamsMitigationProductAccessRules, AttackLayer7TopIndustryParamsMitigationProductIPReputation, AttackLayer7TopIndustryParamsMitigationProductAPIShield, AttackLayer7TopIndustryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TopIndustryResponseEnvelope struct {
	Result  AttackLayer7TopIndustryResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer7TopIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopIndustryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopIndustryResponseEnvelope]
type attackLayer7TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopVerticalParams struct {
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
	Format param.Field[AttackLayer7TopVerticalParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TopVerticalParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TopVerticalParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TopVerticalParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopVerticalParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopVerticalParamsFormat string

const (
	AttackLayer7TopVerticalParamsFormatJson AttackLayer7TopVerticalParamsFormat = "JSON"
	AttackLayer7TopVerticalParamsFormatCsv  AttackLayer7TopVerticalParamsFormat = "CSV"
)

func (r AttackLayer7TopVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsFormatJson, AttackLayer7TopVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopVerticalParamsHTTPMethod string

const (
	AttackLayer7TopVerticalParamsHTTPMethodGet             AttackLayer7TopVerticalParamsHTTPMethod = "GET"
	AttackLayer7TopVerticalParamsHTTPMethodPost            AttackLayer7TopVerticalParamsHTTPMethod = "POST"
	AttackLayer7TopVerticalParamsHTTPMethodDelete          AttackLayer7TopVerticalParamsHTTPMethod = "DELETE"
	AttackLayer7TopVerticalParamsHTTPMethodPut             AttackLayer7TopVerticalParamsHTTPMethod = "PUT"
	AttackLayer7TopVerticalParamsHTTPMethodHead            AttackLayer7TopVerticalParamsHTTPMethod = "HEAD"
	AttackLayer7TopVerticalParamsHTTPMethodPurge           AttackLayer7TopVerticalParamsHTTPMethod = "PURGE"
	AttackLayer7TopVerticalParamsHTTPMethodOptions         AttackLayer7TopVerticalParamsHTTPMethod = "OPTIONS"
	AttackLayer7TopVerticalParamsHTTPMethodPropfind        AttackLayer7TopVerticalParamsHTTPMethod = "PROPFIND"
	AttackLayer7TopVerticalParamsHTTPMethodMkcol           AttackLayer7TopVerticalParamsHTTPMethod = "MKCOL"
	AttackLayer7TopVerticalParamsHTTPMethodPatch           AttackLayer7TopVerticalParamsHTTPMethod = "PATCH"
	AttackLayer7TopVerticalParamsHTTPMethodACL             AttackLayer7TopVerticalParamsHTTPMethod = "ACL"
	AttackLayer7TopVerticalParamsHTTPMethodBcopy           AttackLayer7TopVerticalParamsHTTPMethod = "BCOPY"
	AttackLayer7TopVerticalParamsHTTPMethodBdelete         AttackLayer7TopVerticalParamsHTTPMethod = "BDELETE"
	AttackLayer7TopVerticalParamsHTTPMethodBmove           AttackLayer7TopVerticalParamsHTTPMethod = "BMOVE"
	AttackLayer7TopVerticalParamsHTTPMethodBpropfind       AttackLayer7TopVerticalParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TopVerticalParamsHTTPMethodBproppatch      AttackLayer7TopVerticalParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TopVerticalParamsHTTPMethodCheckin         AttackLayer7TopVerticalParamsHTTPMethod = "CHECKIN"
	AttackLayer7TopVerticalParamsHTTPMethodCheckout        AttackLayer7TopVerticalParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TopVerticalParamsHTTPMethodConnect         AttackLayer7TopVerticalParamsHTTPMethod = "CONNECT"
	AttackLayer7TopVerticalParamsHTTPMethodCopy            AttackLayer7TopVerticalParamsHTTPMethod = "COPY"
	AttackLayer7TopVerticalParamsHTTPMethodLabel           AttackLayer7TopVerticalParamsHTTPMethod = "LABEL"
	AttackLayer7TopVerticalParamsHTTPMethodLock            AttackLayer7TopVerticalParamsHTTPMethod = "LOCK"
	AttackLayer7TopVerticalParamsHTTPMethodMerge           AttackLayer7TopVerticalParamsHTTPMethod = "MERGE"
	AttackLayer7TopVerticalParamsHTTPMethodMkactivity      AttackLayer7TopVerticalParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TopVerticalParamsHTTPMethodMkworkspace     AttackLayer7TopVerticalParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TopVerticalParamsHTTPMethodMove            AttackLayer7TopVerticalParamsHTTPMethod = "MOVE"
	AttackLayer7TopVerticalParamsHTTPMethodNotify          AttackLayer7TopVerticalParamsHTTPMethod = "NOTIFY"
	AttackLayer7TopVerticalParamsHTTPMethodOrderpatch      AttackLayer7TopVerticalParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TopVerticalParamsHTTPMethodPoll            AttackLayer7TopVerticalParamsHTTPMethod = "POLL"
	AttackLayer7TopVerticalParamsHTTPMethodProppatch       AttackLayer7TopVerticalParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TopVerticalParamsHTTPMethodReport          AttackLayer7TopVerticalParamsHTTPMethod = "REPORT"
	AttackLayer7TopVerticalParamsHTTPMethodSearch          AttackLayer7TopVerticalParamsHTTPMethod = "SEARCH"
	AttackLayer7TopVerticalParamsHTTPMethodSubscribe       AttackLayer7TopVerticalParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TopVerticalParamsHTTPMethodTrace           AttackLayer7TopVerticalParamsHTTPMethod = "TRACE"
	AttackLayer7TopVerticalParamsHTTPMethodUncheckout      AttackLayer7TopVerticalParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TopVerticalParamsHTTPMethodUnlock          AttackLayer7TopVerticalParamsHTTPMethod = "UNLOCK"
	AttackLayer7TopVerticalParamsHTTPMethodUnsubscribe     AttackLayer7TopVerticalParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TopVerticalParamsHTTPMethodUpdate          AttackLayer7TopVerticalParamsHTTPMethod = "UPDATE"
	AttackLayer7TopVerticalParamsHTTPMethodVersioncontrol  AttackLayer7TopVerticalParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TopVerticalParamsHTTPMethodBaselinecontrol AttackLayer7TopVerticalParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TopVerticalParamsHTTPMethodXmsenumatts     AttackLayer7TopVerticalParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TopVerticalParamsHTTPMethodRpcOutData      AttackLayer7TopVerticalParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TopVerticalParamsHTTPMethodRpcInData       AttackLayer7TopVerticalParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TopVerticalParamsHTTPMethodJson            AttackLayer7TopVerticalParamsHTTPMethod = "JSON"
	AttackLayer7TopVerticalParamsHTTPMethodCook            AttackLayer7TopVerticalParamsHTTPMethod = "COOK"
	AttackLayer7TopVerticalParamsHTTPMethodTrack           AttackLayer7TopVerticalParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TopVerticalParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsHTTPMethodGet, AttackLayer7TopVerticalParamsHTTPMethodPost, AttackLayer7TopVerticalParamsHTTPMethodDelete, AttackLayer7TopVerticalParamsHTTPMethodPut, AttackLayer7TopVerticalParamsHTTPMethodHead, AttackLayer7TopVerticalParamsHTTPMethodPurge, AttackLayer7TopVerticalParamsHTTPMethodOptions, AttackLayer7TopVerticalParamsHTTPMethodPropfind, AttackLayer7TopVerticalParamsHTTPMethodMkcol, AttackLayer7TopVerticalParamsHTTPMethodPatch, AttackLayer7TopVerticalParamsHTTPMethodACL, AttackLayer7TopVerticalParamsHTTPMethodBcopy, AttackLayer7TopVerticalParamsHTTPMethodBdelete, AttackLayer7TopVerticalParamsHTTPMethodBmove, AttackLayer7TopVerticalParamsHTTPMethodBpropfind, AttackLayer7TopVerticalParamsHTTPMethodBproppatch, AttackLayer7TopVerticalParamsHTTPMethodCheckin, AttackLayer7TopVerticalParamsHTTPMethodCheckout, AttackLayer7TopVerticalParamsHTTPMethodConnect, AttackLayer7TopVerticalParamsHTTPMethodCopy, AttackLayer7TopVerticalParamsHTTPMethodLabel, AttackLayer7TopVerticalParamsHTTPMethodLock, AttackLayer7TopVerticalParamsHTTPMethodMerge, AttackLayer7TopVerticalParamsHTTPMethodMkactivity, AttackLayer7TopVerticalParamsHTTPMethodMkworkspace, AttackLayer7TopVerticalParamsHTTPMethodMove, AttackLayer7TopVerticalParamsHTTPMethodNotify, AttackLayer7TopVerticalParamsHTTPMethodOrderpatch, AttackLayer7TopVerticalParamsHTTPMethodPoll, AttackLayer7TopVerticalParamsHTTPMethodProppatch, AttackLayer7TopVerticalParamsHTTPMethodReport, AttackLayer7TopVerticalParamsHTTPMethodSearch, AttackLayer7TopVerticalParamsHTTPMethodSubscribe, AttackLayer7TopVerticalParamsHTTPMethodTrace, AttackLayer7TopVerticalParamsHTTPMethodUncheckout, AttackLayer7TopVerticalParamsHTTPMethodUnlock, AttackLayer7TopVerticalParamsHTTPMethodUnsubscribe, AttackLayer7TopVerticalParamsHTTPMethodUpdate, AttackLayer7TopVerticalParamsHTTPMethodVersioncontrol, AttackLayer7TopVerticalParamsHTTPMethodBaselinecontrol, AttackLayer7TopVerticalParamsHTTPMethodXmsenumatts, AttackLayer7TopVerticalParamsHTTPMethodRpcOutData, AttackLayer7TopVerticalParamsHTTPMethodRpcInData, AttackLayer7TopVerticalParamsHTTPMethodJson, AttackLayer7TopVerticalParamsHTTPMethodCook, AttackLayer7TopVerticalParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TopVerticalParamsHTTPVersion string

const (
	AttackLayer7TopVerticalParamsHTTPVersionHttPv1 AttackLayer7TopVerticalParamsHTTPVersion = "HTTPv1"
	AttackLayer7TopVerticalParamsHTTPVersionHttPv2 AttackLayer7TopVerticalParamsHTTPVersion = "HTTPv2"
	AttackLayer7TopVerticalParamsHTTPVersionHttPv3 AttackLayer7TopVerticalParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TopVerticalParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsHTTPVersionHttPv1, AttackLayer7TopVerticalParamsHTTPVersionHttPv2, AttackLayer7TopVerticalParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TopVerticalParamsIPVersion string

const (
	AttackLayer7TopVerticalParamsIPVersionIPv4 AttackLayer7TopVerticalParamsIPVersion = "IPv4"
	AttackLayer7TopVerticalParamsIPVersionIPv6 AttackLayer7TopVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer7TopVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsIPVersionIPv4, AttackLayer7TopVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TopVerticalParamsMitigationProduct string

const (
	AttackLayer7TopVerticalParamsMitigationProductDDoS               AttackLayer7TopVerticalParamsMitigationProduct = "DDOS"
	AttackLayer7TopVerticalParamsMitigationProductWAF                AttackLayer7TopVerticalParamsMitigationProduct = "WAF"
	AttackLayer7TopVerticalParamsMitigationProductBotManagement      AttackLayer7TopVerticalParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopVerticalParamsMitigationProductAccessRules        AttackLayer7TopVerticalParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopVerticalParamsMitigationProductIPReputation       AttackLayer7TopVerticalParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopVerticalParamsMitigationProductAPIShield          AttackLayer7TopVerticalParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopVerticalParamsMitigationProductDataLossPrevention AttackLayer7TopVerticalParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopVerticalParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopVerticalParamsMitigationProductDDoS, AttackLayer7TopVerticalParamsMitigationProductWAF, AttackLayer7TopVerticalParamsMitigationProductBotManagement, AttackLayer7TopVerticalParamsMitigationProductAccessRules, AttackLayer7TopVerticalParamsMitigationProductIPReputation, AttackLayer7TopVerticalParamsMitigationProductAPIShield, AttackLayer7TopVerticalParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TopVerticalResponseEnvelope struct {
	Result  AttackLayer7TopVerticalResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer7TopVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopVerticalResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopVerticalResponseEnvelope]
type attackLayer7TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
