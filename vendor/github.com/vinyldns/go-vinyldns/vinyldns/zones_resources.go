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

// ZoneConnection represents a zone connection
type ZoneConnection struct {
	Name          string `json:"name,omitempty"`
	KeyName       string `json:"keyName,omitempty"`
	Key           string `json:"key,omitempty"`
	PrimaryServer string `json:"primaryServer,omitempty"`
}

// ACLRule represents an ACL rule
type ACLRule struct {
	AccessLevel string   `json:"accessLevel"`
	Description string   `json:"description,omitempty"`
	UserID      string   `json:"userId,omitempty"`
	GroupID     string   `json:"groupId,omitempty"`
	RecordMask  string   `json:"recordMask,omitempty"`
	RecordTypes []string `json:"recordTypes"`
}

// ZoneACL represents a zone ACL
type ZoneACL struct {
	Rules []ACLRule `json:"rules"`
}

// Zone represents a zone
type Zone struct {
	Name               string          `json:"name,omitempty"`
	Email              string          `json:"email,omitempty"`
	Status             string          `json:"status,omitempty"`
	Created            string          `json:"created,omitempty"`
	ID                 string          `json:"id,omitempty"`
	AdminGroupID       string          `json:"adminGroupId,omitempty"`
	LatestSync         string          `json:"latestSync,omitempty"`
	Updated            string          `json:"updated,omitempty"`
	Account            string          `json:"account,omitempty"`
	BackendID          string          `json:"backendId,omitempty"`
	AccessLevel        string          `json:"accessLevel,omitempty"`
	Connection         *ZoneConnection `json:"connection,omitempty"`
	TransferConnection *ZoneConnection `json:"transferConnection,omitempty"`
	ACL                *ZoneACL        `json:"acl,omitempty"`
	Shared             bool            `json:"shared,omitempty"`
	IsTest             bool            `json:"isTest,omitempty"`
}

// ZoneResponse represents the JSON response
// from the zone endpoint
type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

// ZoneUpdateResponse represents the JSON
// response from the zone update endpoint
type ZoneUpdateResponse struct {
	Zone       Zone   `json:"zone"`
	UserID     string `json:"userId"`
	ChangeType string `json:"changeType"`
	Status     string `json:"status"`
	Created    string `json:"created"`
	ID         string `json:"id"`
}

// Zones is a slice of zones
type Zones struct {
	Zones     []Zone `json:"zones"`
	StartFrom string `json:"startFrom"`
	MaxItems  int    `json:"maxItems"`
	NextID    string `json:"nextId"`
}

// ZoneChanges represents the zone changes
type ZoneChanges struct {
	ZoneID      string       `json:"zoneId"`
	ZoneChanges []ZoneChange `json:"zoneChanges"`
	StartFrom   string       `json:"startFrom"`
	MaxItems    int          `json:"maxItems"`
	NextID      string       `json:"nextId"`
}

// ZoneChange represents a zone change
type ZoneChange struct {
	Zone       Zone   `json:"zone"`
	UserID     string `json:"userId"`
	ChangeType string `json:"changeType"`
	Status     string `json:"status"`
	Created    string `json:"created"`
	ID         string `json:"id"`
}
