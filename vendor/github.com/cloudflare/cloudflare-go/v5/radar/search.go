// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SearchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Searches for locations, autonomous systems, reports, bots, certificate logs, and
// certificate authorities.
func (r *SearchService) Global(ctx context.Context, query SearchGlobalParams, opts ...option.RequestOption) (res *SearchGlobalResponse, err error) {
	var env SearchGlobalResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/search/global"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SearchGlobalResponse struct {
	Search []SearchGlobalResponseSearch `json:"search,required"`
	JSON   searchGlobalResponseJSON     `json:"-"`
}

// searchGlobalResponseJSON contains the JSON metadata for the struct
// [SearchGlobalResponse]
type searchGlobalResponseJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGlobalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGlobalResponseJSON) RawJSON() string {
	return r.raw
}

type SearchGlobalResponseSearch struct {
	Code string                         `json:"code,required"`
	Name string                         `json:"name,required"`
	Type string                         `json:"type,required"`
	JSON searchGlobalResponseSearchJSON `json:"-"`
}

// searchGlobalResponseSearchJSON contains the JSON metadata for the struct
// [SearchGlobalResponseSearch]
type searchGlobalResponseSearchJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGlobalResponseSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGlobalResponseSearchJSON) RawJSON() string {
	return r.raw
}

type SearchGlobalParams struct {
	// String used to perform the search operation.
	Query param.Field[string] `query:"query,required"`
	// Search types excluded from results.
	Exclude param.Field[[]SearchGlobalParamsExclude] `query:"exclude"`
	// Format in which results will be returned.
	Format param.Field[SearchGlobalParamsFormat] `query:"format"`
	// Search types included in results.
	Include param.Field[[]SearchGlobalParamsInclude] `query:"include"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Limits the number of objects per search category.
	LimitPerGroup param.Field[float64] `query:"limitPerGroup"`
}

// URLQuery serializes [SearchGlobalParams]'s query parameters as `url.Values`.
func (r SearchGlobalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SearchGlobalParamsExclude string

const (
	SearchGlobalParamsExcludeASNs                   SearchGlobalParamsExclude = "ASNS"
	SearchGlobalParamsExcludeBots                   SearchGlobalParamsExclude = "BOTS"
	SearchGlobalParamsExcludeCertificateAuthorities SearchGlobalParamsExclude = "CERTIFICATE_AUTHORITIES"
	SearchGlobalParamsExcludeCertificateLogs        SearchGlobalParamsExclude = "CERTIFICATE_LOGS"
	SearchGlobalParamsExcludeLocations              SearchGlobalParamsExclude = "LOCATIONS"
	SearchGlobalParamsExcludeNotebooks              SearchGlobalParamsExclude = "NOTEBOOKS"
)

func (r SearchGlobalParamsExclude) IsKnown() bool {
	switch r {
	case SearchGlobalParamsExcludeASNs, SearchGlobalParamsExcludeBots, SearchGlobalParamsExcludeCertificateAuthorities, SearchGlobalParamsExcludeCertificateLogs, SearchGlobalParamsExcludeLocations, SearchGlobalParamsExcludeNotebooks:
		return true
	}
	return false
}

// Format in which results will be returned.
type SearchGlobalParamsFormat string

const (
	SearchGlobalParamsFormatJson SearchGlobalParamsFormat = "JSON"
	SearchGlobalParamsFormatCsv  SearchGlobalParamsFormat = "CSV"
)

func (r SearchGlobalParamsFormat) IsKnown() bool {
	switch r {
	case SearchGlobalParamsFormatJson, SearchGlobalParamsFormatCsv:
		return true
	}
	return false
}

type SearchGlobalParamsInclude string

const (
	SearchGlobalParamsIncludeASNs                   SearchGlobalParamsInclude = "ASNS"
	SearchGlobalParamsIncludeBots                   SearchGlobalParamsInclude = "BOTS"
	SearchGlobalParamsIncludeCertificateAuthorities SearchGlobalParamsInclude = "CERTIFICATE_AUTHORITIES"
	SearchGlobalParamsIncludeCertificateLogs        SearchGlobalParamsInclude = "CERTIFICATE_LOGS"
	SearchGlobalParamsIncludeLocations              SearchGlobalParamsInclude = "LOCATIONS"
	SearchGlobalParamsIncludeNotebooks              SearchGlobalParamsInclude = "NOTEBOOKS"
)

func (r SearchGlobalParamsInclude) IsKnown() bool {
	switch r {
	case SearchGlobalParamsIncludeASNs, SearchGlobalParamsIncludeBots, SearchGlobalParamsIncludeCertificateAuthorities, SearchGlobalParamsIncludeCertificateLogs, SearchGlobalParamsIncludeLocations, SearchGlobalParamsIncludeNotebooks:
		return true
	}
	return false
}

type SearchGlobalResponseEnvelope struct {
	Result  SearchGlobalResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    searchGlobalResponseEnvelopeJSON `json:"-"`
}

// searchGlobalResponseEnvelopeJSON contains the JSON metadata for the struct
// [SearchGlobalResponseEnvelope]
type searchGlobalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGlobalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGlobalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
