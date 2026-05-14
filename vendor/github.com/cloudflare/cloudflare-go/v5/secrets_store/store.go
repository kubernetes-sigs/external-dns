// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

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

// StoreService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreService] method instead.
type StoreService struct {
	Options []option.RequestOption
	Secrets *StoreSecretService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r *StoreService) {
	r = &StoreService{}
	r.Options = opts
	r.Secrets = NewStoreSecretService(opts...)
	return
}

// Creates a store in the account
func (r *StoreService) New(ctx context.Context, params StoreNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[StoreNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores", params.AccountID)
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

// Creates a store in the account
func (r *StoreService) NewAutoPaging(ctx context.Context, params StoreNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[StoreNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

// Lists all the stores in an account
func (r *StoreService) List(ctx context.Context, params StoreListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[StoreListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores", params.AccountID)
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

// Lists all the stores in an account
func (r *StoreService) ListAutoPaging(ctx context.Context, params StoreListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[StoreListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a single store
func (r *StoreService) Delete(ctx context.Context, storeID string, body StoreDeleteParams, opts ...option.RequestOption) (res *StoreDeleteResponse, err error) {
	var env StoreDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s", body.AccountID, storeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StoreNewResponse struct {
	// Store Identifier
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the store
	Name string               `json:"name,required"`
	JSON storeNewResponseJSON `json:"-"`
}

// storeNewResponseJSON contains the JSON metadata for the struct
// [StoreNewResponse]
type storeNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeNewResponseJSON) RawJSON() string {
	return r.raw
}

type StoreListResponse struct {
	// Store Identifier
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the store
	Name string                `json:"name,required"`
	JSON storeListResponseJSON `json:"-"`
}

// storeListResponseJSON contains the JSON metadata for the struct
// [StoreListResponse]
type storeListResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeListResponseJSON) RawJSON() string {
	return r.raw
}

type StoreDeleteResponse struct {
	// Store Identifier
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the store
	Name string                  `json:"name,required"`
	JSON storeDeleteResponseJSON `json:"-"`
}

// storeDeleteResponseJSON contains the JSON metadata for the struct
// [StoreDeleteResponse]
type storeDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type StoreNewParams struct {
	// Account Identifier
	AccountID param.Field[string]  `path:"account_id,required"`
	Body      []StoreNewParamsBody `json:"body,required"`
}

func (r StoreNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StoreNewParamsBody struct {
	// The name of the store
	Name param.Field[string] `json:"name,required"`
}

func (r StoreNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StoreListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to sort objects
	Direction param.Field[StoreListParamsDirection] `query:"direction"`
	// Order secrets by values in the given field
	Order param.Field[StoreListParamsOrder] `query:"order"`
	// Page number
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [StoreListParams]'s query parameters as `url.Values`.
func (r StoreListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to sort objects
type StoreListParamsDirection string

const (
	StoreListParamsDirectionAsc  StoreListParamsDirection = "asc"
	StoreListParamsDirectionDesc StoreListParamsDirection = "desc"
)

func (r StoreListParamsDirection) IsKnown() bool {
	switch r {
	case StoreListParamsDirectionAsc, StoreListParamsDirectionDesc:
		return true
	}
	return false
}

// Order secrets by values in the given field
type StoreListParamsOrder string

const (
	StoreListParamsOrderName     StoreListParamsOrder = "name"
	StoreListParamsOrderComment  StoreListParamsOrder = "comment"
	StoreListParamsOrderCreated  StoreListParamsOrder = "created"
	StoreListParamsOrderModified StoreListParamsOrder = "modified"
	StoreListParamsOrderStatus   StoreListParamsOrder = "status"
)

func (r StoreListParamsOrder) IsKnown() bool {
	switch r {
	case StoreListParamsOrderName, StoreListParamsOrderComment, StoreListParamsOrderCreated, StoreListParamsOrderModified, StoreListParamsOrderStatus:
		return true
	}
	return false
}

type StoreDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StoreDeleteResponseEnvelope struct {
	Errors   []StoreDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StoreDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    StoreDeleteResponseEnvelopeSuccess    `json:"success,required"`
	Result     StoreDeleteResponse                   `json:"result"`
	ResultInfo StoreDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storeDeleteResponseEnvelopeJSON       `json:"-"`
}

// storeDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [StoreDeleteResponseEnvelope]
type storeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StoreDeleteResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           StoreDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             storeDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// storeDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StoreDeleteResponseEnvelopeErrors]
type storeDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StoreDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    storeDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// storeDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [StoreDeleteResponseEnvelopeErrorsSource]
type storeDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StoreDeleteResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           StoreDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             storeDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// storeDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StoreDeleteResponseEnvelopeMessages]
type storeDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StoreDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    storeDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// storeDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [StoreDeleteResponseEnvelopeMessagesSource]
type storeDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StoreDeleteResponseEnvelopeSuccess bool

const (
	StoreDeleteResponseEnvelopeSuccessTrue StoreDeleteResponseEnvelopeSuccess = true
)

func (r StoreDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StoreDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StoreDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                   `json:"total_count"`
	JSON       storeDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// storeDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [StoreDeleteResponseEnvelopeResultInfo]
type storeDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
