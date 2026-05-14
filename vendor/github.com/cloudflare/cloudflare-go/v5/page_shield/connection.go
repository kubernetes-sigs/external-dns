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

// ConnectionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// Lists all connections detected by Page Shield.
func (r *ConnectionService) List(ctx context.Context, params ConnectionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Connection], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/connections", params.ZoneID)
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

// Lists all connections detected by Page Shield.
func (r *ConnectionService) ListAutoPaging(ctx context.Context, params ConnectionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Connection] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *ConnectionService) Get(ctx context.Context, connectionID string, query ConnectionGetParams, opts ...option.RequestOption) (res *Connection, err error) {
	var env ConnectionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", query.ZoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Connection struct {
	// Identifier
	ID                        string         `json:"id,required"`
	AddedAt                   time.Time      `json:"added_at,required" format:"date-time"`
	FirstSeenAt               time.Time      `json:"first_seen_at,required" format:"date-time"`
	Host                      string         `json:"host,required"`
	LastSeenAt                time.Time      `json:"last_seen_at,required" format:"date-time"`
	URL                       string         `json:"url,required"`
	URLContainsCDNCGIPath     bool           `json:"url_contains_cdn_cgi_path,required"`
	DomainReportedMalicious   bool           `json:"domain_reported_malicious"`
	FirstPageURL              string         `json:"first_page_url"`
	MaliciousDomainCategories []string       `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string       `json:"malicious_url_categories"`
	PageURLs                  []string       `json:"page_urls"`
	URLReportedMalicious      bool           `json:"url_reported_malicious"`
	JSON                      connectionJSON `json:"-"`
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCDNCGIPath     apijson.Field
	DomainReportedMalicious   apijson.Field
	FirstPageURL              apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionJSON) RawJSON() string {
	return r.raw
}

type ConnectionListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned connections.
	Direction param.Field[ConnectionListParamsDirection] `query:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned
	// connections. The default value is true.
	ExcludeCDNCGI param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of connections as a file.
	Export param.Field[ConnectionListParamsExport] `query:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned connections.
	OrderBy param.Field[ConnectionListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the connections with the applied filters in a single page. This
	// feature is best-effort and it may only work for zones with a low number of
	// connections
	Page param.Field[string] `query:"page"`
	// Includes connections that match one or more page URLs (separated by commas)
	// where they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious connections appear first in the returned connections.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned connections using a comma-separated list of connection
	// statuses. Accepted values: `active`, `infrequent`, and `inactive`. The default
	// value is `active`.
	Status param.Field[string] `query:"status"`
	// Includes connections whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned connections.
type ConnectionListParamsDirection string

const (
	ConnectionListParamsDirectionAsc  ConnectionListParamsDirection = "asc"
	ConnectionListParamsDirectionDesc ConnectionListParamsDirection = "desc"
)

func (r ConnectionListParamsDirection) IsKnown() bool {
	switch r {
	case ConnectionListParamsDirectionAsc, ConnectionListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of connections as a file.
type ConnectionListParamsExport string

const (
	ConnectionListParamsExportCsv ConnectionListParamsExport = "csv"
)

func (r ConnectionListParamsExport) IsKnown() bool {
	switch r {
	case ConnectionListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned connections.
type ConnectionListParamsOrderBy string

const (
	ConnectionListParamsOrderByFirstSeenAt ConnectionListParamsOrderBy = "first_seen_at"
	ConnectionListParamsOrderByLastSeenAt  ConnectionListParamsOrderBy = "last_seen_at"
)

func (r ConnectionListParamsOrderBy) IsKnown() bool {
	switch r {
	case ConnectionListParamsOrderByFirstSeenAt, ConnectionListParamsOrderByLastSeenAt:
		return true
	}
	return false
}

type ConnectionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ConnectionGetResponseEnvelope struct {
	Result Connection `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  ConnectionGetResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo                `json:"errors"`
	Messages []shared.ResponseInfo                `json:"messages"`
	JSON     connectionGetResponseEnvelopeJSON    `json:"-"`
}

// connectionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectionGetResponseEnvelope]
type connectionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConnectionGetResponseEnvelopeSuccess bool

const (
	ConnectionGetResponseEnvelopeSuccessTrue ConnectionGetResponseEnvelopeSuccess = true
)

func (r ConnectionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConnectionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
