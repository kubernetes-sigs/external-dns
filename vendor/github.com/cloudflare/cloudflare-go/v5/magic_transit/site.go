// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SiteService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteService] method instead.
type SiteService struct {
	Options []option.RequestOption
	ACLs    *SiteACLService
	LANs    *SiteLANService
	WANs    *SiteWANService
}

// NewSiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteService(opts ...option.RequestOption) (r *SiteService) {
	r = &SiteService{}
	r.Options = opts
	r.ACLs = NewSiteACLService(opts...)
	r.LANs = NewSiteLANService(opts...)
	r.WANs = NewSiteWANService(opts...)
	return
}

// Creates a new Site
func (r *SiteService) New(ctx context.Context, params SiteNewParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Site.
func (r *SiteService) Update(ctx context.Context, siteID string, params SiteUpdateParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Sites associated with an account. Use connectorid query param to return
// sites where connectorid matches either site.ConnectorID or
// site.SecondaryConnectorID.
func (r *SiteService) List(ctx context.Context, params SiteListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Site], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites", params.AccountID)
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

// Lists Sites associated with an account. Use connectorid query param to return
// sites where connectorid matches either site.ConnectorID or
// site.SecondaryConnectorID.
func (r *SiteService) ListAutoPaging(ctx context.Context, params SiteListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Site] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Remove a specific Site.
func (r *SiteService) Delete(ctx context.Context, siteID string, body SiteDeleteParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", body.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch a specific Site.
func (r *SiteService) Edit(ctx context.Context, siteID string, params SiteEditParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site.
func (r *SiteService) Get(ctx context.Context, siteID string, params SiteGetParams, opts ...option.RequestOption) (res *Site, err error) {
	var env SiteGetResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Site struct {
	// Identifier
	ID string `json:"id"`
	// Magic Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string   `json:"secondary_connector_id"`
	JSON                 siteJSON `json:"-"`
}

// siteJSON contains the JSON metadata for the struct [Site]
type siteJSON struct {
	ID                   apijson.Field
	ConnectorID          apijson.Field
	Description          apijson.Field
	HaMode               apijson.Field
	Location             apijson.Field
	Name                 apijson.Field
	SecondaryConnectorID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Site) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string           `json:"lon"`
	JSON siteLocationJSON `json:"-"`
}

// siteLocationJSON contains the JSON metadata for the struct [SiteLocation]
type siteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLocationJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteLocationParam struct {
	// Latitude
	Lat param.Field[string] `json:"lat"`
	// Longitude
	Lon param.Field[string] `json:"lon"`
}

func (r SiteLocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the site.
	Name param.Field[string] `json:"name,required"`
	// Magic Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode param.Field[bool] `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteLocationParam] `json:"location"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteNewResponseEnvelopeJSON    `json:"-"`
}

// siteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteNewResponseEnvelope]
type siteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteNewResponseEnvelopeSuccess bool

const (
	SiteNewResponseEnvelopeSuccessTrue SiteNewResponseEnvelopeSuccess = true
)

func (r SiteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Magic Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteLocationParam] `json:"location"`
	// The name of the site.
	Name param.Field[string] `json:"name"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteUpdateResponseEnvelope]
type siteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteUpdateResponseEnvelopeSuccess bool

const (
	SiteUpdateResponseEnvelopeSuccessTrue SiteUpdateResponseEnvelopeSuccess = true
)

func (r SiteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Identifier
	Connectorid param.Field[string] `query:"connectorid"`
}

// URLQuery serializes [SiteListParams]'s query parameters as `url.Values`.
func (r SiteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SiteDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteDeleteResponseEnvelope]
type siteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteDeleteResponseEnvelopeSuccess bool

const (
	SiteDeleteResponseEnvelopeSuccessTrue SiteDeleteResponseEnvelopeSuccess = true
)

func (r SiteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Magic Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteLocationParam] `json:"location"`
	// The name of the site.
	Name param.Field[string] `json:"name"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteEditResponseEnvelopeJSON    `json:"-"`
}

// siteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteEditResponseEnvelope]
type siteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteEditResponseEnvelopeSuccess bool

const (
	SiteEditResponseEnvelopeSuccessTrue SiteEditResponseEnvelopeSuccess = true
)

func (r SiteEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteGetParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type SiteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteGetResponseEnvelopeJSON    `json:"-"`
}

// siteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteGetResponseEnvelope]
type siteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteGetResponseEnvelopeSuccess bool

const (
	SiteGetResponseEnvelopeSuccessTrue SiteGetResponseEnvelopeSuccess = true
)

func (r SiteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
