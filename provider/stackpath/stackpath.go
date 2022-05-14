/*
Copyright 2020 The Kubernetes Authors.

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

package stackpath

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	"github.com/wmarchesi123/stackpath-go/pkg/dns"
)

type StackPathProvider struct {
	provider.BaseProvider
	client dns.APIClient
	dryRun bool
}

func (p *StackPathProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (p *StackPathProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}

func (p *StackPathProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return endpoints
}

func (p *StackPathProvider) PropertyValuesEqual(name, previous, current string) bool {
	return previous == current
}

func (p *StackPathProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return endpoint.DomainFilter{}
}
