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

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestAddToSchemeRegistersMetaTypes(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, AddToScheme(scheme))

	codec := runtime.NewParameterCodec(scheme)
	for _, opts := range []runtime.Object{
		&metav1.CreateOptions{},
		&metav1.UpdateOptions{},
		&metav1.PatchOptions{},
		&metav1.DeleteOptions{},
		&metav1.GetOptions{},
		&metav1.ListOptions{},
	} {
		_, err := codec.EncodeParameters(opts, GroupVersion)
		require.NoError(t, err, "encoding %T for %s", opts, GroupVersion)
	}

	// Callers that still register the meta types themselves must keep working.
	metav1.AddToGroupVersion(scheme, GroupVersion)
	_, err := codec.EncodeParameters(&metav1.ListOptions{}, GroupVersion)
	require.NoError(t, err)
}
