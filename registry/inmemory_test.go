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

package registry

import "testing"

func TestInMemoryRegistry(t *testing.T) {
	t.Run("NewInMemoryRegistry", testInMemoryInit)
	t.Run("Records", testInMemoryRecords)
	t.Run("ApplyChanges", testInMemoryApplyChanges)
}

func testInMemoryInit(t *testing.T) {
	_, err := NewInMemoryRegistry("", "")
	if err == nil {
		t.Error("should return error if owner id or zone is empty")
	}
	_, err = NewInMemoryRegistry("owner", "")
	if err == nil {
		t.Error("should return error if owner id or zone is empty")
	}
	_, err = NewInMemoryRegistry("", "zone")
	if err == nil {
		t.Error("should return error if owner id or zone is empty")
	}

	p, err := NewInMemoryRegistry("owner", "zone")
	if err != nil {
		t.Error(err)
	}
	if p.ownerID != "owner" || p.zone != "zone" || p.provider == nil {
		t.Error("in-memory registry incorrectly initialized")
	}
}

func testInMemoryRecords(t *testing.T) {

}

func testInMemoryApplyChanges(t *testing.T) {

}
