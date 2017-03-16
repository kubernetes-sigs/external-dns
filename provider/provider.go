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

package provider

import (
	"net"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// Provider defines the interface DNS providers should implement.
type Provider interface {
	Records(zone string) ([]*endpoint.Endpoint, error)
	ApplyChanges(zone string, changes *plan.Changes) error
}

var metrics = struct {
	APICallsTotal *prometheus.CounterVec
}{
	APICallsTotal: prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Name:      "api_calls_total",
			Help:      "Total number of API calls to DNS providers",
		},
		[]string{"dns_provider"},
	),
}

func init() {
	prometheus.MustRegister(metrics.APICallsTotal)
}

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A for IPs and type CNAME for everything else.
func suitableType(ep *endpoint.Endpoint) string {
	if ep.RecordType != "" {
		return ep.RecordType
	}
	if net.ParseIP(ep.Target) != nil {
		return "A"
	}
	return "CNAME"
}

// ensureTrailingDot ensures that the hostname receives a trailing dot if it hasn't already.
func ensureTrailingDot(hostname string) string {
	if net.ParseIP(hostname) != nil {
		return hostname
	}

	return strings.TrimSuffix(hostname, ".") + "."
}
