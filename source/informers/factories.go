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

package informers

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
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

// Factories holds one informer factory per watched namespace, so RBAC can grant
// namespace-scoped access only (least privilege).
// Cluster-scoped resources need a separate factory.
type Factories[T Factory] struct {
	namespaces []string
	factories  map[string]T
}

// NewFactories builds one factory per namespace, as normalized by NewSourceConfig.
// An empty namespace list yields a single cluster-wide factory.
// Duplicates are dropped so that no namespace is watched twice.
func NewFactories[T Factory](namespaces []string, newFactory func(namespace string) T) *Factories[T] {
	if len(namespaces) == 0 {
		namespaces = []string{v1.NamespaceAll}
	}
	nss := make([]string, 0, len(namespaces))
	factories := make(map[string]T, len(namespaces))
	for _, ns := range namespaces {
		if _, ok := factories[ns]; ok {
			continue
		}
		nss = append(nss, ns)
		factories[ns] = newFactory(ns)
	}
	return &Factories[T]{namespaces: nss, factories: factories}
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

// WaitForCacheSyncAll waits for every factory cache to sync. Namespaces sync in parallel,
// so startup stays within one timeout however many are watched.
func WaitForCacheSyncAll[T interface {
	Factory
	informerFactory
}](ctx context.Context, factories *Factories[T]) error {
	return waitForCacheSyncAll(ctx, factories, func(ctx context.Context, factory T) error {
		return WaitForCacheSync(ctx, factory)
	})
}

// WaitForDynamicCacheSyncAll is WaitForCacheSyncAll for dynamic informer factories.
func WaitForDynamicCacheSyncAll[T interface {
	Factory
	dynamicInformerFactory
}](ctx context.Context, factories *Factories[T]) error {
	return waitForCacheSyncAll(ctx, factories, func(ctx context.Context, factory T) error {
		return WaitForDynamicCacheSync(ctx, factory)
	})
}

func waitForCacheSyncAll[T Factory](ctx context.Context, factories *Factories[T], wait func(context.Context, T) error) error {
	g, gctx := errgroup.WithContext(ctx)
	for _, ns := range factories.namespaces {
		factory := factories.factories[ns]
		g.Go(func() error {
			if err := wait(gctx, factory); err != nil {
				return fmt.Errorf("namespace %s: %w", namespaceName(ns), err)
			}
			return nil
		})
	}
	return g.Wait()
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

// Single wraps one informer as a collection watching every namespace, so that callers
// holding a plain informer can use the same API.
func Single[I SharedInformer](informer I) *Informers[I] {
	return &Informers[I]{
		namespaces: []string{v1.NamespaceAll},
		informers:  map[string]I{v1.NamespaceAll: informer},
	}
}

// For returns the informer watching the given namespace, the cluster-wide one if any.
// A nil collection finds nothing, so a source that skipped building it reports a
// missing informer rather than panicking.
func (i *Informers[I]) For(namespace string) (I, bool) {
	var zero I
	if i == nil {
		return zero, false
	}
	if informer, ok := i.informers[namespace]; ok {
		return informer, true
	}
	if informer, ok := i.informers[v1.NamespaceAll]; ok {
		return informer, true
	}
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
