// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// GatewayListService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayListService] method instead.
type GatewayListService struct {
	Options []option.RequestOption
	Items   *GatewayListItemService
}

// NewGatewayListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayListService(opts ...option.RequestOption) (r *GatewayListService) {
	r = &GatewayListService{}
	r.Options = opts
	r.Items = NewGatewayListItemService(opts...)
	return
}

// Creates a new Zero Trust list.
func (r *GatewayListService) New(ctx context.Context, params GatewayListNewParams, opts ...option.RequestOption) (res *GatewayListNewResponse, err error) {
	var env GatewayListNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust list. Skips updating list items if not included
// in the payload.
func (r *GatewayListService) Update(ctx context.Context, listID string, params GatewayListUpdateParams, opts ...option.RequestOption) (res *GatewayList, err error) {
	var env GatewayListUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) List(ctx context.Context, params GatewayListListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayList], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists", params.AccountID)
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

// Fetches all Zero Trust lists for an account.
func (r *GatewayListService) ListAutoPaging(ctx context.Context, params GatewayListListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayList] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Zero Trust list.
func (r *GatewayListService) Delete(ctx context.Context, listID string, body GatewayListDeleteParams, opts ...option.RequestOption) (res *GatewayListDeleteResponse, err error) {
	var env GatewayListDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Appends or removes an item from a configured Zero Trust list.
func (r *GatewayListService) Edit(ctx context.Context, listID string, params GatewayListEditParams, opts ...option.RequestOption) (res *GatewayList, err error) {
	var env GatewayListEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust list.
func (r *GatewayListService) Get(ctx context.Context, listID string, query GatewayListGetParams, opts ...option.RequestOption) (res *GatewayList, err error) {
	var env GatewayListGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", query.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list item, if present
	Description string `json:"description"`
	// The value of the item in a list.
	Value string          `json:"value"`
	JSON  gatewayItemJSON `json:"-"`
}

// gatewayItemJSON contains the JSON metadata for the struct [GatewayItem]
type gatewayItemJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayItemJSON) RawJSON() string {
	return r.raw
}

type GatewayList struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []GatewayItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      GatewayListType `json:"type"`
	UpdatedAt time.Time       `json:"updated_at" format:"date-time"`
	JSON      gatewayListJSON `json:"-"`
}

// gatewayListJSON contains the JSON metadata for the struct [GatewayList]
type gatewayListJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Items       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListType string

const (
	GatewayListTypeSerial GatewayListType = "SERIAL"
	GatewayListTypeURL    GatewayListType = "URL"
	GatewayListTypeDomain GatewayListType = "DOMAIN"
	GatewayListTypeEmail  GatewayListType = "EMAIL"
	GatewayListTypeIP     GatewayListType = "IP"
)

func (r GatewayListType) IsKnown() bool {
	switch r {
	case GatewayListTypeSerial, GatewayListTypeURL, GatewayListTypeDomain, GatewayListTypeEmail, GatewayListTypeIP:
		return true
	}
	return false
}

type GatewayListNewResponse struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []GatewayItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      GatewayListNewResponseType `json:"type"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      gatewayListNewResponseJSON `json:"-"`
}

// gatewayListNewResponseJSON contains the JSON metadata for the struct
// [GatewayListNewResponse]
type gatewayListNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Items       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type GatewayListNewResponseType string

const (
	GatewayListNewResponseTypeSerial GatewayListNewResponseType = "SERIAL"
	GatewayListNewResponseTypeURL    GatewayListNewResponseType = "URL"
	GatewayListNewResponseTypeDomain GatewayListNewResponseType = "DOMAIN"
	GatewayListNewResponseTypeEmail  GatewayListNewResponseType = "EMAIL"
	GatewayListNewResponseTypeIP     GatewayListNewResponseType = "IP"
)

func (r GatewayListNewResponseType) IsKnown() bool {
	switch r {
	case GatewayListNewResponseTypeSerial, GatewayListNewResponseTypeURL, GatewayListNewResponseTypeDomain, GatewayListNewResponseTypeEmail, GatewayListNewResponseTypeIP:
		return true
	}
	return false
}

type GatewayListDeleteResponse = interface{}

type GatewayListNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[GatewayListNewParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// items to add to the list.
	Items param.Field[[]GatewayListNewParamsItem] `json:"items"`
}

func (r GatewayListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of list.
type GatewayListNewParamsType string

const (
	GatewayListNewParamsTypeSerial GatewayListNewParamsType = "SERIAL"
	GatewayListNewParamsTypeURL    GatewayListNewParamsType = "URL"
	GatewayListNewParamsTypeDomain GatewayListNewParamsType = "DOMAIN"
	GatewayListNewParamsTypeEmail  GatewayListNewParamsType = "EMAIL"
	GatewayListNewParamsTypeIP     GatewayListNewParamsType = "IP"
)

func (r GatewayListNewParamsType) IsKnown() bool {
	switch r {
	case GatewayListNewParamsTypeSerial, GatewayListNewParamsTypeURL, GatewayListNewParamsTypeDomain, GatewayListNewParamsTypeEmail, GatewayListNewParamsTypeIP:
		return true
	}
	return false
}

type GatewayListNewParamsItem struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayListNewResponse                `json:"result"`
	JSON    gatewayListNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayListNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListNewResponseEnvelope]
type gatewayListNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListNewResponseEnvelopeSuccess bool

const (
	GatewayListNewResponseEnvelopeSuccessTrue GatewayListNewResponseEnvelopeSuccess = true
)

func (r GatewayListNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// items to add to the list.
	Items param.Field[[]GatewayListUpdateParamsItem] `json:"items"`
}

func (r GatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListUpdateParamsItem struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayList                              `json:"result"`
	JSON    gatewayListUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayListUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListUpdateResponseEnvelope]
type gatewayListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListUpdateResponseEnvelopeSuccess bool

const (
	GatewayListUpdateResponseEnvelopeSuccessTrue GatewayListUpdateResponseEnvelopeSuccess = true
)

func (r GatewayListUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of list.
	Type param.Field[GatewayListListParamsType] `query:"type"`
}

// URLQuery serializes [GatewayListListParams]'s query parameters as `url.Values`.
func (r GatewayListListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of list.
type GatewayListListParamsType string

const (
	GatewayListListParamsTypeSerial GatewayListListParamsType = "SERIAL"
	GatewayListListParamsTypeURL    GatewayListListParamsType = "URL"
	GatewayListListParamsTypeDomain GatewayListListParamsType = "DOMAIN"
	GatewayListListParamsTypeEmail  GatewayListListParamsType = "EMAIL"
	GatewayListListParamsTypeIP     GatewayListListParamsType = "IP"
)

func (r GatewayListListParamsType) IsKnown() bool {
	switch r {
	case GatewayListListParamsTypeSerial, GatewayListListParamsTypeURL, GatewayListListParamsTypeDomain, GatewayListListParamsTypeEmail, GatewayListListParamsTypeIP:
		return true
	}
	return false
}

type GatewayListDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayListDeleteResponse                `json:"result"`
	JSON    gatewayListDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayListDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListDeleteResponseEnvelope]
type gatewayListDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListDeleteResponseEnvelopeSuccess bool

const (
	GatewayListDeleteResponseEnvelopeSuccessTrue GatewayListDeleteResponseEnvelopeSuccess = true
)

func (r GatewayListDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// items to add to the list.
	Append param.Field[[]GatewayListEditParamsAppend] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r GatewayListEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListEditParamsAppend struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r GatewayListEditParamsAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListEditResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayList                            `json:"result"`
	JSON    gatewayListEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayListEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListEditResponseEnvelope]
type gatewayListEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListEditResponseEnvelopeSuccess bool

const (
	GatewayListEditResponseEnvelopeSuccessTrue GatewayListEditResponseEnvelopeSuccess = true
)

func (r GatewayListEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListGetResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayList                           `json:"result"`
	JSON    gatewayListGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayListGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListGetResponseEnvelope]
type gatewayListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListGetResponseEnvelopeSuccess bool

const (
	GatewayListGetResponseEnvelopeSuccessTrue GatewayListGetResponseEnvelopeSuccess = true
)

func (r GatewayListGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
