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

package annotations

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
)

// AnnotatedObject represents any Kubernetes object with annotations
type AnnotatedObject interface {
	GetAnnotations() map[string]string
}

// Filter filters a slice of objects by annotation selector.
// Returns all items if annotationFilter is empty.
func Filter[T AnnotatedObject](items []T, filter string) ([]T, error) {
	if filter == "" || strings.TrimSpace(filter) == "" {
		return items, nil
	}
	selector, err := ParseFilter(filter)
	if err != nil {
		return nil, err
	}
	if selector.Empty() {
		return items, nil
	}

	filtered := make([]T, 0, len(items))
	for _, item := range items {
		if selector.Matches(labels.Set(item.GetAnnotations())) {
			filtered = append(filtered, item)
		}
	}
	log.Debugf("filtered '%d' services out of '%d' with annotation filter '%s'", len(filtered), len(items), filter)
	return filtered, nil
}
