package dns

import "encoding/json"

type ZoneSearchResult struct {
	Next         string         `json:"next"`
	Limit        int            `json:"limit"`
	TotalResults int            `json:"total_results"`
	Results      []*ZoneSummary `json:"results"`
}

type ZoneSummary struct {
	FQDN   string `json:"fqdn"`
	Handle string `json:"handle"`
}

// New correct types
type RecordSearchResult struct {
	Next         string          `json:"next"`
	Limit        int             `json:"limit"`
	TotalResults int             `json:"total_results"`
	Results      []*RecordSearch `json:"results"`
}

type RecordSearch struct {
	Domain     string          `json:"domain"`
	Type       string          `json:"type"`
	TTL        int             `json:"ttl"`
	ZoneFQDN   string          `json:"zone_fqdn"`
	ZoneHandle string          `json:"zone_handle"`
	Answers    []*SearchAnswer `json:"answers"`
}

type SearchAnswer struct {
	// Answer json.RawMessage `json:"answer"`
	Raw   json.RawMessage `json:"answer"`
	Rdata []string        `json:"-"`
}
