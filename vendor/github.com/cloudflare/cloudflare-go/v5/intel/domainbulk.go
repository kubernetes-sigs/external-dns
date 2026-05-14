// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DomainBulkService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainBulkService] method instead.
type DomainBulkService struct {
	Options []option.RequestOption
}

// NewDomainBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainBulkService(opts ...option.RequestOption) (r *DomainBulkService) {
	r = &DomainBulkService{}
	r.Options = opts
	return
}

// Same as summary.
func (r *DomainBulkService) Get(ctx context.Context, params DomainBulkGetParams, opts ...option.RequestOption) (res *[]DomainBulkGetResponse, err error) {
	var env DomainBulkGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DomainBulkGetResponse struct {
	// Additional information related to the host name.
	AdditionalInformation DomainBulkGetResponseAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application                DomainBulkGetResponseApplication                `json:"application"`
	ContentCategories          []DomainBulkGetResponseContentCategory          `json:"content_categories"`
	Domain                     string                                          `json:"domain"`
	InheritedContentCategories []DomainBulkGetResponseInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                   `json:"inherited_from"`
	InheritedRiskTypes []DomainBulkGetResponseInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                         `json:"risk_score"`
	RiskTypes []DomainBulkGetResponseRiskType `json:"risk_types"`
	JSON      domainBulkGetResponseJSON       `json:"-"`
}

// domainBulkGetResponseJSON contains the JSON metadata for the struct
// [DomainBulkGetResponse]
type domainBulkGetResponseJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DomainBulkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseJSON) RawJSON() string {
	return r.raw
}

// Additional information related to the host name.
type DomainBulkGetResponseAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                         `json:"suspected_malware_family"`
	JSON                   domainBulkGetResponseAdditionalInformationJSON `json:"-"`
}

// domainBulkGetResponseAdditionalInformationJSON contains the JSON metadata for
// the struct [DomainBulkGetResponseAdditionalInformation]
type domainBulkGetResponseAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DomainBulkGetResponseAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseAdditionalInformationJSON) RawJSON() string {
	return r.raw
}

// Application that the hostname belongs to.
type DomainBulkGetResponseApplication struct {
	ID   int64                                `json:"id"`
	Name string                               `json:"name"`
	JSON domainBulkGetResponseApplicationJSON `json:"-"`
}

// domainBulkGetResponseApplicationJSON contains the JSON metadata for the struct
// [DomainBulkGetResponseApplication]
type domainBulkGetResponseApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainBulkGetResponseApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseApplicationJSON) RawJSON() string {
	return r.raw
}

// Current content categories.
type DomainBulkGetResponseContentCategory struct {
	ID              int64                                    `json:"id"`
	Name            string                                   `json:"name"`
	SuperCategoryID int64                                    `json:"super_category_id"`
	JSON            domainBulkGetResponseContentCategoryJSON `json:"-"`
}

// domainBulkGetResponseContentCategoryJSON contains the JSON metadata for the
// struct [DomainBulkGetResponseContentCategory]
type domainBulkGetResponseContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainBulkGetResponseContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseContentCategoryJSON) RawJSON() string {
	return r.raw
}

type DomainBulkGetResponseInheritedContentCategory struct {
	ID              int64                                             `json:"id"`
	Name            string                                            `json:"name"`
	SuperCategoryID int64                                             `json:"super_category_id"`
	JSON            domainBulkGetResponseInheritedContentCategoryJSON `json:"-"`
}

// domainBulkGetResponseInheritedContentCategoryJSON contains the JSON metadata for
// the struct [DomainBulkGetResponseInheritedContentCategory]
type domainBulkGetResponseInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainBulkGetResponseInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseInheritedContentCategoryJSON) RawJSON() string {
	return r.raw
}

type DomainBulkGetResponseInheritedRiskType struct {
	ID              int64                                      `json:"id"`
	Name            string                                     `json:"name"`
	SuperCategoryID int64                                      `json:"super_category_id"`
	JSON            domainBulkGetResponseInheritedRiskTypeJSON `json:"-"`
}

// domainBulkGetResponseInheritedRiskTypeJSON contains the JSON metadata for the
// struct [DomainBulkGetResponseInheritedRiskType]
type domainBulkGetResponseInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainBulkGetResponseInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseInheritedRiskTypeJSON) RawJSON() string {
	return r.raw
}

type DomainBulkGetResponseRiskType struct {
	ID              int64                             `json:"id"`
	Name            string                            `json:"name"`
	SuperCategoryID int64                             `json:"super_category_id"`
	JSON            domainBulkGetResponseRiskTypeJSON `json:"-"`
}

// domainBulkGetResponseRiskTypeJSON contains the JSON metadata for the struct
// [DomainBulkGetResponseRiskType]
type domainBulkGetResponseRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainBulkGetResponseRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseRiskTypeJSON) RawJSON() string {
	return r.raw
}

type DomainBulkGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Accepts multiple values like `?domain=cloudflare.com&domain=example.com`.
	Domain param.Field[[]string] `query:"domain"`
}

// URLQuery serializes [DomainBulkGetParams]'s query parameters as `url.Values`.
func (r DomainBulkGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainBulkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   []DomainBulkGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DomainBulkGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DomainBulkGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       domainBulkGetResponseEnvelopeJSON       `json:"-"`
}

// domainBulkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainBulkGetResponseEnvelope]
type domainBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainBulkGetResponseEnvelopeSuccess bool

const (
	DomainBulkGetResponseEnvelopeSuccessTrue DomainBulkGetResponseEnvelopeSuccess = true
)

func (r DomainBulkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainBulkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainBulkGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                     `json:"total_count"`
	JSON       domainBulkGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// domainBulkGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DomainBulkGetResponseEnvelopeResultInfo]
type domainBulkGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainBulkGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainBulkGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
