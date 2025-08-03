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
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

type Labels struct {
	values map[string]string
}

func (labels *Labels) GetKeysInOrder() []string {
	keys := make([]string, 0, len(labels.values))
	for key := range labels.values {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (labels *Labels) GetValuesOrderedByKey() []string {
	var orderedValues []string
	for _, key := range labels.GetKeysInOrder() {
		orderedValues = append(orderedValues, labels.values[key])
	}

	return orderedValues
}

type LabelOption func(*Labels)

func NewLabels(labelNames []string) *Labels {
	labels := &Labels{
		values: make(map[string]string),
	}

	for _, label := range labelNames {
		labels.values[strings.ToLower(label)] = ""
	}

	return labels
}

func (labels *Labels) WithOptions(options ...LabelOption) {
	for _, option := range options {
		option(labels)
	}
}

func WithLabel(labelName string, labelValue string) LabelOption {
	return func(labels *Labels) {
		if _, ok := labels.values[strings.ToLower(labelName)]; !ok {
			logrus.Errorf("Attempting to set a value for a label that doesn't exist! '%s' does not exist!", labelName)
		} else {
			labels.values[strings.ToLower(labelName)] = labelValue
		}
	}
}
