/*
Copyright 2025 The Kubernetes Authors.

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

package aws

import (
  "context"
  "os"
  "testing"
  "time"

  "github.com/aws/aws-sdk-go-v2/service/route53"
  route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
  "github.com/stretchr/testify/assert"
  "gopkg.in/yaml.v3"
  "sigs.k8s.io/external-dns/endpoint"
  "sigs.k8s.io/external-dns/provider"
)

type HostedZones struct {
  Zones []*HostedZone `yaml:"zones"`
}

type HostedZone struct {
  Name string
  ID   string
  Tags []route53types.Tag `yaml:"tags"`
}

var _ Route53API = &Route53APIFixtureStub{}

type Route53APIFixtureStub struct {
  zones    map[string]*route53types.HostedZone
  zoneTags map[string][]route53types.Tag
  calls    map[string]int
}

func providerFilters(client *Route53APIFixtureStub, options ...func(awsProvider *AWSProvider)) *AWSProvider {
  p := &AWSProvider{
    clients:              map[string]Route53API{defaultAWSProfile: client},
    evaluateTargetHealth: false,
    dryRun:               false,
    domainFilter:         endpoint.NewDomainFilter([]string{}),
    zoneIDFilter:         provider.NewZoneIDFilter([]string{}),
    zoneTypeFilter:       provider.NewZoneTypeFilter(""),
    zoneTagFilter:        provider.NewZoneTagFilter([]string{}),
    zonesCache:           &zonesListCache{duration: 1 * time.Second},
  }
  for _, o := range options {
    o(p)
  }
  return p
}

func WithDomainFilters(filters ...string) func(awsProvider *AWSProvider) {
  return func(awsProvider *AWSProvider) {
    awsProvider.domainFilter = endpoint.NewDomainFilter(filters)
  }
}

func WithZoneIDFilters(filters ...string) func(awsProvider *AWSProvider) {
  return func(awsProvider *AWSProvider) {
    awsProvider.zoneIDFilter = provider.NewZoneIDFilter(filters)
  }
}

func WithZoneTagFilters(filters []string) func(awsProvider *AWSProvider) {
  return func(awsProvider *AWSProvider) {
    awsProvider.zoneTagFilter = provider.NewZoneTagFilter(filters)
  }
}

func NewRoute53APIFixtureStub(zones *HostedZones) *Route53APIFixtureStub {
  route53Zones := make(map[string]*route53types.HostedZone)
  zoneTags := make(map[string][]route53types.Tag)
  for _, zone := range zones.Zones {
    route53Zones[zone.ID] = &route53types.HostedZone{
      Id:   &zone.ID,
      Name: &zone.Name,
    }
    zoneTags[cleanZoneID(zone.ID)] = zone.Tags
  }
  return &Route53APIFixtureStub{
    zones:    route53Zones,
    zoneTags: zoneTags,
    calls:    make(map[string]int),
  }
}

func (r Route53APIFixtureStub) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
  // TODO implement me
  panic("implement me")
}

func (r Route53APIFixtureStub) ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error) {
  // TODO implement me
  panic("implement me")
}

func (r Route53APIFixtureStub) CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput, optFns ...func(*route53.Options)) (*route53.CreateHostedZoneOutput, error) {
  // TODO implement me
  panic("implement me")
}

func (r Route53APIFixtureStub) ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput, optFns ...func(options *route53.Options)) (*route53.ListHostedZonesOutput, error) {
  r.calls["listhostedzones"]++
  output := &route53.ListHostedZonesOutput{}
  for _, zone := range r.zones {
    output.HostedZones = append(output.HostedZones, *zone)
  }
  return output, nil
}

func (r Route53APIFixtureStub) ListTagsForResource(ctx context.Context, input *route53.ListTagsForResourceInput, optFns ...func(options *route53.Options)) (*route53.ListTagsForResourceOutput, error) {
  r.calls["listtagsforresource"]++
  tags := r.zoneTags[*input.ResourceId]
  return &route53.ListTagsForResourceOutput{
    ResourceTagSet: &route53types.ResourceTagSet{
      ResourceId:   input.ResourceId,
      ResourceType: input.ResourceType,
      Tags:         tags,
    },
  }, nil
}

func unmarshalTestHelper(input string, obj any, t *testing.T) {
  t.Helper()
  path, _ := os.Getwd()
  file, err := os.Open(path + input)
  assert.NoError(t, err)
  defer file.Close()
  dec := yaml.NewDecoder(file)
  err = dec.Decode(obj)
  assert.NoError(t, err)
}
