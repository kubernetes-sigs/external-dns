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

// DEXColoService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXColoService] method instead.
type DEXColoService struct {
	Options []option.RequestOption
}

// NewDEXColoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXColoService(opts ...option.RequestOption) (r *DEXColoService) {
	r = &DEXColoService{}
	r.Options = opts
	return
}

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *DEXColoService) List(ctx context.Context, params DEXColoListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DEXColoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/colos", params.AccountID)
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

// List Cloudflare colos that account's devices were connected to during a time
// period, sorted by usage starting from the most used colo. Colos without traffic
// are also returned and sorted alphabetically.
func (r *DEXColoService) ListAutoPaging(ctx context.Context, params DEXColoListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DEXColoListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type DEXColoListResponse = interface{}

type DEXColoListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Start time for connection period in ISO (RFC3339 - ISO 8601) format
	From param.Field[string] `query:"from,required"`
	// End time for connection period in ISO (RFC3339 - ISO 8601) format
	To param.Field[string] `query:"to,required"`
	// Type of usage that colos should be sorted by. If unspecified, returns all
	// Cloudflare colos sorted alphabetically.
	SortBy param.Field[DEXColoListParamsSortBy] `query:"sortBy"`
}

// URLQuery serializes [DEXColoListParams]'s query parameters as `url.Values`.
func (r DEXColoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Type of usage that colos should be sorted by. If unspecified, returns all
// Cloudflare colos sorted alphabetically.
type DEXColoListParamsSortBy string

const (
	DEXColoListParamsSortByFleetStatusUsage      DEXColoListParamsSortBy = "fleet-status-usage"
	DEXColoListParamsSortByApplicationTestsUsage DEXColoListParamsSortBy = "application-tests-usage"
)

func (r DEXColoListParamsSortBy) IsKnown() bool {
	switch r {
	case DEXColoListParamsSortByFleetStatusUsage, DEXColoListParamsSortByApplicationTestsUsage:
		return true
	}
	return false
}
