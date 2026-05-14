// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed

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

// ConfigASNService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigASNService] method instead.
type ConfigASNService struct {
	Options []option.RequestOption
}

// NewConfigASNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigASNService(opts ...option.RequestOption) (r *ConfigASNService) {
	r = &ConfigASNService{}
	r.Options = opts
	return
}

// Delete an ASN from botnet threat feed for a given user.
func (r *ConfigASNService) Delete(ctx context.Context, asnID int64, body ConfigASNDeleteParams, opts ...option.RequestOption) (res *ConfigASNDeleteResponse, err error) {
	var env ConfigASNDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/configs/asn/%v", body.AccountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.
func (r *ConfigASNService) Get(ctx context.Context, query ConfigASNGetParams, opts ...option.RequestOption) (res *ConfigASNGetResponse, err error) {
	var env ConfigASNGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/configs/asn", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigASNDeleteResponse struct {
	ASN  int64                       `json:"asn"`
	JSON configASNDeleteResponseJSON `json:"-"`
}

// configASNDeleteResponseJSON contains the JSON metadata for the struct
// [ConfigASNDeleteResponse]
type configASNDeleteResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigASNGetResponse struct {
	ASN  int64                    `json:"asn"`
	JSON configASNGetResponseJSON `json:"-"`
}

// configASNGetResponseJSON contains the JSON metadata for the struct
// [ConfigASNGetResponse]
type configASNGetResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigASNDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigASNDeleteResponseEnvelope struct {
	Errors   []ConfigASNDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigASNDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ConfigASNDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ConfigASNDeleteResponse                `json:"result"`
	JSON    configASNDeleteResponseEnvelopeJSON    `json:"-"`
}

// configASNDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigASNDeleteResponseEnvelope]
type configASNDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigASNDeleteResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ConfigASNDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             configASNDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// configASNDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConfigASNDeleteResponseEnvelopeErrors]
type configASNDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigASNDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigASNDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    configASNDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configASNDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ConfigASNDeleteResponseEnvelopeErrorsSource]
type configASNDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigASNDeleteResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ConfigASNDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             configASNDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// configASNDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigASNDeleteResponseEnvelopeMessages]
type configASNDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigASNDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigASNDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    configASNDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configASNDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ConfigASNDeleteResponseEnvelopeMessagesSource]
type configASNDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigASNDeleteResponseEnvelopeSuccess bool

const (
	ConfigASNDeleteResponseEnvelopeSuccessTrue ConfigASNDeleteResponseEnvelopeSuccess = true
)

func (r ConfigASNDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigASNDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigASNGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigASNGetResponseEnvelope struct {
	Errors   []ConfigASNGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigASNGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ConfigASNGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ConfigASNGetResponse                `json:"result"`
	JSON    configASNGetResponseEnvelopeJSON    `json:"-"`
}

// configASNGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigASNGetResponseEnvelope]
type configASNGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigASNGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ConfigASNGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             configASNGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// configASNGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigASNGetResponseEnvelopeErrors]
type configASNGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigASNGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigASNGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    configASNGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configASNGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ConfigASNGetResponseEnvelopeErrorsSource]
type configASNGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigASNGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ConfigASNGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             configASNGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// configASNGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigASNGetResponseEnvelopeMessages]
type configASNGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigASNGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigASNGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    configASNGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configASNGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ConfigASNGetResponseEnvelopeMessagesSource]
type configASNGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigASNGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configASNGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigASNGetResponseEnvelopeSuccess bool

const (
	ConfigASNGetResponseEnvelopeSuccessTrue ConfigASNGetResponseEnvelopeSuccess = true
)

func (r ConfigASNGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigASNGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
