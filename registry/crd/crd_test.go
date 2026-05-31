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

package registry

import (
	"context"
	"fmt"
	"maps"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/registry"
)

type CRDSuite struct {
	suite.Suite
}

func (suite *CRDSuite) SetupTest() {
}

// The endpoints needs to be part of the zone otherwise it will be filtered out.
func inMemoryProviderWithEntries(t *testing.T, ctx context.Context, zone string, endpoints ...*endpoint.Endpoint) *inmemory.InMemoryProvider {
	p := inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones([]string{zone}))

	err := p.ApplyChanges(ctx, &plan.Changes{
		Create: endpoints,
	})
	if err != nil {
		t.Fatal("Could not create an in memory provider", err)
	}

	return p
}

func TestCRDSource(t *testing.T) {
	suite.Run(t, new(CRDSuite))
	t.Run("Interface", testCRDSourceImplementsSource)
	t.Run("Constructor", testConstructor)
	t.Run("Records", testRecords)
}

// testCRDSourceImplementsSource tests that crdSource is a valid Source.
func testCRDSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*registry.Registry)(nil), new(CRDRegistry))
}

func testConstructor(t *testing.T) {
	_, err := NewCRDRegistry(nil, "", "", "v1", "", "", time.Second, time.Second)
	assert.Error(t, err, "Expected a new registry to return an error when no ownerID are specified")

	_, err = NewCRDRegistry(nil, "/dev/null", "", "v1", "default", "ownerID", time.Second, time.Second)
	assert.Error(t, err, err.Error()+"Expected a new registry to return an error when there is no kubeconfig")

	_, err = NewCRDRegistry(nil, "", "####", "v1", "default", "ownerID", time.Second, time.Second)
	assert.Error(t, err, err.Error()+"Expected a new registry to return an error when there is an invalid url")
}

func testRecords(t *testing.T) {
	ctx := t.Context()
	t.Run("use the cache if within the time interval", func(t *testing.T) {
		registry := &CRDRegistry{
			recordsCacheRefreshTime: time.Now(),
			cacheInterval:           time.Hour,
			recordsCache: []*endpoint.Endpoint{{
				DNSName:    "cached.mytestdomain.io",
				RecordType: "A",
				Targets:    []string{"127.0.0.1"},
			}},
		}
		endpoints, err := registry.Records(ctx)
		if err != nil {
			t.Error(err)
		}

		if len(endpoints) != 1 {
			t.Error("expected only 1 record from the cache, got: ", len(endpoints))
		}

		if endpoints[0].DNSName != "cached.mytestdomain.io" {
			t.Error("expected DNS Name to be the cached value got: ", endpoints[0].DNSName)
		}
	})

	t.Run("Use k8s api to get records from the registry", func(t *testing.T) {
		// Setup the provider and the mock client for the CRD so that mytestdomain.io can be
		// found on both the provider and the CRD
		provider := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", &endpoint.Endpoint{
			DNSName:       "sub.mytestdomain.io",
			RecordType:    "CNAME",
			SetIdentifier: "myid-1",
		})

		responses := []mockResult{{
			request: mockRequest{
				method:    "GET",
				namespace: "default",
			},
			response: &mockResponse{
				content: apiv1alpha1.DNSRecordList{
					Items: []apiv1alpha1.DNSRecord{{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								apiv1alpha1.RecordOwnerLabel:         "test",
								apiv1alpha1.RecordNameLabel:          "sub.mytestdomain.io",
								apiv1alpha1.RecordTypeLabel:          "CNAME",
								apiv1alpha1.RecordSetIdentifierLabel: "myid-1",
							},
						},
						Spec: apiv1alpha1.DNSRecordSpec{
							Endpoint: endpoint.Endpoint{
								DNSName:       "sub.mytestdomain.io",
								RecordType:    "CNAME",
								SetIdentifier: "myid-1",
								Labels: map[string]string{
									endpoint.ResourceLabelKey: "resource",
								},
							},
						},
					}},
				},
			},
		}}

		client := newMockCRDClient("default", responses...)

		registry := &CRDRegistry{
			provider:  provider,
			namespace: "default",
			client:    client,
			ownerID:   "test",
		}

		// The test
		endpoints, err := registry.Records(ctx)
		if err != nil {
			t.Error(err)
		}

		if len(endpoints) != 1 {
			t.Errorf("expected only 1 endpoint, got %d", len(endpoints))
		}

		if endpoints[0].Labels[endpoint.ResourceLabelKey] != "resource" {
			t.Errorf("endpoint doesn't include the label from the registry: %#v", endpoints[0].Labels)
		}
	})
}

func TestCRDApplyChangesInMemory(t *testing.T) {
	ctx := t.Context()

	testCases := []struct {
		ChangeType string // One of Create, Update, Delete
		Endpoint   *endpoint.Endpoint
		AssertFn   func(c *mockClient)
	}{
		{
			ChangeType: "Create",
			Endpoint: &endpoint.Endpoint{
				DNSName:       "sub.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-1",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels: map[string]string{
					endpoint.OwnerLabelKey: "test",
				},
			},
			AssertFn: func(c *mockClient) {
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{method: "POST", namespace: "default"}))
				assert.True(t, executed)
			},
		},
		{
			ChangeType: "Delete",
			Endpoint: &endpoint.Endpoint{
				DNSName:       "to.be.deleted.mytestdomain.io",
				RecordType:    "A",
				SetIdentifier: "myid-2",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels: map[string]string{
					endpoint.OwnerLabelKey: "test",
				},
			},
			AssertFn: func(c *mockClient) {
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{method: "DELETE", namespace: "default", name: "test-myid-2"}))
				assert.True(t, executed)
			},
		},
		{
			ChangeType: "Update",
			Endpoint: &endpoint.Endpoint{
				DNSName:       "to.be.updated.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-3",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels: map[string]string{
					endpoint.OwnerLabelKey: "test",
				},
			},
			AssertFn: func(c *mockClient) {
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{method: "PUT", namespace: "default", name: "test-myid-3"}))
				assert.True(t, executed)
			},
		},
	}

	for _, testCase := range testCases {
		var seedEndpoints []*endpoint.Endpoint
		var changes plan.Changes
		var responses []mockResult
		switch testCase.ChangeType {
		case "Create":
			dnsrecord := apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "sub-mytestdomain-io",
					Namespace: "default",
					Labels: map[string]string{
						apiv1alpha1.RecordOwnerLabel:         "test",
						apiv1alpha1.RecordNameLabel:          "sub.mytestdomain.io",
						apiv1alpha1.RecordTypeLabel:          "CNAME",
						apiv1alpha1.RecordSetIdentifierLabel: "myid-1",
					},
				},
				Spec: apiv1alpha1.DNSRecordSpec{Endpoint: *testCase.Endpoint},
			}
			responses = append(responses, mockResult{
				request:  mockRequest{method: "POST", namespace: "default"},
				response: &mockResponse{content: dnsrecord},
			}, mockResult{
				request:  mockRequest{method: "GET", namespace: "default"},
				response: &mockResponse{content: apiv1alpha1.DNSRecordList{Items: []apiv1alpha1.DNSRecord{dnsrecord}}},
			})
			changes.Create = []*endpoint.Endpoint{testCase.Endpoint}

		case "Delete":
			dnsrecord := apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{Name: "test-myid-2", Namespace: "default"},
				Spec:       apiv1alpha1.DNSRecordSpec{Endpoint: *testCase.Endpoint},
			}
			responses = append(responses, mockResult{
				request:  mockRequest{method: "GET", namespace: "default"},
				response: &mockResponse{content: apiv1alpha1.DNSRecordList{Items: []apiv1alpha1.DNSRecord{dnsrecord}}},
			}, mockResult{
				request:  mockRequest{method: "DELETE", name: "test-myid-2", namespace: "default"},
				response: &mockResponse{},
			})

			changes.Delete = []*endpoint.Endpoint{testCase.Endpoint}
			seedEndpoints = append(seedEndpoints, testCase.Endpoint)
		case "Update":
			dnsrecord := apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{Name: "test-myid-3", Namespace: "default"},
				Spec:       apiv1alpha1.DNSRecordSpec{Endpoint: *testCase.Endpoint},
			}
			responses = append(responses, mockResult{
				request:  mockRequest{method: "GET", namespace: "default"},
				response: &mockResponse{content: apiv1alpha1.DNSRecordList{Items: []apiv1alpha1.DNSRecord{dnsrecord}}},
			}, mockResult{
				request:  mockRequest{method: "PUT", name: "test-myid-3", namespace: "default"},
				response: &mockResponse{},
			})

			changes.UpdateNew = []*endpoint.Endpoint{testCase.Endpoint}
			changes.UpdateOld = []*endpoint.Endpoint{testCase.Endpoint}
			seedEndpoints = append(seedEndpoints, testCase.Endpoint)
		default:
			t.Errorf("ChangeType not defined: %s", testCase.ChangeType)
		}

		provider := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", seedEndpoints...)
		client := newMockCRDClient("default", responses...)
		registry := &CRDRegistry{
			provider:  provider,
			namespace: "default",
			client:    client,
			ownerID:   "test",
		}

		err := registry.ApplyChanges(ctx, &changes)
		if err != nil {
			t.Error(fmt.Errorf("apply changes failed on %s: %w", testCase.ChangeType, err))
		}

		if testCase.AssertFn != nil {
			testCase.AssertFn(client)
		}

	}
}

type mockProvider struct {
	records  []*endpoint.Endpoint
	addLabel bool
}

func (m *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	return m.records, nil
}

// This applychanges in mocked provider simulate when
// the provider change Labels of the records.
func (m *mockProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	endpoints := changes.Create
	m.records = make([]*endpoint.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		newEp := endpoint.NewEndpointWithTTL(ep.DNSName, ep.RecordType, ep.RecordTTL, ep.Targets...).WithSetIdentifier(ep.SetIdentifier)
		newEp.Labels = endpoint.NewLabels()
		maps.Copy(newEp.Labels, ep.Labels)
		newEp.ProviderSpecific = append(endpoint.ProviderSpecific(nil), ep.ProviderSpecific...)
		// mocked specific change
		if m.addLabel {
			newEp.Labels["prefix"] = "random"
		}
		m.records = append(m.records, newEp)
	}
	return nil
}

func (m *mockProvider) AdjustEndpoints(eps []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return eps, nil
}

func (m *mockProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return &endpoint.DomainFilter{}
}

// Ensure record is updated with expected Labels even when they are modified
// by the provider.ApplyChanges like, for instance, with coredns provider
func TestCRDApplyChangesMockedProvider(t *testing.T) {
	testCases := []struct {
		provider provider.Provider
		assertFn func(c *mockClient)
	}{
		{
			provider: &mockProvider{addLabel: true},
			assertFn: func(c *mockClient) {
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{method: "PUT", name: "sub-mytestdomain-io", namespace: "default"}))
				assert.True(t, executed)
			},
		},
		{
			provider: &mockProvider{addLabel: false},
			assertFn: func(c *mockClient) {
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{method: "PUT", name: "sub-mytestdomain-io", namespace: "default"}))
				assert.False(t, executed)
			},
		},
	}

	for _, testCase := range testCases {

		var changes plan.Changes
		var responses []mockResult
		ep := endpoint.Endpoint{
			DNSName:       "sub.mytestdomain.io",
			RecordType:    "CNAME",
			SetIdentifier: "myid-1",
			Targets:       endpoint.NewTargets("127.0.0.1"),
			Labels:        map[string]string{endpoint.OwnerLabelKey: "owner"},
		}
		dnsrecord := apiv1alpha1.DNSRecord{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sub-mytestdomain-io",
				Namespace: "default",
			},
			Spec: apiv1alpha1.DNSRecordSpec{Endpoint: ep},
		}
		responses = append(responses, mockResult{
			request:  mockRequest{method: "POST", namespace: "default"},
			response: &mockResponse{},
		}, mockResult{
			request:  mockRequest{method: "GET", namespace: "default"},
			response: &mockResponse{content: apiv1alpha1.DNSRecordList{Items: []apiv1alpha1.DNSRecord{dnsrecord}}},
		}, mockResult{
			request:  mockRequest{method: "PUT", name: "sub-mytestdomain-io", namespace: "default"},
			response: &mockResponse{},
		})
		changes.Create = []*endpoint.Endpoint{&ep}

		client := newMockCRDClient("default", responses...)
		registry := &CRDRegistry{
			provider:  testCase.provider,
			namespace: "default",
			client:    client,
			ownerID:   "owner",
		}

		err := registry.ApplyChanges(t.Context(), &changes)
		require.NoError(t, err)

		testCase.assertFn(client)
	}
}

// Mocking tools for the CRD. These are attempt at generating all the right struct
// for the test to be able to simulate requests and return proper objects so that
// tests can be conducted in a controlled fashion to exercise specific behavior without
// requiring a fully configured Kubernetes client. Hopefully it makes it easier to write tests
// while giving enough confidence that the tests represents real-life scenarios.
type mockResult struct {
	request  mockRequest
	response CRDResult
}

type mockClient struct {
	namespace     string
	mockResponses map[mockRequestKey]CRDResult
	requestHit    map[mockRequestKey]struct{}
}

func newMockCRDClient(namespace string, responses ...mockResult) *mockClient {
	mockResponses := map[mockRequestKey]CRDResult{}
	for _, r := range responses {
		mockResponses[keyFromRequest(&r.request)] = r.response
	}

	return &mockClient{
		namespace:     namespace,
		mockResponses: mockResponses,
		requestHit:    map[mockRequestKey]struct{}{},
	}
}

func (mc *mockClient) RequestWasExecuted(key mockRequestKey) bool {
	_, found := mc.requestHit[key]
	return found
}

func (m *mockClient) MockResponses() {
}

func (m *mockClient) Get() CRDRequest {
	return &mockRequest{c: m, namespace: m.namespace, method: "GET"}
}

func (m *mockClient) List() CRDRequest {
	return &mockRequest{c: m, namespace: m.namespace, method: "GET"}
}

func (m *mockClient) Put() CRDRequest {
	return &mockRequest{c: m, namespace: m.namespace, method: "PUT"}
}

func (m *mockClient) Post() CRDRequest {
	return &mockRequest{c: m, namespace: m.namespace, method: "POST"}
}

func (m *mockClient) Delete() CRDRequest {
	return &mockRequest{c: m, namespace: m.namespace, method: "DELETE"}
}

type mockRequestKey struct {
	method    string
	namespace string
	name      string
}

func keyFromRequest(mr *mockRequest) mockRequestKey {
	return mockRequestKey{
		method:    mr.method,
		name:      mr.name,
		namespace: mr.namespace,
	}
}

type mockRequest struct {
	c         *mockClient
	method    string
	namespace string
	name      string
}

func (mr *mockRequest) Name(name string) CRDRequest {
	mr.name = name
	return mr
}

func (mr *mockRequest) Namespace(namespace string) CRDRequest {
	mr.namespace = namespace
	return mr
}

func (mr *mockRequest) Body(any) CRDRequest {
	return mr
}

func (mr *mockRequest) Params(runtime.Object) CRDRequest {
	return mr
}

func (mr *mockRequest) Do(_ context.Context) CRDResult {
	key := keyFromRequest(mr)
	if response, found := mr.c.mockResponses[key]; found {
		mr.c.requestHit[key] = struct{}{}
		return response
	}

	return &mockErrorResponse{request: mr}
}

type mockErrorResponse struct {
	request *mockRequest
}

func (mr *mockErrorResponse) Error() error {
	return fmt.Errorf("Request wasn't mocked: %+v", mr.request)
}

func (mr *mockErrorResponse) Into(_ runtime.Object) error {
	return fmt.Errorf("Request wasn't mocked: %+v", mr.request)
}

type mockResponse struct {
	content any
}

func (mr *mockResponse) Error() error {
	return nil
}

func (mr *mockResponse) Into(obj runtime.Object) error {
	reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(mr.content))
	return nil
}
