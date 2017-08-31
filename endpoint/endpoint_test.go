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

	"github.com/stretchr/testify/assert"
)

func TestNewEndpoint(t *testing.T) {
	e := NewEndpoint("example.org", "foo.com", "CNAME")
	if e.DNSName != "example.org" || e.Targets[0] != "foo.com" || e.RecordType != RecordTypeCNAME {
		t.Error("endpoint is not initialized correctly")
	}
	if e.Labels == nil {
		t.Error("Labels is not initialized")
	}

	w := NewEndpoint("example.org.", "load-balancer.com.", "")
	if w.DNSName != "example.org" || w.Targets[0] != "load-balancer.com" || w.RecordType != "" {
		t.Error("endpoint is not initialized correctly")
	}
}

func TestMergeLabels(t *testing.T) {
	e := NewEndpoint("abc.com", "1.2.3.4", RecordTypeA)
	e.Labels = map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	e.MergeLabels(map[string]string{"baz": "baz", "new": "fox"})
	assert.Equal(t, map[string]string{"foo": "bar", "baz": "qux", "new": "fox"}, e.Labels)
}

func TestTargetSliceEquals(t *testing.T) {
	targets1 := []string{"1.2.3.4", "1.2.3.5", "1.2.3.6"}
	targets2 := []string{"1.2.3.6", "1.2.3.5", "1.2.3.4"}

	assert.True(t, TargetSliceEquals(targets1, targets2), "targets are equal")

	targets3 := []string{"1.2.3.4", "1.2.3.5", "1.2.3.6", "1.2.3.3"}

	assert.False(t, TargetSliceEquals(targets1, targets3), "targets are not equal")

	assert.True(t, TargetSliceEquals([]string{}, []string{}), "targets are equal")
}
