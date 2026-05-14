// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// InstanceStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceStatusService] method instead.
type InstanceStatusService struct {
	Options []option.RequestOption
}

// NewInstanceStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceStatusService(opts ...option.RequestOption) (r *InstanceStatusService) {
	r = &InstanceStatusService{}
	r.Options = opts
	return
}

// Change status of instance
func (r *InstanceStatusService) Edit(ctx context.Context, workflowName string, instanceID string, params InstanceStatusEditParams, opts ...option.RequestOption) (res *InstanceStatusEditResponse, err error) {
	var env InstanceStatusEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s/status", params.AccountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceStatusEditResponse struct {
	Status InstanceStatusEditResponseStatus `json:"status,required"`
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	Timestamp time.Time                      `json:"timestamp,required" format:"date-time"`
	JSON      instanceStatusEditResponseJSON `json:"-"`
}

// instanceStatusEditResponseJSON contains the JSON metadata for the struct
// [InstanceStatusEditResponse]
type instanceStatusEditResponseJSON struct {
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceStatusEditResponseStatus string

const (
	InstanceStatusEditResponseStatusQueued          InstanceStatusEditResponseStatus = "queued"
	InstanceStatusEditResponseStatusRunning         InstanceStatusEditResponseStatus = "running"
	InstanceStatusEditResponseStatusPaused          InstanceStatusEditResponseStatus = "paused"
	InstanceStatusEditResponseStatusErrored         InstanceStatusEditResponseStatus = "errored"
	InstanceStatusEditResponseStatusTerminated      InstanceStatusEditResponseStatus = "terminated"
	InstanceStatusEditResponseStatusComplete        InstanceStatusEditResponseStatus = "complete"
	InstanceStatusEditResponseStatusWaitingForPause InstanceStatusEditResponseStatus = "waitingForPause"
	InstanceStatusEditResponseStatusWaiting         InstanceStatusEditResponseStatus = "waiting"
)

func (r InstanceStatusEditResponseStatus) IsKnown() bool {
	switch r {
	case InstanceStatusEditResponseStatusQueued, InstanceStatusEditResponseStatusRunning, InstanceStatusEditResponseStatusPaused, InstanceStatusEditResponseStatusErrored, InstanceStatusEditResponseStatusTerminated, InstanceStatusEditResponseStatusComplete, InstanceStatusEditResponseStatusWaitingForPause, InstanceStatusEditResponseStatusWaiting:
		return true
	}
	return false
}

type InstanceStatusEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Apply action to instance.
	Status param.Field[InstanceStatusEditParamsStatus] `json:"status,required"`
}

func (r InstanceStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Apply action to instance.
type InstanceStatusEditParamsStatus string

const (
	InstanceStatusEditParamsStatusResume    InstanceStatusEditParamsStatus = "resume"
	InstanceStatusEditParamsStatusPause     InstanceStatusEditParamsStatus = "pause"
	InstanceStatusEditParamsStatusTerminate InstanceStatusEditParamsStatus = "terminate"
)

func (r InstanceStatusEditParamsStatus) IsKnown() bool {
	switch r {
	case InstanceStatusEditParamsStatusResume, InstanceStatusEditParamsStatusPause, InstanceStatusEditParamsStatusTerminate:
		return true
	}
	return false
}

type InstanceStatusEditResponseEnvelope struct {
	Errors     []InstanceStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []InstanceStatusEditResponseEnvelopeMessages `json:"messages,required"`
	Result     InstanceStatusEditResponse                   `json:"result,required"`
	Success    InstanceStatusEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo InstanceStatusEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       instanceStatusEditResponseEnvelopeJSON       `json:"-"`
}

// instanceStatusEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceStatusEditResponseEnvelope]
type instanceStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceStatusEditResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    instanceStatusEditResponseEnvelopeErrorsJSON `json:"-"`
}

// instanceStatusEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InstanceStatusEditResponseEnvelopeErrors]
type instanceStatusEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatusEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InstanceStatusEditResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    instanceStatusEditResponseEnvelopeMessagesJSON `json:"-"`
}

// instanceStatusEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InstanceStatusEditResponseEnvelopeMessages]
type instanceStatusEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatusEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InstanceStatusEditResponseEnvelopeSuccess bool

const (
	InstanceStatusEditResponseEnvelopeSuccessTrue InstanceStatusEditResponseEnvelopeSuccess = true
)

func (r InstanceStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InstanceStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InstanceStatusEditResponseEnvelopeResultInfo struct {
	Count      float64                                          `json:"count,required"`
	PerPage    float64                                          `json:"per_page,required"`
	TotalCount float64                                          `json:"total_count,required"`
	NextCursor string                                           `json:"next_cursor"`
	Page       float64                                          `json:"page"`
	JSON       instanceStatusEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// instanceStatusEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [InstanceStatusEditResponseEnvelopeResultInfo]
type instanceStatusEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	NextCursor  apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatusEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatusEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
