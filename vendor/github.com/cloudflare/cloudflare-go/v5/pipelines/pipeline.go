// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipelines

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// PipelineService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	Options []option.RequestOption
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r *PipelineService) {
	r = &PipelineService{}
	r.Options = opts
	return
}

// Create a new pipeline.
func (r *PipelineService) New(ctx context.Context, params PipelineNewParams, opts ...option.RequestOption) (res *PipelineNewResponse, err error) {
	var env PipelineNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing pipeline.
func (r *PipelineService) Update(ctx context.Context, pipelineName string, params PipelineUpdateParams, opts ...option.RequestOption) (res *PipelineUpdateResponse, err error) {
	var env PipelineUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", params.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, filter, and paginate pipelines in an account.
func (r *PipelineService) List(ctx context.Context, params PipelineListParams, opts ...option.RequestOption) (res *PipelineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a pipeline.
func (r *PipelineService) Delete(ctx context.Context, pipelineName string, body PipelineDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", body.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get configuration of a pipeline.
func (r *PipelineService) Get(ctx context.Context, pipelineName string, query PipelineGetParams, opts ...option.RequestOption) (res *PipelineGetResponse, err error) {
	var env PipelineGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", query.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Describes the configuration of a pipeline.
type PipelineNewResponse struct {
	// Specifies the pipeline identifier.
	ID          string                         `json:"id,required"`
	Destination PipelineNewResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of the pipeline.
	Name   string                      `json:"name,required"`
	Source []PipelineNewResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                 `json:"version,required"`
	JSON    pipelineNewResponseJSON `json:"-"`
}

// pipelineNewResponseJSON contains the JSON metadata for the struct
// [PipelineNewResponse]
type pipelineNewResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestination struct {
	Batch       PipelineNewResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineNewResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineNewResponseDestinationFormat `json:"format,required"`
	Path   PipelineNewResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineNewResponseDestinationType `json:"type,required"`
	JSON pipelineNewResponseDestinationJSON `json:"-"`
}

// pipelineNewResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineNewResponseDestination]
type pipelineNewResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                   `json:"max_rows,required"`
	JSON    pipelineNewResponseDestinationBatchJSON `json:"-"`
}

// pipelineNewResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineNewResponseDestinationBatch]
type pipelineNewResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineNewResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineNewResponseDestinationCompressionJSON `json:"-"`
}

// pipelineNewResponseDestinationCompressionJSON contains the JSON metadata for the
// struct [PipelineNewResponseDestinationCompression]
type pipelineNewResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineNewResponseDestinationCompressionType string

const (
	PipelineNewResponseDestinationCompressionTypeNone    PipelineNewResponseDestinationCompressionType = "none"
	PipelineNewResponseDestinationCompressionTypeGzip    PipelineNewResponseDestinationCompressionType = "gzip"
	PipelineNewResponseDestinationCompressionTypeDeflate PipelineNewResponseDestinationCompressionType = "deflate"
)

func (r PipelineNewResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationCompressionTypeNone, PipelineNewResponseDestinationCompressionTypeGzip, PipelineNewResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineNewResponseDestinationFormat string

const (
	PipelineNewResponseDestinationFormatJson PipelineNewResponseDestinationFormat = "json"
)

func (r PipelineNewResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                 `json:"prefix"`
	JSON   pipelineNewResponseDestinationPathJSON `json:"-"`
}

// pipelineNewResponseDestinationPathJSON contains the JSON metadata for the struct
// [PipelineNewResponseDestinationPath]
type pipelineNewResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineNewResponseDestinationType string

const (
	PipelineNewResponseDestinationTypeR2 PipelineNewResponseDestinationType = "r2"
)

func (r PipelineNewResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineNewResponseSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceFormat `json:"format,required"`
	Type   string                          `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                   `json:"cors"`
	JSON  pipelineNewResponseSourceJSON `json:"-"`
	union PipelineNewResponseSourceUnion
}

// pipelineNewResponseSourceJSON contains the JSON metadata for the struct
// [PipelineNewResponseSource]
type pipelineNewResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineNewResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineNewResponseSourceUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
func (r PipelineNewResponseSource) AsUnion() PipelineNewResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource] or
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
type PipelineNewResponseSourceUnion interface {
	implementsPipelineNewResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineNewResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                       `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool                                                                       `json:"authentication"`
	CORS           PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource]
type pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineNewResponseSource() {
}

// Specifies the format of source data.
type PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                       `json:"origins"`
	JSON    pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                          `json:"type,required"`
	JSON   pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource]
type pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineNewResponseSource() {
}

// Specifies the format of source data.
type PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineNewResponseSourceFormat string

const (
	PipelineNewResponseSourceFormatJson PipelineNewResponseSourceFormat = "json"
)

func (r PipelineNewResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceFormatJson:
		return true
	}
	return false
}

// Describes the configuration of a pipeline.
type PipelineUpdateResponse struct {
	// Specifies the pipeline identifier.
	ID          string                            `json:"id,required"`
	Destination PipelineUpdateResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of the pipeline.
	Name   string                         `json:"name,required"`
	Source []PipelineUpdateResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                    `json:"version,required"`
	JSON    pipelineUpdateResponseJSON `json:"-"`
}

// pipelineUpdateResponseJSON contains the JSON metadata for the struct
// [PipelineUpdateResponse]
type pipelineUpdateResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestination struct {
	Batch       PipelineUpdateResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineUpdateResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineUpdateResponseDestinationFormat `json:"format,required"`
	Path   PipelineUpdateResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineUpdateResponseDestinationType `json:"type,required"`
	JSON pipelineUpdateResponseDestinationJSON `json:"-"`
}

// pipelineUpdateResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseDestination]
type pipelineUpdateResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                      `json:"max_rows,required"`
	JSON    pipelineUpdateResponseDestinationBatchJSON `json:"-"`
}

// pipelineUpdateResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineUpdateResponseDestinationBatch]
type pipelineUpdateResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineUpdateResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineUpdateResponseDestinationCompressionJSON `json:"-"`
}

// pipelineUpdateResponseDestinationCompressionJSON contains the JSON metadata for
// the struct [PipelineUpdateResponseDestinationCompression]
type pipelineUpdateResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineUpdateResponseDestinationCompressionType string

const (
	PipelineUpdateResponseDestinationCompressionTypeNone    PipelineUpdateResponseDestinationCompressionType = "none"
	PipelineUpdateResponseDestinationCompressionTypeGzip    PipelineUpdateResponseDestinationCompressionType = "gzip"
	PipelineUpdateResponseDestinationCompressionTypeDeflate PipelineUpdateResponseDestinationCompressionType = "deflate"
)

func (r PipelineUpdateResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationCompressionTypeNone, PipelineUpdateResponseDestinationCompressionTypeGzip, PipelineUpdateResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineUpdateResponseDestinationFormat string

const (
	PipelineUpdateResponseDestinationFormatJson PipelineUpdateResponseDestinationFormat = "json"
)

func (r PipelineUpdateResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                    `json:"prefix"`
	JSON   pipelineUpdateResponseDestinationPathJSON `json:"-"`
}

// pipelineUpdateResponseDestinationPathJSON contains the JSON metadata for the
// struct [PipelineUpdateResponseDestinationPath]
type pipelineUpdateResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineUpdateResponseDestinationType string

const (
	PipelineUpdateResponseDestinationTypeR2 PipelineUpdateResponseDestinationType = "r2"
)

func (r PipelineUpdateResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineUpdateResponseSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceFormat `json:"format,required"`
	Type   string                             `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                      `json:"cors"`
	JSON  pipelineUpdateResponseSourceJSON `json:"-"`
	union PipelineUpdateResponseSourceUnion
}

// pipelineUpdateResponseSourceJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseSource]
type pipelineUpdateResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineUpdateResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineUpdateResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineUpdateResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineUpdateResponseSourceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
func (r PipelineUpdateResponseSource) AsUnion() PipelineUpdateResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource] or
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
type PipelineUpdateResponseSourceUnion interface {
	implementsPipelineUpdateResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineUpdateResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                          `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool                                                                          `json:"authentication"`
	CORS           PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource]
type pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineUpdateResponseSource() {
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                          `json:"origins"`
	JSON    pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                             `json:"type,required"`
	JSON   pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource]
type pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineUpdateResponseSource() {
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceFormat string

const (
	PipelineUpdateResponseSourceFormatJson PipelineUpdateResponseSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceFormatJson:
		return true
	}
	return false
}

type PipelineListResponse struct {
	ResultInfo PipelineListResponseResultInfo `json:"result_info,required"`
	Results    []PipelineListResponseResult   `json:"results,required"`
	// Indicates whether the API call was successful.
	Success bool                     `json:"success,required"`
	JSON    pipelineListResponseJSON `json:"-"`
}

// pipelineListResponseJSON contains the JSON metadata for the struct
// [PipelineListResponse]
type pipelineListResponseJSON struct {
	ResultInfo  apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultInfo struct {
	// Indicates the number of items on current page.
	Count float64 `json:"count,required"`
	// Indicates the current page number.
	Page float64 `json:"page,required"`
	// Indicates the number of items per page.
	PerPage float64 `json:"per_page,required"`
	// Indicates the total number of items.
	TotalCount float64                            `json:"total_count,required"`
	JSON       pipelineListResponseResultInfoJSON `json:"-"`
}

// pipelineListResponseResultInfoJSON contains the JSON metadata for the struct
// [PipelineListResponseResultInfo]
type pipelineListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Describes the configuration of a pipeline.
type PipelineListResponseResult struct {
	// Specifies the pipeline identifier.
	ID          string                                 `json:"id,required"`
	Destination PipelineListResponseResultsDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of the pipeline.
	Name   string                              `json:"name,required"`
	Source []PipelineListResponseResultsSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                        `json:"version,required"`
	JSON    pipelineListResponseResultJSON `json:"-"`
}

// pipelineListResponseResultJSON contains the JSON metadata for the struct
// [PipelineListResponseResult]
type pipelineListResponseResultJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestination struct {
	Batch       PipelineListResponseResultsDestinationBatch       `json:"batch,required"`
	Compression PipelineListResponseResultsDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineListResponseResultsDestinationFormat `json:"format,required"`
	Path   PipelineListResponseResultsDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineListResponseResultsDestinationType `json:"type,required"`
	JSON pipelineListResponseResultsDestinationJSON `json:"-"`
}

// pipelineListResponseResultsDestinationJSON contains the JSON metadata for the
// struct [PipelineListResponseResultsDestination]
type pipelineListResponseResultsDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                           `json:"max_rows,required"`
	JSON    pipelineListResponseResultsDestinationBatchJSON `json:"-"`
}

// pipelineListResponseResultsDestinationBatchJSON contains the JSON metadata for
// the struct [PipelineListResponseResultsDestinationBatch]
type pipelineListResponseResultsDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineListResponseResultsDestinationCompressionType `json:"type,required"`
	JSON pipelineListResponseResultsDestinationCompressionJSON `json:"-"`
}

// pipelineListResponseResultsDestinationCompressionJSON contains the JSON metadata
// for the struct [PipelineListResponseResultsDestinationCompression]
type pipelineListResponseResultsDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineListResponseResultsDestinationCompressionType string

const (
	PipelineListResponseResultsDestinationCompressionTypeNone    PipelineListResponseResultsDestinationCompressionType = "none"
	PipelineListResponseResultsDestinationCompressionTypeGzip    PipelineListResponseResultsDestinationCompressionType = "gzip"
	PipelineListResponseResultsDestinationCompressionTypeDeflate PipelineListResponseResultsDestinationCompressionType = "deflate"
)

func (r PipelineListResponseResultsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationCompressionTypeNone, PipelineListResponseResultsDestinationCompressionTypeGzip, PipelineListResponseResultsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineListResponseResultsDestinationFormat string

const (
	PipelineListResponseResultsDestinationFormatJson PipelineListResponseResultsDestinationFormat = "json"
)

func (r PipelineListResponseResultsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineListResponseResultsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                         `json:"prefix"`
	JSON   pipelineListResponseResultsDestinationPathJSON `json:"-"`
}

// pipelineListResponseResultsDestinationPathJSON contains the JSON metadata for
// the struct [PipelineListResponseResultsDestinationPath]
type pipelineListResponseResultsDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineListResponseResultsDestinationType string

const (
	PipelineListResponseResultsDestinationTypeR2 PipelineListResponseResultsDestinationType = "r2"
)

func (r PipelineListResponseResultsDestinationType) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineListResponseResultsSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceFormat `json:"format,required"`
	Type   string                                  `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                           `json:"cors"`
	JSON  pipelineListResponseResultsSourceJSON `json:"-"`
	union PipelineListResponseResultsSourceUnion
}

// pipelineListResponseResultsSourceJSON contains the JSON metadata for the struct
// [PipelineListResponseResultsSource]
type pipelineListResponseResultsSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineListResponseResultsSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineListResponseResultsSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineListResponseResultsSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineListResponseResultsSourceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource].
func (r PipelineListResponseResultsSource) AsUnion() PipelineListResponseResultsSourceUnion {
	return r.union
}

// Union satisfied by
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource]
// or
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource].
type PipelineListResponseResultsSourceUnion interface {
	implementsPipelineListResponseResultsSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineListResponseResultsSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                               `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool                                                                               `json:"authentication"`
	CORS           PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource]
type pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineListResponseResultsSource() {
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                               `json:"origins"`
	JSON    pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                                  `json:"type,required"`
	JSON   pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource]
type pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineListResponseResultsSource() {
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceFormat string

const (
	PipelineListResponseResultsSourceFormatJson PipelineListResponseResultsSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceFormatJson:
		return true
	}
	return false
}

// Describes the configuration of a pipeline.
type PipelineGetResponse struct {
	// Specifies the pipeline identifier.
	ID          string                         `json:"id,required"`
	Destination PipelineGetResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of the pipeline.
	Name   string                      `json:"name,required"`
	Source []PipelineGetResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                 `json:"version,required"`
	JSON    pipelineGetResponseJSON `json:"-"`
}

// pipelineGetResponseJSON contains the JSON metadata for the struct
// [PipelineGetResponse]
type pipelineGetResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestination struct {
	Batch       PipelineGetResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineGetResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineGetResponseDestinationFormat `json:"format,required"`
	Path   PipelineGetResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineGetResponseDestinationType `json:"type,required"`
	JSON pipelineGetResponseDestinationJSON `json:"-"`
}

// pipelineGetResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineGetResponseDestination]
type pipelineGetResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                   `json:"max_rows,required"`
	JSON    pipelineGetResponseDestinationBatchJSON `json:"-"`
}

// pipelineGetResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineGetResponseDestinationBatch]
type pipelineGetResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineGetResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineGetResponseDestinationCompressionJSON `json:"-"`
}

// pipelineGetResponseDestinationCompressionJSON contains the JSON metadata for the
// struct [PipelineGetResponseDestinationCompression]
type pipelineGetResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineGetResponseDestinationCompressionType string

const (
	PipelineGetResponseDestinationCompressionTypeNone    PipelineGetResponseDestinationCompressionType = "none"
	PipelineGetResponseDestinationCompressionTypeGzip    PipelineGetResponseDestinationCompressionType = "gzip"
	PipelineGetResponseDestinationCompressionTypeDeflate PipelineGetResponseDestinationCompressionType = "deflate"
)

func (r PipelineGetResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationCompressionTypeNone, PipelineGetResponseDestinationCompressionTypeGzip, PipelineGetResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineGetResponseDestinationFormat string

const (
	PipelineGetResponseDestinationFormatJson PipelineGetResponseDestinationFormat = "json"
)

func (r PipelineGetResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineGetResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                 `json:"prefix"`
	JSON   pipelineGetResponseDestinationPathJSON `json:"-"`
}

// pipelineGetResponseDestinationPathJSON contains the JSON metadata for the struct
// [PipelineGetResponseDestinationPath]
type pipelineGetResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineGetResponseDestinationType string

const (
	PipelineGetResponseDestinationTypeR2 PipelineGetResponseDestinationType = "r2"
)

func (r PipelineGetResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineGetResponseSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceFormat `json:"format,required"`
	Type   string                          `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                   `json:"cors"`
	JSON  pipelineGetResponseSourceJSON `json:"-"`
	union PipelineGetResponseSourceUnion
}

// pipelineGetResponseSourceJSON contains the JSON metadata for the struct
// [PipelineGetResponseSource]
type pipelineGetResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineGetResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineGetResponseSourceUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
func (r PipelineGetResponseSource) AsUnion() PipelineGetResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource] or
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource].
type PipelineGetResponseSourceUnion interface {
	implementsPipelineGetResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineGetResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                       `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication bool                                                                       `json:"authentication"`
	CORS           PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource]
type pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineGetResponseSource() {
}

// Specifies the format of source data.
type PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                       `json:"origins"`
	JSON    pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                          `json:"type,required"`
	JSON   pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource]
type pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineGetResponseSource() {
}

// Specifies the format of source data.
type PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineGetResponseSourceFormat string

const (
	PipelineGetResponseSourceFormatJson PipelineGetResponseSourceFormat = "json"
)

func (r PipelineGetResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewParams struct {
	// Specifies the public ID of the account.
	AccountID   param.Field[string]                       `path:"account_id,required"`
	Destination param.Field[PipelineNewParamsDestination] `json:"destination,required"`
	// Defines the name of the pipeline.
	Name   param.Field[string]                         `json:"name,required"`
	Source param.Field[[]PipelineNewParamsSourceUnion] `json:"source,required"`
}

func (r PipelineNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestination struct {
	Batch       param.Field[PipelineNewParamsDestinationBatch]       `json:"batch,required"`
	Compression param.Field[PipelineNewParamsDestinationCompression] `json:"compression,required"`
	Credentials param.Field[PipelineNewParamsDestinationCredentials] `json:"credentials,required"`
	// Specifies the format of data to deliver.
	Format param.Field[PipelineNewParamsDestinationFormat] `json:"format,required"`
	Path   param.Field[PipelineNewParamsDestinationPath]   `json:"path,required"`
	// Specifies the type of destination.
	Type param.Field[PipelineNewParamsDestinationType] `json:"type,required"`
}

func (r PipelineNewParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes param.Field[int64] `json:"max_bytes"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS param.Field[float64] `json:"max_duration_s"`
	// Specifies rough maximum number of rows per file.
	MaxRows param.Field[int64] `json:"max_rows"`
}

func (r PipelineNewParamsDestinationBatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type param.Field[PipelineNewParamsDestinationCompressionType] `json:"type"`
}

func (r PipelineNewParamsDestinationCompression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the desired compression algorithm and format.
type PipelineNewParamsDestinationCompressionType string

const (
	PipelineNewParamsDestinationCompressionTypeNone    PipelineNewParamsDestinationCompressionType = "none"
	PipelineNewParamsDestinationCompressionTypeGzip    PipelineNewParamsDestinationCompressionType = "gzip"
	PipelineNewParamsDestinationCompressionTypeDeflate PipelineNewParamsDestinationCompressionType = "deflate"
)

func (r PipelineNewParamsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationCompressionTypeNone, PipelineNewParamsDestinationCompressionTypeGzip, PipelineNewParamsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

type PipelineNewParamsDestinationCredentials struct {
	// Specifies the R2 Bucket Access Key Id.
	AccessKeyID param.Field[string] `json:"access_key_id,required"`
	// Specifies the R2 Endpoint.
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Specifies the R2 Bucket Secret Access Key.
	SecretAccessKey param.Field[string] `json:"secret_access_key,required"`
}

func (r PipelineNewParamsDestinationCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the format of data to deliver.
type PipelineNewParamsDestinationFormat string

const (
	PipelineNewParamsDestinationFormatJson PipelineNewParamsDestinationFormat = "json"
)

func (r PipelineNewParamsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineNewParamsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket param.Field[string] `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename param.Field[string] `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath param.Field[string] `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix param.Field[string] `json:"prefix"`
}

func (r PipelineNewParamsDestinationPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of destination.
type PipelineNewParamsDestinationType string

const (
	PipelineNewParamsDestinationTypeR2 PipelineNewParamsDestinationType = "r2"
)

func (r PipelineNewParamsDestinationType) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineNewParamsSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceFormat] `json:"format,required"`
	Type   param.Field[string]                        `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication param.Field[bool]        `json:"authentication"`
	CORS           param.Field[interface{}] `json:"cors"`
}

func (r PipelineNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSource) implementsPipelineNewParamsSourceUnion() {}

// Satisfied by
// [pipelines.PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource],
// [PipelineNewParamsSource].
type PipelineNewParamsSourceUnion interface {
	implementsPipelineNewParamsSourceUnion()
}

type PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                     `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication param.Field[bool]                                                                     `json:"authentication"`
	CORS           param.Field[PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS] `json:"cors"`
}

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineNewParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins param.Field[[]string] `json:"origins"`
}

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                        `json:"type,required"`
}

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineNewParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineNewParamsSourceFormat string

const (
	PipelineNewParamsSourceFormatJson PipelineNewParamsSourceFormat = "json"
)

func (r PipelineNewParamsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseEnvelope struct {
	// Describes the configuration of a pipeline.
	Result PipelineNewResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                            `json:"success,required"`
	JSON    pipelineNewResponseEnvelopeJSON `json:"-"`
}

// pipelineNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineNewResponseEnvelope]
type pipelineNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateParams struct {
	// Specifies the public ID of the account.
	AccountID   param.Field[string]                          `path:"account_id,required"`
	Destination param.Field[PipelineUpdateParamsDestination] `json:"destination,required"`
	// Defines the name of the pipeline.
	Name   param.Field[string]                            `json:"name,required"`
	Source param.Field[[]PipelineUpdateParamsSourceUnion] `json:"source,required"`
}

func (r PipelineUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestination struct {
	Batch       param.Field[PipelineUpdateParamsDestinationBatch]       `json:"batch,required"`
	Compression param.Field[PipelineUpdateParamsDestinationCompression] `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format param.Field[PipelineUpdateParamsDestinationFormat] `json:"format,required"`
	Path   param.Field[PipelineUpdateParamsDestinationPath]   `json:"path,required"`
	// Specifies the type of destination.
	Type        param.Field[PipelineUpdateParamsDestinationType]        `json:"type,required"`
	Credentials param.Field[PipelineUpdateParamsDestinationCredentials] `json:"credentials"`
}

func (r PipelineUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes param.Field[int64] `json:"max_bytes"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS param.Field[float64] `json:"max_duration_s"`
	// Specifies rough maximum number of rows per file.
	MaxRows param.Field[int64] `json:"max_rows"`
}

func (r PipelineUpdateParamsDestinationBatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type param.Field[PipelineUpdateParamsDestinationCompressionType] `json:"type"`
}

func (r PipelineUpdateParamsDestinationCompression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the desired compression algorithm and format.
type PipelineUpdateParamsDestinationCompressionType string

const (
	PipelineUpdateParamsDestinationCompressionTypeNone    PipelineUpdateParamsDestinationCompressionType = "none"
	PipelineUpdateParamsDestinationCompressionTypeGzip    PipelineUpdateParamsDestinationCompressionType = "gzip"
	PipelineUpdateParamsDestinationCompressionTypeDeflate PipelineUpdateParamsDestinationCompressionType = "deflate"
)

func (r PipelineUpdateParamsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationCompressionTypeNone, PipelineUpdateParamsDestinationCompressionTypeGzip, PipelineUpdateParamsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineUpdateParamsDestinationFormat string

const (
	PipelineUpdateParamsDestinationFormatJson PipelineUpdateParamsDestinationFormat = "json"
)

func (r PipelineUpdateParamsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineUpdateParamsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket param.Field[string] `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename param.Field[string] `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath param.Field[string] `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix param.Field[string] `json:"prefix"`
}

func (r PipelineUpdateParamsDestinationPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of destination.
type PipelineUpdateParamsDestinationType string

const (
	PipelineUpdateParamsDestinationTypeR2 PipelineUpdateParamsDestinationType = "r2"
)

func (r PipelineUpdateParamsDestinationType) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineUpdateParamsDestinationCredentials struct {
	// Specifies the R2 Bucket Access Key Id.
	AccessKeyID param.Field[string] `json:"access_key_id,required"`
	// Specifies the R2 Endpoint.
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Specifies the R2 Bucket Secret Access Key.
	SecretAccessKey param.Field[string] `json:"secret_access_key,required"`
}

func (r PipelineUpdateParamsDestinationCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceFormat] `json:"format,required"`
	Type   param.Field[string]                           `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication param.Field[bool]        `json:"authentication"`
	CORS           param.Field[interface{}] `json:"cors"`
}

func (r PipelineUpdateParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSource) implementsPipelineUpdateParamsSourceUnion() {}

// Satisfied by
// [pipelines.PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource],
// [PipelineUpdateParamsSource].
type PipelineUpdateParamsSourceUnion interface {
	implementsPipelineUpdateParamsSourceUnion()
}

type PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                        `json:"type,required"`
	// Specifies whether authentication is required to send to this pipeline via HTTP.
	Authentication param.Field[bool]                                                                        `json:"authentication"`
	CORS           param.Field[PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS] `json:"cors"`
}

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSource) implementsPipelineUpdateParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins param.Field[[]string] `json:"origins"`
}

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesHTTPSourceCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                           `json:"type,required"`
}

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSource) implementsPipelineUpdateParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceCloudflarePipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceFormat string

const (
	PipelineUpdateParamsSourceFormatJson PipelineUpdateParamsSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseEnvelope struct {
	// Describes the configuration of a pipeline.
	Result PipelineUpdateResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                               `json:"success,required"`
	JSON    pipelineUpdateResponseEnvelopeJSON `json:"-"`
}

// pipelineUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseEnvelope]
type pipelineUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PipelineListParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies which page to retrieve.
	Page param.Field[string] `query:"page"`
	// Specifies the number of pipelines per page.
	PerPage param.Field[string] `query:"per_page"`
	// Specifies the prefix of pipeline name to search.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [PipelineListParams]'s query parameters as `url.Values`.
func (r PipelineListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PipelineDeleteParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PipelineGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PipelineGetResponseEnvelope struct {
	// Describes the configuration of a pipeline.
	Result PipelineGetResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                            `json:"success,required"`
	JSON    pipelineGetResponseEnvelopeJSON `json:"-"`
}

// pipelineGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineGetResponseEnvelope]
type pipelineGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
