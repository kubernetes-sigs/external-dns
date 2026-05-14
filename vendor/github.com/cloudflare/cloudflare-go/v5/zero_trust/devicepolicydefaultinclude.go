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

// DevicePolicyDefaultIncludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultIncludeService] method instead.
type DevicePolicyDefaultIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultIncludeService(opts ...option.RequestOption) (r *DevicePolicyDefaultIncludeService) {
	r = &DevicePolicyDefaultIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) Update(ctx context.Context, params DevicePolicyDefaultIncludeUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelInclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", params.AccountID)
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

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) UpdateAutoPaging(ctx context.Context, params DevicePolicyDefaultIncludeUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelInclude] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) Get(ctx context.Context, query DevicePolicyDefaultIncludeGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelInclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", query.AccountID)
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

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) GetAutoPaging(ctx context.Context, query DevicePolicyDefaultIncludeGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelInclude] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type DevicePolicyDefaultIncludeUpdateParams struct {
	AccountID param.Field[string]            `path:"account_id,required"`
	Body      []SplitTunnelIncludeUnionParam `json:"body,required"`
}

func (r DevicePolicyDefaultIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyDefaultIncludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
