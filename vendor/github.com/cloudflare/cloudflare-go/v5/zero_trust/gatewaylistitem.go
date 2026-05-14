// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// GatewayListItemService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayListItemService] method instead.
type GatewayListItemService struct {
	Options []option.RequestOption
}

// NewGatewayListItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayListItemService(opts ...option.RequestOption) (r *GatewayListItemService) {
	r = &GatewayListItemService{}
	r.Options = opts
	return
}

// Fetches all items in a single Zero Trust list.
func (r *GatewayListItemService) List(ctx context.Context, listID string, query GatewayListItemListParams, opts ...option.RequestOption) (res *pagination.SinglePage[[]GatewayItem], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s/items", query.AccountID, listID)
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

// Fetches all items in a single Zero Trust list.
func (r *GatewayListItemService) ListAutoPaging(ctx context.Context, listID string, query GatewayListItemListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[[]GatewayItem] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, listID, query, opts...))
}

type GatewayListItemListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
