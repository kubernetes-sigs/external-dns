// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DevicePostureService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePostureService] method instead.
type DevicePostureService struct {
	Options      []option.RequestOption
	Integrations *DevicePostureIntegrationService
}

// NewDevicePostureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePostureService(opts ...option.RequestOption) (r *DevicePostureService) {
	r = &DevicePostureService{}
	r.Options = opts
	r.Integrations = NewDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	var env DevicePostureNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	var env DevicePostureUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture", query.AccountID)
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

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) ListAutoPaging(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureRule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, body DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	var env DevicePostureDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	var env DevicePostureGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CarbonblackInput = string

type CarbonblackInputParam = string

type ClientCertificateInput struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate.
	Cn   string                     `json:"cn,required"`
	JSON clientCertificateInputJSON `json:"-"`
}

// clientCertificateInputJSON contains the JSON metadata for the struct
// [ClientCertificateInput]
type clientCertificateInputJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ClientCertificateInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateInputJSON) RawJSON() string {
	return r.raw
}

func (r ClientCertificateInput) implementsDeviceInput() {}

type ClientCertificateInputParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate.
	Cn param.Field[string] `json:"cn,required"`
}

func (r ClientCertificateInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ClientCertificateInputParam) implementsDeviceInputUnionParam() {}

type CrowdstrikeInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator.
	Operator CrowdstrikeInputOperator `json:"operator"`
	// Os Version.
	OS string `json:"os"`
	// Overall.
	Overall string `json:"overall"`
	// SensorConfig.
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State CrowdstrikeInputState `json:"state"`
	// Version.
	Version string `json:"version"`
	// Version Operator.
	VersionOperator CrowdstrikeInputVersionOperator `json:"versionOperator"`
	JSON            crowdstrikeInputJSON            `json:"-"`
}

// crowdstrikeInputJSON contains the JSON metadata for the struct
// [CrowdstrikeInput]
type crowdstrikeInputJSON struct {
	ConnectionID    apijson.Field
	LastSeen        apijson.Field
	Operator        apijson.Field
	OS              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	State           apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CrowdstrikeInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crowdstrikeInputJSON) RawJSON() string {
	return r.raw
}

func (r CrowdstrikeInput) implementsDeviceInput() {}

// Operator.
type CrowdstrikeInputOperator string

const (
	CrowdstrikeInputOperatorLess            CrowdstrikeInputOperator = "<"
	CrowdstrikeInputOperatorLessOrEquals    CrowdstrikeInputOperator = "<="
	CrowdstrikeInputOperatorGreater         CrowdstrikeInputOperator = ">"
	CrowdstrikeInputOperatorGreaterOrEquals CrowdstrikeInputOperator = ">="
	CrowdstrikeInputOperatorEquals          CrowdstrikeInputOperator = "=="
)

func (r CrowdstrikeInputOperator) IsKnown() bool {
	switch r {
	case CrowdstrikeInputOperatorLess, CrowdstrikeInputOperatorLessOrEquals, CrowdstrikeInputOperatorGreater, CrowdstrikeInputOperatorGreaterOrEquals, CrowdstrikeInputOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type CrowdstrikeInputState string

const (
	CrowdstrikeInputStateOnline  CrowdstrikeInputState = "online"
	CrowdstrikeInputStateOffline CrowdstrikeInputState = "offline"
	CrowdstrikeInputStateUnknown CrowdstrikeInputState = "unknown"
)

func (r CrowdstrikeInputState) IsKnown() bool {
	switch r {
	case CrowdstrikeInputStateOnline, CrowdstrikeInputStateOffline, CrowdstrikeInputStateUnknown:
		return true
	}
	return false
}

// Version Operator.
type CrowdstrikeInputVersionOperator string

const (
	CrowdstrikeInputVersionOperatorLess            CrowdstrikeInputVersionOperator = "<"
	CrowdstrikeInputVersionOperatorLessOrEquals    CrowdstrikeInputVersionOperator = "<="
	CrowdstrikeInputVersionOperatorGreater         CrowdstrikeInputVersionOperator = ">"
	CrowdstrikeInputVersionOperatorGreaterOrEquals CrowdstrikeInputVersionOperator = ">="
	CrowdstrikeInputVersionOperatorEquals          CrowdstrikeInputVersionOperator = "=="
)

func (r CrowdstrikeInputVersionOperator) IsKnown() bool {
	switch r {
	case CrowdstrikeInputVersionOperatorLess, CrowdstrikeInputVersionOperatorLessOrEquals, CrowdstrikeInputVersionOperatorGreater, CrowdstrikeInputVersionOperatorGreaterOrEquals, CrowdstrikeInputVersionOperatorEquals:
		return true
	}
	return false
}

type CrowdstrikeInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator.
	Operator param.Field[CrowdstrikeInputOperator] `json:"operator"`
	// Os Version.
	OS param.Field[string] `json:"os"`
	// Overall.
	Overall param.Field[string] `json:"overall"`
	// SensorConfig.
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[CrowdstrikeInputState] `json:"state"`
	// Version.
	Version param.Field[string] `json:"version"`
	// Version Operator.
	VersionOperator param.Field[CrowdstrikeInputVersionOperator] `json:"versionOperator"`
}

func (r CrowdstrikeInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrowdstrikeInputParam) implementsDeviceInputUnionParam() {}

// The value to be checked against.
type DeviceInput struct {
	// List ID.
	ID string `json:"id"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey bool `json:"check_private_key"`
	// This field can have the runtime type of [[]CarbonblackInput].
	CheckDisks interface{} `json:"checkDisks"`
	// Common Name that is protected by the certificate.
	Cn string `json:"cn"`
	// Compliance Status.
	ComplianceStatus DeviceInputComplianceStatus `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID string `json:"connection_id"`
	// Count Operator.
	CountOperator DeviceInputCountOperator `json:"countOperator"`
	// Domain.
	Domain string `json:"domain"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Enabled.
	Enabled bool `json:"enabled"`
	// Whether or not file exists.
	Exists bool `json:"exists"`
	// This field can have the runtime type of
	// [[]DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage].
	ExtendedKeyUsage interface{} `json:"extended_key_usage"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// This field can have the runtime type of
	// [DeviceInputTeamsDevicesClientCertificateV2InputRequestLocations].
	Locations interface{} `json:"locations"`
	// Network status of device.
	NetworkStatus DeviceInputNetworkStatus `json:"network_status"`
	// Operating system.
	OperatingSystem DeviceInputOperatingSystem `json:"operating_system"`
	// Agent operational state.
	OperationalState DeviceInputOperationalState `json:"operational_state"`
	// Operator.
	Operator DeviceInputOperator `json:"operator"`
	// Os Version.
	OS string `json:"os"`
	// Operating System Distribution Name (linux only).
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only).
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only).
	OSVersionExtra string `json:"os_version_extra"`
	// Overall.
	Overall string `json:"overall"`
	// File path.
	Path string `json:"path"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DeviceInputRiskLevel `json:"risk_level"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score float64 `json:"score"`
	// Score Operator.
	ScoreOperator DeviceInputScoreOperator `json:"scoreOperator"`
	// SensorConfig.
	SensorConfig string `json:"sensor_config"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DeviceInputState `json:"state"`
	// This field can have the runtime type of [[]string].
	SubjectAlternativeNames interface{} `json:"subject_alternative_names"`
	// Signing certificate thumbprint.
	Thumbprint string `json:"thumbprint"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// Version of OS.
	Version string `json:"version"`
	// Version Operator.
	VersionOperator DeviceInputVersionOperator `json:"versionOperator"`
	JSON            deviceInputJSON            `json:"-"`
	union           DeviceInputUnion
}

// deviceInputJSON contains the JSON metadata for the struct [DeviceInput]
type deviceInputJSON struct {
	ID                      apijson.Field
	ActiveThreats           apijson.Field
	CertificateID           apijson.Field
	CheckPrivateKey         apijson.Field
	CheckDisks              apijson.Field
	Cn                      apijson.Field
	ComplianceStatus        apijson.Field
	ConnectionID            apijson.Field
	CountOperator           apijson.Field
	Domain                  apijson.Field
	EidLastSeen             apijson.Field
	Enabled                 apijson.Field
	Exists                  apijson.Field
	ExtendedKeyUsage        apijson.Field
	Infected                apijson.Field
	IsActive                apijson.Field
	IssueCount              apijson.Field
	LastSeen                apijson.Field
	Locations               apijson.Field
	NetworkStatus           apijson.Field
	OperatingSystem         apijson.Field
	OperationalState        apijson.Field
	Operator                apijson.Field
	OS                      apijson.Field
	OSDistroName            apijson.Field
	OSDistroRevision        apijson.Field
	OSVersionExtra          apijson.Field
	Overall                 apijson.Field
	Path                    apijson.Field
	RequireAll              apijson.Field
	RiskLevel               apijson.Field
	Score                   apijson.Field
	ScoreOperator           apijson.Field
	SensorConfig            apijson.Field
	Sha256                  apijson.Field
	State                   apijson.Field
	SubjectAlternativeNames apijson.Field
	Thumbprint              apijson.Field
	TotalScore              apijson.Field
	Version                 apijson.Field
	VersionOperator         apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r deviceInputJSON) RawJSON() string {
	return r.raw
}

func (r *DeviceInput) UnmarshalJSON(data []byte) (err error) {
	*r = DeviceInput{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeviceInputUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [FileInput], [UniqueClientIDInput],
// [DomainJoinedInput], [OSVersionInput], [FirewallInput], [SentineloneInput],
// [DeviceInputTeamsDevicesCarbonblackInputRequest],
// [DeviceInputTeamsDevicesAccessSerialNumberListInputRequest],
// [DiskEncryptionInput], [DeviceInputTeamsDevicesApplicationInputRequest],
// [ClientCertificateInput],
// [DeviceInputTeamsDevicesClientCertificateV2InputRequest], [WorkspaceOneInput],
// [CrowdstrikeInput], [IntuneInput], [KolideInput], [TaniumInput],
// [SentineloneS2sInput], [DeviceInputTeamsDevicesCustomS2sInputRequest].
func (r DeviceInput) AsUnion() DeviceInputUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by [FileInput], [UniqueClientIDInput], [DomainJoinedInput],
// [OSVersionInput], [FirewallInput], [SentineloneInput],
// [DeviceInputTeamsDevicesCarbonblackInputRequest],
// [DeviceInputTeamsDevicesAccessSerialNumberListInputRequest],
// [DiskEncryptionInput], [DeviceInputTeamsDevicesApplicationInputRequest],
// [ClientCertificateInput],
// [DeviceInputTeamsDevicesClientCertificateV2InputRequest], [WorkspaceOneInput],
// [CrowdstrikeInput], [IntuneInput], [KolideInput], [TaniumInput],
// [SentineloneS2sInput] or [DeviceInputTeamsDevicesCustomS2sInputRequest].
type DeviceInputUnion interface {
	implementsDeviceInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceInputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FileInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UniqueClientIDInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainJoinedInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OSVersionInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SentineloneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesAccessSerialNumberListInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiskEncryptionInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ClientCertificateInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesClientCertificateV2InputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkspaceOneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CrowdstrikeInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntuneInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KolideInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TaniumInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SentineloneS2sInput{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesCustomS2sInputRequest{}),
		},
	)
}

type DeviceInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system.
	OperatingSystem DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesCarbonblackInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesCarbonblackInputRequest]
type deviceInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequest) implementsDeviceInput() {}

// Operating system.
type DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac     DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesAccessSerialNumberListInputRequest struct {
	// UUID of Access List.
	ID   string                                                        `json:"id,required"`
	JSON deviceInputTeamsDevicesAccessSerialNumberListInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesAccessSerialNumberListInputRequestJSON contains the JSON
// metadata for the struct
// [DeviceInputTeamsDevicesAccessSerialNumberListInputRequest]
type deviceInputTeamsDevicesAccessSerialNumberListInputRequestJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesAccessSerialNumberListInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesAccessSerialNumberListInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesAccessSerialNumberListInputRequest) implementsDeviceInput() {}

type DeviceInputTeamsDevicesApplicationInputRequest struct {
	// Operating system.
	OperatingSystem DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesApplicationInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesApplicationInputRequest]
type deviceInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesApplicationInputRequest) implementsDeviceInput() {}

// Operating system.
type DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemWindows DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemLinux   DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemMac     DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemWindows, DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemLinux, DeviceInputTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesClientCertificateV2InputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey bool `json:"check_private_key,required"`
	// Operating system.
	OperatingSystem DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem `json:"operating_system,required"`
	// Certificate Common Name. This may include one or more variables in the ${ }
	// notation. Only ${serial_number} and ${hostname} are valid variables.
	Cn string `json:"cn"`
	// List of values indicating purposes for which the certificate public key can be
	// used.
	ExtendedKeyUsage []DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage `json:"extended_key_usage"`
	Locations        DeviceInputTeamsDevicesClientCertificateV2InputRequestLocations          `json:"locations"`
	// List of certificate Subject Alternative Names.
	SubjectAlternativeNames []string                                                   `json:"subject_alternative_names"`
	JSON                    deviceInputTeamsDevicesClientCertificateV2InputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesClientCertificateV2InputRequestJSON contains the JSON
// metadata for the struct [DeviceInputTeamsDevicesClientCertificateV2InputRequest]
type deviceInputTeamsDevicesClientCertificateV2InputRequestJSON struct {
	CertificateID           apijson.Field
	CheckPrivateKey         apijson.Field
	OperatingSystem         apijson.Field
	Cn                      apijson.Field
	ExtendedKeyUsage        apijson.Field
	Locations               apijson.Field
	SubjectAlternativeNames apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesClientCertificateV2InputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesClientCertificateV2InputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequest) implementsDeviceInput() {}

// Operating system.
type DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemWindows DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemLinux   DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "linux"
	DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemMac     DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemWindows, DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemLinux, DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage string

const (
	DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageClientAuth      DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage = "clientAuth"
	DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageEmailProtection DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage = "emailProtection"
)

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageClientAuth, DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageEmailProtection:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesClientCertificateV2InputRequestLocations struct {
	// List of paths to check for client certificate on linux.
	Paths []string `json:"paths"`
	// List of trust stores to check for client certificate.
	TrustStores []DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore `json:"trust_stores"`
	JSON        deviceInputTeamsDevicesClientCertificateV2InputRequestLocationsJSON         `json:"-"`
}

// deviceInputTeamsDevicesClientCertificateV2InputRequestLocationsJSON contains the
// JSON metadata for the struct
// [DeviceInputTeamsDevicesClientCertificateV2InputRequestLocations]
type deviceInputTeamsDevicesClientCertificateV2InputRequestLocationsJSON struct {
	Paths       apijson.Field
	TrustStores apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesClientCertificateV2InputRequestLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesClientCertificateV2InputRequestLocationsJSON) RawJSON() string {
	return r.raw
}

type DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore string

const (
	DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreSystem DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore = "system"
	DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreUser   DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore = "user"
)

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreSystem, DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreUser:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesCustomS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Operator.
	Operator DeviceInputTeamsDevicesCustomS2sInputRequestOperator `json:"operator,required"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score float64                                          `json:"score,required"`
	JSON  deviceInputTeamsDevicesCustomS2sInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesCustomS2sInputRequestJSON contains the JSON metadata for
// the struct [DeviceInputTeamsDevicesCustomS2sInputRequest]
type deviceInputTeamsDevicesCustomS2sInputRequestJSON struct {
	ConnectionID apijson.Field
	Operator     apijson.Field
	Score        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesCustomS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesCustomS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesCustomS2sInputRequest) implementsDeviceInput() {}

// Operator.
type DeviceInputTeamsDevicesCustomS2sInputRequestOperator string

const (
	DeviceInputTeamsDevicesCustomS2sInputRequestOperatorLess            DeviceInputTeamsDevicesCustomS2sInputRequestOperator = "<"
	DeviceInputTeamsDevicesCustomS2sInputRequestOperatorLessOrEquals    DeviceInputTeamsDevicesCustomS2sInputRequestOperator = "<="
	DeviceInputTeamsDevicesCustomS2sInputRequestOperatorGreater         DeviceInputTeamsDevicesCustomS2sInputRequestOperator = ">"
	DeviceInputTeamsDevicesCustomS2sInputRequestOperatorGreaterOrEquals DeviceInputTeamsDevicesCustomS2sInputRequestOperator = ">="
	DeviceInputTeamsDevicesCustomS2sInputRequestOperatorEquals          DeviceInputTeamsDevicesCustomS2sInputRequestOperator = "=="
)

func (r DeviceInputTeamsDevicesCustomS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesCustomS2sInputRequestOperatorLess, DeviceInputTeamsDevicesCustomS2sInputRequestOperatorLessOrEquals, DeviceInputTeamsDevicesCustomS2sInputRequestOperatorGreater, DeviceInputTeamsDevicesCustomS2sInputRequestOperatorGreaterOrEquals, DeviceInputTeamsDevicesCustomS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

// Compliance Status.
type DeviceInputComplianceStatus string

const (
	DeviceInputComplianceStatusCompliant     DeviceInputComplianceStatus = "compliant"
	DeviceInputComplianceStatusNoncompliant  DeviceInputComplianceStatus = "noncompliant"
	DeviceInputComplianceStatusUnknown       DeviceInputComplianceStatus = "unknown"
	DeviceInputComplianceStatusNotapplicable DeviceInputComplianceStatus = "notapplicable"
	DeviceInputComplianceStatusIngraceperiod DeviceInputComplianceStatus = "ingraceperiod"
	DeviceInputComplianceStatusError         DeviceInputComplianceStatus = "error"
)

func (r DeviceInputComplianceStatus) IsKnown() bool {
	switch r {
	case DeviceInputComplianceStatusCompliant, DeviceInputComplianceStatusNoncompliant, DeviceInputComplianceStatusUnknown, DeviceInputComplianceStatusNotapplicable, DeviceInputComplianceStatusIngraceperiod, DeviceInputComplianceStatusError:
		return true
	}
	return false
}

// Count Operator.
type DeviceInputCountOperator string

const (
	DeviceInputCountOperatorLess            DeviceInputCountOperator = "<"
	DeviceInputCountOperatorLessOrEquals    DeviceInputCountOperator = "<="
	DeviceInputCountOperatorGreater         DeviceInputCountOperator = ">"
	DeviceInputCountOperatorGreaterOrEquals DeviceInputCountOperator = ">="
	DeviceInputCountOperatorEquals          DeviceInputCountOperator = "=="
)

func (r DeviceInputCountOperator) IsKnown() bool {
	switch r {
	case DeviceInputCountOperatorLess, DeviceInputCountOperatorLessOrEquals, DeviceInputCountOperatorGreater, DeviceInputCountOperatorGreaterOrEquals, DeviceInputCountOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DeviceInputNetworkStatus string

const (
	DeviceInputNetworkStatusConnected     DeviceInputNetworkStatus = "connected"
	DeviceInputNetworkStatusDisconnected  DeviceInputNetworkStatus = "disconnected"
	DeviceInputNetworkStatusDisconnecting DeviceInputNetworkStatus = "disconnecting"
	DeviceInputNetworkStatusConnecting    DeviceInputNetworkStatus = "connecting"
)

func (r DeviceInputNetworkStatus) IsKnown() bool {
	switch r {
	case DeviceInputNetworkStatusConnected, DeviceInputNetworkStatusDisconnected, DeviceInputNetworkStatusDisconnecting, DeviceInputNetworkStatusConnecting:
		return true
	}
	return false
}

// Operating system.
type DeviceInputOperatingSystem string

const (
	DeviceInputOperatingSystemWindows  DeviceInputOperatingSystem = "windows"
	DeviceInputOperatingSystemLinux    DeviceInputOperatingSystem = "linux"
	DeviceInputOperatingSystemMac      DeviceInputOperatingSystem = "mac"
	DeviceInputOperatingSystemAndroid  DeviceInputOperatingSystem = "android"
	DeviceInputOperatingSystemIos      DeviceInputOperatingSystem = "ios"
	DeviceInputOperatingSystemChromeos DeviceInputOperatingSystem = "chromeos"
)

func (r DeviceInputOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputOperatingSystemWindows, DeviceInputOperatingSystemLinux, DeviceInputOperatingSystemMac, DeviceInputOperatingSystemAndroid, DeviceInputOperatingSystemIos, DeviceInputOperatingSystemChromeos:
		return true
	}
	return false
}

// Agent operational state.
type DeviceInputOperationalState string

const (
	DeviceInputOperationalStateNa                    DeviceInputOperationalState = "na"
	DeviceInputOperationalStatePartiallyDisabled     DeviceInputOperationalState = "partially_disabled"
	DeviceInputOperationalStateAutoFullyDisabled     DeviceInputOperationalState = "auto_fully_disabled"
	DeviceInputOperationalStateFullyDisabled         DeviceInputOperationalState = "fully_disabled"
	DeviceInputOperationalStateAutoPartiallyDisabled DeviceInputOperationalState = "auto_partially_disabled"
	DeviceInputOperationalStateDisabledError         DeviceInputOperationalState = "disabled_error"
	DeviceInputOperationalStateDBCorruption          DeviceInputOperationalState = "db_corruption"
)

func (r DeviceInputOperationalState) IsKnown() bool {
	switch r {
	case DeviceInputOperationalStateNa, DeviceInputOperationalStatePartiallyDisabled, DeviceInputOperationalStateAutoFullyDisabled, DeviceInputOperationalStateFullyDisabled, DeviceInputOperationalStateAutoPartiallyDisabled, DeviceInputOperationalStateDisabledError, DeviceInputOperationalStateDBCorruption:
		return true
	}
	return false
}

// Operator.
type DeviceInputOperator string

const (
	DeviceInputOperatorLess            DeviceInputOperator = "<"
	DeviceInputOperatorLessOrEquals    DeviceInputOperator = "<="
	DeviceInputOperatorGreater         DeviceInputOperator = ">"
	DeviceInputOperatorGreaterOrEquals DeviceInputOperator = ">="
	DeviceInputOperatorEquals          DeviceInputOperator = "=="
)

func (r DeviceInputOperator) IsKnown() bool {
	switch r {
	case DeviceInputOperatorLess, DeviceInputOperatorLessOrEquals, DeviceInputOperatorGreater, DeviceInputOperatorGreaterOrEquals, DeviceInputOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DeviceInputRiskLevel string

const (
	DeviceInputRiskLevelLow      DeviceInputRiskLevel = "low"
	DeviceInputRiskLevelMedium   DeviceInputRiskLevel = "medium"
	DeviceInputRiskLevelHigh     DeviceInputRiskLevel = "high"
	DeviceInputRiskLevelCritical DeviceInputRiskLevel = "critical"
)

func (r DeviceInputRiskLevel) IsKnown() bool {
	switch r {
	case DeviceInputRiskLevelLow, DeviceInputRiskLevelMedium, DeviceInputRiskLevelHigh, DeviceInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator.
type DeviceInputScoreOperator string

const (
	DeviceInputScoreOperatorLess            DeviceInputScoreOperator = "<"
	DeviceInputScoreOperatorLessOrEquals    DeviceInputScoreOperator = "<="
	DeviceInputScoreOperatorGreater         DeviceInputScoreOperator = ">"
	DeviceInputScoreOperatorGreaterOrEquals DeviceInputScoreOperator = ">="
	DeviceInputScoreOperatorEquals          DeviceInputScoreOperator = "=="
)

func (r DeviceInputScoreOperator) IsKnown() bool {
	switch r {
	case DeviceInputScoreOperatorLess, DeviceInputScoreOperatorLessOrEquals, DeviceInputScoreOperatorGreater, DeviceInputScoreOperatorGreaterOrEquals, DeviceInputScoreOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DeviceInputState string

const (
	DeviceInputStateOnline  DeviceInputState = "online"
	DeviceInputStateOffline DeviceInputState = "offline"
	DeviceInputStateUnknown DeviceInputState = "unknown"
)

func (r DeviceInputState) IsKnown() bool {
	switch r {
	case DeviceInputStateOnline, DeviceInputStateOffline, DeviceInputStateUnknown:
		return true
	}
	return false
}

// Version Operator.
type DeviceInputVersionOperator string

const (
	DeviceInputVersionOperatorLess            DeviceInputVersionOperator = "<"
	DeviceInputVersionOperatorLessOrEquals    DeviceInputVersionOperator = "<="
	DeviceInputVersionOperatorGreater         DeviceInputVersionOperator = ">"
	DeviceInputVersionOperatorGreaterOrEquals DeviceInputVersionOperator = ">="
	DeviceInputVersionOperatorEquals          DeviceInputVersionOperator = "=="
)

func (r DeviceInputVersionOperator) IsKnown() bool {
	switch r {
	case DeviceInputVersionOperatorLess, DeviceInputVersionOperatorLessOrEquals, DeviceInputVersionOperatorGreater, DeviceInputVersionOperatorGreaterOrEquals, DeviceInputVersionOperatorEquals:
		return true
	}
	return false
}

// The value to be checked against.
type DeviceInputParam struct {
	// List ID.
	ID param.Field[string] `json:"id"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey param.Field[bool]        `json:"check_private_key"`
	CheckDisks      param.Field[interface{}] `json:"checkDisks"`
	// Common Name that is protected by the certificate.
	Cn param.Field[string] `json:"cn"`
	// Compliance Status.
	ComplianceStatus param.Field[DeviceInputComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// Count Operator.
	CountOperator param.Field[DeviceInputCountOperator] `json:"countOperator"`
	// Domain.
	Domain param.Field[string] `json:"domain"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether or not file exists.
	Exists           param.Field[bool]        `json:"exists"`
	ExtendedKeyUsage param.Field[interface{}] `json:"extended_key_usage"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen  param.Field[string]      `json:"last_seen"`
	Locations param.Field[interface{}] `json:"locations"`
	// Network status of device.
	NetworkStatus param.Field[DeviceInputNetworkStatus] `json:"network_status"`
	// Operating system.
	OperatingSystem param.Field[DeviceInputOperatingSystem] `json:"operating_system"`
	// Agent operational state.
	OperationalState param.Field[DeviceInputOperationalState] `json:"operational_state"`
	// Operator.
	Operator param.Field[DeviceInputOperator] `json:"operator"`
	// Os Version.
	OS param.Field[string] `json:"os"`
	// Operating System Distribution Name (linux only).
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only).
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only).
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
	// Overall.
	Overall param.Field[string] `json:"overall"`
	// File path.
	Path param.Field[string] `json:"path"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DeviceInputRiskLevel] `json:"risk_level"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score param.Field[float64] `json:"score"`
	// Score Operator.
	ScoreOperator param.Field[DeviceInputScoreOperator] `json:"scoreOperator"`
	// SensorConfig.
	SensorConfig param.Field[string] `json:"sensor_config"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State                   param.Field[DeviceInputState] `json:"state"`
	SubjectAlternativeNames param.Field[interface{}]      `json:"subject_alternative_names"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// Version of OS.
	Version param.Field[string] `json:"version"`
	// Version Operator.
	VersionOperator param.Field[DeviceInputVersionOperator] `json:"versionOperator"`
}

func (r DeviceInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputParam) implementsDeviceInputUnionParam() {}

// The value to be checked against.
//
// Satisfied by [zero_trust.FileInputParam], [zero_trust.UniqueClientIDInputParam],
// [zero_trust.DomainJoinedInputParam], [zero_trust.OSVersionInputParam],
// [zero_trust.FirewallInputParam], [zero_trust.SentineloneInputParam],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesAccessSerialNumberListInputRequestParam],
// [zero_trust.DiskEncryptionInputParam],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequestParam],
// [zero_trust.ClientCertificateInputParam],
// [zero_trust.DeviceInputTeamsDevicesClientCertificateV2InputRequestParam],
// [zero_trust.WorkspaceOneInputParam], [zero_trust.CrowdstrikeInputParam],
// [zero_trust.IntuneInputParam], [zero_trust.KolideInputParam],
// [zero_trust.TaniumInputParam], [zero_trust.SentineloneS2sInputParam],
// [zero_trust.DeviceInputTeamsDevicesCustomS2sInputRequestParam],
// [DeviceInputParam].
type DeviceInputUnionParam interface {
	implementsDeviceInputUnionParam()
}

type DeviceInputTeamsDevicesCarbonblackInputRequestParam struct {
	// Operating system.
	OperatingSystem param.Field[DeviceInputTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) implementsDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesAccessSerialNumberListInputRequestParam struct {
	// UUID of Access List.
	ID param.Field[string] `json:"id,required"`
}

func (r DeviceInputTeamsDevicesAccessSerialNumberListInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesAccessSerialNumberListInputRequestParam) implementsDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesApplicationInputRequestParam struct {
	// Operating system.
	OperatingSystem param.Field[DeviceInputTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) implementsDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesClientCertificateV2InputRequestParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey param.Field[bool] `json:"check_private_key,required"`
	// Operating system.
	OperatingSystem param.Field[DeviceInputTeamsDevicesClientCertificateV2InputRequestOperatingSystem] `json:"operating_system,required"`
	// Certificate Common Name. This may include one or more variables in the ${ }
	// notation. Only ${serial_number} and ${hostname} are valid variables.
	Cn param.Field[string] `json:"cn"`
	// List of values indicating purposes for which the certificate public key can be
	// used.
	ExtendedKeyUsage param.Field[[]DeviceInputTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage] `json:"extended_key_usage"`
	Locations        param.Field[DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsParam]     `json:"locations"`
	// List of certificate Subject Alternative Names.
	SubjectAlternativeNames param.Field[[]string] `json:"subject_alternative_names"`
}

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestParam) implementsDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsParam struct {
	// List of paths to check for client certificate on linux.
	Paths param.Field[[]string] `json:"paths"`
	// List of trust stores to check for client certificate.
	TrustStores param.Field[[]DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore] `json:"trust_stores"`
}

func (r DeviceInputTeamsDevicesClientCertificateV2InputRequestLocationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceInputTeamsDevicesCustomS2sInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Operator.
	Operator param.Field[DeviceInputTeamsDevicesCustomS2sInputRequestOperator] `json:"operator,required"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score param.Field[float64] `json:"score,required"`
}

func (r DeviceInputTeamsDevicesCustomS2sInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesCustomS2sInputRequestParam) implementsDeviceInputUnionParam() {}

type DeviceMatch struct {
	Platform DeviceMatchPlatform `json:"platform"`
	JSON     deviceMatchJSON     `json:"-"`
}

// deviceMatchJSON contains the JSON metadata for the struct [DeviceMatch]
type deviceMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceMatchJSON) RawJSON() string {
	return r.raw
}

type DeviceMatchPlatform string

const (
	DeviceMatchPlatformWindows  DeviceMatchPlatform = "windows"
	DeviceMatchPlatformMac      DeviceMatchPlatform = "mac"
	DeviceMatchPlatformLinux    DeviceMatchPlatform = "linux"
	DeviceMatchPlatformAndroid  DeviceMatchPlatform = "android"
	DeviceMatchPlatformIos      DeviceMatchPlatform = "ios"
	DeviceMatchPlatformChromeos DeviceMatchPlatform = "chromeos"
)

func (r DeviceMatchPlatform) IsKnown() bool {
	switch r {
	case DeviceMatchPlatformWindows, DeviceMatchPlatformMac, DeviceMatchPlatformLinux, DeviceMatchPlatformAndroid, DeviceMatchPlatformIos, DeviceMatchPlatformChromeos:
		return true
	}
	return false
}

type DeviceMatchParam struct {
	Platform param.Field[DeviceMatchPlatform] `json:"platform"`
}

func (r DeviceMatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureRule struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DeviceInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DeviceMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureRuleType `json:"type"`
	JSON devicePostureRuleJSON `json:"-"`
}

// devicePostureRuleJSON contains the JSON metadata for the struct
// [DevicePostureRule]
type devicePostureRuleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRuleJSON) RawJSON() string {
	return r.raw
}

// The type of device posture rule.
type DevicePostureRuleType string

const (
	DevicePostureRuleTypeFile                DevicePostureRuleType = "file"
	DevicePostureRuleTypeApplication         DevicePostureRuleType = "application"
	DevicePostureRuleTypeTanium              DevicePostureRuleType = "tanium"
	DevicePostureRuleTypeGateway             DevicePostureRuleType = "gateway"
	DevicePostureRuleTypeWARP                DevicePostureRuleType = "warp"
	DevicePostureRuleTypeDiskEncryption      DevicePostureRuleType = "disk_encryption"
	DevicePostureRuleTypeSerialNumber        DevicePostureRuleType = "serial_number"
	DevicePostureRuleTypeSentinelone         DevicePostureRuleType = "sentinelone"
	DevicePostureRuleTypeCarbonblack         DevicePostureRuleType = "carbonblack"
	DevicePostureRuleTypeFirewall            DevicePostureRuleType = "firewall"
	DevicePostureRuleTypeOSVersion           DevicePostureRuleType = "os_version"
	DevicePostureRuleTypeDomainJoined        DevicePostureRuleType = "domain_joined"
	DevicePostureRuleTypeClientCertificate   DevicePostureRuleType = "client_certificate"
	DevicePostureRuleTypeClientCertificateV2 DevicePostureRuleType = "client_certificate_v2"
	DevicePostureRuleTypeUniqueClientID      DevicePostureRuleType = "unique_client_id"
	DevicePostureRuleTypeKolide              DevicePostureRuleType = "kolide"
	DevicePostureRuleTypeTaniumS2s           DevicePostureRuleType = "tanium_s2s"
	DevicePostureRuleTypeCrowdstrikeS2s      DevicePostureRuleType = "crowdstrike_s2s"
	DevicePostureRuleTypeIntune              DevicePostureRuleType = "intune"
	DevicePostureRuleTypeWorkspaceOne        DevicePostureRuleType = "workspace_one"
	DevicePostureRuleTypeSentineloneS2s      DevicePostureRuleType = "sentinelone_s2s"
	DevicePostureRuleTypeCustomS2s           DevicePostureRuleType = "custom_s2s"
)

func (r DevicePostureRuleType) IsKnown() bool {
	switch r {
	case DevicePostureRuleTypeFile, DevicePostureRuleTypeApplication, DevicePostureRuleTypeTanium, DevicePostureRuleTypeGateway, DevicePostureRuleTypeWARP, DevicePostureRuleTypeDiskEncryption, DevicePostureRuleTypeSerialNumber, DevicePostureRuleTypeSentinelone, DevicePostureRuleTypeCarbonblack, DevicePostureRuleTypeFirewall, DevicePostureRuleTypeOSVersion, DevicePostureRuleTypeDomainJoined, DevicePostureRuleTypeClientCertificate, DevicePostureRuleTypeClientCertificateV2, DevicePostureRuleTypeUniqueClientID, DevicePostureRuleTypeKolide, DevicePostureRuleTypeTaniumS2s, DevicePostureRuleTypeCrowdstrikeS2s, DevicePostureRuleTypeIntune, DevicePostureRuleTypeWorkspaceOne, DevicePostureRuleTypeSentineloneS2s, DevicePostureRuleTypeCustomS2s:
		return true
	}
	return false
}

type DiskEncryptionInput struct {
	// List of volume names to be checked for encryption.
	CheckDisks []CarbonblackInput `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                    `json:"requireAll"`
	JSON       diskEncryptionInputJSON `json:"-"`
}

// diskEncryptionInputJSON contains the JSON metadata for the struct
// [DiskEncryptionInput]
type diskEncryptionInputJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiskEncryptionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r diskEncryptionInputJSON) RawJSON() string {
	return r.raw
}

func (r DiskEncryptionInput) implementsDeviceInput() {}

type DiskEncryptionInputParam struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]CarbonblackInputParam] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DiskEncryptionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiskEncryptionInputParam) implementsDeviceInputUnionParam() {}

type DomainJoinedInput struct {
	// Operating System.
	OperatingSystem DomainJoinedInputOperatingSystem `json:"operating_system,required"`
	// Domain.
	Domain string                `json:"domain"`
	JSON   domainJoinedInputJSON `json:"-"`
}

// domainJoinedInputJSON contains the JSON metadata for the struct
// [DomainJoinedInput]
type domainJoinedInputJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainJoinedInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJoinedInputJSON) RawJSON() string {
	return r.raw
}

func (r DomainJoinedInput) implementsDeviceInput() {}

// Operating System.
type DomainJoinedInputOperatingSystem string

const (
	DomainJoinedInputOperatingSystemWindows DomainJoinedInputOperatingSystem = "windows"
)

func (r DomainJoinedInputOperatingSystem) IsKnown() bool {
	switch r {
	case DomainJoinedInputOperatingSystemWindows:
		return true
	}
	return false
}

type DomainJoinedInputParam struct {
	// Operating System.
	OperatingSystem param.Field[DomainJoinedInputOperatingSystem] `json:"operating_system,required"`
	// Domain.
	Domain param.Field[string] `json:"domain"`
}

func (r DomainJoinedInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DomainJoinedInputParam) implementsDeviceInputUnionParam() {}

type FileInput struct {
	// Operating system.
	OperatingSystem FileInputOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists.
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string        `json:"thumbprint"`
	JSON       fileInputJSON `json:"-"`
}

// fileInputJSON contains the JSON metadata for the struct [FileInput]
type fileInputJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FileInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileInputJSON) RawJSON() string {
	return r.raw
}

func (r FileInput) implementsDeviceInput() {}

// Operating system.
type FileInputOperatingSystem string

const (
	FileInputOperatingSystemWindows FileInputOperatingSystem = "windows"
	FileInputOperatingSystemLinux   FileInputOperatingSystem = "linux"
	FileInputOperatingSystemMac     FileInputOperatingSystem = "mac"
)

func (r FileInputOperatingSystem) IsKnown() bool {
	switch r {
	case FileInputOperatingSystemWindows, FileInputOperatingSystemLinux, FileInputOperatingSystemMac:
		return true
	}
	return false
}

type FileInputParam struct {
	// Operating system.
	OperatingSystem param.Field[FileInputOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists.
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r FileInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileInputParam) implementsDeviceInputUnionParam() {}

type FirewallInput struct {
	// Enabled.
	Enabled bool `json:"enabled,required"`
	// Operating System.
	OperatingSystem FirewallInputOperatingSystem `json:"operating_system,required"`
	JSON            firewallInputJSON            `json:"-"`
}

// firewallInputJSON contains the JSON metadata for the struct [FirewallInput]
type firewallInputJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FirewallInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallInputJSON) RawJSON() string {
	return r.raw
}

func (r FirewallInput) implementsDeviceInput() {}

// Operating System.
type FirewallInputOperatingSystem string

const (
	FirewallInputOperatingSystemWindows FirewallInputOperatingSystem = "windows"
	FirewallInputOperatingSystemMac     FirewallInputOperatingSystem = "mac"
)

func (r FirewallInputOperatingSystem) IsKnown() bool {
	switch r {
	case FirewallInputOperatingSystemWindows, FirewallInputOperatingSystemMac:
		return true
	}
	return false
}

type FirewallInputParam struct {
	// Enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System.
	OperatingSystem param.Field[FirewallInputOperatingSystem] `json:"operating_system,required"`
}

func (r FirewallInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallInputParam) implementsDeviceInputUnionParam() {}

type IntuneInput struct {
	// Compliance Status.
	ComplianceStatus IntuneInputComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string          `json:"connection_id,required"`
	JSON         intuneInputJSON `json:"-"`
}

// intuneInputJSON contains the JSON metadata for the struct [IntuneInput]
type intuneInputJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IntuneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intuneInputJSON) RawJSON() string {
	return r.raw
}

func (r IntuneInput) implementsDeviceInput() {}

// Compliance Status.
type IntuneInputComplianceStatus string

const (
	IntuneInputComplianceStatusCompliant     IntuneInputComplianceStatus = "compliant"
	IntuneInputComplianceStatusNoncompliant  IntuneInputComplianceStatus = "noncompliant"
	IntuneInputComplianceStatusUnknown       IntuneInputComplianceStatus = "unknown"
	IntuneInputComplianceStatusNotapplicable IntuneInputComplianceStatus = "notapplicable"
	IntuneInputComplianceStatusIngraceperiod IntuneInputComplianceStatus = "ingraceperiod"
	IntuneInputComplianceStatusError         IntuneInputComplianceStatus = "error"
)

func (r IntuneInputComplianceStatus) IsKnown() bool {
	switch r {
	case IntuneInputComplianceStatusCompliant, IntuneInputComplianceStatusNoncompliant, IntuneInputComplianceStatusUnknown, IntuneInputComplianceStatusNotapplicable, IntuneInputComplianceStatusIngraceperiod, IntuneInputComplianceStatusError:
		return true
	}
	return false
}

type IntuneInputParam struct {
	// Compliance Status.
	ComplianceStatus param.Field[IntuneInputComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r IntuneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IntuneInputParam) implementsDeviceInputUnionParam() {}

type KolideInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator.
	CountOperator KolideInputCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string          `json:"issue_count,required"`
	JSON       kolideInputJSON `json:"-"`
}

// kolideInputJSON contains the JSON metadata for the struct [KolideInput]
type kolideInputJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *KolideInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kolideInputJSON) RawJSON() string {
	return r.raw
}

func (r KolideInput) implementsDeviceInput() {}

// Count Operator.
type KolideInputCountOperator string

const (
	KolideInputCountOperatorLess            KolideInputCountOperator = "<"
	KolideInputCountOperatorLessOrEquals    KolideInputCountOperator = "<="
	KolideInputCountOperatorGreater         KolideInputCountOperator = ">"
	KolideInputCountOperatorGreaterOrEquals KolideInputCountOperator = ">="
	KolideInputCountOperatorEquals          KolideInputCountOperator = "=="
)

func (r KolideInputCountOperator) IsKnown() bool {
	switch r {
	case KolideInputCountOperatorLess, KolideInputCountOperatorLessOrEquals, KolideInputCountOperatorGreater, KolideInputCountOperatorGreaterOrEquals, KolideInputCountOperatorEquals:
		return true
	}
	return false
}

type KolideInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator.
	CountOperator param.Field[KolideInputCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r KolideInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KolideInputParam) implementsDeviceInputUnionParam() {}

type OSVersionInput struct {
	// Operating System.
	OperatingSystem OSVersionInputOperatingSystem `json:"operating_system,required"`
	// Operator.
	Operator OSVersionInputOperator `json:"operator,required"`
	// Version of OS.
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only).
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only).
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only).
	OSVersionExtra string             `json:"os_version_extra"`
	JSON           osVersionInputJSON `json:"-"`
}

// osVersionInputJSON contains the JSON metadata for the struct [OSVersionInput]
type osVersionInputJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OSVersionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r osVersionInputJSON) RawJSON() string {
	return r.raw
}

func (r OSVersionInput) implementsDeviceInput() {}

// Operating System.
type OSVersionInputOperatingSystem string

const (
	OSVersionInputOperatingSystemWindows OSVersionInputOperatingSystem = "windows"
)

func (r OSVersionInputOperatingSystem) IsKnown() bool {
	switch r {
	case OSVersionInputOperatingSystemWindows:
		return true
	}
	return false
}

// Operator.
type OSVersionInputOperator string

const (
	OSVersionInputOperatorLess            OSVersionInputOperator = "<"
	OSVersionInputOperatorLessOrEquals    OSVersionInputOperator = "<="
	OSVersionInputOperatorGreater         OSVersionInputOperator = ">"
	OSVersionInputOperatorGreaterOrEquals OSVersionInputOperator = ">="
	OSVersionInputOperatorEquals          OSVersionInputOperator = "=="
)

func (r OSVersionInputOperator) IsKnown() bool {
	switch r {
	case OSVersionInputOperatorLess, OSVersionInputOperatorLessOrEquals, OSVersionInputOperatorGreater, OSVersionInputOperatorGreaterOrEquals, OSVersionInputOperatorEquals:
		return true
	}
	return false
}

type OSVersionInputParam struct {
	// Operating System.
	OperatingSystem param.Field[OSVersionInputOperatingSystem] `json:"operating_system,required"`
	// Operator.
	Operator param.Field[OSVersionInputOperator] `json:"operator,required"`
	// Version of OS.
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only).
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only).
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only).
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r OSVersionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OSVersionInputParam) implementsDeviceInputUnionParam() {}

type SentineloneInput struct {
	// Operating system.
	OperatingSystem SentineloneInputOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string               `json:"thumbprint"`
	JSON       sentineloneInputJSON `json:"-"`
}

// sentineloneInputJSON contains the JSON metadata for the struct
// [SentineloneInput]
type sentineloneInputJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SentineloneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sentineloneInputJSON) RawJSON() string {
	return r.raw
}

func (r SentineloneInput) implementsDeviceInput() {}

// Operating system.
type SentineloneInputOperatingSystem string

const (
	SentineloneInputOperatingSystemWindows SentineloneInputOperatingSystem = "windows"
	SentineloneInputOperatingSystemLinux   SentineloneInputOperatingSystem = "linux"
	SentineloneInputOperatingSystemMac     SentineloneInputOperatingSystem = "mac"
)

func (r SentineloneInputOperatingSystem) IsKnown() bool {
	switch r {
	case SentineloneInputOperatingSystemWindows, SentineloneInputOperatingSystemLinux, SentineloneInputOperatingSystemMac:
		return true
	}
	return false
}

type SentineloneInputParam struct {
	// Operating system.
	OperatingSystem param.Field[SentineloneInputOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r SentineloneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SentineloneInputParam) implementsDeviceInputUnionParam() {}

type SentineloneS2sInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus SentineloneS2sInputNetworkStatus `json:"network_status"`
	// Agent operational state.
	OperationalState SentineloneS2sInputOperationalState `json:"operational_state"`
	// Operator.
	Operator SentineloneS2sInputOperator `json:"operator"`
	JSON     sentineloneS2sInputJSON     `json:"-"`
}

// sentineloneS2sInputJSON contains the JSON metadata for the struct
// [SentineloneS2sInput]
type sentineloneS2sInputJSON struct {
	ConnectionID     apijson.Field
	ActiveThreats    apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	NetworkStatus    apijson.Field
	OperationalState apijson.Field
	Operator         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SentineloneS2sInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sentineloneS2sInputJSON) RawJSON() string {
	return r.raw
}

func (r SentineloneS2sInput) implementsDeviceInput() {}

// Network status of device.
type SentineloneS2sInputNetworkStatus string

const (
	SentineloneS2sInputNetworkStatusConnected     SentineloneS2sInputNetworkStatus = "connected"
	SentineloneS2sInputNetworkStatusDisconnected  SentineloneS2sInputNetworkStatus = "disconnected"
	SentineloneS2sInputNetworkStatusDisconnecting SentineloneS2sInputNetworkStatus = "disconnecting"
	SentineloneS2sInputNetworkStatusConnecting    SentineloneS2sInputNetworkStatus = "connecting"
)

func (r SentineloneS2sInputNetworkStatus) IsKnown() bool {
	switch r {
	case SentineloneS2sInputNetworkStatusConnected, SentineloneS2sInputNetworkStatusDisconnected, SentineloneS2sInputNetworkStatusDisconnecting, SentineloneS2sInputNetworkStatusConnecting:
		return true
	}
	return false
}

// Agent operational state.
type SentineloneS2sInputOperationalState string

const (
	SentineloneS2sInputOperationalStateNa                    SentineloneS2sInputOperationalState = "na"
	SentineloneS2sInputOperationalStatePartiallyDisabled     SentineloneS2sInputOperationalState = "partially_disabled"
	SentineloneS2sInputOperationalStateAutoFullyDisabled     SentineloneS2sInputOperationalState = "auto_fully_disabled"
	SentineloneS2sInputOperationalStateFullyDisabled         SentineloneS2sInputOperationalState = "fully_disabled"
	SentineloneS2sInputOperationalStateAutoPartiallyDisabled SentineloneS2sInputOperationalState = "auto_partially_disabled"
	SentineloneS2sInputOperationalStateDisabledError         SentineloneS2sInputOperationalState = "disabled_error"
	SentineloneS2sInputOperationalStateDBCorruption          SentineloneS2sInputOperationalState = "db_corruption"
)

func (r SentineloneS2sInputOperationalState) IsKnown() bool {
	switch r {
	case SentineloneS2sInputOperationalStateNa, SentineloneS2sInputOperationalStatePartiallyDisabled, SentineloneS2sInputOperationalStateAutoFullyDisabled, SentineloneS2sInputOperationalStateFullyDisabled, SentineloneS2sInputOperationalStateAutoPartiallyDisabled, SentineloneS2sInputOperationalStateDisabledError, SentineloneS2sInputOperationalStateDBCorruption:
		return true
	}
	return false
}

// Operator.
type SentineloneS2sInputOperator string

const (
	SentineloneS2sInputOperatorLess            SentineloneS2sInputOperator = "<"
	SentineloneS2sInputOperatorLessOrEquals    SentineloneS2sInputOperator = "<="
	SentineloneS2sInputOperatorGreater         SentineloneS2sInputOperator = ">"
	SentineloneS2sInputOperatorGreaterOrEquals SentineloneS2sInputOperator = ">="
	SentineloneS2sInputOperatorEquals          SentineloneS2sInputOperator = "=="
)

func (r SentineloneS2sInputOperator) IsKnown() bool {
	switch r {
	case SentineloneS2sInputOperatorLess, SentineloneS2sInputOperatorLessOrEquals, SentineloneS2sInputOperatorGreater, SentineloneS2sInputOperatorGreaterOrEquals, SentineloneS2sInputOperatorEquals:
		return true
	}
	return false
}

type SentineloneS2sInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[SentineloneS2sInputNetworkStatus] `json:"network_status"`
	// Agent operational state.
	OperationalState param.Field[SentineloneS2sInputOperationalState] `json:"operational_state"`
	// Operator.
	Operator param.Field[SentineloneS2sInputOperator] `json:"operator"`
}

func (r SentineloneS2sInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SentineloneS2sInputParam) implementsDeviceInputUnionParam() {}

type TaniumInput struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator TaniumInputOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel TaniumInputRiskLevel `json:"risk_level"`
	// Score Operator.
	ScoreOperator TaniumInputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64         `json:"total_score"`
	JSON       taniumInputJSON `json:"-"`
}

// taniumInputJSON contains the JSON metadata for the struct [TaniumInput]
type taniumInputJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaniumInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taniumInputJSON) RawJSON() string {
	return r.raw
}

func (r TaniumInput) implementsDeviceInput() {}

// Operator to evaluate risk_level or eid_last_seen.
type TaniumInputOperator string

const (
	TaniumInputOperatorLess            TaniumInputOperator = "<"
	TaniumInputOperatorLessOrEquals    TaniumInputOperator = "<="
	TaniumInputOperatorGreater         TaniumInputOperator = ">"
	TaniumInputOperatorGreaterOrEquals TaniumInputOperator = ">="
	TaniumInputOperatorEquals          TaniumInputOperator = "=="
)

func (r TaniumInputOperator) IsKnown() bool {
	switch r {
	case TaniumInputOperatorLess, TaniumInputOperatorLessOrEquals, TaniumInputOperatorGreater, TaniumInputOperatorGreaterOrEquals, TaniumInputOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type TaniumInputRiskLevel string

const (
	TaniumInputRiskLevelLow      TaniumInputRiskLevel = "low"
	TaniumInputRiskLevelMedium   TaniumInputRiskLevel = "medium"
	TaniumInputRiskLevelHigh     TaniumInputRiskLevel = "high"
	TaniumInputRiskLevelCritical TaniumInputRiskLevel = "critical"
)

func (r TaniumInputRiskLevel) IsKnown() bool {
	switch r {
	case TaniumInputRiskLevelLow, TaniumInputRiskLevelMedium, TaniumInputRiskLevelHigh, TaniumInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator.
type TaniumInputScoreOperator string

const (
	TaniumInputScoreOperatorLess            TaniumInputScoreOperator = "<"
	TaniumInputScoreOperatorLessOrEquals    TaniumInputScoreOperator = "<="
	TaniumInputScoreOperatorGreater         TaniumInputScoreOperator = ">"
	TaniumInputScoreOperatorGreaterOrEquals TaniumInputScoreOperator = ">="
	TaniumInputScoreOperatorEquals          TaniumInputScoreOperator = "=="
)

func (r TaniumInputScoreOperator) IsKnown() bool {
	switch r {
	case TaniumInputScoreOperatorLess, TaniumInputScoreOperatorLessOrEquals, TaniumInputScoreOperatorGreater, TaniumInputScoreOperatorGreaterOrEquals, TaniumInputScoreOperatorEquals:
		return true
	}
	return false
}

type TaniumInputParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[TaniumInputOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[TaniumInputRiskLevel] `json:"risk_level"`
	// Score Operator.
	ScoreOperator param.Field[TaniumInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r TaniumInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TaniumInputParam) implementsDeviceInputUnionParam() {}

type UniqueClientIDInput struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System.
	OperatingSystem UniqueClientIDInputOperatingSystem `json:"operating_system,required"`
	JSON            uniqueClientIDInputJSON            `json:"-"`
}

// uniqueClientIDInputJSON contains the JSON metadata for the struct
// [UniqueClientIDInput]
type uniqueClientIDInputJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UniqueClientIDInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uniqueClientIDInputJSON) RawJSON() string {
	return r.raw
}

func (r UniqueClientIDInput) implementsDeviceInput() {}

// Operating System.
type UniqueClientIDInputOperatingSystem string

const (
	UniqueClientIDInputOperatingSystemAndroid  UniqueClientIDInputOperatingSystem = "android"
	UniqueClientIDInputOperatingSystemIos      UniqueClientIDInputOperatingSystem = "ios"
	UniqueClientIDInputOperatingSystemChromeos UniqueClientIDInputOperatingSystem = "chromeos"
)

func (r UniqueClientIDInputOperatingSystem) IsKnown() bool {
	switch r {
	case UniqueClientIDInputOperatingSystemAndroid, UniqueClientIDInputOperatingSystemIos, UniqueClientIDInputOperatingSystemChromeos:
		return true
	}
	return false
}

type UniqueClientIDInputParam struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System.
	OperatingSystem param.Field[UniqueClientIDInputOperatingSystem] `json:"operating_system,required"`
}

func (r UniqueClientIDInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UniqueClientIDInputParam) implementsDeviceInputUnionParam() {}

type WorkspaceOneInput struct {
	// Compliance Status.
	ComplianceStatus WorkspaceOneInputComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                `json:"connection_id,required"`
	JSON         workspaceOneInputJSON `json:"-"`
}

// workspaceOneInputJSON contains the JSON metadata for the struct
// [WorkspaceOneInput]
type workspaceOneInputJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkspaceOneInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceOneInputJSON) RawJSON() string {
	return r.raw
}

func (r WorkspaceOneInput) implementsDeviceInput() {}

// Compliance Status.
type WorkspaceOneInputComplianceStatus string

const (
	WorkspaceOneInputComplianceStatusCompliant    WorkspaceOneInputComplianceStatus = "compliant"
	WorkspaceOneInputComplianceStatusNoncompliant WorkspaceOneInputComplianceStatus = "noncompliant"
	WorkspaceOneInputComplianceStatusUnknown      WorkspaceOneInputComplianceStatus = "unknown"
)

func (r WorkspaceOneInputComplianceStatus) IsKnown() bool {
	switch r {
	case WorkspaceOneInputComplianceStatusCompliant, WorkspaceOneInputComplianceStatusNoncompliant, WorkspaceOneInputComplianceStatusUnknown:
		return true
	}
	return false
}

type WorkspaceOneInputParam struct {
	// Compliance Status.
	ComplianceStatus param.Field[WorkspaceOneInputComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r WorkspaceOneInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkspaceOneInputParam) implementsDeviceInputUnionParam() {}

type DevicePostureDeleteResponse struct {
	// API UUID.
	ID   string                          `json:"id"`
	JSON devicePostureDeleteResponseJSON `json:"-"`
}

// devicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [DevicePostureDeleteResponse]
type devicePostureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureNewParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureNewParamsType string

const (
	DevicePostureNewParamsTypeFile                DevicePostureNewParamsType = "file"
	DevicePostureNewParamsTypeApplication         DevicePostureNewParamsType = "application"
	DevicePostureNewParamsTypeTanium              DevicePostureNewParamsType = "tanium"
	DevicePostureNewParamsTypeGateway             DevicePostureNewParamsType = "gateway"
	DevicePostureNewParamsTypeWARP                DevicePostureNewParamsType = "warp"
	DevicePostureNewParamsTypeDiskEncryption      DevicePostureNewParamsType = "disk_encryption"
	DevicePostureNewParamsTypeSerialNumber        DevicePostureNewParamsType = "serial_number"
	DevicePostureNewParamsTypeSentinelone         DevicePostureNewParamsType = "sentinelone"
	DevicePostureNewParamsTypeCarbonblack         DevicePostureNewParamsType = "carbonblack"
	DevicePostureNewParamsTypeFirewall            DevicePostureNewParamsType = "firewall"
	DevicePostureNewParamsTypeOSVersion           DevicePostureNewParamsType = "os_version"
	DevicePostureNewParamsTypeDomainJoined        DevicePostureNewParamsType = "domain_joined"
	DevicePostureNewParamsTypeClientCertificate   DevicePostureNewParamsType = "client_certificate"
	DevicePostureNewParamsTypeClientCertificateV2 DevicePostureNewParamsType = "client_certificate_v2"
	DevicePostureNewParamsTypeUniqueClientID      DevicePostureNewParamsType = "unique_client_id"
	DevicePostureNewParamsTypeKolide              DevicePostureNewParamsType = "kolide"
	DevicePostureNewParamsTypeTaniumS2s           DevicePostureNewParamsType = "tanium_s2s"
	DevicePostureNewParamsTypeCrowdstrikeS2s      DevicePostureNewParamsType = "crowdstrike_s2s"
	DevicePostureNewParamsTypeIntune              DevicePostureNewParamsType = "intune"
	DevicePostureNewParamsTypeWorkspaceOne        DevicePostureNewParamsType = "workspace_one"
	DevicePostureNewParamsTypeSentineloneS2s      DevicePostureNewParamsType = "sentinelone_s2s"
	DevicePostureNewParamsTypeCustomS2s           DevicePostureNewParamsType = "custom_s2s"
)

func (r DevicePostureNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsTypeFile, DevicePostureNewParamsTypeApplication, DevicePostureNewParamsTypeTanium, DevicePostureNewParamsTypeGateway, DevicePostureNewParamsTypeWARP, DevicePostureNewParamsTypeDiskEncryption, DevicePostureNewParamsTypeSerialNumber, DevicePostureNewParamsTypeSentinelone, DevicePostureNewParamsTypeCarbonblack, DevicePostureNewParamsTypeFirewall, DevicePostureNewParamsTypeOSVersion, DevicePostureNewParamsTypeDomainJoined, DevicePostureNewParamsTypeClientCertificate, DevicePostureNewParamsTypeClientCertificateV2, DevicePostureNewParamsTypeUniqueClientID, DevicePostureNewParamsTypeKolide, DevicePostureNewParamsTypeTaniumS2s, DevicePostureNewParamsTypeCrowdstrikeS2s, DevicePostureNewParamsTypeIntune, DevicePostureNewParamsTypeWorkspaceOne, DevicePostureNewParamsTypeSentineloneS2s, DevicePostureNewParamsTypeCustomS2s:
		return true
	}
	return false
}

type DevicePostureNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureNewResponseEnvelope]
type devicePostureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

func (r DevicePostureNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureUpdateParamsType string

const (
	DevicePostureUpdateParamsTypeFile                DevicePostureUpdateParamsType = "file"
	DevicePostureUpdateParamsTypeApplication         DevicePostureUpdateParamsType = "application"
	DevicePostureUpdateParamsTypeTanium              DevicePostureUpdateParamsType = "tanium"
	DevicePostureUpdateParamsTypeGateway             DevicePostureUpdateParamsType = "gateway"
	DevicePostureUpdateParamsTypeWARP                DevicePostureUpdateParamsType = "warp"
	DevicePostureUpdateParamsTypeDiskEncryption      DevicePostureUpdateParamsType = "disk_encryption"
	DevicePostureUpdateParamsTypeSerialNumber        DevicePostureUpdateParamsType = "serial_number"
	DevicePostureUpdateParamsTypeSentinelone         DevicePostureUpdateParamsType = "sentinelone"
	DevicePostureUpdateParamsTypeCarbonblack         DevicePostureUpdateParamsType = "carbonblack"
	DevicePostureUpdateParamsTypeFirewall            DevicePostureUpdateParamsType = "firewall"
	DevicePostureUpdateParamsTypeOSVersion           DevicePostureUpdateParamsType = "os_version"
	DevicePostureUpdateParamsTypeDomainJoined        DevicePostureUpdateParamsType = "domain_joined"
	DevicePostureUpdateParamsTypeClientCertificate   DevicePostureUpdateParamsType = "client_certificate"
	DevicePostureUpdateParamsTypeClientCertificateV2 DevicePostureUpdateParamsType = "client_certificate_v2"
	DevicePostureUpdateParamsTypeUniqueClientID      DevicePostureUpdateParamsType = "unique_client_id"
	DevicePostureUpdateParamsTypeKolide              DevicePostureUpdateParamsType = "kolide"
	DevicePostureUpdateParamsTypeTaniumS2s           DevicePostureUpdateParamsType = "tanium_s2s"
	DevicePostureUpdateParamsTypeCrowdstrikeS2s      DevicePostureUpdateParamsType = "crowdstrike_s2s"
	DevicePostureUpdateParamsTypeIntune              DevicePostureUpdateParamsType = "intune"
	DevicePostureUpdateParamsTypeWorkspaceOne        DevicePostureUpdateParamsType = "workspace_one"
	DevicePostureUpdateParamsTypeSentineloneS2s      DevicePostureUpdateParamsType = "sentinelone_s2s"
	DevicePostureUpdateParamsTypeCustomS2s           DevicePostureUpdateParamsType = "custom_s2s"
)

func (r DevicePostureUpdateParamsType) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsTypeFile, DevicePostureUpdateParamsTypeApplication, DevicePostureUpdateParamsTypeTanium, DevicePostureUpdateParamsTypeGateway, DevicePostureUpdateParamsTypeWARP, DevicePostureUpdateParamsTypeDiskEncryption, DevicePostureUpdateParamsTypeSerialNumber, DevicePostureUpdateParamsTypeSentinelone, DevicePostureUpdateParamsTypeCarbonblack, DevicePostureUpdateParamsTypeFirewall, DevicePostureUpdateParamsTypeOSVersion, DevicePostureUpdateParamsTypeDomainJoined, DevicePostureUpdateParamsTypeClientCertificate, DevicePostureUpdateParamsTypeClientCertificateV2, DevicePostureUpdateParamsTypeUniqueClientID, DevicePostureUpdateParamsTypeKolide, DevicePostureUpdateParamsTypeTaniumS2s, DevicePostureUpdateParamsTypeCrowdstrikeS2s, DevicePostureUpdateParamsTypeIntune, DevicePostureUpdateParamsTypeWorkspaceOne, DevicePostureUpdateParamsTypeSentineloneS2s, DevicePostureUpdateParamsTypeCustomS2s:
		return true
	}
	return false
}

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePostureUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelope]
type devicePostureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

func (r DevicePostureUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DevicePostureDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelope]
type devicePostureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureGetResponseEnvelope]
type devicePostureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)

func (r DevicePostureGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
