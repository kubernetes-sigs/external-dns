/*
Copyright 2017 The Kubernetes Authors.

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

package internal

import "fmt"

// ListZones dto to read list of zones from API
type ListZones struct {
	Zones []Zone `json:"zones"`
}

// Zone dto to read info from API
type Zone struct {
	Name    string       `json:"name"`
	Records []ZoneRecord `json:"records"`
}

// RRSet dto as part of zone info from API
type RRSet struct {
	TTL     int               `json:"ttl"`
	Records []ResourceRecords `json:"resource_records"`
}

// ResourceRecords dto describe records in RRSet
type ResourceRecords struct {
	Content []string `json:"content"`
}

// ZoneRecord dto describe records in Zone
type ZoneRecord struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	TTL          uint     `json:"ttl"`
	ShortAnswers []string `json:"short_answers"`
}

// APIError customization for API calls
type APIError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error,omitempty"`
}

// Error implementation
func (a APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.StatusCode, a.Message)
}
