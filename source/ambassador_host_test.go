/*
Copyright 2019 The Kubernetes Authors.

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

	ambassador "github.com/datawire/ambassador/pkg/api/getambassador.io/v2"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"
)

type AmbassadorSuite struct {
	suite.Suite
}

func TestAmbassadorSource(t *testing.T) {
	suite.Run(t, new(AmbassadorSuite))
	t.Run("Interface", testAmbassadorSourceImplementsSource)
}

// testAmbassadorSourceImplementsSource tests that ambassadorHostSource is a valid Source.
func testAmbassadorSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(ambassadorHostSource))
}

func TestAmbassadorHostSource(t *testing.T) {
	fakeKubernetesClient := fakeKube.NewSimpleClientset()

	ambassadorScheme := runtime.NewScheme()

	ambassador.AddToScheme(ambassadorScheme)

	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(ambassadorScheme)

	ctx := context.Background()

	namespace := "test"

	host, err := createAmbassadorHost("test-host", "test-service")
	if err != nil {
		t.Fatalf("could not create host resource: %v", err)
	}

	{
		_, err := fakeDynamicClient.Resource(ambHostGVR).Namespace(namespace).Create(ctx, host, v1.CreateOptions{})
		if err != nil {
			t.Fatalf("could not create host: %v", err)
		}
	}

	ambassadorSource, err := NewAmbassadorHostSource(ctx, fakeDynamicClient, fakeKubernetesClient, namespace)
	if err != nil {
		t.Fatalf("could not create ambassador source: %v", err)
	}

	{
		_, err := ambassadorSource.Endpoints(ctx)
		if err != nil {
			t.Fatalf("could not collect ambassador source endpoints: %v", err)
		}
	}

}

func createAmbassadorHost(name, ambassadorService string) (*unstructured.Unstructured, error) {
	host := &ambassador.Host{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				ambHostAnnotation: ambassadorService,
			},
		},
	}
	obj := &unstructured.Unstructured{}
	uc, _ := newUnstructuredConverter()
	err := uc.scheme.Convert(host, obj, nil)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// TestParseAmbLoadBalancerService tests our parsing of Ambassador service info.
func TestParseAmbLoadBalancerService(t *testing.T) {
	vectors := []struct {
		input  string
		ns     string
		svc    string
		errstr string
	}{
		{"svc", "default", "svc", ""},
		{"ns/svc", "ns", "svc", ""},
		{"svc.ns", "ns", "svc", ""},
		{"svc.ns.foo.bar", "ns.foo.bar", "svc", ""},
		{"ns/svc/foo/bar", "", "", "invalid external-dns service: ns/svc/foo/bar"},
		{"ns/svc/foo.bar", "", "", "invalid external-dns service: ns/svc/foo.bar"},
		{"ns.foo/svc/bar", "", "", "invalid external-dns service: ns.foo/svc/bar"},
	}

	for _, v := range vectors {
		ns, svc, err := parseAmbLoadBalancerService(v.input)

		errstr := ""

		if err != nil {
			errstr = err.Error()
		}

		if v.ns != ns {
			t.Errorf("%s: got ns \"%s\", wanted \"%s\"", v.input, ns, v.ns)
		}

		if v.svc != svc {
			t.Errorf("%s: got svc \"%s\", wanted \"%s\"", v.input, svc, v.svc)
		}

		if v.errstr != errstr {
			t.Errorf("%s: got err \"%s\", wanted \"%s\"", v.input, errstr, v.errstr)
		}
	}
}
