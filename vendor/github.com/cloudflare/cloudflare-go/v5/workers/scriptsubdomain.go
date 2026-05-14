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

// ScriptSubdomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSubdomainService] method instead.
type ScriptSubdomainService struct {
	Options []option.RequestOption
}

// NewScriptSubdomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSubdomainService(opts ...option.RequestOption) (r *ScriptSubdomainService) {
	r = &ScriptSubdomainService{}
	r.Options = opts
	return
}

// Enable or disable the Worker on the workers.dev subdomain.
func (r *ScriptSubdomainService) New(ctx context.Context, scriptName string, params ScriptSubdomainNewParams, opts ...option.RequestOption) (res *ScriptSubdomainNewResponse, err error) {
	var env ScriptSubdomainNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable all workers.dev subdomains for a Worker.
func (r *ScriptSubdomainService) Delete(ctx context.Context, scriptName string, body ScriptSubdomainDeleteParams, opts ...option.RequestOption) (res *ScriptSubdomainDeleteResponse, err error) {
	var env ScriptSubdomainDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", body.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get if the Worker is available on the workers.dev subdomain.
func (r *ScriptSubdomainService) Get(ctx context.Context, scriptName string, query ScriptSubdomainGetParams, opts ...option.RequestOption) (res *ScriptSubdomainGetResponse, err error) {
	var env ScriptSubdomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSubdomainNewResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled,required"`
	// Whether the Worker's Preview URLs are available on the workers.dev subdomain.
	PreviewsEnabled bool                           `json:"previews_enabled,required"`
	JSON            scriptSubdomainNewResponseJSON `json:"-"`
}

// scriptSubdomainNewResponseJSON contains the JSON metadata for the struct
// [ScriptSubdomainNewResponse]
type scriptSubdomainNewResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainDeleteResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled,required"`
	// Whether the Worker's Preview URLs are available on the workers.dev subdomain.
	PreviewsEnabled bool                              `json:"previews_enabled,required"`
	JSON            scriptSubdomainDeleteResponseJSON `json:"-"`
}

// scriptSubdomainDeleteResponseJSON contains the JSON metadata for the struct
// [ScriptSubdomainDeleteResponse]
type scriptSubdomainDeleteResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled,required"`
	// Whether the Worker's Preview URLs are available on the workers.dev subdomain.
	PreviewsEnabled bool                           `json:"previews_enabled,required"`
	JSON            scriptSubdomainGetResponseJSON `json:"-"`
}

// scriptSubdomainGetResponseJSON contains the JSON metadata for the struct
// [ScriptSubdomainGetResponse]
type scriptSubdomainGetResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether the Worker should be available on the workers.dev subdomain.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled param.Field[bool] `json:"previews_enabled"`
}

func (r ScriptSubdomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSubdomainNewResponseEnvelope struct {
	Errors   []ScriptSubdomainNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSubdomainNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSubdomainNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptSubdomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSubdomainNewResponseEnvelopeJSON    `json:"-"`
}

// scriptSubdomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSubdomainNewResponseEnvelope]
type scriptSubdomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptSubdomainNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSubdomainNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSubdomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSubdomainNewResponseEnvelopeErrors]
type scriptSubdomainNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptSubdomainNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSubdomainNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptSubdomainNewResponseEnvelopeErrorsSource]
type scriptSubdomainNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptSubdomainNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSubdomainNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSubdomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptSubdomainNewResponseEnvelopeMessages]
type scriptSubdomainNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptSubdomainNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSubdomainNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSubdomainNewResponseEnvelopeMessagesSource]
type scriptSubdomainNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSubdomainNewResponseEnvelopeSuccess bool

const (
	ScriptSubdomainNewResponseEnvelopeSuccessTrue ScriptSubdomainNewResponseEnvelopeSuccess = true
)

func (r ScriptSubdomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSubdomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSubdomainDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSubdomainDeleteResponseEnvelope struct {
	Errors   []ScriptSubdomainDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSubdomainDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSubdomainDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptSubdomainDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSubdomainDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptSubdomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptSubdomainDeleteResponseEnvelope]
type scriptSubdomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ScriptSubdomainDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSubdomainDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSubdomainDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptSubdomainDeleteResponseEnvelopeErrors]
type scriptSubdomainDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    scriptSubdomainDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSubdomainDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptSubdomainDeleteResponseEnvelopeErrorsSource]
type scriptSubdomainDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ScriptSubdomainDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSubdomainDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSubdomainDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptSubdomainDeleteResponseEnvelopeMessages]
type scriptSubdomainDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    scriptSubdomainDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSubdomainDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ScriptSubdomainDeleteResponseEnvelopeMessagesSource]
type scriptSubdomainDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSubdomainDeleteResponseEnvelopeSuccess bool

const (
	ScriptSubdomainDeleteResponseEnvelopeSuccessTrue ScriptSubdomainDeleteResponseEnvelopeSuccess = true
)

func (r ScriptSubdomainDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSubdomainDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSubdomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSubdomainGetResponseEnvelope struct {
	Errors   []ScriptSubdomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSubdomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSubdomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptSubdomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSubdomainGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSubdomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSubdomainGetResponseEnvelope]
type scriptSubdomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptSubdomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSubdomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSubdomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSubdomainGetResponseEnvelopeErrors]
type scriptSubdomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptSubdomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSubdomainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptSubdomainGetResponseEnvelopeErrorsSource]
type scriptSubdomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptSubdomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSubdomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSubdomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptSubdomainGetResponseEnvelopeMessages]
type scriptSubdomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptSubdomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSubdomainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSubdomainGetResponseEnvelopeMessagesSource]
type scriptSubdomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSubdomainGetResponseEnvelopeSuccess bool

const (
	ScriptSubdomainGetResponseEnvelopeSuccessTrue ScriptSubdomainGetResponseEnvelopeSuccess = true
)

func (r ScriptSubdomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSubdomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
