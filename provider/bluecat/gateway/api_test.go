/*
Copyright 2020 The Kubernetes Authors.
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

package api

import (
	"testing"
)

func TestExpandZones(t *testing.T) {
	mockZones := []string{"example.com", "nginx.example.com", "hack.example.com"}
	expected := []string{"zones/com/zones/example/zones/", "zones/com/zones/example/zones/nginx/zones/", "zones/com/zones/example/zones/hack/zones/"}
	for i := range mockZones {
		if expandZone(mockZones[i]) != expected[i] {
			t.Fatalf("%s", expected[i])
		}
	}
}

func TestValidDeployTypes(t *testing.T) {
	validTypes := []string{"no-deploy", "full-deploy"}
	invalidTypes := []string{"anything-else"}
	for _, i := range validTypes {
		if !IsValidDNSDeployType(i) {
			t.Fatalf("%s should be a valid deploy type", i)
		}
	}
	for _, i := range invalidTypes {
		if IsValidDNSDeployType(i) {
			t.Fatalf("%s should be a invalid deploy type", i)
		}
	}
}
