// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Creates a new App for an account
func (r *AppService) New(ctx context.Context, params AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	var env AppNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an Account App
func (r *AppService) Update(ctx context.Context, accountAppID string, params AppUpdateParams, opts ...option.RequestOption) (res *AppUpdateResponse, err error) {
	var env AppUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if accountAppID == "" {
		err = errors.New("missing required account_app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps/%s", params.AccountID, accountAppID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Apps associated with an account.
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AppListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps", query.AccountID)
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

// Lists Apps associated with an account.
func (r *AppService) ListAutoPaging(ctx context.Context, query AppListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AppListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes specific Account App.
func (r *AppService) Delete(ctx context.Context, accountAppID string, body AppDeleteParams, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	var env AppDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if accountAppID == "" {
		err = errors.New("missing required account_app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps/%s", body.AccountID, accountAppID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an Account App
func (r *AppService) Edit(ctx context.Context, accountAppID string, params AppEditParams, opts ...option.RequestOption) (res *AppEditResponse, err error) {
	var env AppEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if accountAppID == "" {
		err = errors.New("missing required account_app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/apps/%s", params.AccountID, accountAppID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Custom app defined for an account.
type AppNewResponse struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string             `json:"type"`
	JSON appNewResponseJSON `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

// Custom app defined for an account.
type AppUpdateResponse struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string                `json:"type"`
	JSON appUpdateResponseJSON `json:"-"`
}

// appUpdateResponseJSON contains the JSON metadata for the struct
// [AppUpdateResponse]
type appUpdateResponseJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Collection of Hostnames and/or IP Subnets to associate with traffic decisions.
type AppListResponse struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id"`
	// This field can have the runtime type of [[]string].
	Hostnames interface{} `json:"hostnames"`
	// This field can have the runtime type of [[]string].
	IPSubnets interface{} `json:"ip_subnets"`
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type  string              `json:"type"`
	JSON  appListResponseJSON `json:"-"`
	union AppListResponseUnion
}

// appListResponseJSON contains the JSON metadata for the struct [AppListResponse]
type appListResponseJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	ManagedAppID apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r appListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AppListResponseMagicAccountApp],
// [AppListResponseMagicManagedApp].
func (r AppListResponse) AsUnion() AppListResponseUnion {
	return r.union
}

// Collection of Hostnames and/or IP Subnets to associate with traffic decisions.
//
// Union satisfied by [AppListResponseMagicAccountApp] or
// [AppListResponseMagicManagedApp].
type AppListResponseUnion interface {
	implementsAppListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppListResponseMagicAccountApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppListResponseMagicManagedApp{}),
		},
	)
}

// Custom app defined for an account.
type AppListResponseMagicAccountApp struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string                             `json:"type"`
	JSON appListResponseMagicAccountAppJSON `json:"-"`
}

// appListResponseMagicAccountAppJSON contains the JSON metadata for the struct
// [AppListResponseMagicAccountApp]
type appListResponseMagicAccountAppJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppListResponseMagicAccountApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseMagicAccountAppJSON) RawJSON() string {
	return r.raw
}

func (r AppListResponseMagicAccountApp) implementsAppListResponse() {}

// Managed app defined by Cloudflare.
type AppListResponseMagicManagedApp struct {
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string                             `json:"type"`
	JSON appListResponseMagicManagedAppJSON `json:"-"`
}

// appListResponseMagicManagedAppJSON contains the JSON metadata for the struct
// [AppListResponseMagicManagedApp]
type appListResponseMagicManagedAppJSON struct {
	ManagedAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppListResponseMagicManagedApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseMagicManagedAppJSON) RawJSON() string {
	return r.raw
}

func (r AppListResponseMagicManagedApp) implementsAppListResponse() {}

// Custom app defined for an account.
type AppDeleteResponse struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string                `json:"type"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Custom app defined for an account.
type AppEditResponse struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames []string `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets []string `json:"ip_subnets"`
	// Display name for the app.
	Name string `json:"name"`
	// Category of the app.
	Type string              `json:"type"`
	JSON appEditResponseJSON `json:"-"`
}

// appEditResponseJSON contains the JSON metadata for the struct [AppEditResponse]
type appEditResponseJSON struct {
	AccountAppID apijson.Field
	Hostnames    apijson.Field
	IPSubnets    apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appEditResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Display name for the app.
	Name param.Field[string] `json:"name,required"`
	// Category of the app.
	Type param.Field[string] `json:"type,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets param.Field[[]string] `json:"ip_subnets"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Custom app defined for an account.
	Result AppNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    appNewResponseEnvelopeJSON    `json:"-"`
}

// appNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelope]
type appNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppNewResponseEnvelopeSuccess bool

const (
	AppNewResponseEnvelopeSuccessTrue AppNewResponseEnvelopeSuccess = true
)

func (r AppNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets param.Field[[]string] `json:"ip_subnets"`
	// Display name for the app.
	Name param.Field[string] `json:"name"`
	// Category of the app.
	Type param.Field[string] `json:"type"`
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Custom app defined for an account.
	Result AppUpdateResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    appUpdateResponseEnvelopeJSON    `json:"-"`
}

// appUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelope]
type appUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppUpdateResponseEnvelopeSuccess bool

const (
	AppUpdateResponseEnvelopeSuccessTrue AppUpdateResponseEnvelopeSuccess = true
)

func (r AppUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AppDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AppDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Custom app defined for an account.
	Result AppDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    appDeleteResponseEnvelopeJSON    `json:"-"`
}

// appDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelope]
type appDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppDeleteResponseEnvelopeSuccess bool

const (
	AppDeleteResponseEnvelopeSuccessTrue AppDeleteResponseEnvelopeSuccess = true
)

func (r AppDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// FQDNs to associate with traffic decisions.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// IPv4 CIDRs to associate with traffic decisions. (IPv6 CIDRs are currently
	// unsupported)
	IPSubnets param.Field[[]string] `json:"ip_subnets"`
	// Display name for the app.
	Name param.Field[string] `json:"name"`
	// Category of the app.
	Type param.Field[string] `json:"type"`
}

func (r AppEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Custom app defined for an account.
	Result AppEditResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    appEditResponseEnvelopeJSON    `json:"-"`
}

// appEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppEditResponseEnvelope]
type appEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppEditResponseEnvelopeSuccess bool

const (
	AppEditResponseEnvelopeSuccessTrue AppEditResponseEnvelopeSuccess = true
)

func (r AppEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
