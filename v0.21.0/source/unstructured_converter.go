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
	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// UnstructuredConverter handles conversions between unstructured.Unstructured and Contour types
type UnstructuredConverter struct {
	// scheme holds an initializer for converting Unstructured to a type
	scheme *runtime.Scheme
}

// NewUnstructuredConverter returns a new UnstructuredConverter initialized
func NewUnstructuredConverter() (*UnstructuredConverter, error) {
	uc := &UnstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Setup converter to understand custom CRD types
	_ = projectcontour.AddToScheme(uc.scheme)

	// Add the core types we need
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}
