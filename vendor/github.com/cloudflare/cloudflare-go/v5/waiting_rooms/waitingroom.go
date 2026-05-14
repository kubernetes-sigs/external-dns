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

// WaitingRoomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaitingRoomService] method instead.
type WaitingRoomService struct {
	Options  []option.RequestOption
	Page     *PageService
	Events   *EventService
	Rules    *RuleService
	Statuses *StatusService
	Settings *SettingService
}

// NewWaitingRoomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaitingRoomService(opts ...option.RequestOption) (r *WaitingRoomService) {
	r = &WaitingRoomService{}
	r.Options = opts
	r.Page = NewPageService(opts...)
	r.Events = NewEventService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Statuses = NewStatusService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Creates a new waiting room.
func (r *WaitingRoomService) New(ctx context.Context, params WaitingRoomNewParams, opts ...option.RequestOption) (res *WaitingRoom, err error) {
	var env WaitingRoomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured waiting room.
func (r *WaitingRoomService) Update(ctx context.Context, waitingRoomID string, params WaitingRoomUpdateParams, opts ...option.RequestOption) (res *WaitingRoom, err error) {
	var env WaitingRoomUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", params.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists waiting rooms for account or zone.
func (r *WaitingRoomService) List(ctx context.Context, params WaitingRoomListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WaitingRoom], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/waiting_rooms", accountOrZone, accountOrZoneID)
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

// Lists waiting rooms for account or zone.
func (r *WaitingRoomService) ListAutoPaging(ctx context.Context, params WaitingRoomListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WaitingRoom] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a waiting room.
func (r *WaitingRoomService) Delete(ctx context.Context, waitingRoomID string, body WaitingRoomDeleteParams, opts ...option.RequestOption) (res *WaitingRoomDeleteResponse, err error) {
	var env WaitingRoomDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", body.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches a configured waiting room.
func (r *WaitingRoomService) Edit(ctx context.Context, waitingRoomID string, params WaitingRoomEditParams, opts ...option.RequestOption) (res *WaitingRoom, err error) {
	var env WaitingRoomEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", params.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single configured waiting room.
func (r *WaitingRoomService) Get(ctx context.Context, waitingRoomID string, query WaitingRoomGetParams, opts ...option.RequestOption) (res *WaitingRoom, err error) {
	var env WaitingRoomGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s", query.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AdditionalRoutes struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host string `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string               `json:"path"`
	JSON additionalRoutesJSON `json:"-"`
}

// additionalRoutesJSON contains the JSON metadata for the struct
// [AdditionalRoutes]
type additionalRoutesJSON struct {
	Host        apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdditionalRoutes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r additionalRoutesJSON) RawJSON() string {
	return r.raw
}

type AdditionalRoutesParam struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host param.Field[string] `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
}

func (r AdditionalRoutesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type CookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite CookieAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure CookieAttributesSecure `json:"secure"`
	JSON   cookieAttributesJSON   `json:"-"`
}

// cookieAttributesJSON contains the JSON metadata for the struct
// [CookieAttributes]
type cookieAttributesJSON struct {
	Samesite    apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CookieAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type CookieAttributesSamesite string

const (
	CookieAttributesSamesiteAuto   CookieAttributesSamesite = "auto"
	CookieAttributesSamesiteLax    CookieAttributesSamesite = "lax"
	CookieAttributesSamesiteNone   CookieAttributesSamesite = "none"
	CookieAttributesSamesiteStrict CookieAttributesSamesite = "strict"
)

func (r CookieAttributesSamesite) IsKnown() bool {
	switch r {
	case CookieAttributesSamesiteAuto, CookieAttributesSamesiteLax, CookieAttributesSamesiteNone, CookieAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type CookieAttributesSecure string

const (
	CookieAttributesSecureAuto   CookieAttributesSecure = "auto"
	CookieAttributesSecureAlways CookieAttributesSecure = "always"
	CookieAttributesSecureNever  CookieAttributesSecure = "never"
)

func (r CookieAttributesSecure) IsKnown() bool {
	switch r {
	case CookieAttributesSecureAuto, CookieAttributesSecureAlways, CookieAttributesSecureNever:
		return true
	}
	return false
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type CookieAttributesParam struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[CookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[CookieAttributesSecure] `json:"secure"`
}

func (r CookieAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueryParam struct {
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host param.Field[string] `json:"host,required"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute,required"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers param.Field[int64] `json:"total_active_users,required"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes param.Field[[]AdditionalRoutesParam] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[CookieAttributesParam] `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix param.Field[string] `json:"cookie_suffix"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage param.Field[QueryDefaultTemplateLanguage] `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description param.Field[string] `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// A list of enabled origin commands.
	EnabledOriginCommands param.Field[[]QueryEnabledOriginCommand] `json:"enabled_origin_commands"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time —\_it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//  24. `turnstile`: Empty when turnstile isn't enabled. String displaying an html
	//     tag to display the Turnstile widget. Please add the `{{{turnstile}}}` tag to
	//     the `custom_html` template to ensure the Turnstile widget appears.
	//  25. `infiniteQueue`: Boolean indicating whether the response is for a user in
	//     the infinite queue.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}
	JsonResponseEnabled param.Field[bool] `json:"json_response_enabled"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll param.Field[bool] `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod param.Field[QueryQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[QueryQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
	// Which action to take when a bot is detected using Turnstile. `log` will have no
	// impact on queueing behavior, simply keeping track of how many bots are detected
	// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
	// state, where they will never reach your origin. `infinite_queue` requires
	// Advanced Waiting Room.
	TurnstileAction param.Field[QueryTurnstileAction] `json:"turnstile_action"`
	// Which Turnstile widget type to use for detecting bot traffic. See
	// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
	// for the definitions of these widget types. Set to `off` to disable the Turnstile
	// integration entirely. Setting this to anything other than `off` or `invisible`
	// requires Advanced Waiting Room.
	TurnstileMode param.Field[QueryTurnstileMode] `json:"turnstile_mode"`
}

func (r QueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type QueryDefaultTemplateLanguage string

const (
	QueryDefaultTemplateLanguageEnUs QueryDefaultTemplateLanguage = "en-US"
	QueryDefaultTemplateLanguageEsEs QueryDefaultTemplateLanguage = "es-ES"
	QueryDefaultTemplateLanguageDeDe QueryDefaultTemplateLanguage = "de-DE"
	QueryDefaultTemplateLanguageFrFr QueryDefaultTemplateLanguage = "fr-FR"
	QueryDefaultTemplateLanguageItIt QueryDefaultTemplateLanguage = "it-IT"
	QueryDefaultTemplateLanguageJaJp QueryDefaultTemplateLanguage = "ja-JP"
	QueryDefaultTemplateLanguageKoKr QueryDefaultTemplateLanguage = "ko-KR"
	QueryDefaultTemplateLanguagePtBr QueryDefaultTemplateLanguage = "pt-BR"
	QueryDefaultTemplateLanguageZhCn QueryDefaultTemplateLanguage = "zh-CN"
	QueryDefaultTemplateLanguageZhTw QueryDefaultTemplateLanguage = "zh-TW"
	QueryDefaultTemplateLanguageNlNl QueryDefaultTemplateLanguage = "nl-NL"
	QueryDefaultTemplateLanguagePlPl QueryDefaultTemplateLanguage = "pl-PL"
	QueryDefaultTemplateLanguageIDID QueryDefaultTemplateLanguage = "id-ID"
	QueryDefaultTemplateLanguageTrTr QueryDefaultTemplateLanguage = "tr-TR"
	QueryDefaultTemplateLanguageArEg QueryDefaultTemplateLanguage = "ar-EG"
	QueryDefaultTemplateLanguageRuRu QueryDefaultTemplateLanguage = "ru-RU"
	QueryDefaultTemplateLanguageFaIr QueryDefaultTemplateLanguage = "fa-IR"
	QueryDefaultTemplateLanguageBgBg QueryDefaultTemplateLanguage = "bg-BG"
	QueryDefaultTemplateLanguageHrHr QueryDefaultTemplateLanguage = "hr-HR"
	QueryDefaultTemplateLanguageCsCz QueryDefaultTemplateLanguage = "cs-CZ"
	QueryDefaultTemplateLanguageDaDK QueryDefaultTemplateLanguage = "da-DK"
	QueryDefaultTemplateLanguageFiFi QueryDefaultTemplateLanguage = "fi-FI"
	QueryDefaultTemplateLanguageLtLt QueryDefaultTemplateLanguage = "lt-LT"
	QueryDefaultTemplateLanguageMsMy QueryDefaultTemplateLanguage = "ms-MY"
	QueryDefaultTemplateLanguageNbNo QueryDefaultTemplateLanguage = "nb-NO"
	QueryDefaultTemplateLanguageRoRo QueryDefaultTemplateLanguage = "ro-RO"
	QueryDefaultTemplateLanguageElGr QueryDefaultTemplateLanguage = "el-GR"
	QueryDefaultTemplateLanguageHeIl QueryDefaultTemplateLanguage = "he-IL"
	QueryDefaultTemplateLanguageHiIn QueryDefaultTemplateLanguage = "hi-IN"
	QueryDefaultTemplateLanguageHuHu QueryDefaultTemplateLanguage = "hu-HU"
	QueryDefaultTemplateLanguageSrBa QueryDefaultTemplateLanguage = "sr-BA"
	QueryDefaultTemplateLanguageSkSk QueryDefaultTemplateLanguage = "sk-SK"
	QueryDefaultTemplateLanguageSlSi QueryDefaultTemplateLanguage = "sl-SI"
	QueryDefaultTemplateLanguageSvSe QueryDefaultTemplateLanguage = "sv-SE"
	QueryDefaultTemplateLanguageTlPh QueryDefaultTemplateLanguage = "tl-PH"
	QueryDefaultTemplateLanguageThTh QueryDefaultTemplateLanguage = "th-TH"
	QueryDefaultTemplateLanguageUkUA QueryDefaultTemplateLanguage = "uk-UA"
	QueryDefaultTemplateLanguageViVn QueryDefaultTemplateLanguage = "vi-VN"
)

func (r QueryDefaultTemplateLanguage) IsKnown() bool {
	switch r {
	case QueryDefaultTemplateLanguageEnUs, QueryDefaultTemplateLanguageEsEs, QueryDefaultTemplateLanguageDeDe, QueryDefaultTemplateLanguageFrFr, QueryDefaultTemplateLanguageItIt, QueryDefaultTemplateLanguageJaJp, QueryDefaultTemplateLanguageKoKr, QueryDefaultTemplateLanguagePtBr, QueryDefaultTemplateLanguageZhCn, QueryDefaultTemplateLanguageZhTw, QueryDefaultTemplateLanguageNlNl, QueryDefaultTemplateLanguagePlPl, QueryDefaultTemplateLanguageIDID, QueryDefaultTemplateLanguageTrTr, QueryDefaultTemplateLanguageArEg, QueryDefaultTemplateLanguageRuRu, QueryDefaultTemplateLanguageFaIr, QueryDefaultTemplateLanguageBgBg, QueryDefaultTemplateLanguageHrHr, QueryDefaultTemplateLanguageCsCz, QueryDefaultTemplateLanguageDaDK, QueryDefaultTemplateLanguageFiFi, QueryDefaultTemplateLanguageLtLt, QueryDefaultTemplateLanguageMsMy, QueryDefaultTemplateLanguageNbNo, QueryDefaultTemplateLanguageRoRo, QueryDefaultTemplateLanguageElGr, QueryDefaultTemplateLanguageHeIl, QueryDefaultTemplateLanguageHiIn, QueryDefaultTemplateLanguageHuHu, QueryDefaultTemplateLanguageSrBa, QueryDefaultTemplateLanguageSkSk, QueryDefaultTemplateLanguageSlSi, QueryDefaultTemplateLanguageSvSe, QueryDefaultTemplateLanguageTlPh, QueryDefaultTemplateLanguageThTh, QueryDefaultTemplateLanguageUkUA, QueryDefaultTemplateLanguageViVn:
		return true
	}
	return false
}

type QueryEnabledOriginCommand string

const (
	QueryEnabledOriginCommandRevoke QueryEnabledOriginCommand = "revoke"
)

func (r QueryEnabledOriginCommand) IsKnown() bool {
	switch r {
	case QueryEnabledOriginCommandRevoke:
		return true
	}
	return false
}

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type QueryQueueingMethod string

const (
	QueryQueueingMethodFifo        QueryQueueingMethod = "fifo"
	QueryQueueingMethodRandom      QueryQueueingMethod = "random"
	QueryQueueingMethodPassthrough QueryQueueingMethod = "passthrough"
	QueryQueueingMethodReject      QueryQueueingMethod = "reject"
)

func (r QueryQueueingMethod) IsKnown() bool {
	switch r {
	case QueryQueueingMethodFifo, QueryQueueingMethodRandom, QueryQueueingMethodPassthrough, QueryQueueingMethodReject:
		return true
	}
	return false
}

// HTTP status code returned to a user while in the queue.
type QueryQueueingStatusCode int64

const (
	QueryQueueingStatusCode200 QueryQueueingStatusCode = 200
	QueryQueueingStatusCode202 QueryQueueingStatusCode = 202
	QueryQueueingStatusCode429 QueryQueueingStatusCode = 429
)

func (r QueryQueueingStatusCode) IsKnown() bool {
	switch r {
	case QueryQueueingStatusCode200, QueryQueueingStatusCode202, QueryQueueingStatusCode429:
		return true
	}
	return false
}

// Which action to take when a bot is detected using Turnstile. `log` will have no
// impact on queueing behavior, simply keeping track of how many bots are detected
// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
// state, where they will never reach your origin. `infinite_queue` requires
// Advanced Waiting Room.
type QueryTurnstileAction string

const (
	QueryTurnstileActionLog           QueryTurnstileAction = "log"
	QueryTurnstileActionInfiniteQueue QueryTurnstileAction = "infinite_queue"
)

func (r QueryTurnstileAction) IsKnown() bool {
	switch r {
	case QueryTurnstileActionLog, QueryTurnstileActionInfiniteQueue:
		return true
	}
	return false
}

// Which Turnstile widget type to use for detecting bot traffic. See
// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
// for the definitions of these widget types. Set to `off` to disable the Turnstile
// integration entirely. Setting this to anything other than `off` or `invisible`
// requires Advanced Waiting Room.
type QueryTurnstileMode string

const (
	QueryTurnstileModeOff                   QueryTurnstileMode = "off"
	QueryTurnstileModeInvisible             QueryTurnstileMode = "invisible"
	QueryTurnstileModeVisibleNonInteractive QueryTurnstileMode = "visible_non_interactive"
	QueryTurnstileModeVisibleManaged        QueryTurnstileMode = "visible_managed"
)

func (r QueryTurnstileMode) IsKnown() bool {
	switch r {
	case QueryTurnstileModeOff, QueryTurnstileModeInvisible, QueryTurnstileModeVisibleNonInteractive, QueryTurnstileModeVisibleManaged:
		return true
	}
	return false
}

type WaitingRoom struct {
	ID string `json:"id"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes []AdditionalRoutes `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes CookieAttributes `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix string    `json:"cookie_suffix"`
	CreatedOn    time.Time `json:"created_on" format:"date-time"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML string `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage WaitingRoomDefaultTemplateLanguage `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description string `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal bool `json:"disable_session_renewal"`
	// A list of enabled origin commands.
	EnabledOriginCommands []WaitingRoomEnabledOriginCommand `json:"enabled_origin_commands"`
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host string `json:"host"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time —\_it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//  24. `turnstile`: Empty when turnstile isn't enabled. String displaying an html
	//     tag to display the Turnstile widget. Please add the `{{{turnstile}}}` tag to
	//     the `custom_html` template to ensure the Turnstile widget appears.
	//  25. `infiniteQueue`: Boolean indicating whether the response is for a user in
	//     the infinite queue.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}
	JsonResponseEnabled bool      `json:"json_response_enabled"`
	ModifiedOn          time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name string `json:"name"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute int64 `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when the next event will begin queueing.
	NextEventPrequeueStartTime string `json:"next_event_prequeue_start_time,nullable"`
	// An ISO 8601 timestamp that marks when the next event will start.
	NextEventStartTime string `json:"next_event_start_time,nullable"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll bool `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod WaitingRoomQueueingMethod `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode WaitingRoomQueueingStatusCode `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration int64 `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended bool `json:"suspended"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers int64 `json:"total_active_users"`
	// Which action to take when a bot is detected using Turnstile. `log` will have no
	// impact on queueing behavior, simply keeping track of how many bots are detected
	// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
	// state, where they will never reach your origin. `infinite_queue` requires
	// Advanced Waiting Room.
	TurnstileAction WaitingRoomTurnstileAction `json:"turnstile_action"`
	// Which Turnstile widget type to use for detecting bot traffic. See
	// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
	// for the definitions of these widget types. Set to `off` to disable the Turnstile
	// integration entirely. Setting this to anything other than `off` or `invisible`
	// requires Advanced Waiting Room.
	TurnstileMode WaitingRoomTurnstileMode `json:"turnstile_mode"`
	JSON          waitingRoomJSON          `json:"-"`
}

// waitingRoomJSON contains the JSON metadata for the struct [WaitingRoom]
type waitingRoomJSON struct {
	ID                         apijson.Field
	AdditionalRoutes           apijson.Field
	CookieAttributes           apijson.Field
	CookieSuffix               apijson.Field
	CreatedOn                  apijson.Field
	CustomPageHTML             apijson.Field
	DefaultTemplateLanguage    apijson.Field
	Description                apijson.Field
	DisableSessionRenewal      apijson.Field
	EnabledOriginCommands      apijson.Field
	Host                       apijson.Field
	JsonResponseEnabled        apijson.Field
	ModifiedOn                 apijson.Field
	Name                       apijson.Field
	NewUsersPerMinute          apijson.Field
	NextEventPrequeueStartTime apijson.Field
	NextEventStartTime         apijson.Field
	Path                       apijson.Field
	QueueAll                   apijson.Field
	QueueingMethod             apijson.Field
	QueueingStatusCode         apijson.Field
	SessionDuration            apijson.Field
	Suspended                  apijson.Field
	TotalActiveUsers           apijson.Field
	TurnstileAction            apijson.Field
	TurnstileMode              apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *WaitingRoom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomJSON) RawJSON() string {
	return r.raw
}

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type WaitingRoomDefaultTemplateLanguage string

const (
	WaitingRoomDefaultTemplateLanguageEnUs WaitingRoomDefaultTemplateLanguage = "en-US"
	WaitingRoomDefaultTemplateLanguageEsEs WaitingRoomDefaultTemplateLanguage = "es-ES"
	WaitingRoomDefaultTemplateLanguageDeDe WaitingRoomDefaultTemplateLanguage = "de-DE"
	WaitingRoomDefaultTemplateLanguageFrFr WaitingRoomDefaultTemplateLanguage = "fr-FR"
	WaitingRoomDefaultTemplateLanguageItIt WaitingRoomDefaultTemplateLanguage = "it-IT"
	WaitingRoomDefaultTemplateLanguageJaJp WaitingRoomDefaultTemplateLanguage = "ja-JP"
	WaitingRoomDefaultTemplateLanguageKoKr WaitingRoomDefaultTemplateLanguage = "ko-KR"
	WaitingRoomDefaultTemplateLanguagePtBr WaitingRoomDefaultTemplateLanguage = "pt-BR"
	WaitingRoomDefaultTemplateLanguageZhCn WaitingRoomDefaultTemplateLanguage = "zh-CN"
	WaitingRoomDefaultTemplateLanguageZhTw WaitingRoomDefaultTemplateLanguage = "zh-TW"
	WaitingRoomDefaultTemplateLanguageNlNl WaitingRoomDefaultTemplateLanguage = "nl-NL"
	WaitingRoomDefaultTemplateLanguagePlPl WaitingRoomDefaultTemplateLanguage = "pl-PL"
	WaitingRoomDefaultTemplateLanguageIDID WaitingRoomDefaultTemplateLanguage = "id-ID"
	WaitingRoomDefaultTemplateLanguageTrTr WaitingRoomDefaultTemplateLanguage = "tr-TR"
	WaitingRoomDefaultTemplateLanguageArEg WaitingRoomDefaultTemplateLanguage = "ar-EG"
	WaitingRoomDefaultTemplateLanguageRuRu WaitingRoomDefaultTemplateLanguage = "ru-RU"
	WaitingRoomDefaultTemplateLanguageFaIr WaitingRoomDefaultTemplateLanguage = "fa-IR"
	WaitingRoomDefaultTemplateLanguageBgBg WaitingRoomDefaultTemplateLanguage = "bg-BG"
	WaitingRoomDefaultTemplateLanguageHrHr WaitingRoomDefaultTemplateLanguage = "hr-HR"
	WaitingRoomDefaultTemplateLanguageCsCz WaitingRoomDefaultTemplateLanguage = "cs-CZ"
	WaitingRoomDefaultTemplateLanguageDaDK WaitingRoomDefaultTemplateLanguage = "da-DK"
	WaitingRoomDefaultTemplateLanguageFiFi WaitingRoomDefaultTemplateLanguage = "fi-FI"
	WaitingRoomDefaultTemplateLanguageLtLt WaitingRoomDefaultTemplateLanguage = "lt-LT"
	WaitingRoomDefaultTemplateLanguageMsMy WaitingRoomDefaultTemplateLanguage = "ms-MY"
	WaitingRoomDefaultTemplateLanguageNbNo WaitingRoomDefaultTemplateLanguage = "nb-NO"
	WaitingRoomDefaultTemplateLanguageRoRo WaitingRoomDefaultTemplateLanguage = "ro-RO"
	WaitingRoomDefaultTemplateLanguageElGr WaitingRoomDefaultTemplateLanguage = "el-GR"
	WaitingRoomDefaultTemplateLanguageHeIl WaitingRoomDefaultTemplateLanguage = "he-IL"
	WaitingRoomDefaultTemplateLanguageHiIn WaitingRoomDefaultTemplateLanguage = "hi-IN"
	WaitingRoomDefaultTemplateLanguageHuHu WaitingRoomDefaultTemplateLanguage = "hu-HU"
	WaitingRoomDefaultTemplateLanguageSrBa WaitingRoomDefaultTemplateLanguage = "sr-BA"
	WaitingRoomDefaultTemplateLanguageSkSk WaitingRoomDefaultTemplateLanguage = "sk-SK"
	WaitingRoomDefaultTemplateLanguageSlSi WaitingRoomDefaultTemplateLanguage = "sl-SI"
	WaitingRoomDefaultTemplateLanguageSvSe WaitingRoomDefaultTemplateLanguage = "sv-SE"
	WaitingRoomDefaultTemplateLanguageTlPh WaitingRoomDefaultTemplateLanguage = "tl-PH"
	WaitingRoomDefaultTemplateLanguageThTh WaitingRoomDefaultTemplateLanguage = "th-TH"
	WaitingRoomDefaultTemplateLanguageUkUA WaitingRoomDefaultTemplateLanguage = "uk-UA"
	WaitingRoomDefaultTemplateLanguageViVn WaitingRoomDefaultTemplateLanguage = "vi-VN"
)

func (r WaitingRoomDefaultTemplateLanguage) IsKnown() bool {
	switch r {
	case WaitingRoomDefaultTemplateLanguageEnUs, WaitingRoomDefaultTemplateLanguageEsEs, WaitingRoomDefaultTemplateLanguageDeDe, WaitingRoomDefaultTemplateLanguageFrFr, WaitingRoomDefaultTemplateLanguageItIt, WaitingRoomDefaultTemplateLanguageJaJp, WaitingRoomDefaultTemplateLanguageKoKr, WaitingRoomDefaultTemplateLanguagePtBr, WaitingRoomDefaultTemplateLanguageZhCn, WaitingRoomDefaultTemplateLanguageZhTw, WaitingRoomDefaultTemplateLanguageNlNl, WaitingRoomDefaultTemplateLanguagePlPl, WaitingRoomDefaultTemplateLanguageIDID, WaitingRoomDefaultTemplateLanguageTrTr, WaitingRoomDefaultTemplateLanguageArEg, WaitingRoomDefaultTemplateLanguageRuRu, WaitingRoomDefaultTemplateLanguageFaIr, WaitingRoomDefaultTemplateLanguageBgBg, WaitingRoomDefaultTemplateLanguageHrHr, WaitingRoomDefaultTemplateLanguageCsCz, WaitingRoomDefaultTemplateLanguageDaDK, WaitingRoomDefaultTemplateLanguageFiFi, WaitingRoomDefaultTemplateLanguageLtLt, WaitingRoomDefaultTemplateLanguageMsMy, WaitingRoomDefaultTemplateLanguageNbNo, WaitingRoomDefaultTemplateLanguageRoRo, WaitingRoomDefaultTemplateLanguageElGr, WaitingRoomDefaultTemplateLanguageHeIl, WaitingRoomDefaultTemplateLanguageHiIn, WaitingRoomDefaultTemplateLanguageHuHu, WaitingRoomDefaultTemplateLanguageSrBa, WaitingRoomDefaultTemplateLanguageSkSk, WaitingRoomDefaultTemplateLanguageSlSi, WaitingRoomDefaultTemplateLanguageSvSe, WaitingRoomDefaultTemplateLanguageTlPh, WaitingRoomDefaultTemplateLanguageThTh, WaitingRoomDefaultTemplateLanguageUkUA, WaitingRoomDefaultTemplateLanguageViVn:
		return true
	}
	return false
}

type WaitingRoomEnabledOriginCommand string

const (
	WaitingRoomEnabledOriginCommandRevoke WaitingRoomEnabledOriginCommand = "revoke"
)

func (r WaitingRoomEnabledOriginCommand) IsKnown() bool {
	switch r {
	case WaitingRoomEnabledOriginCommandRevoke:
		return true
	}
	return false
}

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type WaitingRoomQueueingMethod string

const (
	WaitingRoomQueueingMethodFifo        WaitingRoomQueueingMethod = "fifo"
	WaitingRoomQueueingMethodRandom      WaitingRoomQueueingMethod = "random"
	WaitingRoomQueueingMethodPassthrough WaitingRoomQueueingMethod = "passthrough"
	WaitingRoomQueueingMethodReject      WaitingRoomQueueingMethod = "reject"
)

func (r WaitingRoomQueueingMethod) IsKnown() bool {
	switch r {
	case WaitingRoomQueueingMethodFifo, WaitingRoomQueueingMethodRandom, WaitingRoomQueueingMethodPassthrough, WaitingRoomQueueingMethodReject:
		return true
	}
	return false
}

// HTTP status code returned to a user while in the queue.
type WaitingRoomQueueingStatusCode int64

const (
	WaitingRoomQueueingStatusCode200 WaitingRoomQueueingStatusCode = 200
	WaitingRoomQueueingStatusCode202 WaitingRoomQueueingStatusCode = 202
	WaitingRoomQueueingStatusCode429 WaitingRoomQueueingStatusCode = 429
)

func (r WaitingRoomQueueingStatusCode) IsKnown() bool {
	switch r {
	case WaitingRoomQueueingStatusCode200, WaitingRoomQueueingStatusCode202, WaitingRoomQueueingStatusCode429:
		return true
	}
	return false
}

// Which action to take when a bot is detected using Turnstile. `log` will have no
// impact on queueing behavior, simply keeping track of how many bots are detected
// in Waiting Room Analytics. `infinite_queue` will send bots to a false queueing
// state, where they will never reach your origin. `infinite_queue` requires
// Advanced Waiting Room.
type WaitingRoomTurnstileAction string

const (
	WaitingRoomTurnstileActionLog           WaitingRoomTurnstileAction = "log"
	WaitingRoomTurnstileActionInfiniteQueue WaitingRoomTurnstileAction = "infinite_queue"
)

func (r WaitingRoomTurnstileAction) IsKnown() bool {
	switch r {
	case WaitingRoomTurnstileActionLog, WaitingRoomTurnstileActionInfiniteQueue:
		return true
	}
	return false
}

// Which Turnstile widget type to use for detecting bot traffic. See
// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
// for the definitions of these widget types. Set to `off` to disable the Turnstile
// integration entirely. Setting this to anything other than `off` or `invisible`
// requires Advanced Waiting Room.
type WaitingRoomTurnstileMode string

const (
	WaitingRoomTurnstileModeOff                   WaitingRoomTurnstileMode = "off"
	WaitingRoomTurnstileModeInvisible             WaitingRoomTurnstileMode = "invisible"
	WaitingRoomTurnstileModeVisibleNonInteractive WaitingRoomTurnstileMode = "visible_non_interactive"
	WaitingRoomTurnstileModeVisibleManaged        WaitingRoomTurnstileMode = "visible_managed"
)

func (r WaitingRoomTurnstileMode) IsKnown() bool {
	switch r {
	case WaitingRoomTurnstileModeOff, WaitingRoomTurnstileModeInvisible, WaitingRoomTurnstileModeVisibleNonInteractive, WaitingRoomTurnstileModeVisibleManaged:
		return true
	}
	return false
}

type WaitingRoomDeleteResponse struct {
	ID   string                        `json:"id"`
	JSON waitingRoomDeleteResponseJSON `json:"-"`
}

// waitingRoomDeleteResponseJSON contains the JSON metadata for the struct
// [WaitingRoomDeleteResponse]
type waitingRoomDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Query  QueryParam          `json:"query,required"`
}

func (r WaitingRoomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Query)
}

type WaitingRoomNewResponseEnvelope struct {
	Result WaitingRoom                        `json:"result,required"`
	JSON   waitingRoomNewResponseEnvelopeJSON `json:"-"`
}

// waitingRoomNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomNewResponseEnvelope]
type waitingRoomNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Query  QueryParam          `json:"query,required"`
}

func (r WaitingRoomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Query)
}

type WaitingRoomUpdateResponseEnvelope struct {
	Result WaitingRoom                           `json:"result,required"`
	JSON   waitingRoomUpdateResponseEnvelopeJSON `json:"-"`
}

// waitingRoomUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomUpdateResponseEnvelope]
type waitingRoomUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page. Must be a multiple of 5.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WaitingRoomListParams]'s query parameters as `url.Values`.
func (r WaitingRoomListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WaitingRoomDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WaitingRoomDeleteResponseEnvelope struct {
	Result WaitingRoomDeleteResponse             `json:"result,required"`
	JSON   waitingRoomDeleteResponseEnvelopeJSON `json:"-"`
}

// waitingRoomDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomDeleteResponseEnvelope]
type waitingRoomDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Query  QueryParam          `json:"query,required"`
}

func (r WaitingRoomEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Query)
}

type WaitingRoomEditResponseEnvelope struct {
	Result WaitingRoom                         `json:"result,required"`
	JSON   waitingRoomEditResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomEditResponseEnvelope]
type waitingRoomEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WaitingRoomGetResponseEnvelope struct {
	Result WaitingRoom                        `json:"result,required"`
	JSON   waitingRoomGetResponseEnvelopeJSON `json:"-"`
}

// waitingRoomGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomGetResponseEnvelope]
type waitingRoomGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
