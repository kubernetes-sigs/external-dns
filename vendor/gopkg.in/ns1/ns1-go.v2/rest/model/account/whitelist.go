package account

// IPWhitelist wraps the IP whitelist.
type IPWhitelist struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
