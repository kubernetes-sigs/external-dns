// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

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

// BucketMetricService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketMetricService] method instead.
type BucketMetricService struct {
	Options []option.RequestOption
}

// NewBucketMetricService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketMetricService(opts ...option.RequestOption) (r *BucketMetricService) {
	r = &BucketMetricService{}
	r.Options = opts
	return
}

// Get Storage/Object Count Metrics across all buckets in your account. Note that
// Account-Level Metrics may not immediately reflect the latest data.
func (r *BucketMetricService) List(ctx context.Context, query BucketMetricListParams, opts ...option.RequestOption) (res *BucketMetricListResponse, err error) {
	var env BucketMetricListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/metrics", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Metrics based on the class they belong to.
type BucketMetricListResponse struct {
	// Metrics based on what state they are in(uploaded or published).
	InfrequentAccess BucketMetricListResponseInfrequentAccess `json:"infrequentAccess"`
	// Metrics based on what state they are in(uploaded or published).
	Standard BucketMetricListResponseStandard `json:"standard"`
	JSON     bucketMetricListResponseJSON     `json:"-"`
}

// bucketMetricListResponseJSON contains the JSON metadata for the struct
// [BucketMetricListResponse]
type bucketMetricListResponseJSON struct {
	InfrequentAccess apijson.Field
	Standard         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BucketMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseJSON) RawJSON() string {
	return r.raw
}

// Metrics based on what state they are in(uploaded or published).
type BucketMetricListResponseInfrequentAccess struct {
	// Metrics on number of objects/amount of storage used.
	Published BucketMetricListResponseInfrequentAccessPublished `json:"published"`
	// Metrics on number of objects/amount of storage used.
	Uploaded BucketMetricListResponseInfrequentAccessUploaded `json:"uploaded"`
	JSON     bucketMetricListResponseInfrequentAccessJSON     `json:"-"`
}

// bucketMetricListResponseInfrequentAccessJSON contains the JSON metadata for the
// struct [BucketMetricListResponseInfrequentAccess]
type bucketMetricListResponseInfrequentAccessJSON struct {
	Published   apijson.Field
	Uploaded    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketMetricListResponseInfrequentAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseInfrequentAccessJSON) RawJSON() string {
	return r.raw
}

// Metrics on number of objects/amount of storage used.
type BucketMetricListResponseInfrequentAccessPublished struct {
	// Amount of.
	MetadataSize float64 `json:"metadataSize"`
	// Number of objects stored.
	Objects float64 `json:"objects"`
	// Amount of storage used by object data.
	PayloadSize float64                                               `json:"payloadSize"`
	JSON        bucketMetricListResponseInfrequentAccessPublishedJSON `json:"-"`
}

// bucketMetricListResponseInfrequentAccessPublishedJSON contains the JSON metadata
// for the struct [BucketMetricListResponseInfrequentAccessPublished]
type bucketMetricListResponseInfrequentAccessPublishedJSON struct {
	MetadataSize apijson.Field
	Objects      apijson.Field
	PayloadSize  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketMetricListResponseInfrequentAccessPublished) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseInfrequentAccessPublishedJSON) RawJSON() string {
	return r.raw
}

// Metrics on number of objects/amount of storage used.
type BucketMetricListResponseInfrequentAccessUploaded struct {
	// Amount of.
	MetadataSize float64 `json:"metadataSize"`
	// Number of objects stored.
	Objects float64 `json:"objects"`
	// Amount of storage used by object data.
	PayloadSize float64                                              `json:"payloadSize"`
	JSON        bucketMetricListResponseInfrequentAccessUploadedJSON `json:"-"`
}

// bucketMetricListResponseInfrequentAccessUploadedJSON contains the JSON metadata
// for the struct [BucketMetricListResponseInfrequentAccessUploaded]
type bucketMetricListResponseInfrequentAccessUploadedJSON struct {
	MetadataSize apijson.Field
	Objects      apijson.Field
	PayloadSize  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketMetricListResponseInfrequentAccessUploaded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseInfrequentAccessUploadedJSON) RawJSON() string {
	return r.raw
}

// Metrics based on what state they are in(uploaded or published).
type BucketMetricListResponseStandard struct {
	// Metrics on number of objects/amount of storage used.
	Published BucketMetricListResponseStandardPublished `json:"published"`
	// Metrics on number of objects/amount of storage used.
	Uploaded BucketMetricListResponseStandardUploaded `json:"uploaded"`
	JSON     bucketMetricListResponseStandardJSON     `json:"-"`
}

// bucketMetricListResponseStandardJSON contains the JSON metadata for the struct
// [BucketMetricListResponseStandard]
type bucketMetricListResponseStandardJSON struct {
	Published   apijson.Field
	Uploaded    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketMetricListResponseStandard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseStandardJSON) RawJSON() string {
	return r.raw
}

// Metrics on number of objects/amount of storage used.
type BucketMetricListResponseStandardPublished struct {
	// Amount of.
	MetadataSize float64 `json:"metadataSize"`
	// Number of objects stored.
	Objects float64 `json:"objects"`
	// Amount of storage used by object data.
	PayloadSize float64                                       `json:"payloadSize"`
	JSON        bucketMetricListResponseStandardPublishedJSON `json:"-"`
}

// bucketMetricListResponseStandardPublishedJSON contains the JSON metadata for the
// struct [BucketMetricListResponseStandardPublished]
type bucketMetricListResponseStandardPublishedJSON struct {
	MetadataSize apijson.Field
	Objects      apijson.Field
	PayloadSize  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketMetricListResponseStandardPublished) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseStandardPublishedJSON) RawJSON() string {
	return r.raw
}

// Metrics on number of objects/amount of storage used.
type BucketMetricListResponseStandardUploaded struct {
	// Amount of.
	MetadataSize float64 `json:"metadataSize"`
	// Number of objects stored.
	Objects float64 `json:"objects"`
	// Amount of storage used by object data.
	PayloadSize float64                                      `json:"payloadSize"`
	JSON        bucketMetricListResponseStandardUploadedJSON `json:"-"`
}

// bucketMetricListResponseStandardUploadedJSON contains the JSON metadata for the
// struct [BucketMetricListResponseStandardUploaded]
type bucketMetricListResponseStandardUploadedJSON struct {
	MetadataSize apijson.Field
	Objects      apijson.Field
	PayloadSize  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketMetricListResponseStandardUploaded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseStandardUploadedJSON) RawJSON() string {
	return r.raw
}

type BucketMetricListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type BucketMetricListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// Metrics based on the class they belong to.
	Result BucketMetricListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketMetricListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketMetricListResponseEnvelopeJSON    `json:"-"`
}

// bucketMetricListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketMetricListResponseEnvelope]
type bucketMetricListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketMetricListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketMetricListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketMetricListResponseEnvelopeSuccess bool

const (
	BucketMetricListResponseEnvelopeSuccessTrue BucketMetricListResponseEnvelopeSuccess = true
)

func (r BucketMetricListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketMetricListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
