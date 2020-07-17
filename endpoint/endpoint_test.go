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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestTargetsSame(t *testing.T) {
	tests := []Targets{
		{""},
		{"1.2.3.4"},
		{"8.8.8.8", "8.8.4.4"},
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

func TestDigitalOceanMergeRecordsByNameType(t *testing.T) {
	xs := []*Endpoint{
		NewEndpoint("foo.example.com", "A", "1.2.3.4"),
		NewEndpoint("bar.example.com", "A", "1.2.3.4"),
		NewEndpoint("foo.example.com", "A", "5.6.7.8"),
		NewEndpoint("foo.example.com", "CNAME", "somewhere.out.there.com"),
	}

	merged := MergeEndpointsByNameType(xs)

	assert.Equal(t, 3, len(merged))
	sort.SliceStable(merged, func(i, j int) bool {
		if merged[i].DNSName != merged[j].DNSName {
			return merged[i].DNSName < merged[j].DNSName
		}
		return merged[i].RecordType < merged[j].RecordType
	})
	assert.Equal(t, "bar.example.com", merged[0].DNSName)
	assert.Equal(t, "A", merged[0].RecordType)
	assert.Equal(t, 1, len(merged[0].Targets))
	assert.Equal(t, "1.2.3.4", merged[0].Targets[0])

	assert.Equal(t, "foo.example.com", merged[1].DNSName)
	assert.Equal(t, "A", merged[1].RecordType)
	assert.Equal(t, 2, len(merged[1].Targets))
	assert.ElementsMatch(t, []string{"1.2.3.4", "5.6.7.8"}, merged[1].Targets)

	assert.Equal(t, "foo.example.com", merged[2].DNSName)
	assert.Equal(t, "CNAME", merged[2].RecordType)
	assert.Equal(t, 1, len(merged[2].Targets))
	assert.Equal(t, "somewhere.out.there.com", merged[2].Targets[0])
}
