// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// LogoService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogoService] method instead.
type LogoService struct {
	Options []option.RequestOption
}

// NewLogoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogoService(opts ...option.RequestOption) (r *LogoService) {
	r = &LogoService{}
	r.Options = opts
	return
}

// Return new saved logo queries created from image files
func (r *LogoService) New(ctx context.Context, params LogoNewParams, opts ...option.RequestOption) (res *LogoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/logos", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Return a success message after deleting saved logo queries by ID
func (r *LogoService) Delete(ctx context.Context, logoID string, body LogoDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if logoID == "" {
		err = errors.New("missing required logo_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/logos/%s", body.AccountID, logoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LogoNewResponse struct {
	ID         int64               `json:"id"`
	Tag        string              `json:"tag"`
	UploadPath string              `json:"upload_path"`
	JSON       logoNewResponseJSON `json:"-"`
}

// logoNewResponseJSON contains the JSON metadata for the struct [LogoNewResponse]
type logoNewResponseJSON struct {
	ID          apijson.Field
	Tag         apijson.Field
	UploadPath  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logoNewResponseJSON) RawJSON() string {
	return r.raw
}

type LogoNewParams struct {
	AccountID param.Field[string]    `path:"account_id,required"`
	MatchType param.Field[string]    `query:"match_type"`
	Tag       param.Field[string]    `query:"tag"`
	Threshold param.Field[float64]   `query:"threshold"`
	Image     param.Field[io.Reader] `json:"image" format:"binary"`
}

func (r LogoNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [LogoNewParams]'s query parameters as `url.Values`.
func (r LogoNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogoDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
