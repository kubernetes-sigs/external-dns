// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queues

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// PurgeService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPurgeService] method instead.
type PurgeService struct {
	Options []option.RequestOption
}

// NewPurgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPurgeService(opts ...option.RequestOption) (r *PurgeService) {
	r = &PurgeService{}
	r.Options = opts
	return
}

// Deletes all messages from the Queue.
func (r *PurgeService) Start(ctx context.Context, queueID string, params PurgeStartParams, opts ...option.RequestOption) (res *Queue, err error) {
	var env PurgeStartResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/purge", params.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details about a Queue's purge status.
func (r *PurgeService) Status(ctx context.Context, queueID string, query PurgeStatusParams, opts ...option.RequestOption) (res *PurgeStatusResponse, err error) {
	var env PurgeStatusResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/queues/%s/purge", query.AccountID, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PurgeStatusResponse struct {
	// Indicates if the last purge operation completed successfully.
	Completed string `json:"completed"`
	// Timestamp when the last purge operation started.
	StartedAt string                  `json:"started_at"`
	JSON      purgeStatusResponseJSON `json:"-"`
}

// purgeStatusResponseJSON contains the JSON metadata for the struct
// [PurgeStatusResponse]
type purgeStatusResponseJSON struct {
	Completed   apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r purgeStatusResponseJSON) RawJSON() string {
	return r.raw
}

type PurgeStartParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Confimation that all messages will be deleted permanently.
	DeleteMessagesPermanently param.Field[bool] `json:"delete_messages_permanently"`
}

func (r PurgeStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PurgeStartResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   Queue                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success PurgeStartResponseEnvelopeSuccess `json:"success"`
	JSON    purgeStartResponseEnvelopeJSON    `json:"-"`
}

// purgeStartResponseEnvelopeJSON contains the JSON metadata for the struct
// [PurgeStartResponseEnvelope]
type purgeStartResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeStartResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r purgeStartResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type PurgeStartResponseEnvelopeSuccess bool

const (
	PurgeStartResponseEnvelopeSuccessTrue PurgeStartResponseEnvelopeSuccess = true
)

func (r PurgeStartResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PurgeStartResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PurgeStatusParams struct {
	// A Resource identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PurgeStatusResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   PurgeStatusResponse   `json:"result"`
	// Indicates if the API call was successful or not.
	Success PurgeStatusResponseEnvelopeSuccess `json:"success"`
	JSON    purgeStatusResponseEnvelopeJSON    `json:"-"`
}

// purgeStatusResponseEnvelopeJSON contains the JSON metadata for the struct
// [PurgeStatusResponseEnvelope]
type purgeStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r purgeStatusResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type PurgeStatusResponseEnvelopeSuccess bool

const (
	PurgeStatusResponseEnvelopeSuccessTrue PurgeStatusResponseEnvelopeSuccess = true
)

func (r PurgeStatusResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PurgeStatusResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
