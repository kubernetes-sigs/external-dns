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
	targets := []string{"foo.com", "bar.com"}
	e := NewEndpoint("example.org", targets, "CNAME")
	if e.DNSName != "example.org" || e.RecordType != "CNAME" {
		t.Error("endpoint is not initialized correctly")
	}
	if len(e.Targets) != len(targets) {
		t.Fatalf("expected %d target(s), got %d", len(targets), len(e.Targets))
	}
	for i := range e.Targets {
		if e.Targets[i] != targets[i] {
			t.Error("endpoint is not initialized correctly")
		}
	}
	if e.Labels == nil {
		t.Error("Labels is not initialized")
	}

	targets = []string{"load-balancer.com.", "load-balancer.bar."}
	w := NewEndpoint("example.org.", targets, "")
	if w.DNSName != "example.org" || w.RecordType != "" {
		t.Error("endpoint is not initialized correctly")
	}
	if len(w.Targets) != len(targets) {
		t.Fatalf("expected %d target(s), got %d", len(targets), len(w.Targets))
	}
	for i := range []string{"load-balancer.com", "elb.amazonaws.com"} {
		if w.Targets[i] != targets[i] {
			t.Error("endpoint is not initialized correctly")
		}
	}
}

func TestMergeLabels(t *testing.T) {
	e := NewEndpoint("abc.com", []string{"1.2.3.4"}, "A")
	e.Labels = map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	e.MergeLabels(map[string]string{"baz": "baz", "new": "fox"})
	assert.Equal(t, map[string]string{"foo": "bar", "baz": "qux", "new": "fox"}, e.Labels)
}
