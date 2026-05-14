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

// DLPDatasetUploadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetUploadService] method instead.
type DLPDatasetUploadService struct {
	Options []option.RequestOption
}

// NewDLPDatasetUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDatasetUploadService(opts ...option.RequestOption) (r *DLPDatasetUploadService) {
	r = &DLPDatasetUploadService{}
	r.Options = opts
	return
}

// Prepare to upload a new version of a dataset
func (r *DLPDatasetUploadService) New(ctx context.Context, datasetID string, body DLPDatasetUploadNewParams, opts ...option.RequestOption) (res *NewVersion, err error) {
	var env DLPDatasetUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is used for single-column EDMv1 and Custom Word Lists. The EDM format can
// only be created in the Cloudflare dashboard. For other clients, this operation
// can only be used for non-secret Custom Word Lists. The body must be a UTF-8
// encoded, newline (NL or CRNL) separated list of words to be matched.
func (r *DLPDatasetUploadService) Edit(ctx context.Context, datasetID string, version int64, dataset io.Reader, body DLPDatasetUploadEditParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env DLPDatasetUploadEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", dataset)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", body.AccountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NewVersion struct {
	EncodingVersion int64              `json:"encoding_version,required"`
	MaxCells        int64              `json:"max_cells,required"`
	Version         int64              `json:"version,required"`
	CaseSensitive   bool               `json:"case_sensitive"`
	Columns         []NewVersionColumn `json:"columns"`
	Secret          string             `json:"secret" format:"password"`
	JSON            newVersionJSON     `json:"-"`
}

// newVersionJSON contains the JSON metadata for the struct [NewVersion]
type newVersionJSON struct {
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	CaseSensitive   apijson.Field
	Columns         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NewVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVersionJSON) RawJSON() string {
	return r.raw
}

type NewVersionColumn struct {
	EntryID      string                        `json:"entry_id,required" format:"uuid"`
	HeaderName   string                        `json:"header_name,required"`
	NumCells     int64                         `json:"num_cells,required"`
	UploadStatus NewVersionColumnsUploadStatus `json:"upload_status,required"`
	JSON         newVersionColumnJSON          `json:"-"`
}

// newVersionColumnJSON contains the JSON metadata for the struct
// [NewVersionColumn]
type newVersionColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NewVersionColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVersionColumnJSON) RawJSON() string {
	return r.raw
}

type NewVersionColumnsUploadStatus string

const (
	NewVersionColumnsUploadStatusEmpty      NewVersionColumnsUploadStatus = "empty"
	NewVersionColumnsUploadStatusUploading  NewVersionColumnsUploadStatus = "uploading"
	NewVersionColumnsUploadStatusPending    NewVersionColumnsUploadStatus = "pending"
	NewVersionColumnsUploadStatusProcessing NewVersionColumnsUploadStatus = "processing"
	NewVersionColumnsUploadStatusFailed     NewVersionColumnsUploadStatus = "failed"
	NewVersionColumnsUploadStatusComplete   NewVersionColumnsUploadStatus = "complete"
)

func (r NewVersionColumnsUploadStatus) IsKnown() bool {
	switch r {
	case NewVersionColumnsUploadStatusEmpty, NewVersionColumnsUploadStatusUploading, NewVersionColumnsUploadStatusPending, NewVersionColumnsUploadStatusProcessing, NewVersionColumnsUploadStatusFailed, NewVersionColumnsUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetUploadNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors   []DLPDatasetUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetUploadNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetUploadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  NewVersion                                 `json:"result"`
	JSON    dlpDatasetUploadNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelope]
type dlpDatasetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPDatasetUploadNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetUploadNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelopeErrors]
type dlpDatasetUploadNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadNewResponseEnvelopeErrorsSource]
type dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPDatasetUploadNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetUploadNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadNewResponseEnvelopeMessages]
type dlpDatasetUploadNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadNewResponseEnvelopeMessagesSource]
type dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetUploadNewResponseEnvelopeSuccess bool

const (
	DLPDatasetUploadNewResponseEnvelopeSuccessTrue DLPDatasetUploadNewResponseEnvelopeSuccess = true
)

func (r DLPDatasetUploadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetUploadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDatasetUploadEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

func (r DLPDatasetUploadEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors   []DLPDatasetUploadEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetUploadEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetUploadEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Dataset                                     `json:"result"`
	JSON    dlpDatasetUploadEditResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadEditResponseEnvelope]
type dlpDatasetUploadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPDatasetUploadEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetUploadEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeErrors]
type dlpDatasetUploadEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUploadEditResponseEnvelopeErrorsSource]
type dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPDatasetUploadEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetUploadEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDatasetUploadEditResponseEnvelopeMessages]
type dlpDatasetUploadEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPDatasetUploadEditResponseEnvelopeMessagesSource]
type dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetUploadEditResponseEnvelopeSuccess bool

const (
	DLPDatasetUploadEditResponseEnvelopeSuccessTrue DLPDatasetUploadEditResponseEnvelopeSuccess = true
)

func (r DLPDatasetUploadEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetUploadEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
