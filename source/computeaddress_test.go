/*
Copyright 2018 The Kubernetes Authors.

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
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/rest/fake"
	"net/http"
	"strings"
	"testing"
)

type ComputeAddressSuite struct {
	suite.Suite
}

func (suite *ComputeAddressSuite) SetupTest() {}

func addKnownTypesToSchema(scheme *runtime.Scheme, groupVersion schema.GroupVersion) error {
	scheme.AddKnownTypes(groupVersion,
		&v1beta1.ComputeAddress{},
		&v1beta1.ComputeAddressList{},
	)
	metav1.AddToGroupVersion(scheme, groupVersion)
	return nil
}

func fakeComputeAddressClient(name, namespace, address, addressType, description, ipVersion, location, resourceID string, annotations map[string]string, labels map[string]string, t *testing.T) rest.Interface {
	group := v1beta1.ComputeAddressGVK.Group
	version := v1beta1.ComputeAddressGVK.Version
	apiVersion := group + "/" + version
	kind := v1beta1.ComputeAddressGVK.Kind
	groupVersion := v1beta1.ComputeAddressGVK.GroupVersion()

	scheme := runtime.NewScheme()
	addKnownTypesToSchema(scheme, groupVersion)

	computeAddressList := v1beta1.ComputeAddressList{}
	computeAddress := &v1beta1.ComputeAddress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiVersion,
			Kind:       kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: annotations,
			Labels:      labels,
			Generation:  1,
		},
		Spec: v1beta1.ComputeAddressSpec{
			Address:     &address,
			AddressType: &addressType,
			Description: &description,
			IpVersion:   &ipVersion,
			Location:    location,
			ResourceID:  &resourceID,
		},
	}

	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	client := &fake.RESTClient{
		GroupVersion:         groupVersion,
		VersionedAPIPath:     "/apis/" + apiVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			codec := codecFactory.LegacyCodec(groupVersion)
			switch p, m := req.URL.Path, req.Method; {
			case p == "/apis/"+apiVersion+"/computeaddresses" && m == http.MethodGet:
				fallthrough
			case p == "/apis/"+apiVersion+"/namespaces/"+namespace+"/computeaddresses" && m == http.MethodGet:
				computeAddressList.Items = computeAddressList.Items[:0]
				computeAddressList.Items = append(computeAddressList.Items, *computeAddress)
				return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: objBody(codec, &computeAddressList)}, nil
			case strings.HasPrefix(p, "/apis/"+apiVersion+"/namespaces/") && strings.HasSuffix(p, "computeaddresses") && m == http.MethodGet:
				return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: objBody(codec, &computeAddressList)}, nil
			default:
				return nil, fmt.Errorf("unexpected request: %#v\n%#v", req.URL, req)
			}
		}),
	}

	return client
}

func TestComputeAddressSource(t *testing.T) {
	suite.Run(t, new(ComputeAddressSuite))
	t.Run("Interface", testComputeAddressSourceImplementsSource)
	t.Run("Endpoints", testComputeAddressSourceEndpoints)
}

// testComputeAddressSourceImplementsSource tests that computeAddressSource is a valid Source.
func testComputeAddressSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(computeAddressSource))
}

// testComputeAddressSourceEndpoints tests various scenarios of using ComputeAddress source.
func testComputeAddressSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title               string
		registeredNamespace string
		namespace           string
		name                string
		address             string
		addressType         string
		description         string
		ipVersion           string
		location            string
		resourceID          string
		expectEndpoints     bool
		expectError         bool
		annotationFilter    string
		labelFilter         string
		annotations         map[string]string
		labels              map[string]string
	}{
		{
			title:               "computeaddress within a specific namespace",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			expectEndpoints:     true,
			expectError:         false,
		},
		{
			title:               "no computeaddress within a specific namespace",
			namespace:           "foo",
			registeredNamespace: "bar",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			expectEndpoints:     false,
			expectError:         false,
		},
		{
			title:               "computeaddress with no address",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			expectEndpoints:     false,
			expectError:         false,
		},
		{
			title:               "valid computeaddress gvk with annotation and non-matching annotation filter",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			annotations:         map[string]string{"test": "that"},
			annotationFilter:    "test=filter_something_else",
			expectEndpoints:     false,
			expectError:         false,
		},
		{
			title:               "valid computeaddress gvk with annotation and matching annotation filter",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			annotations:         map[string]string{"test": "that"},
			annotationFilter:    "test=that",
			expectEndpoints:     true,
			expectError:         false,
		},
		{
			title:               "valid computeaddress gvk with label and non-matching label filter",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			labels:              map[string]string{"test": "that"},
			labelFilter:         "test=filter_something_else",
			expectEndpoints:     false,
			expectError:         false,
		},
		{
			title:               "valid computeaddress gvk with label and matching label filter",
			namespace:           "foo",
			registeredNamespace: "foo",
			name:                "foo",
			address:             "1.2.3.4",
			addressType:         "EXTERNAL",
			description:         "Test compute address",
			ipVersion:           "IPV4",
			location:            "global",
			resourceID:          "foo",
			labels:              map[string]string{"test": "that"},
			labelFilter:         "test=that",
			expectEndpoints:     true,
			expectError:         false,
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			restClient := fakeComputeAddressClient(ti.name, ti.namespace, ti.address, ti.addressType, ti.description, ti.ipVersion, ti.location, ti.resourceID, ti.annotations, ti.labels, t)

			scheme := runtime.NewScheme()
			addKnownTypesToSchema(scheme, v1beta1.ComputeAddressGVK.GroupVersion())

			cas, _ := NewComputeAddressSource(restClient, ti.namespace, ti.annotationFilter, ti.labelFilter, scheme)

			receivedEndpoints, err := cas.Endpoints(context.Background())
			if ti.expectError {
				require.Errorf(t, err, "Received err %v", err)
			} else {
				require.NoErrorf(t, err, "Received err %v", err)
			}

			if len(receivedEndpoints) == 0 && !ti.expectEndpoints {
				return
			}

			if err == nil {
				validateComputeAddressResource(t, cas, ti.expectError)
			}
		})
	}
}

func validateComputeAddressResource(t *testing.T, src Source, expectError bool) {
	cas := src.(*computeAddressSource)
	_, err := cas.List(context.Background(), &metav1.ListOptions{})
	if expectError {
		require.Errorf(t, err, "Received err %v", err)
	} else {
		require.NoErrorf(t, err, "Received err %v", err)
	}
}
