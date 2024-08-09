package registry

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/external-dns/crds"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/inmemory"
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
	require.Implements(t, (*Registry)(nil), new(CRDRegistry))
}

func testConstructor(t *testing.T) {
	_, err := NewCRDRegistry(nil, nil, "", time.Second, "default")
	if err == nil {
		t.Error("Expected a new registry to return an error when no ownerID are specified")
	}

	_, err = NewCRDRegistry(nil, nil, "ownerID", time.Second, "namespace")
	if err != nil {
		t.Error("Expected registry to be initialized without error when providing an owner id and a namespace", err)
	}
}

func testRecords(t *testing.T) {
	ctx := context.Background()
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

	t.Run("ALIAS records are converted to CNAME", func(t *testing.T) {
		e := []*endpoint.Endpoint{
			{
				DNSName:    "foo.mytestdomain.io",
				RecordType: "A",
				Targets:    []string{"127.0.0.1"},
				ProviderSpecific: []endpoint.ProviderSpecificProperty{{
					Name:  "alias",
					Value: "true",
				}},
			},
		}
		provider := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", e...)
		responses := []mockResult{{
			request: mockRequest{
				method:    "GET",
				namespace: "default",
			},
			response: &mockResponse{
				content: crds.DNSEntryList{
					Items: []crds.DNSEntry{{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								crds.RegistryResourceLabel: "some-value",
							},
						},
						Spec: crds.DNSEntrySpec{
							Endpoint: *e[0],
						},
					}},
				},
			},
		}}

		registry := &CRDRegistry{
			provider:  provider,
			namespace: "default",
			client:    NewMockCRDClient("default", responses...),
			ownerID:   "test",
		}

		endpoints, err := registry.Records(ctx)
		if err != nil {
			t.Error(err)
		}

		if endpoints[0].RecordType != "CNAME" {
			t.Error("Expected record type to be changed from ALIAS to CNAME: ", endpoints[0].RecordType)
		}
	})

	t.Run("Add existing labels from registry to the record from the provider", func(t *testing.T) {
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
				content: crds.DNSEntryList{
					Items: []crds.DNSEntry{{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								crds.RegistryResourceLabel: "some-value",
							},
						},
						Spec: crds.DNSEntrySpec{
							Endpoint: endpoint.Endpoint{
								DNSName:       "sub.mytestdomain.io",
								RecordType:    "CNAME",
								SetIdentifier: "myid-1",
							},
						},
					}},
				},
			},
		}}

		client := NewMockCRDClient("default", responses...)

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

		if endpoints[0].Labels[endpoint.ResourceLabelKey] != "some-value" {
			t.Errorf("endpoint doesn't include the label from the registry: %#v", endpoints[0].Labels)
		}
	})
}

func TestCRDApplyChanges(t *testing.T) {
	ctx := context.Background()

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
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{
					method:    "POST",
					namespace: "default",
				}))

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
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{
					method:    "DELETE",
					namespace: "default",
					name:      "test-myid-2", // OwnerID = test; IdentifierID = myid-2
				}))

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
				executed := c.RequestWasExecuted(keyFromRequest(&mockRequest{
					method:    "PUT",
					namespace: "default",
					name:      "test-myid-3", // OwnerID = test; IdentifierID = myid-2
				}))

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
			responses = append(responses, mockResult{
				request: mockRequest{
					method:    "POST",
					namespace: "default",
				},
				response: &mockResponse{
					content: crds.DNSEntry{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								crds.RegistryResourceLabel: "some-value",
							},
						},
						Spec: crds.DNSEntrySpec{Endpoint: *testCase.Endpoint},
					},
				},
			})

			changes.Create = []*endpoint.Endpoint{testCase.Endpoint}

		case "Delete":
			responses = append(responses, mockResult{
				request: mockRequest{
					method:    "GET",
					namespace: "default",
				},
				response: &mockResponse{
					content: crds.DNSEntryList{
						Items: []crds.DNSEntry{{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "test-myid-2",
								Namespace: "default",
							},
							Spec: crds.DNSEntrySpec{
								Endpoint: *testCase.Endpoint,
							},
						}},
					},
				},
			}, mockResult{
				request: mockRequest{
					method:    "DELETE",
					name:      "test-myid-2",
					namespace: "default",
				},
				response: &mockResponse{},
			})

			changes.Delete = []*endpoint.Endpoint{testCase.Endpoint}
			seedEndpoints = append(seedEndpoints, testCase.Endpoint)
		case "Update":
			responses = append(responses, mockResult{
				request: mockRequest{
					method:    "GET",
					namespace: "default",
				},
				response: &mockResponse{
					content: crds.DNSEntryList{
						Items: []crds.DNSEntry{{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "test-myid-3",
								Namespace: "default",
							},
							Spec: crds.DNSEntrySpec{
								Endpoint: *testCase.Endpoint,
							},
						}},
					},
				},
			}, mockResult{
				request: mockRequest{
					method:    "PUT",
					name:      "test-myid-3",
					namespace: "default",
				},
				response: &mockResponse{},
			})

			changes.UpdateNew = []*endpoint.Endpoint{testCase.Endpoint}
			changes.UpdateOld = []*endpoint.Endpoint{testCase.Endpoint}
			seedEndpoints = append(seedEndpoints, testCase.Endpoint)
		default:
			t.Errorf("ChangeType not defined: %s", testCase.ChangeType)
		}

		provider := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", seedEndpoints...)
		client := NewMockCRDClient("default", responses...)
		registry := &CRDRegistry{
			provider:  provider,
			namespace: "default",
			client:    client,
			ownerID:   "test",
		}

		err := registry.ApplyChanges(ctx, &changes)
		if err != nil {
			t.Error(err)
		}

		if testCase.AssertFn != nil {
			testCase.AssertFn(client)
		}

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

func NewMockCRDClient(namespace string, responses ...mockResult) *mockClient {
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

func (mc mockClient) RequestWasExecuted(key mockRequestKey) bool {
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

func (mr *mockRequest) Body(interface{}) CRDRequest {
	return mr
}

func (mr *mockRequest) Params(runtime.Object) CRDRequest {
	return mr
}

func (mr *mockRequest) Do(ctx context.Context) CRDResult {
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

func (mr *mockErrorResponse) Into(obj runtime.Object) error {
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
