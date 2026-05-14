package account

// PermissionsMap wraps a User's "permissions" attribute
type PermissionsMap struct {
	DNS        PermissionsDNS        `json:"dns"`
	Data       PermissionsData       `json:"data"`
	Account    PermissionsAccount    `json:"account"`
	Monitoring PermissionsMonitoring `json:"monitoring"`
	Security   *PermissionsSecurity  `json:"security,omitempty"`
	Redirects  PermissionsRedirects  `json:"redirects"`
	Insights   PermissionsInsights   `json:"insights"`
}

// PermissionsDNS wraps a User's "permissions.dns" attribute
type PermissionsDNS struct {
	ViewZones           bool                `json:"view_zones"`
	ManageZones         bool                `json:"manage_zones"`
	ZonesAllowByDefault bool                `json:"zones_allow_by_default"`
	ZonesDeny           []string            `json:"zones_deny"`
	ZonesAllow          []string            `json:"zones_allow"`
	RecordsAllow        []PermissionsRecord `json:"records_allow"`
	RecordsDeny         []PermissionsRecord `json:"records_deny"`
}

// PermissionsData wraps a User's "permissions.data" attribute
type PermissionsData struct {
	PushToDatafeeds   bool `json:"push_to_datafeeds"`
	ManageDatasources bool `json:"manage_datasources"`
	ManageDatafeeds   bool `json:"manage_datafeeds"`
}

// PermissionsAccount wraps a User's "permissions.account" attribute
type PermissionsAccount struct {
	ManageUsers           bool `json:"manage_users"`
	ManagePaymentMethods  bool `json:"manage_payment_methods"`
	ManagePlan            bool `json:"manage_plan"`
	ManageTeams           bool `json:"manage_teams"`
	ManageApikeys         bool `json:"manage_apikeys"`
	ManageAccountSettings bool `json:"manage_account_settings"`
	ViewActivityLog       bool `json:"view_activity_log"`
	ViewInvoices          bool `json:"view_invoices"`
	ManageIPWhitelist     bool `json:"manage_ip_whitelist"`
}

// PermissionsSecurity wraps a User's "permissions.security" attribute.
type PermissionsSecurity struct {
	ManageGlobal2FA bool `json:"manage_global_2fa"`

	// This field is only relevant for DDI and should not be set to true for managed.
	ManageActiveDirectory bool `json:"manage_active_directory,omitempty"`
}

// PermissionsMonitoring wraps a User's "permissions.monitoring" attribute
// Only relevant for the managed product.
type PermissionsMonitoring struct {
	ManageLists bool `json:"manage_lists"`
	ManageJobs  bool `json:"manage_jobs"`
	ViewJobs    bool `json:"view_jobs"`
	CreateJobs  bool `json:"create_jobs"`
	UpdateJobs  bool `json:"update_jobs"`
	DeleteJobs  bool `json:"delete_jobs"`
}

// PermissionsRecord wraps a User's "permissions.record" attribute
type PermissionsRecord struct {
	Domain     string `json:"domain"`
	Subdomains bool   `json:"include_subdomains"`
	Zone       string `json:"zone"`
	RecordType string `json:"type"`
}

// PermissionsRedirects wraps a User's "permissions.redirects" attribute
type PermissionsRedirects struct {
	ManageRedirects bool `json:"manage_redirects"`
}

// PermissionsRedirects wraps a User's "permissions.insights" attribute
type PermissionsInsights struct {
	ManageInsights bool `json:"manage_insights"`
	ViewInsights   bool `json:"view_insights"`
}
