// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

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

// EventDetailService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventDetailService] method instead.
type EventDetailService struct {
	Options []option.RequestOption
}

// NewEventDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventDetailService(opts ...option.RequestOption) (r *EventDetailService) {
	r = &EventDetailService{}
	r.Options = opts
	return
}

// Previews an event's configuration as if it was active. Inherited fields from the
// waiting room will be displayed with their current values.
func (r *EventDetailService) Get(ctx context.Context, waitingRoomID string, eventID string, query EventDetailGetParams, opts ...option.RequestOption) (res *EventDetailGetResponse, err error) {
	var env EventDetailGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s/details", query.ZoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EventQueryParam struct {
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime param.Field[string] `json:"event_end_time,required"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime param.Field[string] `json:"event_start_time,required"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description param.Field[string] `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime param.Field[string] `json:"prequeue_start_time"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod param.Field[string] `json:"queueing_method"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart param.Field[bool] `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended param.Field[bool] `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers param.Field[int64] `json:"total_active_users"`
	// If set, the event will override the waiting room's `turnstile_action` property
	// while it is active. If null, the event will inherit it.
	TurnstileAction param.Field[EventQueryTurnstileAction] `json:"turnstile_action"`
	// If set, the event will override the waiting room's `turnstile_mode` property
	// while it is active. If null, the event will inherit it.
	TurnstileMode param.Field[EventQueryTurnstileMode] `json:"turnstile_mode"`
}

func (r EventQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the event will override the waiting room's `turnstile_action` property
// while it is active. If null, the event will inherit it.
type EventQueryTurnstileAction string

const (
	EventQueryTurnstileActionLog           EventQueryTurnstileAction = "log"
	EventQueryTurnstileActionInfiniteQueue EventQueryTurnstileAction = "infinite_queue"
)

func (r EventQueryTurnstileAction) IsKnown() bool {
	switch r {
	case EventQueryTurnstileActionLog, EventQueryTurnstileActionInfiniteQueue:
		return true
	}
	return false
}

// If set, the event will override the waiting room's `turnstile_mode` property
// while it is active. If null, the event will inherit it.
type EventQueryTurnstileMode string

const (
	EventQueryTurnstileModeOff                   EventQueryTurnstileMode = "off"
	EventQueryTurnstileModeInvisible             EventQueryTurnstileMode = "invisible"
	EventQueryTurnstileModeVisibleNonInteractive EventQueryTurnstileMode = "visible_non_interactive"
	EventQueryTurnstileModeVisibleManaged        EventQueryTurnstileMode = "visible_managed"
)

func (r EventQueryTurnstileMode) IsKnown() bool {
	switch r {
	case EventQueryTurnstileModeOff, EventQueryTurnstileModeInvisible, EventQueryTurnstileModeVisibleNonInteractive, EventQueryTurnstileModeVisibleManaged:
		return true
	}
	return false
}

type EventDetailGetResponse struct {
	ID             string    `json:"id"`
	CreatedOn      time.Time `json:"created_on" format:"date-time"`
	CustomPageHTML string    `json:"custom_page_html"`
	// A note that you can use to add more details about the event.
	Description           string `json:"description"`
	DisableSessionRenewal bool   `json:"disable_session_renewal"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name              string `json:"name"`
	NewUsersPerMinute int64  `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	QueueingMethod    string `json:"queueing_method"`
	SessionDuration   int64  `json:"session_duration"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended        bool                       `json:"suspended"`
	TotalActiveUsers int64                      `json:"total_active_users"`
	JSON             eventDetailGetResponseJSON `json:"-"`
}

// eventDetailGetResponseJSON contains the JSON metadata for the struct
// [EventDetailGetResponse]
type eventDetailGetResponseJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	CustomPageHTML        apijson.Field
	Description           apijson.Field
	DisableSessionRenewal apijson.Field
	EventEndTime          apijson.Field
	EventStartTime        apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NewUsersPerMinute     apijson.Field
	PrequeueStartTime     apijson.Field
	QueueingMethod        apijson.Field
	SessionDuration       apijson.Field
	ShuffleAtEventStart   apijson.Field
	Suspended             apijson.Field
	TotalActiveUsers      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EventDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type EventDetailGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type EventDetailGetResponseEnvelope struct {
	Result EventDetailGetResponse             `json:"result,required"`
	JSON   eventDetailGetResponseEnvelopeJSON `json:"-"`
}

// eventDetailGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventDetailGetResponseEnvelope]
type eventDetailGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDetailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDetailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
