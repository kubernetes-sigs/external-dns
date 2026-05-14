// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

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

// AuthorService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorService] method instead.
type AuthorService struct {
	Options []option.RequestOption
}

// NewAuthorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthorService(opts ...option.RequestOption) (r *AuthorService) {
	r = &AuthorService{}
	r.Options = opts
	return
}

// Author Search
func (r *AuthorService) List(ctx context.Context, query AuthorListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AuthorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/authors/search", query.AccountID)
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

// Author Search
func (r *AuthorService) ListAutoPaging(ctx context.Context, query AuthorListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AuthorListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type AuthorListResponse = interface{}

type AuthorListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
