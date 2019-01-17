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
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"
)

const (
	recordTTL = 300
	// provider specific key that designates whether an AWS ALIAS record has the EvaluateTargetHealth
	// field set to true.
	providerSpecificEvaluateTargetHealth = "aws/evaluate-target-health"
)

var (
	// see: https://docs.aws.amazon.com/general/latest/gr/rande.html#elb_region
	canonicalHostedZones = map[string]string{
		// Application Load Balancers and Classic Load Balancers
		"us-east-2.elb.amazonaws.com":      "Z3AADJGX6KTTL2",
		"us-east-1.elb.amazonaws.com":      "Z35SXDOTRQ7X7K",
		"us-west-1.elb.amazonaws.com":      "Z368ELLRRE2KJ0",
		"us-west-2.elb.amazonaws.com":      "Z1H1FL5HABSF5",
		"ca-central-1.elb.amazonaws.com":   "ZQSVJUPU6J1EY",
		"ap-south-1.elb.amazonaws.com":     "ZP97RAFLXTNZK",
		"ap-northeast-2.elb.amazonaws.com": "ZWKZPGTI48KDX",
		"ap-northeast-3.elb.amazonaws.com": "Z5LXEXXYW11ES",
		"ap-southeast-1.elb.amazonaws.com": "Z1LMS91P8CMLE5",
		"ap-southeast-2.elb.amazonaws.com": "Z1GM3OXH4ZPM65",
		"ap-northeast-1.elb.amazonaws.com": "Z14GRHDCWA56QT",
		"eu-central-1.elb.amazonaws.com":   "Z215JYRZR1TBD5",
		"eu-west-1.elb.amazonaws.com":      "Z32O12XQLNTSW2",
		"eu-west-2.elb.amazonaws.com":      "ZHURV8PSTC4K8",
		"eu-west-3.elb.amazonaws.com":      "Z3Q77PNBQS71R4",
		"eu-north-1.elb.amazonaws.com":     "Z23TAZ6LKFMNIO",
		"sa-east-1.elb.amazonaws.com":      "Z2P70J7HTTTPLU",
		// Network Load Balancers
		"elb.us-east-2.amazonaws.com":      "ZLMOA37VPKANP",
		"elb.us-east-1.amazonaws.com":      "Z26RNL4JYFTOTI",
		"elb.us-west-1.amazonaws.com":      "Z24FKFUX50B4VW",
		"elb.us-west-2.amazonaws.com":      "Z18D5FSROUN65G",
		"elb.ca-central-1.amazonaws.com":   "Z2EPGBW3API2WT",
		"elb.ap-south-1.amazonaws.com":     "ZVDDRBQ08TROA",
		"elb.ap-northeast-2.amazonaws.com": "ZIBE1TIR4HY56",
		"elb.ap-southeast-1.amazonaws.com": "ZKVM4W9LS7TM",
		"elb.ap-southeast-2.amazonaws.com": "ZCT6FZBF4DROD",
		"elb.ap-northeast-1.amazonaws.com": "Z31USIVHYNEOWT",
		"elb.eu-central-1.amazonaws.com":   "Z3F0SRJ5LGBH90",
		"elb.eu-west-1.amazonaws.com":      "Z2IFOLAFXWLO4F",
		"elb.eu-west-2.amazonaws.com":      "ZD4D7Y8KGAS4G",
		"elb.eu-west-3.amazonaws.com":      "Z1CMS0P5QUZ6D5",
		"elb.eu-north-1.amazonaws.com":     "Z1UDT6IFJ4EJM",
		"elb.sa-east-1.amazonaws.com":      "ZTK26PT1VY4CU",
	}
)

// Route53API is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53API interface {
	ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool)) error
	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	ListHostedZonesPages(input *route53.ListHostedZonesInput, fn func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool)) error
	ListTagsForResource(input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)
}

// AWSProvider is an implementation of Provider for AWS Route53.
type AWSProvider struct {
	client               Route53API
	dryRun               bool
	batchChangeSize      int
	batchChangeInterval  time.Duration
	evaluateTargetHealth bool
	// only consider hosted zones managing domains ending in this suffix
	domainFilter DomainFilter
	// filter hosted zones by id
	zoneIDFilter ZoneIDFilter
	// filter hosted zones by type (e.g. private or public)
	zoneTypeFilter ZoneTypeFilter
	// filter hosted zones by tags
	zoneTagFilter ZoneTagFilter
}

// AWSConfig contains configuration to create a new AWS provider.
type AWSConfig struct {
	DomainFilter         DomainFilter
	ZoneIDFilter         ZoneIDFilter
	ZoneTypeFilter       ZoneTypeFilter
	ZoneTagFilter        ZoneTagFilter
	BatchChangeSize      int
	BatchChangeInterval  time.Duration
	EvaluateTargetHealth bool
	AssumeRole           string
	APIRetries           int
	DryRun               bool
}

// NewAWSProvider initializes a new AWS Route53 based Provider.
func NewAWSProvider(awsConfig AWSConfig) (*AWSProvider, error) {
	config := aws.NewConfig().WithMaxRetries(awsConfig.APIRetries)

	config.WithHTTPClient(
		instrumented_http.NewClient(config.HTTPClient, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		}),
	)

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	if awsConfig.AssumeRole != "" {
		log.Infof("Assuming role: %s", awsConfig.AssumeRole)
		session.Config.WithCredentials(stscreds.NewCredentials(session, awsConfig.AssumeRole))
	}

	provider := &AWSProvider{
		client:               route53.New(session),
		domainFilter:         awsConfig.DomainFilter,
		zoneIDFilter:         awsConfig.ZoneIDFilter,
		zoneTypeFilter:       awsConfig.ZoneTypeFilter,
		zoneTagFilter:        awsConfig.ZoneTagFilter,
		batchChangeSize:      awsConfig.BatchChangeSize,
		batchChangeInterval:  awsConfig.BatchChangeInterval,
		evaluateTargetHealth: awsConfig.EvaluateTargetHealth,
		dryRun:               awsConfig.DryRun,
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *AWSProvider) Zones() (map[string]*route53.HostedZone, error) {
	zones := make(map[string]*route53.HostedZone)

	var tagErr error
	f := func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool) {
		for _, zone := range resp.HostedZones {
			if !p.zoneIDFilter.Match(aws.StringValue(zone.Id)) {
				continue
			}

			if !p.zoneTypeFilter.Match(zone) {
				continue
			}

			if !p.domainFilter.Match(aws.StringValue(zone.Name)) {
				continue
			}

			// Only fetch tags if a tag filter was specified
			if !p.zoneTagFilter.IsEmpty() {
				tags, err := p.tagsForZone(*zone.Id)
				if err != nil {
					tagErr = err
					return false
				}
				if !p.zoneTagFilter.Match(tags) {
					continue
				}
			}

			zones[aws.StringValue(zone.Id)] = zone
		}

		return true
	}

	err := p.client.ListHostedZonesPages(&route53.ListHostedZonesInput{}, f)
	if err != nil {
		return nil, err
	}
	if tagErr != nil {
		return nil, tagErr
	}

	for _, zone := range zones {
		log.Debugf("Considering zone: %s (domain: %s)", aws.StringValue(zone.Id), aws.StringValue(zone.Name))
	}

	return zones, nil
}

// wildcardUnescape converts \\052.abc back to *.abc
// Route53 stores wildcards escaped: http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html?shortFooter=true#domain-name-format-asterisk
func wildcardUnescape(s string) string {
	if strings.Contains(s, "\\052") {
		s = strings.Replace(s, "\\052", "*", 1)
	}
	return s
}

// Records returns the list of records in a given hosted zone.
func (p *AWSProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	f := func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool) {
		for _, r := range resp.ResourceRecordSets {
			// TODO(linki, ownership): Remove once ownership system is in place.
			// See: https://github.com/kubernetes-incubator/external-dns/pull/122/files/74e2c3d3e237411e619aefc5aab694742001cdec#r109863370

			if !supportedRecordType(aws.StringValue(r.Type)) {
				continue
			}

			var ttl endpoint.TTL
			if r.TTL != nil {
				ttl = endpoint.TTL(*r.TTL)
			}

			if len(r.ResourceRecords) > 0 {
				targets := make([]string, len(r.ResourceRecords))
				for idx, rr := range r.ResourceRecords {
					targets[idx] = aws.StringValue(rr.Value)
				}

				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(wildcardUnescape(aws.StringValue(r.Name)), aws.StringValue(r.Type), ttl, targets...))
			}

			if r.AliasTarget != nil {
				ep := endpoint.
					NewEndpointWithTTL(wildcardUnescape(aws.StringValue(r.Name)), endpoint.RecordTypeCNAME, ttl, aws.StringValue(r.AliasTarget.DNSName)).
					WithProviderSpecific(providerSpecificEvaluateTargetHealth, fmt.Sprintf("%t", aws.BoolValue(r.AliasTarget.EvaluateTargetHealth)))
				endpoints = append(endpoints, ep)
			}
		}

		return true
	}

	for _, z := range zones {
		params := &route53.ListResourceRecordSetsInput{
			HostedZoneId: z.Id,
		}

		if err := p.client.ListResourceRecordSetsPages(params, f); err != nil {
			return nil, err
		}
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *AWSProvider) CreateRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(p.newChanges(route53.ChangeActionCreate, endpoints))
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *AWSProvider) UpdateRecords(endpoints, _ []*endpoint.Endpoint) error {
	return p.submitChanges(p.newChanges(route53.ChangeActionUpsert, endpoints))
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *AWSProvider) DeleteRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(p.newChanges(route53.ChangeActionDelete, endpoints))
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AWSProvider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*route53.Change, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, p.newChanges(route53.ChangeActionCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, p.newChanges(route53.ChangeActionUpsert, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, p.newChanges(route53.ChangeActionDelete, changes.Delete)...)

	return p.submitChanges(combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *AWSProvider) submitChanges(changes []*route53.Change) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	zones, err := p.Zones()
	if err != nil {
		return err
	}

	// separate into per-zone change sets to be passed to the API.
	changesByZone := changesByZone(zones, changes)
	if len(changesByZone) == 0 {
		log.Info("All records are already up to date, there are no changes for the matching hosted zones")
	}

	var failedZones []string
	for z, cs := range changesByZone {
		var failedUpdate bool

		batchCs := batchChangeSet(cs, p.batchChangeSize)

		for i, b := range batchCs {
			for _, c := range b {
				log.Infof("Desired change: %s %s %s", *c.Action, *c.ResourceRecordSet.Name, *c.ResourceRecordSet.Type)
			}

			if !p.dryRun {
				params := &route53.ChangeResourceRecordSetsInput{
					HostedZoneId: aws.String(z),
					ChangeBatch: &route53.ChangeBatch{
						Changes: b,
					},
				}

				if _, err := p.client.ChangeResourceRecordSets(params); err != nil {
					log.Error(err) //TODO(ideahitme): consider changing the interface in cases when this error might be a concern for other components
					failedUpdate = true
				} else {
					log.Infof("%d record(s) in zone %s were successfully updated", len(b), aws.StringValue(zones[z].Name))
				}

				if i != len(batchCs)-1 {
					time.Sleep(p.batchChangeInterval)
				}
			}
		}

		if failedUpdate {
			failedZones = append(failedZones, z)
		}
	}

	if len(failedZones) > 0 {
		return fmt.Errorf("Failed to submit all changes for the following zones: %v", failedZones)
	}

	return nil
}

// newChanges returns a collection of Changes based on the given records and action.
func (p *AWSProvider) newChanges(action string, endpoints []*endpoint.Endpoint) []*route53.Change {
	changes := make([]*route53.Change, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, p.newChange(action, endpoint))
	}

	return changes
}

// newChange returns a Change of the given record by the given action, e.g.
// action=ChangeActionCreate returns a change for creation of the record and
// action=ChangeActionDelete returns a change for deletion of the record.
func (p *AWSProvider) newChange(action string, endpoint *endpoint.Endpoint) *route53.Change {
	change := &route53.Change{
		Action: aws.String(action),
		ResourceRecordSet: &route53.ResourceRecordSet{
			Name: aws.String(endpoint.DNSName),
		},
	}

	rec, err := p.Records()
	if err != nil {
		log.Infof("getting records failed: %v", err)
	}

	if isAWSLoadBalancer(endpoint) {
		evalTargetHealth := p.evaluateTargetHealth
		if _, ok := endpoint.ProviderSpecific[providerSpecificEvaluateTargetHealth]; ok {
			evalTargetHealth = endpoint.ProviderSpecific[providerSpecificEvaluateTargetHealth] == "true"
		}

		change.ResourceRecordSet.Type = aws.String(route53.RRTypeA)
		change.ResourceRecordSet.AliasTarget = &route53.AliasTarget{
			DNSName:              aws.String(endpoint.Targets[0]),
			HostedZoneId:         aws.String(canonicalHostedZone(endpoint.Targets[0])),
			EvaluateTargetHealth: aws.Bool(evalTargetHealth),
		}
	} else if hostedZone := isAWSAlias(endpoint, rec); hostedZone != "" {
		zones, err := p.Zones()
		if err != nil {
			log.Errorf("getting zones failed: %v", err)
		}
		for _, zone := range zones {
			change.ResourceRecordSet.Type = aws.String(route53.RRTypeA)
			change.ResourceRecordSet.AliasTarget = &route53.AliasTarget{
				DNSName:              aws.String(endpoint.Targets[0]),
				HostedZoneId:         aws.String(cleanZoneID(*zone.Id)),
				EvaluateTargetHealth: aws.Bool(p.evaluateTargetHealth),
			}
		}
	} else {
		change.ResourceRecordSet.Type = aws.String(endpoint.RecordType)
		if !endpoint.RecordTTL.IsConfigured() {
			change.ResourceRecordSet.TTL = aws.Int64(recordTTL)
		} else {
			change.ResourceRecordSet.TTL = aws.Int64(int64(endpoint.RecordTTL))
		}
		change.ResourceRecordSet.ResourceRecords = make([]*route53.ResourceRecord, len(endpoint.Targets))
		for idx, val := range endpoint.Targets {
			change.ResourceRecordSet.ResourceRecords[idx] = &route53.ResourceRecord{
				Value: aws.String(val),
			}
		}
	}

	return change
}

func (p *AWSProvider) tagsForZone(zoneID string) (map[string]string, error) {
	response, err := p.client.ListTagsForResource(&route53.ListTagsForResourceInput{
		ResourceType: aws.String("hostedzone"),
		ResourceId:   aws.String(zoneID),
	})
	if err != nil {
		return nil, err
	}
	tagMap := map[string]string{}
	for _, tag := range response.ResourceTagSet.Tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap, nil
}

func batchChangeSet(cs []*route53.Change, batchSize int) [][]*route53.Change {
	if len(cs) <= batchSize {
		return [][]*route53.Change{cs}
	}

	batchChanges := make([][]*route53.Change, 0)

	changesByName := make(map[string][]*route53.Change, 0)
	for _, v := range cs {
		changesByName[*v.ResourceRecordSet.Name] = append(changesByName[*v.ResourceRecordSet.Name], v)
	}

	names := make([]string, 0)
	for v := range changesByName {
		names = append(names, v)
	}
	sort.Strings(names)

	for _, name := range names {
		totalChangesByName := len(changesByName[name])

		if totalChangesByName > batchSize {
			log.Warnf("Total changes for %s exceeds max batch size of %d, total changes: %d", name,
				batchSize, totalChangesByName)
			continue
		}

		var existingBatch bool
		for i, b := range batchChanges {
			if len(b)+totalChangesByName <= batchSize {
				batchChanges[i] = append(batchChanges[i], changesByName[name]...)
				existingBatch = true
				break
			}
		}
		if !existingBatch {
			batchChanges = append(batchChanges, changesByName[name])
		}
	}

	for i, batch := range batchChanges {
		batchChanges[i] = sortChangesByActionNameType(batch)
	}

	return batchChanges
}

func sortChangesByActionNameType(cs []*route53.Change) []*route53.Change {
	sort.SliceStable(cs, func(i, j int) bool {
		if *cs[i].Action < *cs[j].Action {
			return true
		}
		if *cs[i].Action > *cs[j].Action {
			return false
		}
		if *cs[i].ResourceRecordSet.Name < *cs[j].ResourceRecordSet.Name {
			return true
		}
		if *cs[i].ResourceRecordSet.Name > *cs[j].ResourceRecordSet.Name {
			return false
		}
		return *cs[i].ResourceRecordSet.Type < *cs[j].ResourceRecordSet.Type
	})

	return cs
}

// changesByZone separates a multi-zone change into a single change per zone.
func changesByZone(zones map[string]*route53.HostedZone, changeSet []*route53.Change) map[string][]*route53.Change {
	changes := make(map[string][]*route53.Change)

	for _, z := range zones {
		changes[aws.StringValue(z.Id)] = []*route53.Change{}
	}

	for _, c := range changeSet {
		hostname := ensureTrailingDot(aws.StringValue(c.ResourceRecordSet.Name))

		zones := suitableZones(hostname, zones)
		if len(zones) == 0 {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected ", c.String())
			continue
		}
		for _, z := range zones {
			changes[aws.StringValue(z.Id)] = append(changes[aws.StringValue(z.Id)], c)
			log.Debugf("Adding %s to zone %s [Id: %s]", hostname, aws.StringValue(z.Name), aws.StringValue(z.Id))
		}
	}

	// separating a change could lead to empty sub changes, remove them here.
	for zone, change := range changes {
		if len(change) == 0 {
			delete(changes, zone)
		}
	}

	return changes
}

// suitableZones returns all suitable private zones and the most suitable public zone
//   for a given hostname and a set of zones.
func suitableZones(hostname string, zones map[string]*route53.HostedZone) []*route53.HostedZone {
	var matchingZones []*route53.HostedZone
	var publicZone *route53.HostedZone

	for _, z := range zones {
		if aws.StringValue(z.Name) == hostname || strings.HasSuffix(hostname, "."+aws.StringValue(z.Name)) {
			if z.Config == nil || !aws.BoolValue(z.Config.PrivateZone) {
				// Only select the best matching public zone
				if publicZone == nil || len(aws.StringValue(z.Name)) > len(aws.StringValue(publicZone.Name)) {
					publicZone = z
				}
			} else {
				// Include all private zones
				matchingZones = append(matchingZones, z)
			}
		}
	}

	if publicZone != nil {
		matchingZones = append(matchingZones, publicZone)
	}

	return matchingZones
}

// isAWSLoadBalancer determines if a given hostname belongs to an AWS load balancer.
func isAWSLoadBalancer(ep *endpoint.Endpoint) bool {
	if ep.RecordType == endpoint.RecordTypeCNAME {
		return canonicalHostedZone(ep.Targets[0]) != ""
	}

	return false
}

// isAWSAlias determines if a given hostname belongs to an AWS Alias record by doing an reverse lookup.
func isAWSAlias(ep *endpoint.Endpoint, addrs []*endpoint.Endpoint) string {
	if val, exists := ep.ProviderSpecific["alias"]; ep.RecordType == endpoint.RecordTypeCNAME && exists && val == "true" {
		for _, addr := range addrs {
			if addr.DNSName == ep.Targets[0] {
				if hostedZone := canonicalHostedZone(addr.Targets[0]); hostedZone != "" {
					return hostedZone
				}

			}
		}
	}
	return ""
}

// canonicalHostedZone returns the matching canonical zone for a given hostname.
func canonicalHostedZone(hostname string) string {
	for suffix, zone := range canonicalHostedZones {
		if strings.HasSuffix(hostname, suffix) {
			return zone
		}
	}

	return ""
}

// cleanZoneID removes the "/hostedzone/" prefix
func cleanZoneID(ID string) string {
	if strings.HasPrefix(ID, "/hostedzone/") {
		ID = strings.TrimPrefix(ID, "/hostedzone/")
	}
	return ID
}
