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
	"testing"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/source/informers"
)

// informerFor returns the informer watching the given namespace, failing the test when none does.
func informerFor[I informers.SharedInformer](t *testing.T, set *informers.Informers[I], namespace string) I {
	t.Helper()
	informer, ok := set.For(namespace)
	require.True(t, ok, "no informer watching namespace %q", namespace)
	return informer
}

// firstInformer returns the first informer of the collection, for sources built with a
// single namespace.
func firstInformer[I informers.SharedInformer](set *informers.Informers[I]) I {
	return set.All()[0]
}

// singleInformer wraps one informer as a collection, for sources assembled by hand in tests.
func singleInformer[I informers.SharedInformer](informer I) *informers.Informers[I] {
	return informers.Single(informer)
}
