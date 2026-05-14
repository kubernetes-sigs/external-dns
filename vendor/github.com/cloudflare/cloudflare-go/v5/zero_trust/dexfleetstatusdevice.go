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

// DEXFleetStatusDeviceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXFleetStatusDeviceService] method instead.
type DEXFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewDEXFleetStatusDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXFleetStatusDeviceService(opts ...option.RequestOption) (r *DEXFleetStatusDeviceService) {
	r = &DEXFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *DEXFleetStatusDeviceService) List(ctx context.Context, params DEXFleetStatusDeviceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DEXFleetStatusDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", params.AccountID)
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

// List details for devices using WARP
func (r *DEXFleetStatusDeviceService) ListAutoPaging(ctx context.Context, params DEXFleetStatusDeviceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DEXFleetStatusDeviceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type DEXFleetStatusDeviceListResponse struct {
	// Cloudflare colo
	Colo string `json:"colo,required"`
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId,required"`
	// The mode under which the WARP client is run
	Mode string `json:"mode,required"`
	// Operating system
	Platform string `json:"platform,required"`
	// Network status
	Status string `json:"status,required"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp,required"`
	// WARP client version
	Version         string                                          `json:"version,required"`
	AlwaysOn        bool                                            `json:"alwaysOn,nullable"`
	BatteryCharging bool                                            `json:"batteryCharging,nullable"`
	BatteryCycles   int64                                           `json:"batteryCycles,nullable"`
	BatteryPct      float64                                         `json:"batteryPct,nullable"`
	ConnectionType  string                                          `json:"connectionType,nullable"`
	CPUPct          float64                                         `json:"cpuPct,nullable"`
	CPUPctByApp     [][]DEXFleetStatusDeviceListResponseCPUPctByApp `json:"cpuPctByApp,nullable"`
	DeviceIPV4      DEXFleetStatusDeviceListResponseDeviceIPV4      `json:"deviceIpv4"`
	DeviceIPV6      DEXFleetStatusDeviceListResponseDeviceIPV6      `json:"deviceIpv6"`
	// Device identifier (human readable)
	DeviceName         string                                      `json:"deviceName"`
	DiskReadBps        int64                                       `json:"diskReadBps,nullable"`
	DiskUsagePct       float64                                     `json:"diskUsagePct,nullable"`
	DiskWriteBps       int64                                       `json:"diskWriteBps,nullable"`
	DOHSubdomain       string                                      `json:"dohSubdomain,nullable"`
	EstimatedLossPct   float64                                     `json:"estimatedLossPct,nullable"`
	FirewallEnabled    bool                                        `json:"firewallEnabled,nullable"`
	GatewayIPV4        DEXFleetStatusDeviceListResponseGatewayIPV4 `json:"gatewayIpv4"`
	GatewayIPV6        DEXFleetStatusDeviceListResponseGatewayIPV6 `json:"gatewayIpv6"`
	HandshakeLatencyMs float64                                     `json:"handshakeLatencyMs,nullable"`
	ISPIPV4            DEXFleetStatusDeviceListResponseISPIPV4     `json:"ispIpv4"`
	ISPIPV6            DEXFleetStatusDeviceListResponseISPIPV6     `json:"ispIpv6"`
	Metal              string                                      `json:"metal,nullable"`
	NetworkRcvdBps     int64                                       `json:"networkRcvdBps,nullable"`
	NetworkSentBps     int64                                       `json:"networkSentBps,nullable"`
	NetworkSsid        string                                      `json:"networkSsid,nullable"`
	// User contact email address
	PersonEmail     string                                              `json:"personEmail"`
	RamAvailableKB  int64                                               `json:"ramAvailableKb,nullable"`
	RamUsedPct      float64                                             `json:"ramUsedPct,nullable"`
	RamUsedPctByApp [][]DEXFleetStatusDeviceListResponseRamUsedPctByApp `json:"ramUsedPctByApp,nullable"`
	SwitchLocked    bool                                                `json:"switchLocked,nullable"`
	WifiStrengthDbm int64                                               `json:"wifiStrengthDbm,nullable"`
	JSON            dexFleetStatusDeviceListResponseJSON                `json:"-"`
}

// dexFleetStatusDeviceListResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusDeviceListResponse]
type dexFleetStatusDeviceListResponseJSON struct {
	Colo               apijson.Field
	DeviceID           apijson.Field
	Mode               apijson.Field
	Platform           apijson.Field
	Status             apijson.Field
	Timestamp          apijson.Field
	Version            apijson.Field
	AlwaysOn           apijson.Field
	BatteryCharging    apijson.Field
	BatteryCycles      apijson.Field
	BatteryPct         apijson.Field
	ConnectionType     apijson.Field
	CPUPct             apijson.Field
	CPUPctByApp        apijson.Field
	DeviceIPV4         apijson.Field
	DeviceIPV6         apijson.Field
	DeviceName         apijson.Field
	DiskReadBps        apijson.Field
	DiskUsagePct       apijson.Field
	DiskWriteBps       apijson.Field
	DOHSubdomain       apijson.Field
	EstimatedLossPct   apijson.Field
	FirewallEnabled    apijson.Field
	GatewayIPV4        apijson.Field
	GatewayIPV6        apijson.Field
	HandshakeLatencyMs apijson.Field
	ISPIPV4            apijson.Field
	ISPIPV6            apijson.Field
	Metal              apijson.Field
	NetworkRcvdBps     apijson.Field
	NetworkSentBps     apijson.Field
	NetworkSsid        apijson.Field
	PersonEmail        apijson.Field
	RamAvailableKB     apijson.Field
	RamUsedPct         apijson.Field
	RamUsedPctByApp    apijson.Field
	SwitchLocked       apijson.Field
	WifiStrengthDbm    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseCPUPctByApp struct {
	CPUPct float64                                         `json:"cpu_pct"`
	Name   string                                          `json:"name"`
	JSON   dexFleetStatusDeviceListResponseCPUPctByAppJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseCPUPctByAppJSON contains the JSON metadata for
// the struct [DEXFleetStatusDeviceListResponseCPUPctByApp]
type dexFleetStatusDeviceListResponseCPUPctByAppJSON struct {
	CPUPct      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseCPUPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseCPUPctByAppJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseDeviceIPV4 struct {
	Address  string                                             `json:"address,nullable"`
	ASN      int64                                              `json:"asn,nullable"`
	Aso      string                                             `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseDeviceIPV4Location `json:"location"`
	Netmask  string                                             `json:"netmask,nullable"`
	Version  string                                             `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseDeviceIPV4JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseDeviceIPV4JSON contains the JSON metadata for
// the struct [DEXFleetStatusDeviceListResponseDeviceIPV4]
type dexFleetStatusDeviceListResponseDeviceIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseDeviceIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseDeviceIPV4JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseDeviceIPV4Location struct {
	City       string                                                 `json:"city,nullable"`
	CountryISO string                                                 `json:"country_iso,nullable"`
	StateISO   string                                                 `json:"state_iso,nullable"`
	Zip        string                                                 `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseDeviceIPV4LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseDeviceIPV4LocationJSON contains the JSON
// metadata for the struct [DEXFleetStatusDeviceListResponseDeviceIPV4Location]
type dexFleetStatusDeviceListResponseDeviceIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseDeviceIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseDeviceIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseDeviceIPV6 struct {
	Address  string                                             `json:"address,nullable"`
	ASN      int64                                              `json:"asn,nullable"`
	Aso      string                                             `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseDeviceIPV6Location `json:"location"`
	Netmask  string                                             `json:"netmask,nullable"`
	Version  string                                             `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseDeviceIPV6JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseDeviceIPV6JSON contains the JSON metadata for
// the struct [DEXFleetStatusDeviceListResponseDeviceIPV6]
type dexFleetStatusDeviceListResponseDeviceIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseDeviceIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseDeviceIPV6JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseDeviceIPV6Location struct {
	City       string                                                 `json:"city,nullable"`
	CountryISO string                                                 `json:"country_iso,nullable"`
	StateISO   string                                                 `json:"state_iso,nullable"`
	Zip        string                                                 `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseDeviceIPV6LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseDeviceIPV6LocationJSON contains the JSON
// metadata for the struct [DEXFleetStatusDeviceListResponseDeviceIPV6Location]
type dexFleetStatusDeviceListResponseDeviceIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseDeviceIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseDeviceIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseGatewayIPV4 struct {
	Address  string                                              `json:"address,nullable"`
	ASN      int64                                               `json:"asn,nullable"`
	Aso      string                                              `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseGatewayIPV4Location `json:"location"`
	Netmask  string                                              `json:"netmask,nullable"`
	Version  string                                              `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseGatewayIPV4JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseGatewayIPV4JSON contains the JSON metadata for
// the struct [DEXFleetStatusDeviceListResponseGatewayIPV4]
type dexFleetStatusDeviceListResponseGatewayIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseGatewayIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseGatewayIPV4JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseGatewayIPV4Location struct {
	City       string                                                  `json:"city,nullable"`
	CountryISO string                                                  `json:"country_iso,nullable"`
	StateISO   string                                                  `json:"state_iso,nullable"`
	Zip        string                                                  `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseGatewayIPV4LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseGatewayIPV4LocationJSON contains the JSON
// metadata for the struct [DEXFleetStatusDeviceListResponseGatewayIPV4Location]
type dexFleetStatusDeviceListResponseGatewayIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseGatewayIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseGatewayIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseGatewayIPV6 struct {
	Address  string                                              `json:"address,nullable"`
	ASN      int64                                               `json:"asn,nullable"`
	Aso      string                                              `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseGatewayIPV6Location `json:"location"`
	Netmask  string                                              `json:"netmask,nullable"`
	Version  string                                              `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseGatewayIPV6JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseGatewayIPV6JSON contains the JSON metadata for
// the struct [DEXFleetStatusDeviceListResponseGatewayIPV6]
type dexFleetStatusDeviceListResponseGatewayIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseGatewayIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseGatewayIPV6JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseGatewayIPV6Location struct {
	City       string                                                  `json:"city,nullable"`
	CountryISO string                                                  `json:"country_iso,nullable"`
	StateISO   string                                                  `json:"state_iso,nullable"`
	Zip        string                                                  `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseGatewayIPV6LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseGatewayIPV6LocationJSON contains the JSON
// metadata for the struct [DEXFleetStatusDeviceListResponseGatewayIPV6Location]
type dexFleetStatusDeviceListResponseGatewayIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseGatewayIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseGatewayIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseISPIPV4 struct {
	Address  string                                          `json:"address,nullable"`
	ASN      int64                                           `json:"asn,nullable"`
	Aso      string                                          `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseISPIPV4Location `json:"location"`
	Netmask  string                                          `json:"netmask,nullable"`
	Version  string                                          `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseIspipv4JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseIspipv4JSON contains the JSON metadata for the
// struct [DEXFleetStatusDeviceListResponseISPIPV4]
type dexFleetStatusDeviceListResponseIspipv4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseISPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseIspipv4JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseISPIPV4Location struct {
	City       string                                              `json:"city,nullable"`
	CountryISO string                                              `json:"country_iso,nullable"`
	StateISO   string                                              `json:"state_iso,nullable"`
	Zip        string                                              `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseIspipv4LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseIspipv4LocationJSON contains the JSON metadata
// for the struct [DEXFleetStatusDeviceListResponseISPIPV4Location]
type dexFleetStatusDeviceListResponseIspipv4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseISPIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseIspipv4LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseISPIPV6 struct {
	Address  string                                          `json:"address,nullable"`
	ASN      int64                                           `json:"asn,nullable"`
	Aso      string                                          `json:"aso,nullable"`
	Location DEXFleetStatusDeviceListResponseISPIPV6Location `json:"location"`
	Netmask  string                                          `json:"netmask,nullable"`
	Version  string                                          `json:"version,nullable"`
	JSON     dexFleetStatusDeviceListResponseIspipv6JSON     `json:"-"`
}

// dexFleetStatusDeviceListResponseIspipv6JSON contains the JSON metadata for the
// struct [DEXFleetStatusDeviceListResponseISPIPV6]
type dexFleetStatusDeviceListResponseIspipv6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseISPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseIspipv6JSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseISPIPV6Location struct {
	City       string                                              `json:"city,nullable"`
	CountryISO string                                              `json:"country_iso,nullable"`
	StateISO   string                                              `json:"state_iso,nullable"`
	Zip        string                                              `json:"zip,nullable"`
	JSON       dexFleetStatusDeviceListResponseIspipv6LocationJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseIspipv6LocationJSON contains the JSON metadata
// for the struct [DEXFleetStatusDeviceListResponseISPIPV6Location]
type dexFleetStatusDeviceListResponseIspipv6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseISPIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseIspipv6LocationJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListResponseRamUsedPctByApp struct {
	Name       string                                              `json:"name"`
	RamUsedPct float64                                             `json:"ram_used_pct"`
	JSON       dexFleetStatusDeviceListResponseRamUsedPctByAppJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseRamUsedPctByAppJSON contains the JSON metadata
// for the struct [DEXFleetStatusDeviceListResponseRamUsedPctByApp]
type dexFleetStatusDeviceListResponseRamUsedPctByAppJSON struct {
	Name        apijson.Field
	RamUsedPct  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponseRamUsedPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseRamUsedPctByAppJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time range beginning in ISO format
	From param.Field[string] `query:"from,required"`
	// Page number
	Page param.Field[float64] `query:"page,required"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Time range end in ISO format
	To param.Field[string] `query:"to,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
	// The mode under which the WARP client is run
	Mode param.Field[string] `query:"mode"`
	// Operating system
	Platform param.Field[string] `query:"platform"`
	// Dimension to sort results by
	SortBy param.Field[DEXFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Source:
	//
	// - `hourly` - device details aggregated hourly, up to 7 days prior
	// - `last_seen` - device details, up to 24 hours prior
	// - `raw` - device details, up to 7 days prior
	Source param.Field[DEXFleetStatusDeviceListParamsSource] `query:"source"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [DEXFleetStatusDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Dimension to sort results by
type DEXFleetStatusDeviceListParamsSortBy string

const (
	DEXFleetStatusDeviceListParamsSortByColo      DEXFleetStatusDeviceListParamsSortBy = "colo"
	DEXFleetStatusDeviceListParamsSortByDeviceID  DEXFleetStatusDeviceListParamsSortBy = "device_id"
	DEXFleetStatusDeviceListParamsSortByMode      DEXFleetStatusDeviceListParamsSortBy = "mode"
	DEXFleetStatusDeviceListParamsSortByPlatform  DEXFleetStatusDeviceListParamsSortBy = "platform"
	DEXFleetStatusDeviceListParamsSortByStatus    DEXFleetStatusDeviceListParamsSortBy = "status"
	DEXFleetStatusDeviceListParamsSortByTimestamp DEXFleetStatusDeviceListParamsSortBy = "timestamp"
	DEXFleetStatusDeviceListParamsSortByVersion   DEXFleetStatusDeviceListParamsSortBy = "version"
)

func (r DEXFleetStatusDeviceListParamsSortBy) IsKnown() bool {
	switch r {
	case DEXFleetStatusDeviceListParamsSortByColo, DEXFleetStatusDeviceListParamsSortByDeviceID, DEXFleetStatusDeviceListParamsSortByMode, DEXFleetStatusDeviceListParamsSortByPlatform, DEXFleetStatusDeviceListParamsSortByStatus, DEXFleetStatusDeviceListParamsSortByTimestamp, DEXFleetStatusDeviceListParamsSortByVersion:
		return true
	}
	return false
}

// Source:
//
// - `hourly` - device details aggregated hourly, up to 7 days prior
// - `last_seen` - device details, up to 24 hours prior
// - `raw` - device details, up to 7 days prior
type DEXFleetStatusDeviceListParamsSource string

const (
	DEXFleetStatusDeviceListParamsSourceLastSeen DEXFleetStatusDeviceListParamsSource = "last_seen"
	DEXFleetStatusDeviceListParamsSourceHourly   DEXFleetStatusDeviceListParamsSource = "hourly"
	DEXFleetStatusDeviceListParamsSourceRaw      DEXFleetStatusDeviceListParamsSource = "raw"
)

func (r DEXFleetStatusDeviceListParamsSource) IsKnown() bool {
	switch r {
	case DEXFleetStatusDeviceListParamsSourceLastSeen, DEXFleetStatusDeviceListParamsSourceHourly, DEXFleetStatusDeviceListParamsSourceRaw:
		return true
	}
	return false
}
