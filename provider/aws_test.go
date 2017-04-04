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
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// Compile time check for interface conformance
var _ Route53API = &Route53APIStub{}

// Route53APIStub is a minimal implementation of Route53API, used primarily for unit testing.
// See http://http://docs.aws.amazon.com/sdk-for-go/api/service/route53.html for descriptions
// of all of its methods.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53APIStub struct {
	zones      map[string]*route53.HostedZone
	recordSets map[string]map[string][]*route53.ResourceRecordSet
}

// NewRoute53APIStub returns an initialized Route53APIStub
func NewRoute53APIStub() *Route53APIStub {
	return &Route53APIStub{
		zones:      make(map[string]*route53.HostedZone),
		recordSets: make(map[string]map[string][]*route53.ResourceRecordSet),
	}
}

func (r *Route53APIStub) ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(p *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool)) error {
	output := route53.ListResourceRecordSetsOutput{} // TODO: Support optional input args.
	if len(r.recordSets) <= 0 {
		output.ResourceRecordSets = []*route53.ResourceRecordSet{}
	} else if _, ok := r.recordSets[aws.StringValue(input.HostedZoneId)]; !ok {
		output.ResourceRecordSets = []*route53.ResourceRecordSet{}
	} else {
		for _, rrsets := range r.recordSets[aws.StringValue(input.HostedZoneId)] {
			for _, rrset := range rrsets {
				output.ResourceRecordSets = append(output.ResourceRecordSets, rrset)
			}
		}
	}
	lastPage := true
	fn(&output, lastPage)
	return nil
}

func (r *Route53APIStub) ChangeResourceRecordSets(input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	_, ok := r.zones[aws.StringValue(input.HostedZoneId)]
	if !ok {
		return nil, fmt.Errorf("Hosted zone doesn't exist: %s", aws.StringValue(input.HostedZoneId))
	}

	if len(input.ChangeBatch.Changes) == 0 {
		return nil, fmt.Errorf("ChangeBatch doesn't contain any changes")
	}

	output := &route53.ChangeResourceRecordSetsOutput{}
	recordSets, ok := r.recordSets[aws.StringValue(input.HostedZoneId)]
	if !ok {
		recordSets = make(map[string][]*route53.ResourceRecordSet)
	}

	for _, change := range input.ChangeBatch.Changes {
		key := aws.StringValue(change.ResourceRecordSet.Name) + "::" + aws.StringValue(change.ResourceRecordSet.Type)
		switch aws.StringValue(change.Action) {
		case route53.ChangeActionCreate:
			if _, found := recordSets[key]; found {
				return nil, fmt.Errorf("Attempt to create duplicate rrset %s", key) // TODO: Return AWS errors with codes etc
			}
			recordSets[key] = append(recordSets[key], change.ResourceRecordSet)
		case route53.ChangeActionDelete:
			if _, found := recordSets[key]; !found {
				return nil, fmt.Errorf("Attempt to delete non-existent rrset %s", key) // TODO: Check other fields too
			}
			delete(recordSets, key)
		case route53.ChangeActionUpsert:
			recordSets[key] = []*route53.ResourceRecordSet{change.ResourceRecordSet}
		}
	}
	r.recordSets[aws.StringValue(input.HostedZoneId)] = recordSets
	return output, nil // TODO: We should ideally return status etc, but we don't' use that yet.
}

func (r *Route53APIStub) ListHostedZonesPages(input *route53.ListHostedZonesInput, fn func(p *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool)) error {
	output := &route53.ListHostedZonesOutput{}
	for _, zone := range r.zones {
		output.HostedZones = append(output.HostedZones, zone)
	}
	lastPage := true
	fn(output, lastPage)
	return nil
}

func (r *Route53APIStub) ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	output := &route53.ListHostedZonesByNameOutput{}
	for _, zone := range r.zones {
		if strings.Contains(*input.DNSName, aws.StringValue(zone.Name)) {
			output.HostedZones = append(output.HostedZones, zone)
		}
	}
	return output, nil
}

func (r *Route53APIStub) CreateHostedZone(input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	name := aws.StringValue(input.Name)
	id := "/hostedzone/" + name
	if _, ok := r.zones[id]; ok {
		return nil, fmt.Errorf("Error creating hosted DNS zone: %s already exists", id)
	}
	r.zones[id] = &route53.HostedZone{
		Id:   aws.String(id),
		Name: aws.String(name),
	}
	return &route53.CreateHostedZoneOutput{HostedZone: r.zones[id]}, nil
}

func (r *Route53APIStub) DeleteHostedZone(input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	if _, ok := r.zones[aws.StringValue(input.Id)]; !ok {
		return nil, fmt.Errorf("Error deleting hosted DNS zone: %s does not exist", aws.StringValue(input.Id))
	}
	if len(r.recordSets[aws.StringValue(input.Id)]) > 0 {
		return nil, fmt.Errorf("Error deleting hosted DNS zone: %s has resource records", aws.StringValue(input.Id))
	}
	delete(r.zones, aws.StringValue(input.Id))
	return &route53.DeleteHostedZoneOutput{}, nil
}

func TestAWSZones(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	if len(zones) != 1 {
		t.Fatalf("expected %d zones, got %d", 1, len(zones))
	}

	zone := zones[0]

	if zone != "ext-dns-test.teapot.zalan.do." {
		t.Errorf("expected %s, got %s", "ext-dns-test.teapot.zalan.do.", zone)
	}
}

func TestAWSZone(t *testing.T) {
	provider := newAWSProvider(t, false)

	hostedZone, err := provider.CreateZone("list-ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	zone, err := provider.Zone("list-ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(zone.Id) != aws.StringValue(hostedZone.Id) {
		t.Errorf("expected %s, got %s", aws.StringValue(hostedZone.Id), aws.StringValue(zone.Id))
	}

	if aws.StringValue(zone.Name) != "list-ext-dns-test.teapot.zalan.do." {
		t.Errorf("expected %s, got %s", "list-ext-dns-test.teapot.zalan.do.", aws.StringValue(zone.Name))
	}
}

func TestAWSCreateZone(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, z := range zones {
		if z == "ext-dns-test.teapot.zalan.do." {
			found = true
		}
	}

	if !found {
		t.Fatal("ext-dns-test.teapot.zalan.do. should be there")
	}
}

func TestAWSDeleteZone(t *testing.T) {
	provider := newAWSProvider(t, false)

	zone, err := provider.CreateZone("ext-dns-test-2.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	err = provider.DeleteZone(aws.StringValue(zone.Id))
	if err != nil {
		t.Fatal(err)
	}

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	for _, z := range zones {
		if z == "ext-dns-test-2.teapot.zalan.do." {
			t.Fatal("ext-dns-test-2.teapot.zalan.do.")
		}
	}
}

func TestAWSRecords(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("list-ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	records := []*endpoint.Endpoint{{DNSName: "list-test.list-ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("list-ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("list-ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Errorf("expected %d records, got %d", 1, len(records))
	}

	found := false

	for _, r := range records {
		if r.DNSName == "list-test.list-ext-dns-test.teapot.zalan.do." {
			if r.Target == "8.8.8.8" {
				found = true
			}
		}
	}

	if !found {
		t.Fatal("list-test.list-ext-dns-test.teapot.zalan.do. should be there")
	}
}

func TestAWSCreateRecords(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	records := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "create-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "8.8.8.8" {
				found = true
			}
		}
	}

	if !found {
		t.Fatal("create-test.ext-dns-test.teapot.zalan.do. should be there")
	}
}

func TestAWSUpdateRecords(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	oldRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", oldRecords)
	if err != nil {
		t.Fatal(err)
	}

	newRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"}}

	err = provider.UpdateRecords("ext-dns-test.teapot.zalan.do.", newRecords, oldRecords)
	if err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "update-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "1.2.3.4" {
				found = true
			}
		}
	}

	if !found {
		t.Fatal("update-test.ext-dns-test.teapot.zalan.do. should point to 1.2.3.4")
	}
}

func TestAWSDeleteRecords(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	records := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "20.153.88.175"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	err = provider.DeleteRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "delete-test.ext-dns-test.teapot.zalan.do." {
			found = true
		}
	}

	if found {
		t.Fatal("delete-test.ext-dns-test.teapot.zalan.do. should be gone")
	}
}

func TestAWSApply(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	updateRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", updateRecords)
	if err != nil {
		t.Fatal(err)
	}

	deleteRecords := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "20.153.88.175"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", deleteRecords)
	if err != nil {
		t.Fatal(err)
	}

	createRecords := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}
	updateNewRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"}}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updateNewRecords,
		UpdateOld: updateRecords,
		Delete:    deleteRecords,
	}

	err = provider.ApplyChanges("ext-dns-test.teapot.zalan.do.", changes)
	if err != nil {
		t.Fatal(err)
	}

	// create validation

	records, err := provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "create-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "8.8.8.8" {
				found = true
			}
		}
	}

	if !found {
		t.Fatal("create-test.ext-dns-test.teapot.zalan.do. should be there")
	}

	// update validation

	found = false

	for _, r := range records {
		if r.DNSName == "update-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "1.2.3.4" {
				found = true
			}
		}
	}

	if !found {
		t.Fatal("update-test.ext-dns-test.teapot.zalan.do. should point to 1.2.3.4")
	}

	// delete validation

	found = false

	for _, r := range records {
		if r.DNSName == "delete-test.ext-dns-test.teapot.zalan.do." {
			found = true
		}
	}

	if found {
		t.Fatal("delete-test.ext-dns-test.teapot.zalan.do. should be gone")
	}
}

func TestAWSApplyNoChanges(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	err = provider.ApplyChanges("ext-dns-test.teapot.zalan.do.", &plan.Changes{})
	if err != nil {
		t.Error(err)
	}
}

func TestAWSCreateRecordDryRun(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	provider.DryRun = true

	records := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "create-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "8.8.8.8" {
				found = true
			}
		}
	}

	if found {
		t.Fatal("create-test.ext-dns-test.teapot.zalan.do. should not be there")
	}
}

func TestAWSUpdateRecordDryRun(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	oldRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", oldRecords)
	if err != nil {
		t.Fatal(err)
	}

	provider.DryRun = true

	newRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"}}

	err = provider.UpdateRecords("ext-dns-test.teapot.zalan.do.", newRecords, oldRecords)
	if err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "update-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "1.2.3.4" {
				found = true
			}
		}
	}

	if found {
		t.Fatal("update-test.ext-dns-test.teapot.zalan.do. should not point to 1.2.3.4")
	}
}

func TestAWSDeleteRecordDryRun(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	records := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "20.153.88.175"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	provider.DryRun = true

	err = provider.DeleteRecords("ext-dns-test.teapot.zalan.do.", records)
	if err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "delete-test.ext-dns-test.teapot.zalan.do." {
			found = true
		}
	}

	if !found {
		t.Fatal("delete-test.ext-dns-test.teapot.zalan.do. should not be gone")
	}
}

func TestAWSApplyDryRun(t *testing.T) {
	provider := newAWSProvider(t, false)

	_, err := provider.CreateZone("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	updateRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", updateRecords)
	if err != nil {
		t.Fatal(err)
	}

	deleteRecords := []*endpoint.Endpoint{{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "20.153.88.175"}}

	err = provider.CreateRecords("ext-dns-test.teapot.zalan.do.", deleteRecords)
	if err != nil {
		t.Fatal(err)
	}

	provider.DryRun = true

	createRecords := []*endpoint.Endpoint{{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"}}
	updateNewRecords := []*endpoint.Endpoint{{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"}}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updateNewRecords,
		UpdateOld: updateRecords,
		Delete:    deleteRecords,
	}

	err = provider.ApplyChanges("ext-dns-test.teapot.zalan.do.", changes)
	if err != nil {
		t.Fatal(err)
	}

	// create validation

	records, err := provider.Records("ext-dns-test.teapot.zalan.do.")
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, r := range records {
		if r.DNSName == "create-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "8.8.8.8" {
				found = true
			}
		}
	}

	if found {
		t.Fatal("create-test.ext-dns-test.teapot.zalan.do. should not be there")
	}

	// update validation

	found = false

	for _, r := range records {
		if r.DNSName == "update-test.ext-dns-test.teapot.zalan.do." {
			if r.Target == "1.2.3.4" {
				found = true
			}
		}
	}

	if found {
		t.Fatal("update-test.ext-dns-test.teapot.zalan.do. should not point to 1.2.3.4")
	}

	// delete validation

	found = false

	for _, r := range records {
		if r.DNSName == "delete-test.ext-dns-test.teapot.zalan.do." {
			found = true
		}
	}

	if !found {
		t.Fatal("delete-test.ext-dns-test.teapot.zalan.do. should not be gone")
	}
}

func newAWSProvider(t *testing.T, dryRun bool) *AWSProvider {
	client := NewRoute53APIStub()

	return &AWSProvider{
		Client: client,
		DryRun: dryRun,
	}
}
