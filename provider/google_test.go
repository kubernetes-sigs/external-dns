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
	"net/http"
	"testing"

	"golang.org/x/oauth2/google"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	"google.golang.org/api/dns/v1"
	"google.golang.org/api/googleapi"
)

func TestGoogleZones(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{})

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"zone-1-ext-dns-test-2-gcp-zalan-do": {Name: "zone-1-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-1.ext-dns-test-2.gcp.zalan.do."},
		"zone-2-ext-dns-test-2-gcp-zalan-do": {Name: "zone-2-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-2.ext-dns-test-2.gcp.zalan.do."},
		"zone-3-ext-dns-test-2-gcp-zalan-do": {Name: "zone-3-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-3.ext-dns-test-2.gcp.zalan.do."},
	})
}

func TestGoogleRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
	})

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
	})
}

func TestGoogleCreateRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
	}

	if err := provider.CreateRecords(records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
	})
}

func TestGoogleUpdateRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "4.3.2.1"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "bar.elb.amazonaws.com"},
	}

	if err := provider.UpdateRecords(updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "4.3.2.1"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "bar.elb.amazonaws.com"},
	})
}

func TestGoogleDeleteRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		{DNSName: "delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "baz.elb.amazonaws.com"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "baz.elb.amazonaws.com"},
	}

	if err := provider.DeleteRecords(currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}

func TestGoogleApplyChanges(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "bar.elb.amazonaws.com"},
		{DNSName: "delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "qux.elb.amazonaws.com"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "bar.elb.amazonaws.com"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "4.3.2.1"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "baz.elb.amazonaws.com"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "qux.elb.amazonaws.com"},
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges("_", changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"},
		{DNSName: "create-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "8.8.4.4"},
		{DNSName: "update-test.zone-2.ext-dns-test-2.gcp.zalan.do", Target: "4.3.2.1"},
		{DNSName: "create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "foo.elb.amazonaws.com"},
		{DNSName: "update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", Target: "baz.elb.amazonaws.com"},
	})
}

func TestSeparateChanges(t *testing.T) {
	change := &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.foo.example.org", Ttl: 1},
			{Name: "qux.bar.example.org", Ttl: 2},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.foo.example.org", Ttl: 10},
			{Name: "wambo.bar.example.org", Ttl: 20},
		},
	}

	zones := map[string]*dns.ManagedZone{
		"foo-example-org": &dns.ManagedZone{
			Name:    "foo-example-org",
			DnsName: "foo.example.org.",
		},
		"bar-example-org": &dns.ManagedZone{
			Name:    "bar-example-org",
			DnsName: "bar.example.org.",
		},
		"baz-example-org": &dns.ManagedZone{
			Name:    "baz-example-org",
			DnsName: "baz.example.org.",
		},
	}

	changes := separateChange(zones, change)

	if len(changes) != 2 {
		t.Fatalf("expected %d change(s), got %d", 2, len(changes))
	}

	validateChange(t, changes["foo-example-org"], &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.foo.example.org", Ttl: 1},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.foo.example.org", Ttl: 10},
		},
	})

	validateChange(t, changes["bar-example-org"], &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.bar.example.org", Ttl: 2},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.bar.example.org", Ttl: 20},
		},
	})
}

// GOGOGOGOGOG

// Create with dry run
// func TestGoogleCreateRecordDryRun(t *testing.T) {
// 	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", false, []*endpoint.Endpoint{})
//
// 	records := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"}}
//
// 	err := provider.DeleteRecords("ext-dns-test-2-gcp-zalan-do", records)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	//
//
// 	provider = newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", true, []*endpoint.Endpoint{})
//
// 	err = provider.CreateRecords("ext-dns-test-2-gcp-zalan-do", records)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	records, err = provider.Records("ext-dns-test-2-gcp-zalan-do")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	found := false
//
// 	for _, r := range records {
// 		if r.DNSName == "create-test.ext-dns-test-2.gcp.zalan.do" {
// 			if r.Target == "8.8.8.8" {
// 				found = true
// 			}
// 		}
// 	}
//
// 	if found {
// 		t.Fatal("create-test.ext-dns-test-2.gcp.zalan.do should not be there")
// 	}
// }
//
// // update with dryRun
// func TestGoogleUpdateRecordDryRun(t *testing.T) {
// 	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", false, []*endpoint.Endpoint{})
//
// 	oldRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"}}
// 	newRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"}}
//
// 	err := provider.DeleteRecords("ext-dns-test-2-gcp-zalan-do", newRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	err = provider.CreateRecords("ext-dns-test-2-gcp-zalan-do", oldRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	//
//
// 	provider = newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", true, []*endpoint.Endpoint{})
//
// 	err = provider.UpdateRecords("ext-dns-test-2-gcp-zalan-do", newRecords, oldRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	records, err := provider.Records("ext-dns-test-2-gcp-zalan-do")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	found := false
//
// 	for _, r := range records {
// 		if r.DNSName == "update-test.ext-dns-test-2.gcp.zalan.do" {
// 			if r.Target == "1.2.3.4" {
// 				found = true
// 			}
// 		}
// 	}
//
// 	if found {
// 		t.Fatal("update-test.ext-dns-test-2.gcp.zalan.do should not point to 1.2.3.4")
// 	}
// }
//
// // delete with dryRun
// func TestGoogleDeleteRecordDryRun(t *testing.T) {
// 	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", false, []*endpoint.Endpoint{})
//
// 	records := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test-2.gcp.zalan.do", Target: "20.153.88.175"}}
//
// 	err := provider.CreateRecords("ext-dns-test-2-gcp-zalan-do", records)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	//
//
// 	provider = newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", true, []*endpoint.Endpoint{})
//
// 	err = provider.DeleteRecords("ext-dns-test-2-gcp-zalan-do", records)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	records, err = provider.Records("ext-dns-test-2-gcp-zalan-do")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	found := false
//
// 	for _, r := range records {
// 		if r.DNSName == "delete-test.ext-dns-test-2.gcp.zalan.do" {
// 			found = true
// 		}
// 	}
//
// 	if !found {
// 		t.Fatal("delete-test.ext-dns-test-2.gcp.zalan.do should not be gone")
// 	}
// }
//
// // Apply With DryRun
// func TestGoogleApplyDryRun(t *testing.T) {
// 	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", false, []*endpoint.Endpoint{})
//
// 	// create setup
//
// 	createRecords := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"}}
//
// 	err := provider.DeleteRecords("ext-dns-test-2-gcp-zalan-do", createRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	// update setup
//
// 	oldRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test-2.gcp.zalan.do", Target: "8.8.8.8"}}
// 	newRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test-2.gcp.zalan.do", Target: "1.2.3.4"}}
//
// 	err = provider.DeleteRecords("ext-dns-test-2-gcp-zalan-do", newRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	err = provider.CreateRecords("ext-dns-test-2-gcp-zalan-do", oldRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	// delete setup
//
// 	deleteRecords := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test-2.gcp.zalan.do", Target: "20.153.88.175"}}
//
// 	err = provider.CreateRecords("ext-dns-test-2-gcp-zalan-do", deleteRecords)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	time.Sleep(time.Second)
//
// 	//
//
// 	provider = newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do", true, []*endpoint.Endpoint{})
//
// 	changes := &plan.Changes{
// 		Create:    createRecords,
// 		UpdateNew: newRecords,
// 		UpdateOld: oldRecords,
// 		Delete:    deleteRecords,
// 	}
//
// 	err = provider.ApplyChanges("ext-dns-test-2-gcp-zalan-do", changes)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// create validation
//
// 	records, err := provider.Records("ext-dns-test-2-gcp-zalan-do")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	found := false
//
// 	for _, r := range records {
// 		if r.DNSName == "create-test.ext-dns-test-2.gcp.zalan.do" {
// 			if r.Target == "8.8.8.8" {
// 				found = true
// 			}
// 		}
// 	}
//
// 	if found {
// 		t.Fatal("create-test.ext-dns-test-2.gcp.zalan.do should not be there")
// 	}
//
// 	// update validation
//
// 	found = false
//
// 	for _, r := range records {
// 		if r.DNSName == "update-test.ext-dns-test-2.gcp.zalan.do" {
// 			if r.Target == "1.2.3.4" {
// 				found = true
// 			}
// 		}
// 	}
//
// 	if found {
// 		t.Fatal("update-test.ext-dns-test-2.gcp.zalan.do should not point to 1.2.3.4")
// 	}
//
// 	// delete validation
//
// 	found = false
//
// 	for _, r := range records {
// 		if r.DNSName == "delete-test.ext-dns-test-2.gcp.zalan.do" {
// 			found = true
// 		}
// 	}
//
// 	if !found {
// 		t.Fatal("delete-test.ext-dns-test-2.gcp.zalan.do should not be gone")
// 	}
// }

func validateZones(t *testing.T, zones map[string]*dns.ManagedZone, expected map[string]*dns.ManagedZone) {
	if len(zones) != len(expected) {
		t.Fatalf("expected %d zone(s), got %d", len(expected), len(zones))
	}

	for i, zone := range zones {
		validateZone(t, zone, expected[i])
	}
}

func validateZone(t *testing.T, zone *dns.ManagedZone, expected *dns.ManagedZone) {
	if zone.Name != expected.Name {
		t.Errorf("expected %s, got %s", expected.Name, zone.Name)
	}

	if zone.DnsName != expected.DnsName {
		t.Errorf("expected %s, got %s", expected.DnsName, zone.DnsName)
	}
}

func validateChange(t *testing.T, change *dns.Change, expected *dns.Change) {
	validateChangeRecords(t, change.Additions, expected.Additions)
	validateChangeRecords(t, change.Deletions, expected.Deletions)
}

func validateChangeRecords(t *testing.T, records []*dns.ResourceRecordSet, expected []*dns.ResourceRecordSet) {
	if len(records) != len(expected) {
		t.Fatalf("expected %d change(s), got %d", len(expected), len(records))
	}

	for i := range records {
		validateChangeRecord(t, records[i], expected[i])
	}
}

func validateChangeRecord(t *testing.T, record *dns.ResourceRecordSet, expected *dns.ResourceRecordSet) {
	if record.Name != expected.Name {
		t.Errorf("expected %s, got %s", expected.Name, record.Name)
	}

	if record.Ttl != expected.Ttl {
		t.Errorf("expected %d, got %d", expected.Ttl, record.Ttl)
	}
}

func newGoogleProvider(t *testing.T, domain string, dryRun bool, records []*endpoint.Endpoint) *googleProvider {
	gcloud, err := google.DefaultClient(context.TODO(), dns.NdevClouddnsReadwriteScope)
	if err != nil {
		t.Fatal(err)
	}

	dnsClient, err := dns.New(gcloud)
	if err != nil {
		t.Fatal(err)
	}

	provider := &googleProvider{
		project: "zalando-external-dns-test",
		domain:  domain,
		dryRun:  dryRun,
		resourceRecordSetsClient: resourceRecordSetsService{dnsClient.ResourceRecordSets},
		managedZonesClient:       managedZonesService{dnsClient.ManagedZones},
		changesClient:            changesService{dnsClient.Changes},
	}

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-1-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-1.ext-dns-test-2.gcp.zalan.do.",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-2-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-2.ext-dns-test-2.gcp.zalan.do.",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-3-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-3.ext-dns-test-2.gcp.zalan.do.",
	})

	setupGoogleRecords(t, provider, records)

	provider.dryRun = dryRun

	return provider
}

func createZone(t *testing.T, provider *googleProvider, zone *dns.ManagedZone) {
	zone.Description = "Testing zone for kubernetes.io/external-dns"

	if _, err := provider.managedZonesClient.Create("zalando-external-dns-test", zone).Do(); err != nil {
		if err, ok := err.(*googleapi.Error); !ok || err.Code != http.StatusConflict {
			t.Fatal(err)
		}
	}
}

func setupGoogleRecords(t *testing.T, provider *googleProvider, endpoints []*endpoint.Endpoint) {
	clearGoogleRecords(t, provider, "zone-1-ext-dns-test-2-gcp-zalan-do")
	clearGoogleRecords(t, provider, "zone-2-ext-dns-test-2-gcp-zalan-do")

	records, err := provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})

	if err = provider.CreateRecords(endpoints); err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("_")
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, endpoints)
}

func clearGoogleRecords(t *testing.T, provider *googleProvider, zone string) {
	recordSets := []*dns.ResourceRecordSet{}
	if err := provider.resourceRecordSetsClient.List(provider.project, zone).Pages(context.TODO(), func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			switch r.Type {
			case "A", "CNAME":
				recordSets = append(recordSets, r)
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if len(recordSets) != 0 {
		if _, err := provider.changesClient.Create(provider.project, zone, &dns.Change{
			Deletions: recordSets,
		}).Do(); err != nil {
			t.Fatal(err)
		}
	}
}
