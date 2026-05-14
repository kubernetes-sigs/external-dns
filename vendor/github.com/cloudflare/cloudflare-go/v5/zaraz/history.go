// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

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
)

// HistoryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoryService] method instead.
type HistoryService struct {
	Options []option.RequestOption
	Configs *HistoryConfigService
}

// NewHistoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHistoryService(opts ...option.RequestOption) (r *HistoryService) {
	r = &HistoryService{}
	r.Options = opts
	r.Configs = NewHistoryConfigService(opts...)
	return
}

// Restores a historical published Zaraz configuration by ID for a zone.
func (r *HistoryService) Update(ctx context.Context, params HistoryUpdateParams, opts ...option.RequestOption) (res *Configuration, err error) {
	var env HistoryUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists a history of published Zaraz configuration records for a zone.
func (r *HistoryService) List(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[HistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", params.ZoneID)
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

// Lists a history of published Zaraz configuration records for a zone.
func (r *HistoryService) ListAutoPaging(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[HistoryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type HistoryListResponse struct {
	// ID of the configuration
	ID int64 `json:"id,required"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Configuration description provided by the user who published this configuration
	Description string `json:"description,required"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                  `json:"userId,required"`
	JSON   historyListResponseJSON `json:"-"`
}

// historyListResponseJSON contains the JSON metadata for the struct
// [HistoryListResponse]
type historyListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyListResponseJSON) RawJSON() string {
	return r.raw
}

type HistoryUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// ID of the Zaraz configuration to restore.
	Body int64 `json:"body,required"`
}

func (r HistoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type HistoryUpdateResponseEnvelope struct {
	Errors   []HistoryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HistoryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result Configuration `json:"result,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	JSON    historyUpdateResponseEnvelopeJSON `json:"-"`
}

// historyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [HistoryUpdateResponseEnvelope]
type historyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HistoryUpdateResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           HistoryUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             historyUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// historyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HistoryUpdateResponseEnvelopeErrors]
type historyUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HistoryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HistoryUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    historyUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// historyUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [HistoryUpdateResponseEnvelopeErrorsSource]
type historyUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HistoryUpdateResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           HistoryUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             historyUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// historyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HistoryUpdateResponseEnvelopeMessages]
type historyUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HistoryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HistoryUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    historyUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// historyUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [HistoryUpdateResponseEnvelopeMessagesSource]
type historyUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r historyUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type HistoryListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum amount of results to list. Default value is 10.
	Limit param.Field[int64] `query:"limit"`
	// Ordinal number to start listing the results with. Default value is 0.
	Offset param.Field[int64] `query:"offset"`
	// The field to sort by. Default is updated_at.
	SortField param.Field[HistoryListParamsSortField] `query:"sortField"`
	// Sorting order. Default is DESC.
	SortOrder param.Field[HistoryListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [HistoryListParams]'s query parameters as `url.Values`.
func (r HistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The field to sort by. Default is updated_at.
type HistoryListParamsSortField string

const (
	HistoryListParamsSortFieldID          HistoryListParamsSortField = "id"
	HistoryListParamsSortFieldUserID      HistoryListParamsSortField = "user_id"
	HistoryListParamsSortFieldDescription HistoryListParamsSortField = "description"
	HistoryListParamsSortFieldCreatedAt   HistoryListParamsSortField = "created_at"
	HistoryListParamsSortFieldUpdatedAt   HistoryListParamsSortField = "updated_at"
)

func (r HistoryListParamsSortField) IsKnown() bool {
	switch r {
	case HistoryListParamsSortFieldID, HistoryListParamsSortFieldUserID, HistoryListParamsSortFieldDescription, HistoryListParamsSortFieldCreatedAt, HistoryListParamsSortFieldUpdatedAt:
		return true
	}
	return false
}

// Sorting order. Default is DESC.
type HistoryListParamsSortOrder string

const (
	HistoryListParamsSortOrderDesc HistoryListParamsSortOrder = "DESC"
	HistoryListParamsSortOrderAsc  HistoryListParamsSortOrder = "ASC"
)

func (r HistoryListParamsSortOrder) IsKnown() bool {
	switch r {
	case HistoryListParamsSortOrderDesc, HistoryListParamsSortOrderAsc:
		return true
	}
	return false
}
