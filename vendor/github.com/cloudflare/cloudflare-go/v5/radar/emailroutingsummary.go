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

// EmailRoutingSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingSummaryService] method instead.
type EmailRoutingSummaryService struct {
	Options []option.RequestOption
}

// NewEmailRoutingSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingSummaryService(opts ...option.RequestOption) (r *EmailRoutingSummaryService) {
	r = &EmailRoutingSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation.
func (r *EmailRoutingSummaryService) ARC(ctx context.Context, query EmailRoutingSummaryARCParams, opts ...option.RequestOption) (res *EmailRoutingSummaryARCResponse, err error) {
	var env EmailRoutingSummaryARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation.
func (r *EmailRoutingSummaryService) DKIM(ctx context.Context, query EmailRoutingSummaryDKIMParams, opts ...option.RequestOption) (res *EmailRoutingSummaryDKIMResponse, err error) {
	var env EmailRoutingSummaryDKIMResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation.
func (r *EmailRoutingSummaryService) DMARC(ctx context.Context, query EmailRoutingSummaryDMARCParams, opts ...option.RequestOption) (res *EmailRoutingSummaryDMARCResponse, err error) {
	var env EmailRoutingSummaryDMARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted).
func (r *EmailRoutingSummaryService) Encrypted(ctx context.Context, query EmailRoutingSummaryEncryptedParams, opts ...option.RequestOption) (res *EmailRoutingSummaryEncryptedResponse, err error) {
	var env EmailRoutingSummaryEncryptedResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by IP version.
func (r *EmailRoutingSummaryService) IPVersion(ctx context.Context, query EmailRoutingSummaryIPVersionParams, opts ...option.RequestOption) (res *EmailRoutingSummaryIPVersionResponse, err error) {
	var env EmailRoutingSummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework)
// validation.
func (r *EmailRoutingSummaryService) SPF(ctx context.Context, query EmailRoutingSummarySPFParams, opts ...option.RequestOption) (res *EmailRoutingSummarySPFResponse, err error) {
	var env EmailRoutingSummarySPFResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingSummaryARCResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                  `json:"summary_0,required"`
	JSON     emailRoutingSummaryARCResponseJSON `json:"-"`
}

// emailRoutingSummaryARCResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryARCResponse]
type emailRoutingSummaryARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryARCResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryARCResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryARCResponseMetaJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryARCResponseMeta]
type emailRoutingSummaryARCResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  emailRoutingSummaryARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryARCResponseMetaConfidenceInfo]
type emailRoutingSummaryARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryARCResponseMetaDateRange]
type emailRoutingSummaryARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryARCResponseMetaNormalization string

const (
	EmailRoutingSummaryARCResponseMetaNormalizationPercentage           EmailRoutingSummaryARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryARCResponseMetaNormalizationMin0Max              EmailRoutingSummaryARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryARCResponseMetaNormalizationMinMax               EmailRoutingSummaryARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryARCResponseMetaNormalizationRawValues            EmailRoutingSummaryARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryARCResponseMetaNormalizationPercentageChange     EmailRoutingSummaryARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryARCResponseMetaNormalizationRollingAverage       EmailRoutingSummaryARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryARCResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryARCResponseMetaNormalizationRatio                EmailRoutingSummaryARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCResponseMetaNormalizationPercentage, EmailRoutingSummaryARCResponseMetaNormalizationMin0Max, EmailRoutingSummaryARCResponseMetaNormalizationMinMax, EmailRoutingSummaryARCResponseMetaNormalizationRawValues, EmailRoutingSummaryARCResponseMetaNormalizationPercentageChange, EmailRoutingSummaryARCResponseMetaNormalizationRollingAverage, EmailRoutingSummaryARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryARCResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  emailRoutingSummaryARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryARCResponseMetaUnit]
type emailRoutingSummaryARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryDKIMResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                   `json:"summary_0,required"`
	JSON     emailRoutingSummaryDKIMResponseJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryDKIMResponse]
type emailRoutingSummaryDKIMResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryDKIMResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryDKIMResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryDKIMResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryDKIMResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDKIMResponseMeta]
type emailRoutingSummaryDKIMResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryDKIMResponseMetaConfidenceInfo]
type emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryDKIMResponseMetaDateRange]
type emailRoutingSummaryDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryDKIMResponseMetaNormalization string

const (
	EmailRoutingSummaryDKIMResponseMetaNormalizationPercentage           EmailRoutingSummaryDKIMResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryDKIMResponseMetaNormalizationMin0Max              EmailRoutingSummaryDKIMResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryDKIMResponseMetaNormalizationMinMax               EmailRoutingSummaryDKIMResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryDKIMResponseMetaNormalizationRawValues            EmailRoutingSummaryDKIMResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryDKIMResponseMetaNormalizationPercentageChange     EmailRoutingSummaryDKIMResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryDKIMResponseMetaNormalizationRollingAverage       EmailRoutingSummaryDKIMResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryDKIMResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryDKIMResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryDKIMResponseMetaNormalizationRatio                EmailRoutingSummaryDKIMResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryDKIMResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMResponseMetaNormalizationPercentage, EmailRoutingSummaryDKIMResponseMetaNormalizationMin0Max, EmailRoutingSummaryDKIMResponseMetaNormalizationMinMax, EmailRoutingSummaryDKIMResponseMetaNormalizationRawValues, EmailRoutingSummaryDKIMResponseMetaNormalizationPercentageChange, EmailRoutingSummaryDKIMResponseMetaNormalizationRollingAverage, EmailRoutingSummaryDKIMResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryDKIMResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  emailRoutingSummaryDKIMResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDKIMResponseMetaUnit]
type emailRoutingSummaryDKIMResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryDMARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                    `json:"summary_0,required"`
	JSON     emailRoutingSummaryDMARCResponseJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryDMARCResponse]
type emailRoutingSummaryDMARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryDMARCResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryDMARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryDMARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryDMARCResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDMARCResponseMeta]
type emailRoutingSummaryDMARCResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryDMARCResponseMetaConfidenceInfo]
type emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryDMARCResponseMetaDateRange]
type emailRoutingSummaryDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryDMARCResponseMetaNormalization string

const (
	EmailRoutingSummaryDMARCResponseMetaNormalizationPercentage           EmailRoutingSummaryDMARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryDMARCResponseMetaNormalizationMin0Max              EmailRoutingSummaryDMARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryDMARCResponseMetaNormalizationMinMax               EmailRoutingSummaryDMARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryDMARCResponseMetaNormalizationRawValues            EmailRoutingSummaryDMARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryDMARCResponseMetaNormalizationPercentageChange     EmailRoutingSummaryDMARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryDMARCResponseMetaNormalizationRollingAverage       EmailRoutingSummaryDMARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryDMARCResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryDMARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryDMARCResponseMetaNormalizationRatio                EmailRoutingSummaryDMARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryDMARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCResponseMetaNormalizationPercentage, EmailRoutingSummaryDMARCResponseMetaNormalizationMin0Max, EmailRoutingSummaryDMARCResponseMetaNormalizationMinMax, EmailRoutingSummaryDMARCResponseMetaNormalizationRawValues, EmailRoutingSummaryDMARCResponseMetaNormalizationPercentageChange, EmailRoutingSummaryDMARCResponseMetaNormalizationRollingAverage, EmailRoutingSummaryDMARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryDMARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  emailRoutingSummaryDMARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDMARCResponseMetaUnit]
type emailRoutingSummaryDMARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryEncryptedResponseMeta     `json:"meta,required"`
	Summary0 EmailRoutingSummaryEncryptedResponseSummary0 `json:"summary_0,required"`
	JSON     emailRoutingSummaryEncryptedResponseJSON     `json:"-"`
}

// emailRoutingSummaryEncryptedResponseJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryEncryptedResponse]
type emailRoutingSummaryEncryptedResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryEncryptedResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryEncryptedResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryEncryptedResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryEncryptedResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryEncryptedResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryEncryptedResponseMeta]
type emailRoutingSummaryEncryptedResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo]
type emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryEncryptedResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryEncryptedResponseMetaDateRange]
type emailRoutingSummaryEncryptedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryEncryptedResponseMetaNormalization string

const (
	EmailRoutingSummaryEncryptedResponseMetaNormalizationPercentage           EmailRoutingSummaryEncryptedResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationMin0Max              EmailRoutingSummaryEncryptedResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationMinMax               EmailRoutingSummaryEncryptedResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationRawValues            EmailRoutingSummaryEncryptedResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationPercentageChange     EmailRoutingSummaryEncryptedResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationRollingAverage       EmailRoutingSummaryEncryptedResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryEncryptedResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryEncryptedResponseMetaNormalizationRatio                EmailRoutingSummaryEncryptedResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryEncryptedResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedResponseMetaNormalizationPercentage, EmailRoutingSummaryEncryptedResponseMetaNormalizationMin0Max, EmailRoutingSummaryEncryptedResponseMetaNormalizationMinMax, EmailRoutingSummaryEncryptedResponseMetaNormalizationRawValues, EmailRoutingSummaryEncryptedResponseMetaNormalizationPercentageChange, EmailRoutingSummaryEncryptedResponseMetaNormalizationRollingAverage, EmailRoutingSummaryEncryptedResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryEncryptedResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  emailRoutingSummaryEncryptedResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryEncryptedResponseMetaUnit]
type emailRoutingSummaryEncryptedResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseSummary0 struct {
	// A numeric string.
	Encrypted string `json:"ENCRYPTED,required"`
	// A numeric string.
	NotEncrypted string                                           `json:"NOT_ENCRYPTED,required"`
	JSON         emailRoutingSummaryEncryptedResponseSummary0JSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseSummary0JSON contains the JSON metadata for
// the struct [EmailRoutingSummaryEncryptedResponseSummary0]
type emailRoutingSummaryEncryptedResponseSummary0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 EmailRoutingSummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     emailRoutingSummaryIPVersionResponseJSON     `json:"-"`
}

// emailRoutingSummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryIPVersionResponse]
type emailRoutingSummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummaryIPVersionResponseMeta struct {
	ConfidenceInfo EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummaryIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummaryIPVersionResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummaryIPVersionResponseMetaJSON   `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryIPVersionResponseMeta]
type emailRoutingSummaryIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo]
type emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryIPVersionResponseMetaDateRange]
type emailRoutingSummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummaryIPVersionResponseMetaNormalization string

const (
	EmailRoutingSummaryIPVersionResponseMetaNormalizationPercentage           EmailRoutingSummaryIPVersionResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationMin0Max              EmailRoutingSummaryIPVersionResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationMinMax               EmailRoutingSummaryIPVersionResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationRawValues            EmailRoutingSummaryIPVersionResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationPercentageChange     EmailRoutingSummaryIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationRollingAverage       EmailRoutingSummaryIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationOverlappedPercentage EmailRoutingSummaryIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummaryIPVersionResponseMetaNormalizationRatio                EmailRoutingSummaryIPVersionResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummaryIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionResponseMetaNormalizationPercentage, EmailRoutingSummaryIPVersionResponseMetaNormalizationMin0Max, EmailRoutingSummaryIPVersionResponseMetaNormalizationMinMax, EmailRoutingSummaryIPVersionResponseMetaNormalizationRawValues, EmailRoutingSummaryIPVersionResponseMetaNormalizationPercentageChange, EmailRoutingSummaryIPVersionResponseMetaNormalizationRollingAverage, EmailRoutingSummaryIPVersionResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummaryIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionResponseMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  emailRoutingSummaryIPVersionResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryIPVersionResponseMetaUnit]
type emailRoutingSummaryIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseSummary0 struct {
	// A numeric string.
	IPv4 string `json:"IPv4,required"`
	// A numeric string.
	IPv6 string                                           `json:"IPv6,required"`
	JSON emailRoutingSummaryIPVersionResponseSummary0JSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [EmailRoutingSummaryIPVersionResponseSummary0]
type emailRoutingSummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponse struct {
	// Metadata for the results.
	Meta     EmailRoutingSummarySPFResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                  `json:"summary_0,required"`
	JSON     emailRoutingSummarySPFResponseJSON `json:"-"`
}

// emailRoutingSummarySPFResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummarySPFResponse]
type emailRoutingSummarySPFResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingSummarySPFResponseMeta struct {
	ConfidenceInfo EmailRoutingSummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingSummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingSummarySPFResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingSummarySPFResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingSummarySPFResponseMetaJSON   `json:"-"`
}

// emailRoutingSummarySPFResponseMetaJSON contains the JSON metadata for the struct
// [EmailRoutingSummarySPFResponseMeta]
type emailRoutingSummarySPFResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  emailRoutingSummarySPFResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummarySPFResponseMetaConfidenceInfo]
type emailRoutingSummarySPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummarySPFResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummarySPFResponseMetaDateRange]
type emailRoutingSummarySPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingSummarySPFResponseMetaNormalization string

const (
	EmailRoutingSummarySPFResponseMetaNormalizationPercentage           EmailRoutingSummarySPFResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingSummarySPFResponseMetaNormalizationMin0Max              EmailRoutingSummarySPFResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingSummarySPFResponseMetaNormalizationMinMax               EmailRoutingSummarySPFResponseMetaNormalization = "MIN_MAX"
	EmailRoutingSummarySPFResponseMetaNormalizationRawValues            EmailRoutingSummarySPFResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingSummarySPFResponseMetaNormalizationPercentageChange     EmailRoutingSummarySPFResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingSummarySPFResponseMetaNormalizationRollingAverage       EmailRoutingSummarySPFResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingSummarySPFResponseMetaNormalizationOverlappedPercentage EmailRoutingSummarySPFResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingSummarySPFResponseMetaNormalizationRatio                EmailRoutingSummarySPFResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingSummarySPFResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFResponseMetaNormalizationPercentage, EmailRoutingSummarySPFResponseMetaNormalizationMin0Max, EmailRoutingSummarySPFResponseMetaNormalizationMinMax, EmailRoutingSummarySPFResponseMetaNormalizationRawValues, EmailRoutingSummarySPFResponseMetaNormalizationPercentageChange, EmailRoutingSummarySPFResponseMetaNormalizationRollingAverage, EmailRoutingSummarySPFResponseMetaNormalizationOverlappedPercentage, EmailRoutingSummarySPFResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingSummarySPFResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  emailRoutingSummarySPFResponseMetaUnitJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailRoutingSummarySPFResponseMetaUnit]
type emailRoutingSummarySPFResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummaryARCParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingSummaryARCParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummaryARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummaryARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryARCParamsDKIM string

const (
	EmailRoutingSummaryARCParamsDKIMPass EmailRoutingSummaryARCParamsDKIM = "PASS"
	EmailRoutingSummaryARCParamsDKIMNone EmailRoutingSummaryARCParamsDKIM = "NONE"
	EmailRoutingSummaryARCParamsDKIMFail EmailRoutingSummaryARCParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsDKIMPass, EmailRoutingSummaryARCParamsDKIMNone, EmailRoutingSummaryARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsDMARC string

const (
	EmailRoutingSummaryARCParamsDMARCPass EmailRoutingSummaryARCParamsDMARC = "PASS"
	EmailRoutingSummaryARCParamsDMARCNone EmailRoutingSummaryARCParamsDMARC = "NONE"
	EmailRoutingSummaryARCParamsDMARCFail EmailRoutingSummaryARCParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsDMARCPass, EmailRoutingSummaryARCParamsDMARCNone, EmailRoutingSummaryARCParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsEncrypted string

const (
	EmailRoutingSummaryARCParamsEncryptedEncrypted    EmailRoutingSummaryARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryARCParamsEncryptedNotEncrypted EmailRoutingSummaryARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsEncryptedEncrypted, EmailRoutingSummaryARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryARCParamsFormat string

const (
	EmailRoutingSummaryARCParamsFormatJson EmailRoutingSummaryARCParamsFormat = "JSON"
	EmailRoutingSummaryARCParamsFormatCsv  EmailRoutingSummaryARCParamsFormat = "CSV"
)

func (r EmailRoutingSummaryARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsFormatJson, EmailRoutingSummaryARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsIPVersion string

const (
	EmailRoutingSummaryARCParamsIPVersionIPv4 EmailRoutingSummaryARCParamsIPVersion = "IPv4"
	EmailRoutingSummaryARCParamsIPVersionIPv6 EmailRoutingSummaryARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsIPVersionIPv4, EmailRoutingSummaryARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsSPF string

const (
	EmailRoutingSummaryARCParamsSPFPass EmailRoutingSummaryARCParamsSPF = "PASS"
	EmailRoutingSummaryARCParamsSPFNone EmailRoutingSummaryARCParamsSPF = "NONE"
	EmailRoutingSummaryARCParamsSPFFail EmailRoutingSummaryARCParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsSPFPass, EmailRoutingSummaryARCParamsSPFNone, EmailRoutingSummaryARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCResponseEnvelope struct {
	Result  EmailRoutingSummaryARCResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailRoutingSummaryARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryARCResponseEnvelope]
type emailRoutingSummaryARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummaryDKIMParamsARC] `query:"arc"`
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
	DMARC param.Field[[]EmailRoutingSummaryDKIMParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummaryDKIMParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryDKIMParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummaryDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryDKIMParamsARC string

const (
	EmailRoutingSummaryDKIMParamsARCPass EmailRoutingSummaryDKIMParamsARC = "PASS"
	EmailRoutingSummaryDKIMParamsARCNone EmailRoutingSummaryDKIMParamsARC = "NONE"
	EmailRoutingSummaryDKIMParamsARCFail EmailRoutingSummaryDKIMParamsARC = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsARCPass, EmailRoutingSummaryDKIMParamsARCNone, EmailRoutingSummaryDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsDMARC string

const (
	EmailRoutingSummaryDKIMParamsDMARCPass EmailRoutingSummaryDKIMParamsDMARC = "PASS"
	EmailRoutingSummaryDKIMParamsDMARCNone EmailRoutingSummaryDKIMParamsDMARC = "NONE"
	EmailRoutingSummaryDKIMParamsDMARCFail EmailRoutingSummaryDKIMParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsDMARCPass, EmailRoutingSummaryDKIMParamsDMARCNone, EmailRoutingSummaryDKIMParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsEncrypted string

const (
	EmailRoutingSummaryDKIMParamsEncryptedEncrypted    EmailRoutingSummaryDKIMParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted EmailRoutingSummaryDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryDKIMParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsEncryptedEncrypted, EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryDKIMParamsFormat string

const (
	EmailRoutingSummaryDKIMParamsFormatJson EmailRoutingSummaryDKIMParamsFormat = "JSON"
	EmailRoutingSummaryDKIMParamsFormatCsv  EmailRoutingSummaryDKIMParamsFormat = "CSV"
)

func (r EmailRoutingSummaryDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsFormatJson, EmailRoutingSummaryDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsIPVersion string

const (
	EmailRoutingSummaryDKIMParamsIPVersionIPv4 EmailRoutingSummaryDKIMParamsIPVersion = "IPv4"
	EmailRoutingSummaryDKIMParamsIPVersionIPv6 EmailRoutingSummaryDKIMParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryDKIMParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsIPVersionIPv4, EmailRoutingSummaryDKIMParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsSPF string

const (
	EmailRoutingSummaryDKIMParamsSPFPass EmailRoutingSummaryDKIMParamsSPF = "PASS"
	EmailRoutingSummaryDKIMParamsSPFNone EmailRoutingSummaryDKIMParamsSPF = "NONE"
	EmailRoutingSummaryDKIMParamsSPFFail EmailRoutingSummaryDKIMParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsSPFPass, EmailRoutingSummaryDKIMParamsSPFNone, EmailRoutingSummaryDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMResponseEnvelope struct {
	Result  EmailRoutingSummaryDKIMResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    emailRoutingSummaryDKIMResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDKIMResponseEnvelope]
type emailRoutingSummaryDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummaryDMARCParamsDKIM] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummaryDMARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryDMARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummaryDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryDMARCParamsARC string

const (
	EmailRoutingSummaryDMARCParamsARCPass EmailRoutingSummaryDMARCParamsARC = "PASS"
	EmailRoutingSummaryDMARCParamsARCNone EmailRoutingSummaryDMARCParamsARC = "NONE"
	EmailRoutingSummaryDMARCParamsARCFail EmailRoutingSummaryDMARCParamsARC = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsARCPass, EmailRoutingSummaryDMARCParamsARCNone, EmailRoutingSummaryDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsDKIM string

const (
	EmailRoutingSummaryDMARCParamsDKIMPass EmailRoutingSummaryDMARCParamsDKIM = "PASS"
	EmailRoutingSummaryDMARCParamsDKIMNone EmailRoutingSummaryDMARCParamsDKIM = "NONE"
	EmailRoutingSummaryDMARCParamsDKIMFail EmailRoutingSummaryDMARCParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsDKIMPass, EmailRoutingSummaryDMARCParamsDKIMNone, EmailRoutingSummaryDMARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsEncrypted string

const (
	EmailRoutingSummaryDMARCParamsEncryptedEncrypted    EmailRoutingSummaryDMARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted EmailRoutingSummaryDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryDMARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsEncryptedEncrypted, EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryDMARCParamsFormat string

const (
	EmailRoutingSummaryDMARCParamsFormatJson EmailRoutingSummaryDMARCParamsFormat = "JSON"
	EmailRoutingSummaryDMARCParamsFormatCsv  EmailRoutingSummaryDMARCParamsFormat = "CSV"
)

func (r EmailRoutingSummaryDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsFormatJson, EmailRoutingSummaryDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsIPVersion string

const (
	EmailRoutingSummaryDMARCParamsIPVersionIPv4 EmailRoutingSummaryDMARCParamsIPVersion = "IPv4"
	EmailRoutingSummaryDMARCParamsIPVersionIPv6 EmailRoutingSummaryDMARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryDMARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsIPVersionIPv4, EmailRoutingSummaryDMARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsSPF string

const (
	EmailRoutingSummaryDMARCParamsSPFPass EmailRoutingSummaryDMARCParamsSPF = "PASS"
	EmailRoutingSummaryDMARCParamsSPFNone EmailRoutingSummaryDMARCParamsSPF = "NONE"
	EmailRoutingSummaryDMARCParamsSPFFail EmailRoutingSummaryDMARCParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsSPFPass, EmailRoutingSummaryDMARCParamsSPFNone, EmailRoutingSummaryDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCResponseEnvelope struct {
	Result  EmailRoutingSummaryDMARCResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    emailRoutingSummaryDMARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDMARCResponseEnvelope]
type emailRoutingSummaryDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummaryEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummaryEncryptedParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingSummaryEncryptedParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummaryEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryEncryptedParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryEncryptedParamsARC string

const (
	EmailRoutingSummaryEncryptedParamsARCPass EmailRoutingSummaryEncryptedParamsARC = "PASS"
	EmailRoutingSummaryEncryptedParamsARCNone EmailRoutingSummaryEncryptedParamsARC = "NONE"
	EmailRoutingSummaryEncryptedParamsARCFail EmailRoutingSummaryEncryptedParamsARC = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsARCPass, EmailRoutingSummaryEncryptedParamsARCNone, EmailRoutingSummaryEncryptedParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsDKIM string

const (
	EmailRoutingSummaryEncryptedParamsDKIMPass EmailRoutingSummaryEncryptedParamsDKIM = "PASS"
	EmailRoutingSummaryEncryptedParamsDKIMNone EmailRoutingSummaryEncryptedParamsDKIM = "NONE"
	EmailRoutingSummaryEncryptedParamsDKIMFail EmailRoutingSummaryEncryptedParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsDKIMPass, EmailRoutingSummaryEncryptedParamsDKIMNone, EmailRoutingSummaryEncryptedParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsDMARC string

const (
	EmailRoutingSummaryEncryptedParamsDMARCPass EmailRoutingSummaryEncryptedParamsDMARC = "PASS"
	EmailRoutingSummaryEncryptedParamsDMARCNone EmailRoutingSummaryEncryptedParamsDMARC = "NONE"
	EmailRoutingSummaryEncryptedParamsDMARCFail EmailRoutingSummaryEncryptedParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsDMARCPass, EmailRoutingSummaryEncryptedParamsDMARCNone, EmailRoutingSummaryEncryptedParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryEncryptedParamsFormat string

const (
	EmailRoutingSummaryEncryptedParamsFormatJson EmailRoutingSummaryEncryptedParamsFormat = "JSON"
	EmailRoutingSummaryEncryptedParamsFormatCsv  EmailRoutingSummaryEncryptedParamsFormat = "CSV"
)

func (r EmailRoutingSummaryEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsFormatJson, EmailRoutingSummaryEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsIPVersion string

const (
	EmailRoutingSummaryEncryptedParamsIPVersionIPv4 EmailRoutingSummaryEncryptedParamsIPVersion = "IPv4"
	EmailRoutingSummaryEncryptedParamsIPVersionIPv6 EmailRoutingSummaryEncryptedParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsIPVersionIPv4, EmailRoutingSummaryEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsSPF string

const (
	EmailRoutingSummaryEncryptedParamsSPFPass EmailRoutingSummaryEncryptedParamsSPF = "PASS"
	EmailRoutingSummaryEncryptedParamsSPFNone EmailRoutingSummaryEncryptedParamsSPF = "NONE"
	EmailRoutingSummaryEncryptedParamsSPFFail EmailRoutingSummaryEncryptedParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsSPFPass, EmailRoutingSummaryEncryptedParamsSPFNone, EmailRoutingSummaryEncryptedParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedResponseEnvelope struct {
	Result  EmailRoutingSummaryEncryptedResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    emailRoutingSummaryEncryptedResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryEncryptedResponseEnvelope]
type emailRoutingSummaryEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummaryIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummaryIPVersionParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingSummaryIPVersionParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummaryIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummaryIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingSummaryIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryIPVersionParamsARC string

const (
	EmailRoutingSummaryIPVersionParamsARCPass EmailRoutingSummaryIPVersionParamsARC = "PASS"
	EmailRoutingSummaryIPVersionParamsARCNone EmailRoutingSummaryIPVersionParamsARC = "NONE"
	EmailRoutingSummaryIPVersionParamsARCFail EmailRoutingSummaryIPVersionParamsARC = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsARCPass, EmailRoutingSummaryIPVersionParamsARCNone, EmailRoutingSummaryIPVersionParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsDKIM string

const (
	EmailRoutingSummaryIPVersionParamsDKIMPass EmailRoutingSummaryIPVersionParamsDKIM = "PASS"
	EmailRoutingSummaryIPVersionParamsDKIMNone EmailRoutingSummaryIPVersionParamsDKIM = "NONE"
	EmailRoutingSummaryIPVersionParamsDKIMFail EmailRoutingSummaryIPVersionParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsDKIMPass, EmailRoutingSummaryIPVersionParamsDKIMNone, EmailRoutingSummaryIPVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsDMARC string

const (
	EmailRoutingSummaryIPVersionParamsDMARCPass EmailRoutingSummaryIPVersionParamsDMARC = "PASS"
	EmailRoutingSummaryIPVersionParamsDMARCNone EmailRoutingSummaryIPVersionParamsDMARC = "NONE"
	EmailRoutingSummaryIPVersionParamsDMARCFail EmailRoutingSummaryIPVersionParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsDMARCPass, EmailRoutingSummaryIPVersionParamsDMARCNone, EmailRoutingSummaryIPVersionParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsEncrypted string

const (
	EmailRoutingSummaryIPVersionParamsEncryptedEncrypted    EmailRoutingSummaryIPVersionParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted EmailRoutingSummaryIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsEncryptedEncrypted, EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummaryIPVersionParamsFormat string

const (
	EmailRoutingSummaryIPVersionParamsFormatJson EmailRoutingSummaryIPVersionParamsFormat = "JSON"
	EmailRoutingSummaryIPVersionParamsFormatCsv  EmailRoutingSummaryIPVersionParamsFormat = "CSV"
)

func (r EmailRoutingSummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsFormatJson, EmailRoutingSummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsSPF string

const (
	EmailRoutingSummaryIPVersionParamsSPFPass EmailRoutingSummaryIPVersionParamsSPF = "PASS"
	EmailRoutingSummaryIPVersionParamsSPFNone EmailRoutingSummaryIPVersionParamsSPF = "NONE"
	EmailRoutingSummaryIPVersionParamsSPFFail EmailRoutingSummaryIPVersionParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsSPFPass, EmailRoutingSummaryIPVersionParamsSPFNone, EmailRoutingSummaryIPVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionResponseEnvelope struct {
	Result  EmailRoutingSummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    emailRoutingSummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryIPVersionResponseEnvelope]
type emailRoutingSummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingSummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingSummarySPFParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingSummarySPFParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingSummarySPFParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingSummarySPFParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingSummarySPFParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [EmailRoutingSummarySPFParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummarySPFParamsARC string

const (
	EmailRoutingSummarySPFParamsARCPass EmailRoutingSummarySPFParamsARC = "PASS"
	EmailRoutingSummarySPFParamsARCNone EmailRoutingSummarySPFParamsARC = "NONE"
	EmailRoutingSummarySPFParamsARCFail EmailRoutingSummarySPFParamsARC = "FAIL"
)

func (r EmailRoutingSummarySPFParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsARCPass, EmailRoutingSummarySPFParamsARCNone, EmailRoutingSummarySPFParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsDKIM string

const (
	EmailRoutingSummarySPFParamsDKIMPass EmailRoutingSummarySPFParamsDKIM = "PASS"
	EmailRoutingSummarySPFParamsDKIMNone EmailRoutingSummarySPFParamsDKIM = "NONE"
	EmailRoutingSummarySPFParamsDKIMFail EmailRoutingSummarySPFParamsDKIM = "FAIL"
)

func (r EmailRoutingSummarySPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsDKIMPass, EmailRoutingSummarySPFParamsDKIMNone, EmailRoutingSummarySPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsDMARC string

const (
	EmailRoutingSummarySPFParamsDMARCPass EmailRoutingSummarySPFParamsDMARC = "PASS"
	EmailRoutingSummarySPFParamsDMARCNone EmailRoutingSummarySPFParamsDMARC = "NONE"
	EmailRoutingSummarySPFParamsDMARCFail EmailRoutingSummarySPFParamsDMARC = "FAIL"
)

func (r EmailRoutingSummarySPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsDMARCPass, EmailRoutingSummarySPFParamsDMARCNone, EmailRoutingSummarySPFParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsEncrypted string

const (
	EmailRoutingSummarySPFParamsEncryptedEncrypted    EmailRoutingSummarySPFParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummarySPFParamsEncryptedNotEncrypted EmailRoutingSummarySPFParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummarySPFParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsEncryptedEncrypted, EmailRoutingSummarySPFParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingSummarySPFParamsFormat string

const (
	EmailRoutingSummarySPFParamsFormatJson EmailRoutingSummarySPFParamsFormat = "JSON"
	EmailRoutingSummarySPFParamsFormatCsv  EmailRoutingSummarySPFParamsFormat = "CSV"
)

func (r EmailRoutingSummarySPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsFormatJson, EmailRoutingSummarySPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsIPVersion string

const (
	EmailRoutingSummarySPFParamsIPVersionIPv4 EmailRoutingSummarySPFParamsIPVersion = "IPv4"
	EmailRoutingSummarySPFParamsIPVersionIPv6 EmailRoutingSummarySPFParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummarySPFParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsIPVersionIPv4, EmailRoutingSummarySPFParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummarySPFResponseEnvelope struct {
	Result  EmailRoutingSummarySPFResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailRoutingSummarySPFResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummarySPFResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummarySPFResponseEnvelope]
type emailRoutingSummarySPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
