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

// DevicePolicyCustomFallbackDomainService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomFallbackDomainService] method instead.
type DevicePolicyCustomFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCustomFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyCustomFallbackDomainService) {
	r = &DevicePolicyCustomFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyCustomFallbackDomainService) Update(ctx context.Context, policyID string, params DevicePolicyCustomFallbackDomainUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[FallbackDomain], err error) {
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
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", params.AccountID, policyID)
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

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyCustomFallbackDomainService) UpdateAutoPaging(ctx context.Context, policyID string, params DevicePolicyCustomFallbackDomainUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FallbackDomain] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, policyID, params, opts...))
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyCustomFallbackDomainService) Get(ctx context.Context, policyID string, query DevicePolicyCustomFallbackDomainGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[FallbackDomain], err error) {
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
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", query.AccountID, policyID)
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

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyCustomFallbackDomainService) GetAutoPaging(ctx context.Context, policyID string, query DevicePolicyCustomFallbackDomainGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FallbackDomain] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, policyID, query, opts...))
}

type DevicePolicyCustomFallbackDomainUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Domains   []FallbackDomainParam `json:"domains,required"`
}

func (r DevicePolicyCustomFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Domains)
}

type DevicePolicyCustomFallbackDomainGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
