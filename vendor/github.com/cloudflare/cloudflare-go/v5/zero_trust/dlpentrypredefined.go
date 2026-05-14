// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPEntryPredefinedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryPredefinedService] method instead.
type DLPEntryPredefinedService struct {
	Options []option.RequestOption
}

// NewDLPEntryPredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPEntryPredefinedService(opts ...option.RequestOption) (r *DLPEntryPredefinedService) {
	r = &DLPEntryPredefinedService{}
	r.Options = opts
	return
}

// Predefined entries can't be created, this will update an existing predefined
// entry This is needed for our generated terraform API
func (r *DLPEntryPredefinedService) New(ctx context.Context, params DLPEntryPredefinedNewParams, opts ...option.RequestOption) (res *DLPEntryPredefinedNewResponse, err error) {
	var env DLPEntryPredefinedNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryPredefinedService) Update(ctx context.Context, entryID string, params DLPEntryPredefinedUpdateParams, opts ...option.RequestOption) (res *DLPEntryPredefinedUpdateResponse, err error) {
	var env DLPEntryPredefinedUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is a no-op as predefined entires can't be deleted but is needed for our
// generated terraform API
func (r *DLPEntryPredefinedService) Delete(ctx context.Context, entryID string, body DLPEntryPredefinedDeleteParams, opts ...option.RequestOption) (res *DLPEntryPredefinedDeleteResponse, err error) {
	var env DLPEntryPredefinedDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEntryPredefinedNewResponse struct {
	ID         string                                  `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedNewResponseConfidence `json:"confidence,required"`
	Enabled    bool                                    `json:"enabled,required"`
	Name       string                                  `json:"name,required"`
	ProfileID  string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryPredefinedNewResponseJSON       `json:"-"`
}

// dlpEntryPredefinedNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedNewResponse]
type dlpEntryPredefinedNewResponseJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                        `json:"available,required"`
	JSON      dlpEntryPredefinedNewResponseConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseConfidenceJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedNewResponseConfidence]
type dlpEntryPredefinedNewResponseConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponse struct {
	ID         string                                     `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedUpdateResponseConfidence `json:"confidence,required"`
	Enabled    bool                                       `json:"enabled,required"`
	Name       string                                     `json:"name,required"`
	ProfileID  string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryPredefinedUpdateResponseJSON       `json:"-"`
}

// dlpEntryPredefinedUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedUpdateResponse]
type dlpEntryPredefinedUpdateResponseJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                           `json:"available,required"`
	JSON      dlpEntryPredefinedUpdateResponseConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedUpdateResponseConfidence]
type dlpEntryPredefinedUpdateResponseConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponse = interface{}

type DLPEntryPredefinedNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
	EntryID   param.Field[string] `json:"entry_id,required" format:"uuid"`
	// This field is not actually used as the owning profile for a predefined entry is
	// already set to a predefined profile
	ProfileID param.Field[string] `json:"profile_id" format:"uuid"`
}

func (r DLPEntryPredefinedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryPredefinedNewResponseEnvelope struct {
	Errors   []DLPEntryPredefinedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedNewResponse                `json:"result"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedNewResponseEnvelope]
type dlpEntryPredefinedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPEntryPredefinedNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedNewResponseEnvelopeErrors]
type dlpEntryPredefinedNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedNewResponseEnvelopeErrorsSource]
type dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPEntryPredefinedNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedNewResponseEnvelopeMessages]
type dlpEntryPredefinedNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedNewResponseEnvelopeMessagesSource]
type dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedNewResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedNewResponseEnvelopeSuccessTrue DLPEntryPredefinedNewResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryPredefinedUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
}

func (r DLPEntryPredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryPredefinedUpdateResponseEnvelope struct {
	Errors   []DLPEntryPredefinedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedUpdateResponse                `json:"result"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedUpdateResponseEnvelope]
type dlpEntryPredefinedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedUpdateResponseEnvelopeErrors]
type dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource]
type dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedUpdateResponseEnvelopeMessages]
type dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource]
type dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedUpdateResponseEnvelopeSuccessTrue DLPEntryPredefinedUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryPredefinedDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryPredefinedDeleteResponseEnvelope struct {
	Errors   []DLPEntryPredefinedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedDeleteResponseEnvelope]
type dlpEntryPredefinedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedDeleteResponseEnvelopeErrors]
type dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource]
type dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedDeleteResponseEnvelopeMessages]
type dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource]
type dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedDeleteResponseEnvelopeSuccessTrue DLPEntryPredefinedDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
