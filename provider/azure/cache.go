/*
Copyright 2024 The Kubernetes Authors.

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

package azure

import (
	"time"
)

// zonesCache is a cache for Azure zones(private or public)
type zonesCache[T any] struct {
	age      time.Time
	duration time.Duration
	zones    []T
}

// Reset method to reset the zones and update the age. This will be used to update the cache
// after making a new API call to get the zones.
func (z *zonesCache[T]) Reset(zones []T) {
	if z.duration > time.Duration(0) {
		z.age = time.Now()
		z.zones = zones
	}
}

// Get method to retrieve the cached zones. If cache is not expired, this will be used
// instead of making a new API call to get the zones.
func (z *zonesCache[T]) Get() []T {
	return z.zones
}

// Expired method to check if the cache has expired based on duration or if zones are empty.
// If cache is expired, a new API call will be made to get the zones. If zones are empty, a new
// API call will be made to get the zones. This case comes in at the time of initialization.
func (z *zonesCache[T]) Expired() bool {
	return len(z.zones) < 1 || time.Since(z.age) > z.duration
}
