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

package events

import (
	"fmt"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apiv1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrlruntime "sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/external-dns/internal/sets"
)

func TestNewObjectReference_DoesNotMutateObject(t *testing.T) {
	// Verify that NewObjectReference does NOT mutate the original object
	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
		},
	}
	podCopy := pod.DeepCopy()

	_ = NewObjectReference(pod, "test")

	assert.Equal(t, podCopy, pod)
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantBase string
	}{
		{"mixed case and underscore", "My.Resource_1", "my.resource-1"},
		{"leading invalid chars", "!@#bad*chars", "a---bad-chars"},
		{"leading dash", "-start", "a-start"},
		{"trailing dash", "end-", "end-z"},
		{"leading and trailing dash", "-both-", "a-both-z"},
		{"empty input", "", "a"},
		{"all caps", "ALLCAPS", "allcaps"},
		{"dots preserved", "foo.bar", "foo.bar"},
		{"long input truncated to 253 chars", strings.Repeat("a", 300), strings.Repeat("a", 236)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitize(tt.input, time.Now())
			require.True(t, strings.HasPrefix(result, tt.wantBase+"."),
				"got %q, want prefix %q", result, tt.wantBase+".")
			require.LessOrEqual(t, len(result), 253, "name must be <= 253 chars")
		})
	}
}

func TestObjectReference_Description(t *testing.T) {
	tests := []struct {
		kind      string
		namespace string
		name      string
		expected  string
	}{
		{"Pod", "default", "nginx", "Pod/default/nginx"},
		{"Service", "prod", "api", "Service/prod/api"},
		{"Node", "", "worker-1", "Node/worker-1"},
		{"", "", "", "/"},
	}

	for _, tt := range tests {
		ref := ObjectReference{kind: tt.kind, namespace: tt.namespace, name: tt.name}
		require.Equal(t, tt.expected, ref.description())
	}
}

func TestEvent_NewEvents(t *testing.T) {
	tests := []struct {
		name    string
		event   Event
		asserts func(evs []*eventsv1.Event)
	}{
		{
			name:  "empty event",
			event: NewEvent(nil, "", ActionCreate, RecordReady),
			asserts: func(evs []*eventsv1.Event) {
				require.Empty(t, evs)
			},
		},
		{
			name: "event without uuid",
			event: NewEvent(NewObjectReference(&apiv1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Pod",
					APIVersion: "apiv1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "fake-pod",
					Namespace: apiv1.NamespaceDefault,
				},
			}, "fake"), "", ActionCreate, RecordReady),
			asserts: func(evs []*eventsv1.Event) {
				require.Len(t, evs, 1)
				e := evs[0]
				require.Contains(t, e.Name, "fake-pod.")
				require.Equal(t, apiv1.NamespaceDefault, e.Namespace)
				require.Nil(t, e.Related)
				require.Equal(t, "Pod", e.Regarding.Kind)
				require.Equal(t, "fake-pod", e.Regarding.Name)
				require.Empty(t, e.Regarding.UID)
			},
		},
		{
			name: "event with uuid",
			event: NewEvent(NewObjectReference(&apiv1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Pod",
					APIVersion: "apiv1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "fake-pod",
					Namespace: apiv1.NamespaceDefault,
					UID:       "9de3fc19-8aeb-4e76-865d-ada955403103",
				},
			}, "fake"), "", ActionCreate, RecordReady),
			asserts: func(evs []*eventsv1.Event) {
				require.Len(t, evs, 1)
				e := evs[0]
				require.Contains(t, e.Name, "fake-pod.")
				require.NotNil(t, e.Related)
				require.NotNil(t, e.Regarding)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.asserts(tt.event.events())
		})
	}
}

func TestEvent_Transpose(t *testing.T) {
	ev := NewEvent(&ObjectReference{
		kind:      "Pod",
		namespace: "default",
		name:      "nginx",
	}, "test message", ActionCreate, RecordReady)

	evs := ev.events()
	require.Len(t, evs, 1)
	event := evs[0]
	require.Contains(t, event.ObjectMeta.Name, ev.refs[0].name)
	require.Equal(t, "default", event.ObjectMeta.Namespace)
	require.Equal(t, string(ActionCreate), event.Action)
	require.Equal(t, string(RecordReady), event.Reason)
	require.Equal(t, "test message", event.Note)
	require.Equal(t, apiv1.EventTypeNormal, event.Type)
	require.Equal(t, controllerName, event.ReportingController)
	require.Contains(t, event.ReportingInstance, controllerName+"/source/")

	longMsg := strings.Repeat("a", 2000)
	ev.message = longMsg
	evs = ev.events()
	require.Len(t, evs, 1)
	require.Equal(t, longMsg[:1021]+"...", evs[0].Note)

	ev.refs[0].name = ""
	require.Empty(t, ev.events())
}

func TestEvent_NameAndEventTimeConsistent(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		now := time.Now() // deterministic bubble time: 2000-01-01 00:00:00 UTC
		expectedSuffix := fmt.Sprintf(".%x", now.UnixNano())

		ev := NewEvent(&ObjectReference{
			kind: "Pod", namespace: "default", name: "nginx",
		}, "msg", ActionCreate, RecordReady)

		evs := ev.events()
		require.Len(t, evs, 1)
		k8sEvent := evs[0]

		require.True(t, strings.HasSuffix(k8sEvent.Name, expectedSuffix),
			"name %q must end with timestamp suffix %q", k8sEvent.Name, expectedSuffix)
		require.Equal(t, now.UnixNano(), k8sEvent.EventTime.UnixNano(),
			"EventTime must match the same timestamp used in the name")
	})
}

func TestWithEmitEvents(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected sets.Set[Reason]
		assert   func(c *Config)
	}{
		{
			name:     "valid events",
			input:    []string{string(RecordReady), string(RecordError)},
			expected: sets.New(RecordReady, RecordError),
			assert: func(c *Config) {
				require.Equal(t, sets.New(RecordReady, RecordError), c.emitEvents)
				require.True(t, c.IsEnabled())
			},
		},
		{
			name:     "invalid event",
			input:    []string{"InvalidEvent"},
			expected: sets.New[Reason](),
			assert: func(c *Config) {
				require.Equal(t, sets.New[Reason](), c.emitEvents)
				require.False(t, c.IsEnabled())
			},
		},
		{
			name:     "mixed valid and invalid",
			input:    []string{string(RecordReady), "InvalidEvent"},
			expected: sets.New(RecordReady),
			assert: func(c *Config) {
				require.Equal(t, sets.New(RecordReady), c.emitEvents)
				require.True(t, c.IsEnabled())
			},
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: nil,
			assert: func(c *Config) {
				require.NotNil(t, c)
				require.False(t, c.IsEnabled())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			cfg := &Config{}
			opt := WithEmitEvents(tt.input)
			opt(cfg)
			tt.assert(cfg)
		})
	}
}

// mockEndpointInfo implements EndpointInfo for testing
type mockEndpointInfo struct {
	dnsName    string
	recordType string
	recordTTL  int64
	targets    []string
	owner      string
	refObjects []*ObjectReference
}

func (m *mockEndpointInfo) GetDNSName() string    { return m.dnsName }
func (m *mockEndpointInfo) GetRecordType() string { return m.recordType }
func (m *mockEndpointInfo) GetRecordTTL() int64   { return m.recordTTL }
func (m *mockEndpointInfo) GetTargets() []string  { return m.targets }
func (m *mockEndpointInfo) GetOwner() string      { return m.owner }
func (m *mockEndpointInfo) RefObject() *ObjectReference {
	if len(m.refObjects) == 0 {
		return nil
	}
	return m.refObjects[0]
}
func (m *mockEndpointInfo) RefObjects() []*ObjectReference { return m.refObjects }

func TestNewEventFromEndpoint(t *testing.T) {
	tests := []struct {
		name    string
		ep      EndpointInfo
		action  Action
		reason  Reason
		asserts func(t *testing.T, ev Event)
	}{
		{
			name:   "nil endpoint returns empty event",
			ep:     nil,
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, Event{}, ev)
			},
		},
		{
			name: "endpoint with no RefObjects returns empty event",
			ep: &mockEndpointInfo{
				dnsName:    "example.com",
				recordType: "A",
				recordTTL:  300,
				targets:    []string{"10.0.0.1"},
				owner:      "default",
				refObjects: nil,
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, Event{}, ev)
			},
		},
		{
			name: "valid endpoint with create action",
			ep: &mockEndpointInfo{
				dnsName:    "test.example.com",
				recordType: "A",
				recordTTL:  300,
				targets:    []string{"10.0.0.1", "10.0.0.2"},
				owner:      "my-owner",
				refObjects: []*ObjectReference{{
					kind:      "Service",
					namespace: "default",
					name:      "my-service",
					source:    "service",
				}},
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, ActionCreate, ev.action)
				require.Equal(t, RecordReady, ev.reason)
				require.Equal(t, EventTypeNormal, ev.eType)
				require.Len(t, ev.refs, 1)
				require.Equal(t, "Service", ev.refs[0].kind)
				require.Equal(t, "default", ev.refs[0].namespace)
				require.Equal(t, "my-service", ev.refs[0].name)
				require.Contains(t, ev.message, "record:test.example.com")
				require.Contains(t, ev.message, "owner:my-owner")
				require.Contains(t, ev.message, "type:A")
				require.Contains(t, ev.message, "ttl:300")
				require.Contains(t, ev.message, "targets:10.0.0.1,10.0.0.2")
				require.Contains(t, ev.message, "(external-dns)")
			},
		},
		{
			name: "valid endpoint with delete action",
			ep: &mockEndpointInfo{
				dnsName:    "deleted.example.com",
				recordType: "CNAME",
				recordTTL:  0,
				targets:    []string{"target.example.com"},
				owner:      "",
				refObjects: []*ObjectReference{{
					kind:      "Ingress",
					namespace: "prod",
					name:      "my-ingress",
					source:    "ingress",
				}},
			},
			action: ActionDelete,
			reason: RecordDeleted,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, ActionDelete, ev.action)
				require.Equal(t, RecordDeleted, ev.reason)
				require.Contains(t, ev.message, "record:deleted.example.com")
				require.Contains(t, ev.message, "type:CNAME")
				require.Contains(t, ev.message, "ttl:0")
			},
		},
		{
			name: "endpoint for cluster-scoped resource (Node)",
			ep: &mockEndpointInfo{
				dnsName:    "node1.example.com",
				recordType: "A",
				recordTTL:  60,
				targets:    []string{"192.168.1.1"},
				owner:      "default",
				refObjects: []*ObjectReference{{
					kind:      "Node",
					namespace: "", // cluster-scoped
					name:      "node1",
					source:    "node",
				}},
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, ActionCreate, ev.action)
				require.Empty(t, ev.refs[0].namespace)

				evs := ev.events()
				require.Len(t, evs, 1)
				require.Equal(t, "default", evs[0].Namespace)
			},
		},
		{
			name: "endpoint with multiple RefObjects emits one k8s event per ref",
			ep: &mockEndpointInfo{
				dnsName:    "shared.example.com",
				recordType: "A",
				recordTTL:  300,
				targets:    []string{"1.2.3.4"},
				owner:      "owner",
				refObjects: []*ObjectReference{
					{kind: "DNSEndpoint", namespace: "default", name: "shared-dns", source: "crd"},
					{kind: "Service", namespace: "default", name: "shared-svc", source: "service"},
				},
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Len(t, ev.refs, 2)
				evs := ev.events()
				require.Len(t, evs, 2)
				names := []string{evs[0].Regarding.Name, evs[1].Regarding.Name}
				require.ElementsMatch(t, []string{"shared-dns", "shared-svc"}, names)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := NewEventFromEndpoint(tt.ep, tt.action, tt.reason)
			tt.asserts(t, ev)
		})
	}
}

func TestNewObjectReference(t *testing.T) {
	tests := []struct {
		name     string
		obj      ctrlruntime.Object
		source   string
		expected *ObjectReference
	}{
		{
			name: "Pod with TypeMeta already set",
			obj: &apiv1.Pod{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Pod",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-pod",
					Namespace: "default",
					UID:       "pod-uid-123",
				},
			},
			source: "pod",
			expected: &ObjectReference{
				kind:       "Pod",
				apiVersion: "v1",
				namespace:  "default",
				name:       "my-pod",
				uid:        "pod-uid-123",
				source:     "pod",
			},
		},
		{
			name: "Pod without TypeMeta (simulating informer behavior)",
			obj: &apiv1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "informer-pod",
					Namespace: "kube-system",
					UID:       "informer-uid-456",
				},
			},
			source: "pod",
			expected: &ObjectReference{
				kind:       "Pod",
				apiVersion: "v1",
				namespace:  "kube-system",
				name:       "informer-pod",
				uid:        "informer-uid-456",
				source:     "pod",
			},
		},
		{
			name: "Service without TypeMeta",
			obj: &apiv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-service",
					Namespace: "prod",
					UID:       "svc-uid-789",
				},
			},
			source: "service",
			expected: &ObjectReference{
				kind:       "Service",
				apiVersion: "v1",
				namespace:  "prod",
				name:       "my-service",
				uid:        "svc-uid-789",
				source:     "service",
			},
		},
		{
			name: "Node (cluster-scoped, no namespace)",
			obj: &apiv1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "worker-node-1",
					UID:  "node-uid-abc",
				},
			},
			source: "node",
			expected: &ObjectReference{
				kind:       "Node",
				apiVersion: "v1",
				namespace:  "",
				name:       "worker-node-1",
				uid:        "node-uid-abc",
				source:     "node",
			},
		},
		{
			name: "Endpoints without TypeMeta",
			obj: &apiv1.Endpoints{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-endpoints",
					Namespace: "default",
					UID:       "ep-uid-def",
				},
			},
			source: "endpoints",
			expected: &ObjectReference{
				kind:       "Endpoints",
				apiVersion: "v1",
				namespace:  "default",
				name:       "my-endpoints",
				uid:        "ep-uid-def",
				source:     "endpoints",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewObjectReference(tt.obj, tt.source)
			require.Equal(t, tt.expected.kind, result.kind)
			require.Equal(t, tt.expected.apiVersion, result.apiVersion)
			require.Equal(t, tt.expected.namespace, result.namespace)
			require.Equal(t, tt.expected.name, result.name)
			require.Equal(t, tt.expected.uid, result.uid)
			require.Equal(t, tt.expected.source, result.source)
		})
	}
}

// customObject is a type not registered in the scheme, used to test reflection fallback
type customObject struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

func (c *customObject) DeepCopyObject() runtime.Object {
	return &customObject{
		TypeMeta:   c.TypeMeta,
		ObjectMeta: *c.ObjectMeta.DeepCopy(),
	}
}

func TestEvent_Accessors(t *testing.T) {
	ref := &ObjectReference{kind: "Pod", namespace: "default", name: "nginx"}
	ev := NewEvent(ref, "msg", ActionDelete, RecordDeleted)

	assert.Equal(t, ActionDelete, ev.Action())
	assert.Equal(t, RecordDeleted, ev.Reason())
	assert.Equal(t, EventTypeNormal, ev.EventType())
}

func TestObjectReference_Key(t *testing.T) {
	tests := []struct {
		name     string
		ref      *ObjectReference
		expected string
	}{
		{
			name:     "namespaced resource",
			ref:      &ObjectReference{source: "service", namespace: "default", name: "svc"},
			expected: "service/default/svc",
		},
		{
			name:     "cluster-scoped resource has empty namespace segment",
			ref:      &ObjectReference{source: "node", namespace: "", name: "node-1"},
			expected: "node//node-1",
		},
		{
			name:     "different sources with same namespace and name produce distinct keys",
			ref:      &ObjectReference{source: "crd", namespace: "default", name: "my-dns"},
			expected: "crd/default/my-dns",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.ref.Key())
		})
	}
}

func TestWithDryRun(t *testing.T) {
	cfg := NewConfig(WithDryRun(true))
	assert.True(t, cfg.dryRun)

	cfg = NewConfig(WithDryRun(false))
	assert.False(t, cfg.dryRun)
}

func TestNewObjectReference_ReflectionFallback(t *testing.T) {
	// Test that when object type is not in scheme, reflection is used to get Kind
	obj := &customObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "custom-resource",
			Namespace: "custom-ns",
			UID:       "custom-uid-123",
		},
	}

	ref := NewObjectReference(obj, "custom")

	// Kind should be derived from reflection (struct name)
	require.Equal(t, "customObject", ref.kind)
	// APIVersion will be empty since it's not in scheme
	require.Empty(t, ref.apiVersion)
	require.Equal(t, "custom-ns", ref.namespace)
	require.Equal(t, "custom-resource", ref.name)
	require.Equal(t, "custom-uid-123", string(ref.uid))
	require.Equal(t, "custom", ref.source)
}
