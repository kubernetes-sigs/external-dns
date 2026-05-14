// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

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

// TURNService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTURNService] method instead.
type TURNService struct {
	Options []option.RequestOption
}

// NewTURNService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTURNService(opts ...option.RequestOption) (r *TURNService) {
	r = &TURNService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare Calls TURN key.
func (r *TURNService) New(ctx context.Context, params TURNNewParams, opts ...option.RequestOption) (res *TURNNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Edit details for a single TURN key.
func (r *TURNService) Update(ctx context.Context, keyID string, params TURNUpdateParams, opts ...option.RequestOption) (res *TURNUpdateResponse, err error) {
	var env TURNUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", params.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all TURN keys in the Cloudflare account
func (r *TURNService) List(ctx context.Context, query TURNListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TURNListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", query.AccountID)
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

// Lists all TURN keys in the Cloudflare account
func (r *TURNService) ListAutoPaging(ctx context.Context, query TURNListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TURNListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a TURN key from Cloudflare Calls
func (r *TURNService) Delete(ctx context.Context, keyID string, body TURNDeleteParams, opts ...option.RequestOption) (res *TURNDeleteResponse, err error) {
	var env TURNDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", body.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single TURN key.
func (r *TURNService) Get(ctx context.Context, keyID string, query TURNGetParams, opts ...option.RequestOption) (res *TURNGetResponse, err error) {
	var env TURNGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", query.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TURNNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// Bearer token
	Key string `json:"key"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of a TURN key, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string              `json:"uid"`
	JSON turnNewResponseJSON `json:"-"`
}

// turnNewResponseJSON contains the JSON metadata for the struct [TURNNewResponse]
type turnNewResponseJSON struct {
	Created     apijson.Field
	Key         apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnNewResponseJSON) RawJSON() string {
	return r.raw
}

type TURNUpdateResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                 `json:"uid"`
	JSON turnUpdateResponseJSON `json:"-"`
}

// turnUpdateResponseJSON contains the JSON metadata for the struct
// [TURNUpdateResponse]
type turnUpdateResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type TURNListResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string               `json:"uid"`
	JSON turnListResponseJSON `json:"-"`
}

// turnListResponseJSON contains the JSON metadata for the struct
// [TURNListResponse]
type turnListResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnListResponseJSON) RawJSON() string {
	return r.raw
}

type TURNDeleteResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                 `json:"uid"`
	JSON turnDeleteResponseJSON `json:"-"`
}

// turnDeleteResponseJSON contains the JSON metadata for the struct
// [TURNDeleteResponse]
type turnDeleteResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TURNGetResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string              `json:"uid"`
	JSON turnGetResponseJSON `json:"-"`
}

// turnGetResponseJSON contains the JSON metadata for the struct [TURNGetResponse]
type turnGetResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseJSON) RawJSON() string {
	return r.raw
}

type TURNNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TURNNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TURNUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TURNUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TURNUpdateResponseEnvelope struct {
	Errors   []TURNUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TURNUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TURNUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  TURNUpdateResponse                `json:"result"`
	JSON    turnUpdateResponseEnvelopeJSON    `json:"-"`
}

// turnUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNUpdateResponseEnvelope]
type turnUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TURNUpdateResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           TURNUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             turnUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// turnUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TURNUpdateResponseEnvelopeErrors]
type turnUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TURNUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    turnUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// turnUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TURNUpdateResponseEnvelopeErrorsSource]
type turnUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TURNUpdateResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           TURNUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             turnUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// turnUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TURNUpdateResponseEnvelopeMessages]
type turnUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TURNUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    turnUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// turnUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TURNUpdateResponseEnvelopeMessagesSource]
type turnUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TURNUpdateResponseEnvelopeSuccess bool

const (
	TURNUpdateResponseEnvelopeSuccessTrue TURNUpdateResponseEnvelopeSuccess = true
)

func (r TURNUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TURNListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TURNDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TURNDeleteResponseEnvelope struct {
	Errors   []TURNDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TURNDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TURNDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  TURNDeleteResponse                `json:"result"`
	JSON    turnDeleteResponseEnvelopeJSON    `json:"-"`
}

// turnDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNDeleteResponseEnvelope]
type turnDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TURNDeleteResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           TURNDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             turnDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// turnDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TURNDeleteResponseEnvelopeErrors]
type turnDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TURNDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    turnDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// turnDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TURNDeleteResponseEnvelopeErrorsSource]
type turnDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TURNDeleteResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           TURNDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             turnDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// turnDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TURNDeleteResponseEnvelopeMessages]
type turnDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TURNDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    turnDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// turnDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TURNDeleteResponseEnvelopeMessagesSource]
type turnDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TURNDeleteResponseEnvelopeSuccess bool

const (
	TURNDeleteResponseEnvelopeSuccessTrue TURNDeleteResponseEnvelopeSuccess = true
)

func (r TURNDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TURNGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TURNGetResponseEnvelope struct {
	Errors   []TURNGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TURNGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TURNGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TURNGetResponse                `json:"result"`
	JSON    turnGetResponseEnvelopeJSON    `json:"-"`
}

// turnGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNGetResponseEnvelope]
type turnGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TURNGetResponseEnvelopeErrors struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           TURNGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             turnGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// turnGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TURNGetResponseEnvelopeErrors]
type turnGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TURNGetResponseEnvelopeErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    turnGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// turnGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TURNGetResponseEnvelopeErrorsSource]
type turnGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TURNGetResponseEnvelopeMessages struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           TURNGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             turnGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// turnGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TURNGetResponseEnvelopeMessages]
type turnGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TURNGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TURNGetResponseEnvelopeMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    turnGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// turnGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TURNGetResponseEnvelopeMessagesSource]
type turnGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TURNGetResponseEnvelopeSuccess bool

const (
	TURNGetResponseEnvelopeSuccessTrue TURNGetResponseEnvelopeSuccess = true
)

func (r TURNGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
