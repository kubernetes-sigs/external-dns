package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// DNSDomain represents a domain registered within Civo's infrastructure
type DNSDomain struct {
	// The ID of the domain
	ID string `json:"id"`

	// The ID of the account
	AccountID string `json:"account_id"`

	// The Name of the domain
	Name string `json:"name"`
}

type dnsDomainConfig struct {
	Name string `json:"name"`
}

// DNSRecordType represents the allowed record types: a, cname, mx or txt
type DNSRecordType string

// DNSRecord represents a DNS record registered within Civo's infrastructure
type DNSRecord struct {
	ID          string        `json:"id"`
	AccountID   string        `json:"account_id,omitempty"`
	DNSDomainID string        `json:"domain_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Value       string        `json:"value,omitempty"`
	Type        DNSRecordType `json:"type,omitempty"`
	Priority    int           `json:"priority,omitempty"`
	TTL         int           `json:"ttl,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
}

// DNSRecordConfig describes the parameters for a new DNS record
// none of the fields are mandatory and will be automatically
// set with default values
type DNSRecordConfig struct {
	Type     DNSRecordType `json:"type"`
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Priority int           `json:"priority"`
	TTL      int           `json:"ttl"`
}

const (
	// DNSRecordTypeA represents an A record
	DNSRecordTypeA = "A"

	// DNSRecordTypeCName represents an CNAME record
	DNSRecordTypeCName = "CNAME"

	// DNSRecordTypeMX represents an MX record
	DNSRecordTypeMX = "MX"

	// DNSRecordTypeSRV represents an SRV record
	DNSRecordTypeSRV = "SRV"

	// DNSRecordTypeTXT represents an TXT record
	DNSRecordTypeTXT = "TXT"
)

var (
	// ErrDNSDomainNotFound is returned when the domain is not found
	ErrDNSDomainNotFound = fmt.Errorf("domain not found")

	// ErrDNSRecordNotFound is returned when the record is not found
	ErrDNSRecordNotFound = fmt.Errorf("record not found")
)

// ListDNSDomains returns all Domains owned by the calling API account
func (c *Client) ListDNSDomains() ([]DNSDomain, error) {
	url := "/v2/dns"

	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	var domains = make([]DNSDomain, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&domains); err != nil {
		return nil, err

	}

	return domains, nil
}

// FindDNSDomain finds a domain name by either part of the ID or part of the name
func (c *Client) FindDNSDomain(search string) (*DNSDomain, error) {
	domains, err := c.ListDNSDomains()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := DNSDomain{}

	for _, value := range domains {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// CreateDNSDomain registers a new Domain
func (c *Client) CreateDNSDomain(name string) (*DNSDomain, error) {
	url := "/v2/dns"
	d := &dnsDomainConfig{Name: name}
	body, err := c.SendPostRequest(url, d)
	if err != nil {
		return nil, decodeError(err)
	}

	var n = &DNSDomain{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(n); err != nil {
		return nil, err
	}

	return n, nil
}

// GetDNSDomain returns the DNS Domain that matches the name
func (c *Client) GetDNSDomain(name string) (*DNSDomain, error) {
	ds, err := c.ListDNSDomains()
	if err != nil {
		return nil, decodeError(err)
	}

	for _, d := range ds {
		if d.Name == name {
			return &d, nil
		}
	}

	return nil, ErrDNSDomainNotFound
}

// UpdateDNSDomain updates the provided domain with name
func (c *Client) UpdateDNSDomain(d *DNSDomain, name string) (*DNSDomain, error) {
	url := fmt.Sprintf("/v2/dns/%s", d.ID)
	dc := &dnsDomainConfig{Name: name}
	body, err := c.SendPutRequest(url, dc)
	if err != nil {
		return nil, decodeError(err)
	}

	var r = &DNSDomain{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteDNSDomain deletes the Domain that matches the name
func (c *Client) DeleteDNSDomain(d *DNSDomain) (*SimpleResponse, error) {
	url := fmt.Sprintf("/v2/dns/%s", d.ID)
	resp, err := c.SendDeleteRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// CreateDNSRecord creates a new DNS record
func (c *Client) CreateDNSRecord(domainID string, r *DNSRecordConfig) (*DNSRecord, error) {
	if len(domainID) == 0 {
		return nil, fmt.Errorf("r.DomainID is empty")
	}

	url := fmt.Sprintf("/v2/dns/%s/records", domainID)
	body, err := c.SendPostRequest(url, r)
	if err != nil {
		return nil, decodeError(err)
	}

	var record = &DNSRecord{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(record); err != nil {
		return nil, err
	}

	return record, nil
}

// ListDNSRecords returns all the records associated with domainID
func (c *Client) ListDNSRecords(dnsDomainID string) ([]DNSRecord, error) {
	url := fmt.Sprintf("/v2/dns/%s/records", dnsDomainID)
	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	var rs = make([]DNSRecord, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&rs); err != nil {
		return nil, err

	}

	return rs, nil
}

// GetDNSRecord returns the Record that matches the domain ID and domain record ID
func (c *Client) GetDNSRecord(domainID, domainRecordID string) (*DNSRecord, error) {
	rs, err := c.ListDNSRecords(domainID)
	if err != nil {
		return nil, decodeError(err)
	}

	for _, r := range rs {
		if r.ID == domainRecordID {
			return &r, nil
		}
	}

	return nil, ErrDNSRecordNotFound
}

// UpdateDNSRecord updates the DNS record
func (c *Client) UpdateDNSRecord(r *DNSRecord, rc *DNSRecordConfig) (*DNSRecord, error) {
	url := fmt.Sprintf("/v2/dns/%s/records/%s", r.DNSDomainID, r.ID)
	body, err := c.SendPutRequest(url, rc)
	if err != nil {
		return nil, decodeError(err)
	}

	var dnsRecord = &DNSRecord{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(dnsRecord); err != nil {
		return nil, err
	}

	return dnsRecord, nil
}

// DeleteDNSRecord deletes the DNS record
func (c *Client) DeleteDNSRecord(r *DNSRecord) (*SimpleResponse, error) {
	if len(r.ID) == 0 {
		err := fmt.Errorf("ID is empty")
		return nil, IDisEmptyError.wrap(err)
	}

	if len(r.DNSDomainID) == 0 {
		err := fmt.Errorf("DNSDomainID is empty")
		return nil, IDisEmptyError.wrap(err)
	}

	url := fmt.Sprintf("/v2/dns/%s/records/%s", r.DNSDomainID, r.ID)
	resp, err := c.SendDeleteRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
