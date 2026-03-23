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

/*
Note: currently only supports IP targets (A records), not hostname targets
*/

package source

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/types"
)

// fakeSource is an implementation of Source that provides dummy endpoints for
// testing/dry-running of dns providers without needing an attached Kubernetes cluster.
//
// +externaldns:source:name=fake
// +externaldns:source:category=Testing
// +externaldns:source:description=Provides dummy endpoints for testing and dry-running
// +externaldns:source:resources=Fake Endpoints
// +externaldns:source:filters=
// +externaldns:source:namespace=
// +externaldns:source:fqdn-template=true
// +externaldns:source:events=true
// +externaldns:source:provider-specific=false
type fakeSource struct {
	dnsName string
}

const (
	defaultFQDNTemplate = "example.com"
)

// NewFakeSource creates a new fakeSource with the given config.
func NewFakeSource(fqdnTemplate string) (Source, error) {
	if fqdnTemplate == "" {
		fqdnTemplate = defaultFQDNTemplate
	}

	return &fakeSource{
		dnsName: fqdnTemplate,
	}, nil
}

func (sc *fakeSource) AddEventHandler(_ context.Context, _ func()) {
}

// Endpoints returns endpoint objects.
func (sc *fakeSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 10)

	for i := range 10 {
		endpoints[i] = sc.generateEndpoint()
	}

	return MergeEndpoints(endpoints), nil
}

func (sc *fakeSource) generateEndpoint() *endpoint.Endpoint {
	ep := endpoint.NewEndpoint(
		generateDNSName(4, sc.dnsName),
		endpoint.RecordTypeA,
		generateIPAddress(),
	)
	ep.SetIdentifier = types.Fake
	ep.WithRefObject(events.NewObjectReference(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      types.Fake + "-" + ep.DNSName,
			Namespace: v1.NamespaceDefault,
		},
	}, types.Fake))
	return ep
}

func generateIPAddress() string {
	// 192.0.2.[1-255] is reserved by RFC 5737 for documentation and examples
	return net.IPv4(
		byte(192),
		byte(0),
		byte(2),
		byte(rand.Intn(253)+1),
	).String()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func generateDNSName(prefixLength int, dnsName string) string {
	prefixBytes := make([]rune, prefixLength)

	for i := range prefixBytes {
		prefixBytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	prefixStr := string(prefixBytes)

	return fmt.Sprintf("%s.%s", prefixStr, dnsName)
}
