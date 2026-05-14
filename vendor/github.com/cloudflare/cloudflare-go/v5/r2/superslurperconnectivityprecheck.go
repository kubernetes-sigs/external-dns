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

// SuperSlurperConnectivityPrecheckService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperSlurperConnectivityPrecheckService] method instead.
type SuperSlurperConnectivityPrecheckService struct {
	Options []option.RequestOption
}

// NewSuperSlurperConnectivityPrecheckService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSuperSlurperConnectivityPrecheckService(opts ...option.RequestOption) (r *SuperSlurperConnectivityPrecheckService) {
	r = &SuperSlurperConnectivityPrecheckService{}
	r.Options = opts
	return
}

// Check whether tokens are valid against the source bucket
func (r *SuperSlurperConnectivityPrecheckService) Source(ctx context.Context, params SuperSlurperConnectivityPrecheckSourceParams, opts ...option.RequestOption) (res *SuperSlurperConnectivityPrecheckSourceResponse, err error) {
	var env SuperSlurperConnectivityPrecheckSourceResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/source/connectivity-precheck", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Check whether tokens are valid against the target bucket
func (r *SuperSlurperConnectivityPrecheckService) Target(ctx context.Context, params SuperSlurperConnectivityPrecheckTargetParams, opts ...option.RequestOption) (res *SuperSlurperConnectivityPrecheckTargetResponse, err error) {
	var env SuperSlurperConnectivityPrecheckTargetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/target/connectivity-precheck", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SuperSlurperConnectivityPrecheckSourceResponse struct {
	ConnectivityStatus SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatus `json:"connectivityStatus"`
	JSON               superSlurperConnectivityPrecheckSourceResponseJSON               `json:"-"`
}

// superSlurperConnectivityPrecheckSourceResponseJSON contains the JSON metadata
// for the struct [SuperSlurperConnectivityPrecheckSourceResponse]
type superSlurperConnectivityPrecheckSourceResponseJSON struct {
	ConnectivityStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SuperSlurperConnectivityPrecheckSourceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperConnectivityPrecheckSourceResponseJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatus string

const (
	SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatusSuccess SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatus = "success"
	SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatusError   SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatus = "error"
)

func (r SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatus) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatusSuccess, SuperSlurperConnectivityPrecheckSourceResponseConnectivityStatusError:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckTargetResponse struct {
	ConnectivityStatus SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatus `json:"connectivityStatus"`
	JSON               superSlurperConnectivityPrecheckTargetResponseJSON               `json:"-"`
}

// superSlurperConnectivityPrecheckTargetResponseJSON contains the JSON metadata
// for the struct [SuperSlurperConnectivityPrecheckTargetResponse]
type superSlurperConnectivityPrecheckTargetResponseJSON struct {
	ConnectivityStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SuperSlurperConnectivityPrecheckTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperConnectivityPrecheckTargetResponseJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatus string

const (
	SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatusSuccess SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatus = "success"
	SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatusError   SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatus = "error"
)

func (r SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatus) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatusSuccess, SuperSlurperConnectivityPrecheckTargetResponseConnectivityStatusError:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceParams struct {
	AccountID param.Field[string]                                   `path:"account_id,required"`
	Body      SuperSlurperConnectivityPrecheckSourceParamsBodyUnion `json:"body,required"`
}

func (r SuperSlurperConnectivityPrecheckSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SuperSlurperConnectivityPrecheckSourceParamsBody struct {
	Bucket       param.Field[string]                                                       `json:"bucket"`
	Endpoint     param.Field[string]                                                       `json:"endpoint"`
	Jurisdiction param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[interface{}]                                                  `json:"secret"`
	Vendor       param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyVendor]       `json:"vendor"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBody) implementsSuperSlurperConnectivityPrecheckSourceParamsBodyUnion() {
}

// Satisfied by
// [r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchema],
// [r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchema],
// [r2.SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchema],
// [SuperSlurperConnectivityPrecheckSourceParamsBody].
type SuperSlurperConnectivityPrecheckSourceParamsBodyUnion interface {
	implementsSuperSlurperConnectivityPrecheckSourceParamsBodyUnion()
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchema struct {
	Bucket   param.Field[string]                                                                        `json:"bucket"`
	Endpoint param.Field[string]                                                                        `json:"endpoint"`
	Secret   param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaSecret] `json:"secret"`
	Vendor   param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendor] `json:"vendor"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchema) implementsSuperSlurperConnectivityPrecheckSourceParamsBodyUnion() {
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendor string

const (
	SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendorS3 SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendor = "s3"
)

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperS3SourceSchemaVendorS3:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchema struct {
	Bucket param.Field[string]                                                                         `json:"bucket"`
	Secret param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaSecret] `json:"secret"`
	Vendor param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendor] `json:"vendor"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchema) implementsSuperSlurperConnectivityPrecheckSourceParamsBodyUnion() {
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaSecret struct {
	ClientEmail param.Field[string] `json:"clientEmail"`
	PrivateKey  param.Field[string] `json:"privateKey"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendor string

const (
	SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendorGcs SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendor = "gcs"
)

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperGcsSourceSchemaVendorGcs:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchema struct {
	Bucket       param.Field[string]                                                                              `json:"bucket"`
	Jurisdiction param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaSecret]       `json:"secret"`
	Vendor       param.Field[Provider]                                                                            `json:"vendor"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchema) implementsSuperSlurperConnectivityPrecheckSourceParamsBodyUnion() {
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction string

const (
	SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionDefault SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction = "default"
	SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionEu      SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction = "eu"
	SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionFedramp SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction = "fedramp"
)

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionDefault, SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionEu, SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyR2SlurperR2SourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction string

const (
	SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionDefault SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction = "default"
	SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionEu      SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction = "eu"
	SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionFedramp SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction = "fedramp"
)

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionDefault, SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionEu, SuperSlurperConnectivityPrecheckSourceParamsBodyJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceParamsBodyVendor string

const (
	SuperSlurperConnectivityPrecheckSourceParamsBodyVendorS3  SuperSlurperConnectivityPrecheckSourceParamsBodyVendor = "s3"
	SuperSlurperConnectivityPrecheckSourceParamsBodyVendorGcs SuperSlurperConnectivityPrecheckSourceParamsBodyVendor = "gcs"
	SuperSlurperConnectivityPrecheckSourceParamsBodyVendorR2  SuperSlurperConnectivityPrecheckSourceParamsBodyVendor = "r2"
)

func (r SuperSlurperConnectivityPrecheckSourceParamsBodyVendor) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceParamsBodyVendorS3, SuperSlurperConnectivityPrecheckSourceParamsBodyVendorGcs, SuperSlurperConnectivityPrecheckSourceParamsBodyVendorR2:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckSourceResponseEnvelope struct {
	Errors   []shared.ResponseInfo                          `json:"errors"`
	Messages []string                                       `json:"messages"`
	Result   SuperSlurperConnectivityPrecheckSourceResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccess `json:"success"`
	JSON    superSlurperConnectivityPrecheckSourceResponseEnvelopeJSON    `json:"-"`
}

// superSlurperConnectivityPrecheckSourceResponseEnvelopeJSON contains the JSON
// metadata for the struct [SuperSlurperConnectivityPrecheckSourceResponseEnvelope]
type superSlurperConnectivityPrecheckSourceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperConnectivityPrecheckSourceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperConnectivityPrecheckSourceResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccess bool

const (
	SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccessTrue SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccess = true
)

func (r SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckSourceResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckTargetParams struct {
	AccountID    param.Field[string]                                                   `path:"account_id,required"`
	Bucket       param.Field[string]                                                   `json:"bucket"`
	Jurisdiction param.Field[SuperSlurperConnectivityPrecheckTargetParamsJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[SuperSlurperConnectivityPrecheckTargetParamsSecret]       `json:"secret"`
	Vendor       param.Field[Provider]                                                 `json:"vendor"`
}

func (r SuperSlurperConnectivityPrecheckTargetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperConnectivityPrecheckTargetParamsJurisdiction string

const (
	SuperSlurperConnectivityPrecheckTargetParamsJurisdictionDefault SuperSlurperConnectivityPrecheckTargetParamsJurisdiction = "default"
	SuperSlurperConnectivityPrecheckTargetParamsJurisdictionEu      SuperSlurperConnectivityPrecheckTargetParamsJurisdiction = "eu"
	SuperSlurperConnectivityPrecheckTargetParamsJurisdictionFedramp SuperSlurperConnectivityPrecheckTargetParamsJurisdiction = "fedramp"
)

func (r SuperSlurperConnectivityPrecheckTargetParamsJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckTargetParamsJurisdictionDefault, SuperSlurperConnectivityPrecheckTargetParamsJurisdictionEu, SuperSlurperConnectivityPrecheckTargetParamsJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperConnectivityPrecheckTargetParamsSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperConnectivityPrecheckTargetParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperConnectivityPrecheckTargetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                          `json:"errors"`
	Messages []string                                       `json:"messages"`
	Result   SuperSlurperConnectivityPrecheckTargetResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccess `json:"success"`
	JSON    superSlurperConnectivityPrecheckTargetResponseEnvelopeJSON    `json:"-"`
}

// superSlurperConnectivityPrecheckTargetResponseEnvelopeJSON contains the JSON
// metadata for the struct [SuperSlurperConnectivityPrecheckTargetResponseEnvelope]
type superSlurperConnectivityPrecheckTargetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperConnectivityPrecheckTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperConnectivityPrecheckTargetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccess bool

const (
	SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccessTrue SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccess = true
)

func (r SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SuperSlurperConnectivityPrecheckTargetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
