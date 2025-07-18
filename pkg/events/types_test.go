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
		require.Equal(t, tt.expected, ev.Reference())
	}
}

func TestEvent_Transpose(t *testing.T) {
	ev := Event{
		ref: ObjectReference{
			Kind:      "Pod",
			Namespace: "default",
			Name:      "nginx",
		},
		message: "test message",
		action:  ActionCreate,
		eType:   "Normal",
		reason:  RecordReady,
	}

	event := ev.transpose()
	require.NotNil(t, event)
	require.Contains(t, event.ObjectMeta.Name, ev.ref.Name)
	require.Equal(t, "default", event.ObjectMeta.Namespace)
	require.Equal(t, string(ActionCreate), event.Action)
	require.Equal(t, string(RecordReady), event.Reason)
	require.Equal(t, "test message", event.Note)
	require.Equal(t, "Normal", event.Type)
	require.Contains(t, event.ReportingInstance, "")
	require.Equal(t, controllerName, event.ReportingController)

	longMsg := strings.Repeat("a", 2000)
	ev.message = longMsg
	event = ev.transpose()
	require.Equal(t, longMsg[:1021]+"...", event.Note)

	ev.ref.Name = ""
	require.Nil(t, ev.transpose())
}

func TestWithEmitEvents(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected sets.Set[string]
	}{
		{
			name:     "valid events",
			input:    []string{string(RecordReady), string(RecordError)},
			expected: sets.New[string](string(RecordReady), string(RecordError)),
		},
		{
			name:     "invalid event",
			input:    []string{"InvalidEvent"},
			expected: sets.New[string](),
		},
		{
			name:     "mixed valid and invalid",
			input:    []string{string(RecordReady), "InvalidEvent"},
			expected: sets.New[string](string(RecordReady)),
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{}
			opt := WithEmitEvents(tt.input)
			opt(cfg)
			require.Equal(t, tt.expected, cfg.emitEvents)
		})
	}
}
