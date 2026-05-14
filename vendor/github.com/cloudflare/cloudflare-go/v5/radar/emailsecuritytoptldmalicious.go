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

// EmailSecurityTopTldMaliciousService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldMaliciousService] method instead.
type EmailSecurityTopTldMaliciousService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldMaliciousService(opts ...option.RequestOption) (r *EmailSecurityTopTldMaliciousService) {
	r = &EmailSecurityTopTldMaliciousService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as malicious or not.
func (r *EmailSecurityTopTldMaliciousService) Get(ctx context.Context, malicious EmailSecurityTopTldMaliciousGetParamsMalicious, query EmailSecurityTopTldMaliciousGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldMaliciousGetResponse, err error) {
	var env EmailSecurityTopTldMaliciousGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldMaliciousGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTldMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldMaliciousGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldMaliciousGetResponse]
type emailSecurityTopTldMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTldMaliciousGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTldMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTldMaliciousGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTldMaliciousGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTldMaliciousGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldMaliciousGetResponseMeta]
type emailSecurityTopTldMaliciousGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo]
type emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldMaliciousGetResponseMetaDateRange]
type emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTldMaliciousGetResponseMetaNormalization string

const (
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationPercentage           EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationMin0Max              EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationMinMax               EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRawValues            EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRatio                EmailSecurityTopTldMaliciousGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTldMaliciousGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetResponseMetaNormalizationPercentage, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationMin0Max, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationMinMax, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRawValues, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTldMaliciousGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailSecurityTopTldMaliciousGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldMaliciousGetResponseMetaUnit]
type emailSecurityTopTldMaliciousGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                          `json:"value,required"`
	JSON  emailSecurityTopTldMaliciousGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseTop0JSON contains the JSON metadata for
// the struct [EmailSecurityTopTldMaliciousGetResponseTop0]
type emailSecurityTopTldMaliciousGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTldMaliciousGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTldMaliciousGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTldMaliciousGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTldMaliciousGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTldMaliciousGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[EmailSecurityTopTldMaliciousGetParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTldMaliciousGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldMaliciousGetParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTopTldMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Malicious classification.
type EmailSecurityTopTldMaliciousGetParamsMalicious string

const (
	EmailSecurityTopTldMaliciousGetParamsMaliciousMalicious    EmailSecurityTopTldMaliciousGetParamsMalicious = "MALICIOUS"
	EmailSecurityTopTldMaliciousGetParamsMaliciousNotMalicious EmailSecurityTopTldMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

func (r EmailSecurityTopTldMaliciousGetParamsMalicious) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsMaliciousMalicious, EmailSecurityTopTldMaliciousGetParamsMaliciousNotMalicious:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetParamsARC string

const (
	EmailSecurityTopTldMaliciousGetParamsARCPass EmailSecurityTopTldMaliciousGetParamsARC = "PASS"
	EmailSecurityTopTldMaliciousGetParamsARCNone EmailSecurityTopTldMaliciousGetParamsARC = "NONE"
	EmailSecurityTopTldMaliciousGetParamsARCFail EmailSecurityTopTldMaliciousGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldMaliciousGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsARCPass, EmailSecurityTopTldMaliciousGetParamsARCNone, EmailSecurityTopTldMaliciousGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetParamsDKIM string

const (
	EmailSecurityTopTldMaliciousGetParamsDKIMPass EmailSecurityTopTldMaliciousGetParamsDKIM = "PASS"
	EmailSecurityTopTldMaliciousGetParamsDKIMNone EmailSecurityTopTldMaliciousGetParamsDKIM = "NONE"
	EmailSecurityTopTldMaliciousGetParamsDKIMFail EmailSecurityTopTldMaliciousGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldMaliciousGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsDKIMPass, EmailSecurityTopTldMaliciousGetParamsDKIMNone, EmailSecurityTopTldMaliciousGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetParamsDMARC string

const (
	EmailSecurityTopTldMaliciousGetParamsDMARCPass EmailSecurityTopTldMaliciousGetParamsDMARC = "PASS"
	EmailSecurityTopTldMaliciousGetParamsDMARCNone EmailSecurityTopTldMaliciousGetParamsDMARC = "NONE"
	EmailSecurityTopTldMaliciousGetParamsDMARCFail EmailSecurityTopTldMaliciousGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldMaliciousGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsDMARCPass, EmailSecurityTopTldMaliciousGetParamsDMARCNone, EmailSecurityTopTldMaliciousGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTldMaliciousGetParamsFormat string

const (
	EmailSecurityTopTldMaliciousGetParamsFormatJson EmailSecurityTopTldMaliciousGetParamsFormat = "JSON"
	EmailSecurityTopTldMaliciousGetParamsFormatCsv  EmailSecurityTopTldMaliciousGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldMaliciousGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsFormatJson, EmailSecurityTopTldMaliciousGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetParamsSPF string

const (
	EmailSecurityTopTldMaliciousGetParamsSPFPass EmailSecurityTopTldMaliciousGetParamsSPF = "PASS"
	EmailSecurityTopTldMaliciousGetParamsSPFNone EmailSecurityTopTldMaliciousGetParamsSPF = "NONE"
	EmailSecurityTopTldMaliciousGetParamsSPFFail EmailSecurityTopTldMaliciousGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldMaliciousGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsSPFPass, EmailSecurityTopTldMaliciousGetParamsSPFNone, EmailSecurityTopTldMaliciousGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTldMaliciousGetParamsTldCategory string

const (
	EmailSecurityTopTldMaliciousGetParamsTldCategoryClassic EmailSecurityTopTldMaliciousGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldMaliciousGetParamsTldCategoryCountry EmailSecurityTopTldMaliciousGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldMaliciousGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsTldCategoryClassic, EmailSecurityTopTldMaliciousGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetParamsTLSVersion string

const (
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldMaliciousGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldMaliciousGetResponseEnvelope struct {
	Result  EmailSecurityTopTldMaliciousGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTopTldMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldMaliciousGetResponseEnvelope]
type emailSecurityTopTldMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
