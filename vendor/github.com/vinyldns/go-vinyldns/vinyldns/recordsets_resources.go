/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

// RecordSetChange represents a record
// set change.
type RecordSetChange struct {
	Zone       Zone      `json:"zone"`
	RecordSet  RecordSet `json:"recordSet"`
	UserID     string    `json:"userId"`
	ChangeType string    `json:"changeType"`
	Status     string    `json:"status"`
	Created    string    `json:"created"`
	ID         string    `json:"id"`
}

// RecordSetChanges represents a recordset changes response
type RecordSetChanges struct {
	RecordSetChanges []RecordSetChange `json:"recordSetChanges"`
	ZoneID           string            `json:"zoneId,omitempty"`
	StartFrom        string            `json:"startFrom,omitempty"`
	NextID           string            `json:"nextId,omitempty"`
	MaxItems         int               `json:"maxItems,omitempty"`
	Status           string            `json:"status,omitempty"`
}

// RecordSet represents a DNS record set.
type RecordSet struct {
	ID           string   `json:"id,omitempty"`
	ZoneID       string   `json:"zoneId"`
	OwnerGroupID string   `json:"ownerGroupId,omitempty"`
	Name         string   `json:"name,omitempty"`
	Type         string   `json:"type"`
	Status       string   `json:"status,omitempty"`
	Created      string   `json:"created,omitempty"`
	Updated      string   `json:"updated,omitempty"`
	TTL          int      `json:"ttl"`
	Account      string   `json:"account"`
	Records      []Record `json:"records"`
}

// RecordSetUpdateResponse represents
// a JSON response from the record set update endpoint.
type RecordSetUpdateResponse struct {
	Zone      Zone      `json:"zone"`
	RecordSet RecordSet `json:"recordSet"`
	ChangeID  string    `json:"id"`
	Status    string    `json:"status"`
}

// Record represents a DNS record
type Record struct {
	Address     string `json:"address,omitempty"`
	CName       string `json:"cname,omitempty"`
	Preference  int    `json:"preference,omitempty"`
	Exchange    string `json:"exchange,omitempty"`
	NSDName     string `json:"nsdname,omitempty"`
	PTRDName    string `json:"ptrdname,omitempty"`
	MName       string `json:"mname,omitempty"`
	RName       string `json:"rname,omitempty"`
	Serial      int    `json:"serial,omitempty"`
	Refresh     int    `json:"refresh,omitempty"`
	Retry       int    `json:"retry,omitempty"`
	Expire      int    `json:"expire,omitempty"`
	Minimum     int    `json:"minimum,omitempty"`
	Text        string `json:"text,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	Port        int    `json:"port,omitempty"`
	Target      string `json:"target,omitempty"`
	Algorithm   int    `json:"algorithm,omitempty"`
	Type        string `json:"type,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

// RecordSetResponse represents the JSON
// response from the record set endpoint.
type RecordSetResponse struct {
	RecordSet RecordSet `json:"recordSet"`
}

// RecordSetsResponse represents the JSON
// response from the record sets endpoint.
type RecordSetsResponse struct {
	NextID           string      `json:"nextId,omitempty"`
	MaxItems         int         `json:"maxItems,omitempty"`
	StartFrom        string      `json:"startFrom,omitempty"`
	RecordNameFilter string      `json:"recordNameFilter,omitempty"`
	RecordSets       []RecordSet `json:"recordSets"`
}
