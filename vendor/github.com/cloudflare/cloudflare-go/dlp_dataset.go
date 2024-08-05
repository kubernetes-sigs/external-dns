package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingDatasetID = errors.New("missing required dataset ID")
)

// DLPDatasetUpload represents a single upload version attached to a DLP dataset.
type DLPDatasetUpload struct {
	NumCells int    `json:"num_cells"`
	Status   string `json:"status,omitempty"`
	Version  int    `json:"version"`
}

// DLPDataset represents a DLP Exact Data Match dataset or Custom Word List.
type DLPDataset struct {
	CreatedAt   *time.Time         `json:"created_at,omitempty"`
	Description string             `json:"description,omitempty"`
	ID          string             `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	NumCells    int                `json:"num_cells"`
	Secret      *bool              `json:"secret,omitempty"`
	Status      string             `json:"status,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty"`
	Uploads     []DLPDatasetUpload `json:"uploads"`
}

type ListDLPDatasetsParams struct{}

type DLPDatasetListResponse struct {
	Result []DLPDataset `json:"result"`
	Response
}

// ListDLPDatasets returns all the DLP datasets associated with an account.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-read-all
func (api *API) ListDLPDatasets(ctx context.Context, rc *ResourceContainer, params ListDLPDatasetsParams) ([]DLPDataset, error) {
	if rc.Identifier == "" {
		return nil, nil
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var dlpDatasetListResponse DLPDatasetListResponse
	err = json.Unmarshal(res, &dlpDatasetListResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpDatasetListResponse.Result, nil
}

type DLPDatasetGetResponse struct {
	Result DLPDataset `json:"result"`
	Response
}

// GetDLPDataset returns a DLP dataset based on the dataset ID.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-read
func (api *API) GetDLPDataset(ctx context.Context, rc *ResourceContainer, datasetID string) (DLPDataset, error) {
	if rc.Identifier == "" {
		return DLPDataset{}, nil
	}

	if datasetID == "" {
		return DLPDataset{}, ErrMissingDatasetID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets/%s", rc.Level, rc.Identifier, datasetID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DLPDataset{}, err
	}

	var dlpDatasetGetResponse DLPDatasetGetResponse
	err = json.Unmarshal(res, &dlpDatasetGetResponse)
	if err != nil {
		return DLPDataset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpDatasetGetResponse.Result, nil
}

type CreateDLPDatasetParams struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Secret      *bool  `json:"secret,omitempty"`
}

type CreateDLPDatasetResult struct {
	MaxCells int        `json:"max_cells"`
	Secret   string     `json:"secret"`
	Version  int        `json:"version"`
	Dataset  DLPDataset `json:"dataset"`
}

type CreateDLPDatasetResponse struct {
	Result CreateDLPDatasetResult `json:"result"`
	Response
}

// CreateDLPDataset creates a DLP dataset.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-create
func (api *API) CreateDLPDataset(ctx context.Context, rc *ResourceContainer, params CreateDLPDatasetParams) (CreateDLPDatasetResult, error) {
	if rc.Identifier == "" {
		return CreateDLPDatasetResult{}, nil
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return CreateDLPDatasetResult{}, err
	}

	var CreateDLPDatasetResponse CreateDLPDatasetResponse
	err = json.Unmarshal(res, &CreateDLPDatasetResponse)
	if err != nil {
		return CreateDLPDatasetResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return CreateDLPDatasetResponse.Result, nil
}

// DeleteDLPDataset deletes a DLP dataset.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-delete
func (api *API) DeleteDLPDataset(ctx context.Context, rc *ResourceContainer, datasetID string) error {
	if rc.Identifier == "" {
		return ErrMissingResourceIdentifier
	}

	if datasetID == "" {
		return ErrMissingDatasetID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets/%s", rc.Level, rc.Identifier, datasetID), nil)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

type UpdateDLPDatasetParams struct {
	DatasetID   string
	Description *string `json:"description,omitempty"` // nil to leave descrption as-is
	Name        *string `json:"name,omitempty"`        // nil to leave name as-is
}

type UpdateDLPDatasetResponse struct {
	Result DLPDataset `json:"result"`
	Response
}

// UpdateDLPDataset updates the details of a DLP dataset.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-update
func (api *API) UpdateDLPDataset(ctx context.Context, rc *ResourceContainer, params UpdateDLPDatasetParams) (DLPDataset, error) {
	if rc.Identifier == "" {
		return DLPDataset{}, nil
	}

	if params.DatasetID == "" {
		return DLPDataset{}, ErrMissingDatasetID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets/%s", rc.Level, rc.Identifier, params.DatasetID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return DLPDataset{}, err
	}

	var updateDLPDatasetResponse UpdateDLPDatasetResponse
	err = json.Unmarshal(res, &updateDLPDatasetResponse)
	if err != nil {
		return DLPDataset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateDLPDatasetResponse.Result, nil
}

type CreateDLPDatasetUploadResult struct {
	MaxCells int    `json:"max_cells"`
	Secret   string `json:"secret"`
	Version  int    `json:"version"`
}

type CreateDLPDatasetUploadResponse struct {
	Result CreateDLPDatasetUploadResult `json:"result"`
	Response
}
type CreateDLPDatasetUploadParams struct {
	DatasetID string
}

// CreateDLPDatasetUpload creates a new upload version for the specified DLP dataset.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-create-version
func (api *API) CreateDLPDatasetUpload(ctx context.Context, rc *ResourceContainer, params CreateDLPDatasetUploadParams) (CreateDLPDatasetUploadResult, error) {
	if rc.Identifier == "" {
		return CreateDLPDatasetUploadResult{}, nil
	}

	if params.DatasetID == "" {
		return CreateDLPDatasetUploadResult{}, ErrMissingDatasetID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets/%s/upload", rc.Level, rc.Identifier, params.DatasetID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return CreateDLPDatasetUploadResult{}, err
	}

	var dlpDatasetCreateUploadResponse CreateDLPDatasetUploadResponse
	err = json.Unmarshal(res, &dlpDatasetCreateUploadResponse)
	if err != nil {
		return CreateDLPDatasetUploadResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpDatasetCreateUploadResponse.Result, nil
}

type UploadDLPDatasetVersionParams struct {
	DatasetID string
	Version   int
	Body      interface{}
}

type UploadDLPDatasetVersionResponse struct {
	Result DLPDataset `json:"result"`
	Response
}

// UploadDLPDatasetVersion uploads a new version of the specified DLP dataset.
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-datasets-upload-version
func (api *API) UploadDLPDatasetVersion(ctx context.Context, rc *ResourceContainer, params UploadDLPDatasetVersionParams) (DLPDataset, error) {
	if rc.Identifier == "" {
		return DLPDataset{}, nil
	}

	if params.DatasetID == "" {
		return DLPDataset{}, ErrMissingDatasetID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/datasets/%s/upload/%d", rc.Level, rc.Identifier, params.DatasetID, params.Version), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Body)
	if err != nil {
		return DLPDataset{}, err
	}

	var dlpDatasetUploadVersionResponse UploadDLPDatasetVersionResponse
	err = json.Unmarshal(res, &dlpDatasetUploadVersionResponse)
	if err != nil {
		return DLPDataset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpDatasetUploadVersionResponse.Result, nil
}
