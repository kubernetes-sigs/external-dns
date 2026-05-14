// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// DLPEntryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryService] method instead.
type DLPEntryService struct {
	Options     []option.RequestOption
	Custom      *DLPEntryCustomService
	Predefined  *DLPEntryPredefinedService
	Integration *DLPEntryIntegrationService
}

// NewDLPEntryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPEntryService(opts ...option.RequestOption) (r *DLPEntryService) {
	r = &DLPEntryService{}
	r.Options = opts
	r.Custom = NewDLPEntryCustomService(opts...)
	r.Predefined = NewDLPEntryPredefinedService(opts...)
	r.Integration = NewDLPEntryIntegrationService(opts...)
	return
}

// Creates a DLP custom entry.
func (r *DLPEntryService) New(ctx context.Context, params DLPEntryNewParams, opts ...option.RequestOption) (res *DLPEntryNewResponse, err error) {
	var env DLPEntryNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryService) Update(ctx context.Context, entryID string, params DLPEntryUpdateParams, opts ...option.RequestOption) (res *DLPEntryUpdateResponse, err error) {
	var env DLPEntryUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all DLP entries in an account.
func (r *DLPEntryService) List(ctx context.Context, query DLPEntryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEntryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all DLP entries in an account.
func (r *DLPEntryService) ListAutoPaging(ctx context.Context, query DLPEntryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEntryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a DLP custom entry.
func (r *DLPEntryService) Delete(ctx context.Context, entryID string, body DLPEntryDeleteParams, opts ...option.RequestOption) (res *DLPEntryDeleteResponse, err error) {
	var env DLPEntryDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP entry by ID.
func (r *DLPEntryService) Get(ctx context.Context, entryID string, query DLPEntryGetParams, opts ...option.RequestOption) (res *DLPEntryGetResponse, err error) {
	var env DLPEntryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", query.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEntryNewResponse struct {
	ID        string                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	Enabled   bool                    `json:"enabled,required"`
	Name      string                  `json:"name,required"`
	Pattern   Pattern                 `json:"pattern,required"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
	ProfileID string                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryNewResponseJSON `json:"-"`
}

// dlpEntryNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryNewResponse]
type dlpEntryNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponse struct {
	ID      string                     `json:"id,required" format:"uuid"`
	Enabled bool                       `json:"enabled,required"`
	Name    string                     `json:"name,required"`
	Type    DLPEntryUpdateResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryUpdateResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                `json:"word_list"`
	JSON     dlpEntryUpdateResponseJSON `json:"-"`
	union    DLPEntryUpdateResponseUnion
}

// dlpEntryUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponse]
type dlpEntryUpdateResponseJSON struct {
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

func (r dlpEntryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryUpdateResponseCustomEntry],
// [DLPEntryUpdateResponsePredefinedEntry],
// [DLPEntryUpdateResponseIntegrationEntry],
// [DLPEntryUpdateResponseExactDataEntry],
// [DLPEntryUpdateResponseDocumentFingerprintEntry],
// [DLPEntryUpdateResponseWordListEntry].
func (r DLPEntryUpdateResponse) AsUnion() DLPEntryUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryUpdateResponseCustomEntry],
// [DLPEntryUpdateResponsePredefinedEntry],
// [DLPEntryUpdateResponseIntegrationEntry],
// [DLPEntryUpdateResponseExactDataEntry],
// [DLPEntryUpdateResponseDocumentFingerprintEntry] or
// [DLPEntryUpdateResponseWordListEntry].
type DLPEntryUpdateResponseUnion interface {
	implementsDLPEntryUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseWordListEntry{}),
		},
	)
}

type DLPEntryUpdateResponseCustomEntry struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	Pattern   Pattern                               `json:"pattern,required"`
	Type      DLPEntryUpdateResponseCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseCustomEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseCustomEntry]
type dlpEntryUpdateResponseCustomEntryJSON struct {
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

func (r *DLPEntryUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseCustomEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseCustomEntryType string

const (
	DLPEntryUpdateResponseCustomEntryTypeCustom DLPEntryUpdateResponseCustomEntryType = "custom"
)

func (r DLPEntryUpdateResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedEntry struct {
	ID         string                                          `json:"id,required" format:"uuid"`
	Confidence DLPEntryUpdateResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                            `json:"enabled,required"`
	Name       string                                          `json:"name,required"`
	Type       DLPEntryUpdateResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryUpdateResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryUpdateResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponsePredefinedEntry]
type dlpEntryUpdateResponsePredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponsePredefinedEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                `json:"available,required"`
	JSON      dlpEntryUpdateResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryUpdateResponsePredefinedEntryConfidenceJSON contains the JSON metadata
// for the struct [DLPEntryUpdateResponsePredefinedEntryConfidence]
type dlpEntryUpdateResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponsePredefinedEntryType string

const (
	DLPEntryUpdateResponsePredefinedEntryTypePredefined DLPEntryUpdateResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryUpdateResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryUpdateResponseIntegrationEntry struct {
	ID        string                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                       `json:"enabled,required"`
	Name      string                                     `json:"name,required"`
	Type      DLPEntryUpdateResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseIntegrationEntry]
type dlpEntryUpdateResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseIntegrationEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseIntegrationEntryType string

const (
	DLPEntryUpdateResponseIntegrationEntryTypeIntegration DLPEntryUpdateResponseIntegrationEntryType = "integration"
)

func (r DLPEntryUpdateResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                     `json:"enabled,required"`
	Name          string                                   `json:"name,required"`
	Secret        bool                                     `json:"secret,required"`
	Type          DLPEntryUpdateResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryUpdateResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseExactDataEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseExactDataEntry]
type dlpEntryUpdateResponseExactDataEntryJSON struct {
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

func (r *DLPEntryUpdateResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseExactDataEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseExactDataEntryType string

const (
	DLPEntryUpdateResponseExactDataEntryTypeExactData DLPEntryUpdateResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryUpdateResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryUpdateResponseDocumentFingerprintEntry struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPEntryUpdateResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryUpdateResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseDocumentFingerprintEntryJSON contains the JSON metadata
// for the struct [DLPEntryUpdateResponseDocumentFingerprintEntry]
type dlpEntryUpdateResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseDocumentFingerprintEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseDocumentFingerprintEntryType string

const (
	DLPEntryUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryUpdateResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryUpdateResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryUpdateResponseWordListEntry struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                    `json:"enabled,required"`
	Name      string                                  `json:"name,required"`
	Type      DLPEntryUpdateResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                             `json:"word_list,required"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseWordListEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseWordListEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseWordListEntry]
type dlpEntryUpdateResponseWordListEntryJSON struct {
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

func (r *DLPEntryUpdateResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseWordListEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseWordListEntryType string

const (
	DLPEntryUpdateResponseWordListEntryTypeWordList DLPEntryUpdateResponseWordListEntryType = "word_list"
)

func (r DLPEntryUpdateResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryUpdateResponseType string

const (
	DLPEntryUpdateResponseTypeCustom              DLPEntryUpdateResponseType = "custom"
	DLPEntryUpdateResponseTypePredefined          DLPEntryUpdateResponseType = "predefined"
	DLPEntryUpdateResponseTypeIntegration         DLPEntryUpdateResponseType = "integration"
	DLPEntryUpdateResponseTypeExactData           DLPEntryUpdateResponseType = "exact_data"
	DLPEntryUpdateResponseTypeDocumentFingerprint DLPEntryUpdateResponseType = "document_fingerprint"
	DLPEntryUpdateResponseTypeWordList            DLPEntryUpdateResponseType = "word_list"
)

func (r DLPEntryUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseTypeCustom, DLPEntryUpdateResponseTypePredefined, DLPEntryUpdateResponseTypeIntegration, DLPEntryUpdateResponseTypeExactData, DLPEntryUpdateResponseTypeDocumentFingerprint, DLPEntryUpdateResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryListResponse struct {
	ID      string                   `json:"id,required" format:"uuid"`
	Enabled bool                     `json:"enabled,required"`
	Name    string                   `json:"name,required"`
	Type    DLPEntryListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryListResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}              `json:"word_list"`
	JSON     dlpEntryListResponseJSON `json:"-"`
	union    DLPEntryListResponseUnion
}

// dlpEntryListResponseJSON contains the JSON metadata for the struct
// [DLPEntryListResponse]
type dlpEntryListResponseJSON struct {
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

func (r dlpEntryListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryListResponseCustomEntry],
// [DLPEntryListResponsePredefinedEntry], [DLPEntryListResponseIntegrationEntry],
// [DLPEntryListResponseExactDataEntry],
// [DLPEntryListResponseDocumentFingerprintEntry],
// [DLPEntryListResponseWordListEntry].
func (r DLPEntryListResponse) AsUnion() DLPEntryListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryListResponseCustomEntry],
// [DLPEntryListResponsePredefinedEntry], [DLPEntryListResponseIntegrationEntry],
// [DLPEntryListResponseExactDataEntry],
// [DLPEntryListResponseDocumentFingerprintEntry] or
// [DLPEntryListResponseWordListEntry].
type DLPEntryListResponseUnion interface {
	implementsDLPEntryListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseWordListEntry{}),
		},
	)
}

type DLPEntryListResponseCustomEntry struct {
	ID        string                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                `json:"enabled,required"`
	Name      string                              `json:"name,required"`
	Pattern   Pattern                             `json:"pattern,required"`
	Type      DLPEntryListResponseCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseCustomEntryJSON `json:"-"`
}

// dlpEntryListResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPEntryListResponseCustomEntry]
type dlpEntryListResponseCustomEntryJSON struct {
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

func (r *DLPEntryListResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseCustomEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponseCustomEntryType string

const (
	DLPEntryListResponseCustomEntryTypeCustom DLPEntryListResponseCustomEntryType = "custom"
)

func (r DLPEntryListResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryListResponsePredefinedEntry struct {
	ID         string                                        `json:"id,required" format:"uuid"`
	Confidence DLPEntryListResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                          `json:"enabled,required"`
	Name       string                                        `json:"name,required"`
	Type       DLPEntryListResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                        `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryListResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryListResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPEntryListResponsePredefinedEntry]
type dlpEntryListResponsePredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryListResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponsePredefinedEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                              `json:"available,required"`
	JSON      dlpEntryListResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryListResponsePredefinedEntryConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryListResponsePredefinedEntryConfidence]
type dlpEntryListResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryListResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryListResponsePredefinedEntryType string

const (
	DLPEntryListResponsePredefinedEntryTypePredefined DLPEntryListResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryListResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryListResponseIntegrationEntry struct {
	ID        string                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                     `json:"enabled,required"`
	Name      string                                   `json:"name,required"`
	Type      DLPEntryListResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                `json:"updated_at,required" format:"date-time"`
	ProfileID string                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryListResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryListResponseIntegrationEntry]
type dlpEntryListResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryListResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseIntegrationEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponseIntegrationEntryType string

const (
	DLPEntryListResponseIntegrationEntryTypeIntegration DLPEntryListResponseIntegrationEntryType = "integration"
)

func (r DLPEntryListResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryListResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                   `json:"case_sensitive,required"`
	CreatedAt     time.Time                              `json:"created_at,required" format:"date-time"`
	Enabled       bool                                   `json:"enabled,required"`
	Name          string                                 `json:"name,required"`
	Secret        bool                                   `json:"secret,required"`
	Type          DLPEntryListResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                              `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryListResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryListResponseExactDataEntryJSON contains the JSON metadata for the struct
// [DLPEntryListResponseExactDataEntry]
type dlpEntryListResponseExactDataEntryJSON struct {
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

func (r *DLPEntryListResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseExactDataEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponseExactDataEntryType string

const (
	DLPEntryListResponseExactDataEntryTypeExactData DLPEntryListResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryListResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryListResponseDocumentFingerprintEntry struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Type      DLPEntryListResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryListResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryListResponseDocumentFingerprintEntryJSON contains the JSON metadata for
// the struct [DLPEntryListResponseDocumentFingerprintEntry]
type dlpEntryListResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryListResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseDocumentFingerprintEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponseDocumentFingerprintEntryType string

const (
	DLPEntryListResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryListResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryListResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryListResponseWordListEntry struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	Type      DLPEntryListResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                           `json:"word_list,required"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseWordListEntryJSON `json:"-"`
}

// dlpEntryListResponseWordListEntryJSON contains the JSON metadata for the struct
// [DLPEntryListResponseWordListEntry]
type dlpEntryListResponseWordListEntryJSON struct {
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

func (r *DLPEntryListResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseWordListEntry) implementsDLPEntryListResponse() {}

type DLPEntryListResponseWordListEntryType string

const (
	DLPEntryListResponseWordListEntryTypeWordList DLPEntryListResponseWordListEntryType = "word_list"
)

func (r DLPEntryListResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryListResponseType string

const (
	DLPEntryListResponseTypeCustom              DLPEntryListResponseType = "custom"
	DLPEntryListResponseTypePredefined          DLPEntryListResponseType = "predefined"
	DLPEntryListResponseTypeIntegration         DLPEntryListResponseType = "integration"
	DLPEntryListResponseTypeExactData           DLPEntryListResponseType = "exact_data"
	DLPEntryListResponseTypeDocumentFingerprint DLPEntryListResponseType = "document_fingerprint"
	DLPEntryListResponseTypeWordList            DLPEntryListResponseType = "word_list"
)

func (r DLPEntryListResponseType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseTypeCustom, DLPEntryListResponseTypePredefined, DLPEntryListResponseTypeIntegration, DLPEntryListResponseTypeExactData, DLPEntryListResponseTypeDocumentFingerprint, DLPEntryListResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryDeleteResponse = interface{}

type DLPEntryGetResponse struct {
	ID      string                  `json:"id,required" format:"uuid"`
	Enabled bool                    `json:"enabled,required"`
	Name    string                  `json:"name,required"`
	Type    DLPEntryGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryGetResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}             `json:"word_list"`
	JSON     dlpEntryGetResponseJSON `json:"-"`
	union    DLPEntryGetResponseUnion
}

// dlpEntryGetResponseJSON contains the JSON metadata for the struct
// [DLPEntryGetResponse]
type dlpEntryGetResponseJSON struct {
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

func (r dlpEntryGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryGetResponseCustomEntry],
// [DLPEntryGetResponsePredefinedEntry], [DLPEntryGetResponseIntegrationEntry],
// [DLPEntryGetResponseExactDataEntry],
// [DLPEntryGetResponseDocumentFingerprintEntry],
// [DLPEntryGetResponseWordListEntry].
func (r DLPEntryGetResponse) AsUnion() DLPEntryGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryGetResponseCustomEntry],
// [DLPEntryGetResponsePredefinedEntry], [DLPEntryGetResponseIntegrationEntry],
// [DLPEntryGetResponseExactDataEntry],
// [DLPEntryGetResponseDocumentFingerprintEntry] or
// [DLPEntryGetResponseWordListEntry].
type DLPEntryGetResponseUnion interface {
	implementsDLPEntryGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseWordListEntry{}),
		},
	)
}

type DLPEntryGetResponseCustomEntry struct {
	ID        string                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                               `json:"enabled,required"`
	Name      string                             `json:"name,required"`
	Pattern   Pattern                            `json:"pattern,required"`
	Type      DLPEntryGetResponseCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryGetResponseCustomEntryJSON `json:"-"`
}

// dlpEntryGetResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseCustomEntry]
type dlpEntryGetResponseCustomEntryJSON struct {
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

func (r *DLPEntryGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseCustomEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseCustomEntryType string

const (
	DLPEntryGetResponseCustomEntryTypeCustom DLPEntryGetResponseCustomEntryType = "custom"
)

func (r DLPEntryGetResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryGetResponsePredefinedEntry struct {
	ID         string                                       `json:"id,required" format:"uuid"`
	Confidence DLPEntryGetResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                         `json:"enabled,required"`
	Name       string                                       `json:"name,required"`
	Type       DLPEntryGetResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpEntryGetResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponsePredefinedEntry]
type dlpEntryGetResponsePredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponsePredefinedEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                             `json:"available,required"`
	JSON      dlpEntryGetResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryGetResponsePredefinedEntryConfidence]
type dlpEntryGetResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryGetResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponsePredefinedEntryType string

const (
	DLPEntryGetResponsePredefinedEntryTypePredefined DLPEntryGetResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryGetResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryGetResponseIntegrationEntry struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                    `json:"enabled,required"`
	Name      string                                  `json:"name,required"`
	Type      DLPEntryGetResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryGetResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryGetResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseIntegrationEntry]
type dlpEntryGetResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseIntegrationEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseIntegrationEntryType string

const (
	DLPEntryGetResponseIntegrationEntryTypeIntegration DLPEntryGetResponseIntegrationEntryType = "integration"
)

func (r DLPEntryGetResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryGetResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                  `json:"case_sensitive,required"`
	CreatedAt     time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled       bool                                  `json:"enabled,required"`
	Name          string                                `json:"name,required"`
	Secret        bool                                  `json:"secret,required"`
	Type          DLPEntryGetResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryGetResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryGetResponseExactDataEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseExactDataEntry]
type dlpEntryGetResponseExactDataEntryJSON struct {
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

func (r *DLPEntryGetResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseExactDataEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseExactDataEntryType string

const (
	DLPEntryGetResponseExactDataEntryTypeExactData DLPEntryGetResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryGetResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryGetResponseDocumentFingerprintEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      DLPEntryGetResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryGetResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryGetResponseDocumentFingerprintEntryJSON contains the JSON metadata for
// the struct [DLPEntryGetResponseDocumentFingerprintEntry]
type dlpEntryGetResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseDocumentFingerprintEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseDocumentFingerprintEntryType string

const (
	DLPEntryGetResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryGetResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryGetResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryGetResponseWordListEntry struct {
	ID        string                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                 `json:"enabled,required"`
	Name      string                               `json:"name,required"`
	Type      DLPEntryGetResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                          `json:"word_list,required"`
	ProfileID string                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryGetResponseWordListEntryJSON `json:"-"`
}

// dlpEntryGetResponseWordListEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseWordListEntry]
type dlpEntryGetResponseWordListEntryJSON struct {
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

func (r *DLPEntryGetResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseWordListEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseWordListEntryType string

const (
	DLPEntryGetResponseWordListEntryTypeWordList DLPEntryGetResponseWordListEntryType = "word_list"
)

func (r DLPEntryGetResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryGetResponseType string

const (
	DLPEntryGetResponseTypeCustom              DLPEntryGetResponseType = "custom"
	DLPEntryGetResponseTypePredefined          DLPEntryGetResponseType = "predefined"
	DLPEntryGetResponseTypeIntegration         DLPEntryGetResponseType = "integration"
	DLPEntryGetResponseTypeExactData           DLPEntryGetResponseType = "exact_data"
	DLPEntryGetResponseTypeDocumentFingerprint DLPEntryGetResponseType = "document_fingerprint"
	DLPEntryGetResponseTypeWordList            DLPEntryGetResponseType = "word_list"
)

func (r DLPEntryGetResponseType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseTypeCustom, DLPEntryGetResponseTypePredefined, DLPEntryGetResponseTypeIntegration, DLPEntryGetResponseTypeExactData, DLPEntryGetResponseTypeDocumentFingerprint, DLPEntryGetResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryNewParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Enabled   param.Field[bool]         `json:"enabled,required"`
	Name      param.Field[string]       `json:"name,required"`
	Pattern   param.Field[PatternParam] `json:"pattern,required"`
	ProfileID param.Field[string]       `json:"profile_id,required" format:"uuid"`
}

func (r DLPEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryNewResponseEnvelope struct {
	Errors   []DLPEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryNewResponse                `json:"result"`
	JSON    dlpEntryNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryNewResponseEnvelope]
type dlpEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DLPEntryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DLPEntryNewResponseEnvelopeErrors]
type dlpEntryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    dlpEntryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeErrorsSource]
type dlpEntryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPEntryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeMessages]
type dlpEntryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpEntryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeMessagesSource]
type dlpEntryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryNewResponseEnvelopeSuccess bool

const (
	DLPEntryNewResponseEnvelopeSuccessTrue DLPEntryNewResponseEnvelopeSuccess = true
)

func (r DLPEntryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryUpdateParams struct {
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      DLPEntryUpdateParamsBodyUnion `json:"body,required"`
}

func (r DLPEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPEntryUpdateParamsBody struct {
	Type    param.Field[DLPEntryUpdateParamsBodyType] `json:"type,required"`
	Enabled param.Field[bool]                         `json:"enabled"`
	Name    param.Field[string]                       `json:"name"`
	Pattern param.Field[PatternParam]                 `json:"pattern"`
}

func (r DLPEntryUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBody) implementsDLPEntryUpdateParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPEntryUpdateParamsBodyCustom],
// [zero_trust.DLPEntryUpdateParamsBodyPredefined],
// [zero_trust.DLPEntryUpdateParamsBodyIntegration], [DLPEntryUpdateParamsBody].
type DLPEntryUpdateParamsBodyUnion interface {
	implementsDLPEntryUpdateParamsBodyUnion()
}

type DLPEntryUpdateParamsBodyCustom struct {
	Name    param.Field[string]                             `json:"name,required"`
	Pattern param.Field[PatternParam]                       `json:"pattern,required"`
	Type    param.Field[DLPEntryUpdateParamsBodyCustomType] `json:"type,required"`
	Enabled param.Field[bool]                               `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyCustom) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyCustomType string

const (
	DLPEntryUpdateParamsBodyCustomTypeCustom DLPEntryUpdateParamsBodyCustomType = "custom"
)

func (r DLPEntryUpdateParamsBodyCustomType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyPredefined struct {
	Type    param.Field[DLPEntryUpdateParamsBodyPredefinedType] `json:"type,required"`
	Enabled param.Field[bool]                                   `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyPredefined) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyPredefinedType string

const (
	DLPEntryUpdateParamsBodyPredefinedTypePredefined DLPEntryUpdateParamsBodyPredefinedType = "predefined"
)

func (r DLPEntryUpdateParamsBodyPredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyIntegration struct {
	Type    param.Field[DLPEntryUpdateParamsBodyIntegrationType] `json:"type,required"`
	Enabled param.Field[bool]                                    `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyIntegration) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyIntegrationType string

const (
	DLPEntryUpdateParamsBodyIntegrationTypeIntegration DLPEntryUpdateParamsBodyIntegrationType = "integration"
)

func (r DLPEntryUpdateParamsBodyIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyType string

const (
	DLPEntryUpdateParamsBodyTypeCustom      DLPEntryUpdateParamsBodyType = "custom"
	DLPEntryUpdateParamsBodyTypePredefined  DLPEntryUpdateParamsBodyType = "predefined"
	DLPEntryUpdateParamsBodyTypeIntegration DLPEntryUpdateParamsBodyType = "integration"
)

func (r DLPEntryUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyTypeCustom, DLPEntryUpdateParamsBodyTypePredefined, DLPEntryUpdateParamsBodyTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateResponseEnvelope struct {
	Errors   []DLPEntryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryUpdateResponse                `json:"result"`
	JSON    dlpEntryUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseEnvelope]
type dlpEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPEntryUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseEnvelopeErrors]
type dlpEntryUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpEntryUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponseEnvelopeErrorsSource]
type dlpEntryUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPEntryUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseEnvelopeMessages]
type dlpEntryUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpEntryUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponseEnvelopeMessagesSource]
type dlpEntryUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryUpdateResponseEnvelopeSuccessTrue DLPEntryUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryDeleteResponseEnvelope struct {
	Errors   []DLPEntryDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryDeleteResponseEnvelope]
type dlpEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPEntryDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryDeleteResponseEnvelopeErrors]
type dlpEntryDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpEntryDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryDeleteResponseEnvelopeErrorsSource]
type dlpEntryDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPEntryDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryDeleteResponseEnvelopeMessages]
type dlpEntryDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpEntryDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEntryDeleteResponseEnvelopeMessagesSource]
type dlpEntryDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryDeleteResponseEnvelopeSuccessTrue DLPEntryDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryGetResponseEnvelope struct {
	Errors   []DLPEntryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryGetResponse                `json:"result"`
	JSON    dlpEntryGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseEnvelope]
type dlpEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DLPEntryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseEnvelopeErrors]
type dlpEntryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    dlpEntryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeErrorsSource]
type dlpEntryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPEntryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeMessages]
type dlpEntryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpEntryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeMessagesSource]
type dlpEntryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryGetResponseEnvelopeSuccess bool

const (
	DLPEntryGetResponseEnvelopeSuccessTrue DLPEntryGetResponseEnvelopeSuccess = true
)

func (r DLPEntryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
