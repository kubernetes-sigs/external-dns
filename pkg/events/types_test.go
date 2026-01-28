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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	apiv1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	ctrlruntime "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		input    string
		expected string // expected prefix of sanitized output
	}{
		{"My.Resource_1", "my.resource-1."},
		{"!@#bad*chars", "a---bad-chars."},
		{"-start", "a-start."},
		{"end-", "end-z."},
		{"-both-", "a-both-z."},
		{"", "a."},
		{"ALLCAPS", "allcaps."},
		{"foo.bar", "foo.bar."},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := sanitize(tt.input)
			require.True(t, strings.HasPrefix(result, tt.expected), "expected prefix %q, got %q", tt.expected, result)
			require.Greater(t, len(result), len(tt.expected))
		})
	}
}

func TestEvent_Reference(t *testing.T) {
	tests := []struct {
		kind      string
		namespace string
		name      string
		expected  string
	}{
		{"Pod", "default", "nginx", "Pod/default/nginx"},
		{"Service", "prod", "api", "Service/prod/api"},
		{"", "", "", "//"},
	}

	for _, tt := range tests {
		ev := Event{
			ref: ObjectReference{
				Kind:      tt.kind,
				Namespace: tt.namespace,
				Name:      tt.name,
				Source:    "fake-source",
			},
		}
		require.Equal(t, tt.expected, ev.description())
	}
}

func TestEvent_NewEvents(t *testing.T) {
	tests := []struct {
		name    string
		event   Event
		asserts func(e *eventsv1.Event)
	}{
		{
			name:  "empty event",
			event: NewEvent(nil, "", ActionCreate, RecordReady),
			asserts: func(e *eventsv1.Event) {
				require.Nil(t, e)
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
			asserts: func(e *eventsv1.Event) {
				require.NotNil(t, e)
				require.Contains(t, e.Name, "fake-pod.")
				require.Equal(t, apiv1.NamespaceDefault, e.Namespace)
				require.Nil(t, e.Related)
				require.Equal(t, apiv1.ObjectReference{}, e.Regarding)
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
			asserts: func(e *eventsv1.Event) {
				require.NotNil(t, e)
				require.Contains(t, e.Name, "fake-pod.")
				require.NotNil(t, e.Related)
				require.NotNil(t, e.Regarding)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.asserts(tt.event.event())
		})
	}
}

func TestEvent_Transpose(t *testing.T) {
	ev := NewEvent(&ObjectReference{
		Kind:      "Pod",
		Namespace: "default",
		Name:      "nginx",
	}, "test message", ActionCreate, RecordReady)

	event := ev.event()
	require.NotNil(t, event)
	require.Contains(t, event.ObjectMeta.Name, ev.ref.Name)
	require.Equal(t, "default", event.ObjectMeta.Namespace)
	require.Equal(t, string(ActionCreate), event.Action)
	require.Equal(t, string(RecordReady), event.Reason)
	require.Equal(t, "test message", event.Note)
	require.Equal(t, apiv1.EventTypeNormal, event.Type)
	require.Equal(t, controllerName, event.ReportingController)
	require.Contains(t, event.ReportingInstance, controllerName+"/source/")

	longMsg := strings.Repeat("a", 2000)
	ev.message = longMsg
	event = ev.event()
	require.Equal(t, longMsg[:1021]+"...", event.Note)

	ev.ref.Name = ""
	require.Nil(t, ev.event())
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
			expected: sets.New[Reason](RecordReady, RecordError),
			assert: func(c *Config) {
				require.Equal(t, sets.New[Reason](RecordReady, RecordError), c.emitEvents)
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
			expected: sets.New[Reason](RecordReady),
			assert: func(c *Config) {
				require.Equal(t, sets.New[Reason](RecordReady), c.emitEvents)
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
	refObject  *ObjectReference
}

func (m *mockEndpointInfo) GetDNSName() string          { return m.dnsName }
func (m *mockEndpointInfo) GetRecordType() string       { return m.recordType }
func (m *mockEndpointInfo) GetRecordTTL() int64         { return m.recordTTL }
func (m *mockEndpointInfo) GetTargets() []string        { return m.targets }
func (m *mockEndpointInfo) GetOwner() string            { return m.owner }
func (m *mockEndpointInfo) RefObject() *ObjectReference { return m.refObject }

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
			name: "endpoint with nil RefObject returns empty event",
			ep: &mockEndpointInfo{
				dnsName:    "example.com",
				recordType: "A",
				recordTTL:  300,
				targets:    []string{"10.0.0.1"},
				owner:      "default",
				refObject:  nil,
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
				refObject: &ObjectReference{
					Kind:      "Service",
					Namespace: "default",
					Name:      "my-service",
					Source:    "service",
				},
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, ActionCreate, ev.action)
				require.Equal(t, RecordReady, ev.reason)
				require.Equal(t, EventTypeNormal, ev.eType)
				require.Equal(t, "Service", ev.ref.Kind)
				require.Equal(t, "default", ev.ref.Namespace)
				require.Equal(t, "my-service", ev.ref.Name)
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
				refObject: &ObjectReference{
					Kind:      "Ingress",
					Namespace: "prod",
					Name:      "my-ingress",
					Source:    "ingress",
				},
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
			name: "endpoint for cluster-scoped resource (Node) should handle empty namespace",
			ep: &mockEndpointInfo{
				dnsName:    "node1.example.com",
				recordType: "A",
				recordTTL:  60,
				targets:    []string{"192.168.1.1"},
				owner:      "default",
				refObject: &ObjectReference{
					Kind:      "Node",
					Namespace: "", // cluster-scoped
					Name:      "node1",
					Source:    "node",
				},
			},
			action: ActionCreate,
			reason: RecordReady,
			asserts: func(t *testing.T, ev Event) {
				require.Equal(t, ActionCreate, ev.action)
				require.Empty(t, ev.ref.Namespace)
				k8sEvent := ev.event()
				require.NotNil(t, k8sEvent)
				require.Equal(t, "default", k8sEvent.Namespace)
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
				Kind:       "Pod",
				ApiVersion: "v1",
				Namespace:  "default",
				Name:       "my-pod",
				UID:        "pod-uid-123",
				Source:     "pod",
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
				Kind:       "Pod",
				ApiVersion: "v1",
				Namespace:  "kube-system",
				Name:       "informer-pod",
				UID:        "informer-uid-456",
				Source:     "pod",
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
				Kind:       "Service",
				ApiVersion: "v1",
				Namespace:  "prod",
				Name:       "my-service",
				UID:        "svc-uid-789",
				Source:     "service",
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
				Kind:       "Node",
				ApiVersion: "v1",
				Namespace:  "",
				Name:       "worker-node-1",
				UID:        "node-uid-abc",
				Source:     "node",
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
				Kind:       "Endpoints",
				ApiVersion: "v1",
				Namespace:  "default",
				Name:       "my-endpoints",
				UID:        "ep-uid-def",
				Source:     "endpoints",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewObjectReference(tt.obj, tt.source)
			require.Equal(t, tt.expected.Kind, result.Kind)
			require.Equal(t, tt.expected.ApiVersion, result.ApiVersion)
			require.Equal(t, tt.expected.Namespace, result.Namespace)
			require.Equal(t, tt.expected.Name, result.Name)
			require.Equal(t, tt.expected.UID, result.UID)
			require.Equal(t, tt.expected.Source, result.Source)
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
	require.Equal(t, "customObject", ref.Kind)
	// APIVersion will be empty since it's not in scheme
	require.Empty(t, ref.ApiVersion)
	require.Equal(t, "custom-ns", ref.Namespace)
	require.Equal(t, "custom-resource", ref.Name)
	require.Equal(t, "custom-uid-123", string(ref.UID))
	require.Equal(t, "custom", ref.Source)
}
