// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// OwnershipService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOwnershipService] method instead.
type OwnershipService struct {
	Options []option.RequestOption
}

// NewOwnershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOwnershipService(opts ...option.RequestOption) (r *OwnershipService) {
	r = &OwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *OwnershipService) New(ctx context.Context, params OwnershipNewParams, opts ...option.RequestOption) (res *OwnershipNewResponse, err error) {
	var env OwnershipNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates ownership challenge of the destination.
func (r *OwnershipService) Validate(ctx context.Context, params OwnershipValidateParams, opts ...option.RequestOption) (res *OwnershipValidation, err error) {
	var env OwnershipValidateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership/validate", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OwnershipValidation struct {
	Valid bool                    `json:"valid"`
	JSON  ownershipValidationJSON `json:"-"`
}

// ownershipValidationJSON contains the JSON metadata for the struct
// [OwnershipValidation]
type ownershipValidationJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidationJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponse struct {
	Filename string                   `json:"filename"`
	Message  string                   `json:"message"`
	Valid    bool                     `json:"valid"`
	JSON     ownershipNewResponseJSON `json:"-"`
}

// ownershipNewResponseJSON contains the JSON metadata for the struct
// [OwnershipNewResponse]
type ownershipNewResponseJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data. will be
	// pushed. Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r OwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipNewResponseEnvelope struct {
	Errors   []OwnershipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OwnershipNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OwnershipNewResponse                `json:"result,nullable"`
	JSON    ownershipNewResponseEnvelopeJSON    `json:"-"`
}

// ownershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelope]
type ownershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           OwnershipNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             ownershipNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ownershipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelopeErrors]
type ownershipNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    ownershipNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ownershipNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [OwnershipNewResponseEnvelopeErrorsSource]
type ownershipNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           OwnershipNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             ownershipNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ownershipNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OwnershipNewResponseEnvelopeMessages]
type ownershipNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    ownershipNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ownershipNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OwnershipNewResponseEnvelopeMessagesSource]
type ownershipNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OwnershipNewResponseEnvelopeSuccess bool

const (
	OwnershipNewResponseEnvelopeSuccessTrue OwnershipNewResponseEnvelopeSuccess = true
)

func (r OwnershipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data. will be
	// pushed. Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r OwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipValidateResponseEnvelope struct {
	Errors   []OwnershipValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OwnershipValidateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	Result  OwnershipValidation                      `json:"result,nullable"`
	JSON    ownershipValidateResponseEnvelopeJSON    `json:"-"`
}

// ownershipValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipValidateResponseEnvelope]
type ownershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OwnershipValidateResponseEnvelopeErrorsSource `json:"source"`
	JSON             ownershipValidateResponseEnvelopeErrorsJSON   `json:"-"`
}

// ownershipValidateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OwnershipValidateResponseEnvelopeErrors]
type ownershipValidateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    ownershipValidateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ownershipValidateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OwnershipValidateResponseEnvelopeErrorsSource]
type ownershipValidateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           OwnershipValidateResponseEnvelopeMessagesSource `json:"source"`
	JSON             ownershipValidateResponseEnvelopeMessagesJSON   `json:"-"`
}

// ownershipValidateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OwnershipValidateResponseEnvelopeMessages]
type ownershipValidateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    ownershipValidateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ownershipValidateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OwnershipValidateResponseEnvelopeMessagesSource]
type ownershipValidateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OwnershipValidateResponseEnvelopeSuccess bool

const (
	OwnershipValidateResponseEnvelopeSuccessTrue OwnershipValidateResponseEnvelopeSuccess = true
)

func (r OwnershipValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
