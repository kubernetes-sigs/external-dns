package safedns

import (
	"strconv"
	"time"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// RecordTTL represents the record TTL time in seconds
type RecordTTL int

// Time returns the record TTL time
func (r RecordTTL) Time() time.Time {
	return time.Now().Add(r.Duration())
}

// Duration returns the record TTL duration (seconds)
func (r RecordTTL) Duration() time.Duration {
	return (time.Second * time.Duration(int(r)))
}

func (r RecordTTL) String() string {
	return strconv.Itoa(int(r))
}

type RecordType string

func (s RecordType) String() string {
	return string(s)
}

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCAA   RecordType = "CAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypeSPF   RecordType = "SPF"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeNS    RecordType = "NS"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeAXFR  RecordType = "AXFR"
)

// Zone represents a SafeDNS zone
type Zone struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Record represents a SafeDNS record
type Record struct {
	ID         int                 `json:"id"`
	TemplateID int                 `json:"template_id"`
	Name       string              `json:"name"`
	Type       RecordType          `json:"type"`
	Content    string              `json:"content"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
	TTL        RecordTTL           `json:"ttl"`
	Priority   int                 `json:"priority"`
}

// Note represents a SafeDNS note
type Note struct {
	ID        int                  `json:"id"`
	ContactID int                  `json:"contact_id"`
	Notes     string               `json:"notes"`
	CreatedAt connection.DateTime  `json:"created_at"`
	IP        connection.IPAddress `json:"ip"`
}

// Template represents a SafeDNS template
type Template struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Default   bool            `json:"default"`
	CreatedAt connection.Date `json:"created_at"`
}

// Settings represents SafeDNS account settings/configuration
type Settings struct {
	ID                  int                `json:"id"`
	Email               string             `json:"email"`
	Nameservers         []Nameserver       `json:"nameservers"`
	CustomSOAAllowed    bool               `json:"custom_soa_allowed"`
	CustomBaseNSAllowed bool               `json:"custom_base_ns_allowed"`
	CustomAXFR          CustomAXFRSettings `json:"custom_axfr"`
	DelegationAllowed   bool               `json:"delegation_allowed"`
	Product             string             `json:"product"`
}

// Nameserver represents a SafeDNS nameserver
type Nameserver struct {
	Name string               `json:"name"`
	IP   connection.IPAddress `json:"ip"`
}

// CustomAXFRSettings represents SafeDNS account AXFR settings
type CustomAXFRSettings struct {
	Allowed bool                   `json:"allowed"`
	Name    []string               `json:"name"`
	IP      []connection.IPAddress `json:"ip"`
}
