// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileService] method instead.
type DLPProfileService struct {
	Options    []option.RequestOption
	Custom     *DLPProfileCustomService
	Predefined *DLPProfilePredefinedService
}

// NewDLPProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPProfileService(opts ...option.RequestOption) (r *DLPProfileService) {
	r = &DLPProfileService{}
	r.Options = opts
	r.Custom = NewDLPProfileCustomService(opts...)
	r.Predefined = NewDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) List(ctx context.Context, params DLPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Profile], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) ListAutoPaging(ctx context.Context, params DLPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Profile] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a DLP profile by ID.
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
//
// Deprecated: deprecated
type ContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip SkipConfiguration    `json:"skip,required"`
	JSON contextAwarenessJSON `json:"-"`
}

// contextAwarenessJSON contains the JSON metadata for the struct
// [ContextAwareness]
type contextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
//
// Deprecated: deprecated
type ContextAwarenessParam struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[SkipConfigurationParam] `json:"skip,required"`
}

func (r ContextAwarenessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Profile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of [[]ProfileCustomProfileEntry],
	// [[]ProfilePredefinedProfileEntry], [[]ProfileIntegrationProfileEntry].
	Entries interface{} `json:"entries,required"`
	// The name of the profile.
	Name             string      `json:"name,required"`
	Type             ProfileType `json:"type,required"`
	AIContextEnabled bool        `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                      `json:"allowed_match_count"`
	ConfidenceThreshold ProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	OCREnabled  bool   `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// When the profile was lasted updated.
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	JSON      profileJSON `json:"-"`
	union     ProfileUnion
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	ID                  apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Description         apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	*r = Profile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ProfileCustomProfile],
// [ProfilePredefinedProfile], [ProfileIntegrationProfile].
func (r Profile) AsUnion() ProfileUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfile], [ProfilePredefinedProfile] or
// [ProfileIntegrationProfile].
type ProfileUnion interface {
	implementsProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfile{}),
		},
	)
}

type ProfileCustomProfile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time                   `json:"created_at,required" format:"date-time"`
	Entries   []ProfileCustomProfileEntry `json:"entries,required"`
	// The name of the profile.
	Name       string                   `json:"name,required"`
	OCREnabled bool                     `json:"ocr_enabled,required"`
	Type       ProfileCustomProfileType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                               `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                    `json:"ai_context_enabled"`
	ConfidenceThreshold ProfileCustomProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string                   `json:"description,nullable"`
	JSON        profileCustomProfileJSON `json:"-"`
}

// profileCustomProfileJSON contains the JSON metadata for the struct
// [ProfileCustomProfile]
type profileCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfile) implementsProfile() {}

type ProfileCustomProfileEntry struct {
	ID      string                          `json:"id,required" format:"uuid"`
	Enabled bool                            `json:"enabled,required"`
	Name    string                          `json:"name,required"`
	Type    ProfileCustomProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileCustomProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                   `json:"word_list"`
	JSON     profileCustomProfileEntryJSON `json:"-"`
	union    ProfileCustomProfileEntriesUnion
}

// profileCustomProfileEntryJSON contains the JSON metadata for the struct
// [ProfileCustomProfileEntry]
type profileCustomProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileCustomProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileCustomProfileEntriesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileCustomProfileEntriesCustomEntry],
// [ProfileCustomProfileEntriesPredefinedEntry],
// [ProfileCustomProfileEntriesIntegrationEntry],
// [ProfileCustomProfileEntriesExactDataEntry],
// [ProfileCustomProfileEntriesDocumentFingerprintEntry],
// [ProfileCustomProfileEntriesWordListEntry].
func (r ProfileCustomProfileEntry) AsUnion() ProfileCustomProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfileEntriesCustomEntry],
// [ProfileCustomProfileEntriesPredefinedEntry],
// [ProfileCustomProfileEntriesIntegrationEntry],
// [ProfileCustomProfileEntriesExactDataEntry],
// [ProfileCustomProfileEntriesDocumentFingerprintEntry] or
// [ProfileCustomProfileEntriesWordListEntry].
type ProfileCustomProfileEntriesUnion interface {
	implementsProfileCustomProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileCustomProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesWordListEntry{}),
		},
	)
}

type ProfileCustomProfileEntriesCustomEntry struct {
	ID        string                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                       `json:"enabled,required"`
	Name      string                                     `json:"name,required"`
	Pattern   Pattern                                    `json:"pattern,required"`
	Type      ProfileCustomProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesCustomEntryJSON `json:"-"`
}

// profileCustomProfileEntriesCustomEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesCustomEntry]
type profileCustomProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesCustomEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesCustomEntryType string

const (
	ProfileCustomProfileEntriesCustomEntryTypeCustom ProfileCustomProfileEntriesCustomEntryType = "custom"
)

func (r ProfileCustomProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesPredefinedEntry struct {
	ID         string                                               `json:"id,required" format:"uuid"`
	Confidence ProfileCustomProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                 `json:"enabled,required"`
	Name       string                                               `json:"name,required"`
	Type       ProfileCustomProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON       profileCustomProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// profileCustomProfileEntriesPredefinedEntryJSON contains the JSON metadata for
// the struct [ProfileCustomProfileEntriesPredefinedEntry]
type profileCustomProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesPredefinedEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                     `json:"available,required"`
	JSON      profileCustomProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileCustomProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [ProfileCustomProfileEntriesPredefinedEntryConfidence]
type profileCustomProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomProfileEntriesPredefinedEntryType string

const (
	ProfileCustomProfileEntriesPredefinedEntryTypePredefined ProfileCustomProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfileCustomProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesIntegrationEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      ProfileCustomProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profileCustomProfileEntriesIntegrationEntryJSON contains the JSON metadata for
// the struct [ProfileCustomProfileEntriesIntegrationEntry]
type profileCustomProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesIntegrationEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesIntegrationEntryType string

const (
	ProfileCustomProfileEntriesIntegrationEntryTypeIntegration ProfileCustomProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfileCustomProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                          `json:"case_sensitive,required"`
	CreatedAt     time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled       bool                                          `json:"enabled,required"`
	Name          string                                        `json:"name,required"`
	Secret        bool                                          `json:"secret,required"`
	Type          ProfileCustomProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                     `json:"updated_at,required" format:"date-time"`
	JSON          profileCustomProfileEntriesExactDataEntryJSON `json:"-"`
}

// profileCustomProfileEntriesExactDataEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesExactDataEntry]
type profileCustomProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesExactDataEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesExactDataEntryType string

const (
	ProfileCustomProfileEntriesExactDataEntryTypeExactData ProfileCustomProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfileCustomProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      ProfileCustomProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	JSON      profileCustomProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileCustomProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [ProfileCustomProfileEntriesDocumentFingerprintEntry]
type profileCustomProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesDocumentFingerprintEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesDocumentFingerprintEntryType string

const (
	ProfileCustomProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileCustomProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileCustomProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesWordListEntry struct {
	ID        string                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                         `json:"enabled,required"`
	Name      string                                       `json:"name,required"`
	Type      ProfileCustomProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                    `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                  `json:"word_list,required"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesWordListEntryJSON `json:"-"`
}

// profileCustomProfileEntriesWordListEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesWordListEntry]
type profileCustomProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesWordListEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesWordListEntryType string

const (
	ProfileCustomProfileEntriesWordListEntryTypeWordList ProfileCustomProfileEntriesWordListEntryType = "word_list"
)

func (r ProfileCustomProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesType string

const (
	ProfileCustomProfileEntriesTypeCustom              ProfileCustomProfileEntriesType = "custom"
	ProfileCustomProfileEntriesTypePredefined          ProfileCustomProfileEntriesType = "predefined"
	ProfileCustomProfileEntriesTypeIntegration         ProfileCustomProfileEntriesType = "integration"
	ProfileCustomProfileEntriesTypeExactData           ProfileCustomProfileEntriesType = "exact_data"
	ProfileCustomProfileEntriesTypeDocumentFingerprint ProfileCustomProfileEntriesType = "document_fingerprint"
	ProfileCustomProfileEntriesTypeWordList            ProfileCustomProfileEntriesType = "word_list"
)

func (r ProfileCustomProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesTypeCustom, ProfileCustomProfileEntriesTypePredefined, ProfileCustomProfileEntriesTypeIntegration, ProfileCustomProfileEntriesTypeExactData, ProfileCustomProfileEntriesTypeDocumentFingerprint, ProfileCustomProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileCustomProfileType string

const (
	ProfileCustomProfileTypeCustom ProfileCustomProfileType = "custom"
)

func (r ProfileCustomProfileType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileTypeCustom:
		return true
	}
	return false
}

type ProfileCustomProfileConfidenceThreshold string

const (
	ProfileCustomProfileConfidenceThresholdLow      ProfileCustomProfileConfidenceThreshold = "low"
	ProfileCustomProfileConfidenceThresholdMedium   ProfileCustomProfileConfidenceThreshold = "medium"
	ProfileCustomProfileConfidenceThresholdHigh     ProfileCustomProfileConfidenceThreshold = "high"
	ProfileCustomProfileConfidenceThresholdVeryHigh ProfileCustomProfileConfidenceThreshold = "very_high"
)

func (r ProfileCustomProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfileCustomProfileConfidenceThresholdLow, ProfileCustomProfileConfidenceThresholdMedium, ProfileCustomProfileConfidenceThresholdHigh, ProfileCustomProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfilePredefinedProfile struct {
	// The id of the predefined profile (uuid).
	ID                string                          `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                           `json:"allowed_match_count,required"`
	Entries           []ProfilePredefinedProfileEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                      `json:"name,required"`
	Type                ProfilePredefinedProfileType                `json:"type,required"`
	AIContextEnabled    bool                                        `json:"ai_context_enabled"`
	ConfidenceThreshold ProfilePredefinedProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                         `json:"open_access"`
	JSON       profilePredefinedProfileJSON `json:"-"`
}

// profilePredefinedProfileJSON contains the JSON metadata for the struct
// [ProfilePredefinedProfile]
type profilePredefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfilePredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfile) implementsProfile() {}

type ProfilePredefinedProfileEntry struct {
	ID      string                              `json:"id,required" format:"uuid"`
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	Type    ProfilePredefinedProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfilePredefinedProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                       `json:"word_list"`
	JSON     profilePredefinedProfileEntryJSON `json:"-"`
	union    ProfilePredefinedProfileEntriesUnion
}

// profilePredefinedProfileEntryJSON contains the JSON metadata for the struct
// [ProfilePredefinedProfileEntry]
type profilePredefinedProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profilePredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfilePredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfilePredefinedProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfilePredefinedProfileEntriesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfilePredefinedProfileEntriesCustomEntry],
// [ProfilePredefinedProfileEntriesPredefinedEntry],
// [ProfilePredefinedProfileEntriesIntegrationEntry],
// [ProfilePredefinedProfileEntriesExactDataEntry],
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry],
// [ProfilePredefinedProfileEntriesWordListEntry].
func (r ProfilePredefinedProfileEntry) AsUnion() ProfilePredefinedProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfilePredefinedProfileEntriesCustomEntry],
// [ProfilePredefinedProfileEntriesPredefinedEntry],
// [ProfilePredefinedProfileEntriesIntegrationEntry],
// [ProfilePredefinedProfileEntriesExactDataEntry],
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry] or
// [ProfilePredefinedProfileEntriesWordListEntry].
type ProfilePredefinedProfileEntriesUnion interface {
	implementsProfilePredefinedProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfilePredefinedProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesWordListEntry{}),
		},
	)
}

type ProfilePredefinedProfileEntriesCustomEntry struct {
	ID        string                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                           `json:"enabled,required"`
	Name      string                                         `json:"name,required"`
	Pattern   Pattern                                        `json:"pattern,required"`
	Type      ProfilePredefinedProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesCustomEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesCustomEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesCustomEntry]
type profilePredefinedProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesCustomEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesCustomEntryType string

const (
	ProfilePredefinedProfileEntriesCustomEntryTypeCustom ProfilePredefinedProfileEntriesCustomEntryType = "custom"
)

func (r ProfilePredefinedProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesPredefinedEntry struct {
	ID         string                                                   `json:"id,required" format:"uuid"`
	Confidence ProfilePredefinedProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                     `json:"enabled,required"`
	Name       string                                                   `json:"name,required"`
	Type       ProfilePredefinedProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON       profilePredefinedProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// profilePredefinedProfileEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [ProfilePredefinedProfileEntriesPredefinedEntry]
type profilePredefinedProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesPredefinedEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                         `json:"available,required"`
	JSON      profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [ProfilePredefinedProfileEntriesPredefinedEntryConfidence]
type profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfilePredefinedProfileEntriesPredefinedEntryType string

const (
	ProfilePredefinedProfileEntriesPredefinedEntryTypePredefined ProfilePredefinedProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfilePredefinedProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesIntegrationEntry struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesIntegrationEntryJSON contains the JSON metadata
// for the struct [ProfilePredefinedProfileEntriesIntegrationEntry]
type profilePredefinedProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesIntegrationEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesIntegrationEntryType string

const (
	ProfilePredefinedProfileEntriesIntegrationEntryTypeIntegration ProfilePredefinedProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfilePredefinedProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                              `json:"case_sensitive,required"`
	CreatedAt     time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled       bool                                              `json:"enabled,required"`
	Name          string                                            `json:"name,required"`
	Secret        bool                                              `json:"secret,required"`
	Type          ProfilePredefinedProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                         `json:"updated_at,required" format:"date-time"`
	JSON          profilePredefinedProfileEntriesExactDataEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesExactDataEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesExactDataEntry]
type profilePredefinedProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesExactDataEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesExactDataEntryType string

const (
	ProfilePredefinedProfileEntriesExactDataEntryTypeExactData ProfilePredefinedProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfilePredefinedProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	JSON      profilePredefinedProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry]
type profilePredefinedProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesDocumentFingerprintEntry) implementsProfilePredefinedProfileEntry() {
}

type ProfilePredefinedProfileEntriesDocumentFingerprintEntryType string

const (
	ProfilePredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfilePredefinedProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfilePredefinedProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesWordListEntry struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                      `json:"word_list,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesWordListEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesWordListEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesWordListEntry]
type profilePredefinedProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesWordListEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesWordListEntryType string

const (
	ProfilePredefinedProfileEntriesWordListEntryTypeWordList ProfilePredefinedProfileEntriesWordListEntryType = "word_list"
)

func (r ProfilePredefinedProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesType string

const (
	ProfilePredefinedProfileEntriesTypeCustom              ProfilePredefinedProfileEntriesType = "custom"
	ProfilePredefinedProfileEntriesTypePredefined          ProfilePredefinedProfileEntriesType = "predefined"
	ProfilePredefinedProfileEntriesTypeIntegration         ProfilePredefinedProfileEntriesType = "integration"
	ProfilePredefinedProfileEntriesTypeExactData           ProfilePredefinedProfileEntriesType = "exact_data"
	ProfilePredefinedProfileEntriesTypeDocumentFingerprint ProfilePredefinedProfileEntriesType = "document_fingerprint"
	ProfilePredefinedProfileEntriesTypeWordList            ProfilePredefinedProfileEntriesType = "word_list"
)

func (r ProfilePredefinedProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesTypeCustom, ProfilePredefinedProfileEntriesTypePredefined, ProfilePredefinedProfileEntriesTypeIntegration, ProfilePredefinedProfileEntriesTypeExactData, ProfilePredefinedProfileEntriesTypeDocumentFingerprint, ProfilePredefinedProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedProfileType string

const (
	ProfilePredefinedProfileTypePredefined ProfilePredefinedProfileType = "predefined"
)

func (r ProfilePredefinedProfileType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedProfileConfidenceThreshold string

const (
	ProfilePredefinedProfileConfidenceThresholdLow      ProfilePredefinedProfileConfidenceThreshold = "low"
	ProfilePredefinedProfileConfidenceThresholdMedium   ProfilePredefinedProfileConfidenceThreshold = "medium"
	ProfilePredefinedProfileConfidenceThresholdHigh     ProfilePredefinedProfileConfidenceThreshold = "high"
	ProfilePredefinedProfileConfidenceThresholdVeryHigh ProfilePredefinedProfileConfidenceThreshold = "very_high"
)

func (r ProfilePredefinedProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileConfidenceThresholdLow, ProfilePredefinedProfileConfidenceThresholdMedium, ProfilePredefinedProfileConfidenceThresholdHigh, ProfilePredefinedProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfileIntegrationProfile struct {
	ID        string                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                        `json:"created_at,required" format:"date-time"`
	Entries   []ProfileIntegrationProfileEntry `json:"entries,required"`
	Name      string                           `json:"name,required"`
	Type      ProfileIntegrationProfileType    `json:"type,required"`
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                        `json:"description,nullable"`
	JSON        profileIntegrationProfileJSON `json:"-"`
}

// profileIntegrationProfileJSON contains the JSON metadata for the struct
// [ProfileIntegrationProfile]
type profileIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfile) implementsProfile() {}

type ProfileIntegrationProfileEntry struct {
	ID      string                               `json:"id,required" format:"uuid"`
	Enabled bool                                 `json:"enabled,required"`
	Name    string                               `json:"name,required"`
	Type    ProfileIntegrationProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileIntegrationProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                        `json:"word_list"`
	JSON     profileIntegrationProfileEntryJSON `json:"-"`
	union    ProfileIntegrationProfileEntriesUnion
}

// profileIntegrationProfileEntryJSON contains the JSON metadata for the struct
// [ProfileIntegrationProfileEntry]
type profileIntegrationProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileIntegrationProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileIntegrationProfileEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileIntegrationProfileEntriesCustomEntry],
// [ProfileIntegrationProfileEntriesPredefinedEntry],
// [ProfileIntegrationProfileEntriesIntegrationEntry],
// [ProfileIntegrationProfileEntriesExactDataEntry],
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry],
// [ProfileIntegrationProfileEntriesWordListEntry].
func (r ProfileIntegrationProfileEntry) AsUnion() ProfileIntegrationProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileIntegrationProfileEntriesCustomEntry],
// [ProfileIntegrationProfileEntriesPredefinedEntry],
// [ProfileIntegrationProfileEntriesIntegrationEntry],
// [ProfileIntegrationProfileEntriesExactDataEntry],
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry] or
// [ProfileIntegrationProfileEntriesWordListEntry].
type ProfileIntegrationProfileEntriesUnion interface {
	implementsProfileIntegrationProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileIntegrationProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesWordListEntry{}),
		},
	)
}

type ProfileIntegrationProfileEntriesCustomEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Pattern   Pattern                                         `json:"pattern,required"`
	Type      ProfileIntegrationProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesCustomEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesCustomEntryJSON contains the JSON metadata for
// the struct [ProfileIntegrationProfileEntriesCustomEntry]
type profileIntegrationProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesCustomEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesCustomEntryType string

const (
	ProfileIntegrationProfileEntriesCustomEntryTypeCustom ProfileIntegrationProfileEntriesCustomEntryType = "custom"
)

func (r ProfileIntegrationProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesPredefinedEntry struct {
	ID         string                                                    `json:"id,required" format:"uuid"`
	Confidence ProfileIntegrationProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                      `json:"enabled,required"`
	Name       string                                                    `json:"name,required"`
	Type       ProfileIntegrationProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON       profileIntegrationProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// profileIntegrationProfileEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesPredefinedEntry]
type profileIntegrationProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesPredefinedEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                          `json:"available,required"`
	JSON      profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [ProfileIntegrationProfileEntriesPredefinedEntryConfidence]
type profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationProfileEntriesPredefinedEntryType string

const (
	ProfileIntegrationProfileEntriesPredefinedEntryTypePredefined ProfileIntegrationProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfileIntegrationProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesIntegrationEntry struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesIntegrationEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesIntegrationEntry]
type profileIntegrationProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesIntegrationEntry) implementsProfileIntegrationProfileEntry() {
}

type ProfileIntegrationProfileEntriesIntegrationEntryType string

const (
	ProfileIntegrationProfileEntriesIntegrationEntryTypeIntegration ProfileIntegrationProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfileIntegrationProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                               `json:"case_sensitive,required"`
	CreatedAt     time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled       bool                                               `json:"enabled,required"`
	Name          string                                             `json:"name,required"`
	Secret        bool                                               `json:"secret,required"`
	Type          ProfileIntegrationProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                          `json:"updated_at,required" format:"date-time"`
	JSON          profileIntegrationProfileEntriesExactDataEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesExactDataEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesExactDataEntry]
type profileIntegrationProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesExactDataEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesExactDataEntryType string

const (
	ProfileIntegrationProfileEntriesExactDataEntryTypeExactData ProfileIntegrationProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfileIntegrationProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	JSON      profileIntegrationProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry]
type profileIntegrationProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesDocumentFingerprintEntry) implementsProfileIntegrationProfileEntry() {
}

type ProfileIntegrationProfileEntriesDocumentFingerprintEntryType string

const (
	ProfileIntegrationProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileIntegrationProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileIntegrationProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesWordListEntry struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                       `json:"word_list,required"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesWordListEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesWordListEntryJSON contains the JSON metadata for
// the struct [ProfileIntegrationProfileEntriesWordListEntry]
type profileIntegrationProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesWordListEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesWordListEntryType string

const (
	ProfileIntegrationProfileEntriesWordListEntryTypeWordList ProfileIntegrationProfileEntriesWordListEntryType = "word_list"
)

func (r ProfileIntegrationProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesType string

const (
	ProfileIntegrationProfileEntriesTypeCustom              ProfileIntegrationProfileEntriesType = "custom"
	ProfileIntegrationProfileEntriesTypePredefined          ProfileIntegrationProfileEntriesType = "predefined"
	ProfileIntegrationProfileEntriesTypeIntegration         ProfileIntegrationProfileEntriesType = "integration"
	ProfileIntegrationProfileEntriesTypeExactData           ProfileIntegrationProfileEntriesType = "exact_data"
	ProfileIntegrationProfileEntriesTypeDocumentFingerprint ProfileIntegrationProfileEntriesType = "document_fingerprint"
	ProfileIntegrationProfileEntriesTypeWordList            ProfileIntegrationProfileEntriesType = "word_list"
)

func (r ProfileIntegrationProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesTypeCustom, ProfileIntegrationProfileEntriesTypePredefined, ProfileIntegrationProfileEntriesTypeIntegration, ProfileIntegrationProfileEntriesTypeExactData, ProfileIntegrationProfileEntriesTypeDocumentFingerprint, ProfileIntegrationProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileType string

const (
	ProfileIntegrationProfileTypeIntegration ProfileIntegrationProfileType = "integration"
)

func (r ProfileIntegrationProfileType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type ProfileType string

const (
	ProfileTypeCustom      ProfileType = "custom"
	ProfileTypePredefined  ProfileType = "predefined"
	ProfileTypeIntegration ProfileType = "integration"
)

func (r ProfileType) IsKnown() bool {
	switch r {
	case ProfileTypeCustom, ProfileTypePredefined, ProfileTypeIntegration:
		return true
	}
	return false
}

type ProfileConfidenceThreshold string

const (
	ProfileConfidenceThresholdLow      ProfileConfidenceThreshold = "low"
	ProfileConfidenceThresholdMedium   ProfileConfidenceThreshold = "medium"
	ProfileConfidenceThresholdHigh     ProfileConfidenceThreshold = "high"
	ProfileConfidenceThresholdVeryHigh ProfileConfidenceThreshold = "very_high"
)

func (r ProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfileConfidenceThresholdLow, ProfileConfidenceThresholdMedium, ProfileConfidenceThresholdHigh, ProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

// Content types to exclude from context analysis and return all matches.
type SkipConfiguration struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                  `json:"files,required"`
	JSON  skipConfigurationJSON `json:"-"`
}

// skipConfigurationJSON contains the JSON metadata for the struct
// [SkipConfiguration]
type skipConfigurationJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkipConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipConfigurationJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type SkipConfigurationParam struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r SkipConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Return all profiles, including those that current account does not have access
	// to.
	All param.Field[bool] `query:"all"`
}

// URLQuery serializes [DLPProfileListParams]'s query parameters as `url.Values`.
func (r DLPProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DLPProfileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors   []DLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                              `json:"result"`
	JSON    dlpProfileGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPProfileGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrors]
type dlpProfileGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpProfileGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrorsSource]
type dlpProfileGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPProfileGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeMessages]
type dlpProfileGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpProfileGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseEnvelopeMessagesSource]
type dlpProfileGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileGetResponseEnvelopeSuccess bool

const (
	DLPProfileGetResponseEnvelopeSuccessTrue DLPProfileGetResponseEnvelopeSuccess = true
)

func (r DLPProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
