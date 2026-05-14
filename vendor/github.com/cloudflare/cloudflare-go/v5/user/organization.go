// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/accounts"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Lists organizations the user is associated with.
//
// Deprecated: deprecated
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Organization], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/organizations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Lists organizations the user is associated with.
//
// Deprecated: deprecated
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Organization] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Removes association to an organization.
//
// Deprecated: deprecated
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets a specific organization the user is associated with.
//
// Deprecated: deprecated
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	var env OrganizationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Organization struct {
	// Identifier
	ID string `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Access permissions for this User.
	Permissions []shared.Permission `json:"permissions"`
	// List of roles that a user has within an organization.
	Roles []string `json:"roles"`
	// Whether the user is a member of the organization or has an invitation pending.
	Status accounts.Status  `json:"status"`
	JSON   organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponse struct {
	// Identifier
	ID   string                         `json:"id"`
	JSON organizationDeleteResponseJSON `json:"-"`
}

// organizationDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponse]
type organizationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponse = interface{}

type OrganizationListParams struct {
	// Direction to order organizations.
	Direction param.Field[OrganizationListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[OrganizationListParamsMatch] `query:"match"`
	// Organization name.
	Name param.Field[string] `query:"name"`
	// Field to order organizations by.
	Order param.Field[OrganizationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of organizations per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status param.Field[OrganizationListParamsStatus] `query:"status"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order organizations.
type OrganizationListParamsDirection string

const (
	OrganizationListParamsDirectionAsc  OrganizationListParamsDirection = "asc"
	OrganizationListParamsDirectionDesc OrganizationListParamsDirection = "desc"
)

func (r OrganizationListParamsDirection) IsKnown() bool {
	switch r {
	case OrganizationListParamsDirectionAsc, OrganizationListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any).
type OrganizationListParamsMatch string

const (
	OrganizationListParamsMatchAny OrganizationListParamsMatch = "any"
	OrganizationListParamsMatchAll OrganizationListParamsMatch = "all"
)

func (r OrganizationListParamsMatch) IsKnown() bool {
	switch r {
	case OrganizationListParamsMatchAny, OrganizationListParamsMatchAll:
		return true
	}
	return false
}

// Field to order organizations by.
type OrganizationListParamsOrder string

const (
	OrganizationListParamsOrderID     OrganizationListParamsOrder = "id"
	OrganizationListParamsOrderName   OrganizationListParamsOrder = "name"
	OrganizationListParamsOrderStatus OrganizationListParamsOrder = "status"
)

func (r OrganizationListParamsOrder) IsKnown() bool {
	switch r {
	case OrganizationListParamsOrderID, OrganizationListParamsOrderName, OrganizationListParamsOrderStatus:
		return true
	}
	return false
}

// Whether the user is a member of the organization or has an inivitation pending.
type OrganizationListParamsStatus string

const (
	OrganizationListParamsStatusMember  OrganizationListParamsStatus = "member"
	OrganizationListParamsStatusInvited OrganizationListParamsStatus = "invited"
)

func (r OrganizationListParamsStatus) IsKnown() bool {
	switch r {
	case OrganizationListParamsStatusMember, OrganizationListParamsStatusInvited:
		return true
	}
	return false
}

type OrganizationGetResponseEnvelope struct {
	Errors   []OrganizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OrganizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OrganizationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OrganizationGetResponse                `json:"result"`
	JSON    organizationGetResponseEnvelopeJSON    `json:"-"`
}

// organizationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationGetResponseEnvelope]
type organizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           OrganizationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationGetResponseEnvelopeErrors]
type organizationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    organizationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OrganizationGetResponseEnvelopeErrorsSource]
type organizationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OrganizationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationGetResponseEnvelopeMessages]
type organizationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    organizationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OrganizationGetResponseEnvelopeMessagesSource]
type organizationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationGetResponseEnvelopeSuccess bool

const (
	OrganizationGetResponseEnvelopeSuccessTrue OrganizationGetResponseEnvelopeSuccess = true
)

func (r OrganizationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
