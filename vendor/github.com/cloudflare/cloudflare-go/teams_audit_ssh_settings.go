package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// TeamsList represents a Teams List.
type AuditSSHSettings struct {
	PublicKey string     `json:"public_key"`
	SeedUUID  string     `json:"seed_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type AuditSSHSettingsResponse struct {
	Result AuditSSHSettings `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type GetAuditSSHSettingsParams struct{}

type UpdateAuditSSHSettingsParams struct {
	PublicKey string `json:"public_key"`
}

// GetAuditSSHSettings returns the accounts zt audit ssh settings.
//
// API reference: https://api.cloudflare.com/#zero-trust-get-audit-ssh-settings
func (api *API) GetAuditSSHSettings(ctx context.Context, rc *ResourceContainer, params GetAuditSSHSettingsParams) (AuditSSHSettings, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return AuditSSHSettings{}, ResultInfo{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf("/%s/%s/gateway/audit_ssh_settings", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AuditSSHSettings{}, ResultInfo{}, err
	}

	var auditSSHSettingsResponse AuditSSHSettingsResponse
	err = json.Unmarshal(res, &auditSSHSettingsResponse)
	if err != nil {
		return AuditSSHSettings{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return auditSSHSettingsResponse.Result, auditSSHSettingsResponse.ResultInfo, nil
}

// UpdateAuditSSHSettings updates an existing zt audit ssh setting.
//
// API reference: https://api.cloudflare.com/#zero-trust-update-audit-ssh-settings
func (api *API) UpdateAuditSSHSettings(ctx context.Context, rc *ResourceContainer, params UpdateAuditSSHSettingsParams) (AuditSSHSettings, error) {
	if rc.Level != AccountRouteLevel {
		return AuditSSHSettings{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return AuditSSHSettings{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/audit_ssh_settings",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return AuditSSHSettings{}, err
	}

	var auditSSHSettingsResponse AuditSSHSettingsResponse
	err = json.Unmarshal(res, &auditSSHSettingsResponse)
	if err != nil {
		return AuditSSHSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return auditSSHSettingsResponse.Result, nil
}
