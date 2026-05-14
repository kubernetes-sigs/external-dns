// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

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
)

// LOADocumentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLOADocumentService] method instead.
type LOADocumentService struct {
	Options []option.RequestOption
}

// NewLOADocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLOADocumentService(opts ...option.RequestOption) (r *LOADocumentService) {
	r = &LOADocumentService{}
	r.Options = opts
	return
}

// Submit LOA document (pdf format) under the account.
func (r *LOADocumentService) New(ctx context.Context, params LOADocumentNewParams, opts ...option.RequestOption) (res *LOADocumentNewResponse, err error) {
	var env LOADocumentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Download specified LOA document under the account.
func (r *LOADocumentService) Get(ctx context.Context, loaDocumentID string, query LOADocumentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if loaDocumentID == "" {
		err = errors.New("missing required loa_document_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", query.AccountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LOADocumentNewResponse struct {
	// Identifier for the uploaded LOA document.
	ID string `json:"id,nullable"`
	// Identifier of a Cloudflare account.
	AccountID string    `json:"account_id"`
	Created   time.Time `json:"created" format:"date-time"`
	// Name of LOA document. Max file size 10MB, and supported filetype is pdf.
	Filename string `json:"filename"`
	// File size of the uploaded LOA document.
	SizeBytes int64 `json:"size_bytes"`
	// Whether the LOA has been verified by Cloudflare staff.
	Verified bool `json:"verified"`
	// Timestamp of the moment the LOA was marked as validated.
	VerifiedAt time.Time                  `json:"verified_at,nullable" format:"date-time"`
	JSON       loaDocumentNewResponseJSON `json:"-"`
}

// loaDocumentNewResponseJSON contains the JSON metadata for the struct
// [LOADocumentNewResponse]
type loaDocumentNewResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Filename    apijson.Field
	SizeBytes   apijson.Field
	Verified    apijson.Field
	VerifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// LOA document to upload.
	LOADocument param.Field[string] `json:"loa_document,required"`
}

func (r LOADocumentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type LOADocumentNewResponseEnvelope struct {
	Errors   []LOADocumentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LOADocumentNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success LOADocumentNewResponseEnvelopeSuccess `json:"success,required"`
	Result  LOADocumentNewResponse                `json:"result"`
	JSON    loaDocumentNewResponseEnvelopeJSON    `json:"-"`
}

// loaDocumentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LOADocumentNewResponseEnvelope]
type loaDocumentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           LOADocumentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             loaDocumentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// loaDocumentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LOADocumentNewResponseEnvelopeErrors]
type loaDocumentNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    loaDocumentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// loaDocumentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [LOADocumentNewResponseEnvelopeErrorsSource]
type loaDocumentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           LOADocumentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             loaDocumentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// loaDocumentNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LOADocumentNewResponseEnvelopeMessages]
type loaDocumentNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    loaDocumentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// loaDocumentNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [LOADocumentNewResponseEnvelopeMessagesSource]
type loaDocumentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LOADocumentNewResponseEnvelopeSuccess bool

const (
	LOADocumentNewResponseEnvelopeSuccessTrue LOADocumentNewResponseEnvelopeSuccess = true
)

func (r LOADocumentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LOADocumentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LOADocumentGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}
