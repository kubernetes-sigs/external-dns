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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
)

// startableFactory records the namespace it was built for and whether it was started.
type startableFactory struct {
	namespace   string
	started     bool
	syncResults map[reflect.Type]bool
}

func (f *startableFactory) Start(_ <-chan struct{}) {
	f.started = true
}

func (f *startableFactory) WaitForCacheSync(_ <-chan struct{}) map[reflect.Type]bool {
	return f.syncResults
}

// startableDynamicFactory is the dynamic counterpart of startableFactory.
type startableDynamicFactory struct {
	namespace   string
	syncResults map[schema.GroupVersionResource]bool
}

func (f *startableDynamicFactory) Start(_ <-chan struct{}) {}

func (f *startableDynamicFactory) WaitForCacheSync(_ <-chan struct{}) map[schema.GroupVersionResource]bool {
	return f.syncResults
}

func newStartableFactories(namespaces []string) *Factories[*startableFactory] {
	return NewFactories(namespaces, func(namespace string) *startableFactory {
		return &startableFactory{
			namespace:   namespace,
			syncResults: map[reflect.Type]bool{reflect.TypeFor[string](): true},
		}
	})
}

func TestNormalizeNamespaces(t *testing.T) {
	tests := []struct {
		name       string
		namespaces []string
		want       []string
	}{
		{
			name:       "nil is cluster wide",
			namespaces: nil,
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "empty is cluster wide",
			namespaces: []string{},
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "single namespace is kept",
			namespaces: []string{"team-a"},
			want:       []string{"team-a"},
		},
		{
			name:       "order is preserved",
			namespaces: []string{"team-b", "team-a", "team-c"},
			want:       []string{"team-b", "team-a", "team-c"},
		},
		{
			name:       "duplicates are removed keeping first occurrence",
			namespaces: []string{"team-a", "team-b", "team-a"},
			want:       []string{"team-a", "team-b"},
		},
		{
			name:       "empty entry subsumes every other namespace",
			namespaces: []string{"team-a", "", "team-b"},
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "empty entry first subsumes every other namespace",
			namespaces: []string{"", "team-a"},
			want:       []string{corev1.NamespaceAll},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NormalizeNamespaces(tt.namespaces))
		})
	}
}

func TestNewFactoriesBuildsOnePerNamespace(t *testing.T) {
	factories := newStartableFactories([]string{"team-a", "team-b", "team-a"})

	assert.Equal(t, []string{"team-a", "team-b"}, factories.Namespaces())
	assert.False(t, factories.WatchesAllNamespaces())
	assert.Len(t, factories.All(), 2)

	for _, ns := range factories.Namespaces() {
		factory, ok := factories.Get(ns)
		require.True(t, ok)
		assert.Equal(t, ns, factory.namespace, "factory must be built for its own namespace")
	}

	_, ok := factories.Get("team-c")
	assert.False(t, ok, "unwatched namespace has no factory")
}

func TestNewFactoriesClusterWide(t *testing.T) {
	factories := newStartableFactories(nil)

	assert.Equal(t, []string{corev1.NamespaceAll}, factories.Namespaces())
	assert.True(t, factories.WatchesAllNamespaces())

	factory, ok := factories.Get(corev1.NamespaceAll)
	require.True(t, ok)
	assert.Equal(t, corev1.NamespaceAll, factory.namespace)
}

func TestFactoriesAllIsOrderedByNamespaces(t *testing.T) {
	factories := newStartableFactories([]string{"team-c", "team-a", "team-b"})

	got := make([]string, 0, 3)
	for _, factory := range factories.All() {
		got = append(got, factory.namespace)
	}

	assert.Equal(t, []string{"team-c", "team-a", "team-b"}, got)
}

func TestFactoriesStartStartsEveryFactory(t *testing.T) {
	factories := newStartableFactories([]string{"team-a", "team-b"})

	factories.Start(t.Context().Done())

	for _, factory := range factories.All() {
		assert.True(t, factory.started, "factory of namespace %q must be started", factory.namespace)
	}
}

func TestWaitForCacheSyncAll(t *testing.T) {
	tests := []struct {
		name        string
		namespaces  []string
		unsynced    string
		expectedErr string
	}{
		{
			name:       "all namespaces synced",
			namespaces: []string{"team-a", "team-b"},
		},
		{
			name:       "cluster wide synced",
			namespaces: nil,
		},
		{
			name:        "namespace is reported when its cache does not sync",
			namespaces:  []string{"team-a", "team-b"},
			unsynced:    "team-b",
			expectedErr: "namespace team-b: failed to sync string",
		},
		{
			name:        "cluster wide is reported when its cache does not sync",
			namespaces:  nil,
			unsynced:    corev1.NamespaceAll,
			expectedErr: "namespace (all): failed to sync string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factories := newStartableFactories(tt.namespaces)
			if tt.expectedErr != "" {
				factory, ok := factories.Get(tt.unsynced)
				require.True(t, ok)
				factory.syncResults = map[reflect.Type]bool{reflect.TypeFor[string](): false}
			}

			err := WaitForCacheSyncAll(t.Context(), factories)

			if tt.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func TestWaitForDynamicCacheSyncAll(t *testing.T) {
	newFactories := func(synced bool) *Factories[*startableDynamicFactory] {
		return NewFactories([]string{"team-a", "team-b"}, func(namespace string) *startableDynamicFactory {
			return &startableDynamicFactory{
				namespace:   namespace,
				syncResults: map[schema.GroupVersionResource]bool{{}: synced},
			}
		})
	}

	require.NoError(t, WaitForDynamicCacheSyncAll(t.Context(), newFactories(true)))

	err := WaitForDynamicCacheSyncAll(t.Context(), newFactories(false))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "namespace team-a: failed to sync")
}

func TestInformersFor(t *testing.T) {
	tests := []struct {
		name       string
		namespaces []string
		lookup     string
		wantFound  bool
		wantSource string
	}{
		{
			name:       "exact namespace match",
			namespaces: []string{"team-a", "team-b"},
			lookup:     "team-b",
			wantFound:  true,
			wantSource: "team-b",
		},
		{
			name:       "unwatched namespace is not found",
			namespaces: []string{"team-a", "team-b"},
			lookup:     "team-c",
			wantFound:  false,
		},
		{
			name:       "cluster wide informer serves every namespace",
			namespaces: nil,
			lookup:     "team-c",
			wantFound:  true,
			wantSource: corev1.NamespaceAll,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientset()
			factories := NewFactories(tt.namespaces, func(namespace string) kubeinformers.SharedInformerFactory {
				return kubeinformers.NewSharedInformerFactoryWithOptions(client, 0, kubeinformers.WithNamespace(namespace))
			})
			built := map[coreinformers.ServiceInformer]string{}
			serviceInformers := Map(factories, func(namespace string, factory kubeinformers.SharedInformerFactory) coreinformers.ServiceInformer {
				informer := factory.Core().V1().Services()
				built[informer] = namespace
				return informer
			})

			informer, ok := serviceInformers.For(tt.lookup)

			assert.Equal(t, tt.wantFound, ok)
			if !tt.wantFound {
				assert.Nil(t, informer)
				return
			}
			assert.Equal(t, tt.wantSource, built[informer])
		})
	}
}

func TestInformersListIndexedAllAcrossNamespaces(t *testing.T) {
	ctx := t.Context()
	client := fake.NewClientset()
	for _, ns := range []string{"team-a", "team-b", "team-c"} {
		svc := fakeService()
		svc.Namespace = ns
		_, err := client.CoreV1().Services(ns).Create(ctx, svc, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	factories := NewFactories([]string{"team-a", "team-b"}, func(namespace string) kubeinformers.SharedInformerFactory {
		return kubeinformers.NewSharedInformerFactoryWithOptions(client, 0, kubeinformers.WithNamespace(namespace))
	})
	serviceInformers := Map(factories, func(_ string, factory kubeinformers.SharedInformerFactory) coreinformers.ServiceInformer {
		return factory.Core().V1().Services()
	})
	serviceInformers.MustAddIndexers(IndexerWithOptions[*corev1.Service](
		IndexSelectorWithLabelSelector(labels.SelectorFromSet(labels.Set{"env": "prod"})),
	))
	serviceInformers.MustSetTransform(TransformerWithOptions[*corev1.Service](
		TransformRemoveManagedFields(),
	))

	factories.Start(ctx.Done())
	require.NoError(t, WaitForCacheSyncAll(ctx, factories))

	services := ListIndexedAll[*corev1.Service](serviceInformers.Indexers()...)

	got := make([]string, 0, len(services))
	for _, svc := range services {
		got = append(got, svc.Namespace)
		assert.Empty(t, svc.ManagedFields, "transformer must be set on every informer")
	}
	assert.ElementsMatch(t, []string{"team-a", "team-b"}, got, "services of unwatched namespaces must not be listed")
}

func TestInformersMustAddEventHandlerOnEveryInformer(t *testing.T) {
	ctx := t.Context()
	client := fake.NewClientset()

	factories := NewFactories([]string{"team-a", "team-b"}, func(namespace string) kubeinformers.SharedInformerFactory {
		return kubeinformers.NewSharedInformerFactoryWithOptions(client, 0, kubeinformers.WithNamespace(namespace))
	})
	serviceInformers := Map(factories, func(_ string, factory kubeinformers.SharedInformerFactory) coreinformers.ServiceInformer {
		return factory.Core().V1().Services()
	})
	serviceInformers.MustAddEventHandler(cache.ResourceEventHandlerFuncs{})

	factories.Start(ctx.Done())
	require.NoError(t, WaitForCacheSyncAll(ctx, factories))

	for _, informer := range serviceInformers.All() {
		assert.True(t, informer.Informer().HasSynced())
	}
}

func TestListIndexedAllWithoutIndexers(t *testing.T) {
	assert.Empty(t, ListIndexedAll[*corev1.Service]())
}
