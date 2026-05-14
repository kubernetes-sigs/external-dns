// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// FallbackOriginService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFallbackOriginService] method instead.
type FallbackOriginService struct {
	Options []option.RequestOption
}

// NewFallbackOriginService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFallbackOriginService(opts ...option.RequestOption) (r *FallbackOriginService) {
	r = &FallbackOriginService{}
	r.Options = opts
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Update(ctx context.Context, params FallbackOriginUpdateParams, opts ...option.RequestOption) (res *FallbackOriginUpdateResponse, err error) {
	var env FallbackOriginUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Delete(ctx context.Context, body FallbackOriginDeleteParams, opts ...option.RequestOption) (res *FallbackOriginDeleteResponse, err error) {
	var env FallbackOriginDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Get(ctx context.Context, query FallbackOriginGetParams, opts ...option.RequestOption) (res *FallbackOriginGetResponse, err error) {
	var env FallbackOriginGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FallbackOriginUpdateResponse struct {
	// This is the time the fallback origin was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are errors that were encountered while trying to activate a fallback
	// origin.
	Errors []string `json:"errors"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin string `json:"origin"`
	// Status of the fallback origin's activation.
	Status FallbackOriginUpdateResponseStatus `json:"status"`
	// This is the time the fallback origin was updated.
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      fallbackOriginUpdateResponseJSON `json:"-"`
}

// fallbackOriginUpdateResponseJSON contains the JSON metadata for the struct
// [FallbackOriginUpdateResponse]
type fallbackOriginUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Errors      apijson.Field
	Origin      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the fallback origin's activation.
type FallbackOriginUpdateResponseStatus string

const (
	FallbackOriginUpdateResponseStatusInitializing       FallbackOriginUpdateResponseStatus = "initializing"
	FallbackOriginUpdateResponseStatusPendingDeployment  FallbackOriginUpdateResponseStatus = "pending_deployment"
	FallbackOriginUpdateResponseStatusPendingDeletion    FallbackOriginUpdateResponseStatus = "pending_deletion"
	FallbackOriginUpdateResponseStatusActive             FallbackOriginUpdateResponseStatus = "active"
	FallbackOriginUpdateResponseStatusDeploymentTimedOut FallbackOriginUpdateResponseStatus = "deployment_timed_out"
	FallbackOriginUpdateResponseStatusDeletionTimedOut   FallbackOriginUpdateResponseStatus = "deletion_timed_out"
)

func (r FallbackOriginUpdateResponseStatus) IsKnown() bool {
	switch r {
	case FallbackOriginUpdateResponseStatusInitializing, FallbackOriginUpdateResponseStatusPendingDeployment, FallbackOriginUpdateResponseStatusPendingDeletion, FallbackOriginUpdateResponseStatusActive, FallbackOriginUpdateResponseStatusDeploymentTimedOut, FallbackOriginUpdateResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type FallbackOriginDeleteResponse struct {
	// This is the time the fallback origin was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are errors that were encountered while trying to activate a fallback
	// origin.
	Errors []string `json:"errors"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin string `json:"origin"`
	// Status of the fallback origin's activation.
	Status FallbackOriginDeleteResponseStatus `json:"status"`
	// This is the time the fallback origin was updated.
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      fallbackOriginDeleteResponseJSON `json:"-"`
}

// fallbackOriginDeleteResponseJSON contains the JSON metadata for the struct
// [FallbackOriginDeleteResponse]
type fallbackOriginDeleteResponseJSON struct {
	CreatedAt   apijson.Field
	Errors      apijson.Field
	Origin      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the fallback origin's activation.
type FallbackOriginDeleteResponseStatus string

const (
	FallbackOriginDeleteResponseStatusInitializing       FallbackOriginDeleteResponseStatus = "initializing"
	FallbackOriginDeleteResponseStatusPendingDeployment  FallbackOriginDeleteResponseStatus = "pending_deployment"
	FallbackOriginDeleteResponseStatusPendingDeletion    FallbackOriginDeleteResponseStatus = "pending_deletion"
	FallbackOriginDeleteResponseStatusActive             FallbackOriginDeleteResponseStatus = "active"
	FallbackOriginDeleteResponseStatusDeploymentTimedOut FallbackOriginDeleteResponseStatus = "deployment_timed_out"
	FallbackOriginDeleteResponseStatusDeletionTimedOut   FallbackOriginDeleteResponseStatus = "deletion_timed_out"
)

func (r FallbackOriginDeleteResponseStatus) IsKnown() bool {
	switch r {
	case FallbackOriginDeleteResponseStatusInitializing, FallbackOriginDeleteResponseStatusPendingDeployment, FallbackOriginDeleteResponseStatusPendingDeletion, FallbackOriginDeleteResponseStatusActive, FallbackOriginDeleteResponseStatusDeploymentTimedOut, FallbackOriginDeleteResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type FallbackOriginGetResponse struct {
	// This is the time the fallback origin was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are errors that were encountered while trying to activate a fallback
	// origin.
	Errors []string `json:"errors"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin string `json:"origin"`
	// Status of the fallback origin's activation.
	Status FallbackOriginGetResponseStatus `json:"status"`
	// This is the time the fallback origin was updated.
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      fallbackOriginGetResponseJSON `json:"-"`
}

// fallbackOriginGetResponseJSON contains the JSON metadata for the struct
// [FallbackOriginGetResponse]
type fallbackOriginGetResponseJSON struct {
	CreatedAt   apijson.Field
	Errors      apijson.Field
	Origin      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the fallback origin's activation.
type FallbackOriginGetResponseStatus string

const (
	FallbackOriginGetResponseStatusInitializing       FallbackOriginGetResponseStatus = "initializing"
	FallbackOriginGetResponseStatusPendingDeployment  FallbackOriginGetResponseStatus = "pending_deployment"
	FallbackOriginGetResponseStatusPendingDeletion    FallbackOriginGetResponseStatus = "pending_deletion"
	FallbackOriginGetResponseStatusActive             FallbackOriginGetResponseStatus = "active"
	FallbackOriginGetResponseStatusDeploymentTimedOut FallbackOriginGetResponseStatus = "deployment_timed_out"
	FallbackOriginGetResponseStatusDeletionTimedOut   FallbackOriginGetResponseStatus = "deletion_timed_out"
)

func (r FallbackOriginGetResponseStatus) IsKnown() bool {
	switch r {
	case FallbackOriginGetResponseStatusInitializing, FallbackOriginGetResponseStatusPendingDeployment, FallbackOriginGetResponseStatusPendingDeletion, FallbackOriginGetResponseStatusActive, FallbackOriginGetResponseStatusDeploymentTimedOut, FallbackOriginGetResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type FallbackOriginUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r FallbackOriginUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FallbackOriginUpdateResponseEnvelope struct {
	Errors   []FallbackOriginUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success FallbackOriginUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginUpdateResponse                `json:"result"`
	JSON    fallbackOriginUpdateResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginUpdateResponseEnvelope]
type fallbackOriginUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           FallbackOriginUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             fallbackOriginUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FallbackOriginUpdateResponseEnvelopeErrors]
type fallbackOriginUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    fallbackOriginUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [FallbackOriginUpdateResponseEnvelopeErrorsSource]
type fallbackOriginUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           FallbackOriginUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             fallbackOriginUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FallbackOriginUpdateResponseEnvelopeMessages]
type fallbackOriginUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    fallbackOriginUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [FallbackOriginUpdateResponseEnvelopeMessagesSource]
type fallbackOriginUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FallbackOriginUpdateResponseEnvelopeSuccess bool

const (
	FallbackOriginUpdateResponseEnvelopeSuccessTrue FallbackOriginUpdateResponseEnvelopeSuccess = true
)

func (r FallbackOriginUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FallbackOriginDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FallbackOriginDeleteResponseEnvelope struct {
	Errors   []FallbackOriginDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success FallbackOriginDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginDeleteResponse                `json:"result"`
	JSON    fallbackOriginDeleteResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginDeleteResponseEnvelope]
type fallbackOriginDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           FallbackOriginDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             fallbackOriginDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FallbackOriginDeleteResponseEnvelopeErrors]
type fallbackOriginDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    fallbackOriginDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [FallbackOriginDeleteResponseEnvelopeErrorsSource]
type fallbackOriginDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           FallbackOriginDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             fallbackOriginDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FallbackOriginDeleteResponseEnvelopeMessages]
type fallbackOriginDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    fallbackOriginDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [FallbackOriginDeleteResponseEnvelopeMessagesSource]
type fallbackOriginDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FallbackOriginDeleteResponseEnvelopeSuccess bool

const (
	FallbackOriginDeleteResponseEnvelopeSuccessTrue FallbackOriginDeleteResponseEnvelopeSuccess = true
)

func (r FallbackOriginDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FallbackOriginGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FallbackOriginGetResponseEnvelope struct {
	Errors   []FallbackOriginGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success FallbackOriginGetResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginGetResponse                `json:"result"`
	JSON    fallbackOriginGetResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FallbackOriginGetResponseEnvelope]
type fallbackOriginGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           FallbackOriginGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             fallbackOriginGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// fallbackOriginGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FallbackOriginGetResponseEnvelopeErrors]
type fallbackOriginGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    fallbackOriginGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fallbackOriginGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [FallbackOriginGetResponseEnvelopeErrorsSource]
type fallbackOriginGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           FallbackOriginGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             fallbackOriginGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// fallbackOriginGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FallbackOriginGetResponseEnvelopeMessages]
type fallbackOriginGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    fallbackOriginGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// fallbackOriginGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [FallbackOriginGetResponseEnvelopeMessagesSource]
type fallbackOriginGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FallbackOriginGetResponseEnvelopeSuccess bool

const (
	FallbackOriginGetResponseEnvelopeSuccessTrue FallbackOriginGetResponseEnvelopeSuccess = true
)

func (r FallbackOriginGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
