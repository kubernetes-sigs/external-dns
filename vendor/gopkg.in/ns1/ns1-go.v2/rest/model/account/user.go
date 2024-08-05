package account

// User wraps an NS1 /account/users resource
type User struct {
	// Read-only fields
	LastAccess float64 `json:"last_access"`
	Created    float64 `json:"created"`

	Name                 string               `json:"name"`
	Username             string               `json:"username"`
	Email                string               `json:"email"`
	TeamIDs              []string             `json:"teams"`
	Notify               NotificationSettings `json:"notify"`
	Permissions          PermissionsMap       `json:"permissions"`
	IPWhitelist          []string             `json:"ip_whitelist"`
	IPWhitelistStrict    bool                 `json:"ip_whitelist_strict"`
	TwoFactorAuthEnabled bool                 `json:"2fa_enabled"`
	InviteToken          string               `json:"invite_token,omitempty"`
	SharedAuth           `json:"shared_auth"`
}

// NotificationSettings wraps a User's "notify" attribute
type NotificationSettings struct {
	Billing bool `json:"billing"`
}

// SharedAuth wraps the shared auth object on a User.
type SharedAuth struct {
	SAML `json:"saml"`
}

// SAML wraps the SAML object in SharedAuth.
type SAML struct {
	SSO bool `json:"sso"`
	IDP `json:"idp"`
}

// IDP wraps the IDP object in SAML.
type IDP struct {
	UseMetadataURL *bool   `json:"use_metadata_url"`
	MetadataURL    *string `json:"metadata_url"`
	MetadataFile   *string `json:"metadata_file"`
	Provider       *string `json:"provider"`
}
