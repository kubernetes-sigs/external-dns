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

// DLPEntryIntegrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryIntegrationService] method instead.
type DLPEntryIntegrationService struct {
	Options []option.RequestOption
}

// NewDLPEntryIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPEntryIntegrationService(opts ...option.RequestOption) (r *DLPEntryIntegrationService) {
	r = &DLPEntryIntegrationService{}
	r.Options = opts
	return
}

// Integration entries can't be created, this will update an existing integration
// entry This is needed for our generated terraform API
func (r *DLPEntryIntegrationService) New(ctx context.Context, params DLPEntryIntegrationNewParams, opts ...option.RequestOption) (res *DLPEntryIntegrationNewResponse, err error) {
	var env DLPEntryIntegrationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryIntegrationService) Update(ctx context.Context, entryID string, params DLPEntryIntegrationUpdateParams, opts ...option.RequestOption) (res *DLPEntryIntegrationUpdateResponse, err error) {
	var env DLPEntryIntegrationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is a no-op as integration entires can't be deleted but is needed for our
// generated terraform API
func (r *DLPEntryIntegrationService) Delete(ctx context.Context, entryID string, body DLPEntryIntegrationDeleteParams, opts ...option.RequestOption) (res *DLPEntryIntegrationDeleteResponse, err error) {
	var env DLPEntryIntegrationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEntryIntegrationNewResponse struct {
	ID        string                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                               `json:"enabled,required"`
	Name      string                             `json:"name,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationNewResponseJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationNewResponse]
type dlpEntryIntegrationNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponse struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationUpdateResponseJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationUpdateResponse]
type dlpEntryIntegrationUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponse = interface{}

type DLPEntryIntegrationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
	EntryID   param.Field[string] `json:"entry_id,required" format:"uuid"`
	// This field is not actually used as the owning profile for a predefined entry is
	// already set to a predefined profile
	ProfileID param.Field[string] `json:"profile_id" format:"uuid"`
}

func (r DLPEntryIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryIntegrationNewResponseEnvelope struct {
	Errors   []DLPEntryIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationNewResponse                `json:"result"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationNewResponseEnvelope]
type dlpEntryIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEntryIntegrationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationNewResponseEnvelopeErrors]
type dlpEntryIntegrationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationNewResponseEnvelopeErrorsSource]
type dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryIntegrationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationNewResponseEnvelopeMessages]
type dlpEntryIntegrationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationNewResponseEnvelopeMessagesSource]
type dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationNewResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationNewResponseEnvelopeSuccessTrue DLPEntryIntegrationNewResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryIntegrationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
}

func (r DLPEntryIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryIntegrationUpdateResponseEnvelope struct {
	Errors   []DLPEntryIntegrationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationUpdateResponse                `json:"result"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationUpdateResponseEnvelope]
type dlpEntryIntegrationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationUpdateResponseEnvelopeErrors]
type dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource]
type dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationUpdateResponseEnvelopeMessages]
type dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource]
type dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationUpdateResponseEnvelopeSuccessTrue DLPEntryIntegrationUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryIntegrationDeleteResponseEnvelope struct {
	Errors   []DLPEntryIntegrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationDeleteResponseEnvelope]
type dlpEntryIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationDeleteResponseEnvelopeErrors]
type dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource]
type dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationDeleteResponseEnvelopeMessages]
type dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource]
type dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationDeleteResponseEnvelopeSuccessTrue DLPEntryIntegrationDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
