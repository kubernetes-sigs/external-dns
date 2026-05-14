// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DEXCommandDeviceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandDeviceService] method instead.
type DEXCommandDeviceService struct {
	Options []option.RequestOption
}

// NewDEXCommandDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXCommandDeviceService(opts ...option.RequestOption) (r *DEXCommandDeviceService) {
	r = &DEXCommandDeviceService{}
	r.Options = opts
	return
}

// List devices with WARP client support for remote captures which have been
// connected in the last 1 hour.
func (r *DEXCommandDeviceService) List(ctx context.Context, params DEXCommandDeviceListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[DEXCommandDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/devices", params.AccountID)
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

// List devices with WARP client support for remote captures which have been
// connected in the last 1 hour.
func (r *DEXCommandDeviceService) ListAutoPaging(ctx context.Context, params DEXCommandDeviceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[DEXCommandDeviceListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type DEXCommandDeviceListResponse struct {
	// List of eligible devices
	Devices []DEXCommandDeviceListResponseDevice `json:"devices"`
	JSON    dexCommandDeviceListResponseJSON     `json:"-"`
}

// dexCommandDeviceListResponseJSON contains the JSON metadata for the struct
// [DEXCommandDeviceListResponse]
type dexCommandDeviceListResponseJSON struct {
	Devices     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandDeviceListResponseDevice struct {
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId"`
	// Device identifier (human readable)
	DeviceName string `json:"deviceName"`
	// Whether the device is eligible for remote captures
	Eligible bool `json:"eligible"`
	// If the device is not eligible, the reason why.
	IneligibleReason string `json:"ineligibleReason"`
	// User contact email address
	PersonEmail string `json:"personEmail"`
	// Operating system
	Platform string `json:"platform"`
	// Network status
	Status string `json:"status"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp"`
	// WARP client version
	Version string                                 `json:"version"`
	JSON    dexCommandDeviceListResponseDeviceJSON `json:"-"`
}

// dexCommandDeviceListResponseDeviceJSON contains the JSON metadata for the struct
// [DEXCommandDeviceListResponseDevice]
type dexCommandDeviceListResponseDeviceJSON struct {
	DeviceID         apijson.Field
	DeviceName       apijson.Field
	Eligible         apijson.Field
	IneligibleReason apijson.Field
	PersonEmail      apijson.Field
	Platform         apijson.Field
	Status           apijson.Field
	Timestamp        apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandDeviceListResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandDeviceListResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type DEXCommandDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Filter devices by name or email
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [DEXCommandDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXCommandDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
