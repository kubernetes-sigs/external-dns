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

// BatchRecordChanges represents a list of record changes,
// as returned by the list batch changes VinylDNS API endpoint.
type BatchRecordChanges struct {
	BatchChanges []RecordChange `json:"batchChanges,omitempty"`
}

// RecordChange represents an individual batch record change.
type RecordChange struct {
	ID               string     `json:"id,omitempty"`
	Status           string     `json:"status,omitempty"`
	ChangeType       string     `json:"changeType,omitempty"`
	RecordName       string     `json:"recordName,omitempty"`
	TTL              int        `json:"ttl,omitempty"`
	Type             string     `json:"type,omitempty"`
	ZoneName         string     `json:"zoneName,omitempty"`
	InputName        string     `json:"inputName,omitempty"`
	ZoneID           string     `json:"zoneId,omitempty"`
	TotalChanges     int        `json:"totalChanges,omitempty"`
	UserName         string     `json:"userName,omitempty"`
	Comments         string     `json:"comments,omitempty"`
	UserID           string     `json:"userId,omitempty"`
	CreatedTimestamp string     `json:"createdTimestamp,omitempty"`
	Record           RecordData `json:"record,omitempty"`
	OwnerGroupID     string     `json:"ownerGroupId,omitempty"`
}

// BatchRecordChangeUpdateResponse is represents a batch record change create or update response
type BatchRecordChangeUpdateResponse struct {
	ID                 string         `json:"id,omitempty"`
	UserName           string         `json:"userName,omitempty"`
	UserID             string         `json:"userId,omitempty"`
	Status             string         `json:"status,omitempty"`
	Comments           string         `json:"comments,omitempty"`
	CreatedTimestamp   string         `json:"createdTimestamp,omitempty"`
	OwnerGroupID       string         `json:"ownerGroupId,omitempty"`
	Changes            []RecordChange `json:"changes,omitempty"`
	ApprovalStatus     string         `json:"approvalStatus,omitempty"`
	ReviewerID         string         `json:"reviewerId,omitempty"`
	ReviewerUserName   string         `json:"reviewerUserName,omitempty"`
	ReviewComment      string         `json:"reviewComment,omitempty"`
	ReviewTimestamp    string         `json:"reviewTimestamp,omitempty"`
	ScheduledTime      string         `json:"scheduledTime,omitempty"`
	CancelledTimestamp string         `json:"cancelledTimestamp,omitempty"`
}

// RecordData is represents a batch record change record data.
type RecordData struct {
	Address  string `json:"address,omitempty"`
	CName    string `json:"cname,omitempty"`
	PTRDName string `json:"ptrdname,omitempty"`
}

// BatchRecordChange represents a batch record change API response.
type BatchRecordChange struct {
	ID                 string         `json:"id,omitempty"`
	UserName           string         `json:"userName,omitempty"`
	UserID             string         `json:"userId,omitempty"`
	Status             string         `json:"status,omitempty"`
	Comments           string         `json:"comments,omitempty"`
	CreatedTimestamp   string         `json:"createdTimestamp,omitempty"`
	OwnerGroupID       string         `json:"ownerGroupId,omitempty"`
	Changes            []RecordChange `json:"changes,omitempty"`
	ApprovalStatus     string         `json:"approvalStatus,omitempty"`
	ReviewerID         string         `json:"reviewerId,omitempty"`
	ReviewerUserName   string         `json:"reviewerUserName,omitempty"`
	ReviewComment      string         `json:"reviewComment,omitempty"`
	ReviewTimestamp    string         `json:"reviewTimestamp,omitempty"`
	ScheduledTime      string         `json:"scheduledTime,omitempty"`
	CancelledTimestamp string         `json:"cancelledTimestamp,omitempty"`
}
