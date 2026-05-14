// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// AccessCertificateSettingService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessCertificateSettingService] method instead.
type AccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewAccessCertificateSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessCertificateSettingService(opts ...option.RequestOption) (r *AccessCertificateSettingService) {
	r = &AccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccessCertificateSettingService) Update(ctx context.Context, params AccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[CertificateSettings], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Updates an mTLS certificate's hostname settings.
func (r *AccessCertificateSettingService) UpdateAutoPaging(ctx context.Context, params AccessCertificateSettingUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CertificateSettings] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// List all mTLS hostname settings for this account or zone.
func (r *AccessCertificateSettingService) Get(ctx context.Context, query AccessCertificateSettingGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[CertificateSettings], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
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

// List all mTLS hostname settings for this account or zone.
func (r *AccessCertificateSettingService) GetAutoPaging(ctx context.Context, query AccessCertificateSettingGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CertificateSettings] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type CertificateSettings struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                  `json:"hostname,required"`
	JSON     certificateSettingsJSON `json:"-"`
}

// certificateSettingsJSON contains the JSON metadata for the struct
// [CertificateSettings]
type certificateSettingsJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateSettingsJSON) RawJSON() string {
	return r.raw
}

type CertificateSettingsParam struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork param.Field[bool] `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding param.Field[bool] `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname param.Field[string] `json:"hostname,required"`
}

func (r CertificateSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]CertificateSettingsParam] `json:"settings,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}
