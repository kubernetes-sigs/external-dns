package livedns

import "github.com/go-gandi/go-gandi/internal/client"

// SigningKey holds data about a DNSSEC signing key
type SigningKey struct {
	Status        string `json:"status,omitempty"`
	UUID          string `json:"id,omitempty"`
	Algorithm     int    `json:"algorithm,omitempty"`
	Deleted       *bool  `json:"deleted"`
	AlgorithmName string `json:"algorithm_name,omitempty"`
	FQDN          string `json:"fqdn,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	DS            string `json:"ds,omitempty"`
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

// GetTSIGKeys retrieves all the TSIG keys for the account
func (g *LiveDNS) GetTSIGKeys() (response []TSIGKey, err error) {
	_, err = g.client.Get("axfr/tsig", nil, &response)
	return
}

// GetTSIGKey retrieves the specified TSIG key
func (g *LiveDNS) GetTSIGKey(id string) (response TSIGKey, err error) {
	_, err = g.client.Get("axfr/tsig/"+id, nil, &response)
	return
}

// CreateTSIGKey creates a TSIG key
func (g *LiveDNS) CreateTSIGKey(fqdn string) (response TSIGKey, err error) {
	_, err = g.client.Post("axfr/tsig", nil, &response)
	return
}

// GetDomainTSIGKeys retrieves the specified TSIG key
func (g *LiveDNS) GetDomainTSIGKeys(fqdn string) (response []TSIGKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/axfr/tsig", nil, &response)
	return
}

// AssociateTSIGKeyWithDomain retrieves the specified TSIG key
func (g *LiveDNS) AssociateTSIGKeyWithDomain(fqdn string, id string) (response client.StandardResponse, err error) {
	_, err = g.client.Put("domains/"+fqdn+"/axfr/tsig/"+id, nil, &response)
	return
}

// RemoveTSIGKeyFromDomain retrieves the specified TSIG key
func (g *LiveDNS) RemoveTSIGKeyFromDomain(fqdn string, id string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/axfr/tsig/"+id, nil, nil)
	return
}

// SignDomain creates a DNSKEY and asks Gandi servers to automatically sign the domain
func (g *LiveDNS) SignDomain(fqdn string) (response client.StandardResponse, err error) {
	f := SigningKey{Flags: 257}
	_, err = g.client.Post("domains/"+fqdn+"/keys", f, &response)
	return
}

// GetDomainKeys returns data about the signing keys created for a domain
func (g *LiveDNS) GetDomainKeys(fqdn string) (keys []SigningKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/keys", nil, &keys)
	return
}

// GetDomainKey deletes a signing key from a domain
func (g *LiveDNS) GetDomainKey(fqdn, uuid string) (key SigningKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/keys/"+uuid, nil, &key)
	return
}

// DeleteDomainKey deletes a signing key from a domain
func (g *LiveDNS) DeleteDomainKey(fqdn, uuid string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/keys/"+uuid, nil, nil)
	return
}

// UpdateDomainKey updates a signing key for a domain (only the deleted status, actually...)
func (g *LiveDNS) UpdateDomainKey(fqdn, uuid string, deleted bool) (err error) {
	_, err = g.client.Put("domains/"+fqdn+"/keys/"+uuid, SigningKey{Deleted: &deleted}, nil)
	return
}
