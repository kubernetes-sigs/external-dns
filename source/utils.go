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
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

// ParseIngress parses an ingress string in the format "namespace/name" or "name".
// It returns the namespace and name extracted from the string, or an error if the format is invalid.
// If the namespace is not provided, it defaults to an empty string.
func ParseIngress(ingress string) (string, string, error) {
	var namespace, name string
	var err error
	parts := strings.Split(ingress, "/")
	switch len(parts) {
	case 2:
		namespace, name = parts[0], parts[1]
	case 1:
		name = parts[0]
	default:
		err = fmt.Errorf("invalid ingress name (name or namespace/name) found %q", ingress)
	}

	return namespace, name, err
}

// hasNoEmptyEndpoints checks if the endpoint list is empty and logs
// a debug message if so. Returns true if empty, false otherwise.
func hasNoEmptyEndpoints(
	endpoints []*endpoint.Endpoint,
	rType string, entity metav1.ObjectMetaAccessor,
) bool {
	if len(endpoints) == 0 {
		log.Debugf("No endpoints could be generated from '%s/%s/%s'", rType, entity.GetObjectMeta().GetNamespace(), entity.GetObjectMeta().GetName())
		return true
	}
	return false
}
