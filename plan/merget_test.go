package plan

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sigs.k8s.io/external-dns/endpoint"
	"testing"
)

func TestTargetMerge(t *testing.T) {

	var get = func(targets []string) *endpoint.Endpoint {
		return &endpoint.Endpoint{
			DNSName:    "example.com",
			RecordType: "A",
			RecordTTL:  30,
			Labels:     map[string]string{mergerKey: mergeTargets},
			Targets:    targets,
		}
	}

	var tests = []struct {
		name      string
		candidate *endpoint.Endpoint
		current   *endpoint.Endpoint
		expected  *endpoint.Endpoint
	}{
		{
			name:      "candidate and current has different A targets",
			candidate: get([]string{"10.0.20.1", "10.0.20.2"}),
			current:   get([]string{"10.0.22.3", "10.0.22.4"}),
			expected: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				RecordTTL:  30,
				Labels:     map[string]string{mergerKey: mergeTargets},
				Targets: endpoint.Targets{
					"10.0.20.1",
					"10.0.20.2",
					"10.0.22.3",
					"10.0.22.4",
				},
			},
		},
		{
			name:      "candidate empty and current A targets",
			candidate: get([]string{}),
			current:   get([]string{"10.0.22.3", "10.0.22.4"}),
			expected: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				RecordTTL:  30,
				Labels:     map[string]string{mergerKey: mergeTargets},
				Targets: endpoint.Targets{
					"10.0.22.3",
					"10.0.22.4",
				},
			},
		},
		{
			name:      "candidate A targets and current empty",
			candidate: get([]string{"10.0.20.1", "10.0.20.2"}),
			current:   get([]string{}),
			expected: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				RecordTTL:  30,
				Labels:     map[string]string{mergerKey: mergeTargets},
				Targets: endpoint.Targets{
					"10.0.20.1",
					"10.0.20.2",
				},
			},
		},
		{
			name:      "candidate empty and current empty",
			candidate: get([]string{}),
			current:   get([]string{}),
			expected: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				RecordTTL:  30,
				Labels:     map[string]string{mergerKey: mergeTargets},
				Targets:    endpoint.Targets{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			merger := &TargetMerger{}
			candidate := merger.Merge(test.candidate, test.current)
			assert.True(t, compareTargetsForMerger(test.expected, candidate))
		})
	}

}

func TestDefaultMerge(t *testing.T) {
	var get = func(targets []string) *endpoint.Endpoint {
		return &endpoint.Endpoint{
			DNSName:    "example.com",
			RecordType: "A",
			RecordTTL:  30,
			Labels:     map[string]string{},
			Targets:    targets,
		}
	}

	var tests = []struct {
		name      string
		candidate *endpoint.Endpoint
		current   *endpoint.Endpoint
		expected  *endpoint.Endpoint
	}{
		{
			name:      "candidate empty and current empty",
			candidate: get([]string{}),
			current:   get([]string{}),
			expected:  get([]string{}),
		},
		{
			name:      "candidate A records and current empty",
			candidate: get([]string{"10.0.0.1", "10.0.0.2"}),
			current:   get([]string{}),
			expected:  get([]string{"10.0.0.1", "10.0.0.2"}),
		},
		{
			name:      "candidate empty and current A records",
			candidate: get([]string{}),
			current:   get([]string{"10.0.20.1", "10.0.20.0"}),
			expected:  get([]string{}),
		},
		{
			name:      "candidate and current A records",
			candidate: get([]string{"10.0.0.1", "10.0.0.2"}),
			current:   get([]string{"10.0.20.1", "10.0.20.0"}),
			expected:  get([]string{"10.0.0.1", "10.0.0.2"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			merger := &DefaulMerger{}
			candidate := merger.Merge(test.candidate, test.current)
			assert.True(t, compareTargetsForMerger(test.expected, candidate))
		})
	}

}

func TestResolveMerger(t *testing.T) {

	var tests = []struct {
		name     string
		labels   map[string]string
		expected EndpointMerger
	}{
		{
			name:     mergeTargets,
			labels:   map[string]string{mergerKey: mergeTargets},
			expected: &TargetMerger{},
		},
		{
			name:     "unknown merger",
			labels:   map[string]string{mergerKey: "unknown"},
			expected: &DefaulMerger{},
		},
		{
			name:     "merger not set",
			labels:   map[string]string{"foo": "blah"},
			expected: &DefaulMerger{},
		},
		{
			name:     "nil labels",
			labels:   nil,
			expected: &DefaulMerger{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			candidate := &endpoint.Endpoint{
				Labels: test.labels,
			}
			m := resolveMerger(candidate)
			if reflect.TypeOf(m) != reflect.TypeOf(test.expected) {
				t.Errorf("expected %T, got %T", test.expected, m)
			}
		})
	}
}

// compare targets for Merger strategy
func compareTargetsForMerger(expected, actual *endpoint.Endpoint) bool {
	m := map[string]bool{}
	for _, v := range expected.Targets {
		m[v] = true
	}
	for _, v := range actual.Targets {
		if _, ok := m[v]; !ok {
			m[v] = false
		}
	}
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}
