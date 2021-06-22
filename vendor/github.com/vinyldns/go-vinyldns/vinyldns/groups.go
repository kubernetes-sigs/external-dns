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
	"encoding/json"
	"fmt"
)

// Groups retrieves a list of Groups that the requester is a part of.
func (c *Client) Groups() ([]Group, error) {
	groups := &Groups{}
	err := resourceRequest(c, groupsEP(c), "GET", nil, groups)
	if err != nil {
		return []Group{}, err
	}

	return groups.Groups, nil
}

// GroupsListAll retrieves the complete list of groups with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) GroupsListAll(filter ListFilter) ([]Group, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	groups := []Group{}

	for {
		resp, err := c.groupsList(filter)
		if err != nil {
			return nil, err
		}

		groups = append(groups, resp.Groups...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return groups, nil
		}
	}
}

// GroupCreate creates the Group it's passed.
func (c *Client) GroupCreate(g *Group) (*Group, error) {
	gJSON, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	var resource = &Group{}
	err = resourceRequest(c, groupsEP(c), "POST", gJSON, resource)
	if err != nil {
		return &Group{}, err
	}

	return resource, nil
}

// Group gets the Group whose ID it's passed.
func (c *Client) Group(groupID string) (*Group, error) {
	group := &Group{}
	err := resourceRequest(c, groupEP(c, groupID), "GET", nil, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// GroupDelete deletes the Group whose ID it's passed.
func (c *Client) GroupDelete(groupID string) (*Group, error) {
	group := &Group{}
	err := resourceRequest(c, groupEP(c, groupID), "DELETE", nil, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// GroupUpdate updates the Group whose ID it's passed.
func (c *Client) GroupUpdate(groupID string, g *Group) (*Group, error) {
	gJSON, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	var resource = &Group{}
	err = resourceRequest(c, groupEP(c, groupID), "PUT", gJSON, resource)
	if err != nil {
		return &Group{}, err
	}

	return resource, nil
}

// GroupAdmins returns an array of Users that are admins
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupAdmins(groupID string) ([]User, error) {
	admins := &GroupAdmins{}
	err := resourceRequest(c, groupAdminsEP(c, groupID), "GET", nil, admins)
	if err != nil {
		return nil, err
	}

	return admins.GroupAdmins, nil
}

// GroupMembers returns an array of Users that are members
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupMembers(groupID string) ([]User, error) {
	members := &GroupMembers{}
	err := resourceRequest(c, groupMembersEP(c, groupID), "GET", nil, members)
	if err != nil {
		return nil, err
	}

	return members.GroupMembers, nil
}

// GroupActivity returns group change activity
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupActivity(groupID string) (*GroupChanges, error) {
	activity := &GroupChanges{}
	err := resourceRequest(c, groupActivityEP(c, groupID), "GET", nil, activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}
