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

package dnsprovider

var store = map[string]DNSProvider{}

// Register registers a DNSProvider under a given name.
func Register(name string, provider DNSProvider) {
	store[name] = provider
}

// Lookup returns a DNSProvider by the given name.
func Lookup(name string) DNSProvider {
	return store[name]
}

// LookupMultiple returns multiple DNSProviders given multiple names.
func LookupMultiple(names ...string) (providers []DNSProvider) {
	for _, name := range names {
		providers = append(providers, Lookup(name))
	}

	return providers
}
