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

// LeakedCredentialSummaryService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialSummaryService] method instead.
type LeakedCredentialSummaryService struct {
	Options []option.RequestOption
}

// NewLeakedCredentialSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLeakedCredentialSummaryService(opts ...option.RequestOption) (r *LeakedCredentialSummaryService) {
	r = &LeakedCredentialSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP authentication requests by bot class.
func (r *LeakedCredentialSummaryService) BotClass(ctx context.Context, query LeakedCredentialSummaryBotClassParams, opts ...option.RequestOption) (res *LeakedCredentialSummaryBotClassResponse, err error) {
	var env LeakedCredentialSummaryBotClassResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/summary/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP authentication requests by compromised
// credential status.
func (r *LeakedCredentialSummaryService) Compromised(ctx context.Context, query LeakedCredentialSummaryCompromisedParams, opts ...option.RequestOption) (res *LeakedCredentialSummaryCompromisedResponse, err error) {
	var env LeakedCredentialSummaryCompromisedResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/summary/compromised"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LeakedCredentialSummaryBotClassResponse struct {
	// Metadata for the results.
	Meta     LeakedCredentialSummaryBotClassResponseMeta     `json:"meta,required"`
	Summary0 LeakedCredentialSummaryBotClassResponseSummary0 `json:"summary_0,required"`
	JSON     leakedCredentialSummaryBotClassResponseJSON     `json:"-"`
}

// leakedCredentialSummaryBotClassResponseJSON contains the JSON metadata for the
// struct [LeakedCredentialSummaryBotClassResponse]
type leakedCredentialSummaryBotClassResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type LeakedCredentialSummaryBotClassResponseMeta struct {
	ConfidenceInfo LeakedCredentialSummaryBotClassResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []LeakedCredentialSummaryBotClassResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization LeakedCredentialSummaryBotClassResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []LeakedCredentialSummaryBotClassResponseMetaUnit `json:"units,required"`
	JSON  leakedCredentialSummaryBotClassResponseMetaJSON   `json:"-"`
}

// leakedCredentialSummaryBotClassResponseMetaJSON contains the JSON metadata for
// the struct [LeakedCredentialSummaryBotClassResponseMeta]
type leakedCredentialSummaryBotClassResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseMetaJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryBotClassResponseMetaConfidenceInfo struct {
	Annotations []LeakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  leakedCredentialSummaryBotClassResponseMetaConfidenceInfoJSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [LeakedCredentialSummaryBotClassResponseMetaConfidenceInfo]
type leakedCredentialSummaryBotClassResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type LeakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            leakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [LeakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotation]
type leakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *LeakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryBotClassResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      leakedCredentialSummaryBotClassResponseMetaDateRangeJSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [LeakedCredentialSummaryBotClassResponseMetaDateRange]
type leakedCredentialSummaryBotClassResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type LeakedCredentialSummaryBotClassResponseMetaNormalization string

const (
	LeakedCredentialSummaryBotClassResponseMetaNormalizationPercentage           LeakedCredentialSummaryBotClassResponseMetaNormalization = "PERCENTAGE"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationMin0Max              LeakedCredentialSummaryBotClassResponseMetaNormalization = "MIN0_MAX"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationMinMax               LeakedCredentialSummaryBotClassResponseMetaNormalization = "MIN_MAX"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationRawValues            LeakedCredentialSummaryBotClassResponseMetaNormalization = "RAW_VALUES"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationPercentageChange     LeakedCredentialSummaryBotClassResponseMetaNormalization = "PERCENTAGE_CHANGE"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationRollingAverage       LeakedCredentialSummaryBotClassResponseMetaNormalization = "ROLLING_AVERAGE"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationOverlappedPercentage LeakedCredentialSummaryBotClassResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	LeakedCredentialSummaryBotClassResponseMetaNormalizationRatio                LeakedCredentialSummaryBotClassResponseMetaNormalization = "RATIO"
)

func (r LeakedCredentialSummaryBotClassResponseMetaNormalization) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryBotClassResponseMetaNormalizationPercentage, LeakedCredentialSummaryBotClassResponseMetaNormalizationMin0Max, LeakedCredentialSummaryBotClassResponseMetaNormalizationMinMax, LeakedCredentialSummaryBotClassResponseMetaNormalizationRawValues, LeakedCredentialSummaryBotClassResponseMetaNormalizationPercentageChange, LeakedCredentialSummaryBotClassResponseMetaNormalizationRollingAverage, LeakedCredentialSummaryBotClassResponseMetaNormalizationOverlappedPercentage, LeakedCredentialSummaryBotClassResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type LeakedCredentialSummaryBotClassResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  leakedCredentialSummaryBotClassResponseMetaUnitJSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseMetaUnitJSON contains the JSON metadata
// for the struct [LeakedCredentialSummaryBotClassResponseMetaUnit]
type leakedCredentialSummaryBotClassResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryBotClassResponseSummary0 struct {
	// A numeric string.
	Bot string `json:"bot,required"`
	// A numeric string.
	Human string                                              `json:"human,required"`
	JSON  leakedCredentialSummaryBotClassResponseSummary0JSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseSummary0JSON contains the JSON metadata
// for the struct [LeakedCredentialSummaryBotClassResponseSummary0]
type leakedCredentialSummaryBotClassResponseSummary0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryCompromisedResponse struct {
	// Metadata for the results.
	Meta     LeakedCredentialSummaryCompromisedResponseMeta     `json:"meta,required"`
	Summary0 LeakedCredentialSummaryCompromisedResponseSummary0 `json:"summary_0,required"`
	JSON     leakedCredentialSummaryCompromisedResponseJSON     `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseJSON contains the JSON metadata for
// the struct [LeakedCredentialSummaryCompromisedResponse]
type leakedCredentialSummaryCompromisedResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type LeakedCredentialSummaryCompromisedResponseMeta struct {
	ConfidenceInfo LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []LeakedCredentialSummaryCompromisedResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization LeakedCredentialSummaryCompromisedResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []LeakedCredentialSummaryCompromisedResponseMetaUnit `json:"units,required"`
	JSON  leakedCredentialSummaryCompromisedResponseMetaJSON   `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseMetaJSON contains the JSON metadata
// for the struct [LeakedCredentialSummaryCompromisedResponseMeta]
type leakedCredentialSummaryCompromisedResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseMetaJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfo struct {
	Annotations []LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoJSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfo]
type leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotation]
type leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *LeakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryCompromisedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      leakedCredentialSummaryCompromisedResponseMetaDateRangeJSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [LeakedCredentialSummaryCompromisedResponseMetaDateRange]
type leakedCredentialSummaryCompromisedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type LeakedCredentialSummaryCompromisedResponseMetaNormalization string

const (
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationPercentage           LeakedCredentialSummaryCompromisedResponseMetaNormalization = "PERCENTAGE"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationMin0Max              LeakedCredentialSummaryCompromisedResponseMetaNormalization = "MIN0_MAX"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationMinMax               LeakedCredentialSummaryCompromisedResponseMetaNormalization = "MIN_MAX"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationRawValues            LeakedCredentialSummaryCompromisedResponseMetaNormalization = "RAW_VALUES"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationPercentageChange     LeakedCredentialSummaryCompromisedResponseMetaNormalization = "PERCENTAGE_CHANGE"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationRollingAverage       LeakedCredentialSummaryCompromisedResponseMetaNormalization = "ROLLING_AVERAGE"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationOverlappedPercentage LeakedCredentialSummaryCompromisedResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	LeakedCredentialSummaryCompromisedResponseMetaNormalizationRatio                LeakedCredentialSummaryCompromisedResponseMetaNormalization = "RATIO"
)

func (r LeakedCredentialSummaryCompromisedResponseMetaNormalization) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryCompromisedResponseMetaNormalizationPercentage, LeakedCredentialSummaryCompromisedResponseMetaNormalizationMin0Max, LeakedCredentialSummaryCompromisedResponseMetaNormalizationMinMax, LeakedCredentialSummaryCompromisedResponseMetaNormalizationRawValues, LeakedCredentialSummaryCompromisedResponseMetaNormalizationPercentageChange, LeakedCredentialSummaryCompromisedResponseMetaNormalizationRollingAverage, LeakedCredentialSummaryCompromisedResponseMetaNormalizationOverlappedPercentage, LeakedCredentialSummaryCompromisedResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type LeakedCredentialSummaryCompromisedResponseMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  leakedCredentialSummaryCompromisedResponseMetaUnitJSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseMetaUnitJSON contains the JSON
// metadata for the struct [LeakedCredentialSummaryCompromisedResponseMetaUnit]
type leakedCredentialSummaryCompromisedResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryCompromisedResponseSummary0 struct {
	// A numeric string.
	Clean string `json:"CLEAN,required"`
	// A numeric string.
	Compromised string                                                 `json:"COMPROMISED,required"`
	JSON        leakedCredentialSummaryCompromisedResponseSummary0JSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseSummary0JSON contains the JSON
// metadata for the struct [LeakedCredentialSummaryCompromisedResponseSummary0]
type leakedCredentialSummaryCompromisedResponseSummary0JSON struct {
	Clean       apijson.Field
	Compromised apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryBotClassParams struct {
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]LeakedCredentialSummaryBotClassParamsCompromised] `query:"compromised"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[LeakedCredentialSummaryBotClassParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [LeakedCredentialSummaryBotClassParams]'s query parameters
// as `url.Values`.
func (r LeakedCredentialSummaryBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LeakedCredentialSummaryBotClassParamsCompromised string

const (
	LeakedCredentialSummaryBotClassParamsCompromisedClean       LeakedCredentialSummaryBotClassParamsCompromised = "CLEAN"
	LeakedCredentialSummaryBotClassParamsCompromisedCompromised LeakedCredentialSummaryBotClassParamsCompromised = "COMPROMISED"
)

func (r LeakedCredentialSummaryBotClassParamsCompromised) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryBotClassParamsCompromisedClean, LeakedCredentialSummaryBotClassParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialSummaryBotClassParamsFormat string

const (
	LeakedCredentialSummaryBotClassParamsFormatJson LeakedCredentialSummaryBotClassParamsFormat = "JSON"
	LeakedCredentialSummaryBotClassParamsFormatCsv  LeakedCredentialSummaryBotClassParamsFormat = "CSV"
)

func (r LeakedCredentialSummaryBotClassParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryBotClassParamsFormatJson, LeakedCredentialSummaryBotClassParamsFormatCsv:
		return true
	}
	return false
}

type LeakedCredentialSummaryBotClassResponseEnvelope struct {
	Result  LeakedCredentialSummaryBotClassResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    leakedCredentialSummaryBotClassResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialSummaryBotClassResponseEnvelopeJSON contains the JSON metadata
// for the struct [LeakedCredentialSummaryBotClassResponseEnvelope]
type leakedCredentialSummaryBotClassResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryBotClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryBotClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialSummaryCompromisedParams struct {
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]LeakedCredentialSummaryCompromisedParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[LeakedCredentialSummaryCompromisedParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [LeakedCredentialSummaryCompromisedParams]'s query
// parameters as `url.Values`.
func (r LeakedCredentialSummaryCompromisedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LeakedCredentialSummaryCompromisedParamsBotClass string

const (
	LeakedCredentialSummaryCompromisedParamsBotClassLikelyAutomated LeakedCredentialSummaryCompromisedParamsBotClass = "LIKELY_AUTOMATED"
	LeakedCredentialSummaryCompromisedParamsBotClassLikelyHuman     LeakedCredentialSummaryCompromisedParamsBotClass = "LIKELY_HUMAN"
)

func (r LeakedCredentialSummaryCompromisedParamsBotClass) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryCompromisedParamsBotClassLikelyAutomated, LeakedCredentialSummaryCompromisedParamsBotClassLikelyHuman:
		return true
	}
	return false
}

// Format in which results will be returned.
type LeakedCredentialSummaryCompromisedParamsFormat string

const (
	LeakedCredentialSummaryCompromisedParamsFormatJson LeakedCredentialSummaryCompromisedParamsFormat = "JSON"
	LeakedCredentialSummaryCompromisedParamsFormatCsv  LeakedCredentialSummaryCompromisedParamsFormat = "CSV"
)

func (r LeakedCredentialSummaryCompromisedParamsFormat) IsKnown() bool {
	switch r {
	case LeakedCredentialSummaryCompromisedParamsFormatJson, LeakedCredentialSummaryCompromisedParamsFormatCsv:
		return true
	}
	return false
}

type LeakedCredentialSummaryCompromisedResponseEnvelope struct {
	Result  LeakedCredentialSummaryCompromisedResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    leakedCredentialSummaryCompromisedResponseEnvelopeJSON `json:"-"`
}

// leakedCredentialSummaryCompromisedResponseEnvelopeJSON contains the JSON
// metadata for the struct [LeakedCredentialSummaryCompromisedResponseEnvelope]
type leakedCredentialSummaryCompromisedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialSummaryCompromisedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialSummaryCompromisedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
