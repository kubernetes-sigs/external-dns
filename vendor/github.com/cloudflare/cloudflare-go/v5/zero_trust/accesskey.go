// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// AccessKeyService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessKeyService] method instead.
type AccessKeyService struct {
	Options []option.RequestOption
}

// NewAccessKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessKeyService(opts ...option.RequestOption) (r *AccessKeyService) {
	r = &AccessKeyService{}
	r.Options = opts
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccessKeyService) Update(ctx context.Context, params AccessKeyUpdateParams, opts ...option.RequestOption) (res *AccessKeyUpdateResponse, err error) {
	var env AccessKeyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccessKeyService) Get(ctx context.Context, query AccessKeyGetParams, opts ...option.RequestOption) (res *AccessKeyGetResponse, err error) {
	var env AccessKeyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Perfoms a key rotation for an account.
func (r *AccessKeyService) Rotate(ctx context.Context, body AccessKeyRotateParams, opts ...option.RequestOption) (res *AccessKeyRotateResponse, err error) {
	var env AccessKeyRotateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessKeyUpdateResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                   `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyUpdateResponseJSON `json:"-"`
}

// accessKeyUpdateResponseJSON contains the JSON metadata for the struct
// [AccessKeyUpdateResponse]
type accessKeyUpdateResponseJSON struct {
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessKeyGetResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyGetResponseJSON `json:"-"`
}

// accessKeyGetResponseJSON contains the JSON metadata for the struct
// [AccessKeyGetResponse]
type accessKeyGetResponseJSON struct {
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessKeyRotateResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                   `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyRotateResponseJSON `json:"-"`
}

// accessKeyRotateResponseJSON contains the JSON metadata for the struct
// [AccessKeyRotateResponse]
type accessKeyRotateResponseJSON struct {
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessKeyUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccessKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessKeyUpdateResponseEnvelope struct {
	Errors   []AccessKeyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessKeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessKeyUpdateResponse                `json:"result"`
	JSON    accessKeyUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessKeyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyUpdateResponseEnvelope]
type accessKeyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessKeyUpdateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccessKeyUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessKeyUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessKeyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessKeyUpdateResponseEnvelopeErrors]
type accessKeyUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessKeyUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accessKeyUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessKeyUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessKeyUpdateResponseEnvelopeErrorsSource]
type accessKeyUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessKeyUpdateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessKeyUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessKeyUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessKeyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyUpdateResponseEnvelopeMessages]
type accessKeyUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessKeyUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessKeyUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessKeyUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessKeyUpdateResponseEnvelopeMessagesSource]
type accessKeyUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessKeyUpdateResponseEnvelopeSuccess bool

const (
	AccessKeyUpdateResponseEnvelopeSuccessTrue AccessKeyUpdateResponseEnvelopeSuccess = true
)

func (r AccessKeyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessKeyGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessKeyGetResponseEnvelope struct {
	Errors   []AccessKeyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessKeyGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessKeyGetResponse                `json:"result"`
	JSON    accessKeyGetResponseEnvelopeJSON    `json:"-"`
}

// accessKeyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyGetResponseEnvelope]
type accessKeyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessKeyGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccessKeyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessKeyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessKeyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessKeyGetResponseEnvelopeErrors]
type accessKeyGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessKeyGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accessKeyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessKeyGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AccessKeyGetResponseEnvelopeErrorsSource]
type accessKeyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessKeyGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccessKeyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessKeyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessKeyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyGetResponseEnvelopeMessages]
type accessKeyGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessKeyGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accessKeyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessKeyGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessKeyGetResponseEnvelopeMessagesSource]
type accessKeyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessKeyGetResponseEnvelopeSuccess bool

const (
	AccessKeyGetResponseEnvelopeSuccessTrue AccessKeyGetResponseEnvelopeSuccess = true
)

func (r AccessKeyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessKeyRotateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessKeyRotateResponseEnvelope struct {
	Errors   []AccessKeyRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyRotateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessKeyRotateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessKeyRotateResponse                `json:"result"`
	JSON    accessKeyRotateResponseEnvelopeJSON    `json:"-"`
}

// accessKeyRotateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyRotateResponseEnvelope]
type accessKeyRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessKeyRotateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccessKeyRotateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessKeyRotateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessKeyRotateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessKeyRotateResponseEnvelopeErrors]
type accessKeyRotateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessKeyRotateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accessKeyRotateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessKeyRotateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessKeyRotateResponseEnvelopeErrorsSource]
type accessKeyRotateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessKeyRotateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessKeyRotateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessKeyRotateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessKeyRotateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyRotateResponseEnvelopeMessages]
type accessKeyRotateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessKeyRotateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessKeyRotateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessKeyRotateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessKeyRotateResponseEnvelopeMessagesSource]
type accessKeyRotateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyRotateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessKeyRotateResponseEnvelopeSuccess bool

const (
	AccessKeyRotateResponseEnvelopeSuccessTrue AccessKeyRotateResponseEnvelopeSuccess = true
)

func (r AccessKeyRotateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyRotateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
