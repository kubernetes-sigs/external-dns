// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// BinaryStorageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBinaryStorageService] method instead.
type BinaryStorageService struct {
	Options []option.RequestOption
}

// NewBinaryStorageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBinaryStorageService(opts ...option.RequestOption) (r *BinaryStorageService) {
	r = &BinaryStorageService{}
	r.Options = opts
	return
}

// Posts a file to Binary Storage
func (r *BinaryStorageService) New(ctx context.Context, params BinaryStorageNewParams, opts ...option.RequestOption) (res *BinaryStorageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/binary", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a file from Binary Storage
func (r *BinaryStorageService) Get(ctx context.Context, hash string, query BinaryStorageGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hash == "" {
		err = errors.New("missing required hash parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/binary/%s", query.AccountID, hash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type BinaryStorageNewResponse struct {
	ContentType string                       `json:"content_type,required"`
	Md5         string                       `json:"md5,required"`
	Sha1        string                       `json:"sha1,required"`
	Sha256      string                       `json:"sha256,required"`
	JSON        binaryStorageNewResponseJSON `json:"-"`
}

// binaryStorageNewResponseJSON contains the JSON metadata for the struct
// [BinaryStorageNewResponse]
type binaryStorageNewResponseJSON struct {
	ContentType apijson.Field
	Md5         apijson.Field
	Sha1        apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BinaryStorageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r binaryStorageNewResponseJSON) RawJSON() string {
	return r.raw
}

type BinaryStorageNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The binary file content to upload.
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
}

func (r BinaryStorageNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type BinaryStorageGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
