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

// RobotsTXTTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTTopService] method instead.
type RobotsTXTTopService struct {
	Options    []option.RequestOption
	UserAgents *RobotsTXTTopUserAgentService
}

// NewRobotsTXTTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRobotsTXTTopService(opts ...option.RequestOption) (r *RobotsTXTTopService) {
	r = &RobotsTXTTopService{}
	r.Options = opts
	r.UserAgents = NewRobotsTXTTopUserAgentService(opts...)
	return
}

// Retrieves the top domain categories by the number of robots.txt files parsed.
func (r *RobotsTXTTopService) DomainCategories(ctx context.Context, query RobotsTXTTopDomainCategoriesParams, opts ...option.RequestOption) (res *RobotsTXTTopDomainCategoriesResponse, err error) {
	var env RobotsTXTTopDomainCategoriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/domain_categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RobotsTXTTopDomainCategoriesResponse struct {
	// Metadata for the results.
	Meta RobotsTXTTopDomainCategoriesResponseMeta   `json:"meta,required"`
	Top0 []RobotsTXTTopDomainCategoriesResponseTop0 `json:"top_0,required"`
	JSON robotsTXTTopDomainCategoriesResponseJSON   `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseJSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponse]
type robotsTXTTopDomainCategoriesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RobotsTXTTopDomainCategoriesResponseMeta struct {
	ConfidenceInfo RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RobotsTXTTopDomainCategoriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RobotsTXTTopDomainCategoriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RobotsTXTTopDomainCategoriesResponseMetaUnit `json:"units,required"`
	JSON  robotsTXTTopDomainCategoriesResponseMetaJSON   `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaJSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponseMeta]
type robotsTXTTopDomainCategoriesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo struct {
	Annotations []RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo]
type robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation]
type robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RobotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RobotsTXTTopDomainCategoriesResponseMetaDateRange]
type robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RobotsTXTTopDomainCategoriesResponseMetaNormalization string

const (
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationPercentage           RobotsTXTTopDomainCategoriesResponseMetaNormalization = "PERCENTAGE"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationMin0Max              RobotsTXTTopDomainCategoriesResponseMetaNormalization = "MIN0_MAX"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationMinMax               RobotsTXTTopDomainCategoriesResponseMetaNormalization = "MIN_MAX"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationRawValues            RobotsTXTTopDomainCategoriesResponseMetaNormalization = "RAW_VALUES"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationPercentageChange     RobotsTXTTopDomainCategoriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationRollingAverage       RobotsTXTTopDomainCategoriesResponseMetaNormalization = "ROLLING_AVERAGE"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationOverlappedPercentage RobotsTXTTopDomainCategoriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RobotsTXTTopDomainCategoriesResponseMetaNormalizationRatio                RobotsTXTTopDomainCategoriesResponseMetaNormalization = "RATIO"
)

func (r RobotsTXTTopDomainCategoriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RobotsTXTTopDomainCategoriesResponseMetaNormalizationPercentage, RobotsTXTTopDomainCategoriesResponseMetaNormalizationMin0Max, RobotsTXTTopDomainCategoriesResponseMetaNormalizationMinMax, RobotsTXTTopDomainCategoriesResponseMetaNormalizationRawValues, RobotsTXTTopDomainCategoriesResponseMetaNormalizationPercentageChange, RobotsTXTTopDomainCategoriesResponseMetaNormalizationRollingAverage, RobotsTXTTopDomainCategoriesResponseMetaNormalizationOverlappedPercentage, RobotsTXTTopDomainCategoriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RobotsTXTTopDomainCategoriesResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  robotsTXTTopDomainCategoriesResponseMetaUnitJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseMetaUnitJSON contains the JSON metadata for
// the struct [RobotsTXTTopDomainCategoriesResponseMetaUnit]
type robotsTXTTopDomainCategoriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesResponseTop0 struct {
	Name  string                                       `json:"name,required"`
	Value int64                                        `json:"value,required"`
	JSON  robotsTXTTopDomainCategoriesResponseTop0JSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseTop0JSON contains the JSON metadata for the
// struct [RobotsTXTTopDomainCategoriesResponseTop0]
type robotsTXTTopDomainCategoriesResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopDomainCategoriesParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RobotsTXTTopDomainCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by user agent category.
	UserAgentCategory param.Field[RobotsTXTTopDomainCategoriesParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RobotsTXTTopDomainCategoriesParams]'s query parameters as
// `url.Values`.
func (r RobotsTXTTopDomainCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RobotsTXTTopDomainCategoriesParamsFormat string

const (
	RobotsTXTTopDomainCategoriesParamsFormatJson RobotsTXTTopDomainCategoriesParamsFormat = "JSON"
	RobotsTXTTopDomainCategoriesParamsFormatCsv  RobotsTXTTopDomainCategoriesParamsFormat = "CSV"
)

func (r RobotsTXTTopDomainCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RobotsTXTTopDomainCategoriesParamsFormatJson, RobotsTXTTopDomainCategoriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by user agent category.
type RobotsTXTTopDomainCategoriesParamsUserAgentCategory string

const (
	RobotsTXTTopDomainCategoriesParamsUserAgentCategoryAI RobotsTXTTopDomainCategoriesParamsUserAgentCategory = "AI"
)

func (r RobotsTXTTopDomainCategoriesParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RobotsTXTTopDomainCategoriesParamsUserAgentCategoryAI:
		return true
	}
	return false
}

type RobotsTXTTopDomainCategoriesResponseEnvelope struct {
	Result  RobotsTXTTopDomainCategoriesResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    robotsTXTTopDomainCategoriesResponseEnvelopeJSON `json:"-"`
}

// robotsTXTTopDomainCategoriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RobotsTXTTopDomainCategoriesResponseEnvelope]
type robotsTXTTopDomainCategoriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopDomainCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopDomainCategoriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
