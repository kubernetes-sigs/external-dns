// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Update the current settings for the active account
func (r *SettingService) Update(ctx context.Context, params SettingUpdateParams, opts ...option.RequestOption) (res *SettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get the current settings for the active account
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettingUpdateResponse struct {
	DefaultASN int64                     `json:"default_asn,required"`
	JSON       settingUpdateResponseJSON `json:"-"`
}

// settingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingUpdateResponse]
type settingUpdateResponseJSON struct {
	DefaultASN  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponse struct {
	DefaultASN int64                  `json:"default_asn,required"`
	JSON       settingGetResponseJSON `json:"-"`
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	DefaultASN  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingUpdateParams struct {
	AccountID  param.Field[string] `path:"account_id,required"`
	DefaultASN param.Field[int64]  `json:"default_asn"`
}

func (r SettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
