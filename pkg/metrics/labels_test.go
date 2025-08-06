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

package metrics

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestNewLabels(t *testing.T) {
	tests := []struct {
		name               string
		labelNames         []string
		expectedLabelNames []string
	}{
		{
			name:               "NewLabels initializes Values with labelNames",
			labelNames:         []string{"label1", "label2"},
			expectedLabelNames: []string{"label1", "label2"},
		},
		{
			name:               "NewLabels sets labelNames as lower-case",
			labelNames:         []string{"LABEL1", "label2"},
			expectedLabelNames: []string{"label1", "label2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			labels := NewLabels(tt.labelNames)
			keys := labels.GetKeysInOrder()

			assert.Equal(t, tt.expectedLabelNames, keys)
		})
	}
}

func TestLabelsWithOptions(t *testing.T) {
	tests := []struct {
		name              string
		labelNames        []string
		options           []LabelOption
		expectedValuesMap map[string]string
	}{
		{
			name:       "WithOptions sets label values",
			labelNames: []string{"label1", "label2"},
			options: []LabelOption{
				WithLabel("label1", "alpha"),
				WithLabel("label2", "beta"),
			},
			expectedValuesMap: map[string]string{
				"label1": "alpha",
				"label2": "beta",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			labels := NewLabels(tt.labelNames)
			labels.WithOptions(tt.options...)

			assert.Equal(t, tt.expectedValuesMap, labels.values)
		})
	}
}

func TestLabelsWithLabel(t *testing.T) {
	tests := []struct {
		name             string
		labelNames       []string
		labelName        string
		labelValue       string
		expectedLabels   *Labels
		expectedErrorLog string
	}{
		{
			name:       "WithLabel sets label and value",
			labelNames: []string{"label1"},
			labelName:  "label1",
			labelValue: "alpha",
			expectedLabels: &Labels{
				values: map[string]string{
					"label1": "alpha",
				}},
		},
		{
			name:       "WithLabel sets labelName to lowercase",
			labelNames: []string{"label1"},
			labelName:  "LABEL1",
			labelValue: "alpha",
			expectedLabels: &Labels{
				values: map[string]string{
					"label1": "alpha",
				}},
		},
		{
			name:             "WithLabel errors if label doesn't exist",
			labelNames:       []string{"label1"},
			labelName:        "notreal",
			labelValue:       "",
			expectedErrorLog: "Attempting to set a value for a label that doesn't exist! 'notreal' does not exist!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := testutils.LogsUnderTestWithLogLevel(logrus.WarnLevel, t)

			labels := NewLabels(tt.labelNames)
			labels.WithOptions(WithLabel(tt.labelName, tt.labelValue))

			if tt.expectedLabels != nil {
				assert.Equal(t, tt.expectedLabels, labels)
			}

			if tt.expectedErrorLog != "" {
				testutils.TestHelperLogContains(tt.expectedErrorLog, hook, t)
			}
		})
	}
}

func TestLabelsGetKeysInOrder(t *testing.T) {
	tests := []struct {
		name                string
		labels              *Labels
		expectedKeysInOrder []string
	}{
		{
			"GetKeysInOrder returns keys in alphabetical order",
			NewLabels([]string{"label2", "label1"}),
			[]string{"label1", "label2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderedKeys := tt.labels.GetKeysInOrder()

			assert.Equal(t, tt.expectedKeysInOrder, orderedKeys)
		})
	}
}

func TestLabelsGetValuesOrderedByKey(t *testing.T) {
	tests := []struct {
		name                  string
		labels                *Labels
		labelOptions          []LabelOption
		expectedValuesInOrder []string
	}{
		{
			"GetKeysInOrder returns keys in alphabetical order",
			NewLabels([]string{"label1", "label2"}),
			[]LabelOption{
				WithLabel("label2", "beta"),
				WithLabel("label1", "alpha"),
			},
			[]string{"alpha", "beta"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.labels.WithOptions(tt.labelOptions...)

			orderedValues := tt.labels.GetValuesOrderedByKey()

			assert.Equal(t, tt.expectedValuesInOrder, orderedValues)
		})
	}
}
