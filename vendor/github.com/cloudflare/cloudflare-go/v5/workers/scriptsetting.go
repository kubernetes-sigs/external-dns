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

// ScriptSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSettingService] method instead.
type ScriptSettingService struct {
	Options []option.RequestOption
}

// NewScriptSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSettingService(opts ...option.RequestOption) (r *ScriptSettingService) {
	r = &ScriptSettingService{}
	r.Options = opts
	return
}

// Patch script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Including but not limited to Logpush and Tail Consumers.
func (r *ScriptSettingService) Edit(ctx context.Context, scriptName string, params ScriptSettingEditParams, opts ...option.RequestOption) (res *ScriptSetting, err error) {
	var env ScriptSettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Includes Logpush and Tail Consumers.
func (r *ScriptSettingService) Get(ctx context.Context, scriptName string, query ScriptSettingGetParams, opts ...option.RequestOption) (res *ScriptSetting, err error) {
	var env ScriptSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSettingEditParams struct {
	// Identifier.
	AccountID     param.Field[string] `path:"account_id,required"`
	ScriptSetting ScriptSettingParam  `json:"script_setting,required"`
}

func (r ScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ScriptSetting)
}

type ScriptSettingEditResponseEnvelope struct {
	Errors   []ScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSetting                               `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponseEnvelope]
type scriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ScriptSettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSettingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseEnvelopeErrors]
type scriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    scriptSettingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSettingEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptSettingEditResponseEnvelopeErrorsSource]
type scriptSettingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptSettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSettingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseEnvelopeMessages]
type scriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptSettingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSettingEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSettingEditResponseEnvelopeMessagesSource]
type scriptSettingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSettingEditResponseEnvelopeSuccess bool

const (
	ScriptSettingEditResponseEnvelopeSuccessTrue ScriptSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSettingGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSettingGetResponseEnvelope struct {
	Errors   []ScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSetting                              `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponseEnvelope]
type scriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptSettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptSettingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseEnvelopeErrors]
type scriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptSettingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptSettingGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptSettingGetResponseEnvelopeErrorsSource]
type scriptSettingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptSettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptSettingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseEnvelopeMessages]
type scriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptSettingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptSettingGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptSettingGetResponseEnvelopeMessagesSource]
type scriptSettingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptSettingGetResponseEnvelopeSuccess bool

const (
	ScriptSettingGetResponseEnvelopeSuccessTrue ScriptSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
