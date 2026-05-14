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

// BucketSippyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketSippyService] method instead.
type BucketSippyService struct {
	Options []option.RequestOption
}

// NewBucketSippyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketSippyService(opts ...option.RequestOption) (r *BucketSippyService) {
	r = &BucketSippyService{}
	r.Options = opts
	return
}

// Sets configuration for Sippy for an existing R2 bucket.
func (r *BucketSippyService) Update(ctx context.Context, bucketName string, params BucketSippyUpdateParams, opts ...option.RequestOption) (res *Sippy, err error) {
	var env BucketSippyUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables Sippy on this bucket.
func (r *BucketSippyService) Delete(ctx context.Context, bucketName string, params BucketSippyDeleteParams, opts ...option.RequestOption) (res *BucketSippyDeleteResponse, err error) {
	var env BucketSippyDeleteResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets configuration for Sippy for an existing R2 bucket.
func (r *BucketSippyService) Get(ctx context.Context, bucketName string, params BucketSippyGetParams, opts ...option.RequestOption) (res *Sippy, err error) {
	var env BucketSippyGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Provider string

const (
	ProviderR2 Provider = "r2"
)

func (r Provider) IsKnown() bool {
	switch r {
	case ProviderR2:
		return true
	}
	return false
}

type Sippy struct {
	// Details about the configured destination bucket.
	Destination SippyDestination `json:"destination"`
	// State of Sippy for this bucket.
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket.
	Source SippySource `json:"source"`
	JSON   sippyJSON   `json:"-"`
}

// sippyJSON contains the JSON metadata for the struct [Sippy]
type sippyJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Sippy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket.
type SippyDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket.
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider.
	Bucket   string               `json:"bucket"`
	Provider Provider             `json:"provider"`
	JSON     sippyDestinationJSON `json:"-"`
}

// sippyDestinationJSON contains the JSON metadata for the struct
// [SippyDestination]
type sippyDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDestinationJSON) RawJSON() string {
	return r.raw
}

// Details about the configured source bucket.
type SippySource struct {
	// Name of the bucket on the provider.
	Bucket   string              `json:"bucket"`
	Provider SippySourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only).
	Region string          `json:"region,nullable"`
	JSON   sippySourceJSON `json:"-"`
}

// sippySourceJSON contains the JSON metadata for the struct [SippySource]
type sippySourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippySource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippySourceJSON) RawJSON() string {
	return r.raw
}

type SippySourceProvider string

const (
	SippySourceProviderAws SippySourceProvider = "aws"
	SippySourceProviderGcs SippySourceProvider = "gcs"
)

func (r SippySourceProvider) IsKnown() bool {
	switch r {
	case SippySourceProviderAws, SippySourceProviderGcs:
		return true
	}
	return false
}

type BucketSippyDeleteResponse struct {
	Enabled BucketSippyDeleteResponseEnabled `json:"enabled"`
	JSON    bucketSippyDeleteResponseJSON    `json:"-"`
}

// bucketSippyDeleteResponseJSON contains the JSON metadata for the struct
// [BucketSippyDeleteResponse]
type bucketSippyDeleteResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketSippyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketSippyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BucketSippyDeleteResponseEnabled bool

const (
	BucketSippyDeleteResponseEnabledFalse BucketSippyDeleteResponseEnabled = false
)

func (r BucketSippyDeleteResponseEnabled) IsKnown() bool {
	switch r {
	case BucketSippyDeleteResponseEnabledFalse:
		return true
	}
	return false
}

type BucketSippyUpdateParams struct {
	// Account ID.
	AccountID param.Field[string]              `path:"account_id,required"`
	Body      BucketSippyUpdateParamsBodyUnion `json:"body,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketSippyUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketSippyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type BucketSippyUpdateParamsBody struct {
	Destination param.Field[interface{}] `json:"destination"`
	Source      param.Field[interface{}] `json:"source"`
}

func (r BucketSippyUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketSippyUpdateParamsBody) implementsBucketSippyUpdateParamsBodyUnion() {}

// Satisfied by [r2.BucketSippyUpdateParamsBodyR2EnableSippyAws],
// [r2.BucketSippyUpdateParamsBodyR2EnableSippyGcs], [BucketSippyUpdateParamsBody].
type BucketSippyUpdateParamsBodyUnion interface {
	implementsBucketSippyUpdateParamsBodyUnion()
}

type BucketSippyUpdateParamsBodyR2EnableSippyAws struct {
	// R2 bucket to copy objects to.
	Destination param.Field[BucketSippyUpdateParamsBodyR2EnableSippyAwsDestination] `json:"destination"`
	// AWS S3 bucket to copy objects from.
	Source param.Field[BucketSippyUpdateParamsBodyR2EnableSippyAwsSource] `json:"source"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyAws) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyAws) implementsBucketSippyUpdateParamsBodyUnion() {}

// R2 bucket to copy objects to.
type BucketSippyUpdateParamsBodyR2EnableSippyAwsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API. token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]   `json:"accessKeyId"`
	Provider    param.Field[Provider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API. token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyAwsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AWS S3 bucket to copy objects from.
type BucketSippyUpdateParamsBodyR2EnableSippyAwsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket).
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the AWS S3 bucket.
	Bucket   param.Field[string]                                                    `json:"bucket"`
	Provider param.Field[BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone.
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket).
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyAwsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProvider string

const (
	BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProviderAws BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProvider = "aws"
)

func (r BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProvider) IsKnown() bool {
	switch r {
	case BucketSippyUpdateParamsBodyR2EnableSippyAwsSourceProviderAws:
		return true
	}
	return false
}

type BucketSippyUpdateParamsBodyR2EnableSippyGcs struct {
	// R2 bucket to copy objects to.
	Destination param.Field[BucketSippyUpdateParamsBodyR2EnableSippyGcsDestination] `json:"destination"`
	// GCS bucket to copy objects from.
	Source param.Field[BucketSippyUpdateParamsBodyR2EnableSippyGcsSource] `json:"source"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyGcs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyGcs) implementsBucketSippyUpdateParamsBodyUnion() {}

// R2 bucket to copy objects to.
type BucketSippyUpdateParamsBodyR2EnableSippyGcsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API. token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]   `json:"accessKeyId"`
	Provider    param.Field[Provider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API. token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyGcsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GCS bucket to copy objects from.
type BucketSippyUpdateParamsBodyR2EnableSippyGcsSource struct {
	// Name of the GCS bucket.
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket).
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket).
	PrivateKey param.Field[string]                                                    `json:"privateKey"`
	Provider   param.Field[BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProvider] `json:"provider"`
}

func (r BucketSippyUpdateParamsBodyR2EnableSippyGcsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProvider string

const (
	BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProviderGcs BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProvider = "gcs"
)

func (r BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProvider) IsKnown() bool {
	switch r {
	case BucketSippyUpdateParamsBodyR2EnableSippyGcsSourceProviderGcs:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketSippyUpdateParamsCfR2Jurisdiction string

const (
	BucketSippyUpdateParamsCfR2JurisdictionDefault BucketSippyUpdateParamsCfR2Jurisdiction = "default"
	BucketSippyUpdateParamsCfR2JurisdictionEu      BucketSippyUpdateParamsCfR2Jurisdiction = "eu"
	BucketSippyUpdateParamsCfR2JurisdictionFedramp BucketSippyUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketSippyUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketSippyUpdateParamsCfR2JurisdictionDefault, BucketSippyUpdateParamsCfR2JurisdictionEu, BucketSippyUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketSippyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   Sippy                 `json:"result,required"`
	// Whether the API call was successful.
	Success BucketSippyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketSippyUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketSippyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketSippyUpdateResponseEnvelope]
type bucketSippyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketSippyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketSippyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketSippyUpdateResponseEnvelopeSuccess bool

const (
	BucketSippyUpdateResponseEnvelopeSuccessTrue BucketSippyUpdateResponseEnvelopeSuccess = true
)

func (r BucketSippyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketSippyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketSippyDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketSippyDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketSippyDeleteParamsCfR2Jurisdiction string

const (
	BucketSippyDeleteParamsCfR2JurisdictionDefault BucketSippyDeleteParamsCfR2Jurisdiction = "default"
	BucketSippyDeleteParamsCfR2JurisdictionEu      BucketSippyDeleteParamsCfR2Jurisdiction = "eu"
	BucketSippyDeleteParamsCfR2JurisdictionFedramp BucketSippyDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketSippyDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketSippyDeleteParamsCfR2JurisdictionDefault, BucketSippyDeleteParamsCfR2JurisdictionEu, BucketSippyDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketSippyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []string                  `json:"messages,required"`
	Result   BucketSippyDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketSippyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketSippyDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketSippyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketSippyDeleteResponseEnvelope]
type bucketSippyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketSippyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketSippyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketSippyDeleteResponseEnvelopeSuccess bool

const (
	BucketSippyDeleteResponseEnvelopeSuccessTrue BucketSippyDeleteResponseEnvelopeSuccess = true
)

func (r BucketSippyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketSippyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketSippyGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketSippyGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketSippyGetParamsCfR2Jurisdiction string

const (
	BucketSippyGetParamsCfR2JurisdictionDefault BucketSippyGetParamsCfR2Jurisdiction = "default"
	BucketSippyGetParamsCfR2JurisdictionEu      BucketSippyGetParamsCfR2Jurisdiction = "eu"
	BucketSippyGetParamsCfR2JurisdictionFedramp BucketSippyGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketSippyGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketSippyGetParamsCfR2JurisdictionDefault, BucketSippyGetParamsCfR2JurisdictionEu, BucketSippyGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketSippyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   Sippy                 `json:"result,required"`
	// Whether the API call was successful.
	Success BucketSippyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketSippyGetResponseEnvelopeJSON    `json:"-"`
}

// bucketSippyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketSippyGetResponseEnvelope]
type bucketSippyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketSippyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketSippyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketSippyGetResponseEnvelopeSuccess bool

const (
	BucketSippyGetResponseEnvelopeSuccessTrue BucketSippyGetResponseEnvelopeSuccess = true
)

func (r BucketSippyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketSippyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
