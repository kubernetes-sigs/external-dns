// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

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

// TokenService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options          []option.RequestOption
	PermissionGroups *TokenPermissionGroupService
	Value            *TokenValueService
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	r.PermissionGroups = NewTokenPermissionGroupService(opts...)
	r.Value = NewTokenValueService(opts...)
	return
}

// Create a new Account Owned API token.
func (r *TokenService) New(ctx context.Context, params TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	var env TokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing token.
func (r *TokenService) Update(ctx context.Context, tokenID string, params TokenUpdateParams, opts ...option.RequestOption) (res *shared.Token, err error) {
	var env TokenUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", params.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Account Owned API tokens created for this account.
func (r *TokenService) List(ctx context.Context, params TokenListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.Token], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", params.AccountID)
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

// List all Account Owned API tokens created for this account.
func (r *TokenService) ListAutoPaging(ctx context.Context, params TokenListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.Token] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Destroy an Account Owned API token.
func (r *TokenService) Delete(ctx context.Context, tokenID string, body TokenDeleteParams, opts ...option.RequestOption) (res *TokenDeleteResponse, err error) {
	var env TokenDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", body.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific Account Owned API token.
func (r *TokenService) Get(ctx context.Context, tokenID string, query TokenGetParams, opts ...option.RequestOption) (res *shared.Token, err error) {
	var env TokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test whether a token works.
func (r *TokenService) Verify(ctx context.Context, query TokenVerifyParams, opts ...option.RequestOption) (res *TokenVerifyResponse, err error) {
	var env TokenVerifyResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/verify", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenNewResponse struct {
	// Token identifier tag.
	ID        string                    `json:"id"`
	Condition TokenNewResponseCondition `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time on which the token was created.
	IssuedOn time.Time `json:"issued_on" format:"date-time"`
	// Last time the token was used.
	LastUsedOn time.Time `json:"last_used_on" format:"date-time"`
	// Last time the token was modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Token name.
	Name string `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies []shared.TokenPolicy `json:"policies"`
	// Status of the token.
	Status TokenNewResponseStatus `json:"status"`
	// The token value.
	Value shared.TokenValue    `json:"value"`
	JSON  tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	ExpiresOn   apijson.Field
	IssuedOn    apijson.Field
	LastUsedOn  apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	NotBefore   apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseCondition struct {
	// Client IP restrictions.
	RequestIP TokenNewResponseConditionRequestIP `json:"request_ip"`
	JSON      tokenNewResponseConditionJSON      `json:"-"`
}

// tokenNewResponseConditionJSON contains the JSON metadata for the struct
// [TokenNewResponseCondition]
type tokenNewResponseConditionJSON struct {
	RequestIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseConditionJSON) RawJSON() string {
	return r.raw
}

// Client IP restrictions.
type TokenNewResponseConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In []shared.TokenConditionCIDRList `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn []shared.TokenConditionCIDRList        `json:"not_in"`
	JSON  tokenNewResponseConditionRequestIPJSON `json:"-"`
}

// tokenNewResponseConditionRequestIPJSON contains the JSON metadata for the struct
// [TokenNewResponseConditionRequestIP]
type tokenNewResponseConditionRequestIPJSON struct {
	In          apijson.Field
	NotIn       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseConditionRequestIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseConditionRequestIPJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenNewResponseStatus string

const (
	TokenNewResponseStatusActive   TokenNewResponseStatus = "active"
	TokenNewResponseStatusDisabled TokenNewResponseStatus = "disabled"
	TokenNewResponseStatusExpired  TokenNewResponseStatus = "expired"
)

func (r TokenNewResponseStatus) IsKnown() bool {
	switch r {
	case TokenNewResponseStatusActive, TokenNewResponseStatusDisabled, TokenNewResponseStatusExpired:
		return true
	}
	return false
}

type TokenDeleteResponse struct {
	// Identifier
	ID   string                  `json:"id,required"`
	JSON tokenDeleteResponseJSON `json:"-"`
}

// tokenDeleteResponseJSON contains the JSON metadata for the struct
// [TokenDeleteResponse]
type tokenDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponse struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status TokenVerifyResponseStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time               `json:"not_before" format:"date-time"`
	JSON      tokenVerifyResponseJSON `json:"-"`
}

// tokenVerifyResponseJSON contains the JSON metadata for the struct
// [TokenVerifyResponse]
type tokenVerifyResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenVerifyResponseStatus string

const (
	TokenVerifyResponseStatusActive   TokenVerifyResponseStatus = "active"
	TokenVerifyResponseStatusDisabled TokenVerifyResponseStatus = "disabled"
	TokenVerifyResponseStatusExpired  TokenVerifyResponseStatus = "expired"
)

func (r TokenVerifyResponseStatus) IsKnown() bool {
	switch r {
	case TokenVerifyResponseStatusActive, TokenVerifyResponseStatusDisabled, TokenVerifyResponseStatusExpired:
		return true
	}
	return false
}

type TokenNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]shared.TokenPolicyParam] `json:"policies,required"`
	Condition param.Field[TokenNewParamsCondition]   `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewParamsCondition struct {
	// Client IP restrictions.
	RequestIP param.Field[TokenNewParamsConditionRequestIP] `json:"request_ip"`
}

func (r TokenNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type TokenNewParamsConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]shared.TokenConditionCIDRListParam] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]shared.TokenConditionCIDRListParam] `json:"not_in"`
}

func (r TokenNewParamsConditionRequestIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenNewResponseEnvelope struct {
	Errors   []TokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenNewResponse                `json:"result"`
	JSON    tokenNewResponseEnvelopeJSON    `json:"-"`
}

// tokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelope]
type tokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           TokenNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelopeErrors]
type tokenNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    tokenNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TokenNewResponseEnvelopeErrorsSource]
type tokenNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           TokenNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TokenNewResponseEnvelopeMessages]
type tokenNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenNewResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    tokenNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TokenNewResponseEnvelopeMessagesSource]
type tokenNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenNewResponseEnvelopeSuccess bool

const (
	TokenNewResponseEnvelopeSuccessTrue TokenNewResponseEnvelopeSuccess = true
)

func (r TokenNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Token     shared.TokenParam   `json:"token,required"`
}

func (r TokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Token)
}

type TokenUpdateResponseEnvelope struct {
	Errors   []TokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Token                       `json:"result"`
	JSON    tokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// tokenUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelope]
type tokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TokenUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenUpdateResponseEnvelopeErrors]
type tokenUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    tokenUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TokenUpdateResponseEnvelopeErrorsSource]
type tokenUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TokenUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenUpdateResponseEnvelopeMessages]
type tokenUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    tokenUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TokenUpdateResponseEnvelopeMessagesSource]
type tokenUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenUpdateResponseEnvelopeSuccess bool

const (
	TokenUpdateResponseEnvelopeSuccessTrue TokenUpdateResponseEnvelopeSuccess = true
)

func (r TokenUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order results.
	Direction param.Field[TokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [TokenListParams]'s query parameters as `url.Values`.
func (r TokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type TokenListParamsDirection string

const (
	TokenListParamsDirectionAsc  TokenListParamsDirection = "asc"
	TokenListParamsDirectionDesc TokenListParamsDirection = "desc"
)

func (r TokenListParamsDirection) IsKnown() bool {
	switch r {
	case TokenListParamsDirectionAsc, TokenListParamsDirectionDesc:
		return true
	}
	return false
}

type TokenDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenDeleteResponseEnvelope struct {
	Errors   []TokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenDeleteResponse                `json:"result,nullable"`
	JSON    tokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// tokenDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelope]
type tokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TokenDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenDeleteResponseEnvelopeErrors]
type tokenDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    tokenDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TokenDeleteResponseEnvelopeErrorsSource]
type tokenDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TokenDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenDeleteResponseEnvelopeMessages]
type tokenDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    tokenDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TokenDeleteResponseEnvelopeMessagesSource]
type tokenDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenDeleteResponseEnvelopeSuccess bool

const (
	TokenDeleteResponseEnvelopeSuccessTrue TokenDeleteResponseEnvelopeSuccess = true
)

func (r TokenDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenGetResponseEnvelope struct {
	Errors   []TokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenGetResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Token                    `json:"result"`
	JSON    tokenGetResponseEnvelopeJSON    `json:"-"`
}

// tokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelope]
type tokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           TokenGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelopeErrors]
type tokenGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    tokenGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TokenGetResponseEnvelopeErrorsSource]
type tokenGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           TokenGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TokenGetResponseEnvelopeMessages]
type tokenGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    tokenGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TokenGetResponseEnvelopeMessagesSource]
type tokenGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenGetResponseEnvelopeSuccess bool

const (
	TokenGetResponseEnvelopeSuccessTrue TokenGetResponseEnvelopeSuccess = true
)

func (r TokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenVerifyParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TokenVerifyResponseEnvelope struct {
	Errors   []TokenVerifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenVerifyResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenVerifyResponseEnvelopeSuccess `json:"success,required"`
	Result  TokenVerifyResponse                `json:"result"`
	JSON    tokenVerifyResponseEnvelopeJSON    `json:"-"`
}

// tokenVerifyResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenVerifyResponseEnvelope]
type tokenVerifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           TokenVerifyResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenVerifyResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenVerifyResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TokenVerifyResponseEnvelopeErrors]
type tokenVerifyResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    tokenVerifyResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenVerifyResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [TokenVerifyResponseEnvelopeErrorsSource]
type tokenVerifyResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           TokenVerifyResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenVerifyResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenVerifyResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenVerifyResponseEnvelopeMessages]
type tokenVerifyResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenVerifyResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    tokenVerifyResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenVerifyResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [TokenVerifyResponseEnvelopeMessagesSource]
type tokenVerifyResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenVerifyResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenVerifyResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenVerifyResponseEnvelopeSuccess bool

const (
	TokenVerifyResponseEnvelopeSuccessTrue TokenVerifyResponseEnvelopeSuccess = true
)

func (r TokenVerifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenVerifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
