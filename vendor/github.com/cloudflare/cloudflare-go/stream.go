package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var (
	// ErrMissingUploadURL is for when a URL is required but missing.
	ErrMissingUploadURL = errors.New("required url missing")
	// ErrMissingMaxDuration is for when MaxDuration is required but missing.
	ErrMissingMaxDuration = errors.New("required max duration missing")
	// ErrMissingVideoID is for when VideoID is required but missing.
	ErrMissingVideoID = errors.New("required video id missing")
	// ErrMissingFilePath is for when FilePath is required but missing.
	ErrMissingFilePath = errors.New("required file path missing")
)

// StreamVideo represents a stream video.
type StreamVideo struct {
	AllowedOrigins        []string                 `json:"allowedOrigins,omitempty"`
	Created               *time.Time               `json:"created,omitempty"`
	Duration              int                      `json:"duration,omitempty"`
	Input                 StreamVideoInput         `json:"input,omitempty"`
	MaxDurationSeconds    int                      `json:"maxDurationSeconds,omitempty"`
	Meta                  map[string]interface{}   `json:"meta,omitempty"`
	Modified              *time.Time               `json:"modified,omitempty"`
	UploadExpiry          *time.Time               `json:"uploadExpiry,omitempty"`
	Playback              StreamVideoPlayback      `json:"playback,omitempty"`
	Preview               string                   `json:"preview,omitempty"`
	ReadyToStream         bool                     `json:"readyToStream,omitempty"`
	RequireSignedURLs     bool                     `json:"requireSignedURLs,omitempty"`
	Size                  int                      `json:"size,omitempty"`
	Status                StreamVideoStatus        `json:"status,omitempty"`
	Thumbnail             string                   `json:"thumbnail,omitempty"`
	ThumbnailTimestampPct float64                  `json:"thumbnailTimestampPct,omitempty"`
	UID                   string                   `json:"uid,omitempty"`
	Creator               string                   `json:"creator,omitempty"`
	LiveInput             string                   `json:"liveInput,omitempty"`
	Uploaded              *time.Time               `json:"uploaded,omitempty"`
	Watermark             StreamVideoWatermark     `json:"watermark,omitempty"`
	NFT                   StreamVideoNFTParameters `json:"nft,omitempty"`
}

// StreamVideoInput represents the video input values of a stream video.
type StreamVideoInput struct {
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

// StreamVideoPlayback represents the playback URLs for a video.
type StreamVideoPlayback struct {
	HLS  string `json:"hls,omitempty"`
	Dash string `json:"dash,omitempty"`
}

// StreamVideoStatus represents the status of a stream video.
type StreamVideoStatus struct {
	State           string `json:"state,omitempty"`
	PctComplete     string `json:"pctComplete,omitempty"`
	ErrorReasonCode string `json:"errorReasonCode,omitempty"`
	ErrorReasonText string `json:"errorReasonText,omitempty"`
}

// StreamVideoWatermark represents a watermark for a stream video.
type StreamVideoWatermark struct {
	UID            string     `json:"uid,omitempty"`
	Size           int        `json:"size,omitempty"`
	Height         int        `json:"height,omitempty"`
	Width          int        `json:"width,omitempty"`
	Created        *time.Time `json:"created,omitempty"`
	DownloadedFrom string     `json:"downloadedFrom,omitempty"`
	Name           string     `json:"name,omitempty"`
	Opacity        float64    `json:"opacity,omitempty"`
	Padding        float64    `json:"padding,omitempty"`
	Scale          float64    `json:"scale,omitempty"`
	Position       string     `json:"position,omitempty"`
}

// StreamVideoNFTParameters represents a NFT for a stream video.
type StreamVideoNFTParameters struct {
	AccountID string
	VideoID   string
	Contract  string `json:"contract,omitempty"`
	Token     int    `json:"token,omitempty"`
}

// StreamUploadFromURLParameters are the parameters used when uploading a video from URL.
type StreamUploadFromURLParameters struct {
	AccountID             string
	VideoID               string
	URL                   string                  `json:"url"`
	Creator               string                  `json:"creator,omitempty"`
	ThumbnailTimestampPct float64                 `json:"thumbnailTimestampPct,omitempty"`
	AllowedOrigins        []string                `json:"allowedOrigins,omitempty"`
	RequiredSignedURLs    bool                    `json:"requiredSignedURLs,omitempty"`
	Watermark             UploadVideoURLWatermark `json:"watermark,omitempty"`
}

// StreamCreateVideoParameters are parameters used when creating a video.
type StreamCreateVideoParameters struct {
	AccountID             string
	MaxDurationSeconds    int                     `json:"maxDurationSeconds,omitempty"`
	Expiry                *time.Time              `json:"expiry,omitempty"`
	Creator               string                  `json:"creator,omitempty"`
	ThumbnailTimestampPct float64                 `json:"thumbnailTimestampPct,omitempty"`
	AllowedOrigins        []string                `json:"allowedOrigins,omitempty"`
	RequiredSignedURLs    bool                    `json:"requiredSignedURLs,omitempty"`
	Watermark             UploadVideoURLWatermark `json:"watermark,omitempty"`
}

// UploadVideoURLWatermark represents UID of an existing watermark.
type UploadVideoURLWatermark struct {
	UID string `json:"uid,omitempty"`
}

// StreamVideoCreate represents parameters returned after creating a video.
type StreamVideoCreate struct {
	UploadURL string               `json:"uploadURL,omitempty"`
	UID       string               `json:"uid,omitempty"`
	Watermark StreamVideoWatermark `json:"watermark,omitempty"`
}

// StreamParameters are the basic parameters needed.
type StreamParameters struct {
	AccountID string
	VideoID   string
}

// StreamUploadFileParameters are parameters needed for file upload of a video.
type StreamUploadFileParameters struct {
	AccountID string
	VideoID   string
	FilePath  string
}

// StreamListParameters represents parameters used when listing stream videos.
type StreamListParameters struct {
	AccountID     string
	VideoID       string
	After         *time.Time `url:"after,omitempty"`
	Before        *time.Time `url:"before,omitempty"`
	Creator       string     `url:"creator,omitempty"`
	IncludeCounts bool       `url:"include_counts,omitempty"`
	Search        string     `url:"search,omitempty"`
	Limit         int        `url:"limit,omitempty"`
	Asc           bool       `url:"asc,omitempty"`
	Status        string     `url:"status,omitempty"`
}

// StreamSignedURLParameters represent parameters used when creating a signed URL.
type StreamSignedURLParameters struct {
	AccountID    string
	VideoID      string
	ID           string             `json:"id,omitempty"`
	PEM          string             `json:"pem,omitempty"`
	EXP          int                `json:"exp,omitempty"`
	NBF          int                `json:"nbf,omitempty"`
	Downloadable bool               `json:"downloadable,omitempty"`
	AccessRules  []StreamAccessRule `json:"accessRules,omitempty"`
}

// StreamVideoResponse represents an API response of a stream video.
type StreamVideoResponse struct {
	Response
	Result StreamVideo `json:"result,omitempty"`
}

// StreamVideoCreateResponse represents an API response of creating a stream video.
type StreamVideoCreateResponse struct {
	Response
	Result StreamVideoCreate `json:"result,omitempty"`
}

// StreamListResponse represents the API response from a StreamListRequest.
type StreamListResponse struct {
	Response
	Result []StreamVideo `json:"result,omitempty"`
	Total  string        `json:"total,omitempty"`
	Range  string        `json:"range,omitempty"`
}

// StreamSignedURLResponse represents an API response for a signed URL.
type StreamSignedURLResponse struct {
	Response
	Result struct {
		Token string `json:"token,omitempty"`
	}
}

// StreamAccessRule represents the accessRules when creating a signed URL.
type StreamAccessRule struct {
	Type    string   `json:"type"`
	Country []string `json:"country,omitempty"`
	Action  string   `json:"action"`
	IP      []string `json:"ip,omitempty"`
}

// StreamUploadFromURL send a video URL to it will be downloaded and made available on Stream.
//
// API Reference: https://api.cloudflare.com/#stream-videos-upload-a-video-from-a-url
func (api *API) StreamUploadFromURL(ctx context.Context, params StreamUploadFromURLParameters) (StreamVideo, error) {
	if params.AccountID == "" {
		return StreamVideo{}, ErrMissingAccountID
	}

	if params.URL == "" {
		return StreamVideo{}, ErrMissingUploadURL
	}

	uri := fmt.Sprintf("/accounts/%s/stream/copy", params.AccountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return StreamVideo{}, err
	}

	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamUploadVideoFile uploads a video from a path to the file.
//
// API Reference: https://api.cloudflare.com/#stream-videos-upload-a-video-using-a-single-http-request
func (api *API) StreamUploadVideoFile(ctx context.Context, params StreamUploadFileParameters) (StreamVideo, error) {
	if params.AccountID == "" {
		return StreamVideo{}, ErrMissingAccountID
	}

	if params.FilePath == "" {
		return StreamVideo{}, ErrMissingFilePath
	}

	uri := fmt.Sprintf("/accounts/%s/stream", params.AccountID)

	// Create new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", params.FilePath)
	if err != nil {
		return StreamVideo{}, err
	}
	file, err := os.Open(params.FilePath)
	if err != nil {
		return StreamVideo{}, err
	}
	if _, err := io.Copy(formFile, file); err != nil {
		return StreamVideo{}, err
	}
	if err := writer.Close(); err != nil {
		return StreamVideo{}, err
	}

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPost, uri, body, http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{writer.FormDataContentType()},
	})
	if err != nil {
		return StreamVideo{}, err
	}

	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamCreateVideoDirectURL creates a video and returns an authenticated URL.
//
// API Reference: https://api.cloudflare.com/#stream-videos-create-a-video-and-get-authenticated-direct-upload-url
func (api *API) StreamCreateVideoDirectURL(ctx context.Context, params StreamCreateVideoParameters) (StreamVideoCreate, error) {
	if params.AccountID == "" {
		return StreamVideoCreate{}, ErrMissingAccountID
	}

	if params.MaxDurationSeconds == 0 {
		return StreamVideoCreate{}, ErrMissingMaxDuration
	}

	uri := fmt.Sprintf("/accounts/%s/stream/direct_upload", params.AccountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return StreamVideoCreate{}, err
	}

	var streamVideoCreateResponse StreamVideoCreateResponse
	if err := json.Unmarshal(res, &streamVideoCreateResponse); err != nil {
		return StreamVideoCreate{}, err
	}
	return streamVideoCreateResponse.Result, nil
}

// StreamListVideos list videos currently in stream.
//
// API reference: https://api.cloudflare.com/#stream-videos-list-videos
func (api *API) StreamListVideos(ctx context.Context, params StreamListParameters) ([]StreamVideo, error) {
	if params.AccountID == "" {
		return []StreamVideo{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/stream", params.AccountID), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []StreamVideo{}, err
	}

	var streamListResponse StreamListResponse
	if err := json.Unmarshal(res, &streamListResponse); err != nil {
		return []StreamVideo{}, err
	}
	return streamListResponse.Result, nil
}

// Skipped: https://api.cloudflare.com/#stream-videos-initiate-a-video-upload-using-tus

// StreamGetVideo gets the details for a specific video.
//
// API Reference: https://api.cloudflare.com/#stream-videos-video-details
func (api *API) StreamGetVideo(ctx context.Context, options StreamParameters) (StreamVideo, error) {
	if options.AccountID == "" {
		return StreamVideo{}, ErrMissingAccountID
	}

	if options.VideoID == "" {
		return StreamVideo{}, ErrMissingVideoID
	}

	uri := fmt.Sprintf("/accounts/%s/stream/%s", options.AccountID, options.VideoID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return StreamVideo{}, err
	}
	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamEmbedHTML gets an HTML fragment to embed on a web page.
//
// API Reference: https://api.cloudflare.com/#stream-videos-embed-code-html
func (api *API) StreamEmbedHTML(ctx context.Context, options StreamParameters) (string, error) {
	if options.AccountID == "" {
		return "", ErrMissingAccountID
	}

	if options.VideoID == "" {
		return "", ErrMissingVideoID
	}

	uri := fmt.Sprintf("/accounts/%s/stream/%s/embed", options.AccountID, options.VideoID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return "", err
	}
	return string(res), nil
}

// StreamDeleteVideo deletes a video.
//
// API Reference: https://api.cloudflare.com/#stream-videos-delete-video
func (api *API) StreamDeleteVideo(ctx context.Context, options StreamParameters) error {
	if options.AccountID == "" {
		return ErrMissingAccountID
	}

	if options.VideoID == "" {
		return ErrMissingVideoID
	}

	uri := fmt.Sprintf("/accounts/%s/stream/%s", options.AccountID, options.VideoID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// StreamAssociateNFT associates a video to a token and contract address.
//
// API Reference: https://api.cloudflare.com/#stream-videos-associate-video-to-an-nft
func (api *API) StreamAssociateNFT(ctx context.Context, options StreamVideoNFTParameters) (StreamVideo, error) {
	if options.AccountID == "" {
		return StreamVideo{}, ErrMissingAccountID
	}

	if options.VideoID == "" {
		return StreamVideo{}, ErrMissingVideoID
	}

	uri := fmt.Sprintf("/accounts/%s/stream/%s", options.AccountID, options.VideoID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return StreamVideo{}, err
	}
	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamCreateSignedURL creates a signed URL token for a video.
//
// API Reference: https://api.cloudflare.com/#stream-videos-associate-video-to-an-nft
func (api *API) StreamCreateSignedURL(ctx context.Context, params StreamSignedURLParameters) (string, error) {
	if params.AccountID == "" {
		return "", ErrMissingAccountID
	}
	if params.VideoID == "" {
		return "", ErrMissingVideoID
	}

	uri := fmt.Sprintf("/accounts/%s/stream/%s/token", params.AccountID, params.VideoID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)

	if err != nil {
		return "", err
	}
	var streamSignedResponse StreamSignedURLResponse
	if err := json.Unmarshal(res, &streamSignedResponse); err != nil {
		return "", err
	}
	return streamSignedResponse.Result.Token, nil
}
