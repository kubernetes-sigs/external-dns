// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

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
)

// RankingDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRankingDomainService] method instead.
type RankingDomainService struct {
	Options []option.RequestOption
}

// NewRankingDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRankingDomainService(opts ...option.RequestOption) (r *RankingDomainService) {
	r = &RankingDomainService{}
	r.Options = opts
	return
}

// Retrieves domain rank details. Cloudflare provides an ordered rank for the top
// 100 domains, but for the remainder it only provides ranking buckets like top 200
// thousand, top one million, etc.. These are available through Radar datasets
// endpoints.
func (r *RankingDomainService) Get(ctx context.Context, domain string, query RankingDomainGetParams, opts ...option.RequestOption) (res *RankingDomainGetResponse, err error) {
	var env RankingDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("radar/ranking/domain/%s", domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RankingDomainGetResponse struct {
	Details0 RankingDomainGetResponseDetails0 `json:"details_0,required"`
	Meta     RankingDomainGetResponseMeta     `json:"meta,required"`
	JSON     rankingDomainGetResponseJSON     `json:"-"`
}

// rankingDomainGetResponseJSON contains the JSON metadata for the struct
// [RankingDomainGetResponse]
type rankingDomainGetResponseJSON struct {
	Details0    apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetResponseDetails0 struct {
	Categories []RankingDomainGetResponseDetails0Category `json:"categories,required"`
	// Only available in POPULAR ranking for the most recent ranking.
	Bucket       string                                        `json:"bucket"`
	Rank         int64                                         `json:"rank"`
	TopLocations []RankingDomainGetResponseDetails0TopLocation `json:"top_locations"`
	JSON         rankingDomainGetResponseDetails0JSON          `json:"-"`
}

// rankingDomainGetResponseDetails0JSON contains the JSON metadata for the struct
// [RankingDomainGetResponseDetails0]
type rankingDomainGetResponseDetails0JSON struct {
	Categories   apijson.Field
	Bucket       apijson.Field
	Rank         apijson.Field
	TopLocations apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RankingDomainGetResponseDetails0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseDetails0JSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetResponseDetails0Category struct {
	ID              int64                                        `json:"id,required"`
	Name            string                                       `json:"name,required"`
	SuperCategoryID int64                                        `json:"superCategoryId,required"`
	JSON            rankingDomainGetResponseDetails0CategoryJSON `json:"-"`
}

// rankingDomainGetResponseDetails0CategoryJSON contains the JSON metadata for the
// struct [RankingDomainGetResponseDetails0Category]
type rankingDomainGetResponseDetails0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RankingDomainGetResponseDetails0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseDetails0CategoryJSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetResponseDetails0TopLocation struct {
	LocationCode string                                          `json:"locationCode,required"`
	LocationName string                                          `json:"locationName,required"`
	Rank         int64                                           `json:"rank,required"`
	JSON         rankingDomainGetResponseDetails0TopLocationJSON `json:"-"`
}

// rankingDomainGetResponseDetails0TopLocationJSON contains the JSON metadata for
// the struct [RankingDomainGetResponseDetails0TopLocation]
type rankingDomainGetResponseDetails0TopLocationJSON struct {
	LocationCode apijson.Field
	LocationName apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RankingDomainGetResponseDetails0TopLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseDetails0TopLocationJSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetResponseMeta struct {
	DateRange []RankingDomainGetResponseMetaDateRange `json:"dateRange,required"`
	JSON      rankingDomainGetResponseMetaJSON        `json:"-"`
}

// rankingDomainGetResponseMetaJSON contains the JSON metadata for the struct
// [RankingDomainGetResponseMeta]
type rankingDomainGetResponseMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingDomainGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                 `json:"startTime,required" format:"date-time"`
	JSON      rankingDomainGetResponseMetaDateRangeJSON `json:"-"`
}

// rankingDomainGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RankingDomainGetResponseMetaDateRange]
type rankingDomainGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingDomainGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RankingDomainGetParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RankingDomainGetParamsFormat] `query:"format"`
	// Includes top locations in the response.
	IncludeTopLocations param.Field[bool] `query:"includeTopLocations"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RankingDomainGetParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RankingDomainGetParams]'s query parameters as `url.Values`.
func (r RankingDomainGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingDomainGetParamsFormat string

const (
	RankingDomainGetParamsFormatJson RankingDomainGetParamsFormat = "JSON"
	RankingDomainGetParamsFormatCsv  RankingDomainGetParamsFormat = "CSV"
)

func (r RankingDomainGetParamsFormat) IsKnown() bool {
	switch r {
	case RankingDomainGetParamsFormatJson, RankingDomainGetParamsFormatCsv:
		return true
	}
	return false
}

// The ranking type.
type RankingDomainGetParamsRankingType string

const (
	RankingDomainGetParamsRankingTypePopular        RankingDomainGetParamsRankingType = "POPULAR"
	RankingDomainGetParamsRankingTypeTrendingRise   RankingDomainGetParamsRankingType = "TRENDING_RISE"
	RankingDomainGetParamsRankingTypeTrendingSteady RankingDomainGetParamsRankingType = "TRENDING_STEADY"
)

func (r RankingDomainGetParamsRankingType) IsKnown() bool {
	switch r {
	case RankingDomainGetParamsRankingTypePopular, RankingDomainGetParamsRankingTypeTrendingRise, RankingDomainGetParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RankingDomainGetResponseEnvelope struct {
	Result  RankingDomainGetResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    rankingDomainGetResponseEnvelopeJSON `json:"-"`
}

// rankingDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RankingDomainGetResponseEnvelope]
type rankingDomainGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
