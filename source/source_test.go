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

package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"
)

func TestGetLabelSelector(t *testing.T) {
	tests := []struct {
		name             string
		annotationFilter string
		expectError      bool
		expectedSelector string
	}{
		{
			name:             "Valid label selector",
			annotationFilter: "key1=value1,key2=value2",
			expectedSelector: "key1=value1,key2=value2",
		},
		{
			name:             "Invalid label selector",
			annotationFilter: "key1==value1",
			expectedSelector: "key1=value1",
		},
		{
			name:             "Empty label selector",
			annotationFilter: "",
			expectedSelector: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector, err := getLabelSelector(tt.annotationFilter)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedSelector, selector.String())
		})
	}
}

func TestMatchLabelSelector(t *testing.T) {
	tests := []struct {
		name           string
		selector       labels.Selector
		srcAnnotations map[string]string
		expectedMatch  bool
	}{
		{
			name:           "Matching label selector",
			selector:       labels.SelectorFromSet(labels.Set{"key1": "value1"}),
			srcAnnotations: map[string]string{"key1": "value1", "key2": "value2"},
			expectedMatch:  true,
		},
		{
			name:           "Non-matching label selector",
			selector:       labels.SelectorFromSet(labels.Set{"key1": "value1"}),
			srcAnnotations: map[string]string{"key2": "value2"},
			expectedMatch:  false,
		},
		{
			name:           "Empty label selector",
			selector:       labels.NewSelector(),
			srcAnnotations: map[string]string{"key1": "value1"},
			expectedMatch:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchLabelSelector(tt.selector, tt.srcAnnotations)
			assert.Equal(t, tt.expectedMatch, result)
		})
	}
}
