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
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/source/annotations"
)

func TestIndexerWithOptions_FilterByAnnotation(t *testing.T) {
	indexers := IndexerWithOptions[*unstructured.Unstructured](
		IndexSelectorWithAnnotationFilter("example-annotation"),
	)

	obj := &unstructured.Unstructured{}
	obj.SetAnnotations(map[string]string{"example-annotation": "value"})
	obj.SetNamespace("default")
	obj.SetName("test-object")

	keys, err := indexers[IndexWithSelectors](obj)
	assert.NoError(t, err)
	assert.Equal(t, []string{"default/test-object"}, keys)
}

func TestIndexerWithOptions_FilterByLabel(t *testing.T) {
	labelSelector := labels.SelectorFromSet(labels.Set{"app": "nginx"})
	indexers := IndexerWithOptions[*corev1.Pod](
		IndexSelectorWithLabelSelector(labelSelector),
	)

	obj := &corev1.Pod{}
	obj.SetLabels(map[string]string{"app": "nginx"})
	obj.SetNamespace("default")
	obj.SetName("test-object")

	keys, err := indexers[IndexWithSelectors](obj)
	assert.NoError(t, err)
	assert.Equal(t, []string{"default/test-object"}, keys)
}

func TestIndexerWithOptions_NoMatch(t *testing.T) {
	labelSelector := labels.SelectorFromSet(labels.Set{"app": "nginx"})
	indexers := IndexerWithOptions[*unstructured.Unstructured](
		IndexSelectorWithLabelSelector(labelSelector),
	)

	obj := &unstructured.Unstructured{}
	obj.SetLabels(map[string]string{"app": "apache"})
	obj.SetNamespace("default")
	obj.SetName("test-object")

	keys, err := indexers[IndexWithSelectors](obj)
	assert.NoError(t, err)
	assert.Nil(t, keys)
}

func TestIndexerWithOptions_InvalidType(t *testing.T) {
	indexers := IndexerWithOptions[*unstructured.Unstructured]()

	obj := "invalid-object"

	keys, err := indexers[IndexWithSelectors](obj)
	assert.Error(t, err)
	assert.Nil(t, keys)
	assert.Contains(t, err.Error(), "object is not of type")
}

func TestIndexerWithOptions_EmptyOptions(t *testing.T) {
	indexers := IndexerWithOptions[*unstructured.Unstructured]()

	obj := &unstructured.Unstructured{}
	obj.SetNamespace("default")
	obj.SetName("test-object")

	keys, err := indexers["withSelectors"](obj)
	assert.NoError(t, err)
	assert.Equal(t, []string{"default/test-object"}, keys)
}

func TestIndexerWithOptions_ClusterScopedResource(t *testing.T) {
	indexers := IndexerWithOptions[*corev1.Node]()
	indexFn := indexers[IndexWithSelectors]

	node := &corev1.Node{}
	node.SetName("my-node")

	keys, err := indexFn(node)
	assert.NoError(t, err)
	assert.Equal(t, []string{"my-node"}, keys)
}

func TestIndexerWithOptions_AnnotationFilterNoMatch(t *testing.T) {
	indexers := IndexerWithOptions[*unstructured.Unstructured](
		IndexSelectorWithAnnotationFilter("example-annotation=value"),
	)

	obj := &unstructured.Unstructured{}
	obj.SetAnnotations(map[string]string{"other-annotation": "value"})
	obj.SetNamespace("default")
	obj.SetName("test-object")

	keys, err := indexers[IndexWithSelectors](obj)
	assert.NoError(t, err)
	assert.Nil(t, keys)
}

func TestIndexSelectorWithAnnotationFilter(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedFilter labels.Selector
	}{
		{
			name:           "valid input",
			input:          "key=value",
			expectedFilter: func() labels.Selector { s, _ := annotations.ParseFilter("key=value"); return s }(),
		},
		{
			name:           "empty input",
			input:          "",
			expectedFilter: nil,
		},
		{
			name:           "key only filter",
			input:          "app",
			expectedFilter: func() labels.Selector { s, _ := annotations.ParseFilter("app"); return s }(),
		},
		{
			name:           "poisoned input",
			input:          "=app",
			expectedFilter: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &IndexSelectorOptions{}
			IndexSelectorWithAnnotationFilter(tt.input)(options)
			assert.Equal(t, tt.expectedFilter, options.annotationFilter)
		})
	}
}

func TestIndexerWithOptions_LabelKey(t *testing.T) {
	indexers := IndexerWithOptions[*discoveryv1.EndpointSlice](
		IndexSelectorWithLabelKey(discoveryv1.LabelServiceName),
	)
	indexFn := indexers[IndexWithSelectors]

	t.Run("returns namespace/serviceName when label is set", func(t *testing.T) {
		es := &discoveryv1.EndpointSlice{}
		es.SetNamespace("default")
		es.SetName("my-slice")
		es.SetLabels(map[string]string{discoveryv1.LabelServiceName: "my-service"})

		keys, err := indexFn(es)
		assert.NoError(t, err)
		assert.Equal(t, []string{"default/my-service"}, keys)
	})

	t.Run("returns nil when label is absent", func(t *testing.T) {
		es := &discoveryv1.EndpointSlice{}
		es.SetNamespace("default")
		es.SetName("my-slice")
		es.SetLabels(map[string]string{})

		keys, err := indexFn(es)
		assert.NoError(t, err)
		assert.Nil(t, keys)
	})

	t.Run("wrong type returns error", func(t *testing.T) {
		keys, err := indexFn(&corev1.Service{})
		assert.Error(t, err)
		assert.Nil(t, keys)
	})
}

func TestGetByKey_ObjectExists(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	pod := &corev1.Pod{}
	pod.SetNamespace("default")
	pod.SetName("test-pod")

	err := indexer.Add(pod)
	assert.NoError(t, err)

	result, err := GetByKey[*corev1.Pod](indexer, "default/test-pod")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-pod", result.GetName())
}

func TestGetByKey_ObjectDoesNotExist(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})

	result, err := GetByKey[*corev1.Pod](indexer, "default/non-existent-pod")
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByKey_TypeAssertionFailure(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	service := &corev1.Service{}
	service.SetNamespace("default")
	service.SetName("test-service")

	err := indexer.Add(service)
	assert.NoError(t, err)

	result, err := GetByKey[*corev1.Pod](indexer, "default/test-service")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "object is not of type")
	assert.Nil(t, result)
}

func TestListIndexed_Empty(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, IndexerWithOptions[*corev1.Pod]())
	assert.Empty(t, ListIndexed[*corev1.Pod](indexer))
}

func TestListIndexed_AllMatching(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, IndexerWithOptions[*corev1.Pod]())
	for _, name := range []string{"pod-1", "pod-2", "pod-3"} {
		p := &corev1.Pod{}
		p.SetNamespace("default")
		p.SetName(name)
		require.NoError(t, indexer.Add(p))
	}
	assert.Len(t, ListIndexed[*corev1.Pod](indexer), 3)
}

func TestListIndexed_LabelFilterExcludes(t *testing.T) {
	sel := labels.SelectorFromSet(labels.Set{"app": "nginx"})
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, IndexerWithOptions[*corev1.Pod](
		IndexSelectorWithLabelSelector(sel),
	))

	match := &corev1.Pod{}
	match.SetNamespace("default")
	match.SetName("nginx-pod")
	match.SetLabels(map[string]string{"app": "nginx"})

	noMatch := &corev1.Pod{}
	noMatch.SetNamespace("default")
	noMatch.SetName("other-pod")
	noMatch.SetLabels(map[string]string{"app": "apache"})

	require.NoError(t, indexer.Add(match))
	require.NoError(t, indexer.Add(noMatch))

	result := ListIndexed[*corev1.Pod](indexer)
	assert.Len(t, result, 1)
	assert.Equal(t, "nginx-pod", result[0].GetName())
}

func TestListIndexed_ClusterScopedResource(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, IndexerWithOptions[*corev1.Node]())
	node := &corev1.Node{}
	node.SetName("my-node")
	require.NoError(t, indexer.Add(node))

	result := ListIndexed[*corev1.Node](indexer)
	assert.Len(t, result, 1)
	assert.Equal(t, "my-node", result[0].GetName())
}

func TestIndexSelectorWithFunctions(t *testing.T) {
	makePod := func(name, namespace string, labelsMap, annotationsMap map[string]string) *corev1.Pod {
		p := &corev1.Pod{}
		p.SetName(name)
		p.SetNamespace(namespace)
		p.SetLabels(labelsMap)
		p.SetAnnotations(annotationsMap)
		return p
	}

	type testCase struct {
		name     string
		indexers cache.Indexers
		obj      any
		wantKeys []string
		wantErr  bool
	}

	tests := []testCase{
		{
			name: "single predicate match includes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(func(p *corev1.Pod) bool { return p.GetName() == "wanted" }),
			),
			obj:      makePod("wanted", "default", nil, nil),
			wantKeys: []string{"default/wanted"},
		},
		{
			name: "single predicate mismatch excludes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(func(p *corev1.Pod) bool { return p.GetName() == "wanted" }),
			),
			obj: makePod("other", "default", nil, nil),
		},
		{
			name: "multiple predicates all pass includes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(
					func(p *corev1.Pod) bool { return p.GetNamespace() == "prod" },
					func(p *corev1.Pod) bool { return p.GetLabels()["app"] == "web" },
				),
			),
			obj:      makePod("p", "prod", map[string]string{"app": "web"}, nil),
			wantKeys: []string{"prod/p"},
		},
		{
			name: "multiple predicates one fails excludes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(
					func(p *corev1.Pod) bool { return p.GetNamespace() == "prod" },
					func(p *corev1.Pod) bool { return p.GetLabels()["app"] == "web" },
				),
			),
			obj: makePod("p", "prod", map[string]string{"app": "api"}, nil),
		},
		{
			name: "wrong concrete predicate type excludes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(func(_ *corev1.Service) bool { return true }),
			),
			obj: makePod("p", "default", nil, nil),
		},
		{
			name: "outer type mismatch returns error",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithConditions(func(_ *corev1.Pod) bool { return true }),
			),
			obj:     &corev1.Service{},
			wantErr: true,
		},
		{
			name: "combined annotation and predicate both pass includes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithAnnotationFilter("env=prod"),
				IndexSelectorWithConditions(func(p *corev1.Pod) bool { return p.GetLabels()["app"] == "web" }),
			),
			obj:      makePod("p", "default", map[string]string{"app": "web"}, map[string]string{"env": "prod"}),
			wantKeys: []string{"default/p"},
		},
		{
			name: "combined annotation fails excludes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithAnnotationFilter("env=prod"),
				IndexSelectorWithConditions(func(p *corev1.Pod) bool { return p.GetLabels()["app"] == "web" }),
			),
			obj: makePod("p", "default", map[string]string{"app": "web"}, map[string]string{"env": "staging"}),
		},
		{
			name: "combined predicate fails excludes object",
			indexers: IndexerWithOptions[*corev1.Pod](
				IndexSelectorWithAnnotationFilter("env=prod"),
				IndexSelectorWithConditions(func(p *corev1.Pod) bool { return p.GetLabels()["app"] == "web" }),
			),
			obj: makePod("p", "default", map[string]string{"app": "api"}, map[string]string{"env": "prod"}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, err := tt.indexers[IndexWithSelectors](tt.obj)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, keys)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKeys, keys)
		})
	}
}
