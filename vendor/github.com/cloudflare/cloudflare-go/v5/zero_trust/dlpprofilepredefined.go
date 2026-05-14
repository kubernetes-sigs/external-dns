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

// DLPProfilePredefinedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfilePredefinedService] method instead.
type DLPProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewDLPProfilePredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfilePredefinedService(opts ...option.RequestOption) (r *DLPProfilePredefinedService) {
	r = &DLPProfilePredefinedService{}
	r.Options = opts
	return
}

// Creates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) New(ctx context.Context, params DLPProfilePredefinedNewParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfilePredefinedNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, params DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfilePredefinedUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is a no-op as predefined profiles can't be deleted but is needed for our
// generated terraform API
func (r *DLPProfilePredefinedService) Delete(ctx context.Context, profileID string, body DLPProfilePredefinedDeleteParams, opts ...option.RequestOption) (res *DLPProfilePredefinedDeleteResponse, err error) {
	var env DLPProfilePredefinedDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a predefined DLP profile by id.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfilePredefinedGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPProfilePredefinedDeleteResponse = interface{}

type DLPProfilePredefinedNewParams struct {
	AccountID           param.Field[string] `path:"account_id,required"`
	ProfileID           param.Field[string] `json:"profile_id,required" format:"uuid"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam]                `json:"context_awareness"`
	Entries          param.Field[[]DLPProfilePredefinedNewParamsEntry] `json:"entries"`
	OCREnabled       param.Field[bool]                                 `json:"ocr_enabled"`
}

func (r DLPProfilePredefinedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedNewParamsEntry struct {
	ID      param.Field[string] `json:"id,required" format:"uuid"`
	Enabled param.Field[bool]   `json:"enabled,required"`
}

func (r DLPProfilePredefinedNewParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedNewResponseEnvelope struct {
	Errors   []DLPProfilePredefinedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                        `json:"result"`
	JSON    dlpProfilePredefinedNewResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedNewResponseEnvelope]
type dlpProfilePredefinedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPProfilePredefinedNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedNewResponseEnvelopeErrors]
type dlpProfilePredefinedNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpProfilePredefinedNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedNewResponseEnvelopeErrorsSource]
type dlpProfilePredefinedNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPProfilePredefinedNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedNewResponseEnvelopeMessages]
type dlpProfilePredefinedNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpProfilePredefinedNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedNewResponseEnvelopeMessagesSource]
type dlpProfilePredefinedNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedNewResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedNewResponseEnvelopeSuccessTrue DLPProfilePredefinedNewResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateParams struct {
	AccountID           param.Field[string] `path:"account_id,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam]                   `json:"context_awareness"`
	Entries          param.Field[[]DLPProfilePredefinedUpdateParamsEntry] `json:"entries"`
	OCREnabled       param.Field[bool]                                    `json:"ocr_enabled"`
}

func (r DLPProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateParamsEntry struct {
	ID      param.Field[string] `json:"id,required" format:"uuid"`
	Enabled param.Field[bool]   `json:"enabled,required"`
}

func (r DLPProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateResponseEnvelope struct {
	Errors   []DLPProfilePredefinedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                           `json:"result"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponseEnvelope]
type dlpProfilePredefinedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedUpdateResponseEnvelopeErrors]
type dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource]
type dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEnvelopeMessages]
type dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource]
type dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedUpdateResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedUpdateResponseEnvelopeSuccessTrue DLPProfilePredefinedUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfilePredefinedDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedDeleteResponseEnvelope struct {
	Errors   []DLPProfilePredefinedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfilePredefinedDeleteResponse                `json:"result,nullable"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedDeleteResponseEnvelope]
type dlpProfilePredefinedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedDeleteResponseEnvelopeErrors]
type dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource]
type dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedDeleteResponseEnvelopeMessages]
type dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource]
type dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedDeleteResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedDeleteResponseEnvelopeSuccessTrue DLPProfilePredefinedDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfilePredefinedGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []DLPProfilePredefinedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                        `json:"result"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEnvelope]
type dlpProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPProfilePredefinedGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedGetResponseEnvelopeErrors]
type dlpProfilePredefinedGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEnvelopeErrorsSource]
type dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPProfilePredefinedGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseEnvelopeMessages]
type dlpProfilePredefinedGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEnvelopeMessagesSource]
type dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedGetResponseEnvelopeSuccessTrue DLPProfilePredefinedGetResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
