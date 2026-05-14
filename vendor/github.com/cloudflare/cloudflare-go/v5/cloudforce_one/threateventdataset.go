// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ThreatEventDatasetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventDatasetService] method instead.
type ThreatEventDatasetService struct {
	Options []option.RequestOption
	Health  *ThreatEventDatasetHealthService
}

// NewThreatEventDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventDatasetService(opts ...option.RequestOption) (r *ThreatEventDatasetService) {
	r = &ThreatEventDatasetService{}
	r.Options = opts
	r.Health = NewThreatEventDatasetHealthService(opts...)
	return
}

// Creates a dataset
func (r *ThreatEventDatasetService) New(ctx context.Context, params ThreatEventDatasetNewParams, opts ...option.RequestOption) (res *ThreatEventDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists all datasets in an account
func (r *ThreatEventDatasetService) List(ctx context.Context, query ThreatEventDatasetListParams, opts ...option.RequestOption) (res *[]ThreatEventDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing dataset
func (r *ThreatEventDatasetService) Edit(ctx context.Context, datasetID string, params ThreatEventDatasetEditParams, opts ...option.RequestOption) (res *ThreatEventDatasetEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Reads a dataset
func (r *ThreatEventDatasetService) Get(ctx context.Context, datasetID string, query ThreatEventDatasetGetParams, opts ...option.RequestOption) (res *ThreatEventDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reads data for a raw event
func (r *ThreatEventDatasetService) Raw(ctx context.Context, datasetID string, eventID string, query ThreatEventDatasetRawParams, opts ...option.RequestOption) (res *ThreatEventDatasetRawResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/raw/%s/%s", query.AccountID, datasetID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventDatasetNewResponse struct {
	IsPublic bool                              `json:"isPublic,required"`
	Name     string                            `json:"name,required"`
	UUID     string                            `json:"uuid,required"`
	JSON     threatEventDatasetNewResponseJSON `json:"-"`
}

// threatEventDatasetNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetNewResponse]
type threatEventDatasetNewResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetListResponse struct {
	IsPublic bool                               `json:"isPublic,required"`
	Name     string                             `json:"name,required"`
	UUID     string                             `json:"uuid,required"`
	JSON     threatEventDatasetListResponseJSON `json:"-"`
}

// threatEventDatasetListResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetListResponse]
type threatEventDatasetListResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetEditResponse struct {
	IsPublic bool                               `json:"isPublic,required"`
	Name     string                             `json:"name,required"`
	UUID     string                             `json:"uuid,required"`
	JSON     threatEventDatasetEditResponseJSON `json:"-"`
}

// threatEventDatasetEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetEditResponse]
type threatEventDatasetEditResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetGetResponse struct {
	IsPublic bool                              `json:"isPublic,required"`
	Name     string                            `json:"name,required"`
	UUID     string                            `json:"uuid,required"`
	JSON     threatEventDatasetGetResponseJSON `json:"-"`
}

// threatEventDatasetGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetGetResponse]
type threatEventDatasetGetResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetRawResponse struct {
	ID        string                            `json:"id,required"`
	AccountID float64                           `json:"accountId,required"`
	Created   string                            `json:"created,required"`
	Data      interface{}                       `json:"data,required"`
	Source    string                            `json:"source,required"`
	TLP       string                            `json:"tlp,required"`
	JSON      threatEventDatasetRawResponseJSON `json:"-"`
}

// threatEventDatasetRawResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetRawResponse]
type threatEventDatasetRawResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Source      apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetRawResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// If true, then anyone can search the dataset. If false, then its limited to the
	// account.
	IsPublic param.Field[bool] `json:"isPublic,required"`
	// Used to describe the dataset within the account context.
	Name param.Field[string] `json:"name,required"`
}

func (r ThreatEventDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventDatasetListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventDatasetEditParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// If true, then anyone can search the dataset. If false, then its limited to the
	// account.
	IsPublic param.Field[bool] `json:"isPublic,required"`
	// Used to describe the dataset within the account context.
	Name param.Field[string] `json:"name,required"`
}

func (r ThreatEventDatasetEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventDatasetGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventDatasetRawParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
