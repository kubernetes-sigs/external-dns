/*
Copyright 2025 The Kubernetes Authors.

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

package factory

import (
	"fmt"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/registry/awssd"
	"sigs.k8s.io/external-dns/registry/dynamodb"
	"sigs.k8s.io/external-dns/registry/noop"
	"sigs.k8s.io/external-dns/registry/txt"
)

// RegistryConstructor is a function that creates a Registry from configuration and a provider.
type RegistryConstructor func(cfg *externaldns.Config, p provider.Provider) (registry.Registry, error)

// Select creates a registry based on the given configuration.
func Select(cfg *externaldns.Config, p provider.Provider) (registry.Registry, error) {
	constructor, ok := registries(cfg.Registry)
	if !ok {
		return nil, fmt.Errorf("unknown registry: %s", cfg.Registry)
	}
	return constructor(cfg, p)
}

// registries looks up the constructor for the named registry.
func registries(selector string) (RegistryConstructor, bool) {
	m := map[string]RegistryConstructor{
		externaldns.RegistryDynamoDB: dynamodb.New,
		externaldns.RegistryNoop:     noop.New,
		externaldns.RegistryTXT:      txt.New,
		externaldns.RegistryAWSSD:    awssd.New,
	}
	c, ok := m[selector]
	return c, ok
}
