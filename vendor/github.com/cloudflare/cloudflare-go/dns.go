package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/idna"
)

// DNSRecord represents a DNS record in a zone.
type DNSRecord struct {
	CreatedOn  time.Time   `json:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Content    string      `json:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
}

// DNSRecordResponse represents the response from the DNS endpoint.
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// DNSListResponse represents the response from the list DNS records endpoint.
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// nontransitionalLookup implements the nontransitional processing as specified in
// Unicode Technical Standard 46 with almost all checkings off to maximize user freedom.
var nontransitionalLookup = idna.New(
	idna.MapForLookup(),
	idna.StrictDomainName(false),
	idna.ValidateLabels(false),
)

// toUTS46ASCII tries to convert IDNs (international domain names)
// from Unicode form to Punycode, using non-transitional process specified
// in UTS 46.
//
// Note: conversion errors are silently discarded and partial conversion
// results are used.
func toUTS46ASCII(name string) string {
	name, _ = nontransitionalLookup.ToASCII(name)
	return name
}

// CreateDNSRecord creates a DNS record for the zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
func (api *API) CreateDNSRecord(ctx context.Context, zoneID string, rr DNSRecord) (*DNSRecordResponse, error) {
	rr.Name = toUTS46ASCII(rr.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, rr)
	if err != nil {
		return nil, err
	}

	var recordResp *DNSRecordResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

// DNSRecords returns a slice of DNS records for the given zone identifier.
//
// This takes a DNSRecord to allow filtering of the results returned.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
func (api *API) DNSRecords(ctx context.Context, zoneID string, rr DNSRecord) ([]DNSRecord, error) {
	// Construct a query string
	v := url.Values{}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// Request as many records as possible per page - API max is 100
	v.Set("per_page", "100")
	if rr.Name != "" {
		v.Set("name", rr.Name)
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}

	var query string
	var records []DNSRecord
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		query = "?" + v.Encode()
		uri := "/zones/" + zoneID + "/dns_records" + query
		res, err := api.makeRequest("GET", uri, nil)
		if err != nil {
			return []DNSRecord{}, errors.Wrap(err, errMakeRequestError)
		}
		var r DNSListResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []DNSRecord{}, errors.Wrap(err, errUnmarshalError)
		}
		records = append(records, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}
		// Loop around and fetch the next page
		page++
	}
	return records, nil
}

// DNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) DNSRecord(zoneID, recordID string) (DNSRecord, error) {
	uri := "/zones/" + zoneID + "/dns_records/" + recordID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return DNSRecord{}, errors.Wrap(err, errMakeRequestError)
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(zoneID, recordID string, rr DNSRecord) error {
	rec, err := api.DNSRecord(zoneID, recordID)
	if err != nil {
		return err
	}
	// Populate the record name from the existing one if the update didn't
	// specify it.
	if rr.Name == "" {
		rr.Name = rec.Name
	}
	if rr.Type == "" {
		rr.Type = rec.Type
	}
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Request as many records as possible per page - API max is 50
	v.Set("per_page", "50")
||||||| parent of 5ce8c7613 (update vendored files)
	// Request as many records as possible per page - API max is 50
	v.Set("per_page", "50")
=======
	// Request as many records as possible per page - API max is 100
	v.Set("per_page", "100")
>>>>>>> 5ce8c7613 (update vendored files)
	if rr.Name != "" {
		v.Set("name", rr.Name)
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}

	var query string
	var records []DNSRecord
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		query = "?" + v.Encode()
		uri := "/zones/" + zoneID + "/dns_records" + query
		res, err := api.makeRequest("GET", uri, nil)
		if err != nil {
			return []DNSRecord{}, errors.Wrap(err, errMakeRequestError)
		}
		var r DNSListResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []DNSRecord{}, errors.Wrap(err, errUnmarshalError)
		}
		records = append(records, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}
		// Loop around and fetch the next page
		page++
	}
	return records, nil
}

// DNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) DNSRecord(zoneID, recordID string) (DNSRecord, error) {
	uri := "/zones/" + zoneID + "/dns_records/" + recordID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return DNSRecord{}, errors.Wrap(err, errMakeRequestError)
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(zoneID, recordID string, rr DNSRecord) error {
	rec, err := api.DNSRecord(zoneID, recordID)
	if err != nil {
		return err
	}
	// Populate the record name from the existing one if the update didn't
	// specify it.
	if rr.Name == "" {
		rr.Name = rec.Name
	}
<<<<<<< HEAD
	rr.Type = rec.Type
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	rr.Type = rec.Type
=======
	if rr.Type == "" {
		rr.Type = rec.Type
	}
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Request as many records as possible per page - API max is 50
	v.Set("per_page", "50")
||||||| parent of 6b7ce455e (update vendored files)
	// Request as many records as possible per page - API max is 50
	v.Set("per_page", "50")
=======
	// Request as many records as possible per page - API max is 100
	v.Set("per_page", "100")
>>>>>>> 6b7ce455e (update vendored files)
	if rr.Name != "" {
		v.Set("name", toUTS46ASCII(rr.Name))
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}

	var records []DNSRecord
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/zones/%s/dns_records?%s", zoneID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, err
		}
		var r DNSListResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []DNSRecord{}, errors.Wrap(err, errUnmarshalError)
		}
		records = append(records, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}
		// Loop around and fetch the next page
		page++
	}
	return records, nil
}

// DNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) DNSRecord(ctx context.Context, zoneID, recordID string) (DNSRecord, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(ctx context.Context, zoneID, recordID string, rr DNSRecord) error {
	rr.Name = toUTS46ASCII(rr.Name)

	// Populate the record name from the existing one if the update didn't
	// specify it.
	if rr.Name == "" || rr.Type == "" {
		rec, err := api.DNSRecord(ctx, zoneID, recordID)
		if err != nil {
			return err
		}

		if rr.Name == "" {
			rr.Name = rec.Name
		}
		if rr.Type == "" {
			rr.Type = rec.Type
		}
	}
<<<<<<< HEAD
	rr.Type = rec.Type
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	uri := "/zones/" + zoneID + "/dns_records/" + recordID
	res, err := api.makeRequest("PATCH", uri, rr)
||||||| parent of 6b7ce455e (update vendored files)
	rr.Type = rec.Type
	uri := "/zones/" + zoneID + "/dns_records/" + recordID
	res, err := api.makeRequest("PATCH", uri, rr)
=======
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, rr)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}

// DeleteDNSRecord deletes a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
func (api *API) DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/idna"
)

// DNSRecord represents a DNS record in a zone.
type DNSRecord struct {
	ID         string      `json:"id,omitempty"`
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Content    string      `json:"content,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	CreatedOn  time.Time   `json:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	Meta       interface{} `json:"meta,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
}

// DNSRecordResponse represents the response from the DNS endpoint.
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// DNSListResponse represents the response from the list DNS records endpoint.
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// nontransitionalLookup implements the nontransitional processing as specified in
// Unicode Technical Standard 46 with almost all checkings off to maximize user freedom.
var nontransitionalLookup = idna.New(
	idna.MapForLookup(),
	idna.StrictDomainName(false),
	idna.ValidateLabels(false),
)

// toUTS46ASCII tries to convert IDNs (international domain names)
// from Unicode form to Punycode, using non-transitional process specified
// in UTS 46.
//
// Note: conversion errors are silently discarded and partial conversion
// results are used.
func toUTS46ASCII(name string) string {
	name, _ = nontransitionalLookup.ToASCII(name)
	return name
}

// CreateDNSRecord creates a DNS record for the zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
func (api *API) CreateDNSRecord(ctx context.Context, zoneID string, rr DNSRecord) (*DNSRecordResponse, error) {
	rr.Name = toUTS46ASCII(rr.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, rr)
	if err != nil {
		return nil, err
	}

	var recordResp *DNSRecordResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return recordResp, nil
}

// DNSRecords returns a slice of DNS records for the given zone identifier.
//
// This takes a DNSRecord to allow filtering of the results returned.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
func (api *API) DNSRecords(ctx context.Context, zoneID string, rr DNSRecord) ([]DNSRecord, error) {
	// Construct a query string
	v := url.Values{}
	// Request as many records as possible per page - API max is 100
	v.Set("per_page", "100")
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	// Request as many records as possible per page - API max is 100
	v.Set("per_page", "100")
=======
	// Using default per_page value as specified by the API
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	if rr.Name != "" {
		v.Set("name", toUTS46ASCII(rr.Name))
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}

	var records []DNSRecord
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/zones/%s/dns_records?%s", zoneID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, err
		}
		var r DNSListResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}
		// Loop around and fetch the next page
		page++
	}
	return records, nil
}

// DNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) DNSRecord(ctx context.Context, zoneID, recordID string) (DNSRecord, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(ctx context.Context, zoneID, recordID string, rr DNSRecord) error {
	rr.Name = toUTS46ASCII(rr.Name)

	// Populate the record name from the existing one if the update didn't
	// specify it.
	if rr.Name == "" || rr.Type == "" {
		rec, err := api.DNSRecord(ctx, zoneID, recordID)
		if err != nil {
			return err
		}

		if rr.Name == "" {
			rr.Name = rec.Name
		}
		if rr.Type == "" {
			rr.Type = rec.Type
		}
	}
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, rr)
	if err != nil {
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

// DeleteDNSRecord deletes a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
func (api *API) DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return errors.Wrap(err, errMakeRequestError)
=======
		return err
>>>>>>> 4d7e5ad26 (update vendored files)
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"net/url"
	"strconv"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
	"net/url"
	"strconv"
=======
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"time"

	"github.com/goccy/go-json"
	"golang.org/x/net/idna"
)

// ErrMissingBINDContents is for when the BIND file contents is required but not set.
var ErrMissingBINDContents = errors.New("required BIND config contents missing")

// DNSRecord represents a DNS record in a zone.
type DNSRecord struct {
	CreatedOn  time.Time   `json:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Content    string      `json:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Comment    string      `json:"comment,omitempty"` // the server will omit the comment field when the comment is empty
	Tags       []string    `json:"tags,omitempty"`
}

// DNSRecordResponse represents the response from the DNS endpoint.
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type ListDirection string

const (
	ListDirectionAsc  ListDirection = "asc"
	ListDirectionDesc ListDirection = "desc"
)

type ListDNSRecordsParams struct {
	Type      string        `url:"type,omitempty"`
	Name      string        `url:"name,omitempty"`
	Content   string        `url:"content,omitempty"`
	Proxied   *bool         `url:"proxied,omitempty"`
	Comment   string        `url:"comment,omitempty"` // currently, the server does not support searching for records with an empty comment
	Tags      []string      `url:"tag,omitempty"`     // potentially multiple `tag=`
	TagMatch  string        `url:"tag-match,omitempty"`
	Order     string        `url:"order,omitempty"`
	Direction ListDirection `url:"direction,omitempty"`
	Match     string        `url:"match,omitempty"`
	Priority  *uint16       `url:"-"`

	ResultInfo
}

type UpdateDNSRecordParams struct {
	Type     string      `json:"type,omitempty"`
	Name     string      `json:"name,omitempty"`
	Content  string      `json:"content,omitempty"`
	Data     interface{} `json:"data,omitempty"` // data for: SRV, LOC
	ID       string      `json:"-"`
	Priority *uint16     `json:"priority,omitempty"`
	TTL      int         `json:"ttl,omitempty"`
	Proxied  *bool       `json:"proxied,omitempty"`
	Comment  *string     `json:"comment,omitempty"` // nil will keep the current comment, while StringPtr("") will empty it
	Tags     []string    `json:"tags"`
}

// DNSListResponse represents the response from the list DNS records endpoint.
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// listDNSRecordsDefaultPageSize represents the default per_page size of the API.
var listDNSRecordsDefaultPageSize int = 100

// nontransitionalLookup implements the nontransitional processing as specified in
// Unicode Technical Standard 46 with almost all checkings off to maximize user freedom.
var nontransitionalLookup = idna.New(
	idna.MapForLookup(),
	idna.StrictDomainName(false),
	idna.ValidateLabels(false),
)

// toUTS46ASCII tries to convert IDNs (international domain names)
// from Unicode form to Punycode, using non-transitional process specified
// in UTS 46.
//
// Note: conversion errors are silently discarded and partial conversion
// results are used.
func toUTS46ASCII(name string) string {
	name, _ = nontransitionalLookup.ToASCII(name)
	return name
}

// proxiedRecordsRe is the regular expression for determining if a DNS record
// is proxied or not.
var proxiedRecordsRe = regexp.MustCompile(`(?m)^.*\.\s+1\s+IN\s+CNAME.*$`)

// proxiedRecordImportTemplate is the multipart template for importing *only*
// proxied records. See `nonProxiedRecordImportTemplate` for importing records
// that are not proxied.
var proxiedRecordImportTemplate = `--------------------------BOUNDARY
Content-Disposition: form-data; name="file"; filename="bind.txt"

%s
--------------------------BOUNDARY
Content-Disposition: form-data; name="proxied"

true
--------------------------BOUNDARY--`

// nonProxiedRecordImportTemplate is the multipart template for importing DNS
// records that are not proxed. For importing proxied records, use
// `proxiedRecordImportTemplate`.
var nonProxiedRecordImportTemplate = `--------------------------BOUNDARY
Content-Disposition: form-data; name="file"; filename="bind.txt"

%s
--------------------------BOUNDARY--`

// sanitiseBINDFileInput accepts the BIND file as a string and removes parts
// that are not required for importing or would break the import (like SOA
// records).
func sanitiseBINDFileInput(s string) string {
	// Remove SOA records.
	soaRe := regexp.MustCompile(`(?m)[\r\n]+^.*IN\s+SOA.*$`)
	s = soaRe.ReplaceAllString(s, "")

	// Remove all comments.
	commentRe := regexp.MustCompile(`(?m)[\r\n]+^.*;;.*$`)
	s = commentRe.ReplaceAllString(s, "")

	// Swap all the tabs to spaces.
	r := strings.NewReplacer(
		"\t", " ",
		"\n\n", "\n",
	)
	s = r.Replace(s)
	s = strings.TrimSpace(s)

	return s
}

// extractProxiedRecords accepts a BIND file (as a string) and returns only the
// proxied DNS records.
func extractProxiedRecords(s string) string {
	proxiedOnlyRecords := proxiedRecordsRe.FindAllString(s, -1)
	return strings.Join(proxiedOnlyRecords, "\n")
}

// removeProxiedRecords accepts a BIND file (as a string) and returns the file
// contents without any proxied records included.
func removeProxiedRecords(s string) string {
	return proxiedRecordsRe.ReplaceAllString(s, "")
}

type ExportDNSRecordsParams struct{}
type ImportDNSRecordsParams struct {
	BINDContents string
}

type CreateDNSRecordParams struct {
	CreatedOn  time.Time   `json:"created_on,omitempty" url:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty" url:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty" url:"type,omitempty"`
	Name       string      `json:"name,omitempty" url:"name,omitempty"`
	Content    string      `json:"content,omitempty" url:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty" url:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Comment    string      `json:"comment,omitempty" url:"comment,omitempty"` // to the server, there's no difference between "no comment" and "empty comment"
	Tags       []string    `json:"tags,omitempty"`
}

// CreateDNSRecord creates a DNS record for the zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
func (api *API) CreateDNSRecord(ctx context.Context, rc *ResourceContainer, params CreateDNSRecordParams) (DNSRecord, error) {
	if rc.Identifier == "" {
		return DNSRecord{}, ErrMissingZoneID
	}
	params.Name = toUTS46ASCII(params.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return DNSRecord{}, err
	}

	var recordResp *DNSRecordResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp.Result, nil
}

// ListDNSRecords returns a slice of DNS records for the given zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
func (api *API) ListDNSRecords(ctx context.Context, rc *ResourceContainer, params ListDNSRecordsParams) ([]DNSRecord, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingZoneID
	}

	params.Name = toUTS46ASCII(params.Name)

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = listDNSRecordsDefaultPageSize
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []DNSRecord
	var lastResultInfo ResultInfo

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/dns_records", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, err
		}
		var listResponse DNSListResponse
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		lastResultInfo = listResponse.ResultInfo
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &lastResultInfo, nil
}

// ErrMissingDNSRecordID is for when DNS record ID is needed but not given.
var ErrMissingDNSRecordID = errors.New("required DNS record ID missing")

// GetDNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) GetDNSRecord(ctx context.Context, rc *ResourceContainer, recordID string) (DNSRecord, error) {
	if rc.Identifier == "" {
		return DNSRecord{}, ErrMissingZoneID
	}
	if recordID == "" {
		return DNSRecord{}, ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(ctx context.Context, rc *ResourceContainer, params UpdateDNSRecordParams) (DNSRecord, error) {
	if rc.Identifier == "" {
		return DNSRecord{}, ErrMissingZoneID
	}

	if params.ID == "" {
		return DNSRecord{}, ErrMissingDNSRecordID
	}

	params.Name = toUTS46ASCII(params.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return DNSRecord{}, err
	}

	var recordResp *DNSRecordResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp.Result, nil
}

// DeleteDNSRecord deletes a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
func (api *API) DeleteDNSRecord(ctx context.Context, rc *ResourceContainer, recordID string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}
	if recordID == "" {
		return ErrMissingDNSRecordID
	}

	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
<<<<<<< HEAD
		return errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return errors.Wrap(err, errUnmarshalError)
=======
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	return nil
}

// ExportDNSRecords returns all DNS records for a zone in the BIND format.
//
// API reference: https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-export-dns-records
func (api *API) ExportDNSRecords(ctx context.Context, rc *ResourceContainer, params ExportDNSRecordsParams) (string, error) {
	if rc.Level != ZoneRouteLevel {
		return "", ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return "", ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/dns_records/export", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// ImportDNSRecords takes the contents of a BIND configuration file and imports
// all records at once.
//
// The current state of the API doesn't allow the proxying field to be
// automatically set on records where the TTL is 1. Instead you need to
// explicitly tell the endpoint which records are proxied in the form data. To
// achieve a simpler abstraction, we do the legwork in the method of making the
// two separate API calls (one for proxied and one for non-proxied) instead of
// making the end user know about this detail.
//
// API reference: https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-import-dns-records
func (api *API) ImportDNSRecords(ctx context.Context, rc *ResourceContainer, params ImportDNSRecordsParams) error {
	if rc.Level != ZoneRouteLevel {
		return ErrRequiredZoneLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	if params.BINDContents == "" {
		return ErrMissingBINDContents
	}

	sanitisedBindData := sanitiseBINDFileInput(params.BINDContents)
	nonProxiedRecords := removeProxiedRecords(sanitisedBindData)
	proxiedOnlyRecords := extractProxiedRecords(sanitisedBindData)

	nonProxiedRecordPayload := []byte(fmt.Sprintf(nonProxiedRecordImportTemplate, nonProxiedRecords))
	nonProxiedReqBody := bytes.NewReader(nonProxiedRecordPayload)

	uri := fmt.Sprintf("/zones/%s/dns_records/import", rc.Identifier)
	multipartUploadHeaders := http.Header{
		"Content-Type": {"multipart/form-data; boundary=------------------------BOUNDARY"},
	}

	_, err := api.makeRequestContextWithHeaders(ctx, http.MethodPost, uri, nonProxiedReqBody, multipartUploadHeaders)
	if err != nil {
		return err
	}

	proxiedRecordPayload := []byte(fmt.Sprintf(proxiedRecordImportTemplate, proxiedOnlyRecords))
	proxiedReqBody := bytes.NewReader(proxiedRecordPayload)

	_, err = api.makeRequestContextWithHeaders(ctx, http.MethodPost, uri, proxiedReqBody, multipartUploadHeaders)
	if err != nil {
		return err
	}

	return nil
}
