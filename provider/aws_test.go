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
	"net"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

const (
	testZone = "/hostedzone/ext-dns-test.teapot.zalan.do."
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
		if aws.StringValue(change.ResourceRecordSet.Type) == route53.RRTypeA {
			for _, rrs := range change.ResourceRecordSet.ResourceRecords {
				if net.ParseIP(aws.StringValue(rrs.Value)) == nil {
					return nil, fmt.Errorf("A records must point to IPs")
				}
			}
		}

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

func TestAWSRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "list-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "list-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})
}

func TestAWSCreateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "create-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})
}

func TestAWSUpdateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "1.2.3.4"},
	})
}

func TestAWSDeleteRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{})
}

func TestAWSApplyChanges(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(testZone, changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "create-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "1.2.3.4"},
	})
}

func TestAWSApplyNoChanges(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{})

	if err := provider.ApplyChanges(testZone, &plan.Changes{}); err != nil {
		t.Error(err)
	}
}

func TestAWSCreateRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []map[string]string{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{})
}

func TestAWSUpdateRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})
}

func TestAWSDeleteRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []map[string]string{
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})
}

func TestAWSApplyChangesDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "1.2.3.4"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "8.8.8.8"},
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(testZone, changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "8.8.8.8"},
	})
}

func TestAWSCreateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "foo.elb.amazonaws.com"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "create-test.ext-dns-test.teapot.zalan.do.", "target": "foo.elb.amazonaws.com"},
	})
}

func TestAWSUpdateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "foo.elb.amazonaws.com"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "foo.elb.amazonaws.com"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "bar.elb.amazonaws.com"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "bar.elb.amazonaws.com"},
	})
}

func TestAWSDeleteRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "baz.elb.amazonaws.com"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "baz.elb.amazonaws.com"},
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{})
}

func TestAWSApplyChangesCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []map[string]string{
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "foo.elb.amazonaws.com"},
		{"dnsname": "delete-test.ext-dns-test.teapot.zalan.do.", "target": "qux.elb.amazonaws.com"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do.", Target: "foo.elb.amazonaws.com"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "bar.elb.amazonaws.com"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do.", Target: "baz.elb.amazonaws.com"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do.", Target: "qux.elb.amazonaws.com"},
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(testZone, changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{
		{"dnsname": "create-test.ext-dns-test.teapot.zalan.do.", "target": "foo.elb.amazonaws.com"},
		{"dnsname": "update-test.ext-dns-test.teapot.zalan.do.", "target": "baz.elb.amazonaws.com"},
	})
}

func validateZones(t *testing.T, zones []*route53.HostedZone, expected []map[string]string) {
	if len(zones) != len(expected) {
		t.Fatalf("expected %d zone(s), got %d", len(expected), len(zones))
	}

	for i, zone := range zones {
		validateZone(t, zone, expected[i])
	}
}

func validateZone(t *testing.T, zone *route53.HostedZone, expected map[string]string) {
	if aws.StringValue(zone.Id) != expected["id"] {
		t.Errorf("expected %s, got %s", expected["id"], aws.StringValue(zone.Id))
	}

	if aws.StringValue(zone.Name) != expected["domain"] {
		t.Errorf("expected %s, got %s", expected["domain"], aws.StringValue(zone.Name))
	}
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []map[string]string) {
	if len(endpoints) != len(expected) {
		t.Fatalf("expected %d endpoint(s), got %d", len(expected), len(endpoints))
	}

	sort.Slice(endpoints, func(i, j int) bool { return endpoints[i].DNSName < endpoints[j].DNSName })
	sort.Slice(expected, func(i, j int) bool { return expected[i]["dnsname"] < expected[j]["dnsname"] })

	for i, ep := range endpoints {
		validateEndpoint(t, ep, expected[i])
	}
}

func validateEndpoint(t *testing.T, ep *endpoint.Endpoint, expected map[string]string) {
	if ep.DNSName != expected["dnsname"] {
		t.Errorf("expected %s, got %s", expected["dnsname"], ep.DNSName)
	}

	if ep.Target != expected["target"] {
		t.Errorf("expected %s, got %s", expected["target"], ep.Target)
	}
}

func newAWSProvider(t *testing.T, dryRun bool, records []map[string]string) *AWSProvider {
	client := NewRoute53APIStub()

	if _, err := client.CreateHostedZone(&route53.CreateHostedZoneInput{
		CallerReference: aws.String("external-dns.alpha.kubernetes.io/test-zone"),
		Name:            aws.String("ext-dns-test.teapot.zalan.do."),
	}); err != nil {
		if err, ok := err.(awserr.Error); !ok || err.Code() != route53.ErrCodeHostedZoneAlreadyExists {
			t.Fatal(err)
		}
	}

	provider := &AWSProvider{
		Client: client,
		DryRun: false,
	}

	setupRecords(t, provider, records)

	provider.DryRun = dryRun

	return provider
}

func setupRecords(t *testing.T, provider *AWSProvider, seed []map[string]string) {
	clearRecords(t, provider)

	endpoints := make([]*endpoint.Endpoint, 0, len(seed))

	for _, record := range seed {
		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName: record["dnsname"], Target: record["target"],
		})
	}

	if err := provider.CreateRecords(testZone, endpoints); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, seed)
}

func clearRecords(t *testing.T, provider *AWSProvider) {
	recordSets := []*route53.ResourceRecordSet{}
	if err := provider.Client.ListResourceRecordSetsPages(&route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(testZone),
	}, func(resp *route53.ListResourceRecordSetsOutput, _ bool) bool {
		for _, recordSet := range resp.ResourceRecordSets {
			switch aws.StringValue(recordSet.Type) {
			case "A", "CNAME":
				recordSets = append(recordSets, recordSet)
			}
		}
		return true
	}); err != nil {
		t.Fatal(err)
	}

	changes := make([]*route53.Change, 0, len(recordSets))
	for _, recordSet := range recordSets {
		changes = append(changes, &route53.Change{
			Action:            aws.String(route53.ChangeActionDelete),
			ResourceRecordSet: recordSet,
		})
	}

	if len(changes) != 0 {
		if _, err := provider.Client.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(testZone),
			ChangeBatch: &route53.ChangeBatch{
				Changes: changes,
			},
		}); err != nil {
			t.Fatal(err)
		}
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []map[string]string{})
}
