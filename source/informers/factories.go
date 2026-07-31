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
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

// Factory is the minimal behavior required from a namespaced informer factory.
type Factory interface {
	Start(stopCh <-chan struct{})
}

// SharedInformer is the minimal behavior required from the informers a Factory hands out.
type SharedInformer interface {
	Informer() cache.SharedIndexInformer
}

// NormalizeNamespaces dedups the namespaces to watch, keeping their order.
// NamespaceAll subsumes the others, since watching all and a subset duplicates every object.
func NormalizeNamespaces(namespaces []string) []string {
	seen := make(map[string]struct{}, len(namespaces))
	result := make([]string, 0, len(namespaces))
	for _, ns := range namespaces {
		if ns == v1.NamespaceAll {
			return []string{v1.NamespaceAll}
		}
		if _, ok := seen[ns]; ok {
			continue
		}
		seen[ns] = struct{}{}
		result = append(result, ns)
	}
	if len(result) == 0 {
		return []string{v1.NamespaceAll}
	}
	return result
}

// SingleNamespace returns the only namespace to watch, NamespaceAll when there is none.
// Used by sources watching a single namespace, which ValidateConfig guarantees.
func SingleNamespace(namespaces []string) string {
	return NormalizeNamespaces(namespaces)[0]
}

// Factories holds one informer factory per watched namespace, so that a tenant allowed to
// list only its own namespaces still works. Request cluster-scoped resources from a single
// factory: they ignore its namespace, so every factory would watch the same objects.
type Factories[T Factory] struct {
	namespaces []string
	factories  map[string]T
}

// NewFactories builds one factory per normalized namespace, at least one: an empty
// namespace list yields a single cluster-wide factory.
func NewFactories[T Factory](namespaces []string, newFactory func(namespace string) T) *Factories[T] {
	nss := NormalizeNamespaces(namespaces)
	factories := make(map[string]T, len(nss))
	for _, ns := range nss {
		factories[ns] = newFactory(ns)
	}
	return &Factories[T]{namespaces: nss, factories: factories}
}

// Namespaces returns the watched namespaces, in iteration order.
func (f *Factories[T]) Namespaces() []string {
	return f.namespaces
}

// WatchesAllNamespaces reports whether a single cluster-wide factory is in use.
func (f *Factories[T]) WatchesAllNamespaces() bool {
	return len(f.namespaces) == 1 && f.namespaces[0] == v1.NamespaceAll
}

// Get returns the factory watching the given namespace, if any.
func (f *Factories[T]) Get(namespace string) (T, bool) {
	factory, ok := f.factories[namespace]
	return factory, ok
}

// All returns every factory, ordered by Namespaces.
func (f *Factories[T]) All() []T {
	all := make([]T, 0, len(f.namespaces))
	for _, ns := range f.namespaces {
		all = append(all, f.factories[ns])
	}
	return all
}

// Start starts every factory. All informers must be created and configured beforehand.
func (f *Factories[T]) Start(stopCh <-chan struct{}) {
	for _, factory := range f.All() {
		factory.Start(stopCh)
	}
}

// WaitForCacheSyncAll waits for every factory cache to sync, each with its own timeout.
func WaitForCacheSyncAll[T interface {
	Factory
	informerFactory
}](ctx context.Context, factories *Factories[T]) error {
	for _, ns := range factories.namespaces {
		if err := WaitForCacheSync(ctx, factories.factories[ns]); err != nil {
			return fmt.Errorf("namespace %s: %w", namespaceName(ns), err)
		}
	}
	return nil
}

// WaitForDynamicCacheSyncAll is WaitForCacheSyncAll for dynamic informer factories.
func WaitForDynamicCacheSyncAll[T interface {
	Factory
	dynamicInformerFactory
}](ctx context.Context, factories *Factories[T]) error {
	for _, ns := range factories.namespaces {
		if err := WaitForDynamicCacheSync(ctx, factories.factories[ns]); err != nil {
			return fmt.Errorf("namespace %s: %w", namespaceName(ns), err)
		}
	}
	return nil
}

// Informers holds one informer of a given resource type per watched namespace.
type Informers[I SharedInformer] struct {
	namespaces []string
	informers  map[string]I
}

// Map builds one informer per namespace. build runs before the factories are started,
// so it may register indexers, transformers and event handlers.
func Map[T Factory, I SharedInformer](factories *Factories[T], build func(namespace string, factory T) I) *Informers[I] {
	result := make(map[string]I, len(factories.namespaces))
	for _, ns := range factories.namespaces {
		result[ns] = build(ns, factories.factories[ns])
	}
	return &Informers[I]{namespaces: factories.namespaces, informers: result}
}

// Namespaces returns the namespaces the informers are watching, in iteration order.
func (i *Informers[I]) Namespaces() []string {
	return i.namespaces
}

// For returns the informer watching the given namespace, the cluster-wide one if any.
func (i *Informers[I]) For(namespace string) (I, bool) {
	if informer, ok := i.informers[namespace]; ok {
		return informer, true
	}
	if informer, ok := i.informers[v1.NamespaceAll]; ok {
		return informer, true
	}
	var zero I
	return zero, false
}

// All returns every informer, ordered by Namespaces.
func (i *Informers[I]) All() []I {
	all := make([]I, 0, len(i.namespaces))
	for _, ns := range i.namespaces {
		all = append(all, i.informers[ns])
	}
	return all
}

// Indexers returns the indexer of every informer, to be passed to ListIndexedAll.
func (i *Informers[I]) Indexers() []cache.Indexer {
	indexers := make([]cache.Indexer, 0, len(i.namespaces))
	for _, informer := range i.All() {
		indexers = append(indexers, informer.Informer().GetIndexer())
	}
	return indexers
}

// MustAddIndexers adds the indexers to every informer.
func (i *Informers[I]) MustAddIndexers(indexers cache.Indexers) {
	for _, informer := range i.All() {
		MustAddIndexers(informer.Informer(), indexers)
	}
}

// MustSetTransform sets the transform function on every informer.
func (i *Informers[I]) MustSetTransform(fn cache.TransformFunc) {
	for _, informer := range i.All() {
		MustSetTransform(informer.Informer(), fn)
	}
}

// MustAddEventHandler adds the handler to every informer.
func (i *Informers[I]) MustAddEventHandler(handler cache.ResourceEventHandler) {
	for _, informer := range i.All() {
		MustAddEventHandler(informer.Informer(), handler)
	}
}

// namespaceName renders a namespace for logs and errors, NamespaceAll included.
func namespaceName(namespace string) string {
	if namespace == v1.NamespaceAll {
		return "(all)"
	}
	return namespace
}
