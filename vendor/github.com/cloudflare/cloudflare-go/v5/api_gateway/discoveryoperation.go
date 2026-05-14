// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

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
)

// DiscoveryOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscoveryOperationService] method instead.
type DiscoveryOperationService struct {
	Options []option.RequestOption
}

// NewDiscoveryOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDiscoveryOperationService(opts ...option.RequestOption) (r *DiscoveryOperationService) {
	r = &DiscoveryOperationService{}
	r.Options = opts
	return
}

// Retrieve the most up to date view of discovered operations
func (r *DiscoveryOperationService) List(ctx context.Context, params DiscoveryOperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DiscoveryOperation], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", params.ZoneID)
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

// Retrieve the most up to date view of discovered operations
func (r *DiscoveryOperationService) ListAutoPaging(ctx context.Context, params DiscoveryOperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DiscoveryOperation] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Update the `state` on one or more discovered operations
func (r *DiscoveryOperationService) BulkEdit(ctx context.Context, params DiscoveryOperationBulkEditParams, opts ...option.RequestOption) (res *DiscoveryOperationBulkEditResponse, err error) {
	var env DiscoveryOperationBulkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the `state` on a discovered operation
func (r *DiscoveryOperationService) Edit(ctx context.Context, operationID string, params DiscoveryOperationEditParams, opts ...option.RequestOption) (res *DiscoveryOperationEditResponse, err error) {
	var env DiscoveryOperationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations/%s", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiscoveryOperationBulkEditResponse map[string]DiscoveryOperationBulkEditResponseItem

// Mappings of discovered operations (keys) to objects describing their state
type DiscoveryOperationBulkEditResponseItem struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State DiscoveryOperationBulkEditResponseItemState `json:"state"`
	JSON  discoveryOperationBulkEditResponseItemJSON  `json:"-"`
}

// discoveryOperationBulkEditResponseItemJSON contains the JSON metadata for the
// struct [DiscoveryOperationBulkEditResponseItem]
type discoveryOperationBulkEditResponseItemJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationBulkEditResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationBulkEditResponseItemJSON) RawJSON() string {
	return r.raw
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationBulkEditResponseItemState string

const (
	DiscoveryOperationBulkEditResponseItemStateReview  DiscoveryOperationBulkEditResponseItemState = "review"
	DiscoveryOperationBulkEditResponseItemStateIgnored DiscoveryOperationBulkEditResponseItemState = "ignored"
)

func (r DiscoveryOperationBulkEditResponseItemState) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditResponseItemStateReview, DiscoveryOperationBulkEditResponseItemStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationEditResponse struct {
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State DiscoveryOperationEditResponseState `json:"state"`
	JSON  discoveryOperationEditResponseJSON  `json:"-"`
}

// discoveryOperationEditResponseJSON contains the JSON metadata for the struct
// [DiscoveryOperationEditResponse]
type discoveryOperationEditResponseJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationEditResponseJSON) RawJSON() string {
	return r.raw
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type DiscoveryOperationEditResponseState string

const (
	DiscoveryOperationEditResponseStateReview  DiscoveryOperationEditResponseState = "review"
	DiscoveryOperationEditResponseStateSaved   DiscoveryOperationEditResponseState = "saved"
	DiscoveryOperationEditResponseStateIgnored DiscoveryOperationEditResponseState = "ignored"
)

func (r DiscoveryOperationEditResponseState) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditResponseStateReview, DiscoveryOperationEditResponseStateSaved, DiscoveryOperationEditResponseStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// When `true`, only return API Discovery results that are not saved into API
	// Shield Endpoint Management
	Diff param.Field[bool] `query:"diff"`
	// Direction to order results.
	Direction param.Field[DiscoveryOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by
	Order param.Field[DiscoveryOperationListParamsOrder] `query:"order"`
	// Filter results to only include discovery results sourced from a particular
	// discovery engine
	//
	//   - `ML` - Discovered operations that were sourced using ML API Discovery
	//   - `SessionIdentifier` - Discovered operations that were sourced using Session
	//     Identifier API Discovery
	Origin param.Field[DiscoveryOperationListParamsOrigin] `query:"origin"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter results to only include discovery results in a particular state. States
	// are as follows
	//
	//   - `review` - Discovered operations that are not saved into API Shield Endpoint
	//     Management
	//   - `saved` - Discovered operations that are already saved into API Shield
	//     Endpoint Management
	//   - `ignored` - Discovered operations that have been marked as ignored
	State param.Field[DiscoveryOperationListParamsState] `query:"state"`
}

// URLQuery serializes [DiscoveryOperationListParams]'s query parameters as
// `url.Values`.
func (r DiscoveryOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type DiscoveryOperationListParamsDirection string

const (
	DiscoveryOperationListParamsDirectionAsc  DiscoveryOperationListParamsDirection = "asc"
	DiscoveryOperationListParamsDirectionDesc DiscoveryOperationListParamsDirection = "desc"
)

func (r DiscoveryOperationListParamsDirection) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsDirectionAsc, DiscoveryOperationListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order by
type DiscoveryOperationListParamsOrder string

const (
	DiscoveryOperationListParamsOrderHost                    DiscoveryOperationListParamsOrder = "host"
	DiscoveryOperationListParamsOrderMethod                  DiscoveryOperationListParamsOrder = "method"
	DiscoveryOperationListParamsOrderEndpoint                DiscoveryOperationListParamsOrder = "endpoint"
	DiscoveryOperationListParamsOrderTrafficStatsRequests    DiscoveryOperationListParamsOrder = "traffic_stats.requests"
	DiscoveryOperationListParamsOrderTrafficStatsLastUpdated DiscoveryOperationListParamsOrder = "traffic_stats.last_updated"
)

func (r DiscoveryOperationListParamsOrder) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsOrderHost, DiscoveryOperationListParamsOrderMethod, DiscoveryOperationListParamsOrderEndpoint, DiscoveryOperationListParamsOrderTrafficStatsRequests, DiscoveryOperationListParamsOrderTrafficStatsLastUpdated:
		return true
	}
	return false
}

// Filter results to only include discovery results sourced from a particular
// discovery engine
//
//   - `ML` - Discovered operations that were sourced using ML API Discovery
//   - `SessionIdentifier` - Discovered operations that were sourced using Session
//     Identifier API Discovery
type DiscoveryOperationListParamsOrigin string

const (
	DiscoveryOperationListParamsOriginMl                DiscoveryOperationListParamsOrigin = "ML"
	DiscoveryOperationListParamsOriginSessionIdentifier DiscoveryOperationListParamsOrigin = "SessionIdentifier"
	DiscoveryOperationListParamsOriginLabelDiscovery    DiscoveryOperationListParamsOrigin = "LabelDiscovery"
)

func (r DiscoveryOperationListParamsOrigin) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsOriginMl, DiscoveryOperationListParamsOriginSessionIdentifier, DiscoveryOperationListParamsOriginLabelDiscovery:
		return true
	}
	return false
}

// Filter results to only include discovery results in a particular state. States
// are as follows
//
//   - `review` - Discovered operations that are not saved into API Shield Endpoint
//     Management
//   - `saved` - Discovered operations that are already saved into API Shield
//     Endpoint Management
//   - `ignored` - Discovered operations that have been marked as ignored
type DiscoveryOperationListParamsState string

const (
	DiscoveryOperationListParamsStateReview  DiscoveryOperationListParamsState = "review"
	DiscoveryOperationListParamsStateSaved   DiscoveryOperationListParamsState = "saved"
	DiscoveryOperationListParamsStateIgnored DiscoveryOperationListParamsState = "ignored"
)

func (r DiscoveryOperationListParamsState) IsKnown() bool {
	switch r {
	case DiscoveryOperationListParamsStateReview, DiscoveryOperationListParamsStateSaved, DiscoveryOperationListParamsStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationBulkEditParams struct {
	// Identifier.
	ZoneID param.Field[string]                             `path:"zone_id,required"`
	Body   map[string]DiscoveryOperationBulkEditParamsBody `json:"body,required"`
}

func (r DiscoveryOperationBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Mappings of discovered operations (keys) to objects describing their state
type DiscoveryOperationBulkEditParamsBody struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[DiscoveryOperationBulkEditParamsBodyState] `json:"state"`
}

func (r DiscoveryOperationBulkEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationBulkEditParamsBodyState string

const (
	DiscoveryOperationBulkEditParamsBodyStateReview  DiscoveryOperationBulkEditParamsBodyState = "review"
	DiscoveryOperationBulkEditParamsBodyStateIgnored DiscoveryOperationBulkEditParamsBodyState = "ignored"
)

func (r DiscoveryOperationBulkEditParamsBodyState) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditParamsBodyStateReview, DiscoveryOperationBulkEditParamsBodyStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationBulkEditResponseEnvelope struct {
	Errors   Message                            `json:"errors,required"`
	Messages Message                            `json:"messages,required"`
	Result   DiscoveryOperationBulkEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DiscoveryOperationBulkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryOperationBulkEditResponseEnvelopeJSON    `json:"-"`
}

// discoveryOperationBulkEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [DiscoveryOperationBulkEditResponseEnvelope]
type discoveryOperationBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DiscoveryOperationBulkEditResponseEnvelopeSuccess bool

const (
	DiscoveryOperationBulkEditResponseEnvelopeSuccessTrue DiscoveryOperationBulkEditResponseEnvelopeSuccess = true
)

func (r DiscoveryOperationBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryOperationBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DiscoveryOperationEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[DiscoveryOperationEditParamsState] `json:"state"`
}

func (r DiscoveryOperationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type DiscoveryOperationEditParamsState string

const (
	DiscoveryOperationEditParamsStateReview  DiscoveryOperationEditParamsState = "review"
	DiscoveryOperationEditParamsStateIgnored DiscoveryOperationEditParamsState = "ignored"
)

func (r DiscoveryOperationEditParamsState) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditParamsStateReview, DiscoveryOperationEditParamsStateIgnored:
		return true
	}
	return false
}

type DiscoveryOperationEditResponseEnvelope struct {
	Errors   Message                        `json:"errors,required"`
	Messages Message                        `json:"messages,required"`
	Result   DiscoveryOperationEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success DiscoveryOperationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryOperationEditResponseEnvelopeJSON    `json:"-"`
}

// discoveryOperationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DiscoveryOperationEditResponseEnvelope]
type discoveryOperationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryOperationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryOperationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DiscoveryOperationEditResponseEnvelopeSuccess bool

const (
	DiscoveryOperationEditResponseEnvelopeSuccessTrue DiscoveryOperationEditResponseEnvelopeSuccess = true
)

func (r DiscoveryOperationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryOperationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
