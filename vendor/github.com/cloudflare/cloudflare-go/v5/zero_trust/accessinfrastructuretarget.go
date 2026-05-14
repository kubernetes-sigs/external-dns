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

// AccessInfrastructureTargetService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessInfrastructureTargetService] method instead.
type AccessInfrastructureTargetService struct {
	Options []option.RequestOption
}

// NewAccessInfrastructureTargetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessInfrastructureTargetService(opts ...option.RequestOption) (r *AccessInfrastructureTargetService) {
	r = &AccessInfrastructureTargetService{}
	r.Options = opts
	return
}

// Create new target
func (r *AccessInfrastructureTargetService) New(ctx context.Context, params AccessInfrastructureTargetNewParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetNewResponse, err error) {
	var env AccessInfrastructureTargetNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update target
func (r *AccessInfrastructureTargetService) Update(ctx context.Context, targetID string, params AccessInfrastructureTargetUpdateParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetUpdateResponse, err error) {
	var env AccessInfrastructureTargetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", params.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and sorts an account’s targets. Filters are optional and are ANDed
// together.
func (r *AccessInfrastructureTargetService) List(ctx context.Context, params AccessInfrastructureTargetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessInfrastructureTargetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", params.AccountID)
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

// Lists and sorts an account’s targets. Filters are optional and are ANDed
// together.
func (r *AccessInfrastructureTargetService) ListAutoPaging(ctx context.Context, params AccessInfrastructureTargetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessInfrastructureTargetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete target
func (r *AccessInfrastructureTargetService) Delete(ctx context.Context, targetID string, body AccessInfrastructureTargetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", body.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Removes one or more targets.
//
// Deprecated: deprecated
func (r *AccessInfrastructureTargetService) BulkDelete(ctx context.Context, body AccessInfrastructureTargetBulkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Removes one or more targets.
func (r *AccessInfrastructureTargetService) BulkDeleteV2(ctx context.Context, params AccessInfrastructureTargetBulkDeleteV2Params, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch_delete", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Adds one or more targets.
func (r *AccessInfrastructureTargetService) BulkUpdate(ctx context.Context, params AccessInfrastructureTargetBulkUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessInfrastructureTargetBulkUpdateResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Adds one or more targets.
func (r *AccessInfrastructureTargetService) BulkUpdateAutoPaging(ctx context.Context, params AccessInfrastructureTargetBulkUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessInfrastructureTargetBulkUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkUpdate(ctx, params, opts...))
}

// Get target
func (r *AccessInfrastructureTargetService) Get(ctx context.Context, targetID string, query AccessInfrastructureTargetGetParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetGetResponse, err error) {
	var env AccessInfrastructureTargetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", query.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessInfrastructureTargetNewResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetNewResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetNewResponseJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetNewResponse]
type accessInfrastructureTargetNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetNewResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetNewResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetNewResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetNewResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetNewResponseIP]
type accessInfrastructureTargetNewResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetNewResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetNewResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseIPIPV4]
type accessInfrastructureTargetNewResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetNewResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetNewResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseIPIPV6]
type accessInfrastructureTargetNewResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetUpdateResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                    `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetUpdateResponseJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetUpdateResponse]
type accessInfrastructureTargetUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetUpdateResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetUpdateResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetUpdateResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetUpdateResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIPJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetUpdateResponseIP]
type accessInfrastructureTargetUpdateResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetUpdateResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                             `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetUpdateResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIpipv4JSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseIPIPV4]
type accessInfrastructureTargetUpdateResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetUpdateResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                             `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetUpdateResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIpipv6JSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseIPIPV6]
type accessInfrastructureTargetUpdateResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetListResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetListResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                  `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetListResponseJSON `json:"-"`
}

// accessInfrastructureTargetListResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetListResponse]
type accessInfrastructureTargetListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetListResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetListResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetListResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetListResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetListResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetListResponseIP]
type accessInfrastructureTargetListResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetListResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                           `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetListResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetListResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetListResponseIPIPV4]
type accessInfrastructureTargetListResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetListResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                           `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetListResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetListResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetListResponseIPIPV6]
type accessInfrastructureTargetListResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetBulkUpdateResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetBulkUpdateResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                        `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetBulkUpdateResponseJSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetBulkUpdateResponse]
type accessInfrastructureTargetBulkUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetBulkUpdateResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetBulkUpdateResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetBulkUpdateResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetBulkUpdateResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIPJSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetBulkUpdateResponseIP]
type accessInfrastructureTargetBulkUpdateResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetBulkUpdateResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                                 `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetBulkUpdateResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIpipv4JSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetBulkUpdateResponseIPIPV4]
type accessInfrastructureTargetBulkUpdateResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetBulkUpdateResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                                 `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetBulkUpdateResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIpipv6JSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetBulkUpdateResponseIPIPV6]
type accessInfrastructureTargetBulkUpdateResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetGetResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetGetResponseJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetGetResponse]
type accessInfrastructureTargetGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetGetResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetGetResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetGetResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetGetResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetGetResponseIP]
type accessInfrastructureTargetGetResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetGetResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetGetResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseIPIPV4]
type accessInfrastructureTargetGetResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetGetResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetGetResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseIPIPV6]
type accessInfrastructureTargetGetResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetNewParamsIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetNewParamsIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetNewParamsIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetNewParamsIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetNewParamsIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetNewParamsIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetNewParamsIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetNewParamsIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetNewParamsIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetNewResponseEnvelope struct {
	Errors   []AccessInfrastructureTargetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessInfrastructureTargetNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessInfrastructureTargetNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetNewResponse                `json:"result"`
	JSON    accessInfrastructureTargetNewResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseEnvelope]
type accessInfrastructureTargetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessInfrastructureTargetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessInfrastructureTargetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetNewResponseEnvelopeErrors]
type accessInfrastructureTargetNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessInfrastructureTargetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessInfrastructureTargetNewResponseEnvelopeErrorsSource]
type accessInfrastructureTargetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessInfrastructureTargetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessInfrastructureTargetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetNewResponseEnvelopeMessages]
type accessInfrastructureTargetNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessInfrastructureTargetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessInfrastructureTargetNewResponseEnvelopeMessagesSource]
type accessInfrastructureTargetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessInfrastructureTargetNewResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetNewResponseEnvelopeSuccessTrue AccessInfrastructureTargetNewResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessInfrastructureTargetUpdateParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetUpdateParamsIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetUpdateParamsIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetUpdateParamsIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetUpdateParamsIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetUpdateParamsIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetUpdateParamsIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetUpdateParamsIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetUpdateParamsIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetUpdateParamsIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetUpdateResponseEnvelope struct {
	Errors   []AccessInfrastructureTargetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessInfrastructureTargetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessInfrastructureTargetUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetUpdateResponse                `json:"result"`
	JSON    accessInfrastructureTargetUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseEnvelope]
type accessInfrastructureTargetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessInfrastructureTargetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessInfrastructureTargetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetUpdateResponseEnvelopeErrors]
type accessInfrastructureTargetUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessInfrastructureTargetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccessInfrastructureTargetUpdateResponseEnvelopeErrorsSource]
type accessInfrastructureTargetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AccessInfrastructureTargetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessInfrastructureTargetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AccessInfrastructureTargetUpdateResponseEnvelopeMessages]
type accessInfrastructureTargetUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    accessInfrastructureTargetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessInfrastructureTargetUpdateResponseEnvelopeMessagesSource]
type accessInfrastructureTargetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessInfrastructureTargetUpdateResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetUpdateResponseEnvelopeSuccessTrue AccessInfrastructureTargetUpdateResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessInfrastructureTargetListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Date and time at which the target was created after (inclusive)
	CreatedAfter param.Field[time.Time] `query:"created_after" format:"date-time"`
	// Date and time at which the target was created before (inclusive)
	CreatedBefore param.Field[time.Time] `query:"created_before" format:"date-time"`
	// The sorting direction.
	Direction param.Field[AccessInfrastructureTargetListParamsDirection] `query:"direction"`
	// Hostname of a target
	Hostname param.Field[string] `query:"hostname"`
	// Partial match to the hostname of a target
	HostnameContains param.Field[string] `query:"hostname_contains"`
	// Filters for targets whose IP addresses look like the specified string. Supports
	// `*` as a wildcard character
	IPLike param.Field[string] `query:"ip_like"`
	// IPv4 address of the target
	IPV4 param.Field[string] `query:"ip_v4"`
	// IPv6 address of the target
	IPV6 param.Field[string] `query:"ip_v6"`
	// Filters for targets that have any of the following IP addresses. Specify `ips`
	// multiple times in query parameter to build list of candidates.
	IPs param.Field[[]string] `query:"ips"`
	// Defines an IPv4 filter range's ending value (inclusive). Requires `ipv4_start`
	// to be specified as well.
	IPV4End param.Field[string] `query:"ipv4_end"`
	// Defines an IPv4 filter range's starting value (inclusive). Requires `ipv4_end`
	// to be specified as well.
	IPV4Start param.Field[string] `query:"ipv4_start"`
	// Defines an IPv6 filter range's ending value (inclusive). Requires `ipv6_start`
	// to be specified as well.
	IPV6End param.Field[string] `query:"ipv6_end"`
	// Defines an IPv6 filter range's starting value (inclusive). Requires `ipv6_end`
	// to be specified as well.
	IPV6Start param.Field[string] `query:"ipv6_start"`
	// Date and time at which the target was modified after (inclusive)
	ModifiedAfter param.Field[time.Time] `query:"modified_after" format:"date-time"`
	// Date and time at which the target was modified before (inclusive)
	ModifiedBefore param.Field[time.Time] `query:"modified_before" format:"date-time"`
	// The field to sort by.
	Order param.Field[AccessInfrastructureTargetListParamsOrder] `query:"order"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64] `query:"per_page"`
	// Filters for targets that have any of the following UUIDs. Specify `target_ids`
	// multiple times in query parameter to build list of candidates.
	TargetIDs param.Field[[]string] `query:"target_ids" format:"uuid"`
	// Private virtual network identifier of the target
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccessInfrastructureTargetListParams]'s query parameters as
// `url.Values`.
func (r AccessInfrastructureTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type AccessInfrastructureTargetListParamsDirection string

const (
	AccessInfrastructureTargetListParamsDirectionAsc  AccessInfrastructureTargetListParamsDirection = "asc"
	AccessInfrastructureTargetListParamsDirectionDesc AccessInfrastructureTargetListParamsDirection = "desc"
)

func (r AccessInfrastructureTargetListParamsDirection) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetListParamsDirectionAsc, AccessInfrastructureTargetListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type AccessInfrastructureTargetListParamsOrder string

const (
	AccessInfrastructureTargetListParamsOrderHostname  AccessInfrastructureTargetListParamsOrder = "hostname"
	AccessInfrastructureTargetListParamsOrderCreatedAt AccessInfrastructureTargetListParamsOrder = "created_at"
)

func (r AccessInfrastructureTargetListParamsOrder) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetListParamsOrderHostname, AccessInfrastructureTargetListParamsOrderCreatedAt:
		return true
	}
	return false
}

type AccessInfrastructureTargetDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetBulkDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetBulkDeleteV2Params struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// List of target IDs to bulk delete
	TargetIDs param.Field[[]string] `json:"target_ids,required" format:"uuid"`
}

func (r AccessInfrastructureTargetBulkDeleteV2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetBulkUpdateParams struct {
	// Account identifier
	AccountID param.Field[string]                              `path:"account_id,required"`
	Body      []AccessInfrastructureTargetBulkUpdateParamsBody `json:"body,required"`
}

func (r AccessInfrastructureTargetBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessInfrastructureTargetBulkUpdateParamsBody struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetBulkUpdateParamsBodyIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetGetResponseEnvelope struct {
	Errors   []AccessInfrastructureTargetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessInfrastructureTargetGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessInfrastructureTargetGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetGetResponse                `json:"result"`
	JSON    accessInfrastructureTargetGetResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseEnvelope]
type accessInfrastructureTargetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessInfrastructureTargetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessInfrastructureTargetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetGetResponseEnvelopeErrors]
type accessInfrastructureTargetGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessInfrastructureTargetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessInfrastructureTargetGetResponseEnvelopeErrorsSource]
type accessInfrastructureTargetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessInfrastructureTargetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessInfrastructureTargetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetGetResponseEnvelopeMessages]
type accessInfrastructureTargetGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessInfrastructureTargetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessInfrastructureTargetGetResponseEnvelopeMessagesSource]
type accessInfrastructureTargetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessInfrastructureTargetGetResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetGetResponseEnvelopeSuccessTrue AccessInfrastructureTargetGetResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
