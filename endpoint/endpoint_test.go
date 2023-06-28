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
	"reflect"
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

func TestOwnedRecordFilterMatch(t *testing.T) {
	type fields struct {
		ownerID string
	}
	type args struct {
		ep *Endpoint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "no labels",
			fields: fields{ownerID: "foo"},
			args: args{ep: &Endpoint{
				DNSName:    "foo.com",
				RecordType: RecordTypeA,
			}},
			want: false,
		},
		{
			name:   "no owner label",
			fields: fields{ownerID: "foo"},
			args: args{ep: &Endpoint{
				DNSName:    "foo.com",
				RecordType: RecordTypeA,
				Labels:     NewLabels(),
			}},
			want: false,
		},
		{
			name:   "owner not matched",
			fields: fields{ownerID: "foo"},
			args: args{ep: &Endpoint{
				DNSName:    "foo.com",
				RecordType: RecordTypeA,
				Labels: Labels{
					OwnerLabelKey: "bar",
				},
			}},
			want: false,
		},
		{
			name:   "owner matched",
			fields: fields{ownerID: "foo"},
			args: args{ep: &Endpoint{
				DNSName:    "foo.com",
				RecordType: RecordTypeA,
				Labels: Labels{
					OwnerLabelKey: "foo",
				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ownedRecordFilter{
				ownerID: tt.fields.ownerID,
			}
			if got := f.Match(tt.args.ep); got != tt.want {
				t.Errorf("ownedRecordFilter.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyEndpointFilter(t *testing.T) {
	foo1 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo2 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeCNAME,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	bar := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "bar",
		},
	}
	type args struct {
		filter EndpointFilterInterface
		eps    []*Endpoint
	}
	tests := []struct {
		name string
		args args
		want []*Endpoint
	}{
		{
			name: "filter values",
			args: args{
				filter: NewOwnedRecordFilter("foo"),
				eps: []*Endpoint{
					foo1,
					foo2,
					bar,
				},
			},
			want: []*Endpoint{
				foo1,
				foo2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyEndpointFilter(tt.args.filter, tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyEndpointFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
