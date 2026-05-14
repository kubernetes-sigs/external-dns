// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

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

// WorkflowService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r *WorkflowService) {
	r = &WorkflowService{}
	r.Options = opts
	return
}

// Gets Zaraz workflow for a zone.
func (r *WorkflowService) Get(ctx context.Context, query WorkflowGetParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz workflow
type Workflow string

const (
	WorkflowRealtime Workflow = "realtime"
	WorkflowPreview  Workflow = "preview"
)

func (r Workflow) IsKnown() bool {
	switch r {
	case WorkflowRealtime, WorkflowPreview:
		return true
	}
	return false
}

type WorkflowGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkflowGetResponseEnvelope struct {
	Errors   []WorkflowGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkflowGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result Workflow `json:"result,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	JSON    workflowGetResponseEnvelopeJSON `json:"-"`
}

// workflowGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkflowGetResponseEnvelope]
type workflowGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           WorkflowGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             workflowGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// workflowGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WorkflowGetResponseEnvelopeErrors]
type workflowGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    workflowGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// workflowGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeErrorsSource]
type workflowGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           WorkflowGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             workflowGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// workflowGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeMessages]
type workflowGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    workflowGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// workflowGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeMessagesSource]
type workflowGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
