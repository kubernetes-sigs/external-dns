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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

type TransformOptions struct {
	specSelector    bool
	specExternalIps bool
	statusLb        bool
}

func TransformerWithOptions[T metav1.Object](optFns ...func(options *TransformOptions)) cache.TransformFunc {
	options := TransformOptions{}
	for _, fn := range optFns {
		fn(&options)
	}
	return func(obj any) (any, error) {
		// only transform if the object is a Service at the moment
		entity, ok := obj.(*corev1.Service)
		if !ok {
			return nil, nil
		}
		if entity.UID == "" {
			// Pod was already transformed and we must be idempotent.
			return entity, nil
		}
		svc := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:              entity.Name,
				Namespace:         entity.Namespace,
				DeletionTimestamp: entity.DeletionTimestamp,
			},
			Spec:   corev1.ServiceSpec{},
			Status: corev1.ServiceStatus{},
		}
		if options.specSelector {
			svc.Spec.Selector = entity.Spec.Selector
		}
		if options.specExternalIps {
			svc.Spec.ExternalIPs = entity.Spec.ExternalIPs
		}
		if options.statusLb {
			svc.Status.LoadBalancer = entity.Status.LoadBalancer
		}
		return svc, nil
	}
}

// TransformWithSpecSelector enables copying the Service's .spec.selector field.
func TransformWithSpecSelector() func(options *TransformOptions) {
	return func(options *TransformOptions) {
		options.specSelector = true
	}
}

// TransformWithSpecExternalIPs enables copying the Service's .spec.externalIPs field.
func TransformWithSpecExternalIPs() func(options *TransformOptions) {
	return func(options *TransformOptions) {
		options.specExternalIps = true
	}
}

// TransformWithStatusLoadBalancer enables copying the Service's .status.loadBalancer field.
func TransformWithStatusLoadBalancer() func(options *TransformOptions) {
	return func(options *TransformOptions) {
		options.statusLb = true
	}
}
