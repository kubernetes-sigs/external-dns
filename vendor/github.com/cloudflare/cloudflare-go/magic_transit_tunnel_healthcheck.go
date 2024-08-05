package cloudflare

// MagicTransitTunnelHealthcheck contains information about a tunnel health check.
type MagicTransitTunnelHealthcheck struct {
<<<<<<< HEAD
	Enabled bool   `json:"enabled"`
	Target  string `json:"target,omitempty"`
	Type    string `json:"type,omitempty"`
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	Enabled   bool   `json:"enabled"`
	Target    string `json:"target,omitempty"`
	Type      string `json:"type,omitempty"`
	Rate      string `json:"rate,omitempty"`
	Direction string `json:"direction,omitempty"`
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
