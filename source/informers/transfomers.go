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
	"maps"
	"reflect"
	"slices"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
)

// TransformOptions holds the configuration for TransformerWithOptions.
// All options operate on the metav1.Object interface (or via reflection for Status
// fields) and are therefore applicable to any Kubernetes resource type.
type TransformOptions struct {
	removeManagedFields     bool
	removeLastAppliedConfig bool
	removeStatusConditions  bool
	keepAnnotationPrefixes  []string
}

// TransformRemoveManagedFields strips managedFields from the object's metadata.
// managedFields are written by server-side apply and can be megabytes per object.
func TransformRemoveManagedFields() func(*TransformOptions) {
	return func(o *TransformOptions) {
		o.removeManagedFields = true
	}
}

// TransformRemoveLastAppliedConfig removes the kubectl.kubernetes.io/last-applied-configuration
// annotation, which stores a full JSON snapshot of the resource and can be very large.
func TransformRemoveLastAppliedConfig() func(*TransformOptions) {
	return func(o *TransformOptions) {
		o.removeLastAppliedConfig = true
	}
}

// TransformRemoveStatusConditions clears the Status.Conditions field if the object has one.
// Conditions vary by type (metav1.Condition, corev1.PodCondition, corev1.NodeCondition, …)
// so this is applied via reflection and is a no-op for types without Status.Conditions.
func TransformRemoveStatusConditions() func(*TransformOptions) {
	return func(o *TransformOptions) {
		o.removeStatusConditions = true
	}
}

// TransformKeepAnnotationPrefix retains only annotations whose keys match at least one of
// the given prefixes, discarding all others. Multiple calls accumulate prefixes (OR logic).
func TransformKeepAnnotationPrefix(prefix string) func(*TransformOptions) {
	return func(o *TransformOptions) {
		o.keepAnnotationPrefixes = append(o.keepAnnotationPrefixes, prefix)
	}
}

// TransformerWithOptions returns a cache.TransformFunc that modifies objects of type T
// in place to reduce the memory footprint of the informer cache. All options operate
// on the metav1.Object interface or via reflection, making the transformer applicable
// to any Kubernetes resource type without type-specific logic.
//
// The transformer also populates TypeMeta (Kind/APIVersion) on every object. Kubernetes
// informers strip TypeMeta when returning objects because the client already knows the
// type — populating it here makes cached objects self-describing for templates and logging.
//
// The transform is naturally idempotent: nil-ing an already-nil field and filtering an
// already-filtered map are both no-ops, so calling it multiple times on the same object
// is safe.
//
// Example:
//
//	serviceInformer.Informer().SetTransform(informers.TransformerWithOptions[*corev1.Service](
//	    informers.TransformRemoveManagedFields(),
//	    informers.TransformRemoveLastAppliedConfig(),
//	))
func TransformerWithOptions[T interface {
	metav1.Object
	runtime.Object
}](optFns ...func(*TransformOptions)) cache.TransformFunc {
	options := TransformOptions{}
	for _, fn := range optFns {
		fn(&options)
	}
	return func(obj any) (any, error) {
		entity, ok := obj.(T)
		if !ok {
			return nil, nil
		}
		populateGVK(entity)
		if options.removeManagedFields {
			entity.SetManagedFields(nil)
		}
		if len(options.keepAnnotationPrefixes) > 0 || options.removeLastAppliedConfig {
			anns := entity.GetAnnotations()
			if options.removeLastAppliedConfig {
				delete(anns, corev1.LastAppliedConfigAnnotation)
			}
			if len(options.keepAnnotationPrefixes) > 0 {
				maps.DeleteFunc(anns, func(k, _ string) bool {
					return !slices.ContainsFunc(options.keepAnnotationPrefixes, func(prefix string) bool {
						return strings.HasPrefix(k, prefix)
					})
				})
			}
			entity.SetAnnotations(anns)
		}
		if options.removeStatusConditions {
			clearStatusConditions(entity)
		}
		return entity, nil
	}
}

// populateGVK sets TypeMeta (Kind/APIVersion) on obj if it is missing.
// Kubernetes informers strip TypeMeta from cached objects because the client already
// knows what type it requested. Populating it here makes cached objects self-describing
// for templates and logging without any per-reconciliation overhead.
func populateGVK(obj runtime.Object) {
	if obj.GetObjectKind().GroupVersionKind().Kind != "" {
		return
	}
	gvks, _, err := scheme.Scheme.ObjectKinds(obj)
	if err == nil && len(gvks) > 0 {
		obj.GetObjectKind().SetGroupVersionKind(gvks[0])
	} else {
		// Fallback to reflection for types not registered in the scheme (e.g. CRDs)
		obj.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
			Kind: reflect.TypeOf(obj).Elem().Name(),
		})
	}
}

// clearStatusConditions zeroes out the Status.Conditions field on obj if it exists.
// It handles all condition types (metav1.Condition, corev1.PodCondition, etc.) uniformly.
// Reflection is used because Status.Conditions is a structural convention shared by all
// Kubernetes types but not codified in any interface — element types differ per resource.
// The reflection cost is negligible: paid once per object at cache-population time,
// not on the hot path of every endpoint reconciliation.
func clearStatusConditions(obj any) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if !val.IsValid() {
		return
	}
	statusField := val.FieldByName("Status")
	if !statusField.IsValid() {
		return
	}
	// Status may itself be a pointer (dereference if so)
	if statusField.Kind() == reflect.Ptr {
		if statusField.IsNil() {
			return
		}
		statusField = statusField.Elem()
	}
	condField := statusField.FieldByName("Conditions")
	if condField.IsValid() && condField.CanSet() {
		condField.Set(reflect.Zero(condField.Type()))
	}
}
