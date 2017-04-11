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

package endpoint

import "strings"

const (
	// OwnerLabelKey is the name of the label that defines the owner of an Endpoint.
	OwnerLabelKey = "owner"
)

// Endpoint is a high-level way of a connection between a service and an IP
type Endpoint struct {
	// The hostname of the DNS record
	DNSName string
	// The target the DNS record points to
	Target string
	// Labels stores labels defined for the Endpoint
	Labels map[string]string
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName, target string) *Endpoint {
	return &Endpoint{
		DNSName: strings.TrimSuffix(dnsName, "."),
		Target:  strings.TrimSuffix(target, "."),
		Labels:  map[string]string{},
	}
}

func (e *Endpoint) MergeLabels(labels map[string]string) {
	for k, v := range labels {
		if e.Labels[k] == "" {
			e.Labels[k] = v
		}
	}
}
