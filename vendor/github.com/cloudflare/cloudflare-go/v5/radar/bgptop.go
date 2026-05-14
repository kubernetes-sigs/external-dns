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

// BGPTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPTopService] method instead.
type BGPTopService struct {
	Options []option.RequestOption
	Ases    *BGPTopAseService
}

// NewBGPTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPTopService(opts ...option.RequestOption) (r *BGPTopService) {
	r = &BGPTopService{}
	r.Options = opts
	r.Ases = NewBGPTopAseService(opts...)
	return
}

// Retrieves the top network prefixes by BGP updates.
func (r *BGPTopService) Prefixes(ctx context.Context, query BGPTopPrefixesParams, opts ...option.RequestOption) (res *BGPTopPrefixesResponse, err error) {
	var env BGPTopPrefixesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPTopPrefixesResponse struct {
	Meta BGPTopPrefixesResponseMeta   `json:"meta,required"`
	Top0 []BGPTopPrefixesResponseTop0 `json:"top_0,required"`
	JSON bgpTopPrefixesResponseJSON   `json:"-"`
}

// bgpTopPrefixesResponseJSON contains the JSON metadata for the struct
// [BGPTopPrefixesResponse]
type bgpTopPrefixesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopPrefixesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPTopPrefixesResponseMeta struct {
	DateRange []BGPTopPrefixesResponseMetaDateRange `json:"dateRange,required"`
	JSON      bgpTopPrefixesResponseMetaJSON        `json:"-"`
}

// bgpTopPrefixesResponseMetaJSON contains the JSON metadata for the struct
// [BGPTopPrefixesResponseMeta]
type bgpTopPrefixesResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopPrefixesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopPrefixesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPTopPrefixesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      bgpTopPrefixesResponseMetaDateRangeJSON `json:"-"`
}

// bgpTopPrefixesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BGPTopPrefixesResponseMetaDateRange]
type bgpTopPrefixesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopPrefixesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopPrefixesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BGPTopPrefixesResponseTop0 struct {
	Prefix string `json:"prefix,required"`
	// A numeric string.
	Value string                         `json:"value,required"`
	JSON  bgpTopPrefixesResponseTop0JSON `json:"-"`
}

// bgpTopPrefixesResponseTop0JSON contains the JSON metadata for the struct
// [BGPTopPrefixesResponseTop0]
type bgpTopPrefixesResponseTop0JSON struct {
	Prefix      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopPrefixesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopPrefixesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type BGPTopPrefixesParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BGPTopPrefixesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by BGP update type.
	UpdateType param.Field[[]BGPTopPrefixesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [BGPTopPrefixesParams]'s query parameters as `url.Values`.
func (r BGPTopPrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BGPTopPrefixesParamsFormat string

const (
	BGPTopPrefixesParamsFormatJson BGPTopPrefixesParamsFormat = "JSON"
	BGPTopPrefixesParamsFormatCsv  BGPTopPrefixesParamsFormat = "CSV"
)

func (r BGPTopPrefixesParamsFormat) IsKnown() bool {
	switch r {
	case BGPTopPrefixesParamsFormatJson, BGPTopPrefixesParamsFormatCsv:
		return true
	}
	return false
}

type BGPTopPrefixesParamsUpdateType string

const (
	BGPTopPrefixesParamsUpdateTypeAnnouncement BGPTopPrefixesParamsUpdateType = "ANNOUNCEMENT"
	BGPTopPrefixesParamsUpdateTypeWithdrawal   BGPTopPrefixesParamsUpdateType = "WITHDRAWAL"
)

func (r BGPTopPrefixesParamsUpdateType) IsKnown() bool {
	switch r {
	case BGPTopPrefixesParamsUpdateTypeAnnouncement, BGPTopPrefixesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}

type BGPTopPrefixesResponseEnvelope struct {
	Result  BGPTopPrefixesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    bgpTopPrefixesResponseEnvelopeJSON `json:"-"`
}

// bgpTopPrefixesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPTopPrefixesResponseEnvelope]
type bgpTopPrefixesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTopPrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTopPrefixesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
