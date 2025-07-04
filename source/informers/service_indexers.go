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
	"crypto/sha1"
	"encoding/hex"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	SvcSpecSelectorIndex          = "spec.selector"
	SvcNamespaceSpecSelectorIndex = "namespace/spec.selector"
)

var (
	// ServiceIndexers of indexers to allow fast lookups of services by their spec.selector.
	// This indexer is used to find services that match a specific label selector.
	// Usage:
	//   serviceInformer.Informer().AddIndexers(ServiceIndexers)
	//   serviceInformer.Lister().ByIndex(SvcSpecSelectorIndex, ToSHA(labels.Set(selector).String()))
	ServiceIndexers = cache.Indexers{
		SvcSpecSelectorIndex: func(obj any) ([]string, error) {
			svc, ok := obj.(*corev1.Service)
			if !ok {
				return nil, nil
			}
			return []string{ToSHA(labels.Set(svc.Spec.Selector).String())}, nil
		},
	}
	// ServiceNsSelectorIndexers for namespace/spec.selector to allow fast lookups of services by their spec.selector and namespace.
	// Usage:
	//      serviceInformer.Informer().AddIndexers(ServiceNsSelectorIndexers)
	//      serviceInformer.Lister().ByIndex(SvcNamespaceSpecSelectorIndex, ToSHA(svc.Namespace + "/" + labels.Set(svc.Spec.Selector).String()))
	ServiceNsSelectorIndexers = cache.Indexers{
		SvcNamespaceSpecSelectorIndex: func(obj any) ([]string, error) {
			svc, ok := obj.(*corev1.Service)
			if !ok {
				return nil, nil
			}
			return []string{ToSHA(svc.Namespace + "/" + labels.Set(svc.Spec.Selector).String())}, nil
		},
	}
)

func ServiceWithDefaultOptions(serviceInformer corev1informers.ServiceInformer, namespace string) error {
	_, _ = serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj any) {
				if !log.IsLevelEnabled(log.DebugLevel) {
					return
				}
				u, ok := obj.(*unstructured.Unstructured)
				if !ok {
					log.Debugf("%s added", u.GetKind())
				} else {
					log.WithFields(log.Fields{
						"apiVersion": u.GetAPIVersion(),
						"kind":       u.GetKind(),
						"name":       u.GetName(),
						"namespace":  u.GetNamespace(),
					}).Info("added")
				}
			},
		},
	)

	if namespace == "" {
		return serviceInformer.Informer().AddIndexers(ServiceIndexers)
	}
	return serviceInformer.Informer().AddIndexers(ServiceNsSelectorIndexers)
}

// ToSHA returns the SHA1 hash of the input string as a hex string.
// Using a SHA1 hash of the label selector string (as in ToSHA(labels.Set(selector).String())) is useful:
// - It provides a consistent and compact representation of the selector.
// - It allows for efficient indexing and lookup in Kubernetes informers.
// - It avoids issues with long label selector strings that could exceed index length limits.
func ToSHA(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
