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

package provider

import "strings"

type zoneIDName map[string]string

func (z zoneIDName) Add(zoneID, zoneName string) {
	z[zoneID] = zoneName
}

func (z zoneIDName) FindZone(hostname string) (suitableZoneID, suitableZoneName string) {
	for zoneID, zoneName := range z {
		if hostname == zoneName || strings.HasSuffix(hostname, "."+zoneName) {
			if suitableZoneName == "" || len(zoneName) > len(suitableZoneName) {
				suitableZoneID = zoneID
				suitableZoneName = zoneName
			}
		}
	}
	return
}
