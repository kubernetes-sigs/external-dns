// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AccessApplicationPolicyTestUserService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyTestUserService] method instead.
type AccessApplicationPolicyTestUserService struct {
	Options []option.RequestOption
}

// NewAccessApplicationPolicyTestUserService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyTestUserService(opts ...option.RequestOption) (r *AccessApplicationPolicyTestUserService) {
	r = &AccessApplicationPolicyTestUserService{}
	r.Options = opts
	return
}

// Fetches a single page of user results from an Access policy test.
func (r *AccessApplicationPolicyTestUserService) List(ctx context.Context, policyTestID string, params AccessApplicationPolicyTestUserListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessApplicationPolicyTestUserListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s/users", params.AccountID, policyTestID)
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

// Fetches a single page of user results from an Access policy test.
func (r *AccessApplicationPolicyTestUserService) ListAutoPaging(ctx context.Context, policyTestID string, params AccessApplicationPolicyTestUserListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessApplicationPolicyTestUserListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, policyTestID, params, opts...))
}

type AccessApplicationPolicyTestUserListResponse struct {
	// UUID.
	ID string `json:"id"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// The name of the user.
	Name string `json:"name"`
	// Policy evaluation result for an individual user.
	Status AccessApplicationPolicyTestUserListResponseStatus `json:"status"`
	JSON   accessApplicationPolicyTestUserListResponseJSON   `json:"-"`
}

// accessApplicationPolicyTestUserListResponseJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyTestUserListResponse]
type accessApplicationPolicyTestUserListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestUserListResponseJSON) RawJSON() string {
	return r.raw
}

// Policy evaluation result for an individual user.
type AccessApplicationPolicyTestUserListResponseStatus string

const (
	AccessApplicationPolicyTestUserListResponseStatusApproved AccessApplicationPolicyTestUserListResponseStatus = "approved"
	AccessApplicationPolicyTestUserListResponseStatusBlocked  AccessApplicationPolicyTestUserListResponseStatus = "blocked"
	AccessApplicationPolicyTestUserListResponseStatusError    AccessApplicationPolicyTestUserListResponseStatus = "error"
)

func (r AccessApplicationPolicyTestUserListResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestUserListResponseStatusApproved, AccessApplicationPolicyTestUserListResponseStatusBlocked, AccessApplicationPolicyTestUserListResponseStatusError:
		return true
	}
	return false
}

type AccessApplicationPolicyTestUserListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of results.
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// Filter users by their policy evaluation status.
	Status param.Field[AccessApplicationPolicyTestUserListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccessApplicationPolicyTestUserListParams]'s query
// parameters as `url.Values`.
func (r AccessApplicationPolicyTestUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter users by their policy evaluation status.
type AccessApplicationPolicyTestUserListParamsStatus string

const (
	AccessApplicationPolicyTestUserListParamsStatusSuccess AccessApplicationPolicyTestUserListParamsStatus = "success"
	AccessApplicationPolicyTestUserListParamsStatusFail    AccessApplicationPolicyTestUserListParamsStatus = "fail"
	AccessApplicationPolicyTestUserListParamsStatusError   AccessApplicationPolicyTestUserListParamsStatus = "error"
)

func (r AccessApplicationPolicyTestUserListParamsStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestUserListParamsStatusSuccess, AccessApplicationPolicyTestUserListParamsStatusFail, AccessApplicationPolicyTestUserListParamsStatusError:
		return true
	}
	return false
}
