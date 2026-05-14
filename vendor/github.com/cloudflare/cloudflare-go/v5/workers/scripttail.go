// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// ScriptTailService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptTailService] method instead.
type ScriptTailService struct {
	Options []option.RequestOption
}

// NewScriptTailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptTailService(opts ...option.RequestOption) (r *ScriptTailService) {
	r = &ScriptTailService{}
	r.Options = opts
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *ScriptTailService) New(ctx context.Context, scriptName string, params ScriptTailNewParams, opts ...option.RequestOption) (res *ScriptTailNewResponse, err error) {
	var env ScriptTailNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a tail from a Worker.
func (r *ScriptTailService) Delete(ctx context.Context, scriptName string, id string, body ScriptTailDeleteParams, opts ...option.RequestOption) (res *ScriptTailDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", body.AccountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get list of tails currently deployed on a Worker.
func (r *ScriptTailService) Get(ctx context.Context, scriptName string, query ScriptTailGetParams, opts ...option.RequestOption) (res *ScriptTailGetResponse, err error) {
	var env ScriptTailGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A reference to a script that will consume logs from the attached Worker.
type ConsumerScript struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string             `json:"namespace"`
	JSON      consumerScriptJSON `json:"-"`
}

// consumerScriptJSON contains the JSON metadata for the struct [ConsumerScript]
type consumerScriptJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsumerScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consumerScriptJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ConsumerScriptParam struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ConsumerScriptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptTailNewResponse struct {
	// Identifier.
	ID        string                    `json:"id,required"`
	ExpiresAt string                    `json:"expires_at,required"`
	URL       string                    `json:"url,required"`
	JSON      scriptTailNewResponseJSON `json:"-"`
}

// scriptTailNewResponseJSON contains the JSON metadata for the struct
// [ScriptTailNewResponse]
type scriptTailNewResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponse struct {
	Errors   []ScriptTailDeleteResponseError   `json:"errors,required"`
	Messages []ScriptTailDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptTailDeleteResponseSuccess `json:"success,required"`
	JSON    scriptTailDeleteResponseJSON    `json:"-"`
}

// scriptTailDeleteResponseJSON contains the JSON metadata for the struct
// [ScriptTailDeleteResponse]
type scriptTailDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           ScriptTailDeleteResponseErrorsSource `json:"source"`
	JSON             scriptTailDeleteResponseErrorJSON    `json:"-"`
}

// scriptTailDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ScriptTailDeleteResponseError]
type scriptTailDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    scriptTailDeleteResponseErrorsSourceJSON `json:"-"`
}

// scriptTailDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptTailDeleteResponseErrorsSource]
type scriptTailDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ScriptTailDeleteResponseMessagesSource `json:"source"`
	JSON             scriptTailDeleteResponseMessageJSON    `json:"-"`
}

// scriptTailDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ScriptTailDeleteResponseMessage]
type scriptTailDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ScriptTailDeleteResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    scriptTailDeleteResponseMessagesSourceJSON `json:"-"`
}

// scriptTailDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ScriptTailDeleteResponseMessagesSource]
type scriptTailDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptTailDeleteResponseSuccess bool

const (
	ScriptTailDeleteResponseSuccessTrue ScriptTailDeleteResponseSuccess = true
)

func (r ScriptTailDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ScriptTailDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ScriptTailGetResponse struct {
	// Identifier.
	ID        string                    `json:"id,required"`
	ExpiresAt string                    `json:"expires_at,required"`
	URL       string                    `json:"url,required"`
	JSON      scriptTailGetResponseJSON `json:"-"`
}

// scriptTailGetResponseJSON contains the JSON metadata for the struct
// [ScriptTailGetResponse]
type scriptTailGetResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r ScriptTailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptTailNewResponseEnvelope struct {
	Errors   []ScriptTailNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptTailNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptTailNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptTailNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptTailNewResponseEnvelopeJSON    `json:"-"`
}

// scriptTailNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailNewResponseEnvelope]
type scriptTailNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ScriptTailNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptTailNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptTailNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptTailNewResponseEnvelopeErrors]
type scriptTailNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    scriptTailNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptTailNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptTailNewResponseEnvelopeErrorsSource]
type scriptTailNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ScriptTailNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptTailNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptTailNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptTailNewResponseEnvelopeMessages]
type scriptTailNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    scriptTailNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptTailNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptTailNewResponseEnvelopeMessagesSource]
type scriptTailNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptTailNewResponseEnvelopeSuccess bool

const (
	ScriptTailNewResponseEnvelopeSuccessTrue ScriptTailNewResponseEnvelopeSuccess = true
)

func (r ScriptTailNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptTailDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptTailGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptTailGetResponseEnvelope struct {
	Errors   []ScriptTailGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptTailGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptTailGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptTailGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptTailGetResponseEnvelopeJSON    `json:"-"`
}

// scriptTailGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailGetResponseEnvelope]
type scriptTailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptTailGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ScriptTailGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptTailGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptTailGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptTailGetResponseEnvelopeErrors]
type scriptTailGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptTailGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    scriptTailGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptTailGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptTailGetResponseEnvelopeErrorsSource]
type scriptTailGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptTailGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ScriptTailGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptTailGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptTailGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptTailGetResponseEnvelopeMessages]
type scriptTailGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptTailGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    scriptTailGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptTailGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptTailGetResponseEnvelopeMessagesSource]
type scriptTailGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptTailGetResponseEnvelopeSuccess bool

const (
	ScriptTailGetResponseEnvelopeSuccessTrue ScriptTailGetResponseEnvelopeSuccess = true
)

func (r ScriptTailGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
