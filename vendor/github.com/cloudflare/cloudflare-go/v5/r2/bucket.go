// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// BucketService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
type BucketService struct {
	Options            []option.RequestOption
	Lifecycle          *BucketLifecycleService
	CORS               *BucketCORSService
	Domains            *BucketDomainService
	EventNotifications *BucketEventNotificationService
	Locks              *BucketLockService
	Metrics            *BucketMetricService
	Sippy              *BucketSippyService
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r *BucketService) {
	r = &BucketService{}
	r.Options = opts
	r.Lifecycle = NewBucketLifecycleService(opts...)
	r.CORS = NewBucketCORSService(opts...)
	r.Domains = NewBucketDomainService(opts...)
	r.EventNotifications = NewBucketEventNotificationService(opts...)
	r.Locks = NewBucketLockService(opts...)
	r.Metrics = NewBucketMetricService(opts...)
	r.Sippy = NewBucketSippyService(opts...)
	return
}

// Creates a new R2 bucket.
func (r *BucketService) New(ctx context.Context, params BucketNewParams, opts ...option.RequestOption) (res *Bucket, err error) {
	var env BucketNewResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all R2 buckets on your account.
func (r *BucketService) List(ctx context.Context, params BucketListParams, opts ...option.RequestOption) (res *BucketListResponse, err error) {
	var env BucketListResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing R2 bucket.
func (r *BucketService) Delete(ctx context.Context, bucketName string, params BucketDeleteParams, opts ...option.RequestOption) (res *BucketDeleteResponse, err error) {
	var env BucketDeleteResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates properties of an existing R2 bucket.
func (r *BucketService) Edit(ctx context.Context, bucketName string, params BucketEditParams, opts ...option.RequestOption) (res *Bucket, err error) {
	var env BucketEditResponseEnvelope
	if params.StorageClass.Present {
		opts = append(opts, option.WithHeader("cf-r2-storage-class", fmt.Sprintf("%s", params.StorageClass)))
	}
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets properties of an existing R2 bucket.
func (r *BucketService) Get(ctx context.Context, bucketName string, params BucketGetParams, opts ...option.RequestOption) (res *Bucket, err error) {
	var env BucketGetResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single R2 bucket.
type Bucket struct {
	// Creation timestamp.
	CreationDate string `json:"creation_date"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction BucketJurisdiction `json:"jurisdiction"`
	// Location of the bucket.
	Location BucketLocation `json:"location"`
	// Name of the bucket.
	Name string `json:"name"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass BucketStorageClass `json:"storage_class"`
	JSON         bucketJSON         `json:"-"`
}

// bucketJSON contains the JSON metadata for the struct [Bucket]
type bucketJSON struct {
	CreationDate apijson.Field
	Jurisdiction apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJSON) RawJSON() string {
	return r.raw
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketJurisdiction string

const (
	BucketJurisdictionDefault BucketJurisdiction = "default"
	BucketJurisdictionEu      BucketJurisdiction = "eu"
	BucketJurisdictionFedramp BucketJurisdiction = "fedramp"
)

func (r BucketJurisdiction) IsKnown() bool {
	switch r {
	case BucketJurisdictionDefault, BucketJurisdictionEu, BucketJurisdictionFedramp:
		return true
	}
	return false
}

// Location of the bucket.
type BucketLocation string

const (
	BucketLocationApac BucketLocation = "apac"
	BucketLocationEeur BucketLocation = "eeur"
	BucketLocationEnam BucketLocation = "enam"
	BucketLocationWeur BucketLocation = "weur"
	BucketLocationWnam BucketLocation = "wnam"
	BucketLocationOc   BucketLocation = "oc"
)

func (r BucketLocation) IsKnown() bool {
	switch r {
	case BucketLocationApac, BucketLocationEeur, BucketLocationEnam, BucketLocationWeur, BucketLocationWnam, BucketLocationOc:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketStorageClass string

const (
	BucketStorageClassStandard         BucketStorageClass = "Standard"
	BucketStorageClassInfrequentAccess BucketStorageClass = "InfrequentAccess"
)

func (r BucketStorageClass) IsKnown() bool {
	switch r {
	case BucketStorageClassStandard, BucketStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketListResponse struct {
	Buckets []Bucket               `json:"buckets"`
	JSON    bucketListResponseJSON `json:"-"`
}

// bucketListResponseJSON contains the JSON metadata for the struct
// [BucketListResponse]
type bucketListResponseJSON struct {
	Buckets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDeleteResponse = interface{}

type BucketNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the bucket.
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket.
	LocationHint param.Field[BucketNewParamsLocationHint] `json:"locationHint"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass param.Field[BucketNewParamsStorageClass] `json:"storageClass"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of the bucket.
type BucketNewParamsLocationHint string

const (
	BucketNewParamsLocationHintApac BucketNewParamsLocationHint = "apac"
	BucketNewParamsLocationHintEeur BucketNewParamsLocationHint = "eeur"
	BucketNewParamsLocationHintEnam BucketNewParamsLocationHint = "enam"
	BucketNewParamsLocationHintWeur BucketNewParamsLocationHint = "weur"
	BucketNewParamsLocationHintWnam BucketNewParamsLocationHint = "wnam"
	BucketNewParamsLocationHintOc   BucketNewParamsLocationHint = "oc"
)

func (r BucketNewParamsLocationHint) IsKnown() bool {
	switch r {
	case BucketNewParamsLocationHintApac, BucketNewParamsLocationHintEeur, BucketNewParamsLocationHintEnam, BucketNewParamsLocationHintWeur, BucketNewParamsLocationHintWnam, BucketNewParamsLocationHintOc:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketNewParamsStorageClass string

const (
	BucketNewParamsStorageClassStandard         BucketNewParamsStorageClass = "Standard"
	BucketNewParamsStorageClassInfrequentAccess BucketNewParamsStorageClass = "InfrequentAccess"
)

func (r BucketNewParamsStorageClass) IsKnown() bool {
	switch r {
	case BucketNewParamsStorageClassStandard, BucketNewParamsStorageClassInfrequentAccess:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketNewParamsCfR2Jurisdiction string

const (
	BucketNewParamsCfR2JurisdictionDefault BucketNewParamsCfR2Jurisdiction = "default"
	BucketNewParamsCfR2JurisdictionEu      BucketNewParamsCfR2Jurisdiction = "eu"
	BucketNewParamsCfR2JurisdictionFedramp BucketNewParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketNewParamsCfR2JurisdictionDefault, BucketNewParamsCfR2JurisdictionEu, BucketNewParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// A single R2 bucket.
	Result Bucket `json:"result,required"`
	// Whether the API call was successful.
	Success BucketNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketNewResponseEnvelopeJSON    `json:"-"`
}

// bucketNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketNewResponseEnvelope]
type bucketNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketNewResponseEnvelopeSuccess bool

const (
	BucketNewResponseEnvelopeSuccessTrue BucketNewResponseEnvelopeSuccess = true
)

func (r BucketNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Pagination cursor received during the last List Buckets call. R2 buckets are
	// paginated using cursors instead of page numbers.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order buckets.
	Direction param.Field[BucketListParamsDirection] `query:"direction"`
	// Bucket names to filter by. Only buckets with this phrase in their name will be
	// returned.
	NameContains param.Field[string] `query:"name_contains"`
	// Field to order buckets by.
	Order param.Field[BucketListParamsOrder] `query:"order"`
	// Maximum number of buckets to return in a single call.
	PerPage param.Field[float64] `query:"per_page"`
	// Bucket name to start searching after. Buckets are ordered lexicographically.
	StartAfter param.Field[string] `query:"start_after"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// URLQuery serializes [BucketListParams]'s query parameters as `url.Values`.
func (r BucketListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order buckets.
type BucketListParamsDirection string

const (
	BucketListParamsDirectionAsc  BucketListParamsDirection = "asc"
	BucketListParamsDirectionDesc BucketListParamsDirection = "desc"
)

func (r BucketListParamsDirection) IsKnown() bool {
	switch r {
	case BucketListParamsDirectionAsc, BucketListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order buckets by.
type BucketListParamsOrder string

const (
	BucketListParamsOrderName BucketListParamsOrder = "name"
)

func (r BucketListParamsOrder) IsKnown() bool {
	switch r {
	case BucketListParamsOrderName:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketListParamsCfR2Jurisdiction string

const (
	BucketListParamsCfR2JurisdictionDefault BucketListParamsCfR2Jurisdiction = "default"
	BucketListParamsCfR2JurisdictionEu      BucketListParamsCfR2Jurisdiction = "eu"
	BucketListParamsCfR2JurisdictionFedramp BucketListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketListParamsCfR2JurisdictionDefault, BucketListParamsCfR2JurisdictionEu, BucketListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   BucketListResponse    `json:"result,required"`
	// Whether the API call was successful.
	Success    BucketListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo BucketListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       bucketListResponseEnvelopeJSON       `json:"-"`
}

// bucketListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketListResponseEnvelope]
type bucketListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketListResponseEnvelopeSuccess bool

const (
	BucketListResponseEnvelopeSuccessTrue BucketListResponseEnvelopeSuccess = true
)

func (r BucketListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketListResponseEnvelopeResultInfo struct {
	// A continuation token that should be used to fetch the next page of results.
	Cursor string `json:"cursor"`
	// Maximum number of results on this page.
	PerPage float64                                  `json:"per_page"`
	JSON    bucketListResponseEnvelopeResultInfoJSON `json:"-"`
}

// bucketListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [BucketListResponseEnvelopeResultInfo]
type bucketListResponseEnvelopeResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BucketDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDeleteParamsCfR2Jurisdiction string

const (
	BucketDeleteParamsCfR2JurisdictionDefault BucketDeleteParamsCfR2Jurisdiction = "default"
	BucketDeleteParamsCfR2JurisdictionEu      BucketDeleteParamsCfR2Jurisdiction = "eu"
	BucketDeleteParamsCfR2JurisdictionFedramp BucketDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDeleteParamsCfR2JurisdictionDefault, BucketDeleteParamsCfR2JurisdictionEu, BucketDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   BucketDeleteResponse  `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketDeleteResponseEnvelope]
type bucketDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDeleteResponseEnvelopeSuccess bool

const (
	BucketDeleteResponseEnvelopeSuccessTrue BucketDeleteResponseEnvelopeSuccess = true
)

func (r BucketDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEditParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass param.Field[BucketEditParamsCfR2StorageClass] `header:"cf-r2-storage-class,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEditParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketEditParamsCfR2StorageClass string

const (
	BucketEditParamsCfR2StorageClassStandard         BucketEditParamsCfR2StorageClass = "Standard"
	BucketEditParamsCfR2StorageClassInfrequentAccess BucketEditParamsCfR2StorageClass = "InfrequentAccess"
)

func (r BucketEditParamsCfR2StorageClass) IsKnown() bool {
	switch r {
	case BucketEditParamsCfR2StorageClassStandard, BucketEditParamsCfR2StorageClassInfrequentAccess:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketEditParamsCfR2Jurisdiction string

const (
	BucketEditParamsCfR2JurisdictionDefault BucketEditParamsCfR2Jurisdiction = "default"
	BucketEditParamsCfR2JurisdictionEu      BucketEditParamsCfR2Jurisdiction = "eu"
	BucketEditParamsCfR2JurisdictionFedramp BucketEditParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEditParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEditParamsCfR2JurisdictionDefault, BucketEditParamsCfR2JurisdictionEu, BucketEditParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// A single R2 bucket.
	Result Bucket `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEditResponseEnvelopeJSON    `json:"-"`
}

// bucketEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketEditResponseEnvelope]
type bucketEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEditResponseEnvelopeSuccess bool

const (
	BucketEditResponseEnvelopeSuccessTrue BucketEditResponseEnvelopeSuccess = true
)

func (r BucketEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketGetParamsCfR2Jurisdiction string

const (
	BucketGetParamsCfR2JurisdictionDefault BucketGetParamsCfR2Jurisdiction = "default"
	BucketGetParamsCfR2JurisdictionEu      BucketGetParamsCfR2Jurisdiction = "eu"
	BucketGetParamsCfR2JurisdictionFedramp BucketGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketGetParamsCfR2JurisdictionDefault, BucketGetParamsCfR2JurisdictionEu, BucketGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// A single R2 bucket.
	Result Bucket `json:"result,required"`
	// Whether the API call was successful.
	Success BucketGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketGetResponseEnvelopeJSON    `json:"-"`
}

// bucketGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketGetResponseEnvelope]
type bucketGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketGetResponseEnvelopeSuccess bool

const (
	BucketGetResponseEnvelopeSuccessTrue BucketGetResponseEnvelopeSuccess = true
)

func (r BucketGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
