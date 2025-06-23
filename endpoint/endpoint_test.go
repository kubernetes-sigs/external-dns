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
	"fmt"
	"reflect"
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

func TestNewTargets(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected Targets
	}{
		{
			name:     "no targets",
			input:    []string{},
			expected: Targets{},
		},
		{
			name:     "single target",
			input:    []string{"1.2.3.4"},
			expected: Targets{"1.2.3.4"},
		},
		{
			name:     "multiple targets",
			input:    []string{"example.com", "8.8.8.8", "::0001"},
			expected: Targets{"example.com", "8.8.8.8", "::0001"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Targets := NewTargets(c.input...)
			changedTarget := Targets.String()
			assert.Equal(t, c.expected.String(), changedTarget)

		})
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

func TestGetProviderSpecificProperty(t *testing.T) {
	e := &Endpoint{
		ProviderSpecific: []ProviderSpecificProperty{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	t.Run("key is not present in provider specific", func(t *testing.T) {
		val, ok := e.GetProviderSpecificProperty("hello")
		assert.False(t, ok)
		assert.Empty(t, val)
	})

	t.Run("key is present in provider specific", func(t *testing.T) {
		val, ok := e.GetProviderSpecificProperty("name")
		assert.True(t, ok)
		assert.NotEmpty(t, val)

	})
}

func TestSetProviderSpecficProperty(t *testing.T) {
	cases := []struct {
		name               string
		endpoint           Endpoint
		key                string
		value              string
		expectedIdentifier string
		expected           []ProviderSpecificProperty
	}{
		{
			name:     "endpoint is empty",
			endpoint: Endpoint{},
			key:      "key1",
			value:    "value1",
			expected: []ProviderSpecificProperty{
				{
					Name:  "key1",
					Value: "value1",
				},
			},
		},
		{
			name: "name and key are not matching",
			endpoint: Endpoint{
				DNSName:       "example.org",
				RecordTTL:     TTL(0),
				RecordType:    RecordTypeA,
				SetIdentifier: "newIdentifier",
				Targets: Targets{
					"example.org", "example.com", "1.2.4.5",
				},
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
				},
			},
			expectedIdentifier: "newIdentifier",
			key:                "name2",
			value:              "value2",

			expected: []ProviderSpecificProperty{
				{
					Name:  "name1",
					Value: "value1",
				},
				{
					Name:  "name2",
					Value: "value2",
				},
			},
		},
		{
			name: "some keys are matching and some are not matching ",
			endpoint: Endpoint{
				DNSName:       "example.org",
				RecordTTL:     TTL(0),
				RecordType:    RecordTypeA,
				SetIdentifier: "newIdentifier",
				Targets: Targets{
					"example.org", "example.com", "1.2.4.5",
				},
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
					{
						Name:  "name2",
						Value: "value2",
					},
					{
						Name:  "name3",
						Value: "value3",
					},
				},
			},
			key:                "name2",
			value:              "value2",
			expectedIdentifier: "newIdentifier",
			expected: []ProviderSpecificProperty{
				{
					Name:  "name1",
					Value: "value1",
				},
				{
					Name:  "name2",
					Value: "value2",
				},
				{
					Name:  "name3",
					Value: "value3",
				},
			},
		},
		{
			name: "name and key are not matching",
			endpoint: Endpoint{
				DNSName:       "example.org",
				RecordTTL:     TTL(0),
				RecordType:    RecordTypeA,
				SetIdentifier: "identifier",
				Targets: Targets{
					"example.org", "example.com", "1.2.4.5",
				},
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
				},
			},
			key:                "name1",
			value:              "value2",
			expectedIdentifier: "identifier",
			expected: []ProviderSpecificProperty{
				{
					Name:  "name1",
					Value: "value2",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.endpoint.WithProviderSpecific(c.key, c.value)
			expectedString := fmt.Sprintf("%s %d IN %s %s %s %s", c.endpoint.DNSName, c.endpoint.RecordTTL, c.endpoint.RecordType, c.endpoint.SetIdentifier, c.endpoint.Targets, c.endpoint.ProviderSpecific)
			identifier := c.endpoint.WithSetIdentifier(c.endpoint.SetIdentifier)
			assert.Equal(t, c.expectedIdentifier, identifier.SetIdentifier)
			assert.Equal(t, expectedString, c.endpoint.String())
			if !reflect.DeepEqual([]ProviderSpecificProperty(c.endpoint.ProviderSpecific), c.expected) {
				t.Errorf("unexpected ProviderSpecific:\nGot:      %#v\nExpected: %#v", c.endpoint.ProviderSpecific, c.expected)
			}
		})
	}
}

func TestDeleteProviderSpecificProperty(t *testing.T) {
	cases := []struct {
		name     string
		endpoint Endpoint
		key      string
		expected []ProviderSpecificProperty
	}{
		{
			name: "name and key are not matching",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
				},
			},
			key: "name2",
			expected: []ProviderSpecificProperty{
				{
					Name:  "name1",
					Value: "value1",
				},
			},
		},
		{
			name: "some keys are matching and some keys are not matching",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
					{
						Name:  "name2",
						Value: "value2",
					},
					{
						Name:  "name3",
						Value: "value3",
					},
				},
			},
			key: "name2",
			expected: []ProviderSpecificProperty{
				{
					Name:  "name1",
					Value: "value1",
				},
				{
					Name:  "name3",
					Value: "value3",
				},
			},
		},
		{
			name: "name and key are matching",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{
						Name:  "name1",
						Value: "value1",
					},
				},
			},
			key:      "name1",
			expected: []ProviderSpecificProperty{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.endpoint.DeleteProviderSpecificProperty(c.key)
			if !reflect.DeepEqual([]ProviderSpecificProperty(c.endpoint.ProviderSpecific), c.expected) {
				t.Errorf("unexpected ProviderSpecific:\nGot:      %#v\nExpected: %#v", c.endpoint.ProviderSpecific, c.expected)
			}
		})
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

func TestPDNScheckEndpoint(t *testing.T) {
	tests := []struct {
		description string
		endpoint    Endpoint
		expected    bool
	}{
		{
			description: "Valid MX record target",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 example.com"},
			},
			expected: true,
		},
		{
			description: "Valid MX record with multiple targets",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 example.com", "20 backup.example.com"},
			},
			expected: true,
		},
		{
			description: "MX record with valid and invalid targets",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"example.com", "backup.example.com"},
			},
			expected: false,
		},
		{
			description: "Invalid MX record with missing priority value",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"example.com"},
			},
			expected: false,
		},
		{
			description: "Invalid MX record with too many arguments",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 example.com abc"},
			},
			expected: false,
		},
		{
			description: "Invalid MX record with non-integer priority",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"abc example.com"},
			},
			expected: false,
		},
		{
			description: "Valid SRV record target",
			endpoint: Endpoint{
				DNSName:    "_service._tls.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 20 5060 service.example.com"},
			},
			expected: true,
		},
		{
			description: "Invalid SRV record with missing part",
			endpoint: Endpoint{
				DNSName:    "_service._tls.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 20 5060"},
			},
			expected: false,
		},
		{
			description: "Invalid SRV record with non-integer part",
			endpoint: Endpoint{
				DNSName:    "_service._tls.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 20 abc service.example.com"},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := tt.endpoint.CheckEndpoint()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestNewMXTarget(t *testing.T) {
	tests := []struct {
		description string
		target      string
		expected    *MXTarget
		expectError bool
	}{
		{
			description: "Valid MX record",
			target:      "10 example.com",
			expected:    &MXTarget{priority: 10, host: "example.com"},
			expectError: false,
		},
		{
			description: "Invalid MX record with missing priority",
			target:      "example.com",
			expectError: true,
		},
		{
			description: "Invalid MX record with non-integer priority",
			target:      "abc example.com",
			expectError: true,
		},
		{
			description: "Invalid MX record with too many parts",
			target:      "10 example.com extra",
			expectError: true,
		},
		{
			description: "Missing host",
			target:      "10 ",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := NewMXRecord(tt.target)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestCheckEndpoint(t *testing.T) {
	tests := []struct {
		description string
		endpoint    Endpoint
		expected    bool
	}{
		{
			description: "Valid MX record target",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 example.com"},
			},
			expected: true,
		},
		{
			description: "Invalid MX record target",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"example.com"},
			},
			expected: false,
		},
		{
			description: "Valid SRV record target",
			endpoint: Endpoint{
				DNSName:    "_service._tcp.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 5 5060 example.com"},
			},
			expected: true,
		},
		{
			description: "Invalid SRV record target",
			endpoint: Endpoint{
				DNSName:    "_service._tcp.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 5 example.com"},
			},
			expected: false,
		},
		{
			description: "Non-MX/SRV record type",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeA,
				Targets:    Targets{"192.168.1.1"},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := tt.endpoint.CheckEndpoint()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
