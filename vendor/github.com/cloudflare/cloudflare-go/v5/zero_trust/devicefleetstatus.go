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
)

// DeviceFleetStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceFleetStatusService] method instead.
type DeviceFleetStatusService struct {
	Options []option.RequestOption
}

// NewDeviceFleetStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceFleetStatusService(opts ...option.RequestOption) (r *DeviceFleetStatusService) {
	r = &DeviceFleetStatusService{}
	r.Options = opts
	return
}

// Get the live status of a latest device given device_id from the device_state
// table
func (r *DeviceFleetStatusService) Get(ctx context.Context, deviceID string, params DeviceFleetStatusGetParams, opts ...option.RequestOption) (res *DeviceFleetStatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/%s/fleet-status/live", params.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type DeviceFleetStatusGetResponse struct {
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
	Version         string                                      `json:"version,required"`
	AlwaysOn        bool                                        `json:"alwaysOn,nullable"`
	BatteryCharging bool                                        `json:"batteryCharging,nullable"`
	BatteryCycles   int64                                       `json:"batteryCycles,nullable"`
	BatteryPct      float64                                     `json:"batteryPct,nullable"`
	ConnectionType  string                                      `json:"connectionType,nullable"`
	CPUPct          float64                                     `json:"cpuPct,nullable"`
	CPUPctByApp     [][]DeviceFleetStatusGetResponseCPUPctByApp `json:"cpuPctByApp,nullable"`
	DeviceIPV4      DeviceFleetStatusGetResponseDeviceIPV4      `json:"deviceIpv4"`
	DeviceIPV6      DeviceFleetStatusGetResponseDeviceIPV6      `json:"deviceIpv6"`
	// Device identifier (human readable)
	DeviceName         string                                  `json:"deviceName"`
	DiskReadBps        int64                                   `json:"diskReadBps,nullable"`
	DiskUsagePct       float64                                 `json:"diskUsagePct,nullable"`
	DiskWriteBps       int64                                   `json:"diskWriteBps,nullable"`
	DOHSubdomain       string                                  `json:"dohSubdomain,nullable"`
	EstimatedLossPct   float64                                 `json:"estimatedLossPct,nullable"`
	FirewallEnabled    bool                                    `json:"firewallEnabled,nullable"`
	GatewayIPV4        DeviceFleetStatusGetResponseGatewayIPV4 `json:"gatewayIpv4"`
	GatewayIPV6        DeviceFleetStatusGetResponseGatewayIPV6 `json:"gatewayIpv6"`
	HandshakeLatencyMs float64                                 `json:"handshakeLatencyMs,nullable"`
	ISPIPV4            DeviceFleetStatusGetResponseISPIPV4     `json:"ispIpv4"`
	ISPIPV6            DeviceFleetStatusGetResponseISPIPV6     `json:"ispIpv6"`
	Metal              string                                  `json:"metal,nullable"`
	NetworkRcvdBps     int64                                   `json:"networkRcvdBps,nullable"`
	NetworkSentBps     int64                                   `json:"networkSentBps,nullable"`
	NetworkSsid        string                                  `json:"networkSsid,nullable"`
	// User contact email address
	PersonEmail     string                                          `json:"personEmail"`
	RamAvailableKB  int64                                           `json:"ramAvailableKb,nullable"`
	RamUsedPct      float64                                         `json:"ramUsedPct,nullable"`
	RamUsedPctByApp [][]DeviceFleetStatusGetResponseRamUsedPctByApp `json:"ramUsedPctByApp,nullable"`
	SwitchLocked    bool                                            `json:"switchLocked,nullable"`
	WifiStrengthDbm int64                                           `json:"wifiStrengthDbm,nullable"`
	JSON            deviceFleetStatusGetResponseJSON                `json:"-"`
}

// deviceFleetStatusGetResponseJSON contains the JSON metadata for the struct
// [DeviceFleetStatusGetResponse]
type deviceFleetStatusGetResponseJSON struct {
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

func (r *DeviceFleetStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseCPUPctByApp struct {
	CPUPct float64                                     `json:"cpu_pct"`
	Name   string                                      `json:"name"`
	JSON   deviceFleetStatusGetResponseCPUPctByAppJSON `json:"-"`
}

// deviceFleetStatusGetResponseCPUPctByAppJSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseCPUPctByApp]
type deviceFleetStatusGetResponseCPUPctByAppJSON struct {
	CPUPct      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseCPUPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseCPUPctByAppJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV4 struct {
	Address  string                                         `json:"address,nullable"`
	ASN      int64                                          `json:"asn,nullable"`
	Aso      string                                         `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseDeviceIPV4Location `json:"location"`
	Netmask  string                                         `json:"netmask,nullable"`
	Version  string                                         `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseDeviceIPV4JSON     `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseDeviceIPV4]
type deviceFleetStatusGetResponseDeviceIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV4Location struct {
	City       string                                             `json:"city,nullable"`
	CountryISO string                                             `json:"country_iso,nullable"`
	StateISO   string                                             `json:"state_iso,nullable"`
	Zip        string                                             `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseDeviceIPV4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV4LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseDeviceIPV4Location]
type deviceFleetStatusGetResponseDeviceIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV6 struct {
	Address  string                                         `json:"address,nullable"`
	ASN      int64                                          `json:"asn,nullable"`
	Aso      string                                         `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseDeviceIPV6Location `json:"location"`
	Netmask  string                                         `json:"netmask,nullable"`
	Version  string                                         `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseDeviceIPV6JSON     `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseDeviceIPV6]
type deviceFleetStatusGetResponseDeviceIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseDeviceIPV6Location struct {
	City       string                                             `json:"city,nullable"`
	CountryISO string                                             `json:"country_iso,nullable"`
	StateISO   string                                             `json:"state_iso,nullable"`
	Zip        string                                             `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseDeviceIPV6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseDeviceIPV6LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseDeviceIPV6Location]
type deviceFleetStatusGetResponseDeviceIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseDeviceIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseDeviceIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV4 struct {
	Address  string                                          `json:"address,nullable"`
	ASN      int64                                           `json:"asn,nullable"`
	Aso      string                                          `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseGatewayIPV4Location `json:"location"`
	Netmask  string                                          `json:"netmask,nullable"`
	Version  string                                          `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseGatewayIPV4JSON     `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseGatewayIPV4]
type deviceFleetStatusGetResponseGatewayIPV4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV4Location struct {
	City       string                                              `json:"city,nullable"`
	CountryISO string                                              `json:"country_iso,nullable"`
	StateISO   string                                              `json:"state_iso,nullable"`
	Zip        string                                              `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseGatewayIPV4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV4LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseGatewayIPV4Location]
type deviceFleetStatusGetResponseGatewayIPV4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV6 struct {
	Address  string                                          `json:"address,nullable"`
	ASN      int64                                           `json:"asn,nullable"`
	Aso      string                                          `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseGatewayIPV6Location `json:"location"`
	Netmask  string                                          `json:"netmask,nullable"`
	Version  string                                          `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseGatewayIPV6JSON     `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseGatewayIPV6]
type deviceFleetStatusGetResponseGatewayIPV6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseGatewayIPV6Location struct {
	City       string                                              `json:"city,nullable"`
	CountryISO string                                              `json:"country_iso,nullable"`
	StateISO   string                                              `json:"state_iso,nullable"`
	Zip        string                                              `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseGatewayIPV6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseGatewayIPV6LocationJSON contains the JSON metadata
// for the struct [DeviceFleetStatusGetResponseGatewayIPV6Location]
type deviceFleetStatusGetResponseGatewayIPV6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseGatewayIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseGatewayIPV6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV4 struct {
	Address  string                                      `json:"address,nullable"`
	ASN      int64                                       `json:"asn,nullable"`
	Aso      string                                      `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseISPIPV4Location `json:"location"`
	Netmask  string                                      `json:"netmask,nullable"`
	Version  string                                      `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseIspipv4JSON     `json:"-"`
}

// deviceFleetStatusGetResponseIspipv4JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseISPIPV4]
type deviceFleetStatusGetResponseIspipv4JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv4JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV4Location struct {
	City       string                                          `json:"city,nullable"`
	CountryISO string                                          `json:"country_iso,nullable"`
	StateISO   string                                          `json:"state_iso,nullable"`
	Zip        string                                          `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseIspipv4LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv4LocationJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseISPIPV4Location]
type deviceFleetStatusGetResponseIspipv4LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV4Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv4LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV6 struct {
	Address  string                                      `json:"address,nullable"`
	ASN      int64                                       `json:"asn,nullable"`
	Aso      string                                      `json:"aso,nullable"`
	Location DeviceFleetStatusGetResponseISPIPV6Location `json:"location"`
	Netmask  string                                      `json:"netmask,nullable"`
	Version  string                                      `json:"version,nullable"`
	JSON     deviceFleetStatusGetResponseIspipv6JSON     `json:"-"`
}

// deviceFleetStatusGetResponseIspipv6JSON contains the JSON metadata for the
// struct [DeviceFleetStatusGetResponseISPIPV6]
type deviceFleetStatusGetResponseIspipv6JSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv6JSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseISPIPV6Location struct {
	City       string                                          `json:"city,nullable"`
	CountryISO string                                          `json:"country_iso,nullable"`
	StateISO   string                                          `json:"state_iso,nullable"`
	Zip        string                                          `json:"zip,nullable"`
	JSON       deviceFleetStatusGetResponseIspipv6LocationJSON `json:"-"`
}

// deviceFleetStatusGetResponseIspipv6LocationJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseISPIPV6Location]
type deviceFleetStatusGetResponseIspipv6LocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseISPIPV6Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseIspipv6LocationJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetResponseRamUsedPctByApp struct {
	Name       string                                          `json:"name"`
	RamUsedPct float64                                         `json:"ram_used_pct"`
	JSON       deviceFleetStatusGetResponseRamUsedPctByAppJSON `json:"-"`
}

// deviceFleetStatusGetResponseRamUsedPctByAppJSON contains the JSON metadata for
// the struct [DeviceFleetStatusGetResponseRamUsedPctByApp]
type deviceFleetStatusGetResponseRamUsedPctByAppJSON struct {
	Name        apijson.Field
	RamUsedPct  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceFleetStatusGetResponseRamUsedPctByApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceFleetStatusGetResponseRamUsedPctByAppJSON) RawJSON() string {
	return r.raw
}

type DeviceFleetStatusGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
	// List of data centers to filter results
	Colo param.Field[string] `query:"colo"`
	// Number of minutes before current time
	TimeNow param.Field[string] `query:"time_now"`
}

// URLQuery serializes [DeviceFleetStatusGetParams]'s query parameters as
// `url.Values`.
func (r DeviceFleetStatusGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
