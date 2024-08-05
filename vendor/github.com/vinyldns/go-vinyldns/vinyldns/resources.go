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

import (
	"strconv"
	"strings"
)

// Error represents an error from the
// vinyldns API
type Error struct {
	RequestURL    string
	RequestMethod string
	RequestBody   string
	ResponseBody  string
	ResponseCode  int
}

func (d Error) Error() string {
	components := []string{
		"Request URL:",
		d.RequestURL,
		"Request Method:",
		d.RequestMethod,
		"Request body:",
		d.RequestBody,
		"Response code: ",
		strconv.Itoa(d.ResponseCode),
		"Response body:",
		d.ResponseBody}
	return strings.Join(components, "\n")
}

// ListFilter represents the list query parameters that may be passed to
// VinylDNS API endpoints such as /zones and /zones/${zone_id}/recordsets
type ListFilter struct {
	NameFilter string
	StartFrom  string
	MaxItems   int
}

// NameSort specifies the name sort order for record sets returned by the global list record set response.
// Valid values are ASC (ascending; default) and DESC (descending).
type NameSort string

const (
	// ASC represents an ascending NameSort
	ASC NameSort = "ASC"

	// DESC represents a descending NameSort
	DESC NameSort = "DESC"
)

// GlobalListFilter represents the list query parameters that may be passed to
// VinylDNS API endpoints such as /recordsets
type GlobalListFilter struct {
	RecordNameFilter       string
	RecordTypeFilter       string
	RecordOwnerGroupFilter string
	NameSort               NameSort
	StartFrom              string
	MaxItems               int
}
