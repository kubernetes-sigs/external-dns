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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// StreamService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
	Options      []option.RequestOption
	AudioTracks  *AudioTrackService
	Videos       *VideoService
	Clip         *ClipService
	Copy         *CopyService
	DirectUpload *DirectUploadService
	Keys         *KeyService
	LiveInputs   *LiveInputService
	Watermarks   *WatermarkService
	Webhooks     *WebhookService
	Captions     *CaptionService
	Downloads    *DownloadService
	Embed        *EmbedService
	Token        *TokenService
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r *StreamService) {
	r = &StreamService{}
	r.Options = opts
	r.AudioTracks = NewAudioTrackService(opts...)
	r.Videos = NewVideoService(opts...)
	r.Clip = NewClipService(opts...)
	r.Copy = NewCopyService(opts...)
	r.DirectUpload = NewDirectUploadService(opts...)
	r.Keys = NewKeyService(opts...)
	r.LiveInputs = NewLiveInputService(opts...)
	r.Watermarks = NewWatermarkService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.Captions = NewCaptionService(opts...)
	r.Downloads = NewDownloadService(opts...)
	r.Embed = NewEmbedService(opts...)
	r.Token = NewTokenService(opts...)
	return
}

// Initiates a video upload using the TUS protocol. On success, the server responds
// with a status code 201 (created) and includes a `location` header to indicate
// where the content should be uploaded. Refer to https://tus.io for protocol
// details.
func (r *StreamService) New(ctx context.Context, params StreamNewParams, opts ...option.RequestOption) (err error) {
	if params.TusResumable.Present {
		opts = append(opts, option.WithHeader("Tus-Resumable", fmt.Sprintf("%s", params.TusResumable)))
	}
	if params.UploadLength.Present {
		opts = append(opts, option.WithHeader("Upload-Length", fmt.Sprintf("%s", params.UploadLength)))
	}
	if params.UploadCreator.Present {
		opts = append(opts, option.WithHeader("Upload-Creator", fmt.Sprintf("%s", params.UploadCreator)))
	}
	if params.UploadMetadata.Present {
		opts = append(opts, option.WithHeader("Upload-Metadata", fmt.Sprintf("%s", params.UploadMetadata)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Lists up to 1000 videos from a single request. For a specific range, refer to
// the optional parameters.
func (r *StreamService) List(ctx context.Context, params StreamListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Video], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream", params.AccountID)
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

// Lists up to 1000 videos from a single request. For a specific range, refer to
// the optional parameters.
func (r *StreamService) ListAutoPaging(ctx context.Context, params StreamListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Video] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a video and its copies from Cloudflare Stream.
func (r *StreamService) Delete(ctx context.Context, identifier string, body StreamDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Edit details for a single video.
func (r *StreamService) Edit(ctx context.Context, identifier string, params StreamEditParams, opts ...option.RequestOption) (res *Video, err error) {
	var env StreamEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", params.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single video.
func (r *StreamService) Get(ctx context.Context, identifier string, query StreamGetParams, opts ...option.RequestOption) (res *Video, err error) {
	var env StreamGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AllowedOrigins = string

type AllowedOriginsParam = string

type Video struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []AllowedOrigins `json:"allowedOrigins"`
	// The date and time the media item was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// The duration of the video in seconds. A value of `-1` means the duration is
	// unknown. The duration becomes available after the upload and before the video is
	// ready.
	Duration float64    `json:"duration"`
	Input    VideoInput `json:"input"`
	// The live input ID used to upload a video with Stream Live.
	LiveInput string `json:"liveInput"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the media item was last modified.
	Modified time.Time     `json:"modified" format:"date-time"`
	Playback VideoPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video is playable. The field is empty if the video is not
	// ready for viewing or the live stream is still in progress.
	ReadyToStream bool `json:"readyToStream"`
	// Indicates the time at which the video became playable. The field is empty if the
	// video is not ready for viewing or the live stream is still in progress.
	ReadyToStreamAt time.Time `json:"readyToStreamAt" format:"date-time"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion time.Time `json:"scheduledDeletion" format:"date-time"`
	// The size of the media item in bytes.
	Size float64 `json:"size"`
	// Specifies a detailed status for a video. If the `state` is `inprogress` or
	// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
	// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
	// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
	// and `errorReasonText` provide additional details.
	Status VideoStatus `json:"status"`
	// The media item's thumbnail URI. This field is omitted until encoding is
	// complete.
	Thumbnail string `json:"thumbnail" format:"uri"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64 `json:"thumbnailTimestampPct"`
	// A Cloudflare-generated unique identifier for a media item.
	UID string `json:"uid"`
	// The date and time the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry time.Time `json:"uploadExpiry" format:"date-time"`
	Watermark    Watermark `json:"watermark"`
	JSON         videoJSON `json:"-"`
}

// videoJSON contains the JSON metadata for the struct [Video]
type videoJSON struct {
	AllowedOrigins        apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	Duration              apijson.Field
	Input                 apijson.Field
	LiveInput             apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	ReadyToStream         apijson.Field
	ReadyToStreamAt       apijson.Field
	RequireSignedURLs     apijson.Field
	ScheduledDeletion     apijson.Field
	Size                  apijson.Field
	Status                apijson.Field
	Thumbnail             apijson.Field
	ThumbnailTimestampPct apijson.Field
	UID                   apijson.Field
	Uploaded              apijson.Field
	UploadExpiry          apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Video) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoJSON) RawJSON() string {
	return r.raw
}

type VideoInput struct {
	// The video height in pixels. A value of `-1` means the height is unknown. The
	// value becomes available after the upload and before the video is ready.
	Height int64 `json:"height"`
	// The video width in pixels. A value of `-1` means the width is unknown. The value
	// becomes available after the upload and before the video is ready.
	Width int64          `json:"width"`
	JSON  videoInputJSON `json:"-"`
}

// videoInputJSON contains the JSON metadata for the struct [VideoInput]
type videoInputJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoInputJSON) RawJSON() string {
	return r.raw
}

type VideoPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string            `json:"hls"`
	JSON videoPlaybackJSON `json:"-"`
}

// videoPlaybackJSON contains the JSON metadata for the struct [VideoPlayback]
type videoPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoPlaybackJSON) RawJSON() string {
	return r.raw
}

// Specifies a detailed status for a video. If the `state` is `inprogress` or
// `error`, the `step` field returns `encoding` or `manifest`. If the `state` is
// `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the
// approximate percent of completion. If the `state` is `error`, `errorReasonCode`
// and `errorReasonText` provide additional details.
type VideoStatus struct {
	// Specifies why the video failed to encode. This field is empty if the video is
	// not in an `error` state. Preferred for programmatic use.
	ErrorReasonCode string `json:"errorReasonCode"`
	// Specifies why the video failed to encode using a human readable error message in
	// English. This field is empty if the video is not in an `error` state.
	ErrorReasonText string `json:"errorReasonText"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	PctComplete string `json:"pctComplete"`
	// Specifies the processing status for all quality levels for a video.
	State VideoStatusState `json:"state"`
	JSON  videoStatusJSON  `json:"-"`
}

// videoStatusJSON contains the JSON metadata for the struct [VideoStatus]
type videoStatusJSON struct {
	ErrorReasonCode apijson.Field
	ErrorReasonText apijson.Field
	PctComplete     apijson.Field
	State           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VideoStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoStatusJSON) RawJSON() string {
	return r.raw
}

// Specifies the processing status for all quality levels for a video.
type VideoStatusState string

const (
	VideoStatusStatePendingupload  VideoStatusState = "pendingupload"
	VideoStatusStateDownloading    VideoStatusState = "downloading"
	VideoStatusStateQueued         VideoStatusState = "queued"
	VideoStatusStateInprogress     VideoStatusState = "inprogress"
	VideoStatusStateReady          VideoStatusState = "ready"
	VideoStatusStateError          VideoStatusState = "error"
	VideoStatusStateLiveInprogress VideoStatusState = "live-inprogress"
)

func (r VideoStatusState) IsKnown() bool {
	switch r {
	case VideoStatusStatePendingupload, VideoStatusStateDownloading, VideoStatusStateQueued, VideoStatusStateInprogress, VideoStatusStateReady, VideoStatusStateError, VideoStatusStateLiveInprogress:
		return true
	}
	return false
}

type StreamNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
	// Specifies the TUS protocol version. This value must be included in every upload
	// request. Notes: The only supported version of TUS protocol is 1.0.0.
	TusResumable param.Field[StreamNewParamsTusResumable] `header:"Tus-Resumable,required"`
	// Indicates the size of the entire upload in bytes. The value must be a
	// non-negative integer.
	UploadLength param.Field[int64] `header:"Upload-Length,required"`
	// Provisions a URL to let your end users upload videos directly to Cloudflare
	// Stream without exposing your API token to clients.
	DirectUser param.Field[bool] `query:"direct_user"`
	// A user-defined identifier for the media creator.
	UploadCreator param.Field[string] `header:"Upload-Creator"`
	// Comma-separated key-value pairs following the TUS protocol specification. Values
	// are Base-64 encoded. Supported keys: `name`, `requiresignedurls`,
	// `allowedorigins`, `thumbnailtimestamppct`, `watermark`, `scheduleddeletion`,
	// `maxdurationseconds`.
	UploadMetadata param.Field[string] `header:"Upload-Metadata"`
}

func (r StreamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [StreamNewParams]'s query parameters as `url.Values`.
func (r StreamNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the TUS protocol version. This value must be included in every upload
// request. Notes: The only supported version of TUS protocol is 1.0.0.
type StreamNewParamsTusResumable string

const (
	StreamNewParamsTusResumable1_0_0 StreamNewParamsTusResumable = "1.0.0"
)

func (r StreamNewParamsTusResumable) IsKnown() bool {
	switch r {
	case StreamNewParamsTusResumable1_0_0:
		return true
	}
	return false
}

type StreamListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Lists videos in ascending order of creation.
	Asc param.Field[bool] `query:"asc"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `query:"creator"`
	// Lists videos created before the specified date.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Includes the total number of videos associated with the submitted query
	// parameters.
	IncludeCounts param.Field[bool] `query:"include_counts"`
	// Provides a partial word match of the `name` key in the `meta` field. Slow for
	// medium to large video libraries. May be unavailable for very large libraries.
	Search param.Field[string] `query:"search"`
	// Lists videos created after the specified date.
	Start param.Field[time.Time] `query:"start" format:"date-time"`
	// Specifies the processing status for all quality levels for a video.
	Status param.Field[StreamListParamsStatus] `query:"status"`
	// Specifies whether the video is `vod` or `live`.
	Type param.Field[string] `query:"type"`
	// Provides a fast, exact string match on the `name` key in the `meta` field.
	VideoName param.Field[string] `query:"video_name"`
}

// URLQuery serializes [StreamListParams]'s query parameters as `url.Values`.
func (r StreamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the processing status for all quality levels for a video.
type StreamListParamsStatus string

const (
	StreamListParamsStatusPendingupload  StreamListParamsStatus = "pendingupload"
	StreamListParamsStatusDownloading    StreamListParamsStatus = "downloading"
	StreamListParamsStatusQueued         StreamListParamsStatus = "queued"
	StreamListParamsStatusInprogress     StreamListParamsStatus = "inprogress"
	StreamListParamsStatusReady          StreamListParamsStatus = "ready"
	StreamListParamsStatusError          StreamListParamsStatus = "error"
	StreamListParamsStatusLiveInprogress StreamListParamsStatus = "live-inprogress"
)

func (r StreamListParamsStatus) IsKnown() bool {
	switch r {
	case StreamListParamsStatusPendingupload, StreamListParamsStatusDownloading, StreamListParamsStatusQueued, StreamListParamsStatusInprogress, StreamListParamsStatusReady, StreamListParamsStatusError, StreamListParamsStatusLiveInprogress:
		return true
	}
	return false
}

type StreamDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamEditParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]AllowedOriginsParam] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta param.Field[interface{}] `json:"meta"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// Indicates the date and time at which the video will be deleted. Omit the field
	// to indicate no change, or include with a `null` value to remove an existing
	// scheduled deletion. If specified, must be at least 30 days from upload time.
	ScheduledDeletion param.Field[time.Time] `json:"scheduledDeletion" format:"date-time"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64] `json:"thumbnailTimestampPct"`
	// The date and time when the video upload URL is no longer valid for direct user
	// uploads.
	UploadExpiry param.Field[time.Time] `json:"uploadExpiry" format:"date-time"`
}

func (r StreamEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamEditResponseEnvelope struct {
	Errors   []StreamEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success StreamEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Video                             `json:"result"`
	JSON    streamEditResponseEnvelopeJSON    `json:"-"`
}

// streamEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamEditResponseEnvelope]
type streamEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamEditResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           StreamEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             streamEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// streamEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamEditResponseEnvelopeErrors]
type streamEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamEditResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    streamEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// streamEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [StreamEditResponseEnvelopeErrorsSource]
type streamEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StreamEditResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           StreamEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             streamEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// streamEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [StreamEditResponseEnvelopeMessages]
type streamEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamEditResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    streamEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// streamEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [StreamEditResponseEnvelopeMessagesSource]
type streamEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StreamEditResponseEnvelopeSuccess bool

const (
	StreamEditResponseEnvelopeSuccessTrue StreamEditResponseEnvelopeSuccess = true
)

func (r StreamEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StreamEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type StreamGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamGetResponseEnvelope struct {
	Errors   []StreamGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success StreamGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Video                            `json:"result"`
	JSON    streamGetResponseEnvelopeJSON    `json:"-"`
}

// streamGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelope]
type streamGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           StreamGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             streamGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// streamGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelopeErrors]
type streamGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    streamGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// streamGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [StreamGetResponseEnvelopeErrorsSource]
type streamGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           StreamGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             streamGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// streamGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelopeMessages]
type streamGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    streamGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// streamGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [StreamGetResponseEnvelopeMessagesSource]
type streamGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type StreamGetResponseEnvelopeSuccess bool

const (
	StreamGetResponseEnvelopeSuccessTrue StreamGetResponseEnvelopeSuccess = true
)

func (r StreamGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case StreamGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
