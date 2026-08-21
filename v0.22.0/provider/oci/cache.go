/*
Copyright 2023 The Kubernetes Authors.

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

package oci

import (
	"time"

	"github.com/oracle/oci-go-sdk/v65/dns"
)

type zoneCache struct {
	age      time.Time
	duration time.Duration
	zones    map[string]dns.ZoneSummary
}

func (z *zoneCache) Reset(zones map[string]dns.ZoneSummary) {
	if z.duration > time.Duration(0) {
		z.age = time.Now()
		z.zones = zones
	}
}

func (z *zoneCache) Get() map[string]dns.ZoneSummary {
	return z.zones
}

func (z *zoneCache) Expired() bool {
	return len(z.zones) < 1 || time.Since(z.age) > z.duration
}
