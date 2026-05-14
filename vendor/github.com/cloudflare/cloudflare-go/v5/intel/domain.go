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
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
	Bulks   *DomainBulkService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	r.Bulks = NewDomainBulkService(opts...)
	return
}

// Gets security details and statistics about a domain.
func (r *DomainService) Get(ctx context.Context, params DomainGetParams, opts ...option.RequestOption) (res *Domain, err error) {
	var env DomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/domain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Domain struct {
	// Additional information related to the host name.
	AdditionalInformation DomainAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application                DomainApplication                `json:"application"`
	ContentCategories          []DomainContentCategory          `json:"content_categories"`
	Domain                     string                           `json:"domain"`
	InheritedContentCategories []DomainInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                    `json:"inherited_from"`
	InheritedRiskTypes []DomainInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []DomainResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64          `json:"risk_score"`
	RiskTypes []DomainRiskType `json:"risk_types"`
	JSON      domainJSON       `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	ResolvesToRefs             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJSON) RawJSON() string {
	return r.raw
}

// Additional information related to the host name.
type DomainAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                          `json:"suspected_malware_family"`
	JSON                   domainAdditionalInformationJSON `json:"-"`
}

// domainAdditionalInformationJSON contains the JSON metadata for the struct
// [DomainAdditionalInformation]
type domainAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DomainAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainAdditionalInformationJSON) RawJSON() string {
	return r.raw
}

// Application that the hostname belongs to.
type DomainApplication struct {
	ID   int64                 `json:"id"`
	Name string                `json:"name"`
	JSON domainApplicationJSON `json:"-"`
}

// domainApplicationJSON contains the JSON metadata for the struct
// [DomainApplication]
type domainApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainApplicationJSON) RawJSON() string {
	return r.raw
}

// Current content categories.
type DomainContentCategory struct {
	ID              int64                     `json:"id"`
	Name            string                    `json:"name"`
	SuperCategoryID int64                     `json:"super_category_id"`
	JSON            domainContentCategoryJSON `json:"-"`
}

// domainContentCategoryJSON contains the JSON metadata for the struct
// [DomainContentCategory]
type domainContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainContentCategoryJSON) RawJSON() string {
	return r.raw
}

type DomainInheritedContentCategory struct {
	ID              int64                              `json:"id"`
	Name            string                             `json:"name"`
	SuperCategoryID int64                              `json:"super_category_id"`
	JSON            domainInheritedContentCategoryJSON `json:"-"`
}

// domainInheritedContentCategoryJSON contains the JSON metadata for the struct
// [DomainInheritedContentCategory]
type domainInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainInheritedContentCategoryJSON) RawJSON() string {
	return r.raw
}

type DomainInheritedRiskType struct {
	ID              int64                       `json:"id"`
	Name            string                      `json:"name"`
	SuperCategoryID int64                       `json:"super_category_id"`
	JSON            domainInheritedRiskTypeJSON `json:"-"`
}

// domainInheritedRiskTypeJSON contains the JSON metadata for the struct
// [DomainInheritedRiskType]
type domainInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainInheritedRiskTypeJSON) RawJSON() string {
	return r.raw
}

type DomainResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c.
	ID string `json:"id"`
	// IP address or domain name.
	Value string                  `json:"value"`
	JSON  domainResolvesToRefJSON `json:"-"`
}

// domainResolvesToRefJSON contains the JSON metadata for the struct
// [DomainResolvesToRef]
type domainResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainResolvesToRefJSON) RawJSON() string {
	return r.raw
}

type DomainRiskType struct {
	ID              int64              `json:"id"`
	Name            string             `json:"name"`
	SuperCategoryID int64              `json:"super_category_id"`
	JSON            domainRiskTypeJSON `json:"-"`
}

// domainRiskTypeJSON contains the JSON metadata for the struct [DomainRiskType]
type domainRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainRiskTypeJSON) RawJSON() string {
	return r.raw
}

type DomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [DomainGetParams]'s query parameters as `url.Values`.
func (r DomainGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainGetResponseEnvelope struct {
	Errors   []DomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DomainGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Domain                           `json:"result"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             domainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// domainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeErrors]
type domainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    domainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeErrorsSource]
type domainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             domainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// domainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeMessages]
type domainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    domainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeMessagesSource]
type domainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
