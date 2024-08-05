package dns

import (
	"encoding/json"
	"fmt"
)

// ZoneDNSSEC wraps an NS1 /zone/{zone}/dnssec resource
type ZoneDNSSEC struct {
	Zone       string      `json:"zone,omitempty"`
	Keys       *Keys       `json:"keys,omitempty"`
	Delegation *Delegation `json:"delegation,omitempty"`
}

// Keys holds a list of DNS Keys and a TTL
type Keys struct {
	DNSKey []*Key `json:"dnskey,omitempty"`
	TTL    int    `json:"ttl,omitempty"`
}

// Delegation holds a list of DNS Keys, a list of DS Keys, and a TTL
type Delegation struct {
	DNSKey []*Key `json:"dnskey,omitempty"`
	DS     []*Key `json:"ds,omitempty"`
	TTL    int    `json:"ttl,omitempty"`
}

// Key holds a DNS key
type Key struct {
	Flags     string
	Protocol  string
	Algorithm string
	PublicKey string
}

func (d ZoneDNSSEC) String() string {
	return fmt.Sprintf("%s", d.Zone)
}

// UnmarshalJSON parses a Key from a list into a struct
func (k *Key) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&k.Flags, &k.Protocol, &k.Algorithm, &k.PublicKey}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if l := len(tmp); l != 4 {
		return fmt.Errorf("wrong number of fields in Key: %d != 4", l)
	}
	return nil
}
