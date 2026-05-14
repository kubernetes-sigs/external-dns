// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TokenService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *TokenService) New(ctx context.Context, identifier string, params TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	var env TokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/token", params.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenNewResponse struct {
	// The signed token used with the signed URLs feature.
	Token string               `json:"token"`
	JSON  tokenNewResponseJSON `json:"-"`
}

// tokenNewResponseJSON contains the JSON metadata for the struct
// [TokenNewResponse]
type tokenNewResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The optional ID of a Stream signing key. If present, the `pem` field is also
	// required.
	ID param.Field[string] `json:"id"`
	// The optional list of access rule constraints on the token. Access can be blocked
	// or allowed based on an IP, IP range, or by country. Access rules are evaluated
	// from first to last. If a rule matches, the associated action is applied and no
	// further rules are evaluated.
	AccessRules param.Field[[]TokenNewParamsAccessRule] `json:"accessRules"`
	// The optional boolean value that enables using signed tokens to access MP4
	// download links for a video.
	Downloadable param.Field[bool] `json:"downloadable"`
	// The optional unix epoch timestamp that specficies the time after a token is not
	// accepted. The maximum time specification is 24 hours from issuing time. If this
	// field is not set, the default is one hour after issuing.
	Exp param.Field[int64] `json:"exp"`
	// The optional unix epoch timestamp that specifies the time before a the token is
	// not accepted. If this field is not set, the default is one hour before issuing.
	Nbf param.Field[int64] `json:"nbf"`
	// The optional base64 encoded private key in PEM format associated with a Stream
	// signing key. If present, the `id` field is also required.
	Pem param.Field[string] `json:"pem"`
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines rules for fine-grained control over content than signed URL tokens
// alone. Access rules primarily make tokens conditionally valid based on user
// information. Access Rules are specified on token payloads as the `accessRules`
// property containing an array of Rule objects.
type TokenNewParamsAccessRule struct {
	// The action to take when a request matches a rule. If the action is `block`, the
	// signed token blocks views for viewers matching the rule.
	Action param.Field[TokenNewParamsAccessRulesAction] `json:"action"`
	// An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match
	// requests.
	Country param.Field[[]string] `json:"country"`
	// An array of IPv4 or IPV6 addresses or CIDRs used to match requests.
	IP param.Field[[]string] `json:"ip"`
	// Lists available rule types to match for requests. An `any` type matches all
	// requests and can be used as a wildcard to apply default actions after other
	// rules.
	Type param.Field[TokenNewParamsAccessRulesType] `json:"type"`
}

func (r TokenNewParamsAccessRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when a request matches a rule. If the action is `block`, the
// signed token blocks views for viewers matching the rule.
type TokenNewParamsAccessRulesAction string

const (
	TokenNewParamsAccessRulesActionAllow TokenNewParamsAccessRulesAction = "allow"
	TokenNewParamsAccessRulesActionBlock TokenNewParamsAccessRulesAction = "block"
)

func (r TokenNewParamsAccessRulesAction) IsKnown() bool {
	switch r {
	case TokenNewParamsAccessRulesActionAllow, TokenNewParamsAccessRulesActionBlock:
		return true
	}
	return false
}

// Lists available rule types to match for requests. An `any` type matches all
// requests and can be used as a wildcard to apply default actions after other
// rules.
type TokenNewParamsAccessRulesType string

const (
	TokenNewParamsAccessRulesTypeAny            TokenNewParamsAccessRulesType = "any"
	TokenNewParamsAccessRulesTypeIPSrc          TokenNewParamsAccessRulesType = "ip.src"
	TokenNewParamsAccessRulesTypeIPGeoipCountry TokenNewParamsAccessRulesType = "ip.geoip.country"
)

func (r TokenNewParamsAccessRulesType) IsKnown() bool {
	switch r {
	case TokenNewParamsAccessRulesTypeAny, TokenNewParamsAccessRulesTypeIPSrc, TokenNewParamsAccessRulesTypeIPGeoipCountry:
		return true
	}
	return false
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
