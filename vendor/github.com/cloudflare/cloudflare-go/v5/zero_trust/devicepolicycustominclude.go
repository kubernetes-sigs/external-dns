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

// DevicePolicyCustomIncludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomIncludeService] method instead.
type DevicePolicyCustomIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCustomIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomIncludeService(opts ...option.RequestOption) (r *DevicePolicyCustomIncludeService) {
	r = &DevicePolicyCustomIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) Update(ctx context.Context, policyID string, params DevicePolicyCustomIncludeUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelInclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", params.AccountID, policyID)
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

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) UpdateAutoPaging(ctx context.Context, policyID string, params DevicePolicyCustomIncludeUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelInclude] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, policyID, params, opts...))
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) Get(ctx context.Context, policyID string, query DevicePolicyCustomIncludeGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelInclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", query.AccountID, policyID)
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

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) GetAutoPaging(ctx context.Context, policyID string, query DevicePolicyCustomIncludeGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelInclude] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, policyID, query, opts...))
}

type DevicePolicyCustomIncludeUpdateParams struct {
	AccountID param.Field[string]            `path:"account_id,required"`
	Body      []SplitTunnelIncludeUnionParam `json:"body,required"`
}

func (r DevicePolicyCustomIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyCustomIncludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
