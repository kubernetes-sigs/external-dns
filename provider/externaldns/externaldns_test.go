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

package externaldns

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"testing"
)

func GetCurrentDnsEndpointCr(ctx context.Context) (*endpoint.DNSEndpoint, error) {
	dnsEndpoint := &endpoint.DNSEndpoint{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "test.k8s.io/v1alpha1",
			Kind:       "DNSEndpoint",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:       "endpoints within a specific namespace",
			Namespace:  "test",
			Generation: 1,
		},
		Spec: endpoint.DNSEndpointSpec{
			Endpoints: []*endpoint.Endpoint{
				{DNSName: "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{DNSName: "efg.example.org",
					Targets:    endpoint.Targets{"11.22.33.44"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
		},
	}
	return dnsEndpoint, nil
}

func GetCurrentDnsEndpointCr1(ctx context.Context) (*endpoint.DNSEndpoint, error) {
	error := errors.New("Error while making k8s call")
	return nil, error
}

func invokeGetCurrentDnsEndpointCr1(ctx context.Context, e ExternalDNSInterface) (*endpoint.DNSEndpoint, error) {
	return GetCurrentDnsEndpointCr(ctx)
}

func invokeGetCurrentDnsEndpointCr2(ctx context.Context, e ExternalDNSInterface) (*endpoint.DNSEndpoint, error) {
	return GetCurrentDnsEndpointCr1(ctx)
}

func readCaCertFile(name string) ([]byte, error) {
	return nil, errors.New("")
}

func readNSFile(name string) ([]byte, error) {
	fmt.Println("came here......")
	return nil, errors.New("")
}

func getCrdClient(client kubernetes.Interface, config *rest.Config, version string, kind string) (*rest.RESTClient, *runtime.Scheme, error) {
	crdClient := &rest.RESTClient{}
	return crdClient, nil, nil
}

func readEmptyfile(name string) ([]byte, error) {
	bytes := make([]byte, 0)
	return bytes, nil
}

func TestExternalDnsGetRecords(t *testing.T) {
	ep := ExternalDNSProvider{}
	Cp = invokeGetCurrentDnsEndpointCr1
	endpoints, _ := ep.Records(context.Background())
	assert.Equal(t, 2, len(endpoints))
	assert.Equal(t, "abc.example.org", endpoints[0].DNSName)
	assert.Equal(t, "efg.example.org", endpoints[1].DNSName)

}

func TestExternalDnsNilRecords(t *testing.T) {
	ep := ExternalDNSProvider{}
	Cp = invokeGetCurrentDnsEndpointCr2
	var err error
	endpoints, err := ep.Records(context.Background())
	assert.Equal(t, 0, len(endpoints))
	assert.NotNil(t, err)
}

func refFunction(ctx context.Context, p *ExternalDNSProvider, epCr *endpoint.DNSEndpoint, req *rest.Request) error {

	for _, ep := range epCr.Spec.Endpoints {
		fmt.Println(ep)
	}
	return nil
}

func TestExternalDnsApplyChanges(t *testing.T) {
	plan := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "abc.example.org",
				RecordType: "A",
				Targets:    []string{"1.2.3.4"},
				RecordTTL:  180,
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "efg.example.org",
				RecordType: "A",
				Targets:    []string{""},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
	}
	Cp = invokeGetCurrentDnsEndpointCr1
	ep := ExternalDNSProvider{
		crdClient: &rest.RESTClient{},
	}
	makeReq = refFunction
	err := ep.ApplyChanges(context.Background(), plan)
	assert.NoError(t, err)
}

func TestNewExternalDns(t *testing.T) {
	config := &ExternalDNSProviderConfig{
		Namespace: "namespace",
		CaCrtPath: "CaCrtPath",
		TokenPath: "TokenPath",
	}
	readFile = readEmptyfile
	crdClient = getCrdClient
	provider, _ := NewExternalDNS(config, false)
	assert.NotNil(t, provider)

}

func TestNewExternalDnsWithNSNotExisting(t *testing.T) {
	config := &ExternalDNSProviderConfig{
		Namespace: "namespace",
		CaCrtPath: "CaCrtPath",
		TokenPath: "TokenPath",
	}
	readFile = readNSFile
	provider, _ := NewExternalDNS(config, false)
	assert.Nil(t, provider)

}

func TestNewExternalDnsWithCaCertNotExisting(t *testing.T) {
	config := &ExternalDNSProviderConfig{
		Namespace: "namespace",
		CaCrtPath: "CaCrtPath",
		TokenPath: "TokenPath",
	}
	readFile = readCaCertFile
	provider, _ := NewExternalDNS(config, false)
	assert.Nil(t, provider)

}
