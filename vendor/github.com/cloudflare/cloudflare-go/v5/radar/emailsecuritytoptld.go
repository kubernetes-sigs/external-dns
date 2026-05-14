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

// EmailSecurityTopTldService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldService] method instead.
type EmailSecurityTopTldService struct {
	Options   []option.RequestOption
	Malicious *EmailSecurityTopTldMaliciousService
	Spam      *EmailSecurityTopTldSpamService
	Spoof     *EmailSecurityTopTldSpoofService
}

// NewEmailSecurityTopTldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldService(opts ...option.RequestOption) (r *EmailSecurityTopTldService) {
	r = &EmailSecurityTopTldService{}
	r.Options = opts
	r.Malicious = NewEmailSecurityTopTldMaliciousService(opts...)
	r.Spam = NewEmailSecurityTopTldSpamService(opts...)
	r.Spoof = NewEmailSecurityTopTldSpoofService(opts...)
	return
}

// Retrieves the top TLDs by number of email messages.
func (r *EmailSecurityTopTldService) Get(ctx context.Context, query EmailSecurityTopTldGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldGetResponse, err error) {
	var env EmailSecurityTopTldGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/top/tlds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTldGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldGetResponse]
type emailSecurityTopTldGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTldGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTldGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTldGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTldGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTldGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTldGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTldGetResponseMetaJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldGetResponseMeta]
type emailSecurityTopTldGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  emailSecurityTopTldGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldGetResponseMetaConfidenceInfo]
type emailSecurityTopTldGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldGetResponseMetaDateRange]
type emailSecurityTopTldGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTldGetResponseMetaNormalization string

const (
	EmailSecurityTopTldGetResponseMetaNormalizationPercentage           EmailSecurityTopTldGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTldGetResponseMetaNormalizationMin0Max              EmailSecurityTopTldGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTldGetResponseMetaNormalizationMinMax               EmailSecurityTopTldGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTldGetResponseMetaNormalizationRawValues            EmailSecurityTopTldGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTldGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTldGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTldGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTldGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTldGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTldGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTldGetResponseMetaNormalizationRatio                EmailSecurityTopTldGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTldGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetResponseMetaNormalizationPercentage, EmailSecurityTopTldGetResponseMetaNormalizationMin0Max, EmailSecurityTopTldGetResponseMetaNormalizationMinMax, EmailSecurityTopTldGetResponseMetaNormalizationRawValues, EmailSecurityTopTldGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTldGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTldGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTldGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTldGetResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  emailSecurityTopTldGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldGetResponseMetaUnit]
type emailSecurityTopTldGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                 `json:"value,required"`
	JSON  emailSecurityTopTldGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldGetResponseTop0JSON contains the JSON metadata for the struct
// [EmailSecurityTopTldGetResponseTop0]
type emailSecurityTopTldGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTldGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTldGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTldGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTldGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTldGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[EmailSecurityTopTldGetParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTldGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecurityTopTldGetParamsARC string

const (
	EmailSecurityTopTldGetParamsARCPass EmailSecurityTopTldGetParamsARC = "PASS"
	EmailSecurityTopTldGetParamsARCNone EmailSecurityTopTldGetParamsARC = "NONE"
	EmailSecurityTopTldGetParamsARCFail EmailSecurityTopTldGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsARCPass, EmailSecurityTopTldGetParamsARCNone, EmailSecurityTopTldGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsDKIM string

const (
	EmailSecurityTopTldGetParamsDKIMPass EmailSecurityTopTldGetParamsDKIM = "PASS"
	EmailSecurityTopTldGetParamsDKIMNone EmailSecurityTopTldGetParamsDKIM = "NONE"
	EmailSecurityTopTldGetParamsDKIMFail EmailSecurityTopTldGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsDKIMPass, EmailSecurityTopTldGetParamsDKIMNone, EmailSecurityTopTldGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsDMARC string

const (
	EmailSecurityTopTldGetParamsDMARCPass EmailSecurityTopTldGetParamsDMARC = "PASS"
	EmailSecurityTopTldGetParamsDMARCNone EmailSecurityTopTldGetParamsDMARC = "NONE"
	EmailSecurityTopTldGetParamsDMARCFail EmailSecurityTopTldGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsDMARCPass, EmailSecurityTopTldGetParamsDMARCNone, EmailSecurityTopTldGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTldGetParamsFormat string

const (
	EmailSecurityTopTldGetParamsFormatJson EmailSecurityTopTldGetParamsFormat = "JSON"
	EmailSecurityTopTldGetParamsFormatCsv  EmailSecurityTopTldGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsFormatJson, EmailSecurityTopTldGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsSPF string

const (
	EmailSecurityTopTldGetParamsSPFPass EmailSecurityTopTldGetParamsSPF = "PASS"
	EmailSecurityTopTldGetParamsSPFNone EmailSecurityTopTldGetParamsSPF = "NONE"
	EmailSecurityTopTldGetParamsSPFFail EmailSecurityTopTldGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsSPFPass, EmailSecurityTopTldGetParamsSPFNone, EmailSecurityTopTldGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTldGetParamsTldCategory string

const (
	EmailSecurityTopTldGetParamsTldCategoryClassic EmailSecurityTopTldGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldGetParamsTldCategoryCountry EmailSecurityTopTldGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsTldCategoryClassic, EmailSecurityTopTldGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsTLSVersion string

const (
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldGetResponseEnvelope struct {
	Result  EmailSecurityTopTldGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailSecurityTopTldGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldGetResponseEnvelope]
type emailSecurityTopTldGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
