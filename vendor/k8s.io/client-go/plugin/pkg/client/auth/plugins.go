/*
Copyright 2016 The Kubernetes Authors.

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

package auth

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// Initialize common client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
||||||| parent of 5ce8c7613 (update vendored files)
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
=======
	// Initialize common client auth plugins.
>>>>>>> 5ce8c7613 (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
<<<<<<< HEAD
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
||||||| parent of 6b7ce455e (update vendored files)
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
=======
	// Initialize common client auth plugins.
>>>>>>> 6b7ce455e (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
<<<<<<< HEAD
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
||||||| parent of 4d7e5ad26 (update vendored files)
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
=======
	// Initialize common client auth plugins.
>>>>>>> 4d7e5ad26 (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
<<<<<<< HEAD
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Initialize all known client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
)
