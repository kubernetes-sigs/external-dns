// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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

// PoolReferenceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPoolReferenceService] method instead.
type PoolReferenceService struct {
	Options []option.RequestOption
}

// NewPoolReferenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPoolReferenceService(opts ...option.RequestOption) (r *PoolReferenceService) {
	r = &PoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *PoolReferenceService) Get(ctx context.Context, poolID string, query PoolReferenceGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[PoolReferenceGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", query.AccountID, poolID)
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

// Get the list of resources that reference the provided pool.
func (r *PoolReferenceService) GetAutoPaging(ctx context.Context, poolID string, query PoolReferenceGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PoolReferenceGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, poolID, query, opts...))
}

type PoolReferenceGetResponse struct {
	ReferenceType PoolReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                `json:"resource_id"`
	ResourceName  string                                `json:"resource_name"`
	ResourceType  string                                `json:"resource_type"`
	JSON          poolReferenceGetResponseJSON          `json:"-"`
}

// poolReferenceGetResponseJSON contains the JSON metadata for the struct
// [PoolReferenceGetResponse]
type poolReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PoolReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type PoolReferenceGetResponseReferenceType string

const (
	PoolReferenceGetResponseReferenceTypeStar     PoolReferenceGetResponseReferenceType = "*"
	PoolReferenceGetResponseReferenceTypeReferral PoolReferenceGetResponseReferenceType = "referral"
	PoolReferenceGetResponseReferenceTypeReferrer PoolReferenceGetResponseReferenceType = "referrer"
)

func (r PoolReferenceGetResponseReferenceType) IsKnown() bool {
	switch r {
	case PoolReferenceGetResponseReferenceTypeStar, PoolReferenceGetResponseReferenceTypeReferral, PoolReferenceGetResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type PoolReferenceGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
