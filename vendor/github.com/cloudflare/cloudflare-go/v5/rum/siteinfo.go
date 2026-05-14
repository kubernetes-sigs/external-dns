// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rum

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SiteInfoService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteInfoService] method instead.
type SiteInfoService struct {
	Options []option.RequestOption
}

// NewSiteInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteInfoService(opts ...option.RequestOption) (r *SiteInfoService) {
	r = &SiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *SiteInfoService) New(ctx context.Context, params SiteInfoNewParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteInfoNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Web Analytics site.
func (r *SiteInfoService) Update(ctx context.Context, siteID string, params SiteInfoUpdateParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteInfoUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Web Analytics sites of an account.
func (r *SiteInfoService) List(ctx context.Context, params SiteInfoListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Site], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", params.AccountID)
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

// Lists all Web Analytics sites of an account.
func (r *SiteInfoService) ListAutoPaging(ctx context.Context, params SiteInfoListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Site] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing Web Analytics site.
func (r *SiteInfoService) Delete(ctx context.Context, siteID string, body SiteInfoDeleteParams, opts ...option.RequestOption) (res *SiteInfoDeleteResponse, err error) {
	var env SiteInfoDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", body.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a Web Analytics site.
func (r *SiteInfoService) Get(ctx context.Context, siteID string, query SiteInfoGetParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteInfoGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Site struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RUMRule   `json:"rules"`
	Ruleset SiteRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string   `json:"snippet"`
	JSON    siteJSON `json:"-"`
}

// siteJSON contains the JSON metadata for the struct [Site]
type siteJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Site) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteJSON) RawJSON() string {
	return r.raw
}

type SiteRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string          `json:"zone_tag"`
	JSON    siteRulesetJSON `json:"-"`
}

// siteRulesetJSON contains the JSON metadata for the struct [SiteRuleset]
type siteRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteRulesetJSON) RawJSON() string {
	return r.raw
}

type SiteInfoDeleteResponse struct {
	// The Web Analytics site identifier.
	SiteTag string                     `json:"site_tag"`
	JSON    siteInfoDeleteResponseJSON `json:"-"`
}

// siteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [SiteInfoDeleteResponse]
type siteInfoDeleteResponseJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r SiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteInfoNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  Site                            `json:"result"`
	JSON    siteInfoNewResponseEnvelopeJSON `json:"-"`
}

// siteInfoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoNewResponseEnvelope]
type siteInfoNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// Enables or disables RUM. This option can be used only when auto_install is set
	// to true.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// If enabled, the JavaScript snippet will not be injected for visitors from the
	// EU.
	Lite param.Field[bool] `json:"lite"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r SiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteInfoUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                               `json:"success,required"`
	Result  Site                               `json:"result"`
	JSON    siteInfoUpdateResponseEnvelopeJSON `json:"-"`
}

// siteInfoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoUpdateResponseEnvelope]
type siteInfoUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The property used to sort the list of results.
	OrderBy param.Field[SiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [SiteInfoListParams]'s query parameters as `url.Values`.
func (r SiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The property used to sort the list of results.
type SiteInfoListParamsOrderBy string

const (
	SiteInfoListParamsOrderByHost    SiteInfoListParamsOrderBy = "host"
	SiteInfoListParamsOrderByCreated SiteInfoListParamsOrderBy = "created"
)

func (r SiteInfoListParamsOrderBy) IsKnown() bool {
	switch r {
	case SiteInfoListParamsOrderByHost, SiteInfoListParamsOrderByCreated:
		return true
	}
	return false
}

type SiteInfoDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteInfoDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                               `json:"success,required"`
	Result  SiteInfoDeleteResponse             `json:"result"`
	JSON    siteInfoDeleteResponseEnvelopeJSON `json:"-"`
}

// siteInfoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoDeleteResponseEnvelope]
type siteInfoDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteInfoGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  Site                            `json:"result"`
	JSON    siteInfoGetResponseEnvelopeJSON `json:"-"`
}

// siteInfoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoGetResponseEnvelope]
type siteInfoGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
