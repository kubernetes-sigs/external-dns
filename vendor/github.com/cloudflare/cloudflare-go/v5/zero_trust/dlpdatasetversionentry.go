// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPDatasetVersionEntryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetVersionEntryService] method instead.
type DLPDatasetVersionEntryService struct {
	Options []option.RequestOption
}

// NewDLPDatasetVersionEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDatasetVersionEntryService(opts ...option.RequestOption) (r *DLPDatasetVersionEntryService) {
	r = &DLPDatasetVersionEntryService{}
	r.Options = opts
	return
}

// This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
// created in the Cloudflare dashboard.
func (r *DLPDatasetVersionEntryService) New(ctx context.Context, datasetID string, version int64, entryID string, datasetVersionEntry io.Reader, body DLPDatasetVersionEntryNewParams, opts ...option.RequestOption) (res *DLPDatasetVersionEntryNewResponse, err error) {
	var env DLPDatasetVersionEntryNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", datasetVersionEntry)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/versions/%v/entries/%s", body.AccountID, datasetID, version, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPDatasetVersionEntryNewResponse struct {
	EntryID      string                                        `json:"entry_id,required" format:"uuid"`
	HeaderName   string                                        `json:"header_name,required"`
	NumCells     int64                                         `json:"num_cells,required"`
	UploadStatus DLPDatasetVersionEntryNewResponseUploadStatus `json:"upload_status,required"`
	JSON         dlpDatasetVersionEntryNewResponseJSON         `json:"-"`
}

// dlpDatasetVersionEntryNewResponseJSON contains the JSON metadata for the struct
// [DLPDatasetVersionEntryNewResponse]
type dlpDatasetVersionEntryNewResponseJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionEntryNewResponseUploadStatus string

const (
	DLPDatasetVersionEntryNewResponseUploadStatusEmpty      DLPDatasetVersionEntryNewResponseUploadStatus = "empty"
	DLPDatasetVersionEntryNewResponseUploadStatusUploading  DLPDatasetVersionEntryNewResponseUploadStatus = "uploading"
	DLPDatasetVersionEntryNewResponseUploadStatusPending    DLPDatasetVersionEntryNewResponseUploadStatus = "pending"
	DLPDatasetVersionEntryNewResponseUploadStatusProcessing DLPDatasetVersionEntryNewResponseUploadStatus = "processing"
	DLPDatasetVersionEntryNewResponseUploadStatusFailed     DLPDatasetVersionEntryNewResponseUploadStatus = "failed"
	DLPDatasetVersionEntryNewResponseUploadStatusComplete   DLPDatasetVersionEntryNewResponseUploadStatus = "complete"
)

func (r DLPDatasetVersionEntryNewResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPDatasetVersionEntryNewResponseUploadStatusEmpty, DLPDatasetVersionEntryNewResponseUploadStatusUploading, DLPDatasetVersionEntryNewResponseUploadStatusPending, DLPDatasetVersionEntryNewResponseUploadStatusProcessing, DLPDatasetVersionEntryNewResponseUploadStatusFailed, DLPDatasetVersionEntryNewResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetVersionEntryNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

func (r DLPDatasetVersionEntryNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type DLPDatasetVersionEntryNewResponseEnvelope struct {
	Errors   []DLPDatasetVersionEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetVersionEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetVersionEntryNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPDatasetVersionEntryNewResponse                `json:"result"`
	JSON    dlpDatasetVersionEntryNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetVersionEntryNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetVersionEntryNewResponseEnvelope]
type dlpDatasetVersionEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionEntryNewResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPDatasetVersionEntryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetVersionEntryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetVersionEntryNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPDatasetVersionEntryNewResponseEnvelopeErrors]
type dlpDatasetVersionEntryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionEntryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpDatasetVersionEntryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetVersionEntryNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPDatasetVersionEntryNewResponseEnvelopeErrorsSource]
type dlpDatasetVersionEntryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionEntryNewResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPDatasetVersionEntryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetVersionEntryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetVersionEntryNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPDatasetVersionEntryNewResponseEnvelopeMessages]
type dlpDatasetVersionEntryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionEntryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpDatasetVersionEntryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetVersionEntryNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPDatasetVersionEntryNewResponseEnvelopeMessagesSource]
type dlpDatasetVersionEntryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetVersionEntryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionEntryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetVersionEntryNewResponseEnvelopeSuccess bool

const (
	DLPDatasetVersionEntryNewResponseEnvelopeSuccessTrue DLPDatasetVersionEntryNewResponseEnvelopeSuccess = true
)

func (r DLPDatasetVersionEntryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetVersionEntryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
