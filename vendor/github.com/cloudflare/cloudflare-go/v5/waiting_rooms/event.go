// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// EventService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
	Details *EventDetailService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Details = NewEventDetailService(opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Creates an event for
// a waiting room. An event takes place during a specified period of time,
// temporarily changing the behavior of a waiting room. While the event is active,
// some of the properties in the event's configuration may either override or
// inherit from the waiting room's configuration. Note that events cannot overlap
// with each other, so only one event can be active at a time.
func (r *EventService) New(ctx context.Context, waitingRoomID string, params EventNewParams, opts ...option.RequestOption) (res *Event, err error) {
	var env EventNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events", params.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured event for a waiting room.
func (r *EventService) Update(ctx context.Context, waitingRoomID string, eventID string, params EventUpdateParams, opts ...option.RequestOption) (res *Event, err error) {
	var env EventUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
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
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", params.ZoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists events for a waiting room.
func (r *EventService) List(ctx context.Context, waitingRoomID string, params EventListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Event], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events", params.ZoneID, waitingRoomID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists events for a waiting room.
func (r *EventService) ListAutoPaging(ctx context.Context, waitingRoomID string, params EventListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Event] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, waitingRoomID, params, opts...))
}

// Deletes an event for a waiting room.
func (r *EventService) Delete(ctx context.Context, waitingRoomID string, eventID string, body EventDeleteParams, opts ...option.RequestOption) (res *EventDeleteResponse, err error) {
	var env EventDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
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
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", body.ZoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches a configured event for a waiting room.
func (r *EventService) Edit(ctx context.Context, waitingRoomID string, eventID string, params EventEditParams, opts ...option.RequestOption) (res *Event, err error) {
	var env EventEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
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
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", params.ZoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single configured event for a waiting room.
func (r *EventService) Get(ctx context.Context, waitingRoomID string, eventID string, query EventGetParams, opts ...option.RequestOption) (res *Event, err error) {
	var env EventGetResponseEnvelope
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
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/events/%s", query.ZoneID, waitingRoomID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Event struct {
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// If set, the event will override the waiting room's `custom_page_html` property
	// while it is active. If null, the event will inherit it.
	CustomPageHTML string `json:"custom_page_html,nullable"`
	// A note that you can use to add more details about the event.
	Description string `json:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal`
	// property while it is active. If null, the event will inherit it.
	DisableSessionRenewal bool `json:"disable_session_renewal,nullable"`
	// An ISO 8601 timestamp that marks the end of the event.
	EventEndTime string `json:"event_end_time"`
	// An ISO 8601 timestamp that marks the start of the event. At this time, queued
	// users will be processed with the event's configuration. The start time must be
	// at least one minute before `event_end_time`.
	EventStartTime string    `json:"event_start_time"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and
	// underscores are allowed.
	Name string `json:"name"`
	// If set, the event will override the waiting room's `new_users_per_minute`
	// property while it is active. If null, the event will inherit it. This can only
	// be set if the event's `total_active_users` property is also set.
	NewUsersPerMinute int64 `json:"new_users_per_minute,nullable"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the
	// event starts. The prequeue must start at least five minutes before
	// `event_start_time`.
	PrequeueStartTime string `json:"prequeue_start_time,nullable"`
	// If set, the event will override the waiting room's `queueing_method` property
	// while it is active. If null, the event will inherit it.
	QueueingMethod string `json:"queueing_method,nullable"`
	// If set, the event will override the waiting room's `session_duration` property
	// while it is active. If null, the event will inherit it.
	SessionDuration int64 `json:"session_duration,nullable"`
	// If enabled, users in the prequeue will be shuffled randomly at the
	// `event_start_time`. Requires that `prequeue_start_time` is not null. This is
	// useful for situations when many users will join the event prequeue at the same
	// time and you want to shuffle them to ensure fairness. Naturally, it makes the
	// most sense to enable this feature when the `queueing_method` during the event
	// respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	ShuffleAtEventStart bool `json:"shuffle_at_event_start"`
	// Suspends or allows an event. If set to `true`, the event is ignored and traffic
	// will be handled based on the waiting room configuration.
	Suspended bool `json:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property
	// while it is active. If null, the event will inherit it. This can only be set if
	// the event's `new_users_per_minute` property is also set.
	TotalActiveUsers int64 `json:"total_active_users,nullable"`
	// If set, the event will override the waiting room's `turnstile_action` property
	// while it is active. If null, the event will inherit it.
	TurnstileAction EventTurnstileAction `json:"turnstile_action,nullable"`
	// If set, the event will override the waiting room's `turnstile_mode` property
	// while it is active. If null, the event will inherit it.
	TurnstileMode EventTurnstileMode `json:"turnstile_mode,nullable"`
	JSON          eventJSON          `json:"-"`
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
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
	TurnstileAction       apijson.Field
	TurnstileMode         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventJSON) RawJSON() string {
	return r.raw
}

// If set, the event will override the waiting room's `turnstile_action` property
// while it is active. If null, the event will inherit it.
type EventTurnstileAction string

const (
	EventTurnstileActionLog           EventTurnstileAction = "log"
	EventTurnstileActionInfiniteQueue EventTurnstileAction = "infinite_queue"
)

func (r EventTurnstileAction) IsKnown() bool {
	switch r {
	case EventTurnstileActionLog, EventTurnstileActionInfiniteQueue:
		return true
	}
	return false
}

// If set, the event will override the waiting room's `turnstile_mode` property
// while it is active. If null, the event will inherit it.
type EventTurnstileMode string

const (
	EventTurnstileModeOff                   EventTurnstileMode = "off"
	EventTurnstileModeInvisible             EventTurnstileMode = "invisible"
	EventTurnstileModeVisibleNonInteractive EventTurnstileMode = "visible_non_interactive"
	EventTurnstileModeVisibleManaged        EventTurnstileMode = "visible_managed"
)

func (r EventTurnstileMode) IsKnown() bool {
	switch r {
	case EventTurnstileModeOff, EventTurnstileModeInvisible, EventTurnstileModeVisibleNonInteractive, EventTurnstileModeVisibleManaged:
		return true
	}
	return false
}

type EventDeleteResponse struct {
	ID   string                  `json:"id"`
	JSON eventDeleteResponseJSON `json:"-"`
}

// eventDeleteResponseJSON contains the JSON metadata for the struct
// [EventDeleteResponse]
type eventDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type EventNewParams struct {
	// Identifier.
	ZoneID     param.Field[string] `path:"zone_id,required"`
	EventQuery EventQueryParam     `json:"event_query,required"`
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EventQuery)
}

type EventNewResponseEnvelope struct {
	Result Event                        `json:"result,required"`
	JSON   eventNewResponseEnvelopeJSON `json:"-"`
}

// eventNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventNewResponseEnvelope]
type eventNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EventUpdateParams struct {
	// Identifier.
	ZoneID     param.Field[string] `path:"zone_id,required"`
	EventQuery EventQueryParam     `json:"event_query,required"`
}

func (r EventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EventQuery)
}

type EventUpdateResponseEnvelope struct {
	Result Event                           `json:"result,required"`
	JSON   eventUpdateResponseEnvelopeJSON `json:"-"`
}

// eventUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventUpdateResponseEnvelope]
type eventUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EventListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page. Must be a multiple of 5.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EventDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type EventDeleteResponseEnvelope struct {
	Result EventDeleteResponse             `json:"result,required"`
	JSON   eventDeleteResponseEnvelopeJSON `json:"-"`
}

// eventDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventDeleteResponseEnvelope]
type eventDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EventEditParams struct {
	// Identifier.
	ZoneID     param.Field[string] `path:"zone_id,required"`
	EventQuery EventQueryParam     `json:"event_query,required"`
}

func (r EventEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EventQuery)
}

type EventEditResponseEnvelope struct {
	Result Event                         `json:"result,required"`
	JSON   eventEditResponseEnvelopeJSON `json:"-"`
}

// eventEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventEditResponseEnvelope]
type eventEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EventGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type EventGetResponseEnvelope struct {
	Result Event                        `json:"result,required"`
	JSON   eventGetResponseEnvelopeJSON `json:"-"`
}

// eventGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EventGetResponseEnvelope]
type eventGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
