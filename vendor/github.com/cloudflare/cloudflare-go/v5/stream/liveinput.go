// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

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
)

// LiveInputService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLiveInputService] method instead.
type LiveInputService struct {
	Options []option.RequestOption
	Outputs *LiveInputOutputService
}

// NewLiveInputService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLiveInputService(opts ...option.RequestOption) (r *LiveInputService) {
	r = &LiveInputService{}
	r.Options = opts
	r.Outputs = NewLiveInputOutputService(opts...)
	return
}

// Creates a live input, and returns credentials that you or your users can use to
// stream live video to Cloudflare Stream.
func (r *LiveInputService) New(ctx context.Context, params LiveInputNewParams, opts ...option.RequestOption) (res *LiveInput, err error) {
	var env LiveInputNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specified live input.
func (r *LiveInputService) Update(ctx context.Context, liveInputIdentifier string, params LiveInputUpdateParams, opts ...option.RequestOption) (res *LiveInput, err error) {
	var env LiveInputUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", params.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the live inputs created for an account. To get the credentials needed to
// stream to a specific live input, request a single live input.
func (r *LiveInputService) List(ctx context.Context, params LiveInputListParams, opts ...option.RequestOption) (res *LiveInputListResponse, err error) {
	var env LiveInputListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Prevents a live input from being streamed to and makes the live input
// inaccessible to any future API calls.
func (r *LiveInputService) Delete(ctx context.Context, liveInputIdentifier string, body LiveInputDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", body.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves details of an existing live input.
func (r *LiveInputService) Get(ctx context.Context, liveInputIdentifier string, query LiveInputGetParams, opts ...option.RequestOption) (res *LiveInput, err error) {
	var env LiveInputGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if liveInputIdentifier == "" {
		err = errors.New("missing required live_input_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/live_inputs/%s", query.AccountID, liveInputIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a live input.
type LiveInput struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording LiveInputRecording `json:"recording"`
	// Details for streaming to an live input using RTMPS.
	Rtmps LiveInputRtmps `json:"rtmps"`
	// Details for playback from an live input using RTMPS.
	RtmpsPlayback LiveInputRtmpsPlayback `json:"rtmpsPlayback"`
	// Details for streaming to a live input using SRT.
	Srt LiveInputSrt `json:"srt"`
	// Details for playback from an live input using SRT.
	SrtPlayback LiveInputSrtPlayback `json:"srtPlayback"`
	// The connection status of a live input.
	Status LiveInputStatus `json:"status,nullable"`
	// A unique identifier for a live input.
	UID string `json:"uid"`
	// Details for streaming to a live input using WebRTC.
	WebRtc LiveInputWebRtc `json:"webRTC"`
	// Details for playback from a live input using WebRTC.
	WebRtcPlayback LiveInputWebRtcPlayback `json:"webRTCPlayback"`
	JSON           liveInputJSON           `json:"-"`
}

// liveInputJSON contains the JSON metadata for the struct [LiveInput]
type liveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	Recording                apijson.Field
	Rtmps                    apijson.Field
	RtmpsPlayback            apijson.Field
	Srt                      apijson.Field
	SrtPlayback              apijson.Field
	Status                   apijson.Field
	UID                      apijson.Field
	WebRtc                   apijson.Field
	WebRtcPlayback           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputJSON) RawJSON() string {
	return r.raw
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to
	// `true`.
	HideLiveViewerCount bool `json:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode LiveInputRecordingMode `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds int64                  `json:"timeoutSeconds"`
	JSON           liveInputRecordingJSON `json:"-"`
}

// liveInputRecordingJSON contains the JSON metadata for the struct
// [LiveInputRecording]
type liveInputRecordingJSON struct {
	AllowedOrigins      apijson.Field
	HideLiveViewerCount apijson.Field
	Mode                apijson.Field
	RequireSignedURLs   apijson.Field
	TimeoutSeconds      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LiveInputRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputRecordingJSON) RawJSON() string {
	return r.raw
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputRecordingMode string

const (
	LiveInputRecordingModeOff       LiveInputRecordingMode = "off"
	LiveInputRecordingModeAutomatic LiveInputRecordingMode = "automatic"
)

func (r LiveInputRecordingMode) IsKnown() bool {
	switch r {
	case LiveInputRecordingModeOff, LiveInputRecordingModeAutomatic:
		return true
	}
	return false
}

// Details for streaming to an live input using RTMPS.
type LiveInputRtmps struct {
	// The secret key to use when streaming via RTMPS to a live input.
	StreamKey string `json:"streamKey"`
	// The RTMPS URL you provide to the broadcaster, which they stream live video to.
	URL  string             `json:"url"`
	JSON liveInputRtmpsJSON `json:"-"`
}

// liveInputRtmpsJSON contains the JSON metadata for the struct [LiveInputRtmps]
type liveInputRtmpsJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputRtmps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputRtmpsJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using RTMPS.
type LiveInputRtmpsPlayback struct {
	// The secret key to use for playback via RTMPS.
	StreamKey string `json:"streamKey"`
	// The URL used to play live video over RTMPS.
	URL  string                     `json:"url"`
	JSON liveInputRtmpsPlaybackJSON `json:"-"`
}

// liveInputRtmpsPlaybackJSON contains the JSON metadata for the struct
// [LiveInputRtmpsPlayback]
type liveInputRtmpsPlaybackJSON struct {
	StreamKey   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputRtmpsPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputRtmpsPlaybackJSON) RawJSON() string {
	return r.raw
}

// Details for streaming to a live input using SRT.
type LiveInputSrt struct {
	// The secret key to use when streaming via SRT to a live input.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use when streaming via SRT.
	StreamID string `json:"streamId"`
	// The SRT URL you provide to the broadcaster, which they stream live video to.
	URL  string           `json:"url"`
	JSON liveInputSrtJSON `json:"-"`
}

// liveInputSrtJSON contains the JSON metadata for the struct [LiveInputSrt]
type liveInputSrtJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputSrt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputSrtJSON) RawJSON() string {
	return r.raw
}

// Details for playback from an live input using SRT.
type LiveInputSrtPlayback struct {
	// The secret key to use for playback via SRT.
	Passphrase string `json:"passphrase"`
	// The identifier of the live input to use for playback via SRT.
	StreamID string `json:"streamId"`
	// The URL used to play live video over SRT.
	URL  string                   `json:"url"`
	JSON liveInputSrtPlaybackJSON `json:"-"`
}

// liveInputSrtPlaybackJSON contains the JSON metadata for the struct
// [LiveInputSrtPlayback]
type liveInputSrtPlaybackJSON struct {
	Passphrase  apijson.Field
	StreamID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputSrtPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputSrtPlaybackJSON) RawJSON() string {
	return r.raw
}

// The connection status of a live input.
type LiveInputStatus string

const (
	LiveInputStatusConnected                LiveInputStatus = "connected"
	LiveInputStatusReconnected              LiveInputStatus = "reconnected"
	LiveInputStatusReconnecting             LiveInputStatus = "reconnecting"
	LiveInputStatusClientDisconnect         LiveInputStatus = "client_disconnect"
	LiveInputStatusTTLExceeded              LiveInputStatus = "ttl_exceeded"
	LiveInputStatusFailedToConnect          LiveInputStatus = "failed_to_connect"
	LiveInputStatusFailedToReconnect        LiveInputStatus = "failed_to_reconnect"
	LiveInputStatusNewConfigurationAccepted LiveInputStatus = "new_configuration_accepted"
)

func (r LiveInputStatus) IsKnown() bool {
	switch r {
	case LiveInputStatusConnected, LiveInputStatusReconnected, LiveInputStatusReconnecting, LiveInputStatusClientDisconnect, LiveInputStatusTTLExceeded, LiveInputStatusFailedToConnect, LiveInputStatusFailedToReconnect, LiveInputStatusNewConfigurationAccepted:
		return true
	}
	return false
}

// Details for streaming to a live input using WebRTC.
type LiveInputWebRtc struct {
	// The WebRTC URL you provide to the broadcaster, which they stream live video to.
	URL  string              `json:"url"`
	JSON liveInputWebRtcJSON `json:"-"`
}

// liveInputWebRtcJSON contains the JSON metadata for the struct [LiveInputWebRtc]
type liveInputWebRtcJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputWebRtc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputWebRtcJSON) RawJSON() string {
	return r.raw
}

// Details for playback from a live input using WebRTC.
type LiveInputWebRtcPlayback struct {
	// The URL used to play live video over WebRTC.
	URL  string                      `json:"url"`
	JSON liveInputWebRtcPlaybackJSON `json:"-"`
}

// liveInputWebRtcPlaybackJSON contains the JSON metadata for the struct
// [LiveInputWebRtcPlayback]
type liveInputWebRtcPlaybackJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputWebRtcPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputWebRtcPlaybackJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponse struct {
	LiveInputs []LiveInputListResponseLiveInput `json:"liveInputs"`
	// The total number of remaining live inputs based on cursor position.
	Range int64 `json:"range"`
	// The total number of live inputs that match the provided filters.
	Total int64                     `json:"total"`
	JSON  liveInputListResponseJSON `json:"-"`
}

// liveInputListResponseJSON contains the JSON metadata for the struct
// [LiveInputListResponse]
type liveInputListResponseJSON struct {
	LiveInputs  apijson.Field
	Range       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseLiveInput struct {
	// The date and time the live input was created.
	Created time.Time `json:"created" format:"date-time"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays float64 `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A unique identifier for a live input.
	UID  string                             `json:"uid"`
	JSON liveInputListResponseLiveInputJSON `json:"-"`
}

// liveInputListResponseLiveInputJSON contains the JSON metadata for the struct
// [LiveInputListResponseLiveInput]
type liveInputListResponseLiveInputJSON struct {
	Created                  apijson.Field
	DeleteRecordingAfterDays apijson.Field
	Meta                     apijson.Field
	Modified                 apijson.Field
	UID                      apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LiveInputListResponseLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseLiveInputJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[LiveInputNewParamsRecording] `json:"recording"`
}

func (r LiveInputNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputNewParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to
	// `true`.
	HideLiveViewerCount param.Field[bool] `json:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[LiveInputNewParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r LiveInputNewParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputNewParamsRecordingMode string

const (
	LiveInputNewParamsRecordingModeOff       LiveInputNewParamsRecordingMode = "off"
	LiveInputNewParamsRecordingModeAutomatic LiveInputNewParamsRecordingMode = "automatic"
)

func (r LiveInputNewParamsRecordingMode) IsKnown() bool {
	switch r {
	case LiveInputNewParamsRecordingModeOff, LiveInputNewParamsRecordingModeAutomatic:
		return true
	}
	return false
}

type LiveInputNewResponseEnvelope struct {
	Errors   []LiveInputNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputNewResponseEnvelopeSuccess `json:"success,required"`
	// Details about a live input.
	Result LiveInput                        `json:"result"`
	JSON   liveInputNewResponseEnvelopeJSON `json:"-"`
}

// liveInputNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputNewResponseEnvelope]
type liveInputNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           LiveInputNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LiveInputNewResponseEnvelopeErrors]
type liveInputNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    liveInputNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [LiveInputNewResponseEnvelopeErrorsSource]
type liveInputNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           LiveInputNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputNewResponseEnvelopeMessages]
type liveInputNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    liveInputNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [LiveInputNewResponseEnvelopeMessagesSource]
type liveInputNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputNewResponseEnvelopeSuccess bool

const (
	LiveInputNewResponseEnvelopeSuccessTrue LiveInputNewResponseEnvelopeSuccess = true
)

func (r LiveInputNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Sets the creator ID asssociated with this live input.
	DefaultCreator param.Field[string] `json:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be
	// deleted. When a stream completes and the recording is ready, the value is used
	// to calculate a scheduled deletion date for that recording. Omit the field to
	// indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion.
	DeleteRecordingAfterDays param.Field[float64] `json:"deleteRecordingAfterDays"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing live inputs.
	Meta param.Field[interface{}] `json:"meta"`
	// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
	// most cases, the video will initially be viewable as a live video and transition
	// to on-demand after a condition is satisfied.
	Recording param.Field[LiveInputUpdateParamsRecording] `json:"recording"`
}

func (r LiveInputUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Records the input to a Cloudflare Stream video. Behavior depends on the mode. In
// most cases, the video will initially be viewable as a live video and transition
// to on-demand after a condition is satisfied.
type LiveInputUpdateParamsRecording struct {
	// Lists the origins allowed to display videos created with this input. Enter
	// allowed origin domains in an array and use `*` for wildcard subdomains. An empty
	// array allows videos to be viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to
	// `true`.
	HideLiveViewerCount param.Field[bool] `json:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input. Set this value to `off` to
	// prevent a recording. Set the value to `automatic` to begin a recording and
	// transition to on-demand after Stream Live stops receiving input.
	Mode param.Field[LiveInputUpdateParamsRecordingMode] `json:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property
	// set. Also enforces access controls on any video recording of the livestream with
	// the live input.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Determines the amount of time a live input configured in `automatic` mode should
	// wait before a recording transitions from live to on-demand. `0` is recommended
	// for most use cases and indicates the platform default should be used.
	TimeoutSeconds param.Field[int64] `json:"timeoutSeconds"`
}

func (r LiveInputUpdateParamsRecording) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the recording behavior for the live input. Set this value to `off` to
// prevent a recording. Set the value to `automatic` to begin a recording and
// transition to on-demand after Stream Live stops receiving input.
type LiveInputUpdateParamsRecordingMode string

const (
	LiveInputUpdateParamsRecordingModeOff       LiveInputUpdateParamsRecordingMode = "off"
	LiveInputUpdateParamsRecordingModeAutomatic LiveInputUpdateParamsRecordingMode = "automatic"
)

func (r LiveInputUpdateParamsRecordingMode) IsKnown() bool {
	switch r {
	case LiveInputUpdateParamsRecordingModeOff, LiveInputUpdateParamsRecordingModeAutomatic:
		return true
	}
	return false
}

type LiveInputUpdateResponseEnvelope struct {
	Errors   []LiveInputUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Details about a live input.
	Result LiveInput                           `json:"result"`
	JSON   liveInputUpdateResponseEnvelopeJSON `json:"-"`
}

// liveInputUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputUpdateResponseEnvelope]
type liveInputUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           LiveInputUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseEnvelopeErrors]
type liveInputUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    liveInputUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [LiveInputUpdateResponseEnvelopeErrorsSource]
type liveInputUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           LiveInputUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputUpdateResponseEnvelopeMessages]
type liveInputUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    liveInputUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [LiveInputUpdateResponseEnvelopeMessagesSource]
type liveInputUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputUpdateResponseEnvelopeSuccess bool

const (
	LiveInputUpdateResponseEnvelopeSuccessTrue LiveInputUpdateResponseEnvelopeSuccess = true
)

func (r LiveInputUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
}

// URLQuery serializes [LiveInputListParams]'s query parameters as `url.Values`.
func (r LiveInputListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LiveInputListResponseEnvelope struct {
	Errors   []LiveInputListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputListResponseEnvelopeSuccess `json:"success,required"`
	Result  LiveInputListResponse                `json:"result"`
	JSON    liveInputListResponseEnvelopeJSON    `json:"-"`
}

// liveInputListResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputListResponseEnvelope]
type liveInputListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           LiveInputListResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputListResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LiveInputListResponseEnvelopeErrors]
type liveInputListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    liveInputListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [LiveInputListResponseEnvelopeErrorsSource]
type liveInputListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           LiveInputListResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputListResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputListResponseEnvelopeMessages]
type liveInputListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputListResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    liveInputListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputListResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [LiveInputListResponseEnvelopeMessagesSource]
type liveInputListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputListResponseEnvelopeSuccess bool

const (
	LiveInputListResponseEnvelopeSuccessTrue LiveInputListResponseEnvelopeSuccess = true
)

func (r LiveInputListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LiveInputDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type LiveInputGetResponseEnvelope struct {
	Errors   []LiveInputGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LiveInputGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LiveInputGetResponseEnvelopeSuccess `json:"success,required"`
	// Details about a live input.
	Result LiveInput                        `json:"result"`
	JSON   liveInputGetResponseEnvelopeJSON `json:"-"`
}

// liveInputGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LiveInputGetResponseEnvelope]
type liveInputGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           LiveInputGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             liveInputGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// liveInputGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LiveInputGetResponseEnvelopeErrors]
type liveInputGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    liveInputGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// liveInputGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [LiveInputGetResponseEnvelopeErrorsSource]
type liveInputGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           LiveInputGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             liveInputGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// liveInputGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LiveInputGetResponseEnvelopeMessages]
type liveInputGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LiveInputGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    liveInputGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// liveInputGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [LiveInputGetResponseEnvelopeMessagesSource]
type liveInputGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LiveInputGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveInputGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LiveInputGetResponseEnvelopeSuccess bool

const (
	LiveInputGetResponseEnvelopeSuccessTrue LiveInputGetResponseEnvelopeSuccess = true
)

func (r LiveInputGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LiveInputGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
