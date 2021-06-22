package govultr

// IPv4 struct
type IPv4 struct {
	IP      string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Type    string `json:"type,omitempty"`
	Reverse string `json:"reverse,omitempty"`
}

// IPv6 struct
type IPv6 struct {
	IP          string `json:"ip,omitempty"`
	Network     string `json:"network,omitempty"`
	NetworkSize int    `json:"network_size,omitempty"`
	Type        string `json:"type,omitempty"`
}

type ipBase struct {
	IPv4s []IPv4 `json:"ipv4s,omitempty"`
	IPv6s []IPv6 `json:"ipv6s,omitempty"`
	Meta  *Meta  `json:"meta"`
}
