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
	"net/http"
)

// Zones retrieves the list of zones a user has access to.
func (c *Client) Zones() ([]Zone, error) {
	zones := &Zones{}
	err := resourceRequest(c, zonesEP(c), "GET", nil, zones)
	if err != nil {
		return []Zone{}, err
	}

	return zones.Zones, nil
}

// ZonesListAll retrieves the complete list of zones with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) ZonesListAll(filter ListFilter) ([]Zone, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	zones := []Zone{}

	for {
		resp, err := c.zonesList(filter)
		if err != nil {
			return nil, err
		}

		zones = append(zones, resp.Zones...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return zones, nil
		}
	}
}

// Zone retrieves the Zone whose ID it's passed.
func (c *Client) Zone(id string) (Zone, error) {
	zone := &ZoneResponse{}
	err := resourceRequest(c, zoneEP(c, id), "GET", nil, zone)
	if err != nil {
		return Zone{}, err
	}

	return zone.Zone, nil
}

// ZoneByID retrieves the Zone whose ID it's passed.
// It is a version of Zone whose func name is a bit more explicit.
func (c *Client) ZoneByID(id string) (Zone, error) {
	return c.Zone(id)
}

// ZoneByName retrieves the Zone whose name it's passed.
func (c *Client) ZoneByName(name string) (Zone, error) {
	zone := &ZoneResponse{}
	err := resourceRequest(c, zoneNameEP(c, name), "GET", nil, zone)
	if err != nil {
		return Zone{}, err
	}

	return zone.Zone, nil
}

// ZoneCreate creates the Zone it's passed.
func (c *Client) ZoneCreate(z *Zone) (*ZoneUpdateResponse, error) {
	zJSON, err := json.Marshal(z)
	if err != nil {
		return nil, err
	}
	var resource = &ZoneUpdateResponse{}
	err = resourceRequest(c, zonesEP(c), "POST", zJSON, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneUpdate updates the Zone whose ID it's passed.
func (c *Client) ZoneUpdate(z *Zone) (*ZoneUpdateResponse, error) {
	zJSON, err := json.Marshal(z)
	if err != nil {
		return nil, err
	}
	var resource = &ZoneUpdateResponse{}
	err = resourceRequest(c, zoneEP(c, z.ID), "PUT", zJSON, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneDelete deletes the Zone whose ID it's passed.
func (c *Client) ZoneDelete(zoneID string) (*ZoneUpdateResponse, error) {
	resource := &ZoneUpdateResponse{}
	err := resourceRequest(c, zoneEP(c, zoneID), "DELETE", nil, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneExists returns true if a zone request does not 404
// Otherwise, it returns false
func (c *Client) ZoneExists(id string) (bool, error) {
	zone := &ZoneResponse{}
	err := resourceRequest(c, zoneEP(c, id), "GET", nil, zone)
	if err != nil {
		if vErr, ok := err.(*Error); ok {
			if vErr.ResponseCode == http.StatusNotFound {
				return false, nil
			}

			return false, err
		}
	}

	return true, nil
}

// ZoneNameExists returns true if a zone request does not 404
// Otherwise, it returns false
func (c *Client) ZoneNameExists(name string) (bool, error) {
	zone := &ZoneResponse{}
	err := resourceRequest(c, zoneNameEP(c, name), "GET", nil, zone)
	if err != nil {
		if vErr, ok := err.(*Error); ok {
			if vErr.ResponseCode == http.StatusNotFound {
				return false, nil
			}

			return false, err
		}
	}

	return true, nil
}

// ZoneChanges retrieves the ZoneChanges for the Zone whose ID it's passed.
func (c *Client) ZoneChanges(id string) (*ZoneChanges, error) {
	zh := &ZoneChanges{}
	err := resourceRequest(c, zoneChangesEP(c, id, ListFilter{}), "GET", nil, zh)
	if err != nil {
		return &ZoneChanges{}, err
	}

	return zh, nil
}

// ZoneChangesListAll retrieves the complete list of zone changes with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) ZoneChangesListAll(zoneID string, filter ListFilter) ([]ZoneChange, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	changes := []ZoneChange{}

	for {
		resp, err := c.zoneChangesList(zoneID, filter)
		if err != nil {
			return nil, err
		}

		changes = append(changes, resp.ZoneChanges...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return changes, nil
		}
	}
}

// ZoneChange retrieves the ZoneChange matching the Zone ID and
// and ZoneChange ID it's passed.
func (c *Client) ZoneChange(zoneID, zoneChangeID string) (ZoneChange, error) {
	zc := ZoneChange{}
	history, err := c.ZoneChanges(zoneID)
	if err != nil {
		return zc, err
	}

	for _, each := range history.ZoneChanges {
		fmt.Println(each)
		if each.ID == zoneChangeID {
			return each, nil
		}
	}

	return zc, nil
}

// ZoneSync triggers the sync process of VinyIDNS zone info with existing zone
func (c *Client) ZoneSync(zoneId string) (ZoneChange, error) {
	zc := ZoneChange{}
	err := resourceRequest(c, zoneSyncEP(c, zoneId), "POST", nil, &zc)
	if err != nil {
		return zc, err
	}

	return zc, nil
}
