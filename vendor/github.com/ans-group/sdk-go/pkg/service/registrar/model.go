package registrar

import "github.com/ans-group/sdk-go/pkg/connection"

// Domain represents a UKFast domain
type Domain struct {
	Name         string          `json:"name"`
	Status       string          `json:"status"`
	Registrar    string          `json:"registrar"`
	RegisteredAt connection.Date `json:"registered_at"`
	UpdatedAt    connection.Date `json:"updated_at"`
	RenewalAt    connection.Date `json:"renewal_at"`
	AutoRenew    bool            `json:"auto_renew"`
	WHOISPrivacy bool            `json:"whois_privacy"`
}

// Nameserver represents a nameserver
type Nameserver struct {
	Host string               `json:"host"`
	IP   connection.IPAddress `json:"ip"`
}

// Whois represents WHOIS information
type Whois struct {
	Name        string              `json:"name"`
	Status      []string            `json:"status"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
	ExpiresAt   connection.DateTime `json:"expires_at"`
	Nameservers []Nameserver        `json:"nameservers"`
	Registrar   Registrar           `json:"registrar"`
}

// Registrar represents registrar details
type Registrar struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Tag  string `json:"tag"`
}
