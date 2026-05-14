// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// V1Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options  []option.RequestOption
	Keys     *V1KeyService
	Stats    *V1StatService
	Variants *V1VariantService
	Blobs    *V1BlobService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r *V1Service) {
	r = &V1Service{}
	r.Options = opts
	r.Keys = NewV1KeyService(opts...)
	r.Stats = NewV1StatService(opts...)
	r.Variants = NewV1VariantService(opts...)
	r.Blobs = NewV1BlobService(opts...)
	return
}

// Upload an image with up to 10 Megabytes using a single HTTP POST
// (multipart/form-data) request. An image can be uploaded by sending an image file
// or passing an accessible to an API url.
func (r *V1Service) New(ctx context.Context, params V1NewParams, opts ...option.RequestOption) (res *Image, err error) {
	var env V1NewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
//
// Deprecated: deprecated
func (r *V1Service) List(ctx context.Context, params V1ListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[V1ListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1", params.AccountID)
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

// List up to 100 images with one request. Use the optional parameters below to get
// a specific range of images.
//
// Deprecated: deprecated
func (r *V1Service) ListAutoPaging(ctx context.Context, params V1ListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[V1ListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete an image on Cloudflare Images. On success, all copies of the image are
// deleted and purged from cache.
func (r *V1Service) Delete(ctx context.Context, imageID string, body V1DeleteParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env V1DeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", body.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update image access control. On access control change, all copies of the image
// are purged from cache.
func (r *V1Service) Edit(ctx context.Context, imageID string, params V1EditParams, opts ...option.RequestOption) (res *Image, err error) {
	var env V1EditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", params.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single image.
func (r *V1Service) Get(ctx context.Context, imageID string, query V1GetParams, opts ...option.RequestOption) (res *Image, err error) {
	var env V1GetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s", query.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Image struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Can set the creator field with an internal user ID.
	Creator string `json:"creator,nullable"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []string  `json:"variants" format:"uri"`
	JSON     imageJSON `json:"-"`
}

// imageJSON contains the JSON metadata for the struct [Image]
type imageJSON struct {
	ID                apijson.Field
	Creator           apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Image) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageJSON) RawJSON() string {
	return r.raw
}

type V1ListResponse struct {
	Images []Image            `json:"images"`
	JSON   v1ListResponseJSON `json:"-"`
}

// v1ListResponseJSON contains the JSON metadata for the struct [V1ListResponse]
type v1ListResponseJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ListResponseJSON) RawJSON() string {
	return r.raw
}

type V1NewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional custom unique identifier for your image.
	ID param.Field[string] `json:"id"`
	// Can set the creator field with an internal user ID.
	Creator param.Field[string] `json:"creator"`
	// An image binary data. Only needed when type is uploading a file.
	File param.Field[io.Reader] `json:"file" format:"binary"`
	// User modifiable key-value store. Can use used for keeping references to another
	// system of record for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token for the access.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// A URL to fetch an image from origin. Only needed when type is uploading from a
	// URL.
	URL param.Field[string] `json:"url"`
}

func (r V1NewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type V1NewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Image                 `json:"result,required"`
	// Whether the API call was successful
	Success V1NewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1NewResponseEnvelopeJSON    `json:"-"`
}

// v1NewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1NewResponseEnvelope]
type v1NewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1NewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1NewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1NewResponseEnvelopeSuccess bool

const (
	V1NewResponseEnvelopeSuccessTrue V1NewResponseEnvelopeSuccess = true
)

func (r V1NewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1NewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1ListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Internal user ID set within the creator field. Setting to empty string "" will
	// return images where creator field is not set
	Creator param.Field[string] `query:"creator"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [V1ListParams]'s query parameters as `url.Values`.
func (r V1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type V1DeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1DeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Whether the API call was successful
	Success V1DeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1DeleteResponseEnvelopeJSON    `json:"-"`
}

// v1DeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1DeleteResponseEnvelope]
type v1DeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1DeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1DeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1DeleteResponseEnvelopeSuccess bool

const (
	V1DeleteResponseEnvelopeSuccessTrue V1DeleteResponseEnvelopeSuccess = true
)

func (r V1DeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1DeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1EditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Can set the creator field with an internal user ID.
	Creator param.Field[string] `json:"creator"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. No change if not specified.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image can be accessed using only its UID. If set to
	// `true`, a signed token needs to be generated with a signing key to view the
	// image. Returns a new UID on a change. No change if not specified.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r V1EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1EditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Image                 `json:"result,required"`
	// Whether the API call was successful
	Success V1EditResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1EditResponseEnvelopeJSON    `json:"-"`
}

// v1EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1EditResponseEnvelope]
type v1EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1EditResponseEnvelopeSuccess bool

const (
	V1EditResponseEnvelopeSuccessTrue V1EditResponseEnvelopeSuccess = true
)

func (r V1EditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1EditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1GetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1GetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Image                 `json:"result,required"`
	// Whether the API call was successful
	Success V1GetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1GetResponseEnvelopeJSON    `json:"-"`
}

// v1GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1GetResponseEnvelope]
type v1GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1GetResponseEnvelopeSuccess bool

const (
	V1GetResponseEnvelopeSuccessTrue V1GetResponseEnvelopeSuccess = true
)

func (r V1GetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1GetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
