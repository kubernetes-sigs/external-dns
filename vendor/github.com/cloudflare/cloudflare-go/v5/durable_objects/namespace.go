// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package durable_objects

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options []option.RequestOption
	Objects *NamespaceObjectService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Objects = NewNamespaceObjectService(opts...)
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *NamespaceService) List(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Namespace], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", params.AccountID)
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

// Returns the Durable Object namespaces owned by an account.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Namespace] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type Namespace struct {
	ID        string        `json:"id"`
	Class     string        `json:"class"`
	Name      string        `json:"name"`
	Script    string        `json:"script"`
	UseSqlite bool          `json:"use_sqlite"`
	JSON      namespaceJSON `json:"-"`
}

// namespaceJSON contains the JSON metadata for the struct [Namespace]
type namespaceJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	UseSqlite   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Namespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceJSON) RawJSON() string {
	return r.raw
}

type NamespaceListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
