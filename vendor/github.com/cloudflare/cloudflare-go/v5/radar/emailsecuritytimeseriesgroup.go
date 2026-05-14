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

// EmailSecurityTimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTimeseriesGroupService] method instead.
type EmailSecurityTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTimeseriesGroupService(opts ...option.RequestOption) (r *EmailSecurityTimeseriesGroupService) {
	r = &EmailSecurityTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
func (r *EmailSecurityTimeseriesGroupService) ARC(ctx context.Context, query EmailSecurityTimeseriesGroupARCParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupARCResponse, err error) {
	var env EmailSecurityTimeseriesGroupARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
func (r *EmailSecurityTimeseriesGroupService) DKIM(ctx context.Context, query EmailSecurityTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupDKIMResponse, err error) {
	var env EmailSecurityTimeseriesGroupDKIMResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
func (r *EmailSecurityTimeseriesGroupService) DMARC(ctx context.Context, query EmailSecurityTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupDMARCResponse, err error) {
	var env EmailSecurityTimeseriesGroupDMARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by malicious classification over time.
func (r *EmailSecurityTimeseriesGroupService) Malicious(ctx context.Context, query EmailSecurityTimeseriesGroupMaliciousParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupMaliciousResponse, err error) {
	var env EmailSecurityTimeseriesGroupMaliciousResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by spam classification (spam vs. non-spam)
// over time.
func (r *EmailSecurityTimeseriesGroupService) Spam(ctx context.Context, query EmailSecurityTimeseriesGroupSpamParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSpamResponse, err error) {
	var env EmailSecurityTimeseriesGroupSpamResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
func (r *EmailSecurityTimeseriesGroupService) SPF(ctx context.Context, query EmailSecurityTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSPFResponse, err error) {
	var env EmailSecurityTimeseriesGroupSPFResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by spoof classification (spoof vs.
// non-spoof) over time.
func (r *EmailSecurityTimeseriesGroupService) Spoof(ctx context.Context, query EmailSecurityTimeseriesGroupSpoofParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupSpoofResponse, err error) {
	var env EmailSecurityTimeseriesGroupSpoofResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by threat category over time.
func (r *EmailSecurityTimeseriesGroupService) ThreatCategory(ctx context.Context, query EmailSecurityTimeseriesGroupThreatCategoryParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupThreatCategoryResponse, err error) {
	var env EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by TLS version over time.
func (r *EmailSecurityTimeseriesGroupService) TLSVersion(ctx context.Context, query EmailSecurityTimeseriesGroupTLSVersionParams, opts ...option.RequestOption) (res *EmailSecurityTimeseriesGroupTLSVersionResponse, err error) {
	var env EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTimeseriesGroupARCResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupARCResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupARCResponse]
type emailSecurityTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupARCResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupARCResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupARCResponseMeta]
type emailSecurityTimeseriesGroupARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupARCResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupARCResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupARCResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupARCResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupARCResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupARCResponseMetaDateRange]
type emailSecurityTimeseriesGroupARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupARCResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupARCResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupARCResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupARCResponseMetaUnit]
type emailSecurityTimeseriesGroupARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDKIMResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupDKIMResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                             `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupDKIMResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupDKIMResponse]
type emailSecurityTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupDKIMResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupDKIMResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupDKIMResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupDKIMResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupDKIMResponseMeta]
type emailSecurityTimeseriesGroupDKIMResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupDKIMResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupDKIMResponseMetaDateRange]
type emailSecurityTimeseriesGroupDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupDKIMResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupDKIMResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupDKIMResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDKIMResponseMetaUnit]
type emailSecurityTimeseriesGroupDKIMResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDMARCResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupDMARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                              `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupDMARCResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupDMARCResponse]
type emailSecurityTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupDMARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupDMARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupDMARCResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupDMARCResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupDMARCResponseMeta]
type emailSecurityTimeseriesGroupDMARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupDMARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupDMARCResponseMetaDateRange]
type emailSecurityTimeseriesGroupDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupDMARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupDMARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupDMARCResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDMARCResponseMetaUnit]
type emailSecurityTimeseriesGroupDMARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupMaliciousResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupMaliciousResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupMaliciousResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupMaliciousResponse]
type emailSecurityTimeseriesGroupMaliciousResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupMaliciousResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupMaliciousResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupMaliciousResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupMaliciousResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseMetaJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupMaliciousResponseMeta]
type emailSecurityTimeseriesGroupMaliciousResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupMaliciousResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupMaliciousResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupMaliciousResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupMaliciousResponseMetaDateRange]
type emailSecurityTimeseriesGroupMaliciousResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupMaliciousResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousResponseMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupMaliciousResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupMaliciousResponseMetaUnit]
type emailSecurityTimeseriesGroupMaliciousResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousResponseSerie0 struct {
	Malicious    []string                                                `json:"MALICIOUS,required"`
	NotMalicious []string                                                `json:"NOT_MALICIOUS,required"`
	JSON         emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupMaliciousResponseSerie0]
type emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupSpamResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupSpamResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSpamResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSpamResponse]
type emailSecurityTimeseriesGroupSpamResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupSpamResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupSpamResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupSpamResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupSpamResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupSpamResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupSpamResponseMeta]
type emailSecurityTimeseriesGroupSpamResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupSpamResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupSpamResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupSpamResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupSpamResponseMetaDateRange]
type emailSecurityTimeseriesGroupSpamResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupSpamResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupSpamResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupSpamResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupSpamResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupSpamResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpamResponseMetaUnit]
type emailSecurityTimeseriesGroupSpamResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamResponseSerie0 struct {
	NotSpam []string                                           `json:"NOT_SPAM,required"`
	Spam    []string                                           `json:"SPAM,required"`
	JSON    emailSecurityTimeseriesGroupSpamResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseSerie0JSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpamResponseSerie0]
type emailSecurityTimeseriesGroupSpamResponseSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSPFResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupSPFResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSPFResponseJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSPFResponse]
type emailSecurityTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupSPFResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupSPFResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupSPFResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupSPFResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupSPFResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupSPFResponseMeta]
type emailSecurityTimeseriesGroupSPFResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupSPFResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupSPFResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupSPFResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupSPFResponseMetaDateRange]
type emailSecurityTimeseriesGroupSPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupSPFResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupSPFResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupSPFResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupSPFResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupSPFResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSPFResponseMetaUnit]
type emailSecurityTimeseriesGroupSPFResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupSpoofResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupSpoofResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupSpoofResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTimeseriesGroupSpoofResponse]
type emailSecurityTimeseriesGroupSpoofResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupSpoofResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupSpoofResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupSpoofResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupSpoofResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTimeseriesGroupSpoofResponseMeta]
type emailSecurityTimeseriesGroupSpoofResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupSpoofResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupSpoofResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupSpoofResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupSpoofResponseMetaDateRange]
type emailSecurityTimeseriesGroupSpoofResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupSpoofResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupSpoofResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupSpoofResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpoofResponseMetaUnit]
type emailSecurityTimeseriesGroupSpoofResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofResponseSerie0 struct {
	NotSpoof []string                                            `json:"NOT_SPOOF,required"`
	Spoof    []string                                            `json:"SPOOF,required"`
	JSON     emailSecurityTimeseriesGroupSpoofResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseSerie0JSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpoofResponseSerie0]
type emailSecurityTimeseriesGroupSpoofResponseSerie0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupThreatCategoryResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupThreatCategoryResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupThreatCategoryResponse]
type emailSecurityTimeseriesGroupThreatCategoryResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupThreatCategoryResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupThreatCategoryResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupThreatCategoryResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseMetaJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupThreatCategoryResponseMeta]
type emailSecurityTimeseriesGroupThreatCategoryResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                    `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                          `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRange]
type emailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupThreatCategoryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseMetaUnit struct {
	Name  string                                                         `json:"name,required"`
	Value string                                                         `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupThreatCategoryResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseMetaUnitJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseMetaUnit]
type emailSecurityTimeseriesGroupThreatCategoryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0 struct {
	BrandImpersonation  []string                                                     `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                     `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                     `json:"IdentityDeception,required"`
	Link                []string                                                     `json:"Link,required"`
	JSON                emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0]
type emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionResponse struct {
	// Metadata for the results.
	Meta   EmailSecurityTimeseriesGroupTLSVersionResponseMeta   `json:"meta,required"`
	Serie0 EmailSecurityTimeseriesGroupTLSVersionResponseSerie0 `json:"serie_0,required"`
	JSON   emailSecurityTimeseriesGroupTLSVersionResponseJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupTLSVersionResponse]
type emailSecurityTimeseriesGroupTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTimeseriesGroupTLSVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecurityTimeseriesGroupTLSVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTimeseriesGroupTLSVersionResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTimeseriesGroupTLSVersionResponseMetaJSON   `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseMetaJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseMeta]
type emailSecurityTimeseriesGroupTLSVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval string

const (
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalFifteenMinutes EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneHour        EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_HOUR"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneDay         EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_DAY"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneWeek        EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_WEEK"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneMonth       EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalFifteenMinutes, EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneHour, EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneDay, EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneWeek, EmailSecurityTimeseriesGroupTLSVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfo]
type emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation]
type emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTimeseriesGroupTLSVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [EmailSecurityTimeseriesGroupTLSVersionResponseMetaDateRange]
type emailSecurityTimeseriesGroupTLSVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization string

const (
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationPercentage           EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationMin0Max              EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationMinMax               EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRawValues            EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationPercentageChange     EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRollingAverage       EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationOverlappedPercentage EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRatio                EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationPercentage, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationMin0Max, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationMinMax, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRawValues, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationPercentageChange, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRollingAverage, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationOverlappedPercentage, EmailSecurityTimeseriesGroupTLSVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionResponseMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  emailSecurityTimeseriesGroupTLSVersionResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseMetaUnit]
type emailSecurityTimeseriesGroupTLSVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionResponseSerie0 struct {
	TLS1_0 []string                                                 `json:"TLS 1.0,required"`
	TLS1_1 []string                                                 `json:"TLS 1.1,required"`
	TLS1_2 []string                                                 `json:"TLS 1.2,required"`
	TLS1_3 []string                                                 `json:"TLS 1.3,required"`
	JSON   emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseSerie0]
type emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON struct {
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupARCParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupARCParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupARCParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupARCParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupARCParamsAggInterval15m EmailSecurityTimeseriesGroupARCParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1h  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1d  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupARCParamsAggInterval1w  EmailSecurityTimeseriesGroupARCParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsAggInterval15m, EmailSecurityTimeseriesGroupARCParamsAggInterval1h, EmailSecurityTimeseriesGroupARCParamsAggInterval1d, EmailSecurityTimeseriesGroupARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsDKIM string

const (
	EmailSecurityTimeseriesGroupARCParamsDKIMPass EmailSecurityTimeseriesGroupARCParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupARCParamsDKIMNone EmailSecurityTimeseriesGroupARCParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupARCParamsDKIMFail EmailSecurityTimeseriesGroupARCParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsDKIMPass, EmailSecurityTimeseriesGroupARCParamsDKIMNone, EmailSecurityTimeseriesGroupARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsDMARC string

const (
	EmailSecurityTimeseriesGroupARCParamsDMARCPass EmailSecurityTimeseriesGroupARCParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupARCParamsDMARCNone EmailSecurityTimeseriesGroupARCParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupARCParamsDMARCFail EmailSecurityTimeseriesGroupARCParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsDMARCPass, EmailSecurityTimeseriesGroupARCParamsDMARCNone, EmailSecurityTimeseriesGroupARCParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupARCParamsFormat string

const (
	EmailSecurityTimeseriesGroupARCParamsFormatJson EmailSecurityTimeseriesGroupARCParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupARCParamsFormatCsv  EmailSecurityTimeseriesGroupARCParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsFormatJson, EmailSecurityTimeseriesGroupARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsSPF string

const (
	EmailSecurityTimeseriesGroupARCParamsSPFPass EmailSecurityTimeseriesGroupARCParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupARCParamsSPFNone EmailSecurityTimeseriesGroupARCParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupARCParamsSPFFail EmailSecurityTimeseriesGroupARCParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsSPFPass, EmailSecurityTimeseriesGroupARCParamsSPFNone, EmailSecurityTimeseriesGroupARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupARCResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupARCResponseEnvelope]
type emailSecurityTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDKIMParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsARC] `query:"arc"`
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
	DMARC param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupDKIMParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupDKIMParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDKIMParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1h  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1d  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupDKIMParamsAggInterval1w  EmailSecurityTimeseriesGroupDKIMParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsAggInterval15m, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1h, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1d, EmailSecurityTimeseriesGroupDKIMParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsARC string

const (
	EmailSecurityTimeseriesGroupDKIMParamsARCPass EmailSecurityTimeseriesGroupDKIMParamsARC = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsARCNone EmailSecurityTimeseriesGroupDKIMParamsARC = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsARCFail EmailSecurityTimeseriesGroupDKIMParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsARCPass, EmailSecurityTimeseriesGroupDKIMParamsARCNone, EmailSecurityTimeseriesGroupDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsDMARC string

const (
	EmailSecurityTimeseriesGroupDKIMParamsDMARCPass EmailSecurityTimeseriesGroupDKIMParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsDMARCNone EmailSecurityTimeseriesGroupDKIMParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsDMARCFail EmailSecurityTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsDMARCPass, EmailSecurityTimeseriesGroupDKIMParamsDMARCNone, EmailSecurityTimeseriesGroupDKIMParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupDKIMParamsFormat string

const (
	EmailSecurityTimeseriesGroupDKIMParamsFormatJson EmailSecurityTimeseriesGroupDKIMParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupDKIMParamsFormatCsv  EmailSecurityTimeseriesGroupDKIMParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsFormatJson, EmailSecurityTimeseriesGroupDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsSPF string

const (
	EmailSecurityTimeseriesGroupDKIMParamsSPFPass EmailSecurityTimeseriesGroupDKIMParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupDKIMParamsSPFNone EmailSecurityTimeseriesGroupDKIMParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupDKIMParamsSPFFail EmailSecurityTimeseriesGroupDKIMParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsSPFPass, EmailSecurityTimeseriesGroupDKIMParamsSPFNone, EmailSecurityTimeseriesGroupDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupDKIMParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupDKIMParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupDKIMParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDKIMResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDKIMResponseEnvelope]
type emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupDMARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupDMARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupDMARCParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupDMARCParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1h  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1d  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupDMARCParamsAggInterval1w  EmailSecurityTimeseriesGroupDMARCParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsAggInterval15m, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1h, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1d, EmailSecurityTimeseriesGroupDMARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsARC string

const (
	EmailSecurityTimeseriesGroupDMARCParamsARCPass EmailSecurityTimeseriesGroupDMARCParamsARC = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsARCNone EmailSecurityTimeseriesGroupDMARCParamsARC = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsARCFail EmailSecurityTimeseriesGroupDMARCParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsARCPass, EmailSecurityTimeseriesGroupDMARCParamsARCNone, EmailSecurityTimeseriesGroupDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsDKIM string

const (
	EmailSecurityTimeseriesGroupDMARCParamsDKIMPass EmailSecurityTimeseriesGroupDMARCParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsDKIMNone EmailSecurityTimeseriesGroupDMARCParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsDKIMFail EmailSecurityTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsDKIMPass, EmailSecurityTimeseriesGroupDMARCParamsDKIMNone, EmailSecurityTimeseriesGroupDMARCParamsDKIMFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupDMARCParamsFormat string

const (
	EmailSecurityTimeseriesGroupDMARCParamsFormatJson EmailSecurityTimeseriesGroupDMARCParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupDMARCParamsFormatCsv  EmailSecurityTimeseriesGroupDMARCParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsFormatJson, EmailSecurityTimeseriesGroupDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsSPF string

const (
	EmailSecurityTimeseriesGroupDMARCParamsSPFPass EmailSecurityTimeseriesGroupDMARCParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupDMARCParamsSPFNone EmailSecurityTimeseriesGroupDMARCParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupDMARCParamsSPFFail EmailSecurityTimeseriesGroupDMARCParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsSPFPass, EmailSecurityTimeseriesGroupDMARCParamsSPFNone, EmailSecurityTimeseriesGroupDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupDMARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupDMARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupDMARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupDMARCResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupDMARCResponseEnvelope]
type emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupMaliciousParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupMaliciousParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupMaliciousParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupMaliciousParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupMaliciousParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1d  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1w  EmailSecurityTimeseriesGroupMaliciousParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsAggInterval15m, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1h, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1d, EmailSecurityTimeseriesGroupMaliciousParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsARC string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsARCPass EmailSecurityTimeseriesGroupMaliciousParamsARC = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsARCNone EmailSecurityTimeseriesGroupMaliciousParamsARC = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsARCFail EmailSecurityTimeseriesGroupMaliciousParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsARCPass, EmailSecurityTimeseriesGroupMaliciousParamsARCNone, EmailSecurityTimeseriesGroupMaliciousParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsDKIM string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail EmailSecurityTimeseriesGroupMaliciousParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsDKIMPass, EmailSecurityTimeseriesGroupMaliciousParamsDKIMNone, EmailSecurityTimeseriesGroupMaliciousParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsDMARC string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail EmailSecurityTimeseriesGroupMaliciousParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsDMARCPass, EmailSecurityTimeseriesGroupMaliciousParamsDMARCNone, EmailSecurityTimeseriesGroupMaliciousParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupMaliciousParamsFormat string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsFormatJson EmailSecurityTimeseriesGroupMaliciousParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupMaliciousParamsFormatCsv  EmailSecurityTimeseriesGroupMaliciousParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsFormatJson, EmailSecurityTimeseriesGroupMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsSPF string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsSPFPass EmailSecurityTimeseriesGroupMaliciousParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupMaliciousParamsSPFNone EmailSecurityTimeseriesGroupMaliciousParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupMaliciousParamsSPFFail EmailSecurityTimeseriesGroupMaliciousParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsSPFPass, EmailSecurityTimeseriesGroupMaliciousParamsSPFNone, EmailSecurityTimeseriesGroupMaliciousParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupMaliciousParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupMaliciousParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupMaliciousResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupMaliciousResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupMaliciousResponseEnvelope]
type emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupMaliciousResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupMaliciousResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpamParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSpamParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupSpamParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSpamParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSpamParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupSpamParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupSpamParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSpamParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSpamParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpamParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSpamParamsAggInterval15m EmailSecurityTimeseriesGroupSpamParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1h  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1d  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSpamParamsAggInterval1w  EmailSecurityTimeseriesGroupSpamParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSpamParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsAggInterval15m, EmailSecurityTimeseriesGroupSpamParamsAggInterval1h, EmailSecurityTimeseriesGroupSpamParamsAggInterval1d, EmailSecurityTimeseriesGroupSpamParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsARC string

const (
	EmailSecurityTimeseriesGroupSpamParamsARCPass EmailSecurityTimeseriesGroupSpamParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsARCNone EmailSecurityTimeseriesGroupSpamParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsARCFail EmailSecurityTimeseriesGroupSpamParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsARCPass, EmailSecurityTimeseriesGroupSpamParamsARCNone, EmailSecurityTimeseriesGroupSpamParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSpamParamsDKIMPass EmailSecurityTimeseriesGroupSpamParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsDKIMNone EmailSecurityTimeseriesGroupSpamParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsDKIMFail EmailSecurityTimeseriesGroupSpamParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsDKIMPass, EmailSecurityTimeseriesGroupSpamParamsDKIMNone, EmailSecurityTimeseriesGroupSpamParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSpamParamsDMARCPass EmailSecurityTimeseriesGroupSpamParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsDMARCNone EmailSecurityTimeseriesGroupSpamParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsDMARCFail EmailSecurityTimeseriesGroupSpamParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsDMARCPass, EmailSecurityTimeseriesGroupSpamParamsDMARCNone, EmailSecurityTimeseriesGroupSpamParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupSpamParamsFormat string

const (
	EmailSecurityTimeseriesGroupSpamParamsFormatJson EmailSecurityTimeseriesGroupSpamParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSpamParamsFormatCsv  EmailSecurityTimeseriesGroupSpamParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSpamParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsFormatJson, EmailSecurityTimeseriesGroupSpamParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsSPF string

const (
	EmailSecurityTimeseriesGroupSpamParamsSPFPass EmailSecurityTimeseriesGroupSpamParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupSpamParamsSPFNone EmailSecurityTimeseriesGroupSpamParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupSpamParamsSPFFail EmailSecurityTimeseriesGroupSpamParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpamParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsSPFPass, EmailSecurityTimeseriesGroupSpamParamsSPFNone, EmailSecurityTimeseriesGroupSpamParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSpamParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSpamParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSpamParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpamResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSpamResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpamResponseEnvelope]
type emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpamResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpamResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSPFParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSPFParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSPFParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSPFParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSPFParamsAggInterval15m EmailSecurityTimeseriesGroupSPFParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1h  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1d  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSPFParamsAggInterval1w  EmailSecurityTimeseriesGroupSPFParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSPFParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsAggInterval15m, EmailSecurityTimeseriesGroupSPFParamsAggInterval1h, EmailSecurityTimeseriesGroupSPFParamsAggInterval1d, EmailSecurityTimeseriesGroupSPFParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsARC string

const (
	EmailSecurityTimeseriesGroupSPFParamsARCPass EmailSecurityTimeseriesGroupSPFParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsARCNone EmailSecurityTimeseriesGroupSPFParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsARCFail EmailSecurityTimeseriesGroupSPFParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsARCPass, EmailSecurityTimeseriesGroupSPFParamsARCNone, EmailSecurityTimeseriesGroupSPFParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSPFParamsDKIMPass EmailSecurityTimeseriesGroupSPFParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsDKIMNone EmailSecurityTimeseriesGroupSPFParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsDKIMFail EmailSecurityTimeseriesGroupSPFParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsDKIMPass, EmailSecurityTimeseriesGroupSPFParamsDKIMNone, EmailSecurityTimeseriesGroupSPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSPFParamsDMARCPass EmailSecurityTimeseriesGroupSPFParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSPFParamsDMARCNone EmailSecurityTimeseriesGroupSPFParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSPFParamsDMARCFail EmailSecurityTimeseriesGroupSPFParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsDMARCPass, EmailSecurityTimeseriesGroupSPFParamsDMARCNone, EmailSecurityTimeseriesGroupSPFParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupSPFParamsFormat string

const (
	EmailSecurityTimeseriesGroupSPFParamsFormatJson EmailSecurityTimeseriesGroupSPFParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSPFParamsFormatCsv  EmailSecurityTimeseriesGroupSPFParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsFormatJson, EmailSecurityTimeseriesGroupSPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSPFParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSPFParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSPFParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSPFResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSPFResponseEnvelope]
type emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupSpoofParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupSpoofParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupSpoofParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupSpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupSpoofParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTimeseriesGroupSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupSpoofParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1h  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1d  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupSpoofParamsAggInterval1w  EmailSecurityTimeseriesGroupSpoofParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsAggInterval15m, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1h, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1d, EmailSecurityTimeseriesGroupSpoofParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsARC string

const (
	EmailSecurityTimeseriesGroupSpoofParamsARCPass EmailSecurityTimeseriesGroupSpoofParamsARC = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsARCNone EmailSecurityTimeseriesGroupSpoofParamsARC = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsARCFail EmailSecurityTimeseriesGroupSpoofParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsARCPass, EmailSecurityTimeseriesGroupSpoofParamsARCNone, EmailSecurityTimeseriesGroupSpoofParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsDKIM string

const (
	EmailSecurityTimeseriesGroupSpoofParamsDKIMPass EmailSecurityTimeseriesGroupSpoofParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsDKIMNone EmailSecurityTimeseriesGroupSpoofParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsDKIMFail EmailSecurityTimeseriesGroupSpoofParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsDKIMPass, EmailSecurityTimeseriesGroupSpoofParamsDKIMNone, EmailSecurityTimeseriesGroupSpoofParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsDMARC string

const (
	EmailSecurityTimeseriesGroupSpoofParamsDMARCPass EmailSecurityTimeseriesGroupSpoofParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsDMARCNone EmailSecurityTimeseriesGroupSpoofParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsDMARCFail EmailSecurityTimeseriesGroupSpoofParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsDMARCPass, EmailSecurityTimeseriesGroupSpoofParamsDMARCNone, EmailSecurityTimeseriesGroupSpoofParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupSpoofParamsFormat string

const (
	EmailSecurityTimeseriesGroupSpoofParamsFormatJson EmailSecurityTimeseriesGroupSpoofParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupSpoofParamsFormatCsv  EmailSecurityTimeseriesGroupSpoofParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsFormatJson, EmailSecurityTimeseriesGroupSpoofParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsSPF string

const (
	EmailSecurityTimeseriesGroupSpoofParamsSPFPass EmailSecurityTimeseriesGroupSpoofParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupSpoofParamsSPFNone EmailSecurityTimeseriesGroupSpoofParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupSpoofParamsSPFFail EmailSecurityTimeseriesGroupSpoofParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsSPFPass, EmailSecurityTimeseriesGroupSpoofParamsSPFNone, EmailSecurityTimeseriesGroupSpoofParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupSpoofParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupSpoofParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupSpoofParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupSpoofResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupSpoofResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTimeseriesGroupSpoofResponseEnvelope]
type emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupSpoofResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupSpoofResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupThreatCategoryParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupThreatCategoryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1d  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1w  EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval15m, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1h, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1d, EmailSecurityTimeseriesGroupThreatCategoryParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsARC string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail EmailSecurityTimeseriesGroupThreatCategoryParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsARCPass, EmailSecurityTimeseriesGroupThreatCategoryParamsARCNone, EmailSecurityTimeseriesGroupThreatCategoryParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMPass, EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMNone, EmailSecurityTimeseriesGroupThreatCategoryParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCPass, EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCNone, EmailSecurityTimeseriesGroupThreatCategoryParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupThreatCategoryParamsFormat string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson EmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupThreatCategoryParamsFormatCsv  EmailSecurityTimeseriesGroupThreatCategoryParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsFormatJson, EmailSecurityTimeseriesGroupThreatCategoryParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsSPF string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail EmailSecurityTimeseriesGroupThreatCategoryParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsSPFPass, EmailSecurityTimeseriesGroupThreatCategoryParamsSPFNone, EmailSecurityTimeseriesGroupThreatCategoryParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion string

const (
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_0"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_1"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_2"
	EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_3 EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_0, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_1, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_2, EmailSecurityTimeseriesGroupThreatCategoryParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupThreatCategoryResponse             `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope]
type emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupThreatCategoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupThreatCategoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTimeseriesGroupTLSVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTimeseriesGroupTLSVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTimeseriesGroupTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailSecurityTimeseriesGroupTLSVersionParams]'s query
// parameters as `url.Values`.
func (r EmailSecurityTimeseriesGroupTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "15m"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1h"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1d  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1d"
	EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1w  EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval = "1w"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval15m, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1h, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1d, EmailSecurityTimeseriesGroupTLSVersionParamsAggInterval1w:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsARC string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsARCPass EmailSecurityTimeseriesGroupTLSVersionParamsARC = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsARCNone EmailSecurityTimeseriesGroupTLSVersionParamsARC = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsARCFail EmailSecurityTimeseriesGroupTLSVersionParamsARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsARCPass, EmailSecurityTimeseriesGroupTLSVersionParamsARCNone, EmailSecurityTimeseriesGroupTLSVersionParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsDKIM string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail EmailSecurityTimeseriesGroupTLSVersionParamsDKIM = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsDKIMPass, EmailSecurityTimeseriesGroupTLSVersionParamsDKIMNone, EmailSecurityTimeseriesGroupTLSVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsDMARC string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail EmailSecurityTimeseriesGroupTLSVersionParamsDMARC = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsDMARCPass, EmailSecurityTimeseriesGroupTLSVersionParamsDMARCNone, EmailSecurityTimeseriesGroupTLSVersionParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTimeseriesGroupTLSVersionParamsFormat string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson EmailSecurityTimeseriesGroupTLSVersionParamsFormat = "JSON"
	EmailSecurityTimeseriesGroupTLSVersionParamsFormatCsv  EmailSecurityTimeseriesGroupTLSVersionParamsFormat = "CSV"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsFormatJson, EmailSecurityTimeseriesGroupTLSVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionParamsSPF string

const (
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "PASS"
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "NONE"
	EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail EmailSecurityTimeseriesGroupTLSVersionParamsSPF = "FAIL"
)

func (r EmailSecurityTimeseriesGroupTLSVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTimeseriesGroupTLSVersionParamsSPFPass, EmailSecurityTimeseriesGroupTLSVersionParamsSPFNone, EmailSecurityTimeseriesGroupTLSVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope struct {
	Result  EmailSecurityTimeseriesGroupTLSVersionResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope]
type emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTimeseriesGroupTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTimeseriesGroupTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
