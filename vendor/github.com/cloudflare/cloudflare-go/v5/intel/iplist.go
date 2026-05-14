// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// IPListService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPListService] method instead.
type IPListService struct {
	Options []option.RequestOption
}

// NewIPListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPListService(opts ...option.RequestOption) (r *IPListService) {
	r = &IPListService{}
	r.Options = opts
	return
}

// Get IP Lists.
func (r *IPListService) Get(ctx context.Context, query IPListGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[IPList], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/ip-list", query.AccountID)
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

// Get IP Lists.
func (r *IPListService) GetAutoPaging(ctx context.Context, query IPListGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IPList] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type IPList struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Name        string     `json:"name"`
	JSON        ipListJSON `json:"-"`
}

// ipListJSON contains the JSON metadata for the struct [IPList]
type ipListJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListJSON) RawJSON() string {
	return r.raw
}

type IPListGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
