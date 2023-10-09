package glesys

import (
	"fmt"
	"github.com/glesys/glesys-go/v7"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"sigs.k8s.io/external-dns/endpoint"
	"sort"
	"testing"
)

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

func TestMakeRecord(t *testing.T) {
	recordid := 0
	domain := "example.com"
	host := "foo"
	target := "127.0.0.1"
	recordType := "A"
	ttl := 3600
	actual := makeUpdateRecordParams(recordid, domain, host, target, recordType, ttl)
	expected := glesys.UpdateRecordParams{
		RecordID: recordid,
		Host:     host,
		Data:     target,
		Type:     recordType,
		TTL:      3600,
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateRecord(t *testing.T) {
	recordsByDomain := map[string][]glesys.DNSDomainRecord{
		"example.com": {
			{
				DomainName: "example.com",
				RecordID:   1,
				Host:       "foo",
				Type:       endpoint.RecordTypeA,
				Data:       "1.2.3.4",
				TTL:        glesysRecordTTL,
			},
			{
				DomainName: "example.com",
				RecordID:   2,
				Host:       "foo",
				Type:       endpoint.RecordTypeA,
				Data:       "5.6.7.8",
				TTL:        glesysRecordTTL,
			},
			{
				DomainName: "example.com",
				RecordID:   3,
				Host:       "@",
				Type:       endpoint.RecordTypeCNAME,
				Data:       "foo.example.com.",
				TTL:        glesysRecordTTL,
			},
		},
	}
	updatesByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "10.11.12.13"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "bar.example.com"),
		},
	}
	var changes glesysChanges
	err := processUpdateActions(recordsByDomain, updatesByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 2, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 3, len(changes.Deletes))
	expectedCreates := []*glesys.AddRecordParams{
		{
			DomainName: "example.com",
			Host:       "foo",
			Type:       endpoint.RecordTypeA,
			Data:       "10.11.12.13",
			TTL:        glesysRecordTTL,
		},
		{
			DomainName: "example.com",
			Host:       "@",
			Type:       endpoint.RecordTypeCNAME,
			Data:       "bar.example.com",
			TTL:        glesysRecordTTL,
		},
	}

	if !elementsMatch(t, expectedCreates, changes.Creates) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedCreates, changes.Creates))
	}

	expectedDeletes := []*glesys.DNSDomainRecord{
		{
			RecordID:   2,
			DomainName: "example.com",
			Host:       "foo",
			Type:       endpoint.RecordTypeA,
			Data:       "5.6.7.8",
			TTL:        glesysRecordTTL,
		},
		{
			RecordID:   2,
			DomainName: "example.com",
			Host:       "foo",
			Type:       endpoint.RecordTypeA,
			Data:       "5.6.7.8",
			TTL:        glesysRecordTTL,
		},
		{
			RecordID:   3,
			DomainName: "example.com",
			Host:       "@",
			Type:       endpoint.RecordTypeCNAME,
			Data:       "foo.example.com.",
			TTL:        glesysRecordTTL,
		},
	}
	sort.SliceStable(expectedDeletes, func(i, j int) bool {
		return expectedDeletes[i].RecordID < expectedDeletes[j].RecordID
	})
	sort.SliceStable(changes.Deletes, func(i, j int) bool {
		return changes.Deletes[i].RecordID < changes.Deletes[j].RecordID
	})
	if !elementsMatch(t, expectedDeletes, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDeletes, changes.Deletes))
	}
}

func TestGlesysProcessDeleteActions(t *testing.T) {
	recordsByDomain := map[string][]glesys.DNSDomainRecord{
		"example.com": {
			{
				DomainName: "example.com",
				RecordID:   1,
				Host:       "foo",
				Type:       endpoint.RecordTypeA,
				Data:       "1.2.3.4",
				TTL:        glesysRecordTTL,
			},
			{
				DomainName: "example.com",
				RecordID:   2,
				Host:       "foo",
				Type:       endpoint.RecordTypeA,
				Data:       "5.6.7.8",
				TTL:        glesysRecordTTL,
			},
			{
				DomainName: "example.com",
				RecordID:   3,
				Host:       "@",
				Type:       endpoint.RecordTypeCNAME,
				Data:       "foo.example.com.",
				TTL:        glesysRecordTTL,
			},
		},
	}

	deletesByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "foo.example.com"),
		},
	}

	var changes glesysChanges
	err := processDeleteActions(recordsByDomain, deletesByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 0, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 2, len(changes.Deletes))

	expectedDeletes := []int{1, 2, 3}

	if !elementsMatch(t, expectedDeletes, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDeletes, changes.Deletes))
	}
}
