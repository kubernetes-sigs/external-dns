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

// BucketDomainManagedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketDomainManagedService] method instead.
type BucketDomainManagedService struct {
	Options []option.RequestOption
}

// NewBucketDomainManagedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketDomainManagedService(opts ...option.RequestOption) (r *BucketDomainManagedService) {
	r = &BucketDomainManagedService{}
	r.Options = opts
	return
}

// Updates state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *BucketDomainManagedService) Update(ctx context.Context, bucketName string, params BucketDomainManagedUpdateParams, opts ...option.RequestOption) (res *BucketDomainManagedUpdateResponse, err error) {
	var env BucketDomainManagedUpdateResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *BucketDomainManagedService) List(ctx context.Context, bucketName string, params BucketDomainManagedListParams, opts ...option.RequestOption) (res *BucketDomainManagedListResponse, err error) {
	var env BucketDomainManagedListResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketDomainManagedUpdateResponse struct {
	// Bucket ID.
	BucketID string `json:"bucketId,required"`
	// Domain name of the bucket's r2.dev domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the r2.dev domain.
	Enabled bool                                  `json:"enabled,required"`
	JSON    bucketDomainManagedUpdateResponseJSON `json:"-"`
}

// bucketDomainManagedUpdateResponseJSON contains the JSON metadata for the struct
// [BucketDomainManagedUpdateResponse]
type bucketDomainManagedUpdateResponseJSON struct {
	BucketID    apijson.Field
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainManagedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainManagedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainManagedListResponse struct {
	// Bucket ID.
	BucketID string `json:"bucketId,required"`
	// Domain name of the bucket's r2.dev domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the r2.dev domain.
	Enabled bool                                `json:"enabled,required"`
	JSON    bucketDomainManagedListResponseJSON `json:"-"`
}

// bucketDomainManagedListResponseJSON contains the JSON metadata for the struct
// [BucketDomainManagedListResponse]
type bucketDomainManagedListResponseJSON struct {
	BucketID    apijson.Field
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainManagedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainManagedListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainManagedUpdateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to enable public bucket access at the r2.dev domain.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainManagedUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketDomainManagedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainManagedUpdateParamsCfR2Jurisdiction string

const (
	BucketDomainManagedUpdateParamsCfR2JurisdictionDefault BucketDomainManagedUpdateParamsCfR2Jurisdiction = "default"
	BucketDomainManagedUpdateParamsCfR2JurisdictionEu      BucketDomainManagedUpdateParamsCfR2Jurisdiction = "eu"
	BucketDomainManagedUpdateParamsCfR2JurisdictionFedramp BucketDomainManagedUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainManagedUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainManagedUpdateParamsCfR2JurisdictionDefault, BucketDomainManagedUpdateParamsCfR2JurisdictionEu, BucketDomainManagedUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainManagedUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo             `json:"errors,required"`
	Messages []string                          `json:"messages,required"`
	Result   BucketDomainManagedUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainManagedUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainManagedUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainManagedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainManagedUpdateResponseEnvelope]
type bucketDomainManagedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainManagedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainManagedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainManagedUpdateResponseEnvelopeSuccess bool

const (
	BucketDomainManagedUpdateResponseEnvelopeSuccessTrue BucketDomainManagedUpdateResponseEnvelopeSuccess = true
)

func (r BucketDomainManagedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainManagedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainManagedListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainManagedListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainManagedListParamsCfR2Jurisdiction string

const (
	BucketDomainManagedListParamsCfR2JurisdictionDefault BucketDomainManagedListParamsCfR2Jurisdiction = "default"
	BucketDomainManagedListParamsCfR2JurisdictionEu      BucketDomainManagedListParamsCfR2Jurisdiction = "eu"
	BucketDomainManagedListParamsCfR2JurisdictionFedramp BucketDomainManagedListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainManagedListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainManagedListParamsCfR2JurisdictionDefault, BucketDomainManagedListParamsCfR2JurisdictionEu, BucketDomainManagedListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainManagedListResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []string                        `json:"messages,required"`
	Result   BucketDomainManagedListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainManagedListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainManagedListResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainManagedListResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainManagedListResponseEnvelope]
type bucketDomainManagedListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainManagedListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainManagedListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainManagedListResponseEnvelopeSuccess bool

const (
	BucketDomainManagedListResponseEnvelopeSuccessTrue BucketDomainManagedListResponseEnvelopeSuccess = true
)

func (r BucketDomainManagedListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainManagedListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
