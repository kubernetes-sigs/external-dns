// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vectorize

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// IndexMetadataIndexService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndexMetadataIndexService] method instead.
type IndexMetadataIndexService struct {
	Options []option.RequestOption
}

// NewIndexMetadataIndexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndexMetadataIndexService(opts ...option.RequestOption) (r *IndexMetadataIndexService) {
	r = &IndexMetadataIndexService{}
	r.Options = opts
	return
}

// Enable metadata filtering based on metadata property. Limited to 10 properties.
func (r *IndexMetadataIndexService) New(ctx context.Context, indexName string, params IndexMetadataIndexNewParams, opts ...option.RequestOption) (res *IndexMetadataIndexNewResponse, err error) {
	var env IndexMetadataIndexNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/create", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Metadata Indexes for the specified Vectorize Index.
func (r *IndexMetadataIndexService) List(ctx context.Context, indexName string, query IndexMetadataIndexListParams, opts ...option.RequestOption) (res *IndexMetadataIndexListResponse, err error) {
	var env IndexMetadataIndexListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/list", query.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allow Vectorize to delete the specified metadata index.
func (r *IndexMetadataIndexService) Delete(ctx context.Context, indexName string, params IndexMetadataIndexDeleteParams, opts ...option.RequestOption) (res *IndexMetadataIndexDeleteResponse, err error) {
	var env IndexMetadataIndexDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/delete", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndexMetadataIndexNewResponse struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                            `json:"mutationId"`
	JSON       indexMetadataIndexNewResponseJSON `json:"-"`
}

// indexMetadataIndexNewResponseJSON contains the JSON metadata for the struct
// [IndexMetadataIndexNewResponse]
type indexMetadataIndexNewResponseJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexMetadataIndexNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexNewResponseJSON) RawJSON() string {
	return r.raw
}

type IndexMetadataIndexListResponse struct {
	// Array of indexed metadata properties.
	MetadataIndexes []IndexMetadataIndexListResponseMetadataIndex `json:"metadataIndexes"`
	JSON            indexMetadataIndexListResponseJSON            `json:"-"`
}

// indexMetadataIndexListResponseJSON contains the JSON metadata for the struct
// [IndexMetadataIndexListResponse]
type indexMetadataIndexListResponseJSON struct {
	MetadataIndexes apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IndexMetadataIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexListResponseJSON) RawJSON() string {
	return r.raw
}

type IndexMetadataIndexListResponseMetadataIndex struct {
	// Specifies the type of indexed metadata property.
	IndexType IndexMetadataIndexListResponseMetadataIndexesIndexType `json:"indexType"`
	// Specifies the indexed metadata property.
	PropertyName string                                          `json:"propertyName"`
	JSON         indexMetadataIndexListResponseMetadataIndexJSON `json:"-"`
}

// indexMetadataIndexListResponseMetadataIndexJSON contains the JSON metadata for
// the struct [IndexMetadataIndexListResponseMetadataIndex]
type indexMetadataIndexListResponseMetadataIndexJSON struct {
	IndexType    apijson.Field
	PropertyName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IndexMetadataIndexListResponseMetadataIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexListResponseMetadataIndexJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of indexed metadata property.
type IndexMetadataIndexListResponseMetadataIndexesIndexType string

const (
	IndexMetadataIndexListResponseMetadataIndexesIndexTypeString  IndexMetadataIndexListResponseMetadataIndexesIndexType = "string"
	IndexMetadataIndexListResponseMetadataIndexesIndexTypeNumber  IndexMetadataIndexListResponseMetadataIndexesIndexType = "number"
	IndexMetadataIndexListResponseMetadataIndexesIndexTypeBoolean IndexMetadataIndexListResponseMetadataIndexesIndexType = "boolean"
)

func (r IndexMetadataIndexListResponseMetadataIndexesIndexType) IsKnown() bool {
	switch r {
	case IndexMetadataIndexListResponseMetadataIndexesIndexTypeString, IndexMetadataIndexListResponseMetadataIndexesIndexTypeNumber, IndexMetadataIndexListResponseMetadataIndexesIndexTypeBoolean:
		return true
	}
	return false
}

type IndexMetadataIndexDeleteResponse struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                               `json:"mutationId"`
	JSON       indexMetadataIndexDeleteResponseJSON `json:"-"`
}

// indexMetadataIndexDeleteResponseJSON contains the JSON metadata for the struct
// [IndexMetadataIndexDeleteResponse]
type indexMetadataIndexDeleteResponseJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexMetadataIndexDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IndexMetadataIndexNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the type of metadata property to index.
	IndexType param.Field[IndexMetadataIndexNewParamsIndexType] `json:"indexType,required"`
	// Specifies the metadata property to index.
	PropertyName param.Field[string] `json:"propertyName,required"`
}

func (r IndexMetadataIndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of metadata property to index.
type IndexMetadataIndexNewParamsIndexType string

const (
	IndexMetadataIndexNewParamsIndexTypeString  IndexMetadataIndexNewParamsIndexType = "string"
	IndexMetadataIndexNewParamsIndexTypeNumber  IndexMetadataIndexNewParamsIndexType = "number"
	IndexMetadataIndexNewParamsIndexTypeBoolean IndexMetadataIndexNewParamsIndexType = "boolean"
)

func (r IndexMetadataIndexNewParamsIndexType) IsKnown() bool {
	switch r {
	case IndexMetadataIndexNewParamsIndexTypeString, IndexMetadataIndexNewParamsIndexTypeNumber, IndexMetadataIndexNewParamsIndexTypeBoolean:
		return true
	}
	return false
}

type IndexMetadataIndexNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   IndexMetadataIndexNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexMetadataIndexNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexMetadataIndexNewResponseEnvelopeJSON    `json:"-"`
}

// indexMetadataIndexNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndexMetadataIndexNewResponseEnvelope]
type indexMetadataIndexNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexMetadataIndexNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexMetadataIndexNewResponseEnvelopeSuccess bool

const (
	IndexMetadataIndexNewResponseEnvelopeSuccessTrue IndexMetadataIndexNewResponseEnvelopeSuccess = true
)

func (r IndexMetadataIndexNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexMetadataIndexNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexMetadataIndexListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexMetadataIndexListResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   IndexMetadataIndexListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexMetadataIndexListResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexMetadataIndexListResponseEnvelopeJSON    `json:"-"`
}

// indexMetadataIndexListResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndexMetadataIndexListResponseEnvelope]
type indexMetadataIndexListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexMetadataIndexListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexMetadataIndexListResponseEnvelopeSuccess bool

const (
	IndexMetadataIndexListResponseEnvelopeSuccessTrue IndexMetadataIndexListResponseEnvelopeSuccess = true
)

func (r IndexMetadataIndexListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexMetadataIndexListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexMetadataIndexDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the metadata property for which the index must be deleted.
	PropertyName param.Field[string] `json:"propertyName,required"`
}

func (r IndexMetadataIndexDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexMetadataIndexDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   IndexMetadataIndexDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexMetadataIndexDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexMetadataIndexDeleteResponseEnvelopeJSON    `json:"-"`
}

// indexMetadataIndexDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndexMetadataIndexDeleteResponseEnvelope]
type indexMetadataIndexDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexMetadataIndexDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexMetadataIndexDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexMetadataIndexDeleteResponseEnvelopeSuccess bool

const (
	IndexMetadataIndexDeleteResponseEnvelopeSuccessTrue IndexMetadataIndexDeleteResponseEnvelopeSuccess = true
)

func (r IndexMetadataIndexDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexMetadataIndexDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
