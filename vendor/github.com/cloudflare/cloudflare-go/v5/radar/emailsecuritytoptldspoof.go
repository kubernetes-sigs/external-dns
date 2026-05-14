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

// EmailSecurityTopTldSpoofService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldSpoofService] method instead.
type EmailSecurityTopTldSpoofService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldSpoofService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldSpoofService(opts ...option.RequestOption) (r *EmailSecurityTopTldSpoofService) {
	r = &EmailSecurityTopTldSpoofService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by emails classified as spoof or not.
func (r *EmailSecurityTopTldSpoofService) Get(ctx context.Context, spoof EmailSecurityTopTldSpoofGetParamsSpoof, query EmailSecurityTopTldSpoofGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldSpoofGetResponse, err error) {
	var env EmailSecurityTopTldSpoofGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spoof/%v", spoof)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldSpoofGetResponse struct {
	// Metadata for the results.
	Meta EmailSecurityTopTldSpoofGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldSpoofGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldSpoofGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponse]
type emailSecurityTopTldSpoofGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailSecurityTopTldSpoofGetResponseMeta struct {
	ConfidenceInfo EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []EmailSecurityTopTldSpoofGetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecurityTopTldSpoofGetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecurityTopTldSpoofGetResponseMetaUnit `json:"units,required"`
	JSON  emailSecurityTopTldSpoofGetResponseMetaJSON   `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponseMeta]
type emailSecurityTopTldSpoofGetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo]
type emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldSpoofGetResponseMetaDateRange]
type emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecurityTopTldSpoofGetResponseMetaNormalization string

const (
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationPercentage           EmailSecurityTopTldSpoofGetResponseMetaNormalization = "PERCENTAGE"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationMin0Max              EmailSecurityTopTldSpoofGetResponseMetaNormalization = "MIN0_MAX"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationMinMax               EmailSecurityTopTldSpoofGetResponseMetaNormalization = "MIN_MAX"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationRawValues            EmailSecurityTopTldSpoofGetResponseMetaNormalization = "RAW_VALUES"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationPercentageChange     EmailSecurityTopTldSpoofGetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationRollingAverage       EmailSecurityTopTldSpoofGetResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationOverlappedPercentage EmailSecurityTopTldSpoofGetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecurityTopTldSpoofGetResponseMetaNormalizationRatio                EmailSecurityTopTldSpoofGetResponseMetaNormalization = "RATIO"
)

func (r EmailSecurityTopTldSpoofGetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetResponseMetaNormalizationPercentage, EmailSecurityTopTldSpoofGetResponseMetaNormalizationMin0Max, EmailSecurityTopTldSpoofGetResponseMetaNormalizationMinMax, EmailSecurityTopTldSpoofGetResponseMetaNormalizationRawValues, EmailSecurityTopTldSpoofGetResponseMetaNormalizationPercentageChange, EmailSecurityTopTldSpoofGetResponseMetaNormalizationRollingAverage, EmailSecurityTopTldSpoofGetResponseMetaNormalizationOverlappedPercentage, EmailSecurityTopTldSpoofGetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  emailSecurityTopTldSpoofGetResponseMetaUnitJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpoofGetResponseMetaUnit]
type emailSecurityTopTldSpoofGetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseTop0 struct {
	Name string `json:"name,required"`
	// A numeric string.
	Value string                                      `json:"value,required"`
	JSON  emailSecurityTopTldSpoofGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseTop0JSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponseTop0]
type emailSecurityTopTldSpoofGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecurityTopTldSpoofGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecurityTopTldSpoofGetParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecurityTopTldSpoofGetParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecurityTopTldSpoofGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecurityTopTldSpoofGetParamsSPF] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[EmailSecurityTopTldSpoofGetParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecurityTopTldSpoofGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldSpoofGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldSpoofGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spoof classification.
type EmailSecurityTopTldSpoofGetParamsSpoof string

const (
	EmailSecurityTopTldSpoofGetParamsSpoofSpoof    EmailSecurityTopTldSpoofGetParamsSpoof = "SPOOF"
	EmailSecurityTopTldSpoofGetParamsSpoofNotSpoof EmailSecurityTopTldSpoofGetParamsSpoof = "NOT_SPOOF"
)

func (r EmailSecurityTopTldSpoofGetParamsSpoof) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsSpoofSpoof, EmailSecurityTopTldSpoofGetParamsSpoofNotSpoof:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsARC string

const (
	EmailSecurityTopTldSpoofGetParamsARCPass EmailSecurityTopTldSpoofGetParamsARC = "PASS"
	EmailSecurityTopTldSpoofGetParamsARCNone EmailSecurityTopTldSpoofGetParamsARC = "NONE"
	EmailSecurityTopTldSpoofGetParamsARCFail EmailSecurityTopTldSpoofGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsARCPass, EmailSecurityTopTldSpoofGetParamsARCNone, EmailSecurityTopTldSpoofGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsDKIM string

const (
	EmailSecurityTopTldSpoofGetParamsDKIMPass EmailSecurityTopTldSpoofGetParamsDKIM = "PASS"
	EmailSecurityTopTldSpoofGetParamsDKIMNone EmailSecurityTopTldSpoofGetParamsDKIM = "NONE"
	EmailSecurityTopTldSpoofGetParamsDKIMFail EmailSecurityTopTldSpoofGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsDKIMPass, EmailSecurityTopTldSpoofGetParamsDKIMNone, EmailSecurityTopTldSpoofGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsDMARC string

const (
	EmailSecurityTopTldSpoofGetParamsDMARCPass EmailSecurityTopTldSpoofGetParamsDMARC = "PASS"
	EmailSecurityTopTldSpoofGetParamsDMARCNone EmailSecurityTopTldSpoofGetParamsDMARC = "NONE"
	EmailSecurityTopTldSpoofGetParamsDMARCFail EmailSecurityTopTldSpoofGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsDMARCPass, EmailSecurityTopTldSpoofGetParamsDMARCNone, EmailSecurityTopTldSpoofGetParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecurityTopTldSpoofGetParamsFormat string

const (
	EmailSecurityTopTldSpoofGetParamsFormatJson EmailSecurityTopTldSpoofGetParamsFormat = "JSON"
	EmailSecurityTopTldSpoofGetParamsFormatCsv  EmailSecurityTopTldSpoofGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldSpoofGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsFormatJson, EmailSecurityTopTldSpoofGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsSPF string

const (
	EmailSecurityTopTldSpoofGetParamsSPFPass EmailSecurityTopTldSpoofGetParamsSPF = "PASS"
	EmailSecurityTopTldSpoofGetParamsSPFNone EmailSecurityTopTldSpoofGetParamsSPF = "NONE"
	EmailSecurityTopTldSpoofGetParamsSPFFail EmailSecurityTopTldSpoofGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsSPFPass, EmailSecurityTopTldSpoofGetParamsSPFNone, EmailSecurityTopTldSpoofGetParamsSPFFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type EmailSecurityTopTldSpoofGetParamsTldCategory string

const (
	EmailSecurityTopTldSpoofGetParamsTldCategoryClassic EmailSecurityTopTldSpoofGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldSpoofGetParamsTldCategoryCountry EmailSecurityTopTldSpoofGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldSpoofGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsTldCategoryClassic, EmailSecurityTopTldSpoofGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsTLSVersion string

const (
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldSpoofGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetResponseEnvelope struct {
	Result  EmailSecurityTopTldSpoofGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    emailSecurityTopTldSpoofGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpoofGetResponseEnvelope]
type emailSecurityTopTldSpoofGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
