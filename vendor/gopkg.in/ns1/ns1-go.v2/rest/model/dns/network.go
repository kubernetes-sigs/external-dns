package dns

// Network wraps an NS1 networks/ resource
type Network struct {
	Label     string `json:"label"`
	Name      string `json:"name"`
	NetworkID int    `json:"network_id"`
}
