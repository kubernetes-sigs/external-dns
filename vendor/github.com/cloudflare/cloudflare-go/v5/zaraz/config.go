// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Updates Zaraz configuration for a zone.
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *Configuration, err error) {
	var env ConfigUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets latest Zaraz configuration for a zone. It can be preview or published
// configuration, whichever was the last updated. Secret variables values will not
// be included.
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *Configuration, err error) {
	var env ConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz configuration
type Configuration struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ConfigurationSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ConfigurationTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ConfigurationTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ConfigurationVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Cloudflare Monitoring settings.
	Analytics ConfigurationAnalytics `json:"analytics"`
	// Consent management configuration.
	Consent ConfigurationConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool              `json:"historyChange"`
	JSON          configurationJSON `json:"-"`
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	Settings      apijson.Field
	Tools         apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	Analytics     apijson.Field
	Consent       apijson.Field
	HistoryChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

// General Zaraz settings.
type ConfigurationSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ConfigurationSettingsContextEnricher `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain string `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce bool `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath string `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer bool `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress bool `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams bool `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent bool `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath string `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes bool `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath string `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath string `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath string                    `json:"trackPath"`
	JSON      configurationSettingsJSON `json:"-"`
}

// configurationSettingsJSON contains the JSON metadata for the struct
// [ConfigurationSettings]
type configurationSettingsJSON struct {
	AutoInjectScript    apijson.Field
	ContextEnricher     apijson.Field
	CookieDomain        apijson.Field
	Ecommerce           apijson.Field
	EventsAPIPath       apijson.Field
	HideExternalReferer apijson.Field
	HideIPAddress       apijson.Field
	HideQueryParams     apijson.Field
	HideUserAgent       apijson.Field
	InitPath            apijson.Field
	InjectIframes       apijson.Field
	McRootPath          apijson.Field
	ScriptPath          apijson.Field
	TrackPath           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConfigurationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationSettingsJSON) RawJSON() string {
	return r.raw
}

// Details of the worker that receives and edits Zaraz Context object.
type ConfigurationSettingsContextEnricher struct {
	EscapedWorkerName string                                   `json:"escapedWorkerName,required"`
	WorkerTag         string                                   `json:"workerTag,required"`
	JSON              configurationSettingsContextEnricherJSON `json:"-"`
}

// configurationSettingsContextEnricherJSON contains the JSON metadata for the
// struct [ConfigurationSettingsContextEnricher]
type configurationSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigurationSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationSettingsContextEnricherJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTool struct {
	// This field can have the runtime type of [[]string].
	BlockingTriggers interface{} `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// This field can have the runtime type of
	// [map[string]ConfigurationToolsZarazManagedComponentDefaultFieldsUnion],
	// [map[string]ConfigurationToolsWorkerDefaultFieldsUnion].
	DefaultFields interface{} `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// This field can have the runtime type of [[]string].
	Permissions interface{} `json:"permissions,required"`
	// This field can have the runtime type of
	// [map[string]ConfigurationToolsZarazManagedComponentSettingsUnion],
	// [map[string]ConfigurationToolsWorkerSettingsUnion].
	Settings interface{}            `json:"settings,required"`
	Type     ConfigurationToolsType `json:"type,required"`
	// This field can have the runtime type of [map[string]NeoEvent].
	Actions interface{} `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// This field can have the runtime type of [[]NeoEvent].
	NeoEvents interface{} `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string `json:"vendorPolicyUrl"`
	// This field can have the runtime type of [ConfigurationToolsWorkerWorker].
	Worker interface{}           `json:"worker"`
	JSON   configurationToolJSON `json:"-"`
	union  ConfigurationToolsUnion
}

// configurationToolJSON contains the JSON metadata for the struct
// [ConfigurationTool]
type configurationToolJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	Worker           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r configurationToolJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationTool) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationTool{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationToolsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationToolsZarazManagedComponent], [ConfigurationToolsWorker].
func (r ConfigurationTool) AsUnion() ConfigurationToolsUnion {
	return r.union
}

// Union satisfied by [ConfigurationToolsZarazManagedComponent] or
// [ConfigurationToolsWorker].
type ConfigurationToolsUnion interface {
	implementsConfigurationTool()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationToolsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationToolsZarazManagedComponent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationToolsWorker{}),
		},
	)
}

type ConfigurationToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ConfigurationToolsZarazManagedComponentDefaultFieldsUnion `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ConfigurationToolsZarazManagedComponentSettingsUnion `json:"settings,required"`
	Type     ConfigurationToolsZarazManagedComponentType                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]NeoEvent `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []NeoEvent `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string                                      `json:"vendorPolicyUrl"`
	JSON            configurationToolsZarazManagedComponentJSON `json:"-"`
}

// configurationToolsZarazManagedComponentJSON contains the JSON metadata for the
// struct [ConfigurationToolsZarazManagedComponent]
type configurationToolsZarazManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigurationToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationToolsZarazManagedComponentJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationToolsZarazManagedComponent) implementsConfigurationTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ConfigurationToolsZarazManagedComponentDefaultFieldsUnion interface {
	ImplementsConfigurationToolsZarazManagedComponentDefaultFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationToolsZarazManagedComponentDefaultFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ConfigurationToolsZarazManagedComponentSettingsUnion interface {
	ImplementsConfigurationToolsZarazManagedComponentSettingsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationToolsZarazManagedComponentSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ConfigurationToolsZarazManagedComponentType string

const (
	ConfigurationToolsZarazManagedComponentTypeComponent ConfigurationToolsZarazManagedComponentType = "component"
)

func (r ConfigurationToolsZarazManagedComponentType) IsKnown() bool {
	switch r {
	case ConfigurationToolsZarazManagedComponentTypeComponent:
		return true
	}
	return false
}

type ConfigurationToolsWorker struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ConfigurationToolsWorkerDefaultFieldsUnion `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ConfigurationToolsWorkerSettingsUnion `json:"settings,required"`
	Type     ConfigurationToolsWorkerType                     `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ConfigurationToolsWorkerWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]NeoEvent `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []NeoEvent `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName string `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL string                       `json:"vendorPolicyUrl"`
	JSON            configurationToolsWorkerJSON `json:"-"`
}

// configurationToolsWorkerJSON contains the JSON metadata for the struct
// [ConfigurationToolsWorker]
type configurationToolsWorkerJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Worker           apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	VendorName       apijson.Field
	VendorPolicyURL  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigurationToolsWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationToolsWorkerJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationToolsWorker) implementsConfigurationTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ConfigurationToolsWorkerDefaultFieldsUnion interface {
	ImplementsConfigurationToolsWorkerDefaultFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationToolsWorkerDefaultFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ConfigurationToolsWorkerSettingsUnion interface {
	ImplementsConfigurationToolsWorkerSettingsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationToolsWorkerSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ConfigurationToolsWorkerType string

const (
	ConfigurationToolsWorkerTypeCustomMc ConfigurationToolsWorkerType = "custom-mc"
)

func (r ConfigurationToolsWorkerType) IsKnown() bool {
	switch r {
	case ConfigurationToolsWorkerTypeCustomMc:
		return true
	}
	return false
}

// Cloudflare worker that acts as a managed component
type ConfigurationToolsWorkerWorker struct {
	EscapedWorkerName string                             `json:"escapedWorkerName,required"`
	WorkerTag         string                             `json:"workerTag,required"`
	JSON              configurationToolsWorkerWorkerJSON `json:"-"`
}

// configurationToolsWorkerWorkerJSON contains the JSON metadata for the struct
// [ConfigurationToolsWorkerWorker]
type configurationToolsWorkerWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigurationToolsWorkerWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationToolsWorkerWorkerJSON) RawJSON() string {
	return r.raw
}

type ConfigurationToolsType string

const (
	ConfigurationToolsTypeComponent ConfigurationToolsType = "component"
	ConfigurationToolsTypeCustomMc  ConfigurationToolsType = "custom-mc"
)

func (r ConfigurationToolsType) IsKnown() bool {
	switch r {
	case ConfigurationToolsTypeComponent, ConfigurationToolsTypeCustomMc:
		return true
	}
	return false
}

type ConfigurationTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ConfigurationTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ConfigurationTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                      `json:"description"`
	System      ConfigurationTriggersSystem `json:"system"`
	JSON        configurationTriggerJSON    `json:"-"`
}

// configurationTriggerJSON contains the JSON metadata for the struct
// [ConfigurationTrigger]
type configurationTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConfigurationTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggerJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRule struct {
	ID     string                                  `json:"id,required"`
	Action ConfigurationTriggersExcludeRulesAction `json:"action"`
	Match  string                                  `json:"match"`
	Op     ConfigurationTriggersExcludeRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettings],
	// [ConfigurationTriggersExcludeRulesZarazTimerRuleSettings],
	// [ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleSettings],
	// [ConfigurationTriggersExcludeRulesZarazVariableMatchRuleSettings],
	// [ConfigurationTriggersExcludeRulesZarazScrollDepthRuleSettings],
	// [ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                          `json:"settings"`
	Value    string                               `json:"value"`
	JSON     configurationTriggersExcludeRuleJSON `json:"-"`
	union    ConfigurationTriggersExcludeRulesUnion
}

// configurationTriggersExcludeRuleJSON contains the JSON metadata for the struct
// [ConfigurationTriggersExcludeRule]
type configurationTriggersExcludeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationTriggersExcludeRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationTriggersExcludeRule) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationTriggersExcludeRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationTriggersExcludeRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationTriggersExcludeRulesZarazLoadRule],
// [ConfigurationTriggersExcludeRulesZarazClickListenerRule],
// [ConfigurationTriggersExcludeRulesZarazTimerRule],
// [ConfigurationTriggersExcludeRulesZarazFormSubmissionRule],
// [ConfigurationTriggersExcludeRulesZarazVariableMatchRule],
// [ConfigurationTriggersExcludeRulesZarazScrollDepthRule],
// [ConfigurationTriggersExcludeRulesZarazElementVisibilityRule].
func (r ConfigurationTriggersExcludeRule) AsUnion() ConfigurationTriggersExcludeRulesUnion {
	return r.union
}

// Union satisfied by [ConfigurationTriggersExcludeRulesZarazLoadRule],
// [ConfigurationTriggersExcludeRulesZarazClickListenerRule],
// [ConfigurationTriggersExcludeRulesZarazTimerRule],
// [ConfigurationTriggersExcludeRulesZarazFormSubmissionRule],
// [ConfigurationTriggersExcludeRulesZarazVariableMatchRule],
// [ConfigurationTriggersExcludeRulesZarazScrollDepthRule] or
// [ConfigurationTriggersExcludeRulesZarazElementVisibilityRule].
type ConfigurationTriggersExcludeRulesUnion interface {
	implementsConfigurationTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationTriggersExcludeRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersExcludeRulesZarazElementVisibilityRule{}),
		},
	)
}

type ConfigurationTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                             `json:"id,required"`
	Match string                                             `json:"match,required"`
	Op    ConfigurationTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                             `json:"value,required"`
	JSON  configurationTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazLoadRuleJSON contains the JSON metadata
// for the struct [ConfigurationTriggersExcludeRulesZarazLoadRule]
type configurationTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazLoadRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazLoadRuleOp string

const (
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpContains           ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpEquals             ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpStartsWith         ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpEndsWith           ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpLessThan           ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ConfigurationTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ConfigurationTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigurationTriggersExcludeRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazLoadRuleOpContains, ConfigurationTriggersExcludeRulesZarazLoadRuleOpEquals, ConfigurationTriggersExcludeRulesZarazLoadRuleOpStartsWith, ConfigurationTriggersExcludeRulesZarazLoadRuleOpEndsWith, ConfigurationTriggersExcludeRulesZarazLoadRuleOpMatchRegex, ConfigurationTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex, ConfigurationTriggersExcludeRulesZarazLoadRuleOpGreaterThan, ConfigurationTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual, ConfigurationTriggersExcludeRulesZarazLoadRuleOpLessThan, ConfigurationTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                          `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazClickListenerRuleJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazClickListenerRule]
type configurationTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazClickListenerRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazClickListenerRuleActionClickListener ConfigurationTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ConfigurationTriggersExcludeRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                              `json:"selector,required"`
	Type        ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                               `json:"waitForTags,required"`
	JSON        configurationTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazClickListenerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettings]
type configurationTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCSS   ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath, ConfigurationTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCSS:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                  `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazTimerRuleJSON contains the JSON metadata
// for the struct [ConfigurationTriggersExcludeRulesZarazTimerRule]
type configurationTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazTimerRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazTimerRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazTimerRuleActionTimer ConfigurationTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

func (r ConfigurationTriggersExcludeRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                       `json:"interval,required"`
	Limit    int64                                                       `json:"limit,required"`
	JSON     configurationTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazTimerRuleSettingsJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazTimerRuleSettings]
type configurationTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazFormSubmissionRuleJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazFormSubmissionRule]
type configurationTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazFormSubmissionRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                               `json:"selector,required"`
	Validate bool                                                                 `json:"validate,required"`
	JSON     configurationTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON contains
// the JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type configurationTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                          `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazVariableMatchRuleJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazVariableMatchRule]
type configurationTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazVariableMatchRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ConfigurationTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ConfigurationTriggersExcludeRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                              `json:"match,required"`
	Variable string                                                              `json:"variable,required"`
	JSON     configurationTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazVariableMatchRuleSettings]
type configurationTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                        `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct [ConfigurationTriggersExcludeRulesZarazScrollDepthRule]
type configurationTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazScrollDepthRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ConfigurationTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ConfigurationTriggersExcludeRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                            `json:"positions,required"`
	JSON      configurationTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazScrollDepthRuleSettings]
type configurationTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     configurationTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// configurationTriggersExcludeRulesZarazElementVisibilityRuleJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazElementVisibilityRule]
type configurationTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersExcludeRulesZarazElementVisibilityRule) implementsConfigurationTriggersExcludeRule() {
}

type ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                  `json:"selector,required"`
	JSON     configurationTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// configurationTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON contains
// the JSON metadata for the struct
// [ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type configurationTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersExcludeRulesAction string

const (
	ConfigurationTriggersExcludeRulesActionClickListener     ConfigurationTriggersExcludeRulesAction = "clickListener"
	ConfigurationTriggersExcludeRulesActionTimer             ConfigurationTriggersExcludeRulesAction = "timer"
	ConfigurationTriggersExcludeRulesActionFormSubmission    ConfigurationTriggersExcludeRulesAction = "formSubmission"
	ConfigurationTriggersExcludeRulesActionVariableMatch     ConfigurationTriggersExcludeRulesAction = "variableMatch"
	ConfigurationTriggersExcludeRulesActionScrollDepth       ConfigurationTriggersExcludeRulesAction = "scrollDepth"
	ConfigurationTriggersExcludeRulesActionElementVisibility ConfigurationTriggersExcludeRulesAction = "elementVisibility"
)

func (r ConfigurationTriggersExcludeRulesAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesActionClickListener, ConfigurationTriggersExcludeRulesActionTimer, ConfigurationTriggersExcludeRulesActionFormSubmission, ConfigurationTriggersExcludeRulesActionVariableMatch, ConfigurationTriggersExcludeRulesActionScrollDepth, ConfigurationTriggersExcludeRulesActionElementVisibility:
		return true
	}
	return false
}

type ConfigurationTriggersExcludeRulesOp string

const (
	ConfigurationTriggersExcludeRulesOpContains           ConfigurationTriggersExcludeRulesOp = "CONTAINS"
	ConfigurationTriggersExcludeRulesOpEquals             ConfigurationTriggersExcludeRulesOp = "EQUALS"
	ConfigurationTriggersExcludeRulesOpStartsWith         ConfigurationTriggersExcludeRulesOp = "STARTS_WITH"
	ConfigurationTriggersExcludeRulesOpEndsWith           ConfigurationTriggersExcludeRulesOp = "ENDS_WITH"
	ConfigurationTriggersExcludeRulesOpMatchRegex         ConfigurationTriggersExcludeRulesOp = "MATCH_REGEX"
	ConfigurationTriggersExcludeRulesOpNotMatchRegex      ConfigurationTriggersExcludeRulesOp = "NOT_MATCH_REGEX"
	ConfigurationTriggersExcludeRulesOpGreaterThan        ConfigurationTriggersExcludeRulesOp = "GREATER_THAN"
	ConfigurationTriggersExcludeRulesOpGreaterThanOrEqual ConfigurationTriggersExcludeRulesOp = "GREATER_THAN_OR_EQUAL"
	ConfigurationTriggersExcludeRulesOpLessThan           ConfigurationTriggersExcludeRulesOp = "LESS_THAN"
	ConfigurationTriggersExcludeRulesOpLessThanOrEqual    ConfigurationTriggersExcludeRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigurationTriggersExcludeRulesOp) IsKnown() bool {
	switch r {
	case ConfigurationTriggersExcludeRulesOpContains, ConfigurationTriggersExcludeRulesOpEquals, ConfigurationTriggersExcludeRulesOpStartsWith, ConfigurationTriggersExcludeRulesOpEndsWith, ConfigurationTriggersExcludeRulesOpMatchRegex, ConfigurationTriggersExcludeRulesOpNotMatchRegex, ConfigurationTriggersExcludeRulesOpGreaterThan, ConfigurationTriggersExcludeRulesOpGreaterThanOrEqual, ConfigurationTriggersExcludeRulesOpLessThan, ConfigurationTriggersExcludeRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRule struct {
	ID     string                               `json:"id,required"`
	Action ConfigurationTriggersLoadRulesAction `json:"action"`
	Match  string                               `json:"match"`
	Op     ConfigurationTriggersLoadRulesOp     `json:"op"`
	// This field can have the runtime type of
	// [ConfigurationTriggersLoadRulesZarazClickListenerRuleSettings],
	// [ConfigurationTriggersLoadRulesZarazTimerRuleSettings],
	// [ConfigurationTriggersLoadRulesZarazFormSubmissionRuleSettings],
	// [ConfigurationTriggersLoadRulesZarazVariableMatchRuleSettings],
	// [ConfigurationTriggersLoadRulesZarazScrollDepthRuleSettings],
	// [ConfigurationTriggersLoadRulesZarazElementVisibilityRuleSettings].
	Settings interface{}                       `json:"settings"`
	Value    string                            `json:"value"`
	JSON     configurationTriggersLoadRuleJSON `json:"-"`
	union    ConfigurationTriggersLoadRulesUnion
}

// configurationTriggersLoadRuleJSON contains the JSON metadata for the struct
// [ConfigurationTriggersLoadRule]
type configurationTriggersLoadRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Settings    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationTriggersLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationTriggersLoadRule) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationTriggersLoadRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationTriggersLoadRulesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationTriggersLoadRulesZarazLoadRule],
// [ConfigurationTriggersLoadRulesZarazClickListenerRule],
// [ConfigurationTriggersLoadRulesZarazTimerRule],
// [ConfigurationTriggersLoadRulesZarazFormSubmissionRule],
// [ConfigurationTriggersLoadRulesZarazVariableMatchRule],
// [ConfigurationTriggersLoadRulesZarazScrollDepthRule],
// [ConfigurationTriggersLoadRulesZarazElementVisibilityRule].
func (r ConfigurationTriggersLoadRule) AsUnion() ConfigurationTriggersLoadRulesUnion {
	return r.union
}

// Union satisfied by [ConfigurationTriggersLoadRulesZarazLoadRule],
// [ConfigurationTriggersLoadRulesZarazClickListenerRule],
// [ConfigurationTriggersLoadRulesZarazTimerRule],
// [ConfigurationTriggersLoadRulesZarazFormSubmissionRule],
// [ConfigurationTriggersLoadRulesZarazVariableMatchRule],
// [ConfigurationTriggersLoadRulesZarazScrollDepthRule] or
// [ConfigurationTriggersLoadRulesZarazElementVisibilityRule].
type ConfigurationTriggersLoadRulesUnion interface {
	implementsConfigurationTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationTriggersLoadRulesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazLoadRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazClickListenerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazTimerRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazFormSubmissionRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazVariableMatchRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazScrollDepthRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationTriggersLoadRulesZarazElementVisibilityRule{}),
		},
	)
}

type ConfigurationTriggersLoadRulesZarazLoadRule struct {
	ID    string                                          `json:"id,required"`
	Match string                                          `json:"match,required"`
	Op    ConfigurationTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                          `json:"value,required"`
	JSON  configurationTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazLoadRuleJSON contains the JSON metadata for
// the struct [ConfigurationTriggersLoadRulesZarazLoadRule]
type configurationTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazLoadRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazLoadRule) implementsConfigurationTriggersLoadRule() {}

type ConfigurationTriggersLoadRulesZarazLoadRuleOp string

const (
	ConfigurationTriggersLoadRulesZarazLoadRuleOpContains           ConfigurationTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpEquals             ConfigurationTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpStartsWith         ConfigurationTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpEndsWith           ConfigurationTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpMatchRegex         ConfigurationTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ConfigurationTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpGreaterThan        ConfigurationTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ConfigurationTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpLessThan           ConfigurationTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ConfigurationTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ConfigurationTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigurationTriggersLoadRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazLoadRuleOpContains, ConfigurationTriggersLoadRulesZarazLoadRuleOpEquals, ConfigurationTriggersLoadRulesZarazLoadRuleOpStartsWith, ConfigurationTriggersLoadRulesZarazLoadRuleOpEndsWith, ConfigurationTriggersLoadRulesZarazLoadRuleOpMatchRegex, ConfigurationTriggersLoadRulesZarazLoadRuleOpNotMatchRegex, ConfigurationTriggersLoadRulesZarazLoadRuleOpGreaterThan, ConfigurationTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual, ConfigurationTriggersLoadRulesZarazLoadRuleOpLessThan, ConfigurationTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                       `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazClickListenerRuleJSON contains the JSON
// metadata for the struct [ConfigurationTriggersLoadRulesZarazClickListenerRule]
type configurationTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazClickListenerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazClickListenerRule) implementsConfigurationTriggersLoadRule() {
}

type ConfigurationTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazClickListenerRuleActionClickListener ConfigurationTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ConfigurationTriggersLoadRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                           `json:"selector,required"`
	Type        ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                            `json:"waitForTags,required"`
	JSON        configurationTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazClickListenerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersLoadRulesZarazClickListenerRuleSettings]
type configurationTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazClickListenerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsTypeCSS   ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath, ConfigurationTriggersLoadRulesZarazClickListenerRuleSettingsTypeCSS:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazTimerRule struct {
	ID       string                                               `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazTimerRuleJSON contains the JSON metadata for
// the struct [ConfigurationTriggersLoadRulesZarazTimerRule]
type configurationTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazTimerRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazTimerRule) implementsConfigurationTriggersLoadRule() {}

type ConfigurationTriggersLoadRulesZarazTimerRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazTimerRuleActionTimer ConfigurationTriggersLoadRulesZarazTimerRuleAction = "timer"
)

func (r ConfigurationTriggersLoadRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                    `json:"interval,required"`
	Limit    int64                                                    `json:"limit,required"`
	JSON     configurationTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazTimerRuleSettingsJSON contains the JSON
// metadata for the struct [ConfigurationTriggersLoadRulesZarazTimerRuleSettings]
type configurationTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazTimerRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                        `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazFormSubmissionRuleJSON contains the JSON
// metadata for the struct [ConfigurationTriggersLoadRulesZarazFormSubmissionRule]
type configurationTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazFormSubmissionRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazFormSubmissionRule) implementsConfigurationTriggersLoadRule() {
}

type ConfigurationTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ConfigurationTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ConfigurationTriggersLoadRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                            `json:"selector,required"`
	Validate bool                                                              `json:"validate,required"`
	JSON     configurationTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersLoadRulesZarazFormSubmissionRuleSettings]
type configurationTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                       `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazVariableMatchRuleJSON contains the JSON
// metadata for the struct [ConfigurationTriggersLoadRulesZarazVariableMatchRule]
type configurationTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazVariableMatchRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazVariableMatchRule) implementsConfigurationTriggersLoadRule() {
}

type ConfigurationTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ConfigurationTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ConfigurationTriggersLoadRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                           `json:"match,required"`
	Variable string                                                           `json:"variable,required"`
	JSON     configurationTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazVariableMatchRuleSettingsJSON contains the
// JSON metadata for the struct
// [ConfigurationTriggersLoadRulesZarazVariableMatchRuleSettings]
type configurationTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazVariableMatchRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                     `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazScrollDepthRuleJSON contains the JSON
// metadata for the struct [ConfigurationTriggersLoadRulesZarazScrollDepthRule]
type configurationTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazScrollDepthRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazScrollDepthRule) implementsConfigurationTriggersLoadRule() {
}

type ConfigurationTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ConfigurationTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ConfigurationTriggersLoadRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                         `json:"positions,required"`
	JSON      configurationTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazScrollDepthRuleSettingsJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersLoadRulesZarazScrollDepthRuleSettings]
type configurationTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazScrollDepthRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ConfigurationTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ConfigurationTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     configurationTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// configurationTriggersLoadRulesZarazElementVisibilityRuleJSON contains the JSON
// metadata for the struct
// [ConfigurationTriggersLoadRulesZarazElementVisibilityRule]
type configurationTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazElementVisibilityRuleJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationTriggersLoadRulesZarazElementVisibilityRule) implementsConfigurationTriggersLoadRule() {
}

type ConfigurationTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ConfigurationTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ConfigurationTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ConfigurationTriggersLoadRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                               `json:"selector,required"`
	JSON     configurationTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// configurationTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON contains
// the JSON metadata for the struct
// [ConfigurationTriggersLoadRulesZarazElementVisibilityRuleSettings]
type configurationTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON) RawJSON() string {
	return r.raw
}

type ConfigurationTriggersLoadRulesAction string

const (
	ConfigurationTriggersLoadRulesActionClickListener     ConfigurationTriggersLoadRulesAction = "clickListener"
	ConfigurationTriggersLoadRulesActionTimer             ConfigurationTriggersLoadRulesAction = "timer"
	ConfigurationTriggersLoadRulesActionFormSubmission    ConfigurationTriggersLoadRulesAction = "formSubmission"
	ConfigurationTriggersLoadRulesActionVariableMatch     ConfigurationTriggersLoadRulesAction = "variableMatch"
	ConfigurationTriggersLoadRulesActionScrollDepth       ConfigurationTriggersLoadRulesAction = "scrollDepth"
	ConfigurationTriggersLoadRulesActionElementVisibility ConfigurationTriggersLoadRulesAction = "elementVisibility"
)

func (r ConfigurationTriggersLoadRulesAction) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesActionClickListener, ConfigurationTriggersLoadRulesActionTimer, ConfigurationTriggersLoadRulesActionFormSubmission, ConfigurationTriggersLoadRulesActionVariableMatch, ConfigurationTriggersLoadRulesActionScrollDepth, ConfigurationTriggersLoadRulesActionElementVisibility:
		return true
	}
	return false
}

type ConfigurationTriggersLoadRulesOp string

const (
	ConfigurationTriggersLoadRulesOpContains           ConfigurationTriggersLoadRulesOp = "CONTAINS"
	ConfigurationTriggersLoadRulesOpEquals             ConfigurationTriggersLoadRulesOp = "EQUALS"
	ConfigurationTriggersLoadRulesOpStartsWith         ConfigurationTriggersLoadRulesOp = "STARTS_WITH"
	ConfigurationTriggersLoadRulesOpEndsWith           ConfigurationTriggersLoadRulesOp = "ENDS_WITH"
	ConfigurationTriggersLoadRulesOpMatchRegex         ConfigurationTriggersLoadRulesOp = "MATCH_REGEX"
	ConfigurationTriggersLoadRulesOpNotMatchRegex      ConfigurationTriggersLoadRulesOp = "NOT_MATCH_REGEX"
	ConfigurationTriggersLoadRulesOpGreaterThan        ConfigurationTriggersLoadRulesOp = "GREATER_THAN"
	ConfigurationTriggersLoadRulesOpGreaterThanOrEqual ConfigurationTriggersLoadRulesOp = "GREATER_THAN_OR_EQUAL"
	ConfigurationTriggersLoadRulesOpLessThan           ConfigurationTriggersLoadRulesOp = "LESS_THAN"
	ConfigurationTriggersLoadRulesOpLessThanOrEqual    ConfigurationTriggersLoadRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigurationTriggersLoadRulesOp) IsKnown() bool {
	switch r {
	case ConfigurationTriggersLoadRulesOpContains, ConfigurationTriggersLoadRulesOpEquals, ConfigurationTriggersLoadRulesOpStartsWith, ConfigurationTriggersLoadRulesOpEndsWith, ConfigurationTriggersLoadRulesOpMatchRegex, ConfigurationTriggersLoadRulesOpNotMatchRegex, ConfigurationTriggersLoadRulesOpGreaterThan, ConfigurationTriggersLoadRulesOpGreaterThanOrEqual, ConfigurationTriggersLoadRulesOpLessThan, ConfigurationTriggersLoadRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigurationTriggersSystem string

const (
	ConfigurationTriggersSystemPageload ConfigurationTriggersSystem = "pageload"
)

func (r ConfigurationTriggersSystem) IsKnown() bool {
	switch r {
	case ConfigurationTriggersSystemPageload:
		return true
	}
	return false
}

type ConfigurationVariable struct {
	Name string                     `json:"name,required"`
	Type ConfigurationVariablesType `json:"type,required"`
	// This field can have the runtime type of [string],
	// [ConfigurationVariablesZarazWorkerVariableValue].
	Value interface{}               `json:"value,required"`
	JSON  configurationVariableJSON `json:"-"`
	union ConfigurationVariablesUnion
}

// configurationVariableJSON contains the JSON metadata for the struct
// [ConfigurationVariable]
type configurationVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationVariableJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationVariable) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationVariable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationVariablesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationVariablesZarazStringVariable],
// [ConfigurationVariablesZarazSecretVariable],
// [ConfigurationVariablesZarazWorkerVariable].
func (r ConfigurationVariable) AsUnion() ConfigurationVariablesUnion {
	return r.union
}

// Union satisfied by [ConfigurationVariablesZarazStringVariable],
// [ConfigurationVariablesZarazSecretVariable] or
// [ConfigurationVariablesZarazWorkerVariable].
type ConfigurationVariablesUnion interface {
	implementsConfigurationVariable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationVariablesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConfigurationVariablesZarazStringVariable{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConfigurationVariablesZarazSecretVariable{}),
			DiscriminatorValue: "secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConfigurationVariablesZarazWorkerVariable{}),
			DiscriminatorValue: "worker",
		},
	)
}

type ConfigurationVariablesZarazStringVariable struct {
	Name  string                                        `json:"name,required"`
	Type  ConfigurationVariablesZarazStringVariableType `json:"type,required"`
	Value string                                        `json:"value,required"`
	JSON  configurationVariablesZarazStringVariableJSON `json:"-"`
}

// configurationVariablesZarazStringVariableJSON contains the JSON metadata for the
// struct [ConfigurationVariablesZarazStringVariable]
type configurationVariablesZarazStringVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationVariablesZarazStringVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationVariablesZarazStringVariableJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationVariablesZarazStringVariable) implementsConfigurationVariable() {}

type ConfigurationVariablesZarazStringVariableType string

const (
	ConfigurationVariablesZarazStringVariableTypeString ConfigurationVariablesZarazStringVariableType = "string"
)

func (r ConfigurationVariablesZarazStringVariableType) IsKnown() bool {
	switch r {
	case ConfigurationVariablesZarazStringVariableTypeString:
		return true
	}
	return false
}

type ConfigurationVariablesZarazSecretVariable struct {
	Name  string                                        `json:"name,required"`
	Type  ConfigurationVariablesZarazSecretVariableType `json:"type,required"`
	Value string                                        `json:"value,required"`
	JSON  configurationVariablesZarazSecretVariableJSON `json:"-"`
}

// configurationVariablesZarazSecretVariableJSON contains the JSON metadata for the
// struct [ConfigurationVariablesZarazSecretVariable]
type configurationVariablesZarazSecretVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationVariablesZarazSecretVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationVariablesZarazSecretVariableJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationVariablesZarazSecretVariable) implementsConfigurationVariable() {}

type ConfigurationVariablesZarazSecretVariableType string

const (
	ConfigurationVariablesZarazSecretVariableTypeSecret ConfigurationVariablesZarazSecretVariableType = "secret"
)

func (r ConfigurationVariablesZarazSecretVariableType) IsKnown() bool {
	switch r {
	case ConfigurationVariablesZarazSecretVariableTypeSecret:
		return true
	}
	return false
}

type ConfigurationVariablesZarazWorkerVariable struct {
	Name  string                                         `json:"name,required"`
	Type  ConfigurationVariablesZarazWorkerVariableType  `json:"type,required"`
	Value ConfigurationVariablesZarazWorkerVariableValue `json:"value,required"`
	JSON  configurationVariablesZarazWorkerVariableJSON  `json:"-"`
}

// configurationVariablesZarazWorkerVariableJSON contains the JSON metadata for the
// struct [ConfigurationVariablesZarazWorkerVariable]
type configurationVariablesZarazWorkerVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationVariablesZarazWorkerVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationVariablesZarazWorkerVariableJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationVariablesZarazWorkerVariable) implementsConfigurationVariable() {}

type ConfigurationVariablesZarazWorkerVariableType string

const (
	ConfigurationVariablesZarazWorkerVariableTypeWorker ConfigurationVariablesZarazWorkerVariableType = "worker"
)

func (r ConfigurationVariablesZarazWorkerVariableType) IsKnown() bool {
	switch r {
	case ConfigurationVariablesZarazWorkerVariableTypeWorker:
		return true
	}
	return false
}

type ConfigurationVariablesZarazWorkerVariableValue struct {
	EscapedWorkerName string                                             `json:"escapedWorkerName,required"`
	WorkerTag         string                                             `json:"workerTag,required"`
	JSON              configurationVariablesZarazWorkerVariableValueJSON `json:"-"`
}

// configurationVariablesZarazWorkerVariableValueJSON contains the JSON metadata
// for the struct [ConfigurationVariablesZarazWorkerVariableValue]
type configurationVariablesZarazWorkerVariableValueJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigurationVariablesZarazWorkerVariableValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationVariablesZarazWorkerVariableValueJSON) RawJSON() string {
	return r.raw
}

type ConfigurationVariablesType string

const (
	ConfigurationVariablesTypeString ConfigurationVariablesType = "string"
	ConfigurationVariablesTypeSecret ConfigurationVariablesType = "secret"
	ConfigurationVariablesTypeWorker ConfigurationVariablesType = "worker"
)

func (r ConfigurationVariablesType) IsKnown() bool {
	switch r {
	case ConfigurationVariablesTypeString, ConfigurationVariablesTypeSecret, ConfigurationVariablesTypeWorker:
		return true
	}
	return false
}

// Cloudflare Monitoring settings.
type ConfigurationAnalytics struct {
	// Consent purpose assigned to Monitoring.
	DefaultPurpose string `json:"defaultPurpose"`
	// Whether Advanced Monitoring reports are enabled.
	Enabled bool `json:"enabled"`
	// Session expiration time (seconds).
	SessionExpTime int64                      `json:"sessionExpTime"`
	JSON           configurationAnalyticsJSON `json:"-"`
}

// configurationAnalyticsJSON contains the JSON metadata for the struct
// [ConfigurationAnalytics]
type configurationAnalyticsJSON struct {
	DefaultPurpose apijson.Field
	Enabled        apijson.Field
	SessionExpTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigurationAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationAnalyticsJSON) RawJSON() string {
	return r.raw
}

// Consent management configuration.
type ConfigurationConsent struct {
	Enabled                bool                  `json:"enabled,required"`
	ButtonTextTranslations ButtonTextTranslation `json:"buttonTextTranslations"`
	CompanyEmail           string                `json:"companyEmail"`
	CompanyName            string                `json:"companyName"`
	CompanyStreetAddress   string                `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCSS                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ConfigurationConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ConfigurationConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	TcfCompliant             bool                                                   `json:"tcfCompliant"`
	JSON                     configurationConsentJSON                               `json:"-"`
}

// configurationConsentJSON contains the JSON metadata for the struct
// [ConfigurationConsent]
type configurationConsentJSON struct {
	Enabled                               apijson.Field
	ButtonTextTranslations                apijson.Field
	CompanyEmail                          apijson.Field
	CompanyName                           apijson.Field
	CompanyStreetAddress                  apijson.Field
	ConsentModalIntroHTML                 apijson.Field
	ConsentModalIntroHTMLWithTranslations apijson.Field
	CookieName                            apijson.Field
	CustomCSS                             apijson.Field
	CustomIntroDisclaimerDismissed        apijson.Field
	DefaultLanguage                       apijson.Field
	HideModal                             apijson.Field
	Purposes                              apijson.Field
	PurposesWithTranslations              apijson.Field
	TcfCompliant                          apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *ConfigurationConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationConsentJSON) RawJSON() string {
	return r.raw
}

type ConfigurationConsentPurpose struct {
	Description string                          `json:"description,required"`
	Name        string                          `json:"name,required"`
	JSON        configurationConsentPurposeJSON `json:"-"`
}

// configurationConsentPurposeJSON contains the JSON metadata for the struct
// [ConfigurationConsentPurpose]
type configurationConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationConsentPurposeJSON) RawJSON() string {
	return r.raw
}

type ConfigurationConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                               `json:"name,required"`
	Order int64                                           `json:"order,required"`
	JSON  configurationConsentPurposesWithTranslationJSON `json:"-"`
}

// configurationConsentPurposesWithTranslationJSON contains the JSON metadata for
// the struct [ConfigurationConsentPurposesWithTranslation]
type configurationConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationConsentPurposesWithTranslationJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey,required"`
	// General Zaraz settings.
	Settings param.Field[ConfigUpdateParamsSettings] `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools param.Field[map[string]ConfigUpdateParamsToolsUnion] `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers param.Field[map[string]ConfigUpdateParamsTriggers] `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables param.Field[map[string]ConfigUpdateParamsVariablesUnion] `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion param.Field[int64] `json:"zarazVersion,required"`
	// Cloudflare Monitoring settings.
	Analytics param.Field[ConfigUpdateParamsAnalytics] `json:"analytics"`
	// Consent management configuration.
	Consent param.Field[ConfigUpdateParamsConsent] `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// General Zaraz settings.
type ConfigUpdateParamsSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript param.Field[bool] `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher param.Field[ConfigUpdateParamsSettingsContextEnricher] `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain param.Field[string] `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce param.Field[bool] `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath param.Field[string] `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer param.Field[bool] `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress param.Field[bool] `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams param.Field[bool] `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent param.Field[bool] `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath param.Field[string] `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes param.Field[bool] `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath param.Field[string] `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath param.Field[string] `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath param.Field[string] `json:"trackPath"`
}

func (r ConfigUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ConfigUpdateParamsSettingsContextEnricher struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ConfigUpdateParamsSettingsContextEnricher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTools struct {
	BlockingTriggers param.Field[interface{}] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component     param.Field[string]      `json:"component,required"`
	DefaultFields param.Field[interface{}] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name        param.Field[string]                      `json:"name,required"`
	Permissions param.Field[interface{}]                 `json:"permissions,required"`
	Settings    param.Field[interface{}]                 `json:"settings,required"`
	Type        param.Field[ConfigUpdateParamsToolsType] `json:"type,required"`
	Actions     param.Field[interface{}]                 `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string]      `json:"defaultPurpose"`
	NeoEvents      param.Field[interface{}] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string]      `json:"vendorPolicyUrl"`
	Worker          param.Field[interface{}] `json:"worker"`
}

func (r ConfigUpdateParamsTools) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTools) implementsConfigUpdateParamsToolsUnion() {}

// Satisfied by [zaraz.ConfigUpdateParamsToolsZarazManagedComponent],
// [zaraz.ConfigUpdateParamsToolsWorker], [ConfigUpdateParamsTools].
type ConfigUpdateParamsToolsUnion interface {
	implementsConfigUpdateParamsToolsUnion()
}

type ConfigUpdateParamsToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ConfigUpdateParamsToolsZarazManagedComponentSettingsUnion] `json:"settings,required"`
	Type     param.Field[ConfigUpdateParamsToolsZarazManagedComponentType]                     `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]NeoEventParam] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]NeoEventParam] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string] `json:"vendorPolicyUrl"`
}

func (r ConfigUpdateParamsToolsZarazManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsToolsZarazManagedComponent) implementsConfigUpdateParamsToolsUnion() {}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion interface {
	ImplementsConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ConfigUpdateParamsToolsZarazManagedComponentSettingsUnion interface {
	ImplementsConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()
}

type ConfigUpdateParamsToolsZarazManagedComponentType string

const (
	ConfigUpdateParamsToolsZarazManagedComponentTypeComponent ConfigUpdateParamsToolsZarazManagedComponentType = "component"
)

func (r ConfigUpdateParamsToolsZarazManagedComponentType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsToolsZarazManagedComponentTypeComponent:
		return true
	}
	return false
}

type ConfigUpdateParamsToolsWorker struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ConfigUpdateParamsToolsWorkerDefaultFieldsUnion] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ConfigUpdateParamsToolsWorkerSettingsUnion] `json:"settings,required"`
	Type     param.Field[ConfigUpdateParamsToolsWorkerType]                     `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker param.Field[ConfigUpdateParamsToolsWorkerWorker] `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]NeoEventParam] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]NeoEventParam] `json:"neoEvents"`
	// Vendor name for TCF compliant consent modal, required for Custom Managed
	// Components and Custom HTML tool with a defaultPurpose assigned
	VendorName param.Field[string] `json:"vendorName"`
	// Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom
	// Managed Components and Custom HTML tool with a defaultPurpose assigned
	VendorPolicyURL param.Field[string] `json:"vendorPolicyUrl"`
}

func (r ConfigUpdateParamsToolsWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsToolsWorker) implementsConfigUpdateParamsToolsUnion() {}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ConfigUpdateParamsToolsWorkerDefaultFieldsUnion interface {
	ImplementsConfigUpdateParamsToolsWorkerDefaultFieldsUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ConfigUpdateParamsToolsWorkerSettingsUnion interface {
	ImplementsConfigUpdateParamsToolsWorkerSettingsUnion()
}

type ConfigUpdateParamsToolsWorkerType string

const (
	ConfigUpdateParamsToolsWorkerTypeCustomMc ConfigUpdateParamsToolsWorkerType = "custom-mc"
)

func (r ConfigUpdateParamsToolsWorkerType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsToolsWorkerTypeCustomMc:
		return true
	}
	return false
}

// Cloudflare worker that acts as a managed component
type ConfigUpdateParamsToolsWorkerWorker struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ConfigUpdateParamsToolsWorkerWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsToolsType string

const (
	ConfigUpdateParamsToolsTypeComponent ConfigUpdateParamsToolsType = "component"
	ConfigUpdateParamsToolsTypeCustomMc  ConfigUpdateParamsToolsType = "custom-mc"
)

func (r ConfigUpdateParamsToolsType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsToolsTypeComponent, ConfigUpdateParamsToolsTypeCustomMc:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggers struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules param.Field[[]ConfigUpdateParamsTriggersExcludeRuleUnion] `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules param.Field[[]ConfigUpdateParamsTriggersLoadRuleUnion] `json:"loadRules,required"`
	// Trigger name.
	Name param.Field[string] `json:"name,required"`
	// Trigger description.
	Description param.Field[string]                           `json:"description"`
	System      param.Field[ConfigUpdateParamsTriggersSystem] `json:"system"`
}

func (r ConfigUpdateParamsTriggers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRule struct {
	ID       param.Field[string]                                       `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesAction] `json:"action"`
	Match    param.Field[string]                                       `json:"match"`
	Op       param.Field[ConfigUpdateParamsTriggersExcludeRulesOp]     `json:"op"`
	Settings param.Field[interface{}]                                  `json:"settings"`
	Value    param.Field[string]                                       `json:"value"`
}

func (r ConfigUpdateParamsTriggersExcludeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

// Satisfied by [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazLoadRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazTimerRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule],
// [zaraz.ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule],
// [ConfigUpdateParamsTriggersExcludeRule].
type ConfigUpdateParamsTriggersExcludeRuleUnion interface {
	implementsConfigUpdateParamsTriggersExcludeRuleUnion()
}

type ConfigUpdateParamsTriggersExcludeRulesZarazLoadRule struct {
	ID    param.Field[string]                                                `json:"id,required"`
	Match param.Field[string]                                                `json:"match,required"`
	Op    param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                `json:"value,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains           ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEquals             ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith         ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith           ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThan           ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEquals, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThan, ConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                               `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                   `json:"selector,required"`
	Type        param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                    `json:"waitForTags,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCSS   ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath, ConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCSS:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazTimerRule struct {
	ID       param.Field[string]                                                       `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleActionTimer ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                               `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                             `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                   `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) implementsConfigUpdateParamsTriggersExcludeRuleUnion() {
}

type ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersExcludeRulesAction string

const (
	ConfigUpdateParamsTriggersExcludeRulesActionClickListener     ConfigUpdateParamsTriggersExcludeRulesAction = "clickListener"
	ConfigUpdateParamsTriggersExcludeRulesActionTimer             ConfigUpdateParamsTriggersExcludeRulesAction = "timer"
	ConfigUpdateParamsTriggersExcludeRulesActionFormSubmission    ConfigUpdateParamsTriggersExcludeRulesAction = "formSubmission"
	ConfigUpdateParamsTriggersExcludeRulesActionVariableMatch     ConfigUpdateParamsTriggersExcludeRulesAction = "variableMatch"
	ConfigUpdateParamsTriggersExcludeRulesActionScrollDepth       ConfigUpdateParamsTriggersExcludeRulesAction = "scrollDepth"
	ConfigUpdateParamsTriggersExcludeRulesActionElementVisibility ConfigUpdateParamsTriggersExcludeRulesAction = "elementVisibility"
)

func (r ConfigUpdateParamsTriggersExcludeRulesAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesActionClickListener, ConfigUpdateParamsTriggersExcludeRulesActionTimer, ConfigUpdateParamsTriggersExcludeRulesActionFormSubmission, ConfigUpdateParamsTriggersExcludeRulesActionVariableMatch, ConfigUpdateParamsTriggersExcludeRulesActionScrollDepth, ConfigUpdateParamsTriggersExcludeRulesActionElementVisibility:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersExcludeRulesOp string

const (
	ConfigUpdateParamsTriggersExcludeRulesOpContains           ConfigUpdateParamsTriggersExcludeRulesOp = "CONTAINS"
	ConfigUpdateParamsTriggersExcludeRulesOpEquals             ConfigUpdateParamsTriggersExcludeRulesOp = "EQUALS"
	ConfigUpdateParamsTriggersExcludeRulesOpStartsWith         ConfigUpdateParamsTriggersExcludeRulesOp = "STARTS_WITH"
	ConfigUpdateParamsTriggersExcludeRulesOpEndsWith           ConfigUpdateParamsTriggersExcludeRulesOp = "ENDS_WITH"
	ConfigUpdateParamsTriggersExcludeRulesOpMatchRegex         ConfigUpdateParamsTriggersExcludeRulesOp = "MATCH_REGEX"
	ConfigUpdateParamsTriggersExcludeRulesOpNotMatchRegex      ConfigUpdateParamsTriggersExcludeRulesOp = "NOT_MATCH_REGEX"
	ConfigUpdateParamsTriggersExcludeRulesOpGreaterThan        ConfigUpdateParamsTriggersExcludeRulesOp = "GREATER_THAN"
	ConfigUpdateParamsTriggersExcludeRulesOpGreaterThanOrEqual ConfigUpdateParamsTriggersExcludeRulesOp = "GREATER_THAN_OR_EQUAL"
	ConfigUpdateParamsTriggersExcludeRulesOpLessThan           ConfigUpdateParamsTriggersExcludeRulesOp = "LESS_THAN"
	ConfigUpdateParamsTriggersExcludeRulesOpLessThanOrEqual    ConfigUpdateParamsTriggersExcludeRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigUpdateParamsTriggersExcludeRulesOp) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersExcludeRulesOpContains, ConfigUpdateParamsTriggersExcludeRulesOpEquals, ConfigUpdateParamsTriggersExcludeRulesOpStartsWith, ConfigUpdateParamsTriggersExcludeRulesOpEndsWith, ConfigUpdateParamsTriggersExcludeRulesOpMatchRegex, ConfigUpdateParamsTriggersExcludeRulesOpNotMatchRegex, ConfigUpdateParamsTriggersExcludeRulesOpGreaterThan, ConfigUpdateParamsTriggersExcludeRulesOpGreaterThanOrEqual, ConfigUpdateParamsTriggersExcludeRulesOpLessThan, ConfigUpdateParamsTriggersExcludeRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRule struct {
	ID       param.Field[string]                                    `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesAction] `json:"action"`
	Match    param.Field[string]                                    `json:"match"`
	Op       param.Field[ConfigUpdateParamsTriggersLoadRulesOp]     `json:"op"`
	Settings param.Field[interface{}]                               `json:"settings"`
	Value    param.Field[string]                                    `json:"value"`
}

func (r ConfigUpdateParamsTriggersLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {}

// Satisfied by [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazLoadRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazTimerRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule],
// [zaraz.ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule],
// [ConfigUpdateParamsTriggersLoadRule].
type ConfigUpdateParamsTriggersLoadRuleUnion interface {
	implementsConfigUpdateParamsTriggersLoadRuleUnion()
}

type ConfigUpdateParamsTriggersLoadRulesZarazLoadRule struct {
	ID    param.Field[string]                                             `json:"id,required"`
	Match param.Field[string]                                             `json:"match,required"`
	Op    param.Field[ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                             `json:"value,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazLoadRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains           ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEquals             ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpStartsWith         ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEndsWith           ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex         ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan        ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThan           ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEquals, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpStartsWith, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEndsWith, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThan, ConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                            `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                `json:"selector,required"`
	Type        param.Field[ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                 `json:"waitForTags,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCSS   ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath, ConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCSS:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazTimerRule struct {
	ID       param.Field[string]                                                    `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazTimerRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleActionTimer ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction = "timer"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleActionTimer:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                             `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                            `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                          `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                `json:"id,required"`
	Action   param.Field[ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) implementsConfigUpdateParamsTriggersLoadRuleUnion() {
}

type ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

func (r ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsTriggersLoadRulesAction string

const (
	ConfigUpdateParamsTriggersLoadRulesActionClickListener     ConfigUpdateParamsTriggersLoadRulesAction = "clickListener"
	ConfigUpdateParamsTriggersLoadRulesActionTimer             ConfigUpdateParamsTriggersLoadRulesAction = "timer"
	ConfigUpdateParamsTriggersLoadRulesActionFormSubmission    ConfigUpdateParamsTriggersLoadRulesAction = "formSubmission"
	ConfigUpdateParamsTriggersLoadRulesActionVariableMatch     ConfigUpdateParamsTriggersLoadRulesAction = "variableMatch"
	ConfigUpdateParamsTriggersLoadRulesActionScrollDepth       ConfigUpdateParamsTriggersLoadRulesAction = "scrollDepth"
	ConfigUpdateParamsTriggersLoadRulesActionElementVisibility ConfigUpdateParamsTriggersLoadRulesAction = "elementVisibility"
)

func (r ConfigUpdateParamsTriggersLoadRulesAction) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesActionClickListener, ConfigUpdateParamsTriggersLoadRulesActionTimer, ConfigUpdateParamsTriggersLoadRulesActionFormSubmission, ConfigUpdateParamsTriggersLoadRulesActionVariableMatch, ConfigUpdateParamsTriggersLoadRulesActionScrollDepth, ConfigUpdateParamsTriggersLoadRulesActionElementVisibility:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersLoadRulesOp string

const (
	ConfigUpdateParamsTriggersLoadRulesOpContains           ConfigUpdateParamsTriggersLoadRulesOp = "CONTAINS"
	ConfigUpdateParamsTriggersLoadRulesOpEquals             ConfigUpdateParamsTriggersLoadRulesOp = "EQUALS"
	ConfigUpdateParamsTriggersLoadRulesOpStartsWith         ConfigUpdateParamsTriggersLoadRulesOp = "STARTS_WITH"
	ConfigUpdateParamsTriggersLoadRulesOpEndsWith           ConfigUpdateParamsTriggersLoadRulesOp = "ENDS_WITH"
	ConfigUpdateParamsTriggersLoadRulesOpMatchRegex         ConfigUpdateParamsTriggersLoadRulesOp = "MATCH_REGEX"
	ConfigUpdateParamsTriggersLoadRulesOpNotMatchRegex      ConfigUpdateParamsTriggersLoadRulesOp = "NOT_MATCH_REGEX"
	ConfigUpdateParamsTriggersLoadRulesOpGreaterThan        ConfigUpdateParamsTriggersLoadRulesOp = "GREATER_THAN"
	ConfigUpdateParamsTriggersLoadRulesOpGreaterThanOrEqual ConfigUpdateParamsTriggersLoadRulesOp = "GREATER_THAN_OR_EQUAL"
	ConfigUpdateParamsTriggersLoadRulesOpLessThan           ConfigUpdateParamsTriggersLoadRulesOp = "LESS_THAN"
	ConfigUpdateParamsTriggersLoadRulesOpLessThanOrEqual    ConfigUpdateParamsTriggersLoadRulesOp = "LESS_THAN_OR_EQUAL"
)

func (r ConfigUpdateParamsTriggersLoadRulesOp) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersLoadRulesOpContains, ConfigUpdateParamsTriggersLoadRulesOpEquals, ConfigUpdateParamsTriggersLoadRulesOpStartsWith, ConfigUpdateParamsTriggersLoadRulesOpEndsWith, ConfigUpdateParamsTriggersLoadRulesOpMatchRegex, ConfigUpdateParamsTriggersLoadRulesOpNotMatchRegex, ConfigUpdateParamsTriggersLoadRulesOpGreaterThan, ConfigUpdateParamsTriggersLoadRulesOpGreaterThanOrEqual, ConfigUpdateParamsTriggersLoadRulesOpLessThan, ConfigUpdateParamsTriggersLoadRulesOpLessThanOrEqual:
		return true
	}
	return false
}

type ConfigUpdateParamsTriggersSystem string

const (
	ConfigUpdateParamsTriggersSystemPageload ConfigUpdateParamsTriggersSystem = "pageload"
)

func (r ConfigUpdateParamsTriggersSystem) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsTriggersSystemPageload:
		return true
	}
	return false
}

type ConfigUpdateParamsVariables struct {
	Name  param.Field[string]                          `json:"name,required"`
	Type  param.Field[ConfigUpdateParamsVariablesType] `json:"type,required"`
	Value param.Field[interface{}]                     `json:"value,required"`
}

func (r ConfigUpdateParamsVariables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsVariables) implementsConfigUpdateParamsVariablesUnion() {}

// Satisfied by [zaraz.ConfigUpdateParamsVariablesZarazStringVariable],
// [zaraz.ConfigUpdateParamsVariablesZarazSecretVariable],
// [zaraz.ConfigUpdateParamsVariablesZarazWorkerVariable],
// [ConfigUpdateParamsVariables].
type ConfigUpdateParamsVariablesUnion interface {
	implementsConfigUpdateParamsVariablesUnion()
}

type ConfigUpdateParamsVariablesZarazStringVariable struct {
	Name  param.Field[string]                                             `json:"name,required"`
	Type  param.Field[ConfigUpdateParamsVariablesZarazStringVariableType] `json:"type,required"`
	Value param.Field[string]                                             `json:"value,required"`
}

func (r ConfigUpdateParamsVariablesZarazStringVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsVariablesZarazStringVariable) implementsConfigUpdateParamsVariablesUnion() {
}

type ConfigUpdateParamsVariablesZarazStringVariableType string

const (
	ConfigUpdateParamsVariablesZarazStringVariableTypeString ConfigUpdateParamsVariablesZarazStringVariableType = "string"
)

func (r ConfigUpdateParamsVariablesZarazStringVariableType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsVariablesZarazStringVariableTypeString:
		return true
	}
	return false
}

type ConfigUpdateParamsVariablesZarazSecretVariable struct {
	Name  param.Field[string]                                             `json:"name,required"`
	Type  param.Field[ConfigUpdateParamsVariablesZarazSecretVariableType] `json:"type,required"`
	Value param.Field[string]                                             `json:"value,required"`
}

func (r ConfigUpdateParamsVariablesZarazSecretVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsVariablesZarazSecretVariable) implementsConfigUpdateParamsVariablesUnion() {
}

type ConfigUpdateParamsVariablesZarazSecretVariableType string

const (
	ConfigUpdateParamsVariablesZarazSecretVariableTypeSecret ConfigUpdateParamsVariablesZarazSecretVariableType = "secret"
)

func (r ConfigUpdateParamsVariablesZarazSecretVariableType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsVariablesZarazSecretVariableTypeSecret:
		return true
	}
	return false
}

type ConfigUpdateParamsVariablesZarazWorkerVariable struct {
	Name  param.Field[string]                                              `json:"name,required"`
	Type  param.Field[ConfigUpdateParamsVariablesZarazWorkerVariableType]  `json:"type,required"`
	Value param.Field[ConfigUpdateParamsVariablesZarazWorkerVariableValue] `json:"value,required"`
}

func (r ConfigUpdateParamsVariablesZarazWorkerVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigUpdateParamsVariablesZarazWorkerVariable) implementsConfigUpdateParamsVariablesUnion() {
}

type ConfigUpdateParamsVariablesZarazWorkerVariableType string

const (
	ConfigUpdateParamsVariablesZarazWorkerVariableTypeWorker ConfigUpdateParamsVariablesZarazWorkerVariableType = "worker"
)

func (r ConfigUpdateParamsVariablesZarazWorkerVariableType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsVariablesZarazWorkerVariableTypeWorker:
		return true
	}
	return false
}

type ConfigUpdateParamsVariablesZarazWorkerVariableValue struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ConfigUpdateParamsVariablesZarazWorkerVariableValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsVariablesType string

const (
	ConfigUpdateParamsVariablesTypeString ConfigUpdateParamsVariablesType = "string"
	ConfigUpdateParamsVariablesTypeSecret ConfigUpdateParamsVariablesType = "secret"
	ConfigUpdateParamsVariablesTypeWorker ConfigUpdateParamsVariablesType = "worker"
)

func (r ConfigUpdateParamsVariablesType) IsKnown() bool {
	switch r {
	case ConfigUpdateParamsVariablesTypeString, ConfigUpdateParamsVariablesTypeSecret, ConfigUpdateParamsVariablesTypeWorker:
		return true
	}
	return false
}

// Cloudflare Monitoring settings.
type ConfigUpdateParamsAnalytics struct {
	// Consent purpose assigned to Monitoring.
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// Whether Advanced Monitoring reports are enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Session expiration time (seconds).
	SessionExpTime param.Field[int64] `json:"sessionExpTime"`
}

func (r ConfigUpdateParamsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Consent management configuration.
type ConfigUpdateParamsConsent struct {
	Enabled                param.Field[bool]                       `json:"enabled,required"`
	ButtonTextTranslations param.Field[ButtonTextTranslationParam] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                     `json:"companyEmail"`
	CompanyName            param.Field[string]                     `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                     `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                     `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations param.Field[map[string]string] `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            param.Field[string]            `json:"cookieName"`
	CustomCSS                             param.Field[string]            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        param.Field[bool]              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       param.Field[string]            `json:"defaultLanguage"`
	HideModal                             param.Field[bool]              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes param.Field[map[string]ConfigUpdateParamsConsentPurposes] `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations param.Field[map[string]ConfigUpdateParamsConsentPurposesWithTranslations] `json:"purposesWithTranslations"`
	TcfCompliant             param.Field[bool]                                                         `json:"tcfCompliant"`
}

func (r ConfigUpdateParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsConsentPurposes struct {
	Description param.Field[string] `json:"description,required"`
	Name        param.Field[string] `json:"name,required"`
}

func (r ConfigUpdateParamsConsentPurposes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsConsentPurposesWithTranslations struct {
	// Object where keys are language codes
	Description param.Field[map[string]string] `json:"description,required"`
	// Object where keys are language codes
	Name  param.Field[map[string]string] `json:"name,required"`
	Order param.Field[int64]             `json:"order,required"`
}

func (r ConfigUpdateParamsConsentPurposesWithTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []ConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result Configuration `json:"result,required"`
	// Whether the API call was successful
	Success bool                             `json:"success,required"`
	JSON    configUpdateResponseEnvelopeJSON `json:"-"`
}

// configUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelope]
type configUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ConfigUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             configUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// configUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelopeErrors]
type configUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    configUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseEnvelopeErrorsSource]
type configUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ConfigUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             configUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// configUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseEnvelopeMessages]
type configUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    configUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ConfigUpdateResponseEnvelopeMessagesSource]
type configUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []ConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result Configuration `json:"result,required"`
	// Whether the API call was successful
	Success bool                          `json:"success,required"`
	JSON    configGetResponseEnvelopeJSON `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           ConfigGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             configGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// configGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeErrors]
type configGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    configGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// configGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ConfigGetResponseEnvelopeErrorsSource]
type configGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           ConfigGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             configGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// configGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeMessages]
type configGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    configGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// configGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ConfigGetResponseEnvelopeMessagesSource]
type configGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
