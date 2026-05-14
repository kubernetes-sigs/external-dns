// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// V1BlobService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BlobService] method instead.
type V1BlobService struct {
	Options []option.RequestOption
}

// NewV1BlobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1BlobService(opts ...option.RequestOption) (r *V1BlobService) {
	r = &V1BlobService{}
	r.Options = opts
	return
}

// Fetch base image. For most images this will be the originally uploaded file. For
// larger images it can be a near-lossless version of the original.
func (r *V1BlobService) Get(ctx context.Context, imageID string, query V1BlobGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/*")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/%s/blob", query.AccountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1BlobGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
