// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filters

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// FilterService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFilterService] method instead.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
type FilterService struct {
	Options []option.RequestOption
}

// NewFilterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFilterService(opts ...option.RequestOption) (r *FilterService) {
	r = &FilterService{}
	r.Options = opts
	return
}

// Creates one or more filters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) New(ctx context.Context, params FilterNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallFilter], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Creates one or more filters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) NewAutoPaging(ctx context.Context, params FilterNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallFilter] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

// Updates an existing filter.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) Update(ctx context.Context, filterID string, params FilterUpdateParams, opts ...option.RequestOption) (res *FirewallFilter, err error) {
	var env FilterUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", params.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) List(ctx context.Context, params FilterListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallFilter], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", params.ZoneID)
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

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) ListAutoPaging(ctx context.Context, params FilterListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallFilter] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing filter.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) Delete(ctx context.Context, filterID string, body FilterDeleteParams, opts ...option.RequestOption) (res *FilterDeleteResponse, err error) {
	var env FilterDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", body.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes one or more existing filters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) BulkDelete(ctx context.Context, params FilterBulkDeleteParams, opts ...option.RequestOption) (res *[]FilterBulkDeleteResponse, err error) {
	var env FilterBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more existing filters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) BulkUpdate(ctx context.Context, params FilterBulkUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallFilter], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Updates one or more existing filters.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) BulkUpdateAutoPaging(ctx context.Context, params FilterBulkUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallFilter] {
	return pagination.NewSinglePageAutoPager(r.BulkUpdate(ctx, params, opts...))
}

// Fetches the details of a filter.
//
// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *FilterService) Get(ctx context.Context, filterID string, query FilterGetParams, opts ...option.RequestOption) (res *FirewallFilter, err error) {
	var env FilterGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", query.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string             `json:"ref"`
	JSON firewallFilterJSON `json:"-"`
}

// firewallFilterJSON contains the JSON metadata for the struct [FirewallFilter]
type firewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallFilter) ImplementsFirewallRuleFilter() {}

type FirewallFilterParam struct {
	// An informative summary of the filter.
	Description param.Field[string] `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref param.Field[string] `json:"ref"`
}

func (r FirewallFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FilterDeleteResponse struct {
	// The unique identifier of the filter.
	ID   string                   `json:"id,required"`
	JSON filterDeleteResponseJSON `json:"-"`
}

// filterDeleteResponseJSON contains the JSON metadata for the struct
// [FilterDeleteResponse]
type filterDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FilterBulkDeleteResponse struct {
	// The unique identifier of the filter.
	ID   string                       `json:"id"`
	JSON filterBulkDeleteResponseJSON `json:"-"`
}

// filterBulkDeleteResponseJSON contains the JSON metadata for the struct
// [FilterBulkDeleteResponse]
type filterBulkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FilterNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string]   `path:"zone_id,required"`
	Body   []FirewallFilterParam `json:"body,required"`
}

func (r FilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterUpdateParams struct {
	// Defines an identifier.
	ZoneID         param.Field[string] `path:"zone_id,required"`
	FirewallFilter FirewallFilterParam `json:"firewall_filter,required"`
}

func (r FilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.FirewallFilter)
}

type FilterUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallFilter        `json:"result,required"`
	// Defines whether the API call was successful.
	Success FilterUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterUpdateResponseEnvelopeJSON    `json:"-"`
}

// filterUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterUpdateResponseEnvelope]
type filterUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FilterUpdateResponseEnvelopeSuccess bool

const (
	FilterUpdateResponseEnvelopeSuccessTrue FilterUpdateResponseEnvelopeSuccess = true
)

func (r FilterUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FilterListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The unique identifier of the filter.
	ID param.Field[string] `query:"id"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// A case-insensitive string to find in the expression.
	Expression param.Field[string] `query:"expression"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of filters per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The filter ref (a short reference tag) to search for. Must be an exact match.
	Ref param.Field[string] `query:"ref"`
}

// URLQuery serializes [FilterListParams]'s query parameters as `url.Values`.
func (r FilterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FilterDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FilterDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FilterDeleteResponse  `json:"result,required"`
	// Defines whether the API call was successful.
	Success FilterDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterDeleteResponseEnvelopeJSON    `json:"-"`
}

// filterDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterDeleteResponseEnvelope]
type filterDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FilterDeleteResponseEnvelopeSuccess bool

const (
	FilterDeleteResponseEnvelopeSuccessTrue FilterDeleteResponseEnvelopeSuccess = true
)

func (r FilterDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FilterBulkDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string]   `path:"zone_id,required"`
	ID     param.Field[[]string] `query:"id,required"`
}

// URLQuery serializes [FilterBulkDeleteParams]'s query parameters as `url.Values`.
func (r FilterBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FilterBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   []FilterBulkDeleteResponse `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    FilterBulkDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FilterBulkDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       filterBulkDeleteResponseEnvelopeJSON       `json:"-"`
}

// filterBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterBulkDeleteResponseEnvelope]
type filterBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FilterBulkDeleteResponseEnvelopeSuccess bool

const (
	FilterBulkDeleteResponseEnvelopeSuccessTrue FilterBulkDeleteResponseEnvelopeSuccess = true
)

func (r FilterBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FilterBulkDeleteResponseEnvelopeResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       filterBulkDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// filterBulkDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [FilterBulkDeleteResponseEnvelopeResultInfo]
type filterBulkDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterBulkDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterBulkDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type FilterBulkUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string]          `path:"zone_id,required"`
	Body   []FilterBulkUpdateParamsBody `json:"body,required"`
}

func (r FilterBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterBulkUpdateParamsBody struct {
	// An informative summary of the filter.
	Description param.Field[string] `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref param.Field[string] `json:"ref"`
}

func (r FilterBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FilterGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FilterGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallFilter        `json:"result,required"`
	// Defines whether the API call was successful.
	Success FilterGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterGetResponseEnvelopeJSON    `json:"-"`
}

// filterGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterGetResponseEnvelope]
type filterGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FilterGetResponseEnvelopeSuccess bool

const (
	FilterGetResponseEnvelopeSuccessTrue FilterGetResponseEnvelopeSuccess = true
)

func (r FilterGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
