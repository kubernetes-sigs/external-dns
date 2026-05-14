// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AIToMarkdownService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIToMarkdownService] method instead.
type AIToMarkdownService struct {
	Options []option.RequestOption
}

// NewAIToMarkdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIToMarkdownService(opts ...option.RequestOption) (r *AIToMarkdownService) {
	r = &AIToMarkdownService{}
	r.Options = opts
	return
}

// Convert Files into Markdown
func (r *AIToMarkdownService) New(ctx context.Context, Body io.Reader, body AIToMarkdownNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[AIToMarkdownNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tomarkdown", body.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// Convert Files into Markdown
func (r *AIToMarkdownService) NewAutoPaging(ctx context.Context, Body io.Reader, body AIToMarkdownNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AIToMarkdownNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, Body, body, opts...))
}

type AIToMarkdownNewResponse struct {
	Data     string                      `json:"data,required"`
	Format   string                      `json:"format,required"`
	MimeType string                      `json:"mimeType,required"`
	Name     string                      `json:"name,required"`
	Tokens   string                      `json:"tokens,required"`
	JSON     aiToMarkdownNewResponseJSON `json:"-"`
}

// aiToMarkdownNewResponseJSON contains the JSON metadata for the struct
// [AIToMarkdownNewResponse]
type aiToMarkdownNewResponseJSON struct {
	Data        apijson.Field
	Format      apijson.Field
	MimeType    apijson.Field
	Name        apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIToMarkdownNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiToMarkdownNewResponseJSON) RawJSON() string {
	return r.raw
}

type AIToMarkdownNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Body      io.Reader           `json:"body" format:"binary"`
}

func (r AIToMarkdownNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
