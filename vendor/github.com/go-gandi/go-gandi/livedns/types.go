package livedns

import (
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
)

// LiveDNS is the API client to the Gandi v5 LiveDNS API
type LiveDNS struct {
	client client.Gandi
}

// Domain represents a DNS domain
type Domain struct {
	FQDN               string `json:"fqdn,omitempty"`
	DomainHref         string `json:"domain_href,omitempty"`
	DomainKeysHref     string `json:"domain_keys_href,omitempty"`
	DomainRecordsHref  string `json:"domain_records_href,omitempty"`
	AutomaticSnapshots *bool  `json:"automatic_snapshots,omitempty"`
}

type zone struct {
	TTL int `json:"ttl"`
}

type createDomainRequest struct {
	FQDN string `json:"fqdn"`
	Zone zone   `json:"zone,omitempty"`
}

// Tsig contains tsig data (no kidding!)
type Tsig struct {
	KeyName       string      `json:"key_name,omitempty"`
	Secret        string      `json:"secret,omitempty"`
	UUID          string      `json:"uuid,omitempty"`
	AxfrTsigURL   string      `json:"axfr_tsig_url,omitempty"`
	ConfigSamples interface{} `json:"config_samples,omitempty"`
}

// DomainRecord represents a DNS Record
type DomainRecord struct {
	RrsetType   string   `json:"rrset_type,omitempty"`
	RrsetTTL    int      `json:"rrset_ttl,omitempty"`
	RrsetName   string   `json:"rrset_name,omitempty"`
	RrsetHref   string   `json:"rrset_href,omitempty"`
	RrsetValues []string `json:"rrset_values,omitempty"`
}

// SigningKey holds data about a DNSSEC signing key
type SigningKey struct {
	Status        string `json:"status,omitempty"`
	UUID          string `json:"id,omitempty"`
	Algorithm     int    `json:"algorithm,omitempty"`
	Deleted       *bool  `json:"deleted,omitempty"`
	AlgorithmName string `json:"algorithm_name,omitempty"`
	FQDN          string `json:"fqdn,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	DS            string `json:"ds,omitempty"`
	PublicKey     string `json:"public_key,omitempty"`
	Tag           int    `json:"tag,omitempty"`
	KeyHref       string `json:"key_href,omitempty"`
}

type configSamples struct {
	Bind     string `json:"bind,omitempty"`
	Knot     string `json:"knot,omitempty"`
	NSD      string `json:"nsd,omitempty"`
	PowerDNS string `json:"powerdns,omitempty"`
}

// TSIGKey describes the TSIG key associated with an AXFR secondary
type TSIGKey struct {
	KeyHREF       string        `json:"href,omitempty"`
	ID            string        `json:"id,omitempty"`
	KeyName       string        `json:"key_name,omitempty"`
	Secret        string        `json:"secret,omitempty"`
	ConfigSamples configSamples `json:"config_samples,omitempty"`
}

// Snapshot represents a point in time record of a domain
type Snapshot struct {
	Automatic    *bool          `json:"automatic,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	ID           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	SnapshotHREF string         `json:"snapshot_href,omitempty"`
	ZoneData     []DomainRecord `json:"zone_data,omitempty"`
}

// UpdateDomainRequest contains the params for the UpdateDomain method
type UpdateDomainRequest struct {
	AutomaticSnapshots *bool `json:"automatic_snapshots,omitempty"`
}
