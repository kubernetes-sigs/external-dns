// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// V2DirectUploadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2DirectUploadService] method instead.
type V2DirectUploadService struct {
	Options []option.RequestOption
}

// NewV2DirectUploadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2DirectUploadService(opts ...option.RequestOption) (r *V2DirectUploadService) {
	r = &V2DirectUploadService{}
	r.Options = opts
	return
}

// Direct uploads allow users to upload images without API keys. A common use case
// are web apps, client-side applications, or mobile devices where users upload
// content directly to Cloudflare Images. This method creates a draft record for a
// future image. It returns an upload URL and an image identifier. To verify if the
// image itself has been uploaded, send an image details request
// (accounts/:account_identifier/images/v1/:identifier), and check that the
// `draft: true` property is not present.
func (r *V2DirectUploadService) New(ctx context.Context, params V2DirectUploadNewParams, opts ...option.RequestOption) (res *V2DirectUploadNewResponse, err error) {
	var env V2DirectUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v2/direct_upload", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V2DirectUploadNewResponse struct {
	// Image unique identifier.
	ID string `json:"id"`
	// The URL the unauthenticated upload can be performed to using a single HTTP POST
	// (multipart/form-data) request.
	UploadURL string                        `json:"uploadURL"`
	JSON      v2DirectUploadNewResponseJSON `json:"-"`
}

// v2DirectUploadNewResponseJSON contains the JSON metadata for the struct
// [V2DirectUploadNewResponse]
type v2DirectUploadNewResponseJSON struct {
	ID          apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2DirectUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2DirectUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type V2DirectUploadNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional Image Custom ID. Up to 1024 chars. Can include any number of subpaths,
	// and utf8 characters. Cannot start nor end with a / (forward slash). Cannot be a
	// UUID.
	ID param.Field[string] `json:"id"`
	// Can set the creator field with an internal user ID.
	Creator param.Field[string] `json:"creator"`
	// The date after which the upload will not be accepted. Minimum: Now + 2 minutes.
	// Maximum: Now + 6 hours.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record, for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token to be accessed.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r V2DirectUploadNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type V2DirectUploadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   V2DirectUploadNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success V2DirectUploadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v2DirectUploadNewResponseEnvelopeJSON    `json:"-"`
}

// v2DirectUploadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V2DirectUploadNewResponseEnvelope]
type v2DirectUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2DirectUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2DirectUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V2DirectUploadNewResponseEnvelopeSuccess bool

const (
	V2DirectUploadNewResponseEnvelopeSuccessTrue V2DirectUploadNewResponseEnvelopeSuccess = true
)

func (r V2DirectUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V2DirectUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
