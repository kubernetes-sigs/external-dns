/*
Copyright 2026 The Kubernetes Authors.

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
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	corelisters "k8s.io/client-go/listers/core/v1"
)

type NamespaceFilter struct {
	namespaces           map[string]struct{}
	namespaceLabelFilter labels.Selector
	namespaceLister      corelisters.NamespaceLister
}

func NewNamespaceFilter(namespaces []string, namespaceLabelFilter labels.Selector, namespaceLister corelisters.NamespaceLister) *NamespaceFilter {
	nsMap := make(map[string]struct{})
	for _, ns := range namespaces {
		nsMap[ns] = struct{}{}
	}
	return &NamespaceFilter{
		namespaces:           nsMap,
		namespaceLabelFilter: namespaceLabelFilter,
		namespaceLister:      namespaceLister,
	}
}

func (f *NamespaceFilter) Matches(ns string) bool {
	if ns == corev1.NamespaceAll {
		return true
	}
	if len(f.namespaces) == 0 && (f.namespaceLabelFilter == nil || f.namespaceLabelFilter.Empty()) {
		return true
	}
	if len(f.namespaces) > 0 {
		if _, ok := f.namespaces[ns]; !ok {
			return false
		}
	}
	if f.namespaceLabelFilter != nil && !f.namespaceLabelFilter.Empty() {
		if f.namespaceLister == nil {
			return false
		}
		namespace, err := f.namespaceLister.Get(ns)
		if err != nil {
			log.Errorf("failed to get namespace %s to check labels: %v", ns, err)
			return false
		}
		if !f.namespaceLabelFilter.Matches(labels.Set(namespace.Labels)) {
			return false
		}
	}
	return true
}
