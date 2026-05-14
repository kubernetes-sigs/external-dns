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

// ServiceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *ServiceService) List(ctx context.Context, query ServiceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ServiceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/services", query.AccountID)
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

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *ServiceService) ListAutoPaging(ctx context.Context, query ServiceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ServiceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type ServiceListResponse struct {
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                  `json:"name"`
	JSON serviceListResponseJSON `json:"-"`
}

// serviceListResponseJSON contains the JSON metadata for the struct
// [ServiceListResponse]
type serviceListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}
