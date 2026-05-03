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
	"context"
	"fmt"
	"math/rand"
	"net"

	log "github.com/sirupsen/logrus"
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
	dnsNames []string
	rand     *rand.Rand
}

const (
	defaultFQDNTemplate = "example.com"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	// fakePod is a placeholder Pod used when rendering the FQDN template.
	fakePod = v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      types.Fake,
			Namespace: v1.NamespaceDefault,
		},
	}
)

// NewFakeSource creates a new fakeSource with the given config.
func NewFakeSource(cfg *Config) (Source, error) {
	dnsNames := []string{defaultFQDNTemplate}
	if cfg.TemplateEngine.IsConfigured() {
		hostnames, err := cfg.TemplateEngine.ExecFQDN(&fakePod)
		if err != nil {
			return nil, fmt.Errorf("rendering fqdn template: %w", err)
		}
		if len(hostnames) > 0 {
			dnsNames = hostnames
		}
	}
	rand := rand.New(rand.NewSource(9673))
	return &fakeSource{dnsNames: dnsNames, rand: rand}, nil
}

func (sc *fakeSource) AddEventHandler(_ context.Context, _ func()) {
}

// Endpoints returns one endpoint per supported DNS record type per configured domain.
// A and AAAA records carry one target per domain.
func (sc *fakeSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 0, len(sc.dnsNames)*len(endpoint.KnownRecordTypes))
	for _, dnsName := range sc.dnsNames {
		for _, recordType := range endpoint.KnownRecordTypes {
			ep, err := sc.generateEndpointForType(recordType, dnsName)
			if err != nil {
				return nil, fmt.Errorf("generating %s endpoint: %w", recordType, err)
			}
			if ep != nil {
				endpoints = append(endpoints, ep)
			}
		}
	}
	return MergeEndpoints(endpoints), nil
}

func (sc *fakeSource) generateEndpointForType(recordType, dnsName string) (*endpoint.Endpoint, error) {
	var ep *endpoint.Endpoint

	switch recordType {
	case endpoint.RecordTypeA:
		ep = endpoint.NewEndpoint(sc.generateDNSName(4, dnsName), endpoint.RecordTypeA, generateTargets(len(sc.dnsNames), sc.generateIPv4Address)...)
	case endpoint.RecordTypeAAAA:
		ep = endpoint.NewEndpoint(sc.generateDNSName(4, dnsName), endpoint.RecordTypeAAAA, generateTargets(len(sc.dnsNames), sc.generateIPv6Address)...)
	case endpoint.RecordTypeCNAME:
		ep = endpoint.NewEndpoint(sc.generateDNSName(4, dnsName), endpoint.RecordTypeCNAME, sc.generateDNSName(4, dnsName))
	case endpoint.RecordTypeTXT:
		ep = endpoint.NewEndpoint(sc.generateDNSName(4, dnsName), endpoint.RecordTypeTXT, `"heritage=external-dns,external-dns/owner=fake"`)
	case endpoint.RecordTypeSRV:
		// SRV target format: "priority weight port target." (target must end with a dot per RFC 2782)
		name := sc.generateDNSName(4, dnsName)
		ep = endpoint.NewEndpoint(fmt.Sprintf("_sip._udp.%s", dnsName), endpoint.RecordTypeSRV, fmt.Sprintf("10 20 5060 %s.", name))
	case endpoint.RecordTypeNS:
		ep = endpoint.NewEndpoint(dnsName, endpoint.RecordTypeNS, sc.generateDNSName(3, dnsName))
	case endpoint.RecordTypePTR:
		name := sc.generateDNSName(4, dnsName)
		var err error
		ep, err = endpoint.NewPTREndpoint(sc.generateIPv4Address(), endpoint.TTL(0), name)
		if err != nil {
			return nil, err
		}
	case endpoint.RecordTypeMX:
		ep = endpoint.NewEndpoint(dnsName, endpoint.RecordTypeMX, fmt.Sprintf("10 %s", sc.generateDNSName(4, dnsName)))
	case endpoint.RecordTypeNAPTR:
		// NAPTR target format: "order preference flags service regexp replacement"
		ep = endpoint.NewEndpoint(fmt.Sprintf("_sip._udp.%s", dnsName), endpoint.RecordTypeNAPTR, fmt.Sprintf(`100 10 "u" "E2U+sip" "!^.*$!sip:info@%s!" .`, dnsName))
	default:
		return nil, fmt.Errorf("unsupported record type: %s", recordType)
	}

	if ep != nil {
		pod := fakePod
		pod.Name = fakePodName(ep.DNSName)
		ep.SetIdentifier = types.Fake
		ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("%s/%s/%s", types.Fake, pod.Namespace, ep.DNSName))
		ep.WithRefObject(events.NewObjectReference(&pod, types.Fake))
		log.Debugf("fake source generated %s endpoint: %s -> %v", ep.RecordType, ep.DNSName, ep.Targets)
	}
	return ep, nil
}

func (sc *fakeSource) generateIPv4Address() string {
	// 192.0.2.[1-254] is reserved by RFC 5737 for documentation and examples
	return net.IPv4(
		byte(192),
		byte(0),
		byte(2),
		byte(sc.rand.Intn(253)+1),
	).String()
}

func (sc *fakeSource) generateIPv6Address() string {
	// 2001:db8::/32 is reserved by RFC 3849 for documentation and examples
	return fmt.Sprintf("2001:db8::%x:%x", sc.rand.Intn(0xffff)+1, sc.rand.Intn(0xffff)+1)
}

// fakePodName returns a valid Kubernetes object name for the fake Pod associated
// with an endpoint. Names are capped at 253 characters (RFC 1123 subdomain limit).
func fakePodName(dnsName string) string {
	const prefix = types.Fake + "-"
	const maxLen = 253
	name := prefix + dnsName
	if len(name) > maxLen {
		name = name[:maxLen]
	}
	return name
}

func generateTargets(n int, gen func() string) []string {
	targets := make([]string, n)
	for i := range targets {
		targets[i] = gen()
	}
	return targets
}

func (sc *fakeSource) generateDNSName(prefixLength int, dnsName string) string {
	prefix := make([]rune, prefixLength)
	for i := range prefix {
		prefix[i] = letterRunes[sc.rand.Intn(len(letterRunes))]
	}
	return fmt.Sprintf("%s.%s", string(prefix), dnsName)
}
