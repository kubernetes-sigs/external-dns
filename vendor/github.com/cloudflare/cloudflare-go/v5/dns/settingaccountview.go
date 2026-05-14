// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// SettingAccountViewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAccountViewService] method instead.
type SettingAccountViewService struct {
	Options []option.RequestOption
}

// NewSettingAccountViewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAccountViewService(opts ...option.RequestOption) (r *SettingAccountViewService) {
	r = &SettingAccountViewService{}
	r.Options = opts
	return
}

// Create Internal DNS View for an account
func (r *SettingAccountViewService) New(ctx context.Context, params SettingAccountViewNewParams, opts ...option.RequestOption) (res *SettingAccountViewNewResponse, err error) {
	var env SettingAccountViewNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List DNS Internal Views for an Account
func (r *SettingAccountViewService) List(ctx context.Context, params SettingAccountViewListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingAccountViewListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", params.AccountID)
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

// List DNS Internal Views for an Account
func (r *SettingAccountViewService) ListAutoPaging(ctx context.Context, params SettingAccountViewListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingAccountViewListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an existing Internal DNS View
func (r *SettingAccountViewService) Delete(ctx context.Context, viewID string, body SettingAccountViewDeleteParams, opts ...option.RequestOption) (res *SettingAccountViewDeleteResponse, err error) {
	var env SettingAccountViewDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", body.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing Internal DNS View
func (r *SettingAccountViewService) Edit(ctx context.Context, viewID string, params SettingAccountViewEditParams, opts ...option.RequestOption) (res *SettingAccountViewEditResponse, err error) {
	var env SettingAccountViewEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", params.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get DNS Internal View
func (r *SettingAccountViewService) Get(ctx context.Context, viewID string, query SettingAccountViewGetParams, opts ...option.RequestOption) (res *SettingAccountViewGetResponse, err error) {
	var env SettingAccountViewGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", query.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingAccountViewNewResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                          `json:"zones,required"`
	JSON  settingAccountViewNewResponseJSON `json:"-"`
}

// settingAccountViewNewResponseJSON contains the JSON metadata for the struct
// [SettingAccountViewNewResponse]
type settingAccountViewNewResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountViewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewListResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                           `json:"zones,required"`
	JSON  settingAccountViewListResponseJSON `json:"-"`
}

// settingAccountViewListResponseJSON contains the JSON metadata for the struct
// [SettingAccountViewListResponse]
type settingAccountViewListResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountViewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewDeleteResponse struct {
	// Identifier.
	ID   string                               `json:"id"`
	JSON settingAccountViewDeleteResponseJSON `json:"-"`
}

// settingAccountViewDeleteResponseJSON contains the JSON metadata for the struct
// [SettingAccountViewDeleteResponse]
type settingAccountViewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                           `json:"zones,required"`
	JSON  settingAccountViewEditResponseJSON `json:"-"`
}

// settingAccountViewEditResponseJSON contains the JSON metadata for the struct
// [SettingAccountViewEditResponse]
type settingAccountViewEditResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountViewEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewGetResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                          `json:"zones,required"`
	JSON  settingAccountViewGetResponseJSON `json:"-"`
}

// settingAccountViewGetResponseJSON contains the JSON metadata for the struct
// [SettingAccountViewGetResponse]
type settingAccountViewGetResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountViewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name,required"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones,required"`
}

func (r SettingAccountViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAccountViewNewResponseEnvelope struct {
	Errors   []SettingAccountViewNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAccountViewNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingAccountViewNewResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingAccountViewNewResponse                `json:"result"`
	JSON    settingAccountViewNewResponseEnvelopeJSON    `json:"-"`
}

// settingAccountViewNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAccountViewNewResponseEnvelope]
type settingAccountViewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingAccountViewNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountViewNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountViewNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAccountViewNewResponseEnvelopeErrors]
type settingAccountViewNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingAccountViewNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountViewNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingAccountViewNewResponseEnvelopeErrorsSource]
type settingAccountViewNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingAccountViewNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountViewNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountViewNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingAccountViewNewResponseEnvelopeMessages]
type settingAccountViewNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingAccountViewNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountViewNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAccountViewNewResponseEnvelopeMessagesSource]
type settingAccountViewNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountViewNewResponseEnvelopeSuccess bool

const (
	SettingAccountViewNewResponseEnvelopeSuccessTrue SettingAccountViewNewResponseEnvelopeSuccess = true
)

func (r SettingAccountViewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountViewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAccountViewListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order DNS views in.
	Direction param.Field[SettingAccountViewListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead.
	Match param.Field[SettingAccountViewListParamsMatch] `query:"match"`
	Name  param.Field[SettingAccountViewListParamsName]  `query:"name"`
	// Field to order DNS views by.
	Order param.Field[SettingAccountViewListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS views per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A zone ID that exists in the zones list for the view.
	ZoneID param.Field[string] `query:"zone_id"`
	// A zone name that exists in the zones list for the view.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [SettingAccountViewListParams]'s query parameters as
// `url.Values`.
func (r SettingAccountViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order DNS views in.
type SettingAccountViewListParamsDirection string

const (
	SettingAccountViewListParamsDirectionAsc  SettingAccountViewListParamsDirection = "asc"
	SettingAccountViewListParamsDirectionDesc SettingAccountViewListParamsDirection = "desc"
)

func (r SettingAccountViewListParamsDirection) IsKnown() bool {
	switch r {
	case SettingAccountViewListParamsDirectionAsc, SettingAccountViewListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead.
type SettingAccountViewListParamsMatch string

const (
	SettingAccountViewListParamsMatchAny SettingAccountViewListParamsMatch = "any"
	SettingAccountViewListParamsMatchAll SettingAccountViewListParamsMatch = "all"
)

func (r SettingAccountViewListParamsMatch) IsKnown() bool {
	switch r {
	case SettingAccountViewListParamsMatchAny, SettingAccountViewListParamsMatchAll:
		return true
	}
	return false
}

type SettingAccountViewListParamsName struct {
	// Substring of the DNS view name.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS view name.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS view name.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS view name.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [SettingAccountViewListParamsName]'s query parameters as
// `url.Values`.
func (r SettingAccountViewListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order DNS views by.
type SettingAccountViewListParamsOrder string

const (
	SettingAccountViewListParamsOrderName       SettingAccountViewListParamsOrder = "name"
	SettingAccountViewListParamsOrderCreatedOn  SettingAccountViewListParamsOrder = "created_on"
	SettingAccountViewListParamsOrderModifiedOn SettingAccountViewListParamsOrder = "modified_on"
)

func (r SettingAccountViewListParamsOrder) IsKnown() bool {
	switch r {
	case SettingAccountViewListParamsOrderName, SettingAccountViewListParamsOrderCreatedOn, SettingAccountViewListParamsOrderModifiedOn:
		return true
	}
	return false
}

type SettingAccountViewDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAccountViewDeleteResponseEnvelope struct {
	Result SettingAccountViewDeleteResponse             `json:"result"`
	JSON   settingAccountViewDeleteResponseEnvelopeJSON `json:"-"`
}

// settingAccountViewDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAccountViewDeleteResponseEnvelope]
type settingAccountViewDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones"`
}

func (r SettingAccountViewEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAccountViewEditResponseEnvelope struct {
	Errors   []SettingAccountViewEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAccountViewEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingAccountViewEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingAccountViewEditResponse                `json:"result"`
	JSON    settingAccountViewEditResponseEnvelopeJSON    `json:"-"`
}

// settingAccountViewEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAccountViewEditResponseEnvelope]
type settingAccountViewEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           SettingAccountViewEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountViewEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountViewEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAccountViewEditResponseEnvelopeErrors]
type settingAccountViewEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    settingAccountViewEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountViewEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingAccountViewEditResponseEnvelopeErrorsSource]
type settingAccountViewEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingAccountViewEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountViewEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountViewEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAccountViewEditResponseEnvelopeMessages]
type settingAccountViewEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingAccountViewEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountViewEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAccountViewEditResponseEnvelopeMessagesSource]
type settingAccountViewEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountViewEditResponseEnvelopeSuccess bool

const (
	SettingAccountViewEditResponseEnvelopeSuccessTrue SettingAccountViewEditResponseEnvelopeSuccess = true
)

func (r SettingAccountViewEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountViewEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAccountViewGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAccountViewGetResponseEnvelope struct {
	Errors   []SettingAccountViewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAccountViewGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingAccountViewGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingAccountViewGetResponse                `json:"result"`
	JSON    settingAccountViewGetResponseEnvelopeJSON    `json:"-"`
}

// settingAccountViewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAccountViewGetResponseEnvelope]
type settingAccountViewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingAccountViewGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountViewGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountViewGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAccountViewGetResponseEnvelopeErrors]
type settingAccountViewGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingAccountViewGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountViewGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingAccountViewGetResponseEnvelopeErrorsSource]
type settingAccountViewGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingAccountViewGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountViewGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountViewGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingAccountViewGetResponseEnvelopeMessages]
type settingAccountViewGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountViewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountViewGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingAccountViewGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountViewGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAccountViewGetResponseEnvelopeMessagesSource]
type settingAccountViewGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountViewGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountViewGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountViewGetResponseEnvelopeSuccess bool

const (
	SettingAccountViewGetResponseEnvelopeSuccessTrue SettingAccountViewGetResponseEnvelopeSuccess = true
)

func (r SettingAccountViewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountViewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
