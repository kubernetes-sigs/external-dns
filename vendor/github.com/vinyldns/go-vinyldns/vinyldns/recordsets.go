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
	"io"
)

// RecordSetLimit is the highest number of records the vinyldns server will allow at once
// TODO: is there a way to get this limit directly from vinyldns?
const RecordSetLimit = 100

// RecordSetCollector creates a function to retrieve the next set of recordsets.
// To retrieve *all* recordsets, call that function repeatedly until err == // io.EOF
func (c *Client) RecordSetCollector(zoneID string, limit int) (func() ([]RecordSet, error), error) {
	if limit > RecordSetLimit {
		return nil, fmt.Errorf("Limit must be zero or not greater than %d", RecordSetLimit)
	}

	var nextID string
	var recordSets []RecordSet
	var err error

	return func() ([]RecordSet, error) {
		if err != nil {
			return nil, err
		}

		for {
			rss := &RecordSetsResponse{}
			err = resourceRequest(c, recordSetsListEP(c, zoneID, ListFilter{
				StartFrom: nextID,
				MaxItems:  limit,
			}), "GET", nil, rss)
			if err != nil {
				return nil, err
			}
			recordSets = append(recordSets, rss.RecordSets...)

			nextID = rss.NextID
			if len(nextID) == 0 {
				// keep from trying to get more records
				err = io.EOF

				break
			}
		}

		// return at most `limit` records and remove those returned from recordSets
		max := limit
		if max == 0 || max > len(recordSets) {
			max = len(recordSets)
		}
		r := recordSets[:max]
		recordSets = recordSets[max:]
		return r, err
	}, nil
}

// RecordSets retrieves a list of RecordSets from a Zone.
func (c *Client) RecordSets(id string) ([]RecordSet, error) {
	collector, err := c.RecordSetCollector(id, 0)
	if err != nil {
		return nil, err
	}

	var recordSets []RecordSet
	for {
		rs, err := collector()
		if err != nil && err != io.EOF {
			return nil, err
		}
		recordSets = append(recordSets, rs...)
		if err == io.EOF {
			break
		}
	}

	return recordSets, nil
}

// RecordSetsListAll retrieves the complete list of record sets with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) RecordSetsListAll(zoneID string, filter ListFilter) ([]RecordSet, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	rss := []RecordSet{}

	for {
		resp, err := c.recordSetsList(zoneID, filter)
		if err != nil {
			return nil, err
		}

		rss = append(rss, resp.RecordSets...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return rss, nil
		}
	}
}

// RecordSet retrieves the record matching the Zone ID and RecordSet ID it's passed.
func (c *Client) RecordSet(zoneID, recordSetID string) (RecordSet, error) {
	rs := &RecordSetResponse{}
	err := resourceRequest(c, recordSetEP(c, zoneID, recordSetID), "GET", nil, rs)
	if err != nil {
		return RecordSet{}, err
	}

	return rs.RecordSet, nil
}

// RecordSetCreate creates the RecordSet it's passed in the Zone whose ID it's passed.
func (c *Client) RecordSetCreate(rs *RecordSet) (*RecordSetUpdateResponse, error) {
	rsJSON, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	var resource = &RecordSetUpdateResponse{}
	err = resourceRequest(c, recordSetsEP(c, rs.ZoneID), "POST", rsJSON, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetUpdate updates the RecordSet matching the Zone ID and RecordSetID it's passed.
func (c *Client) RecordSetUpdate(rs *RecordSet) (*RecordSetUpdateResponse, error) {
	rsJSON, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	var resource = &RecordSetUpdateResponse{}
	err = resourceRequest(c, recordSetEP(c, rs.ZoneID, rs.ID), "PUT", rsJSON, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetDelete deletes the RecordSet matching the Zone ID and RecordSet ID it's passed.
func (c *Client) RecordSetDelete(zoneID, recordSetID string) (*RecordSetUpdateResponse, error) {
	resource := &RecordSetUpdateResponse{}
	err := resourceRequest(c, recordSetEP(c, zoneID, recordSetID), "DELETE", nil, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetChanges retrieves the RecordSetChanges response for the Zone and ListFilter it's passed.
func (c *Client) RecordSetChanges(zoneID string, f ListFilter) (*RecordSetChanges, error) {
	rsc := &RecordSetChanges{}
	err := resourceRequest(c, recordSetChangesEP(c, zoneID, f), "GET", nil, rsc)
	if err != nil {
		return &RecordSetChanges{}, err
	}

	return rsc, nil
}

// RecordSetChangesListAll retrieves the complete list of record set changes for the Zone ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) RecordSetChangesListAll(zoneID string, filter ListFilter) ([]RecordSetChange, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	rsc := []RecordSetChange{}

	for {
		resp, err := c.RecordSetChanges(zoneID, filter)
		if err != nil {
			return nil, err
		}

		rsc = append(rsc, resp.RecordSetChanges...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return rsc, nil
		}
	}
}

// RecordSetChange retrieves the RecordSetChange matching the Zone, RecordSet, and Change IDs
// it's passed.
func (c *Client) RecordSetChange(zoneID, recordSetID, changeID string) (*RecordSetChange, error) {
	rsc := &RecordSetChange{}
	err := resourceRequest(c, recordSetChangeEP(c, zoneID, recordSetID, changeID), "GET", nil, rsc)
	if err != nil {
		return &RecordSetChange{}, err
	}

	return rsc, nil
}
