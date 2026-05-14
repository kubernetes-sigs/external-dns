// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// TokenValueService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenValueService] method instead.
type TokenValueService struct {
	Options []option.RequestOption
}

// NewTokenValueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenValueService(opts ...option.RequestOption) (r *TokenValueService) {
	r = &TokenValueService{}
	r.Options = opts
	return
}

// Roll the token secret.
func (r *TokenValueService) Update(ctx context.Context, tokenID string, body TokenValueUpdateParams, opts ...option.RequestOption) (res *shared.TokenValue, err error) {
	var env TokenValueUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("user/tokens/%s/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenValueUpdateParams struct {
	Body interface{} `json:"body,required"`
}

func (r TokenValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TokenValueUpdateResponseEnvelope struct {
	Errors   []TokenValueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenValueUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TokenValueUpdateResponseEnvelopeSuccess `json:"success,required"`
	// The token value.
	Result shared.TokenValue                    `json:"result"`
	JSON   tokenValueUpdateResponseEnvelopeJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenValueUpdateResponseEnvelope]
type tokenValueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           TokenValueUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenValueUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenValueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TokenValueUpdateResponseEnvelopeErrors]
type tokenValueUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    tokenValueUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [TokenValueUpdateResponseEnvelopeErrorsSource]
type tokenValueUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           TokenValueUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenValueUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenValueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TokenValueUpdateResponseEnvelopeMessages]
type tokenValueUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenValueUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    tokenValueUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [TokenValueUpdateResponseEnvelopeMessagesSource]
type tokenValueUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenValueUpdateResponseEnvelopeSuccess bool

const (
	TokenValueUpdateResponseEnvelopeSuccessTrue TokenValueUpdateResponseEnvelopeSuccess = true
)

func (r TokenValueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenValueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
