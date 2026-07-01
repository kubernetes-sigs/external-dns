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

package source

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func TestParseNamespacedName(t *testing.T) {
	tests := []struct {
		name      string
		ingress   string
		wantNS    string
		wantName  string
		wantError bool
	}{
		{
			name:      "valid namespace and name",
			ingress:   "default/test-ingress",
			wantNS:    "default",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "only name provided",
			ingress:   "test-ingress",
			wantNS:    "",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "invalid format",
			ingress:   "default/test/ingress",
			wantNS:    "",
			wantName:  "",
			wantError: true,
		},
		{
			name:      "empty string",
			ingress:   "",
			wantNS:    "",
			wantName:  "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNS, gotName, err := ParseNamespacedName(tt.ingress)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantNS, gotNS)
			assert.Equal(t, tt.wantName, gotName)
		})
	}
}

// informerTransformObjectMeta returns an ObjectMeta populated with
// LastAppliedConfigAnnotation and ManagedFields — the fields that the
// informer transformers are expected to strip.
func informerTransformObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      "test",
		Namespace: metav1.NamespaceDefault,
		Annotations: map[string]string{
			corev1.LastAppliedConfigAnnotation: "test",
		},
		ManagedFields: []metav1.ManagedFieldsEntry{
			{Manager: "test-manager"},
		},
	}
}

// informerTransformHelperConfig is a struct that holds configuration for the informerTransformHelper functions.
type informerTransformHelperConfig struct {
	removedLastAppliedConfigAnnotation bool
	removedManagedFields               bool
	removedStatusConditions            bool
}

// informerTransformHelperOptions are options for configuring the behavior of the informerTransformHelper functions.
type informerTransformHelperOptions func(*informerTransformHelperConfig)

// withRemovedLastAppliedConfigAnnotation indicates that the informer transformer should have removed the
// "metadata.annotations['kubectl.kubernetes.io/last-applied']" annotation
func withRemovedLastAppliedConfigAnnotation() informerTransformHelperOptions {
	return func(config *informerTransformHelperConfig) {
		config.removedLastAppliedConfigAnnotation = true
	}
}

// withRemovedManagedFields indicates that the informer transformer should have removed the "metadata.managedFields" field
func withRemovedManagedFields() informerTransformHelperOptions {
	return func(config *informerTransformHelperConfig) {
		config.removedManagedFields = true
	}
}

// withRemovedStatusConditions indicates that the informer transformer should hav removed the "status.conditions" field
// from the cached object.
func withRemovedStatusConditions() informerTransformHelperOptions {
	return func(config *informerTransformHelperConfig) {
		config.removedStatusConditions = true
	}
}

// testDynamicInformerTransformHelper creates an unstructured object via the
// dynamic client and waits for it to appear in the generic informer cache, then
// asserts whether the informer transformer stripped the fields selected by opts.
func testDynamicInformerTransformHelper(
	t *testing.T,
	gvr schema.GroupVersionResource,
	client dynamic.Interface,
	informer informers.GenericInformer,
	opts ...informerTransformHelperOptions,
) {
	t.Helper()

	cfg := &informerTransformHelperConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	meta := informerTransformObjectMeta()

	obj := &unstructured.Unstructured{}
	obj.SetName(meta.Name)
	obj.SetNamespace(meta.Namespace)
	obj.SetAnnotations(meta.Annotations)
	obj.SetManagedFields(meta.ManagedFields)

	if cfg.removedStatusConditions {
		err := unstructured.SetNestedField(obj.Object, []any{}, "status", "conditions")
		require.NoError(t, err)
	}

	_, err := client.Resource(gvr).Namespace(obj.GetNamespace()).Create(t.Context(), obj, metav1.CreateOptions{})
	require.NoError(t, err)

	var gotObj runtime.Object
	require.Eventually(t, func() bool {
		var err error
		gotObj, err = informer.Lister().ByNamespace(obj.GetNamespace()).Get(obj.GetName())
		return err == nil
	}, 5*time.Second, 10*time.Millisecond)
	require.IsType(t, &unstructured.Unstructured{}, gotObj)

	assertObjectTransformedHelper(t, gotObj.(*unstructured.Unstructured), cfg)
}

// testInformerTransformHelper verifies that the informer transformer stripped
// the fields configured via opts from the single cached object.
func testInformerTransformHelper(t *testing.T, informer cache.SharedIndexInformer, obj metav1.Object, opts ...informerTransformHelperOptions) {
	t.Helper()

	cfg := &informerTransformHelperConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	item, exists, err := informer.GetStore().Get(obj)
	require.NoError(t, err)
	require.True(t, exists)
	require.Implements(t, (*metav1.Object)(nil), item)

	assertObjectTransformedHelper(t, item.(metav1.Object), cfg)
}

// assertObjectTransformedHelper asserts whether the informer transformer stripped the fields selected by cfg from obj.
func assertObjectTransformedHelper(t *testing.T, obj metav1.Object, cfg *informerTransformHelperConfig) {
	t.Helper()

	if cfg.removedLastAppliedConfigAnnotation {
		assert.NotContains(t, obj.GetAnnotations(), corev1.LastAppliedConfigAnnotation)
	} else {
		assert.Contains(t, obj.GetAnnotations(), corev1.LastAppliedConfigAnnotation)
	}

	if cfg.removedManagedFields {
		assert.Empty(t, obj.GetManagedFields())
	} else {
		assert.NotEmpty(t, obj.GetManagedFields())
	}

	if obj, ok := obj.(*unstructured.Unstructured); ok {
		conditions, found, err := unstructured.NestedFieldNoCopy(obj.Object, "status", "conditions")
		require.NoError(t, err)
		if cfg.removedStatusConditions {
			require.False(t, found)
			require.Nil(t, conditions)
		} else if found {
			require.NotNil(t, conditions)
		}
		return
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	require.True(t, val.IsValid(), "object is not valid")
	statusField := val.FieldByName("Status")
	if cfg.removedStatusConditions {
		require.True(t, statusField.IsValid(), "object does not have a Status field")
		condField := statusField.FieldByName("Conditions")
		require.True(t, condField.IsValid(), "Status does not have a Conditions field")
		require.IsType(t, reflect.Slice, condField.Kind())
		assert.Zero(t, condField.Len())
	} else if statusField.IsValid() {
		condField := statusField.FieldByName("Conditions")
		if condField.IsValid() {
			require.IsType(t, reflect.Slice, condField.Kind())
			assert.NotZero(t, condField.Len())
		}
	}
}
