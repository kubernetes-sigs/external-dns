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

// zonesList retrieves the list of zones with the List criteria passed.
func (c *Client) zonesList(filter ListFilter) (*Zones, error) {
	zones := &Zones{}
	err := resourceRequest(c, zonesListEP(c, filter), "GET", nil, zones)
	if err != nil {
		return zones, err
	}

	return zones, nil
}

// zoneChangesList retrieves the list of zone changes with the List criteria passed.
func (c *Client) zoneChangesList(zoneID string, filter ListFilter) (*ZoneChanges, error) {
	changes := &ZoneChanges{}
	err := resourceRequest(c, zoneChangesEP(c, zoneID, filter), "GET", nil, changes)
	if err != nil {
		return changes, err
	}

	return changes, nil
}
