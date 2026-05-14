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

// PageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
	Options []option.RequestOption
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r *PageService) {
	r = &PageService{}
	r.Options = opts
	return
}

// Creates a waiting room page preview. Upload a custom waiting room page for
// preview. You will receive a preview URL in the form
// `http://waitingrooms.dev/preview/<uuid>`. You can use the following query
// parameters to change the state of the preview:
//
//  1. `force_queue`: Boolean indicating if all users will be queued in the waiting
//     room and no one will be let into the origin website (also known as queueAll).
//  2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently
//     full and not accepting new users at the moment.
//  3. `queueing_method`: The queueing method currently used by the waiting room.
//     - **fifo** indicates a FIFO queue.
//     - **random** indicates a Random queue.
//     - **passthrough** indicates a Passthrough queue. Keep in mind that the
//     waiting room page will only be displayed if `force_queue=true` or
//     `event=prequeueing` — for other cases the request will pass through to the
//     origin. For our preview, this will be a fake origin website returning
//     \"Welcome\".
//     - **reject** indicates a Reject queue.
//  4. `event`: Used to preview a waiting room event.
//     - **none** indicates no event is occurring.
//     - **prequeueing** indicates that an event is prequeueing (between
//     `prequeue_start_time` and `event_start_time`).
//     - **started** indicates that an event has started (between `event_start_time`
//     and `event_end_time`).
//  5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users
//     in the prequeue when it starts. This can only be set to **true** if an event
//     is active (`event` is not **none**).
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true` 6.
// `waitTime`: Non-zero, positive integer indicating the estimated wait time in
// minutes. The default value is 10 minutes.
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated
// wait time as 50 minutes.
func (r *PageService) Preview(ctx context.Context, params PagePreviewParams, opts ...option.RequestOption) (res *PagePreviewResponse, err error) {
	var env PagePreviewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/preview", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PagePreviewResponse struct {
	// URL where the custom waiting room page can temporarily be previewed.
	PreviewURL string                  `json:"preview_url"`
	JSON       pagePreviewResponseJSON `json:"-"`
}

// pagePreviewResponseJSON contains the JSON metadata for the struct
// [PagePreviewResponse]
type pagePreviewResponseJSON struct {
	PreviewURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagePreviewResponseJSON) RawJSON() string {
	return r.raw
}

type PagePreviewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	CustomHTML param.Field[string] `json:"custom_html,required"`
}

func (r PagePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PagePreviewResponseEnvelope struct {
	Result PagePreviewResponse             `json:"result,required"`
	JSON   pagePreviewResponseEnvelopeJSON `json:"-"`
}

// pagePreviewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PagePreviewResponseEnvelope]
type pagePreviewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagePreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagePreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
