/*
Copyright 2026 The Kubernetes Authors.

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

package crd

import "sigs.k8s.io/controller-runtime/pkg/client"

// CRDClients bundles the reader (cache-backed, for bulk List) and writer (direct,
// for targeted Get/status writes) built for the crd source, for reuse by other
// components instead of building a second, independent client.
type CRDClients struct {
	Reader client.Reader
	Writer client.Client
}
