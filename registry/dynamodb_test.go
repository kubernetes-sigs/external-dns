/*
Copyright 2023 The Kubernetes Authors.

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

package registry

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

func TestDynamoDBRegistryNew(t *testing.T) {
	api, p := newDynamoDBAPIStub(t, nil)

	_, err := NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "", "", []string{}, []string{}, []byte(""), time.Hour)
	require.NoError(t, err)

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "testPrefix", "", "", []string{}, []string{}, []byte(""), time.Hour)
	require.NoError(t, err)

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "testSuffix", "", []string{}, []string{}, []byte(""), time.Hour)
	require.NoError(t, err)

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "", "testWildcard", []string{}, []string{}, []byte(""), time.Hour)
	require.NoError(t, err)

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "", "testWildcard", []string{}, []string{}, []byte(";k&l)nUC/33:{?d{3)54+,AD?]SX%yh^"), time.Hour)
	require.NoError(t, err)

	_, err = NewDynamoDBRegistry(p, "", api, "test-table", "", "", "", []string{}, []string{}, []byte(""), time.Hour)
	require.EqualError(t, err, "owner id cannot be empty")

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "", "", "", "", []string{}, []string{}, []byte(""), time.Hour)
	require.EqualError(t, err, "table cannot be empty")

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "", "", []string{}, []string{}, []byte(";k&l)nUC/33:{?d{3)54+,AD?]SX%yh^x"), time.Hour)
	require.EqualError(t, err, "the AES Encryption key must have a length of 32 bytes")

	_, err = NewDynamoDBRegistry(p, "test-owner", api, "test-table", "testPrefix", "testSuffix", "", []string{}, []string{}, []byte(""), time.Hour)
	require.EqualError(t, err, "txt-prefix and txt-suffix are mutually exclusive")
}

func TestDynamoDBRegistryRecordsBadTable(t *testing.T) {
	for _, tc := range []struct {
		name     string
		setup    func(desc *dynamodb.TableDescription)
		expected string
	}{
		{
			name: "missing attribute k",
			setup: func(desc *dynamodb.TableDescription) {
				desc.AttributeDefinitions[0].AttributeName = aws.String("wrong")
			},
			expected: "table \"test-table\" must have attribute \"k\" of type \"S\"",
		},
		{
			name: "wrong attribute type",
			setup: func(desc *dynamodb.TableDescription) {
				desc.AttributeDefinitions[0].AttributeType = aws.String("SS")
			},
			expected: "table \"test-table\" attribute \"k\" must have type \"S\"",
		},
		{
			name: "wrong key",
			setup: func(desc *dynamodb.TableDescription) {
				desc.KeySchema[0].AttributeName = aws.String("wrong")
			},
			expected: "table \"test-table\" must have hash key \"k\"",
		},
		{
			name: "has range key",
			setup: func(desc *dynamodb.TableDescription) {
				desc.AttributeDefinitions = append(desc.AttributeDefinitions, &dynamodb.AttributeDefinition{
					AttributeName: aws.String("o"),
					AttributeType: aws.String("S"),
				})
				desc.KeySchema = append(desc.KeySchema, &dynamodb.KeySchemaElement{
					AttributeName: aws.String("o"),
					KeyType:       aws.String("RANGE"),
				})
			},
			expected: "table \"test-table\" must not have a range key",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api, p := newDynamoDBAPIStub(t, nil)
			tc.setup(&api.tableDescription)

			r, _ := NewDynamoDBRegistry(p, "test-owner", api, "test-table", "", "", "", []string{}, []string{}, nil, time.Hour)

			_, err := r.Records(context.Background())
			assert.EqualError(t, err, tc.expected)
		})
	}
}

func TestDynamoDBRegistryRecords(t *testing.T) {
	api, p := newDynamoDBAPIStub(t, nil)

	ctx := context.Background()
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"my-domain.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey:    "test-owner",
				endpoint.ResourceLabelKey: "ingress/default/my-ingress",
			},
		},
		{
			DNSName:       "baz.test-zone.example.org",
			Targets:       endpoint.Targets{"1.1.1.1"},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "set-1",
			Labels: map[string]string{
				endpoint.OwnerLabelKey:    "test-owner",
				endpoint.ResourceLabelKey: "ingress/default/my-ingress",
			},
		},
		{
			DNSName:       "baz.test-zone.example.org",
			Targets:       endpoint.Targets{"2.2.2.2"},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "set-2",
			Labels: map[string]string{
				endpoint.OwnerLabelKey:    "test-owner",
				endpoint.ResourceLabelKey: "ingress/default/other-ingress",
			},
		},
		{
			DNSName:       "migrate.test-zone.example.org",
			Targets:       endpoint.Targets{"3.3.3.3"},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "set-3",
			Labels: map[string]string{
				endpoint.OwnerLabelKey:    "test-owner",
				endpoint.ResourceLabelKey: "ingress/default/other-ingress",
			},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  dynamodbAttributeMigrate,
					Value: "true",
				},
			},
		},
		{
			DNSName:       "txt.orphaned.test-zone.example.org",
			Targets:       endpoint.Targets{"\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/other-ingress\""},
			RecordType:    endpoint.RecordTypeTXT,
			SetIdentifier: "set-3",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "test-owner",
			},
		},
		{
			DNSName:       "txt.baz.test-zone.example.org",
			Targets:       endpoint.Targets{"\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/other-ingress\""},
			RecordType:    endpoint.RecordTypeTXT,
			SetIdentifier: "set-2",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "test-owner",
			},
		},
	}

	r, _ := NewDynamoDBRegistry(p, "test-owner", api, "test-table", "txt.", "", "", []string{}, []string{}, nil, time.Hour)
	_ = p.(*wrappedProvider).Provider.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("migrate.test-zone.example.org", endpoint.RecordTypeA, "3.3.3.3").WithSetIdentifier("set-3"),
			endpoint.NewEndpoint("txt.migrate.test-zone.example.org", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/other-ingress\"").WithSetIdentifier("set-3"),
			endpoint.NewEndpoint("txt.orphaned.test-zone.example.org", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/other-ingress\"").WithSetIdentifier("set-3"),
			endpoint.NewEndpoint("txt.baz.test-zone.example.org", endpoint.RecordTypeTXT, "\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/other-ingress\"").WithSetIdentifier("set-2"),
		},
	})

	records, err := r.Records(ctx)
	require.Nil(t, err)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func TestDynamoDBRegistryApplyChanges(t *testing.T) {
	for _, tc := range []struct {
		name            string
		maxBatchSize    uint8
		stubConfig      DynamoDBStubConfig
		addRecords      []*endpoint.Endpoint
		changes         plan.Changes
		expectedError   string
		expectedRecords []*endpoint.Endpoint
	}{
		{
			name: "create",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "new.test-zone.example.org",
						Targets:       endpoint.Targets{"new.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectInsert: map[string]map[string]string{
					"new.test-zone.example.org#CNAME#set-new": {endpoint.ResourceLabelKey: "ingress/default/new-ingress"},
				},
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
				{
					DNSName:       "new.test-zone.example.org",
					Targets:       endpoint.Targets{"new.loadbalancer.com"},
					RecordType:    endpoint.RecordTypeCNAME,
					SetIdentifier: "set-new",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new-ingress",
					},
				},
			},
		},
		{
			name:         "create more entries than DynamoDB batch size limit",
			maxBatchSize: 2,
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "new1.test-zone.example.org",
						Targets:       endpoint.Targets{"new1.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new1-ingress",
						},
					},
					{
						DNSName:       "new2.test-zone.example.org",
						Targets:       endpoint.Targets{"new2.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new2-ingress",
						},
					},
					{
						DNSName:       "new3.test-zone.example.org",
						Targets:       endpoint.Targets{"new3.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new3-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectInsert: map[string]map[string]string{
					"new1.test-zone.example.org#CNAME#set-new": {endpoint.ResourceLabelKey: "ingress/default/new1-ingress"},
					"new2.test-zone.example.org#CNAME#set-new": {endpoint.ResourceLabelKey: "ingress/default/new2-ingress"},
					"new3.test-zone.example.org#CNAME#set-new": {endpoint.ResourceLabelKey: "ingress/default/new3-ingress"},
				},
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
				{
					DNSName:       "new1.test-zone.example.org",
					Targets:       endpoint.Targets{"new1.loadbalancer.com"},
					RecordType:    endpoint.RecordTypeCNAME,
					SetIdentifier: "set-new",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new1-ingress",
					},
				},
				{
					DNSName:       "new2.test-zone.example.org",
					Targets:       endpoint.Targets{"new2.loadbalancer.com"},
					RecordType:    endpoint.RecordTypeCNAME,
					SetIdentifier: "set-new",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new2-ingress",
					},
				},
				{
					DNSName:       "new3.test-zone.example.org",
					Targets:       endpoint.Targets{"new3.loadbalancer.com"},
					RecordType:    endpoint.RecordTypeCNAME,
					SetIdentifier: "set-new",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new3-ingress",
					},
				},
			},
		},
		{
			name: "create orphaned",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "quux.test-zone.example.org",
						Targets:       endpoint.Targets{"5.5.5.5"},
						RecordType:    endpoint.RecordTypeA,
						SetIdentifier: "set-2",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/quux-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
				{
					DNSName:       "quux.test-zone.example.org",
					Targets:       endpoint.Targets{"5.5.5.5"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/quux-ingress",
					},
				},
			},
		},
		{
			name: "create orphaned change",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "quux.test-zone.example.org",
						Targets:       endpoint.Targets{"5.5.5.5"},
						RecordType:    endpoint.RecordTypeA,
						SetIdentifier: "set-2",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectUpdate: map[string]map[string]string{
					"quux.test-zone.example.org#A#set-2": {endpoint.ResourceLabelKey: "ingress/default/new-ingress"},
				},
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
				{
					DNSName:       "quux.test-zone.example.org",
					Targets:       endpoint.Targets{"5.5.5.5"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new-ingress",
					},
				},
			},
		},
		{
			name: "create duplicate",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "new.test-zone.example.org",
						Targets:       endpoint.Targets{"new.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectInsertError: map[string]string{
					"new.test-zone.example.org#CNAME#set-new": "DuplicateItem",
				},
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
		{
			name: "create error",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:       "new.test-zone.example.org",
						Targets:       endpoint.Targets{"new.loadbalancer.com"},
						RecordType:    endpoint.RecordTypeCNAME,
						SetIdentifier: "set-new",
						Labels: map[string]string{
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectInsertError: map[string]string{
					"new.test-zone.example.org#CNAME#set-new": "TestingError",
				},
			},
			expectedError: "inserting dynamodb record \"new.test-zone.example.org#CNAME#set-new\": TestingError: testing error",
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
		{
			name: "update",
			changes: plan.Changes{
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"new-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"new-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
		{
			name: "update change",
			changes: plan.Changes{
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"new-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
				ExpectUpdate: map[string]map[string]string{
					"bar.test-zone.example.org#CNAME#": {endpoint.ResourceLabelKey: "ingress/default/new-ingress"},
				},
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"new-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
		{
			name: "update migrate",
			addRecords: []*endpoint.Endpoint{
				{
					DNSName:       "txt.bar.test-zone.example.org",
					Targets:       endpoint.Targets{"\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/new-ingress\""},
					RecordType:    endpoint.RecordTypeTXT,
					SetIdentifier: "set-1",
				},
			},
			changes: plan.Changes{
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
						ProviderSpecific: endpoint.ProviderSpecific{
							{
								Name:  dynamodbAttributeMigrate,
								Value: "true",
							},
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectDelete: sets.New("quux.test-zone.example.org#A#set-2"),
				ExpectInsert: map[string]map[string]string{
					"bar.test-zone.example.org#CNAME#": {endpoint.ResourceLabelKey: "ingress/default/new-ingress"},
				},
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/new-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
				{
					DNSName:       "txt.bar.test-zone.example.org",
					Targets:       endpoint.Targets{"\"heritage=external-dns,external-dns/owner=test-owner,external-dns/resource=ingress/default/new-ingress\""},
					RecordType:    endpoint.RecordTypeTXT,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "test-owner",
					},
				},
			},
		},
		{
			name: "update error",
			changes: plan.Changes{
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"new-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/new-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectUpdateError: map[string]string{
					"bar.test-zone.example.org#CNAME#": "TestingError",
				},
			},
			expectedError: "updating dynamodb record \"bar.test-zone.example.org#CNAME#\": TestingError: testing error",
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:    "bar.test-zone.example.org",
					Targets:    endpoint.Targets{"my-domain.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
		{
			name: "delete",
			changes: plan.Changes{
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "bar.test-zone.example.org",
						Targets:    endpoint.Targets{"my-domain.com"},
						RecordType: endpoint.RecordTypeCNAME,
						Labels: map[string]string{
							endpoint.OwnerLabelKey:    "test-owner",
							endpoint.ResourceLabelKey: "ingress/default/my-ingress",
						},
					},
				},
			},
			stubConfig: DynamoDBStubConfig{
				ExpectDelete: sets.New("bar.test-zone.example.org#CNAME#", "quux.test-zone.example.org#A#set-2"),
			},
			expectedRecords: []*endpoint.Endpoint{
				{
					DNSName:    "foo.test-zone.example.org",
					Targets:    endpoint.Targets{"foo.loadbalancer.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: map[string]string{
						endpoint.OwnerLabelKey: "",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"1.1.1.1"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-1",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/my-ingress",
					},
				},
				{
					DNSName:       "baz.test-zone.example.org",
					Targets:       endpoint.Targets{"2.2.2.2"},
					RecordType:    endpoint.RecordTypeA,
					SetIdentifier: "set-2",
					Labels: map[string]string{
						endpoint.OwnerLabelKey:    "test-owner",
						endpoint.ResourceLabelKey: "ingress/default/other-ingress",
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			originalMaxBatchSize := dynamodbMaxBatchSize
			if tc.maxBatchSize > 0 {
				dynamodbMaxBatchSize = tc.maxBatchSize
			}

			api, p := newDynamoDBAPIStub(t, &tc.stubConfig)
			if len(tc.addRecords) > 0 {
				_ = p.(*wrappedProvider).Provider.ApplyChanges(context.Background(), &plan.Changes{
					Create: tc.addRecords,
				})
			}

			ctx := context.Background()

			r, _ := NewDynamoDBRegistry(p, "test-owner", api, "test-table", "txt.", "", "", []string{}, []string{}, nil, time.Hour)
			_, err := r.Records(ctx)
			require.Nil(t, err)

			err = r.ApplyChanges(ctx, &tc.changes)
			if tc.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}

			assert.Empty(t, tc.stubConfig.ExpectInsert, "all expected inserts made")
			assert.Empty(t, tc.stubConfig.ExpectDelete, "all expected deletions made")

			records, err := r.Records(ctx)
			require.Nil(t, err)
			assert.True(t, testutils.SameEndpoints(records, tc.expectedRecords))

			r.recordsCache = nil
			records, err = r.Records(ctx)
			require.Nil(t, err)
			assert.True(t, testutils.SameEndpoints(records, tc.expectedRecords))
			if tc.expectedError == "" {
				assert.Empty(t, r.orphanedLabels)
			}

			dynamodbMaxBatchSize = originalMaxBatchSize
		})
	}
}

// DynamoDBAPIStub is a minimal implementation of DynamoDBAPI, used primarily for unit testing.
type DynamoDBStub struct {
	t                *testing.T
	stubConfig       *DynamoDBStubConfig
	tableDescription dynamodb.TableDescription
	changesApplied   bool
}

type DynamoDBStubConfig struct {
	ExpectInsert      map[string]map[string]string
	ExpectInsertError map[string]string
	ExpectUpdate      map[string]map[string]string
	ExpectUpdateError map[string]string
	ExpectDelete      sets.Set[string]
}

type wrappedProvider struct {
	provider.Provider
	stub *DynamoDBStub
}

func (w *wrappedProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	assert.False(w.stub.t, w.stub.changesApplied, "ApplyChanges already called")
	w.stub.changesApplied = true
	return w.Provider.ApplyChanges(ctx, changes)
}

func newDynamoDBAPIStub(t *testing.T, stubConfig *DynamoDBStubConfig) (*DynamoDBStub, provider.Provider) {
	stub := &DynamoDBStub{
		t:          t,
		stubConfig: stubConfig,
		tableDescription: dynamodb.TableDescription{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("k"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("k"),
					KeyType:       aws.String("HASH"),
				},
			},
		},
	}
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone(testZone)
	_ = p.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("foo.test-zone.example.org", endpoint.RecordTypeCNAME, "foo.loadbalancer.com"),
			endpoint.NewEndpoint("bar.test-zone.example.org", endpoint.RecordTypeCNAME, "my-domain.com"),
			endpoint.NewEndpoint("baz.test-zone.example.org", endpoint.RecordTypeA, "1.1.1.1").WithSetIdentifier("set-1"),
			endpoint.NewEndpoint("baz.test-zone.example.org", endpoint.RecordTypeA, "2.2.2.2").WithSetIdentifier("set-2"),
		},
	})
	return stub, &wrappedProvider{
		Provider: p,
		stub:     stub,
	}
}

func (r *DynamoDBStub) DescribeTableWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	assert.NotNil(r.t, ctx)
	assert.Equal(r.t, "test-table", *input.TableName, "table name")
	return &dynamodb.DescribeTableOutput{
		Table: &r.tableDescription,
	}, nil
}

func (r *DynamoDBStub) ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	assert.NotNil(r.t, ctx)
	assert.Equal(r.t, "test-table", *input.TableName, "table name")
	assert.Equal(r.t, "o = :ownerval", *input.FilterExpression)
	assert.Len(r.t, input.ExpressionAttributeValues, 1)
	assert.Equal(r.t, "test-owner", *input.ExpressionAttributeValues[":ownerval"].S)
	assert.Equal(r.t, "k,l", *input.ProjectionExpression)
	assert.True(r.t, *input.ConsistentRead)
	fn(&dynamodb.ScanOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			{
				"k": &dynamodb.AttributeValue{S: aws.String("bar.test-zone.example.org#CNAME#")},
				"l": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					endpoint.ResourceLabelKey: {S: aws.String("ingress/default/my-ingress")},
				}},
			},
			{
				"k": &dynamodb.AttributeValue{S: aws.String("baz.test-zone.example.org#A#set-1")},
				"l": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					endpoint.ResourceLabelKey: {S: aws.String("ingress/default/my-ingress")},
				}},
			},
			{
				"k": &dynamodb.AttributeValue{S: aws.String("baz.test-zone.example.org#A#set-2")},
				"l": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					endpoint.ResourceLabelKey: {S: aws.String("ingress/default/other-ingress")},
				}},
			},
			{
				"k": &dynamodb.AttributeValue{S: aws.String("quux.test-zone.example.org#A#set-2")},
				"l": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					endpoint.ResourceLabelKey: {S: aws.String("ingress/default/quux-ingress")},
				}},
			},
		},
	}, true)
	return nil
}

func (r *DynamoDBStub) BatchExecuteStatementWithContext(context aws.Context, input *dynamodb.BatchExecuteStatementInput, option ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	assert.NotNil(r.t, context)
	hasDelete := strings.HasPrefix(strings.ToLower(aws.StringValue(input.Statements[0].Statement)), "delete")
	assert.Equal(r.t, hasDelete, r.changesApplied, "delete after provider changes, everything else before")
	assert.LessOrEqual(r.t, len(input.Statements), 25)
	responses := make([]*dynamodb.BatchStatementResponse, 0, len(input.Statements))

	for _, statement := range input.Statements {
		assert.Equal(r.t, hasDelete, strings.HasPrefix(strings.ToLower(aws.StringValue(statement.Statement)), "delete"))
		switch aws.StringValue(statement.Statement) {
		case "DELETE FROM \"test-table\" WHERE \"k\"=? AND \"o\"=?":
			assert.True(r.t, r.changesApplied, "unexpected delete before provider changes")

			key := aws.StringValue(statement.Parameters[0].S)
			assert.True(r.t, r.stubConfig.ExpectDelete.Has(key), "unexpected delete for key %q", key)
			r.stubConfig.ExpectDelete.Delete(key)

			assert.Equal(r.t, "test-owner", aws.StringValue(statement.Parameters[1].S))

			responses = append(responses, &dynamodb.BatchStatementResponse{})

		case "INSERT INTO \"test-table\" VALUE {'k':?, 'o':?, 'l':?}":
			assert.False(r.t, r.changesApplied, "unexpected insert after provider changes")

			key := aws.StringValue(statement.Parameters[0].S)
			if code, exists := r.stubConfig.ExpectInsertError[key]; exists {
				delete(r.stubConfig.ExpectInsertError, key)
				responses = append(responses, &dynamodb.BatchStatementResponse{
					Error: &dynamodb.BatchStatementError{
						Code:    aws.String(code),
						Message: aws.String("testing error"),
					},
				})
				break
			}

			expectedLabels, found := r.stubConfig.ExpectInsert[key]
			assert.True(r.t, found, "unexpected insert for key %q", key)
			delete(r.stubConfig.ExpectInsert, key)

			assert.Equal(r.t, "test-owner", aws.StringValue(statement.Parameters[1].S))

			for label, attribute := range statement.Parameters[2].M {
				value := aws.StringValue(attribute.S)
				expectedValue, found := expectedLabels[label]
				assert.True(r.t, found, "insert for key %q has unexpected label %q", key, label)
				delete(expectedLabels, label)
				assert.Equal(r.t, expectedValue, value, "insert for key %q label %q value", key, label)
			}

			for label := range expectedLabels {
				r.t.Errorf("insert for key %q did not get expected label %q", key, label)
			}

			responses = append(responses, &dynamodb.BatchStatementResponse{})

		case "UPDATE \"test-table\" SET \"l\"=? WHERE \"k\"=?":
			assert.False(r.t, r.changesApplied, "unexpected update after provider changes")

			key := aws.StringValue(statement.Parameters[1].S)
			if code, exists := r.stubConfig.ExpectUpdateError[key]; exists {
				delete(r.stubConfig.ExpectInsertError, key)
				responses = append(responses, &dynamodb.BatchStatementResponse{
					Error: &dynamodb.BatchStatementError{
						Code:    aws.String(code),
						Message: aws.String("testing error"),
					},
				})
				break
			}

			expectedLabels, found := r.stubConfig.ExpectUpdate[key]
			assert.True(r.t, found, "unexpected update for key %q", key)
			delete(r.stubConfig.ExpectUpdate, key)

			for label, attribute := range statement.Parameters[0].M {
				value := aws.StringValue(attribute.S)
				expectedValue, found := expectedLabels[label]
				assert.True(r.t, found, "update for key %q has unexpected label %q", key, label)
				delete(expectedLabels, label)
				assert.Equal(r.t, expectedValue, value, "update for key %q label %q value", key, label)
			}

			for label := range expectedLabels {
				r.t.Errorf("update for key %q did not get expected label %q", key, label)
			}

			responses = append(responses, &dynamodb.BatchStatementResponse{})

		default:
			r.t.Errorf("unexpected statement: %s", aws.StringValue(statement.Statement))
		}
	}

	return &dynamodb.BatchExecuteStatementOutput{
		Responses: responses,
	}, nil
}
