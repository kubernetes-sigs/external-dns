package ipam

// AddrStatus is the status of an address.
type AddrStatus string

// A list of valid address statuses.
// The default status is StatusPlanned.
const (
	StatusPlanned  AddrStatus = "planned"
	StatusAssigned AddrStatus = "assigned"
)

// Address wraps an NS1 /ipam/address resource.
type Address struct {
	ID            int                    `json:"id"`
	Desc          string                 `json:"desc,omitempty"`
	Prefix        string                 `json:"prefix"`
	Name          string                 `json:"name,omitempty"`
	Network       int                    `json:"network_id,omitempty"`
	Status        AddrStatus             `json:"status,omitempty"`
	InheritedTags string                 `json:"inherited_tags,omitempty"`
	Total         string                 `json:"total_addresses"`
	Children      int                    `json:"children"`
	Free          string                 `json:"free_addresses"`
	Used          string                 `json:"used_addresses"`
	KVPS          map[string]interface{} `json:"kvps,omitempty"`
	Tags          map[string]interface{} `json:"tags"`
	DHCPScoped    bool                   `json:"dhcp_scoped"`
	Parent        int                    `json:"parent_id"`
}
