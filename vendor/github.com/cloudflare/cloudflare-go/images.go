package cloudflare

import (
	"bytes"
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"errors"
)

// Image represents a Cloudflare Image.
type Image struct {
	ID                string                 `json:"id"`
	Filename          string                 `json:"filename"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	RequireSignedURLs bool                   `json:"requireSignedURLs"`
	Variants          []string               `json:"variants"`
	Uploaded          time.Time              `json:"uploaded"`
}

// ImageUploadRequest is the data required for an Image Upload request.
type ImageUploadRequest struct {
	File              io.ReadCloser
	Name              string
	RequireSignedURLs bool
	Metadata          map[string]interface{}
}

// write writes the image upload data to a multipart writer, so
// it can be used in an HTTP request.
func (b ImageUploadRequest) write(mpw *multipart.Writer) error {
	if b.File == nil {
		return errors.New("a file to upload must be specified")
	}
	name := b.Name
	part, err := mpw.CreateFormFile("file", name)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, b.File)
	if err != nil {
		_ = b.File.Close()
		return err
	}
	_ = b.File.Close()

	// According to the Cloudflare docs, this field defaults to false.
	// For simplicity, we will only send it if the value is true, however
	// if the default is changed to true, this logic will need to be updated.
	if b.RequireSignedURLs {
		err = mpw.WriteField("requireSignedURLs", "true")
		if err != nil {
			return err
		}
	}

	if b.Metadata != nil {
		part, err = mpw.CreateFormField("metadata")
		if err != nil {
			return err
		}
		err = json.NewEncoder(part).Encode(b.Metadata)
		if err != nil {
			return err
		}
	}

	return nil
}

// ImageUpdateRequest is the data required for an UpdateImage request.
type ImageUpdateRequest struct {
	RequireSignedURLs bool                   `json:"requireSignedURLs"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

// ImageDirectUploadURLRequest is the data required for a CreateImageDirectUploadURL request.
type ImageDirectUploadURLRequest struct {
	Expiry time.Time `json:"expiry"`
}

// ImageDirectUploadURLResponse is the API response for a direct image upload url.
type ImageDirectUploadURLResponse struct {
	Result ImageDirectUploadURL `json:"result"`
	Response
}

// ImageDirectUploadURL .
type ImageDirectUploadURL struct {
	ID        string `json:"id"`
	UploadURL string `json:"uploadURL"`
}

// ImagesListResponse is the API response for listing all images.
type ImagesListResponse struct {
	Result struct {
		Images []Image `json:"images"`
	} `json:"result"`
	Response
}

// ImageDetailsResponse is the API response for getting an image's details.
type ImageDetailsResponse struct {
	Result Image `json:"result"`
	Response
}

// ImagesStatsResponse is the API response for image stats.
type ImagesStatsResponse struct {
	Result struct {
		Count ImagesStatsCount `json:"count"`
	} `json:"result"`
	Response
}

// ImagesStatsCount is the stats attached to a ImagesStatsResponse.
type ImagesStatsCount struct {
	Current int64 `json:"current"`
	Allowed int64 `json:"allowed"`
}

// UploadImage uploads a single image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-upload-an-image-using-a-single-http-request
func (api *API) UploadImage(ctx context.Context, accountID string, upload ImageUploadRequest) (Image, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1", accountID)

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if err := upload.write(w); err != nil {
		_ = w.Close()
		return Image{}, fmt.Errorf("error writing multipart body: %w", err)
	}
	_ = w.Close()

	res, err := api.makeRequestContextWithHeaders(
		ctx,
		http.MethodPost,
		uri,
		body,
		http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{w.FormDataContentType()},
		},
	)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

// UpdateImage updates an existing image's metadata.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-update-image
func (api *API) UpdateImage(ctx context.Context, accountID string, id string, image ImageUpdateRequest) (Image, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", accountID, id)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, image)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

// CreateImageDirectUploadURL creates an authenticated direct upload url.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-create-authenticated-direct-upload-url
func (api *API) CreateImageDirectUploadURL(ctx context.Context, accountID string, params ImageDirectUploadURLRequest) (ImageDirectUploadURL, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/direct_upload", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return ImageDirectUploadURL{}, err
	}

	var imageDirectUploadURLResponse ImageDirectUploadURLResponse
	err = json.Unmarshal(res, &imageDirectUploadURLResponse)
	if err != nil {
		return ImageDirectUploadURL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDirectUploadURLResponse.Result, nil
}

// ListImages lists all images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-list-images
func (api *API) ListImages(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]Image, error) {
	uri := buildURI(fmt.Sprintf("/accounts/%s/images/v1", accountID), pageOpts)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Image{}, err
	}

	var imagesListResponse ImagesListResponse
	err = json.Unmarshal(res, &imagesListResponse)
	if err != nil {
		return []Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imagesListResponse.Result.Images, nil
}

// ImageDetails gets the details of an uploaded image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-image-details
func (api *API) ImageDetails(ctx context.Context, accountID string, id string) (Image, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", accountID, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

// BaseImage gets the base image used to derive variants.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-base-image
func (api *API) BaseImage(ctx context.Context, accountID string, id string) ([]byte, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/%s/blob", accountID, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteImage deletes an image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-delete-image
func (api *API) DeleteImage(ctx context.Context, accountID string, id string) error {
	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", accountID, id)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

// ImagesStats gets an account's statistics for Cloudflare Images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-images-usage-statistics
func (api *API) ImagesStats(ctx context.Context, accountID string) (ImagesStatsCount, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/stats", accountID)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrInvalidImagesAPIVersion = errors.New("invalid images API version")
	ErrMissingImageID          = errors.New("required image ID missing")
)

type ImagesAPIVersion string

const (
	ImagesAPIVersionV1 ImagesAPIVersion = "v1"
	ImagesAPIVersionV2 ImagesAPIVersion = "v2"
)

// Image represents a Cloudflare Image.
type Image struct {
	ID                string                 `json:"id"`
	Filename          string                 `json:"filename"`
	Meta              map[string]interface{} `json:"meta,omitempty"`
	RequireSignedURLs bool                   `json:"requireSignedURLs"`
	Variants          []string               `json:"variants"`
	Uploaded          time.Time              `json:"uploaded"`
}

// UploadImageParams is the data required for an Image Upload request.
type UploadImageParams struct {
	File              io.ReadCloser
	URL               string
	Name              string
	RequireSignedURLs bool
	Metadata          map[string]interface{}
}

// write writes the image upload data to a multipart writer, so
// it can be used in an HTTP request.
func (b UploadImageParams) write(mpw *multipart.Writer) error {
	if b.File == nil && b.URL == "" {
		return errors.New("a file or url to upload must be specified")
	}

	if b.File != nil {
		name := b.Name
		part, err := mpw.CreateFormFile("file", name)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, b.File)
		if err != nil {
			_ = b.File.Close()
			return err
		}
		_ = b.File.Close()
	}

	if b.URL != "" {
		err := mpw.WriteField("url", b.URL)
		if err != nil {
			return err
		}
	}

	// According to the Cloudflare docs, this field defaults to false.
	// For simplicity, we will only send it if the value is true, however
	// if the default is changed to true, this logic will need to be updated.
	if b.RequireSignedURLs {
		err := mpw.WriteField("requireSignedURLs", "true")
		if err != nil {
			return err
		}
	}

	if b.Metadata != nil {
		part, err := mpw.CreateFormField("metadata")
		if err != nil {
			return err
		}
		err = json.NewEncoder(part).Encode(b.Metadata)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateImageParams is the data required for an UpdateImage request.
type UpdateImageParams struct {
	ID                string                 `json:"-"`
	RequireSignedURLs bool                   `json:"requireSignedURLs"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

// CreateImageDirectUploadURLParams is the data required for a CreateImageDirectUploadURL request.
type CreateImageDirectUploadURLParams struct {
	Version           ImagesAPIVersion       `json:"-"`
	Expiry            *time.Time             `json:"expiry,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	RequireSignedURLs *bool                  `json:"requireSignedURLs,omitempty"`
}

// ImageDirectUploadURLResponse is the API response for a direct image upload url.
type ImageDirectUploadURLResponse struct {
	Result ImageDirectUploadURL `json:"result"`
	Response
}

// ImageDirectUploadURL .
type ImageDirectUploadURL struct {
	ID        string `json:"id"`
	UploadURL string `json:"uploadURL"`
}

// ImagesListResponse is the API response for listing all images.
type ImagesListResponse struct {
	Result struct {
		Images []Image `json:"images"`
	} `json:"result"`
	Response
}

// ImageDetailsResponse is the API response for getting an image's details.
type ImageDetailsResponse struct {
	Result Image `json:"result"`
	Response
}

// ImagesStatsResponse is the API response for image stats.
type ImagesStatsResponse struct {
	Result struct {
		Count ImagesStatsCount `json:"count"`
	} `json:"result"`
	Response
}

// ImagesStatsCount is the stats attached to a ImagesStatsResponse.
type ImagesStatsCount struct {
	Current int64 `json:"current"`
	Allowed int64 `json:"allowed"`
}

type ListImagesParams struct {
	ResultInfo
}

// UploadImage uploads a single image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-upload-an-image-using-a-single-http-request
func (api *API) UploadImage(ctx context.Context, rc *ResourceContainer, params UploadImageParams) (Image, error) {
	if rc.Level != AccountRouteLevel {
		return Image{}, ErrRequiredAccountLevelResourceContainer
	}

	if params.File != nil && params.URL != "" {
		return Image{}, errors.New("file and url uploads are mutually exclusive and can only be performed individually")
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1", rc.Identifier)

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if err := params.write(w); err != nil {
		_ = w.Close()
		return Image{}, fmt.Errorf("error writing multipart body: %w", err)
	}
	_ = w.Close()

	res, err := api.makeRequestContextWithHeaders(
		ctx,
		http.MethodPost,
		uri,
		body,
		http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{w.FormDataContentType()},
		},
	)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

// UpdateImage updates an existing image's metadata.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-update-image
func (api *API) UpdateImage(ctx context.Context, rc *ResourceContainer, params UpdateImageParams) (Image, error) {
	if rc.Level != AccountRouteLevel {
		return Image{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", rc.Identifier, params.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

var imagesMultipartBoundary = "----CloudflareImagesGoClientBoundary"

// CreateImageDirectUploadURL creates an authenticated direct upload url.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-create-authenticated-direct-upload-url
func (api *API) CreateImageDirectUploadURL(ctx context.Context, rc *ResourceContainer, params CreateImageDirectUploadURLParams) (ImageDirectUploadURL, error) {
	if rc.Level != AccountRouteLevel {
		return ImageDirectUploadURL{}, ErrRequiredAccountLevelResourceContainer
	}

	if params.Version != "" && params.Version != ImagesAPIVersionV1 && params.Version != ImagesAPIVersionV2 {
		return ImageDirectUploadURL{}, ErrInvalidImagesAPIVersion
	}

	var err error
	var uri string
	var res []byte
	switch params.Version {
	case ImagesAPIVersionV2:
		uri = fmt.Sprintf("/%s/%s/images/%s/direct_upload", rc.Level, rc.Identifier, params.Version)
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		if err := writer.SetBoundary(imagesMultipartBoundary); err != nil {
			return ImageDirectUploadURL{}, fmt.Errorf("error setting multipart boundary")
		}

		if *params.RequireSignedURLs {
			if err = writer.WriteField("requireSignedURLs", "true"); err != nil {
				return ImageDirectUploadURL{}, fmt.Errorf("error writing requireSignedURLs field: %w", err)
			}
		}
		if !params.Expiry.IsZero() {
			if err = writer.WriteField("expiry", params.Expiry.Format(time.RFC3339)); err != nil {
				return ImageDirectUploadURL{}, fmt.Errorf("error writing expiry field: %w", err)
			}
		}
		if params.Metadata != nil {
			var metadataBytes []byte
			if metadataBytes, err = json.Marshal(params.Metadata); err != nil {
				return ImageDirectUploadURL{}, fmt.Errorf("error marshalling metadata to JSON: %w", err)
			}
			if err = writer.WriteField("metadata", string(metadataBytes)); err != nil {
				return ImageDirectUploadURL{}, fmt.Errorf("error writing metadata field: %w", err)
			}
		}
		if err = writer.Close(); err != nil {
			return ImageDirectUploadURL{}, fmt.Errorf("error closing multipart writer: %w", err)
		}

		res, err = api.makeRequestContextWithHeaders(
			ctx,
			http.MethodPost,
			uri,
			body,
			http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{writer.FormDataContentType()},
			},
		)
	case ImagesAPIVersionV1:
	case "":
		uri = fmt.Sprintf("/%s/%s/images/%s/direct_upload", rc.Level, rc.Identifier, ImagesAPIVersionV1)
		res, err = api.makeRequestContext(ctx, http.MethodPost, uri, params)
	default:
		return ImageDirectUploadURL{}, ErrInvalidImagesAPIVersion
	}

	if err != nil {
		return ImageDirectUploadURL{}, err
	}

	var imageDirectUploadURLResponse ImageDirectUploadURLResponse
	err = json.Unmarshal(res, &imageDirectUploadURLResponse)
	if err != nil {
		return ImageDirectUploadURL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDirectUploadURLResponse.Result, nil
}

// ListImages lists all images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-list-images
func (api *API) ListImages(ctx context.Context, rc *ResourceContainer, params ListImagesParams) ([]Image, error) {
	uri := buildURI(fmt.Sprintf("/accounts/%s/images/v1", rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Image{}, err
	}

	var imagesListResponse ImagesListResponse
	err = json.Unmarshal(res, &imagesListResponse)
	if err != nil {
		return []Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imagesListResponse.Result.Images, nil
}

// GetImage gets the details of an uploaded image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-image-details
func (api *API) GetImage(ctx context.Context, rc *ResourceContainer, id string) (Image, error) {
	if rc.Level != AccountRouteLevel {
		return Image{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", rc.Identifier, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Image{}, err
	}

	var imageDetailsResponse ImageDetailsResponse
	err = json.Unmarshal(res, &imageDetailsResponse)
	if err != nil {
		return Image{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imageDetailsResponse.Result, nil
}

// GetBaseImage gets the base image used to derive variants.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-base-image
func (api *API) GetBaseImage(ctx context.Context, rc *ResourceContainer, id string) ([]byte, error) {
	if rc.Level != AccountRouteLevel {
		return []byte{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1/%s/blob", rc.Identifier, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteImage deletes an image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-delete-image
func (api *API) DeleteImage(ctx context.Context, rc *ResourceContainer, id string) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1/%s", rc.Identifier, id)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetImagesStats gets an account's statistics for Cloudflare Images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-images-usage-statistics
func (api *API) GetImagesStats(ctx context.Context, rc *ResourceContainer) (ImagesStatsCount, error) {
	if rc.Level != AccountRouteLevel {
		return ImagesStatsCount{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1/stats", rc.Identifier)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ImagesStatsCount{}, err
	}

	var imagesStatsResponse ImagesStatsResponse
	err = json.Unmarshal(res, &imagesStatsResponse)
	if err != nil {
		return ImagesStatsCount{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return imagesStatsResponse.Result.Count, nil
}
