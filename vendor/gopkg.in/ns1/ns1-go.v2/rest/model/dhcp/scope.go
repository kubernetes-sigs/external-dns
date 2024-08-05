package dhcp

// Scope contains scope related data
type Scope struct {
	ID                   int             `json:"id,omitempty"`
	IDAddress            *int            `json:"address_id,omitempty"`
	IDScopeGroup         *int            `json:"scope_group_id,omitempty"`
	Options              OptionSet       `json:"options,omitempty"`
	AddressDetails       *AddressDetails `json:"address_details,omitempty"`
	ValidLifetimeSecs    *int            `json:"valid_lifetime_secs,omitempty"`
	ClientClass          *string         `json:"client_class,omitempty"`
	RequireClientClasses *[]string       `json:"require_client_classes,omitempty"`
	MatchClientID        []byte          `json:"match_client_id,omitempty"`
	Relays               []string        `json:"relays,omitempty"`
	PingCheckEnabled     *bool           `json:"ping_check_enabled,omitempty"`
	NextServer           *string         `json:"next_server,omitempty"`
	BootFileName         *string         `json:"boot_file_name,omitempty"`
	ServerHostname       *string         `json:"server_hostname,omitempty"`
	Stateless            *bool           `json:"stateless,omitempty"`

	Tags        map[string]string `json:"tags,omitempty"`
	BlockedTags []string          `json:"blocked_tags,omitempty"`
	LocalTags   []string          `json:"local_tags,omitempty"`
}

// AddressDetails scope address details
type AddressDetails struct {
	Name   string `json:"name,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}
