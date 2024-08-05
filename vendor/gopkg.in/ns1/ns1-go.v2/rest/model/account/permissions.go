package account

// PermissionsMap wraps a User's "permissions" attribute
type PermissionsMap struct {
	DNS        PermissionsDNS        `json:"dns"`
	Data       PermissionsData       `json:"data"`
	Account    PermissionsAccount    `json:"account"`
	Monitoring PermissionsMonitoring `json:"monitoring"`
	Security   *PermissionsSecurity  `json:"security,omitempty"`

	// DHCP and IPAM are only relevant for DDI and should not be provided in managed.
	DHCP *PermissionsDHCP `json:"dhcp,omitempty"`
	IPAM *PermissionsIPAM `json:"ipam,omitempty"`
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
}

// PermissionsDHCP wraps a User's "permissions.dhcp" attribute for DDI.
type PermissionsDHCP struct {
	ManageDHCP bool `json:"manage_dhcp"`
	ViewDHCP   bool `json:"view_dhcp"`
	// The fields below are only relevant in DDI v2.5+
	TagsAllow *[]AuthTag `json:"tags_allow,omitempty"`
	TagsDeny  *[]AuthTag `json:"tags_deny,omitempty"`
}

// PermissionsIPAM wraps a User's "permissions.ipam" attribute for DDI.
type PermissionsIPAM struct {
	ManageIPAM bool `json:"manage_ipam"`
	ViewIPAM   bool `json:"view_ipam"`
	// The fields below are only relevant in DDI v2.5+
	TagsAllow *[]AuthTag `json:"tags_allow,omitempty"`
	TagsDeny  *[]AuthTag `json:"tags_deny,omitempty"`
}

// AuthTag wraps the tags used in "tags_allow" and "tags_deny" in DDI and IPAM permissions in DDI.
// Tag Names must start with prefix "auth:"
type AuthTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PermissionsRecord wraps a User's "permissions.record" attribute
type PermissionsRecord struct {
	Domain     string `json:"domain"`
	Subdomains bool   `json:"include_subdomains"`
	Zone       string `json:"zone"`
	RecordType string `json:"type"`
}
