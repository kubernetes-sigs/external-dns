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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// BGPLeakEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPLeakEventService] method instead.
type BGPLeakEventService struct {
	Options []option.RequestOption
}

// NewBGPLeakEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPLeakEventService(opts ...option.RequestOption) (r *BGPLeakEventService) {
	r = &BGPLeakEventService{}
	r.Options = opts
	return
}

// Retrieves the BGP route leak events.
func (r *BGPLeakEventService) List(ctx context.Context, query BGPLeakEventListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[BGPLeakEventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "radar/bgp/leaks/events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieves the BGP route leak events.
func (r *BGPLeakEventService) ListAutoPaging(ctx context.Context, query BGPLeakEventListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[BGPLeakEventListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, query, opts...))
}

type BGPLeakEventListResponse struct {
	ASNInfo []BGPLeakEventListResponseASNInfo `json:"asn_info,required"`
	Events  []BGPLeakEventListResponseEvent   `json:"events,required"`
	JSON    bgpLeakEventListResponseJSON      `json:"-"`
}

// bgpLeakEventListResponseJSON contains the JSON metadata for the struct
// [BGPLeakEventListResponse]
type bgpLeakEventListResponseJSON struct {
	ASNInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventListResponseJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventListResponseASNInfo struct {
	ASN         int64                               `json:"asn,required"`
	CountryCode string                              `json:"country_code,required"`
	OrgName     string                              `json:"org_name,required"`
	JSON        bgpLeakEventListResponseASNInfoJSON `json:"-"`
}

// bgpLeakEventListResponseASNInfoJSON contains the JSON metadata for the struct
// [BGPLeakEventListResponseASNInfo]
type bgpLeakEventListResponseASNInfoJSON struct {
	ASN         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventListResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventListResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventListResponseEvent struct {
	ID          int64                             `json:"id,required"`
	Countries   []string                          `json:"countries,required"`
	DetectedTs  string                            `json:"detected_ts,required"`
	Finished    bool                              `json:"finished,required"`
	LeakASN     int64                             `json:"leak_asn,required"`
	LeakCount   int64                             `json:"leak_count,required"`
	LeakSeg     []int64                           `json:"leak_seg,required"`
	LeakType    int64                             `json:"leak_type,required"`
	MaxTs       string                            `json:"max_ts,required"`
	MinTs       string                            `json:"min_ts,required"`
	OriginCount int64                             `json:"origin_count,required"`
	PeerCount   int64                             `json:"peer_count,required"`
	PrefixCount int64                             `json:"prefix_count,required"`
	JSON        bgpLeakEventListResponseEventJSON `json:"-"`
}

// bgpLeakEventListResponseEventJSON contains the JSON metadata for the struct
// [BGPLeakEventListResponseEvent]
type bgpLeakEventListResponseEventJSON struct {
	ID          apijson.Field
	Countries   apijson.Field
	DetectedTs  apijson.Field
	Finished    apijson.Field
	LeakASN     apijson.Field
	LeakCount   apijson.Field
	LeakSeg     apijson.Field
	LeakType    apijson.Field
	MaxTs       apijson.Field
	MinTs       apijson.Field
	OriginCount apijson.Field
	PeerCount   apijson.Field
	PrefixCount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPLeakEventListResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpLeakEventListResponseEventJSON) RawJSON() string {
	return r.raw
}

type BGPLeakEventListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event.
	EventID param.Field[int64] `query:"eventId"`
	// Format in which results will be returned.
	Format param.Field[BGPLeakEventListParamsFormat] `query:"format"`
	// ASN that is causing or affected by a route leak event.
	InvolvedASN param.Field[int64] `query:"involvedAsn"`
	// Country code of a involved ASN in a route leak event.
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The leaking AS of a route leak event.
	LeakASN param.Field[int64] `query:"leakAsn"`
	// Current page number, starting from 1.
	Page param.Field[int64] `query:"page"`
	// Number of entries per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Sorts results by the specified field.
	SortBy param.Field[BGPLeakEventListParamsSortBy] `query:"sortBy"`
	// Sort order.
	SortOrder param.Field[BGPLeakEventListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [BGPLeakEventListParams]'s query parameters as `url.Values`.
func (r BGPLeakEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BGPLeakEventListParamsFormat string

const (
	BGPLeakEventListParamsFormatJson BGPLeakEventListParamsFormat = "JSON"
	BGPLeakEventListParamsFormatCsv  BGPLeakEventListParamsFormat = "CSV"
)

func (r BGPLeakEventListParamsFormat) IsKnown() bool {
	switch r {
	case BGPLeakEventListParamsFormatJson, BGPLeakEventListParamsFormatCsv:
		return true
	}
	return false
}

// Sorts results by the specified field.
type BGPLeakEventListParamsSortBy string

const (
	BGPLeakEventListParamsSortByID       BGPLeakEventListParamsSortBy = "ID"
	BGPLeakEventListParamsSortByLeaks    BGPLeakEventListParamsSortBy = "LEAKS"
	BGPLeakEventListParamsSortByPeers    BGPLeakEventListParamsSortBy = "PEERS"
	BGPLeakEventListParamsSortByPrefixes BGPLeakEventListParamsSortBy = "PREFIXES"
	BGPLeakEventListParamsSortByOrigins  BGPLeakEventListParamsSortBy = "ORIGINS"
	BGPLeakEventListParamsSortByTime     BGPLeakEventListParamsSortBy = "TIME"
)

func (r BGPLeakEventListParamsSortBy) IsKnown() bool {
	switch r {
	case BGPLeakEventListParamsSortByID, BGPLeakEventListParamsSortByLeaks, BGPLeakEventListParamsSortByPeers, BGPLeakEventListParamsSortByPrefixes, BGPLeakEventListParamsSortByOrigins, BGPLeakEventListParamsSortByTime:
		return true
	}
	return false
}

// Sort order.
type BGPLeakEventListParamsSortOrder string

const (
	BGPLeakEventListParamsSortOrderAsc  BGPLeakEventListParamsSortOrder = "ASC"
	BGPLeakEventListParamsSortOrderDesc BGPLeakEventListParamsSortOrder = "DESC"
)

func (r BGPLeakEventListParamsSortOrder) IsKnown() bool {
	switch r {
	case BGPLeakEventListParamsSortOrderAsc, BGPLeakEventListParamsSortOrderDesc:
		return true
	}
	return false
}
