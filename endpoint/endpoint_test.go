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
		{"dd:dd::01", "::1", "::0001"},
		{"example.org", "EXAMPLE.ORG"},
	}

	for _, d := range tests {
		if d.Same(d) != true {
			t.Errorf("%#v should equal %#v", d, d)
		}
	}
}

func TestSameSuccess(t *testing.T) {
	tests := []struct {
		a Targets
		b Targets
	}{
		{
			[]string{"::1"},
			[]string{"::0001"},
		},
		{
			[]string{"::1", "dd:dd::01"},
			[]string{"dd:00dd::0001", "::0001"},
		},

		{
			[]string{"::1", "dd:dd::01"},
			[]string{"00dd:dd::0001", "::0001"},
		},
		{
			[]string{"::1", "1.1.1.1", "2600.com", "3.3.3.3"},
			[]string{"2600.com", "::0001", "3.3.3.3", "1.1.1.1"},
		},
	}

	for _, d := range tests {
		if d.a.Same(d.b) == false {
			t.Errorf("%#v should equal %#v", d.a, d.b)
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
		{
			[]string{"::1", "2600.com", "3.3.3.3"},
			[]string{"2600.com", "3.3.3.3", "1.1.1.1"},
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

func TestFilterEndpointsByOwnerIDWithRecordTypeA(t *testing.T) {
	foo1 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo2 := &Endpoint{
		DNSName:    "foo2.com",
		RecordType: RecordTypeA,
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
		ownerID string
		eps     []*Endpoint
	}
	tests := []struct {
		name string
		args args
		want []*Endpoint
	}{
		{
			name: "filter values",
			args: args{
				ownerID: "foo",
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
			if got := FilterEndpointsByOwnerID(tt.args.ownerID, tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyEndpointFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEndpointsByOwnerIDWithRecordTypeCNAME(t *testing.T) {
	foo1 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeCNAME,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo2 := &Endpoint{
		DNSName:    "foo2.com",
		RecordType: RecordTypeCNAME,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	bar := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeCNAME,
		Labels: Labels{
			OwnerLabelKey: "bar",
		},
	}
	type args struct {
		ownerID string
		eps     []*Endpoint
	}
	tests := []struct {
		name string
		args args
		want []*Endpoint
	}{
		{
			name: "filter values",
			args: args{
				ownerID: "foo",
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
			if got := FilterEndpointsByOwnerID(tt.args.ownerID, tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyEndpointFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOwnedBy(t *testing.T) {
	type fields struct {
		Labels Labels
	}
	type args struct {
		ownerID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "empty labels",
			fields: fields{Labels: Labels{}},
			args:   args{ownerID: "foo"},
			want:   false,
		},
		{
			name:   "owner label not match",
			fields: fields{Labels: Labels{OwnerLabelKey: "bar"}},
			args:   args{ownerID: "foo"},
			want:   false,
		},
		{
			name:   "owner label match",
			fields: fields{Labels: Labels{OwnerLabelKey: "foo"}},
			args:   args{ownerID: "foo"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				Labels: tt.fields.Labels,
			}
			if got := e.IsOwnedBy(tt.args.ownerID); got != tt.want {
				t.Errorf("Endpoint.IsOwnedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicatedEndpointsWithSimpleZone(t *testing.T) {
	foo1 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo2 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
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
		eps []*Endpoint
	}
	tests := []struct {
		name string
		args args
		want []*Endpoint
	}{
		{
			name: "filter values",
			args: args{
				eps: []*Endpoint{
					foo1,
					foo2,
					bar,
				},
			},
			want: []*Endpoint{
				foo1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicatedEndpointsWithOverlappingZones(t *testing.T) {
	foo1 := &Endpoint{
		DNSName:    "internal.foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo2 := &Endpoint{
		DNSName:    "internal.foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo3 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	foo4 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "foo",
		},
	}
	bar := &Endpoint{
		DNSName:    "internal.foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "bar",
		},
	}
	bar2 := &Endpoint{
		DNSName:    "foo.com",
		RecordType: RecordTypeA,
		Labels: Labels{
			OwnerLabelKey: "bar",
		},
	}

	type args struct {
		eps []*Endpoint
	}
	tests := []struct {
		name string
		args args
		want []*Endpoint
	}{
		{
			name: "filter values",
			args: args{
				eps: []*Endpoint{
					foo1,
					foo2,
					foo3,
					foo4,
					bar,
					bar2,
				},
			},
			want: []*Endpoint{
				foo1,
				foo3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
