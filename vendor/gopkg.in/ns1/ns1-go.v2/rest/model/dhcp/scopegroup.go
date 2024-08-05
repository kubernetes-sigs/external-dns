package dhcp

// SynthesizeDNSRecords ddns configuration
type SynthesizeDNSRecords struct {
	Enabled bool `json:"enabled,omitempty"`
}

// Settings encapsulates common values between SettingsV4 and SettingsV6
type Settings struct {
	// Enabled indicates whether v4 or v6 is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// ValidLifetimeSecs how long leases given out by the server are valid
	ValidLifetimeSecs *int `json:"valid_lifetime_secs,omitempty"`

	// RenewTimerSecs length of DHCP T1 timer
	RenewTimerSecs *int `json:"renew_timer_secs,omitempty"`

	// RebindTimerSecs length of DHCP T2 timer
	RebindTimerSecs *int `json:"rebind_timer_secs,omitempty"`

	// EchoClientID https://tools.ietf.org/html/rfc6842
	EchoClientID *bool `json:"echo_client_id,omitempty"`

	// Options base dhcp options -- will be combined with later levels; should effectively be unique by option name
	Options OptionSet `json:"options"`

	SynthesizeDNSRecords *SynthesizeDNSRecords `json:"nsone-ddns,omitempty"`

	QualifyingSuffix *string `json:"qualifying_suffix,omitempty"`
	GeneratedPrefix  *string `json:"generated_prefix,omitempty"`

	// DeclineProbationPeriod how long lease will stay unavailable for assignment after DHCPDECLINE
	DeclineProbationPeriod *int `json:"decline_probation_period,omitempty"`

	// ReclaimTimerWaitTime interval between reclamation cycles in seconds
	ReclaimTimerWaitTime *int `json:"reclaim_timer_wait_time,omitempty"`

	// FlushReclaimedTimerWaitTime how often the server initiates lease reclamation in seconds
	FlushReclaimedTimerWaitTime *int `json:"flush_reclaimed_timer_wait_time,omitempty"`

	// HoldReclaimedTime how long the lease should be kept after it is reclaimed in seconds
	HoldReclaimedTime *int `json:"hold_reclaimed_time,omitempty"`

	// MaxReclaimLeases maximum number of leases to process at once (zero is unlimited)
	MaxReclaimLeases *int `json:"max_reclaim_leases,omitempty"`

	// MaxReclaimTime upper limit to the length of time a lease reclamation procedure can take
	// (in milliseconds)
	MaxReclaimTime *int `json:"max_reclaim_time,omitempty"`

	// UnwarnedReclaimCycles how many consecutive cycles must end with remaining leases before a warning
	// is printed
	UnwarnedReclaimCycles *int `json:"unwarned_reclaim_cycles,omitempty"`
}

// SettingsV4 defines those DHCPv4 settings which we expose to the user
type SettingsV4 struct {
	Settings

	MatchClientID *bool   `json:"match_client_id"`
	NextServer    *string `json:"next_server"`
	BootFileName  *string `json:"boot_file_name"`
}

// SettingsV6 defines those DHCPv6 settings which we expose to the user
type SettingsV6 struct {
	Settings

	// https://tools.ietf.org/html/rfc3315
	PreferredLifetimeSecs *int `json:"preferred_lifetime_secs,omitempty"`
}

const (
	// PingCheckProbationPeriodDefault is a default value in minutes for the probation_period in the ping check config
	PingCheckProbationPeriodDefault = 60

	// PingCheckNumPingsDefault is a default value for the num_pings in the ping check config
	PingCheckNumPingsDefault = 1

	// PingCheckWaitTimeDefault is a default value in milliseconds for the wait_time in the ping check config
	PingCheckWaitTimeDefault = 500
)

// PingCheckConf represents config for a ping check in a scope group
type PingCheckConf struct {
	Enabled         bool   `json:"enabled,omitempty"`
	NumPings        *int   `json:"num_pings,omitempty"`
	ProbationPeriod *int   `json:"probation_period,omitempty"`
	Type            string `json:"type,omitempty"`
	WaitTime        *int   `json:"wait_time,omitempty"`
}

// ScopeGroup wraps an NS1 /dhcp/scopegroup resource.
type ScopeGroup struct {
	ID             *int           `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	IDDHCPService  *int           `json:"dhcp_service_id,omitempty"`
	DHCP4          SettingsV4     `json:"dhcpv4,omitempty"`
	DHCP6          SettingsV6     `json:"dhcpv6,omitempty"`
	NetworkID      *int           `json:"network_id,omitempty"`
	ReverseDNS     *bool          `json:"reverse_dns,omitempty"`
	ClientClassIds []int          `json:"client_class_ids,omitempty"`
	PingCheck      *PingCheckConf `json:"ping_check,omitempty"`

	Tags        map[string]string `json:"tags,omitempty"`
	LocalTags   []string          `json:"local_tags"`
	BlockedTags []string          `json:"blocked_tags"`

	// TemplateConfig is read-only field
	TemplateConfig []string `json:"template_config,omitempty"`
	Template       *string  `json:"template,omitempty"`
}
