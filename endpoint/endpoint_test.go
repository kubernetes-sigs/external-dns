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
	"slices"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
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

func TestEndpoint_Describe(t *testing.T) {
	ep := &Endpoint{
		DNSName:       "example.com",
		SetIdentifier: "owner-1",
		RecordType:    RecordTypeA,
		Targets:       Targets{"1.2.3.4", "5.6.7.8"},
	}
	assert.Equal(t, "record:example.com, owner:owner-1, type:A, targets:1.2.3.4, 5.6.7.8", ep.Describe())
}

func TestEndpoint_Getters(t *testing.T) {
	ep := &Endpoint{
		DNSName:    "example.com",
		RecordType: RecordTypeA,
		RecordTTL:  TTL(300),
		Targets:    Targets{"1.2.3.4", "5.6.7.8"},
	}
	t.Run("GetDNSName", func(t *testing.T) {
		assert.Equal(t, "example.com", ep.GetDNSName())
	})
	t.Run("GetRecordType", func(t *testing.T) {
		assert.Equal(t, RecordTypeA, ep.GetRecordType())
	})
	t.Run("GetRecordTTL", func(t *testing.T) {
		assert.Equal(t, int64(300), ep.GetRecordTTL())
	})
	t.Run("GetTargets", func(t *testing.T) {
		assert.Equal(t, []string{"1.2.3.4", "5.6.7.8"}, ep.GetTargets())
	})
}

func TestEndpoint_WithLabel(t *testing.T) {
	t.Run("nil Labels map is initialised", func(t *testing.T) {
		ep := &Endpoint{} // Labels is nil
		result := ep.WithLabel("key", "value")
		assert.Equal(t, "value", ep.Labels["key"])
		assert.Same(t, ep, result)
	})

	t.Run("existing Labels map is updated", func(t *testing.T) {
		ep := NewEndpoint("example.com", RecordTypeA, "1.2.3.4") // Labels already initialised
		ep.WithLabel("key", "value")
		assert.Equal(t, "value", ep.Labels["key"])
	})
}

func TestSame_ParseErrorLogged(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	// Two different hostnames: neither parses as IP, triggering both err != nil branches.
	result := Targets{"a.example.com"}.Same(Targets{"b.example.com"})

	assert.False(t, result)
	logtest.TestHelperLogContains("Couldn't parse", hook, t)
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

func TestRetainProviderProperties(t *testing.T) {
	cases := []struct {
		name     string
		endpoint Endpoint
		provider string
		expected []ProviderSpecificProperty
	}{
		{
			name:     "empty provider specific",
			endpoint: Endpoint{},
			provider: "aws",
			expected: nil,
		},
		{
			name: "empty provider, properties untouched",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "aws/evaluate-target-health", Value: "true"},
					{Name: "coredns/group", Value: "my-group"},
				},
			},
			provider: "",
			expected: []ProviderSpecificProperty{
				{Name: "aws/evaluate-target-health", Value: "true"},
				{Name: "coredns/group", Value: "my-group"},
			},
		},
		{
			name: "all properties match provider",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "aws/evaluate-target-health", Value: "true"},
					{Name: "aws/weight", Value: "10"},
				},
			},
			provider: "aws",
			expected: []ProviderSpecificProperty{
				{Name: "aws/evaluate-target-health", Value: "true"},
				{Name: "aws/weight", Value: "10"},
			},
		},
		{
			name: "no properties match provider",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "coredns/group", Value: "my-group"},
				},
			},
			provider: "aws",
			expected: []ProviderSpecificProperty{},
		},
		{
			name: "mixed providers, only configured provider retained",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "aws/evaluate-target-health", Value: "true"},
					{Name: "coredns/group", Value: "my-group"},
					{Name: "aws/weight", Value: "10"},
				},
			},
			provider: "aws",
			expected: []ProviderSpecificProperty{
				{Name: "aws/evaluate-target-health", Value: "true"},
				{Name: "aws/weight", Value: "10"},
			},
		},
		{
			name: "provider agnostic properties without prefix are retained",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: ProviderSpecificAlias, Value: "true"},
					{Name: "aws/evaluate-target-health", Value: "true"},
					{Name: "coredns/group", Value: "my-group"},
				},
			},
			provider: "aws",
			expected: []ProviderSpecificProperty{
				{Name: ProviderSpecificAlias, Value: "true"},
				{Name: "aws/evaluate-target-health", Value: "true"},
			},
		},
		{
			name: "provider prefix must match exactly, not as substring",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "aws-extended/some-prop", Value: "val"},
					{Name: "aws/weight", Value: "10"},
				},
			},
			provider: "aws",
			expected: []ProviderSpecificProperty{
				{Name: "aws/weight", Value: "10"},
			},
		},
		// cloudflare uses annotation-style names (e.g. "external-dns.alpha.kubernetes.io/cloudflare-*")
		// rather than the standard "provider/" prefix, so all properties are retained and only sorted.
		{
			name: "cloudflare retains all properties",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
					{Name: "aws/evaluate-target-health", Value: "true"},
					{Name: ProviderSpecificAlias, Value: "false"},
				},
			},
			provider: "cloudflare",
			expected: []ProviderSpecificProperty{
				{Name: ProviderSpecificAlias, Value: "false"},
				{Name: "aws/evaluate-target-health", Value: "true"},
				{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
			},
		},
		{
			name: "cloudflare properties are sorted",
			endpoint: Endpoint{
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", Value: "true"},
					{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
				},
			},
			provider: "cloudflare",
			expected: []ProviderSpecificProperty{
				{Name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", Value: "true"},
				{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.endpoint.RetainProviderProperties(c.provider)
			require.Equal(t, c.expected, []ProviderSpecificProperty(c.endpoint.ProviderSpecific))
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

func TestFilterEndpointsByOwnerID_Logs(t *testing.T) {
	const (
		msgMismatch = "owner id does not match"
		msgMissing  = "missing owner label"
	)

	matching := &Endpoint{DNSName: "foo.com", RecordType: RecordTypeA, Labels: Labels{OwnerLabelKey: "foo"}}
	mismatch := &Endpoint{DNSName: "bar.com", RecordType: RecordTypeA, Labels: Labels{OwnerLabelKey: "bar"}}
	noLabel := &Endpoint{DNSName: "baz.com", RecordType: RecordTypeA}

	tests := []struct {
		name     string
		eps      []*Endpoint
		wantLogs []string
	}{
		{
			name: "no log: all endpoints match owner",
			eps:  []*Endpoint{matching},
		},
		{
			name:     "logs owner mismatch",
			eps:      []*Endpoint{matching, mismatch},
			wantLogs: []string{msgMismatch},
		},
		{
			name:     "logs missing owner label",
			eps:      []*Endpoint{matching, noLabel},
			wantLogs: []string{msgMissing},
		},
		{
			name:     "logs both mismatch and missing label",
			eps:      []*Endpoint{matching, mismatch, noLabel},
			wantLogs: []string{msgMismatch, msgMissing},
		},
	}

	allMsgs := []string{msgMismatch, msgMissing}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			FilterEndpointsByOwnerID("foo", tt.eps)
			for _, msg := range allMsgs {
				if slices.Contains(tt.wantLogs, msg) {
					logtest.TestHelperLogContainsWithLogLevel(msg, log.DebugLevel, hook, t)
				} else {
					logtest.TestHelperLogNotContains(msg, hook, t)
				}
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

func TestMXTarget_Getters(t *testing.T) {
	m, err := NewMXRecord("10 mail.example.com")
	require.NoError(t, err)
	assert.Equal(t, uint16(10), *m.GetPriority())
	assert.Equal(t, "mail.example.com", *m.GetHost())
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
			description: "Valid AAAA record",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeAAAA,
				Targets:    Targets{"2001:db8::1"},
			},
			expected: true,
		},
		{
			description: "Invalid A record - not an IP",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeA,
				Targets:    Targets{"not-an-ip"},
			},
			expected: false,
		},
		{
			description: "Invalid A record - IPv6 address",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeA,
				Targets:    Targets{"2001:db8::1"},
			},
			expected: false,
		},
		{
			description: "Invalid AAAA record - IPv4 address",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeAAAA,
				Targets:    Targets{"192.168.1.1"},
			},
			expected: false,
		},
		{
			description: "Invalid AAAA record - not an IP",
			endpoint: Endpoint{
				DNSName:    "example.com",
				RecordType: RecordTypeAAAA,
				Targets:    Targets{"not-an-ip"},
			},
			expected: false,
		},
		{
			description: "A record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeA,
				Targets:          Targets{"my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "AAAA record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeAAAA,
				Targets:          Targets{"dualstack.my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "CNAME record with alias=true is valid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeCNAME,
				Targets:          Targets{"d111111abcdef8.cloudfront.net"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: true,
		},
		{
			description: "MX record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeMX,
				Targets:          Targets{"10 mail.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "TXT record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeTXT,
				Targets:          Targets{"v=spf1 ~all"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "NS record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeNS,
				Targets:          Targets{"ns1.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "SRV record with alias=true is invalid",
			endpoint: Endpoint{
				DNSName:          "_sip._tcp.example.com",
				RecordType:       RecordTypeSRV,
				Targets:          Targets{"10 5 5060 sip.example.com."},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			expected: false,
		},
		{
			description: "MX record with alias=false is also invalid",
			endpoint: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeMX,
				Targets:          Targets{"10 mail.example.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "false"}},
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
		{
			description: "Valid PTR record with in-addr.arpa",
			endpoint: Endpoint{
				DNSName:    "2.49.168.192.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			expected: true,
		},
		{
			description: "Valid PTR record with ip6.arpa",
			endpoint: Endpoint{
				DNSName:    "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"v6.example.com"},
			},
			expected: true,
		},
		{
			description: "Valid PTR record with multiple hostname targets",
			endpoint: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"a.example.com", "b.example.com"},
			},
			expected: true,
		},
		{
			description: "Invalid PTR record - DNS name not reverse DNS",
			endpoint: Endpoint{
				DNSName:    "web.example.com",
				RecordType: RecordTypePTR,
				Targets:    Targets{"10.0.0.1"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - target is an IP address",
			endpoint: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"10.0.0.1"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - target is an IPv6 address",
			endpoint: Endpoint{
				DNSName:    "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"2001:db8::1"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - empty target",
			endpoint: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{""},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - no targets",
			endpoint: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - bare in-addr.arpa",
			endpoint: Endpoint{
				DNSName:    "in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - dot-prefixed in-addr.arpa",
			endpoint: Endpoint{
				DNSName:    ".in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - bare ip6.arpa",
			endpoint: Endpoint{
				DNSName:    "ip6.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			expected: false,
		},
		{
			description: "Invalid PTR record - dot-prefixed ip6.arpa",
			endpoint: Endpoint{
				DNSName:    ".ip6.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			expected: false,
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
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
			},
			wantLog: true,
		},
		{
			name: "supported type with alias does not log",
			ep: Endpoint{
				DNSName:          "example.com",
				RecordType:       RecordTypeA,
				Targets:          Targets{"my-elb-123.us-east-1.elb.amazonaws.com"},
				ProviderSpecific: ProviderSpecific{{Name: ProviderSpecificAlias, Value: "true"}},
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
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			tt.ep.CheckEndpoint()

			warnMsg := "does not support alias records"
			if tt.wantLog {
				logtest.TestHelperLogContains(warnMsg, hook, t)
			} else {
				logtest.TestHelperLogNotContains(warnMsg, hook, t)
			}
		})
	}
}

func TestCheckEndpoint_PTRValidationLog(t *testing.T) {
	tests := []struct {
		name    string
		ep      Endpoint
		wantLog string
	}{
		{
			name: "non-reverse DNS name logs invalid",
			ep: Endpoint{
				DNSName:    "web.example.com",
				RecordType: RecordTypePTR,
				Targets:    Targets{"other.example.com"},
			},
			wantLog: "must be a valid reverse DNS name",
		},
		{
			name: "IP address target logs invalid",
			ep: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"10.0.0.1"},
			},
			wantLog: "must be a hostname, not an IP address",
		},
		{
			name: "empty target logs invalid",
			ep: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{""},
			},
			wantLog: "target must not be empty",
		},
		{
			name: "no targets logs invalid",
			ep: Endpoint{
				DNSName:    "1.0.0.10.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{},
			},
			wantLog: "at least one target is required",
		},
		{
			name: "valid PTR does not log",
			ep: Endpoint{
				DNSName:    "2.49.168.192.in-addr.arpa",
				RecordType: RecordTypePTR,
				Targets:    Targets{"web.example.com"},
			},
			wantLog: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			tt.ep.CheckEndpoint()

			if tt.wantLog != "" {
				logtest.TestHelperLogContains(tt.wantLog, hook, t)
			} else {
				logtest.TestHelperLogNotContains("Invalid PTR record", hook, t)
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

func TestGetOwnerId(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *Endpoint
		expected string
	}{
		{
			name: "owner label is set",
			endpoint: &Endpoint{
				Labels: Labels{
					OwnerLabelKey: "my-owner",
				},
			},
			expected: "my-owner",
		},
		{
			name: "owner label is empty string",
			endpoint: &Endpoint{
				Labels: Labels{
					OwnerLabelKey: "",
				},
			},
			expected: "",
		},
		{
			name: "owner label is not set",
			endpoint: &Endpoint{
				Labels: Labels{
					"other-label": "value",
				},
			},
			expected: "",
		},
		{
			name: "labels map is empty",
			endpoint: &Endpoint{
				Labels: Labels{},
			},
			expected: "",
		},
		{
			name: "labels map is nil",
			endpoint: &Endpoint{
				Labels: nil,
			},
			expected: "",
		},
		{
			name: "multiple labels with owner",
			endpoint: &Endpoint{
				Labels: Labels{
					OwnerLabelKey: "owner-123",
					"other-key":   "other-value",
				},
			},
			expected: "owner-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.endpoint.GetOwner()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetNakedDomain(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *Endpoint
		expected string
	}{
		{
			name: "standard subdomain",
			endpoint: &Endpoint{
				DNSName: "www.example.com",
			},
			expected: "example.com",
		},
		{
			name: "nested subdomain",
			endpoint: &Endpoint{
				DNSName: "api.v1.example.com",
			},
			expected: "v1.example.com",
		},
		{
			name: "root domain only",
			endpoint: &Endpoint{
				DNSName: "example.com",
			},
			expected: "example.com",
		},
		{
			name: "single label (no dots)",
			endpoint: &Endpoint{
				DNSName: "localhost",
			},
			expected: "localhost",
		},
		{
			name: "empty DNS name",
			endpoint: &Endpoint{
				DNSName: "",
			},
			expected: "",
		},
		{
			name: "deeply nested subdomain",
			endpoint: &Endpoint{
				DNSName: "a.b.c.d.example.com",
			},
			expected: "b.c.d.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.endpoint.GetNakedDomain()
			assert.Equal(t, tt.expected, result)

		})
	}
}

func TestRequestedRecordType(t *testing.T) {
	ep := NewEndpoint("example.com", RecordTypeA, "1.2.3.4").
		WithProviderSpecific(ProviderSpecificRecordType, "ptr")
	val, ok := ep.RequestedRecordType()
	assert.True(t, ok)
	assert.Equal(t, "ptr", val)

	ep2 := NewEndpoint("example.com", RecordTypeA, "1.2.3.4")
	_, ok = ep2.RequestedRecordType()
	assert.False(t, ok)
}

func TestNewPTREndpoint(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		ttl       TTL
		hostnames []string
		wantName  string
		wantErr   bool
	}{
		{
			name:      "IPv4",
			target:    "192.168.49.2",
			ttl:       300,
			hostnames: []string{"web.example.com"},
			wantName:  "2.49.168.192.in-addr.arpa",
		},
		{
			name:      "IPv6",
			target:    "2001:db8::1",
			ttl:       600,
			hostnames: []string{"v6.example.com"},
			wantName:  "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
		},
		{
			name:      "multiple hostnames",
			target:    "10.0.0.1",
			ttl:       60,
			hostnames: []string{"a.example.com", "b.example.com"},
			wantName:  "1.0.0.10.in-addr.arpa",
		},
		{
			name:    "invalid target",
			target:  "not-an-ip",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep, err := NewPTREndpoint(tt.target, tt.ttl, tt.hostnames...)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantName, ep.DNSName)
			assert.Equal(t, RecordTypePTR, ep.RecordType)
			assert.Equal(t, tt.ttl, ep.RecordTTL)
			assert.Equal(t, Targets(tt.hostnames), ep.Targets)
		})
	}
}

func TestEndpointKey_String(t *testing.T) {
	tests := []struct {
		name string
		key  EndpointKey
		want string
	}{
		{
			name: "empty key",
			key:  EndpointKey{},
			want: `{"" "" "" "0" ""}`},
		{
			name: "complete key",
			key: EndpointKey{
				DNSName:       "example.com",
				RecordType:    RecordTypeA,
				SetIdentifier: "test-set",
				RecordTTL:     300,
				Target:        "127.0.0.1",
			},
			want: `{"example.com" "A" "test-set" "300" "127.0.0.1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.key.String())
		})
	}
}
