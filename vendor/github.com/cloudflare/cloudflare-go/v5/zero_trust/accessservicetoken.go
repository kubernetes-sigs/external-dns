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
)

// AccessServiceTokenService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessServiceTokenService] method instead.
type AccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessServiceTokenService(opts ...option.RequestOption) (r *AccessServiceTokenService) {
	r = &AccessServiceTokenService{}
	r.Options = opts
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccessServiceTokenService) New(ctx context.Context, params AccessServiceTokenNewParams, opts ...option.RequestOption) (res *AccessServiceTokenNewResponse, err error) {
	var env AccessServiceTokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured service token.
func (r *AccessServiceTokenService) Update(ctx context.Context, serviceTokenID string, params AccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all service tokens.
func (r *AccessServiceTokenService) List(ctx context.Context, params AccessServiceTokenListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ServiceToken], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
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

// Lists all service tokens.
func (r *AccessServiceTokenService) ListAutoPaging(ctx context.Context, params AccessServiceTokenListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ServiceToken] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a service token.
func (r *AccessServiceTokenService) Delete(ctx context.Context, serviceTokenID string, body AccessServiceTokenDeleteParams, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single service token.
func (r *AccessServiceTokenService) Get(ctx context.Context, serviceTokenID string, query AccessServiceTokenGetParams, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Refreshes the expiration of a service token.
func (r *AccessServiceTokenService) Refresh(ctx context.Context, serviceTokenID string, body AccessServiceTokenRefreshParams, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenRefreshResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", body.AccountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccessServiceTokenService) Rotate(ctx context.Context, serviceTokenID string, body AccessServiceTokenRotateParams, opts ...option.RequestOption) (res *AccessServiceTokenRotateResponse, err error) {
	var env AccessServiceTokenRotateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", body.AccountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceToken struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration   string    `json:"duration"`
	ExpiresAt  time.Time `json:"expires_at" format:"date-time"`
	LastSeenAt time.Time `json:"last_seen_at" format:"date-time"`
	// The name of the service token.
	Name      string           `json:"name"`
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      serviceTokenJSON `json:"-"`
}

// serviceTokenJSON contains the JSON metadata for the struct [ServiceToken]
type serviceTokenJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	LastSeenAt  apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceTokenJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                            `json:"name"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenNewResponseJSON `json:"-"`
}

// accessServiceTokenNewResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenNewResponse]
type accessServiceTokenNewResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRotateResponseJSON `json:"-"`
}

// accessServiceTokenRotateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenRotateResponse]
type accessServiceTokenRotateResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccessServiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenNewResponseEnvelope struct {
	Errors   []AccessServiceTokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessServiceTokenNewResponse                `json:"result"`
	JSON    accessServiceTokenNewResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenNewResponseEnvelope]
type accessServiceTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccessServiceTokenNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessServiceTokenNewResponseEnvelopeErrors]
type accessServiceTokenNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accessServiceTokenNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessServiceTokenNewResponseEnvelopeErrorsSource]
type accessServiceTokenNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessServiceTokenNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessServiceTokenNewResponseEnvelopeMessages]
type accessServiceTokenNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessServiceTokenNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenNewResponseEnvelopeMessagesSource]
type accessServiceTokenNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenNewResponseEnvelopeSuccess bool

const (
	AccessServiceTokenNewResponseEnvelopeSuccessTrue AccessServiceTokenNewResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenUpdateResponseEnvelope struct {
	Errors   []AccessServiceTokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                    `json:"result"`
	JSON    accessServiceTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseEnvelope]
type accessServiceTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenUpdateResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessServiceTokenUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeErrors]
type accessServiceTokenUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessServiceTokenUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenUpdateResponseEnvelopeErrorsSource]
type accessServiceTokenUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenUpdateResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessServiceTokenUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeMessages]
type accessServiceTokenUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessServiceTokenUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenUpdateResponseEnvelopeMessagesSource]
type accessServiceTokenUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenUpdateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenUpdateResponseEnvelopeSuccessTrue AccessServiceTokenUpdateResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The name of the service token.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Search for service tokens by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessServiceTokenListParams]'s query parameters as
// `url.Values`.
func (r AccessServiceTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessServiceTokenDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessServiceTokenDeleteResponseEnvelope struct {
	Errors   []AccessServiceTokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                    `json:"result"`
	JSON    accessServiceTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseEnvelope]
type accessServiceTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessServiceTokenDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeErrors]
type accessServiceTokenDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessServiceTokenDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenDeleteResponseEnvelopeErrorsSource]
type accessServiceTokenDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessServiceTokenDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeMessages]
type accessServiceTokenDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessServiceTokenDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenDeleteResponseEnvelopeMessagesSource]
type accessServiceTokenDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenDeleteResponseEnvelopeSuccess bool

const (
	AccessServiceTokenDeleteResponseEnvelopeSuccessTrue AccessServiceTokenDeleteResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessServiceTokenGetResponseEnvelope struct {
	Errors   []AccessServiceTokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                 `json:"result"`
	JSON    accessServiceTokenGetResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenGetResponseEnvelope]
type accessServiceTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccessServiceTokenGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessServiceTokenGetResponseEnvelopeErrors]
type accessServiceTokenGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accessServiceTokenGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessServiceTokenGetResponseEnvelopeErrorsSource]
type accessServiceTokenGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessServiceTokenGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessServiceTokenGetResponseEnvelopeMessages]
type accessServiceTokenGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessServiceTokenGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenGetResponseEnvelopeMessagesSource]
type accessServiceTokenGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenGetResponseEnvelopeSuccess bool

const (
	AccessServiceTokenGetResponseEnvelopeSuccessTrue AccessServiceTokenGetResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenRefreshParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessServiceTokenRefreshResponseEnvelope struct {
	Errors   []AccessServiceTokenRefreshResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenRefreshResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenRefreshResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                     `json:"result"`
	JSON    accessServiceTokenRefreshResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRefreshResponseEnvelope]
type accessServiceTokenRefreshResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRefreshResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccessServiceTokenRefreshResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenRefreshResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenRefreshResponseEnvelopeErrors]
type accessServiceTokenRefreshResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRefreshResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accessServiceTokenRefreshResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenRefreshResponseEnvelopeErrorsSource]
type accessServiceTokenRefreshResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRefreshResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessServiceTokenRefreshResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenRefreshResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenRefreshResponseEnvelopeMessages]
type accessServiceTokenRefreshResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRefreshResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessServiceTokenRefreshResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessServiceTokenRefreshResponseEnvelopeMessagesSource]
type accessServiceTokenRefreshResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenRefreshResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRefreshResponseEnvelopeSuccessTrue AccessServiceTokenRefreshResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenRefreshResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenRefreshResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenRotateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessServiceTokenRotateResponseEnvelope struct {
	Errors   []AccessServiceTokenRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenRotateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessServiceTokenRotateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessServiceTokenRotateResponse                `json:"result"`
	JSON    accessServiceTokenRotateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRotateResponseEnvelope]
type accessServiceTokenRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessServiceTokenRotateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessServiceTokenRotateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenRotateResponseEnvelopeErrors]
type accessServiceTokenRotateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessServiceTokenRotateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenRotateResponseEnvelopeErrorsSource]
type accessServiceTokenRotateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessServiceTokenRotateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessServiceTokenRotateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenRotateResponseEnvelopeMessages]
type accessServiceTokenRotateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessServiceTokenRotateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessServiceTokenRotateResponseEnvelopeMessagesSource]
type accessServiceTokenRotateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessServiceTokenRotateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRotateResponseEnvelopeSuccessTrue AccessServiceTokenRotateResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenRotateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenRotateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
