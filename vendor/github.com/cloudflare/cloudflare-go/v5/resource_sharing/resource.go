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

// ResourceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceService] method instead.
type ResourceService struct {
	Options []option.RequestOption
}

// NewResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceService(opts ...option.RequestOption) (r *ResourceService) {
	r = &ResourceService{}
	r.Options = opts
	return
}

// Create a new share resource
func (r *ResourceService) New(ctx context.Context, shareID string, params ResourceNewParams, opts ...option.RequestOption) (res *ResourceNewResponse, err error) {
	var env ResourceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", params.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update is not immediate, an updated share resource object with a new status will
// be returned.
func (r *ResourceService) Update(ctx context.Context, shareID string, resourceID string, params ResourceUpdateParams, opts ...option.RequestOption) (res *ResourceUpdateResponse, err error) {
	var env ResourceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", params.AccountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List share resources by share ID.
func (r *ResourceService) List(ctx context.Context, shareID string, params ResourceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", params.AccountID, shareID)
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

// List share resources by share ID.
func (r *ResourceService) ListAutoPaging(ctx context.Context, shareID string, params ResourceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, shareID, params, opts...))
}

// Deletion is not immediate, an updated share resource object with a new status
// will be returned.
func (r *ResourceService) Delete(ctx context.Context, shareID string, resourceID string, body ResourceDeleteParams, opts ...option.RequestOption) (res *ResourceDeleteResponse, err error) {
	var env ResourceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", body.AccountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get share resource by ID.
func (r *ResourceService) Get(ctx context.Context, shareID string, resourceID string, query ResourceGetParams, opts ...option.RequestOption) (res *ResourceGetResponse, err error) {
	var env ResourceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", query.AccountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceNewResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceNewResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceNewResponseStatus `json:"status,required"`
	JSON   resourceNewResponseJSON   `json:"-"`
}

// resourceNewResponseJSON contains the JSON metadata for the struct
// [ResourceNewResponse]
type resourceNewResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceNewResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceNewResponseResourceType string

const (
	ResourceNewResponseResourceTypeCustomRuleset ResourceNewResponseResourceType = "custom-ruleset"
	ResourceNewResponseResourceTypeWidget        ResourceNewResponseResourceType = "widget"
)

func (r ResourceNewResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceNewResponseResourceTypeCustomRuleset, ResourceNewResponseResourceTypeWidget:
		return true
	}
	return false
}

// Resource Status.
type ResourceNewResponseStatus string

const (
	ResourceNewResponseStatusActive   ResourceNewResponseStatus = "active"
	ResourceNewResponseStatusDeleting ResourceNewResponseStatus = "deleting"
	ResourceNewResponseStatusDeleted  ResourceNewResponseStatus = "deleted"
)

func (r ResourceNewResponseStatus) IsKnown() bool {
	switch r {
	case ResourceNewResponseStatusActive, ResourceNewResponseStatusDeleting, ResourceNewResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceUpdateResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceUpdateResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceUpdateResponseStatus `json:"status,required"`
	JSON   resourceUpdateResponseJSON   `json:"-"`
}

// resourceUpdateResponseJSON contains the JSON metadata for the struct
// [ResourceUpdateResponse]
type resourceUpdateResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceUpdateResponseResourceType string

const (
	ResourceUpdateResponseResourceTypeCustomRuleset ResourceUpdateResponseResourceType = "custom-ruleset"
	ResourceUpdateResponseResourceTypeWidget        ResourceUpdateResponseResourceType = "widget"
)

func (r ResourceUpdateResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceUpdateResponseResourceTypeCustomRuleset, ResourceUpdateResponseResourceTypeWidget:
		return true
	}
	return false
}

// Resource Status.
type ResourceUpdateResponseStatus string

const (
	ResourceUpdateResponseStatusActive   ResourceUpdateResponseStatus = "active"
	ResourceUpdateResponseStatusDeleting ResourceUpdateResponseStatus = "deleting"
	ResourceUpdateResponseStatusDeleted  ResourceUpdateResponseStatus = "deleted"
)

func (r ResourceUpdateResponseStatus) IsKnown() bool {
	switch r {
	case ResourceUpdateResponseStatusActive, ResourceUpdateResponseStatusDeleting, ResourceUpdateResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceListResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceListResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceListResponseStatus `json:"status,required"`
	JSON   resourceListResponseJSON   `json:"-"`
}

// resourceListResponseJSON contains the JSON metadata for the struct
// [ResourceListResponse]
type resourceListResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceListResponseResourceType string

const (
	ResourceListResponseResourceTypeCustomRuleset ResourceListResponseResourceType = "custom-ruleset"
	ResourceListResponseResourceTypeWidget        ResourceListResponseResourceType = "widget"
)

func (r ResourceListResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseResourceTypeCustomRuleset, ResourceListResponseResourceTypeWidget:
		return true
	}
	return false
}

// Resource Status.
type ResourceListResponseStatus string

const (
	ResourceListResponseStatusActive   ResourceListResponseStatus = "active"
	ResourceListResponseStatusDeleting ResourceListResponseStatus = "deleting"
	ResourceListResponseStatusDeleted  ResourceListResponseStatus = "deleted"
)

func (r ResourceListResponseStatus) IsKnown() bool {
	switch r {
	case ResourceListResponseStatusActive, ResourceListResponseStatusDeleting, ResourceListResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceDeleteResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceDeleteResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceDeleteResponseStatus `json:"status,required"`
	JSON   resourceDeleteResponseJSON   `json:"-"`
}

// resourceDeleteResponseJSON contains the JSON metadata for the struct
// [ResourceDeleteResponse]
type resourceDeleteResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceDeleteResponseResourceType string

const (
	ResourceDeleteResponseResourceTypeCustomRuleset ResourceDeleteResponseResourceType = "custom-ruleset"
	ResourceDeleteResponseResourceTypeWidget        ResourceDeleteResponseResourceType = "widget"
)

func (r ResourceDeleteResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceDeleteResponseResourceTypeCustomRuleset, ResourceDeleteResponseResourceTypeWidget:
		return true
	}
	return false
}

// Resource Status.
type ResourceDeleteResponseStatus string

const (
	ResourceDeleteResponseStatusActive   ResourceDeleteResponseStatus = "active"
	ResourceDeleteResponseStatusDeleting ResourceDeleteResponseStatus = "deleting"
	ResourceDeleteResponseStatusDeleted  ResourceDeleteResponseStatus = "deleted"
)

func (r ResourceDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ResourceDeleteResponseStatusActive, ResourceDeleteResponseStatusDeleting, ResourceDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceGetResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceGetResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceGetResponseStatus `json:"status,required"`
	JSON   resourceGetResponseJSON   `json:"-"`
}

// resourceGetResponseJSON contains the JSON metadata for the struct
// [ResourceGetResponse]
type resourceGetResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceGetResponseResourceType string

const (
	ResourceGetResponseResourceTypeCustomRuleset ResourceGetResponseResourceType = "custom-ruleset"
	ResourceGetResponseResourceTypeWidget        ResourceGetResponseResourceType = "widget"
)

func (r ResourceGetResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseResourceTypeCustomRuleset, ResourceGetResponseResourceTypeWidget:
		return true
	}
	return false
}

// Resource Status.
type ResourceGetResponseStatus string

const (
	ResourceGetResponseStatusActive   ResourceGetResponseStatus = "active"
	ResourceGetResponseStatusDeleting ResourceGetResponseStatus = "deleting"
	ResourceGetResponseStatusDeleted  ResourceGetResponseStatus = "deleted"
)

func (r ResourceGetResponseStatus) IsKnown() bool {
	switch r {
	case ResourceGetResponseStatusActive, ResourceGetResponseStatusDeleting, ResourceGetResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceNewParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
	// Account identifier.
	ResourceAccountID param.Field[string] `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID param.Field[string] `json:"resource_id,required"`
	// Resource Type.
	ResourceType param.Field[ResourceNewParamsResourceType] `json:"resource_type,required"`
}

func (r ResourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource Type.
type ResourceNewParamsResourceType string

const (
	ResourceNewParamsResourceTypeCustomRuleset ResourceNewParamsResourceType = "custom-ruleset"
	ResourceNewParamsResourceTypeWidget        ResourceNewParamsResourceType = "widget"
)

func (r ResourceNewParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceNewParamsResourceTypeCustomRuleset, ResourceNewParamsResourceTypeWidget:
		return true
	}
	return false
}

type ResourceNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  ResourceNewResponse             `json:"result"`
	JSON    resourceNewResponseEnvelopeJSON `json:"-"`
}

// resourceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceNewResponseEnvelope]
type resourceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceUpdateParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
}

func (r ResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceUpdateResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                               `json:"success,required"`
	Result  ResourceUpdateResponse             `json:"result"`
	JSON    resourceUpdateResponseEnvelopeJSON `json:"-"`
}

// resourceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceUpdateResponseEnvelope]
type resourceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter share resources by resource_type.
	ResourceType param.Field[ResourceListParamsResourceType] `query:"resource_type"`
	// Filter share resources by status.
	Status param.Field[ResourceListParamsStatus] `query:"status"`
}

// URLQuery serializes [ResourceListParams]'s query parameters as `url.Values`.
func (r ResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter share resources by resource_type.
type ResourceListParamsResourceType string

const (
	ResourceListParamsResourceTypeCustomRuleset ResourceListParamsResourceType = "custom-ruleset"
	ResourceListParamsResourceTypeWidget        ResourceListParamsResourceType = "widget"
)

func (r ResourceListParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceListParamsResourceTypeCustomRuleset, ResourceListParamsResourceTypeWidget:
		return true
	}
	return false
}

// Filter share resources by status.
type ResourceListParamsStatus string

const (
	ResourceListParamsStatusActive   ResourceListParamsStatus = "active"
	ResourceListParamsStatusDeleting ResourceListParamsStatus = "deleting"
	ResourceListParamsStatusDeleted  ResourceListParamsStatus = "deleted"
)

func (r ResourceListParamsStatus) IsKnown() bool {
	switch r {
	case ResourceListParamsStatusActive, ResourceListParamsStatusDeleting, ResourceListParamsStatusDeleted:
		return true
	}
	return false
}

type ResourceDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                               `json:"success,required"`
	Result  ResourceDeleteResponse             `json:"result"`
	JSON    resourceDeleteResponseEnvelopeJSON `json:"-"`
}

// resourceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceDeleteResponseEnvelope]
type resourceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  ResourceGetResponse             `json:"result"`
	JSON    resourceGetResponseEnvelopeJSON `json:"-"`
}

// resourceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceGetResponseEnvelope]
type resourceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
