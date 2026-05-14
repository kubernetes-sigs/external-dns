// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

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
)

// LogoMatchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogoMatchService] method instead.
type LogoMatchService struct {
	Options []option.RequestOption
}

// NewLogoMatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogoMatchService(opts ...option.RequestOption) (r *LogoMatchService) {
	r = &LogoMatchService{}
	r.Options = opts
	return
}

// Return matches as CSV for logo queries based on ID
func (r *LogoMatchService) Download(ctx context.Context, params LogoMatchDownloadParams, opts ...option.RequestOption) (res *LogoMatchDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/logo-matches/download", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Return matches for logo queries based on ID
func (r *LogoMatchService) Get(ctx context.Context, params LogoMatchGetParams, opts ...option.RequestOption) (res *LogoMatchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/logo-matches", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type LogoMatchDownloadResponse struct {
	Matches []map[string]interface{}      `json:"matches"`
	Total   int64                         `json:"total"`
	JSON    logoMatchDownloadResponseJSON `json:"-"`
}

// logoMatchDownloadResponseJSON contains the JSON metadata for the struct
// [LogoMatchDownloadResponse]
type logoMatchDownloadResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogoMatchDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logoMatchDownloadResponseJSON) RawJSON() string {
	return r.raw
}

type LogoMatchGetResponse struct {
	Matches []map[string]interface{} `json:"matches"`
	Total   int64                    `json:"total"`
	JSON    logoMatchGetResponseJSON `json:"-"`
}

// logoMatchGetResponseJSON contains the JSON metadata for the struct
// [LogoMatchGetResponse]
type logoMatchGetResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogoMatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logoMatchGetResponseJSON) RawJSON() string {
	return r.raw
}

type LogoMatchDownloadParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Limit     param.Field[string]   `query:"limit"`
	LogoID    param.Field[[]string] `query:"logo_id"`
	Offset    param.Field[string]   `query:"offset"`
}

// URLQuery serializes [LogoMatchDownloadParams]'s query parameters as
// `url.Values`.
func (r LogoMatchDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogoMatchGetParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Limit     param.Field[string]   `query:"limit"`
	LogoID    param.Field[[]string] `query:"logo_id"`
	Offset    param.Field[string]   `query:"offset"`
}

// URLQuery serializes [LogoMatchGetParams]'s query parameters as `url.Values`.
func (r LogoMatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
