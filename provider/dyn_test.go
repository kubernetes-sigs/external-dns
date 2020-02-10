/*
Copyright 2018 The Kubernetes Authors.

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

package provider

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nesv/go-dynect/dynect"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestDynMerge_NoUpdateOnTTL0Changes(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeA,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	assert.Equal(t, 0, len(merge(updateOld, updateNew)))
}

func TestDynMerge_UpdateOnTTLChanges(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(77),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(10),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 2, len(merged))
	assert.Equal(t, "name1", merged[0].DNSName)
}

func TestDynMerge_AlwaysUpdateTarget(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1-changed"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 1, len(merged))
	assert.Equal(t, "target1-changed", merged[0].Targets[0])
}

func TestDynMerge_NoUpdateIfTTLUnchanged(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 0, len(merged))
}

func TestDyn_endpointToRecord(t *testing.T) {
	tests := []struct {
		ep        *endpoint.Endpoint
		extractor func(*dynect.DataBlock) string
	}{
		{endpoint.NewEndpoint("address", "A", "the-target"), func(b *dynect.DataBlock) string { return b.Address }},
		{endpoint.NewEndpoint("cname", "CNAME", "the-target"), func(b *dynect.DataBlock) string { return b.CName }},
		{endpoint.NewEndpoint("text", "TXT", "the-target"), func(b *dynect.DataBlock) string { return b.TxtData }},
	}

	for _, tc := range tests {
		block := endpointToRecord(tc.ep)
		assert.Equal(t, "the-target", tc.extractor(block))
	}
}

func TestDyn_buildLinkToRecord(t *testing.T) {
	provider := &dynProviderState{
		DynConfig: DynConfig{
			ZoneIDFilter: NewZoneIDFilter([]string{"example.com"}),
			DomainFilter: endpoint.NewDomainFilter([]string{"the-target.example.com"}),
		},
	}

	tests := []struct {
		ep   *endpoint.Endpoint
		link string
	}{
		{endpoint.NewEndpoint("sub.the-target.example.com", "A", "address"), "ARecord/example.com/sub.the-target.example.com/"},
		{endpoint.NewEndpoint("the-target.example.com", "CNAME", "cname"), "CNAMERecord/example.com/the-target.example.com/"},
		{endpoint.NewEndpoint("the-target.example.com", "TXT", "text"), "TXTRecord/example.com/the-target.example.com/"},
		{endpoint.NewEndpoint("the-target.google.com", "TXT", "text"), ""},
		{endpoint.NewEndpoint("mail.example.com", "TXT", "text"), ""},
		{nil, ""},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.link, provider.buildLinkToRecord(tc.ep))
	}
}

func TestDyn_errorOrValue(t *testing.T) {
	e := errors.New("an error")
	val := "value"
	assert.Equal(t, e, errorOrValue(e, val))
	assert.Equal(t, val, errorOrValue(nil, val))
}

func TestDyn_filterAndFixLinks(t *testing.T) {
	links := []string{
		"/REST/ARecord/example.com/the-target.example.com/",
		"/REST/ARecord/example.com/the-target.google.com/",
		"/REST/TXTRecord/example.com/the-target.example.com/",
		"/REST/TXTRecord/example.com/the-target.google.com/",
		"/REST/CNAMERecord/example.com/the-target.google.com/",
		"/REST/CNAMERecord/example.com/the-target.example.com/",
		"/REST/NSRecord/example.com/the-target.google.com/",
		"/REST/NSRecord/example.com/the-target.example.com/",
	}
	filter := endpoint.NewDomainFilter([]string{"example.com"})
	result := filterAndFixLinks(links, filter)

	// should skip non-example.com records and NS records too
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "ARecord/example.com/the-target.example.com/", result[0])
	assert.Equal(t, "TXTRecord/example.com/the-target.example.com/", result[1])
	assert.Equal(t, "CNAMERecord/example.com/the-target.example.com/", result[2])
}

func TestDyn_fixMissingTTL(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%v", dynDefaultTTL), fixMissingTTL(endpoint.TTL(0), 0))

	// nothing to fix
	assert.Equal(t, "111", fixMissingTTL(endpoint.TTL(111), 25))

	// apply min TTL
	assert.Equal(t, "1992", fixMissingTTL(endpoint.TTL(111), 1992))
}

func TestDyn_Snapshot(t *testing.T) {
	snap := ZoneSnapshot{
		serials:   map[string]int{},
		endpoints: map[string][]*endpoint.Endpoint{},
	}

	recs := []*endpoint.Endpoint{
		{
			DNSName:    "name",
			Targets:    endpoint.Targets{"target"},
			RecordTTL:  endpoint.TTL(10000),
			RecordType: "A",
		},
	}

	snap.StoreRecordsForSerial("test", 12, recs)

	cached := snap.GetRecordsForSerial("test", 12)
	assert.Equal(t, recs, cached)

	cached = snap.GetRecordsForSerial("test", 999)
	assert.Nil(t, cached)

	cached = snap.GetRecordsForSerial("sfas", 12)
	assert.Nil(t, cached)

	recs2 := []*endpoint.Endpoint{
		{
			DNSName:    "name",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(100),
			RecordType: "CNAME",
		},
	}

	// update zone with different records and newer serial
	snap.StoreRecordsForSerial("test", 13, recs2)

	cached = snap.GetRecordsForSerial("test", 13)
	assert.Equal(t, recs2, cached)

	cached = snap.GetRecordsForSerial("test", 12)
	assert.Nil(t, cached)
}
