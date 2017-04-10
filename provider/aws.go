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
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

const (
	hostedZonePrefix     = "/hostedzone/"
	elbHostnameSuffix    = ".elb.amazonaws.com"
	evaluateTargetHealth = true
	recordTTL            = 300
)

var (
	// see: https://docs.aws.amazon.com/general/latest/gr/rande.html
	canonicalHostedZones = map[string]string{
		"us-east-1" + elbHostnameSuffix:      "Z35SXDOTRQ7X7K",
		"us-east-2" + elbHostnameSuffix:      "Z3AADJGX6KTTL2",
		"us-west-1" + elbHostnameSuffix:      "Z368ELLRRE2KJ0",
		"us-west-2" + elbHostnameSuffix:      "Z1H1FL5HABSF5",
		"ca-central-1" + elbHostnameSuffix:   "ZQSVJUPU6J1EY",
		"ap-south-1" + elbHostnameSuffix:     "ZP97RAFLXTNZK",
		"ap-northeast-2" + elbHostnameSuffix: "ZWKZPGTI48KDX",
		"ap-southeast-1" + elbHostnameSuffix: "Z1LMS91P8CMLE5",
		"ap-southeast-2" + elbHostnameSuffix: "Z1GM3OXH4ZPM65",
		"ap-northeast-1" + elbHostnameSuffix: "Z14GRHDCWA56QT",
		"eu-central-1" + elbHostnameSuffix:   "Z215JYRZR1TBD5",
		"eu-west-1" + elbHostnameSuffix:      "Z32O12XQLNTSW2",
		"eu-west-2" + elbHostnameSuffix:      "ZHURV8PSTC4K8",
		"sa-east-1" + elbHostnameSuffix:      "Z2P70J7HTTTPLU",
	}
)

// Route53API is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53API interface {
	ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool)) error
	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
}

// AWSProvider is an implementation of Provider for AWS Route53.
type AWSProvider struct {
	Client Route53API
	DryRun bool
}

// NewAWSProvider initializes a new AWS Route53 based Provider.
func NewAWSProvider(dryRun bool) (Provider, error) {
	config := aws.NewConfig()

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	provider := &AWSProvider{
		Client: route53.New(session),
		DryRun: dryRun,
	}

	return provider, nil
}

// Records returns the list of records in a given hosted zone.
func (p *AWSProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	f := func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool) {
		for _, r := range resp.ResourceRecordSets {
			// TODO(linki, ownership): Remove once ownership system is in place.
			// See: https://github.com/kubernetes-incubator/external-dns/pull/122/files/74e2c3d3e237411e619aefc5aab694742001cdec#r109863370
			switch aws.StringValue(r.Type) {
			case route53.RRTypeA, route53.RRTypeCname:
				break
			default:
				continue
			}

			for _, rr := range r.ResourceRecords {
				endpoints = append(endpoints, endpoint.NewEndpoint(aws.StringValue(r.Name), aws.StringValue(rr.Value)))
			}

			if r.AliasTarget != nil {
				endpoints = append(endpoints, endpoint.NewEndpoint(aws.StringValue(r.Name), aws.StringValue(r.AliasTarget.DNSName)))
			}
		}

		return true
	}

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(expandedHostedZoneID(zone)),
	}

	if err := p.Client.ListResourceRecordSetsPages(params, f); err != nil {
		return nil, err
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *AWSProvider) CreateRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newChanges(route53.ChangeActionCreate, endpoints))
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *AWSProvider) UpdateRecords(zone string, endpoints, _ []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newChanges(route53.ChangeActionUpsert, endpoints))
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *AWSProvider) DeleteRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newChanges(route53.ChangeActionDelete, endpoints))
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AWSProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	combinedChanges := make([]*route53.Change, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionUpsert, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionDelete, changes.Delete)...)

	return p.submitChanges(zone, combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *AWSProvider) submitChanges(zone string, changes []*route53.Change) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	if p.DryRun {
		for _, change := range changes {
			log.Infof("Changing records: %s %s", aws.StringValue(change.Action), change.String())
		}

		return nil
	}

	params := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(expandedHostedZoneID(zone)),
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
	}

	if _, err := p.Client.ChangeResourceRecordSets(params); err != nil {
		return err
	}

	return nil
}

// newChanges returns a collection of Changes based on the given records and action.
func newChanges(action string, endpoints []*endpoint.Endpoint) []*route53.Change {
	changes := make([]*route53.Change, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newChange(action, endpoint))
	}

	return changes
}

// newChange returns a Change of the given record by the given action, e.g.
// action=ChangeActionCreate returns a change for creation of the record and
// action=ChangeActionDelete returns a change for deletion of the record.
func newChange(action string, endpoint *endpoint.Endpoint) *route53.Change {
	change := &route53.Change{
		Action: aws.String(action),
		ResourceRecordSet: &route53.ResourceRecordSet{
			Name: aws.String(endpoint.DNSName),
		},
	}

	if isELBHostname(endpoint.Target) {
		change.ResourceRecordSet.Type = aws.String(route53.RRTypeA)
		change.ResourceRecordSet.AliasTarget = &route53.AliasTarget{
			DNSName:              aws.String(endpoint.Target),
			HostedZoneId:         aws.String(canonicalHostedZone(endpoint.Target)),
			EvaluateTargetHealth: aws.Bool(evaluateTargetHealth),
		}
	} else {
		change.ResourceRecordSet.Type = aws.String(suitableType(endpoint.Target))
		change.ResourceRecordSet.TTL = aws.Int64(recordTTL)
		change.ResourceRecordSet.ResourceRecords = []*route53.ResourceRecord{
			{
				Value: aws.String(endpoint.Target),
			},
		}
	}

	return change
}

func isELBHostname(hostname string) bool {
	return canonicalHostedZone(hostname) != ""
}

func canonicalHostedZone(hostname string) string {
	for suffix, zone := range canonicalHostedZones {
		if strings.HasSuffix(hostname, suffix) {
			return zone
		}
	}

	return ""
}

func expandedHostedZoneID(zone string) string {
	return hostedZonePrefix + strings.TrimPrefix(zone, hostedZonePrefix)
}
