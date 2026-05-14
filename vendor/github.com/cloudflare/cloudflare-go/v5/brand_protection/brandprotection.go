// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

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

// BrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandProtectionService] method instead.
type BrandProtectionService struct {
	Options     []option.RequestOption
	Queries     *QueryService
	Matches     *MatchService
	Logos       *LogoService
	LogoMatches *LogoMatchService
}

// NewBrandProtectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandProtectionService(opts ...option.RequestOption) (r *BrandProtectionService) {
	r = &BrandProtectionService{}
	r.Options = opts
	r.Queries = NewQueryService(opts...)
	r.Matches = NewMatchService(opts...)
	r.Logos = NewLogoService(opts...)
	r.LogoMatches = NewLogoMatchService(opts...)
	return
}

// Return new URL submissions
func (r *BrandProtectionService) Submit(ctx context.Context, body BrandProtectionSubmitParams, opts ...option.RequestOption) (res *BrandProtectionSubmitResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Return submitted URLs based on ID
func (r *BrandProtectionService) URLInfo(ctx context.Context, query BrandProtectionURLInfoParams, opts ...option.RequestOption) (res *pagination.SinglePage[BrandProtectionURLInfoResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", query.AccountID)
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

// Return submitted URLs based on ID
func (r *BrandProtectionService) URLInfoAutoPaging(ctx context.Context, query BrandProtectionURLInfoParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[BrandProtectionURLInfoResponse] {
	return pagination.NewSinglePageAutoPager(r.URLInfo(ctx, query, opts...))
}

type BrandProtectionSubmitResponse struct {
	SkippedURLs   []map[string]interface{}          `json:"skipped_urls"`
	SubmittedURLs []map[string]interface{}          `json:"submitted_urls"`
	JSON          brandProtectionSubmitResponseJSON `json:"-"`
}

// brandProtectionSubmitResponseJSON contains the JSON metadata for the struct
// [BrandProtectionSubmitResponse]
type brandProtectionSubmitResponseJSON struct {
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponse map[string]interface{}

type BrandProtectionSubmitParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type BrandProtectionURLInfoParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
