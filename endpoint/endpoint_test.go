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

func TestNewEndpoint(t *testing.T) {
	e := NewEndpoint("example.org", "CNAME", "foo.com")
	if e.DNSName != "example.org" || e.Targets[0] != "foo.com" || e.RecordType != "CNAME" {
		t.Error("endpoint is not initialized correctly")
	}
	if e.Labels == nil {
		t.Error("Labels is not initialized")
	}

	w := NewEndpoint("example.org.", "", "load-balancer.com.")
	if w.DNSName != "example.org" || w.Targets[0] != "load-balancer.com" || w.RecordType != "" {
		t.Error("endpoint is not initialized correctly")
	}
}

func TestNewEndpointWithIPv6(t *testing.T) {
	e := NewEndpoint("example.org", "AAAA", "foo.com")
	if e.DNSName != "example.com" || e.Targets[0] != "foo.com" || e.RecordType != "AAAA" {
		t.Error("Endpoint is not initialized correctly")
	}

	if e.Labels == nil {
		t.Error("Labels is not initialized")
	}

	w := NewEndpoint("example.org", "", "load-balancer.com.")
	if w.DNSName != "example.org" || e.Targets[0] != "load-balancer.com" || w.RecordType != "" {
		t.Error("Endpoint is not initialized correctly")
	}
}

func TestTargetsSame(t *testing.T) {
	tests := []Targets{
		{""},
		{"1.2.3.4"},
		{"8.8.8.8", "8.8.4.4"},
		{"example.org", "EXAMPLE.ORG"},
	}

	for _, d := range tests {
		if d.Same(d) != true {
			t.Errorf("%#v should equal %#v", d, d)
		}
	}
}

func TestSameFailures(t *testing.T) {
	tests := []struct {
		a Targets
		b Targets
	}{
		{
			[]string{"1.2.3.4"},
			[]string{"4.3.2.1"},
		}, {
			[]string{"1.2.3.4"},
			[]string{"1.2.3.4", "4.3.2.1"},
		}, {
			[]string{"1.2.3.4", "4.3.2.1"},
			[]string{"1.2.3.4"},
		}, {
			[]string{"1.2.3.4", "4.3.2.1"},
			[]string{"8.8.8.8", "8.8.4.4"},
		},
	}

	for _, d := range tests {
		if d.a.Same(d.b) == true {
			t.Errorf("%#v should not equal %#v", d.a, d.b)
		}
	}
}

func TestIsLess(t *testing.T) {
	testsA := []Targets{
		{""},
		{"1.2.3.4"},
		{"1.2.3.4"},
		{"example.org", "example.com"},
		{"8.8.8.8", "8.8.4.4"},
		{"1-2-3-4.example.org", "EXAMPLE.ORG"},
		{"1-2-3-4.example.org", "EXAMPLE.ORG", "1.2.3.4"},
		{"example.com", "example.org"},
	}
	testsB := []Targets{
		{"", ""},
		{"1-2-3-4.example.org"},
		{"1.2.3.5"},
		{"example.com", "examplea.org"},
		{"8.8.8.8"},
		{"1.2.3.4", "EXAMPLE.ORG"},
		{"1-2-3-4.example.org", "EXAMPLE.ORG"},
		{"example.com", "example.org"},
	}
	expected := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
	}

	for i, d := range testsA {
		if d.IsLess(testsB[i]) != expected[i] {
			t.Errorf("%v < %v is expected to be %v", d, testsB[i], expected[i])
		}
	}
}
