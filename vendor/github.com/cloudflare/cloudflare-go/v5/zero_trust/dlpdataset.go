// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DLPDatasetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetService] method instead.
type DLPDatasetService struct {
	Options  []option.RequestOption
	Upload   *DLPDatasetUploadService
	Versions *DLPDatasetVersionService
}

// NewDLPDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPDatasetService(opts ...option.RequestOption) (r *DLPDatasetService) {
	r = &DLPDatasetService{}
	r.Options = opts
	r.Upload = NewDLPDatasetUploadService(opts...)
	r.Versions = NewDLPDatasetVersionService(opts...)
	return
}

// Create a new dataset
func (r *DLPDatasetService) New(ctx context.Context, params DLPDatasetNewParams, opts ...option.RequestOption) (res *DatasetCreation, err error) {
	var env DLPDatasetNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update details about a dataset
func (r *DLPDatasetService) Update(ctx context.Context, datasetID string, params DLPDatasetUpdateParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env DLPDatasetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch all datasets
func (r *DLPDatasetService) List(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Dataset], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetch all datasets
func (r *DLPDatasetService) ListAutoPaging(ctx context.Context, query DLPDatasetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Dataset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// This deletes all versions of the dataset.
func (r *DLPDatasetService) Delete(ctx context.Context, datasetID string, body DLPDatasetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", body.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch a specific dataset
func (r *DLPDatasetService) Get(ctx context.Context, datasetID string, query DLPDatasetGetParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env DLPDatasetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Dataset struct {
	ID              string          `json:"id,required" format:"uuid"`
	Columns         []DatasetColumn `json:"columns,required"`
	CreatedAt       time.Time       `json:"created_at,required" format:"date-time"`
	EncodingVersion int64           `json:"encoding_version,required"`
	Name            string          `json:"name,required"`
	NumCells        int64           `json:"num_cells,required"`
	Secret          bool            `json:"secret,required"`
	Status          DatasetStatus   `json:"status,required"`
	// When the dataset was last updated.
	//
	// This includes name or description changes as well as uploads.
	UpdatedAt     time.Time       `json:"updated_at,required" format:"date-time"`
	Uploads       []DatasetUpload `json:"uploads,required"`
	CaseSensitive bool            `json:"case_sensitive"`
	// The description of the dataset.
	Description string      `json:"description,nullable"`
	JSON        datasetJSON `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	ID              apijson.Field
	Columns         apijson.Field
	CreatedAt       apijson.Field
	EncodingVersion apijson.Field
	Name            apijson.Field
	NumCells        apijson.Field
	Secret          apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	Uploads         apijson.Field
	CaseSensitive   apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

type DatasetColumn struct {
	EntryID      string                     `json:"entry_id,required" format:"uuid"`
	HeaderName   string                     `json:"header_name,required"`
	NumCells     int64                      `json:"num_cells,required"`
	UploadStatus DatasetColumnsUploadStatus `json:"upload_status,required"`
	JSON         datasetColumnJSON          `json:"-"`
}

// datasetColumnJSON contains the JSON metadata for the struct [DatasetColumn]
type datasetColumnJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DatasetColumn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetColumnJSON) RawJSON() string {
	return r.raw
}

type DatasetColumnsUploadStatus string

const (
	DatasetColumnsUploadStatusEmpty      DatasetColumnsUploadStatus = "empty"
	DatasetColumnsUploadStatusUploading  DatasetColumnsUploadStatus = "uploading"
	DatasetColumnsUploadStatusPending    DatasetColumnsUploadStatus = "pending"
	DatasetColumnsUploadStatusProcessing DatasetColumnsUploadStatus = "processing"
	DatasetColumnsUploadStatusFailed     DatasetColumnsUploadStatus = "failed"
	DatasetColumnsUploadStatusComplete   DatasetColumnsUploadStatus = "complete"
)

func (r DatasetColumnsUploadStatus) IsKnown() bool {
	switch r {
	case DatasetColumnsUploadStatusEmpty, DatasetColumnsUploadStatusUploading, DatasetColumnsUploadStatusPending, DatasetColumnsUploadStatusProcessing, DatasetColumnsUploadStatusFailed, DatasetColumnsUploadStatusComplete:
		return true
	}
	return false
}

type DatasetStatus string

const (
	DatasetStatusEmpty      DatasetStatus = "empty"
	DatasetStatusUploading  DatasetStatus = "uploading"
	DatasetStatusPending    DatasetStatus = "pending"
	DatasetStatusProcessing DatasetStatus = "processing"
	DatasetStatusFailed     DatasetStatus = "failed"
	DatasetStatusComplete   DatasetStatus = "complete"
)

func (r DatasetStatus) IsKnown() bool {
	switch r {
	case DatasetStatusEmpty, DatasetStatusUploading, DatasetStatusPending, DatasetStatusProcessing, DatasetStatusFailed, DatasetStatusComplete:
		return true
	}
	return false
}

type DatasetUpload struct {
	NumCells int64                `json:"num_cells,required"`
	Status   DatasetUploadsStatus `json:"status,required"`
	Version  int64                `json:"version,required"`
	JSON     datasetUploadJSON    `json:"-"`
}

// datasetUploadJSON contains the JSON metadata for the struct [DatasetUpload]
type datasetUploadJSON struct {
	NumCells    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUploadJSON) RawJSON() string {
	return r.raw
}

type DatasetUploadsStatus string

const (
	DatasetUploadsStatusEmpty      DatasetUploadsStatus = "empty"
	DatasetUploadsStatusUploading  DatasetUploadsStatus = "uploading"
	DatasetUploadsStatusPending    DatasetUploadsStatus = "pending"
	DatasetUploadsStatusProcessing DatasetUploadsStatus = "processing"
	DatasetUploadsStatusFailed     DatasetUploadsStatus = "failed"
	DatasetUploadsStatusComplete   DatasetUploadsStatus = "complete"
)

func (r DatasetUploadsStatus) IsKnown() bool {
	switch r {
	case DatasetUploadsStatusEmpty, DatasetUploadsStatusUploading, DatasetUploadsStatusPending, DatasetUploadsStatusProcessing, DatasetUploadsStatusFailed, DatasetUploadsStatusComplete:
		return true
	}
	return false
}

type DatasetArray []Dataset

type DatasetCreation struct {
	Dataset Dataset `json:"dataset,required"`
	// Encoding version to use for dataset.
	EncodingVersion int64 `json:"encoding_version,required"`
	MaxCells        int64 `json:"max_cells,required"`
	// The version to use when uploading the dataset.
	Version int64 `json:"version,required"`
	// The secret to use for Exact Data Match datasets. This is not present in Custom
	// Wordlists.
	Secret string              `json:"secret" format:"password"`
	JSON   datasetCreationJSON `json:"-"`
}

// datasetCreationJSON contains the JSON metadata for the struct [DatasetCreation]
type datasetCreationJSON struct {
	Dataset         apijson.Field
	EncodingVersion apijson.Field
	MaxCells        apijson.Field
	Version         apijson.Field
	Secret          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DatasetCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetCreationJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if `secret` is true or undefined
	CaseSensitive param.Field[bool] `json:"case_sensitive"`
	// The description of the dataset.
	Description param.Field[string] `json:"description"`
	// Dataset encoding version
	//
	// Non-secret custom word lists with no header are always version 1. Secret EDM
	// lists with no header are version 1. Multicolumn CSV with headers are version 2.
	// Omitting this field provides the default value 0, which is interpreted the same
	// as 1.
	EncodingVersion param.Field[int64] `json:"encoding_version"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder. If
	// false, the response has no secret and the dataset is uploaded in plaintext.
	Secret param.Field[bool] `json:"secret"`
}

func (r DLPDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDatasetNewResponseEnvelope struct {
	Errors   []DLPDatasetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DatasetCreation                      `json:"result"`
	JSON    dlpDatasetNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetNewResponseEnvelope]
type dlpDatasetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPDatasetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeErrors]
type dlpDatasetNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpDatasetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeErrorsSource]
type dlpDatasetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPDatasetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetNewResponseEnvelopeMessages]
type dlpDatasetNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpDatasetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPDatasetNewResponseEnvelopeMessagesSource]
type dlpDatasetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetNewResponseEnvelopeSuccess bool

const (
	DLPDatasetNewResponseEnvelopeSuccessTrue DLPDatasetNewResponseEnvelopeSuccess = true
)

func (r DLPDatasetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDatasetUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Determines if the words should be matched in a case-sensitive manner.
	//
	// Only required for custom word lists.
	CaseSensitive param.Field[bool] `json:"case_sensitive"`
	// The description of the dataset.
	Description param.Field[string] `json:"description"`
	// The name of the dataset, must be unique.
	Name param.Field[string] `json:"name"`
}

func (r DLPDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDatasetUpdateResponseEnvelope struct {
	Errors   []DLPDatasetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Dataset                                 `json:"result"`
	JSON    dlpDatasetUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetUpdateResponseEnvelope]
type dlpDatasetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPDatasetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetUpdateResponseEnvelopeErrors]
type dlpDatasetUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpDatasetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPDatasetUpdateResponseEnvelopeErrorsSource]
type dlpDatasetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPDatasetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetUpdateResponseEnvelopeMessages]
type dlpDatasetUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpDatasetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPDatasetUpdateResponseEnvelopeMessagesSource]
type dlpDatasetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetUpdateResponseEnvelopeSuccess bool

const (
	DLPDatasetUpdateResponseEnvelopeSuccessTrue DLPDatasetUpdateResponseEnvelopeSuccess = true
)

func (r DLPDatasetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDatasetListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPDatasetGetResponseEnvelope struct {
	Errors   []DLPDatasetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPDatasetGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPDatasetGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Dataset                              `json:"result"`
	JSON    dlpDatasetGetResponseEnvelopeJSON    `json:"-"`
}

// dlpDatasetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDatasetGetResponseEnvelope]
type dlpDatasetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPDatasetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDatasetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDatasetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeErrors]
type dlpDatasetGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpDatasetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeErrorsSource]
type dlpDatasetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPDatasetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDatasetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDatasetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDatasetGetResponseEnvelopeMessages]
type dlpDatasetGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpDatasetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDatasetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPDatasetGetResponseEnvelopeMessagesSource]
type dlpDatasetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPDatasetGetResponseEnvelopeSuccess bool

const (
	DLPDatasetGetResponseEnvelopeSuccessTrue DLPDatasetGetResponseEnvelopeSuccess = true
)

func (r DLPDatasetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDatasetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
