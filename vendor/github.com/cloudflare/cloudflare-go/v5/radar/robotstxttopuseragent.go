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

// RobotsTXTTopUserAgentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTTopUserAgentService] method instead.
type RobotsTXTTopUserAgentService struct {
	Options []option.RequestOption
}

// NewRobotsTXTTopUserAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRobotsTXTTopUserAgentService(opts ...option.RequestOption) (r *RobotsTXTTopUserAgentService) {
	r = &RobotsTXTTopUserAgentService{}
	r.Options = opts
	return
}

// Retrieves the top user agents on robots.txt files.
func (r *RobotsTXTTopUserAgentService) Directive(ctx context.Context, query RobotsTXTTopUserAgentDirectiveParams, opts ...option.RequestOption) (res *RobotsTXTTopUserAgentDirectiveResponse, err error) {
	var env RobotsTXTTopUserAgentDirectiveResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/user_agents/directive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RobotsTXTTopUserAgentDirectiveResponse struct {
	// Metadata for the results.
	Meta RobotsTXTTopUserAgentDirectiveResponseMeta   `json:"meta,required"`
	Top0 []RobotsTXTTopUserAgentDirectiveResponseTop0 `json:"top_0,required"`
	JSON robotsTXTTopUserAgentDirectiveResponseJSON   `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseJSON contains the JSON metadata for the
// struct [RobotsTXTTopUserAgentDirectiveResponse]
type robotsTXTTopUserAgentDirectiveResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RobotsTXTTopUserAgentDirectiveResponseMeta struct {
	ConfidenceInfo RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RobotsTXTTopUserAgentDirectiveResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RobotsTXTTopUserAgentDirectiveResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RobotsTXTTopUserAgentDirectiveResponseMetaUnit `json:"units,required"`
	JSON  robotsTXTTopUserAgentDirectiveResponseMetaJSON   `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaJSON contains the JSON metadata for
// the struct [RobotsTXTTopUserAgentDirectiveResponseMeta]
type robotsTXTTopUserAgentDirectiveResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo struct {
	Annotations []RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo]
type robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation]
type robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RobotsTXTTopUserAgentDirectiveResponseMetaDateRange]
type robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RobotsTXTTopUserAgentDirectiveResponseMetaNormalization string

const (
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationPercentage           RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "PERCENTAGE"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationMin0Max              RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "MIN0_MAX"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationMinMax               RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "MIN_MAX"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRawValues            RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "RAW_VALUES"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationPercentageChange     RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRollingAverage       RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "ROLLING_AVERAGE"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationOverlappedPercentage RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRatio                RobotsTXTTopUserAgentDirectiveResponseMetaNormalization = "RATIO"
)

func (r RobotsTXTTopUserAgentDirectiveResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationPercentage, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationMin0Max, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationMinMax, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRawValues, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationPercentageChange, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRollingAverage, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationOverlappedPercentage, RobotsTXTTopUserAgentDirectiveResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RobotsTXTTopUserAgentDirectiveResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON contains the JSON metadata
// for the struct [RobotsTXTTopUserAgentDirectiveResponseMetaUnit]
type robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveResponseTop0 struct {
	Name      string                                         `json:"name,required"`
	Value     int64                                          `json:"value,required"`
	Fully     int64                                          `json:"fully"`
	Partially int64                                          `json:"partially"`
	JSON      robotsTXTTopUserAgentDirectiveResponseTop0JSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseTop0JSON contains the JSON metadata for
// the struct [RobotsTXTTopUserAgentDirectiveResponseTop0]
type robotsTXTTopUserAgentDirectiveResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Fully       apijson.Field
	Partially   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RobotsTXTTopUserAgentDirectiveParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Filters results by robots.txt directive.
	Directive param.Field[RobotsTXTTopUserAgentDirectiveParamsDirective] `query:"directive"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format in which results will be returned.
	Format param.Field[RobotsTXTTopUserAgentDirectiveParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by user agent category.
	UserAgentCategory param.Field[RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RobotsTXTTopUserAgentDirectiveParams]'s query parameters as
// `url.Values`.
func (r RobotsTXTTopUserAgentDirectiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters results by robots.txt directive.
type RobotsTXTTopUserAgentDirectiveParamsDirective string

const (
	RobotsTXTTopUserAgentDirectiveParamsDirectiveAllow    RobotsTXTTopUserAgentDirectiveParamsDirective = "ALLOW"
	RobotsTXTTopUserAgentDirectiveParamsDirectiveDisallow RobotsTXTTopUserAgentDirectiveParamsDirective = "DISALLOW"
)

func (r RobotsTXTTopUserAgentDirectiveParamsDirective) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsDirectiveAllow, RobotsTXTTopUserAgentDirectiveParamsDirectiveDisallow:
		return true
	}
	return false
}

// Format in which results will be returned.
type RobotsTXTTopUserAgentDirectiveParamsFormat string

const (
	RobotsTXTTopUserAgentDirectiveParamsFormatJson RobotsTXTTopUserAgentDirectiveParamsFormat = "JSON"
	RobotsTXTTopUserAgentDirectiveParamsFormatCsv  RobotsTXTTopUserAgentDirectiveParamsFormat = "CSV"
)

func (r RobotsTXTTopUserAgentDirectiveParamsFormat) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsFormatJson, RobotsTXTTopUserAgentDirectiveParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by user agent category.
type RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory string

const (
	RobotsTXTTopUserAgentDirectiveParamsUserAgentCategoryAI RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory = "AI"
)

func (r RobotsTXTTopUserAgentDirectiveParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RobotsTXTTopUserAgentDirectiveParamsUserAgentCategoryAI:
		return true
	}
	return false
}

type RobotsTXTTopUserAgentDirectiveResponseEnvelope struct {
	Result  RobotsTXTTopUserAgentDirectiveResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON `json:"-"`
}

// robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON contains the JSON metadata
// for the struct [RobotsTXTTopUserAgentDirectiveResponseEnvelope]
type robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTTopUserAgentDirectiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTTopUserAgentDirectiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
