package source

import (
	contour "github.com/projectcontour/contour/apis/contour/v1beta1"
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
	_ = contour.AddToScheme(uc.scheme)
	_ = projectcontour.AddToScheme(uc.scheme)

	// Add the core types we need
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}
