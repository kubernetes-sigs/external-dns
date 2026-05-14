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

// DLPProfileCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileCustomService] method instead.
type DLPProfileCustomService struct {
	Options []option.RequestOption
}

// NewDLPProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfileCustomService(opts ...option.RequestOption) (r *DLPProfileCustomService) {
	r = &DLPProfileCustomService{}
	r.Options = opts
	return
}

// Creates a DLP custom profile.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, params DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, body DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	var env DLPProfileCustomDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile by id.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pattern struct {
	Regex string `json:"regex,required"`
	// Deprecated: deprecated
	Validation PatternValidation `json:"validation"`
	JSON       patternJSON       `json:"-"`
}

// patternJSON contains the JSON metadata for the struct [Pattern]
type patternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r patternJSON) RawJSON() string {
	return r.raw
}

type PatternValidation string

const (
	PatternValidationLuhn PatternValidation = "luhn"
)

func (r PatternValidation) IsKnown() bool {
	switch r {
	case PatternValidationLuhn:
		return true
	}
	return false
}

type PatternParam struct {
	Regex param.Field[string] `json:"regex,required"`
	// Deprecated: deprecated
	Validation param.Field[PatternValidation] `json:"validation"`
}

func (r PatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomNewParams struct {
	AccountID        param.Field[string]                                `path:"account_id,required"`
	Entries          param.Field[[]DLPProfileCustomNewParamsEntryUnion] `json:"entries,required"`
	Name             param.Field[string]                                `json:"name,required"`
	AIContextEnabled param.Field[bool]                                  `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
	Words   param.Field[interface{}]  `json:"words"`
}

func (r DLPProfileCustomNewParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntry) implementsDLPProfileCustomNewParamsEntryUnion() {}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsEntry].
type DLPProfileCustomNewParamsEntryUnion interface {
	implementsDLPProfileCustomNewParamsEntryUnion()
}

type DLPProfileCustomNewParamsEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntriesDLPNewCustomEntry) implementsDLPProfileCustomNewParamsEntryUnion() {
}

type DLPProfileCustomNewParamsEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntriesDLPNewWordListEntry) implementsDLPProfileCustomNewParamsEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntry struct {
	Enabled   param.Field[bool]                                            `json:"enabled,required"`
	EntryID   param.Field[string]                                          `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntry) implementsDLPProfileCustomNewParamsSharedEntryUnion() {}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsSharedEntriesCustom],
// [zero_trust.DLPProfileCustomNewParamsSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewParamsSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewParamsSharedEntriesExactData],
// [zero_trust.DLPProfileCustomNewParamsSharedEntriesObject],
// [DLPProfileCustomNewParamsSharedEntry].
type DLPProfileCustomNewParamsSharedEntryUnion interface {
	implementsDLPProfileCustomNewParamsSharedEntryUnion()
}

type DLPProfileCustomNewParamsSharedEntriesCustom struct {
	Enabled   param.Field[bool]                                                  `json:"enabled,required"`
	EntryID   param.Field[string]                                                `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesCustomEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntriesCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntriesCustom) implementsDLPProfileCustomNewParamsSharedEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewParamsSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewParamsSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                      `json:"enabled,required"`
	EntryID   param.Field[string]                                                    `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntriesPredefined) implementsDLPProfileCustomNewParamsSharedEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewParamsSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewParamsSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                       `json:"enabled,required"`
	EntryID   param.Field[string]                                                     `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntriesIntegration) implementsDLPProfileCustomNewParamsSharedEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewParamsSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewParamsSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                     `json:"enabled,required"`
	EntryID   param.Field[string]                                                   `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntriesExactData) implementsDLPProfileCustomNewParamsSharedEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewParamsSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                  `json:"enabled,required"`
	EntryID   param.Field[string]                                                `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsSharedEntriesObject) implementsDLPProfileCustomNewParamsSharedEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntriesObjectEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesObjectEntryTypeDocumentFingerprint DLPProfileCustomNewParamsSharedEntriesObjectEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewParamsSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesObjectEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsSharedEntriesEntryTypeCustom              DLPProfileCustomNewParamsSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsSharedEntriesEntryTypePredefined          DLPProfileCustomNewParamsSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsSharedEntriesEntryTypeIntegration         DLPProfileCustomNewParamsSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsSharedEntriesEntryTypeExactData           DLPProfileCustomNewParamsSharedEntriesEntryType = "exact_data"
	DLPProfileCustomNewParamsSharedEntriesEntryTypeDocumentFingerprint DLPProfileCustomNewParamsSharedEntriesEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsSharedEntriesEntryTypeExactData, DLPProfileCustomNewParamsSharedEntriesEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []DLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                    `json:"result"`
	JSON    dlpProfileCustomNewResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelopeErrors]
type dlpProfileCustomNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeErrorsSource]
type dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeMessages]
type dlpProfileCustomNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeMessagesSource]
type dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomNewResponseEnvelopeSuccess bool

const (
	DLPProfileCustomNewResponseEnvelopeSuccessTrue DLPProfileCustomNewResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParams struct {
	AccountID           param.Field[string] `path:"account_id,required"`
	Name                param.Field[string] `json:"name,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// Custom entries from this profile. If this field is omitted, entries owned by
	// this profile will not be changed.
	Entries    param.Field[[]DLPProfileCustomUpdateParamsEntryUnion] `json:"entries"`
	OCREnabled param.Field[bool]                                     `json:"ocr_enabled"`
	// Other entries, e.g. predefined or integration.
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomUpdateParamsEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
	EntryID param.Field[string]       `json:"entry_id" format:"uuid"`
}

func (r DLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {}

// Satisfied by
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID],
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry],
// [DLPProfileCustomUpdateParamsEntry].
type DLPProfileCustomUpdateParamsEntryUnion interface {
	implementsDLPProfileCustomUpdateParamsEntryUnion()
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	EntryID param.Field[string]       `json:"entry_id,required" format:"uuid"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntry struct {
	Enabled   param.Field[bool]                                               `json:"enabled,required"`
	EntryID   param.Field[string]                                             `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntry) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

// Satisfied by [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesExactData],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObject],
// [DLPProfileCustomUpdateParamsSharedEntry].
type DLPProfileCustomUpdateParamsSharedEntryUnion interface {
	implementsDLPProfileCustomUpdateParamsSharedEntryUnion()
}

type DLPProfileCustomUpdateParamsSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefined) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                          `json:"enabled,required"`
	EntryID   param.Field[string]                                                        `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegration) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                        `json:"enabled,required"`
	EntryID   param.Field[string]                                                      `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesExactData) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                     `json:"enabled,required"`
	EntryID   param.Field[string]                                                   `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesObject) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesObjectEntryTypeDocumentFingerprint DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesObjectEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined          DLPProfileCustomUpdateParamsSharedEntriesEntryType = "predefined"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration         DLPProfileCustomUpdateParamsSharedEntriesEntryType = "integration"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData           DLPProfileCustomUpdateParamsSharedEntriesEntryType = "exact_data"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeDocumentFingerprint DLPProfileCustomUpdateParamsSharedEntriesEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseEnvelope struct {
	Errors   []DLPProfileCustomUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                       `json:"result"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEnvelope]
type dlpProfileCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponseEnvelopeErrors]
type dlpProfileCustomUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeErrorsSource]
type dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseEnvelopeMessages]
type dlpProfileCustomUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeMessagesSource]
type dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomUpdateResponseEnvelopeSuccess bool

const (
	DLPProfileCustomUpdateResponseEnvelopeSuccessTrue DLPProfileCustomUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []DLPProfileCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomDeleteResponse                `json:"result,nullable"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomDeleteResponseEnvelopeErrors]
type dlpProfileCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeErrorsSource]
type dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomDeleteResponseEnvelopeMessages]
type dlpProfileCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeMessagesSource]
type dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDeleteResponseEnvelopeSuccessTrue DLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []DLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                    `json:"result"`
	JSON    dlpProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelopeErrors]
type dlpProfileCustomGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeErrorsSource]
type dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseEnvelopeMessages]
type dlpProfileCustomGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeMessagesSource]
type dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	DLPProfileCustomGetResponseEnvelopeSuccessTrue DLPProfileCustomGetResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
