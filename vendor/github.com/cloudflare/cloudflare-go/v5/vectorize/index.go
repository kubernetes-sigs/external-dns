// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vectorize

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// IndexService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndexService] method instead.
type IndexService struct {
	Options       []option.RequestOption
	MetadataIndex *IndexMetadataIndexService
}

// NewIndexService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIndexService(opts ...option.RequestOption) (r *IndexService) {
	r = &IndexService{}
	r.Options = opts
	r.MetadataIndex = NewIndexMetadataIndexService(opts...)
	return
}

// Creates and returns a new Vectorize Index.
func (r *IndexService) New(ctx context.Context, params IndexNewParams, opts ...option.RequestOption) (res *CreateIndex, err error) {
	var env IndexNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Vectorize Indexes
func (r *IndexService) List(ctx context.Context, query IndexListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CreateIndex], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes", query.AccountID)
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

// Returns a list of Vectorize Indexes
func (r *IndexService) ListAutoPaging(ctx context.Context, query IndexListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CreateIndex] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the specified Vectorize Index.
func (r *IndexService) Delete(ctx context.Context, indexName string, body IndexDeleteParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env IndexDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s", body.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *IndexService) DeleteByIDs(ctx context.Context, indexName string, params IndexDeleteByIDsParams, opts ...option.RequestOption) (res *IndexDeleteByIDsResponse, err error) {
	var env IndexDeleteByIDsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/delete_by_ids", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Vectorize Index.
func (r *IndexService) Get(ctx context.Context, indexName string, query IndexGetParams, opts ...option.RequestOption) (res *CreateIndex, err error) {
	var env IndexGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s", query.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *IndexService) GetByIDs(ctx context.Context, indexName string, params IndexGetByIDsParams, opts ...option.RequestOption) (res *IndexGetByIDsResponse, err error) {
	var env IndexGetByIDsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/get_by_ids", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a vectorize index.
func (r *IndexService) Info(ctx context.Context, indexName string, query IndexInfoParams, opts ...option.RequestOption) (res *IndexInfoResponse, err error) {
	var env IndexInfoResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/info", query.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Inserts vectors into the specified index and returns a mutation id corresponding
// to the vectors enqueued for insertion.
func (r *IndexService) Insert(ctx context.Context, indexName string, params IndexInsertParams, opts ...option.RequestOption) (res *IndexInsertResponse, err error) {
	var env IndexInsertResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/insert", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Finds vectors closest to a given vector in an index.
func (r *IndexService) Query(ctx context.Context, indexName string, params IndexQueryParams, opts ...option.RequestOption) (res *IndexQueryResponse, err error) {
	var env IndexQueryResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/query", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns a mutation id corresponding to the vectors enqueued for upsertion.
func (r *IndexService) Upsert(ctx context.Context, indexName string, params IndexUpsertParams, opts ...option.RequestOption) (res *IndexUpsertResponse, err error) {
	var env IndexUpsertResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/upsert", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CreateIndex struct {
	Config IndexDimensionConfiguration `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn time.Time       `json:"modified_on" format:"date-time"`
	Name       string          `json:"name"`
	JSON       createIndexJSON `json:"-"`
}

// createIndexJSON contains the JSON metadata for the struct [CreateIndex]
type createIndexJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createIndexJSON) RawJSON() string {
	return r.raw
}

type IndexDimensionConfiguration struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric IndexDimensionConfigurationMetric `json:"metric,required"`
	JSON   indexDimensionConfigurationJSON   `json:"-"`
}

// indexDimensionConfigurationJSON contains the JSON metadata for the struct
// [IndexDimensionConfiguration]
type indexDimensionConfigurationJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDimensionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDimensionConfigurationJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of metric to use calculating distance.
type IndexDimensionConfigurationMetric string

const (
	IndexDimensionConfigurationMetricCosine     IndexDimensionConfigurationMetric = "cosine"
	IndexDimensionConfigurationMetricEuclidean  IndexDimensionConfigurationMetric = "euclidean"
	IndexDimensionConfigurationMetricDOTProduct IndexDimensionConfigurationMetric = "dot-product"
)

func (r IndexDimensionConfigurationMetric) IsKnown() bool {
	switch r {
	case IndexDimensionConfigurationMetricCosine, IndexDimensionConfigurationMetricEuclidean, IndexDimensionConfigurationMetricDOTProduct:
		return true
	}
	return false
}

type IndexDimensionConfigurationParam struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[IndexDimensionConfigurationMetric] `json:"metric,required"`
}

func (r IndexDimensionConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexDimensionConfigurationParam) implementsIndexNewParamsConfigUnion() {}

type IndexDeleteByIDsResponse struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                       `json:"mutationId"`
	JSON       indexDeleteByIDsResponseJSON `json:"-"`
}

// indexDeleteByIDsResponseJSON contains the JSON metadata for the struct
// [IndexDeleteByIDsResponse]
type indexDeleteByIDsResponseJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteByIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteByIDsResponseJSON) RawJSON() string {
	return r.raw
}

type IndexGetByIDsResponse = interface{}

type IndexInfoResponse struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions"`
	// Specifies the timestamp the last mutation batch was processed as an ISO8601
	// string.
	ProcessedUpToDatetime time.Time `json:"processedUpToDatetime,nullable" format:"date-time"`
	// The unique identifier for the async mutation operation containing the changeset.
	ProcessedUpToMutation string `json:"processedUpToMutation"`
	// Specifies the number of vectors present in the index
	VectorCount int64                 `json:"vectorCount"`
	JSON        indexInfoResponseJSON `json:"-"`
}

// indexInfoResponseJSON contains the JSON metadata for the struct
// [IndexInfoResponse]
type indexInfoResponseJSON struct {
	Dimensions            apijson.Field
	ProcessedUpToDatetime apijson.Field
	ProcessedUpToMutation apijson.Field
	VectorCount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IndexInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInfoResponseJSON) RawJSON() string {
	return r.raw
}

type IndexInsertResponse struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                  `json:"mutationId"`
	JSON       indexInsertResponseJSON `json:"-"`
}

// indexInsertResponseJSON contains the JSON metadata for the struct
// [IndexInsertResponse]
type indexInsertResponseJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertResponseJSON) RawJSON() string {
	return r.raw
}

type IndexQueryResponse struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []IndexQueryResponseMatch `json:"matches"`
	JSON    indexQueryResponseJSON    `json:"-"`
}

// indexQueryResponseJSON contains the JSON metadata for the struct
// [IndexQueryResponse]
type indexQueryResponseJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseJSON) RawJSON() string {
	return r.raw
}

type IndexQueryResponseMatch struct {
	// Identifier for a Vector
	ID        string      `json:"id"`
	Metadata  interface{} `json:"metadata,nullable"`
	Namespace string      `json:"namespace,nullable"`
	// The score of the vector according to the index's distance metric
	Score  float64                     `json:"score"`
	Values []float64                   `json:"values,nullable"`
	JSON   indexQueryResponseMatchJSON `json:"-"`
}

// indexQueryResponseMatchJSON contains the JSON metadata for the struct
// [IndexQueryResponseMatch]
type indexQueryResponseMatchJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Namespace   apijson.Field
	Score       apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseMatchJSON) RawJSON() string {
	return r.raw
}

type IndexUpsertResponse struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                  `json:"mutationId"`
	JSON       indexUpsertResponseJSON `json:"-"`
}

// indexUpsertResponseJSON contains the JSON metadata for the struct
// [IndexUpsertResponse]
type indexUpsertResponseJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type IndexNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the type of configuration to use for the index.
	Config param.Field[IndexNewParamsConfigUnion] `json:"config,required"`
	Name   param.Field[string]                    `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r IndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of configuration to use for the index.
type IndexNewParamsConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[IndexNewParamsConfigMetric] `json:"metric"`
	// Specifies the preset to use for the index.
	Preset param.Field[IndexNewParamsConfigPreset] `json:"preset"`
}

func (r IndexNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfig) implementsIndexNewParamsConfigUnion() {}

// Specifies the type of configuration to use for the index.
//
// Satisfied by [vectorize.IndexDimensionConfigurationParam],
// [vectorize.IndexNewParamsConfigVectorizeIndexPresetConfiguration],
// [IndexNewParamsConfig].
type IndexNewParamsConfigUnion interface {
	implementsIndexNewParamsConfigUnion()
}

type IndexNewParamsConfigVectorizeIndexPresetConfiguration struct {
	// Specifies the preset to use for the index.
	Preset param.Field[IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset] `json:"preset,required"`
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) implementsIndexNewParamsConfigUnion() {
}

// Specifies the preset to use for the index.
type IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset string

const (
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5        IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-small-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5         IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-base-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5        IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-large-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002   IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "openai/text-embedding-ada-002"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0 IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "cohere/embed-multilingual-v2.0"
)

func (r IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

// Specifies the type of metric to use calculating distance.
type IndexNewParamsConfigMetric string

const (
	IndexNewParamsConfigMetricCosine     IndexNewParamsConfigMetric = "cosine"
	IndexNewParamsConfigMetricEuclidean  IndexNewParamsConfigMetric = "euclidean"
	IndexNewParamsConfigMetricDOTProduct IndexNewParamsConfigMetric = "dot-product"
)

func (r IndexNewParamsConfigMetric) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigMetricCosine, IndexNewParamsConfigMetricEuclidean, IndexNewParamsConfigMetricDOTProduct:
		return true
	}
	return false
}

// Specifies the preset to use for the index.
type IndexNewParamsConfigPreset string

const (
	IndexNewParamsConfigPresetCfBaaiBgeSmallEnV1_5        IndexNewParamsConfigPreset = "@cf/baai/bge-small-en-v1.5"
	IndexNewParamsConfigPresetCfBaaiBgeBaseEnV1_5         IndexNewParamsConfigPreset = "@cf/baai/bge-base-en-v1.5"
	IndexNewParamsConfigPresetCfBaaiBgeLargeEnV1_5        IndexNewParamsConfigPreset = "@cf/baai/bge-large-en-v1.5"
	IndexNewParamsConfigPresetOpenAITextEmbeddingAda002   IndexNewParamsConfigPreset = "openai/text-embedding-ada-002"
	IndexNewParamsConfigPresetCohereEmbedMultilingualV2_0 IndexNewParamsConfigPreset = "cohere/embed-multilingual-v2.0"
)

func (r IndexNewParamsConfigPreset) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigPresetCfBaaiBgeSmallEnV1_5, IndexNewParamsConfigPresetCfBaaiBgeBaseEnV1_5, IndexNewParamsConfigPresetCfBaaiBgeLargeEnV1_5, IndexNewParamsConfigPresetOpenAITextEmbeddingAda002, IndexNewParamsConfigPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

type IndexNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CreateIndex           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexNewResponseEnvelopeJSON    `json:"-"`
}

// indexNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexNewResponseEnvelope]
type indexNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexNewResponseEnvelopeSuccess bool

const (
	IndexNewResponseEnvelopeSuccessTrue IndexNewResponseEnvelopeSuccess = true
)

func (r IndexNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexDeleteResponseEnvelopeJSON    `json:"-"`
}

// indexDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexDeleteResponseEnvelope]
type indexDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexDeleteResponseEnvelopeSuccess bool

const (
	IndexDeleteResponseEnvelopeSuccessTrue IndexDeleteResponseEnvelopeSuccess = true
)

func (r IndexDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexDeleteByIDsParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of vector identifiers to delete from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexDeleteByIDsResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   IndexDeleteByIDsResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexDeleteByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexDeleteByIDsResponseEnvelopeJSON    `json:"-"`
}

// indexDeleteByIDsResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexDeleteByIDsResponseEnvelope]
type indexDeleteByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteByIDsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexDeleteByIDsResponseEnvelopeSuccess bool

const (
	IndexDeleteByIDsResponseEnvelopeSuccessTrue IndexDeleteByIDsResponseEnvelopeSuccess = true
)

func (r IndexDeleteByIDsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexDeleteByIDsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CreateIndex           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexGetResponseEnvelopeJSON    `json:"-"`
}

// indexGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexGetResponseEnvelope]
type indexGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexGetResponseEnvelopeSuccess bool

const (
	IndexGetResponseEnvelopeSuccessTrue IndexGetResponseEnvelopeSuccess = true
)

func (r IndexGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexGetByIDsParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of vector identifiers to retrieve from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexGetByIDsResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Array of vectors with matching ids.
	Result IndexGetByIDsResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexGetByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexGetByIDsResponseEnvelopeJSON    `json:"-"`
}

// indexGetByIDsResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexGetByIDsResponseEnvelope]
type indexGetByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetByIDsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexGetByIDsResponseEnvelopeSuccess bool

const (
	IndexGetByIDsResponseEnvelopeSuccessTrue IndexGetByIDsResponseEnvelopeSuccess = true
)

func (r IndexGetByIDsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexGetByIDsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexInfoParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexInfoResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexInfoResponse     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexInfoResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexInfoResponseEnvelopeJSON    `json:"-"`
}

// indexInfoResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexInfoResponseEnvelope]
type indexInfoResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInfoResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInfoResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexInfoResponseEnvelopeSuccess bool

const (
	IndexInfoResponseEnvelopeSuccessTrue IndexInfoResponseEnvelopeSuccess = true
)

func (r IndexInfoResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexInfoResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexInsertParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// ndjson file containing vectors to insert.
	Body string `json:"body,required"`
	// Behavior for ndjson parse failures.
	UnparsableBehavior param.Field[IndexInsertParamsUnparsableBehavior] `query:"unparsable-behavior"`
}

func (r IndexInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [IndexInsertParams]'s query parameters as `url.Values`.
func (r IndexInsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Behavior for ndjson parse failures.
type IndexInsertParamsUnparsableBehavior string

const (
	IndexInsertParamsUnparsableBehaviorError   IndexInsertParamsUnparsableBehavior = "error"
	IndexInsertParamsUnparsableBehaviorDiscard IndexInsertParamsUnparsableBehavior = "discard"
)

func (r IndexInsertParamsUnparsableBehavior) IsKnown() bool {
	switch r {
	case IndexInsertParamsUnparsableBehaviorError, IndexInsertParamsUnparsableBehaviorDiscard:
		return true
	}
	return false
}

type IndexInsertResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexInsertResponse   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexInsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexInsertResponseEnvelopeJSON    `json:"-"`
}

// indexInsertResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexInsertResponseEnvelope]
type indexInsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexInsertResponseEnvelopeSuccess bool

const (
	IndexInsertResponseEnvelopeSuccessTrue IndexInsertResponseEnvelopeSuccess = true
)

func (r IndexInsertResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexInsertResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexQueryParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector,required"`
	// A metadata filter expression used to limit nearest neighbor results.
	Filter param.Field[interface{}] `json:"filter"`
	// Whether to return no metadata, indexed metadata or all metadata associated with
	// the closest vectors.
	ReturnMetadata param.Field[IndexQueryParamsReturnMetadata] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
}

func (r IndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to return no metadata, indexed metadata or all metadata associated with
// the closest vectors.
type IndexQueryParamsReturnMetadata string

const (
	IndexQueryParamsReturnMetadataNone    IndexQueryParamsReturnMetadata = "none"
	IndexQueryParamsReturnMetadataIndexed IndexQueryParamsReturnMetadata = "indexed"
	IndexQueryParamsReturnMetadataAll     IndexQueryParamsReturnMetadata = "all"
)

func (r IndexQueryParamsReturnMetadata) IsKnown() bool {
	switch r {
	case IndexQueryParamsReturnMetadataNone, IndexQueryParamsReturnMetadataIndexed, IndexQueryParamsReturnMetadataAll:
		return true
	}
	return false
}

type IndexQueryResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexQueryResponse    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexQueryResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexQueryResponseEnvelopeJSON    `json:"-"`
}

// indexQueryResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexQueryResponseEnvelope]
type indexQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexQueryResponseEnvelopeSuccess bool

const (
	IndexQueryResponseEnvelopeSuccessTrue IndexQueryResponseEnvelopeSuccess = true
)

func (r IndexQueryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexQueryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexUpsertParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// ndjson file containing vectors to upsert.
	Body string `json:"body,required"`
	// Behavior for ndjson parse failures.
	UnparsableBehavior param.Field[IndexUpsertParamsUnparsableBehavior] `query:"unparsable-behavior"`
}

func (r IndexUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [IndexUpsertParams]'s query parameters as `url.Values`.
func (r IndexUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Behavior for ndjson parse failures.
type IndexUpsertParamsUnparsableBehavior string

const (
	IndexUpsertParamsUnparsableBehaviorError   IndexUpsertParamsUnparsableBehavior = "error"
	IndexUpsertParamsUnparsableBehaviorDiscard IndexUpsertParamsUnparsableBehavior = "discard"
)

func (r IndexUpsertParamsUnparsableBehavior) IsKnown() bool {
	switch r {
	case IndexUpsertParamsUnparsableBehaviorError, IndexUpsertParamsUnparsableBehaviorDiscard:
		return true
	}
	return false
}

type IndexUpsertResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexUpsertResponse   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexUpsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexUpsertResponseEnvelopeJSON    `json:"-"`
}

// indexUpsertResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexUpsertResponseEnvelope]
type indexUpsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexUpsertResponseEnvelopeSuccess bool

const (
	IndexUpsertResponseEnvelopeSuccessTrue IndexUpsertResponseEnvelopeSuccess = true
)

func (r IndexUpsertResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexUpsertResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
