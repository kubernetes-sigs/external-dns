// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// RequestAssetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestAssetService] method instead.
type RequestAssetService struct {
	Options []option.RequestOption
}

// NewRequestAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestAssetService(opts ...option.RequestOption) (r *RequestAssetService) {
	r = &RequestAssetService{}
	r.Options = opts
	return
}

// List Request Assets
func (r *RequestAssetService) New(ctx context.Context, requestID string, params RequestAssetNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[RequestAssetNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset", params.AccountID, requestID)
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

// List Request Assets
func (r *RequestAssetService) NewAutoPaging(ctx context.Context, requestID string, params RequestAssetNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RequestAssetNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, requestID, params, opts...))
}

// Update a Request Asset
func (r *RequestAssetService) Update(ctx context.Context, requestID string, assetID string, params RequestAssetUpdateParams, opts ...option.RequestOption) (res *RequestAssetUpdateResponse, err error) {
	var env RequestAssetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", params.AccountID, requestID, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Request Asset
func (r *RequestAssetService) Delete(ctx context.Context, requestID string, assetID string, body RequestAssetDeleteParams, opts ...option.RequestOption) (res *RequestAssetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", body.AccountID, requestID, assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a Request Asset
func (r *RequestAssetService) Get(ctx context.Context, requestID string, assetID string, query RequestAssetGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[RequestAssetGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", query.AccountID, requestID, assetID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get a Request Asset
func (r *RequestAssetService) GetAutoPaging(ctx context.Context, requestID string, assetID string, query RequestAssetGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RequestAssetGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, requestID, assetID, query, opts...))
}

type RequestAssetNewResponse struct {
	// Asset ID.
	ID int64 `json:"id,required"`
	// Asset name.
	Name string `json:"name,required"`
	// Defines the asset creation time.
	Created time.Time `json:"created" format:"date-time"`
	// Asset description.
	Description string `json:"description"`
	// Asset file type.
	FileType string                      `json:"file_type"`
	JSON     requestAssetNewResponseJSON `json:"-"`
}

// requestAssetNewResponseJSON contains the JSON metadata for the struct
// [RequestAssetNewResponse]
type requestAssetNewResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetNewResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponse struct {
	// Asset ID.
	ID int64 `json:"id,required"`
	// Asset name.
	Name string `json:"name,required"`
	// Defines the asset creation time.
	Created time.Time `json:"created" format:"date-time"`
	// Asset description.
	Description string `json:"description"`
	// Asset file type.
	FileType string                         `json:"file_type"`
	JSON     requestAssetUpdateResponseJSON `json:"-"`
}

// requestAssetUpdateResponseJSON contains the JSON metadata for the struct
// [RequestAssetUpdateResponse]
type requestAssetUpdateResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponse struct {
	Errors   []RequestAssetDeleteResponseError   `json:"errors,required"`
	Messages []RequestAssetDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestAssetDeleteResponseSuccess `json:"success,required"`
	JSON    requestAssetDeleteResponseJSON    `json:"-"`
}

// requestAssetDeleteResponseJSON contains the JSON metadata for the struct
// [RequestAssetDeleteResponse]
type requestAssetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           RequestAssetDeleteResponseErrorsSource `json:"source"`
	JSON             requestAssetDeleteResponseErrorJSON    `json:"-"`
}

// requestAssetDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RequestAssetDeleteResponseError]
type requestAssetDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestAssetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    requestAssetDeleteResponseErrorsSourceJSON `json:"-"`
}

// requestAssetDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [RequestAssetDeleteResponseErrorsSource]
type requestAssetDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponseMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           RequestAssetDeleteResponseMessagesSource `json:"source"`
	JSON             requestAssetDeleteResponseMessageJSON    `json:"-"`
}

// requestAssetDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RequestAssetDeleteResponseMessage]
type requestAssetDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestAssetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RequestAssetDeleteResponseMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    requestAssetDeleteResponseMessagesSourceJSON `json:"-"`
}

// requestAssetDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [RequestAssetDeleteResponseMessagesSource]
type requestAssetDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestAssetDeleteResponseSuccess bool

const (
	RequestAssetDeleteResponseSuccessTrue RequestAssetDeleteResponseSuccess = true
)

func (r RequestAssetDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RequestAssetDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RequestAssetGetResponse struct {
	// Asset ID.
	ID int64 `json:"id,required"`
	// Asset name.
	Name string `json:"name,required"`
	// Defines the asset creation time.
	Created time.Time `json:"created" format:"date-time"`
	// Asset description.
	Description string `json:"description"`
	// Asset file type.
	FileType string                      `json:"file_type"`
	JSON     requestAssetGetResponseJSON `json:"-"`
}

// requestAssetGetResponseJSON contains the JSON metadata for the struct
// [RequestAssetGetResponse]
type requestAssetGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetGetResponseJSON) RawJSON() string {
	return r.raw
}

type RequestAssetNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of results.
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page.
	PerPage param.Field[int64] `json:"per_page,required"`
}

func (r RequestAssetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestAssetUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Asset file to upload.
	Source param.Field[string] `json:"source"`
}

func (r RequestAssetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestAssetUpdateResponseEnvelope struct {
	Errors   []RequestAssetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RequestAssetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RequestAssetUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RequestAssetUpdateResponse                `json:"result"`
	JSON    requestAssetUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestAssetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestAssetUpdateResponseEnvelope]
type requestAssetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RequestAssetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             requestAssetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// requestAssetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RequestAssetUpdateResponseEnvelopeErrors]
type requestAssetUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    requestAssetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// requestAssetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RequestAssetUpdateResponseEnvelopeErrorsSource]
type requestAssetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RequestAssetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             requestAssetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// requestAssetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RequestAssetUpdateResponseEnvelopeMessages]
type requestAssetUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RequestAssetUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    requestAssetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// requestAssetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RequestAssetUpdateResponseEnvelopeMessagesSource]
type requestAssetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestAssetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestAssetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RequestAssetUpdateResponseEnvelopeSuccess bool

const (
	RequestAssetUpdateResponseEnvelopeSuccessTrue RequestAssetUpdateResponseEnvelopeSuccess = true
)

func (r RequestAssetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestAssetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestAssetDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RequestAssetGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
