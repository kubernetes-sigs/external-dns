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

// User represents a vinyldns user.
type User struct {
	ID        string `json:"id,omitempty"`
	UserName  string `json:"userName,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Created   string `json:"created,omitempty"`
}

// Groups is a slice of groups
type Groups struct {
	Groups          []Group `json:"groups"`
	GroupNameFilter string  `json:"groupNameFilter,omitempty"`
	MaxItems        int     `json:"maxItems,omitempty"`
	NextID          string  `json:"nextId,omitempty"`
	StartFrom       string  `json:"startFrom,omitempty"`
}

// Group represents a vinyldns group.
type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Created     string `json:"created,omitempty"`
	Members     []User `json:"members"`
	Admins      []User `json:"admins"`
}

// GroupAdmins is a slice of Users
type GroupAdmins struct {
	GroupAdmins []User `json:"admins"`
}

// GroupMembers is a slice of Users
type GroupMembers struct {
	GroupMembers []User `json:"members"`
}

// GroupChange represents a group change event object.
type GroupChange struct {
	UserID     string `json:"userId,omitempty"`
	Created    string `json:"created,omitempty"`
	ChangeType string `json:"changeType,omitempty"`
	NewGroup   Group  `json:"newGroup,omitempty"`
	OldGroup   Group  `json:"oldGroup,omitempty"`
}

// GroupChanges is represents the group changes.
type GroupChanges struct {
	Changes []GroupChange `json:"changes"`
}
