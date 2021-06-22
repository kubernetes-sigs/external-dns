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

import "encoding/json"

// BatchRecordChanges returns the list of batch record changes
func (c *Client) BatchRecordChanges() ([]RecordChange, error) {
	changes := &BatchRecordChanges{}
	err := resourceRequest(c, batchRecordChangesEP(c), "GET", nil, changes)
	if err != nil {
		return nil, err
	}

	return changes.BatchChanges, nil
}

// BatchRecordChange returns the batch record change
// associated with the change whose ID it's passed.
func (c *Client) BatchRecordChange(changeID string) (*BatchRecordChange, error) {
	change := &BatchRecordChange{}
	err := resourceRequest(c, batchRecordChangeEP(c, changeID), "GET", nil, change)
	if err != nil {
		return nil, err
	}

	return change, nil
}

// BatchRecordChangeCreate creates the batch record change it's passed.
func (c *Client) BatchRecordChangeCreate(change *BatchRecordChange) (*BatchRecordChangeUpdateResponse, error) {
	cJSON, err := json.Marshal(change)
	if err != nil {
		return nil, err
	}
	var resource = &BatchRecordChangeUpdateResponse{}
	err = resourceRequest(c, batchRecordChangesEP(c), "POST", cJSON, resource)
	if err != nil {
		return &BatchRecordChangeUpdateResponse{}, err
	}

	return resource, nil
}
