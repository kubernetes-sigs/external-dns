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

package registry

import (
	"errors"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"

	"github.com/stretchr/testify/assert"
)

// testTXTRegistryImplementsSource tests that serviceSource is a valid Source.
func TestTXTRegistry(t *testing.T) {
	t.Run("Interface", testTXTRegistryImplementsRegistry)
	t.Run("Records", testTXTRegistryRecords)
	t.Run("RecordsError", testTXTRegistryRecordsReturnsErrors)
	t.Run("ApplyChanges", testTXTRegistryApplyChanges)
	t.Run("Labels", testTXTRegistryLabels)
}

func testTXTRegistryImplementsRegistry(t *testing.T) {
	mockProvider := new(provider.MockProvider)
	reg, _ := NewTXTRegistry(mockProvider, "_")
	assert.Implements(t, (*Registry)(nil), reg)
}

func testTXTRegistryRecords(t *testing.T) {
	for _, tc := range []struct {
		msg      string
		records  []*endpoint.Endpoint
		expected []*endpoint.Endpoint
	}{
		{
			msg: "no owner",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
			},
		},

		{
			msg: "with owner",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
			},
		},

		{
			msg: "with wrong owner",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "owner-for-something-else.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
			},
		},

		{
			msg: "spaces don't matter",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,  external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
			},
		},

		{
			msg: "heritage can be anywhere",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "external-dns/owner=foo,heritage=external-dns", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
			},
		},

		{
			msg: "support arbitrary label name",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
			},
		},

		{
			msg: "require heritage prefix",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
			},
		},

		{
			msg: "look for heritage",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
			},
		},

		{
			msg: "respect heritage of others",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=mate,external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
			},
		},

		{
			msg: "support multiple labels",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
			},
		},

		{
			msg: "label order doesn't matter",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/owner=foo,external-dns/my-label=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
			},
		},

		{
			msg: "heritage and label order doesn't matter",
			records: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A"},
				{DNSName: "foo.example.org", Target: "external-dns/owner=foo,heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
			},
		},
	} {
		t.Run(tc.msg, func(t *testing.T) {
			mockProvider := new(provider.MockProvider)
			mockProvider.On("Records").Return(tc.records, nil)

			reg, _ := NewTXTRegistry(mockProvider, "_")
			records, err := reg.Records()
			assert.NoError(t, err)

			assert.True(t, testutils.SameEndpoints(records, tc.expected))

			mockProvider.AssertExpectations(t)
		})
	}
}

func testTXTRegistryRecordsReturnsErrors(t *testing.T) {
	mockProvider := new(provider.MockProvider)
	mockProvider.On("Records").Return(nil, errors.New("some error"))

	reg, _ := NewTXTRegistry(mockProvider, "_")
	_, err := reg.Records()
	assert.EqualError(t, err, "some error")

	mockProvider.AssertExpectations(t)
}

func testTXTRegistryApplyChanges(t *testing.T) {
	for _, tc := range []struct {
		msg      string
		given    *plan.Changes
		expected *plan.Changes
	}{
		{
			msg: "no owner",
			given: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{}},
				},
			},
			expected: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{}},
				},
			},
		},

		{
			msg: "with owner",
			given: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
				},
			},
			expected: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
					{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo"}},
					{DNSName: "bar.example.org", Target: "heritage=external-dns,external-dns/owner=foo", RecordType: "TXT"},
				},
			},
		},

		{
			msg: "support arbitrary labels",
			given: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
				},
			},
			expected: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
					{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"my-label": "foo"}},
					{DNSName: "bar.example.org", Target: "heritage=external-dns,external-dns/my-label=foo", RecordType: "TXT"},
				},
			},
		},

		{
			msg: "support multiple labels",
			given: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
			},
			expected: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "bar.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
			},
		},

		{
			msg: "multiple labels order shouldn't matter",
			given: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
				},
			},
			expected: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{DNSName: "foo.example.org", Target: "8.8.8.8", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "foo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateOld: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "1.2.3.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				UpdateNew: []*endpoint.Endpoint{
					{DNSName: "wambo.example.org", Target: "4.3.2.1", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "wambo.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
				Delete: []*endpoint.Endpoint{
					{DNSName: "bar.example.org", Target: "8.8.4.4", RecordType: "A", Labels: map[string]string{"owner": "foo", "my-label": "foo"}},
					{DNSName: "bar.example.org", Target: "heritage=external-dns,external-dns/my-label=foo,external-dns/owner=foo", RecordType: "TXT"},
				},
			},
		},
	} {
		t.Run(tc.msg, func(t *testing.T) {
			mockProvider := new(provider.MockProvider)
			mockProvider.On("ApplyChanges", tc.expected).Return(nil)

			reg, _ := NewTXTRegistry(mockProvider, "_")
			err := reg.ApplyChanges(tc.given)
			assert.NoError(t, err)

			mockProvider.AssertExpectations(t)
		})
	}
}

func testTXTRegistryLabels(t *testing.T) {
	for _, tc := range []struct {
		msg      string
		labelStr string
		expected map[string]string
		parsed   string
	}{
		{
			msg:      "",
			labelStr: "foo=bar,qux=wambo",
			expected: map[string]string{
				"foo": "bar",
				"qux": "wambo",
			},
			parsed: "foo=bar,qux=wambo",
		},
		{
			msg:      "",
			labelStr: "foo",
			expected: map[string]string{},
			parsed:   "<none>",
		},
		{
			msg:      "",
			labelStr: "",
			expected: map[string]string{},
			parsed:   "<none>",
		},
		{
			msg:      "",
			labelStr: "foo=,wambo",
			expected: map[string]string{
				"foo": "",
			},
			parsed: "foo=",
		},
		{
			msg:      "",
			labelStr: "foo=bar,wambo",
			expected: map[string]string{
				"foo": "bar",
			},
			parsed: "foo=bar",
		},
		{
			msg:      "",
			labelStr: "foo=bar,wambo=",
			expected: map[string]string{
				"foo":   "bar",
				"wambo": "",
			},
			parsed: "foo=bar,wambo=",
		},
	} {
		t.Run(tc.msg, func(t *testing.T) {
			assert.Equal(t, tc.expected, parseLabels(tc.labelStr))
			assert.Equal(t, tc.parsed, formatLabels(tc.expected))
		})
	}
}
