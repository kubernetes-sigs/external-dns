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

// EmailRoutingTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingTimeseriesGroupService] method instead.
type EmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewEmailRoutingTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *EmailRoutingTimeseriesGroupService) {
	r = &EmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
func (r *EmailRoutingTimeseriesGroupService) ARC(ctx context.Context, query EmailRoutingTimeseriesGroupARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupARCResponse, err error) {
	var env EmailRoutingTimeseriesGroupARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
func (r *EmailRoutingTimeseriesGroupService) DKIM(ctx context.Context, query EmailRoutingTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDKIMResponse, err error) {
	var env EmailRoutingTimeseriesGroupDKIMResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
func (r *EmailRoutingTimeseriesGroupService) DMARC(ctx context.Context, query EmailRoutingTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDMARCResponse, err error) {
	var env EmailRoutingTimeseriesGroupDMARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted) over time.
func (r *EmailRoutingTimeseriesGroupService) Encrypted(ctx context.Context, query EmailRoutingTimeseriesGroupEncryptedParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupEncryptedResponse, err error) {
	var env EmailRoutingTimeseriesGroupEncryptedResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by IP version over time.
func (r *EmailRoutingTimeseriesGroupService) IPVersion(ctx context.Context, query EmailRoutingTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupIPVersionResponse, err error) {
	var env EmailRoutingTimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
func (r *EmailRoutingTimeseriesGroupService) SPF(ctx context.Context, query EmailRoutingTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupSPFResponse, err error) {
	var env EmailRoutingTimeseriesGroupSPFResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingTimeseriesGroupARCResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                           `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupARCResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupARCResponse]
type emailRoutingTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupARCResponseMeta]
type emailRoutingTimeseriesGroupARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupARCResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupARCResponseMetaDateRange]
type emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupARCResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupARCResponseMetaUnit]
type emailRoutingTimeseriesGroupARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupDKIMResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDKIMResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDKIMResponse]
type emailRoutingTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupDKIMResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupDKIMResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupDKIMResponseMeta]
type emailRoutingTimeseriesGroupDKIMResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange]
type emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDKIMResponseMetaUnit]
type emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupDMARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                             `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDMARCResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDMARCResponse]
type emailRoutingTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupDMARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupDMARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupDMARCResponseMeta]
type emailRoutingTimeseriesGroupDMARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange]
type emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseMetaUnit]
type emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupEncryptedResponseMeta   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupEncryptedResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupEncryptedResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupEncryptedResponse]
type emailRoutingTimeseriesGroupEncryptedResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupEncryptedResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupEncryptedResponseMeta]
type emailRoutingTimeseriesGroupEncryptedResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange]
type emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit]
type emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponseSerie0 struct {
	Encrypted    []string                                               `json:"ENCRYPTED,required"`
	NotEncrypted []string                                               `json:"NOT_ENCRYPTED,required"`
	JSON         emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseSerie0]
type emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupIPVersionResponse]
type emailRoutingTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupIPVersionResponseMeta]
type emailRoutingTimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange]
type emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit]
type emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                               `json:"IPv4,required"`
	IPv6 []string                                               `json:"IPv6,required"`
	JSON emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseSerie0]
type emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupSPFResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                           `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupSPFResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupSPFResponse]
type emailRoutingTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupSPFResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupSPFResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupSPFResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupSPFResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupSPFResponseMeta]
type emailRoutingTimeseriesGroupSPFResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupSPFResponseMetaDateRange]
type emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupSPFResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupSPFResponseMetaUnit]
type emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupARCParamsAggInterval15m EmailRoutingTimeseriesGroupARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1h  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1d  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1w  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsAggInterval15m, EmailRoutingTimeseriesGroupARCParamsAggInterval1h, EmailRoutingTimeseriesGroupARCParamsAggInterval1d, EmailRoutingTimeseriesGroupARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupARCParamsDKIMPass EmailRoutingTimeseriesGroupARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDKIMNone EmailRoutingTimeseriesGroupARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDKIMFail EmailRoutingTimeseriesGroupARCParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsDKIMPass, EmailRoutingTimeseriesGroupARCParamsDKIMNone, EmailRoutingTimeseriesGroupARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsDMARC string

const (
	EmailRoutingTimeseriesGroupARCParamsDMARCPass EmailRoutingTimeseriesGroupARCParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDMARCNone EmailRoutingTimeseriesGroupARCParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDMARCFail EmailRoutingTimeseriesGroupARCParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsDMARCPass, EmailRoutingTimeseriesGroupARCParamsDMARCNone, EmailRoutingTimeseriesGroupARCParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupARCParamsFormatJson EmailRoutingTimeseriesGroupARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupARCParamsFormatCsv  EmailRoutingTimeseriesGroupARCParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsFormatJson, EmailRoutingTimeseriesGroupARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4, EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupARCParamsSPFPass EmailRoutingTimeseriesGroupARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupARCParamsSPFNone EmailRoutingTimeseriesGroupARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupARCParamsSPFFail EmailRoutingTimeseriesGroupARCParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsSPFPass, EmailRoutingTimeseriesGroupARCParamsSPFNone, EmailRoutingTimeseriesGroupARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupARCResponseEnvelope]
type emailRoutingTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDKIMParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDKIMParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1d  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1w  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1d, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsARCPass EmailRoutingTimeseriesGroupDKIMParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsARCNone EmailRoutingTimeseriesGroupDKIMParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsARCFail EmailRoutingTimeseriesGroupDKIMParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsARCPass, EmailRoutingTimeseriesGroupDKIMParamsARCNone, EmailRoutingTimeseriesGroupDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsDMARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsDMARCPass EmailRoutingTimeseriesGroupDKIMParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCNone EmailRoutingTimeseriesGroupDKIMParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCFail EmailRoutingTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsDMARCPass, EmailRoutingTimeseriesGroupDKIMParamsDMARCNone, EmailRoutingTimeseriesGroupDKIMParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupDKIMParamsFormat string

const (
	EmailRoutingTimeseriesGroupDKIMParamsFormatJson EmailRoutingTimeseriesGroupDKIMParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDKIMParamsFormatCsv  EmailRoutingTimeseriesGroupDKIMParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsFormatJson, EmailRoutingTimeseriesGroupDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4, EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsSPF string

const (
	EmailRoutingTimeseriesGroupDKIMParamsSPFPass EmailRoutingTimeseriesGroupDKIMParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsSPFNone EmailRoutingTimeseriesGroupDKIMParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsSPFFail EmailRoutingTimeseriesGroupDKIMParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsSPFPass, EmailRoutingTimeseriesGroupDKIMParamsSPFNone, EmailRoutingTimeseriesGroupDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDKIMResponseEnvelope]
type emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDMARCParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDMARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1d  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1w  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1d, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsARC string

const (
	EmailRoutingTimeseriesGroupDMARCParamsARCPass EmailRoutingTimeseriesGroupDMARCParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsARCNone EmailRoutingTimeseriesGroupDMARCParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsARCFail EmailRoutingTimeseriesGroupDMARCParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsARCPass, EmailRoutingTimeseriesGroupDMARCParamsARCNone, EmailRoutingTimeseriesGroupDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupDMARCParamsDKIMPass EmailRoutingTimeseriesGroupDMARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMNone EmailRoutingTimeseriesGroupDMARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMFail EmailRoutingTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsDKIMPass, EmailRoutingTimeseriesGroupDMARCParamsDKIMNone, EmailRoutingTimeseriesGroupDMARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupDMARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupDMARCParamsFormatJson EmailRoutingTimeseriesGroupDMARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDMARCParamsFormatCsv  EmailRoutingTimeseriesGroupDMARCParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsFormatJson, EmailRoutingTimeseriesGroupDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4, EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupDMARCParamsSPFPass EmailRoutingTimeseriesGroupDMARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsSPFNone EmailRoutingTimeseriesGroupDMARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsSPFFail EmailRoutingTimeseriesGroupDMARCParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsSPFPass, EmailRoutingTimeseriesGroupDMARCParamsSPFNone, EmailRoutingTimeseriesGroupDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseEnvelope]
type emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupEncryptedParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupEncryptedParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsARCPass EmailRoutingTimeseriesGroupEncryptedParamsARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsARCNone EmailRoutingTimeseriesGroupEncryptedParamsARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsARCFail EmailRoutingTimeseriesGroupEncryptedParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsARCPass, EmailRoutingTimeseriesGroupEncryptedParamsARCNone, EmailRoutingTimeseriesGroupEncryptedParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsDKIM string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass, EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone, EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsDMARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass, EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone, EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupEncryptedParamsFormat string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsFormatJson EmailRoutingTimeseriesGroupEncryptedParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupEncryptedParamsFormatCsv  EmailRoutingTimeseriesGroupEncryptedParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsFormatJson, EmailRoutingTimeseriesGroupEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4, EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsSPF string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsSPFPass EmailRoutingTimeseriesGroupEncryptedParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFNone EmailRoutingTimeseriesGroupEncryptedParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFFail EmailRoutingTimeseriesGroupEncryptedParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsSPFPass, EmailRoutingTimeseriesGroupEncryptedParamsSPFNone, EmailRoutingTimeseriesGroupEncryptedParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupEncryptedResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseEnvelope]
type emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupIPVersionParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsARCPass EmailRoutingTimeseriesGroupIPVersionParamsARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsARCNone EmailRoutingTimeseriesGroupIPVersionParamsARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsARCFail EmailRoutingTimeseriesGroupIPVersionParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsARCPass, EmailRoutingTimeseriesGroupIPVersionParamsARCNone, EmailRoutingTimeseriesGroupIPVersionParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsDKIM string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass, EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone, EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsDMARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass, EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone, EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupIPVersionParamsFormat string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsFormatJson EmailRoutingTimeseriesGroupIPVersionParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupIPVersionParamsFormatCsv  EmailRoutingTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsFormatJson, EmailRoutingTimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsSPF string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsSPFPass EmailRoutingTimeseriesGroupIPVersionParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFNone EmailRoutingTimeseriesGroupIPVersionParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFFail EmailRoutingTimeseriesGroupIPVersionParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsSPFPass, EmailRoutingTimeseriesGroupIPVersionParamsSPFNone, EmailRoutingTimeseriesGroupIPVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseEnvelope]
type emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupSPFParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupSPFParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupSPFParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupSPFParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupSPFParamsAggInterval15m EmailRoutingTimeseriesGroupSPFParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1h  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1d  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1w  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupSPFParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsAggInterval15m, EmailRoutingTimeseriesGroupSPFParamsAggInterval1h, EmailRoutingTimeseriesGroupSPFParamsAggInterval1d, EmailRoutingTimeseriesGroupSPFParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsARCPass EmailRoutingTimeseriesGroupSPFParamsARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsARCNone EmailRoutingTimeseriesGroupSPFParamsARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsARCFail EmailRoutingTimeseriesGroupSPFParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsARCPass, EmailRoutingTimeseriesGroupSPFParamsARCNone, EmailRoutingTimeseriesGroupSPFParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsDKIM string

const (
	EmailRoutingTimeseriesGroupSPFParamsDKIMPass EmailRoutingTimeseriesGroupSPFParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDKIMNone EmailRoutingTimeseriesGroupSPFParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDKIMFail EmailRoutingTimeseriesGroupSPFParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsDKIMPass, EmailRoutingTimeseriesGroupSPFParamsDKIMNone, EmailRoutingTimeseriesGroupSPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsDMARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsDMARCPass EmailRoutingTimeseriesGroupSPFParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDMARCNone EmailRoutingTimeseriesGroupSPFParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDMARCFail EmailRoutingTimeseriesGroupSPFParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsDMARCPass, EmailRoutingTimeseriesGroupSPFParamsDMARCNone, EmailRoutingTimeseriesGroupSPFParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupSPFParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupSPFParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupSPFParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupSPFParamsFormat string

const (
	EmailRoutingTimeseriesGroupSPFParamsFormatJson EmailRoutingTimeseriesGroupSPFParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupSPFParamsFormatCsv  EmailRoutingTimeseriesGroupSPFParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupSPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsFormatJson, EmailRoutingTimeseriesGroupSPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupSPFParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4, EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupSPFResponseEnvelope]
type emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
