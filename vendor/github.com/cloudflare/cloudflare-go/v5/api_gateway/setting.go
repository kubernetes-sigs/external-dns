// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options          []option.RequestOption
	SchemaValidation *SettingSchemaValidationService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.SchemaValidation = NewSettingSchemaValidationService(opts...)
	return
}

type Settings struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction SettingsValidationDefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction SettingsValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               settingsJSON                               `json:"-"`
}

// settingsJSON contains the JSON metadata for the struct [Settings]
type settingsJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *Settings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsJSON) RawJSON() string {
	return r.raw
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type SettingsValidationDefaultMitigationAction string

const (
	SettingsValidationDefaultMitigationActionNone  SettingsValidationDefaultMitigationAction = "none"
	SettingsValidationDefaultMitigationActionLog   SettingsValidationDefaultMitigationAction = "log"
	SettingsValidationDefaultMitigationActionBlock SettingsValidationDefaultMitigationAction = "block"
)

func (r SettingsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingsValidationDefaultMitigationActionNone, SettingsValidationDefaultMitigationActionLog, SettingsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type SettingsValidationOverrideMitigationAction string

const (
	SettingsValidationOverrideMitigationActionNone SettingsValidationOverrideMitigationAction = "none"
)

func (r SettingsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingsValidationOverrideMitigationActionNone:
		return true
	}
	return false
}
