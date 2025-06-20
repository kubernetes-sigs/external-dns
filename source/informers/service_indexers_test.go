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

package informers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

func TestServiceIndexers_SpecSelectorIndex(t *testing.T) {
	svc := &corev1.Service{}
	svc.Spec.Selector = map[string]string{"app": "nginx", "env": "prod"}

	indexFunc := ServiceIndexers[SvcSpecSelectorIndex]
	indexKeys, err := indexFunc(svc)
	expectedKey := ToSHA(labels.Set(svc.Spec.Selector).String())

	assert.NoError(t, err)
	assert.Len(t, indexKeys, 1)
	assert.Equal(t, expectedKey, indexKeys[0])
}

func TestServiceNsSelectorIndexers_NamespaceSpecSelectorIndex(t *testing.T) {
	svc := &corev1.Service{}
	svc.Namespace = "test-ns"
	svc.Spec.Selector = map[string]string{"app": "nginx", "env": "prod"}

	indexFunc := ServiceNsSelectorIndexers[SvcNamespaceSpecSelectorIndex]
	indexKeys, err := indexFunc(svc)

	expectedKey := ToSHA(svc.Namespace + "/" + labels.Set(svc.Spec.Selector).String())

	assert.NoError(t, err)
	assert.Len(t, indexKeys, 1)
	assert.Equal(t, expectedKey, indexKeys[0])
}

func TestServiceWithDefaultOptions_AddsCorrectIndexers(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		want      cache.Indexers
	}{
		{"empty namespace", "", ServiceIndexers},
		{"non-empty namespace", "ns", ServiceNsSelectorIndexers},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInf := &mockInformer{}

			si := &fakeSvcInformer{
				informer: mockInf,
			}
			err := ServiceWithDefaultOptions(si, tt.namespace)
			assert.NoError(t, err)
			assert.Equal(t, 1, mockInf.addedIndexersTimes, "AddIndexers should be called once")
			assert.Equal(t, 1, mockInf.addedHandlerTimes, "AddEventHandler should be called once")
			assert.Equal(t, tt.want, mockInf.addedIndexers)
		})
	}
}

// mocks
type fakeSvcInformer struct {
	mock.Mock
	informer *mockInformer
}

func (f *fakeSvcInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *fakeSvcInformer) Lister() corev1listers.ServiceLister {
	return nil
}

// mockInformer implements the minimal methods needed for testing.
type mockInformer struct {
	cache.SharedIndexInformer
	addedIndexers      cache.Indexers
	addedIndexersTimes int
	addedHandlerTimes  int
}

func (m *mockInformer) AddEventHandler(_ cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
	m.addedHandlerTimes++
	return nil, nil
}

func (m *mockInformer) AddIndexers(indexers cache.Indexers) error {
	m.addedIndexers = indexers
	m.addedIndexersTimes++
	return nil
}
