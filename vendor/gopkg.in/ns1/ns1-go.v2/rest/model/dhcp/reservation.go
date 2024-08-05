package dhcp

// Reservation contains reservation related data
type Reservation struct {
	ID             *int            `json:"id,omitempty"`
	IDAddress      *int            `json:"address_id,omitempty"`
	IDScopeGroup   *int            `json:"scope_group_id,omitempty"`
	Mac            string          `json:"mac,omitempty"`
	AddressDetails *AddressDetails `json:"address_details,omitempty"`
	Identifier     *Identifier     `json:"identifier,omitempty"`
	Options        OptionSet       `json:"options"`
	DHCPv6         *bool           `json:"dhcpv6,omitempty"`
	ClientClasses  []string        `json:"client_classes,omitempty"`
	NextServer     string          `json:"next_server,omitempty"`
	BootFileName   string          `json:"boot_file_name,omitempty"`
	ServerHostname string          `json:"server_hostname,omitempty"`

	Tags        map[string]string `json:"tags,omitempty"`
	LocalTags   []string          `json:"local_tags,omitempty"`
	BlockedTags []string          `json:"blocked_tags,omitempty"`
}

// IdentifierType is a type of the reservation identifier
type IdentifierType string

// List of available Identifier types
const (
	HWAddressType IdentifierType = "hw-address"
	CircuitIDType IdentifierType = "circuit-id"
	DUIDType      IdentifierType = "duid"
	ClientIDType  IdentifierType = "client-id"
)

// Identifier is a reservation identifier
type Identifier struct {
	Type  IdentifierType `json:"type,omitempty"`
	Value string         `json:"value,omitempty"`
}
