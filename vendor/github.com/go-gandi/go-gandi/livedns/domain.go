package livedns

import "github.com/go-gandi/go-gandi/internal/client"

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

// UpdateDomainRequest contains the params for the UpdateDomain method
type UpdateDomainRequest struct {
	AutomaticSnapshots *bool `json:"automatic_snapshots,omitempty"`
}

// ListDomains lists all domains
func (g *LiveDNS) ListDomains() (domains []Domain, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

// CreateDomain adds a domain to a zone
func (g *LiveDNS) CreateDomain(fqdn string, ttl int) (response client.StandardResponse, err error) {
	_, err = g.client.Post("domains", createDomainRequest{FQDN: fqdn, Zone: zone{TTL: ttl}}, &response)
	return
}

// GetDomain returns a domain
func (g *LiveDNS) GetDomain(fqdn string) (domain Domain, err error) {
	_, err = g.client.Get("domains/"+fqdn, nil, &domain)
	return
}

// UpdateDomain changes the zone associated to a domain
func (g *LiveDNS) UpdateDomain(fqdn string, details UpdateDomainRequest) (response client.StandardResponse, err error) {
	_, err = g.client.Patch("domains/"+fqdn, details, &response)
	return
}

// GetDomainNS returns the list of the nameservers for a domain
func (g *LiveDNS) GetDomainNS(fqdn string) (ns []string, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/nameservers", nil, &ns)
	return
}
