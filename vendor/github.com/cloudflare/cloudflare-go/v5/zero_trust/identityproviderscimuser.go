// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// IdentityProviderSCIMUserService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderSCIMUserService] method instead.
type IdentityProviderSCIMUserService struct {
	Options []option.RequestOption
}

// NewIdentityProviderSCIMUserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityProviderSCIMUserService(opts ...option.RequestOption) (r *IdentityProviderSCIMUserService) {
	r = &IdentityProviderSCIMUserService{}
	r.Options = opts
	return
}

// Lists SCIM User resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMUserService) List(ctx context.Context, identityProviderID string, params IdentityProviderSCIMUserListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessUser], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/users", params.AccountID, identityProviderID)
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

// Lists SCIM User resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMUserService) ListAutoPaging(ctx context.Context, identityProviderID string, params IdentityProviderSCIMUserListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessUser] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, identityProviderID, params, opts...))
}

type IdentityProviderSCIMUserListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique Cloudflare-generated Id of the SCIM User resource; also known as the
	// "Id".
	CfResourceID param.Field[string] `query:"cf_resource_id"`
	// The email address of the SCIM User resource.
	Email param.Field[string] `query:"email"`
	// The IdP-generated Id of the SCIM User resource; also known as the "external Id".
	IdPResourceID param.Field[string] `query:"idp_resource_id"`
	// The name of the SCIM User resource.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The username of the SCIM User resource.
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [IdentityProviderSCIMUserListParams]'s query parameters as
// `url.Values`.
func (r IdentityProviderSCIMUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
