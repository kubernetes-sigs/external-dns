// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SettingImpersonationRegistryService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingImpersonationRegistryService] method instead.
type SettingImpersonationRegistryService struct {
	Options []option.RequestOption
}

// NewSettingImpersonationRegistryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingImpersonationRegistryService(opts ...option.RequestOption) (r *SettingImpersonationRegistryService) {
	r = &SettingImpersonationRegistryService{}
	r.Options = opts
	return
}

// Create an entry in impersonation registry
func (r *SettingImpersonationRegistryService) New(ctx context.Context, params SettingImpersonationRegistryNewParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryNewResponse, err error) {
	var env SettingImpersonationRegistryNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists, searches, and sorts entries in the impersonation registry.
func (r *SettingImpersonationRegistryService) List(ctx context.Context, params SettingImpersonationRegistryListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingImpersonationRegistryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry", params.AccountID)
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

// Lists, searches, and sorts entries in the impersonation registry.
func (r *SettingImpersonationRegistryService) ListAutoPaging(ctx context.Context, params SettingImpersonationRegistryListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingImpersonationRegistryListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an entry from impersonation registry
func (r *SettingImpersonationRegistryService) Delete(ctx context.Context, displayNameID int64, body SettingImpersonationRegistryDeleteParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryDeleteResponse, err error) {
	var env SettingImpersonationRegistryDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", body.AccountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an entry in impersonation registry
func (r *SettingImpersonationRegistryService) Edit(ctx context.Context, displayNameID int64, params SettingImpersonationRegistryEditParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryEditResponse, err error) {
	var env SettingImpersonationRegistryEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", params.AccountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an entry in impersonation registry
func (r *SettingImpersonationRegistryService) Get(ctx context.Context, displayNameID int64, query SettingImpersonationRegistryGetParams, opts ...option.RequestOption) (res *SettingImpersonationRegistryGetResponse, err error) {
	var env SettingImpersonationRegistryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/impersonation_registry/%v", query.AccountID, displayNameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingImpersonationRegistryNewResponse struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                      `json:"external_directory_node_id,nullable"`
	Provenance              string                                      `json:"provenance,nullable"`
	JSON                    settingImpersonationRegistryNewResponseJSON `json:"-"`
}

// settingImpersonationRegistryNewResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryNewResponse]
type settingImpersonationRegistryNewResponseJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryListResponse struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                       `json:"external_directory_node_id,nullable"`
	Provenance              string                                       `json:"provenance,nullable"`
	JSON                    settingImpersonationRegistryListResponseJSON `json:"-"`
}

// settingImpersonationRegistryListResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryListResponse]
type settingImpersonationRegistryListResponseJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryDeleteResponse struct {
	ID   int64                                          `json:"id,required"`
	JSON settingImpersonationRegistryDeleteResponseJSON `json:"-"`
}

// settingImpersonationRegistryDeleteResponseJSON contains the JSON metadata for
// the struct [SettingImpersonationRegistryDeleteResponse]
type settingImpersonationRegistryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditResponse struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                       `json:"external_directory_node_id,nullable"`
	Provenance              string                                       `json:"provenance,nullable"`
	JSON                    settingImpersonationRegistryEditResponseJSON `json:"-"`
}

// settingImpersonationRegistryEditResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryEditResponse]
type settingImpersonationRegistryEditResponseJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetResponse struct {
	ID              int64     `json:"id,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	Email           string    `json:"email,required"`
	IsEmailRegex    bool      `json:"is_email_regex,required"`
	LastModified    time.Time `json:"last_modified,required" format:"date-time"`
	Name            string    `json:"name,required"`
	Comments        string    `json:"comments,nullable"`
	DirectoryID     int64     `json:"directory_id,nullable"`
	DirectoryNodeID int64     `json:"directory_node_id,nullable"`
	// Deprecated: deprecated
	ExternalDirectoryNodeID string                                      `json:"external_directory_node_id,nullable"`
	Provenance              string                                      `json:"provenance,nullable"`
	JSON                    settingImpersonationRegistryGetResponseJSON `json:"-"`
}

// settingImpersonationRegistryGetResponseJSON contains the JSON metadata for the
// struct [SettingImpersonationRegistryGetResponse]
type settingImpersonationRegistryGetResponseJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	Email                   apijson.Field
	IsEmailRegex            apijson.Field
	LastModified            apijson.Field
	Name                    apijson.Field
	Comments                apijson.Field
	DirectoryID             apijson.Field
	DirectoryNodeID         apijson.Field
	ExternalDirectoryNodeID apijson.Field
	Provenance              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryNewParams struct {
	// Account Identifier
	AccountID    param.Field[string] `path:"account_id,required"`
	Email        param.Field[string] `json:"email,required"`
	IsEmailRegex param.Field[bool]   `json:"is_email_regex,required"`
	Name         param.Field[string] `json:"name,required"`
}

func (r SettingImpersonationRegistryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImpersonationRegistryNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                               `json:"errors,required"`
	Messages []shared.ResponseInfo                               `json:"messages,required"`
	Result   SettingImpersonationRegistryNewResponse             `json:"result,required"`
	Success  bool                                                `json:"success,required"`
	JSON     settingImpersonationRegistryNewResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryNewResponseEnvelope]
type settingImpersonationRegistryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The sorting direction.
	Direction param.Field[SettingImpersonationRegistryListParamsDirection] `query:"direction"`
	// The field to sort by.
	Order param.Field[SettingImpersonationRegistryListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage    param.Field[int64]                                            `query:"per_page"`
	Provenance param.Field[SettingImpersonationRegistryListParamsProvenance] `query:"provenance"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingImpersonationRegistryListParams]'s query parameters
// as `url.Values`.
func (r SettingImpersonationRegistryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingImpersonationRegistryListParamsDirection string

const (
	SettingImpersonationRegistryListParamsDirectionAsc  SettingImpersonationRegistryListParamsDirection = "asc"
	SettingImpersonationRegistryListParamsDirectionDesc SettingImpersonationRegistryListParamsDirection = "desc"
)

func (r SettingImpersonationRegistryListParamsDirection) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsDirectionAsc, SettingImpersonationRegistryListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingImpersonationRegistryListParamsOrder string

const (
	SettingImpersonationRegistryListParamsOrderName      SettingImpersonationRegistryListParamsOrder = "name"
	SettingImpersonationRegistryListParamsOrderEmail     SettingImpersonationRegistryListParamsOrder = "email"
	SettingImpersonationRegistryListParamsOrderCreatedAt SettingImpersonationRegistryListParamsOrder = "created_at"
)

func (r SettingImpersonationRegistryListParamsOrder) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsOrderName, SettingImpersonationRegistryListParamsOrderEmail, SettingImpersonationRegistryListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingImpersonationRegistryListParamsProvenance string

const (
	SettingImpersonationRegistryListParamsProvenanceA1SInternal           SettingImpersonationRegistryListParamsProvenance = "A1S_INTERNAL"
	SettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365   SettingImpersonationRegistryListParamsProvenance = "SNOOPY-CASB_OFFICE_365"
	SettingImpersonationRegistryListParamsProvenanceSnoopyOffice365       SettingImpersonationRegistryListParamsProvenance = "SNOOPY-OFFICE_365"
	SettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory SettingImpersonationRegistryListParamsProvenance = "SNOOPY-GOOGLE_DIRECTORY"
)

func (r SettingImpersonationRegistryListParamsProvenance) IsKnown() bool {
	switch r {
	case SettingImpersonationRegistryListParamsProvenanceA1SInternal, SettingImpersonationRegistryListParamsProvenanceSnoopyCasbOffice365, SettingImpersonationRegistryListParamsProvenanceSnoopyOffice365, SettingImpersonationRegistryListParamsProvenanceSnoopyGoogleDirectory:
		return true
	}
	return false
}

type SettingImpersonationRegistryDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingImpersonationRegistryDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                  `json:"errors,required"`
	Messages []shared.ResponseInfo                                  `json:"messages,required"`
	Result   SettingImpersonationRegistryDeleteResponse             `json:"result,required"`
	Success  bool                                                   `json:"success,required"`
	JSON     settingImpersonationRegistryDeleteResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingImpersonationRegistryDeleteResponseEnvelope]
type settingImpersonationRegistryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryEditParams struct {
	// Account Identifier
	AccountID    param.Field[string] `path:"account_id,required"`
	Email        param.Field[string] `json:"email"`
	IsEmailRegex param.Field[bool]   `json:"is_email_regex"`
	Name         param.Field[string] `json:"name"`
}

func (r SettingImpersonationRegistryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImpersonationRegistryEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                `json:"errors,required"`
	Messages []shared.ResponseInfo                                `json:"messages,required"`
	Result   SettingImpersonationRegistryEditResponse             `json:"result,required"`
	Success  bool                                                 `json:"success,required"`
	JSON     settingImpersonationRegistryEditResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryEditResponseEnvelope]
type settingImpersonationRegistryEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImpersonationRegistryGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingImpersonationRegistryGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                               `json:"errors,required"`
	Messages []shared.ResponseInfo                               `json:"messages,required"`
	Result   SettingImpersonationRegistryGetResponse             `json:"result,required"`
	Success  bool                                                `json:"success,required"`
	JSON     settingImpersonationRegistryGetResponseEnvelopeJSON `json:"-"`
}

// settingImpersonationRegistryGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingImpersonationRegistryGetResponseEnvelope]
type settingImpersonationRegistryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImpersonationRegistryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingImpersonationRegistryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
