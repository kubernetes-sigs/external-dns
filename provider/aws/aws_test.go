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

package aws

import (
	"context"
	"fmt"
	"math"
	"net"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	defaultBatchChangeSize       = 4000
	defaultBatchChangeSizeBytes  = 32000
	defaultBatchChangeSizeValues = 1000
	defaultBatchChangeInterval   = time.Second
	defaultEvaluateTargetHealth  = true
)

// Compile time check for interface conformance
var _ Route53API = &Route53APIStub{}

// Route53APIStub is a minimal implementation of Route53API, used primarily for unit testing.
// See http://http://docs.aws.amazon.com/sdk-for-go/api/service/route53.html for descriptions
// of all of its methods.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53APIStub struct {
	zones      map[string]*route53types.HostedZone
	recordSets map[string]map[string][]route53types.ResourceRecordSet
	zoneTags   map[string][]route53types.Tag
	m          dynamicMock
	t          *testing.T
}

// MockMethod starts a description of an expectation of the specified method
// being called.
//
//	Route53APIStub.MockMethod("MyMethod", arg1, arg2)
func (r *Route53APIStub) MockMethod(method string, args ...interface{}) *mock.Call {
	return r.m.On(method, args...)
}

// NewRoute53APIStub returns an initialized Route53APIStub
func NewRoute53APIStub(t *testing.T) *Route53APIStub {
	return &Route53APIStub{
		zones:      make(map[string]*route53types.HostedZone),
		recordSets: make(map[string]map[string][]route53types.ResourceRecordSet),
		zoneTags:   make(map[string][]route53types.Tag),
		t:          t,
	}
}

func (r *Route53APIStub) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
	if r.m.isMocked("ListResourceRecordSets", input) {
		return r.m.ListResourceRecordSets(ctx, input, optFns...)
	}

	output := &route53.ListResourceRecordSetsOutput{} // TODO: Support optional input args.
	require.NotNil(r.t, input.MaxItems)
	assert.EqualValues(r.t, route53PageSize, *input.MaxItems)
	if len(r.recordSets) == 0 {
		output.ResourceRecordSets = []route53types.ResourceRecordSet{}
	} else if _, ok := r.recordSets[*input.HostedZoneId]; !ok {
		output.ResourceRecordSets = []route53types.ResourceRecordSet{}
	} else {
		for _, rrsets := range r.recordSets[*input.HostedZoneId] {
			output.ResourceRecordSets = append(output.ResourceRecordSets, rrsets...)
		}
	}
	return output, nil
}

type Route53APICounter struct {
	wrapped Route53API
	calls   map[string]int
}

func NewRoute53APICounter(w Route53API) *Route53APICounter {
	return &Route53APICounter{
		wrapped: w,
		calls:   map[string]int{},
	}
}

func (c *Route53APICounter) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
	c.calls["ListResourceRecordSetsPages"]++
	return c.wrapped.ListResourceRecordSets(ctx, input, optFns...)
}

func (c *Route53APICounter) ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error) {
	c.calls["ChangeResourceRecordSets"]++
	return c.wrapped.ChangeResourceRecordSets(ctx, input, optFns...)
}

func (c *Route53APICounter) CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput, optFns ...func(*route53.Options)) (*route53.CreateHostedZoneOutput, error) {
	c.calls["CreateHostedZone"]++
	return c.wrapped.CreateHostedZone(ctx, input, optFns...)
}

func (c *Route53APICounter) ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput, optFns ...func(options *route53.Options)) (*route53.ListHostedZonesOutput, error) {
	c.calls["ListHostedZonesPages"]++
	return c.wrapped.ListHostedZones(ctx, input, optFns...)
}

func (c *Route53APICounter) ListTagsForResource(ctx context.Context, input *route53.ListTagsForResourceInput, optFns ...func(options *route53.Options)) (*route53.ListTagsForResourceOutput, error) {
	c.calls["ListTagsForResource"]++
	return c.wrapped.ListTagsForResource(ctx, input, optFns...)
}

// Route53 stores wildcards escaped: http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html?shortFooter=true#domain-name-format-asterisk
func wildcardEscape(s string) string {
	if strings.Contains(s, "*") {
		s = strings.Replace(s, "*", "\\052", 1)
	}
	return s
}

// Route53 octal escapes https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html
func specialCharactersEscape(s string) string {
	var result strings.Builder
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' || char == '.' {
			result.WriteRune(char)
		} else {
			octalCode := fmt.Sprintf("\\%03o", char)
			result.WriteString(octalCode)
		}
	}
	return result.String()
}

func (r *Route53APIStub) ListTagsForResource(ctx context.Context, input *route53.ListTagsForResourceInput, optFns ...func(options *route53.Options)) (*route53.ListTagsForResourceOutput, error) {
	if input.ResourceType == route53types.TagResourceTypeHostedzone {
		tags := r.zoneTags[*input.ResourceId]
		return &route53.ListTagsForResourceOutput{
			ResourceTagSet: &route53types.ResourceTagSet{
				ResourceId:   input.ResourceId,
				ResourceType: input.ResourceType,
				Tags:         tags,
			},
		}, nil
	}
	return &route53.ListTagsForResourceOutput{}, nil
}

func (r *Route53APIStub) ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error) {
	if r.m.isMocked("ChangeResourceRecordSets", input) {
		return r.m.ChangeResourceRecordSets(input)
	}

	_, ok := r.zones[*input.HostedZoneId]
	if !ok {
		return nil, fmt.Errorf("Hosted zone doesn't exist: %s", *input.HostedZoneId)
	}

	if len(input.ChangeBatch.Changes) == 0 {
		return nil, fmt.Errorf("ChangeBatch doesn't contain any changes")
	}

	output := &route53.ChangeResourceRecordSetsOutput{}
	recordSets, ok := r.recordSets[*input.HostedZoneId]
	if !ok {
		recordSets = make(map[string][]route53types.ResourceRecordSet)
	}

	for _, change := range input.ChangeBatch.Changes {
		if change.ResourceRecordSet.Type == route53types.RRTypeA {
			for _, rrs := range change.ResourceRecordSet.ResourceRecords {
				if net.ParseIP(*rrs.Value) == nil {
					return nil, fmt.Errorf("A records must point to IPs")
				}
			}
		}

		change.ResourceRecordSet.Name = aws.String(wildcardEscape(provider.EnsureTrailingDot(*change.ResourceRecordSet.Name)))

		if change.ResourceRecordSet.AliasTarget != nil {
			change.ResourceRecordSet.AliasTarget.DNSName = aws.String(wildcardEscape(provider.EnsureTrailingDot(*change.ResourceRecordSet.AliasTarget.DNSName)))
		}

		setID := ""
		if change.ResourceRecordSet.SetIdentifier != nil {
			setID = *change.ResourceRecordSet.SetIdentifier
		}
		key := *change.ResourceRecordSet.Name + "::" + string(change.ResourceRecordSet.Type) + "::" + setID
		switch change.Action {
		case route53types.ChangeActionCreate:
			if _, found := recordSets[key]; found {
				return nil, fmt.Errorf("Attempt to create duplicate rrset %s", key) // TODO: Return AWS errors with codes etc
			}
			recordSets[key] = append(recordSets[key], *change.ResourceRecordSet)
		case route53types.ChangeActionDelete:
			if _, found := recordSets[key]; !found {
				return nil, fmt.Errorf("Attempt to delete non-existent rrset %s", key) // TODO: Check other fields too
			}
			delete(recordSets, key)
		case route53types.ChangeActionUpsert:
			recordSets[key] = []route53types.ResourceRecordSet{*change.ResourceRecordSet}
		}
	}
	r.recordSets[*input.HostedZoneId] = recordSets
	return output, nil // TODO: We should ideally return status etc, but we don't' use that yet.
}

func (r *Route53APIStub) ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput, optFns ...func(options *route53.Options)) (*route53.ListHostedZonesOutput, error) {
	output := &route53.ListHostedZonesOutput{}
	for _, zone := range r.zones {
		output.HostedZones = append(output.HostedZones, *zone)
	}
	return output, nil
}

func (r *Route53APIStub) CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput, optFns ...func(options *route53.Options)) (*route53.CreateHostedZoneOutput, error) {
	name := *input.Name
	id := "/hostedzone/" + name
	if _, ok := r.zones[id]; ok {
		return nil, fmt.Errorf("Error creating hosted DNS zone: %s already exists", id)
	}
	r.zones[id] = &route53types.HostedZone{
		Id:     aws.String(id),
		Name:   aws.String(name),
		Config: input.HostedZoneConfig,
	}
	return &route53.CreateHostedZoneOutput{HostedZone: r.zones[id]}, nil
}

type dynamicMock struct {
	mock.Mock
}

func (m *dynamicMock) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ListResourceRecordSetsOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*route53.ListResourceRecordSetsOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *dynamicMock) ChangeResourceRecordSets(input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*route53.ChangeResourceRecordSetsOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *dynamicMock) isMocked(method string, arguments ...interface{}) bool {
	for _, call := range m.ExpectedCalls {
		if call.Method == method && call.Repeatability > -1 {
			_, diffCount := call.Arguments.Diff(arguments)
			if diffCount == 0 {
				return true
			}
		}
	}
	return false
}

func TestAWSZones(t *testing.T) {
	publicZones := map[string]*route53types.HostedZone{
		"/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do.": {
			Id:   aws.String("/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."),
			Name: aws.String("zone-1.ext-dns-test-2.teapot.zalan.do."),
		},
		"/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do.": {
			Id:   aws.String("/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do."),
			Name: aws.String("zone-2.ext-dns-test-2.teapot.zalan.do."),
		},
	}

	privateZones := map[string]*route53types.HostedZone{
		"/hostedzone/zone-3.ext-dns-test-2.teapot.zalan.do.": {
			Id:   aws.String("/hostedzone/zone-3.ext-dns-test-2.teapot.zalan.do."),
			Name: aws.String("zone-3.ext-dns-test-2.teapot.zalan.do."),
		},
	}

	allZones := map[string]*route53types.HostedZone{}
	for k, v := range publicZones {
		allZones[k] = v
	}
	for k, v := range privateZones {
		allZones[k] = v
	}

	noZones := map[string]*route53types.HostedZone{}

	for _, ti := range []struct {
		msg            string
		zoneIDFilter   provider.ZoneIDFilter
		zoneTypeFilter provider.ZoneTypeFilter
		zoneTagFilter  provider.ZoneTagFilter
		expectedZones  map[string]*route53types.HostedZone
	}{
		{"no filter", provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), provider.NewZoneTagFilter([]string{}), allZones},
		{"public filter", provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter("public"), provider.NewZoneTagFilter([]string{}), publicZones},
		{"private filter", provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter("private"), provider.NewZoneTagFilter([]string{}), privateZones},
		{"unknown filter", provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter("unknown"), provider.NewZoneTagFilter([]string{}), noZones},
		{"zone id filter", provider.NewZoneIDFilter([]string{"/hostedzone/zone-3.ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneTypeFilter(""), provider.NewZoneTagFilter([]string{}), privateZones},
		{"tag filter", provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), provider.NewZoneTagFilter([]string{"zone=3"}), privateZones},
	} {
		t.Run(ti.msg, func(t *testing.T) {
			provider, _ := newAWSProviderWithTagFilter(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), ti.zoneIDFilter, ti.zoneTypeFilter, ti.zoneTagFilter, defaultEvaluateTargetHealth, false, nil)

			zones, err := provider.Zones(context.Background())
			require.NoError(t, err)

			validateAWSZones(t, zones, ti.expectedZones)
		})
	}
}

func TestAWSRecordsFilter(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.DomainFilter{}, provider.ZoneIDFilter{}, provider.ZoneTypeFilter{}, false, false, nil)
	domainFilter := provider.GetDomainFilter()
	require.NotNil(t, domainFilter)
	require.IsType(t, endpoint.DomainFilter{}, domainFilter)
	count := 0
	filters := domainFilter.(endpoint.DomainFilter).Filters
	for _, tld := range []string{
		"zone-4.ext-dns-test-3.teapot.zalan.do",
		".zone-4.ext-dns-test-3.teapot.zalan.do",
		"zone-2.ext-dns-test-2.teapot.zalan.do",
		".zone-2.ext-dns-test-2.teapot.zalan.do",
		"zone-3.ext-dns-test-2.teapot.zalan.do",
		".zone-3.ext-dns-test-2.teapot.zalan.do",
		"zone-4.ext-dns-test-3.teapot.zalan.do",
		".zone-4.ext-dns-test-3.teapot.zalan.do",
	} {
		assert.Contains(t, filters, tld)
		count++
	}
	assert.Len(t, filters, count)
}

func TestAWSRecords(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), false, false, []route53types.ResourceRecordSet{
		{
			Name:            aws.String("list-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
		},
		{
			Name:            aws.String("list-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
		},
		{
			Name:            aws.String(wildcardEscape("*.wildcard-test.zone-2.ext-dns-test-2.teapot.zalan.do.")),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
		},
		{
			Name:            aws.String(specialCharactersEscape("escape-%!s(<nil>)-codes.zone-2.ext-dns-test-2.teapot.zalan.do.")),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("example")}},
		},
		{
			Name:            aws.String(specialCharactersEscape("escape-%!s(<nil>)-codes-a.zone-2.ext-dns-test-2.teapot.zalan.do.")),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
		},
		{
			Name: aws.String(specialCharactersEscape("escape-%!s(<nil>)-codes-alias.zone-2.ext-dns-test-2.teapot.zalan.do.")),
			Type: route53types.RRTypeA,
			TTL:  aws.Int64(recordTTL),
			AliasTarget: &route53types.AliasTarget{
				DNSName:              aws.String("escape-codes.eu-central-1.elb.amazonaws.com."),
				EvaluateTargetHealth: false,
				HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
			},
		},
		{
			Name: aws.String("list-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type: route53types.RRTypeA,
			AliasTarget: &route53types.AliasTarget{
				DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
				EvaluateTargetHealth: false,
				HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
			},
		},
		{
			Name: aws.String("*.wildcard-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type: route53types.RRTypeA,
			AliasTarget: &route53types.AliasTarget{
				DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
				EvaluateTargetHealth: false,
				HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
			},
		},
		{
			Name: aws.String("list-test-alias-evaluate.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type: route53types.RRTypeA,
			AliasTarget: &route53types.AliasTarget{
				DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
				EvaluateTargetHealth: true,
				HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
			},
		},
		{
			Name:            aws.String("list-test-multiple.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}, {Value: aws.String("8.8.4.4")}},
		},
		{
			Name:            aws.String("prefix-*.wildcard.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeTxt,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("random")}},
		},
		{
			Name:            aws.String("weight-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:   aws.String("test-set-1"),
			Weight:          aws.Int64(10),
		},
		{
			Name:            aws.String("weight-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("4.3.2.1")}},
			SetIdentifier:   aws.String("test-set-2"),
			Weight:          aws.Int64(20),
		},
		{
			Name:            aws.String("latency-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:   aws.String("test-set"),
			Region:          route53types.ResourceRecordSetRegionUsEast1,
		},
		{
			Name:            aws.String("failover-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:   aws.String("test-set"),
			Failover:        route53types.ResourceRecordSetFailoverPrimary,
		},
		{
			Name:             aws.String("multi-value-answer-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:             route53types.RRTypeA,
			TTL:              aws.Int64(recordTTL),
			ResourceRecords:  []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:    aws.String("test-set"),
			MultiValueAnswer: aws.Bool(true),
		},
		{
			Name:            aws.String("geolocation-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:   aws.String("test-set-1"),
			GeoLocation: &route53types.GeoLocation{
				ContinentCode: aws.String("EU"),
			},
		},
		{
			Name:            aws.String("geolocation-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("4.3.2.1")}},
			SetIdentifier:   aws.String("test-set-2"),
			GeoLocation: &route53types.GeoLocation{
				CountryCode: aws.String("DE"),
			},
		},
		{
			Name:            aws.String("geolocation-subdivision-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			SetIdentifier:   aws.String("test-set-1"),
			GeoLocation: &route53types.GeoLocation{
				SubdivisionCode: aws.String("NY"),
			},
		},
		{
			Name:            aws.String("healthcheck-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("foo.example.com")}},
			SetIdentifier:   aws.String("test-set-1"),
			HealthCheckId:   aws.String("foo-bar-healthcheck-id"),
			Weight:          aws.Int64(10),
		},
		{
			Name:            aws.String("healthcheck-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("4.3.2.1")}},
			SetIdentifier:   aws.String("test-set-2"),
			HealthCheckId:   aws.String("abc-def-healthcheck-id"),
			Weight:          aws.Int64(20),
		},
		{
			Name:            aws.String("mail.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeMx,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("10 mailhost1.example.com")}, {Value: aws.String("20 mailhost2.example.com")}},
		},
	})

	records, err := provider.Records(context.Background())
	require.NoError(t, err)

	validateEndpoints(t, provider, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("list-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("list-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("*.wildcard-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("escape-%!s(<nil>)-codes.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "example").WithProviderSpecific(providerSpecificAlias, "false"),
		endpoint.NewEndpointWithTTL("escape-%!s(<nil>)-codes-a.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("escape-%!s(<nil>)-codes-alias.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "escape-codes.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false").WithProviderSpecific(providerSpecificAlias, "true"),
		endpoint.NewEndpointWithTTL("list-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false").WithProviderSpecific(providerSpecificAlias, "true"),
		endpoint.NewEndpointWithTTL("*.wildcard-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false").WithProviderSpecific(providerSpecificAlias, "true"),
		endpoint.NewEndpointWithTTL("list-test-alias-evaluate.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "true").WithProviderSpecific(providerSpecificAlias, "true"),
		endpoint.NewEndpointWithTTL("list-test-multiple.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpointWithTTL("prefix-*.wildcard.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "random"),
		endpoint.NewEndpointWithTTL("weight-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set-1").WithProviderSpecific(providerSpecificWeight, "10"),
		endpoint.NewEndpointWithTTL("weight-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1").WithSetIdentifier("test-set-2").WithProviderSpecific(providerSpecificWeight, "20"),
		endpoint.NewEndpointWithTTL("latency-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set").WithProviderSpecific(providerSpecificRegion, "us-east-1"),
		endpoint.NewEndpointWithTTL("failover-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set").WithProviderSpecific(providerSpecificFailover, "PRIMARY"),
		endpoint.NewEndpointWithTTL("multi-value-answer-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set").WithProviderSpecific(providerSpecificMultiValueAnswer, ""),
		endpoint.NewEndpointWithTTL("geolocation-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set-1").WithProviderSpecific(providerSpecificGeolocationContinentCode, "EU"),
		endpoint.NewEndpointWithTTL("geolocation-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1").WithSetIdentifier("test-set-2").WithProviderSpecific(providerSpecificGeolocationCountryCode, "DE"),
		endpoint.NewEndpointWithTTL("geolocation-subdivision-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4").WithSetIdentifier("test-set-1").WithProviderSpecific(providerSpecificGeolocationSubdivisionCode, "NY"),
		endpoint.NewEndpointWithTTL("healthcheck-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.example.com").WithSetIdentifier("test-set-1").WithProviderSpecific(providerSpecificWeight, "10").WithProviderSpecific(providerSpecificHealthCheckID, "foo-bar-healthcheck-id").WithProviderSpecific(providerSpecificAlias, "false"),
		endpoint.NewEndpointWithTTL("healthcheck-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1").WithSetIdentifier("test-set-2").WithProviderSpecific(providerSpecificWeight, "20").WithProviderSpecific(providerSpecificHealthCheckID, "abc-def-healthcheck-id"),
		endpoint.NewEndpointWithTTL("mail.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, endpoint.TTL(recordTTL), "10 mailhost1.example.com", "20 mailhost2.example.com"),
	})
}

func TestAWSRecordsSoftError(t *testing.T) {
	pvd, subClient := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), false, false, []route53types.ResourceRecordSet{
		{
			Name:            aws.String("list-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
		},
	})

	subClient.MockMethod("ListResourceRecordSets", mock.Anything).Return(nil, fmt.Errorf("Mock route53 failure"))
	_, err := pvd.Records(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
}

func TestAWSAdjustEndpoints(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)

	records := []*endpoint.Endpoint{
		endpoint.NewEndpoint("a-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("cname-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.example.com"),
		endpoint.NewEndpointWithTTL("cname-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, 60, "alias-target.zone-2.ext-dns-test-2.teapot.zalan.do").WithProviderSpecific(providerSpecificAlias, "true"),
		endpoint.NewEndpoint("cname-test-elb.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpoint("cname-test-elb-no-alias.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "false"),
		endpoint.NewEndpoint("cname-test-elb-no-eth.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false"), // eth = evaluate target health
		endpoint.NewEndpoint("cname-test-elb-alias.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "true"),
	}

	records, err := provider.AdjustEndpoints(records)
	assert.NoError(t, err)

	validateEndpoints(t, provider, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("a-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("cname-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.example.com").WithProviderSpecific(providerSpecificAlias, "false"),
		endpoint.NewEndpointWithTTL("cname-test-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, 300, "alias-target.zone-2.ext-dns-test-2.teapot.zalan.do").WithProviderSpecific(providerSpecificAlias, "true").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "true"),
		endpoint.NewEndpoint("cname-test-elb.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "true"),
		endpoint.NewEndpoint("cname-test-elb-no-alias.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "false"),
		endpoint.NewEndpoint("cname-test-elb-no-eth.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "false"), // eth = evaluate target health
		endpoint.NewEndpoint("cname-test-elb-alias.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true").WithProviderSpecific(providerSpecificEvaluateTargetHealth, "true"),
	})
}

func TestAWSApplyChanges(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(p *AWSProvider) context.Context
		listRRSets int
	}{
		{"no cache", func(p *AWSProvider) context.Context { return context.Background() }, 0},
		{"cached", func(p *AWSProvider) context.Context {
			ctx := context.Background()
			records, err := p.Records(ctx)
			require.NoError(t, err)
			return context.WithValue(ctx, provider.RecordsContextKey, records)
		}, 0},
	}

	for _, tt := range tests {
		provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, []route53types.ResourceRecordSet{
			{
				Name:            aws.String("update-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
			},
			{
				Name:            aws.String("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
			},
			{
				Name:            aws.String("update-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.4.4")}},
			},
			{
				Name:            aws.String("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.4.4")}},
			},
			{
				Name:            aws.String("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.1.1.1")}},
			},
			{
				Name: aws.String("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type: route53types.RRTypeA,
				AliasTarget: &route53types.AliasTarget{
					DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
					EvaluateTargetHealth: true,
					HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
				},
			},
			{
				Name:            aws.String("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("bar.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("qux.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("bar.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("qux.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}, {Value: aws.String("8.8.4.4")}},
			},
			{
				Name:            aws.String("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}, {Value: aws.String("4.3.2.1")}},
			},
			{
				Name:            aws.String("weighted-to-simple.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("weighted-to-simple"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("simple-to-weighted.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			},
			{
				Name:            aws.String("policy-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("policy-change"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("set-identifier-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("before"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("set-identifier-no-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("no-change"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("update-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeMx,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("10 mailhost2.bar.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("delete-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeMx,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("30 mailhost1.foo.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String(specialCharactersEscape("escape-%!s(<nil>)-codes.zone-2.ext-dns-test-2.teapot.zalan.do.")),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("no-change"),
				Weight:          aws.Int64(10),
			},
		})

		createRecords := []*endpoint.Endpoint{
			endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
			endpoint.NewEndpoint("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
			endpoint.NewEndpoint("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
			endpoint.NewEndpoint("create-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "10 mailhost1.foo.elb.amazonaws.com"),
		}

		currentRecords := []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.1.1.1"),
			endpoint.NewEndpoint("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "foo.eu-central-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true"),
			endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
			endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
			endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
			endpoint.NewEndpoint("weighted-to-simple.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("weighted-to-simple").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("simple-to-weighted.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("policy-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("policy-change").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("set-identifier-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("before").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("set-identifier-no-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("no-change").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("update-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "10 mailhost2.bar.elb.amazonaws.com"),
			endpoint.NewEndpoint("escape-%!s(<nil>)-codes.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("policy-change").WithSetIdentifier("no-change").WithProviderSpecific(providerSpecificWeight, "10"),
		}
		updatedRecords := []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "4.3.2.1"),
			endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "foo.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true"),
			endpoint.NewEndpoint("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "my-internal-host.example.com"),
			endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
			endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
			endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
			endpoint.NewEndpoint("weighted-to-simple.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("simple-to-weighted.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("simple-to-weighted").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("policy-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("policy-change").WithProviderSpecific(providerSpecificRegion, "us-east-1"),
			endpoint.NewEndpoint("set-identifier-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("after").WithProviderSpecific(providerSpecificWeight, "10"),
			endpoint.NewEndpoint("set-identifier-no-change.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4").WithSetIdentifier("no-change").WithProviderSpecific(providerSpecificWeight, "20"),
			endpoint.NewEndpoint("update-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "20 mailhost3.foo.elb.amazonaws.com"),
		}

		deleteRecords := []*endpoint.Endpoint{
			endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
			endpoint.NewEndpoint("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
			endpoint.NewEndpoint("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
			endpoint.NewEndpoint("delete-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "30 mailhost1.foo.elb.amazonaws.com"),
		}

		changes := &plan.Changes{
			Create:    createRecords,
			UpdateNew: updatedRecords,
			UpdateOld: currentRecords,
			Delete:    deleteRecords,
		}

		ctx := tt.setup(provider)

		provider.zonesCache = &zonesListCache{duration: 0 * time.Minute}
		counter := NewRoute53APICounter(provider.clients[defaultAWSProfile])
		provider.clients[defaultAWSProfile] = counter
		require.NoError(t, provider.ApplyChanges(ctx, changes))

		assert.Equal(t, 1, counter.calls["ListHostedZonesPages"], tt.name)
		assert.Equal(t, tt.listRRSets, counter.calls["ListResourceRecordSetsPages"], tt.name)

		validateRecords(t, listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."), []route53types.ResourceRecordSet{
			{
				Name:            aws.String("create-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
			},
			{
				Name:            aws.String("update-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			},
			{
				Name: aws.String("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type: route53types.RRTypeA,
				AliasTarget: &route53types.AliasTarget{
					DNSName:              aws.String("foo.elb.amazonaws.com."),
					EvaluateTargetHealth: true,
					HostedZoneId:         aws.String("zone-1.ext-dns-test-2.teapot.zalan.do."),
				},
			},
			{
				Name:            aws.String("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("my-internal-host.example.com")}},
			},
			{
				Name:            aws.String("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("foo.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("baz.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("foo.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeCname,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("baz.elb.amazonaws.com")}},
			},
			{
				Name:            aws.String("weighted-to-simple.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
			},
			{
				Name:            aws.String("simple-to-weighted.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("simple-to-weighted"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("policy-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("policy-change"),
				Region:          route53types.ResourceRecordSetRegionUsEast1,
			},
			{
				Name:            aws.String("set-identifier-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("after"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("set-identifier-no-change.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("no-change"),
				Weight:          aws.Int64(20),
			},
			{
				Name:            aws.String("create-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeMx,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("10 mailhost1.foo.elb.amazonaws.com")}},
			},
		})
		validateRecords(t, listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do."), []route53types.ResourceRecordSet{
			{
				Name:            aws.String("escape-\\045\\041s\\050\\074nil\\076\\051-codes.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}},
				SetIdentifier:   aws.String("no-change"),
				Weight:          aws.Int64(10),
			},
			{
				Name:            aws.String("create-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.4.4")}},
			},
			{
				Name:            aws.String("update-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("4.3.2.1")}},
			},
			{
				Name:            aws.String("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}, {Value: aws.String("8.8.4.4")}},
			},
			{
				Name:            aws.String("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeA,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}, {Value: aws.String("4.3.2.1")}},
			},
			{
				Name:            aws.String("update-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do."),
				Type:            route53types.RRTypeMx,
				TTL:             aws.Int64(recordTTL),
				ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("20 mailhost3.foo.elb.amazonaws.com")}},
			},
		})
	}
}

func TestAWSApplyChangesDryRun(t *testing.T) {
	originalRecords := []route53types.ResourceRecordSet{
		{
			Name:            aws.String("update-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
		},
		{
			Name:            aws.String("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}},
		},
		{
			Name:            aws.String("update-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.4.4")}},
		},
		{
			Name:            aws.String("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.4.4")}},
		},
		{
			Name:            aws.String("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.1.1.1")}},
		},
		{
			Name:            aws.String("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("bar.elb.amazonaws.com")}},
		},
		{
			Name:            aws.String("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("qux.elb.amazonaws.com")}},
		},
		{
			Name:            aws.String("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("bar.elb.amazonaws.com")}},
		},
		{
			Name:            aws.String("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeCname,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("qux.elb.amazonaws.com")}},
		},
		{
			Name:            aws.String("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("8.8.8.8")}, {Value: aws.String("8.8.4.4")}},
		},
		{
			Name:            aws.String("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeA,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("1.2.3.4")}, {Value: aws.String("4.3.2.1")}},
		},
		{
			Name:            aws.String("update-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeMx,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("20 mail.foo.elb.amazonaws.com")}},
		},
		{
			Name:            aws.String("delete-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do."),
			Type:            route53types.RRTypeMx,
			TTL:             aws.Int64(recordTTL),
			ResourceRecords: []route53types.ResourceRecord{{Value: aws.String("10 mail.bar.elb.amazonaws.com")}},
		},
	}

	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, true, originalRecords)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
		endpoint.NewEndpoint("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
		endpoint.NewEndpoint("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpoint("create-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "30 mail.foo.elb.amazonaws.com"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.1.1.1"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpoint("update-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "20 mail.foo.elb.amazonaws.com"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "4.3.2.1"),
		endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
		endpoint.NewEndpoint("update-test-mx.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "10 mail.bar.elb.amazonaws.com"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
		endpoint.NewEndpoint("delete-test-mx.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeMX, "10 mail.bar.elb.amazonaws.com"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	ctx := context.Background()

	require.NoError(t, provider.ApplyChanges(ctx, changes))

	validateRecords(t,
		append(
			listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."),
			listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do.")...),
		originalRecords)
}

func TestAWSChangesByZones(t *testing.T) {
	changes := Route53Changes{
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("qux.foo.example.org"), TTL: aws.Int64(1),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("qux.bar.example.org"), TTL: aws.Int64(2),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionDelete,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("wambo.foo.example.org"), TTL: aws.Int64(10),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionDelete,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("wambo.bar.example.org"), TTL: aws.Int64(20),
				},
			},
		},
	}

	zones := map[string]*profiledZone{
		"foo-example-org": {
			profile: defaultAWSProfile,
			zone: &route53types.HostedZone{
				Id:   aws.String("foo-example-org"),
				Name: aws.String("foo.example.org."),
			},
		},
		"bar-example-org": {
			profile: defaultAWSProfile,
			zone: &route53types.HostedZone{
				Id:   aws.String("bar-example-org"),
				Name: aws.String("bar.example.org."),
			},
		},
		"bar-example-org-private": {
			profile: defaultAWSProfile,
			zone: &route53types.HostedZone{
				Id:     aws.String("bar-example-org-private"),
				Name:   aws.String("bar.example.org."),
				Config: &route53types.HostedZoneConfig{PrivateZone: true},
			},
		},
		"baz-example-org": {
			profile: defaultAWSProfile,
			zone: &route53types.HostedZone{
				Id:   aws.String("baz-example-org"),
				Name: aws.String("baz.example.org."),
			},
		},
	}

	changesByZone := changesByZone(zones, changes)
	require.Len(t, changesByZone, 3)

	validateAWSChangeRecords(t, changesByZone["foo-example-org"], Route53Changes{
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("qux.foo.example.org"), TTL: aws.Int64(1),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionDelete,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("wambo.foo.example.org"), TTL: aws.Int64(10),
				},
			},
		},
	})

	validateAWSChangeRecords(t, changesByZone["bar-example-org"], Route53Changes{
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("qux.bar.example.org"), TTL: aws.Int64(2),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionDelete,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("wambo.bar.example.org"), TTL: aws.Int64(20),
				},
			},
		},
	})

	validateAWSChangeRecords(t, changesByZone["bar-example-org-private"], Route53Changes{
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("qux.bar.example.org"), TTL: aws.Int64(2),
				},
			},
		},
		{
			Change: route53types.Change{
				Action: route53types.ChangeActionDelete,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String("wambo.bar.example.org"), TTL: aws.Int64(20),
				},
			},
		},
	})
}

func TestAWSsubmitChanges(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)
	const subnets = 16
	const hosts = defaultBatchChangeSize / subnets

	endpoints := make([]*endpoint.Endpoint, 0)
	for i := 0; i < subnets; i++ {
		for j := 1; j < (hosts + 1); j++ {
			hostname := fmt.Sprintf("subnet%dhost%d.zone-1.ext-dns-test-2.teapot.zalan.do", i, j)
			ip := fmt.Sprintf("1.1.%d.%d", i, j)
			ep := endpoint.NewEndpointWithTTL(hostname, endpoint.RecordTypeA, endpoint.TTL(recordTTL), ip)
			endpoints = append(endpoints, ep)
		}
	}

	ctx := context.Background()
	zones, _ := provider.zones(ctx)
	records, _ := provider.Records(ctx)
	cs := make(Route53Changes, 0, len(endpoints))
	cs = append(cs, provider.newChanges(route53types.ChangeActionCreate, endpoints)...)

	require.NoError(t, provider.submitChanges(ctx, cs, zones))

	records, err := provider.Records(ctx)
	require.NoError(t, err)

	validateEndpoints(t, provider, records, endpoints)
}

func TestAWSsubmitChangesError(t *testing.T) {
	provider, clientStub := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)
	clientStub.MockMethod("ChangeResourceRecordSets", mock.Anything).Return(nil, fmt.Errorf("Mock route53 failure"))

	ctx := context.Background()
	zones, err := provider.zones(ctx)
	require.NoError(t, err)

	ep := endpoint.NewEndpointWithTTL("fail.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.0.0.1")
	cs := provider.newChanges(route53types.ChangeActionCreate, []*endpoint.Endpoint{ep})

	require.Error(t, provider.submitChanges(ctx, cs, zones))
}

func TestAWSsubmitChangesRetryOnError(t *testing.T) {
	provider, clientStub := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)

	ctx := context.Background()
	zones, err := provider.zones(ctx)
	require.NoError(t, err)

	ep1 := endpoint.NewEndpointWithTTL("success.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.0.0.1")
	ep2 := endpoint.NewEndpointWithTTL("fail.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.0.0.2")
	ep3 := endpoint.NewEndpointWithTTL("success2.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.0.0.3")

	ep2txt := endpoint.NewEndpointWithTTL("fail__edns_housekeeping.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "something") // "__edns_housekeeping" is the TXT suffix
	ep2txt.Labels = map[string]string{
		endpoint.OwnedRecordLabelKey: "fail.zone-1.ext-dns-test-2.teapot.zalan.do",
	}

	// "success" and "fail" are created in the first step, both are submitted in the same batch; this should fail
	cs1 := provider.newChanges(route53types.ChangeActionCreate, []*endpoint.Endpoint{ep2, ep2txt, ep1})
	input1 := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String("/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."),
		ChangeBatch: &route53types.ChangeBatch{
			Changes: cs1.Route53Changes(),
		},
	}
	clientStub.MockMethod("ChangeResourceRecordSets", input1).Return(nil, fmt.Errorf("Mock route53 failure"))

	// because of the failure, changes will be retried one by one; make "fail" submitted in its own batch fail as well
	cs2 := provider.newChanges(route53types.ChangeActionCreate, []*endpoint.Endpoint{ep2, ep2txt})
	input2 := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String("/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."),
		ChangeBatch: &route53types.ChangeBatch{
			Changes: cs2.Route53Changes(),
		},
	}
	clientStub.MockMethod("ChangeResourceRecordSets", input2).Return(nil, fmt.Errorf("Mock route53 failure"))

	// "success" should have been created, verify that we still get an error because "fail" failed
	require.Error(t, provider.submitChanges(ctx, cs1, zones))

	// assert that "success" was successfully created and "fail" and its TXT record were not
	records, err := provider.Records(ctx)
	require.NoError(t, err)
	require.True(t, containsRecordWithDNSName(records, "success.zone-1.ext-dns-test-2.teapot.zalan.do"))
	require.False(t, containsRecordWithDNSName(records, "fail.zone-1.ext-dns-test-2.teapot.zalan.do"))
	require.False(t, containsRecordWithDNSName(records, "fail__edns_housekeeping.zone-1.ext-dns-test-2.teapot.zalan.do"))

	// next batch should contain "fail" and "success2", should succeed this time
	cs3 := provider.newChanges(route53types.ChangeActionCreate, []*endpoint.Endpoint{ep2, ep2txt, ep3})
	require.NoError(t, provider.submitChanges(ctx, cs3, zones))

	// verify all records are there
	records, err = provider.Records(ctx)
	require.NoError(t, err)
	require.True(t, containsRecordWithDNSName(records, "success.zone-1.ext-dns-test-2.teapot.zalan.do"))
	require.True(t, containsRecordWithDNSName(records, "fail.zone-1.ext-dns-test-2.teapot.zalan.do"))
	require.True(t, containsRecordWithDNSName(records, "success2.zone-1.ext-dns-test-2.teapot.zalan.do"))
	require.True(t, containsRecordWithDNSName(records, "fail__edns_housekeeping.zone-1.ext-dns-test-2.teapot.zalan.do"))
}

func TestAWSBatchChangeSet(t *testing.T) {
	var cs Route53Changes

	for i := 1; i <= defaultBatchChangeSize; i += 2 {
		cs = append(cs, &Route53Change{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String(fmt.Sprintf("host-%d", i)),
					Type: route53types.RRTypeA,
				},
			},
		})
		cs = append(cs, &Route53Change{
			Change: route53types.Change{
				Action: route53types.ChangeActionCreate,
				ResourceRecordSet: &route53types.ResourceRecordSet{
					Name: aws.String(fmt.Sprintf("host-%d", i)),
					Type: route53types.RRTypeTxt,
				},
			},
		})
	}

	batchCs := batchChangeSet(cs, defaultBatchChangeSize, defaultBatchChangeSizeBytes, defaultBatchChangeSizeValues)

	require.Equal(t, 1, len(batchCs))

	// sorting cs not needed as it should be returned as is
	validateAWSChangeRecords(t, batchCs[0], cs)
}

func TestAWSBatchChangeSetExceeding(t *testing.T) {
	var cs Route53Changes
	const testCount = 50
	const testLimit = 11
	const expectedBatchCount = 5
	const expectedChangesCount = 10

	for i := 1; i <= testCount; i += 2 {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
					},
				},
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
					},
				},
			},
		)
	}

	batchCs := batchChangeSet(cs, testLimit, defaultBatchChangeSizeBytes, defaultBatchChangeSizeValues)

	require.Equal(t, expectedBatchCount, len(batchCs))

	// sorting cs needed to match batchCs
	for i, batch := range batchCs {
		validateAWSChangeRecords(t, batch, sortChangesByActionNameType(cs)[i*expectedChangesCount:expectedChangesCount*(i+1)])
	}
}

func TestAWSBatchChangeSetExceedingNameChange(t *testing.T) {
	var cs Route53Changes
	const testCount = 10
	const testLimit = 1

	for i := 1; i <= testCount; i += 2 {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
					},
				},
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
					},
				},
			},
		)
	}

	batchCs := batchChangeSet(cs, testLimit, defaultBatchChangeSizeBytes, defaultBatchChangeSizeValues)

	require.Equal(t, 0, len(batchCs))
}

func TestAWSBatchChangeSetExceedingBytesLimit(t *testing.T) {
	const (
		testCount = 50
		testLimit = 100
		groupSize = 2
	)

	var (
		cs Route53Changes
		// Bytes for each name
		testBytes = len([]byte("1.2.3.4")) + len([]byte("test-record"))
		// testCount / groupSize / (testLimit // bytes)
		expectedBatchCountFloat = float64(testCount) / float64(groupSize) / float64(testLimit/testBytes)
		// Round up
		expectedBatchCount = int(math.Ceil(expectedBatchCountFloat))
	)

	for i := 1; i <= testCount; i += groupSize {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("1.2.3.4"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("1.2.3.4")),
				sizeValues: 1,
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("txt-record"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("txt-record")),
				sizeValues: 1,
			},
		)
	}

	batchCs := batchChangeSet(cs, defaultBatchChangeSize, testLimit, defaultBatchChangeSizeValues)

	require.Equal(t, expectedBatchCount, len(batchCs))
}

func TestAWSBatchChangeSetExceedingBytesLimitUpsert(t *testing.T) {
	const (
		testCount = 50
		testLimit = 100
		groupSize = 2
	)

	var (
		cs Route53Changes
		// Bytes for each name multiplied by 2 for Upsert records
		testBytes = (len([]byte("1.2.3.4")) + len([]byte("test-record"))) * 2
		// testCount / groupSize / (testLimit // bytes)
		expectedBatchCountFloat = float64(testCount) / float64(groupSize) / float64(testLimit/testBytes)
		// Round up
		expectedBatchCount = int(math.Ceil(expectedBatchCountFloat))
	)

	for i := 1; i <= testCount; i += groupSize {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionUpsert,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("1.2.3.4"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("1.2.3.4")) * 2,
				sizeValues: 1,
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionUpsert,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("txt-record"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("txt-record")) * 2,
				sizeValues: 1,
			},
		)
	}

	batchCs := batchChangeSet(cs, defaultBatchChangeSize, testLimit, defaultBatchChangeSizeValues)

	require.Equal(t, expectedBatchCount, len(batchCs))
}

func TestAWSBatchChangeSetExceedingValuesLimit(t *testing.T) {
	const (
		testCount = 50
		testLimit = 100
		groupSize = 2
		// Values for each group
		testValues = 2
	)

	var (
		cs Route53Changes
		// testCount / groupSize / (testLimit // bytes)
		expectedBatchCountFloat = float64(testCount) / float64(groupSize) / float64(testLimit/testValues)
		// Round up
		expectedBatchCount = int(math.Ceil(expectedBatchCountFloat))
	)

	for i := 1; i <= testCount; i += groupSize {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("1.2.3.4"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("1.2.3.4")),
				sizeValues: 1,
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionCreate,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("txt-record"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("txt-record")),
				sizeValues: 1,
			},
		)
	}

	batchCs := batchChangeSet(cs, defaultBatchChangeSize, defaultBatchChangeSizeBytes, testLimit)

	require.Equal(t, expectedBatchCount, len(batchCs))
}

func TestAWSBatchChangeSetExceedingValuesLimitUpsert(t *testing.T) {
	const (
		testCount = 50
		testLimit = 100
		groupSize = 2
		// Values for each group multiplied by 2 for Upsert records
		testValues = 2 * 2
	)

	var (
		cs Route53Changes
		// testCount / groupSize / (testLimit // bytes)
		expectedBatchCountFloat = float64(testCount) / float64(groupSize) / float64(testLimit/testValues)
		// Round up
		expectedBatchCount = int(math.Ceil(expectedBatchCountFloat))
	)

	for i := 1; i <= testCount; i += groupSize {
		cs = append(cs,
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionUpsert,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeA,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("1.2.3.4"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("1.2.3.4")) * 2,
				sizeValues: 1,
			},
			&Route53Change{
				Change: route53types.Change{
					Action: route53types.ChangeActionUpsert,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("host-%d", i)),
						Type: route53types.RRTypeTxt,
						ResourceRecords: []route53types.ResourceRecord{
							{
								Value: aws.String("txt-record"),
							},
						},
					},
				},
				sizeBytes:  len([]byte("txt-record")) * 2,
				sizeValues: 1,
			},
		)
	}

	batchCs := batchChangeSet(cs, defaultBatchChangeSize, defaultBatchChangeSizeBytes, testLimit)

	require.Equal(t, expectedBatchCount, len(batchCs))
}

func validateEndpoints(t *testing.T, provider *AWSProvider, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "actual and expected endpoints don't match. %+v:%+v", endpoints, expected)

	normalized, err := provider.AdjustEndpoints(endpoints)
	assert.NoError(t, err)
	assert.True(t, testutils.SameEndpoints(normalized, expected), "actual and normalized endpoints don't match. %+v:%+v", endpoints, normalized)
}

func validateAWSZones(t *testing.T, zones map[string]*route53types.HostedZone, expected map[string]*route53types.HostedZone) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		validateAWSZone(t, zone, expected[i])
	}
}

func validateAWSZone(t *testing.T, zone *route53types.HostedZone, expected *route53types.HostedZone) {
	assert.Equal(t, *expected.Id, *zone.Id)
	assert.Equal(t, *expected.Name, *zone.Name)
}

func validateAWSChangeRecords(t *testing.T, records Route53Changes, expected Route53Changes) {
	require.Len(t, records, len(expected))

	for i := range records {
		validateAWSChangeRecord(t, records[i], expected[i])
	}
}

func validateAWSChangeRecord(t *testing.T, record *Route53Change, expected *Route53Change) {
	assert.Equal(t, expected.Action, record.Action)
	assert.Equal(t, *expected.ResourceRecordSet.Name, *record.ResourceRecordSet.Name)
	assert.Equal(t, expected.ResourceRecordSet.Type, record.ResourceRecordSet.Type)
}

func TestAWSCreateRecordsWithCNAME(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)

	records := []*endpoint.Endpoint{
		{DNSName: "create-test.zone-1.ext-dns-test-2.teapot.zalan.do", Targets: endpoint.Targets{"foo.example.org"}, RecordType: endpoint.RecordTypeCNAME},
	}

	adjusted, err := provider.AdjustEndpoints(records)
	require.NoError(t, err)
	require.NoError(t, provider.ApplyChanges(context.Background(), &plan.Changes{
		Create: adjusted,
	}))

	recordSets := listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do.")

	validateRecords(t, recordSets, []route53types.ResourceRecordSet{
		{
			Name: aws.String("create-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
			Type: route53types.RRTypeCname,
			TTL:  aws.Int64(300),
			ResourceRecords: []route53types.ResourceRecord{
				{
					Value: aws.String("foo.example.org"),
				},
			},
		},
	})
}

func TestAWSCreateRecordsWithALIAS(t *testing.T) {
	for key, evaluateTargetHealth := range map[string]bool{
		"true":  true,
		"false": false,
		"":      false,
	} {
		provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)

		// Test dualstack and ipv4 load balancer targets
		records := []*endpoint.Endpoint{
			{
				DNSName:    "create-test.zone-1.ext-dns-test-2.teapot.zalan.do",
				Targets:    endpoint.Targets{"foo.eu-central-1.elb.amazonaws.com"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{
					endpoint.ProviderSpecificProperty{
						Name:  providerSpecificAlias,
						Value: "true",
					},
					endpoint.ProviderSpecificProperty{
						Name:  providerSpecificEvaluateTargetHealth,
						Value: key,
					},
				},
			},
			{
				DNSName:    "create-test-dualstack.zone-1.ext-dns-test-2.teapot.zalan.do",
				Targets:    endpoint.Targets{"bar.eu-central-1.elb.amazonaws.com"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{
					endpoint.ProviderSpecificProperty{
						Name:  providerSpecificAlias,
						Value: "true",
					},
					endpoint.ProviderSpecificProperty{
						Name:  providerSpecificEvaluateTargetHealth,
						Value: key,
					},
				},
				Labels: map[string]string{endpoint.DualstackLabelKey: "true"},
			},
		}
		adjusted, err := provider.AdjustEndpoints(records)
		require.NoError(t, err)
		require.NoError(t, provider.ApplyChanges(context.Background(), &plan.Changes{
			Create: adjusted,
		}))

		recordSets := listAWSRecords(t, provider.clients[defaultAWSProfile], "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do.")

		validateRecords(t, recordSets, []route53types.ResourceRecordSet{
			{
				AliasTarget: &route53types.AliasTarget{
					DNSName:              aws.String("foo.eu-central-1.elb.amazonaws.com."),
					EvaluateTargetHealth: evaluateTargetHealth,
					HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
				},
				Name: aws.String("create-test.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type: route53types.RRTypeA,
			},
			{
				AliasTarget: &route53types.AliasTarget{
					DNSName:              aws.String("bar.eu-central-1.elb.amazonaws.com."),
					EvaluateTargetHealth: evaluateTargetHealth,
					HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
				},
				Name: aws.String("create-test-dualstack.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type: route53types.RRTypeA,
			},
			{
				AliasTarget: &route53types.AliasTarget{
					DNSName:              aws.String("bar.eu-central-1.elb.amazonaws.com."),
					EvaluateTargetHealth: evaluateTargetHealth,
					HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
				},
				Name: aws.String("create-test-dualstack.zone-1.ext-dns-test-2.teapot.zalan.do."),
				Type: route53types.RRTypeAaaa,
			},
		})
	}
}

func TestAWSisLoadBalancer(t *testing.T) {
	for _, tc := range []struct {
		target      string
		recordType  string
		preferCNAME bool
		expected    bool
	}{
		{"bar.eu-central-1.elb.amazonaws.com", endpoint.RecordTypeCNAME, false, true},
		{"bar.eu-central-1.elb.amazonaws.com", endpoint.RecordTypeCNAME, true, false},
		{"foo.example.org", endpoint.RecordTypeCNAME, false, false},
		{"foo.example.org", endpoint.RecordTypeCNAME, true, false},
	} {
		ep := &endpoint.Endpoint{
			Targets:    endpoint.Targets{tc.target},
			RecordType: tc.recordType,
		}
		assert.Equal(t, tc.expected, useAlias(ep, tc.preferCNAME))
	}
}

func TestAWSisAWSAlias(t *testing.T) {
	for _, tc := range []struct {
		target     string
		recordType string
		alias      bool
		hz         string
	}{
		{"foo.example.org", endpoint.RecordTypeA, false, ""},                                 // normal A record
		{"bar.eu-central-1.elb.amazonaws.com", endpoint.RecordTypeA, true, "Z215JYRZR1TBD5"}, // pointing to ELB DNS name
		{"foobar.example.org", endpoint.RecordTypeA, true, "Z1234567890ABC"},                 // HZID retrieved by Route53
		{"baz.example.org", endpoint.RecordTypeA, true, sameZoneAlias},                       // record to be created
	} {
		ep := &endpoint.Endpoint{
			Targets:    endpoint.Targets{tc.target},
			RecordType: tc.recordType,
		}
		if tc.alias {
			ep = ep.WithProviderSpecific(providerSpecificAlias, "true")
			ep = ep.WithProviderSpecific(providerSpecificTargetHostedZone, tc.hz)
		}
		assert.Equal(t, tc.hz, isAWSAlias(ep), "%v", tc)
	}
}

func TestAWSCanonicalHostedZone(t *testing.T) {
	for suffix, id := range canonicalHostedZones {
		zone := canonicalHostedZone(fmt.Sprintf("foo.%s", suffix))
		assert.Equal(t, id, zone)
	}

	zone := canonicalHostedZone("foo.example.org")
	assert.Equal(t, "", zone, "no canonical zone should be returned for a non-aws hostname")
}

func TestAWSSuitableZones(t *testing.T) {
	zones := map[string]*profiledZone{
		// Public domain
		"example-org": {profile: defaultAWSProfile, zone: &route53types.HostedZone{Id: aws.String("example-org"), Name: aws.String("example.org.")}},
		// Public subdomain
		"bar-example-org": {profile: defaultAWSProfile, zone: &route53types.HostedZone{Id: aws.String("bar-example-org"), Name: aws.String("bar.example.org."), Config: &route53types.HostedZoneConfig{PrivateZone: false}}},
		// Public subdomain
		"longfoo-bar-example-org": {profile: defaultAWSProfile, zone: &route53types.HostedZone{Id: aws.String("longfoo-bar-example-org"), Name: aws.String("longfoo.bar.example.org.")}},
		// Private domain
		"example-org-private": {profile: defaultAWSProfile, zone: &route53types.HostedZone{Id: aws.String("example-org-private"), Name: aws.String("example.org."), Config: &route53types.HostedZoneConfig{PrivateZone: true}}},
		// Private subdomain
		"bar-example-org-private": {profile: defaultAWSProfile, zone: &route53types.HostedZone{Id: aws.String("bar-example-org-private"), Name: aws.String("bar.example.org."), Config: &route53types.HostedZoneConfig{PrivateZone: true}}},
	}

	for _, tc := range []struct {
		hostname string
		expected []*profiledZone
	}{
		// bar.example.org is NOT suitable
		{"foobar.example.org.", []*profiledZone{zones["example-org-private"], zones["example-org"]}},

		// all matching private zones are suitable
		// https://github.com/kubernetes-sigs/external-dns/pull/356
		{"bar.example.org.", []*profiledZone{zones["example-org-private"], zones["bar-example-org-private"], zones["bar-example-org"]}},

		{"foo.bar.example.org.", []*profiledZone{zones["example-org-private"], zones["bar-example-org-private"], zones["bar-example-org"]}},
		{"foo.example.org.", []*profiledZone{zones["example-org-private"], zones["example-org"]}},
		{"foo.kubernetes.io.", nil},
	} {
		suitableZones := suitableZones(tc.hostname, zones)
		sort.Slice(suitableZones, func(i, j int) bool {
			return *suitableZones[i].zone.Id < *suitableZones[j].zone.Id
		})
		sort.Slice(tc.expected, func(i, j int) bool {
			return *tc.expected[i].zone.Id < *tc.expected[j].zone.Id
		})
		assert.Equal(t, tc.expected, suitableZones)
	}
}

func createAWSZone(t *testing.T, provider *AWSProvider, zone *route53types.HostedZone) {
	params := &route53.CreateHostedZoneInput{
		CallerReference:  aws.String("external-dns.alpha.kubernetes.io/test-zone"),
		Name:             zone.Name,
		HostedZoneConfig: zone.Config,
	}

	if _, err := provider.clients[defaultAWSProfile].CreateHostedZone(context.Background(), params); err != nil {
		var hzExists *route53types.HostedZoneAlreadyExists
		require.ErrorAs(t, err, &hzExists)
	}
}

func setAWSRecords(t *testing.T, provider *AWSProvider, records []route53types.ResourceRecordSet) {
	dryRun := provider.dryRun
	provider.dryRun = false
	defer func() {
		provider.dryRun = dryRun
	}()

	ctx := context.Background()
	endpoints, err := provider.Records(ctx)
	require.NoError(t, err)

	validateEndpoints(t, provider, endpoints, []*endpoint.Endpoint{})

	var changes Route53Changes
	for _, record := range records {
		changes = append(changes, &Route53Change{
			Change: route53types.Change{
				Action:            route53types.ChangeActionCreate,
				ResourceRecordSet: &record,
			},
		})
	}

	zones, err := provider.zones(ctx)
	require.NoError(t, err)
	err = provider.submitChanges(ctx, changes, zones)
	require.NoError(t, err)

	_, err = provider.Records(ctx)
	require.NoError(t, err)
}

func listAWSRecords(t *testing.T, client Route53API, zone string) []route53types.ResourceRecordSet {
	resp, err := client.ListResourceRecordSets(context.Background(), &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zone),
		MaxItems:     aws.Int32(route53PageSize),
	})
	require.NoError(t, err)

	return resp.ResourceRecordSets
}

func newAWSProvider(t *testing.T, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneTypeFilter provider.ZoneTypeFilter, evaluateTargetHealth, dryRun bool, records []route53types.ResourceRecordSet) (*AWSProvider, *Route53APIStub) {
	return newAWSProviderWithTagFilter(t, domainFilter, zoneIDFilter, zoneTypeFilter, provider.NewZoneTagFilter([]string{}), evaluateTargetHealth, dryRun, records)
}

func newAWSProviderWithTagFilter(t *testing.T, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneTypeFilter provider.ZoneTypeFilter, zoneTagFilter provider.ZoneTagFilter, evaluateTargetHealth, dryRun bool, records []route53types.ResourceRecordSet) (*AWSProvider, *Route53APIStub) {
	client := NewRoute53APIStub(t)

	provider := &AWSProvider{
		clients:               map[string]Route53API{defaultAWSProfile: client},
		batchChangeSize:       defaultBatchChangeSize,
		batchChangeSizeBytes:  defaultBatchChangeSizeBytes,
		batchChangeSizeValues: defaultBatchChangeSizeValues,
		batchChangeInterval:   defaultBatchChangeInterval,
		evaluateTargetHealth:  evaluateTargetHealth,
		domainFilter:          domainFilter,
		zoneIDFilter:          zoneIDFilter,
		zoneTypeFilter:        zoneTypeFilter,
		zoneTagFilter:         zoneTagFilter,
		dryRun:                false,
		zonesCache:            &zonesListCache{duration: 1 * time.Minute},
		failedChangesQueue:    make(map[string]Route53Changes),
	}

	createAWSZone(t, provider, &route53types.HostedZone{
		Id:     aws.String("/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do."),
		Name:   aws.String("zone-1.ext-dns-test-2.teapot.zalan.do."),
		Config: &route53types.HostedZoneConfig{PrivateZone: false},
	})

	createAWSZone(t, provider, &route53types.HostedZone{
		Id:     aws.String("/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do."),
		Name:   aws.String("zone-2.ext-dns-test-2.teapot.zalan.do."),
		Config: &route53types.HostedZoneConfig{PrivateZone: false},
	})

	createAWSZone(t, provider, &route53types.HostedZone{
		Id:     aws.String("/hostedzone/zone-3.ext-dns-test-2.teapot.zalan.do."),
		Name:   aws.String("zone-3.ext-dns-test-2.teapot.zalan.do."),
		Config: &route53types.HostedZoneConfig{PrivateZone: true},
	})

	// filtered out by domain filter
	createAWSZone(t, provider, &route53types.HostedZone{
		Id:     aws.String("/hostedzone/zone-4.ext-dns-test-3.teapot.zalan.do."),
		Name:   aws.String("zone-4.ext-dns-test-3.teapot.zalan.do."),
		Config: &route53types.HostedZoneConfig{PrivateZone: false},
	})

	setupZoneTags(provider.clients[defaultAWSProfile].(*Route53APIStub))

	setAWSRecords(t, provider, records)

	provider.dryRun = dryRun

	return provider, client
}

func setupZoneTags(client *Route53APIStub) {
	addZoneTags(client.zoneTags, "/hostedzone/zone-1.ext-dns-test-2.teapot.zalan.do.", map[string]string{
		"zone-1-tag-1": "tag-1-value",
		"domain":       "test-2",
		"zone":         "1",
	})
	addZoneTags(client.zoneTags, "/hostedzone/zone-2.ext-dns-test-2.teapot.zalan.do.", map[string]string{
		"zone-2-tag-1": "tag-1-value",
		"domain":       "test-2",
		"zone":         "2",
	})
	addZoneTags(client.zoneTags, "/hostedzone/zone-3.ext-dns-test-2.teapot.zalan.do.", map[string]string{
		"zone-3-tag-1": "tag-1-value",
		"domain":       "test-2",
		"zone":         "3",
	})
	addZoneTags(client.zoneTags, "/hostedzone/zone-4.ext-dns-test-2.teapot.zalan.do.", map[string]string{
		"zone-4-tag-1": "tag-1-value",
		"domain":       "test-3",
		"zone":         "4",
	})
}

func addZoneTags(tagMap map[string][]route53types.Tag, zoneID string, tags map[string]string) {
	tagList := make([]route53types.Tag, 0, len(tags))
	for k, v := range tags {
		tagList = append(tagList, route53types.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}
	tagMap[zoneID] = tagList
}

func validateRecords(t *testing.T, records []route53types.ResourceRecordSet, expected []route53types.ResourceRecordSet) {
	assert.ElementsMatch(t, expected, records)
}

func containsRecordWithDNSName(records []*endpoint.Endpoint, dnsName string) bool {
	for _, record := range records {
		if record.DNSName == dnsName {
			return true
		}
	}
	return false
}

func TestRequiresDeleteCreate(t *testing.T) {
	provider, _ := newAWSProvider(t, endpoint.NewDomainFilter([]string{"foo.bar."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), defaultEvaluateTargetHealth, false, nil)

	oldRecordType := endpoint.NewEndpointWithTTL("recordType", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8")
	newRecordType := endpoint.NewEndpointWithTTL("recordType", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar").WithProviderSpecific(providerSpecificAlias, "false")

	assert.False(t, provider.requiresDeleteCreate(oldRecordType, oldRecordType), "actual and expected endpoints don't match. %+v:%+v", oldRecordType, oldRecordType)
	assert.True(t, provider.requiresDeleteCreate(oldRecordType, newRecordType), "actual and expected endpoints don't match. %+v:%+v", oldRecordType, newRecordType)

	oldAtoAlias := endpoint.NewEndpointWithTTL("AtoAlias", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.1.1.1")
	newAtoAlias := endpoint.NewEndpointWithTTL("AtoAlias", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "bar.us-east-1.elb.amazonaws.com").WithProviderSpecific(providerSpecificAlias, "true")

	assert.False(t, provider.requiresDeleteCreate(oldAtoAlias, oldAtoAlias), "actual and expected endpoints don't match. %+v:%+v", oldAtoAlias, oldAtoAlias.DNSName)
	assert.True(t, provider.requiresDeleteCreate(oldAtoAlias, newAtoAlias), "actual and expected endpoints don't match. %+v:%+v", oldAtoAlias, newAtoAlias)

	oldPolicy := endpoint.NewEndpointWithTTL("policy", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8").WithSetIdentifier("nochange").WithProviderSpecific(providerSpecificRegion, "us-east-1")
	newPolicy := endpoint.NewEndpointWithTTL("policy", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8").WithSetIdentifier("nochange").WithProviderSpecific(providerSpecificWeight, "10")

	assert.False(t, provider.requiresDeleteCreate(oldPolicy, oldPolicy), "actual and expected endpoints don't match. %+v:%+v", oldPolicy, oldPolicy)
	assert.True(t, provider.requiresDeleteCreate(oldPolicy, newPolicy), "actual and expected endpoints don't match. %+v:%+v", oldPolicy, newPolicy)

	oldSetIdentifier := endpoint.NewEndpointWithTTL("setIdentifier", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8").WithSetIdentifier("old")
	newSetIdentifier := endpoint.NewEndpointWithTTL("setIdentifier", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8").WithSetIdentifier("new")

	assert.False(t, provider.requiresDeleteCreate(oldSetIdentifier, oldSetIdentifier), "actual and expected endpoints don't match. %+v:%+v", oldSetIdentifier, oldSetIdentifier)
	assert.True(t, provider.requiresDeleteCreate(oldSetIdentifier, newSetIdentifier), "actual and expected endpoints don't match. %+v:%+v", oldSetIdentifier, newSetIdentifier)
}

func TestConvertOctalToAscii(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Characters escaped !\"#$%&'()*+,-/:;",
			input:    "txt-\\041\\042\\043\\044\\045\\046\\047\\050\\051\\052\\053\\054-\\057\\072\\073-test.example.com",
			expected: "txt-!\"#$%&'()*+,-/:;-test.example.com",
		},
		{
			name:     "Characters escaped <=>?@[\\]^_`{|}~",
			input:    "txt-\\074\\075\\076\\077\\100\\133\\134\\135\\136_\\140\\173\\174\\175\\176-test2.example.com",
			expected: "txt-<=>?@[\\]^_`{|}~-test2.example.com",
		},
		{
			name:     "No escaped characters in domain",
			input:    "txt-awesome-test3.example.com",
			expected: "txt-awesome-test3.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := convertOctalToAscii(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
