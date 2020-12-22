/*
Copyright 2020 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package infomaniak

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	infomaniakBaseURL = "https://api.infomaniak.com"
)

// InfomaniakAPI is a basic implementation of an API client to api.infomaniak.com
// It implements only the methods required for the ACME Challenge
type InfomaniakAPI struct {
	apiToken string
}

// ErrorResponse defines the error response format, as described here https://api.infomaniak.com/doc#home
type ErrorResponse struct {
	Code        string            `json:"code"`
	Description string            `json:"description,omitempty"`
	Context     map[string]string `json:"context,omitempty"`
	Errors      []ErrorResponse   `json:"errors,omitempty"`
}

// InfomaniakAPIResponse defines the generic response format, as described here https://api.infomaniak.com/doc#home
type InfomaniakAPIResponse struct {
	Result      string           `json:"result"`
	Data        *json.RawMessage `json:"data,omitempty"`
	ErrResponse ErrorResponse    `json:"error,omitempty"`
}

// NewInfomaniakAPI creates a new infomaniak API client
func NewInfomaniakAPI(apiToken string) *InfomaniakAPI {
	return &InfomaniakAPI{
		apiToken: apiToken,
	}
}

// request builds the raw request
func (ik *InfomaniakAPI) request(method, path string, body io.Reader) (*InfomaniakAPIResponse, error) {
	if path[0] != '/' {
		path = "/" + path
	}
	url := infomaniakBaseURL + path

	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+ik.apiToken)
	req.Header.Set("Content-Type", "application/json")

	log.Infof("%s %s", method, url)
	rawResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var resp InfomaniakAPIResponse
	if err := json.NewDecoder(rawResp.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("%s %s response parsing error: %v", method, path, err)
	}

	rawJSON, _ := json.Marshal(resp)
	log.Infof("Response status: `%s` json response: `%v`", rawResp.Status, string(rawJSON))

	if resp.Result != "success" {
		return nil, fmt.Errorf("%s %s failed: %v", method, path, resp.ErrResponse)
	}

	return &resp, nil
}

// get is a helper to build a bare GET request
func (ik *InfomaniakAPI) get(path string, params url.Values) (*InfomaniakAPIResponse, error) {
	base, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	if params != nil {
		base.RawQuery = params.Encode()
	}
	return ik.request("GET", base.String(), nil)
}

// get is a helper to build a bare POST request
func (ik *InfomaniakAPI) post(path string, body io.Reader) (*InfomaniakAPIResponse, error) {
	return ik.request("POST", path, body)
}

// get is a helper to build a bare PUT request
func (ik *InfomaniakAPI) put(path string, body io.Reader) (*InfomaniakAPIResponse, error) {
	return ik.request("PUT", path, body)
}

// get is a helper to build a bare DELETE request
func (ik *InfomaniakAPI) delete(path string) (*InfomaniakAPIResponse, error) {
	return ik.request("DELETE", path, nil)
}

// InfomaniakDNSDomain defines the format of a Domain object
type InfomaniakDNSDomain struct {
	ID                  uint64           `json:"id,omitempty"`
	AccountID           uint64           `json:"account_id,omitempty"`
	ServiceID           uint64           `json:"service_id,omitempty"`
	ServiceName         string           `json:"service_name,omitempty"`
	CustomerName        string           `json:"customer_name,omitempty"`
	InternalName        string           `json:"internal_name,omitempty,omitempty"`
	CreatedAt           uint64           `json:"created_at,omitempty"`
	ExpiredAt           uint64           `json:"expired_at,omitempty"`
	Version             uint64           `json:"version,omitempty"`
	Maintenance         bool             `json:"maintenance,omitempty"`
	Locked              bool             `json:"locked,omitempty"`
	OperationInProgress bool             `json:"operation_in_progress,omitempty"`
	Tags                *json.RawMessage `json:"tags,omitempty"`
	UniqueID            uint64           `json:"unique_id,omitempty"`
	Description         string           `json:"description,omitempty"`
	Isfree              bool             `json:"is_free,omitempty"`
	Rights              *json.RawMessage `json:"rights,omitempty"`
	Special             bool             `json:"special,omitempty"`
}

// InfomaniakDNSRecord defines the format of a DNSRecord object
type InfomaniakDNSRecord struct {
	ID         string `json:"id,omitempty"`
	Source     string `json:"source,omitempty"`
	SourceIdn  string `json:"source_idn,omitempty"`
	Type       string `json:"type,omitempty"`
	TTL        uint64 `json:"ttl,omitempty"`
	TTLIdn     string `json:"ttl_idn,omitempty"`
	Target     string `json:"target,omitempty"`
	TargetIdn  string `json:"target_idn,omitempty"`
	UpdatedAt  uint64 `json:"updated_at,omitempty"`
	DyndnsID   string `json:"dyndns_id,omitempty,omitempty"`
	Priority   uint64 `json:"priority,omitempty"`
	IsEditable bool   `json:"is_editable,omitempty"`
}

// ErrDomainNotFound
var ErrDomainNotFound = errors.New("domain not found")

// GetDomainByName gather a Domain object from its name
func (ik *InfomaniakAPI) GetDomainByName(name string) (*InfomaniakDNSDomain, error) {
	log.Infof("Getting domain matching `%s`", name)

	// remove trailing . if present
	if strings.HasSuffix(name, ".") {
		name = name[:len(name)-1]
	}

	// Try to find the most specific domain
	// starts with the FQDN, then remove each left label until we have a match
	for {
		i := strings.Index(name, ".")
		if i == -1 {
			break
		}
		params := url.Values{}
		params.Add("service_name", "domain")
		params.Add("customer_name", name)

		resp, err := ik.get("/1/product", params)
		if err != nil {
			return nil, err
		}

		var domains []InfomaniakDNSDomain

		if err = json.Unmarshal(*resp.Data, &domains); err != nil {
			return nil, fmt.Errorf("expected array of Domain, got: %v", string(*resp.Data))
		}

		for _, domain := range domains {
			if domain.CustomerName == name {
				log.Infof("Domain `%s` found, id=`%d`", name, domain.ID)
				return &domain, nil
			}
		}
		log.Infof("Domain `%s` not found, trying with `%s`", name, name[i+1:])
		name = name[i+1:]
	}

	return nil, ErrDomainNotFound
}

// ListDomains gather a list of Domain objects
func (ik *InfomaniakAPI) ListDomains() (*[]InfomaniakDNSDomain, error) {
	log.Infof("Getting list of domains")

	params := url.Values{}
	params.Add("service_name", "domain")

	resp, err := ik.get("/1/product", params)
	if err != nil {
		return nil, err
	}

	var domains []InfomaniakDNSDomain

	if err = json.Unmarshal(*resp.Data, &domains); err != nil {
		return nil, fmt.Errorf("expected array of Domain, got: %v", string(*resp.Data))
	}

	return &domains, nil
}

// getRecordID gather a record id from its specs (domain, source, target, rtype)
func (ik *InfomaniakAPI) getRecordID(domain *InfomaniakDNSDomain, source, target, rtype string) (*string, error) {
	log.Infof("Getting all record for domain=%d, then match source=%s target=%s rtype=%s", domain.ID, source, target, rtype)

	resp, err := ik.get(fmt.Sprintf("/1/domain/%d/dns/record", domain.ID), nil)
	if err != nil {
		return nil, err
	}

	var records []InfomaniakDNSRecord

	if err = json.Unmarshal(*resp.Data, &records); err != nil {
		return nil, fmt.Errorf("expected array of Record, got: %v", string(*resp.Data))
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("no records in zone")
	}

	for _, record := range records {
		if record.Source == source && record.Target == target && record.Type == rtype {
			return &record.ID, nil
		}
	}

	return nil, nil
}

func (ik *InfomaniakAPI) GetRecords(domain *InfomaniakDNSDomain) (*[]InfomaniakDNSRecord, error) {
	resp, err := ik.get(fmt.Sprintf("/1/domain/%d/dns/record", domain.ID), nil)
	if err != nil {
		return nil, err
	}

	var records []InfomaniakDNSRecord

	if err = json.Unmarshal(*resp.Data, &records); err != nil {
		return nil, fmt.Errorf("expected array of Record, got: %v", string(*resp.Data))
	}

	// if len(records) < 1 {
	// 	return nil, fmt.Errorf("no records in zone")
	// }

	return &records, nil
}

// EnsureDNSRecord ensures a record is present with the correct key
func (ik *InfomaniakAPI) EnsureDNSRecord(domain *InfomaniakDNSDomain, source, target, rtype string, ttl uint64) error {
	log.Infof("Ensure record domain=%d source=%s target=%s rtype=%s TTL=%d", domain.ID, source, target, rtype, ttl)

	recordID, err := ik.getRecordID(domain, source, target, rtype)
	if err != nil {
		return err
	}

	if recordID != nil {
		log.Infof("Record already exists (domain=%d record=%s source=%s rtype=%s target=%s), skipping addition", domain.ID, *recordID, source, rtype, target)
		return nil
	}

	// TTL has to be greater or equal to 300
	if ttl < 300 {
		ttl = 300
	}

	record := InfomaniakDNSRecord{Source: source, Target: target, Type: rtype, TTL: ttl}
	rawJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	log.Infof("Adding record domain=%d (source=%s rtype=%s target=%s ttl=%d)", domain.ID, source, rtype, target, ttl)
	_, err = ik.post(fmt.Sprintf("/1/domain/%d/dns/record", domain.ID), bytes.NewBuffer(rawJSON))
	return err
}

// ModifyDNSRecord updates a DNSRecord with a new target
func (ik *InfomaniakAPI) ModifyDNSRecord(domain *InfomaniakDNSDomain, source, oldTarget, newTarget, rtype string, ttl uint64) error {
	log.Infof("Search for record matching domain=%d source=%s target=%s rtype=%s", domain.ID, source, oldTarget, rtype)

	recordID, err := ik.getRecordID(domain, source, oldTarget, rtype)
	if err != nil {
		return err
	}

	if recordID == nil {
		return fmt.Errorf("record does not exist")
	}

	// TTL has to be greater or equal to 300
	if ttl < 300 {
		ttl = 300
	}

	record := InfomaniakDNSRecord{Source: source, Target: newTarget, Type: rtype, TTL: ttl}
	rawJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	log.Infof("Updating record domain=%d (source=%s rtype=%s target=%s ttl=%d)", domain.ID, source, rtype, newTarget, ttl)
	_, err = ik.put(fmt.Sprintf("/1/domain/%d/dns/record/%s", domain.ID, *recordID), bytes.NewBuffer(rawJSON))
	return err

}

// RemoveDNSRecord ensures a record is absent
func (ik *InfomaniakAPI) RemoveDNSRecord(domain *InfomaniakDNSDomain, source, target, rtype string) error {
	log.Infof("Remove record domain=%d source=%s rtype=%s target=%s", domain.ID, source, rtype, target)
	recordID, err := ik.getRecordID(domain, source, target, rtype)
	if err != nil {
		return err
	}

	// the record is already absent doing nothing
	if recordID == nil || len(*recordID) < 1 {
		log.Infof("No record found (domain=%d source=%s rtype=%s target=%s), skipping deletion", domain.ID, source, rtype, target)
		return nil
	}

	log.Infof("Deleting record domain=%d record=%s (source=%s rtype=%s target=%s)", domain.ID, *recordID, source, rtype, target)
	_, err = ik.delete(fmt.Sprintf("/1/domain/%d/dns/record/%s", domain.ID, *recordID))
	return err
}
