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
	"reflect"
	"strings"
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

		change.ResourceRecordSet.Name = aws.String(ensureTrailingDot(aws.StringValue(change.ResourceRecordSet.Name)))

		if change.ResourceRecordSet.AliasTarget != nil {
			change.ResourceRecordSet.AliasTarget.DNSName = aws.String(ensureTrailingDot(aws.StringValue(change.ResourceRecordSet.AliasTarget.DNSName)))
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
		{DNSName: "list-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "list-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
}

func TestAWSCreateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
}

func TestAWSUpdateRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
	})
}

func TestAWSDeleteRecords(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
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
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
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
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
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
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
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
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
}

func TestAWSDeleteRecordsDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}

	if err := provider.DeleteRecords(testZone, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
}

func TestAWSApplyChangesDryRun(t *testing.T) {
	provider := newAWSProvider(t, true, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "1.2.3.4"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
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
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
}

func TestAWSCreateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	})
}

func TestAWSUpdateRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "bar.example.org"},
	}

	if err := provider.UpdateRecords(testZone, updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "bar.example.org"},
	})
}

func TestAWSDeleteRecordsCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "baz.example.org"},
	})

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "baz.example.org"},
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
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "qux.example.org"},
	})

	createRecords := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	}

	currentRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "bar.example.org"},
	}
	updatedRecords := []*endpoint.Endpoint{
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "baz.example.org"},
	}

	deleteRecords := []*endpoint.Endpoint{
		{DNSName: "delete-test.ext-dns-test.teapot.zalan.do", Target: "qux.example.org"},
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
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
		{DNSName: "update-test.ext-dns-test.teapot.zalan.do", Target: "baz.example.org"},
	})
}

func TestAWSCreateRecordsWithCNAME(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.example.org"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	recordSets := listRecords(t, provider.Client)

	validateRecords(t, recordSets, []*route53.ResourceRecordSet{
		{
			Name: aws.String("create-test.ext-dns-test.teapot.zalan.do."),
			Type: aws.String("CNAME"),
			TTL:  aws.Int64(300),
			ResourceRecords: []*route53.ResourceRecord{
				{
					Value: aws.String("foo.example.org"),
				},
			},
		},
	})
}

func TestAWSCreateRecordsWithALIAS(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.ext-dns-test.teapot.zalan.do", Target: "foo.eu-central-1.elb.amazonaws.com"},
	}

	if err := provider.CreateRecords(testZone, records); err != nil {
		t.Fatal(err)
	}

	recordSets := listRecords(t, provider.Client)

	validateRecords(t, recordSets, []*route53.ResourceRecordSet{
		{
			AliasTarget: &route53.AliasTarget{
				DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
				EvaluateTargetHealth: aws.Bool(true),
				HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
			},
			Name: aws.String("create-test.ext-dns-test.teapot.zalan.do."),
			Type: aws.String("A"),
		},
	})
}

func TestAWSIsELBHostname(t *testing.T) {
	for _, tc := range []struct {
		hostname string
		expected bool
	}{
		{"bar.eu-central-1.elb.amazonaws.com", true},
		{"foo.example.org", false},
	} {
		isELB := isELBHostname(tc.hostname)

		if isELB != tc.expected {
			t.Errorf("expected %t, got %t", tc.expected, isELB)
		}
	}
}

func TestAWSCanonicalHostedZone(t *testing.T) {
	for _, tc := range []struct {
		hostname string
		expected string
	}{
		{"foo.us-east-1.elb.amazonaws.com", "Z35SXDOTRQ7X7K"},
		{"foo.us-east-2.elb.amazonaws.com", "Z3AADJGX6KTTL2"},
		{"foo.us-west-1.elb.amazonaws.com", "Z368ELLRRE2KJ0"},
		{"foo.us-west-2.elb.amazonaws.com", "Z1H1FL5HABSF5"},
		{"foo.ca-central-1.elb.amazonaws.com", "ZQSVJUPU6J1EY"},
		{"foo.ap-south-1.elb.amazonaws.com", "ZP97RAFLXTNZK"},
		{"foo.ap-northeast-2.elb.amazonaws.com", "ZWKZPGTI48KDX"},
		{"foo.ap-southeast-1.elb.amazonaws.com", "Z1LMS91P8CMLE5"},
		{"foo.ap-southeast-2.elb.amazonaws.com", "Z1GM3OXH4ZPM65"},
		{"foo.ap-northeast-1.elb.amazonaws.com", "Z14GRHDCWA56QT"},
		{"foo.eu-central-1.elb.amazonaws.com", "Z215JYRZR1TBD5"},
		{"foo.eu-west-1.elb.amazonaws.com", "Z32O12XQLNTSW2"},
		{"foo.eu-west-2.elb.amazonaws.com", "ZHURV8PSTC4K8"},
		{"foo.sa-east-1.elb.amazonaws.com", "Z2P70J7HTTTPLU"},
		{"foo.example.org", ""},
	} {
		zone := canonicalHostedZone(tc.hostname)

		if zone != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, zone)
		}
	}
}

func TestAWSSanitizeZone(t *testing.T) {
	provider := newAWSProvider(t, false, []*endpoint.Endpoint{
		{DNSName: "list-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	records, err := provider.Records(testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "list-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})

	records, err = provider.Records("/hostedzone/" + testZone)
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		{DNSName: "list-test.ext-dns-test.teapot.zalan.do", Target: "8.8.8.8"},
	})
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

func listRecords(t *testing.T, client Route53API) []*route53.ResourceRecordSet {
	recordSets := []*route53.ResourceRecordSet{}
	if err := client.ListResourceRecordSetsPages(&route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(expandedHostedZoneID(testZone)),
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
	return recordSets
}

func clearRecords(t *testing.T, provider *AWSProvider) {
	recordSets := listRecords(t, provider.Client)

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

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	if !testutils.SameEndpoints(endpoints, expected) {
		t.Errorf("expected %v, got %v", expected, endpoints)
	}
}

func validateRecords(t *testing.T, records []*route53.ResourceRecordSet, expected []*route53.ResourceRecordSet) {
	if len(records) != len(expected) {
		t.Errorf("expected %d records, got %d", len(records), len(expected))
	}

	for i := range records {
		if !reflect.DeepEqual(records[i], expected[i]) {
			t.Errorf("record is wrong")
		}
	}
}

func ensureTrailingDot(hostname string) string {
	return strings.TrimSuffix(hostname, ".") + "."
}
