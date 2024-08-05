package dns

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/data"
	"gopkg.in/ns1/ns1-go.v2/rest/model/filter"
)

// Record wraps an NS1 /zone/{zone}/{domain}/{type} resource
type Record struct {
	Meta *data.Meta `json:"meta,omitempty"`

	ID                     string `json:"id,omitempty"`
	Zone                   string `json:"zone"`
	Domain                 string `json:"domain"`
	Type                   string `json:"type"`
	Link                   string `json:"link,omitempty"`
	TTL                    int    `json:"ttl,omitempty"`
	OverrideTTL            *bool  `json:"override_ttl,omitempty"`
	OverrideAddressRecords *bool  `json:"override_address_records,omitempty"`
	UseClientSubnet        *bool  `json:"use_client_subnet,omitempty"`

	// Answers must all be of the same type as the record.
	Answers []*Answer `json:"answers"`
	// The records' filter chain.
	Filters []*filter.Filter `json:"filters"`
	// The records' regions.
	Regions data.Regions `json:"regions,omitempty"`

	// Contains the key/value tag information associated to the record
	Tags map[string]string `json:"tags,omitempty"` // Only relevant for DDI

	// List of tag key names that should not inherit from the parent zone
	BlockedTags []string `json:"blocked_tags,omitempty"` //Only relevant for DDI

	// Read-only fields
	LocalTags []string `json:"local_tags,omitempty"` // Only relevant for DDI
}

// String returns the domain rtype in string format of record
func (r *Record) String() string {
	return fmt.Sprintf("%s %s", r.Domain, r.Type)
}

// NewRecord takes a zone, domain and record type t and creates a *Record with
// UseClientSubnet: true & empty Answers.
func NewRecord(zone string, domain string, t string, tags map[string]string, blockedTags []string) *Record {
	if !strings.HasSuffix(strings.ToLower(domain), strings.ToLower(zone)) {
		domain = fmt.Sprintf("%s.%s", domain, zone)
	}
	return &Record{
		Meta:        &data.Meta{},
		Zone:        zone,
		Domain:      domain,
		Type:        t,
		Answers:     []*Answer{},
		Filters:     []*filter.Filter{},
		Regions:     data.Regions{},
		Tags:        tags,
		BlockedTags: blockedTags,
	}
}

// LinkTo sets a Record Link to an FQDN.
// to is the FQDN of the target record whose config should be used. Does
// not have to be in the same zone.
func (r *Record) LinkTo(to string) {
	r.Meta = nil
	r.Answers = []*Answer{}
	r.Link = to
}

// AddAnswer adds an answer to the record.
func (r *Record) AddAnswer(ans *Answer) {
	if r.Answers == nil {
		r.Answers = []*Answer{}
	}

	r.Answers = append(r.Answers, ans)
}

// AddFilter adds a filter to the records' filter chain(ordering of filters matters).
func (r *Record) AddFilter(fil *filter.Filter) {
	if r.Filters == nil {
		r.Filters = []*filter.Filter{}
	}

	r.Filters = append(r.Filters, fil)
}

// MarshalJSON attempts to convert any Rdata elements that cannot be passed as
// strings to the API to their correct type.
func (r *Record) MarshalJSON() ([]byte, error) {
	if r.Type == "URLFWD" {
		prepared, err := prepareURLFWDRecord(r)
		if err != nil {
			return nil, err
		}
		return json.Marshal(prepared)
	}
	// avoid an infinite loop
	type Alias Record
	return json.Marshal((*Alias)(r))
}

// returns Record with Answers as list of interface, with the Answer RData
// typed correctly for the API.
func prepareURLFWDRecord(r *Record) (interface{}, error) {
	as := []interface{}{}
	for i := range r.Answers {
		a, err := prepareURLFWDAnswer(r.Answers[i])
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	type Alias Record
	prepared := &struct {
		Answers []interface{} `json:"answers"`
		*Alias
	}{
		Answers: as,
		Alias:   (*Alias)(r),
	}
	return prepared, nil
}
