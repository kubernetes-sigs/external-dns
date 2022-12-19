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

package infoblox

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type mockIBConnector struct {
	mockInfobloxZones   *[]ibclient.ZoneAuth
	mockInfobloxObjects *[]ibclient.IBObject
	createdEndpoints    []*endpoint.Endpoint
	deletedEndpoints    []*endpoint.Endpoint
	updatedEndpoints    []*endpoint.Endpoint
}

func (client *mockIBConnector) CreateObject(obj ibclient.IBObject) (ref string, err error) {
	switch obj.ObjectType() {
	case "record:a":
		client.createdEndpoints = append(
			client.createdEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordA).Name,
				endpoint.RecordTypeA,
				obj.(*ibclient.RecordA).Ipv4Addr,
			),
		)
		ref = fmt.Sprintf("%s/%s:%s/default", obj.ObjectType(), base64.StdEncoding.EncodeToString([]byte(obj.(*ibclient.RecordA).Name)), obj.(*ibclient.RecordA).Name)
		obj.(*ibclient.RecordA).Ref = ref
	case "record:cname":
		client.createdEndpoints = append(
			client.createdEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordCNAME).Name,
				endpoint.RecordTypeCNAME,
				obj.(*ibclient.RecordCNAME).Canonical,
			),
		)
		ref = fmt.Sprintf("%s/%s:%s/default", obj.ObjectType(), base64.StdEncoding.EncodeToString([]byte(obj.(*ibclient.RecordCNAME).Name)), obj.(*ibclient.RecordCNAME).Name)
		obj.(*ibclient.RecordCNAME).Ref = ref
	case "record:host":
		for _, i := range obj.(*ibclient.HostRecord).Ipv4Addrs {
			client.createdEndpoints = append(
				client.createdEndpoints,
				endpoint.NewEndpoint(
					obj.(*ibclient.HostRecord).Name,
					endpoint.RecordTypeA,
					i.Ipv4Addr,
				),
			)
		}
		ref = fmt.Sprintf("%s/%s:%s/default", obj.ObjectType(), base64.StdEncoding.EncodeToString([]byte(obj.(*ibclient.HostRecord).Name)), obj.(*ibclient.HostRecord).Name)
		obj.(*ibclient.HostRecord).Ref = ref
	case "record:txt":
		client.createdEndpoints = append(
			client.createdEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordTXT).Name,
				endpoint.RecordTypeTXT,
				obj.(*ibclient.RecordTXT).Text,
			),
		)
		obj.(*ibclient.RecordTXT).Ref = ref
		ref = fmt.Sprintf("%s/%s:%s/default", obj.ObjectType(), base64.StdEncoding.EncodeToString([]byte(obj.(*ibclient.RecordTXT).Name)), obj.(*ibclient.RecordTXT).Name)
	case "record:ptr":
		client.createdEndpoints = append(
			client.createdEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordPTR).PtrdName,
				endpoint.RecordTypePTR,
				obj.(*ibclient.RecordPTR).Ipv4Addr,
			),
		)
		obj.(*ibclient.RecordPTR).Ref = ref
		reverseAddr, err := dns.ReverseAddr(obj.(*ibclient.RecordPTR).Ipv4Addr)
		if err != nil {
			return ref, fmt.Errorf("unable to create reverse addr from %s", obj.(*ibclient.RecordPTR).Ipv4Addr)
		}
		ref = fmt.Sprintf("%s/%s:%s/default", obj.ObjectType(), base64.StdEncoding.EncodeToString([]byte(obj.(*ibclient.RecordPTR).PtrdName)), reverseAddr)
	}
	*client.mockInfobloxObjects = append(
		*client.mockInfobloxObjects,
		obj,
	)
	return ref, nil
}

func (client *mockIBConnector) GetObject(obj ibclient.IBObject, ref string, queryParams *ibclient.QueryParams, res interface{}) (err error) {
	switch obj.ObjectType() {
	case "record:a":
		var result []ibclient.RecordA
		for _, object := range *client.mockInfobloxObjects {
			if object.ObjectType() == "record:a" {
				if ref != "" &&
					ref != object.(*ibclient.RecordA).Ref {
					continue
				}
				if obj.(*ibclient.RecordA).Name != "" &&
					obj.(*ibclient.RecordA).Name != object.(*ibclient.RecordA).Name {
					continue
				}
				result = append(result, *object.(*ibclient.RecordA))
			}
		}
		*res.(*[]ibclient.RecordA) = result
	case "record:cname":
		var result []ibclient.RecordCNAME
		for _, object := range *client.mockInfobloxObjects {
			if object.ObjectType() == "record:cname" {
				if ref != "" &&
					ref != object.(*ibclient.RecordCNAME).Ref {
					continue
				}
				if obj.(*ibclient.RecordCNAME).Name != "" &&
					obj.(*ibclient.RecordCNAME).Name != object.(*ibclient.RecordCNAME).Name {
					continue
				}
				result = append(result, *object.(*ibclient.RecordCNAME))
			}
		}
		*res.(*[]ibclient.RecordCNAME) = result
	case "record:host":
		var result []ibclient.HostRecord
		for _, object := range *client.mockInfobloxObjects {
			if object.ObjectType() == "record:host" {
				if ref != "" &&
					ref != object.(*ibclient.HostRecord).Ref {
					continue
				}
				if obj.(*ibclient.HostRecord).Name != "" &&
					obj.(*ibclient.HostRecord).Name != object.(*ibclient.HostRecord).Name {
					continue
				}
				result = append(result, *object.(*ibclient.HostRecord))
			}
		}
		*res.(*[]ibclient.HostRecord) = result
	case "record:txt":
		var result []ibclient.RecordTXT
		for _, object := range *client.mockInfobloxObjects {
			if object.ObjectType() == "record:txt" {
				if ref != "" &&
					ref != object.(*ibclient.RecordTXT).Ref {
					continue
				}
				if obj.(*ibclient.RecordTXT).Name != "" &&
					obj.(*ibclient.RecordTXT).Name != object.(*ibclient.RecordTXT).Name {
					continue
				}
				result = append(result, *object.(*ibclient.RecordTXT))
			}
		}
		*res.(*[]ibclient.RecordTXT) = result
	case "record:ptr":
		var result []ibclient.RecordPTR
		for _, object := range *client.mockInfobloxObjects {
			if object.ObjectType() == "record:ptr" {
				if ref != "" &&
					ref != object.(*ibclient.RecordPTR).Ref {
					continue
				}
				if obj.(*ibclient.RecordPTR).PtrdName != "" &&
					obj.(*ibclient.RecordPTR).PtrdName != object.(*ibclient.RecordPTR).PtrdName {
					continue
				}
				result = append(result, *object.(*ibclient.RecordPTR))
			}
		}
		*res.(*[]ibclient.RecordPTR) = result
	case "zone_auth":
		*res.(*[]ibclient.ZoneAuth) = *client.mockInfobloxZones
	}
	return
}

func (client *mockIBConnector) DeleteObject(ref string) (refRes string, err error) {
	re := regexp.MustCompile(`([^/]+)/[^:]+:([^/]+)/default`)
	result := re.FindStringSubmatch(ref)

	switch result[1] {
	case "record:a":
		var records []ibclient.RecordA
		obj := ibclient.NewEmptyRecordA()
		obj.Name = result[2]
		client.GetObject(obj, ref, nil, &records)
		for _, record := range records {
			client.deletedEndpoints = append(
				client.deletedEndpoints,
				endpoint.NewEndpoint(
					record.Name,
					endpoint.RecordTypeA,
					"",
				),
			)
		}
	case "record:cname":
		var records []ibclient.RecordCNAME
		obj := ibclient.NewEmptyRecordCNAME()
		obj.Name = result[2]
		client.GetObject(obj, ref, nil, &records)
		for _, record := range records {
			client.deletedEndpoints = append(
				client.deletedEndpoints,
				endpoint.NewEndpoint(
					record.Name,
					endpoint.RecordTypeCNAME,
					"",
				),
			)
		}
	case "record:host":
		var records []ibclient.HostRecord
		obj := ibclient.NewEmptyHostRecord()
		obj.Name = result[2]
		client.GetObject(obj, ref, nil, &records)
		for _, record := range records {
			client.deletedEndpoints = append(
				client.deletedEndpoints,
				endpoint.NewEndpoint(
					record.Name,
					endpoint.RecordTypeA,
					"",
				),
			)
		}
	case "record:txt":
		var records []ibclient.RecordTXT
		obj := ibclient.NewRecordTXT(
			ibclient.RecordTXT{
				Name: result[2],
			},
		)
		client.GetObject(obj, ref, nil, &records)
		for _, record := range records {
			client.deletedEndpoints = append(
				client.deletedEndpoints,
				endpoint.NewEndpoint(
					record.Name,
					endpoint.RecordTypeTXT,
					"",
				),
			)
		}
	case "record:ptr":
		var records []ibclient.RecordPTR
		obj := ibclient.NewEmptyRecordPTR()
		obj.Name = result[2]
		client.GetObject(obj, ref, nil, &records)
		for _, record := range records {
			client.deletedEndpoints = append(
				client.deletedEndpoints,
				endpoint.NewEndpoint(
					record.PtrdName,
					endpoint.RecordTypePTR,
					"",
				),
			)
		}
	}
	return "", nil
}

func (client *mockIBConnector) UpdateObject(obj ibclient.IBObject, ref string) (refRes string, err error) {
	switch obj.ObjectType() {
	case "record:a":
		client.updatedEndpoints = append(
			client.updatedEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordA).Name,
				obj.(*ibclient.RecordA).Ipv4Addr,
				endpoint.RecordTypeA,
			),
		)
	case "record:cname":
		client.updatedEndpoints = append(
			client.updatedEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordCNAME).Name,
				obj.(*ibclient.RecordCNAME).Canonical,
				endpoint.RecordTypeCNAME,
			),
		)
	case "record:host":
		for _, i := range obj.(*ibclient.HostRecord).Ipv4Addrs {
			client.updatedEndpoints = append(
				client.updatedEndpoints,
				endpoint.NewEndpoint(
					obj.(*ibclient.HostRecord).Name,
					i.Ipv4Addr,
					endpoint.RecordTypeA,
				),
			)
		}
	case "record:txt":
		client.updatedEndpoints = append(
			client.updatedEndpoints,
			endpoint.NewEndpoint(
				obj.(*ibclient.RecordTXT).Name,
				obj.(*ibclient.RecordTXT).Text,
				endpoint.RecordTypeTXT,
			),
		)
	}
	return "", nil
}

func createMockInfobloxZone(fqdn string) ibclient.ZoneAuth {
	return ibclient.ZoneAuth{
		Fqdn: fqdn,
	}
}

func createMockInfobloxObject(name, recordType, value string) ibclient.IBObject {
	ref := fmt.Sprintf("record:%s/%s:%s/default", strings.ToLower(recordType), base64.StdEncoding.EncodeToString([]byte(name)), name)
	switch recordType {
	case endpoint.RecordTypeA:
		obj := ibclient.NewEmptyRecordA()
		obj.Name = name
		obj.Ref = ref
		obj.Ipv4Addr = value
		return obj
	case endpoint.RecordTypeCNAME:
		obj := ibclient.NewEmptyRecordCNAME()
		obj.Name = name
		obj.Ref = ref
		obj.Canonical = value
		return obj
	case endpoint.RecordTypeTXT:
		return ibclient.NewRecordTXT(
			ibclient.RecordTXT{
				Ref:  ref,
				Name: name,
				Text: value,
			},
		)
	case "HOST":
		obj := ibclient.NewEmptyHostRecord()
		obj.Name = name
		obj.Ref = ref
		obj.Ipv4Addrs = []ibclient.HostRecordIpv4Addr{
			{
				Ipv4Addr: value,
			},
		}
		return obj
	case endpoint.RecordTypePTR:
		obj := ibclient.NewEmptyRecordPTR()
		obj.PtrdName = name
		obj.Ref = ref
		obj.Ipv4Addr = value
		return obj
	}

	return nil
}

func newInfobloxProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, createPTR bool, client ibclient.IBConnector) *ProviderConfig {
	return &ProviderConfig{
		client:       client,
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		dryRun:       dryRun,
		createPTR:    createPTR,
	}
}

func TestInfobloxRecords(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxZones: &[]ibclient.ZoneAuth{
			createMockInfobloxZone("example.com"),
			createMockInfobloxZone("other.com"),
		},
		mockInfobloxObjects: &[]ibclient.IBObject{
			createMockInfobloxObject("example.com", endpoint.RecordTypeA, "123.123.123.122"),
			createMockInfobloxObject("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockInfobloxObject("nginx.example.com", endpoint.RecordTypeA, "123.123.123.123"),
			createMockInfobloxObject("nginx.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockInfobloxObject("whitespace.example.com", endpoint.RecordTypeA, "123.123.123.124"),
			createMockInfobloxObject("whitespace.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=white space"),
			createMockInfobloxObject("hack.example.com", endpoint.RecordTypeCNAME, "cerberus.infoblox.com"),
			createMockInfobloxObject("multiple.example.com", endpoint.RecordTypeA, "123.123.123.122"),
			createMockInfobloxObject("multiple.example.com", endpoint.RecordTypeA, "123.123.123.121"),
			createMockInfobloxObject("multiple.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockInfobloxObject("existing.example.com", endpoint.RecordTypeA, "124.1.1.1"),
			createMockInfobloxObject("existing.example.com", endpoint.RecordTypeA, "124.1.1.2"),
			createMockInfobloxObject("existing.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=existing"),
			createMockInfobloxObject("host.example.com", "HOST", "125.1.1.1"),
		},
	}

	providerCfg := newInfobloxProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, false, &client)
	actual, err := providerCfg.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=default\""),
		endpoint.NewEndpoint("nginx.example.com", endpoint.RecordTypeA, "123.123.123.123"),
		endpoint.NewEndpoint("nginx.example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=default\""),
		endpoint.NewEndpoint("whitespace.example.com", endpoint.RecordTypeA, "123.123.123.124"),
		endpoint.NewEndpoint("whitespace.example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=white space\""),
		endpoint.NewEndpoint("hack.example.com", endpoint.RecordTypeCNAME, "cerberus.infoblox.com"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeA, "123.123.123.122", "123.123.123.121"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=default\""),
		endpoint.NewEndpoint("existing.example.com", endpoint.RecordTypeA, "124.1.1.1", "124.1.1.2"),
		endpoint.NewEndpoint("existing.example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=existing\""),
		endpoint.NewEndpoint("host.example.com", endpoint.RecordTypeA, "125.1.1.1"),
	}
	validateEndpoints(t, actual, expected)
}

func TestInfobloxAdjustEndpoints(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxZones: &[]ibclient.ZoneAuth{
			createMockInfobloxZone("example.com"),
			createMockInfobloxZone("other.com"),
		},
		mockInfobloxObjects: &[]ibclient.IBObject{
			createMockInfobloxObject("example.com", endpoint.RecordTypeA, "123.123.123.122"),
			createMockInfobloxObject("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockInfobloxObject("hack.example.com", endpoint.RecordTypeCNAME, "cerberus.infoblox.com"),
			createMockInfobloxObject("host.example.com", "HOST", "125.1.1.1"),
		},
	}

	providerCfg := newInfobloxProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, true, &client)
	actual, err := providerCfg.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	providerCfg.AdjustEndpoints(actual)

	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122").WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=default\""),
		endpoint.NewEndpoint("hack.example.com", endpoint.RecordTypeCNAME, "cerberus.infoblox.com"),
		endpoint.NewEndpoint("host.example.com", endpoint.RecordTypeA, "125.1.1.1").WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true"),
	}
	validateEndpoints(t, actual, expected)
}

func TestInfobloxRecordsReverse(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxZones: &[]ibclient.ZoneAuth{
			createMockInfobloxZone("10.0.0.0/24"),
			createMockInfobloxZone("10.0.1.0/24"),
		},
		mockInfobloxObjects: &[]ibclient.IBObject{
			createMockInfobloxObject("example.com", endpoint.RecordTypePTR, "10.0.0.1"),
			createMockInfobloxObject("example2.com", endpoint.RecordTypePTR, "10.0.0.2"),
		},
	}

	providerCfg := newInfobloxProvider(endpoint.NewDomainFilter([]string{"10.0.0.0/24"}), provider.NewZoneIDFilter([]string{""}), true, true, &client)
	actual, err := providerCfg.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypePTR, "10.0.0.1"),
		endpoint.NewEndpoint("example2.com", endpoint.RecordTypePTR, "10.0.0.2"),
	}
	validateEndpoints(t, actual, expected)
}

func TestInfobloxApplyChanges(t *testing.T) {
	client := mockIBConnector{}

	testInfobloxApplyChangesInternal(t, false, false, &client)

	validateEndpoints(t, client.createdEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("new.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.NewEndpoint("newcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeA, "1.2.3.4,3.4.5.6,8.9.10.11"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeTXT, "tag-multiple-A-records"),
	})

	validateEndpoints(t, client.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, ""),
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateEndpoints(t, client.updatedEndpoints, []*endpoint.Endpoint{})
}

func TestInfobloxApplyChangesReverse(t *testing.T) {
	client := mockIBConnector{}

	testInfobloxApplyChangesInternal(t, false, true, &client)

	validateEndpoints(t, client.createdEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypePTR, "1.2.3.4"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypePTR, "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("new.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.NewEndpoint("newcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeA, "1.2.3.4,3.4.5.6,8.9.10.11"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeTXT, "tag-multiple-A-records"),
	})

	validateEndpoints(t, client.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, ""),
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypePTR, ""),
		endpoint.NewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateEndpoints(t, client.updatedEndpoints, []*endpoint.Endpoint{})
}

func TestInfobloxApplyChangesDryRun(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxObjects: &[]ibclient.IBObject{},
	}

	testInfobloxApplyChangesInternal(t, true, false, &client)

	validateEndpoints(t, client.createdEndpoints, []*endpoint.Endpoint{})

	validateEndpoints(t, client.deletedEndpoints, []*endpoint.Endpoint{})

	validateEndpoints(t, client.updatedEndpoints, []*endpoint.Endpoint{})
}

func testInfobloxApplyChangesInternal(t *testing.T, dryRun, createPTR bool, client ibclient.IBConnector) {
	client.(*mockIBConnector).mockInfobloxZones = &[]ibclient.ZoneAuth{
		createMockInfobloxZone("example.com"),
		createMockInfobloxZone("other.com"),
		createMockInfobloxZone("1.2.3.0/24"),
	}
	client.(*mockIBConnector).mockInfobloxObjects = &[]ibclient.IBObject{
		createMockInfobloxObject("deleted.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		createMockInfobloxObject("deleted.example.com", endpoint.RecordTypeTXT, "test-deleting-txt"),
		createMockInfobloxObject("deleted.example.com", endpoint.RecordTypePTR, "121.212.121.212"),
		createMockInfobloxObject("deletedcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		createMockInfobloxObject("old.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		createMockInfobloxObject("oldcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
	}

	providerCfg := newInfobloxProvider(
		endpoint.NewDomainFilter([]string{""}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		createPTR,
		client,
	)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("nope.com", endpoint.RecordTypeA, "4.4.4.4"),
		endpoint.NewEndpoint("nope.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeA, "1.2.3.4,3.4.5.6,8.9.10.11"),
		endpoint.NewEndpoint("multiple.example.com", endpoint.RecordTypeTXT, "tag-multiple-A-records"),
	}

	updateOldRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("old.nope.com", endpoint.RecordTypeA, "121.212.121.212"),
	}

	updateNewRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("new.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.NewEndpoint("newcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("new.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("deleted.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	if createPTR {
		deleteRecords = append(deleteRecords, endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypePTR, "121.212.121.212"))
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updateNewRecords,
		UpdateOld: updateOldRecords,
		Delete:    deleteRecords,
	}

	if err := providerCfg.ApplyChanges(context.Background(), changes); err != nil {
		t.Fatal(err)
	}
}

func TestInfobloxZones(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxZones: &[]ibclient.ZoneAuth{
			createMockInfobloxZone("example.com"),
			createMockInfobloxZone("lvl1-1.example.com"),
			createMockInfobloxZone("lvl2-1.lvl1-1.example.com"),
			createMockInfobloxZone("1.2.3.0/24"),
		},
		mockInfobloxObjects: &[]ibclient.IBObject{},
	}

	providerCfg := newInfobloxProvider(endpoint.NewDomainFilter([]string{"example.com", "1.2.3.0/24"}), provider.NewZoneIDFilter([]string{""}), true, false, &client)
	zones, _ := providerCfg.zones()
	var emptyZoneAuth *ibclient.ZoneAuth
	assert.Equal(t, providerCfg.findZone(zones, "example.com").Fqdn, "example.com")
	assert.Equal(t, providerCfg.findZone(zones, "nomatch-example.com"), emptyZoneAuth)
	assert.Equal(t, providerCfg.findZone(zones, "nginx.example.com").Fqdn, "example.com")
	assert.Equal(t, providerCfg.findZone(zones, "lvl1-1.example.com").Fqdn, "lvl1-1.example.com")
	assert.Equal(t, providerCfg.findZone(zones, "lvl1-2.example.com").Fqdn, "example.com")
	assert.Equal(t, providerCfg.findZone(zones, "lvl2-1.lvl1-1.example.com").Fqdn, "lvl2-1.lvl1-1.example.com")
	assert.Equal(t, providerCfg.findZone(zones, "lvl2-2.lvl1-1.example.com").Fqdn, "lvl1-1.example.com")
	assert.Equal(t, providerCfg.findZone(zones, "lvl2-2.lvl1-2.example.com").Fqdn, "example.com")
	assert.Equal(t, providerCfg.findZone(zones, "1.2.3.0/24").Fqdn, "1.2.3.0/24")
}

func TestInfobloxReverseZones(t *testing.T) {
	client := mockIBConnector{
		mockInfobloxZones: &[]ibclient.ZoneAuth{
			createMockInfobloxZone("example.com"),
			createMockInfobloxZone("1.2.3.0/24"),
			createMockInfobloxZone("10.0.0.0/8"),
		},
		mockInfobloxObjects: &[]ibclient.IBObject{},
	}

	providerCfg := newInfobloxProvider(endpoint.NewDomainFilter([]string{"example.com", "1.2.3.0/24", "10.0.0.0/8"}), provider.NewZoneIDFilter([]string{""}), true, false, &client)
	zones, _ := providerCfg.zones()
	var emptyZoneAuth *ibclient.ZoneAuth
	assert.Equal(t, providerCfg.findReverseZone(zones, "nomatch-example.com"), emptyZoneAuth)
	assert.Equal(t, providerCfg.findReverseZone(zones, "192.168.0.1"), emptyZoneAuth)
	assert.Equal(t, providerCfg.findReverseZone(zones, "1.2.3.4").Fqdn, "1.2.3.0/24")
	assert.Equal(t, providerCfg.findReverseZone(zones, "10.28.29.30").Fqdn, "10.0.0.0/8")
}

func TestExtendedRequestFDQDRegExBuilder(t *testing.T) {
	hostCfg := ibclient.HostConfig{
		Host:    "localhost",
		Port:    "8080",
		Version: "2.3.1",
	}

	authCfg := ibclient.AuthConfig{
		Username: "user",
		Password: "abcd",
	}

	requestBuilder := NewExtendedRequestBuilder(0, "^staging.*test.com$")
	requestBuilder.Init(hostCfg, authCfg)

	obj := ibclient.NewZoneAuth(ibclient.ZoneAuth{})

	req, _ := requestBuilder.BuildRequest(ibclient.GET, obj, "", &ibclient.QueryParams{})

	assert.True(t, req.URL.Query().Get("fqdn~") == "^staging.*test.com$")

	req, _ = requestBuilder.BuildRequest(ibclient.CREATE, obj, "", &ibclient.QueryParams{})

	assert.True(t, req.URL.Query().Get("fqdn~") == "")
}

func TestExtendedRequestMaxResultsBuilder(t *testing.T) {
	hostCfg := ibclient.HostConfig{
		Host:    "localhost",
		Port:    "8080",
		Version: "2.3.1",
	}

	authCfg := ibclient.AuthConfig{
		Username: "user",
		Password: "abcd",
	}

	requestBuilder := NewExtendedRequestBuilder(54321, "")
	requestBuilder.Init(hostCfg, authCfg)

	obj := ibclient.NewEmptyRecordCNAME()
	obj.Zone = "foo.bar.com"

	req, _ := requestBuilder.BuildRequest(ibclient.GET, obj, "", &ibclient.QueryParams{})

	assert.True(t, req.URL.Query().Get("_max_results") == "54321")

	req, _ = requestBuilder.BuildRequest(ibclient.CREATE, obj, "", &ibclient.QueryParams{})

	assert.True(t, req.URL.Query().Get("_max_results") == "")
}

func TestGetObject(t *testing.T) {
	hostCfg := ibclient.HostConfig{}
	authCfg := ibclient.AuthConfig{}
	transportConfig := ibclient.TransportConfig{}
	requestBuilder := NewExtendedRequestBuilder(1000, "mysite.com")
	requestor := mockRequestor{}
	client, _ := ibclient.NewConnector(hostCfg, authCfg, transportConfig, requestBuilder, &requestor)

	providerConfig := newInfobloxProvider(endpoint.NewDomainFilter([]string{"mysite.com"}), provider.NewZoneIDFilter([]string{""}), true, true, client)

	providerConfig.deleteRecords(infobloxChangeMap{
		"myzone.com": []*endpoint.Endpoint{
			endpoint.NewEndpoint("deletethisrecord.com", endpoint.RecordTypeA, "1.2.3.4"),
		},
	})

	requestQuery := requestor.request.URL.Query()
	assert.True(t, requestQuery.Has("name"), "Expected the request to filter objects by name")
}

// Mock requestor that doesn't send request
type mockRequestor struct {
	request *http.Request
}

func (r *mockRequestor) Init(ibclient.AuthConfig, ibclient.TransportConfig) {}
func (r *mockRequestor) SendRequest(req *http.Request) (res []byte, err error) {
	res = []byte("[{}]")
	r.request = req
	return
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "actual and expected endpoints don't match. %s:%s", endpoints, expected)
}
