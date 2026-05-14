// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DevicePolicyDefaultExcludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultExcludeService] method instead.
type DevicePolicyDefaultExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultExcludeService(opts ...option.RequestOption) (r *DevicePolicyDefaultExcludeService) {
	r = &DevicePolicyDefaultExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) Update(ctx context.Context, params DevicePolicyDefaultExcludeUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelExclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", params.AccountID)
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

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) UpdateAutoPaging(ctx context.Context, params DevicePolicyDefaultExcludeUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelExclude] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) Get(ctx context.Context, query DevicePolicyDefaultExcludeGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelExclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) GetAutoPaging(ctx context.Context, query DevicePolicyDefaultExcludeGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelExclude] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type DevicePolicyDefaultExcludeUpdateParams struct {
	AccountID param.Field[string]            `path:"account_id,required"`
	Body      []SplitTunnelExcludeUnionParam `json:"body,required"`
}

func (r DevicePolicyDefaultExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyDefaultExcludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
