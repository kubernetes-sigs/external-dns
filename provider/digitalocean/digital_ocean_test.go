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

package digitalocean

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/digitalocean/godo"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockDigitalOceanClient struct{}

func (m *mockDigitalOceanClient) RecordsByName(context.Context, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanClient) RecordsByTypeAndName(context.Context, string, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanClient) RecordsByType(context.Context, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanClient) List(ctx context.Context, opt *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	if opt == nil || opt.Page == 0 {
		return []godo.Domain{{Name: "foo.com"}, {Name: "example.com"}}, &godo.Response{
			Links: &godo.Links{
				Pages: &godo.Pages{
					Next: "http://example.com/v2/domains/?page=2",
					Last: "1234",
				},
			},
		}, nil
	}
	return []godo.Domain{{Name: "bar.com"}, {Name: "bar.de"}}, nil, nil
}

func (m *mockDigitalOceanClient) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanClient) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1, Name: "new", Type: "CNAME"}, nil, nil
}

func (m *mockDigitalOceanClient) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanClient) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanClient) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanClient) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanClient) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanClient) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	switch domain {
	case "foo.com":
		if opt == nil || opt.Page == 0 {
			return []godo.DomainRecord{
					{ID: 1, Name: "foo.ext-dns-test", Type: "CNAME"},
					{ID: 2, Name: "bar.ext-dns-test", Type: "CNAME"},
					{ID: 3, Name: "@", Type: endpoint.RecordTypeCNAME},
				}, &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{
							Next: "http://example.com/v2/domains/?page=2",
							Last: "1234",
						},
					},
				}, nil
		}
		return []godo.DomainRecord{{ID: 3, Name: "baz.ext-dns-test", Type: "A"}}, nil, nil
	case "example.com":
		if opt == nil || opt.Page == 0 {
			return []godo.DomainRecord{{ID: 1, Name: "new", Type: "CNAME"}}, &godo.Response{
				Links: &godo.Links{
					Pages: &godo.Pages{
						Next: "http://example.com/v2/domains/?page=2",
						Last: "1234",
					},
				},
			}, nil
		}
		return nil, nil, nil
	default:
		return nil, nil, nil
	}
}

type mockDigitalOceanRecordsFail struct{}

func (m *mockDigitalOceanRecordsFail) RecordsByName(context.Context, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanRecordsFail) RecordsByTypeAndName(context.Context, string, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanRecordsFail) RecordsByType(context.Context, string, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	// not used, here only to correctly implement the interface
	return nil, nil, nil
}

func (m *mockDigitalOceanRecordsFail) List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	return []godo.Domain{{Name: "foo.com"}, {Name: "bar.com"}}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Create(context.Context, *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) CreateRecord(context.Context, string, *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Delete(context.Context, string) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanRecordsFail) DeleteRecord(ctx context.Context, domain string, id int) (*godo.Response, error) {
	return nil, nil
}
func (m *mockDigitalOceanRecordsFail) EditRecord(ctx context.Context, domain string, id int, editRequest *godo.DomainRecordEditRequest) (*godo.DomainRecord, *godo.Response, error) {
	return &godo.DomainRecord{ID: 1}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Get(ctx context.Context, name string) (*godo.Domain, *godo.Response, error) {
	return &godo.Domain{Name: "example.com"}, nil, nil
}

func (m *mockDigitalOceanRecordsFail) Record(ctx context.Context, domain string, id int) (*godo.DomainRecord, *godo.Response, error) {
	return nil, nil, fmt.Errorf("Failed to get records")
}

func (m *mockDigitalOceanRecordsFail) Records(ctx context.Context, domain string, opt *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	return []godo.DomainRecord{}, nil, fmt.Errorf("Failed to get records")
}

func isEmpty(xs interface{}) bool {
	if xs != nil {
		objValue := reflect.ValueOf(xs)
		return objValue.Len() == 0
	}
	return true
}

// This function is an adapted copy of the testify package's ElementsMatch function with the
// call to ObjectsAreEqual replaced with cmp.Equal which better handles struct's with pointers to
// other structs. It also ignores ordering when comparing unlike cmp.Equal.
func elementsMatch(t *testing.T, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	if listA == nil && listB == nil {
		return true
	} else if listA == nil {
		return isEmpty(listB)
	} else if listB == nil {
		return isEmpty(listA)
	}

	aKind := reflect.TypeOf(listA).Kind()
	bKind := reflect.TypeOf(listB).Kind()

	if aKind != reflect.Array && aKind != reflect.Slice {
		return assert.Fail(t, fmt.Sprintf("%q has an unsupported type %s", listA, aKind), msgAndArgs...)
	}

	if bKind != reflect.Array && bKind != reflect.Slice {
		return assert.Fail(t, fmt.Sprintf("%q has an unsupported type %s", listB, bKind), msgAndArgs...)
	}

	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)

	aLen := aValue.Len()
	bLen := bValue.Len()

	if aLen != bLen {
		return assert.Fail(t, fmt.Sprintf("lengths don't match: %d != %d", aLen, bLen), msgAndArgs...)
	}

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if cmp.Equal(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return assert.Fail(t, fmt.Sprintf("element %s appears more times in %s than in %s", element, aValue, bValue), msgAndArgs...)
		}
	}

	return true
}

// Test adapted from test in testify library.
// https://github.com/stretchr/testify/blob/b8f7d52a4a7c581d5ed42333572e7fb857c687c2/assert/assertions_test.go#L768-L796
func TestElementsMatch(t *testing.T) {
	mockT := new(testing.T)

	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
	}{
		// matching
		{nil, nil, true},

		{nil, nil, true},
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 1}, []int{1, 1}, true},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{2, 1}, true},
		{[2]int{1, 2}, [2]int{2, 1}, true},
		{[]string{"hello", "world"}, []string{"world", "hello"}, true},
		{[]string{"hello", "hello"}, []string{"hello", "hello"}, true},
		{[]string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}, true},
		{[3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}, true},
		{[]int{}, nil, true},

		// not matching
		{[]int{1}, []int{1, 1}, false},
		{[]int{1, 2}, []int{2, 2}, false},
		{[]string{"hello", "hello"}, []string{"hello"}, false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ElementsMatch(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := elementsMatch(mockT, c.actual, c.expected)

			if res != c.result {
				t.Errorf("elementsMatch(%#v, %#v) should return %v", c.actual, c.expected, c.result)
			}
		})
	}
}

func TestDigitalOceanZones(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client:       &mockDigitalOceanClient{},
		domainFilter: endpoint.NewDomainFilter([]string{"com"}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	validateDigitalOceanZones(t, zones, []godo.Domain{
		{Name: "foo.com"}, {Name: "example.com"}, {Name: "bar.com"},
	})
}

func TestDigitalOceanMakeDomainEditRequest(t *testing.T) {
	// Ensure that records at the root of the zone get `@` as the name.
	r1 := makeDomainEditRequest("example.com", "example.com", endpoint.RecordTypeA,
		"1.2.3.4", digitalOceanRecordTTL)
	assert.Equal(t, &godo.DomainRecordEditRequest{
		Type: endpoint.RecordTypeA,
		Name: "@",
		Data: "1.2.3.4",
		TTL:  digitalOceanRecordTTL,
	}, r1)

	// Ensure the CNAME records have a `.` appended.
	r2 := makeDomainEditRequest("example.com", "foo.example.com", endpoint.RecordTypeCNAME,
		"bar.example.com", digitalOceanRecordTTL)
	assert.Equal(t, &godo.DomainRecordEditRequest{
		Type: endpoint.RecordTypeCNAME,
		Name: "foo",
		Data: "bar.example.com.",
		TTL:  digitalOceanRecordTTL,
	}, r2)

	// Ensure that CNAME records do not have an extra `.` appended if they already have a `.`
	r3 := makeDomainEditRequest("example.com", "foo.example.com", endpoint.RecordTypeCNAME,
		"bar.example.com.", digitalOceanRecordTTL)
	assert.Equal(t, &godo.DomainRecordEditRequest{
		Type: endpoint.RecordTypeCNAME,
		Name: "foo",
		Data: "bar.example.com.",
		TTL:  digitalOceanRecordTTL,
	}, r3)
}

func TestDigitalOceanApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.ext-dns-test.bar.com", Targets: endpoint.Targets{"target"}},
		{DNSName: "new.ext-dns-test-with-ttl.bar.com", Targets: endpoint.Targets{"target"}, RecordTTL: 100},
		{DNSName: "new.ext-dns-test.unexpected.com", Targets: endpoint.Targets{"target"}},
		{DNSName: "bar.com", Targets: endpoint.Targets{"target"}},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.bar.com", Targets: endpoint.Targets{"target"}}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.bar.de", Targets: endpoint.Targets{"target-old"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.foo.com", Targets: endpoint.Targets{"target-new"}, RecordType: "CNAME", RecordTTL: 100}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestDigitalOceanProcessCreateActions(t *testing.T) {
	recordsByDomain := map[string][]godo.DomainRecord{
		"example.com": nil,
	}

	createsByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "foo.example.com"),
		},
	}

	var changes digitalOceanChanges
	err := processCreateActions(recordsByDomain, createsByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 2, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 0, len(changes.Deletes))

	expectedCreates := []*digitalOceanChangeCreate{
		{
			Domain: "example.com",
			Options: &godo.DomainRecordEditRequest{
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "1.2.3.4",
				TTL:  digitalOceanRecordTTL,
			},
		},
		{
			Domain: "example.com",
			Options: &godo.DomainRecordEditRequest{
				Name: "@",
				Type: endpoint.RecordTypeCNAME,
				Data: "foo.example.com.",
				TTL:  digitalOceanRecordTTL,
			},
		},
	}

	if !elementsMatch(t, expectedCreates, changes.Creates) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedCreates, changes.Creates))
	}
}

func TestDigitalOceanProcessUpdateActions(t *testing.T) {
	recordsByDomain := map[string][]godo.DomainRecord{
		"example.com": {
			{
				ID:   1,
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "1.2.3.4",
				TTL:  digitalOceanRecordTTL,
			},
			{
				ID:   2,
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "5.6.7.8",
				TTL:  digitalOceanRecordTTL,
			},
			{
				ID:   3,
				Name: "@",
				Type: endpoint.RecordTypeCNAME,
				Data: "foo.example.com.",
				TTL:  digitalOceanRecordTTL,
			},
		},
	}

	updatesByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "10.11.12.13"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "bar.example.com"),
		},
	}

	var changes digitalOceanChanges
	err := processUpdateActions(recordsByDomain, updatesByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 2, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 3, len(changes.Deletes))

	expectedCreates := []*digitalOceanChangeCreate{
		{
			Domain: "example.com",
			Options: &godo.DomainRecordEditRequest{
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "10.11.12.13",
				TTL:  digitalOceanRecordTTL,
			},
		},
		{
			Domain: "example.com",
			Options: &godo.DomainRecordEditRequest{
				Name: "@",
				Type: endpoint.RecordTypeCNAME,
				Data: "bar.example.com.",
				TTL:  digitalOceanRecordTTL,
			},
		},
	}

	if !elementsMatch(t, expectedCreates, changes.Creates) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedCreates, changes.Creates))
	}

	expectedDeletes := []*digitalOceanChangeDelete{
		{
			Domain:   "example.com",
			RecordID: 1,
		},
		{
			Domain:   "example.com",
			RecordID: 2,
		},
		{
			Domain:   "example.com",
			RecordID: 3,
		},
	}

	if !elementsMatch(t, expectedDeletes, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDeletes, changes.Deletes))
	}
}

func TestDigitalOceanProcessDeleteActions(t *testing.T) {
	recordsByDomain := map[string][]godo.DomainRecord{
		"example.com": {
			{
				ID:   1,
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "1.2.3.4",
				TTL:  digitalOceanRecordTTL,
			},
			// This record will not be deleted because it represents a target not specified to be deleted.
			{
				ID:   2,
				Name: "foo",
				Type: endpoint.RecordTypeA,
				Data: "5.6.7.8",
				TTL:  digitalOceanRecordTTL,
			},
			{
				ID:   3,
				Name: "@",
				Type: endpoint.RecordTypeCNAME,
				Data: "foo.example.com.",
				TTL:  digitalOceanRecordTTL,
			},
		},
	}

	deletesByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "foo.example.com"),
		},
	}

	var changes digitalOceanChanges
	err := processDeleteActions(recordsByDomain, deletesByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 0, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 2, len(changes.Deletes))

	expectedDeletes := []*digitalOceanChangeDelete{
		{
			Domain:   "example.com",
			RecordID: 1,
		},
		{
			Domain:   "example.com",
			RecordID: 3,
		},
	}

	if !elementsMatch(t, expectedDeletes, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDeletes, changes.Deletes))
	}
}

func TestNewDigitalOceanProvider(t *testing.T) {
	_ = os.Setenv("DO_TOKEN", "xxxxxxxxxxxxxxxxx")
	_, err := NewDigitalOceanProvider(context.Background(), endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true, 50)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("DO_TOKEN")
	_, err = NewDigitalOceanProvider(context.Background(), endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true, 50)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestDigitalOceanGetMatchingDomainRecords(t *testing.T) {
	records := []godo.DomainRecord{
		{
			ID:   1,
			Name: "foo",
			Type: endpoint.RecordTypeCNAME,
			Data: "baz.org.",
		},
		{
			ID:   2,
			Name: "baz",
			Type: endpoint.RecordTypeA,
			Data: "1.2.3.4",
		},
		{
			ID:   3,
			Name: "baz",
			Type: endpoint.RecordTypeA,
			Data: "5.6.7.8",
		},
		{
			ID:   4,
			Name: "@",
			Type: endpoint.RecordTypeA,
			Data: "9.10.11.12",
		},
	}

	ep1 := endpoint.NewEndpoint("foo.com", endpoint.RecordTypeCNAME)
	assert.Equal(t, 1, len(getMatchingDomainRecords(records, "com", ep1)))

	ep2 := endpoint.NewEndpoint("foo.com", endpoint.RecordTypeA)
	assert.Equal(t, 0, len(getMatchingDomainRecords(records, "com", ep2)))

	ep3 := endpoint.NewEndpoint("baz.org", endpoint.RecordTypeA)
	r := getMatchingDomainRecords(records, "org", ep3)
	assert.Equal(t, 2, len(r))
	assert.ElementsMatch(t, r, []godo.DomainRecord{
		{
			ID:   2,
			Name: "baz",
			Type: endpoint.RecordTypeA,
			Data: "1.2.3.4",
		},
		{
			ID:   3,
			Name: "baz",
			Type: endpoint.RecordTypeA,
			Data: "5.6.7.8",
		},
	})

	ep4 := endpoint.NewEndpoint("example.com", endpoint.RecordTypeA)
	r2 := getMatchingDomainRecords(records, "example.com", ep4)
	assert.Equal(t, 1, len(r2))
	assert.Equal(t, "9.10.11.12", r2[0].Data)
}

func validateDigitalOceanZones(t *testing.T, zones []godo.Domain, expected []godo.Domain) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		assert.Equal(t, expected[i].Name, zone.Name)
	}
}

func TestDigitalOceanRecord(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}

	records, err := provider.fetchRecords(context.Background(), "example.com")
	if err != nil {
		t.Fatal(err)
	}
	expected := []godo.DomainRecord{{ID: 1, Name: "new", Type: "CNAME"}}
	require.Len(t, records, len(expected))
	for i, record := range records {
		assert.Equal(t, expected[i].Name, record.Name)
	}
}

func TestDigitalOceanAllRecords(t *testing.T) {
	provider := &DigitalOceanProvider{
		Client: &mockDigitalOceanClient{},
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	require.Equal(t, 5, len(records))

	provider.Client = &mockDigitalOceanRecordsFail{}
	_, err = provider.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestDigitalOceanMergeRecordsByNameType(t *testing.T) {
	xs := []*endpoint.Endpoint{
		endpoint.NewEndpoint("foo.example.com", "A", "1.2.3.4"),
		endpoint.NewEndpoint("bar.example.com", "A", "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", "A", "5.6.7.8"),
		endpoint.NewEndpoint("foo.example.com", "CNAME", "somewhere.out.there.com"),
	}

	merged := mergeEndpointsByNameType(xs)

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
