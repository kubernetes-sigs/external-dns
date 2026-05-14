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

// StoreSecretService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreSecretService] method instead.
type StoreSecretService struct {
	Options []option.RequestOption
}

// NewStoreSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreSecretService(opts ...option.RequestOption) (r *StoreSecretService) {
	r = &StoreSecretService{}
	r.Options = opts
	return
}

// Creates a secret in the account
func (r *StoreSecretService) New(ctx context.Context, storeID string, params StoreSecretNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[StoreSecretNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets", params.AccountID, storeID)
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

// Creates a secret in the account
func (r *StoreSecretService) NewAutoPaging(ctx context.Context, storeID string, params StoreSecretNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[StoreSecretNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, storeID, params, opts...))
}

// Lists all store secrets
func (r *StoreSecretService) List(ctx context.Context, storeID string, params StoreSecretListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[StoreSecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets", params.AccountID, storeID)
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

// Lists all store secrets
func (r *StoreSecretService) ListAutoPaging(ctx context.Context, storeID string, params StoreSecretListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[StoreSecretListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, storeID, params, opts...))
}

// Deletes a single secret
func (r *StoreSecretService) Delete(ctx context.Context, storeID string, secretID string, body StoreSecretDeleteParams, opts ...option.RequestOption) (res *StoreSecretDeleteResponse, err error) {
	var env StoreSecretDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets/%s", body.AccountID, storeID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes one or more secrets
func (r *StoreSecretService) BulkDelete(ctx context.Context, storeID string, body StoreSecretBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[StoreSecretBulkDeleteResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets", body.AccountID, storeID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes one or more secrets
func (r *StoreSecretService) BulkDeleteAutoPaging(ctx context.Context, storeID string, body StoreSecretBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[StoreSecretBulkDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, storeID, body, opts...))
}

// Duplicates the secret, keeping the value
func (r *StoreSecretService) Duplicate(ctx context.Context, storeID string, secretID string, params StoreSecretDuplicateParams, opts ...option.RequestOption) (res *StoreSecretDuplicateResponse, err error) {
	var env StoreSecretDuplicateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets/%s/duplicate", params.AccountID, storeID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a single secret
func (r *StoreSecretService) Edit(ctx context.Context, storeID string, secretID string, params StoreSecretEditParams, opts ...option.RequestOption) (res *StoreSecretEditResponse, err error) {
	var env StoreSecretEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets/%s", params.AccountID, storeID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns details of a single secret
func (r *StoreSecretService) Get(ctx context.Context, storeID string, secretID string, query StoreSecretGetParams, opts ...option.RequestOption) (res *StoreSecretGetResponse, err error) {
	var env StoreSecretGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if storeID == "" {
		err = errors.New("missing required store_id parameter")
		return
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/stores/%s/secrets/%s", query.AccountID, storeID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StoreSecretNewResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                       `json:"name,required"`
	Status StoreSecretNewResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                     `json:"comment"`
	JSON    storeSecretNewResponseJSON `json:"-"`
}

// storeSecretNewResponseJSON contains the JSON metadata for the struct
// [StoreSecretNewResponse]
type storeSecretNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretNewResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretNewResponseStatus string

const (
	StoreSecretNewResponseStatusPending StoreSecretNewResponseStatus = "pending"
	StoreSecretNewResponseStatusActive  StoreSecretNewResponseStatus = "active"
	StoreSecretNewResponseStatusDeleted StoreSecretNewResponseStatus = "deleted"
)

func (r StoreSecretNewResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretNewResponseStatusPending, StoreSecretNewResponseStatusActive, StoreSecretNewResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretListResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                        `json:"name,required"`
	Status StoreSecretListResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                      `json:"comment"`
	JSON    storeSecretListResponseJSON `json:"-"`
}

// storeSecretListResponseJSON contains the JSON metadata for the struct
// [StoreSecretListResponse]
type storeSecretListResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretListResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretListResponseStatus string

const (
	StoreSecretListResponseStatusPending StoreSecretListResponseStatus = "pending"
	StoreSecretListResponseStatusActive  StoreSecretListResponseStatus = "active"
	StoreSecretListResponseStatusDeleted StoreSecretListResponseStatus = "deleted"
)

func (r StoreSecretListResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretListResponseStatusPending, StoreSecretListResponseStatusActive, StoreSecretListResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretDeleteResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                          `json:"name,required"`
	Status StoreSecretDeleteResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                        `json:"comment"`
	JSON    storeSecretDeleteResponseJSON `json:"-"`
}

// storeSecretDeleteResponseJSON contains the JSON metadata for the struct
// [StoreSecretDeleteResponse]
type storeSecretDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDeleteResponseStatus string

const (
	StoreSecretDeleteResponseStatusPending StoreSecretDeleteResponseStatus = "pending"
	StoreSecretDeleteResponseStatusActive  StoreSecretDeleteResponseStatus = "active"
	StoreSecretDeleteResponseStatusDeleted StoreSecretDeleteResponseStatus = "deleted"
)

func (r StoreSecretDeleteResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretDeleteResponseStatusPending, StoreSecretDeleteResponseStatusActive, StoreSecretDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretBulkDeleteResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                              `json:"name,required"`
	Status StoreSecretBulkDeleteResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                            `json:"comment"`
	JSON    storeSecretBulkDeleteResponseJSON `json:"-"`
}

// storeSecretBulkDeleteResponseJSON contains the JSON metadata for the struct
// [StoreSecretBulkDeleteResponse]
type storeSecretBulkDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretBulkDeleteResponseStatus string

const (
	StoreSecretBulkDeleteResponseStatusPending StoreSecretBulkDeleteResponseStatus = "pending"
	StoreSecretBulkDeleteResponseStatusActive  StoreSecretBulkDeleteResponseStatus = "active"
	StoreSecretBulkDeleteResponseStatusDeleted StoreSecretBulkDeleteResponseStatus = "deleted"
)

func (r StoreSecretBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretBulkDeleteResponseStatusPending, StoreSecretBulkDeleteResponseStatusActive, StoreSecretBulkDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretDuplicateResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                             `json:"name,required"`
	Status StoreSecretDuplicateResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                           `json:"comment"`
	JSON    storeSecretDuplicateResponseJSON `json:"-"`
}

// storeSecretDuplicateResponseJSON contains the JSON metadata for the struct
// [StoreSecretDuplicateResponse]
type storeSecretDuplicateResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDuplicateResponseStatus string

const (
	StoreSecretDuplicateResponseStatusPending StoreSecretDuplicateResponseStatus = "pending"
	StoreSecretDuplicateResponseStatusActive  StoreSecretDuplicateResponseStatus = "active"
	StoreSecretDuplicateResponseStatusDeleted StoreSecretDuplicateResponseStatus = "deleted"
)

func (r StoreSecretDuplicateResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretDuplicateResponseStatusPending, StoreSecretDuplicateResponseStatusActive, StoreSecretDuplicateResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretEditResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                        `json:"name,required"`
	Status StoreSecretEditResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                      `json:"comment"`
	JSON    storeSecretEditResponseJSON `json:"-"`
}

// storeSecretEditResponseJSON contains the JSON metadata for the struct
// [StoreSecretEditResponse]
type storeSecretEditResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditResponseStatus string

const (
	StoreSecretEditResponseStatusPending StoreSecretEditResponseStatus = "pending"
	StoreSecretEditResponseStatusActive  StoreSecretEditResponseStatus = "active"
	StoreSecretEditResponseStatusDeleted StoreSecretEditResponseStatus = "deleted"
)

func (r StoreSecretEditResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretEditResponseStatusPending, StoreSecretEditResponseStatusActive, StoreSecretEditResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretGetResponse struct {
	// Secret identifier tag.
	ID string `json:"id,required"`
	// Whenthe secret was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the secret was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the secret
	Name   string                       `json:"name,required"`
	Status StoreSecretGetResponseStatus `json:"status,required"`
	// Store Identifier
	StoreID string `json:"store_id,required"`
	// Freeform text describing the secret
	Comment string                     `json:"comment"`
	JSON    storeSecretGetResponseJSON `json:"-"`
}

// storeSecretGetResponseJSON contains the JSON metadata for the struct
// [StoreSecretGetResponse]
type storeSecretGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	StoreID     apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetResponseStatus string

const (
	StoreSecretGetResponseStatusPending StoreSecretGetResponseStatus = "pending"
	StoreSecretGetResponseStatusActive  StoreSecretGetResponseStatus = "active"
	StoreSecretGetResponseStatusDeleted StoreSecretGetResponseStatus = "deleted"
)

func (r StoreSecretGetResponseStatus) IsKnown() bool {
	switch r {
	case StoreSecretGetResponseStatusPending, StoreSecretGetResponseStatusActive, StoreSecretGetResponseStatusDeleted:
		return true
	}
	return false
}

type StoreSecretNewParams struct {
	// Account Identifier
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      []StoreSecretNewParamsBody `json:"body,required"`
}

func (r StoreSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StoreSecretNewParamsBody struct {
	// The name of the secret
	Name param.Field[string] `json:"name,required"`
	// The list of services that can use this secret.
	Scopes param.Field[[]string] `json:"scopes,required"`
	// The value of the secret. Note that this is 'write only' - no API reponse will
	// provide this value, it is only used to create/modify secrets.
	Value param.Field[string] `json:"value,required"`
	// Freeform text describing the secret
	Comment param.Field[string] `json:"comment"`
}

func (r StoreSecretNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StoreSecretListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to sort objects
	Direction param.Field[StoreSecretListParamsDirection] `query:"direction"`
	// Order secrets by values in the given field
	Order param.Field[StoreSecretListParamsOrder] `query:"order"`
	// Page number
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page
	PerPage param.Field[int64] `query:"per_page"`
	// Only secrets with the given scopes will be returned
	Scopes param.Field[[][]string] `query:"scopes"`
	// Search secrets using a filter string, filtering across name and comment
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [StoreSecretListParams]'s query parameters as `url.Values`.
func (r StoreSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to sort objects
type StoreSecretListParamsDirection string

const (
	StoreSecretListParamsDirectionAsc  StoreSecretListParamsDirection = "asc"
	StoreSecretListParamsDirectionDesc StoreSecretListParamsDirection = "desc"
)

func (r StoreSecretListParamsDirection) IsKnown() bool {
	switch r {
	case StoreSecretListParamsDirectionAsc, StoreSecretListParamsDirectionDesc:
		return true
	}
	return false
}

// Order secrets by values in the given field
type StoreSecretListParamsOrder string

const (
	StoreSecretListParamsOrderName     StoreSecretListParamsOrder = "name"
	StoreSecretListParamsOrderComment  StoreSecretListParamsOrder = "comment"
	StoreSecretListParamsOrderCreated  StoreSecretListParamsOrder = "created"
	StoreSecretListParamsOrderModified StoreSecretListParamsOrder = "modified"
	StoreSecretListParamsOrderStatus   StoreSecretListParamsOrder = "status"
)

func (r StoreSecretListParamsOrder) IsKnown() bool {
	switch r {
	case StoreSecretListParamsOrderName, StoreSecretListParamsOrderComment, StoreSecretListParamsOrderCreated, StoreSecretListParamsOrderModified, StoreSecretListParamsOrderStatus:
		return true
	}
	return false
}

type StoreSecretDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StoreSecretDeleteResponseEnvelope struct {
	Errors   []StoreSecretDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StoreSecretDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    StoreSecretDeleteResponseEnvelopeSuccess    `json:"success,required"`
	Result     StoreSecretDeleteResponse                   `json:"result"`
	ResultInfo StoreSecretDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storeSecretDeleteResponseEnvelopeJSON       `json:"-"`
}

// storeSecretDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [StoreSecretDeleteResponseEnvelope]
type storeSecretDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           StoreSecretDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             storeSecretDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// storeSecretDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StoreSecretDeleteResponseEnvelopeErrors]
type storeSecretDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    storeSecretDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// storeSecretDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [StoreSecretDeleteResponseEnvelopeErrorsSource]
type storeSecretDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           StoreSecretDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             storeSecretDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// storeSecretDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StoreSecretDeleteResponseEnvelopeMessages]
type storeSecretDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    storeSecretDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// storeSecretDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [StoreSecretDeleteResponseEnvelopeMessagesSource]
type storeSecretDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StoreSecretDeleteResponseEnvelopeSuccess bool

const (
	StoreSecretDeleteResponseEnvelopeSuccessTrue StoreSecretDeleteResponseEnvelopeSuccess = true
)

func (r StoreSecretDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StoreSecretDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StoreSecretDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                         `json:"total_count"`
	JSON       storeSecretDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// storeSecretDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [StoreSecretDeleteResponseEnvelopeResultInfo]
type storeSecretDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type StoreSecretBulkDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StoreSecretDuplicateParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the secret
	Name param.Field[string] `json:"name,required"`
	// The list of services that can use this secret.
	Scopes param.Field[[]string] `json:"scopes,required"`
	// Freeform text describing the secret
	Comment param.Field[string] `json:"comment"`
}

func (r StoreSecretDuplicateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StoreSecretDuplicateResponseEnvelope struct {
	Errors   []StoreSecretDuplicateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StoreSecretDuplicateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    StoreSecretDuplicateResponseEnvelopeSuccess    `json:"success,required"`
	Result     StoreSecretDuplicateResponse                   `json:"result"`
	ResultInfo StoreSecretDuplicateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storeSecretDuplicateResponseEnvelopeJSON       `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StoreSecretDuplicateResponseEnvelope]
type storeSecretDuplicateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDuplicateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           StoreSecretDuplicateResponseEnvelopeErrorsSource `json:"source"`
	JSON             storeSecretDuplicateResponseEnvelopeErrorsJSON   `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StoreSecretDuplicateResponseEnvelopeErrors]
type storeSecretDuplicateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDuplicateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    storeSecretDuplicateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [StoreSecretDuplicateResponseEnvelopeErrorsSource]
type storeSecretDuplicateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDuplicateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           StoreSecretDuplicateResponseEnvelopeMessagesSource `json:"source"`
	JSON             storeSecretDuplicateResponseEnvelopeMessagesJSON   `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StoreSecretDuplicateResponseEnvelopeMessages]
type storeSecretDuplicateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StoreSecretDuplicateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    storeSecretDuplicateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [StoreSecretDuplicateResponseEnvelopeMessagesSource]
type storeSecretDuplicateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StoreSecretDuplicateResponseEnvelopeSuccess bool

const (
	StoreSecretDuplicateResponseEnvelopeSuccessTrue StoreSecretDuplicateResponseEnvelopeSuccess = true
)

func (r StoreSecretDuplicateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StoreSecretDuplicateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StoreSecretDuplicateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                            `json:"total_count"`
	JSON       storeSecretDuplicateResponseEnvelopeResultInfoJSON `json:"-"`
}

// storeSecretDuplicateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [StoreSecretDuplicateResponseEnvelopeResultInfo]
type storeSecretDuplicateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretDuplicateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretDuplicateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Freeform text describing the secret
	Comment param.Field[string] `json:"comment"`
	// The list of services that can use this secret.
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r StoreSecretEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StoreSecretEditResponseEnvelope struct {
	Errors   []StoreSecretEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StoreSecretEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    StoreSecretEditResponseEnvelopeSuccess    `json:"success,required"`
	Result     StoreSecretEditResponse                   `json:"result"`
	ResultInfo StoreSecretEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storeSecretEditResponseEnvelopeJSON       `json:"-"`
}

// storeSecretEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [StoreSecretEditResponseEnvelope]
type storeSecretEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           StoreSecretEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             storeSecretEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// storeSecretEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StoreSecretEditResponseEnvelopeErrors]
type storeSecretEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    storeSecretEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// storeSecretEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [StoreSecretEditResponseEnvelopeErrorsSource]
type storeSecretEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           StoreSecretEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             storeSecretEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// storeSecretEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StoreSecretEditResponseEnvelopeMessages]
type storeSecretEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StoreSecretEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    storeSecretEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// storeSecretEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [StoreSecretEditResponseEnvelopeMessagesSource]
type storeSecretEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StoreSecretEditResponseEnvelopeSuccess bool

const (
	StoreSecretEditResponseEnvelopeSuccessTrue StoreSecretEditResponseEnvelopeSuccess = true
)

func (r StoreSecretEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StoreSecretEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StoreSecretEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                       `json:"total_count"`
	JSON       storeSecretEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// storeSecretEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [StoreSecretEditResponseEnvelopeResultInfo]
type storeSecretEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StoreSecretGetResponseEnvelope struct {
	Errors   []StoreSecretGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StoreSecretGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    StoreSecretGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     StoreSecretGetResponse                   `json:"result"`
	ResultInfo StoreSecretGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       storeSecretGetResponseEnvelopeJSON       `json:"-"`
}

// storeSecretGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StoreSecretGetResponseEnvelope]
type storeSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           StoreSecretGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             storeSecretGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// storeSecretGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StoreSecretGetResponseEnvelopeErrors]
type storeSecretGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    storeSecretGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// storeSecretGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [StoreSecretGetResponseEnvelopeErrorsSource]
type storeSecretGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           StoreSecretGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             storeSecretGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// storeSecretGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StoreSecretGetResponseEnvelopeMessages]
type storeSecretGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StoreSecretGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    storeSecretGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// storeSecretGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [StoreSecretGetResponseEnvelopeMessagesSource]
type storeSecretGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StoreSecretGetResponseEnvelopeSuccess bool

const (
	StoreSecretGetResponseEnvelopeSuccessTrue StoreSecretGetResponseEnvelopeSuccess = true
)

func (r StoreSecretGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StoreSecretGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StoreSecretGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                      `json:"total_count"`
	JSON       storeSecretGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// storeSecretGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [StoreSecretGetResponseEnvelopeResultInfo]
type storeSecretGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreSecretGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeSecretGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
