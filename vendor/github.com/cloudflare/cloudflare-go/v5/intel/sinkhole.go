// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// SinkholeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSinkholeService] method instead.
type SinkholeService struct {
	Options []option.RequestOption
}

// NewSinkholeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSinkholeService(opts ...option.RequestOption) (r *SinkholeService) {
	r = &SinkholeService{}
	r.Options = opts
	return
}

// List sinkholes owned by this account
func (r *SinkholeService) List(ctx context.Context, query SinkholeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Sinkhole], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", query.AccountID)
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

// List sinkholes owned by this account
func (r *SinkholeService) ListAutoPaging(ctx context.Context, query SinkholeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Sinkhole] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type Sinkhole struct {
	// The unique identifier for the sinkhole
	ID int64 `json:"id"`
	// The account tag that owns this sinkhole
	AccountTag string `json:"account_tag"`
	// The date and time when the sinkhole was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the sinkhole was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the sinkhole
	Name string `json:"name"`
	// The name of the R2 bucket to store results
	R2Bucket string `json:"r2_bucket"`
	// The id of the R2 instance
	R2ID string       `json:"r2_id"`
	JSON sinkholeJSON `json:"-"`
}

// sinkholeJSON contains the JSON metadata for the struct [Sinkhole]
type sinkholeJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	R2Bucket    apijson.Field
	R2ID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Sinkhole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeJSON) RawJSON() string {
	return r.raw
}

type SinkholeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
