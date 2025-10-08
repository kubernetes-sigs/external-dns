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

func TestTransformerWithOptions_Service(t *testing.T) {
	base := fakeService()

	tests := []struct {
		name    string
		options []func(*TransformOptions)
		asserts func(any)
	}{
		{
			name:    "minimalistic object",
			options: nil,
			asserts: func(obj any) {
				svc, ok := obj.(*corev1.Service)
				assert.True(t, ok)
				assert.Empty(t, svc.UID)
				assert.NotEmpty(t, svc.Name)
				assert.NotEmpty(t, svc.Namespace)
			},
		},
		{
			name:    "with selector",
			options: []func(*TransformOptions){TransformWithSpecSelector()},
			asserts: func(obj any) {
				svc, ok := obj.(*corev1.Service)
				assert.True(t, ok)
				assert.NotEmpty(t, svc.Spec.Selector)
				assert.Empty(t, svc.Spec.ExternalIPs)
				assert.Empty(t, svc.Status.LoadBalancer.Ingress)
			},
		},
		{
			name:    "with selector",
			options: []func(*TransformOptions){TransformWithSpecSelector()},
			asserts: func(obj any) {
				svc, ok := obj.(*corev1.Service)
				assert.True(t, ok)
				assert.NotEmpty(t, svc.Spec.Selector)
				assert.Empty(t, svc.Spec.ExternalIPs)
				assert.Empty(t, svc.Status.LoadBalancer.Ingress)
			},
		},
		{
			name:    "with loadBalancer",
			options: []func(*TransformOptions){TransformWithStatusLoadBalancer()},
			asserts: func(obj any) {
				svc, ok := obj.(*corev1.Service)
				assert.True(t, ok)
				assert.Empty(t, svc.Spec.Selector)
				assert.Empty(t, svc.Spec.ExternalIPs)
				assert.NotEmpty(t, svc.Status.LoadBalancer.Ingress)
			},
		},
		{
			name: "all options",
			options: []func(*TransformOptions){
				TransformWithSpecSelector(),
				TransformWithSpecExternalIPs(),
				TransformWithStatusLoadBalancer(),
			},
			asserts: func(obj any) {
				svc, ok := obj.(*corev1.Service)
				assert.True(t, ok)
				assert.NotEmpty(t, svc.Spec.Selector)
				assert.NotEmpty(t, svc.Spec.ExternalIPs)
				assert.NotEmpty(t, svc.Status.LoadBalancer.Ingress)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transform := TransformerWithOptions[*corev1.Service](tt.options...)
			got, err := transform(base)
			require.NoError(t, err)
			tt.asserts(got)
		})
	}

	t.Run("non-service input", func(t *testing.T) {
		transform := TransformerWithOptions[*corev1.Service]()
		out, err := transform("not-a-service")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out != nil {
			t.Errorf("expected nil output for non-service input, got %v", out)
		}
	})
}

func TestTransformer_Service_WithFakeClient(t *testing.T) {
	t.Run("with transformer", func(t *testing.T) {
		ctx := t.Context()
		svc := fakeService()
		fakeClient := fake.NewClientset()

		_, err := fakeClient.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{})
		require.NoError(t, err)

		factory := kubeinformers.NewSharedInformerFactoryWithOptions(fakeClient, 0, kubeinformers.WithNamespace(svc.Namespace))
		serviceInformer := factory.Core().V1().Services()
		err = serviceInformer.Informer().SetTransform(TransformerWithOptions[*corev1.Service](
			TransformWithSpecSelector(),
			TransformWithSpecExternalIPs(),
			TransformWithStatusLoadBalancer(),
		))
		require.NoError(t, err)

		factory.Start(ctx.Done())
		err = WaitForCacheSync(ctx, factory)
		require.NoError(t, err)

		got, err := serviceInformer.Lister().Services(svc.Namespace).Get(svc.Name)
		require.NoError(t, err)

		assert.Equal(t, svc.Spec.Selector, got.Spec.Selector)
		assert.Equal(t, svc.Spec.ExternalIPs, got.Spec.ExternalIPs)
		assert.Equal(t, svc.Status.LoadBalancer.Ingress, got.Status.LoadBalancer.Ingress)
		assert.NotEqual(t, svc.Annotations, got.Annotations)
		assert.NotEqual(t, svc.Labels, got.Labels)
	})

	t.Run("without transformer", func(t *testing.T) {
		ctx := t.Context()
		svc := fakeService()
		fakeClient := fake.NewClientset()

		_, err := fakeClient.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{})
		require.NoError(t, err)

		factory := kubeinformers.NewSharedInformerFactoryWithOptions(fakeClient, 0, kubeinformers.WithNamespace(svc.Namespace))
		serviceInformer := factory.Core().V1().Services()

		err = serviceInformer.Informer().GetIndexer().Add(svc)
		require.NoError(t, err)

		factory.Start(ctx.Done())
		err = WaitForCacheSync(ctx, factory)
		require.NoError(t, err)

		got, err := serviceInformer.Lister().Services(svc.Namespace).Get(svc.Name)
		require.NoError(t, err)

		assert.Equal(t, map[string]string{"app": "demo"}, got.Spec.Selector)
		assert.Equal(t, []string{"1.2.3.4"}, got.Spec.ExternalIPs)
		assert.Equal(t, svc.Status.LoadBalancer.Ingress, got.Status.LoadBalancer.Ingress)
		assert.Equal(t, svc.Annotations, got.Annotations)
		assert.Equal(t, svc.Labels, got.Labels)
	})
}
