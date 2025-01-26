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

package leaderelection

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/uuid"
)

var (
	hostname               = os.Hostname
	inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
)

// getInClusterNamespace retrieves the namespace in which the pod is running by reading the namespace file.
// If the namespace file does not exist or cannot be read, it falls back to the default namespace.
func getNamespace() string {
	// Check whether the namespace file exists.
	// If not, we are not running in cluster so can't guess the namespace.
	namespace, err := os.ReadFile(inClusterNamespacePath)
	if err != nil {
		log.Debugf("not running in-cluster, fallback to %s namespace\n", DefaultNamespace)
		return DefaultNamespace
	}
	return string(namespace)
}

// getLeaderId generates a unique leader ID by combining the hostname and a UUID.
// If the hostname cannot be retrieved, it falls back to using "external-dns" with a UUID.
func getLeaderId() string {
	id, err := hostname()
	if err != nil {
		return fmt.Sprintf("external-dns_%s", uuid.NewUUID())
	}
	return fmt.Sprintf("%s_%s", id, uuid.NewUUID())
}
