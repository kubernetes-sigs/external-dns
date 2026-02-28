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
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/pkg/events"
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
			expected: Targets{"8.8.8.8", "::0001", "example.com"},
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
	t.Run("empty provider specific", func(t *testing.T) {
		e := &Endpoint{}
		val, ok := e.GetProviderSpecificProperty("any")
		assert.False(t, ok)
		assert.Empty(t, val)
	})

	t.Run("key is not present in provider specific", func(t *testing.T) {
		e := &Endpoint{
			ProviderSpecific: []ProviderSpecificProperty{
				{Name: "name", Value: "value"},
			},
		}
		val, ok := e.GetProviderSpecificProperty("hello")
		assert.False(t, ok)
		assert.Empty(t, val)
	})

	t.Run("key is present in provider specific", func(t *testing.T) {
		e := &Endpoint{
			ProviderSpecific: []ProviderSpecificProperty{
				{Name: "name", Value: "value"},
			},
		}
		val, ok := e.GetProviderSpecificProperty("name")
		assert.True(t, ok)
		assert.Equal(t, "value", val)
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
			name:     "empty provider specific",
			endpoint: Endpoint{},
			key:      "any",
			expected: nil,
		},
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
				t.Errorf("Endpoint.isOwnedBy() = %v, want %v", got, tt.want)
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
				Targets:    Targets{"10 20 5060 service.example.com."},
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
				Targets:    Targets{"10 20 abc service.example.com."},
			},
			expected: false,
		},
		{
			description: "Invalid SRV record with missing dot for target host",
			endpoint: Endpoint{
				DNSName:    "_service._tls.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 20 5060 service.example.com"},
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
				Targets:    Targets{"10 5 5060 example.com."},
			},
			expected: true,
		},
		{
			description: "Invalid SRV record target",
			endpoint: Endpoint{
				DNSName:    "_service._tcp.example.com",
				RecordType: RecordTypeSRV,
				Targets:    Targets{"10 5 example.com."},
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
		{
			description: "A record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeA,
				Targets:          Targets{"my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "AAAA record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeAAAA,
				Targets:          Targets{"dualstack.my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "CNAME record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeCNAME,
				Targets:          Targets{"d111111abcdef8.cloudfront.net"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "MX record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeMX,
				Targets:          Targets{"10 mail.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "TXT record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeTXT,
				Targets:          Targets{"v=spf1 ~all"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "NS record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeNS,
				Targets:          Targets{"ns1.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "SRV record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "_sip._tcp.example.com",
				RecordType:       RecordTypeSRV,
				Targets:          Targets{"10 5 5060 sip.example.com."},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "MX record with alias=false is also invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeMX,
				Targets:          Targets{"10 mail.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "false"}},
			},
			expected: false,
		},
		{
			description: "MX record without alias property is valid",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 mail.example.com"},
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

func TestCheckEndpoint_AliasWarningLog(t *testing.T) {
	tests := []struct {
		name    string
		ep      Endpoint
		wantLog bool
	}{
		{
			name: "unsupported type with alias logs warning",
			ep: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeMX,
				Targets:          Targets{"10 mail.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			wantLog: true,
		},
		{
			name: "supported type with alias does not log",
			ep: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeA,
				Targets:          Targets{"my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: providerSpecificAlias, Value: "true"}},
			},
			wantLog: false,
		},
		{
			name: "unsupported type without alias does not log",
			ep: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeMX,
				Targets:    Targets{"10 mail.example.com"},
			},
			wantLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, hook := test.NewNullLogger()
			log.AddHook(hook)
			log.SetOutput(logger.Out)
			log.SetLevel(log.WarnLevel)

			tt.ep.CheckEndpoint()

			warnMsg := "does not support alias records"
			found := false
			for _, entry := range hook.AllEntries() {
				if strings.Contains(entry.Message, warnMsg) && entry.Level == log.WarnLevel {
					found = true
				}
			}

			if tt.wantLog {
				assert.True(t, found, "Expected warning log message not found")
			} else {
				assert.False(t, found, "Unexpected warning log message found")
			}
		})
	}
}

func TestEndpoint_WithRefObject(t *testing.T) {
	ep := &Endpoint{}
	ref := &events.ObjectReference{
		Kind:      "Service",
		Namespace: "default",
		Name:      "my-service",
	}
	result := ep.WithRefObject(ref)

	assert.Equal(t, ref, ep.RefObject(), "refObject should be set")
	assert.Equal(t, ep, result, "should return the same Endpoint pointer")
}

func TestTargets_UniqueOrdered(t *testing.T) {
	tests := []struct {
		name     string
		input    Targets
		expected Targets
	}{
		{
			name:     "no duplicates",
			input:    Targets{"a.example.com", "b.example.com"},
			expected: Targets{"a.example.com", "b.example.com"},
		},
		{
			name:     "with duplicates",
			input:    Targets{"a.example.com", "b.example.com", "a.example.com"},
			expected: Targets{"a.example.com", "b.example.com"},
		},
		{
			name:     "all duplicates",
			input:    []string{"a.example.com", "a.example.com", "a.example.com"},
			expected: Targets{"a.example.com"},
		},
		{
			name:     "already sorted",
			input:    Targets{"a.example.com", "c.example.com", "d.example.com"},
			expected: Targets{"a.example.com", "c.example.com", "d.example.com"},
		},
		{
			name:     "unsorted input",
			input:    Targets{"z.example.com", "a.example.com", "m.example.com"},
			expected: Targets{"a.example.com", "m.example.com", "z.example.com"},
		},
		{
			name:     "empty input",
			input:    Targets{},
			expected: Targets{},
		},
		{
			name:     "single element",
			input:    Targets{"only.example.com"},
			expected: Targets{"only.example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewTargets(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndpoint_WithMinTTL(t *testing.T) {
	tests := []struct {
		name         string
		initialTTL   TTL
		inputTTL     int64
		expectedTTL  TTL
		isConfigured bool
	}{
		{
			name:         "sets TTL when not configured and input > 0",
			initialTTL:   0,
			inputTTL:     300,
			expectedTTL:  300,
			isConfigured: true,
		},
		{
			name:         "does not override when already configured",
			initialTTL:   120,
			inputTTL:     300,
			expectedTTL:  120,
			isConfigured: true,
		},
		{
			name:         "does not set when input is zero",
			initialTTL:   30,
			inputTTL:     0,
			expectedTTL:  30,
			isConfigured: true,
		},
		{
			name:        "does not set when input is negative",
			initialTTL:  0,
			inputTTL:    -10,
			expectedTTL: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := &Endpoint{RecordTTL: tt.initialTTL}
			ep.WithMinTTL(tt.inputTTL)
			assert.Equal(t, tt.expectedTTL, ep.RecordTTL)
			assert.Equal(t, tt.isConfigured, ep.RecordTTL.IsConfigured())
		})
	}
}

// TestNewEndpointWithTTLPreservesDotsInTXTRecords tests that trailing dots are preserved in TXT records
func TestNewEndpointWithTTLPreservesDotsInTXTRecords(t *testing.T) {
	// TXT records should preserve trailing dots (and any arbitrary text)
	txtEndpoint := NewEndpointWithTTL("example.com", RecordTypeTXT, TTL(300),
		"v=1;some_signature=aBx3d5..",
		"text.with.dots...",
		"simple-text")

	require.NotNil(t, txtEndpoint, "TXT endpoint should be created")
	require.Len(t, txtEndpoint.Targets, 3, "should have 3 targets")

	// All dots should be preserved in TXT targets
	assert.Equal(t, "v=1;some_signature=aBx3d5..", txtEndpoint.Targets[0])
	assert.Equal(t, "text.with.dots...", txtEndpoint.Targets[1])
	assert.Equal(t, "simple-text", txtEndpoint.Targets[2])

	// Domain name record types should still have trailing dots trimmed
	aEndpoint := NewEndpointWithTTL("example.com", RecordTypeA, TTL(300), "1.2.3.4.")
	require.NotNil(t, aEndpoint, "A endpoint should be created")
	assert.Equal(t, "1.2.3.4", aEndpoint.Targets[0], "A record should have trailing dot trimmed")

	cnameEndpoint := NewEndpointWithTTL("example.com", RecordTypeCNAME, TTL(300), "target.example.com.")
	require.NotNil(t, cnameEndpoint, "CNAME endpoint should be created")
	assert.Equal(t, "target.example.com", cnameEndpoint.Targets[0], "CNAME record should have trailing dot trimmed")
}

func TestGetBoolProviderSpecificProperty(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       Endpoint
		key            string
		expectedValue  bool
		expectedExists bool
	}{
		{
			name:           "key does not exist",
			endpoint:       Endpoint{},
			key:            "nonexistent",
			expectedValue:  false,
			expectedExists: false,
		},
		{
			name: "key exists with true value",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "enabled", Value: "true"},
				},
			},
			key:            "enabled",
			expectedValue:  true,
			expectedExists: true,
		},
		{
			name: "key exists with false value",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "disabled", Value: "false"},
				},
			},
			key:            "disabled",
			expectedValue:  false,
			expectedExists: true,
		},
		{
			name: "key exists with invalid boolean value",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "invalid", Value: "maybe"},
				},
			},
			key:            "invalid",
			expectedValue:  false,
			expectedExists: true,
		},
		{
			name: "key exists with empty value",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "empty", Value: ""},
				},
			},
			key:            "empty",
			expectedValue:  false,
			expectedExists: true,
		},
		{
			name: "key exists with numeric value",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "numeric", Value: "1"},
				},
			},
			key:            "numeric",
			expectedValue:  false,
			expectedExists: true,
		},
		{
			name: "multiple properties, find correct one",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "first", Value: "invalid"},
					{Name: "second", Value: "true"},
					{Name: "third", Value: "false"},
				},
			},
			key:            "second",
			expectedValue:  true,
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := tt.endpoint.GetBoolProviderSpecificProperty(tt.key)
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedExists, exists)
		})
	}
}
