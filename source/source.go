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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

const (
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
	return input[annotations.AccessKey]
}

func getEndpointsTypeFromAnnotations(annots map[string]string) string {
	return annots[annotations.EndpointsTypeKey]
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

func (fn eventHandlerFunc) OnAdd(_ any, _ bool) { fn() }
func (fn eventHandlerFunc) OnUpdate(_, _ any)   { fn() }
func (fn eventHandlerFunc) OnDelete(_ any)      { fn() }
