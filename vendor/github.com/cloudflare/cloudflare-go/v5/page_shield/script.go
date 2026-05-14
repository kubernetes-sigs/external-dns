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

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptService] method instead.
type ScriptService struct {
	Options []option.RequestOption
}

// NewScriptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScriptService(opts ...option.RequestOption) (r *ScriptService) {
	r = &ScriptService{}
	r.Options = opts
	return
}

// Lists all scripts detected by Page Shield.
func (r *ScriptService) List(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Script], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/scripts", params.ZoneID)
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

// Lists all scripts detected by Page Shield.
func (r *ScriptService) ListAutoPaging(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Script] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a script detected by Page Shield by script ID.
func (r *ScriptService) Get(ctx context.Context, scriptID string, query ScriptGetParams, opts ...option.RequestOption) (res *ScriptGetResponse, err error) {
	var env ScriptGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if scriptID == "" {
		err = errors.New("missing required script_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", query.ZoneID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Script struct {
	// Identifier
	ID                    string    `json:"id,required"`
	AddedAt               time.Time `json:"added_at,required" format:"date-time"`
	FirstSeenAt           time.Time `json:"first_seen_at,required" format:"date-time"`
	Host                  string    `json:"host,required"`
	LastSeenAt            time.Time `json:"last_seen_at,required" format:"date-time"`
	URL                   string    `json:"url,required"`
	URLContainsCDNCGIPath bool      `json:"url_contains_cdn_cgi_path,required"`
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore           int64 `json:"dataflow_score,nullable"`
	DomainReportedMalicious bool  `json:"domain_reported_malicious"`
	// The timestamp of when the script was last fetched.
	FetchedAt    string `json:"fetched_at,nullable"`
	FirstPageURL string `json:"first_page_url"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JSIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore             int64    `json:"magecart_score,nullable"`
	MaliciousDomainCategories []string `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string `json:"malicious_url_categories"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore     int64      `json:"obfuscation_score,nullable"`
	PageURLs             []string   `json:"page_urls"`
	URLReportedMalicious bool       `json:"url_reported_malicious"`
	JSON                 scriptJSON `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCDNCGIPath     apijson.Field
	CryptominingScore         apijson.Field
	DataflowScore             apijson.Field
	DomainReportedMalicious   apijson.Field
	FetchedAt                 apijson.Field
	FirstPageURL              apijson.Field
	Hash                      apijson.Field
	JSIntegrityScore          apijson.Field
	MagecartScore             apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	MalwareScore              apijson.Field
	ObfuscationScore          apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

type ScriptGetResponse struct {
	// Identifier
	ID                    string    `json:"id,required"`
	AddedAt               time.Time `json:"added_at,required" format:"date-time"`
	FirstSeenAt           time.Time `json:"first_seen_at,required" format:"date-time"`
	Host                  string    `json:"host,required"`
	LastSeenAt            time.Time `json:"last_seen_at,required" format:"date-time"`
	URL                   string    `json:"url,required"`
	URLContainsCDNCGIPath bool      `json:"url_contains_cdn_cgi_path,required"`
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore           int64 `json:"dataflow_score,nullable"`
	DomainReportedMalicious bool  `json:"domain_reported_malicious"`
	// The timestamp of when the script was last fetched.
	FetchedAt    string `json:"fetched_at,nullable"`
	FirstPageURL string `json:"first_page_url"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JSIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore             int64    `json:"magecart_score,nullable"`
	MaliciousDomainCategories []string `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string `json:"malicious_url_categories"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore     int64                      `json:"obfuscation_score,nullable"`
	PageURLs             []string                   `json:"page_urls"`
	URLReportedMalicious bool                       `json:"url_reported_malicious"`
	Versions             []ScriptGetResponseVersion `json:"versions,nullable"`
	JSON                 scriptGetResponseJSON      `json:"-"`
}

// scriptGetResponseJSON contains the JSON metadata for the struct
// [ScriptGetResponse]
type scriptGetResponseJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCDNCGIPath     apijson.Field
	CryptominingScore         apijson.Field
	DataflowScore             apijson.Field
	DomainReportedMalicious   apijson.Field
	FetchedAt                 apijson.Field
	FirstPageURL              apijson.Field
	Hash                      apijson.Field
	JSIntegrityScore          apijson.Field
	MagecartScore             apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	MalwareScore              apijson.Field
	ObfuscationScore          apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	Versions                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptGetResponseJSON) RawJSON() string {
	return r.raw
}

// The version of the analyzed script.
type ScriptGetResponseVersion struct {
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore int64 `json:"dataflow_score,nullable"`
	// The timestamp of when the script was last fetched.
	FetchedAt string `json:"fetched_at,nullable"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JSIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore int64 `json:"magecart_score,nullable"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore int64                        `json:"obfuscation_score,nullable"`
	JSON             scriptGetResponseVersionJSON `json:"-"`
}

// scriptGetResponseVersionJSON contains the JSON metadata for the struct
// [ScriptGetResponseVersion]
type scriptGetResponseVersionJSON struct {
	CryptominingScore apijson.Field
	DataflowScore     apijson.Field
	FetchedAt         apijson.Field
	Hash              apijson.Field
	JSIntegrityScore  apijson.Field
	MagecartScore     apijson.Field
	MalwareScore      apijson.Field
	ObfuscationScore  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptGetResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptGetResponseVersionJSON) RawJSON() string {
	return r.raw
}

type ScriptListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned scripts.
	Direction param.Field[ScriptListParamsDirection] `query:"direction"`
	// When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts.
	// The default value is true.
	ExcludeCDNCGI param.Field[bool] `query:"exclude_cdn_cgi"`
	// When true, excludes duplicate scripts. We consider a script duplicate of another
	// if their javascript content matches and they share the same url host and zone
	// hostname. In such case, we return the most recent script for the URL host and
	// zone hostname combination.
	ExcludeDuplicates param.Field[bool] `query:"exclude_duplicates"`
	// Excludes scripts whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of scripts as a file.
	Export param.Field[ScriptListParamsExport] `query:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[ScriptListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the scripts with the applied filters in a single page. This feature
	// is best-effort and it may only work for zones with a low number of scripts
	Page param.Field[string] `query:"page"`
	// Includes scripts that match one or more page URLs (separated by commas) where
	// they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious scripts appear first in the returned scripts.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned scripts using a comma-separated list of scripts statuses.
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is
	// `active`.
	Status param.Field[string] `query:"status"`
	// Includes scripts whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes [ScriptListParams]'s query parameters as `url.Values`.
func (r ScriptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned scripts.
type ScriptListParamsDirection string

const (
	ScriptListParamsDirectionAsc  ScriptListParamsDirection = "asc"
	ScriptListParamsDirectionDesc ScriptListParamsDirection = "desc"
)

func (r ScriptListParamsDirection) IsKnown() bool {
	switch r {
	case ScriptListParamsDirectionAsc, ScriptListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of scripts as a file.
type ScriptListParamsExport string

const (
	ScriptListParamsExportCsv ScriptListParamsExport = "csv"
)

func (r ScriptListParamsExport) IsKnown() bool {
	switch r {
	case ScriptListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned scripts.
type ScriptListParamsOrderBy string

const (
	ScriptListParamsOrderByFirstSeenAt ScriptListParamsOrderBy = "first_seen_at"
	ScriptListParamsOrderByLastSeenAt  ScriptListParamsOrderBy = "last_seen_at"
)

func (r ScriptListParamsOrderBy) IsKnown() bool {
	switch r {
	case ScriptListParamsOrderByFirstSeenAt, ScriptListParamsOrderByLastSeenAt:
		return true
	}
	return false
}

type ScriptGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptGetResponseEnvelope struct {
	Result ScriptGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  ScriptGetResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo            `json:"errors"`
	Messages []shared.ResponseInfo            `json:"messages"`
	JSON     scriptGetResponseEnvelopeJSON    `json:"-"`
}

// scriptGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptGetResponseEnvelope]
type scriptGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptGetResponseEnvelopeSuccess bool

const (
	ScriptGetResponseEnvelopeSuccessTrue ScriptGetResponseEnvelopeSuccess = true
)

func (r ScriptGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
