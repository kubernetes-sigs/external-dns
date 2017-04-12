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
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
)

const (
	// ID of the hosted zone where the tests are running.
	testZone = "ext-dns-test.teapot.zalan.do."
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
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}
	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func TestAWSCreateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func TestAWSUpdateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", "A"),
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", "A"),
	})
}

func TestAWSDeleteRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}

func TestAWSApplyChanges(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", ""),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
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

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", "A"),
	})
}

func TestAWSApplyNoChanges(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	if err := provider.ApplyChanges(testZone, &plan.Changes{}); err != nil {
		t.Error(err)
	}
}

func TestAWSCreateRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}

func TestAWSUpdateRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", ""),
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func TestAWSDeleteRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func TestAWSApplyChangesDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "1.2.3.4", ""),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", ""),
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

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func TestAWSCreateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", ""),
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	})
}

func TestAWSUpdateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", ""),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "bar.elb.amazonaws.com", ""),
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
	})
}

func TestAWSDeleteRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "baz.elb.amazonaws.com", ""),
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}

func TestAWSApplyChangesCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "qux.elb.amazonaws.com", "CNAME"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", ""),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "bar.elb.amazonaws.com", ""),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "baz.elb.amazonaws.com", ""),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.ext-dns-test.teapot.zalan.do", "qux.elb.amazonaws.com", ""),
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

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.ext-dns-test.teapot.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
		endpoint.NewEndpoint("update-test.ext-dns-test.teapot.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	})
}

func TestAWSSanitizeZone(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})

	records, err = provider.Records("/hostedzone/" + testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.ext-dns-test.teapot.zalan.do", "8.8.8.8", "A"),
	})
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	if !testutils.SameEndpoints(endpoints, expected) {
		t.Fatalf("expected %v, got %v", expected, endpoints)
	}
}

func newAWSProvider(t *testing.T, dryRun bool, records []*endpoint.Endpoint) *AWSProvider {
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

func setupRecords(t *testing.T, provider *AWSProvider, endpoints []*endpoint.Endpoint) {
	clearRecords(t, provider)

	if err := provider.CreateRecords(testZone, endpoints); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}
	validateEndpoints(t, records, endpoints)
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

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}
