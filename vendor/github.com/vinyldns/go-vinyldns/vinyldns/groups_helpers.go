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

// groupsList retrieves the list of zones with the List criteria passed.
func (c *Client) groupsList(filter ListFilter) (*Groups, error) {
	groups := &Groups{}
	err := resourceRequest(c, groupsListEP(c, filter), "GET", nil, groups)
	if err != nil {
		return groups, err
	}

	return groups, nil
}
