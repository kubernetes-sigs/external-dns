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

// BucketDomainCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketDomainCustomService] method instead.
type BucketDomainCustomService struct {
	Options []option.RequestOption
}

// NewBucketDomainCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketDomainCustomService(opts ...option.RequestOption) (r *BucketDomainCustomService) {
	r = &BucketDomainCustomService{}
	r.Options = opts
	return
}

// Register a new custom domain for an existing R2 bucket.
func (r *BucketDomainCustomService) New(ctx context.Context, bucketName string, params BucketDomainCustomNewParams, opts ...option.RequestOption) (res *BucketDomainCustomNewResponse, err error) {
	var env BucketDomainCustomNewResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit the configuration for a custom domain on an existing R2 bucket.
func (r *BucketDomainCustomService) Update(ctx context.Context, bucketName string, domain string, params BucketDomainCustomUpdateParams, opts ...option.RequestOption) (res *BucketDomainCustomUpdateResponse, err error) {
	var env BucketDomainCustomUpdateResponseEnvelope
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", params.AccountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all custom domains registered with an existing R2 bucket.
func (r *BucketDomainCustomService) List(ctx context.Context, bucketName string, params BucketDomainCustomListParams, opts ...option.RequestOption) (res *BucketDomainCustomListResponse, err error) {
	var env BucketDomainCustomListResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove custom domain registration from an existing R2 bucket.
func (r *BucketDomainCustomService) Delete(ctx context.Context, bucketName string, domain string, params BucketDomainCustomDeleteParams, opts ...option.RequestOption) (res *BucketDomainCustomDeleteResponse, err error) {
	var env BucketDomainCustomDeleteResponseEnvelope
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", params.AccountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the configuration for a custom domain on an existing R2 bucket.
func (r *BucketDomainCustomService) Get(ctx context.Context, bucketName string, domain string, params BucketDomainCustomGetParams, opts ...option.RequestOption) (res *BucketDomainCustomGetResponse, err error) {
	var env BucketDomainCustomGetResponseEnvelope
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
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", params.AccountID, bucketName, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketDomainCustomNewResponse struct {
	// Domain name of the affected custom domain.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool `json:"enabled,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomNewResponseMinTLS `json:"minTLS"`
	JSON   bucketDomainCustomNewResponseJSON   `json:"-"`
}

// bucketDomainCustomNewResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomNewResponse]
type bucketDomainCustomNewResponseJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Ciphers     apijson.Field
	MinTLS      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomNewResponseMinTLS string

const (
	BucketDomainCustomNewResponseMinTLS1_0 BucketDomainCustomNewResponseMinTLS = "1.0"
	BucketDomainCustomNewResponseMinTLS1_1 BucketDomainCustomNewResponseMinTLS = "1.1"
	BucketDomainCustomNewResponseMinTLS1_2 BucketDomainCustomNewResponseMinTLS = "1.2"
	BucketDomainCustomNewResponseMinTLS1_3 BucketDomainCustomNewResponseMinTLS = "1.3"
)

func (r BucketDomainCustomNewResponseMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewResponseMinTLS1_0, BucketDomainCustomNewResponseMinTLS1_1, BucketDomainCustomNewResponseMinTLS1_2, BucketDomainCustomNewResponseMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomUpdateResponse struct {
	// Domain name of the affected custom domain.
	Domain string `json:"domain,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool `json:"enabled"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomUpdateResponseMinTLS `json:"minTLS"`
	JSON   bucketDomainCustomUpdateResponseJSON   `json:"-"`
}

// bucketDomainCustomUpdateResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomUpdateResponse]
type bucketDomainCustomUpdateResponseJSON struct {
	Domain      apijson.Field
	Ciphers     apijson.Field
	Enabled     apijson.Field
	MinTLS      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomUpdateResponseMinTLS string

const (
	BucketDomainCustomUpdateResponseMinTLS1_0 BucketDomainCustomUpdateResponseMinTLS = "1.0"
	BucketDomainCustomUpdateResponseMinTLS1_1 BucketDomainCustomUpdateResponseMinTLS = "1.1"
	BucketDomainCustomUpdateResponseMinTLS1_2 BucketDomainCustomUpdateResponseMinTLS = "1.2"
	BucketDomainCustomUpdateResponseMinTLS1_3 BucketDomainCustomUpdateResponseMinTLS = "1.3"
)

func (r BucketDomainCustomUpdateResponseMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomUpdateResponseMinTLS1_0, BucketDomainCustomUpdateResponseMinTLS1_1, BucketDomainCustomUpdateResponseMinTLS1_2, BucketDomainCustomUpdateResponseMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomListResponse struct {
	Domains []BucketDomainCustomListResponseDomain `json:"domains,required"`
	JSON    bucketDomainCustomListResponseJSON     `json:"-"`
}

// bucketDomainCustomListResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomListResponse]
type bucketDomainCustomListResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomListResponseDomain struct {
	// Domain name of the custom domain to be added.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool                                        `json:"enabled,required"`
	Status  BucketDomainCustomListResponseDomainsStatus `json:"status,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomListResponseDomainsMinTLS `json:"minTLS"`
	// Zone ID of the custom domain resides in.
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in.
	ZoneName string                                   `json:"zoneName"`
	JSON     bucketDomainCustomListResponseDomainJSON `json:"-"`
}

// bucketDomainCustomListResponseDomainJSON contains the JSON metadata for the
// struct [BucketDomainCustomListResponseDomain]
type bucketDomainCustomListResponseDomainJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	Ciphers     apijson.Field
	MinTLS      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseDomainJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomListResponseDomainsStatus struct {
	// Ownership status of the domain.
	Ownership BucketDomainCustomListResponseDomainsStatusOwnership `json:"ownership,required"`
	// SSL certificate status.
	SSL  BucketDomainCustomListResponseDomainsStatusSSL  `json:"ssl,required"`
	JSON bucketDomainCustomListResponseDomainsStatusJSON `json:"-"`
}

// bucketDomainCustomListResponseDomainsStatusJSON contains the JSON metadata for
// the struct [BucketDomainCustomListResponseDomainsStatus]
type bucketDomainCustomListResponseDomainsStatusJSON struct {
	Ownership   apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseDomainsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseDomainsStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain.
type BucketDomainCustomListResponseDomainsStatusOwnership string

const (
	BucketDomainCustomListResponseDomainsStatusOwnershipPending     BucketDomainCustomListResponseDomainsStatusOwnership = "pending"
	BucketDomainCustomListResponseDomainsStatusOwnershipActive      BucketDomainCustomListResponseDomainsStatusOwnership = "active"
	BucketDomainCustomListResponseDomainsStatusOwnershipDeactivated BucketDomainCustomListResponseDomainsStatusOwnership = "deactivated"
	BucketDomainCustomListResponseDomainsStatusOwnershipBlocked     BucketDomainCustomListResponseDomainsStatusOwnership = "blocked"
	BucketDomainCustomListResponseDomainsStatusOwnershipError       BucketDomainCustomListResponseDomainsStatusOwnership = "error"
	BucketDomainCustomListResponseDomainsStatusOwnershipUnknown     BucketDomainCustomListResponseDomainsStatusOwnership = "unknown"
)

func (r BucketDomainCustomListResponseDomainsStatusOwnership) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsStatusOwnershipPending, BucketDomainCustomListResponseDomainsStatusOwnershipActive, BucketDomainCustomListResponseDomainsStatusOwnershipDeactivated, BucketDomainCustomListResponseDomainsStatusOwnershipBlocked, BucketDomainCustomListResponseDomainsStatusOwnershipError, BucketDomainCustomListResponseDomainsStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status.
type BucketDomainCustomListResponseDomainsStatusSSL string

const (
	BucketDomainCustomListResponseDomainsStatusSSLInitializing BucketDomainCustomListResponseDomainsStatusSSL = "initializing"
	BucketDomainCustomListResponseDomainsStatusSSLPending      BucketDomainCustomListResponseDomainsStatusSSL = "pending"
	BucketDomainCustomListResponseDomainsStatusSSLActive       BucketDomainCustomListResponseDomainsStatusSSL = "active"
	BucketDomainCustomListResponseDomainsStatusSSLDeactivated  BucketDomainCustomListResponseDomainsStatusSSL = "deactivated"
	BucketDomainCustomListResponseDomainsStatusSSLError        BucketDomainCustomListResponseDomainsStatusSSL = "error"
	BucketDomainCustomListResponseDomainsStatusSSLUnknown      BucketDomainCustomListResponseDomainsStatusSSL = "unknown"
)

func (r BucketDomainCustomListResponseDomainsStatusSSL) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsStatusSSLInitializing, BucketDomainCustomListResponseDomainsStatusSSLPending, BucketDomainCustomListResponseDomainsStatusSSLActive, BucketDomainCustomListResponseDomainsStatusSSLDeactivated, BucketDomainCustomListResponseDomainsStatusSSLError, BucketDomainCustomListResponseDomainsStatusSSLUnknown:
		return true
	}
	return false
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomListResponseDomainsMinTLS string

const (
	BucketDomainCustomListResponseDomainsMinTLS1_0 BucketDomainCustomListResponseDomainsMinTLS = "1.0"
	BucketDomainCustomListResponseDomainsMinTLS1_1 BucketDomainCustomListResponseDomainsMinTLS = "1.1"
	BucketDomainCustomListResponseDomainsMinTLS1_2 BucketDomainCustomListResponseDomainsMinTLS = "1.2"
	BucketDomainCustomListResponseDomainsMinTLS1_3 BucketDomainCustomListResponseDomainsMinTLS = "1.3"
)

func (r BucketDomainCustomListResponseDomainsMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsMinTLS1_0, BucketDomainCustomListResponseDomainsMinTLS1_1, BucketDomainCustomListResponseDomainsMinTLS1_2, BucketDomainCustomListResponseDomainsMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomDeleteResponse struct {
	// Name of the removed custom domain.
	Domain string                               `json:"domain,required"`
	JSON   bucketDomainCustomDeleteResponseJSON `json:"-"`
}

// bucketDomainCustomDeleteResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomDeleteResponse]
type bucketDomainCustomDeleteResponseJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomGetResponse struct {
	// Domain name of the custom domain to be added.
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain.
	Enabled bool                                `json:"enabled,required"`
	Status  BucketDomainCustomGetResponseStatus `json:"status,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomGetResponseMinTLS `json:"minTLS"`
	// Zone ID of the custom domain resides in.
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in.
	ZoneName string                            `json:"zoneName"`
	JSON     bucketDomainCustomGetResponseJSON `json:"-"`
}

// bucketDomainCustomGetResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomGetResponse]
type bucketDomainCustomGetResponseJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	Ciphers     apijson.Field
	MinTLS      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomGetResponseStatus struct {
	// Ownership status of the domain.
	Ownership BucketDomainCustomGetResponseStatusOwnership `json:"ownership,required"`
	// SSL certificate status.
	SSL  BucketDomainCustomGetResponseStatusSSL  `json:"ssl,required"`
	JSON bucketDomainCustomGetResponseStatusJSON `json:"-"`
}

// bucketDomainCustomGetResponseStatusJSON contains the JSON metadata for the
// struct [BucketDomainCustomGetResponseStatus]
type bucketDomainCustomGetResponseStatusJSON struct {
	Ownership   apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomGetResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomGetResponseStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain.
type BucketDomainCustomGetResponseStatusOwnership string

const (
	BucketDomainCustomGetResponseStatusOwnershipPending     BucketDomainCustomGetResponseStatusOwnership = "pending"
	BucketDomainCustomGetResponseStatusOwnershipActive      BucketDomainCustomGetResponseStatusOwnership = "active"
	BucketDomainCustomGetResponseStatusOwnershipDeactivated BucketDomainCustomGetResponseStatusOwnership = "deactivated"
	BucketDomainCustomGetResponseStatusOwnershipBlocked     BucketDomainCustomGetResponseStatusOwnership = "blocked"
	BucketDomainCustomGetResponseStatusOwnershipError       BucketDomainCustomGetResponseStatusOwnership = "error"
	BucketDomainCustomGetResponseStatusOwnershipUnknown     BucketDomainCustomGetResponseStatusOwnership = "unknown"
)

func (r BucketDomainCustomGetResponseStatusOwnership) IsKnown() bool {
	switch r {
	case BucketDomainCustomGetResponseStatusOwnershipPending, BucketDomainCustomGetResponseStatusOwnershipActive, BucketDomainCustomGetResponseStatusOwnershipDeactivated, BucketDomainCustomGetResponseStatusOwnershipBlocked, BucketDomainCustomGetResponseStatusOwnershipError, BucketDomainCustomGetResponseStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status.
type BucketDomainCustomGetResponseStatusSSL string

const (
	BucketDomainCustomGetResponseStatusSSLInitializing BucketDomainCustomGetResponseStatusSSL = "initializing"
	BucketDomainCustomGetResponseStatusSSLPending      BucketDomainCustomGetResponseStatusSSL = "pending"
	BucketDomainCustomGetResponseStatusSSLActive       BucketDomainCustomGetResponseStatusSSL = "active"
	BucketDomainCustomGetResponseStatusSSLDeactivated  BucketDomainCustomGetResponseStatusSSL = "deactivated"
	BucketDomainCustomGetResponseStatusSSLError        BucketDomainCustomGetResponseStatusSSL = "error"
	BucketDomainCustomGetResponseStatusSSLUnknown      BucketDomainCustomGetResponseStatusSSL = "unknown"
)

func (r BucketDomainCustomGetResponseStatusSSL) IsKnown() bool {
	switch r {
	case BucketDomainCustomGetResponseStatusSSLInitializing, BucketDomainCustomGetResponseStatusSSLPending, BucketDomainCustomGetResponseStatusSSLActive, BucketDomainCustomGetResponseStatusSSLDeactivated, BucketDomainCustomGetResponseStatusSSLError, BucketDomainCustomGetResponseStatusSSLUnknown:
		return true
	}
	return false
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomGetResponseMinTLS string

const (
	BucketDomainCustomGetResponseMinTLS1_0 BucketDomainCustomGetResponseMinTLS = "1.0"
	BucketDomainCustomGetResponseMinTLS1_1 BucketDomainCustomGetResponseMinTLS = "1.1"
	BucketDomainCustomGetResponseMinTLS1_2 BucketDomainCustomGetResponseMinTLS = "1.2"
	BucketDomainCustomGetResponseMinTLS1_3 BucketDomainCustomGetResponseMinTLS = "1.3"
)

func (r BucketDomainCustomGetResponseMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomGetResponseMinTLS1_0, BucketDomainCustomGetResponseMinTLS1_1, BucketDomainCustomGetResponseMinTLS1_2, BucketDomainCustomGetResponseMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the custom domain to be added.
	Domain param.Field[string] `json:"domain,required"`
	// Whether to enable public bucket access at the custom domain. If undefined, the
	// domain will be enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Zone ID of the custom domain.
	ZoneID param.Field[string] `json:"zoneId,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS param.Field[BucketDomainCustomNewParamsMinTLS] `json:"minTLS"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainCustomNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketDomainCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomNewParamsMinTLS string

const (
	BucketDomainCustomNewParamsMinTLS1_0 BucketDomainCustomNewParamsMinTLS = "1.0"
	BucketDomainCustomNewParamsMinTLS1_1 BucketDomainCustomNewParamsMinTLS = "1.1"
	BucketDomainCustomNewParamsMinTLS1_2 BucketDomainCustomNewParamsMinTLS = "1.2"
	BucketDomainCustomNewParamsMinTLS1_3 BucketDomainCustomNewParamsMinTLS = "1.3"
)

func (r BucketDomainCustomNewParamsMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewParamsMinTLS1_0, BucketDomainCustomNewParamsMinTLS1_1, BucketDomainCustomNewParamsMinTLS1_2, BucketDomainCustomNewParamsMinTLS1_3:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainCustomNewParamsCfR2Jurisdiction string

const (
	BucketDomainCustomNewParamsCfR2JurisdictionDefault BucketDomainCustomNewParamsCfR2Jurisdiction = "default"
	BucketDomainCustomNewParamsCfR2JurisdictionEu      BucketDomainCustomNewParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomNewParamsCfR2JurisdictionFedramp BucketDomainCustomNewParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewParamsCfR2JurisdictionDefault, BucketDomainCustomNewParamsCfR2JurisdictionEu, BucketDomainCustomNewParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   BucketDomainCustomNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainCustomNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomNewResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomNewResponseEnvelope]
type bucketDomainCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainCustomNewResponseEnvelopeSuccess bool

const (
	BucketDomainCustomNewResponseEnvelopeSuccessTrue BucketDomainCustomNewResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainCustomUpdateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether to enable public bucket access at the specified custom domain.
	Enabled param.Field[bool] `json:"enabled"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to previous value.
	MinTLS param.Field[BucketDomainCustomUpdateParamsMinTLS] `json:"minTLS"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainCustomUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketDomainCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to previous value.
type BucketDomainCustomUpdateParamsMinTLS string

const (
	BucketDomainCustomUpdateParamsMinTLS1_0 BucketDomainCustomUpdateParamsMinTLS = "1.0"
	BucketDomainCustomUpdateParamsMinTLS1_1 BucketDomainCustomUpdateParamsMinTLS = "1.1"
	BucketDomainCustomUpdateParamsMinTLS1_2 BucketDomainCustomUpdateParamsMinTLS = "1.2"
	BucketDomainCustomUpdateParamsMinTLS1_3 BucketDomainCustomUpdateParamsMinTLS = "1.3"
)

func (r BucketDomainCustomUpdateParamsMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomUpdateParamsMinTLS1_0, BucketDomainCustomUpdateParamsMinTLS1_1, BucketDomainCustomUpdateParamsMinTLS1_2, BucketDomainCustomUpdateParamsMinTLS1_3:
		return true
	}
	return false
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainCustomUpdateParamsCfR2Jurisdiction string

const (
	BucketDomainCustomUpdateParamsCfR2JurisdictionDefault BucketDomainCustomUpdateParamsCfR2Jurisdiction = "default"
	BucketDomainCustomUpdateParamsCfR2JurisdictionEu      BucketDomainCustomUpdateParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomUpdateParamsCfR2JurisdictionFedramp BucketDomainCustomUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomUpdateParamsCfR2JurisdictionDefault, BucketDomainCustomUpdateParamsCfR2JurisdictionEu, BucketDomainCustomUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []string                         `json:"messages,required"`
	Result   BucketDomainCustomUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomUpdateResponseEnvelope]
type bucketDomainCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainCustomUpdateResponseEnvelopeSuccess bool

const (
	BucketDomainCustomUpdateResponseEnvelopeSuccessTrue BucketDomainCustomUpdateResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainCustomListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainCustomListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainCustomListParamsCfR2Jurisdiction string

const (
	BucketDomainCustomListParamsCfR2JurisdictionDefault BucketDomainCustomListParamsCfR2Jurisdiction = "default"
	BucketDomainCustomListParamsCfR2JurisdictionEu      BucketDomainCustomListParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomListParamsCfR2JurisdictionFedramp BucketDomainCustomListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomListParamsCfR2JurisdictionDefault, BucketDomainCustomListParamsCfR2JurisdictionEu, BucketDomainCustomListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomListResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []string                       `json:"messages,required"`
	Result   BucketDomainCustomListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainCustomListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomListResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomListResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomListResponseEnvelope]
type bucketDomainCustomListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainCustomListResponseEnvelopeSuccess bool

const (
	BucketDomainCustomListResponseEnvelopeSuccessTrue BucketDomainCustomListResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainCustomDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainCustomDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainCustomDeleteParamsCfR2Jurisdiction string

const (
	BucketDomainCustomDeleteParamsCfR2JurisdictionDefault BucketDomainCustomDeleteParamsCfR2Jurisdiction = "default"
	BucketDomainCustomDeleteParamsCfR2JurisdictionEu      BucketDomainCustomDeleteParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomDeleteParamsCfR2JurisdictionFedramp BucketDomainCustomDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomDeleteParamsCfR2JurisdictionDefault, BucketDomainCustomDeleteParamsCfR2JurisdictionEu, BucketDomainCustomDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []string                         `json:"messages,required"`
	Result   BucketDomainCustomDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomDeleteResponseEnvelope]
type bucketDomainCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainCustomDeleteResponseEnvelopeSuccess bool

const (
	BucketDomainCustomDeleteResponseEnvelopeSuccessTrue BucketDomainCustomDeleteResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainCustomGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketDomainCustomGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketDomainCustomGetParamsCfR2Jurisdiction string

const (
	BucketDomainCustomGetParamsCfR2JurisdictionDefault BucketDomainCustomGetParamsCfR2Jurisdiction = "default"
	BucketDomainCustomGetParamsCfR2JurisdictionEu      BucketDomainCustomGetParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomGetParamsCfR2JurisdictionFedramp BucketDomainCustomGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomGetParamsCfR2JurisdictionDefault, BucketDomainCustomGetParamsCfR2JurisdictionEu, BucketDomainCustomGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   BucketDomainCustomGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketDomainCustomGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomGetResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomGetResponseEnvelope]
type bucketDomainCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketDomainCustomGetResponseEnvelopeSuccess bool

const (
	BucketDomainCustomGetResponseEnvelopeSuccessTrue BucketDomainCustomGetResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
