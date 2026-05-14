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

// AccessLogSCIMUpdateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogSCIMUpdateService] method instead.
type AccessLogSCIMUpdateService struct {
	Options []option.RequestOption
}

// NewAccessLogSCIMUpdateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogSCIMUpdateService(opts ...option.RequestOption) (r *AccessLogSCIMUpdateService) {
	r = &AccessLogSCIMUpdateService{}
	r.Options = opts
	return
}

// Lists Access SCIM update logs that maintain a record of updates made to User and
// Group resources synced to Cloudflare via the System for Cross-domain Identity
// Management (SCIM).
func (r *AccessLogSCIMUpdateService) List(ctx context.Context, params AccessLogSCIMUpdateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessLogSCIMUpdateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/logs/scim/updates", params.AccountID)
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

// Lists Access SCIM update logs that maintain a record of updates made to User and
// Group resources synced to Cloudflare via the System for Cross-domain Identity
// Management (SCIM).
func (r *AccessLogSCIMUpdateService) ListAutoPaging(ctx context.Context, params AccessLogSCIMUpdateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessLogSCIMUpdateListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type AccessLogSCIMUpdateListResponse struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	CfResourceID string `json:"cf_resource_id"`
	// The error message which is generated when the status of the SCIM request is
	// 'FAILURE'.
	ErrorDescription string `json:"error_description"`
	// The unique Id of the IdP that has SCIM enabled.
	IdPID string `json:"idp_id"`
	// The IdP-generated Id of the SCIM resource.
	IdPResourceID string    `json:"idp_resource_id"`
	LoggedAt      time.Time `json:"logged_at" format:"date-time"`
	// The JSON-encoded string body of the SCIM request.
	RequestBody string `json:"request_body"`
	// The request method of the SCIM request.
	RequestMethod string `json:"request_method"`
	// The display name of the SCIM Group resource if it exists.
	ResourceGroupName string `json:"resource_group_name"`
	// The resource type of the SCIM request.
	ResourceType string `json:"resource_type"`
	// The email address of the SCIM User resource if it exists.
	ResourceUserEmail string `json:"resource_user_email" format:"email"`
	// The status of the SCIM request.
	Status string                              `json:"status"`
	JSON   accessLogSCIMUpdateListResponseJSON `json:"-"`
}

// accessLogSCIMUpdateListResponseJSON contains the JSON metadata for the struct
// [AccessLogSCIMUpdateListResponse]
type accessLogSCIMUpdateListResponseJSON struct {
	CfResourceID      apijson.Field
	ErrorDescription  apijson.Field
	IdPID             apijson.Field
	IdPResourceID     apijson.Field
	LoggedAt          apijson.Field
	RequestBody       apijson.Field
	RequestMethod     apijson.Field
	ResourceGroupName apijson.Field
	ResourceType      apijson.Field
	ResourceUserEmail apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessLogSCIMUpdateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogSCIMUpdateListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessLogSCIMUpdateListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique Id of the IdP that has SCIM enabled.
	IdPID param.Field[[]string] `query:"idp_id,required"`
	// The unique Cloudflare-generated Id of the SCIM resource.
	CfResourceID param.Field[string] `query:"cf_resource_id"`
	// The chronological order used to sort the logs.
	Direction param.Field[AccessLogSCIMUpdateListParamsDirection] `query:"direction"`
	// The IdP-generated Id of the SCIM resource.
	IdPResourceID param.Field[string] `query:"idp_resource_id"`
	// The maximum number of update logs to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The request method of the SCIM request.
	RequestMethod param.Field[[]AccessLogSCIMUpdateListParamsRequestMethod] `query:"request_method"`
	// The display name of the SCIM Group resource.
	ResourceGroupName param.Field[string] `query:"resource_group_name"`
	// The resource type of the SCIM request.
	ResourceType param.Field[[]AccessLogSCIMUpdateListParamsResourceType] `query:"resource_type"`
	// The email address of the SCIM User resource.
	ResourceUserEmail param.Field[string] `query:"resource_user_email" format:"email"`
	// the timestamp of the earliest update log.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The status of the SCIM request.
	Status param.Field[[]AccessLogSCIMUpdateListParamsStatus] `query:"status"`
	// the timestamp of the most-recent update log.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccessLogSCIMUpdateListParams]'s query parameters as
// `url.Values`.
func (r AccessLogSCIMUpdateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The chronological order used to sort the logs.
type AccessLogSCIMUpdateListParamsDirection string

const (
	AccessLogSCIMUpdateListParamsDirectionDesc AccessLogSCIMUpdateListParamsDirection = "desc"
	AccessLogSCIMUpdateListParamsDirectionAsc  AccessLogSCIMUpdateListParamsDirection = "asc"
)

func (r AccessLogSCIMUpdateListParamsDirection) IsKnown() bool {
	switch r {
	case AccessLogSCIMUpdateListParamsDirectionDesc, AccessLogSCIMUpdateListParamsDirectionAsc:
		return true
	}
	return false
}

type AccessLogSCIMUpdateListParamsRequestMethod string

const (
	AccessLogSCIMUpdateListParamsRequestMethodDelete AccessLogSCIMUpdateListParamsRequestMethod = "DELETE"
	AccessLogSCIMUpdateListParamsRequestMethodPatch  AccessLogSCIMUpdateListParamsRequestMethod = "PATCH"
	AccessLogSCIMUpdateListParamsRequestMethodPost   AccessLogSCIMUpdateListParamsRequestMethod = "POST"
	AccessLogSCIMUpdateListParamsRequestMethodPut    AccessLogSCIMUpdateListParamsRequestMethod = "PUT"
)

func (r AccessLogSCIMUpdateListParamsRequestMethod) IsKnown() bool {
	switch r {
	case AccessLogSCIMUpdateListParamsRequestMethodDelete, AccessLogSCIMUpdateListParamsRequestMethodPatch, AccessLogSCIMUpdateListParamsRequestMethodPost, AccessLogSCIMUpdateListParamsRequestMethodPut:
		return true
	}
	return false
}

type AccessLogSCIMUpdateListParamsResourceType string

const (
	AccessLogSCIMUpdateListParamsResourceTypeUser  AccessLogSCIMUpdateListParamsResourceType = "USER"
	AccessLogSCIMUpdateListParamsResourceTypeGroup AccessLogSCIMUpdateListParamsResourceType = "GROUP"
)

func (r AccessLogSCIMUpdateListParamsResourceType) IsKnown() bool {
	switch r {
	case AccessLogSCIMUpdateListParamsResourceTypeUser, AccessLogSCIMUpdateListParamsResourceTypeGroup:
		return true
	}
	return false
}

type AccessLogSCIMUpdateListParamsStatus string

const (
	AccessLogSCIMUpdateListParamsStatusFailure AccessLogSCIMUpdateListParamsStatus = "FAILURE"
	AccessLogSCIMUpdateListParamsStatusSuccess AccessLogSCIMUpdateListParamsStatus = "SUCCESS"
)

func (r AccessLogSCIMUpdateListParamsStatus) IsKnown() bool {
	switch r {
	case AccessLogSCIMUpdateListParamsStatusFailure, AccessLogSCIMUpdateListParamsStatusSuccess:
		return true
	}
	return false
}
