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

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
)

// TestAddToScheme ensures AddToScheme registers both our API types and the
// standard metav1 types (e.g. ListOptions) for this group version, so that
// client-go's parameter codec can encode them. Regression test for #6787:
// runtime.NewSchemeBuilder does not register meta types automatically, unlike
// controller-runtime's scheme.Builder.
func TestAddToScheme(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, AddToScheme(scheme))

	// Our own types must be registered.
	_, err := scheme.New(GroupVersion.WithKind("DNSEndpoint"))
	require.NoError(t, err)
	_, err = scheme.New(GroupVersion.WithKind("DNSEndpointList"))
	require.NoError(t, err)

	// metav1 types must be registered for the parameter codec to work.
	_, err = scheme.New(GroupVersion.WithKind("ListOptions"))
	assert.NoError(t, err, "metav1 types should be registered via AddToScheme")
}
