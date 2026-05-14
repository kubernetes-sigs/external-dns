// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

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

// RegionalHostnameRegionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalHostnameRegionService] method instead.
type RegionalHostnameRegionService struct {
	Options []option.RequestOption
}

// NewRegionalHostnameRegionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalHostnameRegionService(opts ...option.RequestOption) (r *RegionalHostnameRegionService) {
	r = &RegionalHostnameRegionService{}
	r.Options = opts
	return
}

// List all Regional Services regions available for use by this account.
func (r *RegionalHostnameRegionService) List(ctx context.Context, query RegionalHostnameRegionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RegionalHostnameRegionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/regional_hostnames/regions", query.AccountID)
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

// List all Regional Services regions available for use by this account.
func (r *RegionalHostnameRegionService) ListAutoPaging(ctx context.Context, query RegionalHostnameRegionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RegionalHostnameRegionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type RegionalHostnameRegionListResponse struct {
	// Identifying key for the region
	Key string `json:"key"`
	// Human-readable text label for the region
	Label string                                 `json:"label"`
	JSON  regionalHostnameRegionListResponseJSON `json:"-"`
}

// regionalHostnameRegionListResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameRegionListResponse]
type regionalHostnameRegionListResponseJSON struct {
	Key         apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameRegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameRegionListResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameRegionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
