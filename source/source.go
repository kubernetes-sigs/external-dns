/*
Copyright 2017 The Kubernetes Authors.

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
	"context"
	"fmt"
	"reflect"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

const (
	controllerAnnotationKey       = annotations.ControllerKey
	hostnameAnnotationKey         = annotations.HostnameKey
	accessAnnotationKey           = annotations.AccessKey
	endpointsTypeAnnotationKey    = annotations.EndpointsTypeKey
	targetAnnotationKey           = annotations.TargetKey
	ttlAnnotationKey              = annotations.TtlKey
	aliasAnnotationKey            = annotations.AliasKey
	ingressHostnameSourceKey      = annotations.IngressHostnameSourceKey
	controllerAnnotationValue     = annotations.ControllerValue
	internalHostnameAnnotationKey = annotations.InternalHostnameKey

	EndpointsTypeNodeExternalIP = "NodeExternalIP"
	EndpointsTypeHostIP         = "HostIP"
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error)
	// AddEventHandler adds an event handler that should be triggered if something in source changes
	AddEventHandler(context.Context, func())
}

type kubeObject interface {
	runtime.Object
	metav1.Object
}

func getAccessFromAnnotations(input map[string]string) string {
	return input[accessAnnotationKey]
}

func getEndpointsTypeFromAnnotations(annotations map[string]string) string {
	return annotations[endpointsTypeAnnotationKey]
}

func getLabelSelector(annotationFilter string) (labels.Selector, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	return metav1.LabelSelectorAsSelector(labelSelector)
}

func matchLabelSelector(selector labels.Selector, srcAnnotations map[string]string) bool {
	return selector.Matches(labels.Set(srcAnnotations))
}

type eventHandlerFunc func()

func (fn eventHandlerFunc) OnAdd(obj interface{}, isInInitialList bool) { fn() }
func (fn eventHandlerFunc) OnUpdate(oldObj, newObj interface{})         { fn() }
func (fn eventHandlerFunc) OnDelete(obj interface{})                    { fn() }

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

func waitForCacheSync(ctx context.Context, factory informerFactory) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %v", typ, ctx.Err())
			default:
				return fmt.Errorf("failed to sync %v", typ)
			}
		}
	}
	return nil
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

func waitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %v", typ, ctx.Err())
			default:
				return fmt.Errorf("failed to sync %v", typ)
			}
		}
	}
	return nil
}
