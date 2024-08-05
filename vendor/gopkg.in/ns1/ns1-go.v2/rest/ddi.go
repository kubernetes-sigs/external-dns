package rest

import "gopkg.in/ns1/ns1-go.v2/rest/model/account"

// ddiTeam wraps an NS1 /accounts/teams resource for DDI.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type ddiTeam struct {
	ID          string                `json:"id,omitempty"`
	Name        string                `json:"name"`
	Permissions ddiPermissionsMap     `json:"permissions"`
	IPWhitelist []account.IPWhitelist `json:"ip_whitelist"`
}

// ddiUser wraps an NS1 /account/users resource for DDI.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type ddiUser struct {
	// Read-only fields
	LastAccess float64 `json:"last_access"`

	Name              string                       `json:"name"`
	Username          string                       `json:"username"`
	Email             string                       `json:"email"`
	TeamIDs           []string                     `json:"teams"`
	Notify            account.NotificationSettings `json:"notify"`
	IPWhitelist       []string                     `json:"ip_whitelist"`
	IPWhitelistStrict bool                         `json:"ip_whitelist_strict"`

	Permissions ddiPermissionsMap `json:"permissions"`
}

// ddiAPIKey wraps an NS1 /account/apikeys resource for DDI specifically.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type ddiAPIKey struct {
	// Read-only fields
	ID         string `json:"id,omitempty"`
	Key        string `json:"key,omitempty"`
	LastAccess int    `json:"last_access,omitempty"`

	Name              string   `json:"name"`
	TeamIDs           []string `json:"teams"`
	IPWhitelist       []string `json:"ip_whitelist"`
	IPWhitelistStrict bool     `json:"ip_whitelist_strict"`

	Permissions ddiPermissionsMap `json:"permissions"`
}

// ddiPermissionsMap wraps a User's "permissions" attribute for DDI.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type ddiPermissionsMap struct {
	DNS      account.PermissionsDNS  `json:"dns"`
	Data     account.PermissionsData `json:"data"`
	Account  permissionsDDIAccount   `json:"account"`
	Security permissionsDDISecurity  `json:"security"`
	DHCP     account.PermissionsDHCP `json:"dhcp"`
	IPAM     account.PermissionsIPAM `json:"ipam"`
}

// permissionsDDIAccount wraps a User's "permissions.account" attribute for DDI.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type permissionsDDIAccount struct {
	ManageUsers           bool `json:"manage_users"`
	ManageTeams           bool `json:"manage_teams"`
	ManageApikeys         bool `json:"manage_apikeys"`
	ManageAccountSettings bool `json:"manage_account_settings"`
	ViewActivityLog       bool `json:"view_activity_log"`
}

// permissionsDDISecurity wraps a User's "permissions.security" attribute for DDI.
// Used for internally mapping between DDI permissions to maintain backwards compatibility.
type permissionsDDISecurity struct {
	ManageGlobal2FA       bool `json:"manage_global_2fa"`
	ManageActiveDirectory bool `json:"manage_active_directory"`
}
