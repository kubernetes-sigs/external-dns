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

// NetworkSubnetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetService] method instead.
type NetworkSubnetService struct {
	Options          []option.RequestOption
	CloudflareSource *NetworkSubnetCloudflareSourceService
}

// NewNetworkSubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkSubnetService(opts ...option.RequestOption) (r *NetworkSubnetService) {
	r = &NetworkSubnetService{}
	r.Options = opts
	r.CloudflareSource = NewNetworkSubnetCloudflareSourceService(opts...)
	return
}

// Lists and filters subnets in an account.
func (r *NetworkSubnetService) List(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[NetworkSubnetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets", params.AccountID)
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

// Lists and filters subnets in an account.
func (r *NetworkSubnetService) ListAutoPaging(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NetworkSubnetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type NetworkSubnetListResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetListResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetListResponseJSON       `json:"-"`
}

// networkSubnetListResponseJSON contains the JSON metadata for the struct
// [NetworkSubnetListResponse]
type networkSubnetListResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetListResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetListResponseSubnetType string

const (
	NetworkSubnetListResponseSubnetTypeCloudflareSource NetworkSubnetListResponseSubnetType = "cloudflare_source"
)

func (r NetworkSubnetListResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetListResponseSubnetTypeCloudflareSource:
		return true
	}
	return false
}

type NetworkSubnetListParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// If set, only include subnets in the given address family - `v4` or `v6`
	AddressFamily param.Field[NetworkSubnetListParamsAddressFamily] `query:"address_family"`
	// If set, only list subnets with the given comment.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	// If `true`, only include default subnets. If `false`, exclude default subnets
	// subnets. If not set, all subnets will be included.
	IsDefaultNetwork param.Field[bool] `query:"is_default_network"`
	// If `true`, only include deleted subnets. If `false`, exclude deleted subnets. If
	// not set, all subnets will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// If set, only list subnets with the given name
	Name param.Field[string] `query:"name"`
	// If set, only list the subnet whose network exactly matches the given CIDR.
	Network param.Field[string] `query:"network"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// Sort order of the results. `asc` means oldest to newest, `desc` means newest to
	// oldest. If not set, they will not be in any particular order.
	SortOrder param.Field[NetworkSubnetListParamsSortOrder] `query:"sort_order"`
	// If set, the types of subnets to include, separated by comma.
	SubnetTypes param.Field[NetworkSubnetListParamsSubnetTypes] `query:"subnet_types"`
}

// URLQuery serializes [NetworkSubnetListParams]'s query parameters as
// `url.Values`.
func (r NetworkSubnetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// If set, only include subnets in the given address family - `v4` or `v6`
type NetworkSubnetListParamsAddressFamily string

const (
	NetworkSubnetListParamsAddressFamilyV4 NetworkSubnetListParamsAddressFamily = "v4"
	NetworkSubnetListParamsAddressFamilyV6 NetworkSubnetListParamsAddressFamily = "v6"
)

func (r NetworkSubnetListParamsAddressFamily) IsKnown() bool {
	switch r {
	case NetworkSubnetListParamsAddressFamilyV4, NetworkSubnetListParamsAddressFamilyV6:
		return true
	}
	return false
}

// Sort order of the results. `asc` means oldest to newest, `desc` means newest to
// oldest. If not set, they will not be in any particular order.
type NetworkSubnetListParamsSortOrder string

const (
	NetworkSubnetListParamsSortOrderAsc  NetworkSubnetListParamsSortOrder = "asc"
	NetworkSubnetListParamsSortOrderDesc NetworkSubnetListParamsSortOrder = "desc"
)

func (r NetworkSubnetListParamsSortOrder) IsKnown() bool {
	switch r {
	case NetworkSubnetListParamsSortOrderAsc, NetworkSubnetListParamsSortOrderDesc:
		return true
	}
	return false
}

// If set, the types of subnets to include, separated by comma.
type NetworkSubnetListParamsSubnetTypes string

const (
	NetworkSubnetListParamsSubnetTypesCloudflareSource NetworkSubnetListParamsSubnetTypes = "cloudflare_source"
	NetworkSubnetListParamsSubnetTypesWARP             NetworkSubnetListParamsSubnetTypes = "warp"
)

func (r NetworkSubnetListParamsSubnetTypes) IsKnown() bool {
	switch r {
	case NetworkSubnetListParamsSubnetTypesCloudflareSource, NetworkSubnetListParamsSubnetTypesWARP:
		return true
	}
	return false
}
