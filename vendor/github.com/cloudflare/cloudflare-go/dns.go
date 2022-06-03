package cloudflare

import (
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
	}
	return nil
}
