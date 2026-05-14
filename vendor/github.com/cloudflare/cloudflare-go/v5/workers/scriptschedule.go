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

// ScriptScheduleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptScheduleService] method instead.
type ScriptScheduleService struct {
	Options []option.RequestOption
}

// NewScriptScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptScheduleService(opts ...option.RequestOption) (r *ScriptScheduleService) {
	r = &ScriptScheduleService{}
	r.Options = opts
	return
}

// Updates Cron Triggers for a Worker.
func (r *ScriptScheduleService) Update(ctx context.Context, scriptName string, params ScriptScheduleUpdateParams, opts ...option.RequestOption) (res *ScriptScheduleUpdateResponse, err error) {
	var env ScriptScheduleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Cron Triggers for a Worker.
func (r *ScriptScheduleService) Get(ctx context.Context, scriptName string, query ScriptScheduleGetParams, opts ...option.RequestOption) (res *ScriptScheduleGetResponse, err error) {
	var env ScriptScheduleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptScheduleUpdateResponse struct {
	Schedules []ScriptScheduleUpdateResponseSchedule `json:"schedules,required"`
	JSON      scriptScheduleUpdateResponseJSON       `json:"-"`
}

// scriptScheduleUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptScheduleUpdateResponse]
type scriptScheduleUpdateResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseSchedule struct {
	Cron       string                                   `json:"cron,required"`
	CreatedOn  string                                   `json:"created_on"`
	ModifiedOn string                                   `json:"modified_on"`
	JSON       scriptScheduleUpdateResponseScheduleJSON `json:"-"`
}

// scriptScheduleUpdateResponseScheduleJSON contains the JSON metadata for the
// struct [ScriptScheduleUpdateResponseSchedule]
type scriptScheduleUpdateResponseScheduleJSON struct {
	Cron        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponse struct {
	Schedules []ScriptScheduleGetResponseSchedule `json:"schedules,required"`
	JSON      scriptScheduleGetResponseJSON       `json:"-"`
}

// scriptScheduleGetResponseJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponse]
type scriptScheduleGetResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseSchedule struct {
	Cron       string                                `json:"cron,required"`
	CreatedOn  string                                `json:"created_on"`
	ModifiedOn string                                `json:"modified_on"`
	JSON       scriptScheduleGetResponseScheduleJSON `json:"-"`
}

// scriptScheduleGetResponseScheduleJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponseSchedule]
type scriptScheduleGetResponseScheduleJSON struct {
	Cron        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateParams struct {
	// Identifier.
	AccountID param.Field[string]              `path:"account_id,required"`
	Body      []ScriptScheduleUpdateParamsBody `json:"body,required"`
}

func (r ScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptScheduleUpdateParamsBody struct {
	Cron param.Field[string] `json:"cron,required"`
}

func (r ScriptScheduleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptScheduleUpdateResponseEnvelope struct {
	Errors   []ScriptScheduleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScheduleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptScheduleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptScheduleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptScheduleUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptScheduleUpdateResponseEnvelope]
type scriptScheduleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ScriptScheduleUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptScheduleUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptScheduleUpdateResponseEnvelopeErrors]
type scriptScheduleUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    scriptScheduleUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ScriptScheduleUpdateResponseEnvelopeErrorsSource]
type scriptScheduleUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ScriptScheduleUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptScheduleUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptScheduleUpdateResponseEnvelopeMessages]
type scriptScheduleUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    scriptScheduleUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ScriptScheduleUpdateResponseEnvelopeMessagesSource]
type scriptScheduleUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptScheduleUpdateResponseEnvelopeSuccess bool

const (
	ScriptScheduleUpdateResponseEnvelopeSuccessTrue ScriptScheduleUpdateResponseEnvelopeSuccess = true
)

func (r ScriptScheduleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScheduleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptScheduleGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptScheduleGetResponseEnvelope struct {
	Errors   []ScriptScheduleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScheduleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptScheduleGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptScheduleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptScheduleGetResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponseEnvelope]
type scriptScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ScriptScheduleGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptScheduleGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptScheduleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptScheduleGetResponseEnvelopeErrors]
type scriptScheduleGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    scriptScheduleGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptScheduleGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptScheduleGetResponseEnvelopeErrorsSource]
type scriptScheduleGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ScriptScheduleGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptScheduleGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptScheduleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptScheduleGetResponseEnvelopeMessages]
type scriptScheduleGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    scriptScheduleGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptScheduleGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptScheduleGetResponseEnvelopeMessagesSource]
type scriptScheduleGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptScheduleGetResponseEnvelopeSuccess bool

const (
	ScriptScheduleGetResponseEnvelopeSuccessTrue ScriptScheduleGetResponseEnvelopeSuccess = true
)

func (r ScriptScheduleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScheduleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
