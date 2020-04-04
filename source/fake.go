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
	"fmt"
	"math/rand"
	"net"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
)

// fakeSource is an implementation of Source that provides dummy endpoints for
// testing/dry-running of dns providers without needing an attached Kubernetes cluster.
type fakeSource struct {
	dnsName string
}

const (
	defaultFQDNTemplate = "example.com"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewFakeSource creates a new fakeSource with the given config.
func NewFakeSource(fqdnTemplate string) (Source, error) {
	if fqdnTemplate == "" {
		fqdnTemplate = defaultFQDNTemplate
	}

	return &fakeSource{
		dnsName: fqdnTemplate,
	}, nil
}

func (sc *fakeSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}

// Endpoints returns endpoint objects.
func (sc *fakeSource) Endpoints() ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 10)

	for i := 0; i < 10; i++ {
		endpoints[i], _ = sc.generateEndpoint()
	}

	return endpoints, nil
}

func (sc *fakeSource) generateEndpoint() (*endpoint.Endpoint, error) {
	ep := endpoint.NewEndpoint(
		generateDNSName(4, sc.dnsName),
		endpoint.RecordTypeA,
		generateIPAddress(),
	)

	return ep, nil
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
