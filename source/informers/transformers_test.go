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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
)

func TestTransformRemoveManagedFields(t *testing.T) {
	t.Run("removes managed fields from Service", func(t *testing.T) {
		svc := fakeService()
		require.NotEmpty(t, svc.ManagedFields)

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveManagedFields())
		got, err := transform(svc)
		require.NoError(t, err)
		result := got.(*corev1.Service)
		assert.Empty(t, result.ManagedFields)
		// unrelated fields must be preserved
		assert.NotEmpty(t, result.Name)
		assert.NotEmpty(t, result.Spec.Selector)
		assert.NotEmpty(t, result.Status.LoadBalancer.Ingress)
	})

	t.Run("removes managed fields from Pod", func(t *testing.T) {
		pod := fakePod()
		require.NotEmpty(t, pod.ManagedFields)

		transform := TransformerWithOptions[*corev1.Pod](TransformRemoveManagedFields())
		got, err := transform(pod)
		require.NoError(t, err)
		result := got.(*corev1.Pod)
		assert.Empty(t, result.ManagedFields)
		assert.NotEmpty(t, result.Name)
		assert.NotEmpty(t, result.Spec.NodeName)
	})

	t.Run("idempotent when managed fields already nil", func(t *testing.T) {
		svc := fakeService()
		svc.ManagedFields = nil

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveManagedFields())
		got, err := transform(svc)
		require.NoError(t, err)
		assert.Empty(t, got.(*corev1.Service).ManagedFields)
	})
}

func TestTransformRemoveLastAppliedConfig(t *testing.T) {
	t.Run("removes last-applied-configuration annotation", func(t *testing.T) {
		svc := fakeService()
		require.Contains(t, svc.Annotations, corev1.LastAppliedConfigAnnotation)

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveLastAppliedConfig())
		got, err := transform(svc)
		require.NoError(t, err)
		result := got.(*corev1.Service)
		assert.NotContains(t, result.Annotations, corev1.LastAppliedConfigAnnotation)
		// other annotations must survive
		assert.Contains(t, result.Annotations, "description")
		assert.Contains(t, result.Annotations, "external-dns.alpha.kubernetes.io/hostname")
	})

	t.Run("idempotent when annotation is absent", func(t *testing.T) {
		svc := fakeService()
		delete(svc.Annotations, corev1.LastAppliedConfigAnnotation)

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveLastAppliedConfig())
		got, err := transform(svc)
		require.NoError(t, err)
		assert.NotContains(t, got.(*corev1.Service).Annotations, corev1.LastAppliedConfigAnnotation)
	})
}

func TestTransformRemoveStatusConditions(t *testing.T) {
	t.Run("removes conditions from Service", func(t *testing.T) {
		svc := fakeService()
		require.NotEmpty(t, svc.Status.Conditions)

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveStatusConditions())
		got, err := transform(svc)
		require.NoError(t, err)
		result := got.(*corev1.Service)
		assert.Empty(t, result.Status.Conditions)
		// unrelated status fields must be preserved
		assert.NotEmpty(t, result.Status.LoadBalancer.Ingress)
	})

	t.Run("removes conditions from Pod", func(t *testing.T) {
		pod := fakePod()
		pod.Status.Conditions = []corev1.PodCondition{
			{Type: corev1.PodReady, Status: corev1.ConditionTrue},
		}
		require.NotEmpty(t, pod.Status.Conditions)

		transform := TransformerWithOptions[*corev1.Pod](TransformRemoveStatusConditions())
		got, err := transform(pod)
		require.NoError(t, err)
		assert.Empty(t, got.(*corev1.Pod).Status.Conditions)
	})

	t.Run("removes conditions from Node", func(t *testing.T) {
		node := fakeNode()
		require.NotEmpty(t, node.Status.Conditions)

		transform := TransformerWithOptions[*corev1.Node](TransformRemoveStatusConditions())
		got, err := transform(node)
		require.NoError(t, err)
		result := got.(*corev1.Node)
		assert.Empty(t, result.Status.Conditions)
		// Status.Addresses must be preserved
		assert.NotEmpty(t, result.Status.Addresses)
	})

	t.Run("no-op when conditions are already empty", func(t *testing.T) {
		svc := fakeService()
		svc.Status.Conditions = nil

		transform := TransformerWithOptions[*corev1.Service](TransformRemoveStatusConditions())
		got, err := transform(svc)
		require.NoError(t, err)
		assert.Empty(t, got.(*corev1.Service).Status.Conditions)
	})
}

func TestTransformKeepAnnotationPrefix(t *testing.T) {
	t.Run("keeps only matching prefix", func(t *testing.T) {
		pod := fakePod()
		require.Len(t, pod.Annotations, 3)

		transform := TransformerWithOptions[*corev1.Pod](TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"))
		got, err := transform(pod)
		require.NoError(t, err)
		result := got.(*corev1.Pod)
		assert.Equal(t, map[string]string{
			"external-dns.alpha.kubernetes.io/hostname": "pod.example.com",
		}, result.Annotations)
	})

	t.Run("multiple prefixes use OR logic", func(t *testing.T) {
		pod := fakePod()

		transform := TransformerWithOptions[*corev1.Pod](
			TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"),
			TransformKeepAnnotationPrefix("unrelated.io/"),
		)
		got, err := transform(pod)
		require.NoError(t, err)
		result := got.(*corev1.Pod)
		assert.Contains(t, result.Annotations, "external-dns.alpha.kubernetes.io/hostname")
		assert.Contains(t, result.Annotations, "unrelated.io/annotation")
		assert.NotContains(t, result.Annotations, corev1.LastAppliedConfigAnnotation)
	})

	t.Run("nil annotations map is left unchanged", func(t *testing.T) {
		pod := fakePod()
		pod.Annotations = nil

		transform := TransformerWithOptions[*corev1.Pod](TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"))
		got, err := transform(pod)
		require.NoError(t, err)
		assert.Nil(t, got.(*corev1.Pod).Annotations)
	})
}

func TestTransformerWithOptions_Combined(t *testing.T) {
	svc := fakeService()

	transform := TransformerWithOptions[*corev1.Service](
		TransformRemoveManagedFields(),
		TransformRemoveLastAppliedConfig(),
		TransformRemoveStatusConditions(),
		TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"),
	)
	got, err := transform(svc)
	require.NoError(t, err)
	result := got.(*corev1.Service)

	assert.Empty(t, result.ManagedFields)
	assert.Empty(t, result.Status.Conditions)
	assert.NotContains(t, result.Annotations, corev1.LastAppliedConfigAnnotation)
	assert.NotContains(t, result.Annotations, "description")
	assert.Contains(t, result.Annotations, "external-dns.alpha.kubernetes.io/hostname")
	// Spec and remaining Status fields are fully preserved
	assert.NotEmpty(t, result.Spec.Selector)
	assert.NotEmpty(t, result.Spec.ExternalIPs)
	assert.NotEmpty(t, result.Status.LoadBalancer.Ingress)
}

func TestTransformerWithOptions_TypeMismatch(t *testing.T) {
	t.Run("non-matching type returns nil", func(t *testing.T) {
		transform := TransformerWithOptions[*corev1.Service](TransformRemoveManagedFields())
		got, err := transform(fakePod())
		require.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("non-matching primitive returns nil", func(t *testing.T) {
		transform := TransformerWithOptions[*corev1.Service]()
		got, err := transform("not-a-service")
		require.NoError(t, err)
		assert.Nil(t, got)
	})
}

func TestTransformerWithOptions_Idempotent(t *testing.T) {
	svc := fakeService()
	transform := TransformerWithOptions[*corev1.Service](
		TransformRemoveManagedFields(),
		TransformRemoveLastAppliedConfig(),
		TransformRemoveStatusConditions(),
	)
	first, err := transform(svc)
	require.NoError(t, err)
	second, err := transform(first)
	require.NoError(t, err)

	r1 := first.(*corev1.Service)
	r2 := second.(*corev1.Service)
	assert.Empty(t, r1.ManagedFields)
	assert.Empty(t, r2.ManagedFields)
	assert.NotContains(t, r1.Annotations, corev1.LastAppliedConfigAnnotation)
	assert.NotContains(t, r2.Annotations, corev1.LastAppliedConfigAnnotation)
	assert.Empty(t, r1.Status.Conditions)
	assert.Empty(t, r2.Status.Conditions)
}

func TestTransformerWithOptions_WithFakeClient(t *testing.T) {
	ctx := t.Context()
	svc := fakeService()
	fakeClient := fake.NewClientset()

	_, err := fakeClient.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{})
	require.NoError(t, err)

	factory := kubeinformers.NewSharedInformerFactoryWithOptions(fakeClient, 0, kubeinformers.WithNamespace(svc.Namespace))
	serviceInformer := factory.Core().V1().Services()
	err = serviceInformer.Informer().SetTransform(TransformerWithOptions[*corev1.Service](
		TransformRemoveManagedFields(),
		TransformRemoveLastAppliedConfig(),
		TransformRemoveStatusConditions(),
	))
	require.NoError(t, err)

	factory.Start(ctx.Done())
	err = WaitForCacheSync(ctx, factory)
	require.NoError(t, err)

	got, err := serviceInformer.Lister().Services(svc.Namespace).Get(svc.Name)
	require.NoError(t, err)

	assert.Empty(t, got.ManagedFields)
	assert.Empty(t, got.Status.Conditions)
	assert.NotContains(t, got.Annotations, corev1.LastAppliedConfigAnnotation)
	// TypeMeta is populated by the transformer
	assert.Equal(t, "Service", got.Kind)
	assert.NotEmpty(t, got.APIVersion)
	// Spec and remaining Status are preserved
	assert.Equal(t, svc.Spec.Selector, got.Spec.Selector)
	assert.Equal(t, svc.Spec.ExternalIPs, got.Spec.ExternalIPs)
	assert.Equal(t, svc.Status.LoadBalancer.Ingress, got.Status.LoadBalancer.Ingress)
}

func TestPopulateGVK(t *testing.T) {
	t.Run("populates Kind and APIVersion on Service", func(t *testing.T) {
		svc := fakeService()
		require.Empty(t, svc.Kind)

		populateGVK(svc)
		assert.Equal(t, "Service", svc.Kind)
		assert.NotEmpty(t, svc.APIVersion)
	})

	t.Run("populates Kind and APIVersion on Node", func(t *testing.T) {
		node := fakeNode()
		require.Empty(t, node.Kind)

		populateGVK(node)
		assert.Equal(t, "Node", node.Kind)
		assert.NotEmpty(t, node.APIVersion)
	})

	t.Run("idempotent when GVK already set", func(t *testing.T) {
		svc := fakeService()
		svc.Kind = "Service"
		svc.APIVersion = "v1"

		populateGVK(svc)
		assert.Equal(t, "Service", svc.Kind)
		assert.Equal(t, "v1", svc.APIVersion)
	})
}
