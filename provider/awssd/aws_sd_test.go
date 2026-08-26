/*
Copyright 2018 The Kubernetes Authors.

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

package awssd

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	sdtypes "github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/plan"
)

func TestAWSSDProvider_Records(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"a-srv": {
				Id:          aws.String("a-srv"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(100),
					}},
				},
			},
			"alias-srv": {
				Id:          aws.String("alias-srv"),
				Name:        aws.String("service2"),
				NamespaceId: aws.String("private"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(100),
					}},
				},
			},
			"cname-srv": {
				Id:          aws.String("cname-srv"),
				Name:        aws.String("service3"),
				NamespaceId: aws.String("private"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeCname,
						TTL:  aws.Int64(80),
					}},
				},
			},
			"aaaa-srv": {
				Id:          aws.String("aaaa-srv"),
				Name:        aws.String("service4"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					NamespaceId:   aws.String("private"),
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeAaaa,
						TTL:  aws.Int64(100),
					}},
				},
			},
			"aaaa-srv-not-managed-without-owner-id": {
				Id:          aws.String("aaaa-srv"),
				Name:        aws.String("service5"),
				Description: nil,
				DnsConfig: &sdtypes.DnsConfig{
					NamespaceId:   aws.String("private"),
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeAaaa,
						TTL:  aws.Int64(100),
					}},
				},
			},
		},
	}

	instances := map[string]map[string]*sdtypes.Instance{
		"a-srv": {
			"1.2.3.4": {
				Id: aws.String("1.2.3.4"),
				Attributes: map[string]string{
					sdInstanceAttrIPV4: "1.2.3.4",
				},
			},
			"1.2.3.5": {
				Id: aws.String("1.2.3.5"),
				Attributes: map[string]string{
					sdInstanceAttrIPV4: "1.2.3.5",
				},
			},
		},
		"alias-srv": {
			"load-balancer.us-east-1.elb.amazonaws.com": {
				Id: aws.String("load-balancer.us-east-1.elb.amazonaws.com"),
				Attributes: map[string]string{
					sdInstanceAttrAlias: "load-balancer.us-east-1.elb.amazonaws.com",
				},
			},
		},
		"cname-srv": {
			"cname.target.com": {
				Id: aws.String("cname.target.com"),
				Attributes: map[string]string{
					sdInstanceAttrCname: "cname.target.com",
				},
			},
		},
		"aaaa-srv": {
			"0000:0000:0000:0000:abcd:abcd:abcd:abcd": {
				Id: aws.String("0000:0000:0000:0000:abcd:abcd:abcd:abcd"),
				Attributes: map[string]string{
					sdInstanceAttrIPV6: "0000:0000:0000:0000:abcd:abcd:abcd:abcd",
				},
			},
		},
	}

	expectedEndpoints := []*endpoint.Endpoint{
		{DNSName: "service1.private.com", Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5"}, RecordType: endpoint.RecordTypeA, RecordTTL: 100, Labels: map[string]string{endpoint.AWSSDDescriptionLabel: "owner-id"}},
		{DNSName: "service2.private.com", Targets: endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"}, RecordType: endpoint.RecordTypeCNAME, RecordTTL: 100, Labels: map[string]string{endpoint.AWSSDDescriptionLabel: "owner-id"}},
		{DNSName: "service3.private.com", Targets: endpoint.Targets{"cname.target.com"}, RecordType: endpoint.RecordTypeCNAME, RecordTTL: 80, Labels: map[string]string{endpoint.AWSSDDescriptionLabel: "owner-id"}},
		{DNSName: "service4.private.com", Targets: endpoint.Targets{"0000:0000:0000:0000:abcd:abcd:abcd:abcd"}, RecordType: endpoint.RecordTypeAAAA, RecordTTL: 100, Labels: map[string]string{endpoint.AWSSDDescriptionLabel: "owner-id"}},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
		instances:  instances,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	endpoints, _ := provider.Records(t.Context())

	assert.True(t, testutils.SameEndpoints(expectedEndpoints, endpoints), "expected and actual endpoints don't match, expected=%v, actual=%v", expectedEndpoints, endpoints)
}

func TestAWSSDProvider_ApplyChanges(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	expectedEndpoints := []*endpoint.Endpoint{
		{DNSName: "service1.private.com", Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5"}, RecordType: endpoint.RecordTypeA, RecordTTL: 60},
		{DNSName: "service2.private.com", Targets: endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"}, RecordType: endpoint.RecordTypeCNAME, RecordTTL: 80},
		{DNSName: "service3.private.com", Targets: endpoint.Targets{"cname.target.com"}, RecordType: endpoint.RecordTypeCNAME, RecordTTL: 100},
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	ctx := t.Context()

	// apply creates
	err := provider.ApplyChanges(ctx, &plan.Changes{
		Create: expectedEndpoints,
	})
	assert.NoError(t, err)

	// make sure services were created
	assert.Len(t, api.services["private"], 3)
	existingServices, _ := provider.ListServicesByNamespaceID(t.Context(), namespaces["private"].Id)
	assert.NotNil(t, existingServices["service1"])
	assert.NotNil(t, existingServices["service2"])
	assert.NotNil(t, existingServices["service3"])

	// make sure instances were registered
	endpoints, _ := provider.Records(ctx)
	assert.True(t, testutils.SameEndpoints(expectedEndpoints, endpoints), "expected and actual endpoints don't match, expected=%v, actual=%v", expectedEndpoints, endpoints)

	ctx = t.Context()
	// apply deletes
	err = provider.ApplyChanges(ctx, &plan.Changes{
		Delete: expectedEndpoints,
	})
	assert.NoError(t, err)

	// make sure all instances are gone
	endpoints, _ = provider.Records(ctx)
	assert.Empty(t, endpoints)
}

func TestAWSSDProvider_ApplyChanges_Update(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	oldEndpoints := []*endpoint.Endpoint{
		{DNSName: "service1.private.com", Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5"}, RecordType: endpoint.RecordTypeA, RecordTTL: 60},
	}

	newEndpoints := []*endpoint.Endpoint{
		{DNSName: "service1.private.com", Targets: endpoint.Targets{"1.2.3.4", "1.2.3.6"}, RecordType: endpoint.RecordTypeA, RecordTTL: 60},
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	ctx := t.Context()

	// apply creates
	_ = provider.ApplyChanges(ctx, &plan.Changes{
		Create: oldEndpoints,
	})

	ctx = t.Context()

	// apply update
	_ = provider.ApplyChanges(ctx, &plan.Changes{
		UpdateOld: oldEndpoints,
		UpdateNew: newEndpoints,
	})

	// make sure services were created
	assert.Len(t, api.services["private"], 1)
	existingServices, _ := provider.ListServicesByNamespaceID(ctx, namespaces["private"].Id)
	assert.NotNil(t, existingServices["service1"])

	// make sure instances were registered
	endpoints, _ := provider.Records(ctx)
	assert.True(t, testutils.SameEndpoints(newEndpoints, endpoints), "expected and actual endpoints don't match, expected=%v, actual=%v", newEndpoints, endpoints)

	// make sure only one instance is de-registered
	assert.Len(t, api.deregistered, 1)
	assert.Equal(t, "1.2.3.5", api.deregistered[0], "wrong target de-registered")
}

func TestAWSSDProvider_ApplyChanges_DottedServiceName(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"dev-local": {
			Id:   aws.String("dev-local"),
			Name: aws.String("dev.local"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	createEndpoints := []*endpoint.Endpoint{
		{DNSName: "my-app.elb.dev.local", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA, RecordTTL: 60},
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{"dev.local"}), "", "")

	ctx := t.Context()

	err := provider.ApplyChanges(ctx, &plan.Changes{
		Create: createEndpoints,
	})
	require.NoError(t, err)

	// service must be created with the dotted name "my-app.elb"
	assert.Len(t, api.services["dev-local"], 1)
	existingServices, err := provider.ListServicesByNamespaceID(ctx, namespaces["dev-local"].Id)
	require.NoError(t, err)
	assert.NotNil(t, existingServices["my-app.elb"], "service should be named 'my-app.elb'")

	// verify the record round-trips through Records()
	endpoints, err := provider.Records(ctx)
	require.NoError(t, err)
	assert.True(t, testutils.SameEndpoints(createEndpoints, endpoints),
		"expected and actual endpoints don't match, expected=%v, actual=%v", createEndpoints, endpoints)

	// apply deletes
	err = provider.ApplyChanges(ctx, &plan.Changes{
		Delete: createEndpoints,
	})
	require.NoError(t, err)

	endpoints, _ = provider.Records(ctx)
	assert.Empty(t, endpoints)
}

func TestAWSSDProvider_ListNamespaces(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
		"public": {
			Id:   aws.String("public"),
			Name: aws.String("public.com"),
			Type: sdtypes.NamespaceTypeDnsPublic,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
	}

	for _, tc := range []struct {
		msg                 string
		domainFilter        *endpoint.DomainFilter
		namespaceTypeFilter string
		expectedNamespaces  []*sdtypes.NamespaceSummary
	}{
		{"public filter", endpoint.NewDomainFilter([]string{}), "public", []*sdtypes.NamespaceSummary{namespaceToNamespaceSummary(namespaces["public"])}},
		{"private filter", endpoint.NewDomainFilter([]string{}), "private", []*sdtypes.NamespaceSummary{namespaceToNamespaceSummary(namespaces["private"])}},
		{"optional filter", endpoint.NewDomainFilter([]string{}), "", []*sdtypes.NamespaceSummary{namespaceToNamespaceSummary(namespaces["public"]), namespaceToNamespaceSummary(namespaces["private"])}},
		{"domain filter", endpoint.NewDomainFilter([]string{"public.com"}), "", []*sdtypes.NamespaceSummary{namespaceToNamespaceSummary(namespaces["public"])}},
		{"non-existing domain", endpoint.NewDomainFilter([]string{"xxx.com"}), "", []*sdtypes.NamespaceSummary{}},
	} {
		provider := newTestAWSSDProvider(api, tc.domainFilter, tc.namespaceTypeFilter, "")

		result, err := provider.ListNamespaces(t.Context())
		require.NoError(t, err)

		expectedMap := make(map[string]*sdtypes.NamespaceSummary)
		resultMap := make(map[string]*sdtypes.NamespaceSummary)
		for _, ns := range tc.expectedNamespaces {
			expectedMap[*ns.Id] = ns
		}
		for _, ns := range result {
			resultMap[*ns.Id] = ns
		}

		if !reflect.DeepEqual(resultMap, expectedMap) {
			t.Errorf("AWSSDProvider.ListNamespaces() error = %v, wantErr %v", result, tc.expectedNamespaces)
		}
	}
}

func TestAWSSDProvider_ListServicesByNamespace(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
		"public": {
			Id:   aws.String("public"),
			Name: aws.String("public.com"),
			Type: sdtypes.NamespaceTypeDnsPublic,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
			},
			"srv2": {
				Id:          aws.String("srv2"),
				Name:        aws.String("service2"),
				NamespaceId: aws.String("private"),
			},
		},
		"public": {
			"srv3": {
				Id:          aws.String("srv3"),
				Name:        aws.String("service3"),
				NamespaceId: aws.String("public"),
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	for _, tc := range []struct {
		expectedServices map[string]*sdtypes.Service
	}{
		{map[string]*sdtypes.Service{"service1": services["private"]["srv1"], "service2": services["private"]["srv2"]}},
	} {
		provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

		result, err := provider.ListServicesByNamespaceID(t.Context(), namespaces["private"].Id)
		require.NoError(t, err)
		assert.Equal(t, tc.expectedServices, result)
	}
}

func TestAWSSDProvider_CreateService(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
	}

	expectedServices := make(map[string]*sdtypes.Service)

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	// A type
	_, err := provider.CreateService(t.Context(), aws.String("private"), aws.String("A-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "A-srv",
		},
		RecordType: endpoint.RecordTypeA,
		RecordTTL:  60,
		Targets:    endpoint.Targets{"1.2.3.4"},
	})
	assert.NoError(t, err)

	expectedServices["A-srv"] = &sdtypes.Service{
		Name:        aws.String("A-srv"),
		Description: aws.String("A-srv"),
		DnsConfig: &sdtypes.DnsConfig{
			RoutingPolicy: sdtypes.RoutingPolicyMultivalue,
			DnsRecords: []sdtypes.DnsRecord{{
				Type: sdtypes.RecordTypeA,
				TTL:  aws.Int64(60),
			}},
		},
		NamespaceId: aws.String("private"),
	}

	// AAAA type
	_, err = provider.CreateService(t.Context(), aws.String("private"), aws.String("AAAA-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "AAAA-srv",
		},
		RecordType: endpoint.RecordTypeAAAA,
		RecordTTL:  60,
		Targets:    endpoint.Targets{"::1234:5678:"},
	})
	assert.NoError(t, err)
	expectedServices["AAAA-srv"] = &sdtypes.Service{
		Name:        aws.String("AAAA-srv"),
		Description: aws.String("AAAA-srv"),
		DnsConfig: &sdtypes.DnsConfig{
			RoutingPolicy: sdtypes.RoutingPolicyMultivalue,
			DnsRecords: []sdtypes.DnsRecord{{
				Type: sdtypes.RecordTypeAaaa,
				TTL:  aws.Int64(60),
			}},
		},
		NamespaceId: aws.String("private"),
	}

	// CNAME type
	_, err = provider.CreateService(t.Context(), aws.String("private"), aws.String("CNAME-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "CNAME-srv",
		},
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  80,
		Targets:    endpoint.Targets{"cname.target.com"},
	})
	assert.NoError(t, err)
	expectedServices["CNAME-srv"] = &sdtypes.Service{
		Name:        aws.String("CNAME-srv"),
		Description: aws.String("CNAME-srv"),
		DnsConfig: &sdtypes.DnsConfig{
			RoutingPolicy: sdtypes.RoutingPolicyWeighted,
			DnsRecords: []sdtypes.DnsRecord{{
				Type: sdtypes.RecordTypeCname,
				TTL:  aws.Int64(80),
			}},
		},
		NamespaceId: aws.String("private"),
	}

	// ALIAS type (dual-stack: both A and AAAA)
	_, err = provider.CreateService(t.Context(), aws.String("private"), aws.String("ALIAS-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "ALIAS-srv",
		},
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"},
	})
	assert.NoError(t, err)
	expectedServices["ALIAS-srv"] = &sdtypes.Service{
		Name:        aws.String("ALIAS-srv"),
		Description: aws.String("ALIAS-srv"),
		DnsConfig: &sdtypes.DnsConfig{
			RoutingPolicy: sdtypes.RoutingPolicyWeighted,
			DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(100)},
			},
		},
		NamespaceId: aws.String("private"),
	}

	// ALIAS type with "dualstack." prefixed hostname must be recognized identically -
	// detection does not depend on the "dualstack." substring.
	_, err = provider.CreateService(t.Context(), aws.String("private"), aws.String("ALIAS-dualstack-prefix-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "ALIAS-dualstack-prefix-srv",
		},
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"dualstack.load-balancer.us-east-1.elb.amazonaws.com"},
	})
	assert.NoError(t, err)
	expectedServices["ALIAS-dualstack-prefix-srv"] = &sdtypes.Service{
		Name:        aws.String("ALIAS-dualstack-prefix-srv"),
		Description: aws.String("ALIAS-dualstack-prefix-srv"),
		DnsConfig: &sdtypes.DnsConfig{
			RoutingPolicy: sdtypes.RoutingPolicyWeighted,
			DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(100)},
			},
		},
		NamespaceId: aws.String("private"),
	}

	testHelperAWSSDServicesMapsEqual(t, expectedServices, api.services["private"])
}

func TestAWSSDProvider_CreateServiceDryRun(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")
	provider.dryRun = true

	service, err := provider.CreateService(t.Context(), aws.String("private"), aws.String("A-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			endpoint.AWSSDDescriptionLabel: "A-srv",
		},
		RecordType: endpoint.RecordTypeA,
		RecordTTL:  60,
		Targets:    endpoint.Targets{"1.2.3.4"},
	})
	assert.NoError(t, err)

	assert.NotNil(t, service)
	assert.Equal(t, "dry-run-service", *service.Name)
}

func TestAWSSDProvider_CreateService_LabelNotSet(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   make(map[string]map[string]*sdtypes.Service),
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "owner-123")

	service, err := provider.CreateService(t.Context(), aws.String("private"), aws.String("A-srv"), &endpoint.Endpoint{
		Labels: map[string]string{
			"wrong-unsupported-label": "A-srv",
		},
		RecordType: endpoint.RecordTypeA,
		RecordTTL:  60,
		Targets:    endpoint.Targets{"1.2.3.4"},
	})

	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Empty(t, *service.Description)
}

func TestAWSSDProvider_UpdateService(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyMultivalue,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(60),
					}},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	// update service with different TTL
	err := provider.UpdateService(t.Context(), services["private"]["srv1"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeA,
		RecordTTL:  100,
	})

	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 1)
	assert.Equal(t, int64(100), *api.services["private"]["srv1"].DnsConfig.DnsRecords[0].TTL)
}

func TestAWSSDProvider_UpdateService_DryRun(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyMultivalue,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(60),
					}},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")
	provider.dryRun = true

	// update service with different TTL
	err := provider.UpdateService(t.Context(), services["private"]["srv1"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeAAAA,
		RecordTTL:  100,
	})

	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 1)
	// records should not be updated
	assert.NotEqual(t, 100, api.services["private"]["srv1"].DnsConfig.DnsRecords[0].TTL)
	assert.NotEqual(t, endpoint.RecordTypeAAAA, api.services["private"]["srv1"].DnsConfig.DnsRecords[0].Type)
}

func TestAWSSDProvider_ServiceTypesFromEndpoint(t *testing.T) {
	provider := newTestAWSSDProvider(&AWSSDClientStub{}, endpoint.NewDomainFilter([]string{}), "", "")

	for _, tc := range []struct {
		name     string
		ep       *endpoint.Endpoint
		expected []sdtypes.RecordType
	}{
		{
			name:     "A endpoint",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeA},
		},
		{
			name:     "AAAA endpoint",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"::1234:5678:"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeAaaa},
		},
		{
			name:     "ordinary non-AWS CNAME",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"cname.target.com"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeCname},
		},
		{
			name:     "recognized AWS ELB alias",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeA, sdtypes.RecordTypeAaaa},
		},
		{
			name:     "recognized AWS ELB alias with dualstack. prefix behaves identically",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"dualstack.load-balancer.us-east-1.elb.amazonaws.com"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeA, sdtypes.RecordTypeAaaa},
		},
		{
			name:     "recognized AWS NLB alias",
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"load-balancer.elb.us-west-2.amazonaws.com"}},
			expected: []sdtypes.RecordType{sdtypes.RecordTypeA, sdtypes.RecordTypeAaaa},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, provider.serviceTypesFromEndpoint(tc.ep))
		})
	}
}

func TestAWSSDProvider_UpdateService_Alias(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeCname,
						TTL:  aws.Int64(60),
					}},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	// updating with an AWS LB alias endpoint must build the same [A, AAAA] set that CreateService would
	err := provider.UpdateService(t.Context(), services["private"]["srv1"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"},
	})

	require.NoError(t, err)
	gotRecords := api.services["private"]["srv1"].DnsConfig.DnsRecords
	require.Len(t, gotRecords, 2)
	assert.Equal(t, sdtypes.RecordTypeA, gotRecords[0].Type)
	assert.Equal(t, int64(100), *gotRecords[0].TTL)
	assert.Equal(t, sdtypes.RecordTypeAaaa, gotRecords[1].Type)
	assert.Equal(t, int64(100), *gotRecords[1].TTL)
}

func TestAWSSDProvider_ServiceNeedsUpdate(t *testing.T) {
	provider := newTestAWSSDProvider(&AWSSDClientStub{}, endpoint.NewDomainFilter([]string{}), "", "")

	aliasEndpoint := &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"},
	}
	cnameEndpoint := &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"cname.target.com"},
	}

	for _, tc := range []struct {
		name     string
		srv      *sdtypes.Service
		ep       *endpoint.Endpoint
		expected bool
	}{
		{
			name: "existing A-only AWS alias converges to [A, AAAA]",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
			}}},
			ep:       aliasEndpoint,
			expected: true,
		},
		{
			name: "existing [A, AAAA] alias with matching TTL needs no update",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(100)},
			}}},
			ep:       aliasEndpoint,
			expected: false,
		},
		{
			name: "existing [AAAA, A] order-swapped alias with matching TTL needs no update",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(100)},
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
			}}},
			ep:       aliasEndpoint,
			expected: false,
		},
		{
			name: "existing [A, AAAA] alias transitions to ordinary CNAME",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(100)},
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(100)},
			}}},
			ep:       cnameEndpoint,
			expected: true,
		},
		{
			name: "existing ordinary CNAME transitions to AWS alias",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeCname, TTL: aws.Int64(100)},
			}}},
			ep:       aliasEndpoint,
			expected: true,
		},
		{
			name: "TTL mismatch alone still triggers update, exactly as before",
			srv: &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{
				{Type: sdtypes.RecordTypeA, TTL: aws.Int64(60)},
				{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(60)},
			}}},
			ep:       aliasEndpoint,
			expected: true,
		},
		{
			name:     "unconfigured TTL and matching types needs no update",
			srv:      &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{DnsRecords: []sdtypes.DnsRecord{{Type: sdtypes.RecordTypeCname, TTL: aws.Int64(999)}}}},
			ep:       &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"cname.target.com"}},
			expected: false,
		},
		{
			name:     "nil DnsConfig is handled safely and triggers update",
			srv:      &sdtypes.Service{},
			ep:       aliasEndpoint,
			expected: true,
		},
		{
			name:     "empty DnsRecords is handled safely and triggers update",
			srv:      &sdtypes.Service{DnsConfig: &sdtypes.DnsConfig{}},
			ep:       aliasEndpoint,
			expected: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, provider.serviceNeedsUpdate(tc.srv, tc.ep))
		})
	}
}

// TestAWSSDProvider_ApplyChanges_ExistingAliasConvergesToDualStack proves that once an
// endpoint otherwise reaches submitCreates (here: a Create change for a service that
// already exists in Cloud Map), an existing A-only alias service configuration is
// converged to [A, AAAA]. It does NOT prove - and this patch does not implement -
// automatic migration of a completely unchanged, already-reconciled alias service:
// Records() still reconstructs such a service as a single CNAME endpoint, so if
// nothing about the endpoint changes, the plan stays empty and ApplyChanges returns
// early without ever reaching this code path.
func TestAWSSDProvider_ApplyChanges_ExistingAliasConvergesToDualStack(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"service1": {
				Id:          aws.String("service1"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(100),
					}},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	aliasEndpoint := &endpoint.Endpoint{
		DNSName:    "service1.private.com",
		RecordType: endpoint.RecordTypeCNAME,
		RecordTTL:  100,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"},
	}

	err := provider.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{aliasEndpoint}})
	require.NoError(t, err)

	gotRecords := api.services["private"]["service1"].DnsConfig.DnsRecords
	require.Len(t, gotRecords, 2)
	assert.Equal(t, sdtypes.RecordTypeA, gotRecords[0].Type)
	assert.Equal(t, sdtypes.RecordTypeAaaa, gotRecords[1].Type)

	// applying the exact same (now already-converged) change again must not need
	// a further update - the service configuration is stable at [A, AAAA].
	err = provider.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{aliasEndpoint}})
	require.NoError(t, err)
	assert.Equal(t, gotRecords, api.services["private"]["service1"].DnsConfig.DnsRecords)
}

func TestAWSSDProvider_RegisterInstance_AliasSendsOnlyAliasAttribute(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"alias-srv": {
				Id:          aws.String("alias-srv"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{
						{Type: sdtypes.RecordTypeA, TTL: aws.Int64(60)},
						{Type: sdtypes.RecordTypeAaaa, TTL: aws.Int64(60)},
					},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	err := provider.RegisterInstance(t.Context(), services["private"]["alias-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		DNSName:    "service1.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com"},
	})
	require.NoError(t, err)

	inst := api.instances["alias-srv"]["load-balancer.us-east-1.elb.amazonaws.com"]
	require.NotNil(t, inst)
	assert.Equal(t, map[string]string{sdInstanceAttrAlias: "load-balancer.us-east-1.elb.amazonaws.com"}, inst.Attributes)
}

func TestAWSSDProvider_DeleteService(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Description: aws.String("heritage=external-dns,external-dns/owner=owner-id"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
			},
			"srv2": {
				Id:          aws.String("srv2"),
				Description: aws.String("heritage=external-dns,external-dns/owner=owner-id"),
				Name:        aws.String("service2"),
				NamespaceId: aws.String("private"),
			},
			"srv3": {
				Id:          aws.String("srv3"),
				Description: aws.String("heritage=external-dns,external-dns/owner=owner-id,external-dns/resource=virtualservice/grpc-server/validate-grpc-server"),
				Name:        aws.String("service3"),
				NamespaceId: aws.String("private"),
			},
			"srv4": {
				Id:          aws.String("srv4"),
				Description: nil,
				Name:        aws.String("service4"),
				NamespaceId: aws.String("private"),
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "owner-id")

	// delete first service
	err := provider.DeleteService(t.Context(), services["private"]["srv1"])
	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 3)

	// delete third service
	err = provider.DeleteService(t.Context(), services["private"]["srv3"])
	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 2)

	// delete service with no description
	err = provider.DeleteService(t.Context(), services["private"]["srv4"])
	assert.NoError(t, err)

	expected := map[string]*sdtypes.Service{
		"srv2": {
			Id:          aws.String("srv2"),
			Description: aws.String("heritage=external-dns,external-dns/owner=owner-id"),
			Name:        aws.String("service2"),
			NamespaceId: aws.String("private"),
		},
		"srv4": {
			Id:          aws.String("srv4"),
			Description: nil,
			Name:        aws.String("service4"),
			NamespaceId: aws.String("private"),
		},
	}

	assert.Equal(t, expected, api.services["private"])
}

func TestAWSSDProvider_DeleteServiceEmptyDescription_Logging(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Description: nil,
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
			},
		},
	}

	logs := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "owner-id")

	// delete service
	err := provider.DeleteService(t.Context(), services["private"]["srv1"])
	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 1)

	logtest.TestHelperLogContainsWithLogLevel("Skipping service removal \"service1\" because owner id (service.Description) not set, when should be", log.DebugLevel, logs, t)
}

func TestAWSSDProvider_DeleteServiceDryRun(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:          aws.String("srv1"),
				Description: aws.String("heritage=external-dns,external-dns/owner=owner-id"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "owner-id")
	provider.dryRun = true

	// delete first service
	err := provider.DeleteService(t.Context(), services["private"]["srv1"])
	assert.NoError(t, err)
	assert.Len(t, api.services["private"], 1)
}

func TestAWSSDProvider_RegisterInstance(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"a-srv": {
				Id:          aws.String("a-srv"),
				Name:        aws.String("service1"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(60),
					}},
				},
			},
			"cname-srv": {
				Id:          aws.String("cname-srv"),
				Name:        aws.String("service2"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeCname,
						TTL:  aws.Int64(60),
					}},
				},
			},
			"alias-srv": {
				Id:          aws.String("alias-srv"),
				Name:        aws.String("service3"),
				NamespaceId: aws.String("private"),
				DnsConfig: &sdtypes.DnsConfig{
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeA,
						TTL:  aws.Int64(60),
					}},
				},
			},
			"aaaa-srv": {
				Id:          aws.String("aaaa-srv"),
				Name:        aws.String("service4"),
				Description: aws.String("owner-id"),
				DnsConfig: &sdtypes.DnsConfig{
					NamespaceId:   aws.String("private"),
					RoutingPolicy: sdtypes.RoutingPolicyWeighted,
					DnsRecords: []sdtypes.DnsRecord{{
						Type: sdtypes.RecordTypeAaaa,
						TTL:  aws.Int64(100),
					}},
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
		instances:  make(map[string]map[string]*sdtypes.Instance),
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	expectedInstances := make(map[string]*sdtypes.Instance)

	// IPv4-based instance
	err := provider.RegisterInstance(t.Context(), services["private"]["a-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeA,
		DNSName:    "service1.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"1.2.3.4", "1.2.3.5"},
	})
	assert.NoError(t, err)
	expectedInstances["1.2.3.4"] = &sdtypes.Instance{
		Id: aws.String("1.2.3.4"),
		Attributes: map[string]string{
			sdInstanceAttrIPV4: "1.2.3.4",
		},
	}
	expectedInstances["1.2.3.5"] = &sdtypes.Instance{
		Id: aws.String("1.2.3.5"),
		Attributes: map[string]string{
			sdInstanceAttrIPV4: "1.2.3.5",
		},
	}

	// AWS ELB instance (ALIAS)
	err = provider.RegisterInstance(t.Context(), services["private"]["alias-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		DNSName:    "service1.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"load-balancer.us-east-1.elb.amazonaws.com", "load-balancer.us-west-2.elb.amazonaws.com"},
	})
	assert.NoError(t, err)
	expectedInstances["load-balancer.us-east-1.elb.amazonaws.com"] = &sdtypes.Instance{
		Id: aws.String("load-balancer.us-east-1.elb.amazonaws.com"),
		Attributes: map[string]string{
			sdInstanceAttrAlias: "load-balancer.us-east-1.elb.amazonaws.com",
		},
	}
	expectedInstances["load-balancer.us-west-2.elb.amazonaws.com"] = &sdtypes.Instance{
		Id: aws.String("load-balancer.us-west-2.elb.amazonaws.com"),
		Attributes: map[string]string{
			sdInstanceAttrAlias: "load-balancer.us-west-2.elb.amazonaws.com",
		},
	}

	// AWS NLB instance (ALIAS)
	_ = provider.RegisterInstance(t.Context(), services["private"]["alias-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		DNSName:    "service1.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"load-balancer.elb.us-west-2.amazonaws.com"},
	})
	expectedInstances["load-balancer.elb.us-west-2.amazonaws.com"] = &sdtypes.Instance{
		Id: aws.String("load-balancer.elb.us-west-2.amazonaws.com"),
		Attributes: map[string]string{
			sdInstanceAttrAlias: "load-balancer.elb.us-west-2.amazonaws.com",
		},
	}

	// CNAME instance
	_ = provider.RegisterInstance(t.Context(), services["private"]["cname-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeCNAME,
		DNSName:    "service2.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"cname.target.com"},
	})
	expectedInstances["cname.target.com"] = &sdtypes.Instance{
		Id: aws.String("cname.target.com"),
		Attributes: map[string]string{
			sdInstanceAttrCname: "cname.target.com",
		},
	}

	// IPv6-based instance
	provider.RegisterInstance(t.Context(), services["private"]["aaaa-srv"], &endpoint.Endpoint{
		RecordType: endpoint.RecordTypeAAAA,
		DNSName:    "service4.private.com.",
		RecordTTL:  300,
		Targets:    endpoint.Targets{"0000:0000:0000:0000:abcd:abcd:abcd:abcd"},
	})
	expectedInstances["0000:0000:0000:0000:abcd:abcd:abcd:abcd"] = &sdtypes.Instance{
		Id: aws.String("0000:0000:0000:0000:abcd:abcd:abcd:abcd"),
		Attributes: map[string]string{
			sdInstanceAttrIPV6: "0000:0000:0000:0000:abcd:abcd:abcd:abcd",
		},
	}

	// validate instances
	for _, srvInst := range api.instances {
		for id, inst := range srvInst {
			if !reflect.DeepEqual(*expectedInstances[id], *inst) {
				t.Errorf("Instances don't match, expected = %v, actual %v", *expectedInstances[id], *inst)
			}
		}
	}
}

func TestAWSSDProvider_DeregisterInstance(t *testing.T) {
	namespaces := map[string]*sdtypes.Namespace{
		"private": {
			Id:   aws.String("private"),
			Name: aws.String("private.com"),
			Type: sdtypes.NamespaceTypeDnsPrivate,
		},
	}

	services := map[string]map[string]*sdtypes.Service{
		"private": {
			"srv1": {
				Id:   aws.String("srv1"),
				Name: aws.String("service1"),
			},
		},
	}

	instances := map[string]map[string]*sdtypes.Instance{
		"srv1": {
			"1.2.3.4": {
				Id: aws.String("1.2.3.4"),
				Attributes: map[string]string{
					sdInstanceAttrIPV4: "1.2.3.4",
				},
			},
		},
	}

	api := &AWSSDClientStub{
		namespaces: namespaces,
		services:   services,
		instances:  instances,
	}

	provider := newTestAWSSDProvider(api, endpoint.NewDomainFilter([]string{}), "", "")

	_ = provider.DeregisterInstance(t.Context(), services["private"]["srv1"], endpoint.NewEndpoint("srv1.private.com.", endpoint.RecordTypeA, "1.2.3.4"))

	assert.Empty(t, instances["srv1"])
}

func TestAWSSDProvider_awsTags(t *testing.T) {
	tests := []struct {
		Expectation []sdtypes.Tag
		Input       map[string]string
	}{
		{
			Expectation: []sdtypes.Tag{
				{
					Key:   aws.String("key1"),
					Value: aws.String("value1"),
				},
				{
					Key:   aws.String("key2"),
					Value: aws.String("value2"),
				},
			},
			Input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			Expectation: []sdtypes.Tag{},
			Input:       map[string]string{},
		},
		{
			Expectation: []sdtypes.Tag{},
			Input:       nil,
		},
	}

	for _, test := range tests {
		require.ElementsMatch(t, test.Expectation, awsTags(test.Input))
	}
}

func Test_parseNamespace(t *testing.T) {
	tests := []struct {
		name       string
		hostname   string
		namespaces []*sdtypes.NamespaceSummary
		wantNS     string
	}{
		{
			name:     "simple service name",
			hostname: "foo.dev.local",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("dev.local")},
			},
			wantNS: "dev.local",
		},
		{
			name:     "dotted service name",
			hostname: "foo.bar.dev.local",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("dev.local")},
			},
			wantNS: "dev.local",
		},
		{
			name:     "SRV-style hostname",
			hostname: "_tcp.backend.mynet.internal",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("mynet.internal")},
			},
			wantNS: "mynet.internal",
		},
		{
			name:     "longest namespace match wins",
			hostname: "foo.a.b.c",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("b.c")},
				{Name: aws.String("a.b.c")},
			},
			wantNS: "a.b.c",
		},
		{
			name:     "no matching namespace falls back to first-dot split",
			hostname: "foo.unknown.tld",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("dev.local")},
			},
			wantNS: "unknown.tld",
		},
		{
			name:       "empty namespaces falls back to first-dot split",
			hostname:   "foo.bar.baz",
			namespaces: []*sdtypes.NamespaceSummary{},
			wantNS:     "bar.baz",
		},
		{
			name:       "nil namespaces falls back to first-dot split",
			hostname:   "foo.bar.baz",
			namespaces: nil,
			wantNS:     "bar.baz",
		},
		{
			name:     "trailing dot is stripped before matching",
			hostname: "foo.bar.dev.local.",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("dev.local")},
			},
			wantNS: "dev.local",
		},
		{
			name:     "hostname is namespace only, no service prefix",
			hostname: "dev.local",
			namespaces: []*sdtypes.NamespaceSummary{
				{Name: aws.String("dev.local")},
			},
			wantNS: "local",
		},
		{
			name:       "single label hostname, no dots",
			hostname:   "foo",
			namespaces: nil,
			wantNS:     "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotNS := parseNamespace(tc.hostname, tc.namespaces)
			assert.Equal(t, tc.wantNS, gotNS)
		})
	}
}
