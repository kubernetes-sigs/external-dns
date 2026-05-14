// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// AccessGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessGroupService] method instead.
type AccessGroupService struct {
	Options []option.RequestOption
}

// NewAccessGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessGroupService(opts ...option.RequestOption) (r *AccessGroupService) {
	r = &AccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *AccessGroupService) New(ctx context.Context, params AccessGroupNewParams, opts ...option.RequestOption) (res *AccessGroupNewResponse, err error) {
	var env AccessGroupNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Access group.
func (r *AccessGroupService) Update(ctx context.Context, groupID string, params AccessGroupUpdateParams, opts ...option.RequestOption) (res *AccessGroupUpdateResponse, err error) {
	var env AccessGroupUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *AccessGroupService) List(ctx context.Context, params AccessGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
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

// Lists all Access groups.
func (r *AccessGroupService) ListAutoPaging(ctx context.Context, params AccessGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an Access group.
func (r *AccessGroupService) Delete(ctx context.Context, groupID string, body AccessGroupDeleteParams, opts ...option.RequestOption) (res *AccessGroupDeleteResponse, err error) {
	var env AccessGroupDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access group.
func (r *AccessGroupService) Get(ctx context.Context, groupID string, query AccessGroupGetParams, opts ...option.RequestOption) (res *AccessGroupGetResponse, err error) {
	var env AccessGroupGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGroup struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	ID string `json:"id"`
	// The display name of the SCIM Group resource.
	DisplayName string `json:"displayName"`
	// The IdP-generated Id of the SCIM resource.
	ExternalID string `json:"externalId"`
	// The metadata of the SCIM resource.
	Meta ZeroTrustGroupMeta `json:"meta"`
	// The list of URIs which indicate the attributes contained within a SCIM resource.
	Schemas []string           `json:"schemas"`
	JSON    zeroTrustGroupJSON `json:"-"`
}

// zeroTrustGroupJSON contains the JSON metadata for the struct [ZeroTrustGroup]
type zeroTrustGroupJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Schemas     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupJSON) RawJSON() string {
	return r.raw
}

// The metadata of the SCIM resource.
type ZeroTrustGroupMeta struct {
	// The timestamp of when the SCIM resource was created.
	Created time.Time `json:"created" format:"date-time"`
	// The timestamp of when the SCIM resource was last modified.
	LastModified time.Time              `json:"lastModified" format:"date-time"`
	JSON         zeroTrustGroupMetaJSON `json:"-"`
}

// zeroTrustGroupMetaJSON contains the JSON metadata for the struct
// [ZeroTrustGroupMeta]
type zeroTrustGroupMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupMetaJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewResponse struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessRule `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule               `json:"require"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessGroupNewResponseJSON `json:"-"`
}

// accessGroupNewResponseJSON contains the JSON metadata for the struct
// [AccessGroupNewResponse]
type accessGroupNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupUpdateResponse struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessRule `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                  `json:"require"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      accessGroupUpdateResponseJSON `json:"-"`
}

// accessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponse]
type accessGroupUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupListResponse struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessRule `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule                `json:"require"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      accessGroupListResponseJSON `json:"-"`
}

// accessGroupListResponseJSON contains the JSON metadata for the struct
// [AccessGroupListResponse]
type accessGroupListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponse struct {
	// UUID.
	ID   string                        `json:"id"`
	JSON accessGroupDeleteResponseJSON `json:"-"`
}

// accessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponse]
type accessGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupGetResponse struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessRule `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule               `json:"require"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessGroupGetResponseJSON `json:"-"`
}

// accessGroupGetResponseJSON contains the JSON metadata for the struct
// [AccessGroupGetResponse]
type accessGroupGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupNewResponseEnvelope struct {
	Errors   []AccessGroupNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGroupNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGroupNewResponse                `json:"result"`
	JSON    accessGroupNewResponseEnvelopeJSON    `json:"-"`
}

// accessGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupNewResponseEnvelope]
type accessGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccessGroupNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGroupNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseEnvelopeErrors]
type accessGroupNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accessGroupNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGroupNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseEnvelopeErrorsSource]
type accessGroupNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccessGroupNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGroupNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseEnvelopeMessages]
type accessGroupNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accessGroupNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGroupNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseEnvelopeMessagesSource]
type accessGroupNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGroupNewResponseEnvelopeSuccess bool

const (
	AccessGroupNewResponseEnvelopeSuccessTrue AccessGroupNewResponseEnvelopeSuccess = true
)

func (r AccessGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupUpdateResponseEnvelope struct {
	Errors   []AccessGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGroupUpdateResponse                `json:"result"`
	JSON    accessGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponseEnvelope]
type accessGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGroupUpdateResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessGroupUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGroupUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupUpdateResponseEnvelopeErrors]
type accessGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGroupUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessGroupUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGroupUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessGroupUpdateResponseEnvelopeErrorsSource]
type accessGroupUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGroupUpdateResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccessGroupUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGroupUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGroupUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupUpdateResponseEnvelopeMessages]
type accessGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGroupUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accessGroupUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGroupUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseEnvelopeMessagesSource]
type accessGroupUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGroupUpdateResponseEnvelopeSuccess bool

const (
	AccessGroupUpdateResponseEnvelopeSuccessTrue AccessGroupUpdateResponseEnvelopeSuccess = true
)

func (r AccessGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The name of the group.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Search for groups by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessGroupListParams]'s query parameters as `url.Values`.
func (r AccessGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessGroupDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupDeleteResponseEnvelope struct {
	Errors   []AccessGroupDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGroupDeleteResponse                `json:"result"`
	JSON    accessGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponseEnvelope]
type accessGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessGroupDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGroupDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupDeleteResponseEnvelopeErrors]
type accessGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessGroupDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGroupDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessGroupDeleteResponseEnvelopeErrorsSource]
type accessGroupDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccessGroupDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGroupDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGroupDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupDeleteResponseEnvelopeMessages]
type accessGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accessGroupDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGroupDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessGroupDeleteResponseEnvelopeMessagesSource]
type accessGroupDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGroupDeleteResponseEnvelopeSuccess bool

const (
	AccessGroupDeleteResponseEnvelopeSuccessTrue AccessGroupDeleteResponseEnvelopeSuccess = true
)

func (r AccessGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupGetResponseEnvelope struct {
	Errors   []AccessGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGroupGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGroupGetResponse                `json:"result"`
	JSON    accessGroupGetResponseEnvelopeJSON    `json:"-"`
}

// accessGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupGetResponseEnvelope]
type accessGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGroupGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccessGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseEnvelopeErrors]
type accessGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accessGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseEnvelopeErrorsSource]
type accessGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGroupGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccessGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseEnvelopeMessages]
type accessGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accessGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseEnvelopeMessagesSource]
type accessGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGroupGetResponseEnvelopeSuccess bool

const (
	AccessGroupGetResponseEnvelopeSuccessTrue AccessGroupGetResponseEnvelopeSuccess = true
)

func (r AccessGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
