// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DLPDatasetVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetVersionService] method instead.
type DLPDatasetVersionService struct {
	Options []option.RequestOption
	Entries *DLPDatasetVersionEntryService
}

// NewDLPDatasetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDatasetVersionService(opts ...option.RequestOption) (r *DLPDatasetVersionService) {
	r = &DLPDatasetVersionService{}
	r.Options = opts
	r.Entries = NewDLPDatasetVersionEntryService(opts...)
	return
}

// This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
// created in the Cloudflare dashboard. The columns in the response appear in the
// same order as in the request.
func (r *DLPDatasetVersionService) New(ctx context.Context, datasetID string, version int64, params DLPDatasetVersionNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPDatasetVersionNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/versions/%v", params.AccountID, datasetID, version)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// This is used for multi-column EDMv2 datasets. The EDMv2 format can only be
// created in the Cloudflare dashboard. The columns in the response appear in the
// same order as in the request.
func (r *DLPDatasetVersionService) NewAutoPaging(ctx context.Context, datasetID string, version int64, params DLPDatasetVersionNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPDatasetVersionNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, datasetID, version, params, opts...))
}

type DLPDatasetVersionNewResponse struct {
	EntryID      string                                   `json:"entry_id,required" format:"uuid"`
	HeaderName   string                                   `json:"header_name,required"`
	NumCells     int64                                    `json:"num_cells,required"`
	UploadStatus DLPDatasetVersionNewResponseUploadStatus `json:"upload_status,required"`
	JSON         dlpDatasetVersionNewResponseJSON         `json:"-"`
}

// dlpDatasetVersionNewResponseJSON contains the JSON metadata for the struct
// [DLPDatasetVersionNewResponse]
type dlpDatasetVersionNewResponseJSON struct {
	EntryID      apijson.Field
	HeaderName   apijson.Field
	NumCells     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPDatasetVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetVersionNewResponseUploadStatus string

const (
	DLPDatasetVersionNewResponseUploadStatusEmpty      DLPDatasetVersionNewResponseUploadStatus = "empty"
	DLPDatasetVersionNewResponseUploadStatusUploading  DLPDatasetVersionNewResponseUploadStatus = "uploading"
	DLPDatasetVersionNewResponseUploadStatusPending    DLPDatasetVersionNewResponseUploadStatus = "pending"
	DLPDatasetVersionNewResponseUploadStatusProcessing DLPDatasetVersionNewResponseUploadStatus = "processing"
	DLPDatasetVersionNewResponseUploadStatusFailed     DLPDatasetVersionNewResponseUploadStatus = "failed"
	DLPDatasetVersionNewResponseUploadStatusComplete   DLPDatasetVersionNewResponseUploadStatus = "complete"
)

func (r DLPDatasetVersionNewResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPDatasetVersionNewResponseUploadStatusEmpty, DLPDatasetVersionNewResponseUploadStatusUploading, DLPDatasetVersionNewResponseUploadStatusPending, DLPDatasetVersionNewResponseUploadStatusProcessing, DLPDatasetVersionNewResponseUploadStatusFailed, DLPDatasetVersionNewResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPDatasetVersionNewParams struct {
	AccountID param.Field[string]                   `path:"account_id,required"`
	Body      []DLPDatasetVersionNewParamsBodyUnion `json:"body,required"`
}

func (r DLPDatasetVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPDatasetVersionNewParamsBody struct {
	EntryID    param.Field[string] `json:"entry_id" format:"uuid"`
	EntryName  param.Field[string] `json:"entry_name"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r DLPDatasetVersionNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPDatasetVersionNewParamsBody) implementsDLPDatasetVersionNewParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPDatasetVersionNewParamsBodyExistingColumn],
// [zero_trust.DLPDatasetVersionNewParamsBodyNewColumn],
// [DLPDatasetVersionNewParamsBody].
type DLPDatasetVersionNewParamsBodyUnion interface {
	implementsDLPDatasetVersionNewParamsBodyUnion()
}

type DLPDatasetVersionNewParamsBodyExistingColumn struct {
	EntryID    param.Field[string] `json:"entry_id,required" format:"uuid"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r DLPDatasetVersionNewParamsBodyExistingColumn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPDatasetVersionNewParamsBodyExistingColumn) implementsDLPDatasetVersionNewParamsBodyUnion() {
}

type DLPDatasetVersionNewParamsBodyNewColumn struct {
	EntryName  param.Field[string] `json:"entry_name,required"`
	HeaderName param.Field[string] `json:"header_name"`
	NumCells   param.Field[int64]  `json:"num_cells"`
}

func (r DLPDatasetVersionNewParamsBodyNewColumn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPDatasetVersionNewParamsBodyNewColumn) implementsDLPDatasetVersionNewParamsBodyUnion() {}
