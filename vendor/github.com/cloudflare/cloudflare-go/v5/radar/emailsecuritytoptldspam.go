// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EmailSecurityTopTldSpamService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldSpamService] method instead.
type EmailSecurityTopTldSpamService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldSpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldSpamService(opts ...option.RequestOption) (r *EmailSecurityTopTldSpamService) {
	r = &EmailSecurityTopTldSpamService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as spam or not.
func (r *EmailSecurityTopTldSpamService) Get(ctx context.Context, spam EmailSecurityTopTldSpamGetParamsSpam, query EmailSecurityTopTldSpamGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldSpamGetResponse, err error) {
	var env EmailSecurityTopTldSpamGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldSpamGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTldSpamGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldSpamGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldSpamGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldSpamGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldSpamGetResponse]
type emailSecurityTopTldSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTldSpamGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTldSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTldSpamGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTldSpamGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTldSpamGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpamGetResponseMeta]
type emailSecurityTopTldSpamGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo]
type emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldSpamGetResponseMetaDateRange]
type emailSecurityTopTldSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTldSpamGetResponseMetaNormalization string

const (
	EmailSecurityTopTldSpamGetResponseMetaNormalizationPercentage           EmailSecurityTopTldSpamGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationMin0Max              EmailSecurityTopTldSpamGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationMinMax               EmailSecurityTopTldSpamGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationRawValues            EmailSecurityTopTldSpamGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTldSpamGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTldSpamGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTldSpamGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTldSpamGetResponseMetaNormalizationRatio                EmailSecurityTopTldSpamGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTldSpamGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetResponseMetaNormalizationPercentage, EmailSecurityTopTldSpamGetResponseMetaNormalizationMin0Max, EmailSecurityTopTldSpamGetResponseMetaNormalizationMinMax, EmailSecurityTopTldSpamGetResponseMetaNormalizationRawValues, EmailSecurityTopTldSpamGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTldSpamGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTldSpamGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTldSpamGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  emailSecurityTopTldSpamGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpamGetResponseMetaUnit]
type emailSecurityTopTldSpamGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                     `json:"value,required"`
	JSON  emailSecurityTopTldSpamGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseTop0JSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpamGetResponseTop0]
type emailSecurityTopTldSpamGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTldSpamGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTldSpamGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTldSpamGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTldSpamGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTldSpamGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[EmailSecurityTopTldSpamGetParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTldSpamGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldSpamGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spam classification.
type EmailSecurityTopTldSpamGetParamsSpam string

const (
	EmailSecurityTopTldSpamGetParamsSpamSpam    EmailSecurityTopTldSpamGetParamsSpam = "SPAM"
	EmailSecurityTopTldSpamGetParamsSpamNotSpam EmailSecurityTopTldSpamGetParamsSpam = "NOT_SPAM"
)

func (r EmailSecurityTopTldSpamGetParamsSpam) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsSpamSpam, EmailSecurityTopTldSpamGetParamsSpamNotSpam:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsARC string

const (
	EmailSecurityTopTldSpamGetParamsARCPass EmailSecurityTopTldSpamGetParamsARC = "PASS"
	EmailSecurityTopTldSpamGetParamsARCNone EmailSecurityTopTldSpamGetParamsARC = "NONE"
	EmailSecurityTopTldSpamGetParamsARCFail EmailSecurityTopTldSpamGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsARCPass, EmailSecurityTopTldSpamGetParamsARCNone, EmailSecurityTopTldSpamGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsDKIM string

const (
	EmailSecurityTopTldSpamGetParamsDKIMPass EmailSecurityTopTldSpamGetParamsDKIM = "PASS"
	EmailSecurityTopTldSpamGetParamsDKIMNone EmailSecurityTopTldSpamGetParamsDKIM = "NONE"
	EmailSecurityTopTldSpamGetParamsDKIMFail EmailSecurityTopTldSpamGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsDKIMPass, EmailSecurityTopTldSpamGetParamsDKIMNone, EmailSecurityTopTldSpamGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsDMARC string

const (
	EmailSecurityTopTldSpamGetParamsDMARCPass EmailSecurityTopTldSpamGetParamsDMARC = "PASS"
	EmailSecurityTopTldSpamGetParamsDMARCNone EmailSecurityTopTldSpamGetParamsDMARC = "NONE"
	EmailSecurityTopTldSpamGetParamsDMARCFail EmailSecurityTopTldSpamGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsDMARCPass, EmailSecurityTopTldSpamGetParamsDMARCNone, EmailSecurityTopTldSpamGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTldSpamGetParamsFormat string

const (
	EmailSecurityTopTldSpamGetParamsFormatJson EmailSecurityTopTldSpamGetParamsFormat = "JSON"
	EmailSecurityTopTldSpamGetParamsFormatCsv  EmailSecurityTopTldSpamGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldSpamGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsFormatJson, EmailSecurityTopTldSpamGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsSPF string

const (
	EmailSecurityTopTldSpamGetParamsSPFPass EmailSecurityTopTldSpamGetParamsSPF = "PASS"
	EmailSecurityTopTldSpamGetParamsSPFNone EmailSecurityTopTldSpamGetParamsSPF = "NONE"
	EmailSecurityTopTldSpamGetParamsSPFFail EmailSecurityTopTldSpamGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsSPFPass, EmailSecurityTopTldSpamGetParamsSPFNone, EmailSecurityTopTldSpamGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTldSpamGetParamsTldCategory string

const (
	EmailSecurityTopTldSpamGetParamsTldCategoryClassic EmailSecurityTopTldSpamGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldSpamGetParamsTldCategoryCountry EmailSecurityTopTldSpamGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldSpamGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsTldCategoryClassic, EmailSecurityTopTldSpamGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsTLSVersion string

const (
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldSpamGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetResponseEnvelope struct {
	Result  EmailSecurityTopTldSpamGetResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    emailSecurityTopTldSpamGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpamGetResponseEnvelope]
type emailSecurityTopTldSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
