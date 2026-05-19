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
	"fmt"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/source/annotations"
)

const (
	IndexWithSelectors = "withSelectors"
)

type IndexSelectorOptions struct {
	annotationFilter labels.Selector
	labelSelector    labels.Selector
	// indexByLabelKey, when set, uses the value of this label as the index key
	// instead of the object's own name. Useful for indexing objects by a "parent"
	// resource they reference via a label (e.g. EndpointSlices by Service name).
	indexByLabelKey string
	conditions      []func(metav1.Object) bool
}

func IndexSelectorWithAnnotationFilter(input string) func(options *IndexSelectorOptions) {
	return func(options *IndexSelectorOptions) {
		if input == "" {
			return
		}
		selector, err := annotations.ParseFilter(input)
		if err != nil {
			return
		}
		options.annotationFilter = selector
	}
}

func IndexSelectorWithLabelSelector(input labels.Selector) func(options *IndexSelectorOptions) {
	return func(options *IndexSelectorOptions) {
		options.labelSelector = input
	}
}

// IndexSelectorWithConditions adds typed predicate functions to the index selector.
// Each function receives the concrete object type T; returning false excludes the
// object from the index. The type parameter is inferred from the function arguments.
func IndexSelectorWithConditions[T metav1.Object](fns ...func(T) bool) func(*IndexSelectorOptions) {
	return func(options *IndexSelectorOptions) {
		for _, fn := range fns {
			options.conditions = append(options.conditions, func(obj metav1.Object) bool {
				typed, ok := obj.(T)
				return ok && fn(typed)
			})
		}
	}
}

// IndexSelectorWithLabelKey configures IndexerWithOptions to use the value of the given
// label key as the index key instead of the object's own name. Useful when objects
// reference a "parent" resource via a label (e.g. EndpointSlices carry
// kubernetes.io/service-name pointing at their owning Service).
func IndexSelectorWithLabelKey(key string) func(options *IndexSelectorOptions) {
	return func(options *IndexSelectorOptions) {
		options.indexByLabelKey = key
	}
}

// IndexerWithOptions is a generic function that allows adding multiple indexers
// to a SharedIndexInformer for a specific Kubernetes resource type T. It accepts
// a variadic list of indexer functions, which define custom indexing logic.
//
// Each indexer function is applied to objects of type T, enabling flexible and
// reusable indexing based on annotations, labels, or other criteria.
//
// Example usage:
// err := IndexerWithOptions[*v1.Pod](
//
//	IndexSelectorWithAnnotationFilter("example-annotation"),
//	IndexSelectorWithLabelSelector(labels.SelectorFromSet(labels.Set{"app": "my-app"})),
//
// )
//
// This function ensures type safety and simplifies the process of adding
// custom indexers to informers.
func IndexerWithOptions[T metav1.Object](optFns ...func(options *IndexSelectorOptions)) cache.Indexers {
	options := IndexSelectorOptions{}
	for _, fn := range optFns {
		fn(&options)
	}

	return cache.Indexers{
		IndexWithSelectors: func(obj any) ([]string, error) {
			entity, ok := obj.(T)
			if !ok {
				return nil, fmt.Errorf("object is not of type %T", new(T))
			}

			if options.annotationFilter != nil && !options.annotationFilter.Matches(labels.Set(entity.GetAnnotations())) {
				return nil, nil
			}

			if options.labelSelector != nil && !options.labelSelector.Matches(labels.Set(entity.GetLabels())) {
				return nil, nil
			}
			for _, condition := range options.conditions {
				if !condition(entity) {
					return nil, nil
				}
			}
			name := entity.GetName()
			if options.indexByLabelKey != "" {
				name = entity.GetLabels()[options.indexByLabelKey]
				if name == "" {
					return nil, nil
				}
			}
			key := types.NamespacedName{Namespace: entity.GetNamespace(), Name: name}.String()
			return []string{key}, nil
		},
	}
}

// MustAddIndexers calls AddIndexers on the informer and panics on error.
// AddIndexers only errors if the informer has already been stopped, which is a
// programming error — callers must invoke it before factory.Start().
func MustAddIndexers(informer cache.SharedIndexInformer, indexers cache.Indexers) {
	if err := informer.AddIndexers(indexers); err != nil {
		panic(fmt.Sprintf("AddIndexers called on stopped informer: %v", err))
	}
}

// MustAddEventHandler calls AddEventHandler on the informer and logs a warning on error.
// AddEventHandler only errors if the informer has already been stopped. Unlike
// MustSetTransform and MustAddIndexers (which are called exclusively at setup time),
// AddEventHandler is also called at runtime (e.g. from Source.AddEventHandler), where a
// stopped informer may be a transient shutdown condition rather than a programming error.
func MustAddEventHandler(informer cache.SharedInformer, handler cache.ResourceEventHandler) {
	if _, err := informer.AddEventHandler(handler); err != nil {
		log.Warnf("AddEventHandler called on stopped informer: %v", err)
	}
}

// GetByKey retrieves an object of type T (metav1.Object) from the given cache.Indexer by its key.
// It returns the object and an error if the retrieval or type assertion fails.
// If the object does not exist, it returns the zero value of T and nil.
func GetByKey[T metav1.Object](indexer cache.Indexer, key string) (T, error) {
	var entity T
	obj, exists, err := indexer.GetByKey(key)
	if err != nil || !exists {
		return entity, err
	}

	entity, ok := obj.(T)
	if !ok {
		return entity, fmt.Errorf("object is not of type %T", new(T))
	}
	return entity, nil
}
