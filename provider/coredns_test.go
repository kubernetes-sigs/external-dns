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

package provider

import (
	"context"
	"strings"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const defaultCoreDNSPrefix = "/skydns/"

type fakeETCDClient struct {
	services map[string]*Service
}

func (c fakeETCDClient) GetServices(prefix string) ([]*Service, error) {
	var result []*Service
	for key, value := range c.services {
		if strings.HasPrefix(key, prefix) {
			value.Key = key
			result = append(result, value)
		}
	}
	return result, nil
}

func (c fakeETCDClient) SaveService(service *Service) error {
	c.services[service.Key] = service
	return nil
}

func (c fakeETCDClient) DeleteService(key string) error {
	delete(c.services, key)
	return nil
}

func TestAServiceTranslation(t *testing.T) {
	expectedTarget := "1.2.3.4"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeA

	client := fakeETCDClient{
		map[string]*Service{
			"/skydns/com/example": {Host: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 1 {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}
	if endpoints[0].DNSName != expectedDNSName {
		t.Errorf("got unexpected DNS name: %s != %s", endpoints[0].DNSName, expectedDNSName)
	}
	if endpoints[0].Targets[0] != expectedTarget {
		t.Errorf("got unexpected DNS target: %s != %s", endpoints[0].Targets[0], expectedTarget)
	}
	if endpoints[0].RecordType != expectedRecordType {
		t.Errorf("got unexpected DNS record type: %s != %s", endpoints[0].RecordType, expectedRecordType)
	}
}

func TestCNAMEServiceTranslation(t *testing.T) {
	expectedTarget := "example.net"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeCNAME

	client := fakeETCDClient{
		map[string]*Service{
			"/skydns/com/example": {Host: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 1 {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}
	if endpoints[0].DNSName != expectedDNSName {
		t.Errorf("got unexpected DNS name: %s != %s", endpoints[0].DNSName, expectedDNSName)
	}
	if endpoints[0].Targets[0] != expectedTarget {
		t.Errorf("got unexpected DNS target: %s != %s", endpoints[0].Targets[0], expectedTarget)
	}
	if endpoints[0].RecordType != expectedRecordType {
		t.Errorf("got unexpected DNS record type: %s != %s", endpoints[0].RecordType, expectedRecordType)
	}
}

func TestTXTServiceTranslation(t *testing.T) {
	expectedTarget := "string"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeTXT

	client := fakeETCDClient{
		map[string]*Service{
			"/skydns/com/example": {Text: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 1 {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}
	if endpoints[0].DNSName != expectedDNSName {
		t.Errorf("got unexpected DNS name: %s != %s", endpoints[0].DNSName, expectedDNSName)
	}
	if endpoints[0].Targets[0] != expectedTarget {
		t.Errorf("got unexpected DNS target: %s != %s", endpoints[0].Targets[0], expectedTarget)
	}
	if endpoints[0].RecordType != expectedRecordType {
		t.Errorf("got unexpected DNS record type: %s != %s", endpoints[0].RecordType, expectedRecordType)
	}
}

func TestAWithTXTServiceTranslation(t *testing.T) {
	expectedTargets := map[string]string{
		endpoint.RecordTypeA:   "1.2.3.4",
		endpoint.RecordTypeTXT: "string",
	}
	expectedDNSName := "example.com"

	client := fakeETCDClient{
		map[string]*Service{
			"/skydns/com/example": {Host: "1.2.3.4", Text: "string"},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != len(expectedTargets) {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}

	for _, ep := range endpoints {
		expectedTarget := expectedTargets[ep.RecordType]
		if expectedTarget == "" {
			t.Errorf("got unexpected DNS record type: %s", ep.RecordType)
			continue
		}
		delete(expectedTargets, ep.RecordType)

		if ep.DNSName != expectedDNSName {
			t.Errorf("got unexpected DNS name: %s != %s", ep.DNSName, expectedDNSName)
		}

		if ep.Targets[0] != expectedTarget {
			t.Errorf("got unexpected DNS target: %s != %s", ep.Targets[0], expectedTarget)
		}
	}
}

func TestCNAMEWithTXTServiceTranslation(t *testing.T) {
	expectedTargets := map[string]string{
		endpoint.RecordTypeCNAME: "example.net",
		endpoint.RecordTypeTXT:   "string",
	}
	expectedDNSName := "example.com"

	client := fakeETCDClient{
		map[string]*Service{
			"/skydns/com/example": {Host: "example.net", Text: "string"},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != len(expectedTargets) {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}

	for _, ep := range endpoints {
		expectedTarget := expectedTargets[ep.RecordType]
		if expectedTarget == "" {
			t.Errorf("got unexpected DNS record type: %s", ep.RecordType)
			continue
		}
		delete(expectedTargets, ep.RecordType)

		if ep.DNSName != expectedDNSName {
			t.Errorf("got unexpected DNS name: %s != %s", ep.DNSName, expectedDNSName)
		}

		if ep.Targets[0] != expectedTarget {
			t.Errorf("got unexpected DNS target: %s != %s", ep.Targets[0], expectedTarget)
		}
	}
}

func TestCoreDNSApplyChanges(t *testing.T) {
	client := fakeETCDClient{
		map[string]*Service{},
	}
	coredns := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}

	changes1 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "5.5.5.5"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeTXT, "string1"),
			endpoint.NewEndpoint("domain2.local", endpoint.RecordTypeCNAME, "site.local"),
		},
	}
	coredns.ApplyChanges(context.Background(), changes1)

	expectedServices1 := map[string]*Service{
		"/skydns/local/domain1": {Host: "5.5.5.5", Text: "string1"},
		"/skydns/local/domain2": {Host: "site.local"},
	}
	validateServices(client.services, expectedServices1, t, 1)

	changes2 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain3.local", endpoint.RecordTypeA, "7.7.7.7"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", "A", "6.6.6.6"),
		},
	}
	records, _ := coredns.Records(context.Background())
	for _, ep := range records {
		if ep.DNSName == "domain1.local" {
			changes2.UpdateOld = append(changes2.UpdateOld, ep)
		}
	}
	applyServiceChanges(coredns, changes2)

	expectedServices2 := map[string]*Service{
		"/skydns/local/domain1": {Host: "6.6.6.6", Text: "string1"},
		"/skydns/local/domain2": {Host: "site.local"},
		"/skydns/local/domain3": {Host: "7.7.7.7"},
	}
	validateServices(client.services, expectedServices2, t, 2)

	changes3 := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "6.6.6.6"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeTXT, "string"),
			endpoint.NewEndpoint("domain3.local", endpoint.RecordTypeA, "7.7.7.7"),
		},
	}

	applyServiceChanges(coredns, changes3)

	expectedServices3 := map[string]*Service{
		"/skydns/local/domain2": {Host: "site.local"},
	}
	validateServices(client.services, expectedServices3, t, 3)
}

func applyServiceChanges(provider coreDNSProvider, changes *plan.Changes) {
	ctx := context.Background()
	records, _ := provider.Records(ctx)
	for _, col := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew, changes.Delete} {
		for _, record := range col {
			for _, existingRecord := range records {
				if existingRecord.DNSName == record.DNSName && existingRecord.RecordType == record.RecordType {
					mergeLabels(record, existingRecord.Labels)
				}
			}
		}
	}
	provider.ApplyChanges(ctx, changes)
}

func validateServices(services, expectedServices map[string]*Service, t *testing.T, step int) {
	if len(services) != len(expectedServices) {
		t.Errorf("wrong number of records on step %d: %d != %d", step, len(services), len(expectedServices))
	}
	for key, value := range services {
		keyParts := strings.Split(key, "/")
		expectedKey := strings.Join(keyParts[:len(keyParts)-value.TargetStrip], "/")
		expectedService := expectedServices[expectedKey]
		if expectedService == nil {
			t.Errorf("unexpected service %s", key)
			continue
		}
		delete(expectedServices, key)
		if value.Host != expectedService.Host {
			t.Errorf("wrong host for service %s: %s != %s on step %d", key, value.Host, expectedService.Host, step)
		}
		if value.Text != expectedService.Text {
			t.Errorf("wrong text for service %s: %s != %s on step %d", key, value.Text, expectedService.Text, step)
		}
	}
}

// mergeLabels adds keys to labels if not defined for the endpoint
func mergeLabels(e *endpoint.Endpoint, labels map[string]string) {
	for k, v := range labels {
		if e.Labels[k] == "" {
			e.Labels[k] = v
		}
	}
}
