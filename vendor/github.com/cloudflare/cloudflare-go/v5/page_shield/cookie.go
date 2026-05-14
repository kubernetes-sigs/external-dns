// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// CookieService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCookieService] method instead.
type CookieService struct {
	Options []option.RequestOption
}

// NewCookieService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCookieService(opts ...option.RequestOption) (r *CookieService) {
	r = &CookieService{}
	r.Options = opts
	return
}

// Lists all cookies collected by Page Shield.
func (r *CookieService) List(ctx context.Context, params CookieListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CookieListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/cookies", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all cookies collected by Page Shield.
func (r *CookieService) ListAutoPaging(ctx context.Context, params CookieListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CookieListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a cookie collected by Page Shield by cookie ID.
func (r *CookieService) Get(ctx context.Context, cookieID string, query CookieGetParams, opts ...option.RequestOption) (res *CookieGetResponse, err error) {
	var env CookieGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if cookieID == "" {
		err = errors.New("missing required cookie_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/cookies/%s", query.ZoneID, cookieID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CookieListResponse struct {
	// Identifier
	ID                string                              `json:"id,required"`
	FirstSeenAt       time.Time                           `json:"first_seen_at,required" format:"date-time"`
	Host              string                              `json:"host,required"`
	LastSeenAt        time.Time                           `json:"last_seen_at,required" format:"date-time"`
	Name              string                              `json:"name,required"`
	Type              CookieListResponseType              `json:"type,required"`
	DomainAttribute   string                              `json:"domain_attribute"`
	ExpiresAttribute  time.Time                           `json:"expires_attribute" format:"date-time"`
	HTTPOnlyAttribute bool                                `json:"http_only_attribute"`
	MaxAgeAttribute   int64                               `json:"max_age_attribute"`
	PageURLs          []string                            `json:"page_urls"`
	PathAttribute     string                              `json:"path_attribute"`
	SameSiteAttribute CookieListResponseSameSiteAttribute `json:"same_site_attribute"`
	SecureAttribute   bool                                `json:"secure_attribute"`
	JSON              cookieListResponseJSON              `json:"-"`
}

// cookieListResponseJSON contains the JSON metadata for the struct
// [CookieListResponse]
type cookieListResponseJSON struct {
	ID                apijson.Field
	FirstSeenAt       apijson.Field
	Host              apijson.Field
	LastSeenAt        apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	DomainAttribute   apijson.Field
	ExpiresAttribute  apijson.Field
	HTTPOnlyAttribute apijson.Field
	MaxAgeAttribute   apijson.Field
	PageURLs          apijson.Field
	PathAttribute     apijson.Field
	SameSiteAttribute apijson.Field
	SecureAttribute   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CookieListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieListResponseJSON) RawJSON() string {
	return r.raw
}

type CookieListResponseType string

const (
	CookieListResponseTypeFirstParty CookieListResponseType = "first_party"
	CookieListResponseTypeUnknown    CookieListResponseType = "unknown"
)

func (r CookieListResponseType) IsKnown() bool {
	switch r {
	case CookieListResponseTypeFirstParty, CookieListResponseTypeUnknown:
		return true
	}
	return false
}

type CookieListResponseSameSiteAttribute string

const (
	CookieListResponseSameSiteAttributeLax    CookieListResponseSameSiteAttribute = "lax"
	CookieListResponseSameSiteAttributeStrict CookieListResponseSameSiteAttribute = "strict"
	CookieListResponseSameSiteAttributeNone   CookieListResponseSameSiteAttribute = "none"
)

func (r CookieListResponseSameSiteAttribute) IsKnown() bool {
	switch r {
	case CookieListResponseSameSiteAttributeLax, CookieListResponseSameSiteAttributeStrict, CookieListResponseSameSiteAttributeNone:
		return true
	}
	return false
}

type CookieGetResponse struct {
	// Identifier
	ID                string                             `json:"id,required"`
	FirstSeenAt       time.Time                          `json:"first_seen_at,required" format:"date-time"`
	Host              string                             `json:"host,required"`
	LastSeenAt        time.Time                          `json:"last_seen_at,required" format:"date-time"`
	Name              string                             `json:"name,required"`
	Type              CookieGetResponseType              `json:"type,required"`
	DomainAttribute   string                             `json:"domain_attribute"`
	ExpiresAttribute  time.Time                          `json:"expires_attribute" format:"date-time"`
	HTTPOnlyAttribute bool                               `json:"http_only_attribute"`
	MaxAgeAttribute   int64                              `json:"max_age_attribute"`
	PageURLs          []string                           `json:"page_urls"`
	PathAttribute     string                             `json:"path_attribute"`
	SameSiteAttribute CookieGetResponseSameSiteAttribute `json:"same_site_attribute"`
	SecureAttribute   bool                               `json:"secure_attribute"`
	JSON              cookieGetResponseJSON              `json:"-"`
}

// cookieGetResponseJSON contains the JSON metadata for the struct
// [CookieGetResponse]
type cookieGetResponseJSON struct {
	ID                apijson.Field
	FirstSeenAt       apijson.Field
	Host              apijson.Field
	LastSeenAt        apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	DomainAttribute   apijson.Field
	ExpiresAttribute  apijson.Field
	HTTPOnlyAttribute apijson.Field
	MaxAgeAttribute   apijson.Field
	PageURLs          apijson.Field
	PathAttribute     apijson.Field
	SameSiteAttribute apijson.Field
	SecureAttribute   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CookieGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieGetResponseJSON) RawJSON() string {
	return r.raw
}

type CookieGetResponseType string

const (
	CookieGetResponseTypeFirstParty CookieGetResponseType = "first_party"
	CookieGetResponseTypeUnknown    CookieGetResponseType = "unknown"
)

func (r CookieGetResponseType) IsKnown() bool {
	switch r {
	case CookieGetResponseTypeFirstParty, CookieGetResponseTypeUnknown:
		return true
	}
	return false
}

type CookieGetResponseSameSiteAttribute string

const (
	CookieGetResponseSameSiteAttributeLax    CookieGetResponseSameSiteAttribute = "lax"
	CookieGetResponseSameSiteAttributeStrict CookieGetResponseSameSiteAttribute = "strict"
	CookieGetResponseSameSiteAttributeNone   CookieGetResponseSameSiteAttribute = "none"
)

func (r CookieGetResponseSameSiteAttribute) IsKnown() bool {
	switch r {
	case CookieGetResponseSameSiteAttributeLax, CookieGetResponseSameSiteAttributeStrict, CookieGetResponseSameSiteAttributeNone:
		return true
	}
	return false
}

type CookieListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned cookies.'
	Direction param.Field[CookieListParamsDirection] `query:"direction"`
	// Filters the returned cookies that match the specified domain attribute
	Domain param.Field[string] `query:"domain"`
	// Export the list of cookies as a file.
	Export param.Field[CookieListParamsExport] `query:"export"`
	// Includes cookies that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// Filters the returned cookies that are set with HttpOnly
	HTTPOnly param.Field[bool] `query:"http_only"`
	// Filters the returned cookies that match the specified name. Wildcards are
	// supported at the start and end to support starts with, ends with and contains.
	// e.g. session\*
	Name param.Field[string] `query:"name"`
	// The field used to sort returned cookies.
	OrderBy param.Field[CookieListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the cookies with the applied filters in a single page. This feature
	// is best-effort and it may only work for zones with a low number of cookies
	Page param.Field[string] `query:"page"`
	// Includes connections that match one or more page URLs (separated by commas)
	// where they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// Filters the returned cookies that match the specified path attribute
	Path param.Field[string] `query:"path"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filters the returned cookies that match the specified same_site attribute
	SameSite param.Field[CookieListParamsSameSite] `query:"same_site"`
	// Filters the returned cookies that are set with Secure
	Secure param.Field[bool] `query:"secure"`
	// Filters the returned cookies that match the specified type attribute
	Type param.Field[CookieListParamsType] `query:"type"`
}

// URLQuery serializes [CookieListParams]'s query parameters as `url.Values`.
func (r CookieListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned cookies.'
type CookieListParamsDirection string

const (
	CookieListParamsDirectionAsc  CookieListParamsDirection = "asc"
	CookieListParamsDirectionDesc CookieListParamsDirection = "desc"
)

func (r CookieListParamsDirection) IsKnown() bool {
	switch r {
	case CookieListParamsDirectionAsc, CookieListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of cookies as a file.
type CookieListParamsExport string

const (
	CookieListParamsExportCsv CookieListParamsExport = "csv"
)

func (r CookieListParamsExport) IsKnown() bool {
	switch r {
	case CookieListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned cookies.
type CookieListParamsOrderBy string

const (
	CookieListParamsOrderByFirstSeenAt CookieListParamsOrderBy = "first_seen_at"
	CookieListParamsOrderByLastSeenAt  CookieListParamsOrderBy = "last_seen_at"
)

func (r CookieListParamsOrderBy) IsKnown() bool {
	switch r {
	case CookieListParamsOrderByFirstSeenAt, CookieListParamsOrderByLastSeenAt:
		return true
	}
	return false
}

// Filters the returned cookies that match the specified same_site attribute
type CookieListParamsSameSite string

const (
	CookieListParamsSameSiteLax    CookieListParamsSameSite = "lax"
	CookieListParamsSameSiteStrict CookieListParamsSameSite = "strict"
	CookieListParamsSameSiteNone   CookieListParamsSameSite = "none"
)

func (r CookieListParamsSameSite) IsKnown() bool {
	switch r {
	case CookieListParamsSameSiteLax, CookieListParamsSameSiteStrict, CookieListParamsSameSiteNone:
		return true
	}
	return false
}

// Filters the returned cookies that match the specified type attribute
type CookieListParamsType string

const (
	CookieListParamsTypeFirstParty CookieListParamsType = "first_party"
	CookieListParamsTypeUnknown    CookieListParamsType = "unknown"
)

func (r CookieListParamsType) IsKnown() bool {
	switch r {
	case CookieListParamsTypeFirstParty, CookieListParamsTypeUnknown:
		return true
	}
	return false
}

type CookieGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CookieGetResponseEnvelope struct {
	Result CookieGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  CookieGetResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo            `json:"errors"`
	Messages []shared.ResponseInfo            `json:"messages"`
	JSON     cookieGetResponseEnvelopeJSON    `json:"-"`
}

// cookieGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CookieGetResponseEnvelope]
type cookieGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CookieGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CookieGetResponseEnvelopeSuccess bool

const (
	CookieGetResponseEnvelopeSuccessTrue CookieGetResponseEnvelopeSuccess = true
)

func (r CookieGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CookieGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
