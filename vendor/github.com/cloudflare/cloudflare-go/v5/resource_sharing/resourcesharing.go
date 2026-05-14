// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_sharing

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

// ResourceSharingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceSharingService] method instead.
type ResourceSharingService struct {
	Options    []option.RequestOption
	Recipients *RecipientService
	Resources  *ResourceService
}

// NewResourceSharingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceSharingService(opts ...option.RequestOption) (r *ResourceSharingService) {
	r = &ResourceSharingService{}
	r.Options = opts
	r.Recipients = NewRecipientService(opts...)
	r.Resources = NewResourceService(opts...)
	return
}

// Create a new share
func (r *ResourceSharingService) New(ctx context.Context, params ResourceSharingNewParams, opts ...option.RequestOption) (res *ResourceSharingNewResponse, err error) {
	var env ResourceSharingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating is not immediate, an updated share object with a new status will be
// returned.
func (r *ResourceSharingService) Update(ctx context.Context, shareID string, params ResourceSharingUpdateParams, opts ...option.RequestOption) (res *ResourceSharingUpdateResponse, err error) {
	var env ResourceSharingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", params.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all account shares.
func (r *ResourceSharingService) List(ctx context.Context, params ResourceSharingListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceSharingListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", params.AccountID)
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

// Lists all account shares.
func (r *ResourceSharingService) ListAutoPaging(ctx context.Context, params ResourceSharingListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceSharingListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletion is not immediate, an updated share object with a new status will be
// returned.
func (r *ResourceSharingService) Delete(ctx context.Context, shareID string, body ResourceSharingDeleteParams, opts ...option.RequestOption) (res *ResourceSharingDeleteResponse, err error) {
	var env ResourceSharingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", body.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches share by ID.
func (r *ResourceSharingService) Get(ctx context.Context, shareID string, query ResourceSharingGetParams, opts ...option.RequestOption) (res *ResourceSharingGetResponse, err error) {
	var env ResourceSharingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", query.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceSharingNewResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                               `json:"organization_id,required"`
	Status         ResourceSharingNewResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingNewResponseTargetType `json:"target_type,required"`
	Kind           ResourceSharingNewResponseKind       `json:"kind"`
	JSON           resourceSharingNewResponseJSON       `json:"-"`
}

// resourceSharingNewResponseJSON contains the JSON metadata for the struct
// [ResourceSharingNewResponse]
type resourceSharingNewResponseJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceSharingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingNewResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingNewResponseStatus string

const (
	ResourceSharingNewResponseStatusActive   ResourceSharingNewResponseStatus = "active"
	ResourceSharingNewResponseStatusDeleting ResourceSharingNewResponseStatus = "deleting"
	ResourceSharingNewResponseStatusDeleted  ResourceSharingNewResponseStatus = "deleted"
)

func (r ResourceSharingNewResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseStatusActive, ResourceSharingNewResponseStatusDeleting, ResourceSharingNewResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingNewResponseTargetType string

const (
	ResourceSharingNewResponseTargetTypeAccount      ResourceSharingNewResponseTargetType = "account"
	ResourceSharingNewResponseTargetTypeOrganization ResourceSharingNewResponseTargetType = "organization"
)

func (r ResourceSharingNewResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseTargetTypeAccount, ResourceSharingNewResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingNewResponseKind string

const (
	ResourceSharingNewResponseKindSent     ResourceSharingNewResponseKind = "sent"
	ResourceSharingNewResponseKindReceived ResourceSharingNewResponseKind = "received"
)

func (r ResourceSharingNewResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseKindSent, ResourceSharingNewResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingUpdateResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                  `json:"organization_id,required"`
	Status         ResourceSharingUpdateResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingUpdateResponseTargetType `json:"target_type,required"`
	Kind           ResourceSharingUpdateResponseKind       `json:"kind"`
	JSON           resourceSharingUpdateResponseJSON       `json:"-"`
}

// resourceSharingUpdateResponseJSON contains the JSON metadata for the struct
// [ResourceSharingUpdateResponse]
type resourceSharingUpdateResponseJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceSharingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingUpdateResponseStatus string

const (
	ResourceSharingUpdateResponseStatusActive   ResourceSharingUpdateResponseStatus = "active"
	ResourceSharingUpdateResponseStatusDeleting ResourceSharingUpdateResponseStatus = "deleting"
	ResourceSharingUpdateResponseStatusDeleted  ResourceSharingUpdateResponseStatus = "deleted"
)

func (r ResourceSharingUpdateResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseStatusActive, ResourceSharingUpdateResponseStatusDeleting, ResourceSharingUpdateResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingUpdateResponseTargetType string

const (
	ResourceSharingUpdateResponseTargetTypeAccount      ResourceSharingUpdateResponseTargetType = "account"
	ResourceSharingUpdateResponseTargetTypeOrganization ResourceSharingUpdateResponseTargetType = "organization"
)

func (r ResourceSharingUpdateResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseTargetTypeAccount, ResourceSharingUpdateResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingUpdateResponseKind string

const (
	ResourceSharingUpdateResponseKindSent     ResourceSharingUpdateResponseKind = "sent"
	ResourceSharingUpdateResponseKindReceived ResourceSharingUpdateResponseKind = "received"
)

func (r ResourceSharingUpdateResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseKindSent, ResourceSharingUpdateResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingListResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                `json:"organization_id,required"`
	Status         ResourceSharingListResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingListResponseTargetType `json:"target_type,required"`
	Kind           ResourceSharingListResponseKind       `json:"kind"`
	JSON           resourceSharingListResponseJSON       `json:"-"`
}

// resourceSharingListResponseJSON contains the JSON metadata for the struct
// [ResourceSharingListResponse]
type resourceSharingListResponseJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceSharingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingListResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingListResponseStatus string

const (
	ResourceSharingListResponseStatusActive   ResourceSharingListResponseStatus = "active"
	ResourceSharingListResponseStatusDeleting ResourceSharingListResponseStatus = "deleting"
	ResourceSharingListResponseStatusDeleted  ResourceSharingListResponseStatus = "deleted"
)

func (r ResourceSharingListResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseStatusActive, ResourceSharingListResponseStatusDeleting, ResourceSharingListResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingListResponseTargetType string

const (
	ResourceSharingListResponseTargetTypeAccount      ResourceSharingListResponseTargetType = "account"
	ResourceSharingListResponseTargetTypeOrganization ResourceSharingListResponseTargetType = "organization"
)

func (r ResourceSharingListResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseTargetTypeAccount, ResourceSharingListResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingListResponseKind string

const (
	ResourceSharingListResponseKindSent     ResourceSharingListResponseKind = "sent"
	ResourceSharingListResponseKindReceived ResourceSharingListResponseKind = "received"
)

func (r ResourceSharingListResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseKindSent, ResourceSharingListResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingDeleteResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                  `json:"organization_id,required"`
	Status         ResourceSharingDeleteResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingDeleteResponseTargetType `json:"target_type,required"`
	Kind           ResourceSharingDeleteResponseKind       `json:"kind"`
	JSON           resourceSharingDeleteResponseJSON       `json:"-"`
}

// resourceSharingDeleteResponseJSON contains the JSON metadata for the struct
// [ResourceSharingDeleteResponse]
type resourceSharingDeleteResponseJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceSharingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingDeleteResponseStatus string

const (
	ResourceSharingDeleteResponseStatusActive   ResourceSharingDeleteResponseStatus = "active"
	ResourceSharingDeleteResponseStatusDeleting ResourceSharingDeleteResponseStatus = "deleting"
	ResourceSharingDeleteResponseStatusDeleted  ResourceSharingDeleteResponseStatus = "deleted"
)

func (r ResourceSharingDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseStatusActive, ResourceSharingDeleteResponseStatusDeleting, ResourceSharingDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingDeleteResponseTargetType string

const (
	ResourceSharingDeleteResponseTargetTypeAccount      ResourceSharingDeleteResponseTargetType = "account"
	ResourceSharingDeleteResponseTargetTypeOrganization ResourceSharingDeleteResponseTargetType = "organization"
)

func (r ResourceSharingDeleteResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseTargetTypeAccount, ResourceSharingDeleteResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingDeleteResponseKind string

const (
	ResourceSharingDeleteResponseKindSent     ResourceSharingDeleteResponseKind = "sent"
	ResourceSharingDeleteResponseKindReceived ResourceSharingDeleteResponseKind = "received"
)

func (r ResourceSharingDeleteResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseKindSent, ResourceSharingDeleteResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingGetResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                               `json:"organization_id,required"`
	Status         ResourceSharingGetResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingGetResponseTargetType `json:"target_type,required"`
	Kind           ResourceSharingGetResponseKind       `json:"kind"`
	JSON           resourceSharingGetResponseJSON       `json:"-"`
}

// resourceSharingGetResponseJSON contains the JSON metadata for the struct
// [ResourceSharingGetResponse]
type resourceSharingGetResponseJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceSharingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingGetResponseStatus string

const (
	ResourceSharingGetResponseStatusActive   ResourceSharingGetResponseStatus = "active"
	ResourceSharingGetResponseStatusDeleting ResourceSharingGetResponseStatus = "deleting"
	ResourceSharingGetResponseStatusDeleted  ResourceSharingGetResponseStatus = "deleted"
)

func (r ResourceSharingGetResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseStatusActive, ResourceSharingGetResponseStatusDeleting, ResourceSharingGetResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingGetResponseTargetType string

const (
	ResourceSharingGetResponseTargetTypeAccount      ResourceSharingGetResponseTargetType = "account"
	ResourceSharingGetResponseTargetTypeOrganization ResourceSharingGetResponseTargetType = "organization"
)

func (r ResourceSharingGetResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseTargetTypeAccount, ResourceSharingGetResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingGetResponseKind string

const (
	ResourceSharingGetResponseKindSent     ResourceSharingGetResponseKind = "sent"
	ResourceSharingGetResponseKindReceived ResourceSharingGetResponseKind = "received"
)

func (r ResourceSharingGetResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseKindSent, ResourceSharingGetResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingNewParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the share.
	Name       param.Field[string]                              `json:"name,required"`
	Recipients param.Field[[]ResourceSharingNewParamsRecipient] `json:"recipients,required"`
	Resources  param.Field[[]ResourceSharingNewParamsResource]  `json:"resources,required"`
}

func (r ResourceSharingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account or organization ID must be provided.
type ResourceSharingNewParamsRecipient struct {
	// Account identifier.
	AccountID param.Field[string] `json:"account_id"`
	// Organization identifier.
	OrganizationID param.Field[string] `json:"organization_id"`
}

func (r ResourceSharingNewParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingNewParamsResource struct {
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
	// Account identifier.
	ResourceAccountID param.Field[string] `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID param.Field[string] `json:"resource_id,required"`
	// Resource Type.
	ResourceType param.Field[ResourceSharingNewParamsResourcesResourceType] `json:"resource_type,required"`
}

func (r ResourceSharingNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource Type.
type ResourceSharingNewParamsResourcesResourceType string

const (
	ResourceSharingNewParamsResourcesResourceTypeCustomRuleset ResourceSharingNewParamsResourcesResourceType = "custom-ruleset"
	ResourceSharingNewParamsResourcesResourceTypeWidget        ResourceSharingNewParamsResourcesResourceType = "widget"
)

func (r ResourceSharingNewParamsResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingNewParamsResourcesResourceTypeCustomRuleset, ResourceSharingNewParamsResourcesResourceTypeWidget:
		return true
	}
	return false
}

type ResourceSharingNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  ResourceSharingNewResponse             `json:"result"`
	JSON    resourceSharingNewResponseEnvelopeJSON `json:"-"`
}

// resourceSharingNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceSharingNewResponseEnvelope]
type resourceSharingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingUpdateParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the share.
	Name param.Field[string] `json:"name,required"`
}

func (r ResourceSharingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingUpdateResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	Result  ResourceSharingUpdateResponse             `json:"result"`
	JSON    resourceSharingUpdateResponseEnvelopeJSON `json:"-"`
}

// resourceSharingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceSharingUpdateResponseEnvelope]
type resourceSharingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to sort objects.
	Direction param.Field[ResourceSharingListParamsDirection] `query:"direction"`
	// Filter shares by kind.
	Kind param.Field[ResourceSharingListParamsKind] `query:"kind"`
	// Order shares by values in the given field.
	Order param.Field[ResourceSharingListParamsOrder] `query:"order"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter shares by status.
	Status param.Field[ResourceSharingListParamsStatus] `query:"status"`
	// Filter shares by target_type.
	TargetType param.Field[ResourceSharingListParamsTargetType] `query:"target_type"`
}

// URLQuery serializes [ResourceSharingListParams]'s query parameters as
// `url.Values`.
func (r ResourceSharingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to sort objects.
type ResourceSharingListParamsDirection string

const (
	ResourceSharingListParamsDirectionAsc  ResourceSharingListParamsDirection = "asc"
	ResourceSharingListParamsDirectionDesc ResourceSharingListParamsDirection = "desc"
)

func (r ResourceSharingListParamsDirection) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsDirectionAsc, ResourceSharingListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter shares by kind.
type ResourceSharingListParamsKind string

const (
	ResourceSharingListParamsKindSent     ResourceSharingListParamsKind = "sent"
	ResourceSharingListParamsKindReceived ResourceSharingListParamsKind = "received"
)

func (r ResourceSharingListParamsKind) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsKindSent, ResourceSharingListParamsKindReceived:
		return true
	}
	return false
}

// Order shares by values in the given field.
type ResourceSharingListParamsOrder string

const (
	ResourceSharingListParamsOrderName    ResourceSharingListParamsOrder = "name"
	ResourceSharingListParamsOrderCreated ResourceSharingListParamsOrder = "created"
)

func (r ResourceSharingListParamsOrder) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsOrderName, ResourceSharingListParamsOrderCreated:
		return true
	}
	return false
}

// Filter shares by status.
type ResourceSharingListParamsStatus string

const (
	ResourceSharingListParamsStatusActive   ResourceSharingListParamsStatus = "active"
	ResourceSharingListParamsStatusDeleting ResourceSharingListParamsStatus = "deleting"
	ResourceSharingListParamsStatusDeleted  ResourceSharingListParamsStatus = "deleted"
)

func (r ResourceSharingListParamsStatus) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsStatusActive, ResourceSharingListParamsStatusDeleting, ResourceSharingListParamsStatusDeleted:
		return true
	}
	return false
}

// Filter shares by target_type.
type ResourceSharingListParamsTargetType string

const (
	ResourceSharingListParamsTargetTypeAccount      ResourceSharingListParamsTargetType = "account"
	ResourceSharingListParamsTargetTypeOrganization ResourceSharingListParamsTargetType = "organization"
)

func (r ResourceSharingListParamsTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsTargetTypeAccount, ResourceSharingListParamsTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	Result  ResourceSharingDeleteResponse             `json:"result"`
	JSON    resourceSharingDeleteResponseEnvelopeJSON `json:"-"`
}

// resourceSharingDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceSharingDeleteResponseEnvelope]
type resourceSharingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  ResourceSharingGetResponse             `json:"result"`
	JSON    resourceSharingGetResponseEnvelopeJSON `json:"-"`
}

// resourceSharingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceSharingGetResponseEnvelope]
type resourceSharingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
