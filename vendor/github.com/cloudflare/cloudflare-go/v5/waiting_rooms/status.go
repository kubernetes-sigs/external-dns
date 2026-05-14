// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

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

// StatusService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	Options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r *StatusService) {
	r = &StatusService{}
	r.Options = opts
	return
}

// Fetches the status of a configured waiting room. Response fields include:
//
//  1. `status`: String indicating the status of the waiting room. The possible
//     status are:
//     - **not_queueing** indicates that the configured thresholds have not been met
//     and all users are going through to the origin.
//     - **queueing** indicates that the thresholds have been met and some users are
//     held in the waiting room.
//     - **event_prequeueing** indicates that an event is active and is currently
//     prequeueing users before it starts.
//     - **suspended** indicates that the room is suspended.
//  2. `event_id`: String of the current event's `id` if an event is active,
//     otherwise an empty string.
//  3. `estimated_queued_users`: Integer of the estimated number of users currently
//     waiting in the queue.
//  4. `estimated_total_active_users`: Integer of the estimated number of users
//     currently active on the origin.
//  5. `max_estimated_time_minutes`: Integer of the maximum estimated time currently
//     presented to the users.
func (r *StatusService) Get(ctx context.Context, waitingRoomID string, query StatusGetParams, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	var env StatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/status", query.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StatusGetResponse struct {
	EstimatedQueuedUsers      int64                   `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int64                   `json:"estimated_total_active_users"`
	EventID                   string                  `json:"event_id"`
	MaxEstimatedTimeMinutes   int64                   `json:"max_estimated_time_minutes"`
	Status                    StatusGetResponseStatus `json:"status"`
	JSON                      statusGetResponseJSON   `json:"-"`
}

// statusGetResponseJSON contains the JSON metadata for the struct
// [StatusGetResponse]
type statusGetResponseJSON struct {
	EstimatedQueuedUsers      apijson.Field
	EstimatedTotalActiveUsers apijson.Field
	EventID                   apijson.Field
	MaxEstimatedTimeMinutes   apijson.Field
	Status                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *StatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusGetResponseJSON) RawJSON() string {
	return r.raw
}

type StatusGetResponseStatus string

const (
	StatusGetResponseStatusEventPrequeueing StatusGetResponseStatus = "event_prequeueing"
	StatusGetResponseStatusNotQueueing      StatusGetResponseStatus = "not_queueing"
	StatusGetResponseStatusQueueing         StatusGetResponseStatus = "queueing"
	StatusGetResponseStatusSuspended        StatusGetResponseStatus = "suspended"
)

func (r StatusGetResponseStatus) IsKnown() bool {
	switch r {
	case StatusGetResponseStatusEventPrequeueing, StatusGetResponseStatusNotQueueing, StatusGetResponseStatusQueueing, StatusGetResponseStatusSuspended:
		return true
	}
	return false
}

type StatusGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type StatusGetResponseEnvelope struct {
	Result StatusGetResponse             `json:"result,required"`
	JSON   statusGetResponseEnvelopeJSON `json:"-"`
}

// statusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StatusGetResponseEnvelope]
type statusGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
