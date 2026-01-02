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
	"k8s.io/apimachinery/pkg/util/sets"
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
		t.Run(tt.name, func(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{}
			opt := WithEmitEvents(tt.input)
			opt(cfg)
			tt.assert(cfg)
		})
	}
}
