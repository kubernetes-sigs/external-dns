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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DevicePostureIntegrationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePostureIntegrationService] method instead.
type DevicePostureIntegrationService struct {
	Options []option.RequestOption
}

// NewDevicePostureIntegrationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePostureIntegrationService(opts ...option.RequestOption) (r *DevicePostureIntegrationService) {
	r = &DevicePostureIntegrationService{}
	r.Options = opts
	return
}

// Create a new device posture integration.
func (r *DevicePostureIntegrationService) New(ctx context.Context, params DevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *Integration, err error) {
	var env DevicePostureIntegrationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) List(ctx context.Context, query DevicePostureIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Integration], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", query.AccountID)
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

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) ListAutoPaging(ctx context.Context, query DevicePostureIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Integration] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured device posture integration.
func (r *DevicePostureIntegrationService) Delete(ctx context.Context, integrationID string, body DevicePostureIntegrationDeleteParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env DevicePostureIntegrationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", body.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device posture integration.
func (r *DevicePostureIntegrationService) Edit(ctx context.Context, integrationID string, params DevicePostureIntegrationEditParams, opts ...option.RequestOption) (res *Integration, err error) {
	var env DevicePostureIntegrationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", params.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device posture integration.
func (r *DevicePostureIntegrationService) Get(ctx context.Context, integrationID string, query DevicePostureIntegrationGetParams, opts ...option.RequestOption) (res *Integration, err error) {
	var env DevicePostureIntegrationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", query.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Integration struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config IntegrationConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type IntegrationType `json:"type"`
	JSON integrationJSON `json:"-"`
}

// integrationJSON contains the JSON metadata for the struct [Integration]
type integrationJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Integration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing third-party integration information.
type IntegrationConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                `json:"client_id,required"`
	JSON     integrationConfigJSON `json:"-"`
}

// integrationConfigJSON contains the JSON metadata for the struct
// [IntegrationConfig]
type integrationConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device posture integration.
type IntegrationType string

const (
	IntegrationTypeWorkspaceOne   IntegrationType = "workspace_one"
	IntegrationTypeCrowdstrikeS2s IntegrationType = "crowdstrike_s2s"
	IntegrationTypeUptycs         IntegrationType = "uptycs"
	IntegrationTypeIntune         IntegrationType = "intune"
	IntegrationTypeKolide         IntegrationType = "kolide"
	IntegrationTypeTaniumS2s      IntegrationType = "tanium_s2s"
	IntegrationTypeSentineloneS2s IntegrationType = "sentinelone_s2s"
	IntegrationTypeCustomS2s      IntegrationType = "custom_s2s"
)

func (r IntegrationType) IsKnown() bool {
	switch r {
	case IntegrationTypeWorkspaceOne, IntegrationTypeCrowdstrikeS2s, IntegrationTypeUptycs, IntegrationTypeIntune, IntegrationTypeKolide, IntegrationTypeTaniumS2s, IntegrationTypeSentineloneS2s, IntegrationTypeCustomS2s:
		return true
	}
	return false
}

type DevicePostureIntegrationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationNewParamsConfigUnion] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationNewParamsType] `json:"type,required"`
}

func (r DevicePostureIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationNewParamsConfig struct {
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id"`
}

func (r DevicePostureIntegrationNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfig) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesCustomS2sConfigRequest],
// [DevicePostureIntegrationNewParamsConfig].
type DevicePostureIntegrationNewParamsConfigUnion interface {
	implementsDevicePostureIntegrationNewParamsConfigUnion()
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest struct {
	// The Tanium API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Tanium client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesCustomS2sConfigRequest struct {
	// This id will be passed in the `CF-Access-Client-ID` header when hitting the
	// `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// This secret will be passed in the `CF-Access-Client-Secret` header when hitting
	// the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The Custom Device Posture Integration API URL.
	APIURL param.Field[string] `json:"api_url,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCustomS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCustomS2sConfigRequest) implementsDevicePostureIntegrationNewParamsConfigUnion() {
}

// The type of device posture integration.
type DevicePostureIntegrationNewParamsType string

const (
	DevicePostureIntegrationNewParamsTypeWorkspaceOne   DevicePostureIntegrationNewParamsType = "workspace_one"
	DevicePostureIntegrationNewParamsTypeCrowdstrikeS2s DevicePostureIntegrationNewParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationNewParamsTypeUptycs         DevicePostureIntegrationNewParamsType = "uptycs"
	DevicePostureIntegrationNewParamsTypeIntune         DevicePostureIntegrationNewParamsType = "intune"
	DevicePostureIntegrationNewParamsTypeKolide         DevicePostureIntegrationNewParamsType = "kolide"
	DevicePostureIntegrationNewParamsTypeTaniumS2s      DevicePostureIntegrationNewParamsType = "tanium_s2s"
	DevicePostureIntegrationNewParamsTypeSentineloneS2s DevicePostureIntegrationNewParamsType = "sentinelone_s2s"
	DevicePostureIntegrationNewParamsTypeCustomS2s      DevicePostureIntegrationNewParamsType = "custom_s2s"
)

func (r DevicePostureIntegrationNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationNewParamsTypeWorkspaceOne, DevicePostureIntegrationNewParamsTypeCrowdstrikeS2s, DevicePostureIntegrationNewParamsTypeUptycs, DevicePostureIntegrationNewParamsTypeIntune, DevicePostureIntegrationNewParamsTypeKolide, DevicePostureIntegrationNewParamsTypeTaniumS2s, DevicePostureIntegrationNewParamsTypeSentineloneS2s, DevicePostureIntegrationNewParamsTypeCustomS2s:
		return true
	}
	return false
}

type DevicePostureIntegrationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Integration           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationNewResponseEnvelope]
type devicePostureIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationNewResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationNewResponseEnvelopeSuccessTrue DevicePostureIntegrationNewResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureIntegrationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePostureIntegrationDeleteResponseEnvelope]
type devicePostureIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue DevicePostureIntegrationDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationEditParamsConfigUnion] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationEditParamsType] `json:"type"`
}

func (r DevicePostureIntegrationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationEditParamsConfig struct {
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id"`
}

func (r DevicePostureIntegrationEditParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfig) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesCustomS2sConfigRequest],
// [DevicePostureIntegrationEditParamsConfig].
type DevicePostureIntegrationEditParamsConfigUnion interface {
	implementsDevicePostureIntegrationEditParamsConfigUnion()
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest struct {
	// The Tanium API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Tanium client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesCustomS2sConfigRequest struct {
	// This id will be passed in the `CF-Access-Client-ID` header when hitting the
	// `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// This secret will be passed in the `CF-Access-Client-Secret` header when hitting
	// the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The Custom Device Posture Integration API URL.
	APIURL param.Field[string] `json:"api_url,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCustomS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCustomS2sConfigRequest) implementsDevicePostureIntegrationEditParamsConfigUnion() {
}

// The type of device posture integration.
type DevicePostureIntegrationEditParamsType string

const (
	DevicePostureIntegrationEditParamsTypeWorkspaceOne   DevicePostureIntegrationEditParamsType = "workspace_one"
	DevicePostureIntegrationEditParamsTypeCrowdstrikeS2s DevicePostureIntegrationEditParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationEditParamsTypeUptycs         DevicePostureIntegrationEditParamsType = "uptycs"
	DevicePostureIntegrationEditParamsTypeIntune         DevicePostureIntegrationEditParamsType = "intune"
	DevicePostureIntegrationEditParamsTypeKolide         DevicePostureIntegrationEditParamsType = "kolide"
	DevicePostureIntegrationEditParamsTypeTaniumS2s      DevicePostureIntegrationEditParamsType = "tanium_s2s"
	DevicePostureIntegrationEditParamsTypeSentineloneS2s DevicePostureIntegrationEditParamsType = "sentinelone_s2s"
	DevicePostureIntegrationEditParamsTypeCustomS2s      DevicePostureIntegrationEditParamsType = "custom_s2s"
)

func (r DevicePostureIntegrationEditParamsType) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationEditParamsTypeWorkspaceOne, DevicePostureIntegrationEditParamsTypeCrowdstrikeS2s, DevicePostureIntegrationEditParamsTypeUptycs, DevicePostureIntegrationEditParamsTypeIntune, DevicePostureIntegrationEditParamsTypeKolide, DevicePostureIntegrationEditParamsTypeTaniumS2s, DevicePostureIntegrationEditParamsTypeSentineloneS2s, DevicePostureIntegrationEditParamsTypeCustomS2s:
		return true
	}
	return false
}

type DevicePostureIntegrationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Integration           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationEditResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationEditResponseEnvelope]
type devicePostureIntegrationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationEditResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationEditResponseEnvelopeSuccessTrue DevicePostureIntegrationEditResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureIntegrationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Integration           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationGetResponseEnvelope]
type devicePostureIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationGetResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationGetResponseEnvelopeSuccessTrue DevicePostureIntegrationGetResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
