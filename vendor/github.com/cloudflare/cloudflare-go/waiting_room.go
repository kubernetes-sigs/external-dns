package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WaitingRoom describes a WaitingRoom object.
type WaitingRoom struct {
	CreatedOn                  time.Time  `json:"created_on,omitempty"`
	ModifiedOn                 time.Time  `json:"modified_on,omitempty"`
	Path                       string     `json:"path"`
	Name                       string     `json:"name"`
	Description                string     `json:"description,omitempty"`
	QueueingMethod             string     `json:"queueing_method,omitempty"`
	CustomPageHTML             string     `json:"custom_page_html,omitempty"`
	DefaultTemplateLanguage    string     `json:"default_template_language,omitempty"`
	Host                       string     `json:"host"`
	ID                         string     `json:"id,omitempty"`
	NewUsersPerMinute          int        `json:"new_users_per_minute"`
	TotalActiveUsers           int        `json:"total_active_users"`
	SessionDuration            int        `json:"session_duration"`
	QueueAll                   bool       `json:"queue_all"`
	DisableSessionRenewal      bool       `json:"disable_session_renewal"`
	Suspended                  bool       `json:"suspended"`
	JsonResponseEnabled        bool       `json:"json_response_enabled"`
	NextEventPrequeueStartTime *time.Time `json:"next_event_prequeue_start_time,omitempty"`
	NextEventStartTime         *time.Time `json:"next_event_start_time,omitempty"`
}

// WaitingRoomStatus describes the status of a waiting room.
type WaitingRoomStatus struct {
	Status                    string `json:"status"`
	EventID                   string `json:"event_id"`
	EstimatedQueuedUsers      int    `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int    `json:"estimated_total_active_users"`
	MaxEstimatedTimeMinutes   int    `json:"max_estimated_time_minutes"`
}

// WaitingRoomEvent describes a WaitingRoomEvent object.
type WaitingRoomEvent struct {
	EventEndTime          time.Time  `json:"event_end_time"`
	CreatedOn             time.Time  `json:"created_on,omitempty"`
	ModifiedOn            time.Time  `json:"modified_on,omitempty"`
	PrequeueStartTime     *time.Time `json:"prequeue_start_time,omitempty"`
	EventStartTime        time.Time  `json:"event_start_time"`
	Name                  string     `json:"name"`
	Description           string     `json:"description,omitempty"`
	QueueingMethod        string     `json:"queueing_method,omitempty"`
	ID                    string     `json:"id,omitempty"`
	CustomPageHTML        string     `json:"custom_page_html,omitempty"`
	NewUsersPerMinute     int        `json:"new_users_per_minute,omitempty"`
	TotalActiveUsers      int        `json:"total_active_users,omitempty"`
	SessionDuration       int        `json:"session_duration,omitempty"`
	DisableSessionRenewal *bool      `json:"disable_session_renewal,omitempty"`
	Suspended             bool       `json:"suspended"`
	ShuffleAtEventStart   bool       `json:"shuffle_at_event_start"`
}

// WaitingRoomPagePreviewURL describes a WaitingRoomPagePreviewURL object.
type WaitingRoomPagePreviewURL struct {
	PreviewURL string `json:"preview_url"`
}

// WaitingRoomPagePreviewCustomHTML describes a WaitingRoomPagePreviewCustomHTML object.
type WaitingRoomPagePreviewCustomHTML struct {
	CustomHTML string `json:"custom_html"`
}

// WaitingRoomDetailResponse is the API response, containing a single WaitingRoom.
type WaitingRoomDetailResponse struct {
	Response
	Result WaitingRoom `json:"result"`
}

// WaitingRoomsResponse is the API response, containing an array of WaitingRooms.
type WaitingRoomsResponse struct {
	Response
	Result []WaitingRoom `json:"result"`
}

// WaitingRoomStatusResponse is the API response, containing the status of a waiting room.
type WaitingRoomStatusResponse struct {
	Response
	Result WaitingRoomStatus `json:"result"`
}

// WaitingRoomPagePreviewResponse is the API response, containing the URL to a custom waiting room preview.
type WaitingRoomPagePreviewResponse struct {
	Response
	Result WaitingRoomPagePreviewURL `json:"result"`
}

// WaitingRoomEventDetailResponse is the API response, containing a single WaitingRoomEvent.
type WaitingRoomEventDetailResponse struct {
	Response
	Result WaitingRoomEvent `json:"result"`
}

// WaitingRoomEventsResponse is the API response, containing an array of WaitingRoomEvents.
type WaitingRoomEventsResponse struct {
	Response
	Result []WaitingRoomEvent `json:"result"`
}

// CreateWaitingRoom creates a new Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-waiting-room
func (api *API) CreateWaitingRoom(ctx context.Context, zoneID string, waitingRoom WaitingRoom) (*WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, waitingRoom)
	if err != nil {
		return nil, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

// ListWaitingRooms returns all Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-list-waiting-rooms
func (api *API) ListWaitingRooms(ctx context.Context, zoneID string) ([]WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WaitingRoom{}, err
	}
	var r WaitingRoomsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoom fetches detail about one Waiting room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-waiting-room-details
func (api *API) WaitingRoom(ctx context.Context, zoneID, waitingRoomID string) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ChangeWaitingRoom lets you change individual settings for a Waiting room. This is
// in contrast to UpdateWaitingRoom which replaces the entire Waiting room.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-waiting-room
func (api *API) ChangeWaitingRoom(ctx context.Context, zoneID, waitingRoomID string, waitingRoom WaitingRoom) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, waitingRoom)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateWaitingRoom lets you replace a Waiting Room. This is in contrast to
// ChangeWaitingRoom which lets you change individual settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-waiting-room
func (api *API) UpdateWaitingRoom(ctx context.Context, zoneID string, waitingRoom WaitingRoom) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoom.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, waitingRoom)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteWaitingRoom deletes a Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-delete-waiting-room
func (api *API) DeleteWaitingRoom(ctx context.Context, zoneID, waitingRoomID string) error {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

// WaitingRoomStatus returns the status of one Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-get-waiting-room-status
func (api *API) WaitingRoomStatus(ctx context.Context, zoneID, waitingRoomID string) (WaitingRoomStatus, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/status", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomStatus{}, err
	}
	var r WaitingRoomStatusResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomStatus{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomPagePreview uploads a custom waiting room page for preview and
// returns a preview URL.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-a-custom-waiting-room-page-preview
func (api *API) WaitingRoomPagePreview(ctx context.Context, zoneID, customHTML string) (WaitingRoomPagePreviewURL, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/preview", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, WaitingRoomPagePreviewCustomHTML{CustomHTML: customHTML})

	if err != nil {
		return WaitingRoomPagePreviewURL{}, err
	}
	var r WaitingRoomPagePreviewResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomPagePreviewURL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateWaitingRoomEvent creates a new event for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-event
func (api *API) CreateWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (*WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, waitingRoomEvent)
	if err != nil {
		return nil, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

// ListWaitingRoomEvents returns all Waiting Room Events for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-list-events
func (api *API) ListWaitingRoomEvents(ctx context.Context, zoneID string, waitingRoomID string) ([]WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomEvent fetches detail about one Waiting Room Event for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-event-details
func (api *API) WaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, eventID string) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomEventPreview returns an event's configuration as if it was active.
// Inherited fields from the waiting room will be displayed with their current values.
//
// API reference: https://api.cloudflare.com/#waiting-room-preview-active-event-details
func (api *API) WaitingRoomEventPreview(ctx context.Context, zoneID string, waitingRoomID string, eventID string) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s/details", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ChangeWaitingRoomEvent lets you change individual settings for a Waiting Room Event. This is
// in contrast to UpdateWaitingRoomEvent which replaces the entire Waiting Room Event.
//
// API reference: https://api.cloudflare.com/#waiting-room-patch-event
func (api *API) ChangeWaitingRoomEvent(ctx context.Context, zoneID, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, waitingRoomEvent.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, waitingRoomEvent)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateWaitingRoomEvent lets you replace a Waiting Room Event. This is in contrast to
// ChangeWaitingRoomEvent which lets you change individual settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-event
func (api *API) UpdateWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, waitingRoomEvent.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, waitingRoomEvent)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteWaitingRoomEvent deletes an event for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-delete-event
func (api *API) DeleteWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, eventID string) error {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingWaitingRoomID     = errors.New("missing required waiting room ID")
	ErrMissingWaitingRoomRuleID = errors.New("missing required waiting room rule ID")
)

// WaitingRoom describes a WaitingRoom object.
type WaitingRoom struct {
	CreatedOn                  time.Time           `json:"created_on,omitempty"`
	ModifiedOn                 time.Time           `json:"modified_on,omitempty"`
	Path                       string              `json:"path"`
	Name                       string              `json:"name"`
	Description                string              `json:"description,omitempty"`
	QueueingMethod             string              `json:"queueing_method,omitempty"`
	CustomPageHTML             string              `json:"custom_page_html,omitempty"`
	DefaultTemplateLanguage    string              `json:"default_template_language,omitempty"`
	Host                       string              `json:"host"`
	ID                         string              `json:"id,omitempty"`
	NewUsersPerMinute          int                 `json:"new_users_per_minute"`
	TotalActiveUsers           int                 `json:"total_active_users"`
	SessionDuration            int                 `json:"session_duration"`
	QueueAll                   bool                `json:"queue_all"`
	DisableSessionRenewal      bool                `json:"disable_session_renewal"`
	Suspended                  bool                `json:"suspended"`
	JsonResponseEnabled        bool                `json:"json_response_enabled"`
	NextEventPrequeueStartTime *time.Time          `json:"next_event_prequeue_start_time,omitempty"`
	NextEventStartTime         *time.Time          `json:"next_event_start_time,omitempty"`
	CookieSuffix               string              `json:"cookie_suffix"`
	AdditionalRoutes           []*WaitingRoomRoute `json:"additional_routes,omitempty"`
	QueueingStatusCode         int                 `json:"queueing_status_code"`
}

// WaitingRoomStatus describes the status of a waiting room.
type WaitingRoomStatus struct {
	Status                    string `json:"status"`
	EventID                   string `json:"event_id"`
	EstimatedQueuedUsers      int    `json:"estimated_queued_users"`
	EstimatedTotalActiveUsers int    `json:"estimated_total_active_users"`
	MaxEstimatedTimeMinutes   int    `json:"max_estimated_time_minutes"`
}

// WaitingRoomEvent describes a WaitingRoomEvent object.
type WaitingRoomEvent struct {
	EventEndTime          time.Time  `json:"event_end_time"`
	CreatedOn             time.Time  `json:"created_on,omitempty"`
	ModifiedOn            time.Time  `json:"modified_on,omitempty"`
	PrequeueStartTime     *time.Time `json:"prequeue_start_time,omitempty"`
	EventStartTime        time.Time  `json:"event_start_time"`
	Name                  string     `json:"name"`
	Description           string     `json:"description,omitempty"`
	QueueingMethod        string     `json:"queueing_method,omitempty"`
	ID                    string     `json:"id,omitempty"`
	CustomPageHTML        string     `json:"custom_page_html,omitempty"`
	NewUsersPerMinute     int        `json:"new_users_per_minute,omitempty"`
	TotalActiveUsers      int        `json:"total_active_users,omitempty"`
	SessionDuration       int        `json:"session_duration,omitempty"`
	DisableSessionRenewal *bool      `json:"disable_session_renewal,omitempty"`
	Suspended             bool       `json:"suspended"`
	ShuffleAtEventStart   bool       `json:"shuffle_at_event_start"`
}

type WaitingRoomRule struct {
	ID          string     `json:"id,omitempty"`
	Version     string     `json:"version,omitempty"`
	Action      string     `json:"action"`
	Expression  string     `json:"expression"`
	Description string     `json:"description"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Enabled     *bool      `json:"enabled"`
}

// WaitingRoomSettings describes zone-level waiting room settings.
type WaitingRoomSettings struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on this zone
	SearchEngineCrawlerBypass bool `json:"search_engine_crawler_bypass"`
}

// WaitingRoomPagePreviewURL describes a WaitingRoomPagePreviewURL object.
type WaitingRoomPagePreviewURL struct {
	PreviewURL string `json:"preview_url"`
}

// WaitingRoomPagePreviewCustomHTML describes a WaitingRoomPagePreviewCustomHTML object.
type WaitingRoomPagePreviewCustomHTML struct {
	CustomHTML string `json:"custom_html"`
}

// WaitingRoomRoute describes a WaitingRoomRoute object.
type WaitingRoomRoute struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

// WaitingRoomDetailResponse is the API response, containing a single WaitingRoom.
type WaitingRoomDetailResponse struct {
	Response
	Result WaitingRoom `json:"result"`
}

// WaitingRoomsResponse is the API response, containing an array of WaitingRooms.
type WaitingRoomsResponse struct {
	Response
	Result []WaitingRoom `json:"result"`
}

// WaitingRoomSettingsResponse is the API response, containing zone-level Waiting Room settings.
type WaitingRoomSettingsResponse struct {
	Response
	Result WaitingRoomSettings `json:"result"`
}

// WaitingRoomStatusResponse is the API response, containing the status of a waiting room.
type WaitingRoomStatusResponse struct {
	Response
	Result WaitingRoomStatus `json:"result"`
}

// WaitingRoomPagePreviewResponse is the API response, containing the URL to a custom waiting room preview.
type WaitingRoomPagePreviewResponse struct {
	Response
	Result WaitingRoomPagePreviewURL `json:"result"`
}

// WaitingRoomEventDetailResponse is the API response, containing a single WaitingRoomEvent.
type WaitingRoomEventDetailResponse struct {
	Response
	Result WaitingRoomEvent `json:"result"`
}

// WaitingRoomEventsResponse is the API response, containing an array of WaitingRoomEvents.
type WaitingRoomEventsResponse struct {
	Response
	Result []WaitingRoomEvent `json:"result"`
}

// WaitingRoomRulesResponse is the API response, containing an array of WaitingRoomRule.
type WaitingRoomRulesResponse struct {
	Response
	Result []WaitingRoomRule `json:"result"`
}

// CreateWaitingRoom creates a new Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-waiting-room
func (api *API) CreateWaitingRoom(ctx context.Context, zoneID string, waitingRoom WaitingRoom) (*WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, waitingRoom)
	if err != nil {
		return nil, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

// ListWaitingRooms returns all Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-list-waiting-rooms
func (api *API) ListWaitingRooms(ctx context.Context, zoneID string) ([]WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WaitingRoom{}, err
	}
	var r WaitingRoomsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoom fetches detail about one Waiting room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-waiting-room-details
func (api *API) WaitingRoom(ctx context.Context, zoneID, waitingRoomID string) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ChangeWaitingRoom lets you change individual settings for a Waiting room. This is
// in contrast to UpdateWaitingRoom which replaces the entire Waiting room.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-waiting-room
func (api *API) ChangeWaitingRoom(ctx context.Context, zoneID, waitingRoomID string, waitingRoom WaitingRoom) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, waitingRoom)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateWaitingRoom lets you replace a Waiting Room. This is in contrast to
// ChangeWaitingRoom which lets you change individual settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-waiting-room
func (api *API) UpdateWaitingRoom(ctx context.Context, zoneID string, waitingRoom WaitingRoom) (WaitingRoom, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoom.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, waitingRoom)
	if err != nil {
		return WaitingRoom{}, err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoom{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteWaitingRoom deletes a Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-delete-waiting-room
func (api *API) DeleteWaitingRoom(ctx context.Context, zoneID, waitingRoomID string) error {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r WaitingRoomDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

// WaitingRoomStatus returns the status of one Waiting Room for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-get-waiting-room-status
func (api *API) WaitingRoomStatus(ctx context.Context, zoneID, waitingRoomID string) (WaitingRoomStatus, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/status", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomStatus{}, err
	}
	var r WaitingRoomStatusResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomStatus{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomPagePreview uploads a custom waiting room page for preview and
// returns a preview URL.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-a-custom-waiting-room-page-preview
func (api *API) WaitingRoomPagePreview(ctx context.Context, zoneID, customHTML string) (WaitingRoomPagePreviewURL, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/preview", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, WaitingRoomPagePreviewCustomHTML{CustomHTML: customHTML})

	if err != nil {
		return WaitingRoomPagePreviewURL{}, err
	}
	var r WaitingRoomPagePreviewResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomPagePreviewURL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateWaitingRoomEvent creates a new event for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-event
func (api *API) CreateWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (*WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, waitingRoomEvent)
	if err != nil {
		return nil, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

// ListWaitingRoomEvents returns all Waiting Room Events for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-list-events
func (api *API) ListWaitingRoomEvents(ctx context.Context, zoneID string, waitingRoomID string) ([]WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events", zoneID, waitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomEvent fetches detail about one Waiting Room Event for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-event-details
func (api *API) WaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, eventID string) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// WaitingRoomEventPreview returns an event's configuration as if it was active.
// Inherited fields from the waiting room will be displayed with their current values.
//
// API reference: https://api.cloudflare.com/#waiting-room-preview-active-event-details
func (api *API) WaitingRoomEventPreview(ctx context.Context, zoneID string, waitingRoomID string, eventID string) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s/details", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ChangeWaitingRoomEvent lets you change individual settings for a Waiting Room Event. This is
// in contrast to UpdateWaitingRoomEvent which replaces the entire Waiting Room Event.
//
// API reference: https://api.cloudflare.com/#waiting-room-patch-event
func (api *API) ChangeWaitingRoomEvent(ctx context.Context, zoneID, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, waitingRoomEvent.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, waitingRoomEvent)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateWaitingRoomEvent lets you replace a Waiting Room Event. This is in contrast to
// ChangeWaitingRoomEvent which lets you change individual settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-event
func (api *API) UpdateWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, waitingRoomEvent WaitingRoomEvent) (WaitingRoomEvent, error) {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, waitingRoomEvent.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, waitingRoomEvent)
	if err != nil {
		return WaitingRoomEvent{}, err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomEvent{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteWaitingRoomEvent deletes an event for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-delete-event
func (api *API) DeleteWaitingRoomEvent(ctx context.Context, zoneID string, waitingRoomID string, eventID string) error {
	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/events/%s", zoneID, waitingRoomID, eventID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r WaitingRoomEventDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

type ListWaitingRoomRuleParams struct {
	WaitingRoomID string
}

type CreateWaitingRoomRuleParams struct {
	WaitingRoomID string
	Rule          WaitingRoomRule
}

type ReplaceWaitingRoomRuleParams struct {
	WaitingRoomID string
	Rules         []WaitingRoomRule
}

type UpdateWaitingRoomRuleParams struct {
	WaitingRoomID string
	Rule          WaitingRoomRule
}

type DeleteWaitingRoomRuleParams struct {
	WaitingRoomID string
	RuleID        string
}

// ListWaitingRoomRules lists all rules for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-list-waiting-room-rules
func (api *API) ListWaitingRoomRules(ctx context.Context, rc *ResourceContainer, params ListWaitingRoomRuleParams) ([]WaitingRoomRule, error) {
	if params.WaitingRoomID == "" {
		return nil, ErrMissingWaitingRoomID
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/rules", rc.Identifier, params.WaitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var r WaitingRoomRulesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// CreateWaitingRoomRule creates a new rule for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-create-waiting-room-rule
func (api *API) CreateWaitingRoomRule(ctx context.Context, rc *ResourceContainer, params CreateWaitingRoomRuleParams) ([]WaitingRoomRule, error) {
	if params.WaitingRoomID == "" {
		return nil, ErrMissingWaitingRoomID
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/rules", rc.Identifier, params.WaitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Rule)
	if err != nil {
		return nil, err
	}

	var r WaitingRoomRulesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// ReplaceWaitingRoomRules replaces all rules for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-replace-waiting-room-rules
func (api *API) ReplaceWaitingRoomRules(ctx context.Context, rc *ResourceContainer, params ReplaceWaitingRoomRuleParams) ([]WaitingRoomRule, error) {
	if params.WaitingRoomID == "" {
		return nil, ErrMissingWaitingRoomID
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/rules", rc.Identifier, params.WaitingRoomID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Rules)
	if err != nil {
		return nil, err
	}

	var r WaitingRoomRulesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateWaitingRoomRule updates a rule for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-patch-waiting-room-rule
func (api *API) UpdateWaitingRoomRule(ctx context.Context, rc *ResourceContainer, params UpdateWaitingRoomRuleParams) ([]WaitingRoomRule, error) {
	if params.WaitingRoomID == "" {
		return nil, ErrMissingWaitingRoomID
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/rules/%s", rc.Identifier, params.WaitingRoomID, params.Rule.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params.Rule)
	if err != nil {
		return nil, err
	}

	var r WaitingRoomRulesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteWaitingRoomRule deletes a rule for a Waiting Room.
//
// API reference: https://api.cloudflare.com/#waiting-room-delete-waiting-room-rule
func (api *API) DeleteWaitingRoomRule(ctx context.Context, rc *ResourceContainer, params DeleteWaitingRoomRuleParams) ([]WaitingRoomRule, error) {
	if params.WaitingRoomID == "" {
		return nil, ErrMissingWaitingRoomID
	}

	if params.RuleID == "" {
		return nil, ErrMissingWaitingRoomRuleID
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/%s/rules/%s", rc.Identifier, params.WaitingRoomID, params.RuleID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	var r WaitingRoomRulesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// GetWaitingRoomSettings fetches the Waiting Room zone-level settings for a zone.
//
// API reference: https://api.cloudflare.com/#waiting-room-get-zone-settings
func (api *API) GetWaitingRoomSettings(ctx context.Context, rc *ResourceContainer) (WaitingRoomSettings, error) {
	if rc.Level != ZoneRouteLevel {
		return WaitingRoomSettings{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WaitingRoomSettings{}, err
	}
	var r WaitingRoomSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

type PatchWaitingRoomSettingsParams struct {
	SearchEngineCrawlerBypass *bool `json:"search_engine_crawler_bypass,omitempty"`
}

// PatchWaitingRoomSettings lets you change individual zone-level Waiting Room settings. This is
// in contrast to UpdateWaitingRoomSettings which replaces all settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-patch-zone-settings
func (api *API) PatchWaitingRoomSettings(ctx context.Context, rc *ResourceContainer, params PatchWaitingRoomSettingsParams) (WaitingRoomSettings, error) {
	if rc.Level != ZoneRouteLevel {
		return WaitingRoomSettings{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return WaitingRoomSettings{}, err
	}
	var r WaitingRoomSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

type UpdateWaitingRoomSettingsParams struct {
	SearchEngineCrawlerBypass *bool `json:"search_engine_crawler_bypass,omitempty"`
}

// UpdateWaitingRoomSettings lets you replace all zone-level Waiting Room settings. This is in contrast to
// PatchWaitingRoomSettings which lets you change individual settings.
//
// API reference: https://api.cloudflare.com/#waiting-room-update-zone-settings
func (api *API) UpdateWaitingRoomSettings(ctx context.Context, rc *ResourceContainer, params UpdateWaitingRoomSettingsParams) (WaitingRoomSettings, error) {
	if rc.Level != ZoneRouteLevel {
		return WaitingRoomSettings{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf("/zones/%s/waiting_rooms/settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WaitingRoomSettings{}, err
	}
	var r WaitingRoomSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WaitingRoomSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
