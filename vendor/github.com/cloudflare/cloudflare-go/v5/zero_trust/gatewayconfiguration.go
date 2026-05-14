// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// GatewayConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayConfigurationService] method instead.
type GatewayConfigurationService struct {
	Options           []option.RequestOption
	CustomCertificate *GatewayConfigurationCustomCertificateService
}

// NewGatewayConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayConfigurationService(opts ...option.RequestOption) (r *GatewayConfigurationService) {
	r = &GatewayConfigurationService{}
	r.Options = opts
	r.CustomCertificate = NewGatewayConfigurationCustomCertificateService(opts...)
	return
}

// Updates the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Update(ctx context.Context, params GatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayConfigurationUpdateResponse, err error) {
	var env GatewayConfigurationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `certificate`, without updating the entire configuration object. Returns an
// error if any collection of settings is not properly configured.
func (r *GatewayConfigurationService) Edit(ctx context.Context, params GatewayConfigurationEditParams, opts ...option.RequestOption) (res *GatewayConfigurationEditResponse, err error) {
	var env GatewayConfigurationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Get(ctx context.Context, query GatewayConfigurationGetParams, opts ...option.RequestOption) (res *GatewayConfigurationGetResponse, err error) {
	var env GatewayConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Activity log settings.
type ActivityLogSettings struct {
	// Enable activity logging.
	Enabled bool                    `json:"enabled,nullable"`
	JSON    activityLogSettingsJSON `json:"-"`
}

// activityLogSettingsJSON contains the JSON metadata for the struct
// [ActivityLogSettings]
type activityLogSettingsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivityLogSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activityLogSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type ActivityLogSettingsParam struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ActivityLogSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type AntiVirusSettings struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase,nullable"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase,nullable"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed,nullable"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings NotificationSettings  `json:"notification_settings,nullable"`
	JSON                 antiVirusSettingsJSON `json:"-"`
}

// antiVirusSettingsJSON contains the JSON metadata for the struct
// [AntiVirusSettings]
type antiVirusSettingsJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AntiVirusSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r antiVirusSettingsJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type AntiVirusSettingsParam struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[NotificationSettingsParam] `json:"notification_settings"`
}

func (r AntiVirusSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type BlockPageSettings struct {
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled,nullable"`
	// If mode is customized_block_page: block page footer text.
	FooterText string `json:"footer_text"`
	// If mode is customized_block_page: block page header text.
	HeaderText string `json:"header_text"`
	// If mode is redirect_uri: when enabled, context will be appended to target_uri as
	// query parameters.
	IncludeContext bool `json:"include_context"`
	// If mode is customized_block_page: full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// If mode is customized_block_page: admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// If mode is customized_block_page: subject line for emails created from block
	// page.
	MailtoSubject string `json:"mailto_subject"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode BlockPageSettingsMode `json:"mode"`
	// If mode is customized_block_page: block page title.
	Name string `json:"name"`
	// This setting was shared via the Orgs API and cannot be edited by the current
	// account
	ReadOnly bool `json:"read_only,nullable"`
	// Account tag of account that shared this setting
	SourceAccount string `json:"source_account,nullable"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the
	// block page.
	SuppressFooter bool `json:"suppress_footer"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	TargetURI string `json:"target_uri" format:"uri"`
	// Version number of the setting
	Version int64                 `json:"version,nullable"`
	JSON    blockPageSettingsJSON `json:"-"`
}

// blockPageSettingsJSON contains the JSON metadata for the struct
// [BlockPageSettings]
type blockPageSettingsJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	IncludeContext  apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Mode            apijson.Field
	Name            apijson.Field
	ReadOnly        apijson.Field
	SourceAccount   apijson.Field
	SuppressFooter  apijson.Field
	TargetURI       apijson.Field
	Version         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BlockPageSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockPageSettingsJSON) RawJSON() string {
	return r.raw
}

// Controls whether the user is redirected to a Cloudflare-hosted block page or to
// a customer-provided URI.
type BlockPageSettingsMode string

const (
	BlockPageSettingsModeEmpty               BlockPageSettingsMode = ""
	BlockPageSettingsModeCustomizedBlockPage BlockPageSettingsMode = "customized_block_page"
	BlockPageSettingsModeRedirectURI         BlockPageSettingsMode = "redirect_uri"
)

func (r BlockPageSettingsMode) IsKnown() bool {
	switch r {
	case BlockPageSettingsModeEmpty, BlockPageSettingsModeCustomizedBlockPage, BlockPageSettingsModeRedirectURI:
		return true
	}
	return false
}

// Block page layout settings.
type BlockPageSettingsParam struct {
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled"`
	// If mode is customized_block_page: block page footer text.
	FooterText param.Field[string] `json:"footer_text"`
	// If mode is customized_block_page: block page header text.
	HeaderText param.Field[string] `json:"header_text"`
	// If mode is redirect_uri: when enabled, context will be appended to target_uri as
	// query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
	// If mode is customized_block_page: full URL to the logo file.
	LogoPath param.Field[string] `json:"logo_path"`
	// If mode is customized_block_page: admin email for users to contact.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// If mode is customized_block_page: subject line for emails created from block
	// page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode param.Field[BlockPageSettingsMode] `json:"mode"`
	// If mode is customized_block_page: block page title.
	Name param.Field[string] `json:"name"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the
	// block page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	TargetURI param.Field[string] `json:"target_uri" format:"uri"`
}

func (r BlockPageSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type BodyScanningSettings struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode BodyScanningSettingsInspectionMode `json:"inspection_mode"`
	JSON           bodyScanningSettingsJSON           `json:"-"`
}

// bodyScanningSettingsJSON contains the JSON metadata for the struct
// [BodyScanningSettings]
type bodyScanningSettingsJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BodyScanningSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bodyScanningSettingsJSON) RawJSON() string {
	return r.raw
}

// Set the inspection mode to either `deep` or `shallow`.
type BodyScanningSettingsInspectionMode string

const (
	BodyScanningSettingsInspectionModeDeep    BodyScanningSettingsInspectionMode = "deep"
	BodyScanningSettingsInspectionModeShallow BodyScanningSettingsInspectionMode = "shallow"
)

func (r BodyScanningSettingsInspectionMode) IsKnown() bool {
	switch r {
	case BodyScanningSettingsInspectionModeDeep, BodyScanningSettingsInspectionModeShallow:
		return true
	}
	return false
}

// DLP body scanning settings.
type BodyScanningSettingsParam struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[BodyScanningSettingsInspectionMode] `json:"inspection_mode"`
}

func (r BodyScanningSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type BrowserIsolationSettings struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                         `json:"url_browser_isolation_enabled"`
	JSON                       browserIsolationSettingsJSON `json:"-"`
}

// browserIsolationSettingsJSON contains the JSON metadata for the struct
// [BrowserIsolationSettings]
type browserIsolationSettingsJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *BrowserIsolationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserIsolationSettingsJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type BrowserIsolationSettingsParam struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r BrowserIsolationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI. (deprecated and replaced by
// `certificate`)
//
// Deprecated: deprecated
type CustomCertificateSettings struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required,nullable"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                        `json:"binding_status"`
	UpdatedAt     time.Time                     `json:"updated_at" format:"date-time"`
	JSON          customCertificateSettingsJSON `json:"-"`
}

// customCertificateSettingsJSON contains the JSON metadata for the struct
// [CustomCertificateSettings]
type customCertificateSettingsJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomCertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateSettingsJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI. (deprecated and replaced by
// `certificate`)
//
// Deprecated: deprecated
type CustomCertificateSettingsParam struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r CustomCertificateSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type ExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool `json:"enabled,nullable"`
	// This setting was shared via the Orgs API and cannot be edited by the current
	// account
	ReadOnly bool `json:"read_only"`
	// Account tag of account that shared this setting
	SourceAccount string `json:"source_account"`
	// Version number of the setting
	Version int64                     `json:"version"`
	JSON    extendedEmailMatchingJSON `json:"-"`
}

// extendedEmailMatchingJSON contains the JSON metadata for the struct
// [ExtendedEmailMatching]
type extendedEmailMatchingJSON struct {
	Enabled       apijson.Field
	ReadOnly      apijson.Field
	SourceAccount apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type ExtendedEmailMatchingParam struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ExtendedEmailMatchingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type FipsSettings struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool             `json:"tls"`
	JSON fipsSettingsJSON `json:"-"`
}

// fipsSettingsJSON contains the JSON metadata for the struct [FipsSettings]
type fipsSettingsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FipsSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fipsSettingsJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type FipsSettingsParam struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r FipsSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type GatewayConfigurationSettings struct {
	// Activity log settings.
	ActivityLog ActivityLogSettings `json:"activity_log,nullable"`
	// Anti-virus settings.
	Antivirus AntiVirusSettings `json:"antivirus,nullable"`
	// Block page layout settings.
	BlockPage BlockPageSettings `json:"block_page,nullable"`
	// DLP body scanning settings.
	BodyScanning BodyScanningSettings `json:"body_scanning,nullable"`
	// Browser isolation settings.
	BrowserIsolation BrowserIsolationSettings `json:"browser_isolation,nullable"`
	// Certificate settings for Gateway TLS interception. If not specified, the
	// Cloudflare Root CA will be used.
	Certificate GatewayConfigurationSettingsCertificate `json:"certificate,nullable"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by
	// `certificate`)
	//
	// Deprecated: deprecated
	CustomCertificate CustomCertificateSettings `json:"custom_certificate,nullable"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching ExtendedEmailMatching `json:"extended_email_matching,nullable"`
	// FIPS settings.
	Fips FipsSettings `json:"fips,nullable"`
	// Setting to enable host selector in egress policies.
	HostSelector GatewayConfigurationSettingsHostSelector `json:"host_selector,nullable"`
	// Setting to define inspection settings
	Inspection GatewayConfigurationSettingsInspection `json:"inspection,nullable"`
	// Protocol Detection settings.
	ProtocolDetection ProtocolDetection `json:"protocol_detection,nullable"`
	// Sandbox settings.
	Sandbox GatewayConfigurationSettingsSandbox `json:"sandbox,nullable"`
	// TLS interception settings.
	TLSDecrypt TLSSettings                      `json:"tls_decrypt,nullable"`
	JSON       gatewayConfigurationSettingsJSON `json:"-"`
}

// gatewayConfigurationSettingsJSON contains the JSON metadata for the struct
// [GatewayConfigurationSettings]
type gatewayConfigurationSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	Certificate           apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	HostSelector          apijson.Field
	Inspection            apijson.Field
	ProtocolDetection     apijson.Field
	Sandbox               apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayConfigurationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsJSON) RawJSON() string {
	return r.raw
}

// Certificate settings for Gateway TLS interception. If not specified, the
// Cloudflare Root CA will be used.
type GatewayConfigurationSettingsCertificate struct {
	// UUID of certificate to be used for interception. Certificate must be available
	// (previously called 'active') on the edge. A nil UUID will indicate the
	// Cloudflare Root CA should be used.
	ID   string                                      `json:"id,required"`
	JSON gatewayConfigurationSettingsCertificateJSON `json:"-"`
}

// gatewayConfigurationSettingsCertificateJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsCertificate]
type gatewayConfigurationSettingsCertificateJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsCertificateJSON) RawJSON() string {
	return r.raw
}

// Setting to enable host selector in egress policies.
type GatewayConfigurationSettingsHostSelector struct {
	// Enable filtering via hosts for egress policies.
	Enabled bool                                         `json:"enabled,nullable"`
	JSON    gatewayConfigurationSettingsHostSelectorJSON `json:"-"`
}

// gatewayConfigurationSettingsHostSelectorJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsHostSelector]
type gatewayConfigurationSettingsHostSelectorJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsHostSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsHostSelectorJSON) RawJSON() string {
	return r.raw
}

// Setting to define inspection settings
type GatewayConfigurationSettingsInspection struct {
	// Defines the mode of inspection the proxy will use.
	//
	//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
	//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
	//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
	//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
	Mode GatewayConfigurationSettingsInspectionMode `json:"mode"`
	JSON gatewayConfigurationSettingsInspectionJSON `json:"-"`
}

// gatewayConfigurationSettingsInspectionJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsInspection]
type gatewayConfigurationSettingsInspectionJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsInspection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsInspectionJSON) RawJSON() string {
	return r.raw
}

// Defines the mode of inspection the proxy will use.
//
//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
type GatewayConfigurationSettingsInspectionMode string

const (
	GatewayConfigurationSettingsInspectionModeStatic  GatewayConfigurationSettingsInspectionMode = "static"
	GatewayConfigurationSettingsInspectionModeDynamic GatewayConfigurationSettingsInspectionMode = "dynamic"
)

func (r GatewayConfigurationSettingsInspectionMode) IsKnown() bool {
	switch r {
	case GatewayConfigurationSettingsInspectionModeStatic, GatewayConfigurationSettingsInspectionModeDynamic:
		return true
	}
	return false
}

// Sandbox settings.
type GatewayConfigurationSettingsSandbox struct {
	// Enable sandbox.
	Enabled bool `json:"enabled,nullable"`
	// Action to take when the file cannot be scanned.
	FallbackAction GatewayConfigurationSettingsSandboxFallbackAction `json:"fallback_action"`
	JSON           gatewayConfigurationSettingsSandboxJSON           `json:"-"`
}

// gatewayConfigurationSettingsSandboxJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsSandbox]
type gatewayConfigurationSettingsSandboxJSON struct {
	Enabled        apijson.Field
	FallbackAction apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsSandbox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsSandboxJSON) RawJSON() string {
	return r.raw
}

// Action to take when the file cannot be scanned.
type GatewayConfigurationSettingsSandboxFallbackAction string

const (
	GatewayConfigurationSettingsSandboxFallbackActionAllow GatewayConfigurationSettingsSandboxFallbackAction = "allow"
	GatewayConfigurationSettingsSandboxFallbackActionBlock GatewayConfigurationSettingsSandboxFallbackAction = "block"
)

func (r GatewayConfigurationSettingsSandboxFallbackAction) IsKnown() bool {
	switch r {
	case GatewayConfigurationSettingsSandboxFallbackActionAllow, GatewayConfigurationSettingsSandboxFallbackActionBlock:
		return true
	}
	return false
}

// Account settings
type GatewayConfigurationSettingsParam struct {
	// Activity log settings.
	ActivityLog param.Field[ActivityLogSettingsParam] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[AntiVirusSettingsParam] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[BlockPageSettingsParam] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[BodyScanningSettingsParam] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[BrowserIsolationSettingsParam] `json:"browser_isolation"`
	// Certificate settings for Gateway TLS interception. If not specified, the
	// Cloudflare Root CA will be used.
	Certificate param.Field[GatewayConfigurationSettingsCertificateParam] `json:"certificate"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by
	// `certificate`)
	//
	// Deprecated: deprecated
	CustomCertificate param.Field[CustomCertificateSettingsParam] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[ExtendedEmailMatchingParam] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[FipsSettingsParam] `json:"fips"`
	// Setting to enable host selector in egress policies.
	HostSelector param.Field[GatewayConfigurationSettingsHostSelectorParam] `json:"host_selector"`
	// Setting to define inspection settings
	Inspection param.Field[GatewayConfigurationSettingsInspectionParam] `json:"inspection"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[ProtocolDetectionParam] `json:"protocol_detection"`
	// Sandbox settings.
	Sandbox param.Field[GatewayConfigurationSettingsSandboxParam] `json:"sandbox"`
	// TLS interception settings.
	TLSDecrypt param.Field[TLSSettingsParam] `json:"tls_decrypt"`
}

func (r GatewayConfigurationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate settings for Gateway TLS interception. If not specified, the
// Cloudflare Root CA will be used.
type GatewayConfigurationSettingsCertificateParam struct {
	// UUID of certificate to be used for interception. Certificate must be available
	// (previously called 'active') on the edge. A nil UUID will indicate the
	// Cloudflare Root CA should be used.
	ID param.Field[string] `json:"id,required"`
}

func (r GatewayConfigurationSettingsCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Setting to enable host selector in egress policies.
type GatewayConfigurationSettingsHostSelectorParam struct {
	// Enable filtering via hosts for egress policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationSettingsHostSelectorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Setting to define inspection settings
type GatewayConfigurationSettingsInspectionParam struct {
	// Defines the mode of inspection the proxy will use.
	//
	//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
	//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
	//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
	//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
	Mode param.Field[GatewayConfigurationSettingsInspectionMode] `json:"mode"`
}

func (r GatewayConfigurationSettingsInspectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sandbox settings.
type GatewayConfigurationSettingsSandboxParam struct {
	// Enable sandbox.
	Enabled param.Field[bool] `json:"enabled"`
	// Action to take when the file cannot be scanned.
	FallbackAction param.Field[GatewayConfigurationSettingsSandboxFallbackAction] `json:"fallback_action"`
}

func (r GatewayConfigurationSettingsSandboxParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type NotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                   `json:"support_url"`
	JSON       notificationSettingsJSON `json:"-"`
}

// notificationSettingsJSON contains the JSON metadata for the struct
// [NotificationSettings]
type notificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type NotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r NotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type ProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                  `json:"enabled,nullable"`
	JSON    protocolDetectionJSON `json:"-"`
}

// protocolDetectionJSON contains the JSON metadata for the struct
// [ProtocolDetection]
type protocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r protocolDetectionJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type ProtocolDetectionParam struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ProtocolDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type TLSSettings struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool            `json:"enabled"`
	JSON    tlsSettingsJSON `json:"-"`
}

// tlsSettingsJSON contains the JSON metadata for the struct [TLSSettings]
type tlsSettingsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsSettingsJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type TLSSettingsParam struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TLSSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type GatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account settings
	Settings  GatewayConfigurationSettings           `json:"settings"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationUpdateResponseJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationUpdateResponse]
type gatewayConfigurationUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayConfigurationEditResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account settings
	Settings  GatewayConfigurationSettings         `json:"settings"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationEditResponseJSON `json:"-"`
}

// gatewayConfigurationEditResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationEditResponse]
type gatewayConfigurationEditResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayConfigurationGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account settings
	Settings  GatewayConfigurationSettings        `json:"settings"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationGetResponseJSON `json:"-"`
}

// gatewayConfigurationGetResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationGetResponse]
type gatewayConfigurationGetResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayConfigurationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Account settings
	Settings param.Field[GatewayConfigurationSettingsParam] `json:"settings"`
}

func (r GatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationUpdateResponse             `json:"result"`
	JSON   gatewayConfigurationUpdateResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayConfigurationUpdateResponseEnvelope]
type gatewayConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationUpdateResponseEnvelopeSuccess bool

const (
	GatewayConfigurationUpdateResponseEnvelopeSuccessTrue GatewayConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Account settings
	Settings param.Field[GatewayConfigurationSettingsParam] `json:"settings"`
}

func (r GatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationEditResponse             `json:"result"`
	JSON   gatewayConfigurationEditResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationEditResponseEnvelope]
type gatewayConfigurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationEditResponseEnvelopeSuccess bool

const (
	GatewayConfigurationEditResponseEnvelopeSuccessTrue GatewayConfigurationEditResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	// Account settings
	Result GatewayConfigurationGetResponse             `json:"result"`
	JSON   gatewayConfigurationGetResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationGetResponseEnvelope]
type gatewayConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationGetResponseEnvelopeSuccess bool

const (
	GatewayConfigurationGetResponseEnvelopeSuccessTrue GatewayConfigurationGetResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
