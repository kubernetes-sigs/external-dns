// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

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

// InstanceEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceEventService] method instead.
type InstanceEventService struct {
	Options []option.RequestOption
}

// NewInstanceEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceEventService(opts ...option.RequestOption) (r *InstanceEventService) {
	r = &InstanceEventService{}
	r.Options = opts
	return
}

// Send event to instance
func (r *InstanceEventService) New(ctx context.Context, workflowName string, instanceID string, eventType string, params InstanceEventNewParams, opts ...option.RequestOption) (res *InstanceEventNewResponse, err error) {
	var env InstanceEventNewResponseEnvelope
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
	if eventType == "" {
		err = errors.New("missing required event_type parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s/events/%s", params.AccountID, workflowName, instanceID, eventType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceEventNewResponse = interface{}

type InstanceEventNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body"`
}

func (r InstanceEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type InstanceEventNewResponseEnvelope struct {
	Errors     []InstanceEventNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []InstanceEventNewResponseEnvelopeMessages `json:"messages,required"`
	Success    InstanceEventNewResponseEnvelopeSuccess    `json:"success,required"`
	Result     InstanceEventNewResponse                   `json:"result"`
	ResultInfo InstanceEventNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       instanceEventNewResponseEnvelopeJSON       `json:"-"`
}

// instanceEventNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceEventNewResponseEnvelope]
type instanceEventNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceEventNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceEventNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceEventNewResponseEnvelopeErrors struct {
	Code    float64                                    `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    instanceEventNewResponseEnvelopeErrorsJSON `json:"-"`
}

// instanceEventNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InstanceEventNewResponseEnvelopeErrors]
type instanceEventNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceEventNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceEventNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InstanceEventNewResponseEnvelopeMessages struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    instanceEventNewResponseEnvelopeMessagesJSON `json:"-"`
}

// instanceEventNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InstanceEventNewResponseEnvelopeMessages]
type instanceEventNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceEventNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceEventNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InstanceEventNewResponseEnvelopeSuccess bool

const (
	InstanceEventNewResponseEnvelopeSuccessTrue InstanceEventNewResponseEnvelopeSuccess = true
)

func (r InstanceEventNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InstanceEventNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InstanceEventNewResponseEnvelopeResultInfo struct {
	Count      float64                                        `json:"count,required"`
	PerPage    float64                                        `json:"per_page,required"`
	TotalCount float64                                        `json:"total_count,required"`
	NextCursor string                                         `json:"next_cursor"`
	Page       float64                                        `json:"page"`
	JSON       instanceEventNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// instanceEventNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [InstanceEventNewResponseEnvelopeResultInfo]
type instanceEventNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	NextCursor  apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceEventNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceEventNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
