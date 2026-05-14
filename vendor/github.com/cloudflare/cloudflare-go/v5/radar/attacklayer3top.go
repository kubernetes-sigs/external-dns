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

// AttackLayer3TopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3TopService] method instead.
type AttackLayer3TopService struct {
	Options   []option.RequestOption
	Locations *AttackLayer3TopLocationService
}

// NewAttackLayer3TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer3TopService(opts ...option.RequestOption) (r *AttackLayer3TopService) {
	r = &AttackLayer3TopService{}
	r.Options = opts
	r.Locations = NewAttackLayer3TopLocationService(opts...)
	return
}

// Retrieves the top layer 3 attacks from origin to target location. Values are a
// percentage out of the total layer 3 attacks (with billing country). You can
// optionally limit the number of attacks by origin/target location (useful if all
// the top attacks are from or to the same location).
func (r *AttackLayer3TopService) Attacks(ctx context.Context, query AttackLayer3TopAttacksParams, opts ...option.RequestOption) (res *AttackLayer3TopAttacksResponse, err error) {
	var env AttackLayer3TopAttacksResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/attacks"
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
func (r *AttackLayer3TopService) Industry(ctx context.Context, query AttackLayer3TopIndustryParams, opts ...option.RequestOption) (res *AttackLayer3TopIndustryResponse, err error) {
	var env AttackLayer3TopIndustryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/industry"
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
func (r *AttackLayer3TopService) Vertical(ctx context.Context, query AttackLayer3TopVerticalParams, opts ...option.RequestOption) (res *AttackLayer3TopVerticalResponse, err error) {
	var env AttackLayer3TopVerticalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TopAttacksResponse struct {
	// Metadata for the results.
	Meta AttackLayer3TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopAttacksResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopAttacksResponseJSON   `json:"-"`
}

// attackLayer3TopAttacksResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponse]
type attackLayer3TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TopAttacksResponseMeta struct {
	ConfidenceInfo AttackLayer3TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer3TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TopAttacksResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TopAttacksResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TopAttacksResponseMetaJSON   `json:"-"`
}

// attackLayer3TopAttacksResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponseMeta]
type attackLayer3TopAttacksResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  attackLayer3TopAttacksResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopAttacksResponseMetaConfidenceInfo]
type attackLayer3TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopAttacksResponseMetaDateRange]
type attackLayer3TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopAttacksResponseMetaNormalization string

const (
	AttackLayer3TopAttacksResponseMetaNormalizationPercentage           AttackLayer3TopAttacksResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TopAttacksResponseMetaNormalizationMin0Max              AttackLayer3TopAttacksResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TopAttacksResponseMetaNormalizationMinMax               AttackLayer3TopAttacksResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TopAttacksResponseMetaNormalizationRawValues            AttackLayer3TopAttacksResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TopAttacksResponseMetaNormalizationPercentageChange     AttackLayer3TopAttacksResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TopAttacksResponseMetaNormalizationRollingAverage       AttackLayer3TopAttacksResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TopAttacksResponseMetaNormalizationOverlappedPercentage AttackLayer3TopAttacksResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TopAttacksResponseMetaNormalizationRatio                AttackLayer3TopAttacksResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TopAttacksResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksResponseMetaNormalizationPercentage, AttackLayer3TopAttacksResponseMetaNormalizationMin0Max, AttackLayer3TopAttacksResponseMetaNormalizationMinMax, AttackLayer3TopAttacksResponseMetaNormalizationRawValues, AttackLayer3TopAttacksResponseMetaNormalizationPercentageChange, AttackLayer3TopAttacksResponseMetaNormalizationRollingAverage, AttackLayer3TopAttacksResponseMetaNormalizationOverlappedPercentage, AttackLayer3TopAttacksResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TopAttacksResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  attackLayer3TopAttacksResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3TopAttacksResponseMetaUnit]
type attackLayer3TopAttacksResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                 `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                 `json:"originCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                attackLayer3TopAttacksResponseTop0JSON `json:"-"`
}

// attackLayer3TopAttacksResponseTop0JSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponseTop0]
type attackLayer3TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponse struct {
	// Metadata for the results.
	Meta AttackLayer3TopIndustryResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopIndustryResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopIndustryResponseJSON   `json:"-"`
}

// attackLayer3TopIndustryResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopIndustryResponse]
type attackLayer3TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TopIndustryResponseMeta struct {
	ConfidenceInfo AttackLayer3TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer3TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TopIndustryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TopIndustryResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TopIndustryResponseMetaJSON   `json:"-"`
}

// attackLayer3TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseMeta]
type attackLayer3TopIndustryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  attackLayer3TopIndustryResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopIndustryResponseMetaConfidenceInfo]
type attackLayer3TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopIndustryResponseMetaDateRange]
type attackLayer3TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopIndustryResponseMetaNormalization string

const (
	AttackLayer3TopIndustryResponseMetaNormalizationPercentage           AttackLayer3TopIndustryResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TopIndustryResponseMetaNormalizationMin0Max              AttackLayer3TopIndustryResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TopIndustryResponseMetaNormalizationMinMax               AttackLayer3TopIndustryResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TopIndustryResponseMetaNormalizationRawValues            AttackLayer3TopIndustryResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TopIndustryResponseMetaNormalizationPercentageChange     AttackLayer3TopIndustryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TopIndustryResponseMetaNormalizationRollingAverage       AttackLayer3TopIndustryResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TopIndustryResponseMetaNormalizationOverlappedPercentage AttackLayer3TopIndustryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TopIndustryResponseMetaNormalizationRatio                AttackLayer3TopIndustryResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TopIndustryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopIndustryResponseMetaNormalizationPercentage, AttackLayer3TopIndustryResponseMetaNormalizationMin0Max, AttackLayer3TopIndustryResponseMetaNormalizationMinMax, AttackLayer3TopIndustryResponseMetaNormalizationRawValues, AttackLayer3TopIndustryResponseMetaNormalizationPercentageChange, AttackLayer3TopIndustryResponseMetaNormalizationRollingAverage, AttackLayer3TopIndustryResponseMetaNormalizationOverlappedPercentage, AttackLayer3TopIndustryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TopIndustryResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  attackLayer3TopIndustryResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseMetaUnit]
type attackLayer3TopIndustryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer3TopIndustryResponseTop0JSON `json:"-"`
}

// attackLayer3TopIndustryResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseTop0]
type attackLayer3TopIndustryResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponse struct {
	// Metadata for the results.
	Meta AttackLayer3TopVerticalResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopVerticalResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopVerticalResponseJSON   `json:"-"`
}

// attackLayer3TopVerticalResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopVerticalResponse]
type attackLayer3TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TopVerticalResponseMeta struct {
	ConfidenceInfo AttackLayer3TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer3TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TopVerticalResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TopVerticalResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TopVerticalResponseMetaJSON   `json:"-"`
}

// attackLayer3TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseMeta]
type attackLayer3TopVerticalResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  attackLayer3TopVerticalResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopVerticalResponseMetaConfidenceInfo]
type attackLayer3TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopVerticalResponseMetaDateRange]
type attackLayer3TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopVerticalResponseMetaNormalization string

const (
	AttackLayer3TopVerticalResponseMetaNormalizationPercentage           AttackLayer3TopVerticalResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TopVerticalResponseMetaNormalizationMin0Max              AttackLayer3TopVerticalResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TopVerticalResponseMetaNormalizationMinMax               AttackLayer3TopVerticalResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TopVerticalResponseMetaNormalizationRawValues            AttackLayer3TopVerticalResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TopVerticalResponseMetaNormalizationPercentageChange     AttackLayer3TopVerticalResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TopVerticalResponseMetaNormalizationRollingAverage       AttackLayer3TopVerticalResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TopVerticalResponseMetaNormalizationOverlappedPercentage AttackLayer3TopVerticalResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TopVerticalResponseMetaNormalizationRatio                AttackLayer3TopVerticalResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TopVerticalResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopVerticalResponseMetaNormalizationPercentage, AttackLayer3TopVerticalResponseMetaNormalizationMin0Max, AttackLayer3TopVerticalResponseMetaNormalizationMinMax, AttackLayer3TopVerticalResponseMetaNormalizationRawValues, AttackLayer3TopVerticalResponseMetaNormalizationPercentageChange, AttackLayer3TopVerticalResponseMetaNormalizationRollingAverage, AttackLayer3TopVerticalResponseMetaNormalizationOverlappedPercentage, AttackLayer3TopVerticalResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TopVerticalResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  attackLayer3TopVerticalResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseMetaUnit]
type attackLayer3TopVerticalResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer3TopVerticalResponseTop0JSON `json:"-"`
}

// attackLayer3TopVerticalResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseTop0]
type attackLayer3TopVerticalResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksParams struct {
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
	Format param.Field[AttackLayer3TopAttacksParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TopAttacksParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Specifies whether the `limitPerLocation` applies to the source or target
	// location.
	LimitDirection param.Field[AttackLayer3TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limits the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Orders results based on attack magnitude, defined by total mitigated bytes or
	// total mitigated attacks.
	Magnitude param.Field[AttackLayer3TopAttacksParamsMagnitude] `query:"magnitude"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TopAttacksParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TopAttacksParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopAttacksParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer3TopAttacksParamsFormat string

const (
	AttackLayer3TopAttacksParamsFormatJson AttackLayer3TopAttacksParamsFormat = "JSON"
	AttackLayer3TopAttacksParamsFormatCsv  AttackLayer3TopAttacksParamsFormat = "CSV"
)

func (r AttackLayer3TopAttacksParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsFormatJson, AttackLayer3TopAttacksParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopAttacksParamsIPVersion string

const (
	AttackLayer3TopAttacksParamsIPVersionIPv4 AttackLayer3TopAttacksParamsIPVersion = "IPv4"
	AttackLayer3TopAttacksParamsIPVersionIPv6 AttackLayer3TopAttacksParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopAttacksParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsIPVersionIPv4, AttackLayer3TopAttacksParamsIPVersionIPv6:
		return true
	}
	return false
}

// Specifies whether the `limitPerLocation` applies to the source or target
// location.
type AttackLayer3TopAttacksParamsLimitDirection string

const (
	AttackLayer3TopAttacksParamsLimitDirectionOrigin AttackLayer3TopAttacksParamsLimitDirection = "ORIGIN"
	AttackLayer3TopAttacksParamsLimitDirectionTarget AttackLayer3TopAttacksParamsLimitDirection = "TARGET"
)

func (r AttackLayer3TopAttacksParamsLimitDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsLimitDirectionOrigin, AttackLayer3TopAttacksParamsLimitDirectionTarget:
		return true
	}
	return false
}

// Orders results based on attack magnitude, defined by total mitigated bytes or
// total mitigated attacks.
type AttackLayer3TopAttacksParamsMagnitude string

const (
	AttackLayer3TopAttacksParamsMagnitudeMitigatedBytes   AttackLayer3TopAttacksParamsMagnitude = "MITIGATED_BYTES"
	AttackLayer3TopAttacksParamsMagnitudeMitigatedAttacks AttackLayer3TopAttacksParamsMagnitude = "MITIGATED_ATTACKS"
)

func (r AttackLayer3TopAttacksParamsMagnitude) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsMagnitudeMitigatedBytes, AttackLayer3TopAttacksParamsMagnitudeMitigatedAttacks:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TopAttacksParamsNormalization string

const (
	AttackLayer3TopAttacksParamsNormalizationPercentage AttackLayer3TopAttacksParamsNormalization = "PERCENTAGE"
	AttackLayer3TopAttacksParamsNormalizationMinMax     AttackLayer3TopAttacksParamsNormalization = "MIN_MAX"
)

func (r AttackLayer3TopAttacksParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsNormalizationPercentage, AttackLayer3TopAttacksParamsNormalizationMinMax:
		return true
	}
	return false
}

type AttackLayer3TopAttacksParamsProtocol string

const (
	AttackLayer3TopAttacksParamsProtocolUdp  AttackLayer3TopAttacksParamsProtocol = "UDP"
	AttackLayer3TopAttacksParamsProtocolTCP  AttackLayer3TopAttacksParamsProtocol = "TCP"
	AttackLayer3TopAttacksParamsProtocolIcmp AttackLayer3TopAttacksParamsProtocol = "ICMP"
	AttackLayer3TopAttacksParamsProtocolGRE  AttackLayer3TopAttacksParamsProtocol = "GRE"
)

func (r AttackLayer3TopAttacksParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopAttacksParamsProtocolUdp, AttackLayer3TopAttacksParamsProtocolTCP, AttackLayer3TopAttacksParamsProtocolIcmp, AttackLayer3TopAttacksParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopAttacksResponseEnvelope struct {
	Result  AttackLayer3TopAttacksResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer3TopAttacksResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopAttacksResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopAttacksResponseEnvelope]
type attackLayer3TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryParams struct {
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
	Format param.Field[AttackLayer3TopIndustryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TopIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer3TopIndustryParamsFormat string

const (
	AttackLayer3TopIndustryParamsFormatJson AttackLayer3TopIndustryParamsFormat = "JSON"
	AttackLayer3TopIndustryParamsFormatCsv  AttackLayer3TopIndustryParamsFormat = "CSV"
)

func (r AttackLayer3TopIndustryParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopIndustryParamsFormatJson, AttackLayer3TopIndustryParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopIndustryParamsIPVersion string

const (
	AttackLayer3TopIndustryParamsIPVersionIPv4 AttackLayer3TopIndustryParamsIPVersion = "IPv4"
	AttackLayer3TopIndustryParamsIPVersionIPv6 AttackLayer3TopIndustryParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopIndustryParamsIPVersionIPv4, AttackLayer3TopIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopIndustryParamsProtocol string

const (
	AttackLayer3TopIndustryParamsProtocolUdp  AttackLayer3TopIndustryParamsProtocol = "UDP"
	AttackLayer3TopIndustryParamsProtocolTCP  AttackLayer3TopIndustryParamsProtocol = "TCP"
	AttackLayer3TopIndustryParamsProtocolIcmp AttackLayer3TopIndustryParamsProtocol = "ICMP"
	AttackLayer3TopIndustryParamsProtocolGRE  AttackLayer3TopIndustryParamsProtocol = "GRE"
)

func (r AttackLayer3TopIndustryParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopIndustryParamsProtocolUdp, AttackLayer3TopIndustryParamsProtocolTCP, AttackLayer3TopIndustryParamsProtocolIcmp, AttackLayer3TopIndustryParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopIndustryResponseEnvelope struct {
	Result  AttackLayer3TopIndustryResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer3TopIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopIndustryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseEnvelope]
type attackLayer3TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalParams struct {
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
	Format param.Field[AttackLayer3TopVerticalParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TopVerticalParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TopVerticalParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer3TopVerticalParamsFormat string

const (
	AttackLayer3TopVerticalParamsFormatJson AttackLayer3TopVerticalParamsFormat = "JSON"
	AttackLayer3TopVerticalParamsFormatCsv  AttackLayer3TopVerticalParamsFormat = "CSV"
)

func (r AttackLayer3TopVerticalParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopVerticalParamsFormatJson, AttackLayer3TopVerticalParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopVerticalParamsIPVersion string

const (
	AttackLayer3TopVerticalParamsIPVersionIPv4 AttackLayer3TopVerticalParamsIPVersion = "IPv4"
	AttackLayer3TopVerticalParamsIPVersionIPv6 AttackLayer3TopVerticalParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopVerticalParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopVerticalParamsIPVersionIPv4, AttackLayer3TopVerticalParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopVerticalParamsProtocol string

const (
	AttackLayer3TopVerticalParamsProtocolUdp  AttackLayer3TopVerticalParamsProtocol = "UDP"
	AttackLayer3TopVerticalParamsProtocolTCP  AttackLayer3TopVerticalParamsProtocol = "TCP"
	AttackLayer3TopVerticalParamsProtocolIcmp AttackLayer3TopVerticalParamsProtocol = "ICMP"
	AttackLayer3TopVerticalParamsProtocolGRE  AttackLayer3TopVerticalParamsProtocol = "GRE"
)

func (r AttackLayer3TopVerticalParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopVerticalParamsProtocolUdp, AttackLayer3TopVerticalParamsProtocolTCP, AttackLayer3TopVerticalParamsProtocolIcmp, AttackLayer3TopVerticalParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopVerticalResponseEnvelope struct {
	Result  AttackLayer3TopVerticalResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer3TopVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopVerticalResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseEnvelope]
type attackLayer3TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
