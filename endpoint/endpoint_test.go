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

package endpoint

import (
	"testing"
)

func TestOwner(t *testing.T) {
	e := &Endpoint{}
	if e.Owner() != "" {
		t.Errorf("Owner should be empty if Labels not initialized")
	}
	e.Labels = map[string]string{}
	e.Labels[ownerKey] = ""
	if e.Owner() != "" {
		t.Errorf("Owner should be empty for records without owners")
	}
	e.Labels[ownerKey] = "owner"
	if e.Owner() != "owner" {
		t.Errorf("Incorrect owner is returned")
	}
}

func TestSetOwner(t *testing.T) {
	e := &Endpoint{}
	e.SetOwner("")
	if e.Labels == nil {
		t.Fatalf("Labels map should be initialized")
	}
	if e.Labels[ownerKey] != "" {
		t.Errorf("Owner set incorrectly")
	}
	e.SetOwner("owner")
	if e.Labels[ownerKey] != "owner" {
		t.Errorf("Owner set incorrectly")
	}
}
